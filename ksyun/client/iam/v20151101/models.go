package v20151101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateUserRequest struct {
	*ksyunhttp.BaseRequest
	UserName               *string `json:"UserName,omitempty" name:"UserName"`
	RealName               *string `json:"RealName,omitempty" name:"RealName"`
	Phone                  *string `json:"Phone,omitempty" name:"Phone"`
	Email                  *string `json:"Email,omitempty" name:"Email"`
	Remark                 *string `json:"Remark,omitempty" name:"Remark"`
	Password               *string `json:"Password,omitempty" name:"Password"`
	PasswordResetRequired  *int    `json:"PasswordResetRequired,omitempty" name:"PasswordResetRequired"`
	OpenLoginProtection    *int    `json:"OpenLoginProtection,omitempty" name:"OpenLoginProtection"`
	OpenSecurityProtection *int    `json:"OpenSecurityProtection,omitempty" name:"OpenSecurityProtection"`
	ViewAllProject         *int    `json:"ViewAllProject,omitempty" name:"ViewAllProject"`
	AddProjectId           *int    `json:"AddProjectId,omitempty" name:"AddProjectId"`
}

func (r *CreateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*ksyunhttp.BaseResponse
	CreateUserResult struct {
		User struct {
			UserId                *string `json:"UserId" name:"UserId"`
			Path                  *string `json:"Path" name:"Path"`
			UserName              *string `json:"UserName" name:"UserName"`
			RealName              *string `json:"RealName" name:"RealName"`
			CreateDate            *string `json:"CreateDate" name:"CreateDate"`
			Phone                 *string `json:"Phone" name:"Phone"`
			CountryMobileCode     *string `json:"CountryMobileCode" name:"CountryMobileCode"`
			IsInternational       *int    `json:"isInternational" name:"isInternational"`
			Email                 *string `json:"Email" name:"Email"`
			PhoneVerified         *string `json:"PhoneVerified" name:"PhoneVerified"`
			EmailVerified         *string `json:"EmailVerified" name:"EmailVerified"`
			Remark                *string `json:"Remark" name:"Remark"`
			Krn                   *string `json:"Krn" name:"Krn"`
			PasswordResetRequired *bool   `json:"PasswordResetRequired" name:"PasswordResetRequired"`
			EnableMFA             *int    `json:"EnableMFA" name:"EnableMFA"`
			NeedBindMfa           *string `json:"NeedBindMfa" name:"NeedBindMfa"`
			UpdateDate            *string `json:"UpdateDate" name:"UpdateDate"`
		} `json:"User" name:"User"`
	} `json:"CreateUserResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersRequest struct {
	*ksyunhttp.BaseRequest
	Marker    *string `json:"Marker,omitempty" name:"Marker"`
	MaxItems  *int    `json:"MaxItems,omitempty" name:"MaxItems"`
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`
}

