package v20200702

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-07-02"

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

func NewCreateContainerGroupRequest() (request *CreateContainerGroupRequest) {
	request = &CreateContainerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "CreateContainerGroup")
	return
}

func NewCreateContainerGroupResponse() (response *CreateContainerGroupResponse) {
	response = &CreateContainerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateContainerGroup(request *CreateContainerGroupRequest) string {
	return c.CreateContainerGroupWithContext(context.Background(), request)
}

func (c *Client) CreateContainerGroupSend(request *CreateContainerGroupRequest) (*CreateContainerGroupResponse, error) {
	statusCode, msg, err := c.CreateContainerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateContainerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateContainerGroupWithContext(ctx context.Context, request *CreateContainerGroupRequest) string {
	if request == nil {
		request = NewCreateContainerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateContainerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateContainerGroupWithContextV2(ctx context.Context, request *CreateContainerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateContainerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateContainerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeContainerGroupRequest() (request *DescribeContainerGroupRequest) {
	request = &DescribeContainerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeContainerGroup")
	return
}

func NewDescribeContainerGroupResponse() (response *DescribeContainerGroupResponse) {
	response = &DescribeContainerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeContainerGroup(request *DescribeContainerGroupRequest) string {
	return c.DescribeContainerGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeContainerGroupSend(request *DescribeContainerGroupRequest) (*DescribeContainerGroupResponse, error) {
	statusCode, msg, err := c.DescribeContainerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeContainerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeContainerGroupWithContext(ctx context.Context, request *DescribeContainerGroupRequest) string {
	if request == nil {
		request = NewDescribeContainerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeContainerGroupWithContextV2(ctx context.Context, request *DescribeContainerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeContainerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteContainerGroupRequest() (request *DeleteContainerGroupRequest) {
	request = &DeleteContainerGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DeleteContainerGroup")
	return
}

func NewDeleteContainerGroupResponse() (response *DeleteContainerGroupResponse) {
	response = &DeleteContainerGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) string {
	return c.DeleteContainerGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteContainerGroupSend(request *DeleteContainerGroupRequest) (*DeleteContainerGroupResponse, error) {
	statusCode, msg, err := c.DeleteContainerGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteContainerGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteContainerGroupWithContext(ctx context.Context, request *DeleteContainerGroupRequest) string {
	if request == nil {
		request = NewDeleteContainerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteContainerGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteContainerGroupWithContextV2(ctx context.Context, request *DeleteContainerGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteContainerGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteContainerGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeContainerLogRequest() (request *DescribeContainerLogRequest) {
	request = &DescribeContainerLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeContainerLog")
	return
}

func NewDescribeContainerLogResponse() (response *DescribeContainerLogResponse) {
	response = &DescribeContainerLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeContainerLog(request *DescribeContainerLogRequest) string {
	return c.DescribeContainerLogWithContext(context.Background(), request)
}

func (c *Client) DescribeContainerLogSend(request *DescribeContainerLogRequest) (*DescribeContainerLogResponse, error) {
	statusCode, msg, err := c.DescribeContainerLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeContainerLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeContainerLogWithContext(ctx context.Context, request *DescribeContainerLogRequest) string {
	if request == nil {
		request = NewDescribeContainerLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeContainerLogWithContextV2(ctx context.Context, request *DescribeContainerLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeContainerLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerLogResponse()
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
	request.Init().WithApiInfo("kci", APIVersion, "DescribeRegions")
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
func NewExecContainerCommandRequest() (request *ExecContainerCommandRequest) {
	request = &ExecContainerCommandRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "ExecContainerCommand")
	return
}

func NewExecContainerCommandResponse() (response *ExecContainerCommandResponse) {
	response = &ExecContainerCommandResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ExecContainerCommand(request *ExecContainerCommandRequest) string {
	return c.ExecContainerCommandWithContext(context.Background(), request)
}

func (c *Client) ExecContainerCommandSend(request *ExecContainerCommandRequest) (*ExecContainerCommandResponse, error) {
	statusCode, msg, err := c.ExecContainerCommandWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ExecContainerCommandResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ExecContainerCommandWithContext(ctx context.Context, request *ExecContainerCommandRequest) string {
	if request == nil {
		request = NewExecContainerCommandRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewExecContainerCommandResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ExecContainerCommandWithContextV2(ctx context.Context, request *ExecContainerCommandRequest) (int, string, error) {
	if request == nil {
		request = NewExecContainerCommandRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewExecContainerCommandResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeContainerGroupCountRequest() (request *DescribeContainerGroupCountRequest) {
	request = &DescribeContainerGroupCountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeContainerGroupCount")
	return
}

func NewDescribeContainerGroupCountResponse() (response *DescribeContainerGroupCountResponse) {
	response = &DescribeContainerGroupCountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeContainerGroupCount(request *DescribeContainerGroupCountRequest) string {
	return c.DescribeContainerGroupCountWithContext(context.Background(), request)
}

func (c *Client) DescribeContainerGroupCountSend(request *DescribeContainerGroupCountRequest) (*DescribeContainerGroupCountResponse, error) {
	statusCode, msg, err := c.DescribeContainerGroupCountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeContainerGroupCountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeContainerGroupCountWithContext(ctx context.Context, request *DescribeContainerGroupCountRequest) string {
	if request == nil {
		request = NewDescribeContainerGroupCountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerGroupCountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeContainerGroupCountWithContextV2(ctx context.Context, request *DescribeContainerGroupCountRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeContainerGroupCountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerGroupCountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeContainerGroupEventsRequest() (request *DescribeContainerGroupEventsRequest) {
	request = &DescribeContainerGroupEventsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeContainerGroupEvents")
	return
}

func NewDescribeContainerGroupEventsResponse() (response *DescribeContainerGroupEventsResponse) {
	response = &DescribeContainerGroupEventsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeContainerGroupEvents(request *DescribeContainerGroupEventsRequest) string {
	return c.DescribeContainerGroupEventsWithContext(context.Background(), request)
}

func (c *Client) DescribeContainerGroupEventsSend(request *DescribeContainerGroupEventsRequest) (*DescribeContainerGroupEventsResponse, error) {
	statusCode, msg, err := c.DescribeContainerGroupEventsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeContainerGroupEventsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeContainerGroupEventsWithContext(ctx context.Context, request *DescribeContainerGroupEventsRequest) string {
	if request == nil {
		request = NewDescribeContainerGroupEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerGroupEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeContainerGroupEventsWithContextV2(ctx context.Context, request *DescribeContainerGroupEventsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeContainerGroupEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeContainerGroupEventsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKciPackagesRequest() (request *DescribeKciPackagesRequest) {
	request = &DescribeKciPackagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeKciPackages")
	return
}

func NewDescribeKciPackagesResponse() (response *DescribeKciPackagesResponse) {
	response = &DescribeKciPackagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKciPackages(request *DescribeKciPackagesRequest) string {
	return c.DescribeKciPackagesWithContext(context.Background(), request)
}

func (c *Client) DescribeKciPackagesSend(request *DescribeKciPackagesRequest) (*DescribeKciPackagesResponse, error) {
	statusCode, msg, err := c.DescribeKciPackagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeKciPackagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKciPackagesWithContext(ctx context.Context, request *DescribeKciPackagesRequest) string {
	if request == nil {
		request = NewDescribeKciPackagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeKciPackagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKciPackagesWithContextV2(ctx context.Context, request *DescribeKciPackagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKciPackagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeKciPackagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateImageCacheRequest() (request *CreateImageCacheRequest) {
	request = &CreateImageCacheRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "CreateImageCache")
	return
}

func NewCreateImageCacheResponse() (response *CreateImageCacheResponse) {
	response = &CreateImageCacheResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImageCache(request *CreateImageCacheRequest) string {
	return c.CreateImageCacheWithContext(context.Background(), request)
}

func (c *Client) CreateImageCacheSend(request *CreateImageCacheRequest) (*CreateImageCacheResponse, error) {
	statusCode, msg, err := c.CreateImageCacheWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateImageCacheResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateImageCacheWithContext(ctx context.Context, request *CreateImageCacheRequest) string {
	if request == nil {
		request = NewCreateImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateImageCacheResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateImageCacheWithContextV2(ctx context.Context, request *CreateImageCacheRequest) (int, string, error) {
	if request == nil {
		request = NewCreateImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateImageCacheResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteImageCacheRequest() (request *DeleteImageCacheRequest) {
	request = &DeleteImageCacheRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DeleteImageCache")
	return
}

func NewDeleteImageCacheResponse() (response *DeleteImageCacheResponse) {
	response = &DeleteImageCacheResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteImageCache(request *DeleteImageCacheRequest) string {
	return c.DeleteImageCacheWithContext(context.Background(), request)
}

func (c *Client) DeleteImageCacheSend(request *DeleteImageCacheRequest) (*DeleteImageCacheResponse, error) {
	statusCode, msg, err := c.DeleteImageCacheWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteImageCacheResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteImageCacheWithContext(ctx context.Context, request *DeleteImageCacheRequest) string {
	if request == nil {
		request = NewDeleteImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageCacheResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteImageCacheWithContextV2(ctx context.Context, request *DeleteImageCacheRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageCacheResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImageCacheRequest() (request *DescribeImageCacheRequest) {
	request = &DescribeImageCacheRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeImageCache")
	return
}

func NewDescribeImageCacheResponse() (response *DescribeImageCacheResponse) {
	response = &DescribeImageCacheResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImageCache(request *DescribeImageCacheRequest) string {
	return c.DescribeImageCacheWithContext(context.Background(), request)
}

func (c *Client) DescribeImageCacheSend(request *DescribeImageCacheRequest) (*DescribeImageCacheResponse, error) {
	statusCode, msg, err := c.DescribeImageCacheWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImageCacheResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImageCacheWithContext(ctx context.Context, request *DescribeImageCacheRequest) string {
	if request == nil {
		request = NewDescribeImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageCacheResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImageCacheWithContextV2(ctx context.Context, request *DescribeImageCacheRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageCacheResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewMatchImageCacheRequest() (request *MatchImageCacheRequest) {
	request = &MatchImageCacheRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "MatchImageCache")
	return
}

func NewMatchImageCacheResponse() (response *MatchImageCacheResponse) {
	response = &MatchImageCacheResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) MatchImageCache(request *MatchImageCacheRequest) string {
	return c.MatchImageCacheWithContext(context.Background(), request)
}

func (c *Client) MatchImageCacheSend(request *MatchImageCacheRequest) (*MatchImageCacheResponse, error) {
	statusCode, msg, err := c.MatchImageCacheWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct MatchImageCacheResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) MatchImageCacheWithContext(ctx context.Context, request *MatchImageCacheRequest) string {
	if request == nil {
		request = NewMatchImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewMatchImageCacheResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) MatchImageCacheWithContextV2(ctx context.Context, request *MatchImageCacheRequest) (int, string, error) {
	if request == nil {
		request = NewMatchImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewMatchImageCacheResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImageCacheEventRequest() (request *DescribeImageCacheEventRequest) {
	request = &DescribeImageCacheEventRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "DescribeImageCacheEvent")
	return
}

func NewDescribeImageCacheEventResponse() (response *DescribeImageCacheEventResponse) {
	response = &DescribeImageCacheEventResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImageCacheEvent(request *DescribeImageCacheEventRequest) string {
	return c.DescribeImageCacheEventWithContext(context.Background(), request)
}

func (c *Client) DescribeImageCacheEventSend(request *DescribeImageCacheEventRequest) (*DescribeImageCacheEventResponse, error) {
	statusCode, msg, err := c.DescribeImageCacheEventWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImageCacheEventResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImageCacheEventWithContext(ctx context.Context, request *DescribeImageCacheEventRequest) string {
	if request == nil {
		request = NewDescribeImageCacheEventRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageCacheEventResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImageCacheEventWithContextV2(ctx context.Context, request *DescribeImageCacheEventRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImageCacheEventRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageCacheEventResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateImageCacheRequest() (request *UpdateImageCacheRequest) {
	request = &UpdateImageCacheRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kci", APIVersion, "UpdateImageCache")
	return
}

func NewUpdateImageCacheResponse() (response *UpdateImageCacheResponse) {
	response = &UpdateImageCacheResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateImageCache(request *UpdateImageCacheRequest) string {
	return c.UpdateImageCacheWithContext(context.Background(), request)
}

func (c *Client) UpdateImageCacheSend(request *UpdateImageCacheRequest) (*UpdateImageCacheResponse, error) {
	statusCode, msg, err := c.UpdateImageCacheWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateImageCacheResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateImageCacheWithContext(ctx context.Context, request *UpdateImageCacheRequest) string {
	if request == nil {
		request = NewUpdateImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateImageCacheResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateImageCacheWithContextV2(ctx context.Context, request *UpdateImageCacheRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateImageCacheRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateImageCacheResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
