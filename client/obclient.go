package client

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/config"
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
	serverRoster       route.ObServerRoster
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
		log.Warn("failed to parse full user name", log.String("fullUserName", fullUserName))
		return nil, err
	}
	// 2. Parse config url to get database
	err = cli.parseConfigUrl(configUrl)
	if err != nil {
		log.Warn("failed to parse config url", log.String("configUrl", configUrl))
		return nil, err
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
		log.Warn("invalid full user name", log.String("fullUserName", fullUserName))
		return errors.New("invalid full user name")
	}
	userName := fullUserName[:utIndex]
	tenantName := fullUserName[utIndex+1 : tcIndex]
	clusterName := fullUserName[tcIndex+1:]
	if userName == "" || tenantName == "" || clusterName == "" {
		log.Warn("invalid element in full user name",
			log.String("userName", userName),
			log.String("tenantName", tenantName),
			log.String("clusterName", clusterName))
		return errors.New("invalid element in full user name")
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
			log.Warn("config url not contain database", log.String("config url", configUrl))
			return errors.New("config url not contain database")
		}
	}
	db := configUrl[index+len("database="):]
	if db == "" {
		log.Warn("database is empty", log.String("config url", configUrl))
		return errors.New("database is empty")
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
		log.Warn("nil table name or empty rowKey",
			log.String("tableName", tableName),
			log.Int("rowKey size", len(rowKey)))
		return errors.New("nil table name or empty rowKey")
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
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...ObkvOption) (int64, error) {
	var mutateColNames []string
	var mutateColValues []interface{}
	for _, col := range mutateColumns {
		mutateColNames = append(mutateColNames, col.Name())
		mutateColValues = append(mutateColValues, col.Value())
	}
	res, err := c.execute(
		tableName,
		protocol.Insert,
		rowKey,
		mutateColNames,
		mutateColValues,
		opts...)
	if err != nil {
		log.Warn("failed to execute insert",
			log.String("tableName", tableName),
			log.String("rowkey", table.ColumnsToString(rowKey)),
			log.String("mutateColumns", table.ColumnsToString(mutateColumns)))
		return -1, err
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) InsertOrUpdate(
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...ObkvOption) (int64, error) {
	var mutateColNames []string
	var mutateColValues []interface{}
	for _, col := range mutateColumns {
		mutateColNames = append(mutateColNames, col.Name())
		mutateColValues = append(mutateColValues, col.Value())
	}
	res, err := c.execute(
		tableName,
		protocol.InsertOrUpdate,
		rowKey,
		mutateColNames,
		mutateColValues,
		opts...)
	if err != nil {
		log.Warn("failed to execute insertOrUpdate",
			log.String("tableName", tableName),
			log.String("rowkey", table.ColumnsToString(rowKey)),
			log.String("mutateColumns", table.ColumnsToString(mutateColumns)))
		return -1, err
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) Delete(
	tableName string,
	rowKey []*table.Column,
	opts ...ObkvOption) (int64, error) {
	res, err := c.execute(
		tableName,
		protocol.Del,
		rowKey,
		nil,
		nil,
		opts...)
	if err != nil {
		log.Warn("failed to execute del",
			log.String("tableName", tableName),
			log.String("rowkey", table.ColumnsToString(rowKey)))
		return -1, err
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) Get(
	tableName string,
	rowKey []*table.Column,
	getColumns []string,
	opts ...ObkvOption) (map[string]interface{}, error) {
	res, err := c.execute(
		tableName,
		protocol.Get,
		rowKey,
		getColumns,
		nil,
		opts...)
	if err != nil {
		log.Warn("failed to execute get",
			log.String("tableName", tableName),
			log.String("rowkey", table.ColumnsToString(rowKey)),
			log.String("getColumns", util.StringArrayToString(getColumns)))
		return nil, err
	}
	return res.Entity().GetSimpleProperties(), nil
}

func (c *ObClient) NewBatchExecutor(tableName string) BatchExecutor {
	return newObBatchExecutor(tableName, c)
}

func (c *ObClient) execute(
	tableName string,
	opType protocol.TableOperationType,
	rowKey []*table.Column,
	columns []string,
	properties []interface{},
	opts ...ObkvOption) (*protocol.TableOperationResponse, error) {
	var rowKeyValue []interface{}
	for _, col := range rowKey {
		rowKeyValue = append(rowKeyValue, col.Value())
	}
	// 1. Get table route
	tableParam, err := c.getTableParam(tableName, rowKeyValue, false /* refresh */)
	if err != nil {
		log.Warn("failed to get table param",
			log.String("tableName", tableName),
			log.Int8("opType", int8(opType)))
		return nil, err
	}

	// 2. Construct request.
	request, err := protocol.NewTableOperationRequest(
		tableName,
		tableParam.tableId,
		tableParam.partitionId,
		opType,
		rowKeyValue,
		columns,
		properties,
		c.config.OperationTimeOut,
		c.config.LogLevel,
	)
	if err != nil {
		log.Warn("failed to new operation request",
			log.String("tableName", tableName),
			log.String("tableParam", tableParam.String()),
			log.Int8("opType", int8(opType)))
		return nil, err
	}

	// 3. execute
	result := protocol.NewTableOperationResponse()
	err = tableParam.table.execute(request, result)
	if err != nil {
		log.Warn("failed to execute request", log.String("request", request.String()))
		return nil, errors.WithMessagef(err, "[%s]", request.String())
	}
	return result, nil
}

func (c *ObClient) getTableParam(
	tableName string,
	rowKeyValue []interface{},
	refresh bool) (*ObTableParam, error) {
	entry, err := c.getOrRefreshTableEntry(tableName, refresh, false)
	if err != nil {
		log.Warn("failed to get or refresh table entry", log.String("tableName", tableName))
		return nil, err
	}
	partId, err := c.getPartitionId(entry, rowKeyValue)
	if err != nil {
		log.Warn("failed to get partition id",
			log.String("tableName", tableName),
			log.String("entry", entry.String()))
		return nil, err
	}
	t, err := c.getTable(tableName, entry, partId)
	if err != nil {
		log.Warn("failed to get table",
			log.String("tableName", tableName),
			log.String("entry", entry.String()),
			log.Int64("partId", partId))
		return nil, err
	}

	if util.ObVersion() >= 4 {
		partId, err = entry.PartitionInfo().GetTabletId(partId)
		if err != nil {
			log.Warn("failed to get tablet id",
				log.String("tableName", tableName),
				log.String("entry", entry.String()),
				log.Int64("partId", partId))
			return nil, err
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

func (c *ObClient) getOrRefreshTableEntry(tableName string, refresh bool, waitForRefresh bool) (*route.ObTableEntry, error) {
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
		log.Warn("failed to try lock table to refresh", log.Int64("timeout", timeout))
		return nil, errors.New("failed to try lock table to refresh")
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
			err := c.refreshTableEntry(&entry, tableName)
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
			log.Warn("failed to sync refresh meta data", log.String("tableName", tableName))
			return nil, err
		}
		err = c.refreshTableEntry(&entry, tableName)
		if err != nil {
			log.Warn("failed to refresh table entry", log.String("tableName", tableName))
			return nil, err
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

func (c *ObClient) refreshTableEntry(entry **route.ObTableEntry, tableName string) error {
	var err error
	// 1. Load table entry location or table entry.
	if *entry != nil { // If table entry exist we just need to refresh table locations
		err = c.loadTableEntryLocation(*entry)
		if err != nil {
			log.Warn("failed to load table entry location", log.String("tableName", tableName))
			return err
		}
	} else {
		key := route.NewObTableEntryKey(c.clusterName, c.tenantName, c.database, tableName)
		*entry, err = route.GetTableEntryFromRemote(c.serverRoster.GetServer(), &c.sysUA, key)
		if err != nil {
			log.Warn("failed to get table entry from remote", log.String("key", key.String()))
			return err
		}
	}

	// 2. Set rowKey element to entry.
	if (*entry).IsPartitionTable() {
		rowKeyElement, ok := c.tableRowKeyElement[tableName]
		if !ok {
			log.Warn("failed to get rowKey element by table name", log.String("tableName", tableName))
			return errors.New("failed to get rowKey element by table name")
		}
		(*entry).SetRowKeyElement(rowKeyElement)
	}

	// 3. todo:prepare the table entry for weak read.

	// 4. Put entry to cache.
	c.tableLocations.Store(tableName, entry)

	return nil
}

func (c *ObClient) loadTableEntryLocation(entry *route.ObTableEntry) error {
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
		log.Warn("failed to new db",
			log.String("sysUA", c.sysUA.String()),
			log.String("addr", addr.String()))
		return err
	}
	defer func() {
		_ = db.Close()
	}()

	locEntry, err := route.GetPartLocationEntryFromRemote(db, entry)
	if err != nil {
		log.Warn("failed to get part location entry from remote", log.String("entry", entry.String()))
		return err
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
		log.Warn("failed to lock metadata refreshing timeout", log.Int64("timeout", timeout))
		return errors.New("failed to lock metadata refreshing timeout")
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
		log.Warn("failed fetch meta data")
		return err
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
		log.Warn("failed to load ocp model",
			log.String("configUrl", c.configUrl),
			log.String("localFileName", c.config.RslistLocalFileLocation))
		return err
	}
	c.ocpModel = ocpModel
	addr := c.ocpModel.GetServerAddressRandomly()

	// 2. Get ob cluster version and init route sql
	ver, err := route.GetObVersionFromRemote(addr, &c.sysUA)
	if err != nil {
		log.Warn("failed to get ob version from remote",
			log.String("addr", addr.String()),
			log.String("sysUA", c.sysUA.String()))
		return err
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
	entry, err := route.GetTableEntryFromRemote(addr, &c.sysUA, rootServerKey)
	if err != nil {
		log.Warn("failed to dummy tenant server from remote",
			log.String("addr", addr.String()),
			log.String("sysUA", c.sysUA.String()),
			log.String("key", rootServerKey.String()))
		return err
	}
	// 3.2 Save all tenant server address
	replicaLocations := entry.TableLocation().ReplicaLocations()
	servers := make([]*route.ObServerAddr, 0, len(replicaLocations))
	for _, replicaLoc := range replicaLocations {
		info := replicaLoc.Info()
		addr := replicaLoc.Addr()
		if !info.IsActive() {
			log.Warn("server is not active",
				log.String("server info", info.String()),
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
			log.Warn("fail to init ob table", log.String("obTable", addr.String()))
			return err
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
	if !entry.IsPartitionTable() || entry.PartitionInfo().Level().Index() == route.PartLevelZeroIndex {
		return 0, nil
	}
	if entry.PartitionInfo().Level().Index() == route.PartLevelOneIndex {
		return entry.PartitionInfo().FirstPartDesc().GetPartId(rowKeyValue)
	}
	if entry.PartitionInfo().Level().Index() == route.PartLevelTwoIndex {
		partId1, err := entry.PartitionInfo().FirstPartDesc().GetPartId(rowKeyValue)
		if err != nil {
			log.Warn("failed to get part id from first part desc",
				log.String("first part desc", entry.PartitionInfo().FirstPartDesc().String()))
			return -1, err
		}
		partId2, err := entry.PartitionInfo().SubPartDesc().GetPartId(rowKeyValue)
		if err != nil {
			log.Warn("failed to get part id from sub part desc",
				log.String("sub part desc", entry.PartitionInfo().SubPartDesc().String()))
			return -1, err
		}
		return (int64(partId1) << route.ObPartIdShift) | partId2 | route.ObMask, nil
	}
	log.Warn("unknown partition level", log.String("part info", entry.PartitionInfo().String()))
	return -1, errors.New("unknown partition level")
}

func (c *ObClient) getTable(
	tableName string,
	entry *route.ObTableEntry,
	partId int64) (*ObTable, error) {
	// 1. Get replica location by partition id
	replicaLoc, err := entry.GetPartitionReplicaLocation(partId, route.NewObServerRoute(false))
	if err != nil {
		log.Warn("failed to get partition replica location", log.Int64("partId", partId))
		return nil, err
	}

	// 2. Get table from table Roster by server addr
	t, ok := c.tableRoster.Load(*replicaLoc.Addr())
	if !ok {
		log.Warn("failed to get table by server addr",
			log.String("addr", replicaLoc.Addr().String()))
		return nil, errors.New("failed to get table by server addr")
	}
	// todo: check server addr is expired or not
	tb, _ := t.(*ObTable)
	return tb, nil
}
