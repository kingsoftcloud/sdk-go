package v20200702
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type CreateContainerGroupAdvanceSettingsSystemDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *int    `json:"Size,omitempty" name:"Size"`
}
type CreateContainerGroupAdvanceSettings struct {
	ImageId    *string                                        `json:"ImageId,omitempty" name:"ImageId"`
	DataDiskGb *int                                           `json:"DataDiskGb,omitempty" name:"DataDiskGb"`
	SystemDisk *CreateContainerGroupAdvanceSettingsSystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
}
type CreateContainerGroupMachineDnsConfigOption struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateContainerGroupMachineDnsConfig struct {
	NameServer []*string                                     `json:"NameServer,omitempty" name:"NameServer"`
	Search     []*string                                     `json:"Search,omitempty" name:"Search"`
	Option     []*CreateContainerGroupMachineDnsConfigOption `json:"Option,omitempty" name:"Option"`
}
type CreateContainerGroupMachineHostAliase struct {
	Ip       *string   `json:"Ip,omitempty" name:"Ip"`
	Hostname []*string `json:"Hostname,omitempty" name:"Hostname"`
}
type CreateContainerGroupImageRegistryCredential struct {
	Server   *string `json:"Server,omitempty" name:"Server"`
	Username *string `json:"Username,omitempty" name:"Username"`
	Password *string `json:"Password,omitempty" name:"Password"`
}
type CreateContainerGroupVolumeNFSVolume struct {
	Server   *string `json:"Server,omitempty" name:"Server"`
	Path     *string `json:"Path,omitempty" name:"Path"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type CreateContainerGroupVolumeHostPathVolume struct {
	Path *string `json:"Path,omitempty" name:"Path"`
}
type CreateContainerGroupVolumeEBSVolume struct {
	FsType             *string `json:"FsType,omitempty" name:"FsType"`
	VolumeId           *string `json:"VolumeId,omitempty" name:"VolumeId"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	SnapshotId         *int    `json:"SnapshotId,omitempty" name:"SnapshotId"`
}
type CreateContainerGroupVolumeConfigFileVolumeConfigFileToPath struct {
	Path    *string `json:"Path,omitempty" name:"Path"`
	Content *string `json:"Content,omitempty" name:"Content"`
	Mode    *int    `json:"mode,omitempty" name:"mode"`
}
type CreateContainerGroupVolumeConfigFileVolume struct {
	DefaultMode      *int                                                          `json:"DefaultMode,omitempty" name:"DefaultMode"`
	ConfigFileToPath []*CreateContainerGroupVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" name:"ConfigFileToPath"`
}
type CreateContainerGroupVolume struct {
	Type             *string                                     `json:"Type,omitempty" name:"Type"`
	Name             *string                                     `json:"Name,omitempty" name:"Name"`
	NFSVolume        *CreateContainerGroupVolumeNFSVolume        `json:"NFSVolume,omitempty" name:"NFSVolume"`
	HostPathVolume   *CreateContainerGroupVolumeHostPathVolume   `json:"HostPathVolume,omitempty" name:"HostPathVolume"`
	EBSVolume        *CreateContainerGroupVolumeEBSVolume        `json:"EBSVolume,omitempty" name:"EBSVolume"`
	ConfigFileVolume *CreateContainerGroupVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" name:"ConfigFileVolume"`
}
type CreateContainerGroupContainerLivenessProbeHttpGet struct {
	Port   *int    `json:"Port,omitempty" name:"Port"`
	Path   *string `json:"Path,omitempty" name:"Path"`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}
