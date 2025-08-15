package v20210320

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) CreateFolder(request *CreateFolderRequest) string {
	return c.CreateFolderWithContext(context.Background(), request)
}

func (c *Client) CreateFolderSend(request *CreateFolderRequest) (*CreateFolderResponse, error) {
	statusCode, msg, err := c.CreateFolderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateFolderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) string {
	if request == nil {
		request = NewCreateFolderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFolderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateFolderWithContextV2(ctx context.Context, request *CreateFolderRequest) (int, string, error) {
	if request == nil {
		request = NewCreateFolderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFolderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteFolder(request *DeleteFolderRequest) string {
	return c.DeleteFolderWithContext(context.Background(), request)
}

func (c *Client) DeleteFolderSend(request *DeleteFolderRequest) (*DeleteFolderResponse, error) {
	statusCode, msg, err := c.DeleteFolderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteFolderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest) string {
	if request == nil {
		request = NewDeleteFolderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFolderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFolderWithContextV2(ctx context.Context, request *DeleteFolderRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFolderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFolderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateFolder(request *UpdateFolderRequest) string {
	return c.UpdateFolderWithContext(context.Background(), request)
}

func (c *Client) UpdateFolderSend(request *UpdateFolderRequest) (*UpdateFolderResponse, error) {
	statusCode, msg, err := c.UpdateFolderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateFolderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateFolderWithContext(ctx context.Context, request *UpdateFolderRequest) string {
	if request == nil {
		request = NewUpdateFolderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateFolderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateFolderWithContextV2(ctx context.Context, request *UpdateFolderRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateFolderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateFolderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListAccountsForParent(request *ListAccountsForParentRequest) string {
	return c.ListAccountsForParentWithContext(context.Background(), request)
}

func (c *Client) ListAccountsForParentSend(request *ListAccountsForParentRequest) (*ListAccountsForParentResponse, error) {
	statusCode, msg, err := c.ListAccountsForParentWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListAccountsForParentResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListAccountsForParentWithContext(ctx context.Context, request *ListAccountsForParentRequest) string {
	if request == nil {
		request = NewListAccountsForParentRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAccountsForParentResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListAccountsForParentWithContextV2(ctx context.Context, request *ListAccountsForParentRequest) (int, string, error) {
	if request == nil {
		request = NewListAccountsForParentRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAccountsForParentResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) MoveAccount(request *MoveAccountRequest) string {
	return c.MoveAccountWithContext(context.Background(), request)
}

func (c *Client) MoveAccountSend(request *MoveAccountRequest) (*MoveAccountResponse, error) {
	statusCode, msg, err := c.MoveAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct MoveAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) MoveAccountWithContext(ctx context.Context, request *MoveAccountRequest) string {
	if request == nil {
		request = NewMoveAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewMoveAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) MoveAccountWithContextV2(ctx context.Context, request *MoveAccountRequest) (int, string, error) {
	if request == nil {
		request = NewMoveAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewMoveAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateAccount(request *UpdateAccountRequest) string {
	return c.UpdateAccountWithContext(context.Background(), request)
}

func (c *Client) UpdateAccountSend(request *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	statusCode, msg, err := c.UpdateAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateAccountWithContext(ctx context.Context, request *UpdateAccountRequest) string {
	if request == nil {
		request = NewUpdateAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateAccountWithContextV2(ctx context.Context, request *UpdateAccountRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListAccounts(request *ListAccountsRequest) string {
	return c.ListAccountsWithContext(context.Background(), request)
}

func (c *Client) ListAccountsSend(request *ListAccountsRequest) (*ListAccountsResponse, error) {
	statusCode, msg, err := c.ListAccountsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListAccountsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListAccountsWithContext(ctx context.Context, request *ListAccountsRequest) string {
	if request == nil {
		request = NewListAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAccountsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListAccountsWithContextV2(ctx context.Context, request *ListAccountsRequest) (int, string, error) {
	if request == nil {
		request = NewListAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAccountsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListFolders(request *ListFoldersRequest) string {
	return c.ListFoldersWithContext(context.Background(), request)
}

func (c *Client) ListFoldersSend(request *ListFoldersRequest) (*ListFoldersResponse, error) {
	statusCode, msg, err := c.ListFoldersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListFoldersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListFoldersWithContext(ctx context.Context, request *ListFoldersRequest) string {
	if request == nil {
		request = NewListFoldersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListFoldersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListFoldersWithContextV2(ctx context.Context, request *ListFoldersRequest) (int, string, error) {
	if request == nil {
		request = NewListFoldersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListFoldersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
