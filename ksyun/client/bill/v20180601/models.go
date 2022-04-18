package v20180601
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetMonthBillRequest struct {
    *ksyunhttp.BaseRequest
    Action *string `json:"Action,omitempty" name:"Action"`
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

type GetPostpayDetailBillCSVRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetPostpayDetailBillCSVRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPostpayDetailBillCSVRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPostpayDetailBillCSVRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetPostpayDetailBillCSVResponse struct {
    *ksyunhttp.BaseResponse
}

func (r *GetPostpayDetailBillCSVResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPostpayDetailBillCSVResponse) FromJsonString(s string) error {
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

type getMonthConsumeRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *getMonthConsumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *getMonthConsumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "getMonthConsumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type getMonthConsumeResponse struct {
    *ksyunhttp.BaseResponse
}

func (r *getMonthConsumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *getMonthConsumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type getPostpayDetailConsumeRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *getPostpayDetailConsumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *getPostpayDetailConsumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "getPostpayDetailConsumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type getPostpayDetailConsumeResponse struct {
    *ksyunhttp.BaseResponse
}

func (r *getPostpayDetailConsumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *getPostpayDetailConsumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

