package v20160701

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type ModifyDBParameterGroupParameters struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RestoreToCurInstanceSrcDatabases struct {
}

type RestoreToCurInstanceDstDatabases struct {
}

type RestoreToSgInstanceSrcDatabases struct {
}

type RestoreToSgInstanceDstDatabases struct {
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

type RebootDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *RebootDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RebootDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebootDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerEip                         *string `json:"InnerEip"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *ModifyDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Description          *string `json:"Description"`
			Parameters           struct {
				Delay_key_write                 *string `json:"Delay_key_write"`
				Tx_isolation                    *string `json:"Tx_isolation"`
				Innodb_print_all_deadlocks      *string `json:"Innodb_print_all_deadlocks"`
				Innodb_stats_method             *string `json:"Innodb_stats_method"`
				Innodb_strict_mode              *string `json:"Innodb_strict_mode"`
				Default_time_zone               *string `json:"Default_time_zone"`
				Performance_schema              *string `json:"Performance_schema"`
				Rpl_semi_sync_slave_enabled     *string `json:"Rpl_semi_sync_slave_enabled"`
				Innodb_rollback_on_timeout      *string `json:"Innodb_rollback_on_timeout"`
				Old_alter_table                 *string `json:"Old_alter_table"`
				Binlog_row_image                *string `json:"Binlog_row_image"`
				Query_cache_type                *string `json:"Query_cache_type"`
				Local_infile                    *string `json:"Local_infile"`
				Init_connect                    *string `json:"Init_connect"`
				Binlog_format                   *string `json:"Binlog_format"`
				Log_slave_updates               *string `json:"Log_slave_updates"`
				Innodb_table_locks              *string `json:"Innodb_table_locks"`
				Low_priority_updates            *string `json:"Low_priority_updates"`
				Concurrent_insert               *string `json:"Concurrent_insert"`
				Slow_query_log                  *string `json:"Slow_query_log"`
				Rpl_semi_sync_master_enabled    *string `json:"Rpl_semi_sync_master_enabled"`
				Character_set_server            *string `json:"Character_set_server"`
				Log_slow_admin_statements       *string `json:"Log_slow_admin_statements"`
				Log_bin_trust_function_creators *string `json:"Log_bin_trust_function_creators"`
				Log_queries_not_using_indexes   *string `json:"Log_queries_not_using_indexes"`
				Innodb_stats_on_metadata        *string `json:"Innodb_stats_on_metadata"`
			} `json:"Parameters"`
			Engine *string `json:"Engine"`
		} `json:"DBParameterGroup"`
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

func (r *ResetDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Description          *string `json:"Description"`
			Parameters           struct {
				Delay_key_write                 *string `json:"Delay_key_write"`
				Tx_isolation                    *string `json:"Tx_isolation"`
				Innodb_print_all_deadlocks      *string `json:"Innodb_print_all_deadlocks"`
				Innodb_stats_method             *string `json:"Innodb_stats_method"`
				Innodb_strict_mode              *string `json:"Innodb_strict_mode"`
				Default_time_zone               *string `json:"Default_time_zone"`
				Performance_schema              *string `json:"Performance_schema"`
				Rpl_semi_sync_slave_enabled     *string `json:"Rpl_semi_sync_slave_enabled"`
				Innodb_rollback_on_timeout      *string `json:"Innodb_rollback_on_timeout"`
				Old_alter_table                 *string `json:"Old_alter_table"`
				Binlog_row_image                *string `json:"Binlog_row_image"`
				Query_cache_type                *string `json:"Query_cache_type"`
				Local_infile                    *string `json:"Local_infile"`
				Init_connect                    *string `json:"Init_connect"`
				Binlog_format                   *string `json:"Binlog_format"`
				Log_slave_updates               *string `json:"Log_slave_updates"`
				Innodb_table_locks              *string `json:"Innodb_table_locks"`
				Low_priority_updates            *string `json:"Low_priority_updates"`
				Concurrent_insert               *string `json:"Concurrent_insert"`
				Slow_query_log                  *string `json:"Slow_query_log"`
				Rpl_semi_sync_master_enabled    *string `json:"Rpl_semi_sync_master_enabled"`
				Character_set_server            *string `json:"Character_set_server"`
				Log_slow_admin_statements       *string `json:"Log_slow_admin_statements"`
				Log_bin_trust_function_creators *string `json:"Log_bin_trust_function_creators"`
				Log_queries_not_using_indexes   *string `json:"Log_queries_not_using_indexes"`
				Innodb_stats_on_metadata        *string `json:"Innodb_stats_on_metadata"`
			} `json:"Parameters"`
			Engine *string `json:"Engine"`
		} `json:"DBParameterGroup"`
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

func (r *DescribeDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroups []struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Engine               *string `json:"Engine"`
		} `json:"DBParameterGroups"`
		Source     *string `json:"Source"`
		MaxRecords *int    `json:"MaxRecords"`
		TotalCount *int    `json:"TotalCount"`
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

func (r *DescribeEngineDefaultParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEngineDefaultParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineDefaultParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engine        *string `json:"Engine"`
		EngineVersion *string `json:"EngineVersion"`
		Parameters    struct {
			Connect_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Connect_timeout"`
			Delay_key_write struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Delay_key_write"`
			Innodb_purge_batch_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_purge_batch_size"`
			Myisam_sort_buffer_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Myisam_sort_buffer_size"`
			Bulk_insert_buffer_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Bulk_insert_buffer_size"`
			Div_precision_increment struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Div_precision_increment"`
			Innodb_concurrency_tickets struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_concurrency_tickets"`
			Max_prepared_stmt_count struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Max_prepared_stmt_count"`
			Wait_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Wait_timeout"`
			Tx_isolation struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Alias           *string `json:"Alias"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Tx_isolation"`
			Table_definition_cache struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Table_definition_cache"`
			Auto_increment_increment struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Auto_increment_increment"`
			Ft_query_expansion_limit struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Ft_query_expansion_limit"`
			Innodb_old_blocks_time struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_old_blocks_time"`
			Innodb_stats_sample_pages struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_stats_sample_pages"`
			Innodb_print_all_deadlocks struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Innodb_print_all_deadlocks"`
			Sync_binlog struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Sync_binlog"`
			Innodb_stats_method struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Innodb_stats_method"`
			Lock_wait_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Lock_wait_timeout"`
			Net_read_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Net_read_timeout"`
			Query_prealloc_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Query_prealloc_size"`
			Back_log struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Back_log"`
			Max_error_count struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Max_error_count"`
			Key_cache_age_threshold struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Key_cache_age_threshold"`
			Innodb_log_file_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_log_file_size"`
			Innodb_thread_concurrency struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_thread_concurrency"`
			Innodb_strict_mode struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Innodb_strict_mode"`
			Innodb_flush_log_at_trx_commit struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_flush_log_at_trx_commit"`
			Default_time_zone struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Default_time_zone"`
			Performance_schema struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Performance_schema"`
			Innodb_write_io_threads struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_write_io_threads"`
			Rpl_semi_sync_master_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Rpl_semi_sync_master_timeout"`
			Max_connect_errors struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Max_connect_errors"`
			Join_buffer_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Join_buffer_size"`
			Rpl_semi_sync_slave_enabled struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Rpl_semi_sync_slave_enabled"`
			Innodb_rollback_on_timeout struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Innodb_rollback_on_timeout"`
			Old_alter_table struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Old_alter_table"`
			Binlog_row_image struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Binlog_row_image"`
			Key_cache_block_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Key_cache_block_size"`
			Query_cache_type struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Ignore          *bool   `json:"Ignore"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Query_cache_type"`
			Init_connect struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Init_connect"`
			Local_infile struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Local_infile"`
			Binlog_format struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Binlog_format"`
			Log_slave_updates struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Log_slave_updates"`
			Slow_launch_time struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Slow_launch_time"`
			Innodb_table_locks struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Innodb_table_locks"`
			Net_write_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Net_write_timeout"`
			Query_alloc_block_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Query_alloc_block_size"`
			Lower_case_table_names struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Lower_case_table_names"`
			Tmp_table_size struct {
				Default            *string `json:"Default"`
				Max                *string `json:"Max"`
				RestartRequired    *bool   `json:"RestartRequired"`
				Type               *string `json:"Type"`
				DefaultScaleFactor *string `json:"DefaultScaleFactor"`
				MaxScaleFactor     *string `json:"MaxScaleFactor"`
				Variable           *string `json:"Variable"`
			} `json:"Tmp_table_size"`
			Default_week_format struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Default_week_format"`
			Key_cache_division_limit struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Key_cache_division_limit"`
			Innodb_lock_wait_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_lock_wait_timeout"`
			Delayed_insert_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Delayed_insert_timeout"`
			Net_retry_count struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Net_retry_count"`
			Innodb_purge_threads struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_purge_threads"`
			Binlog_cache_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Binlog_cache_size"`
			Low_priority_updates struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Low_priority_updates"`
			Auto_increment_offset struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Auto_increment_offset"`
			Innodb_max_dirty_pages_pct struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_max_dirty_pages_pct"`
			Innodb_read_ahead_threshold struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_read_ahead_threshold"`
			Query_cache_limit struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Query_cache_limit"`
			Ft_min_word_len struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Ft_min_word_len"`
			Concurrent_insert struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Concurrent_insert"`
			Int_query_time struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Int_query_time"`
			Slow_query_log struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Slow_query_log"`
			Sort_buffer_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Sort_buffer_size"`
			Interactive_timeout struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Interactive_timeout"`
			Query_cache_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Query_cache_size"`
			Innodb_read_io_threads struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_read_io_threads"`
			Max_allowed_packet struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Max_allowed_packet"`
			Rpl_semi_sync_master_enabled struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Rpl_semi_sync_master_enabled"`
			Delayed_insert_limit struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Delayed_insert_limit"`
			Innodb_open_files struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_open_files"`
			Character_set_server struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Character_set_server"`
			Delayed_queue_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Delayed_queue_size"`
			Max_user_connections struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Max_user_connections"`
			Innodb_old_blocks_pct struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_old_blocks_pct"`
			Table_open_cache struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Table_open_cache"`
			Log_slow_admin_statements struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Log_slow_admin_statements"`
			Log_bin_trust_function_creators struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Log_bin_trust_function_creators"`
			Log_queries_not_using_indexes struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Log_queries_not_using_indexes"`
			Innodb_stats_on_metadata struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Innodb_stats_on_metadata"`
			Table_open_cache_instances struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Table_open_cache_instances"`
			Group_concat_max_len struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Group_concat_max_len"`
			Innodb_sort_buffer_size struct {
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Innodb_sort_buffer_size"`
		} `json:"Parameters"`
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

func (r *CreateDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Parameters           struct {
			} `json:"Parameters"`
			Engine *string `json:"Engine"`
		} `json:"DBParameterGroup"`
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

func (r *DeleteDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			Vip                   *string `json:"Vip"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime"`
			MasterUserName        *string `json:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZoneList  struct {
			} `json:"AvailabilityZoneList"`
			Audit                            *bool `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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
}

func (r *RestoreDBInstanceFromDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreDBInstanceFromDBBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreDBInstanceFromDBBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDBInstanceFromDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			Vip                   *string `json:"Vip"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZoneList  struct {
			} `json:"AvailabilityZoneList"`
			Audit                            *bool `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
				DBBackupName         *string `json:"DBBackupName"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *DeleteDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CreateDBInstanceReadReplicaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBInstanceReadReplicaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceReadReplicaResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			Vip                   *string `json:"Vip"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			MasterUserName        *string `json:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZoneList  struct {
			} `json:"AvailabilityZoneList"`
			Audit                            *bool `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
			} `json:"DBSource"`
			ProductId      *string `json:"ProductId"`
			Region         *string `json:"Region"`
			SubOrderId     *string `json:"SubOrderId"`
			SecurityGroups struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *RestoreDBInstanceToPointInTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreDBInstanceToPointInTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDBInstanceToPointInTimeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"Instances"`
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

func (r *DescribeDBInstanceRestorableTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceRestorableTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRestorableTimeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RestorableTime struct {
			Begin *string `json:"Begin"`
			End   *string `json:"End"`
		} `json:"RestorableTime"`
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

func (r *ModifyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"Instances"`
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

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DescribeDBLogFiles []struct {
			LogFileName *string `json:"LogFileName"`
			Date        *string `json:"Date"`
			StartTime   *string `json:"StartTime"`
			EndTime     *string `json:"EndTime"`
		} `json:"DescribeDBLogFiles"`
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

func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBBackup []struct {
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
			DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
			Engine               *string `json:"Engine"`
			EngineVersion        *string `json:"EngineVersion"`
			BackupCreateTime     *string `json:"BackupCreateTime"`
			BackupUpdatedTime    *string `json:"BackupUpdatedTime"`
			DBBackupName         *string `json:"DBBackupName"`
			BackupMode           *string `json:"BackupMode"`
			BackupType           *string `json:"BackupType"`
			Status               *string `json:"Status"`
			BackupSize           *int    `json:"BackupSize"`
			BackupLocationRef    *string `json:"BackupLocationRef"`
			RemotePath           *string `json:"RemotePath"`
			MD5                  *string `json:"MD5"`
		} `json:"DBBackup"`
		TotalCount *int `json:"TotalCount"`
		MaxRecords *int `json:"MaxRecords"`
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

func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"Instances"`
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
	EIPIn                  *string   `json:"EIPIn,omitempty" name:"EIPIn"`
	ExpiryDateLessThan     *int      `json:"ExpiryDateLessThan,omitempty" name:"ExpiryDateLessThan"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
			DBSource             struct {
				DBInstanceName *string `json:"DBInstanceName"`
			} `json:"DBSource"`
			DBInstanceName                   *string `json:"DBInstanceName"`
			DBInstanceStatus                 *string `json:"DBInstanceStatus"`
			DBInstanceType                   *string `json:"DBInstanceType"`
			GroupId                          *string `json:"GroupId"`
			Vip                              *string `json:"Vip"`
			Engine                           *string `json:"Engine"`
			EngineVersion                    *string `json:"EngineVersion"`
			InstanceCreateTime               *string `json:"InstanceCreateTime"`
			MasterUserName                   *string `json:"MasterUserName"`
			VpcId                            *string `json:"VpcId"`
			SubnetId                         *string `json:"SubnetId"`
			PubliclyAccessible               *bool   `json:"PubliclyAccessible"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			OrderSource           *string `json:"OrderSource"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZone      *string `json:"AvailabilityZone"`
			ProductId             *string `json:"ProductId"`
			OrderUse              *string `json:"OrderUse"`
			ProjectName           *string `json:"ProjectName"`
			Region                *string `json:"Region"`
			Tags                  struct {
			} `json:"Tags"`
		} `json:"Instances"`
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

