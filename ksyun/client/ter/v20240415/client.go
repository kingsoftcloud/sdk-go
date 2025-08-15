package v20240415

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-04-15"

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

func NewDescribeStackOutputsRequest() (request *DescribeStackOutputsRequest) {
	request = &DescribeStackOutputsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ter", APIVersion, "DescribeStackOutputs")
	return
}

func NewDescribeStackOutputsResponse() (response *DescribeStackOutputsResponse) {
	response = &DescribeStackOutputsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeStackOutputs(request *DescribeStackOutputsRequest) string {
	return c.DescribeStackOutputsWithContext(context.Background(), request)
}

func (c *Client) DescribeStackOutputsSend(request *DescribeStackOutputsRequest) (*DescribeStackOutputsResponse, error) {
	statusCode, msg, err := c.DescribeStackOutputsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeStackOutputsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeStackOutputsWithContext(ctx context.Context, request *DescribeStackOutputsRequest) string {
	if request == nil {
		request = NewDescribeStackOutputsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStackOutputsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeStackOutputsWithContextV2(ctx context.Context, request *DescribeStackOutputsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeStackOutputsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStackOutputsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeStackEventsRequest() (request *DescribeStackEventsRequest) {
	request = &DescribeStackEventsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ter", APIVersion, "DescribeStackEvents")
	return
}

func NewDescribeStackEventsResponse() (response *DescribeStackEventsResponse) {
	response = &DescribeStackEventsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeStackEvents(request *DescribeStackEventsRequest) string {
	return c.DescribeStackEventsWithContext(context.Background(), request)
}

func (c *Client) DescribeStackEventsSend(request *DescribeStackEventsRequest) (*DescribeStackEventsResponse, error) {
	statusCode, msg, err := c.DescribeStackEventsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeStackEventsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeStackEventsWithContext(ctx context.Context, request *DescribeStackEventsRequest) string {
	if request == nil {
		request = NewDescribeStackEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStackEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeStackEventsWithContextV2(ctx context.Context, request *DescribeStackEventsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeStackEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStackEventsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ter", APIVersion, "DeleteTemplate")
	return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) string {
	return c.DeleteTemplateWithContext(context.Background(), request)
}

func (c *Client) DeleteTemplateSend(request *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	statusCode, msg, err := c.DeleteTemplateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteTemplateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteTemplateWithContext(ctx context.Context, request *DeleteTemplateRequest) string {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteTemplateWithContextV2(ctx context.Context, request *DeleteTemplateRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTemplateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTemplateVersionsRequest() (request *DescribeTemplateVersionsRequest) {
	request = &DescribeTemplateVersionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ter", APIVersion, "DescribeTemplateVersions")
	return
}

func NewDescribeTemplateVersionsResponse() (response *DescribeTemplateVersionsResponse) {
	response = &DescribeTemplateVersionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTemplateVersions(request *DescribeTemplateVersionsRequest) string {
	return c.DescribeTemplateVersionsWithContext(context.Background(), request)
}

func (c *Client) DescribeTemplateVersionsSend(request *DescribeTemplateVersionsRequest) (*DescribeTemplateVersionsResponse, error) {
	statusCode, msg, err := c.DescribeTemplateVersionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeTemplateVersionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTemplateVersionsWithContext(ctx context.Context, request *DescribeTemplateVersionsRequest) string {
	if request == nil {
		request = NewDescribeTemplateVersionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTemplateVersionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTemplateVersionsWithContextV2(ctx context.Context, request *DescribeTemplateVersionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTemplateVersionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTemplateVersionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ter", APIVersion, "DescribeTemplates")
	return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) string {
	return c.DescribeTemplatesWithContext(context.Background(), request)
}

func (c *Client) DescribeTemplatesSend(request *DescribeTemplatesRequest) (*DescribeTemplatesResponse, error) {
	statusCode, msg, err := c.DescribeTemplatesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeTemplatesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTemplatesWithContext(ctx context.Context, request *DescribeTemplatesRequest) string {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTemplatesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTemplatesWithContextV2(ctx context.Context, request *DescribeTemplatesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTemplatesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
