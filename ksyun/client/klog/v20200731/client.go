package v20200731

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-07-31"

type Client struct {
	common.Client
}

func NewClient(credential common.Credentials, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
	request = &CreateProjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateProject")
	return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
	response = &CreateProjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateProject(request *CreateProjectRequest) string {
	return c.CreateProjectWithContext(context.Background(), request)
}

func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) string {
	if request == nil {
		request = NewCreateProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateProjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
	request = &DescribeProjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeProject")
	return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
	response = &DescribeProjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeProject(request *DescribeProjectRequest) string {
	return c.DescribeProjectWithContext(context.Background(), request)
}

func (c *Client) DescribeProjectWithContext(ctx context.Context, request *DescribeProjectRequest) string {
	if request == nil {
		request = NewDescribeProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeProjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateProjectRequest() (request *UpdateProjectRequest) {
	request = &UpdateProjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "UpdateProject")
	return
}

func NewUpdateProjectResponse() (response *UpdateProjectResponse) {
	response = &UpdateProjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateProject(request *UpdateProjectRequest) string {
	return c.UpdateProjectWithContext(context.Background(), request)
}

func (c *Client) UpdateProjectWithContext(ctx context.Context, request *UpdateProjectRequest) string {
	if request == nil {
		request = NewUpdateProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateProjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
	request = &DeleteProjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DeleteProject")
	return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
	response = &DeleteProjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteProject(request *DeleteProjectRequest) string {
	return c.DeleteProjectWithContext(context.Background(), request)
}

func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) string {
	if request == nil {
		request = NewDeleteProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteProjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListProjectsRequest() (request *ListProjectsRequest) {
	request = &ListProjectsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ListProjects")
	return
}

func NewListProjectsResponse() (response *ListProjectsResponse) {
	response = &ListProjectsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListProjects(request *ListProjectsRequest) string {
	return c.ListProjectsWithContext(context.Background(), request)
}

func (c *Client) ListProjectsWithContext(ctx context.Context, request *ListProjectsRequest) string {
	if request == nil {
		request = NewListProjectsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListProjectsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateLogPoolRequest() (request *CreateLogPoolRequest) {
	request = &CreateLogPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateLogPool")
	return
}

func NewCreateLogPoolResponse() (response *CreateLogPoolResponse) {
	response = &CreateLogPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLogPool(request *CreateLogPoolRequest) string {
	return c.CreateLogPoolWithContext(context.Background(), request)
}

func (c *Client) CreateLogPoolWithContext(ctx context.Context, request *CreateLogPoolRequest) string {
	if request == nil {
		request = NewCreateLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateLogPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeLogPoolRequest() (request *DescribeLogPoolRequest) {
	request = &DescribeLogPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeLogPool")
	return
}

func NewDescribeLogPoolResponse() (response *DescribeLogPoolResponse) {
	response = &DescribeLogPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLogPool(request *DescribeLogPoolRequest) string {
	return c.DescribeLogPoolWithContext(context.Background(), request)
}

func (c *Client) DescribeLogPoolWithContext(ctx context.Context, request *DescribeLogPoolRequest) string {
	if request == nil {
		request = NewDescribeLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeLogPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateLogPoolRequest() (request *UpdateLogPoolRequest) {
	request = &UpdateLogPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "UpdateLogPool")
	return
}

func NewUpdateLogPoolResponse() (response *UpdateLogPoolResponse) {
	response = &UpdateLogPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateLogPool(request *UpdateLogPoolRequest) string {
	return c.UpdateLogPoolWithContext(context.Background(), request)
}

func (c *Client) UpdateLogPoolWithContext(ctx context.Context, request *UpdateLogPoolRequest) string {
	if request == nil {
		request = NewUpdateLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateLogPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteLogPoolRequest() (request *DeleteLogPoolRequest) {
	request = &DeleteLogPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DeleteLogPool")
	return
}

func NewDeleteLogPoolResponse() (response *DeleteLogPoolResponse) {
	response = &DeleteLogPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLogPool(request *DeleteLogPoolRequest) string {
	return c.DeleteLogPoolWithContext(context.Background(), request)
}

func (c *Client) DeleteLogPoolWithContext(ctx context.Context, request *DeleteLogPoolRequest) string {
	if request == nil {
		request = NewDeleteLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteLogPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListLogPoolsRequest() (request *ListLogPoolsRequest) {
	request = &ListLogPoolsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ListLogPools")
	return
}

func NewListLogPoolsResponse() (response *ListLogPoolsResponse) {
	response = &ListLogPoolsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListLogPools(request *ListLogPoolsRequest) string {
	return c.ListLogPoolsWithContext(context.Background(), request)
}

func (c *Client) ListLogPoolsWithContext(ctx context.Context, request *ListLogPoolsRequest) string {
	if request == nil {
		request = NewListLogPoolsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListLogPoolsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewPutLogsRequest() (request *PutLogsRequest) {
	request = &PutLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "PutLogs")
	return
}

func NewPutLogsResponse() (response *PutLogsResponse) {
	response = &PutLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) PutLogs(request *PutLogsRequest) string {
	return c.PutLogsWithContext(context.Background(), request)
}

func (c *Client) PutLogsWithContext(ctx context.Context, request *PutLogsRequest) string {
	if request == nil {
		request = NewPutLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewPutLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetLogsRequest() (request *GetLogsRequest) {
	request = &GetLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetLogs")
	return
}

func NewGetLogsResponse() (response *GetLogsResponse) {
	response = &GetLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetLogs(request *GetLogsRequest) string {
	return c.GetLogsWithContext(context.Background(), request)
}

func (c *Client) GetLogsWithContext(ctx context.Context, request *GetLogsRequest) string {
	if request == nil {
		request = NewGetLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateQuickSearchRequest() (request *CreateQuickSearchRequest) {
	request = &CreateQuickSearchRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateQuickSearch")
	return
}

func NewCreateQuickSearchResponse() (response *CreateQuickSearchResponse) {
	response = &CreateQuickSearchResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateQuickSearch(request *CreateQuickSearchRequest) string {
	return c.CreateQuickSearchWithContext(context.Background(), request)
}

func (c *Client) CreateQuickSearchWithContext(ctx context.Context, request *CreateQuickSearchRequest) string {
	if request == nil {
		request = NewCreateQuickSearchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateQuickSearchResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListQuickSearchsRequest() (request *ListQuickSearchsRequest) {
	request = &ListQuickSearchsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ListQuickSearchs")
	return
}

func NewListQuickSearchsResponse() (response *ListQuickSearchsResponse) {
	response = &ListQuickSearchsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListQuickSearchs(request *ListQuickSearchsRequest) string {
	return c.ListQuickSearchsWithContext(context.Background(), request)
}

func (c *Client) ListQuickSearchsWithContext(ctx context.Context, request *ListQuickSearchsRequest) string {
	if request == nil {
		request = NewListQuickSearchsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListQuickSearchsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetQuickSearchRequest() (request *GetQuickSearchRequest) {
	request = &GetQuickSearchRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetQuickSearch")
	return
}

func NewGetQuickSearchResponse() (response *GetQuickSearchResponse) {
	response = &GetQuickSearchResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetQuickSearch(request *GetQuickSearchRequest) string {
	return c.GetQuickSearchWithContext(context.Background(), request)
}

func (c *Client) GetQuickSearchWithContext(ctx context.Context, request *GetQuickSearchRequest) string {
	if request == nil {
		request = NewGetQuickSearchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetQuickSearchResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteQuickSearchsRequest() (request *DeleteQuickSearchsRequest) {
	request = &DeleteQuickSearchsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DeleteQuickSearchs")
	return
}

func NewDeleteQuickSearchsResponse() (response *DeleteQuickSearchsResponse) {
	response = &DeleteQuickSearchsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteQuickSearchs(request *DeleteQuickSearchsRequest) string {
	return c.DeleteQuickSearchsWithContext(context.Background(), request)
}

func (c *Client) DeleteQuickSearchsWithContext(ctx context.Context, request *DeleteQuickSearchsRequest) string {
	if request == nil {
		request = NewDeleteQuickSearchsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteQuickSearchsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDashboardRequest() (request *CreateDashboardRequest) {
	request = &CreateDashboardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateDashboard")
	return
}

func NewCreateDashboardResponse() (response *CreateDashboardResponse) {
	response = &CreateDashboardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDashboard(request *CreateDashboardRequest) string {
	return c.CreateDashboardWithContext(context.Background(), request)
}

func (c *Client) CreateDashboardWithContext(ctx context.Context, request *CreateDashboardRequest) string {
	if request == nil {
		request = NewCreateDashboardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDashboardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDashboardRequest() (request *DeleteDashboardRequest) {
	request = &DeleteDashboardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DeleteDashboard")
	return
}

func NewDeleteDashboardResponse() (response *DeleteDashboardResponse) {
	response = &DeleteDashboardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDashboard(request *DeleteDashboardRequest) string {
	return c.DeleteDashboardWithContext(context.Background(), request)
}

func (c *Client) DeleteDashboardWithContext(ctx context.Context, request *DeleteDashboardRequest) string {
	if request == nil {
		request = NewDeleteDashboardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDashboardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDashboardRequest() (request *DescribeDashboardRequest) {
	request = &DescribeDashboardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeDashboard")
	return
}

func NewDescribeDashboardResponse() (response *DescribeDashboardResponse) {
	response = &DescribeDashboardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDashboard(request *DescribeDashboardRequest) string {
	return c.DescribeDashboardWithContext(context.Background(), request)
}

func (c *Client) DescribeDashboardWithContext(ctx context.Context, request *DescribeDashboardRequest) string {
	if request == nil {
		request = NewDescribeDashboardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDashboardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListDashboardsRequest() (request *ListDashboardsRequest) {
	request = &ListDashboardsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ListDashboards")
	return
}

func NewListDashboardsResponse() (response *ListDashboardsResponse) {
	response = &ListDashboardsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListDashboards(request *ListDashboardsRequest) string {
	return c.ListDashboardsWithContext(context.Background(), request)
}

func (c *Client) ListDashboardsWithContext(ctx context.Context, request *ListDashboardsRequest) string {
	if request == nil {
		request = NewListDashboardsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListDashboardsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateChartRequest() (request *CreateChartRequest) {
	request = &CreateChartRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateChart")
	return
}

func NewCreateChartResponse() (response *CreateChartResponse) {
	response = &CreateChartResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateChart(request *CreateChartRequest) string {
	return c.CreateChartWithContext(context.Background(), request)
}

func (c *Client) CreateChartWithContext(ctx context.Context, request *CreateChartRequest) string {
	if request == nil {
		request = NewCreateChartRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateChartResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteChartRequest() (request *DeleteChartRequest) {
	request = &DeleteChartRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DeleteChart")
	return
}

func NewDeleteChartResponse() (response *DeleteChartResponse) {
	response = &DeleteChartResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteChart(request *DeleteChartRequest) string {
	return c.DeleteChartWithContext(context.Background(), request)
}

func (c *Client) DeleteChartWithContext(ctx context.Context, request *DeleteChartRequest) string {
	if request == nil {
		request = NewDeleteChartRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteChartResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetDashboardNamesByIdsRequest() (request *GetDashboardNamesByIdsRequest) {
	request = &GetDashboardNamesByIdsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetDashboardNamesByIds")
	return
}

func NewGetDashboardNamesByIdsResponse() (response *GetDashboardNamesByIdsResponse) {
	response = &GetDashboardNamesByIdsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDashboardNamesByIds(request *GetDashboardNamesByIdsRequest) string {
	return c.GetDashboardNamesByIdsWithContext(context.Background(), request)
}

func (c *Client) GetDashboardNamesByIdsWithContext(ctx context.Context, request *GetDashboardNamesByIdsRequest) string {
	if request == nil {
		request = NewGetDashboardNamesByIdsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDashboardNamesByIdsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetChartNamesByIdsRequest() (request *GetChartNamesByIdsRequest) {
	request = &GetChartNamesByIdsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetChartNamesByIds")
	return
}

func NewGetChartNamesByIdsResponse() (response *GetChartNamesByIdsResponse) {
	response = &GetChartNamesByIdsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetChartNamesByIds(request *GetChartNamesByIdsRequest) string {
	return c.GetChartNamesByIdsWithContext(context.Background(), request)
}

func (c *Client) GetChartNamesByIdsWithContext(ctx context.Context, request *GetChartNamesByIdsRequest) string {
	if request == nil {
		request = NewGetChartNamesByIdsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetChartNamesByIdsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateDashboardRequest() (request *UpdateDashboardRequest) {
	request = &UpdateDashboardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "UpdateDashboard")
	return
}

func NewUpdateDashboardResponse() (response *UpdateDashboardResponse) {
	response = &UpdateDashboardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDashboard(request *UpdateDashboardRequest) string {
	return c.UpdateDashboardWithContext(context.Background(), request)
}

func (c *Client) UpdateDashboardWithContext(ctx context.Context, request *UpdateDashboardRequest) string {
	if request == nil {
		request = NewUpdateDashboardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDashboardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetUsageRequest() (request *GetUsageRequest) {
	request = &GetUsageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetUsage")
	return
}

func NewGetUsageResponse() (response *GetUsageResponse) {
	response = &GetUsageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUsage(request *GetUsageRequest) string {
	return c.GetUsageWithContext(context.Background(), request)
}

func (c *Client) GetUsageWithContext(ctx context.Context, request *GetUsageRequest) string {
	if request == nil {
		request = NewGetUsageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetUsageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSetIndexTemplateRequest() (request *SetIndexTemplateRequest) {
	request = &SetIndexTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "SetIndexTemplate")
	return
}

func NewSetIndexTemplateResponse() (response *SetIndexTemplateResponse) {
	response = &SetIndexTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetIndexTemplate(request *SetIndexTemplateRequest) string {
	return c.SetIndexTemplateWithContext(context.Background(), request)
}

func (c *Client) SetIndexTemplateWithContext(ctx context.Context, request *SetIndexTemplateRequest) string {
	if request == nil {
		request = NewSetIndexTemplateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetIndexTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetIndexTemplateRequest() (request *GetIndexTemplateRequest) {
	request = &GetIndexTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetIndexTemplate")
	return
}

func NewGetIndexTemplateResponse() (response *GetIndexTemplateResponse) {
	response = &GetIndexTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetIndexTemplate(request *GetIndexTemplateRequest) string {
	return c.GetIndexTemplateWithContext(context.Background(), request)
}

func (c *Client) GetIndexTemplateWithContext(ctx context.Context, request *GetIndexTemplateRequest) string {
	if request == nil {
		request = NewGetIndexTemplateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetIndexTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDownloadTaskRequest() (request *CreateDownloadTaskRequest) {
	request = &CreateDownloadTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateDownloadTask")
	return
}

func NewCreateDownloadTaskResponse() (response *CreateDownloadTaskResponse) {
	response = &CreateDownloadTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDownloadTask(request *CreateDownloadTaskRequest) string {
	return c.CreateDownloadTaskWithContext(context.Background(), request)
}

func (c *Client) CreateDownloadTaskWithContext(ctx context.Context, request *CreateDownloadTaskRequest) string {
	if request == nil {
		request = NewCreateDownloadTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDownloadTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListDownloadTasksRequest() (request *ListDownloadTasksRequest) {
	request = &ListDownloadTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ListDownloadTasks")
	return
}

func NewListDownloadTasksResponse() (response *ListDownloadTasksResponse) {
	response = &ListDownloadTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListDownloadTasks(request *ListDownloadTasksRequest) string {
	return c.ListDownloadTasksWithContext(context.Background(), request)
}

func (c *Client) ListDownloadTasksWithContext(ctx context.Context, request *ListDownloadTasksRequest) string {
	if request == nil {
		request = NewListDownloadTasksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListDownloadTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetDownloadUrlsRequest() (request *GetDownloadUrlsRequest) {
	request = &GetDownloadUrlsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetDownloadUrls")
	return
}

func NewGetDownloadUrlsResponse() (response *GetDownloadUrlsResponse) {
	response = &GetDownloadUrlsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDownloadUrls(request *GetDownloadUrlsRequest) string {
	return c.GetDownloadUrlsWithContext(context.Background(), request)
}

func (c *Client) GetDownloadUrlsWithContext(ctx context.Context, request *GetDownloadUrlsRequest) string {
	if request == nil {
		request = NewGetDownloadUrlsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDownloadUrlsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
	request = &GetMonitorDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetMonitorData")
	return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
	response = &GetMonitorDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMonitorData(request *GetMonitorDataRequest) string {
	return c.GetMonitorDataWithContext(context.Background(), request)
}

func (c *Client) GetMonitorDataWithContext(ctx context.Context, request *GetMonitorDataRequest) string {
	if request == nil {
		request = NewGetMonitorDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMonitorDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeLogErrorReasonRequest() (request *DescribeLogErrorReasonRequest) {
	request = &DescribeLogErrorReasonRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeLogErrorReason")
	return
}

func NewDescribeLogErrorReasonResponse() (response *DescribeLogErrorReasonResponse) {
	response = &DescribeLogErrorReasonResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLogErrorReason(request *DescribeLogErrorReasonRequest) string {
	return c.DescribeLogErrorReasonWithContext(context.Background(), request)
}

func (c *Client) DescribeLogErrorReasonWithContext(ctx context.Context, request *DescribeLogErrorReasonRequest) string {
	if request == nil {
		request = NewDescribeLogErrorReasonRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeLogErrorReasonResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteEtlTaskRequest() (request *DeleteEtlTaskRequest) {
	request = &DeleteEtlTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DeleteEtlTask")
	return
}

func NewDeleteEtlTaskResponse() (response *DeleteEtlTaskResponse) {
	response = &DeleteEtlTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteEtlTask(request *DeleteEtlTaskRequest) string {
	return c.DeleteEtlTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteEtlTaskWithContext(ctx context.Context, request *DeleteEtlTaskRequest) string {
	if request == nil {
		request = NewDeleteEtlTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEtlTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStopEtlTaskRequest() (request *StopEtlTaskRequest) {
	request = &StopEtlTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "StopEtlTask")
	return
}

func NewStopEtlTaskResponse() (response *StopEtlTaskResponse) {
	response = &StopEtlTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopEtlTask(request *StopEtlTaskRequest) string {
	return c.StopEtlTaskWithContext(context.Background(), request)
}

func (c *Client) StopEtlTaskWithContext(ctx context.Context, request *StopEtlTaskRequest) string {
	if request == nil {
		request = NewStopEtlTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopEtlTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStartEtlTaskRequest() (request *StartEtlTaskRequest) {
	request = &StartEtlTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "StartEtlTask")
	return
}

func NewStartEtlTaskResponse() (response *StartEtlTaskResponse) {
	response = &StartEtlTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartEtlTask(request *StartEtlTaskRequest) string {
	return c.StartEtlTaskWithContext(context.Background(), request)
}

func (c *Client) StartEtlTaskWithContext(ctx context.Context, request *StartEtlTaskRequest) string {
	if request == nil {
		request = NewStartEtlTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartEtlTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListEtlTasksRequest() (request *ListEtlTasksRequest) {
	request = &ListEtlTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ListEtlTasks")
	return
}

func NewListEtlTasksResponse() (response *ListEtlTasksResponse) {
	response = &ListEtlTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListEtlTasks(request *ListEtlTasksRequest) string {
	return c.ListEtlTasksWithContext(context.Background(), request)
}

func (c *Client) ListEtlTasksWithContext(ctx context.Context, request *ListEtlTasksRequest) string {
	if request == nil {
		request = NewListEtlTasksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListEtlTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeEtlTaskRequest() (request *DescribeEtlTaskRequest) {
	request = &DescribeEtlTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeEtlTask")
	return
}

func NewDescribeEtlTaskResponse() (response *DescribeEtlTaskResponse) {
	response = &DescribeEtlTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEtlTask(request *DescribeEtlTaskRequest) string {
	return c.DescribeEtlTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeEtlTaskWithContext(ctx context.Context, request *DescribeEtlTaskRequest) string {
	if request == nil {
		request = NewDescribeEtlTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEtlTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyEtlTaskRequest() (request *ModifyEtlTaskRequest) {
	request = &ModifyEtlTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ModifyEtlTask")
	return
}

func NewModifyEtlTaskResponse() (response *ModifyEtlTaskResponse) {
	response = &ModifyEtlTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyEtlTask(request *ModifyEtlTaskRequest) string {
	return c.ModifyEtlTaskWithContext(context.Background(), request)
}

func (c *Client) ModifyEtlTaskWithContext(ctx context.Context, request *ModifyEtlTaskRequest) string {
	if request == nil {
		request = NewModifyEtlTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEtlTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateEtlTaskRequest() (request *CreateEtlTaskRequest) {
	request = &CreateEtlTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "CreateEtlTask")
	return
}

func NewCreateEtlTaskResponse() (response *CreateEtlTaskResponse) {
	response = &CreateEtlTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateEtlTask(request *CreateEtlTaskRequest) string {
	return c.CreateEtlTaskWithContext(context.Background(), request)
}

func (c *Client) CreateEtlTaskWithContext(ctx context.Context, request *CreateEtlTaskRequest) string {
	if request == nil {
		request = NewCreateEtlTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEtlTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeEtlExceptionRequest() (request *DescribeEtlExceptionRequest) {
	request = &DescribeEtlExceptionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeEtlException")
	return
}

func NewDescribeEtlExceptionResponse() (response *DescribeEtlExceptionResponse) {
	response = &DescribeEtlExceptionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEtlException(request *DescribeEtlExceptionRequest) string {
	return c.DescribeEtlExceptionWithContext(context.Background(), request)
}

func (c *Client) DescribeEtlExceptionWithContext(ctx context.Context, request *DescribeEtlExceptionRequest) string {
	if request == nil {
		request = NewDescribeEtlExceptionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEtlExceptionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeEtlStatsRequest() (request *DescribeEtlStatsRequest) {
	request = &DescribeEtlStatsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "DescribeEtlStats")
	return
}

func NewDescribeEtlStatsResponse() (response *DescribeEtlStatsResponse) {
	response = &DescribeEtlStatsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEtlStats(request *DescribeEtlStatsRequest) string {
	return c.DescribeEtlStatsWithContext(context.Background(), request)
}

func (c *Client) DescribeEtlStatsWithContext(ctx context.Context, request *DescribeEtlStatsRequest) string {
	if request == nil {
		request = NewDescribeEtlStatsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEtlStatsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewExecuteEtlDemoRequest() (request *ExecuteEtlDemoRequest) {
	request = &ExecuteEtlDemoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "ExecuteEtlDemo")
	return
}

func NewExecuteEtlDemoResponse() (response *ExecuteEtlDemoResponse) {
	response = &ExecuteEtlDemoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ExecuteEtlDemo(request *ExecuteEtlDemoRequest) string {
	return c.ExecuteEtlDemoWithContext(context.Background(), request)
}

func (c *Client) ExecuteEtlDemoWithContext(ctx context.Context, request *ExecuteEtlDemoRequest) string {
	if request == nil {
		request = NewExecuteEtlDemoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewExecuteEtlDemoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetUserRegionRequest() (request *GetUserRegionRequest) {
	request = &GetUserRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetUserRegion")
	return
}

func NewGetUserRegionResponse() (response *GetUserRegionResponse) {
	response = &GetUserRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUserRegion(request *GetUserRegionRequest) string {
	return c.GetUserRegionWithContext(context.Background(), request)
}

func (c *Client) GetUserRegionWithContext(ctx context.Context, request *GetUserRegionRequest) string {
	if request == nil {
		request = NewGetUserRegionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUserRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetClustersByTypeRequest() (request *GetClustersByTypeRequest) {
	request = &GetClustersByTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("klog", APIVersion, "GetClustersByType")
	return
}

func NewGetClustersByTypeResponse() (response *GetClustersByTypeResponse) {
	response = &GetClustersByTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetClustersByType(request *GetClustersByTypeRequest) string {
	return c.GetClustersByTypeWithContext(context.Background(), request)
}

func (c *Client) GetClustersByTypeWithContext(ctx context.Context, request *GetClustersByTypeRequest) string {
	if request == nil {
		request = NewGetClustersByTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetClustersByTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
