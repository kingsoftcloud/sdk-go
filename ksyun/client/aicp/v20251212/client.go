package v20251212

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-12-12"

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
func NewDescribeTrainJobsRequest() (request *DescribeTrainJobsRequest) {
	request = &DescribeTrainJobsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobs")
	return
}

func NewDescribeTrainJobsResponse() (response *DescribeTrainJobsResponse) {
	response = &DescribeTrainJobsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrainJobs(request *DescribeTrainJobsRequest) string {
	return c.DescribeTrainJobsWithContext(context.Background(), request)
}

func (c *Client) DescribeTrainJobsSend(request *DescribeTrainJobsRequest) (*DescribeTrainJobsResponse, error) {
	statusCode, msg, err := c.DescribeTrainJobsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrainJobsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrainJobsWithContext(ctx context.Context, request *DescribeTrainJobsRequest) string {
	if request == nil {
		request = NewDescribeTrainJobsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrainJobsWithContextV2(ctx context.Context, request *DescribeTrainJobsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrainJobsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyModelAccessRequest() (request *ModifyModelAccessRequest) {
	request = &ModifyModelAccessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyModelAccess")
	return
}

func NewModifyModelAccessResponse() (response *ModifyModelAccessResponse) {
	response = &ModifyModelAccessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyModelAccess(request *ModifyModelAccessRequest) string {
	return c.ModifyModelAccessWithContext(context.Background(), request)
}

func (c *Client) ModifyModelAccessSend(request *ModifyModelAccessRequest) (*ModifyModelAccessResponse, error) {
	statusCode, msg, err := c.ModifyModelAccessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyModelAccessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyModelAccessWithContext(ctx context.Context, request *ModifyModelAccessRequest) string {
	if request == nil {
		request = NewModifyModelAccessRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyModelAccess")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyModelAccessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyModelAccessWithContextV2(ctx context.Context, request *ModifyModelAccessRequest) (int, string, error) {
	if request == nil {
		request = NewModifyModelAccessRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyModelAccess")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyModelAccessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateModelAndVersionRequest() (request *CreateModelAndVersionRequest) {
	request = &CreateModelAndVersionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateModelAndVersion")
	return
}

func NewCreateModelAndVersionResponse() (response *CreateModelAndVersionResponse) {
	response = &CreateModelAndVersionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateModelAndVersion(request *CreateModelAndVersionRequest) string {
	return c.CreateModelAndVersionWithContext(context.Background(), request)
}

func (c *Client) CreateModelAndVersionSend(request *CreateModelAndVersionRequest) (*CreateModelAndVersionResponse, error) {
	statusCode, msg, err := c.CreateModelAndVersionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateModelAndVersionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateModelAndVersionWithContext(ctx context.Context, request *CreateModelAndVersionRequest) string {
	if request == nil {
		request = NewCreateModelAndVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateModelAndVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateModelAndVersionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateModelAndVersionWithContextV2(ctx context.Context, request *CreateModelAndVersionRequest) (int, string, error) {
	if request == nil {
		request = NewCreateModelAndVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateModelAndVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateModelAndVersionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyModelRequest() (request *ModifyModelRequest) {
	request = &ModifyModelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyModel")
	return
}

func NewModifyModelResponse() (response *ModifyModelResponse) {
	response = &ModifyModelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyModel(request *ModifyModelRequest) string {
	return c.ModifyModelWithContext(context.Background(), request)
}

func (c *Client) ModifyModelSend(request *ModifyModelRequest) (*ModifyModelResponse, error) {
	statusCode, msg, err := c.ModifyModelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyModelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyModelWithContext(ctx context.Context, request *ModifyModelRequest) string {
	if request == nil {
		request = NewModifyModelRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyModel")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyModelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyModelWithContextV2(ctx context.Context, request *ModifyModelRequest) (int, string, error) {
	if request == nil {
		request = NewModifyModelRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyModel")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyModelResponse()
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
func NewDeleteModelRequest() (request *DeleteModelRequest) {
	request = &DeleteModelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteModel")
	return
}

func NewDeleteModelResponse() (response *DeleteModelResponse) {
	response = &DeleteModelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteModel(request *DeleteModelRequest) string {
	return c.DeleteModelWithContext(context.Background(), request)
}

func (c *Client) DeleteModelSend(request *DeleteModelRequest) (*DeleteModelResponse, error) {
	statusCode, msg, err := c.DeleteModelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteModelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest) string {
	if request == nil {
		request = NewDeleteModelRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteModel")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteModelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteModelWithContextV2(ctx context.Context, request *DeleteModelRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteModelRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteModel")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteModelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateModelVersionRequest() (request *CreateModelVersionRequest) {
	request = &CreateModelVersionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateModelVersion")
	return
}

func NewCreateModelVersionResponse() (response *CreateModelVersionResponse) {
	response = &CreateModelVersionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateModelVersion(request *CreateModelVersionRequest) string {
	return c.CreateModelVersionWithContext(context.Background(), request)
}

func (c *Client) CreateModelVersionSend(request *CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	statusCode, msg, err := c.CreateModelVersionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateModelVersionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateModelVersionWithContext(ctx context.Context, request *CreateModelVersionRequest) string {
	if request == nil {
		request = NewCreateModelVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateModelVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateModelVersionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateModelVersionWithContextV2(ctx context.Context, request *CreateModelVersionRequest) (int, string, error) {
	if request == nil {
		request = NewCreateModelVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateModelVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateModelVersionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteModelVersionRequest() (request *DeleteModelVersionRequest) {
	request = &DeleteModelVersionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteModelVersion")
	return
}

func NewDeleteModelVersionResponse() (response *DeleteModelVersionResponse) {
	response = &DeleteModelVersionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteModelVersion(request *DeleteModelVersionRequest) string {
	return c.DeleteModelVersionWithContext(context.Background(), request)
}

func (c *Client) DeleteModelVersionSend(request *DeleteModelVersionRequest) (*DeleteModelVersionResponse, error) {
	statusCode, msg, err := c.DeleteModelVersionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteModelVersionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteModelVersionWithContext(ctx context.Context, request *DeleteModelVersionRequest) string {
	if request == nil {
		request = NewDeleteModelVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteModelVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteModelVersionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteModelVersionWithContextV2(ctx context.Context, request *DeleteModelVersionRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteModelVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteModelVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteModelVersionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyModelVersionRequest() (request *ModifyModelVersionRequest) {
	request = &ModifyModelVersionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyModelVersion")
	return
}

func NewModifyModelVersionResponse() (response *ModifyModelVersionResponse) {
	response = &ModifyModelVersionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyModelVersion(request *ModifyModelVersionRequest) string {
	return c.ModifyModelVersionWithContext(context.Background(), request)
}

func (c *Client) ModifyModelVersionSend(request *ModifyModelVersionRequest) (*ModifyModelVersionResponse, error) {
	statusCode, msg, err := c.ModifyModelVersionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyModelVersionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyModelVersionWithContext(ctx context.Context, request *ModifyModelVersionRequest) string {
	if request == nil {
		request = NewModifyModelVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyModelVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyModelVersionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyModelVersionWithContextV2(ctx context.Context, request *ModifyModelVersionRequest) (int, string, error) {
	if request == nil {
		request = NewModifyModelVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyModelVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyModelVersionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModelVersionsRequest() (request *DescribeModelVersionsRequest) {
	request = &DescribeModelVersionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeModelVersions")
	return
}

func NewDescribeModelVersionsResponse() (response *DescribeModelVersionsResponse) {
	response = &DescribeModelVersionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModelVersions(request *DescribeModelVersionsRequest) string {
	return c.DescribeModelVersionsWithContext(context.Background(), request)
}

func (c *Client) DescribeModelVersionsSend(request *DescribeModelVersionsRequest) (*DescribeModelVersionsResponse, error) {
	statusCode, msg, err := c.DescribeModelVersionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeModelVersionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModelVersionsWithContext(ctx context.Context, request *DescribeModelVersionsRequest) string {
	if request == nil {
		request = NewDescribeModelVersionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeModelVersions")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelVersionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModelVersionsWithContextV2(ctx context.Context, request *DescribeModelVersionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModelVersionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeModelVersions")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelVersionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFormatAndFrameworksRequest() (request *DescribeFormatAndFrameworksRequest) {
	request = &DescribeFormatAndFrameworksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeFormatAndFrameworks")
	return
}

func NewDescribeFormatAndFrameworksResponse() (response *DescribeFormatAndFrameworksResponse) {
	response = &DescribeFormatAndFrameworksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFormatAndFrameworks(request *DescribeFormatAndFrameworksRequest) string {
	return c.DescribeFormatAndFrameworksWithContext(context.Background(), request)
}

func (c *Client) DescribeFormatAndFrameworksSend(request *DescribeFormatAndFrameworksRequest) (*DescribeFormatAndFrameworksResponse, error) {
	statusCode, msg, err := c.DescribeFormatAndFrameworksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFormatAndFrameworksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFormatAndFrameworksWithContext(ctx context.Context, request *DescribeFormatAndFrameworksRequest) string {
	if request == nil {
		request = NewDescribeFormatAndFrameworksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeFormatAndFrameworks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFormatAndFrameworksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFormatAndFrameworksWithContextV2(ctx context.Context, request *DescribeFormatAndFrameworksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFormatAndFrameworksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeFormatAndFrameworks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFormatAndFrameworksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
