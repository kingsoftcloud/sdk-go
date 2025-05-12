package v20250501

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateJobSpace_storages struct {
	Space_storage_id       *int    `json:"space_storage_id,omitempty" name:"space_storage_id"`
	Space_storage_location *string `json:"space_storage_location,omitempty" name:"space_storage_location"`
}
type CreateJobEnv struct {
	Key1 *string `json:"key1,omitempty" name:"key1"`
}

type GetImageRequest struct {
	*ksyunhttp.BaseRequest
	Id *int `json:"id,omitempty" name:"id"`
}

func (r *GetImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetImageResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  struct {
		Id          *int    `json:"Id" name:"Id"`
		CreateAt    *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt    *int    `json:"UpdateAt" name:"UpdateAt"`
		ImageName   *string `json:"ImageName" name:"ImageName"`
		RawImageUrl *string `json:"RawImageUrl" name:"RawImageUrl"`
		Description *string `json:"Description" name:"Description"`
		ImageType   *string `json:"ImageType" name:"ImageType"`
		Status      *string `json:"Status" name:"Status"`
		OpType      *string `json:"OpType" name:"OpType"`
	} `json:"Result"`
}

func (r *GetImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListImagesRequest struct {
	*ksyunhttp.BaseRequest
	Image_label *int    `json:"image_label,omitempty" name:"image_label"`
	Status      *string `json:"status,omitempty" name:"status"`
	Page_index  *int    `json:"page_index,omitempty" name:"page_index"`
	Page_size   *int    `json:"page_size,omitempty" name:"page_size"`
}

func (r *ListImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListImagesResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id          *int    `json:"Id" name:"Id"`
		CreateAt    *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt    *int    `json:"UpdateAt" name:"UpdateAt"`
		ImageName   *string `json:"ImageName" name:"ImageName"`
		RawImageUrl *string `json:"RawImageUrl" name:"RawImageUrl"`
		Description *string `json:"Description" name:"Description"`
		ImageType   *string `json:"ImageType" name:"ImageType"`
		Status      *string `json:"Status" name:"Status"`
		OpType      *string `json:"OpType" name:"OpType"`
	} `json:"Result"`
}

