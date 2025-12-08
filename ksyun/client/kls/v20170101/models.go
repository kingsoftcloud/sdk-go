package v20170101
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type ListStreamDurationsRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName    *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App           *string `json:"App,omitempty" name:"App"`
	Pubdomain     *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
	Stream        *string `json:"Stream,omitempty" name:"Stream"`
	StartUnixTime *int    `json:"StartUnixTime,omitempty" name:"StartUnixTime"`
	EndUnixTime   *int    `json:"EndUnixTime,omitempty" name:"EndUnixTime"`
}

func (r *ListStreamDurationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListStreamDurationsResponse struct {
	*ksyunhttp.BaseResponse
	ListStreamDurationsResponse *string `json:"ListStreamDurationsResponse" name:"ListStreamDurationsResponse"`
}

func (r *ListStreamDurationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStreamDurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListHistoryPubStreamsErrInfoRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName    *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App           *string `json:"App,omitempty" name:"App"`
	Pubdomain     *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
	Stream        *string `json:"Stream,omitempty" name:"Stream"`
	OrderTime     *int    `json:"OrderTime,omitempty" name:"OrderTime"`
	Marker        *int    `json:"Marker,omitempty" name:"Marker"`
	Limit         *int    `json:"Limit,omitempty" name:"Limit"`
	StartUnixTime *int    `json:"StartUnixTime,omitempty" name:"StartUnixTime"`
	EndUnixTime   *int    `json:"EndUnixTime,omitempty" name:"EndUnixTime"`
}

func (r *ListHistoryPubStreamsErrInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListHistoryPubStreamsErrInfoResponse struct {
	*ksyunhttp.BaseResponse
	ListHistoryPubStreamsErrInfoResponse *string `json:"ListHistoryPubStreamsErrInfoResponse" name:"ListHistoryPubStreamsErrInfoResponse"`
}

func (r *ListHistoryPubStreamsErrInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHistoryPubStreamsErrInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListHistoryPubStreamsInfoRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName    *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App           *string `json:"App,omitempty" name:"App"`
	Pubdomain     *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
	Stream        *string `json:"Stream,omitempty" name:"Stream"`
	OrderTime     *int    `json:"OrderTime,omitempty" name:"OrderTime"`
	Marker        *int    `json:"Marker,omitempty" name:"Marker"`
	Limit         *int    `json:"Limit,omitempty" name:"Limit"`
	StartUnixTime *int    `json:"StartUnixTime,omitempty" name:"StartUnixTime"`
	EndUnixTime   *int    `json:"EndUnixTime,omitempty" name:"EndUnixTime"`
}

func (r *ListHistoryPubStreamsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListHistoryPubStreamsInfoResponse struct {
	*ksyunhttp.BaseResponse
	ListHistoryPubStreamsInfoResponse *string `json:"ListHistoryPubStreamsInfoResponse" name:"ListHistoryPubStreamsInfoResponse"`
}

func (r *ListHistoryPubStreamsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHistoryPubStreamsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ForbidStreamRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName         *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App                *string `json:"App,omitempty" name:"App"`
	Pubdomain          *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
	Stream             *string `json:"Stream,omitempty" name:"Stream"`
	ForbidTillUnixTime *int    `json:"ForbidTillUnixTime,omitempty" name:"ForbidTillUnixTime"`
}

func (r *ForbidStreamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ForbidStreamResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RetCode *int `json:"RetCode" name:"RetCode"`
		RetMsg *string `json:"RetMsg" name:"RetMsg"`
	} `json:"Data"`
}

func (r *ForbidStreamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ForbidStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ResumeStreamRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App        *string `json:"App,omitempty" name:"App"`
	Pubdomain  *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
	Stream     *string `json:"Stream,omitempty" name:"Stream"`
}

func (r *ResumeStreamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResumeStreamResponse struct {
	*ksyunhttp.BaseResponse
	ResumeStreamResponse *string `json:"ResumeStreamResponse" name:"ResumeStreamResponse"`
}

func (r *ResumeStreamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetBlacklistRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App        *string `json:"App,omitempty" name:"App"`
	Pubdomain  *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
}

func (r *GetBlacklistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBlacklistResponse struct {
	*ksyunhttp.BaseResponse
	GetBlacklistResponse *string `json:"GetBlacklistResponse" name:"GetBlacklistResponse"`
}

func (r *GetBlacklistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBlacklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CheckBlacklistRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App        *string `json:"App,omitempty" name:"App"`
	Pubdomain  *string `json:"Pubdomain,omitempty" name:"Pubdomain"`
	Stream     *string `json:"Stream,omitempty" name:"Stream"`
}

func (r *CheckBlacklistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CheckBlacklistResponse struct {
	*ksyunhttp.BaseResponse
	CheckBlacklistResponse *string `json:"CheckBlacklistResponse" name:"CheckBlacklistResponse"`
}

func (r *CheckBlacklistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBlacklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRealtimeStreamsInfoRequest struct {
	*ksyunhttp.BaseRequest
	UniqueName   *string `json:"UniqueName,omitempty" name:"UniqueName"`
	App          *string `json:"App,omitempty" name:"App"`
	Stream       *string `json:"Stream,omitempty" name:"Stream"`
	DomainIds    *string `json:"DomainIds,omitempty" name:"DomainIds"`
	PullProtocol *string `json:"PullProtocol,omitempty" name:"PullProtocol"`
	Type         *string `json:"Type,omitempty" name:"Type"`
}

func (r *ListRealtimeStreamsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListRealtimeStreamsInfoResponse struct {
	*ksyunhttp.BaseResponse
	ListRealtimeStreamsInfoResponse *string `json:"ListRealtimeStreamsInfoResponse" name:"ListRealtimeStreamsInfoResponse"`
}

func (r *ListRealtimeStreamsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRealtimeStreamsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

