package v20160901

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetRefreshOrPreloadTaskRequest struct {
	*ksyunhttp.BaseRequest
	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *GetRefreshOrPreloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRefreshOrPreloadTaskResponse struct {
	*ksyunhttp.BaseResponse
	GetRefreshOrPreloadTaskResponse *string `json:"GetRefreshOrPreloadTaskResponse" name:"GetRefreshOrPreloadTaskResponse"`
}

func (r *GetRefreshOrPreloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRefreshOrPreloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshCachesRequest struct {
	*ksyunhttp.BaseRequest
	Files *string `json:"Files,omitempty" name:"Files"`
	Dirs  *string `json:"Dirs,omitempty" name:"Dirs"`
}

func (r *RefreshCachesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RefreshCachesResponse struct {
	*ksyunhttp.BaseResponse
	RefreshTaskId *string `json:"RefreshTaskId" name:"RefreshTaskId"`
}

func (r *RefreshCachesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainPidDimensionUsageDataRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetDomainPidDimensionUsageDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDomainPidDimensionUsageDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	Areas     *string `json:"Areas" name:"Areas"`
	Interval  *int    `json:"Interval" name:"Interval"`
	Metric    []struct {
		Time     *string `json:"Time" name:"Time"`
		Value    *int    `json:"Value" name:"Value"`
		Projects []struct {
			Project *string `json:"project" name:"project"`
			Value   *int    `json:"Value" name:"Value"`
		} `json:"Projects" name:"Projects"`
	} `json:"Metric"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	PeakTime  *string `json:"PeakTime" name:"PeakTime"`
}

func (r *GetDomainPidDimensionUsageDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainPidDimensionUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