func (r *ListUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersResponse struct {
	*ksyunhttp.BaseResponse
	ListUserResult struct {
		Users struct {
			Member []struct {
				UserId                *string `json:"UserId" name:"UserId"`
				Path                  *string `json:"Path" name:"Path"`
				UserName              *string `json:"UserName" name:"UserName"`
				RealName              *string `json:"RealName" name:"RealName"`
				CreateDate            *string `json:"CreateDate" name:"CreateDate"`
				Phone                 *string `json:"Phone" name:"Phone"`
				CountryMobileCode     *string `json:"CountryMobileCode" name:"CountryMobileCode"`
				IsInternational       *int    `json:"IsInternational" name:"IsInternational"`
				Email                 *string `json:"Email" name:"Email"`
				PhoneVerified         *string `json:"PhoneVerified" name:"PhoneVerified"`
				EmailVerified         *string `json:"EmailVerified" name:"EmailVerified"`
				Remark                *string `json:"Remark" name:"Remark"`
				Krn                   *string `json:"Krn" name:"Krn"`
				PasswordResetRequired *bool   `json:"PasswordResetRequired" name:"PasswordResetRequired"`
				EnableMFA             *int    `json:"EnableMFA" name:"EnableMFA"`
				NeedBindMfa           *string `json:"NeedBindMfa" name:"NeedBindMfa"`
				UpdateDate            *string `json:"UpdateDate" name:"UpdateDate"`
				Id                    *int    `json:"Id" name:"Id"`
			} `json:"Member"`
		} `json:"Users" name:"Users"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListUserResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserRequest struct {
	*ksyunhttp.BaseRequest
	UserName        *string `json:"UserName,omitempty" name:"UserName"`
	NewUserName     *string `json:"NewUserName,omitempty" name:"NewUserName"`
	NewRealName     *string `json:"NewRealName,omitempty" name:"NewRealName"`
	NewEmail        *string `json:"NewEmail,omitempty" name:"NewEmail"`
	NewPhone        *string `json:"NewPhone,omitempty" name:"NewPhone"`
	IsInternational *int    `json:"IsInternational,omitempty" name:"IsInternational"`
	NewRemark       *string `json:"NewRemark,omitempty" name:"NewRemark"`
}

func (r *UpdateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserResponse struct {
	*ksyunhttp.BaseResponse
	UpdateUserResult struct {
		User struct {
			UserId                *string `json:"UserId" name:"UserId"`
			Path                  *string `json:"Path" name:"Path"`
			UserName              *string `json:"UserName" name:"UserName"`
			RealName              *string `json:"RealName" name:"RealName"`
			CreateDate            *string `json:"CreateDate" name:"CreateDate"`
			Phone                 *string `json:"Phone" name:"Phone"`
			CountryMobileCode     *string `json:"CountryMobileCode" name:"CountryMobileCode"`
			IsInternational       *int    `json:"isInternational" name:"isInternational"`
			Email                 *string `json:"Email" name:"Email"`
			PhoneVerified         *string `json:"PhoneVerified" name:"PhoneVerified"`
			EmailVerified         *string `json:"EmailVerified" name:"EmailVerified"`
			Remark                *string `json:"Remark" name:"Remark"`
			Krn                   *string `json:"Krn" name:"Krn"`
			PasswordResetRequired *string `json:"PasswordResetRequired" name:"PasswordResetRequired"`
			EnableMFA             *int    `json:"EnableMFA" name:"EnableMFA"`
			NeedBindMfa           *string `json:"NeedBindMfa" name:"NeedBindMfa"`
			UpdateDate            *string `json:"UpdateDate" name:"UpdateDate"`
		} `json:"User" name:"User"`
	} `json:"UpdateUserResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *GetUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetUserResponse struct {
	*ksyunhttp.BaseResponse
	GetUserResult struct {
		UserId                *string `json:"UserId" name:"UserId"`
		Path                  *string `json:"Path" name:"Path"`
		UserName              *string `json:"UserName" name:"UserName"`
		CreateDate            *string `json:"CreateDate" name:"CreateDate"`
		Phone                 *int    `json:"Phone" name:"Phone"`
		CountryMobileCode     *int    `json:"CountryMobileCode" name:"CountryMobileCode"`
		IsInternational       *int    `json:"IsInternational" name:"IsInternational"`
		Email                 *string `json:"Email" name:"Email"`
		PhoneVerified         *string `json:"PhoneVerified" name:"PhoneVerified"`
		EmailVerified         *string `json:"EmailVerified" name:"EmailVerified"`
		Remark                *string `json:"Remark" name:"Remark"`
		PasswordResetRequired *string `json:"PasswordResetRequired" name:"PasswordResetRequired"`
		EnableMFA             *int    `json:"EnableMFA" name:"EnableMFA"`
		NeedBindMfa           *int    `json:"NeedBindMfa" name:"NeedBindMfa"`
		UpdateDate            *string `json:"UpdateDate" name:"UpdateDate"`
		ViewAllProject        *string `json:"ViewAllProject" name:"ViewAllProject"`
		Id                    *int    `json:"Id" name:"Id"`
	} `json:"GetUserResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DeleteUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserResponse struct {
	*ksyunhttp.BaseResponse
	Result    *int    `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	UserName  *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserPoliciesRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
	Marker   *string `json:"Marker,omitempty" name:"Marker"`
	MaxItems *string `json:"MaxItems,omitempty" name:"MaxItems"`
}

func (r *ListAttachedUserPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedUserPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAttachedUserPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserPoliciesResponse struct {
	*ksyunhttp.BaseResponse
	ListAttachedUserPoliciesResult struct {
		AttachedPolicies struct {
			Member []struct {
				PolicyKrn     *string `json:"PolicyKrn" name:"PolicyKrn"`
				PolicyName    *string `json:"PolicyName" name:"PolicyName"`
				CreateTime    *string `json:"CreateTime" name:"CreateTime"`
				Description   *string `json:"Description" name:"Description"`
				DescriptionEn *string `json:"DescriptionEn" name:"DescriptionEn"`
				Type          *int    `json:"Type" name:"Type"`
			} `json:"Member"`
		} `json:"AttachedPolicies" name:"AttachedPolicies"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListAttachedUserPoliciesResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAttachedUserPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedUserPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPolicyVersionsRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *ListPolicyVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPolicyVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListPolicyVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListPolicyVersionsResponse struct {
	*ksyunhttp.BaseResponse
	ListPolicyVersionsResult struct {
		Versions struct {
			Member []struct {
				IsDefaultVersion *string   `json:"IsDefaultVersion" name:"IsDefaultVersion"`
				VersionId        *string   `json:"VersionId" name:"VersionId"`
				Document         *string   `json:"Document" name:"Document"`
				PolicyStruct     []*string `json:"PolicyStruct" name:"PolicyStruct"`
				CreateDate       *string   `json:"CreateDate" name:"CreateDate"`
			} `json:"Member"`
		} `json:"Versions" name:"Versions"`
	} `json:"ListPolicyVersionsResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListPolicyVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPolicyVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPolicyVersionRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *SetDefaultPolicyVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultPolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetDefaultPolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPolicyVersionResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetDefaultPolicyVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultPolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	UserName  *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *int    `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyVersionRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DeletePolicyVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyVersionResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeletePolicyVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPolicyVersionRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *GetPolicyVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPolicyVersionResponse struct {
	*ksyunhttp.BaseResponse
	GetPolicyVersionResult struct {
		PolicyVersion struct {
			IsDefaultVersion *string   `json:"IsDefaultVersion" name:"IsDefaultVersion"`
			VersionId        *string   `json:"VersionId" name:"VersionId"`
			Document         *string   `json:"Document" name:"Document"`
			PolicyStruct     []*string `json:"PolicyStruct" name:"PolicyStruct"`
			CreateDate       *string   `json:"CreateDate" name:"CreateDate"`
		} `json:"PolicyVersion" name:"PolicyVersion"`
	} `json:"GetPolicyVersionResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetPolicyVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyVersionRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn      *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
	SetAsDefault   *string `json:"SetAsDefault,omitempty" name:"SetAsDefault"`
	PolicyStruct   *string `json:"PolicyStruct,omitempty" name:"PolicyStruct"`
}

func (r *CreatePolicyVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyVersionResponse struct {
	*ksyunhttp.BaseResponse
	CreatePolicyVersionResult struct {
		PolicyVersion struct {
			IsDefaultVersion *string `json:"IsDefaultVersion" name:"IsDefaultVersion"`
			VersionId        *string `json:"VersionId" name:"VersionId"`
			CreateDate       *string `json:"CreateDate" name:"CreateDate"`
		} `json:"PolicyVersion" name:"PolicyVersion"`
	} `json:"CreatePolicyVersionResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreatePolicyVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesRequest struct {
	*ksyunhttp.BaseRequest
	Marker       *string `json:"Marker,omitempty" name:"Marker"`
	MaxItems     *string `json:"MaxItems,omitempty" name:"MaxItems"`
	OnlyAttached *bool   `json:"OnlyAttached,omitempty" name:"OnlyAttached"`
	Scope        *string `json:"Scope,omitempty" name:"Scope"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesResponse struct {
	*ksyunhttp.BaseResponse
	ListPoliciesResult struct {
		Policies struct {
			Member []struct {
				CreateDate       *string `json:"CreateDate" name:"CreateDate"`
				DefaultVersionId *string `json:"DefaultVersionId" name:"DefaultVersionId"`
				Description      *string `json:"Description" name:"Description"`
				Krn              *string `json:"Krn" name:"Krn"`
				Path             *string `json:"Path" name:"Path"`
				PolicyId         *string `json:"PolicyId" name:"PolicyId"`
				PolicyName       *string `json:"PolicyName" name:"PolicyName"`
				ServiceId        *int    `json:"ServiceId" name:"ServiceId"`
				ServiceName      *string `json:"ServiceName" name:"ServiceName"`
				ServiceViewName  *string `json:"ServiceViewName" name:"ServiceViewName"`
				PolicyType       *int    `json:"PolicyType" name:"PolicyType"`
				CreateMode       *int    `json:"CreateMode" name:"CreateMode"`
				UpdateDate       *string `json:"UpdateDate" name:"UpdateDate"`
				AttachmentCount  *int    `json:"AttachmentCount" name:"AttachmentCount"`
			} `json:"Member"`
		} `json:"Policies" name:"Policies"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListPoliciesResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *GetPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPolicyResponse struct {
	*ksyunhttp.BaseResponse
	GetPolicyResult struct {
		Policy struct {
			CreateDate       *string `json:"CreateDate" name:"CreateDate"`
			DefaultVersionId *string `json:"DefaultVersionId" name:"DefaultVersionId"`
			Description      *string `json:"Description" name:"Description"`
			Krn              *string `json:"Krn" name:"Krn"`
			Path             *string `json:"Path" name:"Path"`
			PolicyId         *string `json:"PolicyId" name:"PolicyId"`
			PolicyName       *string `json:"PolicyName" name:"PolicyName"`
			PolicyType       *int    `json:"PolicyType" name:"PolicyType"`
			CreateMode       *int    `json:"CreateMode" name:"CreateMode"`
			UpdateDate       *string `json:"UpdateDate" name:"UpdateDate"`
		} `json:"Policy" name:"Policy"`
	} `json:"GetPolicyResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyName     *string `json:"PolicyName,omitempty" name:"PolicyName"`
	Description    *string `json:"Description,omitempty" name:"Description"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
	PolicyStruct   *string `json:"PolicyStruct,omitempty" name:"PolicyStruct"`
	CreateMode     *string `json:"CreateMode,omitempty" name:"CreateMode"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyResponse struct {
	*ksyunhttp.BaseResponse
	CreatePolicyResult struct {
		Policy struct {
			AttachmentCount  *string `json:"AttachmentCount" name:"AttachmentCount"`
			CreateDate       *string `json:"CreateDate" name:"CreateDate"`
			DefaultVersionId *string `json:"DefaultVersionId" name:"DefaultVersionId"`
			Description      *string `json:"Description" name:"Description"`
			Krn              *string `json:"Krn" name:"Krn"`
			Path             *string `json:"Path" name:"Path"`
			PolicyId         *string `json:"PolicyId" name:"PolicyId"`
			PolicyName       *string `json:"PolicyName" name:"PolicyName"`
			UpdateDate       *string `json:"UpdateDate" name:"UpdateDate"`
		} `json:"Policy" name:"Policy"`
	} `json:"CreatePolicyResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangePasswordRequest struct {
	*ksyunhttp.BaseRequest
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *ChangePasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ChangePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ChangePasswordResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ChangePasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLoginProfileRequest struct {
	*ksyunhttp.BaseRequest
	UserName               *string `json:"UserName,omitempty" name:"UserName"`
	Password               *string `json:"Password,omitempty" name:"Password"`
	PasswordResetRequired  *bool   `json:"PasswordResetRequired,omitempty" name:"PasswordResetRequired"`
	OpenLoginProtection    *bool   `json:"OpenLoginProtection,omitempty" name:"OpenLoginProtection"`
	OpenSecurityProtection *bool   `json:"OpenSecurityProtection,omitempty" name:"OpenSecurityProtection"`
	ViewAllProject         *bool   `json:"ViewAllProject,omitempty" name:"ViewAllProject"`
}

func (r *UpdateLoginProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLoginProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateLoginProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLoginProfileResponse struct {
	*ksyunhttp.BaseResponse
	CreateLoginProfileResult struct {
		LoginProfile struct {
			UserName               *string `json:"UserName" name:"UserName"`
			OpenLoginProtection    *string `json:"OpenLoginProtection" name:"OpenLoginProtection"`
			OpenSecurityProtection *string `json:"OpenSecurityProtection" name:"OpenSecurityProtection"`
			PasswordResetRequired  *string `json:"PasswordResetRequired" name:"PasswordResetRequired"`
			UpdateDate             *string `json:"UpdateDate" name:"UpdateDate"`
			Password               *string `json:"Password" name:"Password"`
		} `json:"LoginProfile" name:"LoginProfile"`
	} `json:"CreateLoginProfileResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateLoginProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLoginProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginProfileRequest struct {
	*ksyunhttp.BaseRequest
	UserName         *string `json:"UserName,omitempty" name:"UserName"`
	NotCheckPassword *int    `json:"NotCheckPassword,omitempty" name:"NotCheckPassword"`
}

func (r *GetLoginProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetLoginProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginProfileResponse struct {
	*ksyunhttp.BaseResponse
	CreateLoginProfileResult struct {
		LoginProfile struct {
			UserName               *string `json:"UserName" name:"UserName"`
			PasswordResetRequired  *bool   `json:"PasswordResetRequired" name:"PasswordResetRequired"`
			OpenLoginProtection    *int    `json:"OpenLoginProtection" name:"OpenLoginProtection"`
			OpenSecurityProtection *int    `json:"OpenSecurityProtection" name:"OpenSecurityProtection"`
			EnableMfa              *int    `json:"EnableMfa" name:"EnableMfa"`
			ViewAllProject         *int    `json:"ViewAllProject" name:"ViewAllProject"`
			LastLoginDate          *string `json:"LastLoginDate" name:"LastLoginDate"`
			LastLoginIp            *string `json:"LastLoginIp" name:"LastLoginIp"`
			CreateDate             *string `json:"CreateDate" name:"CreateDate"`
			ConsoleLogin           *bool   `json:"ConsoleLogin" name:"ConsoleLogin"`
		} `json:"LoginProfile" name:"LoginProfile"`
	} `json:"CreateLoginProfileResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetLoginProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessKeyRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *CreateAccessKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessKeyResponse struct {
	*ksyunhttp.BaseResponse
	CreateAccessKeyResult struct {
		AccessKey struct {
			UserName        *string `json:"UserName" name:"UserName"`
			AccessKeyId     *string `json:"AccessKeyId" name:"AccessKeyId"`
			SecretAccessKey *string `json:"SecretAccessKey" name:"SecretAccessKey"`
			Status          *string `json:"Status" name:"Status"`
			CreateDate      *string `json:"CreateDate" name:"CreateDate"`
		} `json:"AccessKey" name:"AccessKey"`
	} `json:"CreateAccessKeyResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateAccessKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAccessKeysRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *ListAccessKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAccessKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAccessKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAccessKeysResponse struct {
	*ksyunhttp.BaseResponse
	ListAccessKeyResult struct {
		AccessKeyMetadata struct {
			Member []struct {
				UserName       *string `json:"UserName" name:"UserName"`
				AccessKeyId    *string `json:"AccessKeyId" name:"AccessKeyId"`
				Status         *string `json:"Status" name:"Status"`
				CreateDate     *string `json:"CreateDate" name:"CreateDate"`
				AkLastUsedTime *string `json:"AkLastUsedTime" name:"AkLastUsedTime"`
			} `json:"Member"`
		} `json:"AccessKeyMetadata" name:"AccessKeyMetadata"`
	} `json:"ListAccessKeyResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAccessKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAccessKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccessKeyRequest struct {
	*ksyunhttp.BaseRequest
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	UserName    *string `json:"UserName,omitempty" name:"UserName"`
	Status      *string `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateAccessKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccessKeyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateAccessKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessKeyRequest struct {
	*ksyunhttp.BaseRequest
	UserName    *string `json:"UserName,omitempty" name:"UserName"`
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DeleteAccessKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessKeyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteAccessKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVirtualMFADevicesRequest struct {
	*ksyunhttp.BaseRequest
	AssignmentStatus *string `json:"AssignmentStatus,omitempty" name:"AssignmentStatus"`
	Marker           *string `json:"Marker,omitempty" name:"Marker"`
	MaxItems         *int    `json:"MaxItems,omitempty" name:"MaxItems"`
}

func (r *ListVirtualMFADevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVirtualMFADevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListVirtualMFADevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListVirtualMFADevicesResponse struct {
	*ksyunhttp.BaseResponse
	ListVirtualMFADevicesResult struct {
		VirtualMFADevices struct {
			Member []struct {
				SerialNumber *string `json:"SerialNumber" name:"SerialNumber"`
				EnableDate   *string `json:"EnableDate" name:"EnableDate"`
				User         struct {
					Uuid        *string `json:"Uuid" name:"Uuid"`
					Name        *string `json:"Name" name:"Name"`
					RealName    *string `json:"RealName" name:"RealName"`
					Path        *string `json:"Path" name:"Path"`
					Krn         *string `json:"Krn" name:"Krn"`
					CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
					PwdLastUsed *string `json:"PwdLastUsed" name:"PwdLastUsed"`
				} `json:"User" name:"User"`
			} `json:"Member"`
		} `json:"VirtualMFADevices" name:"VirtualMFADevices"`
		Marker      *string `json:"Marker" name:"Marker"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
	} `json:"ListVirtualMFADevicesResult"`
}

func (r *ListVirtualMFADevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVirtualMFADevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableMFADeviceRequest struct {
	*ksyunhttp.BaseRequest
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" name:"AuthenticationCode1"`
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" name:"AuthenticationCode2"`
	SerialNumber        *string `json:"SerialNumber,omitempty" name:"SerialNumber"`
	UserName            *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *EnableMFADeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableMFADeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EnableMFADeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableMFADeviceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EnableMFADeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableMFADeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeactivateMFADeviceRequest struct {
	*ksyunhttp.BaseRequest
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`
	UserName     *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DeactivateMFADeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeactivateMFADeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeactivateMFADeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeactivateMFADeviceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeactivateMFADeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeactivateMFADeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVirtualMFADeviceRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *GetVirtualMFADeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVirtualMFADeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetVirtualMFADeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetVirtualMFADeviceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	VirtualMFADevice struct {
		SerialNumber *string `json:"SerialNumber" name:"SerialNumber"`
		EnableDate   *string `json:"EnableDate" name:"EnableDate"`
	} `json:"VirtualMFADevice"`
}

func (r *GetVirtualMFADeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVirtualMFADeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*ksyunhttp.BaseRequest
	RoleName      *string `json:"RoleName,omitempty" name:"RoleName"`
	TrustAccounts *string `json:"TrustAccounts,omitempty" name:"TrustAccounts"`
	Description   *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*ksyunhttp.BaseResponse
	CreateRoleResult struct {
		Role struct {
			Path          *string `json:"Path" name:"Path"`
			Krn           *string `json:"Krn" name:"Krn"`
			RoleName      *string `json:"RoleName" name:"RoleName"`
			Description   *string `json:"Description" name:"Description"`
			TrustAccounts *string `json:"TrustAccounts" name:"TrustAccounts"`
			CreateDate    *string `json:"CreateDate" name:"CreateDate"`
			RoleId        *string `json:"RoleId" name:"RoleId"`
		} `json:"Role" name:"Role"`
	} `json:"CreateRoleResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleRequest struct {
	*ksyunhttp.BaseRequest
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRoleRequest struct {
	*ksyunhttp.BaseRequest
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *GetRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRoleResponse struct {
	*ksyunhttp.BaseResponse
	GetRoleResult struct {
		Role struct {
			Id            *int      `json:"Id" name:"Id"`
			Path          *string   `json:"Path" name:"Path"`
			Krn           *string   `json:"Krn" name:"Krn"`
			RoleName      *string   `json:"RoleName" name:"RoleName"`
			Description   *string   `json:"Description" name:"Description"`
			TrustAccounts *string   `json:"TrustAccounts" name:"TrustAccounts"`
			TrustServices *string   `json:"TrustServices" name:"TrustServices"`
			TrustProvider *string   `json:"TrustProvider" name:"TrustProvider"`
			ProviderList  []*string `json:"ProviderList" name:"ProviderList"`
			TrustType     *int      `json:"TrustType" name:"TrustType"`
			CreateDate    *string   `json:"CreateDate" name:"CreateDate"`
			RoleId        *string   `json:"RoleId" name:"RoleId"`
		} `json:"Role" name:"Role"`
	} `json:"GetRoleResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRolesRequest struct {
	*ksyunhttp.BaseRequest
	Marker   *string `json:"Marker,omitempty" name:"Marker"`
	MaxItems *int    `json:"MaxItems,omitempty" name:"MaxItems"`
}

func (r *ListRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListRolesResponse struct {
	*ksyunhttp.BaseResponse
	ListRolesResult struct {
		Roles struct {
			Member []struct {
				Path            *string   `json:"Path" name:"Path"`
				Krn             *string   `json:"Krn" name:"Krn"`
				RoleName        *string   `json:"RoleName" name:"RoleName"`
				Description     *string   `json:"Description" name:"Description"`
				TrustType       *int      `json:"TrustType" name:"TrustType"`
				TrustAccounts   *string   `json:"TrustAccounts" name:"TrustAccounts"`
				TrustServices   []*string `json:"TrustServices" name:"TrustServices"`
				TrustProvider   *string   `json:"TrustProvider" name:"TrustProvider"`
				CreateDate      *string   `json:"CreateDate" name:"CreateDate"`
				RoleId          *string   `json:"RoleId" name:"RoleId"`
				ServiceRoleType *int      `json:"ServiceRoleType" name:"ServiceRoleType"`
			} `json:"Member"`
		} `json:"Roles" name:"Roles"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListRolesResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyRequest struct {
	*ksyunhttp.BaseRequest
	RoleName  *string `json:"RoleName,omitempty" name:"RoleName"`
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *AttachRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *int    `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachRolePolicyRequest struct {
	*ksyunhttp.BaseRequest
	RoleName  *string `json:"RoleName,omitempty" name:"RoleName"`
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *DetachRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachRolePolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DetachRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesRequest struct {
	*ksyunhttp.BaseRequest
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	Marker   *string `json:"Marker,omitempty" name:"Marker"`
	MaxItems *int    `json:"MaxItems,omitempty" name:"MaxItems"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAttachedRolePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesResponse struct {
	*ksyunhttp.BaseResponse
	ListAttachedRolePoliciesResult struct {
		AttachedPolicies struct {
			Member []struct {
				PolicyKrn     *string `json:"PolicyKrn" name:"PolicyKrn"`
				PolicyName    *string `json:"PolicyName" name:"PolicyName"`
				CreateTime    *string `json:"CreateTime" name:"CreateTime"`
				Description   *string `json:"Description" name:"Description"`
				DescriptionEn *string `json:"DescriptionEn" name:"DescriptionEn"`
				Type          *int    `json:"Type" name:"Type"`
			} `json:"Member"`
		} `json:"AttachedPolicies" name:"AttachedPolicies"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListAttachedRolePoliciesResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleTrustAccountsRequest struct {
	*ksyunhttp.BaseRequest
	RoleName         *string `json:"RoleName,omitempty" name:"RoleName"`
	NewTrustAccounts *string `json:"NewTrustAccounts,omitempty" name:"NewTrustAccounts"`
}

func (r *UpdateRoleTrustAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleTrustAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateRoleTrustAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleTrustAccountsResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateRoleTrustAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleTrustAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*ksyunhttp.BaseRequest
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*ksyunhttp.BaseResponse
	Result    *int    `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectInfoRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId   *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
}

func (r *UpdateProjectInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateProjectInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectInfoResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateProjectInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountAllProjectListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetAccountAllProjectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountAllProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetAccountAllProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountAllProjectListResponse struct {
	*ksyunhttp.BaseResponse
	ListProjectResult struct {
		Total       *int `json:"Total" name:"Total"`
		ProjectList []struct {
			ProjectId   *int    `json:"ProjectId" name:"ProjectId"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			ProjectDesc *string `json:"ProjectDesc" name:"ProjectDesc"`
			Status      *int    `json:"Status" name:"Status"`
			Krn         *string `json:"Krn" name:"Krn"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
		} `json:"ProjectList" name:"ProjectList"`
	} `json:"ListProjectResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetAccountAllProjectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountAllProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectInstanceListRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId   *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	ProductLine *string `json:"ProductLine,omitempty" name:"ProductLine"`
	Ps          *int    `json:"Ps,omitempty" name:"Ps"`
	Pn          *int    `json:"Pn,omitempty" name:"Pn"`
}

func (r *GetProjectInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetProjectInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectInstanceListResponse struct {
	*ksyunhttp.BaseResponse
	ListInstanceResult struct {
		Items []struct {
			ProjectId    *string `json:"ProjectId" name:"ProjectId"`
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			ProductGroup *string `json:"ProductGroup" name:"ProductGroup"`
			ProductType  *string `json:"ProductType" name:"ProductType"`
			Region       *string `json:"Region" name:"Region"`
			AccountId    *int    `json:"AccountId" name:"AccountId"`
			Status       *string `json:"Status" name:"Status"`
		} `json:"Items" name:"Items"`
		CurrentPage *int `json:"CurrentPage" name:"CurrentPage"`
		PageSize    *int `json:"PageSize" name:"PageSize"`
		Total       *int `json:"Total" name:"Total"`
	} `json:"ListInstanceResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetProjectInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceProjectIdRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId  *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Action     *string `json:"Action,omitempty" name:"Action"`
	Version    *string `json:"Version,omitempty" name:"Version"`
}

func (r *UpdateInstanceProjectIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceProjectIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateInstanceProjectIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceProjectIdResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateInstanceProjectIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceProjectIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	MaxItems  *int    `json:"MaxItems,omitempty" name:"MaxItems"`
	Marker    *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *ListEntitiesForPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListEntitiesForPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyResponse struct {
	*ksyunhttp.BaseResponse
	ListEntitiesForPolicyResult struct {
		PolicyUsers struct {
			Member []struct {
				UserId       *string `json:"UserId" name:"UserId"`
				UserName     *string `json:"UserName" name:"UserName"`
				UserRealName *string `json:"UserRealName" name:"UserRealName"`
				CreateDate   *string `json:"CreateDate" name:"CreateDate"`
			} `json:"Member"`
		} `json:"PolicyUsers" name:"PolicyUsers"`
		PolicyRoles struct {
			Member []struct {
				RoleId          *string `json:"RoleId" name:"RoleId"`
				RoleName        *string `json:"RoleName" name:"RoleName"`
				RoleDescription *string `json:"RoleDescription" name:"RoleDescription"`
				CreateDate      *string `json:"CreateDate" name:"CreateDate"`
			} `json:"Member"`
		} `json:"PolicyRoles" name:"PolicyRoles"`
		PolicyGroups struct {
			Member []struct {
				GroupId          *string `json:"GroupId" name:"GroupId"`
				GroupName        *string `json:"GroupName" name:"GroupName"`
				CreateDate       *string `json:"CreateDate" name:"CreateDate"`
				GroupDescription *string `json:"GroupDescription" name:"GroupDescription"`
			} `json:"Member"`
		} `json:"PolicyGroups" name:"PolicyGroups"`
		PolicyAccounts struct {
			Member []struct {
				AccountId       *int    `json:"AccountId" name:"AccountId"`
				AccountName     *string `json:"AccountName" name:"AccountName"`
				AccountUsername *string `json:"AccountUsername" name:"AccountUsername"`
				CreateDate      *string `json:"CreateDate" name:"CreateDate"`
			} `json:"Member"`
		} `json:"PolicyAccounts" name:"PolicyAccounts"`
		PolicyResourceDirs struct {
			Member []struct {
				ResourceDirId          *string `json:"ResourceDirId" name:"ResourceDirId"`
				ResourceDirName        *string `json:"ResourceDirName" name:"ResourceDirName"`
				ResourceDirDescription *string `json:"ResourceDirDescription" name:"ResourceDirDescription"`
				CreateDate             *string `json:"CreateDate" name:"CreateDate"`
			} `json:"Member"`
		} `json:"PolicyResourceDirs" name:"PolicyResourceDirs"`
		IsTruncated *bool `json:"IsTruncated" name:"IsTruncated"`
	} `json:"ListEntitiesForPolicyResult"`
	RequestId          *string `json:"RequestId" name:"RequestId"`
	ListPoliciesResult struct {
		Marker *string `json:"Marker" name:"Marker"`
	} `json:"ListPoliciesResult"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectMemberRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId *int `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ListProjectMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListProjectMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectMemberResponse struct {
	*ksyunhttp.BaseResponse
	ListProjectMember []struct {
		MemberId     *int    `json:"MemberId" name:"MemberId"`
		IdentityType *int    `json:"IdentityType" name:"IdentityType"`
		IdentityName *string `json:"IdentityName" name:"IdentityName"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		Krn          *string `json:"Krn" name:"Krn"`
	} `json:"ListProjectMember"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListProjectMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectMemberRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	MemberIds *string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *DeleteProjectMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteProjectMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectMemberResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteProjectMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddProjectMemberRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId    *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	IdentityId   *string `json:"IdentityId,omitempty" name:"IdentityId"`
	IdentityType *int    `json:"IdentityType,omitempty" name:"IdentityType"`
}

func (r *AddProjectMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddProjectMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddProjectMemberResponse struct {
	*ksyunhttp.BaseResponse
	Result    *int    `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddProjectMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProjectMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleRequest struct {
	*ksyunhttp.BaseRequest
	RoleName       *string `json:"RoleName,omitempty" name:"RoleName"`
	NewDescription *string `json:"NewDescription,omitempty" name:"NewDescription"`
}

func (r *UpdateRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*ksyunhttp.BaseRequest
	PolicyKrn      *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
	NewDescription *string `json:"NewDescription,omitempty" name:"NewDescription"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*ksyunhttp.BaseResponse
	UpdatePolicyResult struct {
		Policy *bool `json:"Policy" name:"Policy"`
	} `json:"UpdatePolicyResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*ksyunhttp.BaseRequest
	GroupName   *string `json:"GroupName,omitempty" name:"GroupName"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*ksyunhttp.BaseResponse
	CreateGroupResult struct {
		Group struct {
			GroupId     *string `json:"GroupId" name:"GroupId"`
			Path        *string `json:"Path" name:"Path"`
			GroupName   *string `json:"GroupName" name:"GroupName"`
			Description *string `json:"Description" name:"Description"`
			CreateDate  *string `json:"CreateDate" name:"CreateDate"`
			Krn         *string `json:"Krn" name:"Krn"`
		} `json:"Group" name:"Group"`
	} `json:"CreateGroupResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *DetachGroupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DetachGroupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachGroupPolicyRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	PolicyKrn *string `json:"PolicyKrn,omitempty" name:"PolicyKrn"`
}

func (r *AttachGroupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachGroupPolicyResponse struct {
	*ksyunhttp.BaseResponse
	Result    *int    `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AttachGroupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupPoliciesRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	MaxItems  *string `json:"MaxItems,omitempty" name:"MaxItems"`
	Marker    *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *ListGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupPoliciesResponse struct {
	*ksyunhttp.BaseResponse
	ListGroupPoliciesResult struct {
		AttachedPolicies struct {
			Member []struct {
				PolicyKrn     *string `json:"PolicyKrn" name:"PolicyKrn"`
				PolicyName    *string `json:"PolicyName" name:"PolicyName"`
				CreateTime    *string `json:"CreateTime" name:"CreateTime"`
				Description   *string `json:"Description" name:"Description"`
				DescriptionEn *string `json:"DescriptionEn" name:"DescriptionEn"`
				Type          *int    `json:"Type" name:"Type"`
			} `json:"Member"`
		} `json:"AttachedPolicies" name:"AttachedPolicies"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListGroupPoliciesResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	UserName  *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserToGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddUserToGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserToGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	MaxItems  *string `json:"MaxItems,omitempty" name:"MaxItems"`
	Marker    *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *GetGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupResponse struct {
	*ksyunhttp.BaseResponse
	GetGroupResult struct {
		Group struct {
			Group struct {
				GroupId     *string `json:"GroupId" name:"GroupId"`
				Path        *string `json:"Path" name:"Path"`
				GroupName   *string `json:"GroupName" name:"GroupName"`
				Description *string `json:"Description" name:"Description"`
				CreateDate  *string `json:"CreateDate" name:"CreateDate"`
				Krn         *string `json:"Krn" name:"Krn"`
			} `json:"Group"`
		} `json:"Group" name:"Group"`
	} `json:"GetGroupResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupsForUserRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
	MaxItems *string `json:"MaxItems,omitempty" name:"MaxItems"`
	Marker   *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *ListGroupsForUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListGroupsForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupsForUserResponse struct {
	*ksyunhttp.BaseResponse
	ListGroupsForUserResult struct {
		Groups struct {
			Member []struct {
				Id          *int    `json:"Id" name:"Id"`
				GroupId     *string `json:"GroupId" name:"GroupId"`
				Path        *string `json:"Path" name:"Path"`
				GroupName   *string `json:"GroupName" name:"GroupName"`
				Description *string `json:"Description" name:"Description"`
				CreateDate  *string `json:"CreateDate" name:"CreateDate"`
				Krn         *string `json:"Krn" name:"Krn"`
				PolicyCount *int    `json:"PolicyCount" name:"PolicyCount"`
			} `json:"Member"`
		} `json:"Groups" name:"Groups"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListGroupsForUserResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListGroupsForUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupsRequest struct {
	*ksyunhttp.BaseRequest
	MaxItems *string `json:"MaxItems,omitempty" name:"MaxItems"`
	Marker   *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *ListGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupsResponse struct {
	*ksyunhttp.BaseResponse
	ListGroupsResult struct {
		Groups struct {
			Member []struct {
				GroupId     *string `json:"GroupId" name:"GroupId"`
				Path        *string `json:"Path" name:"Path"`
				GroupName   *string `json:"GroupName" name:"GroupName"`
				Description *string `json:"Description" name:"Description"`
				CreateDate  *string `json:"CreateDate" name:"CreateDate"`
				Krn         *string `json:"Krn" name:"Krn"`
				UserCount   *int    `json:"UserCount" name:"UserCount"`
				PolicyCount *int    `json:"PolicyCount" name:"PolicyCount"`
			} `json:"Member"`
		} `json:"Groups" name:"Groups"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListGroupsResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromGroupRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	UserName  *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RemoveUserFromGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromGroupResponse struct {
	*ksyunhttp.BaseResponse
	Result    *bool   `json:"result" name:"result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveUserFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupRequest struct {
	*ksyunhttp.BaseRequest
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	MaxItems  *int    `json:"MaxItems,omitempty" name:"MaxItems"`
	Page      *int    `json:"Page,omitempty" name:"Page"`
}

func (r *ListUsersForGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersForGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListUsersForGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupResponse struct {
	*ksyunhttp.BaseResponse
	ListUsersForGroupResult struct {
		Users []struct {
			UserId                *string `json:"UserId" name:"UserId"`
			Path                  *string `json:"Path" name:"Path"`
			UserName              *string `json:"UserName" name:"UserName"`
			RealName              *string `json:"RealName" name:"RealName"`
			CreateDate            *string `json:"CreateDate" name:"CreateDate"`
			Phone                 *string `json:"Phone" name:"Phone"`
			CountryMobileCode     *string `json:"CountryMobileCode" name:"CountryMobileCode"`
			IsInternational       *int    `json:"isInternational" name:"isInternational"`
			Email                 *string `json:"Email" name:"Email"`
			PhoneVerified         *string `json:"PhoneVerified" name:"PhoneVerified"`
			EmailVerified         *string `json:"EmailVerified" name:"EmailVerified"`
			Remark                *string `json:"Remark" name:"Remark"`
			Krn                   *string `json:"Krn" name:"Krn"`
			PasswordResetRequired *bool   `json:"PasswordResetRequired" name:"PasswordResetRequired"`
			UpdateDate            *string `json:"UpdateDate" name:"UpdateDate"`
			Id                    *int    `json:"Id" name:"Id"`
		} `json:"Users" name:"Users"`
		IsTruncated *bool   `json:"IsTruncated" name:"IsTruncated"`
		Marker      *string `json:"Marker" name:"Marker"`
	} `json:"ListUsersForGroupResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListUsersForGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersForGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAllUserAccessKeysRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ListAllUserAccessKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllUserAccessKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAllUserAccessKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAllUserAccessKeysResponse struct {
	*ksyunhttp.BaseResponse
	AccessKeyList []struct {
		UserName       *string `json:"UserName" name:"UserName"`
		AccessKey      *string `json:"AccessKey" name:"AccessKey"`
		LastLoginTime  *string `json:"LastLoginTime" name:"LastLoginTime"`
		AkLastUsedTime *string `json:"AkLastUsedTime" name:"AkLastUsedTime"`
	} `json:"AccessKeyList"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAllUserAccessKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllUserAccessKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertInstanceToESRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId    *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	ProductLine  *string `json:"ProductLine,omitempty" name:"ProductLine"`
	ProductGroup *int    `json:"ProductGroup,omitempty" name:"ProductGroup"`
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	RegionEn     *string `json:"RegionEn,omitempty" name:"RegionEn"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *InsertInstanceToESRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertInstanceToESRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "InsertInstanceToESRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InsertInstanceToESResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Status     *int    `json:"Status" name:"Status"`
		Result     *string `json:"Result" name:"Result"`
		Reason     *string `json:"Reason" name:"Reason"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *InsertInstanceToESResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertInstanceToESResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelInstanceFromESRequest struct {
	*ksyunhttp.BaseRequest
	ProductLine *string `json:"ProductLine,omitempty" name:"ProductLine"`
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DelInstanceFromESRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelInstanceFromESRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DelInstanceFromESRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DelInstanceFromESResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Status     *int    `json:"Status" name:"Status"`
		Result     *string `json:"Result" name:"Result"`
		Reason     *string `json:"Reason" name:"Reason"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DelInstanceFromESResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelInstanceFromESResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountAllProjectsByParamsRequest struct {
	*ksyunhttp.BaseRequest
	Ps             *int    `json:"Ps,omitempty" name:"Ps"`
	Pn             *int    `json:"Pn,omitempty" name:"Pn"`
	ParamUserName  *string `json:"ParamUserName,omitempty" name:"ParamUserName"`
	ParamAccessKey *string `json:"ParamAccessKey,omitempty" name:"ParamAccessKey"`
}

func (r *GetAccountAllProjectsByParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountAllProjectsByParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetAccountAllProjectsByParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountAllProjectsByParamsResponse struct {
	*ksyunhttp.BaseResponse
	ListProjectResult struct {
		Total       *int `json:"Total" name:"Total"`
		ProjectList []struct {
			ProjectId   *int    `json:"ProjectId" name:"ProjectId"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			ProjectDesc *string `json:"ProjectDesc" name:"ProjectDesc"`
			Status      *int    `json:"Status" name:"Status"`
			Krn         *string `json:"Krn" name:"Krn"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		} `json:"ProjectList" name:"ProjectList"`
	} `json:"ListProjectResult"`
}

func (r *GetAccountAllProjectsByParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountAllProjectsByParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetUserSsoSettingsRequest struct {
	*ksyunhttp.BaseRequest
	Status   *int    `json:"Status,omitempty" name:"Status"`
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	Domain   *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *SetUserSsoSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUserSsoSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetUserSsoSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetUserSsoSettingsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	UserSsoSettings struct {
		Status   *int    `json:"Status" name:"Status"`
		Metadata *string `json:"Metadata" name:"Metadata"`
		Domain   *string `json:"Domain" name:"Domain"`
	} `json:"UserSsoSettings"`
}

func (r *SetUserSsoSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUserSsoSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserSsoSettingsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetUserSsoSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserSsoSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUserSsoSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetUserSsoSettingsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	UserSsoSettings struct {
		Status   *int    `json:"Status" name:"Status"`
		Metadata *string `json:"Metadata" name:"Metadata"`
		Domain   *string `json:"Domain" name:"Domain"`
	} `json:"UserSsoSettings"`
}

func (r *GetUserSsoSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserSsoSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEffectivePoliciesRequest struct {
	*ksyunhttp.BaseRequest
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	Page       *int    `json:"Page,omitempty" name:"Page"`
	MaxItems   *int    `json:"MaxItems,omitempty" name:"MaxItems"`
}

func (r *GetEffectivePoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEffectivePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetEffectivePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetEffectivePoliciesResponse struct {
	*ksyunhttp.BaseResponse
	ListPoliciesResult struct {
		Policies []struct {
			CreateDate       *string `json:"CreateDate" name:"CreateDate"`
			DefaultVersionId *string `json:"DefaultVersionId" name:"DefaultVersionId"`
			Description      *string `json:"Description" name:"Description"`
			Krn              *string `json:"Krn" name:"Krn"`
			Path             *string `json:"Path" name:"Path"`
			PolicyId         *string `json:"PolicyId" name:"PolicyId"`
			PolicyName       *string `json:"PolicyName" name:"PolicyName"`
			ServiceId        *int    `json:"ServiceId" name:"ServiceId"`
			ServiceName      *string `json:"ServiceName" name:"ServiceName"`
			ServiceViewName  *string `json:"ServiceViewName" name:"ServiceViewName"`
			PolicyType       *int    `json:"PolicyType" name:"PolicyType"`
			CreateMode       *int    `json:"CreateMode" name:"CreateMode"`
			UpdateDate       *string `json:"UpdateDate" name:"UpdateDate"`
			AttachmentCount  *int    `json:"AttachmentCount" name:"AttachmentCount"`
			PolicyDocument   *string `json:"PolicyDocument" name:"PolicyDocument"`
		} `json:"Policies" name:"Policies"`
	} `json:"ListPoliciesResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetEffectivePoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEffectivePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
