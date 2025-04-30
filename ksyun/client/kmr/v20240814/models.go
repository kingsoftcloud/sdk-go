package v20240814

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type StartJobRunSparkSubmitData struct {
	Name                *string   `json:"Name,omitempty" name:"Name"`
	SparkDriverCores    *int      `json:"SparkDriverCores,omitempty" name:"SparkDriverCores"`
	SparkDriverMemory   *string   `json:"SparkDriverMemory,omitempty" name:"SparkDriverMemory"`
	SparkExecutorCores  *int      `json:"SparkExecutorCores,omitempty" name:"SparkExecutorCores"`
	SparkExecutorMemory *string   `json:"SparkExecutorMemory,omitempty" name:"SparkExecutorMemory"`
	SparkNumExecutors   *int      `json:"SparkNumExecutors,omitempty" name:"SparkNumExecutors"`
	Class               *string   `json:"Class,omitempty" name:"Class"`
	AppResource         *string   `json:"AppResource,omitempty" name:"AppResource"`
	ExtraArgs           []*string `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
	Conf                []*string `json:"Conf,omitempty" name:"Conf"`
	Jars                []*string `json:"Jars,omitempty" name:"Jars"`
	Files               []*string `json:"Files,omitempty" name:"Files"`
	PyFiles             []*string `json:"PyFiles,omitempty" name:"PyFiles"`
	Archives            []*string `json:"Archives,omitempty" name:"Archives"`
	Packages            []*string `json:"Packages,omitempty" name:"Packages"`
	CacheFile           []*string `json:"CacheFile,omitempty" name:"CacheFile"`
	Image               *string   `json:"Image,omitempty" name:"Image"`
	HighPriority        *bool     `json:"HighPriority,omitempty" name:"HighPriority"`
}
type CancelJobRunJobRunIds struct {
	JobRunId *string `json:"JobRunId,omitempty" name:"JobRunId"`
	JobType  *string `json:"JobType,omitempty" name:"JobType"`
}
type StartRayJobRunRaySubmitData struct {
	Name               *string `json:"Name,omitempty" name:"Name"`
	RayHeadCores       *int    `json:"RayHeadCores,omitempty" name:"RayHeadCores"`
	RayHeadMemory      *string `json:"RayHeadMemory,omitempty" name:"RayHeadMemory"`
	RayWorkerCores     *int    `json:"RayWorkerCores,omitempty" name:"RayWorkerCores"`
	RayWorkerMemory    *string `json:"RayWorkerMemory,omitempty" name:"RayWorkerMemory"`
	RayWorkerNum       *int    `json:"RayWorkerNum,omitempty" name:"RayWorkerNum"`
	RayWorkerGpus      *int    `json:"RayWorkerGpus,omitempty" name:"RayWorkerGpus"`
	EntrypointCmd      *string `json:"EntrypointCmd,omitempty" name:"EntrypointCmd"`
	EntrypointResource *string `json:"EntrypointResource,omitempty" name:"EntrypointResource"`
	Conf               *string `json:"Conf,omitempty" name:"Conf"`
	Image              *string `json:"Image,omitempty" name:"Image"`
	JuiceFs            *string `json:"JuiceFs,omitempty" name:"JuiceFs"`
	MountPath          *string `json:"MountPath,omitempty" name:"MountPath"`
	RuntimeEnv         *string `json:"RuntimeEnv,omitempty" name:"RuntimeEnv"`
	HighPriority       *bool   `json:"HighPriority,omitempty" name:"HighPriority"`
}
type CancelRayJobRunJobRunIds struct {
	JobRunId *string `json:"JobRunId,omitempty" name:"JobRunId"`
	JobType  *string `json:"JobType,omitempty" name:"JobType"`
}
type StartFlinkJobRunSubmitData struct {
	Name         *string   `json:"Name,omitempty" name:"Name"`
	Image        *string   `json:"Image,omitempty" name:"Image"`
	UpgradeMode  *string   `json:"UpgradeMode,omitempty" name:"UpgradeMode"`
	JobCores     *int      `json:"JobCores,omitempty" name:"JobCores"`
	JobMemory    *string   `json:"JobMemory,omitempty" name:"JobMemory"`
	TaskCores    *int      `json:"TaskCores,omitempty" name:"TaskCores"`
	TaskMemory   *string   `json:"TaskMemory,omitempty" name:"TaskMemory"`
	NumTasks     *int      `json:"NumTasks,omitempty" name:"NumTasks"`
	FlinkConf    []*string `json:"FlinkConf,omitempty" name:"FlinkConf"`
	Dependencies []*string `json:"Dependencies,omitempty" name:"Dependencies"`
	JarUri       *string   `json:"JarUri,omitempty" name:"JarUri"`
	EntryClass   *string   `json:"EntryClass,omitempty" name:"EntryClass"`
	MainArgs     []*string `json:"MainArgs,omitempty" name:"MainArgs"`
}
type CancelFlinkJobRunJobRunIds struct {
	JobRunId *string `json:"JobRunId,omitempty" name:"JobRunId"`
	JobType  *string `json:"JobType,omitempty" name:"JobType"`
}
type QueryMetricsQueryData struct {
	Query *string `json:"Query,omitempty" name:"Query"`
	Start *string `json:"Start,omitempty" name:"Start"`
	End   *string `json:"End,omitempty" name:"End"`
	Step  *string `json:"Step,omitempty" name:"Step"`
}

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
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		WorkspaceId               *string `json:"WorkspaceId" name:"WorkspaceId"`
		WorkspaceName             *string `json:"WorkspaceName" name:"WorkspaceName"`
		Status                    *string `json:"Status" name:"Status"`
		PaymentType               *string `json:"PaymentType" name:"PaymentType"`
		CreateTime                *string `json:"CreateTime" name:"CreateTime"`
		RunningTime               *int    `json:"RunningTime" name:"RunningTime"`
		AccountId                 *string `json:"AccountId" name:"AccountId"`
		Region                    *string `json:"Region" name:"Region"`
		ResourceSpec              *int    `json:"ResourceSpec" name:"ResourceSpec"`
		FixedResourceSpec         *int    `json:"FixedResourceSpec" name:"FixedResourceSpec"`
		ElasticResourceSpec       *int    `json:"ElasticResourceSpec" name:"ElasticResourceSpec"`
		ElasticResourceInstanceId *string `json:"ElasticResourceInstanceId" name:"ElasticResourceInstanceId"`
		ElasticResourceStatus     *string `json:"ElasticResourceStatus" name:"ElasticResourceStatus"`
		ComputeUnit               *string `json:"ComputeUnit" name:"ComputeUnit"`
		HistoryServer             *string `json:"HistoryServer" name:"HistoryServer"`
		AllocatedResources        *int    `json:"AllocatedResources" name:"AllocatedResources"`
		EndPoint                  *string `json:"EndPoint" name:"EndPoint"`
		UsedResources             *int    `json:"UsedResources" name:"UsedResources"`
		MaxCu                     *int    `json:"MaxCu" name:"MaxCu"`
		MinCu                     *int    `json:"MinCu" name:"MinCu"`
		RunningJob                *bool   `json:"RunningJob" name:"RunningJob"`
		GrafanaDashboard          *string `json:"GrafanaDashboard" name:"GrafanaDashboard"`
		FlinkHistoryServer        *string `json:"FlinkHistoryServer" name:"FlinkHistoryServer"`
		FlinkGrafanaDashboard     *string `json:"FlinkGrafanaDashboard" name:"FlinkGrafanaDashboard"`
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
	NameOrId   *string   `json:"NameOrId,omitempty" name:"NameOrId"`
	Status     []*string `json:"Status,omitempty" name:"Status"`
	PageNumber *int      `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize   *int      `json:"PageSize,omitempty" name:"PageSize"`
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
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Workspaces []struct {
			WorkspaceId               *string `json:"WorkspaceId" name:"WorkspaceId"`
			WorkspaceName             *string `json:"WorkspaceName" name:"WorkspaceName"`
			Status                    *string `json:"Status" name:"Status"`
			PaymentType               *string `json:"PaymentType" name:"PaymentType"`
			CreateTime                *string `json:"CreateTime" name:"CreateTime"`
			RunningTime               *int    `json:"RunningTime" name:"RunningTime"`
			AccountId                 *string `json:"AccountId" name:"AccountId"`
			Region                    *string `json:"Region" name:"Region"`
			ResourceSpec              *int    `json:"ResourceSpec" name:"ResourceSpec"`
			FixedResourceSpec         *int    `json:"FixedResourceSpec" name:"FixedResourceSpec"`
			ElasticResourceSpec       *int    `json:"ElasticResourceSpec" name:"ElasticResourceSpec"`
			ElasticResourceInstanceId *string `json:"ElasticResourceInstanceId" name:"ElasticResourceInstanceId"`
			ElasticResourceStatus     *string `json:"ElasticResourceStatus" name:"ElasticResourceStatus"`
			ComputeUnit               *string `json:"ComputeUnit" name:"ComputeUnit"`
			HistoryServer             *string `json:"HistoryServer" name:"HistoryServer"`
			AllocatedResources        *int    `json:"AllocatedResources" name:"AllocatedResources"`
		} `json:"Workspaces" name:"Workspaces"`
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
	WorkspaceId     *string                     `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	AccessKeyId     *string                     `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	AccessKeySecret *string                     `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
	ReleaseVersion  *string                     `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
	SparkSubmitData *StartJobRunSparkSubmitData `json:"SparkSubmitData,omitempty" name:"SparkSubmitData"`
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
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRunId        *string `json:"JobRunId" name:"JobRunId"`
		WorkspaceId     *string `json:"WorkspaceId" name:"WorkspaceId"`
		Name            *string `json:"Name" name:"Name"`
		CodeType        *string `json:"CodeType" name:"CodeType"`
		State           *string `json:"State" name:"State"`
		ErrorReason     *string `json:"ErrorReason" name:"ErrorReason"`
		SubmitTime      *string `json:"SubmitTime" name:"SubmitTime"`
		StartTime       *string `json:"StartTime" name:"StartTime"`
		EndTime         *string `json:"EndTime" name:"EndTime"`
		RunningTime     *int    `json:"RunningTime" name:"RunningTime"`
		WebUI           *string `json:"WebUI" name:"WebUI"`
		ReleaseVersion  *string `json:"ReleaseVersion" name:"ReleaseVersion"`
		SparkSubmitData struct {
			Name                *string   `json:"Name" name:"Name"`
			SparkDriverCores    *int      `json:"SparkDriverCores" name:"SparkDriverCores"`
			SparkDriverMemory   *string   `json:"SparkDriverMemory" name:"SparkDriverMemory"`
			SparkExecutorCores  *int      `json:"SparkExecutorCores" name:"SparkExecutorCores"`
			SparkExecutorMemory *string   `json:"SparkExecutorMemory" name:"SparkExecutorMemory"`
			SparkNumExecutors   *int      `json:"SparkNumExecutors" name:"SparkNumExecutors"`
			Class               *string   `json:"Class" name:"Class"`
			AppResource         *string   `json:"AppResource" name:"AppResource"`
			ExtraArgs           []*string `json:"ExtraArgs" name:"ExtraArgs"`
			Conf                *string   `json:"Conf" name:"Conf"`
			Jars                *string   `json:"Jars" name:"Jars"`
			Files               *string   `json:"Files" name:"Files"`
			PyFiles             *string   `json:"PyFiles" name:"PyFiles"`
			Archives            *string   `json:"Archives" name:"Archives"`
			Packages            *string   `json:"Packages" name:"Packages"`
			CacheFile           *string   `json:"CacheFile" name:"CacheFile"`
			Image               *string   `json:"Image" name:"Image"`
			ProxyUser           *string   `json:"ProxyUser" name:"ProxyUser"`
			Params              *string   `json:"Params" name:"Params"`
		} `json:"SparkSubmitData" name:"SparkSubmitData"`
		ResourceUsage *int `json:"ResourceUsage" name:"ResourceUsage"`
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
	JobRunId    *string `json:"JobRunId,omitempty" name:"JobRunId"`
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
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRunId        *string `json:"JobRunId" name:"JobRunId"`
		WorkspaceId     *string `json:"WorkspaceId" name:"WorkspaceId"`
		Name            *string `json:"Name" name:"Name"`
		CodeType        *string `json:"CodeType" name:"CodeType"`
		State           *string `json:"State" name:"State"`
		ErrorReason     *string `json:"ErrorReason" name:"ErrorReason"`
		SubmitTime      *string `json:"SubmitTime" name:"SubmitTime"`
		StartTime       *string `json:"StartTime" name:"StartTime"`
		EndTime         *string `json:"EndTime" name:"EndTime"`
		RunningTime     *int    `json:"RunningTime" name:"RunningTime"`
		WebUI           *string `json:"WebUI" name:"WebUI"`
		ReleaseVersion  *string `json:"ReleaseVersion" name:"ReleaseVersion"`
		SparkSubmitData struct {
			Name                *string   `json:"Name" name:"Name"`
			SparkDriverCores    *int      `json:"SparkDriverCores" name:"SparkDriverCores"`
			SparkDriverMemory   *string   `json:"SparkDriverMemory" name:"SparkDriverMemory"`
			SparkExecutorCores  *int      `json:"SparkExecutorCores" name:"SparkExecutorCores"`
			SparkExecutorMemory *string   `json:"SparkExecutorMemory" name:"SparkExecutorMemory"`
			SparkNumExecutors   *int      `json:"SparkNumExecutors" name:"SparkNumExecutors"`
			Class               *string   `json:"Class" name:"Class"`
			AppResource         *string   `json:"AppResource" name:"AppResource"`
			ExtraArgs           []*string `json:"ExtraArgs" name:"ExtraArgs"`
			Conf                *string   `json:"Conf" name:"Conf"`
			Jars                *string   `json:"Jars" name:"Jars"`
			Files               *string   `json:"Files" name:"Files"`
			PyFiles             *string   `json:"PyFiles" name:"PyFiles"`
			Archives            *string   `json:"Archives" name:"Archives"`
			Packages            *string   `json:"Packages" name:"Packages"`
			CacheFile           *string   `json:"CacheFile" name:"CacheFile"`
			Image               *string   `json:"Image" name:"Image"`
			ProxyUser           *string   `json:"ProxyUser" name:"ProxyUser"`
			Params              *string   `json:"Params" name:"Params"`
		} `json:"SparkSubmitData" name:"SparkSubmitData"`
		ResourceUsage *int `json:"ResourceUsage" name:"ResourceUsage"`
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
	MaxResults  *int    `json:"MaxResults,omitempty" name:"MaxResults"`
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
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRuns []struct {
			AppId          *string `json:"AppId" name:"AppId"`
			AppName        *string `json:"AppName" name:"AppName"`
			AppStatus      *string `json:"AppStatus" name:"AppStatus"`
			JobRunId       *string `json:"jobRunId" name:"jobRunId"`
			SubmitTime     *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime      *string `json:"StartTime" name:"StartTime"`
			EndTime        *string `json:"EndTime" name:"EndTime"`
			ReleaseVersion *string `json:"ReleaseVersion" name:"ReleaseVersion"`
			UsageResource  *int    `json:"UsageResource" name:"UsageResource"`
		} `json:"JobRuns" name:"JobRuns"`
		MaxResults *int `json:"MaxResults" name:"MaxResults"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
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
	WorkspaceId *string                  `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunIds   []*CancelJobRunJobRunIds `json:"JobRunIds,omitempty" name:"JobRunIds"`
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
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
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
	JobRunId    *string `json:"JobRunId,omitempty" name:"JobRunId"`
	PageNumber  *int    `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize    *int    `json:"PageSize,omitempty" name:"PageSize"`
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
	Status *int `json:"status" name:"status"`
	Data   struct {
		Code      *int    `json:"Code" name:"Code"`
		Message   *string `json:"Message" name:"Message"`
		RequestId *string `json:"RequestId" name:"RequestId"`
		Data      struct {
			Count *int `json:"Count" name:"Count"`
			Total *int `json:"Total" name:"Total"`
			List  []struct {
				Id                *string `json:"Id" name:"Id"`
				Address           *string `json:"Address" name:"Address"`
				Status            *string `json:"Status" name:"Status"`
				RddBlocks         *int    `json:"RddBlocks" name:"RddBlocks"`
				StorageMemory     *string `json:"StorageMemory" name:"StorageMemory"`
				MaxMemory         *string `json:"MaxMemory" name:"MaxMemory"`
				DiskUsed          *string `json:"DiskUsed" name:"DiskUsed"`
				Cores             *int    `json:"Cores" name:"Cores"`
				ActiveTasks       *int    `json:"ActiveTasks" name:"ActiveTasks"`
				FailedTasks       *int    `json:"FailedTasks" name:"FailedTasks"`
				CompletedTasks    *int    `json:"CompletedTasks" name:"CompletedTasks"`
				TotalTasks        *int    `json:"TotalTasks" name:"TotalTasks"`
				TasksTime         *string `json:"TasksTime" name:"TasksTime"`
				TotalGCTime       *string `json:"TotalGCTime" name:"TotalGCTime"`
				TotalInputBytes   *string `json:"TotalInputBytes" name:"TotalInputBytes"`
				TotalShuffleRead  *string `json:"TotalShuffleRead" name:"TotalShuffleRead"`
				TotalShuffleWrite *string `json:"TotalShuffleWrite" name:"TotalShuffleWrite"`
				PodName           *string `json:"PodName" name:"PodName"`
			} `json:"List"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
	Message *string `json:"message" name:"message"`
}

