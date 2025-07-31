package v20210902

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2021-09-02"

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

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DescribeCluster")
	return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCluster(request *DescribeClusterRequest) string {
	return c.DescribeClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterSend(request *DescribeClusterRequest) (*DescribeClusterResponse, error) {
	statusCode, msg, err := c.DescribeClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClusterWithContext(ctx context.Context, request *DescribeClusterRequest) string {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClusterWithContextV2(ctx context.Context, request *DescribeClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewLaunchClusterRequest() (request *LaunchClusterRequest) {
	request = &LaunchClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "LaunchCluster")
	return
}

func NewLaunchClusterResponse() (response *LaunchClusterResponse) {
	response = &LaunchClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) LaunchCluster(request *LaunchClusterRequest) string {
	return c.LaunchClusterWithContext(context.Background(), request)
}

func (c *Client) LaunchClusterSend(request *LaunchClusterRequest) (*LaunchClusterResponse, error) {
	statusCode, msg, err := c.LaunchClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct LaunchClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) LaunchClusterWithContext(ctx context.Context, request *LaunchClusterRequest) string {
	if request == nil {
		request = NewLaunchClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewLaunchClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) LaunchClusterWithContextV2(ctx context.Context, request *LaunchClusterRequest) (int, string, error) {
	if request == nil {
		request = NewLaunchClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewLaunchClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewScaleInInstanceGroupsRequest() (request *ScaleInInstanceGroupsRequest) {
	request = &ScaleInInstanceGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ScaleInInstanceGroups")
	return
}

func NewScaleInInstanceGroupsResponse() (response *ScaleInInstanceGroupsResponse) {
	response = &ScaleInInstanceGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ScaleInInstanceGroups(request *ScaleInInstanceGroupsRequest) string {
	return c.ScaleInInstanceGroupsWithContext(context.Background(), request)
}

func (c *Client) ScaleInInstanceGroupsSend(request *ScaleInInstanceGroupsRequest) (*ScaleInInstanceGroupsResponse, error) {
	statusCode, msg, err := c.ScaleInInstanceGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ScaleInInstanceGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ScaleInInstanceGroupsWithContext(ctx context.Context, request *ScaleInInstanceGroupsRequest) string {
	if request == nil {
		request = NewScaleInInstanceGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewScaleInInstanceGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ScaleInInstanceGroupsWithContextV2(ctx context.Context, request *ScaleInInstanceGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewScaleInInstanceGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewScaleInInstanceGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewScaleOutInstanceGroupsRequest() (request *ScaleOutInstanceGroupsRequest) {
	request = &ScaleOutInstanceGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ScaleOutInstanceGroups")
	return
}

func NewScaleOutInstanceGroupsResponse() (response *ScaleOutInstanceGroupsResponse) {
	response = &ScaleOutInstanceGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ScaleOutInstanceGroups(request *ScaleOutInstanceGroupsRequest) string {
	return c.ScaleOutInstanceGroupsWithContext(context.Background(), request)
}

func (c *Client) ScaleOutInstanceGroupsSend(request *ScaleOutInstanceGroupsRequest) (*ScaleOutInstanceGroupsResponse, error) {
	statusCode, msg, err := c.ScaleOutInstanceGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ScaleOutInstanceGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ScaleOutInstanceGroupsWithContext(ctx context.Context, request *ScaleOutInstanceGroupsRequest) string {
	if request == nil {
		request = NewScaleOutInstanceGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewScaleOutInstanceGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ScaleOutInstanceGroupsWithContextV2(ctx context.Context, request *ScaleOutInstanceGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewScaleOutInstanceGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewScaleOutInstanceGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeClusterInfoRequest() (request *DescribeClusterInfoRequest) {
	request = &DescribeClusterInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DescribeClusterInfo")
	return
}

func NewDescribeClusterInfoResponse() (response *DescribeClusterInfoResponse) {
	response = &DescribeClusterInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusterInfo(request *DescribeClusterInfoRequest) string {
	return c.DescribeClusterInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterInfoSend(request *DescribeClusterInfoRequest) (*DescribeClusterInfoResponse, error) {
	statusCode, msg, err := c.DescribeClusterInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClusterInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClusterInfoWithContext(ctx context.Context, request *DescribeClusterInfoRequest) string {
	if request == nil {
		request = NewDescribeClusterInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClusterInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClusterInfoWithContextV2(ctx context.Context, request *DescribeClusterInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClusterInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListServiceStatusRequest() (request *ListServiceStatusRequest) {
	request = &ListServiceStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListServiceStatus")
	return
}

func NewListServiceStatusResponse() (response *ListServiceStatusResponse) {
	response = &ListServiceStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListServiceStatus(request *ListServiceStatusRequest) string {
	return c.ListServiceStatusWithContext(context.Background(), request)
}

func (c *Client) ListServiceStatusSend(request *ListServiceStatusRequest) (*ListServiceStatusResponse, error) {
	statusCode, msg, err := c.ListServiceStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListServiceStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListServiceStatusWithContext(ctx context.Context, request *ListServiceStatusRequest) string {
	if request == nil {
		request = NewListServiceStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListServiceStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListServiceStatusWithContextV2(ctx context.Context, request *ListServiceStatusRequest) (int, string, error) {
	if request == nil {
		request = NewListServiceStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListServiceStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListClustersRequest() (request *ListClustersRequest) {
	request = &ListClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListClusters")
	return
}

func NewListClustersResponse() (response *ListClustersResponse) {
	response = &ListClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListClusters(request *ListClustersRequest) string {
	return c.ListClustersWithContext(context.Background(), request)
}

func (c *Client) ListClustersSend(request *ListClustersRequest) (*ListClustersResponse, error) {
	statusCode, msg, err := c.ListClustersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListClustersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListClustersWithContext(ctx context.Context, request *ListClustersRequest) string {
	if request == nil {
		request = NewListClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListClustersWithContextV2(ctx context.Context, request *ListClustersRequest) (int, string, error) {
	if request == nil {
		request = NewListClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListClustersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListClusterVersionsRequest() (request *ListClusterVersionsRequest) {
	request = &ListClusterVersionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListClusterVersions")
	return
}

func NewListClusterVersionsResponse() (response *ListClusterVersionsResponse) {
	response = &ListClusterVersionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListClusterVersions(request *ListClusterVersionsRequest) string {
	return c.ListClusterVersionsWithContext(context.Background(), request)
}

func (c *Client) ListClusterVersionsSend(request *ListClusterVersionsRequest) (*ListClusterVersionsResponse, error) {
	statusCode, msg, err := c.ListClusterVersionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListClusterVersionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListClusterVersionsWithContext(ctx context.Context, request *ListClusterVersionsRequest) string {
	if request == nil {
		request = NewListClusterVersionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListClusterVersionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListClusterVersionsWithContextV2(ctx context.Context, request *ListClusterVersionsRequest) (int, string, error) {
	if request == nil {
		request = NewListClusterVersionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListClusterVersionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeService(request *DescribeServiceRequest) string {
	return c.DescribeServiceWithContext(context.Background(), request)
}

func (c *Client) DescribeServiceSend(request *DescribeServiceRequest) (*DescribeServiceResponse, error) {
	statusCode, msg, err := c.DescribeServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeServiceWithContext(ctx context.Context, request *DescribeServiceRequest) string {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeServiceWithContextV2(ctx context.Context, request *DescribeServiceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListConfigurationsRequest() (request *ListConfigurationsRequest) {
	request = &ListConfigurationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListConfigurations")
	return
}

func NewListConfigurationsResponse() (response *ListConfigurationsResponse) {
	response = &ListConfigurationsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListConfigurations(request *ListConfigurationsRequest) string {
	return c.ListConfigurationsWithContext(context.Background(), request)
}

func (c *Client) ListConfigurationsSend(request *ListConfigurationsRequest) (*ListConfigurationsResponse, error) {
	statusCode, msg, err := c.ListConfigurationsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListConfigurationsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListConfigurationsWithContext(ctx context.Context, request *ListConfigurationsRequest) string {
	if request == nil {
		request = NewListConfigurationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListConfigurationsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListConfigurationsWithContextV2(ctx context.Context, request *ListConfigurationsRequest) (int, string, error) {
	if request == nil {
		request = NewListConfigurationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListConfigurationsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListConfigurationHistoryRequest() (request *ListConfigurationHistoryRequest) {
	request = &ListConfigurationHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListConfigurationHistory")
	return
}

func NewListConfigurationHistoryResponse() (response *ListConfigurationHistoryResponse) {
	response = &ListConfigurationHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListConfigurationHistory(request *ListConfigurationHistoryRequest) string {
	return c.ListConfigurationHistoryWithContext(context.Background(), request)
}

func (c *Client) ListConfigurationHistorySend(request *ListConfigurationHistoryRequest) (*ListConfigurationHistoryResponse, error) {
	statusCode, msg, err := c.ListConfigurationHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListConfigurationHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListConfigurationHistoryWithContext(ctx context.Context, request *ListConfigurationHistoryRequest) string {
	if request == nil {
		request = NewListConfigurationHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListConfigurationHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListConfigurationHistoryWithContextV2(ctx context.Context, request *ListConfigurationHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewListConfigurationHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListConfigurationHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTagValuesRequest() (request *ListTagValuesRequest) {
	request = &ListTagValuesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListTagValues")
	return
}

func NewListTagValuesResponse() (response *ListTagValuesResponse) {
	response = &ListTagValuesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTagValues(request *ListTagValuesRequest) string {
	return c.ListTagValuesWithContext(context.Background(), request)
}

func (c *Client) ListTagValuesSend(request *ListTagValuesRequest) (*ListTagValuesResponse, error) {
	statusCode, msg, err := c.ListTagValuesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListTagValuesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest) string {
	if request == nil {
		request = NewListTagValuesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTagValuesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTagValuesWithContextV2(ctx context.Context, request *ListTagValuesRequest) (int, string, error) {
	if request == nil {
		request = NewListTagValuesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTagValuesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTagKeysRequest() (request *ListTagKeysRequest) {
	request = &ListTagKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListTagKeys")
	return
}

func NewListTagKeysResponse() (response *ListTagKeysResponse) {
	response = &ListTagKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTagKeys(request *ListTagKeysRequest) string {
	return c.ListTagKeysWithContext(context.Background(), request)
}

func (c *Client) ListTagKeysSend(request *ListTagKeysRequest) (*ListTagKeysResponse, error) {
	statusCode, msg, err := c.ListTagKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListTagKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest) string {
	if request == nil {
		request = NewListTagKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTagKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTagKeysWithContextV2(ctx context.Context, request *ListTagKeysRequest) (int, string, error) {
	if request == nil {
		request = NewListTagKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListTagKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBindTagsRequest() (request *BindTagsRequest) {
	request = &BindTagsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "BindTags")
	return
}

func NewBindTagsResponse() (response *BindTagsResponse) {
	response = &BindTagsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindTags(request *BindTagsRequest) string {
	return c.BindTagsWithContext(context.Background(), request)
}

func (c *Client) BindTagsSend(request *BindTagsRequest) (*BindTagsResponse, error) {
	statusCode, msg, err := c.BindTagsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct BindTagsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BindTagsWithContext(ctx context.Context, request *BindTagsRequest) string {
	if request == nil {
		request = NewBindTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBindTagsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BindTagsWithContextV2(ctx context.Context, request *BindTagsRequest) (int, string, error) {
	if request == nil {
		request = NewBindTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBindTagsResponse()
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
	request.Init().WithApiInfo("kmr", APIVersion, "StartInstances")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewStartInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestartInstancesRequest() (request *RestartInstancesRequest) {
	request = &RestartInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "RestartInstances")
	return
}

func NewRestartInstancesResponse() (response *RestartInstancesResponse) {
	response = &RestartInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestartInstances(request *RestartInstancesRequest) string {
	return c.RestartInstancesWithContext(context.Background(), request)
}

func (c *Client) RestartInstancesSend(request *RestartInstancesRequest) (*RestartInstancesResponse, error) {
	statusCode, msg, err := c.RestartInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RestartInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestartInstancesWithContext(ctx context.Context, request *RestartInstancesRequest) string {
	if request == nil {
		request = NewRestartInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestartInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestartInstancesWithContextV2(ctx context.Context, request *RestartInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewRestartInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestartInstancesResponse()
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
	request.Init().WithApiInfo("kmr", APIVersion, "StopInstances")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewStopInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListScaleStrategyRequest() (request *ListScaleStrategyRequest) {
	request = &ListScaleStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListScaleStrategy")
	return
}

func NewListScaleStrategyResponse() (response *ListScaleStrategyResponse) {
	response = &ListScaleStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListScaleStrategy(request *ListScaleStrategyRequest) string {
	return c.ListScaleStrategyWithContext(context.Background(), request)
}

func (c *Client) ListScaleStrategySend(request *ListScaleStrategyRequest) (*ListScaleStrategyResponse, error) {
	statusCode, msg, err := c.ListScaleStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListScaleStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListScaleStrategyWithContext(ctx context.Context, request *ListScaleStrategyRequest) string {
	if request == nil {
		request = NewListScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListScaleStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListScaleStrategyWithContextV2(ctx context.Context, request *ListScaleStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewListScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListScaleStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyScaleStrategyRequest() (request *ModifyScaleStrategyRequest) {
	request = &ModifyScaleStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ModifyScaleStrategy")
	return
}

func NewModifyScaleStrategyResponse() (response *ModifyScaleStrategyResponse) {
	response = &ModifyScaleStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyScaleStrategy(request *ModifyScaleStrategyRequest) string {
	return c.ModifyScaleStrategyWithContext(context.Background(), request)
}

func (c *Client) ModifyScaleStrategySend(request *ModifyScaleStrategyRequest) (*ModifyScaleStrategyResponse, error) {
	statusCode, msg, err := c.ModifyScaleStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyScaleStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyScaleStrategyWithContext(ctx context.Context, request *ModifyScaleStrategyRequest) string {
	if request == nil {
		request = NewModifyScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyScaleStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyScaleStrategyWithContextV2(ctx context.Context, request *ModifyScaleStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyScaleStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteScaleStrategyRequest() (request *DeleteScaleStrategyRequest) {
	request = &DeleteScaleStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "DeleteScaleStrategy")
	return
}

func NewDeleteScaleStrategyResponse() (response *DeleteScaleStrategyResponse) {
	response = &DeleteScaleStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteScaleStrategy(request *DeleteScaleStrategyRequest) string {
	return c.DeleteScaleStrategyWithContext(context.Background(), request)
}

func (c *Client) DeleteScaleStrategySend(request *DeleteScaleStrategyRequest) (*DeleteScaleStrategyResponse, error) {
	statusCode, msg, err := c.DeleteScaleStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteScaleStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteScaleStrategyWithContext(ctx context.Context, request *DeleteScaleStrategyRequest) string {
	if request == nil {
		request = NewDeleteScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteScaleStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteScaleStrategyWithContextV2(ctx context.Context, request *DeleteScaleStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteScaleStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListScaleHistoryRequest() (request *ListScaleHistoryRequest) {
	request = &ListScaleHistoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ListScaleHistory")
	return
}

func NewListScaleHistoryResponse() (response *ListScaleHistoryResponse) {
	response = &ListScaleHistoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListScaleHistory(request *ListScaleHistoryRequest) string {
	return c.ListScaleHistoryWithContext(context.Background(), request)
}

func (c *Client) ListScaleHistorySend(request *ListScaleHistoryRequest) (*ListScaleHistoryResponse, error) {
	statusCode, msg, err := c.ListScaleHistoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListScaleHistoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListScaleHistoryWithContext(ctx context.Context, request *ListScaleHistoryRequest) string {
	if request == nil {
		request = NewListScaleHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListScaleHistoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListScaleHistoryWithContextV2(ctx context.Context, request *ListScaleHistoryRequest) (int, string, error) {
	if request == nil {
		request = NewListScaleHistoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListScaleHistoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddLoadBasedScaleStrategyRequest() (request *AddLoadBasedScaleStrategyRequest) {
	request = &AddLoadBasedScaleStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "AddLoadBasedScaleStrategy")
	return
}

func NewAddLoadBasedScaleStrategyResponse() (response *AddLoadBasedScaleStrategyResponse) {
	response = &AddLoadBasedScaleStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddLoadBasedScaleStrategy(request *AddLoadBasedScaleStrategyRequest) string {
	return c.AddLoadBasedScaleStrategyWithContext(context.Background(), request)
}

func (c *Client) AddLoadBasedScaleStrategySend(request *AddLoadBasedScaleStrategyRequest) (*AddLoadBasedScaleStrategyResponse, error) {
	statusCode, msg, err := c.AddLoadBasedScaleStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddLoadBasedScaleStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddLoadBasedScaleStrategyWithContext(ctx context.Context, request *AddLoadBasedScaleStrategyRequest) string {
	if request == nil {
		request = NewAddLoadBasedScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddLoadBasedScaleStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddLoadBasedScaleStrategyWithContextV2(ctx context.Context, request *AddLoadBasedScaleStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewAddLoadBasedScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddLoadBasedScaleStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyLoadBasedScaleStrategyRequest() (request *ModifyLoadBasedScaleStrategyRequest) {
	request = &ModifyLoadBasedScaleStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kmr", APIVersion, "ModifyLoadBasedScaleStrategy")
	return
}

func NewModifyLoadBasedScaleStrategyResponse() (response *ModifyLoadBasedScaleStrategyResponse) {
	response = &ModifyLoadBasedScaleStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyLoadBasedScaleStrategy(request *ModifyLoadBasedScaleStrategyRequest) string {
	return c.ModifyLoadBasedScaleStrategyWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBasedScaleStrategySend(request *ModifyLoadBasedScaleStrategyRequest) (*ModifyLoadBasedScaleStrategyResponse, error) {
	statusCode, msg, err := c.ModifyLoadBasedScaleStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyLoadBasedScaleStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyLoadBasedScaleStrategyWithContext(ctx context.Context, request *ModifyLoadBasedScaleStrategyRequest) string {
	if request == nil {
		request = NewModifyLoadBasedScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyLoadBasedScaleStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyLoadBasedScaleStrategyWithContextV2(ctx context.Context, request *ModifyLoadBasedScaleStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyLoadBasedScaleStrategyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyLoadBasedScaleStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
