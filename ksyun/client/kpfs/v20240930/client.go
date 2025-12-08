package v20240930
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-09-30"

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

func NewUpdatePerformanceOnePosixAclRequest() (request *UpdatePerformanceOnePosixAclRequest) {
	request = &UpdatePerformanceOnePosixAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceOnePosixAcl")
	return
}

func NewUpdatePerformanceOnePosixAclResponse() (response *UpdatePerformanceOnePosixAclResponse) {
	response = &UpdatePerformanceOnePosixAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePerformanceOnePosixAcl(request *UpdatePerformanceOnePosixAclRequest) string {
	return c.UpdatePerformanceOnePosixAclWithContext(context.Background(), request)
}

func (c *Client) UpdatePerformanceOnePosixAclSend(request *UpdatePerformanceOnePosixAclRequest) (*UpdatePerformanceOnePosixAclResponse, error) {
	statusCode, msg, err := c.UpdatePerformanceOnePosixAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdatePerformanceOnePosixAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdatePerformanceOnePosixAclWithContext(ctx context.Context, request *UpdatePerformanceOnePosixAclRequest) string {
	if request == nil {
		request = NewUpdatePerformanceOnePosixAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceOnePosixAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdatePerformanceOnePosixAclWithContextV2(ctx context.Context, request *UpdatePerformanceOnePosixAclRequest) (int, string, error) {
	if request == nil {
		request = NewUpdatePerformanceOnePosixAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceOnePosixAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePerformanceOnePosixAclListRequest() (request *DescribePerformanceOnePosixAclListRequest) {
	request = &DescribePerformanceOnePosixAclListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOnePosixAclList")
	return
}

func NewDescribePerformanceOnePosixAclListResponse() (response *DescribePerformanceOnePosixAclListResponse) {
	response = &DescribePerformanceOnePosixAclListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePerformanceOnePosixAclList(request *DescribePerformanceOnePosixAclListRequest) string {
	return c.DescribePerformanceOnePosixAclListWithContext(context.Background(), request)
}

func (c *Client) DescribePerformanceOnePosixAclListSend(request *DescribePerformanceOnePosixAclListRequest) (*DescribePerformanceOnePosixAclListResponse, error) {
	statusCode, msg, err := c.DescribePerformanceOnePosixAclListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePerformanceOnePosixAclListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePerformanceOnePosixAclListWithContext(ctx context.Context, request *DescribePerformanceOnePosixAclListRequest) string {
	if request == nil {
		request = NewDescribePerformanceOnePosixAclListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePerformanceOnePosixAclListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePerformanceOnePosixAclListWithContextV2(ctx context.Context, request *DescribePerformanceOnePosixAclListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePerformanceOnePosixAclListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePerformanceOnePosixAclListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewManageDataFlowTaskRequest() (request *ManageDataFlowTaskRequest) {
	request = &ManageDataFlowTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowTask")
	return
}

func NewManageDataFlowTaskResponse() (response *ManageDataFlowTaskResponse) {
	response = &ManageDataFlowTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ManageDataFlowTask(request *ManageDataFlowTaskRequest) string {
	return c.ManageDataFlowTaskWithContext(context.Background(), request)
}

func (c *Client) ManageDataFlowTaskSend(request *ManageDataFlowTaskRequest) (*ManageDataFlowTaskResponse, error) {
	statusCode, msg, err := c.ManageDataFlowTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ManageDataFlowTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ManageDataFlowTaskWithContext(ctx context.Context, request *ManageDataFlowTaskRequest) string {
	if request == nil {
		request = NewManageDataFlowTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewManageDataFlowTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ManageDataFlowTaskWithContextV2(ctx context.Context, request *ManageDataFlowTaskRequest) (int, string, error) {
	if request == nil {
		request = NewManageDataFlowTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewManageDataFlowTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDataFlowStrategyRequest() (request *CreateDataFlowStrategyRequest) {
	request = &CreateDataFlowStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowStrategy")
	return
}

func NewCreateDataFlowStrategyResponse() (response *CreateDataFlowStrategyResponse) {
	response = &CreateDataFlowStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataFlowStrategy(request *CreateDataFlowStrategyRequest) string {
	return c.CreateDataFlowStrategyWithContext(context.Background(), request)
}

func (c *Client) CreateDataFlowStrategySend(request *CreateDataFlowStrategyRequest) (*CreateDataFlowStrategyResponse, error) {
	statusCode, msg, err := c.CreateDataFlowStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDataFlowStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDataFlowStrategyWithContext(ctx context.Context, request *CreateDataFlowStrategyRequest) string {
	if request == nil {
		request = NewCreateDataFlowStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDataFlowStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDataFlowStrategyWithContextV2(ctx context.Context, request *CreateDataFlowStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDataFlowStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDataFlowStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDataFlowTaskRequest() (request *ModifyDataFlowTaskRequest) {
	request = &ModifyDataFlowTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ModifyDataFlowTask")
	return
}

func NewModifyDataFlowTaskResponse() (response *ModifyDataFlowTaskResponse) {
	response = &ModifyDataFlowTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDataFlowTask(request *ModifyDataFlowTaskRequest) string {
	return c.ModifyDataFlowTaskWithContext(context.Background(), request)
}

func (c *Client) ModifyDataFlowTaskSend(request *ModifyDataFlowTaskRequest) (*ModifyDataFlowTaskResponse, error) {
	statusCode, msg, err := c.ModifyDataFlowTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyDataFlowTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDataFlowTaskWithContext(ctx context.Context, request *ModifyDataFlowTaskRequest) string {
	if request == nil {
		request = NewModifyDataFlowTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDataFlowTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDataFlowTaskWithContextV2(ctx context.Context, request *ModifyDataFlowTaskRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDataFlowTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDataFlowTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowTaskListRequest() (request *DescribeDataFlowTaskListRequest) {
	request = &DescribeDataFlowTaskListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTaskList")
	return
}

func NewDescribeDataFlowTaskListResponse() (response *DescribeDataFlowTaskListResponse) {
	response = &DescribeDataFlowTaskListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowTaskList(request *DescribeDataFlowTaskListRequest) string {
	return c.DescribeDataFlowTaskListWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowTaskListSend(request *DescribeDataFlowTaskListRequest) (*DescribeDataFlowTaskListResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowTaskListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDataFlowTaskListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowTaskListWithContext(ctx context.Context, request *DescribeDataFlowTaskListRequest) string {
	if request == nil {
		request = NewDescribeDataFlowTaskListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowTaskListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowTaskListWithContextV2(ctx context.Context, request *DescribeDataFlowTaskListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowTaskListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowTaskListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateDataFlowTaskRequest() (request *ActivateDataFlowTaskRequest) {
	request = &ActivateDataFlowTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ActivateDataFlowTask")
	return
}

func NewActivateDataFlowTaskResponse() (response *ActivateDataFlowTaskResponse) {
	response = &ActivateDataFlowTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateDataFlowTask(request *ActivateDataFlowTaskRequest) string {
	return c.ActivateDataFlowTaskWithContext(context.Background(), request)
}

func (c *Client) ActivateDataFlowTaskSend(request *ActivateDataFlowTaskRequest) (*ActivateDataFlowTaskResponse, error) {
	statusCode, msg, err := c.ActivateDataFlowTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ActivateDataFlowTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateDataFlowTaskWithContext(ctx context.Context, request *ActivateDataFlowTaskRequest) string {
	if request == nil {
		request = NewActivateDataFlowTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewActivateDataFlowTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateDataFlowTaskWithContextV2(ctx context.Context, request *ActivateDataFlowTaskRequest) (int, string, error) {
	if request == nil {
		request = NewActivateDataFlowTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewActivateDataFlowTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDataFlowStrategyRequest() (request *DeleteDataFlowStrategyRequest) {
	request = &DeleteDataFlowStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlowStrategy")
	return
}

func NewDeleteDataFlowStrategyResponse() (response *DeleteDataFlowStrategyResponse) {
	response = &DeleteDataFlowStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDataFlowStrategy(request *DeleteDataFlowStrategyRequest) string {
	return c.DeleteDataFlowStrategyWithContext(context.Background(), request)
}

func (c *Client) DeleteDataFlowStrategySend(request *DeleteDataFlowStrategyRequest) (*DeleteDataFlowStrategyResponse, error) {
	statusCode, msg, err := c.DeleteDataFlowStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteDataFlowStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDataFlowStrategyWithContext(ctx context.Context, request *DeleteDataFlowStrategyRequest) string {
	if request == nil {
		request = NewDeleteDataFlowStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataFlowStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDataFlowStrategyWithContextV2(ctx context.Context, request *DeleteDataFlowStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDataFlowStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataFlowStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowStrategyListRequest() (request *DescribeDataFlowStrategyListRequest) {
	request = &DescribeDataFlowStrategyListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategyList")
	return
}

func NewDescribeDataFlowStrategyListResponse() (response *DescribeDataFlowStrategyListResponse) {
	response = &DescribeDataFlowStrategyListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowStrategyList(request *DescribeDataFlowStrategyListRequest) string {
	return c.DescribeDataFlowStrategyListWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowStrategyListSend(request *DescribeDataFlowStrategyListRequest) (*DescribeDataFlowStrategyListResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowStrategyListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDataFlowStrategyListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowStrategyListWithContext(ctx context.Context, request *DescribeDataFlowStrategyListRequest) string {
	if request == nil {
		request = NewDescribeDataFlowStrategyListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowStrategyListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowStrategyListWithContextV2(ctx context.Context, request *DescribeDataFlowStrategyListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowStrategyListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowStrategyListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


