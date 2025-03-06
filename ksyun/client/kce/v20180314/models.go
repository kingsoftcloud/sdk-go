package v20180314

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeRepoNamespaceFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeRepositoryFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribePublicRepositoryFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeTagFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateRepoNamespaceRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	Public    *string `json:"Public,omitempty" name:"Public"`
}

func (r *CreateRepoNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepoNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRepoNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRepoNamespaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateRepoNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepoNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepoNamespaceRequest struct {
	*ksyunhttp.BaseRequest
	Namespace  *string                        `json:"Namespace,omitempty" name:"Namespace"`
	MaxResults *string                        `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *string                        `json:"Marker,omitempty" name:"Marker"`
	Filter     []*DescribeRepoNamespaceFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRepoNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepoNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRepoNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepoNamespaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	Marker       *int    `json:"Marker" name:"Marker"`
	MaxResults   *int    `json:"MaxResults" name:"MaxResults"`
	TotalCount   *int    `json:"TotalCount" name:"TotalCount"`
	NamespaceSet []struct {
		Namespace  *string `json:"Namespace" name:"Namespace"`
		Public     *bool   `json:"Public" name:"Public"`
		RepoCount  *int    `json:"RepoCount" name:"RepoCount"`
		ChartCount *int    `json:"ChartCount" name:"ChartCount"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
	} `json:"NamespaceSet"`
}

func (r *DescribeRepoNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepoNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRepoNamespaceTypeRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	Public    *bool   `json:"Public,omitempty" name:"Public"`
}

func (r *ModifyRepoNamespaceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepoNamespaceTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyRepoNamespaceTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRepoNamespaceTypeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyRepoNamespaceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepoNamespaceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RepoNamespaceExistRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *RepoNamespaceExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RepoNamespaceExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RepoNamespaceExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RepoNamespaceExistResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Exist     *bool   `json:"Exist" name:"Exist"`
}

func (r *RepoNamespaceExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RepoNamespaceExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRepositoryRequest struct {
	*ksyunhttp.BaseRequest
	RepoName    *string `json:"RepoName,omitempty" name:"RepoName"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRepositoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRepositoryRequest struct {
	*ksyunhttp.BaseRequest
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRepositoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryRequest struct {
	*ksyunhttp.BaseRequest
	RepoName   []*string                   `json:"RepoName,omitempty" name:"RepoName"`
	MaxResults *int                        `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int                        `json:"Marker,omitempty" name:"Marker"`
	Filter     []*DescribeRepositoryFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	RepoSet    []struct {
		RepoName    *string `json:"RepoName" name:"RepoName"`
		Public      *bool   `json:"Public" name:"Public"`
		FavorCount  *int    `json:"FavorCount" name:"FavorCount"`
		PullCount   *int    `json:"PullCount" name:"PullCount"`
		Description *string `json:"Description" name:"Description"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
	} `json:"RepoSet"`
}

func (r *DescribeRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicRepositoryRequest struct {
	*ksyunhttp.BaseRequest
	RepoName   []*string                         `json:"RepoName,omitempty" name:"RepoName"`
	MaxResults *int                              `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int                              `json:"Marker,omitempty" name:"Marker"`
	Filter     []*DescribePublicRepositoryFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribePublicRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePublicRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicRepositoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	RepoSet    []struct {
		RepoName    *string `json:"RepoName" name:"RepoName"`
		Public      *bool   `json:"Public" name:"Public"`
		FavorCount  *int    `json:"FavorCount" name:"FavorCount"`
		PullCount   *int    `json:"PullCount" name:"PullCount"`
		Description *string `json:"Description" name:"Description"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
	} `json:"RepoSet"`
}

func (r *DescribePublicRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRepoDescRequest struct {
	*ksyunhttp.BaseRequest
	RepoName    *string `json:"RepoName,omitempty" name:"RepoName"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateRepoDescRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRepoDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateRepoDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRepoDescResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateRepoDescResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRepoDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagRequest struct {
	*ksyunhttp.BaseRequest
	RepoName   *string              `json:"RepoName,omitempty" name:"RepoName"`
	MaxResults *string              `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *string              `json:"Marker,omitempty" name:"Marker"`
	Filter     []*DescribeTagFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	TagSet     []struct {
		RepoName      *string `json:"RepoName" name:"RepoName"`
		TagName       *string `json:"TagName" name:"TagName"`
		ImageId       *string `json:"ImageId" name:"ImageId"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		SizeByte      *int    `json:"SizeByte" name:"SizeByte"`
		Author        *string `json:"Author" name:"Author"`
		DockerVersion *string `json:"DockerVersion" name:"DockerVersion"`
		Architecture  *string `json:"Architecture" name:"Architecture"`
		Os            *string `json:"Os" name:"Os"`
	} `json:"TagSet"`
}

func (r *DescribeTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsRequest struct {
	*ksyunhttp.BaseRequest
	RepoName *string   `json:"RepoName,omitempty" name:"RepoName"`
	Tag      []*string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DeleteTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddFavorRequest struct {
	*ksyunhttp.BaseRequest
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *AddFavorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddFavorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddFavorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddFavorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddFavorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddFavorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFavorRequest struct {
	*ksyunhttp.BaseRequest
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteFavorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteFavorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFavorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteFavorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFavorRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int    `json:"Marker,omitempty" name:"Marker"`
	Keyword    *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *GetFavorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFavorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetFavorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetFavorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	FavorSet   []struct {
		RepoName   *string `json:"RepoName" name:"RepoName"`
		RepoType   *string `json:"RepoType" name:"RepoType"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
	} `json:"FavorSet"`
}

func (r *GetFavorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFavorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterRepositoryAccountRequest struct {
	*ksyunhttp.BaseRequest
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *RegisterRepositoryAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterRepositoryAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RegisterRepositoryAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterRepositoryAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RegisterRepositoryAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterRepositoryAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*ksyunhttp.BaseRequest
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRepoNamespaceRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteRepoNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepoNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRepoNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRepoNamespaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRepoNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepoNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
