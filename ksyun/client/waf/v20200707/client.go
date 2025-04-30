package v20200707
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-07-07"

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

func NewCreateWafRequest() (request *CreateWafRequest) {
	request = &CreateWafRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateWaf")
	return
}

func NewCreateWafResponse() (response *CreateWafResponse) {
	response = &CreateWafResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateWaf(request *CreateWafRequest) string {
	return c.CreateWafWithContext(context.Background(), request)
}

func (c *Client) CreateWafWithContext(ctx context.Context, request *CreateWafRequest) string {
	if request == nil {
		request = NewCreateWafRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateWafResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateDomainRequest() (request *CreateDomainRequest) {
	request = &CreateDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateDomain")
	return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
	response = &CreateDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDomain(request *CreateDomainRequest) string {
	return c.CreateDomainWithContext(context.Background(), request)
}

func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) string {
	if request == nil {
		request = NewCreateDomainRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
	request = &DescribeDomainsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomains")
	return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
	response = &DescribeDomainsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDomains(request *DescribeDomainsRequest) string {
	return c.DescribeDomainsWithContext(context.Background(), request)
}

func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) string {
	if request == nil {
		request = NewDescribeDomainsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDomainsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyDomainRequest() (request *ModifyDomainRequest) {
	request = &ModifyDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDomain")
	return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
	response = &ModifyDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDomain(request *ModifyDomainRequest) string {
	return c.ModifyDomainWithContext(context.Background(), request)
}

func (c *Client) ModifyDomainWithContext(ctx context.Context, request *ModifyDomainRequest) string {
	if request == nil {
		request = NewModifyDomainRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
	request = &DeleteDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteDomain")
	return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
	response = &DeleteDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDomain(request *DeleteDomainRequest) string {
	return c.DeleteDomainWithContext(context.Background(), request)
}

func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) string {
	if request == nil {
		request = NewDeleteDomainRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateAccessControlRuleRequest() (request *CreateAccessControlRuleRequest) {
	request = &CreateAccessControlRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateAccessControlRule")
	return
}

func NewCreateAccessControlRuleResponse() (response *CreateAccessControlRuleResponse) {
	response = &CreateAccessControlRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAccessControlRule(request *CreateAccessControlRuleRequest) string {
	return c.CreateAccessControlRuleWithContext(context.Background(), request)
}

func (c *Client) CreateAccessControlRuleWithContext(ctx context.Context, request *CreateAccessControlRuleRequest) string {
	if request == nil {
		request = NewCreateAccessControlRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAccessControlRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAccessControlRulesRequest() (request *DescribeAccessControlRulesRequest) {
	request = &DescribeAccessControlRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessControlRules")
	return
}

func NewDescribeAccessControlRulesResponse() (response *DescribeAccessControlRulesResponse) {
	response = &DescribeAccessControlRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) string {
	return c.DescribeAccessControlRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeAccessControlRulesWithContext(ctx context.Context, request *DescribeAccessControlRulesRequest) string {
	if request == nil {
		request = NewDescribeAccessControlRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAccessControlRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyAccessControlRuleRequest() (request *ModifyAccessControlRuleRequest) {
	request = &ModifyAccessControlRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAccessControlRule")
	return
}

func NewModifyAccessControlRuleResponse() (response *ModifyAccessControlRuleResponse) {
	response = &ModifyAccessControlRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAccessControlRule(request *ModifyAccessControlRuleRequest) string {
	return c.ModifyAccessControlRuleWithContext(context.Background(), request)
}

func (c *Client) ModifyAccessControlRuleWithContext(ctx context.Context, request *ModifyAccessControlRuleRequest) string {
	if request == nil {
		request = NewModifyAccessControlRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAccessControlRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteAccessControlRuleRequest() (request *DeleteAccessControlRuleRequest) {
	request = &DeleteAccessControlRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAccessControlRule")
	return
}

func NewDeleteAccessControlRuleResponse() (response *DeleteAccessControlRuleResponse) {
	response = &DeleteAccessControlRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAccessControlRule(request *DeleteAccessControlRuleRequest) string {
	return c.DeleteAccessControlRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteAccessControlRuleWithContext(ctx context.Context, request *DeleteAccessControlRuleRequest) string {
	if request == nil {
		request = NewDeleteAccessControlRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAccessControlRuleResponse()
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
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCertificates")
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
func NewCreateIpv6ProtectionRequest() (request *CreateIpv6ProtectionRequest) {
	request = &CreateIpv6ProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateIpv6Protection")
	return
}

func NewCreateIpv6ProtectionResponse() (response *CreateIpv6ProtectionResponse) {
	response = &CreateIpv6ProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateIpv6Protection(request *CreateIpv6ProtectionRequest) string {
	return c.CreateIpv6ProtectionWithContext(context.Background(), request)
}

func (c *Client) CreateIpv6ProtectionWithContext(ctx context.Context, request *CreateIpv6ProtectionRequest) string {
	if request == nil {
		request = NewCreateIpv6ProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateIpv6ProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteIpv6ProtectionRequest() (request *DeleteIpv6ProtectionRequest) {
	request = &DeleteIpv6ProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteIpv6Protection")
	return
}

func NewDeleteIpv6ProtectionResponse() (response *DeleteIpv6ProtectionResponse) {
	response = &DeleteIpv6ProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteIpv6Protection(request *DeleteIpv6ProtectionRequest) string {
	return c.DeleteIpv6ProtectionWithContext(context.Background(), request)
}

func (c *Client) DeleteIpv6ProtectionWithContext(ctx context.Context, request *DeleteIpv6ProtectionRequest) string {
	if request == nil {
		request = NewDeleteIpv6ProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteIpv6ProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyStorageTimeRequest() (request *ModifyStorageTimeRequest) {
	request = &ModifyStorageTimeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyStorageTime")
	return
}

func NewModifyStorageTimeResponse() (response *ModifyStorageTimeResponse) {
	response = &ModifyStorageTimeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyStorageTime(request *ModifyStorageTimeRequest) string {
	return c.ModifyStorageTimeWithContext(context.Background(), request)
}

func (c *Client) ModifyStorageTimeWithContext(ctx context.Context, request *ModifyStorageTimeRequest) string {
	if request == nil {
		request = NewModifyStorageTimeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyStorageTimeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateAlbDomainRequest() (request *CreateAlbDomainRequest) {
	request = &CreateAlbDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateAlbDomain")
	return
}

func NewCreateAlbDomainResponse() (response *CreateAlbDomainResponse) {
	response = &CreateAlbDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlbDomain(request *CreateAlbDomainRequest) string {
	return c.CreateAlbDomainWithContext(context.Background(), request)
}

func (c *Client) CreateAlbDomainWithContext(ctx context.Context, request *CreateAlbDomainRequest) string {
	if request == nil {
		request = NewCreateAlbDomainRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyAlbDomainRequest() (request *ModifyAlbDomainRequest) {
	request = &ModifyAlbDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAlbDomain")
	return
}

func NewModifyAlbDomainResponse() (response *ModifyAlbDomainResponse) {
	response = &ModifyAlbDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAlbDomain(request *ModifyAlbDomainRequest) string {
	return c.ModifyAlbDomainWithContext(context.Background(), request)
}

func (c *Client) ModifyAlbDomainWithContext(ctx context.Context, request *ModifyAlbDomainRequest) string {
	if request == nil {
		request = NewModifyAlbDomainRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteAlbDomainRequest() (request *DeleteAlbDomainRequest) {
	request = &DeleteAlbDomainRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAlbDomain")
	return
}

func NewDeleteAlbDomainResponse() (response *DeleteAlbDomainResponse) {
	response = &DeleteAlbDomainResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlbDomain(request *DeleteAlbDomainRequest) string {
	return c.DeleteAlbDomainWithContext(context.Background(), request)
}

func (c *Client) DeleteAlbDomainWithContext(ctx context.Context, request *DeleteAlbDomainRequest) string {
	if request == nil {
		request = NewDeleteAlbDomainRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAlbDomainResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateAlbWafRequest() (request *CreateAlbWafRequest) {
	request = &CreateAlbWafRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateAlbWaf")
	return
}

func NewCreateAlbWafResponse() (response *CreateAlbWafResponse) {
	response = &CreateAlbWafResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlbWaf(request *CreateAlbWafRequest) string {
	return c.CreateAlbWafWithContext(context.Background(), request)
}

func (c *Client) CreateAlbWafWithContext(ctx context.Context, request *CreateAlbWafRequest) string {
	if request == nil {
		request = NewCreateAlbWafRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlbWafResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyAlbWafRequest() (request *ModifyAlbWafRequest) {
	request = &ModifyAlbWafRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAlbWaf")
	return
}

func NewModifyAlbWafResponse() (response *ModifyAlbWafResponse) {
	response = &ModifyAlbWafResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyAlbWaf(request *ModifyAlbWafRequest) string {
	return c.ModifyAlbWafWithContext(context.Background(), request)
}

func (c *Client) ModifyAlbWafWithContext(ctx context.Context, request *ModifyAlbWafRequest) string {
	if request == nil {
		request = NewModifyAlbWafRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyAlbWafResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAlbCertificatesRequest() (request *DescribeAlbCertificatesRequest) {
	request = &DescribeAlbCertificatesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAlbCertificates")
	return
}

func NewDescribeAlbCertificatesResponse() (response *DescribeAlbCertificatesResponse) {
	response = &DescribeAlbCertificatesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbCertificates(request *DescribeAlbCertificatesRequest) string {
	return c.DescribeAlbCertificatesWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbCertificatesWithContext(ctx context.Context, request *DescribeAlbCertificatesRequest) string {
	if request == nil {
		request = NewDescribeAlbCertificatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbCertificatesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeAlbDomainsRequest() (request *DescribeAlbDomainsRequest) {
	request = &DescribeAlbDomainsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAlbDomains")
	return
}

func NewDescribeAlbDomainsResponse() (response *DescribeAlbDomainsResponse) {
	response = &DescribeAlbDomainsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAlbDomains(request *DescribeAlbDomainsRequest) string {
	return c.DescribeAlbDomainsWithContext(context.Background(), request)
}

func (c *Client) DescribeAlbDomainsWithContext(ctx context.Context, request *DescribeAlbDomainsRequest) string {
	if request == nil {
		request = NewDescribeAlbDomainsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAlbDomainsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


