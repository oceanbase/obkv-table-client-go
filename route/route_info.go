/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package route

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

const refreshTableTimeout = 5 * time.Second
const rslistCheckInterval = 5 * time.Minute

type ObRouteInfo struct {
	clusterVersion   float32
	clusterName      string
	sysUA            *ObUserAuth
	configServerInfo *ObConfigServerInfo
	tableMutexes     ObTableMutexes
	tableLocations   sync.Map // map[tableName]*route.ObTableEntry
	tableRoster      ObTableRoster
	serverRoster     ObServerRoster // all servers which contain current tenant
	taskInfo         *ObRouteTaskInfo
}

// GetTable get table by partition id
func (i *ObRouteInfo) getTableWithRetry(ctx context.Context, server *ObServerAddr) (*ObTable, error) {
	// Get table from table Roster by server addr
	t, ok := i.tableRoster.Get(server.tcpAddr)
	if !ok {
		isCreating := &atomic.Bool{}
		isCreating.Store(false)
		i.taskInfo.createConnPoolServers.AddIfAbsent(server.tcpAddr, isCreating)
		i.taskInfo.TriggerCreateConnectionPool()
		for {
			select {
			case <-ctx.Done():
				return nil, errors.Errorf("get table, server:%s", server.String())
			default:
				t, ok = i.tableRoster.Get(server.tcpAddr)
				if ok {
					return t, nil
				}
				time.Sleep(1 * time.Millisecond)
			}
		}
	}

	return t, nil
}

// GetTable get table by partition id
func (i *ObRouteInfo) GetTable(ctx context.Context, entry *ObTableEntry, partId uint64) (*ObTable, error) {
	// Get replica location by partition id
	replicaLoc, err := entry.GetPartitionLocation(partId, ConsistencyStrong)
	if err != nil {
		return nil, errors.WithMessagef(err, "get partition replica location, partId:%d", partId)
	}

	return i.getTableWithRetry(ctx, replicaLoc.Addr())
}

func (i *ObRouteInfo) ClusterVersion() float32 {
	return i.clusterVersion
}

func NewRouteInfo(sysUA *ObUserAuth) *ObRouteInfo {
	return &ObRouteInfo{
		clusterVersion:   0.0,
		sysUA:            sysUA,
		configServerInfo: NewConfigServerInfo(),
		taskInfo:         NewRouteTaskInfo(),
	}
}

// FetchConfigServerInfo get rslist from config server by http get
func (i *ObRouteInfo) FetchConfigServerInfo(
	configUrl string,
	fileName string,
	timeout time.Duration,
	retryTimes int,
	retryInternal time.Duration) error {
	// 1. init parameter
	i.configServerInfo.configUrl = configUrl
	i.configServerInfo.file = fileName
	i.configServerInfo.timeout = timeout
	i.configServerInfo.retryTimes = retryTimes
	i.configServerInfo.retryInterval = retryInternal

	// 2. get rslist
	rslist, err := i.configServerInfo.FetchRslist()
	if err != nil {
		return errors.WithMessagef(err, "fetch rslist")
	}
	i.configServerInfo.rslist = rslist
	return nil
}

// CheckClusterAndTenant get ob cluster version and check tenant exist
func (i *ObRouteInfo) CheckClusterAndTenant(tenantName string) error {
	// 1. fetch cluster version
	addr, err := i.configServerInfo.GetServerAddressRandomly()
	if err != nil {
		return err
	}
	db, err := NewDB(
		i.sysUA.userName,
		i.sysUA.password,
		addr.ip,
		strconv.Itoa(addr.sqlPort),
		OceanBaseDatabase,
	)
	if err != nil {
		return errors.WithMessagef(err, "new db, sysUA:%s, addr:%s", i.sysUA.String(), addr.String())
	}
	defer func() {
		_ = db.Close()
	}()

	ver, err := GetObVersionFromRemote(db)
	if err != nil {
		return errors.WithMessagef(err, "get cluster version, sysUA:%s, addr:%s", i.sysUA.String(), addr.String())
	}

	i.clusterVersion = ver

	// 2. check tenant exist
	err = CheckTenantExist(db, tenantName)
	if err != nil {
		return err
	}

	return nil
}

