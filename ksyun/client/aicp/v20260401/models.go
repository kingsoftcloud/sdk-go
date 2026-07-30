package v20260401

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateSandboxTemplateEnvs struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateSandboxTemplateImageConfig struct {
	ImageSource        *string `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageEndpoint      *string `json:"ImageEndpoint,omitempty" name:"ImageEndpoint"`
	ImageNamespace     *string `json:"ImageNamespace,omitempty" name:"ImageNamespace"`
	ImageName          *string `json:"ImageName,omitempty" name:"ImageName"`
	ImageUrl           *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	ImageTag           *string `json:"ImageTag,omitempty" name:"ImageTag"`
	CredentialUsername *string `json:"CredentialUsername,omitempty" name:"CredentialUsername"`
	CredentialPwd      *string `json:"CredentialPwd,omitempty" name:"CredentialPwd"`
	RegistryInstanceId *string `json:"RegistryInstanceId,omitempty" name:"RegistryInstanceId"`
}
type CreateSandboxTemplateSkillConfig struct {
	SkillEnable       *bool     `json:"SkillEnable,omitempty" name:"SkillEnable"`
	SkillSpaceIds     []*string `json:"SkillSpaceIds,omitempty" name:"SkillSpaceIds"`
	PublicSkillEnable *bool     `json:"PublicSkillEnable,omitempty" name:"PublicSkillEnable"`
}
type CreateSandboxTemplateNetworkConfigVpcConfiguration struct {
	VpcId    *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
type CreateSandboxTemplateNetworkConfig struct {
	PublicNetworkEnable        *bool                                               `json:"PublicNetworkEnable,omitempty" name:"PublicNetworkEnable"`
	PrivateNetworkEnable       *bool                                               `json:"PrivateNetworkEnable,omitempty" name:"PrivateNetworkEnable"`
	SharedInternetAccessEnable *bool                                               `json:"SharedInternetAccessEnable,omitempty" name:"SharedInternetAccessEnable"`
	VpcConfiguration           *CreateSandboxTemplateNetworkConfigVpcConfiguration `json:"VpcConfiguration,omitempty" name:"VpcConfiguration"`
}
type CreateSandboxTemplateKlogConfig struct {
	KlogEnable *bool `json:"KlogEnable,omitempty" name:"KlogEnable"`
}
type CreateSandboxTemplateKpfsMountConfigKpfsMountPoints struct {
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	RemotePath     *string `json:"RemotePath,omitempty" name:"RemotePath"`
	LocalMountPath *string `json:"LocalMountPath,omitempty" name:"LocalMountPath"`
	ReadOnly       *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type CreateSandboxTemplateKpfsMountConfig struct {
	KpfsEnable      *bool                                                  `json:"KpfsEnable,omitempty" name:"KpfsEnable"`
	KpfsMountPoints []*CreateSandboxTemplateKpfsMountConfigKpfsMountPoints `json:"KpfsMountPoints,omitempty" name:"KpfsMountPoints"`
}
type CreateSandboxTemplateKs3MountConfigKs3MountPoints struct {
	BucketName     *string `json:"BucketName,omitempty" name:"BucketName"`
	RemotePath     *string `json:"RemotePath,omitempty" name:"RemotePath"`
	LocalMountPath *string `json:"LocalMountPath,omitempty" name:"LocalMountPath"`
	ReadOnly       *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type CreateSandboxTemplateKs3MountConfig struct {
	Ks3Enable      *bool                                                `json:"Ks3Enable,omitempty" name:"Ks3Enable"`
	Ks3MountPoints []*CreateSandboxTemplateKs3MountConfigKs3MountPoints `json:"Ks3MountPoints,omitempty" name:"Ks3MountPoints"`
}
type CreateSandboxTemplateKecConfigInstanceSpecsSystemDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *int    `json:"Size,omitempty" name:"Size"`
}
type CreateSandboxTemplateKecConfigInstanceSpecsDataDisks struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}
type CreateSandboxTemplateKecConfigInstanceSpecs struct {
	InstanceType *string                                                 `json:"InstanceType,omitempty" name:"InstanceType"`
	SystemDisk   *CreateSandboxTemplateKecConfigInstanceSpecsSystemDisk  `json:"SystemDisk,omitempty" name:"SystemDisk"`
	DataDisks    []*CreateSandboxTemplateKecConfigInstanceSpecsDataDisks `json:"DataDisks,omitempty" name:"DataDisks"`
}
type CreateSandboxTemplateKecConfig struct {
	KecEnable     *bool                                          `json:"KecEnable,omitempty" name:"KecEnable"`
	InstanceSpecs []*CreateSandboxTemplateKecConfigInstanceSpecs `json:"InstanceSpecs,omitempty" name:"InstanceSpecs"`
}
type CreateSandboxTemplatePreheatConfig struct {
	PreheatEnable *bool `json:"PreheatEnable,omitempty" name:"PreheatEnable"`
	PreheatNumber *int  `json:"PreheatNumber,omitempty" name:"PreheatNumber"`
}
type UpdateSandboxTemplateEnvs struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type UpdateSandboxTemplateImageConfig struct {
	ImageSource        *string `json:"ImageSource,omitempty" name:"ImageSource"`
	ImageEndpoint      *string `json:"ImageEndpoint,omitempty" name:"ImageEndpoint"`
	ImageNamespace     *string `json:"ImageNamespace,omitempty" name:"ImageNamespace"`
	ImageName          *string `json:"ImageName,omitempty" name:"ImageName"`
	ImageUrl           *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	ImageTag           *string `json:"ImageTag,omitempty" name:"ImageTag"`
	CredentialUsername *string `json:"CredentialUsername,omitempty" name:"CredentialUsername"`
	CredentialPwd      *string `json:"CredentialPwd,omitempty" name:"CredentialPwd"`
	RegistryInstanceId *string `json:"RegistryInstanceId,omitempty" name:"RegistryInstanceId"`
}
type UpdateSandboxTemplateSkillConfig struct {
	SkillEnable       *bool     `json:"SkillEnable,omitempty" name:"SkillEnable"`
	SkillSpaceIds     []*string `json:"SkillSpaceIds,omitempty" name:"SkillSpaceIds"`
	PublicSkillEnable *bool     `json:"PublicSkillEnable,omitempty" name:"PublicSkillEnable"`
}
type UpdateSandboxTemplateNetworkConfigVpcConfiguration struct {
	VpcId    *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
type UpdateSandboxTemplateNetworkConfig struct {
	PublicNetworkEnable        *bool                                               `json:"PublicNetworkEnable,omitempty" name:"PublicNetworkEnable"`
	PrivateNetworkEnable       *bool                                               `json:"PrivateNetworkEnable,omitempty" name:"PrivateNetworkEnable"`
	SharedInternetAccessEnable *bool                                               `json:"SharedInternetAccessEnable,omitempty" name:"SharedInternetAccessEnable"`
	VpcConfiguration           *UpdateSandboxTemplateNetworkConfigVpcConfiguration `json:"VpcConfiguration,omitempty" name:"VpcConfiguration"`
}
type UpdateSandboxTemplateKlogConfig struct {
	KlogEnable *bool `json:"KlogEnable,omitempty" name:"KlogEnable"`
}
type UpdateSandboxTemplateKpfsMountConfigKpfsMountPoints struct {
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	RemotePath     *string `json:"RemotePath,omitempty" name:"RemotePath"`
	LocalMountPath *string `json:"LocalMountPath,omitempty" name:"LocalMountPath"`
	ReadOnly       *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type UpdateSandboxTemplateKpfsMountConfig struct {
	KpfsEnable      *bool                                                  `json:"KpfsEnable,omitempty" name:"KpfsEnable"`
	KpfsMountPoints []*UpdateSandboxTemplateKpfsMountConfigKpfsMountPoints `json:"KpfsMountPoints,omitempty" name:"KpfsMountPoints"`
}
type UpdateSandboxTemplateKs3MountConfigKs3MountPoints struct {
	BucketName     *string `json:"BucketName,omitempty" name:"BucketName"`
	RemotePath     *string `json:"RemotePath,omitempty" name:"RemotePath"`
	LocalMountPath *string `json:"LocalMountPath,omitempty" name:"LocalMountPath"`
	ReadOnly       *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type UpdateSandboxTemplateKs3MountConfig struct {
	Ks3Enable      *bool                                                `json:"Ks3Enable,omitempty" name:"Ks3Enable"`
	Ks3MountPoints []*UpdateSandboxTemplateKs3MountConfigKs3MountPoints `json:"Ks3MountPoints,omitempty" name:"Ks3MountPoints"`
}
type UpdateSandboxTemplateKecConfigInstanceSpecsSystemDisk struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Size *int    `json:"Size,omitempty" name:"Size"`
}
type UpdateSandboxTemplateKecConfigInstanceSpecsDataDisks struct {
	Type               *string `json:"Type,omitempty" name:"Type"`
	Size               *int    `json:"Size,omitempty" name:"Size"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}
