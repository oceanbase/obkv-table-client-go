package route

import (
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
)

// write your true config and create table by sql first
var (
	testClusterName = "test"
	testTenantName  = "mysql"
	testDatabase    = "test"
	testTableName   = "test"
	testUserName    = "root"
	testPassword    = ""
	testIp          = "127.0.0.1"
	testSqlPort     = 41101
	testServerPort  = 41100
	testServerAddr  = ObServerAddr{testIp, testSqlPort, testServerPort}
	testUserAuth    = ObUserAuth{testUserName, testPassword}
)

func TestGetTableEntryFromRemote(t *testing.T) {
	err := GetObVersionFromRemote(&testServerAddr, &testUserAuth)
	if err != nil {
		panic(err)
	}
	fmt.Println("ob cluster version is:", ObVersion)
	key := ObTableEntryKey{
		testClusterName,
		testTenantName,
		testDatabase,
		testTableName,
	}
	InitSql(ObVersion)
	entry, err := GetTableEntryFromRemote(&testServerAddr, &testUserAuth, &key)
	if err != nil {
		panic(err)
	}
	fmt.Println("entry is:", entry.ToString())
}

func TestObUserAuth_ToString(t *testing.T) {
	auth := ObUserAuth{}
	assert.Equal(t, "ObUserAuth{userName:, password:}", auth.ToString())
	auth = ObUserAuth{"testUserName", "testPassword"}
	assert.Equal(t, "ObUserAuth{userName:testUserName, password:testPassword}", auth.ToString())
}

func TestObColumnIndexesPair_ToString(t *testing.T) {
	pair := ObColumnIndexesPair{}
	assert.Equal(t, "ObColumnIndexesPair{"+
		"column:ObColumn{"+
		"columnName:, "+
		"index:0, "+
		"objType:nil, "+
		"collationType:ObCollationType{"+"collationType:csTypeInvalid}, "+
		"refColumnNames:[], "+
		"isGenColumn:false, "+
		"columnExpress:nil}, "+
		"indexes:[]}",
		pair.ToString())
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := protocol.NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair = ObColumnIndexesPair{*column, []int{1, 2, 3}}
	assert.Equal(t, "ObColumnIndexesPair{"+
		"column:ObColumn{"+
		"columnName:testColumnName, "+
		"index:0, "+
		"objType:ObObjType{type:ObTinyIntType}, "+
		"collationType:ObCollationType{collationType:csTypeBinary}, "+
		"refColumnNames:[testColumnName], "+
		"isGenColumn:false, "+
		"columnExpress:nil}, "+
		"indexes:[1, 2, 3]}",
		pair.ToString(),
	)
}

