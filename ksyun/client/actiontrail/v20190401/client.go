package v20190401
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2019-04-01"

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

func NewListOperateLogsRequest() (request *ListOperateLogsRequest) {
    request = &ListOperateLogsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("actiontrail", APIVersion, "ListOperateLogs")
    return
}

func NewListOperateLogsResponse() (response *ListOperateLogsResponse) {
    response = &ListOperateLogsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListOperateLogs(request *ListOperateLogsRequest) (response *ListOperateLogsResponse, err error) {
    return c.ListOperateLogsWithContext(context.Background(), request)
}

func (c *Client) ListOperateLogsWithContext(ctx context.Context, request *ListOperateLogsRequest) (response *ListOperateLogsResponse, err error) {
    if request == nil {
        request = NewListOperateLogsRequest()
    }
    request.SetContext(ctx)

    response = NewListOperateLogsResponse()
    err = c.Send(request, response)
    return
}


