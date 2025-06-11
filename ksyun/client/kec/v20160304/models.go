package v20160304
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type DescribeInstancesFilter struct {
	Value []*string `json:"Value,omitempty" name:"Value"`
	Name  []*string `json:"Name,omitempty" name:"Name"`
}
type RunInstancesDataDisk struct {
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	SnapshotId         *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}
type RunInstancesNetworkInterface struct {
	SubnetId         *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId  []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	PrivateIpAddress *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}
type RunInstancesSystemDisk struct {
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	DiskSize *int    `json:"DiskSize,omitempty" name:"DiskSize"`
}
type RunInstancesTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Id    *int    `json:"Id,omitempty" name:"Id"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type RunInstancesInstanceConfigure struct {
	VCPU       *string `json:"VCPU,omitempty" name:"VCPU"`
	MemoryGb   *string `json:"MemoryGb,omitempty" name:"MemoryGb"`
	DataDiskGb *string `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
}
type ModifyInstanceTypeDataDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *string `json:"Size,omitempty" name:"Size"`
}
type CreateDedicatedHostsTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateScalingConfigurationDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	SnapshotId         *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}
type CreateScalingConfigurationTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateScalingGroupSlb struct {
	Id         *string `json:"Id,omitempty" name:"Id"`
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	ServerPort []*int  `json:"ServerPort,omitempty" name:"ServerPort"`
	Weight     *int    `json:"Weight,omitempty" name:"Weight"`
}
type ModifyScalingGroupSlb struct {
	Id         *string   `json:"Id,omitempty" name:"Id"`
	ListenerId *string   `json:"ListenerId,omitempty" name:"ListenerId"`
	ServerPort []*string `json:"ServerPort,omitempty" name:"ServerPort"`
	Weight     *int      `json:"Weight,omitempty" name:"Weight"`
}
type DescribeInstanceTypeConfigsFilter struct {
	Name  []*string `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeFileSystemsFilter struct {
	NameN  *string `json:"Name.N,omitempty" name:"Name.N"`
	ValueN *string `json:"Value.N,omitempty" name:"Value.N"`
}
type CreateModelDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	SnapshotId         *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	SnapshotName       *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}
type CreateModelSystemDisk struct {
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`
}
type CreateModelNetworkInterface struct {
	SubnetId         *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId  []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	PrivateIpAddress *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}
type CreateModelTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Id    *int    `json:"Id,omitempty" name:"Id"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyScalingConfigurationDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}
type ModifyScalingConfigurationTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type DescribePriceDataDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *int    `json:"Size,omitempty" name:"Size"`
}


type DescribeInstancesRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int                       `json:"Marker,omitempty" name:"Marker"`
	InstanceId []*string                  `json:"InstanceId,omitempty" name:"InstanceId"`
	ProjectId  []*string                  `json:"ProjectId,omitempty" name:"ProjectId"`
	Filter     []*DescribeInstancesFilter `json:"Filter,omitempty" name:"Filter"`
	Sort       *string                    `json:"Sort,omitempty" name:"Sort"`
	Search     *string                    `json:"Search,omitempty" name:"Search"`
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
	Marker        *int    `json:"Marker" name:"Marker"`
	InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
	RequestId     *string `json:"RequestId" name:"RequestId"`
	InstancesSet  []struct {
		InstanceId        *string `json:"InstanceId" name:"InstanceId"`
		ProjectId         *int    `json:"ProjectId" name:"ProjectId"`
		ShutdownNoCharge  *bool   `json:"ShutdownNoCharge" name:"ShutdownNoCharge"`
		IsDistributeIpv6  *bool   `json:"IsDistributeIpv6" name:"IsDistributeIpv6"`
		InstanceName      *string `json:"InstanceName" name:"InstanceName"`
		InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		InstanceConfigure struct {
			VCPU         *int    `json:"VCPU" name:"VCPU"`
			GPU          *int    `json:"GPU" name:"GPU"`
			MemoryGb     *int    `json:"MemoryGb" name:"MemoryGb"`
			DataDiskGb   *int    `json:"DataDiskGb" name:"DataDiskGb"`
			RootDiskGb   *int    `json:"RootDiskGb" name:"RootDiskGb"`
			DataDiskType *string `json:"DataDiskType" name:"DataDiskType"`
			VGPU         *string `json:"VGPU" name:"VGPU"`
			Spec         *string `json:"Spec" name:"Spec"`
		} `json:"InstanceConfigure" name:"InstanceConfigure"`
		ImageId       *string `json:"ImageId" name:"ImageId"`
		SubnetId      *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		InstanceState struct {
			Name               *string `json:"Name" name:"Name"`
			OnMigrate          *bool   `json:"OnMigrate" name:"OnMigrate"`
			TaskState          *string `json:"TaskState" name:"TaskState"`
			PreMigrationStatus *string `json:"PreMigrationStatus" name:"PreMigrationStatus"`
			CostTime           *string `json:"CostTime" name:"CostTime"`
			TimeStamp          *string `json:"TimeStamp" name:"TimeStamp"`
		} `json:"InstanceState" name:"InstanceState"`
		IsRunOptimised   *bool `json:"IsRunOptimised" name:"IsRunOptimised"`
		PreMigrateConfig struct {
			VmConfig struct {
				InstanceType *string `json:"InstanceType" name:"InstanceType"`
				Cpu          *int    `json:"Cpu" name:"Cpu"`
				Mem          *int    `json:"Mem" name:"Mem"`
				ProductType  *int    `json:"ProductType" name:"ProductType"`
				SystemType   *string `json:"SystemType" name:"SystemType"`
				SystemSize   *int    `json:"SystemSize" name:"SystemSize"`
				DedicatedHostId *string `json:"DedicatedHostId" name:"DedicatedHostId"`
				Model        *string `json:"Model" name:"Model"`
			} `json:"VmConfig"`
			VolumeConfig struct {
				VolumeType        *string `json:"VolumeType" name:"VolumeType"`
				VolumeSize        *int    `json:"VolumeSize" name:"VolumeSize"`
				VolumeProductType *int    `json:"VolumeProductType" name:"VolumeProductType"`
				VolumeId          *string `json:"VolumeId" name:"VolumeId"`
			} `json:"VolumeConfig"`
		} `json:"PreMigrateConfig" name:"PreMigrateConfig"`
		Monitoring struct {
			State *string `json:"State" name:"State"`
		} `json:"Monitoring" name:"Monitoring"`
		NetworkInterfaceSet []struct {
			NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
			VpcId                *string `json:"VpcId" name:"VpcId"`
			SubnetId             *string `json:"SubnetId" name:"SubnetId"`
			MacAddress           *string `json:"MacAddress" name:"MacAddress"`
			PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			GroupSet             []struct {
				GroupId *string `json:"GroupId" name:"GroupId"`
			} `json:"GroupSet"`
			SecurityGroupSet []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			} `json:"SecurityGroupSet"`
			NetworkInterfaceName *string `json:"NetworkInterfaceName" name:"NetworkInterfaceName"`
		} `json:"NetworkInterfaceSet" name:"NetworkInterfaceSet"`
		SriovNetSupport       *string `json:"SriovNetSupport" name:"SriovNetSupport"`
		IsShowSriovNetSupport *bool   `json:"IsShowSriovNetSupport" name:"IsShowSriovNetSupport"`
		CreationDate          *string `json:"CreationDate" name:"CreationDate"`
		AvailabilityZone      *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		AvailabilityZoneName  *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		DataGuardSet          []struct {
			DataGuardId   *string `json:"DataGuardId" name:"DataGuardId"`
			DataGuardName *string `json:"DataGuardName" name:"DataGuardName"`
			DataGuardType *string `json:"DataGuardType" name:"DataGuardType"`
		} `json:"DataGuardSet" name:"DataGuardSet"`
		DedicatedName    *string `json:"DedicatedName" name:"DedicatedName"`
		DedicatedClusterId *string `json:"DedicatedClusterId" name:"DedicatedClusterId"`
		DedicatedUuid    *string `json:"DedicatedUuid" name:"DedicatedUuid"`
		AutoScalingType  *string `json:"AutoScalingType" name:"AutoScalingType"`
		LargeEBSCapacity *string `json:"LargeEBSCapacity" name:"LargeEBSCapacity"`
		Metadata         []struct {
			Id           *string `json:"Id" name:"Id"`
			Key          *bool   `json:"Key" name:"Key"`
			Value        *string `json:"Value" name:"Value"`
			InstanceUuid *string `json:"Instance_uuid" name:"Instance_uuid"`
		} `json:"Metadata" name:"Metadata"`
		ProductType        *int    `json:"ProductType" name:"ProductType"`
		ProductWhat        *int    `json:"ProductWhat" name:"ProductWhat"`
		LiveUpgradeSupport *bool   `json:"LiveUpgradeSupport" name:"LiveUpgradeSupport"`
		ChargeType         *string `json:"ChargeType" name:"ChargeType"`
		SystemDisk         struct {
			DiskType *string `json:"DiskType" name:"DiskType"`
			DiskSize *int    `json:"DiskSize" name:"DiskSize"`
		} `json:"SystemDisk" name:"SystemDisk"`
		HostName  *string `json:"HostName" name:"HostName"`
		UserData  *string `json:"UserData" name:"UserData"`
		Migration *int    `json:"Migration" name:"Migration"`
		DataDisks []struct {
			DiskId             *string `json:"DiskId" name:"DiskId"`
			DiskType           *string `json:"DiskType" name:"DiskType"`
			DiskSize           *int    `json:"DiskSize" name:"DiskSize"`
			DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
			Encrypted          *bool   `json:"Encrypted" name:"Encrypted"`
		} `json:"DataDisks" name:"DataDisks"`
		VncSupport        *bool   `json:"VncSupport" name:"VncSupport"`
		Platform          *string `json:"Platform" name:"Platform"`
		FailureAutoDelete *bool   `json:"FailureAutoDelete" name:"FailureAutoDelete"`
		SyncTag           *bool   `json:"SyncTag" name:"SyncTag"`
	} `json:"InstancesSet"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RunInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ImageId                    *string                         `json:"ImageId,omitempty" name:"ImageId"`
	DedicatedHostId            *string                         `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
	InstanceType               *string                         `json:"InstanceType,omitempty" name:"InstanceType"`
	DataDiskGb                 *int                            `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	MaxCount                   *int                            `json:"MaxCount,omitempty" name:"MaxCount"`
	MinCount                   *int                            `json:"MinCount,omitempty" name:"MinCount"`
	SubnetId                   *string                         `json:"SubnetId,omitempty" name:"SubnetId"`
	InstancePassword           *string                         `json:"InstancePassword,omitempty" name:"InstancePassword"`
	ChargeType                 *string                         `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime               *int                            `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	SecurityGroupId            []*string                       `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	PrivateIpAddress           *string                         `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	InstanceName               *string                         `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix         *string                         `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	ProjectId                  *int                            `json:"ProjectId,omitempty" name:"ProjectId"`
	DataDisk                   []*RunInstancesDataDisk         `json:"DataDisk,omitempty" name:"DataDisk"`
	NetworkInterface           []*RunInstancesNetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`
	UserData                   *string                         `json:"UserData,omitempty" name:"UserData"`
	SystemDisk                 *RunInstancesSystemDisk         `json:"SystemDisk,omitempty" name:"SystemDisk"`
	ModelId                    *string                         `json:"ModelId,omitempty" name:"ModelId"`
	ModelVersion               *int                            `json:"ModelVersion,omitempty" name:"ModelVersion"`
	AssembledImageDataDiskType *string                         `json:"AssembledImageDataDiskType,omitempty" name:"AssembledImageDataDiskType"`
	AutoCreateEbs              *bool                           `json:"AutoCreateEbs,omitempty" name:"AutoCreateEbs"`
	LineId                     *string                         `json:"LineId,omitempty" name:"LineId"`
	AddressBandWidth           *int                            `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
	AddressChargeType          *string                         `json:"AddressChargeType,omitempty" name:"AddressChargeType"`
	AddressProjectId           *string                         `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
	BandWidthShareId           *string                         `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	AddressPurchaseTime        *int                            `json:"AddressPurchaseTime,omitempty" name:"AddressPurchaseTime"`
	KeyId                      []*string                       `json:"KeyId,omitempty" name:"KeyId"`
	KeepImageLogin             *bool                           `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
	HostName                   *string                         `json:"HostName,omitempty" name:"HostName"`
	HostNameSuffix             *int                            `json:"HostNameSuffix,omitempty" name:"HostNameSuffix"`
	FailureAutoDelete          *bool                           `json:"FailureAutoDelete,omitempty" name:"FailureAutoDelete"`
	Tag                        []*RunInstancesTag              `json:"Tag,omitempty" name:"Tag"`
	DataGuardId                *string                         `json:"DataGuardId,omitempty" name:"DataGuardId"`
	InstanceConfigure          *RunInstancesInstanceConfigure  `json:"InstanceConfigure,omitempty" name:"InstanceConfigure"`
	SriovNetSupport            *bool                           `json:"SriovNetSupport,omitempty" name:"SriovNetSupport"`
	DistributeIpv6             *bool                           `json:"DistributeIpv6,omitempty" name:"DistributeIpv6"`
	LocalVolumeSnapshotId      *string                         `json:"LocalVolumeSnapshotId,omitempty" name:"LocalVolumeSnapshotId"`
	SyncTag                    *bool                           `json:"SyncTag,omitempty" name:"SyncTag"`
	IamRoleName                *string                         `json:"IamRoleName,omitempty" name:"IamRoleName"`
	AutoDeleteTime             *string                         `json:"AutoDeleteTime,omitempty" name:"AutoDeleteTime"`
	AutoDeleteEip              *bool                           `json:"AutoDeleteEip,omitempty" name:"AutoDeleteEip"`
	IsProtect                  *bool                           `json:"IsProtect,omitempty" name:"IsProtect"`
}

func (r *RunInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
	} `json:"InstancesSet"`
}

func (r *RunInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StartInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return  *bool   `json:"Return" name:"Return"`
		Message *string `json:"Message" name:"Message"`
	} `json:"InstancesSet"`
}

func (r *StartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StopInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	ForceStop   *bool     `json:"ForceStop,omitempty" name:"ForceStop"`
	StoppedMode *string   `json:"StoppedMode,omitempty" name:"StoppedMode"`
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
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return  *bool   `json:"Return" name:"Return"`
		Messgae *string `json:"Messgae" name:"Messgae"`
	} `json:"InstancesSet"`
}

func (r *StopInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RebootInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	ForceReboot *bool     `json:"ForceReboot,omitempty" name:"ForceReboot"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return *bool `json:"Return" name:"Return"`
	} `json:"InstancesSet"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyInstanceAttributeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId       *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName     *string `json:"InstanceName,omitempty" name:"InstanceName"`
	InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
	HostName         *string `json:"HostName,omitempty" name:"HostName"`
	RestartMode      *string `json:"RestartMode,omitempty" name:"RestartMode"`
}

func (r *ModifyInstanceAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyInstanceAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyInstanceTypeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId                *string                       `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceType              *string                       `json:"InstanceType,omitempty" name:"InstanceType"`
	InstanceConfigureVCPU     *string                       `json:"InstanceConfigure.VCPU,omitempty" name:"InstanceConfigure.VCPU"`
	InstanceConfigureMemoryGb *string                       `json:"InstanceConfigure.MemoryGb,omitempty" name:"InstanceConfigure.MemoryGb"`
	DataDiskGb                *int                          `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	CrossInstanceMigrate      *bool                         `json:"CrossInstanceMigrate,omitempty" name:"CrossInstanceMigrate"`
	SystemDiskDiskType        *string                       `json:"SystemDisk.DiskType,omitempty" name:"SystemDisk.DiskType"`
	DataDisk                  []*ModifyInstanceTypeDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	StopInstance              *bool                         `json:"StopInstance,omitempty" name:"StopInstance"`
	AutoRestart               *bool                         `json:"AutoRestart,omitempty" name:"AutoRestart"`
	SystemDiskDiskSize        *int                          `json:"SystemDisk.DiskSize,omitempty" name:"SystemDisk.DiskSize"`
	SystemDiskResizeType      *string                       `json:"SystemDisk.ResizeType,omitempty" name:"SystemDisk.ResizeType"`
	InstantAccess             *bool                         `json:"InstantAccess,omitempty" name:"InstantAccess"`
}

func (r *ModifyInstanceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTypeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyInstanceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type TerminateInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	ForceDelete *bool     `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return  *bool   `json:"Return" name:"Return"`
		Message *string `json:"Message" name:"Message"`
	} `json:"InstancesSet"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeImagesRequest struct {
	*ksyunhttp.BaseRequest
	ImageId   *string `json:"ImageId,omitempty" name:"ImageId"`
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImagesSet []struct {
		ImageId       *string `json:"ImageId" name:"ImageId"`
		InstantAccess *bool   `json:"InstantAccess" name:"InstantAccess"`
		ContainerFormat *string `json:"ContainerFormat" name:"ContainerFormat"`
		SnapshotSet   []struct {
			SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
			DiskType   *string `json:"DiskType" name:"DiskType"`
			DiskSize   *int    `json:"DiskSize" name:"DiskSize"`
			Type       *string `json:"Type" name:"Type"`
			VolumeType *string `json:"VolumeType" name:"VolumeType"`
		} `json:"SnapshotSet" name:"SnapshotSet"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		ExportingStatus  *string `json:"ExportingStatus" name:"ExportingStatus"`
		ExportingFinishedTime *string `json:"ExportingFinishedTime" name:"ExportingFinishedTime"`
		Type             *string `json:"Type" name:"Type"`
		Name             *string `json:"Name" name:"Name"`
		ImageState       *string `json:"ImageState" name:"ImageState"`
		CreationDate     *string `json:"CreationDate" name:"CreationDate"`
		Platform         *string `json:"Platform" name:"Platform"`
		IsPublic         *bool   `json:"IsPublic" name:"IsPublic"`
		InstanceId       *string `json:"InstanceId" name:"InstanceId"`
		IsNpe            *bool   `json:"IsNpe" name:"IsNpe"`
		UserCategory     *string `json:"UserCategory" name:"UserCategory"`
		SysDisk          *int    `json:"SysDisk" name:"SysDisk"`
		Progress         *string `json:"Progress" name:"Progress"`
		ImageSource      *string `json:"ImageSource" name:"ImageSource"`
		CloudInitSupport *bool   `json:"CloudInitSupport" name:"CloudInitSupport"`
		Ipv6Support      *bool   `json:"Ipv6Support" name:"Ipv6Support"`
		IsModifyType     *bool   `json:"IsModifyType" name:"IsModifyType"`
		Architecture     *string `json:"Architecture" name:"Architecture"`
		FastBoot         *bool   `json:"FastBoot" name:"FastBoot"`
		OriginalImageId  *string `json:"OriginalImageId" name:"OriginalImageId"`
		IsCloudMarket    *bool   `json:"IsCloudMarket" name:"IsCloudMarket"`
		RealImageId      *string `json:"RealImageId" name:"RealImageId"`
		OnlineExpansion  *bool   `json:"OnlineExpansion" name:"OnlineExpansion"`
		ImageMode        *string `json:"ImageMode" name:"ImageMode"`
	} `json:"ImagesSet"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyImageAttributeRequest struct {
	*ksyunhttp.BaseRequest
	ImageId          *string `json:"ImageId,omitempty" name:"ImageId"`
	Name             *string `json:"Name,omitempty" name:"Name"`
	OsVersion        *string `json:"OsVersion,omitempty" name:"OsVersion"`
	CloudInitSupport *bool   `json:"CloudInitSupport,omitempty" name:"CloudInitSupport"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyImageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyInstanceImageRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId           *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	ImageId              *string   `json:"ImageId,omitempty" name:"ImageId"`
	SystemDiskDiskSize   *int      `json:"SystemDisk.DiskSize,omitempty" name:"SystemDisk.DiskSize"`
	InstancePassword     *string   `json:"InstancePassword,omitempty" name:"InstancePassword"`
	KeyId                []*string `json:"KeyId,omitempty" name:"KeyId"`
	KeepImageLogin       *bool     `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
	SystemDiskDiskType   *string   `json:"SystemDisk.DiskType,omitempty" name:"SystemDisk.DiskType"`
	SystemDiskResizeType *string   `json:"SystemDisk.ResizeType,omitempty" name:"SystemDisk.ResizeType"`
	UserData             *string   `json:"UserData,omitempty" name:"UserData"`
}

func (r *ModifyInstanceImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyInstanceImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateImageRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId    *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	Name          *string   `json:"Name,omitempty" name:"Name"`
	Type          *string   `json:"Type,omitempty" name:"Type"`
	DataDiskIds   []*string `json:"DataDiskIds,omitempty" name:"DataDiskIds"`
	SnapshotIds   []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	InstantAccess *bool     `json:"InstantAccess,omitempty" name:"InstantAccess"`
}

func (r *CreateImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageId   *string `json:"ImageId" name:"ImageId"`
}

func (r *CreateImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RemoveImagesRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *RemoveImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RemoveImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveImagesResponse struct {
	*ksyunhttp.BaseResponse
	ReturnSet []struct {
		ImageId *string `json:"ImageId" name:"ImageId"`
		Return *bool `json:"Return" name:"Return"`
	} `json:"ReturnSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RemoveImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyNetworkInterfaceAttributeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId         *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	NetworkInterfaceId *string   `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	SubnetId           *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId    []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	PrivateIpAddress   *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	DNS1               *string   `json:"DNS1,omitempty" name:"DNS1"`
	DNS2               *string   `json:"DNS2,omitempty" name:"DNS2"`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkInterfaceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNetworkInterfaceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyNetworkInterfaceAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkInterfaceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AttachNetworkInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId         *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	NetworkInterfaceId *string   `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	SubnetId           *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId    []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	PrivateIpAddress   *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	VpcIpv6            *string   `json:"VpcIpv6,omitempty" name:"VpcIpv6"`
	MacAddress         *string   `json:"MacAddress,omitempty" name:"MacAddress"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DetachNetworkInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId         *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeLocalVolumesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceName         *string `json:"InstanceName,omitempty" name:"InstanceName"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults           *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	LocalVolumeId        *string `json:"LocalVolumeId,omitempty" name:"LocalVolumeId"`
	InstanceState        *string `json:"InstanceState,omitempty" name:"InstanceState"`
	LocalVolumeCategory  *string `json:"LocalVolumeCategory,omitempty" name:"LocalVolumeCategory"`
	LocalVolumeSize      *int    `json:"LocalVolumeSize,omitempty" name:"LocalVolumeSize"`
	BindSnapshotPolicy   *bool   `json:"BindSnapshotPolicy,omitempty" name:"BindSnapshotPolicy"`
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	InstanceId           *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLocalVolumesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalVolumesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLocalVolumesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalVolumesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	Marker           *int    `json:"Marker" name:"Marker"`
	LocalVolumeCount *int    `json:"LocalVolumeCount" name:"LocalVolumeCount"`
	LocalVolumeSet   []struct {
		LocalVolumeId       *string `json:"LocalVolumeId" name:"LocalVolumeId"`
		LocalVolumeName     *string `json:"LocalVolumeName" name:"LocalVolumeName"`
		LocalVolumeState    *string `json:"LocalVolumeState" name:"LocalVolumeState"`
		InstanceId          *string `json:"InstanceId" name:"InstanceId"`
		InstanceName        *string `json:"InstanceName" name:"InstanceName"`
		InstanceState       *string `json:"InstanceState" name:"InstanceState"`
		LocalVolumeType     *string `json:"LocalVolumeType" name:"LocalVolumeType"`
		LocalVolumeCategory *string `json:"LocalVolumeCategory" name:"LocalVolumeCategory"`
		LocalVolumeSize     *int    `json:"LocalVolumeSize" name:"LocalVolumeSize"`
		CreationDate        *string `json:"CreationDate" name:"CreationDate"`
		AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId" name:"AutoSnapshotPolicyId"`
	} `json:"LocalVolumeSet"`
}

