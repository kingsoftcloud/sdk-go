package v20160701

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateCacheParameterGroupParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" name:"ParameterName"`
	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`
}
type ModifyCacheParameterGroupParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" name:"ParameterName"`
	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`
}
type InstallPluginsPlugins struct {
	PluginName    *string `json:"PluginName,omitempty" name:"PluginName"`
	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
}

type CreateCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name            *string `json:"Name,omitempty" name:"Name"`
	PassWord        *string `json:"PassWord,omitempty" name:"PassWord"`
	Mode            *int    `json:"Mode,omitempty" name:"Mode"`
	Vip             *string `json:"Vip,omitempty" name:"Vip"`
	Capacity        *int    `json:"Capacity,omitempty" name:"Capacity"`
	VpcId           *string `json:"VpcId,omitempty" name:"VpcId"`
	VnetId          *string `json:"VnetId,omitempty" name:"VnetId"`
	BillType        *int    `json:"BillType,omitempty" name:"BillType"`
	Duration        *int    `json:"Duration,omitempty" name:"Duration"`
	IamProjectId    *string `json:"IamProjectId,omitempty" name:"IamProjectId"`
	Protocol        *string `json:"Protocol,omitempty" name:"Protocol"`
	BackupTimezone  *string `json:"BackupTimezone,omitempty" name:"BackupTimezone"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SlaveNum        *int    `json:"SlaveNum,omitempty" name:"SlaveNum"`
	SlaveVip        *string `json:"SlaveVip,omitempty" name:"SlaveVip"`
	PrepareAzName   *string `json:"PrepareAzName,omitempty" name:"PrepareAzName"`
	RrAzName        *string `json:"RrAzName,omitempty" name:"RrAzName"`
	ShardNum        *int    `json:"ShardNum,omitempty" name:"ShardNum"`
	ShardSize       *int    `json:"ShardSize,omitempty" name:"ShardSize"`
	Separation      *int    `json:"Separation,omitempty" name:"Separation"`
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
		CacheId    *string `json:"CacheId" name:"CacheId"`
		Name       *string `json:"Name" name:"Name"`
		Size       *string `json:"Size" name:"Size"`
		Port       *string `json:"Port" name:"Port"`
		SubOrderId *string `json:"SubOrderId" name:"SubOrderId"`
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
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
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

type DescribeCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
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
		CacheId          *string `json:"CacheId" name:"CacheId"`
		Az               *string `json:"Az" name:"Az"`
		Name             *string `json:"Name" name:"Name"`
		SecurityGroupId  *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Engine           *string `json:"Engine" name:"Engine"`
		Vip              *string `json:"Vip" name:"Vip"`
		SlaveVip         *string `json:"SlaveVip" name:"SlaveVip"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		VpcId            *string `json:"VpcId" name:"VpcId"`
		VnetId           *string `json:"VnetId" name:"VnetId"`
		TimingSwitch     *string `json:"TimingSwitch" name:"TimingSwitch"`
		Timezone         *string `json:"Timezone" name:"Timezone"`
		SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
		ProductId        *string `json:"ProductId" name:"ProductId"`
		ServiceBeginTime *string `json:"ServiceBeginTime" name:"ServiceBeginTime"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		IamProjectId     *string `json:"IamProjectId" name:"IamProjectId"`
		IamProjectName   *string `json:"IamProjectName" name:"IamProjectName"`
		Protocol         *string `json:"Protocol" name:"Protocol"`
	} `json:"Data"`
}

