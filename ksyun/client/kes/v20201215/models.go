package v20201215
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
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
		Id           *string `json:"Id" name:"Id"`
		InstanceGroupType *string `json:"InstanceGroupType" name:"InstanceGroupType"`
		ResourceType *string `json:"ResourceType" name:"ResourceType"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		VolumeSize   *int    `json:"VolumeSize" name:"VolumeSize"`
		VolumeType   *string `json:"VolumeType" name:"VolumeType"`
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
		TagKey *string `json:"TagKey" name:"TagKey"`
		TagValue *string `json:"TagValue" name:"TagValue"`
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
		ClusterId      *string `json:"ClusterId" name:"ClusterId"`
		ClusterName    *string `json:"ClusterName" name:"ClusterName"`
		MainVersion    *string `json:"MainVersion" name:"MainVersion"`
		InstanceGroups []struct {
			Id                *string `json:"Id" name:"Id"`
			InstanceGroupType *string `json:"InstanceGroupType" name:"InstanceGroupType"`
			ResourceType      *string `json:"ResourceType" name:"ResourceType"`
			InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		} `json:"InstanceGroups" name:"InstanceGroups"`
		EnableEip      *bool   `json:"EnableEip" name:"EnableEip"`
		Region         *string `json:"Region" name:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		VpcDomainId    *string `json:"VpcDomainId" name:"VpcDomainId"`
		VpcSubnetId    *string `json:"VpcSubnetId" name:"VpcSubnetId"`
		ChargeType     *string `json:"ChargeType" name:"ChargeType"`
		ClusterStatus  *string `json:"ClusterStatus" name:"ClusterStatus"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime     *string `json:"UpdateTime" name:"UpdateTime"`
		ServingMinutes *int    `json:"ServingMinutes" name:"ServingMinutes"`
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
		Id                 *string `json:"Id" name:"Id"`
		InstanceGroupType  *string `json:"InstanceGroupType" name:"InstanceGroupType"`
		ResourceType       *string `json:"ResourceType" name:"ResourceType"`
		InstanceTypeCode   *string `json:"InstanceTypeCode" name:"InstanceTypeCode"`
		InstanceCount      *int    `json:"InstanceCount" name:"InstanceCount"`
		VolumeType         *string `json:"VolumeType" name:"VolumeType"`
		VolumeSize         *int    `json:"VolumeSize" name:"VolumeSize"`
		VolumeCount        *int    `json:"VolumeCount" name:"VolumeCount"`
		VpcId              *string `json:"VpcId" name:"VpcId"`
		VpcSubnetId        *string `json:"VpcSubnetId" name:"VpcSubnetId"`
		AvalabilityZone    *string `json:"AvalabilityZone" name:"AvalabilityZone"`
		MultiInstanceCount *int    `json:"MultiInstanceCount" name:"MultiInstanceCount"`
		Instances          []struct {
			Id              *string   `json:"Id" name:"Id"`
			InstanceGroupId *string   `json:"InstanceGroupId" name:"InstanceGroupId"`
			InstanceId      *string   `json:"InstanceId" name:"InstanceId"`
			InstanceName    *string   `json:"InstanceName" name:"InstanceName"`
			InternalIp      *string   `json:"InternalIp" name:"InternalIp"`
			Volumes         []*string `json:"Volumes" name:"Volumes"`
		} `json:"Instances" name:"Instances"`
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
	Status *int `json:"status" name:"status"`
	Data   struct {
		RequestId          *string `json:"RequestId" name:"RequestId"`
		OperationId        *string `json:"OperationId" name:"OperationId"`
		HealthCheckHistory []struct {
			StartDate *string `json:"StartDate" name:"StartDate"`
			Red       *int    `json:"Red" name:"Red"`
			Green     *int    `json:"Green" name:"Green"`
			Yellow    *int    `json:"Yellow" name:"Yellow"`
			Failed    *int    `json:"Failed" name:"Failed"`
			Stage     *string `json:"Stage" name:"Stage"`
			Status    []struct {
				Item       *string `json:"Item" name:"Item"`
				Flag       *string `json:"Flag" name:"Flag"`
				Description *string `json:"Description" name:"Description"`
				Suggestion *string `json:"Suggestion" name:"Suggestion"`
				Diagnosis  *string `json:"Diagnosis" name:"Diagnosis"`
			} `json:"Status"`
		} `json:"HealthCheckHistory" name:"HealthCheckHistory"`
		StatusCode *int `json:"StatusCode" name:"StatusCode"`
	} `json:"Data"`
	Message *string `json:"message" name:"message"`
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

