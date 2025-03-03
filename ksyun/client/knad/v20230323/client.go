package v20230323

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2023-03-23"

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

func NewCreateKnadRequest() (request *CreateKnadRequest) {
	request = &CreateKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "CreateKnad")
	return
}

func NewCreateKnadResponse() (response *CreateKnadResponse) {
	response = &CreateKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateKnad(request *CreateKnadRequest) string {
	return c.CreateKnadWithContext(context.Background(), request)
}

func (c *Client) CreateKnadWithContext(ctx context.Context, request *CreateKnadRequest) string {
	if request == nil {
		request = NewCreateKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyKnadRequest() (request *ModifyKnadRequest) {
	request = &ModifyKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyKnad")
	return
}

func NewModifyKnadResponse() (response *ModifyKnadResponse) {
	response = &ModifyKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyKnad(request *ModifyKnadRequest) string {
	return c.ModifyKnadWithContext(context.Background(), request)
}

func (c *Client) ModifyKnadWithContext(ctx context.Context, request *ModifyKnadRequest) string {
	if request == nil {
		request = NewModifyKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewKnadIpListRequest() (request *KnadIpListRequest) {
	request = &KnadIpListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "KnadIpList")
	return
}

func NewKnadIpListResponse() (response *KnadIpListResponse) {
	response = &KnadIpListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) KnadIpList(request *KnadIpListRequest) string {
	return c.KnadIpListWithContext(context.Background(), request)
}

func (c *Client) KnadIpListWithContext(ctx context.Context, request *KnadIpListRequest) string {
	if request == nil {
		request = NewKnadIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewKnadIpListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAssociateIpRequest() (request *AssociateIpRequest) {
	request = &AssociateIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "AssociateIp")
	return
}

func NewAssociateIpResponse() (response *AssociateIpResponse) {
	response = &AssociateIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateIp(request *AssociateIpRequest) string {
	return c.AssociateIpWithContext(context.Background(), request)
}

func (c *Client) AssociateIpWithContext(ctx context.Context, request *AssociateIpRequest) string {
	if request == nil {
		request = NewAssociateIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisassociateIpRequest() (request *DisassociateIpRequest) {
	request = &DisassociateIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DisassociateIp")
	return
}

func NewDisassociateIpResponse() (response *DisassociateIpResponse) {
	response = &DisassociateIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateIp(request *DisassociateIpRequest) string {
	return c.DisassociateIpWithContext(context.Background(), request)
}

func (c *Client) DisassociateIpWithContext(ctx context.Context, request *DisassociateIpRequest) string {
	if request == nil {
		request = NewDisassociateIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeKnadIpRequest() (request *DescribeKnadIpRequest) {
	request = &DescribeKnadIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DescribeKnadIp")
	return
}

func NewDescribeKnadIpResponse() (response *DescribeKnadIpResponse) {
	response = &DescribeKnadIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnadIp(request *DescribeKnadIpRequest) string {
	return c.DescribeKnadIpWithContext(context.Background(), request)
}

func (c *Client) DescribeKnadIpWithContext(ctx context.Context, request *DescribeKnadIpRequest) string {
	if request == nil {
		request = NewDescribeKnadIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnadIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteKnadRequest() (request *DeleteKnadRequest) {
	request = &DeleteKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DeleteKnad")
	return
}

func NewDeleteKnadResponse() (response *DeleteKnadResponse) {
	response = &DeleteKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteKnad(request *DeleteKnadRequest) string {
	return c.DeleteKnadWithContext(context.Background(), request)
}

func (c *Client) DeleteKnadWithContext(ctx context.Context, request *DeleteKnadRequest) string {
	if request == nil {
		request = NewDeleteKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeKnadRequest() (request *DescribeKnadRequest) {
	request = &DescribeKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DescribeKnad")
	return
}

func NewDescribeKnadResponse() (response *DescribeKnadResponse) {
	response = &DescribeKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnad(request *DescribeKnadRequest) string {
	return c.DescribeKnadWithContext(context.Background(), request)
}

func (c *Client) DescribeKnadWithContext(ctx context.Context, request *DescribeKnadRequest) string {
	if request == nil {
		request = NewDescribeKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyKnadNameRequest() (request *ModifyKnadNameRequest) {
	request = &ModifyKnadNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyKnadName")
	return
}

func NewModifyKnadNameResponse() (response *ModifyKnadNameResponse) {
	response = &ModifyKnadNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyKnadName(request *ModifyKnadNameRequest) string {
	return c.ModifyKnadNameWithContext(context.Background(), request)
}

func (c *Client) ModifyKnadNameWithContext(ctx context.Context, request *ModifyKnadNameRequest) string {
	if request == nil {
		request = NewModifyKnadNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKnadNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetBWIpListRequest() (request *GetBWIpListRequest) {
	request = &GetBWIpListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetBWIpList")
	return
}

func NewGetBWIpListResponse() (response *GetBWIpListResponse) {
	response = &GetBWIpListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBWIpList(request *GetBWIpListRequest) string {
	return c.GetBWIpListWithContext(context.Background(), request)
}

func (c *Client) GetBWIpListWithContext(ctx context.Context, request *GetBWIpListRequest) string {
	if request == nil {
		request = NewGetBWIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBWIpListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteBWRequest() (request *DeleteBWRequest) {
	request = &DeleteBWRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DeleteBW")
	return
}

func NewDeleteBWResponse() (response *DeleteBWResponse) {
	response = &DeleteBWResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBW(request *DeleteBWRequest) string {
	return c.DeleteBWWithContext(context.Background(), request)
}

func (c *Client) DeleteBWWithContext(ctx context.Context, request *DeleteBWRequest) string {
	if request == nil {
		request = NewDeleteBWRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBWResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAddBWIpListRequest() (request *AddBWIpListRequest) {
	request = &AddBWIpListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "AddBWIpList")
	return
}

func NewAddBWIpListResponse() (response *AddBWIpListResponse) {
	response = &AddBWIpListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddBWIpList(request *AddBWIpListRequest) string {
	return c.AddBWIpListWithContext(context.Background(), request)
}

func (c *Client) AddBWIpListWithContext(ctx context.Context, request *AddBWIpListRequest) string {
	if request == nil {
		request = NewAddBWIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddBWIpListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetZoneListRequest() (request *GetZoneListRequest) {
	request = &GetZoneListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetZoneList")
	return
}

func NewGetZoneListResponse() (response *GetZoneListResponse) {
	response = &GetZoneListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetZoneList(request *GetZoneListRequest) string {
	return c.GetZoneListWithContext(context.Background(), request)
}

func (c *Client) GetZoneListWithContext(ctx context.Context, request *GetZoneListRequest) string {
	if request == nil {
		request = NewGetZoneListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetZoneListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyPolicyRequest() (request *ModifyPolicyRequest) {
	request = &ModifyPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyPolicy")
	return
}

func NewModifyPolicyResponse() (response *ModifyPolicyResponse) {
	response = &ModifyPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPolicy(request *ModifyPolicyRequest) string {
	return c.ModifyPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyPolicyWithContext(ctx context.Context, request *ModifyPolicyRequest) string {
	if request == nil {
		request = NewModifyPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyBlockLocationRequest() (request *ModifyBlockLocationRequest) {
	request = &ModifyBlockLocationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyBlockLocation")
	return
}

func NewModifyBlockLocationResponse() (response *ModifyBlockLocationResponse) {
	response = &ModifyBlockLocationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBlockLocation(request *ModifyBlockLocationRequest) string {
	return c.ModifyBlockLocationWithContext(context.Background(), request)
}

func (c *Client) ModifyBlockLocationWithContext(ctx context.Context, request *ModifyBlockLocationRequest) string {
	if request == nil {
		request = NewModifyBlockLocationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBlockLocationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetBlockLocationsRequest() (request *GetBlockLocationsRequest) {
	request = &GetBlockLocationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetBlockLocations")
	return
}

func NewGetBlockLocationsResponse() (response *GetBlockLocationsResponse) {
	response = &GetBlockLocationsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBlockLocations(request *GetBlockLocationsRequest) string {
	return c.GetBlockLocationsWithContext(context.Background(), request)
}

func (c *Client) GetBlockLocationsWithContext(ctx context.Context, request *GetBlockLocationsRequest) string {
	if request == nil {
		request = NewGetBlockLocationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBlockLocationsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetAttackLogRequest() (request *GetAttackLogRequest) {
	request = &GetAttackLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetAttackLog")
	return
}

func NewGetAttackLogResponse() (response *GetAttackLogResponse) {
	response = &GetAttackLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetAttackLog(request *GetAttackLogRequest) string {
	return c.GetAttackLogWithContext(context.Background(), request)
}

func (c *Client) GetAttackLogWithContext(ctx context.Context, request *GetAttackLogRequest) string {
	if request == nil {
		request = NewGetAttackLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetAttackLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetKnadPolicyRequest() (request *GetKnadPolicyRequest) {
	request = &GetKnadPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetKnadPolicy")
	return
}

func NewGetKnadPolicyResponse() (response *GetKnadPolicyResponse) {
	response = &GetKnadPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetKnadPolicy(request *GetKnadPolicyRequest) string {
	return c.GetKnadPolicyWithContext(context.Background(), request)
}

func (c *Client) GetKnadPolicyWithContext(ctx context.Context, request *GetKnadPolicyRequest) string {
	if request == nil {
		request = NewGetKnadPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetKnadPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
