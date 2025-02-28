package v20180627

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
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

func (r *CreateCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		CacheId    *string `json:"CacheId"`
		Name       *string `json:"Name"`
		Size       *string `json:"Size"`
		Port       *string `json:"Port"`
		SubOrderId *string `json:"SubOrderId"`
	} `json:"Data"`
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

func (r *DeleteCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *ResizeCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResizeCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResizeCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *DescribeCacheClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheClustersResponse struct {
	*ksyunhttp.BaseResponse
	reqId   *string `json:"reqId" name:"reqId"`
	message *string `json:"message" name:"message"`
	Data    struct {
		List []struct {
			cacheId         *string `json:"cacheId"`
			name            *string `json:"name"`
			securityGroupId *string `json:"securityGroupId"`
			engine          *string `json:"engine"`
			vip             *string `json:"vip"`
			createTime      *string `json:"createTime"`
			vpcId           *string `json:"vpcId"`
			vnetId          *string `json:"vnetId"`
			iamProjectId    *string `json:"iamProjectId"`
			iamProjectName  *string `json:"iamProjectName"`
		} `json:"List"`
	} `json:"Data"`
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

func (r *DescribeCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		CacheId          *string `json:"CacheId"`
		Az               *string `json:"Az"`
		Name             *string `json:"Name"`
		SecurityGroupId  *string `json:"SecurityGroupId"`
		Engine           *string `json:"Engine"`
		Vip              *string `json:"Vip"`
		SlaveVip         *string `json:"SlaveVip"`
		CreateTime       *string `json:"CreateTime"`
		VpcId            *string `json:"VpcId"`
		VnetId           *string `json:"VnetId"`
		TimingSwitch     *string `json:"TimingSwitch"`
		Timezone         *string `json:"Timezone"`
		SubOrderId       *string `json:"SubOrderId"`
		ProductId        *string `json:"ProductId"`
		ServiceBeginTime *string `json:"ServiceBeginTime"`
		ServiceEndTime   *string `json:"ServiceEndTime"`
		IamProjectId     *string `json:"IamProjectId"`
		IamProjectName   *string `json:"IamProjectName"`
		Protocol         *string `json:"Protocol"`
	} `json:"Data"`
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

func (r *FlushCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "FlushCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FlushCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *RenameCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *UpdatePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdatePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordResponse struct {
	*ksyunhttp.BaseResponse
	reqId *string `json:"reqId" name:"reqId"`
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

func (r *DescribeCacheSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheSecurityRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteCacheSecurityRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCacheSecurityRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *SetCacheSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetCacheSecurityRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheSecurityRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*ksyunhttp.BaseResponse
	RegionSet []struct {
		RegionName *string `json:"RegionName"`
		Region     *string `json:"Region"`
	} `json:"RegionSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *DescribeAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailabilityZonesResponse struct {
	*ksyunhttp.BaseResponse
	AvailabilityZoneSet []struct {
		Region           *string `json:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone"`
	} `json:"AvailabilityZoneSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeAvailabilityZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
