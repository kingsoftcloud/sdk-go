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
