package v20200707
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreateWafRequest struct {
	*ksyunhttp.BaseRequest
	WafType       *string `json:"WafType,omitempty" name:"WafType"`
	BillType      *int    `json:"BillType,omitempty" name:"BillType"`
	PurchaseTime  *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	DomainBag     *int    `json:"DomainBag,omitempty" name:"DomainBag"`
	BandWidthIpv4 *int    `json:"BandWidthIpv4,omitempty" name:"BandWidthIpv4"`
	StorageSize   *int    `json:"StorageSize,omitempty" name:"StorageSize"`
	ProjectId     *int    `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateWafRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateWafResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	WafInfo   *string `json:"WafInfo" name:"WafInfo"`
}

func (r *CreateWafResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWafResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteWafRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DeleteWafRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteWafResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteWafResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWafResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateDomainRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecord    *string   `json:"ResourceRecord,omitempty" name:"ResourceRecord"`
	HttpRewrite       *bool     `json:"HttpRewrite,omitempty" name:"HttpRewrite"`
	HttpSource        *bool     `json:"HttpSource,omitempty" name:"HttpSource"`
	CertificateId     *string   `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateRegion *string   `json:"CertificateRegion,omitempty" name:"CertificateRegion"`
	LbMethod          *string   `json:"LbMethod,omitempty" name:"LbMethod"`
	HasProxy          *bool     `json:"HasProxy,omitempty" name:"HasProxy"`
	ProjectId         *int      `json:"ProjectId,omitempty" name:"ProjectId"`
	HeaderMark        *string   `json:"HeaderMark,omitempty" name:"HeaderMark"`
	HeaderValue       *string   `json:"HeaderValue,omitempty" name:"HeaderValue"`
	HealthMonitor     *string   `json:"HealthMonitor,omitempty" name:"HealthMonitor"`
	HttpPort          []*string `json:"HttpPort,omitempty" name:"HttpPort"`
	HttpsPort         []*string `json:"HttpsPort,omitempty" name:"HttpsPort"`
	Sources           *string   `json:"Sources,omitempty" name:"Sources"`
}

func (r *CreateDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateDomainResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDomainsRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecord *string `json:"ResourceRecord,omitempty" name:"ResourceRecord"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDomainsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyDomainRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId  *string   `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
	HttpRewrite       *bool     `json:"HttpRewrite,omitempty" name:"HttpRewrite"`
	HttpSource        *bool     `json:"HttpSource,omitempty" name:"HttpSource"`
	CertificateId     *string   `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateRegion *string   `json:"CertificateRegion,omitempty" name:"CertificateRegion"`
	LbMethod          *string   `json:"LbMethod,omitempty" name:"LbMethod"`
	HasProxy          *bool     `json:"HasProxy,omitempty" name:"HasProxy"`
	ProjectId         *int      `json:"ProjectId,omitempty" name:"ProjectId"`
	HeaderMark        *string   `json:"HeaderMark,omitempty" name:"HeaderMark"`
	HeaderValue       *string   `json:"HeaderValue,omitempty" name:"HeaderValue"`
	HealthMonitor     *string   `json:"HealthMonitor,omitempty" name:"HealthMonitor"`
	HttpPort          []*string `json:"HttpPort,omitempty" name:"HttpPort"`
	HttpsPort         []*string `json:"HttpsPort,omitempty" name:"HttpsPort"`
	Sources           *string   `json:"Sources,omitempty" name:"Sources"`
}

func (r *ModifyDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDomainResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteDomainRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId *string `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDomainResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateAccessControlRuleRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId *string `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
	RuleName         *string `json:"RuleName,omitempty" name:"RuleName"`
	Level            *int    `json:"Level,omitempty" name:"Level"`
	RuleAction       *int    `json:"RuleAction,omitempty" name:"RuleAction"`
	Status           *bool   `json:"Status,omitempty" name:"Status"`
	MatchRuleInfo    *string `json:"MatchRuleInfo,omitempty" name:"MatchRuleInfo"`
	Priority         *int    `json:"Priority,omitempty" name:"Priority"`
}

func (r *CreateAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAccessControlRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	AccessControlRule struct {
		RuleType         *string `json:"RuleType" name:"RuleType"`
		MatchRule        *int    `json:"MatchRule" name:"MatchRule"`
		ArgName          *string `json:"ArgName" name:"ArgName"`
		RuleData         *string `json:"RuleData" name:"RuleData"`
		Status           *bool   `json:"Status" name:"Status"`
		Priority         *int    `json:"Priority" name:"Priority"`
		RuleAction       *int    `json:"RuleAction" name:"RuleAction"`
		MatchRuleInfo    *string `json:"MatchRuleInfo" name:"MatchRuleInfo"`
		Level            *int    `json:"Level" name:"Level"`
		ResourceRecordId *string `json:"ResourceRecordId" name:"ResourceRecordId"`
		RuleName         *string `json:"RuleName" name:"RuleName"`
		RuleId           *string `json:"RuleId" name:"RuleId"`
	} `json:"AccessControlRule"`
}

func (r *CreateAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeAccessControlRulesRequest struct {
	*ksyunhttp.BaseRequest
	RuleId           *string `json:"RuleId,omitempty" name:"RuleId"`
	ResourceRecordId *string `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
	RuleName         *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeAccessControlRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAccessControlRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeAccessControlRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyAccessControlRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleId     *string `json:"RuleId,omitempty" name:"RuleId"`
	RuleName   *string `json:"RuleName,omitempty" name:"RuleName"`
	RuleType   *string `json:"RuleType,omitempty" name:"RuleType"`
	RuleData   *string `json:"RuleData,omitempty" name:"RuleData"`
	MatchRule  *string `json:"MatchRule,omitempty" name:"MatchRule"`
	ArgName    *string `json:"ArgName,omitempty" name:"ArgName"`
	Level      *int    `json:"Level,omitempty" name:"Level"`
	RuleAction *int    `json:"RuleAction,omitempty" name:"RuleAction"`
	Status     *bool   `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyAccessControlRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteAccessControlRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteAccessControlRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeCertificatesRequest struct {
	*ksyunhttp.BaseRequest
	Request *string `json:"Request,omitempty" name:"Request"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCertificatesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateIpv6ProtectionRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId []*string `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
}

func (r *CreateIpv6ProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateIpv6ProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateIpv6ProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIpv6ProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteIpv6ProtectionRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId []*string `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
}

