package v20230101

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateClusterNetworkApiServer struct {
	EipId                 *string `json:"EipId,omitempty" name:"EipId"`
	PublicApiServerEnable *bool   `json:"PublicApiServerEnable,omitempty" name:"PublicApiServerEnable"`
	ReserveSubnetId       *string `json:"ReserveSubnetId,omitempty" name:"ReserveSubnetId"`
}
type CreateClusterNetworkVpcCNI struct {
	Enable        *bool     `json:"Enable,omitempty" name:"Enable"`
	DaemonMode    *string   `json:"DaemonMode,omitempty" name:"DaemonMode"`
	SubnetIds     []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	SecurityGroup *string   `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
}
type CreateClusterNetwork struct {
	VpcId               *string                        `json:"VpcId,omitempty" name:"VpcId"`
	ApiServer           *CreateClusterNetworkApiServer `json:"ApiServer,omitempty" name:"ApiServer"`
	NetworkPluginVeType *string                        `json:"NetworkPluginVeType,omitempty" name:"NetworkPluginVeType"`
	VpcCNI              *CreateClusterNetworkVpcCNI    `json:"VpcCNI,omitempty" name:"VpcCNI"`
	PodCidr             *string                        `json:"PodCidr,omitempty" name:"PodCidr"`
	ServiceCidr         *string                        `json:"ServiceCidr,omitempty" name:"ServiceCidr"`
	MaxPodPerNode       *int                           `json:"MaxPodPerNode,omitempty" name:"MaxPodPerNode"`
	SANs                []*string                      `json:"SANs,omitempty" name:"SANs"`
}
type CreateClusterNodeInstanceSetBasicSettingSystemDisk struct {
	Type       *string `json:"Type,omitempty" name:"Type"`
	Size       *int    `json:"Size,omitempty" name:"Size"`
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`
}
type CreateClusterNodeInstanceSetBasicSettingContainer struct {
	Runtime     *string `json:"Runtime,omitempty" name:"Runtime"`
	Path        *string `json:"Path,omitempty" name:"Path"`
	LogMaxSize  *int    `json:"LogMaxSize,omitempty" name:"LogMaxSize"`
	LogMaxFiles *int    `json:"LogMaxFiles,omitempty" name:"LogMaxFiles"`
}
type CreateClusterNodeInstanceSetBasicSettingDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	FileSystem         *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget        *string `json:"MountTarget,omitempty" name:"MountTarget"`
	Suffix             *string `json:"Suffix,omitempty" name:"Suffix"`
}
type CreateClusterNodeInstanceSetBasicSettingLoginSetting struct {
	Password *string `json:"Password,omitempty" name:"Password"`
	SSHKeyId *string `json:"SSHKeyId,omitempty" name:"SSHKeyId"`
}
type CreateClusterNodeInstanceSetBasicSetting struct {
	Num                  *int                                                  `json:"Num,omitempty" name:"Num"`
	NodeNameMode         *string                                               `json:"NodeNameMode,omitempty" name:"NodeNameMode"`
	ImageId              *string                                               `json:"ImageId,omitempty" name:"ImageId"`
	SecurityGroupId      *string                                               `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SubnetId             *string                                               `json:"SubnetId,omitempty" name:"SubnetId"`
	InstanceType         *string                                               `json:"InstanceType,omitempty" name:"InstanceType"`
	SystemDisk           *CreateClusterNodeInstanceSetBasicSettingSystemDisk   `json:"SystemDisk,omitempty" name:"SystemDisk"`
	Container            *CreateClusterNodeInstanceSetBasicSettingContainer    `json:"Container,omitempty" name:"Container"`
	InstanceName         *string                                               `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceNameSuffix   *int                                                  `json:"InstanceNameSuffix,omitempty" name:"InstanceNameSuffix"`
	AvailabilityZone     *string                                               `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	DataDisk             *CreateClusterNodeInstanceSetBasicSettingDataDisk     `json:"DataDisk,omitempty" name:"DataDisk"`
	LoginSetting         *CreateClusterNodeInstanceSetBasicSettingLoginSetting `json:"LoginSetting,omitempty" name:"LoginSetting"`
	IsNew                *bool                                                 `json:"IsNew,omitempty" name:"IsNew"`
	SecurityGroupID      *string                                               `json:"SecurityGroupID,omitempty" name:"SecurityGroupID"`
	ChargeType           *string                                               `json:"ChargeType,omitempty" name:"ChargeType"`
	Unit                 *string                                               `json:"Unit,omitempty" name:"Unit"`
	PurchaseTime         *string                                               `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	Raid                 *string                                               `json:"Raid,omitempty" name:"Raid"`
	NetworkInterfaceMode *string                                               `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
	BondName             *string                                               `json:"BondName,omitempty" name:"BondName"`
	GpuImageDriverId     *string                                               `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
}
type CreateClusterNodeInstanceSetAdvancedSettingLabels struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateClusterNodeInstanceSetAdvancedSettingTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type CreateClusterNodeInstanceSetAdvancedSettingContainer struct {
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`
	Path    *string `json:"Path,omitempty" name:"Path"`
}
type CreateClusterNodeInstanceSetAdvancedSetting struct {
	PostUserScript *string                                               `json:"PostUserScript,omitempty" name:"PostUserScript"`
	PreUserScript  *string                                               `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Labels         []*CreateClusterNodeInstanceSetAdvancedSettingLabels  `json:"Labels,omitempty" name:"Labels"`
	Taints         []*CreateClusterNodeInstanceSetAdvancedSettingTaints  `json:"Taints,omitempty" name:"Taints"`
	Container      *CreateClusterNodeInstanceSetAdvancedSettingContainer `json:"Container,omitempty" name:"Container"`
}
type CreateClusterNodeInstanceSetComponents struct {
	Type        *string `json:"Type,omitempty" name:"Type"`
	Args        *string `json:"Args,omitempty" name:"Args"`
	LogMaxSize  *int    `json:"LogMaxSize,omitempty" name:"LogMaxSize"`
	LogMaxFiles *int    `json:"LogMaxFiles,omitempty" name:"LogMaxFiles"`
}
type CreateClusterNodeInstanceSet struct {
	BasicSetting    *CreateClusterNodeInstanceSetBasicSetting    `json:"BasicSetting,omitempty" name:"BasicSetting"`
	AdvancedSetting *CreateClusterNodeInstanceSetAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
	Components      []*CreateClusterNodeInstanceSetComponents    `json:"Components,omitempty" name:"Components"`
	Provider        *string                                      `json:"Provider,omitempty" name:"Provider"`
}
type CreateClusterAddons struct {
	Name *string `json:"Name,omitempty" name:"Name"`
}
type ModifyClusterPublicApiServer struct {
	EipId  *string `json:"EipId,omitempty" name:"EipId"`
	Enable *bool   `json:"Enable,omitempty" name:"Enable"`
}
type ModifyClusterVpcCNI struct {
	Enable        *bool     `json:"Enable,omitempty" name:"Enable"`
	DaemonMode    *string   `json:"DaemonMode,omitempty" name:"DaemonMode"`
	SubnetId      []*string `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroup *string   `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
}
type ModifyNodeComponents struct {
	Type *string `json:"Type,omitempty" name:"Type"`
}
type DescribeComponentParamsComponents struct {
	Type         *string `json:"Type,omitempty" name:"Type"`
	ParamVersion *string `json:"ParamVersion,omitempty" name:"ParamVersion"`
}
type AddKecNodesNodeInstanceSetComponents struct {
	Type *string `json:"Type,omitempty" name:"Type"`
}
type AddKecNodesNodeInstanceSetAdvancedSettingContainer struct {
	Runtime     *string `json:"Runtime,omitempty" name:"Runtime"`
	Path        *string `json:"Path,omitempty" name:"Path"`
	LogMaxSize  *int    `json:"LogMaxSize,omitempty" name:"LogMaxSize"`
	LogMaxFiles *int    `json:"LogMaxFiles,omitempty" name:"LogMaxFiles"`
}
type AddKecNodesNodeInstanceSetAdvancedSetting struct {
	PostUserScript *string                                             `json:"PostUserScript,omitempty" name:"PostUserScript"`
	PreUserScript  *string                                             `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Container      *AddKecNodesNodeInstanceSetAdvancedSettingContainer `json:"Container,omitempty" name:"Container"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetComponents struct {
	Type *string `json:"Type,omitempty" name:"Type"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSettingContainer struct {
	Runtime     *string `json:"Runtime,omitempty" name:"Runtime"`
	Path        *string `json:"Path,omitempty" name:"Path"`
	LogMaxSize  *int    `json:"LogMaxSize,omitempty" name:"LogMaxSize"`
	LogMaxFiles *int    `json:"LogMaxFiles,omitempty" name:"LogMaxFiles"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSettingLabels struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSettingTaints struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSetting struct {
	PostUserScript *string                                                            `json:"PostUserScript,omitempty" name:"PostUserScript"`
	PreUserScript  *string                                                            `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Container      *AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSettingContainer `json:"Container,omitempty" name:"Container"`
	Labels         []*AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSettingLabels  `json:"Labels,omitempty" name:"Labels"`
	Taints         *AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSettingTaints    `json:"Taints,omitempty" name:"Taints"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetBasicSettingSystemDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *int    `json:"Size,omitempty" name:"Size"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetBasicSettingDataDisk struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	AutoFormatAndMount *bool   `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetBasicSettingLoginSetting struct {
	Password *string `json:"Password,omitempty" name:"Password"`
	SSHKeyId *string `json:"SSHKeyId,omitempty" name:"SSHKeyId"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSetBasicSetting struct {
	IsNew            *bool                                                              `json:"IsNew,omitempty" name:"IsNew"`
	AvailabilityZone *string                                                            `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ExistedInstances []*string                                                          `json:"ExistedInstances,omitempty" name:"ExistedInstances"`
	Num              *int                                                               `json:"Num,omitempty" name:"Num"`
	InstanceType     *string                                                            `json:"InstanceType,omitempty" name:"InstanceType"`
	ImageID          *string                                                            `json:"ImageID,omitempty" name:"ImageID"`
	SystemDisk       *AddKecNodesNodeInstanceSetNodeInstanceSetBasicSettingSystemDisk   `json:"SystemDisk,omitempty" name:"SystemDisk"`
	DataDisk         []*AddKecNodesNodeInstanceSetNodeInstanceSetBasicSettingDataDisk   `json:"DataDisk,omitempty" name:"DataDisk"`
	InstanceName     *string                                                            `json:"InstanceName,omitempty" name:"InstanceName"`
	SubnetID         *string                                                            `json:"SubnetID,omitempty" name:"SubnetID"`
	NodeNameMode     *string                                                            `json:"NodeNameMode,omitempty" name:"NodeNameMode"`
	LoginSetting     *AddKecNodesNodeInstanceSetNodeInstanceSetBasicSettingLoginSetting `json:"LoginSetting,omitempty" name:"LoginSetting"`
	SecurityGroupID  *string                                                            `json:"SecurityGroupID,omitempty" name:"SecurityGroupID"`
	ChargeType       *string                                                            `json:"ChargeType,omitempty" name:"ChargeType"`
	Unit             *int                                                               `json:"Unit,omitempty" name:"Unit"`
	PurchaseTime     *int                                                               `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}
type AddKecNodesNodeInstanceSetNodeInstanceSet struct {
	Provider        *string                                                   `json:"Provider,omitempty" name:"Provider"`
	Components      []*AddKecNodesNodeInstanceSetNodeInstanceSetComponents    `json:"Components,omitempty" name:"Components"`
	AdvancedSetting *AddKecNodesNodeInstanceSetNodeInstanceSetAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
	BasicSetting    *AddKecNodesNodeInstanceSetNodeInstanceSetBasicSetting    `json:"BasicSetting,omitempty" name:"BasicSetting"`
}
type AddKecNodesNodeInstanceSet struct {
	Provider        *string                                    `json:"Provider,omitempty" name:"Provider"`
	Components      []*AddKecNodesNodeInstanceSetComponents    `json:"Components,omitempty" name:"Components"`
	AdvancedSetting *AddKecNodesNodeInstanceSetAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
	NodeInstanceSet *AddKecNodesNodeInstanceSetNodeInstanceSet `json:"NodeInstanceSet,omitempty" name:"NodeInstanceSet"`
}
type AddEpcNodesNodeInstanceSetComponents struct {
	Type *string `json:"Type,omitempty" name:"Type"`
}
type AddEpcNodesNodeInstanceSetBasicSettingSystemDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *string `json:"Size,omitempty" name:"Size"`
}
type AddEpcNodesNodeInstanceSetBasicSettingDataDisk struct {
	FileSystem  *string `json:"FileSystem,omitempty" name:"FileSystem"`
	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
	Suffix      *string `json:"Suffix,omitempty" name:"Suffix"`
}
type AddEpcNodesNodeInstanceSetBasicSettingLoginSetting struct {
	Password *string `json:"Password,omitempty" name:"Password"`
	SSHKeyId *string `json:"SSHKeyId,omitempty" name:"SSHKeyId"`
}
type AddEpcNodesNodeInstanceSetBasicSetting struct {
	IsNew                *bool                                               `json:"IsNew,omitempty" name:"IsNew"`
	AvailabilityZone     *string                                             `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ExistedInstances     []*string                                           `json:"ExistedInstances,omitempty" name:"ExistedInstances"`
	Num                  *int                                                `json:"Num,omitempty" name:"Num"`
	InstanceType         *string                                             `json:"InstanceType,omitempty" name:"InstanceType"`
	ImageID              *string                                             `json:"ImageID,omitempty" name:"ImageID"`
	SystemDisk           *AddEpcNodesNodeInstanceSetBasicSettingSystemDisk   `json:"SystemDisk,omitempty" name:"SystemDisk"`
	InstanceName         *string                                             `json:"InstanceName,omitempty" name:"InstanceName"`
	SubnetID             *string                                             `json:"SubnetID,omitempty" name:"SubnetID"`
	NodeNameMode         *string                                             `json:"NodeNameMode,omitempty" name:"NodeNameMode"`
	DataDisk             []*AddEpcNodesNodeInstanceSetBasicSettingDataDisk   `json:"DataDisk,omitempty" name:"DataDisk"`
	LoginSetting         *AddEpcNodesNodeInstanceSetBasicSettingLoginSetting `json:"LoginSetting,omitempty" name:"LoginSetting"`
	Series               *string                                             `json:"Series,omitempty" name:"Series"`
	ReInstall            *bool                                               `json:"ReInstall,omitempty" name:"ReInstall"`
	SecurityGroupID      *string                                             `json:"SecurityGroupID,omitempty" name:"SecurityGroupID"`
	Raid                 *string                                             `json:"Raid,omitempty" name:"Raid"`
	NetworkInterfaceMode *string                                             `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
	BondName             *string                                             `json:"BondName,omitempty" name:"BondName"`
	GpuImageDriverId     *string                                             `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
	ChargeType           *string                                             `json:"ChargeType,omitempty" name:"ChargeType"`
	Unit                 *int                                                `json:"Unit,omitempty" name:"Unit"`
	PurchaseTime         *int                                                `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}
type AddEpcNodesNodeInstanceSetAdvancedSettingContainer struct {
	Runtime     *string `json:"Runtime,omitempty" name:"Runtime"`
	Path        *string `json:"Path,omitempty" name:"Path"`
	LogMaxSize  *int    `json:"LogMaxSize,omitempty" name:"LogMaxSize"`
	LogMaxFiles *int    `json:"LogMaxFiles,omitempty" name:"LogMaxFiles"`
}
type AddEpcNodesNodeInstanceSetAdvancedSettingLabels struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type AddEpcNodesNodeInstanceSetAdvancedSettingTaint struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}
type AddEpcNodesNodeInstanceSetAdvancedSetting struct {
	PostUserScript *string                                             `json:"PostUserScript,omitempty" name:"PostUserScript"`
	PreUserScript  *string                                             `json:"PreUserScript,omitempty" name:"PreUserScript"`
	Container      *AddEpcNodesNodeInstanceSetAdvancedSettingContainer `json:"Container,omitempty" name:"Container"`
	Labels         []*AddEpcNodesNodeInstanceSetAdvancedSettingLabels  `json:"Labels,omitempty" name:"Labels"`
	Taint          *AddEpcNodesNodeInstanceSetAdvancedSettingTaint     `json:"Taint,omitempty" name:"Taint"`
}
type AddEpcNodesNodeInstanceSet struct {
	Provider        *string                                    `json:"Provider,omitempty" name:"Provider"`
	Components      []*AddEpcNodesNodeInstanceSetComponents    `json:"Components,omitempty" name:"Components"`
	BasicSetting    *AddEpcNodesNodeInstanceSetBasicSetting    `json:"BasicSetting,omitempty" name:"BasicSetting"`
	AdvancedSetting *AddEpcNodesNodeInstanceSetAdvancedSetting `json:"AdvancedSetting,omitempty" name:"AdvancedSetting"`
}

type CreateClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName       *string                         `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterDesc       *string                         `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	ClusterManageMode *string                         `json:"ClusterManageMode,omitempty" name:"ClusterManageMode"`
	ProjectId         *string                         `json:"ProjectId,omitempty" name:"ProjectId"`
	KubernetesVersion *string                         `json:"KubernetesVersion,omitempty" name:"KubernetesVersion"`
	Network           *CreateClusterNetwork           `json:"Network,omitempty" name:"Network"`
	NodeInstanceSet   []*CreateClusterNodeInstanceSet `json:"NodeInstanceSet,omitempty" name:"NodeInstanceSet"`
	Addons            []*CreateClusterAddons          `json:"Addons,omitempty" name:"Addons"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
	ClusterId *string `json:"ClusterId" name:"ClusterId"`
}

func (r *CreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*ksyunhttp.BaseRequest
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int      `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClustersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxResults *int `json:"MaxResults" name:"MaxResults"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		ClusterSet []struct {
			ClusterId         *string `json:"ClusterId" name:"ClusterId"`
			ClusterName       *string `json:"ClusterName" name:"ClusterName"`
			ClusterManageMode *string `json:"ClusterManageMode" name:"ClusterManageMode"`
			ProjectId         *string `json:"ProjectId" name:"ProjectId"`
			KubernetesVersion *string `json:"KubernetesVersion" name:"KubernetesVersion"`
			Network           struct {
				NetworkPluginType *string `json:"NetworkPluginType" name:"NetworkPluginType"`
				ApiServer         struct {
					PublicApiServerEnable *bool   `json:"PublicApiServerEnable" name:"PublicApiServerEnable"`
					EipId                 *string `json:"EipId" name:"EipId"`
					ReserveSubnetId       *string `json:"ReserveSubnetId" name:"ReserveSubnetId"`
				} `json:"ApiServer" name:"ApiServer"`
				VpcCNI struct {
					Enable        *bool     `json:"Enable" name:"Enable"`
					DaemonMode    *string   `json:"DaemonMode" name:"DaemonMode"`
					SubnetIds     []*string `json:"SubnetIds" name:"SubnetIds"`
					SecurityGroup *string   `json:"SecurityGroup" name:"SecurityGroup"`
				} `json:"VpcCNI" name:"VpcCNI"`
				PodCidr       *string `json:"PodCidr" name:"PodCidr"`
				ServiceCidr   *string `json:"ServiceCidr" name:"ServiceCidr"`
				MaxPodPerNode *int    `json:"MaxPodPerNode" name:"MaxPodPerNode"`
			} `json:"Network"`
			Addons []struct {
				Name *string `json:"Name" name:"Name"`
			} `json:"Addons"`
			ManagedClusterSpec struct {
			} `json:"ManagedClusterSpec"`
			Status struct {
				Phase *string `json:"Phase" name:"Phase"`
			} `json:"Status"`
		} `json:"ClusterSet" name:"ClusterSet"`
	} `json:"Data"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceDelete *bool   `json:"InstanceDelete,omitempty" name:"InstanceDelete"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteClusterResponse struct {
	*ksyunhttp.BaseResponse
	Request *string `json:"Request" name:"Request"`
	Data    struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
	} `json:"Data"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId       *string                       `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterDesc     *string                       `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	SANs            []*string                     `json:"SANs,omitempty" name:"SANs"`
	PublicApiServer *ModifyClusterPublicApiServer `json:"PublicApiServer,omitempty" name:"PublicApiServer"`
	VpcCNI          *ModifyClusterVpcCNI          `json:"VpcCNI,omitempty" name:"VpcCNI"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
	} `json:"Data"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	KceNodeIds  []*string `json:"KceNodeIds,omitempty" name:"KceNodeIds"`
	Marker      *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNodesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxResults *int `json:"MaxResults" name:"MaxResults"`
		Marker     *int `json:"Marker" name:"Marker"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
	InstanceSet struct {
	} `json:"InstanceSet"`
}

func (r *DescribeNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodeRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodeIds        []*string `json:"NodeIds,omitempty" name:"NodeIds"`
	InstanceDelete *bool     `json:"InstanceDelete,omitempty" name:"InstanceDelete"`
	KceNodeIds     *string   `json:"KceNodeIds,omitempty" name:"KceNodeIds"`
	InstanceIds    []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DeleteNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceSet []struct {
			InstanceId *string `json:"InstanceId" name:"InstanceId"`
			KceNodeId  *string `json:"KceNodeId" name:"KceNodeId"`
		} `json:"InstanceSet" name:"InstanceSet"`
	} `json:"Data"`
}

