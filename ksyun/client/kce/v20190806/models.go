package v20190806
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type DescribeClusterInstanceFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type AddClusterInstancesInstanceSet struct {
    NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
    NodePara []*string `json:"NodePara,omitempty" name:"NodePara"`
}

type AddClusterInstancesLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type AddClusterInstancesKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type AddClusterInstancesTaints struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type DescribeEpcForClusterFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type AddClusterEpcInstancesLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type AddClusterEpcInstancesKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type AddClusterEpcInstancesTaints struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type DescribeExistedInstancesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type AddExistedInstancesExistedInstanceKecSet struct {
    NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
    KecPara []*string `json:"KecPara,omitempty" name:"KecPara"`
}

type AddExistedInstancesLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type AddExistedInstancesKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type AddExistedInstancesTaints struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type CreateNodePoolDataDisk struct {
    Type *string `json:"Type,omitempty" name:"Type"`
    Size *int `json:"Size,omitempty" name:"Size"`
    DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

type CreateNodePoolKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type CreateNodePoolEbsTag struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateNodePoolInstanceTag struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateNodePoolLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateNodePoolTaint struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type ModifyNodePoolLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyNodePoolTaint struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type ModifyNodeTemplateDataDisk struct {
    Type *string `json:"Type,omitempty" name:"Type"`
    Size *int `json:"Size,omitempty" name:"Size"`
    DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

type ModifyNodeTemplateLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyNodeTemplateKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type ModifyNodeTemplateTaints struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type ModifyNodeTemplateEbsTag struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyNodeTemplateInstanceTag struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}


type DescribeClusterInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    Filter []*DescribeClusterInstanceFilter `json:"Filter,omitempty" name:"Filter"`
    Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeClusterInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceResponse struct {
    *ksyunhttp.BaseResponse
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		InstanceRole *string `json:"InstanceRole"`
		InstanceStatus *string `json:"InstanceStatus"`
		KecInstancePara struct {
				ProjectId *int `json:"ProjectId"`
				InstanceType *string `json:"InstanceType"`
			InstanceConfigure struct {
				VCPU *int `json:"VCPU"`
				MemoryGb *int `json:"MemoryGb"`
				GPU *int `json:"GPU"`
				DataDiskGb *int `json:"DataDiskGb"`
				RootDiskGb *int `json:"RootDiskGb"`
				DataDiskType *string `json:"DataDiskType"`
			} `json:"InstanceConfigure"`
			SystemDisk struct {
				DiskType *string `json:"DiskType"`
				DiskSize *int `json:"DiskSize"`
			} `json:"SystemDisk"`
				ImageId *string `json:"ImageId"`
				PrivateIpAddress *string `json:"PrivateIpAddress"`
				ChargeType *string `json:"ChargeType"`
				CreateTime *string `json:"CreateTime"`
				AvailabilityZone *string `json:"AvailabilityZone"`
				SubnetId *string `json:"SubnetId"`
				VpcId *string `json:"VpcId"`
			NetworkInterfaceSet []struct {
				NetworkInterfaceId *string `json:"NetworkInterfaceId"`
				NetworkInterfaceType *string `json:"NetworkInterfaceType"`
				SubnetId *string `json:"SubnetId"`
				PrivateIpAddress *string `json:"PrivateIpAddress"`
				MacAddress *string `json:"MacAddress"`
				SecurityGroupSet []struct {
							SecurityGroupId *string `json:"SecurityGroupId"`
					} `json:"SecurityGroupSet"`
				} `json:"NetworkInterfaceSet"`
				DedicatedName *string `json:"DedicatedName"`
				DedicatedId *string `json:"DedicatedId"`
		} `json:"KecInstancePara"`
		UnSchedulable *bool `json:"UnSchedulable"`
		DrainStatus *string `json:"DrainStatus"`
		NodePoolId *string `json:"NodePoolId"`
		AdvancedSetting struct {
			DataDisk struct {
				AutoFormatAndMount *bool `json:"AutoFormatAndMount"`
				FileSystem *string `json:"FileSystem"`
				MountTarget *string `json:"MountTarget"`
			} `json:"DataDisk"`
				ContainerRuntime *string `json:"ContainerRuntime"`
				DockerPath *string `json:"DockerPath"`
				ContainerPath *string `json:"ContainerPath"`
				UserScript *string `json:"UserScript"`
				PreUserScript *string `json:"PreUserScript"`
				Schedulable *bool `json:"Schedulable"`
			Label []struct {
				Key *string `json:"Key"`
				Value *string `json:"Value"`
			} `json:"Label"`
			ExtraArg struct {
				Kubelet []struct {
					} `json:"Kubelet"`
				} `json:"ExtraArg"`
				ContainerLogMaxSize *int `json:"ContainerLogMaxSize"`
				ContainerLogMaxFiles *int `json:"ContainerLogMaxFiles"`
		} `json:"AdvancedSetting"`
		EpcInstancePara struct {
				ProjectId *int `json:"ProjectId"`
				InstanceType *string `json:"InstanceType"`
			Cpu struct {
				Model *string `json:"Model"`
				Frequence *string `json:"Frequence"`
				Count *int `json:"Count"`
				CoreCount *int `json:"CoreCount"`
			} `json:"Cpu"`
				Memory *string `json:"Memory"`
			Gpu struct {
				Model *string `json:"Model"`
				Frequence *string `json:"Frequence"`
				Count *int `json:"Count"`
				CoreCount *int `json:"CoreCount"`
			} `json:"Gpu"`
			DiskSet []struct {
				DiskType *string `json:"DiskType"`
				Raid *string `json:"Raid"`
				Space *string `json:"Space"`
			} `json:"DiskSet"`
				OsName *string `json:"OsName"`
				ImageId *string `json:"ImageId"`
				CreateTime *string `json:"CreateTime"`
				AvailabilityZone *string `json:"AvailabilityZone"`
			NetworkInterfaceSet []struct {
				NetworkInterfaceId *string `json:"NetworkInterfaceId"`
				NetworkInterfaceType *string `json:"NetworkInterfaceType"`
				SubnetId *string `json:"SubnetId"`
				PrivateIpAddress *string `json:"PrivateIpAddress"`
				MacAddress *string `json:"MacAddress"`
				SecurityGroupSet []struct {
							SecurityGroupId *string `json:"SecurityGroupId"`
					} `json:"SecurityGroupSet"`
				} `json:"NetworkInterfaceSet"`
				EnableContainer *bool `json:"EnableContainer"`
				ProductType *string `json:"ProductType"`
		} `json:"EpcInstancePara"`
		Type *string `json:"Type"`
		ErrorMessage *string `json:"ErrorMessage"`
	} `json:"InstanceSet"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    MaxResults *int `json:"MaxResults" name:"MaxResults"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeClusterInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadClusterConfigRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
}