func (r *DescribeCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheClustersRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Name          *string `json:"Name,omitempty" name:"Name"`
	Vip           *string `json:"Vip,omitempty" name:"Vip"`
	VpcId         *string `json:"VpcId,omitempty" name:"VpcId"`
	VnetId        *string `json:"VnetId,omitempty" name:"VnetId"`
	Offset        *string `json:"Offset,omitempty" name:"Offset"`
	Limit         *string `json:"Limit,omitempty" name:"Limit"`
	OrderBy       *string `json:"OrderBy,omitempty" name:"OrderBy"`
	IamProjectId  *string `json:"IamProjectId,omitempty" name:"IamProjectId"`
	FuzzySearch   *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Protocol      *string `json:"Protocol,omitempty" name:"Protocol"`
	TagKey        *string `json:"TagKey,omitempty" name:"TagKey"`
	TagValue      *string `json:"TagValue,omitempty" name:"TagValue"`
	Mode          *string `json:"Mode,omitempty" name:"Mode"`
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
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		List []struct {
			CacheId          *string `json:"cacheId" name:"cacheId"`
			Region           *string `json:"region" name:"region"`
			Az               *string `json:"az" name:"az"`
			Name             *string `json:"name" name:"name"`
			Engine           *string `json:"engine" name:"engine"`
			Vip              *string `json:"vip" name:"vip"`
			CreateTime       *string `json:"createTime" name:"createTime"`
			VpcId            *string `json:"vpcId" name:"vpcId"`
			VnetId           *string `json:"vnetId" name:"vnetId"`
			ServiceBeginTime *string `json:"serviceBeginTime" name:"serviceBeginTime"`
			ServiceEndTime   *string `json:"serviceEndTime" name:"serviceEndTime"`
			IamProjectId     *string `json:"iamProjectId" name:"iamProjectId"`
			IamProjectName   *string `json:"iamProjectName" name:"iamProjectName"`
			Protocol         *string `json:"protocol" name:"protocol"`
			Eip              *string `json:"eip" name:"eip"`
			EipRo            *string `json:"eipRo" name:"eipRo"`
			Tags             struct {
			} `json:"Tags"`
			Area            *string `json:"area" name:"area"`
			DirectVips      *string `json:"directVips" name:"directVips"`
			DirectSupported *string `json:"directSupported" name:"directSupported"`
		} `json:"List" name:"List"`
	} `json:"Data"`
}

func (r *DescribeCacheClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlushCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	DatabaseNo    *string `json:"DatabaseNo,omitempty" name:"DatabaseNo"`
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
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Name          *string `json:"Name,omitempty" name:"Name"`
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
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *RenameCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Capacity      *int    `json:"Capacity,omitempty" name:"Capacity"`
	ShardSize     *int    `json:"ShardSize,omitempty" name:"ShardSize"`
	ShardNum      *int    `json:"ShardNum,omitempty" name:"ShardNum"`
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

type DescribeCacheParametersRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *DescribeCacheParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheParametersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Name         *string `json:"Name" name:"Name"`
		Desc         *string `json:"Desc" name:"Desc"`
		DefaultValue *string `json:"DefaultValue" name:"DefaultValue"`
		CurrentValue *string `json:"CurrentValue" name:"CurrentValue"`
		Validity     struct {
			Type     *string   `json:"type" name:"type"`
			DataType *string   `json:"dataType" name:"dataType"`
			Value    *string   `json:"value" name:"value"`
			Values   []*string `json:"Values" name:"Values"`
			Min      *string   `json:"min" name:"min"`
			Max      *string   `json:"max" name:"max"`
		} `json:"Validity" name:"Validity"`
	} `json:"Data"`
}

func (r *DescribeCacheParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheParametersRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone            *string   `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId                  *string   `json:"CacheId,omitempty" name:"CacheId"`
	Protocol                 *string   `json:"Protocol,omitempty" name:"Protocol"`
	ParametersParameterName  []*string `json:"Parameters.ParameterName,omitempty" name:"Parameters.ParameterName"`
	ParametersParameterValue []*string `json:"Parameters.ParameterValue,omitempty" name:"Parameters.ParameterValue"`
	ResetAllParameters       *bool     `json:"ResetAllParameters,omitempty" name:"ResetAllParameters"`
}

func (r *SetCacheParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCacheParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetCacheParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheParametersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *SetCacheParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCacheParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheDefaultParametersRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	ParamVersion  *string `json:"ParamVersion,omitempty" name:"ParamVersion"`
}

func (r *DescribeCacheDefaultParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheDefaultParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheDefaultParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheDefaultParametersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Name         *string `json:"Name" name:"Name"`
		Desc         *string `json:"Desc" name:"Desc"`
		DefaultValue *string `json:"DefaultValue" name:"DefaultValue"`
		CurrentValue *string `json:"CurrentValue" name:"CurrentValue"`
		Validity     struct {
			Type     *string   `json:"type" name:"type"`
			DataType *string   `json:"dataType" name:"dataType"`
			Value    []*string `json:"Value" name:"Value"`
		} `json:"Validity" name:"Validity"`
	} `json:"Data"`
}

func (r *DescribeCacheDefaultParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheDefaultParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone          *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheParameterGroupIds *string `json:"CacheParameterGroupIds,omitempty" name:"CacheParameterGroupIds"`
	CacheId                *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *SetCacheParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCacheParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetCacheParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetCacheParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCacheParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCacheParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string                                `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name          *string                                `json:"Name,omitempty" name:"Name"`
	Description   *string                                `json:"Description,omitempty" name:"Description"`
	ParamVersion  *string                                `json:"ParamVersion,omitempty" name:"ParamVersion"`
	Parameters    []*CreateCacheParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *CreateCacheParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCacheParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCacheParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCacheParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id          *string `json:"Id" name:"Id"`
		Name        *string `json:"Name" name:"Name"`
		Description *string `json:"Description" name:"Description"`
		Engine      *string `json:"Engine" name:"Engine"`
		Created     *string `json:"Created" name:"Created"`
		Updated     *string `json:"Updated" name:"Updated"`
		Parameters  []struct {
			Name         *string `json:"name" name:"name"`
			Desc         *string `json:"desc" name:"desc"`
			DefaultValue *string `json:"defaultValue" name:"defaultValue"`
			CurrentValue *string `json:"currentValue" name:"currentValue"`
			Validity     struct {
				Type     *string `json:"Type" name:"Type"`
				DataType *string `json:"DataType" name:"DataType"`
				Value    []struct {
				} `json:"Value" name:"Value"`
			} `json:"Validity"`
		} `json:"Parameters" name:"Parameters"`
	} `json:"Data"`
}

