package v20240612

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateStorageConfigKpfsInfo struct {
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	SystemFilePath *string `json:"SystemFilePath,omitempty" name:"SystemFilePath"`
	MntProtocol    *string `json:"MntProtocol,omitempty" name:"MntProtocol"`
}
type CreateStorageConfigKs3Info struct {
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	BucketPath *string `json:"BucketPath,omitempty" name:"BucketPath"`
}
type CreateStorageConfigUsers struct {
	UserId     *string `json:"UserId,omitempty" name:"UserId"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}
type ModifyStorageConfigKs3Info struct {
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	BucketPath *string `json:"BucketPath,omitempty" name:"BucketPath"`
}
type ModifyStorageConfigUsers struct {
	UserId     *string `json:"UserId,omitempty" name:"UserId"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}
type DescribeStorageConfigsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyNotebookStorageConfigs struct {
	StorageConfigId   *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	MountPath         *string `json:"MountPath,omitempty" name:"MountPath"`
	StorageConfigType *string `json:"StorageConfigType,omitempty" name:"StorageConfigType"`
}
type ModifyNotebookServiceConfigs struct {
	Service             *string `json:"Service,omitempty" name:"Service"`
	Port                *int    `json:"Port,omitempty" name:"Port"`
	EnablePublicNetwork *bool   `json:"EnablePublicNetwork,omitempty" name:"EnablePublicNetwork"`
}
type DescribeNotebooksFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateNotebookStorageConfigs struct {
	StorageConfigId   *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	MountPath         *string `json:"MountPath,omitempty" name:"MountPath"`
	StorageConfigType *string `json:"StorageConfigType,omitempty" name:"StorageConfigType"`
}
type CreateNotebookServiceConfigs struct {
	Service             *string `json:"Service,omitempty" name:"Service"`
	Port                *int    `json:"Port,omitempty" name:"Port"`
	EnablePublicNetwork *bool   `json:"EnablePublicNetwork,omitempty" name:"EnablePublicNetwork"`
}
type DescribeImagesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateInferenceCmdOptions struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateInferenceEnv struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateInferenceAutoScaleStrategyMetricStrategy struct {
	CpuUtilization *int `json:"CpuUtilization,omitempty" name:"CpuUtilization"`
	MemUtilization *int `json:"MemUtilization,omitempty" name:"MemUtilization"`
	MinReplicas    *int `json:"MinReplicas,omitempty" name:"MinReplicas"`
	MaxReplicas    *int `json:"MaxReplicas,omitempty" name:"MaxReplicas"`
}
type CreateInferenceAutoScaleStrategy struct {
	Mode           *string                                         `json:"Mode,omitempty" name:"Mode"`
	MetricStrategy *CreateInferenceAutoScaleStrategyMetricStrategy `json:"MetricStrategy,omitempty" name:"MetricStrategy"`
}
type CreateInferenceStorageConfigs struct {
	StorageConfigId   *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	StorageConfigType *string `json:"StorageConfigType,omitempty" name:"StorageConfigType"`
	MountPath         *string `json:"MountPath,omitempty" name:"MountPath"`
}
type DescribeTrainJobPodsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeInferencesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type SetInferenceAutoScaleStrategyAutoScaleStrategyMetricStrategy struct {
	CpuUtilization *int `json:"CpuUtilization,omitempty" name:"CpuUtilization"`
	MemUtilization *int `json:"MemUtilization,omitempty" name:"MemUtilization"`
	MinReplicas    *int `json:"MinReplicas,omitempty" name:"MinReplicas"`
	MaxReplicas    *int `json:"MaxReplicas,omitempty" name:"MaxReplicas"`
}
type SetInferenceAutoScaleStrategyAutoScaleStrategy struct {
	Mode           *string                                                       `json:"Mode,omitempty" name:"Mode"`
	MetricStrategy *SetInferenceAutoScaleStrategyAutoScaleStrategyMetricStrategy `json:"MetricStrategy,omitempty" name:"MetricStrategy"`
}
type ModifyInferenceEnv struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyInferenceCmdOptions struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type DescribeResourcePoolsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeResourcePoolInstancesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateInferenceEndpointRateLimit struct {
	RPM *int `json:"RPM,omitempty" name:"RPM"`
	TPM *int `json:"TPM,omitempty" name:"TPM"`
}
type DescribeInferenceEndpointsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type EnableEndpointRateLimitRateLimit struct {
	TPM *int `json:"TPM,omitempty" name:"TPM"`
	RPM *int `json:"RPM,omitempty" name:"RPM"`
}
type UpdateInferenceEndpointRateLimit struct {
	RPM         *int `json:"RPM,omitempty" name:"RPM"`
	TPM         *int `json:"TPM,omitempty" name:"TPM"`
	Concurrency *int `json:"Concurrency,omitempty" name:"Concurrency"`
	IPM         *int `json:"IPM,omitempty" name:"IPM"`
}
type DescribeQueuesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateStorageConfigRequest struct {
	*ksyunhttp.BaseRequest
	StorageConfigName *string                      `json:"StorageConfigName,omitempty" name:"StorageConfigName"`
	Description       *string                      `json:"Description,omitempty" name:"Description"`
	Type              *string                      `json:"Type,omitempty" name:"Type"`
	MountPath         *string                      `json:"MountPath,omitempty" name:"MountPath"`
	KpfsInfo          *CreateStorageConfigKpfsInfo `json:"KpfsInfo,omitempty" name:"KpfsInfo"`
	Ks3Info           *CreateStorageConfigKs3Info  `json:"Ks3Info,omitempty" name:"Ks3Info"`
	Users             []*CreateStorageConfigUsers  `json:"Users,omitempty" name:"Users"`
	Ak                *string                      `json:"Ak,omitempty" name:"Ak"`
	Sk                *string                      `json:"Sk,omitempty" name:"Sk"`
}

