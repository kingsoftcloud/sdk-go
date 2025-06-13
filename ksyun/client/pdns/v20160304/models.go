package v20160304

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type BindZoneVpcVpcs struct {
	RegionName *string   `json:"RegionName,omitempty" name:"RegionName"`
	VpcId      []*string `json:"VpcId,omitempty" name:"VpcId"`
}
type UnbindZoneVpcVpcs struct {
	RegionName *string   `json:"RegionName,omitempty" name:"RegionName"`
	VpcId      []*string `json:"VpcId,omitempty" name:"VpcId"`
}
type ModifyPdnsFdZoneForwardIp struct {
	Port *string `json:"Port,omitempty" name:"Port"`
	Ip   *string `json:"Ip,omitempty" name:"Ip"`
}
type CreatePdnsFdZoneForwardIp struct {
	Port *string `json:"Port,omitempty" name:"Port"`
	Ip   *string `json:"Ip,omitempty" name:"Ip"`
}
type ModifyEndPointIpConfig struct {
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	SubnetId         *string `json:"SubnetId,omitempty" name:"SubnetId"`
	IP               *string `json:"IP,omitempty" name:"IP"`
}
type CreateEndPointIpConfig struct {
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	SubnetId         *string `json:"SubnetId,omitempty" name:"SubnetId"`
	IP               *string `json:"IP,omitempty" name:"IP"`
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
	ZoneVpc   struct {
		ZoneId     *string `json:"ZoneId" name:"ZoneId"`
		ZoneName   *string `json:"ZoneName" name:"ZoneName"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		ProjectId  *string `json:"ProjectId" name:"ProjectId"`
		ZoneTtl    *string `json:"ZoneTtl" name:"ZoneTtl"`
		BindVpcSet []struct {
			RegionName *string `json:"RegionName" name:"RegionName"`
			VpcId      *string `json:"VpcId" name:"VpcId"`
			Status     *string `json:"Status" name:"Status"`
			VpcName    *string `json:"VpcName" name:"VpcName"`
		} `json:"BindVpcSet" name:"BindVpcSet"`
	} `json:"ZoneVpc"`
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
		ZoneId     *string `json:"ZoneId" name:"ZoneId"`
		ZoneName   *string `json:"ZoneName" name:"ZoneName"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		ProjectId  *string `json:"ProjectId" name:"ProjectId"`
		ZoneTtl    *string `json:"ZoneTtl" name:"ZoneTtl"`
		BindVpcSet []struct {
			RegionName *string `json:"RegionName" name:"RegionName"`
			VpcId      *string `json:"VpcId" name:"VpcId"`
			Status     *string `json:"Status" name:"Status"`
			VpcName    *string `json:"VpcName" name:"VpcName"`
		} `json:"BindVpcSet" name:"BindVpcSet"`
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
	Record    struct {
		RecordId      *string `json:"RecordId" name:"RecordId"`
		RecordName    *string `json:"RecordName" name:"RecordName"`
		Type          *string `json:"Type" name:"Type"`
		RecordTtl     *int    `json:"RecordTtl" name:"RecordTtl"`
		RecordDataSet []struct {
			RecordValue *string `json:"RecordValue" name:"RecordValue"`
			Priority    *int    `json:"Priority" name:"Priority"`
			Weight      *int    `json:"Weight" name:"Weight"`
			Port        *int    `json:"Port" name:"Port"`
		} `json:"RecordDataSet" name:"RecordDataSet"`
	} `json:"Record"`
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
	ZoneId    *string `json:"ZoneId,omitempty" name:"ZoneId"`
	RecordId  *string `json:"RecordId,omitempty" name:"RecordId"`
	RecordTtl *int    `json:"RecordTtl,omitempty" name:"RecordTtl"`
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
	Record    struct {
		RecordId      *string `json:"RecordId" name:"RecordId"`
		RecordName    *string `json:"RecordName" name:"RecordName"`
		Type          *string `json:"Type" name:"Type"`
		RecordTtl     *int    `json:"RecordTtl" name:"RecordTtl"`
		RecordDataSet []struct {
			RecordValue *string `json:"RecordValue" name:"RecordValue"`
			Priority    *int    `json:"Priority" name:"Priority"`
			Weight      *int    `json:"Weight" name:"Weight"`
			Port        *int    `json:"Port" name:"Port"`
		} `json:"RecordDataSet" name:"RecordDataSet"`
	} `json:"Record"`
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
	ZoneId   *string   `json:"ZoneId,omitempty" name:"ZoneId"`
	RecordId []*string `json:"RecordId,omitempty" name:"RecordId"`
	Filter   []*string `json:"Filter,omitempty" name:"Filter"`
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
		RecordId      *string `json:"RecordId" name:"RecordId"`
		RecordName    *string `json:"RecordName" name:"RecordName"`
		Type          *string `json:"Type" name:"Type"`
		RecordTtl     *int    `json:"RecordTtl" name:"RecordTtl"`
		RecordDataSet []struct {
			RecordValue *string `json:"RecordValue" name:"RecordValue"`
			Priority    *int    `json:"Priority" name:"Priority"`
			Weight      *int    `json:"Weight" name:"Weight"`
			Port        *int    `json:"Port" name:"Port"`
		} `json:"RecordDataSet" name:"RecordDataSet"`
	} `json:"RecordSet"`
}