type CreateContainerGroupContainerLivenessProbeTcpSocket struct {
	Port *int `json:"Port,omitempty" name:"Port"`
}
type CreateContainerGroupContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" name:"Command"`
}
type CreateContainerGroupContainerLivenessProbe struct {
	InitialDelaySeconds *int                                                 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`
	PeriodSeconds       *int                                                 `json:"PeriodSeconds,omitempty" name:"PeriodSeconds"`
	TimeoutSeconds      *int                                                 `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`
	SuccessThreshold    *int                                                 `json:"SuccessThreshold,omitempty" name:"SuccessThreshold"`
	FailureThreshold    *int                                                 `json:"FailureThreshold,omitempty" name:"FailureThreshold"`
	HttpGet             *CreateContainerGroupContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" name:"HttpGet"`
	TcpSocket           *CreateContainerGroupContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" name:"TcpSocket"`
	Exec                *CreateContainerGroupContainerLivenessProbeExec      `json:"Exec,omitempty" name:"Exec"`
}
type CreateContainerGroupContainerReadinessProbeHttpGet struct {
	Port   *int    `json:"Port,omitempty" name:"Port"`
	Path   *string `json:"Path,omitempty" name:"Path"`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}
type CreateContainerGroupContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" name:"Command"`
}
type CreateContainerGroupContainerReadinessProbeTcpSocket struct {
	Port *int `json:"Port,omitempty" name:"Port"`
}
type CreateContainerGroupContainerReadinessProbe struct {
	InitialDelaySeconds *int                                                  `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`
	PeriodSeconds       *int                                                  `json:"PeriodSeconds,omitempty" name:"PeriodSeconds"`
	TimeoutSeconds      *int                                                  `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`
	SuccessThreshold    *int                                                  `json:"SuccessThreshold,omitempty" name:"SuccessThreshold"`
	FailureThreshold    *int                                                  `json:"FailureThreshold,omitempty" name:"FailureThreshold"`
	HttpGet             *CreateContainerGroupContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" name:"HttpGet"`
	Exec                *CreateContainerGroupContainerReadinessProbeExec      `json:"Exec,omitempty" name:"Exec"`
	TcpSocket           *CreateContainerGroupContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" name:"TcpSocket"`
}
type CreateContainerGroupContainerEnvironmentVarValueFromFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
}
type CreateContainerGroupContainerEnvironmentVarValueFrom struct {
	FieldRef *CreateContainerGroupContainerEnvironmentVarValueFromFieldRef `json:"FieldRef,omitempty" name:"FieldRef"`
}
type CreateContainerGroupContainerEnvironmentVar struct {
	Key       *string                                               `json:"Key,omitempty" name:"Key"`
	Value     *string                                               `json:"Value,omitempty" name:"Value"`
	ValueFrom *CreateContainerGroupContainerEnvironmentVarValueFrom `json:"ValueFrom,omitempty" name:"ValueFrom"`
}
type CreateContainerGroupContainerPort struct {
	Port     *int    `json:"Port,omitempty" name:"Port"`
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}
type CreateContainerGroupContainerVolumeMount struct {
	Name      *string `json:"Name,omitempty" name:"Name"`
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`
	ReadOnly  *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type CreateContainerGroupContainer struct {
	Name            *string                                        `json:"Name,omitempty" name:"Name"`
	Command         []*string                                      `json:"Command,omitempty" name:"Command"`
	Arg             []*string                                      `json:"Arg,omitempty" name:"Arg"`
	Cpu             *float64                                       `json:"Cpu,omitempty" name:"Cpu"`
	Memory          *string                                        `json:"Memory,omitempty" name:"Memory"`
	Gpu             *float64                                       `json:"Gpu,omitempty" name:"Gpu"`
	WorkingDir      *string                                        `json:"WorkingDir,omitempty" name:"WorkingDir"`
	Image           *string                                        `json:"Image,omitempty" name:"Image"`
	ImagePullPolicy *string                                        `json:"ImagePullPolicy,omitempty" name:"ImagePullPolicy"`
	LivenessProbe   *CreateContainerGroupContainerLivenessProbe    `json:"LivenessProbe,omitempty" name:"LivenessProbe"`
	ReadinessProbe  *CreateContainerGroupContainerReadinessProbe   `json:"ReadinessProbe,omitempty" name:"ReadinessProbe"`
	EnvironmentVar  []*CreateContainerGroupContainerEnvironmentVar `json:"EnvironmentVar,omitempty" name:"EnvironmentVar"`
	Port            []*CreateContainerGroupContainerPort           `json:"Port,omitempty" name:"Port"`
	VolumeMount     []*CreateContainerGroupContainerVolumeMount    `json:"VolumeMount,omitempty" name:"VolumeMount"`
}
type CreateContainerGroupDnsConfigOption struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateContainerGroupDnsConfig struct {
	NameServer []*string                              `json:"NameServer,omitempty" name:"NameServer"`
	Search     []*string                              `json:"Search,omitempty" name:"Search"`
	Option     []*CreateContainerGroupDnsConfigOption `json:"Option,omitempty" name:"Option"`
}
type CreateContainerGroupHostAliase struct {
	Ip       *string   `json:"Ip,omitempty" name:"Ip"`
	Hostname []*string `json:"Hostname,omitempty" name:"Hostname"`
}
type CreateContainerGroupLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateContainerGroupKubeProxy struct {
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}
type CreateContainerGroupDataDisk struct {
	VolumeName         *string `json:"VolumeName,omitempty" name:"VolumeName"`
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	SnapshotId         *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}
type CreateContainerGroupContainerSpec struct {
	RequestCpu *float64 `json:"RequestCpu,omitempty" name:"RequestCpu"`
	RequestMem *float64 `json:"RequestMem,omitempty" name:"RequestMem"`
	LimitCpu   *float64 `json:"LimitCpu,omitempty" name:"LimitCpu"`
	LimitMem   *float64 `json:"LimitMem,omitempty" name:"LimitMem"`
}
type DescribeContainerGroupFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeContainerGroupCountLabel struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateImageCacheImageRegistryCredential struct {
	Server   *string `json:"Server,omitempty" name:"Server"`
	Username *string `json:"Username,omitempty" name:"Username"`
	Password *string `json:"Password,omitempty" name:"Password"`
}
type UpdateImageCacheImageRegistryCredential struct {
	Server   *string `json:"Server,omitempty" name:"Server"`
	Username *string `json:"Username,omitempty" name:"Username"`
	Password *string `json:"Password,omitempty" name:"Password"`
}


type CreateContainerGroupRequest struct {
	*ksyunhttp.BaseRequest
	ContainerGroupName      *string                                        `json:"ContainerGroupName,omitempty" name:"ContainerGroupName"`
	SubnetId                *string                                        `json:"SubnetId,omitempty" name:"SubnetId"`
	MultiSubnetId           []*string                                      `json:"MultiSubnetId,omitempty" name:"MultiSubnetId"`
	SecurityGroupId         []*string                                      `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	KciType                 *string                                        `json:"KciType,omitempty" name:"KciType"`
	InstanceType            *string                                        `json:"InstanceType,omitempty" name:"InstanceType"`
	InstanceFamily          *string                                        `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	ChargeType              *string                                        `json:"ChargeType,omitempty" name:"ChargeType"`
	SpotStrategy            *string                                        `json:"SpotStrategy,omitempty" name:"SpotStrategy"`
	ProjectId               *int                                           `json:"ProjectId,omitempty" name:"ProjectId"`
	Cpu                     *float64                                       `json:"Cpu,omitempty" name:"Cpu"`
	Memory                  *float64                                       `json:"Memory,omitempty" name:"Memory"`
	Gpu                     *float64                                       `json:"Gpu,omitempty" name:"Gpu"`
	KubeConfig              *string                                        `json:"KubeConfig,omitempty" name:"KubeConfig"`
	RetainIp                *bool                                          `json:"RetainIp,omitempty" name:"RetainIp"`
	RetainIpHours           *int                                           `json:"RetainIpHours,omitempty" name:"RetainIpHours"`
	EipAllocationId         *string                                        `json:"EipAllocationId,omitempty" name:"EipAllocationId"`
	MultiEipAllocationId    []*string                                      `json:"MultiEipAllocationId,omitempty" name:"MultiEipAllocationId"`
	AutoMatchImageCache     *bool                                          `json:"AutoMatchImageCache,omitempty" name:"AutoMatchImageCache"`
	ImageCacheId            *string                                        `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
	AdvanceSettings         *CreateContainerGroupAdvanceSettings           `json:"AdvanceSettings,omitempty" name:"AdvanceSettings"`
	MachineDnsConfig        *CreateContainerGroupMachineDnsConfig          `json:"MachineDnsConfig,omitempty" name:"MachineDnsConfig"`
	MachineHostAliase       []*CreateContainerGroupMachineHostAliase       `json:"MachineHostAliase,omitempty" name:"MachineHostAliase"`
	RestartPolicy           *string                                        `json:"RestartPolicy,omitempty" name:"RestartPolicy"`
	ImageRegistryCredential []*CreateContainerGroupImageRegistryCredential `json:"ImageRegistryCredential,omitempty" name:"ImageRegistryCredential"`
	Volume                  []*CreateContainerGroupVolume                  `json:"Volume,omitempty" name:"Volume"`
	Container               []*CreateContainerGroupContainer               `json:"Container,omitempty" name:"Container"`
	DnsConfig               *CreateContainerGroupDnsConfig                 `json:"DnsConfig,omitempty" name:"DnsConfig"`
	HostAliase              []*CreateContainerGroupHostAliase              `json:"HostAliase,omitempty" name:"HostAliase"`
	ClusterDns              *string                                        `json:"ClusterDns,omitempty" name:"ClusterDns"`
	ClusterDomain           *string                                        `json:"ClusterDomain,omitempty" name:"ClusterDomain"`
	Label                   []*CreateContainerGroupLabel                   `json:"Label,omitempty" name:"Label"`
	KubeProxy               *CreateContainerGroupKubeProxy                 `json:"KubeProxy,omitempty" name:"KubeProxy"`
	KlogEnabled             *bool                                          `json:"KlogEnabled,omitempty" name:"KlogEnabled"`
	DataDisk                []*CreateContainerGroupDataDisk                `json:"DataDisk,omitempty" name:"DataDisk"`
	ContainerSpec           []*CreateContainerGroupContainerSpec           `json:"ContainerSpec,omitempty" name:"ContainerSpec"`
}

