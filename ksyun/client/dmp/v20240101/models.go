package v20240101

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeDefaultMonitorItemsRequest struct {
	*ksyunhttp.BaseRequest
	PanelType *string `json:"PanelType,omitempty" name:"PanelType"`
}

func (r *DescribeDefaultMonitorItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDefaultMonitorItemsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MonitorItems struct {
			RdsInnodbRowsInserted              *string `json:"rds.innodb_rows_inserted" name:"rds.innodb_rows_inserted"`
			RdsInnodbBufferPoolPagesTotal      *string `json:"rds.innodb_buffer_pool_pages_total" name:"rds.innodb_buffer_pool_pages_total"`
			RdsSelectScan                      *string `json:"rds.select_scan" name:"rds.select_scan"`
			RdsInnodbOsLogFsyncs               *string `json:"rds.innodb_os_log_fsyncs" name:"rds.innodb_os_log_fsyncs"`
			RdsConnectionUsedPercent           *string `json:"rds.connection_used_percent" name:"rds.connection_used_percent"`
			RdsCpuUsedPercent                  *string `json:"rds.cpu_used_percent" name:"rds.cpu_used_percent"`
			RdsResidentMemorySize              *string `json:"rds.resident_memory_size" name:"rds.resident_memory_size"`
			RdsSpaceUsed                       *string `json:"rds.space_used" name:"rds.space_used"`
			RdsHandlerReadRndNext              *string `json:"rds.handler_read_rnd_next" name:"rds.handler_read_rnd_next"`
			RdsInnodbBufferPoolUseRatio        *string `json:"rds.innodb_buffer_pool_use_ratio" name:"rds.innodb_buffer_pool_use_ratio"`
			RdsComDelete                       *string `json:"rds.com_delete" name:"rds.com_delete"`
			RdsSlowQueries                     *string `json:"rds.slow_queries" name:"rds.slow_queries"`
			RdsInnodbNumOpenFiles              *string `json:"rds.innodb_num_open_files" name:"rds.innodb_num_open_files"`
			RdsSpaceUsedPercent                *string `json:"rds.space_used_percent" name:"rds.space_used_percent"`
			RdsInnodbBufferPoolPagesFree       *string `json:"rds.innodb_buffer_pool_pages_free" name:"rds.innodb_buffer_pool_pages_free"`
			RdsTps                             *string `json:"rds.tps" name:"rds.tps"`
			RdsInnodbBufferPoolPagesFlushedPs  *string `json:"rds.innodb_buffer_pool_pages_flushed_ps" name:"rds.innodb_buffer_pool_pages_flushed_ps"`
			RdsBytesSent                       *string `json:"rds.bytes_sent" name:"rds.bytes_sent"`
			RdsInnodbBufferPoolReads           *string `json:"rds.innodb_buffer_pool_reads" name:"rds.innodb_buffer_pool_reads"`
			RdsQcacheUsedPercent               *string `json:"rds.qcache_used_percent" name:"rds.qcache_used_percent"`
			RdsThreadsCreated                  *string `json:"rds.threads_created" name:"rds.threads_created"`
			RdsRbps                            *string `json:"rds.rbps" name:"rds.rbps"`
			RdsInnodbBufferPoolReadRequests    *string `json:"rds.innodb_buffer_pool_read_requests" name:"rds.innodb_buffer_pool_read_requests"`
			RdsInnodbDataFsyncs                *string `json:"rds.innodb_data_fsyncs" name:"rds.innodb_data_fsyncs"`
			RdsComCommit                       *string `json:"rds.com_commit" name:"rds.com_commit"`
			RdsRiops                           *string `json:"rds.riops" name:"rds.riops"`
			RdsComSelect                       *string `json:"rds.com_select" name:"rds.com_select"`
			RdsQueries                         *string `json:"rds.queries" name:"rds.queries"`
			RdsBytesReceived                   *string `json:"rds.bytes_received" name:"rds.bytes_received"`
			RdsTableLocksWaited                *string `json:"rds.table_locks_waited" name:"rds.table_locks_waited"`
			RdsInnodbRowLockWaits              *string `json:"rds.innodb_row_lock_waits" name:"rds.innodb_row_lock_waits"`
			RdsThreadsConnected                *string `json:"rds.threads_connected" name:"rds.threads_connected"`
			RdsInnodbBufferPoolHitRatio        *string `json:"rds.innodb_buffer_pool_hit_ratio" name:"rds.innodb_buffer_pool_hit_ratio"`
			RdsInnodbRowsUpdated               *string `json:"rds.innodb_rows_updated" name:"rds.innodb_rows_updated"`
			RdsInnodbLogWritesPs               *string `json:"rds.innodb_log_writes_ps" name:"rds.innodb_log_writes_ps"`
			RdsInnodbRowsDeleted               *string `json:"rds.innodb_rows_deleted" name:"rds.innodb_rows_deleted"`
			RdsInnodbRowsRead                  *string `json:"rds.innodb_rows_read" name:"rds.innodb_rows_read"`
			RdsComRollback                     *string `json:"rds.com_rollback" name:"rds.com_rollback"`
			RdsQps                             *string `json:"rds.qps" name:"rds.qps"`
			RdsComReplace                      *string `json:"rds.com_replace" name:"rds.com_replace"`
			RdsThreadsRunning                  *string `json:"rds.threads_running" name:"rds.threads_running"`
			RdsInnodbRowLockTimeAvg            *string `json:"rds.innodb_row_lock_time_avg" name:"rds.innodb_row_lock_time_avg"`
			RdsInnodbRowLockWaitsTotal         *string `json:"rds.innodb_row_lock_waits_total" name:"rds.innodb_row_lock_waits_total"`
			RdsWbps                            *string `json:"rds.wbps" name:"rds.wbps"`
			RdsWiops                           *string `json:"rds.wiops" name:"rds.wiops"`
			RdsComUpdate                       *string `json:"rds.com_update" name:"rds.com_update"`
			RdsComInsert                       *string `json:"rds.com_insert" name:"rds.com_insert"`
			RdsInnodbDataReads                 *string `json:"rds.innodb_data_reads" name:"rds.innodb_data_reads"`
			RdsInnodbDataWrites                *string `json:"rds.innodb_data_writes" name:"rds.innodb_data_writes"`
			RdsMaxConnections                  *string `json:"rds.max_connections" name:"rds.max_connections"`
			RdsTempSpaceUsed                   *string `json:"rds.temp_space_used" name:"rds.temp_space_used"`
			RdsOpenedTables                    *string `json:"rds.opened_tables" name:"rds.opened_tables"`
			RdsInnodbBufferPoolWriteRequestsPs *string `json:"rds.innodb_buffer_pool_write_requests_ps" name:"rds.innodb_buffer_pool_write_requests_ps"`
			RdsMemoryUsedPercent               *string `json:"rds.memory_used_percent" name:"rds.memory_used_percent"`
		} `json:"MonitorItems" name:"MonitorItems"`
		Total *int `json:"Total" name:"Total"`
	} `json:"Data"`
}

func (r *DescribeDefaultMonitorItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultMonitorItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMonitorPanelRequest struct {
	*ksyunhttp.BaseRequest
	PanelId *string `json:"PanelId,omitempty" name:"PanelId"`
}

func (r *DeleteMonitorPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteMonitorPanelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeleteMonitorPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMonitorPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateMonitorPanelRequest struct {
	*ksyunhttp.BaseRequest
	AddInstanceIds     []*string `json:"AddInstanceIds,omitempty" name:"AddInstanceIds"`
	RemoveInstanceIds  []*string `json:"RemoveInstanceIds,omitempty" name:"RemoveInstanceIds"`
	AddMonitorItems    []*string `json:"AddMonitorItems,omitempty" name:"AddMonitorItems"`
	RemoveMonitorItems []*string `json:"RemoveMonitorItems,omitempty" name:"RemoveMonitorItems"`
	PanelId            *string   `json:"PanelId,omitempty" name:"PanelId"`
}

func (r *OperateMonitorPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OperateMonitorPanelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *OperateMonitorPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateMonitorPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorPanelRequest struct {
	*ksyunhttp.BaseRequest
	PanelId *string `json:"PanelId,omitempty" name:"PanelId"`
}

func (r *DescribeMonitorPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMonitorPanelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceList []struct {
			InstanceId     *string `json:"InstanceId" name:"InstanceId"`
			InstanceName   *string `json:"InstanceName" name:"InstanceName"`
			InstanceRegion *string `json:"InstanceRegion" name:"InstanceRegion"`
			DatabaseType   *string `json:"DatabaseType" name:"DatabaseType"`
			Mode           *string `json:"Mode" name:"Mode"`
		} `json:"InstanceList" name:"InstanceList"`
		MonitorItemList []*string `json:"MonitorItemList" name:"MonitorItemList"`
		PanelType       *string   `json:"PanelType" name:"PanelType"`
	} `json:"Data"`
}

func (r *DescribeMonitorPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMonitorPanelInfoRequest struct {
	*ksyunhttp.BaseRequest
	PanelId   *string `json:"PanelId,omitempty" name:"PanelId"`
	PanelName *string `json:"PanelName,omitempty" name:"PanelName"`
}

func (r *ModifyMonitorPanelInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyMonitorPanelInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		PanelId     *string `json:"PanelId" name:"PanelId"`
		PanelName   *string `json:"PanelName" name:"PanelName"`
		PanelType   *string `json:"PanelType" name:"PanelType"`
		Description *string `json:"Description" name:"Description"`
		UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
}

func (r *ModifyMonitorPanelInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMonitorPanelInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMonitorPanelRequest struct {
	*ksyunhttp.BaseRequest
	PanelName *string `json:"PanelName,omitempty" name:"PanelName"`
	PanelType *string `json:"PanelType,omitempty" name:"PanelType"`
}

func (r *CreateMonitorPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMonitorPanelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		PanelId      *string `json:"PanelId" name:"PanelId"`
		PanelName    *string `json:"PanelName" name:"PanelName"`
		DatabaseType *string `json:"DatabaseType" name:"DatabaseType"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
	} `json:"Data"`
}

func (r *CreateMonitorPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMonitorPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DeleteBatchInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBatchInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeleteBatchInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceStatisticsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *InstanceStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type InstanceStatisticsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TypeStatistic []struct {
			Count        *int    `json:"Count" name:"Count"`
			Percent      *string `json:"Percent" name:"Percent"`
			DatabaseType *string `json:"DatabaseType" name:"DatabaseType"`
		} `json:"TypeStatistic" name:"TypeStatistic"`
		RegionStatistic []struct {
			Count          *int    `json:"Count" name:"Count"`
			Percent        *string `json:"Percent" name:"Percent"`
			InstanceRegion *string `json:"InstanceRegion" name:"InstanceRegion"`
		} `json:"RegionStatistic" name:"RegionStatistic"`
		SourceStatistic []struct {
			Count          *int    `json:"Count" name:"Count"`
			Percent        *string `json:"Percent" name:"Percent"`
			InstanceSource *string `json:"InstanceSource" name:"InstanceSource"`
		} `json:"SourceStatistic" name:"SourceStatistic"`
		StatusStatistic struct {
			MySQL struct {
				Normal *int `json:"Normal" name:"Normal"`
				Error  *int `json:"Error" name:"Error"`
			} `json:"MySQL"`
			Redis struct {
				Normal *int `json:"Normal" name:"Normal"`
				Error  *int `json:"Error" name:"Error"`
			} `json:"Redis"`
		} `json:"StatusStatistic" name:"StatusStatistic"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
}

func (r *InstanceStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstanceStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorPanelListRequest struct {
	*ksyunhttp.BaseRequest
	PanelIds  *string `json:"PanelIds,omitempty" name:"PanelIds"`
	PanelType *string `json:"PanelType,omitempty" name:"PanelType"`
	Page      *int    `json:"Page,omitempty" name:"Page"`
	PageSize  *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeMonitorPanelListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMonitorPanelListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MonitorPanels []struct {
			PanelId     *string `json:"PanelId" name:"PanelId"`
			PanelName   *string `json:"PanelName" name:"PanelName"`
			PanelType   *string `json:"PanelType" name:"PanelType"`
			Description *string `json:"Description" name:"Description"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		} `json:"MonitorPanels" name:"MonitorPanels"`
		Page       *int `json:"Page" name:"Page"`
		PageSize   *int `json:"PageSize" name:"PageSize"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
}

func (r *DescribeMonitorPanelListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorPanelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceListRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId     *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseType   *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	InstanceName   *string `json:"InstanceName,omitempty" name:"InstanceName"`
	Ip             *string `json:"Ip,omitempty" name:"Ip"`
	Search         *string `json:"Search,omitempty" name:"Search"`
	Page           *int    `json:"Page,omitempty" name:"Page"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Page         *int `json:"Page" name:"Page"`
		PageSize     *int `json:"PageSize" name:"PageSize"`
		TotalCount   *int `json:"TotalCount" name:"TotalCount"`
		InstanceList []struct {
			InstanceId      *string `json:"InstanceId" name:"InstanceId"`
			InstanceName    *string `json:"InstanceName" name:"InstanceName"`
			DmpStatus       *string `json:"DmpStatus" name:"DmpStatus"`
			InstanceSource  *string `json:"InstanceSource" name:"InstanceSource"`
			InstanceRegion  *string `json:"InstanceRegion" name:"InstanceRegion"`
			DatabaseType    *string `json:"DatabaseType" name:"DatabaseType"`
			DatabaseVersion *string `json:"DatabaseVersion" name:"DatabaseVersion"`
			Mode            *string `json:"Mode" name:"Mode"`
			Ips             *string `json:"Ips" name:"Ips"`
			Port            *int    `json:"Port" name:"Port"`
			Description     *string `json:"Description" name:"Description"`
			ImportTime      *string `json:"ImportTime" name:"ImportTime"`
		} `json:"InstanceList" name:"InstanceList"`
	} `json:"Data"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportInstanceToDmpRequest struct {
	*ksyunhttp.BaseRequest
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	DatabaseType   *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	InstanceId     *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ImportInstanceToDmpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ImportInstanceToDmpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
	} `json:"Data"`
}

func (r *ImportInstanceToDmpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportInstanceToDmpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClustersRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDedicatedClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDedicatedClustersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TotalCount        *int `json:"TotalCount" name:"TotalCount"`
		DedicatedClusters []struct {
			DedicatedClusterId   *string `json:"DedicatedClusterId" name:"DedicatedClusterId"`
			DedicatedClusterName *string `json:"DedicatedClusterName" name:"DedicatedClusterName"`
			Region               *string `json:"Region" name:"Region"`
			Status               *string `json:"Status" name:"Status"`
			DatabaseType         *string `json:"DatabaseType" name:"DatabaseType"`
			HostNum              *int    `json:"HostNum" name:"HostNum"`
			InstanceNum          *int    `json:"InstanceNum" name:"InstanceNum"`
			KceClusterId         *string `json:"KceClusterId" name:"KceClusterId"`
			Ratio                struct {
				MemRatio  *int `json:"MemRatio" name:"MemRatio"`
				DiskRatio *int `json:"DiskRatio" name:"DiskRatio"`
				CpuRatio  *int `json:"CpuRatio" name:"CpuRatio"`
			} `json:"Ratio"`
			CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
		} `json:"DedicatedClusters" name:"DedicatedClusters"`
	} `json:"Data"`
}

func (r *DescribeDedicatedClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedHostsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDedicatedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDedicatedHostsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeDedicatedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseSchemaRequest struct {
	*ksyunhttp.BaseRequest
	DatabaseId *int `json:"DatabaseId,omitempty" name:"DatabaseId"`
}

func (r *DescribeDatabaseSchemaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDatabaseSchemaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DescribeDatabaseSchemaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDatabaseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDatabaseListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeDatabaseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistorySQLRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeHistorySQLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeHistorySQLResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeHistorySQLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistorySQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartExecuteSQLRequest struct {
	*ksyunhttp.BaseRequest
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	Statement    *string `json:"Statement,omitempty" name:"Statement"`
}

func (r *StartExecuteSQLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartExecuteSQLResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *StartExecuteSQLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartExecuteSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UpdateInstanceDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateInstanceDatabaseResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *UpdateInstanceDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDatabaseTableRequest struct {
	*ksyunhttp.BaseRequest
	DatabaseId *string `json:"DatabaseId,omitempty" name:"DatabaseId"`
}

func (r *UpdateDatabaseTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDatabaseTableResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *UpdateDatabaseTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestInstanceConnectionRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseType    *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	DatabaseVersion *string `json:"DatabaseVersion,omitempty" name:"DatabaseVersion"`
	Username        *string `json:"Username,omitempty" name:"Username"`
	Password        *string `json:"Password,omitempty" name:"Password"`
	UseSourceUser   *bool   `json:"UseSourceUser,omitempty" name:"UseSourceUser"`
}

func (r *TestInstanceConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type TestInstanceConnectionResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *TestInstanceConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestInstanceConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
