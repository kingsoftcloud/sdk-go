package v20100525

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
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

type GetMetricStatisticsResponse struct {
	*ksyunhttp.BaseResponse
	GetMetricStatisticsResult struct {
		Datapoints struct {
			Member []struct {
			} `json:"Member"`
		} `json:"Datapoints" name:"Datapoints"`
		Label *string `json:"Label" name:"Label"`
	} `json:"GetMetricStatisticsResult"`
	ResponseMetadata struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
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

type ListMetricsResponse struct {
	*ksyunhttp.BaseResponse
	ListMetricsResult struct {
		Metrics struct {
			Member []struct {
				InstanceId *string `json:"InstanceId" name:"InstanceId"`
				MetricName *string `json:"MetricName" name:"MetricName"`
				MetricDesc *string `json:"MetricDesc" name:"MetricDesc"`
				Namespace  *string `json:"Namespace" name:"Namespace"`
				Interval   *string `json:"Interval" name:"Interval"`
				Type       *string `json:"Type" name:"Type"`
				Unit       *string `json:"Unit" name:"Unit"`
			} `json:"Member"`
		} `json:"Metrics" name:"Metrics"`
	} `json:"ListMetricsResult"`
	ResponseMetadata struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"ResponseMetadata"`
}

func (r *ListMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