func (r *ListExecutorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListExecutorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartRayJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId     *string                      `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	AccessKeyId     *string                      `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	AccessKeySecret *string                      `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
	ReleaseVersion  *string                      `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
	RaySubmitData   *StartRayJobRunRaySubmitData `json:"RaySubmitData,omitempty" name:"RaySubmitData"`
}

func (r *StartRayJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartRayJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartRayJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartRayJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRunId      *string `json:"JobRunId" name:"JobRunId"`
		WorkspaceId   *string `json:"WorkspaceId" name:"WorkspaceId"`
		Name          *string `json:"Name" name:"Name"`
		State         *string `json:"State" name:"State"`
		ErrorReason   *string `json:"ErrorReason" name:"ErrorReason"`
		EndTime       *string `json:"EndTime" name:"EndTime"`
		RaySubmitData struct {
			Name               *string `json:"Name" name:"Name"`
			RayHeadCores       *int    `json:"RayHeadCores" name:"RayHeadCores"`
			RayHeadMemory      *string `json:"RayHeadMemory" name:"RayHeadMemory"`
			RayWorkerCores     *int    `json:"RayWorkerCores" name:"RayWorkerCores"`
			RayWorkerMemory    *string `json:"RayWorkerMemory" name:"RayWorkerMemory"`
			RayWorkerNum       *int    `json:"RayWorkerNum" name:"RayWorkerNum"`
			RayWorkerGpus      *int    `json:"RayWorkerGpus" name:"RayWorkerGpus"`
			EntrypointCmd      *string `json:"EntrypointCmd" name:"EntrypointCmd"`
			EntrypointResource *string `json:"EntrypointResource" name:"EntrypointResource"`
			RuntimeEnv         *string `json:"RuntimeEnv" name:"RuntimeEnv"`
			Conf               *string `json:"Conf" name:"Conf"`
			Image              *string `json:"Image" name:"Image"`
			JuiceFs            *string `json:"JuiceFs" name:"JuiceFs"`
			MountPath          *string `json:"MountPath" name:"MountPath"`
		} `json:"RaySubmitData" name:"RaySubmitData"`
	} `json:"Data"`
}

