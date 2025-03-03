package v20181225

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2018-12-25"

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
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateDBInstance")
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
func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBInstances")
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
func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
	request = &DeleteDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DeleteDBInstance")
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
func NewStatisticDBInstancesRequest() (request *StatisticDBInstancesRequest) {
	request = &StatisticDBInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "StatisticDBInstances")
	return
}

func NewStatisticDBInstancesResponse() (response *StatisticDBInstancesResponse) {
	response = &StatisticDBInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StatisticDBInstances(request *StatisticDBInstancesRequest) string {
	return c.StatisticDBInstancesWithContext(context.Background(), request)
}

func (c *Client) StatisticDBInstancesWithContext(ctx context.Context, request *StatisticDBInstancesRequest) string {
	if request == nil {
		request = NewStatisticDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStatisticDBInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDBInstanceRequest() (request *ModifyDBInstanceRequest) {
	request = &ModifyDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyDBInstance")
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

func (c *Client) ModifyDBInstanceWithContext(ctx context.Context, request *ModifyDBInstanceRequest) string {
	if request == nil {
		request = NewModifyDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBInstanceResponse()
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
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateSecurityGroup")
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
	request.SetContentType("application/json")

	response := NewCreateSecurityGroupResponse()
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
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeSecurityGroup")
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
func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DeleteSecurityGroup")
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
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
	request = &ModifySecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifySecurityGroup")
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
	request.SetContentType("application/json")

	response := NewModifySecurityGroupResponse()
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
	request.Init().WithApiInfo("postgresql", APIVersion, "CloneSecurityGroup")
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
	request.SetContentType("application/json")

	response := NewCloneSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifySecurityGroupRuleRequest() (request *ModifySecurityGroupRuleRequest) {
	request = &ModifySecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifySecurityGroupRule")
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
func NewSecurityGroupRelationRequest() (request *SecurityGroupRelationRequest) {
	request = &SecurityGroupRelationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "SecurityGroupRelation")
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
func NewModifySecurityGroupRuleNameRequest() (request *ModifySecurityGroupRuleNameRequest) {
	request = &ModifySecurityGroupRuleNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifySecurityGroupRuleName")
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
func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
	request = &DescribeDBLogFilesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBLogFiles")
	return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
	response = &DescribeDBLogFilesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) string {
	return c.DescribeDBLogFilesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBLogFilesWithContext(ctx context.Context, request *DescribeDBLogFilesRequest) string {
	if request == nil {
		request = NewDescribeDBLogFilesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBLogFilesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDBBackupRequest() (request *CreateDBBackupRequest) {
	request = &CreateDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateDBBackup")
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
func NewDeleteDBBackupRequest() (request *DeleteDBBackupRequest) {
	request = &DeleteDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DeleteDBBackup")
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
func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
	request = &DescribeDBBackupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBBackups")
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
func NewModifyDBBackupPolicyRequest() (request *ModifyDBBackupPolicyRequest) {
	request = &ModifyDBBackupPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyDBBackupPolicy")
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
func NewOverrideDBInstanceRequest() (request *OverrideDBInstanceRequest) {
	request = &OverrideDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "OverrideDBInstance")
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
func NewCreateDBParameterGroupRequest() (request *CreateDBParameterGroupRequest) {
	request = &CreateDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateDBParameterGroup")
	return
}

func NewCreateDBParameterGroupResponse() (response *CreateDBParameterGroupResponse) {
	response = &CreateDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBParameterGroup(request *CreateDBParameterGroupRequest) string {
	return c.CreateDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) CreateDBParameterGroupWithContext(ctx context.Context, request *CreateDBParameterGroupRequest) string {
	if request == nil {
		request = NewCreateDBParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDBParameterGroupResponse()
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
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyDBParameterGroup")
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
	request.SetContentType("application/json")

	response := NewModifyDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDBParameterGroupRequest() (request *DeleteDBParameterGroupRequest) {
	request = &DeleteDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DeleteDBParameterGroup")
	return
}

func NewDeleteDBParameterGroupResponse() (response *DeleteDBParameterGroupResponse) {
	response = &DeleteDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDBParameterGroup(request *DeleteDBParameterGroupRequest) string {
	return c.DeleteDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteDBParameterGroupWithContext(ctx context.Context, request *DeleteDBParameterGroupRequest) string {
	if request == nil {
		request = NewDeleteDBParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewResetDBParameterGroupRequest() (request *ResetDBParameterGroupRequest) {
	request = &ResetDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ResetDBParameterGroup")
	return
}

func NewResetDBParameterGroupResponse() (response *ResetDBParameterGroupResponse) {
	response = &ResetDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetDBParameterGroup(request *ResetDBParameterGroupRequest) string {
	return c.ResetDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ResetDBParameterGroupWithContext(ctx context.Context, request *ResetDBParameterGroupRequest) string {
	if request == nil {
		request = NewResetDBParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResetDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDBParameterGroupRequest() (request *DescribeDBParameterGroupRequest) {
	request = &DescribeDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBParameterGroup")
	return
}

func NewDescribeDBParameterGroupResponse() (response *DescribeDBParameterGroupResponse) {
	response = &DescribeDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBParameterGroup(request *DescribeDBParameterGroupRequest) string {
	return c.DescribeDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeDBParameterGroupWithContext(ctx context.Context, request *DescribeDBParameterGroupRequest) string {
	if request == nil {
		request = NewDescribeDBParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBParameterGroupResponse()
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
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeEngineDefaultParameters")
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
func NewDescribeDBInstanceParametersRequest() (request *DescribeDBInstanceParametersRequest) {
	request = &DescribeDBInstanceParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBInstanceParameters")
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
func NewRebootDBInstanceRequest() (request *RebootDBInstanceRequest) {
	request = &RebootDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "RebootDBInstance")
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
func NewDescribeDBEngineVersionsRequest() (request *DescribeDBEngineVersionsRequest) {
	request = &DescribeDBEngineVersionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBEngineVersions")
	return
}

func NewDescribeDBEngineVersionsResponse() (response *DescribeDBEngineVersionsResponse) {
	response = &DescribeDBEngineVersionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBEngineVersions(request *DescribeDBEngineVersionsRequest) string {
	return c.DescribeDBEngineVersionsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBEngineVersionsWithContext(ctx context.Context, request *DescribeDBEngineVersionsRequest) string {
	if request == nil {
		request = NewDescribeDBEngineVersionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBEngineVersionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAllocateDBInstanceEipRequest() (request *AllocateDBInstanceEipRequest) {
	request = &AllocateDBInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "AllocateDBInstanceEip")
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
func NewReleaseDBInstanceEipRequest() (request *ReleaseDBInstanceEipRequest) {
	request = &ReleaseDBInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ReleaseDBInstanceEip")
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
func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
	request = &ModifyDBInstanceSpecRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyDBInstanceSpec")
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
func NewRestoreDBInstanceFromDBBackupRequest() (request *RestoreDBInstanceFromDBBackupRequest) {
	request = &RestoreDBInstanceFromDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "RestoreDBInstanceFromDBBackup")
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
func NewSwitchDBInstanceHARequest() (request *SwitchDBInstanceHARequest) {
	request = &SwitchDBInstanceHARequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "SwitchDBInstanceHA")
	return
}

func NewSwitchDBInstanceHAResponse() (response *SwitchDBInstanceHAResponse) {
	response = &SwitchDBInstanceHAResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) string {
	return c.SwitchDBInstanceHAWithContext(context.Background(), request)
}

func (c *Client) SwitchDBInstanceHAWithContext(ctx context.Context, request *SwitchDBInstanceHARequest) string {
	if request == nil {
		request = NewSwitchDBInstanceHARequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSwitchDBInstanceHAResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDBInstanceReadReplicaRequest() (request *CreateDBInstanceReadReplicaRequest) {
	request = &CreateDBInstanceReadReplicaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateDBInstanceReadReplica")
	return
}

func NewCreateDBInstanceReadReplicaResponse() (response *CreateDBInstanceReadReplicaResponse) {
	response = &CreateDBInstanceReadReplicaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBInstanceReadReplica(request *CreateDBInstanceReadReplicaRequest) string {
	return c.CreateDBInstanceReadReplicaWithContext(context.Background(), request)
}

func (c *Client) CreateDBInstanceReadReplicaWithContext(ctx context.Context, request *CreateDBInstanceReadReplicaRequest) string {
	if request == nil {
		request = NewCreateDBInstanceReadReplicaRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDBInstanceReadReplicaResponse()
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
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyInstanceAccountInfo")
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
func NewDescribeInstanceDatabasesRequest() (request *DescribeInstanceDatabasesRequest) {
	request = &DescribeInstanceDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeInstanceDatabases")
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
func NewDescribeDBInstanceExtensionsRequest() (request *DescribeDBInstanceExtensionsRequest) {
	request = &DescribeDBInstanceExtensionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeDBInstanceExtensions")
	return
}

func NewDescribeDBInstanceExtensionsResponse() (response *DescribeDBInstanceExtensionsResponse) {
	response = &DescribeDBInstanceExtensionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstanceExtensions(request *DescribeDBInstanceExtensionsRequest) string {
	return c.DescribeDBInstanceExtensionsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceExtensionsWithContext(ctx context.Context, request *DescribeDBInstanceExtensionsRequest) string {
	if request == nil {
		request = NewDescribeDBInstanceExtensionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDBInstanceExtensionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDBInstanceExtensionRequest() (request *ModifyDBInstanceExtensionRequest) {
	request = &ModifyDBInstanceExtensionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyDBInstanceExtension")
	return
}

func NewModifyDBInstanceExtensionResponse() (response *ModifyDBInstanceExtensionResponse) {
	response = &ModifyDBInstanceExtensionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstanceExtension(request *ModifyDBInstanceExtensionRequest) string {
	return c.ModifyDBInstanceExtensionWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceExtensionWithContext(ctx context.Context, request *ModifyDBInstanceExtensionRequest) string {
	if request == nil {
		request = NewModifyDBInstanceExtensionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBInstanceExtensionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCollationsRequest() (request *DescribeCollationsRequest) {
	request = &DescribeCollationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeCollations")
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
func NewModifyInstanceDatabaseOwnerRequest() (request *ModifyInstanceDatabaseOwnerRequest) {
	request = &ModifyInstanceDatabaseOwnerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyInstanceDatabaseOwner")
	return
}

func NewModifyInstanceDatabaseOwnerResponse() (response *ModifyInstanceDatabaseOwnerResponse) {
	response = &ModifyInstanceDatabaseOwnerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabaseOwner(request *ModifyInstanceDatabaseOwnerRequest) string {
	return c.ModifyInstanceDatabaseOwnerWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabaseOwnerWithContext(ctx context.Context, request *ModifyInstanceDatabaseOwnerRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabaseOwnerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabaseOwnerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInstanceDatabaseRequest() (request *DeleteInstanceDatabaseRequest) {
	request = &DeleteInstanceDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DeleteInstanceDatabase")
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
func NewCreateInstanceDatabaseRequest() (request *CreateInstanceDatabaseRequest) {
	request = &CreateInstanceDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateInstanceDatabase")
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
func NewDescribeInstanceAccountsRequest() (request *DescribeInstanceAccountsRequest) {
	request = &DescribeInstanceAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DescribeInstanceAccounts")
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
func NewCreateInstanceAccountRequest() (request *CreateInstanceAccountRequest) {
	request = &CreateInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "CreateInstanceAccount")
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
func NewDeleteInstanceAccountRequest() (request *DeleteInstanceAccountRequest) {
	request = &DeleteInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "DeleteInstanceAccount")
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
func NewModifyDBNetworkRequest() (request *ModifyDBNetworkRequest) {
	request = &ModifyDBNetworkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("postgresql", APIVersion, "ModifyDBNetwork")
	return
}

func NewModifyDBNetworkResponse() (response *ModifyDBNetworkResponse) {
	response = &ModifyDBNetworkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBNetwork(request *ModifyDBNetworkRequest) string {
	return c.ModifyDBNetworkWithContext(context.Background(), request)
}

func (c *Client) ModifyDBNetworkWithContext(ctx context.Context, request *ModifyDBNetworkRequest) string {
	if request == nil {
		request = NewModifyDBNetworkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBNetworkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
