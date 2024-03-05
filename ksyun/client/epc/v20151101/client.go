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

func NewCreateEpcRequest() (request *CreateEpcRequest) {
    request = &CreateEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateEpc")
    return
}

func NewCreateEpcResponse() (response *CreateEpcResponse) {
    response = &CreateEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateEpc(request *CreateEpcRequest) (string) {
    return c.CreateEpcWithContext(context.Background(), request)
}

func (c *Client) CreateEpcWithContext(ctx context.Context, request *CreateEpcRequest) (string) {
    if request == nil {
        request = NewCreateEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStartEpcRequest() (request *StartEpcRequest) {
    request = &StartEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "StartEpc")
    return
}

func NewStartEpcResponse() (response *StartEpcResponse) {
    response = &StartEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StartEpc(request *StartEpcRequest) (string) {
    return c.StartEpcWithContext(context.Background(), request)
}

func (c *Client) StartEpcWithContext(ctx context.Context, request *StartEpcRequest) (string) {
    if request == nil {
        request = NewStartEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStartEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRebootEpcRequest() (request *RebootEpcRequest) {
    request = &RebootEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "RebootEpc")
    return
}

func NewRebootEpcResponse() (response *RebootEpcResponse) {
    response = &RebootEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RebootEpc(request *RebootEpcRequest) (string) {
    return c.RebootEpcWithContext(context.Background(), request)
}

func (c *Client) RebootEpcWithContext(ctx context.Context, request *RebootEpcRequest) (string) {
    if request == nil {
        request = NewRebootEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRebootEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteEpcRequest() (request *DeleteEpcRequest) {
    request = &DeleteEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DeleteEpc")
    return
}

func NewDeleteEpcResponse() (response *DeleteEpcResponse) {
    response = &DeleteEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteEpc(request *DeleteEpcRequest) (string) {
    return c.DeleteEpcWithContext(context.Background(), request)
}

func (c *Client) DeleteEpcWithContext(ctx context.Context, request *DeleteEpcRequest) (string) {
    if request == nil {
        request = NewDeleteEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewReinstallEpcRequest() (request *ReinstallEpcRequest) {
    request = &ReinstallEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ReinstallEpc")
    return
}

func NewReinstallEpcResponse() (response *ReinstallEpcResponse) {
    response = &ReinstallEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReinstallEpc(request *ReinstallEpcRequest) (string) {
    return c.ReinstallEpcWithContext(context.Background(), request)
}

func (c *Client) ReinstallEpcWithContext(ctx context.Context, request *ReinstallEpcRequest) (string) {
    if request == nil {
        request = NewReinstallEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewReinstallEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
    request = &ModifySecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifySecurityGroup")
    return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
    response = &ModifySecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) (string) {
    return c.ModifySecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupWithContext(ctx context.Context, request *ModifySecurityGroupRequest) (string) {
    if request == nil {
        request = NewModifySecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifySecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateKeyRequest() (request *CreateKeyRequest) {
    request = &CreateKeyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateKey")
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
func NewDescribeEpcsRequest() (request *DescribeEpcsRequest) {
    request = &DescribeEpcsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcs")
    return
}

func NewDescribeEpcsResponse() (response *DescribeEpcsResponse) {
    response = &DescribeEpcsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcs(request *DescribeEpcsRequest) (string) {
    return c.DescribeEpcsWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcsWithContext(ctx context.Context, request *DescribeEpcsRequest) (string) {
    if request == nil {
        request = NewDescribeEpcsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetDynamicCodeRequest() (request *GetDynamicCodeRequest) {
    request = &GetDynamicCodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "GetDynamicCode")
    return
}

func NewGetDynamicCodeResponse() (response *GetDynamicCodeResponse) {
    response = &GetDynamicCodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetDynamicCode(request *GetDynamicCodeRequest) (string) {
    return c.GetDynamicCodeWithContext(context.Background(), request)
}

func (c *Client) GetDynamicCodeWithContext(ctx context.Context, request *GetDynamicCodeRequest) (string) {
    if request == nil {
        request = NewGetDynamicCodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGetDynamicCodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeVpnsRequest() (request *DescribeVpnsRequest) {
    request = &DescribeVpnsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeVpns")
    return
}

func NewDescribeVpnsResponse() (response *DescribeVpnsResponse) {
    response = &DescribeVpnsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeVpns(request *DescribeVpnsRequest) (string) {
    return c.DescribeVpnsWithContext(context.Background(), request)
}

func (c *Client) DescribeVpnsWithContext(ctx context.Context, request *DescribeVpnsRequest) (string) {
    if request == nil {
        request = NewDescribeVpnsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeVpnsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateImage")
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateImage(request *CreateImageRequest) (string) {
    return c.CreateImageWithContext(context.Background(), request)
}

func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (string) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyImageRequest() (request *ModifyImageRequest) {
    request = &ModifyImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyImage")
    return
}

func NewModifyImageResponse() (response *ModifyImageResponse) {
    response = &ModifyImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyImage(request *ModifyImageRequest) (string) {
    return c.ModifyImageWithContext(context.Background(), request)
}

func (c *Client) ModifyImageWithContext(ctx context.Context, request *ModifyImageRequest) (string) {
    if request == nil {
        request = NewModifyImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteImageRequest() (request *DeleteImageRequest) {
    request = &DeleteImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DeleteImage")
    return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
    response = &DeleteImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteImage(request *DeleteImageRequest) (string) {
    return c.DeleteImageWithContext(context.Background(), request)
}

func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) (string) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeImages")
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) (string) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (string) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeImagesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDnsRequest() (request *ModifyDnsRequest) {
    request = &ModifyDnsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyDns")
    return
}

func NewModifyDnsResponse() (response *ModifyDnsResponse) {
    response = &ModifyDnsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDns(request *ModifyDnsRequest) (string) {
    return c.ModifyDnsWithContext(context.Background(), request)
}

func (c *Client) ModifyDnsWithContext(ctx context.Context, request *ModifyDnsRequest) (string) {
    if request == nil {
        request = NewModifyDnsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyDnsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
    request = &ModifyNetworkInterfaceAttributeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyNetworkInterfaceAttribute")
    return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
    response = &ModifyNetworkInterfaceAttributeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) (string) {
    return c.ModifyNetworkInterfaceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) (string) {
    if request == nil {
        request = NewModifyNetworkInterfaceAttributeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyNetworkInterfaceAttributeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribePhysicalMonitorRequest() (request *DescribePhysicalMonitorRequest) {
    request = &DescribePhysicalMonitorRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribePhysicalMonitor")
    return
}

func NewDescribePhysicalMonitorResponse() (response *DescribePhysicalMonitorResponse) {
    response = &DescribePhysicalMonitorResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribePhysicalMonitor(request *DescribePhysicalMonitorRequest) (string) {
    return c.DescribePhysicalMonitorWithContext(context.Background(), request)
}

func (c *Client) DescribePhysicalMonitorWithContext(ctx context.Context, request *DescribePhysicalMonitorRequest) (string) {
    if request == nil {
        request = NewDescribePhysicalMonitorRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribePhysicalMonitorResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcManagementsRequest() (request *DescribeEpcManagementsRequest) {
    request = &DescribeEpcManagementsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcManagements")
    return
}

func NewDescribeEpcManagementsResponse() (response *DescribeEpcManagementsResponse) {
    response = &DescribeEpcManagementsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcManagements(request *DescribeEpcManagementsRequest) (string) {
    return c.DescribeEpcManagementsWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcManagementsWithContext(ctx context.Context, request *DescribeEpcManagementsRequest) (string) {
    if request == nil {
        request = NewDescribeEpcManagementsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcManagementsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeRemoteManagementsRequest() (request *DescribeRemoteManagementsRequest) {
    request = &DescribeRemoteManagementsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeRemoteManagements")
    return
}

func NewDescribeRemoteManagementsResponse() (response *DescribeRemoteManagementsResponse) {
    response = &DescribeRemoteManagementsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeRemoteManagements(request *DescribeRemoteManagementsRequest) (string) {
    return c.DescribeRemoteManagementsWithContext(context.Background(), request)
}

func (c *Client) DescribeRemoteManagementsWithContext(ctx context.Context, request *DescribeRemoteManagementsRequest) (string) {
    if request == nil {
        request = NewDescribeRemoteManagementsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeRemoteManagementsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStopEpcRequest() (request *StopEpcRequest) {
    request = &StopEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "StopEpc")
    return
}

func NewStopEpcResponse() (response *StopEpcResponse) {
    response = &StopEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StopEpc(request *StopEpcRequest) (string) {
    return c.StopEpcWithContext(context.Background(), request)
}

func (c *Client) StopEpcWithContext(ctx context.Context, request *StopEpcRequest) (string) {
    if request == nil {
        request = NewStopEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewStopEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyEpcRequest() (request *ModifyEpcRequest) {
    request = &ModifyEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyEpc")
    return
}

func NewModifyEpcResponse() (response *ModifyEpcResponse) {
    response = &ModifyEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyEpc(request *ModifyEpcRequest) (string) {
    return c.ModifyEpcWithContext(context.Background(), request)
}

func (c *Client) ModifyEpcWithContext(ctx context.Context, request *ModifyEpcRequest) (string) {
    if request == nil {
        request = NewModifyEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyRemoteManagementRequest() (request *ModifyRemoteManagementRequest) {
    request = &ModifyRemoteManagementRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyRemoteManagement")
    return
}

func NewModifyRemoteManagementResponse() (response *ModifyRemoteManagementResponse) {
    response = &ModifyRemoteManagementResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyRemoteManagement(request *ModifyRemoteManagementRequest) (string) {
    return c.ModifyRemoteManagementWithContext(context.Background(), request)
}

func (c *Client) ModifyRemoteManagementWithContext(ctx context.Context, request *ModifyRemoteManagementRequest) (string) {
    if request == nil {
        request = NewModifyRemoteManagementRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyRemoteManagementResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateRemoteManagementRequest() (request *CreateRemoteManagementRequest) {
    request = &CreateRemoteManagementRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateRemoteManagement")
    return
}

func NewCreateRemoteManagementResponse() (response *CreateRemoteManagementResponse) {
    response = &CreateRemoteManagementResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateRemoteManagement(request *CreateRemoteManagementRequest) (string) {
    return c.CreateRemoteManagementWithContext(context.Background(), request)
}

func (c *Client) CreateRemoteManagementWithContext(ctx context.Context, request *CreateRemoteManagementRequest) (string) {
    if request == nil {
        request = NewCreateRemoteManagementRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateRemoteManagementResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewReinstallCustomerEpcRequest() (request *ReinstallCustomerEpcRequest) {
    request = &ReinstallCustomerEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ReinstallCustomerEpc")
    return
}

func NewReinstallCustomerEpcResponse() (response *ReinstallCustomerEpcResponse) {
    response = &ReinstallCustomerEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReinstallCustomerEpc(request *ReinstallCustomerEpcRequest) (string) {
    return c.ReinstallCustomerEpcWithContext(context.Background(), request)
}

func (c *Client) ReinstallCustomerEpcWithContext(ctx context.Context, request *ReinstallCustomerEpcRequest) (string) {
    if request == nil {
        request = NewReinstallCustomerEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewReinstallCustomerEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteRemoteManagementRequest() (request *DeleteRemoteManagementRequest) {
    request = &DeleteRemoteManagementRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DeleteRemoteManagement")
    return
}

func NewDeleteRemoteManagementResponse() (response *DeleteRemoteManagementResponse) {
    response = &DeleteRemoteManagementResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteRemoteManagement(request *DeleteRemoteManagementRequest) (string) {
    return c.DeleteRemoteManagementWithContext(context.Background(), request)
}

func (c *Client) DeleteRemoteManagementWithContext(ctx context.Context, request *DeleteRemoteManagementRequest) (string) {
    if request == nil {
        request = NewDeleteRemoteManagementRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteRemoteManagementResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ResetPassword")
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ResetPassword(request *ResetPasswordRequest) (string) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) (string) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewResetPasswordResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyHyperThreadingRequest() (request *ModifyHyperThreadingRequest) {
    request = &ModifyHyperThreadingRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyHyperThreading")
    return
}

func NewModifyHyperThreadingResponse() (response *ModifyHyperThreadingResponse) {
    response = &ModifyHyperThreadingResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyHyperThreading(request *ModifyHyperThreadingRequest) (string) {
    return c.ModifyHyperThreadingWithContext(context.Background(), request)
}

func (c *Client) ModifyHyperThreadingWithContext(ctx context.Context, request *ModifyHyperThreadingRequest) (string) {
    if request == nil {
        request = NewModifyHyperThreadingRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyHyperThreadingResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAssociateClusterRequest() (request *AssociateClusterRequest) {
    request = &AssociateClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "AssociateCluster")
    return
}

func NewAssociateClusterResponse() (response *AssociateClusterResponse) {
    response = &AssociateClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AssociateCluster(request *AssociateClusterRequest) (string) {
    return c.AssociateClusterWithContext(context.Background(), request)
}

func (c *Client) AssociateClusterWithContext(ctx context.Context, request *AssociateClusterRequest) (string) {
    if request == nil {
        request = NewAssociateClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAssociateClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDisassociateClusterRequest() (request *DisassociateClusterRequest) {
    request = &DisassociateClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DisassociateCluster")
    return
}

func NewDisassociateClusterResponse() (response *DisassociateClusterResponse) {
    response = &DisassociateClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DisassociateCluster(request *DisassociateClusterRequest) (string) {
    return c.DisassociateClusterWithContext(context.Background(), request)
}

func (c *Client) DisassociateClusterWithContext(ctx context.Context, request *DisassociateClusterRequest) (string) {
    if request == nil {
        request = NewDisassociateClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDisassociateClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInspectionsRequest() (request *DescribeInspectionsRequest) {
    request = &DescribeInspectionsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeInspections")
    return
}

func NewDescribeInspectionsResponse() (response *DescribeInspectionsResponse) {
    response = &DescribeInspectionsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInspections(request *DescribeInspectionsRequest) (string) {
    return c.DescribeInspectionsWithContext(context.Background(), request)
}

func (c *Client) DescribeInspectionsWithContext(ctx context.Context, request *DescribeInspectionsRequest) (string) {
    if request == nil {
        request = NewDescribeInspectionsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeInspectionsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcStocksRequest() (request *DescribeEpcStocksRequest) {
    request = &DescribeEpcStocksRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcStocks")
    return
}

func NewDescribeEpcStocksResponse() (response *DescribeEpcStocksResponse) {
    response = &DescribeEpcStocksResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcStocks(request *DescribeEpcStocksRequest) (string) {
    return c.DescribeEpcStocksWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcStocksWithContext(ctx context.Context, request *DescribeEpcStocksRequest) (string) {
    if request == nil {
        request = NewDescribeEpcStocksRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcStocksResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcDeviceAttributesRequest() (request *DescribeEpcDeviceAttributesRequest) {
    request = &DescribeEpcDeviceAttributesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcDeviceAttributes")
    return
}

func NewDescribeEpcDeviceAttributesResponse() (response *DescribeEpcDeviceAttributesResponse) {
    response = &DescribeEpcDeviceAttributesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcDeviceAttributes(request *DescribeEpcDeviceAttributesRequest) (string) {
    return c.DescribeEpcDeviceAttributesWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcDeviceAttributesWithContext(ctx context.Context, request *DescribeEpcDeviceAttributesRequest) (string) {
    if request == nil {
        request = NewDescribeEpcDeviceAttributesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcDeviceAttributesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
    request = &DescribeProcessesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeProcesses")
    return
}

func NewDescribeProcessesResponse() (response *DescribeProcessesResponse) {
    response = &DescribeProcessesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeProcesses(request *DescribeProcessesRequest) (string) {
    return c.DescribeProcessesWithContext(context.Background(), request)
}

func (c *Client) DescribeProcessesWithContext(ctx context.Context, request *DescribeProcessesRequest) (string) {
    if request == nil {
        request = NewDescribeProcessesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeProcessesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateProcessRequest() (request *CreateProcessRequest) {
    request = &CreateProcessRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateProcess")
    return
}

func NewCreateProcessResponse() (response *CreateProcessResponse) {
    response = &CreateProcessResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateProcess(request *CreateProcessRequest) (string) {
    return c.CreateProcessWithContext(context.Background(), request)
}

func (c *Client) CreateProcessWithContext(ctx context.Context, request *CreateProcessRequest) (string) {
    if request == nil {
        request = NewCreateProcessRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateProcessResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteProcessRequest() (request *DeleteProcessRequest) {
    request = &DeleteProcessRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DeleteProcess")
    return
}

func NewDeleteProcessResponse() (response *DeleteProcessResponse) {
    response = &DeleteProcessResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteProcess(request *DeleteProcessRequest) (string) {
    return c.DeleteProcessWithContext(context.Background(), request)
}

func (c *Client) DeleteProcessWithContext(ctx context.Context, request *DeleteProcessRequest) (string) {
    if request == nil {
        request = NewDeleteProcessRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteProcessResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewReplyProcessRequest() (request *ReplyProcessRequest) {
    request = &ReplyProcessRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ReplyProcess")
    return
}

func NewReplyProcessResponse() (response *ReplyProcessResponse) {
    response = &ReplyProcessResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReplyProcess(request *ReplyProcessRequest) (string) {
    return c.ReplyProcessWithContext(context.Background(), request)
}

func (c *Client) ReplyProcessWithContext(ctx context.Context, request *ReplyProcessRequest) (string) {
    if request == nil {
        request = NewReplyProcessRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewReplyProcessResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcTrashesRequest() (request *DescribeEpcTrashesRequest) {
    request = &DescribeEpcTrashesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcTrashes")
    return
}

func NewDescribeEpcTrashesResponse() (response *DescribeEpcTrashesResponse) {
    response = &DescribeEpcTrashesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcTrashes(request *DescribeEpcTrashesRequest) (string) {
    return c.DescribeEpcTrashesWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcTrashesWithContext(ctx context.Context, request *DescribeEpcTrashesRequest) (string) {
    if request == nil {
        request = NewDescribeEpcTrashesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcTrashesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewReturnEpcRequest() (request *ReturnEpcRequest) {
    request = &ReturnEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ReturnEpc")
    return
}

func NewReturnEpcResponse() (response *ReturnEpcResponse) {
    response = &ReturnEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReturnEpc(request *ReturnEpcRequest) (string) {
    return c.ReturnEpcWithContext(context.Background(), request)
}

func (c *Client) ReturnEpcWithContext(ctx context.Context, request *ReturnEpcRequest) (string) {
    if request == nil {
        request = NewReturnEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewReturnEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateResourceRequirementRequest() (request *CreateResourceRequirementRequest) {
    request = &CreateResourceRequirementRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateResourceRequirement")
    return
}

func NewCreateResourceRequirementResponse() (response *CreateResourceRequirementResponse) {
    response = &CreateResourceRequirementResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateResourceRequirement(request *CreateResourceRequirementRequest) (string) {
    return c.CreateResourceRequirementWithContext(context.Background(), request)
}

func (c *Client) CreateResourceRequirementWithContext(ctx context.Context, request *CreateResourceRequirementRequest) (string) {
    if request == nil {
        request = NewCreateResourceRequirementRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateResourceRequirementResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAttachVolumeRequest() (request *AttachVolumeRequest) {
    request = &AttachVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "AttachVolume")
    return
}

func NewAttachVolumeResponse() (response *AttachVolumeResponse) {
    response = &AttachVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AttachVolume(request *AttachVolumeRequest) (string) {
    return c.AttachVolumeWithContext(context.Background(), request)
}

func (c *Client) AttachVolumeWithContext(ctx context.Context, request *AttachVolumeRequest) (string) {
    if request == nil {
        request = NewAttachVolumeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAttachVolumeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDetachVolumeRequest() (request *DetachVolumeRequest) {
    request = &DetachVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DetachVolume")
    return
}

func NewDetachVolumeResponse() (response *DetachVolumeResponse) {
    response = &DetachVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DetachVolume(request *DetachVolumeRequest) (string) {
    return c.DetachVolumeWithContext(context.Background(), request)
}

func (c *Client) DetachVolumeWithContext(ctx context.Context, request *DetachVolumeRequest) (string) {
    if request == nil {
        request = NewDetachVolumeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDetachVolumeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribePriceRequest() (request *DescribePriceRequest) {
    request = &DescribePriceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribePrice")
    return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
    response = &DescribePriceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribePrice(request *DescribePriceRequest) (string) {
    return c.DescribePriceWithContext(context.Background(), request)
}

func (c *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest) (string) {
    if request == nil {
        request = NewDescribePriceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribePriceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyOverclockingAttributeRequest() (request *ModifyOverclockingAttributeRequest) {
    request = &ModifyOverclockingAttributeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ModifyOverclockingAttribute")
    return
}

func NewModifyOverclockingAttributeResponse() (response *ModifyOverclockingAttributeResponse) {
    response = &ModifyOverclockingAttributeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyOverclockingAttribute(request *ModifyOverclockingAttributeRequest) (string) {
    return c.ModifyOverclockingAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyOverclockingAttributeWithContext(ctx context.Context, request *ModifyOverclockingAttributeRequest) (string) {
    if request == nil {
        request = NewModifyOverclockingAttributeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyOverclockingAttributeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCopyImageRequest() (request *CopyImageRequest) {
    request = &CopyImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CopyImage")
    return
}

func NewCopyImageResponse() (response *CopyImageResponse) {
    response = &CopyImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CopyImage(request *CopyImageRequest) (string) {
    return c.CopyImageWithContext(context.Background(), request)
}

func (c *Client) CopyImageWithContext(ctx context.Context, request *CopyImageRequest) (string) {
    if request == nil {
        request = NewCopyImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCopyImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcRaidAttributesRequest() (request *DescribeEpcRaidAttributesRequest) {
    request = &DescribeEpcRaidAttributesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcRaidAttributes")
    return
}

func NewDescribeEpcRaidAttributesResponse() (response *DescribeEpcRaidAttributesResponse) {
    response = &DescribeEpcRaidAttributesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcRaidAttributes(request *DescribeEpcRaidAttributesRequest) (string) {
    return c.DescribeEpcRaidAttributesWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcRaidAttributesWithContext(ctx context.Context, request *DescribeEpcRaidAttributesRequest) (string) {
    if request == nil {
        request = NewDescribeEpcRaidAttributesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcRaidAttributesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeGpuImageDriverRequest() (request *DescribeGpuImageDriverRequest) {
    request = &DescribeGpuImageDriverRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeGpuImageDriver")
    return
}

func NewDescribeGpuImageDriverResponse() (response *DescribeGpuImageDriverResponse) {
    response = &DescribeGpuImageDriverResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeGpuImageDriver(request *DescribeGpuImageDriverRequest) (string) {
    return c.DescribeGpuImageDriverWithContext(context.Background(), request)
}

func (c *Client) DescribeGpuImageDriverWithContext(ctx context.Context, request *DescribeGpuImageDriverRequest) (string) {
    if request == nil {
        request = NewDescribeGpuImageDriverRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeGpuImageDriverResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateShareImageRequest() (request *CreateShareImageRequest) {
    request = &CreateShareImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CreateShareImage")
    return
}

func NewCreateShareImageResponse() (response *CreateShareImageResponse) {
    response = &CreateShareImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateShareImage(request *CreateShareImageRequest) (string) {
    return c.CreateShareImageWithContext(context.Background(), request)
}

func (c *Client) CreateShareImageWithContext(ctx context.Context, request *CreateShareImageRequest) (string) {
    if request == nil {
        request = NewCreateShareImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateShareImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteShareImageRequest() (request *DeleteShareImageRequest) {
    request = &DeleteShareImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DeleteShareImage")
    return
}

func NewDeleteShareImageResponse() (response *DeleteShareImageResponse) {
    response = &DeleteShareImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteShareImage(request *DeleteShareImageRequest) (string) {
    return c.DeleteShareImageWithContext(context.Background(), request)
}

func (c *Client) DeleteShareImageWithContext(ctx context.Context, request *DeleteShareImageRequest) (string) {
    if request == nil {
        request = NewDeleteShareImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteShareImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeShareImageAccountListRequest() (request *DescribeShareImageAccountListRequest) {
    request = &DescribeShareImageAccountListRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeShareImageAccountList")
    return
}

func NewDescribeShareImageAccountListResponse() (response *DescribeShareImageAccountListResponse) {
    response = &DescribeShareImageAccountListResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeShareImageAccountList(request *DescribeShareImageAccountListRequest) (string) {
    return c.DescribeShareImageAccountListWithContext(context.Background(), request)
}

func (c *Client) DescribeShareImageAccountListWithContext(ctx context.Context, request *DescribeShareImageAccountListRequest) (string) {
    if request == nil {
        request = NewDescribeShareImageAccountListRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeShareImageAccountListResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeShareImageRequest() (request *DescribeShareImageRequest) {
    request = &DescribeShareImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeShareImage")
    return
}

func NewDescribeShareImageResponse() (response *DescribeShareImageResponse) {
    response = &DescribeShareImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeShareImage(request *DescribeShareImageRequest) (string) {
    return c.DescribeShareImageWithContext(context.Background(), request)
}

func (c *Client) DescribeShareImageWithContext(ctx context.Context, request *DescribeShareImageRequest) (string) {
    if request == nil {
        request = NewDescribeShareImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeShareImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAcceptShareImageRequest() (request *AcceptShareImageRequest) {
    request = &AcceptShareImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "AcceptShareImage")
    return
}

func NewAcceptShareImageResponse() (response *AcceptShareImageResponse) {
    response = &AcceptShareImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AcceptShareImage(request *AcceptShareImageRequest) (string) {
    return c.AcceptShareImageWithContext(context.Background(), request)
}

func (c *Client) AcceptShareImageWithContext(ctx context.Context, request *AcceptShareImageRequest) (string) {
    if request == nil {
        request = NewAcceptShareImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAcceptShareImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRejectShareImageRequest() (request *RejectShareImageRequest) {
    request = &RejectShareImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "RejectShareImage")
    return
}

func NewRejectShareImageResponse() (response *RejectShareImageResponse) {
    response = &RejectShareImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RejectShareImage(request *RejectShareImageRequest) (string) {
    return c.RejectShareImageWithContext(context.Background(), request)
}

func (c *Client) RejectShareImageWithContext(ctx context.Context, request *RejectShareImageRequest) (string) {
    if request == nil {
        request = NewRejectShareImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRejectShareImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeManagedAccessoryRequest() (request *DescribeManagedAccessoryRequest) {
    request = &DescribeManagedAccessoryRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "DescribeManagedAccessory")
    return
}

func NewDescribeManagedAccessoryResponse() (response *DescribeManagedAccessoryResponse) {
    response = &DescribeManagedAccessoryResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeManagedAccessory(request *DescribeManagedAccessoryRequest) (string) {
    return c.DescribeManagedAccessoryWithContext(context.Background(), request)
}

func (c *Client) DescribeManagedAccessoryWithContext(ctx context.Context, request *DescribeManagedAccessoryRequest) (string) {
    if request == nil {
        request = NewDescribeManagedAccessoryRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeManagedAccessoryResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAutoDeleteEpcRequest() (request *AutoDeleteEpcRequest) {
    request = &AutoDeleteEpcRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "AutoDeleteEpc")
    return
}

func NewAutoDeleteEpcResponse() (response *AutoDeleteEpcResponse) {
    response = &AutoDeleteEpcResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AutoDeleteEpc(request *AutoDeleteEpcRequest) (string) {
    return c.AutoDeleteEpcWithContext(context.Background(), request)
}

func (c *Client) AutoDeleteEpcWithContext(ctx context.Context, request *AutoDeleteEpcRequest) (string) {
    if request == nil {
        request = NewAutoDeleteEpcRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAutoDeleteEpcResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewExportImageRequest() (request *ExportImageRequest) {
    request = &ExportImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "ExportImage")
    return
}

func NewExportImageResponse() (response *ExportImageResponse) {
    response = &ExportImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ExportImage(request *ExportImageRequest) (string) {
    return c.ExportImageWithContext(context.Background(), request)
}

func (c *Client) ExportImageWithContext(ctx context.Context, request *ExportImageRequest) (string) {
    if request == nil {
        request = NewExportImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewExportImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewQueryBucketsRequest() (request *QueryBucketsRequest) {
    request = &QueryBucketsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "QueryBuckets")
    return
}

func NewQueryBucketsResponse() (response *QueryBucketsResponse) {
    response = &QueryBucketsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) QueryBuckets(request *QueryBucketsRequest) (string) {
    return c.QueryBucketsWithContext(context.Background(), request)
}

func (c *Client) QueryBucketsWithContext(ctx context.Context, request *QueryBucketsRequest) (string) {
    if request == nil {
        request = NewQueryBucketsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewQueryBucketsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCancelImageExportRequest() (request *CancelImageExportRequest) {
    request = &CancelImageExportRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("epc", APIVersion, "CancelImageExport")
    return
}

func NewCancelImageExportResponse() (response *CancelImageExportResponse) {
    response = &CancelImageExportResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CancelImageExport(request *CancelImageExportRequest) (string) {
    return c.CancelImageExportWithContext(context.Background(), request)
}

func (c *Client) CancelImageExportWithContext(ctx context.Context, request *CancelImageExportRequest) (string) {
    if request == nil {
        request = NewCancelImageExportRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCancelImageExportResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


