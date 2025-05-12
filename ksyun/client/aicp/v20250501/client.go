package v20250501

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "V1"

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

func NewGetImageRequest() (request *GetImageRequest) {
	request = &GetImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetImage")
	return
}

func NewGetImageResponse() (response *GetImageResponse) {
	response = &GetImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetImage(request *GetImageRequest) string {
	return c.GetImageWithContext(context.Background(), request)
}

func (c *Client) GetImageWithContext(ctx context.Context, request *GetImageRequest) string {
	if request == nil {
		request = NewGetImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListImagesRequest() (request *ListImagesRequest) {
	request = &ListImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListImages")
	return
}

func NewListImagesResponse() (response *ListImagesResponse) {
	response = &ListImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListImages(request *ListImagesRequest) string {
	return c.ListImagesWithContext(context.Background(), request)
}

func (c *Client) ListImagesWithContext(ctx context.Context, request *ListImagesRequest) string {
	if request == nil {
		request = NewListImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteImage(request *DeleteImageRequest) string {
	return c.DeleteImageWithContext(context.Background(), request)
}

func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) string {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImage(request *CreateImageRequest) string {
	return c.CreateImageWithContext(context.Background(), request)
}

func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) string {
	if request == nil {
		request = NewCreateImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListNodeMetricsRequest() (request *ListNodeMetricsRequest) {
	request = &ListNodeMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListNodeMetrics")
	return
}

func NewListNodeMetricsResponse() (response *ListNodeMetricsResponse) {
	response = &ListNodeMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListNodeMetrics(request *ListNodeMetricsRequest) string {
	return c.ListNodeMetricsWithContext(context.Background(), request)
}

func (c *Client) ListNodeMetricsWithContext(ctx context.Context, request *ListNodeMetricsRequest) string {
	if request == nil {
		request = NewListNodeMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListNodeMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListStorageConfigsRequest() (request *ListStorageConfigsRequest) {
	request = &ListStorageConfigsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListStorageConfigs")
	return
}

func NewListStorageConfigsResponse() (response *ListStorageConfigsResponse) {
	response = &ListStorageConfigsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListStorageConfigs(request *ListStorageConfigsRequest) string {
	return c.ListStorageConfigsWithContext(context.Background(), request)
}

func (c *Client) ListStorageConfigsWithContext(ctx context.Context, request *ListStorageConfigsRequest) string {
	if request == nil {
		request = NewListStorageConfigsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListStorageConfigsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListSpaceStoragesRequest() (request *ListSpaceStoragesRequest) {
	request = &ListSpaceStoragesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListSpaceStorages")
	return
}

func NewListSpaceStoragesResponse() (response *ListSpaceStoragesResponse) {
	response = &ListSpaceStoragesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSpaceStorages(request *ListSpaceStoragesRequest) string {
	return c.ListSpaceStoragesWithContext(context.Background(), request)
}

func (c *Client) ListSpaceStoragesWithContext(ctx context.Context, request *ListSpaceStoragesRequest) string {
	if request == nil {
		request = NewListSpaceStoragesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListSpaceStoragesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListWorkspacesRequest() (request *ListWorkspacesRequest) {
	request = &ListWorkspacesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListWorkspaces")
	return
}

func NewListWorkspacesResponse() (response *ListWorkspacesResponse) {
	response = &ListWorkspacesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListWorkspaces(request *ListWorkspacesRequest) string {
	return c.ListWorkspacesWithContext(context.Background(), request)
}

func (c *Client) ListWorkspacesWithContext(ctx context.Context, request *ListWorkspacesRequest) string {
	if request == nil {
		request = NewListWorkspacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListWorkspacesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListGpuTypesRequest() (request *ListGpuTypesRequest) {
	request = &ListGpuTypesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListGpuTypes")
	return
}

func NewListGpuTypesResponse() (response *ListGpuTypesResponse) {
	response = &ListGpuTypesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListGpuTypes(request *ListGpuTypesRequest) string {
	return c.ListGpuTypesWithContext(context.Background(), request)
}

func (c *Client) ListGpuTypesWithContext(ctx context.Context, request *ListGpuTypesRequest) string {
	if request == nil {
		request = NewListGpuTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListGpuTypesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateSpaceStorageRequest() (request *CreateSpaceStorageRequest) {
	request = &CreateSpaceStorageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateSpaceStorage")
	return
}

func NewCreateSpaceStorageResponse() (response *CreateSpaceStorageResponse) {
	response = &CreateSpaceStorageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSpaceStorage(request *CreateSpaceStorageRequest) string {
	return c.CreateSpaceStorageWithContext(context.Background(), request)
}

func (c *Client) CreateSpaceStorageWithContext(ctx context.Context, request *CreateSpaceStorageRequest) string {
	if request == nil {
		request = NewCreateSpaceStorageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSpaceStorageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetWorkspaceQuotaRequest() (request *GetWorkspaceQuotaRequest) {
	request = &GetWorkspaceQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetWorkspaceQuota")
	return
}

func NewGetWorkspaceQuotaResponse() (response *GetWorkspaceQuotaResponse) {
	response = &GetWorkspaceQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetWorkspaceQuota(request *GetWorkspaceQuotaRequest) string {
	return c.GetWorkspaceQuotaWithContext(context.Background(), request)
}

func (c *Client) GetWorkspaceQuotaWithContext(ctx context.Context, request *GetWorkspaceQuotaRequest) string {
	if request == nil {
		request = NewGetWorkspaceQuotaRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetWorkspaceQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateJobRequest() (request *CreateJobRequest) {
	request = &CreateJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateJob")
	return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
	response = &CreateJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateJob(request *CreateJobRequest) string {
	return c.CreateJobWithContext(context.Background(), request)
}

func (c *Client) CreateJobWithContext(ctx context.Context, request *CreateJobRequest) string {
	if request == nil {
		request = NewCreateJobRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListSpaceStorageBindedSpacesRequest() (request *ListSpaceStorageBindedSpacesRequest) {
	request = &ListSpaceStorageBindedSpacesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListSpaceStorageBindedSpaces")
	return
}

func NewListSpaceStorageBindedSpacesResponse() (response *ListSpaceStorageBindedSpacesResponse) {
	response = &ListSpaceStorageBindedSpacesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSpaceStorageBindedSpaces(request *ListSpaceStorageBindedSpacesRequest) string {
	return c.ListSpaceStorageBindedSpacesWithContext(context.Background(), request)
}

func (c *Client) ListSpaceStorageBindedSpacesWithContext(ctx context.Context, request *ListSpaceStorageBindedSpacesRequest) string {
	if request == nil {
		request = NewListSpaceStorageBindedSpacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSpaceStorageBindedSpacesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSpaceStorageBindSpaceRequest() (request *SpaceStorageBindSpaceRequest) {
	request = &SpaceStorageBindSpaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "SpaceStorageBindSpace")
	return
}

func NewSpaceStorageBindSpaceResponse() (response *SpaceStorageBindSpaceResponse) {
	response = &SpaceStorageBindSpaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SpaceStorageBindSpace(request *SpaceStorageBindSpaceRequest) string {
	return c.SpaceStorageBindSpaceWithContext(context.Background(), request)
}

func (c *Client) SpaceStorageBindSpaceWithContext(ctx context.Context, request *SpaceStorageBindSpaceRequest) string {
	if request == nil {
		request = NewSpaceStorageBindSpaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSpaceStorageBindSpaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSpaceStorageUnbindSpaceRequest() (request *SpaceStorageUnbindSpaceRequest) {
	request = &SpaceStorageUnbindSpaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "SpaceStorageUnbindSpace")
	return
}

func NewSpaceStorageUnbindSpaceResponse() (response *SpaceStorageUnbindSpaceResponse) {
	response = &SpaceStorageUnbindSpaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SpaceStorageUnbindSpace(request *SpaceStorageUnbindSpaceRequest) string {
	return c.SpaceStorageUnbindSpaceWithContext(context.Background(), request)
}

func (c *Client) SpaceStorageUnbindSpaceWithContext(ctx context.Context, request *SpaceStorageUnbindSpaceRequest) string {
	if request == nil {
		request = NewSpaceStorageUnbindSpaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSpaceStorageUnbindSpaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteJobRequest() (request *DeleteJobRequest) {
	request = &DeleteJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteJob")
	return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
	response = &DeleteJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteJob(request *DeleteJobRequest) string {
	return c.DeleteJobWithContext(context.Background(), request)
}

func (c *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest) string {
	if request == nil {
		request = NewDeleteJobRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetJobRequest() (request *GetJobRequest) {
	request = &GetJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetJob")
	return
}

func NewGetJobResponse() (response *GetJobResponse) {
	response = &GetJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetJob(request *GetJobRequest) string {
	return c.GetJobWithContext(context.Background(), request)
}

func (c *Client) GetJobWithContext(ctx context.Context, request *GetJobRequest) string {
	if request == nil {
		request = NewGetJobRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetJobEventsRequest() (request *GetJobEventsRequest) {
	request = &GetJobEventsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetJobEvents")
	return
}

func NewGetJobEventsResponse() (response *GetJobEventsResponse) {
	response = &GetJobEventsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetJobEvents(request *GetJobEventsRequest) string {
	return c.GetJobEventsWithContext(context.Background(), request)
}

func (c *Client) GetJobEventsWithContext(ctx context.Context, request *GetJobEventsRequest) string {
	if request == nil {
		request = NewGetJobEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetJobEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetPodEventsRequest() (request *GetPodEventsRequest) {
	request = &GetPodEventsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetPodEvents")
	return
}

func NewGetPodEventsResponse() (response *GetPodEventsResponse) {
	response = &GetPodEventsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPodEvents(request *GetPodEventsRequest) string {
	return c.GetPodEventsWithContext(context.Background(), request)
}

func (c *Client) GetPodEventsWithContext(ctx context.Context, request *GetPodEventsRequest) string {
	if request == nil {
		request = NewGetPodEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetPodEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetWebTerminalRequest() (request *GetWebTerminalRequest) {
	request = &GetWebTerminalRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetWebTerminal")
	return
}

func NewGetWebTerminalResponse() (response *GetWebTerminalResponse) {
	response = &GetWebTerminalResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetWebTerminal(request *GetWebTerminalRequest) string {
	return c.GetWebTerminalWithContext(context.Background(), request)
}

func (c *Client) GetWebTerminalWithContext(ctx context.Context, request *GetWebTerminalRequest) string {
	if request == nil {
		request = NewGetWebTerminalRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetWebTerminalResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetJobMetricsRequest() (request *GetJobMetricsRequest) {
	request = &GetJobMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetJobMetrics")
	return
}

func NewGetJobMetricsResponse() (response *GetJobMetricsResponse) {
	response = &GetJobMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetJobMetrics(request *GetJobMetricsRequest) string {
	return c.GetJobMetricsWithContext(context.Background(), request)
}

func (c *Client) GetJobMetricsWithContext(ctx context.Context, request *GetJobMetricsRequest) string {
	if request == nil {
		request = NewGetJobMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetJobMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListJobsRequest() (request *ListJobsRequest) {
	request = &ListJobsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListJobs")
	return
}

func NewListJobsResponse() (response *ListJobsResponse) {
	response = &ListJobsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListJobs(request *ListJobsRequest) string {
	return c.ListJobsWithContext(context.Background(), request)
}

func (c *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest) string {
	if request == nil {
		request = NewListJobsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListJobsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetPodLogsRequest() (request *GetPodLogsRequest) {
	request = &GetPodLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetPodLogs")
	return
}

func NewGetPodLogsResponse() (response *GetPodLogsResponse) {
	response = &GetPodLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetPodLogs(request *GetPodLogsRequest) string {
	return c.GetPodLogsWithContext(context.Background(), request)
}

func (c *Client) GetPodLogsWithContext(ctx context.Context, request *GetPodLogsRequest) string {
	if request == nil {
		request = NewGetPodLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetPodLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStopJobRequest() (request *StopJobRequest) {
	request = &StopJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopJob")
	return
}

func NewStopJobResponse() (response *StopJobResponse) {
	response = &StopJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopJob(request *StopJobRequest) string {
	return c.StopJobWithContext(context.Background(), request)
}

func (c *Client) StopJobWithContext(ctx context.Context, request *StopJobRequest) string {
	if request == nil {
		request = NewStopJobRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
