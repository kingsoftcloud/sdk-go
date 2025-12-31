package v20200901

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-09-01"

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

func NewCreateUserUsageDataExportTaskRequest() (request *CreateUserUsageDataExportTaskRequest) {
	request = &CreateUserUsageDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "CreateUserUsageDataExportTask")
	return
}

func NewCreateUserUsageDataExportTaskResponse() (response *CreateUserUsageDataExportTaskResponse) {
	response = &CreateUserUsageDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateUserUsageDataExportTask(request *CreateUserUsageDataExportTaskRequest) string {
	return c.CreateUserUsageDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) CreateUserUsageDataExportTaskSend(request *CreateUserUsageDataExportTaskRequest) (*CreateUserUsageDataExportTaskResponse, error) {
	statusCode, msg, err := c.CreateUserUsageDataExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateUserUsageDataExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateUserUsageDataExportTaskWithContext(ctx context.Context, request *CreateUserUsageDataExportTaskRequest) string {
	if request == nil {
		request = NewCreateUserUsageDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "CreateUserUsageDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateUserUsageDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateUserUsageDataExportTaskWithContextV2(ctx context.Context, request *CreateUserUsageDataExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateUserUsageDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "CreateUserUsageDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateUserUsageDataExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetUserUsageDataExportTaskRequest() (request *GetUserUsageDataExportTaskRequest) {
	request = &GetUserUsageDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetUserUsageDataExportTask")
	return
}

func NewGetUserUsageDataExportTaskResponse() (response *GetUserUsageDataExportTaskResponse) {
	response = &GetUserUsageDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUserUsageDataExportTask(request *GetUserUsageDataExportTaskRequest) string {
	return c.GetUserUsageDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) GetUserUsageDataExportTaskSend(request *GetUserUsageDataExportTaskRequest) (*GetUserUsageDataExportTaskResponse, error) {
	statusCode, msg, err := c.GetUserUsageDataExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetUserUsageDataExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetUserUsageDataExportTaskWithContext(ctx context.Context, request *GetUserUsageDataExportTaskRequest) string {
	if request == nil {
		request = NewGetUserUsageDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetUserUsageDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUserUsageDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetUserUsageDataExportTaskWithContextV2(ctx context.Context, request *GetUserUsageDataExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewGetUserUsageDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetUserUsageDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUserUsageDataExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteUserUsageDataExportTaskRequest() (request *DeleteUserUsageDataExportTaskRequest) {
	request = &DeleteUserUsageDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "DeleteUserUsageDataExportTask")
	return
}

func NewDeleteUserUsageDataExportTaskResponse() (response *DeleteUserUsageDataExportTaskResponse) {
	response = &DeleteUserUsageDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteUserUsageDataExportTask(request *DeleteUserUsageDataExportTaskRequest) string {
	return c.DeleteUserUsageDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteUserUsageDataExportTaskSend(request *DeleteUserUsageDataExportTaskRequest) (*DeleteUserUsageDataExportTaskResponse, error) {
	statusCode, msg, err := c.DeleteUserUsageDataExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteUserUsageDataExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteUserUsageDataExportTaskWithContext(ctx context.Context, request *DeleteUserUsageDataExportTaskRequest) string {
	if request == nil {
		request = NewDeleteUserUsageDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteUserUsageDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteUserUsageDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteUserUsageDataExportTaskWithContextV2(ctx context.Context, request *DeleteUserUsageDataExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteUserUsageDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteUserUsageDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteUserUsageDataExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainUsageDataRequest() (request *GetDomainUsageDataRequest) {
	request = &GetDomainUsageDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainUsageData")
	return
}

func NewGetDomainUsageDataResponse() (response *GetDomainUsageDataResponse) {
	response = &GetDomainUsageDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainUsageData(request *GetDomainUsageDataRequest) string {
	return c.GetDomainUsageDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainUsageDataSend(request *GetDomainUsageDataRequest) (*GetDomainUsageDataResponse, error) {
	statusCode, msg, err := c.GetDomainUsageDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetDomainUsageDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainUsageDataWithContext(ctx context.Context, request *GetDomainUsageDataRequest) string {
	if request == nil {
		request = NewGetDomainUsageDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainUsageData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainUsageDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainUsageDataWithContextV2(ctx context.Context, request *GetDomainUsageDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainUsageDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainUsageData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainUsageDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateUsageDetailDataExportTaskRequest() (request *CreateUsageDetailDataExportTaskRequest) {
	request = &CreateUsageDetailDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "CreateUsageDetailDataExportTask")
	return
}

func NewCreateUsageDetailDataExportTaskResponse() (response *CreateUsageDetailDataExportTaskResponse) {
	response = &CreateUsageDetailDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateUsageDetailDataExportTask(request *CreateUsageDetailDataExportTaskRequest) string {
	return c.CreateUsageDetailDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) CreateUsageDetailDataExportTaskSend(request *CreateUsageDetailDataExportTaskRequest) (*CreateUsageDetailDataExportTaskResponse, error) {
	statusCode, msg, err := c.CreateUsageDetailDataExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateUsageDetailDataExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateUsageDetailDataExportTaskWithContext(ctx context.Context, request *CreateUsageDetailDataExportTaskRequest) string {
	if request == nil {
		request = NewCreateUsageDetailDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "CreateUsageDetailDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateUsageDetailDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateUsageDetailDataExportTaskWithContextV2(ctx context.Context, request *CreateUsageDetailDataExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateUsageDetailDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "CreateUsageDetailDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateUsageDetailDataExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetUsageDetailDataExportTaskRequest() (request *GetUsageDetailDataExportTaskRequest) {
	request = &GetUsageDetailDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetUsageDetailDataExportTask")
	return
}

func NewGetUsageDetailDataExportTaskResponse() (response *GetUsageDetailDataExportTaskResponse) {
	response = &GetUsageDetailDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUsageDetailDataExportTask(request *GetUsageDetailDataExportTaskRequest) string {
	return c.GetUsageDetailDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) GetUsageDetailDataExportTaskSend(request *GetUsageDetailDataExportTaskRequest) (*GetUsageDetailDataExportTaskResponse, error) {
	statusCode, msg, err := c.GetUsageDetailDataExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetUsageDetailDataExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetUsageDetailDataExportTaskWithContext(ctx context.Context, request *GetUsageDetailDataExportTaskRequest) string {
	if request == nil {
		request = NewGetUsageDetailDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetUsageDetailDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUsageDetailDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetUsageDetailDataExportTaskWithContextV2(ctx context.Context, request *GetUsageDetailDataExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewGetUsageDetailDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetUsageDetailDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUsageDetailDataExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteUsageDetailDataExportTaskRequest() (request *DeleteUsageDetailDataExportTaskRequest) {
	request = &DeleteUsageDetailDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "DeleteUsageDetailDataExportTask")
	return
}

func NewDeleteUsageDetailDataExportTaskResponse() (response *DeleteUsageDetailDataExportTaskResponse) {
	response = &DeleteUsageDetailDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteUsageDetailDataExportTask(request *DeleteUsageDetailDataExportTaskRequest) string {
	return c.DeleteUsageDetailDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteUsageDetailDataExportTaskSend(request *DeleteUsageDetailDataExportTaskRequest) (*DeleteUsageDetailDataExportTaskResponse, error) {
	statusCode, msg, err := c.DeleteUsageDetailDataExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteUsageDetailDataExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteUsageDetailDataExportTaskWithContext(ctx context.Context, request *DeleteUsageDetailDataExportTaskRequest) string {
	if request == nil {
		request = NewDeleteUsageDetailDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteUsageDetailDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteUsageDetailDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteUsageDetailDataExportTaskWithContextV2(ctx context.Context, request *DeleteUsageDetailDataExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteUsageDetailDataExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteUsageDetailDataExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteUsageDetailDataExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
