package v20220601

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2022-06-01"

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

func NewGetMonthConsumeRequest() (request *GetMonthConsumeRequest) {
	request = &GetMonthConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetMonthConsume")
	return
}

func NewGetMonthConsumeResponse() (response *GetMonthConsumeResponse) {
	response = &GetMonthConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMonthConsume(request *GetMonthConsumeRequest) string {
	return c.GetMonthConsumeWithContext(context.Background(), request)
}

func (c *Client) GetMonthConsumeSend(request *GetMonthConsumeRequest) (*GetMonthConsumeResponse, error) {
	statusCode, msg, err := c.GetMonthConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetMonthConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetMonthConsumeWithContext(ctx context.Context, request *GetMonthConsumeRequest) string {
	if request == nil {
		request = NewGetMonthConsumeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetMonthConsume")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMonthConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetMonthConsumeWithContextV2(ctx context.Context, request *GetMonthConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewGetMonthConsumeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetMonthConsume")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMonthConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPostpayDetailConsumeRequest() (request *GetPostpayDetailConsumeRequest) {
	request = &GetPostpayDetailConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailConsume")
	return
}

func NewGetPostpayDetailConsumeResponse() (response *GetPostpayDetailConsumeResponse) {
	response = &GetPostpayDetailConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPostpayDetailConsume(request *GetPostpayDetailConsumeRequest) string {
	return c.GetPostpayDetailConsumeWithContext(context.Background(), request)
}

func (c *Client) GetPostpayDetailConsumeSend(request *GetPostpayDetailConsumeRequest) (*GetPostpayDetailConsumeResponse, error) {
	statusCode, msg, err := c.GetPostpayDetailConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPostpayDetailConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPostpayDetailConsumeWithContext(ctx context.Context, request *GetPostpayDetailConsumeRequest) string {
	if request == nil {
		request = NewGetPostpayDetailConsumeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailConsume")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetPostpayDetailConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPostpayDetailConsumeWithContextV2(ctx context.Context, request *GetPostpayDetailConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewGetPostpayDetailConsumeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailConsume")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetPostpayDetailConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
