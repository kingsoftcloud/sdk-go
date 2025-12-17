package v20180601

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetMonthBillRequest struct {
	*ksyunhttp.BaseRequest
	BillStartMonth *string `json:"BillStartMonth,omitempty" name:"BillStartMonth"`
	BillEndMonth   *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
}

func (r *GetMonthBillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMonthBillResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetMonthBillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonthBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPostpayDetailBillRequest struct {
	*ksyunhttp.BaseRequest
	BillStartMonth *string `json:"BillStartMonth,omitempty" name:"BillStartMonth"`
	BillEndMonth   *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
	ProductCode    *string `json:"ProductCode,omitempty" name:"ProductCode"`
	ProjectId      *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *GetPostpayDetailBillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetPostpayDetailBillResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetPostpayDetailBillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPostpayDetailBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductCodeRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetProductCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetProductCodeResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetProductCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
