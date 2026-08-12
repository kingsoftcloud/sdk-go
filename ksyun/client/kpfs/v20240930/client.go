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
func NewGetTotalSizeRequest() (request *GetTotalSizeRequest) {
	request = &GetTotalSizeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetTotalSize")
	return
}

func NewGetTotalSizeResponse() (response *GetTotalSizeResponse) {
	response = &GetTotalSizeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetTotalSize(request *GetTotalSizeRequest) string {
	return c.GetTotalSizeWithContext(context.Background(), request)
}

func (c *Client) GetTotalSizeSend(request *GetTotalSizeRequest) (*GetTotalSizeResponse, error) {
	statusCode, msg, err := c.GetTotalSizeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetTotalSizeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetTotalSizeWithContext(ctx context.Context, request *GetTotalSizeRequest) string {
	if request == nil {
		request = NewGetTotalSizeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetTotalSize")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTotalSizeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetTotalSizeWithContextV2(ctx context.Context, request *GetTotalSizeRequest) (int, string, error) {
	if request == nil {
		request = NewGetTotalSizeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetTotalSize")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetTotalSizeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetInodeCountRequest() (request *GetInodeCountRequest) {
	request = &GetInodeCountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetInodeCount")
	return
}

func NewGetInodeCountResponse() (response *GetInodeCountResponse) {
	response = &GetInodeCountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetInodeCount(request *GetInodeCountRequest) string {
	return c.GetInodeCountWithContext(context.Background(), request)
}

func (c *Client) GetInodeCountSend(request *GetInodeCountRequest) (*GetInodeCountResponse, error) {
	statusCode, msg, err := c.GetInodeCountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetInodeCountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetInodeCountWithContext(ctx context.Context, request *GetInodeCountRequest) string {
	if request == nil {
		request = NewGetInodeCountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetInodeCount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetInodeCountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetInodeCountWithContextV2(ctx context.Context, request *GetInodeCountRequest) (int, string, error) {
	if request == nil {
		request = NewGetInodeCountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetInodeCount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetInodeCountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFileSystemClientInfoRequest() (request *DescribeFileSystemClientInfoRequest) {
	request = &DescribeFileSystemClientInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemClientInfo")
	return
}

func NewDescribeFileSystemClientInfoResponse() (response *DescribeFileSystemClientInfoResponse) {
	response = &DescribeFileSystemClientInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileSystemClientInfo(request *DescribeFileSystemClientInfoRequest) string {
	return c.DescribeFileSystemClientInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeFileSystemClientInfoSend(request *DescribeFileSystemClientInfoRequest) (*DescribeFileSystemClientInfoResponse, error) {
	statusCode, msg, err := c.DescribeFileSystemClientInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFileSystemClientInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileSystemClientInfoWithContext(ctx context.Context, request *DescribeFileSystemClientInfoRequest) string {
	if request == nil {
		request = NewDescribeFileSystemClientInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemClientInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemClientInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileSystemClientInfoWithContextV2(ctx context.Context, request *DescribeFileSystemClientInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileSystemClientInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemClientInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemClientInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFileSystemFileListRequest() (request *DescribeFileSystemFileListRequest) {
	request = &DescribeFileSystemFileListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemFileList")
	return
}

func NewDescribeFileSystemFileListResponse() (response *DescribeFileSystemFileListResponse) {
	response = &DescribeFileSystemFileListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileSystemFileList(request *DescribeFileSystemFileListRequest) string {
	return c.DescribeFileSystemFileListWithContext(context.Background(), request)
}

func (c *Client) DescribeFileSystemFileListSend(request *DescribeFileSystemFileListRequest) (*DescribeFileSystemFileListResponse, error) {
	statusCode, msg, err := c.DescribeFileSystemFileListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFileSystemFileListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileSystemFileListWithContext(ctx context.Context, request *DescribeFileSystemFileListRequest) string {
	if request == nil {
		request = NewDescribeFileSystemFileListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemFileList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemFileListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileSystemFileListWithContextV2(ctx context.Context, request *DescribeFileSystemFileListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileSystemFileListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemFileList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemFileListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenewFileSystemRequest() (request *RenewFileSystemRequest) {
	request = &RenewFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "RenewFileSystem")
	return
}

func NewRenewFileSystemResponse() (response *RenewFileSystemResponse) {
	response = &RenewFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenewFileSystem(request *RenewFileSystemRequest) string {
	return c.RenewFileSystemWithContext(context.Background(), request)
}

func (c *Client) RenewFileSystemSend(request *RenewFileSystemRequest) (*RenewFileSystemResponse, error) {
	statusCode, msg, err := c.RenewFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RenewFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenewFileSystemWithContext(ctx context.Context, request *RenewFileSystemRequest) string {
	if request == nil {
		request = NewRenewFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RenewFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenewFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenewFileSystemWithContextV2(ctx context.Context, request *RenewFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewRenewFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RenewFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenewFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpgradeFileSystemRequest() (request *UpgradeFileSystemRequest) {
	request = &UpgradeFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpgradeFileSystem")
	return
}

func NewUpgradeFileSystemResponse() (response *UpgradeFileSystemResponse) {
	response = &UpgradeFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpgradeFileSystem(request *UpgradeFileSystemRequest) string {
	return c.UpgradeFileSystemWithContext(context.Background(), request)
}

func (c *Client) UpgradeFileSystemSend(request *UpgradeFileSystemRequest) (*UpgradeFileSystemResponse, error) {
	statusCode, msg, err := c.UpgradeFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpgradeFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpgradeFileSystemWithContext(ctx context.Context, request *UpgradeFileSystemRequest) string {
	if request == nil {
		request = NewUpgradeFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpgradeFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpgradeFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpgradeFileSystemWithContextV2(ctx context.Context, request *UpgradeFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewUpgradeFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpgradeFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpgradeFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
	request = &CreateFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateFileSystem")
	return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
	response = &CreateFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) string {
	return c.CreateFileSystemWithContext(context.Background(), request)
}

func (c *Client) CreateFileSystemSend(request *CreateFileSystemRequest) (*CreateFileSystemResponse, error) {
	statusCode, msg, err := c.CreateFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) string {
	if request == nil {
		request = NewCreateFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateFileSystemWithContextV2(ctx context.Context, request *CreateFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewCreateFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetCapacityAvailableRequest() (request *GetCapacityAvailableRequest) {
	request = &GetCapacityAvailableRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetCapacityAvailable")
	return
}

func NewGetCapacityAvailableResponse() (response *GetCapacityAvailableResponse) {
	response = &GetCapacityAvailableResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetCapacityAvailable(request *GetCapacityAvailableRequest) string {
	return c.GetCapacityAvailableWithContext(context.Background(), request)
}

func (c *Client) GetCapacityAvailableSend(request *GetCapacityAvailableRequest) (*GetCapacityAvailableResponse, error) {
	statusCode, msg, err := c.GetCapacityAvailableWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetCapacityAvailableResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetCapacityAvailableWithContext(ctx context.Context, request *GetCapacityAvailableRequest) string {
	if request == nil {
		request = NewGetCapacityAvailableRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetCapacityAvailable")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCapacityAvailableResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetCapacityAvailableWithContextV2(ctx context.Context, request *GetCapacityAvailableRequest) (int, string, error) {
	if request == nil {
		request = NewGetCapacityAvailableRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetCapacityAvailable")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCapacityAvailableResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetCapacityTotalRequest() (request *GetCapacityTotalRequest) {
	request = &GetCapacityTotalRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetCapacityTotal")
	return
}

func NewGetCapacityTotalResponse() (response *GetCapacityTotalResponse) {
	response = &GetCapacityTotalResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetCapacityTotal(request *GetCapacityTotalRequest) string {
	return c.GetCapacityTotalWithContext(context.Background(), request)
}

func (c *Client) GetCapacityTotalSend(request *GetCapacityTotalRequest) (*GetCapacityTotalResponse, error) {
	statusCode, msg, err := c.GetCapacityTotalWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetCapacityTotalResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetCapacityTotalWithContext(ctx context.Context, request *GetCapacityTotalRequest) string {
	if request == nil {
		request = NewGetCapacityTotalRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetCapacityTotal")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCapacityTotalResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetCapacityTotalWithContextV2(ctx context.Context, request *GetCapacityTotalRequest) (int, string, error) {
	if request == nil {
		request = NewGetCapacityTotalRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetCapacityTotal")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetCapacityTotalResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetLatencyWriteRequest() (request *GetLatencyWriteRequest) {
	request = &GetLatencyWriteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetLatencyWrite")
	return
}

func NewGetLatencyWriteResponse() (response *GetLatencyWriteResponse) {
	response = &GetLatencyWriteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetLatencyWrite(request *GetLatencyWriteRequest) string {
	return c.GetLatencyWriteWithContext(context.Background(), request)
}

func (c *Client) GetLatencyWriteSend(request *GetLatencyWriteRequest) (*GetLatencyWriteResponse, error) {
	statusCode, msg, err := c.GetLatencyWriteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetLatencyWriteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetLatencyWriteWithContext(ctx context.Context, request *GetLatencyWriteRequest) string {
	if request == nil {
		request = NewGetLatencyWriteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetLatencyWrite")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLatencyWriteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetLatencyWriteWithContextV2(ctx context.Context, request *GetLatencyWriteRequest) (int, string, error) {
	if request == nil {
		request = NewGetLatencyWriteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetLatencyWrite")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLatencyWriteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetLatencyReadRequest() (request *GetLatencyReadRequest) {
	request = &GetLatencyReadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetLatencyRead")
	return
}

func NewGetLatencyReadResponse() (response *GetLatencyReadResponse) {
	response = &GetLatencyReadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetLatencyRead(request *GetLatencyReadRequest) string {
	return c.GetLatencyReadWithContext(context.Background(), request)
}

func (c *Client) GetLatencyReadSend(request *GetLatencyReadRequest) (*GetLatencyReadResponse, error) {
	statusCode, msg, err := c.GetLatencyReadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetLatencyReadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetLatencyReadWithContext(ctx context.Context, request *GetLatencyReadRequest) string {
	if request == nil {
		request = NewGetLatencyReadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetLatencyRead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLatencyReadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetLatencyReadWithContextV2(ctx context.Context, request *GetLatencyReadRequest) (int, string, error) {
	if request == nil {
		request = NewGetLatencyReadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetLatencyRead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLatencyReadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetIopsWriteRequest() (request *GetIopsWriteRequest) {
	request = &GetIopsWriteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetIopsWrite")
	return
}

func NewGetIopsWriteResponse() (response *GetIopsWriteResponse) {
	response = &GetIopsWriteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetIopsWrite(request *GetIopsWriteRequest) string {
	return c.GetIopsWriteWithContext(context.Background(), request)
}

func (c *Client) GetIopsWriteSend(request *GetIopsWriteRequest) (*GetIopsWriteResponse, error) {
	statusCode, msg, err := c.GetIopsWriteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetIopsWriteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetIopsWriteWithContext(ctx context.Context, request *GetIopsWriteRequest) string {
	if request == nil {
		request = NewGetIopsWriteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetIopsWrite")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetIopsWriteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetIopsWriteWithContextV2(ctx context.Context, request *GetIopsWriteRequest) (int, string, error) {
	if request == nil {
		request = NewGetIopsWriteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetIopsWrite")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetIopsWriteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetIopsReadRequest() (request *GetIopsReadRequest) {
	request = &GetIopsReadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetIopsRead")
	return
}

func NewGetIopsReadResponse() (response *GetIopsReadResponse) {
	response = &GetIopsReadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetIopsRead(request *GetIopsReadRequest) string {
	return c.GetIopsReadWithContext(context.Background(), request)
}

func (c *Client) GetIopsReadSend(request *GetIopsReadRequest) (*GetIopsReadResponse, error) {
	statusCode, msg, err := c.GetIopsReadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetIopsReadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetIopsReadWithContext(ctx context.Context, request *GetIopsReadRequest) string {
	if request == nil {
		request = NewGetIopsReadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetIopsRead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetIopsReadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetIopsReadWithContextV2(ctx context.Context, request *GetIopsReadRequest) (int, string, error) {
	if request == nil {
		request = NewGetIopsReadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetIopsRead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetIopsReadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBandwidthWriteRequest() (request *GetBandwidthWriteRequest) {
	request = &GetBandwidthWriteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetBandwidthWrite")
	return
}

func NewGetBandwidthWriteResponse() (response *GetBandwidthWriteResponse) {
	response = &GetBandwidthWriteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBandwidthWrite(request *GetBandwidthWriteRequest) string {
	return c.GetBandwidthWriteWithContext(context.Background(), request)
}

func (c *Client) GetBandwidthWriteSend(request *GetBandwidthWriteRequest) (*GetBandwidthWriteResponse, error) {
	statusCode, msg, err := c.GetBandwidthWriteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetBandwidthWriteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBandwidthWriteWithContext(ctx context.Context, request *GetBandwidthWriteRequest) string {
	if request == nil {
		request = NewGetBandwidthWriteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetBandwidthWrite")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBandwidthWriteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBandwidthWriteWithContextV2(ctx context.Context, request *GetBandwidthWriteRequest) (int, string, error) {
	if request == nil {
		request = NewGetBandwidthWriteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetBandwidthWrite")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBandwidthWriteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBandwidthReadRequest() (request *GetBandwidthReadRequest) {
	request = &GetBandwidthReadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetBandwidthRead")
	return
}

func NewGetBandwidthReadResponse() (response *GetBandwidthReadResponse) {
	response = &GetBandwidthReadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBandwidthRead(request *GetBandwidthReadRequest) string {
	return c.GetBandwidthReadWithContext(context.Background(), request)
}

func (c *Client) GetBandwidthReadSend(request *GetBandwidthReadRequest) (*GetBandwidthReadResponse, error) {
	statusCode, msg, err := c.GetBandwidthReadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetBandwidthReadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBandwidthReadWithContext(ctx context.Context, request *GetBandwidthReadRequest) string {
	if request == nil {
		request = NewGetBandwidthReadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetBandwidthRead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBandwidthReadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBandwidthReadWithContextV2(ctx context.Context, request *GetBandwidthReadRequest) (int, string, error) {
	if request == nil {
		request = NewGetBandwidthReadRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetBandwidthRead")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBandwidthReadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePerformanceOnePosixAclRequest() (request *DeletePerformanceOnePosixAclRequest) {
	request = &DeletePerformanceOnePosixAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeletePerformanceOnePosixAcl")
	return
}

func NewDeletePerformanceOnePosixAclResponse() (response *DeletePerformanceOnePosixAclResponse) {
	response = &DeletePerformanceOnePosixAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePerformanceOnePosixAcl(request *DeletePerformanceOnePosixAclRequest) string {
	return c.DeletePerformanceOnePosixAclWithContext(context.Background(), request)
}

func (c *Client) DeletePerformanceOnePosixAclSend(request *DeletePerformanceOnePosixAclRequest) (*DeletePerformanceOnePosixAclResponse, error) {
	statusCode, msg, err := c.DeletePerformanceOnePosixAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeletePerformanceOnePosixAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePerformanceOnePosixAclWithContext(ctx context.Context, request *DeletePerformanceOnePosixAclRequest) string {
	if request == nil {
		request = NewDeletePerformanceOnePosixAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeletePerformanceOnePosixAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePerformanceOnePosixAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePerformanceOnePosixAclWithContextV2(ctx context.Context, request *DeletePerformanceOnePosixAclRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePerformanceOnePosixAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeletePerformanceOnePosixAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePerformanceOnePosixAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdatePerformanceOnePosixAclRequest() (request *UpdatePerformanceOnePosixAclRequest) {
	request = &UpdatePerformanceOnePosixAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceOnePosixAcl")
	return
}

func NewUpdatePerformanceOnePosixAclResponse() (response *UpdatePerformanceOnePosixAclResponse) {
	response = &UpdatePerformanceOnePosixAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePerformanceOnePosixAcl(request *UpdatePerformanceOnePosixAclRequest) string {
	return c.UpdatePerformanceOnePosixAclWithContext(context.Background(), request)
}

func (c *Client) UpdatePerformanceOnePosixAclSend(request *UpdatePerformanceOnePosixAclRequest) (*UpdatePerformanceOnePosixAclResponse, error) {
	statusCode, msg, err := c.UpdatePerformanceOnePosixAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdatePerformanceOnePosixAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdatePerformanceOnePosixAclWithContext(ctx context.Context, request *UpdatePerformanceOnePosixAclRequest) string {
	if request == nil {
		request = NewUpdatePerformanceOnePosixAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceOnePosixAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceOnePosixAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdatePerformanceOnePosixAclWithContextV2(ctx context.Context, request *UpdatePerformanceOnePosixAclRequest) (int, string, error) {
	if request == nil {
		request = NewUpdatePerformanceOnePosixAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceOnePosixAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceOnePosixAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePerformanceOnePosixAclListRequest() (request *DescribePerformanceOnePosixAclListRequest) {
	request = &DescribePerformanceOnePosixAclListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOnePosixAclList")
	return
}

func NewDescribePerformanceOnePosixAclListResponse() (response *DescribePerformanceOnePosixAclListResponse) {
	response = &DescribePerformanceOnePosixAclListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePerformanceOnePosixAclList(request *DescribePerformanceOnePosixAclListRequest) string {
	return c.DescribePerformanceOnePosixAclListWithContext(context.Background(), request)
}

func (c *Client) DescribePerformanceOnePosixAclListSend(request *DescribePerformanceOnePosixAclListRequest) (*DescribePerformanceOnePosixAclListResponse, error) {
	statusCode, msg, err := c.DescribePerformanceOnePosixAclListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribePerformanceOnePosixAclListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePerformanceOnePosixAclListWithContext(ctx context.Context, request *DescribePerformanceOnePosixAclListRequest) string {
	if request == nil {
		request = NewDescribePerformanceOnePosixAclListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOnePosixAclList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePerformanceOnePosixAclListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePerformanceOnePosixAclListWithContextV2(ctx context.Context, request *DescribePerformanceOnePosixAclListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePerformanceOnePosixAclListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOnePosixAclList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePerformanceOnePosixAclListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetPerformanceOnePosixAclRequest() (request *SetPerformanceOnePosixAclRequest) {
	request = &SetPerformanceOnePosixAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "SetPerformanceOnePosixAcl")
	return
}

func NewSetPerformanceOnePosixAclResponse() (response *SetPerformanceOnePosixAclResponse) {
	response = &SetPerformanceOnePosixAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetPerformanceOnePosixAcl(request *SetPerformanceOnePosixAclRequest) string {
	return c.SetPerformanceOnePosixAclWithContext(context.Background(), request)
}

func (c *Client) SetPerformanceOnePosixAclSend(request *SetPerformanceOnePosixAclRequest) (*SetPerformanceOnePosixAclResponse, error) {
	statusCode, msg, err := c.SetPerformanceOnePosixAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetPerformanceOnePosixAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetPerformanceOnePosixAclWithContext(ctx context.Context, request *SetPerformanceOnePosixAclRequest) string {
	if request == nil {
		request = NewSetPerformanceOnePosixAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetPerformanceOnePosixAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetPerformanceOnePosixAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetPerformanceOnePosixAclWithContextV2(ctx context.Context, request *SetPerformanceOnePosixAclRequest) (int, string, error) {
	if request == nil {
		request = NewSetPerformanceOnePosixAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetPerformanceOnePosixAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetPerformanceOnePosixAclResponse()
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewDescribeDirQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
	request = &DeleteFileSystemRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteFileSystem")
	return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
	response = &DeleteFileSystemResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) string {
	return c.DeleteFileSystemWithContext(context.Background(), request)
}

func (c *Client) DeleteFileSystemSend(request *DeleteFileSystemRequest) (*DeleteFileSystemResponse, error) {
	statusCode, msg, err := c.DeleteFileSystemWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteFileSystemResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) string {
	if request == nil {
		request = NewDeleteFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFileSystemResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFileSystemWithContextV2(ctx context.Context, request *DeleteFileSystemRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFileSystemRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteFileSystem")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteFileSystemResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddPerformanceOnePosixAclIpRequest() (request *AddPerformanceOnePosixAclIpRequest) {
	request = &AddPerformanceOnePosixAclIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "AddPerformanceOnePosixAclIp")
	return
}

func NewAddPerformanceOnePosixAclIpResponse() (response *AddPerformanceOnePosixAclIpResponse) {
	response = &AddPerformanceOnePosixAclIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddPerformanceOnePosixAclIp(request *AddPerformanceOnePosixAclIpRequest) string {
	return c.AddPerformanceOnePosixAclIpWithContext(context.Background(), request)
}

func (c *Client) AddPerformanceOnePosixAclIpSend(request *AddPerformanceOnePosixAclIpRequest) (*AddPerformanceOnePosixAclIpResponse, error) {
	statusCode, msg, err := c.AddPerformanceOnePosixAclIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddPerformanceOnePosixAclIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddPerformanceOnePosixAclIpWithContext(ctx context.Context, request *AddPerformanceOnePosixAclIpRequest) string {
	if request == nil {
		request = NewAddPerformanceOnePosixAclIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "AddPerformanceOnePosixAclIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddPerformanceOnePosixAclIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddPerformanceOnePosixAclIpWithContextV2(ctx context.Context, request *AddPerformanceOnePosixAclIpRequest) (int, string, error) {
	if request == nil {
		request = NewAddPerformanceOnePosixAclIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "AddPerformanceOnePosixAclIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddPerformanceOnePosixAclIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRemovePerformanceOnePosixAclIpRequest() (request *RemovePerformanceOnePosixAclIpRequest) {
	request = &RemovePerformanceOnePosixAclIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "RemovePerformanceOnePosixAclIp")
	return
}

func NewRemovePerformanceOnePosixAclIpResponse() (response *RemovePerformanceOnePosixAclIpResponse) {
	response = &RemovePerformanceOnePosixAclIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemovePerformanceOnePosixAclIp(request *RemovePerformanceOnePosixAclIpRequest) string {
	return c.RemovePerformanceOnePosixAclIpWithContext(context.Background(), request)
}

func (c *Client) RemovePerformanceOnePosixAclIpSend(request *RemovePerformanceOnePosixAclIpRequest) (*RemovePerformanceOnePosixAclIpResponse, error) {
	statusCode, msg, err := c.RemovePerformanceOnePosixAclIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RemovePerformanceOnePosixAclIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RemovePerformanceOnePosixAclIpWithContext(ctx context.Context, request *RemovePerformanceOnePosixAclIpRequest) string {
	if request == nil {
		request = NewRemovePerformanceOnePosixAclIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RemovePerformanceOnePosixAclIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRemovePerformanceOnePosixAclIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RemovePerformanceOnePosixAclIpWithContextV2(ctx context.Context, request *RemovePerformanceOnePosixAclIpRequest) (int, string, error) {
	if request == nil {
		request = NewRemovePerformanceOnePosixAclIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RemovePerformanceOnePosixAclIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRemovePerformanceOnePosixAclIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDataMigrateTaskProgressRequest() (request *GetDataMigrateTaskProgressRequest) {
	request = &GetDataMigrateTaskProgressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetDataMigrateTaskProgress")
	return
}

func NewGetDataMigrateTaskProgressResponse() (response *GetDataMigrateTaskProgressResponse) {
	response = &GetDataMigrateTaskProgressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDataMigrateTaskProgress(request *GetDataMigrateTaskProgressRequest) string {
	return c.GetDataMigrateTaskProgressWithContext(context.Background(), request)
}

func (c *Client) GetDataMigrateTaskProgressSend(request *GetDataMigrateTaskProgressRequest) (*GetDataMigrateTaskProgressResponse, error) {
	statusCode, msg, err := c.GetDataMigrateTaskProgressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetDataMigrateTaskProgressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDataMigrateTaskProgressWithContext(ctx context.Context, request *GetDataMigrateTaskProgressRequest) string {
	if request == nil {
		request = NewGetDataMigrateTaskProgressRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetDataMigrateTaskProgress")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDataMigrateTaskProgressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDataMigrateTaskProgressWithContextV2(ctx context.Context, request *GetDataMigrateTaskProgressRequest) (int, string, error) {
	if request == nil {
		request = NewGetDataMigrateTaskProgressRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetDataMigrateTaskProgress")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDataMigrateTaskProgressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataMigrateTaskListRequest() (request *DescribeDataMigrateTaskListRequest) {
	request = &DescribeDataMigrateTaskListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataMigrateTaskList")
	return
}

func NewDescribeDataMigrateTaskListResponse() (response *DescribeDataMigrateTaskListResponse) {
	response = &DescribeDataMigrateTaskListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataMigrateTaskList(request *DescribeDataMigrateTaskListRequest) string {
	return c.DescribeDataMigrateTaskListWithContext(context.Background(), request)
}

func (c *Client) DescribeDataMigrateTaskListSend(request *DescribeDataMigrateTaskListRequest) (*DescribeDataMigrateTaskListResponse, error) {
	statusCode, msg, err := c.DescribeDataMigrateTaskListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataMigrateTaskListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataMigrateTaskListWithContext(ctx context.Context, request *DescribeDataMigrateTaskListRequest) string {
	if request == nil {
		request = NewDescribeDataMigrateTaskListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataMigrateTaskList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataMigrateTaskListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataMigrateTaskListWithContextV2(ctx context.Context, request *DescribeDataMigrateTaskListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataMigrateTaskListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataMigrateTaskList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataMigrateTaskListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartDataMigrateTaskRequest() (request *StartDataMigrateTaskRequest) {
	request = &StartDataMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "StartDataMigrateTask")
	return
}

func NewStartDataMigrateTaskResponse() (response *StartDataMigrateTaskResponse) {
	response = &StartDataMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartDataMigrateTask(request *StartDataMigrateTaskRequest) string {
	return c.StartDataMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) StartDataMigrateTaskSend(request *StartDataMigrateTaskRequest) (*StartDataMigrateTaskResponse, error) {
	statusCode, msg, err := c.StartDataMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StartDataMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartDataMigrateTaskWithContext(ctx context.Context, request *StartDataMigrateTaskRequest) string {
	if request == nil {
		request = NewStartDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "StartDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartDataMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartDataMigrateTaskWithContextV2(ctx context.Context, request *StartDataMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewStartDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "StartDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartDataMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopDataMigrateTaskRequest() (request *StopDataMigrateTaskRequest) {
	request = &StopDataMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "StopDataMigrateTask")
	return
}

func NewStopDataMigrateTaskResponse() (response *StopDataMigrateTaskResponse) {
	response = &StopDataMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopDataMigrateTask(request *StopDataMigrateTaskRequest) string {
	return c.StopDataMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) StopDataMigrateTaskSend(request *StopDataMigrateTaskRequest) (*StopDataMigrateTaskResponse, error) {
	statusCode, msg, err := c.StopDataMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct StopDataMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopDataMigrateTaskWithContext(ctx context.Context, request *StopDataMigrateTaskRequest) string {
	if request == nil {
		request = NewStopDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "StopDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopDataMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopDataMigrateTaskWithContextV2(ctx context.Context, request *StopDataMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewStopDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "StopDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopDataMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDataMigrateTaskRequest() (request *DeleteDataMigrateTaskRequest) {
	request = &DeleteDataMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataMigrateTask")
	return
}

func NewDeleteDataMigrateTaskResponse() (response *DeleteDataMigrateTaskResponse) {
	response = &DeleteDataMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDataMigrateTask(request *DeleteDataMigrateTaskRequest) string {
	return c.DeleteDataMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteDataMigrateTaskSend(request *DeleteDataMigrateTaskRequest) (*DeleteDataMigrateTaskResponse, error) {
	statusCode, msg, err := c.DeleteDataMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDataMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDataMigrateTaskWithContext(ctx context.Context, request *DeleteDataMigrateTaskRequest) string {
	if request == nil {
		request = NewDeleteDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDataMigrateTaskWithContextV2(ctx context.Context, request *DeleteDataMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDataMigrateTaskRequest() (request *UpdateDataMigrateTaskRequest) {
	request = &UpdateDataMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDataMigrateTask")
	return
}

func NewUpdateDataMigrateTaskResponse() (response *UpdateDataMigrateTaskResponse) {
	response = &UpdateDataMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDataMigrateTask(request *UpdateDataMigrateTaskRequest) string {
	return c.UpdateDataMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) UpdateDataMigrateTaskSend(request *UpdateDataMigrateTaskRequest) (*UpdateDataMigrateTaskResponse, error) {
	statusCode, msg, err := c.UpdateDataMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDataMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDataMigrateTaskWithContext(ctx context.Context, request *UpdateDataMigrateTaskRequest) string {
	if request == nil {
		request = NewUpdateDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDataMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDataMigrateTaskWithContextV2(ctx context.Context, request *UpdateDataMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDataMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDataMigrateTaskRequest() (request *CreateDataMigrateTaskRequest) {
	request = &CreateDataMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataMigrateTask")
	return
}

func NewCreateDataMigrateTaskResponse() (response *CreateDataMigrateTaskResponse) {
	response = &CreateDataMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataMigrateTask(request *CreateDataMigrateTaskRequest) string {
	return c.CreateDataMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) CreateDataMigrateTaskSend(request *CreateDataMigrateTaskRequest) (*CreateDataMigrateTaskResponse, error) {
	statusCode, msg, err := c.CreateDataMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDataMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDataMigrateTaskWithContext(ctx context.Context, request *CreateDataMigrateTaskRequest) string {
	if request == nil {
		request = NewCreateDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDataMigrateTaskWithContextV2(ctx context.Context, request *CreateDataMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDataMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeClientInstallInfoRequest() (request *DescribeClientInstallInfoRequest) {
	request = &DescribeClientInstallInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeClientInstallInfo")
	return
}

func NewDescribeClientInstallInfoResponse() (response *DescribeClientInstallInfoResponse) {
	response = &DescribeClientInstallInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClientInstallInfo(request *DescribeClientInstallInfoRequest) string {
	return c.DescribeClientInstallInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeClientInstallInfoSend(request *DescribeClientInstallInfoRequest) (*DescribeClientInstallInfoResponse, error) {
	statusCode, msg, err := c.DescribeClientInstallInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeClientInstallInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClientInstallInfoWithContext(ctx context.Context, request *DescribeClientInstallInfoRequest) string {
	if request == nil {
		request = NewDescribeClientInstallInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeClientInstallInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClientInstallInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClientInstallInfoWithContextV2(ctx context.Context, request *DescribeClientInstallInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClientInstallInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeClientInstallInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClientInstallInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewManageDataFlowTaskRequest() (request *ManageDataFlowTaskRequest) {
	request = &ManageDataFlowTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowTask")
	return
}

func NewManageDataFlowTaskResponse() (response *ManageDataFlowTaskResponse) {
	response = &ManageDataFlowTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ManageDataFlowTask(request *ManageDataFlowTaskRequest) string {
	return c.ManageDataFlowTaskWithContext(context.Background(), request)
}

func (c *Client) ManageDataFlowTaskSend(request *ManageDataFlowTaskRequest) (*ManageDataFlowTaskResponse, error) {
	statusCode, msg, err := c.ManageDataFlowTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ManageDataFlowTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ManageDataFlowTaskWithContext(ctx context.Context, request *ManageDataFlowTaskRequest) string {
	if request == nil {
		request = NewManageDataFlowTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewManageDataFlowTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ManageDataFlowTaskWithContextV2(ctx context.Context, request *ManageDataFlowTaskRequest) (int, string, error) {
	if request == nil {
		request = NewManageDataFlowTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewManageDataFlowTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDataFlowStrategyRequest() (request *CreateDataFlowStrategyRequest) {
	request = &CreateDataFlowStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowStrategy")
	return
}

func NewCreateDataFlowStrategyResponse() (response *CreateDataFlowStrategyResponse) {
	response = &CreateDataFlowStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataFlowStrategy(request *CreateDataFlowStrategyRequest) string {
	return c.CreateDataFlowStrategyWithContext(context.Background(), request)
}

func (c *Client) CreateDataFlowStrategySend(request *CreateDataFlowStrategyRequest) (*CreateDataFlowStrategyResponse, error) {
	statusCode, msg, err := c.CreateDataFlowStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDataFlowStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDataFlowStrategyWithContext(ctx context.Context, request *CreateDataFlowStrategyRequest) string {
	if request == nil {
		request = NewCreateDataFlowStrategyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowStrategy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataFlowStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDataFlowStrategyWithContextV2(ctx context.Context, request *CreateDataFlowStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDataFlowStrategyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowStrategy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataFlowStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowTaskListRequest() (request *DescribeDataFlowTaskListRequest) {
	request = &DescribeDataFlowTaskListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTaskList")
	return
}

func NewDescribeDataFlowTaskListResponse() (response *DescribeDataFlowTaskListResponse) {
	response = &DescribeDataFlowTaskListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowTaskList(request *DescribeDataFlowTaskListRequest) string {
	return c.DescribeDataFlowTaskListWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowTaskListSend(request *DescribeDataFlowTaskListRequest) (*DescribeDataFlowTaskListResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowTaskListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataFlowTaskListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowTaskListWithContext(ctx context.Context, request *DescribeDataFlowTaskListRequest) string {
	if request == nil {
		request = NewDescribeDataFlowTaskListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTaskList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowTaskListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowTaskListWithContextV2(ctx context.Context, request *DescribeDataFlowTaskListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowTaskListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTaskList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowTaskListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateDataFlowTaskRequest() (request *ActivateDataFlowTaskRequest) {
	request = &ActivateDataFlowTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ActivateDataFlowTask")
	return
}

func NewActivateDataFlowTaskResponse() (response *ActivateDataFlowTaskResponse) {
	response = &ActivateDataFlowTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateDataFlowTask(request *ActivateDataFlowTaskRequest) string {
	return c.ActivateDataFlowTaskWithContext(context.Background(), request)
}

func (c *Client) ActivateDataFlowTaskSend(request *ActivateDataFlowTaskRequest) (*ActivateDataFlowTaskResponse, error) {
	statusCode, msg, err := c.ActivateDataFlowTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ActivateDataFlowTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateDataFlowTaskWithContext(ctx context.Context, request *ActivateDataFlowTaskRequest) string {
	if request == nil {
		request = NewActivateDataFlowTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ActivateDataFlowTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateDataFlowTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateDataFlowTaskWithContextV2(ctx context.Context, request *ActivateDataFlowTaskRequest) (int, string, error) {
	if request == nil {
		request = NewActivateDataFlowTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ActivateDataFlowTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateDataFlowTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDataFlowStrategyRequest() (request *DeleteDataFlowStrategyRequest) {
	request = &DeleteDataFlowStrategyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlowStrategy")
	return
}

func NewDeleteDataFlowStrategyResponse() (response *DeleteDataFlowStrategyResponse) {
	response = &DeleteDataFlowStrategyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDataFlowStrategy(request *DeleteDataFlowStrategyRequest) string {
	return c.DeleteDataFlowStrategyWithContext(context.Background(), request)
}

func (c *Client) DeleteDataFlowStrategySend(request *DeleteDataFlowStrategyRequest) (*DeleteDataFlowStrategyResponse, error) {
	statusCode, msg, err := c.DeleteDataFlowStrategyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDataFlowStrategyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDataFlowStrategyWithContext(ctx context.Context, request *DeleteDataFlowStrategyRequest) string {
	if request == nil {
		request = NewDeleteDataFlowStrategyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlowStrategy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDataFlowStrategyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDataFlowStrategyWithContextV2(ctx context.Context, request *DeleteDataFlowStrategyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDataFlowStrategyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlowStrategy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDataFlowStrategyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowStrategyListRequest() (request *DescribeDataFlowStrategyListRequest) {
	request = &DescribeDataFlowStrategyListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategyList")
	return
}

func NewDescribeDataFlowStrategyListResponse() (response *DescribeDataFlowStrategyListResponse) {
	response = &DescribeDataFlowStrategyListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowStrategyList(request *DescribeDataFlowStrategyListRequest) string {
	return c.DescribeDataFlowStrategyListWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowStrategyListSend(request *DescribeDataFlowStrategyListRequest) (*DescribeDataFlowStrategyListResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowStrategyListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataFlowStrategyListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowStrategyListWithContext(ctx context.Context, request *DescribeDataFlowStrategyListRequest) string {
	if request == nil {
		request = NewDescribeDataFlowStrategyListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategyList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowStrategyListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowStrategyListWithContextV2(ctx context.Context, request *DescribeDataFlowStrategyListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowStrategyListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategyList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowStrategyListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCleanRecycledFilesRequest() (request *CleanRecycledFilesRequest) {
	request = &CleanRecycledFilesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CleanRecycledFiles")
	return
}

func NewCleanRecycledFilesResponse() (response *CleanRecycledFilesResponse) {
	response = &CleanRecycledFilesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CleanRecycledFiles(request *CleanRecycledFilesRequest) string {
	return c.CleanRecycledFilesWithContext(context.Background(), request)
}

func (c *Client) CleanRecycledFilesSend(request *CleanRecycledFilesRequest) (*CleanRecycledFilesResponse, error) {
	statusCode, msg, err := c.CleanRecycledFilesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CleanRecycledFilesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CleanRecycledFilesWithContext(ctx context.Context, request *CleanRecycledFilesRequest) string {
	if request == nil {
		request = NewCleanRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CleanRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCleanRecycledFilesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CleanRecycledFilesWithContextV2(ctx context.Context, request *CleanRecycledFilesRequest) (int, string, error) {
	if request == nil {
		request = NewCleanRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CleanRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCleanRecycledFilesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCleanRecycledFilesRequest() (request *DeleteCleanRecycledFilesRequest) {
	request = &DeleteCleanRecycledFilesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteCleanRecycledFiles")
	return
}

func NewDeleteCleanRecycledFilesResponse() (response *DeleteCleanRecycledFilesResponse) {
	response = &DeleteCleanRecycledFilesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCleanRecycledFiles(request *DeleteCleanRecycledFilesRequest) string {
	return c.DeleteCleanRecycledFilesWithContext(context.Background(), request)
}

func (c *Client) DeleteCleanRecycledFilesSend(request *DeleteCleanRecycledFilesRequest) (*DeleteCleanRecycledFilesResponse, error) {
	statusCode, msg, err := c.DeleteCleanRecycledFilesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCleanRecycledFilesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCleanRecycledFilesWithContext(ctx context.Context, request *DeleteCleanRecycledFilesRequest) string {
	if request == nil {
		request = NewDeleteCleanRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteCleanRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCleanRecycledFilesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCleanRecycledFilesWithContextV2(ctx context.Context, request *DeleteCleanRecycledFilesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCleanRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteCleanRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCleanRecycledFilesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRecycleBinConfigRequest() (request *DeleteRecycleBinConfigRequest) {
	request = &DeleteRecycleBinConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycleBinConfig")
	return
}

func NewDeleteRecycleBinConfigResponse() (response *DeleteRecycleBinConfigResponse) {
	response = &DeleteRecycleBinConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecycleBinConfig(request *DeleteRecycleBinConfigRequest) string {
	return c.DeleteRecycleBinConfigWithContext(context.Background(), request)
}

func (c *Client) DeleteRecycleBinConfigSend(request *DeleteRecycleBinConfigRequest) (*DeleteRecycleBinConfigResponse, error) {
	statusCode, msg, err := c.DeleteRecycleBinConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteRecycleBinConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRecycleBinConfigWithContext(ctx context.Context, request *DeleteRecycleBinConfigRequest) string {
	if request == nil {
		request = NewDeleteRecycleBinConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycleBinConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteRecycleBinConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRecycleBinConfigWithContextV2(ctx context.Context, request *DeleteRecycleBinConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRecycleBinConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycleBinConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteRecycleBinConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRecycledFileListRequest() (request *DeleteRecycledFileListRequest) {
	request = &DeleteRecycledFileListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycledFileList")
	return
}

func NewDeleteRecycledFileListResponse() (response *DeleteRecycledFileListResponse) {
	response = &DeleteRecycledFileListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecycledFileList(request *DeleteRecycledFileListRequest) string {
	return c.DeleteRecycledFileListWithContext(context.Background(), request)
}

func (c *Client) DeleteRecycledFileListSend(request *DeleteRecycledFileListRequest) (*DeleteRecycledFileListResponse, error) {
	statusCode, msg, err := c.DeleteRecycledFileListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteRecycledFileListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRecycledFileListWithContext(ctx context.Context, request *DeleteRecycledFileListRequest) string {
	if request == nil {
		request = NewDeleteRecycledFileListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycledFileList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecycledFileListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRecycledFileListWithContextV2(ctx context.Context, request *DeleteRecycledFileListRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRecycledFileListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycledFileList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecycledFileListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRecycleBinConfigRequest() (request *GetRecycleBinConfigRequest) {
	request = &GetRecycleBinConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRecycleBinConfig")
	return
}

func NewGetRecycleBinConfigResponse() (response *GetRecycleBinConfigResponse) {
	response = &GetRecycleBinConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRecycleBinConfig(request *GetRecycleBinConfigRequest) string {
	return c.GetRecycleBinConfigWithContext(context.Background(), request)
}

func (c *Client) GetRecycleBinConfigSend(request *GetRecycleBinConfigRequest) (*GetRecycleBinConfigResponse, error) {
	statusCode, msg, err := c.GetRecycleBinConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRecycleBinConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRecycleBinConfigWithContext(ctx context.Context, request *GetRecycleBinConfigRequest) string {
	if request == nil {
		request = NewGetRecycleBinConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRecycleBinConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetRecycleBinConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRecycleBinConfigWithContextV2(ctx context.Context, request *GetRecycleBinConfigRequest) (int, string, error) {
	if request == nil {
		request = NewGetRecycleBinConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRecycleBinConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetRecycleBinConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetRecycleBinConfigRequest() (request *SetRecycleBinConfigRequest) {
	request = &SetRecycleBinConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "SetRecycleBinConfig")
	return
}

func NewSetRecycleBinConfigResponse() (response *SetRecycleBinConfigResponse) {
	response = &SetRecycleBinConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetRecycleBinConfig(request *SetRecycleBinConfigRequest) string {
	return c.SetRecycleBinConfigWithContext(context.Background(), request)
}

func (c *Client) SetRecycleBinConfigSend(request *SetRecycleBinConfigRequest) (*SetRecycleBinConfigResponse, error) {
	statusCode, msg, err := c.SetRecycleBinConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetRecycleBinConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetRecycleBinConfigWithContext(ctx context.Context, request *SetRecycleBinConfigRequest) string {
	if request == nil {
		request = NewSetRecycleBinConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetRecycleBinConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetRecycleBinConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetRecycleBinConfigWithContextV2(ctx context.Context, request *SetRecycleBinConfigRequest) (int, string, error) {
	if request == nil {
		request = NewSetRecycleBinConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetRecycleBinConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetRecycleBinConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRecycledFileListRequest() (request *DescribeRecycledFileListRequest) {
	request = &DescribeRecycledFileListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeRecycledFileList")
	return
}

func NewDescribeRecycledFileListResponse() (response *DescribeRecycledFileListResponse) {
	response = &DescribeRecycledFileListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRecycledFileList(request *DescribeRecycledFileListRequest) string {
	return c.DescribeRecycledFileListWithContext(context.Background(), request)
}

func (c *Client) DescribeRecycledFileListSend(request *DescribeRecycledFileListRequest) (*DescribeRecycledFileListResponse, error) {
	statusCode, msg, err := c.DescribeRecycledFileListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeRecycledFileListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRecycledFileListWithContext(ctx context.Context, request *DescribeRecycledFileListRequest) string {
	if request == nil {
		request = NewDescribeRecycledFileListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeRecycledFileList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRecycledFileListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRecycledFileListWithContextV2(ctx context.Context, request *DescribeRecycledFileListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRecycledFileListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeRecycledFileList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRecycledFileListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRecycledFilesRequest() (request *DeleteRecycledFilesRequest) {
	request = &DeleteRecycledFilesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycledFiles")
	return
}

func NewDeleteRecycledFilesResponse() (response *DeleteRecycledFilesResponse) {
	response = &DeleteRecycledFilesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecycledFiles(request *DeleteRecycledFilesRequest) string {
	return c.DeleteRecycledFilesWithContext(context.Background(), request)
}

func (c *Client) DeleteRecycledFilesSend(request *DeleteRecycledFilesRequest) (*DeleteRecycledFilesResponse, error) {
	statusCode, msg, err := c.DeleteRecycledFilesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteRecycledFilesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRecycledFilesWithContext(ctx context.Context, request *DeleteRecycledFilesRequest) string {
	if request == nil {
		request = NewDeleteRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteRecycledFilesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRecycledFilesWithContextV2(ctx context.Context, request *DeleteRecycledFilesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteRecycledFilesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreRecycledFilesRequest() (request *RestoreRecycledFilesRequest) {
	request = &RestoreRecycledFilesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "RestoreRecycledFiles")
	return
}

func NewRestoreRecycledFilesResponse() (response *RestoreRecycledFilesResponse) {
	response = &RestoreRecycledFilesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreRecycledFiles(request *RestoreRecycledFilesRequest) string {
	return c.RestoreRecycledFilesWithContext(context.Background(), request)
}

func (c *Client) RestoreRecycledFilesSend(request *RestoreRecycledFilesRequest) (*RestoreRecycledFilesResponse, error) {
	statusCode, msg, err := c.RestoreRecycledFilesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RestoreRecycledFilesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreRecycledFilesWithContext(ctx context.Context, request *RestoreRecycledFilesRequest) string {
	if request == nil {
		request = NewRestoreRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RestoreRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreRecycledFilesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreRecycledFilesWithContextV2(ctx context.Context, request *RestoreRecycledFilesRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreRecycledFilesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RestoreRecycledFiles")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreRecycledFilesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeClusterInfoRequest() (request *DescribeClusterInfoRequest) {
	request = &DescribeClusterInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeClusterInfo")
	return
}

func NewDescribeClusterInfoResponse() (response *DescribeClusterInfoResponse) {
	response = &DescribeClusterInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeClusterInfo(request *DescribeClusterInfoRequest) string {
	return c.DescribeClusterInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterInfoSend(request *DescribeClusterInfoRequest) (*DescribeClusterInfoResponse, error) {
	statusCode, msg, err := c.DescribeClusterInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeClusterInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeClusterInfoWithContext(ctx context.Context, request *DescribeClusterInfoRequest) string {
	if request == nil {
		request = NewDescribeClusterInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeClusterInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeClusterInfoWithContextV2(ctx context.Context, request *DescribeClusterInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeClusterInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeClusterInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdatePerformanceNfsAclIpRequest() (request *UpdatePerformanceNfsAclIpRequest) {
	request = &UpdatePerformanceNfsAclIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceNfsAclIp")
	return
}

func NewUpdatePerformanceNfsAclIpResponse() (response *UpdatePerformanceNfsAclIpResponse) {
	response = &UpdatePerformanceNfsAclIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePerformanceNfsAclIp(request *UpdatePerformanceNfsAclIpRequest) string {
	return c.UpdatePerformanceNfsAclIpWithContext(context.Background(), request)
}

func (c *Client) UpdatePerformanceNfsAclIpSend(request *UpdatePerformanceNfsAclIpRequest) (*UpdatePerformanceNfsAclIpResponse, error) {
	statusCode, msg, err := c.UpdatePerformanceNfsAclIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdatePerformanceNfsAclIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdatePerformanceNfsAclIpWithContext(ctx context.Context, request *UpdatePerformanceNfsAclIpRequest) string {
	if request == nil {
		request = NewUpdatePerformanceNfsAclIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceNfsAclIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceNfsAclIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdatePerformanceNfsAclIpWithContextV2(ctx context.Context, request *UpdatePerformanceNfsAclIpRequest) (int, string, error) {
	if request == nil {
		request = NewUpdatePerformanceNfsAclIpRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceNfsAclIp")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceNfsAclIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRemovePerformanceNfsAclClientRequest() (request *RemovePerformanceNfsAclClientRequest) {
	request = &RemovePerformanceNfsAclClientRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "RemovePerformanceNfsAclClient")
	return
}

func NewRemovePerformanceNfsAclClientResponse() (response *RemovePerformanceNfsAclClientResponse) {
	response = &RemovePerformanceNfsAclClientResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RemovePerformanceNfsAclClient(request *RemovePerformanceNfsAclClientRequest) string {
	return c.RemovePerformanceNfsAclClientWithContext(context.Background(), request)
}

func (c *Client) RemovePerformanceNfsAclClientSend(request *RemovePerformanceNfsAclClientRequest) (*RemovePerformanceNfsAclClientResponse, error) {
	statusCode, msg, err := c.RemovePerformanceNfsAclClientWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RemovePerformanceNfsAclClientResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RemovePerformanceNfsAclClientWithContext(ctx context.Context, request *RemovePerformanceNfsAclClientRequest) string {
	if request == nil {
		request = NewRemovePerformanceNfsAclClientRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RemovePerformanceNfsAclClient")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRemovePerformanceNfsAclClientResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RemovePerformanceNfsAclClientWithContextV2(ctx context.Context, request *RemovePerformanceNfsAclClientRequest) (int, string, error) {
	if request == nil {
		request = NewRemovePerformanceNfsAclClientRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "RemovePerformanceNfsAclClient")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRemovePerformanceNfsAclClientResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddPerformanceNfsAclClientRequest() (request *AddPerformanceNfsAclClientRequest) {
	request = &AddPerformanceNfsAclClientRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "AddPerformanceNfsAclClient")
	return
}

func NewAddPerformanceNfsAclClientResponse() (response *AddPerformanceNfsAclClientResponse) {
	response = &AddPerformanceNfsAclClientResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddPerformanceNfsAclClient(request *AddPerformanceNfsAclClientRequest) string {
	return c.AddPerformanceNfsAclClientWithContext(context.Background(), request)
}

func (c *Client) AddPerformanceNfsAclClientSend(request *AddPerformanceNfsAclClientRequest) (*AddPerformanceNfsAclClientResponse, error) {
	statusCode, msg, err := c.AddPerformanceNfsAclClientWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AddPerformanceNfsAclClientResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddPerformanceNfsAclClientWithContext(ctx context.Context, request *AddPerformanceNfsAclClientRequest) string {
	if request == nil {
		request = NewAddPerformanceNfsAclClientRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "AddPerformanceNfsAclClient")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddPerformanceNfsAclClientResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddPerformanceNfsAclClientWithContextV2(ctx context.Context, request *AddPerformanceNfsAclClientRequest) (int, string, error) {
	if request == nil {
		request = NewAddPerformanceNfsAclClientRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "AddPerformanceNfsAclClient")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddPerformanceNfsAclClientResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePerformanceOneNfsAclRequest() (request *DeletePerformanceOneNfsAclRequest) {
	request = &DeletePerformanceOneNfsAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeletePerformanceOneNfsAcl")
	return
}

func NewDeletePerformanceOneNfsAclResponse() (response *DeletePerformanceOneNfsAclResponse) {
	response = &DeletePerformanceOneNfsAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePerformanceOneNfsAcl(request *DeletePerformanceOneNfsAclRequest) string {
	return c.DeletePerformanceOneNfsAclWithContext(context.Background(), request)
}

func (c *Client) DeletePerformanceOneNfsAclSend(request *DeletePerformanceOneNfsAclRequest) (*DeletePerformanceOneNfsAclResponse, error) {
	statusCode, msg, err := c.DeletePerformanceOneNfsAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeletePerformanceOneNfsAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePerformanceOneNfsAclWithContext(ctx context.Context, request *DeletePerformanceOneNfsAclRequest) string {
	if request == nil {
		request = NewDeletePerformanceOneNfsAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeletePerformanceOneNfsAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeletePerformanceOneNfsAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePerformanceOneNfsAclWithContextV2(ctx context.Context, request *DeletePerformanceOneNfsAclRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePerformanceOneNfsAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeletePerformanceOneNfsAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeletePerformanceOneNfsAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetPerformanceOneNfsAclRequest() (request *SetPerformanceOneNfsAclRequest) {
	request = &SetPerformanceOneNfsAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "SetPerformanceOneNfsAcl")
	return
}

func NewSetPerformanceOneNfsAclResponse() (response *SetPerformanceOneNfsAclResponse) {
	response = &SetPerformanceOneNfsAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetPerformanceOneNfsAcl(request *SetPerformanceOneNfsAclRequest) string {
	return c.SetPerformanceOneNfsAclWithContext(context.Background(), request)
}

func (c *Client) SetPerformanceOneNfsAclSend(request *SetPerformanceOneNfsAclRequest) (*SetPerformanceOneNfsAclResponse, error) {
	statusCode, msg, err := c.SetPerformanceOneNfsAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetPerformanceOneNfsAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetPerformanceOneNfsAclWithContext(ctx context.Context, request *SetPerformanceOneNfsAclRequest) string {
	if request == nil {
		request = NewSetPerformanceOneNfsAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetPerformanceOneNfsAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetPerformanceOneNfsAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetPerformanceOneNfsAclWithContextV2(ctx context.Context, request *SetPerformanceOneNfsAclRequest) (int, string, error) {
	if request == nil {
		request = NewSetPerformanceOneNfsAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetPerformanceOneNfsAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetPerformanceOneNfsAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePerformanceOneNfsAclListRequest() (request *DescribePerformanceOneNfsAclListRequest) {
	request = &DescribePerformanceOneNfsAclListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOneNfsAclList")
	return
}

func NewDescribePerformanceOneNfsAclListResponse() (response *DescribePerformanceOneNfsAclListResponse) {
	response = &DescribePerformanceOneNfsAclListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePerformanceOneNfsAclList(request *DescribePerformanceOneNfsAclListRequest) string {
	return c.DescribePerformanceOneNfsAclListWithContext(context.Background(), request)
}

func (c *Client) DescribePerformanceOneNfsAclListSend(request *DescribePerformanceOneNfsAclListRequest) (*DescribePerformanceOneNfsAclListResponse, error) {
	statusCode, msg, err := c.DescribePerformanceOneNfsAclListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribePerformanceOneNfsAclListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePerformanceOneNfsAclListWithContext(ctx context.Context, request *DescribePerformanceOneNfsAclListRequest) string {
	if request == nil {
		request = NewDescribePerformanceOneNfsAclListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOneNfsAclList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePerformanceOneNfsAclListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePerformanceOneNfsAclListWithContextV2(ctx context.Context, request *DescribePerformanceOneNfsAclListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePerformanceOneNfsAclListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOneNfsAclList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePerformanceOneNfsAclListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFileSystemNfsClientInfoRequest() (request *DescribeFileSystemNfsClientInfoRequest) {
	request = &DescribeFileSystemNfsClientInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemNfsClientInfo")
	return
}

func NewDescribeFileSystemNfsClientInfoResponse() (response *DescribeFileSystemNfsClientInfoResponse) {
	response = &DescribeFileSystemNfsClientInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileSystemNfsClientInfo(request *DescribeFileSystemNfsClientInfoRequest) string {
	return c.DescribeFileSystemNfsClientInfoWithContext(context.Background(), request)
}

func (c *Client) DescribeFileSystemNfsClientInfoSend(request *DescribeFileSystemNfsClientInfoRequest) (*DescribeFileSystemNfsClientInfoResponse, error) {
	statusCode, msg, err := c.DescribeFileSystemNfsClientInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFileSystemNfsClientInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileSystemNfsClientInfoWithContext(ctx context.Context, request *DescribeFileSystemNfsClientInfoRequest) string {
	if request == nil {
		request = NewDescribeFileSystemNfsClientInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemNfsClientInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemNfsClientInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileSystemNfsClientInfoWithContextV2(ctx context.Context, request *DescribeFileSystemNfsClientInfoRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileSystemNfsClientInfoRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileSystemNfsClientInfo")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileSystemNfsClientInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetFileSystemResourceProtectRequest() (request *SetFileSystemResourceProtectRequest) {
	request = &SetFileSystemResourceProtectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "SetFileSystemResourceProtect")
	return
}

func NewSetFileSystemResourceProtectResponse() (response *SetFileSystemResourceProtectResponse) {
	response = &SetFileSystemResourceProtectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetFileSystemResourceProtect(request *SetFileSystemResourceProtectRequest) string {
	return c.SetFileSystemResourceProtectWithContext(context.Background(), request)
}

func (c *Client) SetFileSystemResourceProtectSend(request *SetFileSystemResourceProtectRequest) (*SetFileSystemResourceProtectResponse, error) {
	statusCode, msg, err := c.SetFileSystemResourceProtectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct SetFileSystemResourceProtectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetFileSystemResourceProtectWithContext(ctx context.Context, request *SetFileSystemResourceProtectRequest) string {
	if request == nil {
		request = NewSetFileSystemResourceProtectRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetFileSystemResourceProtect")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetFileSystemResourceProtectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetFileSystemResourceProtectWithContextV2(ctx context.Context, request *SetFileSystemResourceProtectRequest) (int, string, error) {
	if request == nil {
		request = NewSetFileSystemResourceProtectRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "SetFileSystemResourceProtect")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetFileSystemResourceProtectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFileDeletePolicyListRequest() (request *DescribeFileDeletePolicyListRequest) {
	request = &DescribeFileDeletePolicyListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileDeletePolicyList")
	return
}

func NewDescribeFileDeletePolicyListResponse() (response *DescribeFileDeletePolicyListResponse) {
	response = &DescribeFileDeletePolicyListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileDeletePolicyList(request *DescribeFileDeletePolicyListRequest) string {
	return c.DescribeFileDeletePolicyListWithContext(context.Background(), request)
}

func (c *Client) DescribeFileDeletePolicyListSend(request *DescribeFileDeletePolicyListRequest) (*DescribeFileDeletePolicyListResponse, error) {
	statusCode, msg, err := c.DescribeFileDeletePolicyListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFileDeletePolicyListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileDeletePolicyListWithContext(ctx context.Context, request *DescribeFileDeletePolicyListRequest) string {
	if request == nil {
		request = NewDescribeFileDeletePolicyListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileDeletePolicyList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileDeletePolicyListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileDeletePolicyListWithContextV2(ctx context.Context, request *DescribeFileDeletePolicyListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileDeletePolicyListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileDeletePolicyList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileDeletePolicyListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewEnableFileDeletePolicyRequest() (request *EnableFileDeletePolicyRequest) {
	request = &EnableFileDeletePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "EnableFileDeletePolicy")
	return
}

func NewEnableFileDeletePolicyResponse() (response *EnableFileDeletePolicyResponse) {
	response = &EnableFileDeletePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) EnableFileDeletePolicy(request *EnableFileDeletePolicyRequest) string {
	return c.EnableFileDeletePolicyWithContext(context.Background(), request)
}

func (c *Client) EnableFileDeletePolicySend(request *EnableFileDeletePolicyRequest) (*EnableFileDeletePolicyResponse, error) {
	statusCode, msg, err := c.EnableFileDeletePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct EnableFileDeletePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EnableFileDeletePolicyWithContext(ctx context.Context, request *EnableFileDeletePolicyRequest) string {
	if request == nil {
		request = NewEnableFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "EnableFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableFileDeletePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EnableFileDeletePolicyWithContextV2(ctx context.Context, request *EnableFileDeletePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewEnableFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "EnableFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewEnableFileDeletePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisableFileDeletePolicyRequest() (request *DisableFileDeletePolicyRequest) {
	request = &DisableFileDeletePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DisableFileDeletePolicy")
	return
}

func NewDisableFileDeletePolicyResponse() (response *DisableFileDeletePolicyResponse) {
	response = &DisableFileDeletePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisableFileDeletePolicy(request *DisableFileDeletePolicyRequest) string {
	return c.DisableFileDeletePolicyWithContext(context.Background(), request)
}

func (c *Client) DisableFileDeletePolicySend(request *DisableFileDeletePolicyRequest) (*DisableFileDeletePolicyResponse, error) {
	statusCode, msg, err := c.DisableFileDeletePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisableFileDeletePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisableFileDeletePolicyWithContext(ctx context.Context, request *DisableFileDeletePolicyRequest) string {
	if request == nil {
		request = NewDisableFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DisableFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableFileDeletePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisableFileDeletePolicyWithContextV2(ctx context.Context, request *DisableFileDeletePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDisableFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DisableFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisableFileDeletePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeFileDeletePolicyRequest() (request *DescribeFileDeletePolicyRequest) {
	request = &DescribeFileDeletePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileDeletePolicy")
	return
}

func NewDescribeFileDeletePolicyResponse() (response *DescribeFileDeletePolicyResponse) {
	response = &DescribeFileDeletePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeFileDeletePolicy(request *DescribeFileDeletePolicyRequest) string {
	return c.DescribeFileDeletePolicyWithContext(context.Background(), request)
}

func (c *Client) DescribeFileDeletePolicySend(request *DescribeFileDeletePolicyRequest) (*DescribeFileDeletePolicyResponse, error) {
	statusCode, msg, err := c.DescribeFileDeletePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeFileDeletePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeFileDeletePolicyWithContext(ctx context.Context, request *DescribeFileDeletePolicyRequest) string {
	if request == nil {
		request = NewDescribeFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileDeletePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeFileDeletePolicyWithContextV2(ctx context.Context, request *DescribeFileDeletePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeFileDeletePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteFileDeletePolicyRequest() (request *DeleteFileDeletePolicyRequest) {
	request = &DeleteFileDeletePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteFileDeletePolicy")
	return
}

func NewDeleteFileDeletePolicyResponse() (response *DeleteFileDeletePolicyResponse) {
	response = &DeleteFileDeletePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteFileDeletePolicy(request *DeleteFileDeletePolicyRequest) string {
	return c.DeleteFileDeletePolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteFileDeletePolicySend(request *DeleteFileDeletePolicyRequest) (*DeleteFileDeletePolicyResponse, error) {
	statusCode, msg, err := c.DeleteFileDeletePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteFileDeletePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteFileDeletePolicyWithContext(ctx context.Context, request *DeleteFileDeletePolicyRequest) string {
	if request == nil {
		request = NewDeleteFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteFileDeletePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteFileDeletePolicyWithContextV2(ctx context.Context, request *DeleteFileDeletePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteFileDeletePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateFileDeletePolicyRequest() (request *UpdateFileDeletePolicyRequest) {
	request = &UpdateFileDeletePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdateFileDeletePolicy")
	return
}

func NewUpdateFileDeletePolicyResponse() (response *UpdateFileDeletePolicyResponse) {
	response = &UpdateFileDeletePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateFileDeletePolicy(request *UpdateFileDeletePolicyRequest) string {
	return c.UpdateFileDeletePolicyWithContext(context.Background(), request)
}

func (c *Client) UpdateFileDeletePolicySend(request *UpdateFileDeletePolicyRequest) (*UpdateFileDeletePolicyResponse, error) {
	statusCode, msg, err := c.UpdateFileDeletePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateFileDeletePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateFileDeletePolicyWithContext(ctx context.Context, request *UpdateFileDeletePolicyRequest) string {
	if request == nil {
		request = NewUpdateFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateFileDeletePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateFileDeletePolicyWithContextV2(ctx context.Context, request *UpdateFileDeletePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "UpdateFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateFileDeletePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateFileDeletePolicyRequest() (request *CreateFileDeletePolicyRequest) {
	request = &CreateFileDeletePolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateFileDeletePolicy")
	return
}

func NewCreateFileDeletePolicyResponse() (response *CreateFileDeletePolicyResponse) {
	response = &CreateFileDeletePolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFileDeletePolicy(request *CreateFileDeletePolicyRequest) string {
	return c.CreateFileDeletePolicyWithContext(context.Background(), request)
}

func (c *Client) CreateFileDeletePolicySend(request *CreateFileDeletePolicyRequest) (*CreateFileDeletePolicyResponse, error) {
	statusCode, msg, err := c.CreateFileDeletePolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateFileDeletePolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateFileDeletePolicyWithContext(ctx context.Context, request *CreateFileDeletePolicyRequest) string {
	if request == nil {
		request = NewCreateFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFileDeletePolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateFileDeletePolicyWithContextV2(ctx context.Context, request *CreateFileDeletePolicyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateFileDeletePolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateFileDeletePolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateFileDeletePolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowStrategySubscribeRequest() (request *DescribeDataFlowStrategySubscribeRequest) {
	request = &DescribeDataFlowStrategySubscribeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategySubscribe")
	return
}

func NewDescribeDataFlowStrategySubscribeResponse() (response *DescribeDataFlowStrategySubscribeResponse) {
	response = &DescribeDataFlowStrategySubscribeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowStrategySubscribe(request *DescribeDataFlowStrategySubscribeRequest) string {
	return c.DescribeDataFlowStrategySubscribeWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowStrategySubscribeSend(request *DescribeDataFlowStrategySubscribeRequest) (*DescribeDataFlowStrategySubscribeResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowStrategySubscribeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataFlowStrategySubscribeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowStrategySubscribeWithContext(ctx context.Context, request *DescribeDataFlowStrategySubscribeRequest) string {
	if request == nil {
		request = NewDescribeDataFlowStrategySubscribeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategySubscribe")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowStrategySubscribeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowStrategySubscribeWithContextV2(ctx context.Context, request *DescribeDataFlowStrategySubscribeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowStrategySubscribeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategySubscribe")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowStrategySubscribeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewManageDataFlowStrategySubscribeRequest() (request *ManageDataFlowStrategySubscribeRequest) {
	request = &ManageDataFlowStrategySubscribeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowStrategySubscribe")
	return
}

func NewManageDataFlowStrategySubscribeResponse() (response *ManageDataFlowStrategySubscribeResponse) {
	response = &ManageDataFlowStrategySubscribeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ManageDataFlowStrategySubscribe(request *ManageDataFlowStrategySubscribeRequest) string {
	return c.ManageDataFlowStrategySubscribeWithContext(context.Background(), request)
}

func (c *Client) ManageDataFlowStrategySubscribeSend(request *ManageDataFlowStrategySubscribeRequest) (*ManageDataFlowStrategySubscribeResponse, error) {
	statusCode, msg, err := c.ManageDataFlowStrategySubscribeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ManageDataFlowStrategySubscribeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ManageDataFlowStrategySubscribeWithContext(ctx context.Context, request *ManageDataFlowStrategySubscribeRequest) string {
	if request == nil {
		request = NewManageDataFlowStrategySubscribeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowStrategySubscribe")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewManageDataFlowStrategySubscribeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ManageDataFlowStrategySubscribeWithContextV2(ctx context.Context, request *ManageDataFlowStrategySubscribeRequest) (int, string, error) {
	if request == nil {
		request = NewManageDataFlowStrategySubscribeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ManageDataFlowStrategySubscribe")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewManageDataFlowStrategySubscribeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRemoteCachePutLatencyRequest() (request *GetRemoteCachePutLatencyRequest) {
	request = &GetRemoteCachePutLatencyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCachePutLatency")
	return
}

func NewGetRemoteCachePutLatencyResponse() (response *GetRemoteCachePutLatencyResponse) {
	response = &GetRemoteCachePutLatencyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRemoteCachePutLatency(request *GetRemoteCachePutLatencyRequest) string {
	return c.GetRemoteCachePutLatencyWithContext(context.Background(), request)
}

func (c *Client) GetRemoteCachePutLatencySend(request *GetRemoteCachePutLatencyRequest) (*GetRemoteCachePutLatencyResponse, error) {
	statusCode, msg, err := c.GetRemoteCachePutLatencyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRemoteCachePutLatencyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRemoteCachePutLatencyWithContext(ctx context.Context, request *GetRemoteCachePutLatencyRequest) string {
	if request == nil {
		request = NewGetRemoteCachePutLatencyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCachePutLatency")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCachePutLatencyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRemoteCachePutLatencyWithContextV2(ctx context.Context, request *GetRemoteCachePutLatencyRequest) (int, string, error) {
	if request == nil {
		request = NewGetRemoteCachePutLatencyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCachePutLatency")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCachePutLatencyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRemoteCacheGetLatencyRequest() (request *GetRemoteCacheGetLatencyRequest) {
	request = &GetRemoteCacheGetLatencyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheGetLatency")
	return
}

func NewGetRemoteCacheGetLatencyResponse() (response *GetRemoteCacheGetLatencyResponse) {
	response = &GetRemoteCacheGetLatencyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRemoteCacheGetLatency(request *GetRemoteCacheGetLatencyRequest) string {
	return c.GetRemoteCacheGetLatencyWithContext(context.Background(), request)
}

func (c *Client) GetRemoteCacheGetLatencySend(request *GetRemoteCacheGetLatencyRequest) (*GetRemoteCacheGetLatencyResponse, error) {
	statusCode, msg, err := c.GetRemoteCacheGetLatencyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRemoteCacheGetLatencyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRemoteCacheGetLatencyWithContext(ctx context.Context, request *GetRemoteCacheGetLatencyRequest) string {
	if request == nil {
		request = NewGetRemoteCacheGetLatencyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheGetLatency")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheGetLatencyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRemoteCacheGetLatencyWithContextV2(ctx context.Context, request *GetRemoteCacheGetLatencyRequest) (int, string, error) {
	if request == nil {
		request = NewGetRemoteCacheGetLatencyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheGetLatency")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheGetLatencyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRemoteCachePutThroughputRequest() (request *GetRemoteCachePutThroughputRequest) {
	request = &GetRemoteCachePutThroughputRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCachePutThroughput")
	return
}

func NewGetRemoteCachePutThroughputResponse() (response *GetRemoteCachePutThroughputResponse) {
	response = &GetRemoteCachePutThroughputResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRemoteCachePutThroughput(request *GetRemoteCachePutThroughputRequest) string {
	return c.GetRemoteCachePutThroughputWithContext(context.Background(), request)
}

func (c *Client) GetRemoteCachePutThroughputSend(request *GetRemoteCachePutThroughputRequest) (*GetRemoteCachePutThroughputResponse, error) {
	statusCode, msg, err := c.GetRemoteCachePutThroughputWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRemoteCachePutThroughputResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRemoteCachePutThroughputWithContext(ctx context.Context, request *GetRemoteCachePutThroughputRequest) string {
	if request == nil {
		request = NewGetRemoteCachePutThroughputRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCachePutThroughput")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCachePutThroughputResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRemoteCachePutThroughputWithContextV2(ctx context.Context, request *GetRemoteCachePutThroughputRequest) (int, string, error) {
	if request == nil {
		request = NewGetRemoteCachePutThroughputRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCachePutThroughput")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCachePutThroughputResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRemoteCacheGetThroughputRequest() (request *GetRemoteCacheGetThroughputRequest) {
	request = &GetRemoteCacheGetThroughputRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheGetThroughput")
	return
}

func NewGetRemoteCacheGetThroughputResponse() (response *GetRemoteCacheGetThroughputResponse) {
	response = &GetRemoteCacheGetThroughputResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRemoteCacheGetThroughput(request *GetRemoteCacheGetThroughputRequest) string {
	return c.GetRemoteCacheGetThroughputWithContext(context.Background(), request)
}

func (c *Client) GetRemoteCacheGetThroughputSend(request *GetRemoteCacheGetThroughputRequest) (*GetRemoteCacheGetThroughputResponse, error) {
	statusCode, msg, err := c.GetRemoteCacheGetThroughputWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRemoteCacheGetThroughputResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRemoteCacheGetThroughputWithContext(ctx context.Context, request *GetRemoteCacheGetThroughputRequest) string {
	if request == nil {
		request = NewGetRemoteCacheGetThroughputRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheGetThroughput")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheGetThroughputResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRemoteCacheGetThroughputWithContextV2(ctx context.Context, request *GetRemoteCacheGetThroughputRequest) (int, string, error) {
	if request == nil {
		request = NewGetRemoteCacheGetThroughputRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheGetThroughput")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheGetThroughputResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRemoteCacheIOPSSendRequest() (request *GetRemoteCacheIOPSSendRequest) {
	request = &GetRemoteCacheIOPSSendRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheIOPSSend")
	return
}

func NewGetRemoteCacheIOPSSendResponse() (response *GetRemoteCacheIOPSSendResponse) {
	response = &GetRemoteCacheIOPSSendResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRemoteCacheIOPSSend(request *GetRemoteCacheIOPSSendRequest) string {
	return c.GetRemoteCacheIOPSSendWithContext(context.Background(), request)
}

func (c *Client) GetRemoteCacheIOPSSendSend(request *GetRemoteCacheIOPSSendRequest) (*GetRemoteCacheIOPSSendResponse, error) {
	statusCode, msg, err := c.GetRemoteCacheIOPSSendWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRemoteCacheIOPSSendResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRemoteCacheIOPSSendWithContext(ctx context.Context, request *GetRemoteCacheIOPSSendRequest) string {
	if request == nil {
		request = NewGetRemoteCacheIOPSSendRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheIOPSSend")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheIOPSSendResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRemoteCacheIOPSSendWithContextV2(ctx context.Context, request *GetRemoteCacheIOPSSendRequest) (int, string, error) {
	if request == nil {
		request = NewGetRemoteCacheIOPSSendRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheIOPSSend")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheIOPSSendResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetRemoteCacheIOPSGetRequest() (request *GetRemoteCacheIOPSGetRequest) {
	request = &GetRemoteCacheIOPSGetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheIOPSGet")
	return
}

func NewGetRemoteCacheIOPSGetResponse() (response *GetRemoteCacheIOPSGetResponse) {
	response = &GetRemoteCacheIOPSGetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetRemoteCacheIOPSGet(request *GetRemoteCacheIOPSGetRequest) string {
	return c.GetRemoteCacheIOPSGetWithContext(context.Background(), request)
}

func (c *Client) GetRemoteCacheIOPSGetSend(request *GetRemoteCacheIOPSGetRequest) (*GetRemoteCacheIOPSGetResponse, error) {
	statusCode, msg, err := c.GetRemoteCacheIOPSGetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetRemoteCacheIOPSGetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetRemoteCacheIOPSGetWithContext(ctx context.Context, request *GetRemoteCacheIOPSGetRequest) string {
	if request == nil {
		request = NewGetRemoteCacheIOPSGetRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheIOPSGet")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheIOPSGetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetRemoteCacheIOPSGetWithContextV2(ctx context.Context, request *GetRemoteCacheIOPSGetRequest) (int, string, error) {
	if request == nil {
		request = NewGetRemoteCacheIOPSGetRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "GetRemoteCacheIOPSGet")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetRemoteCacheIOPSGetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowStrategySubscribeFailedRequest() (request *DescribeDataFlowStrategySubscribeFailedRequest) {
	request = &DescribeDataFlowStrategySubscribeFailedRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategySubscribeFailed")
	return
}

func NewDescribeDataFlowStrategySubscribeFailedResponse() (response *DescribeDataFlowStrategySubscribeFailedResponse) {
	response = &DescribeDataFlowStrategySubscribeFailedResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowStrategySubscribeFailed(request *DescribeDataFlowStrategySubscribeFailedRequest) string {
	return c.DescribeDataFlowStrategySubscribeFailedWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowStrategySubscribeFailedSend(request *DescribeDataFlowStrategySubscribeFailedRequest) (*DescribeDataFlowStrategySubscribeFailedResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowStrategySubscribeFailedWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataFlowStrategySubscribeFailedResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowStrategySubscribeFailedWithContext(ctx context.Context, request *DescribeDataFlowStrategySubscribeFailedRequest) string {
	if request == nil {
		request = NewDescribeDataFlowStrategySubscribeFailedRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategySubscribeFailed")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowStrategySubscribeFailedResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowStrategySubscribeFailedWithContextV2(ctx context.Context, request *DescribeDataFlowStrategySubscribeFailedRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowStrategySubscribeFailedRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowStrategySubscribeFailed")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDataFlowStrategySubscribeFailedResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewManageMigrateTaskRequest() (request *ManageMigrateTaskRequest) {
	request = &ManageMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "ManageMigrateTask")
	return
}

func NewManageMigrateTaskResponse() (response *ManageMigrateTaskResponse) {
	response = &ManageMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ManageMigrateTask(request *ManageMigrateTaskRequest) string {
	return c.ManageMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) ManageMigrateTaskSend(request *ManageMigrateTaskRequest) (*ManageMigrateTaskResponse, error) {
	statusCode, msg, err := c.ManageMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ManageMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ManageMigrateTaskWithContext(ctx context.Context, request *ManageMigrateTaskRequest) string {
	if request == nil {
		request = NewManageMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ManageMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewManageMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ManageMigrateTaskWithContextV2(ctx context.Context, request *ManageMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewManageMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "ManageMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewManageMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMigrateTasksRequest() (request *DescribeMigrateTasksRequest) {
	request = &DescribeMigrateTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeMigrateTasks")
	return
}

func NewDescribeMigrateTasksResponse() (response *DescribeMigrateTasksResponse) {
	response = &DescribeMigrateTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMigrateTasks(request *DescribeMigrateTasksRequest) string {
	return c.DescribeMigrateTasksWithContext(context.Background(), request)
}

func (c *Client) DescribeMigrateTasksSend(request *DescribeMigrateTasksRequest) (*DescribeMigrateTasksResponse, error) {
	statusCode, msg, err := c.DescribeMigrateTasksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeMigrateTasksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMigrateTasksWithContext(ctx context.Context, request *DescribeMigrateTasksRequest) string {
	if request == nil {
		request = NewDescribeMigrateTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeMigrateTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMigrateTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMigrateTasksWithContextV2(ctx context.Context, request *DescribeMigrateTasksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMigrateTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeMigrateTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMigrateTasksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMigrateTaskRequest() (request *CreateMigrateTaskRequest) {
	request = &CreateMigrateTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateMigrateTask")
	return
}

func NewCreateMigrateTaskResponse() (response *CreateMigrateTaskResponse) {
	response = &CreateMigrateTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMigrateTask(request *CreateMigrateTaskRequest) string {
	return c.CreateMigrateTaskWithContext(context.Background(), request)
}

func (c *Client) CreateMigrateTaskSend(request *CreateMigrateTaskRequest) (*CreateMigrateTaskResponse, error) {
	statusCode, msg, err := c.CreateMigrateTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateMigrateTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMigrateTaskWithContext(ctx context.Context, request *CreateMigrateTaskRequest) string {
	if request == nil {
		request = NewCreateMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMigrateTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMigrateTaskWithContextV2(ctx context.Context, request *CreateMigrateTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMigrateTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateMigrateTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMigrateTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMigrateRuleRequest() (request *DeleteMigrateRuleRequest) {
	request = &DeleteMigrateRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteMigrateRule")
	return
}

func NewDeleteMigrateRuleResponse() (response *DeleteMigrateRuleResponse) {
	response = &DeleteMigrateRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMigrateRule(request *DeleteMigrateRuleRequest) string {
	return c.DeleteMigrateRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteMigrateRuleSend(request *DeleteMigrateRuleRequest) (*DeleteMigrateRuleResponse, error) {
	statusCode, msg, err := c.DeleteMigrateRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteMigrateRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMigrateRuleWithContext(ctx context.Context, request *DeleteMigrateRuleRequest) string {
	if request == nil {
		request = NewDeleteMigrateRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteMigrateRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMigrateRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMigrateRuleWithContextV2(ctx context.Context, request *DeleteMigrateRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMigrateRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteMigrateRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMigrateRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMigrateRulesRequest() (request *DescribeMigrateRulesRequest) {
	request = &DescribeMigrateRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeMigrateRules")
	return
}

func NewDescribeMigrateRulesResponse() (response *DescribeMigrateRulesResponse) {
	response = &DescribeMigrateRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMigrateRules(request *DescribeMigrateRulesRequest) string {
	return c.DescribeMigrateRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeMigrateRulesSend(request *DescribeMigrateRulesRequest) (*DescribeMigrateRulesResponse, error) {
	statusCode, msg, err := c.DescribeMigrateRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeMigrateRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMigrateRulesWithContext(ctx context.Context, request *DescribeMigrateRulesRequest) string {
	if request == nil {
		request = NewDescribeMigrateRulesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeMigrateRules")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMigrateRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMigrateRulesWithContextV2(ctx context.Context, request *DescribeMigrateRulesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMigrateRulesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeMigrateRules")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMigrateRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMigrateRuleRequest() (request *CreateMigrateRuleRequest) {
	request = &CreateMigrateRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateMigrateRule")
	return
}

func NewCreateMigrateRuleResponse() (response *CreateMigrateRuleResponse) {
	response = &CreateMigrateRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMigrateRule(request *CreateMigrateRuleRequest) string {
	return c.CreateMigrateRuleWithContext(context.Background(), request)
}

func (c *Client) CreateMigrateRuleSend(request *CreateMigrateRuleRequest) (*CreateMigrateRuleResponse, error) {
	statusCode, msg, err := c.CreateMigrateRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateMigrateRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMigrateRuleWithContext(ctx context.Context, request *CreateMigrateRuleRequest) string {
	if request == nil {
		request = NewCreateMigrateRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateMigrateRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMigrateRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMigrateRuleWithContextV2(ctx context.Context, request *CreateMigrateRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMigrateRuleRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateMigrateRule")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMigrateRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