func (r *DescribeZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindFdZoneVpcRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *UnbindFdZoneVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindFdZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnbindFdZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindFdZoneVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *UnbindFdZoneVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindFdZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindFdZoneVpcRequest struct {
	*ksyunhttp.BaseRequest
	FdZoneId   *string `json:"FdZoneId,omitempty" name:"FdZoneId"`
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *BindFdZoneVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindFdZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BindFdZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindFdZoneVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Return     *bool   `json:"Return" name:"Return"`
	BindVpcSet []struct {
		BindVpcId  *string `json:"BindVpcId" name:"BindVpcId"`
		RegionName *string `json:"RegionName" name:"RegionName"`
		VpcId      *string `json:"VpcId" name:"VpcId"`
		FdZoneId   *string `json:"FdZoneId" name:"FdZoneId"`
		Status     *string `json:"Status" name:"Status"`
		Created    *string `json:"Created" name:"Created"`
	} `json:"BindVpcSet"`
}

func (r *BindFdZoneVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindFdZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePdnsFdZoneRequest struct {
	*ksyunhttp.BaseRequest
	FdZoneId []*string `json:"FdZoneId,omitempty" name:"FdZoneId"`
}

func (r *DescribePdnsFdZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePdnsFdZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePdnsFdZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePdnsFdZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	FdZoneSet []struct {
		Id          *string `json:"Id" name:"Id"`
		ZoneName    *string `json:"ZoneName" name:"ZoneName"`
		EndPointId  *string `json:"EndPointId" name:"EndPointId"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		Description *string `json:"Description" name:"Description"`
		BindVpcSet  []struct {
			BindVpcId  *string `json:"BindVpcId" name:"BindVpcId"`
			RegionName *string `json:"RegionName" name:"RegionName"`
			VpcId      *string `json:"VpcId" name:"VpcId"`
			FdZoneId   *string `json:"FdZoneId" name:"FdZoneId"`
			Status     *string `json:"Status" name:"Status"`
			Created    *string `json:"Created" name:"Created"`
		} `json:"BindVpcSet" name:"BindVpcSet"`
	} `json:"FdZoneSet"`
}

func (r *DescribePdnsFdZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePdnsFdZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePdnsFdZoneRequest struct {
	*ksyunhttp.BaseRequest
	FdZoneId *string `json:"FdZoneId,omitempty" name:"FdZoneId"`
}

func (r *DeletePdnsFdZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePdnsFdZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePdnsFdZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePdnsFdZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeletePdnsFdZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePdnsFdZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPdnsFdZoneRequest struct {
	*ksyunhttp.BaseRequest
	FdZoneId    *string                      `json:"FdZoneId,omitempty" name:"FdZoneId"`
	Description *string                      `json:"Description,omitempty" name:"Description"`
	ForwardIp   []*ModifyPdnsFdZoneForwardIp `json:"ForwardIp,omitempty" name:"ForwardIp"`
}

func (r *ModifyPdnsFdZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPdnsFdZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyPdnsFdZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPdnsFdZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	FdZone    struct {
		Id          *string `json:"Id" name:"Id"`
		ZoneName    *string `json:"ZoneName" name:"ZoneName"`
		EndPointId  *string `json:"EndPointId" name:"EndPointId"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		Description *string `json:"Description" name:"Description"`
		BindVpcSet  []struct {
			BindVpcId  *string `json:"BindVpcId" name:"BindVpcId"`
			RegionName *string `json:"RegionName" name:"RegionName"`
			VpcId      *string `json:"VpcId" name:"VpcId"`
			FdZoneId   *string `json:"FdZoneId" name:"FdZoneId"`
			Status     *string `json:"Status" name:"Status"`
			Created    *string `json:"Created" name:"Created"`
		} `json:"BindVpcSet" name:"BindVpcSet"`
	} `json:"FdZone"`
}

func (r *ModifyPdnsFdZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPdnsFdZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePdnsFdZoneRequest struct {
	*ksyunhttp.BaseRequest
	EndPointId  *string                      `json:"EndPointId,omitempty" name:"EndPointId"`
	FdZoneName  *string                      `json:"FdZoneName,omitempty" name:"FdZoneName"`
	Description *string                      `json:"Description,omitempty" name:"Description"`
	ForwardIp   []*CreatePdnsFdZoneForwardIp `json:"ForwardIp,omitempty" name:"ForwardIp"`
}

func (r *CreatePdnsFdZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePdnsFdZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePdnsFdZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePdnsFdZoneResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	FdZone    struct {
		Id          *string `json:"Id" name:"Id"`
		ZoneName    *string `json:"ZoneName" name:"ZoneName"`
		EndPointId  *string `json:"EndPointId" name:"EndPointId"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		Description *string `json:"Description" name:"Description"`
		BindVpcSet  []struct {
			BindVpcId  *string `json:"BindVpcId" name:"BindVpcId"`
			RegionName *string `json:"RegionName" name:"RegionName"`
			VpcId      *string `json:"VpcId" name:"VpcId"`
			FdZoneId   *string `json:"FdZoneId" name:"FdZoneId"`
			Status     *string `json:"Status" name:"Status"`
			Created    *string `json:"Created" name:"Created"`
		} `json:"BindVpcSet" name:"BindVpcSet"`
	} `json:"FdZone"`
}

func (r *CreatePdnsFdZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePdnsFdZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryEndPointRegionAZRequest struct {
	*ksyunhttp.BaseRequest
	Region *string `json:"Region,omitempty" name:"Region"`
}

func (r *QueryEndPointRegionAZRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryEndPointRegionAZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryEndPointRegionAZRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryEndPointRegionAZResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string   `json:"RequestId" name:"RequestId"`
	AvailableAz []*string `json:"AvailableAz" name:"AvailableAz"`
}

func (r *QueryEndPointRegionAZResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryEndPointRegionAZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEndPointsRequest struct {
	*ksyunhttp.BaseRequest
	EndPointId []*string `json:"EndPointId,omitempty" name:"EndPointId"`
}

func (r *DescribeEndPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEndPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEndPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEndPointsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	EndPointSet []struct {
		EndPointId      *string   `json:"EndPointId" name:"EndPointId"`
		Region          *string   `json:"Region" name:"Region"`
		EndPointName    *string   `json:"EndPointName" name:"EndPointName"`
		Type            *string   `json:"Type" name:"Type"`
		VpcId           *string   `json:"VpcId" name:"VpcId"`
		SecurityGroupId *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		FdZoneIds       []*string `json:"FdZoneIds" name:"FdZoneIds"`
		CreateTime      *string   `json:"CreateTime" name:"CreateTime"`
		Status          *string   `json:"Status" name:"Status"`
		Description     *string   `json:"Description" name:"Description"`
		IpConfigSet     []struct {
			AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			SubnetId         *string `json:"SubnetId" name:"SubnetId"`
			Ip               *string `json:"Ip" name:"Ip"`
		} `json:"IpConfigSet" name:"IpConfigSet"`
	} `json:"EndPointSet"`
}

func (r *DescribeEndPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEndPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEndPointRequest struct {
	*ksyunhttp.BaseRequest
	EndPointId *string `json:"EndPointId,omitempty" name:"EndPointId"`
}

func (r *DeleteEndPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEndPointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteEndPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEndPointRequest struct {
	*ksyunhttp.BaseRequest
	EndPointId   *string                   `json:"EndPointId,omitempty" name:"EndPointId"`
	EndPointName *string                   `json:"EndPointName,omitempty" name:"EndPointName"`
	Description  *string                   `json:"Description,omitempty" name:"Description"`
	IpConfig     []*ModifyEndPointIpConfig `json:"IpConfig,omitempty" name:"IpConfig"`
}

func (r *ModifyEndPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEndPointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	EndPoint  struct {
		EndPointId      *string   `json:"EndPointId" name:"EndPointId"`
		Region          *string   `json:"Region" name:"Region"`
		EndPointName    *string   `json:"EndPointName" name:"EndPointName"`
		Type            *string   `json:"Type" name:"Type"`
		VpcId           *string   `json:"VpcId" name:"VpcId"`
		SecurityGroupId *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		FdZoneIds       []*string `json:"FdZoneIds" name:"FdZoneIds"`
		CreateTime      *string   `json:"CreateTime" name:"CreateTime"`
		Status          *string   `json:"Status" name:"Status"`
		Description     *string   `json:"Description" name:"Description"`
		IpConfigSet     []struct {
			AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			SubnetId         *string `json:"SubnetId" name:"SubnetId"`
			Ip               *string `json:"Ip" name:"Ip"`
		} `json:"IpConfigSet" name:"IpConfigSet"`
	} `json:"EndPoint"`
}

func (r *ModifyEndPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEndPointRequest struct {
	*ksyunhttp.BaseRequest
	EndPointName *string                   `json:"EndPointName,omitempty" name:"EndPointName"`
	Region       *string                   `json:"Region,omitempty" name:"Region"`
	VpcId        *string                   `json:"VpcId,omitempty" name:"VpcId"`
	Description  *string                   `json:"Description,omitempty" name:"Description"`
	IpConfig     []*CreateEndPointIpConfig `json:"IpConfig,omitempty" name:"IpConfig"`
}

func (r *CreateEndPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEndPointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	EndPoint  struct {
		EndPointId      *string   `json:"EndPointId" name:"EndPointId"`
		Region          *string   `json:"Region" name:"Region"`
		EndPointName    *string   `json:"EndPointName" name:"EndPointName"`
		Type            *string   `json:"Type" name:"Type"`
		VpcId           *string   `json:"VpcId" name:"VpcId"`
		SecurityGroupId *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		FdZoneIds       []*string `json:"FdZoneIds" name:"FdZoneIds"`
		CreateTime      *string   `json:"CreateTime" name:"CreateTime"`
		Status          *string   `json:"Status" name:"Status"`
		Description     *string   `json:"Description" name:"Description"`
		IpConfigSet     []struct {
			AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			SubnetId         *string `json:"SubnetId" name:"SubnetId"`
			Ip               *string `json:"Ip" name:"Ip"`
		} `json:"IpConfigSet" name:"IpConfigSet"`
	} `json:"EndPoint"`
}

func (r *CreateEndPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
