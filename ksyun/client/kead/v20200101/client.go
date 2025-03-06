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

func (c *Client) DescribeKeadWithContext(ctx context.Context, request *DescribeKeadRequest) string {
	if request == nil {
		request = NewDescribeKeadRequest()
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

func (c *Client) DescribeKeadIpWithContext(ctx context.Context, request *DescribeKeadIpRequest) string {
	if request == nil {
		request = NewDescribeKeadIpRequest()
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

func (c *Client) DescribeBlockIpWithContext(ctx context.Context, request *DescribeBlockIpRequest) string {
	if request == nil {
		request = NewDescribeBlockIpRequest()
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
