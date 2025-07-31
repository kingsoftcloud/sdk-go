package v20180314
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-03-14"

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

func NewCreateRepoNamespaceRequest() (request *CreateRepoNamespaceRequest) {
	request = &CreateRepoNamespaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateRepoNamespace")
	return
}

func NewCreateRepoNamespaceResponse() (response *CreateRepoNamespaceResponse) {
	response = &CreateRepoNamespaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRepoNamespace(request *CreateRepoNamespaceRequest) string {
	return c.CreateRepoNamespaceWithContext(context.Background(), request)
}

func (c *Client) CreateRepoNamespaceSend(request *CreateRepoNamespaceRequest) (*CreateRepoNamespaceResponse, error) {
	statusCode, msg, err := c.CreateRepoNamespaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRepoNamespaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRepoNamespaceWithContext(ctx context.Context, request *CreateRepoNamespaceRequest) string {
	if request == nil {
		request = NewCreateRepoNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRepoNamespaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRepoNamespaceWithContextV2(ctx context.Context, request *CreateRepoNamespaceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRepoNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRepoNamespaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRepoNamespaceRequest() (request *DescribeRepoNamespaceRequest) {
	request = &DescribeRepoNamespaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeRepoNamespace")
	return
}

func NewDescribeRepoNamespaceResponse() (response *DescribeRepoNamespaceResponse) {
	response = &DescribeRepoNamespaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRepoNamespace(request *DescribeRepoNamespaceRequest) string {
	return c.DescribeRepoNamespaceWithContext(context.Background(), request)
}

func (c *Client) DescribeRepoNamespaceSend(request *DescribeRepoNamespaceRequest) (*DescribeRepoNamespaceResponse, error) {
	statusCode, msg, err := c.DescribeRepoNamespaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRepoNamespaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRepoNamespaceWithContext(ctx context.Context, request *DescribeRepoNamespaceRequest) string {
	if request == nil {
		request = NewDescribeRepoNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRepoNamespaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRepoNamespaceWithContextV2(ctx context.Context, request *DescribeRepoNamespaceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRepoNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRepoNamespaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyRepoNamespaceTypeRequest() (request *ModifyRepoNamespaceTypeRequest) {
	request = &ModifyRepoNamespaceTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyRepoNamespaceType")
	return
}

func NewModifyRepoNamespaceTypeResponse() (response *ModifyRepoNamespaceTypeResponse) {
	response = &ModifyRepoNamespaceTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRepoNamespaceType(request *ModifyRepoNamespaceTypeRequest) string {
	return c.ModifyRepoNamespaceTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyRepoNamespaceTypeSend(request *ModifyRepoNamespaceTypeRequest) (*ModifyRepoNamespaceTypeResponse, error) {
	statusCode, msg, err := c.ModifyRepoNamespaceTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyRepoNamespaceTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyRepoNamespaceTypeWithContext(ctx context.Context, request *ModifyRepoNamespaceTypeRequest) string {
	if request == nil {
		request = NewModifyRepoNamespaceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRepoNamespaceTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyRepoNamespaceTypeWithContextV2(ctx context.Context, request *ModifyRepoNamespaceTypeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyRepoNamespaceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRepoNamespaceTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRepoNamespaceExistRequest() (request *RepoNamespaceExistRequest) {
	request = &RepoNamespaceExistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "RepoNamespaceExist")
	return
}

func NewRepoNamespaceExistResponse() (response *RepoNamespaceExistResponse) {
	response = &RepoNamespaceExistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RepoNamespaceExist(request *RepoNamespaceExistRequest) string {
	return c.RepoNamespaceExistWithContext(context.Background(), request)
}

func (c *Client) RepoNamespaceExistSend(request *RepoNamespaceExistRequest) (*RepoNamespaceExistResponse, error) {
	statusCode, msg, err := c.RepoNamespaceExistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RepoNamespaceExistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RepoNamespaceExistWithContext(ctx context.Context, request *RepoNamespaceExistRequest) string {
	if request == nil {
		request = NewRepoNamespaceExistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRepoNamespaceExistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RepoNamespaceExistWithContextV2(ctx context.Context, request *RepoNamespaceExistRequest) (int, string, error) {
	if request == nil {
		request = NewRepoNamespaceExistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRepoNamespaceExistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
	request = &CreateRepositoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateRepository")
	return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
	response = &CreateRepositoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRepository(request *CreateRepositoryRequest) string {
	return c.CreateRepositoryWithContext(context.Background(), request)
}

func (c *Client) CreateRepositorySend(request *CreateRepositoryRequest) (*CreateRepositoryResponse, error) {
	statusCode, msg, err := c.CreateRepositoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRepositoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest) string {
	if request == nil {
		request = NewCreateRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRepositoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRepositoryWithContextV2(ctx context.Context, request *CreateRepositoryRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRepositoryResponse()
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
	request.Init().WithApiInfo("kce", APIVersion, "DeleteRepository")
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
func NewDescribeRepositoryRequest() (request *DescribeRepositoryRequest) {
	request = &DescribeRepositoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeRepository")
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
func NewDescribePublicRepositoryRequest() (request *DescribePublicRepositoryRequest) {
	request = &DescribePublicRepositoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribePublicRepository")
	return
}

func NewDescribePublicRepositoryResponse() (response *DescribePublicRepositoryResponse) {
	response = &DescribePublicRepositoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePublicRepository(request *DescribePublicRepositoryRequest) string {
	return c.DescribePublicRepositoryWithContext(context.Background(), request)
}

func (c *Client) DescribePublicRepositorySend(request *DescribePublicRepositoryRequest) (*DescribePublicRepositoryResponse, error) {
	statusCode, msg, err := c.DescribePublicRepositoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePublicRepositoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePublicRepositoryWithContext(ctx context.Context, request *DescribePublicRepositoryRequest) string {
	if request == nil {
		request = NewDescribePublicRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePublicRepositoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePublicRepositoryWithContextV2(ctx context.Context, request *DescribePublicRepositoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePublicRepositoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePublicRepositoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateRepoDescRequest() (request *UpdateRepoDescRequest) {
	request = &UpdateRepoDescRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdateRepoDesc")
	return
}

func NewUpdateRepoDescResponse() (response *UpdateRepoDescResponse) {
	response = &UpdateRepoDescResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateRepoDesc(request *UpdateRepoDescRequest) string {
	return c.UpdateRepoDescWithContext(context.Background(), request)
}

func (c *Client) UpdateRepoDescSend(request *UpdateRepoDescRequest) (*UpdateRepoDescResponse, error) {
	statusCode, msg, err := c.UpdateRepoDescWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateRepoDescResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateRepoDescWithContext(ctx context.Context, request *UpdateRepoDescRequest) string {
	if request == nil {
		request = NewUpdateRepoDescRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateRepoDescResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateRepoDescWithContextV2(ctx context.Context, request *UpdateRepoDescRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateRepoDescRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateRepoDescResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTagRequest() (request *DescribeTagRequest) {
	request = &DescribeTagRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeTag")
	return
}

func NewDescribeTagResponse() (response *DescribeTagResponse) {
	response = &DescribeTagResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTag(request *DescribeTagRequest) string {
	return c.DescribeTagWithContext(context.Background(), request)
}

func (c *Client) DescribeTagSend(request *DescribeTagRequest) (*DescribeTagResponse, error) {
	statusCode, msg, err := c.DescribeTagWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeTagResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTagWithContext(ctx context.Context, request *DescribeTagRequest) string {
	if request == nil {
		request = NewDescribeTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTagResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTagWithContextV2(ctx context.Context, request *DescribeTagRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTagResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
	request = &DeleteTagsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteTags")
	return
}

func NewDeleteTagsResponse() (response *DeleteTagsResponse) {
	response = &DeleteTagsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTags(request *DeleteTagsRequest) string {
	return c.DeleteTagsWithContext(context.Background(), request)
}

func (c *Client) DeleteTagsSend(request *DeleteTagsRequest) (*DeleteTagsResponse, error) {
	statusCode, msg, err := c.DeleteTagsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteTagsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteTagsWithContext(ctx context.Context, request *DeleteTagsRequest) string {
	if request == nil {
		request = NewDeleteTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTagsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteTagsWithContextV2(ctx context.Context, request *DeleteTagsRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteTagsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddFavorRequest() (request *AddFavorRequest) {
	request = &AddFavorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "AddFavor")
	return
}

func NewAddFavorResponse() (response *AddFavorResponse) {
	response = &AddFavorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddFavor(request *AddFavorRequest) string {
	return c.AddFavorWithContext(context.Background(), request)
}

func (c *Client) AddFavorSend(request *AddFavorRequest) (*AddFavorResponse, error) {
	statusCode, msg, err := c.AddFavorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddFavorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddFavorWithContext(ctx context.Context, request *AddFavorRequest) string {
	if request == nil {
		request = NewAddFavorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddFavorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddFavorWithContextV2(ctx context.Context, request *AddFavorRequest) (int, string, error) {
	if request == nil {
		request = NewAddFavorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddFavorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteFavorRequest() (request *DeleteFavorRequest) {
	request = &DeleteFavorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteFavor")
	return
}

func NewDeleteFavorResponse() (response *DeleteFavorResponse) {
	response = &DeleteFavorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFavor(request *DeleteFavorRequest) string {
	return c.DeleteFavorWithContext(context.Background(), request)
}

func (c *Client) DeleteFavorSend(request *DeleteFavorRequest) (*DeleteFavorResponse, error) {
	statusCode, msg, err := c.DeleteFavorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteFavorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFavorWithContext(ctx context.Context, request *DeleteFavorRequest) string {
	if request == nil {
		request = NewDeleteFavorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFavorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFavorWithContextV2(ctx context.Context, request *DeleteFavorRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFavorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFavorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetFavorRequest() (request *GetFavorRequest) {
	request = &GetFavorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "GetFavor")
	return
}

func NewGetFavorResponse() (response *GetFavorResponse) {
	response = &GetFavorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetFavor(request *GetFavorRequest) string {
	return c.GetFavorWithContext(context.Background(), request)
}

func (c *Client) GetFavorSend(request *GetFavorRequest) (*GetFavorResponse, error) {
	statusCode, msg, err := c.GetFavorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetFavorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetFavorWithContext(ctx context.Context, request *GetFavorRequest) string {
	if request == nil {
		request = NewGetFavorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetFavorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetFavorWithContextV2(ctx context.Context, request *GetFavorRequest) (int, string, error) {
	if request == nil {
		request = NewGetFavorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetFavorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRegisterRepositoryAccountRequest() (request *RegisterRepositoryAccountRequest) {
	request = &RegisterRepositoryAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "RegisterRepositoryAccount")
	return
}

func NewRegisterRepositoryAccountResponse() (response *RegisterRepositoryAccountResponse) {
	response = &RegisterRepositoryAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RegisterRepositoryAccount(request *RegisterRepositoryAccountRequest) string {
	return c.RegisterRepositoryAccountWithContext(context.Background(), request)
}

func (c *Client) RegisterRepositoryAccountSend(request *RegisterRepositoryAccountRequest) (*RegisterRepositoryAccountResponse, error) {
	statusCode, msg, err := c.RegisterRepositoryAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RegisterRepositoryAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RegisterRepositoryAccountWithContext(ctx context.Context, request *RegisterRepositoryAccountRequest) string {
	if request == nil {
		request = NewRegisterRepositoryAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterRepositoryAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RegisterRepositoryAccountWithContextV2(ctx context.Context, request *RegisterRepositoryAccountRequest) (int, string, error) {
	if request == nil {
		request = NewRegisterRepositoryAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRegisterRepositoryAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
	request = &ModifyPasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "ModifyPassword")
	return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
	response = &ModifyPasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPassword(request *ModifyPasswordRequest) string {
	return c.ModifyPasswordWithContext(context.Background(), request)
}

func (c *Client) ModifyPasswordSend(request *ModifyPasswordRequest) (*ModifyPasswordResponse, error) {
	statusCode, msg, err := c.ModifyPasswordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyPasswordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyPasswordWithContext(ctx context.Context, request *ModifyPasswordRequest) string {
	if request == nil {
		request = NewModifyPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyPasswordWithContextV2(ctx context.Context, request *ModifyPasswordRequest) (int, string, error) {
	if request == nil {
		request = NewModifyPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPasswordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRepoNamespaceRequest() (request *DeleteRepoNamespaceRequest) {
	request = &DeleteRepoNamespaceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteRepoNamespace")
	return
}

func NewDeleteRepoNamespaceResponse() (response *DeleteRepoNamespaceResponse) {
	response = &DeleteRepoNamespaceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRepoNamespace(request *DeleteRepoNamespaceRequest) string {
	return c.DeleteRepoNamespaceWithContext(context.Background(), request)
}

func (c *Client) DeleteRepoNamespaceSend(request *DeleteRepoNamespaceRequest) (*DeleteRepoNamespaceResponse, error) {
	statusCode, msg, err := c.DeleteRepoNamespaceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRepoNamespaceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRepoNamespaceWithContext(ctx context.Context, request *DeleteRepoNamespaceRequest) string {
	if request == nil {
		request = NewDeleteRepoNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRepoNamespaceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRepoNamespaceWithContextV2(ctx context.Context, request *DeleteRepoNamespaceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRepoNamespaceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRepoNamespaceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


