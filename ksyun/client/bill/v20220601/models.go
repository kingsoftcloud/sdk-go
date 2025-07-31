package v20220601

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetMonthConsumeRequest struct {
	*ksyunhttp.BaseRequest
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
}

func (r *GetMonthConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMonthConsumeResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetMonthConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonthConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPostpayDetailConsumeRequest struct {
	*ksyunhttp.BaseRequest
	BillMonth   *string `json:"BillMonth,omitempty" name:"BillMonth"`
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	ProjectId   *string `json:"ProjectId,omitempty" name:"ProjectId"`
	PageNo      *int    `json:"PageNo,omitempty" name:"PageNo"`
	PageSize    *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetPostpayDetailConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetPostpayDetailConsumeResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetPostpayDetailConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPostpayDetailConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
