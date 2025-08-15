package v20180627

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	Name         *string `json:"Name,omitempty" name:"Name"`
	Capacity     *string `json:"Capacity,omitempty" name:"Capacity"`
	SlaveNum     *string `json:"SlaveNum,omitempty" name:"SlaveNum"`
	NetType      *string `json:"NetType,omitempty" name:"NetType"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	VnetId       *string `json:"VnetId,omitempty" name:"VnetId"`
	BillType     *string `json:"BillType,omitempty" name:"BillType"`
	Duration     *string `json:"Duration,omitempty" name:"Duration"`
	DurationUnit *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	PassWord     *string `json:"PassWord,omitempty" name:"PassWord"`
	IamProjectId *string `json:"IamProjectId,omitempty" name:"IamProjectId"`
	Engine       *string `json:"Engine,omitempty" name:"Engine"`
}

func (r *CreateCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *CreateCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DeleteCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	CacheId  *string `json:"CacheId,omitempty" name:"CacheId"`
	Capacity *string `json:"Capacity,omitempty" name:"Capacity"`
	Engine   *string `json:"Engine,omitempty" name:"Engine"`
}

func (r *ResizeCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResizeCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ResizeCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheClustersRequest struct {
	*ksyunhttp.BaseRequest
	Engine       *string `json:"Engine,omitempty" name:"Engine"`
	CacheId      *string `json:"CacheId,omitempty" name:"CacheId"`
	Name         *string `json:"Name,omitempty" name:"Name"`
	Vip          *string `json:"Vip,omitempty" name:"Vip"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	VnetId       *string `json:"VnetId,omitempty" name:"VnetId"`
	Offset       *string `json:"Offset,omitempty" name:"Offset"`
	Limit        *string `json:"Limit,omitempty" name:"Limit"`
	OrderBy      *string `json:"OrderBy,omitempty" name:"OrderBy"`
	IamProjectId *string `json:"IamProjectId,omitempty" name:"IamProjectId"`
}

func (r *DescribeCacheClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCacheClustersResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeCacheClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
	Engine  *string `json:"Engine,omitempty" name:"Engine"`
}

func (r *DescribeCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlushCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *FlushCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type FlushCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *FlushCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FlushCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
	Name    *string `json:"Name,omitempty" name:"Name"`
	Engine  *string `json:"Engine,omitempty" name:"Engine"`
}

func (r *RenameCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenameCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *RenameCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordRequest struct {
	*ksyunhttp.BaseRequest
	CacheId  *string `json:"CacheId,omitempty" name:"CacheId"`
	Password *string `json:"Password,omitempty" name:"Password"`
	Engine   *string `json:"Engine,omitempty" name:"Engine"`
}

func (r *UpdatePasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdatePasswordResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *UpdatePasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheSecurityRulesRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeCacheSecurityRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCacheSecurityRulesResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeCacheSecurityRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DeleteCacheSecurityRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCacheSecurityRuleResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteCacheSecurityRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCacheSecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheSecurityRulesRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *SetCacheSecurityRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetCacheSecurityRulesResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *SetCacheSecurityRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCacheSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRegionsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailabilityZonesRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeAvailabilityZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAvailabilityZonesResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeAvailabilityZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
