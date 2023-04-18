package client

import (
	"errors"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
	"strings"
	"sync"
)

func NewClient(
	configUrl string,
	fullUserName string,
	passWord string,
	sysUserName string,
	sysPassWord string,
	cliConfig *config.ClientConfig) (Client, error) {
	// 1. Check args
	if configUrl == "" {
		log.Warn("config url is empty")
		return nil, errors.New("config url is empty")
	}
	if fullUserName == "" || sysUserName == "" {
		log.Warn("user name is empty",
			log.String("fullUserName", fullUserName),
			log.String("sysUserName", sysUserName))
		return nil, errors.New("user name is null")
	}
	if cliConfig == nil {
		log.Warn("client config is nil")
		return nil, errors.New("client config is nil")
	}

	// 2. New client and init
	cli, err := newObClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cliConfig)
	if err != nil {
		log.Warn("failed to new obClient",
			log.String("configUrl", configUrl),
			log.String("fullUserName", fullUserName))
		return nil, err
	}
	err = cli.init()
	if err != nil {
		log.Warn("failed to init client", log.String("client", cli.String()))
		return nil, err
	}

	return cli, nil
}

type ObkvOption struct {
}

type Client interface {
	AddRowkey(tableName string, rowkey []string) error
	Insert(tableName string, rowkey []table.Column, mutateColumns []table.Column, opts ...ObkvOption) (int64, error)
	Get(tableName string, rowkey []table.Column, getColumns []string, opts ...ObkvOption) (map[string]interface{}, error)
}

type ObClient struct {
	configUrl          string
	fullUserName       string
	userName           string
	tenantName         string
	clusterName        string
	password           string
	database           string
	sysUA              route.ObUserAuth
	ocpModel           *ObOcpModel
	tableRoster        sync.Map
	serverRoster       route.ObServerRoster
	tableRowKeyElement map[string]*table.ObRowkeyElement
	config             config.ClientConfig
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
	cli.config = *cliConfig
	cli.tableRowKeyElement = make(map[string]*table.ObRowkeyElement)

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
	// 1. todo: load ocp mode to get RsList
	c.ocpModel = &ObOcpModel{}
	addr := c.ocpModel.getServerAddressRandomly()

	// 2. Get ob version.
	ver, err := route.GetObVersionFromRemote(addr, &c.sysUA)
	if err != nil {
		log.Warn("failed to get ob version from remote",
			log.String("addr", addr.String()),
			log.String("sysUA", c.sysUA.String()))
		return err
	}
	// 2.1 Set ob version and init route sql by ob version.
	if util.ObVersion() != 0.0 {
		util.SetObVersion(ver)
		route.InitSql(ver)
	}

	// 3. All dummy to get tenant server
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
	var tableRoster sync.Map
	var servers []*route.ObServerAddr
	replicaLocations := entry.TableLocation().ReplicaLocations()
	for _, replicaLoc := range replicaLocations {
		info := replicaLoc.Info()
		addr := replicaLoc.Addr()
		if !info.IsActive() {
			log.Warn("server is not active",
				log.String("server info", info.String()),
				log.String("server addr", addr.String()))
			continue
		}
		t, err := table.NewObTable(
			addr.Ip(),
			addr.SvrPort(),
			c.tenantName,
			c.userName,
			c.password,
			c.database,
			c.config.ConnPoolMaxConnSize,
			c.config.ConnTimeOut,
		)
		if err != nil {
			log.Warn("fail to new ob table",
				log.String("addr", addr.String()),
				log.String("tenantName", c.tenantName),
				log.String("userName", c.userName),
				log.String("password", c.password),
				log.String("database", c.database),
				log.Int("ConnPoolMaxConnSize", c.config.ConnPoolMaxConnSize))
			return err
		}
		tableRoster.Store(*addr, t)
		servers = append(servers, addr)
	}
	c.tableRoster = tableRoster
	c.serverRoster.Reset(servers)

	// 4. todo: Get Server LDC info for weak read consistency.
	return nil
}

func (c *ObClient) AddRowkey(tableName string, rowkey []string) error {
	if tableName == "" || len(rowkey) == 0 {
		log.Warn("nil table name or empty rowkey",
			log.String("tableName", tableName),
			log.Int("rowkey size", len(rowkey)))
		return errors.New("nil table name or empty rowkey")
	}
	m := make(map[string]int, 1)
	for i := 0; i < len(rowkey); i++ {
		columnName := rowkey[i]
		m[columnName] = i
	}
	c.tableRowKeyElement[tableName] = table.NewObRowkeyElement(m)
	return nil
}