func (r *DescribeLocalVolumesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalVolumesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateLocalVolumeSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	LocalVolumeId           *string `json:"LocalVolumeId,omitempty" name:"LocalVolumeId"`
	LocalVolumeSnapshotName *string `json:"LocalVolumeSnapshotName,omitempty" name:"LocalVolumeSnapshotName"`
	LocalVolumeSnapshotDesc *string `json:"LocalVolumeSnapshotDesc,omitempty" name:"LocalVolumeSnapshotDesc"`
	InstantAccess           *bool   `json:"InstantAccess,omitempty" name:"InstantAccess"`
}

func (r *CreateLocalVolumeSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLocalVolumeSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateLocalVolumeSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLocalVolumeSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	LocalVolumeSnapshotId *string `json:"LocalVolumeSnapshotId" name:"LocalVolumeSnapshotId"`
}

func (r *CreateLocalVolumeSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLocalVolumeSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeLocalVolumeSnapshotsRequest struct {
	*ksyunhttp.BaseRequest
	Action              *string `json:"Action,omitempty" name:"Action"`
	Version             *string `json:"Version,omitempty" name:"Version"`
	LocalVolumeName     *string `json:"LocalVolumeName,omitempty" name:"LocalVolumeName"`
	SourceLocalVolumeId *string `json:"SourceLocalVolumeId,omitempty" name:"SourceLocalVolumeId"`
}

func (r *DescribeLocalVolumeSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalVolumeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLocalVolumeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalVolumeSnapshotsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	LocalVolumeSnapshotCount *int    `json:"LocalVolumeSnapshotCount" name:"LocalVolumeSnapshotCount"`
	Marker                   *int    `json:"Marker" name:"Marker"`
	LocalVolumeSnapshotSet   []struct {
		LocalVolumeSnapshotId   *string `json:"LocalVolumeSnapshotId" name:"LocalVolumeSnapshotId"`
		InstantAccess           *bool   `json:"InstantAccess" name:"InstantAccess"`
		LocalVolumeSnapshotName *string `json:"LocalVolumeSnapshotName" name:"LocalVolumeSnapshotName"`
		LocalVolumeSnapshotDesc *string `json:"LocalVolumeSnapshotDesc" name:"LocalVolumeSnapshotDesc"`
		SourceLocalVolumeId     *string `json:"SourceLocalVolumeId" name:"SourceLocalVolumeId"`
		SourceLocalVolumeName   *string `json:"SourceLocalVolumeName" name:"SourceLocalVolumeName"`
		ScheduledDeleteTime     *string `json:"ScheduledDeleteTime" name:"ScheduledDeleteTime"`
		CreateImage             *bool   `json:"CreateImage" name:"CreateImage"`
		CopyFromRemote          *bool   `json:"CopyFromRemote" name:"CopyFromRemote"`
		SourceLocalVolumeCategory *string `json:"SourceLocalVolumeCategory" name:"SourceLocalVolumeCategory"`
		SourceLocalVolumeState  *string `json:"SourceLocalVolumeState" name:"SourceLocalVolumeState"`
		State                   *string `json:"State" name:"State"`
		CreationDate            *string `json:"CreationDate" name:"CreationDate"`
		InstanceId              *string `json:"InstanceId" name:"InstanceId"`
		DiskSize                *int    `json:"DiskSize" name:"DiskSize"`
		SnapshotType            *string `json:"SnapshotType" name:"SnapshotType"`
		MinDisk                 *int    `json:"MinDisk" name:"MinDisk"`
	} `json:"LocalVolumeSnapshotSet"`
}

func (r *DescribeLocalVolumeSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalVolumeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RollbackLocalVolumeRequest struct {
	*ksyunhttp.BaseRequest
	LocalVolumeSnapshotId *string `json:"LocalVolumeSnapshotId,omitempty" name:"LocalVolumeSnapshotId"`
}

func (r *RollbackLocalVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackLocalVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RollbackLocalVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RollbackLocalVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RollbackLocalVolumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackLocalVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteLocalVolumeSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	LocalVolumeSnapshotId []*string `json:"LocalVolumeSnapshotId,omitempty" name:"LocalVolumeSnapshotId"`
}

func (r *DeleteLocalVolumeSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLocalVolumeSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteLocalVolumeSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLocalVolumeSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet []struct {
		LocalVolumeSnapshotId *string `json:"LocalVolumeSnapshotId" name:"LocalVolumeSnapshotId"`
		Return    *bool   `json:"Return" name:"Return"`
		Reason    *string `json:"Reason" name:"Reason"`
		ErrorCode *string `json:"ErrorCode" name:"ErrorCode"`
	} `json:"ReturnSet"`
}

func (r *DeleteLocalVolumeSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLocalVolumeSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyDataGuardGroupsRequest struct {
	*ksyunhttp.BaseRequest
	DataGuardId   *string `json:"DataGuardId,omitempty" name:"DataGuardId"`
	DataGuardName *string `json:"DataGuardName,omitempty" name:"DataGuardName"`
}

func (r *ModifyDataGuardGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataGuardGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDataGuardGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataGuardGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyDataGuardGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataGuardGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDataGuardCapacityRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDataGuardCapacityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataGuardCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDataGuardCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataGuardCapacityResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	DataGuardCapacity struct {
		Region   *string `json:"Region" name:"Region"`
		Capacity *int    `json:"Capacity" name:"Capacity"`
	} `json:"DataGuardCapacity"`
	DataGuardCapacitySet []struct {
		DomainTypeCapacity *int `json:"DomainTypeCapacity" name:"DomainTypeCapacity"`
		HostTypeCapacity *int `json:"HostTypeCapacity" name:"HostTypeCapacity"`
	} `json:"DataGuardCapacitySet"`
}

func (r *DescribeDataGuardCapacityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataGuardCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateDataGuardGroupRequest struct {
	*ksyunhttp.BaseRequest
	DataGuardName *string `json:"DataGuardName,omitempty" name:"DataGuardName"`
}

func (r *CreateDataGuardGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataGuardGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDataGuardGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataGuardGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	DataGuardId *string `json:"DataGuardId" name:"DataGuardId"`
}

func (r *CreateDataGuardGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataGuardGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteDataGuardGroupsRequest struct {
	*ksyunhttp.BaseRequest
	DataGuardId []*string `json:"DataGuardId,omitempty" name:"DataGuardId"`
}

func (r *DeleteDataGuardGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataGuardGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDataGuardGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataGuardGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet []struct {
		DataGuardId *string `json:"DataGuardId" name:"DataGuardId"`
		Return *bool `json:"Return" name:"Return"`
	} `json:"ReturnSet"`
}

func (r *DeleteDataGuardGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataGuardGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDataGuardGroupRequest struct {
	*ksyunhttp.BaseRequest
	DataGuardId   *string `json:"DataGuardId,omitempty" name:"DataGuardId"`
	DataGuardName *string `json:"DataGuardName,omitempty" name:"DataGuardName"`
}

func (r *DescribeDataGuardGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataGuardGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDataGuardGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataGuardGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	DataGuardsSet []struct {
		DataGuardId            *string `json:"DataGuardId" name:"DataGuardId"`
		DataGuardName          *string `json:"DataGuardName" name:"DataGuardName"`
		DataGuardLevel         *string `json:"DataGuardLevel" name:"DataGuardLevel"`
		DataGuardCapacity      *int    `json:"DataGuardCapacity" name:"DataGuardCapacity"`
		DataGuardUsedSize      *int    `json:"DataGuardUsedSize" name:"DataGuardUsedSize"`
		DataGuardInstancesList []struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			InstanceName *string `json:"InstanceName" name:"InstanceName"`
		} `json:"DataGuardInstancesList" name:"DataGuardInstancesList"`
		DataGuardType *string `json:"DataGuardType" name:"DataGuardType"`
	} `json:"DataGuardsSet"`
	IsSupportDomain *bool `json:"isSupportDomain" name:"isSupportDomain"`
}

func (r *DescribeDataGuardGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataGuardGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RemoveVmFromDataGuardRequest struct {
	*ksyunhttp.BaseRequest
	DataGuardId *string   `json:"DataGuardId,omitempty" name:"DataGuardId"`
	InstanceId  []*string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RemoveVmFromDataGuardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveVmFromDataGuardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RemoveVmFromDataGuardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveVmFromDataGuardResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return *bool `json:"Return" name:"Return"`
	} `json:"ReturnSet"`
}

func (r *RemoveVmFromDataGuardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveVmFromDataGuardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateDedicatedHostsRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedType      *string                    `json:"DedicatedType,omitempty" name:"DedicatedType"`
	Number             *int                       `json:"Number,omitempty" name:"Number"`
	Name               *string                    `json:"Name,omitempty" name:"Name"`
	ChargeType         *string                    `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime       *int                       `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	InstanceNameSuffix *string                    `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	DedicatedClusterId *string                    `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	Tag                []*CreateDedicatedHostsTag `json:"Tag,omitempty" name:"Tag"`
	AvailabilityZone   *string                    `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ProjectId          *int                       `json:"ProjectId,omitempty" name:"ProjectId"`
	EbsClusterMode     *string                    `json:"EbsClusterMode,omitempty" name:"EbsClusterMode"`
}

func (r *CreateDedicatedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDedicatedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDedicatedHostsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	DedicatedHostSet []struct {
		DedicatedHostId *string `json:"DedicatedHostId" name:"DedicatedHostId"`
		DedicatedHostName *string `json:"DedicatedHostName" name:"DedicatedHostName"`
	} `json:"DedicatedHostSet"`
}

func (r *CreateDedicatedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteDedicatedHostRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedHostId []*string `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
	IsRefund        *bool     `json:"IsRefund,omitempty" name:"IsRefund"`
}

func (r *DeleteDedicatedHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDedicatedHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDedicatedHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDedicatedHostResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		DedicatedHostId *string `json:"DedicatedHostId" name:"DedicatedHostId"`
		Return  *bool   `json:"Return" name:"Return"`
		Message *string `json:"Message" name:"Message"`
	} `json:"ReturnSet"`
}

func (r *DeleteDedicatedHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDedicatedHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDedicatedHostsRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
	Search          *string `json:"search,omitempty" name:"search"`
	ProjectId       []*int  `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeDedicatedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDedicatedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedHostsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	DedicatedHostSet []struct {
		DedicatedHostId      *string   `json:"DedicatedHostId" name:"DedicatedHostId"`
		DedicatedHostName    *string   `json:"DedicatedHostName" name:"DedicatedHostName"`
		State                *string   `json:"State" name:"State"`
		TotalCpu             *int      `json:"TotalCpu" name:"TotalCpu"`
		OriCpu               *int      `json:"OriCpu" name:"OriCpu"`
		AvailableCpu         *int      `json:"AvailableCpu" name:"AvailableCpu"`
		TotalMemory          *int      `json:"TotalMemory" name:"TotalMemory"`
		AvailableMemory      *int      `json:"AvailableMemory" name:"AvailableMemory"`
		TotalDatadisk        *int      `json:"TotalDatadisk" name:"TotalDatadisk"`
		AvailableDatadisk    *int      `json:"AvailableDatadisk" name:"AvailableDatadisk"`
		Instances            []*string `json:"Instances" name:"Instances"`
		CreateDate           *string   `json:"CreateDate" name:"CreateDate"`
		AvailabilityZone     *string   `json:"AvailabilityZone" name:"AvailabilityZone"`
		AvailabilityZoneName *string   `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		Model                *string   `json:"Model" name:"Model"`
		ProjectId            *int      `json:"ProjectId" name:"ProjectId"`
		EbsClusterMode       *string   `json:"EbsClusterMode" name:"EbsClusterMode"`
	} `json:"DedicatedHostSet"`
}

func (r *DescribeDedicatedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScalingConfigurationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingConfigurationName *string   `json:"ScalingConfigurationName,omitempty" name:"ScalingConfigurationName"`
	ScalingConfigurationId   []*string `json:"ScalingConfigurationId,omitempty" name:"ScalingConfigurationId"`
	Marker                   *int      `json:"Marker,omitempty" name:"Marker"`
	ProjectId                []*string `json:"ProjectId,omitempty" name:"ProjectId"`
	MaxResults               *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeScalingConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingConfigurationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"RequestId" name:"RequestId"`
	ScalingConfigurationCount *int    `json:"ScalingConfigurationCount" name:"ScalingConfigurationCount"`
	ScalingConfigurationSet   []struct {
		AvailabilityZone        *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Description             *string `json:"Description" name:"Description"`
		ContainerLabel          *string `json:"ContainerLabel" name:"ContainerLabel"`
		HostName                *string `json:"HostName" name:"HostName"`
		HostNameSuffix          *int    `json:"HostNameSuffix" name:"HostNameSuffix"`
		KeyIds                  *string `json:"KeyIds" name:"KeyIds"`
		DataDiskPerformanceMode *string `json:"DataDiskPerformanceMode" name:"DataDiskPerformanceMode"`
		ScalingConfigurationId  *string `json:"ScalingConfigurationId" name:"ScalingConfigurationId"`
		ScalingConfigurationName *string `json:"ScalingConfigurationName" name:"ScalingConfigurationName"`
		ScalingGroupIdSet       []struct {
			ScalingGroupId   *string `json:"ScalingGroupId" name:"ScalingGroupId"`
			ScalingGroupName *string `json:"ScalingGroupName" name:"ScalingGroupName"`
		} `json:"ScalingGroupIdSet" name:"ScalingGroupIdSet"`
		Cpu               *int      `json:"Cpu" name:"Cpu"`
		Mem               *int      `json:"Mem" name:"Mem"`
		StorageType       *string   `json:"StorageType" name:"StorageType"`
		StorageSize       *int      `json:"StorageSize" name:"StorageSize"`
		RootSize          *int      `json:"RootSize" name:"RootSize"`
		Gpu               *int      `json:"Gpu" name:"Gpu"`
		ImageId           *string   `json:"ImageId" name:"ImageId"`
		NeedMonitorAgent  *int      `json:"NeedMonitorAgent" name:"NeedMonitorAgent"`
		NeedSecurityAgent *int      `json:"NeedSecurityAgent" name:"NeedSecurityAgent"`
		ChargeType        *string   `json:"ChargeType" name:"ChargeType"`
		InstanceType      *string   `json:"InstanceType" name:"InstanceType"`
		InstanceTypeSet   []*string `json:"InstanceTypeSet" name:"InstanceTypeSet"`
		Available         *int      `json:"Available" name:"Available"`
		CreateTime        *string   `json:"CreateTime" name:"CreateTime"`
		Password          *string   `json:"Password" name:"Password"`
		ProductLine       *string   `json:"ProductLine" name:"ProductLine"`
		ContainerParam    *string   `json:"ContainerParam" name:"ContainerParam"`
		InstanceName      *string   `json:"InstanceName" name:"InstanceName"`
		Project           struct {
			ProjectId   *int    `json:"ProjectId" name:"ProjectId"`
			AccountId   *int    `json:"AccountId" name:"AccountId"`
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			ProjectDesc *string `json:"ProjectDesc" name:"ProjectDesc"`
			Krn         *string `json:"Krn" name:"Krn"`
			Status      *int    `json:"Status" name:"Status"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
		} `json:"Project" name:"Project"`
		ProjectId                *int      `json:"ProjectId" name:"ProjectId"`
		KeepImageLogin           *bool     `json:"KeepImageLogin" name:"KeepImageLogin"`
		DataDiskEbsDetail        *string   `json:"DataDiskEbsDetail" name:"DataDiskEbsDetail"`
		SystemDiskType           *string   `json:"SystemDiskType" name:"SystemDiskType"`
		SystemDiskSize           *int      `json:"SystemDiskSize" name:"SystemDiskSize"`
		AvailabilityZoneName     *string   `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		InstanceNameTimeSuffix   *bool     `json:"InstanceNameTimeSuffix" name:"InstanceNameTimeSuffix"`
		SupportAvailabilityZones []*string `json:"SupportAvailabilityZones" name:"SupportAvailabilityZones"`
		BillType                 *int      `json:"BillType" name:"BillType"`
		IsRunOptimised           *bool     `json:"IsRunOptimised" name:"IsRunOptimised"`
		UserData                 *string   `json:"UserData" name:"UserData"`
		AddressBandWidth         *int      `json:"AddressBandWidth" name:"AddressBandWidth"`
		LineId                   *string   `json:"LineId" name:"LineId"`
		AddressProjectId         *int      `json:"AddressProjectId" name:"AddressProjectId"`
		AddressChargeType        *string   `json:"AddressChargeType" name:"AddressChargeType"`
		BandWidthShareId         *string   `json:"BandWidthShareId" name:"BandWidthShareId"`
		LocalVolumeSnapshotId    *string   `json:"LocalVolumeSnapshotId" name:"LocalVolumeSnapshotId"`
		LocalVolumeSnapshotName  *string   `json:"LocalVolumeSnapshotName" name:"LocalVolumeSnapshotName"`
		AddressProject           *string   `json:"AddressProject" name:"AddressProject"`
		TagSet                   []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"TagSet" name:"TagSet"`
		EbsTagSet []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"EbsTagSet" name:"EbsTagSet"`
		SpotStrategy *string `json:"SpotStrategy" name:"SpotStrategy"`
		SpotPriceLimit *float64 `json:"SpotPriceLimit" name:"SpotPriceLimit"`
		SyncTag      *bool   `json:"SyncTag" name:"SyncTag"`
	} `json:"ScalingConfigurationSet"`
}

func (r *DescribeScalingConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateScalingConfigurationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingConfigurationName *string                               `json:"ScalingConfigurationName,omitempty" name:"ScalingConfigurationName"`
	ImageId                  *string                               `json:"ImageId,omitempty" name:"ImageId"`
	Password                 *string                               `json:"Password,omitempty" name:"Password"`
	InstanceType             *string                               `json:"InstanceType,omitempty" name:"InstanceType"`
	ChargeType               *string                               `json:"ChargeType,omitempty" name:"ChargeType"`
	DataDiskGb               *int                                  `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	ProjectId                *int                                  `json:"ProjectId,omitempty" name:"ProjectId"`
	KeepImageLogin           *bool                                 `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
	KeyId                    []*string                             `json:"KeyId,omitempty" name:"KeyId"`
	DataDisk                 []*CreateScalingConfigurationDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	SystemDiskDiskSize       *string                               `json:"SystemDisk.DiskSize,omitempty" name:"SystemDisk.DiskSize"`
	AddressBandWidth         *int                                  `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
	BandWidthShareId         *string                               `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	LineId                   *string                               `json:"LineId,omitempty" name:"LineId"`
	AddressProjectId         *int                                  `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
	InstanceName             *string                               `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix       *string                               `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	UserData                 *string                               `json:"UserData,omitempty" name:"UserData"`
	InstanceNameTimeSuffix   *bool                                 `json:"InstanceNameTimeSuffix,omitempty" name:"InstanceNameTimeSuffix"`
	Tag                      []*CreateScalingConfigurationTag      `json:"Tag,omitempty" name:"Tag"`
	SystemDiskDiskType       *string                               `json:"SystemDisk.DiskType,omitempty" name:"SystemDisk.DiskType"`
	SystemDiskResizeType     *string                               `json:"SystemDisk.ResizeType,omitempty" name:"SystemDisk.ResizeType"`
	SyncTag                  *bool                                 `json:"SyncTag,omitempty" name:"SyncTag"`
}

func (r *CreateScalingConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScalingConfigurationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId" name:"ScalingConfigurationId"`
}

func (r *CreateScalingConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteScalingConfigurationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingConfigurationId []*string `json:"ScalingConfigurationId,omitempty" name:"ScalingConfigurationId"`
}

func (r *DeleteScalingConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingConfigurationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet []struct {
		ScalingConfigurationId *string `json:"ScalingConfigurationId" name:"ScalingConfigurationId"`
		Message *string `json:"Message" name:"Message"`
	} `json:"ReturnSet"`
}

func (r *DeleteScalingConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateScalingGroupRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupName       *string                  `json:"ScalingGroupName,omitempty" name:"ScalingGroupName"`
	ScalingConfigurationId *string                  `json:"ScalingConfigurationId,omitempty" name:"ScalingConfigurationId"`
	MinSize                *string                  `json:"MinSize,omitempty" name:"MinSize"`
	DesiredCapacity        *int                     `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
	RemovePolicy           *string                  `json:"RemovePolicy,omitempty" name:"RemovePolicy"`
	SubnetId               []*string                `json:"SubnetId,omitempty" name:"SubnetId"`
	SubnetStrategy         *string                  `json:"SubnetStrategy,omitempty" name:"SubnetStrategy"`
	SecurityGroupId        *string                  `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Slb                    []*CreateScalingGroupSlb `json:"Slb,omitempty" name:"Slb"`
}

func (r *CreateScalingGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScalingGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingGroupId *string `json:"ScalingGroupId" name:"ScalingGroupId"`
	} `json:"ReturnSet"`
}

func (r *CreateScalingGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScalingGroupRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId         []*string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingGroupName       *string   `json:"ScalingGroupName,omitempty" name:"ScalingGroupName"`
	ScalingConfigurationId *string   `json:"ScalingConfigurationId,omitempty" name:"ScalingConfigurationId"`
	VpcId                  *string   `json:"VpcId,omitempty" name:"VpcId"`
	Marker                 *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults             *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	ScalingActivityId      []*string `json:"ScalingActivityId,omitempty" name:"ScalingActivityId"`
}

func (r *DescribeScalingGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	ScalingGroupCount *int    `json:"ScalingGroupCount" name:"ScalingGroupCount"`
	ScalingGroupSet   []struct {
		ScalingGroupId     *string `json:"ScalingGroupId" name:"ScalingGroupId"`
		BInScalingActivity *int    `json:"BInScalingActivity" name:"BInScalingActivity"`
		SlbId              *string `json:"SlbId" name:"SlbId"`
		ServerPort         *int    `json:"ServerPort" name:"ServerPort"`
		ListenerId         *string `json:"ListenerId" name:"ListenerId"`
		Description        *string `json:"Description" name:"Description"`
		ServerPorts        *string `json:"ServerPorts" name:"ServerPorts"`
		SlbConfigSet       []struct {
			SlbId           *string `json:"SlbId" name:"SlbId"`
			ListenerId      *string `json:"ListenerId" name:"ListenerId"`
			ServerPortSet   []*int  `json:"ServerPortSet" name:"ServerPortSet"`
			Weight          *int    `json:"Weight" name:"Weight"`
			HealthCheckType *string `json:"HealthCheckType" name:"HealthCheckType"`
		} `json:"SlbConfigSet" name:"SlbConfigSet"`
		ScalingGroupName         *string   `json:"ScalingGroupName" name:"ScalingGroupName"`
		ScalingConfigurationId   *string   `json:"ScalingConfigurationId" name:"ScalingConfigurationId"`
		ScalingConfigurationName *string   `json:"ScalingConfigurationName" name:"ScalingConfigurationName"`
		MinSize                  *int      `json:"MinSize" name:"MinSize"`
		MaxSize                  *int      `json:"MaxSize" name:"MaxSize"`
		InstanceNum              *int      `json:"InstanceNum" name:"InstanceNum"`
		CreateTime               *string   `json:"CreateTime" name:"CreateTime"`
		RemovePolicy             *string   `json:"RemovePolicy" name:"RemovePolicy"`
		VpcId                    *string   `json:"VpcId" name:"VpcId"`
		SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		SubnetId                 *string   `json:"SubnetId" name:"SubnetId"`
		Status                   *string   `json:"Status" name:"Status"`
		AvailabilityZone         *string   `json:"AvailabilityZone" name:"AvailabilityZone"`
		ProductLine              *string   `json:"ProductLine" name:"ProductLine"`
		AvailabilityZoneName     *string   `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		DesiredCapacity          *int      `json:"DesiredCapacity" name:"DesiredCapacity"`
		DistributeIpv6           *bool     `json:"DistributeIpv6" name:"DistributeIpv6"`
		SubnetIdSet              []*string `json:"SubnetIdSet" name:"SubnetIdSet"`
		SubnetStrategy           *string   `json:"SubnetStrategy" name:"SubnetStrategy"`
		SecurityGroupIdSet       []*string `json:"SecurityGroupIdSet" name:"SecurityGroupIdSet"`
		LockDesiredCapacity      *bool     `json:"LockDesiredCapacity" name:"LockDesiredCapacity"`
		AllowSlbScalingPolicy    *bool     `json:"AllowSlbScalingPolicy" name:"AllowSlbScalingPolicy"`
	} `json:"ScalingGroupSet"`
}

func (r *DescribeScalingGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyScalingGroupRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId         *string                  `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	MinSize                *int                     `json:"MinSize,omitempty" name:"MinSize"`
	MaxSize                *int                     `json:"MaxSize,omitempty" name:"MaxSize"`
	DesiredCapacity        *int                     `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
	RemovePolicy           *string                  `json:"RemovePolicy,omitempty" name:"RemovePolicy"`
	ScalingGroupName       *string                  `json:"ScalingGroupName,omitempty" name:"ScalingGroupName"`
	ScalingConfigurationId *string                  `json:"ScalingConfigurationId,omitempty" name:"ScalingConfigurationId"`
	SubnetId               []*string                `json:"SubnetId,omitempty" name:"SubnetId"`
	SubnetStrategy         *string                  `json:"SubnetStrategy,omitempty" name:"SubnetStrategy"`
	Slb                    []*ModifyScalingGroupSlb `json:"Slb,omitempty" name:"Slb"`
	ContainerSubnetId      []*string                `json:"ContainerSubnetId,omitempty" name:"ContainerSubnetId"`
}

func (r *ModifyScalingGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
	ReturnSet struct {
		ScalingGroupId *string `json:"ScalingGroupId" name:"ScalingGroupId"`
	} `json:"ReturnSet"`
}

func (r *ModifyScalingGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetKvmProtectedDetachRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId      *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingInstanceId   []*string `json:"ScalingInstanceId,omitempty" name:"ScalingInstanceId"`
	ProtectedFromDetach *int      `json:"ProtectedFromDetach,omitempty" name:"ProtectedFromDetach"`
}

func (r *SetKvmProtectedDetachRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetKvmProtectedDetachRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetKvmProtectedDetachRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetKvmProtectedDetachResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetKvmProtectedDetachResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetKvmProtectedDetachResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScalingInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId    *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingInstanceId []*string `json:"ScalingInstanceId,omitempty" name:"ScalingInstanceId"`
	CreationType      *string   `json:"CreationType,omitempty" name:"CreationType"`
	HealthStatus      *string   `json:"HealthStatus,omitempty" name:"HealthStatus"`
	Marker            *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults        *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeScalingInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScalingInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	ScalingInstanceCount *int    `json:"ScalingInstanceCount" name:"ScalingInstanceCount"`
	ScalingInstanceSet   []struct {
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		InstanceName         *string `json:"InstanceName" name:"InstanceName"`
		HealthStatus         *string `json:"HealthStatus" name:"HealthStatus"`
		CreationType         *string `json:"CreationType" name:"CreationType"`
		AddTime              *string `json:"AddTime" name:"AddTime"`
		ProtectedFromScaleIn *int    `json:"ProtectedFromScaleIn" name:"ProtectedFromScaleIn"`
		HostName             *string `json:"HostName" name:"HostName"`
		LifeCycleState       *string `json:"LifeCycleState" name:"LifeCycleState"`
		AvailabilityZone     *string `json:"AvailabilityZone" name:"AvailabilityZone"`
	} `json:"ScalingInstanceSet"`
}

func (r *DescribeScalingInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AttachInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId    *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingInstanceId []*string `json:"ScalingInstanceId,omitempty" name:"ScalingInstanceId"`
}

func (r *AttachInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DetachInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId    *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingInstanceId []*string `json:"ScalingInstanceId,omitempty" name:"ScalingInstanceId"`
}

func (r *DetachInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScalingActivityRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId    *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingActivityId []*string `json:"ScalingActivityId,omitempty" name:"ScalingActivityId"`
	Marker            *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults        *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	StartTime         *string   `json:"StartTime,omitempty" name:"StartTime"`
	EndTime           *string   `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeScalingActivityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingActivityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScalingActivityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingActivityResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	AvailabilityZone     *string `json:"AvailabilityZone" name:"AvailabilityZone"`
	ScalingActivityCount *int    `json:"ScalingActivityCount" name:"ScalingActivityCount"`
	ScalingActivitySet   []struct {
		ScalingGroupId    *string   `json:"ScalingGroupId" name:"ScalingGroupId"`
		Status            *int      `json:"Status" name:"Status"`
		Cause             *string   `json:"Cause" name:"Cause"`
		Description       *string   `json:"Description" name:"Description"`
		StartTime         *string   `json:"StartTime" name:"StartTime"`
		ScalingActivityId *string   `json:"ScalingActivityId" name:"ScalingActivityId"`
		EndTime           *string   `json:"EndTime" name:"EndTime"`
		Type              *int      `json:"Type" name:"Type"`
		SuccInsList       []*string `json:"SuccInsList" name:"SuccInsList"`
		FailInsList       []*string `json:"FailInsList" name:"FailInsList"`
		ErrorCode         *int      `json:"ErrorCode" name:"ErrorCode"`
	} `json:"ScalingActivitySet"`
}

func (r *DescribeScalingActivityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingActivityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteScalingGroupRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
}

func (r *DeleteScalingGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingGroupId *string `json:"ScalingGroupId" name:"ScalingGroupId"`
	} `json:"ReturnSet"`
}

func (r *DeleteScalingGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DisableScalingGroupRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
}

func (r *DisableScalingGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisableScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisableScalingGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingGroupId *string `json:"ScalingGroupId" name:"ScalingGroupId"`
	} `json:"ReturnSet"`
}

