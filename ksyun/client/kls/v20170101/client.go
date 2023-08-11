package v20170101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func (c *Client) ListStreamDurations(request *ListStreamDurationsRequest) (string) {
    return c.ListStreamDurationsWithContext(context.Background(), request)
}

func (c *Client) ListStreamDurationsWithContext(ctx context.Context, request *ListStreamDurationsRequest) (string) {
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

func (c *Client) ListHistoryPubStreamsErrInfo(request *ListHistoryPubStreamsErrInfoRequest) (string) {
    return c.ListHistoryPubStreamsErrInfoWithContext(context.Background(), request)
}

func (c *Client) ListHistoryPubStreamsErrInfoWithContext(ctx context.Context, request *ListHistoryPubStreamsErrInfoRequest) (string) {
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

func (c *Client) ListHistoryPubStreamsInfo(request *ListHistoryPubStreamsInfoRequest) (string) {
    return c.ListHistoryPubStreamsInfoWithContext(context.Background(), request)
}

func (c *Client) ListHistoryPubStreamsInfoWithContext(ctx context.Context, request *ListHistoryPubStreamsInfoRequest) (string) {
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

func (c *Client) ForbidStream(request *ForbidStreamRequest) (string) {
    return c.ForbidStreamWithContext(context.Background(), request)
}

func (c *Client) ForbidStreamWithContext(ctx context.Context, request *ForbidStreamRequest) (string) {
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

func (c *Client) ResumeStream(request *ResumeStreamRequest) (string) {
    return c.ResumeStreamWithContext(context.Background(), request)
}

func (c *Client) ResumeStreamWithContext(ctx context.Context, request *ResumeStreamRequest) (string) {
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

func (c *Client) GetBlacklist(request *GetBlacklistRequest) (string) {
    return c.GetBlacklistWithContext(context.Background(), request)
}

func (c *Client) GetBlacklistWithContext(ctx context.Context, request *GetBlacklistRequest) (string) {
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

func (c *Client) CheckBlacklist(request *CheckBlacklistRequest) (string) {
    return c.CheckBlacklistWithContext(context.Background(), request)
}

func (c *Client) CheckBlacklistWithContext(ctx context.Context, request *CheckBlacklistRequest) (string) {
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

func (c *Client) ListRealtimeStreamsInfo(request *ListRealtimeStreamsInfoRequest) (string) {
    return c.ListRealtimeStreamsInfoWithContext(context.Background(), request)
}

func (c *Client) ListRealtimeStreamsInfoWithContext(ctx context.Context, request *ListRealtimeStreamsInfoRequest) (string) {
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


