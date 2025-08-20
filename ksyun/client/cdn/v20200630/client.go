package v20200630

import (
	"context"
	"fmt"

	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-06-30"

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
func NewGetServerDataRequest() (request *GetServerDataRequest) {
	request = &GetServerDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetServerData")
	return
}

func NewGetServerDataResponse() (response *GetServerDataResponse) {
	response = &GetServerDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetServerData(request *GetServerDataRequest) string {
	return c.GetServerDataWithContext(context.Background(), request)
}

func (c *Client) GetServerDataSend(request *GetServerDataRequest) (*GetServerDataResponse, error) {
	statusCode, msg, err := c.GetServerDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetServerDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetServerDataWithContext(ctx context.Context, request *GetServerDataRequest) string {
	if request == nil {
		request = NewGetServerDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetServerDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetServerDataWithContextV2(ctx context.Context, request *GetServerDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetServerDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetServerDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainRankingListDataRequest() (request *GetDomainRankingListDataRequest) {
	request = &GetDomainRankingListDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainRankingListData")
	return
}

func NewGetDomainRankingListDataResponse() (response *GetDomainRankingListDataResponse) {
	response = &GetDomainRankingListDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainRankingListData(request *GetDomainRankingListDataRequest) string {
	return c.GetDomainRankingListDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainRankingListDataSend(request *GetDomainRankingListDataRequest) (*GetDomainRankingListDataResponse, error) {
	statusCode, msg, err := c.GetDomainRankingListDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDomainRankingListDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainRankingListDataWithContext(ctx context.Context, request *GetDomainRankingListDataRequest) string {
	if request == nil {
		request = NewGetDomainRankingListDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainRankingListDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainRankingListDataWithContextV2(ctx context.Context, request *GetDomainRankingListDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainRankingListDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainRankingListDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetAreaIspDataRequest() (request *GetAreaIspDataRequest) {
	request = &GetAreaIspDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetAreaIspData")
	return
}

func NewGetAreaIspDataResponse() (response *GetAreaIspDataResponse) {
	response = &GetAreaIspDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetAreaIspData(request *GetAreaIspDataRequest) string {
	return c.GetAreaIspDataWithContext(context.Background(), request)
}

func (c *Client) GetAreaIspDataSend(request *GetAreaIspDataRequest) (*GetAreaIspDataResponse, error) {
	statusCode, msg, err := c.GetAreaIspDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetAreaIspDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetAreaIspDataWithContext(ctx context.Context, request *GetAreaIspDataRequest) string {
	if request == nil {
		request = NewGetAreaIspDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetAreaIspDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetAreaIspDataWithContextV2(ctx context.Context, request *GetAreaIspDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetAreaIspDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetAreaIspDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetTopReferDataRequest() (request *GetTopReferDataRequest) {
	request = &GetTopReferDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetTopReferData")
	return
}

func NewGetTopReferDataResponse() (response *GetTopReferDataResponse) {
	response = &GetTopReferDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetTopReferData(request *GetTopReferDataRequest) string {
	return c.GetTopReferDataWithContext(context.Background(), request)
}

func (c *Client) GetTopReferDataSend(request *GetTopReferDataRequest) (*GetTopReferDataResponse, error) {
	statusCode, msg, err := c.GetTopReferDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetTopReferDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetTopReferDataWithContext(ctx context.Context, request *GetTopReferDataRequest) string {
	if request == nil {
		request = NewGetTopReferDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTopReferDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetTopReferDataWithContextV2(ctx context.Context, request *GetTopReferDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetTopReferDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTopReferDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetTopUrlDataRequest() (request *GetTopUrlDataRequest) {
	request = &GetTopUrlDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetTopUrlData")
	return
}

func NewGetTopUrlDataResponse() (response *GetTopUrlDataResponse) {
	response = &GetTopUrlDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetTopUrlData(request *GetTopUrlDataRequest) string {
	return c.GetTopUrlDataWithContext(context.Background(), request)
}

func (c *Client) GetTopUrlDataSend(request *GetTopUrlDataRequest) (*GetTopUrlDataResponse, error) {
	statusCode, msg, err := c.GetTopUrlDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetTopUrlDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetTopUrlDataWithContext(ctx context.Context, request *GetTopUrlDataRequest) string {
	if request == nil {
		request = NewGetTopUrlDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTopUrlDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetTopUrlDataWithContextV2(ctx context.Context, request *GetTopUrlDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetTopUrlDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTopUrlDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRealTimeHitRateDataRequest() (request *GetRealTimeHitRateDataRequest) {
	request = &GetRealTimeHitRateDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetRealTimeHitRateData")
	return
}

func NewGetRealTimeHitRateDataResponse() (response *GetRealTimeHitRateDataResponse) {
	response = &GetRealTimeHitRateDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRealTimeHitRateData(request *GetRealTimeHitRateDataRequest) string {
	return c.GetRealTimeHitRateDataWithContext(context.Background(), request)
}

func (c *Client) GetRealTimeHitRateDataSend(request *GetRealTimeHitRateDataRequest) (*GetRealTimeHitRateDataResponse, error) {
	statusCode, msg, err := c.GetRealTimeHitRateDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRealTimeHitRateDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRealTimeHitRateDataWithContext(ctx context.Context, request *GetRealTimeHitRateDataRequest) string {
	if request == nil {
		request = NewGetRealTimeHitRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRealTimeHitRateDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRealTimeHitRateDataWithContextV2(ctx context.Context, request *GetRealTimeHitRateDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetRealTimeHitRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRealTimeHitRateDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetReqHitRateDataRequest() (request *GetReqHitRateDataRequest) {
	request = &GetReqHitRateDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetReqHitRateData")
	return
}

func NewGetReqHitRateDataResponse() (response *GetReqHitRateDataResponse) {
	response = &GetReqHitRateDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetReqHitRateData(request *GetReqHitRateDataRequest) string {
	return c.GetReqHitRateDataWithContext(context.Background(), request)
}

func (c *Client) GetReqHitRateDataSend(request *GetReqHitRateDataRequest) (*GetReqHitRateDataResponse, error) {
	statusCode, msg, err := c.GetReqHitRateDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetReqHitRateDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetReqHitRateDataWithContext(ctx context.Context, request *GetReqHitRateDataRequest) string {
	if request == nil {
		request = NewGetReqHitRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetReqHitRateDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetReqHitRateDataWithContextV2(ctx context.Context, request *GetReqHitRateDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetReqHitRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetReqHitRateDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetFlowHitRateDataRequest() (request *GetFlowHitRateDataRequest) {
	request = &GetFlowHitRateDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetFlowHitRateData")
	return
}

func NewGetFlowHitRateDataResponse() (response *GetFlowHitRateDataResponse) {
	response = &GetFlowHitRateDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetFlowHitRateData(request *GetFlowHitRateDataRequest) string {
	return c.GetFlowHitRateDataWithContext(context.Background(), request)
}

func (c *Client) GetFlowHitRateDataSend(request *GetFlowHitRateDataRequest) (*GetFlowHitRateDataResponse, error) {
	statusCode, msg, err := c.GetFlowHitRateDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetFlowHitRateDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetFlowHitRateDataWithContext(ctx context.Context, request *GetFlowHitRateDataRequest) string {
	if request == nil {
		request = NewGetFlowHitRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetFlowHitRateDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetFlowHitRateDataWithContextV2(ctx context.Context, request *GetFlowHitRateDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetFlowHitRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetFlowHitRateDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainRequestPeriodRatioDataRequest() (request *GetDomainRequestPeriodRatioDataRequest) {
	request = &GetDomainRequestPeriodRatioDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainRequestPeriodRatioData")
	return
}

func NewGetDomainRequestPeriodRatioDataResponse() (response *GetDomainRequestPeriodRatioDataResponse) {
	response = &GetDomainRequestPeriodRatioDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainRequestPeriodRatioData(request *GetDomainRequestPeriodRatioDataRequest) string {
	return c.GetDomainRequestPeriodRatioDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainRequestPeriodRatioDataSend(request *GetDomainRequestPeriodRatioDataRequest) (*GetDomainRequestPeriodRatioDataResponse, error) {
	statusCode, msg, err := c.GetDomainRequestPeriodRatioDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDomainRequestPeriodRatioDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainRequestPeriodRatioDataWithContext(ctx context.Context, request *GetDomainRequestPeriodRatioDataRequest) string {
	if request == nil {
		request = NewGetDomainRequestPeriodRatioDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainRequestPeriodRatioDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainRequestPeriodRatioDataWithContextV2(ctx context.Context, request *GetDomainRequestPeriodRatioDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainRequestPeriodRatioDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainRequestPeriodRatioDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetUvDataRequest() (request *GetUvDataRequest) {
	request = &GetUvDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetUvData")
	return
}

func NewGetUvDataResponse() (response *GetUvDataResponse) {
	response = &GetUvDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUvData(request *GetUvDataRequest) string {
	return c.GetUvDataWithContext(context.Background(), request)
}

func (c *Client) GetUvDataSend(request *GetUvDataRequest) (*GetUvDataResponse, error) {
	statusCode, msg, err := c.GetUvDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetUvDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetUvDataWithContext(ctx context.Context, request *GetUvDataRequest) string {
	if request == nil {
		request = NewGetUvDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUvDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetUvDataWithContextV2(ctx context.Context, request *GetUvDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetUvDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUvDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetTopIpDataRequest() (request *GetTopIpDataRequest) {
	request = &GetTopIpDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetTopIpData")
	return
}

func NewGetTopIpDataResponse() (response *GetTopIpDataResponse) {
	response = &GetTopIpDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetTopIpData(request *GetTopIpDataRequest) string {
	return c.GetTopIpDataWithContext(context.Background(), request)
}

func (c *Client) GetTopIpDataSend(request *GetTopIpDataRequest) (*GetTopIpDataResponse, error) {
	statusCode, msg, err := c.GetTopIpDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetTopIpDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetTopIpDataWithContext(ctx context.Context, request *GetTopIpDataRequest) string {
	if request == nil {
		request = NewGetTopIpDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTopIpDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetTopIpDataWithContextV2(ctx context.Context, request *GetTopIpDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetTopIpDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTopIpDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetSrcDomainHttpCodeDetailedDataRequest() (request *GetSrcDomainHttpCodeDetailedDataRequest) {
	request = &GetSrcDomainHttpCodeDetailedDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetSrcDomainHttpCodeDetailedData")
	return
}

func NewGetSrcDomainHttpCodeDetailedDataResponse() (response *GetSrcDomainHttpCodeDetailedDataResponse) {
	response = &GetSrcDomainHttpCodeDetailedDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetSrcDomainHttpCodeDetailedData(request *GetSrcDomainHttpCodeDetailedDataRequest) string {
	return c.GetSrcDomainHttpCodeDetailedDataWithContext(context.Background(), request)
}

func (c *Client) GetSrcDomainHttpCodeDetailedDataSend(request *GetSrcDomainHttpCodeDetailedDataRequest) (*GetSrcDomainHttpCodeDetailedDataResponse, error) {
	statusCode, msg, err := c.GetSrcDomainHttpCodeDetailedDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetSrcDomainHttpCodeDetailedDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetSrcDomainHttpCodeDetailedDataWithContext(ctx context.Context, request *GetSrcDomainHttpCodeDetailedDataRequest) string {
	if request == nil {
		request = NewGetSrcDomainHttpCodeDetailedDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSrcDomainHttpCodeDetailedDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetSrcDomainHttpCodeDetailedDataWithContextV2(ctx context.Context, request *GetSrcDomainHttpCodeDetailedDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetSrcDomainHttpCodeDetailedDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSrcDomainHttpCodeDetailedDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetSrcDomainHttpCodeDataRequest() (request *GetSrcDomainHttpCodeDataRequest) {
	request = &GetSrcDomainHttpCodeDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetSrcDomainHttpCodeData")
	return
}

func NewGetSrcDomainHttpCodeDataResponse() (response *GetSrcDomainHttpCodeDataResponse) {
	response = &GetSrcDomainHttpCodeDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetSrcDomainHttpCodeData(request *GetSrcDomainHttpCodeDataRequest) string {
	return c.GetSrcDomainHttpCodeDataWithContext(context.Background(), request)
}

func (c *Client) GetSrcDomainHttpCodeDataSend(request *GetSrcDomainHttpCodeDataRequest) (*GetSrcDomainHttpCodeDataResponse, error) {
	statusCode, msg, err := c.GetSrcDomainHttpCodeDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetSrcDomainHttpCodeDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetSrcDomainHttpCodeDataWithContext(ctx context.Context, request *GetSrcDomainHttpCodeDataRequest) string {
	if request == nil {
		request = NewGetSrcDomainHttpCodeDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSrcDomainHttpCodeDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetSrcDomainHttpCodeDataWithContextV2(ctx context.Context, request *GetSrcDomainHttpCodeDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetSrcDomainHttpCodeDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetSrcDomainHttpCodeDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainHttpCodeDetailedDataRequest() (request *GetDomainHttpCodeDetailedDataRequest) {
	request = &GetDomainHttpCodeDetailedDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainHttpCodeDetailedData")
	return
}

func NewGetDomainHttpCodeDetailedDataResponse() (response *GetDomainHttpCodeDetailedDataResponse) {
	response = &GetDomainHttpCodeDetailedDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainHttpCodeDetailedData(request *GetDomainHttpCodeDetailedDataRequest) string {
	return c.GetDomainHttpCodeDetailedDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainHttpCodeDetailedDataSend(request *GetDomainHttpCodeDetailedDataRequest) (*GetDomainHttpCodeDetailedDataResponse, error) {
	statusCode, msg, err := c.GetDomainHttpCodeDetailedDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDomainHttpCodeDetailedDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainHttpCodeDetailedDataWithContext(ctx context.Context, request *GetDomainHttpCodeDetailedDataRequest) string {
	if request == nil {
		request = NewGetDomainHttpCodeDetailedDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainHttpCodeDetailedDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainHttpCodeDetailedDataWithContextV2(ctx context.Context, request *GetDomainHttpCodeDetailedDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainHttpCodeDetailedDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainHttpCodeDetailedDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDomainHttpCodeDataRequest() (request *GetDomainHttpCodeDataRequest) {
	request = &GetDomainHttpCodeDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainHttpCodeData")
	return
}

func NewGetDomainHttpCodeDataResponse() (response *GetDomainHttpCodeDataResponse) {
	response = &GetDomainHttpCodeDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainHttpCodeData(request *GetDomainHttpCodeDataRequest) string {
	return c.GetDomainHttpCodeDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainHttpCodeDataSend(request *GetDomainHttpCodeDataRequest) (*GetDomainHttpCodeDataResponse, error) {
	statusCode, msg, err := c.GetDomainHttpCodeDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDomainHttpCodeDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDomainHttpCodeDataWithContext(ctx context.Context, request *GetDomainHttpCodeDataRequest) string {
	if request == nil {
		request = NewGetDomainHttpCodeDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainHttpCodeDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDomainHttpCodeDataWithContextV2(ctx context.Context, request *GetDomainHttpCodeDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetDomainHttpCodeDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainHttpCodeDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetEntryRateDataRequest() (request *GetEntryRateDataRequest) {
	request = &GetEntryRateDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetEntryRateData")
	return
}

func NewGetEntryRateDataResponse() (response *GetEntryRateDataResponse) {
	response = &GetEntryRateDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetEntryRateData(request *GetEntryRateDataRequest) string {
	return c.GetEntryRateDataWithContext(context.Background(), request)
}

func (c *Client) GetEntryRateDataSend(request *GetEntryRateDataRequest) (*GetEntryRateDataResponse, error) {
	statusCode, msg, err := c.GetEntryRateDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetEntryRateDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetEntryRateDataWithContext(ctx context.Context, request *GetEntryRateDataRequest) string {
	if request == nil {
		request = NewGetEntryRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetEntryRateDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetEntryRateDataWithContextV2(ctx context.Context, request *GetEntryRateDataRequest) (int, string, error) {
	if request == nil {
		request = NewGetEntryRateDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetEntryRateDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
