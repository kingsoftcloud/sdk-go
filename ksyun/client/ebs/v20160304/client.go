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

func (c *Client) CreateVolume(request *CreateVolumeRequest) string {
	return c.CreateVolumeWithContext(context.Background(), request)
}

func (c *Client) CreateVolumeWithContext(ctx context.Context, request *CreateVolumeRequest) string {
	if request == nil {
		request = NewCreateVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateVolumeResponse()
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
	request.Init().WithApiInfo("ebs", APIVersion, "AttachVolume")
	return
}

func NewAttachVolumeResponse() (response *AttachVolumeResponse) {
	response = &AttachVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachVolume(request *AttachVolumeRequest) string {
	return c.AttachVolumeWithContext(context.Background(), request)
}

func (c *Client) AttachVolumeWithContext(ctx context.Context, request *AttachVolumeRequest) string {
	if request == nil {
		request = NewAttachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

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
	request.Init().WithApiInfo("ebs", APIVersion, "DetachVolume")
	return
}

func NewDetachVolumeResponse() (response *DetachVolumeResponse) {
	response = &DetachVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachVolume(request *DetachVolumeRequest) string {
	return c.DetachVolumeWithContext(context.Background(), request)
}

func (c *Client) DetachVolumeWithContext(ctx context.Context, request *DetachVolumeRequest) string {
	if request == nil {
		request = NewDetachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDetachVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) DeleteVolume(request *DeleteVolumeRequest) string {
	return c.DeleteVolumeWithContext(context.Background(), request)
}

func (c *Client) DeleteVolumeWithContext(ctx context.Context, request *DeleteVolumeRequest) string {
	if request == nil {
		request = NewDeleteVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) ResizeVolume(request *ResizeVolumeRequest) string {
	return c.ResizeVolumeWithContext(context.Background(), request)
}

func (c *Client) ResizeVolumeWithContext(ctx context.Context, request *ResizeVolumeRequest) string {
	if request == nil {
		request = NewResizeVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResizeVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) DescribeVolumes(request *DescribeVolumesRequest) string {
	return c.DescribeVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeVolumesWithContext(ctx context.Context, request *DescribeVolumesRequest) string {
	if request == nil {
		request = NewDescribeVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeVolumesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) ModifyVolume(request *ModifyVolumeRequest) string {
	return c.ModifyVolumeWithContext(context.Background(), request)
}

func (c *Client) ModifyVolumeWithContext(ctx context.Context, request *ModifyVolumeRequest) string {
	if request == nil {
		request = NewModifyVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) DescribeEbsInstances(request *DescribeEbsInstancesRequest) string {
	return c.DescribeEbsInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeEbsInstancesWithContext(ctx context.Context, request *DescribeEbsInstancesRequest) string {
	if request == nil {
		request = NewDescribeEbsInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeEbsInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) DescribeInstanceVolumes(request *DescribeInstanceVolumesRequest) string {
	return c.DescribeInstanceVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceVolumesWithContext(ctx context.Context, request *DescribeInstanceVolumesRequest) string {
	if request == nil {
		request = NewDescribeInstanceVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceVolumesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) RenewVolume(request *RenewVolumeRequest) string {
	return c.RenewVolumeWithContext(context.Background(), request)
}

func (c *Client) RenewVolumeWithContext(ctx context.Context, request *RenewVolumeRequest) string {
	if request == nil {
		request = NewRenewVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenewVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) UpdateVolumeProject(request *UpdateVolumeProjectRequest) string {
	return c.UpdateVolumeProjectWithContext(context.Background(), request)
}

func (c *Client) UpdateVolumeProjectWithContext(ctx context.Context, request *UpdateVolumeProjectRequest) string {
	if request == nil {
		request = NewUpdateVolumeProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateVolumeProjectResponse()
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
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeSnapshots")
	return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) string {
	return c.DescribeSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) string {
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

func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) string {
	return c.CreateSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest) string {
	if request == nil {
		request = NewCreateSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSnapshotResponse()
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
	request.Init().WithApiInfo("ebs", APIVersion, "DeleteSnapshot")
	return
}

func NewDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
	response = &DeleteSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) string {
	return c.DeleteSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest) string {
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

func (c *Client) RollbackSnapshot(request *RollbackSnapshotRequest) string {
	return c.RollbackSnapshotWithContext(context.Background(), request)
}

func (c *Client) RollbackSnapshotWithContext(ctx context.Context, request *RollbackSnapshotRequest) string {
	if request == nil {
		request = NewRollbackSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRollbackSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) ModifySnapshot(request *ModifySnapshotRequest) string {
	return c.ModifySnapshotWithContext(context.Background(), request)
}

func (c *Client) ModifySnapshotWithContext(ctx context.Context, request *ModifySnapshotRequest) string {
	if request == nil {
		request = NewModifySnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifySnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) RecoveryVolume(request *RecoveryVolumeRequest) string {
	return c.RecoveryVolumeWithContext(context.Background(), request)
}

func (c *Client) RecoveryVolumeWithContext(ctx context.Context, request *RecoveryVolumeRequest) string {
	if request == nil {
		request = NewRecoveryVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRecoveryVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) ValidateAttachInstance(request *ValidateAttachInstanceRequest) string {
	return c.ValidateAttachInstanceWithContext(context.Background(), request)
}

func (c *Client) ValidateAttachInstanceWithContext(ctx context.Context, request *ValidateAttachInstanceRequest) string {
	if request == nil {
		request = NewValidateAttachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewValidateAttachInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCreateVolumePriceRequest() (request *DescribeCreateVolumePriceRequest) {
	request = &DescribeCreateVolumePriceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeCreateVolumePrice")
	return
}

func NewDescribeCreateVolumePriceResponse() (response *DescribeCreateVolumePriceResponse) {
	response = &DescribeCreateVolumePriceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCreateVolumePrice(request *DescribeCreateVolumePriceRequest) string {
	return c.DescribeCreateVolumePriceWithContext(context.Background(), request)
}

func (c *Client) DescribeCreateVolumePriceWithContext(ctx context.Context, request *DescribeCreateVolumePriceRequest) string {
	if request == nil {
		request = NewDescribeCreateVolumePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCreateVolumePriceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifySnapshotTypeRequest() (request *ModifySnapshotTypeRequest) {
	request = &ModifySnapshotTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifySnapshotType")
	return
}

func NewModifySnapshotTypeResponse() (response *ModifySnapshotTypeResponse) {
	response = &ModifySnapshotTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySnapshotType(request *ModifySnapshotTypeRequest) string {
	return c.ModifySnapshotTypeWithContext(context.Background(), request)
}

func (c *Client) ModifySnapshotTypeWithContext(ctx context.Context, request *ModifySnapshotTypeRequest) string {
	if request == nil {
		request = NewModifySnapshotTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifySnapshotTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyVolumeTypeRequest() (request *ModifyVolumeTypeRequest) {
	request = &ModifyVolumeTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifyVolumeType")
	return
}

func NewModifyVolumeTypeResponse() (response *ModifyVolumeTypeResponse) {
	response = &ModifyVolumeTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVolumeType(request *ModifyVolumeTypeRequest) string {
	return c.ModifyVolumeTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyVolumeTypeWithContext(ctx context.Context, request *ModifyVolumeTypeRequest) string {
	if request == nil {
		request = NewModifyVolumeTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyVolumeTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDedicatedBlockStorageClusterAttributeRequest() (request *ModifyDedicatedBlockStorageClusterAttributeRequest) {
	request = &ModifyDedicatedBlockStorageClusterAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifyDedicatedBlockStorageClusterAttribute")
	return
}

func NewModifyDedicatedBlockStorageClusterAttributeResponse() (response *ModifyDedicatedBlockStorageClusterAttributeResponse) {
	response = &ModifyDedicatedBlockStorageClusterAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDedicatedBlockStorageClusterAttribute(request *ModifyDedicatedBlockStorageClusterAttributeRequest) string {
	return c.ModifyDedicatedBlockStorageClusterAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyDedicatedBlockStorageClusterAttributeWithContext(ctx context.Context, request *ModifyDedicatedBlockStorageClusterAttributeRequest) string {
	if request == nil {
		request = NewModifyDedicatedBlockStorageClusterAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDedicatedBlockStorageClusterAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewResizeDedicatedBlockStorageClustersRequest() (request *ResizeDedicatedBlockStorageClustersRequest) {
	request = &ResizeDedicatedBlockStorageClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ResizeDedicatedBlockStorageClusters")
	return
}

func NewResizeDedicatedBlockStorageClustersResponse() (response *ResizeDedicatedBlockStorageClustersResponse) {
	response = &ResizeDedicatedBlockStorageClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResizeDedicatedBlockStorageClusters(request *ResizeDedicatedBlockStorageClustersRequest) string {
	return c.ResizeDedicatedBlockStorageClustersWithContext(context.Background(), request)
}

func (c *Client) ResizeDedicatedBlockStorageClustersWithContext(ctx context.Context, request *ResizeDedicatedBlockStorageClustersRequest) string {
	if request == nil {
		request = NewResizeDedicatedBlockStorageClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResizeDedicatedBlockStorageClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDedicatedBlockStorageClustersRequest() (request *DescribeDedicatedBlockStorageClustersRequest) {
	request = &DescribeDedicatedBlockStorageClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeDedicatedBlockStorageClusters")
	return
}

func NewDescribeDedicatedBlockStorageClustersResponse() (response *DescribeDedicatedBlockStorageClustersResponse) {
	response = &DescribeDedicatedBlockStorageClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedBlockStorageClusters(request *DescribeDedicatedBlockStorageClustersRequest) string {
	return c.DescribeDedicatedBlockStorageClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedBlockStorageClustersWithContext(ctx context.Context, request *DescribeDedicatedBlockStorageClustersRequest) string {
	if request == nil {
		request = NewDescribeDedicatedBlockStorageClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedBlockStorageClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDedicatedBlockStorageClusterRequest() (request *CreateDedicatedBlockStorageClusterRequest) {
	request = &CreateDedicatedBlockStorageClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "CreateDedicatedBlockStorageCluster")
	return
}

func NewCreateDedicatedBlockStorageClusterResponse() (response *CreateDedicatedBlockStorageClusterResponse) {
	response = &CreateDedicatedBlockStorageClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDedicatedBlockStorageCluster(request *CreateDedicatedBlockStorageClusterRequest) string {
	return c.CreateDedicatedBlockStorageClusterWithContext(context.Background(), request)
}

func (c *Client) CreateDedicatedBlockStorageClusterWithContext(ctx context.Context, request *CreateDedicatedBlockStorageClusterRequest) string {
	if request == nil {
		request = NewCreateDedicatedBlockStorageClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedBlockStorageClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
