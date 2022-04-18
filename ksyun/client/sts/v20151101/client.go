package v20151101
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2015-11-01"

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

func NewAssumeRoleRequest() (request *AssumeRoleRequest) {
    request = &AssumeRoleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRole")
    return
}

func NewAssumeRoleResponse() (response *AssumeRoleResponse) {
    response = &AssumeRoleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AssumeRole(request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    return c.AssumeRoleWithContext(context.Background(), request)
}

func (c *Client) AssumeRoleWithContext(ctx context.Context, request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    if request == nil {
        request = NewAssumeRoleRequest()
    }
    request.SetContext(ctx)

    response = NewAssumeRoleResponse()
    err = c.Send(request, response)
    return
}


