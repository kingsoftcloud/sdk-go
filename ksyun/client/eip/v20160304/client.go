package v20160304
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2016-03-04"

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

func NewGetLinesRequest() (request *GetLinesRequest) {
    request = &GetLinesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "GetLines")
    return
}

func NewGetLinesResponse() (response *GetLinesResponse) {
    response = &GetLinesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetLines(request *GetLinesRequest) (string) {
    return c.GetLinesWithContext(context.Background(), request)
}

func (c *Client) GetLinesWithContext(ctx context.Context, request *GetLinesRequest) (string) {
    if request == nil {
        request = NewGetLinesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGetLinesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
    request = &DescribeAddressesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "DescribeAddresses")
    return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
    response = &DescribeAddressesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (string) {
    return c.DescribeAddressesWithContext(context.Background(), request)
}

func (c *Client) DescribeAddressesWithContext(ctx context.Context, request *DescribeAddressesRequest) (string) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeAddressesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAllocateAddressRequest() (request *AllocateAddressRequest) {
    request = &AllocateAddressRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "AllocateAddress")
    return
}

func NewAllocateAddressResponse() (response *AllocateAddressResponse) {
    response = &AllocateAddressResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AllocateAddress(request *AllocateAddressRequest) (string) {
    return c.AllocateAddressWithContext(context.Background(), request)
}

func (c *Client) AllocateAddressWithContext(ctx context.Context, request *AllocateAddressRequest) (string) {
    if request == nil {
        request = NewAllocateAddressRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAllocateAddressResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewReleaseAddressRequest() (request *ReleaseAddressRequest) {
    request = &ReleaseAddressRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "ReleaseAddress")
    return
}

func NewReleaseAddressResponse() (response *ReleaseAddressResponse) {
    response = &ReleaseAddressResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReleaseAddress(request *ReleaseAddressRequest) (string) {
    return c.ReleaseAddressWithContext(context.Background(), request)
}

func (c *Client) ReleaseAddressWithContext(ctx context.Context, request *ReleaseAddressRequest) (string) {
    if request == nil {
        request = NewReleaseAddressRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewReleaseAddressResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
    request = &AssociateAddressRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "AssociateAddress")
    return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
    response = &AssociateAddressResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AssociateAddress(request *AssociateAddressRequest) (string) {
    return c.AssociateAddressWithContext(context.Background(), request)
}

func (c *Client) AssociateAddressWithContext(ctx context.Context, request *AssociateAddressRequest) (string) {
    if request == nil {
        request = NewAssociateAddressRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAssociateAddressResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
    request = &DisassociateAddressRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "DisassociateAddress")
    return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
    response = &DisassociateAddressResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (string) {
    return c.DisassociateAddressWithContext(context.Background(), request)
}

func (c *Client) DisassociateAddressWithContext(ctx context.Context, request *DisassociateAddressRequest) (string) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDisassociateAddressResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyAddressRequest() (request *ModifyAddressRequest) {
    request = &ModifyAddressRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eip", APIVersion, "ModifyAddress")
    return
}

func NewModifyAddressResponse() (response *ModifyAddressResponse) {
    response = &ModifyAddressResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyAddress(request *ModifyAddressRequest) (string) {
    return c.ModifyAddressWithContext(context.Background(), request)
}

func (c *Client) ModifyAddressWithContext(ctx context.Context, request *ModifyAddressRequest) (string) {
    if request == nil {
        request = NewModifyAddressRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyAddressResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


