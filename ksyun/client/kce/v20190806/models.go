package v20190806

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeClusterInstanceFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyClusterInfoControlPlaneLog struct {
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Enable      *bool   `json:"Enable,omitempty" name:"Enable"`
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	Items       *string `json:"Items,omitempty" name:"Items"`
}
type AddClusterInstancesInstanceSetAdvancedSettingDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type AddClusterInstancesInstanceSetAdvancedSettingLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type AddClusterInstancesInstanceSetAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type AddClusterInstancesInstanceSetAdvancedSettingExtraArg struct {
	Kubelet []*AddClusterInstancesInstanceSetAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type AddClusterInstancesInstanceSetAdvancedSettingTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type AddClusterInstancesInstanceSetAdvancedSetting struct {
	DataDisk             *AddClusterInstancesInstanceSetAdvancedSettingDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	ContainerRuntime     *string                                                `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                                `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                                `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                                `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                                  `json:"Schedulable,omitempty" name:"Schedulable"`
	Label                []*AddClusterInstancesInstanceSetAdvancedSettingLabel  `json:"Label,omitempty" name:"Label"`
	ExtraArg             *AddClusterInstancesInstanceSetAdvancedSettingExtraArg `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                                   `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                                   `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
	Taints               []*AddClusterInstancesInstanceSetAdvancedSettingTaints `json:"Taints,omitempty" name:"Taints"`
}
type AddClusterInstancesInstanceSet struct {
	NodeRole        *string                                        `json:"NodeRole,omitempty" name:"NodeRole"`
	NodePara        []*string                                      `json:"NodePara,omitempty" name:"NodePara"`
	AdvancedSetting *AddClusterInstancesInstanceSetAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
}
type DescribeEpcForClusterFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddClusterEpcInstancesAdvancedSettingLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type AddClusterEpcInstancesAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type AddClusterEpcInstancesAdvancedSettingExtraArg struct {
	Kubelet []*AddClusterEpcInstancesAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type AddClusterEpcInstancesAdvancedSettingTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type AddClusterEpcInstancesAdvancedSetting struct {
	ContainerRuntime     *string                                        `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                        `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                        `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                        `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                          `json:"Schedulable,omitempty" name:"Schedulable"`
	Label                []*AddClusterEpcInstancesAdvancedSettingLabel  `json:"Label,omitempty" name:"Label"`
	ExtraArg             *AddClusterEpcInstancesAdvancedSettingExtraArg `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                           `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                           `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
	Taints               []*AddClusterEpcInstancesAdvancedSettingTaints `json:"Taints,omitempty" name:"Taints"`
}
type DescribeExistedInstancesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddExistedInstancesExistedInstanceKecSetAdvancedSettingDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type AddExistedInstancesExistedInstanceKecSetAdvancedSettingLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type AddExistedInstancesExistedInstanceKecSetAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type AddExistedInstancesExistedInstanceKecSetAdvancedSettingExtraArg struct {
	Kubelet []*AddExistedInstancesExistedInstanceKecSetAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type AddExistedInstancesExistedInstanceKecSetAdvancedSettingTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type AddExistedInstancesExistedInstanceKecSetAdvancedSetting struct {
	DataDisk             *AddExistedInstancesExistedInstanceKecSetAdvancedSettingDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	ContainerRuntime     *string                                                          `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                                          `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                                          `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                                          `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                                            `json:"Schedulable,omitempty" name:"Schedulable"`
	Label                []*AddExistedInstancesExistedInstanceKecSetAdvancedSettingLabel  `json:"Label,omitempty" name:"Label"`
	ExtraArg             *AddExistedInstancesExistedInstanceKecSetAdvancedSettingExtraArg `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                                             `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                                             `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
	Taints               []*AddExistedInstancesExistedInstanceKecSetAdvancedSettingTaints `json:"Taints,omitempty" name:"Taints"`
}
type AddExistedInstancesExistedInstanceKecSet struct {
	NodeRole        *string                                                  `json:"NodeRole,omitempty" name:"NodeRole"`
	KecPara         []*string                                                `json:"KecPara,omitempty" name:"KecPara"`
	AdvancedSetting *AddExistedInstancesExistedInstanceKecSetAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
}
type CreateNodePoolNodeTemplateSystemDisk struct {
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	DiskSize *int    `json:"DiskSize,omitempty" name:"DiskSize"`
}
type CreateNodePoolNodeTemplateDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}
type CreateNodePoolNodeTemplateAdvancedSettingDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type CreateNodePoolNodeTemplateAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type CreateNodePoolNodeTemplateAdvancedSettingExtraArg struct {
	Kubelet []*CreateNodePoolNodeTemplateAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type CreateNodePoolNodeTemplateAdvancedSetting struct {
	DataDisk             *CreateNodePoolNodeTemplateAdvancedSettingDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
	ContainerRuntime     *string                                            `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                            `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                            `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                            `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                              `json:"Schedulable,omitempty" name:"Schedulable"`
	ExtraArg             *CreateNodePoolNodeTemplateAdvancedSettingExtraArg `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                               `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                               `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
}
type CreateNodePoolNodeTemplateEbsTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateNodePoolNodeTemplateInstanceTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateNodePoolNodeTemplateMultiDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type CreateNodePoolNodeTemplate struct {
	ChargeType         *string                                    `json:"ChargeType,omitempty" name:"ChargeType"`
	InstanceType       *string                                    `json:"InstanceType,omitempty" name:"InstanceType"`
	SystemDisk         *CreateNodePoolNodeTemplateSystemDisk      `json:"SystemDisk,omitempty" name:"SystemDisk"`
	DataDiskGb         *int                                       `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	DataDisk           []*CreateNodePoolNodeTemplateDataDisk      `json:"DataDisk,omitempty" name:"DataDisk"`
	ImageId            *string                                    `json:"ImageId,omitempty" name:"ImageId"`
	VpcId              *string                                    `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId           []*string                                  `json:"SubnetId,omitempty" name:"SubnetId"`
	SubnetStrategy     *string                                    `json:"SubnetStrategy,omitempty" name:"SubnetStrategy"`
	SecurityGroupId    *string                                    `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	ProjectId          *string                                    `json:"ProjectId,omitempty" name:"ProjectId"`
	Password           *string                                    `json:"Password,omitempty" name:"Password"`
	AdvancedSetting    *CreateNodePoolNodeTemplateAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
	EbsTag             []*CreateNodePoolNodeTemplateEbsTag        `json:"EbsTag,omitempty" name:"EbsTag"`
	InstanceTag        []*CreateNodePoolNodeTemplateInstanceTag   `json:"InstanceTag,omitempty" name:"InstanceTag"`
	Cpu                *string                                    `json:"Cpu,omitempty" name:"Cpu"`
	Mem                *string                                    `json:"Mem,omitempty" name:"Mem"`
	KeyId              []*string                                  `json:"KeyId,omitempty" name:"KeyId"`
	InstanceName       *string                                    `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix *int                                       `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	MultiDataDisk      []*CreateNodePoolNodeTemplateMultiDataDisk `json:"MultiDataDisk,omitempty" name:"MultiDataDisk"`
}
type CreateNodePoolLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateNodePoolTaint struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type ModifyNodePoolLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyNodePoolTaint struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type ModifyNodeTemplateNodeTemplateSystemDisk struct {
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	DiskSize *int    `json:"DiskSize,omitempty" name:"DiskSize"`
}
type ModifyNodeTemplateNodeTemplateDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSettingDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSettingLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSettingExtraArg struct {
	Kubelet []*ModifyNodeTemplateNodeTemplateAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSettingTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSettingMultiDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type ModifyNodeTemplateNodeTemplateAdvancedSetting struct {
	DataDisk             *ModifyNodeTemplateNodeTemplateAdvancedSettingDataDisk        `json:"DataDisk,omitempty" name:"DataDisk"`
	ContainerRuntime     *string                                                       `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                                       `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                                       `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                                       `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                                         `json:"Schedulable,omitempty" name:"Schedulable"`
	Label                []*ModifyNodeTemplateNodeTemplateAdvancedSettingLabel         `json:"Label,omitempty" name:"Label"`
	ExtraArg             *ModifyNodeTemplateNodeTemplateAdvancedSettingExtraArg        `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                                          `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                                          `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
	Taints               []*ModifyNodeTemplateNodeTemplateAdvancedSettingTaints        `json:"Taints,omitempty" name:"Taints"`
	MultiDataDisk        []*ModifyNodeTemplateNodeTemplateAdvancedSettingMultiDataDisk `json:"MultiDataDisk,omitempty" name:"MultiDataDisk"`
}
type ModifyNodeTemplateNodeTemplateEbsTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyNodeTemplateNodeTemplateInstanceTag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyNodeTemplateNodeTemplate struct {
	ChargeType         *string                                        `json:"ChargeType,omitempty" name:"ChargeType"`
	InstanceType       *string                                        `json:"InstanceType,omitempty" name:"InstanceType"`
	InstanceName       *string                                        `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix *int                                           `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	SystemDisk         *ModifyNodeTemplateNodeTemplateSystemDisk      `json:"SystemDisk,omitempty" name:"SystemDisk"`
	DataDiskGb         *int                                           `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	DataDisk           []*ModifyNodeTemplateNodeTemplateDataDisk      `json:"DataDisk,omitempty" name:"DataDisk"`
	ImageId            *string                                        `json:"ImageId,omitempty" name:"ImageId"`
	VpcId              *string                                        `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId           []*string                                      `json:"SubnetId,omitempty" name:"SubnetId"`
	SubnetStrategy     *string                                        `json:"SubnetStrategy,omitempty" name:"SubnetStrategy"`
	SecurityGroupId    *string                                        `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	ProjectId          *string                                        `json:"ProjectId,omitempty" name:"ProjectId"`
	Password           *string                                        `json:"Password,omitempty" name:"Password"`
	KeyId              []*string                                      `json:"KeyId,omitempty" name:"KeyId"`
	AdvancedSetting    *ModifyNodeTemplateNodeTemplateAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
	EbsTag             []*ModifyNodeTemplateNodeTemplateEbsTag        `json:"EbsTag,omitempty" name:"EbsTag"`
	InstanceTag        []*ModifyNodeTemplateNodeTemplateInstanceTag   `json:"InstanceTag,omitempty" name:"InstanceTag"`
	DeleteDataDisk     *bool                                          `json:"DeleteDataDisk,omitempty" name:"DeleteDataDisk"`
	DeleteInstanceTag  *bool                                          `json:"DeleteInstanceTag,omitempty" name:"DeleteInstanceTag"`
	DeleteEbsTag       *bool                                          `json:"DeleteEbsTag,omitempty" name:"DeleteEbsTag"`
	Cpu                *string                                        `json:"Cpu,omitempty" name:"Cpu"`
	Mem                *string                                        `json:"Mem,omitempty" name:"Mem"`
}
type CreateLogRuleInputConfigContainerStandoutSpecifiedWorkloadNamespaceListWorkload struct {
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
}
type CreateLogRuleInputConfigContainerStandoutSpecifiedWorkloadNamespaceList struct {
	NamespaceName *string                                                                            `json:"NamespaceName,omitempty" name:"NamespaceName"`
	AllContainer  *bool                                                                              `json:"AllContainer,omitempty" name:"AllContainer"`
	Workload      []*CreateLogRuleInputConfigContainerStandoutSpecifiedWorkloadNamespaceListWorkload `json:"Workload,omitempty" name:"Workload"`
}
type CreateLogRuleInputConfigContainerStandoutSpecifiedWorkload struct {
	NamespaceList []*CreateLogRuleInputConfigContainerStandoutSpecifiedWorkloadNamespaceList `json:"NamespaceList,omitempty" name:"NamespaceList"`
}
type CreateLogRuleInputConfigContainerStandoutSpecifiedPodLabelContainerLabel struct {
	Key      *string `json:"Key,omitempty" name:"Key"`
	Value    *string `json:"Value,omitempty" name:"Value"`
	Operator *int    `json:"Operator,omitempty" name:"Operator"`
}
type CreateLogRuleInputConfigContainerStandoutSpecifiedPodLabel struct {
	NamespaceOperator     *int                                                                        `json:"NamespaceOperator,omitempty" name:"NamespaceOperator"`
	Namespace             []*string                                                                   `json:"Namespace,omitempty" name:"Namespace"`
	ContainerNameOperator *int                                                                        `json:"ContainerNameOperator,omitempty" name:"ContainerNameOperator"`
	ContainerName         *string                                                                     `json:"ContainerName,omitempty" name:"ContainerName"`
	ContainerLabel        []*CreateLogRuleInputConfigContainerStandoutSpecifiedPodLabelContainerLabel `json:"ContainerLabel,omitempty" name:"ContainerLabel"`
}
type CreateLogRuleInputConfigContainerStandout struct {
	AllContainer      *bool                                                       `json:"AllContainer,omitempty" name:"AllContainer"`
	SpecifiedWorkload *CreateLogRuleInputConfigContainerStandoutSpecifiedWorkload `json:"SpecifiedWorkload,omitempty" name:"SpecifiedWorkload"`
	SpecifiedPodLabel *CreateLogRuleInputConfigContainerStandoutSpecifiedPodLabel `json:"SpecifiedPodLabel,omitempty" name:"SpecifiedPodLabel"`
}
type CreateLogRuleInputConfigContainerFileSpecifiedContainer struct {
	Namespace     *string `json:"Namespace,omitempty" name:"Namespace"`
	WorkloadType  *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	WorkloadName  *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	ContainerPath *string `json:"ContainerPath,omitempty" name:"ContainerPath"`
}
type CreateLogRuleInputConfigContainerFileSpecifiedPodLabelContainerLabel struct {
	Key      *string `json:"Key,omitempty" name:"Key"`
	Value    *string `json:"Value,omitempty" name:"Value"`
	Operator *int    `json:"Operator,omitempty" name:"Operator"`
}
type CreateLogRuleInputConfigContainerFileSpecifiedPodLabel struct {
	NamespaceOperator     *int                                                                    `json:"NamespaceOperator,omitempty" name:"NamespaceOperator"`
	Namespace             []*string                                                               `json:"Namespace,omitempty" name:"Namespace"`
	ContainerNameOperator *int                                                                    `json:"ContainerNameOperator,omitempty" name:"ContainerNameOperator"`
	ContainerName         *string                                                                 `json:"ContainerName,omitempty" name:"ContainerName"`
	ContainerPath         []*string                                                               `json:"ContainerPath,omitempty" name:"ContainerPath"`
	ContainerLabel        []*CreateLogRuleInputConfigContainerFileSpecifiedPodLabelContainerLabel `json:"ContainerLabel,omitempty" name:"ContainerLabel"`
}
type CreateLogRuleInputConfigContainerFile struct {
	SpecifiedContainer *CreateLogRuleInputConfigContainerFileSpecifiedContainer `json:"SpecifiedContainer,omitempty" name:"SpecifiedContainer"`
	SpecifiedPodLabel  *CreateLogRuleInputConfigContainerFileSpecifiedPodLabel  `json:"SpecifiedPodLabel,omitempty" name:"SpecifiedPodLabel"`
}
type CreateLogRuleInputConfigHostFileLabel struct {
	Value *string `json:"Value,omitempty" name:"Value"`
	Key   *string `json:"Key,omitempty" name:"Key"`
}
type CreateLogRuleInputConfigHostFile struct {
	Path  *string                                  `json:"Path,omitempty" name:"Path"`
	Label []*CreateLogRuleInputConfigHostFileLabel `json:"Label,omitempty" name:"Label"`
}
type CreateLogRuleInputConfig struct {
	InputType         *int                                       `json:"InputType,omitempty" name:"InputType"`
	ContainerStandout *CreateLogRuleInputConfigContainerStandout `json:"ContainerStandout,omitempty" name:"ContainerStandout"`
	ContainerFile     *CreateLogRuleInputConfigContainerFile     `json:"ContainerFile,omitempty" name:"ContainerFile"`
	HostFile          *CreateLogRuleInputConfigHostFile          `json:"HostFile,omitempty" name:"HostFile"`
}
type CreateLogRuleOutputConfigKlog struct {
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	PoolName    *string `json:"PoolName,omitempty" name:"PoolName"`
}
type CreateLogRuleOutputConfig struct {
	Klog *CreateLogRuleOutputConfigKlog `json:"Klog,omitempty" name:"Klog"`
}

type DescribeClusterInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId  *string                          `json:"ClusterId,omitempty" name:"ClusterId"`
	MaxResults *int                             `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int                             `json:"Marker,omitempty" name:"Marker"`
	Filter     []*DescribeClusterInstanceFilter `json:"Filter,omitempty" name:"Filter"`
	Search     *string                          `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeClusterInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClusterInstanceResponse struct {
	*ksyunhttp.BaseResponse
	InstanceSet []struct {
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		InstanceName    *string `json:"InstanceName" name:"InstanceName"`
		InstanceRole    *string `json:"InstanceRole" name:"InstanceRole"`
		InstanceStatus  *string `json:"InstanceStatus" name:"InstanceStatus"`
		KecInstancePara struct {
			ProjectId         *int    `json:"ProjectId" name:"ProjectId"`
			InstanceType      *string `json:"InstanceType" name:"InstanceType"`
			InstanceConfigure struct {
				VCPU         *int    `json:"VCPU" name:"VCPU"`
				MemoryGb     *int    `json:"MemoryGb" name:"MemoryGb"`
				GPU          *int    `json:"GPU" name:"GPU"`
				DataDiskGb   *int    `json:"DataDiskGb" name:"DataDiskGb"`
				RootDiskGb   *int    `json:"RootDiskGb" name:"RootDiskGb"`
				DataDiskType *string `json:"DataDiskType" name:"DataDiskType"`
			} `json:"InstanceConfigure"`
			SystemDisk struct {
				DiskType *string `json:"DiskType" name:"DiskType"`
				DiskSize *int    `json:"DiskSize" name:"DiskSize"`
			} `json:"SystemDisk"`
			ImageId             *string `json:"ImageId" name:"ImageId"`
			PrivateIpAddress    *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			ChargeType          *string `json:"ChargeType" name:"ChargeType"`
			CreateTime          *string `json:"CreateTime" name:"CreateTime"`
			AvailabilityZone    *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			SubnetId            *string `json:"SubnetId" name:"SubnetId"`
			VpcId               *string `json:"VpcId" name:"VpcId"`
			NetworkInterfaceSet []struct {
				NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
				NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
				SubnetId             *string `json:"SubnetId" name:"SubnetId"`
				PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
				MacAddress           *string `json:"MacAddress" name:"MacAddress"`
				SecurityGroupSet     []struct {
					SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				} `json:"SecurityGroupSet" name:"SecurityGroupSet"`
			} `json:"NetworkInterfaceSet"`
			DedicatedName *string `json:"DedicatedName" name:"DedicatedName"`
			DedicatedId   *string `json:"DedicatedId" name:"DedicatedId"`
		} `json:"KecInstancePara" name:"KecInstancePara"`
		UnSchedulable   *bool   `json:"UnSchedulable" name:"UnSchedulable"`
		DrainStatus     *string `json:"DrainStatus" name:"DrainStatus"`
		NodePoolId      *string `json:"NodePoolId" name:"NodePoolId"`
		AdvancedSetting struct {
			DataDisk struct {
				AutoFormatAndMount *bool   `json:"AutoFormatAndMount" name:"AutoFormatAndMount"`
				FileSystem         *string `json:"FileSystem" name:"FileSystem"`
				MountTarget        *string `json:"MountTarget" name:"MountTarget"`
			} `json:"DataDisk"`
			ContainerRuntime *string `json:"ContainerRuntime" name:"ContainerRuntime"`
			DockerPath       *string `json:"DockerPath" name:"DockerPath"`
			ContainerPath    *string `json:"ContainerPath" name:"ContainerPath"`
			UserScript       *string `json:"UserScript" name:"UserScript"`
			PreUserScript    *string `json:"PreUserScript" name:"PreUserScript"`
			Schedulable      *bool   `json:"Schedulable" name:"Schedulable"`
			Label            []struct {
				Key   *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
			} `json:"Label"`
			ExtraArg struct {
				Kubelet []*string `json:"Kubelet" name:"Kubelet"`
			} `json:"ExtraArg"`
			ContainerLogMaxSize  *int `json:"ContainerLogMaxSize" name:"ContainerLogMaxSize"`
			ContainerLogMaxFiles *int `json:"ContainerLogMaxFiles" name:"ContainerLogMaxFiles"`
		} `json:"AdvancedSetting" name:"AdvancedSetting"`
		EpcInstancePara struct {
			ProjectId    *int    `json:"ProjectId" name:"ProjectId"`
			InstanceType *string `json:"InstanceType" name:"InstanceType"`
			Cpu          struct {
				Model     *string `json:"Model" name:"Model"`
				Frequence *string `json:"Frequence" name:"Frequence"`
				Count     *int    `json:"Count" name:"Count"`
				CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			} `json:"Cpu"`
			Memory *string `json:"Memory" name:"Memory"`
			Gpu    struct {
				Model     *string `json:"Model" name:"Model"`
				Frequence *string `json:"Frequence" name:"Frequence"`
				Count     *int    `json:"Count" name:"Count"`
				CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			} `json:"Gpu"`
			DiskSet []struct {
				DiskType *string `json:"DiskType" name:"DiskType"`
				Raid     *string `json:"Raid" name:"Raid"`
				Space    *string `json:"Space" name:"Space"`
			} `json:"DiskSet"`
			OsName              *string `json:"OsName" name:"OsName"`
			ImageId             *string `json:"ImageId" name:"ImageId"`
			CreateTime          *string `json:"CreateTime" name:"CreateTime"`
			AvailabilityZone    *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			NetworkInterfaceSet []struct {
				NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
				NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
				SubnetId             *string `json:"SubnetId" name:"SubnetId"`
				PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
				MacAddress           *string `json:"MacAddress" name:"MacAddress"`
				SecurityGroupSet     []struct {
					SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				} `json:"SecurityGroupSet" name:"SecurityGroupSet"`
			} `json:"NetworkInterfaceSet"`
			EnableContainer *bool   `json:"EnableContainer" name:"EnableContainer"`
			ProductType     *string `json:"ProductType" name:"ProductType"`
		} `json:"EpcInstancePara" name:"EpcInstancePara"`
		Type         *string `json:"Type" name:"Type"`
		ErrorMessage *string `json:"ErrorMessage" name:"ErrorMessage"`
		NodeName     *string `json:"NodeName" name:"NodeName"`
	} `json:"InstanceSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
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
	ClusterId          *string `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	IsPublic  *bool   `json:"IsPublic,omitempty" name:"IsPublic"`
}

func (r *DownloadClusterConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DownloadClusterConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	Expose        *bool   `json:"Expose" name:"Expose"`
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
	ClusterId       *string                           `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName     *string                           `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterDesc     *string                           `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	EnableKMSE      *bool                             `json:"EnableKMSE,omitempty" name:"EnableKMSE"`
	ControlPlaneLog *ModifyClusterInfoControlPlaneLog `json:"ControlPlaneLog,omitempty" name:"ControlPlaneLog"`
}

func (r *ModifyClusterInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DescribeInstanceImageResponse struct {
	*ksyunhttp.BaseResponse
	ImageSet []struct {
		ImageId            *string   `json:"ImageId" name:"ImageId"`
		ImageName          *string   `json:"ImageName" name:"ImageName"`
		ImageType          *string   `json:"ImageType" name:"ImageType"`
		MatchedK8sVersions []*string `json:"MatchedK8sVersions" name:"MatchedK8sVersions"`
	} `json:"ImageSet"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
	ClusterId   *string                           `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceSet []*AddClusterInstancesInstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet"`
}

func (r *AddClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddClusterInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
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
	ClusterId          *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceId         []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceDeleteMode *string   `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteClusterInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		Message    *string `json:"Message" name:"Message"`
		Return     *bool   `json:"Return" name:"Return"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
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
	ClusterId    *string                        `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceId   []*string                      `json:"InstanceId,omitempty" name:"InstanceId"`
	Filter       []*DescribeEpcForClusterFilter `json:"Filter,omitempty" name:"Filter"`
	Marker       *int                           `json:"Marker,omitempty" name:"Marker"`
	MaxResults   *int                           `json:"MaxResults,omitempty" name:"MaxResults"`
	OperatorType *string                        `json:"OperatorType,omitempty" name:"OperatorType"`
}

func (r *DescribeEpcForClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcForClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Marker      *int    `json:"Marker" name:"Marker"`
	MaxResults  *int    `json:"MaxResults" name:"MaxResults"`
	TotalCount  *string `json:"TotalCount" name:"TotalCount"`
	InstanceSet []struct {
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		InstanceName    *string `json:"InstanceName" name:"InstanceName"`
		EpcInstancePara struct {
			ProjectId    *int    `json:"ProjectId" name:"ProjectId"`
			InstanceType *string `json:"InstanceType" name:"InstanceType"`
			Cpu          struct {
				Model     *string `json:"Model" name:"Model"`
				Frequence *string `json:"Frequence" name:"Frequence"`
				Count     *int    `json:"Count" name:"Count"`
				CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			} `json:"Cpu"`
			Memory *string `json:"Memory" name:"Memory"`
			Gpu    struct {
				Model     *string `json:"Model" name:"Model"`
				Frequence *string `json:"Frequence" name:"Frequence"`
				Count     *int    `json:"Count" name:"Count"`
				CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			} `json:"Gpu"`
			DiskSet []struct {
				DiskType *string `json:"DiskType" name:"DiskType"`
				Raid     *string `json:"Raid" name:"Raid"`
				Space    *string `json:"Space" name:"Space"`
			} `json:"DiskSet"`
			ImageId             *string `json:"ImageId" name:"ImageId"`
			OsName              *string `json:"OsName" name:"OsName"`
			AvailabilityZone    *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			NetworkInterfaceSet []struct {
				NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
				NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
				SubnetId             *string `json:"SubnetId" name:"SubnetId"`
				PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
				MacAddress           *string `json:"MacAddress" name:"MacAddress"`
				SecurityGroupSet     []struct {
					SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				} `json:"SecurityGroupSet" name:"SecurityGroupSet"`
			} `json:"NetworkInterfaceSet"`
			ProductType     *string `json:"ProductType" name:"ProductType"`
			EnableContainer *bool   `json:"EnableContainer" name:"EnableContainer"`
		} `json:"EpcInstancePara" name:"EpcInstancePara"`
		Type           *string `json:"Type" name:"Type"`
		InstanceStatus *string `json:"InstanceStatus" name:"InstanceStatus"`
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
	ClusterId       *string                                `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceId      []*string                              `json:"InstanceId,omitempty" name:"InstanceId"`
	EpcPara         []*string                              `json:"EpcPara,omitempty" name:"EpcPara"`
	AdvancedSetting *AddClusterEpcInstancesAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
}

func (r *AddClusterEpcInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddClusterEpcInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		Return     *bool   `json:"Return" name:"Return"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
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
	ClusterId  *string                           `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceId []*string                         `json:"InstanceId,omitempty" name:"InstanceId"`
	Marker     *int                              `json:"Marker,omitempty" name:"Marker"`
	MaxResults *string                           `json:"MaxResults,omitempty" name:"MaxResults"`
	Filter     []*DescribeExistedInstancesFilter `json:"Filter,omitempty" name:"Filter"`
	Search     *string                           `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeExistedInstancesResponse struct {
	*ksyunhttp.BaseResponse
	InstanceSet []struct {
		InstanceId        *string `json:"InstanceId" name:"InstanceId"`
		InstanceName      *string `json:"InstanceName" name:"InstanceName"`
		InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		PrivateIpAddress  *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		Available         *bool   `json:"Available" name:"Available"`
		UnavailableReason *string `json:"UnavailableReason" name:"UnavailableReason"`
		ClusterId         *string `json:"ClusterId" name:"ClusterId"`
	} `json:"InstanceSet"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
	ClusterId             *string                                     `json:"ClusterId,omitempty" name:"ClusterId"`
	ExistedInstanceKecSet []*AddExistedInstancesExistedInstanceKecSet `json:"ExistedInstanceKecSet,omitempty" name:"ExistedInstanceKecSet"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddExistedInstancesResponse struct {
	*ksyunhttp.BaseResponse
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return     *bool   `json:"Return" name:"Return"`
		Reason     *string `json:"Reason" name:"Reason"`
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
	NodePoolName        *string                     `json:"NodePoolName,omitempty" name:"NodePoolName"`
	ClusterId           *string                     `json:"ClusterId,omitempty" name:"ClusterId"`
	EnableAutoScale     *bool                       `json:"EnableAutoScale,omitempty" name:"EnableAutoScale"`
	NodeTemplate        *CreateNodePoolNodeTemplate `json:"NodeTemplate,omitempty" name:"NodeTemplate"`
	Label               []*CreateNodePoolLabel      `json:"Label,omitempty" name:"Label"`
	Taint               []*CreateNodePoolTaint      `json:"Taint,omitempty" name:"Taint"`
	MinSize             *int                        `json:"MinSize,omitempty" name:"MinSize"`
	MaxSize             *int                        `json:"MaxSize,omitempty" name:"MaxSize"`
	DesiredCapacity     *int                        `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
	EnableDelProtection *bool                       `json:"EnableDelProtection,omitempty" name:"EnableDelProtection"`
	FailureAutoDelete   *bool                       `json:"FailureAutoDelete,omitempty" name:"FailureAutoDelete"`
}

func (r *CreateNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateNodePoolResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
	ClusterId    *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId   []*string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	Marker       *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults   *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NodePoolName *string   `json:"NodePoolName,omitempty" name:"NodePoolName"`
}

func (r *DescribeNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNodePoolResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	MaxResults  *int    `json:"MaxResults" name:"MaxResults"`
	Marker      *int    `json:"Marker" name:"Marker"`
	TotalCount  *int    `json:"TotalCount" name:"TotalCount"`
	NodePoolSet []struct {
		NodePoolId      *string `json:"NodePoolId" name:"NodePoolId"`
		NodePoolName    *string `json:"NodePoolName" name:"NodePoolName"`
		Status          *string `json:"Status" name:"Status"`
		EnableAutoScale *bool   `json:"EnableAutoScale" name:"EnableAutoScale"`
		NodeTemplate    struct {
			ChargeType         *string `json:"ChargeType" name:"ChargeType"`
			InstanceType       *string `json:"InstanceType" name:"InstanceType"`
			InstanceName       *string `json:"InstanceName" name:"InstanceName"`
			InstanceNameSuffix *int    `json:"InstanceNameSuffix" name:"InstanceNameSuffix"`
			SystemDisk         struct {
				DiskType *string `json:"DiskType" name:"DiskType"`
				DiskSize *int    `json:"DiskSize" name:"DiskSize"`
			} `json:"SystemDisk"`
			DataDiskGb *int `json:"DataDiskGb" name:"DataDiskGb"`
			DataDisk   struct {
				DiskType           *string `json:"DiskType" name:"DiskType"`
				DiskSize           *int    `json:"DiskSize" name:"DiskSize"`
				DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
			} `json:"DataDisk"`
			ImageId         *string   `json:"ImageId" name:"ImageId"`
			VpcId           *string   `json:"VpcId" name:"VpcId"`
			SubnetIdSet     []*string `json:"SubnetIdSet" name:"SubnetIdSet"`
			SubnetStrategy  *string   `json:"SubnetStrategy" name:"SubnetStrategy"`
			SecurityGroupId *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			ProjectId       *int      `json:"ProjectId" name:"ProjectId"`
			Password        *string   `json:"Password" name:"Password"`
			KeyIdSet        []*string `json:"KeyIdSet" name:"KeyIdSet"`
			AdvancedSetting struct {
				DataDisk struct {
					AutoFormatAndMount *bool   `json:"AutoFormatAndMount" name:"AutoFormatAndMount"`
					FileSystem         *string `json:"FileSystem" name:"FileSystem"`
					MountTarget        *string `json:"MountTarget" name:"MountTarget"`
				} `json:"DataDisk" name:"DataDisk"`
				ContainerRuntime *string `json:"ContainerRuntime" name:"ContainerRuntime"`
				DockerPath       *string `json:"DockerPath" name:"DockerPath"`
				ContainerPath    *string `json:"ContainerPath" name:"ContainerPath"`
				UserScript       *string `json:"UserScript" name:"UserScript"`
				PreUserScript    *string `json:"PreUserScript" name:"PreUserScript"`
				Schedulable      *bool   `json:"Schedulable" name:"Schedulable"`
				Label            []struct {
					Key   *string `json:"Key" name:"Key"`
					Value *string `json:"Value" name:"Value"`
				} `json:"Label" name:"Label"`
				ExtraArg struct {
					Kubelet []struct {
						CustomArg *string `json:"CustomArg" name:"CustomArg"`
					} `json:"Kubelet"`
				} `json:"ExtraArg" name:"ExtraArg"`
				ContainerLogMaxSize  *string `json:"ContainerLogMaxSize" name:"ContainerLogMaxSize"`
				ContainerLogMaxFiles *string `json:"ContainerLogMaxFiles" name:"ContainerLogMaxFiles"`
				Taints               []struct {
					Key    *string `json:"Key" name:"Key"`
					Value  *string `json:"Value" name:"Value"`
					Effect *string `json:"Effect" name:"Effect"`
				} `json:"Taints" name:"Taints"`
			} `json:"AdvancedSetting"`
			EbsTags []struct {
				Key   *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
			} `json:"EbsTags"`
			ScalingConfigurationId   *string `json:"ScalingConfigurationId" name:"ScalingConfigurationId"`
			ScalingConfigurationName *string `json:"ScalingConfigurationName" name:"ScalingConfigurationName"`
			Mem                      *string `json:"Mem" name:"Mem"`
			RemovePolicy             *string `json:"RemovePolicy" name:"RemovePolicy"`
			InstanceTags             []struct {
				Key   *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
			} `json:"InstanceTags"`
			DeleteDataDisk    *bool   `json:"DeleteDataDisk" name:"DeleteDataDisk"`
			DeleteInstanceTag *bool   `json:"DeleteInstanceTag" name:"DeleteInstanceTag"`
			DeleteEbsTag      *bool   `json:"DeleteEbsTag" name:"DeleteEbsTag"`
			Cpu               *string `json:"Cpu" name:"Cpu"`
		} `json:"NodeTemplate" name:"NodeTemplate"`
		Labels []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Labels" name:"Labels"`
		Taints []struct {
			Key    *string `json:"Key" name:"Key"`
			Value  *string `json:"Value" name:"Value"`
			Effect *string `json:"Effect" name:"Effect"`
		} `json:"Taints" name:"Taints"`
		MinSize            *int    `json:"MinSize" name:"MinSize"`
		MaxSize            *int    `json:"MaxSize" name:"MaxSize"`
		DesiredCapacity    *int    `json:"DesiredCapacity" name:"DesiredCapacity"`
		InstanceCount      *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		ClusterId          *string `json:"ClusterId" name:"ClusterId"`
		ErrorStatusMessage *string `json:"ErrorStatusMessage" name:"ErrorStatusMessage"`
		FailureAutoDelete  *bool   `json:"FailureAutoDelete" name:"FailureAutoDelete"`
	} `json:"NodePoolSet"`
	ScaleUpPolicy   *string `json:"ScaleUpPolicy" name:"ScaleUpPolicy"`
	ScaleDownPolicy struct {
		ScaleDownEnabled              *bool `json:"ScaleDownEnabled" name:"ScaleDownEnabled"`
		ScaleDownDelayAfterAdd        *int  `json:"ScaleDownDelayAfterAdd" name:"ScaleDownDelayAfterAdd"`
		ScaleDownUnneededTime         *int  `json:"ScaleDownUnneededTime" name:"ScaleDownUnneededTime"`
		ScaleDownUtilizationThreshold *int  `json:"ScaleDownUtilizationThreshold" name:"ScaleDownUtilizationThreshold"`
		MaxEmptyBulkDelete            *int  `json:"MaxEmptyBulkDelete" name:"MaxEmptyBulkDelete"`
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
	ClusterId          *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId         []*string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	InstanceDeleteMode *string   `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteNodePoolResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ReturnSet []struct {
		ScalingGroupId *string `json:"ScalingGroupId" name:"ScalingGroupId"`
		Message        *string `json:"Message" name:"Message"`
		Return         *bool   `json:"Return" name:"Return"`
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
	ClusterId           *string                `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId          *string                `json:"NodePoolId,omitempty" name:"NodePoolId"`
	NodePoolName        *string                `json:"NodePoolName,omitempty" name:"NodePoolName"`
	EnableAutoScale     *bool                  `json:"EnableAutoScale,omitempty" name:"EnableAutoScale"`
	MinSize             *int                   `json:"MinSize,omitempty" name:"MinSize"`
	MaxSize             *int                   `json:"MaxSize,omitempty" name:"MaxSize"`
	DesiredCapacity     *int                   `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
	Label               []*ModifyNodePoolLabel `json:"Label,omitempty" name:"Label"`
	Taint               []*ModifyNodePoolTaint `json:"Taint,omitempty" name:"Taint"`
	UpdateExistingNodes *bool                  `json:"UpdateExistingNodes,omitempty" name:"UpdateExistingNodes"`
}

func (r *ModifyNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	ClusterId    *string                         `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId   *string                         `json:"NodePoolId,omitempty" name:"NodePoolId"`
	NodeTemplate *ModifyNodeTemplateNodeTemplate `json:"NodeTemplate,omitempty" name:"NodeTemplate"`
}

func (r *ModifyNodeTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	ClusterId     *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ScaleUpPolicy *string `json:"ScaleUpPolicy,omitempty" name:"ScaleUpPolicy"`
}

func (r *ModifyNodePoolScaleUpPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	ClusterId                     *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MaxEmptyBulkDelete            *int    `json:"MaxEmptyBulkDelete,omitempty" name:"MaxEmptyBulkDelete"`
	ScaleDownDelayAfterAdd        *int    `json:"ScaleDownDelayAfterAdd,omitempty" name:"ScaleDownDelayAfterAdd"`
	ScaleDownEnabled              *bool   `json:"ScaleDownEnabled,omitempty" name:"ScaleDownEnabled"`
	ScaleDownUnneededTime         *int    `json:"ScaleDownUnneededTime,omitempty" name:"ScaleDownUnneededTime"`
	ScaleDownUtilizationThreshold *int    `json:"ScaleDownUtilizationThreshold,omitempty" name:"ScaleDownUtilizationThreshold"`
}

func (r *ModifyNodePoolScaleDownPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId  *string   `json:"NodePoolId,omitempty" name:"NodePoolId"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AddClusterInstanceToNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddClusterInstanceToNodePoolResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return     *bool   `json:"Return" name:"Return"`
		Message    *string `json:"Message" name:"Message"`
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
	ClusterId              *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId             *string   `json:"NodePoolId,omitempty" name:"NodePoolId"`
	InstanceIds            []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	ProtectedFromScaleDown *bool     `json:"ProtectedFromScaleDown,omitempty" name:"ProtectedFromScaleDown"`
}

func (r *ProtectedFromScaleDownRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ProtectedFromScaleDownResponse struct {
	*ksyunhttp.BaseResponse
	SucceedInstanceIds []*string `json:"SucceedInstanceIds" name:"SucceedInstanceIds"`
	FailedInstanceIds  []*string `json:"FailedInstanceIds" name:"FailedInstanceIds"`
	RequestId          *string   `json:"RequestId" name:"RequestId"`
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
	ClusterId          *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodePoolId         *string   `json:"NodePoolId,omitempty" name:"NodePoolId"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	InstanceDeleteMode *bool     `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterInstancesFromNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteClusterInstancesFromNodePoolResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Return     *bool   `json:"Return" name:"Return"`
		Message    *string `json:"Message" name:"Message"`
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

type DescribeEpcImageResponse struct {
	*ksyunhttp.BaseResponse
	ImageSet []struct {
		ImageId   *string `json:"ImageId" name:"ImageId"`
		ImageName *string `json:"ImageName" name:"ImageName"`
		ImageType *string `json:"ImageType" name:"ImageType"`
	} `json:"ImageSet"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
	ClusterId             *string `json:"ClusterId,omitempty" name:"ClusterId"`
	EnableEventCollecting *bool   `json:"EnableEventCollecting,omitempty" name:"EnableEventCollecting"`
}

func (r *EditEventCollectingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DescribeNodePoolSummaryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NodePools []struct {
		NodePoolId   *string `json:"NodePoolId" name:"NodePoolId"`
		NodePoolName *string `json:"NodePoolName" name:"NodePoolName"`
	} `json:"NodePools"`
}

func (r *DescribeNodePoolSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodePoolSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogRuleRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId    *string                    `json:"ClusterId,omitempty" name:"ClusterId"`
	RuleName     *string                    `json:"RuleName,omitempty" name:"RuleName"`
	InputConfig  *CreateLogRuleInputConfig  `json:"InputConfig,omitempty" name:"InputConfig"`
	OutputConfig *CreateLogRuleOutputConfig `json:"OutputConfig,omitempty" name:"OutputConfig"`
}

func (r *CreateLogRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateLogRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateLogRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeClusterSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClusterSummaryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	ClusterSet struct {
		ClusterId         *string `json:"ClusterId" name:"ClusterId"`
		ClusterName       *string `json:"ClusterName" name:"ClusterName"`
		ClusterManageMode *string `json:"ClusterManageMode" name:"ClusterManageMode"`
		K8sVersion        *string `json:"K8sVersion" name:"K8sVersion"`
		PodCidr           *string `json:"PodCidr" name:"PodCidr"`
		ServiceCidr       *string `json:"ServiceCidr" name:"ServiceCidr"`
		VpcId             *string `json:"VpcId" name:"VpcId"`
		VpcCidr           *string `json:"VpcCidr" name:"VpcCidr"`
		NetworkType       *string `json:"NetworkType" name:"NetworkType"`
		Status            *string `json:"Status" name:"Status"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
	} `json:"ClusterSet"`
}

func (r *DescribeClusterSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNodePoolDelProtectionRequest struct {
	*ksyunhttp.BaseRequest
	NodePoolId          *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	EnableDelProtection *bool   `json:"EnableDelProtection,omitempty" name:"EnableDelProtection"`
}

func (r *UpdateNodePoolDelProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateNodePoolDelProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateNodePoolDelProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNodePoolDelProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReleaseRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Filter    *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeReleaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Releases  []struct {
		ReleaseName    *string `json:"ReleaseName" name:"ReleaseName"`
		StatusCode     *string `json:"StatusCode" name:"StatusCode"`
		Namespace      *string `json:"Namespace" name:"Namespace"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		ChartName      *string `json:"ChartName" name:"ChartName"`
		ChartVersion   *string `json:"ChartVersion" name:"ChartVersion"`
		ChartSource    *string `json:"ChartSource" name:"ChartSource"`
		ChartNamespace *string `json:"ChartNamespace" name:"ChartNamespace"`
	} `json:"Releases"`
	ReleaseVersion *int `json:"ReleaseVersion" name:"ReleaseVersion"`
}

func (r *DescribeReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReleaseHistoryRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ReleaseName *string `json:"ReleaseName,omitempty" name:"ReleaseName"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeReleaseHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeReleaseHistoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	ReleaseVersions []struct {
		ReleaseName    *string `json:"ReleaseName" name:"ReleaseName"`
		StatusCode     *string `json:"StatusCode" name:"StatusCode"`
		ReleaseVersion *int    `json:"ReleaseVersion" name:"ReleaseVersion"`
		DeployTime     *string `json:"DeployTime" name:"DeployTime"`
		Description    *string `json:"Description" name:"Description"`
		ChartName      *string `json:"ChartName" name:"ChartName"`
		ChartVersion   *string `json:"ChartVersion" name:"ChartVersion"`
	} `json:"ReleaseVersions"`
}

func (r *DescribeReleaseHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReleaseHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReleaseDetailRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ReleaseName *string `json:"ReleaseName,omitempty" name:"ReleaseName"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeReleaseDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeReleaseDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	ReleaseDetail struct {
		ReleaseName     *string `json:"ReleaseName" name:"ReleaseName"`
		StatusCode      *string `json:"StatusCode" name:"StatusCode"`
		Namespace       *string `json:"Namespace" name:"Namespace"`
		ReleaseVersion  *int    `json:"ReleaseVersion" name:"ReleaseVersion"`
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		LastDeployTime  *string `json:"LastDeployTime" name:"LastDeployTime"`
		ChartName       *string `json:"ChartName" name:"ChartName"`
		ChartVersion    *string `json:"ChartVersion" name:"ChartVersion"`
		ChartNamespace  *string `json:"ChartNamespace" name:"ChartNamespace"`
		ChartSource     *string `json:"ChartSource" name:"ChartSource"`
		DeployResources []struct {
			Kind      *string `json:"Kind" name:"Kind"`
			Content   *string `json:"Content" name:"Content"`
			Name      *string `json:"Name" name:"Name"`
			Namespace *string `json:"Namespace" name:"Namespace"`
		} `json:"DeployResources" name:"DeployResources"`
		Values       *string `json:"Values" name:"Values"`
		CustomValues *string `json:"CustomValues" name:"CustomValues"`
		Description  *string `json:"Description" name:"Description"`
		ChartUrl     *string `json:"ChartUrl" name:"ChartUrl"`
	} `json:"ReleaseDetail"`
}

func (r *DescribeReleaseDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReleaseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReleaseRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ReleaseName *string `json:"ReleaseName,omitempty" name:"ReleaseName"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteReleaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackReleaseRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ReleaseName    *string `json:"ReleaseName,omitempty" name:"ReleaseName"`
	Namespace      *string `json:"Namespace,omitempty" name:"Namespace"`
	ReleaseVersion *int    `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

func (r *RollbackReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RollbackReleaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RollbackReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallReleaseRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId         *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Namespace         *string `json:"Namespace,omitempty" name:"Namespace"`
	ReleaseName       *string `json:"ReleaseName,omitempty" name:"ReleaseName"`
	ChartSource       *string `json:"ChartSource,omitempty" name:"ChartSource"`
	ChartNamespace    *string `json:"ChartNamespace,omitempty" name:"ChartNamespace"`
	ChartName         *string `json:"ChartName,omitempty" name:"ChartName"`
	ChartVersion      *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
	ChartUrl          *string `json:"ChartUrl,omitempty" name:"ChartUrl"`
	ChartRepoType     *string `json:"ChartRepoType,omitempty" name:"ChartRepoType"`
	ChartRepoUsername *string `json:"ChartRepoUsername,omitempty" name:"ChartRepoUsername"`
	ChartRepoPassword *string `json:"ChartRepoPassword,omitempty" name:"ChartRepoPassword"`
	Values            *string `json:"Values,omitempty" name:"Values"`
}

func (r *InstallReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type InstallReleaseResponse struct {
	*ksyunhttp.BaseResponse
	ReuestId *string `json:"ReuestId" name:"ReuestId"`
}

func (r *InstallReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeReleaseRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId         *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Namespace         *string `json:"Namespace,omitempty" name:"Namespace"`
	ReleaseName       *string `json:"ReleaseName,omitempty" name:"ReleaseName"`
	ChartSource       *string `json:"ChartSource,omitempty" name:"ChartSource"`
	ChartNamespace    *string `json:"ChartNamespace,omitempty" name:"ChartNamespace"`
	ChartName         *string `json:"ChartName,omitempty" name:"ChartName"`
	ChartVersion      *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
	ChartUrl          *string `json:"ChartUrl,omitempty" name:"ChartUrl"`
	ChartRepoType     *string `json:"ChartRepoType,omitempty" name:"ChartRepoType"`
	ChartRepoUsername *string `json:"ChartRepoUsername,omitempty" name:"ChartRepoUsername"`
	ChartRepoPassword *string `json:"ChartRepoPassword,omitempty" name:"ChartRepoPassword"`
	Values            *string `json:"Values,omitempty" name:"Values"`
}

func (r *UpgradeReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpgradeReleaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpgradeReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
