package v20261022

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetModelIdleResourceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetModelIdleResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetModelIdleResourceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string   `json:"RequestId" name:"RequestId"`
	ModelName      *string   `json:"ModelName" name:"ModelName"`
	HasIdle        *bool     `json:"HasIdle" name:"HasIdle"`
	PredictIdleTpm *int64    `json:"PredictIdleTpm" name:"PredictIdleTpm"`
	Notes          []*string `json:"Notes" name:"Notes"`
}

func (r *GetModelIdleResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModelIdleResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
