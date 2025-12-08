package v20160304
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateVolumeTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}


type CreateVolumeRequest struct {
	*ksyunhttp.BaseRequest
	VolumeName       *string            `json:"VolumeName,omitempty" name:"VolumeName"`
	VolumeType       *string            `json:"VolumeType,omitempty" name:"VolumeType"`
	VolumeDesc       *string            `json:"VolumeDesc,omitempty" name:"VolumeDesc"`
	Size             *int               `json:"Size,omitempty" name:"Size"`
	AvailabilityZone *string            `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ChargeType       *string            `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime     *int               `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	ProjectId        *string            `json:"ProjectId,omitempty" name:"ProjectId"`
	SubOrderId       *string            `json:"SubOrderId,omitempty" name:"SubOrderId"`
	SnapshotId       *string            `json:"SnapshotId,omitempty" name:"SnapshotId"`
	ClusterId        *string            `json:"ClusterId,omitempty" name:"ClusterId"`
	Tag              []*CreateVolumeTag `json:"Tag,omitempty" name:"Tag"`
	ProvisionedIops  *int               `json:"ProvisionedIops,omitempty" name:"ProvisionedIops"`
}

func (r *CreateVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	VolumeId  *string `json:"VolumeId" name:"VolumeId"`
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
}

func (r *AttachVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AttachVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *string `json:"Return" name:"Return"`
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
	VolumeId   *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *DetachVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DetachVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
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
}

func (r *DeleteVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *string `json:"Return" name:"Return"`
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
	VolumeId     *string `json:"VolumeId,omitempty" name:"VolumeId"`
	Size         *string `json:"Size,omitempty" name:"Size"`
	OnlineResize *bool   `json:"OnlineResize,omitempty" name:"OnlineResize"`
	SubOrderId   *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
}

func (r *ResizeVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResizeVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
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
	VolumeId              []*string `json:"VolumeId,omitempty" name:"VolumeId"`
	VolumeCategory        *string   `json:"VolumeCategory,omitempty" name:"VolumeCategory"`
	VolumeStatus          *string   `json:"VolumeStatus,omitempty" name:"VolumeStatus"`
	VolumeType            *string   `json:"VolumeType,omitempty" name:"VolumeType"`
	VolumeCreateDate      *string   `json:"VolumeCreateDate,omitempty" name:"VolumeCreateDate"`
	Marker                *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults            *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	TagNKey               *string   `json:"Tag.N.Key,omitempty" name:"Tag.N.Key"`
	TagNValue             *string   `json:"Tag.N.Value,omitempty" name:"Tag.N.Value"`
	VolumeCreateEndDate   *string   `json:"VolumeCreateEndDate,omitempty" name:"VolumeCreateEndDate"`
	VolumeCreateStartDate *string   `json:"VolumeCreateStartDate,omitempty" name:"VolumeCreateStartDate"`
	SourceSnapshotId      *string   `json:"SourceSnapshotId,omitempty" name:"SourceSnapshotId"`
}

func (r *DescribeVolumesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeVolumesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Volumes   []struct {
		VolumeId          *string `json:"VolumeId" name:"VolumeId"`
		VolumeName        *string `json:"VolumeName" name:"VolumeName"`
		VolumeDesc        *string `json:"VolumeDesc" name:"VolumeDesc"`
		Size              *int    `json:"Size" name:"Size"`
		VolumeStatus      *string `json:"VolumeStatus" name:"VolumeStatus"`
		VolumeType        *string `json:"VolumeType" name:"VolumeType"`
		VolumeCategory    *string `json:"VolumeCategory" name:"VolumeCategory"`
		InstanceId        *string `json:"InstanceId" name:"InstanceId"`
		AvailabilityZone  *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		ChargeType        *string `json:"ChargeType" name:"ChargeType"`
		InstanceTradeType *int    `json:"InstanceTradeType" name:"InstanceTradeType"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		Attachment        []struct {
			InstanceId         *string `json:"InstanceId" name:"InstanceId"`
			MountPoint         *string `json:"MountPoint" name:"MountPoint"`
			DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
		} `json:"Attachment" name:"Attachment"`
		ProjectId *int `json:"ProjectId" name:"ProjectId"`
		ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`
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
}

func (r *ModifyVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
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
	VolumeType       *string `json:"VolumeType,omitempty" name:"VolumeType"`
}

func (r *DescribeEbsInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEbsInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Instances []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceIp   *string `json:"InstanceIp" name:"InstanceIp"`
		InstanceEnable *string `json:"InstanceEnable" name:"InstanceEnable"`
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
}

func (r *DescribeInstanceVolumesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceVolumesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Attachments []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		VolumeId *string `json:"VolumeId" name:"VolumeId"`
		MountPoint *string `json:"MountPoint" name:"MountPoint"`
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
}

func (r *RenewVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenewVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	PaidMoney *int    `json:"PaidMoney" name:"PaidMoney"`
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
	VolumeId  []*string `json:"VolumeId,omitempty" name:"VolumeId"`
	ProjectId *string   `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *UpdateVolumeProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateVolumeProjectResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
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
	VolumeId         *string `json:"VolumeId,omitempty" name:"VolumeId"`
	VolumeCategory   *string `json:"VolumeCategory,omitempty" name:"VolumeCategory"`
	SnapshotId       *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	SnapshotName     *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	PageNumber       *int    `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize         *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSnapshotsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Snapshots []struct {
		SnapshotId     *string `json:"SnapshotId" name:"SnapshotId"`
		SnapshotName   *string `json:"SnapshotName" name:"SnapshotName"`
		VolumeId       *string `json:"VolumeId" name:"VolumeId"`
		Size           *int    `json:"Size" name:"Size"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		SnapshotStatus *string `json:"SnapshotStatus" name:"SnapshotStatus"`
		VolumeCategory *string `json:"VolumeCategory" name:"VolumeCategory"`
		VolumeName     *string `json:"VolumeName" name:"VolumeName"`
		VolumeType     *string `json:"VolumeType" name:"VolumeType"`
		Progress       *string `json:"Progress" name:"Progress"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		VolumeStatus   *string `json:"VolumeStatus" name:"VolumeStatus"`
		SnapshotType   *string `json:"SnapshotType" name:"SnapshotType"`
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
	VolumeId            *string `json:"VolumeId,omitempty" name:"VolumeId"`
	SnapshotName        *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	SnapshotDesc        *string `json:"SnapshotDesc,omitempty" name:"SnapshotDesc"`
	SnapshotType        *string `json:"SnapshotType,omitempty" name:"SnapshotType"`
	ScheduledDeleteTime *string `json:"ScheduledDeleteTime,omitempty" name:"ScheduledDeleteTime"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
}