func (r *DisableScalingGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type EnableScalingGroupRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
}

func (r *EnableScalingGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EnableScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableScalingGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingGroupId *string `json:"ScalingGroupId" name:"ScalingGroupId"`
	} `json:"ReturnSet"`
}

func (r *EnableScalingGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScalingNotificationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId        *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingNotificationId []*string `json:"ScalingNotificationId,omitempty" name:"ScalingNotificationId"`
	Marker                *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults            *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeScalingNotificationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScalingNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingNotificationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	ScalingNotificationCount *int    `json:"ScalingNotificationCount" name:"ScalingNotificationCount"`
	ScalingNotificationSet   []struct {
		ScalingNotificationId    *string   `json:"ScalingNotificationId" name:"ScalingNotificationId"`
		ScalingGroupId           *string   `json:"ScalingGroupId" name:"ScalingGroupId"`
		ScalingNotificationTypes []*string `json:"ScalingNotificationTypes" name:"ScalingNotificationTypes"`
		Description              *string   `json:"Description" name:"Description"`
		AvailabilityZone         *string   `json:"AvailabilityZone" name:"AvailabilityZone"`
	} `json:"ScalingNotificationSet"`
}

func (r *DescribeScalingNotificationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateScalingNotificationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingNotificationType []*string `json:"ScalingNotificationType,omitempty" name:"ScalingNotificationType"`
	ScalingGroupId          *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	Description             *string   `json:"Description,omitempty" name:"Description"`
}