func (r *ListImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageRequest struct {
	*ksyunhttp.BaseRequest
	Id *int `json:"id,omitempty" name:"id"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageRequest struct {
	*ksyunhttp.BaseRequest
	Image_name  *string `json:"image_name,omitempty" name:"image_name"`
	Image_url   *string `json:"image_url,omitempty" name:"image_url"`
	Image_type  *string `json:"image_type,omitempty" name:"image_type"`
	Description *string `json:"description,omitempty" name:"description"`
	Mirror_addr *string `json:"mirror_addr,omitempty" name:"mirror_addr"`
	User        *string `json:"user,omitempty" name:"user"`
	Secret      *string `json:"secret,omitempty" name:"secret"`
	Op_type     *string `json:"op_type,omitempty" name:"op_type"`
}

func (r *CreateImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  *string `json:"result" name:"result"`
}

func (r *CreateImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodeMetricsRequest struct {
	*ksyunhttp.BaseRequest
	Ip          *string `json:"ip,omitempty" name:"ip"`
	Show_detail *bool   `json:"show_detail,omitempty" name:"show_detail"`
	Sort_by     *string `json:"sort_by,omitempty" name:"sort_by"`
	Order       *string `json:"order,omitempty" name:"order"`
}

func (r *ListNodeMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodeMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListNodeMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListNodeMetricsResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Ip          *string  `json:"Ip" name:"Ip"`
		GpuCount    *int     `json:"GpuCount" name:"GpuCount"`
		GpuUsed     *int     `json:"GpuUsed" name:"GpuUsed"`
		GpuFree     *int     `json:"GpuFree" name:"GpuFree"`
		CpuSm       *float64 `json:"CpuSm" name:"CpuSm"`
		MemorySm    *float64 `json:"MemorySm" name:"MemorySm"`
		GpuSm       *float64 `json:"GpuSm" name:"GpuSm"`
		GpuMemorySm *float64 `json:"GpuMemorySm" name:"GpuMemorySm"`
		Temperature *float64 `json:"Temperature" name:"Temperature"`
	} `json:"Result"`
	GpuInfo []struct {
		Index       *int     `json:"Index" name:"Index"`
		Sm          *float64 `json:"Sm" name:"Sm"`
		MemorySm    *float64 `json:"MemorySm" name:"MemorySm"`
		Temperature *float64 `json:"Temperature" name:"Temperature"`
	} `json:"GpuInfo"`
}

func (r *ListNodeMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodeMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStorageConfigsRequest struct {
	*ksyunhttp.BaseRequest
	Page_index *int `json:"page_index,omitempty" name:"page_index"`
	Page_size  *int `json:"page_size,omitempty" name:"page_size"`
}

func (r *ListStorageConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStorageConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListStorageConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListStorageConfigsResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id          *int    `json:"Id" name:"Id"`
		CreateAt    *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt    *int    `json:"UpdateAt" name:"UpdateAt"`
		Name        *string `json:"Name" name:"Name"`
		Description *string `json:"Description" name:"Description"`
	} `json:"Result"`
}

func (r *ListStorageConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStorageConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceStoragesRequest struct {
	*ksyunhttp.BaseRequest
	Storage_id        *int `json:"storage_id,omitempty" name:"storage_id"`
	Storage_config_id *int `json:"storage_config_id,omitempty" name:"storage_config_id"`
	Page_index        *int `json:"page_index,omitempty" name:"page_index"`
	Page_size         *int `json:"page_size,omitempty" name:"page_size"`
}

func (r *ListSpaceStoragesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpaceStoragesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSpaceStoragesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceStoragesResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id              *int    `json:"Id" name:"Id"`
		CreateAt        *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt        *int    `json:"UpdateAt" name:"UpdateAt"`
		Name            *string `json:"Name" name:"Name"`
		Description     *string `json:"Description" name:"Description"`
		Path            *string `json:"Path" name:"Path"`
		MaxQuotaByte    *int    `json:"MaxQuotaByte" name:"MaxQuotaByte"`
		CreateUserId    *int    `json:"CreateUserId" name:"CreateUserId"`
		CreateUsername  *string `json:"CreateUsername" name:"CreateUsername"`
		StorageConfigId *int    `json:"StorageConfigId" name:"StorageConfigId"`
	} `json:"Result"`
}

func (r *ListSpaceStoragesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpaceStoragesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkspacesRequest struct {
	*ksyunhttp.BaseRequest
	Name       *string `json:"name,omitempty" name:"name"`
	Page_index *int    `json:"page_index,omitempty" name:"page_index"`
	Page_size  *int    `json:"page_size,omitempty" name:"page_size"`
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
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id          *int    `json:"Id" name:"Id"`
		CreateAt    *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt    *int    `json:"UpdateAt" name:"UpdateAt"`
		Name        *string `json:"Name" name:"Name"`
		Description *string `json:"Description" name:"Description"`
	} `json:"Result"`
}

func (r *ListWorkspacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkspacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGpuTypesRequest struct {
	*ksyunhttp.BaseRequest
	Page_index *int `json:"page_index,omitempty" name:"page_index"`
	Page_size  *int `json:"page_size,omitempty" name:"page_size"`
}

func (r *ListGpuTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGpuTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListGpuTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListGpuTypesResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id           *int    `json:"Id" name:"Id"`
		CreateAt     *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt     *int    `json:"UpdateAt" name:"UpdateAt"`
		Name         *string `json:"Name" name:"Name"`
		Status       *bool   `json:"Status" name:"Status"`
		DeviceMemory *int    `json:"DeviceMemory" name:"DeviceMemory"`
		DeviceType   *string `json:"DeviceType" name:"DeviceType"`
	} `json:"Result"`
}

func (r *ListGpuTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGpuTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSpaceStorageRequest struct {
	*ksyunhttp.BaseRequest
	Type              *string `json:"type,omitempty" name:"type"`
	Max_quota         *int    `json:"max_quota,omitempty" name:"max_quota"`
	Name              *string `json:"name,omitempty" name:"name"`
	Storage_config_id *int    `json:"storage_config_id,omitempty" name:"storage_config_id"`
	Description       *string `json:"description,omitempty" name:"description"`
	Path              *string `json:"path,omitempty" name:"path"`
}

func (r *CreateSpaceStorageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSpaceStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSpaceStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSpaceStorageResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  struct {
		Id              *int    `json:"Id" name:"Id"`
		CreateAt        *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt        *int    `json:"UpdateAt" name:"UpdateAt"`
		Name            *string `json:"Name" name:"Name"`
		Description     *string `json:"Description" name:"Description"`
		Path            *string `json:"Path" name:"Path"`
		MaxQuotaByte    *int    `json:"MaxQuotaByte" name:"MaxQuotaByte"`
		StorageConfigId *int    `json:"StorageConfigId" name:"StorageConfigId"`
	} `json:"Result"`
}

func (r *CreateSpaceStorageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSpaceStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkspaceQuotaRequest struct {
	*ksyunhttp.BaseRequest
	Space_id *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetWorkspaceQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkspaceQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetWorkspaceQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkspaceQuotaResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  struct {
		QuotaLimit struct {
			Exclusive struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Type  *string `json:"type" name:"type"`
					Quota *int    `json:"quota" name:"quota"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"Exclusive"`
			Fixed struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Type  *string `json:"type" name:"type"`
					Quota *int    `json:"quota" name:"quota"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"Fixed"`
			SharedDevelop struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Quota *int    `json:"quota" name:"quota"`
					Type  *string `json:"type" name:"type"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"SharedDevelop"`
			SharedTraining struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Quota *int    `json:"quota" name:"quota"`
					Type  *string `json:"type" name:"type"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"SharedTraining"`
		} `json:"QuotaLimit" name:"QuotaLimit"`
		QuotaRemain struct {
			Exclusive struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Type  *string `json:"type" name:"type"`
					Quota *int    `json:"quota" name:"quota"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"Exclusive"`
			Fixed struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Type  *string `json:"type" name:"type"`
					Quota *int    `json:"quota" name:"quota"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"Fixed"`
			SharedDevelop struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Quota *int    `json:"quota" name:"quota"`
					Type  *string `json:"type" name:"type"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"SharedDevelop"`
			SharedTraining struct {
				CpuQuota  *int `json:"CpuQuota" name:"CpuQuota"`
				GpuQuotas []struct {
					Quota *int    `json:"quota" name:"quota"`
					Type  *string `json:"type" name:"type"`
				} `json:"GpuQuotas" name:"GpuQuotas"`
				MemoryQuota  *int    `json:"MemoryQuota" name:"MemoryQuota"`
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				SpaceId      *int    `json:"SpaceId" name:"SpaceId"`
			} `json:"SharedTraining"`
		} `json:"QuotaRemain" name:"QuotaRemain"`
	} `json:"Result"`
}

func (r *GetWorkspaceQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkspaceQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateJobRequest struct {
	*ksyunhttp.BaseRequest
	Space_id               *int                       `json:"space_id,omitempty" name:"space_id"`
	Name                   *string                    `json:"name,omitempty" name:"name"`
	Priority               *int                       `json:"priority,omitempty" name:"priority"`
	Description            *string                    `json:"description,omitempty" name:"description"`
	Image_type             *int                       `json:"image_type,omitempty" name:"image_type"`
	Image_id               *int                       `json:"image_id,omitempty" name:"image_id"`
	Image_url              *string                    `json:"image_url,omitempty" name:"image_url"`
	Pull_user              *string                    `json:"pull_user,omitempty" name:"pull_user"`
	Pull_password          *string                    `json:"pull_password,omitempty" name:"pull_password"`
	Code_resource          *string                    `json:"code_resource,omitempty" name:"code_resource"`
	Working_folder         *string                    `json:"working_folder,omitempty" name:"working_folder"`
	Git_url                *string                    `json:"git_url,omitempty" name:"git_url"`
	Git_branch             *string                    `json:"git_branch,omitempty" name:"git_branch"`
	Account                *string                    `json:"account,omitempty" name:"account"`
	Access_token           *string                    `json:"access_token,omitempty" name:"access_token"`
	Download_path          *string                    `json:"download_path,omitempty" name:"download_path"`
	Run_command            *string                    `json:"run_command,omitempty" name:"run_command"`
	Training_type          *string                    `json:"training_type,omitempty" name:"training_type"`
	Self_healing           *int                       `json:"self_healing,omitempty" name:"self_healing"`
	Max_self_healing_times *int                       `json:"max_self_healing_times,omitempty" name:"max_self_healing_times"`
	Hanging_duration       *int                       `json:"hanging_duration,omitempty" name:"hanging_duration"`
	Replicas               *int                       `json:"replicas,omitempty" name:"replicas"`
	Device_name            *string                    `json:"device_name,omitempty" name:"device_name"`
	Space_storages         []*CreateJobSpace_storages `json:"space_storages,omitempty" name:"space_storages"`
	Env                    *CreateJobEnv              `json:"env,omitempty" name:"env"`
}

func (r *CreateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateJobResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  struct {
		JobId   *int    `json:"JobId" name:"JobId"`
		TaskId  *int    `json:"TaskId" name:"TaskId"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Result"`
}

func (r *CreateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceStorageBindedSpacesRequest struct {
	*ksyunhttp.BaseRequest
	Storage_id *int `json:"storage_id,omitempty" name:"storage_id"`
}

func (r *ListSpaceStorageBindedSpacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpaceStorageBindedSpacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSpaceStorageBindedSpacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceStorageBindedSpacesResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []*int  `json:"Result" name:"Result"`
}

func (r *ListSpaceStorageBindedSpacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpaceStorageBindedSpacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpaceStorageBindSpaceRequest struct {
	*ksyunhttp.BaseRequest
	Storage_id *int `json:"storage_id,omitempty" name:"storage_id"`
	Space_id   *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *SpaceStorageBindSpaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SpaceStorageBindSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SpaceStorageBindSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SpaceStorageBindSpaceResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  *string `json:"result" name:"result"`
}

func (r *SpaceStorageBindSpaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SpaceStorageBindSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpaceStorageUnbindSpaceRequest struct {
	*ksyunhttp.BaseRequest
	Storage_id *int `json:"storage_id,omitempty" name:"storage_id"`
	Space_id   *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *SpaceStorageUnbindSpaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SpaceStorageUnbindSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SpaceStorageUnbindSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SpaceStorageUnbindSpaceResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  *string `json:"result" name:"result"`
}

func (r *SpaceStorageUnbindSpaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SpaceStorageUnbindSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteJobRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int `json:"job_id,omitempty" name:"job_id"`
	Space_id *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *DeleteJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteJobResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  *string `json:"result" name:"result"`
}

func (r *DeleteJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetJobRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int `json:"job_id,omitempty" name:"job_id"`
	Space_id *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetJobResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  struct {
		Id             *int    `json:"Id" name:"Id"`
		CreateAt       *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt       *int    `json:"UpdateAt" name:"UpdateAt"`
		Name           *string `json:"Name" name:"Name"`
		Priority       *int    `json:"Priority" name:"Priority"`
		Description    *string `json:"Description" name:"Description"`
		Status         *string `json:"Status" name:"Status"`
		StatusMessage  *string `json:"StatusMessage" name:"StatusMessage"`
		SpaceId        *int    `json:"SpaceId" name:"SpaceId"`
		ImageType      *int    `json:"ImageType" name:"ImageType"`
		ImageUrl       *string `json:"ImageUrl" name:"ImageUrl"`
		ImageName      *string `json:"ImageName" name:"ImageName"`
		PullUser       *string `json:"PullUser" name:"PullUser"`
		PullPassword   *string `json:"PullPassword" name:"PullPassword"`
		WorkingFolder  *string `json:"WorkingFolder" name:"WorkingFolder"`
		RunCommand     *string `json:"RunCommand" name:"RunCommand"`
		CreateUserId   *int    `json:"CreateUserId" name:"CreateUserId"`
		CreateUsername *string `json:"CreateUsername" name:"CreateUsername"`
		UpdateUserId   *int    `json:"UpdateUserId" name:"UpdateUserId"`
		TaskId         *int    `json:"TaskId" name:"TaskId"`
		Task           struct {
			Id                  *int    `json:"id" name:"id"`
			CreateAt            *int    `json:"create_at" name:"create_at"`
			UpdateAt            *int    `json:"update_at" name:"update_at"`
			Name                *string `json:"name" name:"name"`
			Description         *string `json:"description" name:"description"`
			TaskType            *string `json:"task_type" name:"task_type"`
			ImageUrl            *string `json:"image_url" name:"image_url"`
			WorkingFolder       *string `json:"working_folder" name:"working_folder"`
			RunCommand          *string `json:"run_command" name:"run_command"`
			SpaceId             *int    `json:"space_id" name:"space_id"`
			JobId               *int    `json:"job_id" name:"job_id"`
			JobName             *string `json:"job_name" name:"job_name"`
			Status              *string `json:"status" name:"status"`
			StatusMessage       *string `json:"status_message" name:"status_message"`
			StartUser           *string `json:"start_user" name:"start_user"`
			StartAt             *int    `json:"start_at" name:"start_at"`
			StopAt              *int    `json:"stop_at" name:"stop_at"`
			Duration            *string `json:"duration" name:"duration"`
			StopUser            *string `json:"stop_user" name:"stop_user"`
			StopWay             *string `json:"stop_way" name:"stop_way"`
			UseGpu              *bool   `json:"use_gpu" name:"use_gpu"`
			CpuPeak             *int    `json:"cpu_peak" name:"cpu_peak"`
			MemPeak             *int    `json:"mem_peak" name:"mem_peak"`
			SelfHealing         *int    `json:"self_healing" name:"self_healing"`
			MaxSelfHealingTimes *int    `json:"max_self_healing_times" name:"max_self_healing_times"`
			HangingDuration     *int    `json:"hanging_duration" name:"hanging_duration"`
			OpSrc               *int    `json:"op_src" name:"op_src"`
			TaskHealth          *string `json:"task_health" name:"task_health"`
			Latest              *int    `json:"latest" name:"latest"`
			PodNameList         *string `json:"pod_name_list" name:"pod_name_list"`
		} `json:"Task" name:"Task"`
		TaskList []struct {
			Id                  *int      `json:"id" name:"id"`
			CreateAt            *int      `json:"create_at" name:"create_at"`
			UpdateAt            *int      `json:"update_at" name:"update_at"`
			Name                *string   `json:"name" name:"name"`
			Description         *string   `json:"description" name:"description"`
			TaskType            *string   `json:"task_type" name:"task_type"`
			ImageUrl            *string   `json:"image_url" name:"image_url"`
			WorkingFolder       *string   `json:"working_folder" name:"working_folder"`
			RunCommand          *string   `json:"run_command" name:"run_command"`
			SpaceId             *int      `json:"space_id" name:"space_id"`
			JobId               *int      `json:"job_id" name:"job_id"`
			JobName             *string   `json:"job_name" name:"job_name"`
			Status              *string   `json:"status" name:"status"`
			StatusMessage       *string   `json:"status_message" name:"status_message"`
			StartUser           *string   `json:"start_user" name:"start_user"`
			StartAt             *int      `json:"start_at" name:"start_at"`
			StopAt              *int      `json:"stop_at" name:"stop_at"`
			Duration            *string   `json:"duration" name:"duration"`
			StopUser            *string   `json:"stop_user" name:"stop_user"`
			StopWay             *string   `json:"stop_way" name:"stop_way"`
			UseGpu              *bool     `json:"use_gpu" name:"use_gpu"`
			CpuPeak             *int      `json:"cpu_peak" name:"cpu_peak"`
			MemPeak             *int      `json:"mem_peak" name:"mem_peak"`
			SelfHealing         *int      `json:"self_healing" name:"self_healing"`
			MaxSelfHealingTimes *int      `json:"max_self_healing_times" name:"max_self_healing_times"`
			HangingDuration     *int      `json:"hanging_duration" name:"hanging_duration"`
			OpSrc               *int      `json:"op_src" name:"op_src"`
			TaskHealth          *string   `json:"task_health" name:"task_health"`
			Latest              *int      `json:"latest" name:"latest"`
			PodNameList         []*string `json:"PodNameList" name:"PodNameList"`
		} `json:"TaskList" name:"TaskList"`
		ResourceType        *string `json:"ResourceType" name:"ResourceType"`
		SpaceName           *string `json:"SpaceName" name:"SpaceName"`
		CpuPeak             *int    `json:"CpuPeak" name:"CpuPeak"`
		MemPeak             *int    `json:"MemPeak" name:"MemPeak"`
		TrainingType        *string `json:"TrainingType" name:"TrainingType"`
		CronEnable          *bool   `json:"CronEnable" name:"CronEnable"`
		IsOutOfQuota        *bool   `json:"IsOutOfQuota" name:"IsOutOfQuota"`
		CodeResource        *string `json:"CodeResource" name:"CodeResource"`
		GitUrl              *string `json:"GitUrl" name:"GitUrl"`
		GitBranch           *string `json:"GitBranch" name:"GitBranch"`
		Account             *string `json:"Account" name:"Account"`
		AccessToken         *string `json:"AccessToken" name:"AccessToken"`
		DownloadPath        *string `json:"DownloadPath" name:"DownloadPath"`
		SelfHealing         *int    `json:"SelfHealing" name:"SelfHealing"`
		MaxSelfHealingTimes *int    `json:"MaxSelfHealingTimes" name:"MaxSelfHealingTimes"`
		HangingDuration     *int    `json:"HangingDuration" name:"HangingDuration"`
		Containers          *string `json:"Containers" name:"Containers"`
	} `json:"Result"`
}

func (r *GetJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetJobEventsRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int `json:"job_id,omitempty" name:"job_id"`
	Task_id  *int `json:"task_id,omitempty" name:"task_id"`
	Space_id *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetJobEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetJobEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetJobEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetJobEventsResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id       *int    `json:"Id" name:"Id"`
		CreateAt *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt *int    `json:"UpdateAt" name:"UpdateAt"`
		TaskId   *int    `json:"TaskId" name:"TaskId"`
		PodName  *string `json:"PodName" name:"PodName"`
		Type     *string `json:"Type" name:"Type"`
		Status   *string `json:"Status" name:"Status"`
		Message  *string `json:"Message" name:"Message"`
		AddAt    *int    `json:"AddAt" name:"AddAt"`
	} `json:"Result"`
}

func (r *GetJobEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetJobEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPodEventsRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int    `json:"job_id,omitempty" name:"job_id"`
	Task_id  *int    `json:"task_id,omitempty" name:"task_id"`
	Pod_name *string `json:"pod_name,omitempty" name:"pod_name"`
	Space_id *int    `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetPodEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPodEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPodEventsResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		Id       *int    `json:"Id" name:"Id"`
		CreateAt *int    `json:"CreateAt" name:"CreateAt"`
		UpdateAt *int    `json:"UpdateAt" name:"UpdateAt"`
		TaskId   *int    `json:"TaskId" name:"TaskId"`
		PodName  *string `json:"PodName" name:"PodName"`
		Type     *string `json:"Type" name:"Type"`
		Status   *string `json:"Status" name:"Status"`
		Message  *string `json:"Message" name:"Message"`
		AddAt    *int    `json:"AddAt" name:"AddAt"`
	} `json:"Result"`
}

func (r *GetPodEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWebTerminalRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int    `json:"job_id,omitempty" name:"job_id"`
	Task_id  *int    `json:"task_id,omitempty" name:"task_id"`
	Pod_name *string `json:"pod_name,omitempty" name:"pod_name"`
	Space_id *int    `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetWebTerminalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWebTerminalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetWebTerminalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetWebTerminalResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  *string `json:"result" name:"result"`
}

func (r *GetWebTerminalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWebTerminalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetJobMetricsRequest struct {
	*ksyunhttp.BaseRequest
	Job_id        *int    `json:"job_id,omitempty" name:"job_id"`
	Task_id       *int    `json:"task_id,omitempty" name:"task_id"`
	Start_time    *string `json:"start_time,omitempty" name:"start_time"`
	End_time      *string `json:"end_time,omitempty" name:"end_time"`
	Time_step     *string `json:"time_step,omitempty" name:"time_step"`
	Metric_type   *string `json:"metric_type,omitempty" name:"metric_type"`
	Metric_detail *int    `json:"metric_detail,omitempty" name:"metric_detail"`
	Space_id      *int    `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetJobMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetJobMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetJobMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetJobMetricsResponse struct {
	*ksyunhttp.BaseResponse
	Code   *int    `json:"code" name:"code"`
	Msg    *string `json:"msg" name:"msg"`
	Result []struct {
		Legend  *string `json:"Legend" name:"Legend"`
		Metrics []struct {
			Timestamp *int `json:"timestamp" name:"timestamp"`
			Value     *int `json:"value" name:"value"`
		} `json:"Metrics" name:"Metrics"`
	} `json:"Result"`
}

func (r *GetJobMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetJobMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListJobsRequest struct {
	*ksyunhttp.BaseRequest
	Page_index *int    `json:"page_index,omitempty" name:"page_index"`
	Page_size  *int    `json:"page_size,omitempty" name:"page_size"`
	Space_id   *int    `json:"space_id,omitempty" name:"space_id"`
	Job_name   *string `json:"job_name,omitempty" name:"job_name"`
	User_name  *string `json:"user_name,omitempty" name:"user_name"`
	Status     *string `json:"status,omitempty" name:"status"`
}

func (r *ListJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListJobsResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  struct {
		Data []struct {
			Id             *int    `json:"id" name:"id"`
			CreateAt       *int    `json:"create_at" name:"create_at"`
			UpdateAt       *int    `json:"update_at" name:"update_at"`
			Name           *string `json:"name" name:"name"`
			Priority       *int    `json:"priority" name:"priority"`
			Description    *string `json:"description" name:"description"`
			Status         *string `json:"status" name:"status"`
			StatusMessage  *string `json:"status_message" name:"status_message"`
			SpaceId        *int    `json:"space_id" name:"space_id"`
			ImageType      *int    `json:"image_type" name:"image_type"`
			ImageUrl       *string `json:"image_url" name:"image_url"`
			ImageName      *string `json:"image_name" name:"image_name"`
			PullUser       *string `json:"pull_user" name:"pull_user"`
			PullPassword   *string `json:"pull_password" name:"pull_password"`
			WorkingFolder  *string `json:"working_folder" name:"working_folder"`
			RunCommand     *string `json:"run_command" name:"run_command"`
			Containers     *string `json:"containers" name:"containers"`
			CreateUserId   *int    `json:"create_user_id" name:"create_user_id"`
			CreateUsername *string `json:"create_username" name:"create_username"`
			UpdateUserId   *int    `json:"update_user_id" name:"update_user_id"`
			TaskId         *int    `json:"task_id" name:"task_id"`
			Task           struct {
				Id                  *int    `json:"Id" name:"Id"`
				CreateAt            *int    `json:"CreateAt" name:"CreateAt"`
				UpdateAt            *int    `json:"UpdateAt" name:"UpdateAt"`
				Name                *string `json:"Name" name:"Name"`
				Description         *string `json:"Description" name:"Description"`
				TaskType            *string `json:"TaskType" name:"TaskType"`
				ImageUrl            *string `json:"ImageUrl" name:"ImageUrl"`
				WorkingFolder       *string `json:"WorkingFolder" name:"WorkingFolder"`
				RunCommand          *string `json:"RunCommand" name:"RunCommand"`
				SpaceId             *int    `json:"SpaceId" name:"SpaceId"`
				JobId               *int    `json:"JobId" name:"JobId"`
				JobName             *string `json:"JobName" name:"JobName"`
				Status              *string `json:"Status" name:"Status"`
				StatusMessage       *string `json:"StatusMessage" name:"StatusMessage"`
				StartUser           *string `json:"StartUser" name:"StartUser"`
				StartAt             *int    `json:"StartAt" name:"StartAt"`
				StopAt              *int    `json:"StopAt" name:"StopAt"`
				Duration            *string `json:"Duration" name:"Duration"`
				StopUser            *string `json:"StopUser" name:"StopUser"`
				StopWay             *string `json:"StopWay" name:"StopWay"`
				UseGpu              *bool   `json:"UseGpu" name:"UseGpu"`
				CpuPeak             *int    `json:"CpuPeak" name:"CpuPeak"`
				MemPeak             *int    `json:"MemPeak" name:"MemPeak"`
				SelfHealing         *int    `json:"SelfHealing" name:"SelfHealing"`
				MaxSelfHealingTimes *int    `json:"MaxSelfHealingTimes" name:"MaxSelfHealingTimes"`
				HangingDuration     *int    `json:"HangingDuration" name:"HangingDuration"`
				OpSrc               *int    `json:"OpSrc" name:"OpSrc"`
				TaskHealth          *string `json:"TaskHealth" name:"TaskHealth"`
				Latest              *int    `json:"Latest" name:"Latest"`
				PodNameList         *string `json:"PodNameList" name:"PodNameList"`
			} `json:"Task"`
			ResourceType        *string `json:"resource_type" name:"resource_type"`
			SpaceName           *string `json:"space_name" name:"space_name"`
			CpuPeak             *int    `json:"cpu_peak" name:"cpu_peak"`
			MemPeak             *int    `json:"mem_peak" name:"mem_peak"`
			TrainingType        *string `json:"training_type" name:"training_type"`
			CronEnable          *bool   `json:"cron_enable" name:"cron_enable"`
			IsOutOfQuota        *bool   `json:"is_out_of_quota" name:"is_out_of_quota"`
			CodeResource        *string `json:"code_resource" name:"code_resource"`
			GitUrl              *string `json:"git_url" name:"git_url"`
			GitBranch           *string `json:"git_branch" name:"git_branch"`
			Account             *string `json:"account" name:"account"`
			AccessToken         *string `json:"access_token" name:"access_token"`
			DownloadPath        *string `json:"download_path" name:"download_path"`
			SelfHealing         *int    `json:"self_healing" name:"self_healing"`
			MaxSelfHealingTimes *int    `json:"max_self_healing_times" name:"max_self_healing_times"`
			HangingDuration     *int    `json:"hanging_duration" name:"hanging_duration"`
		} `json:"Data" name:"Data"`
		PageNum   *int `json:"PageNum" name:"PageNum"`
		PageIndex *int `json:"PageIndex" name:"PageIndex"`
		PageSize  *int `json:"PageSize" name:"PageSize"`
		Total     *int `json:"Total" name:"Total"`
	} `json:"Result"`
}

func (r *ListJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPodLogsRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int    `json:"job_id,omitempty" name:"job_id"`
	Task_id  *int    `json:"task_id,omitempty" name:"task_id"`
	Pod_name *string `json:"pod_name,omitempty" name:"pod_name"`
	Key_word *string `json:"key_word,omitempty" name:"key_word"`
	Tail     *int    `json:"tail,omitempty" name:"tail"`
	Space_id *int    `json:"space_id,omitempty" name:"space_id"`
}

func (r *GetPodLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPodLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPodLogsResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  []struct {
		PodName *string `json:"PodName" name:"PodName"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Result"`
}

func (r *GetPodLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopJobRequest struct {
	*ksyunhttp.BaseRequest
	Job_id   *int `json:"job_id,omitempty" name:"job_id"`
	Space_id *int `json:"space_id,omitempty" name:"space_id"`
}

func (r *StopJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StopJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopJobResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"code" name:"code"`
	Message *string `json:"message" name:"message"`
	Result  *string `json:"result" name:"result"`
}

func (r *StopJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
