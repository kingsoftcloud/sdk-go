package v20230306

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2023-03-06"

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

func NewCreatePrometheusInstanceRequest() (request *CreatePrometheusInstanceRequest) {
	request = &CreatePrometheusInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreatePrometheusInstance")
	return
}

func NewCreatePrometheusInstanceResponse() (response *CreatePrometheusInstanceResponse) {
	response = &CreatePrometheusInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePrometheusInstance(request *CreatePrometheusInstanceRequest) string {
	return c.CreatePrometheusInstanceWithContext(context.Background(), request)
}

func (c *Client) CreatePrometheusInstanceWithContext(ctx context.Context, request *CreatePrometheusInstanceRequest) string {
	if request == nil {
		request = NewCreatePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrometheusInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribePrometheusInstanceRequest() (request *DescribePrometheusInstanceRequest) {
	request = &DescribePrometheusInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribePrometheusInstance")
	return
}

func NewDescribePrometheusInstanceResponse() (response *DescribePrometheusInstanceResponse) {
	response = &DescribePrometheusInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrometheusInstance(request *DescribePrometheusInstanceRequest) string {
	return c.DescribePrometheusInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribePrometheusInstanceWithContext(ctx context.Context, request *DescribePrometheusInstanceRequest) string {
	if request == nil {
		request = NewDescribePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrometheusInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdatePrometheusInstanceRequest() (request *UpdatePrometheusInstanceRequest) {
	request = &UpdatePrometheusInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdatePrometheusInstance")
	return
}

func NewUpdatePrometheusInstanceResponse() (response *UpdatePrometheusInstanceResponse) {
	response = &UpdatePrometheusInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePrometheusInstance(request *UpdatePrometheusInstanceRequest) string {
	return c.UpdatePrometheusInstanceWithContext(context.Background(), request)
}

func (c *Client) UpdatePrometheusInstanceWithContext(ctx context.Context, request *UpdatePrometheusInstanceRequest) string {
	if request == nil {
		request = NewUpdatePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdatePrometheusInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeletePrometheusInstanceRequest() (request *DeletePrometheusInstanceRequest) {
	request = &DeletePrometheusInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeletePrometheusInstance")
	return
}

func NewDeletePrometheusInstanceResponse() (response *DeletePrometheusInstanceResponse) {
	response = &DeletePrometheusInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePrometheusInstance(request *DeletePrometheusInstanceRequest) string {
	return c.DeletePrometheusInstanceWithContext(context.Background(), request)
}

func (c *Client) DeletePrometheusInstanceWithContext(ctx context.Context, request *DeletePrometheusInstanceRequest) string {
	if request == nil {
		request = NewDeletePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrometheusInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewEnableGrafanaRequest() (request *EnableGrafanaRequest) {
	request = &EnableGrafanaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "EnableGrafana")
	return
}

func NewEnableGrafanaResponse() (response *EnableGrafanaResponse) {
	response = &EnableGrafanaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableGrafana(request *EnableGrafanaRequest) string {
	return c.EnableGrafanaWithContext(context.Background(), request)
}

func (c *Client) EnableGrafanaWithContext(ctx context.Context, request *EnableGrafanaRequest) string {
	if request == nil {
		request = NewEnableGrafanaRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableGrafanaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateGrafanaPasswordRequest() (request *UpdateGrafanaPasswordRequest) {
	request = &UpdateGrafanaPasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdateGrafanaPassword")
	return
}

func NewUpdateGrafanaPasswordResponse() (response *UpdateGrafanaPasswordResponse) {
	response = &UpdateGrafanaPasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateGrafanaPassword(request *UpdateGrafanaPasswordRequest) string {
	return c.UpdateGrafanaPasswordWithContext(context.Background(), request)
}

func (c *Client) UpdateGrafanaPasswordWithContext(ctx context.Context, request *UpdateGrafanaPasswordRequest) string {
	if request == nil {
		request = NewUpdateGrafanaPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateGrafanaPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewEnableGrafanaInternetRequest() (request *EnableGrafanaInternetRequest) {
	request = &EnableGrafanaInternetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "EnableGrafanaInternet")
	return
}

func NewEnableGrafanaInternetResponse() (response *EnableGrafanaInternetResponse) {
	response = &EnableGrafanaInternetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableGrafanaInternet(request *EnableGrafanaInternetRequest) string {
	return c.EnableGrafanaInternetWithContext(context.Background(), request)
}

func (c *Client) EnableGrafanaInternetWithContext(ctx context.Context, request *EnableGrafanaInternetRequest) string {
	if request == nil {
		request = NewEnableGrafanaInternetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableGrafanaInternetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeGrafanaWhiteListRequest() (request *DescribeGrafanaWhiteListRequest) {
	request = &DescribeGrafanaWhiteListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeGrafanaWhiteList")
	return
}

func NewDescribeGrafanaWhiteListResponse() (response *DescribeGrafanaWhiteListResponse) {
	response = &DescribeGrafanaWhiteListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeGrafanaWhiteList(request *DescribeGrafanaWhiteListRequest) string {
	return c.DescribeGrafanaWhiteListWithContext(context.Background(), request)
}

func (c *Client) DescribeGrafanaWhiteListWithContext(ctx context.Context, request *DescribeGrafanaWhiteListRequest) string {
	if request == nil {
		request = NewDescribeGrafanaWhiteListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeGrafanaWhiteListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateGrafanaWhiteListRequest() (request *UpdateGrafanaWhiteListRequest) {
	request = &UpdateGrafanaWhiteListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdateGrafanaWhiteList")
	return
}

func NewUpdateGrafanaWhiteListResponse() (response *UpdateGrafanaWhiteListResponse) {
	response = &UpdateGrafanaWhiteListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateGrafanaWhiteList(request *UpdateGrafanaWhiteListRequest) string {
	return c.UpdateGrafanaWhiteListWithContext(context.Background(), request)
}

func (c *Client) UpdateGrafanaWhiteListWithContext(ctx context.Context, request *UpdateGrafanaWhiteListRequest) string {
	if request == nil {
		request = NewUpdateGrafanaWhiteListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateGrafanaWhiteListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAssociateClusterRequest() (request *AssociateClusterRequest) {
	request = &AssociateClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "AssociateCluster")
	return
}

func NewAssociateClusterResponse() (response *AssociateClusterResponse) {
	response = &AssociateClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateCluster(request *AssociateClusterRequest) string {
	return c.AssociateClusterWithContext(context.Background(), request)
}

func (c *Client) AssociateClusterWithContext(ctx context.Context, request *AssociateClusterRequest) string {
	if request == nil {
		request = NewAssociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisassociateClusterRequest() (request *DisassociateClusterRequest) {
	request = &DisassociateClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DisassociateCluster")
	return
}

func NewDisassociateClusterResponse() (response *DisassociateClusterResponse) {
	response = &DisassociateClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateCluster(request *DisassociateClusterRequest) string {
	return c.DisassociateClusterWithContext(context.Background(), request)
}

func (c *Client) DisassociateClusterWithContext(ctx context.Context, request *DisassociateClusterRequest) string {
	if request == nil {
		request = NewDisassociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAssociateClusterListRequest() (request *DescribeAssociateClusterListRequest) {
	request = &DescribeAssociateClusterListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeAssociateClusterList")
	return
}

func NewDescribeAssociateClusterListResponse() (response *DescribeAssociateClusterListResponse) {
	response = &DescribeAssociateClusterListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAssociateClusterList(request *DescribeAssociateClusterListRequest) string {
	return c.DescribeAssociateClusterListWithContext(context.Background(), request)
}

func (c *Client) DescribeAssociateClusterListWithContext(ctx context.Context, request *DescribeAssociateClusterListRequest) string {
	if request == nil {
		request = NewDescribeAssociateClusterListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAssociateClusterListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMonitorListRequest() (request *DescribeMonitorListRequest) {
	request = &DescribeMonitorListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeMonitorList")
	return
}

func NewDescribeMonitorListResponse() (response *DescribeMonitorListResponse) {
	response = &DescribeMonitorListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorList(request *DescribeMonitorListRequest) string {
	return c.DescribeMonitorListWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorListWithContext(ctx context.Context, request *DescribeMonitorListRequest) string {
	if request == nil {
		request = NewDescribeMonitorListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMonitorCollectionConfigRequest() (request *DescribeMonitorCollectionConfigRequest) {
	request = &DescribeMonitorCollectionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeMonitorCollectionConfig")
	return
}

func NewDescribeMonitorCollectionConfigResponse() (response *DescribeMonitorCollectionConfigResponse) {
	response = &DescribeMonitorCollectionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorCollectionConfig(request *DescribeMonitorCollectionConfigRequest) string {
	return c.DescribeMonitorCollectionConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorCollectionConfigWithContext(ctx context.Context, request *DescribeMonitorCollectionConfigRequest) string {
	if request == nil {
		request = NewDescribeMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorCollectionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateMonitorCollectionConfigRequest() (request *UpdateMonitorCollectionConfigRequest) {
	request = &UpdateMonitorCollectionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdateMonitorCollectionConfig")
	return
}

func NewUpdateMonitorCollectionConfigResponse() (response *UpdateMonitorCollectionConfigResponse) {
	response = &UpdateMonitorCollectionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateMonitorCollectionConfig(request *UpdateMonitorCollectionConfigRequest) string {
	return c.UpdateMonitorCollectionConfigWithContext(context.Background(), request)
}

func (c *Client) UpdateMonitorCollectionConfigWithContext(ctx context.Context, request *UpdateMonitorCollectionConfigRequest) string {
	if request == nil {
		request = NewUpdateMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateMonitorCollectionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeMonitorMetricsListRequest() (request *DescribeMonitorMetricsListRequest) {
	request = &DescribeMonitorMetricsListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeMonitorMetricsList")
	return
}

func NewDescribeMonitorMetricsListResponse() (response *DescribeMonitorMetricsListResponse) {
	response = &DescribeMonitorMetricsListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorMetricsList(request *DescribeMonitorMetricsListRequest) string {
	return c.DescribeMonitorMetricsListWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorMetricsListWithContext(ctx context.Context, request *DescribeMonitorMetricsListRequest) string {
	if request == nil {
		request = NewDescribeMonitorMetricsListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorMetricsListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeTargetsListRequest() (request *DescribeTargetsListRequest) {
	request = &DescribeTargetsListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeTargetsList")
	return
}

func NewDescribeTargetsListResponse() (response *DescribeTargetsListResponse) {
	response = &DescribeTargetsListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTargetsList(request *DescribeTargetsListRequest) string {
	return c.DescribeTargetsListWithContext(context.Background(), request)
}

func (c *Client) DescribeTargetsListWithContext(ctx context.Context, request *DescribeTargetsListRequest) string {
	if request == nil {
		request = NewDescribeTargetsListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTargetsListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAgentStatusRequest() (request *DescribeAgentStatusRequest) {
	request = &DescribeAgentStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeAgentStatus")
	return
}

func NewDescribeAgentStatusResponse() (response *DescribeAgentStatusResponse) {
	response = &DescribeAgentStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAgentStatus(request *DescribeAgentStatusRequest) string {
	return c.DescribeAgentStatusWithContext(context.Background(), request)
}

func (c *Client) DescribeAgentStatusWithContext(ctx context.Context, request *DescribeAgentStatusRequest) string {
	if request == nil {
		request = NewDescribeAgentStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAgentStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateMonitorCollectionConfigRequest() (request *CreateMonitorCollectionConfigRequest) {
	request = &CreateMonitorCollectionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateMonitorCollectionConfig")
	return
}

func NewCreateMonitorCollectionConfigResponse() (response *CreateMonitorCollectionConfigResponse) {
	response = &CreateMonitorCollectionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMonitorCollectionConfig(request *CreateMonitorCollectionConfigRequest) string {
	return c.CreateMonitorCollectionConfigWithContext(context.Background(), request)
}

func (c *Client) CreateMonitorCollectionConfigWithContext(ctx context.Context, request *CreateMonitorCollectionConfigRequest) string {
	if request == nil {
		request = NewCreateMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMonitorCollectionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteMonitorCollectionConfigRequest() (request *DeleteMonitorCollectionConfigRequest) {
	request = &DeleteMonitorCollectionConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteMonitorCollectionConfig")
	return
}

func NewDeleteMonitorCollectionConfigResponse() (response *DeleteMonitorCollectionConfigResponse) {
	response = &DeleteMonitorCollectionConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMonitorCollectionConfig(request *DeleteMonitorCollectionConfigRequest) string {
	return c.DeleteMonitorCollectionConfigWithContext(context.Background(), request)
}

func (c *Client) DeleteMonitorCollectionConfigWithContext(ctx context.Context, request *DeleteMonitorCollectionConfigRequest) string {
	if request == nil {
		request = NewDeleteMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMonitorCollectionConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewEnableMetricsRequest() (request *EnableMetricsRequest) {
	request = &EnableMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "EnableMetrics")
	return
}

func NewEnableMetricsResponse() (response *EnableMetricsResponse) {
	response = &EnableMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableMetrics(request *EnableMetricsRequest) string {
	return c.EnableMetricsWithContext(context.Background(), request)
}

func (c *Client) EnableMetricsWithContext(ctx context.Context, request *EnableMetricsRequest) string {
	if request == nil {
		request = NewEnableMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDropMetricsRequest() (request *DropMetricsRequest) {
	request = &DropMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DropMetrics")
	return
}

func NewDropMetricsResponse() (response *DropMetricsResponse) {
	response = &DropMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DropMetrics(request *DropMetricsRequest) string {
	return c.DropMetricsWithContext(context.Background(), request)
}

func (c *Client) DropMetricsWithContext(ctx context.Context, request *DropMetricsRequest) string {
	if request == nil {
		request = NewDropMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDropMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
