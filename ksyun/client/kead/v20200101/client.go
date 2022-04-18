package v20200101
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-01-01"

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

func NewDescribeBlockIpRequest() (request *DescribeBlockIpRequest) {
    request = &DescribeBlockIpRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kead", APIVersion, "DescribeBlockIp")
    return
}

func NewDescribeBlockIpResponse() (response *DescribeBlockIpResponse) {
    response = &DescribeBlockIpResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBlockIp(request *DescribeBlockIpRequest) (response *DescribeBlockIpResponse, err error) {
    return c.DescribeBlockIpWithContext(context.Background(), request)
}

func (c *Client) DescribeBlockIpWithContext(ctx context.Context, request *DescribeBlockIpRequest) (response *DescribeBlockIpResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIpRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeBlockIpResponse()
    err = c.Send(request, response)
    return
}