func (r *CreateCacheParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCacheParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone         *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheParameterGroupId *string `json:"CacheParameterGroupId,omitempty" name:"CacheParameterGroupId"`
}

func (r *DeleteCacheParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCacheParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCacheParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteCacheParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCacheParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCacheParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone         *string                                `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name                  *string                                `json:"Name,omitempty" name:"Name"`
	Description           *string                                `json:"Description,omitempty" name:"Description"`
	ParamVersion          *string                                `json:"ParamVersion,omitempty" name:"ParamVersion"`
	CacheParameterGroupId *string                                `json:"CacheParameterGroupId,omitempty" name:"CacheParameterGroupId"`
	Parameters            []*ModifyCacheParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *ModifyCacheParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCacheParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCacheParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCacheParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id           *string `json:"Id" name:"Id"`
		Name         *string `json:"Name" name:"Name"`
		Description  *string `json:"Description" name:"Description"`
		ParamVersion *string `json:"ParamVersion" name:"ParamVersion"`
		Engine       *string `json:"Engine" name:"Engine"`
		Created      *string `json:"Created" name:"Created"`
		Updated      *string `json:"Updated" name:"Updated"`
		Parameters   []struct {
			Name         *string `json:"name" name:"name"`
			Desc         *string `json:"desc" name:"desc"`
			DefaultValue *string `json:"defaultValue" name:"defaultValue"`
			CurrentValue *string `json:"currentValue" name:"currentValue"`
			Validity     struct {
				Type     *string `json:"Type" name:"Type"`
				DataType *string `json:"DataType" name:"DataType"`
				Value    *string `json:"Value" name:"Value"`
				Values   []struct {
				} `json:"Values" name:"Values"`
				Min *string `json:"Min" name:"Min"`
				Max *string `json:"Max" name:"Max"`
			} `json:"Validity"`
		} `json:"Parameters" name:"Parameters"`
	} `json:"Data"`
}

func (r *ModifyCacheParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCacheParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheParameterGroupsRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone         *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name                  *string `json:"Name,omitempty" name:"Name"`
	CacheParameterGroupId *string `json:"CacheParameterGroupId,omitempty" name:"CacheParameterGroupId"`
	ParamVersion          *string `json:"ParamVersion,omitempty" name:"ParamVersion"`
	Offset                *string `json:"Offset,omitempty" name:"Offset"`
	Limit                 *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCacheParameterGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheParameterGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheParameterGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheParameterGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Engine  *string `json:"Engine" name:"Engine"`
		Created *string `json:"Created" name:"Created"`
		Updated *string `json:"Updated" name:"Updated"`
	} `json:"Data"`
}

func (r *DescribeCacheParameterGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheParameterGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone         *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheParameterGroupId *string `json:"CacheParameterGroupId,omitempty" name:"CacheParameterGroupId"`
}

func (r *DescribeCacheParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id          *string `json:"Id" name:"Id"`
		Name        *string `json:"Name" name:"Name"`
		Description *string `json:"Description" name:"Description"`
		Engine      *string `json:"Engine" name:"Engine"`
		Created     *string `json:"Created" name:"Created"`
		Updated     *string `json:"Updated" name:"Updated"`
		Parameters  []struct {
			Name         *string `json:"name" name:"name"`
			Desc         *string `json:"desc" name:"desc"`
			DefaultValue *string `json:"defaultValue" name:"defaultValue"`
			CurrentValue *string `json:"currentValue" name:"currentValue"`
			Validity     struct {
				Type     *string `json:"Type" name:"Type"`
				DataType *string `json:"DataType" name:"DataType"`
				Min      *string `json:"Min" name:"Min"`
				Max      *string `json:"Max" name:"Max"`
			} `json:"Validity"`
		} `json:"Parameters" name:"Parameters"`
	} `json:"Data"`
}

func (r *DescribeCacheParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetTimingSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	TimingSwitch  *string `json:"TimingSwitch,omitempty" name:"TimingSwitch"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Timezone      *string `json:"Timezone,omitempty" name:"Timezone"`
	Bigkey        *bool   `json:"Bigkey,omitempty" name:"Bigkey"`
}

