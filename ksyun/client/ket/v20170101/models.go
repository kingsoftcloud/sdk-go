package v20170101

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type StartLoopSrcInfo struct {
	Path  *string `json:"Path,omitempty" name:"Path"`
	Index *int    `json:"Index,omitempty" name:"Index"`
}

type PresetRequest struct {
	*ksyunhttp.BaseRequest
	UniqName    *string `json:"UniqName,omitempty" name:"UniqName"`
	App         *string `json:"App,omitempty" name:"App"`
	Preset      *string `json:"Preset,omitempty" name:"Preset"`
	Description *string `json:"Description,omitempty" name:"Description"`
	Output      *string `json:"Output,omitempty" name:"Output"`
	Video       *string `json:"Video,omitempty" name:"Video"`
}

func (r *PresetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type PresetResponse struct {
	*ksyunhttp.BaseResponse
	PresetResponse *string `json:"PresetResponse" name:"PresetResponse"`
}

func (r *PresetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePresetRequest struct {
	*ksyunhttp.BaseRequest
	UniqName    *string `json:"UniqName,omitempty" name:"UniqName"`
	App         *string `json:"App,omitempty" name:"App"`
	Preset      *string `json:"Preset,omitempty" name:"Preset"`
	Description *string `json:"Description,omitempty" name:"Description"`
	Output      *string `json:"Output,omitempty" name:"Output"`
	Video       *string `json:"Video,omitempty" name:"Video"`
}

func (r *UpdatePresetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdatePresetResponse struct {
	*ksyunhttp.BaseResponse
	UpdatePresetResponse *string `json:"UpdatePresetResponse" name:"UpdatePresetResponse"`
}

func (r *UpdatePresetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPresetRequest struct {
	*ksyunhttp.BaseRequest
	UniqName *string `json:"UniqName,omitempty" name:"UniqName"`
	App      *string `json:"App,omitempty" name:"App"`
	Preset   *string `json:"Preset,omitempty" name:"Preset"`
}

func (r *DelPresetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DelPresetResponse struct {
	*ksyunhttp.BaseResponse
	DelPresetResponse *string `json:"DelPresetResponse" name:"DelPresetResponse"`
}

func (r *DelPresetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPresetListRequest struct {
	*ksyunhttp.BaseRequest
	UniqName *string `json:"UniqName,omitempty" name:"UniqName"`
	App      *string `json:"App,omitempty" name:"App"`
}

func (r *GetPresetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetPresetListResponse struct {
	*ksyunhttp.BaseResponse
	GetPresetListResponse *string `json:"GetPresetListResponse" name:"GetPresetListResponse"`
}

func (r *GetPresetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPresetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPresetDetailRequest struct {
	*ksyunhttp.BaseRequest
	UniqName *string `json:"UniqName,omitempty" name:"UniqName"`
	App      *string `json:"App,omitempty" name:"App"`
	Preset   *string `json:"Preset,omitempty" name:"Preset"`
}

func (r *GetPresetDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetPresetDetailResponse struct {
	*ksyunhttp.BaseResponse
	GetPresetDetailResponse *string `json:"GetPresetDetailResponse" name:"GetPresetDetailResponse"`
}

func (r *GetPresetDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPresetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStreamTranListRequest struct {
	*ksyunhttp.BaseRequest
	UniqName *string `json:"UniqName,omitempty" name:"UniqName"`
	App      *string `json:"App,omitempty" name:"App"`
	StreamID *string `json:"StreamID,omitempty" name:"StreamID"`
}

func (r *GetStreamTranListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetStreamTranListResponse struct {
	*ksyunhttp.BaseResponse
	GetStreamTranListResponse *string `json:"GetStreamTranListResponse" name:"GetStreamTranListResponse"`
}

func (r *GetStreamTranListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStreamTranListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartLoopRequest struct {
	*ksyunhttp.BaseRequest
	UniqName      *string             `json:"UniqName,omitempty" name:"UniqName"`
	App           *string             `json:"App,omitempty" name:"App"`
	Preset        *string             `json:"Preset,omitempty" name:"Preset"`
	StreamID      *string             `json:"StreamID,omitempty" name:"StreamID"`
	SrcInfo       []*StartLoopSrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`
	PubDomain     *string             `json:"PubDomain,omitempty" name:"PubDomain"`
	TaskStartTime *string             `json:"TaskStartTime,omitempty" name:"TaskStartTime"`
	TaskStopTime  *string             `json:"TaskStopTime,omitempty" name:"TaskStopTime"`
	LoopTimes     *int                `json:"LoopTimes,omitempty" name:"LoopTimes"`
}

func (r *StartLoopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartLoopResponse struct {
	*ksyunhttp.BaseResponse
	ErrNum *int    `json:"ErrNum" name:"ErrNum"`
	ErrMsg *string `json:"ErrMsg" name:"ErrMsg"`
	List   []struct {
		TaskID *string `json:"TaskID" name:"TaskID"`
		Format *int    `json:"Format" name:"Format"`
	} `json:"List"`
}

func (r *StartLoopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartLoopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopLoopRequest struct {
	*ksyunhttp.BaseRequest
	UniqName *string `json:"UniqName,omitempty" name:"UniqName"`
	App      *string `json:"App,omitempty" name:"App"`
	StreamID *string `json:"StreamID,omitempty" name:"StreamID"`
}

func (r *StopLoopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopLoopResponse struct {
	*ksyunhttp.BaseResponse
	ErrNum *int    `json:"ErrNum" name:"ErrNum"`
	ErrMsg *string `json:"ErrMsg" name:"ErrMsg"`
}

func (r *StopLoopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopLoopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
