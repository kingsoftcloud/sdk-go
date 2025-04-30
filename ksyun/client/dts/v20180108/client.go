package v20180108
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-01-08"

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

func NewSchemaStructRequest() (request *SchemaStructRequest) {
	request = &SchemaStructRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "SchemaStruct")
	return
}

func NewSchemaStructResponse() (response *SchemaStructResponse) {
	response = &SchemaStructResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SchemaStruct(request *SchemaStructRequest) string {
	return c.SchemaStructWithContext(context.Background(), request)
}

func (c *Client) SchemaStructWithContext(ctx context.Context, request *SchemaStructRequest) string {
	if request == nil {
		request = NewSchemaStructRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSchemaStructResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewConnectivityCheckRequest() (request *ConnectivityCheckRequest) {
	request = &ConnectivityCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ConnectivityCheck")
	return
}

func NewConnectivityCheckResponse() (response *ConnectivityCheckResponse) {
	response = &ConnectivityCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ConnectivityCheck(request *ConnectivityCheckRequest) string {
	return c.ConnectivityCheckWithContext(context.Background(), request)
}

func (c *Client) ConnectivityCheckWithContext(ctx context.Context, request *ConnectivityCheckRequest) string {
	if request == nil {
		request = NewConnectivityCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewConnectivityCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreatePrecheckRequest() (request *CreatePrecheckRequest) {
	request = &CreatePrecheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreatePrecheck")
	return
}

func NewCreatePrecheckResponse() (response *CreatePrecheckResponse) {
	response = &CreatePrecheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePrecheck(request *CreatePrecheckRequest) string {
	return c.CreatePrecheckWithContext(context.Background(), request)
}

func (c *Client) CreatePrecheckWithContext(ctx context.Context, request *CreatePrecheckRequest) string {
	if request == nil {
		request = NewCreatePrecheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreatePrecheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateTask")
	return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTask(request *CreateTaskRequest) string {
	return c.CreateTaskWithContext(context.Background(), request)
}

func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) string {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTask(request *DescribeTaskRequest) string {
	return c.DescribeTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) string {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewOperateTaskRequest() (request *OperateTaskRequest) {
	request = &OperateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "OperateTask")
	return
}

func NewOperateTaskResponse() (response *OperateTaskResponse) {
	response = &OperateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperateTask(request *OperateTaskRequest) string {
	return c.OperateTaskWithContext(context.Background(), request)
}

func (c *Client) OperateTaskWithContext(ctx context.Context, request *OperateTaskRequest) string {
	if request == nil {
		request = NewOperateTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeConnConfigRequest() (request *DescribeConnConfigRequest) {
	request = &DescribeConnConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeConnConfig")
	return
}

func NewDescribeConnConfigResponse() (response *DescribeConnConfigResponse) {
	response = &DescribeConnConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeConnConfig(request *DescribeConnConfigRequest) string {
	return c.DescribeConnConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeConnConfigWithContext(ctx context.Context, request *DescribeConnConfigRequest) string {
	if request == nil {
		request = NewDescribeConnConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeConnConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribePrecheckRequest() (request *DescribePrecheckRequest) {
	request = &DescribePrecheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribePrecheck")
	return
}

func NewDescribePrecheckResponse() (response *DescribePrecheckResponse) {
	response = &DescribePrecheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrecheck(request *DescribePrecheckRequest) string {
	return c.DescribePrecheckWithContext(context.Background(), request)
}

func (c *Client) DescribePrecheckWithContext(ctx context.Context, request *DescribePrecheckRequest) string {
	if request == nil {
		request = NewDescribePrecheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePrecheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeSourceUserConfigRequest() (request *DescribeSourceUserConfigRequest) {
	request = &DescribeSourceUserConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUserConfig")
	return
}

func NewDescribeSourceUserConfigResponse() (response *DescribeSourceUserConfigResponse) {
	response = &DescribeSourceUserConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSourceUserConfig(request *DescribeSourceUserConfigRequest) string {
	return c.DescribeSourceUserConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeSourceUserConfigWithContext(ctx context.Context, request *DescribeSourceUserConfigRequest) string {
	if request == nil {
		request = NewDescribeSourceUserConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSourceUserConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSetConsistencyCheckRequest() (request *SetConsistencyCheckRequest) {
	request = &SetConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "SetConsistencyCheck")
	return
}

func NewSetConsistencyCheckResponse() (response *SetConsistencyCheckResponse) {
	response = &SetConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetConsistencyCheck(request *SetConsistencyCheckRequest) string {
	return c.SetConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) SetConsistencyCheckWithContext(ctx context.Context, request *SetConsistencyCheckRequest) string {
	if request == nil {
		request = NewSetConsistencyCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeConsistencyCheckRequest() (request *DescribeConsistencyCheckRequest) {
	request = &DescribeConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeConsistencyCheck")
	return
}

func NewDescribeConsistencyCheckResponse() (response *DescribeConsistencyCheckResponse) {
	response = &DescribeConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeConsistencyCheck(request *DescribeConsistencyCheckRequest) string {
	return c.DescribeConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) DescribeConsistencyCheckWithContext(ctx context.Context, request *DescribeConsistencyCheckRequest) string {
	if request == nil {
		request = NewDescribeConsistencyCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDTSParameterRequest() (request *DescribeDTSParameterRequest) {
	request = &DescribeDTSParameterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameter")
	return
}

func NewDescribeDTSParameterResponse() (response *DescribeDTSParameterResponse) {
	response = &DescribeDTSParameterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDTSParameter(request *DescribeDTSParameterRequest) string {
	return c.DescribeDTSParameterWithContext(context.Background(), request)
}

func (c *Client) DescribeDTSParameterWithContext(ctx context.Context, request *DescribeDTSParameterRequest) string {
	if request == nil {
		request = NewDescribeDTSParameterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDTSParameterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDTSParameterConfigRequest() (request *DescribeDTSParameterConfigRequest) {
	request = &DescribeDTSParameterConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeDTSParameterConfig")
	return
}

func NewDescribeDTSParameterConfigResponse() (response *DescribeDTSParameterConfigResponse) {
	response = &DescribeDTSParameterConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDTSParameterConfig(request *DescribeDTSParameterConfigRequest) string {
	return c.DescribeDTSParameterConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeDTSParameterConfigWithContext(ctx context.Context, request *DescribeDTSParameterConfigRequest) string {
	if request == nil {
		request = NewDescribeDTSParameterConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDTSParameterConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeSourceUserRequest() (request *DescribeSourceUserRequest) {
	request = &DescribeSourceUserRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSourceUser")
	return
}

func NewDescribeSourceUserResponse() (response *DescribeSourceUserResponse) {
	response = &DescribeSourceUserResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSourceUser(request *DescribeSourceUserRequest) string {
	return c.DescribeSourceUserWithContext(context.Background(), request)
}

func (c *Client) DescribeSourceUserWithContext(ctx context.Context, request *DescribeSourceUserRequest) string {
	if request == nil {
		request = NewDescribeSourceUserRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSourceUserResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeSubTaskRequest() (request *DescribeSubTaskRequest) {
	request = &DescribeSubTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubTask")
	return
}

func NewDescribeSubTaskResponse() (response *DescribeSubTaskResponse) {
	response = &DescribeSubTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubTask(request *DescribeSubTaskRequest) string {
	return c.DescribeSubTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeSubTaskWithContext(ctx context.Context, request *DescribeSubTaskRequest) string {
	if request == nil {
		request = NewDescribeSubTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSubTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateConsistencyCheckRequest() (request *CreateConsistencyCheckRequest) {
	request = &CreateConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateConsistencyCheck")
	return
}

func NewCreateConsistencyCheckResponse() (response *CreateConsistencyCheckResponse) {
	response = &CreateConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateConsistencyCheck(request *CreateConsistencyCheckRequest) string {
	return c.CreateConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) CreateConsistencyCheckWithContext(ctx context.Context, request *CreateConsistencyCheckRequest) string {
	if request == nil {
		request = NewCreateConsistencyCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStopConsistencyCheckRequest() (request *StopConsistencyCheckRequest) {
	request = &StopConsistencyCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "StopConsistencyCheck")
	return
}

func NewStopConsistencyCheckResponse() (response *StopConsistencyCheckResponse) {
	response = &StopConsistencyCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopConsistencyCheck(request *StopConsistencyCheckRequest) string {
	return c.StopConsistencyCheckWithContext(context.Background(), request)
}

func (c *Client) StopConsistencyCheckWithContext(ctx context.Context, request *StopConsistencyCheckRequest) string {
	if request == nil {
		request = NewStopConsistencyCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopConsistencyCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRegionConfigRequest() (request *DescribeRegionConfigRequest) {
	request = &DescribeRegionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConfig")
	return
}

func NewDescribeRegionConfigResponse() (response *DescribeRegionConfigResponse) {
	response = &DescribeRegionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegionConfig(request *DescribeRegionConfigRequest) string {
	return c.DescribeRegionConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionConfigWithContext(ctx context.Context, request *DescribeRegionConfigRequest) string {
	if request == nil {
		request = NewDescribeRegionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRegionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewTaskBirdViewRequest() (request *TaskBirdViewRequest) {
	request = &TaskBirdViewRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "TaskBirdView")
	return
}

func NewTaskBirdViewResponse() (response *TaskBirdViewResponse) {
	response = &TaskBirdViewResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TaskBirdView(request *TaskBirdViewRequest) string {
	return c.TaskBirdViewWithContext(context.Background(), request)
}

func (c *Client) TaskBirdViewWithContext(ctx context.Context, request *TaskBirdViewRequest) string {
	if request == nil {
		request = NewTaskBirdViewRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewTaskBirdViewResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


