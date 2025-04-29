package v20250101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-01-01"

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

func NewDescribeAlertingResourcesRequest() (request *DescribeAlertingResourcesRequest) {
	request = &DescribeAlertingResourcesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertingResources")
	return
}

func NewDescribeAlertingResourcesResponse() (response *DescribeAlertingResourcesResponse) {
	response = &DescribeAlertingResourcesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlertingResources(request *DescribeAlertingResourcesRequest) string {
	return c.DescribeAlertingResourcesWithContext(context.Background(), request)
}

func (c *Client) DescribeAlertingResourcesWithContext(ctx context.Context, request *DescribeAlertingResourcesRequest) string {
	if request == nil {
		request = NewDescribeAlertingResourcesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAlertingResourcesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
