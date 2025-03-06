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
