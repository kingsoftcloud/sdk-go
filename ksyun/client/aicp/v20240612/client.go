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

func (c *Client) SaveNotebookImageSend(request *SaveNotebookImageRequest) (*SaveNotebookImageResponse, error) {
	statusCode, msg, err := c.SaveNotebookImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SaveNotebookImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) SaveNotebookImageWithContextV2(ctx context.Context, request *SaveNotebookImageRequest) (int, string, error) {
	if request == nil {
		request = NewSaveNotebookImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSaveNotebookImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ModifyNotebookSend(request *ModifyNotebookRequest) (*ModifyNotebookResponse, error) {
	statusCode, msg, err := c.ModifyNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ModifyNotebookWithContextV2(ctx context.Context, request *ModifyNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteNotebookSend(request *DeleteNotebookRequest) (*DeleteNotebookResponse, error) {
	statusCode, msg, err := c.DeleteNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteNotebookWithContextV2(ctx context.Context, request *DeleteNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNotebooksSend(request *DescribeNotebooksRequest) (*DescribeNotebooksResponse, error) {
	statusCode, msg, err := c.DescribeNotebooksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNotebooksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeNotebooksWithContextV2(ctx context.Context, request *DescribeNotebooksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNotebooksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebooksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateNotebookSend(request *CreateNotebookRequest) (*CreateNotebookResponse, error) {
	statusCode, msg, err := c.CreateNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateNotebookWithContextV2(ctx context.Context, request *CreateNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewCreateNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StopNotebookSend(request *StopNotebookRequest) (*StopNotebookResponse, error) {
	statusCode, msg, err := c.StopNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StopNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StopNotebookWithContextV2(ctx context.Context, request *StopNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewStopNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStopNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StartNotebookSend(request *StartNotebookRequest) (*StartNotebookResponse, error) {
	statusCode, msg, err := c.StartNotebookWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartNotebookResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StartNotebookWithContextV2(ctx context.Context, request *StartNotebookRequest) (int, string, error) {
	if request == nil {
		request = NewStartNotebookRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStartNotebookResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetWebIdeUrlSend(request *GetWebIdeUrlRequest) (*GetWebIdeUrlResponse, error) {
	statusCode, msg, err := c.GetWebIdeUrlWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetWebIdeUrlResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetWebIdeUrlWithContext(ctx context.Context, request *GetWebIdeUrlRequest) string {
	if request == nil {
		request = NewGetWebIdeUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetWebIdeUrlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetWebIdeUrlWithContextV2(ctx context.Context, request *GetWebIdeUrlRequest) (int, string, error) {
	if request == nil {
		request = NewGetWebIdeUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetWebIdeUrlResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeNotebookEventsSend(request *DescribeNotebookEventsRequest) (*DescribeNotebookEventsResponse, error) {
	statusCode, msg, err := c.DescribeNotebookEventsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNotebookEventsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNotebookEventsWithContext(ctx context.Context, request *DescribeNotebookEventsRequest) string {
	if request == nil {
		request = NewDescribeNotebookEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebookEventsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNotebookEventsWithContextV2(ctx context.Context, request *DescribeNotebookEventsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNotebookEventsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebookEventsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNotebookLogRequest() (request *DescribeNotebookLogRequest) {
	request = &DescribeNotebookLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeNotebookLog")
	return
}

func NewDescribeNotebookLogResponse() (response *DescribeNotebookLogResponse) {
	response = &DescribeNotebookLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNotebookLog(request *DescribeNotebookLogRequest) string {
	return c.DescribeNotebookLogWithContext(context.Background(), request)
}

func (c *Client) DescribeNotebookLogSend(request *DescribeNotebookLogRequest) (*DescribeNotebookLogResponse, error) {
	statusCode, msg, err := c.DescribeNotebookLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNotebookLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNotebookLogWithContext(ctx context.Context, request *DescribeNotebookLogRequest) string {
	if request == nil {
		request = NewDescribeNotebookLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebookLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNotebookLogWithContextV2(ctx context.Context, request *DescribeNotebookLogRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNotebookLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNotebookLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopNotebookSavingImageRequest() (request *StopNotebookSavingImageRequest) {
	request = &StopNotebookSavingImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "StopNotebookSavingImage")
	return
}

func NewStopNotebookSavingImageResponse() (response *StopNotebookSavingImageResponse) {
	response = &StopNotebookSavingImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopNotebookSavingImage(request *StopNotebookSavingImageRequest) string {
	return c.StopNotebookSavingImageWithContext(context.Background(), request)
}

func (c *Client) StopNotebookSavingImageSend(request *StopNotebookSavingImageRequest) (*StopNotebookSavingImageResponse, error) {
	statusCode, msg, err := c.StopNotebookSavingImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StopNotebookSavingImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopNotebookSavingImageWithContext(ctx context.Context, request *StopNotebookSavingImageRequest) string {
	if request == nil {
		request = NewStopNotebookSavingImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopNotebookSavingImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopNotebookSavingImageWithContextV2(ctx context.Context, request *StopNotebookSavingImageRequest) (int, string, error) {
	if request == nil {
		request = NewStopNotebookSavingImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopNotebookSavingImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
