package v20221222

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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
