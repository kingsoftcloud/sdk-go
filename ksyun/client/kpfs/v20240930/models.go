package v20240930

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type UpdatePerformanceOnePosixAclFileSystemList struct {
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	VolumePath   *string `json:"VolumePath,omitempty" name:"VolumePath"`
}

type UpdatePerformanceOnePosixAclRequest struct {
	*ksyunhttp.BaseRequest
	PosixAclId     *string                                       `json:"PosixAclId,omitempty" name:"PosixAclId"`
	FileSystemList []*UpdatePerformanceOnePosixAclFileSystemList `json:"FileSystemList,omitempty" name:"FileSystemList"`
	Ips            []*string                                     `json:"Ips,omitempty" name:"Ips"`
	AutoMount      *bool                                         `json:"AutoMount,omitempty" name:"AutoMount"`
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
	PosixAclId     *string `json:"PosixAclId,omitempty" name:"PosixAclId"`
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
			VolumePath     *string `json:"VolumePath" name:"VolumePath"`
			FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
		} `json:"FileSystemList" name:"FileSystemList"`
		Ips       []*string `json:"Ips" name:"Ips"`
		Desc      *string   `json:"Desc" name:"Desc"`
		AutoMount *bool     `json:"AutoMount" name:"AutoMount"`
	} `json:"Data"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	PageNum    *int `json:"PageNum" name:"PageNum"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribePerformanceOnePosixAclListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePerformanceOnePosixAclListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
