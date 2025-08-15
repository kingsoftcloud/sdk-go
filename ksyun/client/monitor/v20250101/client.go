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

func (c *Client) DescribeAlertingResourcesSend(request *DescribeAlertingResourcesRequest) (*DescribeAlertingResourcesResponse, error) {
	statusCode, msg, err := c.DescribeAlertingResourcesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAlertingResourcesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeAlertingResourcesWithContextV2(ctx context.Context, request *DescribeAlertingResourcesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlertingResourcesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAlertingResourcesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSystemEventAttributesRequest() (request *DescribeSystemEventAttributesRequest) {
	request = &DescribeSystemEventAttributesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeSystemEventAttributes")
	return
}

func NewDescribeSystemEventAttributesResponse() (response *DescribeSystemEventAttributesResponse) {
	response = &DescribeSystemEventAttributesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSystemEventAttributes(request *DescribeSystemEventAttributesRequest) string {
	return c.DescribeSystemEventAttributesWithContext(context.Background(), request)
}

func (c *Client) DescribeSystemEventAttributesSend(request *DescribeSystemEventAttributesRequest) (*DescribeSystemEventAttributesResponse, error) {
	statusCode, msg, err := c.DescribeSystemEventAttributesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSystemEventAttributesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSystemEventAttributesWithContext(ctx context.Context, request *DescribeSystemEventAttributesRequest) string {
	if request == nil {
		request = NewDescribeSystemEventAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSystemEventAttributesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSystemEventAttributesWithContextV2(ctx context.Context, request *DescribeSystemEventAttributesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSystemEventAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSystemEventAttributesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
