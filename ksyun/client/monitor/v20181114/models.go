package v20181114

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetMetricStatisticsBatchMetrics struct {
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type GetMetricStatisticsBatchRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string                            `json:"Namespace,omitempty" name:"Namespace"`
	StartTime *string                            `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string                            `json:"EndTime,omitempty" name:"EndTime"`
	Aggregate []*string                          `json:"Aggregate,omitempty" name:"Aggregate"`
	Period    *int                               `json:"Period,omitempty" name:"Period"`
	Metrics   []*GetMetricStatisticsBatchMetrics `json:"Metrics,omitempty" name:"Metrics"`
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
	GetMetricStatisticsBatchResults struct {
		Datapoints struct {
			Member []struct {
				Timestamp     *string `json:"Timestamp"`
				UnixTimestamp *int    `json:"UnixTimestamp"`
				Average       *string `json:"Average"`
				Max           *string `json:"Max"`
				Min           *string `json:"Min"`
			} `json:"Member"`
		} `json:"Datapoints"`
		Label    *string `json:"Label"`
		Instance *string `json:"Instance"`
	} `json:"GetMetricStatisticsBatchResults"`
	ResponseMetadata struct {
		RequestId *string `json:"RequestId"`
	} `json:"ResponseMetadata"`
	errorMessage []*string `json:"errorMessage" name:"errorMessage"`
}

func (r *GetMetricStatisticsBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricStatisticsBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
