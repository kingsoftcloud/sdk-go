package v20210520

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-05-20"

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

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "CreateInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "CreateInstance")
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
func NewListInstanceRequest() (request *ListInstanceRequest) {
	request = &ListInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "ListInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "ListInstance")
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
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeInstance")
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
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteInstance")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteInstance")
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
func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
	request = &RenameInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "RenameInstance")
	return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
	response = &RenameInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameInstance(request *RenameInstanceRequest) string {
	return c.RenameInstanceWithContext(context.Background(), request)
}

func (c *Client) RenameInstanceSend(request *RenameInstanceRequest) (*RenameInstanceResponse, error) {
	statusCode, msg, err := c.RenameInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RenameInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenameInstanceWithContext(ctx context.Context, request *RenameInstanceRequest) string {
	if request == nil {
		request = NewRenameInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "RenameInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenameInstanceWithContextV2(ctx context.Context, request *RenameInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRenameInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "RenameInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListRegionRequest() (request *ListRegionRequest) {
	request = &ListRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListRegion")
	return
}

func NewListRegionResponse() (response *ListRegionResponse) {
	response = &ListRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListRegion(request *ListRegionRequest) string {
	return c.ListRegionWithContext(context.Background(), request)
}

func (c *Client) ListRegionSend(request *ListRegionRequest) (*ListRegionResponse, error) {
	statusCode, msg, err := c.ListRegionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListRegionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListRegionWithContext(ctx context.Context, request *ListRegionRequest) string {
	if request == nil {
		request = NewListRegionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListRegion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListRegionWithContextV2(ctx context.Context, request *ListRegionRequest) (int, string, error) {
	if request == nil {
		request = NewListRegionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListRegion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListRegionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescRegionRequest() (request *DescRegionRequest) {
	request = &DescRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescRegion")
	return
}

func NewDescRegionResponse() (response *DescRegionResponse) {
	response = &DescRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescRegion(request *DescRegionRequest) string {
	return c.DescRegionWithContext(context.Background(), request)
}

func (c *Client) DescRegionSend(request *DescRegionRequest) (*DescRegionResponse, error) {
	statusCode, msg, err := c.DescRegionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescRegionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescRegionWithContext(ctx context.Context, request *DescRegionRequest) string {
	if request == nil {
		request = NewDescRegionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescRegion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescRegionWithContextV2(ctx context.Context, request *DescRegionRequest) (int, string, error) {
	if request == nil {
		request = NewDescRegionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescRegion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescRegionResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityGroup")
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
	request.Init().WithApiInfo("tidb", APIVersion, "ListSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "ListSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "ListSecurityGroup")
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
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

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
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityGroup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityGroup")
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
func NewUpdateSecurityGroupRequest() (request *UpdateSecurityGroupRequest) {
	request = &UpdateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityGroup")
	return
}

func NewUpdateSecurityGroupResponse() (response *UpdateSecurityGroupResponse) {
	response = &UpdateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateSecurityGroup(request *UpdateSecurityGroupRequest) string {
	return c.UpdateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) UpdateSecurityGroupSend(request *UpdateSecurityGroupRequest) (*UpdateSecurityGroupResponse, error) {
	statusCode, msg, err := c.UpdateSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateSecurityGroupWithContext(ctx context.Context, request *UpdateSecurityGroupRequest) string {
	if request == nil {
		request = NewUpdateSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateSecurityGroupWithContextV2(ctx context.Context, request *UpdateSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
	request = &CloneSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CloneSecurityGroup")
	return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
	response = &CloneSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) string {
	return c.CloneSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CloneSecurityGroupSend(request *CloneSecurityGroupRequest) (*CloneSecurityGroupResponse, error) {
	statusCode, msg, err := c.CloneSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CloneSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) string {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CloneSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloneSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloneSecurityGroupWithContextV2(ctx context.Context, request *CloneSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CloneSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloneSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBindSecurityGroupRequest() (request *BindSecurityGroupRequest) {
	request = &BindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "BindSecurityGroup")
	return
}

func NewBindSecurityGroupResponse() (response *BindSecurityGroupResponse) {
	response = &BindSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindSecurityGroup(request *BindSecurityGroupRequest) string {
	return c.BindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) BindSecurityGroupSend(request *BindSecurityGroupRequest) (*BindSecurityGroupResponse, error) {
	statusCode, msg, err := c.BindSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BindSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BindSecurityGroupWithContext(ctx context.Context, request *BindSecurityGroupRequest) string {
	if request == nil {
		request = NewBindSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "BindSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBindSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BindSecurityGroupWithContextV2(ctx context.Context, request *BindSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewBindSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "BindSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBindSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnbindSecurityGroupRequest() (request *UnbindSecurityGroupRequest) {
	request = &UnbindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "UnbindSecurityGroup")
	return
}

func NewUnbindSecurityGroupResponse() (response *UnbindSecurityGroupResponse) {
	response = &UnbindSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindSecurityGroup(request *UnbindSecurityGroupRequest) string {
	return c.UnbindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) UnbindSecurityGroupSend(request *UnbindSecurityGroupRequest) (*UnbindSecurityGroupResponse, error) {
	statusCode, msg, err := c.UnbindSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UnbindSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnbindSecurityGroupWithContext(ctx context.Context, request *UnbindSecurityGroupRequest) string {
	if request == nil {
		request = NewUnbindSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UnbindSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUnbindSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnbindSecurityGroupWithContextV2(ctx context.Context, request *UnbindSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewUnbindSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UnbindSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUnbindSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRebindSecurityGroupRequest() (request *RebindSecurityGroupRequest) {
	request = &RebindSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "RebindSecurityGroup")
	return
}

func NewRebindSecurityGroupResponse() (response *RebindSecurityGroupResponse) {
	response = &RebindSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebindSecurityGroup(request *RebindSecurityGroupRequest) string {
	return c.RebindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) RebindSecurityGroupSend(request *RebindSecurityGroupRequest) (*RebindSecurityGroupResponse, error) {
	statusCode, msg, err := c.RebindSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RebindSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RebindSecurityGroupWithContext(ctx context.Context, request *RebindSecurityGroupRequest) string {
	if request == nil {
		request = NewRebindSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "RebindSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRebindSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RebindSecurityGroupWithContextV2(ctx context.Context, request *RebindSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewRebindSecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "RebindSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRebindSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityRuleRequest() (request *CreateSecurityRuleRequest) {
	request = &CreateSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityRule")
	return
}

func NewCreateSecurityRuleResponse() (response *CreateSecurityRuleResponse) {
	response = &CreateSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityRule(request *CreateSecurityRuleRequest) string {
	return c.CreateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityRuleSend(request *CreateSecurityRuleRequest) (*CreateSecurityRuleResponse, error) {
	statusCode, msg, err := c.CreateSecurityRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSecurityRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityRuleWithContext(ctx context.Context, request *CreateSecurityRuleRequest) string {
	if request == nil {
		request = NewCreateSecurityRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityRuleWithContextV2(ctx context.Context, request *CreateSecurityRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateSecurityRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateSecurityRuleRequest() (request *UpdateSecurityRuleRequest) {
	request = &UpdateSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityRule")
	return
}

func NewUpdateSecurityRuleResponse() (response *UpdateSecurityRuleResponse) {
	response = &UpdateSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateSecurityRule(request *UpdateSecurityRuleRequest) string {
	return c.UpdateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) UpdateSecurityRuleSend(request *UpdateSecurityRuleRequest) (*UpdateSecurityRuleResponse, error) {
	statusCode, msg, err := c.UpdateSecurityRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateSecurityRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateSecurityRuleWithContext(ctx context.Context, request *UpdateSecurityRuleRequest) string {
	if request == nil {
		request = NewUpdateSecurityRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateSecurityRuleWithContextV2(ctx context.Context, request *UpdateSecurityRuleRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateSecurityRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateSecurityRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityRuleRequest() (request *DeleteSecurityRuleRequest) {
	request = &DeleteSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityRule")
	return
}

func NewDeleteSecurityRuleResponse() (response *DeleteSecurityRuleResponse) {
	response = &DeleteSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityRule(request *DeleteSecurityRuleRequest) string {
	return c.DeleteSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityRuleSend(request *DeleteSecurityRuleRequest) (*DeleteSecurityRuleResponse, error) {
	statusCode, msg, err := c.DeleteSecurityRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteSecurityRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityRuleWithContext(ctx context.Context, request *DeleteSecurityRuleRequest) string {
	if request == nil {
		request = NewDeleteSecurityRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityRuleWithContextV2(ctx context.Context, request *DeleteSecurityRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteSecurityRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSecurityRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListSecuredInstanceRequest() (request *ListSecuredInstanceRequest) {
	request = &ListSecuredInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListSecuredInstance")
	return
}

func NewListSecuredInstanceResponse() (response *ListSecuredInstanceResponse) {
	response = &ListSecuredInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSecuredInstance(request *ListSecuredInstanceRequest) string {
	return c.ListSecuredInstanceWithContext(context.Background(), request)
}

func (c *Client) ListSecuredInstanceSend(request *ListSecuredInstanceRequest) (*ListSecuredInstanceResponse, error) {
	statusCode, msg, err := c.ListSecuredInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListSecuredInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListSecuredInstanceWithContext(ctx context.Context, request *ListSecuredInstanceRequest) string {
	if request == nil {
		request = NewListSecuredInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListSecuredInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecuredInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListSecuredInstanceWithContextV2(ctx context.Context, request *ListSecuredInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewListSecuredInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListSecuredInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecuredInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListUnsecuredInstanceRequest() (request *ListUnsecuredInstanceRequest) {
	request = &ListUnsecuredInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListUnsecuredInstance")
	return
}

func NewListUnsecuredInstanceResponse() (response *ListUnsecuredInstanceResponse) {
	response = &ListUnsecuredInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListUnsecuredInstance(request *ListUnsecuredInstanceRequest) string {
	return c.ListUnsecuredInstanceWithContext(context.Background(), request)
}

func (c *Client) ListUnsecuredInstanceSend(request *ListUnsecuredInstanceRequest) (*ListUnsecuredInstanceResponse, error) {
	statusCode, msg, err := c.ListUnsecuredInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListUnsecuredInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListUnsecuredInstanceWithContext(ctx context.Context, request *ListUnsecuredInstanceRequest) string {
	if request == nil {
		request = NewListUnsecuredInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListUnsecuredInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListUnsecuredInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListUnsecuredInstanceWithContextV2(ctx context.Context, request *ListUnsecuredInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewListUnsecuredInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListUnsecuredInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListUnsecuredInstanceResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "ListBackup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "ListBackup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "ListBackup")
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
	request.Init().WithApiInfo("tidb", APIVersion, "CreateBackup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "CreateBackup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "CreateBackup")
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
func NewUpdateBackupRuleRequest() (request *UpdateBackupRuleRequest) {
	request = &UpdateBackupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateBackupRule")
	return
}

func NewUpdateBackupRuleResponse() (response *UpdateBackupRuleResponse) {
	response = &UpdateBackupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateBackupRule(request *UpdateBackupRuleRequest) string {
	return c.UpdateBackupRuleWithContext(context.Background(), request)
}

func (c *Client) UpdateBackupRuleSend(request *UpdateBackupRuleRequest) (*UpdateBackupRuleResponse, error) {
	statusCode, msg, err := c.UpdateBackupRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateBackupRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateBackupRuleWithContext(ctx context.Context, request *UpdateBackupRuleRequest) string {
	if request == nil {
		request = NewUpdateBackupRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateBackupRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateBackupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateBackupRuleWithContextV2(ctx context.Context, request *UpdateBackupRuleRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateBackupRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateBackupRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateBackupRuleResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteBackup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteBackup")
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
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteBackup")
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
func NewCreateRestoreRequest() (request *CreateRestoreRequest) {
	request = &CreateRestoreRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateRestore")
	return
}

func NewCreateRestoreResponse() (response *CreateRestoreResponse) {
	response = &CreateRestoreResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRestore(request *CreateRestoreRequest) string {
	return c.CreateRestoreWithContext(context.Background(), request)
}

func (c *Client) CreateRestoreSend(request *CreateRestoreRequest) (*CreateRestoreResponse, error) {
	statusCode, msg, err := c.CreateRestoreWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateRestoreResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRestoreWithContext(ctx context.Context, request *CreateRestoreRequest) string {
	if request == nil {
		request = NewCreateRestoreRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateRestore")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateRestoreResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRestoreWithContextV2(ctx context.Context, request *CreateRestoreRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRestoreRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateRestore")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateRestoreResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOpenTiMonitorRequest() (request *OpenTiMonitorRequest) {
	request = &OpenTiMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "OpenTiMonitor")
	return
}

func NewOpenTiMonitorResponse() (response *OpenTiMonitorResponse) {
	response = &OpenTiMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OpenTiMonitor(request *OpenTiMonitorRequest) string {
	return c.OpenTiMonitorWithContext(context.Background(), request)
}

func (c *Client) OpenTiMonitorSend(request *OpenTiMonitorRequest) (*OpenTiMonitorResponse, error) {
	statusCode, msg, err := c.OpenTiMonitorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OpenTiMonitorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OpenTiMonitorWithContext(ctx context.Context, request *OpenTiMonitorRequest) string {
	if request == nil {
		request = NewOpenTiMonitorRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "OpenTiMonitor")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOpenTiMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OpenTiMonitorWithContextV2(ctx context.Context, request *OpenTiMonitorRequest) (int, string, error) {
	if request == nil {
		request = NewOpenTiMonitorRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "OpenTiMonitor")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOpenTiMonitorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateTask")
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

func (c *Client) CreateTaskSend(request *CreateTaskRequest) (*CreateTaskResponse, error) {
	statusCode, msg, err := c.CreateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) string {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateTask")
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

func (c *Client) CreateTaskWithContextV2(ctx context.Context, request *CreateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOperationTasksRequest() (request *OperationTasksRequest) {
	request = &OperationTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "OperationTasks")
	return
}

func NewOperationTasksResponse() (response *OperationTasksResponse) {
	response = &OperationTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperationTasks(request *OperationTasksRequest) string {
	return c.OperationTasksWithContext(context.Background(), request)
}

func (c *Client) OperationTasksSend(request *OperationTasksRequest) (*OperationTasksResponse, error) {
	statusCode, msg, err := c.OperationTasksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OperationTasksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OperationTasksWithContext(ctx context.Context, request *OperationTasksRequest) string {
	if request == nil {
		request = NewOperationTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "OperationTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperationTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OperationTasksWithContextV2(ctx context.Context, request *OperationTasksRequest) (int, string, error) {
	if request == nil {
		request = NewOperationTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "OperationTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperationTasksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCheckTaskNameRequest() (request *CheckTaskNameRequest) {
	request = &CheckTaskNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CheckTaskName")
	return
}

func NewCheckTaskNameResponse() (response *CheckTaskNameResponse) {
	response = &CheckTaskNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckTaskName(request *CheckTaskNameRequest) string {
	return c.CheckTaskNameWithContext(context.Background(), request)
}

func (c *Client) CheckTaskNameSend(request *CheckTaskNameRequest) (*CheckTaskNameResponse, error) {
	statusCode, msg, err := c.CheckTaskNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CheckTaskNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CheckTaskNameWithContext(ctx context.Context, request *CheckTaskNameRequest) string {
	if request == nil {
		request = NewCheckTaskNameRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CheckTaskName")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckTaskNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CheckTaskNameWithContextV2(ctx context.Context, request *CheckTaskNameRequest) (int, string, error) {
	if request == nil {
		request = NewCheckTaskNameRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CheckTaskName")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckTaskNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeTask")
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

func (c *Client) DescribeTaskSend(request *DescribeTaskRequest) (*DescribeTaskResponse, error) {
	statusCode, msg, err := c.DescribeTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) string {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeTask")
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

func (c *Client) DescribeTaskWithContextV2(ctx context.Context, request *DescribeTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTasksRequest() (request *ListTasksRequest) {
	request = &ListTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ListTasks")
	return
}

func NewListTasksResponse() (response *ListTasksResponse) {
	response = &ListTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTasks(request *ListTasksRequest) string {
	return c.ListTasksWithContext(context.Background(), request)
}

func (c *Client) ListTasksSend(request *ListTasksRequest) (*ListTasksResponse, error) {
	statusCode, msg, err := c.ListTasksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListTasksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) string {
	if request == nil {
		request = NewListTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTasksWithContextV2(ctx context.Context, request *ListTasksRequest) (int, string, error) {
	if request == nil {
		request = NewListTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ListTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTasksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
	request = &DescribeDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeDatabases")
	return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
	response = &DescribeDatabasesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) string {
	return c.DescribeDatabasesWithContext(context.Background(), request)
}

func (c *Client) DescribeDatabasesSend(request *DescribeDatabasesRequest) (*DescribeDatabasesResponse, error) {
	statusCode, msg, err := c.DescribeDatabasesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDatabasesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) string {
	if request == nil {
		request = NewDescribeDatabasesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeDatabases")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDatabasesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDatabasesWithContextV2(ctx context.Context, request *DescribeDatabasesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDatabasesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeDatabases")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDatabasesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DescribeAccounts")
	return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) string {
	return c.DescribeAccountsWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountsSend(request *DescribeAccountsRequest) (*DescribeAccountsResponse, error) {
	statusCode, msg, err := c.DescribeAccountsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAccountsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) string {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeAccounts")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAccountsWithContextV2(ctx context.Context, request *DescribeAccountsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DescribeAccounts")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "CreateAccount")
	return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAccount(request *CreateAccountRequest) string {
	return c.CreateAccountWithContext(context.Background(), request)
}

func (c *Client) CreateAccountSend(request *CreateAccountRequest) (*CreateAccountResponse, error) {
	statusCode, msg, err := c.CreateAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) string {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAccountWithContextV2(ctx context.Context, request *CreateAccountRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "CreateAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
	request = &DeleteAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "DeleteAccount")
	return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
	response = &DeleteAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAccount(request *DeleteAccountRequest) string {
	return c.DeleteAccountWithContext(context.Background(), request)
}

func (c *Client) DeleteAccountSend(request *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	statusCode, msg, err := c.DeleteAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) string {
	if request == nil {
		request = NewDeleteAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAccountWithContextV2(ctx context.Context, request *DeleteAccountRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "DeleteAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAccountInfoRequest() (request *ModifyAccountInfoRequest) {
	request = &ModifyAccountInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountInfo")
	return
}

func NewModifyAccountInfoResponse() (response *ModifyAccountInfoResponse) {
	response = &ModifyAccountInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAccountInfo(request *ModifyAccountInfoRequest) string {
	return c.ModifyAccountInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyAccountInfoSend(request *ModifyAccountInfoRequest) (*ModifyAccountInfoResponse, error) {
	statusCode, msg, err := c.ModifyAccountInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAccountInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAccountInfoWithContext(ctx context.Context, request *ModifyAccountInfoRequest) string {
	if request == nil {
		request = NewModifyAccountInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAccountInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAccountInfoWithContextV2(ctx context.Context, request *ModifyAccountInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAccountInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAccountInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
	request = &ModifyAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountPrivileges")
	return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
	response = &ModifyAccountPrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) string {
	return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyAccountPrivilegesSend(request *ModifyAccountPrivilegesRequest) (*ModifyAccountPrivilegesResponse, error) {
	statusCode, msg, err := c.ModifyAccountPrivilegesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAccountPrivilegesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) string {
	if request == nil {
		request = NewModifyAccountPrivilegesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountPrivileges")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAccountPrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAccountPrivilegesWithContextV2(ctx context.Context, request *ModifyAccountPrivilegesRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAccountPrivilegesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ModifyAccountPrivileges")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAccountPrivilegesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewConfigurationInstanceEipRequest() (request *ConfigurationInstanceEipRequest) {
	request = &ConfigurationInstanceEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tidb", APIVersion, "ConfigurationInstanceEip")
	return
}

func NewConfigurationInstanceEipResponse() (response *ConfigurationInstanceEipResponse) {
	response = &ConfigurationInstanceEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ConfigurationInstanceEip(request *ConfigurationInstanceEipRequest) string {
	return c.ConfigurationInstanceEipWithContext(context.Background(), request)
}

func (c *Client) ConfigurationInstanceEipSend(request *ConfigurationInstanceEipRequest) (*ConfigurationInstanceEipResponse, error) {
	statusCode, msg, err := c.ConfigurationInstanceEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ConfigurationInstanceEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ConfigurationInstanceEipWithContext(ctx context.Context, request *ConfigurationInstanceEipRequest) string {
	if request == nil {
		request = NewConfigurationInstanceEipRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ConfigurationInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewConfigurationInstanceEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ConfigurationInstanceEipWithContextV2(ctx context.Context, request *ConfigurationInstanceEipRequest) (int, string, error) {
	if request == nil {
		request = NewConfigurationInstanceEipRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("tidb", APIVersion, "ConfigurationInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewConfigurationInstanceEipResponse()
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
	request.Init().WithApiInfo("tidb", APIVersion, "UpdateInstanceTrialOrder")
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
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateInstanceTrialOrder")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

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
		request.Init().WithApiInfo("tidb", APIVersion, "UpdateInstanceTrialOrder")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateInstanceTrialOrderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
