package v20240501

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateJobspace_storages struct {
	Space_storage_id       *int    `json:"space_storage_id,omitempty" name:"space_storage_id"`
	Space_storage_location *string `json:"space_storage_location,omitempty" name:"space_storage_location"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  struct {
		Id            *int    `json:"Id"`
		Create_at     *int    `json:"Create_at"`
		Update_at     *int    `json:"Update_at"`
		Image_name    *string `json:"Image_name"`
		Raw_image_url *string `json:"Raw_image_url"`
		Description   *string `json:"Description"`
		Image_type    *string `json:"Image_type"`
		Status        *string `json:"Status"`
		Message       *string `json:"Message"`
		Op_type       *string `json:"Op_type"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id            *int    `json:"Id"`
		Create_at     *int    `json:"Create_at"`
		Update_at     *int    `json:"Update_at"`
		Image_name    *string `json:"Image_name"`
		Raw_image_url *string `json:"Raw_image_url"`
		Description   *string `json:"Description"`
		Image_type    *string `json:"Image_type"`
		Status        *string `json:"Status"`
		Op_type       *string `json:"Op_type"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  *string `json:"result" name:"result"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Ip        *string `json:"Ip"`
		Gpu_count *int    `json:"Gpu_count"`
		Gpu_used  *int    `json:"Gpu_used"`
		Gpu_free  *int    `json:"Gpu_free"`
	} `json:"Result"`
	Gpu_infos []struct {
		Index *int `json:"Index"`
	} `json:"Gpu_infos"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id          *int    `json:"Id"`
		Create_at   *int    `json:"Create_at"`
		Update_at   *int    `json:"Update_at"`
		Name        *string `json:"Name"`
		Description *string `json:"Description"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id                *int    `json:"Id"`
		Create_at         *int    `json:"Create_at"`
		Update_at         *int    `json:"Update_at"`
		Name              *string `json:"Name"`
		Description       *string `json:"Description"`
		Path              *string `json:"Path"`
		Max_quota_byte    *int    `json:"Max_quota_byte"`
		Create_user_id    *int    `json:"Create_user_id"`
		Create_username   *string `json:"Create_username"`
		Storage_config_id *int    `json:"Storage_config_id"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id          *int    `json:"Id"`
		Create_at   *int    `json:"Create_at"`
		Update_at   *int    `json:"Update_at"`
		Name        *string `json:"Name"`
		Description *string `json:"Description"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id            *int    `json:"Id"`
		Create_at     *int    `json:"Create_at"`
		Update_at     *int    `json:"Update_at"`
		Name          *string `json:"Name"`
		Status        *bool   `json:"Status"`
		Device_memory *int    `json:"Device_memory"`
		Device_type   *string `json:"Device_type"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  struct {
		Id                *int    `json:"Id"`
		Create_at         *int    `json:"Create_at"`
		Update_at         *int    `json:"Update_at"`
		Name              *string `json:"Name"`
		Description       *string `json:"Description"`
		Path              *string `json:"Path"`
		Max_quota_byte    *int    `json:"Max_quota_byte"`
		Create_user_id    *int    `json:"Create_user_id"`
		Create_username   *string `json:"Create_username"`
		Storage_config_id *int    `json:"Storage_config_id"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  struct {
		Quota_limit struct {
			Exclusive struct {
				Cpu_quota  *int `json:"Cpu_quota"`
				Gpu_quotas []struct {
					type *string `json:"type"`
					quota *int `json:"quota"`
				} `json:"Gpu_quotas"`
				Memory_quota  *int    `json:"Memory_quota"`
				Resource_type *string `json:"Resource_type"`
				Space_id      *int    `json:"Space_id"`
			} `json:"Exclusive"`
			Fixed struct {
				Cpu_quota  *int `json:"Cpu_quota"`
				Gpu_quotas []struct {
					type *string `json:"type"`
					quota *int `json:"quota"`
				} `json:"Gpu_quotas"`
				Memory_quota  *int    `json:"Memory_quota"`
				Resource_type *string `json:"Resource_type"`
				Space_id      *int    `json:"Space_id"`
			} `json:"Fixed"`
			Shared-develop struct {
			Cpu_quota *int `json:"Cpu_quota"`
			Gpu_quotas []struct {
			quota *int `json:"quota"`
			type *string `json:"type"`
		} `json:"Gpu_quotas"`
			Memory_quota *int `json:"Memory_quota"`
			Resource_type *string `json:"Resource_type"`
			Space_id *int `json:"Space_id"`
		} `json:"Shared-develop"`
			Shared-training struct {
			Cpu_quota *int `json:"Cpu_quota"`
			Gpu_quotas []struct {
			quota *int `json:"quota"`
			type *string `json:"type"`
		} `json:"Gpu_quotas"`
			Memory_quota *int `json:"Memory_quota"`
			Resource_type *string `json:"Resource_type"`
			Space_id *int `json:"Space_id"`
		} `json:"Shared-training"`
		} `json:"Quota_limit"`
		Quota_remain struct {
			Exclusive struct {
				Cpu_quota  *int `json:"Cpu_quota"`
				Gpu_quotas []struct {
					type *string `json:"type"`
					quota *int `json:"quota"`
				} `json:"Gpu_quotas"`
				Memory_quota  *int    `json:"Memory_quota"`
				Resource_type *string `json:"Resource_type"`
				Space_id      *int    `json:"Space_id"`
			} `json:"Exclusive"`
			Fixed struct {
				Cpu_quota  *int `json:"Cpu_quota"`
				Gpu_quotas []struct {
					type *string `json:"type"`
					quota *int `json:"quota"`
				} `json:"Gpu_quotas"`
				Memory_quota  *int    `json:"Memory_quota"`
				Resource_type *string `json:"Resource_type"`
				Space_id      *int    `json:"Space_id"`
			} `json:"Fixed"`
			Shared-develop struct {
			Cpu_quota *int `json:"Cpu_quota"`
			Gpu_quotas []struct {
			quota *int `json:"quota"`
			type *string `json:"type"`
		} `json:"Gpu_quotas"`
			Memory_quota *int `json:"Memory_quota"`
			Resource_type *string `json:"Resource_type"`
			Space_id *int `json:"Space_id"`
		} `json:"Shared-develop"`
			Shared-training struct {
			Cpu_quota *int `json:"Cpu_quota"`
			Gpu_quotas []struct {
			quota *int `json:"quota"`
			type *string `json:"type"`
		} `json:"Gpu_quotas"`
			Memory_quota *int `json:"Memory_quota"`
			Resource_type *string `json:"Resource_type"`
			Space_id *int `json:"Space_id"`
		} `json:"Shared-training"`
		} `json:"Quota_remain"`
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
	Imgae_url              *string                    `json:"imgae_url,omitempty" name:"imgae_url"`
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
	Space_storages         []*CreateJobspace_storages `json:"space_storages,omitempty" name:"space_storages"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  struct {
		Job_id  *int    `json:"Job_id"`
		Task_id *int    `json:"Task_id"`
		Message *string `json:"Message"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  []*int  `json:"result" name:"result"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  *string `json:"result" name:"result"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  *string `json:"result" name:"result"`
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
	Job_id *int `json:"job_id,omitempty" name:"job_id"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  *string `json:"result" name:"result"`
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
	Job_id *int `json:"job_id,omitempty" name:"job_id"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  struct {
		Id              *int    `json:"Id"`
		Create_at       *int    `json:"Create_at"`
		Update_at       *int    `json:"Update_at"`
		Name            *string `json:"Name"`
		Priority        *int    `json:"Priority"`
		Description     *string `json:"Description"`
		Status          *string `json:"Status"`
		Status_message  *string `json:"Status_message"`
		Space_id        *int    `json:"Space_id"`
		Image_type      *int    `json:"Image_type"`
		Image_url       *string `json:"Image_url"`
		Image_name      *string `json:"Image_name"`
		Pull_user       *string `json:"Pull_user"`
		Pull_password   *string `json:"Pull_password"`
		Working_folder  *string `json:"Working_folder"`
		Run_command     *string `json:"Run_command"`
		Create_user_id  *int    `json:"Create_user_id"`
		Create_username *string `json:"Create_username"`
		Update_user_id  *int    `json:"Update_user_id"`
		Task_id         *int    `json:"Task_id"`
		Task            struct {
			id                     *int    `json:"id"`
			create_at              *int    `json:"create_at"`
			update_at              *int    `json:"update_at"`
			name                   *string `json:"name"`
			description            *string `json:"description"`
			task_type              *string `json:"task_type"`
			image_url              *string `json:"image_url"`
			working_folder         *string `json:"working_folder"`
			run_command            *string `json:"run_command"`
			space_id               *int    `json:"space_id"`
			job_id                 *int    `json:"job_id"`
			job_name               *string `json:"job_name"`
			status                 *string `json:"status"`
			status_message         *string `json:"status_message"`
			start_user             *string `json:"start_user"`
			start_at               *int    `json:"start_at"`
			stop_at                *int    `json:"stop_at"`
			duration               *string `json:"duration"`
			stop_user              *string `json:"stop_user"`
			stop_way               *string `json:"stop_way"`
			use_gpu                *bool   `json:"use_gpu"`
			cpu_peak               *int    `json:"cpu_peak"`
			mem_peak               *int    `json:"mem_peak"`
			self_healing           *int    `json:"self_healing"`
			max_self_healing_times *int    `json:"max_self_healing_times"`
			hanging_duration       *int    `json:"hanging_duration"`
			op_src                 *int    `json:"op_src"`
			task_health            *string `json:"task_health"`
			latest                 *int    `json:"latest"`
			pod_name_list          *string `json:"pod_name_list"`
		} `json:"Task"`
		Task_list []struct {
			id                     *int    `json:"id"`
			create_at              *int    `json:"create_at"`
			update_at              *int    `json:"update_at"`
			name                   *string `json:"name"`
			description            *string `json:"description"`
			task_type              *string `json:"task_type"`
			image_url              *string `json:"image_url"`
			working_folder         *string `json:"working_folder"`
			run_command            *string `json:"run_command"`
			space_id               *int    `json:"space_id"`
			job_id                 *int    `json:"job_id"`
			job_name               *string `json:"job_name"`
			status                 *string `json:"status"`
			status_message         *string `json:"status_message"`
			start_user             *string `json:"start_user"`
			start_at               *int    `json:"start_at"`
			stop_at                *int    `json:"stop_at"`
			duration               *string `json:"duration"`
			stop_user              *string `json:"stop_user"`
			stop_way               *string `json:"stop_way"`
			use_gpu                *bool   `json:"use_gpu"`
			cpu_peak               *int    `json:"cpu_peak"`
			mem_peak               *int    `json:"mem_peak"`
			self_healing           *int    `json:"self_healing"`
			max_self_healing_times *int    `json:"max_self_healing_times"`
			hanging_duration       *int    `json:"hanging_duration"`
			op_src                 *int    `json:"op_src"`
			task_health            *string `json:"task_health"`
			latest                 *int    `json:"latest"`
			Pod_name_list          []struct {
			} `json:"Pod_name_list"`
		} `json:"Task_list"`
		Resource_type          *string `json:"Resource_type"`
		Space_name             *string `json:"Space_name"`
		Cpu_peak               *int    `json:"Cpu_peak"`
		Mem_peak               *int    `json:"Mem_peak"`
		Training_type          *string `json:"Training_type"`
		Cron_enable            *bool   `json:"Cron_enable"`
		Is_out_of_quota        *bool   `json:"Is_out_of_quota"`
		Code_resource          *string `json:"Code_resource"`
		Git_url                *string `json:"Git_url"`
		Git_branch             *string `json:"Git_branch"`
		Account                *string `json:"Account"`
		Access_token           *string `json:"Access_token"`
		Download_path          *string `json:"Download_path"`
		Self_healing           *int    `json:"Self_healing"`
		Max_self_healing_times *int    `json:"Max_self_healing_times"`
		Hanging_duration       *int    `json:"Hanging_duration"`
		Containers             *string `json:"Containers"`
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
	Job_id  *int `json:"job_id,omitempty" name:"job_id"`
	Task_id *int `json:"task_id,omitempty" name:"task_id"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id        *int    `json:"Id"`
		Create_at *int    `json:"Create_at"`
		Update_at *int    `json:"Update_at"`
		Task_id   *int    `json:"Task_id"`
		Pod_name  *string `json:"Pod_name"`
		Type      *string `json:"Type"`
		Status    *string `json:"Status"`
		Message   *string `json:"Message"`
		Add_at    *int    `json:"Add_at"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Id        *int    `json:"Id"`
		Create_at *int    `json:"Create_at"`
		Update_at *int    `json:"Update_at"`
		Task_id   *int    `json:"Task_id"`
		Pod_name  *string `json:"Pod_name"`
		Type      *string `json:"Type"`
		Status    *string `json:"Status"`
		Message   *string `json:"Message"`
		Add_at    *int    `json:"Add_at"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  *string `json:"result" name:"result"`
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
	code   *int    `json:"code" name:"code"`
	msg    *string `json:"msg" name:"msg"`
	Result []struct {
		Legend  *string `json:"Legend"`
		Metrics []struct {
			timestamp *int `json:"timestamp"`
			value     *int `json:"value"`
		} `json:"Metrics"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  struct {
		Data []struct {
			id              *int    `json:"id"`
			create_at       *int    `json:"create_at"`
			update_at       *int    `json:"update_at"`
			name            *string `json:"name"`
			priority        *int    `json:"priority"`
			description     *string `json:"description"`
			status          *string `json:"status"`
			status_message  *string `json:"status_message"`
			space_id        *int    `json:"space_id"`
			image_type      *int    `json:"image_type"`
			image_url       *string `json:"image_url"`
			image_name      *string `json:"image_name"`
			pull_user       *string `json:"pull_user"`
			pull_password   *string `json:"pull_password"`
			working_folder  *string `json:"working_folder"`
			run_command     *string `json:"run_command"`
			containers      *string `json:"containers"`
			create_user_id  *int    `json:"create_user_id"`
			create_username *string `json:"create_username"`
			update_user_id  *int    `json:"update_user_id"`
			task_id         *int    `json:"task_id"`
			Task            struct {
				Id                     *int    `json:"Id"`
				Create_at              *int    `json:"Create_at"`
				Update_at              *int    `json:"Update_at"`
				Name                   *string `json:"Name"`
				Description            *string `json:"Description"`
				Task_type              *string `json:"Task_type"`
				Image_url              *string `json:"Image_url"`
				Working_folder         *string `json:"Working_folder"`
				Run_command            *string `json:"Run_command"`
				Space_id               *int    `json:"Space_id"`
				Job_id                 *int    `json:"Job_id"`
				Job_name               *string `json:"Job_name"`
				Status                 *string `json:"Status"`
				Status_message         *string `json:"Status_message"`
				Start_user             *string `json:"Start_user"`
				Start_at               *int    `json:"Start_at"`
				Stop_at                *int    `json:"Stop_at"`
				Duration               *string `json:"Duration"`
				Stop_user              *string `json:"Stop_user"`
				Stop_way               *string `json:"Stop_way"`
				Use_gpu                *bool   `json:"Use_gpu"`
				Cpu_peak               *int    `json:"Cpu_peak"`
				Mem_peak               *int    `json:"Mem_peak"`
				Self_healing           *int    `json:"Self_healing"`
				Max_self_healing_times *int    `json:"Max_self_healing_times"`
				Hanging_duration       *int    `json:"Hanging_duration"`
				Op_src                 *int    `json:"Op_src"`
				Task_health            *string `json:"Task_health"`
				Latest                 *int    `json:"Latest"`
				Pod_name_list          *string `json:"Pod_name_list"`
			} `json:"Task"`
			resource_type          *string `json:"resource_type"`
			space_name             *string `json:"space_name"`
			cpu_peak               *int    `json:"cpu_peak"`
			mem_peak               *int    `json:"mem_peak"`
			training_type          *string `json:"training_type"`
			cron_enable            *bool   `json:"cron_enable"`
			is_out_of_quota        *bool   `json:"is_out_of_quota"`
			code_resource          *string `json:"code_resource"`
			git_url                *string `json:"git_url"`
			git_branch             *string `json:"git_branch"`
			account                *string `json:"account"`
			access_token           *string `json:"access_token"`
			download_path          *string `json:"download_path"`
			self_healing           *int    `json:"self_healing"`
			max_self_healing_times *int    `json:"max_self_healing_times"`
			hanging_duration       *int    `json:"hanging_duration"`
		} `json:"Data"`
		Page_num   *int `json:"Page_num"`
		Page_index *int `json:"Page_index"`
		Page_size  *int `json:"Page_size"`
		Total      *int `json:"Total"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	Result  []struct {
		Pod_name *string `json:"Pod_name"`
		Message  *string `json:"Message"`
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
	Job_id *int `json:"job_id,omitempty" name:"job_id"`
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
	code    *int    `json:"code" name:"code"`
	message *string `json:"message" name:"message"`
	result  *string `json:"result" name:"result"`
}

func (r *StopJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
