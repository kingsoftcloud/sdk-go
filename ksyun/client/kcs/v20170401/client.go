package v20170401
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2017-04-01"

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

func NewDescribeCacheReadonlyNodeRequest() (request *DescribeCacheReadonlyNodeRequest) {
    request = &DescribeCacheReadonlyNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheReadonlyNode")
    return
}

func NewDescribeCacheReadonlyNodeResponse() (response *DescribeCacheReadonlyNodeResponse) {
    response = &DescribeCacheReadonlyNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheReadonlyNode(request *DescribeCacheReadonlyNodeRequest) (string) {
    return c.DescribeCacheReadonlyNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheReadonlyNodeWithContext(ctx context.Context, request *DescribeCacheReadonlyNodeRequest) (string) {
    if request == nil {
        request = NewDescribeCacheReadonlyNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheReadonlyNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddCacheSlaveNodeRequest() (request *AddCacheSlaveNodeRequest) {
    request = &AddCacheSlaveNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AddCacheSlaveNode")
    return
}

func NewAddCacheSlaveNodeResponse() (response *AddCacheSlaveNodeResponse) {
    response = &AddCacheSlaveNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddCacheSlaveNode(request *AddCacheSlaveNodeRequest) (string) {
    return c.AddCacheSlaveNodeWithContext(context.Background(), request)
}

func (c *Client) AddCacheSlaveNodeWithContext(ctx context.Context, request *AddCacheSlaveNodeRequest) (string) {
    if request == nil {
        request = NewAddCacheSlaveNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewAddCacheSlaveNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteCacheSlaveNodeRequest() (request *DeleteCacheSlaveNodeRequest) {
    request = &DeleteCacheSlaveNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteCacheSlaveNode")
    return
}

func NewDeleteCacheSlaveNodeResponse() (response *DeleteCacheSlaveNodeResponse) {
    response = &DeleteCacheSlaveNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteCacheSlaveNode(request *DeleteCacheSlaveNodeRequest) (string) {
    return c.DeleteCacheSlaveNodeWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheSlaveNodeWithContext(ctx context.Context, request *DeleteCacheSlaveNodeRequest) (string) {
    if request == nil {
        request = NewDeleteCacheSlaveNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteCacheSlaveNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


