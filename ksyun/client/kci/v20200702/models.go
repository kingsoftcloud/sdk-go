package v20200702
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type CreateContainerGroupOption struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateContainerGroupMachineHostAliase struct {
    Ip *string `json:"Ip,omitempty" name:"Ip"`
    Hostname []*string `json:"Hostname,omitempty" name:"Hostname"`
}

type CreateContainerGroupImageRegistryCredential struct {
    Server *string `json:"Server,omitempty" name:"Server"`
    Username *string `json:"Username,omitempty" name:"Username"`
    Password *string `json:"Password,omitempty" name:"Password"`
}

type CreateContainerGroupVolume struct {
    Type *string `json:"Type,omitempty" name:"Type"`
    Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateContainerGroupConfigFileToPath struct {
    Path *string `json:"Path,omitempty" name:"Path"`
    Content *string `json:"Content,omitempty" name:"Content"`
    Mode *int `json:"mode,omitempty" name:"mode"`
}

type CreateContainerGroupContainer struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Command []*string `json:"Command,omitempty" name:"Command"`
    Arg []*string `json:"Arg,omitempty" name:"Arg"`
    Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`
    Memory *string `json:"Memory,omitempty" name:"Memory"`
    Gpu *float64 `json:"Gpu,omitempty" name:"Gpu"`
    WorkingDir *string `json:"WorkingDir,omitempty" name:"WorkingDir"`
    Image *string `json:"Image,omitempty" name:"Image"`
    ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" name:"ImagePullPolicy"`
EnvironmentVar   []struct {
                Key *string `json:"Key,omitempty" name:"Key"`
                        Value *string `json:"Value,omitempty" name:"Value"`
                    } `json:"EnvironmentVar,omitempty" name:"EnvironmentVar"`
Port   []struct {
                Port *int `json:"Port,omitempty" name:"Port"`
                        Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
        } `json:"Port,omitempty" name:"Port"`
VolumeMount   []struct {
                Name *string `json:"Name,omitempty" name:"Name"`
                        MountPath *string `json:"MountPath,omitempty" name:"MountPath"`
                    } `json:"VolumeMount,omitempty" name:"VolumeMount"`
}

