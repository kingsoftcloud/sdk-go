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
