package v20250220

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-02-20"

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

func NewListInstanceSupportBillTypesRequest() (request *ListInstanceSupportBillTypesRequest) {
	request = &ListInstanceSupportBillTypesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "ListInstanceSupportBillTypes")
	return
}

func NewListInstanceSupportBillTypesResponse() (response *ListInstanceSupportBillTypesResponse) {
	response = &ListInstanceSupportBillTypesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstanceSupportBillTypes(request *ListInstanceSupportBillTypesRequest) string {
	return c.ListInstanceSupportBillTypesWithContext(context.Background(), request)
}

func (c *Client) ListInstanceSupportBillTypesSend(request *ListInstanceSupportBillTypesRequest) (*ListInstanceSupportBillTypesResponse, error) {
	statusCode, msg, err := c.ListInstanceSupportBillTypesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListInstanceSupportBillTypesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListInstanceSupportBillTypesWithContext(ctx context.Context, request *ListInstanceSupportBillTypesRequest) string {
	if request == nil {
		request = NewListInstanceSupportBillTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceSupportBillTypesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListInstanceSupportBillTypesWithContextV2(ctx context.Context, request *ListInstanceSupportBillTypesRequest) (int, string, error) {
	if request == nil {
		request = NewListInstanceSupportBillTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceSupportBillTypesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddTrialToBuyTaskRequest() (request *AddTrialToBuyTaskRequest) {
	request = &AddTrialToBuyTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "AddTrialToBuyTask")
	return
}

func NewAddTrialToBuyTaskResponse() (response *AddTrialToBuyTaskResponse) {
	response = &AddTrialToBuyTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddTrialToBuyTask(request *AddTrialToBuyTaskRequest) string {
	return c.AddTrialToBuyTaskWithContext(context.Background(), request)
}

func (c *Client) AddTrialToBuyTaskSend(request *AddTrialToBuyTaskRequest) (*AddTrialToBuyTaskResponse, error) {
	statusCode, msg, err := c.AddTrialToBuyTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddTrialToBuyTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddTrialToBuyTaskWithContext(ctx context.Context, request *AddTrialToBuyTaskRequest) string {
	if request == nil {
		request = NewAddTrialToBuyTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddTrialToBuyTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddTrialToBuyTaskWithContextV2(ctx context.Context, request *AddTrialToBuyTaskRequest) (int, string, error) {
	if request == nil {
		request = NewAddTrialToBuyTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddTrialToBuyTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteTrialToBuyTaskRequest() (request *DeleteTrialToBuyTaskRequest) {
	request = &DeleteTrialToBuyTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "DeleteTrialToBuyTask")
	return
}

func NewDeleteTrialToBuyTaskResponse() (response *DeleteTrialToBuyTaskResponse) {
	response = &DeleteTrialToBuyTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTrialToBuyTask(request *DeleteTrialToBuyTaskRequest) string {
	return c.DeleteTrialToBuyTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteTrialToBuyTaskSend(request *DeleteTrialToBuyTaskRequest) (*DeleteTrialToBuyTaskResponse, error) {
	statusCode, msg, err := c.DeleteTrialToBuyTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteTrialToBuyTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteTrialToBuyTaskWithContext(ctx context.Context, request *DeleteTrialToBuyTaskRequest) string {
	if request == nil {
		request = NewDeleteTrialToBuyTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteTrialToBuyTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteTrialToBuyTaskWithContextV2(ctx context.Context, request *DeleteTrialToBuyTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteTrialToBuyTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteTrialToBuyTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateTrialToBuyNowRequest() (request *CreateTrialToBuyNowRequest) {
	request = &CreateTrialToBuyNowRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "CreateTrialToBuyNow")
	return
}

func NewCreateTrialToBuyNowResponse() (response *CreateTrialToBuyNowResponse) {
	response = &CreateTrialToBuyNowResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTrialToBuyNow(request *CreateTrialToBuyNowRequest) string {
	return c.CreateTrialToBuyNowWithContext(context.Background(), request)
}

func (c *Client) CreateTrialToBuyNowSend(request *CreateTrialToBuyNowRequest) (*CreateTrialToBuyNowResponse, error) {
	statusCode, msg, err := c.CreateTrialToBuyNowWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateTrialToBuyNowResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTrialToBuyNowWithContext(ctx context.Context, request *CreateTrialToBuyNowRequest) string {
	if request == nil {
		request = NewCreateTrialToBuyNowRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrialToBuyNowResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTrialToBuyNowWithContextV2(ctx context.Context, request *CreateTrialToBuyNowRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTrialToBuyNowRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrialToBuyNowResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
