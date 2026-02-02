package v20240930

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-09-30"

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

func NewDescribeFileSystemListRequest() (request *DescribeFileSystemListRequest) {
	request = &DescribeFileSystemListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemList")
	return
}

func NewDescribeFileSystemListResponse() (response *DescribeFileSystemListResponse) {
	response = &DescribeFileSystemListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileSystemList(request *DescribeFileSystemListRequest) string {
	return c.DescribeFileSystemListWithContext(context.Background(), request)
}

func (c *Client) DescribeFileSystemListSend(request *DescribeFileSystemListRequest) (*DescribeFileSystemListResponse, error) {
	statusCode, msg, err := c.DescribeFileSystemListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFileSystemListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileSystemListWithContext(ctx context.Context, request *DescribeFileSystemListRequest) string {
	if request == nil {
		request = NewDescribeFileSystemListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileSystemListWithContextV2(ctx context.Context, request *DescribeFileSystemListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileSystemListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirQuotaListRequest() (request *DescribeDirQuotaListRequest) {
	request = &DescribeDirQuotaListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDirQuotaList")
	return
}

func NewDescribeDirQuotaListResponse() (response *DescribeDirQuotaListResponse) {
	response = &DescribeDirQuotaListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirQuotaList(request *DescribeDirQuotaListRequest) string {
	return c.DescribeDirQuotaListWithContext(context.Background(), request)
}

func (c *Client) DescribeDirQuotaListSend(request *DescribeDirQuotaListRequest) (*DescribeDirQuotaListResponse, error) {
	statusCode, msg, err := c.DescribeDirQuotaListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirQuotaListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirQuotaListWithContext(ctx context.Context, request *DescribeDirQuotaListRequest) string {
	if request == nil {
		request = NewDescribeDirQuotaListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDirQuotaList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirQuotaListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirQuotaListWithContextV2(ctx context.Context, request *DescribeDirQuotaListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirQuotaListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDirQuotaList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirQuotaListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDirQuotaRequest() (request *DeleteDirQuotaRequest) {
	request = &DeleteDirQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDirQuota")
	return
}

func NewDeleteDirQuotaResponse() (response *DeleteDirQuotaResponse) {
	response = &DeleteDirQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDirQuota(request *DeleteDirQuotaRequest) string {
	return c.DeleteDirQuotaWithContext(context.Background(), request)
}

func (c *Client) DeleteDirQuotaSend(request *DeleteDirQuotaRequest) (*DeleteDirQuotaResponse, error) {
	statusCode, msg, err := c.DeleteDirQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDirQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDirQuotaWithContext(ctx context.Context, request *DeleteDirQuotaRequest) string {
	if request == nil {
		request = NewDeleteDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDirQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDirQuotaWithContextV2(ctx context.Context, request *DeleteDirQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDirQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDirQuotaRequest() (request *UpdateDirQuotaRequest) {
	request = &UpdateDirQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDirQuota")
	return
}

func NewUpdateDirQuotaResponse() (response *UpdateDirQuotaResponse) {
	response = &UpdateDirQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDirQuota(request *UpdateDirQuotaRequest) string {
	return c.UpdateDirQuotaWithContext(context.Background(), request)
}

func (c *Client) UpdateDirQuotaSend(request *UpdateDirQuotaRequest) (*UpdateDirQuotaResponse, error) {
	statusCode, msg, err := c.UpdateDirQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDirQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDirQuotaWithContext(ctx context.Context, request *UpdateDirQuotaRequest) string {
	if request == nil {
		request = NewUpdateDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDirQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDirQuotaWithContextV2(ctx context.Context, request *UpdateDirQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDirQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDirQuotaRequest() (request *CreateDirQuotaRequest) {
	request = &CreateDirQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDirQuota")
	return
}

func NewCreateDirQuotaResponse() (response *CreateDirQuotaResponse) {
	response = &CreateDirQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDirQuota(request *CreateDirQuotaRequest) string {
	return c.CreateDirQuotaWithContext(context.Background(), request)
}

func (c *Client) CreateDirQuotaSend(request *CreateDirQuotaRequest) (*CreateDirQuotaResponse, error) {
	statusCode, msg, err := c.CreateDirQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDirQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDirQuotaWithContext(ctx context.Context, request *CreateDirQuotaRequest) string {
	if request == nil {
		request = NewCreateDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDirQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDirQuotaWithContextV2(ctx context.Context, request *CreateDirQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDirQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSubDirListRequest() (request *DescribeSubDirListRequest) {
	request = &DescribeSubDirListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeSubDirList")
	return
}

func NewDescribeSubDirListResponse() (response *DescribeSubDirListResponse) {
	response = &DescribeSubDirListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubDirList(request *DescribeSubDirListRequest) string {
	return c.DescribeSubDirListWithContext(context.Background(), request)
}

func (c *Client) DescribeSubDirListSend(request *DescribeSubDirListRequest) (*DescribeSubDirListResponse, error) {
	statusCode, msg, err := c.DescribeSubDirListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSubDirListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSubDirListWithContext(ctx context.Context, request *DescribeSubDirListRequest) string {
	if request == nil {
		request = NewDescribeSubDirListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeSubDirList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubDirListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSubDirListWithContextV2(ctx context.Context, request *DescribeSubDirListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSubDirListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeSubDirList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSubDirListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDirRequest() (request *DeleteDirRequest) {
	request = &DeleteDirRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDir")
	return
}

func NewDeleteDirResponse() (response *DeleteDirResponse) {
	response = &DeleteDirResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDir(request *DeleteDirRequest) string {
	return c.DeleteDirWithContext(context.Background(), request)
}

func (c *Client) DeleteDirSend(request *DeleteDirRequest) (*DeleteDirResponse, error) {
	statusCode, msg, err := c.DeleteDirWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDirResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDirWithContext(ctx context.Context, request *DeleteDirRequest) string {
	if request == nil {
		request = NewDeleteDirRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDir")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDirResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDirWithContextV2(ctx context.Context, request *DeleteDirRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDirRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDir")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDirResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDirRequest() (request *UpdateDirRequest) {
	request = &UpdateDirRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDir")
	return
}

func NewUpdateDirResponse() (response *UpdateDirResponse) {
	response = &UpdateDirResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDir(request *UpdateDirRequest) string {
	return c.UpdateDirWithContext(context.Background(), request)
}

func (c *Client) UpdateDirSend(request *UpdateDirRequest) (*UpdateDirResponse, error) {
	statusCode, msg, err := c.UpdateDirWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDirResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDirWithContext(ctx context.Context, request *UpdateDirRequest) string {
	if request == nil {
		request = NewUpdateDirRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDir")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDirResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDirWithContextV2(ctx context.Context, request *UpdateDirRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDirRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDir")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDirResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDirRequest() (request *CreateDirRequest) {
	request = &CreateDirRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDir")
	return
}

func NewCreateDirResponse() (response *CreateDirResponse) {
	response = &CreateDirResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDir(request *CreateDirRequest) string {
	return c.CreateDirWithContext(context.Background(), request)
}

func (c *Client) CreateDirSend(request *CreateDirRequest) (*CreateDirResponse, error) {
	statusCode, msg, err := c.CreateDirWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDirResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDirWithContext(ctx context.Context, request *CreateDirRequest) string {
	if request == nil {
		request = NewCreateDirRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDir")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDirResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDirWithContextV2(ctx context.Context, request *CreateDirRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDirRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDir")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDirResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDirQuotaRequest() (request *DescribeDirQuotaRequest) {
	request = &DescribeDirQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDirQuota")
	return
}

func NewDescribeDirQuotaResponse() (response *DescribeDirQuotaResponse) {
	response = &DescribeDirQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDirQuota(request *DescribeDirQuotaRequest) string {
	return c.DescribeDirQuotaWithContext(context.Background(), request)
}

func (c *Client) DescribeDirQuotaSend(request *DescribeDirQuotaRequest) (*DescribeDirQuotaResponse, error) {
	statusCode, msg, err := c.DescribeDirQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDirQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDirQuotaWithContext(ctx context.Context, request *DescribeDirQuotaRequest) string {
	if request == nil {
		request = NewDescribeDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDirQuotaWithContextV2(ctx context.Context, request *DescribeDirQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDirQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDirQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDirQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
