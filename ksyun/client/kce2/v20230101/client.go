package v20230101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) CreateClusterSend(request *CreateClusterRequest) (*CreateClusterResponse, error) {
	statusCode, msg, err := c.CreateClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateClusterWithContextV2(ctx context.Context, request *CreateClusterRequest) (int, string, error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeClustersSend(request *DescribeClustersRequest) (*DescribeClustersResponse, error) {
	statusCode, msg, err := c.DescribeClustersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClustersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeClustersWithContextV2(ctx context.Context, request *DescribeClustersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeClustersResponse()
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

func (c *Client) ModifyClusterSend(request *ModifyClusterRequest) (*ModifyClusterResponse, error) {
	statusCode, msg, err := c.ModifyClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyClusterWithContextV2(ctx context.Context, request *ModifyClusterRequest) (int, string, error) {
	if request == nil {
		request = NewModifyClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNodesSend(request *DescribeNodesRequest) (*DescribeNodesResponse, error) {
	statusCode, msg, err := c.DescribeNodesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNodesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeNodesWithContextV2(ctx context.Context, request *DescribeNodesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteNodeSend(request *DeleteNodeRequest) (*DeleteNodeResponse, error) {
	statusCode, msg, err := c.DeleteNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteNodeWithContextV2(ctx context.Context, request *DeleteNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyNodeSend(request *ModifyNodeRequest) (*ModifyNodeResponse, error) {
	statusCode, msg, err := c.ModifyNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyNodeWithContextV2(ctx context.Context, request *ModifyNodeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeComponentListSend(request *DescribeComponentListRequest) (*DescribeComponentListResponse, error) {
	statusCode, msg, err := c.DescribeComponentListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeComponentListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeComponentListWithContextV2(ctx context.Context, request *DescribeComponentListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeComponentListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNodeComponentsSend(request *DescribeNodeComponentsRequest) (*DescribeNodeComponentsResponse, error) {
	statusCode, msg, err := c.DescribeNodeComponentsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNodeComponentsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeNodeComponentsWithContextV2(ctx context.Context, request *DescribeNodeComponentsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNodeComponentsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodeComponentsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNetworkSend(request *DescribeNetworkRequest) (*DescribeNetworkResponse, error) {
	statusCode, msg, err := c.DescribeNetworkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNetworkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeNetworkWithContextV2(ctx context.Context, request *DescribeNetworkRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNetworkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeComponentParamsSend(request *DescribeComponentParamsRequest) (*DescribeComponentParamsResponse, error) {
	statusCode, msg, err := c.DescribeComponentParamsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeComponentParamsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeComponentParamsWithContextV2(ctx context.Context, request *DescribeComponentParamsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeComponentParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentParamsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeEventLogsSend(request *DescribeEventLogsRequest) (*DescribeEventLogsResponse, error) {
	statusCode, msg, err := c.DescribeEventLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEventLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeEventLogsWithContextV2(ctx context.Context, request *DescribeEventLogsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEventLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEventLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeClusterVersionListSend(request *DescribeClusterVersionListRequest) (*DescribeClusterVersionListResponse, error) {
	statusCode, msg, err := c.DescribeClusterVersionListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeClusterVersionListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeClusterVersionListWithContextV2(ctx context.Context, request *DescribeClusterVersionListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterVersionListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterVersionListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddKecNodesRequest() (request *AddKecNodesRequest) {
	request = &AddKecNodesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "AddKecNodes")
	return
}

func NewAddKecNodesResponse() (response *AddKecNodesResponse) {
	response = &AddKecNodesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddKecNodes(request *AddKecNodesRequest) string {
	return c.AddKecNodesWithContext(context.Background(), request)
}

func (c *Client) AddKecNodesSend(request *AddKecNodesRequest) (*AddKecNodesResponse, error) {
	statusCode, msg, err := c.AddKecNodesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddKecNodesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddKecNodesWithContext(ctx context.Context, request *AddKecNodesRequest) string {
	if request == nil {
		request = NewAddKecNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddKecNodesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddKecNodesWithContextV2(ctx context.Context, request *AddKecNodesRequest) (int, string, error) {
	if request == nil {
		request = NewAddKecNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddKecNodesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddEpcNodesRequest() (request *AddEpcNodesRequest) {
	request = &AddEpcNodesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "AddEpcNodes")
	return
}

func NewAddEpcNodesResponse() (response *AddEpcNodesResponse) {
	response = &AddEpcNodesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddEpcNodes(request *AddEpcNodesRequest) string {
	return c.AddEpcNodesWithContext(context.Background(), request)
}

func (c *Client) AddEpcNodesSend(request *AddEpcNodesRequest) (*AddEpcNodesResponse, error) {
	statusCode, msg, err := c.AddEpcNodesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddEpcNodesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddEpcNodesWithContext(ctx context.Context, request *AddEpcNodesRequest) string {
	if request == nil {
		request = NewAddEpcNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddEpcNodesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddEpcNodesWithContextV2(ctx context.Context, request *AddEpcNodesRequest) (int, string, error) {
	if request == nil {
		request = NewAddEpcNodesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddEpcNodesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
