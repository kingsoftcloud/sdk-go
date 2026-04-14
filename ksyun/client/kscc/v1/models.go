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
		UserName    *string `json:"UserName" name:"UserName"`
		QuotaAmount *int    `json:"QuotaAmount" name:"QuotaAmount"`
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
		QuotaAmount  *int     `json:"QuotaAmount" name:"QuotaAmount"`
		UsageAmount  *float64 `json:"UsageAmount" name:"UsageAmount"`
		UsagePercent *int     `json:"UsagePercent" name:"UsagePercent"`
	} `json:"DescribeUserCostSummaryResult"`
}

func (r *DescribeUserCostSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCostSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
