package v20190401

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2019-04-01"

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

func NewListOperateLogsRequest() (request *ListOperateLogsRequest) {
	request = &ListOperateLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("actiontrail", APIVersion, "ListOperateLogs")
	return
}

func NewListOperateLogsResponse() (response *ListOperateLogsResponse) {
	response = &ListOperateLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListOperateLogs(request *ListOperateLogsRequest) string {
	return c.ListOperateLogsWithContext(context.Background(), request)
}

func (c *Client) ListOperateLogsSend(request *ListOperateLogsRequest) (*ListOperateLogsResponse, error) {
	statusCode, msg, err := c.ListOperateLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListOperateLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListOperateLogsWithContext(ctx context.Context, request *ListOperateLogsRequest) string {
	if request == nil {
		request = NewListOperateLogsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("actiontrail", APIVersion, "ListOperateLogs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListOperateLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListOperateLogsWithContextV2(ctx context.Context, request *ListOperateLogsRequest) (int, string, error) {
	if request == nil {
		request = NewListOperateLogsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("actiontrail", APIVersion, "ListOperateLogs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListOperateLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
