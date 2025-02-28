package v20100525

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetMetricStatisticsRequest struct {
	*ksyunhttp.BaseRequest
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Aggregate  *string `json:"Aggregate,omitempty" name:"Aggregate"`
	Period     *string `json:"Period,omitempty" name:"Period"`
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

type ListMetricsRequest struct {
	*ksyunhttp.BaseRequest
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	PageIndex  *int    `json:"PageIndex,omitempty" name:"PageIndex"`
	PageSize   *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListMetricsResponse struct {
	*ksyunhttp.BaseResponse
	ListMetricsResult struct {
		Metrics struct {
			Member []struct {
				InstanceId *string `json:"InstanceId"`
				MetricName *string `json:"MetricName"`
				MetricDesc *string `json:"MetricDesc"`
				Namespace  *string `json:"Namespace"`
				Interval   *string `json:"Interval"`
				Type       *string `json:"Type"`
				Unit       *string `json:"Unit"`
			} `json:"Member"`
		} `json:"Metrics"`
	} `json:"ListMetricsResult"`
	ResponseMetadata struct {
		RequestId *string `json:"RequestId"`
	} `json:"ResponseMetadata"`
}

func (r *ListMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
