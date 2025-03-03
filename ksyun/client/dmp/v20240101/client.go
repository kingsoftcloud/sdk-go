package v20240101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2024-01-01"

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

func NewDescribeDefaultMonitorItemsRequest() (request *DescribeDefaultMonitorItemsRequest) {
	request = &DescribeDefaultMonitorItemsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDefaultMonitorItems")
	return
}

func NewDescribeDefaultMonitorItemsResponse() (response *DescribeDefaultMonitorItemsResponse) {
	response = &DescribeDefaultMonitorItemsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDefaultMonitorItems(request *DescribeDefaultMonitorItemsRequest) string {
	return c.DescribeDefaultMonitorItemsWithContext(context.Background(), request)
}

func (c *Client) DescribeDefaultMonitorItemsWithContext(ctx context.Context, request *DescribeDefaultMonitorItemsRequest) string {
	if request == nil {
		request = NewDescribeDefaultMonitorItemsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDefaultMonitorItemsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteMonitorPanelRequest() (request *DeleteMonitorPanelRequest) {
	request = &DeleteMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DeleteMonitorPanel")
	return
}

func NewDeleteMonitorPanelResponse() (response *DeleteMonitorPanelResponse) {
	response = &DeleteMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMonitorPanel(request *DeleteMonitorPanelRequest) string {
	return c.DeleteMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) DeleteMonitorPanelWithContext(ctx context.Context, request *DeleteMonitorPanelRequest) string {
	if request == nil {
		request = NewDeleteMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewOperateMonitorPanelRequest() (request *OperateMonitorPanelRequest) {
	request = &OperateMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "OperateMonitorPanel")
	return
}

func NewOperateMonitorPanelResponse() (response *OperateMonitorPanelResponse) {
	response = &OperateMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperateMonitorPanel(request *OperateMonitorPanelRequest) string {
	return c.OperateMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) OperateMonitorPanelWithContext(ctx context.Context, request *OperateMonitorPanelRequest) string {
	if request == nil {
		request = NewOperateMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperateMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMonitorPanelRequest() (request *DescribeMonitorPanelRequest) {
	request = &DescribeMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeMonitorPanel")
	return
}

func NewDescribeMonitorPanelResponse() (response *DescribeMonitorPanelResponse) {
	response = &DescribeMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorPanel(request *DescribeMonitorPanelRequest) string {
	return c.DescribeMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorPanelWithContext(ctx context.Context, request *DescribeMonitorPanelRequest) string {
	if request == nil {
		request = NewDescribeMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyMonitorPanelInfoRequest() (request *ModifyMonitorPanelInfoRequest) {
	request = &ModifyMonitorPanelInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "ModifyMonitorPanelInfo")
	return
}

func NewModifyMonitorPanelInfoResponse() (response *ModifyMonitorPanelInfoResponse) {
	response = &ModifyMonitorPanelInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyMonitorPanelInfo(request *ModifyMonitorPanelInfoRequest) string {
	return c.ModifyMonitorPanelInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyMonitorPanelInfoWithContext(ctx context.Context, request *ModifyMonitorPanelInfoRequest) string {
	if request == nil {
		request = NewModifyMonitorPanelInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyMonitorPanelInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateMonitorPanelRequest() (request *CreateMonitorPanelRequest) {
	request = &CreateMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "CreateMonitorPanel")
	return
}

func NewCreateMonitorPanelResponse() (response *CreateMonitorPanelResponse) {
	response = &CreateMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMonitorPanel(request *CreateMonitorPanelRequest) string {
	return c.CreateMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) CreateMonitorPanelWithContext(ctx context.Context, request *CreateMonitorPanelRequest) string {
	if request == nil {
		request = NewCreateMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteBatchInstancesRequest() (request *DeleteBatchInstancesRequest) {
	request = &DeleteBatchInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DeleteBatchInstances")
	return
}

func NewDeleteBatchInstancesResponse() (response *DeleteBatchInstancesResponse) {
	response = &DeleteBatchInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBatchInstances(request *DeleteBatchInstancesRequest) string {
	return c.DeleteBatchInstancesWithContext(context.Background(), request)
}

func (c *Client) DeleteBatchInstancesWithContext(ctx context.Context, request *DeleteBatchInstancesRequest) string {
	if request == nil {
		request = NewDeleteBatchInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewInstanceStatisticsRequest() (request *InstanceStatisticsRequest) {
	request = &InstanceStatisticsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "InstanceStatistics")
	return
}

func NewInstanceStatisticsResponse() (response *InstanceStatisticsResponse) {
	response = &InstanceStatisticsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) InstanceStatistics(request *InstanceStatisticsRequest) string {
	return c.InstanceStatisticsWithContext(context.Background(), request)
}

func (c *Client) InstanceStatisticsWithContext(ctx context.Context, request *InstanceStatisticsRequest) string {
	if request == nil {
		request = NewInstanceStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewInstanceStatisticsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMonitorPanelListRequest() (request *DescribeMonitorPanelListRequest) {
	request = &DescribeMonitorPanelListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeMonitorPanelList")
	return
}

func NewDescribeMonitorPanelListResponse() (response *DescribeMonitorPanelListResponse) {
	response = &DescribeMonitorPanelListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorPanelList(request *DescribeMonitorPanelListRequest) string {
	return c.DescribeMonitorPanelListWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorPanelListWithContext(ctx context.Context, request *DescribeMonitorPanelListRequest) string {
	if request == nil {
		request = NewDescribeMonitorPanelListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMonitorPanelListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
	request = &DescribeInstanceListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeInstanceList")
	return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
	response = &DescribeInstanceListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) string {
	return c.DescribeInstanceListWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) string {
	if request == nil {
		request = NewDescribeInstanceListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImportInstanceToDmpRequest() (request *ImportInstanceToDmpRequest) {
	request = &ImportInstanceToDmpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "ImportInstanceToDmp")
	return
}

func NewImportInstanceToDmpResponse() (response *ImportInstanceToDmpResponse) {
	response = &ImportInstanceToDmpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportInstanceToDmp(request *ImportInstanceToDmpRequest) string {
	return c.ImportInstanceToDmpWithContext(context.Background(), request)
}

func (c *Client) ImportInstanceToDmpWithContext(ctx context.Context, request *ImportInstanceToDmpRequest) string {
	if request == nil {
		request = NewImportInstanceToDmpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImportInstanceToDmpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDedicatedClustersRequest() (request *DescribeDedicatedClustersRequest) {
	request = &DescribeDedicatedClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDedicatedClusters")
	return
}

func NewDescribeDedicatedClustersResponse() (response *DescribeDedicatedClustersResponse) {
	response = &DescribeDedicatedClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedClusters(request *DescribeDedicatedClustersRequest) string {
	return c.DescribeDedicatedClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedClustersWithContext(ctx context.Context, request *DescribeDedicatedClustersRequest) string {
	if request == nil {
		request = NewDescribeDedicatedClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDedicatedClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDedicatedHostsRequest() (request *DescribeDedicatedHostsRequest) {
	request = &DescribeDedicatedHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDedicatedHosts")
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

func (c *Client) DescribeDedicatedHostsWithContext(ctx context.Context, request *DescribeDedicatedHostsRequest) string {
	if request == nil {
		request = NewDescribeDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDedicatedHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
