package v20151101
import (
    "context"
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

func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    request.SetContext(ctx)

    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    return c.ListUsersWithContext(context.Background(), request)
}

func (c *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest) (response *ListUsersResponse, err error) {
    if request == nil {
        request = NewListUsersRequest()
    }
    request.SetContext(ctx)

    response = NewListUsersResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    return c.UpdateUserWithContext(context.Background(), request)
}

func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateUserResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
    return c.GetUserWithContext(context.Background(), request)
}

func (c *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest) (response *GetUserResponse, err error) {
    if request == nil {
        request = NewGetUserRequest()
    }
    request.SetContext(ctx)

    response = NewGetUserResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListAttachedUserPolicies(request *ListAttachedUserPoliciesRequest) (response *ListAttachedUserPoliciesResponse, err error) {
    return c.ListAttachedUserPoliciesWithContext(context.Background(), request)
}

func (c *Client) ListAttachedUserPoliciesWithContext(ctx context.Context, request *ListAttachedUserPoliciesRequest) (response *ListAttachedUserPoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedUserPoliciesRequest()
    }
    request.SetContext(ctx)

    response = NewListAttachedUserPoliciesResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (response *ListPolicyVersionsResponse, err error) {
    return c.ListPolicyVersionsWithContext(context.Background(), request)
}

func (c *Client) ListPolicyVersionsWithContext(ctx context.Context, request *ListPolicyVersionsRequest) (response *ListPolicyVersionsResponse, err error) {
    if request == nil {
        request = NewListPolicyVersionsRequest()
    }
    request.SetContext(ctx)

    response = NewListPolicyVersionsResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) SetDefaultPolicyVersion(request *SetDefaultPolicyVersionRequest) (response *SetDefaultPolicyVersionResponse, err error) {
    return c.SetDefaultPolicyVersionWithContext(context.Background(), request)
}

func (c *Client) SetDefaultPolicyVersionWithContext(ctx context.Context, request *SetDefaultPolicyVersionRequest) (response *SetDefaultPolicyVersionResponse, err error) {
    if request == nil {
        request = NewSetDefaultPolicyVersionRequest()
    }
    request.SetContext(ctx)

    response = NewSetDefaultPolicyVersionResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) AttachUserPolicy(request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    return c.AttachUserPolicyWithContext(context.Background(), request)
}

func (c *Client) AttachUserPolicyWithContext(ctx context.Context, request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    if request == nil {
        request = NewAttachUserPolicyRequest()
    }
    request.SetContext(ctx)

    response = NewAttachUserPolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (response *DeletePolicyVersionResponse, err error) {
    return c.DeletePolicyVersionWithContext(context.Background(), request)
}

func (c *Client) DeletePolicyVersionWithContext(ctx context.Context, request *DeletePolicyVersionRequest) (response *DeletePolicyVersionResponse, err error) {
    if request == nil {
        request = NewDeletePolicyVersionRequest()
    }
    request.SetContext(ctx)

    response = NewDeletePolicyVersionResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) GetPolicyVersion(request *GetPolicyVersionRequest) (response *GetPolicyVersionResponse, err error) {
    return c.GetPolicyVersionWithContext(context.Background(), request)
}

func (c *Client) GetPolicyVersionWithContext(ctx context.Context, request *GetPolicyVersionRequest) (response *GetPolicyVersionResponse, err error) {
    if request == nil {
        request = NewGetPolicyVersionRequest()
    }
    request.SetContext(ctx)

    response = NewGetPolicyVersionResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (response *CreatePolicyVersionResponse, err error) {
    return c.CreatePolicyVersionWithContext(context.Background(), request)
}

func (c *Client) CreatePolicyVersionWithContext(ctx context.Context, request *CreatePolicyVersionRequest) (response *CreatePolicyVersionResponse, err error) {
    if request == nil {
        request = NewCreatePolicyVersionRequest()
    }
    request.SetContext(ctx)

    response = NewCreatePolicyVersionResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    return c.ListPoliciesWithContext(context.Background(), request)
}

func (c *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    if request == nil {
        request = NewListPoliciesRequest()
    }
    request.SetContext(ctx)

    response = NewListPoliciesResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) GetPolicy(request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
    return c.GetPolicyWithContext(context.Background(), request)
}

func (c *Client) GetPolicyWithContext(ctx context.Context, request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
    if request == nil {
        request = NewGetPolicyRequest()
    }
    request.SetContext(ctx)

    response = NewGetPolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    return c.DeletePolicyWithContext(context.Background(), request)
}

func (c *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    if request == nil {
        request = NewDeletePolicyRequest()
    }
    request.SetContext(ctx)

    response = NewDeletePolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    return c.CreatePolicyWithContext(context.Background(), request)
}

func (c *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    request.SetContext(ctx)

    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ChangePassword(request *ChangePasswordRequest) (response *ChangePasswordResponse, err error) {
    return c.ChangePasswordWithContext(context.Background(), request)
}

func (c *Client) ChangePasswordWithContext(ctx context.Context, request *ChangePasswordRequest) (response *ChangePasswordResponse, err error) {
    if request == nil {
        request = NewChangePasswordRequest()
    }
    request.SetContext(ctx)

    response = NewChangePasswordResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (response *UpdateLoginProfileResponse, err error) {
    return c.UpdateLoginProfileWithContext(context.Background(), request)
}

func (c *Client) UpdateLoginProfileWithContext(ctx context.Context, request *UpdateLoginProfileRequest) (response *UpdateLoginProfileResponse, err error) {
    if request == nil {
        request = NewUpdateLoginProfileRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateLoginProfileResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) GetLoginProfile(request *GetLoginProfileRequest) (response *GetLoginProfileResponse, err error) {
    return c.GetLoginProfileWithContext(context.Background(), request)
}

func (c *Client) GetLoginProfileWithContext(ctx context.Context, request *GetLoginProfileRequest) (response *GetLoginProfileResponse, err error) {
    if request == nil {
        request = NewGetLoginProfileRequest()
    }
    request.SetContext(ctx)

    response = NewGetLoginProfileResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) CreateAccessKey(request *CreateAccessKeyRequest) (response *CreateAccessKeyResponse, err error) {
    return c.CreateAccessKeyWithContext(context.Background(), request)
}

func (c *Client) CreateAccessKeyWithContext(ctx context.Context, request *CreateAccessKeyRequest) (response *CreateAccessKeyResponse, err error) {
    if request == nil {
        request = NewCreateAccessKeyRequest()
    }
    request.SetContext(ctx)

    response = NewCreateAccessKeyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListAccessKeys(request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
    return c.ListAccessKeysWithContext(context.Background(), request)
}

func (c *Client) ListAccessKeysWithContext(ctx context.Context, request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
    if request == nil {
        request = NewListAccessKeysRequest()
    }
    request.SetContext(ctx)

    response = NewListAccessKeysResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) UpdateAccessKey(request *UpdateAccessKeyRequest) (response *UpdateAccessKeyResponse, err error) {
    return c.UpdateAccessKeyWithContext(context.Background(), request)
}

func (c *Client) UpdateAccessKeyWithContext(ctx context.Context, request *UpdateAccessKeyRequest) (response *UpdateAccessKeyResponse, err error) {
    if request == nil {
        request = NewUpdateAccessKeyRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateAccessKeyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (response *DeleteAccessKeyResponse, err error) {
    return c.DeleteAccessKeyWithContext(context.Background(), request)
}

func (c *Client) DeleteAccessKeyWithContext(ctx context.Context, request *DeleteAccessKeyRequest) (response *DeleteAccessKeyResponse, err error) {
    if request == nil {
        request = NewDeleteAccessKeyRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteAccessKeyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    return c.CreateRoleWithContext(context.Background(), request)
}

func (c *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    request.SetContext(ctx)

    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    return c.DeleteRoleWithContext(context.Background(), request)
}

func (c *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteRoleResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) GetRole(request *GetRoleRequest) (response *GetRoleResponse, err error) {
    return c.GetRoleWithContext(context.Background(), request)
}

func (c *Client) GetRoleWithContext(ctx context.Context, request *GetRoleRequest) (response *GetRoleResponse, err error) {
    if request == nil {
        request = NewGetRoleRequest()
    }
    request.SetContext(ctx)

    response = NewGetRoleResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListRoles(request *ListRolesRequest) (response *ListRolesResponse, err error) {
    return c.ListRolesWithContext(context.Background(), request)
}

func (c *Client) ListRolesWithContext(ctx context.Context, request *ListRolesRequest) (response *ListRolesResponse, err error) {
    if request == nil {
        request = NewListRolesRequest()
    }
    request.SetContext(ctx)

    response = NewListRolesResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) AttachRolePolicy(request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
    return c.AttachRolePolicyWithContext(context.Background(), request)
}

func (c *Client) AttachRolePolicyWithContext(ctx context.Context, request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
    if request == nil {
        request = NewAttachRolePolicyRequest()
    }
    request.SetContext(ctx)

    response = NewAttachRolePolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DetachRolePolicy(request *DetachRolePolicyRequest) (response *DetachRolePolicyResponse, err error) {
    return c.DetachRolePolicyWithContext(context.Background(), request)
}

func (c *Client) DetachRolePolicyWithContext(ctx context.Context, request *DetachRolePolicyRequest) (response *DetachRolePolicyResponse, err error) {
    if request == nil {
        request = NewDetachRolePolicyRequest()
    }
    request.SetContext(ctx)

    response = NewDetachRolePolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListAttachedRolePolicies(request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
    return c.ListAttachedRolePoliciesWithContext(context.Background(), request)
}

func (c *Client) ListAttachedRolePoliciesWithContext(ctx context.Context, request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedRolePoliciesRequest()
    }
    request.SetContext(ctx)

    response = NewListAttachedRolePoliciesResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) UpdateRoleTrustAccounts(request *UpdateRoleTrustAccountsRequest) (response *UpdateRoleTrustAccountsResponse, err error) {
    return c.UpdateRoleTrustAccountsWithContext(context.Background(), request)
}

func (c *Client) UpdateRoleTrustAccountsWithContext(ctx context.Context, request *UpdateRoleTrustAccountsRequest) (response *UpdateRoleTrustAccountsResponse, err error) {
    if request == nil {
        request = NewUpdateRoleTrustAccountsRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateRoleTrustAccountsResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListEntityForPolicy(request *ListEntityForPolicyRequest) (response *ListEntityForPolicyResponse, err error) {
    return c.ListEntityForPolicyWithContext(context.Background(), request)
}

func (c *Client) ListEntityForPolicyWithContext(ctx context.Context, request *ListEntityForPolicyRequest) (response *ListEntityForPolicyResponse, err error) {
    if request == nil {
        request = NewListEntityForPolicyRequest()
    }
    request.SetContext(ctx)

    response = NewListEntityForPolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    request.SetContext(ctx)

    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
    return c.ListEntitiesForPolicyWithContext(context.Background(), request)
}

func (c *Client) ListEntitiesForPolicyWithContext(ctx context.Context, request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
    if request == nil {
        request = NewListEntitiesForPolicyRequest()
    }
    request.SetContext(ctx)

    response = NewListEntitiesForPolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListProjectMember(request *ListProjectMemberRequest) (response *ListProjectMemberResponse, err error) {
    return c.ListProjectMemberWithContext(context.Background(), request)
}

func (c *Client) ListProjectMemberWithContext(ctx context.Context, request *ListProjectMemberRequest) (response *ListProjectMemberResponse, err error) {
    if request == nil {
        request = NewListProjectMemberRequest()
    }
    request.SetContext(ctx)

    response = NewListProjectMemberResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeleteProjectMember(request *DeleteProjectMemberRequest) (response *DeleteProjectMemberResponse, err error) {
    return c.DeleteProjectMemberWithContext(context.Background(), request)
}

func (c *Client) DeleteProjectMemberWithContext(ctx context.Context, request *DeleteProjectMemberRequest) (response *DeleteProjectMemberResponse, err error) {
    if request == nil {
        request = NewDeleteProjectMemberRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteProjectMemberResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) AddProjectMember(request *AddProjectMemberRequest) (response *AddProjectMemberResponse, err error) {
    return c.AddProjectMemberWithContext(context.Background(), request)
}

func (c *Client) AddProjectMemberWithContext(ctx context.Context, request *AddProjectMemberRequest) (response *AddProjectMemberResponse, err error) {
    if request == nil {
        request = NewAddProjectMemberRequest()
    }
    request.SetContext(ctx)

    response = NewAddProjectMemberResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) UpdateRole(request *UpdateRoleRequest) (response *UpdateRoleResponse, err error) {
    return c.UpdateRoleWithContext(context.Background(), request)
}

func (c *Client) UpdateRoleWithContext(ctx context.Context, request *UpdateRoleRequest) (response *UpdateRoleResponse, err error) {
    if request == nil {
        request = NewUpdateRoleRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateRoleResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    return c.UpdatePolicyWithContext(context.Background(), request)
}

func (c *Client) UpdatePolicyWithContext(ctx context.Context, request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    request.SetContext(ctx)

    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    request.SetContext(ctx)

    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) DetachGroupPolicy(request *DetachGroupPolicyRequest) (response *DetachGroupPolicyResponse, err error) {
    return c.DetachGroupPolicyWithContext(context.Background(), request)
}

func (c *Client) DetachGroupPolicyWithContext(ctx context.Context, request *DetachGroupPolicyRequest) (response *DetachGroupPolicyResponse, err error) {
    if request == nil {
        request = NewDetachGroupPolicyRequest()
    }
    request.SetContext(ctx)

    response = NewDetachGroupPolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) AttachGroupPolicy(request *AttachGroupPolicyRequest) (response *AttachGroupPolicyResponse, err error) {
    return c.AttachGroupPolicyWithContext(context.Background(), request)
}

func (c *Client) AttachGroupPolicyWithContext(ctx context.Context, request *AttachGroupPolicyRequest) (response *AttachGroupPolicyResponse, err error) {
    if request == nil {
        request = NewAttachGroupPolicyRequest()
    }
    request.SetContext(ctx)

    response = NewAttachGroupPolicyResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListGroupPolicies(request *ListGroupPoliciesRequest) (response *ListGroupPoliciesResponse, err error) {
    return c.ListGroupPoliciesWithContext(context.Background(), request)
}

func (c *Client) ListGroupPoliciesWithContext(ctx context.Context, request *ListGroupPoliciesRequest) (response *ListGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewListGroupPoliciesRequest()
    }
    request.SetContext(ctx)

    response = NewListGroupPoliciesResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    return c.AddUserToGroupWithContext(context.Background(), request)
}

func (c *Client) AddUserToGroupWithContext(ctx context.Context, request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToGroupRequest()
    }
    request.SetContext(ctx)

    response = NewAddUserToGroupResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
    return c.GetGroupWithContext(context.Background(), request)
}

func (c *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest) (response *GetGroupResponse, err error) {
    if request == nil {
        request = NewGetGroupRequest()
    }
    request.SetContext(ctx)

    response = NewGetGroupResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (response *ListGroupsForUserResponse, err error) {
    return c.ListGroupsForUserWithContext(context.Background(), request)
}

func (c *Client) ListGroupsForUserWithContext(ctx context.Context, request *ListGroupsForUserRequest) (response *ListGroupsForUserResponse, err error) {
    if request == nil {
        request = NewListGroupsForUserRequest()
    }
    request.SetContext(ctx)

    response = NewListGroupsForUserResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    return c.ListGroupsWithContext(context.Background(), request)
}

func (c *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    if request == nil {
        request = NewListGroupsRequest()
    }
    request.SetContext(ctx)

    response = NewListGroupsResponse()
    err = c.Send(request, response)
    return
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

func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    return c.RemoveUserFromGroupWithContext(context.Background(), request)
}

func (c *Client) RemoveUserFromGroupWithContext(ctx context.Context, request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromGroupRequest()
    }
    request.SetContext(ctx)

    response = NewRemoveUserFromGroupResponse()
    err = c.Send(request, response)
    return
}


