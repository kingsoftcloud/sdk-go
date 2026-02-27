package v20251114

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-11-14"

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

func NewDescribeKnowledgeBaseModelsRequest() (request *DescribeKnowledgeBaseModelsRequest) {
	request = &DescribeKnowledgeBaseModelsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBaseModels")
	return
}

func NewDescribeKnowledgeBaseModelsResponse() (response *DescribeKnowledgeBaseModelsResponse) {
	response = &DescribeKnowledgeBaseModelsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnowledgeBaseModels(request *DescribeKnowledgeBaseModelsRequest) string {
	return c.DescribeKnowledgeBaseModelsWithContext(context.Background(), request)
}

func (c *Client) DescribeKnowledgeBaseModelsSend(request *DescribeKnowledgeBaseModelsRequest) (*DescribeKnowledgeBaseModelsResponse, error) {
	statusCode, msg, err := c.DescribeKnowledgeBaseModelsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeKnowledgeBaseModelsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKnowledgeBaseModelsWithContext(ctx context.Context, request *DescribeKnowledgeBaseModelsRequest) string {
	if request == nil {
		request = NewDescribeKnowledgeBaseModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBaseModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnowledgeBaseModelsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKnowledgeBaseModelsWithContextV2(ctx context.Context, request *DescribeKnowledgeBaseModelsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKnowledgeBaseModelsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBaseModels")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnowledgeBaseModelsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateKnowledgeBaseServiceRequest() (request *ActivateKnowledgeBaseServiceRequest) {
	request = &ActivateKnowledgeBaseServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ActivateKnowledgeBaseService")
	return
}

func NewActivateKnowledgeBaseServiceResponse() (response *ActivateKnowledgeBaseServiceResponse) {
	response = &ActivateKnowledgeBaseServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateKnowledgeBaseService(request *ActivateKnowledgeBaseServiceRequest) string {
	return c.ActivateKnowledgeBaseServiceWithContext(context.Background(), request)
}

func (c *Client) ActivateKnowledgeBaseServiceSend(request *ActivateKnowledgeBaseServiceRequest) (*ActivateKnowledgeBaseServiceResponse, error) {
	statusCode, msg, err := c.ActivateKnowledgeBaseServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ActivateKnowledgeBaseServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateKnowledgeBaseServiceWithContext(ctx context.Context, request *ActivateKnowledgeBaseServiceRequest) string {
	if request == nil {
		request = NewActivateKnowledgeBaseServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateKnowledgeBaseService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateKnowledgeBaseServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateKnowledgeBaseServiceWithContextV2(ctx context.Context, request *ActivateKnowledgeBaseServiceRequest) (int, string, error) {
	if request == nil {
		request = NewActivateKnowledgeBaseServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateKnowledgeBaseService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateKnowledgeBaseServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRetrieveKnowledgeRequest() (request *RetrieveKnowledgeRequest) {
	request = &RetrieveKnowledgeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "RetrieveKnowledge")
	return
}

func NewRetrieveKnowledgeResponse() (response *RetrieveKnowledgeResponse) {
	response = &RetrieveKnowledgeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RetrieveKnowledge(request *RetrieveKnowledgeRequest) string {
	return c.RetrieveKnowledgeWithContext(context.Background(), request)
}

func (c *Client) RetrieveKnowledgeSend(request *RetrieveKnowledgeRequest) (*RetrieveKnowledgeResponse, error) {
	statusCode, msg, err := c.RetrieveKnowledgeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct RetrieveKnowledgeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RetrieveKnowledgeWithContext(ctx context.Context, request *RetrieveKnowledgeRequest) string {
	if request == nil {
		request = NewRetrieveKnowledgeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "RetrieveKnowledge")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRetrieveKnowledgeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RetrieveKnowledgeWithContextV2(ctx context.Context, request *RetrieveKnowledgeRequest) (int, string, error) {
	if request == nil {
		request = NewRetrieveKnowledgeRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "RetrieveKnowledge")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRetrieveKnowledgeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeChunkRequest() (request *DescribeChunkRequest) {
	request = &DescribeChunkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeChunk")
	return
}

func NewDescribeChunkResponse() (response *DescribeChunkResponse) {
	response = &DescribeChunkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeChunk(request *DescribeChunkRequest) string {
	return c.DescribeChunkWithContext(context.Background(), request)
}

func (c *Client) DescribeChunkSend(request *DescribeChunkRequest) (*DescribeChunkResponse, error) {
	statusCode, msg, err := c.DescribeChunkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeChunkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeChunkWithContext(ctx context.Context, request *DescribeChunkRequest) string {
	if request == nil {
		request = NewDescribeChunkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeChunk")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeChunkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeChunkWithContextV2(ctx context.Context, request *DescribeChunkRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeChunkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeChunk")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeChunkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchDisplayStatusRequest() (request *BatchDisplayStatusRequest) {
	request = &BatchDisplayStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "BatchDisplayStatus")
	return
}

func NewBatchDisplayStatusResponse() (response *BatchDisplayStatusResponse) {
	response = &BatchDisplayStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchDisplayStatus(request *BatchDisplayStatusRequest) string {
	return c.BatchDisplayStatusWithContext(context.Background(), request)
}

func (c *Client) BatchDisplayStatusSend(request *BatchDisplayStatusRequest) (*BatchDisplayStatusResponse, error) {
	statusCode, msg, err := c.BatchDisplayStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct BatchDisplayStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchDisplayStatusWithContext(ctx context.Context, request *BatchDisplayStatusRequest) string {
	if request == nil {
		request = NewBatchDisplayStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "BatchDisplayStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBatchDisplayStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchDisplayStatusWithContextV2(ctx context.Context, request *BatchDisplayStatusRequest) (int, string, error) {
	if request == nil {
		request = NewBatchDisplayStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "BatchDisplayStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewBatchDisplayStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisplayStatusRequest() (request *DisplayStatusRequest) {
	request = &DisplayStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DisplayStatus")
	return
}

func NewDisplayStatusResponse() (response *DisplayStatusResponse) {
	response = &DisplayStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisplayStatus(request *DisplayStatusRequest) string {
	return c.DisplayStatusWithContext(context.Background(), request)
}

func (c *Client) DisplayStatusSend(request *DisplayStatusRequest) (*DisplayStatusResponse, error) {
	statusCode, msg, err := c.DisplayStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisplayStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisplayStatusWithContext(ctx context.Context, request *DisplayStatusRequest) string {
	if request == nil {
		request = NewDisplayStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisplayStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisplayStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisplayStatusWithContextV2(ctx context.Context, request *DisplayStatusRequest) (int, string, error) {
	if request == nil {
		request = NewDisplayStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DisplayStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDisplayStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewIndexingStatusRequest() (request *IndexingStatusRequest) {
	request = &IndexingStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "IndexingStatus")
	return
}

func NewIndexingStatusResponse() (response *IndexingStatusResponse) {
	response = &IndexingStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) IndexingStatus(request *IndexingStatusRequest) string {
	return c.IndexingStatusWithContext(context.Background(), request)
}

func (c *Client) IndexingStatusSend(request *IndexingStatusRequest) (*IndexingStatusResponse, error) {
	statusCode, msg, err := c.IndexingStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct IndexingStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) IndexingStatusWithContext(ctx context.Context, request *IndexingStatusRequest) string {
	if request == nil {
		request = NewIndexingStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "IndexingStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewIndexingStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) IndexingStatusWithContextV2(ctx context.Context, request *IndexingStatusRequest) (int, string, error) {
	if request == nil {
		request = NewIndexingStatusRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "IndexingStatus")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewIndexingStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDocumentRequest() (request *DeleteDocumentRequest) {
	request = &DeleteDocumentRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteDocument")
	return
}

func NewDeleteDocumentResponse() (response *DeleteDocumentResponse) {
	response = &DeleteDocumentResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDocument(request *DeleteDocumentRequest) string {
	return c.DeleteDocumentWithContext(context.Background(), request)
}

func (c *Client) DeleteDocumentSend(request *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	statusCode, msg, err := c.DeleteDocumentWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteDocumentResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDocumentWithContext(ctx context.Context, request *DeleteDocumentRequest) string {
	if request == nil {
		request = NewDeleteDocumentRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteDocument")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDocumentResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDocumentWithContextV2(ctx context.Context, request *DeleteDocumentRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDocumentRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteDocument")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteDocumentResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDocumentRequest() (request *DescribeDocumentRequest) {
	request = &DescribeDocumentRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeDocument")
	return
}

func NewDescribeDocumentResponse() (response *DescribeDocumentResponse) {
	response = &DescribeDocumentResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDocument(request *DescribeDocumentRequest) string {
	return c.DescribeDocumentWithContext(context.Background(), request)
}

func (c *Client) DescribeDocumentSend(request *DescribeDocumentRequest) (*DescribeDocumentResponse, error) {
	statusCode, msg, err := c.DescribeDocumentWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDocumentResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDocumentWithContext(ctx context.Context, request *DescribeDocumentRequest) string {
	if request == nil {
		request = NewDescribeDocumentRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeDocument")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDocumentResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDocumentWithContextV2(ctx context.Context, request *DescribeDocumentRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDocumentRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeDocument")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDocumentResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDocumentsRequest() (request *DescribeDocumentsRequest) {
	request = &DescribeDocumentsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeDocuments")
	return
}

func NewDescribeDocumentsResponse() (response *DescribeDocumentsResponse) {
	response = &DescribeDocumentsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDocuments(request *DescribeDocumentsRequest) string {
	return c.DescribeDocumentsWithContext(context.Background(), request)
}

func (c *Client) DescribeDocumentsSend(request *DescribeDocumentsRequest) (*DescribeDocumentsResponse, error) {
	statusCode, msg, err := c.DescribeDocumentsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeDocumentsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDocumentsWithContext(ctx context.Context, request *DescribeDocumentsRequest) string {
	if request == nil {
		request = NewDescribeDocumentsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeDocuments")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDocumentsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDocumentsWithContextV2(ctx context.Context, request *DescribeDocumentsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDocumentsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeDocuments")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDocumentsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewImportDocumentsRequest() (request *ImportDocumentsRequest) {
	request = &ImportDocumentsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ImportDocuments")
	return
}

func NewImportDocumentsResponse() (response *ImportDocumentsResponse) {
	response = &ImportDocumentsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportDocuments(request *ImportDocumentsRequest) string {
	return c.ImportDocumentsWithContext(context.Background(), request)
}

func (c *Client) ImportDocumentsSend(request *ImportDocumentsRequest) (*ImportDocumentsResponse, error) {
	statusCode, msg, err := c.ImportDocumentsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ImportDocumentsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ImportDocumentsWithContext(ctx context.Context, request *ImportDocumentsRequest) string {
	if request == nil {
		request = NewImportDocumentsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ImportDocuments")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImportDocumentsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ImportDocumentsWithContextV2(ctx context.Context, request *ImportDocumentsRequest) (int, string, error) {
	if request == nil {
		request = NewImportDocumentsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ImportDocuments")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImportDocumentsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteKnowledgeBaseRequest() (request *DeleteKnowledgeBaseRequest) {
	request = &DeleteKnowledgeBaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteKnowledgeBase")
	return
}

func NewDeleteKnowledgeBaseResponse() (response *DeleteKnowledgeBaseResponse) {
	response = &DeleteKnowledgeBaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteKnowledgeBase(request *DeleteKnowledgeBaseRequest) string {
	return c.DeleteKnowledgeBaseWithContext(context.Background(), request)
}

func (c *Client) DeleteKnowledgeBaseSend(request *DeleteKnowledgeBaseRequest) (*DeleteKnowledgeBaseResponse, error) {
	statusCode, msg, err := c.DeleteKnowledgeBaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteKnowledgeBaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteKnowledgeBaseWithContext(ctx context.Context, request *DeleteKnowledgeBaseRequest) string {
	if request == nil {
		request = NewDeleteKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteKnowledgeBaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteKnowledgeBaseWithContextV2(ctx context.Context, request *DeleteKnowledgeBaseRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteKnowledgeBaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyKnowledgeBaseRequest() (request *ModifyKnowledgeBaseRequest) {
	request = &ModifyKnowledgeBaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyKnowledgeBase")
	return
}

func NewModifyKnowledgeBaseResponse() (response *ModifyKnowledgeBaseResponse) {
	response = &ModifyKnowledgeBaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyKnowledgeBase(request *ModifyKnowledgeBaseRequest) string {
	return c.ModifyKnowledgeBaseWithContext(context.Background(), request)
}

func (c *Client) ModifyKnowledgeBaseSend(request *ModifyKnowledgeBaseRequest) (*ModifyKnowledgeBaseResponse, error) {
	statusCode, msg, err := c.ModifyKnowledgeBaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyKnowledgeBaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyKnowledgeBaseWithContext(ctx context.Context, request *ModifyKnowledgeBaseRequest) string {
	if request == nil {
		request = NewModifyKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyKnowledgeBaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyKnowledgeBaseWithContextV2(ctx context.Context, request *ModifyKnowledgeBaseRequest) (int, string, error) {
	if request == nil {
		request = NewModifyKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyKnowledgeBaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKnowledgeBaseRequest() (request *DescribeKnowledgeBaseRequest) {
	request = &DescribeKnowledgeBaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBase")
	return
}

func NewDescribeKnowledgeBaseResponse() (response *DescribeKnowledgeBaseResponse) {
	response = &DescribeKnowledgeBaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnowledgeBase(request *DescribeKnowledgeBaseRequest) string {
	return c.DescribeKnowledgeBaseWithContext(context.Background(), request)
}

func (c *Client) DescribeKnowledgeBaseSend(request *DescribeKnowledgeBaseRequest) (*DescribeKnowledgeBaseResponse, error) {
	statusCode, msg, err := c.DescribeKnowledgeBaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeKnowledgeBaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKnowledgeBaseWithContext(ctx context.Context, request *DescribeKnowledgeBaseRequest) string {
	if request == nil {
		request = NewDescribeKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnowledgeBaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKnowledgeBaseWithContextV2(ctx context.Context, request *DescribeKnowledgeBaseRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnowledgeBaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKnowledgeBasesRequest() (request *DescribeKnowledgeBasesRequest) {
	request = &DescribeKnowledgeBasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBases")
	return
}

func NewDescribeKnowledgeBasesResponse() (response *DescribeKnowledgeBasesResponse) {
	response = &DescribeKnowledgeBasesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnowledgeBases(request *DescribeKnowledgeBasesRequest) string {
	return c.DescribeKnowledgeBasesWithContext(context.Background(), request)
}

func (c *Client) DescribeKnowledgeBasesSend(request *DescribeKnowledgeBasesRequest) (*DescribeKnowledgeBasesResponse, error) {
	statusCode, msg, err := c.DescribeKnowledgeBasesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeKnowledgeBasesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKnowledgeBasesWithContext(ctx context.Context, request *DescribeKnowledgeBasesRequest) string {
	if request == nil {
		request = NewDescribeKnowledgeBasesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBases")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnowledgeBasesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKnowledgeBasesWithContextV2(ctx context.Context, request *DescribeKnowledgeBasesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKnowledgeBasesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeKnowledgeBases")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnowledgeBasesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateKnowledgeBaseRequest() (request *CreateKnowledgeBaseRequest) {
	request = &CreateKnowledgeBaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateKnowledgeBase")
	return
}

func NewCreateKnowledgeBaseResponse() (response *CreateKnowledgeBaseResponse) {
	response = &CreateKnowledgeBaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateKnowledgeBase(request *CreateKnowledgeBaseRequest) string {
	return c.CreateKnowledgeBaseWithContext(context.Background(), request)
}

func (c *Client) CreateKnowledgeBaseSend(request *CreateKnowledgeBaseRequest) (*CreateKnowledgeBaseResponse, error) {
	statusCode, msg, err := c.CreateKnowledgeBaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateKnowledgeBaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateKnowledgeBaseWithContext(ctx context.Context, request *CreateKnowledgeBaseRequest) string {
	if request == nil {
		request = NewCreateKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateKnowledgeBaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateKnowledgeBaseWithContextV2(ctx context.Context, request *CreateKnowledgeBaseRequest) (int, string, error) {
	if request == nil {
		request = NewCreateKnowledgeBaseRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateKnowledgeBase")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateKnowledgeBaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
