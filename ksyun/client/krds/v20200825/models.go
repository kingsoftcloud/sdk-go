package v20200825
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type CreateDBParameterGroupParameters struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyDBParameterGroupParameters struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}


type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupName                       *string   `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupDescription                *string   `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
	SecurityGroupType                       *string   `json:"SecurityGroupType,omitempty" name:"SecurityGroupType"`
	DBInstanceIdentifier                    []*string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	SecurityGroupRuleSecurityGroupRuleNameN *string   `json:"SecurityGroupRule.SecurityGroupRuleName.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleName.N"`
	SecurityGroupRuleSecurityGroupRuleCidrN *string   `json:"SecurityGroupRule.SecurityGroupRuleCidr.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleCidr.N"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			SecurityGroupType        *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created               *string `json:"Created" name:"Created"`
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
	SecurityGroupId   []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupType *string   `json:"SecurityGroupType,omitempty" name:"SecurityGroupType"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSecurityGroupResponse struct {
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
				DBInstanceName *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip            *string `json:"Vip" name:"Vip"`
				Created        *string `json:"Created" name:"Created"`
				DBInstanceType *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created               *string `json:"Created" name:"Created"`
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
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
				DBInstanceName *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip            *string `json:"Vip" name:"Vip"`
				Created        *string `json:"Created" name:"Created"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created               *string `json:"Created" name:"Created"`
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
				DBInstanceName *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip            *string `json:"Vip" name:"Vip"`
				Created        *string `json:"Created" name:"Created"`
				DBInstanceType *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created               *string `json:"Created" name:"Created"`
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
			Instances                struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created               *string `json:"Created" name:"Created"`
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
	SecurityGroupRuleAction                *string   `json:"SecurityGroupRuleAction,omitempty" name:"SecurityGroupRuleAction"`
	SecurityGroupId                        *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupRuleSecurityGroupRuleId   []*string `json:"SecurityGroupRule.SecurityGroupRuleId,omitempty" name:"SecurityGroupRule.SecurityGroupRuleId"`
	SecurityGroupRuleSecurityGroupRuleName []*string `json:"SecurityGroupRule.SecurityGroupRuleName,omitempty" name:"SecurityGroupRule.SecurityGroupRuleName"`
	SecurityGroupRuleSecurityGroupRuleCidr []*string `json:"SecurityGroupRule.SecurityGroupRuleCidr,omitempty" name:"SecurityGroupRule.SecurityGroupRuleCidr"`
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
				DBInstanceName *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip            *string `json:"Vip" name:"Vip"`
				Created        *string `json:"Created" name:"Created"`
				DBInstanceType *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr" name:"SecurityGroupRuleCidr"`
				Created               *string `json:"Created" name:"Created"`
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
	RelationAction       *string `json:"RelationAction,omitempty" name:"RelationAction"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupIds     *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SecurityGroupRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SecurityGroupRelationResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created               *string `json:"Created" name:"Created"`
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
	SecurityGroupId       *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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


type CreateDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupName *string                           `json:"DBParameterGroupName,omitempty" name:"DBParameterGroupName"`
	Engine               *string                           `json:"Engine,omitempty" name:"Engine"`
	EngineVersion        *string                           `json:"EngineVersion,omitempty" name:"EngineVersion"`
	Description          *string                           `json:"Description,omitempty" name:"Description"`
	Parameters           *CreateDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
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

type ModifyDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId   *string                           `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	DBParameterGroupName *string                           `json:"DBParameterGroupName,omitempty" name:"DBParameterGroupName"`
	Description          *string                           `json:"Description,omitempty" name:"Description"`
	Parameters           *ModifyDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
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
				DelayKeyWrite             *string `json:"DelayKeyWrite" name:"DelayKeyWrite"`
				ConnectTimeout            *int    `json:"ConnectTimeout" name:"ConnectTimeout"`
				InnodbPurgeBatchSize      *int    `json:"InnodbPurgeBatchSize" name:"InnodbPurgeBatchSize"`
				MyisamSortBufferSize      *int    `json:"MyisamSortBufferSize" name:"MyisamSortBufferSize"`
				BulkInsertBufferSize      *int    `json:"BulkInsertBufferSize" name:"BulkInsertBufferSize"`
				InnodbConcurrencyTickets  *int    `json:"InnodbConcurrencyTickets" name:"InnodbConcurrencyTickets"`
				DivPrecisionIncrement     *int    `json:"DivPrecisionIncrement" name:"DivPrecisionIncrement"`
				MaxPreparedStmtCount      *int    `json:"MaxPreparedStmtCount" name:"MaxPreparedStmtCount"`
				TxIsolation               *string `json:"TxIsolation" name:"TxIsolation"`
				WaitTimeout               *int    `json:"WaitTimeout" name:"WaitTimeout"`
				TableDefinitionCache      *int    `json:"TableDefinitionCache" name:"TableDefinitionCache"`
				AutoIncrementIncrement    *int    `json:"AutoIncrementIncrement" name:"AutoIncrementIncrement"`
				FtQueryExpansionLimit     *int    `json:"FtQueryExpansionLimit" name:"FtQueryExpansionLimit"`
				InnodbStatsSamplePages    *int    `json:"InnodbStatsSamplePages" name:"InnodbStatsSamplePages"`
				InnodbOldBlocksTime       *int    `json:"InnodbOldBlocksTime" name:"InnodbOldBlocksTime"`
				InnodbPrintAllDeadlocks   *string `json:"InnodbPrintAllDeadlocks" name:"InnodbPrintAllDeadlocks"`
				SyncBinlog                *int    `json:"SyncBinlog" name:"SyncBinlog"`
				InnodbStatsMethod         *string `json:"InnodbStatsMethod" name:"InnodbStatsMethod"`
				LockWaitTimeout           *int    `json:"LockWaitTimeout" name:"LockWaitTimeout"`
				NetReadTimeout            *int    `json:"NetReadTimeout" name:"NetReadTimeout"`
				QueryPreallocSize         *int    `json:"QueryPreallocSize" name:"QueryPreallocSize"`
				MaxErrorCount             *int    `json:"MaxErrorCount" name:"MaxErrorCount"`
				BackLog                   *int    `json:"BackLog" name:"BackLog"`
				KeyCacheAgeThreshold      *int    `json:"KeyCacheAgeThreshold" name:"KeyCacheAgeThreshold"`
				InnodbLogFileSize         *int    `json:"InnodbLogFileSize" name:"InnodbLogFileSize"`
				InnodbThreadConcurrency   *int    `json:"InnodbThreadConcurrency" name:"InnodbThreadConcurrency"`
				InnodbFlushLogAtTrxCommit *int    `json:"InnodbFlushLogAtTrxCommit" name:"InnodbFlushLogAtTrxCommit"`
				InnodbStrictMode          *string `json:"InnodbStrictMode" name:"InnodbStrictMode"`
				DefaultTimeZone           *string `json:"DefaultTimeZone" name:"DefaultTimeZone"`
				PerformanceSchema         *string `json:"PerformanceSchema" name:"PerformanceSchema"`
				InnodbWriteIoThreads      *int    `json:"InnodbWriteIoThreads" name:"InnodbWriteIoThreads"`
				RplSemiSyncMasterTimeout  *int    `json:"RplSemiSyncMasterTimeout" name:"RplSemiSyncMasterTimeout"`
				MaxConnectErrors          *int    `json:"MaxConnectErrors" name:"MaxConnectErrors"`
				JoinBufferSize            *int    `json:"JoinBufferSize" name:"JoinBufferSize"`
				RplSemiSyncSlaveEnabled   *string `json:"RplSemiSyncSlaveEnabled" name:"RplSemiSyncSlaveEnabled"`
				InnodbRollbackOnTimeout   *string `json:"InnodbRollbackOnTimeout" name:"InnodbRollbackOnTimeout"`
				OldAlterTable             *string `json:"OldAlterTable" name:"OldAlterTable"`
				BinlogRowImage            *string `json:"BinlogRowImage" name:"BinlogRowImage"`
				KeyCacheBlockSize         *int    `json:"KeyCacheBlockSize" name:"KeyCacheBlockSize"`
				QueryCacheType            *string `json:"QueryCacheType" name:"QueryCacheType"`
				LocalInfile               *string `json:"LocalInfile" name:"LocalInfile"`
				InitConnect               *string `json:"InitConnect" name:"InitConnect"`
				BinlogFormat              *string `json:"BinlogFormat" name:"BinlogFormat"`
				LogSlaveUpdates           *string `json:"LogSlaveUpdates" name:"LogSlaveUpdates"`
				SlowLaunchTime            *int    `json:"SlowLaunchTime" name:"SlowLaunchTime"`
				NetWriteTimeout           *int    `json:"NetWriteTimeout" name:"NetWriteTimeout"`
				InnodbTableLocks          *string `json:"InnodbTableLocks" name:"InnodbTableLocks"`
				QueryAllocBlockSize       *int    `json:"QueryAllocBlockSize" name:"QueryAllocBlockSize"`
				TmpTableSize              *int    `json:"TmpTableSize" name:"TmpTableSize"`
				LowerCaseTableNames       *int    `json:"LowerCaseTableNames" name:"LowerCaseTableNames"`
				DefaultWeekFormat         *int    `json:"DefaultWeekFormat" name:"DefaultWeekFormat"`
				KeyCacheDivisionLimit     *int    `json:"KeyCacheDivisionLimit" name:"KeyCacheDivisionLimit"`
				InnodbLockWaitTimeout     *int    `json:"InnodbLockWaitTimeout" name:"InnodbLockWaitTimeout"`
				DelayedInsertTimeout      *int    `json:"DelayedInsertTimeout" name:"DelayedInsertTimeout"`
				NetRetryCount             *int    `json:"NetRetryCount" name:"NetRetryCount"`
				InnodbPurgeThreads        *int    `json:"InnodbPurgeThreads" name:"InnodbPurgeThreads"`
				BinlogCacheSize           *int    `json:"BinlogCacheSize" name:"BinlogCacheSize"`
				LowPriorityUpdates        *string `json:"LowPriorityUpdates" name:"LowPriorityUpdates"`
				AutoIncrementOffset       *int    `json:"AutoIncrementOffset" name:"AutoIncrementOffset"`
				QueryCacheLimit           *int    `json:"QueryCacheLimit" name:"QueryCacheLimit"`
				InnodbReadAheadThreshold  *int    `json:"InnodbReadAheadThreshold" name:"InnodbReadAheadThreshold"`
				InnodbMaxDirtyPagesPct    *int    `json:"InnodbMaxDirtyPagesPct" name:"InnodbMaxDirtyPagesPct"`
				FtMinWordLen              *int    `json:"FtMinWordLen" name:"FtMinWordLen"`
				ConcurrentInsert          *string `json:"ConcurrentInsert" name:"ConcurrentInsert"`
				IntQueryTime              *int    `json:"IntQueryTime" name:"IntQueryTime"`
				SlowQueryLog              *string `json:"SlowQueryLog" name:"SlowQueryLog"`
				SortBufferSize            *int    `json:"SortBufferSize" name:"SortBufferSize"`
				InteractiveTimeout        *int    `json:"InteractiveTimeout" name:"InteractiveTimeout"`
				QueryCacheSize            *int    `json:"QueryCacheSize" name:"QueryCacheSize"`
				InnodbReadIoThreads       *int    `json:"InnodbReadIoThreads" name:"InnodbReadIoThreads"`
				RplSemiSyncMasterEnabled  *string `json:"RplSemiSyncMasterEnabled" name:"RplSemiSyncMasterEnabled"`
				MaxAllowedPacket          *int    `json:"MaxAllowedPacket" name:"MaxAllowedPacket"`
				DelayedInsertLimit        *int    `json:"DelayedInsertLimit" name:"DelayedInsertLimit"`
				InnodbOpenFiles           *int    `json:"InnodbOpenFiles" name:"InnodbOpenFiles"`
				CharacterSetServer        *string `json:"CharacterSetServer" name:"CharacterSetServer"`
				DelayedQueueSize          *int    `json:"DelayedQueueSize" name:"DelayedQueueSize"`
				MaxUserConnections        *int    `json:"MaxUserConnections" name:"MaxUserConnections"`
				InnodbOldBlocksPct        *int    `json:"InnodbOldBlocksPct" name:"InnodbOldBlocksPct"`
				TableOpenCache            *int    `json:"TableOpenCache" name:"TableOpenCache"`
				LogSlowAdminStatements    *string `json:"LogSlowAdminStatements" name:"LogSlowAdminStatements"`
				LogBinTrustFunctionCreators *string `json:"LogBinTrustFunctionCreators" name:"LogBinTrustFunctionCreators"`
				LogQueriesNotUsingIndexes *string `json:"LogQueriesNotUsingIndexes" name:"LogQueriesNotUsingIndexes"`
				InnodbStatsOnMetadata     *string `json:"InnodbStatsOnMetadata" name:"InnodbStatsOnMetadata"`
				TableOpenCacheInstances   *int    `json:"TableOpenCacheInstances" name:"TableOpenCacheInstances"`
				GroupConcatMaxLen         *int    `json:"GroupConcatMaxLen" name:"GroupConcatMaxLen"`
				InnodbSortBufferSize      *int    `json:"InnodbSortBufferSize" name:"InnodbSortBufferSize"`
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

