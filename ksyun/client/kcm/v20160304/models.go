package v20160304
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type ApplyCertificateRequest struct {
    *ksyunhttp.BaseRequest
    MainDomain *string `json:"MainDomain,omitempty" name:"MainDomain"`
    CertificateCode *string `json:"CertificateCode,omitempty" name:"CertificateCode"`
    YearLength *int `json:"YearLength,omitempty" name:"YearLength"`
    DomainCount *int `json:"DomainCount,omitempty" name:"DomainCount"`
    WildcardCount *int `json:"WildcardCount,omitempty" name:"WildcardCount"`
    ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
    SubOrderId *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
    ProjectId *int `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ApplyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCertificateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ApplyCertificateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ApplyCertificateResponse struct {
    *ksyunhttp.BaseResponse
	ApplyCertificateResponse struct {
		RequestId *string `json:"RequestId"`
		Certificate struct {
				CertificateId *string `json:"CertificateId"`
				MainDomain *string `json:"MainDomain"`
				CertificateBrand *string `json:"CertificateBrand"`
				CertificateLevel *string `json:"CertificateLevel"`
				CertificateName *string `json:"CertificateName"`
				CertificateCode *string `json:"CertificateCode"`
				CertificateStatus *string `json:"CertificateStatus"`
				YearLength *string `json:"YearLength"`
				DomainCount *string `json:"DomainCount"`
				WildcardCount *string `json:"WildcardCount"`
		} `json:"Certificate"`
	} `json:"ApplyCertificateResponse"`
}

func (r *ApplyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCertificateRequest struct {
    *ksyunhttp.BaseRequest
    CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
    AuthMethod *string `json:"AuthMethod,omitempty" name:"AuthMethod"`
    CSR *string `json:"CSR,omitempty" name:"CSR"`
    ContactId *int `json:"ContactId,omitempty" name:"ContactId"`
    CompanyId *int `json:"CompanyId,omitempty" name:"CompanyId"`
    CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`
    Department *string `json:"Department,omitempty" name:"Department"`
    State *string `json:"State,omitempty" name:"State"`
    City *string `json:"City,omitempty" name:"City"`
    Address *string `json:"Address,omitempty" name:"Address"`
    CompanyPhone *string `json:"CompanyPhone,omitempty" name:"CompanyPhone"`
    PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`
    DcvEmail *string `json:"DcvEmail,omitempty" name:"DcvEmail"`
    AdditionalDomains *string `json:"AdditionalDomains,omitempty" name:"AdditionalDomains"`
    Wildcards *string `json:"Wildcards,omitempty" name:"Wildcards"`
    Contact *string `json:"Contact,omitempty" name:"Contact"`
    IsSubmit *string `json:"IsSubmit,omitempty" name:"IsSubmit"`
    BusinessLicence *string `json:"BusinessLicence,omitempty" name:"BusinessLicence"`
    CsrSource *string `json:"CsrSource,omitempty" name:"CsrSource"`
    Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

func (r *UpdateCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCertificateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateCertificateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCertificateResponse struct {
    *ksyunhttp.BaseResponse
	UpdateCertificateResponse struct {
		RequestId *string `json:"RequestId"`
		Certificate struct {
				CertificateId *string `json:"CertificateId"`
				MainDomain *string `json:"MainDomain"`
				CertificateBrand *string `json:"CertificateBrand"`
				CertificateLevel *string `json:"CertificateLevel"`
				CertificateName *string `json:"CertificateName"`
				CertificateCode *string `json:"CertificateCode"`
				CertificateStatus *string `json:"CertificateStatus"`
				YearLength *string `json:"YearLength"`
				DomainCount *string `json:"DomainCount"`
				WildcardCount *string `json:"WildcardCount"`
				AuthMethod *string `json:"AuthMethod"`
				CSR *string `json:"CSR"`
				State *string `json:"State"`
				City *string `json:"City"`
				Address *string `json:"Address"`
				CompanyPhone *string `json:"CompanyPhone"`
				PostalCode *string `json:"PostalCode"`
				AdditionalDomains *string `json:"AdditionalDomains"`
				Wildcards *string `json:"Wildcards"`
		} `json:"Certificate"`
		Contact struct {
				FirstName *string `json:"FirstName"`
				LastName *string `json:"LastName"`
				Phone *string `json:"Phone"`
				Email *string `json:"Email"`
				Title *string `json:"Title"`
		} `json:"Contact"`
	} `json:"UpdateCertificateResponse"`
}

func (r *UpdateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCertificatesRequest struct {
    *ksyunhttp.BaseRequest
    CertificateId []*string `json:"CertificateId,omitempty" name:"CertificateId"`
    ProjectId []*int `json:"ProjectId,omitempty" name:"ProjectId"`
    Filter []*string `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCertificatesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListCertificatesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListCertificatesResponse struct {
    *ksyunhttp.BaseResponse
	ListCertificatesResponse struct {
		RequestId *string `json:"RequestId"`
		CertificateSet struct {
			Item struct {
				CertificateId *string `json:"CertificateId"`
				MainDomain *string `json:"MainDomain"`
				CertificateBrand *string `json:"CertificateBrand"`
				CertificateLevel *string `json:"CertificateLevel"`
				CertificateName *string `json:"CertificateName"`
				CertificateCode *string `json:"CertificateCode"`
				CertificateStatus *string `json:"CertificateStatus"`
				YearLength *string `json:"YearLength"`
				DomainCount *string `json:"DomainCount"`
				WildcardCount *string `json:"WildcardCount"`
				ExpireTime *string `json:"ExpireTime"`
				AuthMethod *string `json:"AuthMethod"`
				AuthContent struct {
						CheckName *string `json:"CheckName"`
						CheckValue *string `json:"CheckValue"`
						RecordType *string `json:"RecordType"`
				} `json:"AuthContent"`
			} `json:"Item"`
		} `json:"CertificateSet"`
	} `json:"ListCertificatesResponse"`
}

func (r *ListCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCertificatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCertificateDetailRequest struct {
    *ksyunhttp.BaseRequest
    CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *GetCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCertificateDetailRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetCertificateDetailRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetCertificateDetailResponse struct {
    *ksyunhttp.BaseResponse
	GetCertificateDetailResponse struct {
		RequestId *string `json:"RequestId"`
		Certificate struct {
				CertificateId *string `json:"CertificateId"`
				MainDomain *string `json:"MainDomain"`
				CertificateBrand *string `json:"CertificateBrand"`
				CertificateLevel *string `json:"CertificateLevel"`
				CertificateName *string `json:"CertificateName"`
				CertificateCode *string `json:"CertificateCode"`
				CertificateStatus *string `json:"CertificateStatus"`
				YearLength *string `json:"YearLength"`
				DomainCount *string `json:"DomainCount"`
				WildcardCount *string `json:"WildcardCount"`
				Department *string `json:"Department"`
				CSR *string `json:"CSR"`
				State *string `json:"State"`
				City *string `json:"City"`
				Address *string `json:"Address"`
				CompanyPhone *string `json:"CompanyPhone"`
				PostalCode *string `json:"PostalCode"`
				AdditionalDomains *string `json:"AdditionalDomains"`
				Wildcards *string `json:"Wildcards"`
			Contact struct {
				FirstName *string `json:"FirstName"`
				LastName *string `json:"LastName"`
				Phone *string `json:"Phone"`
				Email *string `json:"Email"`
				Title *string `json:"Title"`
			} `json:"Contact"`
		} `json:"Certificate"`
	} `json:"GetCertificateDetailResponse"`
}

func (r *GetCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCertificateDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

