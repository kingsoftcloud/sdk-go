package v20230101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2023-01-01"

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

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCluster(request *CreateClusterRequest) string {
	return c.CreateClusterWithContext(context.Background(), request)
}

func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) string {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeClusters")
	return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
	response = &DescribeClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusters(request *DescribeClustersRequest) string {
	return c.DescribeClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) string {
	if request == nil {
		request = NewDescribeClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DeleteCluster")
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
func NewModifyClusterRequest() (request *ModifyClusterRequest) {
	request = &ModifyClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "ModifyCluster")
	return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
	response = &ModifyClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCluster(request *ModifyClusterRequest) string {
	return c.ModifyClusterWithContext(context.Background(), request)
}

func (c *Client) ModifyClusterWithContext(ctx context.Context, request *ModifyClusterRequest) string {
	if request == nil {
		request = NewModifyClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDownloadClusterConfigRequest() (request *DownloadClusterConfigRequest) {
	request = &DownloadClusterConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DownloadClusterConfig")
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
func NewAddNodesRequest() (request *AddNodesRequest) {
	request = &AddNodesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "AddNodes")
	return
}

func NewAddNodesResponse() (response *AddNodesResponse) {
	response = &AddNodesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddNodes(request *AddNodesRequest) string {
	return c.AddNodesWithContext(context.Background(), request)
}

func (c *Client) AddNodesWithContext(ctx context.Context, request *AddNodesRequest) string {
	if request == nil {
		request = NewAddNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddNodesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNodesRequest() (request *DescribeNodesRequest) {
	request = &DescribeNodesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeNodes")
	return
}

func NewDescribeNodesResponse() (response *DescribeNodesResponse) {
	response = &DescribeNodesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNodes(request *DescribeNodesRequest) string {
	return c.DescribeNodesWithContext(context.Background(), request)
}

func (c *Client) DescribeNodesWithContext(ctx context.Context, request *DescribeNodesRequest) string {
	if request == nil {
		request = NewDescribeNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteNodeRequest() (request *DeleteNodeRequest) {
	request = &DeleteNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DeleteNode")
	return
}

func NewDeleteNodeResponse() (response *DeleteNodeResponse) {
	response = &DeleteNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNode(request *DeleteNodeRequest) string {
	return c.DeleteNodeWithContext(context.Background(), request)
}

func (c *Client) DeleteNodeWithContext(ctx context.Context, request *DeleteNodeRequest) string {
	if request == nil {
		request = NewDeleteNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyNodeRequest() (request *ModifyNodeRequest) {
	request = &ModifyNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "ModifyNode")
	return
}

func NewModifyNodeResponse() (response *ModifyNodeResponse) {
	response = &ModifyNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNode(request *ModifyNodeRequest) string {
	return c.ModifyNodeWithContext(context.Background(), request)
}

func (c *Client) ModifyNodeWithContext(ctx context.Context, request *ModifyNodeRequest) string {
	if request == nil {
		request = NewModifyNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeComponentListRequest() (request *DescribeComponentListRequest) {
	request = &DescribeComponentListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeComponentList")
	return
}

func NewDescribeComponentListResponse() (response *DescribeComponentListResponse) {
	response = &DescribeComponentListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeComponentList(request *DescribeComponentListRequest) string {
	return c.DescribeComponentListWithContext(context.Background(), request)
}

func (c *Client) DescribeComponentListWithContext(ctx context.Context, request *DescribeComponentListRequest) string {
	if request == nil {
		request = NewDescribeComponentListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNodeComponentsRequest() (request *DescribeNodeComponentsRequest) {
	request = &DescribeNodeComponentsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeNodeComponents")
	return
}

func NewDescribeNodeComponentsResponse() (response *DescribeNodeComponentsResponse) {
	response = &DescribeNodeComponentsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNodeComponents(request *DescribeNodeComponentsRequest) string {
	return c.DescribeNodeComponentsWithContext(context.Background(), request)
}

func (c *Client) DescribeNodeComponentsWithContext(ctx context.Context, request *DescribeNodeComponentsRequest) string {
	if request == nil {
		request = NewDescribeNodeComponentsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodeComponentsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNetworkRequest() (request *DescribeNetworkRequest) {
	request = &DescribeNetworkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeNetwork")
	return
}

func NewDescribeNetworkResponse() (response *DescribeNetworkResponse) {
	response = &DescribeNetworkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetwork(request *DescribeNetworkRequest) string {
	return c.DescribeNetworkWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkWithContext(ctx context.Context, request *DescribeNetworkRequest) string {
	if request == nil {
		request = NewDescribeNetworkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeComponentParamsRequest() (request *DescribeComponentParamsRequest) {
	request = &DescribeComponentParamsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeComponentParams")
	return
}

func NewDescribeComponentParamsResponse() (response *DescribeComponentParamsResponse) {
	response = &DescribeComponentParamsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeComponentParams(request *DescribeComponentParamsRequest) string {
	return c.DescribeComponentParamsWithContext(context.Background(), request)
}

func (c *Client) DescribeComponentParamsWithContext(ctx context.Context, request *DescribeComponentParamsRequest) string {
	if request == nil {
		request = NewDescribeComponentParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentParamsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAddonListRequest() (request *DescribeAddonListRequest) {
	request = &DescribeAddonListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeAddonList")
	return
}

func NewDescribeAddonListResponse() (response *DescribeAddonListResponse) {
	response = &DescribeAddonListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAddonList(request *DescribeAddonListRequest) string {
	return c.DescribeAddonListWithContext(context.Background(), request)
}

func (c *Client) DescribeAddonListWithContext(ctx context.Context, request *DescribeAddonListRequest) string {
	if request == nil {
		request = NewDescribeAddonListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddonListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAddonInstancesRequest() (request *DescribeAddonInstancesRequest) {
	request = &DescribeAddonInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeAddonInstances")
	return
}

func NewDescribeAddonInstancesResponse() (response *DescribeAddonInstancesResponse) {
	response = &DescribeAddonInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAddonInstances(request *DescribeAddonInstancesRequest) string {
	return c.DescribeAddonInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeAddonInstancesWithContext(ctx context.Context, request *DescribeAddonInstancesRequest) string {
	if request == nil {
		request = NewDescribeAddonInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddonInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteAddonInstanceRequest() (request *DeleteAddonInstanceRequest) {
	request = &DeleteAddonInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DeleteAddonInstance")
	return
}

func NewDeleteAddonInstanceResponse() (response *DeleteAddonInstanceResponse) {
	response = &DeleteAddonInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAddonInstance(request *DeleteAddonInstanceRequest) string {
	return c.DeleteAddonInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteAddonInstanceWithContext(ctx context.Context, request *DeleteAddonInstanceRequest) string {
	if request == nil {
		request = NewDeleteAddonInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAddonInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateAddonInstanceRequest() (request *CreateAddonInstanceRequest) {
	request = &CreateAddonInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "CreateAddonInstance")
	return
}

func NewCreateAddonInstanceResponse() (response *CreateAddonInstanceResponse) {
	response = &CreateAddonInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAddonInstance(request *CreateAddonInstanceRequest) string {
	return c.CreateAddonInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateAddonInstanceWithContext(ctx context.Context, request *CreateAddonInstanceRequest) string {
	if request == nil {
		request = NewCreateAddonInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAddonInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeEventLogsRequest() (request *DescribeEventLogsRequest) {
	request = &DescribeEventLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeEventLogs")
	return
}

func NewDescribeEventLogsResponse() (response *DescribeEventLogsResponse) {
	response = &DescribeEventLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEventLogs(request *DescribeEventLogsRequest) string {
	return c.DescribeEventLogsWithContext(context.Background(), request)
}

func (c *Client) DescribeEventLogsWithContext(ctx context.Context, request *DescribeEventLogsRequest) string {
	if request == nil {
		request = NewDescribeEventLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEventLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeClusterVersionListRequest() (request *DescribeClusterVersionListRequest) {
	request = &DescribeClusterVersionListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "DescribeClusterVersionList")
	return
}

func NewDescribeClusterVersionListResponse() (response *DescribeClusterVersionListResponse) {
	response = &DescribeClusterVersionListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusterVersionList(request *DescribeClusterVersionListRequest) string {
	return c.DescribeClusterVersionListWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterVersionListWithContext(ctx context.Context, request *DescribeClusterVersionListRequest) string {
	if request == nil {
		request = NewDescribeClusterVersionListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterVersionListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
