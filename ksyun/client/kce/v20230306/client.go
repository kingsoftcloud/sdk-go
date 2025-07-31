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

func (c *Client) CreatePrometheusInstanceSend(request *CreatePrometheusInstanceRequest) (*CreatePrometheusInstanceResponse, error) {
	statusCode, msg, err := c.CreatePrometheusInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreatePrometheusInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreatePrometheusInstanceWithContextV2(ctx context.Context, request *CreatePrometheusInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreatePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrometheusInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribePrometheusInstanceSend(request *DescribePrometheusInstanceRequest) (*DescribePrometheusInstanceResponse, error) {
	statusCode, msg, err := c.DescribePrometheusInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePrometheusInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribePrometheusInstanceWithContextV2(ctx context.Context, request *DescribePrometheusInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrometheusInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdatePrometheusInstanceSend(request *UpdatePrometheusInstanceRequest) (*UpdatePrometheusInstanceResponse, error) {
	statusCode, msg, err := c.UpdatePrometheusInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdatePrometheusInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdatePrometheusInstanceWithContextV2(ctx context.Context, request *UpdatePrometheusInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewUpdatePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdatePrometheusInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeletePrometheusInstanceSend(request *DeletePrometheusInstanceRequest) (*DeletePrometheusInstanceResponse, error) {
	statusCode, msg, err := c.DeletePrometheusInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeletePrometheusInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeletePrometheusInstanceWithContextV2(ctx context.Context, request *DeletePrometheusInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePrometheusInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrometheusInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) EnableGrafanaSend(request *EnableGrafanaRequest) (*EnableGrafanaResponse, error) {
	statusCode, msg, err := c.EnableGrafanaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct EnableGrafanaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) EnableGrafanaWithContextV2(ctx context.Context, request *EnableGrafanaRequest) (int, string, error) {
	if request == nil {
		request = NewEnableGrafanaRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableGrafanaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateGrafanaPasswordSend(request *UpdateGrafanaPasswordRequest) (*UpdateGrafanaPasswordResponse, error) {
	statusCode, msg, err := c.UpdateGrafanaPasswordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateGrafanaPasswordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateGrafanaPasswordWithContextV2(ctx context.Context, request *UpdateGrafanaPasswordRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateGrafanaPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateGrafanaPasswordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) EnableGrafanaInternetSend(request *EnableGrafanaInternetRequest) (*EnableGrafanaInternetResponse, error) {
	statusCode, msg, err := c.EnableGrafanaInternetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct EnableGrafanaInternetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) EnableGrafanaInternetWithContextV2(ctx context.Context, request *EnableGrafanaInternetRequest) (int, string, error) {
	if request == nil {
		request = NewEnableGrafanaInternetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableGrafanaInternetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeGrafanaWhiteListSend(request *DescribeGrafanaWhiteListRequest) (*DescribeGrafanaWhiteListResponse, error) {
	statusCode, msg, err := c.DescribeGrafanaWhiteListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeGrafanaWhiteListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeGrafanaWhiteListWithContextV2(ctx context.Context, request *DescribeGrafanaWhiteListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeGrafanaWhiteListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeGrafanaWhiteListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateGrafanaWhiteListSend(request *UpdateGrafanaWhiteListRequest) (*UpdateGrafanaWhiteListResponse, error) {
	statusCode, msg, err := c.UpdateGrafanaWhiteListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateGrafanaWhiteListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateGrafanaWhiteListWithContextV2(ctx context.Context, request *UpdateGrafanaWhiteListRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateGrafanaWhiteListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateGrafanaWhiteListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) AssociateClusterSend(request *AssociateClusterRequest) (*AssociateClusterResponse, error) {
	statusCode, msg, err := c.AssociateClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssociateClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) AssociateClusterWithContextV2(ctx context.Context, request *AssociateClusterRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DisassociateClusterSend(request *DisassociateClusterRequest) (*DisassociateClusterResponse, error) {
	statusCode, msg, err := c.DisassociateClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DisassociateClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DisassociateClusterWithContextV2(ctx context.Context, request *DisassociateClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeAssociateClusterListSend(request *DescribeAssociateClusterListRequest) (*DescribeAssociateClusterListResponse, error) {
	statusCode, msg, err := c.DescribeAssociateClusterListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAssociateClusterListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeAssociateClusterListWithContextV2(ctx context.Context, request *DescribeAssociateClusterListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAssociateClusterListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAssociateClusterListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeMonitorListSend(request *DescribeMonitorListRequest) (*DescribeMonitorListResponse, error) {
	statusCode, msg, err := c.DescribeMonitorListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMonitorListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeMonitorListWithContextV2(ctx context.Context, request *DescribeMonitorListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMonitorListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeMonitorCollectionConfigSend(request *DescribeMonitorCollectionConfigRequest) (*DescribeMonitorCollectionConfigResponse, error) {
	statusCode, msg, err := c.DescribeMonitorCollectionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMonitorCollectionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeMonitorCollectionConfigWithContextV2(ctx context.Context, request *DescribeMonitorCollectionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorCollectionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateMonitorCollectionConfigSend(request *UpdateMonitorCollectionConfigRequest) (*UpdateMonitorCollectionConfigResponse, error) {
	statusCode, msg, err := c.UpdateMonitorCollectionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateMonitorCollectionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateMonitorCollectionConfigWithContextV2(ctx context.Context, request *UpdateMonitorCollectionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateMonitorCollectionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeMonitorMetricsListSend(request *DescribeMonitorMetricsListRequest) (*DescribeMonitorMetricsListResponse, error) {
	statusCode, msg, err := c.DescribeMonitorMetricsListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMonitorMetricsListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeMonitorMetricsListWithContextV2(ctx context.Context, request *DescribeMonitorMetricsListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMonitorMetricsListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorMetricsListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeTargetsListSend(request *DescribeTargetsListRequest) (*DescribeTargetsListResponse, error) {
	statusCode, msg, err := c.DescribeTargetsListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeTargetsListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeTargetsListWithContextV2(ctx context.Context, request *DescribeTargetsListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTargetsListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTargetsListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeAgentStatusSend(request *DescribeAgentStatusRequest) (*DescribeAgentStatusResponse, error) {
	statusCode, msg, err := c.DescribeAgentStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAgentStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeAgentStatusWithContextV2(ctx context.Context, request *DescribeAgentStatusRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAgentStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAgentStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateMonitorCollectionConfigSend(request *CreateMonitorCollectionConfigRequest) (*CreateMonitorCollectionConfigResponse, error) {
	statusCode, msg, err := c.CreateMonitorCollectionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateMonitorCollectionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateMonitorCollectionConfigWithContextV2(ctx context.Context, request *CreateMonitorCollectionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateMonitorCollectionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteMonitorCollectionConfigSend(request *DeleteMonitorCollectionConfigRequest) (*DeleteMonitorCollectionConfigResponse, error) {
	statusCode, msg, err := c.DeleteMonitorCollectionConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteMonitorCollectionConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteMonitorCollectionConfigWithContextV2(ctx context.Context, request *DeleteMonitorCollectionConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMonitorCollectionConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMonitorCollectionConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) EnableMetricsSend(request *EnableMetricsRequest) (*EnableMetricsResponse, error) {
	statusCode, msg, err := c.EnableMetricsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct EnableMetricsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) EnableMetricsWithContextV2(ctx context.Context, request *EnableMetricsRequest) (int, string, error) {
	if request == nil {
		request = NewEnableMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEnableMetricsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DropMetricsSend(request *DropMetricsRequest) (*DropMetricsResponse, error) {
	statusCode, msg, err := c.DropMetricsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DropMetricsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DropMetricsWithContextV2(ctx context.Context, request *DropMetricsRequest) (int, string, error) {
	if request == nil {
		request = NewDropMetricsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDropMetricsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


