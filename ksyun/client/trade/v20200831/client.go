package v20200831
import (
    "context"
    "fmt"
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

func (c *Client) SetRenewal(request *SetRenewalRequest) (string) {
    return c.SetRenewalWithContext(context.Background(), request)
}

func (c *Client) SetRenewalWithContext(ctx context.Context, request *SetRenewalRequest) (string) {
    if request == nil {
        request = NewSetRenewalRequest()
    }
    request.SetContext(ctx)

    response := NewSetRenewalResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


