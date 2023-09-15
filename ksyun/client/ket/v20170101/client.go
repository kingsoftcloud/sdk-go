package v20170101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2017-01-01"

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

func NewPresetRequest() (request *PresetRequest) {
    request = &PresetRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ket", APIVersion, "Preset")
    return
}

func NewPresetResponse() (response *PresetResponse) {
    response = &PresetResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) Preset(request *PresetRequest) (string) {
    return c.PresetWithContext(context.Background(), request)
}

func (c *Client) PresetWithContext(ctx context.Context, request *PresetRequest) (string) {
    if request == nil {
        request = NewPresetRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewPresetResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdatePresetRequest() (request *UpdatePresetRequest) {
    request = &UpdatePresetRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ket", APIVersion, "UpdatePreset")
    return
}

func NewUpdatePresetResponse() (response *UpdatePresetResponse) {
    response = &UpdatePresetResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdatePreset(request *UpdatePresetRequest) (string) {
    return c.UpdatePresetWithContext(context.Background(), request)
}

func (c *Client) UpdatePresetWithContext(ctx context.Context, request *UpdatePresetRequest) (string) {
    if request == nil {
        request = NewUpdatePresetRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewUpdatePresetResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDelPresetRequest() (request *DelPresetRequest) {
    request = &DelPresetRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ket", APIVersion, "DelPreset")
    return
}

func NewDelPresetResponse() (response *DelPresetResponse) {
    response = &DelPresetResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DelPreset(request *DelPresetRequest) (string) {
    return c.DelPresetWithContext(context.Background(), request)
}

func (c *Client) DelPresetWithContext(ctx context.Context, request *DelPresetRequest) (string) {
    if request == nil {
        request = NewDelPresetRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDelPresetResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetPresetListRequest() (request *GetPresetListRequest) {
    request = &GetPresetListRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ket", APIVersion, "GetPresetList")
    return
}

func NewGetPresetListResponse() (response *GetPresetListResponse) {
    response = &GetPresetListResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPresetList(request *GetPresetListRequest) (string) {
    return c.GetPresetListWithContext(context.Background(), request)
}

func (c *Client) GetPresetListWithContext(ctx context.Context, request *GetPresetListRequest) (string) {
    if request == nil {
        request = NewGetPresetListRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGetPresetListResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetPresetDetailRequest() (request *GetPresetDetailRequest) {
    request = &GetPresetDetailRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ket", APIVersion, "GetPresetDetail")
    return
}

func NewGetPresetDetailResponse() (response *GetPresetDetailResponse) {
    response = &GetPresetDetailResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetPresetDetail(request *GetPresetDetailRequest) (string) {
    return c.GetPresetDetailWithContext(context.Background(), request)
}

func (c *Client) GetPresetDetailWithContext(ctx context.Context, request *GetPresetDetailRequest) (string) {
    if request == nil {
        request = NewGetPresetDetailRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGetPresetDetailResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetStreamTranListRequest() (request *GetStreamTranListRequest) {
    request = &GetStreamTranListRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ket", APIVersion, "GetStreamTranList")
    return
}

func NewGetStreamTranListResponse() (response *GetStreamTranListResponse) {
    response = &GetStreamTranListResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetStreamTranList(request *GetStreamTranListRequest) (string) {
    return c.GetStreamTranListWithContext(context.Background(), request)
}

func (c *Client) GetStreamTranListWithContext(ctx context.Context, request *GetStreamTranListRequest) (string) {
    if request == nil {
        request = NewGetStreamTranListRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGetStreamTranListResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


