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
