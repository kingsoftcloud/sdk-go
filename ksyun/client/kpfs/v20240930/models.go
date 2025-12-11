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


type ManageDataFlowTaskRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *ManageDataFlowTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ManageDataFlowTaskResponse struct {
	*ksyunhttp.BaseResponse
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *CreateDataFlowStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDataFlowStrategyResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *CreateDataFlowStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataFlowStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyDataFlowTaskRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *ModifyDataFlowTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDataFlowTaskResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ModifyDataFlowTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataFlowTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDataFlowTaskListRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeDataFlowTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowTaskListResponse struct {
	*ksyunhttp.BaseResponse
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
}

func (r *ActivateDataFlowTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateDataFlowTaskResponse struct {
	*ksyunhttp.BaseResponse
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
}

func (r *DeleteDataFlowStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDataFlowStrategyResponse struct {
	*ksyunhttp.BaseResponse
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
}

func (r *DescribeDataFlowStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDataFlowStrategyListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeDataFlowStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataFlowStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

