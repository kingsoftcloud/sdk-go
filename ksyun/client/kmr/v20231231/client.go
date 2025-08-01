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

func (c *Client) ListInstancesSend(request *ListInstancesRequest) (*ListInstancesResponse, error) {
	statusCode, msg, err := c.ListInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest) string {
	if request == nil {
		request = NewListInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListInstancesWithContextV2(ctx context.Context, request *ListInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewListInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetInstanceDetailSend(request *GetInstanceDetailRequest) (*GetInstanceDetailResponse, error) {
	statusCode, msg, err := c.GetInstanceDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetInstanceDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetInstanceDetailWithContextV2(ctx context.Context, request *GetInstanceDetailRequest) (int, string, error) {
	if request == nil {
		request = NewGetInstanceDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetInstanceDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyHostsRequest() (request *ModifyHostsRequest) {
	request = &ModifyHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ModifyHosts")
	return
}

func NewModifyHostsResponse() (response *ModifyHostsResponse) {
	response = &ModifyHostsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyHosts(request *ModifyHostsRequest) string {
	return c.ModifyHostsWithContext(context.Background(), request)
}

func (c *Client) ModifyHostsSend(request *ModifyHostsRequest) (*ModifyHostsResponse, error) {
	statusCode, msg, err := c.ModifyHostsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyHostsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyHostsWithContext(ctx context.Context, request *ModifyHostsRequest) string {
	if request == nil {
		request = NewModifyHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyHostsWithContextV2(ctx context.Context, request *ModifyHostsRequest) (int, string, error) {
	if request == nil {
		request = NewModifyHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyHostsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListAutoScaleHistorySend(request *ListAutoScaleHistoryRequest) (*ListAutoScaleHistoryResponse, error) {
	statusCode, msg, err := c.ListAutoScaleHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListAutoScaleHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ListAutoScaleHistoryWithContextV2(ctx context.Context, request *ListAutoScaleHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewListAutoScaleHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAutoScaleHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateAutoScalePolicySend(request *CreateAutoScalePolicyRequest) (*CreateAutoScalePolicyResponse, error) {
	statusCode, msg, err := c.CreateAutoScalePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateAutoScalePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateAutoScalePolicyWithContextV2(ctx context.Context, request *CreateAutoScalePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAutoScalePolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAutoScalePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListAutoScalePolicySend(request *ListAutoScalePolicyRequest) (*ListAutoScalePolicyResponse, error) {
	statusCode, msg, err := c.ListAutoScalePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListAutoScalePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ListAutoScalePolicyWithContextV2(ctx context.Context, request *ListAutoScalePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewListAutoScalePolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListAutoScalePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteAutoScalePolicySend(request *DeleteAutoScalePolicyRequest) (*DeleteAutoScalePolicyResponse, error) {
	statusCode, msg, err := c.DeleteAutoScalePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteAutoScalePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteAutoScalePolicyWithContextV2(ctx context.Context, request *DeleteAutoScalePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAutoScalePolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAutoScalePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


