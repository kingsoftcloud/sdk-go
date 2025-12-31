package v20180601

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-06-01"

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

func NewGetMonthBillRequest() (request *GetMonthBillRequest) {
	request = &GetMonthBillRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetMonthBill")
	return
}

func NewGetMonthBillResponse() (response *GetMonthBillResponse) {
	response = &GetMonthBillResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMonthBill(request *GetMonthBillRequest) string {
	return c.GetMonthBillWithContext(context.Background(), request)
}

func (c *Client) GetMonthBillSend(request *GetMonthBillRequest) (*GetMonthBillResponse, error) {
	statusCode, msg, err := c.GetMonthBillWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetMonthBillResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetMonthBillWithContext(ctx context.Context, request *GetMonthBillRequest) string {
	if request == nil {
		request = NewGetMonthBillRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetMonthBill")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMonthBillResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetMonthBillWithContextV2(ctx context.Context, request *GetMonthBillRequest) (int, string, error) {
	if request == nil {
		request = NewGetMonthBillRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetMonthBill")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMonthBillResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPostpayDetailBillRequest() (request *GetPostpayDetailBillRequest) {
	request = &GetPostpayDetailBillRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailBill")
	return
}

func NewGetPostpayDetailBillResponse() (response *GetPostpayDetailBillResponse) {
	response = &GetPostpayDetailBillResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPostpayDetailBill(request *GetPostpayDetailBillRequest) string {
	return c.GetPostpayDetailBillWithContext(context.Background(), request)
}

func (c *Client) GetPostpayDetailBillSend(request *GetPostpayDetailBillRequest) (*GetPostpayDetailBillResponse, error) {
	statusCode, msg, err := c.GetPostpayDetailBillWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPostpayDetailBillResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPostpayDetailBillWithContext(ctx context.Context, request *GetPostpayDetailBillRequest) string {
	if request == nil {
		request = NewGetPostpayDetailBillRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailBill")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetPostpayDetailBillResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPostpayDetailBillWithContextV2(ctx context.Context, request *GetPostpayDetailBillRequest) (int, string, error) {
	if request == nil {
		request = NewGetPostpayDetailBillRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailBill")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetPostpayDetailBillResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetProductCodeRequest() (request *GetProductCodeRequest) {
	request = &GetProductCodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetProductCode")
	return
}

func NewGetProductCodeResponse() (response *GetProductCodeResponse) {
	response = &GetProductCodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetProductCode(request *GetProductCodeRequest) string {
	return c.GetProductCodeWithContext(context.Background(), request)
}

func (c *Client) GetProductCodeSend(request *GetProductCodeRequest) (*GetProductCodeResponse, error) {
	statusCode, msg, err := c.GetProductCodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetProductCodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetProductCodeWithContext(ctx context.Context, request *GetProductCodeRequest) string {
	if request == nil {
		request = NewGetProductCodeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetProductCode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetProductCodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetProductCodeWithContextV2(ctx context.Context, request *GetProductCodeRequest) (int, string, error) {
	if request == nil {
		request = NewGetProductCodeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill", APIVersion, "GetProductCode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetProductCodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