func (r *SetTimingSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTimingSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetTimingSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetTimingSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetTimingSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTimingSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SnapshotId    *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DeleteSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name          *string `json:"Name,omitempty" name:"Name"`
	SnapshotId    *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *RenameSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenameSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RenameSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Cacheld       *string `json:"Cacheld,omitempty" name:"Cacheld"`
	Type          *string `json:"Type,omitempty" name:"Type"`
	SnapshotId    *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	BucketName    *string `json:"BucketName,omitempty" name:"BucketName"`
	ObjectName    *string `json:"ObjectName,omitempty" name:"ObjectName"`
}

func (r *RestoreSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		SnapshotId   *string `json:"SnapshotId" name:"SnapshotId"`
		Name         *string `json:"Name" name:"Name"`
		CacheId      *string `json:"CacheId" name:"CacheId"`
		Type         *string `json:"Type" name:"Type"`
		Status       *string `json:"Status" name:"Status"`
		ShardId      *string `json:"ShardId" name:"ShardId"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		ResourceSize *string `json:"ResourceSize" name:"ResourceSize"`
	} `json:"Data"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SnapshotId    *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DownloadSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DownloadSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
		Url        *string `json:"Url" name:"Url"`
	} `json:"Data"`
}

func (r *DownloadSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SnapshotId    *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	BucketName    *string `json:"BucketName,omitempty" name:"BucketName"`
	ObjectName    *string `json:"ObjectName,omitempty" name:"ObjectName"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *ExportSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ExportSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ExportSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSnapshotResponse) FromJsonString(s string) error {
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
		RegionName *string `json:"RegionName" name:"RegionName"`
		Region     *string `json:"Region" name:"Region"`
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
		Region           *string `json:"Region" name:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
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

type DescribeCacheByRoleRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Proxy         *string `json:"Proxy,omitempty" name:"Proxy"`
}

