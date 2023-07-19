package v20220101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2022-01-01"

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

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
    request = &CreateAlarmPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
    return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
    response = &CreateAlarmPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (string) {
    return c.CreateAlarmPolicyWithContext(context.Background(), request)
}

func (c *Client) CreateAlarmPolicyWithContext(ctx context.Context, request *CreateAlarmPolicyRequest) (string) {
    if request == nil {
        request = NewCreateAlarmPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewCreateAlarmPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
    request = &DeleteAlarmPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
    return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
    response = &DeleteAlarmPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (string) {
    return c.DeleteAlarmPolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteAlarmPolicyWithContext(ctx context.Context, request *DeleteAlarmPolicyRequest) (string) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteAlarmPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


