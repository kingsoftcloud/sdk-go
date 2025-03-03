package v20210101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2021-01-01"

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

func NewListInstanceRequest() (request *ListInstanceRequest) {
	request = &ListInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ListInstance")
	return
}

func NewListInstanceResponse() (response *ListInstanceResponse) {
	response = &ListInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstance(request *ListInstanceRequest) string {
	return c.ListInstanceWithContext(context.Background(), request)
}

func (c *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest) string {
	if request == nil {
		request = NewListInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceResponse()
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
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeInstance")
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
func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "CreateInstance")
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
	request.Init().WithApiInfo("clickhouse", APIVersion, "DeleteInstance")
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
	request.SetContentType("application/json")

	response := NewDeleteInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
	request = &RenameInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "RenameInstance")
	return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
	response = &RenameInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameInstance(request *RenameInstanceRequest) string {
	return c.RenameInstanceWithContext(context.Background(), request)
}

func (c *Client) RenameInstanceWithContext(ctx context.Context, request *RenameInstanceRequest) string {
	if request == nil {
		request = NewRenameInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenameInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListSecurityGroupRequest() (request *ListSecurityGroupRequest) {
	request = &ListSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ListSecurityGroup")
	return
}

func NewListSecurityGroupResponse() (response *ListSecurityGroupResponse) {
	response = &ListSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSecurityGroup(request *ListSecurityGroupRequest) string {
	return c.ListSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ListSecurityGroupWithContext(ctx context.Context, request *ListSecurityGroupRequest) string {
	if request == nil {
		request = NewListSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeSecurityGroup")
	return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
	response = &DescribeSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) string {
	return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) string {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSecurityGroupResponse()
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
	request.Init().WithApiInfo("clickhouse", APIVersion, "CreateSecurityGroup")
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
	request.Init().WithApiInfo("clickhouse", APIVersion, "DeleteSecurityGroup")
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
	request.SetContentType("application/json")

	response := NewDeleteSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRenameSecurityGroupRequest() (request *RenameSecurityGroupRequest) {
	request = &RenameSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "RenameSecurityGroup")
	return
}

func NewRenameSecurityGroupResponse() (response *RenameSecurityGroupResponse) {
	response = &RenameSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameSecurityGroup(request *RenameSecurityGroupRequest) string {
	return c.RenameSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) RenameSecurityGroupWithContext(ctx context.Context, request *RenameSecurityGroupRequest) string {
	if request == nil {
		request = NewRenameSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenameSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
	request = &CloneSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "CloneSecurityGroup")
	return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
	response = &CloneSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) string {
	return c.CloneSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) string {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewBindSecurityGroupRequest() (request *BindSecurityGroupRequest) {
	request = &BindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "BindSecurityGroup")
	return
}

func NewBindSecurityGroupResponse() (response *BindSecurityGroupResponse) {
	response = &BindSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindSecurityGroup(request *BindSecurityGroupRequest) string {
	return c.BindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) BindSecurityGroupWithContext(ctx context.Context, request *BindSecurityGroupRequest) string {
	if request == nil {
		request = NewBindSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBindSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUnbindSecurityGroupRequest() (request *UnbindSecurityGroupRequest) {
	request = &UnbindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "UnbindSecurityGroup")
	return
}

func NewUnbindSecurityGroupResponse() (response *UnbindSecurityGroupResponse) {
	response = &UnbindSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindSecurityGroup(request *UnbindSecurityGroupRequest) string {
	return c.UnbindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) UnbindSecurityGroupWithContext(ctx context.Context, request *UnbindSecurityGroupRequest) string {
	if request == nil {
		request = NewUnbindSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnbindSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateSecurityRuleRequest() (request *CreateSecurityRuleRequest) {
	request = &CreateSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "CreateSecurityRule")
	return
}

func NewCreateSecurityRuleResponse() (response *CreateSecurityRuleResponse) {
	response = &CreateSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityRule(request *CreateSecurityRuleRequest) string {
	return c.CreateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityRuleWithContext(ctx context.Context, request *CreateSecurityRuleRequest) string {
	if request == nil {
		request = NewCreateSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteSecurityRuleRequest() (request *DeleteSecurityRuleRequest) {
	request = &DeleteSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DeleteSecurityRule")
	return
}

func NewDeleteSecurityRuleResponse() (response *DeleteSecurityRuleResponse) {
	response = &DeleteSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityRule(request *DeleteSecurityRuleRequest) string {
	return c.DeleteSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityRuleWithContext(ctx context.Context, request *DeleteSecurityRuleRequest) string {
	if request == nil {
		request = NewDeleteSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListSecuredInstanceRequest() (request *ListSecuredInstanceRequest) {
	request = &ListSecuredInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ListSecuredInstance")
	return
}

func NewListSecuredInstanceResponse() (response *ListSecuredInstanceResponse) {
	response = &ListSecuredInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSecuredInstance(request *ListSecuredInstanceRequest) string {
	return c.ListSecuredInstanceWithContext(context.Background(), request)
}

func (c *Client) ListSecuredInstanceWithContext(ctx context.Context, request *ListSecuredInstanceRequest) string {
	if request == nil {
		request = NewListSecuredInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecuredInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListUnsecuredInstanceRequest() (request *ListUnsecuredInstanceRequest) {
	request = &ListUnsecuredInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ListUnsecuredInstance")
	return
}

func NewListUnsecuredInstanceResponse() (response *ListUnsecuredInstanceResponse) {
	response = &ListUnsecuredInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListUnsecuredInstance(request *ListUnsecuredInstanceRequest) string {
	return c.ListUnsecuredInstanceWithContext(context.Background(), request)
}

func (c *Client) ListUnsecuredInstanceWithContext(ctx context.Context, request *ListUnsecuredInstanceRequest) string {
	if request == nil {
		request = NewListUnsecuredInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListUnsecuredInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListRecycledInstanceRequest() (request *ListRecycledInstanceRequest) {
	request = &ListRecycledInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ListRecycledInstance")
	return
}

func NewListRecycledInstanceResponse() (response *ListRecycledInstanceResponse) {
	response = &ListRecycledInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListRecycledInstance(request *ListRecycledInstanceRequest) string {
	return c.ListRecycledInstanceWithContext(context.Background(), request)
}

func (c *Client) ListRecycledInstanceWithContext(ctx context.Context, request *ListRecycledInstanceRequest) string {
	if request == nil {
		request = NewListRecycledInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListRecycledInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRecoverRecycledInstanceRequest() (request *RecoverRecycledInstanceRequest) {
	request = &RecoverRecycledInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "RecoverRecycledInstance")
	return
}

func NewRecoverRecycledInstanceResponse() (response *RecoverRecycledInstanceResponse) {
	response = &RecoverRecycledInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RecoverRecycledInstance(request *RecoverRecycledInstanceRequest) string {
	return c.RecoverRecycledInstanceWithContext(context.Background(), request)
}

func (c *Client) RecoverRecycledInstanceWithContext(ctx context.Context, request *RecoverRecycledInstanceRequest) string {
	if request == nil {
		request = NewRecoverRecycledInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRecoverRecycledInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDropRecycledInstanceRequest() (request *DropRecycledInstanceRequest) {
	request = &DropRecycledInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DropRecycledInstance")
	return
}

func NewDropRecycledInstanceResponse() (response *DropRecycledInstanceResponse) {
	response = &DropRecycledInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DropRecycledInstance(request *DropRecycledInstanceRequest) string {
	return c.DropRecycledInstanceWithContext(context.Background(), request)
}

func (c *Client) DropRecycledInstanceWithContext(ctx context.Context, request *DropRecycledInstanceRequest) string {
	if request == nil {
		request = NewDropRecycledInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDropRecycledInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListRegionRequest() (request *ListRegionRequest) {
	request = &ListRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ListRegion")
	return
}

func NewListRegionResponse() (response *ListRegionResponse) {
	response = &ListRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListRegion(request *ListRegionRequest) string {
	return c.ListRegionWithContext(context.Background(), request)
}

func (c *Client) ListRegionWithContext(ctx context.Context, request *ListRegionRequest) string {
	if request == nil {
		request = NewListRegionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescRegionRequest() (request *DescRegionRequest) {
	request = &DescRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescRegion")
	return
}

func NewDescRegionResponse() (response *DescRegionResponse) {
	response = &DescRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescRegion(request *DescRegionRequest) string {
	return c.DescRegionWithContext(context.Background(), request)
}

func (c *Client) DescRegionWithContext(ctx context.Context, request *DescRegionRequest) string {
	if request == nil {
		request = NewDescRegionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateSecurityRuleRequest() (request *UpdateSecurityRuleRequest) {
	request = &UpdateSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "UpdateSecurityRule")
	return
}

func NewUpdateSecurityRuleResponse() (response *UpdateSecurityRuleResponse) {
	response = &UpdateSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateSecurityRule(request *UpdateSecurityRuleRequest) string {
	return c.UpdateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) UpdateSecurityRuleWithContext(ctx context.Context, request *UpdateSecurityRuleRequest) string {
	if request == nil {
		request = NewUpdateSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRebindSecurityGroupRequest() (request *RebindSecurityGroupRequest) {
	request = &RebindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "RebindSecurityGroup")
	return
}

func NewRebindSecurityGroupResponse() (response *RebindSecurityGroupResponse) {
	response = &RebindSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebindSecurityGroup(request *RebindSecurityGroupRequest) string {
	return c.RebindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) RebindSecurityGroupWithContext(ctx context.Context, request *RebindSecurityGroupRequest) string {
	if request == nil {
		request = NewRebindSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebindSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeEngineDefaultParametersRequest() (request *DescribeEngineDefaultParametersRequest) {
	request = &DescribeEngineDefaultParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeEngineDefaultParameters")
	return
}

func NewDescribeEngineDefaultParametersResponse() (response *DescribeEngineDefaultParametersResponse) {
	response = &DescribeEngineDefaultParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEngineDefaultParameters(request *DescribeEngineDefaultParametersRequest) string {
	return c.DescribeEngineDefaultParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeEngineDefaultParametersWithContext(ctx context.Context, request *DescribeEngineDefaultParametersRequest) string {
	if request == nil {
		request = NewDescribeEngineDefaultParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeEngineDefaultParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDBParameterGroupRequest() (request *ModifyDBParameterGroupRequest) {
	request = &ModifyDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ModifyDBParameterGroup")
	return
}

func NewModifyDBParameterGroupResponse() (response *ModifyDBParameterGroupResponse) {
	response = &ModifyDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBParameterGroup(request *ModifyDBParameterGroupRequest) string {
	return c.ModifyDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyDBParameterGroupWithContext(ctx context.Context, request *ModifyDBParameterGroupRequest) string {
	if request == nil {
		request = NewModifyDBParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDBInstanceParametersRequest() (request *DescribeDBInstanceParametersRequest) {
	request = &DescribeDBInstanceParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeDBInstanceParameters")
	return
}

func NewDescribeDBInstanceParametersResponse() (response *DescribeDBInstanceParametersResponse) {
	response = &DescribeDBInstanceParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstanceParameters(request *DescribeDBInstanceParametersRequest) string {
	return c.DescribeDBInstanceParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceParametersWithContext(ctx context.Context, request *DescribeDBInstanceParametersRequest) string {
	if request == nil {
		request = NewDescribeDBInstanceParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBInstanceParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewResetDBParameterRequest() (request *ResetDBParameterRequest) {
	request = &ResetDBParameterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ResetDBParameter")
	return
}

func NewResetDBParameterResponse() (response *ResetDBParameterResponse) {
	response = &ResetDBParameterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetDBParameter(request *ResetDBParameterRequest) string {
	return c.ResetDBParameterWithContext(context.Background(), request)
}

func (c *Client) ResetDBParameterWithContext(ctx context.Context, request *ResetDBParameterRequest) string {
	if request == nil {
		request = NewResetDBParameterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResetDBParameterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeBucketsRequest() (request *DescribeBucketsRequest) {
	request = &DescribeBucketsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeBuckets")
	return
}

func NewDescribeBucketsResponse() (response *DescribeBucketsResponse) {
	response = &DescribeBucketsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBuckets(request *DescribeBucketsRequest) string {
	return c.DescribeBucketsWithContext(context.Background(), request)
}

func (c *Client) DescribeBucketsWithContext(ctx context.Context, request *DescribeBucketsRequest) string {
	if request == nil {
		request = NewDescribeBucketsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeBucketsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewOperateHotAndColdSeparationRequest() (request *OperateHotAndColdSeparationRequest) {
	request = &OperateHotAndColdSeparationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "OperateHotAndColdSeparation")
	return
}

func NewOperateHotAndColdSeparationResponse() (response *OperateHotAndColdSeparationResponse) {
	response = &OperateHotAndColdSeparationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperateHotAndColdSeparation(request *OperateHotAndColdSeparationRequest) string {
	return c.OperateHotAndColdSeparationWithContext(context.Background(), request)
}

func (c *Client) OperateHotAndColdSeparationWithContext(ctx context.Context, request *OperateHotAndColdSeparationRequest) string {
	if request == nil {
		request = NewOperateHotAndColdSeparationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperateHotAndColdSeparationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateInstanceAccountRequest() (request *CreateInstanceAccountRequest) {
	request = &CreateInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "CreateInstanceAccount")
	return
}

func NewCreateInstanceAccountResponse() (response *CreateInstanceAccountResponse) {
	response = &CreateInstanceAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstanceAccount(request *CreateInstanceAccountRequest) string {
	return c.CreateInstanceAccountWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceAccountWithContext(ctx context.Context, request *CreateInstanceAccountRequest) string {
	if request == nil {
		request = NewCreateInstanceAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceAccountPrivilegesRequest() (request *ModifyInstanceAccountPrivilegesRequest) {
	request = &ModifyInstanceAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ModifyInstanceAccountPrivileges")
	return
}

func NewModifyInstanceAccountPrivilegesResponse() (response *ModifyInstanceAccountPrivilegesResponse) {
	response = &ModifyInstanceAccountPrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAccountPrivileges(request *ModifyInstanceAccountPrivilegesRequest) string {
	return c.ModifyInstanceAccountPrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountPrivilegesWithContext(ctx context.Context, request *ModifyInstanceAccountPrivilegesRequest) string {
	if request == nil {
		request = NewModifyInstanceAccountPrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceAccountPrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInstanceAccountRequest() (request *DeleteInstanceAccountRequest) {
	request = &DeleteInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DeleteInstanceAccount")
	return
}

func NewDeleteInstanceAccountResponse() (response *DeleteInstanceAccountResponse) {
	response = &DeleteInstanceAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstanceAccount(request *DeleteInstanceAccountRequest) string {
	return c.DeleteInstanceAccountWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceAccountWithContext(ctx context.Context, request *DeleteInstanceAccountRequest) string {
	if request == nil {
		request = NewDeleteInstanceAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceAccountsRequest() (request *DescribeInstanceAccountsRequest) {
	request = &DescribeInstanceAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeInstanceAccounts")
	return
}

func NewDescribeInstanceAccountsResponse() (response *DescribeInstanceAccountsResponse) {
	response = &DescribeInstanceAccountsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceAccounts(request *DescribeInstanceAccountsRequest) string {
	return c.DescribeInstanceAccountsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceAccountsWithContext(ctx context.Context, request *DescribeInstanceAccountsRequest) string {
	if request == nil {
		request = NewDescribeInstanceAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceAccountsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceDatabasesRequest() (request *DescribeInstanceDatabasesRequest) {
	request = &DescribeInstanceDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "DescribeInstanceDatabases")
	return
}

func NewDescribeInstanceDatabasesResponse() (response *DescribeInstanceDatabasesResponse) {
	response = &DescribeInstanceDatabasesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceDatabases(request *DescribeInstanceDatabasesRequest) string {
	return c.DescribeInstanceDatabasesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceDatabasesWithContext(ctx context.Context, request *DescribeInstanceDatabasesRequest) string {
	if request == nil {
		request = NewDescribeInstanceDatabasesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceDatabasesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceAccountInfoRequest() (request *ModifyInstanceAccountInfoRequest) {
	request = &ModifyInstanceAccountInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clickhouse", APIVersion, "ModifyInstanceAccountInfo")
	return
}

func NewModifyInstanceAccountInfoResponse() (response *ModifyInstanceAccountInfoResponse) {
	response = &ModifyInstanceAccountInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAccountInfo(request *ModifyInstanceAccountInfoRequest) string {
	return c.ModifyInstanceAccountInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountInfoWithContext(ctx context.Context, request *ModifyInstanceAccountInfoRequest) string {
	if request == nil {
		request = NewModifyInstanceAccountInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceAccountInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
