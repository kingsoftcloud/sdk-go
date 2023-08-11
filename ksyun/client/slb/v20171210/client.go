package v20171210
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2017-12-10"

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

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeLoadBalancers")
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (string) {
    return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) (string) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeLoadBalancersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


