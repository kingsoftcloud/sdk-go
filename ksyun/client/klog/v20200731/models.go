package v20200731

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type UpdateLogPoolConfigColumns struct {
	Id      *string `json:"Id,omitempty" name:"Id"`
	Name    *string `json:"Name,omitempty" name:"Name"`
	Type    *string `json:"Type,omitempty" name:"Type"`
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}
type UpdateLogPoolConfig struct {
	Columns []*UpdateLogPoolConfigColumns `json:"Columns,omitempty" name:"Columns"`
}
type ListLogPoolsTags struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateProjectRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName  *string `json:"ProjectName,omitempty" name:"ProjectName"`
	Description  *string `json:"Description,omitempty" name:"Description"`
	IamProjectId *int    `json:"IamProjectId,omitempty" name:"IamProjectId"`
}

func (r *CreateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateProjectResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *CreateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName  *string `json:"ProjectName,omitempty" name:"ProjectName"`
	Description  *string `json:"Description,omitempty" name:"Description"`
	IamProjectId *int    `json:"IamProjectId,omitempty" name:"IamProjectId"`
}

func (r *UpdateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateProjectResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *UpdateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DeleteProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteProjectResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectsRequest struct {
	*ksyunhttp.BaseRequest
	Page *int `json:"Page,omitempty" name:"Page"`
	Size *int `json:"Size,omitempty" name:"Size"`
}

func (r *ListProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListProjectsResponse struct {
	*ksyunhttp.BaseResponse
	Total    *int `json:"Total" name:"Total"`
	Count    *int `json:"Count" name:"Count"`
	Projects []struct {
		ProjectName    *string `json:"ProjectName" name:"ProjectName"`
		IamProjectName *string `json:"IamProjectName" name:"IamProjectName"`
		IamProjectId   *int    `json:"IamProjectId" name:"IamProjectId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime     *string `json:"UpdateTime" name:"UpdateTime"`
		Region         *string `json:"Region" name:"Region"`
		Status         *string `json:"Status" name:"Status"`
		LogPoolNum     *int    `json:"LogPoolNum" name:"LogPoolNum"`
		Tags           []struct {
			ID    *int    `json:"ID" name:"ID"`
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Tags" name:"Tags"`
	} `json:"Projects"`
}

func (r *ListProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogPoolRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName   *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName   *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
	RetentionDays *int    `json:"RetentionDays,omitempty" name:"RetentionDays"`
	Partitions    *int    `json:"Partitions,omitempty" name:"Partitions"`
	Description   *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateLogPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateLogPoolResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *CreateLogPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogPoolRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
}

func (r *DescribeLogPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeLogPoolResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *DescribeLogPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogPoolRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName   *string              `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName   *string              `json:"LogPoolName,omitempty" name:"LogPoolName"`
	RetentionDays *int                 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	Partitions    *int                 `json:"Partitions,omitempty" name:"Partitions"`
	Description   *string              `json:"Description,omitempty" name:"Description"`
	Config        *UpdateLogPoolConfig `json:"Config,omitempty" name:"Config"`
}

func (r *UpdateLogPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateLogPoolResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *UpdateLogPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogPoolRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
}

func (r *DeleteLogPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteLogPoolResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *DeleteLogPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogPoolsRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName *string             `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName *string             `json:"LogPoolName,omitempty" name:"LogPoolName"`
	Page        *int                `json:"Page,omitempty" name:"Page"`
	Size        *int                `json:"Size,omitempty" name:"Size"`
	Tags        []*ListLogPoolsTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *ListLogPoolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListLogPoolsResponse struct {
	*ksyunhttp.BaseResponse
	ProjectName *string `json:"ProjectName" name:"ProjectName"`
	Total       *int    `json:"Total" name:"Total"`
	Count       *int    `json:"Count" name:"Count"`
	LogPools    []struct {
		ProjectName   *string `json:"ProjectName" name:"ProjectName"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		LogPoolName   *string `json:"LogPoolName" name:"LogPoolName"`
		LogPoolId     *string `json:"LogPoolId" name:"LogPoolId"`
		Partitions    *int    `json:"Partitions" name:"Partitions"`
		RetentionDays *int    `json:"RetentionDays" name:"RetentionDays"`
		Tags          []struct {
			Id    *int    `json:"Id" name:"Id"`
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Tags" name:"Tags"`
		WebTracking *bool   `json:"WebTracking" name:"WebTracking"`
		Status      *string `json:"Status" name:"Status"`
		Scene       *string `json:"Scene" name:"Scene"`
	} `json:"LogPools"`
}

func (r *ListLogPoolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogPoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLogsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetLogsResponse struct {
	*ksyunhttp.BaseResponse
	Total  *int  `json:"Total" name:"Total"`
	Count  *int  `json:"Count" name:"Count"`
	HasSql *bool `json:"HasSql" name:"HasSql"`
	Logs   []struct {
		Id        *string `json:"Id" name:"Id"`
		Source    *string `json:"Source" name:"Source"`
		Path      *string `json:"Path" name:"Path"`
		Timestamp *string `json:"Timestamp" name:"Timestamp"`
		Field1    *string `json:"Field1" name:"Field1"`
		Field2    *string `json:"Field2" name:"Field2"`
	} `json:"Logs"`
	Histogram []struct {
		Key      *string `json:"Key" name:"Key"`
		LogCount *int    `json:"LogCount" name:"LogCount"`
	} `json:"Histogram"`
	Keys []*string `json:"Keys" name:"Keys"`
}

func (r *GetLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQuickSearchRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName     *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName     *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
	QuickSearchName *string `json:"QuickSearchName,omitempty" name:"QuickSearchName"`
	Query           *string `json:"Query,omitempty" name:"Query"`
	Description     *string `json:"Description,omitempty" name:"Description"`
	TimeRange       *string `json:"TimeRange,omitempty" name:"TimeRange"`
}

func (r *CreateQuickSearchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateQuickSearchResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *CreateQuickSearchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQuickSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListQuickSearchsRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
	Filter      *string `json:"Filter,omitempty" name:"Filter"`
	Page        *int    `json:"Page,omitempty" name:"Page"`
	Size        *int    `json:"Size,omitempty" name:"Size"`
}

func (r *ListQuickSearchsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListQuickSearchsResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *ListQuickSearchsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListQuickSearchsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQuickSearchsRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName     *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName     *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
	QuickSearchName *string `json:"QuickSearchName,omitempty" name:"QuickSearchName"`
}

func (r *DeleteQuickSearchsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteQuickSearchsResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *DeleteQuickSearchsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQuickSearchsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDownloadTaskRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *CreateDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDownloadTaskResponse struct {
	*ksyunhttp.BaseResponse
	Res *string `json:"Res" name:"Res"`
}

func (r *CreateDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDownloadTasksRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	Page        *string `json:"Page,omitempty" name:"Page"`
	Size        *string `json:"Size,omitempty" name:"Size"`
}

func (r *ListDownloadTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListDownloadTasksResponse struct {
	*ksyunhttp.BaseResponse
	Total *int `json:"Total" name:"Total"`
	Count *int `json:"Count" name:"Count"`
	Data  []struct {
		ProjectName    *string `json:"ProjectName" name:"ProjectName"`
		DownloadTaskId *string `json:"DownloadTaskId" name:"DownloadTaskId"`
		LogPoolNames   *string `json:"LogPoolNames" name:"LogPoolNames"`
		Status         *string `json:"Status" name:"Status"`
		AccountId      *string `json:"AccountId" name:"AccountId"`
		From           *string `json:"From" name:"From"`
		To             *string `json:"To" name:"To"`
		Count          *string `json:"Count" name:"Count"`
		Size           *int    `json:"Size" name:"Size"`
		DownloadUrl    *string `json:"DownloadUrl" name:"DownloadUrl"`
		Query          *string `json:"Query" name:"Query"`
		CreatedAt      *int    `json:"CreatedAt" name:"CreatedAt"`
	} `json:"Data"`
}

func (r *ListDownloadTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDownloadTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
