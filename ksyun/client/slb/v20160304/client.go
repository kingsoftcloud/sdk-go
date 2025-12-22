package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-03-04"

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

	if msg == "" {
		return nil, nil
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
func NewDeleteListenersRequest() (request *DeleteListenersRequest) {
	request = &DeleteListenersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteListeners")
	return
}

func NewDeleteListenersResponse() (response *DeleteListenersResponse) {
	response = &DeleteListenersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteListeners(request *DeleteListenersRequest) string {
	return c.DeleteListenersWithContext(context.Background(), request)
}

func (c *Client) DeleteListenersSend(request *DeleteListenersRequest) (*DeleteListenersResponse, error) {
	statusCode, msg, err := c.DeleteListenersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteListenersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteListenersWithContext(ctx context.Context, request *DeleteListenersRequest) string {
	if request == nil {
		request = NewDeleteListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteListenersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteListenersWithContextV2(ctx context.Context, request *DeleteListenersRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteListenersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyListenersRequest() (request *ModifyListenersRequest) {
	request = &ModifyListenersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyListeners")
	return
}

func NewModifyListenersResponse() (response *ModifyListenersResponse) {
	response = &ModifyListenersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyListeners(request *ModifyListenersRequest) string {
	return c.ModifyListenersWithContext(context.Background(), request)
}

func (c *Client) ModifyListenersSend(request *ModifyListenersRequest) (*ModifyListenersResponse, error) {
	statusCode, msg, err := c.ModifyListenersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyListenersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyListenersWithContext(ctx context.Context, request *ModifyListenersRequest) string {
	if request == nil {
		request = NewModifyListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyListenersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyListenersWithContextV2(ctx context.Context, request *ModifyListenersRequest) (int, string, error) {
	if request == nil {
		request = NewModifyListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyListenersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateListenersRequest() (request *CreateListenersRequest) {
	request = &CreateListenersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateListeners")
	return
}

func NewCreateListenersResponse() (response *CreateListenersResponse) {
	response = &CreateListenersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateListeners(request *CreateListenersRequest) string {
	return c.CreateListenersWithContext(context.Background(), request)
}

func (c *Client) CreateListenersSend(request *CreateListenersRequest) (*CreateListenersResponse, error) {
	statusCode, msg, err := c.CreateListenersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateListenersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateListenersWithContext(ctx context.Context, request *CreateListenersRequest) string {
	if request == nil {
		request = NewCreateListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateListenersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateListenersWithContextV2(ctx context.Context, request *CreateListenersRequest) (int, string, error) {
	if request == nil {
		request = NewCreateListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateListenersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstancesWithListenerRequest() (request *ModifyInstancesWithListenerRequest) {
	request = &ModifyInstancesWithListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyInstancesWithListener")
	return
}

func NewModifyInstancesWithListenerResponse() (response *ModifyInstancesWithListenerResponse) {
	response = &ModifyInstancesWithListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstancesWithListener(request *ModifyInstancesWithListenerRequest) string {
	return c.ModifyInstancesWithListenerWithContext(context.Background(), request)
}

func (c *Client) ModifyInstancesWithListenerSend(request *ModifyInstancesWithListenerRequest) (*ModifyInstancesWithListenerResponse, error) {
	statusCode, msg, err := c.ModifyInstancesWithListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyInstancesWithListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstancesWithListenerWithContext(ctx context.Context, request *ModifyInstancesWithListenerRequest) string {
	if request == nil {
		request = NewModifyInstancesWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstancesWithListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstancesWithListenerWithContextV2(ctx context.Context, request *ModifyInstancesWithListenerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstancesWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstancesWithListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRegisterInstancesWithListenerRequest() (request *RegisterInstancesWithListenerRequest) {
	request = &RegisterInstancesWithListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "RegisterInstancesWithListener")
	return
}

func NewRegisterInstancesWithListenerResponse() (response *RegisterInstancesWithListenerResponse) {
	response = &RegisterInstancesWithListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RegisterInstancesWithListener(request *RegisterInstancesWithListenerRequest) string {
	return c.RegisterInstancesWithListenerWithContext(context.Background(), request)
}

func (c *Client) RegisterInstancesWithListenerSend(request *RegisterInstancesWithListenerRequest) (*RegisterInstancesWithListenerResponse, error) {
	statusCode, msg, err := c.RegisterInstancesWithListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RegisterInstancesWithListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RegisterInstancesWithListenerWithContext(ctx context.Context, request *RegisterInstancesWithListenerRequest) string {
	if request == nil {
		request = NewRegisterInstancesWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterInstancesWithListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RegisterInstancesWithListenerWithContextV2(ctx context.Context, request *RegisterInstancesWithListenerRequest) (int, string, error) {
	if request == nil {
		request = NewRegisterInstancesWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterInstancesWithListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeregisterInstancesFromListenerRequest() (request *DeregisterInstancesFromListenerRequest) {
	request = &DeregisterInstancesFromListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeregisterInstancesFromListener")
	return
}

func NewDeregisterInstancesFromListenerResponse() (response *DeregisterInstancesFromListenerResponse) {
	response = &DeregisterInstancesFromListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeregisterInstancesFromListener(request *DeregisterInstancesFromListenerRequest) string {
	return c.DeregisterInstancesFromListenerWithContext(context.Background(), request)
}

func (c *Client) DeregisterInstancesFromListenerSend(request *DeregisterInstancesFromListenerRequest) (*DeregisterInstancesFromListenerResponse, error) {
	statusCode, msg, err := c.DeregisterInstancesFromListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeregisterInstancesFromListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeregisterInstancesFromListenerWithContext(ctx context.Context, request *DeregisterInstancesFromListenerRequest) string {
	if request == nil {
		request = NewDeregisterInstancesFromListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeregisterInstancesFromListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeregisterInstancesFromListenerWithContextV2(ctx context.Context, request *DeregisterInstancesFromListenerRequest) (int, string, error) {
	if request == nil {
		request = NewDeregisterInstancesFromListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeregisterInstancesFromListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstancesWithListenerRequest() (request *DescribeInstancesWithListenerRequest) {
	request = &DescribeInstancesWithListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeInstancesWithListener")
	return
}

func NewDescribeInstancesWithListenerResponse() (response *DescribeInstancesWithListenerResponse) {
	response = &DescribeInstancesWithListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstancesWithListener(request *DescribeInstancesWithListenerRequest) string {
	return c.DescribeInstancesWithListenerWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesWithListenerSend(request *DescribeInstancesWithListenerRequest) (*DescribeInstancesWithListenerResponse, error) {
	statusCode, msg, err := c.DescribeInstancesWithListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstancesWithListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstancesWithListenerWithContext(ctx context.Context, request *DescribeInstancesWithListenerRequest) string {
	if request == nil {
		request = NewDescribeInstancesWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesWithListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstancesWithListenerWithContextV2(ctx context.Context, request *DescribeInstancesWithListenerRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstancesWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesWithListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyHealthCheckRequest() (request *ModifyHealthCheckRequest) {
	request = &ModifyHealthCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyHealthCheck")
	return
}

func NewModifyHealthCheckResponse() (response *ModifyHealthCheckResponse) {
	response = &ModifyHealthCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyHealthCheck(request *ModifyHealthCheckRequest) string {
	return c.ModifyHealthCheckWithContext(context.Background(), request)
}

func (c *Client) ModifyHealthCheckSend(request *ModifyHealthCheckRequest) (*ModifyHealthCheckResponse, error) {
	statusCode, msg, err := c.ModifyHealthCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyHealthCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyHealthCheckWithContext(ctx context.Context, request *ModifyHealthCheckRequest) string {
	if request == nil {
		request = NewModifyHealthCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyHealthCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyHealthCheckWithContextV2(ctx context.Context, request *ModifyHealthCheckRequest) (int, string, error) {
	if request == nil {
		request = NewModifyHealthCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyHealthCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteHealthCheckRequest() (request *DeleteHealthCheckRequest) {
	request = &DeleteHealthCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteHealthCheck")
	return
}

func NewDeleteHealthCheckResponse() (response *DeleteHealthCheckResponse) {
	response = &DeleteHealthCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteHealthCheck(request *DeleteHealthCheckRequest) string {
	return c.DeleteHealthCheckWithContext(context.Background(), request)
}

func (c *Client) DeleteHealthCheckSend(request *DeleteHealthCheckRequest) (*DeleteHealthCheckResponse, error) {
	statusCode, msg, err := c.DeleteHealthCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteHealthCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteHealthCheckWithContext(ctx context.Context, request *DeleteHealthCheckRequest) string {
	if request == nil {
		request = NewDeleteHealthCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHealthCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteHealthCheckWithContextV2(ctx context.Context, request *DeleteHealthCheckRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteHealthCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHealthCheckResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeHealthChecksRequest() (request *DescribeHealthChecksRequest) {
	request = &DescribeHealthChecksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeHealthChecks")
	return
}

func NewDescribeHealthChecksResponse() (response *DescribeHealthChecksResponse) {
	response = &DescribeHealthChecksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeHealthChecks(request *DescribeHealthChecksRequest) string {
	return c.DescribeHealthChecksWithContext(context.Background(), request)
}

func (c *Client) DescribeHealthChecksSend(request *DescribeHealthChecksRequest) (*DescribeHealthChecksResponse, error) {
	statusCode, msg, err := c.DescribeHealthChecksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeHealthChecksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeHealthChecksWithContext(ctx context.Context, request *DescribeHealthChecksRequest) string {
	if request == nil {
		request = NewDescribeHealthChecksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHealthChecksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeHealthChecksWithContextV2(ctx context.Context, request *DescribeHealthChecksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeHealthChecksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHealthChecksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewConfigureHealthCheckRequest() (request *ConfigureHealthCheckRequest) {
	request = &ConfigureHealthCheckRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ConfigureHealthCheck")
	return
}

func NewConfigureHealthCheckResponse() (response *ConfigureHealthCheckResponse) {
	response = &ConfigureHealthCheckResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ConfigureHealthCheck(request *ConfigureHealthCheckRequest) string {
	return c.ConfigureHealthCheckWithContext(context.Background(), request)
}

func (c *Client) ConfigureHealthCheckSend(request *ConfigureHealthCheckRequest) (*ConfigureHealthCheckResponse, error) {
	statusCode, msg, err := c.ConfigureHealthCheckWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ConfigureHealthCheckResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ConfigureHealthCheckWithContext(ctx context.Context, request *ConfigureHealthCheckRequest) string {
	if request == nil {
		request = NewConfigureHealthCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewConfigureHealthCheckResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ConfigureHealthCheckWithContextV2(ctx context.Context, request *ConfigureHealthCheckRequest) (int, string, error) {
	if request == nil {
		request = NewConfigureHealthCheckRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewConfigureHealthCheckResponse()
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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
func NewModifyLoadBalancerRequest() (request *ModifyLoadBalancerRequest) {
	request = &ModifyLoadBalancerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyLoadBalancer")
	return
}

func NewModifyLoadBalancerResponse() (response *ModifyLoadBalancerResponse) {
	response = &ModifyLoadBalancerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyLoadBalancer(request *ModifyLoadBalancerRequest) string {
	return c.ModifyLoadBalancerWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancerSend(request *ModifyLoadBalancerRequest) (*ModifyLoadBalancerResponse, error) {
	statusCode, msg, err := c.ModifyLoadBalancerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyLoadBalancerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyLoadBalancerWithContext(ctx context.Context, request *ModifyLoadBalancerRequest) string {
	if request == nil {
		request = NewModifyLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyLoadBalancerWithContextV2(ctx context.Context, request *ModifyLoadBalancerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancerResponse()
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

	if msg == "" {
		return nil, nil
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
func NewCreateHostHeaderRequest() (request *CreateHostHeaderRequest) {
	request = &CreateHostHeaderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateHostHeader")
	return
}

func NewCreateHostHeaderResponse() (response *CreateHostHeaderResponse) {
	response = &CreateHostHeaderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateHostHeader(request *CreateHostHeaderRequest) string {
	return c.CreateHostHeaderWithContext(context.Background(), request)
}

func (c *Client) CreateHostHeaderSend(request *CreateHostHeaderRequest) (*CreateHostHeaderResponse, error) {
	statusCode, msg, err := c.CreateHostHeaderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateHostHeaderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateHostHeaderWithContext(ctx context.Context, request *CreateHostHeaderRequest) string {
	if request == nil {
		request = NewCreateHostHeaderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateHostHeaderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateHostHeaderWithContextV2(ctx context.Context, request *CreateHostHeaderRequest) (int, string, error) {
	if request == nil {
		request = NewCreateHostHeaderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateHostHeaderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyHostHeaderRequest() (request *ModifyHostHeaderRequest) {
	request = &ModifyHostHeaderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyHostHeader")
	return
}

func NewModifyHostHeaderResponse() (response *ModifyHostHeaderResponse) {
	response = &ModifyHostHeaderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyHostHeader(request *ModifyHostHeaderRequest) string {
	return c.ModifyHostHeaderWithContext(context.Background(), request)
}

func (c *Client) ModifyHostHeaderSend(request *ModifyHostHeaderRequest) (*ModifyHostHeaderResponse, error) {
	statusCode, msg, err := c.ModifyHostHeaderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyHostHeaderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyHostHeaderWithContext(ctx context.Context, request *ModifyHostHeaderRequest) string {
	if request == nil {
		request = NewModifyHostHeaderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyHostHeaderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyHostHeaderWithContextV2(ctx context.Context, request *ModifyHostHeaderRequest) (int, string, error) {
	if request == nil {
		request = NewModifyHostHeaderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyHostHeaderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteHostHeaderRequest() (request *DeleteHostHeaderRequest) {
	request = &DeleteHostHeaderRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteHostHeader")
	return
}

func NewDeleteHostHeaderResponse() (response *DeleteHostHeaderResponse) {
	response = &DeleteHostHeaderResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteHostHeader(request *DeleteHostHeaderRequest) string {
	return c.DeleteHostHeaderWithContext(context.Background(), request)
}

func (c *Client) DeleteHostHeaderSend(request *DeleteHostHeaderRequest) (*DeleteHostHeaderResponse, error) {
	statusCode, msg, err := c.DeleteHostHeaderWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteHostHeaderResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteHostHeaderWithContext(ctx context.Context, request *DeleteHostHeaderRequest) string {
	if request == nil {
		request = NewDeleteHostHeaderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHostHeaderResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteHostHeaderWithContextV2(ctx context.Context, request *DeleteHostHeaderRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteHostHeaderRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteHostHeaderResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeHostHeadersRequest() (request *DescribeHostHeadersRequest) {
	request = &DescribeHostHeadersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeHostHeaders")
	return
}

func NewDescribeHostHeadersResponse() (response *DescribeHostHeadersResponse) {
	response = &DescribeHostHeadersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeHostHeaders(request *DescribeHostHeadersRequest) string {
	return c.DescribeHostHeadersWithContext(context.Background(), request)
}

func (c *Client) DescribeHostHeadersSend(request *DescribeHostHeadersRequest) (*DescribeHostHeadersResponse, error) {
	statusCode, msg, err := c.DescribeHostHeadersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeHostHeadersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeHostHeadersWithContext(ctx context.Context, request *DescribeHostHeadersRequest) string {
	if request == nil {
		request = NewDescribeHostHeadersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHostHeadersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeHostHeadersWithContextV2(ctx context.Context, request *DescribeHostHeadersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeHostHeadersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHostHeadersResponse()
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

	if msg == "" {
		return nil, nil
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
func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
	request = &DescribeRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeRules")
	return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
	response = &DescribeRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRules(request *DescribeRulesRequest) string {
	return c.DescribeRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeRulesSend(request *DescribeRulesRequest) (*DescribeRulesResponse, error) {
	statusCode, msg, err := c.DescribeRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) string {
	if request == nil {
		request = NewDescribeRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRulesWithContextV2(ctx context.Context, request *DescribeRulesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRulesResponse()
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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
func NewCreateLoadBalancerAclRequest() (request *CreateLoadBalancerAclRequest) {
	request = &CreateLoadBalancerAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateLoadBalancerAcl")
	return
}

func NewCreateLoadBalancerAclResponse() (response *CreateLoadBalancerAclResponse) {
	response = &CreateLoadBalancerAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLoadBalancerAcl(request *CreateLoadBalancerAclRequest) string {
	return c.CreateLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) CreateLoadBalancerAclSend(request *CreateLoadBalancerAclRequest) (*CreateLoadBalancerAclResponse, error) {
	statusCode, msg, err := c.CreateLoadBalancerAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateLoadBalancerAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateLoadBalancerAclWithContext(ctx context.Context, request *CreateLoadBalancerAclRequest) string {
	if request == nil {
		request = NewCreateLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLoadBalancerAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateLoadBalancerAclWithContextV2(ctx context.Context, request *CreateLoadBalancerAclRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLoadBalancerAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteLoadBalancerAclRequest() (request *DeleteLoadBalancerAclRequest) {
	request = &DeleteLoadBalancerAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteLoadBalancerAcl")
	return
}

func NewDeleteLoadBalancerAclResponse() (response *DeleteLoadBalancerAclResponse) {
	response = &DeleteLoadBalancerAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLoadBalancerAcl(request *DeleteLoadBalancerAclRequest) string {
	return c.DeleteLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) DeleteLoadBalancerAclSend(request *DeleteLoadBalancerAclRequest) (*DeleteLoadBalancerAclResponse, error) {
	statusCode, msg, err := c.DeleteLoadBalancerAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteLoadBalancerAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteLoadBalancerAclWithContext(ctx context.Context, request *DeleteLoadBalancerAclRequest) string {
	if request == nil {
		request = NewDeleteLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLoadBalancerAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteLoadBalancerAclWithContextV2(ctx context.Context, request *DeleteLoadBalancerAclRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLoadBalancerAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyLoadBalancerAclRequest() (request *ModifyLoadBalancerAclRequest) {
	request = &ModifyLoadBalancerAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyLoadBalancerAcl")
	return
}

func NewModifyLoadBalancerAclResponse() (response *ModifyLoadBalancerAclResponse) {
	response = &ModifyLoadBalancerAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyLoadBalancerAcl(request *ModifyLoadBalancerAclRequest) string {
	return c.ModifyLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancerAclSend(request *ModifyLoadBalancerAclRequest) (*ModifyLoadBalancerAclResponse, error) {
	statusCode, msg, err := c.ModifyLoadBalancerAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyLoadBalancerAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyLoadBalancerAclWithContext(ctx context.Context, request *ModifyLoadBalancerAclRequest) string {
	if request == nil {
		request = NewModifyLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancerAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyLoadBalancerAclWithContextV2(ctx context.Context, request *ModifyLoadBalancerAclRequest) (int, string, error) {
	if request == nil {
		request = NewModifyLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancerAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateLoadBalancerAclEntryRequest() (request *CreateLoadBalancerAclEntryRequest) {
	request = &CreateLoadBalancerAclEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateLoadBalancerAclEntry")
	return
}

func NewCreateLoadBalancerAclEntryResponse() (response *CreateLoadBalancerAclEntryResponse) {
	response = &CreateLoadBalancerAclEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLoadBalancerAclEntry(request *CreateLoadBalancerAclEntryRequest) string {
	return c.CreateLoadBalancerAclEntryWithContext(context.Background(), request)
}

func (c *Client) CreateLoadBalancerAclEntrySend(request *CreateLoadBalancerAclEntryRequest) (*CreateLoadBalancerAclEntryResponse, error) {
	statusCode, msg, err := c.CreateLoadBalancerAclEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateLoadBalancerAclEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateLoadBalancerAclEntryWithContext(ctx context.Context, request *CreateLoadBalancerAclEntryRequest) string {
	if request == nil {
		request = NewCreateLoadBalancerAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLoadBalancerAclEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateLoadBalancerAclEntryWithContextV2(ctx context.Context, request *CreateLoadBalancerAclEntryRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLoadBalancerAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLoadBalancerAclEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteLoadBalancerAclEntryRequest() (request *DeleteLoadBalancerAclEntryRequest) {
	request = &DeleteLoadBalancerAclEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteLoadBalancerAclEntry")
	return
}

func NewDeleteLoadBalancerAclEntryResponse() (response *DeleteLoadBalancerAclEntryResponse) {
	response = &DeleteLoadBalancerAclEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLoadBalancerAclEntry(request *DeleteLoadBalancerAclEntryRequest) string {
	return c.DeleteLoadBalancerAclEntryWithContext(context.Background(), request)
}

func (c *Client) DeleteLoadBalancerAclEntrySend(request *DeleteLoadBalancerAclEntryRequest) (*DeleteLoadBalancerAclEntryResponse, error) {
	statusCode, msg, err := c.DeleteLoadBalancerAclEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteLoadBalancerAclEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteLoadBalancerAclEntryWithContext(ctx context.Context, request *DeleteLoadBalancerAclEntryRequest) string {
	if request == nil {
		request = NewDeleteLoadBalancerAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLoadBalancerAclEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteLoadBalancerAclEntryWithContextV2(ctx context.Context, request *DeleteLoadBalancerAclEntryRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLoadBalancerAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLoadBalancerAclEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateLoadBalancerAclRequest() (request *AssociateLoadBalancerAclRequest) {
	request = &AssociateLoadBalancerAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AssociateLoadBalancerAcl")
	return
}

func NewAssociateLoadBalancerAclResponse() (response *AssociateLoadBalancerAclResponse) {
	response = &AssociateLoadBalancerAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateLoadBalancerAcl(request *AssociateLoadBalancerAclRequest) string {
	return c.AssociateLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) AssociateLoadBalancerAclSend(request *AssociateLoadBalancerAclRequest) (*AssociateLoadBalancerAclResponse, error) {
	statusCode, msg, err := c.AssociateLoadBalancerAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateLoadBalancerAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateLoadBalancerAclWithContext(ctx context.Context, request *AssociateLoadBalancerAclRequest) string {
	if request == nil {
		request = NewAssociateLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateLoadBalancerAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateLoadBalancerAclWithContextV2(ctx context.Context, request *AssociateLoadBalancerAclRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateLoadBalancerAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateLoadBalancerAclRequest() (request *DisassociateLoadBalancerAclRequest) {
	request = &DisassociateLoadBalancerAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DisassociateLoadBalancerAcl")
	return
}

func NewDisassociateLoadBalancerAclResponse() (response *DisassociateLoadBalancerAclResponse) {
	response = &DisassociateLoadBalancerAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateLoadBalancerAcl(request *DisassociateLoadBalancerAclRequest) string {
	return c.DisassociateLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) DisassociateLoadBalancerAclSend(request *DisassociateLoadBalancerAclRequest) (*DisassociateLoadBalancerAclResponse, error) {
	statusCode, msg, err := c.DisassociateLoadBalancerAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisassociateLoadBalancerAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateLoadBalancerAclWithContext(ctx context.Context, request *DisassociateLoadBalancerAclRequest) string {
	if request == nil {
		request = NewDisassociateLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateLoadBalancerAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateLoadBalancerAclWithContextV2(ctx context.Context, request *DisassociateLoadBalancerAclRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateLoadBalancerAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateLoadBalancerAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeLoadBalancerAclsRequest() (request *DescribeLoadBalancerAclsRequest) {
	request = &DescribeLoadBalancerAclsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeLoadBalancerAcls")
	return
}

func NewDescribeLoadBalancerAclsResponse() (response *DescribeLoadBalancerAclsResponse) {
	response = &DescribeLoadBalancerAclsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLoadBalancerAcls(request *DescribeLoadBalancerAclsRequest) string {
	return c.DescribeLoadBalancerAclsWithContext(context.Background(), request)
}

func (c *Client) DescribeLoadBalancerAclsSend(request *DescribeLoadBalancerAclsRequest) (*DescribeLoadBalancerAclsResponse, error) {
	statusCode, msg, err := c.DescribeLoadBalancerAclsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeLoadBalancerAclsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeLoadBalancerAclsWithContext(ctx context.Context, request *DescribeLoadBalancerAclsRequest) string {
	if request == nil {
		request = NewDescribeLoadBalancerAclsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLoadBalancerAclsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeLoadBalancerAclsWithContextV2(ctx context.Context, request *DescribeLoadBalancerAclsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeLoadBalancerAclsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLoadBalancerAclsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSlbRuleRequest() (request *CreateSlbRuleRequest) {
	request = &CreateSlbRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateSlbRule")
	return
}

func NewCreateSlbRuleResponse() (response *CreateSlbRuleResponse) {
	response = &CreateSlbRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSlbRule(request *CreateSlbRuleRequest) string {
	return c.CreateSlbRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSlbRuleSend(request *CreateSlbRuleRequest) (*CreateSlbRuleResponse, error) {
	statusCode, msg, err := c.CreateSlbRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateSlbRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSlbRuleWithContext(ctx context.Context, request *CreateSlbRuleRequest) string {
	if request == nil {
		request = NewCreateSlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSlbRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSlbRuleWithContextV2(ctx context.Context, request *CreateSlbRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSlbRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySlbRuleRequest() (request *ModifySlbRuleRequest) {
	request = &ModifySlbRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifySlbRule")
	return
}

func NewModifySlbRuleResponse() (response *ModifySlbRuleResponse) {
	response = &ModifySlbRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySlbRule(request *ModifySlbRuleRequest) string {
	return c.ModifySlbRuleWithContext(context.Background(), request)
}

func (c *Client) ModifySlbRuleSend(request *ModifySlbRuleRequest) (*ModifySlbRuleResponse, error) {
	statusCode, msg, err := c.ModifySlbRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifySlbRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySlbRuleWithContext(ctx context.Context, request *ModifySlbRuleRequest) string {
	if request == nil {
		request = NewModifySlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySlbRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySlbRuleWithContextV2(ctx context.Context, request *ModifySlbRuleRequest) (int, string, error) {
	if request == nil {
		request = NewModifySlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySlbRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreatePrivateLinkServerRequest() (request *CreatePrivateLinkServerRequest) {
	request = &CreatePrivateLinkServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreatePrivateLinkServer")
	return
}

func NewCreatePrivateLinkServerResponse() (response *CreatePrivateLinkServerResponse) {
	response = &CreatePrivateLinkServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePrivateLinkServer(request *CreatePrivateLinkServerRequest) string {
	return c.CreatePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) CreatePrivateLinkServerSend(request *CreatePrivateLinkServerRequest) (*CreatePrivateLinkServerResponse, error) {
	statusCode, msg, err := c.CreatePrivateLinkServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreatePrivateLinkServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreatePrivateLinkServerWithContext(ctx context.Context, request *CreatePrivateLinkServerRequest) string {
	if request == nil {
		request = NewCreatePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrivateLinkServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreatePrivateLinkServerWithContextV2(ctx context.Context, request *CreatePrivateLinkServerRequest) (int, string, error) {
	if request == nil {
		request = NewCreatePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrivateLinkServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePrivateLinkServerRequest() (request *DescribePrivateLinkServerRequest) {
	request = &DescribePrivateLinkServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribePrivateLinkServer")
	return
}

func NewDescribePrivateLinkServerResponse() (response *DescribePrivateLinkServerResponse) {
	response = &DescribePrivateLinkServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrivateLinkServer(request *DescribePrivateLinkServerRequest) string {
	return c.DescribePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) DescribePrivateLinkServerSend(request *DescribePrivateLinkServerRequest) (*DescribePrivateLinkServerResponse, error) {
	statusCode, msg, err := c.DescribePrivateLinkServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribePrivateLinkServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePrivateLinkServerWithContext(ctx context.Context, request *DescribePrivateLinkServerRequest) string {
	if request == nil {
		request = NewDescribePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateLinkServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePrivateLinkServerWithContextV2(ctx context.Context, request *DescribePrivateLinkServerRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateLinkServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePrivateLinkServerRequest() (request *DeletePrivateLinkServerRequest) {
	request = &DeletePrivateLinkServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeletePrivateLinkServer")
	return
}

func NewDeletePrivateLinkServerResponse() (response *DeletePrivateLinkServerResponse) {
	response = &DeletePrivateLinkServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePrivateLinkServer(request *DeletePrivateLinkServerRequest) string {
	return c.DeletePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) DeletePrivateLinkServerSend(request *DeletePrivateLinkServerRequest) (*DeletePrivateLinkServerResponse, error) {
	statusCode, msg, err := c.DeletePrivateLinkServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeletePrivateLinkServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePrivateLinkServerWithContext(ctx context.Context, request *DeletePrivateLinkServerRequest) string {
	if request == nil {
		request = NewDeletePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateLinkServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePrivateLinkServerWithContextV2(ctx context.Context, request *DeletePrivateLinkServerRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateLinkServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyPrivateLinkServerRequest() (request *ModifyPrivateLinkServerRequest) {
	request = &ModifyPrivateLinkServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyPrivateLinkServer")
	return
}

func NewModifyPrivateLinkServerResponse() (response *ModifyPrivateLinkServerResponse) {
	response = &ModifyPrivateLinkServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPrivateLinkServer(request *ModifyPrivateLinkServerRequest) string {
	return c.ModifyPrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) ModifyPrivateLinkServerSend(request *ModifyPrivateLinkServerRequest) (*ModifyPrivateLinkServerResponse, error) {
	statusCode, msg, err := c.ModifyPrivateLinkServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyPrivateLinkServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyPrivateLinkServerWithContext(ctx context.Context, request *ModifyPrivateLinkServerRequest) string {
	if request == nil {
		request = NewModifyPrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPrivateLinkServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyPrivateLinkServerWithContextV2(ctx context.Context, request *ModifyPrivateLinkServerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyPrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPrivateLinkServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociatePrivateLinkServerRequest() (request *AssociatePrivateLinkServerRequest) {
	request = &AssociatePrivateLinkServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AssociatePrivateLinkServer")
	return
}

func NewAssociatePrivateLinkServerResponse() (response *AssociatePrivateLinkServerResponse) {
	response = &AssociatePrivateLinkServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociatePrivateLinkServer(request *AssociatePrivateLinkServerRequest) string {
	return c.AssociatePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) AssociatePrivateLinkServerSend(request *AssociatePrivateLinkServerRequest) (*AssociatePrivateLinkServerResponse, error) {
	statusCode, msg, err := c.AssociatePrivateLinkServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociatePrivateLinkServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociatePrivateLinkServerWithContext(ctx context.Context, request *AssociatePrivateLinkServerRequest) string {
	if request == nil {
		request = NewAssociatePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociatePrivateLinkServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociatePrivateLinkServerWithContextV2(ctx context.Context, request *AssociatePrivateLinkServerRequest) (int, string, error) {
	if request == nil {
		request = NewAssociatePrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociatePrivateLinkServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePrivateLinkRequest() (request *DescribePrivateLinkRequest) {
	request = &DescribePrivateLinkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribePrivateLink")
	return
}

func NewDescribePrivateLinkResponse() (response *DescribePrivateLinkResponse) {
	response = &DescribePrivateLinkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrivateLink(request *DescribePrivateLinkRequest) string {
	return c.DescribePrivateLinkWithContext(context.Background(), request)
}

func (c *Client) DescribePrivateLinkSend(request *DescribePrivateLinkRequest) (*DescribePrivateLinkResponse, error) {
	statusCode, msg, err := c.DescribePrivateLinkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribePrivateLinkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePrivateLinkWithContext(ctx context.Context, request *DescribePrivateLinkRequest) string {
	if request == nil {
		request = NewDescribePrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateLinkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePrivateLinkWithContextV2(ctx context.Context, request *DescribePrivateLinkRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateLinkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePrivateLinkRequest() (request *DeletePrivateLinkRequest) {
	request = &DeletePrivateLinkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeletePrivateLink")
	return
}

func NewDeletePrivateLinkResponse() (response *DeletePrivateLinkResponse) {
	response = &DeletePrivateLinkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePrivateLink(request *DeletePrivateLinkRequest) string {
	return c.DeletePrivateLinkWithContext(context.Background(), request)
}

func (c *Client) DeletePrivateLinkSend(request *DeletePrivateLinkRequest) (*DeletePrivateLinkResponse, error) {
	statusCode, msg, err := c.DeletePrivateLinkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeletePrivateLinkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePrivateLinkWithContext(ctx context.Context, request *DeletePrivateLinkRequest) string {
	if request == nil {
		request = NewDeletePrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateLinkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePrivateLinkWithContextV2(ctx context.Context, request *DeletePrivateLinkRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateLinkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyLoadBalancerAclEntryRequest() (request *ModifyLoadBalancerAclEntryRequest) {
	request = &ModifyLoadBalancerAclEntryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyLoadBalancerAclEntry")
	return
}

func NewModifyLoadBalancerAclEntryResponse() (response *ModifyLoadBalancerAclEntryResponse) {
	response = &ModifyLoadBalancerAclEntryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyLoadBalancerAclEntry(request *ModifyLoadBalancerAclEntryRequest) string {
	return c.ModifyLoadBalancerAclEntryWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancerAclEntrySend(request *ModifyLoadBalancerAclEntryRequest) (*ModifyLoadBalancerAclEntryResponse, error) {
	statusCode, msg, err := c.ModifyLoadBalancerAclEntryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyLoadBalancerAclEntryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyLoadBalancerAclEntryWithContext(ctx context.Context, request *ModifyLoadBalancerAclEntryRequest) string {
	if request == nil {
		request = NewModifyLoadBalancerAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancerAclEntryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyLoadBalancerAclEntryWithContextV2(ctx context.Context, request *ModifyLoadBalancerAclEntryRequest) (int, string, error) {
	if request == nil {
		request = NewModifyLoadBalancerAclEntryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancerAclEntryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAcceptPrivateLinkRequest() (request *AcceptPrivateLinkRequest) {
	request = &AcceptPrivateLinkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AcceptPrivateLink")
	return
}

func NewAcceptPrivateLinkResponse() (response *AcceptPrivateLinkResponse) {
	response = &AcceptPrivateLinkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AcceptPrivateLink(request *AcceptPrivateLinkRequest) string {
	return c.AcceptPrivateLinkWithContext(context.Background(), request)
}

func (c *Client) AcceptPrivateLinkSend(request *AcceptPrivateLinkRequest) (*AcceptPrivateLinkResponse, error) {
	statusCode, msg, err := c.AcceptPrivateLinkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AcceptPrivateLinkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AcceptPrivateLinkWithContext(ctx context.Context, request *AcceptPrivateLinkRequest) string {
	if request == nil {
		request = NewAcceptPrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAcceptPrivateLinkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AcceptPrivateLinkWithContextV2(ctx context.Context, request *AcceptPrivateLinkRequest) (int, string, error) {
	if request == nil {
		request = NewAcceptPrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAcceptPrivateLinkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRejectPrivateLinkRequest() (request *RejectPrivateLinkRequest) {
	request = &RejectPrivateLinkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "RejectPrivateLink")
	return
}

func NewRejectPrivateLinkResponse() (response *RejectPrivateLinkResponse) {
	response = &RejectPrivateLinkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RejectPrivateLink(request *RejectPrivateLinkRequest) string {
	return c.RejectPrivateLinkWithContext(context.Background(), request)
}

func (c *Client) RejectPrivateLinkSend(request *RejectPrivateLinkRequest) (*RejectPrivateLinkResponse, error) {
	statusCode, msg, err := c.RejectPrivateLinkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RejectPrivateLinkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RejectPrivateLinkWithContext(ctx context.Context, request *RejectPrivateLinkRequest) string {
	if request == nil {
		request = NewRejectPrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRejectPrivateLinkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RejectPrivateLinkWithContextV2(ctx context.Context, request *RejectPrivateLinkRequest) (int, string, error) {
	if request == nil {
		request = NewRejectPrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRejectPrivateLinkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListPrivateLinkServerRequest() (request *ListPrivateLinkServerRequest) {
	request = &ListPrivateLinkServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ListPrivateLinkServer")
	return
}

func NewListPrivateLinkServerResponse() (response *ListPrivateLinkServerResponse) {
	response = &ListPrivateLinkServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListPrivateLinkServer(request *ListPrivateLinkServerRequest) string {
	return c.ListPrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) ListPrivateLinkServerSend(request *ListPrivateLinkServerRequest) (*ListPrivateLinkServerResponse, error) {
	statusCode, msg, err := c.ListPrivateLinkServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListPrivateLinkServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListPrivateLinkServerWithContext(ctx context.Context, request *ListPrivateLinkServerRequest) string {
	if request == nil {
		request = NewListPrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListPrivateLinkServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListPrivateLinkServerWithContextV2(ctx context.Context, request *ListPrivateLinkServerRequest) (int, string, error) {
	if request == nil {
		request = NewListPrivateLinkServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListPrivateLinkServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRemovePrivateLinkRequest() (request *RemovePrivateLinkRequest) {
	request = &RemovePrivateLinkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "RemovePrivateLink")
	return
}

func NewRemovePrivateLinkResponse() (response *RemovePrivateLinkResponse) {
	response = &RemovePrivateLinkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemovePrivateLink(request *RemovePrivateLinkRequest) string {
	return c.RemovePrivateLinkWithContext(context.Background(), request)
}

func (c *Client) RemovePrivateLinkSend(request *RemovePrivateLinkRequest) (*RemovePrivateLinkResponse, error) {
	statusCode, msg, err := c.RemovePrivateLinkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RemovePrivateLinkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RemovePrivateLinkWithContext(ctx context.Context, request *RemovePrivateLinkRequest) string {
	if request == nil {
		request = NewRemovePrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemovePrivateLinkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RemovePrivateLinkWithContextV2(ctx context.Context, request *RemovePrivateLinkRequest) (int, string, error) {
	if request == nil {
		request = NewRemovePrivateLinkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemovePrivateLinkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAlbRequest() (request *CreateAlbRequest) {
	request = &CreateAlbRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateAlb")
	return
}

func NewCreateAlbResponse() (response *CreateAlbResponse) {
	response = &CreateAlbResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlb(request *CreateAlbRequest) string {
	return c.CreateAlbWithContext(context.Background(), request)
}

func (c *Client) CreateAlbSend(request *CreateAlbRequest) (*CreateAlbResponse, error) {
	statusCode, msg, err := c.CreateAlbWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAlbResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAlbWithContext(ctx context.Context, request *CreateAlbRequest) string {
	if request == nil {
		request = NewCreateAlbRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAlbWithContextV2(ctx context.Context, request *CreateAlbRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAlbRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlbRequest() (request *DeleteAlbRequest) {
	request = &DeleteAlbRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteAlb")
	return
}

func NewDeleteAlbResponse() (response *DeleteAlbResponse) {
	response = &DeleteAlbResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlb(request *DeleteAlbRequest) string {
	return c.DeleteAlbWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbSend(request *DeleteAlbRequest) (*DeleteAlbResponse, error) {
	statusCode, msg, err := c.DeleteAlbWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlbResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlbWithContext(ctx context.Context, request *DeleteAlbRequest) string {
	if request == nil {
		request = NewDeleteAlbRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlbWithContextV2(ctx context.Context, request *DeleteAlbRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlbRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetAlbNameRequest() (request *SetAlbNameRequest) {
	request = &SetAlbNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetAlbName")
	return
}

func NewSetAlbNameResponse() (response *SetAlbNameResponse) {
	response = &SetAlbNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetAlbName(request *SetAlbNameRequest) string {
	return c.SetAlbNameWithContext(context.Background(), request)
}

func (c *Client) SetAlbNameSend(request *SetAlbNameRequest) (*SetAlbNameResponse, error) {
	statusCode, msg, err := c.SetAlbNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetAlbNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetAlbNameWithContext(ctx context.Context, request *SetAlbNameRequest) string {
	if request == nil {
		request = NewSetAlbNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetAlbNameWithContextV2(ctx context.Context, request *SetAlbNameRequest) (int, string, error) {
	if request == nil {
		request = NewSetAlbNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetAlbStatusRequest() (request *SetAlbStatusRequest) {
	request = &SetAlbStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetAlbStatus")
	return
}

func NewSetAlbStatusResponse() (response *SetAlbStatusResponse) {
	response = &SetAlbStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetAlbStatus(request *SetAlbStatusRequest) string {
	return c.SetAlbStatusWithContext(context.Background(), request)
}

func (c *Client) SetAlbStatusSend(request *SetAlbStatusRequest) (*SetAlbStatusResponse, error) {
	statusCode, msg, err := c.SetAlbStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetAlbStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetAlbStatusWithContext(ctx context.Context, request *SetAlbStatusRequest) string {
	if request == nil {
		request = NewSetAlbStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetAlbStatusWithContextV2(ctx context.Context, request *SetAlbStatusRequest) (int, string, error) {
	if request == nil {
		request = NewSetAlbStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlbsRequest() (request *DescribeAlbsRequest) {
	request = &DescribeAlbsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeAlbs")
	return
}

func NewDescribeAlbsResponse() (response *DescribeAlbsResponse) {
	response = &DescribeAlbsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbs(request *DescribeAlbsRequest) string {
	return c.DescribeAlbsWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbsSend(request *DescribeAlbsRequest) (*DescribeAlbsResponse, error) {
	statusCode, msg, err := c.DescribeAlbsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAlbsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlbsWithContext(ctx context.Context, request *DescribeAlbsRequest) string {
	if request == nil {
		request = NewDescribeAlbsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlbsWithContextV2(ctx context.Context, request *DescribeAlbsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlbsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAlbListenerRequest() (request *CreateAlbListenerRequest) {
	request = &CreateAlbListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateAlbListener")
	return
}

func NewCreateAlbListenerResponse() (response *CreateAlbListenerResponse) {
	response = &CreateAlbListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlbListener(request *CreateAlbListenerRequest) string {
	return c.CreateAlbListenerWithContext(context.Background(), request)
}

func (c *Client) CreateAlbListenerSend(request *CreateAlbListenerRequest) (*CreateAlbListenerResponse, error) {
	statusCode, msg, err := c.CreateAlbListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAlbListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAlbListenerWithContext(ctx context.Context, request *CreateAlbListenerRequest) string {
	if request == nil {
		request = NewCreateAlbListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateAlbListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAlbListenerWithContextV2(ctx context.Context, request *CreateAlbListenerRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAlbListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateAlbListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAlbListenerRequest() (request *ModifyAlbListenerRequest) {
	request = &ModifyAlbListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyAlbListener")
	return
}

func NewModifyAlbListenerResponse() (response *ModifyAlbListenerResponse) {
	response = &ModifyAlbListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAlbListener(request *ModifyAlbListenerRequest) string {
	return c.ModifyAlbListenerWithContext(context.Background(), request)
}

func (c *Client) ModifyAlbListenerSend(request *ModifyAlbListenerRequest) (*ModifyAlbListenerResponse, error) {
	statusCode, msg, err := c.ModifyAlbListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAlbListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAlbListenerWithContext(ctx context.Context, request *ModifyAlbListenerRequest) string {
	if request == nil {
		request = NewModifyAlbListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAlbListenerWithContextV2(ctx context.Context, request *ModifyAlbListenerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAlbListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlbListenerRequest() (request *DeleteAlbListenerRequest) {
	request = &DeleteAlbListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteAlbListener")
	return
}

func NewDeleteAlbListenerResponse() (response *DeleteAlbListenerResponse) {
	response = &DeleteAlbListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlbListener(request *DeleteAlbListenerRequest) string {
	return c.DeleteAlbListenerWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbListenerSend(request *DeleteAlbListenerRequest) (*DeleteAlbListenerResponse, error) {
	statusCode, msg, err := c.DeleteAlbListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlbListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlbListenerWithContext(ctx context.Context, request *DeleteAlbListenerRequest) string {
	if request == nil {
		request = NewDeleteAlbListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlbListenerWithContextV2(ctx context.Context, request *DeleteAlbListenerRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlbListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlbListenersRequest() (request *DescribeAlbListenersRequest) {
	request = &DescribeAlbListenersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeAlbListeners")
	return
}

func NewDescribeAlbListenersResponse() (response *DescribeAlbListenersResponse) {
	response = &DescribeAlbListenersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbListeners(request *DescribeAlbListenersRequest) string {
	return c.DescribeAlbListenersWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbListenersSend(request *DescribeAlbListenersRequest) (*DescribeAlbListenersResponse, error) {
	statusCode, msg, err := c.DescribeAlbListenersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAlbListenersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlbListenersWithContext(ctx context.Context, request *DescribeAlbListenersRequest) string {
	if request == nil {
		request = NewDescribeAlbListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbListenersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlbListenersWithContextV2(ctx context.Context, request *DescribeAlbListenersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlbListenersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbListenersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAlbRuleGroupRequest() (request *CreateAlbRuleGroupRequest) {
	request = &CreateAlbRuleGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateAlbRuleGroup")
	return
}

func NewCreateAlbRuleGroupResponse() (response *CreateAlbRuleGroupResponse) {
	response = &CreateAlbRuleGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlbRuleGroup(request *CreateAlbRuleGroupRequest) string {
	return c.CreateAlbRuleGroupWithContext(context.Background(), request)
}

func (c *Client) CreateAlbRuleGroupSend(request *CreateAlbRuleGroupRequest) (*CreateAlbRuleGroupResponse, error) {
	statusCode, msg, err := c.CreateAlbRuleGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAlbRuleGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAlbRuleGroupWithContext(ctx context.Context, request *CreateAlbRuleGroupRequest) string {
	if request == nil {
		request = NewCreateAlbRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateAlbRuleGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAlbRuleGroupWithContextV2(ctx context.Context, request *CreateAlbRuleGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAlbRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateAlbRuleGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlbRuleGroupRequest() (request *DeleteAlbRuleGroupRequest) {
	request = &DeleteAlbRuleGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteAlbRuleGroup")
	return
}

func NewDeleteAlbRuleGroupResponse() (response *DeleteAlbRuleGroupResponse) {
	response = &DeleteAlbRuleGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlbRuleGroup(request *DeleteAlbRuleGroupRequest) string {
	return c.DeleteAlbRuleGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbRuleGroupSend(request *DeleteAlbRuleGroupRequest) (*DeleteAlbRuleGroupResponse, error) {
	statusCode, msg, err := c.DeleteAlbRuleGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlbRuleGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlbRuleGroupWithContext(ctx context.Context, request *DeleteAlbRuleGroupRequest) string {
	if request == nil {
		request = NewDeleteAlbRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbRuleGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlbRuleGroupWithContextV2(ctx context.Context, request *DeleteAlbRuleGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlbRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbRuleGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlbRuleGroupsRequest() (request *DescribeAlbRuleGroupsRequest) {
	request = &DescribeAlbRuleGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeAlbRuleGroups")
	return
}

func NewDescribeAlbRuleGroupsResponse() (response *DescribeAlbRuleGroupsResponse) {
	response = &DescribeAlbRuleGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbRuleGroups(request *DescribeAlbRuleGroupsRequest) string {
	return c.DescribeAlbRuleGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbRuleGroupsSend(request *DescribeAlbRuleGroupsRequest) (*DescribeAlbRuleGroupsResponse, error) {
	statusCode, msg, err := c.DescribeAlbRuleGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAlbRuleGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlbRuleGroupsWithContext(ctx context.Context, request *DescribeAlbRuleGroupsRequest) string {
	if request == nil {
		request = NewDescribeAlbRuleGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbRuleGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlbRuleGroupsWithContextV2(ctx context.Context, request *DescribeAlbRuleGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlbRuleGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbRuleGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAlbRuleGroupRequest() (request *ModifyAlbRuleGroupRequest) {
	request = &ModifyAlbRuleGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyAlbRuleGroup")
	return
}

func NewModifyAlbRuleGroupResponse() (response *ModifyAlbRuleGroupResponse) {
	response = &ModifyAlbRuleGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAlbRuleGroup(request *ModifyAlbRuleGroupRequest) string {
	return c.ModifyAlbRuleGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyAlbRuleGroupSend(request *ModifyAlbRuleGroupRequest) (*ModifyAlbRuleGroupResponse, error) {
	statusCode, msg, err := c.ModifyAlbRuleGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAlbRuleGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAlbRuleGroupWithContext(ctx context.Context, request *ModifyAlbRuleGroupRequest) string {
	if request == nil {
		request = NewModifyAlbRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAlbRuleGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAlbRuleGroupWithContextV2(ctx context.Context, request *ModifyAlbRuleGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAlbRuleGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyAlbRuleGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddAlbRuleRequest() (request *AddAlbRuleRequest) {
	request = &AddAlbRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AddAlbRule")
	return
}

func NewAddAlbRuleResponse() (response *AddAlbRuleResponse) {
	response = &AddAlbRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddAlbRule(request *AddAlbRuleRequest) string {
	return c.AddAlbRuleWithContext(context.Background(), request)
}

func (c *Client) AddAlbRuleSend(request *AddAlbRuleRequest) (*AddAlbRuleResponse, error) {
	statusCode, msg, err := c.AddAlbRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddAlbRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddAlbRuleWithContext(ctx context.Context, request *AddAlbRuleRequest) string {
	if request == nil {
		request = NewAddAlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddAlbRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddAlbRuleWithContextV2(ctx context.Context, request *AddAlbRuleRequest) (int, string, error) {
	if request == nil {
		request = NewAddAlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddAlbRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlbRuleRequest() (request *DeleteAlbRuleRequest) {
	request = &DeleteAlbRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteAlbRule")
	return
}

func NewDeleteAlbRuleResponse() (response *DeleteAlbRuleResponse) {
	response = &DeleteAlbRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlbRule(request *DeleteAlbRuleRequest) string {
	return c.DeleteAlbRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbRuleSend(request *DeleteAlbRuleRequest) (*DeleteAlbRuleResponse, error) {
	statusCode, msg, err := c.DeleteAlbRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlbRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlbRuleWithContext(ctx context.Context, request *DeleteAlbRuleRequest) string {
	if request == nil {
		request = NewDeleteAlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlbRuleWithContextV2(ctx context.Context, request *DeleteAlbRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlbRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAlbListenerCertGroupRequest() (request *CreateAlbListenerCertGroupRequest) {
	request = &CreateAlbListenerCertGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateAlbListenerCertGroup")
	return
}

func NewCreateAlbListenerCertGroupResponse() (response *CreateAlbListenerCertGroupResponse) {
	response = &CreateAlbListenerCertGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlbListenerCertGroup(request *CreateAlbListenerCertGroupRequest) string {
	return c.CreateAlbListenerCertGroupWithContext(context.Background(), request)
}

func (c *Client) CreateAlbListenerCertGroupSend(request *CreateAlbListenerCertGroupRequest) (*CreateAlbListenerCertGroupResponse, error) {
	statusCode, msg, err := c.CreateAlbListenerCertGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAlbListenerCertGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAlbListenerCertGroupWithContext(ctx context.Context, request *CreateAlbListenerCertGroupRequest) string {
	if request == nil {
		request = NewCreateAlbListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbListenerCertGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAlbListenerCertGroupWithContextV2(ctx context.Context, request *CreateAlbListenerCertGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAlbListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbListenerCertGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlbListenerCertGroupRequest() (request *DeleteAlbListenerCertGroupRequest) {
	request = &DeleteAlbListenerCertGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteAlbListenerCertGroup")
	return
}

func NewDeleteAlbListenerCertGroupResponse() (response *DeleteAlbListenerCertGroupResponse) {
	response = &DeleteAlbListenerCertGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlbListenerCertGroup(request *DeleteAlbListenerCertGroupRequest) string {
	return c.DeleteAlbListenerCertGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbListenerCertGroupSend(request *DeleteAlbListenerCertGroupRequest) (*DeleteAlbListenerCertGroupResponse, error) {
	statusCode, msg, err := c.DeleteAlbListenerCertGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlbListenerCertGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlbListenerCertGroupWithContext(ctx context.Context, request *DeleteAlbListenerCertGroupRequest) string {
	if request == nil {
		request = NewDeleteAlbListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbListenerCertGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlbListenerCertGroupWithContextV2(ctx context.Context, request *DeleteAlbListenerCertGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlbListenerCertGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbListenerCertGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlbListenerCertGroupsRequest() (request *DescribeAlbListenerCertGroupsRequest) {
	request = &DescribeAlbListenerCertGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeAlbListenerCertGroups")
	return
}

func NewDescribeAlbListenerCertGroupsResponse() (response *DescribeAlbListenerCertGroupsResponse) {
	response = &DescribeAlbListenerCertGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbListenerCertGroups(request *DescribeAlbListenerCertGroupsRequest) string {
	return c.DescribeAlbListenerCertGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbListenerCertGroupsSend(request *DescribeAlbListenerCertGroupsRequest) (*DescribeAlbListenerCertGroupsResponse, error) {
	statusCode, msg, err := c.DescribeAlbListenerCertGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAlbListenerCertGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlbListenerCertGroupsWithContext(ctx context.Context, request *DescribeAlbListenerCertGroupsRequest) string {
	if request == nil {
		request = NewDescribeAlbListenerCertGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbListenerCertGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlbListenerCertGroupsWithContextV2(ctx context.Context, request *DescribeAlbListenerCertGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlbListenerCertGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbListenerCertGroupsResponse()
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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
func NewSetEnableAlbAccessLogRequest() (request *SetEnableAlbAccessLogRequest) {
	request = &SetEnableAlbAccessLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetEnableAlbAccessLog")
	return
}

func NewSetEnableAlbAccessLogResponse() (response *SetEnableAlbAccessLogResponse) {
	response = &SetEnableAlbAccessLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetEnableAlbAccessLog(request *SetEnableAlbAccessLogRequest) string {
	return c.SetEnableAlbAccessLogWithContext(context.Background(), request)
}

func (c *Client) SetEnableAlbAccessLogSend(request *SetEnableAlbAccessLogRequest) (*SetEnableAlbAccessLogResponse, error) {
	statusCode, msg, err := c.SetEnableAlbAccessLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetEnableAlbAccessLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetEnableAlbAccessLogWithContext(ctx context.Context, request *SetEnableAlbAccessLogRequest) string {
	if request == nil {
		request = NewSetEnableAlbAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetEnableAlbAccessLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetEnableAlbAccessLogWithContextV2(ctx context.Context, request *SetEnableAlbAccessLogRequest) (int, string, error) {
	if request == nil {
		request = NewSetEnableAlbAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetEnableAlbAccessLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetAlbAccessLogRequest() (request *SetAlbAccessLogRequest) {
	request = &SetAlbAccessLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetAlbAccessLog")
	return
}

func NewSetAlbAccessLogResponse() (response *SetAlbAccessLogResponse) {
	response = &SetAlbAccessLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetAlbAccessLog(request *SetAlbAccessLogRequest) string {
	return c.SetAlbAccessLogWithContext(context.Background(), request)
}

func (c *Client) SetAlbAccessLogSend(request *SetAlbAccessLogRequest) (*SetAlbAccessLogResponse, error) {
	statusCode, msg, err := c.SetAlbAccessLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetAlbAccessLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetAlbAccessLogWithContext(ctx context.Context, request *SetAlbAccessLogRequest) string {
	if request == nil {
		request = NewSetAlbAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetAlbAccessLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetAlbAccessLogWithContextV2(ctx context.Context, request *SetAlbAccessLogRequest) (int, string, error) {
	if request == nil {
		request = NewSetAlbAccessLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetAlbAccessLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCloneLoadBalancerRequest() (request *CloneLoadBalancerRequest) {
	request = &CloneLoadBalancerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CloneLoadBalancer")
	return
}

func NewCloneLoadBalancerResponse() (response *CloneLoadBalancerResponse) {
	response = &CloneLoadBalancerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloneLoadBalancer(request *CloneLoadBalancerRequest) string {
	return c.CloneLoadBalancerWithContext(context.Background(), request)
}

func (c *Client) CloneLoadBalancerSend(request *CloneLoadBalancerRequest) (*CloneLoadBalancerResponse, error) {
	statusCode, msg, err := c.CloneLoadBalancerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CloneLoadBalancerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloneLoadBalancerWithContext(ctx context.Context, request *CloneLoadBalancerRequest) string {
	if request == nil {
		request = NewCloneLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneLoadBalancerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloneLoadBalancerWithContextV2(ctx context.Context, request *CloneLoadBalancerRequest) (int, string, error) {
	if request == nil {
		request = NewCloneLoadBalancerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneLoadBalancerResponse()
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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

	if msg == "" {
		return nil, nil
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
func NewCreateAlbBackendServerGroupRequest() (request *CreateAlbBackendServerGroupRequest) {
	request = &CreateAlbBackendServerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "CreateAlbBackendServerGroup")
	return
}

func NewCreateAlbBackendServerGroupResponse() (response *CreateAlbBackendServerGroupResponse) {
	response = &CreateAlbBackendServerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlbBackendServerGroup(request *CreateAlbBackendServerGroupRequest) string {
	return c.CreateAlbBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) CreateAlbBackendServerGroupSend(request *CreateAlbBackendServerGroupRequest) (*CreateAlbBackendServerGroupResponse, error) {
	statusCode, msg, err := c.CreateAlbBackendServerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAlbBackendServerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAlbBackendServerGroupWithContext(ctx context.Context, request *CreateAlbBackendServerGroupRequest) string {
	if request == nil {
		request = NewCreateAlbBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbBackendServerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAlbBackendServerGroupWithContextV2(ctx context.Context, request *CreateAlbBackendServerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAlbBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbBackendServerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlbBackendServerGroupRequest() (request *DeleteAlbBackendServerGroupRequest) {
	request = &DeleteAlbBackendServerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeleteAlbBackendServerGroup")
	return
}

func NewDeleteAlbBackendServerGroupResponse() (response *DeleteAlbBackendServerGroupResponse) {
	response = &DeleteAlbBackendServerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlbBackendServerGroup(request *DeleteAlbBackendServerGroupRequest) string {
	return c.DeleteAlbBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbBackendServerGroupSend(request *DeleteAlbBackendServerGroupRequest) (*DeleteAlbBackendServerGroupResponse, error) {
	statusCode, msg, err := c.DeleteAlbBackendServerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlbBackendServerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlbBackendServerGroupWithContext(ctx context.Context, request *DeleteAlbBackendServerGroupRequest) string {
	if request == nil {
		request = NewDeleteAlbBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbBackendServerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlbBackendServerGroupWithContextV2(ctx context.Context, request *DeleteAlbBackendServerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlbBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbBackendServerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAlbBackendServerGroupRequest() (request *ModifyAlbBackendServerGroupRequest) {
	request = &ModifyAlbBackendServerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyAlbBackendServerGroup")
	return
}

func NewModifyAlbBackendServerGroupResponse() (response *ModifyAlbBackendServerGroupResponse) {
	response = &ModifyAlbBackendServerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAlbBackendServerGroup(request *ModifyAlbBackendServerGroupRequest) string {
	return c.ModifyAlbBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyAlbBackendServerGroupSend(request *ModifyAlbBackendServerGroupRequest) (*ModifyAlbBackendServerGroupResponse, error) {
	statusCode, msg, err := c.ModifyAlbBackendServerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAlbBackendServerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAlbBackendServerGroupWithContext(ctx context.Context, request *ModifyAlbBackendServerGroupRequest) string {
	if request == nil {
		request = NewModifyAlbBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbBackendServerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAlbBackendServerGroupWithContextV2(ctx context.Context, request *ModifyAlbBackendServerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAlbBackendServerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbBackendServerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlbBackendServerGroupsRequest() (request *DescribeAlbBackendServerGroupsRequest) {
	request = &DescribeAlbBackendServerGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeAlbBackendServerGroups")
	return
}

func NewDescribeAlbBackendServerGroupsResponse() (response *DescribeAlbBackendServerGroupsResponse) {
	response = &DescribeAlbBackendServerGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbBackendServerGroups(request *DescribeAlbBackendServerGroupsRequest) string {
	return c.DescribeAlbBackendServerGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbBackendServerGroupsSend(request *DescribeAlbBackendServerGroupsRequest) (*DescribeAlbBackendServerGroupsResponse, error) {
	statusCode, msg, err := c.DescribeAlbBackendServerGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAlbBackendServerGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlbBackendServerGroupsWithContext(ctx context.Context, request *DescribeAlbBackendServerGroupsRequest) string {
	if request == nil {
		request = NewDescribeAlbBackendServerGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbBackendServerGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlbBackendServerGroupsWithContextV2(ctx context.Context, request *DescribeAlbBackendServerGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlbBackendServerGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbBackendServerGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRegisterAlbBackendServerRequest() (request *RegisterAlbBackendServerRequest) {
	request = &RegisterAlbBackendServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "RegisterAlbBackendServer")
	return
}

func NewRegisterAlbBackendServerResponse() (response *RegisterAlbBackendServerResponse) {
	response = &RegisterAlbBackendServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RegisterAlbBackendServer(request *RegisterAlbBackendServerRequest) string {
	return c.RegisterAlbBackendServerWithContext(context.Background(), request)
}

func (c *Client) RegisterAlbBackendServerSend(request *RegisterAlbBackendServerRequest) (*RegisterAlbBackendServerResponse, error) {
	statusCode, msg, err := c.RegisterAlbBackendServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RegisterAlbBackendServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RegisterAlbBackendServerWithContext(ctx context.Context, request *RegisterAlbBackendServerRequest) string {
	if request == nil {
		request = NewRegisterAlbBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterAlbBackendServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RegisterAlbBackendServerWithContextV2(ctx context.Context, request *RegisterAlbBackendServerRequest) (int, string, error) {
	if request == nil {
		request = NewRegisterAlbBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterAlbBackendServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeregisterAlbBackendServerRequest() (request *DeregisterAlbBackendServerRequest) {
	request = &DeregisterAlbBackendServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DeregisterAlbBackendServer")
	return
}

func NewDeregisterAlbBackendServerResponse() (response *DeregisterAlbBackendServerResponse) {
	response = &DeregisterAlbBackendServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeregisterAlbBackendServer(request *DeregisterAlbBackendServerRequest) string {
	return c.DeregisterAlbBackendServerWithContext(context.Background(), request)
}

func (c *Client) DeregisterAlbBackendServerSend(request *DeregisterAlbBackendServerRequest) (*DeregisterAlbBackendServerResponse, error) {
	statusCode, msg, err := c.DeregisterAlbBackendServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeregisterAlbBackendServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeregisterAlbBackendServerWithContext(ctx context.Context, request *DeregisterAlbBackendServerRequest) string {
	if request == nil {
		request = NewDeregisterAlbBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeregisterAlbBackendServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeregisterAlbBackendServerWithContextV2(ctx context.Context, request *DeregisterAlbBackendServerRequest) (int, string, error) {
	if request == nil {
		request = NewDeregisterAlbBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeregisterAlbBackendServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAlbBackendServerRequest() (request *ModifyAlbBackendServerRequest) {
	request = &ModifyAlbBackendServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "ModifyAlbBackendServer")
	return
}

func NewModifyAlbBackendServerResponse() (response *ModifyAlbBackendServerResponse) {
	response = &ModifyAlbBackendServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAlbBackendServer(request *ModifyAlbBackendServerRequest) string {
	return c.ModifyAlbBackendServerWithContext(context.Background(), request)
}

func (c *Client) ModifyAlbBackendServerSend(request *ModifyAlbBackendServerRequest) (*ModifyAlbBackendServerResponse, error) {
	statusCode, msg, err := c.ModifyAlbBackendServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAlbBackendServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAlbBackendServerWithContext(ctx context.Context, request *ModifyAlbBackendServerRequest) string {
	if request == nil {
		request = NewModifyAlbBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbBackendServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAlbBackendServerWithContextV2(ctx context.Context, request *ModifyAlbBackendServerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAlbBackendServerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbBackendServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAlbBackendServersRequest() (request *DescribeAlbBackendServersRequest) {
	request = &DescribeAlbBackendServersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "DescribeAlbBackendServers")
	return
}

func NewDescribeAlbBackendServersResponse() (response *DescribeAlbBackendServersResponse) {
	response = &DescribeAlbBackendServersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbBackendServers(request *DescribeAlbBackendServersRequest) string {
	return c.DescribeAlbBackendServersWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbBackendServersSend(request *DescribeAlbBackendServersRequest) (*DescribeAlbBackendServersResponse, error) {
	statusCode, msg, err := c.DescribeAlbBackendServersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAlbBackendServersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAlbBackendServersWithContext(ctx context.Context, request *DescribeAlbBackendServersRequest) string {
	if request == nil {
		request = NewDescribeAlbBackendServersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbBackendServersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAlbBackendServersWithContextV2(ctx context.Context, request *DescribeAlbBackendServersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAlbBackendServersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbBackendServersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRegisterBackendServerGroupWithListenerRequest() (request *RegisterBackendServerGroupWithListenerRequest) {
	request = &RegisterBackendServerGroupWithListenerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "RegisterBackendServerGroupWithListener")
	return
}

func NewRegisterBackendServerGroupWithListenerResponse() (response *RegisterBackendServerGroupWithListenerResponse) {
	response = &RegisterBackendServerGroupWithListenerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RegisterBackendServerGroupWithListener(request *RegisterBackendServerGroupWithListenerRequest) string {
	return c.RegisterBackendServerGroupWithListenerWithContext(context.Background(), request)
}

func (c *Client) RegisterBackendServerGroupWithListenerSend(request *RegisterBackendServerGroupWithListenerRequest) (*RegisterBackendServerGroupWithListenerResponse, error) {
	statusCode, msg, err := c.RegisterBackendServerGroupWithListenerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RegisterBackendServerGroupWithListenerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RegisterBackendServerGroupWithListenerWithContext(ctx context.Context, request *RegisterBackendServerGroupWithListenerRequest) string {
	if request == nil {
		request = NewRegisterBackendServerGroupWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterBackendServerGroupWithListenerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RegisterBackendServerGroupWithListenerWithContextV2(ctx context.Context, request *RegisterBackendServerGroupWithListenerRequest) (int, string, error) {
	if request == nil {
		request = NewRegisterBackendServerGroupWithListenerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterBackendServerGroupWithListenerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetPrivateLinkDeleteProtectionRequest() (request *SetPrivateLinkDeleteProtectionRequest) {
	request = &SetPrivateLinkDeleteProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetPrivateLinkDeleteProtection")
	return
}

func NewSetPrivateLinkDeleteProtectionResponse() (response *SetPrivateLinkDeleteProtectionResponse) {
	response = &SetPrivateLinkDeleteProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetPrivateLinkDeleteProtection(request *SetPrivateLinkDeleteProtectionRequest) string {
	return c.SetPrivateLinkDeleteProtectionWithContext(context.Background(), request)
}

func (c *Client) SetPrivateLinkDeleteProtectionSend(request *SetPrivateLinkDeleteProtectionRequest) (*SetPrivateLinkDeleteProtectionResponse, error) {
	statusCode, msg, err := c.SetPrivateLinkDeleteProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetPrivateLinkDeleteProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetPrivateLinkDeleteProtectionWithContext(ctx context.Context, request *SetPrivateLinkDeleteProtectionRequest) string {
	if request == nil {
		request = NewSetPrivateLinkDeleteProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetPrivateLinkDeleteProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetPrivateLinkDeleteProtectionWithContextV2(ctx context.Context, request *SetPrivateLinkDeleteProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewSetPrivateLinkDeleteProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetPrivateLinkDeleteProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetAlbDeleteProtectionRequest() (request *SetAlbDeleteProtectionRequest) {
	request = &SetAlbDeleteProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetAlbDeleteProtection")
	return
}

func NewSetAlbDeleteProtectionResponse() (response *SetAlbDeleteProtectionResponse) {
	response = &SetAlbDeleteProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetAlbDeleteProtection(request *SetAlbDeleteProtectionRequest) string {
	return c.SetAlbDeleteProtectionWithContext(context.Background(), request)
}

func (c *Client) SetAlbDeleteProtectionSend(request *SetAlbDeleteProtectionRequest) (*SetAlbDeleteProtectionResponse, error) {
	statusCode, msg, err := c.SetAlbDeleteProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetAlbDeleteProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetAlbDeleteProtectionWithContext(ctx context.Context, request *SetAlbDeleteProtectionRequest) string {
	if request == nil {
		request = NewSetAlbDeleteProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbDeleteProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetAlbDeleteProtectionWithContextV2(ctx context.Context, request *SetAlbDeleteProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewSetAlbDeleteProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbDeleteProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetAlbModificationProtectionRequest() (request *SetAlbModificationProtectionRequest) {
	request = &SetAlbModificationProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "SetAlbModificationProtection")
	return
}

func NewSetAlbModificationProtectionResponse() (response *SetAlbModificationProtectionResponse) {
	response = &SetAlbModificationProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetAlbModificationProtection(request *SetAlbModificationProtectionRequest) string {
	return c.SetAlbModificationProtectionWithContext(context.Background(), request)
}

func (c *Client) SetAlbModificationProtectionSend(request *SetAlbModificationProtectionRequest) (*SetAlbModificationProtectionResponse, error) {
	statusCode, msg, err := c.SetAlbModificationProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetAlbModificationProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetAlbModificationProtectionWithContext(ctx context.Context, request *SetAlbModificationProtectionRequest) string {
	if request == nil {
		request = NewSetAlbModificationProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbModificationProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetAlbModificationProtectionWithContextV2(ctx context.Context, request *SetAlbModificationProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewSetAlbModificationProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetAlbModificationProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddAlbRulesRequest() (request *AddAlbRulesRequest) {
	request = &AddAlbRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("slb", APIVersion, "AddAlbRules")
	return
}

func NewAddAlbRulesResponse() (response *AddAlbRulesResponse) {
	response = &AddAlbRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddAlbRules(request *AddAlbRulesRequest) string {
	return c.AddAlbRulesWithContext(context.Background(), request)
}

func (c *Client) AddAlbRulesSend(request *AddAlbRulesRequest) (*AddAlbRulesResponse, error) {
	statusCode, msg, err := c.AddAlbRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddAlbRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddAlbRulesWithContext(ctx context.Context, request *AddAlbRulesRequest) string {
	if request == nil {
		request = NewAddAlbRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddAlbRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddAlbRulesWithContextV2(ctx context.Context, request *AddAlbRulesRequest) (int, string, error) {
	if request == nil {
		request = NewAddAlbRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddAlbRulesResponse()
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

	if msg == "" {
		return nil, nil
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
