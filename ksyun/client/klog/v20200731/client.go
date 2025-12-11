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

func (c *Client) CreateProjectSend(request *CreateProjectRequest) (*CreateProjectResponse, error) {
	statusCode, msg, err := c.CreateProjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateProjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateProjectWithContextV2(ctx context.Context, request *CreateProjectRequest) (int, string, error) {
	if request == nil {
		request = NewCreateProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateProjectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeProjectSend(request *DescribeProjectRequest) (*DescribeProjectResponse, error) {
	statusCode, msg, err := c.DescribeProjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeProjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeProjectWithContextV2(ctx context.Context, request *DescribeProjectRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeProjectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateProjectSend(request *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	statusCode, msg, err := c.UpdateProjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateProjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateProjectWithContextV2(ctx context.Context, request *UpdateProjectRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateProjectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteProjectSend(request *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	statusCode, msg, err := c.DeleteProjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteProjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteProjectWithContextV2(ctx context.Context, request *DeleteProjectRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteProjectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListProjectsSend(request *ListProjectsRequest) (*ListProjectsResponse, error) {
	statusCode, msg, err := c.ListProjectsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListProjectsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ListProjectsWithContextV2(ctx context.Context, request *ListProjectsRequest) (int, string, error) {
	if request == nil {
		request = NewListProjectsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListProjectsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateLogPoolSend(request *CreateLogPoolRequest) (*CreateLogPoolResponse, error) {
	statusCode, msg, err := c.CreateLogPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateLogPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateLogPoolWithContextV2(ctx context.Context, request *CreateLogPoolRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateLogPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeLogPoolSend(request *DescribeLogPoolRequest) (*DescribeLogPoolResponse, error) {
	statusCode, msg, err := c.DescribeLogPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeLogPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeLogPoolWithContextV2(ctx context.Context, request *DescribeLogPoolRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeLogPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateLogPoolSend(request *UpdateLogPoolRequest) (*UpdateLogPoolResponse, error) {
	statusCode, msg, err := c.UpdateLogPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateLogPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateLogPoolWithContextV2(ctx context.Context, request *UpdateLogPoolRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateLogPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteLogPoolSend(request *DeleteLogPoolRequest) (*DeleteLogPoolResponse, error) {
	statusCode, msg, err := c.DeleteLogPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteLogPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteLogPoolWithContextV2(ctx context.Context, request *DeleteLogPoolRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLogPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteLogPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListLogPoolsSend(request *ListLogPoolsRequest) (*ListLogPoolsResponse, error) {
	statusCode, msg, err := c.ListLogPoolsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListLogPoolsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ListLogPoolsWithContextV2(ctx context.Context, request *ListLogPoolsRequest) (int, string, error) {
	if request == nil {
		request = NewListLogPoolsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListLogPoolsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetLogsSend(request *GetLogsRequest) (*GetLogsResponse, error) {
	statusCode, msg, err := c.GetLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetLogsWithContextV2(ctx context.Context, request *GetLogsRequest) (int, string, error) {
	if request == nil {
		request = NewGetLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateQuickSearchSend(request *CreateQuickSearchRequest) (*CreateQuickSearchResponse, error) {
	statusCode, msg, err := c.CreateQuickSearchWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateQuickSearchResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateQuickSearchWithContextV2(ctx context.Context, request *CreateQuickSearchRequest) (int, string, error) {
	if request == nil {
		request = NewCreateQuickSearchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateQuickSearchResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListQuickSearchsSend(request *ListQuickSearchsRequest) (*ListQuickSearchsResponse, error) {
	statusCode, msg, err := c.ListQuickSearchsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListQuickSearchsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ListQuickSearchsWithContextV2(ctx context.Context, request *ListQuickSearchsRequest) (int, string, error) {
	if request == nil {
		request = NewListQuickSearchsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListQuickSearchsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteQuickSearchsSend(request *DeleteQuickSearchsRequest) (*DeleteQuickSearchsResponse, error) {
	statusCode, msg, err := c.DeleteQuickSearchsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteQuickSearchsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteQuickSearchsWithContextV2(ctx context.Context, request *DeleteQuickSearchsRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteQuickSearchsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteQuickSearchsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateDownloadTaskSend(request *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error) {
	statusCode, msg, err := c.CreateDownloadTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDownloadTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateDownloadTaskWithContextV2(ctx context.Context, request *CreateDownloadTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDownloadTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDownloadTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListDownloadTasksSend(request *ListDownloadTasksRequest) (*ListDownloadTasksResponse, error) {
	statusCode, msg, err := c.ListDownloadTasksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListDownloadTasksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ListDownloadTasksWithContextV2(ctx context.Context, request *ListDownloadTasksRequest) (int, string, error) {
	if request == nil {
		request = NewListDownloadTasksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListDownloadTasksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
