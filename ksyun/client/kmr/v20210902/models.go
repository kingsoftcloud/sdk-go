package v20210902

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type LaunchClusterInstanceGroups struct {
	InstanceGroupType     *string `json:"InstanceGroupType,omitempty" name:"InstanceGroupType"`
	AvailabilityZone      *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	InstanceType          *string `json:"InstanceType,omitempty" name:"InstanceType"`
	ResourceType          *string `json:"ResourceType,omitempty" name:"ResourceType"`
	InstanceCount         *int    `json:"InstanceCount,omitempty" name:"InstanceCount"`
	RaidType              *string `json:"RaidType,omitempty" name:"RaidType"`
	VolumeType            *string `json:"VolumeType,omitempty" name:"VolumeType"`
	VolumeSize            *int    `json:"VolumeSize,omitempty" name:"VolumeSize"`
	VolumeCount           *int    `json:"VolumeCount,omitempty" name:"VolumeCount"`
	SystemDiskType        *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	SystemDiskSize        *int    `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	VpcSubnetId           *string `json:"VpcSubnetId,omitempty" name:"VpcSubnetId"`
	AvailabilityZoneIndex *int    `json:"AvailabilityZoneIndex,omitempty" name:"AvailabilityZoneIndex"`
	InstanceGroupIndex    *int    `json:"InstanceGroupIndex,omitempty" name:"InstanceGroupIndex"`
}

type ScaleInInstanceGroupsInstanceGroups struct {
	Id        *string `json:"Id,omitempty" name:"Id"`
	Instances []struct {
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	} `json:"instances,omitempty" name:"instances"`
}

type ScaleInInstanceGroupsinstances struct {
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ScaleOutInstanceGroupsInstanceGroups struct {
	InstanceGroupType     *string `json:"InstanceGroupType,omitempty" name:"InstanceGroupType"`
	InstanceType          *string `json:"InstanceType,omitempty" name:"InstanceType"`
	ResourceType          *string `json:"ResourceType,omitempty" name:"ResourceType"`
	InstanceCount         *int    `json:"InstanceCount,omitempty" name:"InstanceCount"`
	RaidType              *string `json:"RaidType,omitempty" name:"RaidType"`
	VolumeType            *string `json:"VolumeType,omitempty" name:"VolumeType"`
	VolumeSize            *int    `json:"VolumeSize,omitempty" name:"VolumeSize"`
	VolumeCount           *int    `json:"VolumeCount,omitempty" name:"VolumeCount"`
	SystemDiskType        *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	SystemDiskSize        *int    `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	AvailabilityZone      *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	AvailabilityZoneIndex *int    `json:"AvailabilityZoneIndex,omitempty" name:"AvailabilityZoneIndex"`
	InstanceGroupIndex    *int    `json:"InstanceGroupIndex,omitempty" name:"InstanceGroupIndex"`
}

type BindTagsTags struct {
	TagKey   *string `json:"TagKey,omitempty" name:"TagKey"`
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	TagId    *int    `json:"TagId,omitempty" name:"TagId"`
}

type DescribeClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId      *string `json:"ClusterId" name:"ClusterId"`
	ClusterName    *string `json:"ClusterName" name:"ClusterName"`
	ClusterType    *string `json:"ClusterType" name:"ClusterType"`
	MainVersion    *string `json:"MainVersion" name:"MainVersion"`
	InstanceGroups []struct {
		Id                *string `json:"Id" name:"Id"`
		InstanceGroupType *string `json:"InstanceGroupType" name:"InstanceGroupType"`
		ResourceType      *string `json:"ResourceType" name:"ResourceType"`
		InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		VolumeSize        *int    `json:"VolumeSize" name:"VolumeSize"`
		VolumeType        *string `json:"VolumeType" name:"VolumeType"`
		InstanceIds       []struct {
		} `json:"InstanceIds" name:"InstanceIds"`
	} `json:"InstanceGroups"`
	EnableEip      *bool   `json:"EnableEip" name:"EnableEip"`
	Region         *string `json:"Region" name:"Region"`
	VpcDomainId    *string `json:"VpcDomainId" name:"VpcDomainId"`
	ChargeType     *string `json:"ChargeType" name:"ChargeType"`
	ClusterStatus  *string `json:"ClusterStatus" name:"ClusterStatus"`
	CreateTime     *string `json:"CreateTime" name:"CreateTime"`
	UpdateTime     *string `json:"UpdateTime" name:"UpdateTime"`
	ServingMinutes *int    `json:"ServingMinutes" name:"ServingMinutes"`
	RequestId      *string `json:"RequestId" name:"RequestId"`
	StatusCode     *int    `json:"StatusCode" name:"StatusCode"`
	Tags           []struct {
		TagId    *int    `json:"TagId" name:"TagId"`
		TagKey   *string `json:"TagKey" name:"TagKey"`
		TagValue *string `json:"TagValue" name:"TagValue"`
	} `json:"Tags"`
}

