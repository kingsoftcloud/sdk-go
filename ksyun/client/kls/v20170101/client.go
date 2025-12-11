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

func NewListStreamDurationsRequest() (request *ListStreamDurationsRequest) {
	request = &ListStreamDurationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "ListStreamDurations")
	return
}

func NewListStreamDurationsResponse() (response *ListStreamDurationsResponse) {
	response = &ListStreamDurationsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListStreamDurations(request *ListStreamDurationsRequest) string {
	return c.ListStreamDurationsWithContext(context.Background(), request)
}

func (c *Client) ListStreamDurationsSend(request *ListStreamDurationsRequest) (*ListStreamDurationsResponse, error) {
	statusCode, msg, err := c.ListStreamDurationsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListStreamDurationsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListStreamDurationsWithContext(ctx context.Context, request *ListStreamDurationsRequest) string {
	if request == nil {
		request = NewListStreamDurationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListStreamDurationsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListStreamDurationsWithContextV2(ctx context.Context, request *ListStreamDurationsRequest) (int, string, error) {
	if request == nil {
		request = NewListStreamDurationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListStreamDurationsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListHistoryPubStreamsErrInfoRequest() (request *ListHistoryPubStreamsErrInfoRequest) {
	request = &ListHistoryPubStreamsErrInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "ListHistoryPubStreamsErrInfo")
	return
}

func NewListHistoryPubStreamsErrInfoResponse() (response *ListHistoryPubStreamsErrInfoResponse) {
	response = &ListHistoryPubStreamsErrInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListHistoryPubStreamsErrInfo(request *ListHistoryPubStreamsErrInfoRequest) string {
	return c.ListHistoryPubStreamsErrInfoWithContext(context.Background(), request)
}

func (c *Client) ListHistoryPubStreamsErrInfoSend(request *ListHistoryPubStreamsErrInfoRequest) (*ListHistoryPubStreamsErrInfoResponse, error) {
	statusCode, msg, err := c.ListHistoryPubStreamsErrInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListHistoryPubStreamsErrInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListHistoryPubStreamsErrInfoWithContext(ctx context.Context, request *ListHistoryPubStreamsErrInfoRequest) string {
	if request == nil {
		request = NewListHistoryPubStreamsErrInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListHistoryPubStreamsErrInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListHistoryPubStreamsErrInfoWithContextV2(ctx context.Context, request *ListHistoryPubStreamsErrInfoRequest) (int, string, error) {
	if request == nil {
		request = NewListHistoryPubStreamsErrInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListHistoryPubStreamsErrInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListHistoryPubStreamsInfoRequest() (request *ListHistoryPubStreamsInfoRequest) {
	request = &ListHistoryPubStreamsInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "ListHistoryPubStreamsInfo")
	return
}

func NewListHistoryPubStreamsInfoResponse() (response *ListHistoryPubStreamsInfoResponse) {
	response = &ListHistoryPubStreamsInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListHistoryPubStreamsInfo(request *ListHistoryPubStreamsInfoRequest) string {
	return c.ListHistoryPubStreamsInfoWithContext(context.Background(), request)
}

func (c *Client) ListHistoryPubStreamsInfoSend(request *ListHistoryPubStreamsInfoRequest) (*ListHistoryPubStreamsInfoResponse, error) {
	statusCode, msg, err := c.ListHistoryPubStreamsInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListHistoryPubStreamsInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListHistoryPubStreamsInfoWithContext(ctx context.Context, request *ListHistoryPubStreamsInfoRequest) string {
	if request == nil {
		request = NewListHistoryPubStreamsInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListHistoryPubStreamsInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListHistoryPubStreamsInfoWithContextV2(ctx context.Context, request *ListHistoryPubStreamsInfoRequest) (int, string, error) {
	if request == nil {
		request = NewListHistoryPubStreamsInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListHistoryPubStreamsInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewForbidStreamRequest() (request *ForbidStreamRequest) {
	request = &ForbidStreamRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "ForbidStream")
	return
}

func NewForbidStreamResponse() (response *ForbidStreamResponse) {
	response = &ForbidStreamResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ForbidStream(request *ForbidStreamRequest) string {
	return c.ForbidStreamWithContext(context.Background(), request)
}

func (c *Client) ForbidStreamSend(request *ForbidStreamRequest) (*ForbidStreamResponse, error) {
	statusCode, msg, err := c.ForbidStreamWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ForbidStreamResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ForbidStreamWithContext(ctx context.Context, request *ForbidStreamRequest) string {
	if request == nil {
		request = NewForbidStreamRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewForbidStreamResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ForbidStreamWithContextV2(ctx context.Context, request *ForbidStreamRequest) (int, string, error) {
	if request == nil {
		request = NewForbidStreamRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewForbidStreamResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResumeStreamRequest() (request *ResumeStreamRequest) {
	request = &ResumeStreamRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "ResumeStream")
	return
}

func NewResumeStreamResponse() (response *ResumeStreamResponse) {
	response = &ResumeStreamResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResumeStream(request *ResumeStreamRequest) string {
	return c.ResumeStreamWithContext(context.Background(), request)
}

func (c *Client) ResumeStreamSend(request *ResumeStreamRequest) (*ResumeStreamResponse, error) {
	statusCode, msg, err := c.ResumeStreamWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResumeStreamResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResumeStreamWithContext(ctx context.Context, request *ResumeStreamRequest) string {
	if request == nil {
		request = NewResumeStreamRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResumeStreamResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResumeStreamWithContextV2(ctx context.Context, request *ResumeStreamRequest) (int, string, error) {
	if request == nil {
		request = NewResumeStreamRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResumeStreamResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBlacklistRequest() (request *GetBlacklistRequest) {
	request = &GetBlacklistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "GetBlacklist")
	return
}

func NewGetBlacklistResponse() (response *GetBlacklistResponse) {
	response = &GetBlacklistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBlacklist(request *GetBlacklistRequest) string {
	return c.GetBlacklistWithContext(context.Background(), request)
}

func (c *Client) GetBlacklistSend(request *GetBlacklistRequest) (*GetBlacklistResponse, error) {
	statusCode, msg, err := c.GetBlacklistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetBlacklistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBlacklistWithContext(ctx context.Context, request *GetBlacklistRequest) string {
	if request == nil {
		request = NewGetBlacklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBlacklistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBlacklistWithContextV2(ctx context.Context, request *GetBlacklistRequest) (int, string, error) {
	if request == nil {
		request = NewGetBlacklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBlacklistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCheckBlacklistRequest() (request *CheckBlacklistRequest) {
	request = &CheckBlacklistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "CheckBlacklist")
	return
}

func NewCheckBlacklistResponse() (response *CheckBlacklistResponse) {
	response = &CheckBlacklistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckBlacklist(request *CheckBlacklistRequest) string {
	return c.CheckBlacklistWithContext(context.Background(), request)
}

func (c *Client) CheckBlacklistSend(request *CheckBlacklistRequest) (*CheckBlacklistResponse, error) {
	statusCode, msg, err := c.CheckBlacklistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CheckBlacklistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CheckBlacklistWithContext(ctx context.Context, request *CheckBlacklistRequest) string {
	if request == nil {
		request = NewCheckBlacklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckBlacklistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CheckBlacklistWithContextV2(ctx context.Context, request *CheckBlacklistRequest) (int, string, error) {
	if request == nil {
		request = NewCheckBlacklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckBlacklistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListRealtimeStreamsInfoRequest() (request *ListRealtimeStreamsInfoRequest) {
	request = &ListRealtimeStreamsInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kls", APIVersion, "ListRealtimeStreamsInfo")
	return
}

func NewListRealtimeStreamsInfoResponse() (response *ListRealtimeStreamsInfoResponse) {
	response = &ListRealtimeStreamsInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListRealtimeStreamsInfo(request *ListRealtimeStreamsInfoRequest) string {
	return c.ListRealtimeStreamsInfoWithContext(context.Background(), request)
}

func (c *Client) ListRealtimeStreamsInfoSend(request *ListRealtimeStreamsInfoRequest) (*ListRealtimeStreamsInfoResponse, error) {
	statusCode, msg, err := c.ListRealtimeStreamsInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListRealtimeStreamsInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListRealtimeStreamsInfoWithContext(ctx context.Context, request *ListRealtimeStreamsInfoRequest) string {
	if request == nil {
		request = NewListRealtimeStreamsInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListRealtimeStreamsInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListRealtimeStreamsInfoWithContextV2(ctx context.Context, request *ListRealtimeStreamsInfoRequest) (int, string, error) {
	if request == nil {
		request = NewListRealtimeStreamsInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListRealtimeStreamsInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
