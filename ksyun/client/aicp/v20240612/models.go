package v20240612

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type ModifyNotebookStorageConfigs struct {
	StorageConfigId   *string `json:"StorageConfigId,omitempty" name:"StorageConfigId"`
	MountPath         *string `json:"MountPath,omitempty" name:"MountPath"`
	StorageConfigType *string `json:"StorageConfigType,omitempty" name:"StorageConfigType"`
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

func (r *SaveNotebookImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SaveNotebookImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
	DisplayName            *string                         `json:"DisplayName,omitempty" name:"DisplayName"`
	Description            *string                         `json:"Description,omitempty" name:"Description"`
	ImageId                *string                         `json:"ImageId,omitempty" name:"ImageId"`
	QueueName              *string                         `json:"QueueName,omitempty" name:"QueueName"`
	GPUType                *string                         `json:"GPUType,omitempty" name:"GPUType"`
	GPUNumber              *int                            `json:"GPUNumber,omitempty" name:"GPUNumber"`
	AccessType             *string                         `json:"AccessType,omitempty" name:"AccessType"`
	AllocationId           *string                         `json:"AllocationId,omitempty" name:"AllocationId"`
	EnableSsh              *string                         `json:"EnableSsh,omitempty" name:"EnableSsh"`
	SshPort                *int                            `json:"SshPort,omitempty" name:"SshPort"`
	EnablePublicNetworkSsh *bool                           `json:"EnablePublicNetworkSsh,omitempty" name:"EnablePublicNetworkSsh"`
	SshAuthorizedKeys      *string                         `json:"SshAuthorizedKeys,omitempty" name:"SshAuthorizedKeys"`
	CpuNum                 *int                            `json:"CpuNum,omitempty" name:"CpuNum"`
	Memory                 *int                            `json:"Memory,omitempty" name:"Memory"`
	StorageConfigs         []*ModifyNotebookStorageConfigs `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
}

func (r *ModifyNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
	Status     *string                    `json:"Status,omitempty" name:"Status"`
	Marker     *int                       `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	State      *string                    `json:"State,omitempty" name:"State"`
	Filter     []*DescribeNotebooksFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeNotebooksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNotebooksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNotebooksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebooksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Notebooks []struct {
		NotebookId  *string `json:"NotebookId" name:"NotebookId"`
		Name        *string `json:"Name" name:"Name"`
		Description *string `json:"Description" name:"Description"`
		Status      struct {
			State         *string `json:"State" name:"State"`
			SubmitTime    *string `json:"SubmitTime" name:"SubmitTime"`
			StartTime     *string `json:"StartTime" name:"StartTime"`
			EndTime       *string `json:"EndTime" name:"EndTime"`
			Message       *string `json:"Message" name:"Message"`
			ExecutionTime *string `json:"ExecutionTime" name:"ExecutionTime"`
		} `json:"Status" name:"Status"`
		ClusterId              *string `json:"ClusterId" name:"ClusterId"`
		ImageId                *string `json:"ImageId" name:"ImageId"`
		GPUType                *string `json:"GPUType" name:"GPUType"`
		GPUNumber              *int    `json:"GPUNumber" name:"GPUNumber"`
		CreateUser             *string `json:"CreateUser" name:"CreateUser"`
		ResourcePoolId         *string `json:"ResourcePoolId" name:"ResourcePoolId"`
		QueueName              *string `json:"QueueName" name:"QueueName"`
		AccessType             *string `json:"AccessType" name:"AccessType"`
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime             *string `json:"UpdateTime" name:"UpdateTime"`
		AllocationId           *string `json:"AllocationId" name:"AllocationId"`
		EnableSsh              *bool   `json:"EnableSsh" name:"EnableSsh"`
		EnablePublicNetworkSsh *bool   `json:"EnablePublicNetworkSsh" name:"EnablePublicNetworkSsh"`
		SshPort                *int    `json:"SshPort" name:"SshPort"`
		SshAuthorizedKeys      *string `json:"SshAuthorizedKeys" name:"SshAuthorizedKeys"`
		PodIp                  *string `json:"PodIp" name:"PodIp"`
		ExternalIp             *string `json:"ExternalIp" name:"ExternalIp"`
		CpuNum                 *int    `json:"CpuNum" name:"CpuNum"`
		Memory                 *int    `json:"Memory" name:"Memory"`
		StorageConfigs         []struct {
			StorageConfigId   *string `json:"StorageConfigId" name:"StorageConfigId"`
			MountPath         *string `json:"MountPath" name:"MountPath"`
			StorageConfigType *string `json:"StorageConfigType" name:"StorageConfigType"`
		} `json:"StorageConfigs" name:"StorageConfigs"`
	} `json:"Notebooks"`
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
	DisplayName            *string                         `json:"DisplayName,omitempty" name:"DisplayName"`
	Description            *string                         `json:"Description,omitempty" name:"Description"`
	ImageId                *string                         `json:"ImageId,omitempty" name:"ImageId"`
	QueueName              *string                         `json:"QueueName,omitempty" name:"QueueName"`
	GPUType                *string                         `json:"GPUType,omitempty" name:"GPUType"`
	GPUNumber              *int                            `json:"GPUNumber,omitempty" name:"GPUNumber"`
	AccessType             *string                         `json:"AccessType,omitempty" name:"AccessType"`
	EnablePublicNetworkSsh *bool                           `json:"EnablePublicNetworkSsh,omitempty" name:"EnablePublicNetworkSsh"`
	AllocationId           *string                         `json:"AllocationId,omitempty" name:"AllocationId"`
	CpuNum                 *int                            `json:"CpuNum,omitempty" name:"CpuNum"`
	Memory                 *int                            `json:"Memory,omitempty" name:"Memory"`
	EnableSsh              *bool                           `json:"EnableSsh,omitempty" name:"EnableSsh"`
	SshAuthorizedKeys      *string                         `json:"SshAuthorizedKeys,omitempty" name:"SshAuthorizedKeys"`
	SshPort                *int                            `json:"SshPort,omitempty" name:"SshPort"`
	StorageConfigs         []*CreateNotebookStorageConfigs `json:"StorageConfigs,omitempty" name:"StorageConfigs"`
	ResourcePoolId         *string                         `json:"ResourcePoolId,omitempty" name:"ResourcePoolId"`
}

func (r *CreateNotebookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type ListDatasetHostnameRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ListDatasetHostnameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDatasetHostnameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListDatasetHostnameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListDatasetHostnameResponse struct {
	*ksyunhttp.BaseResponse
	Hostnames []*string `json:"Hostnames" name:"Hostnames"`
	RequestId *string   `json:"RequestId" name:"RequestId"`
}

func (r *ListDatasetHostnameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDatasetHostnameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDatasetTopicRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ListDatasetTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDatasetTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListDatasetTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListDatasetTopicResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ListDatasetTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDatasetTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDatasetRequest struct {
	*ksyunhttp.BaseRequest
	Page        *int    `json:"Page,omitempty" name:"Page"`
	PageSize    *int    `json:"PageSize,omitempty" name:"PageSize"`
	DatasetName *string `json:"DatasetName,omitempty" name:"DatasetName"`
	Topic       *string `json:"Topic,omitempty" name:"Topic"`
	HostName    *string `json:"HostName,omitempty" name:"HostName"`
	Keywords    *string `json:"Keywords,omitempty" name:"Keywords"`
	Source      *string `json:"Source,omitempty" name:"Source"`
}

func (r *ListDatasetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListDatasetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	DatasetList []struct {
		DatasetId   *string   `json:"DatasetId" name:"DatasetId"`
		DatasetName *string   `json:"DatasetName" name:"DatasetName"`
		DatasetDesc *string   `json:"DatasetDesc" name:"DatasetDesc"`
		Topic       *string   `json:"Topic" name:"Topic"`
		TagList     []*string `json:"TagList" name:"TagList"`
		HostName    *string   `json:"HostName" name:"HostName"`
		Keywords    *string   `json:"Keywords" name:"Keywords"`
		CreateTime  *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"DatasetList"`
}

func (r *ListDatasetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatasetRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDatasetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatasetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Dataset   struct {
		DatasetId *string `json:"DatasetId" name:"DatasetId"`
	} `json:"Dataset"`
}

func (r *DescribeDatasetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDatasetRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ImportDatasetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImportDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportDatasetResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ImportDatasetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDatasetResponse) FromJsonString(s string) error {
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

func (r *StopNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StopNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *StartNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *GetWebIdeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetWebIdeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeNotebookEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNotebookEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
