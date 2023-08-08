/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package reroute

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	getReplicaSql      = "SELECT A.table_id as table_id, A.partition_id as partition_id, A.svr_ip as svr_ip, B.svr_port as svr_port, A.role as role FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ? and partition_id in "
	getLDIdSql         = "SELECT ls_id FROM oceanbase.cdb_ob_table_locations WHERE tenant_id = 1 and database_name = ? and table_name = ? and role = 'LEADER'"
	getFollowerAddrSql = "SELECT concat(svr_ip,':',svr_port) AS host FROM oceanbase.__all_virtual_ls_meta_table WHERE tenant_id = 1 and ls_id = ? and role = 2 and replica_status = 'NORMAL' limit 1;"
	switch2FollwerSql  = "ALTER SYSTEM SWITCH REPLICA LEADER ls = %d server= '%s' tenant='sys'"
)

func disableAutoReplicaSwitch() error {
	_, err := test.GlobalDB.Exec("alter system set enable_auto_leader_switch = 'false';")
	if err != nil {
		return errors.WithMessagef(err, "execute disable auto replica switch sql")
	}
	return nil
}

func createInStatement(partNum int) string {
	values := make([]uint64, partNum)
	for i := 0; i < partNum; i++ {
		values[i] = uint64(i)
	}

	// Create inStatement "(0,1,2...partNum);".
	var inStatement string
	inStatement += "("
	for i, v := range values {
		if i > 0 {
			inStatement += ", "
		}
		inStatement += strconv.FormatUint(v, 10)
	}
	inStatement += ");"
	return inStatement
}

func getPartitions(
	ctx context.Context,
	tenantName string,
	databaseName string,
	tableName string,
	partNum int,
	partitions []*Partition) error {

	sql := getReplicaSql + createInStatement(partNum)
	rows, err := test.GlobalDB.QueryContext(ctx, sql, tenantName, databaseName, tableName)
	if err != nil {
		return errors.WithMessagef(err, "sql query, sql:%s", sql)
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		tableId uint64
		partId  uint64
		ip      string
		port    int
		role    int
	)

	isEmpty := true
	for rows.Next() {
		if isEmpty {
			isEmpty = false
		}

		err := rows.Scan(
			&tableId,
			&partId,
			&ip,
			&port,
			&role,
		)
		if err != nil {
			return errors.WithMessagef(err, "scan row")
		}

		replica := &Replica{
			TableId: tableId,
			PartId:  partId,
			Ip:      ip,
			Port:    port,
			Role:    protocol.ObRole(role),
		}

		if replica.Role == protocol.ObRoleLeader {
			for _, partition := range partitions {
				if partition.partId == partId {
					partition.tableId = tableId
					partition.leader = replica
					partition.replicaNum++
				}
			}
		} else { // follower
			for _, partition := range partitions {
				if partition.partId == partId {
					partition.follower = append(partition.follower, replica)
					partition.replicaNum++
				}
			}
		}
	}

	if isEmpty {
		return errors.Errorf("empty set")
	}

	return nil
}

func switchLeader(partitions []*Partition) error {
	partNum := len(partitions)
	for _, partition := range partitions {
		if len(partition.follower) != 2 {
			return errors.Errorf("invalid follower num:%d", len(partition.follower))
		}
	}

	// get a follower ip and port
	ip := partitions[0].follower[0].Ip
	port := partitions[0].follower[0].Port
	server := ip + ":" + strconv.Itoa(port)

	// switch all replica leader to ip:port
	for _, partition := range partitions {
		partIdStr := fmt.Sprintf("%d%%%d@%d", partition.partId, partNum, partition.tableId)
		sql := fmt.Sprintf("ALTER SYSTEM SWITCH REPLICA LEADER PARTITION_ID '%s' SERVER '%s';", partIdStr, server)
		_, err := test.GlobalDB.Exec(sql)
		if err != nil {
			return errors.WithMessagef(err, "execute switch leader sql, partIdStr:%s, server:%s", partIdStr, server)
		}
	}

	return nil
}

