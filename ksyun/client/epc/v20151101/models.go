package v20151101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type DescribeEpcsFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeInspectionsFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeEpcStocksFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeEpcDeviceAttributesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeProcessesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeEpcRaidAttributesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}


type CreateEpcRequest struct {
    *ksyunhttp.BaseRequest
    HostType *string `json:"HostType,omitempty" name:"HostType"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    Raid *string `json:"Raid,omitempty" name:"Raid"`
    RaidId *string `json:"RaidId,omitempty" name:"RaidId"`
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    NetworkInterfaceMode *string `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
    KeyId *string `json:"keyId,omitempty" name:"keyId"`
    SecurityGroupId []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    DNS1 *string `json:"DNS1,omitempty" name:"DNS1"`
    DNS2 *string `json:"DNS2,omitempty" name:"DNS2"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    Sn *string `json:"Sn,omitempty" name:"Sn"`
    PurchaseTime *int `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
    Password *string `json:"Password,omitempty" name:"Password"`
    CloudMonitorAgent *string `json:"CloudMonitorAgent,omitempty" name:"CloudMonitorAgent"`
    ExtensionSubnetId *string `json:"ExtensionSubnetId,omitempty" name:"ExtensionSubnetId"`
    ExtensionPrivateIpAddress *string `json:"ExtensionPrivateIpAddress,omitempty" name:"ExtensionPrivateIpAddress"`
    ExtensionDNS1 *string `json:"ExtensionDNS1,omitempty" name:"ExtensionDNS1"`
    ExtensionDNS2 *string `json:"ExtensionDNS2,omitempty" name:"ExtensionDNS2"`
    Description *string `json:"Description,omitempty" name:"Description"`
    AddressBandWidth *string `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
    LineId *string `json:"LineId,omitempty" name:"LineId"`
    BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
    AddressChargeType *string `json:"AddressChargeType,omitempty" name:"AddressChargeType"`
    AddressPurchaseTime *string `json:"AddressPurchaseTime,omitempty" name:"AddressPurchaseTime"`
    AddressProjectId *string `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
    SystemFileType *string `json:"SystemFileType,omitempty" name:"SystemFileType"`
    DataFileType *string `json:"DataFileType,omitempty" name:"DataFileType"`
    DataDiskCatalogue *string `json:"DataDiskCatalogue,omitempty" name:"DataDiskCatalogue"`
    DataDiskCatalogueSuffix *string `json:"DataDiskCatalogueSuffix,omitempty" name:"DataDiskCatalogueSuffix"`
    HyperThreading *string `json:"HyperThreading,omitempty" name:"HyperThreading"`
    NvmeDataFileType *string `json:"NvmeDataFileType,omitempty" name:"NvmeDataFileType"`
    NvmeDataDiskCatalogue *string `json:"NvmeDataDiskCatalogue,omitempty" name:"NvmeDataDiskCatalogue"`
    NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix,omitempty" name:"NvmeDataDiskCatalogueSuffix"`
    BondAttribute *string `json:"BondAttribute,omitempty" name:"BondAttribute"`
    ContainerAgent *string `json:"ContainerAgent,omitempty" name:"ContainerAgent"`
    KesAgent *string `json:"KesAgent,omitempty" name:"KesAgent"`
    KmrAgent *string `json:"KmrAgent,omitempty" name:"KmrAgent"`
    ComputerName *string `json:"ComputerName,omitempty" name:"ComputerName"`
    OverclockingAttribute *string `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
    GpuImageDriverId *string `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
    SystemVolumeType *string `json:"SystemVolumeType,omitempty" name:"SystemVolumeType"`
    SystemVolumeSize *string `json:"SystemVolumeSize,omitempty" name:"SystemVolumeSize"`
    RoceNetwork *string `json:"RoceNetwork,omitempty" name:"RoceNetwork"`
    ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
    ZoneType *string `json:"ZoneType,omitempty" name:"ZoneType"`
}

