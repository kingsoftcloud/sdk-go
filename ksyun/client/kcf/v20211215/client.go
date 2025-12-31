package v20211215

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-12-15"

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

func NewGetLogDateRequest() (request *GetLogDateRequest) {
	request = &GetLogDateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "GetLogDate")
	return
}

func NewGetLogDateResponse() (response *GetLogDateResponse) {
	response = &GetLogDateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetLogDate(request *GetLogDateRequest) string {
	return c.GetLogDateWithContext(context.Background(), request)
}

func (c *Client) GetLogDateSend(request *GetLogDateRequest) (*GetLogDateResponse, error) {
	statusCode, msg, err := c.GetLogDateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetLogDateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetLogDateWithContext(ctx context.Context, request *GetLogDateRequest) string {
	if request == nil {
		request = NewGetLogDateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "GetLogDate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLogDateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetLogDateWithContextV2(ctx context.Context, request *GetLogDateRequest) (int, string, error) {
	if request == nil {
		request = NewGetLogDateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "GetLogDate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLogDateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateFunctionRequest() (request *CreateFunctionRequest) {
	request = &CreateFunctionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "CreateFunction")
	return
}

func NewCreateFunctionResponse() (response *CreateFunctionResponse) {
	response = &CreateFunctionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFunction(request *CreateFunctionRequest) string {
	return c.CreateFunctionWithContext(context.Background(), request)
}

func (c *Client) CreateFunctionSend(request *CreateFunctionRequest) (*CreateFunctionResponse, error) {
	statusCode, msg, err := c.CreateFunctionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateFunctionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateFunctionWithContext(ctx context.Context, request *CreateFunctionRequest) string {
	if request == nil {
		request = NewCreateFunctionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "CreateFunction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFunctionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateFunctionWithContextV2(ctx context.Context, request *CreateFunctionRequest) (int, string, error) {
	if request == nil {
		request = NewCreateFunctionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "CreateFunction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFunctionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCheckFunctionServiceRequest() (request *CheckFunctionServiceRequest) {
	request = &CheckFunctionServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "CheckFunctionService")
	return
}

func NewCheckFunctionServiceResponse() (response *CheckFunctionServiceResponse) {
	response = &CheckFunctionServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckFunctionService(request *CheckFunctionServiceRequest) string {
	return c.CheckFunctionServiceWithContext(context.Background(), request)
}

func (c *Client) CheckFunctionServiceSend(request *CheckFunctionServiceRequest) (*CheckFunctionServiceResponse, error) {
	statusCode, msg, err := c.CheckFunctionServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CheckFunctionServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CheckFunctionServiceWithContext(ctx context.Context, request *CheckFunctionServiceRequest) string {
	if request == nil {
		request = NewCheckFunctionServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "CheckFunctionService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCheckFunctionServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CheckFunctionServiceWithContextV2(ctx context.Context, request *CheckFunctionServiceRequest) (int, string, error) {
	if request == nil {
		request = NewCheckFunctionServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "CheckFunctionService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCheckFunctionServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOpenFunctionServiceRequest() (request *OpenFunctionServiceRequest) {
	request = &OpenFunctionServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "OpenFunctionService")
	return
}

func NewOpenFunctionServiceResponse() (response *OpenFunctionServiceResponse) {
	response = &OpenFunctionServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OpenFunctionService(request *OpenFunctionServiceRequest) string {
	return c.OpenFunctionServiceWithContext(context.Background(), request)
}

func (c *Client) OpenFunctionServiceSend(request *OpenFunctionServiceRequest) (*OpenFunctionServiceResponse, error) {
	statusCode, msg, err := c.OpenFunctionServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OpenFunctionServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OpenFunctionServiceWithContext(ctx context.Context, request *OpenFunctionServiceRequest) string {
	if request == nil {
		request = NewOpenFunctionServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "OpenFunctionService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOpenFunctionServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OpenFunctionServiceWithContextV2(ctx context.Context, request *OpenFunctionServiceRequest) (int, string, error) {
	if request == nil {
		request = NewOpenFunctionServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "OpenFunctionService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOpenFunctionServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteFunctionRequest() (request *DeleteFunctionRequest) {
	request = &DeleteFunctionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "DeleteFunction")
	return
}

func NewDeleteFunctionResponse() (response *DeleteFunctionResponse) {
	response = &DeleteFunctionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFunction(request *DeleteFunctionRequest) string {
	return c.DeleteFunctionWithContext(context.Background(), request)
}

func (c *Client) DeleteFunctionSend(request *DeleteFunctionRequest) (*DeleteFunctionResponse, error) {
	statusCode, msg, err := c.DeleteFunctionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteFunctionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFunctionWithContext(ctx context.Context, request *DeleteFunctionRequest) string {
	if request == nil {
		request = NewDeleteFunctionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "DeleteFunction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFunctionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFunctionWithContextV2(ctx context.Context, request *DeleteFunctionRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFunctionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "DeleteFunction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFunctionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateTriggerRequest() (request *CreateTriggerRequest) {
	request = &CreateTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "CreateTrigger")
	return
}

func NewCreateTriggerResponse() (response *CreateTriggerResponse) {
	response = &CreateTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTrigger(request *CreateTriggerRequest) string {
	return c.CreateTriggerWithContext(context.Background(), request)
}

func (c *Client) CreateTriggerSend(request *CreateTriggerRequest) (*CreateTriggerResponse, error) {
	statusCode, msg, err := c.CreateTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTriggerWithContext(ctx context.Context, request *CreateTriggerRequest) string {
	if request == nil {
		request = NewCreateTriggerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "CreateTrigger")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTriggerWithContextV2(ctx context.Context, request *CreateTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTriggerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "CreateTrigger")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteTriggerRequest() (request *DeleteTriggerRequest) {
	request = &DeleteTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "DeleteTrigger")
	return
}

func NewDeleteTriggerResponse() (response *DeleteTriggerResponse) {
	response = &DeleteTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTrigger(request *DeleteTriggerRequest) string {
	return c.DeleteTriggerWithContext(context.Background(), request)
}

func (c *Client) DeleteTriggerSend(request *DeleteTriggerRequest) (*DeleteTriggerResponse, error) {
	statusCode, msg, err := c.DeleteTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteTriggerWithContext(ctx context.Context, request *DeleteTriggerRequest) string {
	if request == nil {
		request = NewDeleteTriggerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "DeleteTrigger")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteTriggerWithContextV2(ctx context.Context, request *DeleteTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteTriggerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "DeleteTrigger")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyFunctionRequest() (request *ModifyFunctionRequest) {
	request = &ModifyFunctionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "ModifyFunction")
	return
}

func NewModifyFunctionResponse() (response *ModifyFunctionResponse) {
	response = &ModifyFunctionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyFunction(request *ModifyFunctionRequest) string {
	return c.ModifyFunctionWithContext(context.Background(), request)
}

func (c *Client) ModifyFunctionSend(request *ModifyFunctionRequest) (*ModifyFunctionResponse, error) {
	statusCode, msg, err := c.ModifyFunctionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyFunctionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyFunctionWithContext(ctx context.Context, request *ModifyFunctionRequest) string {
	if request == nil {
		request = NewModifyFunctionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "ModifyFunction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyFunctionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyFunctionWithContextV2(ctx context.Context, request *ModifyFunctionRequest) (int, string, error) {
	if request == nil {
		request = NewModifyFunctionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "ModifyFunction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyFunctionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTriggersRequest() (request *DescribeTriggersRequest) {
	request = &DescribeTriggersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "DescribeTriggers")
	return
}

func NewDescribeTriggersResponse() (response *DescribeTriggersResponse) {
	response = &DescribeTriggersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTriggers(request *DescribeTriggersRequest) string {
	return c.DescribeTriggersWithContext(context.Background(), request)
}

func (c *Client) DescribeTriggersSend(request *DescribeTriggersRequest) (*DescribeTriggersResponse, error) {
	statusCode, msg, err := c.DescribeTriggersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTriggersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTriggersWithContext(ctx context.Context, request *DescribeTriggersRequest) string {
	if request == nil {
		request = NewDescribeTriggersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "DescribeTriggers")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTriggersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTriggersWithContextV2(ctx context.Context, request *DescribeTriggersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTriggersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcf", APIVersion, "DescribeTriggers")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTriggersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
