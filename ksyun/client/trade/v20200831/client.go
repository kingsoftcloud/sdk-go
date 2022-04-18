package v20200831
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-08-31"

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

func NewSetRenewalRequest() (request *SetRenewalRequest) {
    request = &SetRenewalRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trade", APIVersion, "SetRenewal")
    return
}

func NewSetRenewalResponse() (response *SetRenewalResponse) {
    response = &SetRenewalResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetRenewal(request *SetRenewalRequest) (response *SetRenewalResponse, err error) {
    return c.SetRenewalWithContext(context.Background(), request)
}

func (c *Client) SetRenewalWithContext(ctx context.Context, request *SetRenewalRequest) (response *SetRenewalResponse, err error) {
    if request == nil {
        request = NewSetRenewalRequest()
    }
    request.SetContext(ctx)

    response = NewSetRenewalResponse()
    err = c.Send(request, response)
    return
}


