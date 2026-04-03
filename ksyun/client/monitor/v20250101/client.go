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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertingResources")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertingResources")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "DescribeSystemEventAttributes")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "DescribeSystemEventAttributes")
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
func NewListAlarmEffectInstanceRequest() (request *ListAlarmEffectInstanceRequest) {
	request = &ListAlarmEffectInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "ListAlarmEffectInstance")
	return
}

func NewListAlarmEffectInstanceResponse() (response *ListAlarmEffectInstanceResponse) {
	response = &ListAlarmEffectInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListAlarmEffectInstance(request *ListAlarmEffectInstanceRequest) string {
	return c.ListAlarmEffectInstanceWithContext(context.Background(), request)
}

func (c *Client) ListAlarmEffectInstanceSend(request *ListAlarmEffectInstanceRequest) (*ListAlarmEffectInstanceResponse, error) {
	statusCode, msg, err := c.ListAlarmEffectInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListAlarmEffectInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListAlarmEffectInstanceWithContext(ctx context.Context, request *ListAlarmEffectInstanceRequest) string {
	if request == nil {
		request = NewListAlarmEffectInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "ListAlarmEffectInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAlarmEffectInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListAlarmEffectInstanceWithContextV2(ctx context.Context, request *ListAlarmEffectInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewListAlarmEffectInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "ListAlarmEffectInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAlarmEffectInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPrometheusTokenRequest() (request *GetPrometheusTokenRequest) {
	request = &GetPrometheusTokenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetPrometheusToken")
	return
}

func NewGetPrometheusTokenResponse() (response *GetPrometheusTokenResponse) {
	response = &GetPrometheusTokenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPrometheusToken(request *GetPrometheusTokenRequest) string {
	return c.GetPrometheusTokenWithContext(context.Background(), request)
}

func (c *Client) GetPrometheusTokenSend(request *GetPrometheusTokenRequest) (*GetPrometheusTokenResponse, error) {
	statusCode, msg, err := c.GetPrometheusTokenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPrometheusTokenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPrometheusTokenWithContext(ctx context.Context, request *GetPrometheusTokenRequest) string {
	if request == nil {
		request = NewGetPrometheusTokenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "GetPrometheusToken")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPrometheusTokenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPrometheusTokenWithContextV2(ctx context.Context, request *GetPrometheusTokenRequest) (int, string, error) {
	if request == nil {
		request = NewGetPrometheusTokenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "GetPrometheusToken")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPrometheusTokenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewPutDefaultEventPolicyRequest() (request *PutDefaultEventPolicyRequest) {
	request = &PutDefaultEventPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "PutDefaultEventPolicy")
	return
}

func NewPutDefaultEventPolicyResponse() (response *PutDefaultEventPolicyResponse) {
	response = &PutDefaultEventPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) PutDefaultEventPolicy(request *PutDefaultEventPolicyRequest) string {
	return c.PutDefaultEventPolicyWithContext(context.Background(), request)
}

func (c *Client) PutDefaultEventPolicySend(request *PutDefaultEventPolicyRequest) (*PutDefaultEventPolicyResponse, error) {
	statusCode, msg, err := c.PutDefaultEventPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct PutDefaultEventPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) PutDefaultEventPolicyWithContext(ctx context.Context, request *PutDefaultEventPolicyRequest) string {
	if request == nil {
		request = NewPutDefaultEventPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "PutDefaultEventPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewPutDefaultEventPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) PutDefaultEventPolicyWithContextV2(ctx context.Context, request *PutDefaultEventPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewPutDefaultEventPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "PutDefaultEventPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewPutDefaultEventPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