func (r *CreateContainerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateContainerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	ContainerGroupId *string `json:"ContainerGroupId" name:"ContainerGroupId"`
}

func (r *CreateContainerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeContainerGroupRequest struct {
	*ksyunhttp.BaseRequest
	ContainerGroupId []*string                       `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
	MaxResults       *int                            `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker           *int                            `json:"Marker,omitempty" name:"Marker"`
	ProjectId        []*int                          `json:"ProjectId,omitempty" name:"ProjectId"`
	Search           *string                         `json:"Search,omitempty" name:"Search"`
	Filter           []*DescribeContainerGroupFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeContainerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeContainerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	MaxResults      *int    `json:"MaxResults" name:"MaxResults"`
	Marker          *int    `json:"Marker" name:"Marker"`
	TotalCount      *int    `json:"TotalCount" name:"TotalCount"`
	ContainerGroups []struct {
		ContainerGroupId           *string  `json:"ContainerGroupId" name:"ContainerGroupId"`
		ContainerGroupName         *string  `json:"ContainerGroupName" name:"ContainerGroupName"`
		AvailabilityZone           *string  `json:"AvailabilityZone" name:"AvailabilityZone"`
		ChargeType                 *string  `json:"ChargeType" name:"ChargeType"`
		ProjectId                  *int     `json:"ProjectId" name:"ProjectId"`
		KciType                    *string  `json:"KciType" name:"KciType"`
		KciMode                    *string  `json:"KciMode" name:"KciMode"`
		InstanceType               *string  `json:"InstanceType" name:"InstanceType"`
		Status                     *string  `json:"Status" name:"Status"`
		CreateTime                 *string  `json:"CreateTime" name:"CreateTime"`
		Cpu                        *int     `json:"Cpu" name:"Cpu"`
		Memory                     *int     `json:"Memory" name:"Memory"`
		Gpu                        *float64 `json:"Gpu" name:"Gpu"`
		RetainIp                   *bool    `json:"RetainIp" name:"RetainIp"`
		RetainIpHours              *int     `json:"RetainIpHours" name:"RetainIpHours"`
		NetworkInterfaceAttributes []struct {
			NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
			VpcId                *string `json:"VpcId" name:"VpcId"`
			SubnetId             *string `json:"SubnetId" name:"SubnetId"`
			PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			PublicIp             *string `json:"PublicIp" name:"PublicIp"`
			MacAddress           *string `json:"MacAddress" name:"MacAddress"`
			SecurityGroups       []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			} `json:"SecurityGroups"`
		} `json:"NetworkInterfaceAttributes" name:"NetworkInterfaceAttributes"`
		RestartPolicy *string `json:"RestartPolicy" name:"RestartPolicy"`
		SucceededTime *string `json:"SucceededTime" name:"SucceededTime"`
		Labels []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Labels" name:"Labels"`
		DnsConfig struct {
			NameServers []*string `json:"NameServers" name:"NameServers"`
			Searches []*string `json:"Searches" name:"Searches"`
			Options []struct {
				Name *string `json:"Name" name:"Name"`
				Value *string `json:"Value" name:"Value"`
			} `json:"Options"`
		} `json:"DnsConfig" name:"DnsConfig"`
		HostAliases []struct {
			Hostnames []*string `json:"Hostnames" name:"Hostnames"`
			Ip *string `json:"Ip" name:"Ip"`
		} `json:"HostAliases" name:"HostAliases"`
		Volumes []struct {
			Name      *string `json:"Name" name:"Name"`
			Type      *string `json:"Type" name:"Type"`
			NFSVolume struct {
				Path     *string `json:"Path" name:"Path"`
				Server   *string `json:"Server" name:"Server"`
				ReadOnly *bool   `json:"ReadOnly" name:"ReadOnly"`
			} `json:"NFSVolume"`
			HostPathVolume struct {
				Path *string `json:"Path" name:"Path"`
			} `json:"HostPathVolume"`
			EBSVolume struct {
				FsType *string `json:"FsType" name:"FsType"`
				VolumeId *string `json:"VolumeId" name:"VolumeId"`
			} `json:"EBSVolume"`
			ConfigFileVolume struct {
				ConfigFileToPaths []struct {
					Path *string `json:"Path" name:"Path"`
				} `json:"ConfigFileToPaths" name:"ConfigFileToPaths"`
			} `json:"ConfigFileVolume"`
		} `json:"Volumes" name:"Volumes"`
		Containers []struct {
			Name     *string   `json:"Name" name:"Name"`
			Commands []*string `json:"Commands" name:"Commands"`
			Args     []*string `json:"Args" name:"Args"`
			EnvironmentVars []struct {
				Key *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
				ValueFrom struct {
					FieldRef struct {
						FieldPath *string `json:"FieldPath" name:"FieldPath"`
					} `json:"FieldRef"`
				} `json:"ValueFrom" name:"ValueFrom"`
			} `json:"EnvironmentVars"`
			Cpu             *float64 `json:"Cpu" name:"Cpu"`
			Memory          *float64 `json:"Memory" name:"Memory"`
			Gpu             *float64 `json:"Gpu" name:"Gpu"`
			WorkingDir      *string  `json:"WorkingDir" name:"WorkingDir"`
			Image           *string  `json:"Image" name:"Image"`
			ImagePullPolicy *string  `json:"ImagePullPolicy" name:"ImagePullPolicy"`
			RestartCount    *int     `json:"RestartCount" name:"RestartCount"`
			Ports           []struct {
				Protocol *string `json:"Protocol" name:"Protocol"`
				Port *int `json:"Port" name:"Port"`
			} `json:"Ports"`
			VolumeMounts []struct {
				Name     *string `json:"Name" name:"Name"`
				MountPath *string `json:"MountPath" name:"MountPath"`
				ReadOnly *bool   `json:"ReadOnly" name:"ReadOnly"`
			} `json:"VolumeMounts"`
			CurrentState struct {
				StartTime *string `json:"StartTime" name:"StartTime"`
				FinishTime *string `json:"FinishTime" name:"FinishTime"`
				State    *string `json:"State" name:"State"`
				Reason   *string `json:"Reason" name:"Reason"`
				ExitCode *int    `json:"ExitCode" name:"ExitCode"`
			} `json:"CurrentState"`
			PreviousState struct {
				StartTime *string `json:"StartTime" name:"StartTime"`
				FinishTime *string `json:"FinishTime" name:"FinishTime"`
				State    *string `json:"State" name:"State"`
				Reason   *string `json:"Reason" name:"Reason"`
				ExitCode *int    `json:"ExitCode" name:"ExitCode"`
			} `json:"PreviousState"`
			LivenessProbe struct {
				InitialDelaySeconds *int `json:"InitialDelaySeconds" name:"InitialDelaySeconds"`
				PeriodSeconds *int `json:"PeriodSeconds" name:"PeriodSeconds"`
				TimeoutSeconds *int `json:"TimeoutSeconds" name:"TimeoutSeconds"`
				SuccessThreshold *int `json:"SuccessThreshold" name:"SuccessThreshold"`
				FailureThreshold *int `json:"FailureThreshold" name:"FailureThreshold"`
				HttpGet       struct {
					Port   *int    `json:"Port" name:"Port"`
					Path   *string `json:"Path" name:"Path"`
					Scheme *string `json:"Scheme" name:"Scheme"`
				} `json:"HttpGet" name:"HttpGet"`
				Exec struct {
					Commands []*string `json:"Commands" name:"Commands"`
				} `json:"Exec" name:"Exec"`
				TcpSocket struct {
					Port *int `json:"Port" name:"Port"`
				} `json:"TcpSocket" name:"TcpSocket"`
			} `json:"LivenessProbe"`
			ReadinessProbe struct {
				InitialDelaySeconds *int `json:"InitialDelaySeconds" name:"InitialDelaySeconds"`
				PeriodSeconds *int `json:"PeriodSeconds" name:"PeriodSeconds"`
				TimeoutSeconds *int `json:"TimeoutSeconds" name:"TimeoutSeconds"`
				SuccessThreshold *int `json:"SuccessThreshold" name:"SuccessThreshold"`
				FailureThreshold *int `json:"FailureThreshold" name:"FailureThreshold"`
				HttpGet       struct {
					Port   *int    `json:"Port" name:"Port"`
					Path   *string `json:"Path" name:"Path"`
					Scheme *string `json:"Scheme" name:"Scheme"`
				} `json:"HttpGet" name:"HttpGet"`
				Exec struct {
					Commands []*string `json:"Commands" name:"Commands"`
				} `json:"Exec" name:"Exec"`
				TcpSocket struct {
					Port *int `json:"Port" name:"Port"`
				} `json:"TcpSocket" name:"TcpSocket"`
			} `json:"ReadinessProbe"`
		} `json:"Containers" name:"Containers"`
	} `json:"ContainerGroups"`
}