func (r *DescribeCacheByRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheByRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheByRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheByRoleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Name       *string `json:"Name" name:"Name"`
		Ip         *string `json:"Ip" name:"Ip"`
		Status     *string `json:"Status" name:"Status"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		Proxy      *string `json:"Proxy" name:"Proxy"`
		Area       *string `json:"Area" name:"Area"`
	} `json:"Data"`
}

func (r *DescribeCacheByRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheByRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticDBInstancesRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *StatisticDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StatisticDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StatisticDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Partition []struct {
			Name       *string `json:"name" name:"name"`
			Code       *string `json:"code" name:"code"`
			Total      *int    `json:"total" name:"total"`
			Statistics []struct {
				Name  *string `json:"Name" name:"Name"`
				Code  *string `json:"Code" name:"Code"`
				Total *int    `json:"Total" name:"Total"`
			} `json:"Statistics"`
		} `json:"Partition" name:"Partition"`
		Count struct {
			Name       *string `json:"name" name:"name"`
			Code       *string `json:"code" name:"code"`
			Total      *int    `json:"total" name:"total"`
			Statistics []struct {
				Name  *string `json:"Name" name:"Name"`
				Code  *string `json:"Code" name:"Code"`
				Total *int    `json:"Total" name:"Total"`
			} `json:"Statistics"`
		} `json:"Count" name:"Count"`
	} `json:"Data"`
}

func (r *StatisticDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Password      *string `json:"Password,omitempty" name:"Password"`
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
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdatePasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartCacheClusterRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *RestartCacheClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartCacheClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestartCacheClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartCacheClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestartCacheClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartCacheClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateEipRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	InsType       *string `json:"InsType,omitempty" name:"InsType"`
}

func (r *AllocateEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AllocateEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AllocateEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Master   *string `json:"Master" name:"Master"`
		Readonly *string `json:"Readonly" name:"Readonly"`
		Proxy    *string `json:"Proxy" name:"Proxy"`
	} `json:"Data"`
}

func (r *AllocateEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeallocateEipRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	InsType       *string `json:"InsType,omitempty" name:"InsType"`
}

func (r *DeallocateEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeallocateEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeallocateEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeallocateEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeallocateEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeallocateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Offset          *int    `json:"Offset,omitempty" name:"Offset"`
	Limit           *int    `json:"Limit,omitempty" name:"Limit"`
	FilterCache     *bool   `json:"FilterCache,omitempty" name:"FilterCache"`
	SearchKey       *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		List []struct {
			Id      *string `json:"id" name:"id"`
			Name    *string `json:"name" name:"name"`
			Ip      *string `json:"ip" name:"ip"`
			Mode    *int    `json:"mode" name:"mode"`
			Created *string `json:"created" name:"created"`
		} `json:"List" name:"List"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit  *int `json:"Limit" name:"Limit"`
		Total  *int `json:"Total" name:"Total"`
	} `json:"Data"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone       *string   `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SecurityGroupId     *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupRuleId []*string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
}

func (r *DeleteSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRuleRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string   `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Cidrs           []*string `json:"Cidrs,omitempty" name:"Cidrs"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *CreateSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeallocateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string   `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId         []*string `json:"CacheId,omitempty" name:"CacheId"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeallocateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeallocateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeallocateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeallocateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		CacheId *string `json:"CacheId" name:"CacheId"`
		Name    *string `json:"Name" name:"Name"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Data"`
}

func (r *DeallocateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeallocateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string   `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId         []*string `json:"CacheId,omitempty" name:"CacheId"`
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *AllocateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AllocateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AllocateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		CacheId *string `json:"CacheId" name:"CacheId"`
		Name    *string `json:"Name" name:"Name"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Data"`
}

func (r *AllocateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SearchKey       *string `json:"SearchKey,omitempty" name:"SearchKey"`
	Offset          *string `json:"Offset,omitempty" name:"Offset"`
	Limit           *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Name            *string `json:"Name" name:"Name"`
		Description     *string `json:"Description" name:"Description"`
		Rules           []struct {
			Id       *string `json:"id" name:"id"`
			Cidr     *string `json:"cidr" name:"cidr"`
			CreateAt *string `json:"createAt" name:"createAt"`
			Protocol *string `json:"protocol" name:"protocol"`
		} `json:"Rules" name:"Rules"`
		Created *string `json:"Created" name:"Created"`
		Updated *string `json:"Updated" name:"Updated"`
	} `json:"Data"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Offset        *string `json:"Offset,omitempty" name:"Offset"`
	Limit         *int    `json:"Limit,omitempty" name:"Limit"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	FilterCache   *bool   `json:"FilterCache,omitempty" name:"FilterCache"`
	SearchKey     *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		List []struct {
			SecurityGroupId *string `json:"securityGroupId" name:"securityGroupId"`
			Name            *string `json:"name" name:"name"`
			Description     *string `json:"description" name:"description"`
			Created         *string `json:"created" name:"created"`
			Updated         *string `json:"updated" name:"updated"`
		} `json:"List" name:"List"`
	} `json:"Data"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name            *string `json:"Name,omitempty" name:"Name"`
	Description     *string `json:"Description,omitempty" name:"Description"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *ModifySecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone   *string   `json:"AvailableZone,omitempty" name:"AvailableZone"`
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Message         *string `json:"Message" name:"Message"`
	} `json:"Data"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone      *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name               *string `json:"Name,omitempty" name:"Name"`
	Description        *string `json:"Description,omitempty" name:"Description"`
	SrcSecurityGroupId *string `json:"SrcSecurityGroupId,omitempty" name:"SrcSecurityGroupId"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloneSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Name            *string `json:"Name" name:"Name"`
		Description     *string `json:"Description" name:"Description"`
		Created         *string `json:"Created" name:"Created"`
		Updated         *string `json:"Updated" name:"Updated"`
	} `json:"Data"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	Name          *string `json:"Name,omitempty" name:"Name"`
	Description   *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Name            *string `json:"Name" name:"Name"`
		Description     *string `json:"Description" name:"Description"`
		Created         *string `json:"Created" name:"Created"`
		Updated         *string `json:"Updated" name:"Updated"`
	} `json:"Data"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHotKeysRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	Limit         *int    `json:"Limit,omitempty" name:"Limit"`
	Offset        *int    `json:"Offset,omitempty" name:"Offset"`
	SortKey       *string `json:"SortKey,omitempty" name:"SortKey"`
	SortDir       *string `json:"SortDir,omitempty" name:"SortDir"`
}

