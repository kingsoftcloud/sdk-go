package v20221222

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2022-12-22"

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

func NewQueryInstanceConsumeRequest() (request *QueryInstanceConsumeRequest) {
	request = &QueryInstanceConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryInstanceConsume")
	return
}

func NewQueryInstanceConsumeResponse() (response *QueryInstanceConsumeResponse) {
	response = &QueryInstanceConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryInstanceConsume(request *QueryInstanceConsumeRequest) string {
	return c.QueryInstanceConsumeWithContext(context.Background(), request)
}

func (c *Client) QueryInstanceConsumeSend(request *QueryInstanceConsumeRequest) (*QueryInstanceConsumeResponse, error) {
	statusCode, msg, err := c.QueryInstanceConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryInstanceConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryInstanceConsumeWithContext(ctx context.Context, request *QueryInstanceConsumeRequest) string {
	if request == nil {
		request = NewQueryInstanceConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryInstanceConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryInstanceConsumeWithContextV2(ctx context.Context, request *QueryInstanceConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewQueryInstanceConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryInstanceConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryProjectConsumeRequest() (request *QueryProjectConsumeRequest) {
	request = &QueryProjectConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryProjectConsume")
	return
}

func NewQueryProjectConsumeResponse() (response *QueryProjectConsumeResponse) {
	response = &QueryProjectConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryProjectConsume(request *QueryProjectConsumeRequest) string {
	return c.QueryProjectConsumeWithContext(context.Background(), request)
}

func (c *Client) QueryProjectConsumeSend(request *QueryProjectConsumeRequest) (*QueryProjectConsumeResponse, error) {
	statusCode, msg, err := c.QueryProjectConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryProjectConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryProjectConsumeWithContext(ctx context.Context, request *QueryProjectConsumeRequest) string {
	if request == nil {
		request = NewQueryProjectConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryProjectConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryProjectConsumeWithContextV2(ctx context.Context, request *QueryProjectConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewQueryProjectConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryProjectConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryProductConsumeRequest() (request *QueryProductConsumeRequest) {
	request = &QueryProductConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryProductConsume")
	return
}

func NewQueryProductConsumeResponse() (response *QueryProductConsumeResponse) {
	response = &QueryProductConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryProductConsume(request *QueryProductConsumeRequest) string {
	return c.QueryProductConsumeWithContext(context.Background(), request)
}

func (c *Client) QueryProductConsumeSend(request *QueryProductConsumeRequest) (*QueryProductConsumeResponse, error) {
	statusCode, msg, err := c.QueryProductConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryProductConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryProductConsumeWithContext(ctx context.Context, request *QueryProductConsumeRequest) string {
	if request == nil {
		request = NewQueryProductConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryProductConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryProductConsumeWithContextV2(ctx context.Context, request *QueryProductConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewQueryProductConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryProductConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryFinanceUnitConsumeRequest() (request *QueryFinanceUnitConsumeRequest) {
	request = &QueryFinanceUnitConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryFinanceUnitConsume")
	return
}

func NewQueryFinanceUnitConsumeResponse() (response *QueryFinanceUnitConsumeResponse) {
	response = &QueryFinanceUnitConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryFinanceUnitConsume(request *QueryFinanceUnitConsumeRequest) string {
	return c.QueryFinanceUnitConsumeWithContext(context.Background(), request)
}

func (c *Client) QueryFinanceUnitConsumeSend(request *QueryFinanceUnitConsumeRequest) (*QueryFinanceUnitConsumeResponse, error) {
	statusCode, msg, err := c.QueryFinanceUnitConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryFinanceUnitConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryFinanceUnitConsumeWithContext(ctx context.Context, request *QueryFinanceUnitConsumeRequest) string {
	if request == nil {
		request = NewQueryFinanceUnitConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryFinanceUnitConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryFinanceUnitConsumeWithContextV2(ctx context.Context, request *QueryFinanceUnitConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewQueryFinanceUnitConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryFinanceUnitConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryFinanceUnitConsumeOfMonthRequest() (request *QueryFinanceUnitConsumeOfMonthRequest) {
	request = &QueryFinanceUnitConsumeOfMonthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryFinanceUnitConsumeOfMonth")
	return
}

func NewQueryFinanceUnitConsumeOfMonthResponse() (response *QueryFinanceUnitConsumeOfMonthResponse) {
	response = &QueryFinanceUnitConsumeOfMonthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryFinanceUnitConsumeOfMonth(request *QueryFinanceUnitConsumeOfMonthRequest) string {
	return c.QueryFinanceUnitConsumeOfMonthWithContext(context.Background(), request)
}

func (c *Client) QueryFinanceUnitConsumeOfMonthSend(request *QueryFinanceUnitConsumeOfMonthRequest) (*QueryFinanceUnitConsumeOfMonthResponse, error) {
	statusCode, msg, err := c.QueryFinanceUnitConsumeOfMonthWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryFinanceUnitConsumeOfMonthResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryFinanceUnitConsumeOfMonthWithContext(ctx context.Context, request *QueryFinanceUnitConsumeOfMonthRequest) string {
	if request == nil {
		request = NewQueryFinanceUnitConsumeOfMonthRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryFinanceUnitConsumeOfMonthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryFinanceUnitConsumeOfMonthWithContextV2(ctx context.Context, request *QueryFinanceUnitConsumeOfMonthRequest) (int, string, error) {
	if request == nil {
		request = NewQueryFinanceUnitConsumeOfMonthRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryFinanceUnitConsumeOfMonthResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryUserConsumeRequest() (request *QueryUserConsumeRequest) {
	request = &QueryUserConsumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryUserConsume")
	return
}

func NewQueryUserConsumeResponse() (response *QueryUserConsumeResponse) {
	response = &QueryUserConsumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryUserConsume(request *QueryUserConsumeRequest) string {
	return c.QueryUserConsumeWithContext(context.Background(), request)
}

func (c *Client) QueryUserConsumeSend(request *QueryUserConsumeRequest) (*QueryUserConsumeResponse, error) {
	statusCode, msg, err := c.QueryUserConsumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryUserConsumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryUserConsumeWithContext(ctx context.Context, request *QueryUserConsumeRequest) string {
	if request == nil {
		request = NewQueryUserConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryUserConsumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryUserConsumeWithContextV2(ctx context.Context, request *QueryUserConsumeRequest) (int, string, error) {
	if request == nil {
		request = NewQueryUserConsumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryUserConsumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