func (r *CreateScalingNotificationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateScalingNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScalingNotificationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	ScalingNotificationId *string `json:"ScalingNotificationId" name:"ScalingNotificationId"`
}

func (r *CreateScalingNotificationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyScalingNotificationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId        *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingNotificationId *int      `json:"ScalingNotificationId,omitempty" name:"ScalingNotificationId"`
	NotificationType      []*string `json:"NotificationType,omitempty" name:"NotificationType"`
}

func (r *ModifyScalingNotificationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyScalingNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingNotificationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	Return                *bool   `json:"Return" name:"Return"`
	ScalingNotificationId *string `json:"ScalingNotificationId" name:"ScalingNotificationId"`
}

func (r *ModifyScalingNotificationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateScheduledTaskRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId           *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingScheduledTaskName *string `json:"ScalingScheduledTaskName,omitempty" name:"ScalingScheduledTaskName"`
	ReadjustMinSize          *int    `json:"ReadjustMinSize,omitempty" name:"ReadjustMinSize"`
	ReadjustMaxSize          *int    `json:"ReadjustMaxSize,omitempty" name:"ReadjustMaxSize"`
	ReadjustExpectSize       *int    `json:"ReadjustExpectSize,omitempty" name:"ReadjustExpectSize"`
	StartTime                *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime                  *string `json:"EndTime,omitempty" name:"EndTime"`
	Recurrence               *string `json:"Recurrence,omitempty" name:"Recurrence"`
	RepeatUnit               *string `json:"RepeatUnit,omitempty" name:"RepeatUnit"`
	RepeatCycle              *string `json:"RepeatCycle,omitempty" name:"RepeatCycle"`
}

func (r *CreateScheduledTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScheduledTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateScheduledTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingScheduleTaskId *string `json:"ScalingScheduleTaskId" name:"ScalingScheduleTaskId"`
	} `json:"ReturnSet"`
}

func (r *CreateScheduledTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScheduledTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScheduledTaskRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId           *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingScheduledTaskId   []*string `json:"ScalingScheduledTaskId,omitempty" name:"ScalingScheduledTaskId"`
	ScalingScheduledTaskName *string   `json:"ScalingScheduledTaskName,omitempty" name:"ScalingScheduledTaskName"`
	Marker                   *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults               *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeScheduledTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScheduledTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScheduledTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	ScheduledTaskSize      *int    `json:"ScheduledTaskSize" name:"ScheduledTaskSize"`
	ScalingScheduleTaskSet []struct {
		ScalingGroupId         *string `json:"ScalingGroupId" name:"ScalingGroupId"`
		ScalingScheduledTaskId *string `json:"ScalingScheduledTaskId" name:"ScalingScheduledTaskId"`
		ScalingScheduledTaskName *string `json:"ScalingScheduledTaskName" name:"ScalingScheduledTaskName"`
		ReadjustMaxSize        *int    `json:"ReadjustMaxSize" name:"ReadjustMaxSize"`
		ReadjustMinSize        *int    `json:"ReadjustMinSize" name:"ReadjustMinSize"`
		ReadjustExpectSize     *int    `json:"ReadjustExpectSize" name:"ReadjustExpectSize"`
		StartTime              *string `json:"StartTime" name:"StartTime"`
		Recurrence             *string `json:"Recurrence" name:"Recurrence"`
		RepeatUnit             *string `json:"RepeatUnit" name:"RepeatUnit"`
		RepeatCycle            *string `json:"RepeatCycle" name:"RepeatCycle"`
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		EndTime                *string `json:"EndTime" name:"EndTime"`
		AvailabilityZone       *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Description            *string `json:"Description" name:"Description"`
	} `json:"ScalingScheduleTaskSet"`
}

