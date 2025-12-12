package v20250503

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "V3"

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

func NewGetDomainLogsRequest() (request *GetDomainLogsRequest) {
	request = &GetDomainLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainLogs")
	return
}

func NewGetDomainLogsResponse() (response *GetDomainLogsResponse) {
	response = &GetDomainLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainLogs(request *GetDomainLogsRequest) string {
	return c.GetDomainLogsWithContext(context.Background(), request)
}

func (c *Client) GetDomainLogsSend(request *GetDomainLogsRequest) (*GetDomainLogsResponse, error) {
	statusCode, msg, err := c.GetDomainLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDomainLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainLogsWithContext(ctx context.Context, request *GetDomainLogsRequest) string {
	if request == nil {
		request = NewGetDomainLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainLogsWithContextV2(ctx context.Context, request *GetDomainLogsRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetClientRequestDataRequest() (request *GetClientRequestDataRequest) {
	request = &GetClientRequestDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetClientRequestData")
	return
}

func NewGetClientRequestDataResponse() (response *GetClientRequestDataResponse) {
	response = &GetClientRequestDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetClientRequestData(request *GetClientRequestDataRequest) string {
	return c.GetClientRequestDataWithContext(context.Background(), request)
}

func (c *Client) GetClientRequestDataSend(request *GetClientRequestDataRequest) (*GetClientRequestDataResponse, error) {
	statusCode, msg, err := c.GetClientRequestDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetClientRequestDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetClientRequestDataWithContext(ctx context.Context, request *GetClientRequestDataRequest) string {
	if request == nil {
		request = NewGetClientRequestDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetClientRequestDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetClientRequestDataWithContextV2(ctx context.Context, request *GetClientRequestDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetClientRequestDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetClientRequestDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetCdnDomainsRequest() (request *GetCdnDomainsRequest) {
	request = &GetCdnDomainsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetCdnDomains")
	return
}

func NewGetCdnDomainsResponse() (response *GetCdnDomainsResponse) {
	response = &GetCdnDomainsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetCdnDomains(request *GetCdnDomainsRequest) string {
	return c.GetCdnDomainsWithContext(context.Background(), request)
}

func (c *Client) GetCdnDomainsSend(request *GetCdnDomainsRequest) (*GetCdnDomainsResponse, error) {
	statusCode, msg, err := c.GetCdnDomainsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetCdnDomainsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetCdnDomainsWithContext(ctx context.Context, request *GetCdnDomainsRequest) string {
	if request == nil {
		request = NewGetCdnDomainsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCdnDomainsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetCdnDomainsWithContextV2(ctx context.Context, request *GetCdnDomainsRequest) (int, string, error) {
	if request == nil {
		request = NewGetCdnDomainsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCdnDomainsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
