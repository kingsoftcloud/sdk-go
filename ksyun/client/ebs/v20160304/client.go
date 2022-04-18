package v20160304
import (
    "context"
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

func NewCreateVolumeRequest() (request *CreateVolumeRequest) {
    request = &CreateVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "CreateVolume")
    return
}

func NewCreateVolumeResponse() (response *CreateVolumeResponse) {
    response = &CreateVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateVolume(request *CreateVolumeRequest) (response *CreateVolumeResponse, err error) {
    return c.CreateVolumeWithContext(context.Background(), request)
}

func (c *Client) CreateVolumeWithContext(ctx context.Context, request *CreateVolumeRequest) (response *CreateVolumeResponse, err error) {
    if request == nil {
        request = NewCreateVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewCreateVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewAttachVolumeRequest() (request *AttachVolumeRequest) {
    request = &AttachVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "AttachVolume")
    return
}

func NewAttachVolumeResponse() (response *AttachVolumeResponse) {
    response = &AttachVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AttachVolume(request *AttachVolumeRequest) (response *AttachVolumeResponse, err error) {
    return c.AttachVolumeWithContext(context.Background(), request)
}

func (c *Client) AttachVolumeWithContext(ctx context.Context, request *AttachVolumeRequest) (response *AttachVolumeResponse, err error) {
    if request == nil {
        request = NewAttachVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewAttachVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewDetachVolumeRequest() (request *DetachVolumeRequest) {
    request = &DetachVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DetachVolume")
    return
}

func NewDetachVolumeResponse() (response *DetachVolumeResponse) {
    response = &DetachVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DetachVolume(request *DetachVolumeRequest) (response *DetachVolumeResponse, err error) {
    return c.DetachVolumeWithContext(context.Background(), request)
}

func (c *Client) DetachVolumeWithContext(ctx context.Context, request *DetachVolumeRequest) (response *DetachVolumeResponse, err error) {
    if request == nil {
        request = NewDetachVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewDetachVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewDeleteVolumeRequest() (request *DeleteVolumeRequest) {
    request = &DeleteVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DeleteVolume")
    return
}

func NewDeleteVolumeResponse() (response *DeleteVolumeResponse) {
    response = &DeleteVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteVolume(request *DeleteVolumeRequest) (response *DeleteVolumeResponse, err error) {
    return c.DeleteVolumeWithContext(context.Background(), request)
}

func (c *Client) DeleteVolumeWithContext(ctx context.Context, request *DeleteVolumeRequest) (response *DeleteVolumeResponse, err error) {
    if request == nil {
        request = NewDeleteVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewResizeVolumeRequest() (request *ResizeVolumeRequest) {
    request = &ResizeVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "ResizeVolume")
    return
}

func NewResizeVolumeResponse() (response *ResizeVolumeResponse) {
    response = &ResizeVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ResizeVolume(request *ResizeVolumeRequest) (response *ResizeVolumeResponse, err error) {
    return c.ResizeVolumeWithContext(context.Background(), request)
}

func (c *Client) ResizeVolumeWithContext(ctx context.Context, request *ResizeVolumeRequest) (response *ResizeVolumeResponse, err error) {
    if request == nil {
        request = NewResizeVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewResizeVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeVolumesRequest() (request *DescribeVolumesRequest) {
    request = &DescribeVolumesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DescribeVolumes")
    return
}

func NewDescribeVolumesResponse() (response *DescribeVolumesResponse) {
    response = &DescribeVolumesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeVolumes(request *DescribeVolumesRequest) (response *DescribeVolumesResponse, err error) {
    return c.DescribeVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeVolumesWithContext(ctx context.Context, request *DescribeVolumesRequest) (response *DescribeVolumesResponse, err error) {
    if request == nil {
        request = NewDescribeVolumesRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeVolumesResponse()
    err = c.Send(request, response)
    return
}
func NewModifyVolumeRequest() (request *ModifyVolumeRequest) {
    request = &ModifyVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "ModifyVolume")
    return
}

func NewModifyVolumeResponse() (response *ModifyVolumeResponse) {
    response = &ModifyVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyVolume(request *ModifyVolumeRequest) (response *ModifyVolumeResponse, err error) {
    return c.ModifyVolumeWithContext(context.Background(), request)
}

func (c *Client) ModifyVolumeWithContext(ctx context.Context, request *ModifyVolumeRequest) (response *ModifyVolumeResponse, err error) {
    if request == nil {
        request = NewModifyVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewModifyVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeEbsInstancesRequest() (request *DescribeEbsInstancesRequest) {
    request = &DescribeEbsInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DescribeEbsInstances")
    return
}

func NewDescribeEbsInstancesResponse() (response *DescribeEbsInstancesResponse) {
    response = &DescribeEbsInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEbsInstances(request *DescribeEbsInstancesRequest) (response *DescribeEbsInstancesResponse, err error) {
    return c.DescribeEbsInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeEbsInstancesWithContext(ctx context.Context, request *DescribeEbsInstancesRequest) (response *DescribeEbsInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEbsInstancesRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeEbsInstancesResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeInstanceVolumesRequest() (request *DescribeInstanceVolumesRequest) {
    request = &DescribeInstanceVolumesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DescribeInstanceVolumes")
    return
}

func NewDescribeInstanceVolumesResponse() (response *DescribeInstanceVolumesResponse) {
    response = &DescribeInstanceVolumesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceVolumes(request *DescribeInstanceVolumesRequest) (response *DescribeInstanceVolumesResponse, err error) {
    return c.DescribeInstanceVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceVolumesWithContext(ctx context.Context, request *DescribeInstanceVolumesRequest) (response *DescribeInstanceVolumesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVolumesRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeInstanceVolumesResponse()
    err = c.Send(request, response)
    return
}
func NewRenewVolumeRequest() (request *RenewVolumeRequest) {
    request = &RenewVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "RenewVolume")
    return
}

func NewRenewVolumeResponse() (response *RenewVolumeResponse) {
    response = &RenewVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenewVolume(request *RenewVolumeRequest) (response *RenewVolumeResponse, err error) {
    return c.RenewVolumeWithContext(context.Background(), request)
}

func (c *Client) RenewVolumeWithContext(ctx context.Context, request *RenewVolumeRequest) (response *RenewVolumeResponse, err error) {
    if request == nil {
        request = NewRenewVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewRenewVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewUpdateVolumeProjectRequest() (request *UpdateVolumeProjectRequest) {
    request = &UpdateVolumeProjectRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "UpdateVolumeProject")
    return
}

func NewUpdateVolumeProjectResponse() (response *UpdateVolumeProjectResponse) {
    response = &UpdateVolumeProjectResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateVolumeProject(request *UpdateVolumeProjectRequest) (response *UpdateVolumeProjectResponse, err error) {
    return c.UpdateVolumeProjectWithContext(context.Background(), request)
}

func (c *Client) UpdateVolumeProjectWithContext(ctx context.Context, request *UpdateVolumeProjectRequest) (response *UpdateVolumeProjectResponse, err error) {
    if request == nil {
        request = NewUpdateVolumeProjectRequest()
    }
    request.SetContext(ctx)

    response = NewUpdateVolumeProjectResponse()
    err = c.Send(request, response)
    return
}
func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DescribeSnapshots")
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    request.SetContext(ctx)

    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}
func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
    request = &CreateSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "CreateSnapshot")
    return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
    response = &CreateSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    return c.CreateSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotRequest()
    }
    request.SetContext(ctx)

    response = NewCreateSnapshotResponse()
    err = c.Send(request, response)
    return
}
func NewDeleteSnapshotRequest() (request *DeleteSnapshotRequest) {
    request = &DeleteSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "DeleteSnapshot")
    return
}

func NewDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
    response = &DeleteSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (response *DeleteSnapshotResponse, err error) {
    return c.DeleteSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest) (response *DeleteSnapshotResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotRequest()
    }
    request.SetContext(ctx)

    response = NewDeleteSnapshotResponse()
    err = c.Send(request, response)
    return
}
func NewRollbackSnapshotRequest() (request *RollbackSnapshotRequest) {
    request = &RollbackSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "RollbackSnapshot")
    return
}

func NewRollbackSnapshotResponse() (response *RollbackSnapshotResponse) {
    response = &RollbackSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RollbackSnapshot(request *RollbackSnapshotRequest) (response *RollbackSnapshotResponse, err error) {
    return c.RollbackSnapshotWithContext(context.Background(), request)
}

func (c *Client) RollbackSnapshotWithContext(ctx context.Context, request *RollbackSnapshotRequest) (response *RollbackSnapshotResponse, err error) {
    if request == nil {
        request = NewRollbackSnapshotRequest()
    }
    request.SetContext(ctx)

    response = NewRollbackSnapshotResponse()
    err = c.Send(request, response)
    return
}
func NewModifySnapshotRequest() (request *ModifySnapshotRequest) {
    request = &ModifySnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "ModifySnapshot")
    return
}

func NewModifySnapshotResponse() (response *ModifySnapshotResponse) {
    response = &ModifySnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifySnapshot(request *ModifySnapshotRequest) (response *ModifySnapshotResponse, err error) {
    return c.ModifySnapshotWithContext(context.Background(), request)
}

func (c *Client) ModifySnapshotWithContext(ctx context.Context, request *ModifySnapshotRequest) (response *ModifySnapshotResponse, err error) {
    if request == nil {
        request = NewModifySnapshotRequest()
    }
    request.SetContext(ctx)

    response = NewModifySnapshotResponse()
    err = c.Send(request, response)
    return
}
func NewRecoveryVolumeRequest() (request *RecoveryVolumeRequest) {
    request = &RecoveryVolumeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "RecoveryVolume")
    return
}

func NewRecoveryVolumeResponse() (response *RecoveryVolumeResponse) {
    response = &RecoveryVolumeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RecoveryVolume(request *RecoveryVolumeRequest) (response *RecoveryVolumeResponse, err error) {
    return c.RecoveryVolumeWithContext(context.Background(), request)
}

func (c *Client) RecoveryVolumeWithContext(ctx context.Context, request *RecoveryVolumeRequest) (response *RecoveryVolumeResponse, err error) {
    if request == nil {
        request = NewRecoveryVolumeRequest()
    }
    request.SetContext(ctx)

    response = NewRecoveryVolumeResponse()
    err = c.Send(request, response)
    return
}
func NewValidateAttachInstanceRequest() (request *ValidateAttachInstanceRequest) {
    request = &ValidateAttachInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ebs", APIVersion, "ValidateAttachInstance")
    return
}

func NewValidateAttachInstanceResponse() (response *ValidateAttachInstanceResponse) {
    response = &ValidateAttachInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ValidateAttachInstance(request *ValidateAttachInstanceRequest) (response *ValidateAttachInstanceResponse, err error) {
    return c.ValidateAttachInstanceWithContext(context.Background(), request)
}

func (c *Client) ValidateAttachInstanceWithContext(ctx context.Context, request *ValidateAttachInstanceRequest) (response *ValidateAttachInstanceResponse, err error) {
    if request == nil {
        request = NewValidateAttachInstanceRequest()
    }
    request.SetContext(ctx)

    response = NewValidateAttachInstanceResponse()
    err = c.Send(request, response)
    return
}


