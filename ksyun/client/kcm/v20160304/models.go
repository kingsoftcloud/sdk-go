package v20160304

import (
	"encoding/json"

	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateCertificateRequest struct {
	*ksyunhttp.BaseRequest
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`
	PrivateKey      *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	PublicKey       *string `json:"PublicKey,omitempty" name:"PublicKey"`
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`
}

func (r *CreateCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCertificateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Certificate struct {
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		CertificateName *string `json:"CertificateName" name:"CertificateName"`
		CertificateId   *string `json:"CertificateId" name:"CertificateId"`
		ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		CommonName      *string `json:"CommonName" name:"CommonName"`
		CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
	} `json:"Certificate"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId   *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`
	PrivateKey      *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	PublicKey       *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

func (r *ModifyCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Certificate struct {
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		CertificateName *string `json:"CertificateName" name:"CertificateName"`
		CertificateId   *string `json:"CertificateId" name:"CertificateId"`
		ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		CommonName      *string `json:"CommonName" name:"CommonName"`
		CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
	} `json:"Certificate"`
}

func (r *ModifyCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId []*string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CertificateSet []struct {
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		CertificateName *string `json:"CertificateName" name:"CertificateName"`
		CertificateId   *string `json:"CertificateId" name:"CertificateId"`
		ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		CommonName      *string `json:"CommonName" name:"CommonName"`
		CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
		CertType        *string `json:"CertType" name:"CertType"`
		CertificateType *string `json:"CertificateType" name:"CertificateType"`
		PublicKey       *string `json:"PublicKey" name:"PublicKey"`
	} `json:"CertificateSet"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyCertificateRequest struct {
	*ksyunhttp.BaseRequest
	MainDomain      *string `json:"MainDomain,omitempty" name:"MainDomain"`
	CertificateCode *string `json:"CertificateCode,omitempty" name:"CertificateCode"`
	YearLength      *int    `json:"YearLength,omitempty" name:"YearLength"`
	DomainCount     *int    `json:"DomainCount,omitempty" name:"DomainCount"`
	WildcardCount   *int    `json:"WildcardCount,omitempty" name:"WildcardCount"`
	ProductId       *string `json:"ProductId,omitempty" name:"ProductId"`
	SubOrderId      *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
	ProjectId       *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	BillType        *int    `json:"BillType,omitempty" name:"BillType"`
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
		RequestId   *string `json:"RequestId" name:"RequestId"`
		Certificate struct {
			CertificateId     *string `json:"CertificateId" name:"CertificateId"`
			MainDomain        *string `json:"MainDomain" name:"MainDomain"`
			CertificateBrand  *string `json:"CertificateBrand" name:"CertificateBrand"`
			CertificateLevel  *string `json:"CertificateLevel" name:"CertificateLevel"`
			CertificateName   *string `json:"CertificateName" name:"CertificateName"`
			CertificateCode   *string `json:"CertificateCode" name:"CertificateCode"`
			CertificateStatus *string `json:"CertificateStatus" name:"CertificateStatus"`
			YearLength        *string `json:"YearLength" name:"YearLength"`
			DomainCount       *string `json:"DomainCount" name:"DomainCount"`
			WildcardCount     *string `json:"WildcardCount" name:"WildcardCount"`
		} `json:"Certificate" name:"Certificate"`
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
	CertificateId     *string `json:"CertificateId,omitempty" name:"CertificateId"`
	AuthMethod        *string `json:"AuthMethod,omitempty" name:"AuthMethod"`
	CSR               *string `json:"CSR,omitempty" name:"CSR"`
	ContactId         *int    `json:"ContactId,omitempty" name:"ContactId"`
	CompanyId         *int    `json:"CompanyId,omitempty" name:"CompanyId"`
	CompanyName       *string `json:"CompanyName,omitempty" name:"CompanyName"`
	Department        *string `json:"Department,omitempty" name:"Department"`
	State             *string `json:"State,omitempty" name:"State"`
	City              *string `json:"City,omitempty" name:"City"`
	Address           *string `json:"Address,omitempty" name:"Address"`
	CompanyPhone      *string `json:"CompanyPhone,omitempty" name:"CompanyPhone"`
	PostalCode        *string `json:"PostalCode,omitempty" name:"PostalCode"`
	DcvEmail          *string `json:"DcvEmail,omitempty" name:"DcvEmail"`
	AdditionalDomains *string `json:"AdditionalDomains,omitempty" name:"AdditionalDomains"`
	Wildcards         *string `json:"Wildcards,omitempty" name:"Wildcards"`
	Contact           *string `json:"Contact,omitempty" name:"Contact"`
	IsSubmit          *string `json:"IsSubmit,omitempty" name:"IsSubmit"`
	BusinessLicence   *string `json:"BusinessLicence,omitempty" name:"BusinessLicence"`
	CsrSource         *string `json:"CsrSource,omitempty" name:"CsrSource"`
	Algorithm         *string `json:"Algorithm,omitempty" name:"Algorithm"`
	CertSignature     *string `json:"CertSignature,omitempty" name:"CertSignature"`
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
		RequestId   *string `json:"RequestId" name:"RequestId"`
		Certificate struct {
			CertificateId     *string `json:"CertificateId" name:"CertificateId"`
			MainDomain        *string `json:"MainDomain" name:"MainDomain"`
			CertificateBrand  *string `json:"CertificateBrand" name:"CertificateBrand"`
			CertificateLevel  *string `json:"CertificateLevel" name:"CertificateLevel"`
			CertificateName   *string `json:"CertificateName" name:"CertificateName"`
			CertificateCode   *string `json:"CertificateCode" name:"CertificateCode"`
			CertificateStatus *string `json:"CertificateStatus" name:"CertificateStatus"`
			YearLength        *string `json:"YearLength" name:"YearLength"`
			DomainCount       *string `json:"DomainCount" name:"DomainCount"`
			WildcardCount     *string `json:"WildcardCount" name:"WildcardCount"`
			AuthMethod        *string `json:"AuthMethod" name:"AuthMethod"`
			CSR               *string `json:"CSR" name:"CSR"`
			State             *string `json:"State" name:"State"`
			City              *string `json:"City" name:"City"`
			Address           *string `json:"Address" name:"Address"`
			CompanyPhone      *string `json:"CompanyPhone" name:"CompanyPhone"`
			PostalCode        *string `json:"PostalCode" name:"PostalCode"`
			AdditionalDomains *string `json:"AdditionalDomains" name:"AdditionalDomains"`
			Wildcards         *string `json:"Wildcards" name:"Wildcards"`
		} `json:"Certificate" name:"Certificate"`
		Contact struct {
			FirstName *string `json:"FirstName" name:"FirstName"`
			LastName  *string `json:"LastName" name:"LastName"`
			Phone     *string `json:"Phone" name:"Phone"`
			Email     *string `json:"Email" name:"Email"`
			Title     *string `json:"Title" name:"Title"`
		} `json:"Contact" name:"Contact"`
	} `json:"UpdateCertificateResponse"`
}

func (r *UpdateCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReIssueCertificateRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CsrSource     *string `json:"CsrSource,omitempty" name:"CsrSource"`
}

func (r *ReIssueCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReIssueCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReIssueCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReIssueCertificateResponse struct {
	*ksyunhttp.BaseResponse
	ReIssueCertificateResponse *string `json:"ReIssueCertificateResponse" name:"ReIssueCertificateResponse"`
}

func (r *ReIssueCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReIssueCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelTransactionRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CancelTransactionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelTransactionResponse struct {
	*ksyunhttp.BaseResponse
	CancleTransactionResponse struct {
		RequestId   *string `json:"RequestId" name:"RequestId"`
		Certificate struct {
			CertificateId     *string `json:"CertificateId" name:"CertificateId"`
			MainDomain        *string `json:"MainDomain" name:"MainDomain"`
			CertificateBrand  *string `json:"CertificateBrand" name:"CertificateBrand"`
			CertificateLevel  *string `json:"CertificateLevel" name:"CertificateLevel"`
			CertificateName   *string `json:"CertificateName" name:"CertificateName"`
			CertificateCode   *string `json:"CertificateCode" name:"CertificateCode"`
			CertificateStatus *string `json:"CertificateStatus" name:"CertificateStatus"`
			YearLength        *string `json:"YearLength" name:"YearLength"`
			DomainCount       *string `json:"DomainCount" name:"DomainCount"`
			WildcardCount     *string `json:"WildcardCount" name:"WildcardCount"`
		} `json:"Certificate" name:"Certificate"`
	} `json:"CancleTransactionResponse"`
}

func (r *CancelTransactionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCertificatesRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId []*string `json:"CertificateId,omitempty" name:"CertificateId"`
	ProjectId     []*int    `json:"ProjectId,omitempty" name:"ProjectId"`
	Filter        []*string `json:"Filter,omitempty" name:"Filter"`
	Page          *int      `json:"Page,omitempty" name:"Page"`
	PageSize      *int      `json:"PageSize,omitempty" name:"PageSize"`
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
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CertificateSet []struct {
		CertificateId     *string `json:"CertificateId" name:"CertificateId"`
		MainDomain        *string `json:"MainDomain" name:"MainDomain"`
		CertificateBrand  *string `json:"CertificateBrand" name:"CertificateBrand"`
		CertificateLevel  *string `json:"CertificateLevel" name:"CertificateLevel"`
		CertificateName   *string `json:"CertificateName" name:"CertificateName"`
		CertificateCode   *string `json:"CertificateCode" name:"CertificateCode"`
		CertificateStatus *int    `json:"CertificateStatus" name:"CertificateStatus"`
		YearLength        *int    `json:"YearLength" name:"YearLength"`
		DomainCount       *int    `json:"DomainCount" name:"DomainCount"`
		WildcardCount     *int    `json:"WildcardCount" name:"WildcardCount"`
		ExpireTime        *string `json:"ExpireTime" name:"ExpireTime"`
		AuthMethod        *string `json:"AuthMethod" name:"AuthMethod"`
		AuthContent       struct {
			CheckName  *string `json:"CheckName" name:"CheckName"`
			CheckValue *string `json:"CheckValue" name:"CheckValue"`
			RecordType *string `json:"RecordType" name:"RecordType"`
		} `json:"AuthContent" name:"AuthContent"`
	} `json:"CertificateSet"`
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
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Certificate struct {
		CertificateId     *string `json:"CertificateId" name:"CertificateId"`
		MainDomain        *string `json:"MainDomain" name:"MainDomain"`
		CertificateBrand  *string `json:"CertificateBrand" name:"CertificateBrand"`
		CertificateLevel  *string `json:"CertificateLevel" name:"CertificateLevel"`
		CertificateName   *string `json:"CertificateName" name:"CertificateName"`
		CertificateCode   *string `json:"CertificateCode" name:"CertificateCode"`
		CertificateStatus *int    `json:"CertificateStatus" name:"CertificateStatus"`
		YearLength        *int    `json:"YearLength" name:"YearLength"`
		DomainCount       *int    `json:"DomainCount" name:"DomainCount"`
		WildcardCount     *int    `json:"WildcardCount" name:"WildcardCount"`
		CompanyName       *string `json:"CompanyName" name:"CompanyName"`
		CSR               *string `json:"CSR" name:"CSR"`
		State             *string `json:"State" name:"State"`
		City              *string `json:"City" name:"City"`
		Address           *string `json:"Address" name:"Address"`
		CompanyPhone      *string `json:"CompanyPhone" name:"CompanyPhone"`
		PostalCode        *string `json:"PostalCode" name:"PostalCode"`
		Department        *string `json:"Department" name:"Department"`
		AdditionalDomains *string `json:"AdditionalDomains" name:"AdditionalDomains"`
		Wildcards         *string `json:"Wildcards" name:"Wildcards"`
		Contact           struct {
			FirstName *string `json:"FirstName" name:"FirstName"`
			LastName  *string `json:"LastName" name:"LastName"`
			Name      *string `json:"Name" name:"Name"`
			Phone     *string `json:"Phone" name:"Phone"`
			Email     *string `json:"Email" name:"Email"`
			Title     *string `json:"Title" name:"Title"`
		} `json:"Contact" name:"Contact"`
		AuthContent struct {
			CheckName  *string `json:"CheckName" name:"CheckName"`
			CheckValue *string `json:"CheckValue" name:"CheckValue"`
			RecordType *string `json:"RecordType" name:"RecordType"`
		} `json:"AuthContent" name:"AuthContent"`
	} `json:"Certificate"`
}

func (r *GetCertificateDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
