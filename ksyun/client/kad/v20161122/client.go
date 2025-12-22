package v20161122

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-11-22"

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

func NewCreateForwardConfRequest() (request *CreateForwardConfRequest) {
	request = &CreateForwardConfRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "CreateForwardConf")
	return
}

func NewCreateForwardConfResponse() (response *CreateForwardConfResponse) {
	response = &CreateForwardConfResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateForwardConf(request *CreateForwardConfRequest) string {
	return c.CreateForwardConfWithContext(context.Background(), request)
}

func (c *Client) CreateForwardConfSend(request *CreateForwardConfRequest) (*CreateForwardConfResponse, error) {
	statusCode, msg, err := c.CreateForwardConfWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateForwardConfResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateForwardConfWithContext(ctx context.Context, request *CreateForwardConfRequest) string {
	if request == nil {
		request = NewCreateForwardConfRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateForwardConfResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateForwardConfWithContextV2(ctx context.Context, request *CreateForwardConfRequest) (int, string, error) {
	if request == nil {
		request = NewCreateForwardConfRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateForwardConfResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteForwardConfRequest() (request *DeleteForwardConfRequest) {
	request = &DeleteForwardConfRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "DeleteForwardConf")
	return
}

func NewDeleteForwardConfResponse() (response *DeleteForwardConfResponse) {
	response = &DeleteForwardConfResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteForwardConf(request *DeleteForwardConfRequest) string {
	return c.DeleteForwardConfWithContext(context.Background(), request)
}

func (c *Client) DeleteForwardConfSend(request *DeleteForwardConfRequest) (*DeleteForwardConfResponse, error) {
	statusCode, msg, err := c.DeleteForwardConfWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteForwardConfResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteForwardConfWithContext(ctx context.Context, request *DeleteForwardConfRequest) string {
	if request == nil {
		request = NewDeleteForwardConfRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteForwardConfResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteForwardConfWithContextV2(ctx context.Context, request *DeleteForwardConfRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteForwardConfRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteForwardConfResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeForwardConfRequest() (request *DescribeForwardConfRequest) {
	request = &DescribeForwardConfRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "DescribeForwardConf")
	return
}

func NewDescribeForwardConfResponse() (response *DescribeForwardConfResponse) {
	response = &DescribeForwardConfResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeForwardConf(request *DescribeForwardConfRequest) string {
	return c.DescribeForwardConfWithContext(context.Background(), request)
}

func (c *Client) DescribeForwardConfSend(request *DescribeForwardConfRequest) (*DescribeForwardConfResponse, error) {
	statusCode, msg, err := c.DescribeForwardConfWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeForwardConfResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeForwardConfWithContext(ctx context.Context, request *DescribeForwardConfRequest) string {
	if request == nil {
		request = NewDescribeForwardConfRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeForwardConfResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeForwardConfWithContextV2(ctx context.Context, request *DescribeForwardConfRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeForwardConfRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeForwardConfResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateForwardSourceRequest() (request *CreateForwardSourceRequest) {
	request = &CreateForwardSourceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "CreateForwardSource")
	return
}

func NewCreateForwardSourceResponse() (response *CreateForwardSourceResponse) {
	response = &CreateForwardSourceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateForwardSource(request *CreateForwardSourceRequest) string {
	return c.CreateForwardSourceWithContext(context.Background(), request)
}

func (c *Client) CreateForwardSourceSend(request *CreateForwardSourceRequest) (*CreateForwardSourceResponse, error) {
	statusCode, msg, err := c.CreateForwardSourceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateForwardSourceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateForwardSourceWithContext(ctx context.Context, request *CreateForwardSourceRequest) string {
	if request == nil {
		request = NewCreateForwardSourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateForwardSourceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateForwardSourceWithContextV2(ctx context.Context, request *CreateForwardSourceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateForwardSourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateForwardSourceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteForwardSourceRequest() (request *DeleteForwardSourceRequest) {
	request = &DeleteForwardSourceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "DeleteForwardSource")
	return
}

func NewDeleteForwardSourceResponse() (response *DeleteForwardSourceResponse) {
	response = &DeleteForwardSourceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteForwardSource(request *DeleteForwardSourceRequest) string {
	return c.DeleteForwardSourceWithContext(context.Background(), request)
}

func (c *Client) DeleteForwardSourceSend(request *DeleteForwardSourceRequest) (*DeleteForwardSourceResponse, error) {
	statusCode, msg, err := c.DeleteForwardSourceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteForwardSourceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteForwardSourceWithContext(ctx context.Context, request *DeleteForwardSourceRequest) string {
	if request == nil {
		request = NewDeleteForwardSourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteForwardSourceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteForwardSourceWithContextV2(ctx context.Context, request *DeleteForwardSourceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteForwardSourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteForwardSourceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeForwardSourceRequest() (request *DescribeForwardSourceRequest) {
	request = &DescribeForwardSourceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "DescribeForwardSource")
	return
}

func NewDescribeForwardSourceResponse() (response *DescribeForwardSourceResponse) {
	response = &DescribeForwardSourceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeForwardSource(request *DescribeForwardSourceRequest) string {
	return c.DescribeForwardSourceWithContext(context.Background(), request)
}

func (c *Client) DescribeForwardSourceSend(request *DescribeForwardSourceRequest) (*DescribeForwardSourceResponse, error) {
	statusCode, msg, err := c.DescribeForwardSourceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeForwardSourceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeForwardSourceWithContext(ctx context.Context, request *DescribeForwardSourceRequest) string {
	if request == nil {
		request = NewDescribeForwardSourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeForwardSourceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeForwardSourceWithContextV2(ctx context.Context, request *DescribeForwardSourceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeForwardSourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeForwardSourceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetAttackLogRequest() (request *GetAttackLogRequest) {
	request = &GetAttackLogRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "GetAttackLog")
	return
}

func NewGetAttackLogResponse() (response *GetAttackLogResponse) {
	response = &GetAttackLogResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetAttackLog(request *GetAttackLogRequest) string {
	return c.GetAttackLogWithContext(context.Background(), request)
}

func (c *Client) GetAttackLogSend(request *GetAttackLogRequest) (*GetAttackLogResponse, error) {
	statusCode, msg, err := c.GetAttackLogWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetAttackLogResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetAttackLogWithContext(ctx context.Context, request *GetAttackLogRequest) string {
	if request == nil {
		request = NewGetAttackLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetAttackLogResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetAttackLogWithContextV2(ctx context.Context, request *GetAttackLogRequest) (int, string, error) {
	if request == nil {
		request = NewGetAttackLogRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetAttackLogResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeOverviewRequest() (request *DescribeOverviewRequest) {
	request = &DescribeOverviewRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kad", APIVersion, "DescribeOverview")
	return
}

func NewDescribeOverviewResponse() (response *DescribeOverviewResponse) {
	response = &DescribeOverviewResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeOverview(request *DescribeOverviewRequest) string {
	return c.DescribeOverviewWithContext(context.Background(), request)
}

func (c *Client) DescribeOverviewSend(request *DescribeOverviewRequest) (*DescribeOverviewResponse, error) {
	statusCode, msg, err := c.DescribeOverviewWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeOverviewResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeOverviewWithContext(ctx context.Context, request *DescribeOverviewRequest) string {
	if request == nil {
		request = NewDescribeOverviewRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeOverviewResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeOverviewWithContextV2(ctx context.Context, request *DescribeOverviewRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeOverviewRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeOverviewResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
