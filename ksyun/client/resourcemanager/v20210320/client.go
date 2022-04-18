package v20210320
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2021-03-20"

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

func NewCreateFolderRequest() (request *CreateFolderRequest) {
    request = &CreateFolderRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "CreateFolder")
    return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
    response = &CreateFolderResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    return c.CreateFolderWithContext(context.Background(), request)
}

func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    if request == nil {
        request = NewCreateFolderRequest()
    }
    request.SetContext(ctx)

    response = NewCreateFolderResponse()
    err = c.Send(request, response)
    return
}
func NewDeleteFolderRequest() (request *DeleteFolderRequest) {
    request = &DeleteFolderRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "DeleteFolder")
    return
}

func NewDeleteFolderResponse() (response *DeleteFolderResponse) {
    response = &DeleteFolderResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteFolder(request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    return c.DeleteFolderWithContext(context.Background(), request)
}

func (c *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    if request == nil {
        request = NewDeleteFolderRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteFolderResponse()
    err = c.Send(request, response)
    return
}
func NewUpdateFolderRequest() (request *UpdateFolderRequest) {
    request = &UpdateFolderRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "UpdateFolder")
    return
}

func NewUpdateFolderResponse() (response *UpdateFolderResponse) {
    response = &UpdateFolderResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateFolder(request *UpdateFolderRequest) (response *UpdateFolderResponse, err error) {
    return c.UpdateFolderWithContext(context.Background(), request)
}

func (c *Client) UpdateFolderWithContext(ctx context.Context, request *UpdateFolderRequest) (response *UpdateFolderResponse, err error) {
    if request == nil {
        request = NewUpdateFolderRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateFolderResponse()
    err = c.Send(request, response)
    return
}
func NewListAccountsForParentRequest() (request *ListAccountsForParentRequest) {
    request = &ListAccountsForParentRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "ListAccountsForParent")
    return
}

func NewListAccountsForParentResponse() (response *ListAccountsForParentResponse) {
    response = &ListAccountsForParentResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAccountsForParent(request *ListAccountsForParentRequest) (response *ListAccountsForParentResponse, err error) {
    return c.ListAccountsForParentWithContext(context.Background(), request)
}

func (c *Client) ListAccountsForParentWithContext(ctx context.Context, request *ListAccountsForParentRequest) (response *ListAccountsForParentResponse, err error) {
    if request == nil {
        request = NewListAccountsForParentRequest()
    }
    request.SetContext(ctx)

    response = NewListAccountsForParentResponse()
    err = c.Send(request, response)
    return
}
func NewMoveAccountRequest() (request *MoveAccountRequest) {
    request = &MoveAccountRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "MoveAccount")
    return
}

func NewMoveAccountResponse() (response *MoveAccountResponse) {
    response = &MoveAccountResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) MoveAccount(request *MoveAccountRequest) (response *MoveAccountResponse, err error) {
    return c.MoveAccountWithContext(context.Background(), request)
}

func (c *Client) MoveAccountWithContext(ctx context.Context, request *MoveAccountRequest) (response *MoveAccountResponse, err error) {
    if request == nil {
        request = NewMoveAccountRequest()
    }
    request.SetContext(ctx)

    response = NewMoveAccountResponse()
    err = c.Send(request, response)
    return
}
func NewUpdateAccountRequest() (request *UpdateAccountRequest) {
    request = &UpdateAccountRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "UpdateAccount")
    return
}

func NewUpdateAccountResponse() (response *UpdateAccountResponse) {
    response = &UpdateAccountResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateAccount(request *UpdateAccountRequest) (response *UpdateAccountResponse, err error) {
    return c.UpdateAccountWithContext(context.Background(), request)
}

func (c *Client) UpdateAccountWithContext(ctx context.Context, request *UpdateAccountRequest) (response *UpdateAccountResponse, err error) {
    if request == nil {
        request = NewUpdateAccountRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateAccountResponse()
    err = c.Send(request, response)
    return
}
func NewListAccountsRequest() (request *ListAccountsRequest) {
    request = &ListAccountsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "ListAccounts")
    return
}

func NewListAccountsResponse() (response *ListAccountsResponse) {
    response = &ListAccountsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAccounts(request *ListAccountsRequest) (response *ListAccountsResponse, err error) {
    return c.ListAccountsWithContext(context.Background(), request)
}

func (c *Client) ListAccountsWithContext(ctx context.Context, request *ListAccountsRequest) (response *ListAccountsResponse, err error) {
    if request == nil {
        request = NewListAccountsRequest()
    }
    request.SetContext(ctx)

    response = NewListAccountsResponse()
    err = c.Send(request, response)
    return
}
func NewListFoldersRequest() (request *ListFoldersRequest) {
    request = &ListFoldersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("resourcemanager", APIVersion, "ListFolders")
    return
}

func NewListFoldersResponse() (response *ListFoldersResponse) {
    response = &ListFoldersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListFolders(request *ListFoldersRequest) (response *ListFoldersResponse, err error) {
    return c.ListFoldersWithContext(context.Background(), request)
}

func (c *Client) ListFoldersWithContext(ctx context.Context, request *ListFoldersRequest) (response *ListFoldersResponse, err error) {
    if request == nil {
        request = NewListFoldersRequest()
    }
    request.SetContext(ctx)

    response = NewListFoldersResponse()
    err = c.Send(request, response)
    return
}


