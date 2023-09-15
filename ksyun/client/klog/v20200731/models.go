package v20200731
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateProjectRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    Description *string `json:"Description,omitempty" name:"Description"`
    Region *string `json:"Region,omitempty" name:"Region"`
    IamProjectId *int `json:"IamProjectId,omitempty" name:"IamProjectId"`
    IamProjectName *string `json:"IamProjectName,omitempty" name:"IamProjectName"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

type DescribeProjectRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeProjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResponse struct {
    *ksyunhttp.BaseResponse
    CreateTime *string `json:"CreateTime" name:"CreateTime"`
    Description *string `json:"Description" name:"Description"`
    UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
    ProjectName *string `json:"ProjectName" name:"ProjectName"`
    Region *string `json:"Region" name:"Region"`
    IamProjectName *string `json:"IamProjectName" name:"IamProjectName"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    Description *string `json:"Description,omitempty" name:"Description"`
    IamProjectId *int `json:"IamProjectId,omitempty" name:"IamProjectId"`
}

func (r *UpdateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateProjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateProjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

func (r *ListProjectsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListProjectsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListProjectsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
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
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    RetentionDays *int `json:"RetentionDays,omitempty" name:"RetentionDays"`
    Partitions *int `json:"Partitions,omitempty" name:"Partitions"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateLogPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLogPoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateLogPoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeLogPoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLogPoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    RetentionDays *int `json:"RetentionDays,omitempty" name:"RetentionDays"`
    Partitions *int `json:"Partitions,omitempty" name:"Partitions"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateLogPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLogPoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateLogPoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteLogPoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteLogPoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
}

func (r *ListLogPoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListLogPoolsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListLogPoolsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListLogPoolsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *ListLogPoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListLogPoolsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutLogsRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    Time *string `json:"Time,omitempty" name:"Time"`
    Contents *string `json:"Contents,omitempty" name:"Contents"`
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Logs *string `json:"Logs,omitempty" name:"Logs"`
    Filename *string `json:"Filename,omitempty" name:"Filename"`
    Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *PutLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutLogsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "PutLogsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type PutLogsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *PutLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLogsRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    From *int `json:"From,omitempty" name:"From"`
    To *int `json:"To,omitempty" name:"To"`
    Query *string `json:"Query,omitempty" name:"Query"`
    LogPoolId *string `json:"LogPoolId,omitempty" name:"LogPoolId"`
    HitsOpen *bool `json:"HitsOpen,omitempty" name:"HitsOpen"`
    Interval *string `json:"Interval,omitempty" name:"Interval"`
    SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Size *int `json:"Size,omitempty" name:"Size"`
}

func (r *GetLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLogsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetLogsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetLogsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
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
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    QuickSearchName *string `json:"QuickSearchName,omitempty" name:"QuickSearchName"`
    Description *string `json:"Description,omitempty" name:"Description"`
    TimeRange *string `json:"TimeRange,omitempty" name:"TimeRange"`
    Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *CreateQuickSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateQuickSearchRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateQuickSearchRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    Filter *string `json:"Filter,omitempty" name:"Filter"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
}

func (r *ListQuickSearchsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListQuickSearchsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListQuickSearchsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

type GetQuickSearchRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    QuickSearchName *string `json:"QuickSearchName,omitempty" name:"QuickSearchName"`
}

func (r *GetQuickSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetQuickSearchRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetQuickSearchRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetQuickSearchResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *GetQuickSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetQuickSearchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteQuickSearchsRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    QuickSearchName *string `json:"QuickSearchName,omitempty" name:"QuickSearchName"`
}

func (r *DeleteQuickSearchsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteQuickSearchsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteQuickSearchsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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

type CreateDashboardRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
}

func (r *CreateDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDashboardRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDashboardRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *CreateDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardRequest struct {
    *ksyunhttp.BaseRequest
    DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

func (r *DeleteDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDashboardRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDashboardRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *DeleteDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardRequest struct {
    *ksyunhttp.BaseRequest
    DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

func (r *DescribeDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDashboardRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDashboardRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *DescribeDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListDashboardsRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
}

func (r *ListDashboardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListDashboardsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListDashboardsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListDashboardsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *ListDashboardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListDashboardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateChartRequest struct {
    *ksyunhttp.BaseRequest
    DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
    ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
    ChartType *string `json:"ChartType,omitempty" name:"ChartType"`
    Search *string `json:"Search,omitempty" name:"Search"`
    Display *string `json:"Display,omitempty" name:"Display"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    TimeRange *string `json:"TimeRange,omitempty" name:"TimeRange"`
    Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *CreateChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChartRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateChartRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateChartResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *CreateChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChartResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteChartRequest struct {
    *ksyunhttp.BaseRequest
    ChartId *string `json:"ChartId,omitempty" name:"ChartId"`
}

func (r *DeleteChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteChartRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteChartRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteChartResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *DeleteChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteChartResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDashboardNamesByIdsRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    DashboardIds *string `json:"DashboardIds,omitempty" name:"DashboardIds"`
}

func (r *GetDashboardNamesByIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDashboardNamesByIdsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDashboardNamesByIdsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetDashboardNamesByIdsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *GetDashboardNamesByIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDashboardNamesByIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChartNamesByIdsRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    ChartIds *string `json:"ChartIds,omitempty" name:"ChartIds"`
}

func (r *GetChartNamesByIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChartNamesByIdsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetChartNamesByIdsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetChartNamesByIdsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *GetChartNamesByIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChartNamesByIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDashboardRequest struct {
    *ksyunhttp.BaseRequest
    DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
    DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
}

func (r *UpdateDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDashboardRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateDashboardRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDashboardResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *UpdateDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUsageRequest struct {
    *ksyunhttp.BaseRequest
    Projects *string `json:"Projects,omitempty" name:"Projects"`
    Metrics *string `json:"Metrics,omitempty" name:"Metrics"`
    From *string `json:"From,omitempty" name:"From"`
    To *string `json:"To,omitempty" name:"To"`
}

func (r *GetUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUsageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUsageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetUsageResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *GetUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUsageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetIndexTemplateRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    IndexStatus *bool `json:"IndexStatus,omitempty" name:"IndexStatus"`
    FullTextIndex *string `json:"FullTextIndex,omitempty" name:"FullTextIndex"`
    IndexFields *string `json:"IndexFields,omitempty" name:"IndexFields"`
    Lowercase *bool `json:"Lowercase,omitempty" name:"Lowercase"`
    Separator *string `json:"Separator,omitempty" name:"Separator"`
    FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
    FieldType *string `json:"FieldType,omitempty" name:"FieldType"`
    FieldAlias *string `json:"FieldAlias,omitempty" name:"FieldAlias"`
    SeparatorStatus *bool `json:"SeparatorStatus,omitempty" name:"SeparatorStatus"`
}

func (r *SetIndexTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetIndexTemplateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetIndexTemplateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type SetIndexTemplateResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *SetIndexTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetIndexTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetIndexTemplateRequest struct {
    *ksyunhttp.BaseRequest
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
}

func (r *GetIndexTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetIndexTemplateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetIndexTemplateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetIndexTemplateResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *GetIndexTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetIndexTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDownloadTaskRequest struct {
    *ksyunhttp.BaseRequest
    From *string `json:"From,omitempty" name:"From"`
    To *string `json:"To,omitempty" name:"To"`
    LogPoolName *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *CreateDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDownloadTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDownloadTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    Page *string `json:"Page,omitempty" name:"Page"`
    Size *string `json:"Size,omitempty" name:"Size"`
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *ListDownloadTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListDownloadTasksRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListDownloadTasksRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListDownloadTasksResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *ListDownloadTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListDownloadTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDownloadUrlsRequest struct {
    *ksyunhttp.BaseRequest
    DownloadID *string `json:"DownloadID,omitempty" name:"DownloadID"`
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *GetDownloadUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDownloadUrlsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDownloadUrlsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetDownloadUrlsResponse struct {
    *ksyunhttp.BaseResponse
    Res *string `json:"Res" name:"Res"`
}

func (r *GetDownloadUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDownloadUrlsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetMonitorDataRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogErrorReasonRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeLogErrorReasonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogErrorReasonRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLogErrorReasonRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogErrorReasonResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *DescribeLogErrorReasonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogErrorReasonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEtlTaskRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DeleteEtlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEtlTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteEtlTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEtlTaskResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *DeleteEtlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEtlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopEtlTaskRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *StopEtlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopEtlTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StopEtlTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StopEtlTaskResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *StopEtlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopEtlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartEtlTaskRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *StartEtlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartEtlTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartEtlTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StartEtlTaskResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *StartEtlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartEtlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEtlTasksRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *ListEtlTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEtlTasksRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListEtlTasksRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListEtlTasksResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *ListEtlTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEtlTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEtlTaskRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeEtlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEtlTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEtlTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEtlTaskResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *DescribeEtlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEtlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEtlTaskRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *ModifyEtlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEtlTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyEtlTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEtlTaskResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *ModifyEtlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEtlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEtlTaskRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *CreateEtlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEtlTaskRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateEtlTaskRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateEtlTaskResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *CreateEtlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEtlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEtlExceptionRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeEtlExceptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEtlExceptionRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEtlExceptionRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEtlExceptionResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *DescribeEtlExceptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEtlExceptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEtlStatsRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeEtlStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEtlStatsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEtlStatsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEtlStatsResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *DescribeEtlStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEtlStatsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExecuteEtlDemoRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *ExecuteEtlDemoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteEtlDemoRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ExecuteEtlDemoRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ExecuteEtlDemoResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *ExecuteEtlDemoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteEtlDemoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserRegionRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetUserRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserRegionRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUserRegionRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetUserRegionResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *GetUserRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClustersByTypeRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetClustersByTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClustersByTypeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetClustersByTypeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetClustersByTypeResponse struct {
    *ksyunhttp.BaseResponse
    Action *string `json:"Action" name:"Action"`
}

func (r *GetClustersByTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClustersByTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

