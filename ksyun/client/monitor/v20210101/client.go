package v20210101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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
