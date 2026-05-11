package v1

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type UpdateUserQuotaRequest struct {
	*ksyunhttp.BaseRequest
	UserName    *string  `json:"UserName,omitempty" name:"UserName"`
	QuotaAmount *float64 `json:"QuotaAmount,omitempty" name:"QuotaAmount"`
}

func (r *UpdateUserQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateUserQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	UpdateUserQuotaResult struct {
		UserName    *string  `json:"UserName" name:"UserName"`
		QuotaAmount *float64 `json:"QuotaAmount" name:"QuotaAmount"`
	} `json:"UpdateUserQuotaResult"`
}

func (r *UpdateUserQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCostSummaryRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
	Month    *string `json:"Month,omitempty" name:"Month"`
}

func (r *DescribeUserCostSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeUserCostSummaryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                     *string `json:"RequestId" name:"RequestId"`
	DescribeUserCostSummaryResult struct {
		UserCnName   *string  `json:"UserCnName" name:"UserCnName"`
		UserName     *string  `json:"UserName" name:"UserName"`
		DeptId       *string  `json:"DeptId" name:"DeptId"`
		DeptName     *string  `json:"DeptName" name:"DeptName"`
		Month        *string  `json:"Month" name:"Month"`
		QuotaAmount  *float64 `json:"QuotaAmount" name:"QuotaAmount"`
		QuotaLimitBy *string  `json:"QuotaLimitBy" name:"QuotaLimitBy"`
		UsageAmount  *float64 `json:"UsageAmount" name:"UsageAmount"`
		UsagePercent *float64 `json:"UsagePercent" name:"UsagePercent"`
	} `json:"DescribeUserCostSummaryResult"`
}