func (r *StartRayJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartRayJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRayJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunId    *string `json:"JobRunId,omitempty" name:"JobRunId"`
}

func (r *GetRayJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRayJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetRayJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRayJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRunId       *string `json:"JobRunId" name:"JobRunId"`
		WorkspaceId    *string `json:"WorkspaceId" name:"WorkspaceId"`
		Name           *string `json:"Name" name:"Name"`
		State          *string `json:"State" name:"State"`
		ErrorReason    *string `json:"ErrorReason" name:"ErrorReason"`
		SubmitTime     *string `json:"SubmitTime" name:"SubmitTime"`
		StartTime      *string `json:"StartTime" name:"StartTime"`
		EndTime        *string `json:"EndTime" name:"EndTime"`
		RunningTime    *int    `json:"RunningTime" name:"RunningTime"`
		WebUI          *string `json:"WebUI" name:"WebUI"`
		RayLog         *string `json:"RayLog" name:"RayLog"`
		ReleaseVersion *string `json:"ReleaseVersion" name:"ReleaseVersion"`
		RaySubmitData  struct {
			Name               *string `json:"Name" name:"Name"`
			RayHeadCores       *int    `json:"RayHeadCores" name:"RayHeadCores"`
			RayHeadMemory      *string `json:"RayHeadMemory" name:"RayHeadMemory"`
			RayWorkerCores     *int    `json:"RayWorkerCores" name:"RayWorkerCores"`
			RayWorkerMemory    *string `json:"RayWorkerMemory" name:"RayWorkerMemory"`
			RayWorkerNum       *int    `json:"RayWorkerNum" name:"RayWorkerNum"`
			RayWorkerGpus      *int    `json:"RayWorkerGpus" name:"RayWorkerGpus"`
			EntrypointCmd      *string `json:"EntrypointCmd" name:"EntrypointCmd"`
			EntrypointResource *string `json:"EntrypointResource" name:"EntrypointResource"`
			RuntimeEnv         *string `json:"RuntimeEnv" name:"RuntimeEnv"`
			Conf               *string `json:"Conf" name:"Conf"`
			Image              *string `json:"Image" name:"Image"`
			JuiceFs            *string `json:"JuiceFs" name:"JuiceFs"`
			MountPath          *string `json:"MountPath" name:"MountPath"`
		} `json:"RaySubmitData" name:"RaySubmitData"`
		ResourceUsage *int `json:"ResourceUsage" name:"ResourceUsage"`
	} `json:"Data"`
}

