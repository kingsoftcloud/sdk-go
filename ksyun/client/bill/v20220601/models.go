package v20220601

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
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

func (r *GetMonthConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetMonthConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *GetPostpayDetailConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPostpayDetailConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
