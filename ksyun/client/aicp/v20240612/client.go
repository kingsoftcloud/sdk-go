package v20240612

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-06-12"

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

func NewCreateStorageConfigRequest() (request *CreateStorageConfigRequest) {
	request = &CreateStorageConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateStorageConfig")
	return
}

func NewCreateStorageConfigResponse() (response *CreateStorageConfigResponse) {
	response = &CreateStorageConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateStorageConfig(request *CreateStorageConfigRequest) string {
	return c.CreateStorageConfigWithContext(context.Background(), request)
}

func (c *Client) CreateStorageConfigSend(request *CreateStorageConfigRequest) (*CreateStorageConfigResponse, error) {
	statusCode, msg, err := c.CreateStorageConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateStorageConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateStorageConfigWithContext(ctx context.Context, request *CreateStorageConfigRequest) string {
	if request == nil {
		request = NewCreateStorageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateStorageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateStorageConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateStorageConfigWithContextV2(ctx context.Context, request *CreateStorageConfigRequest) (int, string, error) {
	if request == nil {
		request = NewCreateStorageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateStorageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateStorageConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyStorageConfigRequest() (request *ModifyStorageConfigRequest) {
	request = &ModifyStorageConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyStorageConfig")
	return
}

func NewModifyStorageConfigResponse() (response *ModifyStorageConfigResponse) {
	response = &ModifyStorageConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyStorageConfig(request *ModifyStorageConfigRequest) string {
	return c.ModifyStorageConfigWithContext(context.Background(), request)
}

func (c *Client) ModifyStorageConfigSend(request *ModifyStorageConfigRequest) (*ModifyStorageConfigResponse, error) {
	statusCode, msg, err := c.ModifyStorageConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyStorageConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyStorageConfigWithContext(ctx context.Context, request *ModifyStorageConfigRequest) string {
	if request == nil {
		request = NewModifyStorageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyStorageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyStorageConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyStorageConfigWithContextV2(ctx context.Context, request *ModifyStorageConfigRequest) (int, string, error) {
	if request == nil {
		request = NewModifyStorageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyStorageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyStorageConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeStorageConfigsRequest() (request *DescribeStorageConfigsRequest) {
	request = &DescribeStorageConfigsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeStorageConfigs")
	return
}

func NewDescribeStorageConfigsResponse() (response *DescribeStorageConfigsResponse) {
	response = &DescribeStorageConfigsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeStorageConfigs(request *DescribeStorageConfigsRequest) string {
	return c.DescribeStorageConfigsWithContext(context.Background(), request)
}

func (c *Client) DescribeStorageConfigsSend(request *DescribeStorageConfigsRequest) (*DescribeStorageConfigsResponse, error) {
	statusCode, msg, err := c.DescribeStorageConfigsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeStorageConfigsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeStorageConfigsWithContext(ctx context.Context, request *DescribeStorageConfigsRequest) string {
	if request == nil {
		request = NewDescribeStorageConfigsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeStorageConfigs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStorageConfigsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeStorageConfigsWithContextV2(ctx context.Context, request *DescribeStorageConfigsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeStorageConfigsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeStorageConfigs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStorageConfigsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteStorageConfigRequest() (request *DeleteStorageConfigRequest) {
	request = &DeleteStorageConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteStorageConfig")
	return
}

func NewDeleteStorageConfigResponse() (response *DeleteStorageConfigResponse) {
	response = &DeleteStorageConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteStorageConfig(request *DeleteStorageConfigRequest) string {
	return c.DeleteStorageConfigWithContext(context.Background(), request)
}

func (c *Client) DeleteStorageConfigSend(request *DeleteStorageConfigRequest) (*DeleteStorageConfigResponse, error) {
	statusCode, msg, err := c.DeleteStorageConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteStorageConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteStorageConfigWithContext(ctx context.Context, request *DeleteStorageConfigRequest) string {
	if request == nil {
		request = NewDeleteStorageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteStorageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteStorageConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteStorageConfigWithContextV2(ctx context.Context, request *DeleteStorageConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteStorageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteStorageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteStorageConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSaveNotebookImageRequest() (request *SaveNotebookImageRequest) {
	request = &SaveNotebookImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "SaveNotebookImage")
	return
}

func NewSaveNotebookImageResponse() (response *SaveNotebookImageResponse) {
	response = &SaveNotebookImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SaveNotebookImage(request *SaveNotebookImageRequest) string {
	return c.SaveNotebookImageWithContext(context.Background(), request)
}

func (c *Client) SaveNotebookImageSend(request *SaveNotebookImageRequest) (*SaveNotebookImageResponse, error) {
	statusCode, msg, err := c.SaveNotebookImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SaveNotebookImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SaveNotebookImageWithContext(ctx context.Context, request *SaveNotebookImageRequest) string {
	if request == nil {
		request = NewSaveNotebookImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "SaveNotebookImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSaveNotebookImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SaveNotebookImageWithContextV2(ctx context.Context, request *SaveNotebookImageRequest) (int, string, error) {
	if request == nil {
		request = NewSaveNotebookImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "SaveNotebookImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSaveNotebookImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNotebookRequest() (request *ModifyNotebookRequest) {
	request = &ModifyNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyNotebook")
	return
}

func NewModifyNotebookResponse() (response *ModifyNotebookResponse) {
	response = &ModifyNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNotebook(request *ModifyNotebookRequest) string {
	return c.ModifyNotebookWithContext(context.Background(), request)
}

func (c *Client) ModifyNotebookSend(request *ModifyNotebookRequest) (*ModifyNotebookResponse, error) {
	statusCode, msg, err := c.ModifyNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNotebookWithContext(ctx context.Context, request *ModifyNotebookRequest) string {
	if request == nil {
		request = NewModifyNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNotebookWithContextV2(ctx context.Context, request *ModifyNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNotebookRequest() (request *DeleteNotebookRequest) {
	request = &DeleteNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteNotebook")
	return
}

func NewDeleteNotebookResponse() (response *DeleteNotebookResponse) {
	response = &DeleteNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNotebook(request *DeleteNotebookRequest) string {
	return c.DeleteNotebookWithContext(context.Background(), request)
}

func (c *Client) DeleteNotebookSend(request *DeleteNotebookRequest) (*DeleteNotebookResponse, error) {
	statusCode, msg, err := c.DeleteNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNotebookWithContext(ctx context.Context, request *DeleteNotebookRequest) string {
	if request == nil {
		request = NewDeleteNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNotebookWithContextV2(ctx context.Context, request *DeleteNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNotebooksRequest() (request *DescribeNotebooksRequest) {
	request = &DescribeNotebooksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebooks")
	return
}

func NewDescribeNotebooksResponse() (response *DescribeNotebooksResponse) {
	response = &DescribeNotebooksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNotebooks(request *DescribeNotebooksRequest) string {
	return c.DescribeNotebooksWithContext(context.Background(), request)
}

func (c *Client) DescribeNotebooksSend(request *DescribeNotebooksRequest) (*DescribeNotebooksResponse, error) {
	statusCode, msg, err := c.DescribeNotebooksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNotebooksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNotebooksWithContext(ctx context.Context, request *DescribeNotebooksRequest) string {
	if request == nil {
		request = NewDescribeNotebooksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebooks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebooksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNotebooksWithContextV2(ctx context.Context, request *DescribeNotebooksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNotebooksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebooks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebooksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNotebookRequest() (request *CreateNotebookRequest) {
	request = &CreateNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateNotebook")
	return
}

func NewCreateNotebookResponse() (response *CreateNotebookResponse) {
	response = &CreateNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNotebook(request *CreateNotebookRequest) string {
	return c.CreateNotebookWithContext(context.Background(), request)
}

func (c *Client) CreateNotebookSend(request *CreateNotebookRequest) (*CreateNotebookResponse, error) {
	statusCode, msg, err := c.CreateNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNotebookWithContext(ctx context.Context, request *CreateNotebookRequest) string {
	if request == nil {
		request = NewCreateNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNotebookWithContextV2(ctx context.Context, request *CreateNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImage(request *CreateImageRequest) string {
	return c.CreateImageWithContext(context.Background(), request)
}

func (c *Client) CreateImageSend(request *CreateImageRequest) (*CreateImageResponse, error) {
	statusCode, msg, err := c.CreateImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) string {
	if request == nil {
		request = NewCreateImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateImageWithContextV2(ctx context.Context, request *CreateImageRequest) (int, string, error) {
	if request == nil {
		request = NewCreateImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteImage(request *DeleteImageRequest) string {
	return c.DeleteImageWithContext(context.Background(), request)
}

func (c *Client) DeleteImageSend(request *DeleteImageRequest) (*DeleteImageResponse, error) {
	statusCode, msg, err := c.DeleteImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) string {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteImageWithContextV2(ctx context.Context, request *DeleteImageRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyImageRequest() (request *ModifyImageRequest) {
	request = &ModifyImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyImage")
	return
}

func NewModifyImageResponse() (response *ModifyImageResponse) {
	response = &ModifyImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyImage(request *ModifyImageRequest) string {
	return c.ModifyImageWithContext(context.Background(), request)
}

func (c *Client) ModifyImageSend(request *ModifyImageRequest) (*ModifyImageResponse, error) {
	statusCode, msg, err := c.ModifyImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyImageWithContext(ctx context.Context, request *ModifyImageRequest) string {
	if request == nil {
		request = NewModifyImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyImageWithContextV2(ctx context.Context, request *ModifyImageRequest) (int, string, error) {
	if request == nil {
		request = NewModifyImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) string {
	return c.DescribeImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeImagesSend(request *DescribeImagesRequest) (*DescribeImagesResponse, error) {
	statusCode, msg, err := c.DescribeImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) string {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeImages")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImagesWithContextV2(ctx context.Context, request *DescribeImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeImages")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopNotebookRequest() (request *StopNotebookRequest) {
	request = &StopNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopNotebook")
	return
}

func NewStopNotebookResponse() (response *StopNotebookResponse) {
	response = &StopNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopNotebook(request *StopNotebookRequest) string {
	return c.StopNotebookWithContext(context.Background(), request)
}

func (c *Client) StopNotebookSend(request *StopNotebookRequest) (*StopNotebookResponse, error) {
	statusCode, msg, err := c.StopNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopNotebookWithContext(ctx context.Context, request *StopNotebookRequest) string {
	if request == nil {
		request = NewStopNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopNotebookWithContextV2(ctx context.Context, request *StopNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewStopNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartNotebookRequest() (request *StartNotebookRequest) {
	request = &StartNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StartNotebook")
	return
}

func NewStartNotebookResponse() (response *StartNotebookResponse) {
	response = &StartNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartNotebook(request *StartNotebookRequest) string {
	return c.StartNotebookWithContext(context.Background(), request)
}

func (c *Client) StartNotebookSend(request *StartNotebookRequest) (*StartNotebookResponse, error) {
	statusCode, msg, err := c.StartNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartNotebookWithContext(ctx context.Context, request *StartNotebookRequest) string {
	if request == nil {
		request = NewStartNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartNotebookWithContextV2(ctx context.Context, request *StartNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewStartNotebookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartNotebook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetWebIdeUrlRequest() (request *GetWebIdeUrlRequest) {
	request = &GetWebIdeUrlRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetWebIdeUrl")
	return
}

func NewGetWebIdeUrlResponse() (response *GetWebIdeUrlResponse) {
	response = &GetWebIdeUrlResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetWebIdeUrl(request *GetWebIdeUrlRequest) string {
	return c.GetWebIdeUrlWithContext(context.Background(), request)
}

func (c *Client) GetWebIdeUrlSend(request *GetWebIdeUrlRequest) (*GetWebIdeUrlResponse, error) {
	statusCode, msg, err := c.GetWebIdeUrlWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetWebIdeUrlResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetWebIdeUrlWithContext(ctx context.Context, request *GetWebIdeUrlRequest) string {
	if request == nil {
		request = NewGetWebIdeUrlRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetWebIdeUrl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetWebIdeUrlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetWebIdeUrlWithContextV2(ctx context.Context, request *GetWebIdeUrlRequest) (int, string, error) {
	if request == nil {
		request = NewGetWebIdeUrlRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetWebIdeUrl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetWebIdeUrlResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNotebookLogRequest() (request *DescribeNotebookLogRequest) {
	request = &DescribeNotebookLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebookLog")
	return
}

func NewDescribeNotebookLogResponse() (response *DescribeNotebookLogResponse) {
	response = &DescribeNotebookLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNotebookLog(request *DescribeNotebookLogRequest) string {
	return c.DescribeNotebookLogWithContext(context.Background(), request)
}

func (c *Client) DescribeNotebookLogSend(request *DescribeNotebookLogRequest) (*DescribeNotebookLogResponse, error) {
	statusCode, msg, err := c.DescribeNotebookLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNotebookLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNotebookLogWithContext(ctx context.Context, request *DescribeNotebookLogRequest) string {
	if request == nil {
		request = NewDescribeNotebookLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebookLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebookLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNotebookLogWithContextV2(ctx context.Context, request *DescribeNotebookLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNotebookLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebookLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebookLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopNotebookSavingImageRequest() (request *StopNotebookSavingImageRequest) {
	request = &StopNotebookSavingImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopNotebookSavingImage")
	return
}

func NewStopNotebookSavingImageResponse() (response *StopNotebookSavingImageResponse) {
	response = &StopNotebookSavingImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopNotebookSavingImage(request *StopNotebookSavingImageRequest) string {
	return c.StopNotebookSavingImageWithContext(context.Background(), request)
}

func (c *Client) StopNotebookSavingImageSend(request *StopNotebookSavingImageRequest) (*StopNotebookSavingImageResponse, error) {
	statusCode, msg, err := c.StopNotebookSavingImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopNotebookSavingImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopNotebookSavingImageWithContext(ctx context.Context, request *StopNotebookSavingImageRequest) string {
	if request == nil {
		request = NewStopNotebookSavingImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopNotebookSavingImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopNotebookSavingImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopNotebookSavingImageWithContextV2(ctx context.Context, request *StopNotebookSavingImageRequest) (int, string, error) {
	if request == nil {
		request = NewStopNotebookSavingImageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopNotebookSavingImage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopNotebookSavingImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyApikeyRequest() (request *ModifyApikeyRequest) {
	request = &ModifyApikeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyApikey")
	return
}

func NewModifyApikeyResponse() (response *ModifyApikeyResponse) {
	response = &ModifyApikeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyApikey(request *ModifyApikeyRequest) string {
	return c.ModifyApikeyWithContext(context.Background(), request)
}

func (c *Client) ModifyApikeySend(request *ModifyApikeyRequest) (*ModifyApikeyResponse, error) {
	statusCode, msg, err := c.ModifyApikeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyApikeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyApikeyWithContext(ctx context.Context, request *ModifyApikeyRequest) string {
	if request == nil {
		request = NewModifyApikeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyApikey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyApikeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyApikeyWithContextV2(ctx context.Context, request *ModifyApikeyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyApikeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyApikey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyApikeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateApiServiceRequest() (request *ActivateApiServiceRequest) {
	request = &ActivateApiServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ActivateApiService")
	return
}

func NewActivateApiServiceResponse() (response *ActivateApiServiceResponse) {
	response = &ActivateApiServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateApiService(request *ActivateApiServiceRequest) string {
	return c.ActivateApiServiceWithContext(context.Background(), request)
}

func (c *Client) ActivateApiServiceSend(request *ActivateApiServiceRequest) (*ActivateApiServiceResponse, error) {
	statusCode, msg, err := c.ActivateApiServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ActivateApiServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateApiServiceWithContext(ctx context.Context, request *ActivateApiServiceRequest) string {
	if request == nil {
		request = NewActivateApiServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateApiService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateApiServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateApiServiceWithContextV2(ctx context.Context, request *ActivateApiServiceRequest) (int, string, error) {
	if request == nil {
		request = NewActivateApiServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateApiService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateApiServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteApikeyRequest() (request *DeleteApikeyRequest) {
	request = &DeleteApikeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteApikey")
	return
}

func NewDeleteApikeyResponse() (response *DeleteApikeyResponse) {
	response = &DeleteApikeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteApikey(request *DeleteApikeyRequest) string {
	return c.DeleteApikeyWithContext(context.Background(), request)
}

func (c *Client) DeleteApikeySend(request *DeleteApikeyRequest) (*DeleteApikeyResponse, error) {
	statusCode, msg, err := c.DeleteApikeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteApikeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteApikeyWithContext(ctx context.Context, request *DeleteApikeyRequest) string {
	if request == nil {
		request = NewDeleteApikeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteApikey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteApikeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteApikeyWithContextV2(ctx context.Context, request *DeleteApikeyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteApikeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteApikey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteApikeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModelsRequest() (request *DescribeModelsRequest) {
	request = &DescribeModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeModels")
	return
}

func NewDescribeModelsResponse() (response *DescribeModelsResponse) {
	response = &DescribeModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModels(request *DescribeModelsRequest) string {
	return c.DescribeModelsWithContext(context.Background(), request)
}

func (c *Client) DescribeModelsSend(request *DescribeModelsRequest) (*DescribeModelsResponse, error) {
	statusCode, msg, err := c.DescribeModelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeModelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModelsWithContext(ctx context.Context, request *DescribeModelsRequest) string {
	if request == nil {
		request = NewDescribeModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModelsWithContextV2(ctx context.Context, request *DescribeModelsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateApikeyRequest() (request *CreateApikeyRequest) {
	request = &CreateApikeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateApikey")
	return
}

func NewCreateApikeyResponse() (response *CreateApikeyResponse) {
	response = &CreateApikeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateApikey(request *CreateApikeyRequest) string {
	return c.CreateApikeyWithContext(context.Background(), request)
}

func (c *Client) CreateApikeySend(request *CreateApikeyRequest) (*CreateApikeyResponse, error) {
	statusCode, msg, err := c.CreateApikeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateApikeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateApikeyWithContext(ctx context.Context, request *CreateApikeyRequest) string {
	if request == nil {
		request = NewCreateApikeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateApikey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateApikeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateApikeyWithContextV2(ctx context.Context, request *CreateApikeyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateApikeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateApikey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateApikeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetModelDetailRequest() (request *GetModelDetailRequest) {
	request = &GetModelDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetModelDetail")
	return
}

func NewGetModelDetailResponse() (response *GetModelDetailResponse) {
	response = &GetModelDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetModelDetail(request *GetModelDetailRequest) string {
	return c.GetModelDetailWithContext(context.Background(), request)
}

func (c *Client) GetModelDetailSend(request *GetModelDetailRequest) (*GetModelDetailResponse, error) {
	statusCode, msg, err := c.GetModelDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetModelDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetModelDetailWithContext(ctx context.Context, request *GetModelDetailRequest) string {
	if request == nil {
		request = NewGetModelDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetModelDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetModelDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetModelDetailWithContextV2(ctx context.Context, request *GetModelDetailRequest) (int, string, error) {
	if request == nil {
		request = NewGetModelDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetModelDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetModelDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeApikeysRequest() (request *DescribeApikeysRequest) {
	request = &DescribeApikeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeApikeys")
	return
}

func NewDescribeApikeysResponse() (response *DescribeApikeysResponse) {
	response = &DescribeApikeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeApikeys(request *DescribeApikeysRequest) string {
	return c.DescribeApikeysWithContext(context.Background(), request)
}

func (c *Client) DescribeApikeysSend(request *DescribeApikeysRequest) (*DescribeApikeysResponse, error) {
	statusCode, msg, err := c.DescribeApikeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeApikeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeApikeysWithContext(ctx context.Context, request *DescribeApikeysRequest) string {
	if request == nil {
		request = NewDescribeApikeysRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeApikeys")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeApikeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeApikeysWithContextV2(ctx context.Context, request *DescribeApikeysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeApikeysRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeApikeys")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeApikeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableApikeyStatusRequest() (request *DisableApikeyStatusRequest) {
	request = &DisableApikeyStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DisableApikeyStatus")
	return
}

func NewDisableApikeyStatusResponse() (response *DisableApikeyStatusResponse) {
	response = &DisableApikeyStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableApikeyStatus(request *DisableApikeyStatusRequest) string {
	return c.DisableApikeyStatusWithContext(context.Background(), request)
}

func (c *Client) DisableApikeyStatusSend(request *DisableApikeyStatusRequest) (*DisableApikeyStatusResponse, error) {
	statusCode, msg, err := c.DisableApikeyStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisableApikeyStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableApikeyStatusWithContext(ctx context.Context, request *DisableApikeyStatusRequest) string {
	if request == nil {
		request = NewDisableApikeyStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableApikeyStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableApikeyStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableApikeyStatusWithContextV2(ctx context.Context, request *DisableApikeyStatusRequest) (int, string, error) {
	if request == nil {
		request = NewDisableApikeyStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableApikeyStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableApikeyStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetApiServiceRequest() (request *GetApiServiceRequest) {
	request = &GetApiServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetApiService")
	return
}

func NewGetApiServiceResponse() (response *GetApiServiceResponse) {
	response = &GetApiServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetApiService(request *GetApiServiceRequest) string {
	return c.GetApiServiceWithContext(context.Background(), request)
}

func (c *Client) GetApiServiceSend(request *GetApiServiceRequest) (*GetApiServiceResponse, error) {
	statusCode, msg, err := c.GetApiServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetApiServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetApiServiceWithContext(ctx context.Context, request *GetApiServiceRequest) string {
	if request == nil {
		request = NewGetApiServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetApiService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetApiServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetApiServiceWithContextV2(ctx context.Context, request *GetApiServiceRequest) (int, string, error) {
	if request == nil {
		request = NewGetApiServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetApiService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetApiServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBatchInferenceJobsFinalResultDownloadUrlRequest() (request *GetBatchInferenceJobsFinalResultDownloadUrlRequest) {
	request = &GetBatchInferenceJobsFinalResultDownloadUrlRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetBatchInferenceJobsFinalResultDownloadUrl")
	return
}

func NewGetBatchInferenceJobsFinalResultDownloadUrlResponse() (response *GetBatchInferenceJobsFinalResultDownloadUrlResponse) {
	response = &GetBatchInferenceJobsFinalResultDownloadUrlResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBatchInferenceJobsFinalResultDownloadUrl(request *GetBatchInferenceJobsFinalResultDownloadUrlRequest) string {
	return c.GetBatchInferenceJobsFinalResultDownloadUrlWithContext(context.Background(), request)
}

func (c *Client) GetBatchInferenceJobsFinalResultDownloadUrlSend(request *GetBatchInferenceJobsFinalResultDownloadUrlRequest) (*GetBatchInferenceJobsFinalResultDownloadUrlResponse, error) {
	statusCode, msg, err := c.GetBatchInferenceJobsFinalResultDownloadUrlWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetBatchInferenceJobsFinalResultDownloadUrlResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBatchInferenceJobsFinalResultDownloadUrlWithContext(ctx context.Context, request *GetBatchInferenceJobsFinalResultDownloadUrlRequest) string {
	if request == nil {
		request = NewGetBatchInferenceJobsFinalResultDownloadUrlRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetBatchInferenceJobsFinalResultDownloadUrl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBatchInferenceJobsFinalResultDownloadUrlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBatchInferenceJobsFinalResultDownloadUrlWithContextV2(ctx context.Context, request *GetBatchInferenceJobsFinalResultDownloadUrlRequest) (int, string, error) {
	if request == nil {
		request = NewGetBatchInferenceJobsFinalResultDownloadUrlRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetBatchInferenceJobsFinalResultDownloadUrl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBatchInferenceJobsFinalResultDownloadUrlResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInferenceJobsKs3AuthInfoRequest() (request *DescribeInferenceJobsKs3AuthInfoRequest) {
	request = &DescribeInferenceJobsKs3AuthInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeInferenceJobsKs3AuthInfo")
	return
}

func NewDescribeInferenceJobsKs3AuthInfoResponse() (response *DescribeInferenceJobsKs3AuthInfoResponse) {
	response = &DescribeInferenceJobsKs3AuthInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInferenceJobsKs3AuthInfo(request *DescribeInferenceJobsKs3AuthInfoRequest) string {
	return c.DescribeInferenceJobsKs3AuthInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeInferenceJobsKs3AuthInfoSend(request *DescribeInferenceJobsKs3AuthInfoRequest) (*DescribeInferenceJobsKs3AuthInfoResponse, error) {
	statusCode, msg, err := c.DescribeInferenceJobsKs3AuthInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInferenceJobsKs3AuthInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInferenceJobsKs3AuthInfoWithContext(ctx context.Context, request *DescribeInferenceJobsKs3AuthInfoRequest) string {
	if request == nil {
		request = NewDescribeInferenceJobsKs3AuthInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeInferenceJobsKs3AuthInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInferenceJobsKs3AuthInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInferenceJobsKs3AuthInfoWithContextV2(ctx context.Context, request *DescribeInferenceJobsKs3AuthInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInferenceJobsKs3AuthInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeInferenceJobsKs3AuthInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInferenceJobsKs3AuthInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopBatchInferenceJobRequest() (request *StopBatchInferenceJobRequest) {
	request = &StopBatchInferenceJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopBatchInferenceJob")
	return
}

func NewStopBatchInferenceJobResponse() (response *StopBatchInferenceJobResponse) {
	response = &StopBatchInferenceJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopBatchInferenceJob(request *StopBatchInferenceJobRequest) string {
	return c.StopBatchInferenceJobWithContext(context.Background(), request)
}

func (c *Client) StopBatchInferenceJobSend(request *StopBatchInferenceJobRequest) (*StopBatchInferenceJobResponse, error) {
	statusCode, msg, err := c.StopBatchInferenceJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopBatchInferenceJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopBatchInferenceJobWithContext(ctx context.Context, request *StopBatchInferenceJobRequest) string {
	if request == nil {
		request = NewStopBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopBatchInferenceJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopBatchInferenceJobWithContextV2(ctx context.Context, request *StopBatchInferenceJobRequest) (int, string, error) {
	if request == nil {
		request = NewStopBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopBatchInferenceJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateBatchInferenceJobRequest() (request *CreateBatchInferenceJobRequest) {
	request = &CreateBatchInferenceJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateBatchInferenceJob")
	return
}

func NewCreateBatchInferenceJobResponse() (response *CreateBatchInferenceJobResponse) {
	response = &CreateBatchInferenceJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateBatchInferenceJob(request *CreateBatchInferenceJobRequest) string {
	return c.CreateBatchInferenceJobWithContext(context.Background(), request)
}

func (c *Client) CreateBatchInferenceJobSend(request *CreateBatchInferenceJobRequest) (*CreateBatchInferenceJobResponse, error) {
	statusCode, msg, err := c.CreateBatchInferenceJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateBatchInferenceJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateBatchInferenceJobWithContext(ctx context.Context, request *CreateBatchInferenceJobRequest) string {
	if request == nil {
		request = NewCreateBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateBatchInferenceJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateBatchInferenceJobWithContextV2(ctx context.Context, request *CreateBatchInferenceJobRequest) (int, string, error) {
	if request == nil {
		request = NewCreateBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateBatchInferenceJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyBatchInferenceJobRequest() (request *ModifyBatchInferenceJobRequest) {
	request = &ModifyBatchInferenceJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyBatchInferenceJob")
	return
}

func NewModifyBatchInferenceJobResponse() (response *ModifyBatchInferenceJobResponse) {
	response = &ModifyBatchInferenceJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBatchInferenceJob(request *ModifyBatchInferenceJobRequest) string {
	return c.ModifyBatchInferenceJobWithContext(context.Background(), request)
}

func (c *Client) ModifyBatchInferenceJobSend(request *ModifyBatchInferenceJobRequest) (*ModifyBatchInferenceJobResponse, error) {
	statusCode, msg, err := c.ModifyBatchInferenceJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyBatchInferenceJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyBatchInferenceJobWithContext(ctx context.Context, request *ModifyBatchInferenceJobRequest) string {
	if request == nil {
		request = NewModifyBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyBatchInferenceJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyBatchInferenceJobWithContextV2(ctx context.Context, request *ModifyBatchInferenceJobRequest) (int, string, error) {
	if request == nil {
		request = NewModifyBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyBatchInferenceJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBatchInferenceJobsRequest() (request *DescribeBatchInferenceJobsRequest) {
	request = &DescribeBatchInferenceJobsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeBatchInferenceJobs")
	return
}

func NewDescribeBatchInferenceJobsResponse() (response *DescribeBatchInferenceJobsResponse) {
	response = &DescribeBatchInferenceJobsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBatchInferenceJobs(request *DescribeBatchInferenceJobsRequest) string {
	return c.DescribeBatchInferenceJobsWithContext(context.Background(), request)
}

func (c *Client) DescribeBatchInferenceJobsSend(request *DescribeBatchInferenceJobsRequest) (*DescribeBatchInferenceJobsResponse, error) {
	statusCode, msg, err := c.DescribeBatchInferenceJobsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBatchInferenceJobsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBatchInferenceJobsWithContext(ctx context.Context, request *DescribeBatchInferenceJobsRequest) string {
	if request == nil {
		request = NewDescribeBatchInferenceJobsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeBatchInferenceJobs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBatchInferenceJobsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBatchInferenceJobsWithContextV2(ctx context.Context, request *DescribeBatchInferenceJobsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBatchInferenceJobsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeBatchInferenceJobs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBatchInferenceJobsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBatchInferenceJobRequest() (request *DeleteBatchInferenceJobRequest) {
	request = &DeleteBatchInferenceJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteBatchInferenceJob")
	return
}

func NewDeleteBatchInferenceJobResponse() (response *DeleteBatchInferenceJobResponse) {
	response = &DeleteBatchInferenceJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBatchInferenceJob(request *DeleteBatchInferenceJobRequest) string {
	return c.DeleteBatchInferenceJobWithContext(context.Background(), request)
}

func (c *Client) DeleteBatchInferenceJobSend(request *DeleteBatchInferenceJobRequest) (*DeleteBatchInferenceJobResponse, error) {
	statusCode, msg, err := c.DeleteBatchInferenceJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteBatchInferenceJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBatchInferenceJobWithContext(ctx context.Context, request *DeleteBatchInferenceJobRequest) string {
	if request == nil {
		request = NewDeleteBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchInferenceJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBatchInferenceJobWithContextV2(ctx context.Context, request *DeleteBatchInferenceJobRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBatchInferenceJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteBatchInferenceJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchInferenceJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEnableModelsRequest() (request *EnableModelsRequest) {
	request = &EnableModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "EnableModels")
	return
}

func NewEnableModelsResponse() (response *EnableModelsResponse) {
	response = &EnableModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableModels(request *EnableModelsRequest) string {
	return c.EnableModelsWithContext(context.Background(), request)
}

func (c *Client) EnableModelsSend(request *EnableModelsRequest) (*EnableModelsResponse, error) {
	statusCode, msg, err := c.EnableModelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct EnableModelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EnableModelsWithContext(ctx context.Context, request *EnableModelsRequest) string {
	if request == nil {
		request = NewEnableModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "EnableModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EnableModelsWithContextV2(ctx context.Context, request *EnableModelsRequest) (int, string, error) {
	if request == nil {
		request = NewEnableModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "EnableModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableModelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModelQuotasRequest() (request *DescribeModelQuotasRequest) {
	request = &DescribeModelQuotasRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeModelQuotas")
	return
}

func NewDescribeModelQuotasResponse() (response *DescribeModelQuotasResponse) {
	response = &DescribeModelQuotasResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModelQuotas(request *DescribeModelQuotasRequest) string {
	return c.DescribeModelQuotasWithContext(context.Background(), request)
}

func (c *Client) DescribeModelQuotasSend(request *DescribeModelQuotasRequest) (*DescribeModelQuotasResponse, error) {
	statusCode, msg, err := c.DescribeModelQuotasWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeModelQuotasResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModelQuotasWithContext(ctx context.Context, request *DescribeModelQuotasRequest) string {
	if request == nil {
		request = NewDescribeModelQuotasRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeModelQuotas")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelQuotasResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModelQuotasWithContextV2(ctx context.Context, request *DescribeModelQuotasRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModelQuotasRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeModelQuotas")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelQuotasResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableModelsRequest() (request *DisableModelsRequest) {
	request = &DisableModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DisableModels")
	return
}

func NewDisableModelsResponse() (response *DisableModelsResponse) {
	response = &DisableModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableModels(request *DisableModelsRequest) string {
	return c.DisableModelsWithContext(context.Background(), request)
}

func (c *Client) DisableModelsSend(request *DisableModelsRequest) (*DisableModelsResponse, error) {
	statusCode, msg, err := c.DisableModelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisableModelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableModelsWithContext(ctx context.Context, request *DisableModelsRequest) string {
	if request == nil {
		request = NewDisableModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableModelsWithContextV2(ctx context.Context, request *DisableModelsRequest) (int, string, error) {
	if request == nil {
		request = NewDisableModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableModelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEnableOverFreeLimitRequest() (request *EnableOverFreeLimitRequest) {
	request = &EnableOverFreeLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "EnableOverFreeLimit")
	return
}

func NewEnableOverFreeLimitResponse() (response *EnableOverFreeLimitResponse) {
	response = &EnableOverFreeLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableOverFreeLimit(request *EnableOverFreeLimitRequest) string {
	return c.EnableOverFreeLimitWithContext(context.Background(), request)
}

func (c *Client) EnableOverFreeLimitSend(request *EnableOverFreeLimitRequest) (*EnableOverFreeLimitResponse, error) {
	statusCode, msg, err := c.EnableOverFreeLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct EnableOverFreeLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EnableOverFreeLimitWithContext(ctx context.Context, request *EnableOverFreeLimitRequest) string {
	if request == nil {
		request = NewEnableOverFreeLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "EnableOverFreeLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableOverFreeLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EnableOverFreeLimitWithContextV2(ctx context.Context, request *EnableOverFreeLimitRequest) (int, string, error) {
	if request == nil {
		request = NewEnableOverFreeLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "EnableOverFreeLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableOverFreeLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableOverFreeLimitRequest() (request *DisableOverFreeLimitRequest) {
	request = &DisableOverFreeLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DisableOverFreeLimit")
	return
}

func NewDisableOverFreeLimitResponse() (response *DisableOverFreeLimitResponse) {
	response = &DisableOverFreeLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableOverFreeLimit(request *DisableOverFreeLimitRequest) string {
	return c.DisableOverFreeLimitWithContext(context.Background(), request)
}

func (c *Client) DisableOverFreeLimitSend(request *DisableOverFreeLimitRequest) (*DisableOverFreeLimitResponse, error) {
	statusCode, msg, err := c.DisableOverFreeLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisableOverFreeLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableOverFreeLimitWithContext(ctx context.Context, request *DisableOverFreeLimitRequest) string {
	if request == nil {
		request = NewDisableOverFreeLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableOverFreeLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableOverFreeLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableOverFreeLimitWithContextV2(ctx context.Context, request *DisableOverFreeLimitRequest) (int, string, error) {
	if request == nil {
		request = NewDisableOverFreeLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableOverFreeLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableOverFreeLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateTrainJobRequest() (request *CreateTrainJobRequest) {
	request = &CreateTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateTrainJob")
	return
}

func NewCreateTrainJobResponse() (response *CreateTrainJobResponse) {
	response = &CreateTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTrainJob(request *CreateTrainJobRequest) string {
	return c.CreateTrainJobWithContext(context.Background(), request)
}

func (c *Client) CreateTrainJobSend(request *CreateTrainJobRequest) (*CreateTrainJobResponse, error) {
	statusCode, msg, err := c.CreateTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTrainJobWithContext(ctx context.Context, request *CreateTrainJobRequest) string {
	if request == nil {
		request = NewCreateTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTrainJobWithContextV2(ctx context.Context, request *CreateTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTrainJobEventsRequest() (request *DescribeTrainJobEventsRequest) {
	request = &DescribeTrainJobEventsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobEvents")
	return
}

func NewDescribeTrainJobEventsResponse() (response *DescribeTrainJobEventsResponse) {
	response = &DescribeTrainJobEventsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrainJobEvents(request *DescribeTrainJobEventsRequest) string {
	return c.DescribeTrainJobEventsWithContext(context.Background(), request)
}

func (c *Client) DescribeTrainJobEventsSend(request *DescribeTrainJobEventsRequest) (*DescribeTrainJobEventsResponse, error) {
	statusCode, msg, err := c.DescribeTrainJobEventsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrainJobEventsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrainJobEventsWithContext(ctx context.Context, request *DescribeTrainJobEventsRequest) string {
	if request == nil {
		request = NewDescribeTrainJobEventsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobEvents")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrainJobEventsWithContextV2(ctx context.Context, request *DescribeTrainJobEventsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrainJobEventsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobEvents")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobEventsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopTrainJobRequest() (request *StopTrainJobRequest) {
	request = &StopTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopTrainJob")
	return
}

func NewStopTrainJobResponse() (response *StopTrainJobResponse) {
	response = &StopTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopTrainJob(request *StopTrainJobRequest) string {
	return c.StopTrainJobWithContext(context.Background(), request)
}

func (c *Client) StopTrainJobSend(request *StopTrainJobRequest) (*StopTrainJobResponse, error) {
	statusCode, msg, err := c.StopTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopTrainJobWithContext(ctx context.Context, request *StopTrainJobRequest) string {
	if request == nil {
		request = NewStopTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopTrainJobWithContextV2(ctx context.Context, request *StopTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewStopTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTrainJobRequest() (request *DescribeTrainJobRequest) {
	request = &DescribeTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJob")
	return
}

func NewDescribeTrainJobResponse() (response *DescribeTrainJobResponse) {
	response = &DescribeTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrainJob(request *DescribeTrainJobRequest) string {
	return c.DescribeTrainJobWithContext(context.Background(), request)
}

func (c *Client) DescribeTrainJobSend(request *DescribeTrainJobRequest) (*DescribeTrainJobResponse, error) {
	statusCode, msg, err := c.DescribeTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrainJobWithContext(ctx context.Context, request *DescribeTrainJobRequest) string {
	if request == nil {
		request = NewDescribeTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrainJobWithContextV2(ctx context.Context, request *DescribeTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartTrainJobRequest() (request *StartTrainJobRequest) {
	request = &StartTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StartTrainJob")
	return
}

func NewStartTrainJobResponse() (response *StartTrainJobResponse) {
	response = &StartTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartTrainJob(request *StartTrainJobRequest) string {
	return c.StartTrainJobWithContext(context.Background(), request)
}

func (c *Client) StartTrainJobSend(request *StartTrainJobRequest) (*StartTrainJobResponse, error) {
	statusCode, msg, err := c.StartTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartTrainJobWithContext(ctx context.Context, request *StartTrainJobRequest) string {
	if request == nil {
		request = NewStartTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartTrainJobWithContextV2(ctx context.Context, request *StartTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewStartTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteTrainJobRequest() (request *DeleteTrainJobRequest) {
	request = &DeleteTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteTrainJob")
	return
}

func NewDeleteTrainJobResponse() (response *DeleteTrainJobResponse) {
	response = &DeleteTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTrainJob(request *DeleteTrainJobRequest) string {
	return c.DeleteTrainJobWithContext(context.Background(), request)
}

func (c *Client) DeleteTrainJobSend(request *DeleteTrainJobRequest) (*DeleteTrainJobResponse, error) {
	statusCode, msg, err := c.DeleteTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteTrainJobWithContext(ctx context.Context, request *DeleteTrainJobRequest) string {
	if request == nil {
		request = NewDeleteTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteTrainJobWithContextV2(ctx context.Context, request *DeleteTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyTrainJobRequest() (request *ModifyTrainJobRequest) {
	request = &ModifyTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyTrainJob")
	return
}

func NewModifyTrainJobResponse() (response *ModifyTrainJobResponse) {
	response = &ModifyTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyTrainJob(request *ModifyTrainJobRequest) string {
	return c.ModifyTrainJobWithContext(context.Background(), request)
}

func (c *Client) ModifyTrainJobSend(request *ModifyTrainJobRequest) (*ModifyTrainJobResponse, error) {
	statusCode, msg, err := c.ModifyTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyTrainJobWithContext(ctx context.Context, request *ModifyTrainJobRequest) string {
	if request == nil {
		request = NewModifyTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyTrainJobWithContextV2(ctx context.Context, request *ModifyTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewModifyTrainJobRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyTrainJob")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTrainJobPodLogsRequest() (request *DescribeTrainJobPodLogsRequest) {
	request = &DescribeTrainJobPodLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobPodLogs")
	return
}

func NewDescribeTrainJobPodLogsResponse() (response *DescribeTrainJobPodLogsResponse) {
	response = &DescribeTrainJobPodLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrainJobPodLogs(request *DescribeTrainJobPodLogsRequest) string {
	return c.DescribeTrainJobPodLogsWithContext(context.Background(), request)
}

func (c *Client) DescribeTrainJobPodLogsSend(request *DescribeTrainJobPodLogsRequest) (*DescribeTrainJobPodLogsResponse, error) {
	statusCode, msg, err := c.DescribeTrainJobPodLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrainJobPodLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrainJobPodLogsWithContext(ctx context.Context, request *DescribeTrainJobPodLogsRequest) string {
	if request == nil {
		request = NewDescribeTrainJobPodLogsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobPodLogs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobPodLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrainJobPodLogsWithContextV2(ctx context.Context, request *DescribeTrainJobPodLogsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrainJobPodLogsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobPodLogs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobPodLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTrainJobPodsRequest() (request *DescribeTrainJobPodsRequest) {
	request = &DescribeTrainJobPodsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobPods")
	return
}

func NewDescribeTrainJobPodsResponse() (response *DescribeTrainJobPodsResponse) {
	response = &DescribeTrainJobPodsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrainJobPods(request *DescribeTrainJobPodsRequest) string {
	return c.DescribeTrainJobPodsWithContext(context.Background(), request)
}

func (c *Client) DescribeTrainJobPodsSend(request *DescribeTrainJobPodsRequest) (*DescribeTrainJobPodsResponse, error) {
	statusCode, msg, err := c.DescribeTrainJobPodsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrainJobPodsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrainJobPodsWithContext(ctx context.Context, request *DescribeTrainJobPodsRequest) string {
	if request == nil {
		request = NewDescribeTrainJobPodsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobPods")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobPodsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrainJobPodsWithContextV2(ctx context.Context, request *DescribeTrainJobPodsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrainJobPodsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobPods")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobPodsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeResourcePoolsRequest() (request *DescribeResourcePoolsRequest) {
	request = &DescribeResourcePoolsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePools")
	return
}

func NewDescribeResourcePoolsResponse() (response *DescribeResourcePoolsResponse) {
	response = &DescribeResourcePoolsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeResourcePools(request *DescribeResourcePoolsRequest) string {
	return c.DescribeResourcePoolsWithContext(context.Background(), request)
}

func (c *Client) DescribeResourcePoolsSend(request *DescribeResourcePoolsRequest) (*DescribeResourcePoolsResponse, error) {
	statusCode, msg, err := c.DescribeResourcePoolsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeResourcePoolsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeResourcePoolsWithContext(ctx context.Context, request *DescribeResourcePoolsRequest) string {
	if request == nil {
		request = NewDescribeResourcePoolsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePools")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeResourcePoolsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeResourcePoolsWithContextV2(ctx context.Context, request *DescribeResourcePoolsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeResourcePoolsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePools")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeResourcePoolsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeResourcePoolInstancesRequest() (request *DescribeResourcePoolInstancesRequest) {
	request = &DescribeResourcePoolInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePoolInstances")
	return
}

func NewDescribeResourcePoolInstancesResponse() (response *DescribeResourcePoolInstancesResponse) {
	response = &DescribeResourcePoolInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeResourcePoolInstances(request *DescribeResourcePoolInstancesRequest) string {
	return c.DescribeResourcePoolInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeResourcePoolInstancesSend(request *DescribeResourcePoolInstancesRequest) (*DescribeResourcePoolInstancesResponse, error) {
	statusCode, msg, err := c.DescribeResourcePoolInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeResourcePoolInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeResourcePoolInstancesWithContext(ctx context.Context, request *DescribeResourcePoolInstancesRequest) string {
	if request == nil {
		request = NewDescribeResourcePoolInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePoolInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeResourcePoolInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeResourcePoolInstancesWithContextV2(ctx context.Context, request *DescribeResourcePoolInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeResourcePoolInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePoolInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeResourcePoolInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInferenceEndpointRequest() (request *CreateInferenceEndpointRequest) {
	request = &CreateInferenceEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateInferenceEndpoint")
	return
}

func NewCreateInferenceEndpointResponse() (response *CreateInferenceEndpointResponse) {
	response = &CreateInferenceEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInferenceEndpoint(request *CreateInferenceEndpointRequest) string {
	return c.CreateInferenceEndpointWithContext(context.Background(), request)
}

func (c *Client) CreateInferenceEndpointSend(request *CreateInferenceEndpointRequest) (*CreateInferenceEndpointResponse, error) {
	statusCode, msg, err := c.CreateInferenceEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInferenceEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInferenceEndpointWithContext(ctx context.Context, request *CreateInferenceEndpointRequest) string {
	if request == nil {
		request = NewCreateInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInferenceEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInferenceEndpointWithContextV2(ctx context.Context, request *CreateInferenceEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInferenceEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInferenceEndpointsRequest() (request *DescribeInferenceEndpointsRequest) {
	request = &DescribeInferenceEndpointsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeInferenceEndpoints")
	return
}

func NewDescribeInferenceEndpointsResponse() (response *DescribeInferenceEndpointsResponse) {
	response = &DescribeInferenceEndpointsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInferenceEndpoints(request *DescribeInferenceEndpointsRequest) string {
	return c.DescribeInferenceEndpointsWithContext(context.Background(), request)
}

func (c *Client) DescribeInferenceEndpointsSend(request *DescribeInferenceEndpointsRequest) (*DescribeInferenceEndpointsResponse, error) {
	statusCode, msg, err := c.DescribeInferenceEndpointsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInferenceEndpointsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInferenceEndpointsWithContext(ctx context.Context, request *DescribeInferenceEndpointsRequest) string {
	if request == nil {
		request = NewDescribeInferenceEndpointsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeInferenceEndpoints")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInferenceEndpointsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInferenceEndpointsWithContextV2(ctx context.Context, request *DescribeInferenceEndpointsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInferenceEndpointsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeInferenceEndpoints")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInferenceEndpointsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEnableEndpointRateLimitRequest() (request *EnableEndpointRateLimitRequest) {
	request = &EnableEndpointRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "EnableEndpointRateLimit")
	return
}

func NewEnableEndpointRateLimitResponse() (response *EnableEndpointRateLimitResponse) {
	response = &EnableEndpointRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableEndpointRateLimit(request *EnableEndpointRateLimitRequest) string {
	return c.EnableEndpointRateLimitWithContext(context.Background(), request)
}

func (c *Client) EnableEndpointRateLimitSend(request *EnableEndpointRateLimitRequest) (*EnableEndpointRateLimitResponse, error) {
	statusCode, msg, err := c.EnableEndpointRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct EnableEndpointRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EnableEndpointRateLimitWithContext(ctx context.Context, request *EnableEndpointRateLimitRequest) string {
	if request == nil {
		request = NewEnableEndpointRateLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "EnableEndpointRateLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableEndpointRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EnableEndpointRateLimitWithContextV2(ctx context.Context, request *EnableEndpointRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewEnableEndpointRateLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "EnableEndpointRateLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableEndpointRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateInferenceEndpointRequest() (request *UpdateInferenceEndpointRequest) {
	request = &UpdateInferenceEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "UpdateInferenceEndpoint")
	return
}

func NewUpdateInferenceEndpointResponse() (response *UpdateInferenceEndpointResponse) {
	response = &UpdateInferenceEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateInferenceEndpoint(request *UpdateInferenceEndpointRequest) string {
	return c.UpdateInferenceEndpointWithContext(context.Background(), request)
}

func (c *Client) UpdateInferenceEndpointSend(request *UpdateInferenceEndpointRequest) (*UpdateInferenceEndpointResponse, error) {
	statusCode, msg, err := c.UpdateInferenceEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateInferenceEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateInferenceEndpointWithContext(ctx context.Context, request *UpdateInferenceEndpointRequest) string {
	if request == nil {
		request = NewUpdateInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateInferenceEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateInferenceEndpointWithContextV2(ctx context.Context, request *UpdateInferenceEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateInferenceEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartInferenceEndpointRequest() (request *StartInferenceEndpointRequest) {
	request = &StartInferenceEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StartInferenceEndpoint")
	return
}

func NewStartInferenceEndpointResponse() (response *StartInferenceEndpointResponse) {
	response = &StartInferenceEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartInferenceEndpoint(request *StartInferenceEndpointRequest) string {
	return c.StartInferenceEndpointWithContext(context.Background(), request)
}

func (c *Client) StartInferenceEndpointSend(request *StartInferenceEndpointRequest) (*StartInferenceEndpointResponse, error) {
	statusCode, msg, err := c.StartInferenceEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartInferenceEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartInferenceEndpointWithContext(ctx context.Context, request *StartInferenceEndpointRequest) string {
	if request == nil {
		request = NewStartInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartInferenceEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartInferenceEndpointWithContextV2(ctx context.Context, request *StartInferenceEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewStartInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartInferenceEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopInferenceEndpointRequest() (request *StopInferenceEndpointRequest) {
	request = &StopInferenceEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopInferenceEndpoint")
	return
}

func NewStopInferenceEndpointResponse() (response *StopInferenceEndpointResponse) {
	response = &StopInferenceEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopInferenceEndpoint(request *StopInferenceEndpointRequest) string {
	return c.StopInferenceEndpointWithContext(context.Background(), request)
}

func (c *Client) StopInferenceEndpointSend(request *StopInferenceEndpointRequest) (*StopInferenceEndpointResponse, error) {
	statusCode, msg, err := c.StopInferenceEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopInferenceEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopInferenceEndpointWithContext(ctx context.Context, request *StopInferenceEndpointRequest) string {
	if request == nil {
		request = NewStopInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopInferenceEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopInferenceEndpointWithContextV2(ctx context.Context, request *StopInferenceEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewStopInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StopInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopInferenceEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInferenceEndpointRequest() (request *DeleteInferenceEndpointRequest) {
	request = &DeleteInferenceEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteInferenceEndpoint")
	return
}

func NewDeleteInferenceEndpointResponse() (response *DeleteInferenceEndpointResponse) {
	response = &DeleteInferenceEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInferenceEndpoint(request *DeleteInferenceEndpointRequest) string {
	return c.DeleteInferenceEndpointWithContext(context.Background(), request)
}

func (c *Client) DeleteInferenceEndpointSend(request *DeleteInferenceEndpointRequest) (*DeleteInferenceEndpointResponse, error) {
	statusCode, msg, err := c.DeleteInferenceEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInferenceEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInferenceEndpointWithContext(ctx context.Context, request *DeleteInferenceEndpointRequest) string {
	if request == nil {
		request = NewDeleteInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInferenceEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInferenceEndpointWithContextV2(ctx context.Context, request *DeleteInferenceEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInferenceEndpointRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteInferenceEndpoint")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInferenceEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableEndpointRateLimitRequest() (request *DisableEndpointRateLimitRequest) {
	request = &DisableEndpointRateLimitRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DisableEndpointRateLimit")
	return
}

func NewDisableEndpointRateLimitResponse() (response *DisableEndpointRateLimitResponse) {
	response = &DisableEndpointRateLimitResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableEndpointRateLimit(request *DisableEndpointRateLimitRequest) string {
	return c.DisableEndpointRateLimitWithContext(context.Background(), request)
}

func (c *Client) DisableEndpointRateLimitSend(request *DisableEndpointRateLimitRequest) (*DisableEndpointRateLimitResponse, error) {
	statusCode, msg, err := c.DisableEndpointRateLimitWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisableEndpointRateLimitResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableEndpointRateLimitWithContext(ctx context.Context, request *DisableEndpointRateLimitRequest) string {
	if request == nil {
		request = NewDisableEndpointRateLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableEndpointRateLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableEndpointRateLimitResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableEndpointRateLimitWithContextV2(ctx context.Context, request *DisableEndpointRateLimitRequest) (int, string, error) {
	if request == nil {
		request = NewDisableEndpointRateLimitRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisableEndpointRateLimit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableEndpointRateLimitResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeResourcePoolInstanceTasksRequest() (request *DescribeResourcePoolInstanceTasksRequest) {
	request = &DescribeResourcePoolInstanceTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePoolInstanceTasks")
	return
}

func NewDescribeResourcePoolInstanceTasksResponse() (response *DescribeResourcePoolInstanceTasksResponse) {
	response = &DescribeResourcePoolInstanceTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeResourcePoolInstanceTasks(request *DescribeResourcePoolInstanceTasksRequest) string {
	return c.DescribeResourcePoolInstanceTasksWithContext(context.Background(), request)
}

func (c *Client) DescribeResourcePoolInstanceTasksSend(request *DescribeResourcePoolInstanceTasksRequest) (*DescribeResourcePoolInstanceTasksResponse, error) {
	statusCode, msg, err := c.DescribeResourcePoolInstanceTasksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeResourcePoolInstanceTasksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeResourcePoolInstanceTasksWithContext(ctx context.Context, request *DescribeResourcePoolInstanceTasksRequest) string {
	if request == nil {
		request = NewDescribeResourcePoolInstanceTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePoolInstanceTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeResourcePoolInstanceTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeResourcePoolInstanceTasksWithContextV2(ctx context.Context, request *DescribeResourcePoolInstanceTasksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeResourcePoolInstanceTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeResourcePoolInstanceTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeResourcePoolInstanceTasksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetKcrPersonalTokenRequest() (request *SetKcrPersonalTokenRequest) {
	request = &SetKcrPersonalTokenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "SetKcrPersonalToken")
	return
}

func NewSetKcrPersonalTokenResponse() (response *SetKcrPersonalTokenResponse) {
	response = &SetKcrPersonalTokenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetKcrPersonalToken(request *SetKcrPersonalTokenRequest) string {
	return c.SetKcrPersonalTokenWithContext(context.Background(), request)
}

func (c *Client) SetKcrPersonalTokenSend(request *SetKcrPersonalTokenRequest) (*SetKcrPersonalTokenResponse, error) {
	statusCode, msg, err := c.SetKcrPersonalTokenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetKcrPersonalTokenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetKcrPersonalTokenWithContext(ctx context.Context, request *SetKcrPersonalTokenRequest) string {
	if request == nil {
		request = NewSetKcrPersonalTokenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "SetKcrPersonalToken")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetKcrPersonalTokenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetKcrPersonalTokenWithContextV2(ctx context.Context, request *SetKcrPersonalTokenRequest) (int, string, error) {
	if request == nil {
		request = NewSetKcrPersonalTokenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "SetKcrPersonalToken")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetKcrPersonalTokenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeQueuesRequest() (request *DescribeQueuesRequest) {
	request = &DescribeQueuesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeQueues")
	return
}

func NewDescribeQueuesResponse() (response *DescribeQueuesResponse) {
	response = &DescribeQueuesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeQueues(request *DescribeQueuesRequest) string {
	return c.DescribeQueuesWithContext(context.Background(), request)
}

func (c *Client) DescribeQueuesSend(request *DescribeQueuesRequest) (*DescribeQueuesResponse, error) {
	statusCode, msg, err := c.DescribeQueuesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeQueuesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeQueuesWithContext(ctx context.Context, request *DescribeQueuesRequest) string {
	if request == nil {
		request = NewDescribeQueuesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeQueues")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeQueuesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeQueuesWithContextV2(ctx context.Context, request *DescribeQueuesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeQueuesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeQueues")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeQueuesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
