package v20250220

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2025-02-20"

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

func NewListInstanceSupportBillTypesRequest() (request *ListInstanceSupportBillTypesRequest) {
	request = &ListInstanceSupportBillTypesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "ListInstanceSupportBillTypes")
	return
}

func NewListInstanceSupportBillTypesResponse() (response *ListInstanceSupportBillTypesResponse) {
	response = &ListInstanceSupportBillTypesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstanceSupportBillTypes(request *ListInstanceSupportBillTypesRequest) string {
	return c.ListInstanceSupportBillTypesWithContext(context.Background(), request)
}

func (c *Client) ListInstanceSupportBillTypesWithContext(ctx context.Context, request *ListInstanceSupportBillTypesRequest) string {
	if request == nil {
		request = NewListInstanceSupportBillTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceSupportBillTypesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAddTrialToBuyTaskRequest() (request *AddTrialToBuyTaskRequest) {
	request = &AddTrialToBuyTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "AddTrialToBuyTask")
	return
}

func NewAddTrialToBuyTaskResponse() (response *AddTrialToBuyTaskResponse) {
	response = &AddTrialToBuyTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddTrialToBuyTask(request *AddTrialToBuyTaskRequest) string {
	return c.AddTrialToBuyTaskWithContext(context.Background(), request)
}

func (c *Client) AddTrialToBuyTaskWithContext(ctx context.Context, request *AddTrialToBuyTaskRequest) string {
	if request == nil {
		request = NewAddTrialToBuyTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddTrialToBuyTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteTrialToBuyTaskRequest() (request *DeleteTrialToBuyTaskRequest) {
	request = &DeleteTrialToBuyTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "DeleteTrialToBuyTask")
	return
}

func NewDeleteTrialToBuyTaskResponse() (response *DeleteTrialToBuyTaskResponse) {
	response = &DeleteTrialToBuyTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTrialToBuyTask(request *DeleteTrialToBuyTaskRequest) string {
	return c.DeleteTrialToBuyTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteTrialToBuyTaskWithContext(ctx context.Context, request *DeleteTrialToBuyTaskRequest) string {
	if request == nil {
		request = NewDeleteTrialToBuyTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteTrialToBuyTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateTrialToBuyNowRequest() (request *CreateTrialToBuyNowRequest) {
	request = &CreateTrialToBuyNowRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("trade", APIVersion, "CreateTrialToBuyNow")
	return
}

func NewCreateTrialToBuyNowResponse() (response *CreateTrialToBuyNowResponse) {
	response = &CreateTrialToBuyNowResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTrialToBuyNow(request *CreateTrialToBuyNowRequest) string {
	return c.CreateTrialToBuyNowWithContext(context.Background(), request)
}

func (c *Client) CreateTrialToBuyNowWithContext(ctx context.Context, request *CreateTrialToBuyNowRequest) string {
	if request == nil {
		request = NewCreateTrialToBuyNowRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrialToBuyNowResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
