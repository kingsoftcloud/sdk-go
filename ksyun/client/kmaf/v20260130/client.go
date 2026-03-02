package v20260130

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2026-01-30"

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

func NewQueryAnswerRequest() (request *QueryAnswerRequest) {
	request = &QueryAnswerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmaf", APIVersion, "QueryAnswer")
	return
}

func NewQueryAnswerResponse() (response *QueryAnswerResponse) {
	response = &QueryAnswerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryAnswer(request *QueryAnswerRequest) string {
	return c.QueryAnswerWithContext(context.Background(), request)
}

func (c *Client) QueryAnswerSend(request *QueryAnswerRequest) (*QueryAnswerResponse, error) {
	statusCode, msg, err := c.QueryAnswerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryAnswerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryAnswerWithContext(ctx context.Context, request *QueryAnswerRequest) string {
	if request == nil {
		request = NewQueryAnswerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kmaf", APIVersion, "QueryAnswer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryAnswerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryAnswerWithContextV2(ctx context.Context, request *QueryAnswerRequest) (int, string, error) {
	if request == nil {
		request = NewQueryAnswerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kmaf", APIVersion, "QueryAnswer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryAnswerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCheckModerateRequest() (request *CheckModerateRequest) {
	request = &CheckModerateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmaf", APIVersion, "CheckModerate")
	return
}

func NewCheckModerateResponse() (response *CheckModerateResponse) {
	response = &CheckModerateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckModerate(request *CheckModerateRequest) string {
	return c.CheckModerateWithContext(context.Background(), request)
}

func (c *Client) CheckModerateSend(request *CheckModerateRequest) (*CheckModerateResponse, error) {
	statusCode, msg, err := c.CheckModerateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CheckModerateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CheckModerateWithContext(ctx context.Context, request *CheckModerateRequest) string {
	if request == nil {
		request = NewCheckModerateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kmaf", APIVersion, "CheckModerate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckModerateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CheckModerateWithContextV2(ctx context.Context, request *CheckModerateRequest) (int, string, error) {
	if request == nil {
		request = NewCheckModerateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kmaf", APIVersion, "CheckModerate")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckModerateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
