package v20200630

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetClientRequestDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Metric     *string `json:"Metric,omitempty" name:"Metric"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Areas      *string `json:"Areas,omitempty" name:"Areas"`
	Provinces  *string `json:"Provinces,omitempty" name:"Provinces"`
	Isps       *string `json:"Isps,omitempty" name:"Isps"`
	IpType     *string `json:"IpType,omitempty" name:"IpType"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetClientRequestDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClientRequestDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetClientRequestDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetClientRequestDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Metric     *string `json:"Metric" name:"Metric"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Areas      *string `json:"Areas" name:"Areas"`
	Provinces  *string `json:"Provinces" name:"Provinces"`
	Isps       *string `json:"Isps" name:"Isps"`
	IpType     *string `json:"IpType" name:"IpType"`
	Schema     *string `json:"Schema" name:"Schema"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	Datas      []struct {
		Condition struct {
		} `json:"Condition"`
		Data []struct {
			Time *string `json:"Time"`
		} `json:"Data"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetClientRequestDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClientRequestDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServerDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Metric     *string `json:"Metric,omitempty" name:"Metric"`
	DataType   *string `json:"DataType,omitempty" name:"DataType"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Regions    *string `json:"Regions,omitempty" name:"Regions"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetServerDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServerDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetServerDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetServerDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Metric     *string `json:"Metric" name:"Metric"`
	DataType   *string `json:"DataType" name:"DataType"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Regions    *string `json:"Regions" name:"Regions"`
	Schema     *string `json:"Schema" name:"Schema"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	Datas      []struct {
		Condition struct {
			DataType *string `json:"DataType"`
		} `json:"Condition"`
		Data []struct {
			Time *string `json:"Time"`
		} `json:"Data"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetServerDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServerDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainRankingListDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	SortBy    *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *GetDomainRankingListDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainRankingListDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDomainRankingListDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainRankingListDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
	Datas     []struct {
		Domain   *string `json:"Domain"`
		PeakTime *string `json:"PeakTime"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetDomainRankingListDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainRankingListDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAreaIspDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
}

func (r *GetAreaIspDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAreaIspDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetAreaIspDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAreaIspDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	Datas     []struct {
		Area *string `json:"Area"`
		Isps []struct {
			Isp *string `json:"Isp"`
		} `json:"Isps"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetAreaIspDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAreaIspDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopReferDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
	LimitN    *string `json:"LimitN,omitempty" name:"LimitN"`
	SortBy    *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *GetTopReferDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopReferDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetTopReferDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTopReferDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
	Datas     []struct {
		Refer *string `json:"Refer"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetTopReferDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopReferDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopUrlDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
	LimitN    *string `json:"LimitN,omitempty" name:"LimitN"`
	SortBy    *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *GetTopUrlDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopUrlDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetTopUrlDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTopUrlDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
	Datas     []struct {
		Url *string `json:"Url"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetTopUrlDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopUrlDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRealTimeHitRateDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime   *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *string `json:"EndTime,omitempty" name:"EndTime"`
	HitRatetype *string `json:"HitRatetype,omitempty" name:"HitRatetype"`
	CdnType     *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains     *string `json:"Domains,omitempty" name:"Domains"`
	ResultType  *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetRealTimeHitRateDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRealTimeHitRateDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetRealTimeHitRateDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRealTimeHitRateDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime   *string `json:"StartTime" name:"StartTime"`
	EndTime     *string `json:"EndTime" name:"EndTime"`
	HitRatetype *string `json:"HitRatetype" name:"HitRatetype"`
	CdnType     *string `json:"CdnType" name:"CdnType"`
	Domains     *string `json:"Domains" name:"Domains"`
	ResultType  *string `json:"ResultType" name:"ResultType"`
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Datas       []struct {
		Condition struct {
		} `json:"Condition"`
		DetailData struct {
		} `json:"DetailData"`
	} `json:"Datas"`
}

func (r *GetRealTimeHitRateDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRealTimeHitRateDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReqHitRateDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetReqHitRateDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReqHitRateDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetReqHitRateDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetReqHitRateDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Datas      []struct {
		Condition struct {
		} `json:"Condition"`
		DetailData []struct {
			Time *string `json:"Time"`
		} `json:"DetailData"`
	} `json:"Datas"`
}

func (r *GetReqHitRateDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReqHitRateDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowHitRateDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetFlowHitRateDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowHitRateDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetFlowHitRateDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowHitRateDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Datas      []struct {
		Condition struct {
		} `json:"Condition"`
		DetailData []struct {
			Time *string `json:"Time"`
		} `json:"DetailData"`
	} `json:"Datas"`
}

func (r *GetFlowHitRateDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowHitRateDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainRequestPeriodRatioDataRequest struct {
	*ksyunhttp.BaseRequest
	CurrentPeriodStartTime *string `json:"CurrentPeriodStartTime,omitempty" name:"CurrentPeriodStartTime"`
	CurrentPeriodEndTime   *string `json:"CurrentPeriodEndTime,omitempty" name:"CurrentPeriodEndTime"`
	PriorPeriodStartTime   *string `json:"PriorPeriodStartTime,omitempty" name:"PriorPeriodStartTime"`
	PriorPeriodEndTime     *string `json:"PriorPeriodEndTime,omitempty" name:"PriorPeriodEndTime"`
	Metric                 *string `json:"Metric,omitempty" name:"Metric"`
	CdnType                *string `json:"CdnType,omitempty" name:"CdnType"`
	Interval               *string `json:"Interval,omitempty" name:"Interval"`
	Domains                *string `json:"Domains,omitempty" name:"Domains"`
	Areas                  *string `json:"Areas,omitempty" name:"Areas"`
	Provinces              *string `json:"Provinces,omitempty" name:"Provinces"`
	Isps                   *string `json:"Isps,omitempty" name:"Isps"`
	IpType                 *string `json:"IpType,omitempty" name:"IpType"`
	Schema                 *string `json:"Schema,omitempty" name:"Schema"`
	ResultType             *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetDomainRequestPeriodRatioDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainRequestPeriodRatioDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDomainRequestPeriodRatioDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainRequestPeriodRatioDataResponse struct {
	*ksyunhttp.BaseResponse
	CurrentPeriodStartTime *string `json:"CurrentPeriodStartTime" name:"CurrentPeriodStartTime"`
	CurrentPeriodEndTime   *string `json:"CurrentPeriodEndTime" name:"CurrentPeriodEndTime"`
	PriorPeriodStartTime   *string `json:"PriorPeriodStartTime" name:"PriorPeriodStartTime"`
	PriorPeriodEndTime     *string `json:"PriorPeriodEndTime" name:"PriorPeriodEndTime"`
	Metric                 *string `json:"Metric" name:"Metric"`
	CdnType                *string `json:"CdnType" name:"CdnType"`
	Domains                *string `json:"Domains" name:"Domains"`
	Areas                  *string `json:"Areas" name:"Areas"`
	Provinces              *string `json:"Provinces" name:"Provinces"`
	Isps                   *string `json:"Isps" name:"Isps"`
	IpType                 *string `json:"IpType" name:"IpType"`
	Schema                 *string `json:"Schema" name:"Schema"`
	ResultType             *string `json:"ResultType" name:"ResultType"`
	RequestId              *string `json:"RequestId" name:"RequestId"`
	Datas                  []struct {
		Condition struct {
		} `json:"Condition"`
		Data []struct {
			CurrentPeriodTime *string `json:"CurrentPeriodTime"`
			PriorPeriodTime   *string `json:"PriorPeriodTime"`
		} `json:"Data"`
	} `json:"Datas"`
}

func (r *GetDomainRequestPeriodRatioDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainRequestPeriodRatioDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUvDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetUvDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUvDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUvDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetUvDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Datas      []struct {
		Time    *string `json:"Time"`
		Domains []struct {
			Domain *string `json:"Domain"`
		} `json:"Domains"`
	} `json:"Datas"`
}

func (r *GetUvDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUvDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopIpDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
	LimitN    *string `json:"LimitN,omitempty" name:"LimitN"`
	SortBy    *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *GetTopIpDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopIpDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetTopIpDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTopIpDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Datas     []struct {
		Ip *string `json:"Ip"`
	} `json:"Datas"`
	Domains   *string `json:"Domains" name:"Domains"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
}

func (r *GetTopIpDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopIpDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSrcDomainHttpCodeDetailedDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	CodeType   *string `json:"CodeType,omitempty" name:"CodeType"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetSrcDomainHttpCodeDetailedDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSrcDomainHttpCodeDetailedDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetSrcDomainHttpCodeDetailedDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSrcDomainHttpCodeDetailedDataResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetSrcDomainHttpCodeDetailedDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSrcDomainHttpCodeDetailedDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSrcDomainHttpCodeDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetSrcDomainHttpCodeDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSrcDomainHttpCodeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetSrcDomainHttpCodeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSrcDomainHttpCodeDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Schema     *string `json:"Schema" name:"Schema"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	Datas      []struct {
		Condition struct {
			Domains *string `json:"Domains"`
		} `json:"Condition"`
		HttpcodeData []struct {
			CodeType *string `json:"CodeType"`
		} `json:"HttpcodeData"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetSrcDomainHttpCodeDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSrcDomainHttpCodeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainHttpCodeDetailedDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Areas      *string `json:"Areas,omitempty" name:"Areas"`
	Provinces  *string `json:"Provinces,omitempty" name:"Provinces"`
	Isps       *string `json:"Isps,omitempty" name:"Isps"`
	IpType     *string `json:"IpType,omitempty" name:"IpType"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	CodeType   *string `json:"CodeType,omitempty" name:"CodeType"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetDomainHttpCodeDetailedDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainHttpCodeDetailedDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDomainHttpCodeDetailedDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainHttpCodeDetailedDataResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetDomainHttpCodeDetailedDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainHttpCodeDetailedDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainHttpCodeDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Areas      *string `json:"Areas,omitempty" name:"Areas"`
	Provinces  *string `json:"Provinces,omitempty" name:"Provinces"`
	Isps       *string `json:"Isps,omitempty" name:"Isps"`
	IpType     *string `json:"IpType,omitempty" name:"IpType"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetDomainHttpCodeDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainHttpCodeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDomainHttpCodeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainHttpCodeDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Areas      *string `json:"Areas" name:"Areas"`
	Provinces  *string `json:"Provinces" name:"Provinces"`
	Isps       *string `json:"Isps" name:"Isps"`
	IpType     *string `json:"IpType" name:"IpType"`
	Schema     *string `json:"Schema" name:"Schema"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	Datas      []struct {
		Condition struct {
			Domains *string `json:"Domains"`
		} `json:"Condition"`
		HttpcodeData []struct {
			CodeType *string `json:"CodeType"`
		} `json:"HttpcodeData"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetDomainHttpCodeDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainHttpCodeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEntryRateDataRequest struct {
	*ksyunhttp.BaseRequest
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	Province  *string `json:"Province,omitempty" name:"Province"`
	Isp       *string `json:"Isp,omitempty" name:"Isp"`
}

func (r *GetEntryRateDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEntryRateDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetEntryRateDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetEntryRateDataResponse struct {
	*ksyunhttp.BaseResponse
	EcnEntryRateDataResponse struct {
	} `json:"EcnEntryRateDataResponse"`
}

func (r *GetEntryRateDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEntryRateDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
