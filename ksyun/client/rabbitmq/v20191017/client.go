package v20191017

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2019-10-17"

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

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "CreateInstance")
	return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstance(request *CreateInstanceRequest) string {
	return c.CreateInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) string {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstance(request *DeleteInstanceRequest) string {
	return c.DeleteInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) string {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) string {
	return c.DescribeInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) string {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeInstance")
	return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstance(request *DescribeInstanceRequest) string {
	return c.DescribeInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) string {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceNodesRequest() (request *DescribeInstanceNodesRequest) {
	request = &DescribeInstanceNodesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeInstanceNodes")
	return
}

func NewDescribeInstanceNodesResponse() (response *DescribeInstanceNodesResponse) {
	response = &DescribeInstanceNodesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceNodes(request *DescribeInstanceNodesRequest) string {
	return c.DescribeInstanceNodesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceNodesWithContext(ctx context.Context, request *DescribeInstanceNodesRequest) string {
	if request == nil {
		request = NewDescribeInstanceNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceNodesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeValidRegionRequest() (request *DescribeValidRegionRequest) {
	request = &DescribeValidRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeValidRegion")
	return
}

func NewDescribeValidRegionResponse() (response *DescribeValidRegionResponse) {
	response = &DescribeValidRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeValidRegion(request *DescribeValidRegionRequest) string {
	return c.DescribeValidRegionWithContext(context.Background(), request)
}

func (c *Client) DescribeValidRegionWithContext(ctx context.Context, request *DescribeValidRegionRequest) string {
	if request == nil {
		request = NewDescribeValidRegionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeValidRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegions(request *DescribeRegionsRequest) string {
	return c.DescribeRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) string {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeSecurityGroupRulesRequest() (request *DescribeSecurityGroupRulesRequest) {
	request = &DescribeSecurityGroupRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeSecurityGroupRules")
	return
}

func NewDescribeSecurityGroupRulesResponse() (response *DescribeSecurityGroupRulesResponse) {
	response = &DescribeSecurityGroupRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroupRules(request *DescribeSecurityGroupRulesRequest) string {
	return c.DescribeSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupRulesWithContext(ctx context.Context, request *DescribeSecurityGroupRulesRequest) string {
	if request == nil {
		request = NewDescribeSecurityGroupRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSecurityGroupRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAddSecurityGroupRuleRequest() (request *AddSecurityGroupRuleRequest) {
	request = &AddSecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "AddSecurityGroupRule")
	return
}

func NewAddSecurityGroupRuleResponse() (response *AddSecurityGroupRuleResponse) {
	response = &AddSecurityGroupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddSecurityGroupRule(request *AddSecurityGroupRuleRequest) string {
	return c.AddSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) AddSecurityGroupRuleWithContext(ctx context.Context, request *AddSecurityGroupRuleRequest) string {
	if request == nil {
		request = NewAddSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddSecurityGroupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteSecurityGroupRulesRequest() (request *DeleteSecurityGroupRulesRequest) {
	request = &DeleteSecurityGroupRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DeleteSecurityGroupRules")
	return
}

func NewDeleteSecurityGroupRulesResponse() (response *DeleteSecurityGroupRulesResponse) {
	response = &DeleteSecurityGroupRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroupRules(request *DeleteSecurityGroupRulesRequest) string {
	return c.DeleteSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupRulesWithContext(ctx context.Context, request *DeleteSecurityGroupRulesRequest) string {
	if request == nil {
		request = NewDeleteSecurityGroupRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewResetPasswordRequest() (request *ResetPasswordRequest) {
	request = &ResetPasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "ResetPassword")
	return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
	response = &ResetPasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetPassword(request *ResetPasswordRequest) string {
	return c.ResetPasswordWithContext(context.Background(), request)
}

func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) string {
	if request == nil {
		request = NewResetPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRenameRequest() (request *RenameRequest) {
	request = &RenameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "Rename")
	return
}

func NewRenameResponse() (response *RenameResponse) {
	response = &RenameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Rename(request *RenameRequest) string {
	return c.RenameWithContext(context.Background(), request)
}

func (c *Client) RenameWithContext(ctx context.Context, request *RenameRequest) string {
	if request == nil {
		request = NewRenameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAllocateEipRequest() (request *AllocateEipRequest) {
	request = &AllocateEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "AllocateEip")
	return
}

func NewAllocateEipResponse() (response *AllocateEipResponse) {
	response = &AllocateEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateEip(request *AllocateEipRequest) string {
	return c.AllocateEipWithContext(context.Background(), request)
}

func (c *Client) AllocateEipWithContext(ctx context.Context, request *AllocateEipRequest) string {
	if request == nil {
		request = NewAllocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeallocateEipRequest() (request *DeallocateEipRequest) {
	request = &DeallocateEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DeallocateEip")
	return
}

func NewDeallocateEipResponse() (response *DeallocateEipResponse) {
	response = &DeallocateEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeallocateEip(request *DeallocateEipRequest) string {
	return c.DeallocateEipWithContext(context.Background(), request)
}

func (c *Client) DeallocateEipWithContext(ctx context.Context, request *DeallocateEipRequest) string {
	if request == nil {
		request = NewDeallocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSupportPluginsRequest() (request *SupportPluginsRequest) {
	request = &SupportPluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "SupportPlugins")
	return
}

func NewSupportPluginsResponse() (response *SupportPluginsResponse) {
	response = &SupportPluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SupportPlugins(request *SupportPluginsRequest) string {
	return c.SupportPluginsWithContext(context.Background(), request)
}

func (c *Client) SupportPluginsWithContext(ctx context.Context, request *SupportPluginsRequest) string {
	if request == nil {
		request = NewSupportPluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSupportPluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
	request = &RestartInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "RestartInstance")
	return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
	response = &RestartInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestartInstance(request *RestartInstanceRequest) string {
	return c.RestartInstanceWithContext(context.Background(), request)
}

func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) string {
	if request == nil {
		request = NewRestartInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestartInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListInstancePluginsRequest() (request *ListInstancePluginsRequest) {
	request = &ListInstancePluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "ListInstancePlugins")
	return
}

func NewListInstancePluginsResponse() (response *ListInstancePluginsResponse) {
	response = &ListInstancePluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstancePlugins(request *ListInstancePluginsRequest) string {
	return c.ListInstancePluginsWithContext(context.Background(), request)
}

func (c *Client) ListInstancePluginsWithContext(ctx context.Context, request *ListInstancePluginsRequest) string {
	if request == nil {
		request = NewListInstancePluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstancePluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewEnableInstancePluginsRequest() (request *EnableInstancePluginsRequest) {
	request = &EnableInstancePluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "EnableInstancePlugins")
	return
}

func NewEnableInstancePluginsResponse() (response *EnableInstancePluginsResponse) {
	response = &EnableInstancePluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableInstancePlugins(request *EnableInstancePluginsRequest) string {
	return c.EnableInstancePluginsWithContext(context.Background(), request)
}

func (c *Client) EnableInstancePluginsWithContext(ctx context.Context, request *EnableInstancePluginsRequest) string {
	if request == nil {
		request = NewEnableInstancePluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableInstancePluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisableInstancePluginsRequest() (request *DisableInstancePluginsRequest) {
	request = &DisableInstancePluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rabbitmq", APIVersion, "DisableInstancePlugins")
	return
}

func NewDisableInstancePluginsResponse() (response *DisableInstancePluginsResponse) {
	response = &DisableInstancePluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableInstancePlugins(request *DisableInstancePluginsRequest) string {
	return c.DisableInstancePluginsWithContext(context.Background(), request)
}

func (c *Client) DisableInstancePluginsWithContext(ctx context.Context, request *DisableInstancePluginsRequest) string {
	if request == nil {
		request = NewDisableInstancePluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableInstancePluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