func (r *DescribeScheduledTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScheduledTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyScheduledTaskRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ModifyScheduledTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScheduledTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyScheduledTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
	ReturnSet struct {
		ScalingScheduleTaskId *string `json:"ScalingScheduleTaskId" name:"ScalingScheduleTaskId"`
	} `json:"ReturnSet"`
}

func (r *ModifyScheduledTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScheduledTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteScheduledTaskRequest struct {
	*ksyunhttp.BaseRequest
	ScalingScheduledTaskId *string `json:"ScalingScheduledTaskId,omitempty" name:"ScalingScheduledTaskId"`
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
}

func (r *DeleteScheduledTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScheduledTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteScheduledTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScheduledTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingScheduleTaskId *string `json:"ScalingScheduleTaskId" name:"ScalingScheduleTaskId"`
		Message *string `json:"Message" name:"Message"`
	} `json:"ReturnSet"`
}

func (r *DeleteScheduledTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScheduledTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateScalingPolicyRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId    *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`
	Metric            *string `json:"Metric,omitempty" name:"Metric"`
	AdjustmentType    *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`
	AdjustmentValue   *int    `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`
	CoolDown          *int    `json:"CoolDown,omitempty" name:"CoolDown"`
}

func (r *CreateScalingPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScalingPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingPolicyId *string `json:"ScalingPolicyId" name:"ScalingPolicyId"`
	} `json:"ReturnSet"`
}

func (r *CreateScalingPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeScalingPolicyRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId    *string   `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingPolicyId   []*string `json:"ScalingPolicyId,omitempty" name:"ScalingPolicyId"`
	ScalingPolicyName *string   `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`
	Marker            *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults        *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeScalingPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	ScalingPolicySize *int    `json:"ScalingPolicySize" name:"ScalingPolicySize"`
	ScalingPolicySet  []struct {
		ScalingGroupId  *string `json:"ScalingGroupId" name:"ScalingGroupId"`
		ScalingPolicyId *string `json:"ScalingPolicyId" name:"ScalingPolicyId"`
		Description     *string `json:"Description" name:"Description"`
		ScalingPolicyName *string `json:"ScalingPolicyName" name:"ScalingPolicyName"`
		AdjustmentType  *string `json:"AdjustmentType" name:"AdjustmentType"`
		AdjustmentValue *int    `json:"AdjustmentValue" name:"AdjustmentValue"`
		Metric          struct {
			DimensionName      *string `json:"DimensionName" name:"DimensionName"`
			ComparisonOperator *string `json:"ComparisonOperator" name:"ComparisonOperator"`
			Threshold          *int    `json:"Threshold" name:"Threshold"`
			RepeatTimes        *int    `json:"RepeatTimes" name:"RepeatTimes"`
			Period             *int    `json:"Period" name:"Period"`
			Function           *string `json:"Function" name:"Function"`
		} `json:"Metric" name:"Metric"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		CoolDown   *int    `json:"CoolDown" name:"CoolDown"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
	} `json:"ScalingPolicySet"`
}

func (r *DescribeScalingPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyScalingPolicyRequest struct {
	*ksyunhttp.BaseRequest
	ScalingPolicyId   *string `json:"ScalingPolicyId,omitempty" name:"ScalingPolicyId"`
	ScalingGroupId    *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`
	Metric            *string `json:"Metric,omitempty" name:"Metric"`
	AdjustmentType    *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`
	AdjustmentValue   *int    `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`
	CoolDown          *int    `json:"CoolDown,omitempty" name:"CoolDown"`
}

func (r *ModifyScalingPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *string `json:"Return" name:"Return"`
	ReturnSet struct {
		ScalingPolicyId *string `json:"ScalingPolicyId" name:"ScalingPolicyId"`
	} `json:"ReturnSet"`
}

func (r *ModifyScalingPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteScalingPolicyRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
	ScalingPolicyId *string `json:"ScalingPolicyId,omitempty" name:"ScalingPolicyId"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet struct {
		ScalingPolicyId *string `json:"ScalingPolicyId" name:"ScalingPolicyId"`
	} `json:"ReturnSet"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ImportImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageName       *string   `json:"ImageName,omitempty" name:"ImageName"`
	Architecture    *string   `json:"Architecture,omitempty" name:"Architecture"`
	Platform        *string   `json:"Platform,omitempty" name:"Platform"`
	ImageUrl        *string   `json:"ImageUrl,omitempty" name:"ImageUrl"`
	ImageFormat     *string   `json:"ImageFormat,omitempty" name:"ImageFormat"`
	DataImageUrl    []*string `json:"DataImageUrl,omitempty" name:"DataImageUrl"`
	DataImageSize   []*string `json:"DataImageSize,omitempty" name:"DataImageSize"`
	DataImageFormat []*string `json:"DataImageFormat,omitempty" name:"DataImageFormat"`
}

func (r *ImportImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImportImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageId   *string `json:"ImageId" name:"ImageId"`
}

func (r *ImportImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CopyImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId              []*string `json:"ImageId,omitempty" name:"ImageId"`
	DestinationRegion    []*string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	DestinationImageName *string   `json:"DestinationImageName,omitempty" name:"DestinationImageName"`
}

func (r *CopyImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CopyImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CopyImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *CopyImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyImageSharePermissionRequest struct {
	*ksyunhttp.BaseRequest
	ImageId    *string   `json:"ImageId,omitempty" name:"ImageId"`
	AccountId  []*string `json:"AccountId,omitempty" name:"AccountId"`
	Permission *string   `json:"Permission,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionResponse struct {
	*ksyunhttp.BaseResponse
	Return    *bool   `json:"Return" name:"Return"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeImageSharePermissionRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	SharePermissionSet []struct {
		AccountId *string `json:"AccountId" name:"AccountId"`
	} `json:"SharePermissionSet"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
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
		Region *string `json:"Region" name:"Region"`
		RegionName *string `json:"RegionName" name:"RegionName"`
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


type AttachKeyRequest struct {
	*ksyunhttp.BaseRequest
	Action     *string   `json:"Action,omitempty" name:"Action"`
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	KeyId      []*string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *AttachKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Message *string `json:"Message" name:"Message"`
		Return  *bool   `json:"Return" name:"Return"`
	} `json:"InstancesSet"`
}

func (r *AttachKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DetachKeyRequest struct {
	*ksyunhttp.BaseRequest
	Action     *string   `json:"Action,omitempty" name:"Action"`
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	KeyId      []*string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DetachKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	InstancesSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Message *string `json:"Message" name:"Message"`
		Return  *bool   `json:"Return" name:"Return"`
	} `json:"InstancesSet"`
}

func (r *DetachKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachKeyResponse) FromJsonString(s string) error {
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
	RequestId           *string `json:"RequestId" name:"RequestId"`
	AvailabilityZoneSet []struct {
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Region *string `json:"Region" name:"Region"`
	} `json:"AvailabilityZoneSet"`
}

func (r *DescribeAvailabilityZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceTypeConfigsRequest struct {
	*ksyunhttp.BaseRequest
	Filter []*DescribeInstanceTypeConfigsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceTypeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	InstanceTypeConfigSet []struct {
		InstanceType          *string `json:"InstanceType" name:"InstanceType"`
		InstanceFamily        *string `json:"InstanceFamily" name:"InstanceFamily"`
		InstanceFamilyName    *string `json:"InstanceFamilyName" name:"InstanceFamilyName"`
		CPU                   *int    `json:"CPU" name:"CPU"`
		GPU                   *int    `json:"GPU" name:"GPU"`
		Memory                *int    `json:"Memory" name:"Memory"`
		NetworkInterfaceQuota struct {
			NetworkInterfaceCount *int `json:"NetworkInterfaceCount" name:"NetworkInterfaceCount"`
		} `json:"NetworkInterfaceQuota" name:"NetworkInterfaceQuota"`
		PrivateIpQuota struct {
			PrivateIpCount *int `json:"PrivateIpCount" name:"PrivateIpCount"`
		} `json:"PrivateIpQuota" name:"PrivateIpQuota"`
		AvailabilityZoneSet []struct {
			AzCode *string `json:"AzCode" name:"AzCode"`
		} `json:"AvailabilityZoneSet" name:"AvailabilityZoneSet"`
		SystemDiskQuotaSet []struct {
			SystemDiskType *string `json:"SystemDiskType" name:"SystemDiskType"`
		} `json:"SystemDiskQuotaSet" name:"SystemDiskQuotaSet"`
		DataDiskQuotaSet []struct {
			DataDiskType        *string  `json:"DataDiskType" name:"DataDiskType"`
			DataDiskMinSize     *float64 `json:"DataDiskMinSize" name:"DataDiskMinSize"`
			DataDiskMaxsize     *float64 `json:"DataDiskMaxsize" name:"DataDiskMaxsize"`
			DataDiskCount       *int     `json:"DataDiskCount" name:"DataDiskCount"`
			AvailabilityZoneSet []struct {
				AzCode    *string `json:"AzCode" name:"AzCode"`
				IsSellOut *bool   `json:"IsSellOut" name:"IsSellOut"`
			} `json:"AvailabilityZoneSet"`
		} `json:"DataDiskQuotaSet" name:"DataDiskQuotaSet"`
	} `json:"InstanceTypeConfigSet"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceFamilysRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeInstanceFamilysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceFamilysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceFamilysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	InstanceFamilySet []struct {
		AvailabilityZoneSet []struct {
			AzCode *string `json:"AzCode" name:"AzCode"`
		} `json:"AvailabilityZoneSet" name:"AvailabilityZoneSet"`
		InstanceFamilyName *string `json:"InstanceFamilyName" name:"InstanceFamilyName"`
		InstanceFamily *string `json:"InstanceFamily" name:"InstanceFamily"`
	} `json:"InstanceFamilySet"`
}

func (r *DescribeInstanceFamilysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceFamilysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddVmIntoDataGuardRequest struct {
	*ksyunhttp.BaseRequest
	DataGuardId *string   `json:"DataGuardId,omitempty" name:"DataGuardId"`
	InstanceId  []*string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *AddVmIntoDataGuardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVmIntoDataGuardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddVmIntoDataGuardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddVmIntoDataGuardResponse struct {
	*ksyunhttp.BaseResponse
	DataGuardResult struct {
		Flag    *bool   `json:"Flag" name:"Flag"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Message *string `json:"Message" name:"Message"`
	} `json:"DataGuardResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AddVmIntoDataGuardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVmIntoDataGuardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	VpcId            *string `json:"VpcId,omitempty" name:"VpcId"`
	StorageType      *string `json:"StorageType,omitempty" name:"StorageType"`
	ProtocolType     *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
	FileSystemName   *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	ProjectId        *int    `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	CreateFileSystemResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		FileSystemId *string `json:"FileSystemId" name:"FileSystemId"`
	} `json:"CreateFileSystemResponse"`
}

