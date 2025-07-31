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

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) string {
	return c.DescribeInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesSend(request *DescribeInstancesRequest) (*DescribeInstancesResponse, error) {
	statusCode, msg, err := c.DescribeInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) string {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstancesWithContextV2(ctx context.Context, request *DescribeInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RunInstances")
	return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RunInstances(request *RunInstancesRequest) string {
	return c.RunInstancesWithContext(context.Background(), request)
}

func (c *Client) RunInstancesSend(request *RunInstancesRequest) (*RunInstancesResponse, error) {
	statusCode, msg, err := c.RunInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RunInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) string {
	if request == nil {
		request = NewRunInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRunInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RunInstancesWithContextV2(ctx context.Context, request *RunInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewRunInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRunInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "StartInstances")
	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartInstances(request *StartInstancesRequest) string {
	return c.StartInstancesWithContext(context.Background(), request)
}

func (c *Client) StartInstancesSend(request *StartInstancesRequest) (*StartInstancesResponse, error) {
	statusCode, msg, err := c.StartInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) string {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartInstancesWithContextV2(ctx context.Context, request *StartInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "StopInstances")
	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopInstances(request *StopInstancesRequest) string {
	return c.StopInstancesWithContext(context.Background(), request)
}

func (c *Client) StopInstancesSend(request *StopInstancesRequest) (*StopInstancesResponse, error) {
	statusCode, msg, err := c.StopInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StopInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) string {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopInstancesWithContextV2(ctx context.Context, request *StopInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RebootInstances")
	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootInstances(request *RebootInstancesRequest) string {
	return c.RebootInstancesWithContext(context.Background(), request)
}

func (c *Client) RebootInstancesSend(request *RebootInstancesRequest) (*RebootInstancesResponse, error) {
	statusCode, msg, err := c.RebootInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RebootInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) string {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RebootInstancesWithContextV2(ctx context.Context, request *RebootInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAttributeRequest() (request *ModifyInstanceAttributeRequest) {
	request = &ModifyInstanceAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceAttribute")
	return
}

func NewModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
	response = &ModifyInstanceAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) string {
	return c.ModifyInstanceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAttributeSend(request *ModifyInstanceAttributeRequest) (*ModifyInstanceAttributeResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyInstanceAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceAttributeWithContext(ctx context.Context, request *ModifyInstanceAttributeRequest) string {
	if request == nil {
		request = NewModifyInstanceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceAttributeWithContextV2(ctx context.Context, request *ModifyInstanceAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceTypeRequest() (request *ModifyInstanceTypeRequest) {
	request = &ModifyInstanceTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceType")
	return
}

func NewModifyInstanceTypeResponse() (response *ModifyInstanceTypeResponse) {
	response = &ModifyInstanceTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) string {
	return c.ModifyInstanceTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceTypeSend(request *ModifyInstanceTypeRequest) (*ModifyInstanceTypeResponse, error) {
	statusCode, msg, err := c.ModifyInstanceTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyInstanceTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceTypeWithContext(ctx context.Context, request *ModifyInstanceTypeRequest) string {
	if request == nil {
		request = NewModifyInstanceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceTypeWithContextV2(ctx context.Context, request *ModifyInstanceTypeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TerminateInstances(request *TerminateInstancesRequest) string {
	return c.TerminateInstancesWithContext(context.Background(), request)
}

func (c *Client) TerminateInstancesSend(request *TerminateInstancesRequest) (*TerminateInstancesResponse, error) {
	statusCode, msg, err := c.TerminateInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct TerminateInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) string {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTerminateInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) TerminateInstancesWithContextV2(ctx context.Context, request *TerminateInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTerminateInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) string {
	return c.DescribeImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeImagesSend(request *DescribeImagesRequest) (*DescribeImagesResponse, error) {
	statusCode, msg, err := c.DescribeImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) string {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImagesWithContextV2(ctx context.Context, request *DescribeImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
	request = &ModifyImageAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyImageAttribute")
	return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
	response = &ModifyImageAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) string {
	return c.ModifyImageAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyImageAttributeSend(request *ModifyImageAttributeRequest) (*ModifyImageAttributeResponse, error) {
	statusCode, msg, err := c.ModifyImageAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyImageAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest) string {
	if request == nil {
		request = NewModifyImageAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyImageAttributeWithContextV2(ctx context.Context, request *ModifyImageAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyImageAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceImageRequest() (request *ModifyInstanceImageRequest) {
	request = &ModifyInstanceImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceImage")
	return
}

func NewModifyInstanceImageResponse() (response *ModifyInstanceImageResponse) {
	response = &ModifyInstanceImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceImage(request *ModifyInstanceImageRequest) string {
	return c.ModifyInstanceImageWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceImageSend(request *ModifyInstanceImageRequest) (*ModifyInstanceImageResponse, error) {
	statusCode, msg, err := c.ModifyInstanceImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyInstanceImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceImageWithContext(ctx context.Context, request *ModifyInstanceImageRequest) string {
	if request == nil {
		request = NewModifyInstanceImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceImageWithContextV2(ctx context.Context, request *ModifyInstanceImageRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImage(request *CreateImageRequest) string {
	return c.CreateImageWithContext(context.Background(), request)
}

func (c *Client) CreateImageSend(request *CreateImageRequest) (*CreateImageResponse, error) {
	statusCode, msg, err := c.CreateImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) string {
	if request == nil {
		request = NewCreateImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateImageWithContextV2(ctx context.Context, request *CreateImageRequest) (int, string, error) {
	if request == nil {
		request = NewCreateImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRemoveImagesRequest() (request *RemoveImagesRequest) {
	request = &RemoveImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RemoveImages")
	return
}

func NewRemoveImagesResponse() (response *RemoveImagesResponse) {
	response = &RemoveImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemoveImages(request *RemoveImagesRequest) string {
	return c.RemoveImagesWithContext(context.Background(), request)
}

func (c *Client) RemoveImagesSend(request *RemoveImagesRequest) (*RemoveImagesResponse, error) {
	statusCode, msg, err := c.RemoveImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RemoveImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RemoveImagesWithContext(ctx context.Context, request *RemoveImagesRequest) string {
	if request == nil {
		request = NewRemoveImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemoveImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RemoveImagesWithContextV2(ctx context.Context, request *RemoveImagesRequest) (int, string, error) {
	if request == nil {
		request = NewRemoveImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemoveImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
	request = &ModifyNetworkInterfaceAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyNetworkInterfaceAttribute")
	return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
	response = &ModifyNetworkInterfaceAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) string {
	return c.ModifyNetworkInterfaceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyNetworkInterfaceAttributeSend(request *ModifyNetworkInterfaceAttributeRequest) (*ModifyNetworkInterfaceAttributeResponse, error) {
	statusCode, msg, err := c.ModifyNetworkInterfaceAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNetworkInterfaceAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) string {
	if request == nil {
		request = NewModifyNetworkInterfaceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkInterfaceAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNetworkInterfaceAttributeWithContextV2(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNetworkInterfaceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkInterfaceAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
	request = &AttachNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachNetworkInterface")
	return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
	response = &AttachNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) string {
	return c.AttachNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) AttachNetworkInterfaceSend(request *AttachNetworkInterfaceRequest) (*AttachNetworkInterfaceResponse, error) {
	statusCode, msg, err := c.AttachNetworkInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AttachNetworkInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachNetworkInterfaceWithContext(ctx context.Context, request *AttachNetworkInterfaceRequest) string {
	if request == nil {
		request = NewAttachNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachNetworkInterfaceWithContextV2(ctx context.Context, request *AttachNetworkInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewAttachNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachNetworkInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
	request = &DetachNetworkInterfaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachNetworkInterface")
	return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
	response = &DetachNetworkInterfaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) string {
	return c.DetachNetworkInterfaceWithContext(context.Background(), request)
}

func (c *Client) DetachNetworkInterfaceSend(request *DetachNetworkInterfaceRequest) (*DetachNetworkInterfaceResponse, error) {
	statusCode, msg, err := c.DetachNetworkInterfaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachNetworkInterfaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachNetworkInterfaceWithContext(ctx context.Context, request *DetachNetworkInterfaceRequest) string {
	if request == nil {
		request = NewDetachNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachNetworkInterfaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachNetworkInterfaceWithContextV2(ctx context.Context, request *DetachNetworkInterfaceRequest) (int, string, error) {
	if request == nil {
		request = NewDetachNetworkInterfaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachNetworkInterfaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeLocalVolumesRequest() (request *DescribeLocalVolumesRequest) {
	request = &DescribeLocalVolumesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeLocalVolumes")
	return
}

func NewDescribeLocalVolumesResponse() (response *DescribeLocalVolumesResponse) {
	response = &DescribeLocalVolumesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLocalVolumes(request *DescribeLocalVolumesRequest) string {
	return c.DescribeLocalVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeLocalVolumesSend(request *DescribeLocalVolumesRequest) (*DescribeLocalVolumesResponse, error) {
	statusCode, msg, err := c.DescribeLocalVolumesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeLocalVolumesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeLocalVolumesWithContext(ctx context.Context, request *DescribeLocalVolumesRequest) string {
	if request == nil {
		request = NewDescribeLocalVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLocalVolumesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeLocalVolumesWithContextV2(ctx context.Context, request *DescribeLocalVolumesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeLocalVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLocalVolumesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateLocalVolumeSnapshotRequest() (request *CreateLocalVolumeSnapshotRequest) {
	request = &CreateLocalVolumeSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateLocalVolumeSnapshot")
	return
}

func NewCreateLocalVolumeSnapshotResponse() (response *CreateLocalVolumeSnapshotResponse) {
	response = &CreateLocalVolumeSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLocalVolumeSnapshot(request *CreateLocalVolumeSnapshotRequest) string {
	return c.CreateLocalVolumeSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateLocalVolumeSnapshotSend(request *CreateLocalVolumeSnapshotRequest) (*CreateLocalVolumeSnapshotResponse, error) {
	statusCode, msg, err := c.CreateLocalVolumeSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateLocalVolumeSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateLocalVolumeSnapshotWithContext(ctx context.Context, request *CreateLocalVolumeSnapshotRequest) string {
	if request == nil {
		request = NewCreateLocalVolumeSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLocalVolumeSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateLocalVolumeSnapshotWithContextV2(ctx context.Context, request *CreateLocalVolumeSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLocalVolumeSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLocalVolumeSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeLocalVolumeSnapshotsRequest() (request *DescribeLocalVolumeSnapshotsRequest) {
	request = &DescribeLocalVolumeSnapshotsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeLocalVolumeSnapshots")
	return
}

func NewDescribeLocalVolumeSnapshotsResponse() (response *DescribeLocalVolumeSnapshotsResponse) {
	response = &DescribeLocalVolumeSnapshotsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeLocalVolumeSnapshots(request *DescribeLocalVolumeSnapshotsRequest) string {
	return c.DescribeLocalVolumeSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeLocalVolumeSnapshotsSend(request *DescribeLocalVolumeSnapshotsRequest) (*DescribeLocalVolumeSnapshotsResponse, error) {
	statusCode, msg, err := c.DescribeLocalVolumeSnapshotsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeLocalVolumeSnapshotsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeLocalVolumeSnapshotsWithContext(ctx context.Context, request *DescribeLocalVolumeSnapshotsRequest) string {
	if request == nil {
		request = NewDescribeLocalVolumeSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLocalVolumeSnapshotsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeLocalVolumeSnapshotsWithContextV2(ctx context.Context, request *DescribeLocalVolumeSnapshotsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeLocalVolumeSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeLocalVolumeSnapshotsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRollbackLocalVolumeRequest() (request *RollbackLocalVolumeRequest) {
	request = &RollbackLocalVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RollbackLocalVolume")
	return
}

func NewRollbackLocalVolumeResponse() (response *RollbackLocalVolumeResponse) {
	response = &RollbackLocalVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RollbackLocalVolume(request *RollbackLocalVolumeRequest) string {
	return c.RollbackLocalVolumeWithContext(context.Background(), request)
}

func (c *Client) RollbackLocalVolumeSend(request *RollbackLocalVolumeRequest) (*RollbackLocalVolumeResponse, error) {
	statusCode, msg, err := c.RollbackLocalVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RollbackLocalVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RollbackLocalVolumeWithContext(ctx context.Context, request *RollbackLocalVolumeRequest) string {
	if request == nil {
		request = NewRollbackLocalVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackLocalVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RollbackLocalVolumeWithContextV2(ctx context.Context, request *RollbackLocalVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewRollbackLocalVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackLocalVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteLocalVolumeSnapshotRequest() (request *DeleteLocalVolumeSnapshotRequest) {
	request = &DeleteLocalVolumeSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteLocalVolumeSnapshot")
	return
}

func NewDeleteLocalVolumeSnapshotResponse() (response *DeleteLocalVolumeSnapshotResponse) {
	response = &DeleteLocalVolumeSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLocalVolumeSnapshot(request *DeleteLocalVolumeSnapshotRequest) string {
	return c.DeleteLocalVolumeSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteLocalVolumeSnapshotSend(request *DeleteLocalVolumeSnapshotRequest) (*DeleteLocalVolumeSnapshotResponse, error) {
	statusCode, msg, err := c.DeleteLocalVolumeSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteLocalVolumeSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteLocalVolumeSnapshotWithContext(ctx context.Context, request *DeleteLocalVolumeSnapshotRequest) string {
	if request == nil {
		request = NewDeleteLocalVolumeSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLocalVolumeSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteLocalVolumeSnapshotWithContextV2(ctx context.Context, request *DeleteLocalVolumeSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLocalVolumeSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLocalVolumeSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDataGuardGroupsRequest() (request *ModifyDataGuardGroupsRequest) {
	request = &ModifyDataGuardGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyDataGuardGroups")
	return
}

func NewModifyDataGuardGroupsResponse() (response *ModifyDataGuardGroupsResponse) {
	response = &ModifyDataGuardGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDataGuardGroups(request *ModifyDataGuardGroupsRequest) string {
	return c.ModifyDataGuardGroupsWithContext(context.Background(), request)
}

func (c *Client) ModifyDataGuardGroupsSend(request *ModifyDataGuardGroupsRequest) (*ModifyDataGuardGroupsResponse, error) {
	statusCode, msg, err := c.ModifyDataGuardGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyDataGuardGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDataGuardGroupsWithContext(ctx context.Context, request *ModifyDataGuardGroupsRequest) string {
	if request == nil {
		request = NewModifyDataGuardGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDataGuardGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDataGuardGroupsWithContextV2(ctx context.Context, request *ModifyDataGuardGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDataGuardGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDataGuardGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataGuardCapacityRequest() (request *DescribeDataGuardCapacityRequest) {
	request = &DescribeDataGuardCapacityRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDataGuardCapacity")
	return
}

func NewDescribeDataGuardCapacityResponse() (response *DescribeDataGuardCapacityResponse) {
	response = &DescribeDataGuardCapacityResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataGuardCapacity(request *DescribeDataGuardCapacityRequest) string {
	return c.DescribeDataGuardCapacityWithContext(context.Background(), request)
}

func (c *Client) DescribeDataGuardCapacitySend(request *DescribeDataGuardCapacityRequest) (*DescribeDataGuardCapacityResponse, error) {
	statusCode, msg, err := c.DescribeDataGuardCapacityWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDataGuardCapacityResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataGuardCapacityWithContext(ctx context.Context, request *DescribeDataGuardCapacityRequest) string {
	if request == nil {
		request = NewDescribeDataGuardCapacityRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataGuardCapacityResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataGuardCapacityWithContextV2(ctx context.Context, request *DescribeDataGuardCapacityRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataGuardCapacityRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataGuardCapacityResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDataGuardGroupRequest() (request *CreateDataGuardGroupRequest) {
	request = &CreateDataGuardGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateDataGuardGroup")
	return
}

func NewCreateDataGuardGroupResponse() (response *CreateDataGuardGroupResponse) {
	response = &CreateDataGuardGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataGuardGroup(request *CreateDataGuardGroupRequest) string {
	return c.CreateDataGuardGroupWithContext(context.Background(), request)
}

func (c *Client) CreateDataGuardGroupSend(request *CreateDataGuardGroupRequest) (*CreateDataGuardGroupResponse, error) {
	statusCode, msg, err := c.CreateDataGuardGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDataGuardGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDataGuardGroupWithContext(ctx context.Context, request *CreateDataGuardGroupRequest) string {
	if request == nil {
		request = NewCreateDataGuardGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDataGuardGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDataGuardGroupWithContextV2(ctx context.Context, request *CreateDataGuardGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDataGuardGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDataGuardGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDataGuardGroupsRequest() (request *DeleteDataGuardGroupsRequest) {
	request = &DeleteDataGuardGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteDataGuardGroups")
	return
}

func NewDeleteDataGuardGroupsResponse() (response *DeleteDataGuardGroupsResponse) {
	response = &DeleteDataGuardGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDataGuardGroups(request *DeleteDataGuardGroupsRequest) string {
	return c.DeleteDataGuardGroupsWithContext(context.Background(), request)
}

func (c *Client) DeleteDataGuardGroupsSend(request *DeleteDataGuardGroupsRequest) (*DeleteDataGuardGroupsResponse, error) {
	statusCode, msg, err := c.DeleteDataGuardGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteDataGuardGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDataGuardGroupsWithContext(ctx context.Context, request *DeleteDataGuardGroupsRequest) string {
	if request == nil {
		request = NewDeleteDataGuardGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataGuardGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDataGuardGroupsWithContextV2(ctx context.Context, request *DeleteDataGuardGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDataGuardGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataGuardGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataGuardGroupRequest() (request *DescribeDataGuardGroupRequest) {
	request = &DescribeDataGuardGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDataGuardGroup")
	return
}

func NewDescribeDataGuardGroupResponse() (response *DescribeDataGuardGroupResponse) {
	response = &DescribeDataGuardGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataGuardGroup(request *DescribeDataGuardGroupRequest) string {
	return c.DescribeDataGuardGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeDataGuardGroupSend(request *DescribeDataGuardGroupRequest) (*DescribeDataGuardGroupResponse, error) {
	statusCode, msg, err := c.DescribeDataGuardGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDataGuardGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataGuardGroupWithContext(ctx context.Context, request *DescribeDataGuardGroupRequest) string {
	if request == nil {
		request = NewDescribeDataGuardGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataGuardGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataGuardGroupWithContextV2(ctx context.Context, request *DescribeDataGuardGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataGuardGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataGuardGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRemoveVmFromDataGuardRequest() (request *RemoveVmFromDataGuardRequest) {
	request = &RemoveVmFromDataGuardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "RemoveVmFromDataGuard")
	return
}

func NewRemoveVmFromDataGuardResponse() (response *RemoveVmFromDataGuardResponse) {
	response = &RemoveVmFromDataGuardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemoveVmFromDataGuard(request *RemoveVmFromDataGuardRequest) string {
	return c.RemoveVmFromDataGuardWithContext(context.Background(), request)
}

func (c *Client) RemoveVmFromDataGuardSend(request *RemoveVmFromDataGuardRequest) (*RemoveVmFromDataGuardResponse, error) {
	statusCode, msg, err := c.RemoveVmFromDataGuardWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RemoveVmFromDataGuardResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RemoveVmFromDataGuardWithContext(ctx context.Context, request *RemoveVmFromDataGuardRequest) string {
	if request == nil {
		request = NewRemoveVmFromDataGuardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemoveVmFromDataGuardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RemoveVmFromDataGuardWithContextV2(ctx context.Context, request *RemoveVmFromDataGuardRequest) (int, string, error) {
	if request == nil {
		request = NewRemoveVmFromDataGuardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRemoveVmFromDataGuardResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDedicatedHostsRequest() (request *CreateDedicatedHostsRequest) {
	request = &CreateDedicatedHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateDedicatedHosts")
	return
}

func NewCreateDedicatedHostsResponse() (response *CreateDedicatedHostsResponse) {
	response = &CreateDedicatedHostsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDedicatedHosts(request *CreateDedicatedHostsRequest) string {
	return c.CreateDedicatedHostsWithContext(context.Background(), request)
}

func (c *Client) CreateDedicatedHostsSend(request *CreateDedicatedHostsRequest) (*CreateDedicatedHostsResponse, error) {
	statusCode, msg, err := c.CreateDedicatedHostsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDedicatedHostsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDedicatedHostsWithContext(ctx context.Context, request *CreateDedicatedHostsRequest) string {
	if request == nil {
		request = NewCreateDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDedicatedHostsWithContextV2(ctx context.Context, request *CreateDedicatedHostsRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedHostsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDedicatedHostRequest() (request *DeleteDedicatedHostRequest) {
	request = &DeleteDedicatedHostRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteDedicatedHost")
	return
}

func NewDeleteDedicatedHostResponse() (response *DeleteDedicatedHostResponse) {
	response = &DeleteDedicatedHostResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDedicatedHost(request *DeleteDedicatedHostRequest) string {
	return c.DeleteDedicatedHostWithContext(context.Background(), request)
}

func (c *Client) DeleteDedicatedHostSend(request *DeleteDedicatedHostRequest) (*DeleteDedicatedHostResponse, error) {
	statusCode, msg, err := c.DeleteDedicatedHostWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteDedicatedHostResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDedicatedHostWithContext(ctx context.Context, request *DeleteDedicatedHostRequest) string {
	if request == nil {
		request = NewDeleteDedicatedHostRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDedicatedHostResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDedicatedHostWithContextV2(ctx context.Context, request *DeleteDedicatedHostRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDedicatedHostRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDedicatedHostResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDedicatedHostsRequest() (request *DescribeDedicatedHostsRequest) {
	request = &DescribeDedicatedHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDedicatedHosts")
	return
}

func NewDescribeDedicatedHostsResponse() (response *DescribeDedicatedHostsResponse) {
	response = &DescribeDedicatedHostsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedHosts(request *DescribeDedicatedHostsRequest) string {
	return c.DescribeDedicatedHostsWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedHostsSend(request *DescribeDedicatedHostsRequest) (*DescribeDedicatedHostsResponse, error) {
	statusCode, msg, err := c.DescribeDedicatedHostsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDedicatedHostsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDedicatedHostsWithContext(ctx context.Context, request *DescribeDedicatedHostsRequest) string {
	if request == nil {
		request = NewDescribeDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDedicatedHostsWithContextV2(ctx context.Context, request *DescribeDedicatedHostsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedHostsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScalingConfigurationRequest() (request *DescribeScalingConfigurationRequest) {
	request = &DescribeScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingConfiguration")
	return
}

func NewDescribeScalingConfigurationResponse() (response *DescribeScalingConfigurationResponse) {
	response = &DescribeScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingConfiguration(request *DescribeScalingConfigurationRequest) string {
	return c.DescribeScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingConfigurationSend(request *DescribeScalingConfigurationRequest) (*DescribeScalingConfigurationResponse, error) {
	statusCode, msg, err := c.DescribeScalingConfigurationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScalingConfigurationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScalingConfigurationWithContext(ctx context.Context, request *DescribeScalingConfigurationRequest) string {
	if request == nil {
		request = NewDescribeScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScalingConfigurationWithContextV2(ctx context.Context, request *DescribeScalingConfigurationRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingConfigurationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateScalingConfigurationRequest() (request *CreateScalingConfigurationRequest) {
	request = &CreateScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingConfiguration")
	return
}

func NewCreateScalingConfigurationResponse() (response *CreateScalingConfigurationResponse) {
	response = &CreateScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) string {
	return c.CreateScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) CreateScalingConfigurationSend(request *CreateScalingConfigurationRequest) (*CreateScalingConfigurationResponse, error) {
	statusCode, msg, err := c.CreateScalingConfigurationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateScalingConfigurationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateScalingConfigurationWithContext(ctx context.Context, request *CreateScalingConfigurationRequest) string {
	if request == nil {
		request = NewCreateScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateScalingConfigurationWithContextV2(ctx context.Context, request *CreateScalingConfigurationRequest) (int, string, error) {
	if request == nil {
		request = NewCreateScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingConfigurationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteScalingConfigurationRequest() (request *DeleteScalingConfigurationRequest) {
	request = &DeleteScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScalingConfiguration")
	return
}

func NewDeleteScalingConfigurationResponse() (response *DeleteScalingConfigurationResponse) {
	response = &DeleteScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScalingConfiguration(request *DeleteScalingConfigurationRequest) string {
	return c.DeleteScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) DeleteScalingConfigurationSend(request *DeleteScalingConfigurationRequest) (*DeleteScalingConfigurationResponse, error) {
	statusCode, msg, err := c.DeleteScalingConfigurationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteScalingConfigurationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteScalingConfigurationWithContext(ctx context.Context, request *DeleteScalingConfigurationRequest) string {
	if request == nil {
		request = NewDeleteScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteScalingConfigurationWithContextV2(ctx context.Context, request *DeleteScalingConfigurationRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingConfigurationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateScalingGroupRequest() (request *CreateScalingGroupRequest) {
	request = &CreateScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingGroup")
	return
}

func NewCreateScalingGroupResponse() (response *CreateScalingGroupResponse) {
	response = &CreateScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingGroup(request *CreateScalingGroupRequest) string {
	return c.CreateScalingGroupWithContext(context.Background(), request)
}

func (c *Client) CreateScalingGroupSend(request *CreateScalingGroupRequest) (*CreateScalingGroupResponse, error) {
	statusCode, msg, err := c.CreateScalingGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateScalingGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateScalingGroupWithContext(ctx context.Context, request *CreateScalingGroupRequest) string {
	if request == nil {
		request = NewCreateScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateScalingGroupWithContextV2(ctx context.Context, request *CreateScalingGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScalingGroupRequest() (request *DescribeScalingGroupRequest) {
	request = &DescribeScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingGroup")
	return
}

func NewDescribeScalingGroupResponse() (response *DescribeScalingGroupResponse) {
	response = &DescribeScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingGroup(request *DescribeScalingGroupRequest) string {
	return c.DescribeScalingGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingGroupSend(request *DescribeScalingGroupRequest) (*DescribeScalingGroupResponse, error) {
	statusCode, msg, err := c.DescribeScalingGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScalingGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScalingGroupWithContext(ctx context.Context, request *DescribeScalingGroupRequest) string {
	if request == nil {
		request = NewDescribeScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScalingGroupWithContextV2(ctx context.Context, request *DescribeScalingGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyScalingGroupRequest() (request *ModifyScalingGroupRequest) {
	request = &ModifyScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingGroup")
	return
}

func NewModifyScalingGroupResponse() (response *ModifyScalingGroupResponse) {
	response = &ModifyScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) string {
	return c.ModifyScalingGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingGroupSend(request *ModifyScalingGroupRequest) (*ModifyScalingGroupResponse, error) {
	statusCode, msg, err := c.ModifyScalingGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyScalingGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyScalingGroupWithContext(ctx context.Context, request *ModifyScalingGroupRequest) string {
	if request == nil {
		request = NewModifyScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyScalingGroupWithContextV2(ctx context.Context, request *ModifyScalingGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetKvmProtectedDetachRequest() (request *SetKvmProtectedDetachRequest) {
	request = &SetKvmProtectedDetachRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SetKvmProtectedDetach")
	return
}

func NewSetKvmProtectedDetachResponse() (response *SetKvmProtectedDetachResponse) {
	response = &SetKvmProtectedDetachResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetKvmProtectedDetach(request *SetKvmProtectedDetachRequest) string {
	return c.SetKvmProtectedDetachWithContext(context.Background(), request)
}

func (c *Client) SetKvmProtectedDetachSend(request *SetKvmProtectedDetachRequest) (*SetKvmProtectedDetachResponse, error) {
	statusCode, msg, err := c.SetKvmProtectedDetachWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetKvmProtectedDetachResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetKvmProtectedDetachWithContext(ctx context.Context, request *SetKvmProtectedDetachRequest) string {
	if request == nil {
		request = NewSetKvmProtectedDetachRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetKvmProtectedDetachResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetKvmProtectedDetachWithContextV2(ctx context.Context, request *SetKvmProtectedDetachRequest) (int, string, error) {
	if request == nil {
		request = NewSetKvmProtectedDetachRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetKvmProtectedDetachResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScalingInstanceRequest() (request *DescribeScalingInstanceRequest) {
	request = &DescribeScalingInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingInstance")
	return
}

func NewDescribeScalingInstanceResponse() (response *DescribeScalingInstanceResponse) {
	response = &DescribeScalingInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingInstance(request *DescribeScalingInstanceRequest) string {
	return c.DescribeScalingInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingInstanceSend(request *DescribeScalingInstanceRequest) (*DescribeScalingInstanceResponse, error) {
	statusCode, msg, err := c.DescribeScalingInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScalingInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScalingInstanceWithContext(ctx context.Context, request *DescribeScalingInstanceRequest) string {
	if request == nil {
		request = NewDescribeScalingInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScalingInstanceWithContextV2(ctx context.Context, request *DescribeScalingInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScalingInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachInstanceRequest() (request *AttachInstanceRequest) {
	request = &AttachInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachInstance")
	return
}

func NewAttachInstanceResponse() (response *AttachInstanceResponse) {
	response = &AttachInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachInstance(request *AttachInstanceRequest) string {
	return c.AttachInstanceWithContext(context.Background(), request)
}

func (c *Client) AttachInstanceSend(request *AttachInstanceRequest) (*AttachInstanceResponse, error) {
	statusCode, msg, err := c.AttachInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AttachInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachInstanceWithContext(ctx context.Context, request *AttachInstanceRequest) string {
	if request == nil {
		request = NewAttachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachInstanceWithContextV2(ctx context.Context, request *AttachInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewAttachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachInstanceRequest() (request *DetachInstanceRequest) {
	request = &DetachInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachInstance")
	return
}

func NewDetachInstanceResponse() (response *DetachInstanceResponse) {
	response = &DetachInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachInstance(request *DetachInstanceRequest) string {
	return c.DetachInstanceWithContext(context.Background(), request)
}

func (c *Client) DetachInstanceSend(request *DetachInstanceRequest) (*DetachInstanceResponse, error) {
	statusCode, msg, err := c.DetachInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachInstanceWithContext(ctx context.Context, request *DetachInstanceRequest) string {
	if request == nil {
		request = NewDetachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachInstanceWithContextV2(ctx context.Context, request *DetachInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDetachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScalingActivityRequest() (request *DescribeScalingActivityRequest) {
	request = &DescribeScalingActivityRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingActivity")
	return
}

func NewDescribeScalingActivityResponse() (response *DescribeScalingActivityResponse) {
	response = &DescribeScalingActivityResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingActivity(request *DescribeScalingActivityRequest) string {
	return c.DescribeScalingActivityWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingActivitySend(request *DescribeScalingActivityRequest) (*DescribeScalingActivityResponse, error) {
	statusCode, msg, err := c.DescribeScalingActivityWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScalingActivityResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScalingActivityWithContext(ctx context.Context, request *DescribeScalingActivityRequest) string {
	if request == nil {
		request = NewDescribeScalingActivityRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingActivityResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScalingActivityWithContextV2(ctx context.Context, request *DescribeScalingActivityRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScalingActivityRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingActivityResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteScalingGroupRequest() (request *DeleteScalingGroupRequest) {
	request = &DeleteScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScalingGroup")
	return
}

func NewDeleteScalingGroupResponse() (response *DeleteScalingGroupResponse) {
	response = &DeleteScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScalingGroup(request *DeleteScalingGroupRequest) string {
	return c.DeleteScalingGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteScalingGroupSend(request *DeleteScalingGroupRequest) (*DeleteScalingGroupResponse, error) {
	statusCode, msg, err := c.DeleteScalingGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteScalingGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteScalingGroupWithContext(ctx context.Context, request *DeleteScalingGroupRequest) string {
	if request == nil {
		request = NewDeleteScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteScalingGroupWithContextV2(ctx context.Context, request *DeleteScalingGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableScalingGroupRequest() (request *DisableScalingGroupRequest) {
	request = &DisableScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DisableScalingGroup")
	return
}

func NewDisableScalingGroupResponse() (response *DisableScalingGroupResponse) {
	response = &DisableScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableScalingGroup(request *DisableScalingGroupRequest) string {
	return c.DisableScalingGroupWithContext(context.Background(), request)
}

func (c *Client) DisableScalingGroupSend(request *DisableScalingGroupRequest) (*DisableScalingGroupResponse, error) {
	statusCode, msg, err := c.DisableScalingGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DisableScalingGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableScalingGroupWithContext(ctx context.Context, request *DisableScalingGroupRequest) string {
	if request == nil {
		request = NewDisableScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableScalingGroupWithContextV2(ctx context.Context, request *DisableScalingGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDisableScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableScalingGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEnableScalingGroupRequest() (request *EnableScalingGroupRequest) {
	request = &EnableScalingGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "EnableScalingGroup")
	return
}

func NewEnableScalingGroupResponse() (response *EnableScalingGroupResponse) {
	response = &EnableScalingGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableScalingGroup(request *EnableScalingGroupRequest) string {
	return c.EnableScalingGroupWithContext(context.Background(), request)
}

func (c *Client) EnableScalingGroupSend(request *EnableScalingGroupRequest) (*EnableScalingGroupResponse, error) {
	statusCode, msg, err := c.EnableScalingGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct EnableScalingGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EnableScalingGroupWithContext(ctx context.Context, request *EnableScalingGroupRequest) string {
	if request == nil {
		request = NewEnableScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableScalingGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EnableScalingGroupWithContextV2(ctx context.Context, request *EnableScalingGroupRequest) (int, string, error) {
	if request == nil {
		request = NewEnableScalingGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableScalingGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScalingNotificationRequest() (request *DescribeScalingNotificationRequest) {
	request = &DescribeScalingNotificationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingNotification")
	return
}

func NewDescribeScalingNotificationResponse() (response *DescribeScalingNotificationResponse) {
	response = &DescribeScalingNotificationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingNotification(request *DescribeScalingNotificationRequest) string {
	return c.DescribeScalingNotificationWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingNotificationSend(request *DescribeScalingNotificationRequest) (*DescribeScalingNotificationResponse, error) {
	statusCode, msg, err := c.DescribeScalingNotificationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScalingNotificationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScalingNotificationWithContext(ctx context.Context, request *DescribeScalingNotificationRequest) string {
	if request == nil {
		request = NewDescribeScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingNotificationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScalingNotificationWithContextV2(ctx context.Context, request *DescribeScalingNotificationRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingNotificationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateScalingNotificationRequest() (request *CreateScalingNotificationRequest) {
	request = &CreateScalingNotificationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingNotification")
	return
}

func NewCreateScalingNotificationResponse() (response *CreateScalingNotificationResponse) {
	response = &CreateScalingNotificationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingNotification(request *CreateScalingNotificationRequest) string {
	return c.CreateScalingNotificationWithContext(context.Background(), request)
}

func (c *Client) CreateScalingNotificationSend(request *CreateScalingNotificationRequest) (*CreateScalingNotificationResponse, error) {
	statusCode, msg, err := c.CreateScalingNotificationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateScalingNotificationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateScalingNotificationWithContext(ctx context.Context, request *CreateScalingNotificationRequest) string {
	if request == nil {
		request = NewCreateScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingNotificationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateScalingNotificationWithContextV2(ctx context.Context, request *CreateScalingNotificationRequest) (int, string, error) {
	if request == nil {
		request = NewCreateScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingNotificationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyScalingNotificationRequest() (request *ModifyScalingNotificationRequest) {
	request = &ModifyScalingNotificationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingNotification")
	return
}

func NewModifyScalingNotificationResponse() (response *ModifyScalingNotificationResponse) {
	response = &ModifyScalingNotificationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingNotification(request *ModifyScalingNotificationRequest) string {
	return c.ModifyScalingNotificationWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingNotificationSend(request *ModifyScalingNotificationRequest) (*ModifyScalingNotificationResponse, error) {
	statusCode, msg, err := c.ModifyScalingNotificationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyScalingNotificationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyScalingNotificationWithContext(ctx context.Context, request *ModifyScalingNotificationRequest) string {
	if request == nil {
		request = NewModifyScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingNotificationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyScalingNotificationWithContextV2(ctx context.Context, request *ModifyScalingNotificationRequest) (int, string, error) {
	if request == nil {
		request = NewModifyScalingNotificationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingNotificationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateScheduledTaskRequest() (request *CreateScheduledTaskRequest) {
	request = &CreateScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScheduledTask")
	return
}

func NewCreateScheduledTaskResponse() (response *CreateScheduledTaskResponse) {
	response = &CreateScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScheduledTask(request *CreateScheduledTaskRequest) string {
	return c.CreateScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) CreateScheduledTaskSend(request *CreateScheduledTaskRequest) (*CreateScheduledTaskResponse, error) {
	statusCode, msg, err := c.CreateScheduledTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateScheduledTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateScheduledTaskWithContext(ctx context.Context, request *CreateScheduledTaskRequest) string {
	if request == nil {
		request = NewCreateScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateScheduledTaskWithContextV2(ctx context.Context, request *CreateScheduledTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScheduledTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScheduledTaskRequest() (request *DescribeScheduledTaskRequest) {
	request = &DescribeScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScheduledTask")
	return
}

func NewDescribeScheduledTaskResponse() (response *DescribeScheduledTaskResponse) {
	response = &DescribeScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScheduledTask(request *DescribeScheduledTaskRequest) string {
	return c.DescribeScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) DescribeScheduledTaskSend(request *DescribeScheduledTaskRequest) (*DescribeScheduledTaskResponse, error) {
	statusCode, msg, err := c.DescribeScheduledTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScheduledTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScheduledTaskWithContext(ctx context.Context, request *DescribeScheduledTaskRequest) string {
	if request == nil {
		request = NewDescribeScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScheduledTaskWithContextV2(ctx context.Context, request *DescribeScheduledTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScheduledTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyScheduledTaskRequest() (request *ModifyScheduledTaskRequest) {
	request = &ModifyScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScheduledTask")
	return
}

func NewModifyScheduledTaskResponse() (response *ModifyScheduledTaskResponse) {
	response = &ModifyScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScheduledTask(request *ModifyScheduledTaskRequest) string {
	return c.ModifyScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) ModifyScheduledTaskSend(request *ModifyScheduledTaskRequest) (*ModifyScheduledTaskResponse, error) {
	statusCode, msg, err := c.ModifyScheduledTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyScheduledTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyScheduledTaskWithContext(ctx context.Context, request *ModifyScheduledTaskRequest) string {
	if request == nil {
		request = NewModifyScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyScheduledTaskWithContextV2(ctx context.Context, request *ModifyScheduledTaskRequest) (int, string, error) {
	if request == nil {
		request = NewModifyScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScheduledTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteScheduledTaskRequest() (request *DeleteScheduledTaskRequest) {
	request = &DeleteScheduledTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScheduledTask")
	return
}

func NewDeleteScheduledTaskResponse() (response *DeleteScheduledTaskResponse) {
	response = &DeleteScheduledTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScheduledTask(request *DeleteScheduledTaskRequest) string {
	return c.DeleteScheduledTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteScheduledTaskSend(request *DeleteScheduledTaskRequest) (*DeleteScheduledTaskResponse, error) {
	statusCode, msg, err := c.DeleteScheduledTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteScheduledTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteScheduledTaskWithContext(ctx context.Context, request *DeleteScheduledTaskRequest) string {
	if request == nil {
		request = NewDeleteScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScheduledTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteScheduledTaskWithContextV2(ctx context.Context, request *DeleteScheduledTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteScheduledTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScheduledTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateScalingPolicyRequest() (request *CreateScalingPolicyRequest) {
	request = &CreateScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateScalingPolicy")
	return
}

func NewCreateScalingPolicyResponse() (response *CreateScalingPolicyResponse) {
	response = &CreateScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateScalingPolicy(request *CreateScalingPolicyRequest) string {
	return c.CreateScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) CreateScalingPolicySend(request *CreateScalingPolicyRequest) (*CreateScalingPolicyResponse, error) {
	statusCode, msg, err := c.CreateScalingPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateScalingPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateScalingPolicyWithContext(ctx context.Context, request *CreateScalingPolicyRequest) string {
	if request == nil {
		request = NewCreateScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateScalingPolicyWithContextV2(ctx context.Context, request *CreateScalingPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateScalingPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeScalingPolicyRequest() (request *DescribeScalingPolicyRequest) {
	request = &DescribeScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeScalingPolicy")
	return
}

func NewDescribeScalingPolicyResponse() (response *DescribeScalingPolicyResponse) {
	response = &DescribeScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeScalingPolicy(request *DescribeScalingPolicyRequest) string {
	return c.DescribeScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) DescribeScalingPolicySend(request *DescribeScalingPolicyRequest) (*DescribeScalingPolicyResponse, error) {
	statusCode, msg, err := c.DescribeScalingPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeScalingPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeScalingPolicyWithContext(ctx context.Context, request *DescribeScalingPolicyRequest) string {
	if request == nil {
		request = NewDescribeScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeScalingPolicyWithContextV2(ctx context.Context, request *DescribeScalingPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeScalingPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyScalingPolicyRequest() (request *ModifyScalingPolicyRequest) {
	request = &ModifyScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingPolicy")
	return
}

func NewModifyScalingPolicyResponse() (response *ModifyScalingPolicyResponse) {
	response = &ModifyScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingPolicy(request *ModifyScalingPolicyRequest) string {
	return c.ModifyScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingPolicySend(request *ModifyScalingPolicyRequest) (*ModifyScalingPolicyResponse, error) {
	statusCode, msg, err := c.ModifyScalingPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyScalingPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyScalingPolicyWithContext(ctx context.Context, request *ModifyScalingPolicyRequest) string {
	if request == nil {
		request = NewModifyScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyScalingPolicyWithContextV2(ctx context.Context, request *ModifyScalingPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteScalingPolicyRequest() (request *DeleteScalingPolicyRequest) {
	request = &DeleteScalingPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteScalingPolicy")
	return
}

func NewDeleteScalingPolicyResponse() (response *DeleteScalingPolicyResponse) {
	response = &DeleteScalingPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) string {
	return c.DeleteScalingPolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteScalingPolicySend(request *DeleteScalingPolicyRequest) (*DeleteScalingPolicyResponse, error) {
	statusCode, msg, err := c.DeleteScalingPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteScalingPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteScalingPolicyWithContext(ctx context.Context, request *DeleteScalingPolicyRequest) string {
	if request == nil {
		request = NewDeleteScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteScalingPolicyWithContextV2(ctx context.Context, request *DeleteScalingPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteScalingPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteScalingPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewImportImageRequest() (request *ImportImageRequest) {
	request = &ImportImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ImportImage")
	return
}

func NewImportImageResponse() (response *ImportImageResponse) {
	response = &ImportImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportImage(request *ImportImageRequest) string {
	return c.ImportImageWithContext(context.Background(), request)
}

func (c *Client) ImportImageSend(request *ImportImageRequest) (*ImportImageResponse, error) {
	statusCode, msg, err := c.ImportImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ImportImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ImportImageWithContext(ctx context.Context, request *ImportImageRequest) string {
	if request == nil {
		request = NewImportImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewImportImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ImportImageWithContextV2(ctx context.Context, request *ImportImageRequest) (int, string, error) {
	if request == nil {
		request = NewImportImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewImportImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCopyImageRequest() (request *CopyImageRequest) {
	request = &CopyImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CopyImage")
	return
}

func NewCopyImageResponse() (response *CopyImageResponse) {
	response = &CopyImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CopyImage(request *CopyImageRequest) string {
	return c.CopyImageWithContext(context.Background(), request)
}

func (c *Client) CopyImageSend(request *CopyImageRequest) (*CopyImageResponse, error) {
	statusCode, msg, err := c.CopyImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CopyImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CopyImageWithContext(ctx context.Context, request *CopyImageRequest) string {
	if request == nil {
		request = NewCopyImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopyImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CopyImageWithContextV2(ctx context.Context, request *CopyImageRequest) (int, string, error) {
	if request == nil {
		request = NewCopyImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopyImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
	request = &ModifyImageSharePermissionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyImageSharePermission")
	return
}

func NewModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
	response = &ModifyImageSharePermissionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) string {
	return c.ModifyImageSharePermissionWithContext(context.Background(), request)
}

func (c *Client) ModifyImageSharePermissionSend(request *ModifyImageSharePermissionRequest) (*ModifyImageSharePermissionResponse, error) {
	statusCode, msg, err := c.ModifyImageSharePermissionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyImageSharePermissionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyImageSharePermissionWithContext(ctx context.Context, request *ModifyImageSharePermissionRequest) string {
	if request == nil {
		request = NewModifyImageSharePermissionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageSharePermissionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyImageSharePermissionWithContextV2(ctx context.Context, request *ModifyImageSharePermissionRequest) (int, string, error) {
	if request == nil {
		request = NewModifyImageSharePermissionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageSharePermissionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImageSharePermissionRequest() (request *DescribeImageSharePermissionRequest) {
	request = &DescribeImageSharePermissionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeImageSharePermission")
	return
}

func NewDescribeImageSharePermissionResponse() (response *DescribeImageSharePermissionResponse) {
	response = &DescribeImageSharePermissionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) string {
	return c.DescribeImageSharePermissionWithContext(context.Background(), request)
}

func (c *Client) DescribeImageSharePermissionSend(request *DescribeImageSharePermissionRequest) (*DescribeImageSharePermissionResponse, error) {
	statusCode, msg, err := c.DescribeImageSharePermissionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImageSharePermissionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImageSharePermissionWithContext(ctx context.Context, request *DescribeImageSharePermissionRequest) string {
	if request == nil {
		request = NewDescribeImageSharePermissionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageSharePermissionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImageSharePermissionWithContextV2(ctx context.Context, request *DescribeImageSharePermissionRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImageSharePermissionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageSharePermissionResponse()
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
	request.Init().WithApiInfo("kec", APIVersion, "DescribeRegions")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRegionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachKeyRequest() (request *AttachKeyRequest) {
	request = &AttachKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachKey")
	return
}

func NewAttachKeyResponse() (response *AttachKeyResponse) {
	response = &AttachKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachKey(request *AttachKeyRequest) string {
	return c.AttachKeyWithContext(context.Background(), request)
}

func (c *Client) AttachKeySend(request *AttachKeyRequest) (*AttachKeyResponse, error) {
	statusCode, msg, err := c.AttachKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AttachKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachKeyWithContext(ctx context.Context, request *AttachKeyRequest) string {
	if request == nil {
		request = NewAttachKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachKeyWithContextV2(ctx context.Context, request *AttachKeyRequest) (int, string, error) {
	if request == nil {
		request = NewAttachKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachKeyRequest() (request *DetachKeyRequest) {
	request = &DetachKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachKey")
	return
}

func NewDetachKeyResponse() (response *DetachKeyResponse) {
	response = &DetachKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachKey(request *DetachKeyRequest) string {
	return c.DetachKeyWithContext(context.Background(), request)
}

func (c *Client) DetachKeySend(request *DetachKeyRequest) (*DetachKeyResponse, error) {
	statusCode, msg, err := c.DetachKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachKeyWithContext(ctx context.Context, request *DetachKeyRequest) string {
	if request == nil {
		request = NewDetachKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachKeyWithContextV2(ctx context.Context, request *DetachKeyRequest) (int, string, error) {
	if request == nil {
		request = NewDetachKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAvailabilityZonesRequest() (request *DescribeAvailabilityZonesRequest) {
	request = &DescribeAvailabilityZonesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeAvailabilityZones")
	return
}

func NewDescribeAvailabilityZonesResponse() (response *DescribeAvailabilityZonesResponse) {
	response = &DescribeAvailabilityZonesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAvailabilityZones(request *DescribeAvailabilityZonesRequest) string {
	return c.DescribeAvailabilityZonesWithContext(context.Background(), request)
}

func (c *Client) DescribeAvailabilityZonesSend(request *DescribeAvailabilityZonesRequest) (*DescribeAvailabilityZonesResponse, error) {
	statusCode, msg, err := c.DescribeAvailabilityZonesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAvailabilityZonesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, request *DescribeAvailabilityZonesRequest) string {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAvailabilityZonesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAvailabilityZonesWithContextV2(ctx context.Context, request *DescribeAvailabilityZonesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAvailabilityZonesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
	request = &DescribeInstanceTypeConfigsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstanceTypeConfigs")
	return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
	response = &DescribeInstanceTypeConfigsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) string {
	return c.DescribeInstanceTypeConfigsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceTypeConfigsSend(request *DescribeInstanceTypeConfigsRequest) (*DescribeInstanceTypeConfigsResponse, error) {
	statusCode, msg, err := c.DescribeInstanceTypeConfigsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceTypeConfigsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceTypeConfigsWithContext(ctx context.Context, request *DescribeInstanceTypeConfigsRequest) string {
	if request == nil {
		request = NewDescribeInstanceTypeConfigsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceTypeConfigsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceTypeConfigsWithContextV2(ctx context.Context, request *DescribeInstanceTypeConfigsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceTypeConfigsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceTypeConfigsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceFamilysRequest() (request *DescribeInstanceFamilysRequest) {
	request = &DescribeInstanceFamilysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstanceFamilys")
	return
}

func NewDescribeInstanceFamilysResponse() (response *DescribeInstanceFamilysResponse) {
	response = &DescribeInstanceFamilysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceFamilys(request *DescribeInstanceFamilysRequest) string {
	return c.DescribeInstanceFamilysWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceFamilysSend(request *DescribeInstanceFamilysRequest) (*DescribeInstanceFamilysResponse, error) {
	statusCode, msg, err := c.DescribeInstanceFamilysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceFamilysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceFamilysWithContext(ctx context.Context, request *DescribeInstanceFamilysRequest) string {
	if request == nil {
		request = NewDescribeInstanceFamilysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceFamilysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceFamilysWithContextV2(ctx context.Context, request *DescribeInstanceFamilysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceFamilysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceFamilysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddVmIntoDataGuardRequest() (request *AddVmIntoDataGuardRequest) {
	request = &AddVmIntoDataGuardRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AddVmIntoDataGuard")
	return
}

func NewAddVmIntoDataGuardResponse() (response *AddVmIntoDataGuardResponse) {
	response = &AddVmIntoDataGuardResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddVmIntoDataGuard(request *AddVmIntoDataGuardRequest) string {
	return c.AddVmIntoDataGuardWithContext(context.Background(), request)
}

func (c *Client) AddVmIntoDataGuardSend(request *AddVmIntoDataGuardRequest) (*AddVmIntoDataGuardResponse, error) {
	statusCode, msg, err := c.AddVmIntoDataGuardWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddVmIntoDataGuardResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddVmIntoDataGuardWithContext(ctx context.Context, request *AddVmIntoDataGuardRequest) string {
	if request == nil {
		request = NewAddVmIntoDataGuardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddVmIntoDataGuardResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddVmIntoDataGuardWithContextV2(ctx context.Context, request *AddVmIntoDataGuardRequest) (int, string, error) {
	if request == nil {
		request = NewAddVmIntoDataGuardRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddVmIntoDataGuardResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
	request = &CreateFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateFileSystem")
	return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
	response = &CreateFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) string {
	return c.CreateFileSystemWithContext(context.Background(), request)
}

func (c *Client) CreateFileSystemSend(request *CreateFileSystemRequest) (*CreateFileSystemResponse, error) {
	statusCode, msg, err := c.CreateFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) string {
	if request == nil {
		request = NewCreateFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateFileSystemWithContextV2(ctx context.Context, request *CreateFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewCreateFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
	request = &DeleteFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteFileSystem")
	return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
	response = &DeleteFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) string {
	return c.DeleteFileSystemWithContext(context.Background(), request)
}

func (c *Client) DeleteFileSystemSend(request *DeleteFileSystemRequest) (*DeleteFileSystemResponse, error) {
	statusCode, msg, err := c.DeleteFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) string {
	if request == nil {
		request = NewDeleteFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFileSystemWithContextV2(ctx context.Context, request *DeleteFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
	request = &DescribeFileSystemsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeFileSystems")
	return
}

func NewDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
	response = &DescribeFileSystemsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) string {
	return c.DescribeFileSystemsWithContext(context.Background(), request)
}

func (c *Client) DescribeFileSystemsSend(request *DescribeFileSystemsRequest) (*DescribeFileSystemsResponse, error) {
	statusCode, msg, err := c.DescribeFileSystemsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeFileSystemsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest) string {
	if request == nil {
		request = NewDescribeFileSystemsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileSystemsWithContextV2(ctx context.Context, request *DescribeFileSystemsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileSystemsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyFileSystemRequest() (request *ModifyFileSystemRequest) {
	request = &ModifyFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyFileSystem")
	return
}

func NewModifyFileSystemResponse() (response *ModifyFileSystemResponse) {
	response = &ModifyFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyFileSystem(request *ModifyFileSystemRequest) string {
	return c.ModifyFileSystemWithContext(context.Background(), request)
}

func (c *Client) ModifyFileSystemSend(request *ModifyFileSystemRequest) (*ModifyFileSystemResponse, error) {
	statusCode, msg, err := c.ModifyFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyFileSystemWithContext(ctx context.Context, request *ModifyFileSystemRequest) string {
	if request == nil {
		request = NewModifyFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyFileSystemWithContextV2(ctx context.Context, request *ModifyFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewModifyFileSystemRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMountTargetRequest() (request *CreateMountTargetRequest) {
	request = &CreateMountTargetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateMountTarget")
	return
}

func NewCreateMountTargetResponse() (response *CreateMountTargetResponse) {
	response = &CreateMountTargetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMountTarget(request *CreateMountTargetRequest) string {
	return c.CreateMountTargetWithContext(context.Background(), request)
}

func (c *Client) CreateMountTargetSend(request *CreateMountTargetRequest) (*CreateMountTargetResponse, error) {
	statusCode, msg, err := c.CreateMountTargetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateMountTargetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMountTargetWithContext(ctx context.Context, request *CreateMountTargetRequest) string {
	if request == nil {
		request = NewCreateMountTargetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMountTargetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMountTargetWithContextV2(ctx context.Context, request *CreateMountTargetRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMountTargetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMountTargetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMountTargetRequest() (request *DeleteMountTargetRequest) {
	request = &DeleteMountTargetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteMountTarget")
	return
}

func NewDeleteMountTargetResponse() (response *DeleteMountTargetResponse) {
	response = &DeleteMountTargetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMountTarget(request *DeleteMountTargetRequest) string {
	return c.DeleteMountTargetWithContext(context.Background(), request)
}

func (c *Client) DeleteMountTargetSend(request *DeleteMountTargetRequest) (*DeleteMountTargetResponse, error) {
	statusCode, msg, err := c.DeleteMountTargetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteMountTargetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest) string {
	if request == nil {
		request = NewDeleteMountTargetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMountTargetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMountTargetWithContextV2(ctx context.Context, request *DeleteMountTargetRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMountTargetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMountTargetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
	request = &DescribeMountTargetsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeMountTargets")
	return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
	response = &DescribeMountTargetsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) string {
	return c.DescribeMountTargetsWithContext(context.Background(), request)
}

func (c *Client) DescribeMountTargetsSend(request *DescribeMountTargetsRequest) (*DescribeMountTargetsResponse, error) {
	statusCode, msg, err := c.DescribeMountTargetsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMountTargetsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest) string {
	if request == nil {
		request = NewDescribeMountTargetsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMountTargetsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMountTargetsWithContextV2(ctx context.Context, request *DescribeMountTargetsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMountTargetsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMountTargetsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateModelRequest() (request *CreateModelRequest) {
	request = &CreateModelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateModel")
	return
}

func NewCreateModelResponse() (response *CreateModelResponse) {
	response = &CreateModelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateModel(request *CreateModelRequest) string {
	return c.CreateModelWithContext(context.Background(), request)
}

func (c *Client) CreateModelSend(request *CreateModelRequest) (*CreateModelResponse, error) {
	statusCode, msg, err := c.CreateModelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateModelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateModelWithContext(ctx context.Context, request *CreateModelRequest) string {
	if request == nil {
		request = NewCreateModelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateModelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateModelWithContextV2(ctx context.Context, request *CreateModelRequest) (int, string, error) {
	if request == nil {
		request = NewCreateModelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateModelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewTerminateModelsRequest() (request *TerminateModelsRequest) {
	request = &TerminateModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "TerminateModels")
	return
}

func NewTerminateModelsResponse() (response *TerminateModelsResponse) {
	response = &TerminateModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TerminateModels(request *TerminateModelsRequest) string {
	return c.TerminateModelsWithContext(context.Background(), request)
}

func (c *Client) TerminateModelsSend(request *TerminateModelsRequest) (*TerminateModelsResponse, error) {
	statusCode, msg, err := c.TerminateModelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct TerminateModelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) TerminateModelsWithContext(ctx context.Context, request *TerminateModelsRequest) string {
	if request == nil {
		request = NewTerminateModelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTerminateModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) TerminateModelsWithContextV2(ctx context.Context, request *TerminateModelsRequest) (int, string, error) {
	if request == nil {
		request = NewTerminateModelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTerminateModelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModelsRequest() (request *DescribeModelsRequest) {
	request = &DescribeModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeModels")
	return
}

func NewDescribeModelsResponse() (response *DescribeModelsResponse) {
	response = &DescribeModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModels(request *DescribeModelsRequest) string {
	return c.DescribeModelsWithContext(context.Background(), request)
}

func (c *Client) DescribeModelsSend(request *DescribeModelsRequest) (*DescribeModelsResponse, error) {
	statusCode, msg, err := c.DescribeModelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeModelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModelsWithContext(ctx context.Context, request *DescribeModelsRequest) string {
	if request == nil {
		request = NewDescribeModelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModelsWithContextV2(ctx context.Context, request *DescribeModelsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModelsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDedicatedClusterRequest() (request *DescribeDedicatedClusterRequest) {
	request = &DescribeDedicatedClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeDedicatedCluster")
	return
}

func NewDescribeDedicatedClusterResponse() (response *DescribeDedicatedClusterResponse) {
	response = &DescribeDedicatedClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedCluster(request *DescribeDedicatedClusterRequest) string {
	return c.DescribeDedicatedClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedClusterSend(request *DescribeDedicatedClusterRequest) (*DescribeDedicatedClusterResponse, error) {
	statusCode, msg, err := c.DescribeDedicatedClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDedicatedClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDedicatedClusterWithContext(ctx context.Context, request *DescribeDedicatedClusterRequest) string {
	if request == nil {
		request = NewDescribeDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDedicatedClusterWithContextV2(ctx context.Context, request *DescribeDedicatedClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDedicatedClusterRequest() (request *CreateDedicatedClusterRequest) {
	request = &CreateDedicatedClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CreateDedicatedCluster")
	return
}

func NewCreateDedicatedClusterResponse() (response *CreateDedicatedClusterResponse) {
	response = &CreateDedicatedClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDedicatedCluster(request *CreateDedicatedClusterRequest) string {
	return c.CreateDedicatedClusterWithContext(context.Background(), request)
}

func (c *Client) CreateDedicatedClusterSend(request *CreateDedicatedClusterRequest) (*CreateDedicatedClusterResponse, error) {
	statusCode, msg, err := c.CreateDedicatedClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDedicatedClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDedicatedClusterWithContext(ctx context.Context, request *CreateDedicatedClusterRequest) string {
	if request == nil {
		request = NewCreateDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDedicatedClusterWithContextV2(ctx context.Context, request *CreateDedicatedClusterRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDedicatedClusterRequest() (request *DeleteDedicatedClusterRequest) {
	request = &DeleteDedicatedClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DeleteDedicatedCluster")
	return
}

func NewDeleteDedicatedClusterResponse() (response *DeleteDedicatedClusterResponse) {
	response = &DeleteDedicatedClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDedicatedCluster(request *DeleteDedicatedClusterRequest) string {
	return c.DeleteDedicatedClusterWithContext(context.Background(), request)
}

func (c *Client) DeleteDedicatedClusterSend(request *DeleteDedicatedClusterRequest) (*DeleteDedicatedClusterResponse, error) {
	statusCode, msg, err := c.DeleteDedicatedClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteDedicatedClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDedicatedClusterWithContext(ctx context.Context, request *DeleteDedicatedClusterRequest) string {
	if request == nil {
		request = NewDeleteDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDedicatedClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDedicatedClusterWithContextV2(ctx context.Context, request *DeleteDedicatedClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDedicatedClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDedicatedClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetvCPURequest() (request *SetvCPURequest) {
	request = &SetvCPURequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SetvCPU")
	return
}

func NewSetvCPUResponse() (response *SetvCPUResponse) {
	response = &SetvCPUResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetvCPU(request *SetvCPURequest) string {
	return c.SetvCPUWithContext(context.Background(), request)
}

func (c *Client) SetvCPUSend(request *SetvCPURequest) (*SetvCPUResponse, error) {
	statusCode, msg, err := c.SetvCPUWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetvCPUResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetvCPUWithContext(ctx context.Context, request *SetvCPURequest) string {
	if request == nil {
		request = NewSetvCPURequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetvCPUResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetvCPUWithContextV2(ctx context.Context, request *SetvCPURequest) (int, string, error) {
	if request == nil {
		request = NewSetvCPURequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetvCPUResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDedicatedHostMigrateRequest() (request *DedicatedHostMigrateRequest) {
	request = &DedicatedHostMigrateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DedicatedHostMigrate")
	return
}

func NewDedicatedHostMigrateResponse() (response *DedicatedHostMigrateResponse) {
	response = &DedicatedHostMigrateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DedicatedHostMigrate(request *DedicatedHostMigrateRequest) string {
	return c.DedicatedHostMigrateWithContext(context.Background(), request)
}

func (c *Client) DedicatedHostMigrateSend(request *DedicatedHostMigrateRequest) (*DedicatedHostMigrateResponse, error) {
	statusCode, msg, err := c.DedicatedHostMigrateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DedicatedHostMigrateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DedicatedHostMigrateWithContext(ctx context.Context, request *DedicatedHostMigrateRequest) string {
	if request == nil {
		request = NewDedicatedHostMigrateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDedicatedHostMigrateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DedicatedHostMigrateWithContextV2(ctx context.Context, request *DedicatedHostMigrateRequest) (int, string, error) {
	if request == nil {
		request = NewDedicatedHostMigrateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDedicatedHostMigrateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDedicatedClusterNameRequest() (request *ModifyDedicatedClusterNameRequest) {
	request = &ModifyDedicatedClusterNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyDedicatedClusterName")
	return
}

func NewModifyDedicatedClusterNameResponse() (response *ModifyDedicatedClusterNameResponse) {
	response = &ModifyDedicatedClusterNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDedicatedClusterName(request *ModifyDedicatedClusterNameRequest) string {
	return c.ModifyDedicatedClusterNameWithContext(context.Background(), request)
}

func (c *Client) ModifyDedicatedClusterNameSend(request *ModifyDedicatedClusterNameRequest) (*ModifyDedicatedClusterNameResponse, error) {
	statusCode, msg, err := c.ModifyDedicatedClusterNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyDedicatedClusterNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDedicatedClusterNameWithContext(ctx context.Context, request *ModifyDedicatedClusterNameRequest) string {
	if request == nil {
		request = NewModifyDedicatedClusterNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDedicatedClusterNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDedicatedClusterNameWithContextV2(ctx context.Context, request *ModifyDedicatedClusterNameRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDedicatedClusterNameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDedicatedClusterNameResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewInstanceMigrateRequest() (request *InstanceMigrateRequest) {
	request = &InstanceMigrateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "InstanceMigrate")
	return
}

func NewInstanceMigrateResponse() (response *InstanceMigrateResponse) {
	response = &InstanceMigrateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) InstanceMigrate(request *InstanceMigrateRequest) string {
	return c.InstanceMigrateWithContext(context.Background(), request)
}

func (c *Client) InstanceMigrateSend(request *InstanceMigrateRequest) (*InstanceMigrateResponse, error) {
	statusCode, msg, err := c.InstanceMigrateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct InstanceMigrateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) InstanceMigrateWithContext(ctx context.Context, request *InstanceMigrateRequest) string {
	if request == nil {
		request = NewInstanceMigrateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstanceMigrateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) InstanceMigrateWithContextV2(ctx context.Context, request *InstanceMigrateRequest) (int, string, error) {
	if request == nil {
		request = NewInstanceMigrateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstanceMigrateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyInstanceAutoDeleteTimeRequest() (request *ModifyInstanceAutoDeleteTimeRequest) {
	request = &ModifyInstanceAutoDeleteTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyInstanceAutoDeleteTime")
	return
}

func NewModifyInstanceAutoDeleteTimeResponse() (response *ModifyInstanceAutoDeleteTimeResponse) {
	response = &ModifyInstanceAutoDeleteTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceAutoDeleteTime(request *ModifyInstanceAutoDeleteTimeRequest) string {
	return c.ModifyInstanceAutoDeleteTimeWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAutoDeleteTimeSend(request *ModifyInstanceAutoDeleteTimeRequest) (*ModifyInstanceAutoDeleteTimeResponse, error) {
	statusCode, msg, err := c.ModifyInstanceAutoDeleteTimeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyInstanceAutoDeleteTimeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyInstanceAutoDeleteTimeWithContext(ctx context.Context, request *ModifyInstanceAutoDeleteTimeRequest) string {
	if request == nil {
		request = NewModifyInstanceAutoDeleteTimeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAutoDeleteTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyInstanceAutoDeleteTimeWithContextV2(ctx context.Context, request *ModifyInstanceAutoDeleteTimeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceAutoDeleteTimeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceAutoDeleteTimeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyScalingConfigurationRequest() (request *ModifyScalingConfigurationRequest) {
	request = &ModifyScalingConfigurationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyScalingConfiguration")
	return
}

func NewModifyScalingConfigurationResponse() (response *ModifyScalingConfigurationResponse) {
	response = &ModifyScalingConfigurationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScalingConfiguration(request *ModifyScalingConfigurationRequest) string {
	return c.ModifyScalingConfigurationWithContext(context.Background(), request)
}

func (c *Client) ModifyScalingConfigurationSend(request *ModifyScalingConfigurationRequest) (*ModifyScalingConfigurationResponse, error) {
	statusCode, msg, err := c.ModifyScalingConfigurationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyScalingConfigurationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyScalingConfigurationWithContext(ctx context.Context, request *ModifyScalingConfigurationRequest) string {
	if request == nil {
		request = NewModifyScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingConfigurationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyScalingConfigurationWithContextV2(ctx context.Context, request *ModifyScalingConfigurationRequest) (int, string, error) {
	if request == nil {
		request = NewModifyScalingConfigurationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyScalingConfigurationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSpotPriceHistoryRequest() (request *DescribeSpotPriceHistoryRequest) {
	request = &DescribeSpotPriceHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeSpotPriceHistory")
	return
}

func NewDescribeSpotPriceHistoryResponse() (response *DescribeSpotPriceHistoryResponse) {
	response = &DescribeSpotPriceHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSpotPriceHistory(request *DescribeSpotPriceHistoryRequest) string {
	return c.DescribeSpotPriceHistoryWithContext(context.Background(), request)
}

func (c *Client) DescribeSpotPriceHistorySend(request *DescribeSpotPriceHistoryRequest) (*DescribeSpotPriceHistoryResponse, error) {
	statusCode, msg, err := c.DescribeSpotPriceHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSpotPriceHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSpotPriceHistoryWithContext(ctx context.Context, request *DescribeSpotPriceHistoryRequest) string {
	if request == nil {
		request = NewDescribeSpotPriceHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSpotPriceHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSpotPriceHistoryWithContextV2(ctx context.Context, request *DescribeSpotPriceHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSpotPriceHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSpotPriceHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrice(request *DescribePriceRequest) string {
	return c.DescribePriceWithContext(context.Background(), request)
}

func (c *Client) DescribePriceSend(request *DescribePriceRequest) (*DescribePriceResponse, error) {
	statusCode, msg, err := c.DescribePriceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePriceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest) string {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePriceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePriceWithContextV2(ctx context.Context, request *DescribePriceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePriceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEnableImageCachingRequest() (request *EnableImageCachingRequest) {
	request = &EnableImageCachingRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "EnableImageCaching")
	return
}

func NewEnableImageCachingResponse() (response *EnableImageCachingResponse) {
	response = &EnableImageCachingResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableImageCaching(request *EnableImageCachingRequest) string {
	return c.EnableImageCachingWithContext(context.Background(), request)
}

func (c *Client) EnableImageCachingSend(request *EnableImageCachingRequest) (*EnableImageCachingResponse, error) {
	statusCode, msg, err := c.EnableImageCachingWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct EnableImageCachingResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EnableImageCachingWithContext(ctx context.Context, request *EnableImageCachingRequest) string {
	if request == nil {
		request = NewEnableImageCachingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableImageCachingResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EnableImageCachingWithContextV2(ctx context.Context, request *EnableImageCachingRequest) (int, string, error) {
	if request == nil {
		request = NewEnableImageCachingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableImageCachingResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableImageCachingRequest() (request *DisableImageCachingRequest) {
	request = &DisableImageCachingRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DisableImageCaching")
	return
}

func NewDisableImageCachingResponse() (response *DisableImageCachingResponse) {
	response = &DisableImageCachingResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableImageCaching(request *DisableImageCachingRequest) string {
	return c.DisableImageCachingWithContext(context.Background(), request)
}

func (c *Client) DisableImageCachingSend(request *DisableImageCachingRequest) (*DisableImageCachingResponse, error) {
	statusCode, msg, err := c.DisableImageCachingWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DisableImageCachingResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableImageCachingWithContext(ctx context.Context, request *DisableImageCachingRequest) string {
	if request == nil {
		request = NewDisableImageCachingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableImageCachingResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableImageCachingWithContextV2(ctx context.Context, request *DisableImageCachingRequest) (int, string, error) {
	if request == nil {
		request = NewDisableImageCachingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisableImageCachingResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyLoadBalancersRequest() (request *ModifyLoadBalancersRequest) {
	request = &ModifyLoadBalancersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "ModifyLoadBalancers")
	return
}

func NewModifyLoadBalancersResponse() (response *ModifyLoadBalancersResponse) {
	response = &ModifyLoadBalancersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyLoadBalancers(request *ModifyLoadBalancersRequest) string {
	return c.ModifyLoadBalancersWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancersSend(request *ModifyLoadBalancersRequest) (*ModifyLoadBalancersResponse, error) {
	statusCode, msg, err := c.ModifyLoadBalancersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyLoadBalancersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyLoadBalancersWithContext(ctx context.Context, request *ModifyLoadBalancersRequest) string {
	if request == nil {
		request = NewModifyLoadBalancersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyLoadBalancersWithContextV2(ctx context.Context, request *ModifyLoadBalancersRequest) (int, string, error) {
	if request == nil {
		request = NewModifyLoadBalancersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyLoadBalancersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachInstancesIamRoleRequest() (request *AttachInstancesIamRoleRequest) {
	request = &AttachInstancesIamRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "AttachInstancesIamRole")
	return
}

func NewAttachInstancesIamRoleResponse() (response *AttachInstancesIamRoleResponse) {
	response = &AttachInstancesIamRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachInstancesIamRole(request *AttachInstancesIamRoleRequest) string {
	return c.AttachInstancesIamRoleWithContext(context.Background(), request)
}

func (c *Client) AttachInstancesIamRoleSend(request *AttachInstancesIamRoleRequest) (*AttachInstancesIamRoleResponse, error) {
	statusCode, msg, err := c.AttachInstancesIamRoleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AttachInstancesIamRoleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachInstancesIamRoleWithContext(ctx context.Context, request *AttachInstancesIamRoleRequest) string {
	if request == nil {
		request = NewAttachInstancesIamRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachInstancesIamRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachInstancesIamRoleWithContextV2(ctx context.Context, request *AttachInstancesIamRoleRequest) (int, string, error) {
	if request == nil {
		request = NewAttachInstancesIamRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachInstancesIamRoleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachInstancesIamRoleRequest() (request *DetachInstancesIamRoleRequest) {
	request = &DetachInstancesIamRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DetachInstancesIamRole")
	return
}

func NewDetachInstancesIamRoleResponse() (response *DetachInstancesIamRoleResponse) {
	response = &DetachInstancesIamRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachInstancesIamRole(request *DetachInstancesIamRoleRequest) string {
	return c.DetachInstancesIamRoleWithContext(context.Background(), request)
}

func (c *Client) DetachInstancesIamRoleSend(request *DetachInstancesIamRoleRequest) (*DetachInstancesIamRoleResponse, error) {
	statusCode, msg, err := c.DetachInstancesIamRoleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachInstancesIamRoleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachInstancesIamRoleWithContext(ctx context.Context, request *DetachInstancesIamRoleRequest) string {
	if request == nil {
		request = NewDetachInstancesIamRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachInstancesIamRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachInstancesIamRoleWithContextV2(ctx context.Context, request *DetachInstancesIamRoleRequest) (int, string, error) {
	if request == nil {
		request = NewDetachInstancesIamRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachInstancesIamRoleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCopySnapshotRequest() (request *CopySnapshotRequest) {
	request = &CopySnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CopySnapshot")
	return
}

func NewCopySnapshotResponse() (response *CopySnapshotResponse) {
	response = &CopySnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CopySnapshot(request *CopySnapshotRequest) string {
	return c.CopySnapshotWithContext(context.Background(), request)
}

func (c *Client) CopySnapshotSend(request *CopySnapshotRequest) (*CopySnapshotResponse, error) {
	statusCode, msg, err := c.CopySnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CopySnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CopySnapshotWithContext(ctx context.Context, request *CopySnapshotRequest) string {
	if request == nil {
		request = NewCopySnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopySnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CopySnapshotWithContextV2(ctx context.Context, request *CopySnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewCopySnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopySnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewPreMigrateInstanceRequest() (request *PreMigrateInstanceRequest) {
	request = &PreMigrateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "PreMigrateInstance")
	return
}

func NewPreMigrateInstanceResponse() (response *PreMigrateInstanceResponse) {
	response = &PreMigrateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) PreMigrateInstance(request *PreMigrateInstanceRequest) string {
	return c.PreMigrateInstanceWithContext(context.Background(), request)
}

func (c *Client) PreMigrateInstanceSend(request *PreMigrateInstanceRequest) (*PreMigrateInstanceResponse, error) {
	statusCode, msg, err := c.PreMigrateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct PreMigrateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) PreMigrateInstanceWithContext(ctx context.Context, request *PreMigrateInstanceRequest) string {
	if request == nil {
		request = NewPreMigrateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPreMigrateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) PreMigrateInstanceWithContextV2(ctx context.Context, request *PreMigrateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewPreMigrateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewPreMigrateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelPreMigrateInstanceRequest() (request *CancelPreMigrateInstanceRequest) {
	request = &CancelPreMigrateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "CancelPreMigrateInstance")
	return
}

func NewCancelPreMigrateInstanceResponse() (response *CancelPreMigrateInstanceResponse) {
	response = &CancelPreMigrateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelPreMigrateInstance(request *CancelPreMigrateInstanceRequest) string {
	return c.CancelPreMigrateInstanceWithContext(context.Background(), request)
}

func (c *Client) CancelPreMigrateInstanceSend(request *CancelPreMigrateInstanceRequest) (*CancelPreMigrateInstanceResponse, error) {
	statusCode, msg, err := c.CancelPreMigrateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelPreMigrateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelPreMigrateInstanceWithContext(ctx context.Context, request *CancelPreMigrateInstanceRequest) string {
	if request == nil {
		request = NewCancelPreMigrateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelPreMigrateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelPreMigrateInstanceWithContextV2(ctx context.Context, request *CancelPreMigrateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCancelPreMigrateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelPreMigrateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSwitchImageTypeRequest() (request *SwitchImageTypeRequest) {
	request = &SwitchImageTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SwitchImageType")
	return
}

func NewSwitchImageTypeResponse() (response *SwitchImageTypeResponse) {
	response = &SwitchImageTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SwitchImageType(request *SwitchImageTypeRequest) string {
	return c.SwitchImageTypeWithContext(context.Background(), request)
}

func (c *Client) SwitchImageTypeSend(request *SwitchImageTypeRequest) (*SwitchImageTypeResponse, error) {
	statusCode, msg, err := c.SwitchImageTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SwitchImageTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SwitchImageTypeWithContext(ctx context.Context, request *SwitchImageTypeRequest) string {
	if request == nil {
		request = NewSwitchImageTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSwitchImageTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SwitchImageTypeWithContextV2(ctx context.Context, request *SwitchImageTypeRequest) (int, string, error) {
	if request == nil {
		request = NewSwitchImageTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSwitchImageTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetInstanceResourceProtectRequest() (request *SetInstanceResourceProtectRequest) {
	request = &SetInstanceResourceProtectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "SetInstanceResourceProtect")
	return
}

func NewSetInstanceResourceProtectResponse() (response *SetInstanceResourceProtectResponse) {
	response = &SetInstanceResourceProtectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetInstanceResourceProtect(request *SetInstanceResourceProtectRequest) string {
	return c.SetInstanceResourceProtectWithContext(context.Background(), request)
}

func (c *Client) SetInstanceResourceProtectSend(request *SetInstanceResourceProtectRequest) (*SetInstanceResourceProtectResponse, error) {
	statusCode, msg, err := c.SetInstanceResourceProtectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetInstanceResourceProtectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetInstanceResourceProtectWithContext(ctx context.Context, request *SetInstanceResourceProtectRequest) string {
	if request == nil {
		request = NewSetInstanceResourceProtectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetInstanceResourceProtectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetInstanceResourceProtectWithContextV2(ctx context.Context, request *SetInstanceResourceProtectRequest) (int, string, error) {
	if request == nil {
		request = NewSetInstanceResourceProtectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetInstanceResourceProtectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
	request = &DescribeInstanceVncUrlRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kec", APIVersion, "DescribeInstanceVncUrl")
	return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
	response = &DescribeInstanceVncUrlResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) string {
	return c.DescribeInstanceVncUrlWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceVncUrlSend(request *DescribeInstanceVncUrlRequest) (*DescribeInstanceVncUrlResponse, error) {
	statusCode, msg, err := c.DescribeInstanceVncUrlWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceVncUrlResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest) string {
	if request == nil {
		request = NewDescribeInstanceVncUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceVncUrlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceVncUrlWithContextV2(ctx context.Context, request *DescribeInstanceVncUrlRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceVncUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceVncUrlResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