func (r *GetRayJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRayJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRayJobRunsRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string   `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	NameOrId    *string   `json:"NameOrId,omitempty" name:"NameOrId"`
	Status      []*string `json:"Status,omitempty" name:"Status"`
	PageNumber  *int      `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize    *int      `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListRayJobRunsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRayJobRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRayJobRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListRayJobRunsResponse struct {
	*ksyunhttp.BaseResponse
	Status *int `json:"status" name:"status"`
	Data   struct {
		Code      *int    `json:"Code" name:"Code"`
		Message   *string `json:"Message" name:"Message"`
		RequestId *string `json:"RequestId" name:"RequestId"`
		Data      struct {
			MaxResults *int `json:"MaxResults" name:"MaxResults"`
			TotalCount *int `json:"TotalCount" name:"TotalCount"`
			JobRuns    []struct {
				JobRunId       *string `json:"JobRunId" name:"JobRunId"`
				WorkspaceId    *string `json:"WorkspaceId" name:"WorkspaceId"`
				Name           *string `json:"Name" name:"Name"`
				State          *string `json:"State" name:"State"`
				ErrorReason    *string `json:"ErrorReason" name:"ErrorReason"`
				SubmitTime     *string `json:"SubmitTime" name:"SubmitTime"`
				StartTime      *string `json:"StartTime" name:"StartTime"`
				EndTime        *string `json:"EndTime" name:"EndTime"`
				RunningTime    *int    `json:"RunningTime" name:"RunningTime"`
				WebUI          *string `json:"WebUI" name:"WebUI"`
				RayLog         *string `json:"RayLog" name:"RayLog"`
				ReleaseVersion *string `json:"ReleaseVersion" name:"ReleaseVersion"`
				RaySubmitData  struct {
					Name               *string   `json:"Name" name:"Name"`
					RayHeadCores       *int      `json:"RayHeadCores" name:"RayHeadCores"`
					RayHeadMemory      *string   `json:"RayHeadMemory" name:"RayHeadMemory"`
					RayWorkerCores     *int      `json:"RayWorkerCores" name:"RayWorkerCores"`
					RayWorkerMemory    *string   `json:"RayWorkerMemory" name:"RayWorkerMemory"`
					RayWorkerNum       *int      `json:"RayWorkerNum" name:"RayWorkerNum"`
					RayWorkerGpus      *int      `json:"RayWorkerGpus" name:"RayWorkerGpus"`
					EntrypointCmd      *string   `json:"EntrypointCmd" name:"EntrypointCmd"`
					EntrypointResource *string   `json:"EntrypointResource" name:"EntrypointResource"`
					RuntimeEnv         *string   `json:"RuntimeEnv" name:"RuntimeEnv"`
					Conf               []*string `json:"Conf" name:"Conf"`
					Image              *string   `json:"Image" name:"Image"`
					JuiceFs            *string   `json:"JuiceFs" name:"JuiceFs"`
					MountPath          *string   `json:"MountPath" name:"MountPath"`
				} `json:"RaySubmitData" name:"RaySubmitData"`
				ResourceUsage *int `json:"ResourceUsage" name:"ResourceUsage"`
			} `json:"JobRuns"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *ListRayJobRunsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRayJobRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelRayJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string                     `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunIds   []*CancelRayJobRunJobRunIds `json:"JobRunIds,omitempty" name:"JobRunIds"`
}

func (r *CancelRayJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelRayJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelRayJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelRayJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *CancelRayJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelRayJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartFlinkJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId     *string                     `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	AccessKeyId     *string                     `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	AccessKeySecret *string                     `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
	ReleaseVersion  *string                     `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
	SubmitData      *StartFlinkJobRunSubmitData `json:"SubmitData,omitempty" name:"SubmitData"`
}

func (r *StartFlinkJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartFlinkJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartFlinkJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartFlinkJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRunId            *string `json:"JobRunId" name:"JobRunId"`
		WorkspaceId         *string `json:"WorkspaceId" name:"WorkspaceId"`
		Name                *string `json:"Name" name:"Name"`
		State               *string `json:"State" name:"State"`
		ErrorReason         *string `json:"ErrorReason" name:"ErrorReason"`
		EndTime             *string `json:"EndTime" name:"EndTime"`
		FlinkDeploymentData struct {
			Name                *string   `json:"Name" name:"Name"`
			Image               *string   `json:"Image" name:"Image"`
			UpgradeMode         *string   `json:"UpgradeMode" name:"UpgradeMode"`
			JobCores            *int      `json:"JobCores" name:"JobCores"`
			JobMemory           *string   `json:"JobMemory" name:"JobMemory"`
			TaskCores           *int      `json:"TaskCores" name:"TaskCores"`
			TaskMemory          *string   `json:"TaskMemory" name:"TaskMemory"`
			NumTasks            *int      `json:"NumTasks" name:"NumTasks"`
			JarUri              *string   `json:"JarUri" name:"JarUri"`
			Dependencies        []*string `json:"Dependencies" name:"Dependencies"`
			FlinkDeploymentData struct {
				Name         *string   `json:"Name" name:"Name"`
				Image        *string   `json:"Image" name:"Image"`
				UpgradeMode  *string   `json:"UpgradeMode" name:"UpgradeMode"`
				JobCores     *int      `json:"JobCores" name:"JobCores"`
				JobMemory    *string   `json:"JobMemory" name:"JobMemory"`
				NumTasks     *int      `json:"NumTasks" name:"NumTasks"`
				FlinkConf    []*string `json:"FlinkConf" name:"FlinkConf"`
				Dependencies []*string `json:"Dependencies" name:"Dependencies"`
				JarUri       *string   `json:"JarUri" name:"JarUri"`
				EntryClass   *string   `json:"EntryClass" name:"EntryClass"`
				JarUri1      *string   `json:"JarUri1" name:"JarUri1"`
				MainArgs     *string   `json:"MainArgs" name:"MainArgs"`
			} `json:"FlinkDeploymentData"`
		} `json:"FlinkDeploymentData" name:"FlinkDeploymentData"`
	} `json:"Data"`
}

func (r *StartFlinkJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartFlinkJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlinkJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunId    *string `json:"JobRunId,omitempty" name:"JobRunId"`
}

func (r *GetFlinkJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlinkJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetFlinkJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetFlinkJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		JobRunId            *string `json:"JobRunId" name:"JobRunId"`
		WorkspaceId         *string `json:"WorkspaceId" name:"WorkspaceId"`
		Name                *string `json:"Name" name:"Name"`
		State               *string `json:"State" name:"State"`
		ErrorReason         *string `json:"ErrorReason" name:"ErrorReason"`
		SubmitTime          *string `json:"SubmitTime" name:"SubmitTime"`
		StartTime           *string `json:"StartTime" name:"StartTime"`
		EndTime             *string `json:"EndTime" name:"EndTime"`
		RunningTime         *int    `json:"RunningTime" name:"RunningTime"`
		WebUI               *string `json:"WebUI" name:"WebUI"`
		FlinkLog            *string `json:"FlinkLog" name:"FlinkLog"`
		ReleaseVersion      *string `json:"ReleaseVersion" name:"ReleaseVersion"`
		FlinkDeploymentData struct {
			Name         *string   `json:"Name" name:"Name"`
			Image        *string   `json:"Image" name:"Image"`
			UpgradeMode  *string   `json:"UpgradeMode" name:"UpgradeMode"`
			JobCores     *int      `json:"JobCores" name:"JobCores"`
			JobMemory    *string   `json:"JobMemory" name:"JobMemory"`
			TaskCores    *int      `json:"TaskCores" name:"TaskCores"`
			TaskMemory   *string   `json:"TaskMemory" name:"TaskMemory"`
			NumTasks     *int      `json:"NumTasks" name:"NumTasks"`
			FlinkConf    []*string `json:"FlinkConf" name:"FlinkConf"`
			Dependencies []*string `json:"Dependencies" name:"Dependencies"`
			JarUri       *string   `json:"JarUri" name:"JarUri"`
		} `json:"FlinkDeploymentData" name:"FlinkDeploymentData"`
		ResourceUsage *int `json:"ResourceUsage" name:"ResourceUsage"`
	} `json:"Data"`
}

func (r *GetFlinkJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlinkJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFlinkJobRunsRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string   `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	NameOrId    *string   `json:"NameOrId,omitempty" name:"NameOrId"`
	Status      []*string `json:"Status,omitempty" name:"Status"`
	PageNumber  *int      `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize    *int      `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListFlinkJobRunsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFlinkJobRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListFlinkJobRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListFlinkJobRunsResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxResults *int `json:"MaxResults" name:"MaxResults"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		JobRuns    []struct {
			JobRunId            *string `json:"JobRunId" name:"JobRunId"`
			WorkspaceId         *string `json:"WorkspaceId" name:"WorkspaceId"`
			Name                *string `json:"Name" name:"Name"`
			State               *string `json:"State" name:"State"`
			ErrorReason         *string `json:"ErrorReason" name:"ErrorReason"`
			SubmitTime          *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime           *string `json:"StartTime" name:"StartTime"`
			EndTime             *string `json:"EndTime" name:"EndTime"`
			RunningTime         *int    `json:"RunningTime" name:"RunningTime"`
			WebUI               *string `json:"WebUI" name:"WebUI"`
			FlinkLog            *string `json:"FlinkLog" name:"FlinkLog"`
			ReleaseVersion      *string `json:"ReleaseVersion" name:"ReleaseVersion"`
			FlinkDeploymentData struct {
				Name        *string   `json:"Name" name:"Name"`
				Image       *string   `json:"Image" name:"Image"`
				UpgradeMode *string   `json:"UpgradeMode" name:"UpgradeMode"`
				JobCores    *int      `json:"JobCores" name:"JobCores"`
				JobMemory   *string   `json:"JobMemory" name:"JobMemory"`
				TaskCores   *int      `json:"TaskCores" name:"TaskCores"`
				TaskMemory  *string   `json:"TaskMemory" name:"TaskMemory"`
				NumTasks    *int      `json:"NumTasks" name:"NumTasks"`
				EntryClass  *string   `json:"EntryClass" name:"EntryClass"`
				JarUri      *string   `json:"JarUri" name:"JarUri"`
				MainArgs    []*string `json:"MainArgs" name:"MainArgs"`
				FlinkConf   []*string `json:"FlinkConf" name:"FlinkConf"`
			} `json:"FlinkDeploymentData"`
			ResourceUsage *int `json:"ResourceUsage" name:"ResourceUsage"`
		} `json:"JobRuns" name:"JobRuns"`
	} `json:"Data"`
}

func (r *ListFlinkJobRunsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFlinkJobRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelFlinkJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string                       `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunIds   []*CancelFlinkJobRunJobRunIds `json:"JobRunIds,omitempty" name:"JobRunIds"`
}

