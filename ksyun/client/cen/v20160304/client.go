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

func NewCreateCenRequest() (request *CreateCenRequest) {
	request = &CreateCenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCen")
	return
}

func NewCreateCenResponse() (response *CreateCenResponse) {
	response = &CreateCenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCen(request *CreateCenRequest) string {
	return c.CreateCenWithContext(context.Background(), request)
}

func (c *Client) CreateCenSend(request *CreateCenRequest) (*CreateCenResponse, error) {
	statusCode, msg, err := c.CreateCenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCenWithContext(ctx context.Context, request *CreateCenRequest) string {
	if request == nil {
		request = NewCreateCenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCen")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCenWithContextV2(ctx context.Context, request *CreateCenRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCen")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCenRequest() (request *ModifyCenRequest) {
	request = &ModifyCenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "ModifyCen")
	return
}

func NewModifyCenResponse() (response *ModifyCenResponse) {
	response = &ModifyCenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCen(request *ModifyCenRequest) string {
	return c.ModifyCenWithContext(context.Background(), request)
}

func (c *Client) ModifyCenSend(request *ModifyCenRequest) (*ModifyCenResponse, error) {
	statusCode, msg, err := c.ModifyCenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCenWithContext(ctx context.Context, request *ModifyCenRequest) string {
	if request == nil {
		request = NewModifyCenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "ModifyCen")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCenWithContextV2(ctx context.Context, request *ModifyCenRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "ModifyCen")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCenRequest() (request *DeleteCenRequest) {
	request = &DeleteCenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCen")
	return
}

func NewDeleteCenResponse() (response *DeleteCenResponse) {
	response = &DeleteCenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCen(request *DeleteCenRequest) string {
	return c.DeleteCenWithContext(context.Background(), request)
}

func (c *Client) DeleteCenSend(request *DeleteCenRequest) (*DeleteCenResponse, error) {
	statusCode, msg, err := c.DeleteCenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCenWithContext(ctx context.Context, request *DeleteCenRequest) string {
	if request == nil {
		request = NewDeleteCenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCen")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCenWithContextV2(ctx context.Context, request *DeleteCenRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCenRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCen")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCensRequest() (request *DescribeCensRequest) {
	request = &DescribeCensRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCens")
	return
}

func NewDescribeCensResponse() (response *DescribeCensResponse) {
	response = &DescribeCensResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCens(request *DescribeCensRequest) string {
	return c.DescribeCensWithContext(context.Background(), request)
}

func (c *Client) DescribeCensSend(request *DescribeCensRequest) (*DescribeCensResponse, error) {
	statusCode, msg, err := c.DescribeCensWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCensResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCensWithContext(ctx context.Context, request *DescribeCensRequest) string {
	if request == nil {
		request = NewDescribeCensRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCens")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCensResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCensWithContextV2(ctx context.Context, request *DescribeCensRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCensRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCens")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCensResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCenGrantRequest() (request *DeleteCenGrantRequest) {
	request = &DeleteCenGrantRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCenGrant")
	return
}

func NewDeleteCenGrantResponse() (response *DeleteCenGrantResponse) {
	response = &DeleteCenGrantResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCenGrant(request *DeleteCenGrantRequest) string {
	return c.DeleteCenGrantWithContext(context.Background(), request)
}

func (c *Client) DeleteCenGrantSend(request *DeleteCenGrantRequest) (*DeleteCenGrantResponse, error) {
	statusCode, msg, err := c.DeleteCenGrantWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCenGrantResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCenGrantWithContext(ctx context.Context, request *DeleteCenGrantRequest) string {
	if request == nil {
		request = NewDeleteCenGrantRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCenGrant")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenGrantResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCenGrantWithContextV2(ctx context.Context, request *DeleteCenGrantRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCenGrantRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCenGrant")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenGrantResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCenGrantsRequest() (request *DescribeCenGrantsRequest) {
	request = &DescribeCenGrantsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenGrants")
	return
}

func NewDescribeCenGrantsResponse() (response *DescribeCenGrantsResponse) {
	response = &DescribeCenGrantsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenGrants(request *DescribeCenGrantsRequest) string {
	return c.DescribeCenGrantsWithContext(context.Background(), request)
}

func (c *Client) DescribeCenGrantsSend(request *DescribeCenGrantsRequest) (*DescribeCenGrantsResponse, error) {
	statusCode, msg, err := c.DescribeCenGrantsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCenGrantsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCenGrantsWithContext(ctx context.Context, request *DescribeCenGrantsRequest) string {
	if request == nil {
		request = NewDescribeCenGrantsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenGrants")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenGrantsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCenGrantsWithContextV2(ctx context.Context, request *DescribeCenGrantsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCenGrantsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenGrants")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenGrantsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCenBandWidthPackageRequest() (request *CreateCenBandWidthPackageRequest) {
	request = &CreateCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCenBandWidthPackage")
	return
}

func NewCreateCenBandWidthPackageResponse() (response *CreateCenBandWidthPackageResponse) {
	response = &CreateCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCenBandWidthPackage(request *CreateCenBandWidthPackageRequest) string {
	return c.CreateCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) CreateCenBandWidthPackageSend(request *CreateCenBandWidthPackageRequest) (*CreateCenBandWidthPackageResponse, error) {
	statusCode, msg, err := c.CreateCenBandWidthPackageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCenBandWidthPackageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCenBandWidthPackageWithContext(ctx context.Context, request *CreateCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewCreateCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCenBandWidthPackageWithContextV2(ctx context.Context, request *CreateCenBandWidthPackageRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenBandWidthPackageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCenBandWidthPackageRequest() (request *ModifyCenBandWidthPackageRequest) {
	request = &ModifyCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "ModifyCenBandWidthPackage")
	return
}

func NewModifyCenBandWidthPackageResponse() (response *ModifyCenBandWidthPackageResponse) {
	response = &ModifyCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCenBandWidthPackage(request *ModifyCenBandWidthPackageRequest) string {
	return c.ModifyCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) ModifyCenBandWidthPackageSend(request *ModifyCenBandWidthPackageRequest) (*ModifyCenBandWidthPackageResponse, error) {
	statusCode, msg, err := c.ModifyCenBandWidthPackageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCenBandWidthPackageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCenBandWidthPackageWithContext(ctx context.Context, request *ModifyCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewModifyCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "ModifyCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCenBandWidthPackageWithContextV2(ctx context.Context, request *ModifyCenBandWidthPackageRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "ModifyCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenBandWidthPackageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCenBandWidthPackageRequest() (request *DeleteCenBandWidthPackageRequest) {
	request = &DeleteCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCenBandWidthPackage")
	return
}

func NewDeleteCenBandWidthPackageResponse() (response *DeleteCenBandWidthPackageResponse) {
	response = &DeleteCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCenBandWidthPackage(request *DeleteCenBandWidthPackageRequest) string {
	return c.DeleteCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) DeleteCenBandWidthPackageSend(request *DeleteCenBandWidthPackageRequest) (*DeleteCenBandWidthPackageResponse, error) {
	statusCode, msg, err := c.DeleteCenBandWidthPackageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCenBandWidthPackageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCenBandWidthPackageWithContext(ctx context.Context, request *DeleteCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewDeleteCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCenBandWidthPackageWithContextV2(ctx context.Context, request *DeleteCenBandWidthPackageRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenBandWidthPackageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachCenBandWidthPackageRequest() (request *AttachCenBandWidthPackageRequest) {
	request = &AttachCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "AttachCenBandWidthPackage")
	return
}

func NewAttachCenBandWidthPackageResponse() (response *AttachCenBandWidthPackageResponse) {
	response = &AttachCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachCenBandWidthPackage(request *AttachCenBandWidthPackageRequest) string {
	return c.AttachCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) AttachCenBandWidthPackageSend(request *AttachCenBandWidthPackageRequest) (*AttachCenBandWidthPackageResponse, error) {
	statusCode, msg, err := c.AttachCenBandWidthPackageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AttachCenBandWidthPackageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachCenBandWidthPackageWithContext(ctx context.Context, request *AttachCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewAttachCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "AttachCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachCenBandWidthPackageWithContextV2(ctx context.Context, request *AttachCenBandWidthPackageRequest) (int, string, error) {
	if request == nil {
		request = NewAttachCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "AttachCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachCenBandWidthPackageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachCenBandWidthPackageRequest() (request *DetachCenBandWidthPackageRequest) {
	request = &DetachCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DetachCenBandWidthPackage")
	return
}

func NewDetachCenBandWidthPackageResponse() (response *DetachCenBandWidthPackageResponse) {
	response = &DetachCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachCenBandWidthPackage(request *DetachCenBandWidthPackageRequest) string {
	return c.DetachCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) DetachCenBandWidthPackageSend(request *DetachCenBandWidthPackageRequest) (*DetachCenBandWidthPackageResponse, error) {
	statusCode, msg, err := c.DetachCenBandWidthPackageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DetachCenBandWidthPackageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachCenBandWidthPackageWithContext(ctx context.Context, request *DetachCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewDetachCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DetachCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachCenBandWidthPackageWithContextV2(ctx context.Context, request *DetachCenBandWidthPackageRequest) (int, string, error) {
	if request == nil {
		request = NewDetachCenBandWidthPackageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DetachCenBandWidthPackage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachCenBandWidthPackageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCenBandWidthPackagesRequest() (request *DescribeCenBandWidthPackagesRequest) {
	request = &DescribeCenBandWidthPackagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackages")
	return
}

func NewDescribeCenBandWidthPackagesResponse() (response *DescribeCenBandWidthPackagesResponse) {
	response = &DescribeCenBandWidthPackagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenBandWidthPackages(request *DescribeCenBandWidthPackagesRequest) string {
	return c.DescribeCenBandWidthPackagesWithContext(context.Background(), request)
}

func (c *Client) DescribeCenBandWidthPackagesSend(request *DescribeCenBandWidthPackagesRequest) (*DescribeCenBandWidthPackagesResponse, error) {
	statusCode, msg, err := c.DescribeCenBandWidthPackagesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCenBandWidthPackagesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCenBandWidthPackagesWithContext(ctx context.Context, request *DescribeCenBandWidthPackagesRequest) string {
	if request == nil {
		request = NewDescribeCenBandWidthPackagesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackages")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCenBandWidthPackagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCenBandWidthPackagesWithContextV2(ctx context.Context, request *DescribeCenBandWidthPackagesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCenBandWidthPackagesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackages")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCenBandWidthPackagesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCenRegionBandwidthRequest() (request *CreateCenRegionBandwidthRequest) {
	request = &CreateCenRegionBandwidthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCenRegionBandwidth")
	return
}

func NewCreateCenRegionBandwidthResponse() (response *CreateCenRegionBandwidthResponse) {
	response = &CreateCenRegionBandwidthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCenRegionBandwidth(request *CreateCenRegionBandwidthRequest) string {
	return c.CreateCenRegionBandwidthWithContext(context.Background(), request)
}

func (c *Client) CreateCenRegionBandwidthSend(request *CreateCenRegionBandwidthRequest) (*CreateCenRegionBandwidthResponse, error) {
	statusCode, msg, err := c.CreateCenRegionBandwidthWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCenRegionBandwidthResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCenRegionBandwidthWithContext(ctx context.Context, request *CreateCenRegionBandwidthRequest) string {
	if request == nil {
		request = NewCreateCenRegionBandwidthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCenRegionBandwidth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenRegionBandwidthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCenRegionBandwidthWithContextV2(ctx context.Context, request *CreateCenRegionBandwidthRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCenRegionBandwidthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCenRegionBandwidth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenRegionBandwidthResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCenRegionBandwidthRequest() (request *DeleteCenRegionBandwidthRequest) {
	request = &DeleteCenRegionBandwidthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCenRegionBandwidth")
	return
}

func NewDeleteCenRegionBandwidthResponse() (response *DeleteCenRegionBandwidthResponse) {
	response = &DeleteCenRegionBandwidthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCenRegionBandwidth(request *DeleteCenRegionBandwidthRequest) string {
	return c.DeleteCenRegionBandwidthWithContext(context.Background(), request)
}

func (c *Client) DeleteCenRegionBandwidthSend(request *DeleteCenRegionBandwidthRequest) (*DeleteCenRegionBandwidthResponse, error) {
	statusCode, msg, err := c.DeleteCenRegionBandwidthWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteCenRegionBandwidthResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCenRegionBandwidthWithContext(ctx context.Context, request *DeleteCenRegionBandwidthRequest) string {
	if request == nil {
		request = NewDeleteCenRegionBandwidthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCenRegionBandwidth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenRegionBandwidthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCenRegionBandwidthWithContextV2(ctx context.Context, request *DeleteCenRegionBandwidthRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCenRegionBandwidthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DeleteCenRegionBandwidth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenRegionBandwidthResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCenRegionBandwidthRequest() (request *ModifyCenRegionBandwidthRequest) {
	request = &ModifyCenRegionBandwidthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "ModifyCenRegionBandwidth")
	return
}

func NewModifyCenRegionBandwidthResponse() (response *ModifyCenRegionBandwidthResponse) {
	response = &ModifyCenRegionBandwidthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCenRegionBandwidth(request *ModifyCenRegionBandwidthRequest) string {
	return c.ModifyCenRegionBandwidthWithContext(context.Background(), request)
}

func (c *Client) ModifyCenRegionBandwidthSend(request *ModifyCenRegionBandwidthRequest) (*ModifyCenRegionBandwidthResponse, error) {
	statusCode, msg, err := c.ModifyCenRegionBandwidthWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ModifyCenRegionBandwidthResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCenRegionBandwidthWithContext(ctx context.Context, request *ModifyCenRegionBandwidthRequest) string {
	if request == nil {
		request = NewModifyCenRegionBandwidthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "ModifyCenRegionBandwidth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenRegionBandwidthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCenRegionBandwidthWithContextV2(ctx context.Context, request *ModifyCenRegionBandwidthRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCenRegionBandwidthRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "ModifyCenRegionBandwidth")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenRegionBandwidthResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCenRegionBandwidthsRequest() (request *DescribeCenRegionBandwidthsRequest) {
	request = &DescribeCenRegionBandwidthsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRegionBandwidths")
	return
}

func NewDescribeCenRegionBandwidthsResponse() (response *DescribeCenRegionBandwidthsResponse) {
	response = &DescribeCenRegionBandwidthsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenRegionBandwidths(request *DescribeCenRegionBandwidthsRequest) string {
	return c.DescribeCenRegionBandwidthsWithContext(context.Background(), request)
}

func (c *Client) DescribeCenRegionBandwidthsSend(request *DescribeCenRegionBandwidthsRequest) (*DescribeCenRegionBandwidthsResponse, error) {
	statusCode, msg, err := c.DescribeCenRegionBandwidthsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCenRegionBandwidthsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCenRegionBandwidthsWithContext(ctx context.Context, request *DescribeCenRegionBandwidthsRequest) string {
	if request == nil {
		request = NewDescribeCenRegionBandwidthsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRegionBandwidths")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCenRegionBandwidthsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCenRegionBandwidthsWithContextV2(ctx context.Context, request *DescribeCenRegionBandwidthsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCenRegionBandwidthsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRegionBandwidths")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCenRegionBandwidthsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCenRoutesRequest() (request *DescribeCenRoutesRequest) {
	request = &DescribeCenRoutesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRoutes")
	return
}

func NewDescribeCenRoutesResponse() (response *DescribeCenRoutesResponse) {
	response = &DescribeCenRoutesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenRoutes(request *DescribeCenRoutesRequest) string {
	return c.DescribeCenRoutesWithContext(context.Background(), request)
}

func (c *Client) DescribeCenRoutesSend(request *DescribeCenRoutesRequest) (*DescribeCenRoutesResponse, error) {
	statusCode, msg, err := c.DescribeCenRoutesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCenRoutesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCenRoutesWithContext(ctx context.Context, request *DescribeCenRoutesRequest) string {
	if request == nil {
		request = NewDescribeCenRoutesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRoutes")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenRoutesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCenRoutesWithContextV2(ctx context.Context, request *DescribeCenRoutesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCenRoutesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRoutes")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenRoutesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCenBandWidthPackageUsageRequest() (request *DescribeCenBandWidthPackageUsageRequest) {
	request = &DescribeCenBandWidthPackageUsageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackageUsage")
	return
}

func NewDescribeCenBandWidthPackageUsageResponse() (response *DescribeCenBandWidthPackageUsageResponse) {
	response = &DescribeCenBandWidthPackageUsageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenBandWidthPackageUsage(request *DescribeCenBandWidthPackageUsageRequest) string {
	return c.DescribeCenBandWidthPackageUsageWithContext(context.Background(), request)
}

func (c *Client) DescribeCenBandWidthPackageUsageSend(request *DescribeCenBandWidthPackageUsageRequest) (*DescribeCenBandWidthPackageUsageResponse, error) {
	statusCode, msg, err := c.DescribeCenBandWidthPackageUsageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeCenBandWidthPackageUsageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCenBandWidthPackageUsageWithContext(ctx context.Context, request *DescribeCenBandWidthPackageUsageRequest) string {
	if request == nil {
		request = NewDescribeCenBandWidthPackageUsageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackageUsage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCenBandWidthPackageUsageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCenBandWidthPackageUsageWithContextV2(ctx context.Context, request *DescribeCenBandWidthPackageUsageRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCenBandWidthPackageUsageRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackageUsage")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCenBandWidthPackageUsageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNetworkInstancesRequest() (request *DescribeNetworkInstancesRequest) {
	request = &DescribeNetworkInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeNetworkInstances")
	return
}

func NewDescribeNetworkInstancesResponse() (response *DescribeNetworkInstancesResponse) {
	response = &DescribeNetworkInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetworkInstances(request *DescribeNetworkInstancesRequest) string {
	return c.DescribeNetworkInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkInstancesSend(request *DescribeNetworkInstancesRequest) (*DescribeNetworkInstancesResponse, error) {
	statusCode, msg, err := c.DescribeNetworkInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeNetworkInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNetworkInstancesWithContext(ctx context.Context, request *DescribeNetworkInstancesRequest) string {
	if request == nil {
		request = NewDescribeNetworkInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeNetworkInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeNetworkInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNetworkInstancesWithContextV2(ctx context.Context, request *DescribeNetworkInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNetworkInstancesRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeNetworkInstances")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeNetworkInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCenGrantRequest() (request *CreateCenGrantRequest) {
	request = &CreateCenGrantRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCenGrant")
	return
}

func NewCreateCenGrantResponse() (response *CreateCenGrantResponse) {
	response = &CreateCenGrantResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCenGrant(request *CreateCenGrantRequest) string {
	return c.CreateCenGrantWithContext(context.Background(), request)
}

func (c *Client) CreateCenGrantSend(request *CreateCenGrantRequest) (*CreateCenGrantResponse, error) {
	statusCode, msg, err := c.CreateCenGrantWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateCenGrantResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCenGrantWithContext(ctx context.Context, request *CreateCenGrantRequest) string {
	if request == nil {
		request = NewCreateCenGrantRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCenGrant")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenGrantResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCenGrantWithContextV2(ctx context.Context, request *CreateCenGrantRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCenGrantRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CreateCenGrant")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenGrantResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInterAreasRequest() (request *DescribeInterAreasRequest) {
	request = &DescribeInterAreasRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeInterAreas")
	return
}

func NewDescribeInterAreasResponse() (response *DescribeInterAreasResponse) {
	response = &DescribeInterAreasResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInterAreas(request *DescribeInterAreasRequest) string {
	return c.DescribeInterAreasWithContext(context.Background(), request)
}

func (c *Client) DescribeInterAreasSend(request *DescribeInterAreasRequest) (*DescribeInterAreasResponse, error) {
	statusCode, msg, err := c.DescribeInterAreasWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInterAreasResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInterAreasWithContext(ctx context.Context, request *DescribeInterAreasRequest) string {
	if request == nil {
		request = NewDescribeInterAreasRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeInterAreas")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInterAreasResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInterAreasWithContextV2(ctx context.Context, request *DescribeInterAreasRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInterAreasRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeInterAreas")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInterAreasResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInterRegionsRequest() (request *DescribeInterRegionsRequest) {
	request = &DescribeInterRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeInterRegions")
	return
}

func NewDescribeInterRegionsResponse() (response *DescribeInterRegionsResponse) {
	response = &DescribeInterRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInterRegions(request *DescribeInterRegionsRequest) string {
	return c.DescribeInterRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeInterRegionsSend(request *DescribeInterRegionsRequest) (*DescribeInterRegionsResponse, error) {
	statusCode, msg, err := c.DescribeInterRegionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeInterRegionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInterRegionsWithContext(ctx context.Context, request *DescribeInterRegionsRequest) string {
	if request == nil {
		request = NewDescribeInterRegionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeInterRegions")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInterRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInterRegionsWithContextV2(ctx context.Context, request *DescribeInterRegionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInterRegionsRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DescribeInterRegions")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInterRegionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachNetworkInstanceRequest() (request *AttachNetworkInstanceRequest) {
	request = &AttachNetworkInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "AttachNetworkInstance")
	return
}

func NewAttachNetworkInstanceResponse() (response *AttachNetworkInstanceResponse) {
	response = &AttachNetworkInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachNetworkInstance(request *AttachNetworkInstanceRequest) string {
	return c.AttachNetworkInstanceWithContext(context.Background(), request)
}

func (c *Client) AttachNetworkInstanceSend(request *AttachNetworkInstanceRequest) (*AttachNetworkInstanceResponse, error) {
	statusCode, msg, err := c.AttachNetworkInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct AttachNetworkInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachNetworkInstanceWithContext(ctx context.Context, request *AttachNetworkInstanceRequest) string {
	if request == nil {
		request = NewAttachNetworkInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "AttachNetworkInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachNetworkInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachNetworkInstanceWithContextV2(ctx context.Context, request *AttachNetworkInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewAttachNetworkInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "AttachNetworkInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachNetworkInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachNetworkInstanceRequest() (request *DetachNetworkInstanceRequest) {
	request = &DetachNetworkInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DetachNetworkInstance")
	return
}

func NewDetachNetworkInstanceResponse() (response *DetachNetworkInstanceResponse) {
	response = &DetachNetworkInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachNetworkInstance(request *DetachNetworkInstanceRequest) string {
	return c.DetachNetworkInstanceWithContext(context.Background(), request)
}

func (c *Client) DetachNetworkInstanceSend(request *DetachNetworkInstanceRequest) (*DetachNetworkInstanceResponse, error) {
	statusCode, msg, err := c.DetachNetworkInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DetachNetworkInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachNetworkInstanceWithContext(ctx context.Context, request *DetachNetworkInstanceRequest) string {
	if request == nil {
		request = NewDetachNetworkInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DetachNetworkInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachNetworkInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachNetworkInstanceWithContextV2(ctx context.Context, request *DetachNetworkInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDetachNetworkInstanceRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "DetachNetworkInstance")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachNetworkInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCenCidrPublishRequest() (request *CenCidrPublishRequest) {
	request = &CenCidrPublishRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CenCidrPublish")
	return
}

func NewCenCidrPublishResponse() (response *CenCidrPublishResponse) {
	response = &CenCidrPublishResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CenCidrPublish(request *CenCidrPublishRequest) string {
	return c.CenCidrPublishWithContext(context.Background(), request)
}

func (c *Client) CenCidrPublishSend(request *CenCidrPublishRequest) (*CenCidrPublishResponse, error) {
	statusCode, msg, err := c.CenCidrPublishWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CenCidrPublishResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CenCidrPublishWithContext(ctx context.Context, request *CenCidrPublishRequest) string {
	if request == nil {
		request = NewCenCidrPublishRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CenCidrPublish")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCenCidrPublishResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CenCidrPublishWithContextV2(ctx context.Context, request *CenCidrPublishRequest) (int, string, error) {
	if request == nil {
		request = NewCenCidrPublishRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CenCidrPublish")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCenCidrPublishResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCenCidrDeleteRequest() (request *CenCidrDeleteRequest) {
	request = &CenCidrDeleteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CenCidrDelete")
	return
}

func NewCenCidrDeleteResponse() (response *CenCidrDeleteResponse) {
	response = &CenCidrDeleteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CenCidrDelete(request *CenCidrDeleteRequest) string {
	return c.CenCidrDeleteWithContext(context.Background(), request)
}

func (c *Client) CenCidrDeleteSend(request *CenCidrDeleteRequest) (*CenCidrDeleteResponse, error) {
	statusCode, msg, err := c.CenCidrDeleteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CenCidrDeleteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CenCidrDeleteWithContext(ctx context.Context, request *CenCidrDeleteRequest) string {
	if request == nil {
		request = NewCenCidrDeleteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CenCidrDelete")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCenCidrDeleteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CenCidrDeleteWithContextV2(ctx context.Context, request *CenCidrDeleteRequest) (int, string, error) {
	if request == nil {
		request = NewCenCidrDeleteRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("cen", APIVersion, "CenCidrDelete")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCenCidrDeleteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
