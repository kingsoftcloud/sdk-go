package v20190425

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2019-04-25"

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

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDBInstance")
	return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) string {
	return c.CreateDBInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateDBInstanceSend(request *CreateDBInstanceRequest) (*CreateDBInstanceResponse, error) {
	statusCode, msg, err := c.CreateDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) string {
	if request == nil {
		request = NewCreateDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDBInstanceWithContextV2(ctx context.Context, request *CreateDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstances")
	return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
	response = &DescribeDBInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) string {
	return c.DescribeDBInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstancesSend(request *DescribeDBInstancesRequest) (*DescribeDBInstancesResponse, error) {
	statusCode, msg, err := c.DescribeDBInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) string {
	if request == nil {
		request = NewDescribeDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBInstancesWithContextV2(ctx context.Context, request *DescribeDBInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
	request = &DeleteDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDBInstance")
	return
}

func NewDeleteDBInstanceResponse() (response *DeleteDBInstanceResponse) {
	response = &DeleteDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) string {
	return c.DeleteDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteDBInstanceSend(request *DeleteDBInstanceRequest) (*DeleteDBInstanceResponse, error) {
	statusCode, msg, err := c.DeleteDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDBInstanceWithContext(ctx context.Context, request *DeleteDBInstanceRequest) string {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDBInstanceWithContextV2(ctx context.Context, request *DeleteDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceRequest() (request *ModifyDBInstanceRequest) {
	request = &ModifyDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstance")
	return
}

func NewModifyDBInstanceResponse() (response *ModifyDBInstanceResponse) {
	response = &ModifyDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstance(request *ModifyDBInstanceRequest) string {
	return c.ModifyDBInstanceWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceSend(request *ModifyDBInstanceRequest) (*ModifyDBInstanceResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceWithContext(ctx context.Context, request *ModifyDBInstanceRequest) string {
	if request == nil {
		request = NewModifyDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceWithContextV2(ctx context.Context, request *ModifyDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceResponse()
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
	request.Init().WithApiInfo("sqlserver", APIVersion, "CreateSecurityGroup")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewCreateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeSecurityGroup")
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

func (c *Client) DescribeSecurityGroupSend(request *DescribeSecurityGroupRequest) (*DescribeSecurityGroupResponse, error) {
	statusCode, msg, err := c.DescribeSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeSecurityGroupWithContextV2(ctx context.Context, request *DescribeSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSecurityGroupResponse()
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
	request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteSecurityGroup")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewDeleteSecurityGroupResponse()
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
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifySecurityGroup")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewModifySecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
	request = &CloneSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "CloneSecurityGroup")
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

func (c *Client) CloneSecurityGroupSend(request *CloneSecurityGroupRequest) (*CloneSecurityGroupResponse, error) {
	statusCode, msg, err := c.CloneSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CloneSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) string {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloneSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloneSecurityGroupWithContextV2(ctx context.Context, request *CloneSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloneSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRuleRequest() (request *ModifySecurityGroupRuleRequest) {
	request = &ModifySecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifySecurityGroupRule")
	return
}

func NewModifySecurityGroupRuleResponse() (response *ModifySecurityGroupRuleResponse) {
	response = &ModifySecurityGroupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroupRule(request *ModifySecurityGroupRuleRequest) string {
	return c.ModifySecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupRuleSend(request *ModifySecurityGroupRuleRequest) (*ModifySecurityGroupRuleResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySecurityGroupRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupRuleWithContext(ctx context.Context, request *ModifySecurityGroupRuleRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifySecurityGroupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupRuleWithContextV2(ctx context.Context, request *ModifySecurityGroupRuleRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifySecurityGroupRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSecurityGroupRelationRequest() (request *SecurityGroupRelationRequest) {
	request = &SecurityGroupRelationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "SecurityGroupRelation")
	return
}

func NewSecurityGroupRelationResponse() (response *SecurityGroupRelationResponse) {
	response = &SecurityGroupRelationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SecurityGroupRelation(request *SecurityGroupRelationRequest) string {
	return c.SecurityGroupRelationWithContext(context.Background(), request)
}

func (c *Client) SecurityGroupRelationSend(request *SecurityGroupRelationRequest) (*SecurityGroupRelationResponse, error) {
	statusCode, msg, err := c.SecurityGroupRelationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SecurityGroupRelationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SecurityGroupRelationWithContext(ctx context.Context, request *SecurityGroupRelationRequest) string {
	if request == nil {
		request = NewSecurityGroupRelationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSecurityGroupRelationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SecurityGroupRelationWithContextV2(ctx context.Context, request *SecurityGroupRelationRequest) (int, string, error) {
	if request == nil {
		request = NewSecurityGroupRelationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSecurityGroupRelationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRuleNameRequest() (request *ModifySecurityGroupRuleNameRequest) {
	request = &ModifySecurityGroupRuleNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifySecurityGroupRuleName")
	return
}

func NewModifySecurityGroupRuleNameResponse() (response *ModifySecurityGroupRuleNameResponse) {
	response = &ModifySecurityGroupRuleNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroupRuleName(request *ModifySecurityGroupRuleNameRequest) string {
	return c.ModifySecurityGroupRuleNameWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupRuleNameSend(request *ModifySecurityGroupRuleNameRequest) (*ModifySecurityGroupRuleNameResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupRuleNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySecurityGroupRuleNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupRuleNameWithContext(ctx context.Context, request *ModifySecurityGroupRuleNameRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRuleNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifySecurityGroupRuleNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupRuleNameWithContextV2(ctx context.Context, request *ModifySecurityGroupRuleNameRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRuleNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifySecurityGroupRuleNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCollationsRequest() (request *DescribeCollationsRequest) {
	request = &DescribeCollationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCollations")
	return
}

func NewDescribeCollationsResponse() (response *DescribeCollationsResponse) {
	response = &DescribeCollationsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCollations(request *DescribeCollationsRequest) string {
	return c.DescribeCollationsWithContext(context.Background(), request)
}

func (c *Client) DescribeCollationsSend(request *DescribeCollationsRequest) (*DescribeCollationsResponse, error) {
	statusCode, msg, err := c.DescribeCollationsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCollationsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCollationsWithContext(ctx context.Context, request *DescribeCollationsRequest) string {
	if request == nil {
		request = NewDescribeCollationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCollationsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCollationsWithContextV2(ctx context.Context, request *DescribeCollationsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCollationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCollationsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInstanceDatabaseRequest() (request *CreateInstanceDatabaseRequest) {
	request = &CreateInstanceDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "CreateInstanceDatabase")
	return
}

func NewCreateInstanceDatabaseResponse() (response *CreateInstanceDatabaseResponse) {
	response = &CreateInstanceDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstanceDatabase(request *CreateInstanceDatabaseRequest) string {
	return c.CreateInstanceDatabaseWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceDatabaseSend(request *CreateInstanceDatabaseRequest) (*CreateInstanceDatabaseResponse, error) {
	statusCode, msg, err := c.CreateInstanceDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInstanceDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInstanceDatabaseWithContext(ctx context.Context, request *CreateInstanceDatabaseRequest) string {
	if request == nil {
		request = NewCreateInstanceDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInstanceDatabaseWithContextV2(ctx context.Context, request *CreateInstanceDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceDatabasePrivilegesRequest() (request *ModifyInstanceDatabasePrivilegesRequest) {
	request = &ModifyInstanceDatabasePrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceDatabasePrivileges")
	return
}

func NewModifyInstanceDatabasePrivilegesResponse() (response *ModifyInstanceDatabasePrivilegesResponse) {
	response = &ModifyInstanceDatabasePrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabasePrivileges(request *ModifyInstanceDatabasePrivilegesRequest) string {
	return c.ModifyInstanceDatabasePrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabasePrivilegesSend(request *ModifyInstanceDatabasePrivilegesRequest) (*ModifyInstanceDatabasePrivilegesResponse, error) {
	statusCode, msg, err := c.ModifyInstanceDatabasePrivilegesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceDatabasePrivilegesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceDatabasePrivilegesWithContext(ctx context.Context, request *ModifyInstanceDatabasePrivilegesRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabasePrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabasePrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceDatabasePrivilegesWithContextV2(ctx context.Context, request *ModifyInstanceDatabasePrivilegesRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceDatabasePrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabasePrivilegesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceDatabasesRequest() (request *DescribeInstanceDatabasesRequest) {
	request = &DescribeInstanceDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceDatabases")
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

func (c *Client) DescribeInstanceDatabasesSend(request *DescribeInstanceDatabasesRequest) (*DescribeInstanceDatabasesResponse, error) {
	statusCode, msg, err := c.DescribeInstanceDatabasesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstanceDatabasesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInstanceDatabasesWithContextV2(ctx context.Context, request *DescribeInstanceDatabasesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceDatabasesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceDatabasesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInstanceAccountRequest() (request *CreateInstanceAccountRequest) {
	request = &CreateInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "CreateInstanceAccount")
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

func (c *Client) CreateInstanceAccountSend(request *CreateInstanceAccountRequest) (*CreateInstanceAccountResponse, error) {
	statusCode, msg, err := c.CreateInstanceAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInstanceAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateInstanceAccountWithContextV2(ctx context.Context, request *CreateInstanceAccountRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceAccountsRequest() (request *DescribeInstanceAccountsRequest) {
	request = &DescribeInstanceAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceAccounts")
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

func (c *Client) DescribeInstanceAccountsSend(request *DescribeInstanceAccountsRequest) (*DescribeInstanceAccountsResponse, error) {
	statusCode, msg, err := c.DescribeInstanceAccountsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstanceAccountsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInstanceAccountsWithContextV2(ctx context.Context, request *DescribeInstanceAccountsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceAccountsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAccountInfoRequest() (request *ModifyInstanceAccountInfoRequest) {
	request = &ModifyInstanceAccountInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceAccountInfo")
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

func (c *Client) ModifyInstanceAccountInfoSend(request *ModifyInstanceAccountInfoRequest) (*ModifyInstanceAccountInfoResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAccountInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceAccountInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyInstanceAccountInfoWithContextV2(ctx context.Context, request *ModifyInstanceAccountInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAccountInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceAccountInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAccountPrivilegesRequest() (request *ModifyInstanceAccountPrivilegesRequest) {
	request = &ModifyInstanceAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceAccountPrivileges")
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

func (c *Client) ModifyInstanceAccountPrivilegesSend(request *ModifyInstanceAccountPrivilegesRequest) (*ModifyInstanceAccountPrivilegesResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAccountPrivilegesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceAccountPrivilegesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyInstanceAccountPrivilegesWithContextV2(ctx context.Context, request *ModifyInstanceAccountPrivilegesRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAccountPrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceAccountPrivilegesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceAccountRequest() (request *DeleteInstanceAccountRequest) {
	request = &DeleteInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteInstanceAccount")
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

func (c *Client) DeleteInstanceAccountSend(request *DeleteInstanceAccountRequest) (*DeleteInstanceAccountResponse, error) {
	statusCode, msg, err := c.DeleteInstanceAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInstanceAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteInstanceAccountWithContextV2(ctx context.Context, request *DeleteInstanceAccountRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceDatabaseRequest() (request *DeleteInstanceDatabaseRequest) {
	request = &DeleteInstanceDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteInstanceDatabase")
	return
}

func NewDeleteInstanceDatabaseResponse() (response *DeleteInstanceDatabaseResponse) {
	response = &DeleteInstanceDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstanceDatabase(request *DeleteInstanceDatabaseRequest) string {
	return c.DeleteInstanceDatabaseWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceDatabaseSend(request *DeleteInstanceDatabaseRequest) (*DeleteInstanceDatabaseResponse, error) {
	statusCode, msg, err := c.DeleteInstanceDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInstanceDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceDatabaseWithContext(ctx context.Context, request *DeleteInstanceDatabaseRequest) string {
	if request == nil {
		request = NewDeleteInstanceDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceDatabaseWithContextV2(ctx context.Context, request *DeleteInstanceDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceDatabaseInfoRequest() (request *ModifyInstanceDatabaseInfoRequest) {
	request = &ModifyInstanceDatabaseInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceDatabaseInfo")
	return
}

func NewModifyInstanceDatabaseInfoResponse() (response *ModifyInstanceDatabaseInfoResponse) {
	response = &ModifyInstanceDatabaseInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabaseInfo(request *ModifyInstanceDatabaseInfoRequest) string {
	return c.ModifyInstanceDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabaseInfoSend(request *ModifyInstanceDatabaseInfoRequest) (*ModifyInstanceDatabaseInfoResponse, error) {
	statusCode, msg, err := c.ModifyInstanceDatabaseInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceDatabaseInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceDatabaseInfoWithContext(ctx context.Context, request *ModifyInstanceDatabaseInfoRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabaseInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabaseInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceDatabaseInfoWithContextV2(ctx context.Context, request *ModifyInstanceDatabaseInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceDatabaseInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabaseInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOverrideDBInstanceRequest() (request *OverrideDBInstanceRequest) {
	request = &OverrideDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "OverrideDBInstance")
	return
}

func NewOverrideDBInstanceResponse() (response *OverrideDBInstanceResponse) {
	response = &OverrideDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OverrideDBInstance(request *OverrideDBInstanceRequest) string {
	return c.OverrideDBInstanceWithContext(context.Background(), request)
}

func (c *Client) OverrideDBInstanceSend(request *OverrideDBInstanceRequest) (*OverrideDBInstanceResponse, error) {
	statusCode, msg, err := c.OverrideDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OverrideDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OverrideDBInstanceWithContext(ctx context.Context, request *OverrideDBInstanceRequest) string {
	if request == nil {
		request = NewOverrideDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOverrideDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OverrideDBInstanceWithContextV2(ctx context.Context, request *OverrideDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewOverrideDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOverrideDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreDBInstanceFromDBBackupRequest() (request *RestoreDBInstanceFromDBBackupRequest) {
	request = &RestoreDBInstanceFromDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "RestoreDBInstanceFromDBBackup")
	return
}

func NewRestoreDBInstanceFromDBBackupResponse() (response *RestoreDBInstanceFromDBBackupResponse) {
	response = &RestoreDBInstanceFromDBBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreDBInstanceFromDBBackup(request *RestoreDBInstanceFromDBBackupRequest) string {
	return c.RestoreDBInstanceFromDBBackupWithContext(context.Background(), request)
}

func (c *Client) RestoreDBInstanceFromDBBackupSend(request *RestoreDBInstanceFromDBBackupRequest) (*RestoreDBInstanceFromDBBackupResponse, error) {
	statusCode, msg, err := c.RestoreDBInstanceFromDBBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreDBInstanceFromDBBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreDBInstanceFromDBBackupWithContext(ctx context.Context, request *RestoreDBInstanceFromDBBackupRequest) string {
	if request == nil {
		request = NewRestoreDBInstanceFromDBBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreDBInstanceFromDBBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreDBInstanceFromDBBackupWithContextV2(ctx context.Context, request *RestoreDBInstanceFromDBBackupRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreDBInstanceFromDBBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreDBInstanceFromDBBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDBBackupRequest() (request *CreateDBBackupRequest) {
	request = &CreateDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDBBackup")
	return
}

func NewCreateDBBackupResponse() (response *CreateDBBackupResponse) {
	response = &CreateDBBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBBackup(request *CreateDBBackupRequest) string {
	return c.CreateDBBackupWithContext(context.Background(), request)
}

func (c *Client) CreateDBBackupSend(request *CreateDBBackupRequest) (*CreateDBBackupResponse, error) {
	statusCode, msg, err := c.CreateDBBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDBBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDBBackupWithContext(ctx context.Context, request *CreateDBBackupRequest) string {
	if request == nil {
		request = NewCreateDBBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDBBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDBBackupWithContextV2(ctx context.Context, request *CreateDBBackupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDBBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDBBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDBBackupRequest() (request *DeleteDBBackupRequest) {
	request = &DeleteDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDBBackup")
	return
}

func NewDeleteDBBackupResponse() (response *DeleteDBBackupResponse) {
	response = &DeleteDBBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDBBackup(request *DeleteDBBackupRequest) string {
	return c.DeleteDBBackupWithContext(context.Background(), request)
}

func (c *Client) DeleteDBBackupSend(request *DeleteDBBackupRequest) (*DeleteDBBackupResponse, error) {
	statusCode, msg, err := c.DeleteDBBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDBBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDBBackupWithContext(ctx context.Context, request *DeleteDBBackupRequest) string {
	if request == nil {
		request = NewDeleteDBBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDBBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDBBackupWithContextV2(ctx context.Context, request *DeleteDBBackupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDBBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDBBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
	request = &DescribeDBBackupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBBackups")
	return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
	response = &DescribeDBBackupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) string {
	return c.DescribeDBBackupsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBBackupsSend(request *DescribeDBBackupsRequest) (*DescribeDBBackupsResponse, error) {
	statusCode, msg, err := c.DescribeDBBackupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBBackupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBBackupsWithContext(ctx context.Context, request *DescribeDBBackupsRequest) string {
	if request == nil {
		request = NewDescribeDBBackupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBBackupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBBackupsWithContextV2(ctx context.Context, request *DescribeDBBackupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBBackupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBBackupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBBackupPolicyRequest() (request *ModifyDBBackupPolicyRequest) {
	request = &ModifyDBBackupPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBBackupPolicy")
	return
}

func NewModifyDBBackupPolicyResponse() (response *ModifyDBBackupPolicyResponse) {
	response = &ModifyDBBackupPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBBackupPolicy(request *ModifyDBBackupPolicyRequest) string {
	return c.ModifyDBBackupPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyDBBackupPolicySend(request *ModifyDBBackupPolicyRequest) (*ModifyDBBackupPolicyResponse, error) {
	statusCode, msg, err := c.ModifyDBBackupPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBBackupPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBBackupPolicyWithContext(ctx context.Context, request *ModifyDBBackupPolicyRequest) string {
	if request == nil {
		request = NewModifyDBBackupPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBBackupPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBBackupPolicyWithContextV2(ctx context.Context, request *ModifyDBBackupPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBBackupPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBBackupPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateDBInstanceEipRequest() (request *AllocateDBInstanceEipRequest) {
	request = &AllocateDBInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "AllocateDBInstanceEip")
	return
}

func NewAllocateDBInstanceEipResponse() (response *AllocateDBInstanceEipResponse) {
	response = &AllocateDBInstanceEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateDBInstanceEip(request *AllocateDBInstanceEipRequest) string {
	return c.AllocateDBInstanceEipWithContext(context.Background(), request)
}

func (c *Client) AllocateDBInstanceEipSend(request *AllocateDBInstanceEipRequest) (*AllocateDBInstanceEipResponse, error) {
	statusCode, msg, err := c.AllocateDBInstanceEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AllocateDBInstanceEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateDBInstanceEipWithContext(ctx context.Context, request *AllocateDBInstanceEipRequest) string {
	if request == nil {
		request = NewAllocateDBInstanceEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAllocateDBInstanceEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateDBInstanceEipWithContextV2(ctx context.Context, request *AllocateDBInstanceEipRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateDBInstanceEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAllocateDBInstanceEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReleaseDBInstanceEipRequest() (request *ReleaseDBInstanceEipRequest) {
	request = &ReleaseDBInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ReleaseDBInstanceEip")
	return
}

func NewReleaseDBInstanceEipResponse() (response *ReleaseDBInstanceEipResponse) {
	response = &ReleaseDBInstanceEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReleaseDBInstanceEip(request *ReleaseDBInstanceEipRequest) string {
	return c.ReleaseDBInstanceEipWithContext(context.Background(), request)
}

func (c *Client) ReleaseDBInstanceEipSend(request *ReleaseDBInstanceEipRequest) (*ReleaseDBInstanceEipResponse, error) {
	statusCode, msg, err := c.ReleaseDBInstanceEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ReleaseDBInstanceEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReleaseDBInstanceEipWithContext(ctx context.Context, request *ReleaseDBInstanceEipRequest) string {
	if request == nil {
		request = NewReleaseDBInstanceEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewReleaseDBInstanceEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReleaseDBInstanceEipWithContextV2(ctx context.Context, request *ReleaseDBInstanceEipRequest) (int, string, error) {
	if request == nil {
		request = NewReleaseDBInstanceEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewReleaseDBInstanceEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListSlowLogsRequest() (request *ListSlowLogsRequest) {
	request = &ListSlowLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ListSlowLogs")
	return
}

func NewListSlowLogsResponse() (response *ListSlowLogsResponse) {
	response = &ListSlowLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSlowLogs(request *ListSlowLogsRequest) string {
	return c.ListSlowLogsWithContext(context.Background(), request)
}

func (c *Client) ListSlowLogsSend(request *ListSlowLogsRequest) (*ListSlowLogsResponse, error) {
	statusCode, msg, err := c.ListSlowLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListSlowLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListSlowLogsWithContext(ctx context.Context, request *ListSlowLogsRequest) string {
	if request == nil {
		request = NewListSlowLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSlowLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListSlowLogsWithContextV2(ctx context.Context, request *ListSlowLogsRequest) (int, string, error) {
	if request == nil {
		request = NewListSlowLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSlowLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListErrorLogsRequest() (request *ListErrorLogsRequest) {
	request = &ListErrorLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ListErrorLogs")
	return
}

func NewListErrorLogsResponse() (response *ListErrorLogsResponse) {
	response = &ListErrorLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListErrorLogs(request *ListErrorLogsRequest) string {
	return c.ListErrorLogsWithContext(context.Background(), request)
}

func (c *Client) ListErrorLogsSend(request *ListErrorLogsRequest) (*ListErrorLogsResponse, error) {
	statusCode, msg, err := c.ListErrorLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListErrorLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListErrorLogsWithContext(ctx context.Context, request *ListErrorLogsRequest) string {
	if request == nil {
		request = NewListErrorLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListErrorLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListErrorLogsWithContextV2(ctx context.Context, request *ListErrorLogsRequest) (int, string, error) {
	if request == nil {
		request = NewListErrorLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListErrorLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
	request = &ModifyDBInstanceSpecRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceSpec")
	return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
	response = &ModifyDBInstanceSpecResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) string {
	return c.ModifyDBInstanceSpecWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceSpecSend(request *ModifyDBInstanceSpecRequest) (*ModifyDBInstanceSpecResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceSpecWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceSpecResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest) string {
	if request == nil {
		request = NewModifyDBInstanceSpecRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBInstanceSpecResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceSpecWithContextV2(ctx context.Context, request *ModifyDBInstanceSpecRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceSpecRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBInstanceSpecResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImportTaskRequest() (request *DescribeImportTaskRequest) {
	request = &DescribeImportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeImportTask")
	return
}

func NewDescribeImportTaskResponse() (response *DescribeImportTaskResponse) {
	response = &DescribeImportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImportTask(request *DescribeImportTaskRequest) string {
	return c.DescribeImportTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeImportTaskSend(request *DescribeImportTaskRequest) (*DescribeImportTaskResponse, error) {
	statusCode, msg, err := c.DescribeImportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeImportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImportTaskWithContext(ctx context.Context, request *DescribeImportTaskRequest) string {
	if request == nil {
		request = NewDescribeImportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeImportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImportTaskWithContextV2(ctx context.Context, request *DescribeImportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeImportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImportFileRequest() (request *DescribeImportFileRequest) {
	request = &DescribeImportFileRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeImportFile")
	return
}

func NewDescribeImportFileResponse() (response *DescribeImportFileResponse) {
	response = &DescribeImportFileResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImportFile(request *DescribeImportFileRequest) string {
	return c.DescribeImportFileWithContext(context.Background(), request)
}

func (c *Client) DescribeImportFileSend(request *DescribeImportFileRequest) (*DescribeImportFileResponse, error) {
	statusCode, msg, err := c.DescribeImportFileWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeImportFileResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImportFileWithContext(ctx context.Context, request *DescribeImportFileRequest) string {
	if request == nil {
		request = NewDescribeImportFileRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeImportFileResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImportFileWithContextV2(ctx context.Context, request *DescribeImportFileRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImportFileRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeImportFileResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateImportTaskRequest() (request *CreateImportTaskRequest) {
	request = &CreateImportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "CreateImportTask")
	return
}

func NewCreateImportTaskResponse() (response *CreateImportTaskResponse) {
	response = &CreateImportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImportTask(request *CreateImportTaskRequest) string {
	return c.CreateImportTaskWithContext(context.Background(), request)
}

func (c *Client) CreateImportTaskSend(request *CreateImportTaskRequest) (*CreateImportTaskResponse, error) {
	statusCode, msg, err := c.CreateImportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateImportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateImportTaskWithContext(ctx context.Context, request *CreateImportTaskRequest) string {
	if request == nil {
		request = NewCreateImportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateImportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateImportTaskWithContextV2(ctx context.Context, request *CreateImportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateImportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateImportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewFinishImportTaskRequest() (request *FinishImportTaskRequest) {
	request = &FinishImportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "FinishImportTask")
	return
}

func NewFinishImportTaskResponse() (response *FinishImportTaskResponse) {
	response = &FinishImportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) FinishImportTask(request *FinishImportTaskRequest) string {
	return c.FinishImportTaskWithContext(context.Background(), request)
}

func (c *Client) FinishImportTaskSend(request *FinishImportTaskRequest) (*FinishImportTaskResponse, error) {
	statusCode, msg, err := c.FinishImportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct FinishImportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) FinishImportTaskWithContext(ctx context.Context, request *FinishImportTaskRequest) string {
	if request == nil {
		request = NewFinishImportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewFinishImportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) FinishImportTaskWithContextV2(ctx context.Context, request *FinishImportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewFinishImportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewFinishImportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBInstanceRestorableTimeRequest() (request *DescribeDBInstanceRestorableTimeRequest) {
	request = &DescribeDBInstanceRestorableTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstanceRestorableTime")
	return
}

func NewDescribeDBInstanceRestorableTimeResponse() (response *DescribeDBInstanceRestorableTimeResponse) {
	response = &DescribeDBInstanceRestorableTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstanceRestorableTime(request *DescribeDBInstanceRestorableTimeRequest) string {
	return c.DescribeDBInstanceRestorableTimeWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceRestorableTimeSend(request *DescribeDBInstanceRestorableTimeRequest) (*DescribeDBInstanceRestorableTimeResponse, error) {
	statusCode, msg, err := c.DescribeDBInstanceRestorableTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBInstanceRestorableTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBInstanceRestorableTimeWithContext(ctx context.Context, request *DescribeDBInstanceRestorableTimeRequest) string {
	if request == nil {
		request = NewDescribeDBInstanceRestorableTimeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBInstanceRestorableTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBInstanceRestorableTimeWithContextV2(ctx context.Context, request *DescribeDBInstanceRestorableTimeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBInstanceRestorableTimeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBInstanceRestorableTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetHistoryDatabaseInfoRequest() (request *GetHistoryDatabaseInfoRequest) {
	request = &GetHistoryDatabaseInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "GetHistoryDatabaseInfo")
	return
}

func NewGetHistoryDatabaseInfoResponse() (response *GetHistoryDatabaseInfoResponse) {
	response = &GetHistoryDatabaseInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetHistoryDatabaseInfo(request *GetHistoryDatabaseInfoRequest) string {
	return c.GetHistoryDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) GetHistoryDatabaseInfoSend(request *GetHistoryDatabaseInfoRequest) (*GetHistoryDatabaseInfoResponse, error) {
	statusCode, msg, err := c.GetHistoryDatabaseInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetHistoryDatabaseInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetHistoryDatabaseInfoWithContext(ctx context.Context, request *GetHistoryDatabaseInfoRequest) string {
	if request == nil {
		request = NewGetHistoryDatabaseInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetHistoryDatabaseInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetHistoryDatabaseInfoWithContextV2(ctx context.Context, request *GetHistoryDatabaseInfoRequest) (int, string, error) {
	if request == nil {
		request = NewGetHistoryDatabaseInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetHistoryDatabaseInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreToCurInstanceRequest() (request *RestoreToCurInstanceRequest) {
	request = &RestoreToCurInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "RestoreToCurInstance")
	return
}

func NewRestoreToCurInstanceResponse() (response *RestoreToCurInstanceResponse) {
	response = &RestoreToCurInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreToCurInstance(request *RestoreToCurInstanceRequest) string {
	return c.RestoreToCurInstanceWithContext(context.Background(), request)
}

func (c *Client) RestoreToCurInstanceSend(request *RestoreToCurInstanceRequest) (*RestoreToCurInstanceResponse, error) {
	statusCode, msg, err := c.RestoreToCurInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreToCurInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreToCurInstanceWithContext(ctx context.Context, request *RestoreToCurInstanceRequest) string {
	if request == nil {
		request = NewRestoreToCurInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreToCurInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreToCurInstanceWithContextV2(ctx context.Context, request *RestoreToCurInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreToCurInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreToCurInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceDatabaseNameRequest() (request *ModifyInstanceDatabaseNameRequest) {
	request = &ModifyInstanceDatabaseNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceDatabaseName")
	return
}

func NewModifyInstanceDatabaseNameResponse() (response *ModifyInstanceDatabaseNameResponse) {
	response = &ModifyInstanceDatabaseNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabaseName(request *ModifyInstanceDatabaseNameRequest) string {
	return c.ModifyInstanceDatabaseNameWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabaseNameSend(request *ModifyInstanceDatabaseNameRequest) (*ModifyInstanceDatabaseNameResponse, error) {
	statusCode, msg, err := c.ModifyInstanceDatabaseNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceDatabaseNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceDatabaseNameWithContext(ctx context.Context, request *ModifyInstanceDatabaseNameRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabaseNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabaseNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceDatabaseNameWithContextV2(ctx context.Context, request *ModifyInstanceDatabaseNameRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceDatabaseNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabaseNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRebootDBInstanceRequest() (request *RebootDBInstanceRequest) {
	request = &RebootDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "RebootDBInstance")
	return
}

func NewRebootDBInstanceResponse() (response *RebootDBInstanceResponse) {
	response = &RebootDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootDBInstance(request *RebootDBInstanceRequest) string {
	return c.RebootDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RebootDBInstanceSend(request *RebootDBInstanceRequest) (*RebootDBInstanceResponse, error) {
	statusCode, msg, err := c.RebootDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RebootDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RebootDBInstanceWithContext(ctx context.Context, request *RebootDBInstanceRequest) string {
	if request == nil {
		request = NewRebootDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRebootDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RebootDBInstanceWithContextV2(ctx context.Context, request *RebootDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRebootDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRebootDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBBackupPolicyRequest() (request *DescribeDBBackupPolicyRequest) {
	request = &DescribeDBBackupPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBBackupPolicy")
	return
}

func NewDescribeDBBackupPolicyResponse() (response *DescribeDBBackupPolicyResponse) {
	response = &DescribeDBBackupPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBBackupPolicy(request *DescribeDBBackupPolicyRequest) string {
	return c.DescribeDBBackupPolicyWithContext(context.Background(), request)
}

func (c *Client) DescribeDBBackupPolicySend(request *DescribeDBBackupPolicyRequest) (*DescribeDBBackupPolicyResponse, error) {
	statusCode, msg, err := c.DescribeDBBackupPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBBackupPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBBackupPolicyWithContext(ctx context.Context, request *DescribeDBBackupPolicyRequest) string {
	if request == nil {
		request = NewDescribeDBBackupPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBBackupPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBBackupPolicyWithContextV2(ctx context.Context, request *DescribeDBBackupPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBBackupPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBBackupPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDBInstanceOrderRequest() (request *UpdateDBInstanceOrderRequest) {
	request = &UpdateDBInstanceOrderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "UpdateDBInstanceOrder")
	return
}

func NewUpdateDBInstanceOrderResponse() (response *UpdateDBInstanceOrderResponse) {
	response = &UpdateDBInstanceOrderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDBInstanceOrder(request *UpdateDBInstanceOrderRequest) string {
	return c.UpdateDBInstanceOrderWithContext(context.Background(), request)
}

func (c *Client) UpdateDBInstanceOrderSend(request *UpdateDBInstanceOrderRequest) (*UpdateDBInstanceOrderResponse, error) {
	statusCode, msg, err := c.UpdateDBInstanceOrderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDBInstanceOrderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDBInstanceOrderWithContext(ctx context.Context, request *UpdateDBInstanceOrderRequest) string {
	if request == nil {
		request = NewUpdateDBInstanceOrderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDBInstanceOrderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDBInstanceOrderWithContextV2(ctx context.Context, request *UpdateDBInstanceOrderRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDBInstanceOrderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDBInstanceOrderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateResourceProtectionRequest() (request *UpdateResourceProtectionRequest) {
	request = &UpdateResourceProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sqlserver", APIVersion, "UpdateResourceProtection")
	return
}

func NewUpdateResourceProtectionResponse() (response *UpdateResourceProtectionResponse) {
	response = &UpdateResourceProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateResourceProtection(request *UpdateResourceProtectionRequest) string {
	return c.UpdateResourceProtectionWithContext(context.Background(), request)
}

func (c *Client) UpdateResourceProtectionSend(request *UpdateResourceProtectionRequest) (*UpdateResourceProtectionResponse, error) {
	statusCode, msg, err := c.UpdateResourceProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateResourceProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateResourceProtectionWithContext(ctx context.Context, request *UpdateResourceProtectionRequest) string {
	if request == nil {
		request = NewUpdateResourceProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateResourceProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateResourceProtectionWithContextV2(ctx context.Context, request *UpdateResourceProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateResourceProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateResourceProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
