package v20100525
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetMetricStatisticsRequest struct {
    *ksyunhttp.BaseRequest
    Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
    InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
    MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
    StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
    EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
    Aggregate *string `json:"Aggregate,omitempty" name:"Aggregate"`
    Period *int `json:"Period,omitempty" name:"Period"`
}

func (r *GetMetricStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMetricStatisticsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetMetricStatisticsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetMetricStatisticsResponse struct {
    *ksyunhttp.BaseResponse
	GetMetricStatisticsResult struct {
		Datapoints struct {
			Member []struct {
			} `json:"Member"`
		} `json:"Datapoints"`
		Label *string `json:"Label"`
	} `json:"GetMetricStatisticsResult"`
	ResponseMetadata struct {
		RequestId *string `json:"RequestId"`
	} `json:"ResponseMetadata"`
}

func (r *GetMetricStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMetricStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