func (r *CreateEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	AutoReserveInfo struct {
		Reserve *bool `json:"Reserve"`
		HostId *string `json:"HostId"`
	} `json:"AutoReserveInfo"`
	Host struct {
		TorName *string `json:"TorName"`
		RackId *string `json:"RackId"`
		ContainerAgent *string `json:"ContainerAgent"`
		KesAgent *string `json:"KesAgent"`
		BondAttribute *string `json:"BondAttribute"`
		GpuImageDriverId *string `json:"GpuImageDriverId"`
		NvmeDataFileType *string `json:"NvmeDataFileType"`
		Memory *string `json:"Memory"`
		Raid *string `json:"Raid"`
		CloudMonitorAgent *string `json:"CloudMonitorAgent"`
		DataDiskCatalogue *string `json:"DataDiskCatalogue"`
		NetworkInterfaceMode *string `json:"NetworkInterfaceMode"`
		RaidTemplateId *string `json:"RaidTemplateId"`
		DataVolumeSet []struct {
					VolumeId *string `json:"VolumeId"`
					VolumeType *string `json:"VolumeType"`
					DeleteWithInstance *bool `json:"DeleteWithInstance"`
					VolumeSize *string `json:"VolumeSize"`
			} `json:"DataVolumeSet"`
			ImageId *string `json:"ImageId"`
			SystemVolumeType *string `json:"SystemVolumeType"`
			HostName *string `json:"HostName"`
			Tags *string `json:"Tags"`
			SystemFileType *string `json:"SystemFileType"`
			EnableBond *bool `json:"EnableBond"`
			ProductType *string `json:"ProductType"`
			AvailabilityZone *string `json:"AvailabilityZone"`
			HostId *string `json:"HostId"`
			NetworkInterfaceAttributeSet []struct {
						PrivateIpAddress *string `json:"PrivateIpAddress"`
						NetworkInterfaceType *string `json:"NetworkInterfaceType"`
					SecurityGroupSet []struct {
						SecurityGroupId *string `json:"SecurityGroupId"`
					} `json:"SecurityGroupSet"`
						DNS1 *string `json:"DNS1"`
						DNS2 *string `json:"DNS2"`
						SubnetId *string `json:"SubnetId"`
						NetworkInterfaceId *string `json:"NetworkInterfaceId"`
						Mac *string `json:"Mac"`
						VpcId *string `json:"VpcId"`
				} `json:"NetworkInterfaceAttributeSet"`
				ComputerName *string `json:"ComputerName"`
				CabinetId *string `json:"CabinetId"`
				DiskSet []struct {
							DiskType *string `json:"DiskType"`
							Space *string `json:"Space"`
							DiskCount *int `json:"DiskCount"`
							Raid *string `json:"Raid"`
							DiskAttribute *string `json:"DiskAttribute"`
							SystemDiskSpace *string `json:"SystemDiskSpace"`
							DiskSpace *string `json:"DiskSpace"`
					} `json:"DiskSet"`
					DataDiskCatalogueSuffix *string `json:"DataDiskCatalogueSuffix"`
					DataFileType *string `json:"DataFileType"`
					HostType *string `json:"HostType"`
					SystemVolumeSize *string `json:"SystemVolumeSize"`
					NvmeDataDiskCatalogue *string `json:"NvmeDataDiskCatalogue"`
					HostStatus *string `json:"HostStatus"`
					EnableContainer *bool `json:"EnableContainer"`
					ClusterId *string `json:"ClusterId"`
					HyperThreading *string `json:"HyperThreading"`
					CreateTime *string `json:"CreateTime"`
					OsName *string `json:"OsName"`
					CabinetName *string `json:"CabinetName"`
					ProjectId *string `json:"ProjectId"`
					KeyId *string `json:"KeyId"`
					AllowModifyHyperThreading *bool `json:"AllowModifyHyperThreading"`
					ReleasableTime *string `json:"ReleasableTime"`
					RackName *string `json:"RackName"`
					KmrAgent *string `json:"KmrAgent"`
					Sn *string `json:"Sn"`
					NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix"`
					SecurityAgent *string `json:"SecurityAgent"`
					SupportEbs *string `json:"SupportEbs"`
					KplAgent *string `json:"KplAgent"`
					ServiceEndTime *string `json:"ServiceEndTime"`
					ChargeType *string `json:"ChargeType"`
					Cpu struct {
							Model *string `json:"Model"`
							Frequence *string `json:"Frequence"`
							Count *int `json:"Count"`
							CoreCount *int `json:"CoreCount"`
					} `json:"Cpu"`
					Gpu struct {
							Model *string `json:"Model"`
							Frequence *string `json:"Frequence"`
							Count *int `json:"Count"`
							CoreCount *int `json:"CoreCount"`
							GpuCount *int `json:"GpuCount"`
					} `json:"Gpu"`
					Roces []struct {
								Ip *string `json:"Ip"`
								Mask *string `json:"Mask"`
								GateWay *string `json:"GateWay"`
								Type *string `json:"Type"`
						} `json:"Roces"`
						ContractDueTime *string `json:"ContractDueTime"`
						AutoDeleteTime *string `json:"AutoDeleteTime"`
						VpcTrust *int `json:"VpcTrust"`
					} `json:"Host"`
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

func (r *StartEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StartEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StartEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *RebootEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RebootEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RebootEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DeleteEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    KeyId *string `json:"keyId,omitempty" name:"keyId"`
    Password *string `json:"Password,omitempty" name:"Password"`
    NetworkInterfaceMode *string `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
    CloudMonitorAgent *string `json:"CloudMonitorAgent,omitempty" name:"CloudMonitorAgent"`
    Raid *string `json:"Raid,omitempty" name:"Raid"`
    RaidId *string `json:"RaidId,omitempty" name:"RaidId"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    SystemFileType *string `json:"SystemFileType,omitempty" name:"SystemFileType"`
    DataFileType *string `json:"DataFileType,omitempty" name:"DataFileType"`
    DataDiskCatalogue *string `json:"DataDiskCatalogue,omitempty" name:"DataDiskCatalogue"`
    DataDiskCatalogueSuffix *string `json:"DataDiskCatalogueSuffix,omitempty" name:"DataDiskCatalogueSuffix"`
    HyperThreading *string `json:"HyperThreading,omitempty" name:"HyperThreading"`
    NvmeDataFileType *string `json:"NvmeDataFileType,omitempty" name:"NvmeDataFileType"`
    NvmeDataDiskCatalogue *string `json:"NvmeDataDiskCatalogue,omitempty" name:"NvmeDataDiskCatalogue"`
    NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix,omitempty" name:"NvmeDataDiskCatalogueSuffix"`
    BondAttribute *string `json:"BondAttribute,omitempty" name:"BondAttribute"`
    KesAgent *string `json:"KesAgent,omitempty" name:"KesAgent"`
    KmrAgent *string `json:"KmrAgent,omitempty" name:"KmrAgent"`
    ComputerName *string `json:"ComputerName,omitempty" name:"ComputerName"`
    OverclockingAttribute *string `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
    DelayStart *int `json:"DelayStart,omitempty" name:"DelayStart"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    GpuImageDriverId *string `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
    ContainerAgent *string `json:"ContainerAgent,omitempty" name:"ContainerAgent"`
}

func (r *ReinstallEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstallEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReinstallEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ReinstallEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
    SecurityGroupId []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *ModifySecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateKeyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    PrivateKey *string `json:"PrivateKey" name:"PrivateKey"`
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
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    HostId []*string `json:"HostId,omitempty" name:"HostId"`
    Filter []*DescribeEpcsFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcsResponse struct {
    *ksyunhttp.BaseResponse
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	HostSet []struct {
		TorName *string `json:"TorName"`
		KesAgent *string `json:"KesAgent"`
		BondAttribute *string `json:"BondAttribute"`
		GpuImageDriverId *string `json:"GpuImageDriverId"`
		NvmeDataFileType *string `json:"NvmeDataFileType"`
		Memory *string `json:"Memory"`
		Cpu struct {
				Model *string `json:"Model"`
				Count *int `json:"Count"`
				Frequence *string `json:"Frequence"`
				CoreCount *int `json:"CoreCount"`
		} `json:"Cpu"`
		Raid *string `json:"Raid"`
		Gpu struct {
				Model *string `json:"Model"`
				GpuCount *int `json:"GpuCount"`
				Count *int `json:"Count"`
				Frequence *string `json:"Frequence"`
				CoreCount *int `json:"CoreCount"`
		} `json:"Gpu"`
		CloudMonitorAgent *string `json:"CloudMonitorAgent"`
		DataDiskCatalogue *string `json:"DataDiskCatalogue"`
		NetworkInterfaceMode *string `json:"NetworkInterfaceMode"`
		RaidTemplateId *string `json:"RaidTemplateId"`
		DataVolumeSet []struct {
					VolumeId *string `json:"VolumeId"`
					VolumeType *string `json:"VolumeType"`
					DeleteWithInstance *bool `json:"DeleteWithInstance"`
					VolumeSize *string `json:"VolumeSize"`
			} `json:"DataVolumeSet"`
			ImageId *string `json:"ImageId"`
			SystemVolumeType *string `json:"SystemVolumeType"`
			HostName *string `json:"HostName"`
			SystemFileType *string `json:"SystemFileType"`
			EnableBond *bool `json:"EnableBond"`
			ProductType *string `json:"ProductType"`
			AvailabilityZone *string `json:"AvailabilityZone"`
			HostId *string `json:"HostId"`
			NetworkInterfaceAttributeSet []struct {
						PrivateIpAddress *string `json:"PrivateIpAddress"`
						NetworkInterfaceType *string `json:"NetworkInterfaceType"`
					SecurityGroupSet []struct {
						SecurityGroupId *string `json:"SecurityGroupId"`
					} `json:"SecurityGroupSet"`
						DNS1 *string `json:"DNS1"`
						DNS2 *string `json:"DNS2"`
						SubnetId *string `json:"SubnetId"`
						NetworkInterfaceId *string `json:"NetworkInterfaceId"`
						Mac *string `json:"Mac"`
						VpcId *string `json:"VpcId"`
				} `json:"NetworkInterfaceAttributeSet"`
				ComputerName *string `json:"ComputerName"`
				CabinetId *string `json:"CabinetId"`
				DiskSet []struct {
							DiskType *string `json:"DiskType"`
							Space *string `json:"Space"`
							DiskCount *int `json:"DiskCount"`
							Raid *string `json:"Raid"`
							DiskAttribute *string `json:"DiskAttribute"`
							SystemDiskSpace *string `json:"SystemDiskSpace"`
							DiskSpace *string `json:"DiskSpace"`
					} `json:"DiskSet"`
					DataDiskCatalogueSuffix *string `json:"DataDiskCatalogueSuffix"`
					DataFileType *string `json:"DataFileType"`
					HostType *string `json:"HostType"`
					SystemVolumeSize *string `json:"SystemVolumeSize"`
					NvmeDataDiskCatalogue *string `json:"NvmeDataDiskCatalogue"`
					HostStatus *string `json:"HostStatus"`
					EnableContainer *bool `json:"EnableContainer"`
					ClusterId *string `json:"ClusterId"`
					HyperThreading *string `json:"HyperThreading"`
					CreateTime *string `json:"CreateTime"`
					OsName *string `json:"OsName"`
					CabinetName *string `json:"CabinetName"`
					ProjectId *string `json:"ProjectId"`
					KeyId *string `json:"KeyId"`
					AllowModifyHyperThreading *bool `json:"AllowModifyHyperThreading"`
					ReleasableTime *string `json:"ReleasableTime"`
					RackName *string `json:"RackName"`
					KmrAgent *string `json:"KmrAgent"`
					Sn *string `json:"Sn"`
					NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix"`
					SecurityAgent *string `json:"SecurityAgent"`
					SupportEbs *string `json:"SupportEbs"`
					KplAgent *string `json:"KplAgent"`
					RackId *string `json:"RackId"`
					ContainerAgent *string `json:"ContainerAgent"`
					Roces []struct {
								Ip *string `json:"Ip"`
								Mask *string `json:"Mask"`
								GateWay *string `json:"GateWay"`
								Type *string `json:"Type"`
						} `json:"Roces"`
						ContractDueTime *string `json:"ContractDueTime"`
						AutoDeleteTime *string `json:"AutoDeleteTime"`
						VpcTrust *int `json:"VpcTrust"`
					} `json:"HostSet"`
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

func (r *GetDynamicCodeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDynamicCodeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetDynamicCodeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    DynamicCode *string `json:"DynamicCode,omitempty" name:"DynamicCode"`
    Pin *string `json:"Pin,omitempty" name:"Pin"`
    RemoteManagementId *string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
}

func (r *DescribeVpnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpnsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    UserName *string `json:"UserName" name:"UserName"`
    Password *string `json:"Password" name:"Password"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
    ImageMode *string `json:"ImageMode,omitempty" name:"ImageMode"`
    ImageInitialization *string `json:"ImageInitialization,omitempty" name:"ImageInitialization"`
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
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *ModifyImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DeleteImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	ImageSet []struct {
		ImageId *string `json:"ImageId"`
		ImageName *string `json:"ImageName"`
		ImageType *string `json:"ImageType"`
		OsName *string `json:"OsName"`
		OsType *string `json:"OsType"`
		EnableContainer *bool `json:"EnableContainer"`
		CreateTime *string `json:"CreateTime"`
		DiskType *string `json:"DiskType"`
		EbsImageSize *string `json:"EbsImageSize"`
		Status *string `json:"Status"`
		ImageInitialization *string `json:"ImageInitialization"`
		Name *string `json:"Name"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    DNS1 *string `json:"DNS1,omitempty" name:"DNS1"`
    DNS2 *string `json:"DNS2,omitempty" name:"DNS2"`
}

func (r *ModifyDnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDnsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDnsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
    SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`
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
    Return *bool `json:"Return" name:"Return"`
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

func (r *DescribePhysicalMonitorRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePhysicalMonitorRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribePhysicalMonitorResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	CPUInfo []struct {
		Item *string `json:"Item"`
		Status *string `json:"Status"`
	} `json:"CPUInfo"`
	MemoryInfo []struct {
		Item *string `json:"Item"`
		Status *string `json:"Status"`
	} `json:"MemoryInfo"`
	DiskInfo []struct {
		Item *string `json:"Item"`
		Status *string `json:"Status"`
	} `json:"DiskInfo"`
	FanInfo []struct {
		Item *string `json:"Item"`
		Status *string `json:"Status"`
	} `json:"FanInfo"`
	PowerInfo []struct {
		Item *string `json:"Item"`
		Status *string `json:"Status"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
    DynamicCode *string `json:"DynamicCode,omitempty" name:"DynamicCode"`
    Pin *string `json:"Pin,omitempty" name:"Pin"`
    EpcManagementId []*string `json:"EpcManagementId,omitempty" name:"EpcManagementId"`
    RemoteManagementId *string `json:"RemoteManagementId,omitempty" name:"RemoteManagementId"`
}

func (r *DescribeEpcManagementsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcManagementsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcManagementsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcManagementsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	EpcManagementSet []struct {
		EpcManagementId *string `json:"EpcManagementId"`
		EpcManagementIp *string `json:"EpcManagementIp"`
		EpcManagementUserName *string `json:"EpcManagementUserName"`
		Password *string `json:"Password"`
		HostName *string `json:"HostName"`
		Sn *string `json:"Sn"`
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

func (r *DescribeRemoteManagementsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRemoteManagementsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteManagementsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	RemoteManagementSet []struct {
		RemoteManagementId *string `json:"RemoteManagementId"`
		PhoneNumber *string `json:"PhoneNumber"`
		Name *string `json:"Name"`
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

func (r *StopEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StopEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StopEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    DynamicCode *string `json:"DynamicCode,omitempty" name:"DynamicCode"`
    Pin *string `json:"Pin,omitempty" name:"Pin"`
    NewPhoneNumber *string `json:"NewPhoneNumber,omitempty" name:"NewPhoneNumber"`
    NewPin *string `json:"NewPin,omitempty" name:"NewPin"`
    Name *string `json:"Name,omitempty" name:"Name"`
    VersionId *int `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *ModifyRemoteManagementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRemoteManagementRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyRemoteManagementRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteManagementResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Pin *string `json:"Pin,omitempty" name:"Pin"`
    PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
    Name *string `json:"Name,omitempty" name:"Name"`
    VersionId *int `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *CreateRemoteManagementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRemoteManagementRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRemoteManagementRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteManagementResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`
    Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *ReinstallCustomerEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstallCustomerEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReinstallCustomerEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ReinstallCustomerEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DeleteRemoteManagementRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRemoteManagementRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteManagementResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    HyperThreadingStatus *string `json:"HyperThreadingStatus,omitempty" name:"HyperThreadingStatus"`
}

func (r *ModifyHyperThreadingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHyperThreadingRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyHyperThreadingRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHyperThreadingResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AssociateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateClusterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateClusterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AssociateClusterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DisassociateClusterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateClusterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateClusterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
    Filter []*DescribeInspectionsFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeInspectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInspectionsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInspectionsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInspectionsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	InspectionSet []struct {
		HostId *string `json:"HostId"`
		Sn *string `json:"Sn"`
		Region *string `json:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		Status *string `json:"Status"`
		AlarmType *string `json:"AlarmType"`
		AlarmDescription *string `json:"AlarmDescription"`
		CreateTime *string `json:"CreateTime"`
		TorName *string `json:"TorName"`
		CabinetName *string `json:"CabinetName"`
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

func (r *DescribeEpcStocksRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcStocksRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcStocksResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	EpcStockSet []struct {
		InstanceCount *int `json:"InstanceCount"`
		HostType *string `json:"HostType"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		AvailableRaidLevelSet []struct {
			} `json:"AvailableRaidLevelSet"`
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
    Filter []*DescribeEpcDeviceAttributesFilter `json:"Filter,omitempty" name:"Filter"`
    DeviceAttributeId []*string `json:"DeviceAttributeId,omitempty" name:"DeviceAttributeId"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcDeviceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcDeviceAttributesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcDeviceAttributesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcDeviceAttributesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	EpcDeviceAttributeSet []struct {
		DeviceAttributeId *string `json:"DeviceAttributeId"`
		HostType *string `json:"HostType"`
		HostTypeName *string `json:"HostTypeName"`
		Memory *string `json:"Memory"`
		NetworkCard *string `json:"NetworkCard"`
		CpuDeviceSet []struct {
					CpuSpec *string `json:"CpuSpec"`
			} `json:"CpuDeviceSet"`
			GpuDeviceSet []struct {
						GpuModel *string `json:"GpuModel"`
						GpuCount *string `json:"GpuCount"`
				} `json:"GpuDeviceSet"`
				PhysicalDiskDeviceSet []struct {
							DiskAttribute *string `json:"DiskAttribute"`
							DiskCount *string `json:"DiskCount"`
							Space *string `json:"Space"`
					} `json:"PhysicalDiskDeviceSet"`
					PriceSet []struct {
								MonthlylistPrice *string `json:"MonthlylistPrice"`
						} `json:"PriceSet"`
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
    OperationProcessId []*string `json:"OperationProcessId,omitempty" name:"OperationProcessId"`
    Filter []*DescribeProcessesFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeProcessesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	ProcessSet []struct {
		HostId *string `json:"HostId"`
		Sn *string `json:"Sn"`
		Region *string `json:"Region"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		Status *string `json:"Status"`
		Attribute *string `json:"Attribute"`
		ProcessId *string `json:"ProcessId"`
		OperationProcessId *string `json:"OperationProcessId"`
		CreateTime *string `json:"CreateTime"`
		FinishTime *string `json:"FinishTime"`
		Source *string `json:"Source"`
		Content *string `json:"Content"`
		MachineCount *int `json:"MachineCount"`
		Title *string `json:"Title"`
		Type *string `json:"Type"`
		Confirm *string `json:"Confirm"`
		HostTypeName *string `json:"HostTypeName"`
		HostName *string `json:"HostName"`
		CommunicationContentSet []struct {
					Remarks *string `json:"Remarks"`
					Author *string `json:"Author"`
					CreateTime *string `json:"CreateTime"`
			} `json:"CommunicationContentSet"`
			HostType *string `json:"HostType"`
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
    ProcessId *string `json:"ProcessId,omitempty" name:"ProcessId"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Sn *string `json:"Sn,omitempty" name:"Sn"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
    Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
    Content *string `json:"Content,omitempty" name:"Content"`
    MachineCount *int `json:"MachineCount,omitempty" name:"MachineCount"`
    Title *string `json:"Title,omitempty" name:"Title"`
    Type *string `json:"Type,omitempty" name:"Type"`
    Confirm *string `json:"Confirm,omitempty" name:"Confirm"`
    ProcessSource *int `json:"ProcessSource,omitempty" name:"ProcessSource"`
    AutoNocCase *int `json:"autoNocCase,omitempty" name:"autoNocCase"`
}

func (r *CreateProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcessRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateProcessRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcessResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DeleteProcessRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteProcessRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProcessResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
}

func (r *ReplyProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplyProcessRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReplyProcessRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ReplyProcessResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcTrashesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcTrashesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcTrashesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcTrashesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	HostSet []struct {
		CreateTime *string `json:"CreateTime"`
		ComputerName *string `json:"ComputerName"`
		HostId *string `json:"HostId"`
		HostName *string `json:"HostName"`
		HostType *string `json:"HostType"`
		AllowModifyHyperThreading *bool `json:"AllowModifyHyperThreading"`
		ReleasableTime *string `json:"ReleasableTime"`
		TorName *string `json:"TorName"`
		CabinetName *string `json:"CabinetName"`
		RackName *string `json:"RackName"`
		Sn *string `json:"Sn"`
		CabinetId *string `json:"CabinetId"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		Raid *string `json:"Raid"`
		RaidTemplateId *string `json:"RaidTemplateId"`
		ImageId *string `json:"ImageId"`
		KeyId *string `json:"KeyId"`
		NetworkInterfaceMode *string `json:"NetworkInterfaceMode"`
		BondAttribute *string `json:"BondAttribute"`
		EnableBond *bool `json:"EnableBond"`
		SecurityAgent *string `json:"SecurityAgent"`
		CloudMonitorAgent *string `json:"CloudMonitorAgent"`
		SupportEbs *string `json:"SupportEbs"`
		ProductType *string `json:"ProductType"`
		OsName *string `json:"OsName"`
		Memory *string `json:"Memory"`
		HostStatus *string `json:"HostStatus"`
		ClusterId *string `json:"ClusterId"`
		EnableContainer *bool `json:"EnableContainer"`
		ProjectId *string `json:"ProjectId"`
		SystemFileType *string `json:"SystemFileType"`
		DataFileType *string `json:"DataFileType"`
		DataDiskCatalogue *string `json:"DataDiskCatalogue"`
		DataDiskCatalogueSuffix *string `json:"DataDiskCatalogueSuffix"`
		NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix"`
		NvmeDataDiskCatalogue *string `json:"NvmeDataDiskCatalogue"`
		NvmeDataFileType *string `json:"NvmeDataFileType"`
		KesAgent *string `json:"KesAgent"`
		KplAgent *string `json:"KplAgent"`
		KmrAgent *string `json:"KmrAgent"`
		DiskSet []struct {
					DiskType *string `json:"DiskType"`
					SystemDiskSpace *string `json:"SystemDiskSpace"`
					Raid *string `json:"Raid"`
					DiskAttribute *string `json:"DiskAttribute"`
					Space *string `json:"Space"`
					DiskCount *int `json:"DiskCount"`
					DiskSpace *string `json:"DiskSpace"`
			} `json:"DiskSet"`
			NetworkInterfaceAttributeSet []struct {
						NetworkInterfaceId *string `json:"NetworkInterfaceId"`
						NetworkInterfaceType *string `json:"NetworkInterfaceType"`
						SubnetId *string `json:"SubnetId"`
						PrivateIpAddress *string `json:"PrivateIpAddress"`
						DNS1 *string `json:"DNS1"`
						DNS2 *string `json:"DNS2"`
						Mac *string `json:"Mac"`
					SecurityGroupSet []struct {
						SecurityGroupId *string `json:"SecurityGroupId"`
					} `json:"SecurityGroupSet"`
						VpcId *string `json:"VpcId"`
				} `json:"NetworkInterfaceAttributeSet"`
				SystemVolumeType *string `json:"SystemVolumeType"`
				SystemVolumeSize *string `json:"SystemVolumeSize"`
				DataVolumeSet []struct {
							VolumeId *string `json:"VolumeId"`
							VolumeType *string `json:"VolumeType"`
							VolumeSize *string `json:"VolumeSize"`
							DeleteWithInstance *bool `json:"DeleteWithInstance"`
					} `json:"DataVolumeSet"`
					GpuImageDriverId *string `json:"GpuImageDriverId"`
					Tags *string `json:"Tags"`
					HyperThreading *string `json:"HyperThreading"`
					RackId *string `json:"RackId"`
					ContainerAgent *string `json:"ContainerAgent"`
					ServiceEndTime *string `json:"ServiceEndTime"`
					ChargeType *string `json:"ChargeType"`
					Cpu struct {
							Model *string `json:"Model"`
							Frequence *string `json:"Frequence"`
							Count *int `json:"Count"`
							CoreCount *int `json:"CoreCount"`
					} `json:"Cpu"`
					Gpu struct {
							Model *string `json:"Model"`
							Frequence *string `json:"Frequence"`
							Count *int `json:"Count"`
							CoreCount *int `json:"CoreCount"`
							GpuCount *int `json:"GpuCount"`
					} `json:"Gpu"`
					Roces []struct {
								Ip *string `json:"Ip"`
								Mask *string `json:"Mask"`
								GateWay *string `json:"GateWay"`
								Type *string `json:"Type"`
						} `json:"Roces"`
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

func (r *ReturnEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReturnEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ReturnEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    RequirementCount *int `json:"RequirementCount,omitempty" name:"RequirementCount"`
    ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
    UsageDate *string `json:"UsageDate,omitempty" name:"UsageDate"`
    Description *string `json:"Description,omitempty" name:"Description"`
    HostType *string `json:"HostType,omitempty" name:"HostType"`
}

func (r *CreateResourceRequirementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateResourceRequirementRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateResourceRequirementRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateResourceRequirementResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *AttachVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AttachVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    VolumeId *string `json:"VolumeId,omitempty" name:"VolumeId"`
}

func (r *DetachVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachVolumeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachVolumeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DetachVolumeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    HostType *string `json:"HostType,omitempty" name:"HostType"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    PurchaseTime *int `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
    Amount *int `json:"Amount,omitempty" name:"Amount"`
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
}

func (r *DescribePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyOverclockingAttributeRequest struct {
    *ksyunhttp.BaseRequest
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    OverclockingAttribute *string `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
}

func (r *ModifyOverclockingAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyOverclockingAttributeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyOverclockingAttributeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyOverclockingAttributeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    DestinationName *string `json:"DestinationName,omitempty" name:"DestinationName"`
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
    CopyTag *string `json:"CopyTag,omitempty" name:"CopyTag"`
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
    Return *bool `json:"Return" name:"Return"`
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
    Filter []*DescribeEpcRaidAttributesFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeEpcRaidAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEpcRaidAttributesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEpcRaidAttributesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEpcRaidAttributesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	EpcRaidAttributeSet []struct {
		TemplateName *string `json:"TemplateName"`
		CreateTime *string `json:"CreateTime"`
		RaidId *string `json:"RaidId"`
		HostType *string `json:"HostType"`
		DiskSet []struct {
					DiskAttribute *string `json:"DiskAttribute"`
					DiskCount *string `json:"DiskCount"`
					Space *string `json:"Space"`
					SystemDiskSpace *string `json:"SystemDiskSpace"`
					DiskSpace *string `json:"DiskSpace"`
					Raid *string `json:"Raid"`
					DiskType *string `json:"DiskType"`
					DiskId *string `json:"DiskId"`
			} `json:"DiskSet"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    HostType *string `json:"HostType,omitempty" name:"HostType"`
}

func (r *DescribeGpuImageDriverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGpuImageDriverRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeGpuImageDriverRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGpuImageDriverResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	GpuImagesDriverSet []struct {
		GpuImageDriverId *string `json:"GpuImageDriverId"`
		ImageNameSet []struct {
			} `json:"ImageNameSet"`
			GpuModel []struct {
				} `json:"GpuModel"`
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
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    AccountIdN *string `json:"AccountId.N,omitempty" name:"AccountId.N"`
}

func (r *CreateShareImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateShareImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateShareImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateShareImageResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    AccountIdN *string `json:"AccountId.N,omitempty" name:"AccountId.N"`
}

func (r *DeleteShareImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteShareImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteShareImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteShareImageResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DescribeShareImageAccountListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeShareImageAccountListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShareImageAccountListResponse struct {
    *ksyunhttp.BaseResponse
	SharePermissionSet []struct {
		AccountId *int `json:"AccountId"`
		ShareTime *string `json:"ShareTime"`
		Status *string `json:"Status"`
		ImageId *string `json:"ImageId"`
	} `json:"SharePermissionSet"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeShareImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShareImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeShareImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShareImageResponse struct {
    *ksyunhttp.BaseResponse
	SharePermissionSet []struct {
		ImageName *string `json:"ImageName"`
		System *string `json:"System"`
		CreateTime *string `json:"CreateTime"`
		FromId *string `json:"FromId"`
		ImageId *string `json:"ImageId"`
		ShareTime *string `json:"ShareTime"`
		ImageInitialization *string `json:"ImageInitialization"`
		Status *string `json:"Status"`
	} `json:"SharePermissionSet"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
    NextToken *string `json:"NextToken" name:"NextToken"`
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

func (r *AcceptShareImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AcceptShareImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AcceptShareImageResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *RejectShareImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RejectShareImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RejectShareImageResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *DescribeManagedAccessoryRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeManagedAccessoryRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeManagedAccessoryResponse struct {
    *ksyunhttp.BaseResponse
	ManagedAccessorySet []struct {
		SN *string `json:"SN"`
		IDC *string `json:"IDC"`
		Classification *string `json:"Classification"`
		Model *string `json:"Model"`
		Manufacturer *string `json:"Manufacturer"`
		State *string `json:"State"`
		Date *string `json:"Date"`
		Source *string `json:"Source"`
		Notes *string `json:"Notes"`
		Num *int `json:"Num"`
		AlarmNum *int `json:"AlarmNum"`
	} `json:"ManagedAccessorySet"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    AutoDeleteTime *string `json:"AutoDeleteTime,omitempty" name:"AutoDeleteTime"`
    AutoDeleteEip *string `json:"AutoDeleteEip,omitempty" name:"AutoDeleteEip"`
}

func (r *AutoDeleteEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoDeleteEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AutoDeleteEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AutoDeleteEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    Ks3Bucket *string `json:"Ks3Bucket,omitempty" name:"Ks3Bucket"`
    ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
}

func (r *ExportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportImageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ExportImageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ExportImageResponse struct {
    *ksyunhttp.BaseResponse
    Return *bool `json:"Return" name:"Return"`
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

func (r *QueryBucketsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryBucketsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type QueryBucketsResponse struct {
    *ksyunhttp.BaseResponse
	BucketSet []struct {
		BucketName *string `json:"BucketName"`
		BucketHost *string `json:"BucketHost"`
		BucketHostCompatible *string `json:"BucketHostCompatible"`
		BucketInnerHost *string `json:"BucketInnerHost"`
		Endpoint *string `json:"Endpoint"`
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

func (r *CancelImageExportRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelImageExportRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CancelImageExportResponse struct {
    *ksyunhttp.BaseResponse
    Return *bool `json:"Return" name:"Return"`
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
    HostId *string `json:"HostId,omitempty" name:"HostId"`
    HotStandByHostId *string `json:"HotStandByHostId,omitempty" name:"HotStandByHostId"`
}

func (r *UseHotStandByEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UseHotStandByEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UseHotStandByEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UseHotStandByEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
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

func (r *ActivateHotStandbyEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ActivateHotStandbyEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ActivateHotStandbyEpcResponse struct {
    *ksyunhttp.BaseResponse
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
    HostType *string `json:"HostType,omitempty" name:"HostType"`
    AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
    Raid *string `json:"Raid,omitempty" name:"Raid"`
    RaidId *string `json:"RaidId,omitempty" name:"RaidId"`
    ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
    NetworkInterfaceMode *string `json:"NetworkInterfaceMode,omitempty" name:"NetworkInterfaceMode"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    KeyId *string `json:"keyId,omitempty" name:"keyId"`
    SecurityGroupId []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    DNS1 *string `json:"DNS1,omitempty" name:"DNS1"`
    DNS2 *string `json:"DNS2,omitempty" name:"DNS2"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    Sn *string `json:"Sn,omitempty" name:"Sn"`
    PurchaseTime *int `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
    Password *string `json:"Password,omitempty" name:"Password"`
    CloudMonitorAgent *string `json:"CloudMonitorAgent,omitempty" name:"CloudMonitorAgent"`
    ExtensionSubnetId *string `json:"ExtensionSubnetId,omitempty" name:"ExtensionSubnetId"`
    ExtensionDNS1 *string `json:"ExtensionDNS1,omitempty" name:"ExtensionDNS1"`
    ExtensionDNS2 *string `json:"ExtensionDNS2,omitempty" name:"ExtensionDNS2"`
    Description *string `json:"Description,omitempty" name:"Description"`
    AddressBandWidth *string `json:"AddressBandWidth,omitempty" name:"AddressBandWidth"`
    LineId *string `json:"LineId,omitempty" name:"LineId"`
    BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
    AddressChargeType *string `json:"AddressChargeType,omitempty" name:"AddressChargeType"`
    AddressPurchaseTime *string `json:"AddressPurchaseTime,omitempty" name:"AddressPurchaseTime"`
    AddressProjectId *string `json:"AddressProjectId,omitempty" name:"AddressProjectId"`
    SystemFileType *string `json:"SystemFileType,omitempty" name:"SystemFileType"`
    DataFileType *string `json:"DataFileType,omitempty" name:"DataFileType"`
    DataDiskCatalogue *string `json:"DataDiskCatalogue,omitempty" name:"DataDiskCatalogue"`
    DataDiskCatalogueSuffix *string `json:"DataDiskCatalogueSuffix,omitempty" name:"DataDiskCatalogueSuffix"`
    HyperThreading *string `json:"HyperThreading,omitempty" name:"HyperThreading"`
    NvmeDataFileType *string `json:"NvmeDataFileType,omitempty" name:"NvmeDataFileType"`
    NvmeDataDiskCatalogue *string `json:"NvmeDataDiskCatalogue,omitempty" name:"NvmeDataDiskCatalogue"`
    NvmeDataDiskCatalogueSuffix *string `json:"NvmeDataDiskCatalogueSuffix,omitempty" name:"NvmeDataDiskCatalogueSuffix"`
    BondAttribute *string `json:"BondAttribute,omitempty" name:"BondAttribute"`
    ContainerAgent *string `json:"ContainerAgent,omitempty" name:"ContainerAgent"`
    KesAgent *string `json:"KesAgent,omitempty" name:"KesAgent"`
    KmrAgent *string `json:"KmrAgent,omitempty" name:"KmrAgent"`
    ComputerName *string `json:"ComputerName,omitempty" name:"ComputerName"`
    OverclockingAttribute *string `json:"OverclockingAttribute,omitempty" name:"OverclockingAttribute"`
    GpuImageDriverId *string `json:"GpuImageDriverId,omitempty" name:"GpuImageDriverId"`
    SystemVolumeType *string `json:"SystemVolumeType,omitempty" name:"SystemVolumeType"`
    SystemVolumeSize *string `json:"SystemVolumeSize,omitempty" name:"SystemVolumeSize"`
    RoceNetwork *string `json:"RoceNetwork,omitempty" name:"RoceNetwork"`
    ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
    ZoneType *string `json:"ZoneType,omitempty" name:"ZoneType"`
    HostNameStartNo *int `json:"HostNameStartNo,omitempty" name:"HostNameStartNo"`
    ComputerNameStartNo *int `json:"ComputerNameStartNo,omitempty" name:"ComputerNameStartNo"`
    Amount *int `json:"Amount,omitempty" name:"Amount"`
}

func (r *BatchCreateEpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchCreateEpcRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BatchCreateEpcRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type BatchCreateEpcResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Host []struct {
		HostId *string `json:"HostId"`
		HostName *string `json:"HostName"`
	} `json:"Host"`
}

func (r *BatchCreateEpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchCreateEpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

