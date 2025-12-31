package v20231010

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2023-10-10"

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

func NewDeleteSecurityRulesRequest() (request *DeleteSecurityRulesRequest) {
	request = &DeleteSecurityRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "DeleteSecurityRules")
	return
}

func NewDeleteSecurityRulesResponse() (response *DeleteSecurityRulesResponse) {
	response = &DeleteSecurityRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityRules(request *DeleteSecurityRulesRequest) string {
	return c.DeleteSecurityRulesWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityRulesSend(request *DeleteSecurityRulesRequest) (*DeleteSecurityRulesResponse, error) {
	statusCode, msg, err := c.DeleteSecurityRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSecurityRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityRulesWithContext(ctx context.Context, request *DeleteSecurityRulesRequest) string {
	if request == nil {
		request = NewDeleteSecurityRulesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteSecurityRules")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityRulesWithContextV2(ctx context.Context, request *DeleteSecurityRulesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityRulesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteSecurityRules")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityRulesRequest() (request *CreateSecurityRulesRequest) {
	request = &CreateSecurityRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "CreateSecurityRules")
	return
}

func NewCreateSecurityRulesResponse() (response *CreateSecurityRulesResponse) {
	response = &CreateSecurityRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityRules(request *CreateSecurityRulesRequest) string {
	return c.CreateSecurityRulesWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityRulesSend(request *CreateSecurityRulesRequest) (*CreateSecurityRulesResponse, error) {
	statusCode, msg, err := c.CreateSecurityRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSecurityRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityRulesWithContext(ctx context.Context, request *CreateSecurityRulesRequest) string {
	if request == nil {
		request = NewCreateSecurityRulesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateSecurityRules")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSecurityRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityRulesWithContextV2(ctx context.Context, request *CreateSecurityRulesRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityRulesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateSecurityRules")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSecurityRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnbindSecurityGroupInstancesRequest() (request *UnbindSecurityGroupInstancesRequest) {
	request = &UnbindSecurityGroupInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "UnbindSecurityGroupInstances")
	return
}

func NewUnbindSecurityGroupInstancesResponse() (response *UnbindSecurityGroupInstancesResponse) {
	response = &UnbindSecurityGroupInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindSecurityGroupInstances(request *UnbindSecurityGroupInstancesRequest) string {
	return c.UnbindSecurityGroupInstancesWithContext(context.Background(), request)
}

func (c *Client) UnbindSecurityGroupInstancesSend(request *UnbindSecurityGroupInstancesRequest) (*UnbindSecurityGroupInstancesResponse, error) {
	statusCode, msg, err := c.UnbindSecurityGroupInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UnbindSecurityGroupInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnbindSecurityGroupInstancesWithContext(ctx context.Context, request *UnbindSecurityGroupInstancesRequest) string {
	if request == nil {
		request = NewUnbindSecurityGroupInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "UnbindSecurityGroupInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUnbindSecurityGroupInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnbindSecurityGroupInstancesWithContextV2(ctx context.Context, request *UnbindSecurityGroupInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewUnbindSecurityGroupInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "UnbindSecurityGroupInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUnbindSecurityGroupInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBindSecurityGroupInstancesRequest() (request *BindSecurityGroupInstancesRequest) {
	request = &BindSecurityGroupInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "BindSecurityGroupInstances")
	return
}

func NewBindSecurityGroupInstancesResponse() (response *BindSecurityGroupInstancesResponse) {
	response = &BindSecurityGroupInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindSecurityGroupInstances(request *BindSecurityGroupInstancesRequest) string {
	return c.BindSecurityGroupInstancesWithContext(context.Background(), request)
}

func (c *Client) BindSecurityGroupInstancesSend(request *BindSecurityGroupInstancesRequest) (*BindSecurityGroupInstancesResponse, error) {
	statusCode, msg, err := c.BindSecurityGroupInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BindSecurityGroupInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BindSecurityGroupInstancesWithContext(ctx context.Context, request *BindSecurityGroupInstancesRequest) string {
	if request == nil {
		request = NewBindSecurityGroupInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "BindSecurityGroupInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBindSecurityGroupInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BindSecurityGroupInstancesWithContextV2(ctx context.Context, request *BindSecurityGroupInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewBindSecurityGroupInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "BindSecurityGroupInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBindSecurityGroupInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "DeleteSecurityGroup")
	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
	response = &DeleteSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) string {
	return c.DeleteSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupSend(request *DeleteSecurityGroupRequest) (*DeleteSecurityGroupResponse, error) {
	statusCode, msg, err := c.DeleteSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) string {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityGroupWithContextV2(ctx context.Context, request *DeleteSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "DescribeSecurityGroup")
	return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
	response = &DescribeSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) string {
	return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupSend(request *DescribeSecurityGroupRequest) (*DescribeSecurityGroupResponse, error) {
	statusCode, msg, err := c.DescribeSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) string {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DescribeSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSecurityGroupWithContextV2(ctx context.Context, request *DescribeSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DescribeSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "CreateSecurityGroup")
	return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) string {
	return c.CreateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupSend(request *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	statusCode, msg, err := c.CreateSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) string {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityGroupWithContextV2(ctx context.Context, request *CreateSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListSecurityGroupRequest() (request *ListSecurityGroupRequest) {
	request = &ListSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "ListSecurityGroup")
	return
}

func NewListSecurityGroupResponse() (response *ListSecurityGroupResponse) {
	response = &ListSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSecurityGroup(request *ListSecurityGroupRequest) string {
	return c.ListSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ListSecurityGroupSend(request *ListSecurityGroupRequest) (*ListSecurityGroupResponse, error) {
	statusCode, msg, err := c.ListSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListSecurityGroupWithContext(ctx context.Context, request *ListSecurityGroupRequest) string {
	if request == nil {
		request = NewListSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListSecurityGroupWithContextV2(ctx context.Context, request *ListSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewListSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstance(request *DeleteInstanceRequest) string {
	return c.DeleteInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceSend(request *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	statusCode, msg, err := c.DeleteInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) string {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceWithContextV2(ctx context.Context, request *DeleteInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "DescribeInstance")
	return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstance(request *DescribeInstanceRequest) string {
	return c.DescribeInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceSend(request *DescribeInstanceRequest) (*DescribeInstanceResponse, error) {
	statusCode, msg, err := c.DescribeInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) string {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DescribeInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceWithContextV2(ctx context.Context, request *DescribeInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DescribeInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListInstanceRequest() (request *ListInstanceRequest) {
	request = &ListInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "ListInstance")
	return
}

func NewListInstanceResponse() (response *ListInstanceResponse) {
	response = &ListInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstance(request *ListInstanceRequest) string {
	return c.ListInstanceWithContext(context.Background(), request)
}

func (c *Client) ListInstanceSend(request *ListInstanceRequest) (*ListInstanceResponse, error) {
	statusCode, msg, err := c.ListInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest) string {
	if request == nil {
		request = NewListInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListInstanceWithContextV2(ctx context.Context, request *ListInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewListInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "CreateInstance")
	return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstance(request *CreateInstanceRequest) string {
	return c.CreateInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceSend(request *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	statusCode, msg, err := c.CreateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) string {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInstanceWithContextV2(ctx context.Context, request *CreateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReleaseDBInstanceEipRequest() (request *ReleaseDBInstanceEipRequest) {
	request = &ReleaseDBInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "ReleaseDBInstanceEip")
	return
}

func NewReleaseDBInstanceEipResponse() (response *ReleaseDBInstanceEipResponse) {
	response = &ReleaseDBInstanceEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReleaseDBInstanceEip(request *ReleaseDBInstanceEipRequest) string {
	return c.ReleaseDBInstanceEipWithContext(context.Background(), request)
}

func (c *Client) ReleaseDBInstanceEipSend(request *ReleaseDBInstanceEipRequest) (*ReleaseDBInstanceEipResponse, error) {
	statusCode, msg, err := c.ReleaseDBInstanceEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ReleaseDBInstanceEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReleaseDBInstanceEipWithContext(ctx context.Context, request *ReleaseDBInstanceEipRequest) string {
	if request == nil {
		request = NewReleaseDBInstanceEipRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ReleaseDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewReleaseDBInstanceEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReleaseDBInstanceEipWithContextV2(ctx context.Context, request *ReleaseDBInstanceEipRequest) (int, string, error) {
	if request == nil {
		request = NewReleaseDBInstanceEipRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ReleaseDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewReleaseDBInstanceEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateDBInstanceEipRequest() (request *AllocateDBInstanceEipRequest) {
	request = &AllocateDBInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "AllocateDBInstanceEip")
	return
}

func NewAllocateDBInstanceEipResponse() (response *AllocateDBInstanceEipResponse) {
	response = &AllocateDBInstanceEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateDBInstanceEip(request *AllocateDBInstanceEipRequest) string {
	return c.AllocateDBInstanceEipWithContext(context.Background(), request)
}

func (c *Client) AllocateDBInstanceEipSend(request *AllocateDBInstanceEipRequest) (*AllocateDBInstanceEipResponse, error) {
	statusCode, msg, err := c.AllocateDBInstanceEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AllocateDBInstanceEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateDBInstanceEipWithContext(ctx context.Context, request *AllocateDBInstanceEipRequest) string {
	if request == nil {
		request = NewAllocateDBInstanceEipRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "AllocateDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAllocateDBInstanceEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateDBInstanceEipWithContextV2(ctx context.Context, request *AllocateDBInstanceEipRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateDBInstanceEipRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "AllocateDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAllocateDBInstanceEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListBackupRequest() (request *ListBackupRequest) {
	request = &ListBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "ListBackup")
	return
}

func NewListBackupResponse() (response *ListBackupResponse) {
	response = &ListBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListBackup(request *ListBackupRequest) string {
	return c.ListBackupWithContext(context.Background(), request)
}

func (c *Client) ListBackupSend(request *ListBackupRequest) (*ListBackupResponse, error) {
	statusCode, msg, err := c.ListBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListBackupWithContext(ctx context.Context, request *ListBackupRequest) string {
	if request == nil {
		request = NewListBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListBackupWithContextV2(ctx context.Context, request *ListBackupRequest) (int, string, error) {
	if request == nil {
		request = NewListBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateBackupRequest() (request *CreateBackupRequest) {
	request = &CreateBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "CreateBackup")
	return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
	response = &CreateBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateBackup(request *CreateBackupRequest) string {
	return c.CreateBackupWithContext(context.Background(), request)
}

func (c *Client) CreateBackupSend(request *CreateBackupRequest) (*CreateBackupResponse, error) {
	statusCode, msg, err := c.CreateBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) string {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateBackupWithContextV2(ctx context.Context, request *CreateBackupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "CreateBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
	request = &DeleteBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "DeleteBackup")
	return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
	response = &DeleteBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBackup(request *DeleteBackupRequest) string {
	return c.DeleteBackupWithContext(context.Background(), request)
}

func (c *Client) DeleteBackupSend(request *DeleteBackupRequest) (*DeleteBackupResponse, error) {
	statusCode, msg, err := c.DeleteBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBackupWithContext(ctx context.Context, request *DeleteBackupRequest) string {
	if request == nil {
		request = NewDeleteBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBackupWithContextV2(ctx context.Context, request *DeleteBackupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "DeleteBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListCollectionsRequest() (request *ListCollectionsRequest) {
	request = &ListCollectionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "ListCollections")
	return
}

func NewListCollectionsResponse() (response *ListCollectionsResponse) {
	response = &ListCollectionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListCollections(request *ListCollectionsRequest) string {
	return c.ListCollectionsWithContext(context.Background(), request)
}

func (c *Client) ListCollectionsSend(request *ListCollectionsRequest) (*ListCollectionsResponse, error) {
	statusCode, msg, err := c.ListCollectionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListCollectionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListCollectionsWithContext(ctx context.Context, request *ListCollectionsRequest) string {
	if request == nil {
		request = NewListCollectionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListCollections")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListCollectionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListCollectionsWithContextV2(ctx context.Context, request *ListCollectionsRequest) (int, string, error) {
	if request == nil {
		request = NewListCollectionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "ListCollections")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListCollectionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateInstanceTrialOrderRequest() (request *UpdateInstanceTrialOrderRequest) {
	request = &UpdateInstanceTrialOrderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("milvus", APIVersion, "UpdateInstanceTrialOrder")
	return
}

func NewUpdateInstanceTrialOrderResponse() (response *UpdateInstanceTrialOrderResponse) {
	response = &UpdateInstanceTrialOrderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateInstanceTrialOrder(request *UpdateInstanceTrialOrderRequest) string {
	return c.UpdateInstanceTrialOrderWithContext(context.Background(), request)
}

func (c *Client) UpdateInstanceTrialOrderSend(request *UpdateInstanceTrialOrderRequest) (*UpdateInstanceTrialOrderResponse, error) {
	statusCode, msg, err := c.UpdateInstanceTrialOrderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateInstanceTrialOrderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateInstanceTrialOrderWithContext(ctx context.Context, request *UpdateInstanceTrialOrderRequest) string {
	if request == nil {
		request = NewUpdateInstanceTrialOrderRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "UpdateInstanceTrialOrder")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateInstanceTrialOrderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateInstanceTrialOrderWithContextV2(ctx context.Context, request *UpdateInstanceTrialOrderRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateInstanceTrialOrderRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("milvus", APIVersion, "UpdateInstanceTrialOrder")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateInstanceTrialOrderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