func (r *DeleteSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *string `json:"Return" name:"Return"`
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
}

func (r *RollbackSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RollbackSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
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
}

func (r *ModifySnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
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
}

func (r *RecoveryVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RecoveryVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *string `json:"Return" name:"Return"`
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

type ValidateAttachInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	InstanceEnable *string `json:"InstanceEnable" name:"InstanceEnable"`
}

func (r *ValidateAttachInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateAttachInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailabilityZonesRequest struct {
	*ksyunhttp.BaseRequest
	VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`
}

func (r *DescribeAvailabilityZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAvailabilityZonesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string   `json:"RequestId" name:"RequestId"`
	AvailabilityZones []*string `json:"AvailabilityZones" name:"AvailabilityZones"`
}

func (r *DescribeAvailabilityZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeCreateVolumePriceRequest struct {
	*ksyunhttp.BaseRequest
	VolumeType       *string `json:"VolumeType,omitempty" name:"VolumeType"`
	Size             *int    `json:"Size,omitempty" name:"Size"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ChargeType       *string `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime     *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *DescribeCreateVolumePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCreateVolumePriceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	VolumePrice struct {
		CurrencyUnit *string  `json:"CurrencyUnit" name:"CurrencyUnit"`
		TradePrice   *float64 `json:"TradePrice" name:"TradePrice"`
		OriginalPrice *float64 `json:"OriginalPrice" name:"OriginalPrice"`
		DiscountPrice *float64 `json:"DiscountPrice" name:"DiscountPrice"`
	} `json:"VolumePrice"`
}

func (r *DescribeCreateVolumePriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCreateVolumePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifySnapshotTypeRequest struct {
	*ksyunhttp.BaseRequest
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	SnapshotId  *string   `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *ModifySnapshotTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySnapshotTypeResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySnapshotTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyVolumeTypeRequest struct {
	*ksyunhttp.BaseRequest
	VolumeId                       *string `json:"VolumeId,omitempty" name:"VolumeId"`
	PerformanceVolumeSize          *string `json:"PerformanceVolumeSize,omitempty" name:"PerformanceVolumeSize"`
	PerformanceLevelVolumeCategory *string `json:"PerformanceLevelVolumeCategory,omitempty" name:"PerformanceLevelVolumeCategory"`
}

func (r *ModifyVolumeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyVolumeTypeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyVolumeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVolumeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyDedicatedBlockStorageClusterAttributeRequest struct {
	*ksyunhttp.BaseRequest
	DbscId           *string `json:"DbscId,omitempty" name:"DbscId"`
	DbscName         *string `json:"DbscName,omitempty" name:"DbscName"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	DbscDesc         *string `json:"DbscDesc,omitempty" name:"DbscDesc"`
}

func (r *ModifyDedicatedBlockStorageClusterAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDedicatedBlockStorageClusterAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	DbscId    *string `json:"DbscId" name:"DbscId"`
}

func (r *ModifyDedicatedBlockStorageClusterAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDedicatedBlockStorageClusterAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ResizeDedicatedBlockStorageClustersRequest struct {
	*ksyunhttp.BaseRequest
	DbscId           *string `json:"DbscId,omitempty" name:"DbscId"`
	Size             *int    `json:"Size,omitempty" name:"Size"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
}

func (r *ResizeDedicatedBlockStorageClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResizeDedicatedBlockStorageClustersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ResizeDedicatedBlockStorageClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeDedicatedBlockStorageClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedBlockStorageClustersRequest struct {
	*ksyunhttp.BaseRequest
	AvailabilityZone *string   `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	DbscName         *string   `json:"DbscName,omitempty" name:"DbscName"`
	Marker           *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults       *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	DbscCreateDate   *string   `json:"DbscCreateDate,omitempty" name:"DbscCreateDate"`
	DbscId           []*string `json:"DbscId,omitempty" name:"DbscId"`
}

func (r *DescribeDedicatedBlockStorageClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDedicatedBlockStorageClustersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Dbscs     []struct {
		DscId                            *string `json:"DscId" name:"DscId"`
		DbscName                         *string `json:"DbscName" name:"DbscName"`
		DbscDescription                  *string `json:"DbscDescription" name:"DbscDescription"`
		AvailabilityZone                 *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		DbscStatus                       *string `json:"DbscStatus" name:"DbscStatus"`
		DbscType                         *string `json:"DbscType" name:"DbscType"`
		EbsClusterId                     *string `json:"EbsClusterId" name:"EbsClusterId"`
		CreateTime                       *string `json:"CreateTime" name:"CreateTime"`
		DedicatedBlockStorageClusterSize struct {
			AvailableSize *int `json:"AvailableSize" name:"AvailableSize"`
			TotalSize     *int `json:"TotalSize" name:"TotalSize"`
			DeliverySize  *int `json:"DeliverySize" name:"DeliverySize"`
			UsedSize      *int `json:"UsedSize" name:"UsedSize"`
		} `json:"DedicatedBlockStorageClusterSize" name:"DedicatedBlockStorageClusterSize"`
	} `json:"Dbscs"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
	Marker     *int `json:"Marker" name:"Marker"`
	MaxResults *int `json:"MaxResults" name:"MaxResults"`
}

func (r *DescribeDedicatedBlockStorageClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedBlockStorageClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDedicatedBlockStorageClusterRequest struct {
	*ksyunhttp.BaseRequest
	DbscName         *string `json:"DbscName,omitempty" name:"DbscName"`
	DbscType         *string `json:"DbscType,omitempty" name:"DbscType"`
	Size             *int    `json:"Size,omitempty" name:"Size"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	PurchaseTime     *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *CreateDedicatedBlockStorageClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDedicatedBlockStorageClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	DbscId    *string `json:"DbscId" name:"DbscId"`
}

func (r *CreateDedicatedBlockStorageClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedBlockStorageClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVolumePresetRequest struct {
	*ksyunhttp.BaseRequest
	VolumeId        *string `json:"VolumeId,omitempty" name:"VolumeId"`
	ProvisionedIops *int    `json:"ProvisionedIops,omitempty" name:"ProvisionedIops"`
}

func (r *ModifyVolumePresetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyVolumePresetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
}

func (r *ModifyVolumePresetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVolumePresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeVolumeTypeProcessInfoRequest struct {
	*ksyunhttp.BaseRequest
	VolumeId []*string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *GetUpgradeVolumeTypeProcessInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetUpgradeVolumeTypeProcessInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    []struct {
		VolumeId *string `json:"VolumeId" name:"VolumeId"`
		Process  struct {
			PrepareProcess    *int  `json:"PrepareProcess" name:"PrepareProcess"`
			LatestDataProcess *int  `json:"LatestDataProcess" name:"LatestDataProcess"`
			PrepareDone       *bool `json:"PrepareDone" name:"PrepareDone"`
			LatestDataDone    *bool `json:"LatestDataDone" name:"LatestDataDone"`
			EstimatedDuring   *int  `json:"EstimatedDuring" name:"EstimatedDuring"`
		} `json:"Process" name:"Process"`
	} `json:"Result"`
}

func (r *GetUpgradeVolumeTypeProcessInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeVolumeTypeProcessInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