func (r *DownloadClusterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadClusterConfigRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DownloadClusterConfigRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DownloadClusterConfigResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Expose *bool `json:"Expose" name:"Expose"`
    ClusterConfig *string `json:"ClusterConfig" name:"ClusterConfig"`
}

func (r *DownloadClusterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadClusterConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInfoRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
    ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
    EnableKMSE *bool `json:"EnableKMSE,omitempty" name:"EnableKMSE"`
}

func (r *ModifyClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterInfoRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyClusterInfoRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInfoResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceImageRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeInstanceImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceImageResponse struct {
    *ksyunhttp.BaseResponse
	ImageSet []struct {
		ImageId *string `json:"ImageId"`
		ImageName *string `json:"ImageName"`
		ImageType *string `json:"ImageType"`
		MatchedK8sVersions []struct {
			} `json:"MatchedK8sVersions"`
		} `json:"ImageSet"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeInstanceImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstancesRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    InstanceSet []*AddClusterInstancesInstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet"`
}

func (r *AddClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddClusterInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstancesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
	} `json:"InstanceSet"`
}

func (r *AddClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
    InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteClusterInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		Message *string `json:"Message"`
		Return *bool `json:"Return"`
		InstanceId *string `json:"InstanceId"`
	} `json:"InstanceSet"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcForClusterRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
    Filter []*DescribeEpcForClusterFilter `json:"Filter,omitempty" name:"Filter"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeEpcForClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcForClusterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcForClusterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcForClusterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Marker *int `json:"Marker" name:"Marker"`
    MaxResults *int `json:"MaxResults" name:"MaxResults"`
    TotalCount *string `json:"TotalCount" name:"TotalCount"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		EpcInstancePara struct {
				ProjectId *int `json:"ProjectId"`
				InstanceType *string `json:"InstanceType"`
			Cpu struct {
				Model *string `json:"Model"`
				Frequence *string `json:"Frequence"`
				Count *int `json:"Count"`
				CoreCount *int `json:"CoreCount"`
			} `json:"Cpu"`
				Memory *string `json:"Memory"`
			Gpu struct {
				Model *string `json:"Model"`
				Frequence *string `json:"Frequence"`
				Count *int `json:"Count"`
				CoreCount *int `json:"CoreCount"`
			} `json:"Gpu"`
			DiskSet []struct {
				DiskType *string `json:"DiskType"`
				Raid *string `json:"Raid"`
				Space *string `json:"Space"`
			} `json:"DiskSet"`
				ImageId *string `json:"ImageId"`
				OsName *string `json:"OsName"`
				AvailabilityZone *string `json:"AvailabilityZone"`
			NetworkInterfaceSet []struct {
				NetworkInterfaceId *string `json:"NetworkInterfaceId"`
				NetworkInterfaceType *string `json:"NetworkInterfaceType"`
				SubnetId *string `json:"SubnetId"`
				PrivateIpAddress *string `json:"PrivateIpAddress"`
				MacAddress *string `json:"MacAddress"`
				SecurityGroupSet []struct {
							SecurityGroupId *string `json:"SecurityGroupId"`
					} `json:"SecurityGroupSet"`
				} `json:"NetworkInterfaceSet"`
				ProductType *string `json:"ProductType"`
				EnableContainer *bool `json:"EnableContainer"`
		} `json:"EpcInstancePara"`
		Type *string `json:"Type"`
		InstanceStatus *string `json:"InstanceStatus"`
	} `json:"InstanceSet"`
}

func (r *DescribeEpcForClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcForClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterEpcInstancesRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
    EpcPara []*string `json:"EpcPara,omitempty" name:"EpcPara"`
}

func (r *AddClusterEpcInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterEpcInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddClusterEpcInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterEpcInstancesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		Return *bool `json:"Return"`
		InstanceId *string `json:"InstanceId"`
	} `json:"InstanceSet"`
}

func (r *AddClusterEpcInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterEpcInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *string `json:"MaxResults,omitempty" name:"MaxResults"`
    Filter []*DescribeExistedInstancesFilter `json:"Filter,omitempty" name:"Filter"`
    Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExistedInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeExistedInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesResponse struct {
    *ksyunhttp.BaseResponse
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		InstanceType *string `json:"InstanceType"`
		PrivateIpAddress *string `json:"PrivateIpAddress"`
		Available *bool `json:"Available"`
		UnavailableReason *string `json:"UnavailableReason"`
		ClusterId *string `json:"ClusterId"`
	} `json:"InstanceSet"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
    MaxResults *int `json:"MaxResults" name:"MaxResults"`
    Marker *int `json:"Marker" name:"Marker"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExistedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddExistedInstancesRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    ExistedInstanceKecSet []*AddExistedInstancesExistedInstanceKecSet `json:"ExistedInstanceKecSet,omitempty" name:"ExistedInstanceKecSet"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddExistedInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddExistedInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddExistedInstancesResponse struct {
    *ksyunhttp.BaseResponse
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		Return *bool `json:"Return"`
		Reason *string `json:"Reason"`
	} `json:"InstanceSet"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddExistedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNodePoolRequest struct {
    *ksyunhttp.BaseRequest
    NodePoolName *string `json:"NodePoolName,omitempty" name:"NodePoolName"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    EnableAutoScale *bool `json:"EnableAutoScale,omitempty" name:"EnableAutoScale"`
    Label []*CreateNodePoolLabel `json:"Label,omitempty" name:"Label"`
    Taint []*CreateNodePoolTaint `json:"Taint,omitempty" name:"Taint"`
    MinSize *int `json:"MinSize,omitempty" name:"MinSize"`
    MaxSize *int `json:"MaxSize,omitempty" name:"MaxSize"`
    DesiredCapacity *int `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
}

func (r *CreateNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNodePoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNodePoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateNodePoolResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NodePoolId *string `json:"NodePoolId" name:"NodePoolId"`
}

func (r *CreateNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNodePoolResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodePoolRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId []*string `json:"NodePoolId,omitempty" name:"NodePoolId"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NodePoolName *string `json:"NodePoolName,omitempty" name:"NodePoolName"`
}

