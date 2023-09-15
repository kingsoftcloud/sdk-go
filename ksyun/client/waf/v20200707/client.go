package v20200707
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func (c *Client) CreateDomain(request *CreateDomainRequest) (string) {
    return c.CreateDomainWithContext(context.Background(), request)
}

func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (string) {
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

func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (string) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (string) {
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

func (c *Client) ModifyDomain(request *ModifyDomainRequest) (string) {
    return c.ModifyDomainWithContext(context.Background(), request)
}

func (c *Client) ModifyDomainWithContext(ctx context.Context, request *ModifyDomainRequest) (string) {
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

func (c *Client) DeleteDomain(request *DeleteDomainRequest) (string) {
    return c.DeleteDomainWithContext(context.Background(), request)
}

func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) (string) {
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

func (c *Client) CreateAccessControlRule(request *CreateAccessControlRuleRequest) (string) {
    return c.CreateAccessControlRuleWithContext(context.Background(), request)
}

func (c *Client) CreateAccessControlRuleWithContext(ctx context.Context, request *CreateAccessControlRuleRequest) (string) {
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

func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) (string) {
    return c.DescribeAccessControlRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeAccessControlRulesWithContext(ctx context.Context, request *DescribeAccessControlRulesRequest) (string) {
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

func (c *Client) ModifyAccessControlRule(request *ModifyAccessControlRuleRequest) (string) {
    return c.ModifyAccessControlRuleWithContext(context.Background(), request)
}

func (c *Client) ModifyAccessControlRuleWithContext(ctx context.Context, request *ModifyAccessControlRuleRequest) (string) {
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

func (c *Client) DeleteAccessControlRule(request *DeleteAccessControlRuleRequest) (string) {
    return c.DeleteAccessControlRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteAccessControlRuleWithContext(ctx context.Context, request *DeleteAccessControlRuleRequest) (string) {
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

func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (string) {
    return c.DescribeCertificatesWithContext(context.Background(), request)
}

func (c *Client) DescribeCertificatesWithContext(ctx context.Context, request *DescribeCertificatesRequest) (string) {
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

func (c *Client) CreateIpv6Protection(request *CreateIpv6ProtectionRequest) (string) {
    return c.CreateIpv6ProtectionWithContext(context.Background(), request)
}

func (c *Client) CreateIpv6ProtectionWithContext(ctx context.Context, request *CreateIpv6ProtectionRequest) (string) {
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

func (c *Client) DeleteIpv6Protection(request *DeleteIpv6ProtectionRequest) (string) {
    return c.DeleteIpv6ProtectionWithContext(context.Background(), request)
}

func (c *Client) DeleteIpv6ProtectionWithContext(ctx context.Context, request *DeleteIpv6ProtectionRequest) (string) {
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