type UpdateSandboxTemplateKecConfigInstanceSpecs struct {
	InstanceType *string                                                 `json:"InstanceType,omitempty" name:"InstanceType"`
	SystemDisk   *UpdateSandboxTemplateKecConfigInstanceSpecsSystemDisk  `json:"SystemDisk,omitempty" name:"SystemDisk"`
	DataDisks    []*UpdateSandboxTemplateKecConfigInstanceSpecsDataDisks `json:"DataDisks,omitempty" name:"DataDisks"`
}
type UpdateSandboxTemplateKecConfig struct {
	KecEnable     *bool                                          `json:"KecEnable,omitempty" name:"KecEnable"`
	InstanceSpecs []*UpdateSandboxTemplateKecConfigInstanceSpecs `json:"InstanceSpecs,omitempty" name:"InstanceSpecs"`
}
type UpdateSandboxTemplatePreheatConfig struct {
	PreheatEnable *bool `json:"PreheatEnable,omitempty" name:"PreheatEnable"`
	PreheatNumber *int  `json:"PreheatNumber,omitempty" name:"PreheatNumber"`
}
type StartSandboxInstanceKs3MountConfigKs3MountPoints struct {
	BucketName     *string `json:"BucketName,omitempty" name:"BucketName"`
	RemotePath     *string `json:"RemotePath,omitempty" name:"RemotePath"`
	LocalMountPath *string `json:"LocalMountPath,omitempty" name:"LocalMountPath"`
	ReadOnly       *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type StartSandboxInstanceKs3MountConfig struct {
	Ks3Enable      *bool                                               `json:"Ks3Enable,omitempty" name:"Ks3Enable"`
	Ks3MountPoints []*StartSandboxInstanceKs3MountConfigKs3MountPoints `json:"Ks3MountPoints,omitempty" name:"Ks3MountPoints"`
}
type StartSandboxInstanceKpfsMountConfigKpfsMountPoints struct {
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`
	RemotePath     *string `json:"RemotePath,omitempty" name:"RemotePath"`
	LocalMountPath *string `json:"LocalMountPath,omitempty" name:"LocalMountPath"`
	ReadOnly       *bool   `json:"ReadOnly,omitempty" name:"ReadOnly"`
}
type StartSandboxInstanceKpfsMountConfig struct {
	KpfsEnable      *bool                                                 `json:"KpfsEnable,omitempty" name:"KpfsEnable"`
	KpfsMountPoints []*StartSandboxInstanceKpfsMountConfigKpfsMountPoints `json:"KpfsMountPoints,omitempty" name:"KpfsMountPoints"`
}
type StartSandboxInstanceEnvs struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateSandboxTemplateRequest struct {
	*ksyunhttp.BaseRequest
	TemplateName     *string                               `json:"TemplateName,omitempty" name:"TemplateName"`
	Description      *string                               `json:"Description,omitempty" name:"Description"`
	TemplateCategory *string                               `json:"TemplateCategory,omitempty" name:"TemplateCategory"`
	TemplateType     *string                               `json:"TemplateType,omitempty" name:"TemplateType"`
	Command          *string                               `json:"Command,omitempty" name:"Command"`
	Cpu              *int                                  `json:"Cpu,omitempty" name:"Cpu"`
	Memory           *int                                  `json:"Memory,omitempty" name:"Memory"`
	Ports            []*int                                `json:"Ports,omitempty" name:"Ports"`
	Envs             []*CreateSandboxTemplateEnvs          `json:"Envs,omitempty" name:"Envs"`
	ImageConfig      *CreateSandboxTemplateImageConfig     `json:"ImageConfig,omitempty" name:"ImageConfig"`
	SkillConfig      *CreateSandboxTemplateSkillConfig     `json:"SkillConfig,omitempty" name:"SkillConfig"`
	NetworkConfig    *CreateSandboxTemplateNetworkConfig   `json:"NetworkConfig,omitempty" name:"NetworkConfig"`
	KlogConfig       *CreateSandboxTemplateKlogConfig      `json:"KlogConfig,omitempty" name:"KlogConfig"`
	KpfsMountConfig  *CreateSandboxTemplateKpfsMountConfig `json:"KpfsMountConfig,omitempty" name:"KpfsMountConfig"`
	Ks3MountConfig   *CreateSandboxTemplateKs3MountConfig  `json:"Ks3MountConfig,omitempty" name:"Ks3MountConfig"`
	AccessKey        *string                               `json:"AccessKey,omitempty" name:"AccessKey"`
	SecretAccessKey  *string                               `json:"SecretAccessKey,omitempty" name:"SecretAccessKey"`
	KecConfig        *CreateSandboxTemplateKecConfig       `json:"KecConfig,omitempty" name:"KecConfig"`
	PreheatConfig    *CreateSandboxTemplatePreheatConfig   `json:"PreheatConfig,omitempty" name:"PreheatConfig"`
	InstanceQuota    *int                                  `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
}

func (r *CreateSandboxTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSandboxTemplateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TemplateId *string `json:"TemplateId" name:"TemplateId"`
	} `json:"Data"`
}

