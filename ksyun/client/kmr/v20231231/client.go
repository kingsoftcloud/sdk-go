package v20231231

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2023-12-31"

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

func NewListInstancesRequest() (request *ListInstancesRequest) {
	request = &ListInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListInstances")
	return
}

func NewListInstancesResponse() (response *ListInstancesResponse) {
	response = &ListInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstances(request *ListInstancesRequest) string {
	return c.ListInstancesWithContext(context.Background(), request)
}

func (c *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest) string {
	if request == nil {
		request = NewListInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetInstanceDetailRequest() (request *GetInstanceDetailRequest) {
	request = &GetInstanceDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "GetInstanceDetail")
	return
}

func NewGetInstanceDetailResponse() (response *GetInstanceDetailResponse) {
	response = &GetInstanceDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetInstanceDetail(request *GetInstanceDetailRequest) string {
	return c.GetInstanceDetailWithContext(context.Background(), request)
}

func (c *Client) GetInstanceDetailWithContext(ctx context.Context, request *GetInstanceDetailRequest) string {
	if request == nil {
		request = NewGetInstanceDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetInstanceDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListAutoScaleHistoryRequest() (request *ListAutoScaleHistoryRequest) {
	request = &ListAutoScaleHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListAutoScaleHistory")
	return
}

func NewListAutoScaleHistoryResponse() (response *ListAutoScaleHistoryResponse) {
	response = &ListAutoScaleHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListAutoScaleHistory(request *ListAutoScaleHistoryRequest) string {
	return c.ListAutoScaleHistoryWithContext(context.Background(), request)
}

func (c *Client) ListAutoScaleHistoryWithContext(ctx context.Context, request *ListAutoScaleHistoryRequest) string {
	if request == nil {
		request = NewListAutoScaleHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAutoScaleHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateAutoScalePolicyRequest() (request *CreateAutoScalePolicyRequest) {
	request = &CreateAutoScalePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "CreateAutoScalePolicy")
	return
}

func NewCreateAutoScalePolicyResponse() (response *CreateAutoScalePolicyResponse) {
	response = &CreateAutoScalePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAutoScalePolicy(request *CreateAutoScalePolicyRequest) string {
	return c.CreateAutoScalePolicyWithContext(context.Background(), request)
}

func (c *Client) CreateAutoScalePolicyWithContext(ctx context.Context, request *CreateAutoScalePolicyRequest) string {
	if request == nil {
		request = NewCreateAutoScalePolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAutoScalePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListAutoScalePolicyRequest() (request *ListAutoScalePolicyRequest) {
	request = &ListAutoScalePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListAutoScalePolicy")
	return
}

func NewListAutoScalePolicyResponse() (response *ListAutoScalePolicyResponse) {
	response = &ListAutoScalePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListAutoScalePolicy(request *ListAutoScalePolicyRequest) string {
	return c.ListAutoScalePolicyWithContext(context.Background(), request)
}

func (c *Client) ListAutoScalePolicyWithContext(ctx context.Context, request *ListAutoScalePolicyRequest) string {
	if request == nil {
		request = NewListAutoScalePolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAutoScalePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteAutoScalePolicyRequest() (request *DeleteAutoScalePolicyRequest) {
	request = &DeleteAutoScalePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DeleteAutoScalePolicy")
	return
}

func NewDeleteAutoScalePolicyResponse() (response *DeleteAutoScalePolicyResponse) {
	response = &DeleteAutoScalePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAutoScalePolicy(request *DeleteAutoScalePolicyRequest) string {
	return c.DeleteAutoScalePolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteAutoScalePolicyWithContext(ctx context.Context, request *DeleteAutoScalePolicyRequest) string {
	if request == nil {
		request = NewDeleteAutoScalePolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAutoScalePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
