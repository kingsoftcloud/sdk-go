package v20191017
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId        *string `json:"ProjectId,omitempty" name:"ProjectId"`
	InstanceName     *string `json:"InstanceName,omitempty" name:"InstanceName"`
	InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
	VpcId            *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId         *string `json:"SubnetId,omitempty" name:"SubnetId"`
	EngineVersion    *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	BillType         *int    `json:"BillType,omitempty" name:"BillType"`
	Duration         *int    `json:"Duration,omitempty" name:"Duration"`
	Mode             *string `json:"Mode,omitempty" name:"Mode"`
	InstanceType     *string `json:"InstanceType,omitempty" name:"InstanceType"`
	SsdDisk          *int    `json:"SsdDisk,omitempty" name:"SsdDisk"`
	NodeNum          *int    `json:"NodeNum,omitempty" name:"NodeNum"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		SsdDisk    *int    `json:"SsdDisk" name:"SsdDisk"`
		SubOrderId *string `json:"SubOrderId" name:"SubOrderId"`
	} `json:"Data"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	Offset       *int    `json:"Offset,omitempty" name:"Offset"`
	Limit        *int    `json:"Limit,omitempty" name:"Limit"`
	OrderBy      *string `json:"OrderBy,omitempty" name:"OrderBy"`
	ProjectId    *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Offset    *int `json:"Offset" name:"Offset"`
		Limit     *int `json:"Limit" name:"Limit"`
		Total     *int `json:"Total" name:"Total"`
		Instances []struct {
			UserId           *string `json:"UserId" name:"UserId"`
			Region           *string `json:"Region" name:"Region"`
			InstanceName     *string `json:"InstanceName" name:"InstanceName"`
			InstanceId       *string `json:"InstanceId" name:"InstanceId"`
			StatusName       *string `json:"StatusName" name:"StatusName"`
			ExpirationDate   *string `json:"ExpirationDate" name:"ExpirationDate"`
			Status           *string `json:"Status" name:"Status"`
			Vip              *string `json:"Vip" name:"Vip"`
			WebVip           *string `json:"WebVip" name:"WebVip"`
			InstanceType     *string `json:"InstanceType" name:"InstanceType"`
			SsdDisk          *int    `json:"SsdDisk" name:"SsdDisk"`
			Protocol         *string `json:"Protocol" name:"Protocol"`
			SecurityGroupId  *int    `json:"SecurityGroupId" name:"SecurityGroupId"`
			Port             *string `json:"Port" name:"Port"`
			NetworkType      *string `json:"NetworkType" name:"NetworkType"`
			VpcId            *string `json:"VpcId" name:"VpcId"`
			SubnetId         *string `json:"SubnetId" name:"SubnetId"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			BillType         *string `json:"BillType" name:"BillType"`
			CreateDate       *int    `json:"CreateDate" name:"CreateDate"`
			ProjectId        *string `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			NodeNum          *string `json:"NodeNum" name:"NodeNum"`
			AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			ProductWhat      *string `json:"ProductWhat" name:"ProductWhat"`
			Mode             *string `json:"Mode" name:"Mode"`
			ModeName         *string `json:"ModeName" name:"ModeName"`
			Eip              *string `json:"Eip" name:"Eip"`
			WebEip           *string `json:"WebEip" name:"WebEip"`
			EipEgress        *string `json:"EipEgress" name:"EipEgress"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		UserId          *string `json:"UserId" name:"UserId"`
		Region          *string `json:"Region" name:"Region"`
		InstanceName    *string `json:"InstanceName" name:"InstanceName"`
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		StatusName      *string `json:"StatusName" name:"StatusName"`
		Status          *string `json:"Status" name:"Status"`
		Vip             *string `json:"Vip" name:"Vip"`
		WebVip          *string `json:"WebVip" name:"WebVip"`
		InstanceType    *string `json:"InstanceType" name:"InstanceType"`
		SsdDisk         *int    `json:"SsdDisk" name:"SsdDisk"`
		Protocol        *string `json:"Protocol" name:"Protocol"`
		SecurityGroupId *int    `json:"SecurityGroupId" name:"SecurityGroupId"`
		Port            *string `json:"Port" name:"Port"`
		NetworkType     *string `json:"NetworkType" name:"NetworkType"`
		VpcId           *string `json:"VpcId" name:"VpcId"`
		SubnetId        *string `json:"SubnetId" name:"SubnetId"`
		ProductId       *string `json:"ProductId" name:"ProductId"`
		BillType        *string `json:"BillType" name:"BillType"`
		CreateDate      *int    `json:"CreateDate" name:"CreateDate"`
		ProjectId       *string `json:"ProjectId" name:"ProjectId"`
		ProjectName     *string `json:"ProjectName" name:"ProjectName"`
		NodeNum         *string `json:"NodeNum" name:"NodeNum"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		ProductWhat     *string `json:"ProductWhat" name:"ProductWhat"`
		Mode            *string `json:"Mode" name:"Mode"`
		ModeName        *string `json:"ModeName" name:"ModeName"`
		Eip             *string `json:"Eip" name:"Eip"`
		WebEip          *string `json:"WebEip" name:"WebEip"`
		EipEgress       *string `json:"EipEgress" name:"EipEgress"`
	} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceNodesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceNodesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		NodeId *string `json:"NodeId" name:"NodeId"`
		Name   *string `json:"Name" name:"Name"`
		Role   *string `json:"Role" name:"Role"`
		Ip     *string `json:"Ip" name:"Ip"`
		Port   *string `json:"Port" name:"Port"`
		StatusName *string `json:"StatusName" name:"StatusName"`
		Status *string `json:"Status" name:"Status"`
	} `json:"Data"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeValidRegionRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeValidRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeValidRegionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Regions []struct {
			Name         *string `json:"Name" name:"Name"`
			Code         *string `json:"Code" name:"Code"`
			RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
			AreaCode     *string `json:"AreaCode" name:"AreaCode"`
			AreaName     *string `json:"AreaName" name:"AreaName"`
			AreaEnName   *string `json:"AreaEnName" name:"AreaEnName"`
			AvailabilityZones []struct {
				Code *string `json:"Code" name:"Code"`
				Name *string `json:"Name" name:"Name"`
			} `json:"AvailabilityZones"`
		} `json:"Regions" name:"Regions"`
	} `json:"Data"`
}

func (r *DescribeValidRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeValidRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRegionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRegionsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSecurityGroupRulesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSecurityGroupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSecurityGroupRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		Id     *string `json:"Id" name:"Id"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		ToPort *string `json:"ToPort" name:"ToPort"`
		Cidr   *string `json:"Cidr" name:"Cidr"`
	} `json:"Data"`
}

