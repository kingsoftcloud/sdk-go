package v20200630
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func (c *Client) GetClientRequestData(request *GetClientRequestDataRequest) (string) {
    return c.GetClientRequestDataWithContext(context.Background(), request)
}

func (c *Client) GetClientRequestDataWithContext(ctx context.Context, request *GetClientRequestDataRequest) (string) {
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

func (c *Client) GetServerData(request *GetServerDataRequest) (string) {
    return c.GetServerDataWithContext(context.Background(), request)
}

func (c *Client) GetServerDataWithContext(ctx context.Context, request *GetServerDataRequest) (string) {
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

func (c *Client) GetDomainRankingListData(request *GetDomainRankingListDataRequest) (string) {
    return c.GetDomainRankingListDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainRankingListDataWithContext(ctx context.Context, request *GetDomainRankingListDataRequest) (string) {
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

func (c *Client) GetAreaIspData(request *GetAreaIspDataRequest) (string) {
    return c.GetAreaIspDataWithContext(context.Background(), request)
}

func (c *Client) GetAreaIspDataWithContext(ctx context.Context, request *GetAreaIspDataRequest) (string) {
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

func (c *Client) GetTopReferData(request *GetTopReferDataRequest) (string) {
    return c.GetTopReferDataWithContext(context.Background(), request)
}

func (c *Client) GetTopReferDataWithContext(ctx context.Context, request *GetTopReferDataRequest) (string) {
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

func (c *Client) GetTopUrlData(request *GetTopUrlDataRequest) (string) {
    return c.GetTopUrlDataWithContext(context.Background(), request)
}

func (c *Client) GetTopUrlDataWithContext(ctx context.Context, request *GetTopUrlDataRequest) (string) {
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

func (c *Client) GetRealTimeHitRateData(request *GetRealTimeHitRateDataRequest) (string) {
    return c.GetRealTimeHitRateDataWithContext(context.Background(), request)
}

func (c *Client) GetRealTimeHitRateDataWithContext(ctx context.Context, request *GetRealTimeHitRateDataRequest) (string) {
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

func (c *Client) GetReqHitRateData(request *GetReqHitRateDataRequest) (string) {
    return c.GetReqHitRateDataWithContext(context.Background(), request)
}

func (c *Client) GetReqHitRateDataWithContext(ctx context.Context, request *GetReqHitRateDataRequest) (string) {
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

func (c *Client) GetFlowHitRateData(request *GetFlowHitRateDataRequest) (string) {
    return c.GetFlowHitRateDataWithContext(context.Background(), request)
}

func (c *Client) GetFlowHitRateDataWithContext(ctx context.Context, request *GetFlowHitRateDataRequest) (string) {
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

func (c *Client) GetDomainRequestPeriodRatioData(request *GetDomainRequestPeriodRatioDataRequest) (string) {
    return c.GetDomainRequestPeriodRatioDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainRequestPeriodRatioDataWithContext(ctx context.Context, request *GetDomainRequestPeriodRatioDataRequest) (string) {
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

func (c *Client) GetUvData(request *GetUvDataRequest) (string) {
    return c.GetUvDataWithContext(context.Background(), request)
}

func (c *Client) GetUvDataWithContext(ctx context.Context, request *GetUvDataRequest) (string) {
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

func (c *Client) GetTopIpData(request *GetTopIpDataRequest) (string) {
    return c.GetTopIpDataWithContext(context.Background(), request)
}

func (c *Client) GetTopIpDataWithContext(ctx context.Context, request *GetTopIpDataRequest) (string) {
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

func (c *Client) GetSrcDomainHttpCodeDetailedData(request *GetSrcDomainHttpCodeDetailedDataRequest) (string) {
    return c.GetSrcDomainHttpCodeDetailedDataWithContext(context.Background(), request)
}

func (c *Client) GetSrcDomainHttpCodeDetailedDataWithContext(ctx context.Context, request *GetSrcDomainHttpCodeDetailedDataRequest) (string) {
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

func (c *Client) GetSrcDomainHttpCodeData(request *GetSrcDomainHttpCodeDataRequest) (string) {
    return c.GetSrcDomainHttpCodeDataWithContext(context.Background(), request)
}

func (c *Client) GetSrcDomainHttpCodeDataWithContext(ctx context.Context, request *GetSrcDomainHttpCodeDataRequest) (string) {
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

func (c *Client) GetDomainHttpCodeDetailedData(request *GetDomainHttpCodeDetailedDataRequest) (string) {
    return c.GetDomainHttpCodeDetailedDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainHttpCodeDetailedDataWithContext(ctx context.Context, request *GetDomainHttpCodeDetailedDataRequest) (string) {
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

func (c *Client) GetDomainHttpCodeData(request *GetDomainHttpCodeDataRequest) (string) {
    return c.GetDomainHttpCodeDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainHttpCodeDataWithContext(ctx context.Context, request *GetDomainHttpCodeDataRequest) (string) {
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

func (c *Client) GetEntryRateData(request *GetEntryRateDataRequest) (string) {
    return c.GetEntryRateDataWithContext(context.Background(), request)
}

func (c *Client) GetEntryRateDataWithContext(ctx context.Context, request *GetEntryRateDataRequest) (string) {
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


