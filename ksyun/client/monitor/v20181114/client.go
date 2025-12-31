package v20181114

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-11-14"

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

func NewGetMetricStatisticsBatchRequest() (request *GetMetricStatisticsBatchRequest) {
	request = &GetMetricStatisticsBatchRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetMetricStatisticsBatch")
	return
}

func NewGetMetricStatisticsBatchResponse() (response *GetMetricStatisticsBatchResponse) {
	response = &GetMetricStatisticsBatchResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMetricStatisticsBatch(request *GetMetricStatisticsBatchRequest) string {
	return c.GetMetricStatisticsBatchWithContext(context.Background(), request)
}

func (c *Client) GetMetricStatisticsBatchSend(request *GetMetricStatisticsBatchRequest) (*GetMetricStatisticsBatchResponse, error) {
	statusCode, msg, err := c.GetMetricStatisticsBatchWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetMetricStatisticsBatchResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetMetricStatisticsBatchWithContext(ctx context.Context, request *GetMetricStatisticsBatchRequest) string {
	if request == nil {
		request = NewGetMetricStatisticsBatchRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "GetMetricStatisticsBatch")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMetricStatisticsBatchResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetMetricStatisticsBatchWithContextV2(ctx context.Context, request *GetMetricStatisticsBatchRequest) (int, string, error) {
	if request == nil {
		request = NewGetMetricStatisticsBatchRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "GetMetricStatisticsBatch")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMetricStatisticsBatchResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
