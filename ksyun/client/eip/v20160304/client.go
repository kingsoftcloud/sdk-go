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

func NewGetLinesRequest() (request *GetLinesRequest) {
	request = &GetLinesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "GetLines")
	return
}

func NewGetLinesResponse() (response *GetLinesResponse) {
	response = &GetLinesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetLines(request *GetLinesRequest) string {
	return c.GetLinesWithContext(context.Background(), request)
}

func (c *Client) GetLinesSend(request *GetLinesRequest) (*GetLinesResponse, error) {
	statusCode, msg, err := c.GetLinesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GetLinesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetLinesWithContext(ctx context.Context, request *GetLinesRequest) string {
	if request == nil {
		request = NewGetLinesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLinesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetLinesWithContextV2(ctx context.Context, request *GetLinesRequest) (int, string, error) {
	if request == nil {
		request = NewGetLinesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetLinesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
	request = &DescribeAddressesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "DescribeAddresses")
	return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
	response = &DescribeAddressesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) string {
	return c.DescribeAddressesWithContext(context.Background(), request)
}

func (c *Client) DescribeAddressesSend(request *DescribeAddressesRequest) (*DescribeAddressesResponse, error) {
	statusCode, msg, err := c.DescribeAddressesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeAddressesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAddressesWithContext(ctx context.Context, request *DescribeAddressesRequest) string {
	if request == nil {
		request = NewDescribeAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddressesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAddressesWithContextV2(ctx context.Context, request *DescribeAddressesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAddressesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddressesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateAddressRequest() (request *AllocateAddressRequest) {
	request = &AllocateAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "AllocateAddress")
	return
}

func NewAllocateAddressResponse() (response *AllocateAddressResponse) {
	response = &AllocateAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateAddress(request *AllocateAddressRequest) string {
	return c.AllocateAddressWithContext(context.Background(), request)
}

func (c *Client) AllocateAddressSend(request *AllocateAddressRequest) (*AllocateAddressResponse, error) {
	statusCode, msg, err := c.AllocateAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AllocateAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateAddressWithContext(ctx context.Context, request *AllocateAddressRequest) string {
	if request == nil {
		request = NewAllocateAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateAddressWithContextV2(ctx context.Context, request *AllocateAddressRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReleaseAddressRequest() (request *ReleaseAddressRequest) {
	request = &ReleaseAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "ReleaseAddress")
	return
}

func NewReleaseAddressResponse() (response *ReleaseAddressResponse) {
	response = &ReleaseAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReleaseAddress(request *ReleaseAddressRequest) string {
	return c.ReleaseAddressWithContext(context.Background(), request)
}

func (c *Client) ReleaseAddressSend(request *ReleaseAddressRequest) (*ReleaseAddressResponse, error) {
	statusCode, msg, err := c.ReleaseAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ReleaseAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReleaseAddressWithContext(ctx context.Context, request *ReleaseAddressRequest) string {
	if request == nil {
		request = NewReleaseAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReleaseAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReleaseAddressWithContextV2(ctx context.Context, request *ReleaseAddressRequest) (int, string, error) {
	if request == nil {
		request = NewReleaseAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReleaseAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
	request = &AssociateAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "AssociateAddress")
	return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
	response = &AssociateAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateAddress(request *AssociateAddressRequest) string {
	return c.AssociateAddressWithContext(context.Background(), request)
}

func (c *Client) AssociateAddressSend(request *AssociateAddressRequest) (*AssociateAddressResponse, error) {
	statusCode, msg, err := c.AssociateAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AssociateAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateAddressWithContext(ctx context.Context, request *AssociateAddressRequest) string {
	if request == nil {
		request = NewAssociateAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateAddressWithContextV2(ctx context.Context, request *AssociateAddressRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
	request = &DisassociateAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "DisassociateAddress")
	return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
	response = &DisassociateAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) string {
	return c.DisassociateAddressWithContext(context.Background(), request)
}

func (c *Client) DisassociateAddressSend(request *DisassociateAddressRequest) (*DisassociateAddressResponse, error) {
	statusCode, msg, err := c.DisassociateAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DisassociateAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateAddressWithContext(ctx context.Context, request *DisassociateAddressRequest) string {
	if request == nil {
		request = NewDisassociateAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateAddressWithContextV2(ctx context.Context, request *DisassociateAddressRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyAddressRequest() (request *ModifyAddressRequest) {
	request = &ModifyAddressRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "ModifyAddress")
	return
}

func NewModifyAddressResponse() (response *ModifyAddressResponse) {
	response = &ModifyAddressResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAddress(request *ModifyAddressRequest) string {
	return c.ModifyAddressWithContext(context.Background(), request)
}

func (c *Client) ModifyAddressSend(request *ModifyAddressRequest) (*ModifyAddressResponse, error) {
	statusCode, msg, err := c.ModifyAddressWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyAddressResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyAddressWithContext(ctx context.Context, request *ModifyAddressRequest) string {
	if request == nil {
		request = NewModifyAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAddressResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyAddressWithContextV2(ctx context.Context, request *ModifyAddressRequest) (int, string, error) {
	if request == nil {
		request = NewModifyAddressRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAddressResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateEipPoolRequest() (request *CreateEipPoolRequest) {
	request = &CreateEipPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "CreateEipPool")
	return
}

func NewCreateEipPoolResponse() (response *CreateEipPoolResponse) {
	response = &CreateEipPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateEipPool(request *CreateEipPoolRequest) string {
	return c.CreateEipPoolWithContext(context.Background(), request)
}

func (c *Client) CreateEipPoolSend(request *CreateEipPoolRequest) (*CreateEipPoolResponse, error) {
	statusCode, msg, err := c.CreateEipPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateEipPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateEipPoolWithContext(ctx context.Context, request *CreateEipPoolRequest) string {
	if request == nil {
		request = NewCreateEipPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEipPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateEipPoolWithContextV2(ctx context.Context, request *CreateEipPoolRequest) (int, string, error) {
	if request == nil {
		request = NewCreateEipPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEipPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteEipPoolRequest() (request *DeleteEipPoolRequest) {
	request = &DeleteEipPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "DeleteEipPool")
	return
}

func NewDeleteEipPoolResponse() (response *DeleteEipPoolResponse) {
	response = &DeleteEipPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteEipPool(request *DeleteEipPoolRequest) string {
	return c.DeleteEipPoolWithContext(context.Background(), request)
}

func (c *Client) DeleteEipPoolSend(request *DeleteEipPoolRequest) (*DeleteEipPoolResponse, error) {
	statusCode, msg, err := c.DeleteEipPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteEipPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteEipPoolWithContext(ctx context.Context, request *DeleteEipPoolRequest) string {
	if request == nil {
		request = NewDeleteEipPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEipPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteEipPoolWithContextV2(ctx context.Context, request *DeleteEipPoolRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteEipPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEipPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyEipPoolRequest() (request *ModifyEipPoolRequest) {
	request = &ModifyEipPoolRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "ModifyEipPool")
	return
}

func NewModifyEipPoolResponse() (response *ModifyEipPoolResponse) {
	response = &ModifyEipPoolResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyEipPool(request *ModifyEipPoolRequest) string {
	return c.ModifyEipPoolWithContext(context.Background(), request)
}

func (c *Client) ModifyEipPoolSend(request *ModifyEipPoolRequest) (*ModifyEipPoolResponse, error) {
	statusCode, msg, err := c.ModifyEipPoolWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyEipPoolResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyEipPoolWithContext(ctx context.Context, request *ModifyEipPoolRequest) string {
	if request == nil {
		request = NewModifyEipPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEipPoolResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyEipPoolWithContextV2(ctx context.Context, request *ModifyEipPoolRequest) (int, string, error) {
	if request == nil {
		request = NewModifyEipPoolRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEipPoolResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEipPoolsRequest() (request *DescribeEipPoolsRequest) {
	request = &DescribeEipPoolsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "DescribeEipPools")
	return
}

func NewDescribeEipPoolsResponse() (response *DescribeEipPoolsResponse) {
	response = &DescribeEipPoolsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEipPools(request *DescribeEipPoolsRequest) string {
	return c.DescribeEipPoolsWithContext(context.Background(), request)
}

func (c *Client) DescribeEipPoolsSend(request *DescribeEipPoolsRequest) (*DescribeEipPoolsResponse, error) {
	statusCode, msg, err := c.DescribeEipPoolsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeEipPoolsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEipPoolsWithContext(ctx context.Context, request *DescribeEipPoolsRequest) string {
	if request == nil {
		request = NewDescribeEipPoolsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEipPoolsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEipPoolsWithContextV2(ctx context.Context, request *DescribeEipPoolsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEipPoolsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEipPoolsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeIpExistEipPoolUseRequest() (request *DescribeIpExistEipPoolUseRequest) {
	request = &DescribeIpExistEipPoolUseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", APIVersion, "DescribeIpExistEipPoolUse")
	return
}

func NewDescribeIpExistEipPoolUseResponse() (response *DescribeIpExistEipPoolUseResponse) {
	response = &DescribeIpExistEipPoolUseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeIpExistEipPoolUse(request *DescribeIpExistEipPoolUseRequest) string {
	return c.DescribeIpExistEipPoolUseWithContext(context.Background(), request)
}

func (c *Client) DescribeIpExistEipPoolUseSend(request *DescribeIpExistEipPoolUseRequest) (*DescribeIpExistEipPoolUseResponse, error) {
	statusCode, msg, err := c.DescribeIpExistEipPoolUseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeIpExistEipPoolUseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeIpExistEipPoolUseWithContext(ctx context.Context, request *DescribeIpExistEipPoolUseRequest) string {
	if request == nil {
		request = NewDescribeIpExistEipPoolUseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeIpExistEipPoolUseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeIpExistEipPoolUseWithContextV2(ctx context.Context, request *DescribeIpExistEipPoolUseRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeIpExistEipPoolUseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeIpExistEipPoolUseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
