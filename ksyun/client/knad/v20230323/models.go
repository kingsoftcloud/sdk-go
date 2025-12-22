package v20230323

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateKnadRequest struct {
	*ksyunhttp.BaseRequest
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	KnadName  *string `json:"KnadName,omitempty" name:"KnadName"`
	Duration  *int    `json:"Duration,omitempty" name:"Duration"`
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	Band      *int    `json:"Band,omitempty" name:"Band"`
	MaxBand   *int    `json:"MaxBand,omitempty" name:"MaxBand"`
	IpCount   *int    `json:"IpCount,omitempty" name:"IpCount"`
	BillType  *int    `json:"BillType,omitempty" name:"BillType"`
	IdcBand   *int    `json:"IdcBand,omitempty" name:"IdcBand"`
}

func (r *CreateKnadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateKnadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Kid       *string `json:"kid" name:"kid"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CreateKnadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKnadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKnadRequest struct {
	*ksyunhttp.BaseRequest
	KnadId    *string `json:"KnadId,omitempty" name:"KnadId"`
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	IpCount   *int    `json:"IpCount,omitempty" name:"IpCount"`
	Band      *int    `json:"Band,omitempty" name:"Band"`
	MaxBand   *int    `json:"MaxBand,omitempty" name:"MaxBand"`
	IdcBand   *int    `json:"IdcBand,omitempty" name:"IdcBand"`
}

func (r *ModifyKnadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyKnadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyKnadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKnadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindIpListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *UnbindIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UnbindIpListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	EipCount  *int    `json:"EipCount" name:"EipCount"`
	EipSet    []struct {
		Ip           *string `json:"Ip" name:"Ip"`
		EipId        *string `json:"EipId" name:"EipId"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
	} `json:"EipSet"`
}

func (r *UnbindIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateIpRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string   `json:"KnadId,omitempty" name:"KnadId"`
	Ip     []*string `json:"Ip,omitempty" name:"Ip"`
}

func (r *AssociateIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AssociateIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Return    *bool     `json:"Return" name:"Return"`
	FailEips  []*string `json:"FailEips" name:"FailEips"`
}

