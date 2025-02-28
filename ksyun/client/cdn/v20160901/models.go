package v20160901

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetDomainPidDimensionUsageDataRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetDomainPidDimensionUsageDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainPidDimensionUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDomainPidDimensionUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
		Time     *string `json:"Time"`
		Value    *int    `json:"Value"`
		Projects []struct {
			project *string `json:"project"`
			Value   *int    `json:"Value"`
		} `json:"Projects"`
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
