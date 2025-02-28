package v20201215

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type LaunchClusterInstanceGroups struct {
	InstanceGroupType  *string `json:"InstanceGroupType,omitempty" name:"InstanceGroupType"`
	InstanceType       *string `json:"InstanceType,omitempty" name:"InstanceType"`
	ResourceType       *string `json:"ResourceType,omitempty" name:"ResourceType"`
	InstanceCount      *int    `json:"InstanceCount,omitempty" name:"InstanceCount"`
	MultiInstanceCount *int    `json:"MultiInstanceCount,omitempty" name:"MultiInstanceCount"`
	RaidType           *string `json:"RaidType,omitempty" name:"RaidType"`
	VolumeType         *string `json:"VolumeType,omitempty" name:"VolumeType"`
	VolumeSize         *int    `json:"VolumeSize,omitempty" name:"VolumeSize"`
	VolumeCount        *int    `json:"VolumeCount,omitempty" name:"VolumeCount"`
	SystemDiskType     *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	SystemDiskSize     *int    `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
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
		Id                *string `json:"Id"`
		InstanceGroupType *string `json:"InstanceGroupType"`
		ResourceType      *string `json:"ResourceType"`
		InstanceType      *string `json:"InstanceType"`
		VolumeSize        *int    `json:"VolumeSize"`
		VolumeType        *string `json:"VolumeType"`
	} `json:"InstanceGroups"`
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
	ProxyPort        *int    `json:"ProxyPort" name:"ProxyPort"`
	Tags             []struct {
		TagKey   *string `json:"TagKey"`
		TagValue *string `json:"TagValue"`
	} `json:"Tags"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *DescribeClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResponse) FromJsonString(s string) error {
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
		ClusterId      *string `json:"ClusterId"`
		ClusterName    *string `json:"ClusterName"`
		MainVersion    *string `json:"MainVersion"`
		InstanceGroups []struct {
			Id                *string `json:"Id"`
			InstanceGroupType *string `json:"InstanceGroupType"`
			ResourceType      *string `json:"ResourceType"`
			InstanceType      *string `json:"InstanceType"`
		} `json:"InstanceGroups"`
		EnableEip        *bool   `json:"EnableEip"`
		Region           *string `json:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		VpcDomainId      *string `json:"VpcDomainId"`
		VpcSubnetId      *string `json:"VpcSubnetId"`
		ChargeType       *string `json:"ChargeType"`
		ClusterStatus    *string `json:"ClusterStatus"`
		CreateTime       *string `json:"CreateTime"`
		UpdateTime       *string `json:"UpdateTime"`
		ServingMinutes   *int    `json:"ServingMinutes"`
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

type ModifyClusterNameRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LaunchClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName      *string                        `json:"ClusterName,omitempty" name:"ClusterName"`
	MainVersion      *string                        `json:"MainVersion,omitempty" name:"MainVersion"`
	Scenario         *string                        `json:"Scenario,omitempty" name:"Scenario"`
	AvailabilityZone *string                        `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ChargeType       *string                        `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime     *int                           `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	ProjectId        *int                           `json:"ProjectId,omitempty" name:"ProjectId"`
	VpcDomainId      *string                        `json:"VpcDomainId,omitempty" name:"VpcDomainId"`
	VpcSubnetId      *string                        `json:"VpcSubnetId,omitempty" name:"VpcSubnetId"`
	VpcEpcSubnetId   *string                        `json:"VpcEpcSubnetId,omitempty" name:"VpcEpcSubnetId"`
	SecurityGroupId  *string                        `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstanceGroups   []*LaunchClusterInstanceGroups `json:"InstanceGroups,omitempty" name:"InstanceGroups"`
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

type ListInstanceGroupsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListInstanceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListInstanceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceGroupsResponse struct {
	*ksyunhttp.BaseResponse
	InstanceGroups []struct {
		Id                 *string `json:"Id"`
		InstanceGroupType  *string `json:"InstanceGroupType"`
		ResourceType       *string `json:"ResourceType"`
		InstanceTypeCode   *string `json:"InstanceTypeCode"`
		InstanceCount      *int    `json:"InstanceCount"`
		VolumeType         *string `json:"VolumeType"`
		VolumeSize         *int    `json:"VolumeSize"`
		VolumeCount        *int    `json:"VolumeCount"`
		VpcId              *string `json:"VpcId"`
		VpcSubnetId        *string `json:"VpcSubnetId"`
		AvalabilityZone    *string `json:"AvalabilityZone"`
		MultiInstanceCount *int    `json:"MultiInstanceCount"`
		Instances          []struct {
			Id              *string `json:"Id"`
			InstanceGroupId *string `json:"InstanceGroupId"`
			InstanceId      *string `json:"InstanceId"`
			InstanceName    *string `json:"InstanceName"`
			InternalIp      *string `json:"InternalIp"`
			Volumes         []struct {
			} `json:"Volumes"`
		} `json:"Instances"`
	} `json:"InstanceGroups"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ListInstanceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceControlRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ControlType *string `json:"ControlType,omitempty" name:"ControlType"`
	Rolling     *string `json:"Rolling,omitempty" name:"Rolling"`
}

func (r *ServiceControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServiceControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ServiceControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ServiceControlResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *ServiceControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServiceControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterHealthStatisticRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ClusterHealthStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClusterHealthStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ClusterHealthStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClusterHealthStatisticResponse struct {
	*ksyunhttp.BaseResponse
	status *int `json:"status" name:"status"`
	Data   struct {
		RequestId          *string `json:"RequestId"`
		OperationId        *string `json:"OperationId"`
		HealthCheckHistory []struct {
			StartDate *string `json:"StartDate"`
			Red       *int    `json:"Red"`
			Green     *int    `json:"Green"`
			Yellow    *int    `json:"Yellow"`
			Failed    *int    `json:"Failed"`
			Stage     *string `json:"Stage"`
			Status    []struct {
				Item        *string `json:"Item"`
				Flag        *string `json:"Flag"`
				Description *string `json:"Description"`
				Suggestion  *string `json:"Suggestion"`
				Diagnosis   *string `json:"Diagnosis"`
			} `json:"Status"`
		} `json:"HealthCheckHistory"`
		StatusCode *int `json:"StatusCode"`
	} `json:"Data"`
	message *string `json:"message" name:"message"`
}

func (r *ClusterHealthStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClusterHealthStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckClusterHealthRequest struct {
	*ksyunhttp.BaseRequest
	Cluster_id *string   `json:"cluster_id,omitempty" name:"cluster_id"`
	Check_list []*string `json:"check_list,omitempty" name:"check_list"`
}

func (r *CheckClusterHealthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CheckClusterHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckClusterHealthResponse struct {
	*ksyunhttp.BaseResponse
	Return     *string `json:"Return" name:"Return"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	StatusCode *int    `json:"StatusCode" name:"StatusCode"`
}

func (r *CheckClusterHealthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
