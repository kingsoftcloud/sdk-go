package v20230323

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2023-03-23"

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

func NewCreateKnadRequest() (request *CreateKnadRequest) {
	request = &CreateKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "CreateKnad")
	return
}

func NewCreateKnadResponse() (response *CreateKnadResponse) {
	response = &CreateKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateKnad(request *CreateKnadRequest) string {
	return c.CreateKnadWithContext(context.Background(), request)
}

func (c *Client) CreateKnadSend(request *CreateKnadRequest) (*CreateKnadResponse, error) {
	statusCode, msg, err := c.CreateKnadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateKnadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateKnadWithContext(ctx context.Context, request *CreateKnadRequest) string {
	if request == nil {
		request = NewCreateKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateKnadWithContextV2(ctx context.Context, request *CreateKnadRequest) (int, string, error) {
	if request == nil {
		request = NewCreateKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateKnadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyKnadRequest() (request *ModifyKnadRequest) {
	request = &ModifyKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyKnad")
	return
}

func NewModifyKnadResponse() (response *ModifyKnadResponse) {
	response = &ModifyKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyKnad(request *ModifyKnadRequest) string {
	return c.ModifyKnadWithContext(context.Background(), request)
}

func (c *Client) ModifyKnadSend(request *ModifyKnadRequest) (*ModifyKnadResponse, error) {
	statusCode, msg, err := c.ModifyKnadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyKnadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyKnadWithContext(ctx context.Context, request *ModifyKnadRequest) string {
	if request == nil {
		request = NewModifyKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyKnadWithContextV2(ctx context.Context, request *ModifyKnadRequest) (int, string, error) {
	if request == nil {
		request = NewModifyKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKnadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnbindIpListRequest() (request *UnbindIpListRequest) {
	request = &UnbindIpListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "UnbindIpList")
	return
}

func NewUnbindIpListResponse() (response *UnbindIpListResponse) {
	response = &UnbindIpListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindIpList(request *UnbindIpListRequest) string {
	return c.UnbindIpListWithContext(context.Background(), request)
}

func (c *Client) UnbindIpListSend(request *UnbindIpListRequest) (*UnbindIpListResponse, error) {
	statusCode, msg, err := c.UnbindIpListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UnbindIpListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnbindIpListWithContext(ctx context.Context, request *UnbindIpListRequest) string {
	if request == nil {
		request = NewUnbindIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUnbindIpListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnbindIpListWithContextV2(ctx context.Context, request *UnbindIpListRequest) (int, string, error) {
	if request == nil {
		request = NewUnbindIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUnbindIpListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateIpRequest() (request *AssociateIpRequest) {
	request = &AssociateIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "AssociateIp")
	return
}

func NewAssociateIpResponse() (response *AssociateIpResponse) {
	response = &AssociateIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateIp(request *AssociateIpRequest) string {
	return c.AssociateIpWithContext(context.Background(), request)
}

func (c *Client) AssociateIpSend(request *AssociateIpRequest) (*AssociateIpResponse, error) {
	statusCode, msg, err := c.AssociateIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssociateIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateIpWithContext(ctx context.Context, request *AssociateIpRequest) string {
	if request == nil {
		request = NewAssociateIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateIpWithContextV2(ctx context.Context, request *AssociateIpRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateIpRequest() (request *DisassociateIpRequest) {
	request = &DisassociateIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DisassociateIp")
	return
}

func NewDisassociateIpResponse() (response *DisassociateIpResponse) {
	response = &DisassociateIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateIp(request *DisassociateIpRequest) string {
	return c.DisassociateIpWithContext(context.Background(), request)
}

func (c *Client) DisassociateIpSend(request *DisassociateIpRequest) (*DisassociateIpResponse, error) {
	statusCode, msg, err := c.DisassociateIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DisassociateIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateIpWithContext(ctx context.Context, request *DisassociateIpRequest) string {
	if request == nil {
		request = NewDisassociateIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateIpWithContextV2(ctx context.Context, request *DisassociateIpRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKnadIpRequest() (request *DescribeKnadIpRequest) {
	request = &DescribeKnadIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DescribeKnadIp")
	return
}

func NewDescribeKnadIpResponse() (response *DescribeKnadIpResponse) {
	response = &DescribeKnadIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnadIp(request *DescribeKnadIpRequest) string {
	return c.DescribeKnadIpWithContext(context.Background(), request)
}

func (c *Client) DescribeKnadIpSend(request *DescribeKnadIpRequest) (*DescribeKnadIpResponse, error) {
	statusCode, msg, err := c.DescribeKnadIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeKnadIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKnadIpWithContext(ctx context.Context, request *DescribeKnadIpRequest) string {
	if request == nil {
		request = NewDescribeKnadIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeKnadIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKnadIpWithContextV2(ctx context.Context, request *DescribeKnadIpRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKnadIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeKnadIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteKnadRequest() (request *DeleteKnadRequest) {
	request = &DeleteKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DeleteKnad")
	return
}

func NewDeleteKnadResponse() (response *DeleteKnadResponse) {
	response = &DeleteKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteKnad(request *DeleteKnadRequest) string {
	return c.DeleteKnadWithContext(context.Background(), request)
}

func (c *Client) DeleteKnadSend(request *DeleteKnadRequest) (*DeleteKnadResponse, error) {
	statusCode, msg, err := c.DeleteKnadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteKnadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteKnadWithContext(ctx context.Context, request *DeleteKnadRequest) string {
	if request == nil {
		request = NewDeleteKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteKnadWithContextV2(ctx context.Context, request *DeleteKnadRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteKnadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKnadRequest() (request *DescribeKnadRequest) {
	request = &DescribeKnadRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DescribeKnad")
	return
}

func NewDescribeKnadResponse() (response *DescribeKnadResponse) {
	response = &DescribeKnadResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKnad(request *DescribeKnadRequest) string {
	return c.DescribeKnadWithContext(context.Background(), request)
}

func (c *Client) DescribeKnadSend(request *DescribeKnadRequest) (*DescribeKnadResponse, error) {
	statusCode, msg, err := c.DescribeKnadWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeKnadResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKnadWithContext(ctx context.Context, request *DescribeKnadRequest) string {
	if request == nil {
		request = NewDescribeKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnadResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKnadWithContextV2(ctx context.Context, request *DescribeKnadRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKnadRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeKnadResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBWIpListRequest() (request *GetBWIpListRequest) {
	request = &GetBWIpListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetBWIpList")
	return
}

func NewGetBWIpListResponse() (response *GetBWIpListResponse) {
	response = &GetBWIpListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBWIpList(request *GetBWIpListRequest) string {
	return c.GetBWIpListWithContext(context.Background(), request)
}

func (c *Client) GetBWIpListSend(request *GetBWIpListRequest) (*GetBWIpListResponse, error) {
	statusCode, msg, err := c.GetBWIpListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetBWIpListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBWIpListWithContext(ctx context.Context, request *GetBWIpListRequest) string {
	if request == nil {
		request = NewGetBWIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBWIpListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBWIpListWithContextV2(ctx context.Context, request *GetBWIpListRequest) (int, string, error) {
	if request == nil {
		request = NewGetBWIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetBWIpListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBWRequest() (request *DeleteBWRequest) {
	request = &DeleteBWRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "DeleteBW")
	return
}

func NewDeleteBWResponse() (response *DeleteBWResponse) {
	response = &DeleteBWResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBW(request *DeleteBWRequest) string {
	return c.DeleteBWWithContext(context.Background(), request)
}

func (c *Client) DeleteBWSend(request *DeleteBWRequest) (*DeleteBWResponse, error) {
	statusCode, msg, err := c.DeleteBWWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteBWResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBWWithContext(ctx context.Context, request *DeleteBWRequest) string {
	if request == nil {
		request = NewDeleteBWRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBWResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBWWithContextV2(ctx context.Context, request *DeleteBWRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBWRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBWResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddBWIpListRequest() (request *AddBWIpListRequest) {
	request = &AddBWIpListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "AddBWIpList")
	return
}

func NewAddBWIpListResponse() (response *AddBWIpListResponse) {
	response = &AddBWIpListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddBWIpList(request *AddBWIpListRequest) string {
	return c.AddBWIpListWithContext(context.Background(), request)
}

func (c *Client) AddBWIpListSend(request *AddBWIpListRequest) (*AddBWIpListResponse, error) {
	statusCode, msg, err := c.AddBWIpListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddBWIpListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddBWIpListWithContext(ctx context.Context, request *AddBWIpListRequest) string {
	if request == nil {
		request = NewAddBWIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddBWIpListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddBWIpListWithContextV2(ctx context.Context, request *AddBWIpListRequest) (int, string, error) {
	if request == nil {
		request = NewAddBWIpListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAddBWIpListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetZoneListRequest() (request *GetZoneListRequest) {
	request = &GetZoneListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetZoneList")
	return
}

func NewGetZoneListResponse() (response *GetZoneListResponse) {
	response = &GetZoneListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetZoneList(request *GetZoneListRequest) string {
	return c.GetZoneListWithContext(context.Background(), request)
}

func (c *Client) GetZoneListSend(request *GetZoneListRequest) (*GetZoneListResponse, error) {
	statusCode, msg, err := c.GetZoneListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetZoneListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetZoneListWithContext(ctx context.Context, request *GetZoneListRequest) string {
	if request == nil {
		request = NewGetZoneListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetZoneListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetZoneListWithContextV2(ctx context.Context, request *GetZoneListRequest) (int, string, error) {
	if request == nil {
		request = NewGetZoneListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetZoneListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyPolicyRequest() (request *ModifyPolicyRequest) {
	request = &ModifyPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyPolicy")
	return
}

func NewModifyPolicyResponse() (response *ModifyPolicyResponse) {
	response = &ModifyPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPolicy(request *ModifyPolicyRequest) string {
	return c.ModifyPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyPolicySend(request *ModifyPolicyRequest) (*ModifyPolicyResponse, error) {
	statusCode, msg, err := c.ModifyPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyPolicyWithContext(ctx context.Context, request *ModifyPolicyRequest) string {
	if request == nil {
		request = NewModifyPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyPolicyWithContextV2(ctx context.Context, request *ModifyPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyBlockLocationRequest() (request *ModifyBlockLocationRequest) {
	request = &ModifyBlockLocationRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "ModifyBlockLocation")
	return
}

func NewModifyBlockLocationResponse() (response *ModifyBlockLocationResponse) {
	response = &ModifyBlockLocationResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBlockLocation(request *ModifyBlockLocationRequest) string {
	return c.ModifyBlockLocationWithContext(context.Background(), request)
}

func (c *Client) ModifyBlockLocationSend(request *ModifyBlockLocationRequest) (*ModifyBlockLocationResponse, error) {
	statusCode, msg, err := c.ModifyBlockLocationWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyBlockLocationResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyBlockLocationWithContext(ctx context.Context, request *ModifyBlockLocationRequest) string {
	if request == nil {
		request = NewModifyBlockLocationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBlockLocationResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyBlockLocationWithContextV2(ctx context.Context, request *ModifyBlockLocationRequest) (int, string, error) {
	if request == nil {
		request = NewModifyBlockLocationRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBlockLocationResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetBlockLocationsRequest() (request *GetBlockLocationsRequest) {
	request = &GetBlockLocationsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetBlockLocations")
	return
}

func NewGetBlockLocationsResponse() (response *GetBlockLocationsResponse) {
	response = &GetBlockLocationsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetBlockLocations(request *GetBlockLocationsRequest) string {
	return c.GetBlockLocationsWithContext(context.Background(), request)
}

func (c *Client) GetBlockLocationsSend(request *GetBlockLocationsRequest) (*GetBlockLocationsResponse, error) {
	statusCode, msg, err := c.GetBlockLocationsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetBlockLocationsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetBlockLocationsWithContext(ctx context.Context, request *GetBlockLocationsRequest) string {
	if request == nil {
		request = NewGetBlockLocationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBlockLocationsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetBlockLocationsWithContextV2(ctx context.Context, request *GetBlockLocationsRequest) (int, string, error) {
	if request == nil {
		request = NewGetBlockLocationsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetBlockLocationsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetKnadPolicyRequest() (request *GetKnadPolicyRequest) {
	request = &GetKnadPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("knad", APIVersion, "GetKnadPolicy")
	return
}

func NewGetKnadPolicyResponse() (response *GetKnadPolicyResponse) {
	response = &GetKnadPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetKnadPolicy(request *GetKnadPolicyRequest) string {
	return c.GetKnadPolicyWithContext(context.Background(), request)
}

func (c *Client) GetKnadPolicySend(request *GetKnadPolicyRequest) (*GetKnadPolicyResponse, error) {
	statusCode, msg, err := c.GetKnadPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetKnadPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetKnadPolicyWithContext(ctx context.Context, request *GetKnadPolicyRequest) string {
	if request == nil {
		request = NewGetKnadPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetKnadPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetKnadPolicyWithContextV2(ctx context.Context, request *GetKnadPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewGetKnadPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetKnadPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
