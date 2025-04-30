package v20210520
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-05-20"

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
	request.Init().WithApiInfo("tidb", APIVersion, "CreateInstance")
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
	request.SetContentType("application/json")

	response := NewCreateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListInstanceRequest() (request *ListInstanceRequest) {
	request = &ListInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListInstance")
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
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeInstance")
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
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteInstance")
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
	request.Init().WithApiInfo("tidb", APIVersion, "RenameInstance")
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
	request.SetContentType("application/json")

	response := NewRenameInstanceResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "ListRegion")
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
	request.Init().WithApiInfo("tidb", APIVersion, "DescRegion")
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
func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityGroup")
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
func NewListSecurityGroupRequest() (request *ListSecurityGroupRequest) {
	request = &ListSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListSecurityGroup")
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
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeSecurityGroup")
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
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityGroup")
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
func NewUpdateSecurityGroupRequest() (request *UpdateSecurityGroupRequest) {
	request = &UpdateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityGroup")
	return
}

func NewUpdateSecurityGroupResponse() (response *UpdateSecurityGroupResponse) {
	response = &UpdateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateSecurityGroup(request *UpdateSecurityGroupRequest) string {
	return c.UpdateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) UpdateSecurityGroupWithContext(ctx context.Context, request *UpdateSecurityGroupRequest) string {
	if request == nil {
		request = NewUpdateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSecurityGroupResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "CloneSecurityGroup")
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
func NewBindSecurityGroupRequest() (request *BindSecurityGroupRequest) {
	request = &BindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "BindSecurityGroup")
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
	request.SetContentType("application/json")

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
	request.Init().WithApiInfo("tidb", APIVersion, "UnbindSecurityGroup")
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
	request.SetContentType("application/json")

	response := NewUnbindSecurityGroupResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "RebindSecurityGroup")
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
	request.SetContentType("application/json")

	response := NewRebindSecurityGroupResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityRule")
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
	request.SetContentType("application/json")

	response := NewCreateSecurityRuleResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityRule")
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
	request.SetContentType("application/json")

	response := NewUpdateSecurityRuleResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityRule")
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
	request.Init().WithApiInfo("tidb", APIVersion, "ListSecuredInstance")
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
	request.Init().WithApiInfo("tidb", APIVersion, "ListUnsecuredInstance")
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
func NewListBackupRequest() (request *ListBackupRequest) {
	request = &ListBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListBackup")
	return
}

func NewListBackupResponse() (response *ListBackupResponse) {
	response = &ListBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListBackup(request *ListBackupRequest) string {
	return c.ListBackupWithContext(context.Background(), request)
}

func (c *Client) ListBackupWithContext(ctx context.Context, request *ListBackupRequest) string {
	if request == nil {
		request = NewListBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateBackupRequest() (request *CreateBackupRequest) {
	request = &CreateBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateBackup")
	return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
	response = &CreateBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateBackup(request *CreateBackupRequest) string {
	return c.CreateBackupWithContext(context.Background(), request)
}

func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) string {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateBackupRuleRequest() (request *UpdateBackupRuleRequest) {
	request = &UpdateBackupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateBackupRule")
	return
}

func NewUpdateBackupRuleResponse() (response *UpdateBackupRuleResponse) {
	response = &UpdateBackupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateBackupRule(request *UpdateBackupRuleRequest) string {
	return c.UpdateBackupRuleWithContext(context.Background(), request)
}

func (c *Client) UpdateBackupRuleWithContext(ctx context.Context, request *UpdateBackupRuleRequest) string {
	if request == nil {
		request = NewUpdateBackupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateBackupRuleResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "ResetPassword")
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
	request.SetContentType("application/json")

	response := NewResetPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
	request = &DeleteBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteBackup")
	return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
	response = &DeleteBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBackup(request *DeleteBackupRequest) string {
	return c.DeleteBackupWithContext(context.Background(), request)
}

func (c *Client) DeleteBackupWithContext(ctx context.Context, request *DeleteBackupRequest) string {
	if request == nil {
		request = NewDeleteBackupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateRestoreRequest() (request *CreateRestoreRequest) {
	request = &CreateRestoreRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateRestore")
	return
}

func NewCreateRestoreResponse() (response *CreateRestoreResponse) {
	response = &CreateRestoreResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRestore(request *CreateRestoreRequest) string {
	return c.CreateRestoreWithContext(context.Background(), request)
}

func (c *Client) CreateRestoreWithContext(ctx context.Context, request *CreateRestoreRequest) string {
	if request == nil {
		request = NewCreateRestoreRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateRestoreResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewOpenTiMonitorRequest() (request *OpenTiMonitorRequest) {
	request = &OpenTiMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "OpenTiMonitor")
	return
}

func NewOpenTiMonitorResponse() (response *OpenTiMonitorResponse) {
	response = &OpenTiMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OpenTiMonitor(request *OpenTiMonitorRequest) string {
	return c.OpenTiMonitorWithContext(context.Background(), request)
}

func (c *Client) OpenTiMonitorWithContext(ctx context.Context, request *OpenTiMonitorRequest) string {
	if request == nil {
		request = NewOpenTiMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOpenTiMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateTask")
	return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTask(request *CreateTaskRequest) string {
	return c.CreateTaskWithContext(context.Background(), request)
}

func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) string {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewOperationTasksRequest() (request *OperationTasksRequest) {
	request = &OperationTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "OperationTasks")
	return
}

func NewOperationTasksResponse() (response *OperationTasksResponse) {
	response = &OperationTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperationTasks(request *OperationTasksRequest) string {
	return c.OperationTasksWithContext(context.Background(), request)
}

func (c *Client) OperationTasksWithContext(ctx context.Context, request *OperationTasksRequest) string {
	if request == nil {
		request = NewOperationTasksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperationTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCheckTaskNameRequest() (request *CheckTaskNameRequest) {
	request = &CheckTaskNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CheckTaskName")
	return
}

func NewCheckTaskNameResponse() (response *CheckTaskNameResponse) {
	response = &CheckTaskNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckTaskName(request *CheckTaskNameRequest) string {
	return c.CheckTaskNameWithContext(context.Background(), request)
}

func (c *Client) CheckTaskNameWithContext(ctx context.Context, request *CheckTaskNameRequest) string {
	if request == nil {
		request = NewCheckTaskNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckTaskNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTask(request *DescribeTaskRequest) string {
	return c.DescribeTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) string {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListTasksRequest() (request *ListTasksRequest) {
	request = &ListTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListTasks")
	return
}

func NewListTasksResponse() (response *ListTasksResponse) {
	response = &ListTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTasks(request *ListTasksRequest) string {
	return c.ListTasksWithContext(context.Background(), request)
}

func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) string {
	if request == nil {
		request = NewListTasksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTasksResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeDatabases")
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
func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeAccounts")
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
	request.Init().WithApiInfo("tidb", APIVersion, "CreateAccount")
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
	request.SetContentType("application/json")

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
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteAccount")
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
	request.SetContentType("application/json")

	response := NewDeleteAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyAccountInfoRequest() (request *ModifyAccountInfoRequest) {
	request = &ModifyAccountInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountInfo")
	return
}

func NewModifyAccountInfoResponse() (response *ModifyAccountInfoResponse) {
	response = &ModifyAccountInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAccountInfo(request *ModifyAccountInfoRequest) string {
	return c.ModifyAccountInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyAccountInfoWithContext(ctx context.Context, request *ModifyAccountInfoRequest) string {
	if request == nil {
		request = NewModifyAccountInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAccountInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
	request = &ModifyAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountPrivileges")
	return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
	response = &ModifyAccountPrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) string {
	return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) string {
	if request == nil {
		request = NewModifyAccountPrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAccountPrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewConfigurationInstanceEipRequest() (request *ConfigurationInstanceEipRequest) {
	request = &ConfigurationInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ConfigurationInstanceEip")
	return
}

func NewConfigurationInstanceEipResponse() (response *ConfigurationInstanceEipResponse) {
	response = &ConfigurationInstanceEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ConfigurationInstanceEip(request *ConfigurationInstanceEipRequest) string {
	return c.ConfigurationInstanceEipWithContext(context.Background(), request)
}

func (c *Client) ConfigurationInstanceEipWithContext(ctx context.Context, request *ConfigurationInstanceEipRequest) string {
	if request == nil {
		request = NewConfigurationInstanceEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewConfigurationInstanceEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


