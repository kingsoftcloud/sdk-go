package v20200731

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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