func (c *ObClient) Insert(
	tableName string,
	rowkey []table.Column,
	mutateColumns []table.Column,
	opts ...ObkvOption) (int64, error) {
	var mutateColNames []string
	var mutateColValues []interface{}
	for _, col := range mutateColumns {
		mutateColNames = append(mutateColNames, col.Name)
		mutateColValues = append(mutateColValues, col.Value)
	}
	res, err := c.execute(
		tableName,
		protocol.ObTableOperationTypeInsert,
		rowkey,
		mutateColNames,
		mutateColValues,
		opts...)
	if err != nil {
		log.Warn("failed to execute insert",
			log.String("tableName", tableName),
			log.String("rowkey", columnsToString(rowkey)),
			log.String("mutateColumns", columnsToString(mutateColumns)))
		return -1, err
	}
	return res.AffectedRows(), nil
}

func (c *ObClient) Get(
	tableName string,
	rowkey []table.Column,
	getColumns []string,
	opts ...ObkvOption) (map[string]interface{}, error) {
	res, err := c.execute(
		tableName,
		protocol.ObTableOperationTypeGet,
		rowkey,
		getColumns,
		nil,
		opts...)
	if err != nil {
		log.Warn("failed to execute get",
			log.String("tableName", tableName),
			log.String("rowkey", columnsToString(rowkey)),
			log.String("getColumns", util.StringArrayToString(getColumns)))
		return nil, err
	}
	return res.Entity().Properties(), nil
}

func (c *ObClient) execute(
	tableName string,
	opType protocol.ObTableOperationType,
	rowkey []table.Column,
	columns []string,
	properties []interface{},
	opts ...ObkvOption) (*protocol.ObTableOperationResult, error) {
	var rowkeyValue []interface{}
	for _, col := range rowkey {
		rowkeyValue = append(rowkeyValue, col.Value)
	}
	// 1. Get table route
	tableParam, err := c.getTableParam(tableName, rowkeyValue, false /* refresh */)
	if err != nil {
		log.Warn("failed to get table param",
			log.String("tableName", tableName),
			log.Int8("opType", int8(opType)))
		return nil, err
	}

	// 2. Construct request.
	request, err := protocol.NewObTableOperationRequest(
		tableName,
		tableParam.TableId(),
		tableParam.PartitionId(),
		opType,
		rowkeyValue,
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
	result := new(protocol.ObTableOperationResult)
	err = tableParam.Table().Execute(request, result)
	if err != nil {
		log.Warn("failed to execute request", log.String("request", request.String()))
		return nil, err
	}
	return result, nil
}

func (c *ObClient) getTableParam(
	tableName string,
	rowkeyValue []interface{},
	refresh bool) (*table.ObTableParam, error) {
	entry, err := c.getOrRefreshTableEntry(tableName, refresh)
	if err != nil {
		log.Warn("failed to get or refresh table entry", log.String("tableName", tableName))
		return nil, err
	}
	partId, err := c.getPartitionId(entry, rowkeyValue)
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
	return table.NewObTableParam(t, entry.TableId(), partId), nil
}

func (c *ObClient) getOrRefreshTableEntry(tableName string, refresh bool) (*route.ObTableEntry, error) {
	key := route.NewObTableEntryKey(c.clusterName, c.tenantName, c.database, tableName)
	// todo：add refresh logic
	entry, err := route.GetTableEntryFromRemote(c.serverRoster.GetServer(), &c.sysUA, key)
	if err != nil {
		log.Warn("failed to get table entry from remote",
			log.String("server", c.serverRoster.GetServer().String()),
			log.String("sysUA", c.sysUA.String()),
			log.String("entry key", key.String()))
		return nil, err
	}

	if entry.IsPartitionTable() {
		rowkeyElement, ok := c.tableRowKeyElement[tableName]
		if !ok {
			log.Warn("failed to get rowkey element by table name", log.String("tableName", tableName))
			return nil, errors.New("failed to get rowkey element by table name")
		}
		entry.SetRowKeyElement(rowkeyElement)
	}

	return entry, nil
}

// get partition id by rowkey
func (c *ObClient) getPartitionId(entry *route.ObTableEntry, rowkeyValue []interface{}) (int64, error) {
	if !entry.IsPartitionTable() || entry.PartitionInfo().Level().Index() == route.PartLevelZeroIndex {
		return 0, nil
	}
	if entry.PartitionInfo().Level().Index() == route.PartLevelOneIndex {
		return entry.PartitionInfo().FirstPartDesc().GetPartId(rowkeyValue)
	}
	if entry.PartitionInfo().Level().Index() == route.PartLevelTwoIndex {
		partId1, err := entry.PartitionInfo().FirstPartDesc().GetPartId(rowkeyValue)
		if err != nil {
			log.Warn("failed to get part id from first part desc",
				log.String("first part desc", entry.PartitionInfo().FirstPartDesc().String()))
			return -1, err
		}
		partId2, err := entry.PartitionInfo().SubPartDesc().GetPartId(rowkeyValue)
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
	partId int64) (*table.ObTable, error) {
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
	tb, _ := t.(*table.ObTable)
	return tb, nil
}