func (r *CancelFlinkJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelFlinkJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelFlinkJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelFlinkJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *CancelFlinkJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelFlinkJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuspendFlinkJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunId    *string `json:"JobRunId,omitempty" name:"JobRunId"`
}

func (r *SuspendFlinkJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SuspendFlinkJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SuspendFlinkJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SuspendFlinkJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *SuspendFlinkJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SuspendFlinkJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartFlinkJobRunRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	JobRunId    *string `json:"JobRunId,omitempty" name:"JobRunId"`
}

func (r *RestartFlinkJobRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartFlinkJobRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestartFlinkJobRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartFlinkJobRunResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *RestartFlinkJobRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartFlinkJobRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricListRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *DescribeMetricListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMetricListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricListResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Count   *int `json:"Count" name:"Count"`
		Metrics []struct {
			MetricName  *string `json:"MetricName" name:"MetricName"`
			Unit        *string `json:"Unit" name:"Unit"`
			Period      *int    `json:"Period" name:"Period"`
			Statistics  *string `json:"Statistics" name:"Statistics"`
			Dimensions  *string `json:"Dimensions" name:"Dimensions"`
			Description *string `json:"Description" name:"Description"`
		} `json:"Metrics" name:"Metrics"`
	} `json:"Data"`
	IsPartial *bool   `json:"isPartial" name:"isPartial"`
	Status    *string `json:"status" name:"status"`
}

