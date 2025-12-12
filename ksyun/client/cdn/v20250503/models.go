package v20250503

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetDomainLogsRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *GetDomainLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDomainLogsResponse struct {
	*ksyunhttp.BaseResponse
	GetDomainLogsResponse *string `json:"GetDomainLogsResponse" name:"GetDomainLogsResponse"`
}

func (r *GetDomainLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

type GetCdnDomainsRequest struct {
	*ksyunhttp.BaseRequest
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber   *int    `json:"PageNumber,omitempty" name:"PageNumber"`
	DomainName   *string `json:"DomainName,omitempty" name:"DomainName"`
	ProjectId    *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	DomainStatus *string `json:"DomainStatus,omitempty" name:"DomainStatus"`
	CdnType      *string `json:"CdnType,omitempty" name:"CdnType"`
	FuzzyMatch   *string `json:"FuzzyMatch,omitempty" name:"FuzzyMatch"`
}

func (r *GetCdnDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetCdnDomainsResponse struct {
	*ksyunhttp.BaseResponse
	PageNumber *int `json:"PageNumber" name:"PageNumber"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Domains    []struct {
		DomainName      *string `json:"DomainName" name:"DomainName"`
		DomainId        *string `json:"DomainId" name:"DomainId"`
		Cname           *string `json:"Cname" name:"Cname"`
		CdnType         *string `json:"CdnType" name:"CdnType"`
		CdnSubType      *string `json:"CdnSubType" name:"CdnSubType"`
		IcpRegistration *string `json:"IcpRegistration" name:"IcpRegistration"`
		DomainStatus    *string `json:"DomainStatus" name:"DomainStatus"`
		CreatedTime     *string `json:"CreatedTime" name:"CreatedTime"`
		ModifiedTime    *string `json:"ModifiedTime" name:"ModifiedTime"`
		Description     *string `json:"Description" name:"Description"`
		Projectld       *int    `json:"Projectld" name:"Projectld"`
		ProjectName     *string `json:"ProjectName" name:"ProjectName"`
		Region          *string `json:"Region" name:"Region"`
	} `json:"Domains"`
}

func (r *GetCdnDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCdnDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