func (r *AssociateIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateIpRequest struct {
	*ksyunhttp.BaseRequest
	Ip []*string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DisassociateIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisassociateIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	DeletedEipSet []struct {
		KnadId     *string `json:"KnadId" name:"KnadId"`
		OldIpCount *int    `json:"OldIpCount" name:"OldIpCount"`
		IpCount    *int    `json:"IpCount" name:"IpCount"`
		Return     *bool   `json:"Return" name:"Return"`
	} `json:"DeletedEipSet"`
}

func (r *DisassociateIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnadIpRequest struct {
	*ksyunhttp.BaseRequest
	IpSort       *string `json:"IpSort,omitempty" name:"IpSort"`
	IpStatusSort *string `json:"IpStatusSort,omitempty" name:"IpStatusSort"`
	BandSort     *string `json:"BandSort,omitempty" name:"BandSort"`
	Ip           *string `json:"Ip,omitempty" name:"Ip"`
	KnadId       *string `json:"KnadId,omitempty" name:"KnadId"`
	ProjectId    []*int  `json:"ProjectId,omitempty" name:"ProjectId"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
	OffSet       *int    `json:"OffSet,omitempty" name:"OffSet"`
}

func (r *DescribeKnadIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnadIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	KnadIpSet []struct {
		KnadId       *string `json:"KnadId" name:"KnadId"`
		KnadName     *string `json:"KnadName" name:"KnadName"`
		EipId        *string `json:"EipId" name:"EipId"`
		KnadIp       *string `json:"KnadIp" name:"KnadIp"`
		StatusDesc   *string `json:"StatusDesc" name:"StatusDesc"`
		ProjectId    *int    `json:"ProjectId" name:"ProjectId"`
		Band         *int    `json:"Band" name:"Band"`
		MaxBand      *int    `json:"MaxBand" name:"MaxBand"`
		TemplateId   *int    `json:"TemplateId" name:"TemplateId"`
		TemplateName *string `json:"TemplateName" name:"TemplateName"`
		EipInfo      struct {
			InstanceType *string `json:"InstanceType" name:"InstanceType"`
		} `json:"EipInfo" name:"EipInfo"`
	} `json:"KnadIpSet"`
	KnadIpCount *int `json:"KnadIpCount" name:"KnadIpCount"`
}

func (r *DescribeKnadIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnadIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKnadRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string `json:"KnadId,omitempty" name:"KnadId"`
}

func (r *DeleteKnadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteKnadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *string `json:"Return" name:"Return"`
}

func (r *DeleteKnadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKnadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnadRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
	KnadId    []*string `json:"KnadId,omitempty" name:"KnadId"`
}

func (r *DescribeKnadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	KnadSet   []struct {
		KnadId      *string `json:"KnadId" name:"KnadId"`
		KnadName    *string `json:"KnadName" name:"KnadName"`
		Band        *int    `json:"Band" name:"Band"`
		MaxBand     *int    `json:"MaxBand" name:"MaxBand"`
		ProjectId   *string `json:"ProjectId" name:"ProjectId"`
		UsedIpCount *int    `json:"UsedIpCount" name:"UsedIpCount"`
		IpCount     *int    `json:"IpCount" name:"IpCount"`
		TemplateId  *int    `json:"TemplateId" name:"TemplateId"`
		ExprieTime  *string `json:"ExprieTime" name:"ExprieTime"`
		UdpBlock    *int    `json:"UdpBlock" name:"UdpBlock"`
		OnlyChinaIp *int    `json:"OnlyChinaIp" name:"OnlyChinaIp"`
		IdcBand     *int    `json:"IdcBand" name:"IdcBand"`
		ServiceId   *string `json:"ServiceId" name:"ServiceId"`
	} `json:"KnadSet"`
}

func (r *DescribeKnadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBWIpListRequest struct {
	*ksyunhttp.BaseRequest
	KnadId   *string `json:"KnadId,omitempty" name:"KnadId"`
	Type     *int    `json:"type,omitempty" name:"type"`
	Ip       *string `json:"Ip,omitempty" name:"Ip"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
	OffSet   *int    `json:"OffSet,omitempty" name:"OffSet"`
}

func (r *GetBWIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBWIpListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	IpList    []struct {
		Id *string `json:"Id" name:"Id"`
		Ip *string `json:"Ip" name:"Ip"`
	} `json:"IpList"`
	Total *int `json:"total" name:"total"`
}

func (r *GetBWIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBWIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBWRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string   `json:"KnadId,omitempty" name:"KnadId"`
	IpId   []*string `json:"IpId,omitempty" name:"IpId"`
	Type   *int      `json:"type,omitempty" name:"type"`
}

func (r *DeleteBWRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBWResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteBWResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBWResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddBWIpListRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string   `json:"KnadId,omitempty" name:"KnadId"`
	Ip     []*string `json:"Ip,omitempty" name:"Ip"`
	Type   *int      `json:"type,omitempty" name:"type"`
}

func (r *AddBWIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddBWIpListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AddBWIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBWIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZoneListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetZoneListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetZoneListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetZoneListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyRequest struct {
	*ksyunhttp.BaseRequest
	KnadId     *string `json:"KnadId,omitempty" name:"KnadId"`
	TemplateId *int    `json:"TemplateId,omitempty" name:"TemplateId"`
	UdpBlock   *int    `json:"UdpBlock,omitempty" name:"UdpBlock"`
	TcpBlock   *int    `json:"TcpBlock,omitempty" name:"TcpBlock"`
	IcmpBlock  *int    `json:"IcmpBlock,omitempty" name:"IcmpBlock"`
}

func (r *ModifyPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockLocationRequest struct {
	*ksyunhttp.BaseRequest
	KnadId        *string   `json:"KnadId,omitempty" name:"KnadId"`
	LocationBlock *int      `json:"LocationBlock,omitempty" name:"LocationBlock"`
	Location      []*string `json:"Location,omitempty" name:"Location"`
}

func (r *ModifyBlockLocationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyBlockLocationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyBlockLocationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBlockLocationsRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string `json:"KnadId,omitempty" name:"KnadId"`
}

func (r *GetBlockLocationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBlockLocationsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	LocationBlock    *int    `json:"LocationBlock" name:"LocationBlock"`
	BlockLocationSet []struct {
		ZoneKey  *string `json:"ZoneKey" name:"ZoneKey"`
		ZoneName *string `json:"ZoneName" name:"ZoneName"`
	} `json:"BlockLocationSet"`
}

func (r *GetBlockLocationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBlockLocationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetKnadPolicyRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string `json:"KnadId,omitempty" name:"KnadId"`
}

func (r *GetKnadPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetKnadPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TemplateId *int    `json:"TemplateId" name:"TemplateId"`
	UdpBlock   *int    `json:"UdpBlock" name:"UdpBlock"`
	TcpBlock   *int    `json:"TcpBlock" name:"TcpBlock"`
	IcmpBlock  *int    `json:"IcmpBlock" name:"IcmpBlock"`
}

func (r *GetKnadPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetKnadPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertEipsRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string   `json:"KnadId,omitempty" name:"KnadId"`
	Ip     []*string `json:"Ip,omitempty" name:"Ip"`
}

func (r *InsertEipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type InsertEipsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *InsertEipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
