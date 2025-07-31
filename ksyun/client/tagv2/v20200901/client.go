package v20200901
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-09-01"

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

func NewCreateTagRequest() (request *CreateTagRequest) {
	request = &CreateTagRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "CreateTag")
	return
}

func NewCreateTagResponse() (response *CreateTagResponse) {
	response = &CreateTagResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTag(request *CreateTagRequest) string {
	return c.CreateTagWithContext(context.Background(), request)
}

func (c *Client) CreateTagSend(request *CreateTagRequest) (*CreateTagResponse, error) {
	statusCode, msg, err := c.CreateTagWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateTagResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTagWithContext(ctx context.Context, request *CreateTagRequest) string {
	if request == nil {
		request = NewCreateTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTagResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTagWithContextV2(ctx context.Context, request *CreateTagRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTagResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteTagRequest() (request *DeleteTagRequest) {
	request = &DeleteTagRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "DeleteTag")
	return
}

func NewDeleteTagResponse() (response *DeleteTagResponse) {
	response = &DeleteTagResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteTag(request *DeleteTagRequest) string {
	return c.DeleteTagWithContext(context.Background(), request)
}

func (c *Client) DeleteTagSend(request *DeleteTagRequest) (*DeleteTagResponse, error) {
	statusCode, msg, err := c.DeleteTagWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteTagResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteTagWithContext(ctx context.Context, request *DeleteTagRequest) string {
	if request == nil {
		request = NewDeleteTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteTagResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteTagWithContextV2(ctx context.Context, request *DeleteTagRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteTagRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteTagResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTagsRequest() (request *ListTagsRequest) {
	request = &ListTagsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "ListTags")
	return
}

func NewListTagsResponse() (response *ListTagsResponse) {
	response = &ListTagsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTags(request *ListTagsRequest) string {
	return c.ListTagsWithContext(context.Background(), request)
}

func (c *Client) ListTagsSend(request *ListTagsRequest) (*ListTagsResponse, error) {
	statusCode, msg, err := c.ListTagsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListTagsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest) string {
	if request == nil {
		request = NewListTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTagsWithContextV2(ctx context.Context, request *ListTagsRequest) (int, string, error) {
	if request == nil {
		request = NewListTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTagKeysRequest() (request *ListTagKeysRequest) {
	request = &ListTagKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "ListTagKeys")
	return
}

func NewListTagKeysResponse() (response *ListTagKeysResponse) {
	response = &ListTagKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTagKeys(request *ListTagKeysRequest) string {
	return c.ListTagKeysWithContext(context.Background(), request)
}

func (c *Client) ListTagKeysSend(request *ListTagKeysRequest) (*ListTagKeysResponse, error) {
	statusCode, msg, err := c.ListTagKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListTagKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest) string {
	if request == nil {
		request = NewListTagKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTagKeysWithContextV2(ctx context.Context, request *ListTagKeysRequest) (int, string, error) {
	if request == nil {
		request = NewListTagKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTagValuesRequest() (request *ListTagValuesRequest) {
	request = &ListTagValuesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "ListTagValues")
	return
}

func NewListTagValuesResponse() (response *ListTagValuesResponse) {
	response = &ListTagValuesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTagValues(request *ListTagValuesRequest) string {
	return c.ListTagValuesWithContext(context.Background(), request)
}

func (c *Client) ListTagValuesSend(request *ListTagValuesRequest) (*ListTagValuesResponse, error) {
	statusCode, msg, err := c.ListTagValuesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListTagValuesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest) string {
	if request == nil {
		request = NewListTagValuesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagValuesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTagValuesWithContextV2(ctx context.Context, request *ListTagValuesRequest) (int, string, error) {
	if request == nil {
		request = NewListTagValuesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagValuesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListResourcesRequest() (request *ListResourcesRequest) {
	request = &ListResourcesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "ListResources")
	return
}

func NewListResourcesResponse() (response *ListResourcesResponse) {
	response = &ListResourcesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListResources(request *ListResourcesRequest) string {
	return c.ListResourcesWithContext(context.Background(), request)
}

func (c *Client) ListResourcesSend(request *ListResourcesRequest) (*ListResourcesResponse, error) {
	statusCode, msg, err := c.ListResourcesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListResourcesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest) string {
	if request == nil {
		request = NewListResourcesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListResourcesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListResourcesWithContextV2(ctx context.Context, request *ListResourcesRequest) (int, string, error) {
	if request == nil {
		request = NewListResourcesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListResourcesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListTagsByResourceIdsRequest() (request *ListTagsByResourceIdsRequest) {
	request = &ListTagsByResourceIdsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "ListTagsByResourceIds")
	return
}

func NewListTagsByResourceIdsResponse() (response *ListTagsByResourceIdsResponse) {
	response = &ListTagsByResourceIdsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListTagsByResourceIds(request *ListTagsByResourceIdsRequest) string {
	return c.ListTagsByResourceIdsWithContext(context.Background(), request)
}

func (c *Client) ListTagsByResourceIdsSend(request *ListTagsByResourceIdsRequest) (*ListTagsByResourceIdsResponse, error) {
	statusCode, msg, err := c.ListTagsByResourceIdsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListTagsByResourceIdsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListTagsByResourceIdsWithContext(ctx context.Context, request *ListTagsByResourceIdsRequest) string {
	if request == nil {
		request = NewListTagsByResourceIdsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagsByResourceIdsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListTagsByResourceIdsWithContextV2(ctx context.Context, request *ListTagsByResourceIdsRequest) (int, string, error) {
	if request == nil {
		request = NewListTagsByResourceIdsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListTagsByResourceIdsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReplaceResourcesTagsRequest() (request *ReplaceResourcesTagsRequest) {
	request = &ReplaceResourcesTagsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "ReplaceResourcesTags")
	return
}

func NewReplaceResourcesTagsResponse() (response *ReplaceResourcesTagsResponse) {
	response = &ReplaceResourcesTagsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReplaceResourcesTags(request *ReplaceResourcesTagsRequest) string {
	return c.ReplaceResourcesTagsWithContext(context.Background(), request)
}

func (c *Client) ReplaceResourcesTagsSend(request *ReplaceResourcesTagsRequest) (*ReplaceResourcesTagsResponse, error) {
	statusCode, msg, err := c.ReplaceResourcesTagsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReplaceResourcesTagsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReplaceResourcesTagsWithContext(ctx context.Context, request *ReplaceResourcesTagsRequest) string {
	if request == nil {
		request = NewReplaceResourcesTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewReplaceResourcesTagsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReplaceResourcesTagsWithContextV2(ctx context.Context, request *ReplaceResourcesTagsRequest) (int, string, error) {
	if request == nil {
		request = NewReplaceResourcesTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewReplaceResourcesTagsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachResourceTagsRequest() (request *DetachResourceTagsRequest) {
	request = &DetachResourceTagsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "DetachResourceTags")
	return
}

func NewDetachResourceTagsResponse() (response *DetachResourceTagsResponse) {
	response = &DetachResourceTagsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachResourceTags(request *DetachResourceTagsRequest) string {
	return c.DetachResourceTagsWithContext(context.Background(), request)
}

func (c *Client) DetachResourceTagsSend(request *DetachResourceTagsRequest) (*DetachResourceTagsResponse, error) {
	statusCode, msg, err := c.DetachResourceTagsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachResourceTagsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachResourceTagsWithContext(ctx context.Context, request *DetachResourceTagsRequest) string {
	if request == nil {
		request = NewDetachResourceTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachResourceTagsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachResourceTagsWithContextV2(ctx context.Context, request *DetachResourceTagsRequest) (int, string, error) {
	if request == nil {
		request = NewDetachResourceTagsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachResourceTagsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateTagAndAttachResourceRequest() (request *CreateTagAndAttachResourceRequest) {
	request = &CreateTagAndAttachResourceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tagv2", APIVersion, "CreateTagAndAttachResource")
	return
}

func NewCreateTagAndAttachResourceResponse() (response *CreateTagAndAttachResourceResponse) {
	response = &CreateTagAndAttachResourceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTagAndAttachResource(request *CreateTagAndAttachResourceRequest) string {
	return c.CreateTagAndAttachResourceWithContext(context.Background(), request)
}

func (c *Client) CreateTagAndAttachResourceSend(request *CreateTagAndAttachResourceRequest) (*CreateTagAndAttachResourceResponse, error) {
	statusCode, msg, err := c.CreateTagAndAttachResourceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateTagAndAttachResourceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTagAndAttachResourceWithContext(ctx context.Context, request *CreateTagAndAttachResourceRequest) string {
	if request == nil {
		request = NewCreateTagAndAttachResourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateTagAndAttachResourceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTagAndAttachResourceWithContextV2(ctx context.Context, request *CreateTagAndAttachResourceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTagAndAttachResourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateTagAndAttachResourceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


