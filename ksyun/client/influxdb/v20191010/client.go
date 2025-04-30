package v20191010
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2019-10-10"

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
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateInstance")
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
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteInstance")
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
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeInstance")
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
func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeInstances")
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
func NewDescribeInstanceNodeRequest() (request *DescribeInstanceNodeRequest) {
	request = &DescribeInstanceNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeInstanceNode")
	return
}

func NewDescribeInstanceNodeResponse() (response *DescribeInstanceNodeResponse) {
	response = &DescribeInstanceNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceNode(request *DescribeInstanceNodeRequest) string {
	return c.DescribeInstanceNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceNodeWithContext(ctx context.Context, request *DescribeInstanceNodeRequest) string {
	if request == nil {
		request = NewDescribeInstanceNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceNodeResponse()
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
	request.Init().WithApiInfo("Influxdb", APIVersion, "RenameInstance")
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
func NewDescribeValidRegionsRequest() (request *DescribeValidRegionsRequest) {
	request = &DescribeValidRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeValidRegions")
	return
}

func NewDescribeValidRegionsResponse() (response *DescribeValidRegionsResponse) {
	response = &DescribeValidRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeValidRegions(request *DescribeValidRegionsRequest) string {
	return c.DescribeValidRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeValidRegionsWithContext(ctx context.Context, request *DescribeValidRegionsRequest) string {
	if request == nil {
		request = NewDescribeValidRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeValidRegionsResponse()
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
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeSecurityGroup")
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
func NewCreateSecurityRuleRequest() (request *CreateSecurityRuleRequest) {
	request = &CreateSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateSecurityRule")
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
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteSecurityRule")
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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
	request = &DescribeDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeDatabases")
	return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
	response = &DescribeDatabasesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) string {
	return c.DescribeDatabasesWithContext(context.Background(), request)
}

func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) string {
	if request == nil {
		request = NewDescribeDatabasesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDatabasesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDatabaseRequest() (request *CreateDatabaseRequest) {
	request = &CreateDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateDatabase")
	return
}

func NewCreateDatabaseResponse() (response *CreateDatabaseResponse) {
	response = &CreateDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDatabase(request *CreateDatabaseRequest) string {
	return c.CreateDatabaseWithContext(context.Background(), request)
}

func (c *Client) CreateDatabaseWithContext(ctx context.Context, request *CreateDatabaseRequest) string {
	if request == nil {
		request = NewCreateDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDatabaseRequest() (request *DeleteDatabaseRequest) {
	request = &DeleteDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteDatabase")
	return
}

func NewDeleteDatabaseResponse() (response *DeleteDatabaseResponse) {
	response = &DeleteDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDatabase(request *DeleteDatabaseRequest) string {
	return c.DeleteDatabaseWithContext(context.Background(), request)
}

func (c *Client) DeleteDatabaseWithContext(ctx context.Context, request *DeleteDatabaseRequest) string {
	if request == nil {
		request = NewDeleteDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRetentionPolicyListRequest() (request *DescribeRetentionPolicyListRequest) {
	request = &DescribeRetentionPolicyListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeRetentionPolicyList")
	return
}

func NewDescribeRetentionPolicyListResponse() (response *DescribeRetentionPolicyListResponse) {
	response = &DescribeRetentionPolicyListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRetentionPolicyList(request *DescribeRetentionPolicyListRequest) string {
	return c.DescribeRetentionPolicyListWithContext(context.Background(), request)
}

func (c *Client) DescribeRetentionPolicyListWithContext(ctx context.Context, request *DescribeRetentionPolicyListRequest) string {
	if request == nil {
		request = NewDescribeRetentionPolicyListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRetentionPolicyListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateRetentionPolicyRequest() (request *CreateRetentionPolicyRequest) {
	request = &CreateRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateRetentionPolicy")
	return
}

func NewCreateRetentionPolicyResponse() (response *CreateRetentionPolicyResponse) {
	response = &CreateRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRetentionPolicy(request *CreateRetentionPolicyRequest) string {
	return c.CreateRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) CreateRetentionPolicyWithContext(ctx context.Context, request *CreateRetentionPolicyRequest) string {
	if request == nil {
		request = NewCreateRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteRetentionPolicyRequest() (request *DeleteRetentionPolicyRequest) {
	request = &DeleteRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteRetentionPolicy")
	return
}

func NewDeleteRetentionPolicyResponse() (response *DeleteRetentionPolicyResponse) {
	response = &DeleteRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRetentionPolicy(request *DeleteRetentionPolicyRequest) string {
	return c.DeleteRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteRetentionPolicyWithContext(ctx context.Context, request *DeleteRetentionPolicyRequest) string {
	if request == nil {
		request = NewDeleteRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyRetentionPolicyRequest() (request *ModifyRetentionPolicyRequest) {
	request = &ModifyRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "ModifyRetentionPolicy")
	return
}

func NewModifyRetentionPolicyResponse() (response *ModifyRetentionPolicyResponse) {
	response = &ModifyRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRetentionPolicy(request *ModifyRetentionPolicyRequest) string {
	return c.ModifyRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyRetentionPolicyWithContext(ctx context.Context, request *ModifyRetentionPolicyRequest) string {
	if request == nil {
		request = NewModifyRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMeasurementsRequest() (request *DescribeMeasurementsRequest) {
	request = &DescribeMeasurementsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeMeasurements")
	return
}

func NewDescribeMeasurementsResponse() (response *DescribeMeasurementsResponse) {
	response = &DescribeMeasurementsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMeasurements(request *DescribeMeasurementsRequest) string {
	return c.DescribeMeasurementsWithContext(context.Background(), request)
}

func (c *Client) DescribeMeasurementsWithContext(ctx context.Context, request *DescribeMeasurementsRequest) string {
	if request == nil {
		request = NewDescribeMeasurementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMeasurementsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteMeasurementRequest() (request *DeleteMeasurementRequest) {
	request = &DeleteMeasurementRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteMeasurement")
	return
}

func NewDeleteMeasurementResponse() (response *DeleteMeasurementResponse) {
	response = &DeleteMeasurementResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMeasurement(request *DeleteMeasurementRequest) string {
	return c.DeleteMeasurementWithContext(context.Background(), request)
}

func (c *Client) DeleteMeasurementWithContext(ctx context.Context, request *DeleteMeasurementRequest) string {
	if request == nil {
		request = NewDeleteMeasurementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMeasurementResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeAccounts")
	return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) string {
	return c.DescribeAccountsWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) string {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateAccount")
	return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAccount(request *CreateAccountRequest) string {
	return c.CreateAccountWithContext(context.Background(), request)
}

func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) string {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
	request = &DeleteAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteAccount")
	return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
	response = &DeleteAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAccount(request *DeleteAccountRequest) string {
	return c.DeleteAccountWithContext(context.Background(), request)
}

func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) string {
	if request == nil {
		request = NewDeleteAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
	request = &DescribeAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeAccountPrivileges")
	return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
	response = &DescribeAccountPrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) string {
	return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) string {
	if request == nil {
		request = NewDescribeAccountPrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountPrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGrantAccountPrivilegeRequest() (request *GrantAccountPrivilegeRequest) {
	request = &GrantAccountPrivilegeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "GrantAccountPrivilege")
	return
}

func NewGrantAccountPrivilegeResponse() (response *GrantAccountPrivilegeResponse) {
	response = &GrantAccountPrivilegeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) string {
	return c.GrantAccountPrivilegeWithContext(context.Background(), request)
}

func (c *Client) GrantAccountPrivilegeWithContext(ctx context.Context, request *GrantAccountPrivilegeRequest) string {
	if request == nil {
		request = NewGrantAccountPrivilegeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGrantAccountPrivilegeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRevokeAccountPrivilegeRequest() (request *RevokeAccountPrivilegeRequest) {
	request = &RevokeAccountPrivilegeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "RevokeAccountPrivilege")
	return
}

func NewRevokeAccountPrivilegeResponse() (response *RevokeAccountPrivilegeResponse) {
	response = &RevokeAccountPrivilegeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RevokeAccountPrivilege(request *RevokeAccountPrivilegeRequest) string {
	return c.RevokeAccountPrivilegeWithContext(context.Background(), request)
}

func (c *Client) RevokeAccountPrivilegeWithContext(ctx context.Context, request *RevokeAccountPrivilegeRequest) string {
	if request == nil {
		request = NewRevokeAccountPrivilegeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRevokeAccountPrivilegeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
	request = &ResetAccountPasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "ResetAccountPassword")
	return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
	response = &ResetAccountPasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) string {
	return c.ResetAccountPasswordWithContext(context.Background(), request)
}

func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) string {
	if request == nil {
		request = NewResetAccountPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetAccountPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAccountDetailListRequest() (request *DescribeAccountDetailListRequest) {
	request = &DescribeAccountDetailListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeAccountDetailList")
	return
}

func NewDescribeAccountDetailListResponse() (response *DescribeAccountDetailListResponse) {
	response = &DescribeAccountDetailListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccountDetailList(request *DescribeAccountDetailListRequest) string {
	return c.DescribeAccountDetailListWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountDetailListWithContext(ctx context.Context, request *DescribeAccountDetailListRequest) string {
	if request == nil {
		request = NewDescribeAccountDetailListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountDetailListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


