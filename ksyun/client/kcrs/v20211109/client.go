package v20211109

import (
	"context"
	"fmt"

	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) CreateNamespaceSend(request *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	statusCode, msg, err := c.CreateNamespaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateNamespaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateNamespaceWithContextV2(ctx context.Context, request *CreateNamespaceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateNamespaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNamespaceSend(request *DescribeNamespaceRequest) (*DescribeNamespaceResponse, error) {
	statusCode, msg, err := c.DescribeNamespaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNamespaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeNamespaceWithContextV2(ctx context.Context, request *DescribeNamespaceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNamespaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyNamespaceTypeSend(request *ModifyNamespaceTypeRequest) (*ModifyNamespaceTypeResponse, error) {
	statusCode, msg, err := c.ModifyNamespaceTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNamespaceTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyNamespaceTypeWithContextV2(ctx context.Context, request *ModifyNamespaceTypeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNamespaceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNamespaceTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNamespaceExistSend(request *DescribeNamespaceExistRequest) (*DescribeNamespaceExistResponse, error) {
	statusCode, msg, err := c.DescribeNamespaceExistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNamespaceExistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeNamespaceExistWithContextV2(ctx context.Context, request *DescribeNamespaceExistRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNamespaceExistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeNamespaceExistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteNamespaceSend(request *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	statusCode, msg, err := c.DeleteNamespaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteNamespaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteNamespaceWithContextV2(ctx context.Context, request *DeleteNamespaceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNamespaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeImagesSend(request *DescribeImagesRequest) (*DescribeImagesResponse, error) {
	statusCode, msg, err := c.DescribeImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeImagesWithContextV2(ctx context.Context, request *DescribeImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteImagesSend(request *DeleteImagesRequest) (*DeleteImagesResponse, error) {
	statusCode, msg, err := c.DeleteImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteImagesWithContextV2(ctx context.Context, request *DeleteImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteRepoTagSend(request *DeleteRepoTagRequest) (*DeleteRepoTagResponse, error) {
	statusCode, msg, err := c.DeleteRepoTagWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRepoTagResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteRepoTagWithContextV2(ctx context.Context, request *DeleteRepoTagRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRepoTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRepoTagResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeRepositorySend(request *DescribeRepositoryRequest) (*DescribeRepositoryResponse, error) {
	statusCode, msg, err := c.DescribeRepositoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRepositoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeRepositoryWithContextV2(ctx context.Context, request *DescribeRepositoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRepositoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyRepoDescSend(request *ModifyRepoDescRequest) (*ModifyRepoDescResponse, error) {
	statusCode, msg, err := c.ModifyRepoDescWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyRepoDescResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyRepoDescWithContextV2(ctx context.Context, request *ModifyRepoDescRequest) (int, string, error) {
	if request == nil {
		request = NewModifyRepoDescRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRepoDescResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteRepositorySend(request *DeleteRepositoryRequest) (*DeleteRepositoryResponse, error) {
	statusCode, msg, err := c.DeleteRepositoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRepositoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteRepositoryWithContextV2(ctx context.Context, request *DeleteRepositoryRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRepositoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StartImageScanSend(request *StartImageScanRequest) (*StartImageScanResponse, error) {
	statusCode, msg, err := c.StartImageScanWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartImageScanResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StartImageScanWithContextV2(ctx context.Context, request *StartImageScanRequest) (int, string, error) {
	if request == nil {
		request = NewStartImageScanRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartImageScanResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeImageScanSend(request *DescribeImageScanRequest) (*DescribeImageScanResponse, error) {
	statusCode, msg, err := c.DescribeImageScanWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImageScanResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeImageScanWithContextV2(ctx context.Context, request *DescribeImageScanRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImageScanRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImageScanResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateInstanceTokenSend(request *CreateInstanceTokenRequest) (*CreateInstanceTokenResponse, error) {
	statusCode, msg, err := c.CreateInstanceTokenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateInstanceTokenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateInstanceTokenWithContextV2(ctx context.Context, request *CreateInstanceTokenRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceTokenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeInternalEndpointSend(request *DescribeInternalEndpointRequest) (*DescribeInternalEndpointResponse, error) {
	statusCode, msg, err := c.DescribeInternalEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInternalEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInternalEndpointWithContextV2(ctx context.Context, request *DescribeInternalEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInternalEndpointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInternalEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeInstanceTokenSend(request *DescribeInstanceTokenRequest) (*DescribeInstanceTokenResponse, error) {
	statusCode, msg, err := c.DescribeInstanceTokenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceTokenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInstanceTokenWithContextV2(ctx context.Context, request *DescribeInstanceTokenRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceTokenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateInternalEndpointSend(request *CreateInternalEndpointRequest) (*CreateInternalEndpointResponse, error) {
	statusCode, msg, err := c.CreateInternalEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateInternalEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateInternalEndpointWithContextV2(ctx context.Context, request *CreateInternalEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInternalEndpointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInternalEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyInstanceTokenStatusSend(request *ModifyInstanceTokenStatusRequest) (*ModifyInstanceTokenStatusResponse, error) {
	statusCode, msg, err := c.ModifyInstanceTokenStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyInstanceTokenStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyInstanceTokenStatusWithContextV2(ctx context.Context, request *ModifyInstanceTokenStatusRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceTokenStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTokenStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteInternalEndpointSend(request *DeleteInternalEndpointRequest) (*DeleteInternalEndpointResponse, error) {
	statusCode, msg, err := c.DeleteInternalEndpointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteInternalEndpointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteInternalEndpointWithContextV2(ctx context.Context, request *DeleteInternalEndpointRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInternalEndpointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInternalEndpointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyInstanceTokenInformationSend(request *ModifyInstanceTokenInformationRequest) (*ModifyInstanceTokenInformationResponse, error) {
	statusCode, msg, err := c.ModifyInstanceTokenInformationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyInstanceTokenInformationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyInstanceTokenInformationWithContextV2(ctx context.Context, request *ModifyInstanceTokenInformationRequest) (int, string, error) {
	if request == nil {
		request = NewModifyInstanceTokenInformationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyInstanceTokenInformationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeInternalEndpointDnsSend(request *DescribeInternalEndpointDnsRequest) (*DescribeInternalEndpointDnsResponse, error) {
	statusCode, msg, err := c.DescribeInternalEndpointDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInternalEndpointDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInternalEndpointDnsWithContextV2(ctx context.Context, request *DescribeInternalEndpointDnsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInternalEndpointDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInternalEndpointDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteInstanceTokenSend(request *DeleteInstanceTokenRequest) (*DeleteInstanceTokenResponse, error) {
	statusCode, msg, err := c.DeleteInstanceTokenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteInstanceTokenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteInstanceTokenWithContextV2(ctx context.Context, request *DeleteInstanceTokenRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceTokenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateInternalEndpointDnsSend(request *CreateInternalEndpointDnsRequest) (*CreateInternalEndpointDnsResponse, error) {
	statusCode, msg, err := c.CreateInternalEndpointDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateInternalEndpointDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateInternalEndpointDnsWithContextV2(ctx context.Context, request *CreateInternalEndpointDnsRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInternalEndpointDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInternalEndpointDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteInternalEndpointDnsSend(request *DeleteInternalEndpointDnsRequest) (*DeleteInternalEndpointDnsResponse, error) {
	statusCode, msg, err := c.DeleteInternalEndpointDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteInternalEndpointDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteInternalEndpointDnsWithContextV2(ctx context.Context, request *DeleteInternalEndpointDnsRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInternalEndpointDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInternalEndpointDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateInstanceSend(request *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	statusCode, msg, err := c.CreateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateInstanceWithContextV2(ctx context.Context, request *CreateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteInstanceSend(request *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	statusCode, msg, err := c.DeleteInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) string {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceWithContextV2(ctx context.Context, request *DeleteInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeInstanceUsageSend(request *DescribeInstanceUsageRequest) (*DescribeInstanceUsageResponse, error) {
	statusCode, msg, err := c.DescribeInstanceUsageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceUsageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInstanceUsageWithContextV2(ctx context.Context, request *DescribeInstanceUsageRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceUsageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceUsageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeInstanceSend(request *DescribeInstanceRequest) (*DescribeInstanceResponse, error) {
	statusCode, msg, err := c.DescribeInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeInstanceWithContextV2(ctx context.Context, request *DescribeInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateWebhookTriggerSend(request *CreateWebhookTriggerRequest) (*CreateWebhookTriggerResponse, error) {
	statusCode, msg, err := c.CreateWebhookTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateWebhookTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateWebhookTriggerWithContextV2(ctx context.Context, request *CreateWebhookTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewCreateWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateWebhookTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeWebhookTriggerSend(request *DescribeWebhookTriggerRequest) (*DescribeWebhookTriggerResponse, error) {
	statusCode, msg, err := c.DescribeWebhookTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeWebhookTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeWebhookTriggerWithContextV2(ctx context.Context, request *DescribeWebhookTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeWebhookTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyWebhookTriggerSend(request *ModifyWebhookTriggerRequest) (*ModifyWebhookTriggerResponse, error) {
	statusCode, msg, err := c.ModifyWebhookTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyWebhookTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyWebhookTriggerWithContextV2(ctx context.Context, request *ModifyWebhookTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyWebhookTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeWebhookTriggerLogSend(request *DescribeWebhookTriggerLogRequest) (*DescribeWebhookTriggerLogResponse, error) {
	statusCode, msg, err := c.DescribeWebhookTriggerLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeWebhookTriggerLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeWebhookTriggerLogWithContextV2(ctx context.Context, request *DescribeWebhookTriggerLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeWebhookTriggerLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeWebhookTriggerLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteWebhookTriggerSend(request *DeleteWebhookTriggerRequest) (*DeleteWebhookTriggerResponse, error) {
	statusCode, msg, err := c.DeleteWebhookTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteWebhookTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteWebhookTriggerWithContextV2(ctx context.Context, request *DeleteWebhookTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteWebhookTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteWebhookTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateRetentionRuleSend(request *CreateRetentionRuleRequest) (*CreateRetentionRuleResponse, error) {
	statusCode, msg, err := c.CreateRetentionRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRetentionRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateRetentionRuleWithContextV2(ctx context.Context, request *CreateRetentionRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRetentionRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateRetentionRuleSend(request *UpdateRetentionRuleRequest) (*UpdateRetentionRuleResponse, error) {
	statusCode, msg, err := c.UpdateRetentionRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateRetentionRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateRetentionRuleWithContextV2(ctx context.Context, request *UpdateRetentionRuleRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateRetentionRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteRetentionRuleSend(request *DeleteRetentionRuleRequest) (*DeleteRetentionRuleResponse, error) {
	statusCode, msg, err := c.DeleteRetentionRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRetentionRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteRetentionRuleWithContextV2(ctx context.Context, request *DeleteRetentionRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRetentionRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeRetentionRuleSend(request *DescribeRetentionRuleRequest) (*DescribeRetentionRuleResponse, error) {
	statusCode, msg, err := c.DescribeRetentionRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRetentionRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeRetentionRuleWithContextV2(ctx context.Context, request *DescribeRetentionRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRetentionRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRetentionRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) RunRetentionPolicySend(request *RunRetentionPolicyRequest) (*RunRetentionPolicyResponse, error) {
	statusCode, msg, err := c.RunRetentionPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RunRetentionPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) RunRetentionPolicyWithContextV2(ctx context.Context, request *RunRetentionPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewRunRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRunRetentionPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetRetentionPolicyLogsSend(request *GetRetentionPolicyLogsRequest) (*GetRetentionPolicyLogsResponse, error) {
	statusCode, msg, err := c.GetRetentionPolicyLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRetentionPolicyLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetRetentionPolicyLogsWithContextV2(ctx context.Context, request *GetRetentionPolicyLogsRequest) (int, string, error) {
	if request == nil {
		request = NewGetRetentionPolicyLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionPolicyLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetRetentionPolicyLogDetailSend(request *GetRetentionPolicyLogDetailRequest) (*GetRetentionPolicyLogDetailResponse, error) {
	statusCode, msg, err := c.GetRetentionPolicyLogDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRetentionPolicyLogDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetRetentionPolicyLogDetailWithContextV2(ctx context.Context, request *GetRetentionPolicyLogDetailRequest) (int, string, error) {
	if request == nil {
		request = NewGetRetentionPolicyLogDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionPolicyLogDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetRetentionPolicyLogSend(request *GetRetentionPolicyLogRequest) (*GetRetentionPolicyLogResponse, error) {
	statusCode, msg, err := c.GetRetentionPolicyLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRetentionPolicyLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetRetentionPolicyLogWithContextV2(ctx context.Context, request *GetRetentionPolicyLogRequest) (int, string, error) {
	if request == nil {
		request = NewGetRetentionPolicyLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionPolicyLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetRetentionTriggerSend(request *GetRetentionTriggerRequest) (*GetRetentionTriggerResponse, error) {
	statusCode, msg, err := c.GetRetentionTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetRetentionTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetRetentionTriggerWithContextV2(ctx context.Context, request *GetRetentionTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewGetRetentionTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRetentionTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateRetentionTriggerSend(request *UpdateRetentionTriggerRequest) (*UpdateRetentionTriggerResponse, error) {
	statusCode, msg, err := c.UpdateRetentionTriggerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateRetentionTriggerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateRetentionTriggerWithContextV2(ctx context.Context, request *UpdateRetentionTriggerRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateRetentionTriggerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateRetentionTriggerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ScheduleSend(request *ScheduleRequest) (*ScheduleResponse, error) {
	statusCode, msg, err := c.ScheduleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ScheduleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ScheduleWithContextV2(ctx context.Context, request *ScheduleRequest) (int, string, error) {
	if request == nil {
		request = NewScheduleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewScheduleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
