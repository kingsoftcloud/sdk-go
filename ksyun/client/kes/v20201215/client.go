package v20201215

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func (c *Client) ModifyClusterNameWithContext(ctx context.Context, request *ModifyClusterNameRequest) string {
	if request == nil {
		request = NewModifyClusterNameRequest()
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

func (c *Client) ListInstanceGroupsWithContext(ctx context.Context, request *ListInstanceGroupsRequest) string {
	if request == nil {
		request = NewListInstanceGroupsRequest()
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

func (c *Client) ServiceControlWithContext(ctx context.Context, request *ServiceControlRequest) string {
	if request == nil {
		request = NewServiceControlRequest()
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

func (c *Client) ClusterHealthStatisticWithContext(ctx context.Context, request *ClusterHealthStatisticRequest) string {
	if request == nil {
		request = NewClusterHealthStatisticRequest()
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

func (c *Client) CheckClusterHealthWithContext(ctx context.Context, request *CheckClusterHealthRequest) string {
	if request == nil {
		request = NewCheckClusterHealthRequest()
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
