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
	"fmt"
	"runtime"
	"strings"
	"time"

	"golang.org/x/sys/unix"

	"github.com/oceanbase/obkv-table-client-go/log"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/client/option"

	"github.com/oceanbase/obkv-table-client-go/config"
	oberror "github.com/oceanbase/obkv-table-client-go/error"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type obClient struct {
	config       *config.ClientConfig
	configUrl    string
	fullUserName string
	userName     string
	tenantName   string
	clusterName  string
	password     string
	database     string
	sysUA        *route.ObUserAuth

	// for odp client
	odpTable   *route.ObTable
	odpIP      string
	odpRpcPort int

	routeInfo *route.ObRouteInfo
}

func newObClient(
	configUrl string,
	fullUserName string,
	passWord string,
	sysUserName string,
	sysPassWord string,
	cliConfig *config.ClientConfig) (*obClient, error) {
	cli := new(obClient)
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
	cli.sysUA = route.NewObUserAuth(sysUserName, sysPassWord)
	cli.config = cliConfig
	cli.routeInfo = route.NewRouteInfo(cli.sysUA)

	return cli, nil
}

func newOdpClient(
	fullUserName string,
	passWord string,
	odpIP string,
	odpRpcPort int,
	database string,
	cliConfig *config.ClientConfig) (*obClient, error) {
	cli := new(obClient)
	// 1. Parse full username to get userName/tenantName/clusterName
	err := cli.parseOdpFullUserName(fullUserName)
	if err != nil {
		return nil, errors.WithMessagef(err, "parse full user name, fullUserName:%s", fullUserName)
	}
	// 2. init other members
	cli.password = passWord
	cli.odpIP = odpIP
	cli.odpRpcPort = odpRpcPort
	cli.database = database
	cli.config = cliConfig
	cli.routeInfo = nil

	return cli, nil
}

func (c *obClient) String() string {
	var configStr = "nil"
	if c.config != nil {
		configStr = c.config.String()
	}

	var sysUAStr = "nil"
	if c.sysUA != nil {
		sysUAStr = c.sysUA.String()
	}
	return "obClient{" +
		"config:" + configStr + ", " +
		"configUrl:" + c.configUrl + ", " +
		"fullUserName:" + c.fullUserName + ", " +
		"userName:" + c.userName + ", " +
		"tenantName:" + c.tenantName + ", " +
		"clusterName:" + c.clusterName + ", " +
		"database:" + c.database + ", " +
		"sysUA:" + sysUAStr +
		"}"
}

// GetRpcFlag get rpc header flag
func (c *obClient) GetRpcFlag() uint16 {
	rpcFlag := protocol.RpcHeaderDefaultFlag
	if c.config.EnableRerouting {
		rpcFlag |= protocol.RequireReroutingFlag
	}
	return rpcFlag
}

