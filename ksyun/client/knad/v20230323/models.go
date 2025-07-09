package v20230323

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
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

func (r *CreateKnadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateKnadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateKnadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Kid       *string `json:"kid" name:"kid"`
	Return    *string `json:"Return" name:"Return"`
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
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	IpCount   *int    `json:"IpCount,omitempty" name:"IpCount"`
	KnadId    *string `json:"KnadId,omitempty" name:"KnadId"`
	Band      *int    `json:"Band,omitempty" name:"Band"`
	MaxBand   *int    `json:"MaxBand,omitempty" name:"MaxBand"`
	IdcBand   *int    `json:"IdcBand,omitempty" name:"IdcBand"`
}

func (r *ModifyKnadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKnadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyKnadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKnadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *string `json:"Return" name:"Return"`
}

func (r *ModifyKnadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKnadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KnadIpListRequest struct {
	*ksyunhttp.BaseRequest
	KnadId *string `json:"KnadId,omitempty" name:"KnadId"`
}

func (r *KnadIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KnadIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "KnadIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type KnadIpListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	EipSet    []struct {
		Ip           *string `json:"Ip" name:"Ip"`
		EipId        *string `json:"EipId" name:"EipId"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
	} `json:"EipSet"`
}

func (r *KnadIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KnadIpListResponse) FromJsonString(s string) error {
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

func (r *AssociateIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DisassociateIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeKnadIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeKnadIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteKnadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteKnadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeKnadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeKnadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type ModifyKnadNameRequest struct {
	*ksyunhttp.BaseRequest
	KnadId   *string `json:"KnadId,omitempty" name:"KnadId"`
	KnadName *string `json:"KnadName,omitempty" name:"KnadName"`
}

func (r *ModifyKnadNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKnadNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyKnadNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKnadNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyKnadNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKnadNameResponse) FromJsonString(s string) error {
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

func (r *GetBWIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetBWIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteBWRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteBWRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *AddBWIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddBWIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *GetZoneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetZoneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyBlockLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyBlockLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *GetBlockLocationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetBlockLocationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type GetAttackLogRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	Ip        *string `json:"Ip,omitempty" name:"Ip"`
	PageSize  *int    `json:"PageSize,omitempty" name:"PageSize"`
	OffSet    *int    `json:"OffSet,omitempty" name:"OffSet"`
}

func (r *GetAttackLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetAttackLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	AttackLogSet []struct {
		Ip         *string `json:"Ip" name:"Ip"`
		StartTime  *string `json:"StartTime" name:"StartTime"`
		EndTime    *string `json:"EndTime" name:"EndTime"`
		AttackType *string `json:"AttackType" name:"AttackType"`
		MaxBps     *string `json:"MaxBps" name:"MaxBps"`
	} `json:"AttackLogSet"`
	LogCount *int `json:"LogCount" name:"LogCount"`
}

func (r *GetAttackLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackLogResponse) FromJsonString(s string) error {
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

func (r *GetKnadPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetKnadPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