// FetchServerRoster get all servers which contains tenant
func (i *ObRouteInfo) FetchServerRoster(clusterName, tenantName string) error {
	i.clusterName = clusterName
	i.tableRoster.tenantName = tenantName
	key := NewObTableEntryKey(
		clusterName,
		tenantName,
		OceanBaseDatabase,
		AllDummyTable,
	)
	addr, err := i.configServerInfo.GetServerAddressRandomly()
	if err != nil {
		return err
	}
	entry, err := GetTableEntryFromRemote(context.TODO(), addr, i.sysUA, key)
	if err != nil {
		return errors.WithMessagef(err, "dummy tenant server from remote, addr:%s, sysUA:%s, key:%s",
			addr.String(), i.sysUA.String(), key.String())
	}

	replicaLocations := entry.TableLocation().ReplicaLocations()
	for _, replicaLoc := range replicaLocations {
		if !replicaLoc.SvrStatus().IsActive() {
			log.Warn("server is not active",
				log.String("server info", replicaLoc.SvrStatus().String()),
				log.String("server addr", addr.String()))
			continue
		}
		i.serverRoster.servers = append(i.serverRoster.servers, replicaLoc.addr)
	}

	return nil
}

// ConstructTableRoster tableRoster mean all partition table which contain a connection pool
func (i *ObRouteInfo) ConstructTableRoster(
	userName string,
	password string,
	database string,
	connPoolSize int,
	connectTimeout time.Duration,
	loginTimeout time.Duration) error {
	// 1. init login info
	i.tableRoster.userName = userName
	i.tableRoster.password = password
	i.tableRoster.database = database
	i.tableRoster.connPoolSize = connPoolSize
	i.tableRoster.connectTimeout = connectTimeout
	i.tableRoster.loginTimeout = loginTimeout
	tenantName := i.tableRoster.tenantName

	// 2. init table for each server and add to tableRoster
	for _, server := range i.serverRoster.servers {
		table := NewObTable(server.Ip(), server.SvrPort(), tenantName, userName, password, database)
		err := table.Init(connPoolSize, connectTimeout, loginTimeout)
		if err != nil {
			return errors.WithMessagef(err, "init ob table, obTable:%s", table.String())
		}
		i.tableRoster.Add(tcpAddr{ip: server.Ip(), port: server.SvrPort()}, table)
	}
	return nil
}

// Close all connection pool
func (i *ObRouteInfo) Close() {
	i.tableRoster.Close()
	i.taskInfo.checkRslistTicker.Stop()
}

// get partition id by rowKey
func (i *ObRouteInfo) getPartitionId(entry *ObTableEntry, rowKeyValue []*table.Column) (uint64, error) {
	if !entry.IsPartitionTable() || entry.PartitionInfo().Level() == PartLevelZero {
		return 0, nil
	}
	if entry.PartitionInfo().Level() == PartLevelOne {
		return entry.PartitionInfo().FirstPartDesc().GetPartId(rowKeyValue)
	}
	if entry.PartitionInfo().Level() == PartLevelTwo {
		partId1, err := entry.PartitionInfo().FirstPartDesc().GetPartId(rowKeyValue)
		if err != nil {
			return ObInvalidPartId, errors.WithMessagef(err, "get part id from first part desc, firstDesc:%s",
				entry.PartitionInfo().FirstPartDesc().String())
		}
		partId2, err := entry.PartitionInfo().SubPartDesc().GetPartId(rowKeyValue)
		if err != nil {
			return ObInvalidPartId, errors.WithMessagef(err, "get part id from sub part desc, firstDesc:%s",
				entry.PartitionInfo().SubPartDesc().String())
		}
		return (partId1)<<ObPartIdShift | partId2 | ObMask, nil
	}
	return ObInvalidPartId, errors.Errorf("unknown partition level, partInfo:%s", entry.PartitionInfo().String())
}

