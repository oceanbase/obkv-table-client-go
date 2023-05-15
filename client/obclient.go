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

package client

import (
	"context"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/config"
	oberror "github.com/oceanbase/obkv-table-client-go/error"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObClient struct {
	config       *config.ClientConfig
	configUrl    string
	fullUserName string
	userName     string
	tenantName   string
	clusterName  string
	password     string
	database     string
	sysUA        route.ObUserAuth

	ocpModel *route.ObOcpModel

	tableMutexes       sync.Map // map[tableName]sync.RWMutex
	tableLocations     sync.Map // map[tableName]*route.ObTableEntry
	tableRoster        sync.Map
	serverRoster       obServerRoster
	tableRowKeyElement map[string]*table.ObRowKeyElement

	lastRefreshMetadataTimestamp atomic.Int64
	refreshMetadataLock          sync.Mutex
}

func newObClient(
	configUrl string,
	fullUserName string,
	passWord string,
	sysUserName string,
	sysPassWord string,
	cliConfig *config.ClientConfig) (*ObClient, error) {
	cli := new(ObClient)
	// 1. Parse full username to get userName/tenantName/clusterName
	err := cli.parseFullUserName(fullUserName)
	if err != nil {
		return nil, errors.WithMessagef(err, "parse full user name, fullUserName:%s", fullUserName)
	}
	// 2. Parse config url to get database
	err = cli.parseConfigUrl(configUrl)
	if err != nil {
		return nil, errors.WithMessagef(err, "parse config url, configUrl:%s", configUrl)
	}

	// 3. init other members
	cli.password = passWord
	cli.sysUA = *route.NewObUserAuth(sysUserName, sysPassWord)
	cli.config = cliConfig
	cli.tableRowKeyElement = make(map[string]*table.ObRowKeyElement)

	return cli, nil
}

func (c *ObClient) String() string {
	return "ObClient{" +
		"configUrl:" + c.configUrl + ", " +
		"fullUserName:" + c.fullUserName + ", " +
		"userName:" + c.userName + ", " +
		"tenantName:" + c.tenantName + ", " +
		"clusterName:" + c.clusterName + ", " +
		"sysUA:" + c.sysUA.String() + ", " +
		"configUrl:" + c.configUrl + ", " +
		"configUrl:" + c.configUrl + ", " +
		"config:" + c.config.String() +
		"}"
}

// standard format: user_name@tenant_name#cluster_name
func (c *ObClient) parseFullUserName(fullUserName string) error {
	utIndex := strings.Index(fullUserName, "@")
	tcIndex := strings.Index(fullUserName, "#")
	if utIndex == -1 || tcIndex == -1 || tcIndex <= utIndex {
		return errors.Errorf("invalid full user name, fullUserName:%s", fullUserName)
	}
	userName := fullUserName[:utIndex]
	tenantName := fullUserName[utIndex+1 : tcIndex]
	clusterName := fullUserName[tcIndex+1:]
	if userName == "" || tenantName == "" || clusterName == "" {
		return errors.Errorf("invalid element in full user name, userName:%s, tenantName:%s, clusterName:%s",
			userName, tenantName, clusterName)
	}
	c.userName = userName
	c.tenantName = tenantName
	c.clusterName = clusterName
	c.fullUserName = fullUserName
	return nil
}

// format: http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx
func (c *ObClient) parseConfigUrl(configUrl string) error {
	index := strings.Index(configUrl, "database=")
	if index == -1 {
		index = strings.Index(configUrl, "DATABASE=")
		if index == -1 {
			return errors.Errorf("config url not contain database, configUrl:%s", configUrl)
		}
	}
	db := configUrl[index+len("database="):]
	if db == "" {
		return errors.Errorf("database is empty, configUrl:%s", configUrl)
	}
	c.configUrl = configUrl
	c.database = db
	return nil
}

func (c *ObClient) init() error {
	return c.fetchMetadata()
}

func (c *ObClient) AddRowKey(tableName string, rowKey []string) error {
	if tableName == "" || len(rowKey) == 0 {
		return errors.Errorf("nil table name or empty rowKey, tableName:%s, rowKey size:%d", tableName, len(rowKey))
	}
	m := make(map[string]int, 1)
	for i := 0; i < len(rowKey); i++ {
		columnName := rowKey[i]
		m[columnName] = i
	}
	c.tableRowKeyElement[tableName] = table.NewObRowKeyElement(m)
	return nil
}

func (c *ObClient) Insert(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...ObkvOption) (int64, error) {
	res, err := c.execute(
		ctx,
		tableName,
		protocol.Insert,
		rowKey,
		mutateColumns,
		opts...)
	if err != nil {
		return -1, errors.WithMessagef(err, "execute insert, tableName:%s, rowKey:%s, mutateColumns:%s",
			tableName, table.ColumnsToString(rowKey), table.ColumnsToString(mutateColumns))
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) Update(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...ObkvOption) (int64, error) {
	res, err := c.execute(
		ctx,
		tableName,
		protocol.Update,
		rowKey,
		mutateColumns,
		opts...)
	if err != nil {
		return -1, errors.WithMessagef(err, "execute update, tableName:%s, rowKey:%s, mutateColumns:%s",
			tableName, table.ColumnsToString(rowKey), table.ColumnsToString(mutateColumns))
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) InsertOrUpdate(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...ObkvOption) (int64, error) {
	res, err := c.execute(
		ctx,
		tableName,
		protocol.InsertOrUpdate,
		rowKey,
		mutateColumns,
		opts...)
	if err != nil {
		return -1, errors.WithMessagef(err, "execute insert or update, tableName:%s, rowKey:%s, mutateColumns:%s",
			tableName, table.ColumnsToString(rowKey), table.ColumnsToString(mutateColumns))
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) Delete(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	opts ...ObkvOption) (int64, error) {
	res, err := c.execute(
		ctx,
		tableName,
		protocol.Del,
		rowKey,
		nil,
		opts...)
	if err != nil {
		return -1, errors.WithMessagef(err, "execute delete, tableName:%s, rowKey:%s",
			tableName, table.ColumnsToString(rowKey))
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) Get(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	getColumns []string,
	opts ...ObkvOption) (map[string]interface{}, error) {
	var columns []*table.Column
	for _, columnName := range getColumns {
		columns = append(columns, table.NewColumn(columnName, nil))
	}
	res, err := c.execute(
		ctx,
		tableName,
		protocol.Get,
		rowKey,
		columns,
		opts...)
	if err != nil {
		return nil, errors.WithMessagef(err, "execute get, tableName:%s, rowKey:%s, getColumns:%s",
			tableName, table.ColumnsToString(rowKey), util.StringArrayToString(getColumns))
	}
	return res.Entity().GetSimpleProperties(), nil
}

func (c *ObClient) NewBatchExecutor(tableName string) BatchExecutor {
	return newObBatchExecutor(tableName, c)
}

func (c *ObClient) execute(
	ctx context.Context,
	tableName string,
	opType protocol.TableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	opts ...ObkvOption) (*protocol.TableOperationResponse, error) {
	var rowKeyValue []interface{}
	for _, col := range rowKey {
		rowKeyValue = append(rowKeyValue, col.Value())
	}
	// 1. Get table route
	tableParam, err := c.getTableParam(ctx, tableName, rowKeyValue, false /* refresh */)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table param, tableName:%s, opType:%d", tableName, opType)
	}

	// 2. Construct request.
	request, err := protocol.NewTableOperationRequest(
		tableName,
		tableParam.tableId,
		tableParam.partitionId,
		opType,
		rowKey,
		columns,
		c.config.OperationTimeOut,
		c.config.LogLevel,
	)
	if err != nil {
		return nil, errors.WithMessagef(err, "new operation request, tableName:%s, tableParam:%s, opType:%d",
			tableName, tableParam.String(), opType)
	}

	// 3. execute
	result := protocol.NewTableOperationResponse()
	err = tableParam.table.execute(request, result)
	if err != nil {
		return nil, errors.WithMessagef(err, "execute request, request:%s", request.String())
	}

	if oberror.ObErrorCode(result.Header().ErrorNo()) != oberror.ObSuccess {
		return nil, oberror.NewProtocolError(
			tableParam.table.ip,
			tableParam.table.port,
			oberror.ObErrorCode(result.Header().ErrorNo()),
			result.UniqueId(),
			result.Sequence(),
			tableName,
		)
	}

	return result, nil
}

func (c *ObClient) getTableParam(
	ctx context.Context,
	tableName string,
	rowKeyValue []interface{},
	refresh bool) (*ObTableParam, error) {
	entry, err := c.getOrRefreshTableEntry(ctx, tableName, refresh, false)
	if err != nil {
		return nil, errors.WithMessagef(err, "get or refresh table entry, tableName:%s", tableName)
	}
	partId, err := c.getPartitionId(entry, rowKeyValue)
	if err != nil {
		return nil, errors.WithMessagef(err, "get partition id, tableName:%s, tableEntry:%s", tableName, entry.String())
	}
	t, err := c.getTable(entry, partId)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table, tableName:%s, tableEntry:%s, partId:%d",
			tableName, entry.String(), partId)
	}

	if util.ObVersion() >= 4 {
		partId, err = entry.PartitionInfo().GetTabletId(partId)
		if err != nil {
			return nil, errors.WithMessagef(err, "get tablet id, tableName:%s, tableEntry:%s, partId:%d",
				tableName, entry.String(), partId)
		}
	}

	return NewObTableParam(t, entry.TableId(), partId), nil
}

func (c *ObClient) needRefreshTableEntry(entry *route.ObTableEntry) (int64, bool) {
	ratio := math.Pow(2, float64(c.serverRoster.MaxPriority()))
	intervalMs := float64(c.config.TableEntryRefreshIntervalBase) / ratio
	ceilingMs := float64(c.config.TableEntryRefreshIntervalCeiling)
	intervalMs = math.Min(intervalMs, ceilingMs)
	return int64(intervalMs) - (time.Now().UnixMilli() - entry.RefreshTimeMills()),
		float64(time.Now().UnixMilli()-entry.RefreshTimeMills()) >= intervalMs
}

func (c *ObClient) getOrRefreshTableEntry(
	ctx context.Context,
	tableName string,
	refresh bool,
	waitForRefresh bool) (*route.ObTableEntry, error) {
	// 1. Get entry from cache
	entry := c.getTableEntryFromCache(tableName)
	if entry != nil {
		if !refresh { // If the refresh is false indicates that user tolerate not the latest data
			return entry, nil
		}
		sleepTime, needRefresh := c.needRefreshTableEntry(entry)
		if needRefresh {
			if waitForRefresh {
				time.Sleep(time.Millisecond * time.Duration(sleepTime))
			} else {
				return entry, nil
			}
		}
	}

	// 2. Cache entry not exist, get from remote
	// 2.1 Lock table firstly
	var lock *sync.RWMutex
	tmpLock := new(sync.RWMutex)
	v, loaded := c.tableMutexes.LoadOrStore(tableName, tmpLock)
	if loaded {
		lock = v.(*sync.RWMutex)
	} else {
		lock = tmpLock
	}

	var timeout int64 = 0
	for ; timeout < c.config.TableEntryRefreshLockTimeout.Milliseconds() && !lock.TryLock(); timeout++ {
		time.Sleep(time.Millisecond)
	}
	if timeout == c.config.TableEntryRefreshLockTimeout.Milliseconds() {
		return nil, errors.Errorf("failed to try lock table to refresh, timeout:%d", timeout)
	}
	defer func() {
		lock.Unlock()
	}()

	// 2.2 Double check whether we need to do refreshing or not, other goroutine may have refreshed
	entry = c.getTableEntryFromCache(tableName)
	if entry != nil {
		if _, needRefresh := c.needRefreshTableEntry(entry); !needRefresh {
			return entry, nil
		}
	}

	if entry == nil || refresh {
		refreshTryTimes := int(math.Min(float64(c.serverRoster.Size()), float64(c.config.TableEntryRefreshTryTimes)))
		for i := 0; i < refreshTryTimes; i++ {
			err := c.refreshTableEntry(ctx, &entry, tableName)
			if err != nil {
				log.Warn("failed to refresh table entry",
					log.Int("times", i),
					log.String("tableName", tableName))
			} else {
				return entry, nil
			}
		}
		log.Info("refresh table entry has tried specific times failure and will sync refresh metadata",
			log.Int("refreshTryTimes", refreshTryTimes),
			log.String("tableName", tableName))
		err := c.syncRefreshMetadata()
		if err != nil {
			return nil, errors.WithMessagef(err, "sync refresh meta data, tableName:%s", tableName)
		}
		err = c.refreshTableEntry(ctx, &entry, tableName)
		if err != nil {
			return nil, errors.WithMessagef(err, "refresh table entry, tableName:%s", tableName)
		}
		return entry, nil
	}

	// entry != nil but entry maybe is expired
	return entry, nil
}

func (c *ObClient) getTableEntryFromCache(tableName string) *route.ObTableEntry {
	v, ok := c.tableLocations.Load(tableName)
	if ok {
		entry, _ := v.(*route.ObTableEntry)
		return entry
	}
	return nil
}

func (c *ObClient) refreshTableEntry(ctx context.Context, entry **route.ObTableEntry, tableName string) error {
	var err error
	// 1. Load table entry location or table entry.
	if *entry != nil { // If table entry exist we just need to refresh table locations
		err = c.loadTableEntryLocation(ctx, *entry)
		if err != nil {
			return errors.WithMessagef(err, "load table entry location, tableName:%s", tableName)
		}
	} else {
		key := route.NewObTableEntryKey(c.clusterName, c.tenantName, c.database, tableName)
		*entry, err = route.GetTableEntryFromRemote(ctx, c.serverRoster.GetServer(), &c.sysUA, key)
		if err != nil {
			return errors.WithMessagef(err, "get table entry from remote, key:%s", key.String())
		}
	}

	// 2. Set rowKey element to entry.
	if (*entry).IsPartitionTable() {
		rowKeyElement, ok := c.tableRowKeyElement[tableName]
		if !ok {
			return errors.Errorf("failed to get rowKey element by table name, tableName:%s", tableName)
		}
		(*entry).SetRowKeyElement(rowKeyElement)
	}

	// 3. todo:prepare the table entry for weak read.

	// 4. Put entry to cache.
	c.tableLocations.Store(tableName, entry)

	return nil
}

func (c *ObClient) loadTableEntryLocation(ctx context.Context, entry *route.ObTableEntry) error {
	addr := c.serverRoster.GetServer()
	// 1. Get db handle
	db, err := route.NewDB(
		c.sysUA.UserName(),
		c.sysUA.Password(),
		addr.Ip(),
		strconv.Itoa(addr.SvrPort()),
		route.OceanbaseDatabase,
	)
	if err != nil {
		return errors.WithMessagef(err, "new db, sysUA:%s, addr:%s", c.sysUA.String(), addr.String())
	}
	defer func() {
		_ = db.Close()
	}()

	locEntry, err := route.GetPartLocationEntryFromRemote(ctx, db, entry)
	if err != nil {
		return errors.WithMessagef(err, "get part location entry from remote, tableEntry:%s", entry.String())
	}
	entry.SetPartLocationEntry(locEntry)
	return nil
}

func (c *ObClient) isMetaAlreadyRefreshed() bool {
	return time.Now().UnixMilli()-c.lastRefreshMetadataTimestamp.Load() < c.config.MetadataRefreshInterval.Milliseconds()
}

func (c *ObClient) syncRefreshMetadata() error {
	// 1. Check whether the meta has been refreshed or not
	if c.isMetaAlreadyRefreshed() {
		log.Info("try to lock metadata refreshing, it has refresh",
			log.Int64("lastRefreshTime", c.lastRefreshMetadataTimestamp.Load()),
			log.Duration("metadataRefreshInterval", c.config.MetadataRefreshInterval))
		return nil
	}

	// 2. try lock
	var timeout int64 = 0
	for ; timeout < c.config.MetadataRefreshLockTimeout.Milliseconds() && !c.refreshMetadataLock.TryLock(); timeout++ {
		time.Sleep(time.Millisecond)
	}
	if timeout == c.config.MetadataRefreshLockTimeout.Milliseconds() {
		return errors.Errorf("failed to lock metadata refreshing timeout, timeout:%d", timeout)
	}
	defer func() {
		c.refreshMetadataLock.Unlock()
	}()

	// 3. Double check after lock
	if c.isMetaAlreadyRefreshed() {
		log.Info("try to lock metadata refreshing, it has refresh",
			log.Int64("lastRefreshTime", c.lastRefreshMetadataTimestamp.Load()),
			log.Duration("metadataRefreshInterval", c.config.MetadataRefreshInterval))
		return nil
	}

	// 4. fetch meta data
	err := c.fetchMetadata()
	if err != nil {
		return errors.WithMessagef(err, "fetch meta data")
	}
	return nil
}

func (c *ObClient) fetchMetadata() error {
	// 1. Load ocp mode to get RsList
	ocpModel, err := route.LoadOcpModel(
		c.configUrl,
		c.config.RslistLocalFileLocation,
		c.config.RslistHttpGetTimeout,
		c.config.RslistHttpGetRetryTimes,
		c.config.RslistHttpGetRetryInterval,
	)
	if err != nil {
		return errors.WithMessagef(err, "load ocp model, configUrl:%s, localFileName:%s",
			c.configUrl, c.config.RslistLocalFileLocation)
	}
	c.ocpModel = ocpModel
	addr := c.ocpModel.GetServerAddressRandomly()

	// 2. Get ob cluster version and init route sql
	ver, err := route.GetObVersionFromRemote(addr, &c.sysUA)
	if err != nil {
		return errors.WithMessagef(err, "get ob version from remote, addr:%s, sysUA:%s",
			addr.String(), c.sysUA.String())
	}
	// 2.1 Set ob version and init route sql by ob version.
	if util.ObVersion() == 0.0 {
		util.SetObVersion(ver)
		route.InitSql(ver)
	}

	// 3. Dummy route to get tenant server and create ObTable for each observer node.
	//    Each ObTable contains a connection pool.
	// 3.1 Get table entry with specific table name("__all_dummy")
	rootServerKey := route.NewObTableEntryKey(
		c.clusterName,
		c.tenantName,
		route.OceanbaseDatabase,
		route.AllDummyTable,
	)
	entry, err := route.GetTableEntryFromRemote(context.TODO(), addr, &c.sysUA, rootServerKey)
	if err != nil {
		return errors.WithMessagef(err, "dummy tenant server from remote, addr:%s, sysUA:%s, key:%s",
			addr.String(), c.sysUA.String(), rootServerKey.String())
	}
	// 3.2 Save all tenant server address
	replicaLocations := entry.TableLocation().ReplicaLocations()
	servers := make([]*route.ObServerAddr, 0, len(replicaLocations))
	for _, replicaLoc := range replicaLocations {
		svrStatus := replicaLoc.SvrStatus()
		addr := replicaLoc.Addr()
		if !svrStatus.IsActive() {
			log.Warn("server is not active",
				log.String("server info", svrStatus.String()),
				log.String("server addr", addr.String()))
			continue
		}
		servers = append(servers, addr)

		if _, ok := c.tableRoster.Load(*addr); ok { // already exist, continue
			continue
		}

		t := NewObTable(addr.Ip(), addr.SvrPort(), c.tenantName, c.userName, c.password, c.database)
		err = t.init(c.config.ConnPoolMaxConnSize, c.config.RpcConnectTimeOut)
		if err != nil {
			return errors.WithMessagef(err, "init ob table, obTable:%s", t.String())
		}
		_, loaded := c.tableRoster.LoadOrStore(*addr, t)
		if loaded { // Already stored by other goroutine, close table
			t.close()
		}
	}

	// 3.3 Clean useless ob table
	c.tableRoster.Range(func(k, v interface{}) bool {
		contain := false
		for _, addr := range servers {
			if *addr == k {
				contain = true
				break
			}
		}
		if !contain {
			v, loaded := c.tableRoster.LoadAndDelete(k)
			if loaded {
				t := v.(*ObTable)
				t.close()
			}
		}
		return true
	})

	// 3.4 Update server roster
	c.serverRoster.Reset(servers)

	// 4. todo: Get Server LDC info for weak read consistency.
	// 5. Update lastRefreshMetadataTimestamp
	c.lastRefreshMetadataTimestamp.Store(time.Now().UnixMilli())
	return nil
}

// get partition id by rowKey
func (c *ObClient) getPartitionId(entry *route.ObTableEntry, rowKeyValue []interface{}) (int64, error) {
	if !entry.IsPartitionTable() || entry.PartitionInfo().Level() == route.PartLevelZero {
		return 0, nil
	}
	if entry.PartitionInfo().Level() == route.PartLevelOne {
		return entry.PartitionInfo().FirstPartDesc().GetPartId(rowKeyValue)
	}
	if entry.PartitionInfo().Level() == route.PartLevelTwo {
		partId1, err := entry.PartitionInfo().FirstPartDesc().GetPartId(rowKeyValue)
		if err != nil {
			return -1, errors.WithMessagef(err, "get part id from first part desc, firstDesc:%s",
				entry.PartitionInfo().FirstPartDesc().String())
		}
		partId2, err := entry.PartitionInfo().SubPartDesc().GetPartId(rowKeyValue)
		if err != nil {
			return -1, errors.WithMessagef(err, "get part id from sub part desc, firstDesc:%s",
				entry.PartitionInfo().SubPartDesc().String())
		}
		return (partId1)<<route.ObPartIdShift | partId2 | route.ObMask, nil
	}
	return -1, errors.Errorf("unknown partition level, partInfo:%s", entry.PartitionInfo().String())
}

func (c *ObClient) getTable(entry *route.ObTableEntry, partId int64) (*ObTable, error) {
	// 1. Get replica location by partition id
	replicaLoc, err := entry.GetPartitionReplicaLocation(partId, route.ConsistencyStrong)
	if err != nil {
		return nil, errors.WithMessagef(err, "get partition replica location, partId:%d", partId)
	}

	// 2. Get table from table Roster by server addr
	t, ok := c.tableRoster.Load(*replicaLoc.Addr())
	if !ok {
		return nil, errors.Errorf("failed to get table by server addr, addr:%s", replicaLoc.Addr().String())
	}
	// todo: check server addr is expired or not
	tb, _ := t.(*ObTable)
	return tb, nil
}