func SwitchReplicaLeaderRandomly(
	tenantName string,
	databaseName string,
	tableName string,
	partNum int) error {

	if partNum != 2 {
		return errors.Errorf("invalid partition num:%d", partNum)
	}

	// 1. disable auto replica switch
	err := disableAutoReplicaSwitch()
	if err != nil {
		return errors.WithMessagef(err, "disable auto replica switch")
	}

	// 2. query replica
	partitions := make([]*Partition, 0, partNum)
	for i := 0; i < partNum; i++ {
		partitions = append(partitions, NewPartition(uint64(i)))
	}
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Second) // 5s
	err = getPartitions(ctx, tenantName, databaseName, tableName, partNum, partitions)
	if err != nil {
		return errors.WithMessagef(err, "get replicas")
	}

	// 3. check replica
	if len(partitions) != partNum {
		return errors.Errorf("invalid partition num:%d", len(partitions))
	}

	// 4. switch leader
	err = switchLeader(partitions)
	if err != nil {
		return errors.WithMessagef(err, "switch leader")
	}

	return nil
}

func getLsId(ctx context.Context, databaseName string, tableName string) (uint64, error) {
	rows, err := test.GlobalDB.QueryContext(ctx, getLDIdSql, databaseName, tableName)
	if err != nil {
		return 0, errors.WithMessagef(err, "sql query, sql:%s", getLDIdSql)
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		lsId     uint64
		prevLsId uint64
	)
	prevLsId = math.MaxUint64

	isEmpty := true
	for rows.Next() {
		if isEmpty {
			isEmpty = false
		}

		err := rows.Scan(&lsId)
		if err != nil {
			return 0, errors.WithMessagef(err, "scan row")
		}
		if prevLsId == math.MaxUint64 {
			prevLsId = lsId
		} else {
			if lsId != prevLsId {
				return 0, errors.Errorf("suppose ls is the same for every partition, lsId:%d, prevLsId:%d", lsId, prevLsId)
			}
		}
	}
	if isEmpty {
		return 0, errors.Errorf("empty set")
	}
	return lsId, nil
}

func getFollowerAddr(ctx context.Context, LsId uint64) (string, error) {
	rows, err := test.GlobalDB.QueryContext(ctx, getFollowerAddrSql, LsId)
	if err != nil {
		return "", errors.WithMessagef(err, "sql query, sql:%s", getFollowerAddrSql)
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		addr string
	)
	addr = ""

	isEmpty := true
	for rows.Next() {
		if isEmpty {
			isEmpty = false
		}

		err := rows.Scan(&addr)
		if err != nil {
			return "", errors.WithMessagef(err, "scan row")
		}
	}
	if isEmpty {
		return "", errors.Errorf("empty set")
	}
	return addr, nil
}

func switch2Follwer(ctx context.Context, LsId uint64, addr string) error {
	_, err := test.GlobalDB.Exec(fmt.Sprintf(switch2FollwerSql, LsId, addr))
	if err != nil {
		return errors.WithMessagef(err, "sql query, sql:%s", switch2FollwerSql)
	}
	return nil
}

func SwitchReplicaLeaderRandomly4x(tenantName string, databaseName string, tableName string) error {
	// select tenant_id from dba_ob_tenants where status = 'NORMAL' and tenant_name = 'sys';
	// Since we use sys tenant, the tenant_id is 1.

	// get lsId
	lsId, err := getLsId(context.Background(), databaseName, tableName)
	if err != nil {
		return errors.WithMessagef(err, "get lsId")
	}

	// get follower addr
	addr, err := getFollowerAddr(context.Background(), lsId)
	if err != nil {
		return errors.WithMessagef(err, "get follower addr")
	}

	// switch leader to follower
	err = switch2Follwer(context.Background(), lsId, addr)
	if err != nil {
		return errors.WithMessagef(err, "switch leader to follower")
	}

	return nil
}
