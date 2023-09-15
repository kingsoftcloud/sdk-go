package v20160701
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2016-07-01"

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

func NewCreateCacheClusterRequest() (request *CreateCacheClusterRequest) {
    request = &CreateCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "CreateCacheCluster")
    return
}

func NewCreateCacheClusterResponse() (response *CreateCacheClusterResponse) {
    response = &CreateCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateCacheCluster(request *CreateCacheClusterRequest) (string) {
    return c.CreateCacheClusterWithContext(context.Background(), request)
}

func (c *Client) CreateCacheClusterWithContext(ctx context.Context, request *CreateCacheClusterRequest) (string) {
    if request == nil {
        request = NewCreateCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteCacheClusterRequest() (request *DeleteCacheClusterRequest) {
    request = &DeleteCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteCacheCluster")
    return
}

func NewDeleteCacheClusterResponse() (response *DeleteCacheClusterResponse) {
    response = &DeleteCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteCacheCluster(request *DeleteCacheClusterRequest) (string) {
    return c.DeleteCacheClusterWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheClusterWithContext(ctx context.Context, request *DeleteCacheClusterRequest) (string) {
    if request == nil {
        request = NewDeleteCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheClusterRequest() (request *DescribeCacheClusterRequest) {
    request = &DescribeCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheCluster")
    return
}

func NewDescribeCacheClusterResponse() (response *DescribeCacheClusterResponse) {
    response = &DescribeCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheCluster(request *DescribeCacheClusterRequest) (string) {
    return c.DescribeCacheClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheClusterWithContext(ctx context.Context, request *DescribeCacheClusterRequest) (string) {
    if request == nil {
        request = NewDescribeCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheClustersRequest() (request *DescribeCacheClustersRequest) {
    request = &DescribeCacheClustersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheClusters")
    return
}

func NewDescribeCacheClustersResponse() (response *DescribeCacheClustersResponse) {
    response = &DescribeCacheClustersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheClusters(request *DescribeCacheClustersRequest) (string) {
    return c.DescribeCacheClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheClustersWithContext(ctx context.Context, request *DescribeCacheClustersRequest) (string) {
    if request == nil {
        request = NewDescribeCacheClustersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheClustersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewFlushCacheClusterRequest() (request *FlushCacheClusterRequest) {
    request = &FlushCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "FlushCacheCluster")
    return
}

func NewFlushCacheClusterResponse() (response *FlushCacheClusterResponse) {
    response = &FlushCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) FlushCacheCluster(request *FlushCacheClusterRequest) (string) {
    return c.FlushCacheClusterWithContext(context.Background(), request)
}

func (c *Client) FlushCacheClusterWithContext(ctx context.Context, request *FlushCacheClusterRequest) (string) {
    if request == nil {
        request = NewFlushCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewFlushCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRenameCacheClusterRequest() (request *RenameCacheClusterRequest) {
    request = &RenameCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "RenameCacheCluster")
    return
}

func NewRenameCacheClusterResponse() (response *RenameCacheClusterResponse) {
    response = &RenameCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenameCacheCluster(request *RenameCacheClusterRequest) (string) {
    return c.RenameCacheClusterWithContext(context.Background(), request)
}

func (c *Client) RenameCacheClusterWithContext(ctx context.Context, request *RenameCacheClusterRequest) (string) {
    if request == nil {
        request = NewRenameCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRenameCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewResizeCacheClusterRequest() (request *ResizeCacheClusterRequest) {
    request = &ResizeCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "ResizeCacheCluster")
    return
}

func NewResizeCacheClusterResponse() (response *ResizeCacheClusterResponse) {
    response = &ResizeCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ResizeCacheCluster(request *ResizeCacheClusterRequest) (string) {
    return c.ResizeCacheClusterWithContext(context.Background(), request)
}

func (c *Client) ResizeCacheClusterWithContext(ctx context.Context, request *ResizeCacheClusterRequest) (string) {
    if request == nil {
        request = NewResizeCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewResizeCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheParametersRequest() (request *DescribeCacheParametersRequest) {
    request = &DescribeCacheParametersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheParameters")
    return
}

func NewDescribeCacheParametersResponse() (response *DescribeCacheParametersResponse) {
    response = &DescribeCacheParametersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheParameters(request *DescribeCacheParametersRequest) (string) {
    return c.DescribeCacheParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheParametersWithContext(ctx context.Context, request *DescribeCacheParametersRequest) (string) {
    if request == nil {
        request = NewDescribeCacheParametersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheParametersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetCacheParametersRequest() (request *SetCacheParametersRequest) {
    request = &SetCacheParametersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "SetCacheParameters")
    return
}

func NewSetCacheParametersResponse() (response *SetCacheParametersResponse) {
    response = &SetCacheParametersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetCacheParameters(request *SetCacheParametersRequest) (string) {
    return c.SetCacheParametersWithContext(context.Background(), request)
}

func (c *Client) SetCacheParametersWithContext(ctx context.Context, request *SetCacheParametersRequest) (string) {
    if request == nil {
        request = NewSetCacheParametersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSetCacheParametersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheDefaultParametersRequest() (request *DescribeCacheDefaultParametersRequest) {
    request = &DescribeCacheDefaultParametersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheDefaultParameters")
    return
}

func NewDescribeCacheDefaultParametersResponse() (response *DescribeCacheDefaultParametersResponse) {
    response = &DescribeCacheDefaultParametersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheDefaultParameters(request *DescribeCacheDefaultParametersRequest) (string) {
    return c.DescribeCacheDefaultParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheDefaultParametersWithContext(ctx context.Context, request *DescribeCacheDefaultParametersRequest) (string) {
    if request == nil {
        request = NewDescribeCacheDefaultParametersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheDefaultParametersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetCacheParameterGroupRequest() (request *SetCacheParameterGroupRequest) {
    request = &SetCacheParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "SetCacheParameterGroup")
    return
}

func NewSetCacheParameterGroupResponse() (response *SetCacheParameterGroupResponse) {
    response = &SetCacheParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetCacheParameterGroup(request *SetCacheParameterGroupRequest) (string) {
    return c.SetCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) SetCacheParameterGroupWithContext(ctx context.Context, request *SetCacheParameterGroupRequest) (string) {
    if request == nil {
        request = NewSetCacheParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSetCacheParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateCacheParameterGroupRequest() (request *CreateCacheParameterGroupRequest) {
    request = &CreateCacheParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "CreateCacheParameterGroup")
    return
}

func NewCreateCacheParameterGroupResponse() (response *CreateCacheParameterGroupResponse) {
    response = &CreateCacheParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateCacheParameterGroup(request *CreateCacheParameterGroupRequest) (string) {
    return c.CreateCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) CreateCacheParameterGroupWithContext(ctx context.Context, request *CreateCacheParameterGroupRequest) (string) {
    if request == nil {
        request = NewCreateCacheParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateCacheParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteCacheParameterGroupRequest() (request *DeleteCacheParameterGroupRequest) {
    request = &DeleteCacheParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteCacheParameterGroup")
    return
}

func NewDeleteCacheParameterGroupResponse() (response *DeleteCacheParameterGroupResponse) {
    response = &DeleteCacheParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteCacheParameterGroup(request *DeleteCacheParameterGroupRequest) (string) {
    return c.DeleteCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheParameterGroupWithContext(ctx context.Context, request *DeleteCacheParameterGroupRequest) (string) {
    if request == nil {
        request = NewDeleteCacheParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteCacheParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyCacheParameterGroupRequest() (request *ModifyCacheParameterGroupRequest) {
    request = &ModifyCacheParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "ModifyCacheParameterGroup")
    return
}

func NewModifyCacheParameterGroupResponse() (response *ModifyCacheParameterGroupResponse) {
    response = &ModifyCacheParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyCacheParameterGroup(request *ModifyCacheParameterGroupRequest) (string) {
    return c.ModifyCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyCacheParameterGroupWithContext(ctx context.Context, request *ModifyCacheParameterGroupRequest) (string) {
    if request == nil {
        request = NewModifyCacheParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyCacheParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheParameterGroupsRequest() (request *DescribeCacheParameterGroupsRequest) {
    request = &DescribeCacheParameterGroupsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheParameterGroups")
    return
}

func NewDescribeCacheParameterGroupsResponse() (response *DescribeCacheParameterGroupsResponse) {
    response = &DescribeCacheParameterGroupsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheParameterGroups(request *DescribeCacheParameterGroupsRequest) (string) {
    return c.DescribeCacheParameterGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheParameterGroupsWithContext(ctx context.Context, request *DescribeCacheParameterGroupsRequest) (string) {
    if request == nil {
        request = NewDescribeCacheParameterGroupsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheParameterGroupsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheParameterGroupRequest() (request *DescribeCacheParameterGroupRequest) {
    request = &DescribeCacheParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheParameterGroup")
    return
}

func NewDescribeCacheParameterGroupResponse() (response *DescribeCacheParameterGroupResponse) {
    response = &DescribeCacheParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheParameterGroup(request *DescribeCacheParameterGroupRequest) (string) {
    return c.DescribeCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheParameterGroupWithContext(ctx context.Context, request *DescribeCacheParameterGroupRequest) (string) {
    if request == nil {
        request = NewDescribeCacheParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetTimingSnapshotRequest() (request *SetTimingSnapshotRequest) {
    request = &SetTimingSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "SetTimingSnapshot")
    return
}

func NewSetTimingSnapshotResponse() (response *SetTimingSnapshotResponse) {
    response = &SetTimingSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetTimingSnapshot(request *SetTimingSnapshotRequest) (string) {
    return c.SetTimingSnapshotWithContext(context.Background(), request)
}

func (c *Client) SetTimingSnapshotWithContext(ctx context.Context, request *SetTimingSnapshotRequest) (string) {
    if request == nil {
        request = NewSetTimingSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSetTimingSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteSnapshotRequest() (request *DeleteSnapshotRequest) {
    request = &DeleteSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteSnapshot")
    return
}

func NewDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
    response = &DeleteSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (string) {
    return c.DeleteSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest) (string) {
    if request == nil {
        request = NewDeleteSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRenameSnapshotRequest() (request *RenameSnapshotRequest) {
    request = &RenameSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "RenameSnapshot")
    return
}

func NewRenameSnapshotResponse() (response *RenameSnapshotResponse) {
    response = &RenameSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenameSnapshot(request *RenameSnapshotRequest) (string) {
    return c.RenameSnapshotWithContext(context.Background(), request)
}

func (c *Client) RenameSnapshotWithContext(ctx context.Context, request *RenameSnapshotRequest) (string) {
    if request == nil {
        request = NewRenameSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRenameSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestoreSnapshotRequest() (request *RestoreSnapshotRequest) {
    request = &RestoreSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "RestoreSnapshot")
    return
}

func NewRestoreSnapshotResponse() (response *RestoreSnapshotResponse) {
    response = &RestoreSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestoreSnapshot(request *RestoreSnapshotRequest) (string) {
    return c.RestoreSnapshotWithContext(context.Background(), request)
}

func (c *Client) RestoreSnapshotWithContext(ctx context.Context, request *RestoreSnapshotRequest) (string) {
    if request == nil {
        request = NewRestoreSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRestoreSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeSnapshots")
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (string) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (string) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSnapshotsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDownloadSnapshotRequest() (request *DownloadSnapshotRequest) {
    request = &DownloadSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DownloadSnapshot")
    return
}

func NewDownloadSnapshotResponse() (response *DownloadSnapshotResponse) {
    response = &DownloadSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DownloadSnapshot(request *DownloadSnapshotRequest) (string) {
    return c.DownloadSnapshotWithContext(context.Background(), request)
}

func (c *Client) DownloadSnapshotWithContext(ctx context.Context, request *DownloadSnapshotRequest) (string) {
    if request == nil {
        request = NewDownloadSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDownloadSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewExportSnapshotRequest() (request *ExportSnapshotRequest) {
    request = &ExportSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "ExportSnapshot")
    return
}

func NewExportSnapshotResponse() (response *ExportSnapshotResponse) {
    response = &ExportSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ExportSnapshot(request *ExportSnapshotRequest) (string) {
    return c.ExportSnapshotWithContext(context.Background(), request)
}

func (c *Client) ExportSnapshotWithContext(ctx context.Context, request *ExportSnapshotRequest) (string) {
    if request == nil {
        request = NewExportSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewExportSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (string) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (string) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeRegionsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeAvailabilityZonesRequest() (request *DescribeAvailabilityZonesRequest) {
    request = &DescribeAvailabilityZonesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeAvailabilityZones")
    return
}

func NewDescribeAvailabilityZonesResponse() (response *DescribeAvailabilityZonesResponse) {
    response = &DescribeAvailabilityZonesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeAvailabilityZones(request *DescribeAvailabilityZonesRequest) (string) {
    return c.DescribeAvailabilityZonesWithContext(context.Background(), request)
}

func (c *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, request *DescribeAvailabilityZonesRequest) (string) {
    if request == nil {
        request = NewDescribeAvailabilityZonesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeAvailabilityZonesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCacheByRoleRequest() (request *DescribeCacheByRoleRequest) {
    request = &DescribeCacheByRoleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheByRole")
    return
}

func NewDescribeCacheByRoleResponse() (response *DescribeCacheByRoleResponse) {
    response = &DescribeCacheByRoleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCacheByRole(request *DescribeCacheByRoleRequest) (string) {
    return c.DescribeCacheByRoleWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheByRoleWithContext(ctx context.Context, request *DescribeCacheByRoleRequest) (string) {
    if request == nil {
        request = NewDescribeCacheByRoleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCacheByRoleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStatisticDBInstancesRequest() (request *StatisticDBInstancesRequest) {
    request = &StatisticDBInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "StatisticDBInstances")
    return
}

func NewStatisticDBInstancesResponse() (response *StatisticDBInstancesResponse) {
    response = &StatisticDBInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StatisticDBInstances(request *StatisticDBInstancesRequest) (string) {
    return c.StatisticDBInstancesWithContext(context.Background(), request)
}

func (c *Client) StatisticDBInstancesWithContext(ctx context.Context, request *StatisticDBInstancesRequest) (string) {
    if request == nil {
        request = NewStatisticDBInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStatisticDBInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdatePasswordRequest() (request *UpdatePasswordRequest) {
    request = &UpdatePasswordRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "UpdatePassword")
    return
}

func NewUpdatePasswordResponse() (response *UpdatePasswordResponse) {
    response = &UpdatePasswordResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdatePassword(request *UpdatePasswordRequest) (string) {
    return c.UpdatePasswordWithContext(context.Background(), request)
}

func (c *Client) UpdatePasswordWithContext(ctx context.Context, request *UpdatePasswordRequest) (string) {
    if request == nil {
        request = NewUpdatePasswordRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewUpdatePasswordResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestartCacheClusterRequest() (request *RestartCacheClusterRequest) {
    request = &RestartCacheClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "RestartCacheCluster")
    return
}

func NewRestartCacheClusterResponse() (response *RestartCacheClusterResponse) {
    response = &RestartCacheClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestartCacheCluster(request *RestartCacheClusterRequest) (string) {
    return c.RestartCacheClusterWithContext(context.Background(), request)
}

func (c *Client) RestartCacheClusterWithContext(ctx context.Context, request *RestartCacheClusterRequest) (string) {
    if request == nil {
        request = NewRestartCacheClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRestartCacheClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAllocateEipRequest() (request *AllocateEipRequest) {
    request = &AllocateEipRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AllocateEip")
    return
}

func NewAllocateEipResponse() (response *AllocateEipResponse) {
    response = &AllocateEipResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AllocateEip(request *AllocateEipRequest) (string) {
    return c.AllocateEipWithContext(context.Background(), request)
}

func (c *Client) AllocateEipWithContext(ctx context.Context, request *AllocateEipRequest) (string) {
    if request == nil {
        request = NewAllocateEipRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAllocateEipResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeallocateEipRequest() (request *DeallocateEipRequest) {
    request = &DeallocateEipRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeallocateEip")
    return
}

func NewDeallocateEipResponse() (response *DeallocateEipResponse) {
    response = &DeallocateEipResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeallocateEip(request *DeallocateEipRequest) (string) {
    return c.DeallocateEipWithContext(context.Background(), request)
}

func (c *Client) DeallocateEipWithContext(ctx context.Context, request *DeallocateEipRequest) (string) {
    if request == nil {
        request = NewDeallocateEipRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeallocateEipResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (string) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (string) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
    request = &DeleteSecurityGroupRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteSecurityGroupRule")
    return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
    response = &DeleteSecurityGroupRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (string) {
    return c.DeleteSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupRuleWithContext(ctx context.Context, request *DeleteSecurityGroupRuleRequest) (string) {
    if request == nil {
        request = NewDeleteSecurityGroupRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteSecurityGroupRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateSecurityGroupRuleRequest() (request *CreateSecurityGroupRuleRequest) {
    request = &CreateSecurityGroupRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "CreateSecurityGroupRule")
    return
}

func NewCreateSecurityGroupRuleResponse() (response *CreateSecurityGroupRuleResponse) {
    response = &CreateSecurityGroupRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateSecurityGroupRule(request *CreateSecurityGroupRuleRequest) (string) {
    return c.CreateSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupRuleWithContext(ctx context.Context, request *CreateSecurityGroupRuleRequest) (string) {
    if request == nil {
        request = NewCreateSecurityGroupRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateSecurityGroupRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeallocateSecurityGroupRequest() (request *DeallocateSecurityGroupRequest) {
    request = &DeallocateSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeallocateSecurityGroup")
    return
}

func NewDeallocateSecurityGroupResponse() (response *DeallocateSecurityGroupResponse) {
    response = &DeallocateSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeallocateSecurityGroup(request *DeallocateSecurityGroupRequest) (string) {
    return c.DeallocateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeallocateSecurityGroupWithContext(ctx context.Context, request *DeallocateSecurityGroupRequest) (string) {
    if request == nil {
        request = NewDeallocateSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeallocateSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAllocateSecurityGroupRequest() (request *AllocateSecurityGroupRequest) {
    request = &AllocateSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AllocateSecurityGroup")
    return
}

func NewAllocateSecurityGroupResponse() (response *AllocateSecurityGroupResponse) {
    response = &AllocateSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AllocateSecurityGroup(request *AllocateSecurityGroupRequest) (string) {
    return c.AllocateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) AllocateSecurityGroupWithContext(ctx context.Context, request *AllocateSecurityGroupRequest) (string) {
    if request == nil {
        request = NewAllocateSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAllocateSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
    request = &DescribeSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeSecurityGroup")
    return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
    response = &DescribeSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (string) {
    return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) (string) {
    if request == nil {
        request = NewDescribeSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
    request = &DescribeSecurityGroupsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeSecurityGroups")
    return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
    response = &DescribeSecurityGroupsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (string) {
    return c.DescribeSecurityGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest) (string) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSecurityGroupsResponse()
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
    request.Init().WithApiInfo("kcs", APIVersion, "ModifySecurityGroup")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifySecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
    request = &DeleteSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteSecurityGroup")
    return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
    response = &DeleteSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (string) {
    return c.DeleteSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) (string) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
    request = &CloneSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "CloneSecurityGroup")
    return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
    response = &CloneSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) (string) {
    return c.CloneSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) (string) {
    if request == nil {
        request = NewCloneSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCloneSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
    request = &CreateSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "CreateSecurityGroup")
    return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
    response = &CreateSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (string) {
    return c.CreateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) (string) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeHotKeysRequest() (request *DescribeHotKeysRequest) {
    request = &DescribeHotKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeHotKeys")
    return
}

func NewDescribeHotKeysResponse() (response *DescribeHotKeysResponse) {
    response = &DescribeHotKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeHotKeys(request *DescribeHotKeysRequest) (string) {
    return c.DescribeHotKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeHotKeysWithContext(ctx context.Context, request *DescribeHotKeysRequest) (string) {
    if request == nil {
        request = NewDescribeHotKeysRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeHotKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAnalyzeHotKeysRequest() (request *AnalyzeHotKeysRequest) {
    request = &AnalyzeHotKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeHotKeys")
    return
}

func NewAnalyzeHotKeysResponse() (response *AnalyzeHotKeysResponse) {
    response = &AnalyzeHotKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AnalyzeHotKeys(request *AnalyzeHotKeysRequest) (string) {
    return c.AnalyzeHotKeysWithContext(context.Background(), request)
}

func (c *Client) AnalyzeHotKeysWithContext(ctx context.Context, request *AnalyzeHotKeysRequest) (string) {
    if request == nil {
        request = NewAnalyzeHotKeysRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewAnalyzeHotKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCloseDirectAccessToClusterRequest() (request *CloseDirectAccessToClusterRequest) {
    request = &CloseDirectAccessToClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "CloseDirectAccessToCluster")
    return
}

func NewCloseDirectAccessToClusterResponse() (response *CloseDirectAccessToClusterResponse) {
    response = &CloseDirectAccessToClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CloseDirectAccessToCluster(request *CloseDirectAccessToClusterRequest) (string) {
    return c.CloseDirectAccessToClusterWithContext(context.Background(), request)
}

func (c *Client) CloseDirectAccessToClusterWithContext(ctx context.Context, request *CloseDirectAccessToClusterRequest) (string) {
    if request == nil {
        request = NewCloseDirectAccessToClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCloseDirectAccessToClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewOpenDirectAccessToClusterRequest() (request *OpenDirectAccessToClusterRequest) {
    request = &OpenDirectAccessToClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "OpenDirectAccessToCluster")
    return
}

func NewOpenDirectAccessToClusterResponse() (response *OpenDirectAccessToClusterResponse) {
    response = &OpenDirectAccessToClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) OpenDirectAccessToCluster(request *OpenDirectAccessToClusterRequest) (string) {
    return c.OpenDirectAccessToClusterWithContext(context.Background(), request)
}

func (c *Client) OpenDirectAccessToClusterWithContext(ctx context.Context, request *OpenDirectAccessToClusterRequest) (string) {
    if request == nil {
        request = NewOpenDirectAccessToClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewOpenDirectAccessToClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeParentBackUpsSnapshotsRequest() (request *DescribeParentBackUpsSnapshotsRequest) {
    request = &DescribeParentBackUpsSnapshotsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeParentBackUpsSnapshots")
    return
}

func NewDescribeParentBackUpsSnapshotsResponse() (response *DescribeParentBackUpsSnapshotsResponse) {
    response = &DescribeParentBackUpsSnapshotsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeParentBackUpsSnapshots(request *DescribeParentBackUpsSnapshotsRequest) (string) {
    return c.DescribeParentBackUpsSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeParentBackUpsSnapshotsWithContext(ctx context.Context, request *DescribeParentBackUpsSnapshotsRequest) (string) {
    if request == nil {
        request = NewDescribeParentBackUpsSnapshotsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeParentBackUpsSnapshotsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeBackUpsSnapshotsDetailRequest() (request *DescribeBackUpsSnapshotsDetailRequest) {
    request = &DescribeBackUpsSnapshotsDetailRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeBackUpsSnapshotsDetail")
    return
}

func NewDescribeBackUpsSnapshotsDetailResponse() (response *DescribeBackUpsSnapshotsDetailResponse) {
    response = &DescribeBackUpsSnapshotsDetailResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBackUpsSnapshotsDetail(request *DescribeBackUpsSnapshotsDetailRequest) (string) {
    return c.DescribeBackUpsSnapshotsDetailWithContext(context.Background(), request)
}

func (c *Client) DescribeBackUpsSnapshotsDetailWithContext(ctx context.Context, request *DescribeBackUpsSnapshotsDetailRequest) (string) {
    if request == nil {
        request = NewDescribeBackUpsSnapshotsDetailRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeBackUpsSnapshotsDetailResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteLevelSnapshotsRequest() (request *DeleteLevelSnapshotsRequest) {
    request = &DeleteLevelSnapshotsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteLevelSnapshots")
    return
}

func NewDeleteLevelSnapshotsResponse() (response *DeleteLevelSnapshotsResponse) {
    response = &DeleteLevelSnapshotsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteLevelSnapshots(request *DeleteLevelSnapshotsRequest) (string) {
    return c.DeleteLevelSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DeleteLevelSnapshotsWithContext(ctx context.Context, request *DeleteLevelSnapshotsRequest) (string) {
    if request == nil {
        request = NewDeleteLevelSnapshotsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteLevelSnapshotsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDownloadLevelSnapshotRequest() (request *DownloadLevelSnapshotRequest) {
    request = &DownloadLevelSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DownloadLevelSnapshot")
    return
}

func NewDownloadLevelSnapshotResponse() (response *DownloadLevelSnapshotResponse) {
    response = &DownloadLevelSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DownloadLevelSnapshot(request *DownloadLevelSnapshotRequest) (string) {
    return c.DownloadLevelSnapshotWithContext(context.Background(), request)
}

func (c *Client) DownloadLevelSnapshotWithContext(ctx context.Context, request *DownloadLevelSnapshotRequest) (string) {
    if request == nil {
        request = NewDownloadLevelSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDownloadLevelSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeBigKeysRequest() (request *DescribeBigKeysRequest) {
    request = &DescribeBigKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeBigKeys")
    return
}

func NewDescribeBigKeysResponse() (response *DescribeBigKeysResponse) {
    response = &DescribeBigKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBigKeys(request *DescribeBigKeysRequest) (string) {
    return c.DescribeBigKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeBigKeysWithContext(ctx context.Context, request *DescribeBigKeysRequest) (string) {
    if request == nil {
        request = NewDescribeBigKeysRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeBigKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteBigKeysAnalyseResultRequest() (request *DeleteBigKeysAnalyseResultRequest) {
    request = &DeleteBigKeysAnalyseResultRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DeleteBigKeysAnalyseResult")
    return
}

func NewDeleteBigKeysAnalyseResultResponse() (response *DeleteBigKeysAnalyseResultResponse) {
    response = &DeleteBigKeysAnalyseResultResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteBigKeysAnalyseResult(request *DeleteBigKeysAnalyseResultRequest) (string) {
    return c.DeleteBigKeysAnalyseResultWithContext(context.Background(), request)
}

func (c *Client) DeleteBigKeysAnalyseResultWithContext(ctx context.Context, request *DeleteBigKeysAnalyseResultRequest) (string) {
    if request == nil {
        request = NewDeleteBigKeysAnalyseResultRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteBigKeysAnalyseResultResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAnalyzeBigKeysRequest() (request *AnalyzeBigKeysRequest) {
    request = &AnalyzeBigKeysRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeBigKeys")
    return
}

func NewAnalyzeBigKeysResponse() (response *AnalyzeBigKeysResponse) {
    response = &AnalyzeBigKeysResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AnalyzeBigKeys(request *AnalyzeBigKeysRequest) (string) {
    return c.AnalyzeBigKeysWithContext(context.Background(), request)
}

func (c *Client) AnalyzeBigKeysWithContext(ctx context.Context, request *AnalyzeBigKeysRequest) (string) {
    if request == nil {
        request = NewAnalyzeBigKeysRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewAnalyzeBigKeysResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCreateSnapshotStatusRequest() (request *DescribeCreateSnapshotStatusRequest) {
    request = &DescribeCreateSnapshotStatusRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "DescribeCreateSnapshotStatus")
    return
}

func NewDescribeCreateSnapshotStatusResponse() (response *DescribeCreateSnapshotStatusResponse) {
    response = &DescribeCreateSnapshotStatusResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCreateSnapshotStatus(request *DescribeCreateSnapshotStatusRequest) (string) {
    return c.DescribeCreateSnapshotStatusWithContext(context.Background(), request)
}

func (c *Client) DescribeCreateSnapshotStatusWithContext(ctx context.Context, request *DescribeCreateSnapshotStatusRequest) (string) {
    if request == nil {
        request = NewDescribeCreateSnapshotStatusRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCreateSnapshotStatusResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetDailyAnalyzeSwitchStateRequest() (request *GetDailyAnalyzeSwitchStateRequest) {
    request = &GetDailyAnalyzeSwitchStateRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "GetDailyAnalyzeSwitchState")
    return
}

func NewGetDailyAnalyzeSwitchStateResponse() (response *GetDailyAnalyzeSwitchStateResponse) {
    response = &GetDailyAnalyzeSwitchStateResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetDailyAnalyzeSwitchState(request *GetDailyAnalyzeSwitchStateRequest) (string) {
    return c.GetDailyAnalyzeSwitchStateWithContext(context.Background(), request)
}

func (c *Client) GetDailyAnalyzeSwitchStateWithContext(ctx context.Context, request *GetDailyAnalyzeSwitchStateRequest) (string) {
    if request == nil {
        request = NewGetDailyAnalyzeSwitchStateRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewGetDailyAnalyzeSwitchStateResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAnalyzeDailyRequest() (request *AnalyzeDailyRequest) {
    request = &AnalyzeDailyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeDaily")
    return
}

func NewAnalyzeDailyResponse() (response *AnalyzeDailyResponse) {
    response = &AnalyzeDailyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AnalyzeDaily(request *AnalyzeDailyRequest) (string) {
    return c.AnalyzeDailyWithContext(context.Background(), request)
}

func (c *Client) AnalyzeDailyWithContext(ctx context.Context, request *AnalyzeDailyRequest) (string) {
    if request == nil {
        request = NewAnalyzeDailyRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAnalyzeDailyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAnalyzeSlowDailyRequest() (request *AnalyzeSlowDailyRequest) {
    request = &AnalyzeSlowDailyRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeSlowDaily")
    return
}

func NewAnalyzeSlowDailyResponse() (response *AnalyzeSlowDailyResponse) {
    response = &AnalyzeSlowDailyResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AnalyzeSlowDaily(request *AnalyzeSlowDailyRequest) (string) {
    return c.AnalyzeSlowDailyWithContext(context.Background(), request)
}

func (c *Client) AnalyzeSlowDailyWithContext(ctx context.Context, request *AnalyzeSlowDailyRequest) (string) {
    if request == nil {
        request = NewAnalyzeSlowDailyRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAnalyzeSlowDailyResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAnalyzeDailySwitchRequest() (request *AnalyzeDailySwitchRequest) {
    request = &AnalyzeDailySwitchRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeDailySwitch")
    return
}

func NewAnalyzeDailySwitchResponse() (response *AnalyzeDailySwitchResponse) {
    response = &AnalyzeDailySwitchResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AnalyzeDailySwitch(request *AnalyzeDailySwitchRequest) (string) {
    return c.AnalyzeDailySwitchWithContext(context.Background(), request)
}

func (c *Client) AnalyzeDailySwitchWithContext(ctx context.Context, request *AnalyzeDailySwitchRequest) (string) {
    if request == nil {
        request = NewAnalyzeDailySwitchRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAnalyzeDailySwitchResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


