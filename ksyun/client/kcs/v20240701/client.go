package v20240701

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-07-01"

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

func NewDescribeCacheByRoleRequest() (request *DescribeCacheByRoleRequest) {
	request = &DescribeCacheByRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheByRole")
	return
}

func NewDescribeCacheByRoleResponse() (response *DescribeCacheByRoleResponse) {
	response = &DescribeCacheByRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheByRole(request *DescribeCacheByRoleRequest) string {
	return c.DescribeCacheByRoleWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheByRoleSend(request *DescribeCacheByRoleRequest) (*DescribeCacheByRoleResponse, error) {
	statusCode, msg, err := c.DescribeCacheByRoleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCacheByRoleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheByRoleWithContext(ctx context.Context, request *DescribeCacheByRoleRequest) string {
	if request == nil {
		request = NewDescribeCacheByRoleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheByRole")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheByRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheByRoleWithContextV2(ctx context.Context, request *DescribeCacheByRoleRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheByRoleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheByRole")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheByRoleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