// GetPartitionIds get partition ids by rowKeyPair
func (i *ObRouteInfo) GetPartitionIds(entry *ObTableEntry, rowKeyPair *table.RangePair) ([]uint64, error) {
	if !entry.IsPartitionTable() || entry.PartitionInfo().Level() == PartLevelZero {
		return []uint64{0}, nil
	}
	if entry.PartitionInfo().Level() == PartLevelOne {
		return entry.PartitionInfo().FirstPartDesc().GetPartIds(rowKeyPair)
	}
	if entry.PartitionInfo().Level() == PartLevelTwo {
		partIds1, err := entry.PartitionInfo().FirstPartDesc().GetPartIds(rowKeyPair)
		if err != nil {
			return nil, errors.WithMessagef(err, "get part id from first part desc, firstDesc:%s",
				entry.PartitionInfo().FirstPartDesc().String())
		}
		partIds2, err := entry.PartitionInfo().SubPartDesc().GetPartIds(rowKeyPair)
		if err != nil {
			return nil, errors.WithMessagef(err, "get part id from sub part desc, firstDesc:%s",
				entry.PartitionInfo().SubPartDesc().String())
		}
		// do cartesian product
		partIds := make([]uint64, 0, len(partIds1)*len(partIds2))
		for _, partId1 := range partIds1 {
			for _, partId2 := range partIds2 {
				partIds = append(partIds, (partId1)<<ObPartIdShift|partId2|ObMask)
			}
		}
		return partIds, nil
	}
	return nil, errors.Errorf("unknown partition level, partInfo:%s", entry.PartitionInfo().String())
}

// GetTableParam get table param by rowkey, ObTableParam means a partition table
func (i *ObRouteInfo) GetTableParam(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	opdTable *ObTable) (*ObTableParam, error) {
	// odp table
	if opdTable != nil {
		return NewObTableParam(opdTable, 0, 0), nil
	}
	entry, err := i.GetTableEntry(ctx, tableName)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table entry, tableName:%s", tableName)
	}
	partId, err := i.getPartitionId(entry, rowKey)
	if err != nil {
		return nil, errors.WithMessagef(err, "get partition id, tableName:%s, tableEntry:%s", tableName, entry.String())
	}
	t, err := i.GetTable(ctx, entry, partId)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table, tableName:%s, tableEntry:%s, partId:%d",
			tableName, entry.String(), partId)
	}

	if util.ObVersion() >= 4 && entry.IsPartitionTable() {
		partId, err = entry.PartitionInfo().GetTabletId(partId)
		if err != nil {
			return nil, errors.WithMessagef(err, "get tablet id, tableName:%s, tableEntry:%s, partId:%d",
				tableName, entry.String(), partId)
		}
	}

	return NewObTableParam(t, entry.TableId(), partId), nil
}

func (i *ObRouteInfo) getTableEntryFromCache(tableName string) *ObTableEntry {
	v, ok := i.tableLocations.Load(tableName)
	if ok {
		entry, _ := v.(*ObTableEntry)
		return entry
	}
	return nil
}

// GetTableEntry get table entry from cache or remote
func (i *ObRouteInfo) GetTableEntry(ctx context.Context, tableName string) (*ObTableEntry, error) {
	// 1. Get entry from cache
	entry := i.getTableEntryFromCache(tableName)
	if entry != nil {
		return entry, nil
	}

	// 2. Cache entry not exist, get from remote
	// 2.1 Lock table firstly
	i.tableMutexes.Lock(tableName)
	defer i.tableMutexes.Unlock(tableName)

	// 2.2 Double check whether we need to do refreshing or not, other goroutine may have fetch
	entry = i.getTableEntryFromCache(tableName)
	if entry != nil {
		return entry, nil
	}

	// 2.3 Fetch table entry
	key := NewObTableEntryKey(
		i.clusterName,
		i.tableRoster.tenantName,
		i.tableRoster.database,
		tableName,
	)
	entry, err := GetTableEntryFromRemote(ctx, i.serverRoster.GetServer(), i.sysUA, key)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table entry from remote, key:%s", key.String())
	}

	// 3. Store cache
	i.tableLocations.Store(tableName, entry)

	return entry, nil
}

func (i *ObRouteInfo) refreshTableLocations(addr *tcpAddr) error {
	i.tableLocations.Range(func(key, value interface{}) bool {
		tableName := key.(string)
		entry := value.(*ObTableEntry)
		for _, replica := range entry.tableLocation.replicaLocations {
			if replica.addr.tcpAddr.Equal(addr) {
				// trigger refresh table
				isRefreshing := &atomic.Bool{}
				isRefreshing.Store(false)
				i.taskInfo.tables.AddIfAbsent(tableName, isRefreshing)
				i.taskInfo.TriggerRefreshTable()
			}
		}
		return true
	})
	return nil
}

