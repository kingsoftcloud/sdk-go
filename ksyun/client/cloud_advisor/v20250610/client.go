package v20250610

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-06-10"

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

func NewGetReportRequest() (request *GetReportRequest) {
	request = &GetReportRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloud-advisor", APIVersion, "GetReport")
	return
}

func NewGetReportResponse() (response *GetReportResponse) {
	response = &GetReportResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetReport(request *GetReportRequest) string {
	return c.GetReportWithContext(context.Background(), request)
}

func (c *Client) GetReportSend(request *GetReportRequest) (*GetReportResponse, error) {
	statusCode, msg, err := c.GetReportWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetReportResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetReportWithContext(ctx context.Context, request *GetReportRequest) string {
	if request == nil {
		request = NewGetReportRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cloud-advisor", APIVersion, "GetReport")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetReportResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetReportWithContextV2(ctx context.Context, request *GetReportRequest) (int, string, error) {
	if request == nil {
		request = NewGetReportRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cloud-advisor", APIVersion, "GetReport")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetReportResponse()
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
	request.Init().WithApiInfo("cloud-advisor", APIVersion, "CreateTask")
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
		request.Init().WithApiInfo("cloud-advisor", APIVersion, "CreateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

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
		request.Init().WithApiInfo("cloud-advisor", APIVersion, "CreateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListInspectionItemRequest() (request *ListInspectionItemRequest) {
	request = &ListInspectionItemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloud-advisor", APIVersion, "ListInspectionItem")
	return
}

func NewListInspectionItemResponse() (response *ListInspectionItemResponse) {
	response = &ListInspectionItemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInspectionItem(request *ListInspectionItemRequest) string {
	return c.ListInspectionItemWithContext(context.Background(), request)
}

func (c *Client) ListInspectionItemSend(request *ListInspectionItemRequest) (*ListInspectionItemResponse, error) {
	statusCode, msg, err := c.ListInspectionItemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListInspectionItemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListInspectionItemWithContext(ctx context.Context, request *ListInspectionItemRequest) string {
	if request == nil {
		request = NewListInspectionItemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cloud-advisor", APIVersion, "ListInspectionItem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInspectionItemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListInspectionItemWithContextV2(ctx context.Context, request *ListInspectionItemRequest) (int, string, error) {
	if request == nil {
		request = NewListInspectionItemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cloud-advisor", APIVersion, "ListInspectionItem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInspectionItemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
