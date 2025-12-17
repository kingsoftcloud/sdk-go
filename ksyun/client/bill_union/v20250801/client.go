package v20250801

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-08-01"

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

func NewQueryItemBillsRequest() (request *QueryItemBillsRequest) {
	request = &QueryItemBillsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryItemBills")
	return
}

func NewQueryItemBillsResponse() (response *QueryItemBillsResponse) {
	response = &QueryItemBillsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryItemBills(request *QueryItemBillsRequest) string {
	return c.QueryItemBillsWithContext(context.Background(), request)
}

func (c *Client) QueryItemBillsSend(request *QueryItemBillsRequest) (*QueryItemBillsResponse, error) {
	statusCode, msg, err := c.QueryItemBillsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryItemBillsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryItemBillsWithContext(ctx context.Context, request *QueryItemBillsRequest) string {
	if request == nil {
		request = NewQueryItemBillsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryItemBillsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryItemBillsWithContextV2(ctx context.Context, request *QueryItemBillsRequest) (int, string, error) {
	if request == nil {
		request = NewQueryItemBillsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryItemBillsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryProductTypesRequest() (request *QueryProductTypesRequest) {
	request = &QueryProductTypesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "QueryProductTypes")
	return
}

func NewQueryProductTypesResponse() (response *QueryProductTypesResponse) {
	response = &QueryProductTypesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryProductTypes(request *QueryProductTypesRequest) string {
	return c.QueryProductTypesWithContext(context.Background(), request)
}

func (c *Client) QueryProductTypesSend(request *QueryProductTypesRequest) (*QueryProductTypesResponse, error) {
	statusCode, msg, err := c.QueryProductTypesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryProductTypesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryProductTypesWithContext(ctx context.Context, request *QueryProductTypesRequest) string {
	if request == nil {
		request = NewQueryProductTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryProductTypesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryProductTypesWithContextV2(ctx context.Context, request *QueryProductTypesRequest) (int, string, error) {
	if request == nil {
		request = NewQueryProductTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryProductTypesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