func (i *ObRouteInfo) RunBackgroundTask() {
	for {
		select {
		case <-i.taskInfo.createConnPoolChan:
			i.runCreateConnPoolTask()
		case <-i.taskInfo.dropConnPoolChan:
			i.runDropConnPoolTask()
		case <-i.taskInfo.refreshTableChan:
			i.runRefreshTableTask()
		case <-i.taskInfo.configServerChan:
			i.runCheckRslistTask()
		}
	}
}

func (i *ObRouteInfo) RunTickerTask() {
	for {
		select {
		case <-i.taskInfo.checkRslistTicker.C:
			i.taskInfo.TriggerCheckRslist()
		}
	}
}

func (i *ObRouteInfo) reroute(
	ctx context.Context,
	moveRsp *protocol.ObTableMoveResponse,
	request protocol.ObPayload,
	result protocol.ObPayload) error {

	// 1. Get table
	addr := moveRsp.ReplicaInfo().Server()
	port := moveRsp.ReplicaInfo().Server().Port()
	server := NewObServerAddr(
		addr.IpToString(),
		0,
		int(port),
	)
	table, err := i.getTableWithRetry(ctx, server)
	if err != nil {
		return errors.WithMessagef(err, "get table, server:%s", server.String())
	}

	// 2. Execute
	_, err = table.Execute(ctx, request, result)
	return err
}

func (i *ObRouteInfo) Execute(
	ctx context.Context,
	tableName string,
	table *ObTable,
	request protocol.ObPayload,
	result protocol.ObPayload) (error, bool) {

	needRetry := false
	needRefreshTable := false
	needReroute := false
	moveRsp, err := table.Execute(ctx, request, result)
	if err != nil {
		if table.IsDisconnected() {
			isDropping := &atomic.Bool{}
			isDropping.Store(false)
			server := tcpAddr{table.ip, table.port}
			i.taskInfo.dropConnPoolServers.AddIfAbsent(server, isDropping)
			i.taskInfo.TriggerDropConnectionPool()
		}

		if result.Flag()&protocol.RpcBadRoutingFlag != 0 {
			needRefreshTable = true
		}

		if moveRsp != nil {
			needReroute = true
		}
	}

	if needRefreshTable {
		// add table and trigger background refresh table task
		isRefreshing := &atomic.Bool{}
		isRefreshing.Store(false)
		i.taskInfo.tables.AddIfAbsent(tableName, isRefreshing)
		i.taskInfo.TriggerRefreshTable()
		needRetry = true
	}

	if needReroute {
		log.Info(fmt.Sprintf("route, to:%s", moveRsp.ReplicaInfo().Server().String()))
		err = i.reroute(ctx, moveRsp, request, result)
	}

	if err != nil {
		if result.RemoteAddr() != nil {
			return errors.WithMessagef(err, "obtable remote:[%s] execute", result.RemoteAddr().String()), needRetry
		} else {
			return errors.WithMessagef(err, "obtable (remote unknown) execute"), needRetry
		}
	}

	return nil, false
}

func (i *ObRouteInfo) addTable(addr tcpAddr) error {
	_, ok := i.tableRoster.Get(addr)
	if ok {
		return nil // exist, no need to add
	}

	table := NewObTable(
		addr.ip,
		addr.port,
		i.tableRoster.tenantName,
		i.tableRoster.userName,
		i.tableRoster.password,
		i.tableRoster.database,
	)
	err := table.Init(i.tableRoster.connPoolSize, i.tableRoster.connectTimeout, i.tableRoster.loginTimeout)
	if err != nil {
		return errors.WithMessagef(err, "init ob table, obTable:%s", table.String())
	}
	i.tableRoster.Add(addr, table)

	return nil
}

func (i *ObRouteInfo) dropTable(addr tcpAddr) {
	v, load := i.tableRoster.LoadAndDelete(addr)
	if load {
		table, ok := v.(*ObTable)
		if ok {
			table.Close()
		}
	}
}