func (r *OverrideDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "OverrideDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OverrideDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"Instances"`
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

func (r *DescribeDBEngineVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBEngineVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEngineVersionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engines struct {
			MySQL []struct {
				Engine        *string `json:"Engine"`
				EngineVersion *string `json:"EngineVersion"`
			} `json:"MySQL"`
			Consistent_mysql []struct {
				Engine        *string `json:"Engine"`
				EngineVersion *string `json:"EngineVersion"`
			} `json:"Consistent_mysql"`
			Percona []struct {
				Engine        *string `json:"Engine"`
				EngineVersion *string `json:"EngineVersion"`
			} `json:"Percona"`
			Ebs_mysql []struct {
				Engine        *string `json:"Engine"`
				EngineVersion *string `json:"EngineVersion"`
			} `json:"Ebs_mysql"`
		} `json:"Engines"`
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

func (r *UpgradeDBInstanceEngineVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceEngineVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"Instances"`
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

func (r *ModifyDBInstanceTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceTypeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			SecurityGroupId       *string `json:"SecurityGroupId"`
			Vip                   *string `json:"Vip"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime"`
			MasterUserName        *string `json:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZone      *string `json:"AvailabilityZone"`
			AvailabilityZoneList  []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
				PointInTime          *string `json:"PointInTime"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceEndTime   *string `json:"ServiceEndTime"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *DescribeDBInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engine        *string `json:"Engine"`
		EngineVersion *string `json:"EngineVersion"`
		Parameters    struct {
			delay_key_write                 *string `json:"delay_key_write"`
			tx_isolation                    *string `json:"tx_isolation"`
			innodb_print_all_deadlocks      *string `json:"innodb_print_all_deadlocks"`
			innodb_stats_method             *string `json:"innodb_stats_method"`
			innodb_strict_mode              *string `json:"innodb_strict_mode"`
			default_time_zone               *string `json:"default_time_zone"`
			performance_schema              *string `json:"performance_schema"`
			rpl_semi_sync_slave_enabled     *string `json:"rpl_semi_sync_slave_enabled"`
			innodb_rollback_on_timeout      *string `json:"innodb_rollback_on_timeout"`
			old_alter_table                 *string `json:"old_alter_table"`
			binlog_row_image                *string `json:"binlog_row_image"`
			query_cache_type                *string `json:"query_cache_type"`
			local_infile                    *string `json:"local_infile"`
			init_connect                    *string `json:"init_connect"`
			binlog_format                   *string `json:"binlog_format"`
			log_slave_updates               *string `json:"log_slave_updates"`
			innodb_table_locks              *string `json:"innodb_table_locks"`
			low_priority_updates            *string `json:"low_priority_updates"`
			concurrent_insert               *string `json:"concurrent_insert"`
			slow_query_log                  *string `json:"slow_query_log"`
			rpl_semi_sync_master_enabled    *string `json:"rpl_semi_sync_master_enabled"`
			character_set_server            *string `json:"character_set_server"`
			log_slow_admin_statements       *string `json:"log_slow_admin_statements"`
			log_bin_trust_function_creators *string `json:"log_bin_trust_function_creators"`
			log_queries_not_using_indexes   *string `json:"log_queries_not_using_indexes"`
			innodb_stats_on_metadata        *string `json:"innodb_stats_on_metadata"`
		} `json:"Parameters"`
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

type DeleteDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBBackupIdentifier *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
}

func (r *DeleteDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDBBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CreateDBBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	DBBackup struct {
		DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
		DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
		Engine               *string `json:"Engine"`
		EngineVersion        *string `json:"EngineVersion"`
		BackupCreateTime     *string `json:"BackupCreateTime"`
		BackupUpdatedTime    *string `json:"BackupUpdatedTime"`
		DBBackupName         *string `json:"DBBackupName"`
		BackupType           *string `json:"BackupType"`
		Status               *string `json:"Status"`
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

type RenewDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Duration             *int    `json:"Duration,omitempty" name:"Duration"`
	DurationUnit         *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	BillType             *string `json:"BillType,omitempty" name:"BillType"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *RenewDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenewDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string `json:"DBInstanceIdentifier"`
			PreferredBackupTime              *string `json:"PreferredBackupTime"`
			DBInstanceName                   *string `json:"DBInstanceName"`
			DBInstanceStatus                 *string `json:"DBInstanceStatus"`
			DBInstanceType                   *string `json:"DBInstanceType"`
			DBParameterGroupId               *string `json:"DBParameterGroupId"`
			GroupId                          *string `json:"GroupId"`
			SecurityGroupId                  *string `json:"SecurityGroupId"`
			Vip                              *string `json:"Vip"`
			Engine                           *string `json:"Engine"`
			EngineVersion                    *string `json:"EngineVersion"`
			InstanceCreateTime               *string `json:"InstanceCreateTime"`
			MasterUserName                   *string `json:"MasterUserName"`
			DatastoreVersionId               *string `json:"DatastoreVersionId"`
			Region                           *string `json:"Region"`
			VpcId                            *string `json:"VpcId"`
			PubliclyAccessible               *bool   `json:"PubliclyAccessible"`
			ReadReplicaDBInstanceIdentifiers []struct {
				Id                              *string `json:"Id"`
				Vip                             *string `json:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			OrderSource            *string `json:"OrderSource"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			ProductId              *string `json:"ProductId"`
			ServiceEndTime         *string `json:"ServiceEndTime"`
			OrderUse               *string `json:"OrderUse"`
			ProjectName            *string `json:"ProjectName"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RenewDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewDBInstanceResponse) FromJsonString(s string) error {
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

func (r *SwitchDBInstanceHARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SwitchDBInstanceHARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceHAResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *GenerateDBAdminURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GenerateDBAdminURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type StatisticDBInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ExpiryDateLessThan *int `json:"ExpiryDateLessThan,omitempty" name:"ExpiryDateLessThan"`
}

func (r *StatisticDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StatisticDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StatisticDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Statistic struct {
		} `json:"Statistic"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StatisticDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticDBInstancesResponse) FromJsonString(s string) error {
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

func (r *AllocateDBInstanceEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AllocateDBInstanceEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AllocateDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
				Vip                             *string `json:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier"`
				Id                              *string `json:"Id"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *ReleaseDBInstanceEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReleaseDBInstanceEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
				Vip                             *string `json:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier"`
				Id                              *string `json:"Id"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupName *string `json:"SecurityGroupName"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

func (r *ModifyDBInstanceAvailabilityZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceAvailabilityZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceAvailabilityZoneResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id *string `json:"Id"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
		} `json:"DBInstance"`
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

type DescribeDBInstanceRegionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDBInstanceRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRegionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Regions []struct {
			Code              *string `json:"Code"`
			Name              *string `json:"Name"`
			RegionEnName      *string `json:"RegionEnName"`
			AreaCode          *string `json:"AreaCode"`
			AreaName          *string `json:"AreaName"`
			AreaEnName        *string `json:"AreaEnName"`
			AvailabilityZones []struct {
				Code *string `json:"Code"`
				Name *string `json:"Name"`
			} `json:"AvailabilityZones"`
		} `json:"Regions"`
	} `json:"Data"`
}

func (r *DescribeDBInstanceRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancePackagesRequest struct {
	*ksyunhttp.BaseRequest
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
}

func (r *DescribeDBInstancePackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstancePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancePackagesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Packages []struct {
			Region struct {
				Code *string `json:"Code"`
				Name *string `json:"Name"`
			} `json:"Region"`
			PackageSet []struct {
				Memory            *int    `json:"Memory"`
				DataDiskMax       *int    `json:"DataDiskMax"`
				DataDiskMin       *int    `json:"DataDiskMin"`
				DateDiskStep      *int    `json:"DateDiskStep"`
				PackageCode       *string `json:"PackageCode"`
				DataDiskRecommend *string `json:"DataDiskRecommend"`
			} `json:"PackageSet"`
		} `json:"Packages"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstancePackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancePackagesResponse) FromJsonString(s string) error {
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

func (r *DescribeLastLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLastLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLastLogResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		LogFiles []struct {
			LogFileName *string `json:"LogFileName"`
			Size        *string `json:"Size"`
			DBLogType   *string `json:"DBLogType"`
			Status      *string `json:"Status"`
		} `json:"LogFiles"`
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

func (r *StartAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *StopAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StopAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAuditResponse struct {
	*ksyunhttp.BaseResponse
	AuditRows []struct {
		Id                 *string `json:"Id"`
		InstanceId         *string `json:"InstanceId"`
		AccessUsername     *string `json:"AccessUsername"`
		SourceIp           *string `json:"SourceIp"`
		AccessSqlExt       *string `json:"AccessSqlExt"`
		AccessSqlStatement *string `json:"AccessSqlStatement"`
		AccessSqlLanguage  *string `json:"AccessSqlLanguage"`
		AccessDBName       *string `json:"AccessDBName"`
		RunResult          *bool   `json:"RunResult"`
		Duration           *string `json:"Duration"`
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

func (r *AuditStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AuditStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AuditStatisticResponse struct {
	*ksyunhttp.BaseResponse
	AuditStatistic struct {
		AccessSqlLanguage struct {
		} `json:"AccessSqlLanguage"`
		AccessSqlStatement struct {
		} `json:"AccessSqlStatement"`
	} `json:"AuditStatistic"`
}

func (r *AuditStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditStatisticResponse) FromJsonString(s string) error {
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

func (r *GetTableRestorableTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetTableRestorableTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTableRestorableTimeResponse struct {
	*ksyunhttp.BaseResponse
	RestorableTime struct {
		Begin *string `json:"Begin"`
		End   *string `json:"End"`
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

func (r *GetHistoryDatabaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetHistoryDatabaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetHistoryDatabaseInfoResponse struct {
	*ksyunhttp.BaseResponse
	Databases []struct {
		DatabaseName *string `json:"DatabaseName"`
		TableNames   []struct {
		} `json:"TableNames"`
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

func (r *OverrideDBInstanceByPointInTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "OverrideDBInstanceByPointInTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *RestoreToCurInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreToCurInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *RestoreToSgInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreToSgInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeAuditHotCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAuditHotCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditHotCountResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		DbName    *string `json:"DbName"`
		TableName *string `json:"TableName"`
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

func (r *DescribeAuditHotDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAuditHotDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditHotDurationResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		DbName    *string `json:"DbName"`
		TableName *string `json:"TableName"`
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

func (r *SqlAuditReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SqlAuditReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SqlAuditReportResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		DbName      *string `json:"DbName"`
		SqlTemplate *string `json:"SqlTemplate"`
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

func (r *SqlAuditLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SqlAuditLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SqlAuditLineChartResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		ProductId *string `json:"ProductId"`
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

func (r *SlowLogReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SlowLogReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogReportResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		ProcessingTime  *int    `json:"ProcessingTime"`
		Id              *int    `json:"Id"`
		Checksum        *string `json:"Checksum"`
		Fingerprint     *string `json:"Fingerprint"`
		QueryCount      *int    `json:"QueryCount"`
		QueryTimeAvg    *int    `json:"QueryTimeAvg"`
		QueryTimeMax    *int    `json:"QueryTimeMax"`
		QueryTimeSum    *int    `json:"QueryTimeSum"`
		LockTimeAvg     *int    `json:"LockTimeAvg"`
		LockTimeMax     *int    `json:"LockTimeMax"`
		LockTimeSum     *int    `json:"LockTimeSum"`
		RowsExaminedAvg *int    `json:"RowsExaminedAvg"`
		RowsExaminedMax *int    `json:"RowsExaminedMax"`
		RowsExaminedSum *int    `json:"RowsExaminedSum"`
		RowsSentAvg     *int    `json:"RowsSentAvg"`
		RowsSentMax     *int    `json:"RowsSentMax"`
		RowsSentSum     *int    `json:"RowsSentSum"`
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

func (r *SlowLogLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SlowLogLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogLineChartResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		ReportTime  *int    `json:"ReportTime"`
		ReportCount *int    `json:"ReportCount"`
		Percentage  *string `json:"Percentage"`
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

func (r *SlowLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SlowLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogDetailResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		Id               *string `json:"Id"`
		SqlExecTimeStamp *string `json:"SqlExecTimeStamp"`
		AggTime          *string `json:"AggTime"`
		AuthUser         *string `json:"AuthUser"`
		CurrentUser      *string `json:"CurrentUser"`
		Hostname         *string `json:"Hostname"`
		QueryTime        *string `json:"QueryTime"`
		LockTime         *string `json:"LockTime"`
		RowsSent         *string `json:"RowsSent"`
		RowsExamined     *string `json:"RowsExamined"`
		SqlContent       *string `json:"SqlContent"`
		ThreadId         *string `json:"ThreadId"`
		ProductId        *string `json:"ProductId"`
		TenantId         *string `json:"TenantId"`
		Checksum         *string `json:"Checksum"`
		Fingerprint      *string `json:"Fingerprint"`
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

func (r *StartAuditDetailExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartAuditDetailExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListAuditDetailExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAuditDetailExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAuditDetailExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskResps []struct {
		StartTime   *string `json:"StartTime"`
		EndTime     *string `json:"EndTime"`
		S3FileNames []struct {
		} `json:"S3FileNames"`
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

type DescribeInstanceAccountsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeInstanceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Accounts []struct {
			InstanceAccountName        *string `json:"InstanceAccountName"`
			InstanceAccountDescription *string `json:"InstanceAccountDescription"`
			Created                    *string `json:"Created"`
			InstanceAccountType        *string `json:"InstanceAccountType"`
			InstanceAccountPrivileges  []struct {
				InstanceDatabaseName *string `json:"InstanceDatabaseName"`
				Privilege            *string `json:"Privilege"`
			} `json:"InstanceAccountPrivileges"`
		} `json:"Accounts"`
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

func (r *ModifyInstanceAccountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type DescribeCollationsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeCollationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCollationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCollationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCollationsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Collations struct {
			Utf8 []struct {
			} `json:"Utf8"`
			Gbk []struct {
			} `json:"Gbk"`
			Latin1 []struct {
			} `json:"Latin1"`
			Utf8mb4 []struct {
			} `json:"Utf8mb4"`
			Gb18030 []struct {
			} `json:"Gb18030"`
			Koi8u []struct {
			} `json:"Koi8u"`
			Koi8r []struct {
			} `json:"Koi8r"`
			Cp850 []struct {
			} `json:"Cp850"`
			Macce []struct {
			} `json:"Macce"`
			Ujis []struct {
			} `json:"Ujis"`
			Hebrew []struct {
			} `json:"Hebrew"`
			Cp932 []struct {
			} `json:"Cp932"`
			Ascii []struct {
			} `json:"Ascii"`
			Binary []struct {
			} `json:"Binary"`
			Sjis []struct {
			} `json:"Sjis"`
			Armscii8 []struct {
			} `json:"Armscii8"`
			Cp852 []struct {
			} `json:"Cp852"`
			Keybcs2 []struct {
			} `json:"Keybcs2"`
			Cp866 []struct {
			} `json:"Cp866"`
			Geostd8 []struct {
			} `json:"Geostd8"`
			Cp1257 []struct {
			} `json:"Cp1257"`
			Ucs2 []struct {
			} `json:"Ucs2"`
			Dec8 []struct {
			} `json:"Dec8"`
			Cp1250 []struct {
			} `json:"Cp1250"`
			Tis620 []struct {
			} `json:"Tis620"`
			Utf32 []struct {
			} `json:"Utf32"`
			Latin5 []struct {
			} `json:"Latin5"`
			Hp8 []struct {
			} `json:"Hp8"`
			Utf16le []struct {
			} `json:"Utf16le"`
			Latin2 []struct {
			} `json:"Latin2"`
			Macroman []struct {
			} `json:"Macroman"`
			Eucjpms []struct {
			} `json:"Eucjpms"`
			Gb2312 []struct {
			} `json:"Gb2312"`
			Cp1256 []struct {
			} `json:"Cp1256"`
			Big5 []struct {
			} `json:"Big5"`
			Greek []struct {
			} `json:"Greek"`
			Euckr []struct {
			} `json:"Euckr"`
			Cp1251 []struct {
			} `json:"Cp1251"`
			Utf16 []struct {
			} `json:"Utf16"`
			Swe7 []struct {
			} `json:"Swe7"`
			Latin7 []struct {
			} `json:"Latin7"`
		} `json:"Collations"`
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

func (r *CreateInstanceDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateInstanceDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyInstanceDatabasePrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceDatabasePrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeInstanceDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		InstanceDatabases []struct {
			InstanceDatabaseName         *string `json:"InstanceDatabaseName"`
			InstanceDatabaseCollation    *string `json:"InstanceDatabaseCollation"`
			InstanceDatabaseCollationSet *string `json:"InstanceDatabaseCollationSet"`
			InstanceDatabaseDescription  *string `json:"InstanceDatabaseDescription"`
			InstanceDatabaseStatus       *string `json:"InstanceDatabaseStatus"`
			InstanceDatabasePrivileges   struct {
			} `json:"InstanceDatabasePrivileges"`
		} `json:"InstanceDatabases"`
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

func (r *ModifyInstanceDatabaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceDatabaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *StartSlowLogDetailExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartSlowLogDetailExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListSlowLogDetailExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSlowLogDetailExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSlowLogDetailExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskResps []struct {
		StartTime    *string `json:"StartTime"`
		EndTime      *string `json:"EndTime"`
		Status       *int    `json:"Status"`
		CreateTime   *int    `json:"CreateTime"`
		RecordNumber *int    `json:"RecordNumber"`
		S3FileNames  []struct {
		} `json:"S3FileNames"`
	} `json:"TaskResps"`
	totalCount *int    `json:"totalCount" name:"totalCount"`
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

func (r *CreateInstanceAccountActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateInstanceAccountActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountActionResponse struct {
	*ksyunhttp.BaseResponse
	DBInstanceIdentifier       *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string `json:"InstanceAccountName" name:"InstanceAccountName"`
	InstanceAccountPassword    *string `json:"InstanceAccountPassword" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string `json:"InstanceAccountDescription" name:"InstanceAccountDescription"`
	InstanceAccountPrivileges  []struct {
		InstanceDatabaseName *string `json:"InstanceDatabaseName"`
		Privilege            *string `json:"Privilege"`
	} `json:"InstanceAccountPrivileges"`
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

func (r *ModifyInstanceAccountPrivilegesActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountPrivilegesActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteInstanceAccountActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteInstanceAccountActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteInstanceDatabaseActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteInstanceDatabaseActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyDBNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeEngineParametersModifyHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEngineParametersModifyHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineParametersModifyHistoryResponse struct {
	*ksyunhttp.BaseResponse
	request_id *string `json:"request_id" name:"request_id"`
	History    []struct {
		Id                *string `json:"Id"`
		Configuration_key *string `json:"Configuration_key"`
		Old_value         *string `json:"Old_value"`
		New_value         *string `json:"New_value"`
		Created           *string `json:"Created"`
	} `json:"History"`
}

func (r *DescribeEngineParametersModifyHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineParametersModifyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
