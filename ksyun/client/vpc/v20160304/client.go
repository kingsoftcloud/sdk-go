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
func NewCreateDcNatIpRequest() (request *CreateDcNatIpRequest) {
	request = &CreateDcNatIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDcNatIp")
	return
}

func NewCreateDcNatIpResponse() (response *CreateDcNatIpResponse) {
	response = &CreateDcNatIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDcNatIp(request *CreateDcNatIpRequest) string {
	return c.CreateDcNatIpWithContext(context.Background(), request)
}

func (c *Client) CreateDcNatIpWithContext(ctx context.Context, request *CreateDcNatIpRequest) string {
	if request == nil {
		request = NewCreateDcNatIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDcNatIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDcNatIpRequest() (request *DeleteDcNatIpRequest) {
	request = &DeleteDcNatIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDcNatIp")
	return
}

func NewDeleteDcNatIpResponse() (response *DeleteDcNatIpResponse) {
	response = &DeleteDcNatIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDcNatIp(request *DeleteDcNatIpRequest) string {
	return c.DeleteDcNatIpWithContext(context.Background(), request)
}

func (c *Client) DeleteDcNatIpWithContext(ctx context.Context, request *DeleteDcNatIpRequest) string {
	if request == nil {
		request = NewDeleteDcNatIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDcNatIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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
