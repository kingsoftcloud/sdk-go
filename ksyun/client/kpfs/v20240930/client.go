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
func NewDeleteDataFlowRequest() (request *DeleteDataFlowRequest) {
	request = &DeleteDataFlowRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlow")
	return
}

func NewDeleteDataFlowResponse() (response *DeleteDataFlowResponse) {
	response = &DeleteDataFlowResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDataFlow(request *DeleteDataFlowRequest) string {
	return c.DeleteDataFlowWithContext(context.Background(), request)
}

func (c *Client) DeleteDataFlowSend(request *DeleteDataFlowRequest) (*DeleteDataFlowResponse, error) {
	statusCode, msg, err := c.DeleteDataFlowWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDataFlowResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDataFlowWithContext(ctx context.Context, request *DeleteDataFlowRequest) string {
	if request == nil {
		request = NewDeleteDataFlowRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlow")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataFlowResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDataFlowWithContextV2(ctx context.Context, request *DeleteDataFlowRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDataFlowRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DeleteDataFlow")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDataFlowResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowTasksRequest() (request *DescribeDataFlowTasksRequest) {
	request = &DescribeDataFlowTasksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTasks")
	return
}

func NewDescribeDataFlowTasksResponse() (response *DescribeDataFlowTasksResponse) {
	response = &DescribeDataFlowTasksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlowTasks(request *DescribeDataFlowTasksRequest) string {
	return c.DescribeDataFlowTasksWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowTasksSend(request *DescribeDataFlowTasksRequest) (*DescribeDataFlowTasksResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowTasksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataFlowTasksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowTasksWithContext(ctx context.Context, request *DescribeDataFlowTasksRequest) string {
	if request == nil {
		request = NewDescribeDataFlowTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowTasksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowTasksWithContextV2(ctx context.Context, request *DescribeDataFlowTasksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowTasksRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlowTasks")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowTasksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDataFlowsRequest() (request *DescribeDataFlowsRequest) {
	request = &DescribeDataFlowsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlows")
	return
}

func NewDescribeDataFlowsResponse() (response *DescribeDataFlowsResponse) {
	response = &DescribeDataFlowsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataFlows(request *DescribeDataFlowsRequest) string {
	return c.DescribeDataFlowsWithContext(context.Background(), request)
}

func (c *Client) DescribeDataFlowsSend(request *DescribeDataFlowsRequest) (*DescribeDataFlowsResponse, error) {
	statusCode, msg, err := c.DescribeDataFlowsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDataFlowsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDataFlowsWithContext(ctx context.Context, request *DescribeDataFlowsRequest) string {
	if request == nil {
		request = NewDescribeDataFlowsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlows")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDataFlowsWithContextV2(ctx context.Context, request *DescribeDataFlowsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDataFlowsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "DescribeDataFlows")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDataFlowsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDataFlowTaskRequest() (request *CreateDataFlowTaskRequest) {
	request = &CreateDataFlowTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowTask")
	return
}

func NewCreateDataFlowTaskResponse() (response *CreateDataFlowTaskResponse) {
	response = &CreateDataFlowTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataFlowTask(request *CreateDataFlowTaskRequest) string {
	return c.CreateDataFlowTaskWithContext(context.Background(), request)
}

func (c *Client) CreateDataFlowTaskSend(request *CreateDataFlowTaskRequest) (*CreateDataFlowTaskResponse, error) {
	statusCode, msg, err := c.CreateDataFlowTaskWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDataFlowTaskResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDataFlowTaskWithContext(ctx context.Context, request *CreateDataFlowTaskRequest) string {
	if request == nil {
		request = NewCreateDataFlowTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataFlowTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDataFlowTaskWithContextV2(ctx context.Context, request *CreateDataFlowTaskRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDataFlowTaskRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlowTask")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataFlowTaskResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDataFlowRequest() (request *CreateDataFlowRequest) {
	request = &CreateDataFlowRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlow")
	return
}

func NewCreateDataFlowResponse() (response *CreateDataFlowResponse) {
	response = &CreateDataFlowResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDataFlow(request *CreateDataFlowRequest) string {
	return c.CreateDataFlowWithContext(context.Background(), request)
}

func (c *Client) CreateDataFlowSend(request *CreateDataFlowRequest) (*CreateDataFlowResponse, error) {
	statusCode, msg, err := c.CreateDataFlowWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateDataFlowResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDataFlowWithContext(ctx context.Context, request *CreateDataFlowRequest) string {
	if request == nil {
		request = NewCreateDataFlowRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlow")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataFlowResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDataFlowWithContextV2(ctx context.Context, request *CreateDataFlowRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDataFlowRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kpfs", APIVersion, "CreateDataFlow")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateDataFlowResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
