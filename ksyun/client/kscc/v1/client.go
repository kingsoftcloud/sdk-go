package v1

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "V1"

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

func NewUpdateUserQuotaRequest() (request *UpdateUserQuotaRequest) {
	request = &UpdateUserQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "UpdateUserQuota")
	return
}

func NewUpdateUserQuotaResponse() (response *UpdateUserQuotaResponse) {
	response = &UpdateUserQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateUserQuota(request *UpdateUserQuotaRequest) string {
	return c.UpdateUserQuotaWithContext(context.Background(), request)
}

func (c *Client) UpdateUserQuotaSend(request *UpdateUserQuotaRequest) (*UpdateUserQuotaResponse, error) {
	statusCode, msg, err := c.UpdateUserQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateUserQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateUserQuotaWithContext(ctx context.Context, request *UpdateUserQuotaRequest) string {
	if request == nil {
		request = NewUpdateUserQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateUserQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateUserQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateUserQuotaWithContextV2(ctx context.Context, request *UpdateUserQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateUserQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateUserQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateUserQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeUserCostSummaryRequest() (request *DescribeUserCostSummaryRequest) {
	request = &DescribeUserCostSummaryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserCostSummary")
	return
}

func NewDescribeUserCostSummaryResponse() (response *DescribeUserCostSummaryResponse) {
	response = &DescribeUserCostSummaryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeUserCostSummary(request *DescribeUserCostSummaryRequest) string {
	return c.DescribeUserCostSummaryWithContext(context.Background(), request)
}

func (c *Client) DescribeUserCostSummarySend(request *DescribeUserCostSummaryRequest) (*DescribeUserCostSummaryResponse, error) {
	statusCode, msg, err := c.DescribeUserCostSummaryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeUserCostSummaryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeUserCostSummaryWithContext(ctx context.Context, request *DescribeUserCostSummaryRequest) string {
	if request == nil {
		request = NewDescribeUserCostSummaryRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserCostSummary")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserCostSummaryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeUserCostSummaryWithContextV2(ctx context.Context, request *DescribeUserCostSummaryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeUserCostSummaryRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserCostSummary")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserCostSummaryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAiLogDetailByIdsRequest() (request *DescribeAiLogDetailByIdsRequest) {
	request = &DescribeAiLogDetailByIdsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeAiLogDetailByIds")
	return
}

func NewDescribeAiLogDetailByIdsResponse() (response *DescribeAiLogDetailByIdsResponse) {
	response = &DescribeAiLogDetailByIdsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAiLogDetailByIds(request *DescribeAiLogDetailByIdsRequest) string {
	return c.DescribeAiLogDetailByIdsWithContext(context.Background(), request)
}

func (c *Client) DescribeAiLogDetailByIdsSend(request *DescribeAiLogDetailByIdsRequest) (*DescribeAiLogDetailByIdsResponse, error) {
	statusCode, msg, err := c.DescribeAiLogDetailByIdsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAiLogDetailByIdsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAiLogDetailByIdsWithContext(ctx context.Context, request *DescribeAiLogDetailByIdsRequest) string {
	if request == nil {
		request = NewDescribeAiLogDetailByIdsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeAiLogDetailByIds")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAiLogDetailByIdsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAiLogDetailByIdsWithContextV2(ctx context.Context, request *DescribeAiLogDetailByIdsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAiLogDetailByIdsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeAiLogDetailByIds")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAiLogDetailByIdsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAiLogDetailRequest() (request *DescribeAiLogDetailRequest) {
	request = &DescribeAiLogDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeAiLogDetail")
	return
}

func NewDescribeAiLogDetailResponse() (response *DescribeAiLogDetailResponse) {
	response = &DescribeAiLogDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAiLogDetail(request *DescribeAiLogDetailRequest) string {
	return c.DescribeAiLogDetailWithContext(context.Background(), request)
}

func (c *Client) DescribeAiLogDetailSend(request *DescribeAiLogDetailRequest) (*DescribeAiLogDetailResponse, error) {
	statusCode, msg, err := c.DescribeAiLogDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAiLogDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAiLogDetailWithContext(ctx context.Context, request *DescribeAiLogDetailRequest) string {
	if request == nil {
		request = NewDescribeAiLogDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeAiLogDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAiLogDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAiLogDetailWithContextV2(ctx context.Context, request *DescribeAiLogDetailRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAiLogDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeAiLogDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAiLogDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeOrganizationTreeRequest() (request *DescribeOrganizationTreeRequest) {
	request = &DescribeOrganizationTreeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeOrganizationTree")
	return
}

func NewDescribeOrganizationTreeResponse() (response *DescribeOrganizationTreeResponse) {
	response = &DescribeOrganizationTreeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeOrganizationTree(request *DescribeOrganizationTreeRequest) string {
	return c.DescribeOrganizationTreeWithContext(context.Background(), request)
}

func (c *Client) DescribeOrganizationTreeSend(request *DescribeOrganizationTreeRequest) (*DescribeOrganizationTreeResponse, error) {
	statusCode, msg, err := c.DescribeOrganizationTreeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeOrganizationTreeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeOrganizationTreeWithContext(ctx context.Context, request *DescribeOrganizationTreeRequest) string {
	if request == nil {
		request = NewDescribeOrganizationTreeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeOrganizationTree")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeOrganizationTreeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeOrganizationTreeWithContextV2(ctx context.Context, request *DescribeOrganizationTreeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeOrganizationTreeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeOrganizationTree")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeOrganizationTreeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModelMetricsRequest() (request *DescribeModelMetricsRequest) {
	request = &DescribeModelMetricsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeModelMetrics")
	return
}

func NewDescribeModelMetricsResponse() (response *DescribeModelMetricsResponse) {
	response = &DescribeModelMetricsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModelMetrics(request *DescribeModelMetricsRequest) string {
	return c.DescribeModelMetricsWithContext(context.Background(), request)
}

func (c *Client) DescribeModelMetricsSend(request *DescribeModelMetricsRequest) (*DescribeModelMetricsResponse, error) {
	statusCode, msg, err := c.DescribeModelMetricsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeModelMetricsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModelMetricsWithContext(ctx context.Context, request *DescribeModelMetricsRequest) string {
	if request == nil {
		request = NewDescribeModelMetricsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeModelMetrics")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelMetricsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModelMetricsWithContextV2(ctx context.Context, request *DescribeModelMetricsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModelMetricsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeModelMetrics")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelMetricsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeUserTokenUsageRequest() (request *DescribeUserTokenUsageRequest) {
	request = &DescribeUserTokenUsageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserTokenUsage")
	return
}

func NewDescribeUserTokenUsageResponse() (response *DescribeUserTokenUsageResponse) {
	response = &DescribeUserTokenUsageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeUserTokenUsage(request *DescribeUserTokenUsageRequest) string {
	return c.DescribeUserTokenUsageWithContext(context.Background(), request)
}

func (c *Client) DescribeUserTokenUsageSend(request *DescribeUserTokenUsageRequest) (*DescribeUserTokenUsageResponse, error) {
	statusCode, msg, err := c.DescribeUserTokenUsageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeUserTokenUsageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeUserTokenUsageWithContext(ctx context.Context, request *DescribeUserTokenUsageRequest) string {
	if request == nil {
		request = NewDescribeUserTokenUsageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserTokenUsage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserTokenUsageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeUserTokenUsageWithContextV2(ctx context.Context, request *DescribeUserTokenUsageRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeUserTokenUsageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserTokenUsage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserTokenUsageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeUserQuotaListRequest() (request *DescribeUserQuotaListRequest) {
	request = &DescribeUserQuotaListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserQuotaList")
	return
}

func NewDescribeUserQuotaListResponse() (response *DescribeUserQuotaListResponse) {
	response = &DescribeUserQuotaListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeUserQuotaList(request *DescribeUserQuotaListRequest) string {
	return c.DescribeUserQuotaListWithContext(context.Background(), request)
}

func (c *Client) DescribeUserQuotaListSend(request *DescribeUserQuotaListRequest) (*DescribeUserQuotaListResponse, error) {
	statusCode, msg, err := c.DescribeUserQuotaListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeUserQuotaListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeUserQuotaListWithContext(ctx context.Context, request *DescribeUserQuotaListRequest) string {
	if request == nil {
		request = NewDescribeUserQuotaListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserQuotaList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserQuotaListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeUserQuotaListWithContextV2(ctx context.Context, request *DescribeUserQuotaListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeUserQuotaListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeUserQuotaList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUserQuotaListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeQuotaGlobalConfigRequest() (request *DescribeQuotaGlobalConfigRequest) {
	request = &DescribeQuotaGlobalConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeQuotaGlobalConfig")
	return
}

func NewDescribeQuotaGlobalConfigResponse() (response *DescribeQuotaGlobalConfigResponse) {
	response = &DescribeQuotaGlobalConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeQuotaGlobalConfig(request *DescribeQuotaGlobalConfigRequest) string {
	return c.DescribeQuotaGlobalConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeQuotaGlobalConfigSend(request *DescribeQuotaGlobalConfigRequest) (*DescribeQuotaGlobalConfigResponse, error) {
	statusCode, msg, err := c.DescribeQuotaGlobalConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeQuotaGlobalConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeQuotaGlobalConfigWithContext(ctx context.Context, request *DescribeQuotaGlobalConfigRequest) string {
	if request == nil {
		request = NewDescribeQuotaGlobalConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeQuotaGlobalConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeQuotaGlobalConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeQuotaGlobalConfigWithContextV2(ctx context.Context, request *DescribeQuotaGlobalConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeQuotaGlobalConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeQuotaGlobalConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeQuotaGlobalConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateQuotaGlobalConfigRequest() (request *UpdateQuotaGlobalConfigRequest) {
	request = &UpdateQuotaGlobalConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "UpdateQuotaGlobalConfig")
	return
}

func NewUpdateQuotaGlobalConfigResponse() (response *UpdateQuotaGlobalConfigResponse) {
	response = &UpdateQuotaGlobalConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateQuotaGlobalConfig(request *UpdateQuotaGlobalConfigRequest) string {
	return c.UpdateQuotaGlobalConfigWithContext(context.Background(), request)
}

func (c *Client) UpdateQuotaGlobalConfigSend(request *UpdateQuotaGlobalConfigRequest) (*UpdateQuotaGlobalConfigResponse, error) {
	statusCode, msg, err := c.UpdateQuotaGlobalConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateQuotaGlobalConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateQuotaGlobalConfigWithContext(ctx context.Context, request *UpdateQuotaGlobalConfigRequest) string {
	if request == nil {
		request = NewUpdateQuotaGlobalConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateQuotaGlobalConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateQuotaGlobalConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateQuotaGlobalConfigWithContextV2(ctx context.Context, request *UpdateQuotaGlobalConfigRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateQuotaGlobalConfigRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateQuotaGlobalConfig")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateQuotaGlobalConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAccountQuotaRequest() (request *DescribeAccountQuotaRequest) {
	request = &DescribeAccountQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeAccountQuota")
	return
}

func NewDescribeAccountQuotaResponse() (response *DescribeAccountQuotaResponse) {
	response = &DescribeAccountQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccountQuota(request *DescribeAccountQuotaRequest) string {
	return c.DescribeAccountQuotaWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountQuotaSend(request *DescribeAccountQuotaRequest) (*DescribeAccountQuotaResponse, error) {
	statusCode, msg, err := c.DescribeAccountQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAccountQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAccountQuotaWithContext(ctx context.Context, request *DescribeAccountQuotaRequest) string {
	if request == nil {
		request = NewDescribeAccountQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeAccountQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAccountQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAccountQuotaWithContextV2(ctx context.Context, request *DescribeAccountQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAccountQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeAccountQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAccountQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateAccountQuotaRequest() (request *UpdateAccountQuotaRequest) {
	request = &UpdateAccountQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "UpdateAccountQuota")
	return
}

func NewUpdateAccountQuotaResponse() (response *UpdateAccountQuotaResponse) {
	response = &UpdateAccountQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateAccountQuota(request *UpdateAccountQuotaRequest) string {
	return c.UpdateAccountQuotaWithContext(context.Background(), request)
}

func (c *Client) UpdateAccountQuotaSend(request *UpdateAccountQuotaRequest) (*UpdateAccountQuotaResponse, error) {
	statusCode, msg, err := c.UpdateAccountQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateAccountQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateAccountQuotaWithContext(ctx context.Context, request *UpdateAccountQuotaRequest) string {
	if request == nil {
		request = NewUpdateAccountQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateAccountQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateAccountQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateAccountQuotaWithContextV2(ctx context.Context, request *UpdateAccountQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateAccountQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateAccountQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateAccountQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAccountQuotaRequest() (request *DeleteAccountQuotaRequest) {
	request = &DeleteAccountQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DeleteAccountQuota")
	return
}

func NewDeleteAccountQuotaResponse() (response *DeleteAccountQuotaResponse) {
	response = &DeleteAccountQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAccountQuota(request *DeleteAccountQuotaRequest) string {
	return c.DeleteAccountQuotaWithContext(context.Background(), request)
}

func (c *Client) DeleteAccountQuotaSend(request *DeleteAccountQuotaRequest) (*DeleteAccountQuotaResponse, error) {
	statusCode, msg, err := c.DeleteAccountQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAccountQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAccountQuotaWithContext(ctx context.Context, request *DeleteAccountQuotaRequest) string {
	if request == nil {
		request = NewDeleteAccountQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteAccountQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteAccountQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAccountQuotaWithContextV2(ctx context.Context, request *DeleteAccountQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAccountQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteAccountQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteAccountQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDefaultMemberQuotaRequest() (request *UpdateDefaultMemberQuotaRequest) {
	request = &UpdateDefaultMemberQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "UpdateDefaultMemberQuota")
	return
}

func NewUpdateDefaultMemberQuotaResponse() (response *UpdateDefaultMemberQuotaResponse) {
	response = &UpdateDefaultMemberQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDefaultMemberQuota(request *UpdateDefaultMemberQuotaRequest) string {
	return c.UpdateDefaultMemberQuotaWithContext(context.Background(), request)
}

func (c *Client) UpdateDefaultMemberQuotaSend(request *UpdateDefaultMemberQuotaRequest) (*UpdateDefaultMemberQuotaResponse, error) {
	statusCode, msg, err := c.UpdateDefaultMemberQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDefaultMemberQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDefaultMemberQuotaWithContext(ctx context.Context, request *UpdateDefaultMemberQuotaRequest) string {
	if request == nil {
		request = NewUpdateDefaultMemberQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateDefaultMemberQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDefaultMemberQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDefaultMemberQuotaWithContextV2(ctx context.Context, request *UpdateDefaultMemberQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDefaultMemberQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateDefaultMemberQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDefaultMemberQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDefaultMemberQuotaRequest() (request *DeleteDefaultMemberQuotaRequest) {
	request = &DeleteDefaultMemberQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DeleteDefaultMemberQuota")
	return
}

func NewDeleteDefaultMemberQuotaResponse() (response *DeleteDefaultMemberQuotaResponse) {
	response = &DeleteDefaultMemberQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDefaultMemberQuota(request *DeleteDefaultMemberQuotaRequest) string {
	return c.DeleteDefaultMemberQuotaWithContext(context.Background(), request)
}

func (c *Client) DeleteDefaultMemberQuotaSend(request *DeleteDefaultMemberQuotaRequest) (*DeleteDefaultMemberQuotaResponse, error) {
	statusCode, msg, err := c.DeleteDefaultMemberQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDefaultMemberQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDefaultMemberQuotaWithContext(ctx context.Context, request *DeleteDefaultMemberQuotaRequest) string {
	if request == nil {
		request = NewDeleteDefaultMemberQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteDefaultMemberQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDefaultMemberQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDefaultMemberQuotaWithContextV2(ctx context.Context, request *DeleteDefaultMemberQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDefaultMemberQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteDefaultMemberQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDefaultMemberQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDeptQuotaListRequest() (request *DescribeDeptQuotaListRequest) {
	request = &DescribeDeptQuotaListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DescribeDeptQuotaList")
	return
}

func NewDescribeDeptQuotaListResponse() (response *DescribeDeptQuotaListResponse) {
	response = &DescribeDeptQuotaListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDeptQuotaList(request *DescribeDeptQuotaListRequest) string {
	return c.DescribeDeptQuotaListWithContext(context.Background(), request)
}

func (c *Client) DescribeDeptQuotaListSend(request *DescribeDeptQuotaListRequest) (*DescribeDeptQuotaListResponse, error) {
	statusCode, msg, err := c.DescribeDeptQuotaListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDeptQuotaListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDeptQuotaListWithContext(ctx context.Context, request *DescribeDeptQuotaListRequest) string {
	if request == nil {
		request = NewDescribeDeptQuotaListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeDeptQuotaList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDeptQuotaListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDeptQuotaListWithContextV2(ctx context.Context, request *DescribeDeptQuotaListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDeptQuotaListRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DescribeDeptQuotaList")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDeptQuotaListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDeptQuotaRequest() (request *UpdateDeptQuotaRequest) {
	request = &UpdateDeptQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "UpdateDeptQuota")
	return
}

func NewUpdateDeptQuotaResponse() (response *UpdateDeptQuotaResponse) {
	response = &UpdateDeptQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDeptQuota(request *UpdateDeptQuotaRequest) string {
	return c.UpdateDeptQuotaWithContext(context.Background(), request)
}

func (c *Client) UpdateDeptQuotaSend(request *UpdateDeptQuotaRequest) (*UpdateDeptQuotaResponse, error) {
	statusCode, msg, err := c.UpdateDeptQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateDeptQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDeptQuotaWithContext(ctx context.Context, request *UpdateDeptQuotaRequest) string {
	if request == nil {
		request = NewUpdateDeptQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateDeptQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDeptQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDeptQuotaWithContextV2(ctx context.Context, request *UpdateDeptQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDeptQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "UpdateDeptQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateDeptQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDeptQuotaRequest() (request *DeleteDeptQuotaRequest) {
	request = &DeleteDeptQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DeleteDeptQuota")
	return
}

func NewDeleteDeptQuotaResponse() (response *DeleteDeptQuotaResponse) {
	response = &DeleteDeptQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDeptQuota(request *DeleteDeptQuotaRequest) string {
	return c.DeleteDeptQuotaWithContext(context.Background(), request)
}

func (c *Client) DeleteDeptQuotaSend(request *DeleteDeptQuotaRequest) (*DeleteDeptQuotaResponse, error) {
	statusCode, msg, err := c.DeleteDeptQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDeptQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDeptQuotaWithContext(ctx context.Context, request *DeleteDeptQuotaRequest) string {
	if request == nil {
		request = NewDeleteDeptQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteDeptQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDeptQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDeptQuotaWithContextV2(ctx context.Context, request *DeleteDeptQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDeptQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteDeptQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDeptQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteUserQuotaRequest() (request *DeleteUserQuotaRequest) {
	request = &DeleteUserQuotaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kscc", APIVersion, "DeleteUserQuota")
	return
}

func NewDeleteUserQuotaResponse() (response *DeleteUserQuotaResponse) {
	response = &DeleteUserQuotaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteUserQuota(request *DeleteUserQuotaRequest) string {
	return c.DeleteUserQuotaWithContext(context.Background(), request)
}

func (c *Client) DeleteUserQuotaSend(request *DeleteUserQuotaRequest) (*DeleteUserQuotaResponse, error) {
	statusCode, msg, err := c.DeleteUserQuotaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteUserQuotaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteUserQuotaWithContext(ctx context.Context, request *DeleteUserQuotaRequest) string {
	if request == nil {
		request = NewDeleteUserQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteUserQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteUserQuotaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteUserQuotaWithContextV2(ctx context.Context, request *DeleteUserQuotaRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteUserQuotaRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kscc", APIVersion, "DeleteUserQuota")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteUserQuotaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
