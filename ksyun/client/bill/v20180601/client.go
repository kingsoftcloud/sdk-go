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

func (c *Client) GetMonthBillWithContext(ctx context.Context, request *GetMonthBillRequest) string {
	if request == nil {
		request = NewGetMonthBillRequest()
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

func (c *Client) GetPostpayDetailBillWithContext(ctx context.Context, request *GetPostpayDetailBillRequest) string {
	if request == nil {
		request = NewGetPostpayDetailBillRequest()
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

func (c *Client) GetProductCodeWithContext(ctx context.Context, request *GetProductCodeRequest) string {
	if request == nil {
		request = NewGetProductCodeRequest()
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