func (r *DescribeNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodePoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNodePoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodePoolResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    MaxResults *int `json:"MaxResults" name:"MaxResults"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	NodePoolSet []struct {
		NodePoolId *string `json:"NodePoolId"`
		NodePoolName *string `json:"NodePoolName"`
		Status *string `json:"Status"`
		EnableAutoScale *bool `json:"EnableAutoScale"`
		NodeTemplate struct {
				ChargeType *string `json:"ChargeType"`
				InstanceType *string `json:"InstanceType"`
				InstanceName *string `json:"InstanceName"`
				InstanceNameSuffix *string `json:"InstanceNameSuffix"`
			SystemDisk struct {
				DiskType *string `json:"DiskType"`
				DiskSize *int `json:"DiskSize"`
			} `json:"SystemDisk"`
				DataDiskGb *int `json:"DataDiskGb"`
			DataDisk struct {
				DiskType *string `json:"DiskType"`
				DiskSize *int `json:"DiskSize"`
				DeleteWithInstance *bool `json:"DeleteWithInstance"`
			} `json:"DataDisk"`
				ImageId *string `json:"ImageId"`
				VpcId *string `json:"VpcId"`
			SubnetIdSet []struct {
			} `json:"SubnetIdSet"`
				SubnetStrategy *string `json:"SubnetStrategy"`
				SecurityGroupId *string `json:"SecurityGroupId"`
				ProjectId *int `json:"ProjectId"`
				Password *string `json:"Password"`
			KeyIdSet []struct {
			} `json:"KeyIdSet"`
			AdvancedSetting struct {
				DataDisk struct {
						AutoFormatAndMount *bool `json:"AutoFormatAndMount"`
						FileSystem *string `json:"FileSystem"`
						MountTarget *string `json:"MountTarget"`
				} `json:"DataDisk"`
				ContainerRuntime *string `json:"ContainerRuntime"`
				DockerPath *string `json:"DockerPath"`
				ContainerPath *string `json:"ContainerPath"`
				UserScript *string `json:"UserScript"`
				PreUserScript *string `json:"PreUserScript"`
				Schedulable *bool `json:"Schedulable"`
				Label []struct {
							Key *string `json:"Key"`
							Value *string `json:"Value"`
					} `json:"Label"`
					ExtraArg struct {
						Kubelet []struct {
							CustomArg *string `json:"CustomArg"`
						} `json:"Kubelet"`
					} `json:"ExtraArg"`
					ContainerLogMaxSize *int `json:"ContainerLogMaxSize"`
					ContainerLogMaxFiles *int `json:"ContainerLogMaxFiles"`
					Taints []struct {
								Key *string `json:"Key"`
								Value *string `json:"Value"`
								Effect *string `json:"Effect"`
						} `json:"Taints"`
					} `json:"AdvancedSetting"`
			EbsTags []struct {
				Key *string `json:"Key"`
				Value *string `json:"Value"`
			} `json:"EbsTags"`
				ScalingConfigurationId *string `json:"ScalingConfigurationId"`
				ScalingConfigurationName *string `json:"ScalingConfigurationName"`
				Mem *string `json:"Mem"`
				RemovePolicy *string `json:"RemovePolicy"`
			InstanceTags []struct {
				Key *string `json:"Key"`
				Value *string `json:"Value"`
			} `json:"InstanceTags"`
				DeleteDataDisk *bool `json:"DeleteDataDisk"`
				DeleteInstanceTag *bool `json:"DeleteInstanceTag"`
				DeleteEbsTag *bool `json:"DeleteEbsTag"`
				Cpu *string `json:"Cpu"`
		} `json:"NodeTemplate"`
		Labels []struct {
					Key *string `json:"Key"`
					Value *string `json:"Value"`
			} `json:"Labels"`
			Taints []struct {
						Key *string `json:"Key"`
						Value *string `json:"Value"`
						Effect *string `json:"Effect"`
				} `json:"Taints"`
				MinSize *int `json:"MinSize"`
				MaxSize *int `json:"MaxSize"`
				DesiredCapacity *int `json:"DesiredCapacity"`
				InstanceCount *int `json:"InstanceCount"`
				CreateTime *string `json:"CreateTime"`
				ClusterId *string `json:"ClusterId"`
				ErrorStatusMessage *string `json:"ErrorStatusMessage"`
			} `json:"NodePoolSet"`
    ScaleUpPolicy *string `json:"ScaleUpPolicy" name:"ScaleUpPolicy"`
	ScaleDownPolicy struct {
		ScaleDownEnabled *bool `json:"ScaleDownEnabled"`
		ScaleDownDelayAfterAdd *int `json:"ScaleDownDelayAfterAdd"`
		ScaleDownUnneededTime *int `json:"ScaleDownUnneededTime"`
		ScaleDownUtilizationThreshold *int `json:"ScaleDownUtilizationThreshold"`
		MaxEmptyBulkDelete *int `json:"MaxEmptyBulkDelete"`
	} `json:"ScaleDownPolicy"`
}

func (r *DescribeNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodePoolResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNodePoolRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId []*string `json:"NodePoolId,omitempty" name:"NodePoolId"`
    InstanceDeleteMode *bool `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNodePoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNodePoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNodePoolResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet []struct {
		ScalingGroupId *string `json:"ScalingGroupId"`
		Message *string `json:"Message"`
		Return *bool `json:"Return"`
	} `json:"ReturnSet"`
}

