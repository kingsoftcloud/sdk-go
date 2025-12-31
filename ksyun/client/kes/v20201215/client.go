package v20201215

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-12-15"

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
	request.Init().WithApiInfo("kes", APIVersion, "DescribeCluster")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "DescribeCluster")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "DescribeCluster")
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
func NewListClustersRequest() (request *ListClustersRequest) {
	request = &ListClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kes", APIVersion, "ListClusters")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ListClusters")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ListClusters")
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
func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
	request = &ModifyClusterNameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kes", APIVersion, "ModifyClusterName")
	return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
	response = &ModifyClusterNameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) string {
	return c.ModifyClusterNameWithContext(context.Background(), request)
}

func (c *Client) ModifyClusterNameSend(request *ModifyClusterNameRequest) (*ModifyClusterNameResponse, error) {
	statusCode, msg, err := c.ModifyClusterNameWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyClusterNameResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyClusterNameWithContext(ctx context.Context, request *ModifyClusterNameRequest) string {
	if request == nil {
		request = NewModifyClusterNameRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ModifyClusterName")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyClusterNameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyClusterNameWithContextV2(ctx context.Context, request *ModifyClusterNameRequest) (int, string, error) {
	if request == nil {
		request = NewModifyClusterNameRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ModifyClusterName")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyClusterNameResponse()
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
	request.Init().WithApiInfo("kes", APIVersion, "LaunchCluster")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "LaunchCluster")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "LaunchCluster")
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
func NewListInstanceGroupsRequest() (request *ListInstanceGroupsRequest) {
	request = &ListInstanceGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kes", APIVersion, "ListInstanceGroups")
	return
}

func NewListInstanceGroupsResponse() (response *ListInstanceGroupsResponse) {
	response = &ListInstanceGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListInstanceGroups(request *ListInstanceGroupsRequest) string {
	return c.ListInstanceGroupsWithContext(context.Background(), request)
}

func (c *Client) ListInstanceGroupsSend(request *ListInstanceGroupsRequest) (*ListInstanceGroupsResponse, error) {
	statusCode, msg, err := c.ListInstanceGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListInstanceGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListInstanceGroupsWithContext(ctx context.Context, request *ListInstanceGroupsRequest) string {
	if request == nil {
		request = NewListInstanceGroupsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ListInstanceGroups")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListInstanceGroupsWithContextV2(ctx context.Context, request *ListInstanceGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewListInstanceGroupsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ListInstanceGroups")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListInstanceGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewServiceControlRequest() (request *ServiceControlRequest) {
	request = &ServiceControlRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kes", APIVersion, "ServiceControl")
	return
}

func NewServiceControlResponse() (response *ServiceControlResponse) {
	response = &ServiceControlResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ServiceControl(request *ServiceControlRequest) string {
	return c.ServiceControlWithContext(context.Background(), request)
}

func (c *Client) ServiceControlSend(request *ServiceControlRequest) (*ServiceControlResponse, error) {
	statusCode, msg, err := c.ServiceControlWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ServiceControlResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ServiceControlWithContext(ctx context.Context, request *ServiceControlRequest) string {
	if request == nil {
		request = NewServiceControlRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ServiceControl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewServiceControlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ServiceControlWithContextV2(ctx context.Context, request *ServiceControlRequest) (int, string, error) {
	if request == nil {
		request = NewServiceControlRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ServiceControl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewServiceControlResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewClusterHealthStatisticRequest() (request *ClusterHealthStatisticRequest) {
	request = &ClusterHealthStatisticRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kes", APIVersion, "ClusterHealthStatistic")
	return
}

func NewClusterHealthStatisticResponse() (response *ClusterHealthStatisticResponse) {
	response = &ClusterHealthStatisticResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ClusterHealthStatistic(request *ClusterHealthStatisticRequest) string {
	return c.ClusterHealthStatisticWithContext(context.Background(), request)
}

func (c *Client) ClusterHealthStatisticSend(request *ClusterHealthStatisticRequest) (*ClusterHealthStatisticResponse, error) {
	statusCode, msg, err := c.ClusterHealthStatisticWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ClusterHealthStatisticResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ClusterHealthStatisticWithContext(ctx context.Context, request *ClusterHealthStatisticRequest) string {
	if request == nil {
		request = NewClusterHealthStatisticRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ClusterHealthStatistic")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewClusterHealthStatisticResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ClusterHealthStatisticWithContextV2(ctx context.Context, request *ClusterHealthStatisticRequest) (int, string, error) {
	if request == nil {
		request = NewClusterHealthStatisticRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "ClusterHealthStatistic")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewClusterHealthStatisticResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCheckClusterHealthRequest() (request *CheckClusterHealthRequest) {
	request = &CheckClusterHealthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kes", APIVersion, "CheckClusterHealth")
	return
}

func NewCheckClusterHealthResponse() (response *CheckClusterHealthResponse) {
	response = &CheckClusterHealthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CheckClusterHealth(request *CheckClusterHealthRequest) string {
	return c.CheckClusterHealthWithContext(context.Background(), request)
}

func (c *Client) CheckClusterHealthSend(request *CheckClusterHealthRequest) (*CheckClusterHealthResponse, error) {
	statusCode, msg, err := c.CheckClusterHealthWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CheckClusterHealthResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CheckClusterHealthWithContext(ctx context.Context, request *CheckClusterHealthRequest) string {
	if request == nil {
		request = NewCheckClusterHealthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "CheckClusterHealth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckClusterHealthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CheckClusterHealthWithContextV2(ctx context.Context, request *CheckClusterHealthRequest) (int, string, error) {
	if request == nil {
		request = NewCheckClusterHealthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kes", APIVersion, "CheckClusterHealth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCheckClusterHealthResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
