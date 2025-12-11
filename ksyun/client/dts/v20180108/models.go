package v20180108

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreatePrecheckDTSParameter struct {
	DBParameter            *string `json:"DBParameter,omitempty" name:"DBParameter"`
	TargetDBParameterValue *string `json:"TargetDBParameterValue,omitempty" name:"TargetDBParameterValue"`
}
type CreatePrecheckSourceUser struct {
	Username   *string `json:"Username,omitempty" name:"Username"`
	SourceHost *string `json:"SourceHost,omitempty" name:"SourceHost"`
}

type SchemaStructRequest struct {
	*ksyunhttp.BaseRequest
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	SourceType       *string `json:"SourceType,omitempty" name:"SourceType"`
	SourceUsername   *string `json:"SourceUsername,omitempty" name:"SourceUsername"`
	SourcePassword   *string `json:"SourcePassword,omitempty" name:"SourcePassword"`
	SourceRegion     *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

func (r *SchemaStructRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SchemaStructResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Schemas struct {
			Schema []struct {
				Source    *string `json:"Source" name:"Source"`
				Target    *string `json:"Target" name:"Target"`
				Renamable *bool   `json:"Renamable" name:"Renamable"`
				IsFull    *bool   `json:"IsFull" name:"IsFull"`
				Children  []struct {
					Source    *string `json:"source" name:"source"`
					Target    *string `json:"target" name:"target"`
					Renamable *bool   `json:"renamable" name:"renamable"`
					IsFull    *bool   `json:"is_full" name:"is_full"`
					Children  []struct {
						Source    *string   `json:"Source" name:"Source"`
						Target    *string   `json:"Target" name:"Target"`
						Renamable *bool     `json:"Renamable" name:"Renamable"`
						IsFull    *bool     `json:"IsFull" name:"IsFull"`
						Children  []*string `json:"Children" name:"Children"`
					} `json:"Children"`
				} `json:"Children" name:"Children"`
			} `json:"Schema"`
		} `json:"Schemas" name:"Schemas"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SchemaStructResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SchemaStructResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectivityCheckRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ConnectivityCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ConnectivityCheckResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Status  *string `json:"Status" name:"Status"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ConnectivityCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConnectivityCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrecheckRequest struct {
	*ksyunhttp.BaseRequest
	SourceType       *string                       `json:"SourceType,omitempty" name:"SourceType"`
	TargetType       *string                       `json:"TargetType,omitempty" name:"TargetType"`
	TargetRegion     *string                       `json:"TargetRegion,omitempty" name:"TargetRegion"`
	SourceRegion     *string                       `json:"SourceRegion,omitempty" name:"SourceRegion"`
	DbSchema         *string                       `json:"DbSchema,omitempty" name:"DbSchema"`
	SubTasks         *string                       `json:"SubTasks,omitempty" name:"SubTasks"`
	SourceInstanceId *string                       `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	TargetInstanceId *string                       `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`
	SourceUsername   *string                       `json:"SourceUsername,omitempty" name:"SourceUsername"`
	SourcePassword   *string                       `json:"SourcePassword,omitempty" name:"SourcePassword"`
	Type             *string                       `json:"Type,omitempty" name:"Type"`
	DTSParameter     []*CreatePrecheckDTSParameter `json:"DTSParameter,omitempty" name:"DTSParameter"`
	SourceUser       []*CreatePrecheckSourceUser   `json:"SourceUser,omitempty" name:"SourceUser"`
}

func (r *CreatePrecheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreatePrecheckResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		PrecheckId *string `json:"PrecheckId" name:"PrecheckId"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreatePrecheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrecheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskName   *string `json:"TaskName,omitempty" name:"TaskName"`
	SubTask    *string `json:"SubTask,omitempty" name:"SubTask"`
	TaskType   *string `json:"TaskType,omitempty" name:"TaskType"`
	PrecheckId *string `json:"PrecheckId,omitempty" name:"PrecheckId"`
	BillType   *int    `json:"BillType,omitempty" name:"BillType"`
	Duration   *int    `json:"Duration,omitempty" name:"Duration"`
}

func (r *CreateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *string `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      struct {
		TaskId   *string `json:"TaskId" name:"TaskId"`
		TaskName *string `json:"TaskName" name:"TaskName"`
		TaskType *string `json:"TaskType" name:"TaskType"`
		OrderId  *string `json:"OrderId" name:"OrderId"`
	} `json:"Data"`
}

func (r *CreateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId           *string   `json:"TaskId,omitempty" name:"TaskId"`
	TargetType       *string   `json:"TargetType,omitempty" name:"TargetType"`
	TaskType         *string   `json:"TaskType,omitempty" name:"TaskType"`
	TaskStatus       []*string `json:"TaskStatus,omitempty" name:"TaskStatus"`
	TaskName         *string   `json:"TaskName,omitempty" name:"TaskName"`
	TargetInstanceId *string   `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`
	SourceInstanceId *string   `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	OrderByType      *string   `json:"OrderByType,omitempty" name:"OrderByType"`
	Marker           *int      `json:"Marker,omitempty" name:"Marker"`
	MaxRecords       *int      `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTaskResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Tasks []struct {
			TaskName                     *string `json:"TaskName" name:"TaskName"`
			TaskId                       *string `json:"TaskId" name:"TaskId"`
			TaskStatus                   *string `json:"TaskStatus" name:"TaskStatus"`
			LatestConsistencyCheckStatus *string `json:"LatestConsistencyCheckStatus" name:"LatestConsistencyCheckStatus"`
			Created                      *string `json:"Created" name:"Created"`
			ConnConfigId                 *string `json:"ConnConfigId" name:"ConnConfigId"`
			SourceType                   *string `json:"SourceType" name:"SourceType"`
			SubTask                      *string `json:"SubTask" name:"SubTask"`
			StartTime                    *string `json:"StartTime" name:"StartTime"`
			MeasureValue                 *int    `json:"MeasureValue" name:"MeasureValue"`
			SubTasks                     []struct {
				Id           *string `json:"Id" name:"Id"`
				TaskId       *string `json:"TaskId" name:"TaskId"`
				ConnConfigId *string `json:"ConnConfigId" name:"ConnConfigId"`
				Name         *string `json:"Name" name:"Name"`
				ActionOn     *string `json:"ActionOn" name:"ActionOn"`
				Status       *string `json:"Status" name:"Status"`
				AgentStage   *string `json:"AgentStage" name:"AgentStage"`
				Judging      *int    `json:"Judging" name:"Judging"`
				Message      *string `json:"Message" name:"Message"`
				Node         *string `json:"Node" name:"Node"`
				Region       *string `json:"Region" name:"Region"`
				Created      *string `json:"Created" name:"Created"`
				Updated      *string `json:"Updated" name:"Updated"`
				Deleted      *int    `json:"Deleted" name:"Deleted"`
				Params       *string `json:"Params" name:"Params"`
				Progress     *int    `json:"Progress" name:"Progress"`
				Latency      *int    `json:"Latency" name:"Latency"`
				AccountId    *string `json:"AccountId" name:"AccountId"`
				StartTime    *string `json:"StartTime" name:"StartTime"`
				EndTime      *string `json:"EndTime" name:"EndTime"`
				ReloadTimes  *int    `json:"ReloadTimes" name:"ReloadTimes"`
				FailureNum   *int    `json:"FailureNum" name:"FailureNum"`
				ErrSkip      *int    `json:"ErrSkip" name:"ErrSkip"`
			} `json:"SubTasks"`
		} `json:"Tasks" name:"Tasks"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Marker     *int `json:"Marker" name:"Marker"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId     *string `json:"TaskId,omitempty" name:"TaskId"`
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	TaskType   *string `json:"TaskType,omitempty" name:"TaskType"`
	ErrSkip    *int    `json:"ErrSkip,omitempty" name:"ErrSkip"`
}

func (r *OperateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OperateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *OperateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnConfigRequest struct {
	*ksyunhttp.BaseRequest
	ConnConfigId *string `json:"ConnConfigId,omitempty" name:"ConnConfigId"`
}

func (r *DescribeConnConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeConnConfigResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		ConnConfig struct {
			Id               *string `json:"Id" name:"Id"`
			AccountId        *string `json:"AccountId" name:"AccountId"`
			Region           *string `json:"Region" name:"Region"`
			SourceRegion     *string `json:"sourceRegion" name:"sourceRegion"`
			TargetRegion     *string `json:"targetRegion" name:"targetRegion"`
			SourceHost       *string `json:"SourceHost" name:"SourceHost"`
			SourcePort       *string `json:"SourcePort" name:"SourcePort"`
			SourceUsername   *string `json:"SourceUsername" name:"SourceUsername"`
			SourceInstanceId *string `json:"SourceInstanceId" name:"SourceInstanceId"`
			TargetInstanceId *string `json:"TargetInstanceId" name:"TargetInstanceId"`
			TargetTopic      *string `json:"TargetTopic" name:"TargetTopic"`
			SourceType       *string `json:"SourceType" name:"SourceType"`
			TargetType       *string `json:"TargetType" name:"TargetType"`
			DbSchema         struct {
				IsFull *bool `json:"IsFull" name:"IsFull"`
			} `json:"DbSchema"`
		} `json:"ConnConfig" name:"ConnConfig"`
	} `json:"Data"`
}

func (r *DescribeConnConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrecheckRequest struct {
	*ksyunhttp.BaseRequest
	PrecheckId *string `json:"PrecheckId,omitempty" name:"PrecheckId"`
}

func (r *DescribePrecheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribePrecheckResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Progress    *int `json:"Progress" name:"Progress"`
		SubPrecheck []struct {
			Id          *string `json:"Id" name:"Id"`
			PrecheckId  *string `json:"PrecheckId" name:"PrecheckId"`
			Name        *string `json:"Name" name:"Name"`
			Status      *string `json:"Status" name:"Status"`
			Description *string `json:"Description" name:"Description"`
			ChineseName *string `json:"ChineseName" name:"ChineseName"`
			Solution    *string `json:"Solution" name:"Solution"`
		} `json:"SubPrecheck" name:"SubPrecheck"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribePrecheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrecheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceUserConfigRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeSourceUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSourceUserConfigResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SourceUserConfig []*string `json:"SourceUserConfig" name:"SourceUserConfig"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeSourceUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetConsistencyCheckRequest struct {
	*ksyunhttp.BaseRequest
	TaskId                    *string `json:"TaskId,omitempty" name:"TaskId"`
	ConsistencyCheckAuto      *bool   `json:"ConsistencyCheckAuto,omitempty" name:"ConsistencyCheckAuto"`
	ConsistencyCheckCycle     *int    `json:"ConsistencyCheckCycle,omitempty" name:"ConsistencyCheckCycle"`
	ConsistencyCheckFixedTime *string `json:"ConsistencyCheckFixedTime,omitempty" name:"ConsistencyCheckFixedTime"`
}

func (r *SetConsistencyCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetConsistencyCheckResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetConsistencyCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetConsistencyCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsistencyCheckRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeConsistencyCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeConsistencyCheckResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		TaskId                    *string `json:"TaskId" name:"TaskId"`
		ConsistencyCheckCycle     *int    `json:"ConsistencyCheckCycle" name:"ConsistencyCheckCycle"`
		ConsistencyCheckFixedTime *string `json:"ConsistencyCheckFixedTime" name:"ConsistencyCheckFixedTime"`
		ConsistencyCheckAuto      *string `json:"ConsistencyCheckAuto" name:"ConsistencyCheckAuto"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeConsistencyCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsistencyCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDTSParameterRequest struct {
	*ksyunhttp.BaseRequest
	SourceType       *string `json:"SourceType,omitempty" name:"SourceType"`
	TargetType       *string `json:"TargetType,omitempty" name:"TargetType"`
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`
	TargetRegion     *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	SourceHost       *string `json:"SourceHost,omitempty" name:"SourceHost"`
	SourcePort       *string `json:"SourcePort,omitempty" name:"SourcePort"`
	SourceUsername   *string `json:"SourceUsername,omitempty" name:"SourceUsername"`
	SourcePassword   *string `json:"SourcePassword,omitempty" name:"SourcePassword"`
	SourceRegion     *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

func (r *DescribeDTSParameterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDTSParameterResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DTSParameter []struct {
			DBParameter                *string   `json:"DBParameter" name:"DBParameter"`
			DBParameterDefault         *string   `json:"DBParameterDefault" name:"DBParameterDefault"`
			DBParameterType            *string   `json:"DBParameterType" name:"DBParameterType"`
			DBParameterMax             *string   `json:"DBParameterMax" name:"DBParameterMax"`
			DBParameterMin             *string   `json:"DBParameterMin" name:"DBParameterMin"`
			DBParameterRestartRequired *bool     `json:"DBParameterRestartRequired" name:"DBParameterRestartRequired"`
			SourceDBParameterValue     *string   `json:"SourceDBParameterValue" name:"SourceDBParameterValue"`
			TargetDBParameterValue     *string   `json:"TargetDBParameterValue" name:"TargetDBParameterValue"`
			DBParameterEnums           []*string `json:"DBParameterEnums" name:"DBParameterEnums"`
		} `json:"DTSParameter" name:"DTSParameter"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDTSParameterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDTSParameterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDTSParameterConfigRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDTSParameterConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDTSParameterConfigResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DTSParameterConfig []*string `json:"DTSParameterConfig" name:"DTSParameterConfig"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDTSParameterConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDTSParameterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceUserRequest struct {
	*ksyunhttp.BaseRequest
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	SourceType       *string `json:"SourceType,omitempty" name:"SourceType"`
	SourceUsername   *string `json:"SourceUsername,omitempty" name:"SourceUsername"`
	SourcePassword   *string `json:"SourcePassword,omitempty" name:"SourcePassword"`
	SourceRegion     *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	TargetType       *string `json:"TargetType,omitempty" name:"TargetType"`
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`
	TargetRegion     *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
}

func (r *DescribeSourceUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSourceUserResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SourceUser []struct {
			Username      *string   `json:"Username" name:"Username"`
			Host          *string   `json:"Host" name:"Host"`
			Description   *string   `json:"Description" name:"Description"`
			Priv          []*string `json:"Priv" name:"Priv"`
			UserAvailable *string   `json:"UserAvailable" name:"UserAvailable"`
		} `json:"SourceUser" name:"SourceUser"`
	} `json:"Data"`
}

func (r *DescribeSourceUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId       *string `json:"TaskId,omitempty" name:"TaskId"`
	SubTaskId    *string `json:"SubTaskId,omitempty" name:"SubTaskId"`
	SubTaskName  *string `json:"SubTaskName,omitempty" name:"SubTaskName"`
	OrderByType  *string `json:"OrderByType,omitempty" name:"OrderByType"`
	ObjectStatus *string `json:"ObjectStatus,omitempty" name:"ObjectStatus"`
	Marker       *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords   *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeSubTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSubTaskResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Id       *string `json:"Id" name:"Id"`
		Status   *string `json:"Status" name:"Status"`
		Created  *string `json:"Created" name:"Created"`
		Updated  *string `json:"Updated" name:"Updated"`
		Progress *int    `json:"Progress" name:"Progress"`
		Object   []struct {
			ObjectName         *string `json:"ObjectName" name:"ObjectName"`
			ObjectType         *string `json:"ObjectType" name:"ObjectType"`
			ObjectStatus       *string `json:"ObjectStatus" name:"ObjectStatus"`
			ObjectSourceSchema *string `json:"ObjectSourceSchema" name:"ObjectSourceSchema"`
			ObjectTargetSchema *string `json:"ObjectTargetSchema" name:"ObjectTargetSchema"`
			ObjectTotalCount   *int    `json:"ObjectTotalCount" name:"ObjectTotalCount"`
			ObjectTargetCount  *int    `json:"ObjectTargetCount" name:"ObjectTargetCount"`
			TimeUsed           *int    `json:"TimeUsed" name:"TimeUsed"`
		} `json:"Object" name:"Object"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Marker     *int `json:"Marker" name:"Marker"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
	} `json:"Data"`
}

func (r *DescribeSubTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsistencyCheckRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CreateConsistencyCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateConsistencyCheckResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateConsistencyCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsistencyCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopConsistencyCheckRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopConsistencyCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopConsistencyCheckResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StopConsistencyCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopConsistencyCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfigRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeRegionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRegionConfigResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RegionConfig struct {
			SubscriptionKrds []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"SubscriptionKrds"`
			ClusterKcs []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"ClusterKcs"`
			Kpg []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"Kpg"`
			PublicPg []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"PublicPg"`
			Krds []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"Krds"`
			Kmgo []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"Kmgo"`
			PublicRedis []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"PublicRedis"`
			Kcs []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"Kcs"`
			Public []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"Public"`
			PublicMgo []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"PublicMgo"`
			PublicClusterRedis []struct {
				RegionCode   *string `json:"RegionCode" name:"RegionCode"`
				RegionName   *string `json:"RegionName" name:"RegionName"`
				RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
				AreaCode     *string `json:"AreaCode" name:"AreaCode"`
				AreaName     *string `json:"AreaName" name:"AreaName"`
				AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			} `json:"PublicClusterRedis"`
		} `json:"RegionConfig" name:"RegionConfig"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeRegionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskBirdViewRequest struct {
	*ksyunhttp.BaseRequest
	TaskType *string `json:"taskType,omitempty" name:"taskType"`
}

func (r *TaskBirdViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type TaskBirdViewResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Global struct {
			Total     *int `json:"total" name:"total"`
			Running   *int `json:"running" name:"running"`
			Finished  *int `json:"finished" name:"finished"`
			Erroneous *int `json:"erroneous" name:"erroneous"`
			Others    *int `json:"others" name:"others"`
		} `json:"Global" name:"Global"`
		Regions []struct {
			RegionId   *int    `json:"regionId" name:"regionId"`
			RegionCode *string `json:"regionCode" name:"regionCode"`
			RegionName *string `json:"regionName" name:"regionName"`
			Total      *int    `json:"total" name:"total"`
			Statistic  struct {
				Running   *int `json:"Running" name:"Running"`
				Finished  *int `json:"Finished" name:"Finished"`
				Erroneous *int `json:"Erroneous" name:"Erroneous"`
				Others    *int `json:"Others" name:"Others"`
			} `json:"Statistic"`
		} `json:"Regions" name:"Regions"`
	} `json:"Data"`
}

func (r *TaskBirdViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TaskBirdViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
