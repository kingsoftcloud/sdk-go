package v20161122
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2016-11-22"

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

func NewCreateForwardConfRequest() (request *CreateForwardConfRequest) {
    request = &CreateForwardConfRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kad", APIVersion, "CreateForwardConf")
    return
}

func NewCreateForwardConfResponse() (response *CreateForwardConfResponse) {
    response = &CreateForwardConfResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateForwardConf(request *CreateForwardConfRequest) (response *CreateForwardConfResponse, err error) {
    return c.CreateForwardConfWithContext(context.Background(), request)
}

func (c *Client) CreateForwardConfWithContext(ctx context.Context, request *CreateForwardConfRequest) (response *CreateForwardConfResponse, err error) {
    if request == nil {
        request = NewCreateForwardConfRequest()
    }
    request.SetContext(ctx)

    response = NewCreateForwardConfResponse()
    err = c.Send(request, response)
    return
}
func NewDeleteForwardConfRequest() (request *DeleteForwardConfRequest) {
    request = &DeleteForwardConfRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kad", APIVersion, "DeleteForwardConf")
    return
}

func NewDeleteForwardConfResponse() (response *DeleteForwardConfResponse) {
    response = &DeleteForwardConfResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteForwardConf(request *DeleteForwardConfRequest) (response *DeleteForwardConfResponse, err error) {
    return c.DeleteForwardConfWithContext(context.Background(), request)
}

func (c *Client) DeleteForwardConfWithContext(ctx context.Context, request *DeleteForwardConfRequest) (response *DeleteForwardConfResponse, err error) {
    if request == nil {
        request = NewDeleteForwardConfRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteForwardConfResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeForwardConfRequest() (request *DescribeForwardConfRequest) {
    request = &DescribeForwardConfRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kad", APIVersion, "DescribeForwardConf")
    return
}

func NewDescribeForwardConfResponse() (response *DescribeForwardConfResponse) {
    response = &DescribeForwardConfResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeForwardConf(request *DescribeForwardConfRequest) (response *DescribeForwardConfResponse, err error) {
    return c.DescribeForwardConfWithContext(context.Background(), request)
}

func (c *Client) DescribeForwardConfWithContext(ctx context.Context, request *DescribeForwardConfRequest) (response *DescribeForwardConfResponse, err error) {
    if request == nil {
        request = NewDescribeForwardConfRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeForwardConfResponse()
    err = c.Send(request, response)
    return
}
func NewCreateForwardSourceRequest() (request *CreateForwardSourceRequest) {
    request = &CreateForwardSourceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kad", APIVersion, "CreateForwardSource")
    return
}

func NewCreateForwardSourceResponse() (response *CreateForwardSourceResponse) {
    response = &CreateForwardSourceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateForwardSource(request *CreateForwardSourceRequest) (response *CreateForwardSourceResponse, err error) {
    return c.CreateForwardSourceWithContext(context.Background(), request)
}

func (c *Client) CreateForwardSourceWithContext(ctx context.Context, request *CreateForwardSourceRequest) (response *CreateForwardSourceResponse, err error) {
    if request == nil {
        request = NewCreateForwardSourceRequest()
    }
    request.SetContext(ctx)

    response = NewCreateForwardSourceResponse()
    err = c.Send(request, response)
    return
}
func NewDeleteForwardSourceRequest() (request *DeleteForwardSourceRequest) {
    request = &DeleteForwardSourceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kad", APIVersion, "DeleteForwardSource")
    return
}

func NewDeleteForwardSourceResponse() (response *DeleteForwardSourceResponse) {
    response = &DeleteForwardSourceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteForwardSource(request *DeleteForwardSourceRequest) (response *DeleteForwardSourceResponse, err error) {
    return c.DeleteForwardSourceWithContext(context.Background(), request)
}

func (c *Client) DeleteForwardSourceWithContext(ctx context.Context, request *DeleteForwardSourceRequest) (response *DeleteForwardSourceResponse, err error) {
    if request == nil {
        request = NewDeleteForwardSourceRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteForwardSourceResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeForwardSourceRequest() (request *DescribeForwardSourceRequest) {
    request = &DescribeForwardSourceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kad", APIVersion, "DescribeForwardSource")
    return
}

func NewDescribeForwardSourceResponse() (response *DescribeForwardSourceResponse) {
    response = &DescribeForwardSourceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeForwardSource(request *DescribeForwardSourceRequest) (response *DescribeForwardSourceResponse, err error) {
    return c.DescribeForwardSourceWithContext(context.Background(), request)
}

func (c *Client) DescribeForwardSourceWithContext(ctx context.Context, request *DescribeForwardSourceRequest) (response *DescribeForwardSourceResponse, err error) {
    if request == nil {
        request = NewDescribeForwardSourceRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeForwardSourceResponse()
    err = c.Send(request, response)
    return
}


