package v20190806
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2019-08-06"

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

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
    request = &DescribeClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeCluster")
    return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
    response = &DescribeClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCluster(request *DescribeClusterRequest) (string) {
    return c.DescribeClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterWithContext(ctx context.Context, request *DescribeClusterRequest) (string) {
    if request == nil {
        request = NewDescribeClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeClusterInstanceRequest() (request *DescribeClusterInstanceRequest) {
    request = &DescribeClusterInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeClusterInstance")
    return
}

func NewDescribeClusterInstanceResponse() (response *DescribeClusterInstanceResponse) {
    response = &DescribeClusterInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeClusterInstance(request *DescribeClusterInstanceRequest) (string) {
    return c.DescribeClusterInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterInstanceWithContext(ctx context.Context, request *DescribeClusterInstanceRequest) (string) {
    if request == nil {
        request = NewDescribeClusterInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeClusterInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DeleteCluster")
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteCluster(request *DeleteClusterRequest) (string) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (string) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDownloadClusterConfigRequest() (request *DownloadClusterConfigRequest) {
    request = &DownloadClusterConfigRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DownloadClusterConfig")
    return
}

func NewDownloadClusterConfigResponse() (response *DownloadClusterConfigResponse) {
    response = &DownloadClusterConfigResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DownloadClusterConfig(request *DownloadClusterConfigRequest) (string) {
    return c.DownloadClusterConfigWithContext(context.Background(), request)
}

func (c *Client) DownloadClusterConfigWithContext(ctx context.Context, request *DownloadClusterConfigRequest) (string) {
    if request == nil {
        request = NewDownloadClusterConfigRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDownloadClusterConfigResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyClusterInfoRequest() (request *ModifyClusterInfoRequest) {
    request = &ModifyClusterInfoRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ModifyClusterInfo")
    return
}

func NewModifyClusterInfoResponse() (response *ModifyClusterInfoResponse) {
    response = &ModifyClusterInfoResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyClusterInfo(request *ModifyClusterInfoRequest) (string) {
    return c.ModifyClusterInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyClusterInfoWithContext(ctx context.Context, request *ModifyClusterInfoRequest) (string) {
    if request == nil {
        request = NewModifyClusterInfoRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyClusterInfoResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstanceImageRequest() (request *DescribeInstanceImageRequest) {
    request = &DescribeInstanceImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeInstanceImage")
    return
}

func NewDescribeInstanceImageResponse() (response *DescribeInstanceImageResponse) {
    response = &DescribeInstanceImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceImage(request *DescribeInstanceImageRequest) (string) {
    return c.DescribeInstanceImageWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceImageWithContext(ctx context.Context, request *DescribeInstanceImageRequest) (string) {
    if request == nil {
        request = NewDescribeInstanceImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeInstanceImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
    request = &AddClusterInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "AddClusterInstances")
    return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
    response = &AddClusterInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (string) {
    return c.AddClusterInstancesWithContext(context.Background(), request)
}

func (c *Client) AddClusterInstancesWithContext(ctx context.Context, request *AddClusterInstancesRequest) (string) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAddClusterInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
    request = &DeleteClusterInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DeleteClusterInstances")
    return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
    response = &DeleteClusterInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) (string) {
    return c.DeleteClusterInstancesWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterInstancesWithContext(ctx context.Context, request *DeleteClusterInstancesRequest) (string) {
    if request == nil {
        request = NewDeleteClusterInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteClusterInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcForClusterRequest() (request *DescribeEpcForClusterRequest) {
    request = &DescribeEpcForClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeEpcForCluster")
    return
}

func NewDescribeEpcForClusterResponse() (response *DescribeEpcForClusterResponse) {
    response = &DescribeEpcForClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcForCluster(request *DescribeEpcForClusterRequest) (string) {
    return c.DescribeEpcForClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcForClusterWithContext(ctx context.Context, request *DescribeEpcForClusterRequest) (string) {
    if request == nil {
        request = NewDescribeEpcForClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeEpcForClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddClusterEpcInstancesRequest() (request *AddClusterEpcInstancesRequest) {
    request = &AddClusterEpcInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "AddClusterEpcInstances")
    return
}

func NewAddClusterEpcInstancesResponse() (response *AddClusterEpcInstancesResponse) {
    response = &AddClusterEpcInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddClusterEpcInstances(request *AddClusterEpcInstancesRequest) (string) {
    return c.AddClusterEpcInstancesWithContext(context.Background(), request)
}

func (c *Client) AddClusterEpcInstancesWithContext(ctx context.Context, request *AddClusterEpcInstancesRequest) (string) {
    if request == nil {
        request = NewAddClusterEpcInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAddClusterEpcInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeExistedInstancesRequest() (request *DescribeExistedInstancesRequest) {
    request = &DescribeExistedInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeExistedInstances")
    return
}

func NewDescribeExistedInstancesResponse() (response *DescribeExistedInstancesResponse) {
    response = &DescribeExistedInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeExistedInstances(request *DescribeExistedInstancesRequest) (string) {
    return c.DescribeExistedInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeExistedInstancesWithContext(ctx context.Context, request *DescribeExistedInstancesRequest) (string) {
    if request == nil {
        request = NewDescribeExistedInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeExistedInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddExistedInstancesRequest() (request *AddExistedInstancesRequest) {
    request = &AddExistedInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "AddExistedInstances")
    return
}

func NewAddExistedInstancesResponse() (response *AddExistedInstancesResponse) {
    response = &AddExistedInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddExistedInstances(request *AddExistedInstancesRequest) (string) {
    return c.AddExistedInstancesWithContext(context.Background(), request)
}

func (c *Client) AddExistedInstancesWithContext(ctx context.Context, request *AddExistedInstancesRequest) (string) {
    if request == nil {
        request = NewAddExistedInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAddExistedInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewForceRemoveClusterInstanceRequest() (request *ForceRemoveClusterInstanceRequest) {
    request = &ForceRemoveClusterInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ForceRemoveClusterInstance")
    return
}

func NewForceRemoveClusterInstanceResponse() (response *ForceRemoveClusterInstanceResponse) {
    response = &ForceRemoveClusterInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ForceRemoveClusterInstance(request *ForceRemoveClusterInstanceRequest) (string) {
    return c.ForceRemoveClusterInstanceWithContext(context.Background(), request)
}

func (c *Client) ForceRemoveClusterInstanceWithContext(ctx context.Context, request *ForceRemoveClusterInstanceRequest) (string) {
    if request == nil {
        request = NewForceRemoveClusterInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewForceRemoveClusterInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateNodePoolRequest() (request *CreateNodePoolRequest) {
    request = &CreateNodePoolRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "CreateNodePool")
    return
}

func NewCreateNodePoolResponse() (response *CreateNodePoolResponse) {
    response = &CreateNodePoolResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateNodePool(request *CreateNodePoolRequest) (string) {
    return c.CreateNodePoolWithContext(context.Background(), request)
}

func (c *Client) CreateNodePoolWithContext(ctx context.Context, request *CreateNodePoolRequest) (string) {
    if request == nil {
        request = NewCreateNodePoolRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateNodePoolResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeNodePoolRequest() (request *DescribeNodePoolRequest) {
    request = &DescribeNodePoolRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeNodePool")
    return
}

func NewDescribeNodePoolResponse() (response *DescribeNodePoolResponse) {
    response = &DescribeNodePoolResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeNodePool(request *DescribeNodePoolRequest) (string) {
    return c.DescribeNodePoolWithContext(context.Background(), request)
}

func (c *Client) DescribeNodePoolWithContext(ctx context.Context, request *DescribeNodePoolRequest) (string) {
    if request == nil {
        request = NewDescribeNodePoolRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeNodePoolResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteNodePoolRequest() (request *DeleteNodePoolRequest) {
    request = &DeleteNodePoolRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DeleteNodePool")
    return
}

func NewDeleteNodePoolResponse() (response *DeleteNodePoolResponse) {
    response = &DeleteNodePoolResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteNodePool(request *DeleteNodePoolRequest) (string) {
    return c.DeleteNodePoolWithContext(context.Background(), request)
}

func (c *Client) DeleteNodePoolWithContext(ctx context.Context, request *DeleteNodePoolRequest) (string) {
    if request == nil {
        request = NewDeleteNodePoolRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteNodePoolResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyNodePoolRequest() (request *ModifyNodePoolRequest) {
    request = &ModifyNodePoolRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ModifyNodePool")
    return
}

func NewModifyNodePoolResponse() (response *ModifyNodePoolResponse) {
    response = &ModifyNodePoolResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyNodePool(request *ModifyNodePoolRequest) (string) {
    return c.ModifyNodePoolWithContext(context.Background(), request)
}

func (c *Client) ModifyNodePoolWithContext(ctx context.Context, request *ModifyNodePoolRequest) (string) {
    if request == nil {
        request = NewModifyNodePoolRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyNodePoolResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyNodeTemplateRequest() (request *ModifyNodeTemplateRequest) {
    request = &ModifyNodeTemplateRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ModifyNodeTemplate")
    return
}

func NewModifyNodeTemplateResponse() (response *ModifyNodeTemplateResponse) {
    response = &ModifyNodeTemplateResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyNodeTemplate(request *ModifyNodeTemplateRequest) (string) {
    return c.ModifyNodeTemplateWithContext(context.Background(), request)
}

func (c *Client) ModifyNodeTemplateWithContext(ctx context.Context, request *ModifyNodeTemplateRequest) (string) {
    if request == nil {
        request = NewModifyNodeTemplateRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyNodeTemplateResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyNodePoolScaleUpPolicyRequest() (request *ModifyNodePoolScaleUpPolicyRequest) {
    request = &ModifyNodePoolScaleUpPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ModifyNodePoolScaleUpPolicy")
    return
}

func NewModifyNodePoolScaleUpPolicyResponse() (response *ModifyNodePoolScaleUpPolicyResponse) {
    response = &ModifyNodePoolScaleUpPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyNodePoolScaleUpPolicy(request *ModifyNodePoolScaleUpPolicyRequest) (string) {
    return c.ModifyNodePoolScaleUpPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyNodePoolScaleUpPolicyWithContext(ctx context.Context, request *ModifyNodePoolScaleUpPolicyRequest) (string) {
    if request == nil {
        request = NewModifyNodePoolScaleUpPolicyRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyNodePoolScaleUpPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyNodePoolScaleDownPolicyRequest() (request *ModifyNodePoolScaleDownPolicyRequest) {
    request = &ModifyNodePoolScaleDownPolicyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ModifyNodePoolScaleDownPolicy")
    return
}

func NewModifyNodePoolScaleDownPolicyResponse() (response *ModifyNodePoolScaleDownPolicyResponse) {
    response = &ModifyNodePoolScaleDownPolicyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyNodePoolScaleDownPolicy(request *ModifyNodePoolScaleDownPolicyRequest) (string) {
    return c.ModifyNodePoolScaleDownPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyNodePoolScaleDownPolicyWithContext(ctx context.Context, request *ModifyNodePoolScaleDownPolicyRequest) (string) {
    if request == nil {
        request = NewModifyNodePoolScaleDownPolicyRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyNodePoolScaleDownPolicyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddClusterInstanceToNodePoolRequest() (request *AddClusterInstanceToNodePoolRequest) {
    request = &AddClusterInstanceToNodePoolRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "AddClusterInstanceToNodePool")
    return
}

func NewAddClusterInstanceToNodePoolResponse() (response *AddClusterInstanceToNodePoolResponse) {
    response = &AddClusterInstanceToNodePoolResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddClusterInstanceToNodePool(request *AddClusterInstanceToNodePoolRequest) (string) {
    return c.AddClusterInstanceToNodePoolWithContext(context.Background(), request)
}

func (c *Client) AddClusterInstanceToNodePoolWithContext(ctx context.Context, request *AddClusterInstanceToNodePoolRequest) (string) {
    if request == nil {
        request = NewAddClusterInstanceToNodePoolRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAddClusterInstanceToNodePoolResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewProtectedFromScaleDownRequest() (request *ProtectedFromScaleDownRequest) {
    request = &ProtectedFromScaleDownRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "ProtectedFromScaleDown")
    return
}

func NewProtectedFromScaleDownResponse() (response *ProtectedFromScaleDownResponse) {
    response = &ProtectedFromScaleDownResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ProtectedFromScaleDown(request *ProtectedFromScaleDownRequest) (string) {
    return c.ProtectedFromScaleDownWithContext(context.Background(), request)
}

func (c *Client) ProtectedFromScaleDownWithContext(ctx context.Context, request *ProtectedFromScaleDownRequest) (string) {
    if request == nil {
        request = NewProtectedFromScaleDownRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewProtectedFromScaleDownResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteClusterInstancesFromNodePoolRequest() (request *DeleteClusterInstancesFromNodePoolRequest) {
    request = &DeleteClusterInstancesFromNodePoolRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DeleteClusterInstancesFromNodePool")
    return
}

func NewDeleteClusterInstancesFromNodePoolResponse() (response *DeleteClusterInstancesFromNodePoolResponse) {
    response = &DeleteClusterInstancesFromNodePoolResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteClusterInstancesFromNodePool(request *DeleteClusterInstancesFromNodePoolRequest) (string) {
    return c.DeleteClusterInstancesFromNodePoolWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterInstancesFromNodePoolWithContext(ctx context.Context, request *DeleteClusterInstancesFromNodePoolRequest) (string) {
    if request == nil {
        request = NewDeleteClusterInstancesFromNodePoolRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteClusterInstancesFromNodePoolResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEpcImageRequest() (request *DescribeEpcImageRequest) {
    request = &DescribeEpcImageRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kce", APIVersion, "DescribeEpcImage")
    return
}

func NewDescribeEpcImageResponse() (response *DescribeEpcImageResponse) {
    response = &DescribeEpcImageResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEpcImage(request *DescribeEpcImageRequest) (string) {
    return c.DescribeEpcImageWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcImageWithContext(ctx context.Context, request *DescribeEpcImageRequest) (string) {
    if request == nil {
        request = NewDescribeEpcImageRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeEpcImageResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


