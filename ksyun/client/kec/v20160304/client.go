package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) string {
	return c.DescribeInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) string {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RunInstances")
	return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RunInstances(request *RunInstancesRequest) string {
	return c.RunInstancesWithContext(context.Background(), request)
}

func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) string {
	if request == nil {
		request = NewRunInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRunInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "StartInstances")
	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartInstances(request *StartInstancesRequest) string {
	return c.StartInstancesWithContext(context.Background(), request)
}

func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) string {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "StopInstances")
	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopInstances(request *StopInstancesRequest) string {
	return c.StopInstancesWithContext(context.Background(), request)
}

func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) string {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RebootInstances")
	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootInstances(request *RebootInstancesRequest) string {
	return c.RebootInstancesWithContext(context.Background(), request)
}

func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) string {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceAttributeRequest() (request *ModifyInstanceAttributeRequest) {
	request = &ModifyInstanceAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceAttribute")
	return
}

func NewModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
	response = &ModifyInstanceAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) string {
	return c.ModifyInstanceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAttributeWithContext(ctx context.Context, request *ModifyInstanceAttributeRequest) string {
	if request == nil {
		request = NewModifyInstanceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceTypeRequest() (request *ModifyInstanceTypeRequest) {
	request = &ModifyInstanceTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceType")
	return
}

func NewModifyInstanceTypeResponse() (response *ModifyInstanceTypeResponse) {
	response = &ModifyInstanceTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) string {
	return c.ModifyInstanceTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceTypeWithContext(ctx context.Context, request *ModifyInstanceTypeRequest) string {
	if request == nil {
		request = NewModifyInstanceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TerminateInstances(request *TerminateInstancesRequest) string {
	return c.TerminateInstancesWithContext(context.Background(), request)
}

func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) string {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTerminateInstancesResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) string {
	return c.DescribeImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) string {
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
func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
	request = &ModifyImageAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyImageAttribute")
	return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
	response = &ModifyImageAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) string {
	return c.ModifyImageAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest) string {
	if request == nil {
		request = NewModifyImageAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceImageRequest() (request *ModifyInstanceImageRequest) {
	request = &ModifyInstanceImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceImage")
	return
}

func NewModifyInstanceImageResponse() (response *ModifyInstanceImageResponse) {
	response = &ModifyInstanceImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceImage(request *ModifyInstanceImageRequest) string {
	return c.ModifyInstanceImageWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceImageWithContext(ctx context.Context, request *ModifyInstanceImageRequest) string {
	if request == nil {
		request = NewModifyInstanceImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceImageResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImage(request *CreateImageRequest) string {
	return c.CreateImageWithContext(context.Background(), request)
}

func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) string {
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
func NewRemoveImagesRequest() (request *RemoveImagesRequest) {
	request = &RemoveImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RemoveImages")
	return
}

func NewRemoveImagesResponse() (response *RemoveImagesResponse) {
	response = &RemoveImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemoveImages(request *RemoveImagesRequest) string {
	return c.RemoveImagesWithContext(context.Background(), request)
}

func (c *Client) RemoveImagesWithContext(ctx context.Context, request *RemoveImagesRequest) string {
	if request == nil {
		request = NewRemoveImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemoveImagesResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "ModifyNetworkInterfaceAttribute")
	return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
	response = &ModifyNetworkInterfaceAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) string {
	return c.ModifyNetworkInterfaceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) string {
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
func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
	request = &AttachNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachNetworkInterface")
	return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
	response = &AttachNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) string {
	return c.AttachNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) AttachNetworkInterfaceWithContext(ctx context.Context, request *AttachNetworkInterfaceRequest) string {
	if request == nil {
		request = NewAttachNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
	request = &DetachNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachNetworkInterface")
	return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
	response = &DetachNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) string {
	return c.DetachNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) DetachNetworkInterfaceWithContext(ctx context.Context, request *DetachNetworkInterfaceRequest) string {
	if request == nil {
		request = NewDetachNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeLocalVolumesRequest() (request *DescribeLocalVolumesRequest) {
	request = &DescribeLocalVolumesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeLocalVolumes")
	return
}

func NewDescribeLocalVolumesResponse() (response *DescribeLocalVolumesResponse) {
	response = &DescribeLocalVolumesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLocalVolumes(request *DescribeLocalVolumesRequest) string {
	return c.DescribeLocalVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeLocalVolumesWithContext(ctx context.Context, request *DescribeLocalVolumesRequest) string {
	if request == nil {
		request = NewDescribeLocalVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLocalVolumesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateLocalVolumeSnapshotRequest() (request *CreateLocalVolumeSnapshotRequest) {
	request = &CreateLocalVolumeSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateLocalVolumeSnapshot")
	return
}

func NewCreateLocalVolumeSnapshotResponse() (response *CreateLocalVolumeSnapshotResponse) {
	response = &CreateLocalVolumeSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLocalVolumeSnapshot(request *CreateLocalVolumeSnapshotRequest) string {
	return c.CreateLocalVolumeSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateLocalVolumeSnapshotWithContext(ctx context.Context, request *CreateLocalVolumeSnapshotRequest) string {
	if request == nil {
		request = NewCreateLocalVolumeSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLocalVolumeSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeLocalVolumeSnapshotsRequest() (request *DescribeLocalVolumeSnapshotsRequest) {
	request = &DescribeLocalVolumeSnapshotsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeLocalVolumeSnapshots")
	return
}

func NewDescribeLocalVolumeSnapshotsResponse() (response *DescribeLocalVolumeSnapshotsResponse) {
	response = &DescribeLocalVolumeSnapshotsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLocalVolumeSnapshots(request *DescribeLocalVolumeSnapshotsRequest) string {
	return c.DescribeLocalVolumeSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeLocalVolumeSnapshotsWithContext(ctx context.Context, request *DescribeLocalVolumeSnapshotsRequest) string {
	if request == nil {
		request = NewDescribeLocalVolumeSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLocalVolumeSnapshotsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRollbackLocalVolumeRequest() (request *RollbackLocalVolumeRequest) {
	request = &RollbackLocalVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RollbackLocalVolume")
	return
}

func NewRollbackLocalVolumeResponse() (response *RollbackLocalVolumeResponse) {
	response = &RollbackLocalVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RollbackLocalVolume(request *RollbackLocalVolumeRequest) string {
	return c.RollbackLocalVolumeWithContext(context.Background(), request)
}

func (c *Client) RollbackLocalVolumeWithContext(ctx context.Context, request *RollbackLocalVolumeRequest) string {
	if request == nil {
		request = NewRollbackLocalVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackLocalVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteLocalVolumeSnapshotRequest() (request *DeleteLocalVolumeSnapshotRequest) {
	request = &DeleteLocalVolumeSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteLocalVolumeSnapshot")
	return
}

func NewDeleteLocalVolumeSnapshotResponse() (response *DeleteLocalVolumeSnapshotResponse) {
	response = &DeleteLocalVolumeSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLocalVolumeSnapshot(request *DeleteLocalVolumeSnapshotRequest) string {
	return c.DeleteLocalVolumeSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteLocalVolumeSnapshotWithContext(ctx context.Context, request *DeleteLocalVolumeSnapshotRequest) string {
	if request == nil {
		request = NewDeleteLocalVolumeSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLocalVolumeSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDataGuardGroupsRequest() (request *ModifyDataGuardGroupsRequest) {
	request = &ModifyDataGuardGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyDataGuardGroups")
	return
}

func NewModifyDataGuardGroupsResponse() (response *ModifyDataGuardGroupsResponse) {
	response = &ModifyDataGuardGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDataGuardGroups(request *ModifyDataGuardGroupsRequest) string {
	return c.ModifyDataGuardGroupsWithContext(context.Background(), request)
}

func (c *Client) ModifyDataGuardGroupsWithContext(ctx context.Context, request *ModifyDataGuardGroupsRequest) string {
	if request == nil {
		request = NewModifyDataGuardGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDataGuardGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDataGuardCapacityRequest() (request *DescribeDataGuardCapacityRequest) {
	request = &DescribeDataGuardCapacityRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDataGuardCapacity")
	return
}

func NewDescribeDataGuardCapacityResponse() (response *DescribeDataGuardCapacityResponse) {
	response = &DescribeDataGuardCapacityResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataGuardCapacity(request *DescribeDataGuardCapacityRequest) string {
	return c.DescribeDataGuardCapacityWithContext(context.Background(), request)
}

func (c *Client) DescribeDataGuardCapacityWithContext(ctx context.Context, request *DescribeDataGuardCapacityRequest) string {
	if request == nil {
		request = NewDescribeDataGuardCapacityRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataGuardCapacityResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDataGuardGroupRequest() (request *CreateDataGuardGroupRequest) {
	request = &CreateDataGuardGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateDataGuardGroup")
	return
}

func NewCreateDataGuardGroupResponse() (response *CreateDataGuardGroupResponse) {
	response = &CreateDataGuardGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataGuardGroup(request *CreateDataGuardGroupRequest) string {
	return c.CreateDataGuardGroupWithContext(context.Background(), request)
}

func (c *Client) CreateDataGuardGroupWithContext(ctx context.Context, request *CreateDataGuardGroupRequest) string {
	if request == nil {
		request = NewCreateDataGuardGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDataGuardGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDataGuardGroupsRequest() (request *DeleteDataGuardGroupsRequest) {
	request = &DeleteDataGuardGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteDataGuardGroups")
	return
}

func NewDeleteDataGuardGroupsResponse() (response *DeleteDataGuardGroupsResponse) {
	response = &DeleteDataGuardGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDataGuardGroups(request *DeleteDataGuardGroupsRequest) string {
	return c.DeleteDataGuardGroupsWithContext(context.Background(), request)
}

func (c *Client) DeleteDataGuardGroupsWithContext(ctx context.Context, request *DeleteDataGuardGroupsRequest) string {
	if request == nil {
		request = NewDeleteDataGuardGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataGuardGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDataGuardGroupRequest() (request *DescribeDataGuardGroupRequest) {
	request = &DescribeDataGuardGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDataGuardGroup")
	return
}

func NewDescribeDataGuardGroupResponse() (response *DescribeDataGuardGroupResponse) {
	response = &DescribeDataGuardGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataGuardGroup(request *DescribeDataGuardGroupRequest) string {
	return c.DescribeDataGuardGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeDataGuardGroupWithContext(ctx context.Context, request *DescribeDataGuardGroupRequest) string {
	if request == nil {
		request = NewDescribeDataGuardGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataGuardGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRemoveVmFromDataGuardRequest() (request *RemoveVmFromDataGuardRequest) {
	request = &RemoveVmFromDataGuardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RemoveVmFromDataGuard")
	return
}

func NewRemoveVmFromDataGuardResponse() (response *RemoveVmFromDataGuardResponse) {
	response = &RemoveVmFromDataGuardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemoveVmFromDataGuard(request *RemoveVmFromDataGuardRequest) string {
	return c.RemoveVmFromDataGuardWithContext(context.Background(), request)
}

func (c *Client) RemoveVmFromDataGuardWithContext(ctx context.Context, request *RemoveVmFromDataGuardRequest) string {
	if request == nil {
		request = NewRemoveVmFromDataGuardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemoveVmFromDataGuardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDedicatedHostsRequest() (request *CreateDedicatedHostsRequest) {
	request = &CreateDedicatedHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateDedicatedHosts")
	return
}

func NewCreateDedicatedHostsResponse() (response *CreateDedicatedHostsResponse) {
	response = &CreateDedicatedHostsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDedicatedHosts(request *CreateDedicatedHostsRequest) string {
	return c.CreateDedicatedHostsWithContext(context.Background(), request)
}

func (c *Client) CreateDedicatedHostsWithContext(ctx context.Context, request *CreateDedicatedHostsRequest) string {
	if request == nil {
		request = NewCreateDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDedicatedHostRequest() (request *DeleteDedicatedHostRequest) {
	request = &DeleteDedicatedHostRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteDedicatedHost")
	return
}

func NewDeleteDedicatedHostResponse() (response *DeleteDedicatedHostResponse) {
	response = &DeleteDedicatedHostResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDedicatedHost(request *DeleteDedicatedHostRequest) string {
	return c.DeleteDedicatedHostWithContext(context.Background(), request)
}

func (c *Client) DeleteDedicatedHostWithContext(ctx context.Context, request *DeleteDedicatedHostRequest) string {
	if request == nil {
		request = NewDeleteDedicatedHostRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDedicatedHostResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDedicatedHostsRequest() (request *DescribeDedicatedHostsRequest) {
	request = &DescribeDedicatedHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDedicatedHosts")
	return
}

func NewDescribeDedicatedHostsResponse() (response *DescribeDedicatedHostsResponse) {
	response = &DescribeDedicatedHostsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedHosts(request *DescribeDedicatedHostsRequest) string {
	return c.DescribeDedicatedHostsWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedHostsWithContext(ctx context.Context, request *DescribeDedicatedHostsRequest) string {
	if request == nil {
		request = NewDescribeDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScalingConfigurationRequest() (request *DescribeScalingConfigurationRequest) {
	request = &DescribeScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingConfiguration")
	return
}

func NewDescribeScalingConfigurationResponse() (response *DescribeScalingConfigurationResponse) {
	response = &DescribeScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingConfiguration(request *DescribeScalingConfigurationRequest) string {
	return c.DescribeScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingConfigurationWithContext(ctx context.Context, request *DescribeScalingConfigurationRequest) string {
	if request == nil {
		request = NewDescribeScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateScalingConfigurationRequest() (request *CreateScalingConfigurationRequest) {
	request = &CreateScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingConfiguration")
	return
}

func NewCreateScalingConfigurationResponse() (response *CreateScalingConfigurationResponse) {
	response = &CreateScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) string {
	return c.CreateScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) CreateScalingConfigurationWithContext(ctx context.Context, request *CreateScalingConfigurationRequest) string {
	if request == nil {
		request = NewCreateScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteScalingConfigurationRequest() (request *DeleteScalingConfigurationRequest) {
	request = &DeleteScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScalingConfiguration")
	return
}

func NewDeleteScalingConfigurationResponse() (response *DeleteScalingConfigurationResponse) {
	response = &DeleteScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScalingConfiguration(request *DeleteScalingConfigurationRequest) string {
	return c.DeleteScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) DeleteScalingConfigurationWithContext(ctx context.Context, request *DeleteScalingConfigurationRequest) string {
	if request == nil {
		request = NewDeleteScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateScalingGroupRequest() (request *CreateScalingGroupRequest) {
	request = &CreateScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingGroup")
	return
}

func NewCreateScalingGroupResponse() (response *CreateScalingGroupResponse) {
	response = &CreateScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingGroup(request *CreateScalingGroupRequest) string {
	return c.CreateScalingGroupWithContext(context.Background(), request)
}

func (c *Client) CreateScalingGroupWithContext(ctx context.Context, request *CreateScalingGroupRequest) string {
	if request == nil {
		request = NewCreateScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScalingGroupRequest() (request *DescribeScalingGroupRequest) {
	request = &DescribeScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingGroup")
	return
}

func NewDescribeScalingGroupResponse() (response *DescribeScalingGroupResponse) {
	response = &DescribeScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingGroup(request *DescribeScalingGroupRequest) string {
	return c.DescribeScalingGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingGroupWithContext(ctx context.Context, request *DescribeScalingGroupRequest) string {
	if request == nil {
		request = NewDescribeScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyScalingGroupRequest() (request *ModifyScalingGroupRequest) {
	request = &ModifyScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingGroup")
	return
}

func NewModifyScalingGroupResponse() (response *ModifyScalingGroupResponse) {
	response = &ModifyScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) string {
	return c.ModifyScalingGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingGroupWithContext(ctx context.Context, request *ModifyScalingGroupRequest) string {
	if request == nil {
		request = NewModifyScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSetKvmProtectedDetachRequest() (request *SetKvmProtectedDetachRequest) {
	request = &SetKvmProtectedDetachRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SetKvmProtectedDetach")
	return
}

func NewSetKvmProtectedDetachResponse() (response *SetKvmProtectedDetachResponse) {
	response = &SetKvmProtectedDetachResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetKvmProtectedDetach(request *SetKvmProtectedDetachRequest) string {
	return c.SetKvmProtectedDetachWithContext(context.Background(), request)
}

func (c *Client) SetKvmProtectedDetachWithContext(ctx context.Context, request *SetKvmProtectedDetachRequest) string {
	if request == nil {
		request = NewSetKvmProtectedDetachRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetKvmProtectedDetachResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScalingInstanceRequest() (request *DescribeScalingInstanceRequest) {
	request = &DescribeScalingInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingInstance")
	return
}

func NewDescribeScalingInstanceResponse() (response *DescribeScalingInstanceResponse) {
	response = &DescribeScalingInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingInstance(request *DescribeScalingInstanceRequest) string {
	return c.DescribeScalingInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingInstanceWithContext(ctx context.Context, request *DescribeScalingInstanceRequest) string {
	if request == nil {
		request = NewDescribeScalingInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAttachInstanceRequest() (request *AttachInstanceRequest) {
	request = &AttachInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachInstance")
	return
}

func NewAttachInstanceResponse() (response *AttachInstanceResponse) {
	response = &AttachInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachInstance(request *AttachInstanceRequest) string {
	return c.AttachInstanceWithContext(context.Background(), request)
}

func (c *Client) AttachInstanceWithContext(ctx context.Context, request *AttachInstanceRequest) string {
	if request == nil {
		request = NewAttachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachInstanceRequest() (request *DetachInstanceRequest) {
	request = &DetachInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachInstance")
	return
}

func NewDetachInstanceResponse() (response *DetachInstanceResponse) {
	response = &DetachInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachInstance(request *DetachInstanceRequest) string {
	return c.DetachInstanceWithContext(context.Background(), request)
}

func (c *Client) DetachInstanceWithContext(ctx context.Context, request *DetachInstanceRequest) string {
	if request == nil {
		request = NewDetachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScalingActivityRequest() (request *DescribeScalingActivityRequest) {
	request = &DescribeScalingActivityRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingActivity")
	return
}

func NewDescribeScalingActivityResponse() (response *DescribeScalingActivityResponse) {
	response = &DescribeScalingActivityResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingActivity(request *DescribeScalingActivityRequest) string {
	return c.DescribeScalingActivityWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingActivityWithContext(ctx context.Context, request *DescribeScalingActivityRequest) string {
	if request == nil {
		request = NewDescribeScalingActivityRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingActivityResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteScalingGroupRequest() (request *DeleteScalingGroupRequest) {
	request = &DeleteScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScalingGroup")
	return
}

func NewDeleteScalingGroupResponse() (response *DeleteScalingGroupResponse) {
	response = &DeleteScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScalingGroup(request *DeleteScalingGroupRequest) string {
	return c.DeleteScalingGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteScalingGroupWithContext(ctx context.Context, request *DeleteScalingGroupRequest) string {
	if request == nil {
		request = NewDeleteScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisableScalingGroupRequest() (request *DisableScalingGroupRequest) {
	request = &DisableScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DisableScalingGroup")
	return
}

func NewDisableScalingGroupResponse() (response *DisableScalingGroupResponse) {
	response = &DisableScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableScalingGroup(request *DisableScalingGroupRequest) string {
	return c.DisableScalingGroupWithContext(context.Background(), request)
}

func (c *Client) DisableScalingGroupWithContext(ctx context.Context, request *DisableScalingGroupRequest) string {
	if request == nil {
		request = NewDisableScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewEnableScalingGroupRequest() (request *EnableScalingGroupRequest) {
	request = &EnableScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "EnableScalingGroup")
	return
}

func NewEnableScalingGroupResponse() (response *EnableScalingGroupResponse) {
	response = &EnableScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableScalingGroup(request *EnableScalingGroupRequest) string {
	return c.EnableScalingGroupWithContext(context.Background(), request)
}

func (c *Client) EnableScalingGroupWithContext(ctx context.Context, request *EnableScalingGroupRequest) string {
	if request == nil {
		request = NewEnableScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScalingNotificationRequest() (request *DescribeScalingNotificationRequest) {
	request = &DescribeScalingNotificationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingNotification")
	return
}

func NewDescribeScalingNotificationResponse() (response *DescribeScalingNotificationResponse) {
	response = &DescribeScalingNotificationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingNotification(request *DescribeScalingNotificationRequest) string {
	return c.DescribeScalingNotificationWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingNotificationWithContext(ctx context.Context, request *DescribeScalingNotificationRequest) string {
	if request == nil {
		request = NewDescribeScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingNotificationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateScalingNotificationRequest() (request *CreateScalingNotificationRequest) {
	request = &CreateScalingNotificationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingNotification")
	return
}

func NewCreateScalingNotificationResponse() (response *CreateScalingNotificationResponse) {
	response = &CreateScalingNotificationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingNotification(request *CreateScalingNotificationRequest) string {
	return c.CreateScalingNotificationWithContext(context.Background(), request)
}

func (c *Client) CreateScalingNotificationWithContext(ctx context.Context, request *CreateScalingNotificationRequest) string {
	if request == nil {
		request = NewCreateScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingNotificationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyScalingNotificationRequest() (request *ModifyScalingNotificationRequest) {
	request = &ModifyScalingNotificationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingNotification")
	return
}

func NewModifyScalingNotificationResponse() (response *ModifyScalingNotificationResponse) {
	response = &ModifyScalingNotificationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingNotification(request *ModifyScalingNotificationRequest) string {
	return c.ModifyScalingNotificationWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingNotificationWithContext(ctx context.Context, request *ModifyScalingNotificationRequest) string {
	if request == nil {
		request = NewModifyScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingNotificationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateScheduledTaskRequest() (request *CreateScheduledTaskRequest) {
	request = &CreateScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScheduledTask")
	return
}

func NewCreateScheduledTaskResponse() (response *CreateScheduledTaskResponse) {
	response = &CreateScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScheduledTask(request *CreateScheduledTaskRequest) string {
	return c.CreateScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) CreateScheduledTaskWithContext(ctx context.Context, request *CreateScheduledTaskRequest) string {
	if request == nil {
		request = NewCreateScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScheduledTaskRequest() (request *DescribeScheduledTaskRequest) {
	request = &DescribeScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScheduledTask")
	return
}

func NewDescribeScheduledTaskResponse() (response *DescribeScheduledTaskResponse) {
	response = &DescribeScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScheduledTask(request *DescribeScheduledTaskRequest) string {
	return c.DescribeScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeScheduledTaskWithContext(ctx context.Context, request *DescribeScheduledTaskRequest) string {
	if request == nil {
		request = NewDescribeScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyScheduledTaskRequest() (request *ModifyScheduledTaskRequest) {
	request = &ModifyScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScheduledTask")
	return
}

func NewModifyScheduledTaskResponse() (response *ModifyScheduledTaskResponse) {
	response = &ModifyScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScheduledTask(request *ModifyScheduledTaskRequest) string {
	return c.ModifyScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) ModifyScheduledTaskWithContext(ctx context.Context, request *ModifyScheduledTaskRequest) string {
	if request == nil {
		request = NewModifyScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteScheduledTaskRequest() (request *DeleteScheduledTaskRequest) {
	request = &DeleteScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScheduledTask")
	return
}

func NewDeleteScheduledTaskResponse() (response *DeleteScheduledTaskResponse) {
	response = &DeleteScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScheduledTask(request *DeleteScheduledTaskRequest) string {
	return c.DeleteScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteScheduledTaskWithContext(ctx context.Context, request *DeleteScheduledTaskRequest) string {
	if request == nil {
		request = NewDeleteScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateScalingPolicyRequest() (request *CreateScalingPolicyRequest) {
	request = &CreateScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingPolicy")
	return
}

func NewCreateScalingPolicyResponse() (response *CreateScalingPolicyResponse) {
	response = &CreateScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingPolicy(request *CreateScalingPolicyRequest) string {
	return c.CreateScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) CreateScalingPolicyWithContext(ctx context.Context, request *CreateScalingPolicyRequest) string {
	if request == nil {
		request = NewCreateScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeScalingPolicyRequest() (request *DescribeScalingPolicyRequest) {
	request = &DescribeScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingPolicy")
	return
}

func NewDescribeScalingPolicyResponse() (response *DescribeScalingPolicyResponse) {
	response = &DescribeScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingPolicy(request *DescribeScalingPolicyRequest) string {
	return c.DescribeScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingPolicyWithContext(ctx context.Context, request *DescribeScalingPolicyRequest) string {
	if request == nil {
		request = NewDescribeScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyScalingPolicyRequest() (request *ModifyScalingPolicyRequest) {
	request = &ModifyScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingPolicy")
	return
}

func NewModifyScalingPolicyResponse() (response *ModifyScalingPolicyResponse) {
	response = &ModifyScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingPolicy(request *ModifyScalingPolicyRequest) string {
	return c.ModifyScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingPolicyWithContext(ctx context.Context, request *ModifyScalingPolicyRequest) string {
	if request == nil {
		request = NewModifyScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteScalingPolicyRequest() (request *DeleteScalingPolicyRequest) {
	request = &DeleteScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScalingPolicy")
	return
}

func NewDeleteScalingPolicyResponse() (response *DeleteScalingPolicyResponse) {
	response = &DeleteScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) string {
	return c.DeleteScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteScalingPolicyWithContext(ctx context.Context, request *DeleteScalingPolicyRequest) string {
	if request == nil {
		request = NewDeleteScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImportImageRequest() (request *ImportImageRequest) {
	request = &ImportImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ImportImage")
	return
}

func NewImportImageResponse() (response *ImportImageResponse) {
	response = &ImportImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportImage(request *ImportImageRequest) string {
	return c.ImportImageWithContext(context.Background(), request)
}

func (c *Client) ImportImageWithContext(ctx context.Context, request *ImportImageRequest) string {
	if request == nil {
		request = NewImportImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewImportImageResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "CopyImage")
	return
}

func NewCopyImageResponse() (response *CopyImageResponse) {
	response = &CopyImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CopyImage(request *CopyImageRequest) string {
	return c.CopyImageWithContext(context.Background(), request)
}

func (c *Client) CopyImageWithContext(ctx context.Context, request *CopyImageRequest) string {
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
func NewModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
	request = &ModifyImageSharePermissionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyImageSharePermission")
	return
}

func NewModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
	response = &ModifyImageSharePermissionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) string {
	return c.ModifyImageSharePermissionWithContext(context.Background(), request)
}

func (c *Client) ModifyImageSharePermissionWithContext(ctx context.Context, request *ModifyImageSharePermissionRequest) string {
	if request == nil {
		request = NewModifyImageSharePermissionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageSharePermissionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeImageSharePermissionRequest() (request *DescribeImageSharePermissionRequest) {
	request = &DescribeImageSharePermissionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeImageSharePermission")
	return
}

func NewDescribeImageSharePermissionResponse() (response *DescribeImageSharePermissionResponse) {
	response = &DescribeImageSharePermissionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) string {
	return c.DescribeImageSharePermissionWithContext(context.Background(), request)
}

func (c *Client) DescribeImageSharePermissionWithContext(ctx context.Context, request *DescribeImageSharePermissionRequest) string {
	if request == nil {
		request = NewDescribeImageSharePermissionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageSharePermissionResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegions(request *DescribeRegionsRequest) string {
	return c.DescribeRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) string {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAttachKeyRequest() (request *AttachKeyRequest) {
	request = &AttachKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachKey")
	return
}

func NewAttachKeyResponse() (response *AttachKeyResponse) {
	response = &AttachKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachKey(request *AttachKeyRequest) string {
	return c.AttachKeyWithContext(context.Background(), request)
}

func (c *Client) AttachKeyWithContext(ctx context.Context, request *AttachKeyRequest) string {
	if request == nil {
		request = NewAttachKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachKeyRequest() (request *DetachKeyRequest) {
	request = &DetachKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachKey")
	return
}

func NewDetachKeyResponse() (response *DetachKeyResponse) {
	response = &DetachKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachKey(request *DetachKeyRequest) string {
	return c.DetachKeyWithContext(context.Background(), request)
}

func (c *Client) DetachKeyWithContext(ctx context.Context, request *DetachKeyRequest) string {
	if request == nil {
		request = NewDetachKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachKeyResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "DescribeAvailabilityZones")
	return
}

func NewDescribeAvailabilityZonesResponse() (response *DescribeAvailabilityZonesResponse) {
	response = &DescribeAvailabilityZonesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAvailabilityZones(request *DescribeAvailabilityZonesRequest) string {
	return c.DescribeAvailabilityZonesWithContext(context.Background(), request)
}

func (c *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, request *DescribeAvailabilityZonesRequest) string {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAvailabilityZonesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
	request = &DescribeInstanceTypeConfigsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstanceTypeConfigs")
	return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
	response = &DescribeInstanceTypeConfigsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) string {
	return c.DescribeInstanceTypeConfigsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceTypeConfigsWithContext(ctx context.Context, request *DescribeInstanceTypeConfigsRequest) string {
	if request == nil {
		request = NewDescribeInstanceTypeConfigsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceTypeConfigsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceFamilysRequest() (request *DescribeInstanceFamilysRequest) {
	request = &DescribeInstanceFamilysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstanceFamilys")
	return
}

func NewDescribeInstanceFamilysResponse() (response *DescribeInstanceFamilysResponse) {
	response = &DescribeInstanceFamilysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceFamilys(request *DescribeInstanceFamilysRequest) string {
	return c.DescribeInstanceFamilysWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceFamilysWithContext(ctx context.Context, request *DescribeInstanceFamilysRequest) string {
	if request == nil {
		request = NewDescribeInstanceFamilysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceFamilysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAddVmIntoDataGuardRequest() (request *AddVmIntoDataGuardRequest) {
	request = &AddVmIntoDataGuardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AddVmIntoDataGuard")
	return
}

func NewAddVmIntoDataGuardResponse() (response *AddVmIntoDataGuardResponse) {
	response = &AddVmIntoDataGuardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddVmIntoDataGuard(request *AddVmIntoDataGuardRequest) string {
	return c.AddVmIntoDataGuardWithContext(context.Background(), request)
}

func (c *Client) AddVmIntoDataGuardWithContext(ctx context.Context, request *AddVmIntoDataGuardRequest) string {
	if request == nil {
		request = NewAddVmIntoDataGuardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddVmIntoDataGuardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
	request = &CreateFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateFileSystem")
	return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
	response = &CreateFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) string {
	return c.CreateFileSystemWithContext(context.Background(), request)
}

func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) string {
	if request == nil {
		request = NewCreateFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
	request = &DeleteFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteFileSystem")
	return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
	response = &DeleteFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) string {
	return c.DeleteFileSystemWithContext(context.Background(), request)
}

func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) string {
	if request == nil {
		request = NewDeleteFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
	request = &DescribeFileSystemsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeFileSystems")
	return
}

func NewDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
	response = &DescribeFileSystemsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) string {
	return c.DescribeFileSystemsWithContext(context.Background(), request)
}

func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest) string {
	if request == nil {
		request = NewDescribeFileSystemsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyFileSystemRequest() (request *ModifyFileSystemRequest) {
	request = &ModifyFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyFileSystem")
	return
}

func NewModifyFileSystemResponse() (response *ModifyFileSystemResponse) {
	response = &ModifyFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyFileSystem(request *ModifyFileSystemRequest) string {
	return c.ModifyFileSystemWithContext(context.Background(), request)
}

func (c *Client) ModifyFileSystemWithContext(ctx context.Context, request *ModifyFileSystemRequest) string {
	if request == nil {
		request = NewModifyFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateMountTargetRequest() (request *CreateMountTargetRequest) {
	request = &CreateMountTargetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateMountTarget")
	return
}

func NewCreateMountTargetResponse() (response *CreateMountTargetResponse) {
	response = &CreateMountTargetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMountTarget(request *CreateMountTargetRequest) string {
	return c.CreateMountTargetWithContext(context.Background(), request)
}

func (c *Client) CreateMountTargetWithContext(ctx context.Context, request *CreateMountTargetRequest) string {
	if request == nil {
		request = NewCreateMountTargetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMountTargetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteMountTargetRequest() (request *DeleteMountTargetRequest) {
	request = &DeleteMountTargetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteMountTarget")
	return
}

func NewDeleteMountTargetResponse() (response *DeleteMountTargetResponse) {
	response = &DeleteMountTargetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMountTarget(request *DeleteMountTargetRequest) string {
	return c.DeleteMountTargetWithContext(context.Background(), request)
}

func (c *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest) string {
	if request == nil {
		request = NewDeleteMountTargetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMountTargetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
	request = &DescribeMountTargetsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeMountTargets")
	return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
	response = &DescribeMountTargetsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) string {
	return c.DescribeMountTargetsWithContext(context.Background(), request)
}

func (c *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest) string {
	if request == nil {
		request = NewDescribeMountTargetsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMountTargetsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateModelRequest() (request *CreateModelRequest) {
	request = &CreateModelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateModel")
	return
}

func NewCreateModelResponse() (response *CreateModelResponse) {
	response = &CreateModelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateModel(request *CreateModelRequest) string {
	return c.CreateModelWithContext(context.Background(), request)
}

func (c *Client) CreateModelWithContext(ctx context.Context, request *CreateModelRequest) string {
	if request == nil {
		request = NewCreateModelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateModelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewTerminateModelsRequest() (request *TerminateModelsRequest) {
	request = &TerminateModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "TerminateModels")
	return
}

func NewTerminateModelsResponse() (response *TerminateModelsResponse) {
	response = &TerminateModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TerminateModels(request *TerminateModelsRequest) string {
	return c.TerminateModelsWithContext(context.Background(), request)
}

func (c *Client) TerminateModelsWithContext(ctx context.Context, request *TerminateModelsRequest) string {
	if request == nil {
		request = NewTerminateModelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTerminateModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeModelsRequest() (request *DescribeModelsRequest) {
	request = &DescribeModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeModels")
	return
}

func NewDescribeModelsResponse() (response *DescribeModelsResponse) {
	response = &DescribeModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModels(request *DescribeModelsRequest) string {
	return c.DescribeModelsWithContext(context.Background(), request)
}

func (c *Client) DescribeModelsWithContext(ctx context.Context, request *DescribeModelsRequest) string {
	if request == nil {
		request = NewDescribeModelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDedicatedClusterRequest() (request *DescribeDedicatedClusterRequest) {
	request = &DescribeDedicatedClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDedicatedCluster")
	return
}

func NewDescribeDedicatedClusterResponse() (response *DescribeDedicatedClusterResponse) {
	response = &DescribeDedicatedClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedCluster(request *DescribeDedicatedClusterRequest) string {
	return c.DescribeDedicatedClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedClusterWithContext(ctx context.Context, request *DescribeDedicatedClusterRequest) string {
	if request == nil {
		request = NewDescribeDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDedicatedClusterRequest() (request *CreateDedicatedClusterRequest) {
	request = &CreateDedicatedClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateDedicatedCluster")
	return
}

func NewCreateDedicatedClusterResponse() (response *CreateDedicatedClusterResponse) {
	response = &CreateDedicatedClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDedicatedCluster(request *CreateDedicatedClusterRequest) string {
	return c.CreateDedicatedClusterWithContext(context.Background(), request)
}

func (c *Client) CreateDedicatedClusterWithContext(ctx context.Context, request *CreateDedicatedClusterRequest) string {
	if request == nil {
		request = NewCreateDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDedicatedClusterRequest() (request *DeleteDedicatedClusterRequest) {
	request = &DeleteDedicatedClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteDedicatedCluster")
	return
}

func NewDeleteDedicatedClusterResponse() (response *DeleteDedicatedClusterResponse) {
	response = &DeleteDedicatedClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDedicatedCluster(request *DeleteDedicatedClusterRequest) string {
	return c.DeleteDedicatedClusterWithContext(context.Background(), request)
}

func (c *Client) DeleteDedicatedClusterWithContext(ctx context.Context, request *DeleteDedicatedClusterRequest) string {
	if request == nil {
		request = NewDeleteDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDedicatedClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSetvCPURequest() (request *SetvCPURequest) {
	request = &SetvCPURequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SetvCPU")
	return
}

func NewSetvCPUResponse() (response *SetvCPUResponse) {
	response = &SetvCPUResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetvCPU(request *SetvCPURequest) string {
	return c.SetvCPUWithContext(context.Background(), request)
}

func (c *Client) SetvCPUWithContext(ctx context.Context, request *SetvCPURequest) string {
	if request == nil {
		request = NewSetvCPURequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetvCPUResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDedicatedHostMigrateRequest() (request *DedicatedHostMigrateRequest) {
	request = &DedicatedHostMigrateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DedicatedHostMigrate")
	return
}

func NewDedicatedHostMigrateResponse() (response *DedicatedHostMigrateResponse) {
	response = &DedicatedHostMigrateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DedicatedHostMigrate(request *DedicatedHostMigrateRequest) string {
	return c.DedicatedHostMigrateWithContext(context.Background(), request)
}

func (c *Client) DedicatedHostMigrateWithContext(ctx context.Context, request *DedicatedHostMigrateRequest) string {
	if request == nil {
		request = NewDedicatedHostMigrateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDedicatedHostMigrateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDedicatedClusterNameRequest() (request *ModifyDedicatedClusterNameRequest) {
	request = &ModifyDedicatedClusterNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyDedicatedClusterName")
	return
}

func NewModifyDedicatedClusterNameResponse() (response *ModifyDedicatedClusterNameResponse) {
	response = &ModifyDedicatedClusterNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDedicatedClusterName(request *ModifyDedicatedClusterNameRequest) string {
	return c.ModifyDedicatedClusterNameWithContext(context.Background(), request)
}

func (c *Client) ModifyDedicatedClusterNameWithContext(ctx context.Context, request *ModifyDedicatedClusterNameRequest) string {
	if request == nil {
		request = NewModifyDedicatedClusterNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDedicatedClusterNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewInstanceMigrateRequest() (request *InstanceMigrateRequest) {
	request = &InstanceMigrateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "InstanceMigrate")
	return
}

func NewInstanceMigrateResponse() (response *InstanceMigrateResponse) {
	response = &InstanceMigrateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) InstanceMigrate(request *InstanceMigrateRequest) string {
	return c.InstanceMigrateWithContext(context.Background(), request)
}

func (c *Client) InstanceMigrateWithContext(ctx context.Context, request *InstanceMigrateRequest) string {
	if request == nil {
		request = NewInstanceMigrateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstanceMigrateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceAutoDeleteTimeRequest() (request *ModifyInstanceAutoDeleteTimeRequest) {
	request = &ModifyInstanceAutoDeleteTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceAutoDeleteTime")
	return
}

func NewModifyInstanceAutoDeleteTimeResponse() (response *ModifyInstanceAutoDeleteTimeResponse) {
	response = &ModifyInstanceAutoDeleteTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAutoDeleteTime(request *ModifyInstanceAutoDeleteTimeRequest) string {
	return c.ModifyInstanceAutoDeleteTimeWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAutoDeleteTimeWithContext(ctx context.Context, request *ModifyInstanceAutoDeleteTimeRequest) string {
	if request == nil {
		request = NewModifyInstanceAutoDeleteTimeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAutoDeleteTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyScalingConfigurationRequest() (request *ModifyScalingConfigurationRequest) {
	request = &ModifyScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingConfiguration")
	return
}

func NewModifyScalingConfigurationResponse() (response *ModifyScalingConfigurationResponse) {
	response = &ModifyScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingConfiguration(request *ModifyScalingConfigurationRequest) string {
	return c.ModifyScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingConfigurationWithContext(ctx context.Context, request *ModifyScalingConfigurationRequest) string {
	if request == nil {
		request = NewModifyScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeSpotPriceHistoryRequest() (request *DescribeSpotPriceHistoryRequest) {
	request = &DescribeSpotPriceHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeSpotPriceHistory")
	return
}

func NewDescribeSpotPriceHistoryResponse() (response *DescribeSpotPriceHistoryResponse) {
	response = &DescribeSpotPriceHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSpotPriceHistory(request *DescribeSpotPriceHistoryRequest) string {
	return c.DescribeSpotPriceHistoryWithContext(context.Background(), request)
}

func (c *Client) DescribeSpotPriceHistoryWithContext(ctx context.Context, request *DescribeSpotPriceHistoryRequest) string {
	if request == nil {
		request = NewDescribeSpotPriceHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSpotPriceHistoryResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrice(request *DescribePriceRequest) string {
	return c.DescribePriceWithContext(context.Background(), request)
}

func (c *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest) string {
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
func NewEnableImageCachingRequest() (request *EnableImageCachingRequest) {
	request = &EnableImageCachingRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "EnableImageCaching")
	return
}

func NewEnableImageCachingResponse() (response *EnableImageCachingResponse) {
	response = &EnableImageCachingResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableImageCaching(request *EnableImageCachingRequest) string {
	return c.EnableImageCachingWithContext(context.Background(), request)
}

func (c *Client) EnableImageCachingWithContext(ctx context.Context, request *EnableImageCachingRequest) string {
	if request == nil {
		request = NewEnableImageCachingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableImageCachingResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisableImageCachingRequest() (request *DisableImageCachingRequest) {
	request = &DisableImageCachingRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DisableImageCaching")
	return
}

func NewDisableImageCachingResponse() (response *DisableImageCachingResponse) {
	response = &DisableImageCachingResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableImageCaching(request *DisableImageCachingRequest) string {
	return c.DisableImageCachingWithContext(context.Background(), request)
}

func (c *Client) DisableImageCachingWithContext(ctx context.Context, request *DisableImageCachingRequest) string {
	if request == nil {
		request = NewDisableImageCachingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableImageCachingResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyLoadBalancersRequest() (request *ModifyLoadBalancersRequest) {
	request = &ModifyLoadBalancersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyLoadBalancers")
	return
}

func NewModifyLoadBalancersResponse() (response *ModifyLoadBalancersResponse) {
	response = &ModifyLoadBalancersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyLoadBalancers(request *ModifyLoadBalancersRequest) string {
	return c.ModifyLoadBalancersWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancersWithContext(ctx context.Context, request *ModifyLoadBalancersRequest) string {
	if request == nil {
		request = NewModifyLoadBalancersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAttachInstancesIamRoleRequest() (request *AttachInstancesIamRoleRequest) {
	request = &AttachInstancesIamRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachInstancesIamRole")
	return
}

func NewAttachInstancesIamRoleResponse() (response *AttachInstancesIamRoleResponse) {
	response = &AttachInstancesIamRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachInstancesIamRole(request *AttachInstancesIamRoleRequest) string {
	return c.AttachInstancesIamRoleWithContext(context.Background(), request)
}

func (c *Client) AttachInstancesIamRoleWithContext(ctx context.Context, request *AttachInstancesIamRoleRequest) string {
	if request == nil {
		request = NewAttachInstancesIamRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachInstancesIamRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachInstancesIamRoleRequest() (request *DetachInstancesIamRoleRequest) {
	request = &DetachInstancesIamRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachInstancesIamRole")
	return
}

func NewDetachInstancesIamRoleResponse() (response *DetachInstancesIamRoleResponse) {
	response = &DetachInstancesIamRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachInstancesIamRole(request *DetachInstancesIamRoleRequest) string {
	return c.DetachInstancesIamRoleWithContext(context.Background(), request)
}

func (c *Client) DetachInstancesIamRoleWithContext(ctx context.Context, request *DetachInstancesIamRoleRequest) string {
	if request == nil {
		request = NewDetachInstancesIamRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachInstancesIamRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCopySnapshotRequest() (request *CopySnapshotRequest) {
	request = &CopySnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CopySnapshot")
	return
}

func NewCopySnapshotResponse() (response *CopySnapshotResponse) {
	response = &CopySnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CopySnapshot(request *CopySnapshotRequest) string {
	return c.CopySnapshotWithContext(context.Background(), request)
}

func (c *Client) CopySnapshotWithContext(ctx context.Context, request *CopySnapshotRequest) string {
	if request == nil {
		request = NewCopySnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopySnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewPreMigrateInstanceRequest() (request *PreMigrateInstanceRequest) {
	request = &PreMigrateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "PreMigrateInstance")
	return
}

func NewPreMigrateInstanceResponse() (response *PreMigrateInstanceResponse) {
	response = &PreMigrateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) PreMigrateInstance(request *PreMigrateInstanceRequest) string {
	return c.PreMigrateInstanceWithContext(context.Background(), request)
}

func (c *Client) PreMigrateInstanceWithContext(ctx context.Context, request *PreMigrateInstanceRequest) string {
	if request == nil {
		request = NewPreMigrateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPreMigrateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCancelPreMigrateInstanceRequest() (request *CancelPreMigrateInstanceRequest) {
	request = &CancelPreMigrateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CancelPreMigrateInstance")
	return
}

func NewCancelPreMigrateInstanceResponse() (response *CancelPreMigrateInstanceResponse) {
	response = &CancelPreMigrateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelPreMigrateInstance(request *CancelPreMigrateInstanceRequest) string {
	return c.CancelPreMigrateInstanceWithContext(context.Background(), request)
}

func (c *Client) CancelPreMigrateInstanceWithContext(ctx context.Context, request *CancelPreMigrateInstanceRequest) string {
	if request == nil {
		request = NewCancelPreMigrateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelPreMigrateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetVNCAddressRequest() (request *GetVNCAddressRequest) {
	request = &GetVNCAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "GetVNCAddress")
	return
}

func NewGetVNCAddressResponse() (response *GetVNCAddressResponse) {
	response = &GetVNCAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetVNCAddress(request *GetVNCAddressRequest) string {
	return c.GetVNCAddressWithContext(context.Background(), request)
}

func (c *Client) GetVNCAddressWithContext(ctx context.Context, request *GetVNCAddressRequest) string {
	if request == nil {
		request = NewGetVNCAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetVNCAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSwitchImageTypeRequest() (request *SwitchImageTypeRequest) {
	request = &SwitchImageTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SwitchImageType")
	return
}

func NewSwitchImageTypeResponse() (response *SwitchImageTypeResponse) {
	response = &SwitchImageTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SwitchImageType(request *SwitchImageTypeRequest) string {
	return c.SwitchImageTypeWithContext(context.Background(), request)
}

func (c *Client) SwitchImageTypeWithContext(ctx context.Context, request *SwitchImageTypeRequest) string {
	if request == nil {
		request = NewSwitchImageTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSwitchImageTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
