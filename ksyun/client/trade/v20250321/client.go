package v20250321

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-03-21"

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

func NewQueryUnPayOrdersRequest() (request *QueryUnPayOrdersRequest) {
	request = &QueryUnPayOrdersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "QueryUnPayOrders")
	return
}

func NewQueryUnPayOrdersResponse() (response *QueryUnPayOrdersResponse) {
	response = &QueryUnPayOrdersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryUnPayOrders(request *QueryUnPayOrdersRequest) string {
	return c.QueryUnPayOrdersWithContext(context.Background(), request)
}

func (c *Client) QueryUnPayOrdersSend(request *QueryUnPayOrdersRequest) (*QueryUnPayOrdersResponse, error) {
	statusCode, msg, err := c.QueryUnPayOrdersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryUnPayOrdersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryUnPayOrdersWithContext(ctx context.Context, request *QueryUnPayOrdersRequest) string {
	if request == nil {
		request = NewQueryUnPayOrdersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryUnPayOrdersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryUnPayOrdersWithContextV2(ctx context.Context, request *QueryUnPayOrdersRequest) (int, string, error) {
	if request == nil {
		request = NewQueryUnPayOrdersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryUnPayOrdersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryOrderInfoRequest() (request *QueryOrderInfoRequest) {
	request = &QueryOrderInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "QueryOrderInfo")
	return
}

func NewQueryOrderInfoResponse() (response *QueryOrderInfoResponse) {
	response = &QueryOrderInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryOrderInfo(request *QueryOrderInfoRequest) string {
	return c.QueryOrderInfoWithContext(context.Background(), request)
}

func (c *Client) QueryOrderInfoSend(request *QueryOrderInfoRequest) (*QueryOrderInfoResponse, error) {
	statusCode, msg, err := c.QueryOrderInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryOrderInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryOrderInfoWithContext(ctx context.Context, request *QueryOrderInfoRequest) string {
	if request == nil {
		request = NewQueryOrderInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryOrderInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryOrderInfoWithContextV2(ctx context.Context, request *QueryOrderInfoRequest) (int, string, error) {
	if request == nil {
		request = NewQueryOrderInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryOrderInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelOrderRequest() (request *CancelOrderRequest) {
	request = &CancelOrderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "CancelOrder")
	return
}

func NewCancelOrderResponse() (response *CancelOrderResponse) {
	response = &CancelOrderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelOrder(request *CancelOrderRequest) string {
	return c.CancelOrderWithContext(context.Background(), request)
}

func (c *Client) CancelOrderSend(request *CancelOrderRequest) (*CancelOrderResponse, error) {
	statusCode, msg, err := c.CancelOrderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CancelOrderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelOrderWithContext(ctx context.Context, request *CancelOrderRequest) string {
	if request == nil {
		request = NewCancelOrderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelOrderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelOrderWithContextV2(ctx context.Context, request *CancelOrderRequest) (int, string, error) {
	if request == nil {
		request = NewCancelOrderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelOrderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewLaunchPayOrderRequest() (request *LaunchPayOrderRequest) {
	request = &LaunchPayOrderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "LaunchPayOrder")
	return
}

func NewLaunchPayOrderResponse() (response *LaunchPayOrderResponse) {
	response = &LaunchPayOrderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) LaunchPayOrder(request *LaunchPayOrderRequest) string {
	return c.LaunchPayOrderWithContext(context.Background(), request)
}

func (c *Client) LaunchPayOrderSend(request *LaunchPayOrderRequest) (*LaunchPayOrderResponse, error) {
	statusCode, msg, err := c.LaunchPayOrderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct LaunchPayOrderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) LaunchPayOrderWithContext(ctx context.Context, request *LaunchPayOrderRequest) string {
	if request == nil {
		request = NewLaunchPayOrderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewLaunchPayOrderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) LaunchPayOrderWithContextV2(ctx context.Context, request *LaunchPayOrderRequest) (int, string, error) {
	if request == nil {
		request = NewLaunchPayOrderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewLaunchPayOrderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
