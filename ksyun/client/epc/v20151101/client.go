package v20151101
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2015-11-01"

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

func NewCreateEpcRequest() (request *CreateEpcRequest) {
	request = &CreateEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateEpc")
	return
}

func NewCreateEpcResponse() (response *CreateEpcResponse) {
	response = &CreateEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateEpc(request *CreateEpcRequest) string {
	return c.CreateEpcWithContext(context.Background(), request)
}

func (c *Client) CreateEpcSend(request *CreateEpcRequest) (*CreateEpcResponse, error) {
	statusCode, msg, err := c.CreateEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateEpcWithContext(ctx context.Context, request *CreateEpcRequest) string {
	if request == nil {
		request = NewCreateEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateEpcWithContextV2(ctx context.Context, request *CreateEpcRequest) (int, string, error) {
	if request == nil {
		request = NewCreateEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartEpcRequest() (request *StartEpcRequest) {
	request = &StartEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "StartEpc")
	return
}

func NewStartEpcResponse() (response *StartEpcResponse) {
	response = &StartEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartEpc(request *StartEpcRequest) string {
	return c.StartEpcWithContext(context.Background(), request)
}

func (c *Client) StartEpcSend(request *StartEpcRequest) (*StartEpcResponse, error) {
	statusCode, msg, err := c.StartEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartEpcWithContext(ctx context.Context, request *StartEpcRequest) string {
	if request == nil {
		request = NewStartEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartEpcWithContextV2(ctx context.Context, request *StartEpcRequest) (int, string, error) {
	if request == nil {
		request = NewStartEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRebootEpcRequest() (request *RebootEpcRequest) {
	request = &RebootEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "RebootEpc")
	return
}

func NewRebootEpcResponse() (response *RebootEpcResponse) {
	response = &RebootEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootEpc(request *RebootEpcRequest) string {
	return c.RebootEpcWithContext(context.Background(), request)
}

func (c *Client) RebootEpcSend(request *RebootEpcRequest) (*RebootEpcResponse, error) {
	statusCode, msg, err := c.RebootEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RebootEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RebootEpcWithContext(ctx context.Context, request *RebootEpcRequest) string {
	if request == nil {
		request = NewRebootEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RebootEpcWithContextV2(ctx context.Context, request *RebootEpcRequest) (int, string, error) {
	if request == nil {
		request = NewRebootEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteEpcRequest() (request *DeleteEpcRequest) {
	request = &DeleteEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteEpc")
	return
}

func NewDeleteEpcResponse() (response *DeleteEpcResponse) {
	response = &DeleteEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteEpc(request *DeleteEpcRequest) string {
	return c.DeleteEpcWithContext(context.Background(), request)
}

func (c *Client) DeleteEpcSend(request *DeleteEpcRequest) (*DeleteEpcResponse, error) {
	statusCode, msg, err := c.DeleteEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteEpcWithContext(ctx context.Context, request *DeleteEpcRequest) string {
	if request == nil {
		request = NewDeleteEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteEpcWithContextV2(ctx context.Context, request *DeleteEpcRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReinstallEpcRequest() (request *ReinstallEpcRequest) {
	request = &ReinstallEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ReinstallEpc")
	return
}

func NewReinstallEpcResponse() (response *ReinstallEpcResponse) {
	response = &ReinstallEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReinstallEpc(request *ReinstallEpcRequest) string {
	return c.ReinstallEpcWithContext(context.Background(), request)
}

func (c *Client) ReinstallEpcSend(request *ReinstallEpcRequest) (*ReinstallEpcResponse, error) {
	statusCode, msg, err := c.ReinstallEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReinstallEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReinstallEpcWithContext(ctx context.Context, request *ReinstallEpcRequest) string {
	if request == nil {
		request = NewReinstallEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReinstallEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReinstallEpcWithContextV2(ctx context.Context, request *ReinstallEpcRequest) (int, string, error) {
	if request == nil {
		request = NewReinstallEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReinstallEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
	request = &ModifySecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifySecurityGroup")
	return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
	response = &ModifySecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) string {
	return c.ModifySecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupSend(request *ModifySecurityGroupRequest) (*ModifySecurityGroupResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupWithContext(ctx context.Context, request *ModifySecurityGroupRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupWithContextV2(ctx context.Context, request *ModifySecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateKeyRequest() (request *CreateKeyRequest) {
	request = &CreateKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateKey")
	return
}

func NewCreateKeyResponse() (response *CreateKeyResponse) {
	response = &CreateKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateKey(request *CreateKeyRequest) string {
	return c.CreateKeyWithContext(context.Background(), request)
}

func (c *Client) CreateKeySend(request *CreateKeyRequest) (*CreateKeyResponse, error) {
	statusCode, msg, err := c.CreateKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateKeyWithContext(ctx context.Context, request *CreateKeyRequest) string {
	if request == nil {
		request = NewCreateKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateKeyWithContextV2(ctx context.Context, request *CreateKeyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateKeyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcsRequest() (request *DescribeEpcsRequest) {
	request = &DescribeEpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcs")
	return
}

func NewDescribeEpcsResponse() (response *DescribeEpcsResponse) {
	response = &DescribeEpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcs(request *DescribeEpcsRequest) string {
	return c.DescribeEpcsWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcsSend(request *DescribeEpcsRequest) (*DescribeEpcsResponse, error) {
	statusCode, msg, err := c.DescribeEpcsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcsWithContext(ctx context.Context, request *DescribeEpcsRequest) string {
	if request == nil {
		request = NewDescribeEpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcsWithContextV2(ctx context.Context, request *DescribeEpcsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDynamicCodeRequest() (request *GetDynamicCodeRequest) {
	request = &GetDynamicCodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "GetDynamicCode")
	return
}

func NewGetDynamicCodeResponse() (response *GetDynamicCodeResponse) {
	response = &GetDynamicCodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDynamicCode(request *GetDynamicCodeRequest) string {
	return c.GetDynamicCodeWithContext(context.Background(), request)
}

func (c *Client) GetDynamicCodeSend(request *GetDynamicCodeRequest) (*GetDynamicCodeResponse, error) {
	statusCode, msg, err := c.GetDynamicCodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDynamicCodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDynamicCodeWithContext(ctx context.Context, request *GetDynamicCodeRequest) string {
	if request == nil {
		request = NewGetDynamicCodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDynamicCodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDynamicCodeWithContextV2(ctx context.Context, request *GetDynamicCodeRequest) (int, string, error) {
	if request == nil {
		request = NewGetDynamicCodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDynamicCodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVpnsRequest() (request *DescribeVpnsRequest) {
	request = &DescribeVpnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeVpns")
	return
}

func NewDescribeVpnsResponse() (response *DescribeVpnsResponse) {
	response = &DescribeVpnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpns(request *DescribeVpnsRequest) string {
	return c.DescribeVpnsWithContext(context.Background(), request)
}

func (c *Client) DescribeVpnsSend(request *DescribeVpnsRequest) (*DescribeVpnsResponse, error) {
	statusCode, msg, err := c.DescribeVpnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeVpnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVpnsWithContext(ctx context.Context, request *DescribeVpnsRequest) string {
	if request == nil {
		request = NewDescribeVpnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVpnsWithContextV2(ctx context.Context, request *DescribeVpnsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVpnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVpnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateImage(request *CreateImageRequest) string {
	return c.CreateImageWithContext(context.Background(), request)
}

func (c *Client) CreateImageSend(request *CreateImageRequest) (*CreateImageResponse, error) {
	statusCode, msg, err := c.CreateImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) string {
	if request == nil {
		request = NewCreateImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateImageWithContextV2(ctx context.Context, request *CreateImageRequest) (int, string, error) {
	if request == nil {
		request = NewCreateImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyImageRequest() (request *ModifyImageRequest) {
	request = &ModifyImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyImage")
	return
}

func NewModifyImageResponse() (response *ModifyImageResponse) {
	response = &ModifyImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyImage(request *ModifyImageRequest) string {
	return c.ModifyImageWithContext(context.Background(), request)
}

func (c *Client) ModifyImageSend(request *ModifyImageRequest) (*ModifyImageResponse, error) {
	statusCode, msg, err := c.ModifyImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyImageWithContext(ctx context.Context, request *ModifyImageRequest) string {
	if request == nil {
		request = NewModifyImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyImageWithContextV2(ctx context.Context, request *ModifyImageRequest) (int, string, error) {
	if request == nil {
		request = NewModifyImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteImage(request *DeleteImageRequest) string {
	return c.DeleteImageWithContext(context.Background(), request)
}

func (c *Client) DeleteImageSend(request *DeleteImageRequest) (*DeleteImageResponse, error) {
	statusCode, msg, err := c.DeleteImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) string {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteImageWithContextV2(ctx context.Context, request *DeleteImageRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) string {
	return c.DescribeImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeImagesSend(request *DescribeImagesRequest) (*DescribeImagesResponse, error) {
	statusCode, msg, err := c.DescribeImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) string {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeImagesWithContextV2(ctx context.Context, request *DescribeImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDnsRequest() (request *ModifyDnsRequest) {
	request = &ModifyDnsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyDns")
	return
}

func NewModifyDnsResponse() (response *ModifyDnsResponse) {
	response = &ModifyDnsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDns(request *ModifyDnsRequest) string {
	return c.ModifyDnsWithContext(context.Background(), request)
}

func (c *Client) ModifyDnsSend(request *ModifyDnsRequest) (*ModifyDnsResponse, error) {
	statusCode, msg, err := c.ModifyDnsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyDnsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDnsWithContext(ctx context.Context, request *ModifyDnsRequest) string {
	if request == nil {
		request = NewModifyDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDnsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDnsWithContextV2(ctx context.Context, request *ModifyDnsRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDnsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDnsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
	request = &ModifyNetworkInterfaceAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyNetworkInterfaceAttribute")
	return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
	response = &ModifyNetworkInterfaceAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) string {
	return c.ModifyNetworkInterfaceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyNetworkInterfaceAttributeSend(request *ModifyNetworkInterfaceAttributeRequest) (*ModifyNetworkInterfaceAttributeResponse, error) {
	statusCode, msg, err := c.ModifyNetworkInterfaceAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyNetworkInterfaceAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) string {
	if request == nil {
		request = NewModifyNetworkInterfaceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkInterfaceAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyNetworkInterfaceAttributeWithContextV2(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyNetworkInterfaceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyNetworkInterfaceAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePhysicalMonitorRequest() (request *DescribePhysicalMonitorRequest) {
	request = &DescribePhysicalMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribePhysicalMonitor")
	return
}

func NewDescribePhysicalMonitorResponse() (response *DescribePhysicalMonitorResponse) {
	response = &DescribePhysicalMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePhysicalMonitor(request *DescribePhysicalMonitorRequest) string {
	return c.DescribePhysicalMonitorWithContext(context.Background(), request)
}

func (c *Client) DescribePhysicalMonitorSend(request *DescribePhysicalMonitorRequest) (*DescribePhysicalMonitorResponse, error) {
	statusCode, msg, err := c.DescribePhysicalMonitorWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePhysicalMonitorResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePhysicalMonitorWithContext(ctx context.Context, request *DescribePhysicalMonitorRequest) string {
	if request == nil {
		request = NewDescribePhysicalMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePhysicalMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePhysicalMonitorWithContextV2(ctx context.Context, request *DescribePhysicalMonitorRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePhysicalMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePhysicalMonitorResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcManagementsRequest() (request *DescribeEpcManagementsRequest) {
	request = &DescribeEpcManagementsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcManagements")
	return
}

func NewDescribeEpcManagementsResponse() (response *DescribeEpcManagementsResponse) {
	response = &DescribeEpcManagementsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcManagements(request *DescribeEpcManagementsRequest) string {
	return c.DescribeEpcManagementsWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcManagementsSend(request *DescribeEpcManagementsRequest) (*DescribeEpcManagementsResponse, error) {
	statusCode, msg, err := c.DescribeEpcManagementsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcManagementsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcManagementsWithContext(ctx context.Context, request *DescribeEpcManagementsRequest) string {
	if request == nil {
		request = NewDescribeEpcManagementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcManagementsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcManagementsWithContextV2(ctx context.Context, request *DescribeEpcManagementsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcManagementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcManagementsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRemoteManagementsRequest() (request *DescribeRemoteManagementsRequest) {
	request = &DescribeRemoteManagementsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeRemoteManagements")
	return
}

func NewDescribeRemoteManagementsResponse() (response *DescribeRemoteManagementsResponse) {
	response = &DescribeRemoteManagementsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRemoteManagements(request *DescribeRemoteManagementsRequest) string {
	return c.DescribeRemoteManagementsWithContext(context.Background(), request)
}

func (c *Client) DescribeRemoteManagementsSend(request *DescribeRemoteManagementsRequest) (*DescribeRemoteManagementsResponse, error) {
	statusCode, msg, err := c.DescribeRemoteManagementsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRemoteManagementsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRemoteManagementsWithContext(ctx context.Context, request *DescribeRemoteManagementsRequest) string {
	if request == nil {
		request = NewDescribeRemoteManagementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRemoteManagementsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRemoteManagementsWithContextV2(ctx context.Context, request *DescribeRemoteManagementsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRemoteManagementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRemoteManagementsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopEpcRequest() (request *StopEpcRequest) {
	request = &StopEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "StopEpc")
	return
}

func NewStopEpcResponse() (response *StopEpcResponse) {
	response = &StopEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopEpc(request *StopEpcRequest) string {
	return c.StopEpcWithContext(context.Background(), request)
}

func (c *Client) StopEpcSend(request *StopEpcRequest) (*StopEpcResponse, error) {
	statusCode, msg, err := c.StopEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StopEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopEpcWithContext(ctx context.Context, request *StopEpcRequest) string {
	if request == nil {
		request = NewStopEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopEpcWithContextV2(ctx context.Context, request *StopEpcRequest) (int, string, error) {
	if request == nil {
		request = NewStopEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyEpcRequest() (request *ModifyEpcRequest) {
	request = &ModifyEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyEpc")
	return
}

func NewModifyEpcResponse() (response *ModifyEpcResponse) {
	response = &ModifyEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyEpc(request *ModifyEpcRequest) string {
	return c.ModifyEpcWithContext(context.Background(), request)
}

func (c *Client) ModifyEpcSend(request *ModifyEpcRequest) (*ModifyEpcResponse, error) {
	statusCode, msg, err := c.ModifyEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyEpcWithContext(ctx context.Context, request *ModifyEpcRequest) string {
	if request == nil {
		request = NewModifyEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyEpcWithContextV2(ctx context.Context, request *ModifyEpcRequest) (int, string, error) {
	if request == nil {
		request = NewModifyEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyRemoteManagementRequest() (request *ModifyRemoteManagementRequest) {
	request = &ModifyRemoteManagementRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyRemoteManagement")
	return
}

func NewModifyRemoteManagementResponse() (response *ModifyRemoteManagementResponse) {
	response = &ModifyRemoteManagementResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRemoteManagement(request *ModifyRemoteManagementRequest) string {
	return c.ModifyRemoteManagementWithContext(context.Background(), request)
}

func (c *Client) ModifyRemoteManagementSend(request *ModifyRemoteManagementRequest) (*ModifyRemoteManagementResponse, error) {
	statusCode, msg, err := c.ModifyRemoteManagementWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyRemoteManagementResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyRemoteManagementWithContext(ctx context.Context, request *ModifyRemoteManagementRequest) string {
	if request == nil {
		request = NewModifyRemoteManagementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRemoteManagementResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyRemoteManagementWithContextV2(ctx context.Context, request *ModifyRemoteManagementRequest) (int, string, error) {
	if request == nil {
		request = NewModifyRemoteManagementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRemoteManagementResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRemoteManagementRequest() (request *CreateRemoteManagementRequest) {
	request = &CreateRemoteManagementRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateRemoteManagement")
	return
}

func NewCreateRemoteManagementResponse() (response *CreateRemoteManagementResponse) {
	response = &CreateRemoteManagementResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRemoteManagement(request *CreateRemoteManagementRequest) string {
	return c.CreateRemoteManagementWithContext(context.Background(), request)
}

func (c *Client) CreateRemoteManagementSend(request *CreateRemoteManagementRequest) (*CreateRemoteManagementResponse, error) {
	statusCode, msg, err := c.CreateRemoteManagementWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRemoteManagementResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRemoteManagementWithContext(ctx context.Context, request *CreateRemoteManagementRequest) string {
	if request == nil {
		request = NewCreateRemoteManagementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRemoteManagementResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRemoteManagementWithContextV2(ctx context.Context, request *CreateRemoteManagementRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRemoteManagementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRemoteManagementResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReinstallCustomerEpcRequest() (request *ReinstallCustomerEpcRequest) {
	request = &ReinstallCustomerEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ReinstallCustomerEpc")
	return
}

func NewReinstallCustomerEpcResponse() (response *ReinstallCustomerEpcResponse) {
	response = &ReinstallCustomerEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReinstallCustomerEpc(request *ReinstallCustomerEpcRequest) string {
	return c.ReinstallCustomerEpcWithContext(context.Background(), request)
}

func (c *Client) ReinstallCustomerEpcSend(request *ReinstallCustomerEpcRequest) (*ReinstallCustomerEpcResponse, error) {
	statusCode, msg, err := c.ReinstallCustomerEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReinstallCustomerEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReinstallCustomerEpcWithContext(ctx context.Context, request *ReinstallCustomerEpcRequest) string {
	if request == nil {
		request = NewReinstallCustomerEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReinstallCustomerEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReinstallCustomerEpcWithContextV2(ctx context.Context, request *ReinstallCustomerEpcRequest) (int, string, error) {
	if request == nil {
		request = NewReinstallCustomerEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReinstallCustomerEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRemoteManagementRequest() (request *DeleteRemoteManagementRequest) {
	request = &DeleteRemoteManagementRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteRemoteManagement")
	return
}

func NewDeleteRemoteManagementResponse() (response *DeleteRemoteManagementResponse) {
	response = &DeleteRemoteManagementResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRemoteManagement(request *DeleteRemoteManagementRequest) string {
	return c.DeleteRemoteManagementWithContext(context.Background(), request)
}

func (c *Client) DeleteRemoteManagementSend(request *DeleteRemoteManagementRequest) (*DeleteRemoteManagementResponse, error) {
	statusCode, msg, err := c.DeleteRemoteManagementWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRemoteManagementResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRemoteManagementWithContext(ctx context.Context, request *DeleteRemoteManagementRequest) string {
	if request == nil {
		request = NewDeleteRemoteManagementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRemoteManagementResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRemoteManagementWithContextV2(ctx context.Context, request *DeleteRemoteManagementRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRemoteManagementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRemoteManagementResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResetPasswordRequest() (request *ResetPasswordRequest) {
	request = &ResetPasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ResetPassword")
	return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
	response = &ResetPasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetPassword(request *ResetPasswordRequest) string {
	return c.ResetPasswordWithContext(context.Background(), request)
}

func (c *Client) ResetPasswordSend(request *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	statusCode, msg, err := c.ResetPasswordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResetPasswordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) string {
	if request == nil {
		request = NewResetPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResetPasswordWithContextV2(ctx context.Context, request *ResetPasswordRequest) (int, string, error) {
	if request == nil {
		request = NewResetPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetPasswordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyHyperThreadingRequest() (request *ModifyHyperThreadingRequest) {
	request = &ModifyHyperThreadingRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyHyperThreading")
	return
}

func NewModifyHyperThreadingResponse() (response *ModifyHyperThreadingResponse) {
	response = &ModifyHyperThreadingResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyHyperThreading(request *ModifyHyperThreadingRequest) string {
	return c.ModifyHyperThreadingWithContext(context.Background(), request)
}

func (c *Client) ModifyHyperThreadingSend(request *ModifyHyperThreadingRequest) (*ModifyHyperThreadingResponse, error) {
	statusCode, msg, err := c.ModifyHyperThreadingWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyHyperThreadingResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyHyperThreadingWithContext(ctx context.Context, request *ModifyHyperThreadingRequest) string {
	if request == nil {
		request = NewModifyHyperThreadingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyHyperThreadingResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyHyperThreadingWithContextV2(ctx context.Context, request *ModifyHyperThreadingRequest) (int, string, error) {
	if request == nil {
		request = NewModifyHyperThreadingRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyHyperThreadingResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAssociateClusterRequest() (request *AssociateClusterRequest) {
	request = &AssociateClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "AssociateCluster")
	return
}

func NewAssociateClusterResponse() (response *AssociateClusterResponse) {
	response = &AssociateClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateCluster(request *AssociateClusterRequest) string {
	return c.AssociateClusterWithContext(context.Background(), request)
}

func (c *Client) AssociateClusterSend(request *AssociateClusterRequest) (*AssociateClusterResponse, error) {
	statusCode, msg, err := c.AssociateClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssociateClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AssociateClusterWithContext(ctx context.Context, request *AssociateClusterRequest) string {
	if request == nil {
		request = NewAssociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AssociateClusterWithContextV2(ctx context.Context, request *AssociateClusterRequest) (int, string, error) {
	if request == nil {
		request = NewAssociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDisassociateClusterRequest() (request *DisassociateClusterRequest) {
	request = &DisassociateClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DisassociateCluster")
	return
}

func NewDisassociateClusterResponse() (response *DisassociateClusterResponse) {
	response = &DisassociateClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateCluster(request *DisassociateClusterRequest) string {
	return c.DisassociateClusterWithContext(context.Background(), request)
}

func (c *Client) DisassociateClusterSend(request *DisassociateClusterRequest) (*DisassociateClusterResponse, error) {
	statusCode, msg, err := c.DisassociateClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DisassociateClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DisassociateClusterWithContext(ctx context.Context, request *DisassociateClusterRequest) string {
	if request == nil {
		request = NewDisassociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DisassociateClusterWithContextV2(ctx context.Context, request *DisassociateClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDisassociateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInspectionsRequest() (request *DescribeInspectionsRequest) {
	request = &DescribeInspectionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeInspections")
	return
}

func NewDescribeInspectionsResponse() (response *DescribeInspectionsResponse) {
	response = &DescribeInspectionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInspections(request *DescribeInspectionsRequest) string {
	return c.DescribeInspectionsWithContext(context.Background(), request)
}

func (c *Client) DescribeInspectionsSend(request *DescribeInspectionsRequest) (*DescribeInspectionsResponse, error) {
	statusCode, msg, err := c.DescribeInspectionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInspectionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInspectionsWithContext(ctx context.Context, request *DescribeInspectionsRequest) string {
	if request == nil {
		request = NewDescribeInspectionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInspectionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInspectionsWithContextV2(ctx context.Context, request *DescribeInspectionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInspectionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInspectionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcStocksRequest() (request *DescribeEpcStocksRequest) {
	request = &DescribeEpcStocksRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcStocks")
	return
}

func NewDescribeEpcStocksResponse() (response *DescribeEpcStocksResponse) {
	response = &DescribeEpcStocksResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcStocks(request *DescribeEpcStocksRequest) string {
	return c.DescribeEpcStocksWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcStocksSend(request *DescribeEpcStocksRequest) (*DescribeEpcStocksResponse, error) {
	statusCode, msg, err := c.DescribeEpcStocksWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcStocksResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcStocksWithContext(ctx context.Context, request *DescribeEpcStocksRequest) string {
	if request == nil {
		request = NewDescribeEpcStocksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcStocksResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcStocksWithContextV2(ctx context.Context, request *DescribeEpcStocksRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcStocksRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcStocksResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcDeviceAttributesRequest() (request *DescribeEpcDeviceAttributesRequest) {
	request = &DescribeEpcDeviceAttributesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcDeviceAttributes")
	return
}

func NewDescribeEpcDeviceAttributesResponse() (response *DescribeEpcDeviceAttributesResponse) {
	response = &DescribeEpcDeviceAttributesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcDeviceAttributes(request *DescribeEpcDeviceAttributesRequest) string {
	return c.DescribeEpcDeviceAttributesWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcDeviceAttributesSend(request *DescribeEpcDeviceAttributesRequest) (*DescribeEpcDeviceAttributesResponse, error) {
	statusCode, msg, err := c.DescribeEpcDeviceAttributesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcDeviceAttributesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcDeviceAttributesWithContext(ctx context.Context, request *DescribeEpcDeviceAttributesRequest) string {
	if request == nil {
		request = NewDescribeEpcDeviceAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcDeviceAttributesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcDeviceAttributesWithContextV2(ctx context.Context, request *DescribeEpcDeviceAttributesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcDeviceAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcDeviceAttributesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
	request = &DescribeProcessesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeProcesses")
	return
}

func NewDescribeProcessesResponse() (response *DescribeProcessesResponse) {
	response = &DescribeProcessesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeProcesses(request *DescribeProcessesRequest) string {
	return c.DescribeProcessesWithContext(context.Background(), request)
}

func (c *Client) DescribeProcessesSend(request *DescribeProcessesRequest) (*DescribeProcessesResponse, error) {
	statusCode, msg, err := c.DescribeProcessesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeProcessesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeProcessesWithContext(ctx context.Context, request *DescribeProcessesRequest) string {
	if request == nil {
		request = NewDescribeProcessesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeProcessesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeProcessesWithContextV2(ctx context.Context, request *DescribeProcessesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeProcessesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeProcessesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateProcessRequest() (request *CreateProcessRequest) {
	request = &CreateProcessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateProcess")
	return
}

func NewCreateProcessResponse() (response *CreateProcessResponse) {
	response = &CreateProcessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateProcess(request *CreateProcessRequest) string {
	return c.CreateProcessWithContext(context.Background(), request)
}

func (c *Client) CreateProcessSend(request *CreateProcessRequest) (*CreateProcessResponse, error) {
	statusCode, msg, err := c.CreateProcessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateProcessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateProcessWithContext(ctx context.Context, request *CreateProcessRequest) string {
	if request == nil {
		request = NewCreateProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateProcessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateProcessWithContextV2(ctx context.Context, request *CreateProcessRequest) (int, string, error) {
	if request == nil {
		request = NewCreateProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateProcessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteProcessRequest() (request *DeleteProcessRequest) {
	request = &DeleteProcessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteProcess")
	return
}

func NewDeleteProcessResponse() (response *DeleteProcessResponse) {
	response = &DeleteProcessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteProcess(request *DeleteProcessRequest) string {
	return c.DeleteProcessWithContext(context.Background(), request)
}

func (c *Client) DeleteProcessSend(request *DeleteProcessRequest) (*DeleteProcessResponse, error) {
	statusCode, msg, err := c.DeleteProcessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteProcessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteProcessWithContext(ctx context.Context, request *DeleteProcessRequest) string {
	if request == nil {
		request = NewDeleteProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteProcessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteProcessWithContextV2(ctx context.Context, request *DeleteProcessRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteProcessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReplyProcessRequest() (request *ReplyProcessRequest) {
	request = &ReplyProcessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ReplyProcess")
	return
}

func NewReplyProcessResponse() (response *ReplyProcessResponse) {
	response = &ReplyProcessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReplyProcess(request *ReplyProcessRequest) string {
	return c.ReplyProcessWithContext(context.Background(), request)
}

func (c *Client) ReplyProcessSend(request *ReplyProcessRequest) (*ReplyProcessResponse, error) {
	statusCode, msg, err := c.ReplyProcessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReplyProcessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReplyProcessWithContext(ctx context.Context, request *ReplyProcessRequest) string {
	if request == nil {
		request = NewReplyProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReplyProcessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReplyProcessWithContextV2(ctx context.Context, request *ReplyProcessRequest) (int, string, error) {
	if request == nil {
		request = NewReplyProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReplyProcessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcTrashesRequest() (request *DescribeEpcTrashesRequest) {
	request = &DescribeEpcTrashesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcTrashes")
	return
}

func NewDescribeEpcTrashesResponse() (response *DescribeEpcTrashesResponse) {
	response = &DescribeEpcTrashesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcTrashes(request *DescribeEpcTrashesRequest) string {
	return c.DescribeEpcTrashesWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcTrashesSend(request *DescribeEpcTrashesRequest) (*DescribeEpcTrashesResponse, error) {
	statusCode, msg, err := c.DescribeEpcTrashesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcTrashesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcTrashesWithContext(ctx context.Context, request *DescribeEpcTrashesRequest) string {
	if request == nil {
		request = NewDescribeEpcTrashesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcTrashesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcTrashesWithContextV2(ctx context.Context, request *DescribeEpcTrashesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcTrashesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcTrashesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReturnEpcRequest() (request *ReturnEpcRequest) {
	request = &ReturnEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ReturnEpc")
	return
}

func NewReturnEpcResponse() (response *ReturnEpcResponse) {
	response = &ReturnEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReturnEpc(request *ReturnEpcRequest) string {
	return c.ReturnEpcWithContext(context.Background(), request)
}

func (c *Client) ReturnEpcSend(request *ReturnEpcRequest) (*ReturnEpcResponse, error) {
	statusCode, msg, err := c.ReturnEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReturnEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReturnEpcWithContext(ctx context.Context, request *ReturnEpcRequest) string {
	if request == nil {
		request = NewReturnEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReturnEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReturnEpcWithContextV2(ctx context.Context, request *ReturnEpcRequest) (int, string, error) {
	if request == nil {
		request = NewReturnEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReturnEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateResourceRequirementRequest() (request *CreateResourceRequirementRequest) {
	request = &CreateResourceRequirementRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateResourceRequirement")
	return
}

func NewCreateResourceRequirementResponse() (response *CreateResourceRequirementResponse) {
	response = &CreateResourceRequirementResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateResourceRequirement(request *CreateResourceRequirementRequest) string {
	return c.CreateResourceRequirementWithContext(context.Background(), request)
}

func (c *Client) CreateResourceRequirementSend(request *CreateResourceRequirementRequest) (*CreateResourceRequirementResponse, error) {
	statusCode, msg, err := c.CreateResourceRequirementWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateResourceRequirementResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateResourceRequirementWithContext(ctx context.Context, request *CreateResourceRequirementRequest) string {
	if request == nil {
		request = NewCreateResourceRequirementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateResourceRequirementResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateResourceRequirementWithContextV2(ctx context.Context, request *CreateResourceRequirementRequest) (int, string, error) {
	if request == nil {
		request = NewCreateResourceRequirementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateResourceRequirementResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachVolumeRequest() (request *AttachVolumeRequest) {
	request = &AttachVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "AttachVolume")
	return
}

func NewAttachVolumeResponse() (response *AttachVolumeResponse) {
	response = &AttachVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachVolume(request *AttachVolumeRequest) string {
	return c.AttachVolumeWithContext(context.Background(), request)
}

func (c *Client) AttachVolumeSend(request *AttachVolumeRequest) (*AttachVolumeResponse, error) {
	statusCode, msg, err := c.AttachVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AttachVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachVolumeWithContext(ctx context.Context, request *AttachVolumeRequest) string {
	if request == nil {
		request = NewAttachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachVolumeWithContextV2(ctx context.Context, request *AttachVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewAttachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachVolumeRequest() (request *DetachVolumeRequest) {
	request = &DetachVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DetachVolume")
	return
}

func NewDetachVolumeResponse() (response *DetachVolumeResponse) {
	response = &DetachVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachVolume(request *DetachVolumeRequest) string {
	return c.DetachVolumeWithContext(context.Background(), request)
}

func (c *Client) DetachVolumeSend(request *DetachVolumeRequest) (*DetachVolumeResponse, error) {
	statusCode, msg, err := c.DetachVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachVolumeWithContext(ctx context.Context, request *DetachVolumeRequest) string {
	if request == nil {
		request = NewDetachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachVolumeWithContextV2(ctx context.Context, request *DetachVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewDetachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePrice(request *DescribePriceRequest) string {
	return c.DescribePriceWithContext(context.Background(), request)
}

func (c *Client) DescribePriceSend(request *DescribePriceRequest) (*DescribePriceResponse, error) {
	statusCode, msg, err := c.DescribePriceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePriceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest) string {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePriceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePriceWithContextV2(ctx context.Context, request *DescribePriceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePriceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateConfirmRequest() (request *UpdateConfirmRequest) {
	request = &UpdateConfirmRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "UpdateConfirm")
	return
}

func NewUpdateConfirmResponse() (response *UpdateConfirmResponse) {
	response = &UpdateConfirmResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateConfirm(request *UpdateConfirmRequest) string {
	return c.UpdateConfirmWithContext(context.Background(), request)
}

func (c *Client) UpdateConfirmSend(request *UpdateConfirmRequest) (*UpdateConfirmResponse, error) {
	statusCode, msg, err := c.UpdateConfirmWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateConfirmResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateConfirmWithContext(ctx context.Context, request *UpdateConfirmRequest) string {
	if request == nil {
		request = NewUpdateConfirmRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateConfirmResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateConfirmWithContextV2(ctx context.Context, request *UpdateConfirmRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateConfirmRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateConfirmResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyOverclockingAttributeRequest() (request *ModifyOverclockingAttributeRequest) {
	request = &ModifyOverclockingAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyOverclockingAttribute")
	return
}

func NewModifyOverclockingAttributeResponse() (response *ModifyOverclockingAttributeResponse) {
	response = &ModifyOverclockingAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyOverclockingAttribute(request *ModifyOverclockingAttributeRequest) string {
	return c.ModifyOverclockingAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyOverclockingAttributeSend(request *ModifyOverclockingAttributeRequest) (*ModifyOverclockingAttributeResponse, error) {
	statusCode, msg, err := c.ModifyOverclockingAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyOverclockingAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyOverclockingAttributeWithContext(ctx context.Context, request *ModifyOverclockingAttributeRequest) string {
	if request == nil {
		request = NewModifyOverclockingAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyOverclockingAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyOverclockingAttributeWithContextV2(ctx context.Context, request *ModifyOverclockingAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyOverclockingAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyOverclockingAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCopyImageRequest() (request *CopyImageRequest) {
	request = &CopyImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CopyImage")
	return
}

func NewCopyImageResponse() (response *CopyImageResponse) {
	response = &CopyImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CopyImage(request *CopyImageRequest) string {
	return c.CopyImageWithContext(context.Background(), request)
}

func (c *Client) CopyImageSend(request *CopyImageRequest) (*CopyImageResponse, error) {
	statusCode, msg, err := c.CopyImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CopyImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CopyImageWithContext(ctx context.Context, request *CopyImageRequest) string {
	if request == nil {
		request = NewCopyImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopyImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CopyImageWithContextV2(ctx context.Context, request *CopyImageRequest) (int, string, error) {
	if request == nil {
		request = NewCopyImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCopyImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEpcRaidAttributesRequest() (request *DescribeEpcRaidAttributesRequest) {
	request = &DescribeEpcRaidAttributesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeEpcRaidAttributes")
	return
}

func NewDescribeEpcRaidAttributesResponse() (response *DescribeEpcRaidAttributesResponse) {
	response = &DescribeEpcRaidAttributesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEpcRaidAttributes(request *DescribeEpcRaidAttributesRequest) string {
	return c.DescribeEpcRaidAttributesWithContext(context.Background(), request)
}

func (c *Client) DescribeEpcRaidAttributesSend(request *DescribeEpcRaidAttributesRequest) (*DescribeEpcRaidAttributesResponse, error) {
	statusCode, msg, err := c.DescribeEpcRaidAttributesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEpcRaidAttributesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEpcRaidAttributesWithContext(ctx context.Context, request *DescribeEpcRaidAttributesRequest) string {
	if request == nil {
		request = NewDescribeEpcRaidAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcRaidAttributesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEpcRaidAttributesWithContextV2(ctx context.Context, request *DescribeEpcRaidAttributesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEpcRaidAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEpcRaidAttributesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeGpuImageDriverRequest() (request *DescribeGpuImageDriverRequest) {
	request = &DescribeGpuImageDriverRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeGpuImageDriver")
	return
}

func NewDescribeGpuImageDriverResponse() (response *DescribeGpuImageDriverResponse) {
	response = &DescribeGpuImageDriverResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeGpuImageDriver(request *DescribeGpuImageDriverRequest) string {
	return c.DescribeGpuImageDriverWithContext(context.Background(), request)
}

func (c *Client) DescribeGpuImageDriverSend(request *DescribeGpuImageDriverRequest) (*DescribeGpuImageDriverResponse, error) {
	statusCode, msg, err := c.DescribeGpuImageDriverWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeGpuImageDriverResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeGpuImageDriverWithContext(ctx context.Context, request *DescribeGpuImageDriverRequest) string {
	if request == nil {
		request = NewDescribeGpuImageDriverRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeGpuImageDriverResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeGpuImageDriverWithContextV2(ctx context.Context, request *DescribeGpuImageDriverRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeGpuImageDriverRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeGpuImageDriverResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateShareImageRequest() (request *CreateShareImageRequest) {
	request = &CreateShareImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateShareImage")
	return
}

func NewCreateShareImageResponse() (response *CreateShareImageResponse) {
	response = &CreateShareImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateShareImage(request *CreateShareImageRequest) string {
	return c.CreateShareImageWithContext(context.Background(), request)
}

func (c *Client) CreateShareImageSend(request *CreateShareImageRequest) (*CreateShareImageResponse, error) {
	statusCode, msg, err := c.CreateShareImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateShareImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateShareImageWithContext(ctx context.Context, request *CreateShareImageRequest) string {
	if request == nil {
		request = NewCreateShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateShareImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateShareImageWithContextV2(ctx context.Context, request *CreateShareImageRequest) (int, string, error) {
	if request == nil {
		request = NewCreateShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateShareImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteShareImageRequest() (request *DeleteShareImageRequest) {
	request = &DeleteShareImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteShareImage")
	return
}

func NewDeleteShareImageResponse() (response *DeleteShareImageResponse) {
	response = &DeleteShareImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteShareImage(request *DeleteShareImageRequest) string {
	return c.DeleteShareImageWithContext(context.Background(), request)
}

func (c *Client) DeleteShareImageSend(request *DeleteShareImageRequest) (*DeleteShareImageResponse, error) {
	statusCode, msg, err := c.DeleteShareImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteShareImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteShareImageWithContext(ctx context.Context, request *DeleteShareImageRequest) string {
	if request == nil {
		request = NewDeleteShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteShareImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteShareImageWithContextV2(ctx context.Context, request *DeleteShareImageRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteShareImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeShareImageAccountListRequest() (request *DescribeShareImageAccountListRequest) {
	request = &DescribeShareImageAccountListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeShareImageAccountList")
	return
}

func NewDescribeShareImageAccountListResponse() (response *DescribeShareImageAccountListResponse) {
	response = &DescribeShareImageAccountListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeShareImageAccountList(request *DescribeShareImageAccountListRequest) string {
	return c.DescribeShareImageAccountListWithContext(context.Background(), request)
}

func (c *Client) DescribeShareImageAccountListSend(request *DescribeShareImageAccountListRequest) (*DescribeShareImageAccountListResponse, error) {
	statusCode, msg, err := c.DescribeShareImageAccountListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeShareImageAccountListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeShareImageAccountListWithContext(ctx context.Context, request *DescribeShareImageAccountListRequest) string {
	if request == nil {
		request = NewDescribeShareImageAccountListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeShareImageAccountListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeShareImageAccountListWithContextV2(ctx context.Context, request *DescribeShareImageAccountListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeShareImageAccountListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeShareImageAccountListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeShareImageRequest() (request *DescribeShareImageRequest) {
	request = &DescribeShareImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeShareImage")
	return
}

func NewDescribeShareImageResponse() (response *DescribeShareImageResponse) {
	response = &DescribeShareImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeShareImage(request *DescribeShareImageRequest) string {
	return c.DescribeShareImageWithContext(context.Background(), request)
}

func (c *Client) DescribeShareImageSend(request *DescribeShareImageRequest) (*DescribeShareImageResponse, error) {
	statusCode, msg, err := c.DescribeShareImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeShareImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeShareImageWithContext(ctx context.Context, request *DescribeShareImageRequest) string {
	if request == nil {
		request = NewDescribeShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeShareImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeShareImageWithContextV2(ctx context.Context, request *DescribeShareImageRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeShareImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAcceptShareImageRequest() (request *AcceptShareImageRequest) {
	request = &AcceptShareImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "AcceptShareImage")
	return
}

func NewAcceptShareImageResponse() (response *AcceptShareImageResponse) {
	response = &AcceptShareImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AcceptShareImage(request *AcceptShareImageRequest) string {
	return c.AcceptShareImageWithContext(context.Background(), request)
}

func (c *Client) AcceptShareImageSend(request *AcceptShareImageRequest) (*AcceptShareImageResponse, error) {
	statusCode, msg, err := c.AcceptShareImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AcceptShareImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AcceptShareImageWithContext(ctx context.Context, request *AcceptShareImageRequest) string {
	if request == nil {
		request = NewAcceptShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAcceptShareImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AcceptShareImageWithContextV2(ctx context.Context, request *AcceptShareImageRequest) (int, string, error) {
	if request == nil {
		request = NewAcceptShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAcceptShareImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRejectShareImageRequest() (request *RejectShareImageRequest) {
	request = &RejectShareImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "RejectShareImage")
	return
}

func NewRejectShareImageResponse() (response *RejectShareImageResponse) {
	response = &RejectShareImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RejectShareImage(request *RejectShareImageRequest) string {
	return c.RejectShareImageWithContext(context.Background(), request)
}

func (c *Client) RejectShareImageSend(request *RejectShareImageRequest) (*RejectShareImageResponse, error) {
	statusCode, msg, err := c.RejectShareImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RejectShareImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RejectShareImageWithContext(ctx context.Context, request *RejectShareImageRequest) string {
	if request == nil {
		request = NewRejectShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRejectShareImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RejectShareImageWithContextV2(ctx context.Context, request *RejectShareImageRequest) (int, string, error) {
	if request == nil {
		request = NewRejectShareImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRejectShareImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeManagedAccessoryRequest() (request *DescribeManagedAccessoryRequest) {
	request = &DescribeManagedAccessoryRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeManagedAccessory")
	return
}

func NewDescribeManagedAccessoryResponse() (response *DescribeManagedAccessoryResponse) {
	response = &DescribeManagedAccessoryResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeManagedAccessory(request *DescribeManagedAccessoryRequest) string {
	return c.DescribeManagedAccessoryWithContext(context.Background(), request)
}

func (c *Client) DescribeManagedAccessorySend(request *DescribeManagedAccessoryRequest) (*DescribeManagedAccessoryResponse, error) {
	statusCode, msg, err := c.DescribeManagedAccessoryWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeManagedAccessoryResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeManagedAccessoryWithContext(ctx context.Context, request *DescribeManagedAccessoryRequest) string {
	if request == nil {
		request = NewDescribeManagedAccessoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeManagedAccessoryResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeManagedAccessoryWithContextV2(ctx context.Context, request *DescribeManagedAccessoryRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeManagedAccessoryRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeManagedAccessoryResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAutoDeleteEpcRequest() (request *AutoDeleteEpcRequest) {
	request = &AutoDeleteEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "AutoDeleteEpc")
	return
}

func NewAutoDeleteEpcResponse() (response *AutoDeleteEpcResponse) {
	response = &AutoDeleteEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AutoDeleteEpc(request *AutoDeleteEpcRequest) string {
	return c.AutoDeleteEpcWithContext(context.Background(), request)
}

func (c *Client) AutoDeleteEpcSend(request *AutoDeleteEpcRequest) (*AutoDeleteEpcResponse, error) {
	statusCode, msg, err := c.AutoDeleteEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AutoDeleteEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AutoDeleteEpcWithContext(ctx context.Context, request *AutoDeleteEpcRequest) string {
	if request == nil {
		request = NewAutoDeleteEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAutoDeleteEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AutoDeleteEpcWithContextV2(ctx context.Context, request *AutoDeleteEpcRequest) (int, string, error) {
	if request == nil {
		request = NewAutoDeleteEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAutoDeleteEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewExportImageRequest() (request *ExportImageRequest) {
	request = &ExportImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ExportImage")
	return
}

func NewExportImageResponse() (response *ExportImageResponse) {
	response = &ExportImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ExportImage(request *ExportImageRequest) string {
	return c.ExportImageWithContext(context.Background(), request)
}

func (c *Client) ExportImageSend(request *ExportImageRequest) (*ExportImageResponse, error) {
	statusCode, msg, err := c.ExportImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ExportImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ExportImageWithContext(ctx context.Context, request *ExportImageRequest) string {
	if request == nil {
		request = NewExportImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewExportImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ExportImageWithContextV2(ctx context.Context, request *ExportImageRequest) (int, string, error) {
	if request == nil {
		request = NewExportImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewExportImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewQueryBucketsRequest() (request *QueryBucketsRequest) {
	request = &QueryBucketsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "QueryBuckets")
	return
}

func NewQueryBucketsResponse() (response *QueryBucketsResponse) {
	response = &QueryBucketsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryBuckets(request *QueryBucketsRequest) string {
	return c.QueryBucketsWithContext(context.Background(), request)
}

func (c *Client) QueryBucketsSend(request *QueryBucketsRequest) (*QueryBucketsResponse, error) {
	statusCode, msg, err := c.QueryBucketsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryBucketsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryBucketsWithContext(ctx context.Context, request *QueryBucketsRequest) string {
	if request == nil {
		request = NewQueryBucketsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryBucketsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryBucketsWithContextV2(ctx context.Context, request *QueryBucketsRequest) (int, string, error) {
	if request == nil {
		request = NewQueryBucketsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryBucketsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelImageExportRequest() (request *CancelImageExportRequest) {
	request = &CancelImageExportRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CancelImageExport")
	return
}

func NewCancelImageExportResponse() (response *CancelImageExportResponse) {
	response = &CancelImageExportResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelImageExport(request *CancelImageExportRequest) string {
	return c.CancelImageExportWithContext(context.Background(), request)
}

func (c *Client) CancelImageExportSend(request *CancelImageExportRequest) (*CancelImageExportResponse, error) {
	statusCode, msg, err := c.CancelImageExportWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelImageExportResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelImageExportWithContext(ctx context.Context, request *CancelImageExportRequest) string {
	if request == nil {
		request = NewCancelImageExportRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelImageExportResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelImageExportWithContextV2(ctx context.Context, request *CancelImageExportRequest) (int, string, error) {
	if request == nil {
		request = NewCancelImageExportRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelImageExportResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUseHotStandByEpcRequest() (request *UseHotStandByEpcRequest) {
	request = &UseHotStandByEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "UseHotStandByEpc")
	return
}

func NewUseHotStandByEpcResponse() (response *UseHotStandByEpcResponse) {
	response = &UseHotStandByEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UseHotStandByEpc(request *UseHotStandByEpcRequest) string {
	return c.UseHotStandByEpcWithContext(context.Background(), request)
}

func (c *Client) UseHotStandByEpcSend(request *UseHotStandByEpcRequest) (*UseHotStandByEpcResponse, error) {
	statusCode, msg, err := c.UseHotStandByEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UseHotStandByEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UseHotStandByEpcWithContext(ctx context.Context, request *UseHotStandByEpcRequest) string {
	if request == nil {
		request = NewUseHotStandByEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUseHotStandByEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UseHotStandByEpcWithContextV2(ctx context.Context, request *UseHotStandByEpcRequest) (int, string, error) {
	if request == nil {
		request = NewUseHotStandByEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUseHotStandByEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewActivateHotStandbyEpcRequest() (request *ActivateHotStandbyEpcRequest) {
	request = &ActivateHotStandbyEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ActivateHotStandbyEpc")
	return
}

func NewActivateHotStandbyEpcResponse() (response *ActivateHotStandbyEpcResponse) {
	response = &ActivateHotStandbyEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ActivateHotStandbyEpc(request *ActivateHotStandbyEpcRequest) string {
	return c.ActivateHotStandbyEpcWithContext(context.Background(), request)
}

func (c *Client) ActivateHotStandbyEpcSend(request *ActivateHotStandbyEpcRequest) (*ActivateHotStandbyEpcResponse, error) {
	statusCode, msg, err := c.ActivateHotStandbyEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ActivateHotStandbyEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ActivateHotStandbyEpcWithContext(ctx context.Context, request *ActivateHotStandbyEpcRequest) string {
	if request == nil {
		request = NewActivateHotStandbyEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewActivateHotStandbyEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ActivateHotStandbyEpcWithContextV2(ctx context.Context, request *ActivateHotStandbyEpcRequest) (int, string, error) {
	if request == nil {
		request = NewActivateHotStandbyEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewActivateHotStandbyEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchCreateEpcRequest() (request *BatchCreateEpcRequest) {
	request = &BatchCreateEpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "BatchCreateEpc")
	return
}

func NewBatchCreateEpcResponse() (response *BatchCreateEpcResponse) {
	response = &BatchCreateEpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchCreateEpc(request *BatchCreateEpcRequest) string {
	return c.BatchCreateEpcWithContext(context.Background(), request)
}

func (c *Client) BatchCreateEpcSend(request *BatchCreateEpcRequest) (*BatchCreateEpcResponse, error) {
	statusCode, msg, err := c.BatchCreateEpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct BatchCreateEpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchCreateEpcWithContext(ctx context.Context, request *BatchCreateEpcRequest) string {
	if request == nil {
		request = NewBatchCreateEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchCreateEpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchCreateEpcWithContextV2(ctx context.Context, request *BatchCreateEpcRequest) (int, string, error) {
	if request == nil {
		request = NewBatchCreateEpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchCreateEpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeUseHotStandbyRecordsRequest() (request *DescribeUseHotStandbyRecordsRequest) {
	request = &DescribeUseHotStandbyRecordsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeUseHotStandbyRecords")
	return
}

func NewDescribeUseHotStandbyRecordsResponse() (response *DescribeUseHotStandbyRecordsResponse) {
	response = &DescribeUseHotStandbyRecordsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeUseHotStandbyRecords(request *DescribeUseHotStandbyRecordsRequest) string {
	return c.DescribeUseHotStandbyRecordsWithContext(context.Background(), request)
}

func (c *Client) DescribeUseHotStandbyRecordsSend(request *DescribeUseHotStandbyRecordsRequest) (*DescribeUseHotStandbyRecordsResponse, error) {
	statusCode, msg, err := c.DescribeUseHotStandbyRecordsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeUseHotStandbyRecordsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeUseHotStandbyRecordsWithContext(ctx context.Context, request *DescribeUseHotStandbyRecordsRequest) string {
	if request == nil {
		request = NewDescribeUseHotStandbyRecordsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUseHotStandbyRecordsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeUseHotStandbyRecordsWithContextV2(ctx context.Context, request *DescribeUseHotStandbyRecordsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeUseHotStandbyRecordsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeUseHotStandbyRecordsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeGpuRoceTopologyRequest() (request *DescribeGpuRoceTopologyRequest) {
	request = &DescribeGpuRoceTopologyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeGpuRoceTopology")
	return
}

func NewDescribeGpuRoceTopologyResponse() (response *DescribeGpuRoceTopologyResponse) {
	response = &DescribeGpuRoceTopologyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeGpuRoceTopology(request *DescribeGpuRoceTopologyRequest) string {
	return c.DescribeGpuRoceTopologyWithContext(context.Background(), request)
}

func (c *Client) DescribeGpuRoceTopologySend(request *DescribeGpuRoceTopologyRequest) (*DescribeGpuRoceTopologyResponse, error) {
	statusCode, msg, err := c.DescribeGpuRoceTopologyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeGpuRoceTopologyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeGpuRoceTopologyWithContext(ctx context.Context, request *DescribeGpuRoceTopologyRequest) string {
	if request == nil {
		request = NewDescribeGpuRoceTopologyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeGpuRoceTopologyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeGpuRoceTopologyWithContextV2(ctx context.Context, request *DescribeGpuRoceTopologyRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeGpuRoceTopologyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeGpuRoceTopologyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyProcessRequest() (request *ModifyProcessRequest) {
	request = &ModifyProcessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifyProcess")
	return
}

func NewModifyProcessResponse() (response *ModifyProcessResponse) {
	response = &ModifyProcessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyProcess(request *ModifyProcessRequest) string {
	return c.ModifyProcessWithContext(context.Background(), request)
}

func (c *Client) ModifyProcessSend(request *ModifyProcessRequest) (*ModifyProcessResponse, error) {
	statusCode, msg, err := c.ModifyProcessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyProcessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyProcessWithContext(ctx context.Context, request *ModifyProcessRequest) string {
	if request == nil {
		request = NewModifyProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyProcessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyProcessWithContextV2(ctx context.Context, request *ModifyProcessRequest) (int, string, error) {
	if request == nil {
		request = NewModifyProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyProcessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewConfirmProcessRequest() (request *ConfirmProcessRequest) {
	request = &ConfirmProcessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ConfirmProcess")
	return
}

func NewConfirmProcessResponse() (response *ConfirmProcessResponse) {
	response = &ConfirmProcessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ConfirmProcess(request *ConfirmProcessRequest) string {
	return c.ConfirmProcessWithContext(context.Background(), request)
}

func (c *Client) ConfirmProcessSend(request *ConfirmProcessRequest) (*ConfirmProcessResponse, error) {
	statusCode, msg, err := c.ConfirmProcessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ConfirmProcessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ConfirmProcessWithContext(ctx context.Context, request *ConfirmProcessRequest) string {
	if request == nil {
		request = NewConfirmProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewConfirmProcessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ConfirmProcessWithContextV2(ctx context.Context, request *ConfirmProcessRequest) (int, string, error) {
	if request == nil {
		request = NewConfirmProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewConfirmProcessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeModelConfigRequest() (request *DescribeModelConfigRequest) {
	request = &DescribeModelConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeModelConfig")
	return
}

func NewDescribeModelConfigResponse() (response *DescribeModelConfigResponse) {
	response = &DescribeModelConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeModelConfig(request *DescribeModelConfigRequest) string {
	return c.DescribeModelConfigWithContext(context.Background(), request)
}

func (c *Client) DescribeModelConfigSend(request *DescribeModelConfigRequest) (*DescribeModelConfigResponse, error) {
	statusCode, msg, err := c.DescribeModelConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeModelConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeModelConfigWithContext(ctx context.Context, request *DescribeModelConfigRequest) string {
	if request == nil {
		request = NewDescribeModelConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeModelConfigWithContextV2(ctx context.Context, request *DescribeModelConfigRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeModelConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeModelConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRoceEventRequest() (request *DescribeRoceEventRequest) {
	request = &DescribeRoceEventRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeRoceEvent")
	return
}

func NewDescribeRoceEventResponse() (response *DescribeRoceEventResponse) {
	response = &DescribeRoceEventResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRoceEvent(request *DescribeRoceEventRequest) string {
	return c.DescribeRoceEventWithContext(context.Background(), request)
}

func (c *Client) DescribeRoceEventSend(request *DescribeRoceEventRequest) (*DescribeRoceEventResponse, error) {
	statusCode, msg, err := c.DescribeRoceEventWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRoceEventResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRoceEventWithContext(ctx context.Context, request *DescribeRoceEventRequest) string {
	if request == nil {
		request = NewDescribeRoceEventRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRoceEventResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRoceEventWithContextV2(ctx context.Context, request *DescribeRoceEventRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRoceEventRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRoceEventResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRoceEventDetailsRequest() (request *DescribeRoceEventDetailsRequest) {
	request = &DescribeRoceEventDetailsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeRoceEventDetails")
	return
}

func NewDescribeRoceEventDetailsResponse() (response *DescribeRoceEventDetailsResponse) {
	response = &DescribeRoceEventDetailsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRoceEventDetails(request *DescribeRoceEventDetailsRequest) string {
	return c.DescribeRoceEventDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeRoceEventDetailsSend(request *DescribeRoceEventDetailsRequest) (*DescribeRoceEventDetailsResponse, error) {
	statusCode, msg, err := c.DescribeRoceEventDetailsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRoceEventDetailsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRoceEventDetailsWithContext(ctx context.Context, request *DescribeRoceEventDetailsRequest) string {
	if request == nil {
		request = NewDescribeRoceEventDetailsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRoceEventDetailsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRoceEventDetailsWithContextV2(ctx context.Context, request *DescribeRoceEventDetailsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRoceEventDetailsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRoceEventDetailsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewBatchCreateProcessRequest() (request *BatchCreateProcessRequest) {
	request = &BatchCreateProcessRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "BatchCreateProcess")
	return
}

func NewBatchCreateProcessResponse() (response *BatchCreateProcessResponse) {
	response = &BatchCreateProcessResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) BatchCreateProcess(request *BatchCreateProcessRequest) string {
	return c.BatchCreateProcessWithContext(context.Background(), request)
}

func (c *Client) BatchCreateProcessSend(request *BatchCreateProcessRequest) (*BatchCreateProcessResponse, error) {
	statusCode, msg, err := c.BatchCreateProcessWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct BatchCreateProcessResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) BatchCreateProcessWithContext(ctx context.Context, request *BatchCreateProcessRequest) string {
	if request == nil {
		request = NewBatchCreateProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchCreateProcessResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) BatchCreateProcessWithContextV2(ctx context.Context, request *BatchCreateProcessRequest) (int, string, error) {
	if request == nil {
		request = NewBatchCreateProcessRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewBatchCreateProcessResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateInspectHostRequest() (request *CreateInspectHostRequest) {
	request = &CreateInspectHostRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateInspectHost")
	return
}

func NewCreateInspectHostResponse() (response *CreateInspectHostResponse) {
	response = &CreateInspectHostResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInspectHost(request *CreateInspectHostRequest) (string) {
	return c.CreateInspectHostWithContext(context.Background(), request)
}

func (c *Client) CreateInspectHostSend(request *CreateInspectHostRequest) (*CreateInspectHostResponse, error) {
	statusCode, msg, err := c.CreateInspectHostWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateInspectHostResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInspectHostWithContext(ctx context.Context, request *CreateInspectHostRequest) (string) {
	if request == nil {
		request = NewCreateInspectHostRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInspectHostResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInspectHostWithContextV2(ctx context.Context, request *CreateInspectHostRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInspectHostRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInspectHostResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInspectHostResultsRequest() (request *DescribeInspectHostResultsRequest) {
	request = &DescribeInspectHostResultsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeInspectHostResults")
	return
}

func NewDescribeInspectHostResultsResponse() (response *DescribeInspectHostResultsResponse) {
	response = &DescribeInspectHostResultsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInspectHostResults(request *DescribeInspectHostResultsRequest) (string) {
	return c.DescribeInspectHostResultsWithContext(context.Background(), request)
}

func (c *Client) DescribeInspectHostResultsSend(request *DescribeInspectHostResultsRequest) (*DescribeInspectHostResultsResponse, error) {
	statusCode, msg, err := c.DescribeInspectHostResultsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInspectHostResultsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInspectHostResultsWithContext(ctx context.Context, request *DescribeInspectHostResultsRequest) (string) {
	if request == nil {
		request = NewDescribeInspectHostResultsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInspectHostResultsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInspectHostResultsWithContextV2(ctx context.Context, request *DescribeInspectHostResultsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInspectHostResultsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInspectHostResultsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeXidDetailsRequest() (request *DescribeXidDetailsRequest) {
	request = &DescribeXidDetailsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeXidDetails")
	return
}

func NewDescribeXidDetailsResponse() (response *DescribeXidDetailsResponse) {
	response = &DescribeXidDetailsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeXidDetails(request *DescribeXidDetailsRequest) (string) {
	return c.DescribeXidDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeXidDetailsSend(request *DescribeXidDetailsRequest) (*DescribeXidDetailsResponse, error) {
	statusCode, msg, err := c.DescribeXidDetailsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeXidDetailsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeXidDetailsWithContext(ctx context.Context, request *DescribeXidDetailsRequest) (string) {
	if request == nil {
		request = NewDescribeXidDetailsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeXidDetailsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeXidDetailsWithContextV2(ctx context.Context, request *DescribeXidDetailsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeXidDetailsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeXidDetailsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoImagesRequest() (request *DescribeSoImagesRequest) {
	request = &DescribeSoImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoImages")
	return
}

func NewDescribeSoImagesResponse() (response *DescribeSoImagesResponse) {
	response = &DescribeSoImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoImages(request *DescribeSoImagesRequest) (string) {
	return c.DescribeSoImagesWithContext(context.Background(), request)
}

func (c *Client) DescribeSoImagesSend(request *DescribeSoImagesRequest) (*DescribeSoImagesResponse, error) {
	statusCode, msg, err := c.DescribeSoImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoImagesWithContext(ctx context.Context, request *DescribeSoImagesRequest) (string) {
	if request == nil {
		request = NewDescribeSoImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoImagesWithContextV2(ctx context.Context, request *DescribeSoImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRebootSoInstanceRequest() (request *RebootSoInstanceRequest) {
	request = &RebootSoInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "RebootSoInstance")
	return
}

func NewRebootSoInstanceResponse() (response *RebootSoInstanceResponse) {
	response = &RebootSoInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RebootSoInstance(request *RebootSoInstanceRequest) (string) {
	return c.RebootSoInstanceWithContext(context.Background(), request)
}

func (c *Client) RebootSoInstanceSend(request *RebootSoInstanceRequest) (*RebootSoInstanceResponse, error) {
	statusCode, msg, err := c.RebootSoInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RebootSoInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RebootSoInstanceWithContext(ctx context.Context, request *RebootSoInstanceRequest) (string) {
	if request == nil {
		request = NewRebootSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootSoInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RebootSoInstanceWithContextV2(ctx context.Context, request *RebootSoInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRebootSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRebootSoInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSoImagesRequest() (request *DeleteSoImagesRequest) {
	request = &DeleteSoImagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteSoImages")
	return
}

func NewDeleteSoImagesResponse() (response *DeleteSoImagesResponse) {
	response = &DeleteSoImagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSoImages(request *DeleteSoImagesRequest) (string) {
	return c.DeleteSoImagesWithContext(context.Background(), request)
}

func (c *Client) DeleteSoImagesSend(request *DeleteSoImagesRequest) (*DeleteSoImagesResponse, error) {
	statusCode, msg, err := c.DeleteSoImagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSoImagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSoImagesWithContext(ctx context.Context, request *DeleteSoImagesRequest) (string) {
	if request == nil {
		request = NewDeleteSoImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoImagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSoImagesWithContextV2(ctx context.Context, request *DeleteSoImagesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSoImagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoImagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSoVpcRequest() (request *DeleteSoVpcRequest) {
	request = &DeleteSoVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteSoVpc")
	return
}

func NewDeleteSoVpcResponse() (response *DeleteSoVpcResponse) {
	response = &DeleteSoVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSoVpc(request *DeleteSoVpcRequest) (string) {
	return c.DeleteSoVpcWithContext(context.Background(), request)
}

func (c *Client) DeleteSoVpcSend(request *DeleteSoVpcRequest) (*DeleteSoVpcResponse, error) {
	statusCode, msg, err := c.DeleteSoVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSoVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSoVpcWithContext(ctx context.Context, request *DeleteSoVpcRequest) (string) {
	if request == nil {
		request = NewDeleteSoVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSoVpcWithContextV2(ctx context.Context, request *DeleteSoVpcRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSoVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoInstanceTypeRequest() (request *DescribeSoInstanceTypeRequest) {
	request = &DescribeSoInstanceTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoInstanceType")
	return
}

func NewDescribeSoInstanceTypeResponse() (response *DescribeSoInstanceTypeResponse) {
	response = &DescribeSoInstanceTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoInstanceType(request *DescribeSoInstanceTypeRequest) (string) {
	return c.DescribeSoInstanceTypeWithContext(context.Background(), request)
}

func (c *Client) DescribeSoInstanceTypeSend(request *DescribeSoInstanceTypeRequest) (*DescribeSoInstanceTypeResponse, error) {
	statusCode, msg, err := c.DescribeSoInstanceTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoInstanceTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoInstanceTypeWithContext(ctx context.Context, request *DescribeSoInstanceTypeRequest) (string) {
	if request == nil {
		request = NewDescribeSoInstanceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoInstanceTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoInstanceTypeWithContextV2(ctx context.Context, request *DescribeSoInstanceTypeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoInstanceTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoInstanceTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoAvailableResourceRequest() (request *DescribeSoAvailableResourceRequest) {
	request = &DescribeSoAvailableResourceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoAvailableResource")
	return
}

func NewDescribeSoAvailableResourceResponse() (response *DescribeSoAvailableResourceResponse) {
	response = &DescribeSoAvailableResourceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoAvailableResource(request *DescribeSoAvailableResourceRequest) (string) {
	return c.DescribeSoAvailableResourceWithContext(context.Background(), request)
}

func (c *Client) DescribeSoAvailableResourceSend(request *DescribeSoAvailableResourceRequest) (*DescribeSoAvailableResourceResponse, error) {
	statusCode, msg, err := c.DescribeSoAvailableResourceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoAvailableResourceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoAvailableResourceWithContext(ctx context.Context, request *DescribeSoAvailableResourceRequest) (string) {
	if request == nil {
		request = NewDescribeSoAvailableResourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoAvailableResourceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoAvailableResourceWithContextV2(ctx context.Context, request *DescribeSoAvailableResourceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoAvailableResourceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoAvailableResourceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoInstancesRequest() (request *DescribeSoInstancesRequest) {
	request = &DescribeSoInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoInstances")
	return
}

func NewDescribeSoInstancesResponse() (response *DescribeSoInstancesResponse) {
	response = &DescribeSoInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoInstances(request *DescribeSoInstancesRequest) (string) {
	return c.DescribeSoInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeSoInstancesSend(request *DescribeSoInstancesRequest) (*DescribeSoInstancesResponse, error) {
	statusCode, msg, err := c.DescribeSoInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoInstancesWithContext(ctx context.Context, request *DescribeSoInstancesRequest) (string) {
	if request == nil {
		request = NewDescribeSoInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoInstancesWithContextV2(ctx context.Context, request *DescribeSoInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSoInstanceRequest() (request *DeleteSoInstanceRequest) {
	request = &DeleteSoInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteSoInstance")
	return
}

func NewDeleteSoInstanceResponse() (response *DeleteSoInstanceResponse) {
	response = &DeleteSoInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSoInstance(request *DeleteSoInstanceRequest) (string) {
	return c.DeleteSoInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteSoInstanceSend(request *DeleteSoInstanceRequest) (*DeleteSoInstanceResponse, error) {
	statusCode, msg, err := c.DeleteSoInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSoInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSoInstanceWithContext(ctx context.Context, request *DeleteSoInstanceRequest) (string) {
	if request == nil {
		request = NewDeleteSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSoInstanceWithContextV2(ctx context.Context, request *DeleteSoInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoSecurityGroupsRequest() (request *DescribeSoSecurityGroupsRequest) {
	request = &DescribeSoSecurityGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoSecurityGroups")
	return
}

func NewDescribeSoSecurityGroupsResponse() (response *DescribeSoSecurityGroupsResponse) {
	response = &DescribeSoSecurityGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoSecurityGroups(request *DescribeSoSecurityGroupsRequest) (string) {
	return c.DescribeSoSecurityGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeSoSecurityGroupsSend(request *DescribeSoSecurityGroupsRequest) (*DescribeSoSecurityGroupsResponse, error) {
	statusCode, msg, err := c.DescribeSoSecurityGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoSecurityGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoSecurityGroupsWithContext(ctx context.Context, request *DescribeSoSecurityGroupsRequest) (string) {
	if request == nil {
		request = NewDescribeSoSecurityGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoSecurityGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoSecurityGroupsWithContextV2(ctx context.Context, request *DescribeSoSecurityGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoSecurityGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoSecurityGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSoVpcRequest() (request *CreateSoVpcRequest) {
	request = &CreateSoVpcRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateSoVpc")
	return
}

func NewCreateSoVpcResponse() (response *CreateSoVpcResponse) {
	response = &CreateSoVpcResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSoVpc(request *CreateSoVpcRequest) (string) {
	return c.CreateSoVpcWithContext(context.Background(), request)
}

func (c *Client) CreateSoVpcSend(request *CreateSoVpcRequest) (*CreateSoVpcResponse, error) {
	statusCode, msg, err := c.CreateSoVpcWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSoVpcResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSoVpcWithContext(ctx context.Context, request *CreateSoVpcRequest) (string) {
	if request == nil {
		request = NewCreateSoVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoVpcResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSoVpcWithContextV2(ctx context.Context, request *CreateSoVpcRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSoVpcRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoVpcResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSoSubnetRequest() (request *DeleteSoSubnetRequest) {
	request = &DeleteSoSubnetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteSoSubnet")
	return
}

func NewDeleteSoSubnetResponse() (response *DeleteSoSubnetResponse) {
	response = &DeleteSoSubnetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSoSubnet(request *DeleteSoSubnetRequest) (string) {
	return c.DeleteSoSubnetWithContext(context.Background(), request)
}

func (c *Client) DeleteSoSubnetSend(request *DeleteSoSubnetRequest) (*DeleteSoSubnetResponse, error) {
	statusCode, msg, err := c.DeleteSoSubnetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSoSubnetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSoSubnetWithContext(ctx context.Context, request *DeleteSoSubnetRequest) (string) {
	if request == nil {
		request = NewDeleteSoSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoSubnetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSoSubnetWithContextV2(ctx context.Context, request *DeleteSoSubnetRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSoSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoSubnetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoKeyPairsRequest() (request *DescribeSoKeyPairsRequest) {
	request = &DescribeSoKeyPairsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoKeyPairs")
	return
}

func NewDescribeSoKeyPairsResponse() (response *DescribeSoKeyPairsResponse) {
	response = &DescribeSoKeyPairsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoKeyPairs(request *DescribeSoKeyPairsRequest) (string) {
	return c.DescribeSoKeyPairsWithContext(context.Background(), request)
}

func (c *Client) DescribeSoKeyPairsSend(request *DescribeSoKeyPairsRequest) (*DescribeSoKeyPairsResponse, error) {
	statusCode, msg, err := c.DescribeSoKeyPairsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoKeyPairsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoKeyPairsWithContext(ctx context.Context, request *DescribeSoKeyPairsRequest) (string) {
	if request == nil {
		request = NewDescribeSoKeyPairsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoKeyPairsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoKeyPairsWithContextV2(ctx context.Context, request *DescribeSoKeyPairsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoKeyPairsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoKeyPairsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartSoInstanceRequest() (request *StartSoInstanceRequest) {
	request = &StartSoInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "StartSoInstance")
	return
}

func NewStartSoInstanceResponse() (response *StartSoInstanceResponse) {
	response = &StartSoInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartSoInstance(request *StartSoInstanceRequest) (string) {
	return c.StartSoInstanceWithContext(context.Background(), request)
}

func (c *Client) StartSoInstanceSend(request *StartSoInstanceRequest) (*StartSoInstanceResponse, error) {
	statusCode, msg, err := c.StartSoInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartSoInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartSoInstanceWithContext(ctx context.Context, request *StartSoInstanceRequest) (string) {
	if request == nil {
		request = NewStartSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartSoInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartSoInstanceWithContextV2(ctx context.Context, request *StartSoInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewStartSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartSoInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoInstanceTypesRequest() (request *DescribeSoInstanceTypesRequest) {
	request = &DescribeSoInstanceTypesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoInstanceTypes")
	return
}

func NewDescribeSoInstanceTypesResponse() (response *DescribeSoInstanceTypesResponse) {
	response = &DescribeSoInstanceTypesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoInstanceTypes(request *DescribeSoInstanceTypesRequest) (string) {
	return c.DescribeSoInstanceTypesWithContext(context.Background(), request)
}

func (c *Client) DescribeSoInstanceTypesSend(request *DescribeSoInstanceTypesRequest) (*DescribeSoInstanceTypesResponse, error) {
	statusCode, msg, err := c.DescribeSoInstanceTypesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoInstanceTypesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoInstanceTypesWithContext(ctx context.Context, request *DescribeSoInstanceTypesRequest) (string) {
	if request == nil {
		request = NewDescribeSoInstanceTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoInstanceTypesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoInstanceTypesWithContextV2(ctx context.Context, request *DescribeSoInstanceTypesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoInstanceTypesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoInstanceTypesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySoSubnetAttributesRequest() (request *ModifySoSubnetAttributesRequest) {
	request = &ModifySoSubnetAttributesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifySoSubnetAttributes")
	return
}

func NewModifySoSubnetAttributesResponse() (response *ModifySoSubnetAttributesResponse) {
	response = &ModifySoSubnetAttributesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySoSubnetAttributes(request *ModifySoSubnetAttributesRequest) (string) {
	return c.ModifySoSubnetAttributesWithContext(context.Background(), request)
}

func (c *Client) ModifySoSubnetAttributesSend(request *ModifySoSubnetAttributesRequest) (*ModifySoSubnetAttributesResponse, error) {
	statusCode, msg, err := c.ModifySoSubnetAttributesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySoSubnetAttributesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySoSubnetAttributesWithContext(ctx context.Context, request *ModifySoSubnetAttributesRequest) (string) {
	if request == nil {
		request = NewModifySoSubnetAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoSubnetAttributesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySoSubnetAttributesWithContextV2(ctx context.Context, request *ModifySoSubnetAttributesRequest) (int, string, error) {
	if request == nil {
		request = NewModifySoSubnetAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoSubnetAttributesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoSubnetRequest() (request *DescribeSoSubnetRequest) {
	request = &DescribeSoSubnetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoSubnet")
	return
}

func NewDescribeSoSubnetResponse() (response *DescribeSoSubnetResponse) {
	response = &DescribeSoSubnetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoSubnet(request *DescribeSoSubnetRequest) (string) {
	return c.DescribeSoSubnetWithContext(context.Background(), request)
}

func (c *Client) DescribeSoSubnetSend(request *DescribeSoSubnetRequest) (*DescribeSoSubnetResponse, error) {
	statusCode, msg, err := c.DescribeSoSubnetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoSubnetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoSubnetWithContext(ctx context.Context, request *DescribeSoSubnetRequest) (string) {
	if request == nil {
		request = NewDescribeSoSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoSubnetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoSubnetWithContextV2(ctx context.Context, request *DescribeSoSubnetRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoSubnetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySoKeyPairAttributeRequest() (request *ModifySoKeyPairAttributeRequest) {
	request = &ModifySoKeyPairAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifySoKeyPairAttribute")
	return
}

func NewModifySoKeyPairAttributeResponse() (response *ModifySoKeyPairAttributeResponse) {
	response = &ModifySoKeyPairAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySoKeyPairAttribute(request *ModifySoKeyPairAttributeRequest) (string) {
	return c.ModifySoKeyPairAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifySoKeyPairAttributeSend(request *ModifySoKeyPairAttributeRequest) (*ModifySoKeyPairAttributeResponse, error) {
	statusCode, msg, err := c.ModifySoKeyPairAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySoKeyPairAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySoKeyPairAttributeWithContext(ctx context.Context, request *ModifySoKeyPairAttributeRequest) (string) {
	if request == nil {
		request = NewModifySoKeyPairAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoKeyPairAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySoKeyPairAttributeWithContextV2(ctx context.Context, request *ModifySoKeyPairAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifySoKeyPairAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoKeyPairAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySoImageAttributeRequest() (request *ModifySoImageAttributeRequest) {
	request = &ModifySoImageAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifySoImageAttribute")
	return
}

func NewModifySoImageAttributeResponse() (response *ModifySoImageAttributeResponse) {
	response = &ModifySoImageAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySoImageAttribute(request *ModifySoImageAttributeRequest) (string) {
	return c.ModifySoImageAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifySoImageAttributeSend(request *ModifySoImageAttributeRequest) (*ModifySoImageAttributeResponse, error) {
	statusCode, msg, err := c.ModifySoImageAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySoImageAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySoImageAttributeWithContext(ctx context.Context, request *ModifySoImageAttributeRequest) (string) {
	if request == nil {
		request = NewModifySoImageAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoImageAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySoImageAttributeWithContextV2(ctx context.Context, request *ModifySoImageAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifySoImageAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoImageAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySoVpcAttributesRequest() (request *ModifySoVpcAttributesRequest) {
	request = &ModifySoVpcAttributesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifySoVpcAttributes")
	return
}

func NewModifySoVpcAttributesResponse() (response *ModifySoVpcAttributesResponse) {
	response = &ModifySoVpcAttributesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySoVpcAttributes(request *ModifySoVpcAttributesRequest) (string) {
	return c.ModifySoVpcAttributesWithContext(context.Background(), request)
}

func (c *Client) ModifySoVpcAttributesSend(request *ModifySoVpcAttributesRequest) (*ModifySoVpcAttributesResponse, error) {
	statusCode, msg, err := c.ModifySoVpcAttributesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySoVpcAttributesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySoVpcAttributesWithContext(ctx context.Context, request *ModifySoVpcAttributesRequest) (string) {
	if request == nil {
		request = NewModifySoVpcAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoVpcAttributesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySoVpcAttributesWithContextV2(ctx context.Context, request *ModifySoVpcAttributesRequest) (int, string, error) {
	if request == nil {
		request = NewModifySoVpcAttributesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoVpcAttributesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReplaceSoSystemVolumeRequest() (request *ReplaceSoSystemVolumeRequest) {
	request = &ReplaceSoSystemVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ReplaceSoSystemVolume")
	return
}

func NewReplaceSoSystemVolumeResponse() (response *ReplaceSoSystemVolumeResponse) {
	response = &ReplaceSoSystemVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReplaceSoSystemVolume(request *ReplaceSoSystemVolumeRequest) (string) {
	return c.ReplaceSoSystemVolumeWithContext(context.Background(), request)
}

func (c *Client) ReplaceSoSystemVolumeSend(request *ReplaceSoSystemVolumeRequest) (*ReplaceSoSystemVolumeResponse, error) {
	statusCode, msg, err := c.ReplaceSoSystemVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReplaceSoSystemVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReplaceSoSystemVolumeWithContext(ctx context.Context, request *ReplaceSoSystemVolumeRequest) (string) {
	if request == nil {
		request = NewReplaceSoSystemVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReplaceSoSystemVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReplaceSoSystemVolumeWithContextV2(ctx context.Context, request *ReplaceSoSystemVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewReplaceSoSystemVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReplaceSoSystemVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSoSubnetRequest() (request *CreateSoSubnetRequest) {
	request = &CreateSoSubnetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateSoSubnet")
	return
}

func NewCreateSoSubnetResponse() (response *CreateSoSubnetResponse) {
	response = &CreateSoSubnetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSoSubnet(request *CreateSoSubnetRequest) (string) {
	return c.CreateSoSubnetWithContext(context.Background(), request)
}

func (c *Client) CreateSoSubnetSend(request *CreateSoSubnetRequest) (*CreateSoSubnetResponse, error) {
	statusCode, msg, err := c.CreateSoSubnetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSoSubnetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSoSubnetWithContext(ctx context.Context, request *CreateSoSubnetRequest) (string) {
	if request == nil {
		request = NewCreateSoSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoSubnetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSoSubnetWithContextV2(ctx context.Context, request *CreateSoSubnetRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSoSubnetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoSubnetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSoVpcsRequest() (request *DescribeSoVpcsRequest) {
	request = &DescribeSoVpcsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DescribeSoVpcs")
	return
}

func NewDescribeSoVpcsResponse() (response *DescribeSoVpcsResponse) {
	response = &DescribeSoVpcsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSoVpcs(request *DescribeSoVpcsRequest) string {
	return c.DescribeSoVpcsWithContext(context.Background(), request)
}

func (c *Client) DescribeSoVpcsSend(request *DescribeSoVpcsRequest) (*DescribeSoVpcsResponse, error) {
	statusCode, msg, err := c.DescribeSoVpcsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSoVpcsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSoVpcsWithContext(ctx context.Context, request *DescribeSoVpcsRequest) string {
	if request == nil {
		request = NewDescribeSoVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoVpcsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSoVpcsWithContextV2(ctx context.Context, request *DescribeSoVpcsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSoVpcsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSoVpcsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStopSoInstanceRequest() (request *StopSoInstanceRequest) {
	request = &StopSoInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "StopSoInstance")
	return
}

func NewStopSoInstanceResponse() (response *StopSoInstanceResponse) {
	response = &StopSoInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StopSoInstance(request *StopSoInstanceRequest) string {
	return c.StopSoInstanceWithContext(context.Background(), request)
}

func (c *Client) StopSoInstanceSend(request *StopSoInstanceRequest) (*StopSoInstanceResponse, error) {
	statusCode, msg, err := c.StopSoInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StopSoInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StopSoInstanceWithContext(ctx context.Context, request *StopSoInstanceRequest) string {
	if request == nil {
		request = NewStopSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopSoInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StopSoInstanceWithContextV2(ctx context.Context, request *StopSoInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewStopSoInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStopSoInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSoKeyPairsRequest() (request *DeleteSoKeyPairsRequest) {
	request = &DeleteSoKeyPairsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "DeleteSoKeyPairs")
	return
}

func NewDeleteSoKeyPairsResponse() (response *DeleteSoKeyPairsResponse) {
	response = &DeleteSoKeyPairsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSoKeyPairs(request *DeleteSoKeyPairsRequest) string {
	return c.DeleteSoKeyPairsWithContext(context.Background(), request)
}

func (c *Client) DeleteSoKeyPairsSend(request *DeleteSoKeyPairsRequest) (*DeleteSoKeyPairsResponse, error) {
	statusCode, msg, err := c.DeleteSoKeyPairsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSoKeyPairsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSoKeyPairsWithContext(ctx context.Context, request *DeleteSoKeyPairsRequest) string {
	if request == nil {
		request = NewDeleteSoKeyPairsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoKeyPairsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSoKeyPairsWithContextV2(ctx context.Context, request *DeleteSoKeyPairsRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSoKeyPairsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSoKeyPairsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSoImageRequest() (request *CreateSoImageRequest) {
	request = &CreateSoImageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateSoImage")
	return
}

func NewCreateSoImageResponse() (response *CreateSoImageResponse) {
	response = &CreateSoImageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSoImage(request *CreateSoImageRequest) string {
	return c.CreateSoImageWithContext(context.Background(), request)
}

func (c *Client) CreateSoImageSend(request *CreateSoImageRequest) (*CreateSoImageResponse, error) {
	statusCode, msg, err := c.CreateSoImageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSoImageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSoImageWithContext(ctx context.Context, request *CreateSoImageRequest) string {
	if request == nil {
		request = NewCreateSoImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoImageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSoImageWithContextV2(ctx context.Context, request *CreateSoImageRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSoImageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoImageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySoInstanceAttributeRequest() (request *ModifySoInstanceAttributeRequest) {
	request = &ModifySoInstanceAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "ModifySoInstanceAttribute")
	return
}

func NewModifySoInstanceAttributeResponse() (response *ModifySoInstanceAttributeResponse) {
	response = &ModifySoInstanceAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySoInstanceAttribute(request *ModifySoInstanceAttributeRequest) string {
	return c.ModifySoInstanceAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifySoInstanceAttributeSend(request *ModifySoInstanceAttributeRequest) (*ModifySoInstanceAttributeResponse, error) {
	statusCode, msg, err := c.ModifySoInstanceAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySoInstanceAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySoInstanceAttributeWithContext(ctx context.Context, request *ModifySoInstanceAttributeRequest) string {
	if request == nil {
		request = NewModifySoInstanceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoInstanceAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySoInstanceAttributeWithContextV2(ctx context.Context, request *ModifySoInstanceAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifySoInstanceAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySoInstanceAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSoKeyPairRequest() (request *CreateSoKeyPairRequest) {
	request = &CreateSoKeyPairRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("epc", APIVersion, "CreateSoKeyPair")
	return
}

func NewCreateSoKeyPairResponse() (response *CreateSoKeyPairResponse) {
	response = &CreateSoKeyPairResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSoKeyPair(request *CreateSoKeyPairRequest) string {
	return c.CreateSoKeyPairWithContext(context.Background(), request)
}

func (c *Client) CreateSoKeyPairSend(request *CreateSoKeyPairRequest) (*CreateSoKeyPairResponse, error) {
	statusCode, msg, err := c.CreateSoKeyPairWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSoKeyPairResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSoKeyPairWithContext(ctx context.Context, request *CreateSoKeyPairRequest) string {
	if request == nil {
		request = NewCreateSoKeyPairRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoKeyPairResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSoKeyPairWithContextV2(ctx context.Context, request *CreateSoKeyPairRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSoKeyPairRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSoKeyPairResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


