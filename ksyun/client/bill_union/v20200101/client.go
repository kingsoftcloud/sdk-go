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

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByPayMode")
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProduct")
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
    request = &DescribeBillSummaryByProjectRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProject")
    return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
    response = &DescribeBillSummaryByProjectResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeBillSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeInstanceSummaryBillsRequest() (request *DescribeInstanceSummaryBillsRequest) {
    request = &DescribeInstanceSummaryBillsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeInstanceSummaryBills")
    return
}

func NewDescribeInstanceSummaryBillsResponse() (response *DescribeInstanceSummaryBillsResponse) {
    response = &DescribeInstanceSummaryBillsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceSummaryBills(request *DescribeInstanceSummaryBillsRequest) (response *DescribeInstanceSummaryBillsResponse, err error) {
    return c.DescribeInstanceSummaryBillsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceSummaryBillsWithContext(ctx context.Context, request *DescribeInstanceSummaryBillsRequest) (response *DescribeInstanceSummaryBillsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSummaryBillsRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeInstanceSummaryBillsResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeProductCodeRequest() (request *DescribeProductCodeRequest) {
    request = &DescribeProductCodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeProductCode")
    return
}

func NewDescribeProductCodeResponse() (response *DescribeProductCodeResponse) {
    response = &DescribeProductCodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeProductCode(request *DescribeProductCodeRequest) (response *DescribeProductCodeResponse, err error) {
    return c.DescribeProductCodeWithContext(context.Background(), request)
}

func (c *Client) DescribeProductCodeWithContext(ctx context.Context, request *DescribeProductCodeRequest) (response *DescribeProductCodeResponse, err error) {
    if request == nil {
        request = NewDescribeProductCodeRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeProductCodeResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeSplitItemBillDetailsRequest() (request *DescribeSplitItemBillDetailsRequest) {
    request = &DescribeSplitItemBillDetailsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemBillDetails")
    return
}

func NewDescribeSplitItemBillDetailsResponse() (response *DescribeSplitItemBillDetailsResponse) {
    response = &DescribeSplitItemBillDetailsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSplitItemBillDetails(request *DescribeSplitItemBillDetailsRequest) (response *DescribeSplitItemBillDetailsResponse, err error) {
    return c.DescribeSplitItemBillDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeSplitItemBillDetailsWithContext(ctx context.Context, request *DescribeSplitItemBillDetailsRequest) (response *DescribeSplitItemBillDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeSplitItemBillDetailsRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeSplitItemBillDetailsResponse()
    err = c.Send(request, response)
    return
}