func (r *DescribeContainerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteContainerGroupRequest struct {
	*ksyunhttp.BaseRequest
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
}

func (r *DeleteContainerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteContainerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteContainerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeContainerLogRequest struct {
	*ksyunhttp.BaseRequest
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
	ContainerName    *string `json:"ContainerName,omitempty" name:"ContainerName"`
	Tail             *int    `json:"Tail,omitempty" name:"Tail"`
}

func (r *DescribeContainerLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeContainerLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Content   *string `json:"Content" name:"Content"`
}

func (r *DescribeContainerLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRegionsRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRegionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Regions   []struct {
		Region            *string   `json:"Region" name:"Region"`
		RegionName        *string   `json:"RegionName" name:"RegionName"`
		AvailabilityZones []*string `json:"AvailabilityZones" name:"AvailabilityZones"`
	} `json:"Regions"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ExecContainerCommandRequest struct {
	*ksyunhttp.BaseRequest
	ContainerGroupId *string   `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
	ContainerName    *string   `json:"ContainerName,omitempty" name:"ContainerName"`
	Command          []*string `json:"Command,omitempty" name:"Command"`
	TTY              *bool     `json:"TTY,omitempty" name:"TTY"`
}

func (r *ExecContainerCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ExecContainerCommandResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	WebSocketUri *string `json:"WebSocketUri" name:"WebSocketUri"`
}

func (r *ExecContainerCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecContainerCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeContainerGroupCountRequest struct {
	*ksyunhttp.BaseRequest
	Label *DescribeContainerGroupCountLabel `json:"Label,omitempty" name:"Label"`
}

func (r *DescribeContainerGroupCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeContainerGroupCountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ClusterPodCounts []struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
		Count *int `json:"Count" name:"Count"`
	} `json:"ClusterPodCounts"`
}

func (r *DescribeContainerGroupCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerGroupCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeContainerGroupEventsRequest struct {
	*ksyunhttp.BaseRequest
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
}

func (r *DescribeContainerGroupEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeContainerGroupEventsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Events    []struct {
		FirstTimestamp *string `json:"FirstTimestamp" name:"FirstTimestamp"`
		LastTimestamp *string `json:"LastTimestamp" name:"LastTimestamp"`
		Count  *int    `json:"Count" name:"Count"`
		Type   *string `json:"Type" name:"Type"`
		Reason *string `json:"Reason" name:"Reason"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Events"`
}

func (r *DescribeContainerGroupEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerGroupEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeKciPackagesRequest struct {
	*ksyunhttp.BaseRequest
	ChargeType       *string `json:"ChargeType,omitempty" name:"ChargeType"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	KciType          *string `json:"KciType,omitempty" name:"KciType"`
}

func (r *DescribeKciPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKciPackagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Packages  []struct {
		Cpu *float64   `json:"Cpu" name:"Cpu"`
		Mem []*float64 `json:"Mem" name:"Mem"`
	} `json:"Packages"`
}

func (r *DescribeKciPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKciPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateImageCacheRequest struct {
	*ksyunhttp.BaseRequest
	ImageCacheName          *string                                    `json:"ImageCacheName,omitempty" name:"ImageCacheName"`
	SubnetId                *string                                    `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId         *string                                    `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	ImageCacheSize          *int                                       `json:"ImageCacheSize,omitempty" name:"ImageCacheSize"`
	RetentionDays           *int                                       `json:"RetentionDays,omitempty" name:"RetentionDays"`
	Image                   []*string                                  `json:"Image,omitempty" name:"Image"`
	ImageRegistryCredential []*CreateImageCacheImageRegistryCredential `json:"ImageRegistryCredential,omitempty" name:"ImageRegistryCredential"`
	ImageCacheType          *string                                    `json:"ImageCacheType,omitempty" name:"ImageCacheType"`
	EnableWarm              *bool                                      `json:"EnableWarm,omitempty" name:"EnableWarm"`
}

func (r *CreateImageCacheRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateImageCacheResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	ImageCacheId *string `json:"ImageCacheId" name:"ImageCacheId"`
}

func (r *CreateImageCacheResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteImageCacheRequest struct {
	*ksyunhttp.BaseRequest
	ImageCacheId *string `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
}

func (r *DeleteImageCacheRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteImageCacheResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteImageCacheResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeImageCacheRequest struct {
	*ksyunhttp.BaseRequest
	ImageCacheId   []*string `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
	ImageCacheName *string   `json:"ImageCacheName,omitempty" name:"ImageCacheName"`
	Image          *string   `json:"Image,omitempty" name:"Image"`
	Marker         *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults     *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeImageCacheRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeImageCacheResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	MaxResults  *int    `json:"MaxResults" name:"MaxResults"`
	Marker      *int    `json:"Marker" name:"Marker"`
	TotalCount  *int    `json:"TotalCount" name:"TotalCount"`
	ImageCaches []struct {
		ImageCacheId     *string   `json:"ImageCacheId" name:"ImageCacheId"`
		ImageCacheName   *string   `json:"ImageCacheName" name:"ImageCacheName"`
		ExpireTime       *string   `json:"ExpireTime" name:"ExpireTime"`
		CreateTime       *string   `json:"CreateTime" name:"CreateTime"`
		Status           *string   `json:"Status" name:"Status"`
		ImageCacheSize   *int      `json:"ImageCacheSize" name:"ImageCacheSize"`
		SnapshotId       *string   `json:"SnapshotId" name:"SnapshotId"`
		ContainerGroupId *string   `json:"ContainerGroupId" name:"ContainerGroupId"`
		Images           []*string `json:"Images" name:"Images"`
		Reason           *string   `json:"Reason" name:"Reason"`
		ImageCacheType   *string   `json:"ImageCacheType" name:"ImageCacheType"`
	} `json:"ImageCaches"`
}

func (r *DescribeImageCacheResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type MatchImageCacheRequest struct {
	*ksyunhttp.BaseRequest
	Image []*string `json:"Image,omitempty" name:"Image"`
}

func (r *MatchImageCacheRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type MatchImageCacheResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ImageCaches []struct {
		ImageCacheId *string `json:"ImageCacheId" name:"ImageCacheId"`
		ImageCacheName *string `json:"ImageCacheName" name:"ImageCacheName"`
	} `json:"ImageCaches"`
}

func (r *MatchImageCacheResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MatchImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeImageCacheEventRequest struct {
	*ksyunhttp.BaseRequest
	ImageCacheId *string `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
}

func (r *DescribeImageCacheEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeImageCacheEventResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Events    []struct {
		FirstTimestamp *string `json:"FirstTimestamp" name:"FirstTimestamp"`
		LastTimestamp *string `json:"LastTimestamp" name:"LastTimestamp"`
		Count  *int    `json:"Count" name:"Count"`
		Type   *string `json:"Type" name:"Type"`
		Reason *string `json:"Reason" name:"Reason"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Events"`
}

func (r *DescribeImageCacheEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageCacheEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateImageCacheRequest struct {
	*ksyunhttp.BaseRequest
	ImageCacheId            *string                                    `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
	ImageCacheName          *string                                    `json:"ImageCacheName,omitempty" name:"ImageCacheName"`
	SubnetId                *string                                    `json:"SubnetId,omitempty" name:"SubnetId"`
	SecurityGroupId         *string                                    `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	ImageCacheSize          *int                                       `json:"ImageCacheSize,omitempty" name:"ImageCacheSize"`
	RetentionDays           *int                                       `json:"RetentionDays,omitempty" name:"RetentionDays"`
	Image                   []*string                                  `json:"Image,omitempty" name:"Image"`
	ImageCacheType          *string                                    `json:"ImageCacheType,omitempty" name:"ImageCacheType"`
	EnableWarm              *bool                                      `json:"EnableWarm,omitempty" name:"EnableWarm"`
	ImageRegistryCredential []*UpdateImageCacheImageRegistryCredential `json:"ImageRegistryCredential,omitempty" name:"ImageRegistryCredential"`
}

func (r *UpdateImageCacheRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateImageCacheResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	ImageCacheId *string `json:"ImageCacheId" name:"ImageCacheId"`
}

func (r *UpdateImageCacheResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

