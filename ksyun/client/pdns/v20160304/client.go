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

func NewCreatePrivateDnsRequest() (request *CreatePrivateDnsRequest) {
	request = &CreatePrivateDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreatePrivateDns")
	return
}

func NewCreatePrivateDnsResponse() (response *CreatePrivateDnsResponse) {
	response = &CreatePrivateDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePrivateDns(request *CreatePrivateDnsRequest) string {
	return c.CreatePrivateDnsWithContext(context.Background(), request)
}

func (c *Client) CreatePrivateDnsSend(request *CreatePrivateDnsRequest) (*CreatePrivateDnsResponse, error) {
	statusCode, msg, err := c.CreatePrivateDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreatePrivateDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreatePrivateDnsWithContext(ctx context.Context, request *CreatePrivateDnsRequest) string {
	if request == nil {
		request = NewCreatePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrivateDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreatePrivateDnsWithContextV2(ctx context.Context, request *CreatePrivateDnsRequest) (int, string, error) {
	if request == nil {
		request = NewCreatePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePrivateDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePrivateDnsRequest() (request *DeletePrivateDnsRequest) {
	request = &DeletePrivateDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeletePrivateDns")
	return
}

func NewDeletePrivateDnsResponse() (response *DeletePrivateDnsResponse) {
	response = &DeletePrivateDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePrivateDns(request *DeletePrivateDnsRequest) string {
	return c.DeletePrivateDnsWithContext(context.Background(), request)
}

func (c *Client) DeletePrivateDnsSend(request *DeletePrivateDnsRequest) (*DeletePrivateDnsResponse, error) {
	statusCode, msg, err := c.DeletePrivateDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeletePrivateDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePrivateDnsWithContext(ctx context.Context, request *DeletePrivateDnsRequest) string {
	if request == nil {
		request = NewDeletePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePrivateDnsWithContextV2(ctx context.Context, request *DeletePrivateDnsRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePrivateDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePrivateDnsRequest() (request *DescribePrivateDnsRequest) {
	request = &DescribePrivateDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribePrivateDns")
	return
}

func NewDescribePrivateDnsResponse() (response *DescribePrivateDnsResponse) {
	response = &DescribePrivateDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrivateDns(request *DescribePrivateDnsRequest) string {
	return c.DescribePrivateDnsWithContext(context.Background(), request)
}

func (c *Client) DescribePrivateDnsSend(request *DescribePrivateDnsRequest) (*DescribePrivateDnsResponse, error) {
	statusCode, msg, err := c.DescribePrivateDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePrivateDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePrivateDnsWithContext(ctx context.Context, request *DescribePrivateDnsRequest) string {
	if request == nil {
		request = NewDescribePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePrivateDnsWithContextV2(ctx context.Context, request *DescribePrivateDnsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePrivateDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePrivateDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateVpcsRequest() (request *AssociateVpcsRequest) {
	request = &AssociateVpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "AssociateVpcs")
	return
}

func NewAssociateVpcsResponse() (response *AssociateVpcsResponse) {
	response = &AssociateVpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateVpcs(request *AssociateVpcsRequest) string {
	return c.AssociateVpcsWithContext(context.Background(), request)
}

func (c *Client) AssociateVpcsSend(request *AssociateVpcsRequest) (*AssociateVpcsResponse, error) {
	statusCode, msg, err := c.AssociateVpcsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssociateVpcsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateVpcsWithContext(ctx context.Context, request *AssociateVpcsRequest) string {
	if request == nil {
		request = NewAssociateVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateVpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateVpcsWithContextV2(ctx context.Context, request *AssociateVpcsRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateVpcsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateVpcsRequest() (request *DisassociateVpcsRequest) {
	request = &DisassociateVpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DisassociateVpcs")
	return
}

func NewDisassociateVpcsResponse() (response *DisassociateVpcsResponse) {
	response = &DisassociateVpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateVpcs(request *DisassociateVpcsRequest) string {
	return c.DisassociateVpcsWithContext(context.Background(), request)
}

func (c *Client) DisassociateVpcsSend(request *DisassociateVpcsRequest) (*DisassociateVpcsResponse, error) {
	statusCode, msg, err := c.DisassociateVpcsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DisassociateVpcsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateVpcsWithContext(ctx context.Context, request *DisassociateVpcsRequest) string {
	if request == nil {
		request = NewDisassociateVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateVpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateVpcsWithContextV2(ctx context.Context, request *DisassociateVpcsRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateVpcsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateZoneRequest() (request *CreateZoneRequest) {
	request = &CreateZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateZone")
	return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
	response = &CreateZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateZone(request *CreateZoneRequest) string {
	return c.CreateZoneWithContext(context.Background(), request)
}

func (c *Client) CreateZoneSend(request *CreateZoneRequest) (*CreateZoneResponse, error) {
	statusCode, msg, err := c.CreateZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) string {
	if request == nil {
		request = NewCreateZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateZoneWithContextV2(ctx context.Context, request *CreateZoneRequest) (int, string, error) {
	if request == nil {
		request = NewCreateZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
	request = &DeleteZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteZone")
	return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
	response = &DeleteZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteZone(request *DeleteZoneRequest) string {
	return c.DeleteZoneWithContext(context.Background(), request)
}

func (c *Client) DeleteZoneSend(request *DeleteZoneRequest) (*DeleteZoneResponse, error) {
	statusCode, msg, err := c.DeleteZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) string {
	if request == nil {
		request = NewDeleteZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteZoneWithContextV2(ctx context.Context, request *DeleteZoneRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyZoneRequest() (request *ModifyZoneRequest) {
	request = &ModifyZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyZone")
	return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
	response = &ModifyZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyZone(request *ModifyZoneRequest) string {
	return c.ModifyZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyZoneSend(request *ModifyZoneRequest) (*ModifyZoneResponse, error) {
	statusCode, msg, err := c.ModifyZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) string {
	if request == nil {
		request = NewModifyZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyZoneWithContextV2(ctx context.Context, request *ModifyZoneRequest) (int, string, error) {
	if request == nil {
		request = NewModifyZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeZoneRequest() (request *DescribeZoneRequest) {
	request = &DescribeZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeZone")
	return
}

func NewDescribeZoneResponse() (response *DescribeZoneResponse) {
	response = &DescribeZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeZone(request *DescribeZoneRequest) string {
	return c.DescribeZoneWithContext(context.Background(), request)
}

func (c *Client) DescribeZoneSend(request *DescribeZoneRequest) (*DescribeZoneResponse, error) {
	statusCode, msg, err := c.DescribeZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeZoneWithContext(ctx context.Context, request *DescribeZoneRequest) string {
	if request == nil {
		request = NewDescribeZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeZoneWithContextV2(ctx context.Context, request *DescribeZoneRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRecordRequest() (request *CreateRecordRequest) {
	request = &CreateRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateRecord")
	return
}

func NewCreateRecordResponse() (response *CreateRecordResponse) {
	response = &CreateRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRecord(request *CreateRecordRequest) string {
	return c.CreateRecordWithContext(context.Background(), request)
}

func (c *Client) CreateRecordSend(request *CreateRecordRequest) (*CreateRecordResponse, error) {
	statusCode, msg, err := c.CreateRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRecordWithContext(ctx context.Context, request *CreateRecordRequest) string {
	if request == nil {
		request = NewCreateRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRecordWithContextV2(ctx context.Context, request *CreateRecordRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
	request = &DeleteRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteRecord")
	return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
	response = &DeleteRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecord(request *DeleteRecordRequest) string {
	return c.DeleteRecordWithContext(context.Background(), request)
}

func (c *Client) DeleteRecordSend(request *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	statusCode, msg, err := c.DeleteRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) string {
	if request == nil {
		request = NewDeleteRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRecordWithContextV2(ctx context.Context, request *DeleteRecordRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
	request = &DescribeRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeRecord")
	return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
	response = &DescribeRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRecord(request *DescribeRecordRequest) string {
	return c.DescribeRecordWithContext(context.Background(), request)
}

func (c *Client) DescribeRecordSend(request *DescribeRecordRequest) (*DescribeRecordResponse, error) {
	statusCode, msg, err := c.DescribeRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRecordWithContext(ctx context.Context, request *DescribeRecordRequest) string {
	if request == nil {
		request = NewDescribeRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRecordWithContextV2(ctx context.Context, request *DescribeRecordRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRecordDataRequest() (request *CreateRecordDataRequest) {
	request = &CreateRecordDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateRecordData")
	return
}

func NewCreateRecordDataResponse() (response *CreateRecordDataResponse) {
	response = &CreateRecordDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRecordData(request *CreateRecordDataRequest) string {
	return c.CreateRecordDataWithContext(context.Background(), request)
}

func (c *Client) CreateRecordDataSend(request *CreateRecordDataRequest) (*CreateRecordDataResponse, error) {
	statusCode, msg, err := c.CreateRecordDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRecordDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRecordDataWithContext(ctx context.Context, request *CreateRecordDataRequest) string {
	if request == nil {
		request = NewCreateRecordDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRecordDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRecordDataWithContextV2(ctx context.Context, request *CreateRecordDataRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRecordDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRecordDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRecordDataRequest() (request *DeleteRecordDataRequest) {
	request = &DeleteRecordDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteRecordData")
	return
}

func NewDeleteRecordDataResponse() (response *DeleteRecordDataResponse) {
	response = &DeleteRecordDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRecordData(request *DeleteRecordDataRequest) string {
	return c.DeleteRecordDataWithContext(context.Background(), request)
}

func (c *Client) DeleteRecordDataSend(request *DeleteRecordDataRequest) (*DeleteRecordDataResponse, error) {
	statusCode, msg, err := c.DeleteRecordDataWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRecordDataResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRecordDataWithContext(ctx context.Context, request *DeleteRecordDataRequest) string {
	if request == nil {
		request = NewDeleteRecordDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecordDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRecordDataWithContextV2(ctx context.Context, request *DeleteRecordDataRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRecordDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRecordDataResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreatePdnsZoneRequest() (request *CreatePdnsZoneRequest) {
	request = &CreatePdnsZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreatePdnsZone")
	return
}

func NewCreatePdnsZoneResponse() (response *CreatePdnsZoneResponse) {
	response = &CreatePdnsZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePdnsZone(request *CreatePdnsZoneRequest) string {
	return c.CreatePdnsZoneWithContext(context.Background(), request)
}

func (c *Client) CreatePdnsZoneSend(request *CreatePdnsZoneRequest) (*CreatePdnsZoneResponse, error) {
	statusCode, msg, err := c.CreatePdnsZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreatePdnsZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreatePdnsZoneWithContext(ctx context.Context, request *CreatePdnsZoneRequest) string {
	if request == nil {
		request = NewCreatePdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePdnsZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreatePdnsZoneWithContextV2(ctx context.Context, request *CreatePdnsZoneRequest) (int, string, error) {
	if request == nil {
		request = NewCreatePdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePdnsZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyPdnsZoneRequest() (request *ModifyPdnsZoneRequest) {
	request = &ModifyPdnsZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyPdnsZone")
	return
}

func NewModifyPdnsZoneResponse() (response *ModifyPdnsZoneResponse) {
	response = &ModifyPdnsZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPdnsZone(request *ModifyPdnsZoneRequest) string {
	return c.ModifyPdnsZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyPdnsZoneSend(request *ModifyPdnsZoneRequest) (*ModifyPdnsZoneResponse, error) {
	statusCode, msg, err := c.ModifyPdnsZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyPdnsZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyPdnsZoneWithContext(ctx context.Context, request *ModifyPdnsZoneRequest) string {
	if request == nil {
		request = NewModifyPdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPdnsZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyPdnsZoneWithContextV2(ctx context.Context, request *ModifyPdnsZoneRequest) (int, string, error) {
	if request == nil {
		request = NewModifyPdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPdnsZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePdnsZoneRequest() (request *DeletePdnsZoneRequest) {
	request = &DeletePdnsZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeletePdnsZone")
	return
}

func NewDeletePdnsZoneResponse() (response *DeletePdnsZoneResponse) {
	response = &DeletePdnsZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePdnsZone(request *DeletePdnsZoneRequest) string {
	return c.DeletePdnsZoneWithContext(context.Background(), request)
}

func (c *Client) DeletePdnsZoneSend(request *DeletePdnsZoneRequest) (*DeletePdnsZoneResponse, error) {
	statusCode, msg, err := c.DeletePdnsZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeletePdnsZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePdnsZoneWithContext(ctx context.Context, request *DeletePdnsZoneRequest) string {
	if request == nil {
		request = NewDeletePdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePdnsZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePdnsZoneWithContextV2(ctx context.Context, request *DeletePdnsZoneRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePdnsZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePdnsZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePdnsZonesRequest() (request *DescribePdnsZonesRequest) {
	request = &DescribePdnsZonesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribePdnsZones")
	return
}

func NewDescribePdnsZonesResponse() (response *DescribePdnsZonesResponse) {
	response = &DescribePdnsZonesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePdnsZones(request *DescribePdnsZonesRequest) string {
	return c.DescribePdnsZonesWithContext(context.Background(), request)
}

func (c *Client) DescribePdnsZonesSend(request *DescribePdnsZonesRequest) (*DescribePdnsZonesResponse, error) {
	statusCode, msg, err := c.DescribePdnsZonesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePdnsZonesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePdnsZonesWithContext(ctx context.Context, request *DescribePdnsZonesRequest) string {
	if request == nil {
		request = NewDescribePdnsZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePdnsZonesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePdnsZonesWithContextV2(ctx context.Context, request *DescribePdnsZonesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePdnsZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePdnsZonesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBindZoneVpcRequest() (request *BindZoneVpcRequest) {
	request = &BindZoneVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "BindZoneVpc")
	return
}

func NewBindZoneVpcResponse() (response *BindZoneVpcResponse) {
	response = &BindZoneVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindZoneVpc(request *BindZoneVpcRequest) string {
	return c.BindZoneVpcWithContext(context.Background(), request)
}

func (c *Client) BindZoneVpcSend(request *BindZoneVpcRequest) (*BindZoneVpcResponse, error) {
	statusCode, msg, err := c.BindZoneVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct BindZoneVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BindZoneVpcWithContext(ctx context.Context, request *BindZoneVpcRequest) string {
	if request == nil {
		request = NewBindZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBindZoneVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BindZoneVpcWithContextV2(ctx context.Context, request *BindZoneVpcRequest) (int, string, error) {
	if request == nil {
		request = NewBindZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBindZoneVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnbindZoneVpcRequest() (request *UnbindZoneVpcRequest) {
	request = &UnbindZoneVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "UnbindZoneVpc")
	return
}

func NewUnbindZoneVpcResponse() (response *UnbindZoneVpcResponse) {
	response = &UnbindZoneVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindZoneVpc(request *UnbindZoneVpcRequest) string {
	return c.UnbindZoneVpcWithContext(context.Background(), request)
}

func (c *Client) UnbindZoneVpcSend(request *UnbindZoneVpcRequest) (*UnbindZoneVpcResponse, error) {
	statusCode, msg, err := c.UnbindZoneVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UnbindZoneVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnbindZoneVpcWithContext(ctx context.Context, request *UnbindZoneVpcRequest) string {
	if request == nil {
		request = NewUnbindZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnbindZoneVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnbindZoneVpcWithContextV2(ctx context.Context, request *UnbindZoneVpcRequest) (int, string, error) {
	if request == nil {
		request = NewUnbindZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnbindZoneVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateZoneRecordRequest() (request *CreateZoneRecordRequest) {
	request = &CreateZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateZoneRecord")
	return
}

func NewCreateZoneRecordResponse() (response *CreateZoneRecordResponse) {
	response = &CreateZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateZoneRecord(request *CreateZoneRecordRequest) string {
	return c.CreateZoneRecordWithContext(context.Background(), request)
}

func (c *Client) CreateZoneRecordSend(request *CreateZoneRecordRequest) (*CreateZoneRecordResponse, error) {
	statusCode, msg, err := c.CreateZoneRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateZoneRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateZoneRecordWithContext(ctx context.Context, request *CreateZoneRecordRequest) string {
	if request == nil {
		request = NewCreateZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateZoneRecordWithContextV2(ctx context.Context, request *CreateZoneRecordRequest) (int, string, error) {
	if request == nil {
		request = NewCreateZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateZoneRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteZoneRecordRequest() (request *DeleteZoneRecordRequest) {
	request = &DeleteZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteZoneRecord")
	return
}

func NewDeleteZoneRecordResponse() (response *DeleteZoneRecordResponse) {
	response = &DeleteZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteZoneRecord(request *DeleteZoneRecordRequest) string {
	return c.DeleteZoneRecordWithContext(context.Background(), request)
}

func (c *Client) DeleteZoneRecordSend(request *DeleteZoneRecordRequest) (*DeleteZoneRecordResponse, error) {
	statusCode, msg, err := c.DeleteZoneRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteZoneRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteZoneRecordWithContext(ctx context.Context, request *DeleteZoneRecordRequest) string {
	if request == nil {
		request = NewDeleteZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteZoneRecordWithContextV2(ctx context.Context, request *DeleteZoneRecordRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteZoneRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyZoneRecordRequest() (request *ModifyZoneRecordRequest) {
	request = &ModifyZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyZoneRecord")
	return
}

func NewModifyZoneRecordResponse() (response *ModifyZoneRecordResponse) {
	response = &ModifyZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyZoneRecord(request *ModifyZoneRecordRequest) string {
	return c.ModifyZoneRecordWithContext(context.Background(), request)
}

func (c *Client) ModifyZoneRecordSend(request *ModifyZoneRecordRequest) (*ModifyZoneRecordResponse, error) {
	statusCode, msg, err := c.ModifyZoneRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyZoneRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyZoneRecordWithContext(ctx context.Context, request *ModifyZoneRecordRequest) string {
	if request == nil {
		request = NewModifyZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyZoneRecordWithContextV2(ctx context.Context, request *ModifyZoneRecordRequest) (int, string, error) {
	if request == nil {
		request = NewModifyZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyZoneRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeZoneRecordRequest() (request *DescribeZoneRecordRequest) {
	request = &DescribeZoneRecordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeZoneRecord")
	return
}

func NewDescribeZoneRecordResponse() (response *DescribeZoneRecordResponse) {
	response = &DescribeZoneRecordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeZoneRecord(request *DescribeZoneRecordRequest) string {
	return c.DescribeZoneRecordWithContext(context.Background(), request)
}

func (c *Client) DescribeZoneRecordSend(request *DescribeZoneRecordRequest) (*DescribeZoneRecordResponse, error) {
	statusCode, msg, err := c.DescribeZoneRecordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeZoneRecordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeZoneRecordWithContext(ctx context.Context, request *DescribeZoneRecordRequest) string {
	if request == nil {
		request = NewDescribeZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeZoneRecordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeZoneRecordWithContextV2(ctx context.Context, request *DescribeZoneRecordRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeZoneRecordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeZoneRecordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUnbindFdZoneVpcRequest() (request *UnbindFdZoneVpcRequest) {
	request = &UnbindFdZoneVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "UnbindFdZoneVpc")
	return
}

func NewUnbindFdZoneVpcResponse() (response *UnbindFdZoneVpcResponse) {
	response = &UnbindFdZoneVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UnbindFdZoneVpc(request *UnbindFdZoneVpcRequest) string {
	return c.UnbindFdZoneVpcWithContext(context.Background(), request)
}

func (c *Client) UnbindFdZoneVpcSend(request *UnbindFdZoneVpcRequest) (*UnbindFdZoneVpcResponse, error) {
	statusCode, msg, err := c.UnbindFdZoneVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UnbindFdZoneVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UnbindFdZoneVpcWithContext(ctx context.Context, request *UnbindFdZoneVpcRequest) string {
	if request == nil {
		request = NewUnbindFdZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnbindFdZoneVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UnbindFdZoneVpcWithContextV2(ctx context.Context, request *UnbindFdZoneVpcRequest) (int, string, error) {
	if request == nil {
		request = NewUnbindFdZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUnbindFdZoneVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBindFdZoneVpcRequest() (request *BindFdZoneVpcRequest) {
	request = &BindFdZoneVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "BindFdZoneVpc")
	return
}

func NewBindFdZoneVpcResponse() (response *BindFdZoneVpcResponse) {
	response = &BindFdZoneVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BindFdZoneVpc(request *BindFdZoneVpcRequest) string {
	return c.BindFdZoneVpcWithContext(context.Background(), request)
}

func (c *Client) BindFdZoneVpcSend(request *BindFdZoneVpcRequest) (*BindFdZoneVpcResponse, error) {
	statusCode, msg, err := c.BindFdZoneVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct BindFdZoneVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BindFdZoneVpcWithContext(ctx context.Context, request *BindFdZoneVpcRequest) string {
	if request == nil {
		request = NewBindFdZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBindFdZoneVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BindFdZoneVpcWithContextV2(ctx context.Context, request *BindFdZoneVpcRequest) (int, string, error) {
	if request == nil {
		request = NewBindFdZoneVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBindFdZoneVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePdnsFdZoneRequest() (request *DescribePdnsFdZoneRequest) {
	request = &DescribePdnsFdZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribePdnsFdZone")
	return
}

func NewDescribePdnsFdZoneResponse() (response *DescribePdnsFdZoneResponse) {
	response = &DescribePdnsFdZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePdnsFdZone(request *DescribePdnsFdZoneRequest) string {
	return c.DescribePdnsFdZoneWithContext(context.Background(), request)
}

func (c *Client) DescribePdnsFdZoneSend(request *DescribePdnsFdZoneRequest) (*DescribePdnsFdZoneResponse, error) {
	statusCode, msg, err := c.DescribePdnsFdZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePdnsFdZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePdnsFdZoneWithContext(ctx context.Context, request *DescribePdnsFdZoneRequest) string {
	if request == nil {
		request = NewDescribePdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePdnsFdZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePdnsFdZoneWithContextV2(ctx context.Context, request *DescribePdnsFdZoneRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePdnsFdZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeletePdnsFdZoneRequest() (request *DeletePdnsFdZoneRequest) {
	request = &DeletePdnsFdZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeletePdnsFdZone")
	return
}

func NewDeletePdnsFdZoneResponse() (response *DeletePdnsFdZoneResponse) {
	response = &DeletePdnsFdZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeletePdnsFdZone(request *DeletePdnsFdZoneRequest) string {
	return c.DeletePdnsFdZoneWithContext(context.Background(), request)
}

func (c *Client) DeletePdnsFdZoneSend(request *DeletePdnsFdZoneRequest) (*DeletePdnsFdZoneResponse, error) {
	statusCode, msg, err := c.DeletePdnsFdZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeletePdnsFdZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeletePdnsFdZoneWithContext(ctx context.Context, request *DeletePdnsFdZoneRequest) string {
	if request == nil {
		request = NewDeletePdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePdnsFdZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeletePdnsFdZoneWithContextV2(ctx context.Context, request *DeletePdnsFdZoneRequest) (int, string, error) {
	if request == nil {
		request = NewDeletePdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeletePdnsFdZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyPdnsFdZoneRequest() (request *ModifyPdnsFdZoneRequest) {
	request = &ModifyPdnsFdZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyPdnsFdZone")
	return
}

func NewModifyPdnsFdZoneResponse() (response *ModifyPdnsFdZoneResponse) {
	response = &ModifyPdnsFdZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyPdnsFdZone(request *ModifyPdnsFdZoneRequest) string {
	return c.ModifyPdnsFdZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyPdnsFdZoneSend(request *ModifyPdnsFdZoneRequest) (*ModifyPdnsFdZoneResponse, error) {
	statusCode, msg, err := c.ModifyPdnsFdZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyPdnsFdZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyPdnsFdZoneWithContext(ctx context.Context, request *ModifyPdnsFdZoneRequest) string {
	if request == nil {
		request = NewModifyPdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPdnsFdZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyPdnsFdZoneWithContextV2(ctx context.Context, request *ModifyPdnsFdZoneRequest) (int, string, error) {
	if request == nil {
		request = NewModifyPdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyPdnsFdZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreatePdnsFdZoneRequest() (request *CreatePdnsFdZoneRequest) {
	request = &CreatePdnsFdZoneRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreatePdnsFdZone")
	return
}

func NewCreatePdnsFdZoneResponse() (response *CreatePdnsFdZoneResponse) {
	response = &CreatePdnsFdZoneResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatePdnsFdZone(request *CreatePdnsFdZoneRequest) string {
	return c.CreatePdnsFdZoneWithContext(context.Background(), request)
}

func (c *Client) CreatePdnsFdZoneSend(request *CreatePdnsFdZoneRequest) (*CreatePdnsFdZoneResponse, error) {
	statusCode, msg, err := c.CreatePdnsFdZoneWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreatePdnsFdZoneResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreatePdnsFdZoneWithContext(ctx context.Context, request *CreatePdnsFdZoneRequest) string {
	if request == nil {
		request = NewCreatePdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePdnsFdZoneResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreatePdnsFdZoneWithContextV2(ctx context.Context, request *CreatePdnsFdZoneRequest) (int, string, error) {
	if request == nil {
		request = NewCreatePdnsFdZoneRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatePdnsFdZoneResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryEndPointRegionAZRequest() (request *QueryEndPointRegionAZRequest) {
	request = &QueryEndPointRegionAZRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "QueryEndPointRegionAZ")
	return
}

func NewQueryEndPointRegionAZResponse() (response *QueryEndPointRegionAZResponse) {
	response = &QueryEndPointRegionAZResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryEndPointRegionAZ(request *QueryEndPointRegionAZRequest) string {
	return c.QueryEndPointRegionAZWithContext(context.Background(), request)
}

func (c *Client) QueryEndPointRegionAZSend(request *QueryEndPointRegionAZRequest) (*QueryEndPointRegionAZResponse, error) {
	statusCode, msg, err := c.QueryEndPointRegionAZWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryEndPointRegionAZResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryEndPointRegionAZWithContext(ctx context.Context, request *QueryEndPointRegionAZRequest) string {
	if request == nil {
		request = NewQueryEndPointRegionAZRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryEndPointRegionAZResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryEndPointRegionAZWithContextV2(ctx context.Context, request *QueryEndPointRegionAZRequest) (int, string, error) {
	if request == nil {
		request = NewQueryEndPointRegionAZRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryEndPointRegionAZResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEndPointsRequest() (request *DescribeEndPointsRequest) {
	request = &DescribeEndPointsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DescribeEndPoints")
	return
}

func NewDescribeEndPointsResponse() (response *DescribeEndPointsResponse) {
	response = &DescribeEndPointsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEndPoints(request *DescribeEndPointsRequest) string {
	return c.DescribeEndPointsWithContext(context.Background(), request)
}

func (c *Client) DescribeEndPointsSend(request *DescribeEndPointsRequest) (*DescribeEndPointsResponse, error) {
	statusCode, msg, err := c.DescribeEndPointsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEndPointsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEndPointsWithContext(ctx context.Context, request *DescribeEndPointsRequest) string {
	if request == nil {
		request = NewDescribeEndPointsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEndPointsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEndPointsWithContextV2(ctx context.Context, request *DescribeEndPointsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEndPointsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEndPointsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteEndPointRequest() (request *DeleteEndPointRequest) {
	request = &DeleteEndPointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "DeleteEndPoint")
	return
}

func NewDeleteEndPointResponse() (response *DeleteEndPointResponse) {
	response = &DeleteEndPointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteEndPoint(request *DeleteEndPointRequest) string {
	return c.DeleteEndPointWithContext(context.Background(), request)
}

func (c *Client) DeleteEndPointSend(request *DeleteEndPointRequest) (*DeleteEndPointResponse, error) {
	statusCode, msg, err := c.DeleteEndPointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteEndPointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteEndPointWithContext(ctx context.Context, request *DeleteEndPointRequest) string {
	if request == nil {
		request = NewDeleteEndPointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEndPointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteEndPointWithContextV2(ctx context.Context, request *DeleteEndPointRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteEndPointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEndPointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyEndPointRequest() (request *ModifyEndPointRequest) {
	request = &ModifyEndPointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "ModifyEndPoint")
	return
}

func NewModifyEndPointResponse() (response *ModifyEndPointResponse) {
	response = &ModifyEndPointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyEndPoint(request *ModifyEndPointRequest) string {
	return c.ModifyEndPointWithContext(context.Background(), request)
}

func (c *Client) ModifyEndPointSend(request *ModifyEndPointRequest) (*ModifyEndPointResponse, error) {
	statusCode, msg, err := c.ModifyEndPointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyEndPointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyEndPointWithContext(ctx context.Context, request *ModifyEndPointRequest) string {
	if request == nil {
		request = NewModifyEndPointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEndPointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyEndPointWithContextV2(ctx context.Context, request *ModifyEndPointRequest) (int, string, error) {
	if request == nil {
		request = NewModifyEndPointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEndPointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateEndPointRequest() (request *CreateEndPointRequest) {
	request = &CreateEndPointRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("pdns", APIVersion, "CreateEndPoint")
	return
}

func NewCreateEndPointResponse() (response *CreateEndPointResponse) {
	response = &CreateEndPointResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateEndPoint(request *CreateEndPointRequest) string {
	return c.CreateEndPointWithContext(context.Background(), request)
}

func (c *Client) CreateEndPointSend(request *CreateEndPointRequest) (*CreateEndPointResponse, error) {
	statusCode, msg, err := c.CreateEndPointWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateEndPointResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateEndPointWithContext(ctx context.Context, request *CreateEndPointRequest) string {
	if request == nil {
		request = NewCreateEndPointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEndPointResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateEndPointWithContextV2(ctx context.Context, request *CreateEndPointRequest) (int, string, error) {
	if request == nil {
		request = NewCreateEndPointRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEndPointResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


