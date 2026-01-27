package v20251212

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateTrainJobStorageConfigs struct {
	StorageConfigId *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	MountType       *string `json:"MountType,omitempty" name:"MountType"`
	MountPath       *string `json:"MountPath,omitempty" name:"MountPath"`
}
type CreateTrainJobRolesImageConfig struct {
	ImageId         *string `json:"ImageId,omitempty" name:"ImageId"`
	ImageSource     *string `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageRegistryId *string `json:"ImageRegistryId,omitempty" name:"ImageRegistryId"`
	ImageRepoId     *string `json:"ImageRepoId,omitempty" name:"ImageRepoId"`
	ImageTagId      *string `json:"ImageTagId,omitempty" name:"ImageTagId"`
}
type CreateTrainJobRolesResourceConfig struct {
	GPUType   *string `json:"GPUType,omitempty" name:"GPUType"`
	GPUNumber *int    `json:"GPUNumber,omitempty" name:"GPUNumber"`
	CPUNum    *int    `json:"CPUNum,omitempty" name:"CPUNum"`
	Memory    *int    `json:"Memory,omitempty" name:"Memory"`
}
type CreateTrainJobRolesEnvs struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateTrainJobRoles struct {
	RoleName       *string                            `json:"RoleName,omitempty" name:"RoleName"`
	Replicas       *int                               `json:"Replicas,omitempty" name:"Replicas"`
	ImageConfig    *CreateTrainJobRolesImageConfig    `json:"ImageConfig,omitempty" name:"ImageConfig"`
	ResourceConfig *CreateTrainJobRolesResourceConfig `json:"ResourceConfig,omitempty" name:"ResourceConfig"`
	RunCommand     *string                            `json:"RunCommand,omitempty" name:"RunCommand"`
	RestartPolicy  *string                            `json:"RestartPolicy,omitempty" name:"RestartPolicy"`
	Envs           []*CreateTrainJobRolesEnvs         `json:"Envs,omitempty" name:"Envs"`
}
type DescribeTrainJobsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyModelAccessUsers struct {
	UserId     *string `json:"UserId,omitempty" name:"UserId"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}