func (r *DescribeUserCostSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCostSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAiLogDetailByIdsRequest struct {
	*ksyunhttp.BaseRequest
	MessageIds *string `json:"MessageIds,omitempty" name:"MessageIds"`
}

func (r *DescribeAiLogDetailByIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAiLogDetailByIdsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                      *string `json:"RequestId" name:"RequestId"`
	DescribeAiLogDetailByIdsResult struct {
		Total *int `json:"Total" name:"Total"`
		Items []struct {
			MessageId    *string `json:"MessageId" name:"MessageId"`
			Status       *string `json:"Status" name:"Status"`
			StatusCode   *int    `json:"StatusCode" name:"StatusCode"`
			RequestJson  *string `json:"RequestJson" name:"RequestJson"`
			ResponseJson *string `json:"ResponseJson" name:"ResponseJson"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeAiLogDetailByIdsResult"`
}

func (r *DescribeAiLogDetailByIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAiLogDetailByIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAiLogDetailRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	UserName  *string `json:"UserName,omitempty" name:"UserName"`
	ModelList *string `json:"ModelList,omitempty" name:"ModelList"`
	Page      *int    `json:"Page,omitempty" name:"Page"`
	Size      *int    `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeAiLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAiLogDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"RequestId" name:"RequestId"`
	DescribeAiLogDetailResult struct {
		Total      *int `json:"Total" name:"Total"`
		Page       *int `json:"Page" name:"Page"`
		Size       *int `json:"Size" name:"Size"`
		TotalPages *int `json:"TotalPages" name:"TotalPages"`
		Items      []struct {
			MessageId   *string  `json:"MessageId" name:"MessageId"`
			CreateTime  *string  `json:"CreateTime" name:"CreateTime"`
			UserName    *string  `json:"UserName" name:"UserName"`
			Model       *string  `json:"Model" name:"Model"`
			Tokens      *int     `json:"Tokens" name:"Tokens"`
			QuotaAmount *float64 `json:"QuotaAmount" name:"QuotaAmount"`
			StatusCode  *int     `json:"StatusCode" name:"StatusCode"`
			Status      *string  `json:"Status" name:"Status"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeAiLogDetailResult"`
}

func (r *DescribeAiLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAiLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationTreeRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeOrganizationTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeOrganizationTreeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                      *string `json:"RequestId" name:"RequestId"`
	DescribeOrganizationTreeResult struct {
		Total *int `json:"Total" name:"Total"`
		Items []struct {
			Id        *string `json:"Id" name:"Id"`
			ParentId  *string `json:"ParentId" name:"ParentId"`
			Name      *string `json:"Name" name:"Name"`
			AbsPath   *string `json:"AbsPath" name:"AbsPath"`
			AiEnabled *int    `json:"AiEnabled" name:"AiEnabled"`
			Children  []struct {
				Id   *string `json:"Id" name:"Id"`
				Name *string `json:"Name" name:"Name"`
			} `json:"Children"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeOrganizationTreeResult"`
}

func (r *DescribeOrganizationTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelMetricsRequest struct {
	*ksyunhttp.BaseRequest
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	ModelName    *string `json:"ModelName,omitempty" name:"ModelName"`
	TimeInterval *string `json:"TimeInterval,omitempty" name:"TimeInterval"`
}

func (r *DescribeModelMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModelMetricsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                  *string `json:"RequestId" name:"RequestId"`
	DescribeModelMetricsResult struct {
		TimeInterval *string `json:"TimeInterval" name:"TimeInterval"`
		Items        []struct {
			ModelName        *string  `json:"ModelName" name:"ModelName"`
			TimeStr          *string  `json:"TimeStr" name:"TimeStr"`
			RequestCounter   *int     `json:"RequestCounter" name:"RequestCounter"`
			TotalTokens      *int     `json:"TotalTokens" name:"TotalTokens"`
			PromptTokens     *int     `json:"PromptTokens" name:"PromptTokens"`
			CompletionTokens *int     `json:"CompletionTokens" name:"CompletionTokens"`
			Tpm              *float64 `json:"Tpm" name:"Tpm"`
			Rpm              *float64 `json:"Rpm" name:"Rpm"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeModelMetricsResult"`
}

func (r *DescribeModelMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTokenUsageRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	UserName   *string `json:"UserName,omitempty" name:"UserName"`
	Department *string `json:"Department,omitempty" name:"Department"`
}

func (r *DescribeUserTokenUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeUserTokenUsageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                    *string `json:"RequestId" name:"RequestId"`
	DescribeUserTokenUsageResult struct {
		Total *int `json:"Total" name:"Total"`
		Items []struct {
			UserName     *string  `json:"UserName" name:"UserName"`
			StatDate     *string  `json:"StatDate" name:"StatDate"`
			TotalTokens  *int     `json:"TotalTokens" name:"TotalTokens"`
			InputTokens  *int     `json:"InputTokens" name:"InputTokens"`
			OutputTokens *int     `json:"OutputTokens" name:"OutputTokens"`
			CachedTokens *int     `json:"CachedTokens" name:"CachedTokens"`
			QuotaAmount  *float64 `json:"QuotaAmount" name:"QuotaAmount"`
			AbsPath      *string  `json:"AbsPath" name:"AbsPath"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeUserTokenUsageResult"`
}

func (r *DescribeUserTokenUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTokenUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaListRequest struct {
	*ksyunhttp.BaseRequest
	Keyword  *string `json:"Keyword,omitempty" name:"Keyword"`
	PageNum  *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
	Month    *string `json:"Month,omitempty" name:"Month"`
	SortKey  *string `json:"SortKey,omitempty" name:"SortKey"`
	SortType *string `json:"SortType,omitempty" name:"SortType"`
}

func (r *DescribeUserQuotaListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeUserQuotaListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                   *string `json:"RequestId" name:"RequestId"`
	DescribeUserQuotaListResult struct {
		Total    *int `json:"Total" name:"Total"`
		PageNum  *int `json:"PageNum" name:"PageNum"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		Items    []struct {
			UserName     *string  `json:"UserName" name:"UserName"`
			UserCnName   *string  `json:"UserCnName" name:"UserCnName"`
			DeptId       *string  `json:"DeptId" name:"DeptId"`
			DeptName     *string  `json:"DeptName" name:"DeptName"`
			QuotaAmount  *float64 `json:"QuotaAmount" name:"QuotaAmount"`
			UsageAmount  *float64 `json:"UsageAmount" name:"UsageAmount"`
			UsagePercent *float64 `json:"UsagePercent" name:"UsagePercent"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeUserQuotaListResult"`
}

func (r *DescribeUserQuotaListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserQuotaListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaGlobalConfigRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeQuotaGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeQuotaGlobalConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                       *string `json:"RequestId" name:"RequestId"`
	DescribeQuotaGlobalConfigResult struct {
		WarnThreshold      *int     `json:"WarnThreshold" name:"WarnThreshold"`
		Discount           *float64 `json:"Discount" name:"Discount"`
		ContactsJson       *string  `json:"ContactsJson" name:"ContactsJson"`
		AccountEnabled     *int     `json:"AccountEnabled" name:"AccountEnabled"`
		DeptEnabled        *int     `json:"DeptEnabled" name:"DeptEnabled"`
		MemberEnabled      *int     `json:"MemberEnabled" name:"MemberEnabled"`
		DefaultMemberQuota *float64 `json:"DefaultMemberQuota" name:"DefaultMemberQuota"`
	} `json:"DescribeQuotaGlobalConfigResult"`
}

func (r *DescribeQuotaGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateQuotaGlobalConfigRequest struct {
	*ksyunhttp.BaseRequest
	WarnThreshold      *int     `json:"WarnThreshold,omitempty" name:"WarnThreshold"`
	Discount           *float64 `json:"Discount,omitempty" name:"Discount"`
	ContactsJson       *string  `json:"ContactsJson,omitempty" name:"ContactsJson"`
	AccountEnabled     *int     `json:"AccountEnabled,omitempty" name:"AccountEnabled"`
	DeptEnabled        *int     `json:"DeptEnabled,omitempty" name:"DeptEnabled"`
	MemberEnabled      *int     `json:"MemberEnabled,omitempty" name:"MemberEnabled"`
	DefaultMemberQuota *float64 `json:"DefaultMemberQuota,omitempty" name:"DefaultMemberQuota"`
}

func (r *UpdateQuotaGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateQuotaGlobalConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                     *string `json:"RequestId" name:"RequestId"`
	UpdateQuotaGlobalConfigResult struct {
		WarnThreshold      *int     `json:"WarnThreshold" name:"WarnThreshold"`
		Discount           *float64 `json:"Discount" name:"Discount"`
		ContactsJson       *string  `json:"ContactsJson" name:"ContactsJson"`
		AccountEnabled     *int     `json:"AccountEnabled" name:"AccountEnabled"`
		DeptEnabled        *int     `json:"DeptEnabled" name:"DeptEnabled"`
		MemberEnabled      *int     `json:"MemberEnabled" name:"MemberEnabled"`
		DefaultMemberQuota *float64 `json:"DefaultMemberQuota" name:"DefaultMemberQuota"`
	} `json:"UpdateQuotaGlobalConfigResult"`
}

func (r *UpdateQuotaGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateQuotaGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountQuotaRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeAccountQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAccountQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                  *string `json:"RequestId" name:"RequestId"`
	DescribeAccountQuotaResult struct {
		QuotaAmount  *float64 `json:"QuotaAmount" name:"QuotaAmount"`
		UsageAmount  *float64 `json:"UsageAmount" name:"UsageAmount"`
		UsagePercent *float64 `json:"UsagePercent" name:"UsagePercent"`
	} `json:"DescribeAccountQuotaResult"`
}

func (r *DescribeAccountQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountQuotaRequest struct {
	*ksyunhttp.BaseRequest
	QuotaAmount *float64 `json:"QuotaAmount,omitempty" name:"QuotaAmount"`
}

func (r *UpdateAccountQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateAccountQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	UpdateAccountQuotaResult struct {
		QuotaAmount  *float64 `json:"QuotaAmount" name:"QuotaAmount"`
		UsageAmount  *float64 `json:"UsageAmount" name:"UsageAmount"`
		UsagePercent *float64 `json:"UsagePercent" name:"UsagePercent"`
	} `json:"UpdateAccountQuotaResult"`
}

func (r *UpdateAccountQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccountQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountQuotaRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DeleteAccountQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteAccountQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	DeleteAccountQuotaResult struct {
		AccountId *string `json:"AccountId" name:"AccountId"`
	} `json:"DeleteAccountQuotaResult"`
}

func (r *DeleteAccountQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDefaultMemberQuotaRequest struct {
	*ksyunhttp.BaseRequest
	DefaultMemberQuota *float64 `json:"DefaultMemberQuota,omitempty" name:"DefaultMemberQuota"`
}

func (r *UpdateDefaultMemberQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDefaultMemberQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                      *string `json:"RequestId" name:"RequestId"`
	UpdateDefaultMemberQuotaResult struct {
		WarnThreshold      *int     `json:"WarnThreshold" name:"WarnThreshold"`
		Discount           *float64 `json:"Discount" name:"Discount"`
		ContactsJson       *string  `json:"ContactsJson" name:"ContactsJson"`
		AccountEnabled     *int     `json:"AccountEnabled" name:"AccountEnabled"`
		DeptEnabled        *int     `json:"DeptEnabled" name:"DeptEnabled"`
		MemberEnabled      *int     `json:"MemberEnabled" name:"MemberEnabled"`
		DefaultMemberQuota *float64 `json:"DefaultMemberQuota" name:"DefaultMemberQuota"`
	} `json:"UpdateDefaultMemberQuotaResult"`
}

func (r *UpdateDefaultMemberQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefaultMemberQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDefaultMemberQuotaRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DeleteDefaultMemberQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDefaultMemberQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                      *string `json:"RequestId" name:"RequestId"`
	DeleteDefaultMemberQuotaResult struct {
		DefaultMemberQuota *float64 `json:"DefaultMemberQuota" name:"DefaultMemberQuota"`
	} `json:"DeleteDefaultMemberQuotaResult"`
}

func (r *DeleteDefaultMemberQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDefaultMemberQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeptQuotaListRequest struct {
	*ksyunhttp.BaseRequest
	Keyword  *string `json:"Keyword,omitempty" name:"Keyword"`
	PageNum  *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
	SortKey  *string `json:"SortKey,omitempty" name:"SortKey"`
	SortType *string `json:"SortType,omitempty" name:"SortType"`
}

func (r *DescribeDeptQuotaListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDeptQuotaListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                   *string `json:"RequestId" name:"RequestId"`
	DescribeDeptQuotaListResult struct {
		Total    *int `json:"Total" name:"Total"`
		PageNum  *int `json:"PageNum" name:"PageNum"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		Items    []struct {
			DeptId       *string  `json:"DeptId" name:"DeptId"`
			DeptName     *string  `json:"DeptName" name:"DeptName"`
			QuotaAmount  *float64 `json:"QuotaAmount" name:"QuotaAmount"`
			UsageAmount  *float64 `json:"UsageAmount" name:"UsageAmount"`
			UsagePercent *float64 `json:"UsagePercent" name:"UsagePercent"`
		} `json:"Items" name:"Items"`
	} `json:"DescribeDeptQuotaListResult"`
}

func (r *DescribeDeptQuotaListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeptQuotaListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDeptQuotaRequest struct {
	*ksyunhttp.BaseRequest
	DeptId      *string  `json:"DeptId,omitempty" name:"DeptId"`
	QuotaAmount *float64 `json:"QuotaAmount,omitempty" name:"QuotaAmount"`
}

func (r *UpdateDeptQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDeptQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	UpdateDeptQuotaResult struct {
		DeptId      *string  `json:"DeptId" name:"DeptId"`
		QuotaAmount *float64 `json:"QuotaAmount" name:"QuotaAmount"`
	} `json:"UpdateDeptQuotaResult"`
}

func (r *UpdateDeptQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDeptQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeptQuotaRequest struct {
	*ksyunhttp.BaseRequest
	DeptId *string `json:"DeptId,omitempty" name:"DeptId"`
}

func (r *DeleteDeptQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDeptQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	DeleteDeptQuotaResult struct {
		DeptId *string `json:"DeptId" name:"DeptId"`
	} `json:"DeleteDeptQuotaResult"`
}

func (r *DeleteDeptQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeptQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserQuotaRequest struct {
	*ksyunhttp.BaseRequest
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DeleteUserQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteUserQuotaResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	DeleteUserQuotaResult struct {
		UserName *string `json:"UserName" name:"UserName"`
	} `json:"DeleteUserQuotaResult"`
}

func (r *DeleteUserQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
