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

func NewCreateKeyRequest() (request *CreateKeyRequest) {
    request = &CreateKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kkms", APIVersion, "CreateKey")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateKeyResponse()
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
    request.Init().WithApiInfo("kkms", APIVersion, "ModifyKey")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyKeyStateRequest() (request *ModifyKeyStateRequest) {
    request = &ModifyKeyStateRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kkms", APIVersion, "ModifyKeyState")
    return
}

func NewModifyKeyStateResponse() (response *ModifyKeyStateResponse) {
    response = &ModifyKeyStateResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyKeyState(request *ModifyKeyStateRequest) (string) {
    return c.ModifyKeyStateWithContext(context.Background(), request)
}

func (c *Client) ModifyKeyStateWithContext(ctx context.Context, request *ModifyKeyStateRequest) (string) {
    if request == nil {
        request = NewModifyKeyStateRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyKeyStateResponse()
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
    request.Init().WithApiInfo("kkms", APIVersion, "DeleteKey")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteKeyResponse()
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
    request.Init().WithApiInfo("kkms", APIVersion, "DescribeKeys")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewEncryptRequest() (request *EncryptRequest) {
    request = &EncryptRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kkms", APIVersion, "Encrypt")
    return
}

func NewEncryptResponse() (response *EncryptResponse) {
    response = &EncryptResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) Encrypt(request *EncryptRequest) (string) {
    return c.EncryptWithContext(context.Background(), request)
}

func (c *Client) EncryptWithContext(ctx context.Context, request *EncryptRequest) (string) {
    if request == nil {
        request = NewEncryptRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewEncryptResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDecryptRequest() (request *DecryptRequest) {
    request = &DecryptRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kkms", APIVersion, "Decrypt")
    return
}

func NewDecryptResponse() (response *DecryptResponse) {
    response = &DecryptResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) Decrypt(request *DecryptRequest) (string) {
    return c.DecryptWithContext(context.Background(), request)
}

func (c *Client) DecryptWithContext(ctx context.Context, request *DecryptRequest) (string) {
    if request == nil {
        request = NewDecryptRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDecryptResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGenerateDataKeyRequest() (request *GenerateDataKeyRequest) {
    request = &GenerateDataKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kkms", APIVersion, "GenerateDataKey")
    return
}

func NewGenerateDataKeyResponse() (response *GenerateDataKeyResponse) {
    response = &GenerateDataKeyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GenerateDataKey(request *GenerateDataKeyRequest) (string) {
    return c.GenerateDataKeyWithContext(context.Background(), request)
}

func (c *Client) GenerateDataKeyWithContext(ctx context.Context, request *GenerateDataKeyRequest) (string) {
    if request == nil {
        request = NewGenerateDataKeyRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGenerateDataKeyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


