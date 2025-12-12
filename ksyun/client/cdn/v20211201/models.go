package v20211201

import (
	"encoding/json"

	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type GetRefreshOrPreloadTaskRequest struct {
	*ksyunhttp.BaseRequest
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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