func (r *CreateStorageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateStorageConfigResponse struct {
	*ksyunhttp.BaseResponse
	StorageConfigId *string `json:"StorageConfigId" name:"StorageConfigId"`
	RequestId       *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateStorageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStorageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStorageConfigRequest struct {
	*ksyunhttp.BaseRequest
	StorageConfigId   *string                     `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	StorageConfigName *string                     `json:"StorageConfigName,omitempty" name:"StorageConfigName"`
	Description       *string                     `json:"Description,omitempty" name:"Description"`
	MountPath         *string                     `json:"MountPath,omitempty" name:"MountPath"`
	Ks3Info           *ModifyStorageConfigKs3Info `json:"Ks3Info,omitempty" name:"Ks3Info"`
	Users             []*ModifyStorageConfigUsers `json:"Users,omitempty" name:"Users"`
	Ak                *string                     `json:"Ak,omitempty" name:"Ak"`
	Sk                *string                     `json:"Sk,omitempty" name:"Sk"`
}

func (r *ModifyStorageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyStorageConfigResponse struct {
	*ksyunhttp.BaseResponse
	StorageConfigId *string `json:"StorageConfigId" name:"StorageConfigId"`
	RequestId       *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyStorageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStorageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageConfigsRequest struct {
	*ksyunhttp.BaseRequest
	StorageConfigId []*string                       `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	Filter          []*DescribeStorageConfigsFilter `json:"Filter,omitempty" name:"Filter"`
	PageSize        *int                            `json:"PageSize,omitempty" name:"PageSize"`
	Page            *int                            `json:"Page,omitempty" name:"Page"`
}

func (r *DescribeStorageConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeStorageConfigsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	PageSize         *int    `json:"PageSize" name:"PageSize"`
	Page             *int    `json:"Page" name:"Page"`
	TotalCount       *int    `json:"TotalCount" name:"TotalCount"`
	StorageConfigSet []struct {
		StorageConfigId   *string `json:"StorageConfigId" name:"StorageConfigId"`
		StorageConfigName *string `json:"StorageConfigName" name:"StorageConfigName"`
		Description       *string `json:"Description" name:"Description"`
		Type              *string `json:"Type" name:"Type"`
		MountPath         *string `json:"MountPath" name:"MountPath"`
		Ak                *string `json:"Ak" name:"Ak"`
		KpfsInfo          struct {
			FileSystemId   *string `json:"FileSystemId" name:"FileSystemId"`
			FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
			SystemFilePath *string `json:"SystemFilePath" name:"SystemFilePath"`
			MntProtocol    *string `json:"MntProtocol" name:"MntProtocol"`
			ClusterName    *string `json:"ClusterName" name:"ClusterName"`
		} `json:"KpfsInfo" name:"KpfsInfo"`
		Ks3Info struct {
			BucketName *string `json:"BucketName" name:"BucketName"`
			BucketPath *string `json:"BucketPath" name:"BucketPath"`
			Endpoint   *string `json:"Endpoint" name:"Endpoint"`
		} `json:"Ks3Info" name:"Ks3Info"`
		CreatorId  *string `json:"CreatorId" name:"CreatorId"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		Users      []struct {
			UserId     *string `json:"UserId" name:"UserId"`
			Permission *string `json:"Permission" name:"Permission"`
		} `json:"Users" name:"Users"`
	} `json:"StorageConfigSet"`
}

func (r *DescribeStorageConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStorageConfigRequest struct {
	*ksyunhttp.BaseRequest
	StorageConfigId *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
}

func (r *DeleteStorageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteStorageConfigResponse struct {
	*ksyunhttp.BaseResponse
	StorageConfigId *string `json:"StorageConfigId" name:"StorageConfigId"`
	RequestId       *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteStorageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStorageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveNotebookImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageName           *string `json:"ImageName,omitempty" name:"ImageName"`
	Description         *string `json:"Description,omitempty" name:"Description"`
	ImageType           *string `json:"ImageType,omitempty" name:"ImageType"`
	Namespace           *string `json:"Namespace,omitempty" name:"Namespace"`
	NamespacePermission *string `json:"NamespacePermission,omitempty" name:"NamespacePermission"`
	ImageRepo           *string `json:"ImageRepo,omitempty" name:"ImageRepo"`
	ImageVersion        *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
	OfficialInstance    *string `json:"OfficialInstance,omitempty" name:"OfficialInstance"`
	UserName            *string `json:"UserName,omitempty" name:"UserName"`
	Password            *string `json:"Password,omitempty" name:"Password"`
	ImagePermission     *string `json:"ImagePermission,omitempty" name:"ImagePermission"`
	NotebookId          *string `json:"NotebookId,omitempty" name:"NotebookId"`
	ImageDomain         *string `json:"ImageDomain,omitempty" name:"ImageDomain"`
}

func (r *SaveNotebookImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SaveNotebookImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageId   *string `json:"ImageId" name:"ImageId"`
}

func (r *SaveNotebookImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveNotebookImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNotebookRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId             *string                         `json:"NotebookId,omitempty" name:"NotebookId"`
	NotebookName           *string                         `json:"NotebookName,omitempty" name:"NotebookName"`
	Description            *string                         `json:"Description,omitempty" name:"Description"`
	ImageId                *string                         `json:"ImageId,omitempty" name:"ImageId"`
	QueueName              *string                         `json:"QueueName,omitempty" name:"QueueName"`
	GPUType                *string                         `json:"GPUType,omitempty" name:"GPUType"`
	GPUNumber              *string                         `json:"GPUNumber,omitempty" name:"GPUNumber"`
	CPUNum                 *int                            `json:"CPUNum,omitempty" name:"CPUNum"`
	Memory                 *int                            `json:"Memory,omitempty" name:"Memory"`
	AccessType             *string                         `json:"AccessType,omitempty" name:"AccessType"`
	EnablePublicNetworkSsh *bool                           `json:"EnablePublicNetworkSsh,omitempty" name:"EnablePublicNetworkSsh"`
	SshAuthorizedKeys      *string                         `json:"SshAuthorizedKeys,omitempty" name:"SshAuthorizedKeys"`
	StorageConfigs         []*ModifyNotebookStorageConfigs `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
	ServiceConfigs         []*ModifyNotebookServiceConfigs `json:"ServiceConfigs,omitempty" name:"ServiceConfigs"`
	AutoSave               *bool                           `json:"AutoSave,omitempty" name:"AutoSave"`
	RunOnCPU               *string                         `json:"RunOnCPU,omitempty" name:"RunOnCPU"`
	EnableSSH              *string                         `json:"EnableSSH,omitempty" name:"EnableSSH"`
	SSHPort                *int                            `json:"SSHPort,omitempty" name:"SSHPort"`
	SSHAuthorizedKeys      *string                         `json:"SSHAuthorizedKeys,omitempty" name:"SSHAuthorizedKeys"`
	EnablePublicNetworkSSH *bool                           `json:"EnablePublicNetworkSSH,omitempty" name:"EnablePublicNetworkSSH"`
	AllocationId           *string                         `json:"AllocationId,omitempty" name:"AllocationId"`
	ImageTagId             *string                         `json:"ImageTagId,omitempty" name:"ImageTagId"`
	ImageSource            *string                         `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageRepoId            *string                         `json:"ImageRepoId,omitempty" name:"ImageRepoId"`
	ImageRegistryId        *string                         `json:"ImageRegistryId,omitempty" name:"ImageRegistryId"`
}

func (r *ModifyNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyNotebookResponse struct {
	*ksyunhttp.BaseResponse
	NotebookId *string `json:"NotebookId" name:"NotebookId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNotebookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNotebookRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId *string `json:"NotebookId,omitempty" name:"NotebookId"`
}

func (r *DeleteNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteNotebookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NotebookId *string `json:"NotebookId" name:"NotebookId"`
}

func (r *DeleteNotebookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebooksRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId []*string                  `json:"NotebookId,omitempty" name:"NotebookId"`
	Name       *string                    `json:"Name,omitempty" name:"Name"`
	Page       *int                       `json:"Page,omitempty" name:"Page"`
	PageSize   *int                       `json:"PageSize,omitempty" name:"PageSize"`
	Filter     []*DescribeNotebooksFilter `json:"Filter,omitempty" name:"Filter"`
	QueueId    *string                    `json:"QueueId,omitempty" name:"QueueId"`
}

func (r *DescribeNotebooksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNotebooksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Notebooks []struct {
		NotebookId   *string `json:"NotebookId" name:"NotebookId"`
		NotebookName *string `json:"NotebookName" name:"NotebookName"`
		Description  *string `json:"Description" name:"Description"`
		Status       struct {
			State         *string `json:"State" name:"State"`
			SubmitTime    *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime     *string `json:"StartTime" name:"StartTime"`
			EndTime       *string `json:"EndTime" name:"EndTime"`
			Message       *string `json:"Message" name:"Message"`
			ExecutionTime *string `json:"ExecutionTime" name:"ExecutionTime"`
		} `json:"Status" name:"Status"`
		AutoSave               *bool   `json:"AutoSave" name:"AutoSave"`
		ImageSaveStatus        *string `json:"ImageSaveStatus" name:"ImageSaveStatus"`
		ImageId                *string `json:"ImageId" name:"ImageId"`
		GPUType                *string `json:"GPUType" name:"GPUType"`
		GPUNumber              *string `json:"GPUNumber" name:"GPUNumber"`
		CreateUser             *string `json:"CreateUser" name:"CreateUser"`
		Namespace              *string `json:"Namespace" name:"Namespace"`
		ResourcePoolId         *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		QueueName              *string `json:"QueueName" name:"QueueName"`
		AccessType             *string `json:"AccessType" name:"AccessType"`
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime             *string `json:"UpdateTime" name:"UpdateTime"`
		AllocationId           *string `json:"AllocationId" name:"AllocationId"`
		EnableSSH              *bool   `json:"EnableSSH" name:"EnableSSH"`
		EnablePublicNetworkSSH *bool   `json:"EnablePublicNetworkSSH" name:"EnablePublicNetworkSSH"`
		SSHPort                *int    `json:"SSHPort" name:"SSHPort"`
		SSHAuthorizedKeys      *string `json:"SSHAuthorizedKeys" name:"SSHAuthorizedKeys"`
		PodIp                  *string `json:"PodIp" name:"PodIp"`
		ExternalIp             *string `json:"ExternalIp" name:"ExternalIp"`
		CPUNum                 *int    `json:"CPUNum" name:"CPUNum"`
		Memory                 *int    `json:"Memory" name:"Memory"`
		ImageUrl               *string `json:"ImageUrl" name:"ImageUrl"`
		ImageSource            *string `json:"ImageSource" name:"ImageSource"`
		ImageRegistryId        *string `json:"ImageRegistryId" name:"ImageRegistryId"`
		ImageRepoId            *string `json:"ImageRepoId" name:"ImageRepoId"`
		ImageTagId             *string `json:"ImageTagId" name:"ImageTagId"`
		ImageRegistryName      *string `json:"ImageRegistryName" name:"ImageRegistryName"`
		ImageRepoName          *string `json:"ImageRepoName" name:"ImageRepoName"`
		ImageTagName           *string `json:"ImageTagName" name:"ImageTagName"`
		StorageConfigs         []struct {
			StorageConfigId   *string `json:"StorageConfigId" name:"StorageConfigId"`
			MountPath         *string `json:"MountPath" name:"MountPath"`
			StorageConfigType *string `json:"StorageConfigType" name:"StorageConfigType"`
		} `json:"StorageConfigs" name:"StorageConfigs"`
		ServiceConfigs []struct {
			Service             *string `json:"Service" name:"Service"`
			Port                *int    `json:"Port" name:"Port"`
			EnablePublicNetwork *bool   `json:"EnablePublicNetwork" name:"EnablePublicNetwork"`
		} `json:"ServiceConfigs" name:"ServiceConfigs"`
		Label struct {
			TerminatePolicyId *string `json:"TerminatePolicyId" name:"TerminatePolicyId"`
		} `json:"Label" name:"Label"`
		Version            *string  `json:"Version" name:"Version"`
		EstimatedImageSize *float64 `json:"EstimatedImageSize" name:"EstimatedImageSize"`
		ComputeStatus      *string  `json:"ComputeStatus" name:"ComputeStatus"`
		RunOnCPU           *bool    `json:"RunOnCPU" name:"RunOnCPU"`
		NodeIp             *string  `json:"NodeIp" name:"NodeIp"`
	} `json:"Notebooks"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Page       *int `json:"Page" name:"Page"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
}

func (r *DescribeNotebooksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNotebooksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNotebookRequest struct {
	*ksyunhttp.BaseRequest
	NotebookName           *string                         `json:"NotebookName,omitempty" name:"NotebookName"`
	Description            *string                         `json:"Description,omitempty" name:"Description"`
	ResourcePoolId         *string                         `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	QueueName              *string                         `json:"QueueName,omitempty" name:"QueueName"`
	GPUType                *string                         `json:"GPUType,omitempty" name:"GPUType"`
	GPUNumber              *string                         `json:"GPUNumber,omitempty" name:"GPUNumber"`
	CPUNum                 *int                            `json:"CPUNum,omitempty" name:"CPUNum"`
	Memory                 *int                            `json:"Memory,omitempty" name:"Memory"`
	AccessType             *string                         `json:"AccessType,omitempty" name:"AccessType"`
	StorageConfigs         []*CreateNotebookStorageConfigs `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
	AutoSave               *bool                           `json:"AutoSave,omitempty" name:"AutoSave"`
	ServiceConfigs         []*CreateNotebookServiceConfigs `json:"ServiceConfigs,omitempty" name:"ServiceConfigs"`
	ImageSource            *string                         `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageId                *string                         `json:"ImageId,omitempty" name:"ImageId"`
	ImageRegistryId        *string                         `json:"ImageRegistryId,omitempty" name:"ImageRegistryId"`
	ImageRepoId            *string                         `json:"ImageRepoId,omitempty" name:"ImageRepoId"`
	ImageTagId             *string                         `json:"ImageTagId,omitempty" name:"ImageTagId"`
	EnableSSH              *bool                           `json:"EnableSSH,omitempty" name:"EnableSSH"`
	SSHAuthorizedKeys      *string                         `json:"SSHAuthorizedKeys,omitempty" name:"SSHAuthorizedKeys"`
	SSHPort                *int                            `json:"SSHPort,omitempty" name:"SSHPort"`
	EnablePublicNetworkSSH *bool                           `json:"EnablePublicNetworkSSH,omitempty" name:"EnablePublicNetworkSSH"`
	AllocationId           *string                         `json:"AllocationId,omitempty" name:"AllocationId"`
	RunOnCPU               *string                         `json:"RunOnCPU,omitempty" name:"RunOnCPU"`
}

func (r *CreateNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateNotebookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NotebookId *string `json:"NotebookId" name:"NotebookId"`
}

func (r *CreateNotebookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageName           *string `json:"ImageName,omitempty" name:"ImageName"`
	Description         *string `json:"Description,omitempty" name:"Description"`
	ImageType           *string `json:"ImageType,omitempty" name:"ImageType"`
	Namespace           *string `json:"Namespace,omitempty" name:"Namespace"`
	NamespacePermission *string `json:"NamespacePermission,omitempty" name:"NamespacePermission"`
	ImageRepo           *string `json:"ImageRepo,omitempty" name:"ImageRepo"`
	ImageVersion        *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
	OfficialInstance    *string `json:"OfficialInstance,omitempty" name:"OfficialInstance"`
	UserName            *string `json:"UserName,omitempty" name:"UserName"`
	Password            *string `json:"Password,omitempty" name:"Password"`
	ImagePermission     *string `json:"ImagePermission,omitempty" name:"ImagePermission"`
}

func (r *CreateImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageId   *string `json:"ImageId" name:"ImageId"`
}

func (r *CreateImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageId   *string `json:"ImageId" name:"ImageId"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId         *string `json:"ImageId,omitempty" name:"ImageId"`
	ImageName       *string `json:"ImageName,omitempty" name:"ImageName"`
	ImagePermission *string `json:"ImagePermission,omitempty" name:"ImagePermission"`
}

func (r *ModifyImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageId   *string `json:"ImageId" name:"ImageId"`
}

func (r *ModifyImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*ksyunhttp.BaseRequest
	Page                *int                    `json:"Page,omitempty" name:"Page"`
	PageSize            *int                    `json:"PageSize,omitempty" name:"PageSize"`
	ImageSource         *string                 `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageStatus         *string                 `json:"ImageStatus,omitempty" name:"ImageStatus"`
	ImageType           *string                 `json:"ImageType,omitempty" name:"ImageType"`
	ApplicationScenario *string                 `json:"ApplicationScenario,omitempty" name:"ApplicationScenario"`
	ImageId             []*string               `json:"ImageId,omitempty" name:"ImageId"`
	ImageName           *string                 `json:"ImageName,omitempty" name:"ImageName"`
	Filter              []*DescribeImagesFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeImagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageSet  []struct {
		ImageId             *string   `json:"ImageId" name:"ImageId"`
		ImageName           *string   `json:"ImageName" name:"ImageName"`
		ImageSource         *string   `json:"ImageSource" name:"ImageSource"`
		ImageType           *string   `json:"ImageType" name:"ImageType"`
		ImageFrame          []*string `json:"ImageFrame" name:"ImageFrame"`
		PythonVersion       *string   `json:"PythonVersion" name:"PythonVersion"`
		ApplicationScenario *string   `json:"ApplicationScenario" name:"ApplicationScenario"`
		CudaVersion         *string   `json:"CudaVersion" name:"CudaVersion"`
		ImageSize           *string   `json:"ImageSize" name:"ImageSize"`
		Description         *string   `json:"Description" name:"Description"`
		Namespace           *string   `json:"Namespace" name:"Namespace"`
		NamespacePermission *string   `json:"NamespacePermission" name:"NamespacePermission"`
		ImageRepo           *string   `json:"ImageRepo" name:"ImageRepo"`
		ImageVersion        *string   `json:"ImageVersion" name:"ImageVersion"`
		ImageDomain         *string   `json:"ImageDomain" name:"ImageDomain"`
		OfficialInstance    *string   `json:"OfficialInstance" name:"OfficialInstance"`
		ImagePermission     *string   `json:"ImagePermission" name:"ImagePermission"`
		ImageStatus         *string   `json:"ImageStatus" name:"ImageStatus"`
		ImageAbnormalReason *string   `json:"ImageAbnormalReason" name:"ImageAbnormalReason"`
		ImageStatusName     *string   `json:"ImageStatusName" name:"ImageStatusName"`
		CreateUser          *string   `json:"CreateUser" name:"CreateUser"`
		CreateTime          *string   `json:"CreateTime" name:"CreateTime"`
		UpdateTime          *string   `json:"UpdateTime" name:"UpdateTime"`
	} `json:"ImageSet"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Page       *int `json:"Page" name:"Page"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInferenceRequest struct {
	*ksyunhttp.BaseRequest
	InferenceName       *string                           `json:"InferenceName,omitempty" name:"InferenceName"`
	Description         *string                           `json:"Description,omitempty" name:"Description"`
	ResourcePoolId      *string                           `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	QueueName           *string                           `json:"QueueName,omitempty" name:"QueueName"`
	Replicas            *int                              `json:"Replicas,omitempty" name:"Replicas"`
	AccessType          *string                           `json:"AccessType,omitempty" name:"AccessType"`
	DeploymentType      *string                           `json:"DeploymentType,omitempty" name:"DeploymentType"`
	Engine              *string                           `json:"Engine,omitempty" name:"Engine"`
	ModelName           *string                           `json:"ModelName,omitempty" name:"ModelName"`
	CmdOptions          []*CreateInferenceCmdOptions      `json:"CmdOptions,omitempty" name:"CmdOptions"`
	ModelStorageEnabled *bool                             `json:"ModelStorageEnabled,omitempty" name:"ModelStorageEnabled"`
	ModelStoragePath    *string                           `json:"ModelStoragePath,omitempty" name:"ModelStoragePath"`
	EntryPoint          *string                           `json:"EntryPoint,omitempty" name:"EntryPoint"`
	ImageSource         *string                           `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageId             *string                           `json:"ImageId,omitempty" name:"ImageId"`
	ImageRegistryId     *string                           `json:"ImageRegistryId,omitempty" name:"ImageRegistryId"`
	ImageRepoId         *string                           `json:"ImageRepoId,omitempty" name:"ImageRepoId"`
	ImageTagId          *string                           `json:"ImageTagId,omitempty" name:"ImageTagId"`
	SubnetId            *string                           `json:"SubnetId,omitempty" name:"SubnetId"`
	Port                *int                              `json:"Port,omitempty" name:"Port"`
	Env                 []*CreateInferenceEnv             `json:"Env,omitempty" name:"Env"`
	GPUType             *string                           `json:"GPUType,omitempty" name:"GPUType"`
	GPUNum              *string                           `json:"GPUNum,omitempty" name:"GPUNum"`
	CPUNum              *int                              `json:"CPUNum,omitempty" name:"CPUNum"`
	Memory              *int                              `json:"Memory,omitempty" name:"Memory"`
	AutoScaleEnable     *bool                             `json:"AutoScaleEnable,omitempty" name:"AutoScaleEnable"`
	AutoScaleStrategy   *CreateInferenceAutoScaleStrategy `json:"AutoScaleStrategy,omitempty" name:"AutoScaleStrategy"`
	RunOnCPU            *bool                             `json:"RunOnCPU,omitempty" name:"RunOnCPU"`
	Distributed         *bool                             `json:"Distributed,omitempty" name:"Distributed"`
	NodeNum             *bool                             `json:"NodeNum,omitempty" name:"NodeNum"`
	StorageConfigs      []*CreateInferenceStorageConfigs  `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
}

func (r *CreateInferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInferenceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *CreateInferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInferenceModelsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetInferenceModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetInferenceModelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	TotalCount      *int    `json:"TotalCount" name:"TotalCount"`
	InferenceModels []struct {
		Name        *string `json:"Name" name:"Name"`
		Group       *string `json:"Group" name:"Group"`
		StorageType *string `json:"StorageType" name:"StorageType"`
		Description *string `json:"Description" name:"Description"`
		PassParams  struct {
			GPUNum            *string `json:"GPUNum" name:"GPUNum"`
			CPUCore           *int    `json:"CPUCore" name:"CPUCore"`
			Memory            *int    `json:"Memory" name:"Memory"`
			SupportGPUNumList *string `json:"SupportGPUNumList" name:"SupportGPUNumList"`
			SupportSingleNode *bool   `json:"SupportSingleNode" name:"SupportSingleNode"`
			SupportMultiNode  *bool   `json:"SupportMultiNode" name:"SupportMultiNode"`
			SupportVLLM       *bool   `json:"SupportVLLM" name:"SupportVLLM"`
			SupportSGLang     *bool   `json:"SupportSGLang" name:"SupportSGLang"`
			SupportSGLangPD   *bool   `json:"SupportSGLangPD" name:"SupportSGLangPD"`
			SupportOpenPi     *bool   `json:"SupportOpenPi" name:"SupportOpenPi"`
		} `json:"PassParams" name:"PassParams"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"InferenceModels"`
}

func (r *GetInferenceModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInferenceModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInferencePodsRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
	State       *string `json:"State,omitempty" name:"State"`
}

func (r *GetInferencePodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetInferencePodsResponse struct {
	*ksyunhttp.BaseResponse
	Pods []struct {
		PodName     *string `json:"PodName" name:"PodName"`
		InferenceId *string `json:"InferenceId" name:"InferenceId"`
		Namespace   *string `json:"Namespace" name:"Namespace"`
		ClusterId   *string `json:"ClusterId" name:"ClusterId"`
		State       *string `json:"State" name:"State"`
	} `json:"Pods"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetInferencePodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInferencePodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInferenceLogsRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId  *string `json:"InferenceId,omitempty" name:"InferenceId"`
	PodName      *string `json:"PodName,omitempty" name:"PodName"`
	SinceSeconds *int    `json:"SinceSeconds,omitempty" name:"SinceSeconds"`
	TailLines    *int    `json:"TailLines,omitempty" name:"TailLines"`
}

func (r *GetInferenceLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetInferenceLogsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	PodLogs   *string `json:"PodLogs" name:"PodLogs"`
}

func (r *GetInferenceLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInferenceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopNotebookRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId *string `json:"NotebookId,omitempty" name:"NotebookId"`
}

func (r *StopNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopNotebookResponse struct {
	*ksyunhttp.BaseResponse
	NotebookId *string `json:"NotebookId" name:"NotebookId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *StopNotebookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartNotebookRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId *string `json:"NotebookId,omitempty" name:"NotebookId"`
}

func (r *StartNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartNotebookResponse struct {
	*ksyunhttp.BaseResponse
	NotebookId *string `json:"NotebookId" name:"NotebookId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *StartNotebookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWebIdeUrlRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId       *string `json:"NotebookId,omitempty" name:"NotebookId"`
	ExpirationMinute *string `json:"ExpirationMinute,omitempty" name:"ExpirationMinute"`
}

func (r *GetWebIdeUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetWebIdeUrlResponse struct {
	*ksyunhttp.BaseResponse
	JupyterWebUrl *string `json:"JupyterWebUrl" name:"JupyterWebUrl"`
	VscodeWebUrl  *string `json:"VscodeWebUrl" name:"VscodeWebUrl"`
	NotebookId    *string `json:"NotebookId" name:"NotebookId"`
	RequestId     *string `json:"RequestId" name:"RequestId"`
}

func (r *GetWebIdeUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWebIdeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookLogRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId   *string `json:"NotebookId,omitempty" name:"NotebookId"`
	SinceSeconds *int    `json:"SinceSeconds,omitempty" name:"SinceSeconds"`
	TailLines    *string `json:"TailLines,omitempty" name:"TailLines"`
}

func (r *DescribeNotebookLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNotebookLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	PodLogs   *string `json:"PodLogs" name:"PodLogs"`
}

func (r *DescribeNotebookLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNotebookLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInferenceAutoScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
}

func (r *GetInferenceAutoScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetInferenceAutoScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	AutoScaleStrategy struct {
		Mode           *string `json:"Mode" name:"Mode"`
		MetricStrategy struct {
			CpuUtilization *int `json:"CpuUtilization" name:"CpuUtilization"`
			MemUtilization *int `json:"MemUtilization" name:"MemUtilization"`
			MinReplicas    *int `json:"MinReplicas" name:"MinReplicas"`
			MaxReplicas    *int `json:"MaxReplicas" name:"MaxReplicas"`
		} `json:"MetricStrategy" name:"MetricStrategy"`
	} `json:"AutoScaleStrategy"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetInferenceAutoScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInferenceAutoScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopNotebookSavingImageRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId *string `json:"NotebookId,omitempty" name:"NotebookId"`
}

func (r *StopNotebookSavingImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopNotebookSavingImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NotebookId *string `json:"NotebookId" name:"NotebookId"`
}

func (r *StopNotebookSavingImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopNotebookSavingImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableApikeyStatusRequest struct {
	*ksyunhttp.BaseRequest
	KeyId  *string `json:"KeyId,omitempty" name:"KeyId"`
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *EnableApikeyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EnableApikeyStatusResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EnableApikeyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableApikeyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApikeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyId              *string   `json:"KeyId,omitempty" name:"KeyId"`
	Name               *string   `json:"Name,omitempty" name:"Name"`
	Description        *string   `json:"Description,omitempty" name:"Description"`
	AssociatedModelIds []*string `json:"AssociatedModelIds,omitempty" name:"AssociatedModelIds"`
	AllAssociatedModel *bool     `json:"AllAssociatedModel,omitempty" name:"AllAssociatedModel"`
}

func (r *ModifyApikeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyApikeyResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyApikeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApikeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateApiServiceRequest struct {
	*ksyunhttp.BaseRequest
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ActivateApiServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateApiServiceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"Success" name:"Success"`
	Error     struct {
		Code    *string `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Error"`
}

func (r *ActivateApiServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateApiServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApikeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteApikeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteApikeyResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteApikeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApikeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelsRequest struct {
	*ksyunhttp.BaseRequest
	Marker        *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults    *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	ModelCategory []*string `json:"ModelCategory,omitempty" name:"ModelCategory"`
	Provider      []*string `json:"Provider,omitempty" name:"Provider"`
	ContextLength []*int    `json:"ContextLength,omitempty" name:"ContextLength"`
	ModelName     *string   `json:"ModelName,omitempty" name:"ModelName"`
}

func (r *DescribeModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModelsResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
	Data       []struct {
		ModelId     *string   `json:"ModelId" name:"ModelId"`
		ModelName   *string   `json:"ModelName" name:"ModelName"`
		Description *string   `json:"Description" name:"Description"`
		IconUrl     *string   `json:"IconUrl" name:"IconUrl"`
		Tags        []*string `json:"Tags" name:"Tags"`
		UpdateTime  *int64    `json:"UpdateTime" name:"UpdateTime"`
		CodeCase    struct {
			VideoModel    *bool     `json:"VideoModel" name:"VideoModel"`
			Batch         *bool     `json:"Batch" name:"Batch"`
			ModelKind     *int      `json:"ModelKind" name:"ModelKind"`
			InputType     []*string `json:"InputType" name:"InputType"`
			OutputType    []*string `json:"OutputType" name:"OutputType"`
			SupportTryout *string   `json:"SupportTryout" name:"SupportTryout"`
			ContextLength *string   `json:"ContextLength" name:"ContextLength"`
		} `json:"CodeCase" name:"CodeCase"`
		IsOverFreeDisable *bool   `json:"IsOverFreeDisable" name:"IsOverFreeDisable"`
		FreeTotalQuota    *int64  `json:"FreeTotalQuota" name:"FreeTotalQuota"`
		FreeUsedQuota     *int64  `json:"FreeUsedQuota" name:"FreeUsedQuota"`
		Tpm               *int    `json:"Tpm" name:"Tpm"`
		Rpm               *int    `json:"Rpm" name:"Rpm"`
		ActiveStatus      *int    `json:"ActiveStatus" name:"ActiveStatus"`
		ModelType         *string `json:"ModelType" name:"ModelType"`
		Provider          *string `json:"Provider" name:"Provider"`
	} `json:"Data"`
}

func (r *DescribeModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApikeyRequest struct {
	*ksyunhttp.BaseRequest
	Name               *string   `json:"Name,omitempty" name:"Name"`
	Description        *string   `json:"Description,omitempty" name:"Description"`
	ProjectId          *int64    `json:"ProjectId,omitempty" name:"ProjectId"`
	AssociatedModelIds []*string `json:"AssociatedModelIds,omitempty" name:"AssociatedModelIds"`
	AllAssociatedModel *bool     `json:"AllAssociatedModel,omitempty" name:"AllAssociatedModel"`
	AllowedIps         []*string `json:"AllowedIps,omitempty" name:"AllowedIps"`
}

func (r *CreateApikeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateApikeyResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateApikeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApikeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModelDetailRequest struct {
	*ksyunhttp.BaseRequest
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *GetModelDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetModelDetailResponse struct {
	*ksyunhttp.BaseResponse
	ModelApiModelDataWebResp struct {
		ModelId     *string   `json:"ModelId" name:"ModelId"`
		ModelName   *string   `json:"ModelName" name:"ModelName"`
		Description *string   `json:"Description" name:"Description"`
		IconUrl     *string   `json:"IconUrl" name:"IconUrl"`
		Tags        []*string `json:"Tags" name:"Tags"`
		UpdateTime  *int64    `json:"UpdateTime" name:"UpdateTime"`
		CodeCase    struct {
			BaseUrl                  *string `json:"BaseUrl" name:"BaseUrl"`
			Model                    *string `json:"Model" name:"Model"`
			RestApiExternalbaseUrl   *string `json:"RestApiExternalbaseUrl" name:"RestApiExternalbaseUrl"`
			RestApiInternalbaseUrl   *string `json:"RestApiInternalbaseUrl" name:"RestApiInternalbaseUrl"`
			OpenSdkExternalbaseUrl   *string `json:"OpenSdkExternalbaseUrl" name:"OpenSdkExternalbaseUrl"`
			OpenSdkInternalbaseUrl   *string `json:"OpenSdkInternalbaseUrl" name:"OpenSdkInternalbaseUrl"`
			VideoModel               *bool   `json:"VideoModel" name:"VideoModel"`
			FunctionCall             *bool   `json:"FunctionCall" name:"FunctionCall"`
			StructuredOutput         *bool   `json:"StructuredOutput" name:"StructuredOutput"`
			InternetSearch           *bool   `json:"InternetSearch" name:"InternetSearch"`
			DeepThinking             *bool   `json:"DeepThinking" name:"DeepThinking"`
			Batch                    *bool   `json:"Batch" name:"Batch"`
			CacheContext             *bool   `json:"CacheContext" name:"CacheContext"`
			ModelProvider            *string `json:"ModelProvider" name:"ModelProvider"`
			OnlineCalculateInput     *string `json:"OnlineCalculateInput" name:"OnlineCalculateInput"`
			OnlineCalculateOutput    *string `json:"OnlineCalculateOutput" name:"OnlineCalculateOutput"`
			BatchCalculateInput      *string `json:"BatchCalculateInput" name:"BatchCalculateInput"`
			BatchCalculateOutput     *string `json:"BatchCalculateOutput" name:"BatchCalculateOutput"`
			OnlineCalculateRight     *string `json:"OnlineCalculateRight" name:"OnlineCalculateRight"`
			BatchCalculateCacheRight *string `json:"BatchCalculateCacheRight" name:"BatchCalculateCacheRight"`
			RpmLimit                 *string `json:"RpmLimit" name:"RpmLimit"`
			TpmLimit                 *string `json:"TpmLimit" name:"TpmLimit"`
			ContextLength            *string `json:"ContextLength" name:"ContextLength"`
			InputMaxLength           *string `json:"InputMaxLength" name:"InputMaxLength"`
			OutputMaxLength          *string `json:"OutputMaxLength" name:"OutputMaxLength"`
			DeepThinkMaxLength       *string `json:"DeepThinkMaxLength" name:"DeepThinkMaxLength"`
			VideoGeneration          *bool   `json:"VideoGeneration" name:"VideoGeneration"`
			ImageToVideo             *bool   `json:"ImageToVideo" name:"ImageToVideo"`
			VideoModelModeSet        []struct {
				VideoModelMode  *string `json:"VideoModelMode" name:"VideoModelMode"`
				VideoModelPrice *string `json:"VideoModelPrice" name:"VideoModelPrice"`
				Resolution      *string `json:"Resolution" name:"Resolution"`
				FrameRate       *string `json:"FrameRate" name:"FrameRate"`
				Duration        *string `json:"Duration" name:"Duration"`
				Concurrency     *string `json:"Concurrency" name:"Concurrency"`
			} `json:"VideoModelModeSet"`
			SingleExternalUrl  *string   `json:"SingleExternalUrl" name:"SingleExternalUrl"`
			SingleInternalUrl  *string   `json:"SingleInternalUrl" name:"SingleInternalUrl"`
			DeepThinkingSwitch *bool     `json:"DeepThinkingSwitch" name:"DeepThinkingSwitch"`
			ImageModelModeMap  *string   `json:"ImageModelModeMap" name:"ImageModelModeMap"`
			InputType          []*string `json:"InputType" name:"InputType"`
			OutputType         []*string `json:"OutputType" name:"OutputType"`
			SupportTryout      *string   `json:"SupportTryout" name:"SupportTryout"`
			TextToSpeech       *bool     `json:"TextToSpeech" name:"TextToSpeech"`
			VoicePrice         *string   `json:"VoicePrice" name:"VoicePrice"`
			PricingRules       struct {
				OnlineCalculateDisplayCache []struct {
				} `json:"OnlineCalculateDisplayCache" name:"OnlineCalculateDisplayCache"`
			} `json:"PricingRules"`
		} `json:"CodeCase" name:"CodeCase"`
		RequestId    *string `json:"RequestId" name:"RequestId"`
		ActiveStatus *int    `json:"ActiveStatus" name:"ActiveStatus"`
		ExtraFields  *string `json:"ExtraFields" name:"ExtraFields"`
	} `json:"ModelApiModelDataWebResp"`
}

func (r *GetModelDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModelDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApikeysRequest struct {
	*ksyunhttp.BaseRequest
	Marker            *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults        *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	AssociatedModelId []*string `json:"AssociatedModelId,omitempty" name:"AssociatedModelId"`
	Status            []*string `json:"Status,omitempty" name:"Status"`
	Namekeyword       *string   `json:"Namekeyword,omitempty" name:"Namekeyword"`
	DefaultKey        *bool     `json:"DefaultKey,omitempty" name:"DefaultKey"`
	KeyId             []*string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeApikeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeApikeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		KeyId              *string   `json:"KeyId" name:"KeyId"`
		KeyValue           *string   `json:"KeyValue" name:"KeyValue"`
		AssociatedModelIds []*string `json:"AssociatedModelIds" name:"AssociatedModelIds"`
		Name               *string   `json:"Name" name:"Name"`
		Description        *string   `json:"Description" name:"Description"`
		Status             *string   `json:"Status" name:"Status"`
		CreateTimestamp    *int64    `json:"CreateTimestamp" name:"CreateTimestamp"`
		ProjectId          *string   `json:"ProjectId" name:"ProjectId"`
		ProjectName        *string   `json:"ProjectName" name:"ProjectName"`
		CreateUserId       *string   `json:"CreateUserId" name:"CreateUserId"`
		CreateUserName     *string   `json:"CreateUserName" name:"CreateUserName"`
		AssociatedModels   []struct {
			ModelId   *string `json:"ModelId" name:"ModelId"`
			ModelName *string `json:"ModelName" name:"ModelName"`
		} `json:"AssociatedModels" name:"AssociatedModels"`
		AllAssociatedModel *bool     `json:"AllAssociatedModel" name:"AllAssociatedModel"`
		AllowedIps         []*string `json:"AllowedIps" name:"AllowedIps"`
	} `json:"Data"`
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeApikeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApikeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableApikeyStatusRequest struct {
	*ksyunhttp.BaseRequest
	KeyId  *string `json:"KeyId,omitempty" name:"KeyId"`
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DisableApikeyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisableApikeyStatusResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DisableApikeyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableApikeyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiServiceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetApiServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetApiServiceResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"Status" name:"Status"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetApiServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBatchInferenceJobsFinalResultDownloadUrlRequest struct {
	*ksyunhttp.BaseRequest
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *GetBatchInferenceJobsFinalResultDownloadUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBatchInferenceJobsFinalResultDownloadUrlResponse struct {
	*ksyunhttp.BaseResponse
	DownloadUrl *string `json:"DownloadUrl" name:"DownloadUrl"`
}

func (r *GetBatchInferenceJobsFinalResultDownloadUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBatchInferenceJobsFinalResultDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInferenceJobsKs3AuthInfoRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeInferenceJobsKs3AuthInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInferenceJobsKs3AuthInfoResponse struct {
	*ksyunhttp.BaseResponse
	Id             *string `json:"Id" name:"Id"`
	Ak             *string `json:"Ak" name:"Ak"`
	Sk             *string `json:"Sk" name:"Sk"`
	NeedCreateRole *bool   `json:"NeedCreateRole" name:"NeedCreateRole"`
}

func (r *DescribeInferenceJobsKs3AuthInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInferenceJobsKs3AuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopBatchInferenceJobRequest struct {
	*ksyunhttp.BaseRequest
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *StopBatchInferenceJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopBatchInferenceJobResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StopBatchInferenceJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBatchInferenceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBatchInferenceJobRequest struct {
	*ksyunhttp.BaseRequest
	JobName          *string `json:"JobName,omitempty" name:"JobName"`
	JobDesc          *string `json:"JobDesc,omitempty" name:"JobDesc"`
	ApikeyId         *string `json:"ApikeyId,omitempty" name:"ApikeyId"`
	Model            *string `json:"Model,omitempty" name:"Model"`
	ExecuteTimeoutMs *int64  `json:"ExecuteTimeoutMs,omitempty" name:"ExecuteTimeoutMs"`
	InputDataType    *string `json:"InputDataType,omitempty" name:"InputDataType"`
	Ks3Region        *string `json:"Ks3Region,omitempty" name:"Ks3Region"`
	Ks3Ak            *string `json:"Ks3Ak,omitempty" name:"Ks3Ak"`
	Ks3Sk            *string `json:"Ks3Sk,omitempty" name:"Ks3Sk"`
	InBucket         *string `json:"InBucket,omitempty" name:"InBucket"`
	OutBucket        *string `json:"OutBucket,omitempty" name:"OutBucket"`
	InObjectName     *string `json:"InObjectName,omitempty" name:"InObjectName"`
	OutObjectName    *string `json:"OutObjectName,omitempty" name:"OutObjectName"`
}

func (r *CreateBatchInferenceJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateBatchInferenceJobResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateBatchInferenceJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBatchInferenceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBatchInferenceJobRequest struct {
	*ksyunhttp.BaseRequest
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	JobName *string `json:"JobName,omitempty" name:"JobName"`
	JobDesc *string `json:"JobDesc,omitempty" name:"JobDesc"`
}

func (r *ModifyBatchInferenceJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyBatchInferenceJobResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyBatchInferenceJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBatchInferenceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchInferenceJobsRequest struct {
	*ksyunhttp.BaseRequest
	Marker         *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults     *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	JobNameKeyword *string   `json:"JobNameKeyword,omitempty" name:"JobNameKeyword"`
	Status         []*string `json:"Status,omitempty" name:"Status"`
	BatchId        *string   `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *DescribeBatchInferenceJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBatchInferenceJobsResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
	Data       []struct {
		BatchId        *string `json:"BatchId" name:"BatchId"`
		JobName        *string `json:"JobName" name:"JobName"`
		JobDesc        *string `json:"JobDesc" name:"JobDesc"`
		Status         *string `json:"Status" name:"Status"`
		CreateTime     *int64  `json:"CreateTime" name:"CreateTime"`
		CreateUserName *string `json:"CreateUserName" name:"CreateUserName"`
		JobMetadata    struct {
			Total            *int    `json:"Total" name:"Total"`
			Executed         *int    `json:"Executed" name:"Executed"`
			ApikeyId         *string `json:"ApikeyId" name:"ApikeyId"`
			Model            *string `json:"Model" name:"Model"`
			ExecuteTimeoutMs *int64  `json:"ExecuteTimeoutMs" name:"ExecuteTimeoutMs"`
			InputDataType    *string `json:"InputDataType" name:"InputDataType"`
			Ks3Region        *string `json:"Ks3Region" name:"Ks3Region"`
			InBucket         *string `json:"InBucket" name:"InBucket"`
			OutBucket        *string `json:"OutBucket" name:"OutBucket"`
			InObjectName     *string `json:"InObjectName" name:"InObjectName"`
			OutObjectName    *string `json:"OutObjectName" name:"OutObjectName"`
			FailReasonCode   *string `json:"FailReasonCode" name:"FailReasonCode"`
			FailReason       *string `json:"FailReason" name:"FailReason"`
			FailReasonDesc   *string `json:"FailReasonDesc" name:"FailReasonDesc"`
		} `json:"JobMetadata" name:"JobMetadata"`
	} `json:"Data"`
}

func (r *DescribeBatchInferenceJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchInferenceJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchInferenceJobRequest struct {
	*ksyunhttp.BaseRequest
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *DeleteBatchInferenceJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBatchInferenceJobResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteBatchInferenceJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchInferenceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableModelsRequest struct {
	*ksyunhttp.BaseRequest
	ModelIds []*string `json:"ModelIds,omitempty" name:"ModelIds"`
}

func (r *EnableModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EnableModelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"Success" name:"Success"`
}

func (r *EnableModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelQuotasRequest struct {
	*ksyunhttp.BaseRequest
	Marker     *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	Keyword    *string `json:"Keyword,omitempty" name:"Keyword"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeModelQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModelQuotasResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
	Data       []struct {
		ModelId         *string `json:"ModelId" name:"ModelId"`
		ModelName       *string `json:"ModelName" name:"ModelName"`
		Description     *string `json:"Description" name:"Description"`
		Icon            *string `json:"Icon" name:"Icon"`
		FreeQuotaStatus *string `json:"FreeQuotaStatus" name:"FreeQuotaStatus"`
		ActiveStatus    *int    `json:"ActiveStatus" name:"ActiveStatus"`
		FreeTotalQuota  *int64  `json:"FreeTotalQuota" name:"FreeTotalQuota"`
		FreeUsedQuota   *int64  `json:"FreeUsedQuota" name:"FreeUsedQuota"`
		BatchSupported  *bool   `json:"BatchSupported" name:"BatchSupported"`
		VideoModel      *bool   `json:"VideoModel" name:"VideoModel"`
		Tpm             *int    `json:"Tpm" name:"Tpm"`
		Rpm             *int    `json:"Rpm" name:"Rpm"`
	} `json:"Data"`
}

func (r *DescribeModelQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableModelsRequest struct {
	*ksyunhttp.BaseRequest
	ModelIds []*string `json:"ModelIds,omitempty" name:"ModelIds"`
}

func (r *DisableModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisableModelsResponse struct {
	*ksyunhttp.BaseResponse
	Success   *bool   `json:"Success" name:"Success"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DisableModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableOverFreeLimitRequest struct {
	*ksyunhttp.BaseRequest
	ModelIds []*string `json:"ModelIds,omitempty" name:"ModelIds"`
}

func (r *EnableOverFreeLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EnableOverFreeLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"Success" name:"Success"`
}

func (r *EnableOverFreeLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableOverFreeLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableOverFreeLimitRequest struct {
	*ksyunhttp.BaseRequest
	ModelIds []*string `json:"ModelIds,omitempty" name:"ModelIds"`
}

func (r *DisableOverFreeLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisableOverFreeLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"Success" name:"Success"`
}

func (r *DisableOverFreeLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableOverFreeLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrainJobEventsRequest struct {
	*ksyunhttp.BaseRequest
	ResourcePoolId *string `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	TrainJobId     *string `json:"TrainJobId,omitempty" name:"TrainJobId"`
}

func (r *DescribeTrainJobEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTrainJobEventsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	DataSet   []struct {
		FirstSeen *string `json:"FirstSeen" name:"FirstSeen"`
		LastSeen  *string `json:"LastSeen" name:"LastSeen"`
		Type      *string `json:"Type" name:"Type"`
		Object    struct {
			Kind      *string `json:"Kind" name:"Kind"`
			Namespace *string `json:"Namespace" name:"Namespace"`
			Name      *string `json:"Name" name:"Name"`
		} `json:"Object" name:"Object"`
		Reason  *string `json:"Reason" name:"Reason"`
		Message *string `json:"Message" name:"Message"`
		Source  struct {
			Component *string `json:"component" name:"component"`
		} `json:"Source" name:"Source"`
		Count *int `json:"Count" name:"Count"`
	} `json:"DataSet"`
}

func (r *DescribeTrainJobEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrainJobEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobId *string `json:"TrainJobId,omitempty" name:"TrainJobId"`
}

func (r *StopTrainJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopTrainJobResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TrainJobId *string `json:"TrainJobId" name:"TrainJobId"`
}

func (r *StopTrainJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopTrainJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobId *string `json:"TrainJobId,omitempty" name:"TrainJobId"`
}

func (r *StartTrainJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartTrainJobResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TrainJobId *string `json:"TrainJobId" name:"TrainJobId"`
}

func (r *StartTrainJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartTrainJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobId *string `json:"TrainJobId,omitempty" name:"TrainJobId"`
}

func (r *DeleteTrainJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteTrainJobResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TrainJobId *string `json:"TrainJobId" name:"TrainJobId"`
}

func (r *DeleteTrainJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrainJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobId *string `json:"TrainJobId,omitempty" name:"TrainJobId"`
	Priority   *string `json:"Priority,omitempty" name:"Priority"`
}

func (r *ModifyTrainJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyTrainJobResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TrainJobId *string `json:"TrainJobId" name:"TrainJobId"`
}

func (r *ModifyTrainJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrainJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrainJobPodLogsRequest struct {
	*ksyunhttp.BaseRequest
	ResourcePoolId *string `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	TrainJobId     *string `json:"TrainJobId,omitempty" name:"TrainJobId"`
	PodName        *string `json:"PodName,omitempty" name:"PodName"`
	SinceSeconds   *int    `json:"SinceSeconds,omitempty" name:"SinceSeconds"`
	TailLines      *int    `json:"TailLines,omitempty" name:"TailLines"`
}

func (r *DescribeTrainJobPodLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTrainJobPodLogsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	PodLogs   *string `json:"PodLogs" name:"PodLogs"`
}

func (r *DescribeTrainJobPodLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrainJobPodLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrainJobPodsRequest struct {
	*ksyunhttp.BaseRequest
	Marker     *string                       `json:"Marker,omitempty" name:"Marker"`
	MaxResults *string                       `json:"MaxResults,omitempty" name:"MaxResults"`
	TrainJobId *string                       `json:"TrainJobId,omitempty" name:"TrainJobId"`
	Filter     []*DescribeTrainJobPodsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeTrainJobPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTrainJobPodsResponse struct {
	*ksyunhttp.BaseResponse
	MaxResults     *int `json:"MaxResults" name:"MaxResults"`
	TrainJobPodSet []struct {
		TrainJobId *string `json:"TrainJobId" name:"TrainJobId"`
		Role       *string `json:"Role" name:"Role"`
		StartTimes *int    `json:"StartTimes" name:"StartTimes"`
		PodName    *string `json:"PodName" name:"PodName"`
		Status     struct {
			State          *string `json:"State" name:"State"`
			ContainerState *string `json:"ContainerState" name:"ContainerState"`
			SubmitTime     *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime      *string `json:"StartTime" name:"StartTime"`
			EndTime        *string `json:"EndTime" name:"EndTime"`
			Ip             *string `json:"Ip" name:"Ip"`
			RestartCount   *int    `json:"RestartCount" name:"RestartCount"`
		} `json:"Status" name:"Status"`
		ContainerName *string `json:"ContainerName" name:"ContainerName"`
		Kind          *string `json:"Kind" name:"Kind"`
		NameSpace     *string `json:"NameSpace" name:"NameSpace"`
	} `json:"TrainJobPodSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Marker     *int    `json:"Marker" name:"Marker"`
}

func (r *DescribeTrainJobPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrainJobPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInferencesRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId []*string                   `json:"InferenceId,omitempty" name:"InferenceId"`
	Filter      []*DescribeInferencesFilter `json:"Filter,omitempty" name:"Filter"`
	PageSize    *int                        `json:"PageSize,omitempty" name:"PageSize"`
	Page        *int                        `json:"Page,omitempty" name:"Page"`
}

func (r *DescribeInferencesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInferencesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	PageSize     *int    `json:"PageSize" name:"PageSize"`
	Page         *int    `json:"Page" name:"Page"`
	TotalCount   *int    `json:"TotalCount" name:"TotalCount"`
	InferenceSet []struct {
		InferenceName  *string `json:"InferenceName" name:"InferenceName"`
		InferenceId    *string `json:"InferenceId" name:"InferenceId"`
		Status         *string `json:"Status" name:"Status"`
		IngressEnabled *bool   `json:"IngressEnabled" name:"IngressEnabled"`
		Description    *string `json:"Description" name:"Description"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		Duration       *string `json:"Duration" name:"Duration"`
		UpdateTime     *string `json:"UpdateTime" name:"UpdateTime"`
		ModelName      *string `json:"ModelName" name:"ModelName"`
		BaseImage      *string `json:"BaseImage" name:"BaseImage"`
		CustomImage    *string `json:"CustomImage" name:"CustomImage"`
		EntryPoint     *string `json:"EntryPoint" name:"EntryPoint"`
		AccessAddress  *string `json:"AccessAddress" name:"AccessAddress"`
		Port           *int    `json:"Port" name:"Port"`
		Env            []struct {
			Name  *string `json:"Name" name:"Name"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Env" name:"Env"`
		CmdOptions []struct {
			Name  *string `json:"Name" name:"Name"`
			Value *string `json:"Value" name:"Value"`
		} `json:"CmdOptions" name:"CmdOptions"`
		ResourcePoolId  *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		DeploymentType  *string `json:"DeploymentType" name:"DeploymentType"`
		Engine          *string `json:"Engine" name:"Engine"`
		AccessType      *string `json:"AccessType" name:"AccessType"`
		ClusterId       *string `json:"ClusterId" name:"ClusterId"`
		QueueName       *string `json:"QueueName" name:"QueueName"`
		GPUType         *string `json:"GPUType" name:"GPUType"`
		GPUNum          *string `json:"GPUNum" name:"GPUNum"`
		CPUNum          *int    `json:"CPUNum" name:"CPUNum"`
		Memory          *int    `json:"Memory" name:"Memory"`
		Distributed     *bool   `json:"Distributed" name:"Distributed"`
		NodeNum         *int    `json:"NodeNum" name:"NodeNum"`
		Replicas        *int    `json:"Replicas" name:"Replicas"`
		CurrentReplicas *int    `json:"CurrentReplicas" name:"CurrentReplicas"`
		VpcId           *string `json:"VpcId" name:"VpcId"`
		SubnetId        *string `json:"SubnetId" name:"SubnetId"`
		StorageConfigs  []struct {
			StorageConfigId   *string `json:"StorageConfigId" name:"StorageConfigId"`
			StorageConfigName *string `json:"StorageConfigName" name:"StorageConfigName"`
			StorageConfigType *string `json:"StorageConfigType" name:"StorageConfigType"`
			StorageType       *string `json:"StorageType" name:"StorageType"`
			MountPath         *string `json:"MountPath" name:"MountPath"`
		} `json:"StorageConfigs" name:"StorageConfigs"`
		AutoScaleEnable      *bool   `json:"AutoScaleEnable" name:"AutoScaleEnable"`
		ModelStorageEnabled  *bool   `json:"ModelStorageEnabled" name:"ModelStorageEnabled"`
		ModelStoragePath     *string `json:"ModelStoragePath" name:"ModelStoragePath"`
		ImageUrl             *string `json:"ImageUrl" name:"ImageUrl"`
		ImageId              *string `json:"ImageId" name:"ImageId"`
		ImageSource          *string `json:"ImageSource" name:"ImageSource"`
		ImageRegistryId      *string `json:"ImageRegistryId" name:"ImageRegistryId"`
		ImageRepoId          *string `json:"ImageRepoId" name:"ImageRepoId"`
		ImageTagId           *string `json:"ImageTagId" name:"ImageTagId"`
		ImageRegistryName    *string `json:"ImageRegistryName" name:"ImageRegistryName"`
		ImageRepoName        *string `json:"ImageRepoName" name:"ImageRepoName"`
		ImageTagName         *string `json:"ImageTagName" name:"ImageTagName"`
		RunOnCPU             *bool   `json:"RunOnCPU" name:"RunOnCPU"`
		HostNetworkEnabled   *bool   `json:"HostNetworkEnabled" name:"HostNetworkEnabled"`
		ConsistentHashEnable *bool   `json:"ConsistentHashEnable" name:"ConsistentHashEnable"`
		ConsistentHashHeader *string `json:"ConsistentHashHeader" name:"ConsistentHashHeader"`
	} `json:"InferenceSet"`
}

func (r *DescribeInferencesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInferencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetInferenceAutoScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId       *string                                         `json:"InferenceId,omitempty" name:"InferenceId"`
	AutoScaleEnable   *bool                                           `json:"AutoScaleEnable,omitempty" name:"AutoScaleEnable"`
	AutoScaleStrategy *SetInferenceAutoScaleStrategyAutoScaleStrategy `json:"AutoScaleStrategy,omitempty" name:"AutoScaleStrategy"`
}

func (r *SetInferenceAutoScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetInferenceAutoScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *SetInferenceAutoScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetInferenceAutoScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInferenceRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
}

func (r *DeleteInferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInferenceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *DeleteInferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInferenceRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
}

func (r *StopInferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopInferenceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *StopInferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInferenceDetailRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
}

func (r *GetInferenceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetInferenceDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	InferenceDetail struct {
		InferenceName  *string `json:"InferenceName" name:"InferenceName"`
		InferenceId    *string `json:"InferenceId" name:"InferenceId"`
		Status         *string `json:"Status" name:"Status"`
		IngressEnabled *bool   `json:"IngressEnabled" name:"IngressEnabled"`
		Description    *string `json:"Description" name:"Description"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		Duration       *string `json:"Duration" name:"Duration"`
		UpdateTime     *string `json:"UpdateTime" name:"UpdateTime"`
		ModelName      *string `json:"ModelName" name:"ModelName"`
		BaseImage      *string `json:"BaseImage" name:"BaseImage"`
		CustomImage    *string `json:"CustomImage" name:"CustomImage"`
		EntryPoint     *string `json:"EntryPoint" name:"EntryPoint"`
		AccessAddress  *string `json:"AccessAddress" name:"AccessAddress"`
		Port           *int    `json:"Port" name:"Port"`
		Env            []struct {
			Name  *string `json:"Name" name:"Name"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Env" name:"Env"`
		CmdOptions []struct {
			Name  *string `json:"Name" name:"Name"`
			Value *string `json:"Value" name:"Value"`
		} `json:"CmdOptions" name:"CmdOptions"`
		ResourcePoolId  *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		DeploymentType  *string `json:"DeploymentType" name:"DeploymentType"`
		Engine          *string `json:"Engine" name:"Engine"`
		AccessType      *string `json:"AccessType" name:"AccessType"`
		ClusterId       *string `json:"ClusterId" name:"ClusterId"`
		QueueName       *string `json:"QueueName" name:"QueueName"`
		GPUType         *string `json:"GPUType" name:"GPUType"`
		GPUNum          *string `json:"GPUNum" name:"GPUNum"`
		CPUNum          *int    `json:"CPUNum" name:"CPUNum"`
		Memory          *int    `json:"Memory" name:"Memory"`
		Distributed     *bool   `json:"Distributed" name:"Distributed"`
		NodeNum         *int    `json:"NodeNum" name:"NodeNum"`
		Replicas        *int    `json:"Replicas" name:"Replicas"`
		CurrentReplicas *int    `json:"CurrentReplicas" name:"CurrentReplicas"`
		VpcId           *string `json:"VpcId" name:"VpcId"`
		SubnetId        *string `json:"SubnetId" name:"SubnetId"`
		StorageConfigs  []struct {
			StorageConfigId   *string `json:"StorageConfigId" name:"StorageConfigId"`
			StorageConfigName *string `json:"StorageConfigName" name:"StorageConfigName"`
			StorageConfigType *string `json:"StorageConfigType" name:"StorageConfigType"`
			StorageType       *string `json:"StorageType" name:"StorageType"`
			MountPath         *string `json:"MountPath" name:"MountPath"`
		} `json:"StorageConfigs" name:"StorageConfigs"`
		AutoScaleEnable      *bool   `json:"AutoScaleEnable" name:"AutoScaleEnable"`
		ModelStorageEnabled  *bool   `json:"ModelStorageEnabled" name:"ModelStorageEnabled"`
		ModelStoragePath     *string `json:"ModelStoragePath" name:"ModelStoragePath"`
		ImageUrl             *string `json:"ImageUrl" name:"ImageUrl"`
		ImageId              *string `json:"ImageId" name:"ImageId"`
		ImageSource          *string `json:"ImageSource" name:"ImageSource"`
		ImageRegistryId      *string `json:"ImageRegistryId" name:"ImageRegistryId"`
		ImageRepoId          *string `json:"ImageRepoId" name:"ImageRepoId"`
		ImageTagId           *string `json:"ImageTagId" name:"ImageTagId"`
		ImageRegistryName    *string `json:"ImageRegistryName" name:"ImageRegistryName"`
		ImageRepoName        *string `json:"ImageRepoName" name:"ImageRepoName"`
		ImageTagName         *string `json:"ImageTagName" name:"ImageTagName"`
		RunOnCPU             *bool   `json:"RunOnCPU" name:"RunOnCPU"`
		HostNetworkEnabled   *bool   `json:"HostNetworkEnabled" name:"HostNetworkEnabled"`
		ConsistentHashEnable *bool   `json:"ConsistentHashEnable" name:"ConsistentHashEnable"`
		ConsistentHashHeader *string `json:"ConsistentHashHeader" name:"ConsistentHashHeader"`
	} `json:"InferenceDetail"`
}

func (r *GetInferenceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInferenceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInferenceRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
}

func (r *StartInferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartInferenceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *StartInferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInferenceRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId        *string                      `json:"InferenceId,omitempty" name:"InferenceId"`
	InferenceName      *string                      `json:"InferenceName,omitempty" name:"InferenceName"`
	Description        *string                      `json:"Description,omitempty" name:"Description"`
	EntryPoint         *string                      `json:"EntryPoint,omitempty" name:"EntryPoint"`
	ImageSource        *string                      `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageId            *string                      `json:"ImageId,omitempty" name:"ImageId"`
	ImageRegistryId    *string                      `json:"ImageRegistryId,omitempty" name:"ImageRegistryId"`
	ImageRepoId        *string                      `json:"ImageRepoId,omitempty" name:"ImageRepoId"`
	ImageTagId         *string                      `json:"ImageTagId,omitempty" name:"ImageTagId"`
	Env                []*ModifyInferenceEnv        `json:"Env,omitempty" name:"Env"`
	CmdOptions         []*ModifyInferenceCmdOptions `json:"CmdOptions,omitempty" name:"CmdOptions"`
	HostNetworkEnabled *bool                        `json:"HostNetworkEnabled,omitempty" name:"HostNetworkEnabled"`
}

func (r *ModifyInferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInferenceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *ModifyInferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetInferenceReplicasRequest struct {
	*ksyunhttp.BaseRequest
	InferenceId *string `json:"InferenceId,omitempty" name:"InferenceId"`
	Replicas    *int    `json:"Replicas,omitempty" name:"Replicas"`
}

func (r *SetInferenceReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetInferenceReplicasResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InferenceId *string `json:"InferenceId" name:"InferenceId"`
}

func (r *SetInferenceReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetInferenceReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolsRequest struct {
	*ksyunhttp.BaseRequest
	Sort             *string                        `json:"Sort,omitempty" name:"Sort"`
	Page             *int                           `json:"Page,omitempty" name:"Page"`
	PageSize         *int                           `json:"PageSize,omitempty" name:"PageSize"`
	ResourcePoolName *string                        `json:"ResourcePoolName,omitempty" name:"ResourcePoolName"`
	Component        *string                        `json:"Component,omitempty" name:"Component"`
	ResourcePoolId   []*string                      `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	Filter           []*DescribeResourcePoolsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeResourcePoolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeResourcePoolsResponse struct {
	*ksyunhttp.BaseResponse
	ResourcePoolSet []struct {
		ResourcePoolId        *string   `json:"ResourcePoolId" name:"ResourcePoolId"`
		ResourcePoolName      *string   `json:"ResourcePoolName" name:"ResourcePoolName"`
		Description           *string   `json:"Description" name:"Description"`
		ResourcePoolType      *string   `json:"ResourcePoolType" name:"ResourcePoolType"`
		VpcId                 *string   `json:"VpcId" name:"VpcId"`
		Status                *string   `json:"Status" name:"Status"`
		CreateTime            *string   `json:"CreateTime" name:"CreateTime"`
		GpuNodeNum            *int      `json:"GpuNodeNum" name:"GpuNodeNum"`
		AvailableGpuNodeNum   *int      `json:"AvailableGpuNodeNum" name:"AvailableGpuNodeNum"`
		GpuNum                *int      `json:"GpuNum" name:"GpuNum"`
		AvailableGpuNum       *int      `json:"AvailableGpuNum" name:"AvailableGpuNum"`
		PrometheusId          *string   `json:"PrometheusId" name:"PrometheusId"`
		ClusterId             *string   `json:"ClusterId" name:"ClusterId"`
		Purpose               *string   `json:"Purpose" name:"Purpose"`
		Components            []*string `json:"Components" name:"Components"`
		LogProjectName        *string   `json:"LogProjectName" name:"LogProjectName"`
		LogPoolName           *string   `json:"LogPoolName" name:"LogPoolName"`
		FileSystemId          *string   `json:"FileSystemId" name:"FileSystemId"`
		EnableKPFSPerformance *bool     `json:"EnableKPFSPerformance" name:"EnableKPFSPerformance"`
		EnableKlog            *bool     `json:"EnableKlog" name:"EnableKlog"`
	} `json:"ResourcePoolSet"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
	Page       *int    `json:"Page" name:"Page"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeResourcePoolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ResourcePoolId *string                                `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	PageSize       *int                                   `json:"PageSize,omitempty" name:"PageSize"`
	Page           *int                                   `json:"Page,omitempty" name:"Page"`
	InstanceName   *string                                `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceId     []*string                              `json:"InstanceId,omitempty" name:"InstanceId"`
	ProjectId      []*string                              `json:"ProjectId,omitempty" name:"ProjectId"`
	Filter         []*DescribeResourcePoolInstancesFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeResourcePoolInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeResourcePoolInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	PageSize                *int    `json:"PageSize" name:"PageSize"`
	Page                    *int    `json:"Page" name:"Page"`
	TotalCount              *int    `json:"TotalCount" name:"TotalCount"`
	ResourcePoolInstanceSet []struct {
		ResourcePoolId      *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		InstanceId          *string `json:"InstanceId" name:"InstanceId"`
		InstanceName        *string `json:"InstanceName" name:"InstanceName"`
		K8sRole             *string `json:"K8sRole" name:"K8sRole"`
		InstanceStatus      *string `json:"InstanceStatus" name:"InstanceStatus"`
		InstanceType        *string `json:"InstanceType" name:"InstanceType"`
		NodeType            *string `json:"NodeType" name:"NodeType"`
		InstanceIp          *string `json:"InstanceIp" name:"InstanceIp"`
		VpcId               *string `json:"VpcId" name:"VpcId"`
		AvailabilityZone    *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		ChargeType          *string `json:"ChargeType" name:"ChargeType"`
		ProjectId           *string `json:"ProjectId" name:"ProjectId"`
		NodeIcon            *string `json:"NodeIcon" name:"NodeIcon"`
		ServiceEndTime      *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		IsTrial             *bool   `json:"IsTrial" name:"IsTrial"`
		IsGpu               *bool   `json:"IsGpu" name:"IsGpu"`
		ResourcePoolType    *string `json:"ResourcePoolType" name:"ResourcePoolType"`
		GpuType             *string `json:"GpuType" name:"GpuType"`
		NetworkInterfaceSet []struct {
			VpcId                *string `json:"VpcId" name:"VpcId"`
			SubnetId             *string `json:"SubnetId" name:"SubnetId"`
			PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
		} `json:"NetworkInterfaceSet" name:"NetworkInterfaceSet"`
		Cpu struct {
			Allocated   *string `json:"Allocated" name:"Allocated"`
			Allocatable *string `json:"Allocatable" name:"Allocatable"`
		} `json:"Cpu" name:"Cpu"`
		Gpu struct {
			Allocated   *string `json:"Allocated" name:"Allocated"`
			Allocatable *string `json:"Allocatable" name:"Allocatable"`
		} `json:"Gpu" name:"Gpu"`
		Memory struct {
			Allocated   *string `json:"Allocated" name:"Allocated"`
			Allocatable *string `json:"Allocatable" name:"Allocatable"`
		} `json:"Memory" name:"Memory"`
		UnSchedulable *bool `json:"UnSchedulable" name:"UnSchedulable"`
	} `json:"ResourcePoolInstanceSet"`
}

func (r *DescribeResourcePoolInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePoolInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInferenceEndpointRequest struct {
	*ksyunhttp.BaseRequest
	EndpointName *string                           `json:"EndpointName,omitempty" name:"EndpointName"`
	ProjectId    *string                           `json:"ProjectId,omitempty" name:"ProjectId"`
	ModelName    *string                           `json:"ModelName,omitempty" name:"ModelName"`
	RateLimit    *CreateInferenceEndpointRateLimit `json:"RateLimit ,omitempty" name:"RateLimit "`
	ModelId      *string                           `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *CreateInferenceEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInferenceEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *CreateInferenceEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInferenceEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInferenceEndpointsRequest struct {
	*ksyunhttp.BaseRequest
	EndpointId   []*string                           `json:"EndpointId,omitempty" name:"EndpointId"`
	EndpointName *string                             `json:"EndpointName,omitempty" name:"EndpointName"`
	Marker       *int                                `json:"Marker,omitempty" name:"Marker"`
	MaxResults   *int                                `json:"MaxResults,omitempty" name:"MaxResults"`
	ProjectId    []*string                           `json:"ProjectId,omitempty" name:"ProjectId"`
	Filter       []*DescribeInferenceEndpointsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeInferenceEndpointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInferenceEndpointsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Endpoints  []struct {
		ModelName        *string `json:"ModelName" name:"ModelName"`
		ModelType        *int    `json:"ModelType" name:"ModelType"`
		ModelSource      *string `json:"ModelSource" name:"ModelSource"`
		Status           *string `json:"Status" name:"Status"`
		RateLimitEnabled *bool   `json:"RateLimitEnabled" name:"RateLimitEnabled"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		CreateUserId     *string `json:"CreateUserId" name:"CreateUserId"`
		CreateUserName   *string `json:"CreateUserName" name:"CreateUserName"`
		RateLimitConfig  []struct {
			RPM *int `json:"RPM" name:"RPM"`
			TPM *int `json:"TPM" name:"TPM"`
		} `json:"RateLimitConfig" name:"RateLimitConfig"`
		EndpointId   *string `json:"EndpointId" name:"EndpointId"`
		EndpointName *string `json:"EndpointName" name:"EndpointName"`
		Description  *string `json:"Description" name:"Description"`
		ProjectId    *string `json:"ProjectId" name:"ProjectId"`
		ProjectName  *string `json:"ProjectName" name:"ProjectName"`
	} `json:"Endpoints"`
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
}

func (r *DescribeInferenceEndpointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInferenceEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableEndpointRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	EndpointId *string                           `json:"EndpointId,omitempty" name:"EndpointId"`
	RateLimit  *EnableEndpointRateLimitRateLimit `json:"RateLimit,omitempty" name:"RateLimit"`
}

func (r *EnableEndpointRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EnableEndpointRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *EnableEndpointRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableEndpointRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInferenceEndpointRequest struct {
	*ksyunhttp.BaseRequest
	EndpointName *string                           `json:"EndpointName,omitempty" name:"EndpointName"`
	ProjectId    *string                           `json:"ProjectId,omitempty" name:"ProjectId"`
	ModelName    *string                           `json:"ModelName,omitempty" name:"ModelName"`
	RateLimit    *UpdateInferenceEndpointRateLimit `json:"RateLimit ,omitempty" name:"RateLimit "`
	EndpointId   *string                           `json:"EndpointId,omitempty" name:"EndpointId"`
}

func (r *UpdateInferenceEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateInferenceEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *UpdateInferenceEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInferenceEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInferenceEndpointRequest struct {
	*ksyunhttp.BaseRequest
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`
}

func (r *StartInferenceEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartInferenceEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *StartInferenceEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInferenceEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInferenceEndpointRequest struct {
	*ksyunhttp.BaseRequest
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`
}

func (r *StopInferenceEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopInferenceEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *StopInferenceEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInferenceEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInferenceEndpointRequest struct {
	*ksyunhttp.BaseRequest
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`
}

func (r *DeleteInferenceEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInferenceEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *DeleteInferenceEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInferenceEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableEndpointRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`
}

func (r *DisableEndpointRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisableEndpointRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	EndpointId *string `json:"EndpointId" name:"EndpointId"`
}

func (r *DisableEndpointRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableEndpointRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolInstanceTasksRequest struct {
	*ksyunhttp.BaseRequest
	ResourcePoolId *string `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	InstanceId     *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TaskType       *string `json:"TaskType,omitempty" name:"TaskType"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
	Page           *int    `json:"Page,omitempty" name:"Page"`
}

func (r *DescribeResourcePoolInstanceTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeResourcePoolInstanceTasksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	PageSize        *int    `json:"PageSize" name:"PageSize"`
	Page            *int    `json:"Page" name:"Page"`
	TotalCount      *int    `json:"TotalCount" name:"TotalCount"`
	InstanceTaskSet []struct {
		PodInstanceId *string `json:"PodInstanceId" name:"PodInstanceId"`
		TaskName      *string `json:"TaskName" name:"TaskName"`
		TaskId        *string `json:"TaskId" name:"TaskId"`
		TaskType      *string `json:"TaskType" name:"TaskType"`
		GPUNum        *int    `json:"GPUNum" name:"GPUNum"`
		CPUNum        *int    `json:"CPUNum" name:"CPUNum"`
		MemeryNum     *int    `json:"MemeryNum" name:"MemeryNum"`
		CreateUser    *string `json:"CreateUser" name:"CreateUser"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		StartTime     *string `json:"StartTime" name:"StartTime"`
		EndTime       *string `json:"EndTime" name:"EndTime"`
		Status        *string `json:"Status" name:"Status"`
	} `json:"InstanceTaskSet"`
}

func (r *DescribeResourcePoolInstanceTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePoolInstanceTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetKcrPersonalTokenRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *SetKcrPersonalTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetKcrPersonalTokenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetKcrPersonalTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetKcrPersonalTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueuesRequest struct {
	*ksyunhttp.BaseRequest
	QueueId  []*string               `json:"QueueId,omitempty" name:"QueueId"`
	Page     *int                    `json:"Page,omitempty" name:"Page"`
	PageSize *int                    `json:"PageSize,omitempty" name:"PageSize"`
	Filter   []*DescribeQueuesFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeQueuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeQueuesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
	Page       *int    `json:"Page" name:"Page"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	QueueSet   []struct {
		QueueId        *string `json:"QueueId" name:"QueueId"`
		ResourcePoolId *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		QueueName      *string `json:"QueueName" name:"QueueName"`
		Description    *string `json:"Description" name:"Description"`
		Reclaimable    *bool   `json:"Reclaimable" name:"Reclaimable"`
		AllowBorrowing *bool   `json:"AllowBorrowing" name:"AllowBorrowing"`
		Status         struct {
			State   *string `json:"State" name:"State"`
			Running *int    `json:"Running" name:"Running"`
			Inqueue *int    `json:"Inqueue" name:"Inqueue"`
			Pending *int    `json:"Pending" name:"Pending"`
			Queuing *int    `json:"Queuing" name:"Queuing"`
		} `json:"Status" name:"Status"`
		CreatorId  *string `json:"CreatorId" name:"CreatorId"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
		AccessList []struct {
			UserId     *string `json:"UserId" name:"UserId"`
			Permission *string `json:"Permission" name:"Permission"`
		} `json:"AccessList" name:"AccessList"`
		Capability struct {
			CPUNum    *int `json:"CPUNum" name:"CPUNum"`
			MemoryNum *int `json:"MemoryNum" name:"MemoryNum"`
			GPUInfos  []struct {
				GPUType *string `json:"GPUType" name:"GPUType"`
				GPUNum  *int    `json:"GPUNum" name:"GPUNum"`
			} `json:"GPUInfos"`
		} `json:"Capability" name:"Capability"`
		Allocated struct {
			CPUNum    *int    `json:"CPUNum" name:"CPUNum"`
			MemoryNum *int    `json:"MemoryNum" name:"MemoryNum"`
			GPUNum    *string `json:"GPUNum" name:"GPUNum"`
		} `json:"Allocated" name:"Allocated"`
	} `json:"QueueSet"`
}

func (r *DescribeQueuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
