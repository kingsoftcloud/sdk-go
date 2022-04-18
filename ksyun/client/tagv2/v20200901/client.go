package v20200901
import (
    "context"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-09-01"

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

func NewCreateTagRequest() (request *CreateTagRequest) {
    request = &CreateTagRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "CreateTag")
    return
}

func NewCreateTagResponse() (response *CreateTagResponse) {
    response = &CreateTagResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateTag(request *CreateTagRequest) (response *CreateTagResponse, err error) {
    return c.CreateTagWithContext(context.Background(), request)
}

func (c *Client) CreateTagWithContext(ctx context.Context, request *CreateTagRequest) (response *CreateTagResponse, err error) {
    if request == nil {
        request = NewCreateTagRequest()
    }
    request.SetContext(ctx)

    response = NewCreateTagResponse()
    err = c.Send(request, response)
    return
}
func NewDeleteTagRequest() (request *DeleteTagRequest) {
    request = &DeleteTagRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "DeleteTag")
    return
}

func NewDeleteTagResponse() (response *DeleteTagResponse) {
    response = &DeleteTagResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteTag(request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
    return c.DeleteTagWithContext(context.Background(), request)
}

func (c *Client) DeleteTagWithContext(ctx context.Context, request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
    if request == nil {
        request = NewDeleteTagRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteTagResponse()
    err = c.Send(request, response)
    return
}
func NewListTagsRequest() (request *ListTagsRequest) {
    request = &ListTagsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "ListTags")
    return
}

func NewListTagsResponse() (response *ListTagsResponse) {
    response = &ListTagsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListTags(request *ListTagsRequest) (response *ListTagsResponse, err error) {
    return c.ListTagsWithContext(context.Background(), request)
}

func (c *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest) (response *ListTagsResponse, err error) {
    if request == nil {
        request = NewListTagsRequest()
    }
    request.SetContext(ctx)

    response = NewListTagsResponse()
    err = c.Send(request, response)
    return
}
func NewListTagKeysRequest() (request *ListTagKeysRequest) {
    request = &ListTagKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "ListTagKeys")
    return
}

func NewListTagKeysResponse() (response *ListTagKeysResponse) {
    response = &ListTagKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListTagKeys(request *ListTagKeysRequest) (response *ListTagKeysResponse, err error) {
    return c.ListTagKeysWithContext(context.Background(), request)
}

func (c *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest) (response *ListTagKeysResponse, err error) {
    if request == nil {
        request = NewListTagKeysRequest()
    }
    request.SetContext(ctx)

    response = NewListTagKeysResponse()
    err = c.Send(request, response)
    return
}
func NewListTagValuesRequest() (request *ListTagValuesRequest) {
    request = &ListTagValuesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "ListTagValues")
    return
}

func NewListTagValuesResponse() (response *ListTagValuesResponse) {
    response = &ListTagValuesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListTagValues(request *ListTagValuesRequest) (response *ListTagValuesResponse, err error) {
    return c.ListTagValuesWithContext(context.Background(), request)
}

func (c *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest) (response *ListTagValuesResponse, err error) {
    if request == nil {
        request = NewListTagValuesRequest()
    }
    request.SetContext(ctx)

    response = NewListTagValuesResponse()
    err = c.Send(request, response)
    return
}
func NewListResourcesRequest() (request *ListResourcesRequest) {
    request = &ListResourcesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "ListResources")
    return
}

func NewListResourcesResponse() (response *ListResourcesResponse) {
    response = &ListResourcesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListResources(request *ListResourcesRequest) (response *ListResourcesResponse, err error) {
    return c.ListResourcesWithContext(context.Background(), request)
}

func (c *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest) (response *ListResourcesResponse, err error) {
    if request == nil {
        request = NewListResourcesRequest()
    }
    request.SetContext(ctx)

    response = NewListResourcesResponse()
    err = c.Send(request, response)
    return
}
func NewReplaceResourcesTagsRequest() (request *ReplaceResourcesTagsRequest) {
    request = &ReplaceResourcesTagsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "ReplaceResourcesTags")
    return
}

func NewReplaceResourcesTagsResponse() (response *ReplaceResourcesTagsResponse) {
    response = &ReplaceResourcesTagsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReplaceResourcesTags(request *ReplaceResourcesTagsRequest) (response *ReplaceResourcesTagsResponse, err error) {
    return c.ReplaceResourcesTagsWithContext(context.Background(), request)
}

func (c *Client) ReplaceResourcesTagsWithContext(ctx context.Context, request *ReplaceResourcesTagsRequest) (response *ReplaceResourcesTagsResponse, err error) {
    if request == nil {
        request = NewReplaceResourcesTagsRequest()
    }
    request.SetContext(ctx)

    response = NewReplaceResourcesTagsResponse()
    err = c.Send(request, response)
    return
}
func NewDetachResourceTagsRequest() (request *DetachResourceTagsRequest) {
    request = &DetachResourceTagsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tagv2", APIVersion, "DetachResourceTags")
    return
}

func NewDetachResourceTagsResponse() (response *DetachResourceTagsResponse) {
    response = &DetachResourceTagsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DetachResourceTags(request *DetachResourceTagsRequest) (response *DetachResourceTagsResponse, err error) {
    return c.DetachResourceTagsWithContext(context.Background(), request)
}

func (c *Client) DetachResourceTagsWithContext(ctx context.Context, request *DetachResourceTagsRequest) (response *DetachResourceTagsResponse, err error) {
    if request == nil {
        request = NewDetachResourceTagsRequest()
    }
    request.SetContext(ctx)

    response = NewDetachResourceTagsResponse()
    err = c.Send(request, response)
    return
}