func (r *DescribeMetricListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMetricsRequest struct {
	*ksyunhttp.BaseRequest
	WorkspaceId *string                `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
	ProductType *string                `json:"ProductType,omitempty" name:"ProductType"`
	QueryData   *QueryMetricsQueryData `json:"QueryData,omitempty" name:"QueryData"`
}

func (r *QueryMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMetricsResponse struct {
	*ksyunhttp.BaseResponse
	Status    *string `json:"status" name:"status"`
	IsPartial *bool   `json:"isPartial" name:"isPartial"`
	Data      struct {
		ResultType *string `json:"ResultType" name:"ResultType"`
		Result     []struct {
			Metric struct {
				Name        *string `json:"Name" name:"Name"`
				JobRunId    *string `json:"JobRunId" name:"JobRunId"`
				Component   *string `json:"Component" name:"Component"`
				WorkspaceId *string `json:"WorkspaceId" name:"WorkspaceId"`
			} `json:"Metric"`
			Values []*string `json:"Values" name:"Values"`
		} `json:"Result" name:"Result"`
	} `json:"Data"`
	Stats struct {
		SeriesFetched     *string `json:"SeriesFetched" name:"SeriesFetched"`
		ExecutionTimeMsec *int    `json:"ExecutionTimeMsec" name:"ExecutionTimeMsec"`
	} `json:"Stats"`
}

func (r *QueryMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
