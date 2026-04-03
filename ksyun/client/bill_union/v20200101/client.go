package v20200101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-01-01"

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

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
	request = &DescribeBillSummaryByPayModeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByPayMode")
	return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
	response = &DescribeBillSummaryByPayModeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) string {
	return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByPayModeSend(request *DescribeBillSummaryByPayModeRequest) (*DescribeBillSummaryByPayModeResponse, error) {
	statusCode, msg, err := c.DescribeBillSummaryByPayModeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBillSummaryByPayModeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) string {
	if request == nil {
		request = NewDescribeBillSummaryByPayModeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByPayMode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBillSummaryByPayModeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBillSummaryByPayModeWithContextV2(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBillSummaryByPayModeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByPayMode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBillSummaryByPayModeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
	request = &DescribeBillSummaryByProductRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProduct")
	return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
	response = &DescribeBillSummaryByProductResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) string {
	return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByProductSend(request *DescribeBillSummaryByProductRequest) (*DescribeBillSummaryByProductResponse, error) {
	statusCode, msg, err := c.DescribeBillSummaryByProductWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBillSummaryByProductResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) string {
	if request == nil {
		request = NewDescribeBillSummaryByProductRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProduct")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeBillSummaryByProductResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBillSummaryByProductWithContextV2(ctx context.Context, request *DescribeBillSummaryByProductRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBillSummaryByProductRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProduct")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeBillSummaryByProductResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
	request = &DescribeBillSummaryByProjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProject")
	return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
	response = &DescribeBillSummaryByProjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) string {
	return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByProjectSend(request *DescribeBillSummaryByProjectRequest) (*DescribeBillSummaryByProjectResponse, error) {
	statusCode, msg, err := c.DescribeBillSummaryByProjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeBillSummaryByProjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) string {
	if request == nil {
		request = NewDescribeBillSummaryByProjectRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProject")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBillSummaryByProjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBillSummaryByProjectWithContextV2(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBillSummaryByProjectRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProject")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBillSummaryByProjectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceSummaryBillsRequest() (request *DescribeInstanceSummaryBillsRequest) {
	request = &DescribeInstanceSummaryBillsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeInstanceSummaryBills")
	return
}

func NewDescribeInstanceSummaryBillsResponse() (response *DescribeInstanceSummaryBillsResponse) {
	response = &DescribeInstanceSummaryBillsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceSummaryBills(request *DescribeInstanceSummaryBillsRequest) string {
	return c.DescribeInstanceSummaryBillsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceSummaryBillsSend(request *DescribeInstanceSummaryBillsRequest) (*DescribeInstanceSummaryBillsResponse, error) {
	statusCode, msg, err := c.DescribeInstanceSummaryBillsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInstanceSummaryBillsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceSummaryBillsWithContext(ctx context.Context, request *DescribeInstanceSummaryBillsRequest) string {
	if request == nil {
		request = NewDescribeInstanceSummaryBillsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeInstanceSummaryBills")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceSummaryBillsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceSummaryBillsWithContextV2(ctx context.Context, request *DescribeInstanceSummaryBillsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceSummaryBillsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeInstanceSummaryBills")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceSummaryBillsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeProductCodeRequest() (request *DescribeProductCodeRequest) {
	request = &DescribeProductCodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeProductCode")
	return
}

func NewDescribeProductCodeResponse() (response *DescribeProductCodeResponse) {
	response = &DescribeProductCodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeProductCode(request *DescribeProductCodeRequest) string {
	return c.DescribeProductCodeWithContext(context.Background(), request)
}

func (c *Client) DescribeProductCodeSend(request *DescribeProductCodeRequest) (*DescribeProductCodeResponse, error) {
	statusCode, msg, err := c.DescribeProductCodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeProductCodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeProductCodeWithContext(ctx context.Context, request *DescribeProductCodeRequest) string {
	if request == nil {
		request = NewDescribeProductCodeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeProductCode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeProductCodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeProductCodeWithContextV2(ctx context.Context, request *DescribeProductCodeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeProductCodeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeProductCode")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeProductCodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSplitItemBillDetailsRequest() (request *DescribeSplitItemBillDetailsRequest) {
	request = &DescribeSplitItemBillDetailsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemBillDetails")
	return
}

func NewDescribeSplitItemBillDetailsResponse() (response *DescribeSplitItemBillDetailsResponse) {
	response = &DescribeSplitItemBillDetailsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSplitItemBillDetails(request *DescribeSplitItemBillDetailsRequest) string {
	return c.DescribeSplitItemBillDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeSplitItemBillDetailsSend(request *DescribeSplitItemBillDetailsRequest) (*DescribeSplitItemBillDetailsResponse, error) {
	statusCode, msg, err := c.DescribeSplitItemBillDetailsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSplitItemBillDetailsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSplitItemBillDetailsWithContext(ctx context.Context, request *DescribeSplitItemBillDetailsRequest) string {
	if request == nil {
		request = NewDescribeSplitItemBillDetailsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemBillDetails")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSplitItemBillDetailsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSplitItemBillDetailsWithContextV2(ctx context.Context, request *DescribeSplitItemBillDetailsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSplitItemBillDetailsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemBillDetails")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSplitItemBillDetailsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMiItemBillsRequest() (request *DescribeMiItemBillsRequest) {
	request = &DescribeMiItemBillsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeMiItemBills")
	return
}

func NewDescribeMiItemBillsResponse() (response *DescribeMiItemBillsResponse) {
	response = &DescribeMiItemBillsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMiItemBills(request *DescribeMiItemBillsRequest) string {
	return c.DescribeMiItemBillsWithContext(context.Background(), request)
}

func (c *Client) DescribeMiItemBillsSend(request *DescribeMiItemBillsRequest) (*DescribeMiItemBillsResponse, error) {
	statusCode, msg, err := c.DescribeMiItemBillsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeMiItemBillsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMiItemBillsWithContext(ctx context.Context, request *DescribeMiItemBillsRequest) string {
	if request == nil {
		request = NewDescribeMiItemBillsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeMiItemBills")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMiItemBillsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMiItemBillsWithContextV2(ctx context.Context, request *DescribeMiItemBillsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMiItemBillsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeMiItemBills")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMiItemBillsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSplitItemDayBillDetailsRequest() (request *DescribeSplitItemDayBillDetailsRequest) {
	request = &DescribeSplitItemDayBillDetailsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemDayBillDetails")
	return
}

func NewDescribeSplitItemDayBillDetailsResponse() (response *DescribeSplitItemDayBillDetailsResponse) {
	response = &DescribeSplitItemDayBillDetailsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSplitItemDayBillDetails(request *DescribeSplitItemDayBillDetailsRequest) string {
	return c.DescribeSplitItemDayBillDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeSplitItemDayBillDetailsSend(request *DescribeSplitItemDayBillDetailsRequest) (*DescribeSplitItemDayBillDetailsResponse, error) {
	statusCode, msg, err := c.DescribeSplitItemDayBillDetailsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeSplitItemDayBillDetailsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSplitItemDayBillDetailsWithContext(ctx context.Context, request *DescribeSplitItemDayBillDetailsRequest) string {
	if request == nil {
		request = NewDescribeSplitItemDayBillDetailsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemDayBillDetails")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSplitItemDayBillDetailsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSplitItemDayBillDetailsWithContextV2(ctx context.Context, request *DescribeSplitItemDayBillDetailsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSplitItemDayBillDetailsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemDayBillDetails")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSplitItemDayBillDetailsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListProductGroupsRequest() (request *ListProductGroupsRequest) {
	request = &ListProductGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill-union", APIVersion, "ListProductGroups")
	return
}

func NewListProductGroupsResponse() (response *ListProductGroupsResponse) {
	response = &ListProductGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListProductGroups(request *ListProductGroupsRequest) string {
	return c.ListProductGroupsWithContext(context.Background(), request)
}

func (c *Client) ListProductGroupsSend(request *ListProductGroupsRequest) (*ListProductGroupsResponse, error) {
	statusCode, msg, err := c.ListProductGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListProductGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListProductGroupsWithContext(ctx context.Context, request *ListProductGroupsRequest) string {
	if request == nil {
		request = NewListProductGroupsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "ListProductGroups")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListProductGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListProductGroupsWithContextV2(ctx context.Context, request *ListProductGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewListProductGroupsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("bill-union", APIVersion, "ListProductGroups")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListProductGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