func (r *DeleteIpv6ProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteIpv6ProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteIpv6ProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpv6ProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyStorageTimeRequest struct {
	*ksyunhttp.BaseRequest
	StorageTime *int `json:"StorageTime,omitempty" name:"StorageTime"`
}

func (r *ModifyStorageTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyStorageTimeResponse struct {
	*ksyunhttp.BaseResponse
	StorageTime *int `json:"StorageTime" name:"StorageTime"`
}

func (r *ModifyStorageTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStorageTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateAlbDomainRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecord    *string   `json:"ResourceRecord,omitempty" name:"ResourceRecord"`
	HttpRewrite       *bool     `json:"HttpRewrite,omitempty" name:"HttpRewrite"`
	HttpSource        *bool     `json:"HttpSource,omitempty" name:"HttpSource"`
	CertificateId     *string   `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateRegion *string   `json:"CertificateRegion,omitempty" name:"CertificateRegion"`
	LbMethod          *string   `json:"LbMethod,omitempty" name:"LbMethod"`
	HasProxy          *bool     `json:"HasProxy,omitempty" name:"HasProxy"`
	ProjectId         *int      `json:"ProjectId,omitempty" name:"ProjectId"`
	HeaderMark        *string   `json:"HeaderMark,omitempty" name:"HeaderMark"`
	HeaderValue       *string   `json:"HeaderValue,omitempty" name:"HeaderValue"`
	HealthMonitor     *string   `json:"HealthMonitor,omitempty" name:"HealthMonitor"`
	HttpPort          []*string `json:"HttpPort,omitempty" name:"HttpPort"`
	HttpsPort         []*string `json:"HttpsPort,omitempty" name:"HttpsPort"`
	Sources           *string   `json:"Sources,omitempty" name:"Sources"`
}

func (r *CreateAlbDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAlbDomainResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateAlbDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyAlbDomainRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId  *string   `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
	HttpRewrite       *bool     `json:"HttpRewrite,omitempty" name:"HttpRewrite"`
	HttpSource        *bool     `json:"HttpSource,omitempty" name:"HttpSource"`
	CertificateId     *string   `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateRegion *string   `json:"CertificateRegion,omitempty" name:"CertificateRegion"`
	LbMethod          *string   `json:"LbMethod,omitempty" name:"LbMethod"`
	HasProxy          *bool     `json:"HasProxy,omitempty" name:"HasProxy"`
	ProjectId         *int      `json:"ProjectId,omitempty" name:"ProjectId"`
	HeaderMark        *string   `json:"HeaderMark,omitempty" name:"HeaderMark"`
	HeaderValue       *string   `json:"HeaderValue,omitempty" name:"HeaderValue"`
	HealthMonitor     *string   `json:"HealthMonitor,omitempty" name:"HealthMonitor"`
	HttpPort          []*string `json:"HttpPort,omitempty" name:"HttpPort"`
	HttpsPort         []*string `json:"HttpsPort,omitempty" name:"HttpsPort"`
	Sources           *string   `json:"Sources,omitempty" name:"Sources"`
}

func (r *ModifyAlbDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyAlbDomainResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyAlbDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteAlbDomainRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecordId *string `json:"ResourceRecordId,omitempty" name:"ResourceRecordId"`
}

func (r *DeleteAlbDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteAlbDomainResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteAlbDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbWafRequest struct {
	*ksyunhttp.BaseRequest
	WafType      *string `json:"WafType,omitempty" name:"WafType"`
	BillType     *int    `json:"BillType,omitempty" name:"BillType"`
	ProjectId    *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	AlbNumber    *int    `json:"AlbNumber,omitempty" name:"AlbNumber"`
	DomainBag    *int    `json:"DomainBag,omitempty" name:"DomainBag"`
	QpsBag       *int    `json:"QpsBag,omitempty" name:"QpsBag"`
	PurchaseTime *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *CreateAlbWafRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAlbWafResponse struct {
	*ksyunhttp.BaseResponse
	WafInfo struct {
	} `json:"WafInfo"`
}

func (r *CreateAlbWafResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbWafResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbWafRequest struct {
	*ksyunhttp.BaseRequest
	WafType   *string `json:"WafType,omitempty" name:"WafType"`
	DomainBag *int    `json:"DomainBag,omitempty" name:"DomainBag"`
	AlbNumber *int    `json:"AlbNumber,omitempty" name:"AlbNumber"`
	QpsBag    *int    `json:"QpsBag,omitempty" name:"QpsBag"`
}

func (r *ModifyAlbWafRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyAlbWafResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ModifyAlbWafResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbWafResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbDomainsRequest struct {
	*ksyunhttp.BaseRequest
	ResourceRecord *string `json:"ResourceRecord,omitempty" name:"ResourceRecord"`
}

func (r *DescribeAlbDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAlbDomainsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeAlbDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