func (r *CreateSandboxTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSandboxTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSandboxTemplateRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId       *string                               `json:"TemplateId,omitempty" name:"TemplateId"`
	TemplateName     *string                               `json:"TemplateName,omitempty" name:"TemplateName"`
	Description      *string                               `json:"Description,omitempty" name:"Description"`
	Envs             []*UpdateSandboxTemplateEnvs          `json:"Envs,omitempty" name:"Envs"`
	TemplateType     *string                               `json:"TemplateType,omitempty" name:"TemplateType"`
	TemplateCategory *string                               `json:"TemplateCategory,omitempty" name:"TemplateCategory"`
	Command          *string                               `json:"Command,omitempty" name:"Command"`
	Ports            []*int                                `json:"Ports,omitempty" name:"Ports"`
	ImageConfig      *UpdateSandboxTemplateImageConfig     `json:"ImageConfig,omitempty" name:"ImageConfig"`
	SkillConfig      *UpdateSandboxTemplateSkillConfig     `json:"SkillConfig,omitempty" name:"SkillConfig"`
	NetworkConfig    *UpdateSandboxTemplateNetworkConfig   `json:"NetworkConfig,omitempty" name:"NetworkConfig"`
	KlogConfig       *UpdateSandboxTemplateKlogConfig      `json:"KlogConfig,omitempty" name:"KlogConfig"`
	KpfsMountConfig  *UpdateSandboxTemplateKpfsMountConfig `json:"KpfsMountConfig,omitempty" name:"KpfsMountConfig"`
	Ks3MountConfig   *UpdateSandboxTemplateKs3MountConfig  `json:"Ks3MountConfig,omitempty" name:"Ks3MountConfig"`
	AccessKey        *string                               `json:"AccessKey,omitempty" name:"AccessKey"`
	SecretAccessKey  *string                               `json:"SecretAccessKey,omitempty" name:"SecretAccessKey"`
	KecConfig        *UpdateSandboxTemplateKecConfig       `json:"KecConfig,omitempty" name:"KecConfig"`
	PreheatConfig    *UpdateSandboxTemplatePreheatConfig   `json:"PreheatConfig,omitempty" name:"PreheatConfig"`
	InstanceQuota    *int                                  `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
	Cpu              *int                                  `json:"Cpu,omitempty" name:"Cpu"`
	Memory           *int                                  `json:"Memory,omitempty" name:"Memory"`
}

func (r *UpdateSandboxTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateSandboxTemplateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TemplateId *string `json:"TemplateId" name:"TemplateId"`
	} `json:"Data"`
}

func (r *UpdateSandboxTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSandboxTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSandboxInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DeleteSandboxInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSandboxInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		SuccessDeleted []*string `json:"SuccessDeleted" name:"SuccessDeleted"`
		FailedDeleted  []struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			FailedReason *string `json:"FailedReason" name:"FailedReason"`
		} `json:"FailedDeleted" name:"FailedDeleted"`
	} `json:"Data"`
}

func (r *DeleteSandboxInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSandboxInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetSandboxInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetSandboxInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceId          *string `json:"InstanceId" name:"InstanceId"`
		TemplateId          *string `json:"TemplateId" name:"TemplateId"`
		TemplateType        *string `json:"TemplateType" name:"TemplateType"`
		Status              *string `json:"Status" name:"Status"`
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		CustomConfiguration struct {
			ImageUrl *string `json:"ImageUrl" name:"ImageUrl"`
			Port     *int    `json:"Port" name:"Port"`
			Command  *string `json:"Command" name:"Command"`
			Envs     []struct {
				Key   *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
			} `json:"Envs"`
		} `json:"CustomConfiguration" name:"CustomConfiguration"`
		Timeout *int    `json:"Timeout" name:"Timeout"`
		EndTime *string `json:"EndTime" name:"EndTime"`
		Envs    []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Envs" name:"Envs"`
		Ks3MountConfig struct {
			Ks3Enable      *bool `json:"Ks3Enable" name:"Ks3Enable"`
			Ks3MountPoints []struct {
				BucketName     *string `json:"BucketName" name:"BucketName"`
				RemotePath     *string `json:"RemotePath" name:"RemotePath"`
				LocalMountPath *string `json:"LocalMountPath" name:"LocalMountPath"`
				ReadOnly       *bool   `json:"ReadOnly" name:"ReadOnly"`
			} `json:"Ks3MountPoints"`
		} `json:"Ks3MountConfig" name:"Ks3MountConfig"`
		KpfsMountConfig struct {
			KpfsEnable      *bool `json:"KpfsEnable" name:"KpfsEnable"`
			KpfsMountPoints []struct {
				FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
				RemotePath     *string `json:"RemotePath" name:"RemotePath"`
				LocalMountPath *string `json:"LocalMountPath" name:"LocalMountPath"`
				ReadOnly       *bool   `json:"ReadOnly" name:"ReadOnly"`
			} `json:"KpfsMountPoints"`
		} `json:"KpfsMountConfig" name:"KpfsMountConfig"`
	} `json:"Data"`
}

