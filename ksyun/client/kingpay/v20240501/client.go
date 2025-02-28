package v20240501

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func NewQueryCashWalletActionRequest() (request *QueryCashWalletActionRequest) {
	request = &QueryCashWalletActionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kingpay", APIVersion, "QueryCashWalletAction")
	return
}

func NewQueryCashWalletActionResponse() (response *QueryCashWalletActionResponse) {
	response = &QueryCashWalletActionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryCashWalletAction(request *QueryCashWalletActionRequest) string {
	return c.QueryCashWalletActionWithContext(context.Background(), request)
}

func (c *Client) QueryCashWalletActionWithContext(ctx context.Context, request *QueryCashWalletActionRequest) string {
	if request == nil {
		request = NewQueryCashWalletActionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryCashWalletActionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
