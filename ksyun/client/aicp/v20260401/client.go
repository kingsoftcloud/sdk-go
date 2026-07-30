package v20260401

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2026-04-01"

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

func NewCreateSandboxTemplateRequest() (request *CreateSandboxTemplateRequest) {
	request = &CreateSandboxTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateSandboxTemplate")
	return
}

func NewCreateSandboxTemplateResponse() (response *CreateSandboxTemplateResponse) {
	response = &CreateSandboxTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSandboxTemplate(request *CreateSandboxTemplateRequest) string {
	return c.CreateSandboxTemplateWithContext(context.Background(), request)
}

func (c *Client) CreateSandboxTemplateSend(request *CreateSandboxTemplateRequest) (*CreateSandboxTemplateResponse, error) {
	statusCode, msg, err := c.CreateSandboxTemplateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSandboxTemplateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSandboxTemplateWithContext(ctx context.Context, request *CreateSandboxTemplateRequest) string {
	if request == nil {
		request = NewCreateSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSandboxTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSandboxTemplateWithContextV2(ctx context.Context, request *CreateSandboxTemplateRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSandboxTemplateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateSandboxTemplateRequest() (request *UpdateSandboxTemplateRequest) {
	request = &UpdateSandboxTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "UpdateSandboxTemplate")
	return
}

func NewUpdateSandboxTemplateResponse() (response *UpdateSandboxTemplateResponse) {
	response = &UpdateSandboxTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateSandboxTemplate(request *UpdateSandboxTemplateRequest) string {
	return c.UpdateSandboxTemplateWithContext(context.Background(), request)
}

func (c *Client) UpdateSandboxTemplateSend(request *UpdateSandboxTemplateRequest) (*UpdateSandboxTemplateResponse, error) {
	statusCode, msg, err := c.UpdateSandboxTemplateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateSandboxTemplateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateSandboxTemplateWithContext(ctx context.Context, request *UpdateSandboxTemplateRequest) string {
	if request == nil {
		request = NewUpdateSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSandboxTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateSandboxTemplateWithContextV2(ctx context.Context, request *UpdateSandboxTemplateRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSandboxTemplateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSandboxInstanceRequest() (request *DeleteSandboxInstanceRequest) {
	request = &DeleteSandboxInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteSandboxInstance")
	return
}

func NewDeleteSandboxInstanceResponse() (response *DeleteSandboxInstanceResponse) {
	response = &DeleteSandboxInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSandboxInstance(request *DeleteSandboxInstanceRequest) string {
	return c.DeleteSandboxInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteSandboxInstanceSend(request *DeleteSandboxInstanceRequest) (*DeleteSandboxInstanceResponse, error) {
	statusCode, msg, err := c.DeleteSandboxInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSandboxInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSandboxInstanceWithContext(ctx context.Context, request *DeleteSandboxInstanceRequest) string {
	if request == nil {
		request = NewDeleteSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSandboxInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSandboxInstanceWithContextV2(ctx context.Context, request *DeleteSandboxInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSandboxInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetSandboxInstanceRequest() (request *GetSandboxInstanceRequest) {
	request = &GetSandboxInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxInstance")
	return
}

func NewGetSandboxInstanceResponse() (response *GetSandboxInstanceResponse) {
	response = &GetSandboxInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetSandboxInstance(request *GetSandboxInstanceRequest) string {
	return c.GetSandboxInstanceWithContext(context.Background(), request)
}

func (c *Client) GetSandboxInstanceSend(request *GetSandboxInstanceRequest) (*GetSandboxInstanceResponse, error) {
	statusCode, msg, err := c.GetSandboxInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetSandboxInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetSandboxInstanceWithContext(ctx context.Context, request *GetSandboxInstanceRequest) string {
	if request == nil {
		request = NewGetSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetSandboxInstanceWithContextV2(ctx context.Context, request *GetSandboxInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewGetSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetSandboxInstanceListRequest() (request *GetSandboxInstanceListRequest) {
	request = &GetSandboxInstanceListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxInstanceList")
	return
}

func NewGetSandboxInstanceListResponse() (response *GetSandboxInstanceListResponse) {
	response = &GetSandboxInstanceListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetSandboxInstanceList(request *GetSandboxInstanceListRequest) string {
	return c.GetSandboxInstanceListWithContext(context.Background(), request)
}

func (c *Client) GetSandboxInstanceListSend(request *GetSandboxInstanceListRequest) (*GetSandboxInstanceListResponse, error) {
	statusCode, msg, err := c.GetSandboxInstanceListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetSandboxInstanceListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetSandboxInstanceListWithContext(ctx context.Context, request *GetSandboxInstanceListRequest) string {
	if request == nil {
		request = NewGetSandboxInstanceListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxInstanceList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxInstanceListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetSandboxInstanceListWithContextV2(ctx context.Context, request *GetSandboxInstanceListRequest) (int, string, error) {
	if request == nil {
		request = NewGetSandboxInstanceListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxInstanceList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxInstanceListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetSandboxTemplateListRequest() (request *GetSandboxTemplateListRequest) {
	request = &GetSandboxTemplateListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxTemplateList")
	return
}

func NewGetSandboxTemplateListResponse() (response *GetSandboxTemplateListResponse) {
	response = &GetSandboxTemplateListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetSandboxTemplateList(request *GetSandboxTemplateListRequest) string {
	return c.GetSandboxTemplateListWithContext(context.Background(), request)
}

func (c *Client) GetSandboxTemplateListSend(request *GetSandboxTemplateListRequest) (*GetSandboxTemplateListResponse, error) {
	statusCode, msg, err := c.GetSandboxTemplateListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetSandboxTemplateListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetSandboxTemplateListWithContext(ctx context.Context, request *GetSandboxTemplateListRequest) string {
	if request == nil {
		request = NewGetSandboxTemplateListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxTemplateList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxTemplateListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetSandboxTemplateListWithContextV2(ctx context.Context, request *GetSandboxTemplateListRequest) (int, string, error) {
	if request == nil {
		request = NewGetSandboxTemplateListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxTemplateList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxTemplateListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartSandboxInstanceRequest() (request *StartSandboxInstanceRequest) {
	request = &StartSandboxInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StartSandboxInstance")
	return
}

func NewStartSandboxInstanceResponse() (response *StartSandboxInstanceResponse) {
	response = &StartSandboxInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartSandboxInstance(request *StartSandboxInstanceRequest) string {
	return c.StartSandboxInstanceWithContext(context.Background(), request)
}

func (c *Client) StartSandboxInstanceSend(request *StartSandboxInstanceRequest) (*StartSandboxInstanceResponse, error) {
	statusCode, msg, err := c.StartSandboxInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartSandboxInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartSandboxInstanceWithContext(ctx context.Context, request *StartSandboxInstanceRequest) string {
	if request == nil {
		request = NewStartSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartSandboxInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartSandboxInstanceWithContextV2(ctx context.Context, request *StartSandboxInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewStartSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "StartSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartSandboxInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSandboxTemplateRequest() (request *DeleteSandboxTemplateRequest) {
	request = &DeleteSandboxTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteSandboxTemplate")
	return
}

func NewDeleteSandboxTemplateResponse() (response *DeleteSandboxTemplateResponse) {
	response = &DeleteSandboxTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSandboxTemplate(request *DeleteSandboxTemplateRequest) string {
	return c.DeleteSandboxTemplateWithContext(context.Background(), request)
}

func (c *Client) DeleteSandboxTemplateSend(request *DeleteSandboxTemplateRequest) (*DeleteSandboxTemplateResponse, error) {
	statusCode, msg, err := c.DeleteSandboxTemplateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSandboxTemplateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSandboxTemplateWithContext(ctx context.Context, request *DeleteSandboxTemplateRequest) string {
	if request == nil {
		request = NewDeleteSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSandboxTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSandboxTemplateWithContextV2(ctx context.Context, request *DeleteSandboxTemplateRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSandboxTemplateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetSandboxTemplateRequest() (request *GetSandboxTemplateRequest) {
	request = &GetSandboxTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxTemplate")
	return
}

func NewGetSandboxTemplateResponse() (response *GetSandboxTemplateResponse) {
	response = &GetSandboxTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetSandboxTemplate(request *GetSandboxTemplateRequest) string {
	return c.GetSandboxTemplateWithContext(context.Background(), request)
}

func (c *Client) GetSandboxTemplateSend(request *GetSandboxTemplateRequest) (*GetSandboxTemplateResponse, error) {
	statusCode, msg, err := c.GetSandboxTemplateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetSandboxTemplateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetSandboxTemplateWithContext(ctx context.Context, request *GetSandboxTemplateRequest) string {
	if request == nil {
		request = NewGetSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetSandboxTemplateWithContextV2(ctx context.Context, request *GetSandboxTemplateRequest) (int, string, error) {
	if request == nil {
		request = NewGetSandboxTemplateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetSandboxTemplate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSandboxTemplateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPublicImageListRequest() (request *GetPublicImageListRequest) {
	request = &GetPublicImageListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetPublicImageList")
	return
}

func NewGetPublicImageListResponse() (response *GetPublicImageListResponse) {
	response = &GetPublicImageListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPublicImageList(request *GetPublicImageListRequest) string {
	return c.GetPublicImageListWithContext(context.Background(), request)
}

func (c *Client) GetPublicImageListSend(request *GetPublicImageListRequest) (*GetPublicImageListResponse, error) {
	statusCode, msg, err := c.GetPublicImageListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPublicImageListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPublicImageListWithContext(ctx context.Context, request *GetPublicImageListRequest) string {
	if request == nil {
		request = NewGetPublicImageListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetPublicImageList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPublicImageListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPublicImageListWithContextV2(ctx context.Context, request *GetPublicImageListRequest) (int, string, error) {
	if request == nil {
		request = NewGetPublicImageListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetPublicImageList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPublicImageListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateSandboxInstanceRequest() (request *UpdateSandboxInstanceRequest) {
	request = &UpdateSandboxInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "UpdateSandboxInstance")
	return
}

func NewUpdateSandboxInstanceResponse() (response *UpdateSandboxInstanceResponse) {
	response = &UpdateSandboxInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateSandboxInstance(request *UpdateSandboxInstanceRequest) string {
	return c.UpdateSandboxInstanceWithContext(context.Background(), request)
}

func (c *Client) UpdateSandboxInstanceSend(request *UpdateSandboxInstanceRequest) (*UpdateSandboxInstanceResponse, error) {
	statusCode, msg, err := c.UpdateSandboxInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateSandboxInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateSandboxInstanceWithContext(ctx context.Context, request *UpdateSandboxInstanceRequest) string {
	if request == nil {
		request = NewUpdateSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateSandboxInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateSandboxInstanceWithContextV2(ctx context.Context, request *UpdateSandboxInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateSandboxInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateSandboxInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateSandboxInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