func (r *GetSandboxInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSandboxInstanceListRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId   *string `json:"TemplateId,omitempty" name:"TemplateId"`
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetSandboxInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetSandboxInstanceListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceSet []struct {
			InstanceId *string `json:"InstanceId" name:"InstanceId"`
			TemplateId *string `json:"TemplateId" name:"TemplateId"`
			AccessUrl  struct {
				CdpUrl   *string `json:"CdpUrl" name:"CdpUrl"`
				NoVncUrl *string `json:"NoVncUrl" name:"NoVncUrl"`
				CodeUrl  *string `json:"CodeUrl" name:"CodeUrl"`
				AppUrl   *string `json:"AppUrl" name:"AppUrl"`
			} `json:"AccessUrl"`
			TemplateType   *string `json:"TemplateType" name:"TemplateType"`
			Status         *string `json:"Status" name:"Status"`
			EndTime        *string `json:"EndTime" name:"EndTime"`
			CreateTime     *string `json:"CreateTime" name:"CreateTime"`
			ContainerId    *string `json:"ContainerId" name:"ContainerId"`
			Ks3MountConfig struct {
				Ks3Enable      *bool `json:"Ks3Enable" name:"Ks3Enable"`
				Ks3MountPoints []struct {
					BucketName     *string `json:"BucketName" name:"BucketName"`
					RemotePath     *string `json:"RemotePath" name:"RemotePath"`
					LocalMountPath *string `json:"LocalMountPath" name:"LocalMountPath"`
					ReadOnly       *bool   `json:"ReadOnly" name:"ReadOnly"`
				} `json:"Ks3MountPoints" name:"Ks3MountPoints"`
			} `json:"Ks3MountConfig"`
			KpfsMountConfig struct {
				KpfsEnable      *bool `json:"KpfsEnable" name:"KpfsEnable"`
				KpfsMountPoints []struct {
					FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
					RemotePath     *string `json:"RemotePath" name:"RemotePath"`
					LocalMountPath *string `json:"LocalMountPath" name:"LocalMountPath"`
					ReadOnly       *bool   `json:"ReadOnly" name:"ReadOnly"`
				} `json:"KpfsMountPoints" name:"KpfsMountPoints"`
			} `json:"KpfsMountConfig"`
		} `json:"InstanceSet" name:"InstanceSet"`
		TotalCount *int    `json:"TotalCount" name:"TotalCount"`
		NextToken  *string `json:"NextToken" name:"NextToken"`
		MaxResults *int    `json:"MaxResults" name:"MaxResults"`
		PageNum    *int    `json:"PageNum" name:"PageNum"`
		PageSize   *int    `json:"PageSize" name:"PageSize"`
	} `json:"Data"`
}

