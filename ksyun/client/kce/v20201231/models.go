package v20201231

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateClusterManagedClusterMultiMaster struct {
	SubnetId        *string `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSettingDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSettingLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSettingExtraArg struct {
	Kubelet []*CreateClusterInstanceForNodeNodeConfigAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSettingTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSettingMultiDataDisk struct {
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
}
type CreateClusterInstanceForNodeNodeConfigAdvancedSetting struct {
	DataDisk             *CreateClusterInstanceForNodeNodeConfigAdvancedSettingDataDisk        `json:"DataDisk,omitempty" name:"DataDisk"`
	ContainerRuntime     *string                                                               `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                                               `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                                               `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                                               `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                                                 `json:"Schedulable,omitempty" name:"Schedulable"`
	Label                []*CreateClusterInstanceForNodeNodeConfigAdvancedSettingLabel         `json:"Label,omitempty" name:"Label"`
	ExtraArg             *CreateClusterInstanceForNodeNodeConfigAdvancedSettingExtraArg        `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                                                  `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                                                  `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
	Taints               []*CreateClusterInstanceForNodeNodeConfigAdvancedSettingTaints        `json:"Taints,omitempty" name:"Taints"`
	MultiDataDisk        []*CreateClusterInstanceForNodeNodeConfigAdvancedSettingMultiDataDisk `json:"MultiDataDisk,omitempty" name:"MultiDataDisk"`
}
type CreateClusterInstanceForNodeNodeConfig struct {
	Para            *string                                                `json:"Para,omitempty" name:"Para"`
	AdvancedSetting *CreateClusterInstanceForNodeNodeConfigAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
}
type CreateClusterInstanceForNode struct {
	NodeRole   *string                                   `json:"NodeRole,omitempty" name:"NodeRole"`
	NodeConfig []*CreateClusterInstanceForNodeNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
}
type CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingExtraArgKubelet struct {
	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}
type CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingExtraArg struct {
	Kubelet []*CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingExtraArgKubelet `json:"Kubelet,omitempty" name:"Kubelet"`
}
type CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingTaint struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type CreateClusterExistedInstanceForEpcEpcConfigAdvancedSetting struct {
	ContainerRuntime     *string                                                             `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	ContainerPath        *string                                                             `json:"ContainerPath,omitempty" name:"ContainerPath"`
	UserScript           *string                                                             `json:"UserScript,omitempty" name:"UserScript"`
	PreUserScript        *string                                                             `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Schedulable          *bool                                                               `json:"Schedulable,omitempty" name:"Schedulable"`
	Label                []*CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingLabel  `json:"Label,omitempty" name:"Label"`
	ExtraArg             *CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingExtraArg `json:"ExtraArg,omitempty" name:"ExtraArg"`
	ContainerLogMaxSize  *int                                                                `json:"ContainerLogMaxSize,omitempty" name:"ContainerLogMaxSize"`
	ContainerLogMaxFiles *int                                                                `json:"ContainerLogMaxFiles,omitempty" name:"ContainerLogMaxFiles"`
	Taint                []*CreateClusterExistedInstanceForEpcEpcConfigAdvancedSettingTaint  `json:"Taint,omitempty" name:"Taint"`
}
type CreateClusterExistedInstanceForEpcEpcConfig struct {
	Para            *string                                                     `json:"Para,omitempty" name:"Para"`
	AdvancedSetting *CreateClusterExistedInstanceForEpcEpcConfigAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
}
type CreateClusterExistedInstanceForEpc struct {
	NodeRole  *string                                        `json:"NodeRole,omitempty" name:"NodeRole"`
	EpcConfig []*CreateClusterExistedInstanceForEpcEpcConfig `json:"EpcConfig,omitempty" name:"EpcConfig"`
}
type CreateClusterComponent struct {
	Name   *string `json:"Name,omitempty" name:"Name"`
	Config *string `json:"Config,omitempty" name:"Config"`
}
type CreateClusterControlPlaneLog struct {
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Enable      *bool   `json:"Enable,omitempty" name:"Enable"`
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	Items       *string `json:"Items,omitempty" name:"Items"`
}

type CreateClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName               *string                                   `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterType               *string                                   `json:"ClusterType,omitempty" name:"ClusterType"`
	ClusterManageMode         *string                                   `json:"ClusterManageMode,omitempty" name:"ClusterManageMode"`
	ClusterDesc               *string                                   `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	VpcId                     *string                                   `json:"VpcId,omitempty" name:"VpcId"`
	PodCidr                   *string                                   `json:"PodCidr,omitempty" name:"PodCidr"`
	ServiceCidr               *string                                   `json:"ServiceCidr,omitempty" name:"ServiceCidr"`
	NetworkType               *string                                   `json:"NetworkType,omitempty" name:"NetworkType"`
	K8sVersion                *string                                   `json:"K8sVersion,omitempty" name:"K8sVersion"`
	ReserveSubnetId           *string                                   `json:"ReserveSubnetId,omitempty" name:"ReserveSubnetId"`
	PublicApiServer           *string                                   `json:"PublicApiServer,omitempty" name:"PublicApiServer"`
	ExposePublicApiServer     *bool                                     `json:"ExposePublicApiServer,omitempty" name:"ExposePublicApiServer"`
	MaxPodPerNode             *string                                   `json:"MaxPodPerNode,omitempty" name:"MaxPodPerNode"`
	MasterEtcdSeparate        *bool                                     `json:"MasterEtcdSeparate,omitempty" name:"MasterEtcdSeparate"`
	ManagedClusterMultiMaster []*CreateClusterManagedClusterMultiMaster `json:"ManagedClusterMultiMaster,omitempty" name:"ManagedClusterMultiMaster"`
	InstanceForNode           []*CreateClusterInstanceForNode           `json:"InstanceForNode,omitempty" name:"InstanceForNode"`
	ExistedInstanceForEpc     []*CreateClusterExistedInstanceForEpc     `json:"ExistedInstanceForEpc,omitempty" name:"ExistedInstanceForEpc"`
	Component                 []*CreateClusterComponent                 `json:"Component,omitempty" name:"Component"`
	ControlPlaneLog           *CreateClusterControlPlaneLog             `json:"ControlPlaneLog,omitempty" name:"ControlPlaneLog"`
	EnableDelProtection       *bool                                     `json:"EnableDelProtection,omitempty" name:"EnableDelProtection"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ClusterId *string `json:"ClusterId" name:"ClusterId"`
}

func (r *CreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
