package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2016-03-04"

type Client struct {
	common.Client
}

func NewClient(credential common.Credentials, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewCreatePrivateDnsRequest() (request *CreatePrivateDnsRequest) {
	request = &CreatePrivateDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreatePrivateDns")
	return
}

func NewCreatePrivateDnsResponse() (response *CreatePrivateDnsResponse) {
	response = &CreatePrivateDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePrivateDns(request *CreatePrivateDnsRequest) string {
	return c.CreatePrivateDnsWithContext(context.Background(), request)
}

func (c *Client) CreatePrivateDnsWithContext(ctx context.Context, request *CreatePrivateDnsRequest) string {
	if request == nil {
		request = NewCreatePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrivateDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeletePrivateDnsRequest() (request *DeletePrivateDnsRequest) {
	request = &DeletePrivateDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeletePrivateDns")
	return
}

func NewDeletePrivateDnsResponse() (response *DeletePrivateDnsResponse) {
	response = &DeletePrivateDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePrivateDns(request *DeletePrivateDnsRequest) string {
	return c.DeletePrivateDnsWithContext(context.Background(), request)
}

func (c *Client) DeletePrivateDnsWithContext(ctx context.Context, request *DeletePrivateDnsRequest) string {
	if request == nil {
		request = NewDeletePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribePrivateDnsRequest() (request *DescribePrivateDnsRequest) {
	request = &DescribePrivateDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribePrivateDns")
	return
}

func NewDescribePrivateDnsResponse() (response *DescribePrivateDnsResponse) {
	response = &DescribePrivateDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrivateDns(request *DescribePrivateDnsRequest) string {
	return c.DescribePrivateDnsWithContext(context.Background(), request)
}

func (c *Client) DescribePrivateDnsWithContext(ctx context.Context, request *DescribePrivateDnsRequest) string {
	if request == nil {
		request = NewDescribePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAssociateVpcsRequest() (request *AssociateVpcsRequest) {
	request = &AssociateVpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "AssociateVpcs")
	return
}

func NewAssociateVpcsResponse() (response *AssociateVpcsResponse) {
	response = &AssociateVpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateVpcs(request *AssociateVpcsRequest) string {
	return c.AssociateVpcsWithContext(context.Background(), request)
}

func (c *Client) AssociateVpcsWithContext(ctx context.Context, request *AssociateVpcsRequest) string {
	if request == nil {
		request = NewAssociateVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateVpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisassociateVpcsRequest() (request *DisassociateVpcsRequest) {
	request = &DisassociateVpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DisassociateVpcs")
	return
}

func NewDisassociateVpcsResponse() (response *DisassociateVpcsResponse) {
	response = &DisassociateVpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateVpcs(request *DisassociateVpcsRequest) string {
	return c.DisassociateVpcsWithContext(context.Background(), request)
}

func (c *Client) DisassociateVpcsWithContext(ctx context.Context, request *DisassociateVpcsRequest) string {
	if request == nil {
		request = NewDisassociateVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateVpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateZoneRequest() (request *CreateZoneRequest) {
	request = &CreateZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateZone")
	return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
	response = &CreateZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateZone(request *CreateZoneRequest) string {
	return c.CreateZoneWithContext(context.Background(), request)
}

func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) string {
	if request == nil {
		request = NewCreateZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
	request = &DeleteZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteZone")
	return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
	response = &DeleteZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteZone(request *DeleteZoneRequest) string {
	return c.DeleteZoneWithContext(context.Background(), request)
}

func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) string {
	if request == nil {
		request = NewDeleteZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyZoneRequest() (request *ModifyZoneRequest) {
	request = &ModifyZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyZone")
	return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
	response = &ModifyZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyZone(request *ModifyZoneRequest) string {
	return c.ModifyZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) string {
	if request == nil {
		request = NewModifyZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeZoneRequest() (request *DescribeZoneRequest) {
	request = &DescribeZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeZone")
	return
}

func NewDescribeZoneResponse() (response *DescribeZoneResponse) {
	response = &DescribeZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeZone(request *DescribeZoneRequest) string {
	return c.DescribeZoneWithContext(context.Background(), request)
}

func (c *Client) DescribeZoneWithContext(ctx context.Context, request *DescribeZoneRequest) string {
	if request == nil {
		request = NewDescribeZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateRecordRequest() (request *CreateRecordRequest) {
	request = &CreateRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateRecord")
	return
}

func NewCreateRecordResponse() (response *CreateRecordResponse) {
	response = &CreateRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRecord(request *CreateRecordRequest) string {
	return c.CreateRecordWithContext(context.Background(), request)
}

func (c *Client) CreateRecordWithContext(ctx context.Context, request *CreateRecordRequest) string {
	if request == nil {
		request = NewCreateRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
	request = &DeleteRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteRecord")
	return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
	response = &DeleteRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecord(request *DeleteRecordRequest) string {
	return c.DeleteRecordWithContext(context.Background(), request)
}

func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) string {
	if request == nil {
		request = NewDeleteRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyRecordRequest() (request *ModifyRecordRequest) {
	request = &ModifyRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyRecord")
	return
}

func NewModifyRecordResponse() (response *ModifyRecordResponse) {
	response = &ModifyRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRecord(request *ModifyRecordRequest) string {
	return c.ModifyRecordWithContext(context.Background(), request)
}

func (c *Client) ModifyRecordWithContext(ctx context.Context, request *ModifyRecordRequest) string {
	if request == nil {
		request = NewModifyRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
	request = &DescribeRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeRecord")
	return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
	response = &DescribeRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRecord(request *DescribeRecordRequest) string {
	return c.DescribeRecordWithContext(context.Background(), request)
}

func (c *Client) DescribeRecordWithContext(ctx context.Context, request *DescribeRecordRequest) string {
	if request == nil {
		request = NewDescribeRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateRecordDataRequest() (request *CreateRecordDataRequest) {
	request = &CreateRecordDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateRecordData")
	return
}

func NewCreateRecordDataResponse() (response *CreateRecordDataResponse) {
	response = &CreateRecordDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRecordData(request *CreateRecordDataRequest) string {
	return c.CreateRecordDataWithContext(context.Background(), request)
}

func (c *Client) CreateRecordDataWithContext(ctx context.Context, request *CreateRecordDataRequest) string {
	if request == nil {
		request = NewCreateRecordDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRecordDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteRecordDataRequest() (request *DeleteRecordDataRequest) {
	request = &DeleteRecordDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteRecordData")
	return
}

func NewDeleteRecordDataResponse() (response *DeleteRecordDataResponse) {
	response = &DeleteRecordDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecordData(request *DeleteRecordDataRequest) string {
	return c.DeleteRecordDataWithContext(context.Background(), request)
}

func (c *Client) DeleteRecordDataWithContext(ctx context.Context, request *DeleteRecordDataRequest) string {
	if request == nil {
		request = NewDeleteRecordDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecordDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreatePdnsZoneRequest() (request *CreatePdnsZoneRequest) {
	request = &CreatePdnsZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreatePdnsZone")
	return
}

func NewCreatePdnsZoneResponse() (response *CreatePdnsZoneResponse) {
	response = &CreatePdnsZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePdnsZone(request *CreatePdnsZoneRequest) string {
	return c.CreatePdnsZoneWithContext(context.Background(), request)
}

func (c *Client) CreatePdnsZoneWithContext(ctx context.Context, request *CreatePdnsZoneRequest) string {
	if request == nil {
		request = NewCreatePdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePdnsZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyPdnsZoneRequest() (request *ModifyPdnsZoneRequest) {
	request = &ModifyPdnsZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyPdnsZone")
	return
}

func NewModifyPdnsZoneResponse() (response *ModifyPdnsZoneResponse) {
	response = &ModifyPdnsZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPdnsZone(request *ModifyPdnsZoneRequest) string {
	return c.ModifyPdnsZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyPdnsZoneWithContext(ctx context.Context, request *ModifyPdnsZoneRequest) string {
	if request == nil {
		request = NewModifyPdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPdnsZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeletePdnsZoneRequest() (request *DeletePdnsZoneRequest) {
	request = &DeletePdnsZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeletePdnsZone")
	return
}

func NewDeletePdnsZoneResponse() (response *DeletePdnsZoneResponse) {
	response = &DeletePdnsZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePdnsZone(request *DeletePdnsZoneRequest) string {
	return c.DeletePdnsZoneWithContext(context.Background(), request)
}

func (c *Client) DeletePdnsZoneWithContext(ctx context.Context, request *DeletePdnsZoneRequest) string {
	if request == nil {
		request = NewDeletePdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePdnsZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribePdnsZonesRequest() (request *DescribePdnsZonesRequest) {
	request = &DescribePdnsZonesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribePdnsZones")
	return
}

func NewDescribePdnsZonesResponse() (response *DescribePdnsZonesResponse) {
	response = &DescribePdnsZonesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePdnsZones(request *DescribePdnsZonesRequest) string {
	return c.DescribePdnsZonesWithContext(context.Background(), request)
}

func (c *Client) DescribePdnsZonesWithContext(ctx context.Context, request *DescribePdnsZonesRequest) string {
	if request == nil {
		request = NewDescribePdnsZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePdnsZonesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewBindZoneVpcRequest() (request *BindZoneVpcRequest) {
	request = &BindZoneVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "BindZoneVpc")
	return
}

func NewBindZoneVpcResponse() (response *BindZoneVpcResponse) {
	response = &BindZoneVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindZoneVpc(request *BindZoneVpcRequest) string {
	return c.BindZoneVpcWithContext(context.Background(), request)
}

func (c *Client) BindZoneVpcWithContext(ctx context.Context, request *BindZoneVpcRequest) string {
	if request == nil {
		request = NewBindZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBindZoneVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUnbindZoneVpcRequest() (request *UnbindZoneVpcRequest) {
	request = &UnbindZoneVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "UnbindZoneVpc")
	return
}

func NewUnbindZoneVpcResponse() (response *UnbindZoneVpcResponse) {
	response = &UnbindZoneVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindZoneVpc(request *UnbindZoneVpcRequest) string {
	return c.UnbindZoneVpcWithContext(context.Background(), request)
}

func (c *Client) UnbindZoneVpcWithContext(ctx context.Context, request *UnbindZoneVpcRequest) string {
	if request == nil {
		request = NewUnbindZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnbindZoneVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateZoneRecordRequest() (request *CreateZoneRecordRequest) {
	request = &CreateZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateZoneRecord")
	return
}

func NewCreateZoneRecordResponse() (response *CreateZoneRecordResponse) {
	response = &CreateZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateZoneRecord(request *CreateZoneRecordRequest) string {
	return c.CreateZoneRecordWithContext(context.Background(), request)
}

func (c *Client) CreateZoneRecordWithContext(ctx context.Context, request *CreateZoneRecordRequest) string {
	if request == nil {
		request = NewCreateZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteZoneRecordRequest() (request *DeleteZoneRecordRequest) {
	request = &DeleteZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteZoneRecord")
	return
}

func NewDeleteZoneRecordResponse() (response *DeleteZoneRecordResponse) {
	response = &DeleteZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteZoneRecord(request *DeleteZoneRecordRequest) string {
	return c.DeleteZoneRecordWithContext(context.Background(), request)
}

func (c *Client) DeleteZoneRecordWithContext(ctx context.Context, request *DeleteZoneRecordRequest) string {
	if request == nil {
		request = NewDeleteZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyZoneRecordRequest() (request *ModifyZoneRecordRequest) {
	request = &ModifyZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyZoneRecord")
	return
}

func NewModifyZoneRecordResponse() (response *ModifyZoneRecordResponse) {
	response = &ModifyZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyZoneRecord(request *ModifyZoneRecordRequest) string {
	return c.ModifyZoneRecordWithContext(context.Background(), request)
}

func (c *Client) ModifyZoneRecordWithContext(ctx context.Context, request *ModifyZoneRecordRequest) string {
	if request == nil {
		request = NewModifyZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeZoneRecordRequest() (request *DescribeZoneRecordRequest) {
	request = &DescribeZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeZoneRecord")
	return
}

func NewDescribeZoneRecordResponse() (response *DescribeZoneRecordResponse) {
	response = &DescribeZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeZoneRecord(request *DescribeZoneRecordRequest) string {
	return c.DescribeZoneRecordWithContext(context.Background(), request)
}

func (c *Client) DescribeZoneRecordWithContext(ctx context.Context, request *DescribeZoneRecordRequest) string {
	if request == nil {
		request = NewDescribeZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
