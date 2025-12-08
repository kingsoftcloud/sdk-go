package v20210101
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-01-01"

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

func NewListAlarmPolicyRequest() (request *ListAlarmPolicyRequest) {
	request = &ListAlarmPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "ListAlarmPolicy")
	return
}

func NewListAlarmPolicyResponse() (response *ListAlarmPolicyResponse) {
	response = &ListAlarmPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListAlarmPolicy(request *ListAlarmPolicyRequest) string {
	return c.ListAlarmPolicyWithContext(context.Background(), request)
}

func (c *Client) ListAlarmPolicySend(request *ListAlarmPolicyRequest) (*ListAlarmPolicyResponse, error) {
	statusCode, msg, err := c.ListAlarmPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListAlarmPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListAlarmPolicyWithContext(ctx context.Context, request *ListAlarmPolicyRequest) string {
	if request == nil {
		request = NewListAlarmPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListAlarmPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListAlarmPolicyWithContextV2(ctx context.Context, request *ListAlarmPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewListAlarmPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListAlarmPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlarmPolicyRequest() (request *DescribeAlarmPolicyRequest) {
	request = &DescribeAlarmPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicy")
	return
}

func NewDescribeAlarmPolicyResponse() (response *DescribeAlarmPolicyResponse) {
	response = &DescribeAlarmPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) string {
	return c.DescribeAlarmPolicyWithContext(context.Background(), request)
}

func (c *Client) DescribeAlarmPolicySend(request *DescribeAlarmPolicyRequest) (*DescribeAlarmPolicyResponse, error) {
	statusCode, msg, err := c.DescribeAlarmPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAlarmPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlarmPolicyWithContext(ctx context.Context, request *DescribeAlarmPolicyRequest) string {
	if request == nil {
		request = NewDescribeAlarmPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAlarmPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlarmPolicyWithContextV2(ctx context.Context, request *DescribeAlarmPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlarmPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAlarmPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePolicyObjectRequest() (request *DescribePolicyObjectRequest) {
	request = &DescribePolicyObjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyObject")
	return
}

func NewDescribePolicyObjectResponse() (response *DescribePolicyObjectResponse) {
	response = &DescribePolicyObjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePolicyObject(request *DescribePolicyObjectRequest) string {
	return c.DescribePolicyObjectWithContext(context.Background(), request)
}

func (c *Client) DescribePolicyObjectSend(request *DescribePolicyObjectRequest) (*DescribePolicyObjectResponse, error) {
	statusCode, msg, err := c.DescribePolicyObjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePolicyObjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePolicyObjectWithContext(ctx context.Context, request *DescribePolicyObjectRequest) string {
	if request == nil {
		request = NewDescribePolicyObjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePolicyObjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePolicyObjectWithContextV2(ctx context.Context, request *DescribePolicyObjectRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePolicyObjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePolicyObjectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlarmReceivesRequest() (request *DescribeAlarmReceivesRequest) {
	request = &DescribeAlarmReceivesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmReceives")
	return
}

func NewDescribeAlarmReceivesResponse() (response *DescribeAlarmReceivesResponse) {
	response = &DescribeAlarmReceivesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlarmReceives(request *DescribeAlarmReceivesRequest) string {
	return c.DescribeAlarmReceivesWithContext(context.Background(), request)
}

func (c *Client) DescribeAlarmReceivesSend(request *DescribeAlarmReceivesRequest) (*DescribeAlarmReceivesResponse, error) {
	statusCode, msg, err := c.DescribeAlarmReceivesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAlarmReceivesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlarmReceivesWithContext(ctx context.Context, request *DescribeAlarmReceivesRequest) string {
	if request == nil {
		request = NewDescribeAlarmReceivesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAlarmReceivesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlarmReceivesWithContextV2(ctx context.Context, request *DescribeAlarmReceivesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlarmReceivesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAlarmReceivesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddAlarmReceivesRequest() (request *AddAlarmReceivesRequest) {
	request = &AddAlarmReceivesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "AddAlarmReceives")
	return
}

func NewAddAlarmReceivesResponse() (response *AddAlarmReceivesResponse) {
	response = &AddAlarmReceivesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddAlarmReceives(request *AddAlarmReceivesRequest) string {
	return c.AddAlarmReceivesWithContext(context.Background(), request)
}

func (c *Client) AddAlarmReceivesSend(request *AddAlarmReceivesRequest) (*AddAlarmReceivesResponse, error) {
	statusCode, msg, err := c.AddAlarmReceivesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddAlarmReceivesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddAlarmReceivesWithContext(ctx context.Context, request *AddAlarmReceivesRequest) string {
	if request == nil {
		request = NewAddAlarmReceivesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddAlarmReceivesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddAlarmReceivesWithContextV2(ctx context.Context, request *AddAlarmReceivesRequest) (int, string, error) {
	if request == nil {
		request = NewAddAlarmReceivesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddAlarmReceivesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlarmReceivesRequest() (request *DeleteAlarmReceivesRequest) {
	request = &DeleteAlarmReceivesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmReceives")
	return
}

func NewDeleteAlarmReceivesResponse() (response *DeleteAlarmReceivesResponse) {
	response = &DeleteAlarmReceivesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlarmReceives(request *DeleteAlarmReceivesRequest) string {
	return c.DeleteAlarmReceivesWithContext(context.Background(), request)
}

func (c *Client) DeleteAlarmReceivesSend(request *DeleteAlarmReceivesRequest) (*DeleteAlarmReceivesResponse, error) {
	statusCode, msg, err := c.DeleteAlarmReceivesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteAlarmReceivesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlarmReceivesWithContext(ctx context.Context, request *DeleteAlarmReceivesRequest) string {
	if request == nil {
		request = NewDeleteAlarmReceivesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlarmReceivesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlarmReceivesWithContextV2(ctx context.Context, request *DeleteAlarmReceivesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlarmReceivesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlarmReceivesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetUserGroupRequest() (request *GetUserGroupRequest) {
	request = &GetUserGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetUserGroup")
	return
}

func NewGetUserGroupResponse() (response *GetUserGroupResponse) {
	response = &GetUserGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUserGroup(request *GetUserGroupRequest) string {
	return c.GetUserGroupWithContext(context.Background(), request)
}

func (c *Client) GetUserGroupSend(request *GetUserGroupRequest) (*GetUserGroupResponse, error) {
	statusCode, msg, err := c.GetUserGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetUserGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetUserGroupWithContext(ctx context.Context, request *GetUserGroupRequest) string {
	if request == nil {
		request = NewGetUserGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetUserGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetUserGroupWithContextV2(ctx context.Context, request *GetUserGroupRequest) (int, string, error) {
	if request == nil {
		request = NewGetUserGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetUserGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetAlertUserRequest() (request *GetAlertUserRequest) {
	request = &GetAlertUserRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetAlertUser")
	return
}

func NewGetAlertUserResponse() (response *GetAlertUserResponse) {
	response = &GetAlertUserResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetAlertUser(request *GetAlertUserRequest) string {
	return c.GetAlertUserWithContext(context.Background(), request)
}

func (c *Client) GetAlertUserSend(request *GetAlertUserRequest) (*GetAlertUserResponse, error) {
	statusCode, msg, err := c.GetAlertUserWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetAlertUserResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetAlertUserWithContext(ctx context.Context, request *GetAlertUserRequest) string {
	if request == nil {
		request = NewGetAlertUserRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetAlertUserResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetAlertUserWithContextV2(ctx context.Context, request *GetAlertUserRequest) (int, string, error) {
	if request == nil {
		request = NewGetAlertUserRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetAlertUserResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateAlertUserStatusRequest() (request *UpdateAlertUserStatusRequest) {
	request = &UpdateAlertUserStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "UpdateAlertUserStatus")
	return
}

func NewUpdateAlertUserStatusResponse() (response *UpdateAlertUserStatusResponse) {
	response = &UpdateAlertUserStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateAlertUserStatus(request *UpdateAlertUserStatusRequest) string {
	return c.UpdateAlertUserStatusWithContext(context.Background(), request)
}

func (c *Client) UpdateAlertUserStatusSend(request *UpdateAlertUserStatusRequest) (*UpdateAlertUserStatusResponse, error) {
	statusCode, msg, err := c.UpdateAlertUserStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateAlertUserStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateAlertUserStatusWithContext(ctx context.Context, request *UpdateAlertUserStatusRequest) string {
	if request == nil {
		request = NewUpdateAlertUserStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateAlertUserStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateAlertUserStatusWithContextV2(ctx context.Context, request *UpdateAlertUserStatusRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateAlertUserStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateAlertUserStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSysEventGroupListRequest() (request *DescribeSysEventGroupListRequest) {
	request = &DescribeSysEventGroupListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeSysEventGroupList")
	return
}

func NewDescribeSysEventGroupListResponse() (response *DescribeSysEventGroupListResponse) {
	response = &DescribeSysEventGroupListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSysEventGroupList(request *DescribeSysEventGroupListRequest) string {
	return c.DescribeSysEventGroupListWithContext(context.Background(), request)
}

func (c *Client) DescribeSysEventGroupListSend(request *DescribeSysEventGroupListRequest) (*DescribeSysEventGroupListResponse, error) {
	statusCode, msg, err := c.DescribeSysEventGroupListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSysEventGroupListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSysEventGroupListWithContext(ctx context.Context, request *DescribeSysEventGroupListRequest) string {
	if request == nil {
		request = NewDescribeSysEventGroupListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSysEventGroupListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSysEventGroupListWithContextV2(ctx context.Context, request *DescribeSysEventGroupListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSysEventGroupListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSysEventGroupListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMonitorProductListRequest() (request *DescribeMonitorProductListRequest) {
	request = &DescribeMonitorProductListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorProductList")
	return
}

func NewDescribeMonitorProductListResponse() (response *DescribeMonitorProductListResponse) {
	response = &DescribeMonitorProductListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorProductList(request *DescribeMonitorProductListRequest) string {
	return c.DescribeMonitorProductListWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorProductListSend(request *DescribeMonitorProductListRequest) (*DescribeMonitorProductListResponse, error) {
	statusCode, msg, err := c.DescribeMonitorProductListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMonitorProductListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMonitorProductListWithContext(ctx context.Context, request *DescribeMonitorProductListRequest) string {
	if request == nil {
		request = NewDescribeMonitorProductListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMonitorProductListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMonitorProductListWithContextV2(ctx context.Context, request *DescribeMonitorProductListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMonitorProductListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMonitorProductListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