func (i *ObRouteInfo) runCreateConnPoolTask() {
	for idx := 0; idx < i.taskInfo.createConnPoolServers.Size(); idx++ {
		i.taskInfo.createConnPoolServers.Range(func(key, value interface{}) {
			isCreating := value.(*atomic.Bool)
			swapped := isCreating.CompareAndSwap(false, true)
			if swapped {
				go func() {
					addr := key.(tcpAddr)
					err := i.addTable(addr)
					if err != nil {
						log.Warn("add table", log.String("server", addr.String()))
					}
					i.taskInfo.createConnPoolServers.Remove(key)
					log.Info("[runCreateConnPoolTask] connection pool for server has been created", log.String("server", addr.String()))
				}()
			}
		})
	}
}

func (i *ObRouteInfo) runDropConnPoolTask() {
	for idx := 0; idx < i.taskInfo.dropConnPoolServers.Size(); idx++ {
		i.taskInfo.dropConnPoolServers.Range(func(key, value interface{}) {
			isDropping := value.(*atomic.Bool)
			swapped := isDropping.CompareAndSwap(false, true)
			if swapped {
				go func() {
					// 1. drop old table
					addr := key.(tcpAddr)
					i.dropTable(addr)
					i.taskInfo.dropConnPoolServers.Remove(key)
					log.Info("[runDropConnPoolTask] connection pool has been dropped", log.String("addr", addr.String()))

					// 2. refresh table locations which contain the dropping server
					err := i.refreshTableLocations(&addr)
					if err != nil {
						log.Warn("refresh table locations", log.String("server", addr.String()))
					}
					log.Info("[runDropConnPoolTask] table contains server pool has been refreshed", log.String("server", addr.String()))
				}()
			}
		})
	}
}

func (i *ObRouteInfo) refreshTableEntry(tableName string) error {
	// 1. Check need refresh or not
	value, ok := i.tableLocations.Load(tableName)
	if ok {
		entry, ok := value.(*ObTableEntry)
		if ok && !entry.NeedRefresh() { // no need to refresh
			log.Info("no need to refresh table", log.String("table", tableName))
			return nil
		}
	}

	// 2. Delete from cache
	i.tableLocations.Delete(tableName)

	// 3. Refresh
	ctx, _ := context.WithTimeout(context.Background(), refreshTableTimeout)
	_, err := i.GetTableEntry(ctx, tableName)
	if err != nil {
		log.Warn("get table entry", log.String("tableName", tableName))
		return err
	}

	return nil
}

func (i *ObRouteInfo) runRefreshTableTask() {
	for idx := 0; idx < i.taskInfo.tables.Size(); idx++ {
		i.taskInfo.tables.Range(func(key, value interface{}) {
			isRefreshing := value.(*atomic.Bool)
			swapped := isRefreshing.CompareAndSwap(false, true)
			if swapped {
				go func() {
					tableName := key.(string)
					log.Info("[runRefreshTableTask] refresh table entry", log.String("table", tableName))
					err := i.refreshTableEntry(tableName)
					if err != nil {
						log.Warn("refresh table entry", log.String("tableName", tableName))
					}
					i.taskInfo.tables.Remove(key)
				}()
			}
		})
	}
}

func (i *ObRouteInfo) runCheckRslistTask() {
	go func() {
		newRslist, err := i.configServerInfo.FetchRslist()
		if err != nil {
			log.Warn("fetch rslist")
		}
		if !i.configServerInfo.rslist.Equal(newRslist) {
			missServers := i.configServerInfo.rslist.FindMissingElements(newRslist)
			log.Info(fmt.Sprintf("[runCheckRslistTask] missServers size:%d", len(missServers)))
			for _, server := range missServers {
				log.Info(fmt.Sprintf("[runCheckRslistTask] missServer:%s", server.tcpAddr.String()))
				t, ok := i.tableRoster.Get(server.tcpAddr)
				if ok {
					if t.IsDisconnected() {
						log.Info(fmt.Sprintf("[runCheckRslistTask] connection pool need close, pool:%s", t.String()))
						isDropping := &atomic.Bool{}
						isDropping.Store(false)
						server := tcpAddr{t.ip, t.port}
						i.taskInfo.dropConnPoolServers.AddIfAbsent(server, isDropping)
						i.taskInfo.TriggerDropConnectionPool()
					}
				}
			}
			i.configServerInfo.rslist = newRslist
		}
	}()
}
