package v20180601
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2018-06-01"

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

func NewGetMonthBillRequest() (request *GetMonthBillRequest) {
    request = &GetMonthBillRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetMonthBill")
    return
}

func NewGetMonthBillResponse() (response *GetMonthBillResponse) {
    response = &GetMonthBillResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetMonthBill(request *GetMonthBillRequest) (response *GetMonthBillResponse, err error) {
    return c.GetMonthBillWithContext(context.Background(), request)
}

func (c *Client) GetMonthBillWithContext(ctx context.Context, request *GetMonthBillRequest) (response *GetMonthBillResponse, err error) {
    if request == nil {
        request = NewGetMonthBillRequest()
    }
    request.SetContext(ctx)

    response = NewGetMonthBillResponse()
    err = c.Send(request, response)
    return
}
func NewGetPostpayDetailBillRequest() (request *GetPostpayDetailBillRequest) {
    request = &GetPostpayDetailBillRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailBill")
    return
}

func NewGetPostpayDetailBillResponse() (response *GetPostpayDetailBillResponse) {
    response = &GetPostpayDetailBillResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPostpayDetailBill(request *GetPostpayDetailBillRequest) (response *GetPostpayDetailBillResponse, err error) {
    return c.GetPostpayDetailBillWithContext(context.Background(), request)
}

func (c *Client) GetPostpayDetailBillWithContext(ctx context.Context, request *GetPostpayDetailBillRequest) (response *GetPostpayDetailBillResponse, err error) {
    if request == nil {
        request = NewGetPostpayDetailBillRequest()
    }
    request.SetContext(ctx)

    response = NewGetPostpayDetailBillResponse()
    err = c.Send(request, response)
    return
}
func NewGetPostpayDetailBillCSVRequest() (request *GetPostpayDetailBillCSVRequest) {
    request = &GetPostpayDetailBillCSVRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetPostpayDetailBillCSV")
    return
}

func NewGetPostpayDetailBillCSVResponse() (response *GetPostpayDetailBillCSVResponse) {
    response = &GetPostpayDetailBillCSVResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPostpayDetailBillCSV(request *GetPostpayDetailBillCSVRequest) (response *GetPostpayDetailBillCSVResponse, err error) {
    return c.GetPostpayDetailBillCSVWithContext(context.Background(), request)
}

func (c *Client) GetPostpayDetailBillCSVWithContext(ctx context.Context, request *GetPostpayDetailBillCSVRequest) (response *GetPostpayDetailBillCSVResponse, err error) {
    if request == nil {
        request = NewGetPostpayDetailBillCSVRequest()
    }
    request.SetContext(ctx)

    response = NewGetPostpayDetailBillCSVResponse()
    err = c.Send(request, response)
    return
}
func NewGetProductCodeRequest() (request *GetProductCodeRequest) {
    request = &GetProductCodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetProductCode")
    return
}

func NewGetProductCodeResponse() (response *GetProductCodeResponse) {
    response = &GetProductCodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetProductCode(request *GetProductCodeRequest) (response *GetProductCodeResponse, err error) {
    return c.GetProductCodeWithContext(context.Background(), request)
}

func (c *Client) GetProductCodeWithContext(ctx context.Context, request *GetProductCodeRequest) (response *GetProductCodeResponse, err error) {
    if request == nil {
        request = NewGetProductCodeRequest()
    }
    request.SetContext(ctx)

    response = NewGetProductCodeResponse()
    err = c.Send(request, response)
    return
}
func NewgetMonthConsumeRequest() (request *getMonthConsumeRequest) {
    request = &getMonthConsumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "getMonthConsume")
    return
}

func NewgetMonthConsumeResponse() (response *getMonthConsumeResponse) {
    response = &getMonthConsumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) getMonthConsume(request *getMonthConsumeRequest) (response *getMonthConsumeResponse, err error) {
    return c.getMonthConsumeWithContext(context.Background(), request)
}

func (c *Client) getMonthConsumeWithContext(ctx context.Context, request *getMonthConsumeRequest) (response *getMonthConsumeResponse, err error) {
    if request == nil {
        request = NewgetMonthConsumeRequest()
    }
    request.SetContext(ctx)

    response = NewgetMonthConsumeResponse()
    err = c.Send(request, response)
    return
}
func NewgetPostpayDetailConsumeRequest() (request *getPostpayDetailConsumeRequest) {
    request = &getPostpayDetailConsumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "getPostpayDetailConsume")
    return
}

func NewgetPostpayDetailConsumeResponse() (response *getPostpayDetailConsumeResponse) {
    response = &getPostpayDetailConsumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) getPostpayDetailConsume(request *getPostpayDetailConsumeRequest) (response *getPostpayDetailConsumeResponse, err error) {
    return c.getPostpayDetailConsumeWithContext(context.Background(), request)
}

func (c *Client) getPostpayDetailConsumeWithContext(ctx context.Context, request *getPostpayDetailConsumeRequest) (response *getPostpayDetailConsumeResponse, err error) {
    if request == nil {
        request = NewgetPostpayDetailConsumeRequest()
    }
    request.SetContext(ctx)

    response = NewgetPostpayDetailConsumeResponse()
    err = c.Send(request, response)
    return
}


