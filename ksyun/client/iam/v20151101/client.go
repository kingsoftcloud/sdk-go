package v20151101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2015-11-01"

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

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreateUser")
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateUser(request *CreateUserRequest) (string) {
    return c.CreateUserWithContext(context.Background(), request)
}

func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (string) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    request.SetContext(ctx)

    response := NewCreateUserResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListUsers")
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListUsers(request *ListUsersRequest) (string) {
    return c.ListUsersWithContext(context.Background(), request)
}

func (c *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest) (string) {
    if request == nil {
        request = NewListUsersRequest()
    }
    request.SetContext(ctx)

    response := NewListUsersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "UpdateUser")
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateUser(request *UpdateUserRequest) (string) {
    return c.UpdateUserWithContext(context.Background(), request)
}

func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (string) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    request.SetContext(ctx)

    response := NewUpdateUserResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetUserRequest() (request *GetUserRequest) {
    request = &GetUserRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "GetUser")
    return
}

func NewGetUserResponse() (response *GetUserResponse) {
    response = &GetUserResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetUser(request *GetUserRequest) (string) {
    return c.GetUserWithContext(context.Background(), request)
}

func (c *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest) (string) {
    if request == nil {
        request = NewGetUserRequest()
    }
    request.SetContext(ctx)

    response := NewGetUserResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeleteUser")
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteUser(request *DeleteUserRequest) (string) {
    return c.DeleteUserWithContext(context.Background(), request)
}

func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (string) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteUserResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListAttachedUserPoliciesRequest() (request *ListAttachedUserPoliciesRequest) {
    request = &ListAttachedUserPoliciesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListAttachedUserPolicies")
    return
}

func NewListAttachedUserPoliciesResponse() (response *ListAttachedUserPoliciesResponse) {
    response = &ListAttachedUserPoliciesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAttachedUserPolicies(request *ListAttachedUserPoliciesRequest) (string) {
    return c.ListAttachedUserPoliciesWithContext(context.Background(), request)
}

func (c *Client) ListAttachedUserPoliciesWithContext(ctx context.Context, request *ListAttachedUserPoliciesRequest) (string) {
    if request == nil {
        request = NewListAttachedUserPoliciesRequest()
    }
    request.SetContext(ctx)

    response := NewListAttachedUserPoliciesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListPolicyVersionsRequest() (request *ListPolicyVersionsRequest) {
    request = &ListPolicyVersionsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListPolicyVersions")
    return
}

func NewListPolicyVersionsResponse() (response *ListPolicyVersionsResponse) {
    response = &ListPolicyVersionsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (string) {
    return c.ListPolicyVersionsWithContext(context.Background(), request)
}

func (c *Client) ListPolicyVersionsWithContext(ctx context.Context, request *ListPolicyVersionsRequest) (string) {
    if request == nil {
        request = NewListPolicyVersionsRequest()
    }
    request.SetContext(ctx)

    response := NewListPolicyVersionsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetDefaultPolicyVersionRequest() (request *SetDefaultPolicyVersionRequest) {
    request = &SetDefaultPolicyVersionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "SetDefaultPolicyVersion")
    return
}

func NewSetDefaultPolicyVersionResponse() (response *SetDefaultPolicyVersionResponse) {
    response = &SetDefaultPolicyVersionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetDefaultPolicyVersion(request *SetDefaultPolicyVersionRequest) (string) {
    return c.SetDefaultPolicyVersionWithContext(context.Background(), request)
}

func (c *Client) SetDefaultPolicyVersionWithContext(ctx context.Context, request *SetDefaultPolicyVersionRequest) (string) {
    if request == nil {
        request = NewSetDefaultPolicyVersionRequest()
    }
    request.SetContext(ctx)

    response := NewSetDefaultPolicyVersionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAttachUserPolicyRequest() (request *AttachUserPolicyRequest) {
    request = &AttachUserPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "AttachUserPolicy")
    return
}

func NewAttachUserPolicyResponse() (response *AttachUserPolicyResponse) {
    response = &AttachUserPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AttachUserPolicy(request *AttachUserPolicyRequest) (string) {
    return c.AttachUserPolicyWithContext(context.Background(), request)
}

func (c *Client) AttachUserPolicyWithContext(ctx context.Context, request *AttachUserPolicyRequest) (string) {
    if request == nil {
        request = NewAttachUserPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewAttachUserPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeletePolicyVersionRequest() (request *DeletePolicyVersionRequest) {
    request = &DeletePolicyVersionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeletePolicyVersion")
    return
}

func NewDeletePolicyVersionResponse() (response *DeletePolicyVersionResponse) {
    response = &DeletePolicyVersionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (string) {
    return c.DeletePolicyVersionWithContext(context.Background(), request)
}

func (c *Client) DeletePolicyVersionWithContext(ctx context.Context, request *DeletePolicyVersionRequest) (string) {
    if request == nil {
        request = NewDeletePolicyVersionRequest()
    }
    request.SetContext(ctx)

    response := NewDeletePolicyVersionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetPolicyVersionRequest() (request *GetPolicyVersionRequest) {
    request = &GetPolicyVersionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "GetPolicyVersion")
    return
}

func NewGetPolicyVersionResponse() (response *GetPolicyVersionResponse) {
    response = &GetPolicyVersionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPolicyVersion(request *GetPolicyVersionRequest) (string) {
    return c.GetPolicyVersionWithContext(context.Background(), request)
}

func (c *Client) GetPolicyVersionWithContext(ctx context.Context, request *GetPolicyVersionRequest) (string) {
    if request == nil {
        request = NewGetPolicyVersionRequest()
    }
    request.SetContext(ctx)

    response := NewGetPolicyVersionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreatePolicyVersionRequest() (request *CreatePolicyVersionRequest) {
    request = &CreatePolicyVersionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreatePolicyVersion")
    return
}

func NewCreatePolicyVersionResponse() (response *CreatePolicyVersionResponse) {
    response = &CreatePolicyVersionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (string) {
    return c.CreatePolicyVersionWithContext(context.Background(), request)
}

func (c *Client) CreatePolicyVersionWithContext(ctx context.Context, request *CreatePolicyVersionRequest) (string) {
    if request == nil {
        request = NewCreatePolicyVersionRequest()
    }
    request.SetContext(ctx)

    response := NewCreatePolicyVersionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListPoliciesRequest() (request *ListPoliciesRequest) {
    request = &ListPoliciesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListPolicies")
    return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
    response = &ListPoliciesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListPolicies(request *ListPoliciesRequest) (string) {
    return c.ListPoliciesWithContext(context.Background(), request)
}

func (c *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest) (string) {
    if request == nil {
        request = NewListPoliciesRequest()
    }
    request.SetContext(ctx)

    response := NewListPoliciesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetPolicyRequest() (request *GetPolicyRequest) {
    request = &GetPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "GetPolicy")
    return
}

func NewGetPolicyResponse() (response *GetPolicyResponse) {
    response = &GetPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPolicy(request *GetPolicyRequest) (string) {
    return c.GetPolicyWithContext(context.Background(), request)
}

func (c *Client) GetPolicyWithContext(ctx context.Context, request *GetPolicyRequest) (string) {
    if request == nil {
        request = NewGetPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewGetPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
    request = &DeletePolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeletePolicy")
    return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
    response = &DeletePolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeletePolicy(request *DeletePolicyRequest) (string) {
    return c.DeletePolicyWithContext(context.Background(), request)
}

func (c *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest) (string) {
    if request == nil {
        request = NewDeletePolicyRequest()
    }
    request.SetContext(ctx)

    response := NewDeletePolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
    request = &CreatePolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreatePolicy")
    return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
    response = &CreatePolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreatePolicy(request *CreatePolicyRequest) (string) {
    return c.CreatePolicyWithContext(context.Background(), request)
}

func (c *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest) (string) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    request.SetContext(ctx)

    response := NewCreatePolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewChangePasswordRequest() (request *ChangePasswordRequest) {
    request = &ChangePasswordRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ChangePassword")
    return
}

func NewChangePasswordResponse() (response *ChangePasswordResponse) {
    response = &ChangePasswordResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ChangePassword(request *ChangePasswordRequest) (string) {
    return c.ChangePasswordWithContext(context.Background(), request)
}

func (c *Client) ChangePasswordWithContext(ctx context.Context, request *ChangePasswordRequest) (string) {
    if request == nil {
        request = NewChangePasswordRequest()
    }
    request.SetContext(ctx)

    response := NewChangePasswordResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateLoginProfileRequest() (request *UpdateLoginProfileRequest) {
    request = &UpdateLoginProfileRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "UpdateLoginProfile")
    return
}

func NewUpdateLoginProfileResponse() (response *UpdateLoginProfileResponse) {
    response = &UpdateLoginProfileResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (string) {
    return c.UpdateLoginProfileWithContext(context.Background(), request)
}

func (c *Client) UpdateLoginProfileWithContext(ctx context.Context, request *UpdateLoginProfileRequest) (string) {
    if request == nil {
        request = NewUpdateLoginProfileRequest()
    }
    request.SetContext(ctx)

    response := NewUpdateLoginProfileResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetLoginProfileRequest() (request *GetLoginProfileRequest) {
    request = &GetLoginProfileRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "GetLoginProfile")
    return
}

func NewGetLoginProfileResponse() (response *GetLoginProfileResponse) {
    response = &GetLoginProfileResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetLoginProfile(request *GetLoginProfileRequest) (string) {
    return c.GetLoginProfileWithContext(context.Background(), request)
}

func (c *Client) GetLoginProfileWithContext(ctx context.Context, request *GetLoginProfileRequest) (string) {
    if request == nil {
        request = NewGetLoginProfileRequest()
    }
    request.SetContext(ctx)

    response := NewGetLoginProfileResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateAccessKeyRequest() (request *CreateAccessKeyRequest) {
    request = &CreateAccessKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreateAccessKey")
    return
}

func NewCreateAccessKeyResponse() (response *CreateAccessKeyResponse) {
    response = &CreateAccessKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateAccessKey(request *CreateAccessKeyRequest) (string) {
    return c.CreateAccessKeyWithContext(context.Background(), request)
}

func (c *Client) CreateAccessKeyWithContext(ctx context.Context, request *CreateAccessKeyRequest) (string) {
    if request == nil {
        request = NewCreateAccessKeyRequest()
    }
    request.SetContext(ctx)

    response := NewCreateAccessKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListAccessKeysRequest() (request *ListAccessKeysRequest) {
    request = &ListAccessKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListAccessKeys")
    return
}

func NewListAccessKeysResponse() (response *ListAccessKeysResponse) {
    response = &ListAccessKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAccessKeys(request *ListAccessKeysRequest) (string) {
    return c.ListAccessKeysWithContext(context.Background(), request)
}

func (c *Client) ListAccessKeysWithContext(ctx context.Context, request *ListAccessKeysRequest) (string) {
    if request == nil {
        request = NewListAccessKeysRequest()
    }
    request.SetContext(ctx)

    response := NewListAccessKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateAccessKeyRequest() (request *UpdateAccessKeyRequest) {
    request = &UpdateAccessKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "UpdateAccessKey")
    return
}

func NewUpdateAccessKeyResponse() (response *UpdateAccessKeyResponse) {
    response = &UpdateAccessKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateAccessKey(request *UpdateAccessKeyRequest) (string) {
    return c.UpdateAccessKeyWithContext(context.Background(), request)
}

func (c *Client) UpdateAccessKeyWithContext(ctx context.Context, request *UpdateAccessKeyRequest) (string) {
    if request == nil {
        request = NewUpdateAccessKeyRequest()
    }
    request.SetContext(ctx)

    response := NewUpdateAccessKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteAccessKeyRequest() (request *DeleteAccessKeyRequest) {
    request = &DeleteAccessKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeleteAccessKey")
    return
}

func NewDeleteAccessKeyResponse() (response *DeleteAccessKeyResponse) {
    response = &DeleteAccessKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (string) {
    return c.DeleteAccessKeyWithContext(context.Background(), request)
}

func (c *Client) DeleteAccessKeyWithContext(ctx context.Context, request *DeleteAccessKeyRequest) (string) {
    if request == nil {
        request = NewDeleteAccessKeyRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteAccessKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreateRole")
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateRole(request *CreateRoleRequest) (string) {
    return c.CreateRoleWithContext(context.Background(), request)
}

func (c *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest) (string) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    request.SetContext(ctx)

    response := NewCreateRoleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeleteRole")
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteRole(request *DeleteRoleRequest) (string) {
    return c.DeleteRoleWithContext(context.Background(), request)
}

func (c *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest) (string) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteRoleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetRoleRequest() (request *GetRoleRequest) {
    request = &GetRoleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "GetRole")
    return
}

func NewGetRoleResponse() (response *GetRoleResponse) {
    response = &GetRoleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetRole(request *GetRoleRequest) (string) {
    return c.GetRoleWithContext(context.Background(), request)
}

func (c *Client) GetRoleWithContext(ctx context.Context, request *GetRoleRequest) (string) {
    if request == nil {
        request = NewGetRoleRequest()
    }
    request.SetContext(ctx)

    response := NewGetRoleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListRolesRequest() (request *ListRolesRequest) {
    request = &ListRolesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListRoles")
    return
}

func NewListRolesResponse() (response *ListRolesResponse) {
    response = &ListRolesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListRoles(request *ListRolesRequest) (string) {
    return c.ListRolesWithContext(context.Background(), request)
}

func (c *Client) ListRolesWithContext(ctx context.Context, request *ListRolesRequest) (string) {
    if request == nil {
        request = NewListRolesRequest()
    }
    request.SetContext(ctx)

    response := NewListRolesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAttachRolePolicyRequest() (request *AttachRolePolicyRequest) {
    request = &AttachRolePolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "AttachRolePolicy")
    return
}

func NewAttachRolePolicyResponse() (response *AttachRolePolicyResponse) {
    response = &AttachRolePolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AttachRolePolicy(request *AttachRolePolicyRequest) (string) {
    return c.AttachRolePolicyWithContext(context.Background(), request)
}

func (c *Client) AttachRolePolicyWithContext(ctx context.Context, request *AttachRolePolicyRequest) (string) {
    if request == nil {
        request = NewAttachRolePolicyRequest()
    }
    request.SetContext(ctx)

    response := NewAttachRolePolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDetachRolePolicyRequest() (request *DetachRolePolicyRequest) {
    request = &DetachRolePolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DetachRolePolicy")
    return
}

func NewDetachRolePolicyResponse() (response *DetachRolePolicyResponse) {
    response = &DetachRolePolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DetachRolePolicy(request *DetachRolePolicyRequest) (string) {
    return c.DetachRolePolicyWithContext(context.Background(), request)
}

func (c *Client) DetachRolePolicyWithContext(ctx context.Context, request *DetachRolePolicyRequest) (string) {
    if request == nil {
        request = NewDetachRolePolicyRequest()
    }
    request.SetContext(ctx)

    response := NewDetachRolePolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListAttachedRolePoliciesRequest() (request *ListAttachedRolePoliciesRequest) {
    request = &ListAttachedRolePoliciesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListAttachedRolePolicies")
    return
}

func NewListAttachedRolePoliciesResponse() (response *ListAttachedRolePoliciesResponse) {
    response = &ListAttachedRolePoliciesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAttachedRolePolicies(request *ListAttachedRolePoliciesRequest) (string) {
    return c.ListAttachedRolePoliciesWithContext(context.Background(), request)
}

func (c *Client) ListAttachedRolePoliciesWithContext(ctx context.Context, request *ListAttachedRolePoliciesRequest) (string) {
    if request == nil {
        request = NewListAttachedRolePoliciesRequest()
    }
    request.SetContext(ctx)

    response := NewListAttachedRolePoliciesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateRoleTrustAccountsRequest() (request *UpdateRoleTrustAccountsRequest) {
    request = &UpdateRoleTrustAccountsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "UpdateRoleTrustAccounts")
    return
}

func NewUpdateRoleTrustAccountsResponse() (response *UpdateRoleTrustAccountsResponse) {
    response = &UpdateRoleTrustAccountsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateRoleTrustAccounts(request *UpdateRoleTrustAccountsRequest) (string) {
    return c.UpdateRoleTrustAccountsWithContext(context.Background(), request)
}

func (c *Client) UpdateRoleTrustAccountsWithContext(ctx context.Context, request *UpdateRoleTrustAccountsRequest) (string) {
    if request == nil {
        request = NewUpdateRoleTrustAccountsRequest()
    }
    request.SetContext(ctx)

    response := NewUpdateRoleTrustAccountsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListEntityForPolicyRequest() (request *ListEntityForPolicyRequest) {
    request = &ListEntityForPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListEntityForPolicy")
    return
}

func NewListEntityForPolicyResponse() (response *ListEntityForPolicyResponse) {
    response = &ListEntityForPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListEntityForPolicy(request *ListEntityForPolicyRequest) (string) {
    return c.ListEntityForPolicyWithContext(context.Background(), request)
}

func (c *Client) ListEntityForPolicyWithContext(ctx context.Context, request *ListEntityForPolicyRequest) (string) {
    if request == nil {
        request = NewListEntityForPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewListEntityForPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreateProject")
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateProject(request *CreateProjectRequest) (string) {
    return c.CreateProjectWithContext(context.Background(), request)
}

func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (string) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    request.SetContext(ctx)

    response := NewCreateProjectResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListEntitiesForPolicyRequest() (request *ListEntitiesForPolicyRequest) {
    request = &ListEntitiesForPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListEntitiesForPolicy")
    return
}

func NewListEntitiesForPolicyResponse() (response *ListEntitiesForPolicyResponse) {
    response = &ListEntitiesForPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (string) {
    return c.ListEntitiesForPolicyWithContext(context.Background(), request)
}

func (c *Client) ListEntitiesForPolicyWithContext(ctx context.Context, request *ListEntitiesForPolicyRequest) (string) {
    if request == nil {
        request = NewListEntitiesForPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewListEntitiesForPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListProjectMemberRequest() (request *ListProjectMemberRequest) {
    request = &ListProjectMemberRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListProjectMember")
    return
}

func NewListProjectMemberResponse() (response *ListProjectMemberResponse) {
    response = &ListProjectMemberResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListProjectMember(request *ListProjectMemberRequest) (string) {
    return c.ListProjectMemberWithContext(context.Background(), request)
}

func (c *Client) ListProjectMemberWithContext(ctx context.Context, request *ListProjectMemberRequest) (string) {
    if request == nil {
        request = NewListProjectMemberRequest()
    }
    request.SetContext(ctx)

    response := NewListProjectMemberResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteProjectMemberRequest() (request *DeleteProjectMemberRequest) {
    request = &DeleteProjectMemberRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeleteProjectMember")
    return
}

func NewDeleteProjectMemberResponse() (response *DeleteProjectMemberResponse) {
    response = &DeleteProjectMemberResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteProjectMember(request *DeleteProjectMemberRequest) (string) {
    return c.DeleteProjectMemberWithContext(context.Background(), request)
}

func (c *Client) DeleteProjectMemberWithContext(ctx context.Context, request *DeleteProjectMemberRequest) (string) {
    if request == nil {
        request = NewDeleteProjectMemberRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteProjectMemberResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddProjectMemberRequest() (request *AddProjectMemberRequest) {
    request = &AddProjectMemberRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "AddProjectMember")
    return
}

func NewAddProjectMemberResponse() (response *AddProjectMemberResponse) {
    response = &AddProjectMemberResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddProjectMember(request *AddProjectMemberRequest) (string) {
    return c.AddProjectMemberWithContext(context.Background(), request)
}

func (c *Client) AddProjectMemberWithContext(ctx context.Context, request *AddProjectMemberRequest) (string) {
    if request == nil {
        request = NewAddProjectMemberRequest()
    }
    request.SetContext(ctx)

    response := NewAddProjectMemberResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateRoleRequest() (request *UpdateRoleRequest) {
    request = &UpdateRoleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "UpdateRole")
    return
}

func NewUpdateRoleResponse() (response *UpdateRoleResponse) {
    response = &UpdateRoleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateRole(request *UpdateRoleRequest) (string) {
    return c.UpdateRoleWithContext(context.Background(), request)
}

func (c *Client) UpdateRoleWithContext(ctx context.Context, request *UpdateRoleRequest) (string) {
    if request == nil {
        request = NewUpdateRoleRequest()
    }
    request.SetContext(ctx)

    response := NewUpdateRoleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "UpdatePolicy")
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (string) {
    return c.UpdatePolicyWithContext(context.Background(), request)
}

func (c *Client) UpdatePolicyWithContext(ctx context.Context, request *UpdatePolicyRequest) (string) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    request.SetContext(ctx)

    response := NewUpdatePolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateGroup(request *CreateGroupRequest) (string) {
    return c.CreateGroupWithContext(context.Background(), request)
}

func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (string) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    request.SetContext(ctx)

    response := NewCreateGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteGroup(request *DeleteGroupRequest) (string) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (string) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDetachGroupPolicyRequest() (request *DetachGroupPolicyRequest) {
    request = &DetachGroupPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "DetachGroupPolicy")
    return
}

func NewDetachGroupPolicyResponse() (response *DetachGroupPolicyResponse) {
    response = &DetachGroupPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DetachGroupPolicy(request *DetachGroupPolicyRequest) (string) {
    return c.DetachGroupPolicyWithContext(context.Background(), request)
}

func (c *Client) DetachGroupPolicyWithContext(ctx context.Context, request *DetachGroupPolicyRequest) (string) {
    if request == nil {
        request = NewDetachGroupPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewDetachGroupPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAttachGroupPolicyRequest() (request *AttachGroupPolicyRequest) {
    request = &AttachGroupPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "AttachGroupPolicy")
    return
}

func NewAttachGroupPolicyResponse() (response *AttachGroupPolicyResponse) {
    response = &AttachGroupPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AttachGroupPolicy(request *AttachGroupPolicyRequest) (string) {
    return c.AttachGroupPolicyWithContext(context.Background(), request)
}

func (c *Client) AttachGroupPolicyWithContext(ctx context.Context, request *AttachGroupPolicyRequest) (string) {
    if request == nil {
        request = NewAttachGroupPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewAttachGroupPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListGroupPoliciesRequest() (request *ListGroupPoliciesRequest) {
    request = &ListGroupPoliciesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListGroupPolicies")
    return
}

func NewListGroupPoliciesResponse() (response *ListGroupPoliciesResponse) {
    response = &ListGroupPoliciesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListGroupPolicies(request *ListGroupPoliciesRequest) (string) {
    return c.ListGroupPoliciesWithContext(context.Background(), request)
}

func (c *Client) ListGroupPoliciesWithContext(ctx context.Context, request *ListGroupPoliciesRequest) (string) {
    if request == nil {
        request = NewListGroupPoliciesRequest()
    }
    request.SetContext(ctx)

    response := NewListGroupPoliciesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
    request = &AddUserToGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "AddUserToGroup")
    return
}

func NewAddUserToGroupResponse() (response *AddUserToGroupResponse) {
    response = &AddUserToGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (string) {
    return c.AddUserToGroupWithContext(context.Background(), request)
}

func (c *Client) AddUserToGroupWithContext(ctx context.Context, request *AddUserToGroupRequest) (string) {
    if request == nil {
        request = NewAddUserToGroupRequest()
    }
    request.SetContext(ctx)

    response := NewAddUserToGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetGroupRequest() (request *GetGroupRequest) {
    request = &GetGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "GetGroup")
    return
}

func NewGetGroupResponse() (response *GetGroupResponse) {
    response = &GetGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetGroup(request *GetGroupRequest) (string) {
    return c.GetGroupWithContext(context.Background(), request)
}

func (c *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest) (string) {
    if request == nil {
        request = NewGetGroupRequest()
    }
    request.SetContext(ctx)

    response := NewGetGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListGroupsForUserRequest() (request *ListGroupsForUserRequest) {
    request = &ListGroupsForUserRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListGroupsForUser")
    return
}

func NewListGroupsForUserResponse() (response *ListGroupsForUserResponse) {
    response = &ListGroupsForUserResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (string) {
    return c.ListGroupsForUserWithContext(context.Background(), request)
}

func (c *Client) ListGroupsForUserWithContext(ctx context.Context, request *ListGroupsForUserRequest) (string) {
    if request == nil {
        request = NewListGroupsForUserRequest()
    }
    request.SetContext(ctx)

    response := NewListGroupsForUserResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListGroupsRequest() (request *ListGroupsRequest) {
    request = &ListGroupsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "ListGroups")
    return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
    response = &ListGroupsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListGroups(request *ListGroupsRequest) (string) {
    return c.ListGroupsWithContext(context.Background(), request)
}

func (c *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest) (string) {
    if request == nil {
        request = NewListGroupsRequest()
    }
    request.SetContext(ctx)

    response := NewListGroupsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
    request = &RemoveUserFromGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iam", APIVersion, "RemoveUserFromGroup")
    return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
    response = &RemoveUserFromGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (string) {
    return c.RemoveUserFromGroupWithContext(context.Background(), request)
}

func (c *Client) RemoveUserFromGroupWithContext(ctx context.Context, request *RemoveUserFromGroupRequest) (string) {
    if request == nil {
        request = NewRemoveUserFromGroupRequest()
    }
    request.SetContext(ctx)

    response := NewRemoveUserFromGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


