package v20220601
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2022-06-01"

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

func NewGetMonthConsumeRequest() (request *GetMonthConsumeRequest) {
    request = &GetMonthConsumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetMonthConsume")
    return
}

func NewGetMonthConsumeResponse() (response *GetMonthConsumeResponse) {
    response = &GetMonthConsumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetMonthConsume(request *GetMonthConsumeRequest) (string) {
    return c.GetMonthConsumeWithContext(context.Background(), request)
}

func (c *Client) GetMonthConsumeWithContext(ctx context.Context, request *GetMonthConsumeRequest) (string) {
    if request == nil {
        request = NewGetMonthConsumeRequest()
    }
    request.SetContext(ctx)

    response := NewGetMonthConsumeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetPostpayDetailConsumeRequest() (request *GetPostpayDetailConsumeRequest) {
    request = &GetPostpayDetailConsumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailConsume")
    return
}

func NewGetPostpayDetailConsumeResponse() (response *GetPostpayDetailConsumeResponse) {
    response = &GetPostpayDetailConsumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPostpayDetailConsume(request *GetPostpayDetailConsumeRequest) (string) {
    return c.GetPostpayDetailConsumeWithContext(context.Background(), request)
}

func (c *Client) GetPostpayDetailConsumeWithContext(ctx context.Context, request *GetPostpayDetailConsumeRequest) (string) {
    if request == nil {
        request = NewGetPostpayDetailConsumeRequest()
    }
    request.SetContext(ctx)

    response := NewGetPostpayDetailConsumeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


