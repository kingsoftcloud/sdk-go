package v20180601
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetMonthBillRequest struct {
    *ksyunhttp.BaseRequest
    BillStartMonth *string `json:"BillStartMonth,omitempty" name:"BillStartMonth"`
    BillEndMonth *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
}

func (r *GetMonthBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonthBillRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetMonthBillRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    BillEndMonth *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
    ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *GetPostpayDetailBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPostpayDetailBillRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPostpayDetailBillRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetProductCodeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetProductCodeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

