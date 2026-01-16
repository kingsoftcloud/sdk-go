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

func NewDeleteBatchCfwAddrbookRequest() (request *DeleteBatchCfwAddrbookRequest) {
	request = &DeleteBatchCfwAddrbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteBatchCfwAddrbook")
	return
}

func NewDeleteBatchCfwAddrbookResponse() (response *DeleteBatchCfwAddrbookResponse) {
	response = &DeleteBatchCfwAddrbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBatchCfwAddrbook(request *DeleteBatchCfwAddrbookRequest) string {
	return c.DeleteBatchCfwAddrbookWithContext(context.Background(), request)
}

func (c *Client) DeleteBatchCfwAddrbookSend(request *DeleteBatchCfwAddrbookRequest) (*DeleteBatchCfwAddrbookResponse, error) {
	statusCode, msg, err := c.DeleteBatchCfwAddrbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteBatchCfwAddrbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBatchCfwAddrbookWithContext(ctx context.Context, request *DeleteBatchCfwAddrbookRequest) string {
	if request == nil {
		request = NewDeleteBatchCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteBatchCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchCfwAddrbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBatchCfwAddrbookWithContextV2(ctx context.Context, request *DeleteBatchCfwAddrbookRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBatchCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteBatchCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchCfwAddrbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteServiceGroupBatchRequest() (request *DeleteServiceGroupBatchRequest) {
	request = &DeleteServiceGroupBatchRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteServiceGroupBatch")
	return
}

func NewDeleteServiceGroupBatchResponse() (response *DeleteServiceGroupBatchResponse) {
	response = &DeleteServiceGroupBatchResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteServiceGroupBatch(request *DeleteServiceGroupBatchRequest) string {
	return c.DeleteServiceGroupBatchWithContext(context.Background(), request)
}

func (c *Client) DeleteServiceGroupBatchSend(request *DeleteServiceGroupBatchRequest) (*DeleteServiceGroupBatchResponse, error) {
	statusCode, msg, err := c.DeleteServiceGroupBatchWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteServiceGroupBatchResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteServiceGroupBatchWithContext(ctx context.Context, request *DeleteServiceGroupBatchRequest) string {
	if request == nil {
		request = NewDeleteServiceGroupBatchRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteServiceGroupBatch")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteServiceGroupBatchResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteServiceGroupBatchWithContextV2(ctx context.Context, request *DeleteServiceGroupBatchRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteServiceGroupBatchRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteServiceGroupBatch")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteServiceGroupBatchResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBatchHostbookRequest() (request *DeleteBatchHostbookRequest) {
	request = &DeleteBatchHostbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteBatchHostbook")
	return
}

func NewDeleteBatchHostbookResponse() (response *DeleteBatchHostbookResponse) {
	response = &DeleteBatchHostbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBatchHostbook(request *DeleteBatchHostbookRequest) string {
	return c.DeleteBatchHostbookWithContext(context.Background(), request)
}

func (c *Client) DeleteBatchHostbookSend(request *DeleteBatchHostbookRequest) (*DeleteBatchHostbookResponse, error) {
	statusCode, msg, err := c.DeleteBatchHostbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteBatchHostbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBatchHostbookWithContext(ctx context.Context, request *DeleteBatchHostbookRequest) string {
	if request == nil {
		request = NewDeleteBatchHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteBatchHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchHostbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBatchHostbookWithContextV2(ctx context.Context, request *DeleteBatchHostbookRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBatchHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteBatchHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchHostbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyHostbookRequest() (request *ModifyHostbookRequest) {
	request = &ModifyHostbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ModifyHostbook")
	return
}

func NewModifyHostbookResponse() (response *ModifyHostbookResponse) {
	response = &ModifyHostbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyHostbook(request *ModifyHostbookRequest) string {
	return c.ModifyHostbookWithContext(context.Background(), request)
}

func (c *Client) ModifyHostbookSend(request *ModifyHostbookRequest) (*ModifyHostbookResponse, error) {
	statusCode, msg, err := c.ModifyHostbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyHostbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyHostbookWithContext(ctx context.Context, request *ModifyHostbookRequest) string {
	if request == nil {
		request = NewModifyHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyHostbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyHostbookWithContextV2(ctx context.Context, request *ModifyHostbookRequest) (int, string, error) {
	if request == nil {
		request = NewModifyHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyHostbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateHostbookRequest() (request *CreateHostbookRequest) {
	request = &CreateHostbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "CreateHostbook")
	return
}

func NewCreateHostbookResponse() (response *CreateHostbookResponse) {
	response = &CreateHostbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateHostbook(request *CreateHostbookRequest) string {
	return c.CreateHostbookWithContext(context.Background(), request)
}

func (c *Client) CreateHostbookSend(request *CreateHostbookRequest) (*CreateHostbookResponse, error) {
	statusCode, msg, err := c.CreateHostbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateHostbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateHostbookWithContext(ctx context.Context, request *CreateHostbookRequest) string {
	if request == nil {
		request = NewCreateHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateHostbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateHostbookWithContextV2(ctx context.Context, request *CreateHostbookRequest) (int, string, error) {
	if request == nil {
		request = NewCreateHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateHostbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeHostbookRequest() (request *DescribeHostbookRequest) {
	request = &DescribeHostbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeHostbook")
	return
}

func NewDescribeHostbookResponse() (response *DescribeHostbookResponse) {
	response = &DescribeHostbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeHostbook(request *DescribeHostbookRequest) string {
	return c.DescribeHostbookWithContext(context.Background(), request)
}

func (c *Client) DescribeHostbookSend(request *DescribeHostbookRequest) (*DescribeHostbookResponse, error) {
	statusCode, msg, err := c.DescribeHostbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeHostbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeHostbookWithContext(ctx context.Context, request *DescribeHostbookRequest) string {
	if request == nil {
		request = NewDescribeHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeHostbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeHostbookWithContextV2(ctx context.Context, request *DescribeHostbookRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeHostbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeHostbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeHostbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryCfwInstanceDetailRequest() (request *QueryCfwInstanceDetailRequest) {
	request = &QueryCfwInstanceDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "QueryCfwInstanceDetail")
	return
}

func NewQueryCfwInstanceDetailResponse() (response *QueryCfwInstanceDetailResponse) {
	response = &QueryCfwInstanceDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryCfwInstanceDetail(request *QueryCfwInstanceDetailRequest) string {
	return c.QueryCfwInstanceDetailWithContext(context.Background(), request)
}

func (c *Client) QueryCfwInstanceDetailSend(request *QueryCfwInstanceDetailRequest) (*QueryCfwInstanceDetailResponse, error) {
	statusCode, msg, err := c.QueryCfwInstanceDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryCfwInstanceDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryCfwInstanceDetailWithContext(ctx context.Context, request *QueryCfwInstanceDetailRequest) string {
	if request == nil {
		request = NewQueryCfwInstanceDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "QueryCfwInstanceDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryCfwInstanceDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryCfwInstanceDetailWithContextV2(ctx context.Context, request *QueryCfwInstanceDetailRequest) (int, string, error) {
	if request == nil {
		request = NewQueryCfwInstanceDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "QueryCfwInstanceDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryCfwInstanceDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCloudFireWallInstanceRequest() (request *DeleteCloudFireWallInstanceRequest) {
	request = &DeleteCloudFireWallInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteCloudFireWallInstance")
	return
}

func NewDeleteCloudFireWallInstanceResponse() (response *DeleteCloudFireWallInstanceResponse) {
	response = &DeleteCloudFireWallInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCloudFireWallInstance(request *DeleteCloudFireWallInstanceRequest) string {
	return c.DeleteCloudFireWallInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteCloudFireWallInstanceSend(request *DeleteCloudFireWallInstanceRequest) (*DeleteCloudFireWallInstanceResponse, error) {
	statusCode, msg, err := c.DeleteCloudFireWallInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCloudFireWallInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCloudFireWallInstanceWithContext(ctx context.Context, request *DeleteCloudFireWallInstanceRequest) string {
	if request == nil {
		request = NewDeleteCloudFireWallInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCloudFireWallInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCloudFireWallInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCloudFireWallInstanceWithContextV2(ctx context.Context, request *DeleteCloudFireWallInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCloudFireWallInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCloudFireWallInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCloudFireWallInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryOverviewDetailRequest() (request *QueryOverviewDetailRequest) {
	request = &QueryOverviewDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "QueryOverviewDetail")
	return
}

func NewQueryOverviewDetailResponse() (response *QueryOverviewDetailResponse) {
	response = &QueryOverviewDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryOverviewDetail(request *QueryOverviewDetailRequest) string {
	return c.QueryOverviewDetailWithContext(context.Background(), request)
}

func (c *Client) QueryOverviewDetailSend(request *QueryOverviewDetailRequest) (*QueryOverviewDetailResponse, error) {
	statusCode, msg, err := c.QueryOverviewDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryOverviewDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryOverviewDetailWithContext(ctx context.Context, request *QueryOverviewDetailRequest) string {
	if request == nil {
		request = NewQueryOverviewDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "QueryOverviewDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryOverviewDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryOverviewDetailWithContextV2(ctx context.Context, request *QueryOverviewDetailRequest) (int, string, error) {
	if request == nil {
		request = NewQueryOverviewDetailRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "QueryOverviewDetail")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryOverviewDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTrafficLogRequest() (request *DescribeTrafficLogRequest) {
	request = &DescribeTrafficLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeTrafficLog")
	return
}

func NewDescribeTrafficLogResponse() (response *DescribeTrafficLogResponse) {
	response = &DescribeTrafficLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrafficLog(request *DescribeTrafficLogRequest) string {
	return c.DescribeTrafficLogWithContext(context.Background(), request)
}

func (c *Client) DescribeTrafficLogSend(request *DescribeTrafficLogRequest) (*DescribeTrafficLogResponse, error) {
	statusCode, msg, err := c.DescribeTrafficLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrafficLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrafficLogWithContext(ctx context.Context, request *DescribeTrafficLogRequest) string {
	if request == nil {
		request = NewDescribeTrafficLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeTrafficLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeTrafficLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrafficLogWithContextV2(ctx context.Context, request *DescribeTrafficLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrafficLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeTrafficLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeTrafficLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAttackLogRequest() (request *DescribeAttackLogRequest) {
	request = &DescribeAttackLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeAttackLog")
	return
}

func NewDescribeAttackLogResponse() (response *DescribeAttackLogResponse) {
	response = &DescribeAttackLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAttackLog(request *DescribeAttackLogRequest) string {
	return c.DescribeAttackLogWithContext(context.Background(), request)
}

func (c *Client) DescribeAttackLogSend(request *DescribeAttackLogRequest) (*DescribeAttackLogResponse, error) {
	statusCode, msg, err := c.DescribeAttackLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAttackLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAttackLogWithContext(ctx context.Context, request *DescribeAttackLogRequest) string {
	if request == nil {
		request = NewDescribeAttackLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeAttackLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAttackLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAttackLogWithContextV2(ctx context.Context, request *DescribeAttackLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAttackLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeAttackLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAttackLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAclLogRequest() (request *DescribeAclLogRequest) {
	request = &DescribeAclLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeAclLog")
	return
}

func NewDescribeAclLogResponse() (response *DescribeAclLogResponse) {
	response = &DescribeAclLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAclLog(request *DescribeAclLogRequest) string {
	return c.DescribeAclLogWithContext(context.Background(), request)
}

func (c *Client) DescribeAclLogSend(request *DescribeAclLogRequest) (*DescribeAclLogResponse, error) {
	statusCode, msg, err := c.DescribeAclLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAclLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAclLogWithContext(ctx context.Context, request *DescribeAclLogRequest) string {
	if request == nil {
		request = NewDescribeAclLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeAclLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAclLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAclLogWithContextV2(ctx context.Context, request *DescribeAclLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAclLogRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeAclLog")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAclLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCloudFireWallFeatureRequest() (request *ModifyCloudFireWallFeatureRequest) {
	request = &ModifyCloudFireWallFeatureRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ModifyCloudFireWallFeature")
	return
}

func NewModifyCloudFireWallFeatureResponse() (response *ModifyCloudFireWallFeatureResponse) {
	response = &ModifyCloudFireWallFeatureResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCloudFireWallFeature(request *ModifyCloudFireWallFeatureRequest) string {
	return c.ModifyCloudFireWallFeatureWithContext(context.Background(), request)
}

func (c *Client) ModifyCloudFireWallFeatureSend(request *ModifyCloudFireWallFeatureRequest) (*ModifyCloudFireWallFeatureResponse, error) {
	statusCode, msg, err := c.ModifyCloudFireWallFeatureWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCloudFireWallFeatureResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCloudFireWallFeatureWithContext(ctx context.Context, request *ModifyCloudFireWallFeatureRequest) string {
	if request == nil {
		request = NewModifyCloudFireWallFeatureRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCloudFireWallFeature")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCloudFireWallFeatureResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCloudFireWallFeatureWithContextV2(ctx context.Context, request *ModifyCloudFireWallFeatureRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCloudFireWallFeatureRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCloudFireWallFeature")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCloudFireWallFeatureResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCfwAddrbookRequest() (request *DescribeCfwAddrbookRequest) {
	request = &DescribeCfwAddrbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAddrbook")
	return
}

func NewDescribeCfwAddrbookResponse() (response *DescribeCfwAddrbookResponse) {
	response = &DescribeCfwAddrbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfwAddrbook(request *DescribeCfwAddrbookRequest) string {
	return c.DescribeCfwAddrbookWithContext(context.Background(), request)
}

func (c *Client) DescribeCfwAddrbookSend(request *DescribeCfwAddrbookRequest) (*DescribeCfwAddrbookResponse, error) {
	statusCode, msg, err := c.DescribeCfwAddrbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCfwAddrbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCfwAddrbookWithContext(ctx context.Context, request *DescribeCfwAddrbookRequest) string {
	if request == nil {
		request = NewDescribeCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCfwAddrbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCfwAddrbookWithContextV2(ctx context.Context, request *DescribeCfwAddrbookRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCfwAddrbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCfwAddrbookRequest() (request *DeleteCfwAddrbookRequest) {
	request = &DeleteCfwAddrbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwAddrbook")
	return
}

func NewDeleteCfwAddrbookResponse() (response *DeleteCfwAddrbookResponse) {
	response = &DeleteCfwAddrbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCfwAddrbook(request *DeleteCfwAddrbookRequest) string {
	return c.DeleteCfwAddrbookWithContext(context.Background(), request)
}

func (c *Client) DeleteCfwAddrbookSend(request *DeleteCfwAddrbookRequest) (*DeleteCfwAddrbookResponse, error) {
	statusCode, msg, err := c.DeleteCfwAddrbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCfwAddrbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCfwAddrbookWithContext(ctx context.Context, request *DeleteCfwAddrbookRequest) string {
	if request == nil {
		request = NewDeleteCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCfwAddrbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCfwAddrbookWithContextV2(ctx context.Context, request *DeleteCfwAddrbookRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCfwAddrbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCfwAddrbookRequest() (request *ModifyCfwAddrbookRequest) {
	request = &ModifyCfwAddrbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwAddrbook")
	return
}

func NewModifyCfwAddrbookResponse() (response *ModifyCfwAddrbookResponse) {
	response = &ModifyCfwAddrbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCfwAddrbook(request *ModifyCfwAddrbookRequest) string {
	return c.ModifyCfwAddrbookWithContext(context.Background(), request)
}

func (c *Client) ModifyCfwAddrbookSend(request *ModifyCfwAddrbookRequest) (*ModifyCfwAddrbookResponse, error) {
	statusCode, msg, err := c.ModifyCfwAddrbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCfwAddrbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCfwAddrbookWithContext(ctx context.Context, request *ModifyCfwAddrbookRequest) string {
	if request == nil {
		request = NewModifyCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwAddrbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCfwAddrbookWithContextV2(ctx context.Context, request *ModifyCfwAddrbookRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwAddrbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCfwAddrbookRequest() (request *CreateCfwAddrbookRequest) {
	request = &CreateCfwAddrbookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwAddrbook")
	return
}

func NewCreateCfwAddrbookResponse() (response *CreateCfwAddrbookResponse) {
	response = &CreateCfwAddrbookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCfwAddrbook(request *CreateCfwAddrbookRequest) string {
	return c.CreateCfwAddrbookWithContext(context.Background(), request)
}

func (c *Client) CreateCfwAddrbookSend(request *CreateCfwAddrbookRequest) (*CreateCfwAddrbookResponse, error) {
	statusCode, msg, err := c.CreateCfwAddrbookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCfwAddrbookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCfwAddrbookWithContext(ctx context.Context, request *CreateCfwAddrbookRequest) string {
	if request == nil {
		request = NewCreateCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCfwAddrbookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCfwAddrbookWithContextV2(ctx context.Context, request *CreateCfwAddrbookRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCfwAddrbookRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwAddrbook")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCfwAddrbookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAlterCfwAclStatusRequest() (request *AlterCfwAclStatusRequest) {
	request = &AlterCfwAclStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "AlterCfwAclStatus")
	return
}

func NewAlterCfwAclStatusResponse() (response *AlterCfwAclStatusResponse) {
	response = &AlterCfwAclStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AlterCfwAclStatus(request *AlterCfwAclStatusRequest) string {
	return c.AlterCfwAclStatusWithContext(context.Background(), request)
}

func (c *Client) AlterCfwAclStatusSend(request *AlterCfwAclStatusRequest) (*AlterCfwAclStatusResponse, error) {
	statusCode, msg, err := c.AlterCfwAclStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AlterCfwAclStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AlterCfwAclStatusWithContext(ctx context.Context, request *AlterCfwAclStatusRequest) string {
	if request == nil {
		request = NewAlterCfwAclStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "AlterCfwAclStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAlterCfwAclStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AlterCfwAclStatusWithContextV2(ctx context.Context, request *AlterCfwAclStatusRequest) (int, string, error) {
	if request == nil {
		request = NewAlterCfwAclStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "AlterCfwAclStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAlterCfwAclStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResetCfwAclHitCountRequest() (request *ResetCfwAclHitCountRequest) {
	request = &ResetCfwAclHitCountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ResetCfwAclHitCount")
	return
}

func NewResetCfwAclHitCountResponse() (response *ResetCfwAclHitCountResponse) {
	response = &ResetCfwAclHitCountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetCfwAclHitCount(request *ResetCfwAclHitCountRequest) string {
	return c.ResetCfwAclHitCountWithContext(context.Background(), request)
}

func (c *Client) ResetCfwAclHitCountSend(request *ResetCfwAclHitCountRequest) (*ResetCfwAclHitCountResponse, error) {
	statusCode, msg, err := c.ResetCfwAclHitCountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ResetCfwAclHitCountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResetCfwAclHitCountWithContext(ctx context.Context, request *ResetCfwAclHitCountRequest) string {
	if request == nil {
		request = NewResetCfwAclHitCountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ResetCfwAclHitCount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResetCfwAclHitCountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResetCfwAclHitCountWithContextV2(ctx context.Context, request *ResetCfwAclHitCountRequest) (int, string, error) {
	if request == nil {
		request = NewResetCfwAclHitCountRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ResetCfwAclHitCount")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResetCfwAclHitCountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAlterAclPriorityRequest() (request *AlterAclPriorityRequest) {
	request = &AlterAclPriorityRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "AlterAclPriority")
	return
}

func NewAlterAclPriorityResponse() (response *AlterAclPriorityResponse) {
	response = &AlterAclPriorityResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AlterAclPriority(request *AlterAclPriorityRequest) string {
	return c.AlterAclPriorityWithContext(context.Background(), request)
}

func (c *Client) AlterAclPrioritySend(request *AlterAclPriorityRequest) (*AlterAclPriorityResponse, error) {
	statusCode, msg, err := c.AlterAclPriorityWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AlterAclPriorityResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AlterAclPriorityWithContext(ctx context.Context, request *AlterAclPriorityRequest) string {
	if request == nil {
		request = NewAlterAclPriorityRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "AlterAclPriority")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAlterAclPriorityResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AlterAclPriorityWithContextV2(ctx context.Context, request *AlterAclPriorityRequest) (int, string, error) {
	if request == nil {
		request = NewAlterAclPriorityRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "AlterAclPriority")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAlterAclPriorityResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCfwAclRequest() (request *DeleteCfwAclRequest) {
	request = &DeleteCfwAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwAcl")
	return
}

func NewDeleteCfwAclResponse() (response *DeleteCfwAclResponse) {
	response = &DeleteCfwAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCfwAcl(request *DeleteCfwAclRequest) string {
	return c.DeleteCfwAclWithContext(context.Background(), request)
}

func (c *Client) DeleteCfwAclSend(request *DeleteCfwAclRequest) (*DeleteCfwAclResponse, error) {
	statusCode, msg, err := c.DeleteCfwAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCfwAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCfwAclWithContext(ctx context.Context, request *DeleteCfwAclRequest) string {
	if request == nil {
		request = NewDeleteCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCfwAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCfwAclWithContextV2(ctx context.Context, request *DeleteCfwAclRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCfwAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCfwAclRequest() (request *ModifyCfwAclRequest) {
	request = &ModifyCfwAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwAcl")
	return
}

func NewModifyCfwAclResponse() (response *ModifyCfwAclResponse) {
	response = &ModifyCfwAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCfwAcl(request *ModifyCfwAclRequest) string {
	return c.ModifyCfwAclWithContext(context.Background(), request)
}

func (c *Client) ModifyCfwAclSend(request *ModifyCfwAclRequest) (*ModifyCfwAclResponse, error) {
	statusCode, msg, err := c.ModifyCfwAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCfwAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCfwAclWithContext(ctx context.Context, request *ModifyCfwAclRequest) string {
	if request == nil {
		request = NewModifyCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCfwAclWithContextV2(ctx context.Context, request *ModifyCfwAclRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCfwAclRequest() (request *DescribeCfwAclRequest) {
	request = &DescribeCfwAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAcl")
	return
}

func NewDescribeCfwAclResponse() (response *DescribeCfwAclResponse) {
	response = &DescribeCfwAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfwAcl(request *DescribeCfwAclRequest) string {
	return c.DescribeCfwAclWithContext(context.Background(), request)
}

func (c *Client) DescribeCfwAclSend(request *DescribeCfwAclRequest) (*DescribeCfwAclResponse, error) {
	statusCode, msg, err := c.DescribeCfwAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCfwAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCfwAclWithContext(ctx context.Context, request *DescribeCfwAclRequest) string {
	if request == nil {
		request = NewDescribeCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCfwAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCfwAclWithContextV2(ctx context.Context, request *DescribeCfwAclRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCfwAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCfwAclRequest() (request *CreateCfwAclRequest) {
	request = &CreateCfwAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwAcl")
	return
}

func NewCreateCfwAclResponse() (response *CreateCfwAclResponse) {
	response = &CreateCfwAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCfwAcl(request *CreateCfwAclRequest) string {
	return c.CreateCfwAclWithContext(context.Background(), request)
}

func (c *Client) CreateCfwAclSend(request *CreateCfwAclRequest) (*CreateCfwAclResponse, error) {
	statusCode, msg, err := c.CreateCfwAclWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCfwAclResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCfwAclWithContext(ctx context.Context, request *CreateCfwAclRequest) string {
	if request == nil {
		request = NewCreateCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCfwAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCfwAclWithContextV2(ctx context.Context, request *CreateCfwAclRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCfwAclRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwAcl")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCfwAclResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCfwEipProtectRequest() (request *ModifyCfwEipProtectRequest) {
	request = &ModifyCfwEipProtectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwEipProtect")
	return
}

func NewModifyCfwEipProtectResponse() (response *ModifyCfwEipProtectResponse) {
	response = &ModifyCfwEipProtectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCfwEipProtect(request *ModifyCfwEipProtectRequest) string {
	return c.ModifyCfwEipProtectWithContext(context.Background(), request)
}

func (c *Client) ModifyCfwEipProtectSend(request *ModifyCfwEipProtectRequest) (*ModifyCfwEipProtectResponse, error) {
	statusCode, msg, err := c.ModifyCfwEipProtectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCfwEipProtectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCfwEipProtectWithContext(ctx context.Context, request *ModifyCfwEipProtectRequest) string {
	if request == nil {
		request = NewModifyCfwEipProtectRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwEipProtect")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwEipProtectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCfwEipProtectWithContextV2(ctx context.Context, request *ModifyCfwEipProtectRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCfwEipProtectRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwEipProtect")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwEipProtectResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCfwEipsRequest() (request *DescribeCfwEipsRequest) {
	request = &DescribeCfwEipsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwEips")
	return
}

func NewDescribeCfwEipsResponse() (response *DescribeCfwEipsResponse) {
	response = &DescribeCfwEipsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfwEips(request *DescribeCfwEipsRequest) string {
	return c.DescribeCfwEipsWithContext(context.Background(), request)
}

func (c *Client) DescribeCfwEipsSend(request *DescribeCfwEipsRequest) (*DescribeCfwEipsResponse, error) {
	statusCode, msg, err := c.DescribeCfwEipsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCfwEipsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCfwEipsWithContext(ctx context.Context, request *DescribeCfwEipsRequest) string {
	if request == nil {
		request = NewDescribeCfwEipsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwEips")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCfwEipsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCfwEipsWithContextV2(ctx context.Context, request *DescribeCfwEipsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCfwEipsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwEips")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCfwEipsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeServiceGroupRequest() (request *DescribeServiceGroupRequest) {
	request = &DescribeServiceGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeServiceGroup")
	return
}

func NewDescribeServiceGroupResponse() (response *DescribeServiceGroupResponse) {
	response = &DescribeServiceGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeServiceGroup(request *DescribeServiceGroupRequest) string {
	return c.DescribeServiceGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeServiceGroupSend(request *DescribeServiceGroupRequest) (*DescribeServiceGroupResponse, error) {
	statusCode, msg, err := c.DescribeServiceGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeServiceGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeServiceGroupWithContext(ctx context.Context, request *DescribeServiceGroupRequest) string {
	if request == nil {
		request = NewDescribeServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeServiceGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeServiceGroupWithContextV2(ctx context.Context, request *DescribeServiceGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeServiceGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCfwServiceGroupRequest() (request *ModifyCfwServiceGroupRequest) {
	request = &ModifyCfwServiceGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwServiceGroup")
	return
}

func NewModifyCfwServiceGroupResponse() (response *ModifyCfwServiceGroupResponse) {
	response = &ModifyCfwServiceGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCfwServiceGroup(request *ModifyCfwServiceGroupRequest) string {
	return c.ModifyCfwServiceGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyCfwServiceGroupSend(request *ModifyCfwServiceGroupRequest) (*ModifyCfwServiceGroupResponse, error) {
	statusCode, msg, err := c.ModifyCfwServiceGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCfwServiceGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCfwServiceGroupWithContext(ctx context.Context, request *ModifyCfwServiceGroupRequest) string {
	if request == nil {
		request = NewModifyCfwServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwServiceGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCfwServiceGroupWithContextV2(ctx context.Context, request *ModifyCfwServiceGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCfwServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "ModifyCfwServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCfwServiceGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCfwServiceGroupRequest() (request *DeleteCfwServiceGroupRequest) {
	request = &DeleteCfwServiceGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwServiceGroup")
	return
}

func NewDeleteCfwServiceGroupResponse() (response *DeleteCfwServiceGroupResponse) {
	response = &DeleteCfwServiceGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCfwServiceGroup(request *DeleteCfwServiceGroupRequest) string {
	return c.DeleteCfwServiceGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteCfwServiceGroupSend(request *DeleteCfwServiceGroupRequest) (*DeleteCfwServiceGroupResponse, error) {
	statusCode, msg, err := c.DeleteCfwServiceGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCfwServiceGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCfwServiceGroupWithContext(ctx context.Context, request *DeleteCfwServiceGroupRequest) string {
	if request == nil {
		request = NewDeleteCfwServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCfwServiceGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCfwServiceGroupWithContextV2(ctx context.Context, request *DeleteCfwServiceGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCfwServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DeleteCfwServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCfwServiceGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCfwServiceGroupRequest() (request *CreateCfwServiceGroupRequest) {
	request = &CreateCfwServiceGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwServiceGroup")
	return
}

func NewCreateCfwServiceGroupResponse() (response *CreateCfwServiceGroupResponse) {
	response = &CreateCfwServiceGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCfwServiceGroup(request *CreateCfwServiceGroupRequest) string {
	return c.CreateCfwServiceGroupWithContext(context.Background(), request)
}

func (c *Client) CreateCfwServiceGroupSend(request *CreateCfwServiceGroupRequest) (*CreateCfwServiceGroupResponse, error) {
	statusCode, msg, err := c.CreateCfwServiceGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCfwServiceGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCfwServiceGroupWithContext(ctx context.Context, request *CreateCfwServiceGroupRequest) string {
	if request == nil {
		request = NewCreateCfwServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCfwServiceGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCfwServiceGroupWithContextV2(ctx context.Context, request *CreateCfwServiceGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCfwServiceGroupRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "CreateCfwServiceGroup")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCfwServiceGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCloudFireWallInstanceRequest() (request *DescribeCloudFireWallInstanceRequest) {
	request = &DescribeCloudFireWallInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeCloudFireWallInstance")
	return
}

func NewDescribeCloudFireWallInstanceResponse() (response *DescribeCloudFireWallInstanceResponse) {
	response = &DescribeCloudFireWallInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCloudFireWallInstance(request *DescribeCloudFireWallInstanceRequest) string {
	return c.DescribeCloudFireWallInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeCloudFireWallInstanceSend(request *DescribeCloudFireWallInstanceRequest) (*DescribeCloudFireWallInstanceResponse, error) {
	statusCode, msg, err := c.DescribeCloudFireWallInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCloudFireWallInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCloudFireWallInstanceWithContext(ctx context.Context, request *DescribeCloudFireWallInstanceRequest) string {
	if request == nil {
		request = NewDescribeCloudFireWallInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCloudFireWallInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCloudFireWallInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCloudFireWallInstanceWithContextV2(ctx context.Context, request *DescribeCloudFireWallInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCloudFireWallInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kfw", APIVersion, "DescribeCloudFireWallInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCloudFireWallInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
