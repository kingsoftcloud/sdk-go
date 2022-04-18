package v20211209
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeCostBillRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeCostBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostBillRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCostBillRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostBillResponse struct {
    *ksyunhttp.BaseResponse
    status *int `json:"status" name:"status"`
}

func (r *DescribeCostBillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostBillResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

