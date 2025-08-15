package v20240117

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-01-17"

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

func NewGetDwsuMetricRequest() (request *GetDwsuMetricRequest) {
	request = &GetDwsuMetricRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("relyt", APIVersion, "GetDwsuMetric")
	return
}

func NewGetDwsuMetricResponse() (response *GetDwsuMetricResponse) {
	response = &GetDwsuMetricResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDwsuMetric(request *GetDwsuMetricRequest) string {
	return c.GetDwsuMetricWithContext(context.Background(), request)
}

func (c *Client) GetDwsuMetricSend(request *GetDwsuMetricRequest) (*GetDwsuMetricResponse, error) {
	statusCode, msg, err := c.GetDwsuMetricWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDwsuMetricResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDwsuMetricWithContext(ctx context.Context, request *GetDwsuMetricRequest) string {
	if request == nil {
		request = NewGetDwsuMetricRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDwsuMetricResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDwsuMetricWithContextV2(ctx context.Context, request *GetDwsuMetricRequest) (int, string, error) {
	if request == nil {
		request = NewGetDwsuMetricRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDwsuMetricResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
