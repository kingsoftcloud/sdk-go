package v20211201

import (
	"context"
	"fmt"

	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-12-01"

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

func NewGetRefreshOrPreloadTaskRequest() (request *GetRefreshOrPreloadTaskRequest) {
	request = &GetRefreshOrPreloadTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetRefreshOrPreloadTask")
	return
}

func NewGetRefreshOrPreloadTaskResponse() (response *GetRefreshOrPreloadTaskResponse) {
	response = &GetRefreshOrPreloadTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRefreshOrPreloadTask(request *GetRefreshOrPreloadTaskRequest) string {
	return c.GetRefreshOrPreloadTaskWithContext(context.Background(), request)
}

func (c *Client) GetRefreshOrPreloadTaskSend(request *GetRefreshOrPreloadTaskRequest) (*GetRefreshOrPreloadTaskResponse, error) {
	statusCode, msg, err := c.GetRefreshOrPreloadTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRefreshOrPreloadTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRefreshOrPreloadTaskWithContext(ctx context.Context, request *GetRefreshOrPreloadTaskRequest) string {
	if request == nil {
		request = NewGetRefreshOrPreloadTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetRefreshOrPreloadTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRefreshOrPreloadTaskWithContextV2(ctx context.Context, request *GetRefreshOrPreloadTaskRequest) (int, string, error) {
	if request == nil {
		request = NewGetRefreshOrPreloadTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetRefreshOrPreloadTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
