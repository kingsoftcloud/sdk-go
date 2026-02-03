package v20160901

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetRefreshOrPreloadTaskRequest struct {
	*ksyunhttp.BaseRequest
	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *GetRefreshOrPreloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRefreshOrPreloadTaskResponse struct {
	*ksyunhttp.BaseResponse
	GetRefreshOrPreloadTaskResponse *string `json:"GetRefreshOrPreloadTaskResponse" name:"GetRefreshOrPreloadTaskResponse"`
}

func (r *GetRefreshOrPreloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRefreshOrPreloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshCachesRequest struct {
	*ksyunhttp.BaseRequest
	Files *string `json:"Files,omitempty" name:"Files"`
	Dirs  *string `json:"Dirs,omitempty" name:"Dirs"`
}

func (r *RefreshCachesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RefreshCachesResponse struct {
	*ksyunhttp.BaseResponse
	RefreshTaskId *string `json:"RefreshTaskId" name:"RefreshTaskId"`
}

func (r *RefreshCachesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