type CreateContainerGroupEnvironmentVar struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateContainerGroupPort struct {
    Port *int `json:"Port,omitempty" name:"Port"`
    Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type CreateContainerGroupVolumeMount struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    MountPath *string `json:"MountPath,omitempty" name:"MountPath"`
    ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type CreateContainerGroupOption struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateContainerGroupHostAliase struct {
    Ip *string `json:"Ip,omitempty" name:"Ip"`
    Hostname []*string `json:"Hostname,omitempty" name:"Hostname"`
}

type CreateContainerGroupLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateContainerGroupDataDisk struct {
    VolumeName *string `json:"VolumeName,omitempty" name:"VolumeName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    Size *int `json:"Size,omitempty" name:"Size"`
    DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
    SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type CreateContainerGroupContainerSpec struct {
    RequestCpu *float64 `json:"RequestCpu,omitempty" name:"RequestCpu"`
    RequestMem *float64 `json:"RequestMem,omitempty" name:"RequestMem"`
    LimitCpu *float64 `json:"LimitCpu,omitempty" name:"LimitCpu"`
    LimitMem *float64 `json:"LimitMem,omitempty" name:"LimitMem"`
}

type DescribeContainerGroupFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateImageCacheImageRegistryCredential struct {
    Server *string `json:"Server,omitempty" name:"Server"`
    Username *string `json:"Username,omitempty" name:"Username"`
    Password *string `json:"Password,omitempty" name:"Password"`
}


type CreateContainerGroupRequest struct {
    *ksyunhttp.BaseRequest
    ContainerGroupName *string `json:"ContainerGroupName,omitempty" name:"ContainerGroupName"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    MultiSubnetId []*string `json:"MultiSubnetId,omitempty" name:"MultiSubnetId"`
    SecurityGroupId []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    KciType *string `json:"KciType,omitempty" name:"KciType"`
    InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
    InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    SpotStrategy *string `json:"SpotStrategy,omitempty" name:"SpotStrategy"`
    ProjectId *int `json:"ProjectId,omitempty" name:"ProjectId"`
    Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`
    Memory *float64 `json:"Memory,omitempty" name:"Memory"`
    Gpu *float64 `json:"Gpu,omitempty" name:"Gpu"`
    KubeConfig *string `json:"KubeConfig,omitempty" name:"KubeConfig"`
    RetainIp *bool `json:"RetainIp,omitempty" name:"RetainIp"`
    RetainIpHours *int `json:"RetainIpHours,omitempty" name:"RetainIpHours"`
    EipAllocationId *string `json:"EipAllocationId,omitempty" name:"EipAllocationId"`
    AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" name:"AutoMatchImageCache"`
    ImageCacheId *string `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
    MachineHostAliase []*CreateContainerGroupMachineHostAliase `json:"MachineHostAliase,omitempty" name:"MachineHostAliase"`
    RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`
    ImageRegistryCredential []*CreateContainerGroupImageRegistryCredential `json:"ImageRegistryCredential,omitempty" name:"ImageRegistryCredential"`
    Volume []*CreateContainerGroupVolume `json:"Volume,omitempty" name:"Volume"`
    Container []*CreateContainerGroupContainer `json:"Container,omitempty" name:"Container"`
    HostAliase []*CreateContainerGroupHostAliase `json:"HostAliase,omitempty" name:"HostAliase"`
    ClusterDns *string `json:"ClusterDns,omitempty" name:"ClusterDns"`
    ClusterDomain *string `json:"ClusterDomain,omitempty" name:"ClusterDomain"`
    Label []*CreateContainerGroupLabel `json:"Label,omitempty" name:"Label"`
    KlogEnabled *bool `json:"KlogEnabled,omitempty" name:"KlogEnabled"`
    DataDisk []*CreateContainerGroupDataDisk `json:"DataDisk,omitempty" name:"DataDisk"`
    ContainerSpec []*CreateContainerGroupContainerSpec `json:"ContainerSpec,omitempty" name:"ContainerSpec"`
}

func (r *CreateContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContainerGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateContainerGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateContainerGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    ContainerGroupId []*string `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    ProjectId []*int `json:"ProjectId,omitempty" name:"ProjectId"`
    Search *string `json:"Search,omitempty" name:"Search"`
    Filter []*DescribeContainerGroupFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeContainerGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    MaxResults *int `json:"MaxResults" name:"MaxResults"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	ContainerGroups []struct {
		ContainerGroupId *string `json:"ContainerGroupId"`
		ContainerGroupName *string `json:"ContainerGroupName"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		ChargeType *string `json:"ChargeType"`
		ProjectId *int `json:"ProjectId"`
		KciType *string `json:"KciType"`
		KciMode *string `json:"KciMode"`
		InstanceType *string `json:"InstanceType"`
		Status *string `json:"Status"`
		CreateTime *string `json:"CreateTime"`
		Cpu *int `json:"Cpu"`
		Memory *int `json:"Memory"`
		RetainIp *bool `json:"RetainIp"`
		RetainIpHours *int `json:"RetainIpHours"`
		NetworkInterfaceAttributes []struct {
					NetworkInterfaceId *string `json:"NetworkInterfaceId"`
					NetworkInterfaceType *string `json:"NetworkInterfaceType"`
					VpcId *string `json:"VpcId"`
					SubnetId *string `json:"SubnetId"`
					PrivateIpAddress *string `json:"PrivateIpAddress"`
					PublicIp *string `json:"PublicIp"`
					MacAddress *string `json:"MacAddress"`
				SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
				} `json:"SecurityGroups"`
			} `json:"NetworkInterfaceAttributes"`
			RestartPolicy *string `json:"RestartPolicy"`
			SucceededTime *string `json:"SucceededTime"`
			Labels []struct {
						Key *string `json:"Key"`
						Value *string `json:"Value"`
				} `json:"Labels"`
				DnsConfig struct {
					NameServers []struct {
					} `json:"NameServers"`
					Searches []struct {
					} `json:"Searches"`
					Options []struct {
						Name *string `json:"Name"`
						Value *string `json:"Value"`
					} `json:"Options"`
				} `json:"DnsConfig"`
				HostAliases []struct {
						Hostnames []struct {
						} `json:"Hostnames"`
							Ip *string `json:"Ip"`
					} `json:"HostAliases"`
					Volumes []struct {
								Name *string `json:"Name"`
								Type *string `json:"Type"`
							NFSVolume struct {
								Path *string `json:"Path"`
								Server *string `json:"Server"`
								ReadOnly *bool `json:"ReadOnly"`
							} `json:"NFSVolume"`
							HostPathVolume struct {
								Path *string `json:"Path"`
							} `json:"HostPathVolume"`
							EBSVolume struct {
								FsType *string `json:"FsType"`
								VolumeId *string `json:"VolumeId"`
							} `json:"EBSVolume"`
							ConfigFileVolume struct {
								ConfigFileToPaths []struct {
											Path *string `json:"Path"`
									} `json:"ConfigFileToPaths"`
								} `json:"ConfigFileVolume"`
						} `json:"Volumes"`
						Containers []struct {
									Name *string `json:"Name"`
								Commands []struct {
								} `json:"Commands"`
								Args []struct {
								} `json:"Args"`
								EnvironmentVars []struct {
									Key *string `json:"Key"`
									Value *string `json:"Value"`
									ValueFrom struct {
										FieldRef struct {
											FieldPath *string `json:"FieldPath"`
										} `json:"FieldRef"`
									} `json:"ValueFrom"`
								} `json:"EnvironmentVars"`
									WorkingDir *string `json:"WorkingDir"`
									Image *string `json:"Image"`
									ImagePullPolicy *string `json:"ImagePullPolicy"`
									RestartCount *int `json:"RestartCount"`
								Ports []struct {
									Protocol *string `json:"Protocol"`
									Port *int `json:"Port"`
								} `json:"Ports"`
								VolumeMounts []struct {
									Name *string `json:"Name"`
									MountPath *string `json:"MountPath"`
									ReadOnly *bool `json:"ReadOnly"`
								} `json:"VolumeMounts"`
								CurrentState struct {
									StartTime *string `json:"StartTime"`
									FinishTime *string `json:"FinishTime"`
									State *string `json:"State"`
									Reason *string `json:"Reason"`
									ExitCode *int `json:"ExitCode"`
								} `json:"CurrentState"`
								PreviousState struct {
									StartTime *string `json:"StartTime"`
									FinishTime *string `json:"FinishTime"`
									State *string `json:"State"`
									Reason *string `json:"Reason"`
									ExitCode *int `json:"ExitCode"`
								} `json:"PreviousState"`
								LivenessProbe struct {
									InitialDelaySeconds *int `json:"InitialDelaySeconds"`
									PeriodSeconds *int `json:"PeriodSeconds"`
									TimeoutSeconds *int `json:"TimeoutSeconds"`
									SuccessThreshold *int `json:"SuccessThreshold"`
									FailureThreshold *int `json:"FailureThreshold"`
									HttpGet struct {
											Port *int `json:"Port"`
											Path *string `json:"Path"`
											Scheme *string `json:"Scheme"`
									} `json:"HttpGet"`
									Exec struct {
										Commands []struct {
										} `json:"Commands"`
									} `json:"Exec"`
									TcpSocket struct {
											Port *int `json:"Port"`
									} `json:"TcpSocket"`
								} `json:"LivenessProbe"`
								ReadinessProbe struct {
									InitialDelaySeconds *int `json:"InitialDelaySeconds"`
									PeriodSeconds *int `json:"PeriodSeconds"`
									TimeoutSeconds *int `json:"TimeoutSeconds"`
									SuccessThreshold *int `json:"SuccessThreshold"`
									FailureThreshold *int `json:"FailureThreshold"`
									HttpGet struct {
											Port *int `json:"Port"`
											Path *string `json:"Path"`
											Scheme *string `json:"Scheme"`
									} `json:"HttpGet"`
									Exec struct {
										Commands []struct {
										} `json:"Commands"`
									} `json:"Exec"`
									TcpSocket struct {
											Port *int `json:"Port"`
									} `json:"TcpSocket"`
								} `json:"ReadinessProbe"`
							} `json:"Containers"`
						} `json:"ContainerGroups"`
}

func (r *DescribeContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupListRequest struct {
    *ksyunhttp.BaseRequest
    Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeContainerGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeContainerGroupListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeContainerGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupListResponse) FromJsonString(s string) error {
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

func (r *DeleteContainerGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteContainerGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
    Tail *int `json:"Tail,omitempty" name:"Tail"`
}

func (r *DescribeContainerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerLogRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeContainerLogRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerLogResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Content *string `json:"Content" name:"Content"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
	Regions []struct {
		Region *string `json:"Region"`
		RegionName *string `json:"RegionName"`
		AvailabilityZones []struct {
			} `json:"AvailabilityZones"`
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
    ContainerGroupId *string `json:"ContainerGroupId,omitempty" name:"ContainerGroupId"`
    ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
    Command []*string `json:"Command,omitempty" name:"Command"`
    TTY *bool `json:"TTY,omitempty" name:"TTY"`
}

func (r *ExecContainerCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecContainerCommandRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ExecContainerCommandRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ExecContainerCommandResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
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
}

func (r *DescribeContainerGroupCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupCountRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeContainerGroupCountRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupCountResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	ClusterPodCounts []struct {
		ClusterId *string `json:"ClusterId"`
		Count *int `json:"Count"`
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

func (r *DescribeContainerGroupEventsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeContainerGroupEventsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupEventsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Events []struct {
		FirstTimestamp *string `json:"FirstTimestamp"`
		LastTimestamp *string `json:"LastTimestamp"`
		Count *int `json:"Count"`
		Type *string `json:"Type"`
		Reason *string `json:"Reason"`
		Message *string `json:"Message"`
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
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    KciType *string `json:"KciType,omitempty" name:"KciType"`
}

func (r *DescribeKciPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKciPackagesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeKciPackagesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKciPackagesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Packages []struct {
		Mem []struct {
			} `json:"Mem"`
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
    ImageCacheName *string `json:"ImageCacheName,omitempty" name:"ImageCacheName"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    ImageCacheSize *int `json:"ImageCacheSize,omitempty" name:"ImageCacheSize"`
    RetentionDays *int `json:"RetentionDays,omitempty" name:"RetentionDays"`
    Image []*string `json:"Image,omitempty" name:"Image"`
    ImageRegistryCredential []*CreateImageCacheImageRegistryCredential `json:"ImageRegistryCredential,omitempty" name:"ImageRegistryCredential"`
    ImageCacheType *string `json:"ImageCacheType,omitempty" name:"ImageCacheType"`
    EnableWarm *bool `json:"EnableWarm,omitempty" name:"EnableWarm"`
}

func (r *CreateImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageCacheRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateImageCacheRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageCacheResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
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

func (r *DeleteImageCacheRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteImageCacheRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
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
    ImageCacheId []*string `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
    ImageCacheName *string `json:"ImageCacheName,omitempty" name:"ImageCacheName"`
    Image *string `json:"Image,omitempty" name:"Image"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageCacheRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeImageCacheRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageCacheResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    MaxResults *int `json:"MaxResults" name:"MaxResults"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	ImageCaches []struct {
		ImageCacheId *string `json:"ImageCacheId"`
		ImageCacheName *string `json:"ImageCacheName"`
		ExpireTime *string `json:"ExpireTime"`
		CreateTime *string `json:"CreateTime"`
		Status *string `json:"Status"`
		ImageCacheSize *int `json:"ImageCacheSize"`
		SnapshotId *string `json:"SnapshotId"`
		ContainerGroupId *string `json:"ContainerGroupId"`
		Images []struct {
			} `json:"Images"`
			Reason *string `json:"Reason"`
			ImageCacheType *string `json:"ImageCacheType"`
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

func (r *MatchImageCacheRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "MatchImageCacheRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type MatchImageCacheResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	ImageCaches []struct {
		ImageCacheId *string `json:"ImageCacheId"`
		ImageCacheName *string `json:"ImageCacheName"`
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

func (r *DescribeImageCacheEventRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeImageCacheEventRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageCacheEventResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Events []struct {
		FirstTimestamp *string `json:"FirstTimestamp"`
		LastTimestamp *string `json:"LastTimestamp"`
		Count *int `json:"Count"`
		Type *string `json:"Type"`
		Reason *string `json:"Reason"`
		Message *string `json:"Message"`
	} `json:"Events"`
}

func (r *DescribeImageCacheEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageCacheEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

