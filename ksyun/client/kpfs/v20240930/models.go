package v20240930

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type UpdatePerformanceOnePosixAclFileSystemList struct {
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	VolumePath   *string `json:"VolumePath,omitempty" name:"VolumePath"`
}
type SetPerformanceOnePosixAclFileSystemList struct {
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	VolumePath   *string `json:"VolumePath,omitempty" name:"VolumePath"`
}
type UpdateDataMigrateTaskExportTaskPeriodConfig struct {
	FrequencyUnit    *string   `json:"FrequencyUnit,omitempty" name:"FrequencyUnit"`
	IndexOfFrequency []*int    `json:"IndexOfFrequency,omitempty" name:"IndexOfFrequency"`
	TimePoints       []*string `json:"TimePoints,omitempty" name:"TimePoints"`
}
type CreateDataMigrateTaskBucketConfig struct {
	Bucket       *string `json:"Bucket,omitempty" name:"Bucket"`
	BucektPrefix *string `json:"BucektPrefix,omitempty" name:"BucektPrefix"`
}
type CreateDataMigrateTaskExportTaskPeriodConfig struct {
	FrequencyUnit    *string   `json:"FrequencyUnit,omitempty" name:"FrequencyUnit"`
	IndexOfFrequency []*int    `json:"IndexOfFrequency,omitempty" name:"IndexOfFrequency"`
	TimePoints       []*string `json:"TimePoints,omitempty" name:"TimePoints"`
}
type DeleteRecycledFilesFiles struct {
	FileName   *string `json:"FileName,omitempty" name:"FileName"`
	DeleteTime *int64  `json:"DeleteTime,omitempty" name:"DeleteTime"`
	Position   *string `json:"Position,omitempty" name:"Position"`
}
type RestoreRecycledFilesFiles struct {
	FileName   *string `json:"FileName,omitempty" name:"FileName"`
	DeleteTime *int64  `json:"DeleteTime,omitempty" name:"DeleteTime"`
	Position   *string `json:"Position,omitempty" name:"Position"`
	Type       *string `json:"Type,omitempty" name:"Type"`
	Length     *int64  `json:"Length,omitempty" name:"Length"`
}
type UpdatePerformanceNfsAclIpIps struct {
	Ip         *string `json:"Ip,omitempty" name:"Ip"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
	RootSquash *string `json:"RootSquash,omitempty" name:"RootSquash"`
}
type AddPerformanceNfsAclClientIps struct {
	Ip         *string `json:"Ip,omitempty" name:"Ip"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
	RootSquash *string `json:"RootSquash,omitempty" name:"RootSquash"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}
type SetPerformanceOneNfsAclIps struct {
	Ip         *string `json:"Ip,omitempty" name:"Ip"`
	Permission *string `json:"Permission,omitempty" name:"Permission"`
	RootSquash *string `json:"RootSquash,omitempty" name:"RootSquash"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}
type UpdateFileDeletePolicyFrequencyTimePointsStart struct {
	Hour *int `json:"Hour,omitempty" name:"Hour"`
}
type UpdateFileDeletePolicyFrequencyTimePointsEnd struct {
	Hour *int `json:"Hour,omitempty" name:"Hour"`
}
type UpdateFileDeletePolicyFrequencyTimePoints struct {
	Start *UpdateFileDeletePolicyFrequencyTimePointsStart `json:"Start,omitempty" name:"Start"`
	End   *UpdateFileDeletePolicyFrequencyTimePointsEnd   `json:"End,omitempty" name:"End"`
}
type UpdateFileDeletePolicyFileNameRule struct {
	Rule *string `json:"Rule,omitempty" name:"Rule"`
}
type UpdateFileDeletePolicyFileSizeRule struct {
	Rule     *string `json:"Rule,omitempty" name:"Rule"`
	MaxValue *int    `json:"MaxValue,omitempty" name:"MaxValue"`
	Unit     *string `json:"Unit,omitempty" name:"Unit"`
	MinValue *int    `json:"MinValue,omitempty" name:"MinValue"`
}
type UpdateFileDeletePolicyTimeRuleParameters struct {
	Type   *string `json:"Type,omitempty" name:"Type"`
	OpType *string `json:"OpType,omitempty" name:"OpType"`
	Unit   *string `json:"Unit,omitempty" name:"Unit"`
	Value  *int    `json:"Value,omitempty" name:"Value"`
}
type CreateFileDeletePolicyFrequencyTimePointsStart struct {
	Hour *int `json:"Hour,omitempty" name:"Hour"`
}
type CreateFileDeletePolicyFrequencyTimePointsEnd struct {
	Hour *int `json:"Hour,omitempty" name:"Hour"`
}
type CreateFileDeletePolicyFrequencyTimePoints struct {
	Start *CreateFileDeletePolicyFrequencyTimePointsStart `json:"Start,omitempty" name:"Start"`
	End   *CreateFileDeletePolicyFrequencyTimePointsEnd   `json:"End,omitempty" name:"End"`
}
type CreateFileDeletePolicyFileNameRule struct {
	Rule *string `json:"Rule,omitempty" name:"Rule"`
}
type CreateFileDeletePolicyFileSizeRule struct {
	Rule     *string `json:"Rule,omitempty" name:"Rule"`
	MaxValue *int    `json:"MaxValue,omitempty" name:"MaxValue"`
	MinValue *int    `json:"MinValue,omitempty" name:"MinValue"`
	Unit     *string `json:"Unit,omitempty" name:"Unit"`
}
type CreateFileDeletePolicyTimeRuleParameters struct {
	Type   *string `json:"Type,omitempty" name:"Type"`
	OpType *string `json:"OpType,omitempty" name:"OpType"`
	Unit   *string `json:"Unit,omitempty" name:"Unit"`
	Value  *int    `json:"Value,omitempty" name:"Value"`
}
type CreateMigrateRuleSrcData struct {
	StorageType  *string `json:"StorageType,omitempty" name:"StorageType"`
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DirPath      *string `json:"DirPath,omitempty" name:"DirPath"`
	BucketName   *string `json:"BucketName,omitempty" name:"BucketName"`
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`
}
type CreateMigrateRuleDstData struct {
	StorageType  *string `json:"StorageType,omitempty" name:"StorageType"`
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DirPath      *string `json:"DirPath,omitempty" name:"DirPath"`
	BucketName   *string `json:"BucketName,omitempty" name:"BucketName"`
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`
}

type DescribeFileSystemListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	FileSystemIds  *string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
	StoreClasses   *string `json:"StoreClasses,omitempty" name:"StoreClasses"`
	ProjectId      *string `json:"ProjectId,omitempty" name:"ProjectId"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFileSystemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFileSystemListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Status     *string `json:"Status" name:"Status"`
		ChargeInfo struct {
			ChargeType  *string `json:"ChargeType" name:"ChargeType"`
			ExpiredTime *string `json:"ExpiredTime" name:"ExpiredTime"`
		} `json:"ChargeInfo" name:"ChargeInfo"`
		FileSystemInfo struct {
			FileSystemName     *string `json:"FileSystemName" name:"FileSystemName"`
			Capacity           *int64  `json:"Capacity" name:"Capacity"`
			Region             *string `json:"Region" name:"Region"`
			RegionName         *string `json:"RegionName" name:"RegionName"`
			AvailZone          *string `json:"AvailZone" name:"AvailZone"`
			FileSystemId       *string `json:"FileSystemId" name:"FileSystemId"`
			CreateTime         *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime         *string `json:"UpdateTime" name:"UpdateTime"`
			StoreClass         *string `json:"StoreClass" name:"StoreClass"`
			StorePoolType      *string `json:"StorePoolType" name:"StorePoolType"`
			ClientMountCommand *string `json:"ClientMountCommand" name:"ClientMountCommand"`
			ChunkSize          *string `json:"ChunkSize" name:"ChunkSize"`
			SRoceCluster       *string `json:"SRoceCluster" name:"SRoceCluster"`
			ClusterName        *string `json:"ClusterName" name:"ClusterName"`
			ClusterCode        *string `json:"ClusterCode" name:"ClusterCode"`
		} `json:"FileSystemInfo" name:"FileSystemInfo"`
		AccessRules []struct {
			Token *string `json:"Token" name:"Token"`
		} `json:"AccessRules" name:"AccessRules"`
		ProjectId  *string `json:"ProjectId" name:"ProjectId"`
		VolumeInfo struct {
			Inodes      *int64 `json:"Inodes" name:"Inodes"`
			UseCapacity *int64 `json:"UseCapacity" name:"UseCapacity"`
		} `json:"VolumeInfo" name:"VolumeInfo"`
	} `json:"Data"`
	PageSize   *int64 `json:"PageSize" name:"PageSize"`
	PageNum    *int64 `json:"PageNum" name:"PageNum"`
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeFileSystemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSystemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTotalSizeRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	DirPath      *string `json:"DirPath,omitempty" name:"DirPath"`
}

func (r *GetTotalSizeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetTotalSizeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetTotalSizeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTotalSizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInodeCountRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	DirPath      *string `json:"DirPath,omitempty" name:"DirPath"`
}

func (r *GetInodeCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetInodeCountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetInodeCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInodeCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemClientInfoRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" name:"HostNamePrefix"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeFileSystemClientInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFileSystemClientInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		HostName       *string `json:"HostName" name:"HostName"`
		Ip             *string `json:"Ip" name:"Ip"`
		VpcIp          *string `json:"VpcIp" name:"VpcIp"`
		MountPoint     *string `json:"MountPoint" name:"MountPoint"`
		Version        *string `json:"Version" name:"Version"`
		FileHandles    *int    `json:"FileHandles" name:"FileHandles"`
		RunningTime    *int64  `json:"RunningTime" name:"RunningTime"`
		CmdArgs        *string `json:"CmdArgs" name:"CmdArgs"`
		CacheGroup     *string `json:"CacheGroup" name:"CacheGroup"`
		CacheGroupRole *string `json:"CacheGroupRole" name:"CacheGroupRole"`
		ClientId       *string `json:"ClientId" name:"ClientId"`
	} `json:"Data"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeFileSystemClientInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSystemClientInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemFileListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	Dir          *string `json:"Dir,omitempty" name:"Dir"`
	FileName     *string `json:"FileName,omitempty" name:"FileName"`
	PageNum      *int64  `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int64  `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFileSystemFileListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFileSystemFileListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Name       *string `json:"Name" name:"Name"`
		Type       *string `json:"Type" name:"Type"`
		Length     *int64  `json:"Length" name:"Length"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
	PageSize   *int64 `json:"PageSize" name:"PageSize"`
	PageNum    *int64 `json:"PageNum" name:"PageNum"`
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeFileSystemFileListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSystemFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	PurchaseTime *int64  `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *RenewFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenewFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RenewFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	Capacity     *int64  `json:"Capacity,omitempty" name:"Capacity"`
}

func (r *UpgradeFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpgradeFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpgradeFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	Region         *string `json:"Region,omitempty" name:"Region"`
	AvailZone      *string `json:"AvailZone,omitempty" name:"AvailZone"`
	ChargeType     *string `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime   *int64  `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	StoreClass     *string `json:"StoreClass,omitempty" name:"StoreClass"`
	Capacity       *int64  `json:"Capacity,omitempty" name:"Capacity"`
	ChunkSize      *int64  `json:"ChunkSize,omitempty" name:"ChunkSize"`
	ClusterCode    *string `json:"ClusterCode,omitempty" name:"ClusterCode"`
}