func (r *GetSandboxInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSandboxInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSandboxTemplateListRequest struct {
	*ksyunhttp.BaseRequest
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	PageNum      *int    `json:"PageNum,omitempty" name:"PageNum"`
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetSandboxTemplateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetSandboxTemplateListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Templates []struct {
			TemplateId       *string `json:"TemplateId" name:"TemplateId"`
			TemplateName     *string `json:"TemplateName" name:"TemplateName"`
			Description      *string `json:"Description" name:"Description"`
			TemplateType     *string `json:"TemplateType" name:"TemplateType"`
			Status           *string `json:"Status" name:"Status"`
			CreatedAt        *string `json:"CreatedAt" name:"CreatedAt"`
			CanDelete        *bool   `json:"CanDelete" name:"CanDelete"`
			TemplateCategory *string `json:"TemplateCategory" name:"TemplateCategory"`
			KlogConfig       struct {
				KlogEnable      *bool   `json:"KlogEnable" name:"KlogEnable"`
				KlogProjectName *string `json:"KlogProjectName" name:"KlogProjectName"`
				KlogPoolName    *string `json:"KlogPoolName" name:"KlogPoolName"`
			} `json:"KlogConfig"`
		} `json:"Templates" name:"Templates"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		PageNum    *int `json:"PageNum" name:"PageNum"`
		PageSize   *int `json:"PageSize" name:"PageSize"`
	} `json:"Data"`
}

func (r *GetSandboxTemplateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSandboxTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSandboxInstanceRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId      *string                              `json:"TemplateId,omitempty" name:"TemplateId"`
	Timeout         *int                                 `json:"Timeout,omitempty" name:"Timeout"`
	Ks3MountConfig  *StartSandboxInstanceKs3MountConfig  `json:"Ks3MountConfig,omitempty" name:"Ks3MountConfig"`
	KpfsMountConfig *StartSandboxInstanceKpfsMountConfig `json:"KpfsMountConfig,omitempty" name:"KpfsMountConfig"`
	AccessKey       *string                              `json:"AccessKey,omitempty" name:"AccessKey"`
	SecretAccessKey *string                              `json:"SecretAccessKey,omitempty" name:"SecretAccessKey"`
	Envs            []*StartSandboxInstanceEnvs          `json:"Envs,omitempty" name:"Envs"`
}

func (r *StartSandboxInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartSandboxInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Endpoint   *string `json:"Endpoint" name:"Endpoint"`
		TemplateId *string `json:"TemplateId" name:"TemplateId"`
		Token      *string `json:"Token" name:"Token"`
		Timeout    *int    `json:"Timeout" name:"Timeout"`
	} `json:"Data"`
}

func (r *StartSandboxInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSandboxTemplateRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteSandboxTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSandboxTemplateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TemplateId *string `json:"TemplateId" name:"TemplateId"`
	} `json:"Data"`
}

func (r *DeleteSandboxTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSandboxTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSandboxTemplateRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *GetSandboxTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetSandboxTemplateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		TemplateId   *string `json:"TemplateId" name:"TemplateId"`
		TemplateName *string `json:"TemplateName" name:"TemplateName"`
		Command      *string `json:"Command" name:"Command"`
		CreatedAt    *string `json:"CreatedAt" name:"CreatedAt"`
		Description  *string `json:"Description" name:"Description"`
		Envs         []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"Envs" name:"Envs"`
		Status           *string `json:"Status" name:"Status"`
		TemplateType     *string `json:"TemplateType" name:"TemplateType"`
		CanDelete        *bool   `json:"CanDelete" name:"CanDelete"`
		TemplateCategory *string `json:"TemplateCategory" name:"TemplateCategory"`
		Ks3MountConfig   struct {
			Ks3Enable      *bool `json:"Ks3Enable" name:"Ks3Enable"`
			Ks3MountPoints []struct {
				BucketName     *string `json:"BucketName" name:"BucketName"`
				RemotePath     *string `json:"RemotePath" name:"RemotePath"`
				LocalMountPath *string `json:"LocalMountPath" name:"LocalMountPath"`
				ReadOnly       *bool   `json:"ReadOnly" name:"ReadOnly"`
			} `json:"Ks3MountPoints"`
		} `json:"Ks3MountConfig" name:"Ks3MountConfig"`
		KpfsMountConfig struct {
			KpfsEnable      *bool `json:"KpfsEnable" name:"KpfsEnable"`
			KpfsMountPoints []struct {
				FileSystemName *string `json:"FileSystemName" name:"FileSystemName"`
				RemotePath     *string `json:"RemotePath" name:"RemotePath"`
				LocalMountPath *string `json:"LocalMountPath" name:"LocalMountPath"`
				ReadOnly       *bool   `json:"ReadOnly" name:"ReadOnly"`
			} `json:"KpfsMountPoints"`
		} `json:"KpfsMountConfig" name:"KpfsMountConfig"`
		Ports         []*int `json:"Ports" name:"Ports"`
		PreheatConfig struct {
			PreheatEnable           *bool `json:"PreheatEnable" name:"PreheatEnable"`
			PreheatNumber           *int  `json:"PreheatNumber" name:"PreheatNumber"`
			PreheatedInstanceNumber *int  `json:"PreheatedInstanceNumber" name:"PreheatedInstanceNumber"`
		} `json:"PreheatConfig" name:"PreheatConfig"`
		UpdatedAt  *string `json:"UpdatedAt" name:"UpdatedAt"`
		KlogConfig struct {
			KlogEnable      *bool   `json:"KlogEnable" name:"KlogEnable"`
			KlogProjectName *string `json:"KlogProjectName" name:"KlogProjectName"`
			KlogPoolName    *string `json:"KlogPoolName" name:"KlogPoolName"`
		} `json:"KlogConfig" name:"KlogConfig"`
		NetworkConfig struct {
			PublicNetworkEnable        *bool `json:"PublicNetworkEnable" name:"PublicNetworkEnable"`
			PrivateNetworkEnable       *bool `json:"PrivateNetworkEnable" name:"PrivateNetworkEnable"`
			SharedInternetAccessEnable *bool `json:"SharedInternetAccessEnable" name:"SharedInternetAccessEnable"`
			VpcConfiguration           struct {
				VpcId    *string `json:"VpcId" name:"VpcId"`
				SubnetId *string `json:"SubnetId" name:"SubnetId"`
			} `json:"VpcConfiguration"`
		} `json:"NetworkConfig" name:"NetworkConfig"`
		ImageConfig struct {
			ImageUrl    *string `json:"ImageUrl" name:"ImageUrl"`
			ImageTag    *string `json:"ImageTag" name:"ImageTag"`
			ImageSource *string `json:"ImageSource" name:"ImageSource"`
		} `json:"ImageConfig" name:"ImageConfig"`
		SkillConfig struct {
			SkillEnable       *bool     `json:"SkillEnable" name:"SkillEnable"`
			SkillSpaceIds     []*string `json:"SkillSpaceIds" name:"SkillSpaceIds"`
			PublicSkillEnable *bool     `json:"PublicSkillEnable" name:"PublicSkillEnable"`
		} `json:"SkillConfig" name:"SkillConfig"`
		AccessKey                    *string `json:"AccessKey" name:"AccessKey"`
		InstanceQuota                *int    `json:"InstanceQuota" name:"InstanceQuota"`
		RemainingInstanceQuota       *int    `json:"RemainingInstanceQuota" name:"RemainingInstanceQuota"`
		RemainingSystemInstanceQuota *int    `json:"RemainingSystemInstanceQuota" name:"RemainingSystemInstanceQuota"`
		KecConfig                    struct {
			KecEnable     *bool `json:"KecEnable" name:"KecEnable"`
			InstanceSpecs []struct {
				InstanceType *string `json:"InstanceType" name:"InstanceType"`
				Cpu          *int    `json:"Cpu" name:"Cpu"`
				Memory       *int    `json:"Memory" name:"Memory"`
				SystemDisk   struct {
					Type *string `json:"Type" name:"Type"`
					Size *int    `json:"Size" name:"Size"`
				} `json:"SystemDisk" name:"SystemDisk"`
				DataDisks []struct {
					Type               *string `json:"Type" name:"Type"`
					Size               *int    `json:"Size" name:"Size"`
					DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
					Path               *string `json:"Path" name:"Path"`
				} `json:"DataDisks" name:"DataDisks"`
			} `json:"InstanceSpecs"`
		} `json:"KecConfig" name:"KecConfig"`
	} `json:"Data"`
}

func (r *GetSandboxTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSandboxTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPublicImageListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetPublicImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetPublicImageListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ImageSet []struct {
			ImageUrl     *string  `json:"ImageUrl" name:"ImageUrl"`
			ImageType    *string  `json:"ImageType" name:"ImageType"`
			Description  *string  `json:"Description" name:"Description"`
			ImageSize    *float64 `json:"ImageSize" name:"ImageSize"`
			ImageVersion *string  `json:"ImageVersion" name:"ImageVersion"`
		} `json:"ImageSet" name:"ImageSet"`
	} `json:"Data"`
}

func (r *GetPublicImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPublicImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSandboxInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Timeout    *int    `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *UpdateSandboxInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateSandboxInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Timeout    *int    `json:"Timeout" name:"Timeout"`
	} `json:"Data"`
}

func (r *UpdateSandboxInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
