package v20160304

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type BindZoneVpcVpcs struct {
	RegionName *string   `json:"RegionName,omitempty" name:"RegionName"`
	VpcId      []*string `json:"VpcId,omitempty" name:"VpcId"`
}

type UnbindZoneVpcVpcs struct {
	RegionName *string   `json:"RegionName,omitempty" name:"RegionName"`
	VpcId      []*string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeZoneRecordFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreatePrivateDnsRequest struct {
	*ksyunhttp.BaseRequest
	Action    *string `json:"Action,omitempty" name:"Action"`
	Version   *string `json:"Version,omitempty" name:"Version"`
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreatePrivateDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateDnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePrivateDnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateDnsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *CreatePrivateDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateDnsRequest struct {
	*ksyunhttp.BaseRequest
	Action    *string `json:"Action,omitempty" name:"Action"`
	Version   *string `json:"Version,omitempty" name:"Version"`
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeletePrivateDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateDnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePrivateDnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateDnsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeletePrivateDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateDnsRequest struct {
	*ksyunhttp.BaseRequest
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
}

func (r *DescribePrivateDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateDnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePrivateDnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateDnsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribePrivateDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateVpcsRequest struct {
	*ksyunhttp.BaseRequest
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
	VpcId   *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *AssociateVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateVpcsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *AssociateVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateVpcsRequest struct {
	*ksyunhttp.BaseRequest
	Action  *string `json:"Action,omitempty" name:"Action"`
	Version *string `json:"Version,omitempty" name:"Version"`
	VpcId   *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DisassociateVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateVpcsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DisassociateVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneRequest struct {
	*ksyunhttp.BaseRequest
	Action   *string `json:"Action,omitempty" name:"Action"`
	Version  *string `json:"Version,omitempty" name:"Version"`
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	ZoneTtl  *int    `json:"ZoneTtl,omitempty" name:"ZoneTtl"`
}

func (r *CreateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *CreateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *ModifyZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ModifyZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *CreateRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *CreateRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DeleteRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordRequest struct {
	*ksyunhttp.BaseRequest
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
}

func (r *ModifyRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ModifyRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordDataRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *CreateRecordDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRecordDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordDataResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *CreateRecordDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordDataRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DeleteRecordDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRecordDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordDataResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteRecordDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePdnsZoneRequest struct {
	*ksyunhttp.BaseRequest
	ZoneName   *string `json:"ZoneName,omitempty" name:"ZoneName"`
	ZoneTtl    *int    `json:"ZoneTtl,omitempty" name:"ZoneTtl"`
	ProjectId  *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreatePdnsZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePdnsZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePdnsZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePdnsZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreatePdnsZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePdnsZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPdnsZoneRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId  *string `json:"ZoneId,omitempty" name:"ZoneId"`
	ZoneTtl *int    `json:"ZoneTtl,omitempty" name:"ZoneTtl"`
}

func (r *ModifyPdnsZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPdnsZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyPdnsZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPdnsZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyPdnsZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPdnsZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePdnsZoneRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DeletePdnsZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePdnsZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePdnsZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePdnsZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeletePdnsZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePdnsZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePdnsZonesRequest struct {
	*ksyunhttp.BaseRequest
	Filter     []*string `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribePdnsZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePdnsZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePdnsZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePdnsZonesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	ZoneSet   []struct {
		ZoneId     *string `json:"ZoneId"`
		ZoneName   *string `json:"ZoneName"`
		CreateTime *string `json:"CreateTime"`
		ProjectId  *string `json:"ProjectId"`
		ZoneTtl    *string `json:"ZoneTtl"`
		BindVpcSet []struct {
			RegionName *string `json:"RegionName"`
			VpcId      *string `json:"VpcId"`
			Status     *string `json:"Status"`
			VpcName    *string `json:"VpcName"`
		} `json:"BindVpcSet"`
	} `json:"ZoneSet"`
}

func (r *DescribePdnsZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePdnsZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindZoneVpcRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId *string            `json:"ZoneId,omitempty" name:"ZoneId"`
	Vpcs   []*BindZoneVpcVpcs `json:"Vpcs,omitempty" name:"Vpcs"`
}

func (r *BindZoneVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BindZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindZoneVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *BindZoneVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindZoneVpcRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId *string              `json:"ZoneId,omitempty" name:"ZoneId"`
	Vpcs   []*UnbindZoneVpcVpcs `json:"Vpcs,omitempty" name:"Vpcs"`
}

func (r *UnbindZoneVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnbindZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindZoneVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *UnbindZoneVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneRecordRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId      *string `json:"ZoneId,omitempty" name:"ZoneId"`
	RecordName  *string `json:"RecordName,omitempty" name:"RecordName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
	RecordTtl   *int    `json:"RecordTtl,omitempty" name:"RecordTtl"`
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
}

func (r *CreateZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneRecordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneRecordRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId      *string `json:"ZoneId,omitempty" name:"ZoneId"`
	RecordId    *string `json:"RecordId,omitempty" name:"RecordId"`
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
	Priority    *string `json:"Priority,omitempty" name:"Priority"`
	Weight      *string `json:"Weight,omitempty" name:"Weight"`
	Port        *string `json:"Port,omitempty" name:"Port"`
}

func (r *DeleteZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneRecordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneRecordRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId      *string `json:"ZoneId,omitempty" name:"ZoneId"`
	RecordId    *string `json:"RecordId,omitempty" name:"RecordId"`
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
	RecordTtl   *int    `json:"RecordTtl,omitempty" name:"RecordTtl"`
}

func (r *ModifyZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneRecordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneRecordRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId   *string                     `json:"ZoneId,omitempty" name:"ZoneId"`
	RecordId []*string                   `json:"RecordId,omitempty" name:"RecordId"`
	Filter   []*DescribeZoneRecordFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneRecordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RecordSet []struct {
		RecordId      *string `json:"RecordId"`
		RecordName    *string `json:"RecordName"`
		Type          *string `json:"Type"`
		RecordTtl     *int    `json:"RecordTtl"`
		RecordDataSet []struct {
			RecordValue *string `json:"RecordValue"`
			Priority    *int    `json:"Priority"`
			Weight      *int    `json:"Weight"`
			Port        *int    `json:"Port"`
		} `json:"RecordDataSet"`
	} `json:"RecordSet"`
}

func (r *DescribeZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
