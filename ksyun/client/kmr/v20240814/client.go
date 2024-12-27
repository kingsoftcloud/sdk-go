package v20240814
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func (c *Client) DetailWorkspace(request *DetailWorkspaceRequest) (string) {
    return c.DetailWorkspaceWithContext(context.Background(), request)
}

func (c *Client) DetailWorkspaceWithContext(ctx context.Context, request *DetailWorkspaceRequest) (string) {
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

func (c *Client) ListWorkspaces(request *ListWorkspacesRequest) (string) {
    return c.ListWorkspacesWithContext(context.Background(), request)
}

func (c *Client) ListWorkspacesWithContext(ctx context.Context, request *ListWorkspacesRequest) (string) {
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

func (c *Client) StartJobRun(request *StartJobRunRequest) (string) {
    return c.StartJobRunWithContext(context.Background(), request)
}

func (c *Client) StartJobRunWithContext(ctx context.Context, request *StartJobRunRequest) (string) {
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

func (c *Client) GetJobRun(request *GetJobRunRequest) (string) {
    return c.GetJobRunWithContext(context.Background(), request)
}

func (c *Client) GetJobRunWithContext(ctx context.Context, request *GetJobRunRequest) (string) {
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

func (c *Client) ListJobRuns(request *ListJobRunsRequest) (string) {
    return c.ListJobRunsWithContext(context.Background(), request)
}

func (c *Client) ListJobRunsWithContext(ctx context.Context, request *ListJobRunsRequest) (string) {
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

func (c *Client) CancelJobRun(request *CancelJobRunRequest) (string) {
    return c.CancelJobRunWithContext(context.Background(), request)
}

func (c *Client) CancelJobRunWithContext(ctx context.Context, request *CancelJobRunRequest) (string) {
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

func (c *Client) ListExecutor(request *ListExecutorRequest) (string) {
    return c.ListExecutorWithContext(context.Background(), request)
}

func (c *Client) ListExecutorWithContext(ctx context.Context, request *ListExecutorRequest) (string) {
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


