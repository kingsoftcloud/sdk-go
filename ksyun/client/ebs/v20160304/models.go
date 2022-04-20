package v20160304
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeName *string `json:"VolumeName,omitempty" name:"VolumeName"`
    VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`
    VolumeDesc *string `json:"VolumeDesc,omitempty" name:"VolumeDesc"`
    Size *int `json:"Size,omitempty" name:"Size"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    PurchaseTime *int `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    SubOrderId *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
}

func (r *CreateVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    VolumeId *string `json:"VolumeId" name:"VolumeId"`
}

func (r *CreateVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    DeleteWithInstance *string `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

func (r *AttachVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AttachVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *string `json:"Return" name:"Return"`
}

func (r *AttachVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachVolumeRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *DetachVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DetachVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *DetachVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *string `json:"Return" name:"Return"`
}

func (r *DeleteVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    Size *string `json:"Size,omitempty" name:"Size"`
    OnlineResize *bool `json:"OnlineResize,omitempty" name:"OnlineResize"`
    SubOrderId *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
}

func (r *ResizeVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResizeVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ResizeVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *ResizeVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVolumesRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId []*string `json:"VolumeId,omitempty" name:"VolumeId"`
    VolumeCategory *string `json:"VolumeCategory,omitempty" name:"VolumeCategory"`
    VolumeStatus *string `json:"VolumeStatus,omitempty" name:"VolumeStatus"`
    VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`
    VolumeCreateDate *string `json:"VolumeCreateDate,omitempty" name:"VolumeCreateDate"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeVolumesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVolumesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVolumesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVolumesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Volumes []struct {
		VolumeId *string `json:"VolumeId"`
		VolumeName *string `json:"VolumeName"`
		VolumeDesc *string `json:"VolumeDesc"`
		Size *int `json:"Size"`
		VolumeStatus *string `json:"VolumeStatus"`
		VolumeType *string `json:"VolumeType"`
		VolumeCategory *string `json:"VolumeCategory"`
		InstanceId *string `json:"InstanceId"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		ChargeType *string `json:"ChargeType"`
		InstanceTradeType *int `json:"InstanceTradeType"`
		CreateTime *string `json:"CreateTime"`
		Attachment []struct {
					InstanceId *string `json:"InstanceId"`
					MountPoint *string `json:"MountPoint"`
					DeleteWithInstance *bool `json:"DeleteWithInstance"`
			} `json:"Attachment"`
			ProjectId *string `json:"ProjectId"`
			ExpireTime *string `json:"ExpireTime"`
		} `json:"Volumes"`
}

func (r *DescribeVolumesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVolumesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    VolumeName *string `json:"VolumeName,omitempty" name:"VolumeName"`
    VolumeDesc *string `json:"VolumeDesc,omitempty" name:"VolumeDesc"`
    DeleteWithInstance *string `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

func (r *ModifyVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *ModifyVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEbsInstancesRequest struct {
    *ksyunhttp.BaseRequest
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`
}

func (r *DescribeEbsInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEbsInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEbsInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEbsInstancesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Instances []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		InstanceIp *string `json:"InstanceIp"`
		InstanceEnable *string `json:"InstanceEnable"`
	} `json:"Instances"`
}

func (r *DescribeEbsInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEbsInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVolumesRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVolumesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceVolumesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceVolumesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVolumesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Attachments []struct {
		InstanceId *string `json:"InstanceId"`
		VolumeId *string `json:"VolumeId"`
		MountPoint *string `json:"MountPoint"`
	} `json:"Attachments"`
}

func (r *DescribeInstanceVolumesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceVolumesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    PurchaseTime *int `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *RenewVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenewVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RenewVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    PaidMoney *int `json:"PaidMoney" name:"PaidMoney"`
}

func (r *RenewVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateVolumeProjectRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId []*string `json:"VolumeId,omitempty" name:"VolumeId"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *UpdateVolumeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateVolumeProjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateVolumeProjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateVolumeProjectResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *UpdateVolumeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateVolumeProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    VolumeCategory *string `json:"VolumeCategory,omitempty" name:"VolumeCategory"`
    SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
    PageNumber *int `json:"PageNumber,omitempty" name:"PageNumber"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Snapshots []struct {
		SnapshotId *string `json:"SnapshotId"`
		SnapshotName *string `json:"SnapshotName"`
		VolumeId *string `json:"VolumeId"`
		Size *int `json:"Size"`
		CreateTime *string `json:"CreateTime"`
		SnapshotStatus *string `json:"SnapshotStatus"`
		VolumeCategory *string `json:"VolumeCategory"`
		VolumeName *string `json:"VolumeName"`
		VolumeType *string `json:"VolumeType"`
		Progress *string `json:"Progress"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		VolumeStatus *string `json:"VolumeStatus"`
		SnapshotType *string `json:"SnapshotType"`
	} `json:"Snapshots"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
    SnapshotDesc *string `json:"SnapshotDesc,omitempty" name:"SnapshotDesc"`
    SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSnapshotRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSnapshotRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSnapshotResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotRequest struct {
    *ksyunhttp.BaseRequest
    SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DeleteSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSnapshotRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSnapshotRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *string `json:"Return" name:"Return"`
}

func (r *DeleteSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSnapshotResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackSnapshotRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
    SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *RollbackSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackSnapshotRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RollbackSnapshotRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RollbackSnapshotResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *RollbackSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackSnapshotResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotRequest struct {
    *ksyunhttp.BaseRequest
    SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
    SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
    SnapshotDesc *string `json:"SnapshotDesc,omitempty" name:"SnapshotDesc"`
}

func (r *ModifySnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySnapshotRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySnapshotRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *ModifySnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySnapshotResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecoveryVolumeRequest struct {
    *ksyunhttp.BaseRequest
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *RecoveryVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoveryVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RecoveryVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RecoveryVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *string `json:"Return" name:"Return"`
}

func (r *RecoveryVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoveryVolumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateAttachInstanceRequest struct {
    *ksyunhttp.BaseRequest
    VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ValidateAttachInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateAttachInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ValidateAttachInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ValidateAttachInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    InstanceEnable *string `json:"InstanceEnable" name:"InstanceEnable"`
}

func (r *ValidateAttachInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateAttachInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

