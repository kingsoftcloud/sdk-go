package v20200630
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
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

type GetClientRequestDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Metric     *string `json:"Metric" name:"Metric"`
	Interval   *int    `json:"Interval" name:"Interval"`
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
		} `json:"Condition" name:"Condition"`
		Data []struct {
			Time *string  `json:"Time" name:"Time"`
			Flow *float64 `json:"Flow" name:"Flow"`
		} `json:"Data" name:"Data"`
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

type GetServerDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Metric     *string `json:"Metric" name:"Metric"`
	DataType   *string `json:"DataType" name:"DataType"`
	Interval   *int    `json:"Interval" name:"Interval"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Regions    *string `json:"Regions" name:"Regions"`
	Schema     *string `json:"Schema" name:"Schema"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	Datas      []struct {
		Condition struct {
			DataType *string `json:"DataType" name:"DataType"`
		} `json:"Condition" name:"Condition"`
		Data []struct {
			Time *string  `json:"Time" name:"Time"`
			Flow *float64 `json:"Flow" name:"Flow"`
		} `json:"Data" name:"Data"`
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

type GetDomainRankingListDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
	Datas     []struct {
		Domain *string `json:"Domain" name:"Domain"`
		Rank   *int    `json:"Rank" name:"Rank"`
		Flow   *int    `json:"Flow" name:"Flow"`
		FlowProportion *float64 `json:"FlowProportion" name:"FlowProportion"`
		Bw     *int    `json:"Bw" name:"Bw"`
		PeakTime *string `json:"PeakTime" name:"PeakTime"`
		Pv     *int    `json:"Pv" name:"Pv"`
		PvProportion *float64 `json:"PvProportion" name:"PvProportion"`
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

type GetAreaIspDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	Datas     []struct {
		Area *string `json:"Area" name:"Area"`
		Flow *int `json:"Flow" name:"Flow"`
		Pv   *int `json:"Pv" name:"Pv"`
		FlowProportion *float64 `json:"FlowProportion" name:"FlowProportion"`
		PvProportion *float64 `json:"PvProportion" name:"PvProportion"`
		Isps []struct {
			Isp            *string  `json:"Isp" name:"Isp"`
			Flow           *int     `json:"Flow" name:"Flow"`
			Pv             *int     `json:"Pv" name:"Pv"`
			FlowProportion *float64 `json:"FlowProportion" name:"FlowProportion"`
			PvProportion   *float64 `json:"PvProportion" name:"PvProportion"`
		} `json:"Isps" name:"Isps"`
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

type GetTopReferDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	LimitN    *int    `json:"LimitN" name:"LimitN"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
	Datas     []struct {
		Refer *string `json:"Refer" name:"Refer"`
		Rank *int `json:"Rank" name:"Rank"`
		Pv   *int `json:"Pv" name:"Pv"`
		PvProportion *float64 `json:"PvProportion" name:"PvProportion"`
		Flow *int `json:"Flow" name:"Flow"`
		FlowProportion *float64 `json:"FlowProportion" name:"FlowProportion"`
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

type GetTopUrlDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	LimitN    *int    `json:"LimitN" name:"LimitN"`
	SortBy    *string `json:"SortBy" name:"SortBy"`
	Datas     []struct {
		Url  *string `json:"Url" name:"Url"`
		Rank *int    `json:"Rank" name:"Rank"`
		Pv   *int    `json:"Pv" name:"Pv"`
		PvProportion *float64 `json:"PvProportion" name:"PvProportion"`
		Flow *int    `json:"Flow" name:"Flow"`
		FlowProportion *float64 `json:"FlowProportion" name:"FlowProportion"`
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
		} `json:"Condition" name:"Condition"`
		DetailData struct {
			HitFlow            *int     `json:"HitFlow" name:"HitFlow"`
			MissFlow           *int     `json:"MissFlow" name:"MissFlow"`
			MissFlowProportion *float64 `json:"MissFlowProportion" name:"MissFlowProportion"`
			HitFlowProportion  *float64 `json:"HitFlowProportion" name:"HitFlowProportion"`
		} `json:"DetailData" name:"DetailData"`
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

type GetReqHitRateDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Interval   *int    `json:"Interval" name:"Interval"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Datas      []struct {
		Condition struct {
		} `json:"Condition" name:"Condition"`
		DetailData []struct {
			Time               *string  `json:"Time" name:"Time"`
			HitFlow            *int     `json:"HitFlow" name:"HitFlow"`
			MissFlow           *int     `json:"MissFlow" name:"MissFlow"`
			MissFlowProportion *float64 `json:"MissFlowProportion" name:"MissFlowProportion"`
			HitFlowProportion  *float64 `json:"HitFlowProportion" name:"HitFlowProportion"`
		} `json:"DetailData" name:"DetailData"`
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

type GetFlowHitRateDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Interval   *int    `json:"Interval" name:"Interval"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Datas      []struct {
		Condition struct {
		} `json:"Condition" name:"Condition"`
		DetailData []struct {
			Time               *string  `json:"Time" name:"Time"`
			HitFlow            *int     `json:"HitFlow" name:"HitFlow"`
			MissFlow           *int     `json:"MissFlow" name:"MissFlow"`
			MissFlowProportion *float64 `json:"MissFlowProportion" name:"MissFlowProportion"`
			HitFlowProportion  *float64 `json:"HitFlowProportion" name:"HitFlowProportion"`
		} `json:"DetailData" name:"DetailData"`
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

type GetDomainRequestPeriodRatioDataResponse struct {
	*ksyunhttp.BaseResponse
	CurrentPeriodStartTime *string `json:"CurrentPeriodStartTime" name:"CurrentPeriodStartTime"`
	CurrentPeriodEndTime   *string `json:"CurrentPeriodEndTime" name:"CurrentPeriodEndTime"`
	PriorPeriodStartTime   *string `json:"PriorPeriodStartTime" name:"PriorPeriodStartTime"`
	PriorPeriodEndTime     *string `json:"PriorPeriodEndTime" name:"PriorPeriodEndTime"`
	Metric                 *string `json:"Metric" name:"Metric"`
	Interval               *int    `json:"Interval" name:"Interval"`
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
		} `json:"Condition" name:"Condition"`
		Data []struct {
			CurrentPeriodTime  *string  `json:"CurrentPeriodTime" name:"CurrentPeriodTime"`
			PriorPeriodTime    *string  `json:"PriorPeriodTime" name:"PriorPeriodTime"`
			CurrentPeriodValue *int     `json:"CurrentPeriodValue" name:"CurrentPeriodValue"`
			PriorPeriodValue   *int     `json:"PriorPeriodValue" name:"PriorPeriodValue"`
			PeriodRatio        *float64 `json:"PeriodRatio" name:"PeriodRatio"`
		} `json:"Data" name:"Data"`
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

type GetUvDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Interval   *int    `json:"Interval" name:"Interval"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Datas      []struct {
		Time    *string `json:"Time" name:"Time"`
		Uv      *int    `json:"Uv" name:"Uv"`
		Domains []struct {
			Domain *string `json:"Domain" name:"Domain"`
			Uv     *int    `json:"Uv" name:"Uv"`
		} `json:"Domains" name:"Domains"`
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

type GetTopIpDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	LimitN    *int    `json:"LimitN" name:"LimitN"`
	Datas     []struct {
		Ip   *string `json:"Ip" name:"Ip"`
		Pv   *int    `json:"Pv" name:"Pv"`
		Flow *int    `json:"Flow" name:"Flow"`
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
			Domains *string `json:"Domains" name:"Domains"`
		} `json:"Condition" name:"Condition"`
		HttpcodeData []struct {
			CodeType   *string  `json:"CodeType" name:"CodeType"`
			PV         *int     `json:"PV" name:"PV"`
			Proportion *float64 `json:"Proportion" name:"Proportion"`
		} `json:"HttpcodeData" name:"HttpcodeData"`
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

type GetDomainHttpCodeDetailedDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
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
			Domains *string `json:"Domains" name:"Domains"`
		} `json:"Condition" name:"Condition"`
		HttpcodeData []struct {
			CodeType   *string  `json:"CodeType" name:"CodeType"`
			PV         *int     `json:"PV" name:"PV"`
			Proportion *float64 `json:"Proportion" name:"Proportion"`
		} `json:"HttpcodeData" name:"HttpcodeData"`
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

