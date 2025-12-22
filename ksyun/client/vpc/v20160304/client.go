package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func NewCreateVpcRequest() (request *CreateVpcRequest) {
	request = &CreateVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpc")
	return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
	response = &CreateVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpc(request *CreateVpcRequest) string {
	return c.CreateVpcWithContext(context.Background(), request)
}

func (c *Client) CreateVpcSend(request *CreateVpcRequest) (*CreateVpcResponse, error) {
	statusCode, msg, err := c.CreateVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateVpcWithContext(ctx context.Context, request *CreateVpcRequest) string {
	if request == nil {
		request = NewCreateVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateVpcWithContextV2(ctx context.Context, request *CreateVpcRequest) (int, string, error) {
	if request == nil {
		request = NewCreateVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
	request = &DeleteVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpc")
	return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
	response = &DeleteVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpc(request *DeleteVpcRequest) string {
	return c.DeleteVpcWithContext(context.Background(), request)
}

func (c *Client) DeleteVpcSend(request *DeleteVpcRequest) (*DeleteVpcResponse, error) {
	statusCode, msg, err := c.DeleteVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteVpcWithContext(ctx context.Context, request *DeleteVpcRequest) string {
	if request == nil {
		request = NewDeleteVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteVpcWithContextV2(ctx context.Context, request *DeleteVpcRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcs")
	return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) string {
	return c.DescribeVpcsWithContext(context.Background(), request)
}

func (c *Client) DescribeVpcsSend(request *DescribeVpcsRequest) (*DescribeVpcsResponse, error) {
	statusCode, msg, err := c.DescribeVpcsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeVpcsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest) string {
	if request == nil {
		request = NewDescribeVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpcsWithContextV2(ctx context.Context, request *DescribeVpcsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpcsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
	request = &CreateSubnetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSubnet")
	return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
	response = &CreateSubnetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSubnet(request *CreateSubnetRequest) string {
	return c.CreateSubnetWithContext(context.Background(), request)
}

func (c *Client) CreateSubnetSend(request *CreateSubnetRequest) (*CreateSubnetResponse, error) {
	statusCode, msg, err := c.CreateSubnetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSubnetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest) string {
	if request == nil {
		request = NewCreateSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSubnetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSubnetWithContextV2(ctx context.Context, request *CreateSubnetRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSubnetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
	request = &DeleteSubnetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteSubnet")
	return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
	response = &DeleteSubnetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) string {
	return c.DeleteSubnetWithContext(context.Background(), request)
}

func (c *Client) DeleteSubnetSend(request *DeleteSubnetRequest) (*DeleteSubnetResponse, error) {
	statusCode, msg, err := c.DeleteSubnetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSubnetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest) string {
	if request == nil {
		request = NewDeleteSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSubnetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSubnetWithContextV2(ctx context.Context, request *DeleteSubnetRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSubnetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
	request = &DescribeSubnetsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnets")
	return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
	response = &DescribeSubnetsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) string {
	return c.DescribeSubnetsWithContext(context.Background(), request)
}

func (c *Client) DescribeSubnetsSend(request *DescribeSubnetsRequest) (*DescribeSubnetsResponse, error) {
	statusCode, msg, err := c.DescribeSubnetsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSubnetsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSubnetsWithContext(ctx context.Context, request *DescribeSubnetsRequest) string {
	if request == nil {
		request = NewDescribeSubnetsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubnetsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSubnetsWithContextV2(ctx context.Context, request *DescribeSubnetsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSubnetsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubnetsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateNetworkAclRequest() (request *AssociateNetworkAclRequest) {
	request = &AssociateNetworkAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateNetworkAcl")
	return
}

func NewAssociateNetworkAclResponse() (response *AssociateNetworkAclResponse) {
	response = &AssociateNetworkAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateNetworkAcl(request *AssociateNetworkAclRequest) string {
	return c.AssociateNetworkAclWithContext(context.Background(), request)
}

func (c *Client) AssociateNetworkAclSend(request *AssociateNetworkAclRequest) (*AssociateNetworkAclResponse, error) {
	statusCode, msg, err := c.AssociateNetworkAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateNetworkAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateNetworkAclWithContext(ctx context.Context, request *AssociateNetworkAclRequest) string {
	if request == nil {
		request = NewAssociateNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateNetworkAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateNetworkAclWithContextV2(ctx context.Context, request *AssociateNetworkAclRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateNetworkAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateNetworkAclRequest() (request *DisassociateNetworkAclRequest) {
	request = &DisassociateNetworkAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNetworkAcl")
	return
}

func NewDisassociateNetworkAclResponse() (response *DisassociateNetworkAclResponse) {
	response = &DisassociateNetworkAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateNetworkAcl(request *DisassociateNetworkAclRequest) string {
	return c.DisassociateNetworkAclWithContext(context.Background(), request)
}

func (c *Client) DisassociateNetworkAclSend(request *DisassociateNetworkAclRequest) (*DisassociateNetworkAclResponse, error) {
	statusCode, msg, err := c.DisassociateNetworkAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisassociateNetworkAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateNetworkAclWithContext(ctx context.Context, request *DisassociateNetworkAclRequest) string {
	if request == nil {
		request = NewDisassociateNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateNetworkAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateNetworkAclWithContextV2(ctx context.Context, request *DisassociateNetworkAclRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateNetworkAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateRoute")
	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRoute(request *CreateRouteRequest) string {
	return c.CreateRouteWithContext(context.Background(), request)
}

func (c *Client) CreateRouteSend(request *CreateRouteRequest) (*CreateRouteResponse, error) {
	statusCode, msg, err := c.CreateRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRouteWithContext(ctx context.Context, request *CreateRouteRequest) string {
	if request == nil {
		request = NewCreateRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRouteWithContextV2(ctx context.Context, request *CreateRouteRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
	request = &DeleteRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoute")
	return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
	response = &DeleteRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRoute(request *DeleteRouteRequest) string {
	return c.DeleteRouteWithContext(context.Background(), request)
}

func (c *Client) DeleteRouteSend(request *DeleteRouteRequest) (*DeleteRouteResponse, error) {
	statusCode, msg, err := c.DeleteRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRouteWithContext(ctx context.Context, request *DeleteRouteRequest) string {
	if request == nil {
		request = NewDeleteRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRouteWithContextV2(ctx context.Context, request *DeleteRouteRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRoutesRequest() (request *DescribeRoutesRequest) {
	request = &DescribeRoutesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRoutes")
	return
}

func NewDescribeRoutesResponse() (response *DescribeRoutesResponse) {
	response = &DescribeRoutesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRoutes(request *DescribeRoutesRequest) string {
	return c.DescribeRoutesWithContext(context.Background(), request)
}

func (c *Client) DescribeRoutesSend(request *DescribeRoutesRequest) (*DescribeRoutesResponse, error) {
	statusCode, msg, err := c.DescribeRoutesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeRoutesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRoutesWithContext(ctx context.Context, request *DescribeRoutesRequest) string {
	if request == nil {
		request = NewDescribeRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRoutesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRoutesWithContextV2(ctx context.Context, request *DescribeRoutesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRoutesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNetworkAclRequest() (request *CreateNetworkAclRequest) {
	request = &CreateNetworkAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkAcl")
	return
}

func NewCreateNetworkAclResponse() (response *CreateNetworkAclResponse) {
	response = &CreateNetworkAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNetworkAcl(request *CreateNetworkAclRequest) string {
	return c.CreateNetworkAclWithContext(context.Background(), request)
}

func (c *Client) CreateNetworkAclSend(request *CreateNetworkAclRequest) (*CreateNetworkAclResponse, error) {
	statusCode, msg, err := c.CreateNetworkAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateNetworkAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNetworkAclWithContext(ctx context.Context, request *CreateNetworkAclRequest) string {
	if request == nil {
		request = NewCreateNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNetworkAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNetworkAclWithContextV2(ctx context.Context, request *CreateNetworkAclRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNetworkAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNetworkAclRequest() (request *DeleteNetworkAclRequest) {
	request = &DeleteNetworkAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkAcl")
	return
}

func NewDeleteNetworkAclResponse() (response *DeleteNetworkAclResponse) {
	response = &DeleteNetworkAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNetworkAcl(request *DeleteNetworkAclRequest) string {
	return c.DeleteNetworkAclWithContext(context.Background(), request)
}

func (c *Client) DeleteNetworkAclSend(request *DeleteNetworkAclRequest) (*DeleteNetworkAclResponse, error) {
	statusCode, msg, err := c.DeleteNetworkAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNetworkAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNetworkAclWithContext(ctx context.Context, request *DeleteNetworkAclRequest) string {
	if request == nil {
		request = NewDeleteNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNetworkAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNetworkAclWithContextV2(ctx context.Context, request *DeleteNetworkAclRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNetworkAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNetworkAclEntryRequest() (request *CreateNetworkAclEntryRequest) {
	request = &CreateNetworkAclEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkAclEntry")
	return
}

func NewCreateNetworkAclEntryResponse() (response *CreateNetworkAclEntryResponse) {
	response = &CreateNetworkAclEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNetworkAclEntry(request *CreateNetworkAclEntryRequest) string {
	return c.CreateNetworkAclEntryWithContext(context.Background(), request)
}

func (c *Client) CreateNetworkAclEntrySend(request *CreateNetworkAclEntryRequest) (*CreateNetworkAclEntryResponse, error) {
	statusCode, msg, err := c.CreateNetworkAclEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateNetworkAclEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNetworkAclEntryWithContext(ctx context.Context, request *CreateNetworkAclEntryRequest) string {
	if request == nil {
		request = NewCreateNetworkAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNetworkAclEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNetworkAclEntryWithContextV2(ctx context.Context, request *CreateNetworkAclEntryRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNetworkAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNetworkAclEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNetworkAclEntryRequest() (request *DeleteNetworkAclEntryRequest) {
	request = &DeleteNetworkAclEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkAclEntry")
	return
}

func NewDeleteNetworkAclEntryResponse() (response *DeleteNetworkAclEntryResponse) {
	response = &DeleteNetworkAclEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNetworkAclEntry(request *DeleteNetworkAclEntryRequest) string {
	return c.DeleteNetworkAclEntryWithContext(context.Background(), request)
}

func (c *Client) DeleteNetworkAclEntrySend(request *DeleteNetworkAclEntryRequest) (*DeleteNetworkAclEntryResponse, error) {
	statusCode, msg, err := c.DeleteNetworkAclEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNetworkAclEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNetworkAclEntryWithContext(ctx context.Context, request *DeleteNetworkAclEntryRequest) string {
	if request == nil {
		request = NewDeleteNetworkAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNetworkAclEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNetworkAclEntryWithContextV2(ctx context.Context, request *DeleteNetworkAclEntryRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNetworkAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNetworkAclEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNetworkAclsRequest() (request *DescribeNetworkAclsRequest) {
	request = &DescribeNetworkAclsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkAcls")
	return
}

func NewDescribeNetworkAclsResponse() (response *DescribeNetworkAclsResponse) {
	response = &DescribeNetworkAclsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetworkAcls(request *DescribeNetworkAclsRequest) string {
	return c.DescribeNetworkAclsWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkAclsSend(request *DescribeNetworkAclsRequest) (*DescribeNetworkAclsResponse, error) {
	statusCode, msg, err := c.DescribeNetworkAclsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNetworkAclsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNetworkAclsWithContext(ctx context.Context, request *DescribeNetworkAclsRequest) string {
	if request == nil {
		request = NewDescribeNetworkAclsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkAclsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNetworkAclsWithContextV2(ctx context.Context, request *DescribeNetworkAclsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNetworkAclsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkAclsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroup")
	return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) string {
	return c.CreateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupSend(request *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	statusCode, msg, err := c.CreateSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) string {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityGroupWithContextV2(ctx context.Context, request *CreateSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroup")
	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
	response = &DeleteSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) string {
	return c.DeleteSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupSend(request *DeleteSecurityGroupRequest) (*DeleteSecurityGroupResponse, error) {
	statusCode, msg, err := c.DeleteSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) string {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityGroupWithContextV2(ctx context.Context, request *DeleteSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAuthorizeSecurityGroupEntryRequest() (request *AuthorizeSecurityGroupEntryRequest) {
	request = &AuthorizeSecurityGroupEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AuthorizeSecurityGroupEntry")
	return
}

func NewAuthorizeSecurityGroupEntryResponse() (response *AuthorizeSecurityGroupEntryResponse) {
	response = &AuthorizeSecurityGroupEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AuthorizeSecurityGroupEntry(request *AuthorizeSecurityGroupEntryRequest) string {
	return c.AuthorizeSecurityGroupEntryWithContext(context.Background(), request)
}

func (c *Client) AuthorizeSecurityGroupEntrySend(request *AuthorizeSecurityGroupEntryRequest) (*AuthorizeSecurityGroupEntryResponse, error) {
	statusCode, msg, err := c.AuthorizeSecurityGroupEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AuthorizeSecurityGroupEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AuthorizeSecurityGroupEntryWithContext(ctx context.Context, request *AuthorizeSecurityGroupEntryRequest) string {
	if request == nil {
		request = NewAuthorizeSecurityGroupEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAuthorizeSecurityGroupEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AuthorizeSecurityGroupEntryWithContextV2(ctx context.Context, request *AuthorizeSecurityGroupEntryRequest) (int, string, error) {
	if request == nil {
		request = NewAuthorizeSecurityGroupEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAuthorizeSecurityGroupEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRevokeSecurityGroupEntryRequest() (request *RevokeSecurityGroupEntryRequest) {
	request = &RevokeSecurityGroupEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RevokeSecurityGroupEntry")
	return
}

func NewRevokeSecurityGroupEntryResponse() (response *RevokeSecurityGroupEntryResponse) {
	response = &RevokeSecurityGroupEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RevokeSecurityGroupEntry(request *RevokeSecurityGroupEntryRequest) string {
	return c.RevokeSecurityGroupEntryWithContext(context.Background(), request)
}

func (c *Client) RevokeSecurityGroupEntrySend(request *RevokeSecurityGroupEntryRequest) (*RevokeSecurityGroupEntryResponse, error) {
	statusCode, msg, err := c.RevokeSecurityGroupEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RevokeSecurityGroupEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RevokeSecurityGroupEntryWithContext(ctx context.Context, request *RevokeSecurityGroupEntryRequest) string {
	if request == nil {
		request = NewRevokeSecurityGroupEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRevokeSecurityGroupEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RevokeSecurityGroupEntryWithContextV2(ctx context.Context, request *RevokeSecurityGroupEntryRequest) (int, string, error) {
	if request == nil {
		request = NewRevokeSecurityGroupEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRevokeSecurityGroupEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
	request = &DescribeSecurityGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroups")
	return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
	response = &DescribeSecurityGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) string {
	return c.DescribeSecurityGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupsSend(request *DescribeSecurityGroupsRequest) (*DescribeSecurityGroupsResponse, error) {
	statusCode, msg, err := c.DescribeSecurityGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSecurityGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest) string {
	if request == nil {
		request = NewDescribeSecurityGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSecurityGroupsWithContextV2(ctx context.Context, request *DescribeSecurityGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSecurityGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNatRequest() (request *CreateNatRequest) {
	request = &CreateNatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNat")
	return
}

func NewCreateNatResponse() (response *CreateNatResponse) {
	response = &CreateNatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNat(request *CreateNatRequest) string {
	return c.CreateNatWithContext(context.Background(), request)
}

func (c *Client) CreateNatSend(request *CreateNatRequest) (*CreateNatResponse, error) {
	statusCode, msg, err := c.CreateNatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateNatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNatWithContext(ctx context.Context, request *CreateNatRequest) string {
	if request == nil {
		request = NewCreateNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNatWithContextV2(ctx context.Context, request *CreateNatRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNatRequest() (request *DeleteNatRequest) {
	request = &DeleteNatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNat")
	return
}

func NewDeleteNatResponse() (response *DeleteNatResponse) {
	response = &DeleteNatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNat(request *DeleteNatRequest) string {
	return c.DeleteNatWithContext(context.Background(), request)
}

func (c *Client) DeleteNatSend(request *DeleteNatRequest) (*DeleteNatResponse, error) {
	statusCode, msg, err := c.DeleteNatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNatWithContext(ctx context.Context, request *DeleteNatRequest) string {
	if request == nil {
		request = NewDeleteNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNatWithContextV2(ctx context.Context, request *DeleteNatRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNatsRequest() (request *DescribeNatsRequest) {
	request = &DescribeNatsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNats")
	return
}

func NewDescribeNatsResponse() (response *DescribeNatsResponse) {
	response = &DescribeNatsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNats(request *DescribeNatsRequest) string {
	return c.DescribeNatsWithContext(context.Background(), request)
}

func (c *Client) DescribeNatsSend(request *DescribeNatsRequest) (*DescribeNatsResponse, error) {
	statusCode, msg, err := c.DescribeNatsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNatsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNatsWithContext(ctx context.Context, request *DescribeNatsRequest) string {
	if request == nil {
		request = NewDescribeNatsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNatsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNatsWithContextV2(ctx context.Context, request *DescribeNatsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNatsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNatsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateNatRequest() (request *AssociateNatRequest) {
	request = &AssociateNatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateNat")
	return
}

func NewAssociateNatResponse() (response *AssociateNatResponse) {
	response = &AssociateNatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateNat(request *AssociateNatRequest) string {
	return c.AssociateNatWithContext(context.Background(), request)
}

func (c *Client) AssociateNatSend(request *AssociateNatRequest) (*AssociateNatResponse, error) {
	statusCode, msg, err := c.AssociateNatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateNatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateNatWithContext(ctx context.Context, request *AssociateNatRequest) string {
	if request == nil {
		request = NewAssociateNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateNatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateNatWithContextV2(ctx context.Context, request *AssociateNatRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateNatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateNatRequest() (request *DisassociateNatRequest) {
	request = &DisassociateNatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNat")
	return
}

func NewDisassociateNatResponse() (response *DisassociateNatResponse) {
	response = &DisassociateNatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateNat(request *DisassociateNatRequest) string {
	return c.DisassociateNatWithContext(context.Background(), request)
}

func (c *Client) DisassociateNatSend(request *DisassociateNatRequest) (*DisassociateNatResponse, error) {
	statusCode, msg, err := c.DisassociateNatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisassociateNatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateNatWithContext(ctx context.Context, request *DisassociateNatRequest) string {
	if request == nil {
		request = NewDisassociateNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateNatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateNatWithContextV2(ctx context.Context, request *DisassociateNatRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateNatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInternetGatewaysRequest() (request *DescribeInternetGatewaysRequest) {
	request = &DescribeInternetGatewaysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeInternetGateways")
	return
}

func NewDescribeInternetGatewaysResponse() (response *DescribeInternetGatewaysResponse) {
	response = &DescribeInternetGatewaysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInternetGateways(request *DescribeInternetGatewaysRequest) string {
	return c.DescribeInternetGatewaysWithContext(context.Background(), request)
}

func (c *Client) DescribeInternetGatewaysSend(request *DescribeInternetGatewaysRequest) (*DescribeInternetGatewaysResponse, error) {
	statusCode, msg, err := c.DescribeInternetGatewaysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInternetGatewaysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInternetGatewaysWithContext(ctx context.Context, request *DescribeInternetGatewaysRequest) string {
	if request == nil {
		request = NewDescribeInternetGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInternetGatewaysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInternetGatewaysWithContextV2(ctx context.Context, request *DescribeInternetGatewaysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInternetGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInternetGatewaysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateVpcPeeringConnectionRequest() (request *CreateVpcPeeringConnectionRequest) {
	request = &CreateVpcPeeringConnectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcPeeringConnection")
	return
}

func NewCreateVpcPeeringConnectionResponse() (response *CreateVpcPeeringConnectionResponse) {
	response = &CreateVpcPeeringConnectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpcPeeringConnection(request *CreateVpcPeeringConnectionRequest) string {
	return c.CreateVpcPeeringConnectionWithContext(context.Background(), request)
}

func (c *Client) CreateVpcPeeringConnectionSend(request *CreateVpcPeeringConnectionRequest) (*CreateVpcPeeringConnectionResponse, error) {
	statusCode, msg, err := c.CreateVpcPeeringConnectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateVpcPeeringConnectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateVpcPeeringConnectionWithContext(ctx context.Context, request *CreateVpcPeeringConnectionRequest) string {
	if request == nil {
		request = NewCreateVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpcPeeringConnectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateVpcPeeringConnectionWithContextV2(ctx context.Context, request *CreateVpcPeeringConnectionRequest) (int, string, error) {
	if request == nil {
		request = NewCreateVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpcPeeringConnectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteVpcPeeringConnectionRequest() (request *DeleteVpcPeeringConnectionRequest) {
	request = &DeleteVpcPeeringConnectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcPeeringConnection")
	return
}

func NewDeleteVpcPeeringConnectionResponse() (response *DeleteVpcPeeringConnectionResponse) {
	response = &DeleteVpcPeeringConnectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpcPeeringConnection(request *DeleteVpcPeeringConnectionRequest) string {
	return c.DeleteVpcPeeringConnectionWithContext(context.Background(), request)
}

func (c *Client) DeleteVpcPeeringConnectionSend(request *DeleteVpcPeeringConnectionRequest) (*DeleteVpcPeeringConnectionResponse, error) {
	statusCode, msg, err := c.DeleteVpcPeeringConnectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteVpcPeeringConnectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteVpcPeeringConnectionWithContext(ctx context.Context, request *DeleteVpcPeeringConnectionRequest) string {
	if request == nil {
		request = NewDeleteVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpcPeeringConnectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteVpcPeeringConnectionWithContextV2(ctx context.Context, request *DeleteVpcPeeringConnectionRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpcPeeringConnectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpcPeeringConnectionsRequest() (request *DescribeVpcPeeringConnectionsRequest) {
	request = &DescribeVpcPeeringConnectionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcPeeringConnections")
	return
}

func NewDescribeVpcPeeringConnectionsResponse() (response *DescribeVpcPeeringConnectionsResponse) {
	response = &DescribeVpcPeeringConnectionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcPeeringConnections(request *DescribeVpcPeeringConnectionsRequest) string {
	return c.DescribeVpcPeeringConnectionsWithContext(context.Background(), request)
}

func (c *Client) DescribeVpcPeeringConnectionsSend(request *DescribeVpcPeeringConnectionsRequest) (*DescribeVpcPeeringConnectionsResponse, error) {
	statusCode, msg, err := c.DescribeVpcPeeringConnectionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeVpcPeeringConnectionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpcPeeringConnectionsWithContext(ctx context.Context, request *DescribeVpcPeeringConnectionsRequest) string {
	if request == nil {
		request = NewDescribeVpcPeeringConnectionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpcPeeringConnectionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpcPeeringConnectionsWithContextV2(ctx context.Context, request *DescribeVpcPeeringConnectionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpcPeeringConnectionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpcPeeringConnectionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNetworkAclRequest() (request *ModifyNetworkAclRequest) {
	request = &ModifyNetworkAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkAcl")
	return
}

func NewModifyNetworkAclResponse() (response *ModifyNetworkAclResponse) {
	response = &ModifyNetworkAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNetworkAcl(request *ModifyNetworkAclRequest) string {
	return c.ModifyNetworkAclWithContext(context.Background(), request)
}

func (c *Client) ModifyNetworkAclSend(request *ModifyNetworkAclRequest) (*ModifyNetworkAclResponse, error) {
	statusCode, msg, err := c.ModifyNetworkAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyNetworkAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNetworkAclWithContext(ctx context.Context, request *ModifyNetworkAclRequest) string {
	if request == nil {
		request = NewModifyNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNetworkAclWithContextV2(ctx context.Context, request *ModifyNetworkAclRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNetworkAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
	request = &ModifySecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroup")
	return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
	response = &ModifySecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) string {
	return c.ModifySecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupSend(request *ModifySecurityGroupRequest) (*ModifySecurityGroupResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupWithContext(ctx context.Context, request *ModifySecurityGroupRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupWithContextV2(ctx context.Context, request *ModifySecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySubnetRequest() (request *ModifySubnetRequest) {
	request = &ModifySubnetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifySubnet")
	return
}

func NewModifySubnetResponse() (response *ModifySubnetResponse) {
	response = &ModifySubnetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySubnet(request *ModifySubnetRequest) string {
	return c.ModifySubnetWithContext(context.Background(), request)
}

func (c *Client) ModifySubnetSend(request *ModifySubnetRequest) (*ModifySubnetResponse, error) {
	statusCode, msg, err := c.ModifySubnetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySubnetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySubnetWithContext(ctx context.Context, request *ModifySubnetRequest) string {
	if request == nil {
		request = NewModifySubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySubnetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySubnetWithContextV2(ctx context.Context, request *ModifySubnetRequest) (int, string, error) {
	if request == nil {
		request = NewModifySubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySubnetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNatRequest() (request *ModifyNatRequest) {
	request = &ModifyNatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNat")
	return
}

func NewModifyNatResponse() (response *ModifyNatResponse) {
	response = &ModifyNatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNat(request *ModifyNatRequest) string {
	return c.ModifyNatWithContext(context.Background(), request)
}

func (c *Client) ModifyNatSend(request *ModifyNatRequest) (*ModifyNatResponse, error) {
	statusCode, msg, err := c.ModifyNatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyNatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNatWithContext(ctx context.Context, request *ModifyNatRequest) string {
	if request == nil {
		request = NewModifyNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNatWithContextV2(ctx context.Context, request *ModifyNatRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
	request = &DescribeNetworkInterfacesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaces")
	return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
	response = &DescribeNetworkInterfacesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) string {
	return c.DescribeNetworkInterfacesWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkInterfacesSend(request *DescribeNetworkInterfacesRequest) (*DescribeNetworkInterfacesResponse, error) {
	statusCode, msg, err := c.DescribeNetworkInterfacesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNetworkInterfacesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNetworkInterfacesWithContext(ctx context.Context, request *DescribeNetworkInterfacesRequest) string {
	if request == nil {
		request = NewDescribeNetworkInterfacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkInterfacesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNetworkInterfacesWithContextV2(ctx context.Context, request *DescribeNetworkInterfacesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNetworkInterfacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkInterfacesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSubnetAvailableAddressesRequest() (request *DescribeSubnetAvailableAddressesRequest) {
	request = &DescribeSubnetAvailableAddressesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnetAvailableAddresses")
	return
}

func NewDescribeSubnetAvailableAddressesResponse() (response *DescribeSubnetAvailableAddressesResponse) {
	response = &DescribeSubnetAvailableAddressesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubnetAvailableAddresses(request *DescribeSubnetAvailableAddressesRequest) string {
	return c.DescribeSubnetAvailableAddressesWithContext(context.Background(), request)
}

func (c *Client) DescribeSubnetAvailableAddressesSend(request *DescribeSubnetAvailableAddressesRequest) (*DescribeSubnetAvailableAddressesResponse, error) {
	statusCode, msg, err := c.DescribeSubnetAvailableAddressesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSubnetAvailableAddressesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSubnetAvailableAddressesWithContext(ctx context.Context, request *DescribeSubnetAvailableAddressesRequest) string {
	if request == nil {
		request = NewDescribeSubnetAvailableAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubnetAvailableAddressesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSubnetAvailableAddressesWithContextV2(ctx context.Context, request *DescribeSubnetAvailableAddressesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSubnetAvailableAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubnetAvailableAddressesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVpcRequest() (request *ModifyVpcRequest) {
	request = &ModifyVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpc")
	return
}

func NewModifyVpcResponse() (response *ModifyVpcResponse) {
	response = &ModifyVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpc(request *ModifyVpcRequest) string {
	return c.ModifyVpcWithContext(context.Background(), request)
}

func (c *Client) ModifyVpcSend(request *ModifyVpcRequest) (*ModifyVpcResponse, error) {
	statusCode, msg, err := c.ModifyVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVpcWithContext(ctx context.Context, request *ModifyVpcRequest) string {
	if request == nil {
		request = NewModifyVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVpcWithContextV2(ctx context.Context, request *ModifyVpcRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAcceptVpcPeeringConnectionRequest() (request *AcceptVpcPeeringConnectionRequest) {
	request = &AcceptVpcPeeringConnectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AcceptVpcPeeringConnection")
	return
}

func NewAcceptVpcPeeringConnectionResponse() (response *AcceptVpcPeeringConnectionResponse) {
	response = &AcceptVpcPeeringConnectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AcceptVpcPeeringConnection(request *AcceptVpcPeeringConnectionRequest) string {
	return c.AcceptVpcPeeringConnectionWithContext(context.Background(), request)
}

func (c *Client) AcceptVpcPeeringConnectionSend(request *AcceptVpcPeeringConnectionRequest) (*AcceptVpcPeeringConnectionResponse, error) {
	statusCode, msg, err := c.AcceptVpcPeeringConnectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AcceptVpcPeeringConnectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AcceptVpcPeeringConnectionWithContext(ctx context.Context, request *AcceptVpcPeeringConnectionRequest) string {
	if request == nil {
		request = NewAcceptVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAcceptVpcPeeringConnectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AcceptVpcPeeringConnectionWithContextV2(ctx context.Context, request *AcceptVpcPeeringConnectionRequest) (int, string, error) {
	if request == nil {
		request = NewAcceptVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAcceptVpcPeeringConnectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRejectVpcPeeringConnectionRequest() (request *RejectVpcPeeringConnectionRequest) {
	request = &RejectVpcPeeringConnectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RejectVpcPeeringConnection")
	return
}

func NewRejectVpcPeeringConnectionResponse() (response *RejectVpcPeeringConnectionResponse) {
	response = &RejectVpcPeeringConnectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RejectVpcPeeringConnection(request *RejectVpcPeeringConnectionRequest) string {
	return c.RejectVpcPeeringConnectionWithContext(context.Background(), request)
}

func (c *Client) RejectVpcPeeringConnectionSend(request *RejectVpcPeeringConnectionRequest) (*RejectVpcPeeringConnectionResponse, error) {
	statusCode, msg, err := c.RejectVpcPeeringConnectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RejectVpcPeeringConnectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RejectVpcPeeringConnectionWithContext(ctx context.Context, request *RejectVpcPeeringConnectionRequest) string {
	if request == nil {
		request = NewRejectVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRejectVpcPeeringConnectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RejectVpcPeeringConnectionWithContextV2(ctx context.Context, request *RejectVpcPeeringConnectionRequest) (int, string, error) {
	if request == nil {
		request = NewRejectVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRejectVpcPeeringConnectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVpcPeeringConnectionRequest() (request *ModifyVpcPeeringConnectionRequest) {
	request = &ModifyVpcPeeringConnectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcPeeringConnection")
	return
}

func NewModifyVpcPeeringConnectionResponse() (response *ModifyVpcPeeringConnectionResponse) {
	response = &ModifyVpcPeeringConnectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpcPeeringConnection(request *ModifyVpcPeeringConnectionRequest) string {
	return c.ModifyVpcPeeringConnectionWithContext(context.Background(), request)
}

func (c *Client) ModifyVpcPeeringConnectionSend(request *ModifyVpcPeeringConnectionRequest) (*ModifyVpcPeeringConnectionResponse, error) {
	statusCode, msg, err := c.ModifyVpcPeeringConnectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyVpcPeeringConnectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVpcPeeringConnectionWithContext(ctx context.Context, request *ModifyVpcPeeringConnectionRequest) string {
	if request == nil {
		request = NewModifyVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpcPeeringConnectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVpcPeeringConnectionWithContextV2(ctx context.Context, request *ModifyVpcPeeringConnectionRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVpcPeeringConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpcPeeringConnectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAvailabilityZonesRequest() (request *DescribeAvailabilityZonesRequest) {
	request = &DescribeAvailabilityZonesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAvailabilityZones")
	return
}

func NewDescribeAvailabilityZonesResponse() (response *DescribeAvailabilityZonesResponse) {
	response = &DescribeAvailabilityZonesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAvailabilityZones(request *DescribeAvailabilityZonesRequest) string {
	return c.DescribeAvailabilityZonesWithContext(context.Background(), request)
}

func (c *Client) DescribeAvailabilityZonesSend(request *DescribeAvailabilityZonesRequest) (*DescribeAvailabilityZonesResponse, error) {
	statusCode, msg, err := c.DescribeAvailabilityZonesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAvailabilityZonesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, request *DescribeAvailabilityZonesRequest) string {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAvailabilityZonesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAvailabilityZonesWithContextV2(ctx context.Context, request *DescribeAvailabilityZonesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAvailabilityZonesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirectConnectsRequest() (request *DescribeDirectConnectsRequest) {
	request = &DescribeDirectConnectsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnects")
	return
}

func NewDescribeDirectConnectsResponse() (response *DescribeDirectConnectsResponse) {
	response = &DescribeDirectConnectsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirectConnects(request *DescribeDirectConnectsRequest) string {
	return c.DescribeDirectConnectsWithContext(context.Background(), request)
}

func (c *Client) DescribeDirectConnectsSend(request *DescribeDirectConnectsRequest) (*DescribeDirectConnectsResponse, error) {
	statusCode, msg, err := c.DescribeDirectConnectsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirectConnectsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirectConnectsWithContext(ctx context.Context, request *DescribeDirectConnectsRequest) string {
	if request == nil {
		request = NewDescribeDirectConnectsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirectConnectsWithContextV2(ctx context.Context, request *DescribeDirectConnectsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirectConnectsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDirectConnectInterfaceRequest() (request *CreateDirectConnectInterfaceRequest) {
	request = &CreateDirectConnectInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectInterface")
	return
}

func NewCreateDirectConnectInterfaceResponse() (response *CreateDirectConnectInterfaceResponse) {
	response = &CreateDirectConnectInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDirectConnectInterface(request *CreateDirectConnectInterfaceRequest) string {
	return c.CreateDirectConnectInterfaceWithContext(context.Background(), request)
}

func (c *Client) CreateDirectConnectInterfaceSend(request *CreateDirectConnectInterfaceRequest) (*CreateDirectConnectInterfaceResponse, error) {
	statusCode, msg, err := c.CreateDirectConnectInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDirectConnectInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDirectConnectInterfaceWithContext(ctx context.Context, request *CreateDirectConnectInterfaceRequest) string {
	if request == nil {
		request = NewCreateDirectConnectInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDirectConnectInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDirectConnectInterfaceWithContextV2(ctx context.Context, request *CreateDirectConnectInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDirectConnectInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDirectConnectInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDirectConnectInterfaceRequest() (request *DeleteDirectConnectInterfaceRequest) {
	request = &DeleteDirectConnectInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectInterface")
	return
}

func NewDeleteDirectConnectInterfaceResponse() (response *DeleteDirectConnectInterfaceResponse) {
	response = &DeleteDirectConnectInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDirectConnectInterface(request *DeleteDirectConnectInterfaceRequest) string {
	return c.DeleteDirectConnectInterfaceWithContext(context.Background(), request)
}

func (c *Client) DeleteDirectConnectInterfaceSend(request *DeleteDirectConnectInterfaceRequest) (*DeleteDirectConnectInterfaceResponse, error) {
	statusCode, msg, err := c.DeleteDirectConnectInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDirectConnectInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDirectConnectInterfaceWithContext(ctx context.Context, request *DeleteDirectConnectInterfaceRequest) string {
	if request == nil {
		request = NewDeleteDirectConnectInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDirectConnectInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDirectConnectInterfaceWithContextV2(ctx context.Context, request *DeleteDirectConnectInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDirectConnectInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDirectConnectInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirectConnectInterfacesRequest() (request *DescribeDirectConnectInterfacesRequest) {
	request = &DescribeDirectConnectInterfacesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectInterfaces")
	return
}

func NewDescribeDirectConnectInterfacesResponse() (response *DescribeDirectConnectInterfacesResponse) {
	response = &DescribeDirectConnectInterfacesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirectConnectInterfaces(request *DescribeDirectConnectInterfacesRequest) string {
	return c.DescribeDirectConnectInterfacesWithContext(context.Background(), request)
}

func (c *Client) DescribeDirectConnectInterfacesSend(request *DescribeDirectConnectInterfacesRequest) (*DescribeDirectConnectInterfacesResponse, error) {
	statusCode, msg, err := c.DescribeDirectConnectInterfacesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirectConnectInterfacesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirectConnectInterfacesWithContext(ctx context.Context, request *DescribeDirectConnectInterfacesRequest) string {
	if request == nil {
		request = NewDescribeDirectConnectInterfacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectInterfacesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirectConnectInterfacesWithContextV2(ctx context.Context, request *DescribeDirectConnectInterfacesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirectConnectInterfacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectInterfacesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDirectConnectGatewayRequest() (request *CreateDirectConnectGatewayRequest) {
	request = &CreateDirectConnectGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectGateway")
	return
}

func NewCreateDirectConnectGatewayResponse() (response *CreateDirectConnectGatewayResponse) {
	response = &CreateDirectConnectGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDirectConnectGateway(request *CreateDirectConnectGatewayRequest) string {
	return c.CreateDirectConnectGatewayWithContext(context.Background(), request)
}

func (c *Client) CreateDirectConnectGatewaySend(request *CreateDirectConnectGatewayRequest) (*CreateDirectConnectGatewayResponse, error) {
	statusCode, msg, err := c.CreateDirectConnectGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDirectConnectGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDirectConnectGatewayWithContext(ctx context.Context, request *CreateDirectConnectGatewayRequest) string {
	if request == nil {
		request = NewCreateDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDirectConnectGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDirectConnectGatewayWithContextV2(ctx context.Context, request *CreateDirectConnectGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDirectConnectGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDirectConnectGatewayRequest() (request *DeleteDirectConnectGatewayRequest) {
	request = &DeleteDirectConnectGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectGateway")
	return
}

func NewDeleteDirectConnectGatewayResponse() (response *DeleteDirectConnectGatewayResponse) {
	response = &DeleteDirectConnectGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDirectConnectGateway(request *DeleteDirectConnectGatewayRequest) string {
	return c.DeleteDirectConnectGatewayWithContext(context.Background(), request)
}

func (c *Client) DeleteDirectConnectGatewaySend(request *DeleteDirectConnectGatewayRequest) (*DeleteDirectConnectGatewayResponse, error) {
	statusCode, msg, err := c.DeleteDirectConnectGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDirectConnectGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDirectConnectGatewayWithContext(ctx context.Context, request *DeleteDirectConnectGatewayRequest) string {
	if request == nil {
		request = NewDeleteDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDirectConnectGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDirectConnectGatewayWithContextV2(ctx context.Context, request *DeleteDirectConnectGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDirectConnectGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirectConnectGatewaysRequest() (request *DescribeDirectConnectGatewaysRequest) {
	request = &DescribeDirectConnectGatewaysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectGateways")
	return
}

func NewDescribeDirectConnectGatewaysResponse() (response *DescribeDirectConnectGatewaysResponse) {
	response = &DescribeDirectConnectGatewaysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirectConnectGateways(request *DescribeDirectConnectGatewaysRequest) string {
	return c.DescribeDirectConnectGatewaysWithContext(context.Background(), request)
}

func (c *Client) DescribeDirectConnectGatewaysSend(request *DescribeDirectConnectGatewaysRequest) (*DescribeDirectConnectGatewaysResponse, error) {
	statusCode, msg, err := c.DescribeDirectConnectGatewaysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirectConnectGatewaysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirectConnectGatewaysWithContext(ctx context.Context, request *DescribeDirectConnectGatewaysRequest) string {
	if request == nil {
		request = NewDescribeDirectConnectGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectGatewaysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirectConnectGatewaysWithContextV2(ctx context.Context, request *DescribeDirectConnectGatewaysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirectConnectGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectGatewaysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachDirectConnectGatewayRequest() (request *AttachDirectConnectGatewayRequest) {
	request = &AttachDirectConnectGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AttachDirectConnectGateway")
	return
}

func NewAttachDirectConnectGatewayResponse() (response *AttachDirectConnectGatewayResponse) {
	response = &AttachDirectConnectGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachDirectConnectGateway(request *AttachDirectConnectGatewayRequest) string {
	return c.AttachDirectConnectGatewayWithContext(context.Background(), request)
}

func (c *Client) AttachDirectConnectGatewaySend(request *AttachDirectConnectGatewayRequest) (*AttachDirectConnectGatewayResponse, error) {
	statusCode, msg, err := c.AttachDirectConnectGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AttachDirectConnectGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachDirectConnectGatewayWithContext(ctx context.Context, request *AttachDirectConnectGatewayRequest) string {
	if request == nil {
		request = NewAttachDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachDirectConnectGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachDirectConnectGatewayWithContextV2(ctx context.Context, request *AttachDirectConnectGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewAttachDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachDirectConnectGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachDirectConnectGatewayRequest() (request *DetachDirectConnectGatewayRequest) {
	request = &DetachDirectConnectGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DetachDirectConnectGateway")
	return
}

func NewDetachDirectConnectGatewayResponse() (response *DetachDirectConnectGatewayResponse) {
	response = &DetachDirectConnectGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachDirectConnectGateway(request *DetachDirectConnectGatewayRequest) string {
	return c.DetachDirectConnectGatewayWithContext(context.Background(), request)
}

func (c *Client) DetachDirectConnectGatewaySend(request *DetachDirectConnectGatewayRequest) (*DetachDirectConnectGatewayResponse, error) {
	statusCode, msg, err := c.DetachDirectConnectGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DetachDirectConnectGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachDirectConnectGatewayWithContext(ctx context.Context, request *DetachDirectConnectGatewayRequest) string {
	if request == nil {
		request = NewDetachDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachDirectConnectGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachDirectConnectGatewayWithContextV2(ctx context.Context, request *DetachDirectConnectGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewDetachDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachDirectConnectGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDirectConnectInterfaceRequest() (request *ModifyDirectConnectInterfaceRequest) {
	request = &ModifyDirectConnectInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyDirectConnectInterface")
	return
}

func NewModifyDirectConnectInterfaceResponse() (response *ModifyDirectConnectInterfaceResponse) {
	response = &ModifyDirectConnectInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDirectConnectInterface(request *ModifyDirectConnectInterfaceRequest) string {
	return c.ModifyDirectConnectInterfaceWithContext(context.Background(), request)
}

func (c *Client) ModifyDirectConnectInterfaceSend(request *ModifyDirectConnectInterfaceRequest) (*ModifyDirectConnectInterfaceResponse, error) {
	statusCode, msg, err := c.ModifyDirectConnectInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDirectConnectInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDirectConnectInterfaceWithContext(ctx context.Context, request *ModifyDirectConnectInterfaceRequest) string {
	if request == nil {
		request = NewModifyDirectConnectInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDirectConnectInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDirectConnectInterfaceWithContextV2(ctx context.Context, request *ModifyDirectConnectInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDirectConnectInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDirectConnectInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDirectConnectGatewayRequest() (request *ModifyDirectConnectGatewayRequest) {
	request = &ModifyDirectConnectGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyDirectConnectGateway")
	return
}

func NewModifyDirectConnectGatewayResponse() (response *ModifyDirectConnectGatewayResponse) {
	response = &ModifyDirectConnectGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDirectConnectGateway(request *ModifyDirectConnectGatewayRequest) string {
	return c.ModifyDirectConnectGatewayWithContext(context.Background(), request)
}

func (c *Client) ModifyDirectConnectGatewaySend(request *ModifyDirectConnectGatewayRequest) (*ModifyDirectConnectGatewayResponse, error) {
	statusCode, msg, err := c.ModifyDirectConnectGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDirectConnectGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDirectConnectGatewayWithContext(ctx context.Context, request *ModifyDirectConnectGatewayRequest) string {
	if request == nil {
		request = NewModifyDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDirectConnectGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDirectConnectGatewayWithContextV2(ctx context.Context, request *ModifyDirectConnectGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDirectConnectGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDirectConnectGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateVpnGatewayRequest() (request *CreateVpnGatewayRequest) {
	request = &CreateVpnGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnGateway")
	return
}

func NewCreateVpnGatewayResponse() (response *CreateVpnGatewayResponse) {
	response = &CreateVpnGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpnGateway(request *CreateVpnGatewayRequest) string {
	return c.CreateVpnGatewayWithContext(context.Background(), request)
}

func (c *Client) CreateVpnGatewaySend(request *CreateVpnGatewayRequest) (*CreateVpnGatewayResponse, error) {
	statusCode, msg, err := c.CreateVpnGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateVpnGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateVpnGatewayWithContext(ctx context.Context, request *CreateVpnGatewayRequest) string {
	if request == nil {
		request = NewCreateVpnGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpnGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateVpnGatewayWithContextV2(ctx context.Context, request *CreateVpnGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewCreateVpnGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpnGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVpnGatewayRequest() (request *ModifyVpnGatewayRequest) {
	request = &ModifyVpnGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGateway")
	return
}

func NewModifyVpnGatewayResponse() (response *ModifyVpnGatewayResponse) {
	response = &ModifyVpnGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpnGateway(request *ModifyVpnGatewayRequest) string {
	return c.ModifyVpnGatewayWithContext(context.Background(), request)
}

func (c *Client) ModifyVpnGatewaySend(request *ModifyVpnGatewayRequest) (*ModifyVpnGatewayResponse, error) {
	statusCode, msg, err := c.ModifyVpnGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyVpnGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVpnGatewayWithContext(ctx context.Context, request *ModifyVpnGatewayRequest) string {
	if request == nil {
		request = NewModifyVpnGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpnGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVpnGatewayWithContextV2(ctx context.Context, request *ModifyVpnGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVpnGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpnGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteVpnGatewayRequest() (request *DeleteVpnGatewayRequest) {
	request = &DeleteVpnGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGateway")
	return
}

func NewDeleteVpnGatewayResponse() (response *DeleteVpnGatewayResponse) {
	response = &DeleteVpnGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpnGateway(request *DeleteVpnGatewayRequest) string {
	return c.DeleteVpnGatewayWithContext(context.Background(), request)
}

func (c *Client) DeleteVpnGatewaySend(request *DeleteVpnGatewayRequest) (*DeleteVpnGatewayResponse, error) {
	statusCode, msg, err := c.DeleteVpnGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteVpnGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteVpnGatewayWithContext(ctx context.Context, request *DeleteVpnGatewayRequest) string {
	if request == nil {
		request = NewDeleteVpnGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpnGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteVpnGatewayWithContextV2(ctx context.Context, request *DeleteVpnGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteVpnGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpnGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
	request = &DescribeVpnGatewaysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGateways")
	return
}

func NewDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
	response = &DescribeVpnGatewaysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) string {
	return c.DescribeVpnGatewaysWithContext(context.Background(), request)
}

func (c *Client) DescribeVpnGatewaysSend(request *DescribeVpnGatewaysRequest) (*DescribeVpnGatewaysResponse, error) {
	statusCode, msg, err := c.DescribeVpnGatewaysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeVpnGatewaysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpnGatewaysWithContext(ctx context.Context, request *DescribeVpnGatewaysRequest) string {
	if request == nil {
		request = NewDescribeVpnGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnGatewaysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpnGatewaysWithContextV2(ctx context.Context, request *DescribeVpnGatewaysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpnGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnGatewaysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateVpnTunnelRequest() (request *CreateVpnTunnelRequest) {
	request = &CreateVpnTunnelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnTunnel")
	return
}

func NewCreateVpnTunnelResponse() (response *CreateVpnTunnelResponse) {
	response = &CreateVpnTunnelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpnTunnel(request *CreateVpnTunnelRequest) string {
	return c.CreateVpnTunnelWithContext(context.Background(), request)
}

func (c *Client) CreateVpnTunnelSend(request *CreateVpnTunnelRequest) (*CreateVpnTunnelResponse, error) {
	statusCode, msg, err := c.CreateVpnTunnelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateVpnTunnelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateVpnTunnelWithContext(ctx context.Context, request *CreateVpnTunnelRequest) string {
	if request == nil {
		request = NewCreateVpnTunnelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpnTunnelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateVpnTunnelWithContextV2(ctx context.Context, request *CreateVpnTunnelRequest) (int, string, error) {
	if request == nil {
		request = NewCreateVpnTunnelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpnTunnelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVpnTunnelRequest() (request *ModifyVpnTunnelRequest) {
	request = &ModifyVpnTunnelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnTunnel")
	return
}

func NewModifyVpnTunnelResponse() (response *ModifyVpnTunnelResponse) {
	response = &ModifyVpnTunnelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpnTunnel(request *ModifyVpnTunnelRequest) string {
	return c.ModifyVpnTunnelWithContext(context.Background(), request)
}

func (c *Client) ModifyVpnTunnelSend(request *ModifyVpnTunnelRequest) (*ModifyVpnTunnelResponse, error) {
	statusCode, msg, err := c.ModifyVpnTunnelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyVpnTunnelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVpnTunnelWithContext(ctx context.Context, request *ModifyVpnTunnelRequest) string {
	if request == nil {
		request = NewModifyVpnTunnelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpnTunnelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVpnTunnelWithContextV2(ctx context.Context, request *ModifyVpnTunnelRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVpnTunnelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpnTunnelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteVpnTunnelRequest() (request *DeleteVpnTunnelRequest) {
	request = &DeleteVpnTunnelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnTunnel")
	return
}

func NewDeleteVpnTunnelResponse() (response *DeleteVpnTunnelResponse) {
	response = &DeleteVpnTunnelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpnTunnel(request *DeleteVpnTunnelRequest) string {
	return c.DeleteVpnTunnelWithContext(context.Background(), request)
}

func (c *Client) DeleteVpnTunnelSend(request *DeleteVpnTunnelRequest) (*DeleteVpnTunnelResponse, error) {
	statusCode, msg, err := c.DeleteVpnTunnelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteVpnTunnelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteVpnTunnelWithContext(ctx context.Context, request *DeleteVpnTunnelRequest) string {
	if request == nil {
		request = NewDeleteVpnTunnelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpnTunnelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteVpnTunnelWithContextV2(ctx context.Context, request *DeleteVpnTunnelRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteVpnTunnelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpnTunnelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpnTunnelsRequest() (request *DescribeVpnTunnelsRequest) {
	request = &DescribeVpnTunnelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnTunnels")
	return
}

func NewDescribeVpnTunnelsResponse() (response *DescribeVpnTunnelsResponse) {
	response = &DescribeVpnTunnelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpnTunnels(request *DescribeVpnTunnelsRequest) string {
	return c.DescribeVpnTunnelsWithContext(context.Background(), request)
}

func (c *Client) DescribeVpnTunnelsSend(request *DescribeVpnTunnelsRequest) (*DescribeVpnTunnelsResponse, error) {
	statusCode, msg, err := c.DescribeVpnTunnelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeVpnTunnelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpnTunnelsWithContext(ctx context.Context, request *DescribeVpnTunnelsRequest) string {
	if request == nil {
		request = NewDescribeVpnTunnelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnTunnelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpnTunnelsWithContextV2(ctx context.Context, request *DescribeVpnTunnelsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpnTunnelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnTunnelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCustomerGatewayRequest() (request *CreateCustomerGatewayRequest) {
	request = &CreateCustomerGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateCustomerGateway")
	return
}

func NewCreateCustomerGatewayResponse() (response *CreateCustomerGatewayResponse) {
	response = &CreateCustomerGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCustomerGateway(request *CreateCustomerGatewayRequest) string {
	return c.CreateCustomerGatewayWithContext(context.Background(), request)
}

func (c *Client) CreateCustomerGatewaySend(request *CreateCustomerGatewayRequest) (*CreateCustomerGatewayResponse, error) {
	statusCode, msg, err := c.CreateCustomerGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCustomerGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCustomerGatewayWithContext(ctx context.Context, request *CreateCustomerGatewayRequest) string {
	if request == nil {
		request = NewCreateCustomerGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCustomerGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCustomerGatewayWithContextV2(ctx context.Context, request *CreateCustomerGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCustomerGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCustomerGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCustomerGatewayRequest() (request *ModifyCustomerGatewayRequest) {
	request = &ModifyCustomerGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyCustomerGateway")
	return
}

func NewModifyCustomerGatewayResponse() (response *ModifyCustomerGatewayResponse) {
	response = &ModifyCustomerGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCustomerGateway(request *ModifyCustomerGatewayRequest) string {
	return c.ModifyCustomerGatewayWithContext(context.Background(), request)
}

func (c *Client) ModifyCustomerGatewaySend(request *ModifyCustomerGatewayRequest) (*ModifyCustomerGatewayResponse, error) {
	statusCode, msg, err := c.ModifyCustomerGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCustomerGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCustomerGatewayWithContext(ctx context.Context, request *ModifyCustomerGatewayRequest) string {
	if request == nil {
		request = NewModifyCustomerGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCustomerGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCustomerGatewayWithContextV2(ctx context.Context, request *ModifyCustomerGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCustomerGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCustomerGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCustomerGatewayRequest() (request *DeleteCustomerGatewayRequest) {
	request = &DeleteCustomerGatewayRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteCustomerGateway")
	return
}

func NewDeleteCustomerGatewayResponse() (response *DeleteCustomerGatewayResponse) {
	response = &DeleteCustomerGatewayResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCustomerGateway(request *DeleteCustomerGatewayRequest) string {
	return c.DeleteCustomerGatewayWithContext(context.Background(), request)
}

func (c *Client) DeleteCustomerGatewaySend(request *DeleteCustomerGatewayRequest) (*DeleteCustomerGatewayResponse, error) {
	statusCode, msg, err := c.DeleteCustomerGatewayWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCustomerGatewayResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCustomerGatewayWithContext(ctx context.Context, request *DeleteCustomerGatewayRequest) string {
	if request == nil {
		request = NewDeleteCustomerGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCustomerGatewayResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCustomerGatewayWithContextV2(ctx context.Context, request *DeleteCustomerGatewayRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCustomerGatewayRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCustomerGatewayResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDirectConnectRequest() (request *ModifyDirectConnectRequest) {
	request = &ModifyDirectConnectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyDirectConnect")
	return
}

func NewModifyDirectConnectResponse() (response *ModifyDirectConnectResponse) {
	response = &ModifyDirectConnectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDirectConnect(request *ModifyDirectConnectRequest) string {
	return c.ModifyDirectConnectWithContext(context.Background(), request)
}

func (c *Client) ModifyDirectConnectSend(request *ModifyDirectConnectRequest) (*ModifyDirectConnectResponse, error) {
	statusCode, msg, err := c.ModifyDirectConnectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDirectConnectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDirectConnectWithContext(ctx context.Context, request *ModifyDirectConnectRequest) string {
	if request == nil {
		request = NewModifyDirectConnectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDirectConnectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDirectConnectWithContextV2(ctx context.Context, request *ModifyDirectConnectRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDirectConnectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDirectConnectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCustomerGatewaysRequest() (request *DescribeCustomerGatewaysRequest) {
	request = &DescribeCustomerGatewaysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCustomerGateways")
	return
}

func NewDescribeCustomerGatewaysResponse() (response *DescribeCustomerGatewaysResponse) {
	response = &DescribeCustomerGatewaysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) string {
	return c.DescribeCustomerGatewaysWithContext(context.Background(), request)
}

func (c *Client) DescribeCustomerGatewaysSend(request *DescribeCustomerGatewaysRequest) (*DescribeCustomerGatewaysResponse, error) {
	statusCode, msg, err := c.DescribeCustomerGatewaysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCustomerGatewaysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCustomerGatewaysWithContext(ctx context.Context, request *DescribeCustomerGatewaysRequest) string {
	if request == nil {
		request = NewDescribeCustomerGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCustomerGatewaysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCustomerGatewaysWithContextV2(ctx context.Context, request *DescribeCustomerGatewaysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCustomerGatewaysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCustomerGatewaysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSubnetAllocatedIpAddressesRequest() (request *DescribeSubnetAllocatedIpAddressesRequest) {
	request = &DescribeSubnetAllocatedIpAddressesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnetAllocatedIpAddresses")
	return
}

func NewDescribeSubnetAllocatedIpAddressesResponse() (response *DescribeSubnetAllocatedIpAddressesResponse) {
	response = &DescribeSubnetAllocatedIpAddressesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubnetAllocatedIpAddresses(request *DescribeSubnetAllocatedIpAddressesRequest) string {
	return c.DescribeSubnetAllocatedIpAddressesWithContext(context.Background(), request)
}

func (c *Client) DescribeSubnetAllocatedIpAddressesSend(request *DescribeSubnetAllocatedIpAddressesRequest) (*DescribeSubnetAllocatedIpAddressesResponse, error) {
	statusCode, msg, err := c.DescribeSubnetAllocatedIpAddressesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSubnetAllocatedIpAddressesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSubnetAllocatedIpAddressesWithContext(ctx context.Context, request *DescribeSubnetAllocatedIpAddressesRequest) string {
	if request == nil {
		request = NewDescribeSubnetAllocatedIpAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubnetAllocatedIpAddressesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSubnetAllocatedIpAddressesWithContextV2(ctx context.Context, request *DescribeSubnetAllocatedIpAddressesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSubnetAllocatedIpAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubnetAllocatedIpAddressesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddNatIpRequest() (request *AddNatIpRequest) {
	request = &AddNatIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AddNatIp")
	return
}

func NewAddNatIpResponse() (response *AddNatIpResponse) {
	response = &AddNatIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddNatIp(request *AddNatIpRequest) string {
	return c.AddNatIpWithContext(context.Background(), request)
}

func (c *Client) AddNatIpSend(request *AddNatIpRequest) (*AddNatIpResponse, error) {
	statusCode, msg, err := c.AddNatIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddNatIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddNatIpWithContext(ctx context.Context, request *AddNatIpRequest) string {
	if request == nil {
		request = NewAddNatIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddNatIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddNatIpWithContextV2(ctx context.Context, request *AddNatIpRequest) (int, string, error) {
	if request == nil {
		request = NewAddNatIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddNatIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNatIpRequest() (request *DeleteNatIpRequest) {
	request = &DeleteNatIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatIp")
	return
}

func NewDeleteNatIpResponse() (response *DeleteNatIpResponse) {
	response = &DeleteNatIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNatIp(request *DeleteNatIpRequest) string {
	return c.DeleteNatIpWithContext(context.Background(), request)
}

func (c *Client) DeleteNatIpSend(request *DeleteNatIpRequest) (*DeleteNatIpResponse, error) {
	statusCode, msg, err := c.DeleteNatIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNatIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNatIpWithContext(ctx context.Context, request *DeleteNatIpRequest) string {
	if request == nil {
		request = NewDeleteNatIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNatIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNatIpWithContextV2(ctx context.Context, request *DeleteNatIpRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNatIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNatIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateVpcCidrBlockRequest() (request *AssociateVpcCidrBlockRequest) {
	request = &AssociateVpcCidrBlockRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateVpcCidrBlock")
	return
}

func NewAssociateVpcCidrBlockResponse() (response *AssociateVpcCidrBlockResponse) {
	response = &AssociateVpcCidrBlockResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateVpcCidrBlock(request *AssociateVpcCidrBlockRequest) string {
	return c.AssociateVpcCidrBlockWithContext(context.Background(), request)
}

func (c *Client) AssociateVpcCidrBlockSend(request *AssociateVpcCidrBlockRequest) (*AssociateVpcCidrBlockResponse, error) {
	statusCode, msg, err := c.AssociateVpcCidrBlockWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateVpcCidrBlockResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateVpcCidrBlockWithContext(ctx context.Context, request *AssociateVpcCidrBlockRequest) string {
	if request == nil {
		request = NewAssociateVpcCidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateVpcCidrBlockResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateVpcCidrBlockWithContextV2(ctx context.Context, request *AssociateVpcCidrBlockRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateVpcCidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateVpcCidrBlockResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeIpv6PublicIpAddressesRequest() (request *DescribeIpv6PublicIpAddressesRequest) {
	request = &DescribeIpv6PublicIpAddressesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpv6PublicIpAddresses")
	return
}

func NewDescribeIpv6PublicIpAddressesResponse() (response *DescribeIpv6PublicIpAddressesResponse) {
	response = &DescribeIpv6PublicIpAddressesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeIpv6PublicIpAddresses(request *DescribeIpv6PublicIpAddressesRequest) string {
	return c.DescribeIpv6PublicIpAddressesWithContext(context.Background(), request)
}

func (c *Client) DescribeIpv6PublicIpAddressesSend(request *DescribeIpv6PublicIpAddressesRequest) (*DescribeIpv6PublicIpAddressesResponse, error) {
	statusCode, msg, err := c.DescribeIpv6PublicIpAddressesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeIpv6PublicIpAddressesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeIpv6PublicIpAddressesWithContext(ctx context.Context, request *DescribeIpv6PublicIpAddressesRequest) string {
	if request == nil {
		request = NewDescribeIpv6PublicIpAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeIpv6PublicIpAddressesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeIpv6PublicIpAddressesWithContextV2(ctx context.Context, request *DescribeIpv6PublicIpAddressesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeIpv6PublicIpAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeIpv6PublicIpAddressesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeIpv6NetworkInterfacesRequest() (request *DescribeIpv6NetworkInterfacesRequest) {
	request = &DescribeIpv6NetworkInterfacesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpv6NetworkInterfaces")
	return
}

func NewDescribeIpv6NetworkInterfacesResponse() (response *DescribeIpv6NetworkInterfacesResponse) {
	response = &DescribeIpv6NetworkInterfacesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeIpv6NetworkInterfaces(request *DescribeIpv6NetworkInterfacesRequest) string {
	return c.DescribeIpv6NetworkInterfacesWithContext(context.Background(), request)
}

func (c *Client) DescribeIpv6NetworkInterfacesSend(request *DescribeIpv6NetworkInterfacesRequest) (*DescribeIpv6NetworkInterfacesResponse, error) {
	statusCode, msg, err := c.DescribeIpv6NetworkInterfacesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeIpv6NetworkInterfacesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeIpv6NetworkInterfacesWithContext(ctx context.Context, request *DescribeIpv6NetworkInterfacesRequest) string {
	if request == nil {
		request = NewDescribeIpv6NetworkInterfacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeIpv6NetworkInterfacesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeIpv6NetworkInterfacesWithContextV2(ctx context.Context, request *DescribeIpv6NetworkInterfacesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeIpv6NetworkInterfacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeIpv6NetworkInterfacesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateIpv6PublicIpRequest() (request *CreateIpv6PublicIpRequest) {
	request = &CreateIpv6PublicIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateIpv6PublicIp")
	return
}

func NewCreateIpv6PublicIpResponse() (response *CreateIpv6PublicIpResponse) {
	response = &CreateIpv6PublicIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateIpv6PublicIp(request *CreateIpv6PublicIpRequest) string {
	return c.CreateIpv6PublicIpWithContext(context.Background(), request)
}

func (c *Client) CreateIpv6PublicIpSend(request *CreateIpv6PublicIpRequest) (*CreateIpv6PublicIpResponse, error) {
	statusCode, msg, err := c.CreateIpv6PublicIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateIpv6PublicIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateIpv6PublicIpWithContext(ctx context.Context, request *CreateIpv6PublicIpRequest) string {
	if request == nil {
		request = NewCreateIpv6PublicIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateIpv6PublicIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateIpv6PublicIpWithContextV2(ctx context.Context, request *CreateIpv6PublicIpRequest) (int, string, error) {
	if request == nil {
		request = NewCreateIpv6PublicIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateIpv6PublicIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReleaseIpv6PublicIpRequest() (request *ReleaseIpv6PublicIpRequest) {
	request = &ReleaseIpv6PublicIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReleaseIpv6PublicIp")
	return
}

func NewReleaseIpv6PublicIpResponse() (response *ReleaseIpv6PublicIpResponse) {
	response = &ReleaseIpv6PublicIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReleaseIpv6PublicIp(request *ReleaseIpv6PublicIpRequest) string {
	return c.ReleaseIpv6PublicIpWithContext(context.Background(), request)
}

func (c *Client) ReleaseIpv6PublicIpSend(request *ReleaseIpv6PublicIpRequest) (*ReleaseIpv6PublicIpResponse, error) {
	statusCode, msg, err := c.ReleaseIpv6PublicIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ReleaseIpv6PublicIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReleaseIpv6PublicIpWithContext(ctx context.Context, request *ReleaseIpv6PublicIpRequest) string {
	if request == nil {
		request = NewReleaseIpv6PublicIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReleaseIpv6PublicIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReleaseIpv6PublicIpWithContextV2(ctx context.Context, request *ReleaseIpv6PublicIpRequest) (int, string, error) {
	if request == nil {
		request = NewReleaseIpv6PublicIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReleaseIpv6PublicIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAlterIpv6PublicIpStateRequest() (request *AlterIpv6PublicIpStateRequest) {
	request = &AlterIpv6PublicIpStateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AlterIpv6PublicIpState")
	return
}

func NewAlterIpv6PublicIpStateResponse() (response *AlterIpv6PublicIpStateResponse) {
	response = &AlterIpv6PublicIpStateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AlterIpv6PublicIpState(request *AlterIpv6PublicIpStateRequest) string {
	return c.AlterIpv6PublicIpStateWithContext(context.Background(), request)
}

func (c *Client) AlterIpv6PublicIpStateSend(request *AlterIpv6PublicIpStateRequest) (*AlterIpv6PublicIpStateResponse, error) {
	statusCode, msg, err := c.AlterIpv6PublicIpStateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AlterIpv6PublicIpStateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AlterIpv6PublicIpStateWithContext(ctx context.Context, request *AlterIpv6PublicIpStateRequest) string {
	if request == nil {
		request = NewAlterIpv6PublicIpStateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAlterIpv6PublicIpStateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AlterIpv6PublicIpStateWithContextV2(ctx context.Context, request *AlterIpv6PublicIpStateRequest) (int, string, error) {
	if request == nil {
		request = NewAlterIpv6PublicIpStateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAlterIpv6PublicIpStateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyIpv6PublicIpRequest() (request *ModifyIpv6PublicIpRequest) {
	request = &ModifyIpv6PublicIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyIpv6PublicIp")
	return
}

func NewModifyIpv6PublicIpResponse() (response *ModifyIpv6PublicIpResponse) {
	response = &ModifyIpv6PublicIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyIpv6PublicIp(request *ModifyIpv6PublicIpRequest) string {
	return c.ModifyIpv6PublicIpWithContext(context.Background(), request)
}

func (c *Client) ModifyIpv6PublicIpSend(request *ModifyIpv6PublicIpRequest) (*ModifyIpv6PublicIpResponse, error) {
	statusCode, msg, err := c.ModifyIpv6PublicIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyIpv6PublicIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyIpv6PublicIpWithContext(ctx context.Context, request *ModifyIpv6PublicIpRequest) string {
	if request == nil {
		request = NewModifyIpv6PublicIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyIpv6PublicIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyIpv6PublicIpWithContextV2(ctx context.Context, request *ModifyIpv6PublicIpRequest) (int, string, error) {
	if request == nil {
		request = NewModifyIpv6PublicIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyIpv6PublicIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyPrivateIpAddressAttributeRequest() (request *ModifyPrivateIpAddressAttributeRequest) {
	request = &ModifyPrivateIpAddressAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyPrivateIpAddressAttribute")
	return
}

func NewModifyPrivateIpAddressAttributeResponse() (response *ModifyPrivateIpAddressAttributeResponse) {
	response = &ModifyPrivateIpAddressAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPrivateIpAddressAttribute(request *ModifyPrivateIpAddressAttributeRequest) string {
	return c.ModifyPrivateIpAddressAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyPrivateIpAddressAttributeSend(request *ModifyPrivateIpAddressAttributeRequest) (*ModifyPrivateIpAddressAttributeResponse, error) {
	statusCode, msg, err := c.ModifyPrivateIpAddressAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyPrivateIpAddressAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyPrivateIpAddressAttributeWithContext(ctx context.Context, request *ModifyPrivateIpAddressAttributeRequest) string {
	if request == nil {
		request = NewModifyPrivateIpAddressAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPrivateIpAddressAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyPrivateIpAddressAttributeWithContextV2(ctx context.Context, request *ModifyPrivateIpAddressAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyPrivateIpAddressAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPrivateIpAddressAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirectConnectRoutesRequest() (request *DescribeDirectConnectRoutesRequest) {
	request = &DescribeDirectConnectRoutesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectRoutes")
	return
}

func NewDescribeDirectConnectRoutesResponse() (response *DescribeDirectConnectRoutesResponse) {
	response = &DescribeDirectConnectRoutesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirectConnectRoutes(request *DescribeDirectConnectRoutesRequest) string {
	return c.DescribeDirectConnectRoutesWithContext(context.Background(), request)
}

func (c *Client) DescribeDirectConnectRoutesSend(request *DescribeDirectConnectRoutesRequest) (*DescribeDirectConnectRoutesResponse, error) {
	statusCode, msg, err := c.DescribeDirectConnectRoutesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirectConnectRoutesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirectConnectRoutesWithContext(ctx context.Context, request *DescribeDirectConnectRoutesRequest) string {
	if request == nil {
		request = NewDescribeDirectConnectRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectRoutesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirectConnectRoutesWithContextV2(ctx context.Context, request *DescribeDirectConnectRoutesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirectConnectRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirectConnectRoutesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewPublishDirectConnectRouteToBgpRequest() (request *PublishDirectConnectRouteToBgpRequest) {
	request = &PublishDirectConnectRouteToBgpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "PublishDirectConnectRouteToBgp")
	return
}

func NewPublishDirectConnectRouteToBgpResponse() (response *PublishDirectConnectRouteToBgpResponse) {
	response = &PublishDirectConnectRouteToBgpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) PublishDirectConnectRouteToBgp(request *PublishDirectConnectRouteToBgpRequest) string {
	return c.PublishDirectConnectRouteToBgpWithContext(context.Background(), request)
}

func (c *Client) PublishDirectConnectRouteToBgpSend(request *PublishDirectConnectRouteToBgpRequest) (*PublishDirectConnectRouteToBgpResponse, error) {
	statusCode, msg, err := c.PublishDirectConnectRouteToBgpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct PublishDirectConnectRouteToBgpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) PublishDirectConnectRouteToBgpWithContext(ctx context.Context, request *PublishDirectConnectRouteToBgpRequest) string {
	if request == nil {
		request = NewPublishDirectConnectRouteToBgpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPublishDirectConnectRouteToBgpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) PublishDirectConnectRouteToBgpWithContextV2(ctx context.Context, request *PublishDirectConnectRouteToBgpRequest) (int, string, error) {
	if request == nil {
		request = NewPublishDirectConnectRouteToBgpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPublishDirectConnectRouteToBgpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelDirectConnectRouteToBgpRequest() (request *CancelDirectConnectRouteToBgpRequest) {
	request = &CancelDirectConnectRouteToBgpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CancelDirectConnectRouteToBgp")
	return
}

func NewCancelDirectConnectRouteToBgpResponse() (response *CancelDirectConnectRouteToBgpResponse) {
	response = &CancelDirectConnectRouteToBgpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelDirectConnectRouteToBgp(request *CancelDirectConnectRouteToBgpRequest) string {
	return c.CancelDirectConnectRouteToBgpWithContext(context.Background(), request)
}

func (c *Client) CancelDirectConnectRouteToBgpSend(request *CancelDirectConnectRouteToBgpRequest) (*CancelDirectConnectRouteToBgpResponse, error) {
	statusCode, msg, err := c.CancelDirectConnectRouteToBgpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CancelDirectConnectRouteToBgpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelDirectConnectRouteToBgpWithContext(ctx context.Context, request *CancelDirectConnectRouteToBgpRequest) string {
	if request == nil {
		request = NewCancelDirectConnectRouteToBgpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelDirectConnectRouteToBgpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelDirectConnectRouteToBgpWithContextV2(ctx context.Context, request *CancelDirectConnectRouteToBgpRequest) (int, string, error) {
	if request == nil {
		request = NewCancelDirectConnectRouteToBgpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelDirectConnectRouteToBgpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachDirectConnectGatewayWithVpcRequest() (request *DetachDirectConnectGatewayWithVpcRequest) {
	request = &DetachDirectConnectGatewayWithVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DetachDirectConnectGatewayWithVpc")
	return
}

func NewDetachDirectConnectGatewayWithVpcResponse() (response *DetachDirectConnectGatewayWithVpcResponse) {
	response = &DetachDirectConnectGatewayWithVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachDirectConnectGatewayWithVpc(request *DetachDirectConnectGatewayWithVpcRequest) string {
	return c.DetachDirectConnectGatewayWithVpcWithContext(context.Background(), request)
}

func (c *Client) DetachDirectConnectGatewayWithVpcSend(request *DetachDirectConnectGatewayWithVpcRequest) (*DetachDirectConnectGatewayWithVpcResponse, error) {
	statusCode, msg, err := c.DetachDirectConnectGatewayWithVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DetachDirectConnectGatewayWithVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachDirectConnectGatewayWithVpcWithContext(ctx context.Context, request *DetachDirectConnectGatewayWithVpcRequest) string {
	if request == nil {
		request = NewDetachDirectConnectGatewayWithVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachDirectConnectGatewayWithVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachDirectConnectGatewayWithVpcWithContextV2(ctx context.Context, request *DetachDirectConnectGatewayWithVpcRequest) (int, string, error) {
	if request == nil {
		request = NewDetachDirectConnectGatewayWithVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachDirectConnectGatewayWithVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachDirectConnectGatewayWithVpcRequest() (request *AttachDirectConnectGatewayWithVpcRequest) {
	request = &AttachDirectConnectGatewayWithVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AttachDirectConnectGatewayWithVpc")
	return
}

func NewAttachDirectConnectGatewayWithVpcResponse() (response *AttachDirectConnectGatewayWithVpcResponse) {
	response = &AttachDirectConnectGatewayWithVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachDirectConnectGatewayWithVpc(request *AttachDirectConnectGatewayWithVpcRequest) string {
	return c.AttachDirectConnectGatewayWithVpcWithContext(context.Background(), request)
}

func (c *Client) AttachDirectConnectGatewayWithVpcSend(request *AttachDirectConnectGatewayWithVpcRequest) (*AttachDirectConnectGatewayWithVpcResponse, error) {
	statusCode, msg, err := c.AttachDirectConnectGatewayWithVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AttachDirectConnectGatewayWithVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachDirectConnectGatewayWithVpcWithContext(ctx context.Context, request *AttachDirectConnectGatewayWithVpcRequest) string {
	if request == nil {
		request = NewAttachDirectConnectGatewayWithVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachDirectConnectGatewayWithVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachDirectConnectGatewayWithVpcWithContextV2(ctx context.Context, request *AttachDirectConnectGatewayWithVpcRequest) (int, string, error) {
	if request == nil {
		request = NewAttachDirectConnectGatewayWithVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachDirectConnectGatewayWithVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateSubnetIpv6CidrBlockRequest() (request *AllocateSubnetIpv6CidrBlockRequest) {
	request = &AllocateSubnetIpv6CidrBlockRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AllocateSubnetIpv6CidrBlock")
	return
}

func NewAllocateSubnetIpv6CidrBlockResponse() (response *AllocateSubnetIpv6CidrBlockResponse) {
	response = &AllocateSubnetIpv6CidrBlockResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateSubnetIpv6CidrBlock(request *AllocateSubnetIpv6CidrBlockRequest) string {
	return c.AllocateSubnetIpv6CidrBlockWithContext(context.Background(), request)
}

func (c *Client) AllocateSubnetIpv6CidrBlockSend(request *AllocateSubnetIpv6CidrBlockRequest) (*AllocateSubnetIpv6CidrBlockResponse, error) {
	statusCode, msg, err := c.AllocateSubnetIpv6CidrBlockWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AllocateSubnetIpv6CidrBlockResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateSubnetIpv6CidrBlockWithContext(ctx context.Context, request *AllocateSubnetIpv6CidrBlockRequest) string {
	if request == nil {
		request = NewAllocateSubnetIpv6CidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateSubnetIpv6CidrBlockResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateSubnetIpv6CidrBlockWithContextV2(ctx context.Context, request *AllocateSubnetIpv6CidrBlockRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateSubnetIpv6CidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateSubnetIpv6CidrBlockResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRouteTableRequest() (request *CreateRouteTableRequest) {
	request = &CreateRouteTableRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateRouteTable")
	return
}

func NewCreateRouteTableResponse() (response *CreateRouteTableResponse) {
	response = &CreateRouteTableResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) string {
	return c.CreateRouteTableWithContext(context.Background(), request)
}

func (c *Client) CreateRouteTableSend(request *CreateRouteTableRequest) (*CreateRouteTableResponse, error) {
	statusCode, msg, err := c.CreateRouteTableWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateRouteTableResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRouteTableWithContext(ctx context.Context, request *CreateRouteTableRequest) string {
	if request == nil {
		request = NewCreateRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRouteTableResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRouteTableWithContextV2(ctx context.Context, request *CreateRouteTableRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRouteTableResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRouteTableRequest() (request *DeleteRouteTableRequest) {
	request = &DeleteRouteTableRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteRouteTable")
	return
}

func NewDeleteRouteTableResponse() (response *DeleteRouteTableResponse) {
	response = &DeleteRouteTableResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) string {
	return c.DeleteRouteTableWithContext(context.Background(), request)
}

func (c *Client) DeleteRouteTableSend(request *DeleteRouteTableRequest) (*DeleteRouteTableResponse, error) {
	statusCode, msg, err := c.DeleteRouteTableWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteRouteTableResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRouteTableWithContext(ctx context.Context, request *DeleteRouteTableRequest) string {
	if request == nil {
		request = NewDeleteRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRouteTableResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRouteTableWithContextV2(ctx context.Context, request *DeleteRouteTableRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRouteTableResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyRouteTableRequest() (request *ModifyRouteTableRequest) {
	request = &ModifyRouteTableRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyRouteTable")
	return
}

func NewModifyRouteTableResponse() (response *ModifyRouteTableResponse) {
	response = &ModifyRouteTableResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRouteTable(request *ModifyRouteTableRequest) string {
	return c.ModifyRouteTableWithContext(context.Background(), request)
}

func (c *Client) ModifyRouteTableSend(request *ModifyRouteTableRequest) (*ModifyRouteTableResponse, error) {
	statusCode, msg, err := c.ModifyRouteTableWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyRouteTableResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyRouteTableWithContext(ctx context.Context, request *ModifyRouteTableRequest) string {
	if request == nil {
		request = NewModifyRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRouteTableResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyRouteTableWithContextV2(ctx context.Context, request *ModifyRouteTableRequest) (int, string, error) {
	if request == nil {
		request = NewModifyRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRouteTableResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
	request = &DescribeRouteTablesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRouteTables")
	return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
	response = &DescribeRouteTablesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) string {
	return c.DescribeRouteTablesWithContext(context.Background(), request)
}

func (c *Client) DescribeRouteTablesSend(request *DescribeRouteTablesRequest) (*DescribeRouteTablesResponse, error) {
	statusCode, msg, err := c.DescribeRouteTablesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeRouteTablesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRouteTablesWithContext(ctx context.Context, request *DescribeRouteTablesRequest) string {
	if request == nil {
		request = NewDescribeRouteTablesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRouteTablesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRouteTablesWithContextV2(ctx context.Context, request *DescribeRouteTablesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRouteTablesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRouteTablesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateRouteTableRequest() (request *AssociateRouteTableRequest) {
	request = &AssociateRouteTableRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateRouteTable")
	return
}

func NewAssociateRouteTableResponse() (response *AssociateRouteTableResponse) {
	response = &AssociateRouteTableResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateRouteTable(request *AssociateRouteTableRequest) string {
	return c.AssociateRouteTableWithContext(context.Background(), request)
}

func (c *Client) AssociateRouteTableSend(request *AssociateRouteTableRequest) (*AssociateRouteTableResponse, error) {
	statusCode, msg, err := c.AssociateRouteTableWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateRouteTableResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateRouteTableWithContext(ctx context.Context, request *AssociateRouteTableRequest) string {
	if request == nil {
		request = NewAssociateRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateRouteTableResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateRouteTableWithContextV2(ctx context.Context, request *AssociateRouteTableRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateRouteTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateRouteTableResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
	request = &DeleteNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkInterface")
	return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNetworkInterfaceResponse) {
	response = &DeleteNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) string {
	return c.DeleteNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) DeleteNetworkInterfaceSend(request *DeleteNetworkInterfaceRequest) (*DeleteNetworkInterfaceResponse, error) {
	statusCode, msg, err := c.DeleteNetworkInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNetworkInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNetworkInterfaceWithContext(ctx context.Context, request *DeleteNetworkInterfaceRequest) string {
	if request == nil {
		request = NewDeleteNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNetworkInterfaceWithContextV2(ctx context.Context, request *DeleteNetworkInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNetworkInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
	request = &CreateNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkInterface")
	return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
	response = &CreateNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) string {
	return c.CreateNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) CreateNetworkInterfaceSend(request *CreateNetworkInterfaceRequest) (*CreateNetworkInterfaceResponse, error) {
	statusCode, msg, err := c.CreateNetworkInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateNetworkInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNetworkInterfaceWithContext(ctx context.Context, request *CreateNetworkInterfaceRequest) string {
	if request == nil {
		request = NewCreateNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNetworkInterfaceWithContextV2(ctx context.Context, request *CreateNetworkInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNetworkInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNetworkInterfaceRequest() (request *ModifyNetworkInterfaceRequest) {
	request = &ModifyNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkInterface")
	return
}

func NewModifyNetworkInterfaceResponse() (response *ModifyNetworkInterfaceResponse) {
	response = &ModifyNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNetworkInterface(request *ModifyNetworkInterfaceRequest) string {
	return c.ModifyNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) ModifyNetworkInterfaceSend(request *ModifyNetworkInterfaceRequest) (*ModifyNetworkInterfaceResponse, error) {
	statusCode, msg, err := c.ModifyNetworkInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyNetworkInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNetworkInterfaceWithContext(ctx context.Context, request *ModifyNetworkInterfaceRequest) string {
	if request == nil {
		request = NewModifyNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNetworkInterfaceWithContextV2(ctx context.Context, request *ModifyNetworkInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNatRateLimitRequest() (request *CreateNatRateLimitRequest) {
	request = &CreateNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNatRateLimit")
	return
}

func NewCreateNatRateLimitResponse() (response *CreateNatRateLimitResponse) {
	response = &CreateNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNatRateLimit(request *CreateNatRateLimitRequest) string {
	return c.CreateNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) CreateNatRateLimitSend(request *CreateNatRateLimitRequest) (*CreateNatRateLimitResponse, error) {
	statusCode, msg, err := c.CreateNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNatRateLimitWithContext(ctx context.Context, request *CreateNatRateLimitRequest) string {
	if request == nil {
		request = NewCreateNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNatRateLimitWithContextV2(ctx context.Context, request *CreateNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNatRateLimitRequest() (request *DescribeNatRateLimitRequest) {
	request = &DescribeNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatRateLimit")
	return
}

func NewDescribeNatRateLimitResponse() (response *DescribeNatRateLimitResponse) {
	response = &DescribeNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNatRateLimit(request *DescribeNatRateLimitRequest) string {
	return c.DescribeNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) DescribeNatRateLimitSend(request *DescribeNatRateLimitRequest) (*DescribeNatRateLimitResponse, error) {
	statusCode, msg, err := c.DescribeNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNatRateLimitWithContext(ctx context.Context, request *DescribeNatRateLimitRequest) string {
	if request == nil {
		request = NewDescribeNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNatRateLimitWithContextV2(ctx context.Context, request *DescribeNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNatRateLimitRequest() (request *ModifyNatRateLimitRequest) {
	request = &ModifyNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatRateLimit")
	return
}

func NewModifyNatRateLimitResponse() (response *ModifyNatRateLimitResponse) {
	response = &ModifyNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNatRateLimit(request *ModifyNatRateLimitRequest) string {
	return c.ModifyNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) ModifyNatRateLimitSend(request *ModifyNatRateLimitRequest) (*ModifyNatRateLimitResponse, error) {
	statusCode, msg, err := c.ModifyNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNatRateLimitWithContext(ctx context.Context, request *ModifyNatRateLimitRequest) string {
	if request == nil {
		request = NewModifyNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNatRateLimitWithContextV2(ctx context.Context, request *ModifyNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNatRateLimitRequest() (request *DeleteNatRateLimitRequest) {
	request = &DeleteNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatRateLimit")
	return
}

func NewDeleteNatRateLimitResponse() (response *DeleteNatRateLimitResponse) {
	response = &DeleteNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNatRateLimit(request *DeleteNatRateLimitRequest) string {
	return c.DeleteNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) DeleteNatRateLimitSend(request *DeleteNatRateLimitRequest) (*DeleteNatRateLimitResponse, error) {
	statusCode, msg, err := c.DeleteNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNatRateLimitWithContext(ctx context.Context, request *DeleteNatRateLimitRequest) string {
	if request == nil {
		request = NewDeleteNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNatRateLimitWithContextV2(ctx context.Context, request *DeleteNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDnatRequest() (request *CreateDnatRequest) {
	request = &CreateDnatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDnat")
	return
}

func NewCreateDnatResponse() (response *CreateDnatResponse) {
	response = &CreateDnatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDnat(request *CreateDnatRequest) string {
	return c.CreateDnatWithContext(context.Background(), request)
}

func (c *Client) CreateDnatSend(request *CreateDnatRequest) (*CreateDnatResponse, error) {
	statusCode, msg, err := c.CreateDnatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDnatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDnatWithContext(ctx context.Context, request *CreateDnatRequest) string {
	if request == nil {
		request = NewCreateDnatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDnatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDnatWithContextV2(ctx context.Context, request *CreateDnatRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDnatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDnatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDnatRequest() (request *DeleteDnatRequest) {
	request = &DeleteDnatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDnat")
	return
}

func NewDeleteDnatResponse() (response *DeleteDnatResponse) {
	response = &DeleteDnatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDnat(request *DeleteDnatRequest) string {
	return c.DeleteDnatWithContext(context.Background(), request)
}

func (c *Client) DeleteDnatSend(request *DeleteDnatRequest) (*DeleteDnatResponse, error) {
	statusCode, msg, err := c.DeleteDnatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDnatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDnatWithContext(ctx context.Context, request *DeleteDnatRequest) string {
	if request == nil {
		request = NewDeleteDnatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDnatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDnatWithContextV2(ctx context.Context, request *DeleteDnatRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDnatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDnatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDnatsRequest() (request *DescribeDnatsRequest) {
	request = &DescribeDnatsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDnats")
	return
}

func NewDescribeDnatsResponse() (response *DescribeDnatsResponse) {
	response = &DescribeDnatsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDnats(request *DescribeDnatsRequest) string {
	return c.DescribeDnatsWithContext(context.Background(), request)
}

func (c *Client) DescribeDnatsSend(request *DescribeDnatsRequest) (*DescribeDnatsResponse, error) {
	statusCode, msg, err := c.DescribeDnatsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDnatsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDnatsWithContext(ctx context.Context, request *DescribeDnatsRequest) string {
	if request == nil {
		request = NewDescribeDnatsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDnatsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDnatsWithContextV2(ctx context.Context, request *DescribeDnatsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDnatsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDnatsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDnatRequest() (request *ModifyDnatRequest) {
	request = &ModifyDnatRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyDnat")
	return
}

func NewModifyDnatResponse() (response *ModifyDnatResponse) {
	response = &ModifyDnatResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDnat(request *ModifyDnatRequest) string {
	return c.ModifyDnatWithContext(context.Background(), request)
}

func (c *Client) ModifyDnatSend(request *ModifyDnatRequest) (*ModifyDnatResponse, error) {
	statusCode, msg, err := c.ModifyDnatWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDnatResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDnatWithContext(ctx context.Context, request *ModifyDnatRequest) string {
	if request == nil {
		request = NewModifyDnatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDnatResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDnatWithContextV2(ctx context.Context, request *ModifyDnatRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDnatRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDnatResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateInstanceRequest() (request *AssociateInstanceRequest) {
	request = &AssociateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateInstance")
	return
}

func NewAssociateInstanceResponse() (response *AssociateInstanceResponse) {
	response = &AssociateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateInstance(request *AssociateInstanceRequest) string {
	return c.AssociateInstanceWithContext(context.Background(), request)
}

func (c *Client) AssociateInstanceSend(request *AssociateInstanceRequest) (*AssociateInstanceResponse, error) {
	statusCode, msg, err := c.AssociateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateInstanceWithContext(ctx context.Context, request *AssociateInstanceRequest) string {
	if request == nil {
		request = NewAssociateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateInstanceWithContextV2(ctx context.Context, request *AssociateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateInstanceRequest() (request *DisassociateInstanceRequest) {
	request = &DisassociateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateInstance")
	return
}

func NewDisassociateInstanceResponse() (response *DisassociateInstanceResponse) {
	response = &DisassociateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateInstance(request *DisassociateInstanceRequest) string {
	return c.DisassociateInstanceWithContext(context.Background(), request)
}

func (c *Client) DisassociateInstanceSend(request *DisassociateInstanceRequest) (*DisassociateInstanceResponse, error) {
	statusCode, msg, err := c.DisassociateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisassociateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateInstanceWithContext(ctx context.Context, request *DisassociateInstanceRequest) string {
	if request == nil {
		request = NewDisassociateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateInstanceWithContextV2(ctx context.Context, request *DisassociateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateHaVipRequest() (request *CreateHaVipRequest) {
	request = &CreateHaVipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateHaVip")
	return
}

func NewCreateHaVipResponse() (response *CreateHaVipResponse) {
	response = &CreateHaVipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateHaVip(request *CreateHaVipRequest) string {
	return c.CreateHaVipWithContext(context.Background(), request)
}

func (c *Client) CreateHaVipSend(request *CreateHaVipRequest) (*CreateHaVipResponse, error) {
	statusCode, msg, err := c.CreateHaVipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateHaVipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateHaVipWithContext(ctx context.Context, request *CreateHaVipRequest) string {
	if request == nil {
		request = NewCreateHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateHaVipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateHaVipWithContextV2(ctx context.Context, request *CreateHaVipRequest) (int, string, error) {
	if request == nil {
		request = NewCreateHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateHaVipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteHaVipRequest() (request *DeleteHaVipRequest) {
	request = &DeleteHaVipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteHaVip")
	return
}

func NewDeleteHaVipResponse() (response *DeleteHaVipResponse) {
	response = &DeleteHaVipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) string {
	return c.DeleteHaVipWithContext(context.Background(), request)
}

func (c *Client) DeleteHaVipSend(request *DeleteHaVipRequest) (*DeleteHaVipResponse, error) {
	statusCode, msg, err := c.DeleteHaVipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteHaVipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteHaVipWithContext(ctx context.Context, request *DeleteHaVipRequest) string {
	if request == nil {
		request = NewDeleteHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHaVipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteHaVipWithContextV2(ctx context.Context, request *DeleteHaVipRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHaVipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateHaVipRequest() (request *AssociateHaVipRequest) {
	request = &AssociateHaVipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateHaVip")
	return
}

func NewAssociateHaVipResponse() (response *AssociateHaVipResponse) {
	response = &AssociateHaVipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateHaVip(request *AssociateHaVipRequest) string {
	return c.AssociateHaVipWithContext(context.Background(), request)
}

func (c *Client) AssociateHaVipSend(request *AssociateHaVipRequest) (*AssociateHaVipResponse, error) {
	statusCode, msg, err := c.AssociateHaVipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateHaVipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateHaVipWithContext(ctx context.Context, request *AssociateHaVipRequest) string {
	if request == nil {
		request = NewAssociateHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateHaVipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateHaVipWithContextV2(ctx context.Context, request *AssociateHaVipRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateHaVipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnAssociateHaVipRequest() (request *UnAssociateHaVipRequest) {
	request = &UnAssociateHaVipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnAssociateHaVip")
	return
}

func NewUnAssociateHaVipResponse() (response *UnAssociateHaVipResponse) {
	response = &UnAssociateHaVipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnAssociateHaVip(request *UnAssociateHaVipRequest) string {
	return c.UnAssociateHaVipWithContext(context.Background(), request)
}

func (c *Client) UnAssociateHaVipSend(request *UnAssociateHaVipRequest) (*UnAssociateHaVipResponse, error) {
	statusCode, msg, err := c.UnAssociateHaVipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UnAssociateHaVipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnAssociateHaVipWithContext(ctx context.Context, request *UnAssociateHaVipRequest) string {
	if request == nil {
		request = NewUnAssociateHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnAssociateHaVipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnAssociateHaVipWithContextV2(ctx context.Context, request *UnAssociateHaVipRequest) (int, string, error) {
	if request == nil {
		request = NewUnAssociateHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnAssociateHaVipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeHaVipRequest() (request *DescribeHaVipRequest) {
	request = &DescribeHaVipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeHaVip")
	return
}

func NewDescribeHaVipResponse() (response *DescribeHaVipResponse) {
	response = &DescribeHaVipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeHaVip(request *DescribeHaVipRequest) string {
	return c.DescribeHaVipWithContext(context.Background(), request)
}

func (c *Client) DescribeHaVipSend(request *DescribeHaVipRequest) (*DescribeHaVipResponse, error) {
	statusCode, msg, err := c.DescribeHaVipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeHaVipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeHaVipWithContext(ctx context.Context, request *DescribeHaVipRequest) string {
	if request == nil {
		request = NewDescribeHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHaVipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeHaVipWithContextV2(ctx context.Context, request *DescribeHaVipRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeHaVipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHaVipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDirectConnectGatewayRouteRequest() (request *CreateDirectConnectGatewayRouteRequest) {
	request = &CreateDirectConnectGatewayRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectGatewayRoute")
	return
}

func NewCreateDirectConnectGatewayRouteResponse() (response *CreateDirectConnectGatewayRouteResponse) {
	response = &CreateDirectConnectGatewayRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDirectConnectGatewayRoute(request *CreateDirectConnectGatewayRouteRequest) string {
	return c.CreateDirectConnectGatewayRouteWithContext(context.Background(), request)
}

func (c *Client) CreateDirectConnectGatewayRouteSend(request *CreateDirectConnectGatewayRouteRequest) (*CreateDirectConnectGatewayRouteResponse, error) {
	statusCode, msg, err := c.CreateDirectConnectGatewayRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDirectConnectGatewayRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDirectConnectGatewayRouteWithContext(ctx context.Context, request *CreateDirectConnectGatewayRouteRequest) string {
	if request == nil {
		request = NewCreateDirectConnectGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDirectConnectGatewayRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDirectConnectGatewayRouteWithContextV2(ctx context.Context, request *CreateDirectConnectGatewayRouteRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDirectConnectGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDirectConnectGatewayRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDirectConnectGatewayRouteRequest() (request *DeleteDirectConnectGatewayRouteRequest) {
	request = &DeleteDirectConnectGatewayRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectGatewayRoute")
	return
}

func NewDeleteDirectConnectGatewayRouteResponse() (response *DeleteDirectConnectGatewayRouteResponse) {
	response = &DeleteDirectConnectGatewayRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDirectConnectGatewayRoute(request *DeleteDirectConnectGatewayRouteRequest) string {
	return c.DeleteDirectConnectGatewayRouteWithContext(context.Background(), request)
}

func (c *Client) DeleteDirectConnectGatewayRouteSend(request *DeleteDirectConnectGatewayRouteRequest) (*DeleteDirectConnectGatewayRouteResponse, error) {
	statusCode, msg, err := c.DeleteDirectConnectGatewayRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDirectConnectGatewayRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDirectConnectGatewayRouteWithContext(ctx context.Context, request *DeleteDirectConnectGatewayRouteRequest) string {
	if request == nil {
		request = NewDeleteDirectConnectGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDirectConnectGatewayRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDirectConnectGatewayRouteWithContextV2(ctx context.Context, request *DeleteDirectConnectGatewayRouteRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDirectConnectGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDirectConnectGatewayRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirectConnectGatewayRouteRequest() (request *DescribeDirectConnectGatewayRouteRequest) {
	request = &DescribeDirectConnectGatewayRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectGatewayRoute")
	return
}

func NewDescribeDirectConnectGatewayRouteResponse() (response *DescribeDirectConnectGatewayRouteResponse) {
	response = &DescribeDirectConnectGatewayRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirectConnectGatewayRoute(request *DescribeDirectConnectGatewayRouteRequest) string {
	return c.DescribeDirectConnectGatewayRouteWithContext(context.Background(), request)
}

func (c *Client) DescribeDirectConnectGatewayRouteSend(request *DescribeDirectConnectGatewayRouteRequest) (*DescribeDirectConnectGatewayRouteResponse, error) {
	statusCode, msg, err := c.DescribeDirectConnectGatewayRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirectConnectGatewayRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirectConnectGatewayRouteWithContext(ctx context.Context, request *DescribeDirectConnectGatewayRouteRequest) string {
	if request == nil {
		request = NewDescribeDirectConnectGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDirectConnectGatewayRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirectConnectGatewayRouteWithContextV2(ctx context.Context, request *DescribeDirectConnectGatewayRouteRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirectConnectGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDirectConnectGatewayRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewPublishDirectConnectRouteRequest() (request *PublishDirectConnectRouteRequest) {
	request = &PublishDirectConnectRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "PublishDirectConnectRoute")
	return
}

func NewPublishDirectConnectRouteResponse() (response *PublishDirectConnectRouteResponse) {
	response = &PublishDirectConnectRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) PublishDirectConnectRoute(request *PublishDirectConnectRouteRequest) string {
	return c.PublishDirectConnectRouteWithContext(context.Background(), request)
}

func (c *Client) PublishDirectConnectRouteSend(request *PublishDirectConnectRouteRequest) (*PublishDirectConnectRouteResponse, error) {
	statusCode, msg, err := c.PublishDirectConnectRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct PublishDirectConnectRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) PublishDirectConnectRouteWithContext(ctx context.Context, request *PublishDirectConnectRouteRequest) string {
	if request == nil {
		request = NewPublishDirectConnectRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPublishDirectConnectRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) PublishDirectConnectRouteWithContextV2(ctx context.Context, request *PublishDirectConnectRouteRequest) (int, string, error) {
	if request == nil {
		request = NewPublishDirectConnectRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPublishDirectConnectRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnpublishDirectConnectRouteRequest() (request *UnpublishDirectConnectRouteRequest) {
	request = &UnpublishDirectConnectRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnpublishDirectConnectRoute")
	return
}

func NewUnpublishDirectConnectRouteResponse() (response *UnpublishDirectConnectRouteResponse) {
	response = &UnpublishDirectConnectRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnpublishDirectConnectRoute(request *UnpublishDirectConnectRouteRequest) string {
	return c.UnpublishDirectConnectRouteWithContext(context.Background(), request)
}

func (c *Client) UnpublishDirectConnectRouteSend(request *UnpublishDirectConnectRouteRequest) (*UnpublishDirectConnectRouteResponse, error) {
	statusCode, msg, err := c.UnpublishDirectConnectRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UnpublishDirectConnectRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnpublishDirectConnectRouteWithContext(ctx context.Context, request *UnpublishDirectConnectRouteRequest) string {
	if request == nil {
		request = NewUnpublishDirectConnectRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnpublishDirectConnectRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnpublishDirectConnectRouteWithContextV2(ctx context.Context, request *UnpublishDirectConnectRouteRequest) (int, string, error) {
	if request == nil {
		request = NewUnpublishDirectConnectRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnpublishDirectConnectRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddSecondaryCidrBlockRequest() (request *AddSecondaryCidrBlockRequest) {
	request = &AddSecondaryCidrBlockRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AddSecondaryCidrBlock")
	return
}

func NewAddSecondaryCidrBlockResponse() (response *AddSecondaryCidrBlockResponse) {
	response = &AddSecondaryCidrBlockResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddSecondaryCidrBlock(request *AddSecondaryCidrBlockRequest) string {
	return c.AddSecondaryCidrBlockWithContext(context.Background(), request)
}

func (c *Client) AddSecondaryCidrBlockSend(request *AddSecondaryCidrBlockRequest) (*AddSecondaryCidrBlockResponse, error) {
	statusCode, msg, err := c.AddSecondaryCidrBlockWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddSecondaryCidrBlockResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddSecondaryCidrBlockWithContext(ctx context.Context, request *AddSecondaryCidrBlockRequest) string {
	if request == nil {
		request = NewAddSecondaryCidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddSecondaryCidrBlockResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddSecondaryCidrBlockWithContextV2(ctx context.Context, request *AddSecondaryCidrBlockRequest) (int, string, error) {
	if request == nil {
		request = NewAddSecondaryCidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddSecondaryCidrBlockResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecondaryCidrBlockRequest() (request *DeleteSecondaryCidrBlockRequest) {
	request = &DeleteSecondaryCidrBlockRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecondaryCidrBlock")
	return
}

func NewDeleteSecondaryCidrBlockResponse() (response *DeleteSecondaryCidrBlockResponse) {
	response = &DeleteSecondaryCidrBlockResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecondaryCidrBlock(request *DeleteSecondaryCidrBlockRequest) string {
	return c.DeleteSecondaryCidrBlockWithContext(context.Background(), request)
}

func (c *Client) DeleteSecondaryCidrBlockSend(request *DeleteSecondaryCidrBlockRequest) (*DeleteSecondaryCidrBlockResponse, error) {
	statusCode, msg, err := c.DeleteSecondaryCidrBlockWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSecondaryCidrBlockResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecondaryCidrBlockWithContext(ctx context.Context, request *DeleteSecondaryCidrBlockRequest) string {
	if request == nil {
		request = NewDeleteSecondaryCidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecondaryCidrBlockResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecondaryCidrBlockWithContextV2(ctx context.Context, request *DeleteSecondaryCidrBlockRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecondaryCidrBlockRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecondaryCidrBlockResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssignPrivateIpAddressRequest() (request *AssignPrivateIpAddressRequest) {
	request = &AssignPrivateIpAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssignPrivateIpAddress")
	return
}

func NewAssignPrivateIpAddressResponse() (response *AssignPrivateIpAddressResponse) {
	response = &AssignPrivateIpAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssignPrivateIpAddress(request *AssignPrivateIpAddressRequest) string {
	return c.AssignPrivateIpAddressWithContext(context.Background(), request)
}

func (c *Client) AssignPrivateIpAddressSend(request *AssignPrivateIpAddressRequest) (*AssignPrivateIpAddressResponse, error) {
	statusCode, msg, err := c.AssignPrivateIpAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssignPrivateIpAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssignPrivateIpAddressWithContext(ctx context.Context, request *AssignPrivateIpAddressRequest) string {
	if request == nil {
		request = NewAssignPrivateIpAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssignPrivateIpAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssignPrivateIpAddressWithContextV2(ctx context.Context, request *AssignPrivateIpAddressRequest) (int, string, error) {
	if request == nil {
		request = NewAssignPrivateIpAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssignPrivateIpAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnassignPrivateIpAddressRequest() (request *UnassignPrivateIpAddressRequest) {
	request = &UnassignPrivateIpAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnassignPrivateIpAddress")
	return
}

func NewUnassignPrivateIpAddressResponse() (response *UnassignPrivateIpAddressResponse) {
	response = &UnassignPrivateIpAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnassignPrivateIpAddress(request *UnassignPrivateIpAddressRequest) string {
	return c.UnassignPrivateIpAddressWithContext(context.Background(), request)
}

func (c *Client) UnassignPrivateIpAddressSend(request *UnassignPrivateIpAddressRequest) (*UnassignPrivateIpAddressResponse, error) {
	statusCode, msg, err := c.UnassignPrivateIpAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UnassignPrivateIpAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnassignPrivateIpAddressWithContext(ctx context.Context, request *UnassignPrivateIpAddressRequest) string {
	if request == nil {
		request = NewUnassignPrivateIpAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnassignPrivateIpAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnassignPrivateIpAddressWithContextV2(ctx context.Context, request *UnassignPrivateIpAddressRequest) (int, string, error) {
	if request == nil {
		request = NewUnassignPrivateIpAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnassignPrivateIpAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchCreateNatRateLimitRequest() (request *BatchCreateNatRateLimitRequest) {
	request = &BatchCreateNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "BatchCreateNatRateLimit")
	return
}

func NewBatchCreateNatRateLimitResponse() (response *BatchCreateNatRateLimitResponse) {
	response = &BatchCreateNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchCreateNatRateLimit(request *BatchCreateNatRateLimitRequest) string {
	return c.BatchCreateNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) BatchCreateNatRateLimitSend(request *BatchCreateNatRateLimitRequest) (*BatchCreateNatRateLimitResponse, error) {
	statusCode, msg, err := c.BatchCreateNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BatchCreateNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchCreateNatRateLimitWithContext(ctx context.Context, request *BatchCreateNatRateLimitRequest) string {
	if request == nil {
		request = NewBatchCreateNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchCreateNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchCreateNatRateLimitWithContextV2(ctx context.Context, request *BatchCreateNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewBatchCreateNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchCreateNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchModifyNatRateLimitRequest() (request *BatchModifyNatRateLimitRequest) {
	request = &BatchModifyNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "BatchModifyNatRateLimit")
	return
}

func NewBatchModifyNatRateLimitResponse() (response *BatchModifyNatRateLimitResponse) {
	response = &BatchModifyNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchModifyNatRateLimit(request *BatchModifyNatRateLimitRequest) string {
	return c.BatchModifyNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) BatchModifyNatRateLimitSend(request *BatchModifyNatRateLimitRequest) (*BatchModifyNatRateLimitResponse, error) {
	statusCode, msg, err := c.BatchModifyNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BatchModifyNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchModifyNatRateLimitWithContext(ctx context.Context, request *BatchModifyNatRateLimitRequest) string {
	if request == nil {
		request = NewBatchModifyNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchModifyNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchModifyNatRateLimitWithContextV2(ctx context.Context, request *BatchModifyNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewBatchModifyNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchModifyNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchDeleteNatRateLimitRequest() (request *BatchDeleteNatRateLimitRequest) {
	request = &BatchDeleteNatRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "BatchDeleteNatRateLimit")
	return
}

func NewBatchDeleteNatRateLimitResponse() (response *BatchDeleteNatRateLimitResponse) {
	response = &BatchDeleteNatRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchDeleteNatRateLimit(request *BatchDeleteNatRateLimitRequest) string {
	return c.BatchDeleteNatRateLimitWithContext(context.Background(), request)
}

func (c *Client) BatchDeleteNatRateLimitSend(request *BatchDeleteNatRateLimitRequest) (*BatchDeleteNatRateLimitResponse, error) {
	statusCode, msg, err := c.BatchDeleteNatRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BatchDeleteNatRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchDeleteNatRateLimitWithContext(ctx context.Context, request *BatchDeleteNatRateLimitRequest) string {
	if request == nil {
		request = NewBatchDeleteNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchDeleteNatRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchDeleteNatRateLimitWithContextV2(ctx context.Context, request *BatchDeleteNatRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewBatchDeleteNatRateLimitRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchDeleteNatRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpnGatewayRoutesRequest() (request *DescribeVpnGatewayRoutesRequest) {
	request = &DescribeVpnGatewayRoutesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewayRoutes")
	return
}

func NewDescribeVpnGatewayRoutesResponse() (response *DescribeVpnGatewayRoutesResponse) {
	response = &DescribeVpnGatewayRoutesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpnGatewayRoutes(request *DescribeVpnGatewayRoutesRequest) string {
	return c.DescribeVpnGatewayRoutesWithContext(context.Background(), request)
}

func (c *Client) DescribeVpnGatewayRoutesSend(request *DescribeVpnGatewayRoutesRequest) (*DescribeVpnGatewayRoutesResponse, error) {
	statusCode, msg, err := c.DescribeVpnGatewayRoutesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeVpnGatewayRoutesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpnGatewayRoutesWithContext(ctx context.Context, request *DescribeVpnGatewayRoutesRequest) string {
	if request == nil {
		request = NewDescribeVpnGatewayRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnGatewayRoutesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpnGatewayRoutesWithContextV2(ctx context.Context, request *DescribeVpnGatewayRoutesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpnGatewayRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnGatewayRoutesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateVpnGatewayRouteRequest() (request *CreateVpnGatewayRouteRequest) {
	request = &CreateVpnGatewayRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnGatewayRoute")
	return
}

func NewCreateVpnGatewayRouteResponse() (response *CreateVpnGatewayRouteResponse) {
	response = &CreateVpnGatewayRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpnGatewayRoute(request *CreateVpnGatewayRouteRequest) string {
	return c.CreateVpnGatewayRouteWithContext(context.Background(), request)
}

func (c *Client) CreateVpnGatewayRouteSend(request *CreateVpnGatewayRouteRequest) (*CreateVpnGatewayRouteResponse, error) {
	statusCode, msg, err := c.CreateVpnGatewayRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateVpnGatewayRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateVpnGatewayRouteWithContext(ctx context.Context, request *CreateVpnGatewayRouteRequest) string {
	if request == nil {
		request = NewCreateVpnGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpnGatewayRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateVpnGatewayRouteWithContextV2(ctx context.Context, request *CreateVpnGatewayRouteRequest) (int, string, error) {
	if request == nil {
		request = NewCreateVpnGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVpnGatewayRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteVpnGatewayRouteRequest() (request *DeleteVpnGatewayRouteRequest) {
	request = &DeleteVpnGatewayRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGatewayRoute")
	return
}

func NewDeleteVpnGatewayRouteResponse() (response *DeleteVpnGatewayRouteResponse) {
	response = &DeleteVpnGatewayRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpnGatewayRoute(request *DeleteVpnGatewayRouteRequest) string {
	return c.DeleteVpnGatewayRouteWithContext(context.Background(), request)
}

func (c *Client) DeleteVpnGatewayRouteSend(request *DeleteVpnGatewayRouteRequest) (*DeleteVpnGatewayRouteResponse, error) {
	statusCode, msg, err := c.DeleteVpnGatewayRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteVpnGatewayRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteVpnGatewayRouteWithContext(ctx context.Context, request *DeleteVpnGatewayRouteRequest) string {
	if request == nil {
		request = NewDeleteVpnGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpnGatewayRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteVpnGatewayRouteWithContextV2(ctx context.Context, request *DeleteVpnGatewayRouteRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteVpnGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVpnGatewayRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpnTunnelIpsecStatusRequest() (request *DescribeVpnTunnelIpsecStatusRequest) {
	request = &DescribeVpnTunnelIpsecStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnTunnelIpsecStatus")
	return
}

func NewDescribeVpnTunnelIpsecStatusResponse() (response *DescribeVpnTunnelIpsecStatusResponse) {
	response = &DescribeVpnTunnelIpsecStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpnTunnelIpsecStatus(request *DescribeVpnTunnelIpsecStatusRequest) string {
	return c.DescribeVpnTunnelIpsecStatusWithContext(context.Background(), request)
}

func (c *Client) DescribeVpnTunnelIpsecStatusSend(request *DescribeVpnTunnelIpsecStatusRequest) (*DescribeVpnTunnelIpsecStatusResponse, error) {
	statusCode, msg, err := c.DescribeVpnTunnelIpsecStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeVpnTunnelIpsecStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpnTunnelIpsecStatusWithContext(ctx context.Context, request *DescribeVpnTunnelIpsecStatusRequest) string {
	if request == nil {
		request = NewDescribeVpnTunnelIpsecStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnTunnelIpsecStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpnTunnelIpsecStatusWithContextV2(ctx context.Context, request *DescribeVpnTunnelIpsecStatusRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpnTunnelIpsecStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnTunnelIpsecStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryNatTopVifMonitorRequest() (request *QueryNatTopVifMonitorRequest) {
	request = &QueryNatTopVifMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "QueryNatTopVifMonitor")
	return
}

func NewQueryNatTopVifMonitorResponse() (response *QueryNatTopVifMonitorResponse) {
	response = &QueryNatTopVifMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryNatTopVifMonitor(request *QueryNatTopVifMonitorRequest) string {
	return c.QueryNatTopVifMonitorWithContext(context.Background(), request)
}

func (c *Client) QueryNatTopVifMonitorSend(request *QueryNatTopVifMonitorRequest) (*QueryNatTopVifMonitorResponse, error) {
	statusCode, msg, err := c.QueryNatTopVifMonitorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryNatTopVifMonitorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryNatTopVifMonitorWithContext(ctx context.Context, request *QueryNatTopVifMonitorRequest) string {
	if request == nil {
		request = NewQueryNatTopVifMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryNatTopVifMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryNatTopVifMonitorWithContextV2(ctx context.Context, request *QueryNatTopVifMonitorRequest) (int, string, error) {
	if request == nil {
		request = NewQueryNatTopVifMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryNatTopVifMonitorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNatIpStatusRequest() (request *ModifyNatIpStatusRequest) {
	request = &ModifyNatIpStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatIpStatus")
	return
}

func NewModifyNatIpStatusResponse() (response *ModifyNatIpStatusResponse) {
	response = &ModifyNatIpStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNatIpStatus(request *ModifyNatIpStatusRequest) string {
	return c.ModifyNatIpStatusWithContext(context.Background(), request)
}

func (c *Client) ModifyNatIpStatusSend(request *ModifyNatIpStatusRequest) (*ModifyNatIpStatusResponse, error) {
	statusCode, msg, err := c.ModifyNatIpStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyNatIpStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNatIpStatusWithContext(ctx context.Context, request *ModifyNatIpStatusRequest) string {
	if request == nil {
		request = NewModifyNatIpStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNatIpStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNatIpStatusWithContextV2(ctx context.Context, request *ModifyNatIpStatusRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNatIpStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNatIpStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryPeerTopVifMonitorRequest() (request *QueryPeerTopVifMonitorRequest) {
	request = &QueryPeerTopVifMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "QueryPeerTopVifMonitor")
	return
}

func NewQueryPeerTopVifMonitorResponse() (response *QueryPeerTopVifMonitorResponse) {
	response = &QueryPeerTopVifMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryPeerTopVifMonitor(request *QueryPeerTopVifMonitorRequest) string {
	return c.QueryPeerTopVifMonitorWithContext(context.Background(), request)
}

func (c *Client) QueryPeerTopVifMonitorSend(request *QueryPeerTopVifMonitorRequest) (*QueryPeerTopVifMonitorResponse, error) {
	statusCode, msg, err := c.QueryPeerTopVifMonitorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryPeerTopVifMonitorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryPeerTopVifMonitorWithContext(ctx context.Context, request *QueryPeerTopVifMonitorRequest) string {
	if request == nil {
		request = NewQueryPeerTopVifMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryPeerTopVifMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryPeerTopVifMonitorWithContextV2(ctx context.Context, request *QueryPeerTopVifMonitorRequest) (int, string, error) {
	if request == nil {
		request = NewQueryPeerTopVifMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryPeerTopVifMonitorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVpnGatewayRouteRequest() (request *ModifyVpnGatewayRouteRequest) {
	request = &ModifyVpnGatewayRouteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewayRoute")
	return
}

func NewModifyVpnGatewayRouteResponse() (response *ModifyVpnGatewayRouteResponse) {
	response = &ModifyVpnGatewayRouteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpnGatewayRoute(request *ModifyVpnGatewayRouteRequest) string {
	return c.ModifyVpnGatewayRouteWithContext(context.Background(), request)
}

func (c *Client) ModifyVpnGatewayRouteSend(request *ModifyVpnGatewayRouteRequest) (*ModifyVpnGatewayRouteResponse, error) {
	statusCode, msg, err := c.ModifyVpnGatewayRouteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyVpnGatewayRouteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVpnGatewayRouteWithContext(ctx context.Context, request *ModifyVpnGatewayRouteRequest) string {
	if request == nil {
		request = NewModifyVpnGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpnGatewayRouteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVpnGatewayRouteWithContextV2(ctx context.Context, request *ModifyVpnGatewayRouteRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVpnGatewayRouteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVpnGatewayRouteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirectConnectInterfacesBgpStatusRequest() (request *DescribeDirectConnectInterfacesBgpStatusRequest) {
	request = &DescribeDirectConnectInterfacesBgpStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectInterfacesBgpStatus")
	return
}

func NewDescribeDirectConnectInterfacesBgpStatusResponse() (response *DescribeDirectConnectInterfacesBgpStatusResponse) {
	response = &DescribeDirectConnectInterfacesBgpStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirectConnectInterfacesBgpStatus(request *DescribeDirectConnectInterfacesBgpStatusRequest) string {
	return c.DescribeDirectConnectInterfacesBgpStatusWithContext(context.Background(), request)
}

func (c *Client) DescribeDirectConnectInterfacesBgpStatusSend(request *DescribeDirectConnectInterfacesBgpStatusRequest) (*DescribeDirectConnectInterfacesBgpStatusResponse, error) {
	statusCode, msg, err := c.DescribeDirectConnectInterfacesBgpStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirectConnectInterfacesBgpStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirectConnectInterfacesBgpStatusWithContext(ctx context.Context, request *DescribeDirectConnectInterfacesBgpStatusRequest) string {
	if request == nil {
		request = NewDescribeDirectConnectInterfacesBgpStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDirectConnectInterfacesBgpStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirectConnectInterfacesBgpStatusWithContextV2(ctx context.Context, request *DescribeDirectConnectInterfacesBgpStatusRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirectConnectInterfacesBgpStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDirectConnectInterfacesBgpStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeactiveFlowLogRequest() (request *DeactiveFlowLogRequest) {
	request = &DeactiveFlowLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeactiveFlowLog")
	return
}

func NewDeactiveFlowLogResponse() (response *DeactiveFlowLogResponse) {
	response = &DeactiveFlowLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeactiveFlowLog(request *DeactiveFlowLogRequest) string {
	return c.DeactiveFlowLogWithContext(context.Background(), request)
}

func (c *Client) DeactiveFlowLogSend(request *DeactiveFlowLogRequest) (*DeactiveFlowLogResponse, error) {
	statusCode, msg, err := c.DeactiveFlowLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeactiveFlowLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeactiveFlowLogWithContext(ctx context.Context, request *DeactiveFlowLogRequest) string {
	if request == nil {
		request = NewDeactiveFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeactiveFlowLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeactiveFlowLogWithContextV2(ctx context.Context, request *DeactiveFlowLogRequest) (int, string, error) {
	if request == nil {
		request = NewDeactiveFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeactiveFlowLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActiveFlowLogRequest() (request *ActiveFlowLogRequest) {
	request = &ActiveFlowLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ActiveFlowLog")
	return
}

func NewActiveFlowLogResponse() (response *ActiveFlowLogResponse) {
	response = &ActiveFlowLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActiveFlowLog(request *ActiveFlowLogRequest) string {
	return c.ActiveFlowLogWithContext(context.Background(), request)
}

func (c *Client) ActiveFlowLogSend(request *ActiveFlowLogRequest) (*ActiveFlowLogResponse, error) {
	statusCode, msg, err := c.ActiveFlowLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ActiveFlowLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActiveFlowLogWithContext(ctx context.Context, request *ActiveFlowLogRequest) string {
	if request == nil {
		request = NewActiveFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewActiveFlowLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActiveFlowLogWithContextV2(ctx context.Context, request *ActiveFlowLogRequest) (int, string, error) {
	if request == nil {
		request = NewActiveFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewActiveFlowLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteFlowLogRequest() (request *DeleteFlowLogRequest) {
	request = &DeleteFlowLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteFlowLog")
	return
}

func NewDeleteFlowLogResponse() (response *DeleteFlowLogResponse) {
	response = &DeleteFlowLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFlowLog(request *DeleteFlowLogRequest) string {
	return c.DeleteFlowLogWithContext(context.Background(), request)
}

func (c *Client) DeleteFlowLogSend(request *DeleteFlowLogRequest) (*DeleteFlowLogResponse, error) {
	statusCode, msg, err := c.DeleteFlowLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteFlowLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFlowLogWithContext(ctx context.Context, request *DeleteFlowLogRequest) string {
	if request == nil {
		request = NewDeleteFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFlowLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFlowLogWithContextV2(ctx context.Context, request *DeleteFlowLogRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFlowLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyFlowLogRequest() (request *ModifyFlowLogRequest) {
	request = &ModifyFlowLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyFlowLog")
	return
}

func NewModifyFlowLogResponse() (response *ModifyFlowLogResponse) {
	response = &ModifyFlowLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyFlowLog(request *ModifyFlowLogRequest) string {
	return c.ModifyFlowLogWithContext(context.Background(), request)
}

func (c *Client) ModifyFlowLogSend(request *ModifyFlowLogRequest) (*ModifyFlowLogResponse, error) {
	statusCode, msg, err := c.ModifyFlowLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyFlowLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyFlowLogWithContext(ctx context.Context, request *ModifyFlowLogRequest) string {
	if request == nil {
		request = NewModifyFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyFlowLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyFlowLogWithContextV2(ctx context.Context, request *ModifyFlowLogRequest) (int, string, error) {
	if request == nil {
		request = NewModifyFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyFlowLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
	request = &DescribeFlowLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeFlowLogs")
	return
}

func NewDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
	response = &DescribeFlowLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) string {
	return c.DescribeFlowLogsWithContext(context.Background(), request)
}

func (c *Client) DescribeFlowLogsSend(request *DescribeFlowLogsRequest) (*DescribeFlowLogsResponse, error) {
	statusCode, msg, err := c.DescribeFlowLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFlowLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFlowLogsWithContext(ctx context.Context, request *DescribeFlowLogsRequest) string {
	if request == nil {
		request = NewDescribeFlowLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFlowLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFlowLogsWithContextV2(ctx context.Context, request *DescribeFlowLogsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFlowLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFlowLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateFlowLogRequest() (request *CreateFlowLogRequest) {
	request = &CreateFlowLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateFlowLog")
	return
}

func NewCreateFlowLogResponse() (response *CreateFlowLogResponse) {
	response = &CreateFlowLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFlowLog(request *CreateFlowLogRequest) string {
	return c.CreateFlowLogWithContext(context.Background(), request)
}

func (c *Client) CreateFlowLogSend(request *CreateFlowLogRequest) (*CreateFlowLogResponse, error) {
	statusCode, msg, err := c.CreateFlowLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateFlowLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateFlowLogWithContext(ctx context.Context, request *CreateFlowLogRequest) string {
	if request == nil {
		request = NewCreateFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFlowLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateFlowLogWithContextV2(ctx context.Context, request *CreateFlowLogRequest) (int, string, error) {
	if request == nil {
		request = NewCreateFlowLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFlowLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
