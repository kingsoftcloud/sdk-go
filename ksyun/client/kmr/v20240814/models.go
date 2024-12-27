package v20240814
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DetailWorkspaceRequest struct {
    *ksyunhttp.BaseRequest
    WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
}

func (r *DetailWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetailWorkspaceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetailWorkspaceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DetailWorkspaceResponse struct {
    *ksyunhttp.BaseResponse
    Code *int `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		WorkspaceId *string `json:"WorkspaceId"`
		WorkspaceName *string `json:"WorkspaceName"`
		Status *string `json:"Status"`
		PaymentType *string `json:"PaymentType"`
		CreateTime *string `json:"CreateTime"`
		RunningTime *int `json:"RunningTime"`
		AccountId *string `json:"AccountId"`
		Region *string `json:"Region"`
		ResourceSpec *int `json:"ResourceSpec"`
		FixedResourceSpec *int `json:"FixedResourceSpec"`
		ElasticResourceSpec *int `json:"ElasticResourceSpec"`
		ElasticResourceInstanceId *string `json:"ElasticResourceInstanceId"`
		ElasticResourceStatus *string `json:"ElasticResourceStatus"`
		ComputeUnit *string `json:"ComputeUnit"`
		HistoryServer *string `json:"HistoryServer"`
		AllocatedResources *int `json:"AllocatedResources"`
		EndPoint *string `json:"EndPoint"`
		UsedResources *int `json:"UsedResources"`
		MaxCu *int `json:"MaxCu"`
		MinCu *int `json:"MinCu"`
		RunningJob *bool `json:"RunningJob"`
	} `json:"Data"`
}

func (r *DetailWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetailWorkspaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListWorkspacesRequest struct {
    *ksyunhttp.BaseRequest
    NameOrId *string `json:"NameOrId,omitempty" name:"NameOrId"`
    Status []*string `json:"Status,omitempty" name:"Status"`
    PageNumber *int `json:"PageNumber,omitempty" name:"PageNumber"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListWorkspacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListWorkspacesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListWorkspacesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListWorkspacesResponse struct {
    *ksyunhttp.BaseResponse
    Code *int `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		TotalCount *int `json:"TotalCount"`
		Workspaces []struct {
					WorkspaceId *string `json:"WorkspaceId"`
					WorkspaceName *string `json:"WorkspaceName"`
					Status *string `json:"Status"`
					PaymentType *string `json:"PaymentType"`
					CreateTime *string `json:"CreateTime"`
					RunningTime *int `json:"RunningTime"`
					AccountId *string `json:"AccountId"`
					Region *string `json:"Region"`
					ResourceSpec *int `json:"ResourceSpec"`
					FixedResourceSpec *int `json:"FixedResourceSpec"`
					ElasticResourceSpec *int `json:"ElasticResourceSpec"`
					ElasticResourceInstanceId *string `json:"ElasticResourceInstanceId"`
					ElasticResourceStatus *string `json:"ElasticResourceStatus"`
					ComputeUnit *string `json:"ComputeUnit"`
					HistoryServer *string `json:"HistoryServer"`
					AllocatedResources *int `json:"AllocatedResources"`
			} `json:"Workspaces"`
		} `json:"Data"`
}

func (r *ListWorkspacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListWorkspacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartJobRunRequest struct {
    *ksyunhttp.BaseRequest
    WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
    AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
    AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
    ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

func (r *StartJobRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartJobRunRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartJobRunRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StartJobRunResponse struct {
    *ksyunhttp.BaseResponse
    Code *int `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		JobRunId *string `json:"JobRunId"`
		WorkspaceId *string `json:"WorkspaceId"`
		Name *string `json:"Name"`
		CodeType *string `json:"CodeType"`
		State *string `json:"State"`
		ErrorReason *string `json:"ErrorReason"`
		SubmitTime *string `json:"SubmitTime"`
		StartTime *string `json:"StartTime"`
		EndTime *string `json:"EndTime"`
		RunningTime *int `json:"RunningTime"`
		WebUI *string `json:"WebUI"`
		ReleaseVersion *string `json:"ReleaseVersion"`
		SparkSubmitData struct {
				Name *string `json:"Name"`
				SparkDriverCores *int `json:"SparkDriverCores"`
				SparkDriverMemory *string `json:"SparkDriverMemory"`
				SparkExecutorCores *int `json:"SparkExecutorCores"`
				SparkExecutorMemory *string `json:"SparkExecutorMemory"`
				SparkNumExecutors *int `json:"SparkNumExecutors"`
				Class *string `json:"Class"`
				AppResource *string `json:"AppResource"`
			ExtraArgs []struct {
			} `json:"ExtraArgs"`
				Conf *string `json:"Conf"`
				Jars *string `json:"Jars"`
				Files *string `json:"Files"`
				PyFiles *string `json:"PyFiles"`
				Archives *string `json:"Archives"`
				Packages *string `json:"Packages"`
				CacheFile *string `json:"CacheFile"`
				Image *string `json:"Image"`
				ProxyUser *string `json:"ProxyUser"`
				Params *string `json:"Params"`
		} `json:"SparkSubmitData"`
		ResourceUsage *int `json:"ResourceUsage"`
	} `json:"Data"`
}

func (r *StartJobRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartJobRunResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetJobRunRequest struct {
    *ksyunhttp.BaseRequest
    WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
    JobRunId *string `json:"JobRunId,omitempty" name:"JobRunId"`
}

func (r *GetJobRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetJobRunRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetJobRunRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetJobRunResponse struct {
    *ksyunhttp.BaseResponse
    Code *int `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		JobRunId *string `json:"JobRunId"`
		WorkspaceId *string `json:"WorkspaceId"`
		Name *string `json:"Name"`
		CodeType *string `json:"CodeType"`
		State *string `json:"State"`
		ErrorReason *string `json:"ErrorReason"`
		SubmitTime *string `json:"SubmitTime"`
		StartTime *string `json:"StartTime"`
		EndTime *string `json:"EndTime"`
		RunningTime *int `json:"RunningTime"`
		WebUI *string `json:"WebUI"`
		ReleaseVersion *string `json:"ReleaseVersion"`
		SparkSubmitData struct {
				Name *string `json:"Name"`
				SparkDriverCores *int `json:"SparkDriverCores"`
				SparkDriverMemory *string `json:"SparkDriverMemory"`
				SparkExecutorCores *int `json:"SparkExecutorCores"`
				SparkExecutorMemory *string `json:"SparkExecutorMemory"`
				SparkNumExecutors *int `json:"SparkNumExecutors"`
				Class *string `json:"Class"`
				AppResource *string `json:"AppResource"`
			ExtraArgs []struct {
			} `json:"ExtraArgs"`
				Conf *string `json:"Conf"`
				Jars *string `json:"Jars"`
				Files *string `json:"Files"`
				PyFiles *string `json:"PyFiles"`
				Archives *string `json:"Archives"`
				Packages *string `json:"Packages"`
				CacheFile *string `json:"CacheFile"`
				Image *string `json:"Image"`
				ProxyUser *string `json:"ProxyUser"`
				Params *string `json:"Params"`
		} `json:"SparkSubmitData"`
		ResourceUsage *int `json:"ResourceUsage"`
	} `json:"Data"`
}

func (r *GetJobRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetJobRunResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListJobRunsRequest struct {
    *ksyunhttp.BaseRequest
    WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *ListJobRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListJobRunsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListJobRunsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListJobRunsResponse struct {
    *ksyunhttp.BaseResponse
    Code *int `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		JobRuns []struct {
					AppId *string `json:"AppId"`
					AppName *string `json:"AppName"`
					AppStatus *string `json:"AppStatus"`
					jobRunId *string `json:"jobRunId"`
					SubmitTime *string `json:"SubmitTime"`
					StartTime *string `json:"StartTime"`
					EndTime *string `json:"EndTime"`
					ReleaseVersion *string `json:"ReleaseVersion"`
					UsageResource *int `json:"UsageResource"`
			} `json:"JobRuns"`
			MaxResults *int `json:"MaxResults"`
			TotalCount *int `json:"TotalCount"`
		} `json:"Data"`
}

func (r *ListJobRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListJobRunsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelJobRunRequest struct {
    *ksyunhttp.BaseRequest
    WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
    JobRunId *string `json:"JobRunId,omitempty" name:"JobRunId"`
}

func (r *CancelJobRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelJobRunRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelJobRunRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CancelJobRunResponse struct {
    *ksyunhttp.BaseResponse
    Code *int `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		JobRunId *string `json:"JobRunId"`
	} `json:"Data"`
}

func (r *CancelJobRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelJobRunResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListExecutorRequest struct {
    *ksyunhttp.BaseRequest
    WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
    JobRunId *string `json:"JobRunId,omitempty" name:"JobRunId"`
    PageNumber *int `json:"PageNumber,omitempty" name:"PageNumber"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListExecutorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListExecutorRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListExecutorRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListExecutorResponse struct {
    *ksyunhttp.BaseResponse
    status *int `json:"status" name:"status"`
	Data struct {
		Code *int `json:"Code"`
		Message *string `json:"Message"`
		RequestId *string `json:"RequestId"`
		Data struct {
				Count *int `json:"Count"`
				Total *int `json:"Total"`
			List []struct {
				Id *string `json:"Id"`
				Address *string `json:"Address"`
				Status *string `json:"Status"`
				RddBlocks *int `json:"RddBlocks"`
				StorageMemory *string `json:"StorageMemory"`
				MaxMemory *string `json:"MaxMemory"`
				DiskUsed *string `json:"DiskUsed"`
				Cores *int `json:"Cores"`
				ActiveTasks *int `json:"ActiveTasks"`
				FailedTasks *int `json:"FailedTasks"`
				CompletedTasks *int `json:"CompletedTasks"`
				TotalTasks *int `json:"TotalTasks"`
				TasksTime *string `json:"TasksTime"`
				TotalGCTime *string `json:"TotalGCTime"`
				TotalInputBytes *string `json:"TotalInputBytes"`
				TotalShuffleRead *string `json:"TotalShuffleRead"`
				TotalShuffleWrite *string `json:"TotalShuffleWrite"`
				PodName *string `json:"PodName"`
			} `json:"List"`
		} `json:"Data"`
	} `json:"Data"`
    message *string `json:"message" name:"message"`
}

func (r *ListExecutorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListExecutorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

