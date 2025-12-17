package v20240814

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-08-14"

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

func NewDetailWorkspaceRequest() (request *DetailWorkspaceRequest) {
	request = &DetailWorkspaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DetailWorkspace")
	return
}

func NewDetailWorkspaceResponse() (response *DetailWorkspaceResponse) {
	response = &DetailWorkspaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetailWorkspace(request *DetailWorkspaceRequest) string {
	return c.DetailWorkspaceWithContext(context.Background(), request)
}

func (c *Client) DetailWorkspaceSend(request *DetailWorkspaceRequest) (*DetailWorkspaceResponse, error) {
	statusCode, msg, err := c.DetailWorkspaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetailWorkspaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetailWorkspaceWithContext(ctx context.Context, request *DetailWorkspaceRequest) string {
	if request == nil {
		request = NewDetailWorkspaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDetailWorkspaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetailWorkspaceWithContextV2(ctx context.Context, request *DetailWorkspaceRequest) (int, string, error) {
	if request == nil {
		request = NewDetailWorkspaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDetailWorkspaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListWorkspacesRequest() (request *ListWorkspacesRequest) {
	request = &ListWorkspacesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListWorkspaces")
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

func (c *Client) ListWorkspacesSend(request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	statusCode, msg, err := c.ListWorkspacesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListWorkspacesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListWorkspacesWithContext(ctx context.Context, request *ListWorkspacesRequest) string {
	if request == nil {
		request = NewListWorkspacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListWorkspacesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListWorkspacesWithContextV2(ctx context.Context, request *ListWorkspacesRequest) (int, string, error) {
	if request == nil {
		request = NewListWorkspacesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListWorkspacesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartJobRunRequest() (request *StartJobRunRequest) {
	request = &StartJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "StartJobRun")
	return
}

func NewStartJobRunResponse() (response *StartJobRunResponse) {
	response = &StartJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartJobRun(request *StartJobRunRequest) string {
	return c.StartJobRunWithContext(context.Background(), request)
}

func (c *Client) StartJobRunSend(request *StartJobRunRequest) (*StartJobRunResponse, error) {
	statusCode, msg, err := c.StartJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartJobRunWithContext(ctx context.Context, request *StartJobRunRequest) string {
	if request == nil {
		request = NewStartJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartJobRunWithContextV2(ctx context.Context, request *StartJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewStartJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetJobRunRequest() (request *GetJobRunRequest) {
	request = &GetJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "GetJobRun")
	return
}

func NewGetJobRunResponse() (response *GetJobRunResponse) {
	response = &GetJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetJobRun(request *GetJobRunRequest) string {
	return c.GetJobRunWithContext(context.Background(), request)
}

func (c *Client) GetJobRunSend(request *GetJobRunRequest) (*GetJobRunResponse, error) {
	statusCode, msg, err := c.GetJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetJobRunWithContext(ctx context.Context, request *GetJobRunRequest) string {
	if request == nil {
		request = NewGetJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetJobRunWithContextV2(ctx context.Context, request *GetJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewGetJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListJobRunsRequest() (request *ListJobRunsRequest) {
	request = &ListJobRunsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListJobRuns")
	return
}

func NewListJobRunsResponse() (response *ListJobRunsResponse) {
	response = &ListJobRunsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListJobRuns(request *ListJobRunsRequest) string {
	return c.ListJobRunsWithContext(context.Background(), request)
}

func (c *Client) ListJobRunsSend(request *ListJobRunsRequest) (*ListJobRunsResponse, error) {
	statusCode, msg, err := c.ListJobRunsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListJobRunsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListJobRunsWithContext(ctx context.Context, request *ListJobRunsRequest) string {
	if request == nil {
		request = NewListJobRunsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListJobRunsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListJobRunsWithContextV2(ctx context.Context, request *ListJobRunsRequest) (int, string, error) {
	if request == nil {
		request = NewListJobRunsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListJobRunsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelJobRunRequest() (request *CancelJobRunRequest) {
	request = &CancelJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "CancelJobRun")
	return
}

func NewCancelJobRunResponse() (response *CancelJobRunResponse) {
	response = &CancelJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelJobRun(request *CancelJobRunRequest) string {
	return c.CancelJobRunWithContext(context.Background(), request)
}

func (c *Client) CancelJobRunSend(request *CancelJobRunRequest) (*CancelJobRunResponse, error) {
	statusCode, msg, err := c.CancelJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelJobRunWithContext(ctx context.Context, request *CancelJobRunRequest) string {
	if request == nil {
		request = NewCancelJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelJobRunWithContextV2(ctx context.Context, request *CancelJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewCancelJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListExecutorRequest() (request *ListExecutorRequest) {
	request = &ListExecutorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListExecutor")
	return
}

func NewListExecutorResponse() (response *ListExecutorResponse) {
	response = &ListExecutorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListExecutor(request *ListExecutorRequest) string {
	return c.ListExecutorWithContext(context.Background(), request)
}

func (c *Client) ListExecutorSend(request *ListExecutorRequest) (*ListExecutorResponse, error) {
	statusCode, msg, err := c.ListExecutorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListExecutorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListExecutorWithContext(ctx context.Context, request *ListExecutorRequest) string {
	if request == nil {
		request = NewListExecutorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListExecutorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListExecutorWithContextV2(ctx context.Context, request *ListExecutorRequest) (int, string, error) {
	if request == nil {
		request = NewListExecutorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListExecutorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartRayJobRunRequest() (request *StartRayJobRunRequest) {
	request = &StartRayJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "StartRayJobRun")
	return
}

func NewStartRayJobRunResponse() (response *StartRayJobRunResponse) {
	response = &StartRayJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartRayJobRun(request *StartRayJobRunRequest) string {
	return c.StartRayJobRunWithContext(context.Background(), request)
}

func (c *Client) StartRayJobRunSend(request *StartRayJobRunRequest) (*StartRayJobRunResponse, error) {
	statusCode, msg, err := c.StartRayJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartRayJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartRayJobRunWithContext(ctx context.Context, request *StartRayJobRunRequest) string {
	if request == nil {
		request = NewStartRayJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartRayJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartRayJobRunWithContextV2(ctx context.Context, request *StartRayJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewStartRayJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartRayJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRayJobRunRequest() (request *GetRayJobRunRequest) {
	request = &GetRayJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "GetRayJobRun")
	return
}

func NewGetRayJobRunResponse() (response *GetRayJobRunResponse) {
	response = &GetRayJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRayJobRun(request *GetRayJobRunRequest) string {
	return c.GetRayJobRunWithContext(context.Background(), request)
}

func (c *Client) GetRayJobRunSend(request *GetRayJobRunRequest) (*GetRayJobRunResponse, error) {
	statusCode, msg, err := c.GetRayJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRayJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRayJobRunWithContext(ctx context.Context, request *GetRayJobRunRequest) string {
	if request == nil {
		request = NewGetRayJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetRayJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRayJobRunWithContextV2(ctx context.Context, request *GetRayJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewGetRayJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetRayJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListRayJobRunsRequest() (request *ListRayJobRunsRequest) {
	request = &ListRayJobRunsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListRayJobRuns")
	return
}

func NewListRayJobRunsResponse() (response *ListRayJobRunsResponse) {
	response = &ListRayJobRunsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListRayJobRuns(request *ListRayJobRunsRequest) string {
	return c.ListRayJobRunsWithContext(context.Background(), request)
}

func (c *Client) ListRayJobRunsSend(request *ListRayJobRunsRequest) (*ListRayJobRunsResponse, error) {
	statusCode, msg, err := c.ListRayJobRunsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListRayJobRunsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListRayJobRunsWithContext(ctx context.Context, request *ListRayJobRunsRequest) string {
	if request == nil {
		request = NewListRayJobRunsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListRayJobRunsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListRayJobRunsWithContextV2(ctx context.Context, request *ListRayJobRunsRequest) (int, string, error) {
	if request == nil {
		request = NewListRayJobRunsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListRayJobRunsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelRayJobRunRequest() (request *CancelRayJobRunRequest) {
	request = &CancelRayJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "CancelRayJobRun")
	return
}

func NewCancelRayJobRunResponse() (response *CancelRayJobRunResponse) {
	response = &CancelRayJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelRayJobRun(request *CancelRayJobRunRequest) string {
	return c.CancelRayJobRunWithContext(context.Background(), request)
}

func (c *Client) CancelRayJobRunSend(request *CancelRayJobRunRequest) (*CancelRayJobRunResponse, error) {
	statusCode, msg, err := c.CancelRayJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelRayJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelRayJobRunWithContext(ctx context.Context, request *CancelRayJobRunRequest) string {
	if request == nil {
		request = NewCancelRayJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelRayJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelRayJobRunWithContextV2(ctx context.Context, request *CancelRayJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewCancelRayJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelRayJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartFlinkJobRunRequest() (request *StartFlinkJobRunRequest) {
	request = &StartFlinkJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "StartFlinkJobRun")
	return
}

func NewStartFlinkJobRunResponse() (response *StartFlinkJobRunResponse) {
	response = &StartFlinkJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartFlinkJobRun(request *StartFlinkJobRunRequest) string {
	return c.StartFlinkJobRunWithContext(context.Background(), request)
}

func (c *Client) StartFlinkJobRunSend(request *StartFlinkJobRunRequest) (*StartFlinkJobRunResponse, error) {
	statusCode, msg, err := c.StartFlinkJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartFlinkJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartFlinkJobRunWithContext(ctx context.Context, request *StartFlinkJobRunRequest) string {
	if request == nil {
		request = NewStartFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartFlinkJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartFlinkJobRunWithContextV2(ctx context.Context, request *StartFlinkJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewStartFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartFlinkJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetFlinkJobRunRequest() (request *GetFlinkJobRunRequest) {
	request = &GetFlinkJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "GetFlinkJobRun")
	return
}

func NewGetFlinkJobRunResponse() (response *GetFlinkJobRunResponse) {
	response = &GetFlinkJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetFlinkJobRun(request *GetFlinkJobRunRequest) string {
	return c.GetFlinkJobRunWithContext(context.Background(), request)
}

func (c *Client) GetFlinkJobRunSend(request *GetFlinkJobRunRequest) (*GetFlinkJobRunResponse, error) {
	statusCode, msg, err := c.GetFlinkJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetFlinkJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetFlinkJobRunWithContext(ctx context.Context, request *GetFlinkJobRunRequest) string {
	if request == nil {
		request = NewGetFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetFlinkJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetFlinkJobRunWithContextV2(ctx context.Context, request *GetFlinkJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewGetFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetFlinkJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListFlinkJobRunsRequest() (request *ListFlinkJobRunsRequest) {
	request = &ListFlinkJobRunsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListFlinkJobRuns")
	return
}

func NewListFlinkJobRunsResponse() (response *ListFlinkJobRunsResponse) {
	response = &ListFlinkJobRunsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListFlinkJobRuns(request *ListFlinkJobRunsRequest) string {
	return c.ListFlinkJobRunsWithContext(context.Background(), request)
}

func (c *Client) ListFlinkJobRunsSend(request *ListFlinkJobRunsRequest) (*ListFlinkJobRunsResponse, error) {
	statusCode, msg, err := c.ListFlinkJobRunsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListFlinkJobRunsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListFlinkJobRunsWithContext(ctx context.Context, request *ListFlinkJobRunsRequest) string {
	if request == nil {
		request = NewListFlinkJobRunsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListFlinkJobRunsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListFlinkJobRunsWithContextV2(ctx context.Context, request *ListFlinkJobRunsRequest) (int, string, error) {
	if request == nil {
		request = NewListFlinkJobRunsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListFlinkJobRunsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelFlinkJobRunRequest() (request *CancelFlinkJobRunRequest) {
	request = &CancelFlinkJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "CancelFlinkJobRun")
	return
}

func NewCancelFlinkJobRunResponse() (response *CancelFlinkJobRunResponse) {
	response = &CancelFlinkJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelFlinkJobRun(request *CancelFlinkJobRunRequest) string {
	return c.CancelFlinkJobRunWithContext(context.Background(), request)
}

func (c *Client) CancelFlinkJobRunSend(request *CancelFlinkJobRunRequest) (*CancelFlinkJobRunResponse, error) {
	statusCode, msg, err := c.CancelFlinkJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelFlinkJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelFlinkJobRunWithContext(ctx context.Context, request *CancelFlinkJobRunRequest) string {
	if request == nil {
		request = NewCancelFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelFlinkJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelFlinkJobRunWithContextV2(ctx context.Context, request *CancelFlinkJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewCancelFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelFlinkJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSuspendFlinkJobRunRequest() (request *SuspendFlinkJobRunRequest) {
	request = &SuspendFlinkJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "SuspendFlinkJobRun")
	return
}

func NewSuspendFlinkJobRunResponse() (response *SuspendFlinkJobRunResponse) {
	response = &SuspendFlinkJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SuspendFlinkJobRun(request *SuspendFlinkJobRunRequest) string {
	return c.SuspendFlinkJobRunWithContext(context.Background(), request)
}

func (c *Client) SuspendFlinkJobRunSend(request *SuspendFlinkJobRunRequest) (*SuspendFlinkJobRunResponse, error) {
	statusCode, msg, err := c.SuspendFlinkJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SuspendFlinkJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SuspendFlinkJobRunWithContext(ctx context.Context, request *SuspendFlinkJobRunRequest) string {
	if request == nil {
		request = NewSuspendFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSuspendFlinkJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SuspendFlinkJobRunWithContextV2(ctx context.Context, request *SuspendFlinkJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewSuspendFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSuspendFlinkJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestartFlinkJobRunRequest() (request *RestartFlinkJobRunRequest) {
	request = &RestartFlinkJobRunRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "RestartFlinkJobRun")
	return
}

func NewRestartFlinkJobRunResponse() (response *RestartFlinkJobRunResponse) {
	response = &RestartFlinkJobRunResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestartFlinkJobRun(request *RestartFlinkJobRunRequest) string {
	return c.RestartFlinkJobRunWithContext(context.Background(), request)
}

func (c *Client) RestartFlinkJobRunSend(request *RestartFlinkJobRunRequest) (*RestartFlinkJobRunResponse, error) {
	statusCode, msg, err := c.RestartFlinkJobRunWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RestartFlinkJobRunResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestartFlinkJobRunWithContext(ctx context.Context, request *RestartFlinkJobRunRequest) string {
	if request == nil {
		request = NewRestartFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestartFlinkJobRunResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestartFlinkJobRunWithContextV2(ctx context.Context, request *RestartFlinkJobRunRequest) (int, string, error) {
	if request == nil {
		request = NewRestartFlinkJobRunRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestartFlinkJobRunResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMetricListRequest() (request *DescribeMetricListRequest) {
	request = &DescribeMetricListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DescribeMetricList")
	return
}

func NewDescribeMetricListResponse() (response *DescribeMetricListResponse) {
	response = &DescribeMetricListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMetricList(request *DescribeMetricListRequest) string {
	return c.DescribeMetricListWithContext(context.Background(), request)
}

func (c *Client) DescribeMetricListSend(request *DescribeMetricListRequest) (*DescribeMetricListResponse, error) {
	statusCode, msg, err := c.DescribeMetricListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMetricListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMetricListWithContext(ctx context.Context, request *DescribeMetricListRequest) string {
	if request == nil {
		request = NewDescribeMetricListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMetricListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMetricListWithContextV2(ctx context.Context, request *DescribeMetricListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMetricListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMetricListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryMetricsRequest() (request *QueryMetricsRequest) {
	request = &QueryMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "QueryMetrics")
	return
}

func NewQueryMetricsResponse() (response *QueryMetricsResponse) {
	response = &QueryMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryMetrics(request *QueryMetricsRequest) string {
	return c.QueryMetricsWithContext(context.Background(), request)
}

func (c *Client) QueryMetricsSend(request *QueryMetricsRequest) (*QueryMetricsResponse, error) {
	statusCode, msg, err := c.QueryMetricsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryMetricsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryMetricsWithContext(ctx context.Context, request *QueryMetricsRequest) string {
	if request == nil {
		request = NewQueryMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryMetricsWithContextV2(ctx context.Context, request *QueryMetricsRequest) (int, string, error) {
	if request == nil {
		request = NewQueryMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryMetricsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
