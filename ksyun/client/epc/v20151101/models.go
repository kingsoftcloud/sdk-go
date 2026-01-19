package v20151101

import (
	"encoding/json"

	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeEpcsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeImagesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeInspectionsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeEpcStocksFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeEpcDeviceAttributesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeProcessesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeEpcRaidAttributesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeUseHotStandbyRecordsFilterN struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeRoceEventFilterN struct {
	Name   *string `json:"Name,omitempty" name:"Name"`
	ValueM *string `json:"Value.M,omitempty" name:"Value.M"`
}
type RunSoInstancesVolumes struct {
	Size []*int `json:"Size,omitempty" name:"Size"`
}

type CreateEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostType                        *string   `json:"HostType,omitempty" name:"HostType"`
	AvailabilityZone                *string   `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Raid                            *string   `json:"Raid,omitempty" name:"Raid"`
	RaidId                          *string   `json:"RaidId,omitempty" name:"RaidId"`
	ImageId                         *string   `json:"ImageId,omitempty" name:"ImageId"`
	NetworkInterfaceMode            *string   `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
	SubnetId                        *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	PrivateIpAddress                *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	KeyId                           *string   `json:"keyId,omitempty" name:"keyId"`
	SecurityGroupId                 []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	DNS1                            *string   `json:"DNS1,omitempty" name:"DNS1"`
	DNS2                            *string   `json:"DNS2,omitempty" name:"DNS2"`
	HostName                        *string   `json:"HostName,omitempty" name:"HostName"`
	ProjectId                       *string   `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType                      *string   `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime                    *int      `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	Password                        *string   `json:"Password,omitempty" name:"Password"`
	CloudMonitorAgent               *string   `json:"CloudMonitorAgent,omitempty" name:"CloudMonitorAgent"`
	ExtensionSubnetId               *string   `json:"ExtensionSubnetId,omitempty" name:"ExtensionSubnetId"`
	ExtensionPrivateIpAddress       *string   `json:"ExtensionPrivateIpAddress,omitempty" name:"ExtensionPrivateIpAddress"`
	ExtensionDNS1                   *string   `json:"ExtensionDNS1,omitempty" name:"ExtensionDNS1"`
	ExtensionDNS2                   *string   `json:"ExtensionDNS2,omitempty" name:"ExtensionDNS2"`
	Description                     *string   `json:"Description,omitempty" name:"Description"`
	AddressBandWidth                *string   `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
	LineId                          *string   `json:"LineId,omitempty" name:"LineId"`
	BandWidthShareId                *string   `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	AddressChargeType               *string   `json:"AddressChargeType,omitempty" name:"AddressChargeType"`
	AddressPurchaseTime             *string   `json:"AddressPurchaseTime,omitempty" name:"AddressPurchaseTime"`
	AddressProjectId                *string   `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
	SystemFileType                  *string   `json:"SystemFileType,omitempty" name:"SystemFileType"`
	DataFileType                    *string   `json:"DataFileType,omitempty" name:"DataFileType"`
	DataDiskCatalogue               *string   `json:"DataDiskCatalogue,omitempty" name:"DataDiskCatalogue"`
	DataDiskCatalogueSuffix         *string   `json:"DataDiskCatalogueSuffix,omitempty" name:"DataDiskCatalogueSuffix"`
	HyperThreading                  *string   `json:"HyperThreading,omitempty" name:"HyperThreading"`
	NvmeDataFileType                *string   `json:"NvmeDataFileType,omitempty" name:"NvmeDataFileType"`
	NvmeDataDiskCatalogue           *string   `json:"NvmeDataDiskCatalogue,omitempty" name:"NvmeDataDiskCatalogue"`
	NvmeDataDiskCatalogueSuffix     *string   `json:"NvmeDataDiskCatalogueSuffix,omitempty" name:"NvmeDataDiskCatalogueSuffix"`
	BondAttribute                   *string   `json:"BondAttribute,omitempty" name:"BondAttribute"`
	ContainerAgent                  *string   `json:"ContainerAgent,omitempty" name:"ContainerAgent"`
	KesAgent                        *string   `json:"KesAgent,omitempty" name:"KesAgent"`
	KmrAgent                        *string   `json:"KmrAgent,omitempty" name:"KmrAgent"`
	ComputerName                    *string   `json:"ComputerName,omitempty" name:"ComputerName"`
	OverclockingAttribute           *string   `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
	GpuImageDriverId                *string   `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
	SystemVolumeType                *string   `json:"SystemVolumeType,omitempty" name:"SystemVolumeType"`
	SystemVolumeSize                *string   `json:"SystemVolumeSize,omitempty" name:"SystemVolumeSize"`
	RoceNetwork                     *string   `json:"RoceNetwork,omitempty" name:"RoceNetwork"`
	ZoneId                          *string   `json:"ZoneId,omitempty" name:"ZoneId"`
	ZoneType                        *string   `json:"ZoneType,omitempty" name:"ZoneType"`
	UseHotStandby                   *string   `json:"UseHotStandby,omitempty" name:"UseHotStandby"`
	TimedRegularization             *string   `json:"TimedRegularization,omitempty" name:"TimedRegularization"`
	PasswordInherit                 *string   `json:"PasswordInherit,omitempty" name:"PasswordInherit"`
	DataDiskMount                   *string   `json:"DataDiskMount,omitempty" name:"DataDiskMount"`
	StorageRoceNetworkCardName      *string   `json:"StorageRoceNetworkCardName,omitempty" name:"StorageRoceNetworkCardName"`
	AnacondaN                       *string   `json:"Anaconda.N,omitempty" name:"Anaconda.N"`
	FrameworkN                      *string   `json:"Framework.N,omitempty" name:"Framework.N"`
	EngineN                         *string   `json:"Engine.N,omitempty" name:"Engine.N"`
	AiModelN                        *string   `json:"AiModel.N,omitempty" name:"AiModel.N"`
	UserData                        *string   `json:"UserData,omitempty" name:"UserData"`
	StorageRoceNetworkInterfaceMode *string   `json:"StorageRoceNetworkInterfaceMode,omitempty" name:"StorageRoceNetworkInterfaceMode"`
	RoceCluster                     *string   `json:"RoceCluster,omitempty" name:"RoceCluster"`
	SRoceCluster                    *string   `json:"SRoceCluster,omitempty" name:"SRoceCluster"`
	UserDefinedData                 *string   `json:"UserDefinedData,omitempty" name:"UserDefinedData"`
}

func (r *CreateEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	AutoReserveInfo struct {
		Reserve *bool   `json:"Reserve" name:"Reserve"`
		HostId  *string `json:"HostId" name:"HostId"`
	} `json:"AutoReserveInfo"`
	Host struct {
		TorName              *string `json:"TorName" name:"TorName"`
		RackId               *string `json:"RackId" name:"RackId"`
		ContainerAgent       *string `json:"ContainerAgent" name:"ContainerAgent"`
		KesAgent             *string `json:"KesAgent" name:"KesAgent"`
		BondAttribute        *string `json:"BondAttribute" name:"BondAttribute"`
		GpuImageDriverId     *string `json:"GpuImageDriverId" name:"GpuImageDriverId"`
		NvmeDataFileType     *string `json:"NvmeDataFileType" name:"NvmeDataFileType"`
		Memory               *string `json:"Memory" name:"Memory"`
		Raid                 *string `json:"Raid" name:"Raid"`
		CloudMonitorAgent    *string `json:"CloudMonitorAgent" name:"CloudMonitorAgent"`
		DataDiskCatalogue    *string `json:"DataDiskCatalogue" name:"DataDiskCatalogue"`
		NetworkInterfaceMode *string `json:"NetworkInterfaceMode" name:"NetworkInterfaceMode"`
		RaidTemplateId       *string `json:"RaidTemplateId" name:"RaidTemplateId"`
		DataVolumeSet        []struct {
			VolumeId           *string `json:"VolumeId" name:"VolumeId"`
			VolumeType         *string `json:"VolumeType" name:"VolumeType"`
			DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
			VolumeSize         *string `json:"VolumeSize" name:"VolumeSize"`
		} `json:"DataVolumeSet" name:"DataVolumeSet"`
		ImageId                      *string `json:"ImageId" name:"ImageId"`
		SystemVolumeType             *string `json:"SystemVolumeType" name:"SystemVolumeType"`
		HostName                     *string `json:"HostName" name:"HostName"`
		Tags                         *string `json:"Tags" name:"Tags"`
		SystemFileType               *string `json:"SystemFileType" name:"SystemFileType"`
		EnableBond                   *bool   `json:"EnableBond" name:"EnableBond"`
		ProductType                  *string `json:"ProductType" name:"ProductType"`
		AvailabilityZone             *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		HostId                       *string `json:"HostId" name:"HostId"`
		NetworkInterfaceAttributeSet []struct {
			PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
			SecurityGroupSet     []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			} `json:"SecurityGroupSet"`
			DNS1               *string `json:"DNS1" name:"DNS1"`
			DNS2               *string `json:"DNS2" name:"DNS2"`
			SubnetId           *string `json:"SubnetId" name:"SubnetId"`
			NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			Mac                *string `json:"Mac" name:"Mac"`
			VpcId              *string `json:"VpcId" name:"VpcId"`
			Ipv6Address        *string `json:"Ipv6Address" name:"Ipv6Address"`
		} `json:"NetworkInterfaceAttributeSet" name:"NetworkInterfaceAttributeSet"`
		ComputerName *string `json:"ComputerName" name:"ComputerName"`
		CabinetId    *string `json:"CabinetId" name:"CabinetId"`
		DiskSet      []struct {
			DiskType        *string `json:"DiskType" name:"DiskType"`
			Space           *string `json:"Space" name:"Space"`
			DiskCount       *int    `json:"DiskCount" name:"DiskCount"`
			Raid            *string `json:"Raid" name:"Raid"`
			DiskAttribute   *string `json:"DiskAttribute" name:"DiskAttribute"`
			SystemDiskSpace *string `json:"SystemDiskSpace" name:"SystemDiskSpace"`
			DiskSpace       *string `json:"DiskSpace" name:"DiskSpace"`
		} `json:"DiskSet" name:"DiskSet"`
		DataDiskCatalogueSuffix     *string `json:"DataDiskCatalogueSuffix" name:"DataDiskCatalogueSuffix"`
		DataFileType                *string `json:"DataFileType" name:"DataFileType"`
		HostType                    *string `json:"HostType" name:"HostType"`
		SystemVolumeSize            *string `json:"SystemVolumeSize" name:"SystemVolumeSize"`
		NvmeDataDiskCatalogue       *string `json:"NvmeDataDiskCatalogue" name:"NvmeDataDiskCatalogue"`
		HostStatus                  *string `json:"HostStatus" name:"HostStatus"`
		EnableContainer             *bool   `json:"EnableContainer" name:"EnableContainer"`
		ClusterId                   *string `json:"ClusterId" name:"ClusterId"`
		HyperThreading              *string `json:"HyperThreading" name:"HyperThreading"`
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		OsName                      *string `json:"OsName" name:"OsName"`
		CabinetName                 *string `json:"CabinetName" name:"CabinetName"`
		ProjectId                   *string `json:"ProjectId" name:"ProjectId"`
		KeyId                       *string `json:"KeyId" name:"KeyId"`
		AllowModifyHyperThreading   *bool   `json:"AllowModifyHyperThreading" name:"AllowModifyHyperThreading"`
		ReleasableTime              *string `json:"ReleasableTime" name:"ReleasableTime"`
		RackName                    *string `json:"RackName" name:"RackName"`
		KmrAgent                    *string `json:"KmrAgent" name:"KmrAgent"`
		Sn                          *string `json:"Sn" name:"Sn"`
		NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix" name:"NvmeDataDiskCatalogueSuffix"`
		SecurityAgent               *string `json:"SecurityAgent" name:"SecurityAgent"`
		SupportEbs                  *string `json:"SupportEbs" name:"SupportEbs"`
		KplAgent                    *string `json:"KplAgent" name:"KplAgent"`
		ServiceEndTime              *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		ChargeType                  *string `json:"ChargeType" name:"ChargeType"`
		Cpu                         struct {
			Model     *string `json:"Model" name:"Model"`
			Frequence *string `json:"Frequence" name:"Frequence"`
			Count     *int    `json:"Count" name:"Count"`
			CoreCount *int    `json:"CoreCount" name:"CoreCount"`
		} `json:"Cpu" name:"Cpu"`
		Gpu struct {
			Model     *string `json:"Model" name:"Model"`
			Frequence *string `json:"Frequence" name:"Frequence"`
			Count     *int    `json:"Count" name:"Count"`
			CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			GpuCount  *int    `json:"GpuCount" name:"GpuCount"`
		} `json:"Gpu" name:"Gpu"`
		Roces []struct {
			Ip      *string `json:"Ip" name:"Ip"`
			Mask    *string `json:"Mask" name:"Mask"`
			GateWay *string `json:"GateWay" name:"GateWay"`
			Type    *string `json:"Type" name:"Type"`
		} `json:"Roces" name:"Roces"`
		ContractDueTime         *string   `json:"ContractDueTime" name:"ContractDueTime"`
		AutoDeleteTime          *string   `json:"AutoDeleteTime" name:"AutoDeleteTime"`
		VpcTrust                *int      `json:"VpcTrust" name:"VpcTrust"`
		TimedRegularizationTime *string   `json:"TimedRegularizationTime" name:"TimedRegularizationTime"`
		Anacondas               []*string `json:"Anacondas" name:"Anacondas"`
		Frameworks              []*string `json:"Frameworks" name:"Frameworks"`
		Engines                 []*string `json:"Engines" name:"Engines"`
		AiModels                []*string `json:"AiModels" name:"AiModels"`
	} `json:"Host"`
	SRoceCluster *string `json:"SRoceCluster" name:"SRoceCluster"`
	RoceCluster  *string `json:"RoceCluster" name:"RoceCluster"`
}

