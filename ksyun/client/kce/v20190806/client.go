package v20190806
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2019-08-06"

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

func NewDescribeClusterInstanceRequest() (request *DescribeClusterInstanceRequest) {
	request = &DescribeClusterInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeClusterInstance")
	return
}

func NewDescribeClusterInstanceResponse() (response *DescribeClusterInstanceResponse) {
	response = &DescribeClusterInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusterInstance(request *DescribeClusterInstanceRequest) string {
	return c.DescribeClusterInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterInstanceSend(request *DescribeClusterInstanceRequest) (*DescribeClusterInstanceResponse, error) {
	statusCode, msg, err := c.DescribeClusterInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClusterInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClusterInstanceWithContext(ctx context.Context, request *DescribeClusterInstanceRequest) string {
	if request == nil {
		request = NewDescribeClusterInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClusterInstanceWithContextV2(ctx context.Context, request *DescribeClusterInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteCluster")
	return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCluster(request *DeleteClusterRequest) string {
	return c.DeleteClusterWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterSend(request *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	statusCode, msg, err := c.DeleteClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) string {
	if request == nil {
		request = NewDeleteClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteClusterWithContextV2(ctx context.Context, request *DeleteClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDownloadClusterConfigRequest() (request *DownloadClusterConfigRequest) {
	request = &DownloadClusterConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DownloadClusterConfig")
	return
}

func NewDownloadClusterConfigResponse() (response *DownloadClusterConfigResponse) {
	response = &DownloadClusterConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DownloadClusterConfig(request *DownloadClusterConfigRequest) string {
	return c.DownloadClusterConfigWithContext(context.Background(), request)
}

func (c *Client) DownloadClusterConfigSend(request *DownloadClusterConfigRequest) (*DownloadClusterConfigResponse, error) {
	statusCode, msg, err := c.DownloadClusterConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DownloadClusterConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DownloadClusterConfigWithContext(ctx context.Context, request *DownloadClusterConfigRequest) string {
	if request == nil {
		request = NewDownloadClusterConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDownloadClusterConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DownloadClusterConfigWithContextV2(ctx context.Context, request *DownloadClusterConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDownloadClusterConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDownloadClusterConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyClusterInfoRequest() (request *ModifyClusterInfoRequest) {
	request = &ModifyClusterInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyClusterInfo")
	return
}

func NewModifyClusterInfoResponse() (response *ModifyClusterInfoResponse) {
	response = &ModifyClusterInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyClusterInfo(request *ModifyClusterInfoRequest) string {
	return c.ModifyClusterInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyClusterInfoSend(request *ModifyClusterInfoRequest) (*ModifyClusterInfoResponse, error) {
	statusCode, msg, err := c.ModifyClusterInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyClusterInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyClusterInfoWithContext(ctx context.Context, request *ModifyClusterInfoRequest) string {
	if request == nil {
		request = NewModifyClusterInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyClusterInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyClusterInfoWithContextV2(ctx context.Context, request *ModifyClusterInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyClusterInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyClusterInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceImageRequest() (request *DescribeInstanceImageRequest) {
	request = &DescribeInstanceImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeInstanceImage")
	return
}

func NewDescribeInstanceImageResponse() (response *DescribeInstanceImageResponse) {
	response = &DescribeInstanceImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceImage(request *DescribeInstanceImageRequest) string {
	return c.DescribeInstanceImageWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceImageSend(request *DescribeInstanceImageRequest) (*DescribeInstanceImageResponse, error) {
	statusCode, msg, err := c.DescribeInstanceImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceImageWithContext(ctx context.Context, request *DescribeInstanceImageRequest) string {
	if request == nil {
		request = NewDescribeInstanceImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceImageWithContextV2(ctx context.Context, request *DescribeInstanceImageRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
	request = &AddClusterInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "AddClusterInstances")
	return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
	response = &AddClusterInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) string {
	return c.AddClusterInstancesWithContext(context.Background(), request)
}

func (c *Client) AddClusterInstancesSend(request *AddClusterInstancesRequest) (*AddClusterInstancesResponse, error) {
	statusCode, msg, err := c.AddClusterInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddClusterInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddClusterInstancesWithContext(ctx context.Context, request *AddClusterInstancesRequest) string {
	if request == nil {
		request = NewAddClusterInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddClusterInstancesWithContextV2(ctx context.Context, request *AddClusterInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewAddClusterInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
	request = &DeleteClusterInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteClusterInstances")
	return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
	response = &DeleteClusterInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) string {
	return c.DeleteClusterInstancesWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterInstancesSend(request *DeleteClusterInstancesRequest) (*DeleteClusterInstancesResponse, error) {
	statusCode, msg, err := c.DeleteClusterInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteClusterInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteClusterInstancesWithContext(ctx context.Context, request *DeleteClusterInstancesRequest) string {
	if request == nil {
		request = NewDeleteClusterInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteClusterInstancesWithContextV2(ctx context.Context, request *DeleteClusterInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteClusterInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcForClusterRequest() (request *DescribeEpcForClusterRequest) {
	request = &DescribeEpcForClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeEpcForCluster")
	return
}

func NewDescribeEpcForClusterResponse() (response *DescribeEpcForClusterResponse) {
	response = &DescribeEpcForClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcForCluster(request *DescribeEpcForClusterRequest) string {
	return c.DescribeEpcForClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcForClusterSend(request *DescribeEpcForClusterRequest) (*DescribeEpcForClusterResponse, error) {
	statusCode, msg, err := c.DescribeEpcForClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcForClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcForClusterWithContext(ctx context.Context, request *DescribeEpcForClusterRequest) string {
	if request == nil {
		request = NewDescribeEpcForClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcForClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcForClusterWithContextV2(ctx context.Context, request *DescribeEpcForClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcForClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcForClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddClusterEpcInstancesRequest() (request *AddClusterEpcInstancesRequest) {
	request = &AddClusterEpcInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "AddClusterEpcInstances")
	return
}

func NewAddClusterEpcInstancesResponse() (response *AddClusterEpcInstancesResponse) {
	response = &AddClusterEpcInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddClusterEpcInstances(request *AddClusterEpcInstancesRequest) string {
	return c.AddClusterEpcInstancesWithContext(context.Background(), request)
}

func (c *Client) AddClusterEpcInstancesSend(request *AddClusterEpcInstancesRequest) (*AddClusterEpcInstancesResponse, error) {
	statusCode, msg, err := c.AddClusterEpcInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddClusterEpcInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddClusterEpcInstancesWithContext(ctx context.Context, request *AddClusterEpcInstancesRequest) string {
	if request == nil {
		request = NewAddClusterEpcInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterEpcInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddClusterEpcInstancesWithContextV2(ctx context.Context, request *AddClusterEpcInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewAddClusterEpcInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterEpcInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeExistedInstancesRequest() (request *DescribeExistedInstancesRequest) {
	request = &DescribeExistedInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeExistedInstances")
	return
}

func NewDescribeExistedInstancesResponse() (response *DescribeExistedInstancesResponse) {
	response = &DescribeExistedInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeExistedInstances(request *DescribeExistedInstancesRequest) string {
	return c.DescribeExistedInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeExistedInstancesSend(request *DescribeExistedInstancesRequest) (*DescribeExistedInstancesResponse, error) {
	statusCode, msg, err := c.DescribeExistedInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeExistedInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeExistedInstancesWithContext(ctx context.Context, request *DescribeExistedInstancesRequest) string {
	if request == nil {
		request = NewDescribeExistedInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeExistedInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeExistedInstancesWithContextV2(ctx context.Context, request *DescribeExistedInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeExistedInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeExistedInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddExistedInstancesRequest() (request *AddExistedInstancesRequest) {
	request = &AddExistedInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "AddExistedInstances")
	return
}

func NewAddExistedInstancesResponse() (response *AddExistedInstancesResponse) {
	response = &AddExistedInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddExistedInstances(request *AddExistedInstancesRequest) string {
	return c.AddExistedInstancesWithContext(context.Background(), request)
}

func (c *Client) AddExistedInstancesSend(request *AddExistedInstancesRequest) (*AddExistedInstancesResponse, error) {
	statusCode, msg, err := c.AddExistedInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddExistedInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddExistedInstancesWithContext(ctx context.Context, request *AddExistedInstancesRequest) string {
	if request == nil {
		request = NewAddExistedInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddExistedInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddExistedInstancesWithContextV2(ctx context.Context, request *AddExistedInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewAddExistedInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddExistedInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateNodePoolRequest() (request *CreateNodePoolRequest) {
	request = &CreateNodePoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateNodePool")
	return
}

func NewCreateNodePoolResponse() (response *CreateNodePoolResponse) {
	response = &CreateNodePoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNodePool(request *CreateNodePoolRequest) string {
	return c.CreateNodePoolWithContext(context.Background(), request)
}

func (c *Client) CreateNodePoolSend(request *CreateNodePoolRequest) (*CreateNodePoolResponse, error) {
	statusCode, msg, err := c.CreateNodePoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateNodePoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateNodePoolWithContext(ctx context.Context, request *CreateNodePoolRequest) string {
	if request == nil {
		request = NewCreateNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNodePoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateNodePoolWithContextV2(ctx context.Context, request *CreateNodePoolRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNodePoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNodePoolRequest() (request *DescribeNodePoolRequest) {
	request = &DescribeNodePoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeNodePool")
	return
}

func NewDescribeNodePoolResponse() (response *DescribeNodePoolResponse) {
	response = &DescribeNodePoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNodePool(request *DescribeNodePoolRequest) string {
	return c.DescribeNodePoolWithContext(context.Background(), request)
}

func (c *Client) DescribeNodePoolSend(request *DescribeNodePoolRequest) (*DescribeNodePoolResponse, error) {
	statusCode, msg, err := c.DescribeNodePoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNodePoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNodePoolWithContext(ctx context.Context, request *DescribeNodePoolRequest) string {
	if request == nil {
		request = NewDescribeNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodePoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNodePoolWithContextV2(ctx context.Context, request *DescribeNodePoolRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodePoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteNodePoolRequest() (request *DeleteNodePoolRequest) {
	request = &DeleteNodePoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteNodePool")
	return
}

func NewDeleteNodePoolResponse() (response *DeleteNodePoolResponse) {
	response = &DeleteNodePoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNodePool(request *DeleteNodePoolRequest) string {
	return c.DeleteNodePoolWithContext(context.Background(), request)
}

func (c *Client) DeleteNodePoolSend(request *DeleteNodePoolRequest) (*DeleteNodePoolResponse, error) {
	statusCode, msg, err := c.DeleteNodePoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteNodePoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteNodePoolWithContext(ctx context.Context, request *DeleteNodePoolRequest) string {
	if request == nil {
		request = NewDeleteNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNodePoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteNodePoolWithContextV2(ctx context.Context, request *DeleteNodePoolRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNodePoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNodePoolRequest() (request *ModifyNodePoolRequest) {
	request = &ModifyNodePoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyNodePool")
	return
}

func NewModifyNodePoolResponse() (response *ModifyNodePoolResponse) {
	response = &ModifyNodePoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNodePool(request *ModifyNodePoolRequest) string {
	return c.ModifyNodePoolWithContext(context.Background(), request)
}

func (c *Client) ModifyNodePoolSend(request *ModifyNodePoolRequest) (*ModifyNodePoolResponse, error) {
	statusCode, msg, err := c.ModifyNodePoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNodePoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNodePoolWithContext(ctx context.Context, request *ModifyNodePoolRequest) string {
	if request == nil {
		request = NewModifyNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodePoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNodePoolWithContextV2(ctx context.Context, request *ModifyNodePoolRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodePoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNodeTemplateRequest() (request *ModifyNodeTemplateRequest) {
	request = &ModifyNodeTemplateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyNodeTemplate")
	return
}

func NewModifyNodeTemplateResponse() (response *ModifyNodeTemplateResponse) {
	response = &ModifyNodeTemplateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNodeTemplate(request *ModifyNodeTemplateRequest) string {
	return c.ModifyNodeTemplateWithContext(context.Background(), request)
}

func (c *Client) ModifyNodeTemplateSend(request *ModifyNodeTemplateRequest) (*ModifyNodeTemplateResponse, error) {
	statusCode, msg, err := c.ModifyNodeTemplateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNodeTemplateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNodeTemplateWithContext(ctx context.Context, request *ModifyNodeTemplateRequest) string {
	if request == nil {
		request = NewModifyNodeTemplateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodeTemplateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNodeTemplateWithContextV2(ctx context.Context, request *ModifyNodeTemplateRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNodeTemplateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodeTemplateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNodePoolScaleUpPolicyRequest() (request *ModifyNodePoolScaleUpPolicyRequest) {
	request = &ModifyNodePoolScaleUpPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyNodePoolScaleUpPolicy")
	return
}

func NewModifyNodePoolScaleUpPolicyResponse() (response *ModifyNodePoolScaleUpPolicyResponse) {
	response = &ModifyNodePoolScaleUpPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNodePoolScaleUpPolicy(request *ModifyNodePoolScaleUpPolicyRequest) string {
	return c.ModifyNodePoolScaleUpPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyNodePoolScaleUpPolicySend(request *ModifyNodePoolScaleUpPolicyRequest) (*ModifyNodePoolScaleUpPolicyResponse, error) {
	statusCode, msg, err := c.ModifyNodePoolScaleUpPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNodePoolScaleUpPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNodePoolScaleUpPolicyWithContext(ctx context.Context, request *ModifyNodePoolScaleUpPolicyRequest) string {
	if request == nil {
		request = NewModifyNodePoolScaleUpPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodePoolScaleUpPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNodePoolScaleUpPolicyWithContextV2(ctx context.Context, request *ModifyNodePoolScaleUpPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNodePoolScaleUpPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodePoolScaleUpPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNodePoolScaleDownPolicyRequest() (request *ModifyNodePoolScaleDownPolicyRequest) {
	request = &ModifyNodePoolScaleDownPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyNodePoolScaleDownPolicy")
	return
}

func NewModifyNodePoolScaleDownPolicyResponse() (response *ModifyNodePoolScaleDownPolicyResponse) {
	response = &ModifyNodePoolScaleDownPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNodePoolScaleDownPolicy(request *ModifyNodePoolScaleDownPolicyRequest) string {
	return c.ModifyNodePoolScaleDownPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyNodePoolScaleDownPolicySend(request *ModifyNodePoolScaleDownPolicyRequest) (*ModifyNodePoolScaleDownPolicyResponse, error) {
	statusCode, msg, err := c.ModifyNodePoolScaleDownPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNodePoolScaleDownPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNodePoolScaleDownPolicyWithContext(ctx context.Context, request *ModifyNodePoolScaleDownPolicyRequest) string {
	if request == nil {
		request = NewModifyNodePoolScaleDownPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodePoolScaleDownPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNodePoolScaleDownPolicyWithContextV2(ctx context.Context, request *ModifyNodePoolScaleDownPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNodePoolScaleDownPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodePoolScaleDownPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddClusterInstanceToNodePoolRequest() (request *AddClusterInstanceToNodePoolRequest) {
	request = &AddClusterInstanceToNodePoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "AddClusterInstanceToNodePool")
	return
}

func NewAddClusterInstanceToNodePoolResponse() (response *AddClusterInstanceToNodePoolResponse) {
	response = &AddClusterInstanceToNodePoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddClusterInstanceToNodePool(request *AddClusterInstanceToNodePoolRequest) string {
	return c.AddClusterInstanceToNodePoolWithContext(context.Background(), request)
}

func (c *Client) AddClusterInstanceToNodePoolSend(request *AddClusterInstanceToNodePoolRequest) (*AddClusterInstanceToNodePoolResponse, error) {
	statusCode, msg, err := c.AddClusterInstanceToNodePoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddClusterInstanceToNodePoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddClusterInstanceToNodePoolWithContext(ctx context.Context, request *AddClusterInstanceToNodePoolRequest) string {
	if request == nil {
		request = NewAddClusterInstanceToNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterInstanceToNodePoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddClusterInstanceToNodePoolWithContextV2(ctx context.Context, request *AddClusterInstanceToNodePoolRequest) (int, string, error) {
	if request == nil {
		request = NewAddClusterInstanceToNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddClusterInstanceToNodePoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewProtectedFromScaleDownRequest() (request *ProtectedFromScaleDownRequest) {
	request = &ProtectedFromScaleDownRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ProtectedFromScaleDown")
	return
}

func NewProtectedFromScaleDownResponse() (response *ProtectedFromScaleDownResponse) {
	response = &ProtectedFromScaleDownResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ProtectedFromScaleDown(request *ProtectedFromScaleDownRequest) string {
	return c.ProtectedFromScaleDownWithContext(context.Background(), request)
}

func (c *Client) ProtectedFromScaleDownSend(request *ProtectedFromScaleDownRequest) (*ProtectedFromScaleDownResponse, error) {
	statusCode, msg, err := c.ProtectedFromScaleDownWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ProtectedFromScaleDownResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ProtectedFromScaleDownWithContext(ctx context.Context, request *ProtectedFromScaleDownRequest) string {
	if request == nil {
		request = NewProtectedFromScaleDownRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewProtectedFromScaleDownResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ProtectedFromScaleDownWithContextV2(ctx context.Context, request *ProtectedFromScaleDownRequest) (int, string, error) {
	if request == nil {
		request = NewProtectedFromScaleDownRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewProtectedFromScaleDownResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteClusterInstancesFromNodePoolRequest() (request *DeleteClusterInstancesFromNodePoolRequest) {
	request = &DeleteClusterInstancesFromNodePoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteClusterInstancesFromNodePool")
	return
}

func NewDeleteClusterInstancesFromNodePoolResponse() (response *DeleteClusterInstancesFromNodePoolResponse) {
	response = &DeleteClusterInstancesFromNodePoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteClusterInstancesFromNodePool(request *DeleteClusterInstancesFromNodePoolRequest) string {
	return c.DeleteClusterInstancesFromNodePoolWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterInstancesFromNodePoolSend(request *DeleteClusterInstancesFromNodePoolRequest) (*DeleteClusterInstancesFromNodePoolResponse, error) {
	statusCode, msg, err := c.DeleteClusterInstancesFromNodePoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteClusterInstancesFromNodePoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteClusterInstancesFromNodePoolWithContext(ctx context.Context, request *DeleteClusterInstancesFromNodePoolRequest) string {
	if request == nil {
		request = NewDeleteClusterInstancesFromNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterInstancesFromNodePoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteClusterInstancesFromNodePoolWithContextV2(ctx context.Context, request *DeleteClusterInstancesFromNodePoolRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteClusterInstancesFromNodePoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteClusterInstancesFromNodePoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcImageRequest() (request *DescribeEpcImageRequest) {
	request = &DescribeEpcImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeEpcImage")
	return
}

func NewDescribeEpcImageResponse() (response *DescribeEpcImageResponse) {
	response = &DescribeEpcImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcImage(request *DescribeEpcImageRequest) string {
	return c.DescribeEpcImageWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcImageSend(request *DescribeEpcImageRequest) (*DescribeEpcImageResponse, error) {
	statusCode, msg, err := c.DescribeEpcImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcImageWithContext(ctx context.Context, request *DescribeEpcImageRequest) string {
	if request == nil {
		request = NewDescribeEpcImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeEpcImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcImageWithContextV2(ctx context.Context, request *DescribeEpcImageRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeEpcImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEditEventCollectingRequest() (request *EditEventCollectingRequest) {
	request = &EditEventCollectingRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "EditEventCollecting")
	return
}

func NewEditEventCollectingResponse() (response *EditEventCollectingResponse) {
	response = &EditEventCollectingResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EditEventCollecting(request *EditEventCollectingRequest) string {
	return c.EditEventCollectingWithContext(context.Background(), request)
}

func (c *Client) EditEventCollectingSend(request *EditEventCollectingRequest) (*EditEventCollectingResponse, error) {
	statusCode, msg, err := c.EditEventCollectingWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct EditEventCollectingResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EditEventCollectingWithContext(ctx context.Context, request *EditEventCollectingRequest) string {
	if request == nil {
		request = NewEditEventCollectingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEditEventCollectingResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EditEventCollectingWithContextV2(ctx context.Context, request *EditEventCollectingRequest) (int, string, error) {
	if request == nil {
		request = NewEditEventCollectingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEditEventCollectingResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNodePoolSummaryRequest() (request *DescribeNodePoolSummaryRequest) {
	request = &DescribeNodePoolSummaryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeNodePoolSummary")
	return
}

func NewDescribeNodePoolSummaryResponse() (response *DescribeNodePoolSummaryResponse) {
	response = &DescribeNodePoolSummaryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNodePoolSummary(request *DescribeNodePoolSummaryRequest) string {
	return c.DescribeNodePoolSummaryWithContext(context.Background(), request)
}

func (c *Client) DescribeNodePoolSummarySend(request *DescribeNodePoolSummaryRequest) (*DescribeNodePoolSummaryResponse, error) {
	statusCode, msg, err := c.DescribeNodePoolSummaryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNodePoolSummaryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNodePoolSummaryWithContext(ctx context.Context, request *DescribeNodePoolSummaryRequest) string {
	if request == nil {
		request = NewDescribeNodePoolSummaryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodePoolSummaryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNodePoolSummaryWithContextV2(ctx context.Context, request *DescribeNodePoolSummaryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNodePoolSummaryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodePoolSummaryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateLogRuleRequest() (request *CreateLogRuleRequest) {
	request = &CreateLogRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateLogRule")
	return
}

func NewCreateLogRuleResponse() (response *CreateLogRuleResponse) {
	response = &CreateLogRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLogRule(request *CreateLogRuleRequest) string {
	return c.CreateLogRuleWithContext(context.Background(), request)
}

func (c *Client) CreateLogRuleSend(request *CreateLogRuleRequest) (*CreateLogRuleResponse, error) {
	statusCode, msg, err := c.CreateLogRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateLogRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateLogRuleWithContext(ctx context.Context, request *CreateLogRuleRequest) string {
	if request == nil {
		request = NewCreateLogRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLogRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateLogRuleWithContextV2(ctx context.Context, request *CreateLogRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLogRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateLogRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeClusterSummaryRequest() (request *DescribeClusterSummaryRequest) {
	request = &DescribeClusterSummaryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeClusterSummary")
	return
}

func NewDescribeClusterSummaryResponse() (response *DescribeClusterSummaryResponse) {
	response = &DescribeClusterSummaryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusterSummary(request *DescribeClusterSummaryRequest) string {
	return c.DescribeClusterSummaryWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterSummarySend(request *DescribeClusterSummaryRequest) (*DescribeClusterSummaryResponse, error) {
	statusCode, msg, err := c.DescribeClusterSummaryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClusterSummaryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClusterSummaryWithContext(ctx context.Context, request *DescribeClusterSummaryRequest) string {
	if request == nil {
		request = NewDescribeClusterSummaryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterSummaryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClusterSummaryWithContextV2(ctx context.Context, request *DescribeClusterSummaryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterSummaryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterSummaryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateNodePoolDelProtectionRequest() (request *UpdateNodePoolDelProtectionRequest) {
	request = &UpdateNodePoolDelProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdateNodePoolDelProtection")
	return
}

func NewUpdateNodePoolDelProtectionResponse() (response *UpdateNodePoolDelProtectionResponse) {
	response = &UpdateNodePoolDelProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateNodePoolDelProtection(request *UpdateNodePoolDelProtectionRequest) string {
	return c.UpdateNodePoolDelProtectionWithContext(context.Background(), request)
}

func (c *Client) UpdateNodePoolDelProtectionSend(request *UpdateNodePoolDelProtectionRequest) (*UpdateNodePoolDelProtectionResponse, error) {
	statusCode, msg, err := c.UpdateNodePoolDelProtectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateNodePoolDelProtectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateNodePoolDelProtectionWithContext(ctx context.Context, request *UpdateNodePoolDelProtectionRequest) string {
	if request == nil {
		request = NewUpdateNodePoolDelProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateNodePoolDelProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateNodePoolDelProtectionWithContextV2(ctx context.Context, request *UpdateNodePoolDelProtectionRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateNodePoolDelProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateNodePoolDelProtectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeReleaseRequest() (request *DescribeReleaseRequest) {
	request = &DescribeReleaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeRelease")
	return
}

func NewDescribeReleaseResponse() (response *DescribeReleaseResponse) {
	response = &DescribeReleaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRelease(request *DescribeReleaseRequest) string {
	return c.DescribeReleaseWithContext(context.Background(), request)
}

func (c *Client) DescribeReleaseSend(request *DescribeReleaseRequest) (*DescribeReleaseResponse, error) {
	statusCode, msg, err := c.DescribeReleaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeReleaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeReleaseWithContext(ctx context.Context, request *DescribeReleaseRequest) string {
	if request == nil {
		request = NewDescribeReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeReleaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeReleaseWithContextV2(ctx context.Context, request *DescribeReleaseRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeReleaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeReleaseHistoryRequest() (request *DescribeReleaseHistoryRequest) {
	request = &DescribeReleaseHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeReleaseHistory")
	return
}

func NewDescribeReleaseHistoryResponse() (response *DescribeReleaseHistoryResponse) {
	response = &DescribeReleaseHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeReleaseHistory(request *DescribeReleaseHistoryRequest) string {
	return c.DescribeReleaseHistoryWithContext(context.Background(), request)
}

func (c *Client) DescribeReleaseHistorySend(request *DescribeReleaseHistoryRequest) (*DescribeReleaseHistoryResponse, error) {
	statusCode, msg, err := c.DescribeReleaseHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeReleaseHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeReleaseHistoryWithContext(ctx context.Context, request *DescribeReleaseHistoryRequest) string {
	if request == nil {
		request = NewDescribeReleaseHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeReleaseHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeReleaseHistoryWithContextV2(ctx context.Context, request *DescribeReleaseHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeReleaseHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeReleaseHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeReleaseDetailRequest() (request *DescribeReleaseDetailRequest) {
	request = &DescribeReleaseDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeReleaseDetail")
	return
}

func NewDescribeReleaseDetailResponse() (response *DescribeReleaseDetailResponse) {
	response = &DescribeReleaseDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeReleaseDetail(request *DescribeReleaseDetailRequest) string {
	return c.DescribeReleaseDetailWithContext(context.Background(), request)
}

func (c *Client) DescribeReleaseDetailSend(request *DescribeReleaseDetailRequest) (*DescribeReleaseDetailResponse, error) {
	statusCode, msg, err := c.DescribeReleaseDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeReleaseDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeReleaseDetailWithContext(ctx context.Context, request *DescribeReleaseDetailRequest) string {
	if request == nil {
		request = NewDescribeReleaseDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeReleaseDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeReleaseDetailWithContextV2(ctx context.Context, request *DescribeReleaseDetailRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeReleaseDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeReleaseDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteReleaseRequest() (request *DeleteReleaseRequest) {
	request = &DeleteReleaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteRelease")
	return
}

func NewDeleteReleaseResponse() (response *DeleteReleaseResponse) {
	response = &DeleteReleaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRelease(request *DeleteReleaseRequest) string {
	return c.DeleteReleaseWithContext(context.Background(), request)
}

func (c *Client) DeleteReleaseSend(request *DeleteReleaseRequest) (*DeleteReleaseResponse, error) {
	statusCode, msg, err := c.DeleteReleaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteReleaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteReleaseWithContext(ctx context.Context, request *DeleteReleaseRequest) string {
	if request == nil {
		request = NewDeleteReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteReleaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteReleaseWithContextV2(ctx context.Context, request *DeleteReleaseRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteReleaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRollbackReleaseRequest() (request *RollbackReleaseRequest) {
	request = &RollbackReleaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "RollbackRelease")
	return
}

func NewRollbackReleaseResponse() (response *RollbackReleaseResponse) {
	response = &RollbackReleaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RollbackRelease(request *RollbackReleaseRequest) string {
	return c.RollbackReleaseWithContext(context.Background(), request)
}

func (c *Client) RollbackReleaseSend(request *RollbackReleaseRequest) (*RollbackReleaseResponse, error) {
	statusCode, msg, err := c.RollbackReleaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RollbackReleaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RollbackReleaseWithContext(ctx context.Context, request *RollbackReleaseRequest) string {
	if request == nil {
		request = NewRollbackReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackReleaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RollbackReleaseWithContextV2(ctx context.Context, request *RollbackReleaseRequest) (int, string, error) {
	if request == nil {
		request = NewRollbackReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackReleaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewInstallReleaseRequest() (request *InstallReleaseRequest) {
	request = &InstallReleaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "InstallRelease")
	return
}

func NewInstallReleaseResponse() (response *InstallReleaseResponse) {
	response = &InstallReleaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) InstallRelease(request *InstallReleaseRequest) string {
	return c.InstallReleaseWithContext(context.Background(), request)
}

func (c *Client) InstallReleaseSend(request *InstallReleaseRequest) (*InstallReleaseResponse, error) {
	statusCode, msg, err := c.InstallReleaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct InstallReleaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) InstallReleaseWithContext(ctx context.Context, request *InstallReleaseRequest) string {
	if request == nil {
		request = NewInstallReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstallReleaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) InstallReleaseWithContextV2(ctx context.Context, request *InstallReleaseRequest) (int, string, error) {
	if request == nil {
		request = NewInstallReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstallReleaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpgradeReleaseRequest() (request *UpgradeReleaseRequest) {
	request = &UpgradeReleaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpgradeRelease")
	return
}

func NewUpgradeReleaseResponse() (response *UpgradeReleaseResponse) {
	response = &UpgradeReleaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpgradeRelease(request *UpgradeReleaseRequest) string {
	return c.UpgradeReleaseWithContext(context.Background(), request)
}

func (c *Client) UpgradeReleaseSend(request *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error) {
	statusCode, msg, err := c.UpgradeReleaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpgradeReleaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpgradeReleaseWithContext(ctx context.Context, request *UpgradeReleaseRequest) string {
	if request == nil {
		request = NewUpgradeReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpgradeReleaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpgradeReleaseWithContextV2(ctx context.Context, request *UpgradeReleaseRequest) (int, string, error) {
	if request == nil {
		request = NewUpgradeReleaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpgradeReleaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


