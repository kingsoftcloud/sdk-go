package v20211209
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2021-12-09"

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

func NewDescribeCostBillRequest() (request *DescribeCostBillRequest) {
    request = &DescribeCostBillRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeCostBill")
    return
}

func NewDescribeCostBillResponse() (response *DescribeCostBillResponse) {
    response = &DescribeCostBillResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCostBill(request *DescribeCostBillRequest) (response *DescribeCostBillResponse, err error) {
    return c.DescribeCostBillWithContext(context.Background(), request)
}

func (c *Client) DescribeCostBillWithContext(ctx context.Context, request *DescribeCostBillRequest) (response *DescribeCostBillResponse, err error) {
    if request == nil {
        request = NewDescribeCostBillRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeCostBillResponse()
    err = c.Send(request, response)
    return
}


