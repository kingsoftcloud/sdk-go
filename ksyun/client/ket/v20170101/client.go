package v20170101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2017-01-01"

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

func NewPresetRequest() (request *PresetRequest) {
	request = &PresetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "Preset")
	return
}

func NewPresetResponse() (response *PresetResponse) {
	response = &PresetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Preset(request *PresetRequest) string {
	return c.PresetWithContext(context.Background(), request)
}

func (c *Client) PresetSend(request *PresetRequest) (*PresetResponse, error) {
	statusCode, msg, err := c.PresetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct PresetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) PresetWithContext(ctx context.Context, request *PresetRequest) string {
	if request == nil {
		request = NewPresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewPresetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) PresetWithContextV2(ctx context.Context, request *PresetRequest) (int, string, error) {
	if request == nil {
		request = NewPresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewPresetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdatePresetRequest() (request *UpdatePresetRequest) {
	request = &UpdatePresetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "UpdatePreset")
	return
}

func NewUpdatePresetResponse() (response *UpdatePresetResponse) {
	response = &UpdatePresetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePreset(request *UpdatePresetRequest) string {
	return c.UpdatePresetWithContext(context.Background(), request)
}

func (c *Client) UpdatePresetSend(request *UpdatePresetRequest) (*UpdatePresetResponse, error) {
	statusCode, msg, err := c.UpdatePresetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdatePresetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdatePresetWithContext(ctx context.Context, request *UpdatePresetRequest) string {
	if request == nil {
		request = NewUpdatePresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdatePresetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdatePresetWithContextV2(ctx context.Context, request *UpdatePresetRequest) (int, string, error) {
	if request == nil {
		request = NewUpdatePresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdatePresetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDelPresetRequest() (request *DelPresetRequest) {
	request = &DelPresetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "DelPreset")
	return
}

func NewDelPresetResponse() (response *DelPresetResponse) {
	response = &DelPresetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DelPreset(request *DelPresetRequest) string {
	return c.DelPresetWithContext(context.Background(), request)
}

func (c *Client) DelPresetSend(request *DelPresetRequest) (*DelPresetResponse, error) {
	statusCode, msg, err := c.DelPresetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DelPresetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DelPresetWithContext(ctx context.Context, request *DelPresetRequest) string {
	if request == nil {
		request = NewDelPresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDelPresetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DelPresetWithContextV2(ctx context.Context, request *DelPresetRequest) (int, string, error) {
	if request == nil {
		request = NewDelPresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDelPresetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPresetListRequest() (request *GetPresetListRequest) {
	request = &GetPresetListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "GetPresetList")
	return
}

func NewGetPresetListResponse() (response *GetPresetListResponse) {
	response = &GetPresetListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPresetList(request *GetPresetListRequest) string {
	return c.GetPresetListWithContext(context.Background(), request)
}

func (c *Client) GetPresetListSend(request *GetPresetListRequest) (*GetPresetListResponse, error) {
	statusCode, msg, err := c.GetPresetListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPresetListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPresetListWithContext(ctx context.Context, request *GetPresetListRequest) string {
	if request == nil {
		request = NewGetPresetListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPresetListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPresetListWithContextV2(ctx context.Context, request *GetPresetListRequest) (int, string, error) {
	if request == nil {
		request = NewGetPresetListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPresetListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetPresetDetailRequest() (request *GetPresetDetailRequest) {
	request = &GetPresetDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "GetPresetDetail")
	return
}

func NewGetPresetDetailResponse() (response *GetPresetDetailResponse) {
	response = &GetPresetDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPresetDetail(request *GetPresetDetailRequest) string {
	return c.GetPresetDetailWithContext(context.Background(), request)
}

func (c *Client) GetPresetDetailSend(request *GetPresetDetailRequest) (*GetPresetDetailResponse, error) {
	statusCode, msg, err := c.GetPresetDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetPresetDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetPresetDetailWithContext(ctx context.Context, request *GetPresetDetailRequest) string {
	if request == nil {
		request = NewGetPresetDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPresetDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetPresetDetailWithContextV2(ctx context.Context, request *GetPresetDetailRequest) (int, string, error) {
	if request == nil {
		request = NewGetPresetDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetPresetDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetStreamTranListRequest() (request *GetStreamTranListRequest) {
	request = &GetStreamTranListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "GetStreamTranList")
	return
}

func NewGetStreamTranListResponse() (response *GetStreamTranListResponse) {
	response = &GetStreamTranListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetStreamTranList(request *GetStreamTranListRequest) string {
	return c.GetStreamTranListWithContext(context.Background(), request)
}

func (c *Client) GetStreamTranListSend(request *GetStreamTranListRequest) (*GetStreamTranListResponse, error) {
	statusCode, msg, err := c.GetStreamTranListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetStreamTranListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetStreamTranListWithContext(ctx context.Context, request *GetStreamTranListRequest) string {
	if request == nil {
		request = NewGetStreamTranListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetStreamTranListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetStreamTranListWithContextV2(ctx context.Context, request *GetStreamTranListRequest) (int, string, error) {
	if request == nil {
		request = NewGetStreamTranListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetStreamTranListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartLoopRequest() (request *StartLoopRequest) {
	request = &StartLoopRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "StartLoop")
	return
}

func NewStartLoopResponse() (response *StartLoopResponse) {
	response = &StartLoopResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartLoop(request *StartLoopRequest) string {
	return c.StartLoopWithContext(context.Background(), request)
}

func (c *Client) StartLoopSend(request *StartLoopRequest) (*StartLoopResponse, error) {
	statusCode, msg, err := c.StartLoopWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartLoopResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartLoopWithContext(ctx context.Context, request *StartLoopRequest) string {
	if request == nil {
		request = NewStartLoopRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartLoopResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartLoopWithContextV2(ctx context.Context, request *StartLoopRequest) (int, string, error) {
	if request == nil {
		request = NewStartLoopRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartLoopResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopLoopRequest() (request *StopLoopRequest) {
	request = &StopLoopRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ket", APIVersion, "StopLoop")
	return
}

func NewStopLoopResponse() (response *StopLoopResponse) {
	response = &StopLoopResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopLoop(request *StopLoopRequest) string {
	return c.StopLoopWithContext(context.Background(), request)
}

func (c *Client) StopLoopSend(request *StopLoopRequest) (*StopLoopResponse, error) {
	statusCode, msg, err := c.StopLoopWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopLoopResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopLoopWithContext(ctx context.Context, request *StopLoopRequest) string {
	if request == nil {
		request = NewStopLoopRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopLoopResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopLoopWithContextV2(ctx context.Context, request *StopLoopRequest) (int, string, error) {
	if request == nil {
		request = NewStopLoopRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopLoopResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