func (r *DescribeHotKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHotKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeHotKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHotKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Updated    *string `json:"Updated" name:"Updated"`
		TaskStatus *string `json:"TaskStatus" name:"TaskStatus"`
		Hotkeys    []struct {
			Elements          *string `json:"elements" name:"elements"`
			Key               *string `json:"key" name:"key"`
			Encoding          *string `json:"encoding" name:"encoding"`
			Type              *string `json:"type" name:"type"`
			LenLargestElement *string `json:"len_largest_element" name:"len_largest_element"`
		} `json:"Hotkeys" name:"Hotkeys"`
		Created *string `json:"Created" name:"Created"`
	} `json:"Data"`
}

func (r *DescribeHotKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHotKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeHotKeysRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *AnalyzeHotKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeHotKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AnalyzeHotKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeHotKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AnalyzeHotKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeHotKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseDirectAccessToClusterRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *CloseDirectAccessToClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseDirectAccessToClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloseDirectAccessToClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseDirectAccessToClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CloseDirectAccessToClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseDirectAccessToClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenDirectAccessToClusterRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	SubnetId      *string `json:"SubnetId,omitempty" name:"SubnetId"`
	VpcId         *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *OpenDirectAccessToClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenDirectAccessToClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "OpenDirectAccessToClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenDirectAccessToClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *OpenDirectAccessToClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenDirectAccessToClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParentBackUpsSnapshotsRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *DescribeParentBackUpsSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeParentBackUpsSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeParentBackUpsSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParentBackUpsSnapshotsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
		ResourceId *string `json:"ResourceId" name:"ResourceId"`
		BackupType *string `json:"BackupType" name:"BackupType"`
		BackupName *string `json:"BackupName" name:"BackupName"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		BackUpTime *string `json:"BackUpTime" name:"BackUpTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
		TaskStatus *string `json:"TaskStatus" name:"TaskStatus"`
	} `json:"Data"`
}

