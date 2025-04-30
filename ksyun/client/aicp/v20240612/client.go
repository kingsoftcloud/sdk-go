package v20240612

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-06-12"

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

func NewSaveNotebookImageRequest() (request *SaveNotebookImageRequest) {
	request = &SaveNotebookImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "SaveNotebookImage")
	return
}

func NewSaveNotebookImageResponse() (response *SaveNotebookImageResponse) {
	response = &SaveNotebookImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SaveNotebookImage(request *SaveNotebookImageRequest) string {
	return c.SaveNotebookImageWithContext(context.Background(), request)
}

func (c *Client) SaveNotebookImageWithContext(ctx context.Context, request *SaveNotebookImageRequest) string {
	if request == nil {
		request = NewSaveNotebookImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSaveNotebookImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyNotebookRequest() (request *ModifyNotebookRequest) {
	request = &ModifyNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ModifyNotebook")
	return
}

func NewModifyNotebookResponse() (response *ModifyNotebookResponse) {
	response = &ModifyNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNotebook(request *ModifyNotebookRequest) string {
	return c.ModifyNotebookWithContext(context.Background(), request)
}

func (c *Client) ModifyNotebookWithContext(ctx context.Context, request *ModifyNotebookRequest) string {
	if request == nil {
		request = NewModifyNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteNotebookRequest() (request *DeleteNotebookRequest) {
	request = &DeleteNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DeleteNotebook")
	return
}

func NewDeleteNotebookResponse() (response *DeleteNotebookResponse) {
	response = &DeleteNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNotebook(request *DeleteNotebookRequest) string {
	return c.DeleteNotebookWithContext(context.Background(), request)
}

func (c *Client) DeleteNotebookWithContext(ctx context.Context, request *DeleteNotebookRequest) string {
	if request == nil {
		request = NewDeleteNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNotebooksRequest() (request *DescribeNotebooksRequest) {
	request = &DescribeNotebooksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebooks")
	return
}

func NewDescribeNotebooksResponse() (response *DescribeNotebooksResponse) {
	response = &DescribeNotebooksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNotebooks(request *DescribeNotebooksRequest) string {
	return c.DescribeNotebooksWithContext(context.Background(), request)
}

func (c *Client) DescribeNotebooksWithContext(ctx context.Context, request *DescribeNotebooksRequest) string {
	if request == nil {
		request = NewDescribeNotebooksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebooksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateNotebookRequest() (request *CreateNotebookRequest) {
	request = &CreateNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateNotebook")
	return
}

func NewCreateNotebookResponse() (response *CreateNotebookResponse) {
	response = &CreateNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateNotebook(request *CreateNotebookRequest) string {
	return c.CreateNotebookWithContext(context.Background(), request)
}

func (c *Client) CreateNotebookWithContext(ctx context.Context, request *CreateNotebookRequest) string {
	if request == nil {
		request = NewCreateNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListDatasetHostnameRequest() (request *ListDatasetHostnameRequest) {
	request = &ListDatasetHostnameRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListDatasetHostname")
	return
}

func NewListDatasetHostnameResponse() (response *ListDatasetHostnameResponse) {
	response = &ListDatasetHostnameResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListDatasetHostname(request *ListDatasetHostnameRequest) string {
	return c.ListDatasetHostnameWithContext(context.Background(), request)
}

func (c *Client) ListDatasetHostnameWithContext(ctx context.Context, request *ListDatasetHostnameRequest) string {
	if request == nil {
		request = NewListDatasetHostnameRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListDatasetHostnameResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListDatasetTopicRequest() (request *ListDatasetTopicRequest) {
	request = &ListDatasetTopicRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListDatasetTopic")
	return
}

func NewListDatasetTopicResponse() (response *ListDatasetTopicResponse) {
	response = &ListDatasetTopicResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListDatasetTopic(request *ListDatasetTopicRequest) string {
	return c.ListDatasetTopicWithContext(context.Background(), request)
}

func (c *Client) ListDatasetTopicWithContext(ctx context.Context, request *ListDatasetTopicRequest) string {
	if request == nil {
		request = NewListDatasetTopicRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListDatasetTopicResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListDatasetRequest() (request *ListDatasetRequest) {
	request = &ListDatasetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ListDataset")
	return
}

func NewListDatasetResponse() (response *ListDatasetResponse) {
	response = &ListDatasetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListDataset(request *ListDatasetRequest) string {
	return c.ListDatasetWithContext(context.Background(), request)
}

func (c *Client) ListDatasetWithContext(ctx context.Context, request *ListDatasetRequest) string {
	if request == nil {
		request = NewListDatasetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListDatasetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDatasetRequest() (request *DescribeDatasetRequest) {
	request = &DescribeDatasetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeDataset")
	return
}

func NewDescribeDatasetResponse() (response *DescribeDatasetResponse) {
	response = &DescribeDatasetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDataset(request *DescribeDatasetRequest) string {
	return c.DescribeDatasetWithContext(context.Background(), request)
}

func (c *Client) DescribeDatasetWithContext(ctx context.Context, request *DescribeDatasetRequest) string {
	if request == nil {
		request = NewDescribeDatasetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDatasetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImportDatasetRequest() (request *ImportDatasetRequest) {
	request = &ImportDatasetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "ImportDataset")
	return
}

func NewImportDatasetResponse() (response *ImportDatasetResponse) {
	response = &ImportDatasetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportDataset(request *ImportDatasetRequest) string {
	return c.ImportDatasetWithContext(context.Background(), request)
}

func (c *Client) ImportDatasetWithContext(ctx context.Context, request *ImportDatasetRequest) string {
	if request == nil {
		request = NewImportDatasetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImportDatasetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStopNotebookRequest() (request *StopNotebookRequest) {
	request = &StopNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopNotebook")
	return
}

func NewStopNotebookResponse() (response *StopNotebookResponse) {
	response = &StopNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopNotebook(request *StopNotebookRequest) string {
	return c.StopNotebookWithContext(context.Background(), request)
}

func (c *Client) StopNotebookWithContext(ctx context.Context, request *StopNotebookRequest) string {
	if request == nil {
		request = NewStopNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStartNotebookRequest() (request *StartNotebookRequest) {
	request = &StartNotebookRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StartNotebook")
	return
}

func NewStartNotebookResponse() (response *StartNotebookResponse) {
	response = &StartNotebookResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartNotebook(request *StartNotebookRequest) string {
	return c.StartNotebookWithContext(context.Background(), request)
}

func (c *Client) StartNotebookWithContext(ctx context.Context, request *StartNotebookRequest) string {
	if request == nil {
		request = NewStartNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartNotebookResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetWebIdeUrlRequest() (request *GetWebIdeUrlRequest) {
	request = &GetWebIdeUrlRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "GetWebIdeUrl")
	return
}

func NewGetWebIdeUrlResponse() (response *GetWebIdeUrlResponse) {
	response = &GetWebIdeUrlResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetWebIdeUrl(request *GetWebIdeUrlRequest) string {
	return c.GetWebIdeUrlWithContext(context.Background(), request)
}

func (c *Client) GetWebIdeUrlWithContext(ctx context.Context, request *GetWebIdeUrlRequest) string {
	if request == nil {
		request = NewGetWebIdeUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetWebIdeUrlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNotebookEventsRequest() (request *DescribeNotebookEventsRequest) {
	request = &DescribeNotebookEventsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebookEvents")
	return
}

func NewDescribeNotebookEventsResponse() (response *DescribeNotebookEventsResponse) {
	response = &DescribeNotebookEventsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNotebookEvents(request *DescribeNotebookEventsRequest) string {
	return c.DescribeNotebookEventsWithContext(context.Background(), request)
}

func (c *Client) DescribeNotebookEventsWithContext(ctx context.Context, request *DescribeNotebookEventsRequest) string {
	if request == nil {
		request = NewDescribeNotebookEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeNotebookEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
