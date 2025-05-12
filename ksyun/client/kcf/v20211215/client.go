package v20211215

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-12-15"

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

func NewGetLogDateRequest() (request *GetLogDateRequest) {
	request = &GetLogDateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "GetLogDate")
	return
}

func NewGetLogDateResponse() (response *GetLogDateResponse) {
	response = &GetLogDateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetLogDate(request *GetLogDateRequest) string {
	return c.GetLogDateWithContext(context.Background(), request)
}

func (c *Client) GetLogDateWithContext(ctx context.Context, request *GetLogDateRequest) string {
	if request == nil {
		request = NewGetLogDateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLogDateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateFunctionRequest() (request *CreateFunctionRequest) {
	request = &CreateFunctionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "CreateFunction")
	return
}

func NewCreateFunctionResponse() (response *CreateFunctionResponse) {
	response = &CreateFunctionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFunction(request *CreateFunctionRequest) string {
	return c.CreateFunctionWithContext(context.Background(), request)
}

func (c *Client) CreateFunctionWithContext(ctx context.Context, request *CreateFunctionRequest) string {
	if request == nil {
		request = NewCreateFunctionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFunctionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCheckFunctionServiceRequest() (request *CheckFunctionServiceRequest) {
	request = &CheckFunctionServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "CheckFunctionService")
	return
}

func NewCheckFunctionServiceResponse() (response *CheckFunctionServiceResponse) {
	response = &CheckFunctionServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckFunctionService(request *CheckFunctionServiceRequest) string {
	return c.CheckFunctionServiceWithContext(context.Background(), request)
}

func (c *Client) CheckFunctionServiceWithContext(ctx context.Context, request *CheckFunctionServiceRequest) string {
	if request == nil {
		request = NewCheckFunctionServiceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCheckFunctionServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewOpenFunctionServiceRequest() (request *OpenFunctionServiceRequest) {
	request = &OpenFunctionServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "OpenFunctionService")
	return
}

func NewOpenFunctionServiceResponse() (response *OpenFunctionServiceResponse) {
	response = &OpenFunctionServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OpenFunctionService(request *OpenFunctionServiceRequest) string {
	return c.OpenFunctionServiceWithContext(context.Background(), request)
}

func (c *Client) OpenFunctionServiceWithContext(ctx context.Context, request *OpenFunctionServiceRequest) string {
	if request == nil {
		request = NewOpenFunctionServiceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOpenFunctionServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteFunctionRequest() (request *DeleteFunctionRequest) {
	request = &DeleteFunctionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "DeleteFunction")
	return
}

func NewDeleteFunctionResponse() (response *DeleteFunctionResponse) {
	response = &DeleteFunctionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFunction(request *DeleteFunctionRequest) string {
	return c.DeleteFunctionWithContext(context.Background(), request)
}

func (c *Client) DeleteFunctionWithContext(ctx context.Context, request *DeleteFunctionRequest) string {
	if request == nil {
		request = NewDeleteFunctionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFunctionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateTriggerRequest() (request *CreateTriggerRequest) {
	request = &CreateTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "CreateTrigger")
	return
}

func NewCreateTriggerResponse() (response *CreateTriggerResponse) {
	response = &CreateTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTrigger(request *CreateTriggerRequest) string {
	return c.CreateTriggerWithContext(context.Background(), request)
}

func (c *Client) CreateTriggerWithContext(ctx context.Context, request *CreateTriggerRequest) string {
	if request == nil {
		request = NewCreateTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteTriggerRequest() (request *DeleteTriggerRequest) {
	request = &DeleteTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "DeleteTrigger")
	return
}

func NewDeleteTriggerResponse() (response *DeleteTriggerResponse) {
	response = &DeleteTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTrigger(request *DeleteTriggerRequest) string {
	return c.DeleteTriggerWithContext(context.Background(), request)
}

func (c *Client) DeleteTriggerWithContext(ctx context.Context, request *DeleteTriggerRequest) string {
	if request == nil {
		request = NewDeleteTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyFunctionRequest() (request *ModifyFunctionRequest) {
	request = &ModifyFunctionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "ModifyFunction")
	return
}

func NewModifyFunctionResponse() (response *ModifyFunctionResponse) {
	response = &ModifyFunctionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyFunction(request *ModifyFunctionRequest) string {
	return c.ModifyFunctionWithContext(context.Background(), request)
}

func (c *Client) ModifyFunctionWithContext(ctx context.Context, request *ModifyFunctionRequest) string {
	if request == nil {
		request = NewModifyFunctionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyFunctionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeTriggersRequest() (request *DescribeTriggersRequest) {
	request = &DescribeTriggersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcf", APIVersion, "DescribeTriggers")
	return
}

func NewDescribeTriggersResponse() (response *DescribeTriggersResponse) {
	response = &DescribeTriggersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTriggers(request *DescribeTriggersRequest) string {
	return c.DescribeTriggersWithContext(context.Background(), request)
}

func (c *Client) DescribeTriggersWithContext(ctx context.Context, request *DescribeTriggersRequest) string {
	if request == nil {
		request = NewDescribeTriggersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTriggersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
