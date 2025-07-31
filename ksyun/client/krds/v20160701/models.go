package v20160701

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type ModifyDBParameterGroupParameters struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateSecurityGroupSecurityGroupRule struct {
	SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol,omitempty" name:"SecurityGroupRuleProtocol"`
	SecurityGroupRuleName     *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}
type RestoreToCurInstanceSrcDatabases struct {
	DatabaseName  *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	WholeDatabase *string `json:"WholeDatabase,omitempty" name:"WholeDatabase"`
	TableNames    *string `json:"TableNames,omitempty" name:"TableNames"`
}
type RestoreToCurInstanceDstDatabases struct {
	TableNames    []*string `json:"TableNames,omitempty" name:"TableNames"`
	WholeDatabase *string   `json:"WholeDatabase,omitempty" name:"WholeDatabase"`
	DatabaseName  *string   `json:"DatabaseName,omitempty" name:"DatabaseName"`
}
type RestoreToSgInstanceSrcDatabases struct {
	DatabaseName  *string   `json:"DatabaseName,omitempty" name:"DatabaseName"`
	WholeDatabase *string   `json:"WholeDatabase,omitempty" name:"WholeDatabase"`
	TableNames    []*string `json:"TableNames,omitempty" name:"TableNames"`
}
type RestoreToSgInstanceDstDatabases struct {
	DatabaseName  *string   `json:"DatabaseName,omitempty" name:"DatabaseName"`
	WholeDatabase *string   `json:"WholeDatabase,omitempty" name:"WholeDatabase"`
	TableNames    []*string `json:"TableNames,omitempty" name:"TableNames"`
}
type ModifyInstanceAccountPrivilegesInstanceAccountPrivileges struct {
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Privilege            *string `json:"Privilege,omitempty" name:"Privilege"`
}
type ModifyInstanceDatabasePrivilegesInstanceDatabasePrivileges struct {
	InstanceAccountName *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	Privilege           *string `json:"Privilege,omitempty" name:"Privilege"`
}
type CreateInstanceAccountActionInstanceAccountPrivileges struct {
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Privilege            *string `json:"Privilege,omitempty" name:"Privilege"`
}
type ModifyInstanceAccountPrivilegesActionInstanceAccountPrivileges struct {
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Privilege            *string `json:"Privilege,omitempty" name:"Privilege"`
}
type SetUpProxyInstanceReadOnlyInstanceList struct {
	Id     *string `json:"Id,omitempty" name:"Id"`
	Weight *int    `json:"Weight,omitempty" name:"Weight"`
}
type ModifyInstanceDatabasePrivilegesActionInstanceDatabasePrivileges struct {
	InstanceAccountName *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	Privilege           *string `json:"Privilege,omitempty" name:"Privilege"`
}

type RebootDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *RebootDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RebootDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerEip                         *string  `json:"InnerEip" name:"InnerEip"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RebootDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId   *string                             `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	Parameters           []*ModifyDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
	DBParameterGroupName *string                             `json:"DBParameterGroupName,omitempty" name:"DBParameterGroupName"`
	Description          *string                             `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Description          *string `json:"Description" name:"Description"`
			Parameters           struct {
				DelayKeyWrite               *string  `json:"DelayKeyWrite" name:"DelayKeyWrite"`
				TxIsolation                 *string  `json:"TxIsolation" name:"TxIsolation"`
				InnodbPrintAllDeadlocks     *string  `json:"InnodbPrintAllDeadlocks" name:"InnodbPrintAllDeadlocks"`
				InnodbStatsMethod           *string  `json:"InnodbStatsMethod" name:"InnodbStatsMethod"`
				InnodbStrictMode            *string  `json:"InnodbStrictMode" name:"InnodbStrictMode"`
				DefaultTimeZone             *string  `json:"DefaultTimeZone" name:"DefaultTimeZone"`
				PerformanceSchema           *string  `json:"PerformanceSchema" name:"PerformanceSchema"`
				RplSemiSyncSlaveEnabled     *string  `json:"RplSemiSyncSlaveEnabled" name:"RplSemiSyncSlaveEnabled"`
				InnodbRollbackOnTimeout     *string  `json:"InnodbRollbackOnTimeout" name:"InnodbRollbackOnTimeout"`
				OldAlterTable               *string  `json:"OldAlterTable" name:"OldAlterTable"`
				BinlogRowImage              *string  `json:"BinlogRowImage" name:"BinlogRowImage"`
				QueryCacheType              *string  `json:"QueryCacheType" name:"QueryCacheType"`
				LocalInfile                 *string  `json:"LocalInfile" name:"LocalInfile"`
				InitConnect                 *string  `json:"InitConnect" name:"InitConnect"`
				BinlogFormat                *string  `json:"BinlogFormat" name:"BinlogFormat"`
				LogSlaveUpdates             *string  `json:"LogSlaveUpdates" name:"LogSlaveUpdates"`
				InnodbTableLocks            *string  `json:"InnodbTableLocks" name:"InnodbTableLocks"`
				LowPriorityUpdates          *string  `json:"LowPriorityUpdates" name:"LowPriorityUpdates"`
				ConcurrentInsert            *string  `json:"ConcurrentInsert" name:"ConcurrentInsert"`
				IntQueryTime                *float64 `json:"IntQueryTime" name:"IntQueryTime"`
				SlowQueryLog                *string  `json:"SlowQueryLog" name:"SlowQueryLog"`
				RplSemiSyncMasterEnabled    *string  `json:"RplSemiSyncMasterEnabled" name:"RplSemiSyncMasterEnabled"`
				CharacterSetServer          *string  `json:"CharacterSetServer" name:"CharacterSetServer"`
				LogSlowAdminStatements      *string  `json:"LogSlowAdminStatements" name:"LogSlowAdminStatements"`
				LogBinTrustFunctionCreators *string  `json:"LogBinTrustFunctionCreators" name:"LogBinTrustFunctionCreators"`
				LogQueriesNotUsingIndexes   *string  `json:"LogQueriesNotUsingIndexes" name:"LogQueriesNotUsingIndexes"`
				InnodbStatsOnMetadata       *string  `json:"InnodbStatsOnMetadata" name:"InnodbStatsOnMetadata"`
			} `json:"Parameters"`
			Engine *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroup" name:"DBParameterGroup"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
}

func (r *ResetDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResetDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Description          *string `json:"Description" name:"Description"`
			Parameters           struct {
				DelayKeyWrite               *string `json:"DelayKeyWrite" name:"DelayKeyWrite"`
				TxIsolation                 *string `json:"TxIsolation" name:"TxIsolation"`
				InnodbPrintAllDeadlocks     *string `json:"InnodbPrintAllDeadlocks" name:"InnodbPrintAllDeadlocks"`
				InnodbStatsMethod           *string `json:"InnodbStatsMethod" name:"InnodbStatsMethod"`
				InnodbStrictMode            *string `json:"InnodbStrictMode" name:"InnodbStrictMode"`
				DefaultTimeZone             *string `json:"DefaultTimeZone" name:"DefaultTimeZone"`
				PerformanceSchema           *string `json:"PerformanceSchema" name:"PerformanceSchema"`
				RplSemiSyncSlaveEnabled     *string `json:"RplSemiSyncSlaveEnabled" name:"RplSemiSyncSlaveEnabled"`
				InnodbRollbackOnTimeout     *string `json:"InnodbRollbackOnTimeout" name:"InnodbRollbackOnTimeout"`
				OldAlterTable               *string `json:"OldAlterTable" name:"OldAlterTable"`
				BinlogRowImage              *string `json:"BinlogRowImage" name:"BinlogRowImage"`
				QueryCacheType              *string `json:"QueryCacheType" name:"QueryCacheType"`
				LocalInfile                 *string `json:"LocalInfile" name:"LocalInfile"`
				InitConnect                 *string `json:"InitConnect" name:"InitConnect"`
				BinlogFormat                *string `json:"BinlogFormat" name:"BinlogFormat"`
				LogSlaveUpdates             *string `json:"LogSlaveUpdates" name:"LogSlaveUpdates"`
				InnodbTableLocks            *string `json:"InnodbTableLocks" name:"InnodbTableLocks"`
				LowPriorityUpdates          *string `json:"LowPriorityUpdates" name:"LowPriorityUpdates"`
				ConcurrentInsert            *string `json:"ConcurrentInsert" name:"ConcurrentInsert"`
				SlowQueryLog                *string `json:"SlowQueryLog" name:"SlowQueryLog"`
				RplSemiSyncMasterEnabled    *string `json:"RplSemiSyncMasterEnabled" name:"RplSemiSyncMasterEnabled"`
				CharacterSetServer          *string `json:"CharacterSetServer" name:"CharacterSetServer"`
				LogSlowAdminStatements      *string `json:"LogSlowAdminStatements" name:"LogSlowAdminStatements"`
				LogBinTrustFunctionCreators *string `json:"LogBinTrustFunctionCreators" name:"LogBinTrustFunctionCreators"`
				LogQueriesNotUsingIndexes   *string `json:"LogQueriesNotUsingIndexes" name:"LogQueriesNotUsingIndexes"`
				InnodbStatsOnMetadata       *string `json:"InnodbStatsOnMetadata" name:"InnodbStatsOnMetadata"`
			} `json:"Parameters"`
			Engine *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroup" name:"DBParameterGroup"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ResetDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	Keyword            *string `json:"Keyword,omitempty" name:"Keyword"`
	Marker             *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords         *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
	Source             *string `json:"Source,omitempty" name:"Source"`
}

func (r *DescribeDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroups []struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Engine               *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroups" name:"DBParameterGroups"`
		Source     *string `json:"Source" name:"Source"`
		MaxRecords *int    `json:"MaxRecords" name:"MaxRecords"`
		TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineDefaultParametersRequest struct {
	*ksyunhttp.BaseRequest
	Engine        *string `json:"Engine,omitempty" name:"Engine"`
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *DescribeEngineDefaultParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEngineDefaultParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engine        *string `json:"Engine" name:"Engine"`
		EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
		Parameters    struct {
			ConnectTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"ConnectTimeout"`
			DelayKeyWrite struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"DelayKeyWrite"`
			InnodbPurgeBatchSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbPurgeBatchSize"`
			MyisamSortBufferSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"MyisamSortBufferSize"`
			BulkInsertBufferSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"BulkInsertBufferSize"`
			DivPrecisionIncrement struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"DivPrecisionIncrement"`
			InnodbConcurrencyTickets struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbConcurrencyTickets"`
			MaxPreparedStmtCount struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"MaxPreparedStmtCount"`
			WaitTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"WaitTimeout"`
			TxIsolation struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Alias           *string   `json:"Alias" name:"Alias"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"TxIsolation"`
			TableDefinitionCache struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"TableDefinitionCache"`
			AutoIncrementIncrement struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"AutoIncrementIncrement"`
			FtQueryExpansionLimit struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"FtQueryExpansionLimit"`
			InnodbOldBlocksTime struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbOldBlocksTime"`
			InnodbStatsSamplePages struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbStatsSamplePages"`
			InnodbPrintAllDeadlocks struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InnodbPrintAllDeadlocks"`
			SyncBinlog struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"SyncBinlog"`
			InnodbStatsMethod struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InnodbStatsMethod"`
			LockWaitTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"LockWaitTimeout"`
			NetReadTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"NetReadTimeout"`
			QueryPreallocSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"QueryPreallocSize"`
			BackLog struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"BackLog"`
			MaxErrorCount struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"MaxErrorCount"`
			KeyCacheAgeThreshold struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"KeyCacheAgeThreshold"`
			InnodbLogFileSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbLogFileSize"`
			InnodbThreadConcurrency struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbThreadConcurrency"`
			InnodbStrictMode struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InnodbStrictMode"`
			InnodbFlushLogAtTrxCommit struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbFlushLogAtTrxCommit"`
			DefaultTimeZone struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"DefaultTimeZone"`
			PerformanceSchema struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"PerformanceSchema"`
			InnodbWriteIoThreads struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbWriteIoThreads"`
			RplSemiSyncMasterTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"RplSemiSyncMasterTimeout"`
			MaxConnectErrors struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"MaxConnectErrors"`
			JoinBufferSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"JoinBufferSize"`
			RplSemiSyncSlaveEnabled struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"RplSemiSyncSlaveEnabled"`
			InnodbRollbackOnTimeout struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InnodbRollbackOnTimeout"`
			OldAlterTable struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"OldAlterTable"`
			BinlogRowImage struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"BinlogRowImage"`
			KeyCacheBlockSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"KeyCacheBlockSize"`
			QueryCacheType struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Ignore          *bool     `json:"Ignore" name:"Ignore"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"QueryCacheType"`
			InitConnect struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InitConnect"`
			LocalInfile struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"LocalInfile"`
			BinlogFormat struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"BinlogFormat"`
			LogSlaveUpdates struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"LogSlaveUpdates"`
			SlowLaunchTime struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"SlowLaunchTime"`
			InnodbTableLocks struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InnodbTableLocks"`
			NetWriteTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"NetWriteTimeout"`
			QueryAllocBlockSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"QueryAllocBlockSize"`
			LowerCaseTableNames struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"LowerCaseTableNames"`
			TmpTableSize struct {
				Default            *string `json:"Default" name:"Default"`
				Max                *string `json:"Max" name:"Max"`
				RestartRequired    *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type               *string `json:"Type" name:"Type"`
				DefaultScaleFactor *string `json:"DefaultScaleFactor" name:"DefaultScaleFactor"`
				MaxScaleFactor     *string `json:"MaxScaleFactor" name:"MaxScaleFactor"`
				Variable           *string `json:"Variable" name:"Variable"`
			} `json:"TmpTableSize"`
			DefaultWeekFormat struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"DefaultWeekFormat"`
			KeyCacheDivisionLimit struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"KeyCacheDivisionLimit"`
			InnodbLockWaitTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbLockWaitTimeout"`
			DelayedInsertTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"DelayedInsertTimeout"`
			NetRetryCount struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"NetRetryCount"`
			InnodbPurgeThreads struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbPurgeThreads"`
			BinlogCacheSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"BinlogCacheSize"`
			LowPriorityUpdates struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"LowPriorityUpdates"`
			AutoIncrementOffset struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"AutoIncrementOffset"`
			InnodbMaxDirtyPagesPct struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbMaxDirtyPagesPct"`
			InnodbReadAheadThreshold struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbReadAheadThreshold"`
			QueryCacheLimit struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"QueryCacheLimit"`
			FtMinWordLen struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"FtMinWordLen"`
			ConcurrentInsert struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"ConcurrentInsert"`
			IntQueryTime struct {
				Default         *float64 `json:"Default" name:"Default"`
				RestartRequired *bool    `json:"RestartRequired" name:"RestartRequired"`
				Type            *string  `json:"Type" name:"Type"`
			} `json:"IntQueryTime"`
			SlowQueryLog struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"SlowQueryLog"`
			SortBufferSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"SortBufferSize"`
			InteractiveTimeout struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InteractiveTimeout"`
			QueryCacheSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"QueryCacheSize"`
			InnodbReadIoThreads struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbReadIoThreads"`
			MaxAllowedPacket struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"MaxAllowedPacket"`
			RplSemiSyncMasterEnabled struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"RplSemiSyncMasterEnabled"`
			DelayedInsertLimit struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"DelayedInsertLimit"`
			InnodbOpenFiles struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbOpenFiles"`
			CharacterSetServer struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"CharacterSetServer"`
			DelayedQueueSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"DelayedQueueSize"`
			MaxUserConnections struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"MaxUserConnections"`
			InnodbOldBlocksPct struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbOldBlocksPct"`
			TableOpenCache struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"TableOpenCache"`
			LogSlowAdminStatements struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"LogSlowAdminStatements"`
			LogBinTrustFunctionCreators struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"LogBinTrustFunctionCreators"`
			LogQueriesNotUsingIndexes struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"LogQueriesNotUsingIndexes"`
			InnodbStatsOnMetadata struct {
				Default         *string   `json:"Default" name:"Default"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				Type            *string   `json:"Type" name:"Type"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"InnodbStatsOnMetadata"`
			TableOpenCacheInstances struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"TableOpenCacheInstances"`
			GroupConcatMaxLen struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"GroupConcatMaxLen"`
			InnodbSortBufferSize struct {
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				Type            *string `json:"Type" name:"Type"`
			} `json:"InnodbSortBufferSize"`
		} `json:"Parameters" name:"Parameters"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeEngineDefaultParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineDefaultParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupName *string   `json:"DBParameterGroupName,omitempty" name:"DBParameterGroupName"`
	Engine               *string   `json:"Engine,omitempty" name:"Engine"`
	EngineVersion        *string   `json:"EngineVersion,omitempty" name:"EngineVersion"`
	ParametersName       []*string `json:"Parameters.Name,omitempty" name:"Parameters.Name"`
	ParametersValue      []*string `json:"Parameters.Value,omitempty" name:"Parameters.Value"`
	Description          *string   `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Parameters           struct {
			} `json:"Parameters"`
			Engine *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroup" name:"DBParameterGroup"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
}

func (r *DeleteDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	Mem                        *int      `json:"Mem,omitempty" name:"Mem"`
	Disk                       *int      `json:"Disk,omitempty" name:"Disk"`
	DBInstanceName             *string   `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	Engine                     *string   `json:"Engine,omitempty" name:"Engine"`
	EngineVersion              *string   `json:"EngineVersion,omitempty" name:"EngineVersion"`
	MasterUserPassword         *string   `json:"MasterUserPassword,omitempty" name:"MasterUserPassword"`
	MasterUserName             *string   `json:"MasterUserName,omitempty" name:"MasterUserName"`
	DBInstanceType             *string   `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	VpcId                      *string   `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId                   *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	PreferredBackupTime        *string   `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	DBParameterGroupId         *string   `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	SecurityGroupId            *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Port                       *string   `json:"Port,omitempty" name:"Port"`
	BillType                   *string   `json:"BillType,omitempty" name:"BillType"`
	Duration                   *string   `json:"Duration,omitempty" name:"Duration"`
	DurationUnit               *string   `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone           []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ProjectId                  *int      `json:"ProjectId,omitempty" name:"ProjectId"`
	TableNamesAreCaseSensitive *int      `json:"TableNamesAreCaseSensitive,omitempty" name:"TableNamesAreCaseSensitive"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId" name:"GroupId"`
			Vip                   *string `json:"Vip" name:"Vip"`
			Engine                *string `json:"Engine" name:"Engine"`
			EngineVersion         *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName        *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId" name:"VpcId"`
			SubnetId              *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType              *string `json:"BillType" name:"BillType"`
			OrderType             *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList  struct {
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDBInstanceFromDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceName     *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	DBBackupIdentifier *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	DBInstanceType     *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	ProjectId          *string `json:"ProjectId,omitempty" name:"ProjectId"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Duration           *int    `json:"Duration,omitempty" name:"Duration"`
	DurationUnit       *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	Port               *int    `json:"Port,omitempty" name:"Port"`
	BillType           *string `json:"BillType,omitempty" name:"BillType"`
}

func (r *RestoreDBInstanceFromDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestoreDBInstanceFromDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId" name:"GroupId"`
			Vip                   *string `json:"Vip" name:"Vip"`
			Engine                *string `json:"Engine" name:"Engine"`
			EngineVersion         *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			DatastoreVersionId    *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId" name:"VpcId"`
			SubnetId              *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType              *string `json:"BillType" name:"BillType"`
			OrderType             *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList  struct {
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
				DBBackupName         *string `json:"DBBackupName" name:"DBBackupName"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreDBInstanceFromDBBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreDBInstanceFromDBBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DeleteDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceReadReplicaRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	AttachedVipId        *string `json:"AttachedVipId,omitempty" name:"AttachedVipId"`
	BillType             *string `json:"BillType,omitempty" name:"BillType"`
	Duration             *string `json:"Duration,omitempty" name:"Duration"`
	DurationUnit         *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone1    *string `json:"AvailabilityZone.1,omitempty" name:"AvailabilityZone.1"`
	ProjectId            *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	Vip                  *string `json:"Vip,omitempty" name:"Vip"`
	Mem                  *int    `json:"Mem,omitempty" name:"Mem"`
	Disk                 *int    `json:"Disk,omitempty" name:"Disk"`
}

func (r *CreateDBInstanceReadReplicaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDBInstanceReadReplicaResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId" name:"GroupId"`
			Vip                   *string `json:"Vip" name:"Vip"`
			Engine                *string `json:"Engine" name:"Engine"`
			EngineVersion         *string `json:"EngineVersion" name:"EngineVersion"`
			MasterUserName        *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId" name:"VpcId"`
			SubnetId              *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType              *string `json:"BillType" name:"BillType"`
			OrderType             *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList  struct {
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"DBSource"`
			ProductId      *string `json:"ProductId" name:"ProductId"`
			Region         *string `json:"Region" name:"Region"`
			SubOrderId     *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBInstanceReadReplicaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBInstanceReadReplicaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDBInstanceToPointInTimeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	RestorableTime       *string `json:"RestorableTime,omitempty" name:"RestorableTime"`
}

func (r *RestoreDBInstanceToPointInTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestoreDBInstanceToPointInTimeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreDBInstanceToPointInTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreDBInstanceToPointInTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRestorableTimeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBInstanceRestorableTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBInstanceRestorableTimeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RestorableTime struct {
			Begin *string `json:"Begin" name:"Begin"`
			End   *string `json:"End" name:"End"`
		} `json:"RestorableTime" name:"RestorableTime"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstanceRestorableTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRestorableTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	DBParameterGroupId   *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
}

func (r *ModifyDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBLogType            *string `json:"DBLogType,omitempty" name:"DBLogType"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBLogFilesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DescribeDBLogFiles []struct {
			LogFileName *string  `json:"LogFileName" name:"LogFileName"`
			Size        *float64 `json:"Size" name:"Size"`
			RawSize     *float64 `json:"RawSize" name:"RawSize"`
			Date        *string  `json:"Date" name:"Date"`
			StartTime   *string  `json:"StartTime" name:"StartTime"`
			EndTime     *string  `json:"EndTime" name:"EndTime"`
		} `json:"DescribeDBLogFiles" name:"DescribeDBLogFiles"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	BackupType           *string `json:"BackupType,omitempty" name:"BackupType"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
	Marker               *string `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBBackupsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBBackup []struct {
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
			Engine               *string `json:"Engine" name:"Engine"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			BackupCreateTime     *string `json:"BackupCreateTime" name:"BackupCreateTime"`
			BackupUpdatedTime    *string `json:"BackupUpdatedTime" name:"BackupUpdatedTime"`
			DBBackupName         *string `json:"DBBackupName" name:"DBBackupName"`
			BackupMode           *string `json:"BackupMode" name:"BackupMode"`
			BackupType           *string `json:"BackupType" name:"BackupType"`
			Status               *string `json:"Status" name:"Status"`
			BackupSize           *int    `json:"BackupSize" name:"BackupSize"`
			BackupLocationRef    *string `json:"BackupLocationRef" name:"BackupLocationRef"`
			RemotePath           *string `json:"RemotePath" name:"RemotePath"`
			MD5                  *string `json:"MD5" name:"MD5"`
		} `json:"DBBackup" name:"DBBackup"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Mem                  *int    `json:"Mem,omitempty" name:"Mem"`
	Disk                 *int    `json:"Disk,omitempty" name:"Disk"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBInstanceSpecResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier   *string   `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBInstanceType         *string   `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	DBInstanceStatus       *string   `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`
	Keyword                *string   `json:"Keyword,omitempty" name:"Keyword"`
	Order                  *string   `json:"Order,omitempty" name:"Order"`
	ProjectId              *string   `json:"ProjectId,omitempty" name:"ProjectId"`
	Marker                 *int      `json:"Marker,omitempty" name:"Marker"`
	MaxRecords             *int      `json:"MaxRecords,omitempty" name:"MaxRecords"`
	DBInstanceIdentifierIn []*string `json:"DBInstanceIdentifierIn,omitempty" name:"DBInstanceIdentifierIn"`
	DBInstanceNameIn       []*string `json:"DBInstanceNameIn,omitempty" name:"DBInstanceNameIn"`
	VipIn                  []*string `json:"VipIn,omitempty" name:"VipIn"`
	EIPIn                  []*string `json:"EIPIn,omitempty" name:"EIPIn"`
	ExpiryDateLessThan     *int      `json:"ExpiryDateLessThan,omitempty" name:"ExpiryDateLessThan"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			MiniVersion                      *string   `json:"MiniVersion" name:"MiniVersion"`
			SecurityGroups                   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			NetworkType   *int  `json:"NetworkType" name:"NetworkType"`
			SupportIPV6   *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BindInstances []struct {
				Status *string `json:"Status" name:"Status"`
				Weight *int    `json:"Weight" name:"Weight"`
				Name   *string `json:"Name" name:"Name"`
				Vpc    *string `json:"Vpc" name:"Vpc"`
				Type   *string `json:"Type" name:"Type"`
				Id     *string `json:"Id" name:"Id"`
			} `json:"BindInstances"`
			ProxyNodeInfo []struct {
				Name *string `json:"Name" name:"Name"`
				Id   *string `json:"Id" name:"Id"`
			} `json:"ProxyNodeInfo"`
			ProxyInfo struct {
			} `json:"ProxyInfo"`
			AutoSwitch *int `json:"AutoSwitch" name:"AutoSwitch"`
			BillTypeId *int `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverrideDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
}

func (r *OverrideDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OverrideDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *OverrideDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverrideDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEngineVersionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDBEngineVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBEngineVersionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engines struct {
			MySQL []struct {
				Engine        *string `json:"Engine" name:"Engine"`
				EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
			} `json:"MySQL"`
			ConsistentMysql []struct {
				Engine        *string `json:"Engine" name:"Engine"`
				EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
			} `json:"ConsistentMysql"`
			Percona []struct {
				Engine        *string `json:"Engine" name:"Engine"`
				EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
			} `json:"Percona"`
			EbsMysql []struct {
				Engine        *string `json:"Engine" name:"Engine"`
				EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
			} `json:"EbsMysql"`
		} `json:"Engines" name:"Engines"`
	} `json:"Data"`
}

func (r *DescribeDBEngineVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEngineVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Engine               *string `json:"Engine,omitempty" name:"Engine"`
	EngineVersion        *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *UpgradeDBInstanceEngineVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpgradeDBInstanceEngineVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeDBInstanceEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceTypeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBInstanceType       *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	BillType             *string `json:"BillType,omitempty" name:"BillType"`
	Duration             *int    `json:"Duration,omitempty" name:"Duration"`
	DurationUnit         *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone     *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
}

func (r *ModifyDBInstanceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBInstanceTypeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId       *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                   *string `json:"Vip" name:"Vip"`
			Engine                *string `json:"Engine" name:"Engine"`
			EngineVersion         *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName        *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId" name:"VpcId"`
			SubnetId              *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType              *string `json:"BillType" name:"BillType"`
			OrderType             *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZone      *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			AvailabilityZoneList  []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
				PointInTime          *string `json:"PointInTime" name:"PointInTime"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBInstanceParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBInstanceParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engine        *string `json:"Engine" name:"Engine"`
		EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
		Parameters    struct {
			DelayKeyWrite               *string  `json:"delay_key_write" name:"delay_key_write"`
			TxIsolation                 *string  `json:"tx_isolation" name:"tx_isolation"`
			InnodbPrintAllDeadlocks     *string  `json:"innodb_print_all_deadlocks" name:"innodb_print_all_deadlocks"`
			InnodbStatsMethod           *string  `json:"innodb_stats_method" name:"innodb_stats_method"`
			InnodbStrictMode            *string  `json:"innodb_strict_mode" name:"innodb_strict_mode"`
			DefaultTimeZone             *string  `json:"default_time_zone" name:"default_time_zone"`
			PerformanceSchema           *string  `json:"performance_schema" name:"performance_schema"`
			RplSemiSyncSlaveEnabled     *string  `json:"rpl_semi_sync_slave_enabled" name:"rpl_semi_sync_slave_enabled"`
			InnodbRollbackOnTimeout     *string  `json:"innodb_rollback_on_timeout" name:"innodb_rollback_on_timeout"`
			OldAlterTable               *string  `json:"old_alter_table" name:"old_alter_table"`
			BinlogRowImage              *string  `json:"binlog_row_image" name:"binlog_row_image"`
			QueryCacheType              *string  `json:"query_cache_type" name:"query_cache_type"`
			LocalInfile                 *string  `json:"local_infile" name:"local_infile"`
			InitConnect                 *string  `json:"init_connect" name:"init_connect"`
			BinlogFormat                *string  `json:"binlog_format" name:"binlog_format"`
			LogSlaveUpdates             *string  `json:"log_slave_updates" name:"log_slave_updates"`
			InnodbTableLocks            *string  `json:"innodb_table_locks" name:"innodb_table_locks"`
			LowPriorityUpdates          *string  `json:"low_priority_updates" name:"low_priority_updates"`
			ConcurrentInsert            *string  `json:"concurrent_insert" name:"concurrent_insert"`
			IntQueryTime                *float64 `json:"Int_query_time" name:"Int_query_time"`
			SlowQueryLog                *string  `json:"slow_query_log" name:"slow_query_log"`
			RplSemiSyncMasterEnabled    *string  `json:"rpl_semi_sync_master_enabled" name:"rpl_semi_sync_master_enabled"`
			CharacterSetServer          *string  `json:"character_set_server" name:"character_set_server"`
			LogSlowAdminStatements      *string  `json:"log_slow_admin_statements" name:"log_slow_admin_statements"`
			LogBinTrustFunctionCreators *string  `json:"log_bin_trust_function_creators" name:"log_bin_trust_function_creators"`
			LogQueriesNotUsingIndexes   *string  `json:"log_queries_not_using_indexes" name:"log_queries_not_using_indexes"`
			InnodbStatsOnMetadata       *string  `json:"innodb_stats_on_metadata" name:"innodb_stats_on_metadata"`
		} `json:"Parameters" name:"Parameters"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstanceParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBBackupPolicyRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier   *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	PreferredBackupTime    *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	ExpireAfter            *int    `json:"ExpireAfter,omitempty" name:"ExpireAfter"`
	IncrementalBackupCycle *string `json:"IncrementalBackupCycle,omitempty" name:"IncrementalBackupCycle"`
	FullBackupCycle        *string `json:"FullBackupCycle,omitempty" name:"FullBackupCycle"`
	BinlogExpireAfter      *int    `json:"BinlogExpireAfter,omitempty" name:"BinlogExpireAfter"`
	HighFrequencyBackup    *bool   `json:"HighFrequencyBackup,omitempty" name:"HighFrequencyBackup"`
}

func (r *ModifyDBBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBBackupPolicyResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupPolicyRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBBackupPolicyResponse struct {
	*ksyunhttp.BaseResponse
	BackupConfig struct {
		ExpireAfter            *int    `json:"ExpireAfter" name:"ExpireAfter"`
		AutobackupAt           *int    `json:"AutobackupAt" name:"AutobackupAt"`
		IncrementalBackupCycle *int    `json:"IncrementalBackupCycle" name:"IncrementalBackupCycle"`
		Duration               *int    `json:"Duration" name:"Duration"`
		FullBackupCycle        *string `json:"FullBackupCycle" name:"FullBackupCycle"`
		GroupId                *string `json:"GroupId" name:"GroupId"`
		BinlogExpireAfter      *int    `json:"BinlogExpireAfter" name:"BinlogExpireAfter"`
		HighFrequencyBackup    *bool   `json:"HighFrequencyBackup" name:"HighFrequencyBackup"`
	} `json:"BackupConfig"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBBackupIdentifier *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
}

func (r *DeleteDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupName         *string `json:"DBBackupName,omitempty" name:"DBBackupName"`
	Description          *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	DBBackup struct {
		DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
		DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
		Engine               *string `json:"Engine" name:"Engine"`
		EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
		BackupCreateTime     *string `json:"BackupCreateTime" name:"BackupCreateTime"`
		BackupUpdatedTime    *string `json:"BackupUpdatedTime" name:"BackupUpdatedTime"`
		DBBackupName         *string `json:"DBBackupName" name:"DBBackupName"`
		BackupType           *string `json:"BackupType" name:"BackupType"`
		Status               *string `json:"Status" name:"Status"`
	} `json:"DBBackup"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceHARequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SwitchDBInstanceHARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SwitchDBInstanceHAResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SwitchDBInstanceHAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDBInstanceHAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateDBAdminURLRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *GenerateDBAdminURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GenerateDBAdminURLResponse struct {
	*ksyunhttp.BaseResponse
	DBAdminURL *string `json:"DBAdminURL" name:"DBAdminURL"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *GenerateDBAdminURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDBAdminURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Port                 *string `json:"Port,omitempty" name:"Port"`
}

func (r *AllocateDBInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AllocateDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
				Vip                             *string `json:"Vip" name:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier" name:"ReadReplicaDBInstanceIdentifier"`
				Id                              *string `json:"Id" name:"Id"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AllocateDBInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateDBInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *ReleaseDBInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReleaseDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
				Vip                             *string `json:"Vip" name:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier" name:"ReadReplicaDBInstanceIdentifier"`
				Id                              *string `json:"Id" name:"Id"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ReleaseDBInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseDBInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceAvailabilityZoneRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AvailabilityZone1    *string `json:"AvailabilityZone.1,omitempty" name:"AvailabilityZone.1"`
	AvailabilityZone2    *string `json:"AvailabilityZone.2,omitempty" name:"AvailabilityZone.2"`
}

func (r *ModifyDBInstanceAvailabilityZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBInstanceAvailabilityZoneResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id" name:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *float64 `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string  `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool    `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
		} `json:"DBInstance" name:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceAvailabilityZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceAvailabilityZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupName        *string                                 `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupRule        []*CreateSecurityGroupSecurityGroupRule `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
	DBInstanceIdentifier     *string                                 `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	SecurityGroupDescription *string                                 `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string   `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr     *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Type            *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupIdList *string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr     *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId          *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName        *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
}

func (r *ModifySecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr     *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId          *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName        *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloneSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr     *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupRuleAction                 *string `json:"SecurityGroupRuleAction,omitempty" name:"SecurityGroupRuleAction"`
	SecurityGroupId                         *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupRuleSecurityGroupRuleIdN   *string `json:"SecurityGroupRule.SecurityGroupRuleId.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleId.N"`
	SecurityGroupRuleSecurityGroupRuleNameN *string `json:"SecurityGroupRule.SecurityGroupRuleName.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleName.N"`
	SecurityGroupRuleSecurityGroupRuleCidrN *string `json:"SecurityGroupRule.SecurityGroupRuleCidr.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleCidr.N"`
}

func (r *ModifySecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr     *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupRelationRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId      *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	RelationAction       *string   `json:"RelationAction,omitempty" name:"RelationAction"`
	DBInstanceIdentifier []*string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SecurityGroupRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SecurityGroupRelationResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr     *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SecurityGroupRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SecurityGroupRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleNameRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupRuleId   *string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
	SecurityGroupRuleName *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}

func (r *ModifySecurityGroupRuleNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySecurityGroupRuleNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupRuleNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRuleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLastLogRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBLogType            *string `json:"DBLogType,omitempty" name:"DBLogType"`
}

func (r *DescribeLastLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeLastLogResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		LogFiles []struct {
			LogFileName *string `json:"LogFileName" name:"LogFileName"`
			Size        *string `json:"Size" name:"Size"`
			DBLogType   *string `json:"DBLogType" name:"DBLogType"`
			Status      *string `json:"Status" name:"Status"`
		} `json:"LogFiles" name:"LogFiles"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeLastLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLastLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartAuditRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *StartAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartAuditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StartAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopAuditRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *StopAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopAuditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StopAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAuditRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AccessSqlStatement   *string `json:"AccessSqlStatement,omitempty" name:"AccessSqlStatement"`
	AccessSqlLanguage    *string `json:"AccessSqlLanguage,omitempty" name:"AccessSqlLanguage"`
	AccessDBName         *string `json:"AccessDBName,omitempty" name:"AccessDBName"`
	SourceIp             *string `json:"SourceIp,omitempty" name:"SourceIp"`
	AccessUsername       *string `json:"AccessUsername,omitempty" name:"AccessUsername"`
	AuditBeginTime       *string `json:"AuditBeginTime,omitempty" name:"AuditBeginTime"`
	AuditEndTime         *string `json:"AuditEndTime,omitempty" name:"AuditEndTime"`
	RunResult            *string `json:"RunResult,omitempty" name:"RunResult"`
	KeyWord              *string `json:"KeyWord,omitempty" name:"KeyWord"`
	SortBy               *string `json:"SortBy,omitempty" name:"SortBy"`
	SortWay              *string `json:"SortWay,omitempty" name:"SortWay"`
	Marker               *string `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *string `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *ListAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListAuditResponse struct {
	*ksyunhttp.BaseResponse
	AuditRows []struct {
		Id                 *string `json:"Id" name:"Id"`
		InstanceId         *string `json:"InstanceId" name:"InstanceId"`
		AccessUsername     *string `json:"AccessUsername" name:"AccessUsername"`
		SourceIp           *string `json:"SourceIp" name:"SourceIp"`
		AccessSqlExt       *string `json:"AccessSqlExt" name:"AccessSqlExt"`
		AccessSqlStatement *string `json:"AccessSqlStatement" name:"AccessSqlStatement"`
		AccessSqlLanguage  *string `json:"AccessSqlLanguage" name:"AccessSqlLanguage"`
		AccessDBName       *string `json:"AccessDBName" name:"AccessDBName"`
		RunResult          *bool   `json:"RunResult" name:"RunResult"`
		Duration           *string `json:"Duration" name:"Duration"`
	} `json:"AuditRows"`
}

func (r *ListAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditStatisticRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier    *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AuditStatisticBeginTime *string `json:"AuditStatisticBeginTime,omitempty" name:"AuditStatisticBeginTime"`
	AuditStatisticEndTime   *string `json:"AuditStatisticEndTime,omitempty" name:"AuditStatisticEndTime"`
}

func (r *AuditStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AuditStatisticResponse struct {
	*ksyunhttp.BaseResponse
	AuditStatistic struct {
		AccessSqlLanguage struct {
		} `json:"AccessSqlLanguage" name:"AccessSqlLanguage"`
		AccessSqlStatement struct {
		} `json:"AccessSqlStatement" name:"AccessSqlStatement"`
	} `json:"AuditStatistic"`
}

func (r *AuditStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCurrentDatabaseInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *GetCurrentDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetCurrentDatabaseInfoResponse struct {
	*ksyunhttp.BaseResponse
	Databases []struct {
		DatabaseName *string   `json:"DatabaseName" name:"DatabaseName"`
		TableNames   []*string `json:"TableNames" name:"TableNames"`
	} `json:"Databases"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetCurrentDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCurrentDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTableRestorableTimeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *GetTableRestorableTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetTableRestorableTimeResponse struct {
	*ksyunhttp.BaseResponse
	RestorableTime struct {
		Begin *string `json:"Begin" name:"Begin"`
		End   *string `json:"End" name:"End"`
	} `json:"RestorableTime"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetTableRestorableTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTableRestorableTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHistoryDatabaseInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	RestorableTime       *string `json:"RestorableTime,omitempty" name:"RestorableTime"`
}

func (r *GetHistoryDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetHistoryDatabaseInfoResponse struct {
	*ksyunhttp.BaseResponse
	Databases []struct {
		DatabaseName *string   `json:"DatabaseName" name:"DatabaseName"`
		TableNames   []*string `json:"TableNames" name:"TableNames"`
	} `json:"Databases"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetHistoryDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHistoryDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverrideDBInstanceByPointInTimeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	RestorableTime       *string `json:"RestorableTime,omitempty" name:"RestorableTime"`
}

func (r *OverrideDBInstanceByPointInTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OverrideDBInstanceByPointInTimeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *OverrideDBInstanceByPointInTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverrideDBInstanceByPointInTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreToCurInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string                             `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string                             `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	RestorableTime       *string                             `json:"RestorableTime,omitempty" name:"RestorableTime"`
	SrcDatabases         []*RestoreToCurInstanceSrcDatabases `json:"SrcDatabases,omitempty" name:"SrcDatabases"`
	DstDatabases         []*RestoreToCurInstanceDstDatabases `json:"DstDatabases,omitempty" name:"DstDatabases"`
}

func (r *RestoreToCurInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestoreToCurInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreToCurInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreToCurInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreToSgInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string                            `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string                            `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	RestorableTime       *string                            `json:"RestorableTime,omitempty" name:"RestorableTime"`
	SrcDatabases         []*RestoreToSgInstanceSrcDatabases `json:"SrcDatabases,omitempty" name:"SrcDatabases"`
	DstDatabases         []*RestoreToSgInstanceDstDatabases `json:"DstDatabases,omitempty" name:"DstDatabases"`
}

func (r *RestoreToSgInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestoreToSgInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestStageId *string `json:"RequestStageId" name:"RequestStageId"`
	RequestId      *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreToSgInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreToSgInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditHotCountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	TimeRange            *string `json:"TimeRange,omitempty" name:"TimeRange"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAuditHotCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAuditHotCountResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		DbName     *string  `json:"DbName" name:"DbName"`
		TableName  *string  `json:"TableName" name:"TableName"`
		CountRatio *float64 `json:"CountRatio" name:"CountRatio"`
	} `json:"Data"`
}

func (r *DescribeAuditHotCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditHotCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditHotDurationRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	TimeRange            *string `json:"TimeRange,omitempty" name:"TimeRange"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAuditHotDurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAuditHotDurationResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		DbName        *string  `json:"DbName" name:"DbName"`
		TableName     *string  `json:"TableName" name:"TableName"`
		Duration      *float64 `json:"Duration" name:"Duration"`
		DurationRatio *float64 `json:"DurationRatio" name:"DurationRatio"`
	} `json:"Data"`
}

func (r *DescribeAuditHotDurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditHotDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SqlAuditReportRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	TimeRange            *string `json:"TimeRange,omitempty" name:"TimeRange"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
	DbName               *string `json:"DbName,omitempty" name:"DbName"`
	SortBy               *string `json:"SortBy,omitempty" name:"SortBy"`
	SortWay              *int    `json:"SortWay,omitempty" name:"SortWay"`
	Page                 *int    `json:"Page,omitempty" name:"Page"`
	Size                 *int    `json:"Size,omitempty" name:"Size"`
}

func (r *SqlAuditReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SqlAuditReportResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		DbName        *string  `json:"DbName" name:"DbName"`
		SqlTemplate   *string  `json:"SqlTemplate" name:"SqlTemplate"`
		CountRatio    *float64 `json:"CountRatio" name:"CountRatio"`
		DurationRatio *float64 `json:"DurationRatio" name:"DurationRatio"`
		DurationAvg   *float64 `json:"DurationAvg" name:"DurationAvg"`
	} `json:"Data"`
}

func (r *SqlAuditReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SqlAuditReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SqlAuditLineChartRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	TimeRange            *string `json:"TimeRange,omitempty" name:"TimeRange"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SqlAuditLineChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SqlAuditLineChartResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		ProductId *string `json:"ProductId" name:"ProductId"`
	} `json:"Data"`
}

func (r *SqlAuditLineChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SqlAuditLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogReportRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
	SortBy               *string `json:"SortBy,omitempty" name:"SortBy"`
	SortWay              *string `json:"SortWay,omitempty" name:"SortWay"`
	HeadKey              *string `json:"HeadKey,omitempty" name:"HeadKey"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *SlowLogReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SlowLogReportResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		ProcessingTime  *int    `json:"ProcessingTime" name:"ProcessingTime"`
		Id              *int    `json:"Id" name:"Id"`
		Checksum        *string `json:"Checksum" name:"Checksum"`
		Fingerprint     *string `json:"Fingerprint" name:"Fingerprint"`
		QueryCount      *int    `json:"QueryCount" name:"QueryCount"`
		QueryTimeAvg    *int    `json:"QueryTimeAvg" name:"QueryTimeAvg"`
		QueryTimeMax    *int    `json:"QueryTimeMax" name:"QueryTimeMax"`
		QueryTimeSum    *int    `json:"QueryTimeSum" name:"QueryTimeSum"`
		LockTimeAvg     *int    `json:"LockTimeAvg" name:"LockTimeAvg"`
		LockTimeMax     *int    `json:"LockTimeMax" name:"LockTimeMax"`
		LockTimeSum     *int    `json:"LockTimeSum" name:"LockTimeSum"`
		RowsExaminedAvg *int    `json:"RowsExaminedAvg" name:"RowsExaminedAvg"`
		RowsExaminedMax *int    `json:"RowsExaminedMax" name:"RowsExaminedMax"`
		RowsExaminedSum *int    `json:"RowsExaminedSum" name:"RowsExaminedSum"`
		RowsSentAvg     *int    `json:"RowsSentAvg" name:"RowsSentAvg"`
		RowsSentMax     *int    `json:"RowsSentMax" name:"RowsSentMax"`
		RowsSentSum     *int    `json:"RowsSentSum" name:"RowsSentSum"`
	} `json:"Data"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Marker     *int `json:"Marker" name:"Marker"`
}

func (r *SlowLogReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SlowLogReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogLineChartRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SlowLogLineChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SlowLogLineChartResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		ReportTime  *int    `json:"ReportTime" name:"ReportTime"`
		ReportCount *int    `json:"ReportCount" name:"ReportCount"`
		Percentage  *string `json:"Percentage" name:"Percentage"`
	} `json:"Data"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	QuerySum   *int `json:"QuerySum" name:"QuerySum"`
}

func (r *SlowLogLineChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SlowLogLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogDetailRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
	SortBy               *string `json:"SortBy,omitempty" name:"SortBy"`
	SortWay              *string `json:"SortWay,omitempty" name:"SortWay"`
	HeadKey              *string `json:"HeadKey,omitempty" name:"HeadKey"`
	FingerPrint          *string `json:"FingerPrint,omitempty" name:"FingerPrint"`
	Checksum             *string `json:"checksum,omitempty" name:"checksum"`
}

func (r *SlowLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SlowLogDetailResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		Id               *string `json:"Id" name:"Id"`
		SqlExecTimeStamp *string `json:"SqlExecTimeStamp" name:"SqlExecTimeStamp"`
		AggTime          *string `json:"AggTime" name:"AggTime"`
		AuthUser         *string `json:"AuthUser" name:"AuthUser"`
		CurrentUser      *string `json:"CurrentUser" name:"CurrentUser"`
		Hostname         *string `json:"Hostname" name:"Hostname"`
		QueryTime        *string `json:"QueryTime" name:"QueryTime"`
		LockTime         *string `json:"LockTime" name:"LockTime"`
		RowsSent         *string `json:"RowsSent" name:"RowsSent"`
		RowsExamined     *string `json:"RowsExamined" name:"RowsExamined"`
		SqlContent       *string `json:"SqlContent" name:"SqlContent"`
		ThreadId         *string `json:"ThreadId" name:"ThreadId"`
		ProductId        *string `json:"ProductId" name:"ProductId"`
		TenantId         *string `json:"TenantId" name:"TenantId"`
		Checksum         *string `json:"Checksum" name:"Checksum"`
		Fingerprint      *string `json:"Fingerprint" name:"Fingerprint"`
	} `json:"Data"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Marker     *int `json:"Marker" name:"Marker"`
}

func (r *SlowLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SlowLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartAuditDetailExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	ExportFileds         *string `json:"ExportFileds,omitempty" name:"ExportFileds"`
	AuditBeginTime       *string `json:"AuditBeginTime,omitempty" name:"AuditBeginTime"`
	AuditEndTime         *string `json:"AuditEndTime,omitempty" name:"AuditEndTime"`
}

func (r *StartAuditDetailExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartAuditDetailExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StartAuditDetailExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartAuditDetailExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAuditDetailExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Marker               *string `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *string `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *ListAuditDetailExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListAuditDetailExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskResps []struct {
		StartTime   *string   `json:"StartTime" name:"StartTime"`
		EndTime     *string   `json:"EndTime" name:"EndTime"`
		S3FileNames []*string `json:"S3FileNames" name:"S3FileNames"`
	} `json:"TaskResps"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAuditDetailExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAuditDetailExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPassword    *string `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeInstanceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceAccountsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Accounts []struct {
			InstanceAccountName        *string `json:"InstanceAccountName" name:"InstanceAccountName"`
			InstanceAccountDescription *string `json:"InstanceAccountDescription" name:"InstanceAccountDescription"`
			Created                    *string `json:"Created" name:"Created"`
			InstanceAccountType        *string `json:"InstanceAccountType" name:"InstanceAccountType"`
			InstanceAccountPrivileges  []struct {
				InstanceDatabaseName *string `json:"InstanceDatabaseName" name:"InstanceDatabaseName"`
				Privilege            *string `json:"Privilege" name:"Privilege"`
			} `json:"InstanceAccountPrivileges"`
		} `json:"Accounts" name:"Accounts"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeInstanceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPassword    *string `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
}

func (r *ModifyInstanceAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceAccountInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountPrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier      *string                                                     `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName       *string                                                     `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPrivileges []*ModifyInstanceAccountPrivilegesInstanceAccountPrivileges `json:"InstanceAccountPrivileges,omitempty" name:"InstanceAccountPrivileges"`
}

func (r *ModifyInstanceAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceAccountPrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName  *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCollationsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeCollationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCollationsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Collations struct {
			Utf8     []*string `json:"Utf8" name:"Utf8"`
			Gbk      []*string `json:"Gbk" name:"Gbk"`
			Latin1   []*string `json:"Latin1" name:"Latin1"`
			Utf8mb4  []*string `json:"Utf8mb4" name:"Utf8mb4"`
			Gb18030  []*string `json:"Gb18030" name:"Gb18030"`
			Koi8u    []*string `json:"Koi8u" name:"Koi8u"`
			Koi8r    []*string `json:"Koi8r" name:"Koi8r"`
			Cp850    []*string `json:"Cp850" name:"Cp850"`
			Macce    []*string `json:"Macce" name:"Macce"`
			Ujis     []*string `json:"Ujis" name:"Ujis"`
			Hebrew   []*string `json:"Hebrew" name:"Hebrew"`
			Cp932    []*string `json:"Cp932" name:"Cp932"`
			Ascii    []*string `json:"Ascii" name:"Ascii"`
			Binary   []*string `json:"Binary" name:"Binary"`
			Sjis     []*string `json:"Sjis" name:"Sjis"`
			Armscii8 []*string `json:"Armscii8" name:"Armscii8"`
			Cp852    []*string `json:"Cp852" name:"Cp852"`
			Keybcs2  []*string `json:"Keybcs2" name:"Keybcs2"`
			Cp866    []*string `json:"Cp866" name:"Cp866"`
			Geostd8  []*string `json:"Geostd8" name:"Geostd8"`
			Cp1257   []*string `json:"Cp1257" name:"Cp1257"`
			Ucs2     []*string `json:"Ucs2" name:"Ucs2"`
			Dec8     []*string `json:"Dec8" name:"Dec8"`
			Cp1250   []*string `json:"Cp1250" name:"Cp1250"`
			Tis620   []*string `json:"Tis620" name:"Tis620"`
			Utf32    []*string `json:"Utf32" name:"Utf32"`
			Latin5   []*string `json:"Latin5" name:"Latin5"`
			Hp8      []*string `json:"Hp8" name:"Hp8"`
			Utf16le  []*string `json:"Utf16le" name:"Utf16le"`
			Latin2   []*string `json:"Latin2" name:"Latin2"`
			Macroman []*string `json:"Macroman" name:"Macroman"`
			Eucjpms  []*string `json:"Eucjpms" name:"Eucjpms"`
			Gb2312   []*string `json:"Gb2312" name:"Gb2312"`
			Cp1256   []*string `json:"Cp1256" name:"Cp1256"`
			Big5     []*string `json:"Big5" name:"Big5"`
			Greek    []*string `json:"Greek" name:"Greek"`
			Euckr    []*string `json:"Euckr" name:"Euckr"`
			Cp1251   []*string `json:"Cp1251" name:"Cp1251"`
			Utf16    []*string `json:"Utf16" name:"Utf16"`
			Swe7     []*string `json:"Swe7" name:"Swe7"`
			Latin7   []*string `json:"Latin7" name:"Latin7"`
		} `json:"Collations" name:"Collations"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeCollationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCollationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier        *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName        *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabaseCollation   *string `json:"InstanceDatabaseCollation,omitempty" name:"InstanceDatabaseCollation"`
	InstanceDatabaseDescription *string `json:"InstanceDatabaseDescription,omitempty" name:"InstanceDatabaseDescription"`
}

func (r *CreateInstanceDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInstanceDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabasePrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string                                                       `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName       *string                                                       `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabasePrivileges []*ModifyInstanceDatabasePrivilegesInstanceDatabasePrivileges `json:"InstanceDatabasePrivileges,omitempty" name:"InstanceDatabasePrivileges"`
}

func (r *ModifyInstanceDatabasePrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceDatabasePrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabasePrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabasePrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDatabasesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeInstanceDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		InstanceDatabases []struct {
			InstanceDatabaseName         *string `json:"InstanceDatabaseName" name:"InstanceDatabaseName"`
			InstanceDatabaseCollation    *string `json:"InstanceDatabaseCollation" name:"InstanceDatabaseCollation"`
			InstanceDatabaseCollationSet *string `json:"InstanceDatabaseCollationSet" name:"InstanceDatabaseCollationSet"`
			InstanceDatabaseDescription  *string `json:"InstanceDatabaseDescription" name:"InstanceDatabaseDescription"`
			InstanceDatabaseStatus       *string `json:"InstanceDatabaseStatus" name:"InstanceDatabaseStatus"`
			InstanceDatabasePrivileges   struct {
			} `json:"InstanceDatabasePrivileges"`
		} `json:"InstanceDatabases" name:"InstanceDatabases"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeInstanceDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabaseInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier        *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName        *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabaseDescription *string `json:"InstanceDatabaseDescription,omitempty" name:"InstanceDatabaseDescription"`
}

func (r *ModifyInstanceDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceDatabaseInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSlowLogDetailExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AuditBeginTime       *string `json:"AuditBeginTime,omitempty" name:"AuditBeginTime"`
	AuditEndTime         *string `json:"AuditEndTime,omitempty" name:"AuditEndTime"`
	AccessSqlStatement   *string `json:"AccessSqlStatement,omitempty" name:"AccessSqlStatement"`
}

func (r *StartSlowLogDetailExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartSlowLogDetailExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StartSlowLogDetailExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSlowLogDetailExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSlowLogDetailExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Marker               *string `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *string `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *ListSlowLogDetailExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListSlowLogDetailExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskResps []struct {
		StartTime    *string   `json:"StartTime" name:"StartTime"`
		EndTime      *string   `json:"EndTime" name:"EndTime"`
		Status       *int      `json:"Status" name:"Status"`
		CreateTime   *int      `json:"CreateTime" name:"CreateTime"`
		RecordNumber *int      `json:"RecordNumber" name:"RecordNumber"`
		S3FileNames  []*string `json:"S3FileNames" name:"S3FileNames"`
	} `json:"TaskResps"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *ListSlowLogDetailExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSlowLogDetailExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountActionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string                                                 `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string                                                 `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPassword    *string                                                 `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string                                                 `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
	InstanceAccountPrivileges  []*CreateInstanceAccountActionInstanceAccountPrivileges `json:"InstanceAccountPrivileges,omitempty" name:"InstanceAccountPrivileges"`
}

func (r *CreateInstanceAccountActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceAccountActionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInstanceAccountActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceAccountActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountPrivilegesActionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier      *string                                                           `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName       *string                                                           `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPrivileges []*ModifyInstanceAccountPrivilegesActionInstanceAccountPrivileges `json:"InstanceAccountPrivileges,omitempty" name:"InstanceAccountPrivileges"`
}

func (r *ModifyInstanceAccountPrivilegesActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceAccountPrivilegesActionResponse struct {
	*ksyunhttp.BaseResponse
	DBInstanceIdentifier      *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
	InstanceAccountName       *string   `json:"InstanceAccountName" name:"InstanceAccountName"`
	InstanceAccountPrivileges []*string `json:"InstanceAccountPrivileges" name:"InstanceAccountPrivileges"`
}

func (r *ModifyInstanceAccountPrivilegesActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountPrivilegesActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceAccountActionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName  *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
}

func (r *DeleteInstanceAccountActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceAccountActionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceAccountActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceAccountActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceDatabaseActionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
}

func (r *DeleteInstanceDatabaseActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceDatabaseActionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceDatabaseActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceDatabaseActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBNetworkRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	VpcId                *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId             *string `json:"SubnetId,omitempty" name:"SubnetId"`
	Vip                  *string `json:"Vip,omitempty" name:"Vip"`
	Port                 *string `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyDBNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBNetworkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceMonitorPeriodRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBInstanceMonitorPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBInstanceMonitorPeriodResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Period *int `json:"Period" name:"Period"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstanceMonitorPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceMonitorPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceMonitorPeriodRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Period               *string `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyDBInstanceMonitorPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBInstanceMonitorPeriodResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceMonitorPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceMonitorPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineParametersModifyHistoryRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	Name               *string `json:"Name,omitempty" name:"Name"`
	MaxRecords         *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
	Marker             *int    `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeEngineParametersModifyHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEngineParametersModifyHistoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"request_id" name:"request_id"`
	History   []struct {
		Id               *string `json:"Id" name:"Id"`
		ConfigurationKey *string `json:"ConfigurationKey" name:"ConfigurationKey"`
		OldValue         *string `json:"OldValue" name:"OldValue"`
		NewValue         *string `json:"NewValue" name:"NewValue"`
		Created          *string `json:"Created" name:"Created"`
	} `json:"History"`
}

func (r *DescribeEngineParametersModifyHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineParametersModifyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchApplyDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBParameterGroupId   *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
}

func (r *BatchApplyDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type BatchApplyDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	SucceededCount *int      `json:"succeededCount" name:"succeededCount"`
	FailedCount    *int      `json:"failedCount" name:"failedCount"`
	Succeeded      []*string `json:"Succeeded" name:"Succeeded"`
	Failed         []struct {
		Result *string `json:"Result" name:"Result"`
		Id     *string `json:"Id" name:"Id"`
		Name   *string `json:"Name" name:"Name"`
	} `json:"Failed"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *BatchApplyDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchApplyDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceLatesVersionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *UpgradeDBInstanceLatesVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpgradeDBInstanceLatesVersionResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			MiniVersion          *string `json:"MiniVersion" name:"MiniVersion"`
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus     *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime  *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId              *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId      *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                  *string `json:"Vip" name:"Vip"`
			Port                 *int    `json:"Port" name:"Port"`
			Engine               *string `json:"Engine" name:"Engine"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpgradeDBInstanceLatesVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeDBInstanceLatesVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeProxyInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeProxyInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			MiniVersion                      *string   `json:"MiniVersion" name:"MiniVersion"`
			SecurityGroups                   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			NetworkType   *int      `json:"NetworkType" name:"NetworkType"`
			SupportIPV6   *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BindInstances []*string `json:"BindInstances" name:"BindInstances"`
			ProxyNodeInfo []*string `json:"ProxyNodeInfo" name:"ProxyNodeInfo"`
			ProxyInfo     struct {
			} `json:"ProxyInfo"`
			AutoSwitch *int `json:"AutoSwitch" name:"AutoSwitch"`
			BillTypeId *int `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeProxyInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProxyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetUpProxyInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string                                   `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	ReadOnlyInstanceList []*SetUpProxyInstanceReadOnlyInstanceList `json:"ReadOnlyInstanceList,omitempty" name:"ReadOnlyInstanceList"`
}

func (r *SetUpProxyInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetUpProxyInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetUpProxyInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUpProxyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemporaryCloseSwitchoverRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AutoSwitch           *bool   `json:"AutoSwitch,omitempty" name:"AutoSwitch"`
	ExpireTime           *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *TemporaryCloseSwitchoverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type TemporaryCloseSwitchoverResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *TemporaryCloseSwitchoverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TemporaryCloseSwitchoverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeBackupOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBackupOverviewResponse struct {
	*ksyunhttp.BaseResponse
	BackupOverview struct {
		Total struct {
			Size        *string `json:"Size" name:"Size"`
			Num         *string `json:"Num" name:"Num"`
			FreeSize    *string `json:"FreeSize" name:"FreeSize"`
			UsedSize    *string `json:"UsedSize" name:"UsedSize"`
			ChargedSize *string `json:"ChargedSize" name:"ChargedSize"`
		} `json:"Total" name:"Total"`
		Backup struct {
			Size             *string `json:"Size" name:"Size"`
			Num              *string `json:"Num" name:"Num"`
			AutoBackupSize   *string `json:"AutoBackupSize" name:"AutoBackupSize"`
			AutoBackupNum    *string `json:"AutoBackupNum" name:"AutoBackupNum"`
			ManualBackupSize *string `json:"ManualBackupSize" name:"ManualBackupSize"`
			ManualBackupNum  *string `json:"ManualBackupNum" name:"ManualBackupNum"`
		} `json:"Backup" name:"Backup"`
		Binlog struct {
			Size *string `json:"Size" name:"Size"`
			Num  *string `json:"Num" name:"Num"`
		} `json:"Binlog" name:"Binlog"`
		Deleted struct {
			Size *string `json:"Size" name:"Size"`
			Num  *string `json:"Num" name:"Num"`
		} `json:"Deleted" name:"Deleted"`
	} `json:"BackupOverview"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticBackupDetailsRequest struct {
	*ksyunhttp.BaseRequest
	DataType   *string `json:"DataType,omitempty" name:"DataType"`
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	Marker     *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeStatisticBackupDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeStatisticBackupDetailsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Backups []struct {
			Name                 *string `json:"Name" name:"Name"`
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			BackupSize           *int    `json:"BackupSize" name:"BackupSize"`
			Type                 *string `json:"Type" name:"Type"`
			Status               *string `json:"Status" name:"Status"`
			DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
			BackupCreateTime     *string `json:"BackupCreateTime" name:"BackupCreateTime"`
			BackupUpdatedTime    *string `json:"BackupUpdatedTime" name:"BackupUpdatedTime"`
			RemotePath           *string `json:"RemotePath" name:"RemotePath"`
			BackupLocationRef    *string `json:"BackupLocationRef" name:"BackupLocationRef"`
			BackupSizeUnit       *string `json:"BackupSizeUnit" name:"BackupSizeUnit"`
		} `json:"Backups" name:"Backups"`
		BinlogTasks []struct {
			BinlogName           *string `json:"BinlogName" name:"BinlogName"`
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			FirstEventCtime      *string `json:"FirstEventCtime" name:"FirstEventCtime"`
			LastEventCtime       *string `json:"LastEventCtime" name:"LastEventCtime"`
			Size                 *int    `json:"Size" name:"Size"`
			SyncStatus           *string `json:"SyncStatus" name:"SyncStatus"`
			LogFileName          *string `json:"LogFileName" name:"LogFileName"`
			SizeUnit             *string `json:"SizeUnit" name:"SizeUnit"`
		} `json:"BinlogTasks" name:"BinlogTasks"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeStatisticBackupDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticBackupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMaintenanceTimeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string  `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	StartTime            *string  `json:"StartTime,omitempty" name:"StartTime"`
	Duration             *float64 `json:"Duration,omitempty" name:"Duration"`
}

func (r *ModifyMaintenanceTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyMaintenanceTimeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyMaintenanceTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMaintenanceTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabasePrivilegesActionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string                                                             `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName       *string                                                             `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabasePrivileges []*ModifyInstanceDatabasePrivilegesActionInstanceDatabasePrivileges `json:"InstanceDatabasePrivileges,omitempty" name:"InstanceDatabasePrivileges"`
}

func (r *ModifyInstanceDatabasePrivilegesActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceDatabasePrivilegesActionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabasePrivilegesActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabasePrivilegesActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
