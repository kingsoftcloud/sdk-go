package v20160701

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-07-01"

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

func NewRebootDBInstanceRequest() (request *RebootDBInstanceRequest) {
	request = &RebootDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "RebootDBInstance")
	return
}

func NewRebootDBInstanceResponse() (response *RebootDBInstanceResponse) {
	response = &RebootDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootDBInstance(request *RebootDBInstanceRequest) string {
	return c.RebootDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RebootDBInstanceSend(request *RebootDBInstanceRequest) (*RebootDBInstanceResponse, error) {
	statusCode, msg, err := c.RebootDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RebootDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RebootDBInstanceWithContext(ctx context.Context, request *RebootDBInstanceRequest) string {
	if request == nil {
		request = NewRebootDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RebootDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RebootDBInstanceWithContextV2(ctx context.Context, request *RebootDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRebootDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RebootDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBParameterGroupRequest() (request *ModifyDBParameterGroupRequest) {
	request = &ModifyDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBParameterGroup")
	return
}

func NewModifyDBParameterGroupResponse() (response *ModifyDBParameterGroupResponse) {
	response = &ModifyDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBParameterGroup(request *ModifyDBParameterGroupRequest) string {
	return c.ModifyDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyDBParameterGroupSend(request *ModifyDBParameterGroupRequest) (*ModifyDBParameterGroupResponse, error) {
	statusCode, msg, err := c.ModifyDBParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBParameterGroupWithContext(ctx context.Context, request *ModifyDBParameterGroupRequest) string {
	if request == nil {
		request = NewModifyDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBParameterGroupWithContextV2(ctx context.Context, request *ModifyDBParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResetDBParameterGroupRequest() (request *ResetDBParameterGroupRequest) {
	request = &ResetDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ResetDBParameterGroup")
	return
}

func NewResetDBParameterGroupResponse() (response *ResetDBParameterGroupResponse) {
	response = &ResetDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetDBParameterGroup(request *ResetDBParameterGroupRequest) string {
	return c.ResetDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ResetDBParameterGroupSend(request *ResetDBParameterGroupRequest) (*ResetDBParameterGroupResponse, error) {
	statusCode, msg, err := c.ResetDBParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ResetDBParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResetDBParameterGroupWithContext(ctx context.Context, request *ResetDBParameterGroupRequest) string {
	if request == nil {
		request = NewResetDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ResetDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResetDBParameterGroupWithContextV2(ctx context.Context, request *ResetDBParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewResetDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ResetDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetDBParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBParameterGroupRequest() (request *DescribeDBParameterGroupRequest) {
	request = &DescribeDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBParameterGroup")
	return
}

func NewDescribeDBParameterGroupResponse() (response *DescribeDBParameterGroupResponse) {
	response = &DescribeDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBParameterGroup(request *DescribeDBParameterGroupRequest) string {
	return c.DescribeDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeDBParameterGroupSend(request *DescribeDBParameterGroupRequest) (*DescribeDBParameterGroupResponse, error) {
	statusCode, msg, err := c.DescribeDBParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBParameterGroupWithContext(ctx context.Context, request *DescribeDBParameterGroupRequest) string {
	if request == nil {
		request = NewDescribeDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBParameterGroupWithContextV2(ctx context.Context, request *DescribeDBParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEngineDefaultParametersRequest() (request *DescribeEngineDefaultParametersRequest) {
	request = &DescribeEngineDefaultParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineDefaultParameters")
	return
}

func NewDescribeEngineDefaultParametersResponse() (response *DescribeEngineDefaultParametersResponse) {
	response = &DescribeEngineDefaultParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEngineDefaultParameters(request *DescribeEngineDefaultParametersRequest) string {
	return c.DescribeEngineDefaultParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeEngineDefaultParametersSend(request *DescribeEngineDefaultParametersRequest) (*DescribeEngineDefaultParametersResponse, error) {
	statusCode, msg, err := c.DescribeEngineDefaultParametersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeEngineDefaultParametersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEngineDefaultParametersWithContext(ctx context.Context, request *DescribeEngineDefaultParametersRequest) string {
	if request == nil {
		request = NewDescribeEngineDefaultParametersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineDefaultParameters")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEngineDefaultParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEngineDefaultParametersWithContextV2(ctx context.Context, request *DescribeEngineDefaultParametersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEngineDefaultParametersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineDefaultParameters")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEngineDefaultParametersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDBParameterGroupRequest() (request *CreateDBParameterGroupRequest) {
	request = &CreateDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateDBParameterGroup")
	return
}

func NewCreateDBParameterGroupResponse() (response *CreateDBParameterGroupResponse) {
	response = &CreateDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBParameterGroup(request *CreateDBParameterGroupRequest) string {
	return c.CreateDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) CreateDBParameterGroupSend(request *CreateDBParameterGroupRequest) (*CreateDBParameterGroupResponse, error) {
	statusCode, msg, err := c.CreateDBParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDBParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDBParameterGroupWithContext(ctx context.Context, request *CreateDBParameterGroupRequest) string {
	if request == nil {
		request = NewCreateDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDBParameterGroupWithContextV2(ctx context.Context, request *CreateDBParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDBParameterGroupRequest() (request *DeleteDBParameterGroupRequest) {
	request = &DeleteDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DeleteDBParameterGroup")
	return
}

func NewDeleteDBParameterGroupResponse() (response *DeleteDBParameterGroupResponse) {
	response = &DeleteDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDBParameterGroup(request *DeleteDBParameterGroupRequest) string {
	return c.DeleteDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteDBParameterGroupSend(request *DeleteDBParameterGroupRequest) (*DeleteDBParameterGroupResponse, error) {
	statusCode, msg, err := c.DeleteDBParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDBParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDBParameterGroupWithContext(ctx context.Context, request *DeleteDBParameterGroupRequest) string {
	if request == nil {
		request = NewDeleteDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDBParameterGroupWithContextV2(ctx context.Context, request *DeleteDBParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDBParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstance")
	return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) string {
	return c.CreateDBInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateDBInstanceSend(request *CreateDBInstanceRequest) (*CreateDBInstanceResponse, error) {
	statusCode, msg, err := c.CreateDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) string {
	if request == nil {
		request = NewCreateDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDBInstanceWithContextV2(ctx context.Context, request *CreateDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreDBInstanceFromDBBackupRequest() (request *RestoreDBInstanceFromDBBackupRequest) {
	request = &RestoreDBInstanceFromDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceFromDBBackup")
	return
}

func NewRestoreDBInstanceFromDBBackupResponse() (response *RestoreDBInstanceFromDBBackupResponse) {
	response = &RestoreDBInstanceFromDBBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreDBInstanceFromDBBackup(request *RestoreDBInstanceFromDBBackupRequest) string {
	return c.RestoreDBInstanceFromDBBackupWithContext(context.Background(), request)
}

func (c *Client) RestoreDBInstanceFromDBBackupSend(request *RestoreDBInstanceFromDBBackupRequest) (*RestoreDBInstanceFromDBBackupResponse, error) {
	statusCode, msg, err := c.RestoreDBInstanceFromDBBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreDBInstanceFromDBBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreDBInstanceFromDBBackupWithContext(ctx context.Context, request *RestoreDBInstanceFromDBBackupRequest) string {
	if request == nil {
		request = NewRestoreDBInstanceFromDBBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceFromDBBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestoreDBInstanceFromDBBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreDBInstanceFromDBBackupWithContextV2(ctx context.Context, request *RestoreDBInstanceFromDBBackupRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreDBInstanceFromDBBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceFromDBBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestoreDBInstanceFromDBBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
	request = &DeleteDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DeleteDBInstance")
	return
}

func NewDeleteDBInstanceResponse() (response *DeleteDBInstanceResponse) {
	response = &DeleteDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) string {
	return c.DeleteDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteDBInstanceSend(request *DeleteDBInstanceRequest) (*DeleteDBInstanceResponse, error) {
	statusCode, msg, err := c.DeleteDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDBInstanceWithContext(ctx context.Context, request *DeleteDBInstanceRequest) string {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDBInstanceWithContextV2(ctx context.Context, request *DeleteDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDBInstanceReadReplicaRequest() (request *CreateDBInstanceReadReplicaRequest) {
	request = &CreateDBInstanceReadReplicaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstanceReadReplica")
	return
}

func NewCreateDBInstanceReadReplicaResponse() (response *CreateDBInstanceReadReplicaResponse) {
	response = &CreateDBInstanceReadReplicaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBInstanceReadReplica(request *CreateDBInstanceReadReplicaRequest) string {
	return c.CreateDBInstanceReadReplicaWithContext(context.Background(), request)
}

func (c *Client) CreateDBInstanceReadReplicaSend(request *CreateDBInstanceReadReplicaRequest) (*CreateDBInstanceReadReplicaResponse, error) {
	statusCode, msg, err := c.CreateDBInstanceReadReplicaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDBInstanceReadReplicaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDBInstanceReadReplicaWithContext(ctx context.Context, request *CreateDBInstanceReadReplicaRequest) string {
	if request == nil {
		request = NewCreateDBInstanceReadReplicaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstanceReadReplica")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBInstanceReadReplicaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDBInstanceReadReplicaWithContextV2(ctx context.Context, request *CreateDBInstanceReadReplicaRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDBInstanceReadReplicaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstanceReadReplica")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBInstanceReadReplicaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreDBInstanceToPointInTimeRequest() (request *RestoreDBInstanceToPointInTimeRequest) {
	request = &RestoreDBInstanceToPointInTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceToPointInTime")
	return
}

func NewRestoreDBInstanceToPointInTimeResponse() (response *RestoreDBInstanceToPointInTimeResponse) {
	response = &RestoreDBInstanceToPointInTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreDBInstanceToPointInTime(request *RestoreDBInstanceToPointInTimeRequest) string {
	return c.RestoreDBInstanceToPointInTimeWithContext(context.Background(), request)
}

func (c *Client) RestoreDBInstanceToPointInTimeSend(request *RestoreDBInstanceToPointInTimeRequest) (*RestoreDBInstanceToPointInTimeResponse, error) {
	statusCode, msg, err := c.RestoreDBInstanceToPointInTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreDBInstanceToPointInTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreDBInstanceToPointInTimeWithContext(ctx context.Context, request *RestoreDBInstanceToPointInTimeRequest) string {
	if request == nil {
		request = NewRestoreDBInstanceToPointInTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceToPointInTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestoreDBInstanceToPointInTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreDBInstanceToPointInTimeWithContextV2(ctx context.Context, request *RestoreDBInstanceToPointInTimeRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreDBInstanceToPointInTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceToPointInTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestoreDBInstanceToPointInTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBInstanceRestorableTimeRequest() (request *DescribeDBInstanceRestorableTimeRequest) {
	request = &DescribeDBInstanceRestorableTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceRestorableTime")
	return
}

func NewDescribeDBInstanceRestorableTimeResponse() (response *DescribeDBInstanceRestorableTimeResponse) {
	response = &DescribeDBInstanceRestorableTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstanceRestorableTime(request *DescribeDBInstanceRestorableTimeRequest) string {
	return c.DescribeDBInstanceRestorableTimeWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceRestorableTimeSend(request *DescribeDBInstanceRestorableTimeRequest) (*DescribeDBInstanceRestorableTimeResponse, error) {
	statusCode, msg, err := c.DescribeDBInstanceRestorableTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBInstanceRestorableTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBInstanceRestorableTimeWithContext(ctx context.Context, request *DescribeDBInstanceRestorableTimeRequest) string {
	if request == nil {
		request = NewDescribeDBInstanceRestorableTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceRestorableTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstanceRestorableTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBInstanceRestorableTimeWithContextV2(ctx context.Context, request *DescribeDBInstanceRestorableTimeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBInstanceRestorableTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceRestorableTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstanceRestorableTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceRequest() (request *ModifyDBInstanceRequest) {
	request = &ModifyDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstance")
	return
}

func NewModifyDBInstanceResponse() (response *ModifyDBInstanceResponse) {
	response = &ModifyDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstance(request *ModifyDBInstanceRequest) string {
	return c.ModifyDBInstanceWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceSend(request *ModifyDBInstanceRequest) (*ModifyDBInstanceResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceWithContext(ctx context.Context, request *ModifyDBInstanceRequest) string {
	if request == nil {
		request = NewModifyDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceWithContextV2(ctx context.Context, request *ModifyDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
	request = &DescribeDBLogFilesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBLogFiles")
	return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
	response = &DescribeDBLogFilesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) string {
	return c.DescribeDBLogFilesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBLogFilesSend(request *DescribeDBLogFilesRequest) (*DescribeDBLogFilesResponse, error) {
	statusCode, msg, err := c.DescribeDBLogFilesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBLogFilesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBLogFilesWithContext(ctx context.Context, request *DescribeDBLogFilesRequest) string {
	if request == nil {
		request = NewDescribeDBLogFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBLogFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBLogFilesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBLogFilesWithContextV2(ctx context.Context, request *DescribeDBLogFilesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBLogFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBLogFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBLogFilesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
	request = &DescribeDBBackupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackups")
	return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
	response = &DescribeDBBackupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) string {
	return c.DescribeDBBackupsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBBackupsSend(request *DescribeDBBackupsRequest) (*DescribeDBBackupsResponse, error) {
	statusCode, msg, err := c.DescribeDBBackupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBBackupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBBackupsWithContext(ctx context.Context, request *DescribeDBBackupsRequest) string {
	if request == nil {
		request = NewDescribeDBBackupsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackups")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBBackupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBBackupsWithContextV2(ctx context.Context, request *DescribeDBBackupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBBackupsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackups")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBBackupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
	request = &ModifyDBInstanceSpecRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceSpec")
	return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
	response = &ModifyDBInstanceSpecResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) string {
	return c.ModifyDBInstanceSpecWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceSpecSend(request *ModifyDBInstanceSpecRequest) (*ModifyDBInstanceSpecResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceSpecWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceSpecResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest) string {
	if request == nil {
		request = NewModifyDBInstanceSpecRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceSpec")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceSpecResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceSpecWithContextV2(ctx context.Context, request *ModifyDBInstanceSpecRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceSpecRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceSpec")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceSpecResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstances")
	return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
	response = &DescribeDBInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) string {
	return c.DescribeDBInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstancesSend(request *DescribeDBInstancesRequest) (*DescribeDBInstancesResponse, error) {
	statusCode, msg, err := c.DescribeDBInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) string {
	if request == nil {
		request = NewDescribeDBInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBInstancesWithContextV2(ctx context.Context, request *DescribeDBInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOverrideDBInstanceRequest() (request *OverrideDBInstanceRequest) {
	request = &OverrideDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstance")
	return
}

func NewOverrideDBInstanceResponse() (response *OverrideDBInstanceResponse) {
	response = &OverrideDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OverrideDBInstance(request *OverrideDBInstanceRequest) string {
	return c.OverrideDBInstanceWithContext(context.Background(), request)
}

func (c *Client) OverrideDBInstanceSend(request *OverrideDBInstanceRequest) (*OverrideDBInstanceResponse, error) {
	statusCode, msg, err := c.OverrideDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OverrideDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OverrideDBInstanceWithContext(ctx context.Context, request *OverrideDBInstanceRequest) string {
	if request == nil {
		request = NewOverrideDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOverrideDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OverrideDBInstanceWithContextV2(ctx context.Context, request *OverrideDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewOverrideDBInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOverrideDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBEngineVersionsRequest() (request *DescribeDBEngineVersionsRequest) {
	request = &DescribeDBEngineVersionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBEngineVersions")
	return
}

func NewDescribeDBEngineVersionsResponse() (response *DescribeDBEngineVersionsResponse) {
	response = &DescribeDBEngineVersionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBEngineVersions(request *DescribeDBEngineVersionsRequest) string {
	return c.DescribeDBEngineVersionsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBEngineVersionsSend(request *DescribeDBEngineVersionsRequest) (*DescribeDBEngineVersionsResponse, error) {
	statusCode, msg, err := c.DescribeDBEngineVersionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBEngineVersionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBEngineVersionsWithContext(ctx context.Context, request *DescribeDBEngineVersionsRequest) string {
	if request == nil {
		request = NewDescribeDBEngineVersionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBEngineVersions")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBEngineVersionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBEngineVersionsWithContextV2(ctx context.Context, request *DescribeDBEngineVersionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBEngineVersionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBEngineVersions")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBEngineVersionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpgradeDBInstanceEngineVersionRequest() (request *UpgradeDBInstanceEngineVersionRequest) {
	request = &UpgradeDBInstanceEngineVersionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceEngineVersion")
	return
}

func NewUpgradeDBInstanceEngineVersionResponse() (response *UpgradeDBInstanceEngineVersionResponse) {
	response = &UpgradeDBInstanceEngineVersionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) string {
	return c.UpgradeDBInstanceEngineVersionWithContext(context.Background(), request)
}

func (c *Client) UpgradeDBInstanceEngineVersionSend(request *UpgradeDBInstanceEngineVersionRequest) (*UpgradeDBInstanceEngineVersionResponse, error) {
	statusCode, msg, err := c.UpgradeDBInstanceEngineVersionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpgradeDBInstanceEngineVersionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpgradeDBInstanceEngineVersionWithContext(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest) string {
	if request == nil {
		request = NewUpgradeDBInstanceEngineVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceEngineVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpgradeDBInstanceEngineVersionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpgradeDBInstanceEngineVersionWithContextV2(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest) (int, string, error) {
	if request == nil {
		request = NewUpgradeDBInstanceEngineVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceEngineVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpgradeDBInstanceEngineVersionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceTypeRequest() (request *ModifyDBInstanceTypeRequest) {
	request = &ModifyDBInstanceTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceType")
	return
}

func NewModifyDBInstanceTypeResponse() (response *ModifyDBInstanceTypeResponse) {
	response = &ModifyDBInstanceTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstanceType(request *ModifyDBInstanceTypeRequest) string {
	return c.ModifyDBInstanceTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceTypeSend(request *ModifyDBInstanceTypeRequest) (*ModifyDBInstanceTypeResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceTypeWithContext(ctx context.Context, request *ModifyDBInstanceTypeRequest) string {
	if request == nil {
		request = NewModifyDBInstanceTypeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceType")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceTypeWithContextV2(ctx context.Context, request *ModifyDBInstanceTypeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceTypeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceType")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBInstanceParametersRequest() (request *DescribeDBInstanceParametersRequest) {
	request = &DescribeDBInstanceParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceParameters")
	return
}

func NewDescribeDBInstanceParametersResponse() (response *DescribeDBInstanceParametersResponse) {
	response = &DescribeDBInstanceParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstanceParameters(request *DescribeDBInstanceParametersRequest) string {
	return c.DescribeDBInstanceParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceParametersSend(request *DescribeDBInstanceParametersRequest) (*DescribeDBInstanceParametersResponse, error) {
	statusCode, msg, err := c.DescribeDBInstanceParametersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBInstanceParametersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBInstanceParametersWithContext(ctx context.Context, request *DescribeDBInstanceParametersRequest) string {
	if request == nil {
		request = NewDescribeDBInstanceParametersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceParameters")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstanceParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBInstanceParametersWithContextV2(ctx context.Context, request *DescribeDBInstanceParametersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBInstanceParametersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceParameters")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstanceParametersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBBackupPolicyRequest() (request *ModifyDBBackupPolicyRequest) {
	request = &ModifyDBBackupPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBBackupPolicy")
	return
}

func NewModifyDBBackupPolicyResponse() (response *ModifyDBBackupPolicyResponse) {
	response = &ModifyDBBackupPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBBackupPolicy(request *ModifyDBBackupPolicyRequest) string {
	return c.ModifyDBBackupPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyDBBackupPolicySend(request *ModifyDBBackupPolicyRequest) (*ModifyDBBackupPolicyResponse, error) {
	statusCode, msg, err := c.ModifyDBBackupPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBBackupPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBBackupPolicyWithContext(ctx context.Context, request *ModifyDBBackupPolicyRequest) string {
	if request == nil {
		request = NewModifyDBBackupPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBBackupPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBBackupPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBBackupPolicyWithContextV2(ctx context.Context, request *ModifyDBBackupPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBBackupPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBBackupPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBBackupPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBBackupPolicyRequest() (request *DescribeDBBackupPolicyRequest) {
	request = &DescribeDBBackupPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackupPolicy")
	return
}

func NewDescribeDBBackupPolicyResponse() (response *DescribeDBBackupPolicyResponse) {
	response = &DescribeDBBackupPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBBackupPolicy(request *DescribeDBBackupPolicyRequest) string {
	return c.DescribeDBBackupPolicyWithContext(context.Background(), request)
}

func (c *Client) DescribeDBBackupPolicySend(request *DescribeDBBackupPolicyRequest) (*DescribeDBBackupPolicyResponse, error) {
	statusCode, msg, err := c.DescribeDBBackupPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBBackupPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBBackupPolicyWithContext(ctx context.Context, request *DescribeDBBackupPolicyRequest) string {
	if request == nil {
		request = NewDescribeDBBackupPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackupPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBBackupPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBBackupPolicyWithContextV2(ctx context.Context, request *DescribeDBBackupPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBBackupPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackupPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBBackupPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDBBackupRequest() (request *DeleteDBBackupRequest) {
	request = &DeleteDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DeleteDBBackup")
	return
}

func NewDeleteDBBackupResponse() (response *DeleteDBBackupResponse) {
	response = &DeleteDBBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDBBackup(request *DeleteDBBackupRequest) string {
	return c.DeleteDBBackupWithContext(context.Background(), request)
}

func (c *Client) DeleteDBBackupSend(request *DeleteDBBackupRequest) (*DeleteDBBackupResponse, error) {
	statusCode, msg, err := c.DeleteDBBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDBBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDBBackupWithContext(ctx context.Context, request *DeleteDBBackupRequest) string {
	if request == nil {
		request = NewDeleteDBBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteDBBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDBBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDBBackupWithContextV2(ctx context.Context, request *DeleteDBBackupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDBBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteDBBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDBBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDBBackupRequest() (request *CreateDBBackupRequest) {
	request = &CreateDBBackupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateDBBackup")
	return
}

func NewCreateDBBackupResponse() (response *CreateDBBackupResponse) {
	response = &CreateDBBackupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDBBackup(request *CreateDBBackupRequest) string {
	return c.CreateDBBackupWithContext(context.Background(), request)
}

func (c *Client) CreateDBBackupSend(request *CreateDBBackupRequest) (*CreateDBBackupResponse, error) {
	statusCode, msg, err := c.CreateDBBackupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDBBackupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDBBackupWithContext(ctx context.Context, request *CreateDBBackupRequest) string {
	if request == nil {
		request = NewCreateDBBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBBackupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDBBackupWithContextV2(ctx context.Context, request *CreateDBBackupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDBBackupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateDBBackup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDBBackupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSwitchDBInstanceHARequest() (request *SwitchDBInstanceHARequest) {
	request = &SwitchDBInstanceHARequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SwitchDBInstanceHA")
	return
}

func NewSwitchDBInstanceHAResponse() (response *SwitchDBInstanceHAResponse) {
	response = &SwitchDBInstanceHAResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) string {
	return c.SwitchDBInstanceHAWithContext(context.Background(), request)
}

func (c *Client) SwitchDBInstanceHASend(request *SwitchDBInstanceHARequest) (*SwitchDBInstanceHAResponse, error) {
	statusCode, msg, err := c.SwitchDBInstanceHAWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SwitchDBInstanceHAResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SwitchDBInstanceHAWithContext(ctx context.Context, request *SwitchDBInstanceHARequest) string {
	if request == nil {
		request = NewSwitchDBInstanceHARequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SwitchDBInstanceHA")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSwitchDBInstanceHAResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SwitchDBInstanceHAWithContextV2(ctx context.Context, request *SwitchDBInstanceHARequest) (int, string, error) {
	if request == nil {
		request = NewSwitchDBInstanceHARequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SwitchDBInstanceHA")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSwitchDBInstanceHAResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGenerateDBAdminURLRequest() (request *GenerateDBAdminURLRequest) {
	request = &GenerateDBAdminURLRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "GenerateDBAdminURL")
	return
}

func NewGenerateDBAdminURLResponse() (response *GenerateDBAdminURLResponse) {
	response = &GenerateDBAdminURLResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GenerateDBAdminURL(request *GenerateDBAdminURLRequest) string {
	return c.GenerateDBAdminURLWithContext(context.Background(), request)
}

func (c *Client) GenerateDBAdminURLSend(request *GenerateDBAdminURLRequest) (*GenerateDBAdminURLResponse, error) {
	statusCode, msg, err := c.GenerateDBAdminURLWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GenerateDBAdminURLResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GenerateDBAdminURLWithContext(ctx context.Context, request *GenerateDBAdminURLRequest) string {
	if request == nil {
		request = NewGenerateDBAdminURLRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GenerateDBAdminURL")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGenerateDBAdminURLResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GenerateDBAdminURLWithContextV2(ctx context.Context, request *GenerateDBAdminURLRequest) (int, string, error) {
	if request == nil {
		request = NewGenerateDBAdminURLRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GenerateDBAdminURL")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGenerateDBAdminURLResponse()
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
	request.Init().WithApiInfo("krds", APIVersion, "AllocateDBInstanceEip")
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
		request.Init().WithApiInfo("krds", APIVersion, "AllocateDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

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
		request.Init().WithApiInfo("krds", APIVersion, "AllocateDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateDBInstanceEipResponse()
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
	request.Init().WithApiInfo("krds", APIVersion, "ReleaseDBInstanceEip")
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
		request.Init().WithApiInfo("krds", APIVersion, "ReleaseDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

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
		request.Init().WithApiInfo("krds", APIVersion, "ReleaseDBInstanceEip")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReleaseDBInstanceEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceAvailabilityZoneRequest() (request *ModifyDBInstanceAvailabilityZoneRequest) {
	request = &ModifyDBInstanceAvailabilityZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceAvailabilityZone")
	return
}

func NewModifyDBInstanceAvailabilityZoneResponse() (response *ModifyDBInstanceAvailabilityZoneResponse) {
	response = &ModifyDBInstanceAvailabilityZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstanceAvailabilityZone(request *ModifyDBInstanceAvailabilityZoneRequest) string {
	return c.ModifyDBInstanceAvailabilityZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceAvailabilityZoneSend(request *ModifyDBInstanceAvailabilityZoneRequest) (*ModifyDBInstanceAvailabilityZoneResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceAvailabilityZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceAvailabilityZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceAvailabilityZoneWithContext(ctx context.Context, request *ModifyDBInstanceAvailabilityZoneRequest) string {
	if request == nil {
		request = NewModifyDBInstanceAvailabilityZoneRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceAvailabilityZone")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceAvailabilityZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceAvailabilityZoneWithContextV2(ctx context.Context, request *ModifyDBInstanceAvailabilityZoneRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceAvailabilityZoneRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceAvailabilityZone")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceAvailabilityZoneResponse()
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
	request.Init().WithApiInfo("krds", APIVersion, "CreateSecurityGroup")
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
		request.Init().WithApiInfo("krds", APIVersion, "CreateSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

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
		request.Init().WithApiInfo("krds", APIVersion, "CreateSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupResponse()
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
	request.Init().WithApiInfo("krds", APIVersion, "DescribeSecurityGroup")
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
		request.Init().WithApiInfo("krds", APIVersion, "DescribeSecurityGroup")
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
		request.Init().WithApiInfo("krds", APIVersion, "DescribeSecurityGroup")
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
	request.Init().WithApiInfo("krds", APIVersion, "DeleteSecurityGroup")
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
		request.Init().WithApiInfo("krds", APIVersion, "DeleteSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

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
		request.Init().WithApiInfo("krds", APIVersion, "DeleteSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
	request = &ModifySecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroup")
	return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
	response = &ModifySecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) string {
	return c.ModifySecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupSend(request *ModifySecurityGroupRequest) (*ModifySecurityGroupResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupWithContext(ctx context.Context, request *ModifySecurityGroupRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupWithContextV2(ctx context.Context, request *ModifySecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
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
	request.Init().WithApiInfo("krds", APIVersion, "CloneSecurityGroup")
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
		request.Init().WithApiInfo("krds", APIVersion, "CloneSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

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
		request.Init().WithApiInfo("krds", APIVersion, "CloneSecurityGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRuleRequest() (request *ModifySecurityGroupRuleRequest) {
	request = &ModifySecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRule")
	return
}

func NewModifySecurityGroupRuleResponse() (response *ModifySecurityGroupRuleResponse) {
	response = &ModifySecurityGroupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroupRule(request *ModifySecurityGroupRuleRequest) string {
	return c.ModifySecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupRuleSend(request *ModifySecurityGroupRuleRequest) (*ModifySecurityGroupRuleResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySecurityGroupRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupRuleWithContext(ctx context.Context, request *ModifySecurityGroupRuleRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupRuleWithContextV2(ctx context.Context, request *ModifySecurityGroupRuleRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSecurityGroupRelationRequest() (request *SecurityGroupRelationRequest) {
	request = &SecurityGroupRelationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SecurityGroupRelation")
	return
}

func NewSecurityGroupRelationResponse() (response *SecurityGroupRelationResponse) {
	response = &SecurityGroupRelationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SecurityGroupRelation(request *SecurityGroupRelationRequest) string {
	return c.SecurityGroupRelationWithContext(context.Background(), request)
}

func (c *Client) SecurityGroupRelationSend(request *SecurityGroupRelationRequest) (*SecurityGroupRelationResponse, error) {
	statusCode, msg, err := c.SecurityGroupRelationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SecurityGroupRelationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SecurityGroupRelationWithContext(ctx context.Context, request *SecurityGroupRelationRequest) string {
	if request == nil {
		request = NewSecurityGroupRelationRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SecurityGroupRelation")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSecurityGroupRelationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SecurityGroupRelationWithContextV2(ctx context.Context, request *SecurityGroupRelationRequest) (int, string, error) {
	if request == nil {
		request = NewSecurityGroupRelationRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SecurityGroupRelation")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSecurityGroupRelationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRuleNameRequest() (request *ModifySecurityGroupRuleNameRequest) {
	request = &ModifySecurityGroupRuleNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRuleName")
	return
}

func NewModifySecurityGroupRuleNameResponse() (response *ModifySecurityGroupRuleNameResponse) {
	response = &ModifySecurityGroupRuleNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroupRuleName(request *ModifySecurityGroupRuleNameRequest) string {
	return c.ModifySecurityGroupRuleNameWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupRuleNameSend(request *ModifySecurityGroupRuleNameRequest) (*ModifySecurityGroupRuleNameResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupRuleNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySecurityGroupRuleNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupRuleNameWithContext(ctx context.Context, request *ModifySecurityGroupRuleNameRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRuleNameRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRuleName")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupRuleNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupRuleNameWithContextV2(ctx context.Context, request *ModifySecurityGroupRuleNameRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRuleNameRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRuleName")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupRuleNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeLastLogRequest() (request *DescribeLastLogRequest) {
	request = &DescribeLastLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeLastLog")
	return
}

func NewDescribeLastLogResponse() (response *DescribeLastLogResponse) {
	response = &DescribeLastLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLastLog(request *DescribeLastLogRequest) string {
	return c.DescribeLastLogWithContext(context.Background(), request)
}

func (c *Client) DescribeLastLogSend(request *DescribeLastLogRequest) (*DescribeLastLogResponse, error) {
	statusCode, msg, err := c.DescribeLastLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeLastLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeLastLogWithContext(ctx context.Context, request *DescribeLastLogRequest) string {
	if request == nil {
		request = NewDescribeLastLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeLastLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLastLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeLastLogWithContextV2(ctx context.Context, request *DescribeLastLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeLastLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeLastLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLastLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartAuditRequest() (request *StartAuditRequest) {
	request = &StartAuditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "StartAudit")
	return
}

func NewStartAuditResponse() (response *StartAuditResponse) {
	response = &StartAuditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartAudit(request *StartAuditRequest) string {
	return c.StartAuditWithContext(context.Background(), request)
}

func (c *Client) StartAuditSend(request *StartAuditRequest) (*StartAuditResponse, error) {
	statusCode, msg, err := c.StartAuditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartAuditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartAuditWithContext(ctx context.Context, request *StartAuditRequest) string {
	if request == nil {
		request = NewStartAuditRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StartAudit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartAuditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartAuditWithContextV2(ctx context.Context, request *StartAuditRequest) (int, string, error) {
	if request == nil {
		request = NewStartAuditRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StartAudit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartAuditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopAuditRequest() (request *StopAuditRequest) {
	request = &StopAuditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "StopAudit")
	return
}

func NewStopAuditResponse() (response *StopAuditResponse) {
	response = &StopAuditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopAudit(request *StopAuditRequest) string {
	return c.StopAuditWithContext(context.Background(), request)
}

func (c *Client) StopAuditSend(request *StopAuditRequest) (*StopAuditResponse, error) {
	statusCode, msg, err := c.StopAuditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopAuditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopAuditWithContext(ctx context.Context, request *StopAuditRequest) string {
	if request == nil {
		request = NewStopAuditRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StopAudit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopAuditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopAuditWithContextV2(ctx context.Context, request *StopAuditRequest) (int, string, error) {
	if request == nil {
		request = NewStopAuditRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StopAudit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopAuditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListAuditRequest() (request *ListAuditRequest) {
	request = &ListAuditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ListAudit")
	return
}

func NewListAuditResponse() (response *ListAuditResponse) {
	response = &ListAuditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListAudit(request *ListAuditRequest) string {
	return c.ListAuditWithContext(context.Background(), request)
}

func (c *Client) ListAuditSend(request *ListAuditRequest) (*ListAuditResponse, error) {
	statusCode, msg, err := c.ListAuditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListAuditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListAuditWithContext(ctx context.Context, request *ListAuditRequest) string {
	if request == nil {
		request = NewListAuditRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ListAudit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAuditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListAuditWithContextV2(ctx context.Context, request *ListAuditRequest) (int, string, error) {
	if request == nil {
		request = NewListAuditRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ListAudit")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAuditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAuditStatisticRequest() (request *AuditStatisticRequest) {
	request = &AuditStatisticRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "AuditStatistic")
	return
}

func NewAuditStatisticResponse() (response *AuditStatisticResponse) {
	response = &AuditStatisticResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AuditStatistic(request *AuditStatisticRequest) string {
	return c.AuditStatisticWithContext(context.Background(), request)
}

func (c *Client) AuditStatisticSend(request *AuditStatisticRequest) (*AuditStatisticResponse, error) {
	statusCode, msg, err := c.AuditStatisticWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AuditStatisticResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AuditStatisticWithContext(ctx context.Context, request *AuditStatisticRequest) string {
	if request == nil {
		request = NewAuditStatisticRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "AuditStatistic")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAuditStatisticResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AuditStatisticWithContextV2(ctx context.Context, request *AuditStatisticRequest) (int, string, error) {
	if request == nil {
		request = NewAuditStatisticRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "AuditStatistic")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAuditStatisticResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetCurrentDatabaseInfoRequest() (request *GetCurrentDatabaseInfoRequest) {
	request = &GetCurrentDatabaseInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "GetCurrentDatabaseInfo")
	return
}

func NewGetCurrentDatabaseInfoResponse() (response *GetCurrentDatabaseInfoResponse) {
	response = &GetCurrentDatabaseInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetCurrentDatabaseInfo(request *GetCurrentDatabaseInfoRequest) string {
	return c.GetCurrentDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) GetCurrentDatabaseInfoSend(request *GetCurrentDatabaseInfoRequest) (*GetCurrentDatabaseInfoResponse, error) {
	statusCode, msg, err := c.GetCurrentDatabaseInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetCurrentDatabaseInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetCurrentDatabaseInfoWithContext(ctx context.Context, request *GetCurrentDatabaseInfoRequest) string {
	if request == nil {
		request = NewGetCurrentDatabaseInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GetCurrentDatabaseInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCurrentDatabaseInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetCurrentDatabaseInfoWithContextV2(ctx context.Context, request *GetCurrentDatabaseInfoRequest) (int, string, error) {
	if request == nil {
		request = NewGetCurrentDatabaseInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GetCurrentDatabaseInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCurrentDatabaseInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetTableRestorableTimeRequest() (request *GetTableRestorableTimeRequest) {
	request = &GetTableRestorableTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "GetTableRestorableTime")
	return
}

func NewGetTableRestorableTimeResponse() (response *GetTableRestorableTimeResponse) {
	response = &GetTableRestorableTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetTableRestorableTime(request *GetTableRestorableTimeRequest) string {
	return c.GetTableRestorableTimeWithContext(context.Background(), request)
}

func (c *Client) GetTableRestorableTimeSend(request *GetTableRestorableTimeRequest) (*GetTableRestorableTimeResponse, error) {
	statusCode, msg, err := c.GetTableRestorableTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetTableRestorableTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetTableRestorableTimeWithContext(ctx context.Context, request *GetTableRestorableTimeRequest) string {
	if request == nil {
		request = NewGetTableRestorableTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GetTableRestorableTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTableRestorableTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetTableRestorableTimeWithContextV2(ctx context.Context, request *GetTableRestorableTimeRequest) (int, string, error) {
	if request == nil {
		request = NewGetTableRestorableTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GetTableRestorableTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTableRestorableTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetHistoryDatabaseInfoRequest() (request *GetHistoryDatabaseInfoRequest) {
	request = &GetHistoryDatabaseInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "GetHistoryDatabaseInfo")
	return
}

func NewGetHistoryDatabaseInfoResponse() (response *GetHistoryDatabaseInfoResponse) {
	response = &GetHistoryDatabaseInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetHistoryDatabaseInfo(request *GetHistoryDatabaseInfoRequest) string {
	return c.GetHistoryDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) GetHistoryDatabaseInfoSend(request *GetHistoryDatabaseInfoRequest) (*GetHistoryDatabaseInfoResponse, error) {
	statusCode, msg, err := c.GetHistoryDatabaseInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetHistoryDatabaseInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetHistoryDatabaseInfoWithContext(ctx context.Context, request *GetHistoryDatabaseInfoRequest) string {
	if request == nil {
		request = NewGetHistoryDatabaseInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GetHistoryDatabaseInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetHistoryDatabaseInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetHistoryDatabaseInfoWithContextV2(ctx context.Context, request *GetHistoryDatabaseInfoRequest) (int, string, error) {
	if request == nil {
		request = NewGetHistoryDatabaseInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "GetHistoryDatabaseInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetHistoryDatabaseInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOverrideDBInstanceByPointInTimeRequest() (request *OverrideDBInstanceByPointInTimeRequest) {
	request = &OverrideDBInstanceByPointInTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstanceByPointInTime")
	return
}

func NewOverrideDBInstanceByPointInTimeResponse() (response *OverrideDBInstanceByPointInTimeResponse) {
	response = &OverrideDBInstanceByPointInTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OverrideDBInstanceByPointInTime(request *OverrideDBInstanceByPointInTimeRequest) string {
	return c.OverrideDBInstanceByPointInTimeWithContext(context.Background(), request)
}

func (c *Client) OverrideDBInstanceByPointInTimeSend(request *OverrideDBInstanceByPointInTimeRequest) (*OverrideDBInstanceByPointInTimeResponse, error) {
	statusCode, msg, err := c.OverrideDBInstanceByPointInTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct OverrideDBInstanceByPointInTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OverrideDBInstanceByPointInTimeWithContext(ctx context.Context, request *OverrideDBInstanceByPointInTimeRequest) string {
	if request == nil {
		request = NewOverrideDBInstanceByPointInTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstanceByPointInTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOverrideDBInstanceByPointInTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OverrideDBInstanceByPointInTimeWithContextV2(ctx context.Context, request *OverrideDBInstanceByPointInTimeRequest) (int, string, error) {
	if request == nil {
		request = NewOverrideDBInstanceByPointInTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstanceByPointInTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOverrideDBInstanceByPointInTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreToCurInstanceRequest() (request *RestoreToCurInstanceRequest) {
	request = &RestoreToCurInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "RestoreToCurInstance")
	return
}

func NewRestoreToCurInstanceResponse() (response *RestoreToCurInstanceResponse) {
	response = &RestoreToCurInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreToCurInstance(request *RestoreToCurInstanceRequest) string {
	return c.RestoreToCurInstanceWithContext(context.Background(), request)
}

func (c *Client) RestoreToCurInstanceSend(request *RestoreToCurInstanceRequest) (*RestoreToCurInstanceResponse, error) {
	statusCode, msg, err := c.RestoreToCurInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreToCurInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreToCurInstanceWithContext(ctx context.Context, request *RestoreToCurInstanceRequest) string {
	if request == nil {
		request = NewRestoreToCurInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreToCurInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreToCurInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreToCurInstanceWithContextV2(ctx context.Context, request *RestoreToCurInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreToCurInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreToCurInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreToCurInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreToSgInstanceRequest() (request *RestoreToSgInstanceRequest) {
	request = &RestoreToSgInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "RestoreToSgInstance")
	return
}

func NewRestoreToSgInstanceResponse() (response *RestoreToSgInstanceResponse) {
	response = &RestoreToSgInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreToSgInstance(request *RestoreToSgInstanceRequest) string {
	return c.RestoreToSgInstanceWithContext(context.Background(), request)
}

func (c *Client) RestoreToSgInstanceSend(request *RestoreToSgInstanceRequest) (*RestoreToSgInstanceResponse, error) {
	statusCode, msg, err := c.RestoreToSgInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreToSgInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreToSgInstanceWithContext(ctx context.Context, request *RestoreToSgInstanceRequest) string {
	if request == nil {
		request = NewRestoreToSgInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreToSgInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreToSgInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreToSgInstanceWithContextV2(ctx context.Context, request *RestoreToSgInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreToSgInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "RestoreToSgInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreToSgInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAuditHotCountRequest() (request *DescribeAuditHotCountRequest) {
	request = &DescribeAuditHotCountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotCount")
	return
}

func NewDescribeAuditHotCountResponse() (response *DescribeAuditHotCountResponse) {
	response = &DescribeAuditHotCountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAuditHotCount(request *DescribeAuditHotCountRequest) string {
	return c.DescribeAuditHotCountWithContext(context.Background(), request)
}

func (c *Client) DescribeAuditHotCountSend(request *DescribeAuditHotCountRequest) (*DescribeAuditHotCountResponse, error) {
	statusCode, msg, err := c.DescribeAuditHotCountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAuditHotCountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAuditHotCountWithContext(ctx context.Context, request *DescribeAuditHotCountRequest) string {
	if request == nil {
		request = NewDescribeAuditHotCountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotCount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAuditHotCountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAuditHotCountWithContextV2(ctx context.Context, request *DescribeAuditHotCountRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAuditHotCountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotCount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAuditHotCountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAuditHotDurationRequest() (request *DescribeAuditHotDurationRequest) {
	request = &DescribeAuditHotDurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotDuration")
	return
}

func NewDescribeAuditHotDurationResponse() (response *DescribeAuditHotDurationResponse) {
	response = &DescribeAuditHotDurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAuditHotDuration(request *DescribeAuditHotDurationRequest) string {
	return c.DescribeAuditHotDurationWithContext(context.Background(), request)
}

func (c *Client) DescribeAuditHotDurationSend(request *DescribeAuditHotDurationRequest) (*DescribeAuditHotDurationResponse, error) {
	statusCode, msg, err := c.DescribeAuditHotDurationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAuditHotDurationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAuditHotDurationWithContext(ctx context.Context, request *DescribeAuditHotDurationRequest) string {
	if request == nil {
		request = NewDescribeAuditHotDurationRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotDuration")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAuditHotDurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAuditHotDurationWithContextV2(ctx context.Context, request *DescribeAuditHotDurationRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAuditHotDurationRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotDuration")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAuditHotDurationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSqlAuditReportRequest() (request *SqlAuditReportRequest) {
	request = &SqlAuditReportRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SqlAuditReport")
	return
}

func NewSqlAuditReportResponse() (response *SqlAuditReportResponse) {
	response = &SqlAuditReportResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SqlAuditReport(request *SqlAuditReportRequest) string {
	return c.SqlAuditReportWithContext(context.Background(), request)
}

func (c *Client) SqlAuditReportSend(request *SqlAuditReportRequest) (*SqlAuditReportResponse, error) {
	statusCode, msg, err := c.SqlAuditReportWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SqlAuditReportResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SqlAuditReportWithContext(ctx context.Context, request *SqlAuditReportRequest) string {
	if request == nil {
		request = NewSqlAuditReportRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SqlAuditReport")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSqlAuditReportResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SqlAuditReportWithContextV2(ctx context.Context, request *SqlAuditReportRequest) (int, string, error) {
	if request == nil {
		request = NewSqlAuditReportRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SqlAuditReport")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSqlAuditReportResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSqlAuditLineChartRequest() (request *SqlAuditLineChartRequest) {
	request = &SqlAuditLineChartRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SqlAuditLineChart")
	return
}

func NewSqlAuditLineChartResponse() (response *SqlAuditLineChartResponse) {
	response = &SqlAuditLineChartResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SqlAuditLineChart(request *SqlAuditLineChartRequest) string {
	return c.SqlAuditLineChartWithContext(context.Background(), request)
}

func (c *Client) SqlAuditLineChartSend(request *SqlAuditLineChartRequest) (*SqlAuditLineChartResponse, error) {
	statusCode, msg, err := c.SqlAuditLineChartWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SqlAuditLineChartResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SqlAuditLineChartWithContext(ctx context.Context, request *SqlAuditLineChartRequest) string {
	if request == nil {
		request = NewSqlAuditLineChartRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SqlAuditLineChart")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSqlAuditLineChartResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SqlAuditLineChartWithContextV2(ctx context.Context, request *SqlAuditLineChartRequest) (int, string, error) {
	if request == nil {
		request = NewSqlAuditLineChartRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SqlAuditLineChart")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSqlAuditLineChartResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSlowLogReportRequest() (request *SlowLogReportRequest) {
	request = &SlowLogReportRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SlowLogReport")
	return
}

func NewSlowLogReportResponse() (response *SlowLogReportResponse) {
	response = &SlowLogReportResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SlowLogReport(request *SlowLogReportRequest) string {
	return c.SlowLogReportWithContext(context.Background(), request)
}

func (c *Client) SlowLogReportSend(request *SlowLogReportRequest) (*SlowLogReportResponse, error) {
	statusCode, msg, err := c.SlowLogReportWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SlowLogReportResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SlowLogReportWithContext(ctx context.Context, request *SlowLogReportRequest) string {
	if request == nil {
		request = NewSlowLogReportRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SlowLogReport")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSlowLogReportResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SlowLogReportWithContextV2(ctx context.Context, request *SlowLogReportRequest) (int, string, error) {
	if request == nil {
		request = NewSlowLogReportRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SlowLogReport")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSlowLogReportResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSlowLogLineChartRequest() (request *SlowLogLineChartRequest) {
	request = &SlowLogLineChartRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SlowLogLineChart")
	return
}

func NewSlowLogLineChartResponse() (response *SlowLogLineChartResponse) {
	response = &SlowLogLineChartResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SlowLogLineChart(request *SlowLogLineChartRequest) string {
	return c.SlowLogLineChartWithContext(context.Background(), request)
}

func (c *Client) SlowLogLineChartSend(request *SlowLogLineChartRequest) (*SlowLogLineChartResponse, error) {
	statusCode, msg, err := c.SlowLogLineChartWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SlowLogLineChartResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SlowLogLineChartWithContext(ctx context.Context, request *SlowLogLineChartRequest) string {
	if request == nil {
		request = NewSlowLogLineChartRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SlowLogLineChart")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSlowLogLineChartResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SlowLogLineChartWithContextV2(ctx context.Context, request *SlowLogLineChartRequest) (int, string, error) {
	if request == nil {
		request = NewSlowLogLineChartRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SlowLogLineChart")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSlowLogLineChartResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSlowLogDetailRequest() (request *SlowLogDetailRequest) {
	request = &SlowLogDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SlowLogDetail")
	return
}

func NewSlowLogDetailResponse() (response *SlowLogDetailResponse) {
	response = &SlowLogDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SlowLogDetail(request *SlowLogDetailRequest) string {
	return c.SlowLogDetailWithContext(context.Background(), request)
}

func (c *Client) SlowLogDetailSend(request *SlowLogDetailRequest) (*SlowLogDetailResponse, error) {
	statusCode, msg, err := c.SlowLogDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SlowLogDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SlowLogDetailWithContext(ctx context.Context, request *SlowLogDetailRequest) string {
	if request == nil {
		request = NewSlowLogDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SlowLogDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSlowLogDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SlowLogDetailWithContextV2(ctx context.Context, request *SlowLogDetailRequest) (int, string, error) {
	if request == nil {
		request = NewSlowLogDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SlowLogDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSlowLogDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartAuditDetailExportTaskRequest() (request *StartAuditDetailExportTaskRequest) {
	request = &StartAuditDetailExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "StartAuditDetailExportTask")
	return
}

func NewStartAuditDetailExportTaskResponse() (response *StartAuditDetailExportTaskResponse) {
	response = &StartAuditDetailExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartAuditDetailExportTask(request *StartAuditDetailExportTaskRequest) string {
	return c.StartAuditDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) StartAuditDetailExportTaskSend(request *StartAuditDetailExportTaskRequest) (*StartAuditDetailExportTaskResponse, error) {
	statusCode, msg, err := c.StartAuditDetailExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartAuditDetailExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartAuditDetailExportTaskWithContext(ctx context.Context, request *StartAuditDetailExportTaskRequest) string {
	if request == nil {
		request = NewStartAuditDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StartAuditDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartAuditDetailExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartAuditDetailExportTaskWithContextV2(ctx context.Context, request *StartAuditDetailExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewStartAuditDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StartAuditDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartAuditDetailExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListAuditDetailExportTaskRequest() (request *ListAuditDetailExportTaskRequest) {
	request = &ListAuditDetailExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ListAuditDetailExportTask")
	return
}

func NewListAuditDetailExportTaskResponse() (response *ListAuditDetailExportTaskResponse) {
	response = &ListAuditDetailExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListAuditDetailExportTask(request *ListAuditDetailExportTaskRequest) string {
	return c.ListAuditDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) ListAuditDetailExportTaskSend(request *ListAuditDetailExportTaskRequest) (*ListAuditDetailExportTaskResponse, error) {
	statusCode, msg, err := c.ListAuditDetailExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListAuditDetailExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListAuditDetailExportTaskWithContext(ctx context.Context, request *ListAuditDetailExportTaskRequest) string {
	if request == nil {
		request = NewListAuditDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ListAuditDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAuditDetailExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListAuditDetailExportTaskWithContextV2(ctx context.Context, request *ListAuditDetailExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewListAuditDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ListAuditDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAuditDetailExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInstanceAccountRequest() (request *CreateInstanceAccountRequest) {
	request = &CreateInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccount")
	return
}

func NewCreateInstanceAccountResponse() (response *CreateInstanceAccountResponse) {
	response = &CreateInstanceAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstanceAccount(request *CreateInstanceAccountRequest) string {
	return c.CreateInstanceAccountWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceAccountSend(request *CreateInstanceAccountRequest) (*CreateInstanceAccountResponse, error) {
	statusCode, msg, err := c.CreateInstanceAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInstanceAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInstanceAccountWithContext(ctx context.Context, request *CreateInstanceAccountRequest) string {
	if request == nil {
		request = NewCreateInstanceAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInstanceAccountWithContextV2(ctx context.Context, request *CreateInstanceAccountRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceAccountsRequest() (request *DescribeInstanceAccountsRequest) {
	request = &DescribeInstanceAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceAccounts")
	return
}

func NewDescribeInstanceAccountsResponse() (response *DescribeInstanceAccountsResponse) {
	response = &DescribeInstanceAccountsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceAccounts(request *DescribeInstanceAccountsRequest) string {
	return c.DescribeInstanceAccountsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceAccountsSend(request *DescribeInstanceAccountsRequest) (*DescribeInstanceAccountsResponse, error) {
	statusCode, msg, err := c.DescribeInstanceAccountsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstanceAccountsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceAccountsWithContext(ctx context.Context, request *DescribeInstanceAccountsRequest) string {
	if request == nil {
		request = NewDescribeInstanceAccountsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceAccounts")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceAccountsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceAccountsWithContextV2(ctx context.Context, request *DescribeInstanceAccountsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceAccountsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceAccounts")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceAccountsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAccountInfoRequest() (request *ModifyInstanceAccountInfoRequest) {
	request = &ModifyInstanceAccountInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountInfo")
	return
}

func NewModifyInstanceAccountInfoResponse() (response *ModifyInstanceAccountInfoResponse) {
	response = &ModifyInstanceAccountInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAccountInfo(request *ModifyInstanceAccountInfoRequest) string {
	return c.ModifyInstanceAccountInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountInfoSend(request *ModifyInstanceAccountInfoRequest) (*ModifyInstanceAccountInfoResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAccountInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceAccountInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceAccountInfoWithContext(ctx context.Context, request *ModifyInstanceAccountInfoRequest) string {
	if request == nil {
		request = NewModifyInstanceAccountInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAccountInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceAccountInfoWithContextV2(ctx context.Context, request *ModifyInstanceAccountInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAccountInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAccountInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAccountPrivilegesRequest() (request *ModifyInstanceAccountPrivilegesRequest) {
	request = &ModifyInstanceAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivileges")
	return
}

func NewModifyInstanceAccountPrivilegesResponse() (response *ModifyInstanceAccountPrivilegesResponse) {
	response = &ModifyInstanceAccountPrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAccountPrivileges(request *ModifyInstanceAccountPrivilegesRequest) string {
	return c.ModifyInstanceAccountPrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountPrivilegesSend(request *ModifyInstanceAccountPrivilegesRequest) (*ModifyInstanceAccountPrivilegesResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAccountPrivilegesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceAccountPrivilegesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceAccountPrivilegesWithContext(ctx context.Context, request *ModifyInstanceAccountPrivilegesRequest) string {
	if request == nil {
		request = NewModifyInstanceAccountPrivilegesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivileges")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAccountPrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceAccountPrivilegesWithContextV2(ctx context.Context, request *ModifyInstanceAccountPrivilegesRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAccountPrivilegesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivileges")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAccountPrivilegesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceAccountRequest() (request *DeleteInstanceAccountRequest) {
	request = &DeleteInstanceAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccount")
	return
}

func NewDeleteInstanceAccountResponse() (response *DeleteInstanceAccountResponse) {
	response = &DeleteInstanceAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstanceAccount(request *DeleteInstanceAccountRequest) string {
	return c.DeleteInstanceAccountWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceAccountSend(request *DeleteInstanceAccountRequest) (*DeleteInstanceAccountResponse, error) {
	statusCode, msg, err := c.DeleteInstanceAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInstanceAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceAccountWithContext(ctx context.Context, request *DeleteInstanceAccountRequest) string {
	if request == nil {
		request = NewDeleteInstanceAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceAccountWithContextV2(ctx context.Context, request *DeleteInstanceAccountRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceAccountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCollationsRequest() (request *DescribeCollationsRequest) {
	request = &DescribeCollationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeCollations")
	return
}

func NewDescribeCollationsResponse() (response *DescribeCollationsResponse) {
	response = &DescribeCollationsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCollations(request *DescribeCollationsRequest) string {
	return c.DescribeCollationsWithContext(context.Background(), request)
}

func (c *Client) DescribeCollationsSend(request *DescribeCollationsRequest) (*DescribeCollationsResponse, error) {
	statusCode, msg, err := c.DescribeCollationsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCollationsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCollationsWithContext(ctx context.Context, request *DescribeCollationsRequest) string {
	if request == nil {
		request = NewDescribeCollationsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeCollations")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCollationsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCollationsWithContextV2(ctx context.Context, request *DescribeCollationsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCollationsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeCollations")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCollationsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInstanceDatabaseRequest() (request *CreateInstanceDatabaseRequest) {
	request = &CreateInstanceDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceDatabase")
	return
}

func NewCreateInstanceDatabaseResponse() (response *CreateInstanceDatabaseResponse) {
	response = &CreateInstanceDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstanceDatabase(request *CreateInstanceDatabaseRequest) string {
	return c.CreateInstanceDatabaseWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceDatabaseSend(request *CreateInstanceDatabaseRequest) (*CreateInstanceDatabaseResponse, error) {
	statusCode, msg, err := c.CreateInstanceDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInstanceDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInstanceDatabaseWithContext(ctx context.Context, request *CreateInstanceDatabaseRequest) string {
	if request == nil {
		request = NewCreateInstanceDatabaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceDatabase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInstanceDatabaseWithContextV2(ctx context.Context, request *CreateInstanceDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceDatabaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceDatabase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceDatabasePrivilegesRequest() (request *ModifyInstanceDatabasePrivilegesRequest) {
	request = &ModifyInstanceDatabasePrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivileges")
	return
}

func NewModifyInstanceDatabasePrivilegesResponse() (response *ModifyInstanceDatabasePrivilegesResponse) {
	response = &ModifyInstanceDatabasePrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabasePrivileges(request *ModifyInstanceDatabasePrivilegesRequest) string {
	return c.ModifyInstanceDatabasePrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabasePrivilegesSend(request *ModifyInstanceDatabasePrivilegesRequest) (*ModifyInstanceDatabasePrivilegesResponse, error) {
	statusCode, msg, err := c.ModifyInstanceDatabasePrivilegesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceDatabasePrivilegesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceDatabasePrivilegesWithContext(ctx context.Context, request *ModifyInstanceDatabasePrivilegesRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabasePrivilegesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivileges")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceDatabasePrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceDatabasePrivilegesWithContextV2(ctx context.Context, request *ModifyInstanceDatabasePrivilegesRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceDatabasePrivilegesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivileges")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceDatabasePrivilegesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceDatabasesRequest() (request *DescribeInstanceDatabasesRequest) {
	request = &DescribeInstanceDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceDatabases")
	return
}

func NewDescribeInstanceDatabasesResponse() (response *DescribeInstanceDatabasesResponse) {
	response = &DescribeInstanceDatabasesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceDatabases(request *DescribeInstanceDatabasesRequest) string {
	return c.DescribeInstanceDatabasesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceDatabasesSend(request *DescribeInstanceDatabasesRequest) (*DescribeInstanceDatabasesResponse, error) {
	statusCode, msg, err := c.DescribeInstanceDatabasesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstanceDatabasesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceDatabasesWithContext(ctx context.Context, request *DescribeInstanceDatabasesRequest) string {
	if request == nil {
		request = NewDescribeInstanceDatabasesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceDatabases")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceDatabasesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceDatabasesWithContextV2(ctx context.Context, request *DescribeInstanceDatabasesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceDatabasesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceDatabases")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceDatabasesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceDatabaseInfoRequest() (request *ModifyInstanceDatabaseInfoRequest) {
	request = &ModifyInstanceDatabaseInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabaseInfo")
	return
}

func NewModifyInstanceDatabaseInfoResponse() (response *ModifyInstanceDatabaseInfoResponse) {
	response = &ModifyInstanceDatabaseInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabaseInfo(request *ModifyInstanceDatabaseInfoRequest) string {
	return c.ModifyInstanceDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabaseInfoSend(request *ModifyInstanceDatabaseInfoRequest) (*ModifyInstanceDatabaseInfoResponse, error) {
	statusCode, msg, err := c.ModifyInstanceDatabaseInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceDatabaseInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceDatabaseInfoWithContext(ctx context.Context, request *ModifyInstanceDatabaseInfoRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabaseInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabaseInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceDatabaseInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceDatabaseInfoWithContextV2(ctx context.Context, request *ModifyInstanceDatabaseInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceDatabaseInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabaseInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceDatabaseInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartSlowLogDetailExportTaskRequest() (request *StartSlowLogDetailExportTaskRequest) {
	request = &StartSlowLogDetailExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "StartSlowLogDetailExportTask")
	return
}

func NewStartSlowLogDetailExportTaskResponse() (response *StartSlowLogDetailExportTaskResponse) {
	response = &StartSlowLogDetailExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartSlowLogDetailExportTask(request *StartSlowLogDetailExportTaskRequest) string {
	return c.StartSlowLogDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) StartSlowLogDetailExportTaskSend(request *StartSlowLogDetailExportTaskRequest) (*StartSlowLogDetailExportTaskResponse, error) {
	statusCode, msg, err := c.StartSlowLogDetailExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartSlowLogDetailExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartSlowLogDetailExportTaskWithContext(ctx context.Context, request *StartSlowLogDetailExportTaskRequest) string {
	if request == nil {
		request = NewStartSlowLogDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StartSlowLogDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartSlowLogDetailExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartSlowLogDetailExportTaskWithContextV2(ctx context.Context, request *StartSlowLogDetailExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewStartSlowLogDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "StartSlowLogDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartSlowLogDetailExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListSlowLogDetailExportTaskRequest() (request *ListSlowLogDetailExportTaskRequest) {
	request = &ListSlowLogDetailExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ListSlowLogDetailExportTask")
	return
}

func NewListSlowLogDetailExportTaskResponse() (response *ListSlowLogDetailExportTaskResponse) {
	response = &ListSlowLogDetailExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSlowLogDetailExportTask(request *ListSlowLogDetailExportTaskRequest) string {
	return c.ListSlowLogDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) ListSlowLogDetailExportTaskSend(request *ListSlowLogDetailExportTaskRequest) (*ListSlowLogDetailExportTaskResponse, error) {
	statusCode, msg, err := c.ListSlowLogDetailExportTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListSlowLogDetailExportTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListSlowLogDetailExportTaskWithContext(ctx context.Context, request *ListSlowLogDetailExportTaskRequest) string {
	if request == nil {
		request = NewListSlowLogDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ListSlowLogDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListSlowLogDetailExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListSlowLogDetailExportTaskWithContextV2(ctx context.Context, request *ListSlowLogDetailExportTaskRequest) (int, string, error) {
	if request == nil {
		request = NewListSlowLogDetailExportTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ListSlowLogDetailExportTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListSlowLogDetailExportTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInstanceAccountActionRequest() (request *CreateInstanceAccountActionRequest) {
	request = &CreateInstanceAccountActionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccountAction")
	return
}

func NewCreateInstanceAccountActionResponse() (response *CreateInstanceAccountActionResponse) {
	response = &CreateInstanceAccountActionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstanceAccountAction(request *CreateInstanceAccountActionRequest) string {
	return c.CreateInstanceAccountActionWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceAccountActionSend(request *CreateInstanceAccountActionRequest) (*CreateInstanceAccountActionResponse, error) {
	statusCode, msg, err := c.CreateInstanceAccountActionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateInstanceAccountActionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInstanceAccountActionWithContext(ctx context.Context, request *CreateInstanceAccountActionRequest) string {
	if request == nil {
		request = NewCreateInstanceAccountActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccountAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceAccountActionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInstanceAccountActionWithContextV2(ctx context.Context, request *CreateInstanceAccountActionRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceAccountActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccountAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateInstanceAccountActionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAccountPrivilegesActionRequest() (request *ModifyInstanceAccountPrivilegesActionRequest) {
	request = &ModifyInstanceAccountPrivilegesActionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivilegesAction")
	return
}

func NewModifyInstanceAccountPrivilegesActionResponse() (response *ModifyInstanceAccountPrivilegesActionResponse) {
	response = &ModifyInstanceAccountPrivilegesActionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAccountPrivilegesAction(request *ModifyInstanceAccountPrivilegesActionRequest) string {
	return c.ModifyInstanceAccountPrivilegesActionWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountPrivilegesActionSend(request *ModifyInstanceAccountPrivilegesActionRequest) (*ModifyInstanceAccountPrivilegesActionResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAccountPrivilegesActionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceAccountPrivilegesActionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceAccountPrivilegesActionWithContext(ctx context.Context, request *ModifyInstanceAccountPrivilegesActionRequest) string {
	if request == nil {
		request = NewModifyInstanceAccountPrivilegesActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivilegesAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceAccountPrivilegesActionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceAccountPrivilegesActionWithContextV2(ctx context.Context, request *ModifyInstanceAccountPrivilegesActionRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAccountPrivilegesActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivilegesAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceAccountPrivilegesActionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceAccountActionRequest() (request *DeleteInstanceAccountActionRequest) {
	request = &DeleteInstanceAccountActionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccountAction")
	return
}

func NewDeleteInstanceAccountActionResponse() (response *DeleteInstanceAccountActionResponse) {
	response = &DeleteInstanceAccountActionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstanceAccountAction(request *DeleteInstanceAccountActionRequest) string {
	return c.DeleteInstanceAccountActionWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceAccountActionSend(request *DeleteInstanceAccountActionRequest) (*DeleteInstanceAccountActionResponse, error) {
	statusCode, msg, err := c.DeleteInstanceAccountActionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInstanceAccountActionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceAccountActionWithContext(ctx context.Context, request *DeleteInstanceAccountActionRequest) string {
	if request == nil {
		request = NewDeleteInstanceAccountActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccountAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceAccountActionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceAccountActionWithContextV2(ctx context.Context, request *DeleteInstanceAccountActionRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceAccountActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccountAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceAccountActionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceDatabaseActionRequest() (request *DeleteInstanceDatabaseActionRequest) {
	request = &DeleteInstanceDatabaseActionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceDatabaseAction")
	return
}

func NewDeleteInstanceDatabaseActionResponse() (response *DeleteInstanceDatabaseActionResponse) {
	response = &DeleteInstanceDatabaseActionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstanceDatabaseAction(request *DeleteInstanceDatabaseActionRequest) string {
	return c.DeleteInstanceDatabaseActionWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceDatabaseActionSend(request *DeleteInstanceDatabaseActionRequest) (*DeleteInstanceDatabaseActionResponse, error) {
	statusCode, msg, err := c.DeleteInstanceDatabaseActionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteInstanceDatabaseActionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceDatabaseActionWithContext(ctx context.Context, request *DeleteInstanceDatabaseActionRequest) string {
	if request == nil {
		request = NewDeleteInstanceDatabaseActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceDatabaseAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceDatabaseActionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceDatabaseActionWithContextV2(ctx context.Context, request *DeleteInstanceDatabaseActionRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceDatabaseActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceDatabaseAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceDatabaseActionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBNetworkRequest() (request *ModifyDBNetworkRequest) {
	request = &ModifyDBNetworkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBNetwork")
	return
}

func NewModifyDBNetworkResponse() (response *ModifyDBNetworkResponse) {
	response = &ModifyDBNetworkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBNetwork(request *ModifyDBNetworkRequest) string {
	return c.ModifyDBNetworkWithContext(context.Background(), request)
}

func (c *Client) ModifyDBNetworkSend(request *ModifyDBNetworkRequest) (*ModifyDBNetworkResponse, error) {
	statusCode, msg, err := c.ModifyDBNetworkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBNetworkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBNetworkWithContext(ctx context.Context, request *ModifyDBNetworkRequest) string {
	if request == nil {
		request = NewModifyDBNetworkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBNetwork")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBNetworkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBNetworkWithContextV2(ctx context.Context, request *ModifyDBNetworkRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBNetworkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBNetwork")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyDBNetworkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDBInstanceMonitorPeriodRequest() (request *DescribeDBInstanceMonitorPeriodRequest) {
	request = &DescribeDBInstanceMonitorPeriodRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceMonitorPeriod")
	return
}

func NewDescribeDBInstanceMonitorPeriodResponse() (response *DescribeDBInstanceMonitorPeriodResponse) {
	response = &DescribeDBInstanceMonitorPeriodResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDBInstanceMonitorPeriod(request *DescribeDBInstanceMonitorPeriodRequest) string {
	return c.DescribeDBInstanceMonitorPeriodWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceMonitorPeriodSend(request *DescribeDBInstanceMonitorPeriodRequest) (*DescribeDBInstanceMonitorPeriodResponse, error) {
	statusCode, msg, err := c.DescribeDBInstanceMonitorPeriodWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDBInstanceMonitorPeriodResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDBInstanceMonitorPeriodWithContext(ctx context.Context, request *DescribeDBInstanceMonitorPeriodRequest) string {
	if request == nil {
		request = NewDescribeDBInstanceMonitorPeriodRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceMonitorPeriod")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstanceMonitorPeriodResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDBInstanceMonitorPeriodWithContextV2(ctx context.Context, request *DescribeDBInstanceMonitorPeriodRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDBInstanceMonitorPeriodRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceMonitorPeriod")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDBInstanceMonitorPeriodResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDBInstanceMonitorPeriodRequest() (request *ModifyDBInstanceMonitorPeriodRequest) {
	request = &ModifyDBInstanceMonitorPeriodRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceMonitorPeriod")
	return
}

func NewModifyDBInstanceMonitorPeriodResponse() (response *ModifyDBInstanceMonitorPeriodResponse) {
	response = &ModifyDBInstanceMonitorPeriodResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDBInstanceMonitorPeriod(request *ModifyDBInstanceMonitorPeriodRequest) string {
	return c.ModifyDBInstanceMonitorPeriodWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceMonitorPeriodSend(request *ModifyDBInstanceMonitorPeriodRequest) (*ModifyDBInstanceMonitorPeriodResponse, error) {
	statusCode, msg, err := c.ModifyDBInstanceMonitorPeriodWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyDBInstanceMonitorPeriodResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDBInstanceMonitorPeriodWithContext(ctx context.Context, request *ModifyDBInstanceMonitorPeriodRequest) string {
	if request == nil {
		request = NewModifyDBInstanceMonitorPeriodRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceMonitorPeriod")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceMonitorPeriodResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDBInstanceMonitorPeriodWithContextV2(ctx context.Context, request *ModifyDBInstanceMonitorPeriodRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDBInstanceMonitorPeriodRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceMonitorPeriod")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDBInstanceMonitorPeriodResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEngineParametersModifyHistoryRequest() (request *DescribeEngineParametersModifyHistoryRequest) {
	request = &DescribeEngineParametersModifyHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineParametersModifyHistory")
	return
}

func NewDescribeEngineParametersModifyHistoryResponse() (response *DescribeEngineParametersModifyHistoryResponse) {
	response = &DescribeEngineParametersModifyHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEngineParametersModifyHistory(request *DescribeEngineParametersModifyHistoryRequest) string {
	return c.DescribeEngineParametersModifyHistoryWithContext(context.Background(), request)
}

func (c *Client) DescribeEngineParametersModifyHistorySend(request *DescribeEngineParametersModifyHistoryRequest) (*DescribeEngineParametersModifyHistoryResponse, error) {
	statusCode, msg, err := c.DescribeEngineParametersModifyHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeEngineParametersModifyHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEngineParametersModifyHistoryWithContext(ctx context.Context, request *DescribeEngineParametersModifyHistoryRequest) string {
	if request == nil {
		request = NewDescribeEngineParametersModifyHistoryRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineParametersModifyHistory")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeEngineParametersModifyHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEngineParametersModifyHistoryWithContextV2(ctx context.Context, request *DescribeEngineParametersModifyHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEngineParametersModifyHistoryRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineParametersModifyHistory")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeEngineParametersModifyHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchApplyDBParameterGroupRequest() (request *BatchApplyDBParameterGroupRequest) {
	request = &BatchApplyDBParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "BatchApplyDBParameterGroup")
	return
}

func NewBatchApplyDBParameterGroupResponse() (response *BatchApplyDBParameterGroupResponse) {
	response = &BatchApplyDBParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchApplyDBParameterGroup(request *BatchApplyDBParameterGroupRequest) string {
	return c.BatchApplyDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) BatchApplyDBParameterGroupSend(request *BatchApplyDBParameterGroupRequest) (*BatchApplyDBParameterGroupResponse, error) {
	statusCode, msg, err := c.BatchApplyDBParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BatchApplyDBParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchApplyDBParameterGroupWithContext(ctx context.Context, request *BatchApplyDBParameterGroupRequest) string {
	if request == nil {
		request = NewBatchApplyDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "BatchApplyDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBatchApplyDBParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchApplyDBParameterGroupWithContextV2(ctx context.Context, request *BatchApplyDBParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewBatchApplyDBParameterGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "BatchApplyDBParameterGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBatchApplyDBParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpgradeDBInstanceLatesVersionRequest() (request *UpgradeDBInstanceLatesVersionRequest) {
	request = &UpgradeDBInstanceLatesVersionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceLatesVersion")
	return
}

func NewUpgradeDBInstanceLatesVersionResponse() (response *UpgradeDBInstanceLatesVersionResponse) {
	response = &UpgradeDBInstanceLatesVersionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpgradeDBInstanceLatesVersion(request *UpgradeDBInstanceLatesVersionRequest) string {
	return c.UpgradeDBInstanceLatesVersionWithContext(context.Background(), request)
}

func (c *Client) UpgradeDBInstanceLatesVersionSend(request *UpgradeDBInstanceLatesVersionRequest) (*UpgradeDBInstanceLatesVersionResponse, error) {
	statusCode, msg, err := c.UpgradeDBInstanceLatesVersionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpgradeDBInstanceLatesVersionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpgradeDBInstanceLatesVersionWithContext(ctx context.Context, request *UpgradeDBInstanceLatesVersionRequest) string {
	if request == nil {
		request = NewUpgradeDBInstanceLatesVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceLatesVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpgradeDBInstanceLatesVersionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpgradeDBInstanceLatesVersionWithContextV2(ctx context.Context, request *UpgradeDBInstanceLatesVersionRequest) (int, string, error) {
	if request == nil {
		request = NewUpgradeDBInstanceLatesVersionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceLatesVersion")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpgradeDBInstanceLatesVersionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeProxyInstanceRequest() (request *DescribeProxyInstanceRequest) {
	request = &DescribeProxyInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeProxyInstance")
	return
}

func NewDescribeProxyInstanceResponse() (response *DescribeProxyInstanceResponse) {
	response = &DescribeProxyInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeProxyInstance(request *DescribeProxyInstanceRequest) string {
	return c.DescribeProxyInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeProxyInstanceSend(request *DescribeProxyInstanceRequest) (*DescribeProxyInstanceResponse, error) {
	statusCode, msg, err := c.DescribeProxyInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeProxyInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeProxyInstanceWithContext(ctx context.Context, request *DescribeProxyInstanceRequest) string {
	if request == nil {
		request = NewDescribeProxyInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeProxyInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeProxyInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeProxyInstanceWithContextV2(ctx context.Context, request *DescribeProxyInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeProxyInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeProxyInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeProxyInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetUpProxyInstanceRequest() (request *SetUpProxyInstanceRequest) {
	request = &SetUpProxyInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "SetUpProxyInstance")
	return
}

func NewSetUpProxyInstanceResponse() (response *SetUpProxyInstanceResponse) {
	response = &SetUpProxyInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetUpProxyInstance(request *SetUpProxyInstanceRequest) string {
	return c.SetUpProxyInstanceWithContext(context.Background(), request)
}

func (c *Client) SetUpProxyInstanceSend(request *SetUpProxyInstanceRequest) (*SetUpProxyInstanceResponse, error) {
	statusCode, msg, err := c.SetUpProxyInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetUpProxyInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetUpProxyInstanceWithContext(ctx context.Context, request *SetUpProxyInstanceRequest) string {
	if request == nil {
		request = NewSetUpProxyInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SetUpProxyInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetUpProxyInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetUpProxyInstanceWithContextV2(ctx context.Context, request *SetUpProxyInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewSetUpProxyInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "SetUpProxyInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetUpProxyInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewTemporaryCloseSwitchoverRequest() (request *TemporaryCloseSwitchoverRequest) {
	request = &TemporaryCloseSwitchoverRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "TemporaryCloseSwitchover")
	return
}

func NewTemporaryCloseSwitchoverResponse() (response *TemporaryCloseSwitchoverResponse) {
	response = &TemporaryCloseSwitchoverResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TemporaryCloseSwitchover(request *TemporaryCloseSwitchoverRequest) string {
	return c.TemporaryCloseSwitchoverWithContext(context.Background(), request)
}

func (c *Client) TemporaryCloseSwitchoverSend(request *TemporaryCloseSwitchoverRequest) (*TemporaryCloseSwitchoverResponse, error) {
	statusCode, msg, err := c.TemporaryCloseSwitchoverWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct TemporaryCloseSwitchoverResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) TemporaryCloseSwitchoverWithContext(ctx context.Context, request *TemporaryCloseSwitchoverRequest) string {
	if request == nil {
		request = NewTemporaryCloseSwitchoverRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "TemporaryCloseSwitchover")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTemporaryCloseSwitchoverResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) TemporaryCloseSwitchoverWithContextV2(ctx context.Context, request *TemporaryCloseSwitchoverRequest) (int, string, error) {
	if request == nil {
		request = NewTemporaryCloseSwitchoverRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "TemporaryCloseSwitchover")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTemporaryCloseSwitchoverResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBackupOverviewRequest() (request *DescribeBackupOverviewRequest) {
	request = &DescribeBackupOverviewRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeBackupOverview")
	return
}

func NewDescribeBackupOverviewResponse() (response *DescribeBackupOverviewResponse) {
	response = &DescribeBackupOverviewResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBackupOverview(request *DescribeBackupOverviewRequest) string {
	return c.DescribeBackupOverviewWithContext(context.Background(), request)
}

func (c *Client) DescribeBackupOverviewSend(request *DescribeBackupOverviewRequest) (*DescribeBackupOverviewResponse, error) {
	statusCode, msg, err := c.DescribeBackupOverviewWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBackupOverviewResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBackupOverviewWithContext(ctx context.Context, request *DescribeBackupOverviewRequest) string {
	if request == nil {
		request = NewDescribeBackupOverviewRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeBackupOverview")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackupOverviewResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBackupOverviewWithContextV2(ctx context.Context, request *DescribeBackupOverviewRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBackupOverviewRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeBackupOverview")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackupOverviewResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeStatisticBackupDetailsRequest() (request *DescribeStatisticBackupDetailsRequest) {
	request = &DescribeStatisticBackupDetailsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "DescribeStatisticBackupDetails")
	return
}

func NewDescribeStatisticBackupDetailsResponse() (response *DescribeStatisticBackupDetailsResponse) {
	response = &DescribeStatisticBackupDetailsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeStatisticBackupDetails(request *DescribeStatisticBackupDetailsRequest) string {
	return c.DescribeStatisticBackupDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeStatisticBackupDetailsSend(request *DescribeStatisticBackupDetailsRequest) (*DescribeStatisticBackupDetailsResponse, error) {
	statusCode, msg, err := c.DescribeStatisticBackupDetailsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeStatisticBackupDetailsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeStatisticBackupDetailsWithContext(ctx context.Context, request *DescribeStatisticBackupDetailsRequest) string {
	if request == nil {
		request = NewDescribeStatisticBackupDetailsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeStatisticBackupDetails")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStatisticBackupDetailsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeStatisticBackupDetailsWithContextV2(ctx context.Context, request *DescribeStatisticBackupDetailsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeStatisticBackupDetailsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "DescribeStatisticBackupDetails")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeStatisticBackupDetailsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyMaintenanceTimeRequest() (request *ModifyMaintenanceTimeRequest) {
	request = &ModifyMaintenanceTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyMaintenanceTime")
	return
}

func NewModifyMaintenanceTimeResponse() (response *ModifyMaintenanceTimeResponse) {
	response = &ModifyMaintenanceTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyMaintenanceTime(request *ModifyMaintenanceTimeRequest) string {
	return c.ModifyMaintenanceTimeWithContext(context.Background(), request)
}

func (c *Client) ModifyMaintenanceTimeSend(request *ModifyMaintenanceTimeRequest) (*ModifyMaintenanceTimeResponse, error) {
	statusCode, msg, err := c.ModifyMaintenanceTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyMaintenanceTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyMaintenanceTimeWithContext(ctx context.Context, request *ModifyMaintenanceTimeRequest) string {
	if request == nil {
		request = NewModifyMaintenanceTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyMaintenanceTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyMaintenanceTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyMaintenanceTimeWithContextV2(ctx context.Context, request *ModifyMaintenanceTimeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyMaintenanceTimeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyMaintenanceTime")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyMaintenanceTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceDatabasePrivilegesActionRequest() (request *ModifyInstanceDatabasePrivilegesActionRequest) {
	request = &ModifyInstanceDatabasePrivilegesActionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivilegesAction")
	return
}

func NewModifyInstanceDatabasePrivilegesActionResponse() (response *ModifyInstanceDatabasePrivilegesActionResponse) {
	response = &ModifyInstanceDatabasePrivilegesActionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceDatabasePrivilegesAction(request *ModifyInstanceDatabasePrivilegesActionRequest) string {
	return c.ModifyInstanceDatabasePrivilegesActionWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabasePrivilegesActionSend(request *ModifyInstanceDatabasePrivilegesActionRequest) (*ModifyInstanceDatabasePrivilegesActionResponse, error) {
	statusCode, msg, err := c.ModifyInstanceDatabasePrivilegesActionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstanceDatabasePrivilegesActionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceDatabasePrivilegesActionWithContext(ctx context.Context, request *ModifyInstanceDatabasePrivilegesActionRequest) string {
	if request == nil {
		request = NewModifyInstanceDatabasePrivilegesActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivilegesAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabasePrivilegesActionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceDatabasePrivilegesActionWithContextV2(ctx context.Context, request *ModifyInstanceDatabasePrivilegesActionRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceDatabasePrivilegesActionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivilegesAction")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyInstanceDatabasePrivilegesActionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDBInstanceOrderRequest() (request *UpdateDBInstanceOrderRequest) {
	request = &UpdateDBInstanceOrderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "UpdateDBInstanceOrder")
	return
}

func NewUpdateDBInstanceOrderResponse() (response *UpdateDBInstanceOrderResponse) {
	response = &UpdateDBInstanceOrderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDBInstanceOrder(request *UpdateDBInstanceOrderRequest) string {
	return c.UpdateDBInstanceOrderWithContext(context.Background(), request)
}

func (c *Client) UpdateDBInstanceOrderSend(request *UpdateDBInstanceOrderRequest) (*UpdateDBInstanceOrderResponse, error) {
	statusCode, msg, err := c.UpdateDBInstanceOrderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDBInstanceOrderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDBInstanceOrderWithContext(ctx context.Context, request *UpdateDBInstanceOrderRequest) string {
	if request == nil {
		request = NewUpdateDBInstanceOrderRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpdateDBInstanceOrder")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDBInstanceOrderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDBInstanceOrderWithContextV2(ctx context.Context, request *UpdateDBInstanceOrderRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDBInstanceOrderRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpdateDBInstanceOrder")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDBInstanceOrderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateResourceProtectionRequest() (request *UpdateResourceProtectionRequest) {
	request = &UpdateResourceProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("krds", APIVersion, "UpdateResourceProtection")
	return
}

func NewUpdateResourceProtectionResponse() (response *UpdateResourceProtectionResponse) {
	response = &UpdateResourceProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateResourceProtection(request *UpdateResourceProtectionRequest) string {
	return c.UpdateResourceProtectionWithContext(context.Background(), request)
}

func (c *Client) UpdateResourceProtectionSend(request *UpdateResourceProtectionRequest) (*UpdateResourceProtectionResponse, error) {
	statusCode, msg, err := c.UpdateResourceProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateResourceProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateResourceProtectionWithContext(ctx context.Context, request *UpdateResourceProtectionRequest) string {
	if request == nil {
		request = NewUpdateResourceProtectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpdateResourceProtection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateResourceProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateResourceProtectionWithContextV2(ctx context.Context, request *UpdateResourceProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateResourceProtectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("krds", APIVersion, "UpdateResourceProtection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateResourceProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
