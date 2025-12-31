package v3

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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainLogs")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainLogs")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetClientRequestData")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetClientRequestData")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetCdnDomains")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetCdnDomains")
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
func NewDeleteCdnDomainRequest() (request *DeleteCdnDomainRequest) {
	request = &DeleteCdnDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "DeleteCdnDomain")
	return
}

func NewDeleteCdnDomainResponse() (response *DeleteCdnDomainResponse) {
	response = &DeleteCdnDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) string {
	return c.DeleteCdnDomainWithContext(context.Background(), request)
}

func (c *Client) DeleteCdnDomainSend(request *DeleteCdnDomainRequest) (*DeleteCdnDomainResponse, error) {
	statusCode, msg, err := c.DeleteCdnDomainWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCdnDomainResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCdnDomainWithContext(ctx context.Context, request *DeleteCdnDomainRequest) string {
	if request == nil {
		request = NewDeleteCdnDomainRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteCdnDomain")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCdnDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCdnDomainWithContextV2(ctx context.Context, request *DeleteCdnDomainRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCdnDomainRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteCdnDomain")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCdnDomainResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetCdnDomainBasicInfoRequest() (request *GetCdnDomainBasicInfoRequest) {
	request = &GetCdnDomainBasicInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetCdnDomainBasicInfo")
	return
}

func NewGetCdnDomainBasicInfoResponse() (response *GetCdnDomainBasicInfoResponse) {
	response = &GetCdnDomainBasicInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetCdnDomainBasicInfo(request *GetCdnDomainBasicInfoRequest) string {
	return c.GetCdnDomainBasicInfoWithContext(context.Background(), request)
}

func (c *Client) GetCdnDomainBasicInfoSend(request *GetCdnDomainBasicInfoRequest) (*GetCdnDomainBasicInfoResponse, error) {
	statusCode, msg, err := c.GetCdnDomainBasicInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetCdnDomainBasicInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetCdnDomainBasicInfoWithContext(ctx context.Context, request *GetCdnDomainBasicInfoRequest) string {
	if request == nil {
		request = NewGetCdnDomainBasicInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetCdnDomainBasicInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCdnDomainBasicInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetCdnDomainBasicInfoWithContextV2(ctx context.Context, request *GetCdnDomainBasicInfoRequest) (int, string, error) {
	if request == nil {
		request = NewGetCdnDomainBasicInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetCdnDomainBasicInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCdnDomainBasicInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCdnDomainBasicInfoRequest() (request *ModifyCdnDomainBasicInfoRequest) {
	request = &ModifyCdnDomainBasicInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "ModifyCdnDomainBasicInfo")
	return
}

func NewModifyCdnDomainBasicInfoResponse() (response *ModifyCdnDomainBasicInfoResponse) {
	response = &ModifyCdnDomainBasicInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCdnDomainBasicInfo(request *ModifyCdnDomainBasicInfoRequest) string {
	return c.ModifyCdnDomainBasicInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyCdnDomainBasicInfoSend(request *ModifyCdnDomainBasicInfoRequest) (*ModifyCdnDomainBasicInfoResponse, error) {
	statusCode, msg, err := c.ModifyCdnDomainBasicInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCdnDomainBasicInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCdnDomainBasicInfoWithContext(ctx context.Context, request *ModifyCdnDomainBasicInfoRequest) string {
	if request == nil {
		request = NewModifyCdnDomainBasicInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "ModifyCdnDomainBasicInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCdnDomainBasicInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCdnDomainBasicInfoWithContextV2(ctx context.Context, request *ModifyCdnDomainBasicInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCdnDomainBasicInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "ModifyCdnDomainBasicInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCdnDomainBasicInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddCdnDomainRequest() (request *AddCdnDomainRequest) {
	request = &AddCdnDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "AddCdnDomain")
	return
}

func NewAddCdnDomainResponse() (response *AddCdnDomainResponse) {
	response = &AddCdnDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddCdnDomain(request *AddCdnDomainRequest) string {
	return c.AddCdnDomainWithContext(context.Background(), request)
}

func (c *Client) AddCdnDomainSend(request *AddCdnDomainRequest) (*AddCdnDomainResponse, error) {
	statusCode, msg, err := c.AddCdnDomainWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddCdnDomainResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddCdnDomainWithContext(ctx context.Context, request *AddCdnDomainRequest) string {
	if request == nil {
		request = NewAddCdnDomainRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "AddCdnDomain")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddCdnDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddCdnDomainWithContextV2(ctx context.Context, request *AddCdnDomainRequest) (int, string, error) {
	if request == nil {
		request = NewAddCdnDomainRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "AddCdnDomain")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddCdnDomainResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainConfigsRequest() (request *GetDomainConfigsRequest) {
	request = &GetDomainConfigsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainConfigs")
	return
}

func NewGetDomainConfigsResponse() (response *GetDomainConfigsResponse) {
	response = &GetDomainConfigsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainConfigs(request *GetDomainConfigsRequest) string {
	return c.GetDomainConfigsWithContext(context.Background(), request)
}

func (c *Client) GetDomainConfigsSend(request *GetDomainConfigsRequest) (*GetDomainConfigsResponse, error) {
	statusCode, msg, err := c.GetDomainConfigsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetDomainConfigsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainConfigsWithContext(ctx context.Context, request *GetDomainConfigsRequest) string {
	if request == nil {
		request = NewGetDomainConfigsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainConfigs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainConfigsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainConfigsWithContextV2(ctx context.Context, request *GetDomainConfigsRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainConfigsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainConfigs")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainConfigsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartStopCdnDomainRequest() (request *StartStopCdnDomainRequest) {
	request = &StartStopCdnDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "StartStopCdnDomain")
	return
}

func NewStartStopCdnDomainResponse() (response *StartStopCdnDomainResponse) {
	response = &StartStopCdnDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartStopCdnDomain(request *StartStopCdnDomainRequest) string {
	return c.StartStopCdnDomainWithContext(context.Background(), request)
}

func (c *Client) StartStopCdnDomainSend(request *StartStopCdnDomainRequest) (*StartStopCdnDomainResponse, error) {
	statusCode, msg, err := c.StartStopCdnDomainWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartStopCdnDomainResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartStopCdnDomainWithContext(ctx context.Context, request *StartStopCdnDomainRequest) string {
	if request == nil {
		request = NewStartStopCdnDomainRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "StartStopCdnDomain")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartStopCdnDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartStopCdnDomainWithContextV2(ctx context.Context, request *StartStopCdnDomainRequest) (int, string, error) {
	if request == nil {
		request = NewStartStopCdnDomainRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "StartStopCdnDomain")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartStopCdnDomainResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetCacheRuleConfigRequest() (request *SetCacheRuleConfigRequest) {
	request = &SetCacheRuleConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetCacheRuleConfig")
	return
}

func NewSetCacheRuleConfigResponse() (response *SetCacheRuleConfigResponse) {
	response = &SetCacheRuleConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetCacheRuleConfig(request *SetCacheRuleConfigRequest) string {
	return c.SetCacheRuleConfigWithContext(context.Background(), request)
}

func (c *Client) SetCacheRuleConfigSend(request *SetCacheRuleConfigRequest) (*SetCacheRuleConfigResponse, error) {
	statusCode, msg, err := c.SetCacheRuleConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetCacheRuleConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetCacheRuleConfigWithContext(ctx context.Context, request *SetCacheRuleConfigRequest) string {
	if request == nil {
		request = NewSetCacheRuleConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetCacheRuleConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetCacheRuleConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetCacheRuleConfigWithContextV2(ctx context.Context, request *SetCacheRuleConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetCacheRuleConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetCacheRuleConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetCacheRuleConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetBackOriginHostConfigRequest() (request *SetBackOriginHostConfigRequest) {
	request = &SetBackOriginHostConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetBackOriginHostConfig")
	return
}

func NewSetBackOriginHostConfigResponse() (response *SetBackOriginHostConfigResponse) {
	response = &SetBackOriginHostConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetBackOriginHostConfig(request *SetBackOriginHostConfigRequest) string {
	return c.SetBackOriginHostConfigWithContext(context.Background(), request)
}

func (c *Client) SetBackOriginHostConfigSend(request *SetBackOriginHostConfigRequest) (*SetBackOriginHostConfigResponse, error) {
	statusCode, msg, err := c.SetBackOriginHostConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetBackOriginHostConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetBackOriginHostConfigWithContext(ctx context.Context, request *SetBackOriginHostConfigRequest) string {
	if request == nil {
		request = NewSetBackOriginHostConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetBackOriginHostConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetBackOriginHostConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetBackOriginHostConfigWithContextV2(ctx context.Context, request *SetBackOriginHostConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetBackOriginHostConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetBackOriginHostConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetBackOriginHostConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetValidDomainListRequest() (request *GetValidDomainListRequest) {
	request = &GetValidDomainListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetValidDomainList")
	return
}

func NewGetValidDomainListResponse() (response *GetValidDomainListResponse) {
	response = &GetValidDomainListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetValidDomainList(request *GetValidDomainListRequest) string {
	return c.GetValidDomainListWithContext(context.Background(), request)
}

func (c *Client) GetValidDomainListSend(request *GetValidDomainListRequest) (*GetValidDomainListResponse, error) {
	statusCode, msg, err := c.GetValidDomainListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetValidDomainListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetValidDomainListWithContext(ctx context.Context, request *GetValidDomainListRequest) string {
	if request == nil {
		request = NewGetValidDomainListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetValidDomainList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetValidDomainListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetValidDomainListWithContextV2(ctx context.Context, request *GetValidDomainListRequest) (int, string, error) {
	if request == nil {
		request = NewGetValidDomainListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetValidDomainList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetValidDomainListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainAuthContentRequest() (request *GetDomainAuthContentRequest) {
	request = &GetDomainAuthContentRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainAuthContent")
	return
}

func NewGetDomainAuthContentResponse() (response *GetDomainAuthContentResponse) {
	response = &GetDomainAuthContentResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainAuthContent(request *GetDomainAuthContentRequest) string {
	return c.GetDomainAuthContentWithContext(context.Background(), request)
}

func (c *Client) GetDomainAuthContentSend(request *GetDomainAuthContentRequest) (*GetDomainAuthContentResponse, error) {
	statusCode, msg, err := c.GetDomainAuthContentWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetDomainAuthContentResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainAuthContentWithContext(ctx context.Context, request *GetDomainAuthContentRequest) string {
	if request == nil {
		request = NewGetDomainAuthContentRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainAuthContent")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainAuthContentResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainAuthContentWithContextV2(ctx context.Context, request *GetDomainAuthContentRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainAuthContentRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetDomainAuthContent")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainAuthContentResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetVideoSeekConfigRequest() (request *SetVideoSeekConfigRequest) {
	request = &SetVideoSeekConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetVideoSeekConfig")
	return
}

func NewSetVideoSeekConfigResponse() (response *SetVideoSeekConfigResponse) {
	response = &SetVideoSeekConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetVideoSeekConfig(request *SetVideoSeekConfigRequest) string {
	return c.SetVideoSeekConfigWithContext(context.Background(), request)
}

func (c *Client) SetVideoSeekConfigSend(request *SetVideoSeekConfigRequest) (*SetVideoSeekConfigResponse, error) {
	statusCode, msg, err := c.SetVideoSeekConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetVideoSeekConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetVideoSeekConfigWithContext(ctx context.Context, request *SetVideoSeekConfigRequest) string {
	if request == nil {
		request = NewSetVideoSeekConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetVideoSeekConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetVideoSeekConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetVideoSeekConfigWithContextV2(ctx context.Context, request *SetVideoSeekConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetVideoSeekConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetVideoSeekConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetVideoSeekConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetPageCompressConfigRequest() (request *SetPageCompressConfigRequest) {
	request = &SetPageCompressConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetPageCompressConfig")
	return
}

func NewSetPageCompressConfigResponse() (response *SetPageCompressConfigResponse) {
	response = &SetPageCompressConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetPageCompressConfig(request *SetPageCompressConfigRequest) string {
	return c.SetPageCompressConfigWithContext(context.Background(), request)
}

func (c *Client) SetPageCompressConfigSend(request *SetPageCompressConfigRequest) (*SetPageCompressConfigResponse, error) {
	statusCode, msg, err := c.SetPageCompressConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetPageCompressConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetPageCompressConfigWithContext(ctx context.Context, request *SetPageCompressConfigRequest) string {
	if request == nil {
		request = NewSetPageCompressConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetPageCompressConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetPageCompressConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetPageCompressConfigWithContextV2(ctx context.Context, request *SetPageCompressConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetPageCompressConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetPageCompressConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetPageCompressConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetBrCompressConfigRequest() (request *SetBrCompressConfigRequest) {
	request = &SetBrCompressConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetBrCompressConfig")
	return
}

func NewSetBrCompressConfigResponse() (response *SetBrCompressConfigResponse) {
	response = &SetBrCompressConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetBrCompressConfig(request *SetBrCompressConfigRequest) string {
	return c.SetBrCompressConfigWithContext(context.Background(), request)
}

func (c *Client) SetBrCompressConfigSend(request *SetBrCompressConfigRequest) (*SetBrCompressConfigResponse, error) {
	statusCode, msg, err := c.SetBrCompressConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetBrCompressConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetBrCompressConfigWithContext(ctx context.Context, request *SetBrCompressConfigRequest) string {
	if request == nil {
		request = NewSetBrCompressConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetBrCompressConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetBrCompressConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetBrCompressConfigWithContextV2(ctx context.Context, request *SetBrCompressConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetBrCompressConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetBrCompressConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetBrCompressConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetIgnoreQueryStringConfigRequest() (request *SetIgnoreQueryStringConfigRequest) {
	request = &SetIgnoreQueryStringConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetIgnoreQueryStringConfig")
	return
}

func NewSetIgnoreQueryStringConfigResponse() (response *SetIgnoreQueryStringConfigResponse) {
	response = &SetIgnoreQueryStringConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetIgnoreQueryStringConfig(request *SetIgnoreQueryStringConfigRequest) string {
	return c.SetIgnoreQueryStringConfigWithContext(context.Background(), request)
}

func (c *Client) SetIgnoreQueryStringConfigSend(request *SetIgnoreQueryStringConfigRequest) (*SetIgnoreQueryStringConfigResponse, error) {
	statusCode, msg, err := c.SetIgnoreQueryStringConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetIgnoreQueryStringConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetIgnoreQueryStringConfigWithContext(ctx context.Context, request *SetIgnoreQueryStringConfigRequest) string {
	if request == nil {
		request = NewSetIgnoreQueryStringConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetIgnoreQueryStringConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetIgnoreQueryStringConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetIgnoreQueryStringConfigWithContextV2(ctx context.Context, request *SetIgnoreQueryStringConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetIgnoreQueryStringConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetIgnoreQueryStringConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetIgnoreQueryStringConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetSetOriginAdvancedConfigRequest() (request *SetSetOriginAdvancedConfigRequest) {
	request = &SetSetOriginAdvancedConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetSetOriginAdvancedConfig")
	return
}

func NewSetSetOriginAdvancedConfigResponse() (response *SetSetOriginAdvancedConfigResponse) {
	response = &SetSetOriginAdvancedConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetSetOriginAdvancedConfig(request *SetSetOriginAdvancedConfigRequest) string {
	return c.SetSetOriginAdvancedConfigWithContext(context.Background(), request)
}

func (c *Client) SetSetOriginAdvancedConfigSend(request *SetSetOriginAdvancedConfigRequest) (*SetSetOriginAdvancedConfigResponse, error) {
	statusCode, msg, err := c.SetSetOriginAdvancedConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetSetOriginAdvancedConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetSetOriginAdvancedConfigWithContext(ctx context.Context, request *SetSetOriginAdvancedConfigRequest) string {
	if request == nil {
		request = NewSetSetOriginAdvancedConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetSetOriginAdvancedConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetSetOriginAdvancedConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetSetOriginAdvancedConfigWithContextV2(ctx context.Context, request *SetSetOriginAdvancedConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetSetOriginAdvancedConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetSetOriginAdvancedConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetSetOriginAdvancedConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewValidateDomainOwnerRequest() (request *ValidateDomainOwnerRequest) {
	request = &ValidateDomainOwnerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "ValidateDomainOwner")
	return
}

func NewValidateDomainOwnerResponse() (response *ValidateDomainOwnerResponse) {
	response = &ValidateDomainOwnerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ValidateDomainOwner(request *ValidateDomainOwnerRequest) string {
	return c.ValidateDomainOwnerWithContext(context.Background(), request)
}

func (c *Client) ValidateDomainOwnerSend(request *ValidateDomainOwnerRequest) (*ValidateDomainOwnerResponse, error) {
	statusCode, msg, err := c.ValidateDomainOwnerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ValidateDomainOwnerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ValidateDomainOwnerWithContext(ctx context.Context, request *ValidateDomainOwnerRequest) string {
	if request == nil {
		request = NewValidateDomainOwnerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "ValidateDomainOwner")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewValidateDomainOwnerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ValidateDomainOwnerWithContextV2(ctx context.Context, request *ValidateDomainOwnerRequest) (int, string, error) {
	if request == nil {
		request = NewValidateDomainOwnerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "ValidateDomainOwner")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewValidateDomainOwnerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetHttp2OptionConfigRequest() (request *SetHttp2OptionConfigRequest) {
	request = &SetHttp2OptionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetHttp2OptionConfig")
	return
}

func NewSetHttp2OptionConfigResponse() (response *SetHttp2OptionConfigResponse) {
	response = &SetHttp2OptionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetHttp2OptionConfig(request *SetHttp2OptionConfigRequest) string {
	return c.SetHttp2OptionConfigWithContext(context.Background(), request)
}

func (c *Client) SetHttp2OptionConfigSend(request *SetHttp2OptionConfigRequest) (*SetHttp2OptionConfigResponse, error) {
	statusCode, msg, err := c.SetHttp2OptionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetHttp2OptionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetHttp2OptionConfigWithContext(ctx context.Context, request *SetHttp2OptionConfigRequest) string {
	if request == nil {
		request = NewSetHttp2OptionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetHttp2OptionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetHttp2OptionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetHttp2OptionConfigWithContextV2(ctx context.Context, request *SetHttp2OptionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetHttp2OptionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetHttp2OptionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetHttp2OptionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetReferProtectionConfigRequest() (request *SetReferProtectionConfigRequest) {
	request = &SetReferProtectionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetReferProtectionConfig")
	return
}

func NewSetReferProtectionConfigResponse() (response *SetReferProtectionConfigResponse) {
	response = &SetReferProtectionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetReferProtectionConfig(request *SetReferProtectionConfigRequest) string {
	return c.SetReferProtectionConfigWithContext(context.Background(), request)
}

func (c *Client) SetReferProtectionConfigSend(request *SetReferProtectionConfigRequest) (*SetReferProtectionConfigResponse, error) {
	statusCode, msg, err := c.SetReferProtectionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetReferProtectionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetReferProtectionConfigWithContext(ctx context.Context, request *SetReferProtectionConfigRequest) string {
	if request == nil {
		request = NewSetReferProtectionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetReferProtectionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetReferProtectionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetReferProtectionConfigWithContextV2(ctx context.Context, request *SetReferProtectionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetReferProtectionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetReferProtectionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetReferProtectionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetIpProtectionConfigRequest() (request *SetIpProtectionConfigRequest) {
	request = &SetIpProtectionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetIpProtectionConfig")
	return
}

func NewSetIpProtectionConfigResponse() (response *SetIpProtectionConfigResponse) {
	response = &SetIpProtectionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetIpProtectionConfig(request *SetIpProtectionConfigRequest) string {
	return c.SetIpProtectionConfigWithContext(context.Background(), request)
}

func (c *Client) SetIpProtectionConfigSend(request *SetIpProtectionConfigRequest) (*SetIpProtectionConfigResponse, error) {
	statusCode, msg, err := c.SetIpProtectionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetIpProtectionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetIpProtectionConfigWithContext(ctx context.Context, request *SetIpProtectionConfigRequest) string {
	if request == nil {
		request = NewSetIpProtectionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetIpProtectionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetIpProtectionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetIpProtectionConfigWithContextV2(ctx context.Context, request *SetIpProtectionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetIpProtectionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetIpProtectionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetIpProtectionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetHttpHeadersConfigRequest() (request *SetHttpHeadersConfigRequest) {
	request = &SetHttpHeadersConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetHttpHeadersConfig")
	return
}

func NewSetHttpHeadersConfigResponse() (response *SetHttpHeadersConfigResponse) {
	response = &SetHttpHeadersConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetHttpHeadersConfig(request *SetHttpHeadersConfigRequest) string {
	return c.SetHttpHeadersConfigWithContext(context.Background(), request)
}

func (c *Client) SetHttpHeadersConfigSend(request *SetHttpHeadersConfigRequest) (*SetHttpHeadersConfigResponse, error) {
	statusCode, msg, err := c.SetHttpHeadersConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetHttpHeadersConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetHttpHeadersConfigWithContext(ctx context.Context, request *SetHttpHeadersConfigRequest) string {
	if request == nil {
		request = NewSetHttpHeadersConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetHttpHeadersConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetHttpHeadersConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetHttpHeadersConfigWithContextV2(ctx context.Context, request *SetHttpHeadersConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetHttpHeadersConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetHttpHeadersConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetHttpHeadersConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteHttpHeadersConfigRequest() (request *DeleteHttpHeadersConfigRequest) {
	request = &DeleteHttpHeadersConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "DeleteHttpHeadersConfig")
	return
}

func NewDeleteHttpHeadersConfigResponse() (response *DeleteHttpHeadersConfigResponse) {
	response = &DeleteHttpHeadersConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteHttpHeadersConfig(request *DeleteHttpHeadersConfigRequest) string {
	return c.DeleteHttpHeadersConfigWithContext(context.Background(), request)
}

func (c *Client) DeleteHttpHeadersConfigSend(request *DeleteHttpHeadersConfigRequest) (*DeleteHttpHeadersConfigResponse, error) {
	statusCode, msg, err := c.DeleteHttpHeadersConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteHttpHeadersConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteHttpHeadersConfigWithContext(ctx context.Context, request *DeleteHttpHeadersConfigRequest) string {
	if request == nil {
		request = NewDeleteHttpHeadersConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteHttpHeadersConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHttpHeadersConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteHttpHeadersConfigWithContextV2(ctx context.Context, request *DeleteHttpHeadersConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteHttpHeadersConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "DeleteHttpHeadersConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHttpHeadersConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetHttpHeaderListRequest() (request *GetHttpHeaderListRequest) {
	request = &GetHttpHeaderListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetHttpHeaderList")
	return
}

func NewGetHttpHeaderListResponse() (response *GetHttpHeaderListResponse) {
	response = &GetHttpHeaderListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetHttpHeaderList(request *GetHttpHeaderListRequest) string {
	return c.GetHttpHeaderListWithContext(context.Background(), request)
}

func (c *Client) GetHttpHeaderListSend(request *GetHttpHeaderListRequest) (*GetHttpHeaderListResponse, error) {
	statusCode, msg, err := c.GetHttpHeaderListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetHttpHeaderListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetHttpHeaderListWithContext(ctx context.Context, request *GetHttpHeaderListRequest) string {
	if request == nil {
		request = NewGetHttpHeaderListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetHttpHeaderList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetHttpHeaderListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetHttpHeaderListWithContextV2(ctx context.Context, request *GetHttpHeaderListRequest) (int, string, error) {
	if request == nil {
		request = NewGetHttpHeaderListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetHttpHeaderList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetHttpHeaderListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetRequestAuthConfigRequest() (request *SetRequestAuthConfigRequest) {
	request = &SetRequestAuthConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetRequestAuthConfig")
	return
}

func NewSetRequestAuthConfigResponse() (response *SetRequestAuthConfigResponse) {
	response = &SetRequestAuthConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetRequestAuthConfig(request *SetRequestAuthConfigRequest) string {
	return c.SetRequestAuthConfigWithContext(context.Background(), request)
}

func (c *Client) SetRequestAuthConfigSend(request *SetRequestAuthConfigRequest) (*SetRequestAuthConfigResponse, error) {
	statusCode, msg, err := c.SetRequestAuthConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetRequestAuthConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetRequestAuthConfigWithContext(ctx context.Context, request *SetRequestAuthConfigRequest) string {
	if request == nil {
		request = NewSetRequestAuthConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetRequestAuthConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetRequestAuthConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetRequestAuthConfigWithContextV2(ctx context.Context, request *SetRequestAuthConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetRequestAuthConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetRequestAuthConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetRequestAuthConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetForceRedirectConfigRequest() (request *SetForceRedirectConfigRequest) {
	request = &SetForceRedirectConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetForceRedirectConfig")
	return
}

func NewSetForceRedirectConfigResponse() (response *SetForceRedirectConfigResponse) {
	response = &SetForceRedirectConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetForceRedirectConfig(request *SetForceRedirectConfigRequest) string {
	return c.SetForceRedirectConfigWithContext(context.Background(), request)
}

func (c *Client) SetForceRedirectConfigSend(request *SetForceRedirectConfigRequest) (*SetForceRedirectConfigResponse, error) {
	statusCode, msg, err := c.SetForceRedirectConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetForceRedirectConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetForceRedirectConfigWithContext(ctx context.Context, request *SetForceRedirectConfigRequest) string {
	if request == nil {
		request = NewSetForceRedirectConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetForceRedirectConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetForceRedirectConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetForceRedirectConfigWithContextV2(ctx context.Context, request *SetForceRedirectConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetForceRedirectConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetForceRedirectConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetForceRedirectConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetErrorPageConfigRequest() (request *SetErrorPageConfigRequest) {
	request = &SetErrorPageConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetErrorPageConfig")
	return
}

func NewSetErrorPageConfigResponse() (response *SetErrorPageConfigResponse) {
	response = &SetErrorPageConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetErrorPageConfig(request *SetErrorPageConfigRequest) string {
	return c.SetErrorPageConfigWithContext(context.Background(), request)
}

func (c *Client) SetErrorPageConfigSend(request *SetErrorPageConfigRequest) (*SetErrorPageConfigResponse, error) {
	statusCode, msg, err := c.SetErrorPageConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetErrorPageConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetErrorPageConfigWithContext(ctx context.Context, request *SetErrorPageConfigRequest) string {
	if request == nil {
		request = NewSetErrorPageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetErrorPageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetErrorPageConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetErrorPageConfigWithContextV2(ctx context.Context, request *SetErrorPageConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetErrorPageConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetErrorPageConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetErrorPageConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetTLSVersionConfigRequest() (request *SetTLSVersionConfigRequest) {
	request = &SetTLSVersionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "SetTLSVersionConfig")
	return
}

func NewSetTLSVersionConfigResponse() (response *SetTLSVersionConfigResponse) {
	response = &SetTLSVersionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetTLSVersionConfig(request *SetTLSVersionConfigRequest) string {
	return c.SetTLSVersionConfigWithContext(context.Background(), request)
}

func (c *Client) SetTLSVersionConfigSend(request *SetTLSVersionConfigRequest) (*SetTLSVersionConfigResponse, error) {
	statusCode, msg, err := c.SetTLSVersionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetTLSVersionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetTLSVersionConfigWithContext(ctx context.Context, request *SetTLSVersionConfigRequest) string {
	if request == nil {
		request = NewSetTLSVersionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetTLSVersionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetTLSVersionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetTLSVersionConfigWithContextV2(ctx context.Context, request *SetTLSVersionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetTLSVersionConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "SetTLSVersionConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetTLSVersionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBillingModeRequest() (request *GetBillingModeRequest) {
	request = &GetBillingModeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetBillingMode")
	return
}

func NewGetBillingModeResponse() (response *GetBillingModeResponse) {
	response = &GetBillingModeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBillingMode(request *GetBillingModeRequest) string {
	return c.GetBillingModeWithContext(context.Background(), request)
}

func (c *Client) GetBillingModeSend(request *GetBillingModeRequest) (*GetBillingModeResponse, error) {
	statusCode, msg, err := c.GetBillingModeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetBillingModeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBillingModeWithContext(ctx context.Context, request *GetBillingModeRequest) string {
	if request == nil {
		request = NewGetBillingModeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetBillingMode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBillingModeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBillingModeWithContextV2(ctx context.Context, request *GetBillingModeRequest) (int, string, error) {
	if request == nil {
		request = NewGetBillingModeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetBillingMode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBillingModeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBlockUrlQuotaRequest() (request *GetBlockUrlQuotaRequest) {
	request = &GetBlockUrlQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetBlockUrlQuota")
	return
}

func NewGetBlockUrlQuotaResponse() (response *GetBlockUrlQuotaResponse) {
	response = &GetBlockUrlQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBlockUrlQuota(request *GetBlockUrlQuotaRequest) string {
	return c.GetBlockUrlQuotaWithContext(context.Background(), request)
}

func (c *Client) GetBlockUrlQuotaSend(request *GetBlockUrlQuotaRequest) (*GetBlockUrlQuotaResponse, error) {
	statusCode, msg, err := c.GetBlockUrlQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetBlockUrlQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBlockUrlQuotaWithContext(ctx context.Context, request *GetBlockUrlQuotaRequest) string {
	if request == nil {
		request = NewGetBlockUrlQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetBlockUrlQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBlockUrlQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBlockUrlQuotaWithContextV2(ctx context.Context, request *GetBlockUrlQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewGetBlockUrlQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetBlockUrlQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBlockUrlQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBandwidthDataRequest() (request *GetBandwidthDataRequest) {
	request = &GetBandwidthDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetBandwidthData")
	return
}

func NewGetBandwidthDataResponse() (response *GetBandwidthDataResponse) {
	response = &GetBandwidthDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBandwidthData(request *GetBandwidthDataRequest) string {
	return c.GetBandwidthDataWithContext(context.Background(), request)
}

func (c *Client) GetBandwidthDataSend(request *GetBandwidthDataRequest) (*GetBandwidthDataResponse, error) {
	statusCode, msg, err := c.GetBandwidthDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetBandwidthDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBandwidthDataWithContext(ctx context.Context, request *GetBandwidthDataRequest) string {
	if request == nil {
		request = NewGetBandwidthDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetBandwidthData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBandwidthDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBandwidthDataWithContextV2(ctx context.Context, request *GetBandwidthDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetBandwidthDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetBandwidthData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBandwidthDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetFlowDataRequest() (request *GetFlowDataRequest) {
	request = &GetFlowDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetFlowData")
	return
}

func NewGetFlowDataResponse() (response *GetFlowDataResponse) {
	response = &GetFlowDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetFlowData(request *GetFlowDataRequest) string {
	return c.GetFlowDataWithContext(context.Background(), request)
}

func (c *Client) GetFlowDataSend(request *GetFlowDataRequest) (*GetFlowDataResponse, error) {
	statusCode, msg, err := c.GetFlowDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetFlowDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetFlowDataWithContext(ctx context.Context, request *GetFlowDataRequest) string {
	if request == nil {
		request = NewGetFlowDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetFlowData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetFlowDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetFlowDataWithContextV2(ctx context.Context, request *GetFlowDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetFlowDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetFlowData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetFlowDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPvDataRequest() (request *GetPvDataRequest) {
	request = &GetPvDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetPvData")
	return
}

func NewGetPvDataResponse() (response *GetPvDataResponse) {
	response = &GetPvDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPvData(request *GetPvDataRequest) string {
	return c.GetPvDataWithContext(context.Background(), request)
}

func (c *Client) GetPvDataSend(request *GetPvDataRequest) (*GetPvDataResponse, error) {
	statusCode, msg, err := c.GetPvDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPvDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPvDataWithContext(ctx context.Context, request *GetPvDataRequest) string {
	if request == nil {
		request = NewGetPvDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetPvData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPvDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPvDataWithContextV2(ctx context.Context, request *GetPvDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetPvDataRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cdn", APIVersion, "GetPvData")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPvDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
