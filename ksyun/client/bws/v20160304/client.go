package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-03-04"

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

func NewCreateBandWidthShareRequest() (request *CreateBandWidthShareRequest) {
	request = &CreateBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "CreateBandWidthShare")
	return
}

func NewCreateBandWidthShareResponse() (response *CreateBandWidthShareResponse) {
	response = &CreateBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateBandWidthShare(request *CreateBandWidthShareRequest) string {
	return c.CreateBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) CreateBandWidthShareSend(request *CreateBandWidthShareRequest) (*CreateBandWidthShareResponse, error) {
	statusCode, msg, err := c.CreateBandWidthShareWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateBandWidthShareResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateBandWidthShareWithContext(ctx context.Context, request *CreateBandWidthShareRequest) string {
	if request == nil {
		request = NewCreateBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "CreateBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateBandWidthShareWithContextV2(ctx context.Context, request *CreateBandWidthShareRequest) (int, string, error) {
	if request == nil {
		request = NewCreateBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "CreateBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateBandWidthShareResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBandWidthSharesRequest() (request *DescribeBandWidthSharesRequest) {
	request = &DescribeBandWidthSharesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "DescribeBandWidthShares")
	return
}

func NewDescribeBandWidthSharesResponse() (response *DescribeBandWidthSharesResponse) {
	response = &DescribeBandWidthSharesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBandWidthShares(request *DescribeBandWidthSharesRequest) string {
	return c.DescribeBandWidthSharesWithContext(context.Background(), request)
}

func (c *Client) DescribeBandWidthSharesSend(request *DescribeBandWidthSharesRequest) (*DescribeBandWidthSharesResponse, error) {
	statusCode, msg, err := c.DescribeBandWidthSharesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBandWidthSharesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBandWidthSharesWithContext(ctx context.Context, request *DescribeBandWidthSharesRequest) string {
	if request == nil {
		request = NewDescribeBandWidthSharesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "DescribeBandWidthShares")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBandWidthSharesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBandWidthSharesWithContextV2(ctx context.Context, request *DescribeBandWidthSharesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBandWidthSharesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "DescribeBandWidthShares")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBandWidthSharesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateBandWidthShareRequest() (request *AssociateBandWidthShareRequest) {
	request = &AssociateBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "AssociateBandWidthShare")
	return
}

func NewAssociateBandWidthShareResponse() (response *AssociateBandWidthShareResponse) {
	response = &AssociateBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateBandWidthShare(request *AssociateBandWidthShareRequest) string {
	return c.AssociateBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) AssociateBandWidthShareSend(request *AssociateBandWidthShareRequest) (*AssociateBandWidthShareResponse, error) {
	statusCode, msg, err := c.AssociateBandWidthShareWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateBandWidthShareResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateBandWidthShareWithContext(ctx context.Context, request *AssociateBandWidthShareRequest) string {
	if request == nil {
		request = NewAssociateBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "AssociateBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateBandWidthShareWithContextV2(ctx context.Context, request *AssociateBandWidthShareRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "AssociateBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateBandWidthShareResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateBandWidthShareRequest() (request *DisassociateBandWidthShareRequest) {
	request = &DisassociateBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "DisassociateBandWidthShare")
	return
}

func NewDisassociateBandWidthShareResponse() (response *DisassociateBandWidthShareResponse) {
	response = &DisassociateBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateBandWidthShare(request *DisassociateBandWidthShareRequest) string {
	return c.DisassociateBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) DisassociateBandWidthShareSend(request *DisassociateBandWidthShareRequest) (*DisassociateBandWidthShareResponse, error) {
	statusCode, msg, err := c.DisassociateBandWidthShareWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisassociateBandWidthShareResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateBandWidthShareWithContext(ctx context.Context, request *DisassociateBandWidthShareRequest) string {
	if request == nil {
		request = NewDisassociateBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "DisassociateBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateBandWidthShareWithContextV2(ctx context.Context, request *DisassociateBandWidthShareRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "DisassociateBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateBandWidthShareResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyBandWidthShareRequest() (request *ModifyBandWidthShareRequest) {
	request = &ModifyBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "ModifyBandWidthShare")
	return
}

func NewModifyBandWidthShareResponse() (response *ModifyBandWidthShareResponse) {
	response = &ModifyBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBandWidthShare(request *ModifyBandWidthShareRequest) string {
	return c.ModifyBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) ModifyBandWidthShareSend(request *ModifyBandWidthShareRequest) (*ModifyBandWidthShareResponse, error) {
	statusCode, msg, err := c.ModifyBandWidthShareWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyBandWidthShareResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyBandWidthShareWithContext(ctx context.Context, request *ModifyBandWidthShareRequest) string {
	if request == nil {
		request = NewModifyBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "ModifyBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyBandWidthShareWithContextV2(ctx context.Context, request *ModifyBandWidthShareRequest) (int, string, error) {
	if request == nil {
		request = NewModifyBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "ModifyBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBandWidthShareResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBandWidthShareRequest() (request *DeleteBandWidthShareRequest) {
	request = &DeleteBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "DeleteBandWidthShare")
	return
}

func NewDeleteBandWidthShareResponse() (response *DeleteBandWidthShareResponse) {
	response = &DeleteBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBandWidthShare(request *DeleteBandWidthShareRequest) string {
	return c.DeleteBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) DeleteBandWidthShareSend(request *DeleteBandWidthShareRequest) (*DeleteBandWidthShareResponse, error) {
	statusCode, msg, err := c.DeleteBandWidthShareWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteBandWidthShareResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBandWidthShareWithContext(ctx context.Context, request *DeleteBandWidthShareRequest) string {
	if request == nil {
		request = NewDeleteBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "DeleteBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBandWidthShareWithContextV2(ctx context.Context, request *DeleteBandWidthShareRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBandWidthShareRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "DeleteBandWidthShare")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBandWidthShareResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryBwsTopEipMonitorRequest() (request *QueryBwsTopEipMonitorRequest) {
	request = &QueryBwsTopEipMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "QueryBwsTopEipMonitor")
	return
}

func NewQueryBwsTopEipMonitorResponse() (response *QueryBwsTopEipMonitorResponse) {
	response = &QueryBwsTopEipMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryBwsTopEipMonitor(request *QueryBwsTopEipMonitorRequest) string {
	return c.QueryBwsTopEipMonitorWithContext(context.Background(), request)
}

func (c *Client) QueryBwsTopEipMonitorSend(request *QueryBwsTopEipMonitorRequest) (*QueryBwsTopEipMonitorResponse, error) {
	statusCode, msg, err := c.QueryBwsTopEipMonitorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryBwsTopEipMonitorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryBwsTopEipMonitorWithContext(ctx context.Context, request *QueryBwsTopEipMonitorRequest) string {
	if request == nil {
		request = NewQueryBwsTopEipMonitorRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "QueryBwsTopEipMonitor")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryBwsTopEipMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryBwsTopEipMonitorWithContextV2(ctx context.Context, request *QueryBwsTopEipMonitorRequest) (int, string, error) {
	if request == nil {
		request = NewQueryBwsTopEipMonitorRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bws", APIVersion, "QueryBwsTopEipMonitor")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryBwsTopEipMonitorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
