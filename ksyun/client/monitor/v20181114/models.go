package v20181114
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetMetricStatisticsBatchRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetMetricStatisticsBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMetricStatisticsBatchRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetMetricStatisticsBatchRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetMetricStatisticsBatchResponse struct {
    *ksyunhttp.BaseResponse
    status *int `json:"status" name:"status"`
}

func (r *GetMetricStatisticsBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMetricStatisticsBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

