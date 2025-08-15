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

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
	request = &CreateCertificateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "CreateCertificate")
	return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
	response = &CreateCertificateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCertificate(request *CreateCertificateRequest) string {
	return c.CreateCertificateWithContext(context.Background(), request)
}

func (c *Client) CreateCertificateSend(request *CreateCertificateRequest) (*CreateCertificateResponse, error) {
	statusCode, msg, err := c.CreateCertificateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateCertificateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCertificateWithContext(ctx context.Context, request *CreateCertificateRequest) string {
	if request == nil {
		request = NewCreateCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCertificateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCertificateWithContextV2(ctx context.Context, request *CreateCertificateRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCertificateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
	request = &DeleteCertificateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "DeleteCertificate")
	return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
	response = &DeleteCertificateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) string {
	return c.DeleteCertificateWithContext(context.Background(), request)
}

func (c *Client) DeleteCertificateSend(request *DeleteCertificateRequest) (*DeleteCertificateResponse, error) {
	statusCode, msg, err := c.DeleteCertificateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteCertificateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCertificateWithContext(ctx context.Context, request *DeleteCertificateRequest) string {
	if request == nil {
		request = NewDeleteCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCertificateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCertificateWithContextV2(ctx context.Context, request *DeleteCertificateRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCertificateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCertificateRequest() (request *ModifyCertificateRequest) {
	request = &ModifyCertificateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "ModifyCertificate")
	return
}

func NewModifyCertificateResponse() (response *ModifyCertificateResponse) {
	response = &ModifyCertificateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCertificate(request *ModifyCertificateRequest) string {
	return c.ModifyCertificateWithContext(context.Background(), request)
}

func (c *Client) ModifyCertificateSend(request *ModifyCertificateRequest) (*ModifyCertificateResponse, error) {
	statusCode, msg, err := c.ModifyCertificateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyCertificateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCertificateWithContext(ctx context.Context, request *ModifyCertificateRequest) string {
	if request == nil {
		request = NewModifyCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCertificateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCertificateWithContextV2(ctx context.Context, request *ModifyCertificateRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCertificateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
	request = &DescribeCertificatesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "DescribeCertificates")
	return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
	response = &DescribeCertificatesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) string {
	return c.DescribeCertificatesWithContext(context.Background(), request)
}

func (c *Client) DescribeCertificatesSend(request *DescribeCertificatesRequest) (*DescribeCertificatesResponse, error) {
	statusCode, msg, err := c.DescribeCertificatesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCertificatesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCertificatesWithContext(ctx context.Context, request *DescribeCertificatesRequest) string {
	if request == nil {
		request = NewDescribeCertificatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCertificatesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCertificatesWithContextV2(ctx context.Context, request *DescribeCertificatesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCertificatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCertificatesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewApplyCertificateRequest() (request *ApplyCertificateRequest) {
	request = &ApplyCertificateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "ApplyCertificate")
	return
}

func NewApplyCertificateResponse() (response *ApplyCertificateResponse) {
	response = &ApplyCertificateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ApplyCertificate(request *ApplyCertificateRequest) string {
	return c.ApplyCertificateWithContext(context.Background(), request)
}

func (c *Client) ApplyCertificateSend(request *ApplyCertificateRequest) (*ApplyCertificateResponse, error) {
	statusCode, msg, err := c.ApplyCertificateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ApplyCertificateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ApplyCertificateWithContext(ctx context.Context, request *ApplyCertificateRequest) string {
	if request == nil {
		request = NewApplyCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewApplyCertificateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ApplyCertificateWithContextV2(ctx context.Context, request *ApplyCertificateRequest) (int, string, error) {
	if request == nil {
		request = NewApplyCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewApplyCertificateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateCertificateRequest() (request *UpdateCertificateRequest) {
	request = &UpdateCertificateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "UpdateCertificate")
	return
}

func NewUpdateCertificateResponse() (response *UpdateCertificateResponse) {
	response = &UpdateCertificateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateCertificate(request *UpdateCertificateRequest) string {
	return c.UpdateCertificateWithContext(context.Background(), request)
}

func (c *Client) UpdateCertificateSend(request *UpdateCertificateRequest) (*UpdateCertificateResponse, error) {
	statusCode, msg, err := c.UpdateCertificateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateCertificateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateCertificateWithContext(ctx context.Context, request *UpdateCertificateRequest) string {
	if request == nil {
		request = NewUpdateCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateCertificateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateCertificateWithContextV2(ctx context.Context, request *UpdateCertificateRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateCertificateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewReIssueCertificateRequest() (request *ReIssueCertificateRequest) {
	request = &ReIssueCertificateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "ReIssueCertificate")
	return
}

func NewReIssueCertificateResponse() (response *ReIssueCertificateResponse) {
	response = &ReIssueCertificateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ReIssueCertificate(request *ReIssueCertificateRequest) string {
	return c.ReIssueCertificateWithContext(context.Background(), request)
}

func (c *Client) ReIssueCertificateSend(request *ReIssueCertificateRequest) (*ReIssueCertificateResponse, error) {
	statusCode, msg, err := c.ReIssueCertificateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ReIssueCertificateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ReIssueCertificateWithContext(ctx context.Context, request *ReIssueCertificateRequest) string {
	if request == nil {
		request = NewReIssueCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReIssueCertificateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ReIssueCertificateWithContextV2(ctx context.Context, request *ReIssueCertificateRequest) (int, string, error) {
	if request == nil {
		request = NewReIssueCertificateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewReIssueCertificateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCancelTransactionRequest() (request *CancelTransactionRequest) {
	request = &CancelTransactionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "CancelTransaction")
	return
}

func NewCancelTransactionResponse() (response *CancelTransactionResponse) {
	response = &CancelTransactionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelTransaction(request *CancelTransactionRequest) string {
	return c.CancelTransactionWithContext(context.Background(), request)
}

func (c *Client) CancelTransactionSend(request *CancelTransactionRequest) (*CancelTransactionResponse, error) {
	statusCode, msg, err := c.CancelTransactionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelTransactionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CancelTransactionWithContext(ctx context.Context, request *CancelTransactionRequest) string {
	if request == nil {
		request = NewCancelTransactionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelTransactionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CancelTransactionWithContextV2(ctx context.Context, request *CancelTransactionRequest) (int, string, error) {
	if request == nil {
		request = NewCancelTransactionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCancelTransactionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewListCertificatesRequest() (request *ListCertificatesRequest) {
	request = &ListCertificatesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "ListCertificates")
	return
}

func NewListCertificatesResponse() (response *ListCertificatesResponse) {
	response = &ListCertificatesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListCertificates(request *ListCertificatesRequest) string {
	return c.ListCertificatesWithContext(context.Background(), request)
}

func (c *Client) ListCertificatesSend(request *ListCertificatesRequest) (*ListCertificatesResponse, error) {
	statusCode, msg, err := c.ListCertificatesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListCertificatesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListCertificatesWithContext(ctx context.Context, request *ListCertificatesRequest) string {
	if request == nil {
		request = NewListCertificatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListCertificatesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListCertificatesWithContextV2(ctx context.Context, request *ListCertificatesRequest) (int, string, error) {
	if request == nil {
		request = NewListCertificatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListCertificatesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetCertificateDetailRequest() (request *GetCertificateDetailRequest) {
	request = &GetCertificateDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcm", APIVersion, "GetCertificateDetail")
	return
}

func NewGetCertificateDetailResponse() (response *GetCertificateDetailResponse) {
	response = &GetCertificateDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetCertificateDetail(request *GetCertificateDetailRequest) string {
	return c.GetCertificateDetailWithContext(context.Background(), request)
}

func (c *Client) GetCertificateDetailSend(request *GetCertificateDetailRequest) (*GetCertificateDetailResponse, error) {
	statusCode, msg, err := c.GetCertificateDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetCertificateDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetCertificateDetailWithContext(ctx context.Context, request *GetCertificateDetailRequest) string {
	if request == nil {
		request = NewGetCertificateDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetCertificateDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetCertificateDetailWithContextV2(ctx context.Context, request *GetCertificateDetailRequest) (int, string, error) {
	if request == nil {
		request = NewGetCertificateDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetCertificateDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
