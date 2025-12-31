package v20180108

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-01-08"

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

func NewSchemaStructRequest() (request *SchemaStructRequest) {
	request = &SchemaStructRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "SchemaStruct")
	return
}

func NewSchemaStructResponse() (response *SchemaStructResponse) {
	response = &SchemaStructResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SchemaStruct(request *SchemaStructRequest) string {
	return c.SchemaStructWithContext(context.Background(), request)
}

func (c *Client) SchemaStructSend(request *SchemaStructRequest) (*SchemaStructResponse, error) {
	statusCode, msg, err := c.SchemaStructWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SchemaStructResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SchemaStructWithContext(ctx context.Context, request *SchemaStructRequest) string {
	if request == nil {
		request = NewSchemaStructRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "SchemaStruct")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSchemaStructResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SchemaStructWithContextV2(ctx context.Context, request *SchemaStructRequest) (int, string, error) {
	if request == nil {
		request = NewSchemaStructRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "SchemaStruct")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSchemaStructResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewConnectivityCheckRequest() (request *ConnectivityCheckRequest) {
	request = &ConnectivityCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ConnectivityCheck")
	return
}

func NewConnectivityCheckResponse() (response *ConnectivityCheckResponse) {
	response = &ConnectivityCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ConnectivityCheck(request *ConnectivityCheckRequest) string {
	return c.ConnectivityCheckWithContext(context.Background(), request)
}

func (c *Client) ConnectivityCheckSend(request *ConnectivityCheckRequest) (*ConnectivityCheckResponse, error) {
	statusCode, msg, err := c.ConnectivityCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ConnectivityCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ConnectivityCheckWithContext(ctx context.Context, request *ConnectivityCheckRequest) string {
	if request == nil {
		request = NewConnectivityCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "ConnectivityCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewConnectivityCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ConnectivityCheckWithContextV2(ctx context.Context, request *ConnectivityCheckRequest) (int, string, error) {
	if request == nil {
		request = NewConnectivityCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "ConnectivityCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewConnectivityCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreatePrecheckRequest() (request *CreatePrecheckRequest) {
	request = &CreatePrecheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreatePrecheck")
	return
}

func NewCreatePrecheckResponse() (response *CreatePrecheckResponse) {
	response = &CreatePrecheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePrecheck(request *CreatePrecheckRequest) string {
	return c.CreatePrecheckWithContext(context.Background(), request)
}

func (c *Client) CreatePrecheckSend(request *CreatePrecheckRequest) (*CreatePrecheckResponse, error) {
	statusCode, msg, err := c.CreatePrecheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreatePrecheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreatePrecheckWithContext(ctx context.Context, request *CreatePrecheckRequest) string {
	if request == nil {
		request = NewCreatePrecheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "CreatePrecheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreatePrecheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreatePrecheckWithContextV2(ctx context.Context, request *CreatePrecheckRequest) (int, string, error) {
	if request == nil {
		request = NewCreatePrecheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "CreatePrecheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreatePrecheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateTask")
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

func (c *Client) CreateTaskSend(request *CreateTaskRequest) (*CreateTaskResponse, error) {
	statusCode, msg, err := c.CreateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) string {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "CreateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTaskWithContextV2(ctx context.Context, request *CreateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "CreateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeTask")
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

func (c *Client) DescribeTaskSend(request *DescribeTaskRequest) (*DescribeTaskResponse, error) {
	statusCode, msg, err := c.DescribeTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) string {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeTask")
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

func (c *Client) DescribeTaskWithContextV2(ctx context.Context, request *DescribeTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOperateTaskRequest() (request *OperateTaskRequest) {
	request = &OperateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "OperateTask")
	return
}

func NewOperateTaskResponse() (response *OperateTaskResponse) {
	response = &OperateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperateTask(request *OperateTaskRequest) string {
	return c.OperateTaskWithContext(context.Background(), request)
}

func (c *Client) OperateTaskSend(request *OperateTaskRequest) (*OperateTaskResponse, error) {
	statusCode, msg, err := c.OperateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OperateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OperateTaskWithContext(ctx context.Context, request *OperateTaskRequest) string {
	if request == nil {
		request = NewOperateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "OperateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOperateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OperateTaskWithContextV2(ctx context.Context, request *OperateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewOperateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "OperateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOperateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeConnConfigRequest() (request *DescribeConnConfigRequest) {
	request = &DescribeConnConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeConnConfig")
	return
}

func NewDescribeConnConfigResponse() (response *DescribeConnConfigResponse) {
	response = &DescribeConnConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeConnConfig(request *DescribeConnConfigRequest) string {
	return c.DescribeConnConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeConnConfigSend(request *DescribeConnConfigRequest) (*DescribeConnConfigResponse, error) {
	statusCode, msg, err := c.DescribeConnConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeConnConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeConnConfigWithContext(ctx context.Context, request *DescribeConnConfigRequest) string {
	if request == nil {
		request = NewDescribeConnConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeConnConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeConnConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeConnConfigWithContextV2(ctx context.Context, request *DescribeConnConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeConnConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeConnConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeConnConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePrecheckRequest() (request *DescribePrecheckRequest) {
	request = &DescribePrecheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribePrecheck")
	return
}

func NewDescribePrecheckResponse() (response *DescribePrecheckResponse) {
	response = &DescribePrecheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrecheck(request *DescribePrecheckRequest) string {
	return c.DescribePrecheckWithContext(context.Background(), request)
}

func (c *Client) DescribePrecheckSend(request *DescribePrecheckRequest) (*DescribePrecheckResponse, error) {
	statusCode, msg, err := c.DescribePrecheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribePrecheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePrecheckWithContext(ctx context.Context, request *DescribePrecheckRequest) string {
	if request == nil {
		request = NewDescribePrecheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribePrecheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrecheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePrecheckWithContextV2(ctx context.Context, request *DescribePrecheckRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePrecheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribePrecheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrecheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSourceUserConfigRequest() (request *DescribeSourceUserConfigRequest) {
	request = &DescribeSourceUserConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUserConfig")
	return
}

func NewDescribeSourceUserConfigResponse() (response *DescribeSourceUserConfigResponse) {
	response = &DescribeSourceUserConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSourceUserConfig(request *DescribeSourceUserConfigRequest) string {
	return c.DescribeSourceUserConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeSourceUserConfigSend(request *DescribeSourceUserConfigRequest) (*DescribeSourceUserConfigResponse, error) {
	statusCode, msg, err := c.DescribeSourceUserConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSourceUserConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSourceUserConfigWithContext(ctx context.Context, request *DescribeSourceUserConfigRequest) string {
	if request == nil {
		request = NewDescribeSourceUserConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUserConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSourceUserConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSourceUserConfigWithContextV2(ctx context.Context, request *DescribeSourceUserConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSourceUserConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUserConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSourceUserConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetConsistencyCheckRequest() (request *SetConsistencyCheckRequest) {
	request = &SetConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "SetConsistencyCheck")
	return
}

func NewSetConsistencyCheckResponse() (response *SetConsistencyCheckResponse) {
	response = &SetConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetConsistencyCheck(request *SetConsistencyCheckRequest) string {
	return c.SetConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) SetConsistencyCheckSend(request *SetConsistencyCheckRequest) (*SetConsistencyCheckResponse, error) {
	statusCode, msg, err := c.SetConsistencyCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetConsistencyCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetConsistencyCheckWithContext(ctx context.Context, request *SetConsistencyCheckRequest) string {
	if request == nil {
		request = NewSetConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "SetConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetConsistencyCheckWithContextV2(ctx context.Context, request *SetConsistencyCheckRequest) (int, string, error) {
	if request == nil {
		request = NewSetConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "SetConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetConsistencyCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeConsistencyCheckRequest() (request *DescribeConsistencyCheckRequest) {
	request = &DescribeConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeConsistencyCheck")
	return
}

func NewDescribeConsistencyCheckResponse() (response *DescribeConsistencyCheckResponse) {
	response = &DescribeConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeConsistencyCheck(request *DescribeConsistencyCheckRequest) string {
	return c.DescribeConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) DescribeConsistencyCheckSend(request *DescribeConsistencyCheckRequest) (*DescribeConsistencyCheckResponse, error) {
	statusCode, msg, err := c.DescribeConsistencyCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeConsistencyCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeConsistencyCheckWithContext(ctx context.Context, request *DescribeConsistencyCheckRequest) string {
	if request == nil {
		request = NewDescribeConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeConsistencyCheckWithContextV2(ctx context.Context, request *DescribeConsistencyCheckRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeConsistencyCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDTSParameterRequest() (request *DescribeDTSParameterRequest) {
	request = &DescribeDTSParameterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameter")
	return
}

func NewDescribeDTSParameterResponse() (response *DescribeDTSParameterResponse) {
	response = &DescribeDTSParameterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDTSParameter(request *DescribeDTSParameterRequest) string {
	return c.DescribeDTSParameterWithContext(context.Background(), request)
}

func (c *Client) DescribeDTSParameterSend(request *DescribeDTSParameterRequest) (*DescribeDTSParameterResponse, error) {
	statusCode, msg, err := c.DescribeDTSParameterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDTSParameterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDTSParameterWithContext(ctx context.Context, request *DescribeDTSParameterRequest) string {
	if request == nil {
		request = NewDescribeDTSParameterRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameter")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDTSParameterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDTSParameterWithContextV2(ctx context.Context, request *DescribeDTSParameterRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDTSParameterRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameter")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDTSParameterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDTSParameterConfigRequest() (request *DescribeDTSParameterConfigRequest) {
	request = &DescribeDTSParameterConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameterConfig")
	return
}

func NewDescribeDTSParameterConfigResponse() (response *DescribeDTSParameterConfigResponse) {
	response = &DescribeDTSParameterConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDTSParameterConfig(request *DescribeDTSParameterConfigRequest) string {
	return c.DescribeDTSParameterConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeDTSParameterConfigSend(request *DescribeDTSParameterConfigRequest) (*DescribeDTSParameterConfigResponse, error) {
	statusCode, msg, err := c.DescribeDTSParameterConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDTSParameterConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDTSParameterConfigWithContext(ctx context.Context, request *DescribeDTSParameterConfigRequest) string {
	if request == nil {
		request = NewDescribeDTSParameterConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameterConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDTSParameterConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDTSParameterConfigWithContextV2(ctx context.Context, request *DescribeDTSParameterConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDTSParameterConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameterConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDTSParameterConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSourceUserRequest() (request *DescribeSourceUserRequest) {
	request = &DescribeSourceUserRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUser")
	return
}

func NewDescribeSourceUserResponse() (response *DescribeSourceUserResponse) {
	response = &DescribeSourceUserResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSourceUser(request *DescribeSourceUserRequest) string {
	return c.DescribeSourceUserWithContext(context.Background(), request)
}

func (c *Client) DescribeSourceUserSend(request *DescribeSourceUserRequest) (*DescribeSourceUserResponse, error) {
	statusCode, msg, err := c.DescribeSourceUserWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSourceUserResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSourceUserWithContext(ctx context.Context, request *DescribeSourceUserRequest) string {
	if request == nil {
		request = NewDescribeSourceUserRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUser")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSourceUserResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSourceUserWithContextV2(ctx context.Context, request *DescribeSourceUserRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSourceUserRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUser")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSourceUserResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSubTaskRequest() (request *DescribeSubTaskRequest) {
	request = &DescribeSubTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubTask")
	return
}

func NewDescribeSubTaskResponse() (response *DescribeSubTaskResponse) {
	response = &DescribeSubTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubTask(request *DescribeSubTaskRequest) string {
	return c.DescribeSubTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeSubTaskSend(request *DescribeSubTaskRequest) (*DescribeSubTaskResponse, error) {
	statusCode, msg, err := c.DescribeSubTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSubTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSubTaskWithContext(ctx context.Context, request *DescribeSubTaskRequest) string {
	if request == nil {
		request = NewDescribeSubTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeSubTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSubTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSubTaskWithContextV2(ctx context.Context, request *DescribeSubTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSubTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeSubTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSubTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateConsistencyCheckRequest() (request *CreateConsistencyCheckRequest) {
	request = &CreateConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateConsistencyCheck")
	return
}

func NewCreateConsistencyCheckResponse() (response *CreateConsistencyCheckResponse) {
	response = &CreateConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateConsistencyCheck(request *CreateConsistencyCheckRequest) string {
	return c.CreateConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) CreateConsistencyCheckSend(request *CreateConsistencyCheckRequest) (*CreateConsistencyCheckResponse, error) {
	statusCode, msg, err := c.CreateConsistencyCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateConsistencyCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateConsistencyCheckWithContext(ctx context.Context, request *CreateConsistencyCheckRequest) string {
	if request == nil {
		request = NewCreateConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "CreateConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateConsistencyCheckWithContextV2(ctx context.Context, request *CreateConsistencyCheckRequest) (int, string, error) {
	if request == nil {
		request = NewCreateConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "CreateConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateConsistencyCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopConsistencyCheckRequest() (request *StopConsistencyCheckRequest) {
	request = &StopConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "StopConsistencyCheck")
	return
}

func NewStopConsistencyCheckResponse() (response *StopConsistencyCheckResponse) {
	response = &StopConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopConsistencyCheck(request *StopConsistencyCheckRequest) string {
	return c.StopConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) StopConsistencyCheckSend(request *StopConsistencyCheckRequest) (*StopConsistencyCheckResponse, error) {
	statusCode, msg, err := c.StopConsistencyCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopConsistencyCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopConsistencyCheckWithContext(ctx context.Context, request *StopConsistencyCheckRequest) string {
	if request == nil {
		request = NewStopConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "StopConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopConsistencyCheckWithContextV2(ctx context.Context, request *StopConsistencyCheckRequest) (int, string, error) {
	if request == nil {
		request = NewStopConsistencyCheckRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "StopConsistencyCheck")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopConsistencyCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRegionConfigRequest() (request *DescribeRegionConfigRequest) {
	request = &DescribeRegionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConfig")
	return
}

func NewDescribeRegionConfigResponse() (response *DescribeRegionConfigResponse) {
	response = &DescribeRegionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegionConfig(request *DescribeRegionConfigRequest) string {
	return c.DescribeRegionConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionConfigSend(request *DescribeRegionConfigRequest) (*DescribeRegionConfigResponse, error) {
	statusCode, msg, err := c.DescribeRegionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeRegionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRegionConfigWithContext(ctx context.Context, request *DescribeRegionConfigRequest) string {
	if request == nil {
		request = NewDescribeRegionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRegionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRegionConfigWithContextV2(ctx context.Context, request *DescribeRegionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRegionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRegionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewTaskBirdViewRequest() (request *TaskBirdViewRequest) {
	request = &TaskBirdViewRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "TaskBirdView")
	return
}

func NewTaskBirdViewResponse() (response *TaskBirdViewResponse) {
	response = &TaskBirdViewResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TaskBirdView(request *TaskBirdViewRequest) string {
	return c.TaskBirdViewWithContext(context.Background(), request)
}

func (c *Client) TaskBirdViewSend(request *TaskBirdViewRequest) (*TaskBirdViewResponse, error) {
	statusCode, msg, err := c.TaskBirdViewWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct TaskBirdViewResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) TaskBirdViewWithContext(ctx context.Context, request *TaskBirdViewRequest) string {
	if request == nil {
		request = NewTaskBirdViewRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "TaskBirdView")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewTaskBirdViewResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) TaskBirdViewWithContextV2(ctx context.Context, request *TaskBirdViewRequest) (int, string, error) {
	if request == nil {
		request = NewTaskBirdViewRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("dts", APIVersion, "TaskBirdView")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewTaskBirdViewResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
