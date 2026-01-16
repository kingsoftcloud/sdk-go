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
