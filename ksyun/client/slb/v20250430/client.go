package v20250430
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-04-30"

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

func NewDescribeBackendServersRequest() (request *DescribeBackendServersRequest) {
	request = &DescribeBackendServersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeBackendServers")
	return
}

func NewDescribeBackendServersResponse() (response *DescribeBackendServersResponse) {
	response = &DescribeBackendServersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBackendServers(request *DescribeBackendServersRequest) string {
	return c.DescribeBackendServersWithContext(context.Background(), request)
}

func (c *Client) DescribeBackendServersSend(request *DescribeBackendServersRequest) (*DescribeBackendServersResponse, error) {
	statusCode, msg, err := c.DescribeBackendServersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeBackendServersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBackendServersWithContext(ctx context.Context, request *DescribeBackendServersRequest) string {
	if request == nil {
		request = NewDescribeBackendServersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackendServersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBackendServersWithContextV2(ctx context.Context, request *DescribeBackendServersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBackendServersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackendServersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyBackendServerRequest() (request *ModifyBackendServerRequest) {
	request = &ModifyBackendServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyBackendServer")
	return
}

func NewModifyBackendServerResponse() (response *ModifyBackendServerResponse) {
	response = &ModifyBackendServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBackendServer(request *ModifyBackendServerRequest) string {
	return c.ModifyBackendServerWithContext(context.Background(), request)
}

func (c *Client) ModifyBackendServerSend(request *ModifyBackendServerRequest) (*ModifyBackendServerResponse, error) {
	statusCode, msg, err := c.ModifyBackendServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyBackendServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyBackendServerWithContext(ctx context.Context, request *ModifyBackendServerRequest) string {
	if request == nil {
		request = NewModifyBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBackendServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyBackendServerWithContextV2(ctx context.Context, request *ModifyBackendServerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBackendServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeregisterBackendServerRequest() (request *DeregisterBackendServerRequest) {
	request = &DeregisterBackendServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeregisterBackendServer")
	return
}

func NewDeregisterBackendServerResponse() (response *DeregisterBackendServerResponse) {
	response = &DeregisterBackendServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeregisterBackendServer(request *DeregisterBackendServerRequest) string {
	return c.DeregisterBackendServerWithContext(context.Background(), request)
}

func (c *Client) DeregisterBackendServerSend(request *DeregisterBackendServerRequest) (*DeregisterBackendServerResponse, error) {
	statusCode, msg, err := c.DeregisterBackendServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeregisterBackendServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeregisterBackendServerWithContext(ctx context.Context, request *DeregisterBackendServerRequest) string {
	if request == nil {
		request = NewDeregisterBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeregisterBackendServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeregisterBackendServerWithContextV2(ctx context.Context, request *DeregisterBackendServerRequest) (int, string, error) {
	if request == nil {
		request = NewDeregisterBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeregisterBackendServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRegisterBackendServerRequest() (request *RegisterBackendServerRequest) {
	request = &RegisterBackendServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "RegisterBackendServer")
	return
}

func NewRegisterBackendServerResponse() (response *RegisterBackendServerResponse) {
	response = &RegisterBackendServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RegisterBackendServer(request *RegisterBackendServerRequest) string {
	return c.RegisterBackendServerWithContext(context.Background(), request)
}

func (c *Client) RegisterBackendServerSend(request *RegisterBackendServerRequest) (*RegisterBackendServerResponse, error) {
	statusCode, msg, err := c.RegisterBackendServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RegisterBackendServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RegisterBackendServerWithContext(ctx context.Context, request *RegisterBackendServerRequest) string {
	if request == nil {
		request = NewRegisterBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterBackendServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RegisterBackendServerWithContextV2(ctx context.Context, request *RegisterBackendServerRequest) (int, string, error) {
	if request == nil {
		request = NewRegisterBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterBackendServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBackendServerGroupsRequest() (request *DescribeBackendServerGroupsRequest) {
	request = &DescribeBackendServerGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeBackendServerGroups")
	return
}

func NewDescribeBackendServerGroupsResponse() (response *DescribeBackendServerGroupsResponse) {
	response = &DescribeBackendServerGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBackendServerGroups(request *DescribeBackendServerGroupsRequest) string {
	return c.DescribeBackendServerGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeBackendServerGroupsSend(request *DescribeBackendServerGroupsRequest) (*DescribeBackendServerGroupsResponse, error) {
	statusCode, msg, err := c.DescribeBackendServerGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeBackendServerGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBackendServerGroupsWithContext(ctx context.Context, request *DescribeBackendServerGroupsRequest) string {
	if request == nil {
		request = NewDescribeBackendServerGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackendServerGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBackendServerGroupsWithContextV2(ctx context.Context, request *DescribeBackendServerGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBackendServerGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackendServerGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyBackendServerGroupRequest() (request *ModifyBackendServerGroupRequest) {
	request = &ModifyBackendServerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyBackendServerGroup")
	return
}

func NewModifyBackendServerGroupResponse() (response *ModifyBackendServerGroupResponse) {
	response = &ModifyBackendServerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBackendServerGroup(request *ModifyBackendServerGroupRequest) string {
	return c.ModifyBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyBackendServerGroupSend(request *ModifyBackendServerGroupRequest) (*ModifyBackendServerGroupResponse, error) {
	statusCode, msg, err := c.ModifyBackendServerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyBackendServerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyBackendServerGroupWithContext(ctx context.Context, request *ModifyBackendServerGroupRequest) string {
	if request == nil {
		request = NewModifyBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBackendServerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyBackendServerGroupWithContextV2(ctx context.Context, request *ModifyBackendServerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBackendServerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBackendServerGroupRequest() (request *DeleteBackendServerGroupRequest) {
	request = &DeleteBackendServerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteBackendServerGroup")
	return
}

func NewDeleteBackendServerGroupResponse() (response *DeleteBackendServerGroupResponse) {
	response = &DeleteBackendServerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBackendServerGroup(request *DeleteBackendServerGroupRequest) string {
	return c.DeleteBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteBackendServerGroupSend(request *DeleteBackendServerGroupRequest) (*DeleteBackendServerGroupResponse, error) {
	statusCode, msg, err := c.DeleteBackendServerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteBackendServerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBackendServerGroupWithContext(ctx context.Context, request *DeleteBackendServerGroupRequest) string {
	if request == nil {
		request = NewDeleteBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBackendServerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBackendServerGroupWithContextV2(ctx context.Context, request *DeleteBackendServerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBackendServerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateBackendServerGroupRequest() (request *CreateBackendServerGroupRequest) {
	request = &CreateBackendServerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateBackendServerGroup")
	return
}

func NewCreateBackendServerGroupResponse() (response *CreateBackendServerGroupResponse) {
	response = &CreateBackendServerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateBackendServerGroup(request *CreateBackendServerGroupRequest) string {
	return c.CreateBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) CreateBackendServerGroupSend(request *CreateBackendServerGroupRequest) (*CreateBackendServerGroupResponse, error) {
	statusCode, msg, err := c.CreateBackendServerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateBackendServerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateBackendServerGroupWithContext(ctx context.Context, request *CreateBackendServerGroupRequest) string {
	if request == nil {
		request = NewCreateBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateBackendServerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateBackendServerGroupWithContextV2(ctx context.Context, request *CreateBackendServerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateBackendServerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
	request = &DescribeListenersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeListeners")
	return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
	response = &DescribeListenersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeListeners(request *DescribeListenersRequest) string {
	return c.DescribeListenersWithContext(context.Background(), request)
}

func (c *Client) DescribeListenersSend(request *DescribeListenersRequest) (*DescribeListenersResponse, error) {
	statusCode, msg, err := c.DescribeListenersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeListenersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) string {
	if request == nil {
		request = NewDescribeListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeListenersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeListenersWithContextV2(ctx context.Context, request *DescribeListenersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeListenersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyListenerRequest() (request *ModifyListenerRequest) {
	request = &ModifyListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyListener")
	return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
	response = &ModifyListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyListener(request *ModifyListenerRequest) string {
	return c.ModifyListenerWithContext(context.Background(), request)
}

func (c *Client) ModifyListenerSend(request *ModifyListenerRequest) (*ModifyListenerResponse, error) {
	statusCode, msg, err := c.ModifyListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyListenerWithContext(ctx context.Context, request *ModifyListenerRequest) string {
	if request == nil {
		request = NewModifyListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyListenerWithContextV2(ctx context.Context, request *ModifyListenerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
	request = &DeleteListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteListener")
	return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
	response = &DeleteListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteListener(request *DeleteListenerRequest) string {
	return c.DeleteListenerWithContext(context.Background(), request)
}

func (c *Client) DeleteListenerSend(request *DeleteListenerRequest) (*DeleteListenerResponse, error) {
	statusCode, msg, err := c.DeleteListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) string {
	if request == nil {
		request = NewDeleteListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteListenerWithContextV2(ctx context.Context, request *DeleteListenerRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateListenerRequest() (request *CreateListenerRequest) {
	request = &CreateListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateListener")
	return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
	response = &CreateListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateListener(request *CreateListenerRequest) string {
	return c.CreateListenerWithContext(context.Background(), request)
}

func (c *Client) CreateListenerSend(request *CreateListenerRequest) (*CreateListenerResponse, error) {
	statusCode, msg, err := c.CreateListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) string {
	if request == nil {
		request = NewCreateListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateListenerWithContextV2(ctx context.Context, request *CreateListenerRequest) (int, string, error) {
	if request == nil {
		request = NewCreateListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetAccessLogRequest() (request *SetAccessLogRequest) {
	request = &SetAccessLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetAccessLog")
	return
}

func NewSetAccessLogResponse() (response *SetAccessLogResponse) {
	response = &SetAccessLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetAccessLog(request *SetAccessLogRequest) string {
	return c.SetAccessLogWithContext(context.Background(), request)
}

func (c *Client) SetAccessLogSend(request *SetAccessLogRequest) (*SetAccessLogResponse, error) {
	statusCode, msg, err := c.SetAccessLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetAccessLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetAccessLogWithContext(ctx context.Context, request *SetAccessLogRequest) string {
	if request == nil {
		request = NewSetAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAccessLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetAccessLogWithContextV2(ctx context.Context, request *SetAccessLogRequest) (int, string, error) {
	if request == nil {
		request = NewSetAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAccessLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetEnableAccessLogRequest() (request *SetEnableAccessLogRequest) {
	request = &SetEnableAccessLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetEnableAccessLog")
	return
}

func NewSetEnableAccessLogResponse() (response *SetEnableAccessLogResponse) {
	response = &SetEnableAccessLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetEnableAccessLog(request *SetEnableAccessLogRequest) string {
	return c.SetEnableAccessLogWithContext(context.Background(), request)
}

func (c *Client) SetEnableAccessLogSend(request *SetEnableAccessLogRequest) (*SetEnableAccessLogResponse, error) {
	statusCode, msg, err := c.SetEnableAccessLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetEnableAccessLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetEnableAccessLogWithContext(ctx context.Context, request *SetEnableAccessLogRequest) string {
	if request == nil {
		request = NewSetEnableAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetEnableAccessLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetEnableAccessLogWithContextV2(ctx context.Context, request *SetEnableAccessLogRequest) (int, string, error) {
	if request == nil {
		request = NewSetEnableAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetEnableAccessLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetLbProtocolLayersRequest() (request *SetLbProtocolLayersRequest) {
	request = &SetLbProtocolLayersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetLbProtocolLayers")
	return
}

func NewSetLbProtocolLayersResponse() (response *SetLbProtocolLayersResponse) {
	response = &SetLbProtocolLayersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetLbProtocolLayers(request *SetLbProtocolLayersRequest) string {
	return c.SetLbProtocolLayersWithContext(context.Background(), request)
}

func (c *Client) SetLbProtocolLayersSend(request *SetLbProtocolLayersRequest) (*SetLbProtocolLayersResponse, error) {
	statusCode, msg, err := c.SetLbProtocolLayersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetLbProtocolLayersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetLbProtocolLayersWithContext(ctx context.Context, request *SetLbProtocolLayersRequest) string {
	if request == nil {
		request = NewSetLbProtocolLayersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLbProtocolLayersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetLbProtocolLayersWithContextV2(ctx context.Context, request *SetLbProtocolLayersRequest) (int, string, error) {
	if request == nil {
		request = NewSetLbProtocolLayersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLbProtocolLayersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetLoadBalancerStatusRequest() (request *SetLoadBalancerStatusRequest) {
	request = &SetLoadBalancerStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetLoadBalancerStatus")
	return
}

func NewSetLoadBalancerStatusResponse() (response *SetLoadBalancerStatusResponse) {
	response = &SetLoadBalancerStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetLoadBalancerStatus(request *SetLoadBalancerStatusRequest) string {
	return c.SetLoadBalancerStatusWithContext(context.Background(), request)
}

func (c *Client) SetLoadBalancerStatusSend(request *SetLoadBalancerStatusRequest) (*SetLoadBalancerStatusResponse, error) {
	statusCode, msg, err := c.SetLoadBalancerStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetLoadBalancerStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetLoadBalancerStatusWithContext(ctx context.Context, request *SetLoadBalancerStatusRequest) string {
	if request == nil {
		request = NewSetLoadBalancerStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLoadBalancerStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetLoadBalancerStatusWithContextV2(ctx context.Context, request *SetLoadBalancerStatusRequest) (int, string, error) {
	if request == nil {
		request = NewSetLoadBalancerStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLoadBalancerStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetLoadBalancerNameRequest() (request *SetLoadBalancerNameRequest) {
	request = &SetLoadBalancerNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetLoadBalancerName")
	return
}

func NewSetLoadBalancerNameResponse() (response *SetLoadBalancerNameResponse) {
	response = &SetLoadBalancerNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetLoadBalancerName(request *SetLoadBalancerNameRequest) string {
	return c.SetLoadBalancerNameWithContext(context.Background(), request)
}

func (c *Client) SetLoadBalancerNameSend(request *SetLoadBalancerNameRequest) (*SetLoadBalancerNameResponse, error) {
	statusCode, msg, err := c.SetLoadBalancerNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetLoadBalancerNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetLoadBalancerNameWithContext(ctx context.Context, request *SetLoadBalancerNameRequest) string {
	if request == nil {
		request = NewSetLoadBalancerNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLoadBalancerNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetLoadBalancerNameWithContextV2(ctx context.Context, request *SetLoadBalancerNameRequest) (int, string, error) {
	if request == nil {
		request = NewSetLoadBalancerNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLoadBalancerNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
	request = &DescribeLoadBalancersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeLoadBalancers")
	return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
	response = &DescribeLoadBalancersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) string {
	return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

func (c *Client) DescribeLoadBalancersSend(request *DescribeLoadBalancersRequest) (*DescribeLoadBalancersResponse, error) {
	statusCode, msg, err := c.DescribeLoadBalancersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeLoadBalancersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) string {
	if request == nil {
		request = NewDescribeLoadBalancersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLoadBalancersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeLoadBalancersWithContextV2(ctx context.Context, request *DescribeLoadBalancersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeLoadBalancersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLoadBalancersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteLoadBalancer")
	return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
	response = &DeleteLoadBalancerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) string {
	return c.DeleteLoadBalancerWithContext(context.Background(), request)
}

func (c *Client) DeleteLoadBalancerSend(request *DeleteLoadBalancerRequest) (*DeleteLoadBalancerResponse, error) {
	statusCode, msg, err := c.DeleteLoadBalancerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteLoadBalancerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteLoadBalancerWithContext(ctx context.Context, request *DeleteLoadBalancerRequest) string {
	if request == nil {
		request = NewDeleteLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLoadBalancerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteLoadBalancerWithContextV2(ctx context.Context, request *DeleteLoadBalancerRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLoadBalancerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateLoadBalancer")
	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) string {
	return c.CreateLoadBalancerWithContext(context.Background(), request)
}

func (c *Client) CreateLoadBalancerSend(request *CreateLoadBalancerRequest) (*CreateLoadBalancerResponse, error) {
	statusCode, msg, err := c.CreateLoadBalancerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateLoadBalancerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) string {
	if request == nil {
		request = NewCreateLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLoadBalancerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateLoadBalancerWithContextV2(ctx context.Context, request *CreateLoadBalancerRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLoadBalancerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCertificateWithGroupRequest() (request *ModifyCertificateWithGroupRequest) {
	request = &ModifyCertificateWithGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyCertificateWithGroup")
	return
}

func NewModifyCertificateWithGroupResponse() (response *ModifyCertificateWithGroupResponse) {
	response = &ModifyCertificateWithGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCertificateWithGroup(request *ModifyCertificateWithGroupRequest) string {
	return c.ModifyCertificateWithGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyCertificateWithGroupSend(request *ModifyCertificateWithGroupRequest) (*ModifyCertificateWithGroupResponse, error) {
	statusCode, msg, err := c.ModifyCertificateWithGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyCertificateWithGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCertificateWithGroupWithContext(ctx context.Context, request *ModifyCertificateWithGroupRequest) string {
	if request == nil {
		request = NewModifyCertificateWithGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCertificateWithGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCertificateWithGroupWithContextV2(ctx context.Context, request *ModifyCertificateWithGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCertificateWithGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCertificateWithGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDissociateCertificateWithGroupRequest() (request *DissociateCertificateWithGroupRequest) {
	request = &DissociateCertificateWithGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DissociateCertificateWithGroup")
	return
}

func NewDissociateCertificateWithGroupResponse() (response *DissociateCertificateWithGroupResponse) {
	response = &DissociateCertificateWithGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DissociateCertificateWithGroup(request *DissociateCertificateWithGroupRequest) string {
	return c.DissociateCertificateWithGroupWithContext(context.Background(), request)
}

func (c *Client) DissociateCertificateWithGroupSend(request *DissociateCertificateWithGroupRequest) (*DissociateCertificateWithGroupResponse, error) {
	statusCode, msg, err := c.DissociateCertificateWithGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DissociateCertificateWithGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DissociateCertificateWithGroupWithContext(ctx context.Context, request *DissociateCertificateWithGroupRequest) string {
	if request == nil {
		request = NewDissociateCertificateWithGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDissociateCertificateWithGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DissociateCertificateWithGroupWithContextV2(ctx context.Context, request *DissociateCertificateWithGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDissociateCertificateWithGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDissociateCertificateWithGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateCertificateWithGroupRequest() (request *AssociateCertificateWithGroupRequest) {
	request = &AssociateCertificateWithGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AssociateCertificateWithGroup")
	return
}

func NewAssociateCertificateWithGroupResponse() (response *AssociateCertificateWithGroupResponse) {
	response = &AssociateCertificateWithGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateCertificateWithGroup(request *AssociateCertificateWithGroupRequest) string {
	return c.AssociateCertificateWithGroupWithContext(context.Background(), request)
}

func (c *Client) AssociateCertificateWithGroupSend(request *AssociateCertificateWithGroupRequest) (*AssociateCertificateWithGroupResponse, error) {
	statusCode, msg, err := c.AssociateCertificateWithGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssociateCertificateWithGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateCertificateWithGroupWithContext(ctx context.Context, request *AssociateCertificateWithGroupRequest) string {
	if request == nil {
		request = NewAssociateCertificateWithGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateCertificateWithGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateCertificateWithGroupWithContextV2(ctx context.Context, request *AssociateCertificateWithGroupRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateCertificateWithGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateCertificateWithGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeListenerCertGroupsRequest() (request *DescribeListenerCertGroupsRequest) {
	request = &DescribeListenerCertGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeListenerCertGroups")
	return
}

func NewDescribeListenerCertGroupsResponse() (response *DescribeListenerCertGroupsResponse) {
	response = &DescribeListenerCertGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeListenerCertGroups(request *DescribeListenerCertGroupsRequest) string {
	return c.DescribeListenerCertGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeListenerCertGroupsSend(request *DescribeListenerCertGroupsRequest) (*DescribeListenerCertGroupsResponse, error) {
	statusCode, msg, err := c.DescribeListenerCertGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeListenerCertGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeListenerCertGroupsWithContext(ctx context.Context, request *DescribeListenerCertGroupsRequest) string {
	if request == nil {
		request = NewDescribeListenerCertGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeListenerCertGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeListenerCertGroupsWithContextV2(ctx context.Context, request *DescribeListenerCertGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeListenerCertGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeListenerCertGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteListenerCertGroupRequest() (request *DeleteListenerCertGroupRequest) {
	request = &DeleteListenerCertGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteListenerCertGroup")
	return
}

func NewDeleteListenerCertGroupResponse() (response *DeleteListenerCertGroupResponse) {
	response = &DeleteListenerCertGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteListenerCertGroup(request *DeleteListenerCertGroupRequest) string {
	return c.DeleteListenerCertGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteListenerCertGroupSend(request *DeleteListenerCertGroupRequest) (*DeleteListenerCertGroupResponse, error) {
	statusCode, msg, err := c.DeleteListenerCertGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteListenerCertGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteListenerCertGroupWithContext(ctx context.Context, request *DeleteListenerCertGroupRequest) string {
	if request == nil {
		request = NewDeleteListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteListenerCertGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteListenerCertGroupWithContextV2(ctx context.Context, request *DeleteListenerCertGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteListenerCertGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateListenerCertGroupRequest() (request *CreateListenerCertGroupRequest) {
	request = &CreateListenerCertGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateListenerCertGroup")
	return
}

func NewCreateListenerCertGroupResponse() (response *CreateListenerCertGroupResponse) {
	response = &CreateListenerCertGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateListenerCertGroup(request *CreateListenerCertGroupRequest) string {
	return c.CreateListenerCertGroupWithContext(context.Background(), request)
}

func (c *Client) CreateListenerCertGroupSend(request *CreateListenerCertGroupRequest) (*CreateListenerCertGroupResponse, error) {
	statusCode, msg, err := c.CreateListenerCertGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateListenerCertGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateListenerCertGroupWithContext(ctx context.Context, request *CreateListenerCertGroupRequest) string {
	if request == nil {
		request = NewCreateListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateListenerCertGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateListenerCertGroupWithContextV2(ctx context.Context, request *CreateListenerCertGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateListenerCertGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddRulesRequest() (request *AddRulesRequest) {
	request = &AddRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AddRules")
	return
}

func NewAddRulesResponse() (response *AddRulesResponse) {
	response = &AddRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddRules(request *AddRulesRequest) string {
	return c.AddRulesWithContext(context.Background(), request)
}

func (c *Client) AddRulesSend(request *AddRulesRequest) (*AddRulesResponse, error) {
	statusCode, msg, err := c.AddRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddRulesWithContext(ctx context.Context, request *AddRulesRequest) string {
	if request == nil {
		request = NewAddRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddRulesWithContextV2(ctx context.Context, request *AddRulesRequest) (int, string, error) {
	if request == nil {
		request = NewAddRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
	request = &DeleteRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteRule")
	return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
	response = &DeleteRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRule(request *DeleteRuleRequest) string {
	return c.DeleteRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteRuleSend(request *DeleteRuleRequest) (*DeleteRuleResponse, error) {
	statusCode, msg, err := c.DeleteRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) string {
	if request == nil {
		request = NewDeleteRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRuleWithContextV2(ctx context.Context, request *DeleteRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddRuleRequest() (request *AddRuleRequest) {
	request = &AddRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AddRule")
	return
}

func NewAddRuleResponse() (response *AddRuleResponse) {
	response = &AddRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddRule(request *AddRuleRequest) string {
	return c.AddRuleWithContext(context.Background(), request)
}

func (c *Client) AddRuleSend(request *AddRuleRequest) (*AddRuleResponse, error) {
	statusCode, msg, err := c.AddRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddRuleWithContext(ctx context.Context, request *AddRuleRequest) string {
	if request == nil {
		request = NewAddRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddRuleWithContextV2(ctx context.Context, request *AddRuleRequest) (int, string, error) {
	if request == nil {
		request = NewAddRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyRuleGroupRequest() (request *ModifyRuleGroupRequest) {
	request = &ModifyRuleGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyRuleGroup")
	return
}

func NewModifyRuleGroupResponse() (response *ModifyRuleGroupResponse) {
	response = &ModifyRuleGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRuleGroup(request *ModifyRuleGroupRequest) string {
	return c.ModifyRuleGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyRuleGroupSend(request *ModifyRuleGroupRequest) (*ModifyRuleGroupResponse, error) {
	statusCode, msg, err := c.ModifyRuleGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyRuleGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyRuleGroupWithContext(ctx context.Context, request *ModifyRuleGroupRequest) string {
	if request == nil {
		request = NewModifyRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRuleGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyRuleGroupWithContextV2(ctx context.Context, request *ModifyRuleGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRuleGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRuleGroupsRequest() (request *DescribeRuleGroupsRequest) {
	request = &DescribeRuleGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeRuleGroups")
	return
}

func NewDescribeRuleGroupsResponse() (response *DescribeRuleGroupsResponse) {
	response = &DescribeRuleGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRuleGroups(request *DescribeRuleGroupsRequest) string {
	return c.DescribeRuleGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeRuleGroupsSend(request *DescribeRuleGroupsRequest) (*DescribeRuleGroupsResponse, error) {
	statusCode, msg, err := c.DescribeRuleGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRuleGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRuleGroupsWithContext(ctx context.Context, request *DescribeRuleGroupsRequest) string {
	if request == nil {
		request = NewDescribeRuleGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRuleGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRuleGroupsWithContextV2(ctx context.Context, request *DescribeRuleGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRuleGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRuleGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRuleGroupRequest() (request *DeleteRuleGroupRequest) {
	request = &DeleteRuleGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteRuleGroup")
	return
}

func NewDeleteRuleGroupResponse() (response *DeleteRuleGroupResponse) {
	response = &DeleteRuleGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRuleGroup(request *DeleteRuleGroupRequest) string {
	return c.DeleteRuleGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteRuleGroupSend(request *DeleteRuleGroupRequest) (*DeleteRuleGroupResponse, error) {
	statusCode, msg, err := c.DeleteRuleGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRuleGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRuleGroupWithContext(ctx context.Context, request *DeleteRuleGroupRequest) string {
	if request == nil {
		request = NewDeleteRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRuleGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRuleGroupWithContextV2(ctx context.Context, request *DeleteRuleGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRuleGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRuleGroupRequest() (request *CreateRuleGroupRequest) {
	request = &CreateRuleGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateRuleGroup")
	return
}

func NewCreateRuleGroupResponse() (response *CreateRuleGroupResponse) {
	response = &CreateRuleGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRuleGroup(request *CreateRuleGroupRequest) string {
	return c.CreateRuleGroupWithContext(context.Background(), request)
}

func (c *Client) CreateRuleGroupSend(request *CreateRuleGroupRequest) (*CreateRuleGroupResponse, error) {
	statusCode, msg, err := c.CreateRuleGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRuleGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRuleGroupWithContext(ctx context.Context, request *CreateRuleGroupRequest) string {
	if request == nil {
		request = NewCreateRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRuleGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRuleGroupWithContextV2(ctx context.Context, request *CreateRuleGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRuleGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetLBModificationProtectionRequest() (request *SetLBModificationProtectionRequest) {
	request = &SetLBModificationProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetLBModificationProtection")
	return
}

func NewSetLBModificationProtectionResponse() (response *SetLBModificationProtectionResponse) {
	response = &SetLBModificationProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetLBModificationProtection(request *SetLBModificationProtectionRequest) string {
	return c.SetLBModificationProtectionWithContext(context.Background(), request)
}

func (c *Client) SetLBModificationProtectionSend(request *SetLBModificationProtectionRequest) (*SetLBModificationProtectionResponse, error) {
	statusCode, msg, err := c.SetLBModificationProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetLBModificationProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetLBModificationProtectionWithContext(ctx context.Context, request *SetLBModificationProtectionRequest) string {
	if request == nil {
		request = NewSetLBModificationProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLBModificationProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetLBModificationProtectionWithContextV2(ctx context.Context, request *SetLBModificationProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewSetLBModificationProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLBModificationProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetLBDeleteProtectionRequest() (request *SetLBDeleteProtectionRequest) {
	request = &SetLBDeleteProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetLBDeleteProtection")
	return
}

func NewSetLBDeleteProtectionResponse() (response *SetLBDeleteProtectionResponse) {
	response = &SetLBDeleteProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetLBDeleteProtection(request *SetLBDeleteProtectionRequest) string {
	return c.SetLBDeleteProtectionWithContext(context.Background(), request)
}

func (c *Client) SetLBDeleteProtectionSend(request *SetLBDeleteProtectionRequest) (*SetLBDeleteProtectionResponse, error) {
	statusCode, msg, err := c.SetLBDeleteProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetLBDeleteProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetLBDeleteProtectionWithContext(ctx context.Context, request *SetLBDeleteProtectionRequest) string {
	if request == nil {
		request = NewSetLBDeleteProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLBDeleteProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetLBDeleteProtectionWithContextV2(ctx context.Context, request *SetLBDeleteProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewSetLBDeleteProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetLBDeleteProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