func (r *DescribeClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LaunchClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName     *string                        `json:"ClusterName,omitempty" name:"ClusterName"`
	Distribution    *string                        `json:"Distribution,omitempty" name:"Distribution"`
	MainVersion     *string                        `json:"MainVersion,omitempty" name:"MainVersion"`
	ChargeType      *string                        `json:"ChargeType,omitempty" name:"ChargeType"`
	Services        []*string                      `json:"Services,omitempty" name:"Services"`
	ProjectId       *int                           `json:"ProjectId,omitempty" name:"ProjectId"`
	VpcDomainId     *string                        `json:"VpcDomainId,omitempty" name:"VpcDomainId"`
	SecurityGroupId *string                        `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstanceGroups  []*LaunchClusterInstanceGroups `json:"InstanceGroups,omitempty" name:"InstanceGroups"`
}

func (r *LaunchClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LaunchClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "LaunchClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type LaunchClusterResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *LaunchClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LaunchClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleInInstanceGroupsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId              *string                                `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceGroups         []*ScaleInInstanceGroupsInstanceGroups `json:"InstanceGroups,omitempty" name:"InstanceGroups"`
	GracefulScaleIn        *bool                                  `json:"GracefulScaleIn,omitempty" name:"GracefulScaleIn"`
	GracefulScaleInTimeout *int                                   `json:"GracefulScaleInTimeout,omitempty" name:"GracefulScaleInTimeout"`
}