// standard format: user_name@tenant_name#cluster_name
func (c *obClient) parseFullUserName(fullUserName string) error {
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

// standard format: user_name@tenant_name#cluster_name or user_name for VIP
func (c *obClient) parseOdpFullUserName(fullUserName string) error {
	utIndex := strings.Index(fullUserName, "@")
	tcIndex := strings.Index(fullUserName, "#")
	if utIndex == -1 || tcIndex == -1 || tcIndex <= utIndex {
		c.userName = fullUserName
		c.fullUserName = fullUserName
	} else {
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
	}
	return nil
}

// format: http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx
func (c *obClient) parseConfigUrl(configUrl string) error {
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

func (c *obClient) init() error {
	// 1. get rslist from config server by http get
	err := c.routeInfo.FetchConfigServerInfo(
		c.configUrl,
		c.config.RsListLocalFileLocation,
		c.config.RsListHttpGetTimeout,
		c.config.RsListHttpGetRetryTimes,
		c.config.RsListHttpGetRetryInterval,
	)
	if err != nil {
		return errors.WithMessagef(err, "get rslist, url:%s", c.configUrl)
	}

	// 2. fetch cluster version, accessing route table depends on the cluster version. And check tenant exist.
	err = c.routeInfo.CheckClusterAndTenant(c.tenantName)
	if err != nil {
		return errors.WithMessagef(err, "get cluster version")
	}
	if util.ObVersion() == 0.0 {
		util.SetObVersion(c.routeInfo.ClusterVersion())
		route.InitSql(c.routeInfo.ClusterVersion())
	}

	// 3. fetch server roster which means the server that contains the tenant
	err = c.routeInfo.FetchServerRoster(c.clusterName, c.tenantName)
	if err != nil {
		return errors.WithMessagef(err, "get server roster")
	}

	// 4. construct table roster which means creating connection pool for each server
	err = c.routeInfo.ConstructTableRoster(
		c.userName,
		c.password,
		c.database,
		c.config.ConnPoolMaxConnSize,
		c.config.ConnConnectTimeOut,
		c.config.ConnLoginTimeout,
	)
	if err != nil {
		return err
	}

	// 5. Run background task
	go c.routeInfo.RunBackgroundTask()
	go c.routeInfo.RunTickerTask()

	return nil
}

func (c *obClient) initOdp() error {
	// 1. Init odp table
	t := route.NewObTable(c.odpIP, c.odpRpcPort, c.tenantName, c.fullUserName, c.password, c.database)
	t.SetMaxConnectionAge(c.config.MaxConnectionAge)
	t.SetEnableSLBLoadBalance(c.config.EnableSLBLoadBalance)
	err := t.Init(c.config.ConnPoolMaxConnSize, c.config.ConnConnectTimeOut, c.config.ConnLoginTimeout)
	// 2. Init sql
	// ObVersion will be set when login in init()
	route.InitSql(util.ObVersion())
	if err != nil {
		return err
	}
	// 3. Set ODP Table
	c.odpTable = t

	return nil
}

func (c *obClient) Insert(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (int64, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	if operationOptions.TableFilter == nil {
		res, err := c.executeWithRetry(
			ctx,
			tableName,
			protocol.ObTableOperationInsert,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return -1, err
		}
		return res.AffectedRows(), nil
	} else {
		res, err := c.executeWithFilterAndRetry(
			ctx,
			tableName,
			protocol.ObTableOperationInsert,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return -1, err
		}
		return res.AffectedRows(), nil
	}
}

func (c *obClient) Update(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (int64, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	if operationOptions.TableFilter == nil {
		res, err := c.executeWithRetry(
			ctx,
			tableName,
			protocol.ObTableOperationUpdate,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return -1, err
		}
		return res.AffectedRows(), nil
	} else {
		res, err := c.executeWithFilterAndRetry(
			ctx,
			tableName,
			protocol.ObTableOperationUpdate,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return -1, err
		}
		return res.AffectedRows(), nil
	}
}

func (c *obClient) InsertOrUpdate(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (int64, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	res, err := c.executeWithRetry(
		ctx,
		tableName,
		protocol.ObTableOperationInsertOrUpdate,
		rowKey,
		mutateColumns,
		operationOptions)
	if err != nil {
		return -1, err
	}
	return res.AffectedRows(), nil
}

func (c *obClient) InsertOrUpdateWithResult(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (SingleResult, error) {
	operationOptions := c.getOperationOptions(opts...)
	res, err := c.executeWithRetry(
		ctx,
		tableName,
		protocol.ObTableOperationInsertOrUpdate,
		rowKey,
		mutateColumns,
		operationOptions)
	if err != nil {
		return nil, errors.WithMessagef(err, "execute insert or update, tableName:%s, rowKey:%s, mutateColumns:%s",
			tableName, table.ColumnsToString(rowKey), table.ColumnsToString(mutateColumns))
	}
	return newObSingleResult(res.AffectedRows(), nil, res.Flags()), nil
}

func (c *obClient) Replace(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (int64, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	res, err := c.executeWithRetry(
		ctx,
		tableName,
		protocol.ObTableOperationReplace,
		rowKey,
		mutateColumns,
		operationOptions)
	if err != nil {
		return -1, err
	}
	return res.AffectedRows(), nil
}

func (c *obClient) Increment(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (SingleResult, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	if operationOptions.TableFilter == nil {
		res, err := c.executeWithRetry(
			ctx,
			tableName,
			protocol.ObTableOperationIncrement,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return nil, err
		}
		return newObSingleResult(res.AffectedRows(), res.Entity(), res.Flags()), nil
	} else {
		res, err := c.executeWithFilterAndRetry(
			ctx,
			tableName,
			protocol.ObTableOperationIncrement,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return nil, err
		}
		return newObSingleResult(res.AffectedRows(), nil, 0), nil
	}
}

func (c *obClient) Append(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	mutateColumns []*table.Column,
	opts ...option.ObOperationOption) (SingleResult, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	if operationOptions.TableFilter == nil {
		res, err := c.executeWithRetry(
			ctx,
			tableName,
			protocol.ObTableOperationAppend,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return nil, err
		}
		return newObSingleResult(res.AffectedRows(), res.Entity(), res.Flags()), nil
	} else {
		res, err := c.executeWithFilterAndRetry(
			ctx,
			tableName,
			protocol.ObTableOperationAppend,
			rowKey,
			mutateColumns,
			operationOptions)
		if err != nil {
			return nil, err
		}
		return newObSingleResult(res.AffectedRows(), nil, 0), nil
	}
}

func (c *obClient) Delete(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	opts ...option.ObOperationOption) (int64, error) {
	log.InitTraceId(&ctx)
	operationOptions := c.getOperationOptions(opts...)
	if operationOptions.TableFilter == nil {
		res, err := c.executeWithRetry(
			ctx,
			tableName,
			protocol.ObTableOperationDel,
			rowKey,
			nil,
			operationOptions)
		if err != nil {
			return -1, err
		}
		return res.AffectedRows(), nil
	} else {
		res, err := c.executeWithFilterAndRetry(
			ctx,
			tableName,
			protocol.ObTableOperationDel,
			rowKey,
			nil,
			operationOptions)
		if err != nil {
			return -1, err
		}
		return res.AffectedRows(), nil
	}
}

func (c *obClient) Get(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column,
	getColumns []string,
	opts ...option.ObOperationOption) (SingleResult, error) {
	log.InitTraceId(&ctx)
	var columns = make([]*table.Column, 0, len(getColumns))
	for _, columnName := range getColumns {
		columns = append(columns, table.NewColumn(columnName, nil))
	}
	operationOptions := c.getOperationOptions(opts...)
	res, err := c.executeWithRetry(
		ctx,
		tableName,
		protocol.ObTableOperationGet,
		rowKey,
		columns,
		operationOptions)
	if err != nil {
		return nil, err
	}
	return newObSingleResult(res.AffectedRows(), res.Entity(), res.Flags()), nil
}

func (c *obClient) Query(ctx context.Context, tableName string, rangePairs []*table.RangePair, opts ...option.ObQueryOption) (QueryResultIterator, error) {
	log.InitTraceId(&ctx)
	queryOpts := c.getObQueryOptions(opts...)
	queryExecutor := newObQueryExecutorWithParams(tableName, c)
	queryExecutor.addKeyRanges(rangePairs)
	queryExecutor.setQueryOptions(queryOpts)
	res, err := queryExecutor.init(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *obClient) NewBatchExecutor(tableName string, opts ...option.ObBatchOption) BatchExecutor {
	batchOpts := c.getObBatchOptions(opts...)
	batchExecutor := newObBatchExecutor(tableName, c)
	batchExecutor.setBatchOptions(batchOpts)
	return batchExecutor
}

func (c *obClient) NewAggExecutor(tableName string, rangePairs []*table.RangePair, opts ...option.ObQueryOption) AggExecutor {
	queryOpts := c.getObQueryOptions(opts...)
	queryExecutor := newObQueryExecutorWithParams(tableName, c)
	queryExecutor.addKeyRanges(rangePairs)
	queryExecutor.setQueryOptions(queryOpts)
	return newObAggExecutor(queryExecutor)
}

func (c *obClient) Close() {
	if c.routeInfo != nil {
		c.routeInfo.Close()
	}

	if c.odpTable != nil {
		c.odpTable.Close()
	}
	_ = log.Sync()
}

func (c *obClient) getOperationOptions(opts ...option.ObOperationOption) *option.ObOperationOptions {
	operationOptions := option.NewOperationOptions()
	for _, opt := range opts {
		opt.Apply(operationOptions)
	}
	return operationOptions
}

func (c *obClient) getObQueryOptions(options ...option.ObQueryOption) *option.ObQueryOptions {
	opts := option.NewObQueryOption()
	for _, op := range options {
		op.Apply(opts)
	}
	return opts
}

func (c *obClient) getObBatchOptions(options ...option.ObBatchOption) *option.ObBatchOptions {
	opts := option.NewObBatchOption()
	for _, op := range options {
		op.Apply(opts)
	}
	return opts
}

func (c *obClient) executeWithRetry(
	ctx context.Context,
	tableName string,
	opType protocol.ObTableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	operationOptions *option.ObOperationOptions) (*protocol.ObTableOperationResponse, error) {

	if _, ok := ctx.Deadline(); !ok {
		ctx, _ = context.WithTimeout(ctx, c.config.OperationTimeOut) // default timeout operation timeout
	}

	res, err, needRetry := c.execute(ctx, tableName, opType, rowKey, columns, operationOptions)
	for err != nil && needRetry {
		select {
		case <-ctx.Done():
			return nil, errors.WithMessage(err, "retry and timeout")
		default:
			res, err, needRetry = c.execute(ctx, tableName, opType, rowKey, columns, operationOptions)
			time.Sleep(1 * time.Millisecond)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *obClient) execute(
	ctx context.Context,
	tableName string,
	opType protocol.ObTableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	operationOptions *option.ObOperationOptions) (*protocol.ObTableOperationResponse, error, bool) {

	needRetry := false
	// 1. Get table route
	tableParam, err := c.GetTableParam(ctx, tableName, rowKey)
	if err != nil {
		log.Error("Runtime", ctx.Value(log.ObkvTraceIdName), "error occur in execute",
			log.Int64("opType", int64(opType)), log.String("tableName", tableName), log.String("tableParam", tableParam.String()))
		return nil, errors.WithMessagef(err, "get table param, tableName:%s, opType:%d", tableName, opType), needRetry
	}

	// 2. Construct request.
	request, err := protocol.NewObTableOperationRequestWithParams(
		tableName,
		tableParam.TableId(),
		tableParam.PartitionId(),
		opType,
		rowKey,
		columns,
		operationOptions.ReturnRowKey,
		operationOptions.ReturnAffectedEntity,
		c.config.OperationTimeOut,
		c.GetRpcFlag(),
	)
	if err != nil {
		log.Error("Runtime", ctx.Value(log.ObkvTraceIdName), "error occur in execute",
			log.Int64("opType", int64(opType)), log.String("tableName", tableName), log.String("tableParam", tableParam.String()))
		return nil, errors.WithMessagef(err, "new operation request, tableName:%s, tableParam:%s, opType:%d",
			tableName, tableParam.String(), opType), needRetry
	}

	// 3. execute
	result := protocol.NewObTableOperationResponse()
	err, needRetry = c.executeInternal(ctx, tableName, tableParam.Table(), request, result)
	if err != nil {
		trace := fmt.Sprintf("Y%X-%016X", result.UniqueId(), result.Sequence())
		log.Error("Runtime", ctx.Value(log.ObkvTraceIdName), "error occur in execute", log.String("observerTraceId", trace))
		return nil, err, needRetry
	}

	if oberror.ObErrorCode(result.Header().ErrorNo()) != oberror.ObSuccess {
		trace := fmt.Sprintf("Y%X-%016X", result.UniqueId(), result.Sequence())
		log.Error("Runtime", ctx.Value(log.ObkvTraceIdName), "error occur in execute", log.String("observerTraceId", trace))
		return nil, protocol.NewProtocolError(
			result.RemoteAddr().String(),
			oberror.ObErrorCode(result.Header().ErrorNo()),
			result.Header().Msg(),
			result.Sequence(),
			result.UniqueId(),
		), needRetry
	}

	return result, nil, needRetry
}

func (c *obClient) executeWithFilterAndRetry(
	ctx context.Context,
	tableName string,
	opType protocol.ObTableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	operationOptions *option.ObOperationOptions) (*protocol.ObTableQueryAndMutateResponse, error) {

	if _, ok := ctx.Deadline(); !ok {
		ctx, _ = context.WithTimeout(ctx, c.config.OperationTimeOut) // default timeout operation timeout
	}

	res, err, needRetry := c.executeWithFilter(ctx, tableName, opType, rowKey, columns, operationOptions)
	for err != nil && needRetry {
		select {
		case <-ctx.Done():
			return nil, errors.WithMessage(err, "retry and timeout")
		default:
			res, err, needRetry = c.executeWithFilter(ctx, tableName, opType, rowKey, columns, operationOptions)
			time.Sleep(1 * time.Millisecond)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *obClient) executeWithFilter(
	ctx context.Context,
	tableName string,
	opType protocol.ObTableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	operationOptions *option.ObOperationOptions) (*protocol.ObTableQueryAndMutateResponse, error, bool) {

	needRetry := false

	// 1. Get table route
	tableParam, err := c.GetTableParam(ctx, tableName, rowKey)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table param, tableName:%s, opType:%d", tableName, opType), needRetry
	}

	// 2. Construct request.
	request, err := protocol.NewObTableQueryAndMutateRequestWithRowKeyAndParams(
		tableName,
		tableParam.TableId(),
		tableParam.PartitionId(),
		opType,
		rowKey,
		columns,
		c.config.OperationTimeOut,
		c.GetRpcFlag(),
	)
	if err != nil {
		return nil, errors.WithMessagef(err, "new operation request, tableName:%s, tableParam:%s, opType:%d",
			tableName, tableParam.String(), opType), needRetry
	}

	request.TableQueryAndMutate().TableQuery().SetFilterString(operationOptions.TableFilter.String())

	if opType == protocol.ObTableOperationInsert {
		// set query range into table query
		keyRanges, err := TransferQueryRange(operationOptions.ScanRange)
		if err != nil {
			return nil, errors.WithMessage(err, "transfer query range"), needRetry
		}
		request.TableQueryAndMutate().TableQuery().SetKeyRanges(keyRanges)
	}

	// 3. execute
	result := protocol.NewObTableQueryAndMutateResponse()
	err, needRetry = c.executeInternal(ctx, tableName, tableParam.Table(), request, result)
	if err != nil {
		return nil, err, needRetry
	}

	return result, nil, needRetry
}

func (c *obClient) executeInternal(
	ctx context.Context,
	tableName string,
	table *route.ObTable,
	request protocol.ObPayload,
	result protocol.ObPayload) (error, bool) {

	if c.routeInfo != nil {
		return c.routeInfo.Execute(ctx, tableName, table, request, result)
	}

	_, err := table.Execute(ctx, request, result)
	return err, false
}

func (c *obClient) GetTableParam(
	ctx context.Context,
	tableName string,
	rowKey []*table.Column) (*route.ObTableParam, error) {

	if c.odpTable != nil {
		return route.NewObTableParam(c.odpTable, 0, 0), nil
	}

	return c.routeInfo.GetTableParam(ctx, tableName, rowKey)
}

func MonitorSlowQuery(executeTime int64, slowQueryThreshold int64, tableName string, clientTraceId any) {
	if executeTime > slowQueryThreshold {
		pId := unix.Getpid()
		buf := make([]byte, 64)
		n := runtime.Stack(buf, false)
		id := buf[:n]
		var goroutineID uint64
		fmt.Sscanf(string(id), "goroutine %d", &goroutineID)
		log.Info("Monitor", clientTraceId, "SlowQuery", log.String("tableName", tableName),
			log.Int64("executeTime", executeTime), log.Int64("slowQueryThreshold", slowQueryThreshold),
			log.Int("pId", pId), log.Uint64("goroutineID", goroutineID))
	}
}
