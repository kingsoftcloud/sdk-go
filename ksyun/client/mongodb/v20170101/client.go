package v20170101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2017-01-01"

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

func NewCreateMongoDBInstanceRequest() (request *CreateMongoDBInstanceRequest) {
	request = &CreateMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "CreateMongoDBInstance")
	return
}

func NewCreateMongoDBInstanceResponse() (response *CreateMongoDBInstanceResponse) {
	response = &CreateMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMongoDBInstance(request *CreateMongoDBInstanceRequest) string {
	return c.CreateMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateMongoDBInstanceSend(request *CreateMongoDBInstanceRequest) (*CreateMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.CreateMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMongoDBInstanceWithContext(ctx context.Context, request *CreateMongoDBInstanceRequest) string {
	if request == nil {
		request = NewCreateMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMongoDBInstanceWithContextV2(ctx context.Context, request *CreateMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMongoDBInstanceRequest() (request *DeleteMongoDBInstanceRequest) {
	request = &DeleteMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DeleteMongoDBInstance")
	return
}

func NewDeleteMongoDBInstanceResponse() (response *DeleteMongoDBInstanceResponse) {
	response = &DeleteMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMongoDBInstance(request *DeleteMongoDBInstanceRequest) string {
	return c.DeleteMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteMongoDBInstanceSend(request *DeleteMongoDBInstanceRequest) (*DeleteMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.DeleteMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMongoDBInstanceWithContext(ctx context.Context, request *DeleteMongoDBInstanceRequest) string {
	if request == nil {
		request = NewDeleteMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMongoDBInstanceWithContextV2(ctx context.Context, request *DeleteMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMongoDBInstanceRequest() (request *DescribeMongoDBInstanceRequest) {
	request = &DescribeMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBInstance")
	return
}

func NewDescribeMongoDBInstanceResponse() (response *DescribeMongoDBInstanceResponse) {
	response = &DescribeMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMongoDBInstance(request *DescribeMongoDBInstanceRequest) string {
	return c.DescribeMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBInstanceSend(request *DescribeMongoDBInstanceRequest) (*DescribeMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.DescribeMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMongoDBInstanceWithContext(ctx context.Context, request *DescribeMongoDBInstanceRequest) string {
	if request == nil {
		request = NewDescribeMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMongoDBInstanceWithContextV2(ctx context.Context, request *DescribeMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMongoDBInstancesRequest() (request *DescribeMongoDBInstancesRequest) {
	request = &DescribeMongoDBInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBInstances")
	return
}

func NewDescribeMongoDBInstancesResponse() (response *DescribeMongoDBInstancesResponse) {
	response = &DescribeMongoDBInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMongoDBInstances(request *DescribeMongoDBInstancesRequest) string {
	return c.DescribeMongoDBInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBInstancesSend(request *DescribeMongoDBInstancesRequest) (*DescribeMongoDBInstancesResponse, error) {
	statusCode, msg, err := c.DescribeMongoDBInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMongoDBInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMongoDBInstancesWithContext(ctx context.Context, request *DescribeMongoDBInstancesRequest) string {
	if request == nil {
		request = NewDescribeMongoDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMongoDBInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMongoDBInstancesWithContextV2(ctx context.Context, request *DescribeMongoDBInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMongoDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMongoDBInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMongoDBInstanceNodeRequest() (request *DescribeMongoDBInstanceNodeRequest) {
	request = &DescribeMongoDBInstanceNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBInstanceNode")
	return
}

func NewDescribeMongoDBInstanceNodeResponse() (response *DescribeMongoDBInstanceNodeResponse) {
	response = &DescribeMongoDBInstanceNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMongoDBInstanceNode(request *DescribeMongoDBInstanceNodeRequest) string {
	return c.DescribeMongoDBInstanceNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBInstanceNodeSend(request *DescribeMongoDBInstanceNodeRequest) (*DescribeMongoDBInstanceNodeResponse, error) {
	statusCode, msg, err := c.DescribeMongoDBInstanceNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMongoDBInstanceNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMongoDBInstanceNodeWithContext(ctx context.Context, request *DescribeMongoDBInstanceNodeRequest) string {
	if request == nil {
		request = NewDescribeMongoDBInstanceNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMongoDBInstanceNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMongoDBInstanceNodeWithContextV2(ctx context.Context, request *DescribeMongoDBInstanceNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMongoDBInstanceNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMongoDBInstanceNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenameMongoDBInstanceRequest() (request *RenameMongoDBInstanceRequest) {
	request = &RenameMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "RenameMongoDBInstance")
	return
}

func NewRenameMongoDBInstanceResponse() (response *RenameMongoDBInstanceResponse) {
	response = &RenameMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameMongoDBInstance(request *RenameMongoDBInstanceRequest) string {
	return c.RenameMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RenameMongoDBInstanceSend(request *RenameMongoDBInstanceRequest) (*RenameMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.RenameMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RenameMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenameMongoDBInstanceWithContext(ctx context.Context, request *RenameMongoDBInstanceRequest) string {
	if request == nil {
		request = NewRenameMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenameMongoDBInstanceWithContextV2(ctx context.Context, request *RenameMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRenameMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResetPasswordMongoDBInstanceRequest() (request *ResetPasswordMongoDBInstanceRequest) {
	request = &ResetPasswordMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "ResetPasswordMongoDBInstance")
	return
}

func NewResetPasswordMongoDBInstanceResponse() (response *ResetPasswordMongoDBInstanceResponse) {
	response = &ResetPasswordMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetPasswordMongoDBInstance(request *ResetPasswordMongoDBInstanceRequest) string {
	return c.ResetPasswordMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) ResetPasswordMongoDBInstanceSend(request *ResetPasswordMongoDBInstanceRequest) (*ResetPasswordMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.ResetPasswordMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResetPasswordMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResetPasswordMongoDBInstanceWithContext(ctx context.Context, request *ResetPasswordMongoDBInstanceRequest) string {
	if request == nil {
		request = NewResetPasswordMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResetPasswordMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResetPasswordMongoDBInstanceWithContextV2(ctx context.Context, request *ResetPasswordMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewResetPasswordMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResetPasswordMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestartMongoDBInstanceRequest() (request *RestartMongoDBInstanceRequest) {
	request = &RestartMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "RestartMongoDBInstance")
	return
}

func NewRestartMongoDBInstanceResponse() (response *RestartMongoDBInstanceResponse) {
	response = &RestartMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestartMongoDBInstance(request *RestartMongoDBInstanceRequest) string {
	return c.RestartMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RestartMongoDBInstanceSend(request *RestartMongoDBInstanceRequest) (*RestartMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.RestartMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RestartMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestartMongoDBInstanceWithContext(ctx context.Context, request *RestartMongoDBInstanceRequest) string {
	if request == nil {
		request = NewRestartMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestartMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestartMongoDBInstanceWithContextV2(ctx context.Context, request *RestartMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRestartMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestartMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMongoDBSnapshotRequest() (request *CreateMongoDBSnapshotRequest) {
	request = &CreateMongoDBSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "CreateMongoDBSnapshot")
	return
}

func NewCreateMongoDBSnapshotResponse() (response *CreateMongoDBSnapshotResponse) {
	response = &CreateMongoDBSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMongoDBSnapshot(request *CreateMongoDBSnapshotRequest) string {
	return c.CreateMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateMongoDBSnapshotSend(request *CreateMongoDBSnapshotRequest) (*CreateMongoDBSnapshotResponse, error) {
	statusCode, msg, err := c.CreateMongoDBSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateMongoDBSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMongoDBSnapshotWithContext(ctx context.Context, request *CreateMongoDBSnapshotRequest) string {
	if request == nil {
		request = NewCreateMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMongoDBSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMongoDBSnapshotWithContextV2(ctx context.Context, request *CreateMongoDBSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMongoDBSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetMongoDBTimingSnapshotRequest() (request *SetMongoDBTimingSnapshotRequest) {
	request = &SetMongoDBTimingSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "SetMongoDBTimingSnapshot")
	return
}

func NewSetMongoDBTimingSnapshotResponse() (response *SetMongoDBTimingSnapshotResponse) {
	response = &SetMongoDBTimingSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetMongoDBTimingSnapshot(request *SetMongoDBTimingSnapshotRequest) string {
	return c.SetMongoDBTimingSnapshotWithContext(context.Background(), request)
}

func (c *Client) SetMongoDBTimingSnapshotSend(request *SetMongoDBTimingSnapshotRequest) (*SetMongoDBTimingSnapshotResponse, error) {
	statusCode, msg, err := c.SetMongoDBTimingSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetMongoDBTimingSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetMongoDBTimingSnapshotWithContext(ctx context.Context, request *SetMongoDBTimingSnapshotRequest) string {
	if request == nil {
		request = NewSetMongoDBTimingSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetMongoDBTimingSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetMongoDBTimingSnapshotWithContextV2(ctx context.Context, request *SetMongoDBTimingSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewSetMongoDBTimingSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetMongoDBTimingSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMongoDBSnapshotRequest() (request *DescribeMongoDBSnapshotRequest) {
	request = &DescribeMongoDBSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBSnapshot")
	return
}

func NewDescribeMongoDBSnapshotResponse() (response *DescribeMongoDBSnapshotResponse) {
	response = &DescribeMongoDBSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMongoDBSnapshot(request *DescribeMongoDBSnapshotRequest) string {
	return c.DescribeMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBSnapshotSend(request *DescribeMongoDBSnapshotRequest) (*DescribeMongoDBSnapshotResponse, error) {
	statusCode, msg, err := c.DescribeMongoDBSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMongoDBSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMongoDBSnapshotWithContext(ctx context.Context, request *DescribeMongoDBSnapshotRequest) string {
	if request == nil {
		request = NewDescribeMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMongoDBSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMongoDBSnapshotWithContextV2(ctx context.Context, request *DescribeMongoDBSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMongoDBSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMongoDBSnapshotRequest() (request *DeleteMongoDBSnapshotRequest) {
	request = &DeleteMongoDBSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DeleteMongoDBSnapshot")
	return
}

func NewDeleteMongoDBSnapshotResponse() (response *DeleteMongoDBSnapshotResponse) {
	response = &DeleteMongoDBSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMongoDBSnapshot(request *DeleteMongoDBSnapshotRequest) string {
	return c.DeleteMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteMongoDBSnapshotSend(request *DeleteMongoDBSnapshotRequest) (*DeleteMongoDBSnapshotResponse, error) {
	statusCode, msg, err := c.DeleteMongoDBSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteMongoDBSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMongoDBSnapshotWithContext(ctx context.Context, request *DeleteMongoDBSnapshotRequest) string {
	if request == nil {
		request = NewDeleteMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMongoDBSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMongoDBSnapshotWithContextV2(ctx context.Context, request *DeleteMongoDBSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMongoDBSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenameMongoDBSnapshotRequest() (request *RenameMongoDBSnapshotRequest) {
	request = &RenameMongoDBSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "RenameMongoDBSnapshot")
	return
}

func NewRenameMongoDBSnapshotResponse() (response *RenameMongoDBSnapshotResponse) {
	response = &RenameMongoDBSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameMongoDBSnapshot(request *RenameMongoDBSnapshotRequest) string {
	return c.RenameMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) RenameMongoDBSnapshotSend(request *RenameMongoDBSnapshotRequest) (*RenameMongoDBSnapshotResponse, error) {
	statusCode, msg, err := c.RenameMongoDBSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RenameMongoDBSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenameMongoDBSnapshotWithContext(ctx context.Context, request *RenameMongoDBSnapshotRequest) string {
	if request == nil {
		request = NewRenameMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameMongoDBSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenameMongoDBSnapshotWithContextV2(ctx context.Context, request *RenameMongoDBSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewRenameMongoDBSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameMongoDBSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddSecurityGroupRuleRequest() (request *AddSecurityGroupRuleRequest) {
	request = &AddSecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "AddSecurityGroupRule")
	return
}

func NewAddSecurityGroupRuleResponse() (response *AddSecurityGroupRuleResponse) {
	response = &AddSecurityGroupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddSecurityGroupRule(request *AddSecurityGroupRuleRequest) string {
	return c.AddSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) AddSecurityGroupRuleSend(request *AddSecurityGroupRuleRequest) (*AddSecurityGroupRuleResponse, error) {
	statusCode, msg, err := c.AddSecurityGroupRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddSecurityGroupRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddSecurityGroupRuleWithContext(ctx context.Context, request *AddSecurityGroupRuleRequest) string {
	if request == nil {
		request = NewAddSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddSecurityGroupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddSecurityGroupRuleWithContextV2(ctx context.Context, request *AddSecurityGroupRuleRequest) (int, string, error) {
	if request == nil {
		request = NewAddSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddSecurityGroupRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityGroupRulesRequest() (request *DeleteSecurityGroupRulesRequest) {
	request = &DeleteSecurityGroupRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DeleteSecurityGroupRules")
	return
}

func NewDeleteSecurityGroupRulesResponse() (response *DeleteSecurityGroupRulesResponse) {
	response = &DeleteSecurityGroupRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroupRules(request *DeleteSecurityGroupRulesRequest) string {
	return c.DeleteSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupRulesSend(request *DeleteSecurityGroupRulesRequest) (*DeleteSecurityGroupRulesResponse, error) {
	statusCode, msg, err := c.DeleteSecurityGroupRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSecurityGroupRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityGroupRulesWithContext(ctx context.Context, request *DeleteSecurityGroupRulesRequest) string {
	if request == nil {
		request = NewDeleteSecurityGroupRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityGroupRulesWithContextV2(ctx context.Context, request *DeleteSecurityGroupRulesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityGroupRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListSecurityGroupRulesRequest() (request *ListSecurityGroupRulesRequest) {
	request = &ListSecurityGroupRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "ListSecurityGroupRules")
	return
}

func NewListSecurityGroupRulesResponse() (response *ListSecurityGroupRulesResponse) {
	response = &ListSecurityGroupRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListSecurityGroupRules(request *ListSecurityGroupRulesRequest) string {
	return c.ListSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) ListSecurityGroupRulesSend(request *ListSecurityGroupRulesRequest) (*ListSecurityGroupRulesResponse, error) {
	statusCode, msg, err := c.ListSecurityGroupRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListSecurityGroupRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListSecurityGroupRulesWithContext(ctx context.Context, request *ListSecurityGroupRulesRequest) string {
	if request == nil {
		request = NewListSecurityGroupRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecurityGroupRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListSecurityGroupRulesWithContextV2(ctx context.Context, request *ListSecurityGroupRulesRequest) (int, string, error) {
	if request == nil {
		request = NewListSecurityGroupRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListSecurityGroupRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateMongoDBInstanceRequest() (request *UpdateMongoDBInstanceRequest) {
	request = &UpdateMongoDBInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "UpdateMongoDBInstance")
	return
}

func NewUpdateMongoDBInstanceResponse() (response *UpdateMongoDBInstanceResponse) {
	response = &UpdateMongoDBInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateMongoDBInstance(request *UpdateMongoDBInstanceRequest) string {
	return c.UpdateMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) UpdateMongoDBInstanceSend(request *UpdateMongoDBInstanceRequest) (*UpdateMongoDBInstanceResponse, error) {
	statusCode, msg, err := c.UpdateMongoDBInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateMongoDBInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateMongoDBInstanceWithContext(ctx context.Context, request *UpdateMongoDBInstanceRequest) string {
	if request == nil {
		request = NewUpdateMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateMongoDBInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateMongoDBInstanceWithContextV2(ctx context.Context, request *UpdateMongoDBInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateMongoDBInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateMongoDBInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddSecondaryInstanceRequest() (request *AddSecondaryInstanceRequest) {
	request = &AddSecondaryInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "AddSecondaryInstance")
	return
}

func NewAddSecondaryInstanceResponse() (response *AddSecondaryInstanceResponse) {
	response = &AddSecondaryInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddSecondaryInstance(request *AddSecondaryInstanceRequest) string {
	return c.AddSecondaryInstanceWithContext(context.Background(), request)
}

func (c *Client) AddSecondaryInstanceSend(request *AddSecondaryInstanceRequest) (*AddSecondaryInstanceResponse, error) {
	statusCode, msg, err := c.AddSecondaryInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddSecondaryInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddSecondaryInstanceWithContext(ctx context.Context, request *AddSecondaryInstanceRequest) string {
	if request == nil {
		request = NewAddSecondaryInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddSecondaryInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddSecondaryInstanceWithContextV2(ctx context.Context, request *AddSecondaryInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewAddSecondaryInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddSecondaryInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMongoDBShardNodeRequest() (request *DescribeMongoDBShardNodeRequest) {
	request = &DescribeMongoDBShardNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBShardNode")
	return
}

func NewDescribeMongoDBShardNodeResponse() (response *DescribeMongoDBShardNodeResponse) {
	response = &DescribeMongoDBShardNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMongoDBShardNode(request *DescribeMongoDBShardNodeRequest) string {
	return c.DescribeMongoDBShardNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBShardNodeSend(request *DescribeMongoDBShardNodeRequest) (*DescribeMongoDBShardNodeResponse, error) {
	statusCode, msg, err := c.DescribeMongoDBShardNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMongoDBShardNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMongoDBShardNodeWithContext(ctx context.Context, request *DescribeMongoDBShardNodeRequest) string {
	if request == nil {
		request = NewDescribeMongoDBShardNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMongoDBShardNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMongoDBShardNodeWithContextV2(ctx context.Context, request *DescribeMongoDBShardNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMongoDBShardNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMongoDBShardNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeValidRegionRequest() (request *DescribeValidRegionRequest) {
	request = &DescribeValidRegionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeValidRegion")
	return
}

func NewDescribeValidRegionResponse() (response *DescribeValidRegionResponse) {
	response = &DescribeValidRegionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeValidRegion(request *DescribeValidRegionRequest) string {
	return c.DescribeValidRegionWithContext(context.Background(), request)
}

func (c *Client) DescribeValidRegionSend(request *DescribeValidRegionRequest) (*DescribeValidRegionResponse, error) {
	statusCode, msg, err := c.DescribeValidRegionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeValidRegionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeValidRegionWithContext(ctx context.Context, request *DescribeValidRegionRequest) string {
	if request == nil {
		request = NewDescribeValidRegionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeValidRegionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeValidRegionWithContextV2(ctx context.Context, request *DescribeValidRegionRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeValidRegionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeValidRegionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateEipRequest() (request *AllocateEipRequest) {
	request = &AllocateEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "AllocateEip")
	return
}

func NewAllocateEipResponse() (response *AllocateEipResponse) {
	response = &AllocateEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateEip(request *AllocateEipRequest) string {
	return c.AllocateEipWithContext(context.Background(), request)
}

func (c *Client) AllocateEipSend(request *AllocateEipRequest) (*AllocateEipResponse, error) {
	statusCode, msg, err := c.AllocateEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AllocateEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateEipWithContext(ctx context.Context, request *AllocateEipRequest) string {
	if request == nil {
		request = NewAllocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateEipWithContextV2(ctx context.Context, request *AllocateEipRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeallocateEipRequest() (request *DeallocateEipRequest) {
	request = &DeallocateEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DeallocateEip")
	return
}

func NewDeallocateEipResponse() (response *DeallocateEipResponse) {
	response = &DeallocateEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeallocateEip(request *DeallocateEipRequest) string {
	return c.DeallocateEipWithContext(context.Background(), request)
}

func (c *Client) DeallocateEipSend(request *DeallocateEipRequest) (*DeallocateEipResponse, error) {
	statusCode, msg, err := c.DeallocateEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeallocateEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeallocateEipWithContext(ctx context.Context, request *DeallocateEipRequest) string {
	if request == nil {
		request = NewDeallocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeallocateEipWithContextV2(ctx context.Context, request *DeallocateEipRequest) (int, string, error) {
	if request == nil {
		request = NewDeallocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegions(request *DescribeRegionsRequest) string {
	return c.DescribeRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionsSend(request *DescribeRegionsRequest) (*DescribeRegionsResponse, error) {
	statusCode, msg, err := c.DescribeRegionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRegionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) string {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRegionsWithContextV2(ctx context.Context, request *DescribeRegionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRegionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMongoDBShardInstanceRequest() (request *CreateMongoDBShardInstanceRequest) {
	request = &CreateMongoDBShardInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "CreateMongoDBShardInstance")
	return
}

func NewCreateMongoDBShardInstanceResponse() (response *CreateMongoDBShardInstanceResponse) {
	response = &CreateMongoDBShardInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMongoDBShardInstance(request *CreateMongoDBShardInstanceRequest) string {
	return c.CreateMongoDBShardInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateMongoDBShardInstanceSend(request *CreateMongoDBShardInstanceRequest) (*CreateMongoDBShardInstanceResponse, error) {
	statusCode, msg, err := c.CreateMongoDBShardInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateMongoDBShardInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMongoDBShardInstanceWithContext(ctx context.Context, request *CreateMongoDBShardInstanceRequest) string {
	if request == nil {
		request = NewCreateMongoDBShardInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMongoDBShardInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMongoDBShardInstanceWithContextV2(ctx context.Context, request *CreateMongoDBShardInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMongoDBShardInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMongoDBShardInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDownloadSnapshotRequest() (request *DownloadSnapshotRequest) {
	request = &DownloadSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DownloadSnapshot")
	return
}

func NewDownloadSnapshotResponse() (response *DownloadSnapshotResponse) {
	response = &DownloadSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DownloadSnapshot(request *DownloadSnapshotRequest) string {
	return c.DownloadSnapshotWithContext(context.Background(), request)
}

func (c *Client) DownloadSnapshotSend(request *DownloadSnapshotRequest) (*DownloadSnapshotResponse, error) {
	statusCode, msg, err := c.DownloadSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DownloadSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DownloadSnapshotWithContext(ctx context.Context, request *DownloadSnapshotRequest) string {
	if request == nil {
		request = NewDownloadSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDownloadSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DownloadSnapshotWithContextV2(ctx context.Context, request *DownloadSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDownloadSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDownloadSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCloneInstanceRequest() (request *CloneInstanceRequest) {
	request = &CloneInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "CloneInstance")
	return
}

func NewCloneInstanceResponse() (response *CloneInstanceResponse) {
	response = &CloneInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloneInstance(request *CloneInstanceRequest) string {
	return c.CloneInstanceWithContext(context.Background(), request)
}

func (c *Client) CloneInstanceSend(request *CloneInstanceRequest) (*CloneInstanceResponse, error) {
	statusCode, msg, err := c.CloneInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloneInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloneInstanceWithContext(ctx context.Context, request *CloneInstanceRequest) string {
	if request == nil {
		request = NewCloneInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloneInstanceWithContextV2(ctx context.Context, request *CloneInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCloneInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeShardNodeRequest() (request *DescribeShardNodeRequest) {
	request = &DescribeShardNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeShardNode")
	return
}

func NewDescribeShardNodeResponse() (response *DescribeShardNodeResponse) {
	response = &DescribeShardNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeShardNode(request *DescribeShardNodeRequest) string {
	return c.DescribeShardNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeShardNodeSend(request *DescribeShardNodeRequest) (*DescribeShardNodeResponse, error) {
	statusCode, msg, err := c.DescribeShardNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeShardNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeShardNodeWithContext(ctx context.Context, request *DescribeShardNodeRequest) string {
	if request == nil {
		request = NewDescribeShardNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeShardNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeShardNodeWithContextV2(ctx context.Context, request *DescribeShardNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeShardNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeShardNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceStatisticRequest() (request *DescribeInstanceStatisticRequest) {
	request = &DescribeInstanceStatisticRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceStatistic")
	return
}

func NewDescribeInstanceStatisticResponse() (response *DescribeInstanceStatisticResponse) {
	response = &DescribeInstanceStatisticResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceStatistic(request *DescribeInstanceStatisticRequest) string {
	return c.DescribeInstanceStatisticWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceStatisticSend(request *DescribeInstanceStatisticRequest) (*DescribeInstanceStatisticResponse, error) {
	statusCode, msg, err := c.DescribeInstanceStatisticWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceStatisticResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceStatisticWithContext(ctx context.Context, request *DescribeInstanceStatisticRequest) string {
	if request == nil {
		request = NewDescribeInstanceStatisticRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceStatisticResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceStatisticWithContextV2(ctx context.Context, request *DescribeInstanceStatisticRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceStatisticRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceStatisticResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddClusterNodeRequest() (request *AddClusterNodeRequest) {
	request = &AddClusterNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "AddClusterNode")
	return
}

func NewAddClusterNodeResponse() (response *AddClusterNodeResponse) {
	response = &AddClusterNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddClusterNode(request *AddClusterNodeRequest) string {
	return c.AddClusterNodeWithContext(context.Background(), request)
}

func (c *Client) AddClusterNodeSend(request *AddClusterNodeRequest) (*AddClusterNodeResponse, error) {
	statusCode, msg, err := c.AddClusterNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddClusterNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddClusterNodeWithContext(ctx context.Context, request *AddClusterNodeRequest) string {
	if request == nil {
		request = NewAddClusterNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddClusterNodeWithContextV2(ctx context.Context, request *AddClusterNodeRequest) (int, string, error) {
	if request == nil {
		request = NewAddClusterNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteClusterNodeRequest() (request *DeleteClusterNodeRequest) {
	request = &DeleteClusterNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DeleteClusterNode")
	return
}

func NewDeleteClusterNodeResponse() (response *DeleteClusterNodeResponse) {
	response = &DeleteClusterNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteClusterNode(request *DeleteClusterNodeRequest) string {
	return c.DeleteClusterNodeWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterNodeSend(request *DeleteClusterNodeRequest) (*DeleteClusterNodeResponse, error) {
	statusCode, msg, err := c.DeleteClusterNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteClusterNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteClusterNodeWithContext(ctx context.Context, request *DeleteClusterNodeRequest) string {
	if request == nil {
		request = NewDeleteClusterNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteClusterNodeWithContextV2(ctx context.Context, request *DeleteClusterNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteClusterNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSlowLogDetailRequest() (request *DescribeSlowLogDetailRequest) {
	request = &DescribeSlowLogDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogDetail")
	return
}

func NewDescribeSlowLogDetailResponse() (response *DescribeSlowLogDetailResponse) {
	response = &DescribeSlowLogDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSlowLogDetail(request *DescribeSlowLogDetailRequest) string {
	return c.DescribeSlowLogDetailWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogDetailSend(request *DescribeSlowLogDetailRequest) (*DescribeSlowLogDetailResponse, error) {
	statusCode, msg, err := c.DescribeSlowLogDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSlowLogDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSlowLogDetailWithContext(ctx context.Context, request *DescribeSlowLogDetailRequest) string {
	if request == nil {
		request = NewDescribeSlowLogDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSlowLogDetailWithContextV2(ctx context.Context, request *DescribeSlowLogDetailRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSlowLogDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSlowLogStatisticsRequest() (request *DescribeSlowLogStatisticsRequest) {
	request = &DescribeSlowLogStatisticsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogStatistics")
	return
}

func NewDescribeSlowLogStatisticsResponse() (response *DescribeSlowLogStatisticsResponse) {
	response = &DescribeSlowLogStatisticsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSlowLogStatistics(request *DescribeSlowLogStatisticsRequest) string {
	return c.DescribeSlowLogStatisticsWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogStatisticsSend(request *DescribeSlowLogStatisticsRequest) (*DescribeSlowLogStatisticsResponse, error) {
	statusCode, msg, err := c.DescribeSlowLogStatisticsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSlowLogStatisticsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSlowLogStatisticsWithContext(ctx context.Context, request *DescribeSlowLogStatisticsRequest) string {
	if request == nil {
		request = NewDescribeSlowLogStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogStatisticsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSlowLogStatisticsWithContextV2(ctx context.Context, request *DescribeSlowLogStatisticsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSlowLogStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogStatisticsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSlowLogDatabaseRequest() (request *DescribeSlowLogDatabaseRequest) {
	request = &DescribeSlowLogDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogDatabase")
	return
}

func NewDescribeSlowLogDatabaseResponse() (response *DescribeSlowLogDatabaseResponse) {
	response = &DescribeSlowLogDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSlowLogDatabase(request *DescribeSlowLogDatabaseRequest) string {
	return c.DescribeSlowLogDatabaseWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogDatabaseSend(request *DescribeSlowLogDatabaseRequest) (*DescribeSlowLogDatabaseResponse, error) {
	statusCode, msg, err := c.DescribeSlowLogDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSlowLogDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSlowLogDatabaseWithContext(ctx context.Context, request *DescribeSlowLogDatabaseRequest) string {
	if request == nil {
		request = NewDescribeSlowLogDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSlowLogDatabaseWithContextV2(ctx context.Context, request *DescribeSlowLogDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSlowLogDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSlowLogLineChartRequest() (request *DescribeSlowLogLineChartRequest) {
	request = &DescribeSlowLogLineChartRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogLineChart")
	return
}

func NewDescribeSlowLogLineChartResponse() (response *DescribeSlowLogLineChartResponse) {
	response = &DescribeSlowLogLineChartResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSlowLogLineChart(request *DescribeSlowLogLineChartRequest) string {
	return c.DescribeSlowLogLineChartWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogLineChartSend(request *DescribeSlowLogLineChartRequest) (*DescribeSlowLogLineChartResponse, error) {
	statusCode, msg, err := c.DescribeSlowLogLineChartWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSlowLogLineChartResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSlowLogLineChartWithContext(ctx context.Context, request *DescribeSlowLogLineChartRequest) string {
	if request == nil {
		request = NewDescribeSlowLogLineChartRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogLineChartResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSlowLogLineChartWithContextV2(ctx context.Context, request *DescribeSlowLogLineChartRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSlowLogLineChartRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSlowLogLineChartResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateMongoDBInstanceClusterRequest() (request *UpdateMongoDBInstanceClusterRequest) {
	request = &UpdateMongoDBInstanceClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "UpdateMongoDBInstanceCluster")
	return
}

func NewUpdateMongoDBInstanceClusterResponse() (response *UpdateMongoDBInstanceClusterResponse) {
	response = &UpdateMongoDBInstanceClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateMongoDBInstanceCluster(request *UpdateMongoDBInstanceClusterRequest) string {
	return c.UpdateMongoDBInstanceClusterWithContext(context.Background(), request)
}

func (c *Client) UpdateMongoDBInstanceClusterSend(request *UpdateMongoDBInstanceClusterRequest) (*UpdateMongoDBInstanceClusterResponse, error) {
	statusCode, msg, err := c.UpdateMongoDBInstanceClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateMongoDBInstanceClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateMongoDBInstanceClusterWithContext(ctx context.Context, request *UpdateMongoDBInstanceClusterRequest) string {
	if request == nil {
		request = NewUpdateMongoDBInstanceClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateMongoDBInstanceClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateMongoDBInstanceClusterWithContextV2(ctx context.Context, request *UpdateMongoDBInstanceClusterRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateMongoDBInstanceClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateMongoDBInstanceClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeClusterForRestoreRequest() (request *DescribeClusterForRestoreRequest) {
	request = &DescribeClusterForRestoreRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeClusterForRestore")
	return
}

func NewDescribeClusterForRestoreResponse() (response *DescribeClusterForRestoreResponse) {
	response = &DescribeClusterForRestoreResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusterForRestore(request *DescribeClusterForRestoreRequest) string {
	return c.DescribeClusterForRestoreWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterForRestoreSend(request *DescribeClusterForRestoreRequest) (*DescribeClusterForRestoreResponse, error) {
	statusCode, msg, err := c.DescribeClusterForRestoreWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClusterForRestoreResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClusterForRestoreWithContext(ctx context.Context, request *DescribeClusterForRestoreRequest) string {
	if request == nil {
		request = NewDescribeClusterForRestoreRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClusterForRestoreResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClusterForRestoreWithContextV2(ctx context.Context, request *DescribeClusterForRestoreRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterForRestoreRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClusterForRestoreResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDefaultParamsRequest() (request *DescribeDefaultParamsRequest) {
	request = &DescribeDefaultParamsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDefaultParams")
	return
}

func NewDescribeDefaultParamsResponse() (response *DescribeDefaultParamsResponse) {
	response = &DescribeDefaultParamsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDefaultParams(request *DescribeDefaultParamsRequest) string {
	return c.DescribeDefaultParamsWithContext(context.Background(), request)
}

func (c *Client) DescribeDefaultParamsSend(request *DescribeDefaultParamsRequest) (*DescribeDefaultParamsResponse, error) {
	statusCode, msg, err := c.DescribeDefaultParamsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDefaultParamsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDefaultParamsWithContext(ctx context.Context, request *DescribeDefaultParamsRequest) string {
	if request == nil {
		request = NewDescribeDefaultParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDefaultParamsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDefaultParamsWithContextV2(ctx context.Context, request *DescribeDefaultParamsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDefaultParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDefaultParamsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateParamGroupRequest() (request *CreateParamGroupRequest) {
	request = &CreateParamGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "CreateParamGroup")
	return
}

func NewCreateParamGroupResponse() (response *CreateParamGroupResponse) {
	response = &CreateParamGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateParamGroup(request *CreateParamGroupRequest) string {
	return c.CreateParamGroupWithContext(context.Background(), request)
}

func (c *Client) CreateParamGroupSend(request *CreateParamGroupRequest) (*CreateParamGroupResponse, error) {
	statusCode, msg, err := c.CreateParamGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateParamGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateParamGroupWithContext(ctx context.Context, request *CreateParamGroupRequest) string {
	if request == nil {
		request = NewCreateParamGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateParamGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateParamGroupWithContextV2(ctx context.Context, request *CreateParamGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateParamGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateParamGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeParamGroupListRequest() (request *DescribeParamGroupListRequest) {
	request = &DescribeParamGroupListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeParamGroupList")
	return
}

func NewDescribeParamGroupListResponse() (response *DescribeParamGroupListResponse) {
	response = &DescribeParamGroupListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeParamGroupList(request *DescribeParamGroupListRequest) string {
	return c.DescribeParamGroupListWithContext(context.Background(), request)
}

func (c *Client) DescribeParamGroupListSend(request *DescribeParamGroupListRequest) (*DescribeParamGroupListResponse, error) {
	statusCode, msg, err := c.DescribeParamGroupListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeParamGroupListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeParamGroupListWithContext(ctx context.Context, request *DescribeParamGroupListRequest) string {
	if request == nil {
		request = NewDescribeParamGroupListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeParamGroupListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeParamGroupListWithContextV2(ctx context.Context, request *DescribeParamGroupListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeParamGroupListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeParamGroupListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeParamGroupInfoRequest() (request *DescribeParamGroupInfoRequest) {
	request = &DescribeParamGroupInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeParamGroupInfo")
	return
}

func NewDescribeParamGroupInfoResponse() (response *DescribeParamGroupInfoResponse) {
	response = &DescribeParamGroupInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeParamGroupInfo(request *DescribeParamGroupInfoRequest) string {
	return c.DescribeParamGroupInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeParamGroupInfoSend(request *DescribeParamGroupInfoRequest) (*DescribeParamGroupInfoResponse, error) {
	statusCode, msg, err := c.DescribeParamGroupInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeParamGroupInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeParamGroupInfoWithContext(ctx context.Context, request *DescribeParamGroupInfoRequest) string {
	if request == nil {
		request = NewDescribeParamGroupInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeParamGroupInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeParamGroupInfoWithContextV2(ctx context.Context, request *DescribeParamGroupInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeParamGroupInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeParamGroupInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModifyHistoryRequest() (request *DescribeModifyHistoryRequest) {
	request = &DescribeModifyHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeModifyHistory")
	return
}

func NewDescribeModifyHistoryResponse() (response *DescribeModifyHistoryResponse) {
	response = &DescribeModifyHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModifyHistory(request *DescribeModifyHistoryRequest) string {
	return c.DescribeModifyHistoryWithContext(context.Background(), request)
}

func (c *Client) DescribeModifyHistorySend(request *DescribeModifyHistoryRequest) (*DescribeModifyHistoryResponse, error) {
	statusCode, msg, err := c.DescribeModifyHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeModifyHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModifyHistoryWithContext(ctx context.Context, request *DescribeModifyHistoryRequest) string {
	if request == nil {
		request = NewDescribeModifyHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModifyHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModifyHistoryWithContextV2(ctx context.Context, request *DescribeModifyHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModifyHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModifyHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
	request = &DescribeInstanceParamsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceParams")
	return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
	response = &DescribeInstanceParamsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) string {
	return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceParamsSend(request *DescribeInstanceParamsRequest) (*DescribeInstanceParamsResponse, error) {
	statusCode, msg, err := c.DescribeInstanceParamsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceParamsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) string {
	if request == nil {
		request = NewDescribeInstanceParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceParamsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceParamsWithContextV2(ctx context.Context, request *DescribeInstanceParamsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceParamsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyParamGroupRequest() (request *ModifyParamGroupRequest) {
	request = &ModifyParamGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "ModifyParamGroup")
	return
}

func NewModifyParamGroupResponse() (response *ModifyParamGroupResponse) {
	response = &ModifyParamGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyParamGroup(request *ModifyParamGroupRequest) string {
	return c.ModifyParamGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyParamGroupSend(request *ModifyParamGroupRequest) (*ModifyParamGroupResponse, error) {
	statusCode, msg, err := c.ModifyParamGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyParamGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyParamGroupWithContext(ctx context.Context, request *ModifyParamGroupRequest) string {
	if request == nil {
		request = NewModifyParamGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyParamGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyParamGroupWithContextV2(ctx context.Context, request *ModifyParamGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyParamGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyParamGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteParamGroupRequest() (request *DeleteParamGroupRequest) {
	request = &DeleteParamGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("mongodb", APIVersion, "DeleteParamGroup")
	return
}

func NewDeleteParamGroupResponse() (response *DeleteParamGroupResponse) {
	response = &DeleteParamGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteParamGroup(request *DeleteParamGroupRequest) string {
	return c.DeleteParamGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteParamGroupSend(request *DeleteParamGroupRequest) (*DeleteParamGroupResponse, error) {
	statusCode, msg, err := c.DeleteParamGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteParamGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteParamGroupWithContext(ctx context.Context, request *DeleteParamGroupRequest) string {
	if request == nil {
		request = NewDeleteParamGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteParamGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteParamGroupWithContextV2(ctx context.Context, request *DeleteParamGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteParamGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteParamGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
