package v20200901
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type DeleteTagTags struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ListResourcesTagFilters struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ReplaceResourcesTagsReplaceTags struct {
	ResourceUuids *string `json:"ResourceUuids,omitempty" name:"ResourceUuids"`
	TagIds        *string `json:"TagIds,omitempty" name:"TagIds"`
}


type CreateTagRequest struct {
	*ksyunhttp.BaseRequest
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *CreateTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTagResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	TagId     *int    `json:"TagId" name:"TagId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteTagRequest struct {
	*ksyunhttp.BaseRequest
	Tags []*DeleteTagTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *DeleteTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteTagResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListTagsRequest struct {
	*ksyunhttp.BaseRequest
	Page     *int    `json:"Page,omitempty" name:"Page"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
	Key      *string `json:"Key,omitempty" name:"Key"`
	Value    *string `json:"Value,omitempty" name:"Value"`
}

func (r *ListTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListTagsResponse struct {
	*ksyunhttp.BaseResponse
	Tags []struct {
		Id        *int    `json:"Id" name:"Id"`
		Key       *string `json:"Key" name:"Key"`
		Value     *string `json:"Value" name:"Value"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		CanDelete *int    `json:"CanDelete" name:"CanDelete"`
		IsBillTag *int    `json:"IsBillTag" name:"IsBillTag"`
	} `json:"Tags"`
	Page      *int    `json:"Page" name:"Page"`
	PageSize  *int    `json:"PageSize" name:"PageSize"`
	Total     *int    `json:"Total" name:"Total"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListTagKeysRequest struct {
	*ksyunhttp.BaseRequest
	TagKey   *string `json:"TagKey,omitempty" name:"TagKey"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListTagKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListTagKeysResponse struct {
	*ksyunhttp.BaseResponse
	TagKeys   []*string `json:"TagKeys" name:"TagKeys"`
	Page      *int      `json:"Page" name:"Page"`
	PageSize  *int      `json:"PageSize" name:"PageSize"`
	Total     *int      `json:"Total" name:"Total"`
	RequestId *string   `json:"RequestId" name:"RequestId"`
}

func (r *ListTagKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListTagValuesRequest struct {
	*ksyunhttp.BaseRequest
	TagKeys  *string `json:"TagKeys,omitempty" name:"TagKeys"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListTagValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListTagValuesResponse struct {
	*ksyunhttp.BaseResponse
	TagValues []struct {
		ID      *string `json:"ID" name:"ID"`
		Key     *string `json:"Key" name:"Key"`
		Value   *string `json:"Value" name:"Value"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		BindNum *int    `json:"BindNum" name:"BindNum"`
	} `json:"TagValues"`
	Page      *int    `json:"Page" name:"Page"`
	PageSize  *int    `json:"PageSize" name:"PageSize"`
	Total     *int    `json:"Total" name:"Total"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListTagValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListResourcesRequest struct {
	*ksyunhttp.BaseRequest
	ResourceType  *string                    `json:"ResourceType,omitempty" name:"ResourceType"`
	ProjectIds    *string                    `json:"ProjectIds,omitempty" name:"ProjectIds"`
	RegionCodes   *string                    `json:"RegionCodes,omitempty" name:"RegionCodes"`
	ResourceUuids *string                    `json:"ResourceUuids,omitempty" name:"ResourceUuids"`
	ResourceName  *string                    `json:"ResourceName,omitempty" name:"ResourceName"`
	TagFilters    []*ListResourcesTagFilters `json:"TagFilters,omitempty" name:"TagFilters"`
	Page          *int                       `json:"Page,omitempty" name:"Page"`
	PageSize      *int                       `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListResourcesResponse struct {
	*ksyunhttp.BaseResponse
	Resources *string `json:"Resources" name:"Resources"`
	Page      *int    `json:"Page" name:"Page"`
	PageSize  *int    `json:"PageSize" name:"PageSize"`
	Total     *int    `json:"Total" name:"Total"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListTagsByResourceIdsRequest struct {
	*ksyunhttp.BaseRequest
	ResourceType  *string `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceUuids *string `json:"ResourceUuids,omitempty" name:"ResourceUuids"`
}

func (r *ListTagsByResourceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListTagsByResourceIdsResponse struct {
	*ksyunhttp.BaseResponse
	Tags []struct {
		ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
		TagId  *int    `json:"TagId" name:"TagId"`
		TagKey *string `json:"TagKey" name:"TagKey"`
		TagValue *string `json:"TagValue" name:"TagValue"`
	} `json:"Tags"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListTagsByResourceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagsByResourceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ReplaceResourcesTagsRequest struct {
	*ksyunhttp.BaseRequest
	ResourceType *string                            `json:"ResourceType,omitempty" name:"ResourceType"`
	ReplaceTags  []*ReplaceResourcesTagsReplaceTags `json:"ReplaceTags,omitempty" name:"ReplaceTags"`
}

func (r *ReplaceResourcesTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReplaceResourcesTagsResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ReplaceResourcesTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceResourcesTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DetachResourceTagsRequest struct {
	*ksyunhttp.BaseRequest
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceUuid *string `json:"ResourceUuid,omitempty" name:"ResourceUuid"`
	TagIds       *string `json:"TagIds,omitempty" name:"TagIds"`
}

func (r *DetachResourceTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DetachResourceTagsResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DetachResourceTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateTagAndAttachResourceRequest struct {
	*ksyunhttp.BaseRequest
	TagKey       *string `json:"TagKey,omitempty" name:"TagKey"`
	TagValue     *string `json:"TagValue,omitempty" name:"TagValue"`
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceUuid *string `json:"ResourceUuid,omitempty" name:"ResourceUuid"`
}

func (r *CreateTagAndAttachResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTagAndAttachResourceResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateTagAndAttachResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTagAndAttachResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

