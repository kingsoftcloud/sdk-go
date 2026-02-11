package v20240930

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

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

type DescribeFileSystemListRequest struct {
	*ksyunhttp.BaseRequest
	Region         *string `json:"Region,omitempty" name:"Region"`
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
	PageSize   *int `json:"PageSize" name:"PageSize"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
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
		LogicalHardThreshold *int    `json:"LogicalHardThreshold" name:"LogicalHardThreshold"`
		LogicalUsedCapacity  *int    `json:"LogicalUsedCapacity" name:"LogicalUsedCapacity"`
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
	LogicalHardThreshold *int64  `json:"LogicalHardThreshold,omitempty" name:"LogicalHardThreshold"`
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
	LogicalHardThreshold *int64  `json:"LogicalHardThreshold,omitempty" name:"LogicalHardThreshold"`
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
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Data       []struct {
		Name       *string `json:"Name" name:"Name"`
		UpdateTime *int64  `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
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
}

func (r *DescribeDirQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirQuotaResponse) FromJsonString(s string) error {
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
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
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
			MountDomain    *string `json:"MountDomain" name:"MountDomain"`
		} `json:"FileSystemList" name:"FileSystemList"`
		Ips []struct {
			Ip         *string `json:"Ip" name:"Ip"`
			Permission *string `json:"Permission" name:"Permission"`
			RootSquash *string `json:"RootSquash" name:"RootSquash"`
			Hostname   *string `json:"Hostname" name:"Hostname"`
			Type       *string `json:"Type" name:"Type"`
		} `json:"Ips" name:"Ips"`
		Desc *string `json:"Desc" name:"Desc"`
	} `json:"Data"`
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

type DeleteDataFlowRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DataFlowId   *string `json:"DataFlowId,omitempty" name:"DataFlowId"`
}

func (r *DeleteDataFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDataFlowResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDataFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataFlowTasksRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DataFlowId   *string `json:"DataFlowId,omitempty" name:"DataFlowId"`
	TaskIds      *string `json:"TaskIds,omitempty" name:"TaskIds"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeDataFlowTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowTasksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		FileSystemId *string `json:"FileSystemId" name:"FileSystemId"`
		DataFlowId   *string `json:"DataFlowId" name:"DataFlowId"`
		Name         *string `json:"Name" name:"Name"`
		DirPath      *string `json:"DirPath" name:"DirPath"`
		Bucket       *string `json:"Bucket" name:"Bucket"`
		BucketPrefix *string `json:"BucketPrefix" name:"BucketPrefix"`
		TaskId       *string `json:"TaskId" name:"TaskId"`
		TaskAction   *string `json:"TaskAction" name:"TaskAction"`
		SrcDirectory *string `json:"SrcDirectory" name:"SrcDirectory"`
		DstDirectory *string `json:"DstDirectory" name:"DstDirectory"`
		Status       *string `json:"Status" name:"Status"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		StartTime    *string `json:"StartTime" name:"StartTime"`
		EndTime      *string `json:"EndTime" name:"EndTime"`
		ExcuteTime   *string `json:"ExcuteTime" name:"ExcuteTime"`
	} `json:"Data"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeDataFlowTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataFlowsRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DataFlowId   *string `json:"DataFlowId,omitempty" name:"DataFlowId"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeDataFlowsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		DataFlowId   *string `json:"DataFlowId" name:"DataFlowId"`
		Name         *string `json:"Name" name:"Name"`
		FileSystemId *string `json:"FileSystemId" name:"FileSystemId"`
		DirPath      *string `json:"DirPath" name:"DirPath"`
		Bucket       *string `json:"Bucket" name:"Bucket"`
		BucketPrefix *string `json:"BucketPrefix" name:"BucketPrefix"`
		Status       *string `json:"Status" name:"Status"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime   *string `json:"UpdateTime" name:"UpdateTime"`
		Description  *string `json:"Description" name:"Description"`
	} `json:"Data"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeDataFlowsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataFlowTaskRequest struct {
	*ksyunhttp.BaseRequest
	DataFlowId   *string   `json:"DataFlowId,omitempty" name:"DataFlowId"`
	TaskAction   *string   `json:"TaskAction,omitempty" name:"TaskAction"`
	SrcDirectory *string   `json:"SrcDirectory,omitempty" name:"SrcDirectory"`
	DstDirectory *string   `json:"DstDirectory,omitempty" name:"DstDirectory"`
	EntryList    []*string `json:"EntryList,omitempty" name:"EntryList"`
	Bandwidth    *int      `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

func (r *CreateDataFlowTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDataFlowTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	TaskId    *string `json:"TaskId" name:"TaskId"`
}

func (r *CreateDataFlowTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataFlowTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataFlowRequest struct {
	*ksyunhttp.BaseRequest
	Name         *string `json:"Name,omitempty" name:"Name"`
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DirPath      *string `json:"DirPath,omitempty" name:"DirPath"`
	Bucket       *string `json:"Bucket,omitempty" name:"Bucket"`
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`
	Description  *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDataFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDataFlowResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	DataFlowId *string `json:"DataFlowId" name:"DataFlowId"`
}

func (r *CreateDataFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
