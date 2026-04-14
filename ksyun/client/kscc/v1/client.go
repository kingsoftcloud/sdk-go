package v1

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "V1"

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

func NewUpdateUserQuotaRequest() (request *UpdateUserQuotaRequest) {
	request = &UpdateUserQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "UpdateUserQuota")
	return
}

func NewUpdateUserQuotaResponse() (response *UpdateUserQuotaResponse) {
	response = &UpdateUserQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateUserQuota(request *UpdateUserQuotaRequest) string {
	return c.UpdateUserQuotaWithContext(context.Background(), request)
}

func (c *Client) UpdateUserQuotaSend(request *UpdateUserQuotaRequest) (*UpdateUserQuotaResponse, error) {
	statusCode, msg, err := c.UpdateUserQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateUserQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateUserQuotaWithContext(ctx context.Context, request *UpdateUserQuotaRequest) string {
	if request == nil {
		request = NewUpdateUserQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateUserQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateUserQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateUserQuotaWithContextV2(ctx context.Context, request *UpdateUserQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateUserQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateUserQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateUserQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeUserCostSummaryRequest() (request *DescribeUserCostSummaryRequest) {
	request = &DescribeUserCostSummaryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserCostSummary")
	return
}

func NewDescribeUserCostSummaryResponse() (response *DescribeUserCostSummaryResponse) {
	response = &DescribeUserCostSummaryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeUserCostSummary(request *DescribeUserCostSummaryRequest) string {
	return c.DescribeUserCostSummaryWithContext(context.Background(), request)
}

func (c *Client) DescribeUserCostSummarySend(request *DescribeUserCostSummaryRequest) (*DescribeUserCostSummaryResponse, error) {
	statusCode, msg, err := c.DescribeUserCostSummaryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeUserCostSummaryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeUserCostSummaryWithContext(ctx context.Context, request *DescribeUserCostSummaryRequest) string {
	if request == nil {
		request = NewDescribeUserCostSummaryRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserCostSummary")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserCostSummaryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeUserCostSummaryWithContextV2(ctx context.Context, request *DescribeUserCostSummaryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeUserCostSummaryRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserCostSummary")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserCostSummaryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