func (r *DescribeParentBackUpsSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeParentBackUpsSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackUpsSnapshotsDetailRequest struct {
	*ksyunhttp.BaseRequest
	CacheId    *string `json:"CacheId,omitempty" name:"CacheId"`
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DescribeBackUpsSnapshotsDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackUpsSnapshotsDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBackUpsSnapshotsDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackUpsSnapshotsDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		SnapshotId   *string `json:"SnapshotId" name:"SnapshotId"`
		ResourceId   *string `json:"ResourceId" name:"ResourceId"`
		ResourceSize *int    `json:"ResourceSize" name:"ResourceSize"`
		BackupName   *string `json:"BackupName" name:"BackupName"`
		BackupType   *string `json:"BackupType" name:"BackupType"`
		TaskStatus   *string `json:"TaskStatus" name:"TaskStatus"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime   *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
}

func (r *DescribeBackUpsSnapshotsDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackUpsSnapshotsDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLevelSnapshotsRequest struct {
	*ksyunhttp.BaseRequest
	CacheId    *string `json:"CacheId,omitempty" name:"CacheId"`
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DeleteLevelSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLevelSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteLevelSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLevelSnapshotsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteLevelSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLevelSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadLevelSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	CacheId    *string `json:"CacheId,omitempty" name:"CacheId"`
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DownloadLevelSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadLevelSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DownloadLevelSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadLevelSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		DownloadUrl []struct {
		} `json:"DownloadUrl" name:"DownloadUrl"`
	} `json:"Data"`
}

func (r *DownloadLevelSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadLevelSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBigKeysRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *DescribeBigKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBigKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBigKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBigKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id         *string `json:"Id" name:"Id"`
		ResourceId *string `json:"ResourceId" name:"ResourceId"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
		TaskType   *string `json:"TaskType" name:"TaskType"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		TaskStatus *string `json:"TaskStatus" name:"TaskStatus"`
	} `json:"Data"`
}

func (r *DescribeBigKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBigKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBigKeysAnalyseResultRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
	TaskId  *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteBigKeysAnalyseResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBigKeysAnalyseResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteBigKeysAnalyseResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBigKeysAnalyseResultResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteBigKeysAnalyseResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBigKeysAnalyseResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeBigKeysRequest struct {
	*ksyunhttp.BaseRequest
	TaskId    *string `json:"TaskId,omitempty" name:"TaskId"`
	QueryType *int    `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *AnalyzeBigKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeBigKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AnalyzeBigKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeBigKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		STRING struct {
			KeyType *string `json:"keyType" name:"keyType"`
			Percent *string `json:"percent" name:"percent"`
		} `json:"STRING" name:"STRING"`
		LIST struct {
			KeyType *string `json:"keyType" name:"keyType"`
			Percent *string `json:"percent" name:"percent"`
		} `json:"LIST" name:"LIST"`
		SET struct {
			KeyType *string `json:"keyType" name:"keyType"`
			Percent *string `json:"percent" name:"percent"`
		} `json:"SET" name:"SET"`
		HASH struct {
			KeyType *string `json:"keyType" name:"keyType"`
			Percent *string `json:"percent" name:"percent"`
		} `json:"HASH" name:"HASH"`
		SORT struct {
			KeyType *string `json:"keyType" name:"keyType"`
			Percent *string `json:"percent" name:"percent"`
		} `json:"SORT" name:"SORT"`
	} `json:"Data"`
}

func (r *AnalyzeBigKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeBigKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCreateSnapshotStatusRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeCreateSnapshotStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCreateSnapshotStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCreateSnapshotStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCreateSnapshotStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		BackupResource struct {
			ResourceId     *string `json:"resourceId" name:"resourceId"`
			BackupType     *int    `json:"backup_type" name:"backup_type"`
			BackupTimezone *string `json:"backup_timezone" name:"backup_timezone"`
			ResourceMode   *string `json:"resource_mode" name:"resource_mode"`
			TaskStatus     *string `json:"task_status" name:"task_status"`
			Metadata       struct {
				Bigkey *int `json:"Bigkey" name:"Bigkey"`
			} `json:"Metadata"`
			ResourceType *string `json:"resource_type" name:"resource_type"`
		} `json:"BackupResource" name:"BackupResource"`
	} `json:"Data"`
}

func (r *DescribeCreateSnapshotStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCreateSnapshotStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDailyAnalyzeSwitchStateRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *GetDailyAnalyzeSwitchStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDailyAnalyzeSwitchStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDailyAnalyzeSwitchStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDailyAnalyzeSwitchStateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Slowlog    *bool `json:"Slowlog" name:"Slowlog"`
		Servicelog *bool `json:"Servicelog" name:"Servicelog"`
	} `json:"Data"`
}

func (r *GetDailyAnalyzeSwitchStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDailyAnalyzeSwitchStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeDailyRequest struct {
	*ksyunhttp.BaseRequest
	CacheId        *string `json:"CacheId,omitempty" name:"CacheId"`
	StartQueryTime *string `json:"StartQueryTime,omitempty" name:"StartQueryTime"`
	EndQueryTime   *string `json:"EndQueryTime,omitempty" name:"EndQueryTime"`
	PageNum        *string `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *string `json:"PageSize,omitempty" name:"PageSize"`
	ShardId        *string `json:"ShardId,omitempty" name:"ShardId"`
	OperationType  *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *AnalyzeDailyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeDailyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AnalyzeDailyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeDailyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *AnalyzeDailyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeDailyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeSlowDailyRequest struct {
	*ksyunhttp.BaseRequest
	CacheId        *string `json:"CacheId,omitempty" name:"CacheId"`
	StartQueryTime *string `json:"StartQueryTime,omitempty" name:"StartQueryTime"`
	EndQueryTime   *string `json:"EndQueryTime,omitempty" name:"EndQueryTime"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
	ShardId        *string `json:"ShardId,omitempty" name:"ShardId"`
	IsProxy        *bool   `json:"IsProxy,omitempty" name:"IsProxy"`
	OperationType  *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *AnalyzeSlowDailyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeSlowDailyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AnalyzeSlowDailyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeSlowDailyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *AnalyzeSlowDailyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeSlowDailyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeDailySwitchRequest struct {
	*ksyunhttp.BaseRequest
	CacheId    *string `json:"CacheId,omitempty" name:"CacheId"`
	ServiceLog *bool   `json:"ServiceLog,omitempty" name:"ServiceLog"`
	SlowLog    *bool   `json:"SlowLog,omitempty" name:"SlowLog"`
}

func (r *AnalyzeDailySwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeDailySwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AnalyzeDailySwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeDailySwitchResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *AnalyzeDailySwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyzeDailySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreByTimePointSwitchRequest struct {
	*ksyunhttp.BaseRequest
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	RestoreSwitch *string `json:"RestoreSwitch,omitempty" name:"RestoreSwitch"`
}

func (r *RestoreByTimePointSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreByTimePointSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreByTimePointSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreByTimePointSwitchResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *RestoreByTimePointSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreByTimePointSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRestoreTimePointsRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
}

func (r *DescribeRestoreTimePointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRestoreTimePointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRestoreTimePointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRestoreTimePointsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeRestoreTimePointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRestoreTimePointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBigHotKeysRequest struct {
	*ksyunhttp.BaseRequest
	CacheId        *string `json:"CacheId,omitempty" name:"CacheId"`
	NodeId         *string `json:"NodeId,omitempty" name:"NodeId"`
	QueryType      *string `json:"QueryType,omitempty" name:"QueryType"`
	KeyType        *int    `json:"KeyType,omitempty" name:"KeyType"`
	KeyName        *string `json:"KeyName,omitempty" name:"KeyName"`
	StartQueryTime *string `json:"StartQueryTime,omitempty" name:"StartQueryTime"`
	EndQueryTime   *string `json:"EndQueryTime,omitempty" name:"EndQueryTime"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNum        *int    `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeBigHotKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBigHotKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBigHotKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBigHotKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		KeyInfos []struct {
			KeyName        *string `json:"KeyName" name:"KeyName"`
			DB             *int    `json:"DB" name:"DB"`
			DataType       *string `json:"DataType" name:"DataType"`
			NodeId         *string `json:"NodeId" name:"NodeId"`
			StatisticValue *int    `json:"StatisticValue" name:"StatisticValue"`
			CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		} `json:"KeyInfos" name:"KeyInfos"`
		TotalCount  *int `json:"TotalCount" name:"TotalCount"`
		TotalPage   *int `json:"TotalPage" name:"TotalPage"`
		CurrentPage *int `json:"CurrentPage" name:"CurrentPage"`
		PageSize    *int `json:"PageSize" name:"PageSize"`
	} `json:"Data"`
}

func (r *DescribeBigHotKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBigHotKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginsRequest struct {
	*ksyunhttp.BaseRequest
	CacheId     *string `json:"CacheId,omitempty" name:"CacheId"`
	Installed   *bool   `json:"Installed,omitempty" name:"Installed"`
	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset      *int    `json:"Offset,omitempty" name:"Offset"`
	Limit       *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Plugins []struct {
			PluginName      *string   `json:"PluginName" name:"PluginName"`
			EnableUninstall *bool     `json:"EnableUninstall" name:"EnableUninstall"`
			UpgradeVersions []*string `json:"UpgradeVersions" name:"UpgradeVersions"`
			Description     *string   `json:"Description" name:"Description"`
		} `json:"Plugins" name:"Plugins"`
		Total *int `json:"Total" name:"Total"`
	} `json:"Data"`
}

func (r *DescribePluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallPluginsRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string                  `json:"CacheId,omitempty" name:"CacheId"`
	Plugins []*InstallPluginsPlugins `json:"Plugins,omitempty" name:"Plugins"`
}

func (r *InstallPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "InstallPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InstallPluginsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *InstallPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallPluginsRequest struct {
	*ksyunhttp.BaseRequest
	CacheId     *string   `json:"CacheId,omitempty" name:"CacheId"`
	PluginNames []*string `json:"PluginNames,omitempty" name:"PluginNames"`
}

func (r *UninstallPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UninstallPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UninstallPluginsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *UninstallPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