func (r *CreateFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	FileSystemId *string `json:"FileSystemId" name:"FileSystemId"`
	RequestId    *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCapacityAvailableRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *GetCapacityAvailableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetCapacityAvailableResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetCapacityAvailableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCapacityAvailableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCapacityTotalRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *GetCapacityTotalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetCapacityTotalResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetCapacityTotalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCapacityTotalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLatencyWriteRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	ClientNm     *string `json:"ClientNm,omitempty" name:"ClientNm"`
	VpcIp        *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *GetLatencyWriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetLatencyWriteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetLatencyWriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLatencyWriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLatencyReadRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	ClientNm     *string `json:"ClientNm,omitempty" name:"ClientNm"`
	VpcIp        *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *GetLatencyReadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetLatencyReadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetLatencyReadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLatencyReadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetIopsWriteRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	ClientNm     *string `json:"ClientNm,omitempty" name:"ClientNm"`
	VpcIp        *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *GetIopsWriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetIopsWriteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetIopsWriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetIopsWriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetIopsReadRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	ClientNm     *string `json:"ClientNm,omitempty" name:"ClientNm"`
	VpcIp        *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *GetIopsReadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetIopsReadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetIopsReadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetIopsReadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBandwidthWriteRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	ClientNm     *string `json:"ClientNm,omitempty" name:"ClientNm"`
	VpcIp        *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *GetBandwidthWriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBandwidthWriteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetBandwidthWriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBandwidthWriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBandwidthReadRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval     *string `json:"Interval,omitempty" name:"Interval"`
	ClientNm     *string `json:"ClientNm,omitempty" name:"ClientNm"`
	VpcIp        *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *GetBandwidthReadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBandwidthReadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64 `json:"Time" name:"Time"`
		Value *int64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetBandwidthReadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBandwidthReadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePerformanceOnePosixAclRequest struct {
	*ksyunhttp.BaseRequest
	PosixAclId *string `json:"PosixAclId,omitempty" name:"PosixAclId"`
}

func (r *DeletePerformanceOnePosixAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeletePerformanceOnePosixAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeletePerformanceOnePosixAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePerformanceOnePosixAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePerformanceOnePosixAclRequest struct {
	*ksyunhttp.BaseRequest
	PosixAclId     *string                                       `json:"PosixAclId,omitempty" name:"PosixAclId"`
	FileSystemList []*UpdatePerformanceOnePosixAclFileSystemList `json:"FileSystemList,omitempty" name:"FileSystemList"`
	AutoMount      *bool                                         `json:"AutoMount,omitempty" name:"AutoMount"`
	Ips            []*string                                     `json:"Ips,omitempty" name:"Ips"`
	Desc           *string                                       `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdatePerformanceOnePosixAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdatePerformanceOnePosixAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdatePerformanceOnePosixAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePerformanceOnePosixAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePerformanceOnePosixAclListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	Ip             *string `json:"Ip,omitempty" name:"Ip"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribePerformanceOnePosixAclListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribePerformanceOnePosixAclListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		PosixAclId     *string `json:"PosixAclId" name:"PosixAclId"`
		FileSystemList []struct {
			FileSystemId   *string `json:"FileSystemId" name:"FileSystemId"`
			FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
			VolumePath     *string `json:"VolumePath" name:"VolumePath"`
		} `json:"FileSystemList" name:"FileSystemList"`
		Ips       []*string `json:"Ips" name:"Ips"`
		Desc      *string   `json:"Desc" name:"Desc"`
		AutoMount *bool     `json:"AutoMount" name:"AutoMount"`
	} `json:"Data"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribePerformanceOnePosixAclListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePerformanceOnePosixAclListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPerformanceOnePosixAclRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemList []*SetPerformanceOnePosixAclFileSystemList `json:"FileSystemList,omitempty" name:"FileSystemList"`
	AutoMount      *bool                                      `json:"AutoMount,omitempty" name:"AutoMount"`
	Ips            []*string                                  `json:"Ips,omitempty" name:"Ips"`
	Desc           *string                                    `json:"Desc,omitempty" name:"Desc"`
}

func (r *SetPerformanceOnePosixAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetPerformanceOnePosixAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	PosixAclId *string `json:"PosixAclId" name:"PosixAclId"`
}

func (r *SetPerformanceOnePosixAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPerformanceOnePosixAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirQuotaListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass     *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName    *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath        *string `json:"DirPath,omitempty" name:"DirPath"`
	FuzzySearch    *bool   `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeDirQuotaListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDirQuotaListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	PageNum    *int    `json:"PageNum" name:"PageNum"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Data       []struct {
		DirPath              *string `json:"DirPath" name:"DirPath"`
		LogicalCapacityType  *string `json:"LogicalCapacityType" name:"LogicalCapacityType"`
		LogicalHardThreshold *int64  `json:"LogicalHardThreshold" name:"LogicalHardThreshold"`
		LogicalUsedCapacity  *int64  `json:"LogicalUsedCapacity" name:"LogicalUsedCapacity"`
		LogicalInodesType    *string `json:"LogicalInodesType" name:"LogicalInodesType"`
		LogicalHardInodes    *int64  `json:"LogicalHardInodes" name:"LogicalHardInodes"`
		LogicalUsedInodes    *int64  `json:"LogicalUsedInodes" name:"LogicalUsedInodes"`
	} `json:"Data"`
}

func (r *DescribeDirQuotaListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirQuotaListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirQuotaRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass     *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName    *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath        *string `json:"DirPath,omitempty" name:"DirPath"`
}

func (r *DeleteDirQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDirQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDirQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDirQuotaRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId         *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass           *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName          *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName       *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath              *string `json:"DirPath,omitempty" name:"DirPath"`
	LogicalCapacityType  *string `json:"LogicalCapacityType,omitempty" name:"LogicalCapacityType"`
	LogicalHardThreshold *int64  `json:"LogicalHardThreshold,omitempty" name:"LogicalHardThreshold"`
	LogicalInodesType    *string `json:"LogicalInodesType,omitempty" name:"LogicalInodesType"`
	LogicalHardInodes    *int64  `json:"LogicalHardInodes,omitempty" name:"LogicalHardInodes"`
}

func (r *UpdateDirQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDirQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateDirQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDirQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirQuotaRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId         *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass           *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName          *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName       *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath              *string `json:"DirPath,omitempty" name:"DirPath"`
	LogicalCapacityType  *string `json:"LogicalCapacityType,omitempty" name:"LogicalCapacityType"`
	LogicalHardThreshold *int64  `json:"LogicalHardThreshold,omitempty" name:"LogicalHardThreshold"`
	LogicalInodesType    *string `json:"LogicalInodesType,omitempty" name:"LogicalInodesType"`
	LogicalHardInodes    *int64  `json:"LogicalHardInodes,omitempty" name:"LogicalHardInodes"`
}

func (r *CreateDirQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDirQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDirQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubDirListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass     *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName    *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath        *string `json:"DirPath,omitempty" name:"DirPath"`
	Name           *string `json:"Name,omitempty" name:"Name"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeSubDirListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSubDirListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Name       *string `json:"Name" name:"Name"`
		UpdateTime *int64  `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeSubDirListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubDirListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass     *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName    *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath        *string `json:"DirPath,omitempty" name:"DirPath"`
}

func (r *DeleteDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDirResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDirRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId           *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass             *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName            *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName         *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath                *string `json:"DirPath,omitempty" name:"DirPath"`
	FileSysPosixPermission *int    `json:"FileSysPosixPermission,omitempty" name:"FileSysPosixPermission"`
	FileSysOwnerUserId     *int    `json:"FileSysOwnerUserId,omitempty" name:"FileSysOwnerUserId"`
	FileSysOwnerGroupId    *int    `json:"FileSysOwnerGroupId,omitempty" name:"FileSysOwnerGroupId"`
}

func (r *UpdateDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDirResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId           *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass             *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName            *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName         *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath                *string `json:"DirPath,omitempty" name:"DirPath"`
	FileSysOwnerUserId     *int    `json:"FileSysOwnerUserId,omitempty" name:"FileSysOwnerUserId"`
	FileSysOwnerGroupId    *int    `json:"FileSysOwnerGroupId,omitempty" name:"FileSysOwnerGroupId"`
	FileSysPosixPermission *int    `json:"FileSysPosixPermission,omitempty" name:"FileSysPosixPermission"`
}

func (r *CreateDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDirResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirQuotaRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StoreClass     *string `json:"StoreClass,omitempty" name:"StoreClass"`
	ClusterName    *string `json:"ClusterName,omitempty" name:"ClusterName"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	DirPath        *string `json:"DirPath,omitempty" name:"DirPath"`
}

func (r *DescribeDirQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDirQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	LogicalHardThreshold *int64  `json:"LogicalHardThreshold" name:"LogicalHardThreshold"`
	LogicalUsedCapacity  *int64  `json:"LogicalUsedCapacity" name:"LogicalUsedCapacity"`
	LogicalCapacityType  *string `json:"LogicalCapacityType" name:"LogicalCapacityType"`
	LogicalInodesType    *string `json:"LogicalInodesType" name:"LogicalInodesType"`
	LogicalHardInodes    *int64  `json:"LogicalHardInodes" name:"LogicalHardInodes"`
	LogicalUsedInodes    *int64  `json:"LogicalUsedInodes" name:"LogicalUsedInodes"`
}

func (r *DescribeDirQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPerformanceOnePosixAclIpRequest struct {
	*ksyunhttp.BaseRequest
	PosixAclId *string `json:"PosixAclId,omitempty" name:"PosixAclId"`
	Ip         *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *AddPerformanceOnePosixAclIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddPerformanceOnePosixAclIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddPerformanceOnePosixAclIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPerformanceOnePosixAclIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemovePerformanceOnePosixAclIpRequest struct {
	*ksyunhttp.BaseRequest
	PosixAclId *string `json:"PosixAclId,omitempty" name:"PosixAclId"`
	Ip         *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *RemovePerformanceOnePosixAclIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RemovePerformanceOnePosixAclIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RemovePerformanceOnePosixAclIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemovePerformanceOnePosixAclIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataMigrateTaskProgressRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GetDataMigrateTaskProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDataMigrateTaskProgressResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	ExecStatus    *string `json:"ExecStatus" name:"ExecStatus"`
	ExecCount     *int    `json:"ExecCount" name:"ExecCount"`
	ExecStartTime *string `json:"ExecStartTime" name:"ExecStartTime"`
	ExecStatistic struct {
		IsActive          *bool   `json:"IsActive" name:"IsActive"`
		ScanObjects       *int64  `json:"ScanObjects" name:"ScanObjects"`
		TransSuccessCount *int64  `json:"TransSuccessCount" name:"TransSuccessCount"`
		TransFailedCount  *int64  `json:"TransFailedCount" name:"TransFailedCount"`
		TransTotalSize    *string `json:"TransTotalSize" name:"TransTotalSize"`
		TransTime         *string `json:"TransTime" name:"TransTime"`
	} `json:"ExecStatistic"`
	ProgressInfo struct {
		WTotalTransfers  *string `json:"WTotalTransfers" name:"WTotalTransfers"`
		WTotalSize       *string `json:"WTotalSize" name:"WTotalSize"`
		VTotalSpeed      *string `json:"VTotalSpeed" name:"VTotalSpeed"`
		VTotalTransfers  *int64  `json:"VTotalTransfers" name:"VTotalTransfers"`
		WTransfers       *int64  `json:"WTransfers" name:"WTransfers"`
		WElapsedTime     *string `json:"WElapsedTime" name:"WElapsedTime"`
		WTotalEta        *string `json:"WTotalEta" name:"WTotalEta"`
		WTotalPercentage *string `json:"WTotalPercentage" name:"WTotalPercentage"`
		Transferring     []struct {
			Bytes      *string `json:"Bytes" name:"Bytes"`
			Eta        *string `json:"Eta" name:"Eta"`
			FileName   *string `json:"FileName" name:"FileName"`
			Percentage *string `json:"Percentage" name:"Percentage"`
			Speed      *string `json:"Speed" name:"Speed"`
			SpeedAvg   *string `json:"SpeedAvg" name:"SpeedAvg"`
			FileSize   *string `json:"FileSize" name:"FileSize"`
		} `json:"Transferring" name:"Transferring"`
	} `json:"ProgressInfo"`
}

func (r *GetDataMigrateTaskProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataMigrateTaskProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataMigrateTaskListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	TaskIds      *string `json:"TaskIds,omitempty" name:"TaskIds"`
	TaskName     *string `json:"TaskName,omitempty" name:"TaskName"`
	TaskType     *string `json:"TaskType,omitempty" name:"TaskType"`
	DirPath      *string `json:"DirPath,omitempty" name:"DirPath"`
	Bucket       *string `json:"Bucket,omitempty" name:"Bucket"`
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeDataMigrateTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataMigrateTaskListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		FileSystemId *string `json:"FileSystemId" name:"FileSystemId"`
		TaskId       *string `json:"TaskId" name:"TaskId"`
		TaskType     *string `json:"TaskType" name:"TaskType"`
		TaskName     *string `json:"TaskName" name:"TaskName"`
		BucketConfig struct {
			Bucket       *string `json:"Bucket" name:"Bucket"`
			BucektPrefix *string `json:"BucektPrefix" name:"BucektPrefix"`
		} `json:"BucketConfig" name:"BucketConfig"`
		DirPath                 *string `json:"DirPath" name:"DirPath"`
		BandWidthLimit          *string `json:"BandWidthLimit" name:"BandWidthLimit"`
		CleanSourceFile         *bool   `json:"CleanSourceFile" name:"CleanSourceFile"`
		ExportTaskPeriodEnabled *string `json:"ExportTaskPeriodEnabled" name:"ExportTaskPeriodEnabled"`
		Description             *string `json:"Description" name:"Description"`
		ExportTaskPeriodConfig  struct {
			FrequencyUnit    *string   `json:"FrequencyUnit" name:"FrequencyUnit"`
			IndexOfFrequency []*int    `json:"IndexOfFrequency" name:"IndexOfFrequency"`
			TimePoints       []*string `json:"TimePoints" name:"TimePoints"`
		} `json:"ExportTaskPeriodConfig" name:"ExportTaskPeriodConfig"`
		ExecCount           *int    `json:"ExecCount" name:"ExecCount"`
		ExecResultMsg       *string `json:"ExecResultMsg" name:"ExecResultMsg"`
		ExecResultErrorCode *string `json:"ExecResultErrorCode" name:"ExecResultErrorCode"`
		ExecStartTime       *string `json:"ExecStartTime" name:"ExecStartTime"`
		ExecStatus          *string `json:"ExecStatus" name:"ExecStatus"`
		CreateTime          *int64  `json:"CreateTime" name:"CreateTime"`
		UpdateTime          *int64  `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeDataMigrateTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataMigrateTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartDataMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StartDataMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartDataMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StartDataMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartDataMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopDataMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopDataMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopDataMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StopDataMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopDataMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	TaskIds      *string `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DeleteDataMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDataMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDataMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDataMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId                  *string                                      `json:"TaskId,omitempty" name:"TaskId"`
	TaskName                *string                                      `json:"TaskName,omitempty" name:"TaskName"`
	DirPath                 *string                                      `json:"DirPath,omitempty" name:"DirPath"`
	Description             *string                                      `json:"Description,omitempty" name:"Description"`
	BandWidthLimit          *int                                         `json:"BandWidthLimit,omitempty" name:"BandWidthLimit"`
	CleanSourceFile         *bool                                        `json:"CleanSourceFile,omitempty" name:"CleanSourceFile"`
	ExportTaskPeriodEnabled *string                                      `json:"ExportTaskPeriodEnabled,omitempty" name:"ExportTaskPeriodEnabled"`
	ExportTaskPeriodConfig  *UpdateDataMigrateTaskExportTaskPeriodConfig `json:"ExportTaskPeriodConfig,omitempty" name:"ExportTaskPeriodConfig"`
}

func (r *UpdateDataMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDataMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateDataMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDataMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId            *string                                      `json:"FileSystemId,omitempty" name:"FileSystemId"`
	TaskName                *string                                      `json:"TaskName,omitempty" name:"TaskName"`
	TaskType                *string                                      `json:"TaskType,omitempty" name:"TaskType"`
	BucketConfig            *CreateDataMigrateTaskBucketConfig           `json:"BucketConfig,omitempty" name:"BucketConfig"`
	DirPath                 *string                                      `json:"DirPath,omitempty" name:"DirPath"`
	Description             *string                                      `json:"Description,omitempty" name:"Description"`
	BandWidthLimit          *int                                         `json:"BandWidthLimit,omitempty" name:"BandWidthLimit"`
	CleanSourceFile         *bool                                        `json:"CleanSourceFile,omitempty" name:"CleanSourceFile"`
	ExportTaskPeriodEnabled *string                                      `json:"ExportTaskPeriodEnabled,omitempty" name:"ExportTaskPeriodEnabled"`
	ExportTaskPeriodConfig  *CreateDataMigrateTaskExportTaskPeriodConfig `json:"ExportTaskPeriodConfig,omitempty" name:"ExportTaskPeriodConfig"`
}

func (r *CreateDataMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDataMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskId    *string `json:"TaskId" name:"TaskId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDataMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClientInstallInfoRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeClientInstallInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClientInstallInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	ClusterDataIP *string `json:"ClusterDataIP" name:"ClusterDataIP"`
	Data          []struct {
		DownloadUrl   *string `json:"DownloadUrl" name:"DownloadUrl"`
		OsVersion     *string `json:"OsVersion" name:"OsVersion"`
		KernelVersion *string `json:"KernelVersion" name:"KernelVersion"`
		NicDriver     *string `json:"NicDriver" name:"NicDriver"`
	} `json:"Data"`
}

func (r *DescribeClientInstallInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClientInstallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageDataFlowTaskRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
	TaskId       *string `json:"TaskId,omitempty" name:"TaskId"`
	Operation    *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ManageDataFlowTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ManageDataFlowTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ManageDataFlowTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageDataFlowTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataFlowStrategyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId     *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyName     *string `json:"StrategyName,omitempty" name:"StrategyName"`
	StrategyType     *string `json:"StrategyType,omitempty" name:"StrategyType"`
	Bind             *string `json:"Bind,omitempty" name:"Bind"`
	DataLoadingMode  *string `json:"DataLoadingMode,omitempty" name:"DataLoadingMode"`
	DirPath          *string `json:"DirPath,omitempty" name:"DirPath"`
	Bucket           *string `json:"Bucket,omitempty" name:"Bucket"`
	BucketPrefix     *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`
	DuplicateProcess *string `json:"DuplicateProcess,omitempty" name:"DuplicateProcess"`
	Subscribe        *string `json:"Subscribe,omitempty" name:"Subscribe"`
	CleanSourceFile  *bool   `json:"CleanSourceFile,omitempty" name:"CleanSourceFile"`
	BandWidthLimit   *string `json:"BandWidthLimit,omitempty" name:"BandWidthLimit"`
	ArchiveRule      *int    `json:"ArchiveRule,omitempty" name:"ArchiveRule"`
}

func (r *CreateDataFlowStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDataFlowStrategyResponse struct {
	*ksyunhttp.BaseResponse
	StrategyId *string `json:"StrategyId" name:"StrategyId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDataFlowStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataFlowStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataFlowTaskListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeDataFlowTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowTaskListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		TaskId             *string `json:"TaskId" name:"TaskId"`
		Status             *string `json:"Status" name:"Status"`
		TransferFileCount  *int64  `json:"TransferFileCount" name:"TransferFileCount"`
		TransferFileData   *int64  `json:"TransferFileData" name:"TransferFileData"`
		TransferThroughput *int64  `json:"TransferThroughput" name:"TransferThroughput"`
		TransferOps        *int64  `json:"TransferOps" name:"TransferOps"`
		StartTime          *string `json:"StartTime" name:"StartTime"`
		EndTime            *string `json:"EndTime" name:"EndTime"`
	} `json:"Data"`
}

func (r *DescribeDataFlowTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateDataFlowTaskRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *ActivateDataFlowTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateDataFlowTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ActivateDataFlowTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateDataFlowTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataFlowStrategyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteDataFlowStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDataFlowStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDataFlowStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataFlowStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataFlowStrategyListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeDataFlowStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowStrategyListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
	PageNum    *int    `json:"PageNum" name:"PageNum"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Data       []struct {
		StrategyType     *string `json:"StrategyType" name:"StrategyType"`
		Bind             *string `json:"Bind" name:"Bind"`
		StrategyId       *string `json:"StrategyId" name:"StrategyId"`
		DataLoadingMode  *string `json:"DataLoadingMode" name:"DataLoadingMode"`
		StrategyName     *string `json:"StrategyName" name:"StrategyName"`
		FileSystemId     *string `json:"FileSystemId" name:"FileSystemId"`
		DirPath          *string `json:"DirPath" name:"DirPath"`
		Bucket           *string `json:"Bucket" name:"Bucket"`
		BucketPrefix     *string `json:"BucketPrefix" name:"BucketPrefix"`
		DuplicateProcess *string `json:"DuplicateProcess" name:"DuplicateProcess"`
		Subscribe        *string `json:"Subscribe" name:"Subscribe"`
		Status           *string `json:"Status" name:"Status"`
		UpdateTime       *string `json:"UpdateTime" name:"UpdateTime"`
		CleanSourceFile  *bool   `json:"CleanSourceFile" name:"CleanSourceFile"`
		BandWidthLimit   *string `json:"BandWidthLimit" name:"BandWidthLimit"`
		ArchiveRule      *string `json:"ArchiveRule" name:"ArchiveRule"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
	} `json:"Data"`
}

func (r *DescribeDataFlowStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CleanRecycledFilesRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *CleanRecycledFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CleanRecycledFilesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CleanRecycledFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanRecycledFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCleanRecycledFilesRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteCleanRecycledFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCleanRecycledFilesResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteCleanRecycledFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCleanRecycledFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecycleBinConfigRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteRecycleBinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRecycleBinConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRecycleBinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecycleBinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecycledFileListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteRecycledFileListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRecycledFileListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteRecycledFileListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecycledFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecycleBinConfigRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *GetRecycleBinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRecycleBinConfigResponse struct {
	*ksyunhttp.BaseResponse
	Enabled    *string `json:"Enabled" name:"Enabled"`
	ExpireTime *int    `json:"ExpireTime" name:"ExpireTime"`
	ExpireType *string `json:"ExpireType" name:"ExpireType"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *GetRecycleBinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecycleBinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRecycleBinConfigRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	Enabled      *string `json:"Enabled,omitempty" name:"Enabled"`
	ExpireTime   *int    `json:"ExpireTime,omitempty" name:"ExpireTime"`
	ExpireType   *string `json:"ExpireType,omitempty" name:"ExpireType"`
}

func (r *SetRecycleBinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetRecycleBinConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetRecycleBinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRecycleBinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecycledFileListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	RecycledPath *string `json:"RecycledPath,omitempty" name:"RecycledPath"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRecycledFileListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRecycledFileListResponse struct {
	*ksyunhttp.BaseResponse
	Files []struct {
		FileName     *string `json:"FileName" name:"FileName"`
		RecycledPath *string `json:"RecycledPath" name:"RecycledPath"`
		DeleteTime   *int64  `json:"DeleteTime" name:"DeleteTime"`
		Position     *string `json:"Position" name:"Position"`
		Type         *string `json:"Type" name:"Type"`
		Length       *int64  `json:"Length" name:"Length"`
		Inode        *int    `json:"Inode" name:"Inode"`
	} `json:"Files"`
	TotalBytes *int    `json:"TotalBytes" name:"TotalBytes"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeRecycledFileListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecycledFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecycledFilesRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string                     `json:"FileSystemId,omitempty" name:"FileSystemId"`
	RecycledPath *string                     `json:"RecycledPath,omitempty" name:"RecycledPath"`
	Files        []*DeleteRecycledFilesFiles `json:"Files,omitempty" name:"Files"`
	Inodes       []*int                      `json:"Inodes,omitempty" name:"Inodes"`
}

func (r *DeleteRecycledFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRecycledFilesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRecycledFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecycledFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreRecycledFilesRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string                      `json:"FileSystemId,omitempty" name:"FileSystemId"`
	RecycledPath *string                      `json:"RecycledPath,omitempty" name:"RecycledPath"`
	Files        []*RestoreRecycledFilesFiles `json:"Files,omitempty" name:"Files"`
	Inodes       []*int                       `json:"Inodes,omitempty" name:"Inodes"`
}

func (r *RestoreRecycledFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestoreRecycledFilesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreRecycledFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreRecycledFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInfoRequest struct {
	*ksyunhttp.BaseRequest
	Region        *string `json:"Region,omitempty" name:"Region"`
	AvailZone     *string `json:"AvailZone,omitempty" name:"AvailZone"`
	StoreClass    *string `json:"StoreClass,omitempty" name:"StoreClass"`
	SRoceCluster  *string `json:"SRoceCluster,omitempty" name:"SRoceCluster"`
	StorePoolType *string `json:"StorePoolType,omitempty" name:"StorePoolType"`
}

func (r *DescribeClusterInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClusterInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Region        *string   `json:"Region" name:"Region"`
		AvailZone     *string   `json:"AvailZone" name:"AvailZone"`
		StoreClasses  []*string `json:"StoreClasses" name:"StoreClasses"`
		StorePoolType *string   `json:"StorePoolType" name:"StorePoolType"`
		SRoceCluster  *string   `json:"SRoceCluster" name:"SRoceCluster"`
		ClusterCode   *string   `json:"ClusterCode" name:"ClusterCode"`
	} `json:"Data"`
}

func (r *DescribeClusterInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePerformanceNfsAclIpRequest struct {
	*ksyunhttp.BaseRequest
	NfsAclId *string                         `json:"NfsAclId,omitempty" name:"NfsAclId"`
	Ips      []*UpdatePerformanceNfsAclIpIps `json:"Ips,omitempty" name:"Ips"`
}

func (r *UpdatePerformanceNfsAclIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdatePerformanceNfsAclIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdatePerformanceNfsAclIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePerformanceNfsAclIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemovePerformanceNfsAclClientRequest struct {
	*ksyunhttp.BaseRequest
	NfsAclId *string   `json:"NfsAclId,omitempty" name:"NfsAclId"`
	Ips      []*string `json:"Ips,omitempty" name:"Ips"`
}

func (r *RemovePerformanceNfsAclClientRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RemovePerformanceNfsAclClientResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RemovePerformanceNfsAclClientResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemovePerformanceNfsAclClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPerformanceNfsAclClientRequest struct {
	*ksyunhttp.BaseRequest
	NfsAclId *string                          `json:"NfsAclId,omitempty" name:"NfsAclId"`
	Ips      []*AddPerformanceNfsAclClientIps `json:"Ips,omitempty" name:"Ips"`
}

func (r *AddPerformanceNfsAclClientRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddPerformanceNfsAclClientResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddPerformanceNfsAclClientResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPerformanceNfsAclClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePerformanceOneNfsAclRequest struct {
	*ksyunhttp.BaseRequest
	NfsAclId *string `json:"NfsAclId,omitempty" name:"NfsAclId"`
}

func (r *DeletePerformanceOneNfsAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeletePerformanceOneNfsAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeletePerformanceOneNfsAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePerformanceOneNfsAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPerformanceOneNfsAclRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string                       `json:"FileSystemId,omitempty" name:"FileSystemId"`
	ExportPath   *string                       `json:"ExportPath,omitempty" name:"ExportPath"`
	Ips          []*SetPerformanceOneNfsAclIps `json:"Ips,omitempty" name:"Ips"`
	Desc         *string                       `json:"Desc,omitempty" name:"Desc"`
}

func (r *SetPerformanceOneNfsAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetPerformanceOneNfsAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NfsAclId  *string `json:"NfsAclId" name:"NfsAclId"`
}

func (r *SetPerformanceOneNfsAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPerformanceOneNfsAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePerformanceOneNfsAclListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	NfsAclId       *string `json:"NfsAclId,omitempty" name:"NfsAclId"`
	PageNum        *int64  `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *int64  `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribePerformanceOneNfsAclListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribePerformanceOneNfsAclListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		NfsAclId       *string `json:"NfsAclId" name:"NfsAclId"`
		ExportPath     *string `json:"ExportPath" name:"ExportPath"`
		FileSystemList []struct {
			FileSystemId   *string `json:"FileSystemId" name:"FileSystemId"`
			Region         *string `json:"Region" name:"Region"`
			FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
			RegionName     *string `json:"RegionName" name:"RegionName"`
		} `json:"FileSystemList" name:"FileSystemList"`
		MountDomain *string `json:"MountDomain" name:"MountDomain"`
		Ips         []struct {
			Ip         *string `json:"Ip" name:"Ip"`
			Permission *string `json:"Permission" name:"Permission"`
			RootSquash *string `json:"RootSquash" name:"RootSquash"`
			Hostname   *string `json:"Hostname" name:"Hostname"`
			Type       *string `json:"Type" name:"Type"`
		} `json:"Ips" name:"Ips"`
		Desc *string `json:"Desc" name:"Desc"`
	} `json:"Data"`
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribePerformanceOneNfsAclListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePerformanceOneNfsAclListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemNfsClientInfoRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
	Action       *string `json:"Action,omitempty" name:"Action"`
	Version      *string `json:"Version,omitempty" name:"Version"`
}

func (r *DescribeFileSystemNfsClientInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFileSystemNfsClientInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string   `json:"RequestId" name:"RequestId"`
	Data       []*string `json:"Data" name:"Data"`
	PageNum    *int64    `json:"PageNum" name:"PageNum"`
	PageSize   *int64    `json:"PageSize" name:"PageSize"`
	TotalCount *int64    `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeFileSystemNfsClientInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSystemNfsClientInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetFileSystemResourceProtectRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemIds []*string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
	IsProtection  *bool     `json:"IsProtection,omitempty" name:"IsProtection"`
}

func (r *SetFileSystemResourceProtectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetFileSystemResourceProtectResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetFileSystemResourceProtectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetFileSystemResourceProtectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileDeletePolicyListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId       *string   `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DeletePolicyStatus *string   `json:"DeletePolicyStatus,omitempty" name:"DeletePolicyStatus"`
	DirPath            *string   `json:"DirPath,omitempty" name:"DirPath"`
	FileDeletePolicyId []*string `json:"FileDeletePolicyId,omitempty" name:"FileDeletePolicyId"`
	PageNum            *int      `json:"PageNum,omitempty" name:"PageNum"`
	PageSize           *int      `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFileDeletePolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFileDeletePolicyListResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount           *int `json:"TotalCount" name:"TotalCount"`
	PageSize             *int `json:"PageSize" name:"PageSize"`
	PageNumber           *int `json:"PageNumber" name:"PageNumber"`
	FileDeletePolicyList []struct {
		FileSystemId        *string `json:"FileSystemId" name:"FileSystemId"`
		FileSystemName      *string `json:"FileSystemName" name:"FileSystemName"`
		FileDeletePolicyId  *string `json:"FileDeletePolicyId" name:"FileDeletePolicyId"`
		DirPath             *string `json:"DirPath" name:"DirPath"`
		ExecutionType       *string `json:"ExecutionType" name:"ExecutionType"`
		FrequencyUnit       *string `json:"FrequencyUnit" name:"FrequencyUnit"`
		IndexOfFrequency    []*int  `json:"IndexOfFrequency" name:"IndexOfFrequency"`
		FrequencyTimePoints []struct {
			Start struct {
				Hour *int `json:"Hour" name:"Hour"`
			} `json:"Start"`
			End struct {
				Hour *int `json:"Hour" name:"Hour"`
			} `json:"End"`
		} `json:"FrequencyTimePoints" name:"FrequencyTimePoints"`
		FileSizeRule struct {
			Rule     *string `json:"Rule" name:"Rule"`
			MaxValue *int    `json:"MaxValue" name:"MaxValue"`
			Unit     *string `json:"Unit" name:"Unit"`
			MinValue *int    `json:"MinValue" name:"MinValue"`
		} `json:"FileSizeRule" name:"FileSizeRule"`
		TimeRuleParameters []struct {
			Type   *string `json:"Type" name:"Type"`
			OpType *string `json:"OpType" name:"OpType"`
			Unit   *string `json:"Unit" name:"Unit"`
			Value  *int    `json:"Value" name:"Value"`
		} `json:"TimeRuleParameters" name:"TimeRuleParameters"`
		DeletePolicyStatus *string `json:"DeletePolicyStatus" name:"DeletePolicyStatus"`
		CreateTime         *int64  `json:"CreateTime" name:"CreateTime"`
		Description        *string `json:"Description" name:"Description"`
		FileNameRule       struct {
			Rule *string `json:"Rule" name:"Rule"`
		} `json:"FileNameRule" name:"FileNameRule"`
		FileDeletePolicyName *string `json:"FileDeletePolicyName" name:"FileDeletePolicyName"`
	} `json:"FileDeletePolicyList"`
}

func (r *DescribeFileDeletePolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileDeletePolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableFileDeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId       *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileDeletePolicyId *string `json:"FileDeletePolicyId,omitempty" name:"FileDeletePolicyId"`
}

func (r *EnableFileDeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EnableFileDeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EnableFileDeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableFileDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableFileDeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId       *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileDeletePolicyId *string `json:"FileDeletePolicyId,omitempty" name:"FileDeletePolicyId"`
}

func (r *DisableFileDeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisableFileDeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DisableFileDeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableFileDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileDeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	FileDeletePolicyId *string `json:"FileDeletePolicyId,omitempty" name:"FileDeletePolicyId"`
	FileSystemId       *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeFileDeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFileDeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	FileSystemId         *string `json:"FileSystemId" name:"FileSystemId"`
	FileSystemName       *string `json:"FileSystemName" name:"FileSystemName"`
	FileDeletePolicyId   *string `json:"FileDeletePolicyId" name:"FileDeletePolicyId"`
	FileDeletePolicyName *string `json:"FileDeletePolicyName" name:"FileDeletePolicyName"`
	DirPath              *string `json:"DirPath" name:"DirPath"`
	ExecutionType        *string `json:"ExecutionType" name:"ExecutionType"`
	FrequencyUnit        *string `json:"FrequencyUnit" name:"FrequencyUnit"`
	IndexOfFrequency     []*int  `json:"IndexOfFrequency" name:"IndexOfFrequency"`
	FrequencyTimePoints  []struct {
		Start struct {
			Hour *int `json:"Hour" name:"Hour"`
		} `json:"Start" name:"Start"`
		End struct {
			Hour *int `json:"Hour" name:"Hour"`
		} `json:"End" name:"End"`
	} `json:"FrequencyTimePoints"`
	FileSizeRule struct {
		Rule     *string `json:"Rule" name:"Rule"`
		MaxValue *int    `json:"MaxValue" name:"MaxValue"`
		MinValue *int    `json:"MinValue" name:"MinValue"`
		Unit     *string `json:"Unit" name:"Unit"`
	} `json:"FileSizeRule"`
	TimeRuleParameters []struct {
		Type   *string `json:"Type" name:"Type"`
		OpType *string `json:"OpType" name:"OpType"`
		Unit   *string `json:"Unit" name:"Unit"`
		Value  *int    `json:"Value" name:"Value"`
	} `json:"TimeRuleParameters"`
	DeletePolicyStatus *string `json:"DeletePolicyStatus" name:"DeletePolicyStatus"`
	CreateTime         *int64  `json:"CreateTime" name:"CreateTime"`
	Description        *string `json:"Description" name:"Description"`
	FileNameRule       struct {
		Rule *string `json:"Rule" name:"Rule"`
	} `json:"FileNameRule"`
}

func (r *DescribeFileDeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFileDeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId       *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileDeletePolicyId *string `json:"FileDeletePolicyId,omitempty" name:"FileDeletePolicyId"`
}

func (r *DeleteFileDeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteFileDeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteFileDeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFileDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFileDeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId         *string                                      `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileDeletePolicyId   *string                                      `json:"FileDeletePolicyId,omitempty" name:"FileDeletePolicyId"`
	FileDeletePolicyName *string                                      `json:"FileDeletePolicyName,omitempty" name:"FileDeletePolicyName"`
	ExecutionType        *string                                      `json:"ExecutionType,omitempty" name:"ExecutionType"`
	FrequencyUnit        *string                                      `json:"FrequencyUnit,omitempty" name:"FrequencyUnit"`
	IndexOfFrequency     []*int                                       `json:"IndexOfFrequency,omitempty" name:"IndexOfFrequency"`
	FrequencyTimePoints  []*UpdateFileDeletePolicyFrequencyTimePoints `json:"FrequencyTimePoints,omitempty" name:"FrequencyTimePoints"`
	FileNameRule         *UpdateFileDeletePolicyFileNameRule          `json:"FileNameRule,omitempty" name:"FileNameRule"`
	FileSizeRule         *UpdateFileDeletePolicyFileSizeRule          `json:"FileSizeRule,omitempty" name:"FileSizeRule"`
	TimeRuleParameters   []*UpdateFileDeletePolicyTimeRuleParameters  `json:"TimeRuleParameters,omitempty" name:"TimeRuleParameters"`
	Description          *string                                      `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateFileDeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateFileDeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateFileDeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFileDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFileDeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId         *string                                      `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileDeletePolicyName *string                                      `json:"FileDeletePolicyName,omitempty" name:"FileDeletePolicyName"`
	DirPath              *string                                      `json:"DirPath,omitempty" name:"DirPath"`
	ExecutionType        *string                                      `json:"ExecutionType,omitempty" name:"ExecutionType"`
	FrequencyUnit        *string                                      `json:"FrequencyUnit,omitempty" name:"FrequencyUnit"`
	IndexOfFrequency     []*int                                       `json:"IndexOfFrequency,omitempty" name:"IndexOfFrequency"`
	FrequencyTimePoints  []*CreateFileDeletePolicyFrequencyTimePoints `json:"FrequencyTimePoints,omitempty" name:"FrequencyTimePoints"`
	FileNameRule         *CreateFileDeletePolicyFileNameRule          `json:"FileNameRule,omitempty" name:"FileNameRule"`
	FileSizeRule         *CreateFileDeletePolicyFileSizeRule          `json:"FileSizeRule,omitempty" name:"FileSizeRule"`
	TimeRuleParameters   []*CreateFileDeletePolicyTimeRuleParameters  `json:"TimeRuleParameters,omitempty" name:"TimeRuleParameters"`
	Description          *string                                      `json:"Description,omitempty" name:"Description"`
}

func (r *CreateFileDeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateFileDeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	FileDeletePolicyId *string `json:"FileDeletePolicyId" name:"FileDeletePolicyId"`
	RequestId          *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateFileDeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFileDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataFlowStrategySubscribeRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeDataFlowStrategySubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowStrategySubscribeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		SubscribeId *string `json:"SubscribeId" name:"SubscribeId"`
		StartTime   *string `json:"StartTime" name:"StartTime"`
		EndTime     *string `json:"EndTime" name:"EndTime"`
	} `json:"Data"`
}

func (r *DescribeDataFlowStrategySubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowStrategySubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageDataFlowStrategySubscribeRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
	Operation    *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ManageDataFlowStrategySubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ManageDataFlowStrategySubscribeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ManageDataFlowStrategySubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageDataFlowStrategySubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRemoteCachePutLatencyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime      *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval       *string `json:"Interval,omitempty" name:"Interval"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	ClientNm       *string `json:"ClientNm,omitempty" name:"ClientNm"`
}

func (r *GetRemoteCachePutLatencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRemoteCachePutLatencyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64   `json:"Time" name:"Time"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetRemoteCachePutLatencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRemoteCachePutLatencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRemoteCacheGetLatencyRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime      *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval       *string `json:"Interval,omitempty" name:"Interval"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	ClientNm       *string `json:"ClientNm,omitempty" name:"ClientNm"`
}

func (r *GetRemoteCacheGetLatencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRemoteCacheGetLatencyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64   `json:"Time" name:"Time"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetRemoteCacheGetLatencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRemoteCacheGetLatencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRemoteCachePutThroughputRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime      *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval       *string `json:"Interval,omitempty" name:"Interval"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	ClientNm       *string `json:"ClientNm,omitempty" name:"ClientNm"`
}

func (r *GetRemoteCachePutThroughputRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRemoteCachePutThroughputResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64   `json:"Time" name:"Time"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetRemoteCachePutThroughputResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRemoteCachePutThroughputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRemoteCacheGetThroughputRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime      *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval       *string `json:"Interval,omitempty" name:"Interval"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	ClientNm       *string `json:"ClientNm,omitempty" name:"ClientNm"`
}

func (r *GetRemoteCacheGetThroughputRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRemoteCacheGetThroughputResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64   `json:"Time" name:"Time"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetRemoteCacheGetThroughputResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRemoteCacheGetThroughputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRemoteCacheIOPSSendRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime      *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval       *string `json:"Interval,omitempty" name:"Interval"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	ClientNm       *string `json:"ClientNm,omitempty" name:"ClientNm"`
}

func (r *GetRemoteCacheIOPSSendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRemoteCacheIOPSSendResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64   `json:"Time" name:"Time"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetRemoteCacheIOPSSendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRemoteCacheIOPSSendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRemoteCacheIOPSGetRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StartTime      *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *string `json:"EndTime,omitempty" name:"EndTime"`
	Interval       *string `json:"Interval,omitempty" name:"Interval"`
	CacheGroup     *string `json:"CacheGroup,omitempty" name:"CacheGroup"`
	CacheGroupRole *string `json:"CacheGroupRole,omitempty" name:"CacheGroupRole"`
	ClientNm       *string `json:"ClientNm,omitempty" name:"ClientNm"`
}

func (r *GetRemoteCacheIOPSGetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRemoteCacheIOPSGetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Time  *int64   `json:"Time" name:"Time"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetRemoteCacheIOPSGetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRemoteCacheIOPSGetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataFlowStrategySubscribeFailedRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	StrategyId   *string `json:"StrategyId,omitempty" name:"StrategyId"`
	SubscribeId  *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeDataFlowStrategySubscribeFailedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowStrategySubscribeFailedResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		EventTime    *string `json:"EventTime" name:"EventTime"`
		ObjectKey    *string `json:"ObjectKey" name:"ObjectKey"`
		EventType    *string `json:"EventType" name:"EventType"`
		FailedReason *string `json:"FailedReason" name:"FailedReason"`
		Content      *string `json:"Content" name:"Content"`
	} `json:"Data"`
}

func (r *DescribeDataFlowStrategySubscribeFailedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowStrategySubscribeFailedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskId    *string `json:"TaskId,omitempty" name:"TaskId"`
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ManageMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ManageMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ManageMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateTasksRequest struct {
	*ksyunhttp.BaseRequest
	RuleId   *string `json:"RuleId,omitempty" name:"RuleId"`
	TaskId   *string `json:"TaskId,omitempty" name:"TaskId"`
	PageSize *int64  `json:"PageSize,omitempty" name:"PageSize"`
	PageNum  *int64  `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeMigrateTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMigrateTasksResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		TaskId             *string `json:"TaskId" name:"TaskId"`
		SrcDirectory       *string `json:"SrcDirectory" name:"SrcDirectory"`
		DstDirectory       *string `json:"DstDirectory" name:"DstDirectory"`
		Status             *string `json:"Status" name:"Status"`
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		StartTime          *string `json:"StartTime" name:"StartTime"`
		EndTime            *string `json:"EndTime" name:"EndTime"`
		ExcuteTime         *string `json:"ExcuteTime" name:"ExcuteTime"`
		ScanFileCount      *int64  `json:"ScanFileCount" name:"ScanFileCount"`
		CompletedFileCount *int64  `json:"CompletedFileCount" name:"CompletedFileCount"`
		FailedFileCount    *int64  `json:"FailedFileCount" name:"FailedFileCount"`
		ScanFileBytes      *int64  `json:"ScanFileBytes" name:"ScanFileBytes"`
		CompletedFileBytes *int64  `json:"CompletedFileBytes" name:"CompletedFileBytes"`
		TaskIops           *int64  `json:"TaskIops" name:"TaskIops"`
		TaskBandwidth      *int64  `json:"TaskBandwidth" name:"TaskBandwidth"`
		RuleId             *string `json:"RuleId" name:"RuleId"`
		Bandwidth          *int    `json:"Bandwidth" name:"Bandwidth"`
		UpdateTime         *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
	PageSize   *int64  `json:"PageSize" name:"PageSize"`
	PageNum    *int64  `json:"PageNum" name:"PageNum"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeMigrateTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateTaskRequest struct {
	*ksyunhttp.BaseRequest
	RuleId       *string   `json:"RuleId,omitempty" name:"RuleId"`
	SrcDirectory *string   `json:"SrcDirectory,omitempty" name:"SrcDirectory"`
	DstDirectory *string   `json:"DstDirectory,omitempty" name:"DstDirectory"`
	EntryList    []*string `json:"EntryList,omitempty" name:"EntryList"`
}

func (r *CreateMigrateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMigrateTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskId    *string `json:"TaskId" name:"TaskId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateMigrateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMigrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteMigrateRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteMigrateRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteMigrateRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMigrateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateRulesRequest struct {
	*ksyunhttp.BaseRequest
	RuleId   *string `json:"RuleId,omitempty" name:"RuleId"`
	Region   *string `json:"Region,omitempty" name:"Region"`
	PageSize *int64  `json:"PageSize,omitempty" name:"PageSize"`
	PageNum  *int64  `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeMigrateRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMigrateRulesResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		RuleId        *string `json:"RuleId" name:"RuleId"`
		RuleName      *string `json:"RuleName" name:"RuleName"`
		Region        *string `json:"Region" name:"Region"`
		SourceAddress struct {
			AddressType    *string `json:"AddressType" name:"AddressType"`
			FileSystemId   *string `json:"FileSystemId" name:"FileSystemId"`
			FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
			Path           *string `json:"Path" name:"Path"`
			BucketName     *string `json:"BucketName" name:"BucketName"`
			BucketPrefix   *string `json:"BucketPrefix" name:"BucketPrefix"`
		} `json:"SourceAddress" name:"SourceAddress"`
		TargetAddress struct {
			AddressType    *string `json:"AddressType" name:"AddressType"`
			FileSystemId   *string `json:"FileSystemId" name:"FileSystemId"`
			FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
			Path           *string `json:"Path" name:"Path"`
			BucketName     *string `json:"BucketName" name:"BucketName"`
			BucketPrefix   *string `json:"BucketPrefix" name:"BucketPrefix"`
		} `json:"TargetAddress" name:"TargetAddress"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
	} `json:"Data"`
	PageSize   *int64  `json:"PageSize" name:"PageSize"`
	PageNum    *int64  `json:"PageNum" name:"PageNum"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeMigrateRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateRuleRequest struct {
	*ksyunhttp.BaseRequest
	Name    *string                   `json:"Name,omitempty" name:"Name"`
	Region  *string                   `json:"Region,omitempty" name:"Region"`
	SrcData *CreateMigrateRuleSrcData `json:"SrcData,omitempty" name:"SrcData"`
	DstData *CreateMigrateRuleDstData `json:"DstData,omitempty" name:"DstData"`
}

func (r *CreateMigrateRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMigrateRuleResponse struct {
	*ksyunhttp.BaseResponse
	RuleId    *string `json:"RuleId" name:"RuleId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateMigrateRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMigrateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
