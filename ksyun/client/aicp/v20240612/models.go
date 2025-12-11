package v20240612

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

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
type CreateTrainJobFrameworkReplicas struct {
	Chief     *int `json:"Chief,omitempty" name:"Chief"`
	Evaluator *int `json:"Evaluator,omitempty" name:"Evaluator"`
	PS        *int `json:"PS,omitempty" name:"PS"`
	Master    *int `json:"Master,omitempty" name:"Master"`
}
type CreateTrainJobEnvs struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateTrainJobStorageConfigs struct {
	StorageConfigId   *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	MountPath         *string `json:"MountPath,omitempty" name:"MountPath"`
	StorageConfigType *string `json:"StorageConfigType,omitempty" name:"StorageConfigType"`
}
type DescribeTrainJobFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeTrainJobPodsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeResourcePoolsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeResourcePoolInstancesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
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
	GPUNumber              *int                            `json:"GPUNumber,omitempty" name:"GPUNumber"`
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
	Marker     *int                       `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	State      *string                    `json:"State,omitempty" name:"State"`
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
		GPUNumber              *int    `json:"GPUNumber" name:"GPUNumber"`
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
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
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
	GPUNumber              *int                            `json:"GPUNumber,omitempty" name:"GPUNumber"`
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

type DescribeNotebookEventsRequest struct {
	*ksyunhttp.BaseRequest
	NotebookId *string `json:"NotebookId,omitempty" name:"NotebookId"`
}

func (r *DescribeNotebookEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNotebookEventsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	DataSet   []struct {
		FirstSeen *string `json:"FirstSeen" name:"FirstSeen"`
		LastSeen  *string `json:"LastSeen" name:"LastSeen"`
		Type      *string `json:"Type" name:"Type"`
		Object    struct {
			Kind *string `json:"Kind" name:"Kind"`
			Name *string `json:"Name" name:"Name"`
		} `json:"Object" name:"Object"`
		Reason  *string `json:"Reason" name:"Reason"`
		Message *string `json:"Message" name:"Message"`
		Source  struct {
			Component *string `json:"Component" name:"Component"`
		} `json:"Source" name:"Source"`
		Count *int `json:"Count" name:"Count"`
	} `json:"DataSet"`
}

func (r *DescribeNotebookEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNotebookEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookLogRequest struct {
	*ksyunhttp.BaseRequest
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

type CreateTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobName       *string                          `json:"TrainJobName,omitempty" name:"TrainJobName"`
	Description        *string                          `json:"Description,omitempty" name:"Description"`
	ResourcePoolId     *string                          `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
	QueueName          *string                          `json:"QueueName,omitempty" name:"QueueName"`
	Priority           *string                          `json:"Priority,omitempty" name:"Priority"`
	Command            *string                          `json:"Command,omitempty" name:"Command"`
	Framework          *string                          `json:"Framework,omitempty" name:"Framework"`
	ImageSource        *string                          `json:"ImageSource,omitempty" name:"ImageSource"`
	FrameworkReplicas  *CreateTrainJobFrameworkReplicas `json:"FrameworkReplicas,omitempty" name:"FrameworkReplicas"`
	RestartPolicy      *string                          `json:"RestartPolicy,omitempty" name:"RestartPolicy"`
	Envs               []*CreateTrainJobEnvs            `json:"Envs,omitempty" name:"Envs"`
	SupportTensorboard *bool                            `json:"SupportTensorboard,omitempty" name:"SupportTensorboard"`
	ImageId            *string                          `json:"ImageId,omitempty" name:"ImageId"`
	ImageRegistryId    *string                          `json:"ImageRegistryId,omitempty" name:"ImageRegistryId"`
	ImageRepoId        *string                          `json:"ImageRepoId,omitempty" name:"ImageRepoId"`
	ImageTagId         *string                          `json:"ImageTagId,omitempty" name:"ImageTagId"`
	GPUType            *string                          `json:"GPUType,omitempty" name:"GPUType"`
	GPUNumber          *int                             `json:"GPUNumber,omitempty" name:"GPUNumber"`
	CPUNum             *int                             `json:"CPUNum,omitempty" name:"CPUNum"`
	Memory             *int                             `json:"Memory,omitempty" name:"Memory"`
	StorageConfigs     []*CreateTrainJobStorageConfigs  `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
	AccessType         *string                          `json:"AccessType,omitempty" name:"AccessType"`
	MaxRuntime         *int                             `json:"MaxRuntime,omitempty" name:"MaxRuntime"`
	SelfHealing        *bool                            `json:"SelfHealing,omitempty" name:"SelfHealing"`
	RunOnCPU           *bool                            `json:"RunOnCPU,omitempty" name:"RunOnCPU"`
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
			Kind            *string `json:"Kind" name:"Kind"`
			Namespace       *string `json:"Namespace" name:"Namespace"`
			Name            *string `json:"Name" name:"Name"`
			UID             *string `json:"UID" name:"UID"`
			APIVersion      *string `json:"APIVersion" name:"APIVersion"`
			ResourceVersion *string `json:"ResourceVersion" name:"ResourceVersion"`
			FieldPath       *string `json:"FieldPath" name:"FieldPath"`
		} `json:"Object" name:"Object"`
		Reason  *string `json:"Reason" name:"Reason"`
		Message *string `json:"Message" name:"Message"`
		Source  struct {
			Component *string `json:"component" name:"component"`
			Host      *string `json:"host" name:"host"`
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

type DescribeTrainJobRequest struct {
	*ksyunhttp.BaseRequest
	TrainJobId []*string                 `json:"TrainJobId,omitempty" name:"TrainJobId"`
	Filter     []*DescribeTrainJobFilter `json:"Filter,omitempty" name:"Filter"`
	Marker     *int                      `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int                      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeTrainJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTrainJobResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	TotalCount  *int    `json:"TotalCount" name:"TotalCount"`
	Marker      *int    `json:"Marker" name:"Marker"`
	MaxResults  *int    `json:"MaxResults" name:"MaxResults"`
	TrainJobSet []struct {
		TrainJobName      *string `json:"TrainJobName" name:"TrainJobName"`
		TrainJobId        *string `json:"TrainJobId" name:"TrainJobId"`
		ResourcePoolName  *string `json:"ResourcePoolName" name:"ResourcePoolName"`
		ResourcePoolId    *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		ResourcePoolType  *string `json:"ResourcePoolType" name:"ResourcePoolType"`
		Namespace         *string `json:"Namespace" name:"Namespace"`
		QueueName         *string `json:"QueueName" name:"QueueName"`
		Description       *string `json:"Description" name:"Description"`
		Priority          *string `json:"Priority" name:"Priority"`
		Framework         *string `json:"Framework" name:"Framework"`
		AccessType        *string `json:"AccessType" name:"AccessType"`
		ImageId           *string `json:"ImageId" name:"ImageId"`
		ImageSource       *string `json:"ImageSource" name:"ImageSource"`
		ImageRegistryId   *string `json:"ImageRegistryId" name:"ImageRegistryId"`
		ImageRepoId       *string `json:"ImageRepoId" name:"ImageRepoId"`
		ImageTagId        *string `json:"ImageTagId" name:"ImageTagId"`
		CreateUserName    *string `json:"CreateUserName" name:"CreateUserName"`
		FrameworkReplicas struct {
			Master    *int `json:"Master" name:"Master"`
			Worker    *int `json:"Worker" name:"Worker"`
			Chief     *int `json:"Chief" name:"Chief"`
			Evaluator *int `json:"Evaluator" name:"Evaluator"`
			PS        *int `json:"PS" name:"PS"`
		} `json:"FrameworkReplicas" name:"FrameworkReplicas"`
		GPUType       *string `json:"GPUType" name:"GPUType"`
		GPUNumber     *int    `json:"GPUNumber" name:"GPUNumber"`
		CPUNum        *int    `json:"CPUNum" name:"CPUNum"`
		Memory        *int    `json:"Memory" name:"Memory"`
		RestartPolicy *string `json:"RestartPolicy" name:"RestartPolicy"`
		Status        struct {
			State         *string `json:"State" name:"State"`
			PodSucceedNum *int    `json:"PodSucceedNum" name:"PodSucceedNum"`
			PodFailedNum  *int    `json:"PodFailedNum" name:"PodFailedNum"`
			SubmitTime    *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime     *string `json:"StartTime" name:"StartTime"`
			QueueTime     *int    `json:"QueueTime" name:"QueueTime"`
			EndTime       *string `json:"EndTime" name:"EndTime"`
			Message       *string `json:"Message" name:"Message"`
		} `json:"Status" name:"Status"`
		CreateUserId *string `json:"CreateUserId" name:"CreateUserId"`
		SelfHealing  *bool   `json:"SelfHealing" name:"SelfHealing"`
		Envs         []struct {
			Name  *string `json:"Name" name:"Name"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Envs" name:"Envs"`
		Command            *string `json:"Command" name:"Command"`
		MaxRuntime         *int    `json:"MaxRuntime" name:"MaxRuntime"`
		EnableHostNetwork  *bool   `json:"EnableHostNetwork" name:"EnableHostNetwork"`
		RunOnCPU           *bool   `json:"RunOnCPU" name:"RunOnCPU"`
		SupportTensorboard *bool   `json:"SupportTensorboard" name:"SupportTensorboard"`
		StorageConfigs     []struct {
			StorageConfigId   *string `json:"StorageConfigId" name:"StorageConfigId"`
			StorageConfigType *string `json:"StorageConfigType" name:"StorageConfigType"`
			MountPath         *string `json:"MountPath" name:"MountPath"`
		} `json:"StorageConfigs" name:"StorageConfigs"`
	} `json:"TrainJobSet"`
}

func (r *DescribeTrainJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrainJobResponse) FromJsonString(s string) error {
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
			Allocated   *int `json:"Allocated" name:"Allocated"`
			Allocatable *int `json:"Allocatable" name:"Allocatable"`
		} `json:"Cpu" name:"Cpu"`
		Gpu struct {
			Allocated   *int `json:"Allocated" name:"Allocated"`
			Allocatable *int `json:"Allocatable" name:"Allocatable"`
		} `json:"Gpu" name:"Gpu"`
		Memory struct {
			Allocated   *int `json:"Allocated" name:"Allocated"`
			Allocatable *int `json:"Allocatable" name:"Allocatable"`
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
