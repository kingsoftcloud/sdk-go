package v20100525

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2010-05-25"

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

func NewGetMetricStatisticsRequest() (request *GetMetricStatisticsRequest) {
	request = &GetMetricStatisticsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetMetricStatistics")
	return
}

func NewGetMetricStatisticsResponse() (response *GetMetricStatisticsResponse) {
	response = &GetMetricStatisticsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMetricStatistics(request *GetMetricStatisticsRequest) string {
	return c.GetMetricStatisticsWithContext(context.Background(), request)
}

func (c *Client) GetMetricStatisticsSend(request *GetMetricStatisticsRequest) (*GetMetricStatisticsResponse, error) {
	statusCode, msg, err := c.GetMetricStatisticsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetMetricStatisticsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetMetricStatisticsWithContext(ctx context.Context, request *GetMetricStatisticsRequest) string {
	if request == nil {
		request = NewGetMetricStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMetricStatisticsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetMetricStatisticsWithContextV2(ctx context.Context, request *GetMetricStatisticsRequest) (int, string, error) {
	if request == nil {
		request = NewGetMetricStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMetricStatisticsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListMetricsRequest() (request *ListMetricsRequest) {
	request = &ListMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "ListMetrics")
	return
}

func NewListMetricsResponse() (response *ListMetricsResponse) {
	response = &ListMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListMetrics(request *ListMetricsRequest) string {
	return c.ListMetricsWithContext(context.Background(), request)
}

func (c *Client) ListMetricsSend(request *ListMetricsRequest) (*ListMetricsResponse, error) {
	statusCode, msg, err := c.ListMetricsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListMetricsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListMetricsWithContext(ctx context.Context, request *ListMetricsRequest) string {
	if request == nil {
		request = NewListMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListMetricsWithContextV2(ctx context.Context, request *ListMetricsRequest) (int, string, error) {
	if request == nil {
		request = NewListMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListMetricsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