func (r *CreateFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	DeleteFileSystemResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		Return *string `json:"Return" name:"Return"`
	} `json:"DeleteFileSystemResponse"`
}

func (r *DeleteFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeFileSystemsRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId []*string                    `json:"FileSystemId,omitempty" name:"FileSystemId"`
	MaxResults   *int                         `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker       *int                         `json:"Marker,omitempty" name:"Marker"`
	ProjectId    []*int                       `json:"ProjectId,omitempty" name:"ProjectId"`
	IncludeDel   *bool                        `json:"IncludeDel,omitempty" name:"IncludeDel"`
	Filter       []*DescribeFileSystemsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeFileSystemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	Marker          *int    `json:"Marker" name:"Marker"`
	FileSystemCount *int    `json:"FileSystemCount" name:"FileSystemCount"`
	FileSystems     []struct {
		FileSystemId     *string `json:"FileSystemId" name:"FileSystemId"`
		FileSystemName   *string `json:"FileSystemName" name:"FileSystemName"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		AvailabilityZoneName *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		StorageType      *string `json:"StorageType" name:"StorageType"`
		ProtocolType     *string `json:"ProtocolType" name:"ProtocolType"`
		VpcId            *string `json:"VpcId" name:"VpcId"`
		FileSystemState  *string `json:"FileSystemState" name:"FileSystemState"`
		CreationDate     *string `json:"CreationDate" name:"CreationDate"`
		MountTargets     []struct {
			MountTargetId    *string `json:"MountTargetId" name:"MountTargetId"`
			SubnetId         *string `json:"SubnetId" name:"SubnetId"`
			IpAddress        *string `json:"IpAddress" name:"IpAddress"`
			MountTargetState *string `json:"MountTargetState" name:"MountTargetState"`
			CreationDate     *string `json:"CreationDate" name:"CreationDate"`
		} `json:"MountTargets" name:"MountTargets"`
		Size *int `json:"Size" name:"Size"`
		UsedSize *int `json:"UsedSize" name:"UsedSize"`
	} `json:"FileSystems"`
}

func (r *DescribeFileSystemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyFileSystemRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId   *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
}

func (r *ModifyFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileSystemResponse struct {
	*ksyunhttp.BaseResponse
	ModifyFileSystemResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		Return *string `json:"Return" name:"Return"`
	} `json:"ModifyFileSystemResponse"`
}

func (r *ModifyFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateMountTargetRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	IpVersion    *string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *CreateMountTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMountTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateMountTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMountTargetResponse struct {
	*ksyunhttp.BaseResponse
	CreateMountTargetResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		MountTargetId *string `json:"MountTargetId" name:"MountTargetId"`
		IpAddress *string `json:"IpAddress" name:"IpAddress"`
	} `json:"CreateMountTargetResponse"`
}