func (r *CreateEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *StartEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *StartEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *RebootEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RebootEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RebootEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *DeleteEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReinstallEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId                      *string `json:"HostId,omitempty" name:"HostId"`
	ImageId                     *string `json:"ImageId,omitempty" name:"ImageId"`
	KeyId                       *string `json:"keyId,omitempty" name:"keyId"`
	Password                    *string `json:"Password,omitempty" name:"Password"`
	NetworkInterfaceMode        *string `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
	CloudMonitorAgent           *string `json:"CloudMonitorAgent,omitempty" name:"CloudMonitorAgent"`
	Raid                        *string `json:"Raid,omitempty" name:"Raid"`
	RaidId                      *string `json:"RaidId,omitempty" name:"RaidId"`
	HostName                    *string `json:"HostName,omitempty" name:"HostName"`
	SystemFileType              *string `json:"SystemFileType,omitempty" name:"SystemFileType"`
	DataFileType                *string `json:"DataFileType,omitempty" name:"DataFileType"`
	DataDiskCatalogue           *string `json:"DataDiskCatalogue,omitempty" name:"DataDiskCatalogue"`
	DataDiskCatalogueSuffix     *string `json:"DataDiskCatalogueSuffix,omitempty" name:"DataDiskCatalogueSuffix"`
	HyperThreading              *string `json:"HyperThreading,omitempty" name:"HyperThreading"`
	NvmeDataFileType            *string `json:"NvmeDataFileType,omitempty" name:"NvmeDataFileType"`
	NvmeDataDiskCatalogue       *string `json:"NvmeDataDiskCatalogue,omitempty" name:"NvmeDataDiskCatalogue"`
	NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix,omitempty" name:"NvmeDataDiskCatalogueSuffix"`
	BondAttribute               *string `json:"BondAttribute,omitempty" name:"BondAttribute"`
	KesAgent                    *string `json:"KesAgent,omitempty" name:"KesAgent"`
	KmrAgent                    *string `json:"KmrAgent,omitempty" name:"KmrAgent"`
	ComputerName                *string `json:"ComputerName,omitempty" name:"ComputerName"`
	OverclockingAttribute       *string `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
	DelayStart                  *int    `json:"DelayStart,omitempty" name:"DelayStart"`
	AvailabilityZone            *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	GpuImageDriverId            *string `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
	ContainerAgent              *string `json:"ContainerAgent,omitempty" name:"ContainerAgent"`
	PasswordInherit             *string `json:"PasswordInherit,omitempty" name:"PasswordInherit"`
	DataDiskMount               *string `json:"DataDiskMount,omitempty" name:"DataDiskMount"`
	StorageRoceNetworkCardName  *string `json:"StorageRoceNetworkCardName,omitempty" name:"StorageRoceNetworkCardName"`
	UserDefinedData             *string `json:"UserDefinedData,omitempty" name:"UserDefinedData"`
}

func (r *ReinstallEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReinstallEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ReinstallEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReinstallEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	HostId             *string   `json:"HostId,omitempty" name:"HostId"`
	NetworkInterfaceId *string   `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	SecurityGroupId    []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *ModifySecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifySecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyName     *string `json:"KeyName,omitempty" name:"KeyName"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	PrivateKey *string `json:"PrivateKey" name:"PrivateKey"`
	Key        struct {
		KeyName    *string `json:"KeyName" name:"KeyName"`
		PublicKey  *string `json:"PublicKey" name:"PublicKey"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		KeyId      *string `json:"KeyId" name:"KeyId"`
	} `json:"Key"`
}

func (r *CreateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcsRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId  []*string             `json:"ProjectId,omitempty" name:"ProjectId"`
	HostId     []*string             `json:"HostId,omitempty" name:"HostId"`
	Filter     []*DescribeEpcsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                  `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string               `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcsResponse struct {
	*ksyunhttp.BaseResponse
	HostSet []struct {
		ComputerName                *string `json:"ComputerName" name:"ComputerName"`
		KmrAgent                    *string `json:"KmrAgent" name:"KmrAgent"`
		RackId                      *string `json:"RackId" name:"RackId"`
		GroupName                   *string `json:"GroupName" name:"GroupName"`
		ContainerAgent              *string `json:"ContainerAgent" name:"ContainerAgent"`
		KesAgent                    *string `json:"KesAgent" name:"KesAgent"`
		BondAttribute               *string `json:"BondAttribute" name:"BondAttribute"`
		NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix" name:"NvmeDataDiskCatalogueSuffix"`
		NvmeDataDiskCatalogue       *string `json:"NvmeDataDiskCatalogue" name:"NvmeDataDiskCatalogue"`
		NvmeDataFileType            *string `json:"NvmeDataFileType" name:"NvmeDataFileType"`
		HyperThreading              *string `json:"HyperThreading" name:"HyperThreading"`
		DataDiskCatalogueSuffix     *string `json:"DataDiskCatalogueSuffix" name:"DataDiskCatalogueSuffix"`
		SystemFileType              *string `json:"SystemFileType" name:"SystemFileType"`
		DataFileType                *string `json:"DataFileType" name:"DataFileType"`
		DataDiskCatalogue           *string `json:"DataDiskCatalogue" name:"DataDiskCatalogue"`
		ReleasableTime              *string `json:"ReleasableTime" name:"ReleasableTime"`
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		HostName                    *string `json:"HostName" name:"HostName"`
		HostType                    *string `json:"HostType" name:"HostType"`
		GroupHostType               *string `json:"GroupHostType" name:"GroupHostType"`
		HostId                      *string `json:"HostId" name:"HostId"`
		Sn                          *string `json:"Sn" name:"Sn"`
		Roces                       []struct {
			Ip              *string `json:"Ip" name:"Ip"`
			Mask            *string `json:"Mask" name:"Mask"`
			GateWay         *string `json:"GateWay" name:"GateWay"`
			Type            *string `json:"Type" name:"Type"`
			SwName          *string `json:"SwName" name:"SwName"`
			SwPort          *string `json:"SwPort" name:"SwPort"`
			NetworkCardName *string `json:"NetworkCardName" name:"NetworkCardName"`
		} `json:"Roces" name:"Roces"`
		ContractDueTime           *string `json:"ContractDueTime" name:"ContractDueTime"`
		DependencyUserId          *string `json:"DependencyUserId" name:"DependencyUserId"`
		AutoDeleteTime            *string `json:"AutoDeleteTime" name:"AutoDeleteTime"`
		TimedRegularizationTime   *string `json:"TimedRegularizationTime" name:"TimedRegularizationTime"`
		CabinetId                 *string `json:"CabinetId" name:"CabinetId"`
		AvailabilityZone          *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Raid                      *string `json:"Raid" name:"Raid"`
		RaidTemplateId            *string `json:"RaidTemplateId" name:"RaidTemplateId"`
		ImageId                   *string `json:"ImageId" name:"ImageId"`
		KeyId                     *string `json:"KeyId" name:"KeyId"`
		NetworkInterfaceMode      *string `json:"NetworkInterfaceMode" name:"NetworkInterfaceMode"`
		PublicNetworkLimit        *bool   `json:"PublicNetworkLimit" name:"PublicNetworkLimit"`
		EnableBond                *bool   `json:"EnableBond" name:"EnableBond"`
		AllowModifyHyperThreading *bool   `json:"AllowModifyHyperThreading" name:"AllowModifyHyperThreading"`
		SecurityAgent             *string `json:"SecurityAgent" name:"SecurityAgent"`
		CloudMonitorAgent         *string `json:"CloudMonitorAgent" name:"CloudMonitorAgent"`
		ProductType               *string `json:"ProductType" name:"ProductType"`
		OsName                    *string `json:"OsName" name:"OsName"`
		Memory                    *string `json:"Memory" name:"Memory"`
		HostStatus                *string `json:"HostStatus" name:"HostStatus"`
		Cpu                       struct {
			Model     *string `json:"Model" name:"Model"`
			Frequence *string `json:"Frequence" name:"Frequence"`
			Count     *int    `json:"Count" name:"Count"`
			CoreCount *int    `json:"CoreCount" name:"CoreCount"`
		} `json:"Cpu" name:"Cpu"`
		Gpu struct {
			Model     *string `json:"Model" name:"Model"`
			Frequence *string `json:"Frequence" name:"Frequence"`
			Count     *int    `json:"Count" name:"Count"`
			CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			GpuCount  *int    `json:"GpuCount" name:"GpuCount"`
		} `json:"Gpu" name:"Gpu"`
		DiskSet []struct {
			DiskType      *string `json:"DiskType" name:"DiskType"`
			Raid          *string `json:"Raid" name:"Raid"`
			DiskSpace     *string `json:"DiskSpace" name:"DiskSpace"`
			Space         *string `json:"Space" name:"Space"`
			DiskAttribute *string `json:"DiskAttribute" name:"DiskAttribute"`
			DiskCount     *int    `json:"DiskCount" name:"DiskCount"`
		} `json:"DiskSet" name:"DiskSet"`
		NetworkInterfaceAttributeSet []struct {
			VpcId                *string `json:"VpcId" name:"VpcId"`
			NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
			SubnetId             *string `json:"SubnetId" name:"SubnetId"`
			PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			DNS1                 *string `json:"DNS1" name:"DNS1"`
			DNS2                 *string `json:"DNS2" name:"DNS2"`
			Mac                  *string `json:"Mac" name:"Mac"`
			SecurityGroupSet     []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			} `json:"SecurityGroupSet"`
			EipAddress struct {
				BandWidth      *int    `json:"BandWidth" name:"BandWidth"`
				CreateTime     *string `json:"CreateTime" name:"CreateTime"`
				AllocationId   *string `json:"AllocationId" name:"AllocationId"`
				PublicIp       *string `json:"PublicIp" name:"PublicIp"`
				IpState        *string `json:"IpState" name:"IpState"`
				ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
				IpVersion      *string `json:"IpVersion" name:"IpVersion"`
			} `json:"EipAddress"`
			Ipv6Address *string `json:"Ipv6Address" name:"Ipv6Address"`
		} `json:"NetworkInterfaceAttributeSet" name:"NetworkInterfaceAttributeSet"`
		NetworkCardSet []struct {
			Type      *string `json:"Type" name:"Type"`
			Attribute *string `json:"Attribute" name:"Attribute"`
			Num       *int    `json:"Num" name:"Num"`
			Kind      *string `json:"Kind" name:"Kind"`
		} `json:"NetworkCardSet" name:"NetworkCardSet"`
		ProjectId        *string `json:"ProjectId" name:"ProjectId"`
		CabinetName      *string `json:"CabinetName" name:"CabinetName"`
		RackName         *string `json:"RackName" name:"RackName"`
		TorName          *string `json:"TorName" name:"TorName"`
		ClusterId        *string `json:"ClusterId" name:"ClusterId"`
		EnableContainer  *bool   `json:"EnableContainer" name:"EnableContainer"`
		SupportEbs       *string `json:"SupportEbs" name:"SupportEbs"`
		SystemVolumeType *string `json:"SystemVolumeType" name:"SystemVolumeType"`
		SystemVolumeSize *string `json:"SystemVolumeSize" name:"SystemVolumeSize"`
		DataVolumeSet    []struct {
			VolumeId           *string `json:"VolumeId" name:"VolumeId"`
			VolumeType         *string `json:"VolumeType" name:"VolumeType"`
			VolumeSize         *string `json:"VolumeSize" name:"VolumeSize"`
			DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
		} `json:"DataVolumeSet" name:"DataVolumeSet"`
		GpuImageDriverId *string   `json:"GpuImageDriverId" name:"GpuImageDriverId"`
		VpcTrust         *int      `json:"VpcTrust" name:"VpcTrust"`
		ServiceEndTime   *string   `json:"ServiceEndTime" name:"ServiceEndTime"`
		ChargeType       *string   `json:"ChargeType" name:"ChargeType"`
		Anacondas        []*string `json:"Anacondas" name:"Anacondas"`
		Frameworks       []*string `json:"Frameworks" name:"Frameworks"`
		Engines          []*string `json:"Engines" name:"Engines"`
		AiModels         []*string `json:"AiModels" name:"AiModels"`
		RoceCluster      *string   `json:"RoceCluster" name:"RoceCluster"`
		RoceClusterGroup *string   `json:"RoceClusterGroup" name:"RoceClusterGroup"`
		SRoceCluster     *string   `json:"SRoceCluster" name:"SRoceCluster"`
	} `json:"HostSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	Tags       []struct {
		ResourceId *string `json:"ResourceId" name:"ResourceId"`
		TagId      *string `json:"TagId" name:"TagId"`
		Key        *string `json:"Key" name:"Key"`
		Value      *string `json:"Value" name:"Value"`
	} `json:"Tags"`
}

func (r *DescribeEpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDynamicCodeRequest struct {
	*ksyunhttp.BaseRequest
	RemoteManagementId *string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
}

func (r *GetDynamicCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDynamicCodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *GetDynamicCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDynamicCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnsRequest struct {
	*ksyunhttp.BaseRequest
	DynamicCode        *string `json:"DynamicCode,omitempty" name:"DynamicCode"`
	Pin                *string `json:"Pin,omitempty" name:"Pin"`
	RemoteManagementId *string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
}

func (r *DescribeVpnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeVpnsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	UserName  *string `json:"UserName" name:"UserName"`
	Password  *string `json:"Password" name:"Password"`
}

func (r *DescribeVpnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageRequest struct {
	*ksyunhttp.BaseRequest
	HostId              *string `json:"HostId,omitempty" name:"HostId"`
	ImageName           *string `json:"ImageName,omitempty" name:"ImageName"`
	ImageMode           *string `json:"ImageMode,omitempty" name:"ImageMode"`
	ImageInitialization *string `json:"ImageInitialization,omitempty" name:"ImageInitialization"`
}

func (r *CreateImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Image     struct {
		ImageId             *string `json:"ImageId" name:"ImageId"`
		ImageName           *string `json:"ImageName" name:"ImageName"`
		ImageType           *string `json:"ImageType" name:"ImageType"`
		OsName              *string `json:"OsName" name:"OsName"`
		OsType              *string `json:"OsType" name:"OsType"`
		EnableContainer     *bool   `json:"EnableContainer" name:"EnableContainer"`
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		DiskType            *string `json:"DiskType" name:"DiskType"`
		EbsImageSize        *string `json:"EbsImageSize" name:"EbsImageSize"`
		Status              *string `json:"Status" name:"Status"`
		ImageInitialization *string `json:"ImageInitialization" name:"ImageInitialization"`
		Name                *string `json:"Name" name:"Name"`
	} `json:"Image"`
}

func (r *CreateImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	ImageId   *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *ModifyImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int                    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                 `json:"NextToken,omitempty" name:"NextToken"`
	ImageIdN   *string                 `json:"ImageId.N,omitempty" name:"ImageId.N"`
	Filter     []*DescribeImagesFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeImagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	ImageSet   []struct {
		ImageId             *string `json:"ImageId" name:"ImageId"`
		ImageName           *string `json:"ImageName" name:"ImageName"`
		ImageType           *string `json:"ImageType" name:"ImageType"`
		OsName              *string `json:"OsName" name:"OsName"`
		OsType              *string `json:"OsType" name:"OsType"`
		EnableContainer     *bool   `json:"EnableContainer" name:"EnableContainer"`
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		DiskType            *string `json:"DiskType" name:"DiskType"`
		EbsImageSize        *string `json:"EbsImageSize" name:"EbsImageSize"`
		Status              *string `json:"Status" name:"Status"`
		ImageInitialization *string `json:"ImageInitialization" name:"ImageInitialization"`
		Name                *string `json:"Name" name:"Name"`
		ExportStatus        *string `json:"ExportStatus" name:"ExportStatus"`
		ExportProgress      *string `json:"ExportProgress" name:"ExportProgress"`
		Source              *string `json:"Source" name:"Source"`
		KernelVersion       *string `json:"KernelVersion" name:"KernelVersion"`
	} `json:"ImageSet"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	HostId             *string `json:"HostId,omitempty" name:"HostId"`
	DNS1               *string `json:"DNS1,omitempty" name:"DNS1"`
	DNS2               *string `json:"DNS2,omitempty" name:"DNS2"`
}

