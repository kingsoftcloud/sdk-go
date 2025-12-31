package v20250828

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-08-28"

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

func NewQueryInstancesRequest() (request *QueryInstancesRequest) {
	request = &QueryInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "QueryInstances")
	return
}

func NewQueryInstancesResponse() (response *QueryInstancesResponse) {
	response = &QueryInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryInstances(request *QueryInstancesRequest) string {
	return c.QueryInstancesWithContext(context.Background(), request)
}

func (c *Client) QueryInstancesSend(request *QueryInstancesRequest) (*QueryInstancesResponse, error) {
	statusCode, msg, err := c.QueryInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryInstancesWithContext(ctx context.Context, request *QueryInstancesRequest) string {
	if request == nil {
		request = NewQueryInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("trade", APIVersion, "QueryInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryInstancesWithContextV2(ctx context.Context, request *QueryInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewQueryInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("trade", APIVersion, "QueryInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
