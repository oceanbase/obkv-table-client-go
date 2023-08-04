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
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	getReplicaSql = "SELECT A.table_id as table_id, A.partition_id as partition_id, A.svr_ip as svr_ip, B.svr_port as svr_port, A.role as role FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ? and partition_id in "
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