func (r *ModifyDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDnsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId  *string   `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	HostId              *string   `json:"HostId,omitempty" name:"HostId"`
	SubnetId            *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	IpAddress           *string   `json:"IpAddress,omitempty" name:"IpAddress"`
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`
	SecurityGroupId     []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DescribePhysicalMonitorRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *DescribePhysicalMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribePhysicalMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	CPUInfo   []struct {
		Item   *string `json:"Item" name:"Item"`
		Status *string `json:"Status" name:"Status"`
	} `json:"CPUInfo"`
	MemoryInfo []struct {
		Item   *string `json:"Item" name:"Item"`
		Status *string `json:"Status" name:"Status"`
	} `json:"MemoryInfo"`
	DiskInfo []struct {
		Item   *string `json:"Item" name:"Item"`
		Status *string `json:"Status" name:"Status"`
	} `json:"DiskInfo"`
	FanInfo []struct {
		Item   *string `json:"Item" name:"Item"`
		Status *string `json:"Status" name:"Status"`
	} `json:"FanInfo"`
	PowerInfo []struct {
		Item   *string `json:"Item" name:"Item"`
		Status *string `json:"Status" name:"Status"`
	} `json:"PowerInfo"`
}

func (r *DescribePhysicalMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysicalMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcManagementsRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults         *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken          *string   `json:"NextToken,omitempty" name:"NextToken"`
	DynamicCode        *string   `json:"DynamicCode,omitempty" name:"DynamicCode"`
	Pin                *string   `json:"Pin,omitempty" name:"Pin"`
	EpcManagementId    []*string `json:"EpcManagementId,omitempty" name:"EpcManagementId"`
	RemoteManagementId *string   `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
	ProjectId          []*string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeEpcManagementsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcManagementsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	NextToken        *string `json:"NextToken" name:"NextToken"`
	TotalCount       *int64  `json:"TotalCount" name:"TotalCount"`
	EpcManagementSet []struct {
		EpcManagementId       *string `json:"EpcManagementId" name:"EpcManagementId"`
		EpcManagementIp       *string `json:"EpcManagementIp" name:"EpcManagementIp"`
		EpcManagementUserName *string `json:"EpcManagementUserName" name:"EpcManagementUserName"`
		Password              *string `json:"Password" name:"Password"`
		HostName              *string `json:"HostName" name:"HostName"`
		Sn                    *string `json:"Sn" name:"Sn"`
		ProjectId             *int64  `json:"ProjectId" name:"ProjectId"`
	} `json:"EpcManagementSet"`
}

func (r *DescribeEpcManagementsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEpcManagementsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteManagementsRequest struct {
	*ksyunhttp.BaseRequest
	RemoteManagementId []*string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
}

func (r *DescribeRemoteManagementsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRemoteManagementsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId           *string `json:"RequestId" name:"RequestId"`
	RemoteManagementSet []struct {
		RemoteManagementId *string `json:"RemoteManagementId" name:"RemoteManagementId"`
		PhoneNumber        *string `json:"PhoneNumber" name:"PhoneNumber"`
		Name               *string `json:"Name" name:"Name"`
	} `json:"RemoteManagementSet"`
}

func (r *DescribeRemoteManagementsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteManagementsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *StopEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *StopEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId      *string `json:"HostId,omitempty" name:"HostId"`
	HostName    *string `json:"HostName,omitempty" name:"HostName"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteManagementRequest struct {
	*ksyunhttp.BaseRequest
	RemoteManagementId *string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
	DynamicCode        *string `json:"DynamicCode,omitempty" name:"DynamicCode"`
	Pin                *string `json:"Pin,omitempty" name:"Pin"`
	NewPhoneNumber     *string `json:"NewPhoneNumber,omitempty" name:"NewPhoneNumber"`
	NewPin             *string `json:"NewPin,omitempty" name:"NewPin"`
	Name               *string `json:"Name,omitempty" name:"Name"`
	VersionId          *int64  `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *ModifyRemoteManagementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyRemoteManagementResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	RemoteManagement struct {
		RemoteManagementId *string `json:"RemoteManagementId" name:"RemoteManagementId"`
		PhoneNumber        *string `json:"PhoneNumber" name:"PhoneNumber"`
		Name               *string `json:"Name" name:"Name"`
	} `json:"RemoteManagement"`
}

func (r *ModifyRemoteManagementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteManagementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteManagementRequest struct {
	*ksyunhttp.BaseRequest
	DynamicCode *string `json:"DynamicCode,omitempty" name:"DynamicCode"`
	Pin         *string `json:"Pin,omitempty" name:"Pin"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	Name        *string `json:"Name,omitempty" name:"Name"`
	VersionId   *int64  `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *CreateRemoteManagementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateRemoteManagementResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	RemoteManagement struct {
		RemoteManagementId *string `json:"RemoteManagementId" name:"RemoteManagementId"`
		PhoneNumber        *string `json:"PhoneNumber" name:"PhoneNumber"`
		Name               *string `json:"Name" name:"Name"`
	} `json:"RemoteManagement"`
}

func (r *CreateRemoteManagementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteManagementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReinstallCustomerEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId   *string `json:"HostId,omitempty" name:"HostId"`
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`
	Path     *string `json:"Path,omitempty" name:"Path"`
}

func (r *ReinstallCustomerEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReinstallCustomerEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ReinstallCustomerEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReinstallCustomerEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteManagementRequest struct {
	*ksyunhttp.BaseRequest
	RemoteManagementId *string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
}

func (r *DeleteRemoteManagementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRemoteManagementResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteRemoteManagementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteManagementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordRequest struct {
	*ksyunhttp.BaseRequest
	HostId   *string `json:"HostId,omitempty" name:"HostId"`
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResetPasswordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHyperThreadingRequest struct {
	*ksyunhttp.BaseRequest
	HostId               *string `json:"HostId,omitempty" name:"HostId"`
	HyperThreadingStatus *string `json:"HyperThreadingStatus,omitempty" name:"HyperThreadingStatus"`
}

func (r *ModifyHyperThreadingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyHyperThreadingResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyHyperThreadingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHyperThreadingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateClusterRequest struct {
	*ksyunhttp.BaseRequest
	HostId    *string `json:"HostId,omitempty" name:"HostId"`
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AssociateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AssociateClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateClusterRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *DisassociateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisassociateClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DisassociateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInspectionsRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int                         `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                      `json:"NextToken,omitempty" name:"NextToken"`
	Filter     []*DescribeInspectionsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeInspectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInspectionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	NextToken     *string `json:"NextToken" name:"NextToken"`
	TotalCount    *int64  `json:"TotalCount" name:"TotalCount"`
	InspectionSet []struct {
		HostId           *string `json:"HostId" name:"HostId"`
		Sn               *string `json:"Sn" name:"Sn"`
		Region           *string `json:"Region" name:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Status           *string `json:"Status" name:"Status"`
		AlarmType        *string `json:"AlarmType" name:"AlarmType"`
		AlarmDescription *string `json:"AlarmDescription" name:"AlarmDescription"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		TorName          *string `json:"TorName" name:"TorName"`
		CabinetName      *string `json:"CabinetName" name:"CabinetName"`
	} `json:"InspectionSet"`
}

func (r *DescribeInspectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInspectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcStocksRequest struct {
	*ksyunhttp.BaseRequest
	Filter []*DescribeEpcStocksFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeEpcStocksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcStocksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	EpcStockSet []struct {
		InstanceCount         *int      `json:"InstanceCount" name:"InstanceCount"`
		HostType              *string   `json:"HostType" name:"HostType"`
		AvailabilityZone      *string   `json:"AvailabilityZone" name:"AvailabilityZone"`
		AvailableRaidLevelSet []*string `json:"AvailableRaidLevelSet" name:"AvailableRaidLevelSet"`
		Attribute             *string   `json:"Attribute" name:"Attribute"`
	} `json:"EpcStockSet"`
}

func (r *DescribeEpcStocksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEpcStocksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcDeviceAttributesRequest struct {
	*ksyunhttp.BaseRequest
	Filter            []*DescribeEpcDeviceAttributesFilter `json:"Filter,omitempty" name:"Filter"`
	DeviceAttributeId []*string                            `json:"DeviceAttributeId,omitempty" name:"DeviceAttributeId"`
	MaxResults        *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken         *string                              `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcDeviceAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcDeviceAttributesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	NextToken             *string `json:"NextToken" name:"NextToken"`
	EpcDeviceAttributeSet []struct {
		DeviceAttributeId *string `json:"DeviceAttributeId" name:"DeviceAttributeId"`
		HostType          *string `json:"HostType" name:"HostType"`
		HostTypeName      *string `json:"HostTypeName" name:"HostTypeName"`
		Memory            *string `json:"Memory" name:"Memory"`
		NetworkCard       *string `json:"NetworkCard" name:"NetworkCard"`
		CpuDeviceSet      []struct {
			CpuSpec *string `json:"CpuSpec" name:"CpuSpec"`
		} `json:"CpuDeviceSet" name:"CpuDeviceSet"`
		GpuDeviceSet []struct {
			GpuModel *string `json:"GpuModel" name:"GpuModel"`
			GpuCount *string `json:"GpuCount" name:"GpuCount"`
		} `json:"GpuDeviceSet" name:"GpuDeviceSet"`
		PhysicalDiskDeviceSet []struct {
			DiskAttribute *string `json:"DiskAttribute" name:"DiskAttribute"`
			DiskCount     *string `json:"DiskCount" name:"DiskCount"`
			Space         *string `json:"Space" name:"Space"`
		} `json:"PhysicalDiskDeviceSet" name:"PhysicalDiskDeviceSet"`
		PriceSet []struct {
			MonthlylistPrice *string `json:"MonthlylistPrice" name:"MonthlylistPrice"`
		} `json:"PriceSet" name:"PriceSet"`
		IsGroup                  *bool `json:"IsGroup" name:"IsGroup"`
		SubEpcDeviceAttributeSet []struct {
			DeviceAttributeId *string `json:"DeviceAttributeId" name:"DeviceAttributeId"`
			HostType          *string `json:"HostType" name:"HostType"`
			HostTypeName      *string `json:"HostTypeName" name:"HostTypeName"`
			Memory            *string `json:"Memory" name:"Memory"`
			CpuDeviceSet      []struct {
				CpuSpec *string `json:"CpuSpec" name:"CpuSpec"`
			} `json:"CpuDeviceSet"`
			GpuDeviceSet []struct {
				GpuModel *string `json:"GpuModel" name:"GpuModel"`
				GpuCount *string `json:"GpuCount" name:"GpuCount"`
			} `json:"GpuDeviceSet"`
			PhysicalDiskDeviceSet []struct {
				DiskAttribute *string `json:"DiskAttribute" name:"DiskAttribute"`
				DiskCount     *string `json:"DiskCount" name:"DiskCount"`
				Space         *string `json:"Space" name:"Space"`
			} `json:"PhysicalDiskDeviceSet"`
			SubHostType     *string `json:"SubHostType" name:"SubHostType"`
			SubHostTypeName *string `json:"SubHostTypeName" name:"SubHostTypeName"`
			IsGroup         *bool   `json:"IsGroup" name:"IsGroup"`
			VpcNetworkCard  []struct {
				Type *string `json:"Type" name:"Type"`
				Num  *int    `json:"Num" name:"Num"`
			} `json:"VpcNetworkCard"`
			RdmaNetworkCard []struct {
				Type    *string `json:"Type" name:"Type"`
				Num     *int    `json:"Num" name:"Num"`
				Kind    *string `json:"Kind" name:"Kind"`
				UseType *string `json:"UseType" name:"UseType"`
			} `json:"RdmaNetworkCard"`
		} `json:"SubEpcDeviceAttributeSet" name:"SubEpcDeviceAttributeSet"`
		VpcNetworkCard []struct {
			Type *string `json:"Type" name:"Type"`
			Num  *int    `json:"Num" name:"Num"`
		} `json:"VpcNetworkCard" name:"VpcNetworkCard"`
		RdmaNetworkCard []struct {
			Type    *string `json:"Type" name:"Type"`
			Num     *int    `json:"Num" name:"Num"`
			Kind    *string `json:"Kind" name:"Kind"`
			UseType *string `json:"UseType" name:"UseType"`
		} `json:"RdmaNetworkCard" name:"RdmaNetworkCard"`
	} `json:"EpcDeviceAttributeSet"`
}

func (r *DescribeEpcDeviceAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEpcDeviceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesRequest struct {
	*ksyunhttp.BaseRequest
	OperationProcessId []*string                  `json:"OperationProcessId,omitempty" name:"OperationProcessId"`
	Filter             []*DescribeProcessesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults         *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken          *string                    `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeProcessesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	ProcessSet []struct {
		HostId                  *string `json:"HostId" name:"HostId"`
		Sn                      *string `json:"Sn" name:"Sn"`
		Region                  *string `json:"Region" name:"Region"`
		AvailabilityZone        *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Status                  *string `json:"Status" name:"Status"`
		Attribute               *string `json:"Attribute" name:"Attribute"`
		ProcessId               *string `json:"ProcessId" name:"ProcessId"`
		OperationProcessId      *string `json:"OperationProcessId" name:"OperationProcessId"`
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		FinishTime              *string `json:"FinishTime" name:"FinishTime"`
		Source                  *string `json:"Source" name:"Source"`
		Content                 *string `json:"Content" name:"Content"`
		MachineCount            *int    `json:"MachineCount" name:"MachineCount"`
		Title                   *string `json:"Title" name:"Title"`
		Type                    *string `json:"Type" name:"Type"`
		Confirm                 *string `json:"Confirm" name:"Confirm"`
		HostTypeName            *string `json:"HostTypeName" name:"HostTypeName"`
		HostName                *string `json:"HostName" name:"HostName"`
		CommunicationContentSet []struct {
			Remarks    *string `json:"Remarks" name:"Remarks"`
			Author     *string `json:"Author" name:"Author"`
			CreateTime *string `json:"CreateTime" name:"CreateTime"`
		} `json:"CommunicationContentSet" name:"CommunicationContentSet"`
		HostType                *string `json:"HostType" name:"HostType"`
		FaultTypeLv1            *string `json:"FaultTypeLv1" name:"FaultTypeLv1"`
		FaultTypeLv1Name        *string `json:"FaultTypeLv1Name" name:"FaultTypeLv1Name"`
		FaultTypeLv2            *string `json:"FaultTypeLv2" name:"FaultTypeLv2"`
		FaultTypeLv2Name        *string `json:"FaultTypeLv2Name" name:"FaultTypeLv2Name"`
		FaultTypeLv4            *string `json:"FaultTypeLv4" name:"FaultTypeLv4"`
		FaultTypeLv4Name        *string `json:"FaultTypeLv4Name" name:"FaultTypeLv4Name"`
		FaultTypeLv3            *string `json:"FaultTypeLv3" name:"FaultTypeLv3"`
		FaultTypeLv3Name        *string `json:"FaultTypeLv3Name" name:"FaultTypeLv3Name"`
		FaultMsg                *string `json:"FaultMsg" name:"FaultMsg"`
		ProcessSettlementType   *string `json:"ProcessSettlementType" name:"ProcessSettlementType"`
		ProcessSettlementReason *string `json:"ProcessSettlementReason" name:"ProcessSettlementReason"`
	} `json:"ProcessSet"`
}

func (r *DescribeProcessesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProcessesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessRequest struct {
	*ksyunhttp.BaseRequest
	ProcessId             *string   `json:"ProcessId,omitempty" name:"ProcessId"`
	InstanceId            *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	Sn                    *string   `json:"Sn,omitempty" name:"Sn"`
	AvailabilityZone      *string   `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	CreateTime            *string   `json:"CreateTime,omitempty" name:"CreateTime"`
	Attribute             *string   `json:"Attribute,omitempty" name:"Attribute"`
	Content               *string   `json:"Content,omitempty" name:"Content"`
	MachineCount          *int      `json:"MachineCount,omitempty" name:"MachineCount"`
	Title                 *string   `json:"Title,omitempty" name:"Title"`
	Type                  *string   `json:"Type,omitempty" name:"Type"`
	Confirm               *string   `json:"Confirm,omitempty" name:"Confirm"`
	ProcessSource         *int      `json:"ProcessSource,omitempty" name:"ProcessSource"`
	AutoNocCase           *int      `json:"autoNocCase,omitempty" name:"autoNocCase"`
	LogFileName           []*string `json:"LogFileName,omitempty" name:"LogFileName"`
	LogFile               []*string `json:"LogFile,omitempty" name:"LogFile"`
	LogUrl                []*string `json:"LogUrl,omitempty" name:"LogUrl"`
	AuthorizeCableReplace *int      `json:"AuthorizeCableReplace,omitempty" name:"AuthorizeCableReplace"`
	EventId               *string   `json:"EventId,omitempty" name:"EventId"`
}

func (r *CreateProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateProcessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
	ProcessId *string `json:"ProcessId" name:"ProcessId"`
}

func (r *CreateProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProcessRequest struct {
	*ksyunhttp.BaseRequest
	OperationProcessId *string `json:"OperationProcessId,omitempty" name:"OperationProcessId"`
}

func (r *DeleteProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteProcessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplyProcessRequest struct {
	*ksyunhttp.BaseRequest
	OperationProcessId *string `json:"OperationProcessId,omitempty" name:"OperationProcessId"`
	Remarks            *string `json:"Remarks,omitempty" name:"Remarks"`
}

func (r *ReplyProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReplyProcessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ReplyProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplyProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcTrashesRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcTrashesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcTrashesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	HostSet    []struct {
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		ComputerName                *string `json:"ComputerName" name:"ComputerName"`
		HostId                      *string `json:"HostId" name:"HostId"`
		HostName                    *string `json:"HostName" name:"HostName"`
		HostType                    *string `json:"HostType" name:"HostType"`
		AllowModifyHyperThreading   *bool   `json:"AllowModifyHyperThreading" name:"AllowModifyHyperThreading"`
		ReleasableTime              *string `json:"ReleasableTime" name:"ReleasableTime"`
		TorName                     *string `json:"TorName" name:"TorName"`
		CabinetName                 *string `json:"CabinetName" name:"CabinetName"`
		RackName                    *string `json:"RackName" name:"RackName"`
		Sn                          *string `json:"Sn" name:"Sn"`
		CabinetId                   *string `json:"CabinetId" name:"CabinetId"`
		AvailabilityZone            *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Raid                        *string `json:"Raid" name:"Raid"`
		RaidTemplateId              *string `json:"RaidTemplateId" name:"RaidTemplateId"`
		ImageId                     *string `json:"ImageId" name:"ImageId"`
		KeyId                       *string `json:"KeyId" name:"KeyId"`
		NetworkInterfaceMode        *string `json:"NetworkInterfaceMode" name:"NetworkInterfaceMode"`
		BondAttribute               *string `json:"BondAttribute" name:"BondAttribute"`
		EnableBond                  *bool   `json:"EnableBond" name:"EnableBond"`
		SecurityAgent               *string `json:"SecurityAgent" name:"SecurityAgent"`
		CloudMonitorAgent           *string `json:"CloudMonitorAgent" name:"CloudMonitorAgent"`
		SupportEbs                  *string `json:"SupportEbs" name:"SupportEbs"`
		ProductType                 *string `json:"ProductType" name:"ProductType"`
		OsName                      *string `json:"OsName" name:"OsName"`
		Memory                      *string `json:"Memory" name:"Memory"`
		HostStatus                  *string `json:"HostStatus" name:"HostStatus"`
		ClusterId                   *string `json:"ClusterId" name:"ClusterId"`
		EnableContainer             *bool   `json:"EnableContainer" name:"EnableContainer"`
		ProjectId                   *string `json:"ProjectId" name:"ProjectId"`
		SystemFileType              *string `json:"SystemFileType" name:"SystemFileType"`
		DataFileType                *string `json:"DataFileType" name:"DataFileType"`
		DataDiskCatalogue           *string `json:"DataDiskCatalogue" name:"DataDiskCatalogue"`
		DataDiskCatalogueSuffix     *string `json:"DataDiskCatalogueSuffix" name:"DataDiskCatalogueSuffix"`
		NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix" name:"NvmeDataDiskCatalogueSuffix"`
		NvmeDataDiskCatalogue       *string `json:"NvmeDataDiskCatalogue" name:"NvmeDataDiskCatalogue"`
		NvmeDataFileType            *string `json:"NvmeDataFileType" name:"NvmeDataFileType"`
		KesAgent                    *string `json:"KesAgent" name:"KesAgent"`
		KplAgent                    *string `json:"KplAgent" name:"KplAgent"`
		KmrAgent                    *string `json:"KmrAgent" name:"KmrAgent"`
		DiskSet                     []struct {
			DiskType        *string `json:"DiskType" name:"DiskType"`
			SystemDiskSpace *string `json:"SystemDiskSpace" name:"SystemDiskSpace"`
			Raid            *string `json:"Raid" name:"Raid"`
			DiskAttribute   *string `json:"DiskAttribute" name:"DiskAttribute"`
			Space           *string `json:"Space" name:"Space"`
			DiskCount       *int    `json:"DiskCount" name:"DiskCount"`
			DiskSpace       *string `json:"DiskSpace" name:"DiskSpace"`
		} `json:"DiskSet" name:"DiskSet"`
		NetworkInterfaceAttributeSet []struct {
			NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
			SubnetId             *string `json:"SubnetId" name:"SubnetId"`
			PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			DNS1                 *string `json:"DNS1" name:"DNS1"`
			DNS2                 *string `json:"DNS2" name:"DNS2"`
			Mac                  *string `json:"Mac" name:"Mac"`
			SecurityGroupSet     []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			} `json:"SecurityGroupSet"`
			VpcId *string `json:"VpcId" name:"VpcId"`
		} `json:"NetworkInterfaceAttributeSet" name:"NetworkInterfaceAttributeSet"`
		SystemVolumeType *string `json:"SystemVolumeType" name:"SystemVolumeType"`
		SystemVolumeSize *string `json:"SystemVolumeSize" name:"SystemVolumeSize"`
		DataVolumeSet    []struct {
			VolumeId           *string `json:"VolumeId" name:"VolumeId"`
			VolumeType         *string `json:"VolumeType" name:"VolumeType"`
			VolumeSize         *string `json:"VolumeSize" name:"VolumeSize"`
			DeleteWithInstance *bool   `json:"DeleteWithInstance" name:"DeleteWithInstance"`
		} `json:"DataVolumeSet" name:"DataVolumeSet"`
		GpuImageDriverId *string `json:"GpuImageDriverId" name:"GpuImageDriverId"`
		Tags             *string `json:"Tags" name:"Tags"`
		HyperThreading   *string `json:"HyperThreading" name:"HyperThreading"`
		RackId           *string `json:"RackId" name:"RackId"`
		ContainerAgent   *string `json:"ContainerAgent" name:"ContainerAgent"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		ChargeType       *string `json:"ChargeType" name:"ChargeType"`
		Cpu              struct {
			Model     *string `json:"Model" name:"Model"`
			Frequence *string `json:"Frequence" name:"Frequence"`
			Count     *int    `json:"Count" name:"Count"`
			CoreCount *int    `json:"CoreCount" name:"CoreCount"`
		} `json:"Cpu" name:"Cpu"`
		Gpu struct {
			Model     *string `json:"Model" name:"Model"`
			Frequence *string `json:"Frequence" name:"Frequence"`
			Count     *int    `json:"Count" name:"Count"`
			CoreCount *int    `json:"CoreCount" name:"CoreCount"`
			GpuCount  *int    `json:"GpuCount" name:"GpuCount"`
		} `json:"Gpu" name:"Gpu"`
		Roces []struct {
			Ip      *string `json:"Ip" name:"Ip"`
			Mask    *string `json:"Mask" name:"Mask"`
			GateWay *string `json:"GateWay" name:"GateWay"`
			Type    *string `json:"Type" name:"Type"`
		} `json:"Roces" name:"Roces"`
	} `json:"HostSet"`
}

func (r *DescribeEpcTrashesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEpcTrashesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *ReturnEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReturnEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ReturnEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceRequirementRequest struct {
	*ksyunhttp.BaseRequest
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	RequirementCount *int    `json:"RequirementCount,omitempty" name:"RequirementCount"`
	ProjectName      *string `json:"ProjectName,omitempty" name:"ProjectName"`
	UsageDate        *string `json:"UsageDate,omitempty" name:"UsageDate"`
	Description      *string `json:"Description,omitempty" name:"Description"`
	HostType         *string `json:"HostType,omitempty" name:"HostType"`
}

func (r *CreateResourceRequirementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateResourceRequirementResponse struct {
	*ksyunhttp.BaseResponse
	RequestId           *string `json:"RequestId" name:"RequestId"`
	ResourceRequirement struct {
		ResourceRequirementId                      *string `json:"ResourceRequirementId" name:"ResourceRequirementId"`
		HostType                                   *string `json:"HostType" name:"HostType"`
		Region                                     *string `json:"Region" name:"Region"`
		AvailabilityZone                           *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		RequirementCount                           *int    `json:"RequirementCount" name:"RequirementCount"`
		ReserveCount                               *int    `json:"ReserveCount" name:"ReserveCount"`
		ProjectName                                *string `json:"ProjectName" name:"ProjectName"`
		UsageDate                                  *string `json:"UsageDate" name:"UsageDate"`
		Description                                *string `json:"Description" name:"Description"`
		CreateTime                                 *string `json:"CreateTime" name:"CreateTime"`
		Progress                                   *string `json:"Progress" name:"Progress"`
		ResourceRequirementCommunicationContentSet []struct {
			Remarks    *string `json:"Remarks" name:"Remarks"`
			Author     *string `json:"Author" name:"Author"`
			CreateTime *string `json:"CreateTime" name:"CreateTime"`
		} `json:"ResourceRequirementCommunicationContentSet" name:"ResourceRequirementCommunicationContentSet"`
	} `json:"ResourceRequirement"`
}

func (r *CreateResourceRequirementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceRequirementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachVolumeRequest struct {
	*ksyunhttp.BaseRequest
	HostId   *string `json:"HostId,omitempty" name:"HostId"`
	VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *AttachVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AttachVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachVolumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachVolumeRequest struct {
	*ksyunhttp.BaseRequest
	HostId   *string `json:"HostId,omitempty" name:"HostId"`
	VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *DetachVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DetachVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachVolumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceRequest struct {
	*ksyunhttp.BaseRequest
	ChargeType       *string `json:"ChargeType,omitempty" name:"ChargeType"`
	HostType         *string `json:"HostType,omitempty" name:"HostType"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	PurchaseTime     *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	Amount           *int    `json:"Amount,omitempty" name:"Amount"`
}

func (r *DescribePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribePriceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	PriceInfo struct {
		PriceSet []struct {
			Currency      *string `json:"Currency" name:"Currency"`
			DiscountPrice *int    `json:"DiscountPrice" name:"DiscountPrice"`
			OriginalPrice *int    `json:"OriginalPrice" name:"OriginalPrice"`
			TradePrice    *int    `json:"TradePrice" name:"TradePrice"`
		} `json:"PriceSet" name:"PriceSet"`
	} `json:"PriceInfo"`
}

func (r *DescribePriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConfirmRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *UpdateConfirmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateConfirmResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *UpdateConfirmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConfirmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOverclockingAttributeRequest struct {
	*ksyunhttp.BaseRequest
	HostId                *string `json:"HostId,omitempty" name:"HostId"`
	OverclockingAttribute *string `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
}

func (r *ModifyOverclockingAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyOverclockingAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyOverclockingAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOverclockingAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyImageRequest struct {
	*ksyunhttp.BaseRequest
	DestinationName   *string `json:"DestinationName,omitempty" name:"DestinationName"`
	ImageId           *string `json:"ImageId,omitempty" name:"ImageId"`
	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	CopyTag           *string `json:"CopyTag,omitempty" name:"CopyTag"`
}

func (r *CopyImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CopyImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CopyImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcRaidAttributesRequest struct {
	*ksyunhttp.BaseRequest
	Filter     []*DescribeEpcRaidAttributesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                               `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                            `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcRaidAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEpcRaidAttributesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId           *string `json:"RequestId" name:"RequestId"`
	NextToken           *string `json:"NextToken" name:"NextToken"`
	TotalCount          *int64  `json:"TotalCount" name:"TotalCount"`
	EpcRaidAttributeSet []struct {
		TemplateName *string `json:"TemplateName" name:"TemplateName"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		RaidId       *string `json:"RaidId" name:"RaidId"`
		HostType     *string `json:"HostType" name:"HostType"`
		DiskSet      []struct {
			DiskAttribute   *string `json:"DiskAttribute" name:"DiskAttribute"`
			DiskCount       *string `json:"DiskCount" name:"DiskCount"`
			Space           *string `json:"Space" name:"Space"`
			SystemDiskSpace *string `json:"SystemDiskSpace" name:"SystemDiskSpace"`
			DiskSpace       *string `json:"DiskSpace" name:"DiskSpace"`
			Raid            *string `json:"Raid" name:"Raid"`
			DiskType        *string `json:"DiskType" name:"DiskType"`
			DiskId          *string `json:"DiskId" name:"DiskId"`
		} `json:"DiskSet" name:"DiskSet"`
	} `json:"EpcRaidAttributeSet"`
}

func (r *DescribeEpcRaidAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEpcRaidAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGpuImageDriverRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string `json:"NextToken,omitempty" name:"NextToken"`
	ImageId    *string `json:"ImageId,omitempty" name:"ImageId"`
	HostType   *string `json:"HostType,omitempty" name:"HostType"`
}

func (r *DescribeGpuImageDriverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeGpuImageDriverResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NextToken          *string `json:"NextToken" name:"NextToken"`
	TotalCount         *int64  `json:"TotalCount" name:"TotalCount"`
	GpuImagesDriverSet []struct {
		GpuImageDriverId *string   `json:"GpuImageDriverId" name:"GpuImageDriverId"`
		ImageNameSet     []*string `json:"ImageNameSet" name:"ImageNameSet"`
		GpuModel         []*string `json:"GpuModel" name:"GpuModel"`
	} `json:"GpuImagesDriverSet"`
}

func (r *DescribeGpuImageDriverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGpuImageDriverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateShareImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId   *string   `json:"ImageId,omitempty" name:"ImageId"`
	AccountId []*string `json:"AccountId,omitempty" name:"AccountId"`
}

func (r *CreateShareImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateShareImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CreateShareImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateShareImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShareImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId   *string   `json:"ImageId,omitempty" name:"ImageId"`
	AccountId []*string `json:"AccountId,omitempty" name:"AccountId"`
}

func (r *DeleteShareImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteShareImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteShareImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteShareImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShareImageAccountListRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeShareImageAccountListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeShareImageAccountListResponse struct {
	*ksyunhttp.BaseResponse
	SharePermissionSet []struct {
		AccountId *int64  `json:"AccountId" name:"AccountId"`
		ShareTime *string `json:"ShareTime" name:"ShareTime"`
		Status    *string `json:"Status" name:"Status"`
		ImageId   *string `json:"ImageId" name:"ImageId"`
	} `json:"SharePermissionSet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DescribeShareImageAccountListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShareImageAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShareImageRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeShareImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeShareImageResponse struct {
	*ksyunhttp.BaseResponse
	SharePermissionSet []struct {
		ImageName           *string `json:"ImageName" name:"ImageName"`
		System              *string `json:"System" name:"System"`
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		FromId              *string `json:"FromId" name:"FromId"`
		ImageId             *string `json:"ImageId" name:"ImageId"`
		ShareTime           *string `json:"ShareTime" name:"ShareTime"`
		ImageInitialization *string `json:"ImageInitialization" name:"ImageInitialization"`
		Status              *string `json:"Status" name:"Status"`
	} `json:"SharePermissionSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeShareImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShareImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptShareImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *AcceptShareImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AcceptShareImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AcceptShareImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptShareImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectShareImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *RejectShareImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RejectShareImageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RejectShareImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectShareImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeManagedAccessoryRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeManagedAccessoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeManagedAccessoryResponse struct {
	*ksyunhttp.BaseResponse
	ManagedAccessorySet []struct {
		SN             *string `json:"SN" name:"SN"`
		IDC            *string `json:"IDC" name:"IDC"`
		Classification *string `json:"Classification" name:"Classification"`
		Model          *string `json:"Model" name:"Model"`
		Manufacturer   *string `json:"Manufacturer" name:"Manufacturer"`
		State          *string `json:"State" name:"State"`
		Date           *string `json:"Date" name:"Date"`
		Source         *string `json:"Source" name:"Source"`
		Notes          *string `json:"Notes" name:"Notes"`
		Num            *int    `json:"Num" name:"Num"`
		AlarmNum       *int    `json:"AlarmNum" name:"AlarmNum"`
	} `json:"ManagedAccessorySet"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DescribeManagedAccessoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeManagedAccessoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoDeleteEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId         *string `json:"HostId,omitempty" name:"HostId"`
	AutoDeleteTime *string `json:"AutoDeleteTime,omitempty" name:"AutoDeleteTime"`
	AutoDeleteEip  *string `json:"AutoDeleteEip,omitempty" name:"AutoDeleteEip"`
}

func (r *AutoDeleteEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AutoDeleteEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AutoDeleteEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AutoDeleteEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportImageRequest struct {
	*ksyunhttp.BaseRequest
	ImageId    *string `json:"ImageId,omitempty" name:"ImageId"`
	Ks3Bucket  *string `json:"Ks3Bucket,omitempty" name:"Ks3Bucket"`
	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
}

func (r *ExportImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ExportImageResponse struct {
	*ksyunhttp.BaseResponse
	Return    *bool   `json:"Return" name:"Return"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ExportImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBucketsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *QueryBucketsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryBucketsResponse struct {
	*ksyunhttp.BaseResponse
	BucketSet []struct {
		BucketName           *string `json:"BucketName" name:"BucketName"`
		BucketHost           *string `json:"BucketHost" name:"BucketHost"`
		BucketHostCompatible *string `json:"BucketHostCompatible" name:"BucketHostCompatible"`
		BucketInnerHost      *string `json:"BucketInnerHost" name:"BucketInnerHost"`
		Endpoint             *string `json:"Endpoint" name:"Endpoint"`
	} `json:"BucketSet"`
	Code *int `json:"Code" name:"Code"`
}

func (r *QueryBucketsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBucketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelImageExportRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *CancelImageExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CancelImageExportResponse struct {
	*ksyunhttp.BaseResponse
	Return    *bool   `json:"Return" name:"Return"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CancelImageExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelImageExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UseHotStandByEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId             *string `json:"HostId,omitempty" name:"HostId"`
	HotStandByHostId   *string `json:"HotStandByHostId,omitempty" name:"HotStandByHostId"`
	RetainInstanceInfo *string `json:"RetainInstanceInfo,omitempty" name:"RetainInstanceInfo"`
}

func (r *UseHotStandByEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UseHotStandByEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *UseHotStandByEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UseHotStandByEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateHotStandbyEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *ActivateHotStandbyEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateHotStandbyEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ActivateHotStandbyEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateHotStandbyEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateEpcRequest struct {
	*ksyunhttp.BaseRequest
	HostType                    *string   `json:"HostType,omitempty" name:"HostType"`
	AvailabilityZone            *string   `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Raid                        *string   `json:"Raid,omitempty" name:"Raid"`
	RaidId                      *string   `json:"RaidId,omitempty" name:"RaidId"`
	ImageId                     *string   `json:"ImageId,omitempty" name:"ImageId"`
	NetworkInterfaceMode        *string   `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
	SubnetId                    *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	KeyId                       *string   `json:"keyId,omitempty" name:"keyId"`
	SecurityGroupId             []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	DNS1                        *string   `json:"DNS1,omitempty" name:"DNS1"`
	DNS2                        *string   `json:"DNS2,omitempty" name:"DNS2"`
	HostName                    *string   `json:"HostName,omitempty" name:"HostName"`
	ProjectId                   *string   `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType                  *string   `json:"ChargeType,omitempty" name:"ChargeType"`
	Sn                          *string   `json:"Sn,omitempty" name:"Sn"`
	PurchaseTime                *int      `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	Password                    *string   `json:"Password,omitempty" name:"Password"`
	CloudMonitorAgent           *string   `json:"CloudMonitorAgent,omitempty" name:"CloudMonitorAgent"`
	ExtensionSubnetId           *string   `json:"ExtensionSubnetId,omitempty" name:"ExtensionSubnetId"`
	ExtensionDNS1               *string   `json:"ExtensionDNS1,omitempty" name:"ExtensionDNS1"`
	ExtensionDNS2               *string   `json:"ExtensionDNS2,omitempty" name:"ExtensionDNS2"`
	Description                 *string   `json:"Description,omitempty" name:"Description"`
	AddressBandWidth            *string   `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
	LineId                      *string   `json:"LineId,omitempty" name:"LineId"`
	BandWidthShareId            *string   `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	AddressChargeType           *string   `json:"AddressChargeType,omitempty" name:"AddressChargeType"`
	AddressPurchaseTime         *string   `json:"AddressPurchaseTime,omitempty" name:"AddressPurchaseTime"`
	AddressProjectId            *string   `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
	SystemFileType              *string   `json:"SystemFileType,omitempty" name:"SystemFileType"`
	DataFileType                *string   `json:"DataFileType,omitempty" name:"DataFileType"`
	DataDiskCatalogue           *string   `json:"DataDiskCatalogue,omitempty" name:"DataDiskCatalogue"`
	DataDiskCatalogueSuffix     *string   `json:"DataDiskCatalogueSuffix,omitempty" name:"DataDiskCatalogueSuffix"`
	HyperThreading              *string   `json:"HyperThreading,omitempty" name:"HyperThreading"`
	NvmeDataFileType            *string   `json:"NvmeDataFileType,omitempty" name:"NvmeDataFileType"`
	NvmeDataDiskCatalogue       *string   `json:"NvmeDataDiskCatalogue,omitempty" name:"NvmeDataDiskCatalogue"`
	NvmeDataDiskCatalogueSuffix *string   `json:"NvmeDataDiskCatalogueSuffix,omitempty" name:"NvmeDataDiskCatalogueSuffix"`
	BondAttribute               *string   `json:"BondAttribute,omitempty" name:"BondAttribute"`
	ContainerAgent              *string   `json:"ContainerAgent,omitempty" name:"ContainerAgent"`
	KesAgent                    *string   `json:"KesAgent,omitempty" name:"KesAgent"`
	KmrAgent                    *string   `json:"KmrAgent,omitempty" name:"KmrAgent"`
	ComputerName                *string   `json:"ComputerName,omitempty" name:"ComputerName"`
	OverclockingAttribute       *string   `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
	GpuImageDriverId            *string   `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
	SystemVolumeType            *string   `json:"SystemVolumeType,omitempty" name:"SystemVolumeType"`
	SystemVolumeSize            *string   `json:"SystemVolumeSize,omitempty" name:"SystemVolumeSize"`
	RoceNetwork                 *string   `json:"RoceNetwork,omitempty" name:"RoceNetwork"`
	ZoneId                      *string   `json:"ZoneId,omitempty" name:"ZoneId"`
	ZoneType                    *string   `json:"ZoneType,omitempty" name:"ZoneType"`
	HostNameStartNo             *int      `json:"HostNameStartNo,omitempty" name:"HostNameStartNo"`
	ComputerNameStartNo         *int      `json:"ComputerNameStartNo,omitempty" name:"ComputerNameStartNo"`
	Amount                      *int      `json:"Amount,omitempty" name:"Amount"`
	TimedRegularization         *string   `json:"TimedRegularization,omitempty" name:"TimedRegularization"`
	PasswordInherit             *string   `json:"PasswordInherit,omitempty" name:"PasswordInherit"`
	DataDiskMount               *string   `json:"DataDiskMount,omitempty" name:"DataDiskMount"`
	StorageRoceNetworkCardName  *string   `json:"StorageRoceNetworkCardName,omitempty" name:"StorageRoceNetworkCardName"`
	SRoceCluster                *string   `json:"SRoceCluster,omitempty" name:"SRoceCluster"`
	RoceCluster                 *string   `json:"RoceCluster,omitempty" name:"RoceCluster"`
}

func (r *BatchCreateEpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type BatchCreateEpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Host      []struct {
		HostId   *string `json:"HostId" name:"HostId"`
		HostName *string `json:"HostName" name:"HostName"`
	} `json:"Host"`
}

func (r *BatchCreateEpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateEpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUseHotStandbyRecordsRequest struct {
	*ksyunhttp.BaseRequest
	FilterN    *DescribeUseHotStandbyRecordsFilterN `json:"Filter.N,omitempty" name:"Filter.N"`
	MaxResults *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                              `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeUseHotStandbyRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeUseHotStandbyRecordsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	NextToken               *string `json:"NextToken" name:"NextToken"`
	TotalCount              *int    `json:"TotalCount" name:"TotalCount"`
	UseHotStandbyRecordsSet []struct {
		FaultHostName        *string `json:"FaultHostName" name:"FaultHostName"`
		FaultInstanceId      *string `json:"FaultInstanceId" name:"FaultInstanceId"`
		FaultSn              *string `json:"FaultSn" name:"FaultSn"`
		FaultHostType        *string `json:"FaultHostType" name:"FaultHostType"`
		FaultPrivateIp       *string `json:"FaultPrivateIp" name:"FaultPrivateIp"`
		FaultEip             *string `json:"FaultEip" name:"FaultEip"`
		ReplaceType          *string `json:"ReplaceType" name:"ReplaceType"`
		HotStandbyHostName   *string `json:"HotStandbyHostName" name:"HotStandbyHostName"`
		HotStandbyInstanceId *string `json:"HotStandbyInstanceId" name:"HotStandbyInstanceId"`
		HotStandbySn         *string `json:"HotStandbySn" name:"HotStandbySn"`
		HotStandbyHostType   *string `json:"HotStandbyHostType" name:"HotStandbyHostType"`
		HotStandbyPrivateIp  *string `json:"HotStandbyPrivateIp" name:"HotStandbyPrivateIp"`
		HotStandbyEip        *string `json:"HotStandbyEip" name:"HotStandbyEip"`
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
	} `json:"UseHotStandbyRecordsSet"`
}

func (r *DescribeUseHotStandbyRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUseHotStandbyRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGpuRoceTopologyRequest struct {
	*ksyunhttp.BaseRequest
	SpineName *string `json:"SpineName,omitempty" name:"SpineName"`
}

func (r *DescribeGpuRoceTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeGpuRoceTopologyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	DescribePorts struct {
		LeafDatas []struct {
			SpineName *string `json:"SpineName" name:"SpineName"`
			LeafName  *string `json:"LeafName" name:"LeafName"`
			SpinePort *string `json:"SpinePort" name:"SpinePort"`
			LeafPort  *string `json:"LeafPort" name:"LeafPort"`
		} `json:"LeafDatas" name:"LeafDatas"`
		CoreDatas []struct {
			CoreName  *string `json:"CoreName" name:"CoreName"`
			SpineName *string `json:"SpineName" name:"SpineName"`
			CorePort  *string `json:"CorePort" name:"CorePort"`
			SpinePort *string `json:"SpinePort" name:"SpinePort"`
		} `json:"CoreDatas" name:"CoreDatas"`
	} `json:"DescribePorts"`
}

func (r *DescribeGpuRoceTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGpuRoceTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProcessRequest struct {
	*ksyunhttp.BaseRequest
	OperationProcessId *string `json:"OperationProcessId,omitempty" name:"OperationProcessId"`
	Confirm            *string `json:"Confirm,omitempty" name:"Confirm"`
	Status             *string `json:"Status,omitempty" name:"Status"`
	Content            *string `json:"Content,omitempty" name:"Content"`
}

func (r *ModifyProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyProcessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ConfirmTime *string `json:"ConfirmTime" name:"ConfirmTime"`
	Status      *string `json:"Status" name:"Status"`
	Return      *bool   `json:"Return" name:"Return"`
}

func (r *ModifyProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmProcessRequest struct {
	*ksyunhttp.BaseRequest
	OperationProcessId   *string `json:"OperationProcessId,omitempty" name:"OperationProcessId"`
	UserConfirmAvailable *string `json:"UserConfirmAvailable,omitempty" name:"UserConfirmAvailable"`
}

func (r *ConfirmProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ConfirmProcessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	UserConfirmAvailable     *string `json:"UserConfirmAvailable" name:"UserConfirmAvailable"`
	UserConfirmAvailableTime *string `json:"UserConfirmAvailableTime" name:"UserConfirmAvailableTime"`
	Return                   *bool   `json:"Return" name:"Return"`
}

func (r *ConfirmProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelConfigRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults       *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken        *string `json:"NextToken,omitempty" name:"NextToken"`
	ImageId          *string `json:"ImageId,omitempty" name:"ImageId"`
	HostType         *string `json:"HostType,omitempty" name:"HostType"`
	GpuImageDriverId *string `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
}

func (r *DescribeModelConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModelConfigResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount     *int    `json:"TotalCount" name:"TotalCount"`
	NextToken      *string `json:"NextToken" name:"NextToken"`
	ModelConfigSet []struct {
		GpuImageDriverId *string `json:"GpuImageDriverId" name:"GpuImageDriverId"`
		ImageName        *string `json:"ImageName" name:"ImageName"`
		GpuModel         *string `json:"GpuModel" name:"GpuModel"`
		Anaconda         *string `json:"Anaconda" name:"Anaconda"`
		Framework        *string `json:"Framework" name:"Framework"`
		Engine           *string `json:"Engine" name:"Engine"`
		AiModel          *string `json:"AiModel" name:"AiModel"`
	} `json:"ModelConfigSet"`
}

func (r *DescribeModelConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModelConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoceEventRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int                      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                   `json:"NextToken,omitempty" name:"NextToken"`
	FilterN    *DescribeRoceEventFilterN `json:"Filter.N,omitempty" name:"Filter.N"`
	HostIdN    *string                   `json:"HostId.N,omitempty" name:"HostId.N"`
}

func (r *DescribeRoceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRoceEventResponse struct {
	*ksyunhttp.BaseResponse
	RoceEventResponseSet []struct {
		EventTime         *string `json:"EventTime" name:"EventTime"`
		EventName         *string `json:"EventName" name:"EventName"`
		EventStatus       *int    `json:"EventStatus" name:"EventStatus"`
		EventId           *string `json:"EventId" name:"EventId"`
		InstanceName      *string `json:"InstanceName" name:"InstanceName"`
		InstanceId        *string `json:"InstanceId" name:"InstanceId"`
		Sn                *string `json:"Sn" name:"Sn"`
		EventDesc         *string `json:"EventDesc" name:"EventDesc"`
		RoceIp            *string `json:"RoceIp" name:"RoceIp"`
		RoceType          *string `json:"RoceType" name:"RoceType"`
		RoceInterfaceName *string `json:"RoceInterfaceName" name:"RoceInterfaceName"`
		AvailabilityZone  *string `json:"AvailabilityZone" name:"AvailabilityZone"`
	} `json:"RoceEventResponseSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeRoceEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoceEventDetailsRequest struct {
	*ksyunhttp.BaseRequest
	EventId    *string `json:"EventId,omitempty" name:"EventId"`
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeRoceEventDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRoceEventDetailsResponse struct {
	*ksyunhttp.BaseResponse
	RoceEventDetailResponseSet []struct {
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		Sn           *string `json:"Sn" name:"Sn"`
		EventStatus  *string `json:"EventStatus" name:"EventStatus"`
		EventTime    *string `json:"EventTime" name:"EventTime"`
		EventName    *string `json:"EventName" name:"EventName"`
		HostName     *string `json:"HostName" name:"HostName"`
		PortName     *string `json:"PortName" name:"PortName"`
	} `json:"RoceEventDetailResponseSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeRoceEventDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoceEventDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateProcessRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId       []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	AvailabilityZone *string   `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Attribute        *string   `json:"Attribute,omitempty" name:"Attribute"`
	Content          *string   `json:"Content,omitempty" name:"Content"`
	LogFileName      []*string `json:"LogFileName,omitempty" name:"LogFileName"`
	LogFile          []*string `json:"LogFile,omitempty" name:"LogFile"`
	LogUrl           []*string `json:"LogUrl,omitempty" name:"LogUrl"`
	MachineCount     *int      `json:"MachineCount,omitempty" name:"MachineCount"`
	Title            *string   `json:"Title,omitempty" name:"Title"`
	Type             *string   `json:"Type,omitempty" name:"Type"`
	Confirm          *string   `json:"Confirm,omitempty" name:"Confirm"`
	ProcessSource    *int      `json:"ProcessSource,omitempty" name:"ProcessSource"`
	AutoNocCase      *int      `json:"AutoNocCase,omitempty" name:"AutoNocCase"`
}

func (r *BatchCreateProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type BatchCreateProcessResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ProcessInfo []struct {
		ProcessId  *string `json:"ProcessId" name:"ProcessId"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Sn         *string `json:"Sn" name:"Sn"`
	} `json:"ProcessInfo"`
	Return *bool `json:"Return" name:"Return"`
}

func (r *BatchCreateProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInspectHostRequest struct {
	*ksyunhttp.BaseRequest
	InspectType *string   `json:"InspectType,omitempty" name:"InspectType"`
	InspectName *string   `json:"InspectName,omitempty" name:"InspectName"`
	HostId      []*string `json:"HostId,omitempty" name:"HostId"`
}

func (r *CreateInspectHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInspectHostResponse struct {
	*ksyunhttp.BaseResponse
	InspectId *string `json:"InspectId" name:"InspectId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CreateInspectHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInspectHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInspectHostResultsRequest struct {
	*ksyunhttp.BaseRequest
	InspectId  []*string `json:"InspectId,omitempty" name:"InspectId"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeInspectHostResultsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInspectHostResultsResponse struct {
	*ksyunhttp.BaseResponse
	InspectResultSet []struct {
		InspectId              *string `json:"InspectId" name:"InspectId"`
		InspectName            *string `json:"InspectName" name:"InspectName"`
		GpuModel               *string `json:"GpuModel" name:"GpuModel"`
		Result                 *string `json:"Result" name:"Result"`
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime             *string `json:"UpdateTime" name:"UpdateTime"`
		SuccessCount           *int    `json:"SuccessCount" name:"SuccessCount"`
		FailCount              *int    `json:"FailCount" name:"FailCount"`
		InspectResultDetailSet []struct {
			HostId        *string `json:"HostId" name:"HostId"`
			Sn            *string `json:"Sn" name:"Sn"`
			Result        *string `json:"Result" name:"Result"`
			InspectStatus *string `json:"InspectStatus" name:"InspectStatus"`
			CreateTime    *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		} `json:"InspectResultDetailSet" name:"InspectResultDetailSet"`
	} `json:"InspectResultSet"`
	InspectHostCount *int    `json:"InspectHostCount" name:"InspectHostCount"`
	RequestId        *string `json:"RequestId" name:"RequestId"`
	TotalCount       *int    `json:"TotalCount" name:"TotalCount"`
	NextToken        *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeInspectHostResultsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInspectHostResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeXidDetailsRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Name       *string `json:"Name,omitempty" name:"Name"`
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeXidDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeXidDetailsResponse struct {
	*ksyunhttp.BaseResponse
	XidDetailSet []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		Sn           *string `json:"Sn" name:"Sn"`
		EventName    *string `json:"EventName" name:"EventName"`
		EventTime    *string `json:"EventTime" name:"EventTime"`
		KcmName      *string `json:"KcmName" name:"KcmName"`
		Xid          *string `json:"Xid" name:"Xid"`
		EventMsg     *string `json:"EventMsg" name:"EventMsg"`
	} `json:"XidDetailSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeXidDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeXidDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSoInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ImageId                *string                `json:"ImageId,omitempty" name:"ImageId"`
	InstanceName           *string                `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceTypeId         *string                `json:"InstanceTypeId,omitempty" name:"InstanceTypeId"`
	SecurityGroupId        []*string              `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SubnetId               *string                `json:"SubnetId,omitempty" name:"SubnetId"`
	Volumes                *RunSoInstancesVolumes `json:"Volumes,omitempty" name:"Volumes"`
	ZoneId                 *string                `json:"ZoneId,omitempty" name:"ZoneId"`
	Description            *string                `json:"Description,omitempty" name:"Description"`
	HostName               *string                `json:"HostName,omitempty" name:"HostName"`
	InstanceChargeType     *string                `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	KeepImageCredential    *bool                  `json:"KeepImageCredential,omitempty" name:"KeepImageCredential"`
	KeyPairName            *string                `json:"KeyPairName,omitempty" name:"KeyPairName"`
	Password               *string                `json:"Password,omitempty" name:"Password"`
	Period                 *int                   `json:"Period,omitempty" name:"Period"`
	SuffixIndex            *int                   `json:"SuffixIndex,omitempty" name:"SuffixIndex"`
	UniqueSuffix           *bool                  `json:"UniqueSuffix,omitempty" name:"UniqueSuffix"`
	InstallRunCommandAgent *bool                  `json:"InstallRunCommandAgent,omitempty" name:"InstallRunCommandAgent"`
	Count                  *int                   `json:"Count,omitempty" name:"Count"`
	SoZoneId               *string                `json:"SoZoneId,omitempty" name:"SoZoneId"`
	UserData               *string                `json:"UserData,omitempty" name:"UserData"`
}

func (r *RunSoInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RunSoInstancesResponse struct {
	*ksyunhttp.BaseResponse
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds"`
	RequestId   *string   `json:"RequestId" name:"RequestId"`
	Return      *bool     `json:"Return" name:"Return"`
}

func (r *RunSoInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunSoInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoImagesRequest struct {
	*ksyunhttp.BaseRequest
	ImageId            []*string `json:"ImageId,omitempty" name:"ImageId"`
	ImageName          *string   `json:"ImageName,omitempty" name:"ImageName"`
	IsSupportCloudInit *bool     `json:"IsSupportCloudInit,omitempty" name:"IsSupportCloudInit"`
	MaxResults         *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken          *string   `json:"NextToken,omitempty" name:"NextToken"`
	OsType             *string   `json:"OsType,omitempty" name:"OsType"`
	Platform           *string   `json:"Platform,omitempty" name:"Platform"`
	Status             []*string `json:"Status,omitempty" name:"Status"`
	Visibility         *string   `json:"Visibility,omitempty" name:"Visibility"`
	SoZoneId           *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoImagesResponse struct {
	*ksyunhttp.BaseResponse
	Images []struct {
		Architecture       *string `json:"Architecture" name:"Architecture"`
		BootMode           *string `json:"BootMode" name:"BootMode"`
		CreatedAt          *string `json:"CreatedAt" name:"CreatedAt"`
		Description        *string `json:"Description" name:"Description"`
		ImageId            *string `json:"ImageId" name:"ImageId"`
		ImageName          *string `json:"ImageName" name:"ImageName"`
		IsSupportCloudInit *bool   `json:"IsSupportCloudInit" name:"IsSupportCloudInit"`
		Kernel             *string `json:"Kernel" name:"Kernel"`
		OsName             *string `json:"OsName" name:"OsName"`
		OsType             *string `json:"OsType" name:"OsType"`
		Status             *string `json:"Status" name:"Status"`
		Platform           *string `json:"Platform" name:"Platform"`
		Size               *int    `json:"Size" name:"Size"`
		UpdatedAt          *string `json:"UpdatedAt" name:"UpdatedAt"`
		VirtualSize        *int    `json:"VirtualSize" name:"VirtualSize"`
		Visibility         *string `json:"Visibility" name:"Visibility"`
	} `json:"Images"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DescribeSoImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootSoInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ForceStop   *bool     `json:"ForceStop,omitempty" name:"ForceStop"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *RebootSoInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RebootSoInstanceResponse struct {
	*ksyunhttp.BaseResponse
	ImageId   *string `json:"ImageId" name:"ImageId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *RebootSoInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootSoInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSoImagesRequest struct {
	*ksyunhttp.BaseRequest
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	SoZoneId *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DeleteSoImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSoImagesResponse struct {
	*ksyunhttp.BaseResponse
	OperationDetails []struct {
		ImageId *string `json:"ImageId" name:"ImageId"`
		Error   struct {
			Code    *string `json:"Code" name:"Code"`
			Message *string `json:"Message" name:"Message"`
		} `json:"Error" name:"Error"`
	} `json:"OperationDetails"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteSoImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSoVpcRequest struct {
	*ksyunhttp.BaseRequest
	VpcId    *string `json:"VpcId,omitempty" name:"VpcId"`
	SoZoneId *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DeleteSoVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSoVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteSoVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoAvailableResourceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	InstanceTypeId     *string `json:"InstanceTypeId,omitempty" name:"InstanceTypeId"`
	ZoneId             *string `json:"ZoneId,omitempty" name:"ZoneId"`
	SoZoneId           *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoAvailableResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoAvailableResourceResponse struct {
	*ksyunhttp.BaseResponse
	AvailableZones []struct {
		AvailableResources []struct {
			SupportedResources []struct {
				Status *string `json:"Status" name:"Status"`
				Value  *string `json:"Value" name:"Value"`
			} `json:"SupportedResources"`
			Type *string `json:"Type" name:"Type"`
		} `json:"AvailableResources" name:"AvailableResources"`
		RegionId *string `json:"RegionId" name:"RegionId"`
		Status   *string `json:"Status" name:"Status"`
		ZoneId   *string `json:"ZoneId" name:"ZoneId"`
	} `json:"AvailableZones"`
}

func (r *DescribeSoAvailableResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoAvailableResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceChargeType *string   `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	InstanceTypeId     []*string `json:"InstanceTypeId,omitempty" name:"InstanceTypeId"`
	KeyPairName        *string   `json:"KeyPairName,omitempty" name:"KeyPairName"`
	MaxResults         *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken          *string   `json:"NextToken,omitempty" name:"NextToken"`
	PrimaryIpAddress   *string   `json:"PrimaryIpAddress,omitempty" name:"PrimaryIpAddress"`
	Status             *string   `json:"Status,omitempty" name:"Status"`
	VpcId              *string   `json:"VpcId,omitempty" name:"VpcId"`
	ZoneId             *string   `json:"ZoneId,omitempty" name:"ZoneId"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	SoZoneId           *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Instances []struct {
		CreatedAt         *string `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt         *string `json:"UpdatedAt" name:"UpdatedAt"`
		ZoneId            *string `json:"ZoneId" name:"ZoneId"`
		ImageId           *string `json:"ImageId" name:"ImageId"`
		Status            *string `json:"Status" name:"Status"`
		InstanceName      *string `json:"InstanceName" name:"InstanceName"`
		Description       *string `json:"Description" name:"Description"`
		VpcId             *string `json:"VpcId" name:"VpcId"`
		NetworkInterfaces []struct {
			VpcId            *string `json:"VpcId" name:"VpcId"`
			SubnetId         *string `json:"SubnetId" name:"SubnetId"`
			PrimaryIpAddress *string `json:"PrimaryIpAddress" name:"PrimaryIpAddress"`
			Type             *string `json:"Type" name:"Type"`
			MacAddress       *string `json:"MacAddress" name:"MacAddress"`
			Ipv6Addresses    []struct {
			} `json:"Ipv6Addresses"`
			SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds"`
		} `json:"NetworkInterfaces" name:"NetworkInterfaces"`
		RdmaIpAddresses []struct {
		} `json:"RdmaIpAddresses" name:"RdmaIpAddresses"`
		KeyPairName        *string `json:"KeyPairName" name:"KeyPairName"`
		KeyPairId          *string `json:"KeyPairId" name:"KeyPairId"`
		StoppedMode        *string `json:"StoppedMode" name:"StoppedMode"`
		InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
		InstanceTypeId     *string `json:"InstanceTypeId" name:"InstanceTypeId"`
		ExpiredAt          *string `json:"ExpiredAt" name:"ExpiredAt"`
		OsType             *string `json:"OsType" name:"OsType"`
		OsName             *string `json:"OsName" name:"OsName"`
		Cpus               *int    `json:"Cpus" name:"Cpus"`
		MemorySize         *int    `json:"MemorySize" name:"MemorySize"`
		InstanceId         *string `json:"InstanceId" name:"InstanceId"`
		Hostname           *string `json:"Hostname" name:"Hostname"`
		Uuid               *string `json:"Uuid" name:"Uuid"`
		LocalVolumes       []struct {
			VolumeType *string `json:"VolumeType" name:"VolumeType"`
			Size       *int    `json:"Size" name:"Size"`
			Count      *int    `json:"Count" name:"Count"`
		} `json:"LocalVolumes" name:"LocalVolumes"`
		CpuOptions struct {
			CoreCount      *int `json:"CoreCount" name:"CoreCount"`
			ThreadsPerCore *int `json:"ThreadsPerCore" name:"ThreadsPerCore"`
		} `json:"CpuOptions" name:"CpuOptions"`
		DeletionProtection *bool `json:"DeletionProtection" name:"DeletionProtection"`
		SystemDisk         struct {
			Type   *string `json:"Type" name:"Type"`
			DiskPl *string `json:"DiskPl" name:"DiskPl"`
			Size   *int    `json:"Size" name:"Size"`
			Count  *int    `json:"Count" name:"Count"`
		} `json:"SystemDisk" name:"SystemDisk"`
		DataDisk struct {
			Type   *string `json:"Type" name:"Type"`
			DiskPl *string `json:"DiskPl" name:"DiskPl"`
			Size   *int    `json:"Size" name:"Size"`
			Count  *int    `json:"Count" name:"Count"`
		} `json:"DataDisk" name:"DataDisk"`
	} `json:"Instances"`
}

func (r *DescribeSoInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSoInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DeleteSoInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSoInstanceResponse struct {
	*ksyunhttp.BaseResponse
	OperationDetails []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Error      struct {
			Code    *string `json:"Code" name:"Code"`
			Message *string `json:"Message" name:"Message"`
		} `json:"Error" name:"Error"`
	} `json:"OperationDetails"`
}

func (r *DeleteSoInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoSecurityGroupsRequest struct {
	*ksyunhttp.BaseRequest
	VpcId            *string   `json:"VpcId,omitempty" name:"VpcId"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	NextToken        *string   `json:"NextToken,omitempty" name:"NextToken"`
	MaxResults       *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	SoZoneId         *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoSecurityGroupsResponse struct {
	*ksyunhttp.BaseResponse
	SecurityGroups []struct {
		VpcId             *string `json:"VpcId" name:"VpcId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		Description       *string `json:"Description" name:"Description"`
		SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Status            *string `json:"Status" name:"Status"`
		CreationTime      *string `json:"CreationTime" name:"CreationTime"`
		Type              *string `json:"Type" name:"Type"`
	} `json:"SecurityGroups"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DescribeSoSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoVpcRequest struct {
	*ksyunhttp.BaseRequest
	VpcName     *string   `json:"VpcName,omitempty" name:"VpcName"`
	Description *string   `json:"Description,omitempty" name:"Description"`
	CidrBlock   *string   `json:"CidrBlock,omitempty" name:"CidrBlock"`
	DnsServers  []*string `json:"DnsServers,omitempty" name:"DnsServers"`
	AttachVpcId *string   `json:"AttachVpcId,omitempty" name:"AttachVpcId"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *CreateSoVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSoVpcResponse struct {
	*ksyunhttp.BaseResponse
	VpcId     *string `json:"VpcId" name:"VpcId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CreateSoVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSoSubnetRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	SoZoneId *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DeleteSoSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSoSubnetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSoSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoKeyPairsRequest struct {
	*ksyunhttp.BaseRequest
	FingerPrint  *string   `json:"FingerPrint,omitempty" name:"FingerPrint"`
	KeyPairIds   []*string `json:"KeyPairIds,omitempty" name:"KeyPairIds"`
	KeyPairName  *string   `json:"KeyPairName,omitempty" name:"KeyPairName"`
	KeyPairNames []*string `json:"KeyPairNames,omitempty" name:"KeyPairNames"`
	MaxResults   *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string   `json:"NextToken,omitempty" name:"NextToken"`
	SoZoneId     *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoKeyPairsResponse struct {
	*ksyunhttp.BaseResponse
	KeyPairs []struct {
		CreatedAt   *string `json:"CreatedAt" name:"CreatedAt"`
		Description *string `json:"Description" name:"Description"`
		FingerPrint *string `json:"FingerPrint" name:"FingerPrint"`
		KeyPairId   *string `json:"KeyPairId" name:"KeyPairId"`
		KeyPairName *string `json:"KeyPairName" name:"KeyPairName"`
		UpdatedAt   *string `json:"UpdatedAt" name:"UpdatedAt"`
	} `json:"KeyPairs"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DescribeSoKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSoInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *StartSoInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartSoInstanceResponse struct {
	*ksyunhttp.BaseResponse
	OperationDetails []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Error      struct {
			Code    *string `json:"Code" name:"Code"`
			Message *string `json:"Message" name:"Message"`
		} `json:"Error" name:"Error"`
	} `json:"OperationDetails"`
}

func (r *StartSoInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSoInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoInstanceTypesRequest struct {
	*ksyunhttp.BaseRequest
	ImageId        *string   `json:"ImageId,omitempty" name:"ImageId"`
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" name:"InstanceTypeId"`
	SoZoneId       *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoInstanceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoInstanceTypesResponse struct {
	*ksyunhttp.BaseResponse
	InstanceTypes []struct {
		Gpu struct {
		} `json:"Gpu" name:"Gpu"`
		Rdma struct {
		} `json:"Rdma" name:"Rdma"`
		Processor struct {
			Cpus           *int     `json:"Cpus" name:"Cpus"`
			Model          *string  `json:"Model" name:"Model"`
			BaseFrequency  *float64 `json:"BaseFrequency" name:"BaseFrequency"`
			TurboFrequency *int     `json:"TurboFrequency" name:"TurboFrequency"`
		} `json:"Processor" name:"Processor"`
		Memory struct {
			Size          *int `json:"Size" name:"Size"`
			EncryptedSize *int `json:"EncryptedSize" name:"EncryptedSize"`
		} `json:"Memory" name:"Memory"`
		LocalVolumes []struct {
		} `json:"LocalVolumes" name:"LocalVolumes"`
	} `json:"InstanceTypes"`
	NextToken *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeSoInstanceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoSubnetAttributesRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId    *string `json:"SubnetId,omitempty" name:"SubnetId"`
	SubnetName  *string `json:"SubnetName,omitempty" name:"SubnetName"`
	Description *string `json:"Description,omitempty" name:"Description"`
	SoZoneId    *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *ModifySoSubnetAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySoSubnetAttributesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifySoSubnetAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoSubnetAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoSubnetRequest struct {
	*ksyunhttp.BaseRequest
	ZoneId     *string   `json:"ZoneId,omitempty" name:"ZoneId"`
	SubnetName *string   `json:"SubnetName,omitempty" name:"SubnetName"`
	VpcId      *string   `json:"VpcId,omitempty" name:"VpcId"`
	SubnetIds  []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	SoZoneId   *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoSubnetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Subnets   []struct {
		AvailableIpAddressCount *int    `json:"AvailableIpAddressCount" name:"AvailableIpAddressCount"`
		CidrBlock               *string `json:"CidrBlock" name:"CidrBlock"`
		Status                  *string `json:"Status" name:"Status"`
		SubnetId                *string `json:"SubnetId" name:"SubnetId"`
		SubnetName              *string `json:"SubnetName" name:"SubnetName"`
		TotalIpv4Count          *int    `json:"TotalIpv4Count" name:"TotalIpv4Count"`
		UpdateTime              *string `json:"UpdateTime" name:"UpdateTime"`
		VpcId                   *string `json:"VpcId" name:"VpcId"`
		ZoneId                  *string `json:"ZoneId" name:"ZoneId"`
	} `json:"Subnets"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeSoSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoKeyPairAttributeRequest struct {
	*ksyunhttp.BaseRequest
	Description *string `json:"Description,omitempty" name:"Description"`
	KeyPairId   *string `json:"KeyPairId,omitempty" name:"KeyPairId"`
	KeyPairName *string `json:"KeyPairName,omitempty" name:"KeyPairName"`
	SoZoneId    *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *ModifySoKeyPairAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySoKeyPairAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	KeyPairName *string `json:"KeyPairName" name:"KeyPairName"`
	Return      *bool   `json:"Return" name:"Return"`
}

func (r *ModifySoKeyPairAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoImageAttributeRequest struct {
	*ksyunhttp.BaseRequest
	BootMode    *string `json:"BootMode,omitempty" name:"BootMode"`
	Description *string `json:"Description,omitempty" name:"Description"`
	ImageId     *string `json:"ImageId,omitempty" name:"ImageId"`
	ImageName   *string `json:"ImageName,omitempty" name:"ImageName"`
	SoZoneId    *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *ModifySoImageAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySoImageAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifySoImageAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoVpcAttributesRequest struct {
	*ksyunhttp.BaseRequest
	VpcId       *string   `json:"VpcId,omitempty" name:"VpcId"`
	Description *string   `json:"Description,omitempty" name:"Description"`
	DnsServers  []*string `json:"DnsServers,omitempty" name:"DnsServers"`
	VpcName     *string   `json:"VpcName,omitempty" name:"VpcName"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *ModifySoVpcAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySoVpcAttributesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifySoVpcAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoVpcAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceSoSystemVolumeRequest struct {
	*ksyunhttp.BaseRequest
	ImageId             *string `json:"ImageId,omitempty" name:"ImageId"`
	InstanceId          *string `json:"InstanceId,omitempty" name:"InstanceId"`
	KeepImageCredential *bool   `json:"KeepImageCredential,omitempty" name:"KeepImageCredential"`
	KeyPairName         *string `json:"KeyPairName,omitempty" name:"KeyPairName"`
	Password            *string `json:"Password,omitempty" name:"Password"`
	SoZoneId            *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *ReplaceSoSystemVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReplaceSoSystemVolumeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ReplaceSoSystemVolumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceSoSystemVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoSubnetRequest struct {
	*ksyunhttp.BaseRequest
	VpcId       *string `json:"VpcId,omitempty" name:"VpcId"`
	ZoneId      *string `json:"ZoneId,omitempty" name:"ZoneId"`
	SubnetName  *string `json:"SubnetName,omitempty" name:"SubnetName"`
	Description *string `json:"Description,omitempty" name:"Description"`
	CidrBlock   *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	SoZoneId    *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *CreateSoSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSoSubnetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
	SubnetId  *string `json:"SubnetId" name:"SubnetId"`
}

func (r *CreateSoSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoVpcsRequest struct {
	*ksyunhttp.BaseRequest
	VpcName    *string   `json:"VpcName,omitempty" name:"VpcName"`
	VpcIds     []*string `json:"VpcIds,omitempty" name:"VpcIds"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	SoZoneId   *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DescribeSoVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoVpcsResponse struct {
	*ksyunhttp.BaseResponse
	Vpcs []struct {
		VpcId            *string   `json:"VpcId" name:"VpcId"`
		AttachVpcId      *string   `json:"AttachVpcId" name:"AttachVpcId"`
		AttachVpcName    *string   `json:"AttachVpcName" name:"AttachVpcName"`
		VpcName          *string   `json:"VpcName" name:"VpcName"`
		Description      *string   `json:"Description" name:"Description"`
		CidrBlock        *string   `json:"CidrBlock" name:"CidrBlock"`
		SubnetIds        []*string `json:"SubnetIds" name:"SubnetIds"`
		SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds"`
		DnsServers       []*string `json:"DnsServers" name:"DnsServers"`
		Status           *string   `json:"Status" name:"Status"`
		CreationTime     *string   `json:"CreationTime" name:"CreationTime"`
		UpdateTime       *string   `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Vpcs"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DescribeSoVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopSoInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ForceStop   *bool     `json:"ForceStop,omitempty" name:"ForceStop"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *StopSoInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StopSoInstanceResponse struct {
	*ksyunhttp.BaseResponse
	OperationDetails []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Error      struct {
			Code    *string `json:"Code" name:"Code"`
			Message *string `json:"Message" name:"Message"`
		} `json:"Error" name:"Error"`
	} `json:"OperationDetails"`
}

func (r *StopSoInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopSoInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSoKeyPairsRequest struct {
	*ksyunhttp.BaseRequest
	KeyPairNames []*string `json:"KeyPairNames,omitempty" name:"KeyPairNames"`
	SoZoneId     *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *DeleteSoKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSoKeyPairsResponse struct {
	*ksyunhttp.BaseResponse
	OperationDetails []struct {
		KeyPairName *string `json:"KeyPairName" name:"KeyPairName"`
		Error       struct {
			Code    *string `json:"Code" name:"Code"`
			Message *string `json:"Message" name:"Message"`
		} `json:"Error" name:"Error"`
	} `json:"OperationDetails"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteSoKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoImageRequest struct {
	*ksyunhttp.BaseRequest
	ForceStop   *bool     `json:"ForceStop,omitempty" name:"ForceStop"`
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	SoZoneId    *string   `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *CreateSoImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSoImageResponse struct {
	*ksyunhttp.BaseResponse
	ImageId   *string `json:"ImageId" name:"ImageId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *CreateSoImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoInstanceAttributeRequest struct {
	*ksyunhttp.BaseRequest
	DeletionProtection *string `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
	Description        *string `json:"Description,omitempty" name:"Description"`
	Hostname           *string `json:"Hostname,omitempty" name:"Hostname"`
	InstanceId         *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName       *string `json:"InstanceName,omitempty" name:"InstanceName"`
	Password           *string `json:"Password,omitempty" name:"Password"`
	SoZoneId           *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
	UserData           *string `json:"UserData,omitempty" name:"UserData"`
}

func (r *ModifySoInstanceAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySoInstanceAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifySoInstanceAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoKeyPairRequest struct {
	*ksyunhttp.BaseRequest
	KeyPairName *string `json:"KeyPairName,omitempty" name:"KeyPairName"`
	Description *string `json:"Description,omitempty" name:"Description"`
	SoZoneId    *string `json:"SoZoneId,omitempty" name:"SoZoneId"`
}

func (r *CreateSoKeyPairRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSoKeyPairResponse struct {
	*ksyunhttp.BaseResponse
	KeyPairName *string `json:"KeyPairName" name:"KeyPairName"`
	PrivateKey  *string `json:"PrivateKey" name:"PrivateKey"`
	KeyPairId   *string `json:"KeyPairId" name:"KeyPairId"`
	FingerPrint *string `json:"FingerPrint" name:"FingerPrint"`
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Return      *bool   `json:"Return" name:"Return"`
}

func (r *CreateSoKeyPairResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallAgentRequest struct {
	*ksyunhttp.BaseRequest
	HostId   *string `json:"HostId,omitempty" name:"HostId"`
	AgentId  *string `json:"AgentId,omitempty" name:"AgentId"`
	Username *string `json:"Username,omitempty" name:"Username"`
	Password *string `json:"Password,omitempty" name:"Password"`
	Key      *string `json:"Key,omitempty" name:"Key"`
}

func (r *InstallAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type InstallAgentResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *InstallAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentRequest struct {
	*ksyunhttp.BaseRequest
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`
	AgentId   *string `json:"AgentId,omitempty" name:"AgentId"`
	AgentType *string `json:"AgentType,omitempty" name:"AgentType"`
}

func (r *DescribeAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAgentResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	AgentInfo []struct {
		AgentName *string `json:"AgentName" name:"AgentName"`
		AgentId   *string `json:"AgentId" name:"AgentId"`
		AgentType *string `json:"AgentType" name:"AgentType"`
		Version   *string `json:"Version" name:"Version"`
	} `json:"AgentInfo"`
}

func (r *DescribeAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentInstallStatusRequest struct {
	*ksyunhttp.BaseRequest
	HostId     []*string `json:"HostId,omitempty" name:"HostId"`
	AgentId    *string   `json:"AgentId,omitempty" name:"AgentId"`
	Status     *string   `json:"Status,omitempty" name:"Status"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeAgentInstallStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAgentInstallStatusResponse struct {
	*ksyunhttp.BaseResponse
	AgentSet []struct {
		HostId     *string `json:"HostId" name:"HostId"`
		HostName   *string `json:"HostName" name:"HostName"`
		AgentId    *string `json:"AgentId" name:"AgentId"`
		Status     *string `json:"Status" name:"Status"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"AgentSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeAgentInstallStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInstallStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoUserDataRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSoUserDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSoUserDataResponse struct {
	*ksyunhttp.BaseResponse
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	UserData   *string `json:"UserData" name:"UserData"`
}

func (r *DescribeSoUserDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoUserDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDataRequest struct {
	*ksyunhttp.BaseRequest
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *DescribeUserDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeUserDataResponse struct {
	*ksyunhttp.BaseResponse
	HostId          *string `json:"HostId" name:"HostId"`
	UserDefinedData *string `json:"UserDefinedData" name:"UserDefinedData"`
}

func (r *DescribeUserDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
