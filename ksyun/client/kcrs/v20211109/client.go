package v20211109

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2021-11-09"

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

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
	request = &CreateNamespaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateNamespace")
	return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
	response = &CreateNamespaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNamespace(request *CreateNamespaceRequest) string {
	return c.CreateNamespaceWithContext(context.Background(), request)
}

func (c *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest) string {
	if request == nil {
		request = NewCreateNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNamespaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNamespaceRequest() (request *DescribeNamespaceRequest) {
	request = &DescribeNamespaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeNamespace")
	return
}

func NewDescribeNamespaceResponse() (response *DescribeNamespaceResponse) {
	response = &DescribeNamespaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNamespace(request *DescribeNamespaceRequest) string {
	return c.DescribeNamespaceWithContext(context.Background(), request)
}

func (c *Client) DescribeNamespaceWithContext(ctx context.Context, request *DescribeNamespaceRequest) string {
	if request == nil {
		request = NewDescribeNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNamespaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyNamespaceTypeRequest() (request *ModifyNamespaceTypeRequest) {
	request = &ModifyNamespaceTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "ModifyNamespaceType")
	return
}

func NewModifyNamespaceTypeResponse() (response *ModifyNamespaceTypeResponse) {
	response = &ModifyNamespaceTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNamespaceType(request *ModifyNamespaceTypeRequest) string {
	return c.ModifyNamespaceTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyNamespaceTypeWithContext(ctx context.Context, request *ModifyNamespaceTypeRequest) string {
	if request == nil {
		request = NewModifyNamespaceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNamespaceTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNamespaceExistRequest() (request *DescribeNamespaceExistRequest) {
	request = &DescribeNamespaceExistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeNamespaceExist")
	return
}

func NewDescribeNamespaceExistResponse() (response *DescribeNamespaceExistResponse) {
	response = &DescribeNamespaceExistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNamespaceExist(request *DescribeNamespaceExistRequest) string {
	return c.DescribeNamespaceExistWithContext(context.Background(), request)
}

func (c *Client) DescribeNamespaceExistWithContext(ctx context.Context, request *DescribeNamespaceExistRequest) string {
	if request == nil {
		request = NewDescribeNamespaceExistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeNamespaceExistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
	request = &DeleteNamespaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteNamespace")
	return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
	response = &DeleteNamespaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) string {
	return c.DeleteNamespaceWithContext(context.Background(), request)
}

func (c *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest) string {
	if request == nil {
		request = NewDeleteNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNamespaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) string {
	return c.DescribeImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) string {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteImages")
	return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteImages(request *DeleteImagesRequest) string {
	return c.DeleteImagesWithContext(context.Background(), request)
}

func (c *Client) DeleteImagesWithContext(ctx context.Context, request *DeleteImagesRequest) string {
	if request == nil {
		request = NewDeleteImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteRepoTagRequest() (request *DeleteRepoTagRequest) {
	request = &DeleteRepoTagRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteRepoTag")
	return
}

func NewDeleteRepoTagResponse() (response *DeleteRepoTagResponse) {
	response = &DeleteRepoTagResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRepoTag(request *DeleteRepoTagRequest) string {
	return c.DeleteRepoTagWithContext(context.Background(), request)
}

func (c *Client) DeleteRepoTagWithContext(ctx context.Context, request *DeleteRepoTagRequest) string {
	if request == nil {
		request = NewDeleteRepoTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRepoTagResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRepositoryRequest() (request *DescribeRepositoryRequest) {
	request = &DescribeRepositoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeRepository")
	return
}

func NewDescribeRepositoryResponse() (response *DescribeRepositoryResponse) {
	response = &DescribeRepositoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRepository(request *DescribeRepositoryRequest) string {
	return c.DescribeRepositoryWithContext(context.Background(), request)
}

func (c *Client) DescribeRepositoryWithContext(ctx context.Context, request *DescribeRepositoryRequest) string {
	if request == nil {
		request = NewDescribeRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRepositoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyRepoDescRequest() (request *ModifyRepoDescRequest) {
	request = &ModifyRepoDescRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "ModifyRepoDesc")
	return
}

func NewModifyRepoDescResponse() (response *ModifyRepoDescResponse) {
	response = &ModifyRepoDescResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRepoDesc(request *ModifyRepoDescRequest) string {
	return c.ModifyRepoDescWithContext(context.Background(), request)
}

func (c *Client) ModifyRepoDescWithContext(ctx context.Context, request *ModifyRepoDescRequest) string {
	if request == nil {
		request = NewModifyRepoDescRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRepoDescResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
	request = &DeleteRepositoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteRepository")
	return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
	response = &DeleteRepositoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) string {
	return c.DeleteRepositoryWithContext(context.Background(), request)
}

func (c *Client) DeleteRepositoryWithContext(ctx context.Context, request *DeleteRepositoryRequest) string {
	if request == nil {
		request = NewDeleteRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRepositoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStartImageScanRequest() (request *StartImageScanRequest) {
	request = &StartImageScanRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "StartImageScan")
	return
}

func NewStartImageScanResponse() (response *StartImageScanResponse) {
	response = &StartImageScanResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartImageScan(request *StartImageScanRequest) string {
	return c.StartImageScanWithContext(context.Background(), request)
}

func (c *Client) StartImageScanWithContext(ctx context.Context, request *StartImageScanRequest) string {
	if request == nil {
		request = NewStartImageScanRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartImageScanResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeImageScanRequest() (request *DescribeImageScanRequest) {
	request = &DescribeImageScanRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeImageScan")
	return
}

func NewDescribeImageScanResponse() (response *DescribeImageScanResponse) {
	response = &DescribeImageScanResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImageScan(request *DescribeImageScanRequest) string {
	return c.DescribeImageScanWithContext(context.Background(), request)
}

func (c *Client) DescribeImageScanWithContext(ctx context.Context, request *DescribeImageScanRequest) string {
	if request == nil {
		request = NewDescribeImageScanRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageScanResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateInstanceTokenRequest() (request *CreateInstanceTokenRequest) {
	request = &CreateInstanceTokenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateInstanceToken")
	return
}

func NewCreateInstanceTokenResponse() (response *CreateInstanceTokenResponse) {
	response = &CreateInstanceTokenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstanceToken(request *CreateInstanceTokenRequest) string {
	return c.CreateInstanceTokenWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceTokenWithContext(ctx context.Context, request *CreateInstanceTokenRequest) string {
	if request == nil {
		request = NewCreateInstanceTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceTokenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInternalEndpointRequest() (request *DescribeInternalEndpointRequest) {
	request = &DescribeInternalEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeInternalEndpoint")
	return
}

func NewDescribeInternalEndpointResponse() (response *DescribeInternalEndpointResponse) {
	response = &DescribeInternalEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInternalEndpoint(request *DescribeInternalEndpointRequest) string {
	return c.DescribeInternalEndpointWithContext(context.Background(), request)
}

func (c *Client) DescribeInternalEndpointWithContext(ctx context.Context, request *DescribeInternalEndpointRequest) string {
	if request == nil {
		request = NewDescribeInternalEndpointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInternalEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceTokenRequest() (request *DescribeInstanceTokenRequest) {
	request = &DescribeInstanceTokenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeInstanceToken")
	return
}

func NewDescribeInstanceTokenResponse() (response *DescribeInstanceTokenResponse) {
	response = &DescribeInstanceTokenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceToken(request *DescribeInstanceTokenRequest) string {
	return c.DescribeInstanceTokenWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceTokenWithContext(ctx context.Context, request *DescribeInstanceTokenRequest) string {
	if request == nil {
		request = NewDescribeInstanceTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceTokenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateInternalEndpointRequest() (request *CreateInternalEndpointRequest) {
	request = &CreateInternalEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateInternalEndpoint")
	return
}

func NewCreateInternalEndpointResponse() (response *CreateInternalEndpointResponse) {
	response = &CreateInternalEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInternalEndpoint(request *CreateInternalEndpointRequest) string {
	return c.CreateInternalEndpointWithContext(context.Background(), request)
}

func (c *Client) CreateInternalEndpointWithContext(ctx context.Context, request *CreateInternalEndpointRequest) string {
	if request == nil {
		request = NewCreateInternalEndpointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInternalEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceTokenStatusRequest() (request *ModifyInstanceTokenStatusRequest) {
	request = &ModifyInstanceTokenStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "ModifyInstanceTokenStatus")
	return
}

func NewModifyInstanceTokenStatusResponse() (response *ModifyInstanceTokenStatusResponse) {
	response = &ModifyInstanceTokenStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceTokenStatus(request *ModifyInstanceTokenStatusRequest) string {
	return c.ModifyInstanceTokenStatusWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceTokenStatusWithContext(ctx context.Context, request *ModifyInstanceTokenStatusRequest) string {
	if request == nil {
		request = NewModifyInstanceTokenStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTokenStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInternalEndpointRequest() (request *DeleteInternalEndpointRequest) {
	request = &DeleteInternalEndpointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteInternalEndpoint")
	return
}

func NewDeleteInternalEndpointResponse() (response *DeleteInternalEndpointResponse) {
	response = &DeleteInternalEndpointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInternalEndpoint(request *DeleteInternalEndpointRequest) string {
	return c.DeleteInternalEndpointWithContext(context.Background(), request)
}

func (c *Client) DeleteInternalEndpointWithContext(ctx context.Context, request *DeleteInternalEndpointRequest) string {
	if request == nil {
		request = NewDeleteInternalEndpointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInternalEndpointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyInstanceTokenInformationRequest() (request *ModifyInstanceTokenInformationRequest) {
	request = &ModifyInstanceTokenInformationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "ModifyInstanceTokenInformation")
	return
}

func NewModifyInstanceTokenInformationResponse() (response *ModifyInstanceTokenInformationResponse) {
	response = &ModifyInstanceTokenInformationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceTokenInformation(request *ModifyInstanceTokenInformationRequest) string {
	return c.ModifyInstanceTokenInformationWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceTokenInformationWithContext(ctx context.Context, request *ModifyInstanceTokenInformationRequest) string {
	if request == nil {
		request = NewModifyInstanceTokenInformationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTokenInformationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInternalEndpointDnsRequest() (request *DescribeInternalEndpointDnsRequest) {
	request = &DescribeInternalEndpointDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeInternalEndpointDns")
	return
}

func NewDescribeInternalEndpointDnsResponse() (response *DescribeInternalEndpointDnsResponse) {
	response = &DescribeInternalEndpointDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInternalEndpointDns(request *DescribeInternalEndpointDnsRequest) string {
	return c.DescribeInternalEndpointDnsWithContext(context.Background(), request)
}

func (c *Client) DescribeInternalEndpointDnsWithContext(ctx context.Context, request *DescribeInternalEndpointDnsRequest) string {
	if request == nil {
		request = NewDescribeInternalEndpointDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInternalEndpointDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInstanceTokenRequest() (request *DeleteInstanceTokenRequest) {
	request = &DeleteInstanceTokenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteInstanceToken")
	return
}

func NewDeleteInstanceTokenResponse() (response *DeleteInstanceTokenResponse) {
	response = &DeleteInstanceTokenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstanceToken(request *DeleteInstanceTokenRequest) string {
	return c.DeleteInstanceTokenWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceTokenWithContext(ctx context.Context, request *DeleteInstanceTokenRequest) string {
	if request == nil {
		request = NewDeleteInstanceTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceTokenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateInternalEndpointDnsRequest() (request *CreateInternalEndpointDnsRequest) {
	request = &CreateInternalEndpointDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateInternalEndpointDns")
	return
}

func NewCreateInternalEndpointDnsResponse() (response *CreateInternalEndpointDnsResponse) {
	response = &CreateInternalEndpointDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInternalEndpointDns(request *CreateInternalEndpointDnsRequest) string {
	return c.CreateInternalEndpointDnsWithContext(context.Background(), request)
}

func (c *Client) CreateInternalEndpointDnsWithContext(ctx context.Context, request *CreateInternalEndpointDnsRequest) string {
	if request == nil {
		request = NewCreateInternalEndpointDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInternalEndpointDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInternalEndpointDnsRequest() (request *DeleteInternalEndpointDnsRequest) {
	request = &DeleteInternalEndpointDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteInternalEndpointDns")
	return
}

func NewDeleteInternalEndpointDnsResponse() (response *DeleteInternalEndpointDnsResponse) {
	response = &DeleteInternalEndpointDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInternalEndpointDns(request *DeleteInternalEndpointDnsRequest) string {
	return c.DeleteInternalEndpointDnsWithContext(context.Background(), request)
}

func (c *Client) DeleteInternalEndpointDnsWithContext(ctx context.Context, request *DeleteInternalEndpointDnsRequest) string {
	if request == nil {
		request = NewDeleteInternalEndpointDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInternalEndpointDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateInstance")
	return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstance(request *CreateInstanceRequest) string {
	return c.CreateInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) string {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstance(request *DeleteInstanceRequest) string {
	return c.DeleteInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) string {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceUsageRequest() (request *DescribeInstanceUsageRequest) {
	request = &DescribeInstanceUsageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeInstanceUsage")
	return
}

func NewDescribeInstanceUsageResponse() (response *DescribeInstanceUsageResponse) {
	response = &DescribeInstanceUsageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceUsage(request *DescribeInstanceUsageRequest) string {
	return c.DescribeInstanceUsageWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceUsageWithContext(ctx context.Context, request *DescribeInstanceUsageRequest) string {
	if request == nil {
		request = NewDescribeInstanceUsageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceUsageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeInstance")
	return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstance(request *DescribeInstanceRequest) string {
	return c.DescribeInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) string {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateWebhookTriggerRequest() (request *CreateWebhookTriggerRequest) {
	request = &CreateWebhookTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateWebhookTrigger")
	return
}

func NewCreateWebhookTriggerResponse() (response *CreateWebhookTriggerResponse) {
	response = &CreateWebhookTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateWebhookTrigger(request *CreateWebhookTriggerRequest) string {
	return c.CreateWebhookTriggerWithContext(context.Background(), request)
}

func (c *Client) CreateWebhookTriggerWithContext(ctx context.Context, request *CreateWebhookTriggerRequest) string {
	if request == nil {
		request = NewCreateWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateWebhookTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeWebhookTriggerRequest() (request *DescribeWebhookTriggerRequest) {
	request = &DescribeWebhookTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeWebhookTrigger")
	return
}

func NewDescribeWebhookTriggerResponse() (response *DescribeWebhookTriggerResponse) {
	response = &DescribeWebhookTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeWebhookTrigger(request *DescribeWebhookTriggerRequest) string {
	return c.DescribeWebhookTriggerWithContext(context.Background(), request)
}

func (c *Client) DescribeWebhookTriggerWithContext(ctx context.Context, request *DescribeWebhookTriggerRequest) string {
	if request == nil {
		request = NewDescribeWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeWebhookTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyWebhookTriggerRequest() (request *ModifyWebhookTriggerRequest) {
	request = &ModifyWebhookTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "ModifyWebhookTrigger")
	return
}

func NewModifyWebhookTriggerResponse() (response *ModifyWebhookTriggerResponse) {
	response = &ModifyWebhookTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyWebhookTrigger(request *ModifyWebhookTriggerRequest) string {
	return c.ModifyWebhookTriggerWithContext(context.Background(), request)
}

func (c *Client) ModifyWebhookTriggerWithContext(ctx context.Context, request *ModifyWebhookTriggerRequest) string {
	if request == nil {
		request = NewModifyWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyWebhookTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeWebhookTriggerLogRequest() (request *DescribeWebhookTriggerLogRequest) {
	request = &DescribeWebhookTriggerLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeWebhookTriggerLog")
	return
}

func NewDescribeWebhookTriggerLogResponse() (response *DescribeWebhookTriggerLogResponse) {
	response = &DescribeWebhookTriggerLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeWebhookTriggerLog(request *DescribeWebhookTriggerLogRequest) string {
	return c.DescribeWebhookTriggerLogWithContext(context.Background(), request)
}

func (c *Client) DescribeWebhookTriggerLogWithContext(ctx context.Context, request *DescribeWebhookTriggerLogRequest) string {
	if request == nil {
		request = NewDescribeWebhookTriggerLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeWebhookTriggerLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteWebhookTriggerRequest() (request *DeleteWebhookTriggerRequest) {
	request = &DeleteWebhookTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteWebhookTrigger")
	return
}

func NewDeleteWebhookTriggerResponse() (response *DeleteWebhookTriggerResponse) {
	response = &DeleteWebhookTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteWebhookTrigger(request *DeleteWebhookTriggerRequest) string {
	return c.DeleteWebhookTriggerWithContext(context.Background(), request)
}

func (c *Client) DeleteWebhookTriggerWithContext(ctx context.Context, request *DeleteWebhookTriggerRequest) string {
	if request == nil {
		request = NewDeleteWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteWebhookTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateRetentionRuleRequest() (request *CreateRetentionRuleRequest) {
	request = &CreateRetentionRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "CreateRetentionRule")
	return
}

func NewCreateRetentionRuleResponse() (response *CreateRetentionRuleResponse) {
	response = &CreateRetentionRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRetentionRule(request *CreateRetentionRuleRequest) string {
	return c.CreateRetentionRuleWithContext(context.Background(), request)
}

func (c *Client) CreateRetentionRuleWithContext(ctx context.Context, request *CreateRetentionRuleRequest) string {
	if request == nil {
		request = NewCreateRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRetentionRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateRetentionRuleRequest() (request *UpdateRetentionRuleRequest) {
	request = &UpdateRetentionRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "UpdateRetentionRule")
	return
}

func NewUpdateRetentionRuleResponse() (response *UpdateRetentionRuleResponse) {
	response = &UpdateRetentionRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateRetentionRule(request *UpdateRetentionRuleRequest) string {
	return c.UpdateRetentionRuleWithContext(context.Background(), request)
}

func (c *Client) UpdateRetentionRuleWithContext(ctx context.Context, request *UpdateRetentionRuleRequest) string {
	if request == nil {
		request = NewUpdateRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateRetentionRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteRetentionRuleRequest() (request *DeleteRetentionRuleRequest) {
	request = &DeleteRetentionRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DeleteRetentionRule")
	return
}

func NewDeleteRetentionRuleResponse() (response *DeleteRetentionRuleResponse) {
	response = &DeleteRetentionRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRetentionRule(request *DeleteRetentionRuleRequest) string {
	return c.DeleteRetentionRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteRetentionRuleWithContext(ctx context.Context, request *DeleteRetentionRuleRequest) string {
	if request == nil {
		request = NewDeleteRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRetentionRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRetentionRuleRequest() (request *DescribeRetentionRuleRequest) {
	request = &DescribeRetentionRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "DescribeRetentionRule")
	return
}

func NewDescribeRetentionRuleResponse() (response *DescribeRetentionRuleResponse) {
	response = &DescribeRetentionRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRetentionRule(request *DescribeRetentionRuleRequest) string {
	return c.DescribeRetentionRuleWithContext(context.Background(), request)
}

func (c *Client) DescribeRetentionRuleWithContext(ctx context.Context, request *DescribeRetentionRuleRequest) string {
	if request == nil {
		request = NewDescribeRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRetentionRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRunRetentionPolicyRequest() (request *RunRetentionPolicyRequest) {
	request = &RunRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "RunRetentionPolicy")
	return
}

func NewRunRetentionPolicyResponse() (response *RunRetentionPolicyResponse) {
	response = &RunRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RunRetentionPolicy(request *RunRetentionPolicyRequest) string {
	return c.RunRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) RunRetentionPolicyWithContext(ctx context.Context, request *RunRetentionPolicyRequest) string {
	if request == nil {
		request = NewRunRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRunRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetRetentionPolicyLogsRequest() (request *GetRetentionPolicyLogsRequest) {
	request = &GetRetentionPolicyLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "GetRetentionPolicyLogs")
	return
}

func NewGetRetentionPolicyLogsResponse() (response *GetRetentionPolicyLogsResponse) {
	response = &GetRetentionPolicyLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRetentionPolicyLogs(request *GetRetentionPolicyLogsRequest) string {
	return c.GetRetentionPolicyLogsWithContext(context.Background(), request)
}

func (c *Client) GetRetentionPolicyLogsWithContext(ctx context.Context, request *GetRetentionPolicyLogsRequest) string {
	if request == nil {
		request = NewGetRetentionPolicyLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionPolicyLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetRetentionPolicyLogDetailRequest() (request *GetRetentionPolicyLogDetailRequest) {
	request = &GetRetentionPolicyLogDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "GetRetentionPolicyLogDetail")
	return
}

func NewGetRetentionPolicyLogDetailResponse() (response *GetRetentionPolicyLogDetailResponse) {
	response = &GetRetentionPolicyLogDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRetentionPolicyLogDetail(request *GetRetentionPolicyLogDetailRequest) string {
	return c.GetRetentionPolicyLogDetailWithContext(context.Background(), request)
}

func (c *Client) GetRetentionPolicyLogDetailWithContext(ctx context.Context, request *GetRetentionPolicyLogDetailRequest) string {
	if request == nil {
		request = NewGetRetentionPolicyLogDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionPolicyLogDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetRetentionPolicyLogRequest() (request *GetRetentionPolicyLogRequest) {
	request = &GetRetentionPolicyLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "GetRetentionPolicyLog")
	return
}

func NewGetRetentionPolicyLogResponse() (response *GetRetentionPolicyLogResponse) {
	response = &GetRetentionPolicyLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRetentionPolicyLog(request *GetRetentionPolicyLogRequest) string {
	return c.GetRetentionPolicyLogWithContext(context.Background(), request)
}

func (c *Client) GetRetentionPolicyLogWithContext(ctx context.Context, request *GetRetentionPolicyLogRequest) string {
	if request == nil {
		request = NewGetRetentionPolicyLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionPolicyLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetRetentionTriggerRequest() (request *GetRetentionTriggerRequest) {
	request = &GetRetentionTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "GetRetentionTrigger")
	return
}

func NewGetRetentionTriggerResponse() (response *GetRetentionTriggerResponse) {
	response = &GetRetentionTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRetentionTrigger(request *GetRetentionTriggerRequest) string {
	return c.GetRetentionTriggerWithContext(context.Background(), request)
}

func (c *Client) GetRetentionTriggerWithContext(ctx context.Context, request *GetRetentionTriggerRequest) string {
	if request == nil {
		request = NewGetRetentionTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateRetentionTriggerRequest() (request *UpdateRetentionTriggerRequest) {
	request = &UpdateRetentionTriggerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "UpdateRetentionTrigger")
	return
}

func NewUpdateRetentionTriggerResponse() (response *UpdateRetentionTriggerResponse) {
	response = &UpdateRetentionTriggerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateRetentionTrigger(request *UpdateRetentionTriggerRequest) string {
	return c.UpdateRetentionTriggerWithContext(context.Background(), request)
}

func (c *Client) UpdateRetentionTriggerWithContext(ctx context.Context, request *UpdateRetentionTriggerRequest) string {
	if request == nil {
		request = NewUpdateRetentionTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateRetentionTriggerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewScheduleRequest() (request *ScheduleRequest) {
	request = &ScheduleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcrs", APIVersion, "Schedule")
	return
}

func NewScheduleResponse() (response *ScheduleResponse) {
	response = &ScheduleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Schedule(request *ScheduleRequest) string {
	return c.ScheduleWithContext(context.Background(), request)
}

func (c *Client) ScheduleWithContext(ctx context.Context, request *ScheduleRequest) string {
	if request == nil {
		request = NewScheduleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewScheduleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