type CreateModelAndVersionUsers struct {
	UserId     *string `json:"UserId,omitempty" name:"UserId"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

type CreateTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobName       *string                         `json:"TrainJobName,omitempty" name:"TrainJobName"`
	Description        *string                         `json:"Description,omitempty" name:"Description"`
	ResourcePoolId     *string                         `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	Priority           *string                         `json:"Priority,omitempty" name:"Priority"`
	QueueName          *string                         `json:"QueueName,omitempty" name:"QueueName"`
	Framework          *string                         `json:"Framework,omitempty" name:"Framework"`
	AccessType         *string                         `json:"AccessType,omitempty" name:"AccessType"`
	SelfHealing        *bool                           `json:"SelfHealing,omitempty" name:"SelfHealing"`
	MaxRuntimeHour     *int64                          `json:"MaxRuntimeHour,omitempty" name:"MaxRuntimeHour"`
	JobRunOnCPU        *bool                           `json:"JobRunOnCPU,omitempty" name:"JobRunOnCPU"`
	SupportTensorboard *bool                           `json:"SupportTensorboard,omitempty" name:"SupportTensorboard"`
	StorageConfigs     []*CreateTrainJobStorageConfigs `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
	Roles              []*CreateTrainJobRoles          `json:"Roles,omitempty" name:"Roles"`
}

func (r *CreateTrainJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTrainJobResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TrainJobId *string `json:"TrainJobId" name:"TrainJobId"`
}

func (r *CreateTrainJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrainJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrainJobsRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobId   []*string                  `json:"TrainJobId,omitempty" name:"TrainJobId"`
	Filter       []*DescribeTrainJobsFilter `json:"Filter,omitempty" name:"Filter"`
	PageSize     *int                       `json:"PageSize,omitempty" name:"PageSize"`
	Page         *int                       `json:"Page,omitempty" name:"Page"`
	TrainJobName *string                    `json:"TrainJobName,omitempty" name:"TrainJobName"`
}

func (r *DescribeTrainJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTrainJobsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	PageSize    *int    `json:"PageSize" name:"PageSize"`
	Page        *int    `json:"Page" name:"Page"`
	TotalCount  *int    `json:"TotalCount" name:"TotalCount"`
	TrainJobSet []struct {
		TrainJobName       *string `json:"TrainJobName" name:"TrainJobName"`
		Description        *string `json:"Description" name:"Description"`
		ResourcePoolId     *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		TrainJobId         *string `json:"TrainJobId" name:"TrainJobId"`
		Namespace          *string `json:"Namespace" name:"Namespace"`
		QueueName          *string `json:"QueueName" name:"QueueName"`
		Priority           *string `json:"Priority" name:"Priority"`
		Framework          *string `json:"Framework" name:"Framework"`
		AccessType         *string `json:"AccessType" name:"AccessType"`
		CreateUserId       *string `json:"CreateUserId" name:"CreateUserId"`
		SelfHealing        *bool   `json:"SelfHealing" name:"SelfHealing"`
		MaxRuntimeHour     *int64  `json:"MaxRuntimeHour" name:"MaxRuntimeHour"`
		JobRunOnCPU        *bool   `json:"JobRunOnCPU" name:"JobRunOnCPU"`
		SupportTensorboard *bool   `json:"SupportTensorboard" name:"SupportTensorboard"`
		RebootNumber       *int    `json:"RebootNumber" name:"RebootNumber"`
		TotalPodNum        *int    `json:"TotalPodNum" name:"TotalPodNum"`
		JobStatus          struct {
			Status        *string `json:"Status" name:"Status"`
			SubmitTime    *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime     *string `json:"StartTime" name:"StartTime"`
			QueueTime     *int    `json:"QueueTime" name:"QueueTime"`
			EndTime       *string `json:"EndTime" name:"EndTime"`
			Message       *string `json:"Message" name:"Message"`
			ExecutionTime *string `json:"ExecutionTime" name:"ExecutionTime"`
		} `json:"JobStatus" name:"JobStatus"`
		StorageConfigs []struct {
			StorageConfigId *string `json:"StorageConfigId" name:"StorageConfigId"`
			MountType       *string `json:"MountType" name:"MountType"`
			MountPath       *string `json:"MountPath" name:"MountPath"`
		} `json:"StorageConfigs" name:"StorageConfigs"`
		Roles []struct {
			RoleName    *string `json:"RoleName" name:"RoleName"`
			Replicas    *int    `json:"Replicas" name:"Replicas"`
			ImageConfig struct {
				ImageId           *string `json:"ImageId" name:"ImageId"`
				ImageSource       *string `json:"ImageSource" name:"ImageSource"`
				ImageRegistryId   *string `json:"ImageRegistryId" name:"ImageRegistryId"`
				ImageRepoId       *string `json:"ImageRepoId" name:"ImageRepoId"`
				ImageTagId        *string `json:"ImageTagId" name:"ImageTagId"`
				ImageRegistryName *string `json:"ImageRegistryName" name:"ImageRegistryName"`
				ImageRepoName     *string `json:"ImageRepoName" name:"ImageRepoName"`
				ImageTagName      *string `json:"ImageTagName" name:"ImageTagName"`
			} `json:"ImageConfig"`
			ResourceConfig struct {
				GPUType   *string `json:"GPUType" name:"GPUType"`
				GPUNumber *int    `json:"GPUNumber" name:"GPUNumber"`
				CPUNum    *int    `json:"CPUNum" name:"CPUNum"`
				Memory    *int    `json:"Memory" name:"Memory"`
			} `json:"ResourceConfig"`
			RunCommand    *string `json:"RunCommand" name:"RunCommand"`
			RestartPolicy *string `json:"RestartPolicy" name:"RestartPolicy"`
			Envs          []struct {
				Name  *string `json:"Name" name:"Name"`
				Value *string `json:"Value" name:"Value"`
			} `json:"Envs"`
		} `json:"Roles" name:"Roles"`
	} `json:"TrainJobSet"`
}

func (r *DescribeTrainJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrainJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyModelAccessRequest struct {
	*ksyunhttp.BaseRequest
	ModelId *string                   `json:"ModelId,omitempty" name:"ModelId"`
	Users   []*ModifyModelAccessUsers `json:"Users,omitempty" name:"Users"`
}

func (r *ModifyModelAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyModelAccessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ModelId   *string `json:"ModelId" name:"ModelId"`
}

func (r *ModifyModelAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModelAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateModelAndVersionRequest struct {
	*ksyunhttp.BaseRequest
	ModelName               *string                       `json:"ModelName,omitempty" name:"ModelName"`
	ModelDescription        *string                       `json:"ModelDescription,omitempty" name:"ModelDescription"`
	ModelVersionName        *string                       `json:"ModelVersionName,omitempty" name:"ModelVersionName"`
	ModelVersionDescription *string                       `json:"ModelVersionDescription,omitempty" name:"ModelVersionDescription"`
	SourceType              *string                       `json:"SourceType,omitempty" name:"SourceType"`
	StorageConfigId         *string                       `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	Format                  *string                       `json:"Format,omitempty" name:"Format"`
	Framework               *string                       `json:"Framework,omitempty" name:"Framework"`
	Users                   []*CreateModelAndVersionUsers `json:"Users,omitempty" name:"Users"`
}

func (r *CreateModelAndVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateModelAndVersionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	ModelId        *string `json:"ModelId" name:"ModelId"`
	ModelVersionId *string `json:"ModelVersionId" name:"ModelVersionId"`
}