func (r *ScaleInInstanceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleInInstanceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ScaleInInstanceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScaleInInstanceGroupsResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ScaleInInstanceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleInInstanceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleOutInstanceGroupsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string                                 `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceGroups []*ScaleOutInstanceGroupsInstanceGroups `json:"InstanceGroups,omitempty" name:"InstanceGroups"`
	ProjectId      *int                                    `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ScaleOutInstanceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleOutInstanceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ScaleOutInstanceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScaleOutInstanceGroupsResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ScaleOutInstanceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleOutInstanceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInfoRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	GaeaInfo  struct {
		Ks3Path *string `json:"Ks3Path" name:"Ks3Path"`
		Ks3Type *string `json:"Ks3Type" name:"Ks3Type"`
	} `json:"GaeaInfo"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *DescribeClusterInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListServiceStatusRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListServiceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListServiceStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Services  []struct {
		ClusterId     *string `json:"ClusterId" name:"ClusterId"`
		ServiceId     *string `json:"ServiceId" name:"ServiceId"`
		ServiceName   *string `json:"ServiceName" name:"ServiceName"`
		Status        *int    `json:"Status" name:"Status"`
		Steady        *bool   `json:"Steady" name:"Steady"`
		Version       *string `json:"Version" name:"Version"`
		AbnormalCount *int    `json:"AbnormalCount" name:"AbnormalCount"`
		ComponentInfo []struct {
		} `json:"ComponentInfo" name:"ComponentInfo"`
		LastStartTime *string `json:"LastStartTime" name:"LastStartTime"`
		WebUrl        *string `json:"WebUrl" name:"WebUrl"`
		Port          []struct {
		} `json:"Port" name:"Port"`
		WebInfo []struct {
		} `json:"WebInfo" name:"WebInfo"`
	} `json:"Services"`
	EnableEip      *bool `json:"EnableEip" name:"EnableEip"`
	IsConfComplete *bool `json:"IsConfComplete" name:"IsConfComplete"`
	StatusCode     *int  `json:"StatusCode" name:"StatusCode"`
}

func (r *ListServiceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClustersRequest struct {
	*ksyunhttp.BaseRequest
	Marker *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *ListClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListClustersResponse struct {
	*ksyunhttp.BaseResponse
	Clusters []struct {
		ClusterId      *string `json:"ClusterId" name:"ClusterId"`
		ClusterName    *string `json:"ClusterName" name:"ClusterName"`
		MainVersion    *string `json:"MainVersion" name:"MainVersion"`
		InstanceGroups []struct {
			Id                *string `json:"Id" name:"Id"`
			InstanceGroupType *string `json:"InstanceGroupType" name:"InstanceGroupType"`
			ResourceType      *string `json:"ResourceType" name:"ResourceType"`
			InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		} `json:"InstanceGroups" name:"InstanceGroups"`
		EnableEip        *bool   `json:"EnableEip" name:"EnableEip"`
		Region           *string `json:"Region" name:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		VpcDomainId      *string `json:"VpcDomainId" name:"VpcDomainId"`
		VpcSubnetId      *string `json:"VpcSubnetId" name:"VpcSubnetId"`
		ChargeType       *string `json:"ChargeType" name:"ChargeType"`
		ClusterStatus    *string `json:"ClusterStatus" name:"ClusterStatus"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime       *string `json:"UpdateTime" name:"UpdateTime"`
		ServingMinutes   *int    `json:"ServingMinutes" name:"ServingMinutes"`
	} `json:"Clusters"`
	Total      *int    `json:"Total" name:"Total"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ListClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterVersionsRequest struct {
	*ksyunhttp.BaseRequest
	DistributionVersion *string `json:"DistributionVersion,omitempty" name:"DistributionVersion"`
}

func (r *ListClusterVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListClusterVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterVersionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Versions  []struct {
		Id       *string `json:"Id" name:"Id"`
		Version  *string `json:"Version" name:"Version"`
		Services []struct {
			Name               *string `json:"Name" name:"Name"`
			Version            *string `json:"Version" name:"Version"`
			Essential          *bool   `json:"Essential" name:"Essential"`
			IsSystem           *bool   `json:"IsSystem" name:"IsSystem"`
			DependencyServices *string `json:"DependencyServices" name:"DependencyServices"`
		} `json:"Services" name:"Services"`
	} `json:"Versions"`
	InstanceGroupRequirements []struct {
		Name      *string `json:"Name" name:"Name"`
		Essential *bool   `json:"Essential" name:"Essential"`
		IsPrimary *bool   `json:"IsPrimary" name:"IsPrimary"`
		Isomerism *bool   `json:"Isomerism" name:"Isomerism"`
		Max       *int    `json:"Max" name:"Max"`
		Min       *int    `json:"Min" name:"Min"`
		Step      *int    `json:"Step" name:"Step"`
	} `json:"InstanceGroupRequirements"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *ListClusterVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	Services        *string `json:"Services" name:"Services"`
	ComponentStatus []struct {
		ClusterId             *string `json:"ClusterId" name:"ClusterId"`
		ComponentId           *string `json:"ComponentId" name:"ComponentId"`
		ComponentName         *string `json:"ComponentName" name:"ComponentName"`
		Role                  *string `json:"Role" name:"Role"`
		InstanceId            *string `json:"InstanceId" name:"InstanceId"`
		InstanceName          *string `json:"InstanceName" name:"InstanceName"`
		InternalIp            *string `json:"InternalIp" name:"InternalIp"`
		InstanceGroupType     *string `json:"InstanceGroupType" name:"InstanceGroupType"`
		InstanceGroupIndex    *int    `json:"InstanceGroupIndex" name:"InstanceGroupIndex"`
		AvailabilityZoneIndex *int    `json:"AvailabilityZoneIndex" name:"AvailabilityZoneIndex"`
		Installed             *bool   `json:"Installed" name:"Installed"`
		State                 *int    `json:"State" name:"State"`
		TargetState           *int    `json:"TargetState" name:"TargetState"`
		Steady                *bool   `json:"Steady" name:"Steady"`
		LastStartTime         *string `json:"LastStartTime" name:"LastStartTime"`
		ErrorType             *int    `json:"ErrorType" name:"ErrorType"`
		ErrorMsg              *string `json:"ErrorMsg" name:"ErrorMsg"`
		RestartRequired       *bool   `json:"RestartRequired" name:"RestartRequired"`
		ExtraInfo             []struct {
		} `json:"ExtraInfo" name:"ExtraInfo"`
	} `json:"ComponentStatus"`
	Total      *int `json:"Total" name:"Total"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConfigurationsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *ListConfigurationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListConfigurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListConfigurationsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	ClusterId    *string `json:"ClusterId" name:"ClusterId"`
	ServiceName  *string `json:"ServiceName" name:"ServiceName"`
	ConfigGroups []struct {
		Id              *string `json:"Id" name:"Id"`
		InstanceGroupId *string `json:"InstanceGroupId" name:"InstanceGroupId"`
		ConfigTags      []struct {
			Tag            *string `json:"Tag" name:"Tag"`
			Configurations []struct {
				Key         *string `json:"Key" name:"Key"`
				Value       *string `json:"Value" name:"Value"`
				Custom      *bool   `json:"Custom" name:"Custom"`
				Description *string `json:"Description" name:"Description"`
				RelatedKey  *string `json:"RelatedKey" name:"RelatedKey"`
				UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
			} `json:"Configurations"`
		} `json:"ConfigTags" name:"ConfigTags"`
	} `json:"ConfigGroups"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *ListConfigurationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConfigurationHistoryRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	Marker      *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *ListConfigurationHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigurationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListConfigurationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListConfigurationHistoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ClusterId   *string `json:"ClusterId" name:"ClusterId"`
	ServiceName *string `json:"ServiceName" name:"ServiceName"`
	Total       *int    `json:"Total" name:"Total"`
	Data        []struct {
		Key                   *string `json:"Key" name:"Key"`
		InstanceGroupType     *string `json:"InstanceGroupType" name:"InstanceGroupType"`
		InstanceGroupIndex    *int    `json:"InstanceGroupIndex" name:"InstanceGroupIndex"`
		AvailabilityZoneIndex *int    `json:"AvailabilityZoneIndex" name:"AvailabilityZoneIndex"`
		UpdateTime            *string `json:"UpdateTime" name:"UpdateTime"`
		SourceValue           *string `json:"SourceValue" name:"SourceValue"`
		TargetValue           *string `json:"TargetValue" name:"TargetValue"`
		Activated             *bool   `json:"Activated" name:"Activated"`
	} `json:"Data"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *ListConfigurationHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigurationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTagValuesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListTagValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListTagValuesResponse struct {
	*ksyunhttp.BaseResponse
	TagValues []struct {
		Id         *int    `json:"Id" name:"Id"`
		Key        *string `json:"Key" name:"Key"`
		Value      *string `json:"Value" name:"Value"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		BindNum    *int    `json:"BindNum" name:"BindNum"`
	} `json:"TagValues"`
	Page       *int    `json:"Page" name:"Page"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
	Total      *int    `json:"Total" name:"Total"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ListTagValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTagKeysRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListTagKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListTagKeysResponse struct {
	*ksyunhttp.BaseResponse
	TagKeys    []*string `json:"TagKeys" name:"TagKeys"`
	Page       *int      `json:"Page" name:"Page"`
	PageSize   *int      `json:"PageSize" name:"PageSize"`
	Total      *int      `json:"Total" name:"Total"`
	RequestId  *string   `json:"RequestId" name:"RequestId"`
	StatusCode *int      `json:"StatusCode" name:"StatusCode"`
}

func (r *ListTagKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindTagsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string         `json:"ClusterId,omitempty" name:"ClusterId"`
	Tags      []*BindTagsTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *BindTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BindTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindTagsResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *BindTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *StartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceIds    []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	Interval       *int      `json:"Interval,omitempty" name:"Interval"`
	RollingRestart *bool     `json:"RollingRestart,omitempty" name:"RollingRestart"`
	ForceReboot    *bool     `json:"ForceReboot,omitempty" name:"ForceReboot"`
}

func (r *RestartInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartInstancesResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *RestartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StopInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *StopInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	Id             *string `json:"Id" name:"Id"`
	Name           *string `json:"Name" name:"Name"`
	StrategyModule *string `json:"StrategyModule" name:"StrategyModule"`
	Strategies     []struct {
		Id                 *string `json:"Id" name:"Id"`
		Name               *string `json:"Name" name:"Name"`
		ClusterId          *string `json:"ClusterId" name:"ClusterId"`
		Type               *string `json:"Type" name:"Type"`
		Enabled            *bool   `json:"Enabled" name:"Enabled"`
		Priority           *int    `json:"Priority" name:"Priority"`
		MaxInstance        *int    `json:"MaxInstance" name:"MaxInstance"`
		MinInstance        *int    `json:"MinInstance" name:"MinInstance"`
		ChargeType         *string `json:"ChargeType" name:"ChargeType"`
		PeriodType         *string `json:"PeriodType" name:"PeriodType"`
		PeriodParam        *string `json:"PeriodParam" name:"PeriodParam"`
		PeriodTime         *string `json:"PeriodTime" name:"PeriodTime"`
		Step               *int    `json:"Step" name:"Step"`
		NewTaskGroup       *bool   `json:"NewTaskGroup" name:"NewTaskGroup"`
		InstanceGroupId    *string `json:"InstanceGroupId" name:"InstanceGroupId"`
		InstanceGroupAlias *string `json:"InstanceGroupAlias" name:"InstanceGroupAlias"`
		InstanceGroupTypes struct {
			VpcSubnetId     *string `json:"VpcSubnetId" name:"VpcSubnetId"`
			SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			ResourceTypes   *string `json:"ResourceTypes" name:"ResourceTypes"`
		} `json:"InstanceGroupTypes" name:"InstanceGroupTypes"`
		ResourceBackUp    *bool   `json:"ResourceBackUp" name:"ResourceBackUp"`
		AutoRestore       *string `json:"AutoRestore" name:"AutoRestore"`
		ActivateTime      *string `json:"ActivateTime" name:"ActivateTime"`
		ExpireTime        *string `json:"ExpireTime" name:"ExpireTime"`
		InstanceGroupInfo struct {
			SystemDiskType *string `json:"SystemDiskType" name:"SystemDiskType"`
			SystemDiskSize *int    `json:"SystemDiskSize" name:"SystemDiskSize"`
		} `json:"InstanceGroupInfo" name:"InstanceGroupInfo"`
		GracefulScaleIn        *bool   `json:"GracefulScaleIn" name:"GracefulScaleIn"`
		GracefulScaleInTimeout *int    `json:"GracefulScaleInTimeout" name:"GracefulScaleInTimeout"`
		InstancePassword       *string `json:"InstancePassword" name:"InstancePassword"`
		LabelId                *string `json:"LabelId" name:"LabelId"`
		CreatedAt              *string `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt              *string `json:"UpdatedAt" name:"UpdatedAt"`
	} `json:"Strategies"`
	ConflictStrategies  *string `json:"ConflictStrategies" name:"ConflictStrategies"`
	LoadBasedStrategies []struct {
		Id                 *string `json:"Id" name:"Id"`
		ClusterId          *string `json:"ClusterId" name:"ClusterId"`
		ActivateTime       *string `json:"ActivateTime" name:"ActivateTime"`
		ChargeType         *string `json:"ChargeType" name:"ChargeType"`
		Enabled            *bool   `json:"Enabled" name:"Enabled"`
		ExpireTime         *string `json:"ExpireTime" name:"ExpireTime"`
		NewTaskGroup       *bool   `json:"NewTaskGroup" name:"NewTaskGroup"`
		InstanceGroupId    *string `json:"InstanceGroupId" name:"InstanceGroupId"`
		InstanceGroupAlias *string `json:"InstanceGroupAlias" name:"InstanceGroupAlias"`
		InstanceGroupInfo  struct {
			SystemDiskType *string `json:"SystemDiskType" name:"SystemDiskType"`
			SystemDiskSize *int    `json:"SystemDiskSize" name:"SystemDiskSize"`
		} `json:"InstanceGroupInfo" name:"InstanceGroupInfo"`
		InstanceGroupTypes struct {
			VpcSubnetId     *string `json:"VpcSubnetId" name:"VpcSubnetId"`
			SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			ResourceTypes   []struct {
				TypeCode       *string `json:"TypeCode" name:"TypeCode"`
				DiskType       *string `json:"DiskType" name:"DiskType"`
				DiskSize       *int    `json:"DiskSize" name:"DiskSize"`
				VolumeCount    *int    `json:"VolumeCount" name:"VolumeCount"`
				SystemDiskType *string `json:"SystemDiskType" name:"SystemDiskType"`
				SystemDiskSize *int    `json:"SystemDiskSize" name:"SystemDiskSize"`
			} `json:"ResourceTypes"`
		} `json:"InstanceGroupTypes" name:"InstanceGroupTypes"`
		MaxInstance    *int    `json:"MaxInstance" name:"MaxInstance"`
		MinInstance    *int    `json:"MinInstance" name:"MinInstance"`
		Name           *string `json:"Name" name:"Name"`
		Step           *int    `json:"Step" name:"Step"`
		Type           *string `json:"Type" name:"Type"`
		StatisticsRule struct {
			LoadMetricId   *string `json:"LoadMetricId" name:"LoadMetricId"`
			Period         *int    `json:"Period" name:"Period"`
			CompareType    *string `json:"CompareType" name:"CompareType"`
			Threshold      *int    `json:"Threshold" name:"Threshold"`
			RepeatCount    *int    `json:"RepeatCount" name:"RepeatCount"`
			CoolDownPeriod *int    `json:"CoolDownPeriod" name:"CoolDownPeriod"`
		} `json:"StatisticsRule" name:"StatisticsRule"`
		TriggerTimes           *int    `json:"TriggerTimes" name:"TriggerTimes"`
		GracefulScaleIn        *bool   `json:"GracefulScaleIn" name:"GracefulScaleIn"`
		GracefulScaleInTimeout *int    `json:"GracefulScaleInTimeout" name:"GracefulScaleInTimeout"`
		InstancePassword       *string `json:"InstancePassword" name:"InstancePassword"`
		CreatedAt              *string `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt              *string `json:"UpdatedAt" name:"UpdatedAt"`
	} `json:"LoadBasedStrategies"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *ListScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	Succeed    *bool   `json:"Succeed" name:"Succeed"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ModifyScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string `json:"ClusterId,omitempty" name:"ClusterId"`
	StrategyId     *string `json:"StrategyId,omitempty" name:"StrategyId"`
	StrategyModule *string `json:"StrategyModule,omitempty" name:"StrategyModule"`
}

func (r *DeleteScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	Succeed    *bool   `json:"Succeed" name:"Succeed"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *DeleteScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListScaleHistoryRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListScaleHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScaleHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListScaleHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListScaleHistoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	ClusterId      *string `json:"ClusterId" name:"ClusterId"`
	ScaleHistories []struct {
		StartTime                *string `json:"StartTime" name:"StartTime"`
		ScaleHistoryId           *string `json:"ScaleHistoryId" name:"ScaleHistoryId"`
		StrategyName             *string `json:"StrategyName" name:"StrategyName"`
		Action                   *string `json:"Action" name:"Action"`
		StrategyModule           *string `json:"StrategyModule" name:"StrategyModule"`
		InstanceGroupId          *string `json:"InstanceGroupId" name:"InstanceGroupId"`
		InstanceCount            *int    `json:"InstanceCount" name:"InstanceCount"`
		Status                   *string `json:"Status" name:"Status"`
		Result                   *string `json:"Result" name:"Result"`
		PreviousCount            *int    `json:"PreviousCount" name:"PreviousCount"`
		FinalCount               *int    `json:"FinalCount" name:"FinalCount"`
		AutoCreatedInstanceGroup *bool   `json:"AutoCreatedInstanceGroup" name:"AutoCreatedInstanceGroup"`
		LabelId                  *string `json:"LabelId" name:"LabelId"`
		ResourceType             *string `json:"ResourceType" name:"ResourceType"`
	} `json:"ScaleHistories"`
	Total      *int `json:"Total" name:"Total"`
	StatusCode *int `json:"StatusCode" name:"StatusCode"`
}

func (r *ListScaleHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScaleHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLoadBasedScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AddLoadBasedScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLoadBasedScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddLoadBasedScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddLoadBasedScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	Succeed    *bool   `json:"Succeed" name:"Succeed"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *AddLoadBasedScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLoadBasedScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBasedScaleStrategyRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyLoadBasedScaleStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBasedScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyLoadBasedScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBasedScaleStrategyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	Succeed    *bool   `json:"Succeed" name:"Succeed"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ModifyLoadBasedScaleStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBasedScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
