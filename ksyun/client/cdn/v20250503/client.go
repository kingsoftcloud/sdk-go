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
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetHttp2OptionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