func (r *CreateModelAndVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateModelAndVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyModelRequest struct {
	*ksyunhttp.BaseRequest
	ModelId          *string `json:"ModelId,omitempty" name:"ModelId"`
	ModelName        *string `json:"ModelName,omitempty" name:"ModelName"`
	ModelDescription *string `json:"ModelDescription,omitempty" name:"ModelDescription"`
}

func (r *ModifyModelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyModelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ModelId   *string `json:"ModelId" name:"ModelId"`
}

func (r *ModifyModelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelsRequest struct {
	*ksyunhttp.BaseRequest
	ModelIdN  []*string `json:"ModelId.N,omitempty" name:"ModelId.N"`
	ModelName *string   `json:"ModelName,omitempty" name:"ModelName"`
	Page      *int      `json:"Page,omitempty" name:"Page"`
	PageSize  *int      `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Page       *int    `json:"Page" name:"Page"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
	ModelSet   []struct {
		ModelId           *string `json:"ModelId" name:"ModelId"`
		ModelName         *string `json:"ModelName" name:"ModelName"`
		ModelDescription  *string `json:"ModelDescription" name:"ModelDescription"`
		ModelVersionCount *int    `json:"ModelVersionCount" name:"ModelVersionCount"`
		CreateUser        *string `json:"CreateUser" name:"CreateUser"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime        *string `json:"UpdateTime" name:"UpdateTime"`
		AccessList        []struct {
			UserId     *string `json:"UserId" name:"UserId"`
			Permission *string `json:"Permission" name:"Permission"`
		} `json:"AccessList" name:"AccessList"`
	} `json:"ModelSet"`
}

func (r *DescribeModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModelRequest struct {
	*ksyunhttp.BaseRequest
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *DeleteModelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteModelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ModelId   *string `json:"ModelId" name:"ModelId"`
}

func (r *DeleteModelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateModelVersionRequest struct {
	*ksyunhttp.BaseRequest
	ModelId                 *string `json:"ModelId,omitempty" name:"ModelId"`
	ModelVersionName        *string `json:"ModelVersionName,omitempty" name:"ModelVersionName"`
	ModelVersionDescription *string `json:"ModelVersionDescription,omitempty" name:"ModelVersionDescription"`
	SourceType              *string `json:"SourceType,omitempty" name:"SourceType"`
	StorageConfigId         *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	Format                  *string `json:"Format,omitempty" name:"Format"`
	Framework               *string `json:"Framework,omitempty" name:"Framework"`
}

func (r *CreateModelVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateModelVersionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	ModelVersionId *string `json:"ModelVersionId" name:"ModelVersionId"`
}

func (r *CreateModelVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModelVersionRequest struct {
	*ksyunhttp.BaseRequest
	ModelVersionId *string `json:"ModelVersionId,omitempty" name:"ModelVersionId"`
}

func (r *DeleteModelVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteModelVersionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	ModelVersionId *string `json:"ModelVersionId" name:"ModelVersionId"`
}

func (r *DeleteModelVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyModelVersionRequest struct {
	*ksyunhttp.BaseRequest
	ModelVersionId          *string `json:"ModelVersionId,omitempty" name:"ModelVersionId"`
	ModelVersionName        *string `json:"ModelVersionName,omitempty" name:"ModelVersionName"`
	ModelVersionDescription *string `json:"ModelVersionDescription,omitempty" name:"ModelVersionDescription"`
	Format                  *string `json:"Format,omitempty" name:"Format"`
	Framework               *string `json:"Framework,omitempty" name:"Framework"`
	SourceType              *string `json:"SourceType,omitempty" name:"SourceType"`
	StorageConfigId         *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
}

func (r *ModifyModelVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyModelVersionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	ModelVersionId *string `json:"ModelVersionId" name:"ModelVersionId"`
}

func (r *ModifyModelVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelVersionsRequest struct {
	*ksyunhttp.BaseRequest
	ModelVersionIdN  []*string `json:"ModelVersionId.N,omitempty" name:"ModelVersionId.N"`
	ModelVersionName *string   `json:"ModelVersionName,omitempty" name:"ModelVersionName"`
	ModelId          *string   `json:"ModelId,omitempty" name:"ModelId"`
	Page             *int      `json:"Page,omitempty" name:"Page"`
	PageSize         *int      `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeModelVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModelVersionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	TotalCount      *int    `json:"TotalCount" name:"TotalCount"`
	Page            *int    `json:"Page" name:"Page"`
	PageSize        *int    `json:"PageSize" name:"PageSize"`
	ModelVersionSet []struct {
		ModelId                 *string `json:"ModelId" name:"ModelId"`
		ModelVersionId          *string `json:"ModelVersionId" name:"ModelVersionId"`
		ModelVersionName        *string `json:"ModelVersionName" name:"ModelVersionName"`
		ModelVersionDescription *string `json:"ModelVersionDescription" name:"ModelVersionDescription"`
		SourceType              *string `json:"SourceType" name:"SourceType"`
		Format                  *string `json:"Format" name:"Format"`
		Framework               *string `json:"Framework" name:"Framework"`
		StorageConfigId         *string `json:"StorageConfigId" name:"StorageConfigId"`
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime              *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"ModelVersionSet"`
}

func (r *DescribeModelVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFormatAndFrameworksRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeFormatAndFrameworksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFormatAndFrameworksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId           *string `json:"RequestId" name:"RequestId"`
	FormatAndFrameworks []struct {
		Format     *string   `json:"Format" name:"Format"`
		Frameworks []*string `json:"Frameworks" name:"Frameworks"`
	} `json:"FormatAndFrameworks"`
}

func (r *DescribeFormatAndFrameworksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFormatAndFrameworksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
