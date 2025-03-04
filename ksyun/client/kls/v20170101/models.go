package v20170101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type ListStreamDurationsRequest struct {
	*ksyunhttp.BaseRequest
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
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
