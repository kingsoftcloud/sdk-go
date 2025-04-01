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
