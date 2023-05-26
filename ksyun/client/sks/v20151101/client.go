package v20151101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2015-11-01"

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

func NewCreateKeyRequest() (request *CreateKeyRequest) {
    request = &CreateKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sks", APIVersion, "CreateKey")
    return
}

func NewCreateKeyResponse() (response *CreateKeyResponse) {
    response = &CreateKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateKey(request *CreateKeyRequest) (string) {
    return c.CreateKeyWithContext(context.Background(), request)
}

func (c *Client) CreateKeyWithContext(ctx context.Context, request *CreateKeyRequest) (string) {
    if request == nil {
        request = NewCreateKeyRequest()
    }
    request.SetContext(ctx)

    response := NewCreateKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewImportKeyRequest() (request *ImportKeyRequest) {
    request = &ImportKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sks", APIVersion, "ImportKey")
    return
}

func NewImportKeyResponse() (response *ImportKeyResponse) {
    response = &ImportKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ImportKey(request *ImportKeyRequest) (string) {
    return c.ImportKeyWithContext(context.Background(), request)
}

func (c *Client) ImportKeyWithContext(ctx context.Context, request *ImportKeyRequest) (string) {
    if request == nil {
        request = NewImportKeyRequest()
    }
    request.SetContext(ctx)

    response := NewImportKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteKeyRequest() (request *DeleteKeyRequest) {
    request = &DeleteKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sks", APIVersion, "DeleteKey")
    return
}

func NewDeleteKeyResponse() (response *DeleteKeyResponse) {
    response = &DeleteKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteKey(request *DeleteKeyRequest) (string) {
    return c.DeleteKeyWithContext(context.Background(), request)
}

func (c *Client) DeleteKeyWithContext(ctx context.Context, request *DeleteKeyRequest) (string) {
    if request == nil {
        request = NewDeleteKeyRequest()
    }
    request.SetContext(ctx)

    response := NewDeleteKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyKeyRequest() (request *ModifyKeyRequest) {
    request = &ModifyKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sks", APIVersion, "ModifyKey")
    return
}

func NewModifyKeyResponse() (response *ModifyKeyResponse) {
    response = &ModifyKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyKey(request *ModifyKeyRequest) (string) {
    return c.ModifyKeyWithContext(context.Background(), request)
}

func (c *Client) ModifyKeyWithContext(ctx context.Context, request *ModifyKeyRequest) (string) {
    if request == nil {
        request = NewModifyKeyRequest()
    }
    request.SetContext(ctx)

    response := NewModifyKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeKeysRequest() (request *DescribeKeysRequest) {
    request = &DescribeKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sks", APIVersion, "DescribeKeys")
    return
}

func NewDescribeKeysResponse() (response *DescribeKeysResponse) {
    response = &DescribeKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeKeys(request *DescribeKeysRequest) (string) {
    return c.DescribeKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeKeysWithContext(ctx context.Context, request *DescribeKeysRequest) (string) {
    if request == nil {
        request = NewDescribeKeysRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