func (r *DeleteNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeRequest struct {
	*ksyunhttp.BaseRequest
	ClusteId   *string                 `json:"ClusteId,omitempty" name:"ClusteId"`
	KceNodeId  *string                 `json:"KceNodeId,omitempty" name:"KceNodeId"`
	InstanceId *string                 `json:"InstanceId,omitempty" name:"InstanceId"`
	Components []*ModifyNodeComponents `json:"Components,omitempty" name:"Components"`
}

func (r *ModifyNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyNodeResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId  *string `json:"ClusterId" name:"ClusterId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	Components []struct {
		Type *string `json:"Type" name:"Type"`
	} `json:"Components"`
}

func (r *ModifyNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentListRequest struct {
	*ksyunhttp.BaseRequest
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
}

func (r *DescribeComponentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeComponentListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeComponentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeComponentsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	NodeNames   []*string `json:"NodeNames,omitempty" name:"NodeNames"`
	NodeIds     *string   `json:"NodeIds,omitempty" name:"NodeIds"`
	Marker      *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeNodeComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNodeComponentsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId      *string `json:"ClusterId" name:"ClusterId"`
		NodeComponents []struct {
			NodeId          *string `json:"NodeId" name:"NodeId"`
			ComponentStatus []struct {
				Type        *string `json:"Type" name:"Type"`
				CurVersion  *string `json:"CurVersion" name:"CurVersion"`
				SpecVersion *string `json:"SpecVersion" name:"SpecVersion"`
			} `json:"ComponentStatus"`
		} `json:"NodeComponents" name:"NodeComponents"`
	} `json:"Data"`
}

func (r *DescribeNodeComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNetworkResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId *string `json:"ClusterId" name:"ClusterId"`
	NetworkId *string `json:"NetworkId" name:"NetworkId"`
	PublicSLB struct {
		Phase  *string `json:"Phase" name:"Phase"`
		Reason *string `json:"Reason" name:"Reason"`
	} `json:"PublicSLB"`
	PrivateSLB struct {
		SLBId  *string `json:"SLBId" name:"SLBId"`
		SLBIp  *string `json:"SLBIp" name:"SLBIp"`
		Phase  *string `json:"Phase" name:"Phase"`
		Reason *string `json:"Reason" name:"Reason"`
	} `json:"PrivateSLB"`
	PrivateLink struct {
		LinkIp   *string `json:"LinkIp" name:"LinkIp"`
		LinkPort *string `json:"LinkPort" name:"LinkPort"`
		Phase    *string `json:"Phase" name:"Phase"`
		Reason   *string `json:"Reason" name:"Reason"`
	} `json:"PrivateLink"`
	EIP struct {
		EIPId  *string `json:"EIPId" name:"EIPId"`
		Phase  *string `json:"Phase" name:"Phase"`
		Reason *string `json:"Reason" name:"Reason"`
	} `json:"EIP"`
	PublicAccess *bool   `json:"PublicAccess" name:"PublicAccess"`
	CreatedTime  *string `json:"CreatedTime" name:"CreatedTime"`
	UpdatedTime  *string `json:"UpdatedTime" name:"UpdatedTime"`
}

func (r *DescribeNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentParamsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId  *string                              `json:"ClusterId,omitempty" name:"ClusterId"`
	Components []*DescribeComponentParamsComponents `json:"Components,omitempty" name:"Components"`
	Marker     *int                                 `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeComponentParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeComponentParamsResponse struct {
	*ksyunhttp.BaseResponse
	RequeestId *string `json:"RequeestId" name:"RequeestId"`
	Data       struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
	} `json:"Data"`
	Components []struct {
		Type    *string `json:"Type" name:"Type"`
		Version *string `json:"Version" name:"Version"`
		Args    *string `json:"Args" name:"Args"`
	} `json:"Components"`
}

func (r *DescribeComponentParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddonInstancesRequest struct {
	*ksyunhttp.BaseRequest
	CulsterId   *string   `json:"CulsterId,omitempty" name:"CulsterId"`
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	Name        *string   `json:"Name,omitempty" name:"Name"`
	AddonIds    []*string `json:"AddonIds,omitempty" name:"AddonIds"`
}

func (r *DescribeAddonInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAddonInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Addons []struct {
			ClusterId   *string `json:"ClusterId" name:"ClusterId"`
			AddonId     *string `json:"AddonId" name:"AddonId"`
			InstanceId  *string `json:"InstanceId" name:"InstanceId"`
			Type        *string `json:"Type" name:"Type"`
			ToDelete    *bool   `json:"ToDelete" name:"ToDelete"`
			Phase       *string `json:"Phase" name:"Phase"`
			CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
			UpdatedTime *string `json:"UpdatedTime" name:"UpdatedTime"`
		} `json:"Addons" name:"Addons"`
	} `json:"Data"`
	Args struct {
	} `json:"Args"`
}

func (r *DescribeAddonInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddonInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventLogsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	NodeId      *string `json:"NodeId,omitempty" name:"NodeId"`
	NodeName    *string `json:"NodeName,omitempty" name:"NodeName"`
	Inner       *bool   `json:"Inner,omitempty" name:"Inner"`
	Marker      *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int    `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeEventLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEventLogsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeEventLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVersionListRequest struct {
	*ksyunhttp.BaseRequest
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
}

func (r *DescribeClusterVersionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClusterVersionListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *DescribeClusterVersionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddKecNodesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId       *string                     `json:"ClusterId,omitempty" name:"ClusterId"`
	NodeInstanceSet *AddKecNodesNodeInstanceSet `json:"NodeInstanceSet,omitempty" name:"NodeInstanceSet"`
}

func (r *AddKecNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddKecNodesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RequestId   *string `json:"RequestId" name:"RequestId"`
		InstanceSet []struct {
			KceNodeId *string `json:"KceNodeId" name:"KceNodeId"`
			Code      *int    `json:"Code" name:"Code"`
		} `json:"InstanceSet" name:"InstanceSet"`
	} `json:"Data"`
}

func (r *AddKecNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddKecNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddEpcNodesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId       *string                       `json:"ClusterId,omitempty" name:"ClusterId"`
	NodeInstanceSet []*AddEpcNodesNodeInstanceSet `json:"NodeInstanceSet,omitempty" name:"NodeInstanceSet"`
}

func (r *AddEpcNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddEpcNodesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RequestId   *string `json:"RequestId" name:"RequestId"`
		InstanceSet []struct {
			KceNodeId *string `json:"KceNodeId" name:"KceNodeId"`
			Code      *int    `json:"Code" name:"Code"`
		} `json:"InstanceSet" name:"InstanceSet"`
	} `json:"Data"`
}

func (r *AddEpcNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddEpcNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
