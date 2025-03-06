package v20170101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
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

func (r *ListStreamDurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListStreamDurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListHistoryPubStreamsErrInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListHistoryPubStreamsErrInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListHistoryPubStreamsErrInfoResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		App       *string `json:"App" name:"App"`
		Pubdomain *string `json:"Pubdomain" name:"Pubdomain"`
		Result    []struct {
			Detail []struct {
				Clientip      *string `json:"Clientip" name:"Clientip"`
				StatusMessage *string `json:"StatusMessage" name:"StatusMessage"`
			} `json:"Detail"`
			Stream *string `json:"Stream" name:"Stream"`
		} `json:"Result" name:"Result"`
		RetMsg     *string `json:"RetMsg" name:"RetMsg"`
		UniqueName *string `json:"UniqueName" name:"UniqueName"`
	} `json:"Data"`
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

func (r *ListHistoryPubStreamsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListHistoryPubStreamsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListHistoryPubStreamsInfoResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		App       *string `json:"App" name:"App"`
		Pubdomain *string `json:"Pubdomain" name:"Pubdomain"`
		Result    []struct {
			Detail []struct {
				Clientip      *string `json:"Clientip" name:"Clientip"`
				StatusMessage *string `json:"StatusMessage" name:"StatusMessage"`
			} `json:"Detail"`
			Stream *string `json:"Stream" name:"Stream"`
		} `json:"Result" name:"Result"`
		RetMsg     *string `json:"RetMsg" name:"RetMsg"`
		UniqueName *string `json:"UniqueName" name:"UniqueName"`
	} `json:"Data"`
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

func (r *ForbidStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ForbidStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ForbidStreamResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
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

func (r *ResumeStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResumeStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResumeStreamResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RetMsg *string `json:"RetMsg" name:"RetMsg"`
	} `json:"Data"`
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

func (r *GetBlacklistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetBlacklistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBlacklistResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		App       *string `json:"App" name:"App"`
		Pubdomain *string `json:"Pubdomain" name:"Pubdomain"`
		Recs      []struct {
			Stream *string `json:"Stream" name:"Stream"`
		} `json:"Recs" name:"Recs"`
		RetMsg     *string `json:"RetMsg" name:"RetMsg"`
		UniqueName *string `json:"UniqueName" name:"UniqueName"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *CheckBlacklistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CheckBlacklistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckBlacklistResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RetMsg *string `json:"RetMsg" name:"RetMsg"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *ListRealtimeStreamsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRealtimeStreamsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
