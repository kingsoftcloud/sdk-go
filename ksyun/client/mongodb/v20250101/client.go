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

func NewDescribeDefaultParamsRequest() (request *DescribeDefaultParamsRequest) {
	request = &DescribeDefaultParamsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDefaultParams")
	return
}

func NewDescribeDefaultParamsResponse() (response *DescribeDefaultParamsResponse) {
	response = &DescribeDefaultParamsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDefaultParams(request *DescribeDefaultParamsRequest) string {
	return c.DescribeDefaultParamsWithContext(context.Background(), request)
}

func (c *Client) DescribeDefaultParamsSend(request *DescribeDefaultParamsRequest) (*DescribeDefaultParamsResponse, error) {
	statusCode, msg, err := c.DescribeDefaultParamsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDefaultParamsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDefaultParamsWithContext(ctx context.Context, request *DescribeDefaultParamsRequest) string {
	if request == nil {
		request = NewDescribeDefaultParamsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDefaultParams")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDefaultParamsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDefaultParamsWithContextV2(ctx context.Context, request *DescribeDefaultParamsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDefaultParamsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDefaultParams")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDefaultParamsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
