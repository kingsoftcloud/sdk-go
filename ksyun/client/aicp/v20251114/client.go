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
func NewCreateMemorySdkRequest() (request *CreateMemorySdkRequest) {
	request = &CreateMemorySdkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateMemorySdk")
	return
}

func NewCreateMemorySdkResponse() (response *CreateMemorySdkResponse) {
	response = &CreateMemorySdkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMemorySdk(request *CreateMemorySdkRequest) string {
	return c.CreateMemorySdkWithContext(context.Background(), request)
}

func (c *Client) CreateMemorySdkSend(request *CreateMemorySdkRequest) (*CreateMemorySdkResponse, error) {
	statusCode, msg, err := c.CreateMemorySdkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateMemorySdkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMemorySdkWithContext(ctx context.Context, request *CreateMemorySdkRequest) string {
	if request == nil {
		request = NewCreateMemorySdkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateMemorySdk")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMemorySdkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMemorySdkWithContextV2(ctx context.Context, request *CreateMemorySdkRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMemorySdkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateMemorySdk")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMemorySdkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryMemorySdkRequest() (request *QueryMemorySdkRequest) {
	request = &QueryMemorySdkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "QueryMemorySdk")
	return
}

func NewQueryMemorySdkResponse() (response *QueryMemorySdkResponse) {
	response = &QueryMemorySdkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryMemorySdk(request *QueryMemorySdkRequest) string {
	return c.QueryMemorySdkWithContext(context.Background(), request)
}

func (c *Client) QueryMemorySdkSend(request *QueryMemorySdkRequest) (*QueryMemorySdkResponse, error) {
	statusCode, msg, err := c.QueryMemorySdkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct QueryMemorySdkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryMemorySdkWithContext(ctx context.Context, request *QueryMemorySdkRequest) string {
	if request == nil {
		request = NewQueryMemorySdkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "QueryMemorySdk")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryMemorySdkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryMemorySdkWithContextV2(ctx context.Context, request *QueryMemorySdkRequest) (int, string, error) {
	if request == nil {
		request = NewQueryMemorySdkRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "QueryMemorySdk")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryMemorySdkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMemoryCollectionRequest() (request *CreateMemoryCollectionRequest) {
	request = &CreateMemoryCollectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateMemoryCollection")
	return
}

func NewCreateMemoryCollectionResponse() (response *CreateMemoryCollectionResponse) {
	response = &CreateMemoryCollectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMemoryCollection(request *CreateMemoryCollectionRequest) string {
	return c.CreateMemoryCollectionWithContext(context.Background(), request)
}

func (c *Client) CreateMemoryCollectionSend(request *CreateMemoryCollectionRequest) (*CreateMemoryCollectionResponse, error) {
	statusCode, msg, err := c.CreateMemoryCollectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateMemoryCollectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMemoryCollectionWithContext(ctx context.Context, request *CreateMemoryCollectionRequest) string {
	if request == nil {
		request = NewCreateMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMemoryCollectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMemoryCollectionWithContextV2(ctx context.Context, request *CreateMemoryCollectionRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMemoryCollectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetMemoryCollectionRequest() (request *GetMemoryCollectionRequest) {
	request = &GetMemoryCollectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetMemoryCollection")
	return
}

func NewGetMemoryCollectionResponse() (response *GetMemoryCollectionResponse) {
	response = &GetMemoryCollectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMemoryCollection(request *GetMemoryCollectionRequest) string {
	return c.GetMemoryCollectionWithContext(context.Background(), request)
}

func (c *Client) GetMemoryCollectionSend(request *GetMemoryCollectionRequest) (*GetMemoryCollectionResponse, error) {
	statusCode, msg, err := c.GetMemoryCollectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetMemoryCollectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetMemoryCollectionWithContext(ctx context.Context, request *GetMemoryCollectionRequest) string {
	if request == nil {
		request = NewGetMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMemoryCollectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetMemoryCollectionWithContextV2(ctx context.Context, request *GetMemoryCollectionRequest) (int, string, error) {
	if request == nil {
		request = NewGetMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMemoryCollectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListMemoryCollectionsRequest() (request *ListMemoryCollectionsRequest) {
	request = &ListMemoryCollectionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListMemoryCollections")
	return
}

func NewListMemoryCollectionsResponse() (response *ListMemoryCollectionsResponse) {
	response = &ListMemoryCollectionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListMemoryCollections(request *ListMemoryCollectionsRequest) string {
	return c.ListMemoryCollectionsWithContext(context.Background(), request)
}

func (c *Client) ListMemoryCollectionsSend(request *ListMemoryCollectionsRequest) (*ListMemoryCollectionsResponse, error) {
	statusCode, msg, err := c.ListMemoryCollectionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ListMemoryCollectionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListMemoryCollectionsWithContext(ctx context.Context, request *ListMemoryCollectionsRequest) string {
	if request == nil {
		request = NewListMemoryCollectionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ListMemoryCollections")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListMemoryCollectionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListMemoryCollectionsWithContextV2(ctx context.Context, request *ListMemoryCollectionsRequest) (int, string, error) {
	if request == nil {
		request = NewListMemoryCollectionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ListMemoryCollections")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListMemoryCollectionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMemoryCollectionRequest() (request *DeleteMemoryCollectionRequest) {
	request = &DeleteMemoryCollectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteMemoryCollection")
	return
}

func NewDeleteMemoryCollectionResponse() (response *DeleteMemoryCollectionResponse) {
	response = &DeleteMemoryCollectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMemoryCollection(request *DeleteMemoryCollectionRequest) string {
	return c.DeleteMemoryCollectionWithContext(context.Background(), request)
}

func (c *Client) DeleteMemoryCollectionSend(request *DeleteMemoryCollectionRequest) (*DeleteMemoryCollectionResponse, error) {
	statusCode, msg, err := c.DeleteMemoryCollectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteMemoryCollectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMemoryCollectionWithContext(ctx context.Context, request *DeleteMemoryCollectionRequest) string {
	if request == nil {
		request = NewDeleteMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMemoryCollectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMemoryCollectionWithContextV2(ctx context.Context, request *DeleteMemoryCollectionRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMemoryCollectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetMemoryBaseServiceRequest() (request *GetMemoryBaseServiceRequest) {
	request = &GetMemoryBaseServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetMemoryBaseService")
	return
}

func NewGetMemoryBaseServiceResponse() (response *GetMemoryBaseServiceResponse) {
	response = &GetMemoryBaseServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetMemoryBaseService(request *GetMemoryBaseServiceRequest) string {
	return c.GetMemoryBaseServiceWithContext(context.Background(), request)
}

func (c *Client) GetMemoryBaseServiceSend(request *GetMemoryBaseServiceRequest) (*GetMemoryBaseServiceResponse, error) {
	statusCode, msg, err := c.GetMemoryBaseServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetMemoryBaseServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetMemoryBaseServiceWithContext(ctx context.Context, request *GetMemoryBaseServiceRequest) string {
	if request == nil {
		request = NewGetMemoryBaseServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetMemoryBaseService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMemoryBaseServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetMemoryBaseServiceWithContextV2(ctx context.Context, request *GetMemoryBaseServiceRequest) (int, string, error) {
	if request == nil {
		request = NewGetMemoryBaseServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "GetMemoryBaseService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetMemoryBaseServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateMemoryBaseServiceRequest() (request *ActivateMemoryBaseServiceRequest) {
	request = &ActivateMemoryBaseServiceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ActivateMemoryBaseService")
	return
}

func NewActivateMemoryBaseServiceResponse() (response *ActivateMemoryBaseServiceResponse) {
	response = &ActivateMemoryBaseServiceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateMemoryBaseService(request *ActivateMemoryBaseServiceRequest) string {
	return c.ActivateMemoryBaseServiceWithContext(context.Background(), request)
}

func (c *Client) ActivateMemoryBaseServiceSend(request *ActivateMemoryBaseServiceRequest) (*ActivateMemoryBaseServiceResponse, error) {
	statusCode, msg, err := c.ActivateMemoryBaseServiceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ActivateMemoryBaseServiceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateMemoryBaseServiceWithContext(ctx context.Context, request *ActivateMemoryBaseServiceRequest) string {
	if request == nil {
		request = NewActivateMemoryBaseServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateMemoryBaseService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateMemoryBaseServiceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateMemoryBaseServiceWithContextV2(ctx context.Context, request *ActivateMemoryBaseServiceRequest) (int, string, error) {
	if request == nil {
		request = NewActivateMemoryBaseServiceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateMemoryBaseService")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateMemoryBaseServiceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateMemoryCollectionRequest() (request *UpdateMemoryCollectionRequest) {
	request = &UpdateMemoryCollectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "UpdateMemoryCollection")
	return
}

func NewUpdateMemoryCollectionResponse() (response *UpdateMemoryCollectionResponse) {
	response = &UpdateMemoryCollectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateMemoryCollection(request *UpdateMemoryCollectionRequest) string {
	return c.UpdateMemoryCollectionWithContext(context.Background(), request)
}

func (c *Client) UpdateMemoryCollectionSend(request *UpdateMemoryCollectionRequest) (*UpdateMemoryCollectionResponse, error) {
	statusCode, msg, err := c.UpdateMemoryCollectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct UpdateMemoryCollectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateMemoryCollectionWithContext(ctx context.Context, request *UpdateMemoryCollectionRequest) string {
	if request == nil {
		request = NewUpdateMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateMemoryCollectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateMemoryCollectionWithContextV2(ctx context.Context, request *UpdateMemoryCollectionRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateMemoryCollectionRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "UpdateMemoryCollection")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateMemoryCollectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMcpServerRequest() (request *DeleteMcpServerRequest) {
	request = &DeleteMcpServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteMcpServer")
	return
}

func NewDeleteMcpServerResponse() (response *DeleteMcpServerResponse) {
	response = &DeleteMcpServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMcpServer(request *DeleteMcpServerRequest) string {
	return c.DeleteMcpServerWithContext(context.Background(), request)
}

func (c *Client) DeleteMcpServerSend(request *DeleteMcpServerRequest) (*DeleteMcpServerResponse, error) {
	statusCode, msg, err := c.DeleteMcpServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteMcpServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMcpServerWithContext(ctx context.Context, request *DeleteMcpServerRequest) string {
	if request == nil {
		request = NewDeleteMcpServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteMcpServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMcpServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMcpServerWithContextV2(ctx context.Context, request *DeleteMcpServerRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMcpServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeleteMcpServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMcpServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyMcpServerRequest() (request *ModifyMcpServerRequest) {
	request = &ModifyMcpServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyMcpServer")
	return
}

func NewModifyMcpServerResponse() (response *ModifyMcpServerResponse) {
	response = &ModifyMcpServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyMcpServer(request *ModifyMcpServerRequest) string {
	return c.ModifyMcpServerWithContext(context.Background(), request)
}

func (c *Client) ModifyMcpServerSend(request *ModifyMcpServerRequest) (*ModifyMcpServerResponse, error) {
	statusCode, msg, err := c.ModifyMcpServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyMcpServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyMcpServerWithContext(ctx context.Context, request *ModifyMcpServerRequest) string {
	if request == nil {
		request = NewModifyMcpServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyMcpServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyMcpServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyMcpServerWithContextV2(ctx context.Context, request *ModifyMcpServerRequest) (int, string, error) {
	if request == nil {
		request = NewModifyMcpServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ModifyMcpServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyMcpServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMcpServerRequest() (request *CreateMcpServerRequest) {
	request = &CreateMcpServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateMcpServer")
	return
}

func NewCreateMcpServerResponse() (response *CreateMcpServerResponse) {
	response = &CreateMcpServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMcpServer(request *CreateMcpServerRequest) string {
	return c.CreateMcpServerWithContext(context.Background(), request)
}

func (c *Client) CreateMcpServerSend(request *CreateMcpServerRequest) (*CreateMcpServerResponse, error) {
	statusCode, msg, err := c.CreateMcpServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateMcpServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMcpServerWithContext(ctx context.Context, request *CreateMcpServerRequest) string {
	if request == nil {
		request = NewCreateMcpServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateMcpServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMcpServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMcpServerWithContextV2(ctx context.Context, request *CreateMcpServerRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMcpServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "CreateMcpServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMcpServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMcpServersRequest() (request *DescribeMcpServersRequest) {
	request = &DescribeMcpServersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpServers")
	return
}

func NewDescribeMcpServersResponse() (response *DescribeMcpServersResponse) {
	response = &DescribeMcpServersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMcpServers(request *DescribeMcpServersRequest) string {
	return c.DescribeMcpServersWithContext(context.Background(), request)
}

func (c *Client) DescribeMcpServersSend(request *DescribeMcpServersRequest) (*DescribeMcpServersResponse, error) {
	statusCode, msg, err := c.DescribeMcpServersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeMcpServersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMcpServersWithContext(ctx context.Context, request *DescribeMcpServersRequest) string {
	if request == nil {
		request = NewDescribeMcpServersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpServers")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMcpServersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMcpServersWithContextV2(ctx context.Context, request *DescribeMcpServersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMcpServersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpServers")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMcpServersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMcpOfficialServersRequest() (request *DescribeMcpOfficialServersRequest) {
	request = &DescribeMcpOfficialServersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpOfficialServers")
	return
}

func NewDescribeMcpOfficialServersResponse() (response *DescribeMcpOfficialServersResponse) {
	response = &DescribeMcpOfficialServersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMcpOfficialServers(request *DescribeMcpOfficialServersRequest) string {
	return c.DescribeMcpOfficialServersWithContext(context.Background(), request)
}

func (c *Client) DescribeMcpOfficialServersSend(request *DescribeMcpOfficialServersRequest) (*DescribeMcpOfficialServersResponse, error) {
	statusCode, msg, err := c.DescribeMcpOfficialServersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeMcpOfficialServersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMcpOfficialServersWithContext(ctx context.Context, request *DescribeMcpOfficialServersRequest) string {
	if request == nil {
		request = NewDescribeMcpOfficialServersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpOfficialServers")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMcpOfficialServersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMcpOfficialServersWithContextV2(ctx context.Context, request *DescribeMcpOfficialServersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMcpOfficialServersRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpOfficialServers")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMcpOfficialServersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeactivateMcpOfficialServerRequest() (request *DeactivateMcpOfficialServerRequest) {
	request = &DeactivateMcpOfficialServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeactivateMcpOfficialServer")
	return
}

func NewDeactivateMcpOfficialServerResponse() (response *DeactivateMcpOfficialServerResponse) {
	response = &DeactivateMcpOfficialServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeactivateMcpOfficialServer(request *DeactivateMcpOfficialServerRequest) string {
	return c.DeactivateMcpOfficialServerWithContext(context.Background(), request)
}

func (c *Client) DeactivateMcpOfficialServerSend(request *DeactivateMcpOfficialServerRequest) (*DeactivateMcpOfficialServerResponse, error) {
	statusCode, msg, err := c.DeactivateMcpOfficialServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeactivateMcpOfficialServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeactivateMcpOfficialServerWithContext(ctx context.Context, request *DeactivateMcpOfficialServerRequest) string {
	if request == nil {
		request = NewDeactivateMcpOfficialServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeactivateMcpOfficialServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeactivateMcpOfficialServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeactivateMcpOfficialServerWithContextV2(ctx context.Context, request *DeactivateMcpOfficialServerRequest) (int, string, error) {
	if request == nil {
		request = NewDeactivateMcpOfficialServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DeactivateMcpOfficialServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeactivateMcpOfficialServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateMcpOfficialServerRequest() (request *ActivateMcpOfficialServerRequest) {
	request = &ActivateMcpOfficialServerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ActivateMcpOfficialServer")
	return
}

func NewActivateMcpOfficialServerResponse() (response *ActivateMcpOfficialServerResponse) {
	response = &ActivateMcpOfficialServerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateMcpOfficialServer(request *ActivateMcpOfficialServerRequest) string {
	return c.ActivateMcpOfficialServerWithContext(context.Background(), request)
}

func (c *Client) ActivateMcpOfficialServerSend(request *ActivateMcpOfficialServerRequest) (*ActivateMcpOfficialServerResponse, error) {
	statusCode, msg, err := c.ActivateMcpOfficialServerWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ActivateMcpOfficialServerResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateMcpOfficialServerWithContext(ctx context.Context, request *ActivateMcpOfficialServerRequest) string {
	if request == nil {
		request = NewActivateMcpOfficialServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateMcpOfficialServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateMcpOfficialServerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateMcpOfficialServerWithContextV2(ctx context.Context, request *ActivateMcpOfficialServerRequest) (int, string, error) {
	if request == nil {
		request = NewActivateMcpOfficialServerRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "ActivateMcpOfficialServer")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewActivateMcpOfficialServerResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMcpSquaresRequest() (request *DescribeMcpSquaresRequest) {
	request = &DescribeMcpSquaresRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpSquares")
	return
}

func NewDescribeMcpSquaresResponse() (response *DescribeMcpSquaresResponse) {
	response = &DescribeMcpSquaresResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMcpSquares(request *DescribeMcpSquaresRequest) string {
	return c.DescribeMcpSquaresWithContext(context.Background(), request)
}

func (c *Client) DescribeMcpSquaresSend(request *DescribeMcpSquaresRequest) (*DescribeMcpSquaresResponse, error) {
	statusCode, msg, err := c.DescribeMcpSquaresWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeMcpSquaresResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMcpSquaresWithContext(ctx context.Context, request *DescribeMcpSquaresRequest) string {
	if request == nil {
		request = NewDescribeMcpSquaresRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpSquares")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMcpSquaresResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMcpSquaresWithContextV2(ctx context.Context, request *DescribeMcpSquaresRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMcpSquaresRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("aicp", APIVersion, "DescribeMcpSquares")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMcpSquaresResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