func (r *DescribeSecurityGroupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddSecurityGroupRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Cidrs      *string `json:"Cidrs,omitempty" name:"Cidrs"`
}

func (r *AddSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddSecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		Id     *string `json:"Id" name:"Id"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		ToPort *string `json:"ToPort" name:"ToPort"`
		Cidr   *string `json:"Cidr" name:"Cidr"`
	} `json:"Data"`
}

func (r *AddSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteSecurityGroupRulesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Cidrs      *string `json:"Cidrs,omitempty" name:"Cidrs"`
}

func (r *DeleteSecurityGroupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSecurityGroupRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		Id     *string `json:"Id" name:"Id"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		ToPort *string `json:"ToPort" name:"ToPort"`
		Cidr   *string `json:"Cidr" name:"Cidr"`
	} `json:"Data"`
}

func (r *DeleteSecurityGroupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ResetPasswordRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId       *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResetPasswordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RenameRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *RenameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *RenameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AllocateEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *AllocateEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AllocateEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *AllocateEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeallocateEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeallocateEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeallocateEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeallocateEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeallocateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SupportPluginsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *SupportPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SupportPluginsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		PluginName    *string `json:"PluginName" name:"PluginName"`
		NeedToRestart *bool   `json:"NeedToRestart" name:"NeedToRestart"`
	} `json:"Data"`
}

func (r *SupportPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SupportPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RestartInstanceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *RestartInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestartInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Instances []struct {
			InstanceId *string `json:"InstanceId" name:"InstanceId"`
			OperStatus *string `json:"OperStatus" name:"OperStatus"`
			Msg        *string `json:"Msg" name:"Msg"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListInstancePluginsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ListInstancePluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListInstancePluginsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		InstanceStatus *string `json:"InstanceStatus" name:"InstanceStatus"`
		Plugins    []struct {
			PluginName   *string `json:"PluginName" name:"PluginName"`
			PluginStatus *int    `json:"PluginStatus" name:"PluginStatus"`
		} `json:"Plugins" name:"Plugins"`
	} `json:"Data"`
}

func (r *ListInstancePluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstancePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableInstancePluginsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *EnableInstancePluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EnableInstancePluginsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Plugins []struct {
			PluginName   *string `json:"PluginName" name:"PluginName"`
			PluginStatus *int    `json:"PluginStatus" name:"PluginStatus"`
		} `json:"Plugins" name:"Plugins"`
	} `json:"Data"`
}

func (r *EnableInstancePluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableInstancePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableInstancePluginsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId     *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	DisablePlugins []*string `json:"DisablePlugins,omitempty" name:"DisablePlugins"`
}

func (r *DisableInstancePluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisableInstancePluginsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Plugins []struct {
			PluginName   *string `json:"PluginName" name:"PluginName"`
			PluginStatus *int    `json:"PluginStatus" name:"PluginStatus"`
		} `json:"Plugins" name:"Plugins"`
	} `json:"Data"`
}

func (r *DisableInstancePluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableInstancePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