func (r *DeleteNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNodePoolResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
    NodePoolName *string `json:"NodePoolName,omitempty" name:"NodePoolName"`
    EnableAutoScale *bool `json:"EnableAutoScale,omitempty" name:"EnableAutoScale"`
    MinSize *int `json:"MinSize,omitempty" name:"MinSize"`
    MaxSize *int `json:"MaxSize,omitempty" name:"MaxSize"`
    DesiredCapacity *int `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
    Label []*ModifyNodePoolLabel `json:"Label,omitempty" name:"Label"`
    Taint []*ModifyNodePoolTaint `json:"Taint,omitempty" name:"Taint"`
    UpdateExistingNodes *bool `json:"UpdateExistingNodes,omitempty" name:"UpdateExistingNodes"`
}

func (r *ModifyNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodePoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNodePoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodePoolResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeTemplateRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

func (r *ModifyNodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodeTemplateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNodeTemplateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeTemplateResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolScaleUpPolicyRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    ScaleUpPolicy *string `json:"ScaleUpPolicy,omitempty" name:"ScaleUpPolicy"`
}

func (r *ModifyNodePoolScaleUpPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodePoolScaleUpPolicyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNodePoolScaleUpPolicyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolScaleUpPolicyResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNodePoolScaleUpPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodePoolScaleUpPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolScaleDownPolicyRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MaxEmptyBulkDelete *int `json:"MaxEmptyBulkDelete,omitempty" name:"MaxEmptyBulkDelete"`
    ScaleDownDelayAfterAdd *int `json:"ScaleDownDelayAfterAdd,omitempty" name:"ScaleDownDelayAfterAdd"`
    ScaleDownEnabled *bool `json:"ScaleDownEnabled,omitempty" name:"ScaleDownEnabled"`
    ScaleDownUnneededTime *int `json:"ScaleDownUnneededTime,omitempty" name:"ScaleDownUnneededTime"`
    ScaleDownUtilizationThreshold *int `json:"ScaleDownUtilizationThreshold,omitempty" name:"ScaleDownUtilizationThreshold"`
}

func (r *ModifyNodePoolScaleDownPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodePoolScaleDownPolicyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNodePoolScaleDownPolicyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolScaleDownPolicyResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNodePoolScaleDownPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNodePoolScaleDownPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstanceToNodePoolRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AddClusterInstanceToNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstanceToNodePoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddClusterInstanceToNodePoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstanceToNodePoolResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		Return *bool `json:"Return"`
		Message *string `json:"Message"`
	} `json:"InstanceSet"`
}

func (r *AddClusterInstanceToNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstanceToNodePoolResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProtectedFromScaleDownRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
    ProtectedFromScaleDown *bool `json:"ProtectedFromScaleDown,omitempty" name:"ProtectedFromScaleDown"`
}

func (r *ProtectedFromScaleDownRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProtectedFromScaleDownRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ProtectedFromScaleDownRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ProtectedFromScaleDownResponse struct {
    *ksyunhttp.BaseResponse
    SucceedInstanceIds []*string `json:"SucceedInstanceIds" name:"SucceedInstanceIds"`
    FailedInstanceIds []*string `json:"FailedInstanceIds" name:"FailedInstanceIds"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ProtectedFromScaleDownResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProtectedFromScaleDownResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesFromNodePoolRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
    InstanceDeleteMode *bool `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterInstancesFromNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesFromNodePoolRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteClusterInstancesFromNodePoolRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesFromNodePoolResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		Return *bool `json:"Return"`
		Message *string `json:"Message"`
	} `json:"InstanceSet"`
}

func (r *DeleteClusterInstancesFromNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesFromNodePoolResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcImageRequest struct {
    *ksyunhttp.BaseRequest
    ImageId []*string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeEpcImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcImageResponse struct {
    *ksyunhttp.BaseResponse
	ImageSet []struct {
		ImageId *string `json:"ImageId"`
		ImageName *string `json:"ImageName"`
		ImageType *string `json:"ImageType"`
	} `json:"ImageSet"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeEpcImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditEventCollectingRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    EnableEventCollecting *bool `json:"EnableEventCollecting,omitempty" name:"EnableEventCollecting"`
}

func (r *EditEventCollectingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditEventCollectingRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EditEventCollectingRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type EditEventCollectingResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EditEventCollectingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditEventCollectingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodePoolSummaryRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeNodePoolSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodePoolSummaryRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNodePoolSummaryRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodePoolSummaryResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	NodePools []struct {
		NodePoolId *string `json:"NodePoolId"`
		NodePoolName *string `json:"NodePoolName"`
	} `json:"NodePools"`
}

func (r *DescribeNodePoolSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodePoolSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *DescribeClusterSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSummaryRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterSummaryRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	ClusterSet struct {
		ClusterId *string `json:"ClusterId"`
		ClusterName *string `json:"ClusterName"`
		ClusterManageMode *string `json:"ClusterManageMode"`
		K8sVersion *string `json:"K8sVersion"`
		PodCidr *string `json:"PodCidr"`
		ServiceCidr *string `json:"ServiceCidr"`
		VpcId *string `json:"VpcId"`
		VpcCidr *string `json:"VpcCidr"`
		NetworkType *string `json:"NetworkType"`
		Status *string `json:"Status"`
		CreateTime *string `json:"CreateTime"`
	} `json:"ClusterSet"`
}

func (r *DescribeClusterSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

