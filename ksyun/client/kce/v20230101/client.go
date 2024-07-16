package v20230101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2023-01-01"

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

func NewDescribeEventLogsRequest() (request *DescribeEventLogsRequest) {
    request = &DescribeEventLogsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeEventLogs")
    return
}

func NewDescribeEventLogsResponse() (response *DescribeEventLogsResponse) {
    response = &DescribeEventLogsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEventLogs(request *DescribeEventLogsRequest) (string) {
    return c.DescribeEventLogsWithContext(context.Background(), request)
}

func (c *Client) DescribeEventLogsWithContext(ctx context.Context, request *DescribeEventLogsRequest) (string) {
    if request == nil {
        request = NewDescribeEventLogsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEventLogsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateAddonInstanceRequest() (request *CreateAddonInstanceRequest) {
    request = &CreateAddonInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "CreateAddonInstance")
    return
}

func NewCreateAddonInstanceResponse() (response *CreateAddonInstanceResponse) {
    response = &CreateAddonInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateAddonInstance(request *CreateAddonInstanceRequest) (string) {
    return c.CreateAddonInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateAddonInstanceWithContext(ctx context.Context, request *CreateAddonInstanceRequest) (string) {
    if request == nil {
        request = NewCreateAddonInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateAddonInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteAddonInstanceRequest() (request *DeleteAddonInstanceRequest) {
    request = &DeleteAddonInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DeleteAddonInstance")
    return
}

func NewDeleteAddonInstanceResponse() (response *DeleteAddonInstanceResponse) {
    response = &DeleteAddonInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteAddonInstance(request *DeleteAddonInstanceRequest) (string) {
    return c.DeleteAddonInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteAddonInstanceWithContext(ctx context.Context, request *DeleteAddonInstanceRequest) (string) {
    if request == nil {
        request = NewDeleteAddonInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteAddonInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeAddonInstancesRequest() (request *DescribeAddonInstancesRequest) {
    request = &DescribeAddonInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeAddonInstances")
    return
}

func NewDescribeAddonInstancesResponse() (response *DescribeAddonInstancesResponse) {
    response = &DescribeAddonInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeAddonInstances(request *DescribeAddonInstancesRequest) (string) {
    return c.DescribeAddonInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeAddonInstancesWithContext(ctx context.Context, request *DescribeAddonInstancesRequest) (string) {
    if request == nil {
        request = NewDescribeAddonInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeAddonInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeAddonListRequest() (request *DescribeAddonListRequest) {
    request = &DescribeAddonListRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeAddonList")
    return
}

func NewDescribeAddonListResponse() (response *DescribeAddonListResponse) {
    response = &DescribeAddonListResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeAddonList(request *DescribeAddonListRequest) (string) {
    return c.DescribeAddonListWithContext(context.Background(), request)
}

func (c *Client) DescribeAddonListWithContext(ctx context.Context, request *DescribeAddonListRequest) (string) {
    if request == nil {
        request = NewDescribeAddonListRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeAddonListResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeComponentParamsRequest() (request *DescribeComponentParamsRequest) {
    request = &DescribeComponentParamsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeComponentParams")
    return
}

func NewDescribeComponentParamsResponse() (response *DescribeComponentParamsResponse) {
    response = &DescribeComponentParamsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeComponentParams(request *DescribeComponentParamsRequest) (string) {
    return c.DescribeComponentParamsWithContext(context.Background(), request)
}

func (c *Client) DescribeComponentParamsWithContext(ctx context.Context, request *DescribeComponentParamsRequest) (string) {
    if request == nil {
        request = NewDescribeComponentParamsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeComponentParamsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeNetworkRequest() (request *DescribeNetworkRequest) {
    request = &DescribeNetworkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeNetwork")
    return
}

func NewDescribeNetworkResponse() (response *DescribeNetworkResponse) {
    response = &DescribeNetworkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeNetwork(request *DescribeNetworkRequest) (string) {
    return c.DescribeNetworkWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkWithContext(ctx context.Context, request *DescribeNetworkRequest) (string) {
    if request == nil {
        request = NewDescribeNetworkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeNetworkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeNodeComponentsRequest() (request *DescribeNodeComponentsRequest) {
    request = &DescribeNodeComponentsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeNodeComponents")
    return
}

func NewDescribeNodeComponentsResponse() (response *DescribeNodeComponentsResponse) {
    response = &DescribeNodeComponentsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeNodeComponents(request *DescribeNodeComponentsRequest) (string) {
    return c.DescribeNodeComponentsWithContext(context.Background(), request)
}

func (c *Client) DescribeNodeComponentsWithContext(ctx context.Context, request *DescribeNodeComponentsRequest) (string) {
    if request == nil {
        request = NewDescribeNodeComponentsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeNodeComponentsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeComponentListRequest() (request *DescribeComponentListRequest) {
    request = &DescribeComponentListRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeComponentList")
    return
}

func NewDescribeComponentListResponse() (response *DescribeComponentListResponse) {
    response = &DescribeComponentListResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeComponentList(request *DescribeComponentListRequest) (string) {
    return c.DescribeComponentListWithContext(context.Background(), request)
}

func (c *Client) DescribeComponentListWithContext(ctx context.Context, request *DescribeComponentListRequest) (string) {
    if request == nil {
        request = NewDescribeComponentListRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeComponentListResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