func (r *CreateMountTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMountTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteMountTargetRequest struct {
	*ksyunhttp.BaseRequest
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
}

func (r *DeleteMountTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMountTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteMountTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMountTargetResponse struct {
	*ksyunhttp.BaseResponse
	DeleteMountTargetResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		Return *string `json:"Return" name:"Return"`
	} `json:"DeleteMountTargetResponse"`
}

func (r *DeleteMountTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMountTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMountTargetsRequest struct {
	*ksyunhttp.BaseRequest
	FileSystemId  *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
	MaxResults    *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker        *int    `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeMountTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMountTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMountTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMountTargetsResponse struct {
	*ksyunhttp.BaseResponse
	DescribeMountTargetsResponse struct {
		RequestId    *string `json:"RequestId" name:"RequestId"`
		Marker       *string `json:"Marker" name:"Marker"`
		FileSystemCount *string `json:"FileSystemCount" name:"FileSystemCount"`
		MountTargets struct {
			Item struct {
				MountTargetId *string `json:"MountTargetId" name:"MountTargetId"`
				SubnetId      *string `json:"SubnetId" name:"SubnetId"`
				IpAddress     *string `json:"IpAddress" name:"IpAddress"`
				MountTargetState *string `json:"MountTargetState" name:"MountTargetState"`
				CreationDate  *string `json:"CreationDate" name:"CreationDate"`
			} `json:"Item"`
		} `json:"MountTargets" name:"MountTargets"`
	} `json:"DescribeMountTargetsResponse"`
}

func (r *DescribeMountTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMountTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateModelRequest struct {
	*ksyunhttp.BaseRequest
	ImageId                    *string                        `json:"ImageId,omitempty" name:"ImageId"`
	InstanceType               *string                        `json:"InstanceType,omitempty" name:"InstanceType"`
	DataDiskGb                 *int                           `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	SubnetId                   *string                        `json:"SubnetId,omitempty" name:"SubnetId"`
	DataDisk                   []*CreateModelDataDisk         `json:"DataDisk,omitempty" name:"DataDisk"`
	KeepImageLogin             *bool                          `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
	KeyId                      *string                        `json:"KeyId,omitempty" name:"KeyId"`
	ChargeType                 *string                        `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime               *int                           `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	SecurityGroupId            []*string                      `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	PrivateIpAddress           *string                        `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	InstanceName               *string                        `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix         *string                        `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	SriovNetSupport            *string                        `json:"SriovNetSupport,omitempty" name:"SriovNetSupport"`
	ProjectId                  *int                           `json:"ProjectId,omitempty" name:"ProjectId"`
	DataGuardId                *string                        `json:"DataGuardId,omitempty" name:"DataGuardId"`
	AddressBandWidth           *int                           `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
	LineId                     *string                        `json:"LineId,omitempty" name:"LineId"`
	AddressChargeType          *string                        `json:"AddressChargeType,omitempty" name:"AddressChargeType"`
	AddressPurchaseTime        *int                           `json:"AddressPurchaseTime,omitempty" name:"AddressPurchaseTime"`
	AddressProjectId           *string                        `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
	ModelName                  *string                        `json:"ModelName,omitempty" name:"ModelName"`
	FailureAutoDelete          *bool                          `json:"FailureAutoDelete,omitempty" name:"FailureAutoDelete"`
	SystemDisk                 *CreateModelSystemDisk         `json:"SystemDisk,omitempty" name:"SystemDisk"`
	HostName                   *string                        `json:"HostName,omitempty" name:"HostName"`
	HostNameSuffix             *string                        `json:"HostNameSuffix,omitempty" name:"HostNameSuffix"`
	UserData                   *string                        `json:"UserData,omitempty" name:"UserData"`
	NetworkInterface           []*CreateModelNetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`
	Tag                        []*CreateModelTag              `json:"Tag,omitempty" name:"Tag"`
	AllocateAddress            *bool                          `json:"AllocateAddress,omitempty" name:"AllocateAddress"`
	IsDistributeIpv6           *bool                          `json:"IsDistributeIpv6,omitempty" name:"IsDistributeIpv6"`
	Mem                        *string                        `json:"Mem,omitempty" name:"Mem"`
	Cpu                        *string                        `json:"Cpu,omitempty" name:"Cpu"`
	IamRoleName                *string                        `json:"IamRoleName,omitempty" name:"IamRoleName"`
	AssembledImageDataDiskType *string                        `json:"AssembledImageDataDiskType,omitempty" name:"AssembledImageDataDiskType"`
	LocalVolumeSnapshotId      *string                        `json:"LocalVolumeSnapshotId,omitempty" name:"LocalVolumeSnapshotId"`
	SyncTag                    *bool                          `json:"SyncTag,omitempty" name:"SyncTag"`
}

func (r *CreateModelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateModelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ModelId   *string `json:"ModelId" name:"ModelId"`
	ModelName *string `json:"ModelName" name:"ModelName"`
}

func (r *CreateModelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type TerminateModelsRequest struct {
	*ksyunhttp.BaseRequest
	ModelId      *string `json:"ModelId,omitempty" name:"ModelId"`
	ModelVersion *int    `json:"ModelVersion,omitempty" name:"ModelVersion"`
}

func (r *TerminateModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "TerminateModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TerminateModelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *TerminateModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeModelsRequest struct {
	*ksyunhttp.BaseRequest
	ModelId    []*string `json:"ModelId,omitempty" name:"ModelId"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int      `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
	ModelSet      []struct {
		ModelName        *string `json:"ModelName" name:"ModelName"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		AvailabilityZoneName *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		ModelId          *string `json:"ModelId" name:"ModelId"`
		ModifiedTime     *string `json:"ModifiedTime" name:"ModifiedTime"`
		ImageId          *string `json:"ImageId" name:"ImageId"`
		InstanceType     *string `json:"InstanceType" name:"InstanceType"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		ChargeType       *string `json:"ChargeType" name:"ChargeType"`
		PurchaseTime     *string `json:"PurchaseTime" name:"PurchaseTime"`
		SecurityGroupId  *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		InstanceName     *string `json:"InstanceName" name:"InstanceName"`
		SriovNetSupport  *string `json:"SriovNetSupport" name:"SriovNetSupport"`
		KeepImageLogin   *string `json:"KeepImageLogin" name:"KeepImageLogin"`
		SystemDisk       struct {
			DiskType *string `json:"DiskType" name:"DiskType"`
			DiskSize *string `json:"DiskSize" name:"DiskSize"`
		} `json:"SystemDisk" name:"SystemDisk"`
		FailureAutoDelete *bool `json:"FailureAutoDelete" name:"FailureAutoDelete"`
		SyncTag *bool `json:"SyncTag" name:"SyncTag"`
	} `json:"ModelSet"`
}

func (r *DescribeModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDedicatedClusterRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedClusterId []*string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	DedicatedClustersSet struct {
		DedicatedClusterId *string `json:"DedicatedClusterId" name:"DedicatedClusterId"`
		DedicatedClusterName *string `json:"DedicatedClusterName" name:"DedicatedClusterName"`
		CreationDate       *string `json:"CreationDate" name:"CreationDate"`
		AvailabilityZone   *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		AvailabilityZoneName *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		Model              *string `json:"Model" name:"Model"`
		DedicatedHostsSet  []struct {
			DedicatedHostId   *string `json:"DedicatedHostId" name:"DedicatedHostId"`
			DedicatedHostName *string `json:"DedicatedHostName" name:"DedicatedHostName"`
		} `json:"DedicatedHostsSet" name:"DedicatedHostsSet"`
	} `json:"DedicatedClustersSet"`
}

func (r *DescribeDedicatedClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateDedicatedClusterRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedClusterName *string `json:"DedicatedClusterName,omitempty" name:"DedicatedClusterName"`
	Model                *string `json:"Model,omitempty" name:"Model"`
	AvailabilityZone     *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
}

func (r *CreateDedicatedClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDedicatedClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	DedicatedClusterId *string `json:"DedicatedClusterId" name:"DedicatedClusterId"`
}

func (r *CreateDedicatedClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteDedicatedClusterRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedClusterId []*string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *DeleteDedicatedClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDedicatedClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDedicatedClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDedicatedClusterResponse struct {
	*ksyunhttp.BaseResponse
	DedicatedClustersSet []struct {
		DedicatedClusterId *string `json:"DedicatedClusterId" name:"DedicatedClusterId"`
		Return  *string `json:"Return" name:"Return"`
		Message *string `json:"Message" name:"Message"`
	} `json:"DedicatedClustersSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDedicatedClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDedicatedClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetvCPURequest struct {
	*ksyunhttp.BaseRequest
	DedicatedHostId []*string `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
	VCPU            *int      `json:"VCPU,omitempty" name:"VCPU"`
}

func (r *SetvCPURequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetvCPURequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetvCPURequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetvCPUResponse struct {
	*ksyunhttp.BaseResponse
	ReturnSet []struct {
		DedicatedHostId *string `json:"DedicatedHostId" name:"DedicatedHostId"`
		Return  *bool   `json:"Return" name:"Return"`
		Message *string `json:"Message" name:"Message"`
	} `json:"ReturnSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetvCPUResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetvCPUResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DedicatedHostMigrateRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedClusterId *string   `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	DedicatedHostId    []*string `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
}

func (r *DedicatedHostMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DedicatedHostMigrateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DedicatedHostMigrateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DedicatedHostMigrateResponse struct {
	*ksyunhttp.BaseResponse
	ReturnSet []struct {
		DedicatedHostId *string `json:"DedicatedHostId" name:"DedicatedHostId"`
		Return  *string `json:"Return" name:"Return"`
		Message *string `json:"Message" name:"Message"`
	} `json:"ReturnSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DedicatedHostMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DedicatedHostMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyDedicatedClusterNameRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedClusterId   *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	DedicatedClusterName *string `json:"DedicatedClusterName,omitempty" name:"DedicatedClusterName"`
}

func (r *ModifyDedicatedClusterNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDedicatedClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDedicatedClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDedicatedClusterNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *string `json:"Return" name:"Return"`
}

func (r *ModifyDedicatedClusterNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDedicatedClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type InstanceMigrateRequest struct {
	*ksyunhttp.BaseRequest
	DedicatedHostId *string   `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
	InstanceId      *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceType    *string   `json:"InstanceType,omitempty" name:"InstanceType"`
	DataDisk        []*string `json:"DataDisk,omitempty" name:"DataDisk"`
}

func (r *InstanceMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstanceMigrateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "InstanceMigrateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InstanceMigrateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *InstanceMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstanceMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyInstanceAutoDeleteTimeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId     []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	AutoDeleteTime *string   `json:"AutoDeleteTime,omitempty" name:"AutoDeleteTime"`
	AutoDeleteEip  *bool     `json:"AutoDeleteEip,omitempty" name:"AutoDeleteEip"`
}

func (r *ModifyInstanceAutoDeleteTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAutoDeleteTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceAutoDeleteTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAutoDeleteTimeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyInstanceAutoDeleteTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAutoDeleteTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyScalingConfigurationRequest struct {
	*ksyunhttp.BaseRequest
	ScalingConfigurationId   *string                               `json:"ScalingConfigurationId,omitempty" name:"ScalingConfigurationId"`
	ScalingConfigurationName *string                               `json:"ScalingConfigurationName,omitempty" name:"ScalingConfigurationName"`
	ImageId                  *string                               `json:"ImageId,omitempty" name:"ImageId"`
	Password                 *string                               `json:"Password,omitempty" name:"Password"`
	InstanceType             *string                               `json:"InstanceType,omitempty" name:"InstanceType"`
	ChargeType               *string                               `json:"ChargeType,omitempty" name:"ChargeType"`
	DataDiskGb               *int                                  `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	ProjectId                *int                                  `json:"ProjectId,omitempty" name:"ProjectId"`
	KeepImageLogin           *bool                                 `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
	KeyId                    []*string                             `json:"KeyId,omitempty" name:"KeyId"`
	DataDisk                 []*ModifyScalingConfigurationDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	SystemDiskDiskSize       *int                                  `json:"SystemDisk.DiskSize,omitempty" name:"SystemDisk.DiskSize"`
	AddressBandWidth         *int                                  `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
	BandWidthShareId         *string                               `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	LineId                   *string                               `json:"LineId,omitempty" name:"LineId"`
	AddressProjectId         *int                                  `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
	InstanceName             *string                               `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix       *string                               `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	UserData                 *string                               `json:"UserData,omitempty" name:"UserData"`
	InstanceNameTimeSuffix   *bool                                 `json:"InstanceNameTimeSuffix,omitempty" name:"InstanceNameTimeSuffix"`
	Tag                      []*ModifyScalingConfigurationTag      `json:"Tag,omitempty" name:"Tag"`
	LoginSetAfter            *bool                                 `json:"LoginSetAfter,omitempty" name:"LoginSetAfter"`
	IpBindAfter              *bool                                 `json:"IpBindAfter,omitempty" name:"IpBindAfter"`
	InstanceNameRandom       *bool                                 `json:"InstanceNameRandom,omitempty" name:"InstanceNameRandom"`
	SystemDiskDiskType       *string                               `json:"SystemDisk.DiskType,omitempty" name:"SystemDisk.DiskType"`
	SystemDiskResizeType     *string                               `json:"SystemDisk.ResizeType,omitempty" name:"SystemDisk.ResizeType"`
	SyncTag                  *bool                                 `json:"SyncTag,omitempty" name:"SyncTag"`
}

func (r *ModifyScalingConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingConfigurationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyScalingConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSpotPriceHistoryRequest struct {
	*ksyunhttp.BaseRequest
	InstanceType     *string `json:"InstanceType,omitempty" name:"InstanceType"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	StartTime        *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime          *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSpotPriceHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpotPriceHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSpotPriceHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpotPriceHistoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	InstanceType     *string `json:"InstanceType" name:"InstanceType"`
	AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
	SpotPrices       []struct {
		SpotPrice *float64 `json:"SpotPrice" name:"SpotPrice"`
		OriginPrice *float64 `json:"OriginPrice" name:"OriginPrice"`
		Timestamp *string  `json:"Timestamp" name:"Timestamp"`
	} `json:"SpotPrices"`
}

func (r *DescribeSpotPriceHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpotPriceHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceType       *string                  `json:"InstanceType,omitempty" name:"InstanceType"`
	SystemDiskDiskSize *string                  `json:"SystemDisk.DiskSize,omitempty" name:"SystemDisk.DiskSize"`
	ImageId            *string                  `json:"ImageId,omitempty" name:"ImageId"`
	ChargeType         *string                  `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime       *int                     `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	DataDiskGb         *int                     `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	DataDisk           []*DescribePriceDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	MaxCount           *int                     `json:"MaxCount,omitempty" name:"MaxCount"`
	SystemDiskDiskType *string                  `json:"SystemDisk.DiskType,omitempty" name:"SystemDisk.DiskType"`
}

func (r *DescribePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	PriceInfo struct {
		InstancePrice struct {
			DiscountPrice *float64 `json:"DiscountPrice" name:"DiscountPrice"`
			OriginalPrice *float64 `json:"OriginalPrice" name:"OriginalPrice"`
			TradePrice    *float64 `json:"TradePrice" name:"TradePrice"`
			PriceUnit     *string  `json:"PriceUnit" name:"PriceUnit"`
			UserDiscount  *float64 `json:"UserDiscount" name:"UserDiscount"`
		} `json:"InstancePrice" name:"InstancePrice"`
		EBSPrice struct {
			DiscountPrice *float64 `json:"DiscountPrice" name:"DiscountPrice"`
			OriginalPrice *float64 `json:"OriginalPrice" name:"OriginalPrice"`
			TradePrice    *float64 `json:"TradePrice" name:"TradePrice"`
			PriceUnit     *string  `json:"PriceUnit" name:"PriceUnit"`
			UserDiscount  *float64 `json:"UserDiscount" name:"UserDiscount"`
		} `json:"EBSPrice" name:"EBSPrice"`
	} `json:"PriceInfo"`
}

func (r *DescribePriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableImageCachingRequest struct {
	*ksyunhttp.BaseRequest
	ImageId []*string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *EnableImageCachingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableImageCachingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EnableImageCachingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableImageCachingResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *EnableImageCachingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableImageCachingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableImageCachingRequest struct {
	*ksyunhttp.BaseRequest
	ImageId []*string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DisableImageCachingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableImageCachingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisableImageCachingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisableImageCachingResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *string `json:"Return" name:"Return"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *DisableImageCachingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableImageCachingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancersRequest struct {
	*ksyunhttp.BaseRequest
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" name:"ScalingGroupId"`
}

func (r *ModifyLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancersResponse struct {
	*ksyunhttp.BaseResponse
	ModifyLoadBalancersResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		Return *string `json:"Return" name:"Return"`
	} `json:"ModifyLoadBalancersResponse"`
}

func (r *ModifyLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachInstancesIamRoleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	IamRoleName *string   `json:"IamRoleName,omitempty" name:"IamRoleName"`
}

func (r *AttachInstancesIamRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachInstancesIamRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachInstancesIamRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachInstancesIamRoleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	FailCount   *int    `json:"FailCount" name:"FailCount"`
	TotalCount  *int    `json:"TotalCount" name:"TotalCount"`
	IamRoleName *string `json:"IamRoleName" name:"IamRoleName"`
	Results     struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Result *string `json:"Result" name:"Result"`
	} `json:"Results"`
}

func (r *AttachInstancesIamRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachInstancesIamRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachInstancesIamRoleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachInstancesIamRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachInstancesIamRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachInstancesIamRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachInstancesIamRoleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	FailCount  *int    `json:"FailCount" name:"FailCount"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Results    struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Result *bool `json:"Result" name:"Result"`
	} `json:"Results"`
}

func (r *DetachInstancesIamRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachInstancesIamRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopySnapshotRequest struct {
	*ksyunhttp.BaseRequest
	SnapshotId              []*string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	DestinationRegion       []*string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	DestinationSnapshotName *string   `json:"DestinationSnapshotName,omitempty" name:"DestinationSnapshotName"`
	DestinationSnapshotDesc *string   `json:"DestinationSnapshotDesc,omitempty" name:"DestinationSnapshotDesc"`
}

func (r *CopySnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopySnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CopySnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CopySnapshotResponse struct {
	*ksyunhttp.BaseResponse
	Return    *bool   `json:"Return" name:"Return"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CopySnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopySnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreMigrateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId     *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceType   *string `json:"InstanceType,omitempty" name:"InstanceType"`
	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	DataDiskType   *string `json:"DataDiskType,omitempty" name:"DataDiskType"`
	InstantAccess  *bool   `json:"InstantAccess,omitempty" name:"InstantAccess"`
}

func (r *PreMigrateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreMigrateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "PreMigrateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PreMigrateInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *PreMigrateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreMigrateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelPreMigrateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CancelPreMigrateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelPreMigrateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelPreMigrateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelPreMigrateInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CancelPreMigrateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelPreMigrateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVNCAddressRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetVNCAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVNCAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetVNCAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetVNCAddressResponse struct {
	*ksyunhttp.BaseResponse
	VNCAddress struct {
		Port *string `json:"Port" name:"Port"`
		Host *string `json:"Host" name:"Host"`
	} `json:"VNCAddress"`
	Cookies []struct {
		CookieKey *string `json:"CookieKey" name:"CookieKey"`
		CookieValue *string `json:"CookieValue" name:"CookieValue"`
	} `json:"Cookies"`
	Domain    *string `json:"Domain" name:"Domain"`
	VNCPort   *string `json:"VNCPort" name:"VNCPort"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetVNCAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVNCAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchImageTypeRequest struct {
	*ksyunhttp.BaseRequest
	ImageId []*string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *SwitchImageTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchImageTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SwitchImageTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchImageTypeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId           *string `json:"RequestId" name:"RequestId"`
	ImageErrorResponses []struct {
		Result *bool   `json:"Result" name:"Result"`
		Reason *string `json:"Reason" name:"Reason"`
		ImageId *string `json:"ImageId" name:"ImageId"`
	} `json:"ImageErrorResponses"`
}

func (r *SwitchImageTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchImageTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetInstanceResourceProtectRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	IsProtection *bool     `json:"IsProtection,omitempty" name:"IsProtection"`
}

func (r *SetInstanceResourceProtectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetInstanceResourceProtectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetInstanceResourceProtectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetInstanceResourceProtectResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetInstanceResourceProtectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetInstanceResourceProtectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

