package v20240930

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

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
