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

func NewCreateKeyRequest() (request *CreateKeyRequest) {
	request = &CreateKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "CreateKey")
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
		request.Init().WithApiInfo("kkms", APIVersion, "CreateKey")
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
		request.Init().WithApiInfo("kkms", APIVersion, "CreateKey")
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
func NewModifyKeyRequest() (request *ModifyKeyRequest) {
	request = &ModifyKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "ModifyKey")
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
		request.Init().WithApiInfo("kkms", APIVersion, "ModifyKey")
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
		request.Init().WithApiInfo("kkms", APIVersion, "ModifyKey")
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
func NewModifyKeyStateRequest() (request *ModifyKeyStateRequest) {
	request = &ModifyKeyStateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "ModifyKeyState")
	return
}

func NewModifyKeyStateResponse() (response *ModifyKeyStateResponse) {
	response = &ModifyKeyStateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyKeyState(request *ModifyKeyStateRequest) string {
	return c.ModifyKeyStateWithContext(context.Background(), request)
}

func (c *Client) ModifyKeyStateSend(request *ModifyKeyStateRequest) (*ModifyKeyStateResponse, error) {
	statusCode, msg, err := c.ModifyKeyStateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyKeyStateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyKeyStateWithContext(ctx context.Context, request *ModifyKeyStateRequest) string {
	if request == nil {
		request = NewModifyKeyStateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "ModifyKeyState")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKeyStateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyKeyStateWithContextV2(ctx context.Context, request *ModifyKeyStateRequest) (int, string, error) {
	if request == nil {
		request = NewModifyKeyStateRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "ModifyKeyState")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyKeyStateResponse()
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
	request.Init().WithApiInfo("kkms", APIVersion, "DeleteKey")
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
		request.Init().WithApiInfo("kkms", APIVersion, "DeleteKey")
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
		request.Init().WithApiInfo("kkms", APIVersion, "DeleteKey")
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
func NewDescribeKeysRequest() (request *DescribeKeysRequest) {
	request = &DescribeKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "DescribeKeys")
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
		request.Init().WithApiInfo("kkms", APIVersion, "DescribeKeys")
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
		request.Init().WithApiInfo("kkms", APIVersion, "DescribeKeys")
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
func NewEncryptRequest() (request *EncryptRequest) {
	request = &EncryptRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "Encrypt")
	return
}

func NewEncryptResponse() (response *EncryptResponse) {
	response = &EncryptResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Encrypt(request *EncryptRequest) string {
	return c.EncryptWithContext(context.Background(), request)
}

func (c *Client) EncryptSend(request *EncryptRequest) (*EncryptResponse, error) {
	statusCode, msg, err := c.EncryptWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct EncryptResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) EncryptWithContext(ctx context.Context, request *EncryptRequest) string {
	if request == nil {
		request = NewEncryptRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "Encrypt")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEncryptResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) EncryptWithContextV2(ctx context.Context, request *EncryptRequest) (int, string, error) {
	if request == nil {
		request = NewEncryptRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "Encrypt")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewEncryptResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDecryptRequest() (request *DecryptRequest) {
	request = &DecryptRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "Decrypt")
	return
}

func NewDecryptResponse() (response *DecryptResponse) {
	response = &DecryptResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Decrypt(request *DecryptRequest) string {
	return c.DecryptWithContext(context.Background(), request)
}

func (c *Client) DecryptSend(request *DecryptRequest) (*DecryptResponse, error) {
	statusCode, msg, err := c.DecryptWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DecryptResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DecryptWithContext(ctx context.Context, request *DecryptRequest) string {
	if request == nil {
		request = NewDecryptRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "Decrypt")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDecryptResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DecryptWithContextV2(ctx context.Context, request *DecryptRequest) (int, string, error) {
	if request == nil {
		request = NewDecryptRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "Decrypt")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDecryptResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGenerateDataKeyRequest() (request *GenerateDataKeyRequest) {
	request = &GenerateDataKeyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kkms", APIVersion, "GenerateDataKey")
	return
}

func NewGenerateDataKeyResponse() (response *GenerateDataKeyResponse) {
	response = &GenerateDataKeyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GenerateDataKey(request *GenerateDataKeyRequest) string {
	return c.GenerateDataKeyWithContext(context.Background(), request)
}

func (c *Client) GenerateDataKeySend(request *GenerateDataKeyRequest) (*GenerateDataKeyResponse, error) {
	statusCode, msg, err := c.GenerateDataKeyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct GenerateDataKeyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GenerateDataKeyWithContext(ctx context.Context, request *GenerateDataKeyRequest) string {
	if request == nil {
		request = NewGenerateDataKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "GenerateDataKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGenerateDataKeyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GenerateDataKeyWithContextV2(ctx context.Context, request *GenerateDataKeyRequest) (int, string, error) {
	if request == nil {
		request = NewGenerateDataKeyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kkms", APIVersion, "GenerateDataKey")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGenerateDataKeyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
