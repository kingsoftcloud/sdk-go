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

func NewCreateKeyRequest() (request *CreateKeyRequest) {
	request = &CreateKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sks", APIVersion, "CreateKey")
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

	if msg == "" {
		return nil, nil
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "CreateKey")
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
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "CreateKey")
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
func NewImportKeyRequest() (request *ImportKeyRequest) {
	request = &ImportKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sks", APIVersion, "ImportKey")
	return
}

func NewImportKeyResponse() (response *ImportKeyResponse) {
	response = &ImportKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportKey(request *ImportKeyRequest) string {
	return c.ImportKeyWithContext(context.Background(), request)
}

func (c *Client) ImportKeySend(request *ImportKeyRequest) (*ImportKeyResponse, error) {
	statusCode, msg, err := c.ImportKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ImportKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ImportKeyWithContext(ctx context.Context, request *ImportKeyRequest) string {
	if request == nil {
		request = NewImportKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "ImportKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewImportKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ImportKeyWithContextV2(ctx context.Context, request *ImportKeyRequest) (int, string, error) {
	if request == nil {
		request = NewImportKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "ImportKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewImportKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteKeyRequest() (request *DeleteKeyRequest) {
	request = &DeleteKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sks", APIVersion, "DeleteKey")
	return
}

func NewDeleteKeyResponse() (response *DeleteKeyResponse) {
	response = &DeleteKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteKey(request *DeleteKeyRequest) string {
	return c.DeleteKeyWithContext(context.Background(), request)
}

func (c *Client) DeleteKeySend(request *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	statusCode, msg, err := c.DeleteKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteKeyWithContext(ctx context.Context, request *DeleteKeyRequest) string {
	if request == nil {
		request = NewDeleteKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "DeleteKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteKeyWithContextV2(ctx context.Context, request *DeleteKeyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "DeleteKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyKeyRequest() (request *ModifyKeyRequest) {
	request = &ModifyKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sks", APIVersion, "ModifyKey")
	return
}

func NewModifyKeyResponse() (response *ModifyKeyResponse) {
	response = &ModifyKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyKey(request *ModifyKeyRequest) string {
	return c.ModifyKeyWithContext(context.Background(), request)
}

func (c *Client) ModifyKeySend(request *ModifyKeyRequest) (*ModifyKeyResponse, error) {
	statusCode, msg, err := c.ModifyKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyKeyWithContext(ctx context.Context, request *ModifyKeyRequest) string {
	if request == nil {
		request = NewModifyKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "ModifyKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyKeyWithContextV2(ctx context.Context, request *ModifyKeyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "ModifyKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeKeysRequest() (request *DescribeKeysRequest) {
	request = &DescribeKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sks", APIVersion, "DescribeKeys")
	return
}

func NewDescribeKeysResponse() (response *DescribeKeysResponse) {
	response = &DescribeKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeKeys(request *DescribeKeysRequest) string {
	return c.DescribeKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeKeysSend(request *DescribeKeysRequest) (*DescribeKeysResponse, error) {
	statusCode, msg, err := c.DescribeKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeKeysWithContext(ctx context.Context, request *DescribeKeysRequest) string {
	if request == nil {
		request = NewDescribeKeysRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "DescribeKeys")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeKeysWithContextV2(ctx context.Context, request *DescribeKeysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeKeysRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("sks", APIVersion, "DescribeKeys")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