func TestObPartDescCommon_ToString(t *testing.T) {
	comm := ObPartDescCommon{}
	assert.Equal(t, "ObPartDescCommon{"+
		"partFuncType:ObPartFuncType{name:, index:0}, "+
		"partExpr:, "+
		"orderedPartColumnNames:, "+
		"orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[], "+
		"rowKeyElement:{}}",
		comm.ToString(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := protocol.NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := ObColumnIndexesPair{*column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeHashIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []ObColumnIndexesPair{pair}
	partColumns := []protocol.ObColumn{*column}
	rowKeyElement := make(map[string]int, 3)
	rowKeyElement["c1"] = 0
	comm = ObPartDescCommon{partFuncType: partFuncType,
		partExpr:                            partExpr,
		orderedPartColumnNames:              orderedPartColumnNames,
		orderedPartRefColumnRowKeyRelations: orderedPartRefColumnRowKeyRelations,
		partColumns:                         partColumns,
		rowKeyElement:                       rowKeyElement,
	}
	assert.Equal(t, "ObPartDescCommon{"+
		"partFuncType:ObPartFuncType{name:HASH, index:0}, "+
		"partExpr:c1, c2, "+
		"orderedPartColumnNames:c1,c2, "+
		"orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], "+
		"partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:{m[c1]=0}}",
		comm.ToString(),
	)
}

func TestObRangePartDesc_ToString(t *testing.T) {
	desc := ObRangePartDesc{}
	assert.Equal(t, "ObRangePartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:, index:0}, partExpr:, orderedPartColumnNames:, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:{}}, "+
		"orderedCompareColumns:[], "+
		"orderedCompareColumnTypes:[]}",
		desc.ToString(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := protocol.NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := ObColumnIndexesPair{*column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeRangeIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []ObColumnIndexesPair{pair}
	partColumns := []protocol.ObColumn{*column}
	rowKeyElement := make(map[string]int, 3)
	rowKeyElement["c1"] = 0
	comm := ObPartDescCommon{partFuncType: partFuncType,
		partExpr:                            partExpr,
		orderedPartColumnNames:              orderedPartColumnNames,
		orderedPartRefColumnRowKeyRelations: orderedPartRefColumnRowKeyRelations,
		partColumns:                         partColumns,
		rowKeyElement:                       rowKeyElement,
	}
	desc = ObRangePartDesc{
		comm:                      comm,
		orderedCompareColumns:     []protocol.ObColumn{*column, *column},
		orderedCompareColumnTypes: []protocol.ObObjType{objType, objType},
	}
	assert.Equal(t, "ObRangePartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:RANGE, index:3}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:{m[c1]=0}}, "+
		"orderedCompareColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], "+
		"orderedCompareColumnTypes:[ObObjType{type:ObTinyIntType}, ObObjType{type:ObTinyIntType}]}",
		desc.ToString(),
	)
}

func TestObHashPartDesc_ToString(t *testing.T) {
	desc := ObHashPartDesc{}
	assert.Equal(t, "ObHashPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:, index:0}, partExpr:, orderedPartColumnNames:, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:{}}, "+
		"completeWorks:[], "+
		"partSpace:0, "+
		"partNum:0, "+
		"partNameIdMap:{}}",
		desc.ToString(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := protocol.NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := ObColumnIndexesPair{*column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeHashIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []ObColumnIndexesPair{pair}
	partColumns := []protocol.ObColumn{*column}
	rowKeyElement := make(map[string]int, 3)
	rowKeyElement["c1"] = 0
	comm := ObPartDescCommon{partFuncType: partFuncType,
		partExpr:                            partExpr,
		orderedPartColumnNames:              orderedPartColumnNames,
		orderedPartRefColumnRowKeyRelations: orderedPartRefColumnRowKeyRelations,
		partColumns:                         partColumns,
		rowKeyElement:                       rowKeyElement,
	}
	partNameIdMap := make(map[string]int64)
	partNameIdMap["p0"] = 0
	desc = ObHashPartDesc{
		comm:          comm,
		completeWorks: []int64{1, 2, 3},
		partSpace:     0,
		partNum:       10,
		partNameIdMap: partNameIdMap,
	}
	assert.Equal(t, "ObHashPartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH, index:0}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:{m[c1]=0}}, "+
		"completeWorks:[1, 2, 3], "+
		"partSpace:0, "+
		"partNum:10, "+
		"partNameIdMap:{m[p0]=0}}",
		desc.ToString(),
	)
}

func TestObKeyPartDesc_ToString(t *testing.T) {
	desc := ObKeyPartDesc{}
	assert.Equal(t, "ObKeyPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:, index:0}, partExpr:, orderedPartColumnNames:, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:{}}, "+
		"partSpace:0, "+
		"partNum:0, "+
		"partNameIdMap:{}}",
		desc.ToString(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := protocol.NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := ObColumnIndexesPair{*column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeKeyV2Index)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []ObColumnIndexesPair{pair}
	partColumns := []protocol.ObColumn{*column}
	rowKeyElement := make(map[string]int, 3)
	rowKeyElement["c1"] = 0
	comm := ObPartDescCommon{partFuncType: partFuncType,
		partExpr:                            partExpr,
		orderedPartColumnNames:              orderedPartColumnNames,
		orderedPartRefColumnRowKeyRelations: orderedPartRefColumnRowKeyRelations,
		partColumns:                         partColumns,
		rowKeyElement:                       rowKeyElement,
	}
	partNameIdMap := make(map[string]int64)
	partNameIdMap["p0"] = 0
	desc = ObKeyPartDesc{
		comm:          comm,
		partSpace:     0,
		partNum:       10,
		partNameIdMap: partNameIdMap,
	}
	assert.Equal(t, "ObKeyPartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:KEY_V2, index:6}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:{m[c1]=0}}, "+
		"partSpace:0, "+
		"partNum:10, "+
		"partNameIdMap:{m[p0]=0}}",
		desc.ToString(),
	)
}

func TestObPartitionLevel_ToString(t *testing.T) {
	level := ObPartitionLevel{}
	assert.Equal(t, "ObPartitionLevel{name:, index:0}", level.ToString())
	level = newObPartitionLevel(partLevelZeroIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelZero, index:0}", level.ToString())
	level = newObPartitionLevel(partLevelOneIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelOne, index:1}", level.ToString())
	level = newObPartitionLevel(partLevelTwoIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelTwo, index:2}", level.ToString())
	level = newObPartitionLevel(partLevelUnknownIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelUnknown, index:-1}", level.ToString())

}

func TestObPartFuncType_ToString(t *testing.T) {
	part := ObPartFuncType{}
	assert.Equal(t, "ObPartFuncType{name:, index:0}", part.ToString())
	part = newObPartFuncType(partFuncTypeHashIndex)
	assert.Equal(t, "ObPartFuncType{name:HASH, index:0}", part.ToString())
	part = newObPartFuncType(partFuncTypeKeyIndex)
	assert.Equal(t, "ObPartFuncType{name:KEY, index:1}", part.ToString())
	part = newObPartFuncType(partFuncTypeKeyImplIndex)
	assert.Equal(t, "ObPartFuncType{name:KEY_IMPLICIT, index:2}", part.ToString())
	part = newObPartFuncType(partFuncTypeRangeIndex)
	assert.Equal(t, "ObPartFuncType{name:RANGE, index:3}", part.ToString())
	part = newObPartFuncType(partFuncTypeRangeColIndex)
	assert.Equal(t, "ObPartFuncType{name:RANGE_COLUMNS, index:4}", part.ToString())
	part = newObPartFuncType(partFuncTypeListIndex)
	assert.Equal(t, "ObPartFuncType{name:LIST, index:5}", part.ToString())
	part = newObPartFuncType(partFuncTypeKeyV2Index)
	assert.Equal(t, "ObPartFuncType{name:KEY_V2, index:6}", part.ToString())
	part = newObPartFuncType(partFuncTypeListColIndex)
	assert.Equal(t, "ObPartFuncType{name:LIST_COLUMNS, index:7}", part.ToString())
	part = newObPartFuncType(partFuncTypeHashV2Index)
	assert.Equal(t, "ObPartFuncType{name:HASH_V2, index:8}", part.ToString())
	part = newObPartFuncType(partFuncTypeKeyV3Index)
	assert.Equal(t, "ObPartFuncType{name:KEY_V3, index:9}", part.ToString())
	part = newObPartFuncType(partFuncTypeUnknownIndex)
	assert.Equal(t, "ObPartFuncType{name:UNKNOWN, index:-1}", part.ToString())
}

func TestObServerAddr_ToString(t *testing.T) {
	addr := ObServerAddr{}
	assert.Equal(t, "ObServerAddr{ip:, sqlPort:0, svrPort:0}", addr.ToString())
	addr = ObServerAddr{"127.0.0.1", 8080, 1227}
	assert.Equal(t, "ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}", addr.ToString())
}

func TestObServerInfo_ToString(t *testing.T) {
	info := ObServerInfo{}
	assert.Equal(t, "ObServerInfo{stopTime:0, status:}", info.ToString())
	info = ObServerInfo{0, "Active"}
	assert.Equal(t, info.isActive(), true)
	assert.Equal(t, "ObServerInfo{stopTime:0, status:Active}", info.ToString())
}

func TestObServerRole_ToString(t *testing.T) {
	role := ObServerRole{}
	assert.Equal(t, "ObServerRole{name:, index:0}", role.ToString())
	role = newObServerRole(ServerRoleLeaderIndex)
	assert.Equal(t, "ObServerRole{name:LEADER, index:1}", role.ToString())
	role = newObServerRole(ServerRoleFollowerIndex)
	assert.Equal(t, "ObServerRole{name:FOLLOWER, index:2}", role.ToString())
	role = newObServerRole(ServerRoleInvalidIndex)
	assert.Equal(t, "ObServerRole{name:INVALID_ROLE, index:-1}", role.ToString())
}

func TestObReplicaType_ToString(t *testing.T) {
	replica := ObReplicaType{}
	assert.Equal(t, "ObReplicaType{name:, index:0}", replica.ToString())
	replica = newObReplicaType(ReplicaTypeFullIndex)
	assert.Equal(t, "ObReplicaType{name:FULL, index:0}", replica.ToString())
	replica = newObReplicaType(ReplicaTypeLogOnlyIndex)
	assert.Equal(t, "ObReplicaType{name:LOGONLY, index:5}", replica.ToString())
	replica = newObReplicaType(ReplicaTypeReadOnlyIndex)
	assert.Equal(t, "ObReplicaType{name:READONLY, index:16}", replica.ToString())
	replica = newObReplicaType(ReplicaTypeInvalidIndex)
	assert.Equal(t, "ObReplicaType{name:INVALID, index:-1}", replica.ToString())
}

func TestObReplicaLocation_ToString(t *testing.T) {
	replica := ObReplicaLocation{}
	assert.Equal(t, "ObReplicaLocation{"+
		"addr:ObServerAddr{ip:, sqlPort:0, svrPort:0}, "+
		"info:ObServerInfo{stopTime:0, status:}, "+
		"role:ObServerRole{name:, index:0}, "+
		"replicaType:ObReplicaType{name:, index:0}}",
		replica.ToString(),
	)
	replica = ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	assert.Equal(t, "ObReplicaLocation{"+
		"addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, "+
		"info:ObServerInfo{stopTime:0, status:Active}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}",
		replica.ToString(),
	)
}

func TestObTableLocation_ToString(t *testing.T) {
	loc := ObTableLocation{}
	assert.Equal(t, "ObTableLocation{replicaLocations:[]}", loc.ToString())
	replica := ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	loc = ObTableLocation{[]ObReplicaLocation{replica, replica}}
	assert.Equal(t, "ObTableLocation{"+
		"replicaLocations:["+
		"ObReplicaLocation{"+
		"addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, "+
		"info:ObServerInfo{stopTime:0, status:Active}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{"+
		"addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, "+
		"info:ObServerInfo{stopTime:0, status:Active}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}]}", loc.ToString(),
	)
}

func TestObPartitionInfo_ToString(t *testing.T) {
	info := ObPartitionInfo{}
	assert.Equal(t, "ObPartitionInfo{"+
		"level:ObPartitionLevel{name:, index:0}, "+
		"firstPartDesc:nil, "+
		"subPartDesc:nil, "+
		"partColumns:[], "+
		"partTabletIdMap:{}, "+
		"partNameIdMap:{}, "+
		"rowKeyElement:{}}",
		info.ToString(),
	)
	level := newObPartitionLevel(partLevelZeroIndex)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := protocol.NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := ObColumnIndexesPair{*column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeHashIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []ObColumnIndexesPair{pair}
	partColumns := []protocol.ObColumn{*column}
	rowKeyElement := make(map[string]int, 3)
	rowKeyElement["c1"] = 0
	comm := ObPartDescCommon{partFuncType: partFuncType,
		partExpr:                            partExpr,
		orderedPartColumnNames:              orderedPartColumnNames,
		orderedPartRefColumnRowKeyRelations: orderedPartRefColumnRowKeyRelations,
		partColumns:                         partColumns,
		rowKeyElement:                       rowKeyElement,
	}
	partNameIdMap := make(map[string]int64)
	partNameIdMap["p0"] = 0
	desc := ObHashPartDesc{
		comm:          comm,
		completeWorks: []int64{1, 2, 3},
		partSpace:     0,
		partNum:       10,
		partNameIdMap: partNameIdMap,
	}
	partTabletIdMap := make(map[int64]int64)
	partTabletIdMap[0] = 500021
	info = ObPartitionInfo{
		level:           level,
		firstPartDesc:   desc,
		subPartDesc:     desc,
		partColumns:     partColumns,
		partTabletIdMap: partTabletIdMap,
		partNameIdMap:   partNameIdMap,
		rowKeyElement:   rowKeyElement,
	}
	assert.Equal(t, "ObPartitionInfo{"+
		"level:ObPartitionLevel{name:partLevelZero, index:0}, "+
		"firstPartDesc:ObHashPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH, index:0}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:{m[c1]=0}}, completeWorks:[1, 2, 3], partSpace:0, partNum:10, partNameIdMap:{m[p0]=0}}, "+
		"subPartDesc:ObHashPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH, index:0}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:{m[c1]=0}}, completeWorks:[1, 2, 3], partSpace:0, partNum:10, partNameIdMap:{m[p0]=0}}, "+
		"partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:csTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], "+
		"partTabletIdMap:{m[0]=500021}, "+
		"partNameIdMap:{m[p0]=0}, "+
		"rowKeyElement:{m[c1]=0}}",
		info.ToString(),
	)
}

func TestObPartitionLocation_ToString(t *testing.T) {
	loc := ObPartitionLocation{}
	assert.Equal(t, "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:, sqlPort:0, svrPort:0}, info:ObServerInfo{stopTime:0, status:}, role:ObServerRole{name:, index:0}, replicaType:ObReplicaType{name:, index:0}}, "+
		"replicas:[]}",
		loc.ToString(),
	)
	leader := ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	follower := ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	loc = ObPartitionLocation{
		leader,
		[]ObReplicaLocation{follower, follower},
	}
	assert.Equal(t, "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, info:ObServerInfo{stopTime:0, status:Active}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, info:ObServerInfo{stopTime:0, status:Active}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, info:ObServerInfo{stopTime:0, status:Active}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}",
		loc.ToString(),
	)
}

func TestObPartLocationEntry_ToString(t *testing.T) {
	entry := ObPartLocationEntry{}
	assert.Equal(t, "ObPartLocationEntry{partLocations:{}}", entry.ToString())
	loc := ObPartitionLocation{}
	m := make(map[int]*ObPartitionLocation, 1)
	m[0] = &loc
	entry = ObPartLocationEntry{m}
	assert.Equal(t, "ObPartLocationEntry{"+
		"partLocations:{m[0]=ObPartitionLocation{leader:ObReplicaLocation{addr:ObServerAddr{ip:, sqlPort:0, svrPort:0}, info:ObServerInfo{stopTime:0, status:}, role:ObServerRole{name:, index:0}, replicaType:ObReplicaType{name:, index:0}}, replicas:[]}}}",
		entry.ToString(),
	)
}

func TestObTableEntryKey_ToString(t *testing.T) {
	key := ObTableEntryKey{}
	assert.Equal(t, "ObTableEntryKey{clusterName:, tenantNane:, databaseName:, tableName:}", key.ToString())
	key = ObTableEntryKey{
		"testClusterName",
		"testTenantNane",
		"testDatabaseName",
		"testTableName",
	}
	assert.Equal(t, "ObTableEntryKey{"+
		"clusterName:testClusterName, "+
		"tenantNane:testDatabaseName, "+
		"databaseName:testDatabaseName, "+
		"tableName:testTableName}",
		key.ToString(),
	)
}

func TestObTableEntry_ToString(t *testing.T) {
	entry := ObTableEntry{}
	assert.Equal(t, "ObTableEntry{"+
		"tableId:0, "+
		"partNum:0, "+
		"replicaNum:0, "+
		"refreshTimeMills:0, "+
		"tableEntryKey:ObTableEntryKey{clusterName:, tenantNane:, databaseName:, tableName:}, "+
		"partitionInfo:nil, "+
		"tableLocation:nil, "+
		"partitionEntry:nil}",
		entry.ToString(),
	)
}
