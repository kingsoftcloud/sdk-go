package v20250501
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

func (c *Client) QueryCashWalletActionSend(request *QueryCashWalletActionRequest) (*QueryCashWalletActionResponse, error) {
	statusCode, msg, err := c.QueryCashWalletActionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryCashWalletActionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) QueryCashWalletActionWithContextV2(ctx context.Context, request *QueryCashWalletActionRequest) (int, string, error) {
	if request == nil {
		request = NewQueryCashWalletActionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryCashWalletActionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


