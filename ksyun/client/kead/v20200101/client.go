package v20200101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-01-01"

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

func NewDescribeKeadRequest() (request *DescribeKeadRequest) {
	request = &DescribeKeadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kead", APIVersion, "DescribeKead")
	return
}

func NewDescribeKeadResponse() (response *DescribeKeadResponse) {
	response = &DescribeKeadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKead(request *DescribeKeadRequest) string {
	return c.DescribeKeadWithContext(context.Background(), request)
}

func (c *Client) DescribeKeadSend(request *DescribeKeadRequest) (*DescribeKeadResponse, error) {
	statusCode, msg, err := c.DescribeKeadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeKeadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKeadWithContext(ctx context.Context, request *DescribeKeadRequest) string {
	if request == nil {
		request = NewDescribeKeadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kead", APIVersion, "DescribeKead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKeadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKeadWithContextV2(ctx context.Context, request *DescribeKeadRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKeadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kead", APIVersion, "DescribeKead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKeadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKeadIpRequest() (request *DescribeKeadIpRequest) {
	request = &DescribeKeadIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kead", APIVersion, "DescribeKeadIp")
	return
}

func NewDescribeKeadIpResponse() (response *DescribeKeadIpResponse) {
	response = &DescribeKeadIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKeadIp(request *DescribeKeadIpRequest) string {
	return c.DescribeKeadIpWithContext(context.Background(), request)
}

func (c *Client) DescribeKeadIpSend(request *DescribeKeadIpRequest) (*DescribeKeadIpResponse, error) {
	statusCode, msg, err := c.DescribeKeadIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeKeadIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKeadIpWithContext(ctx context.Context, request *DescribeKeadIpRequest) string {
	if request == nil {
		request = NewDescribeKeadIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kead", APIVersion, "DescribeKeadIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKeadIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKeadIpWithContextV2(ctx context.Context, request *DescribeKeadIpRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKeadIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kead", APIVersion, "DescribeKeadIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKeadIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBlockIpRequest() (request *DescribeBlockIpRequest) {
	request = &DescribeBlockIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kead", APIVersion, "DescribeBlockIp")
	return
}

func NewDescribeBlockIpResponse() (response *DescribeBlockIpResponse) {
	response = &DescribeBlockIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBlockIp(request *DescribeBlockIpRequest) string {
	return c.DescribeBlockIpWithContext(context.Background(), request)
}

func (c *Client) DescribeBlockIpSend(request *DescribeBlockIpRequest) (*DescribeBlockIpResponse, error) {
	statusCode, msg, err := c.DescribeBlockIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBlockIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBlockIpWithContext(ctx context.Context, request *DescribeBlockIpRequest) string {
	if request == nil {
		request = NewDescribeBlockIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kead", APIVersion, "DescribeBlockIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeBlockIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBlockIpWithContextV2(ctx context.Context, request *DescribeBlockIpRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBlockIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kead", APIVersion, "DescribeBlockIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeBlockIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
