package v20250501
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type StrategyruleeditPolicies struct {
	Description  *string `json:"description,omitempty" name:"description"`
	Direction    *string `json:"direction,omitempty" name:"direction"`
	CidrBlock    *string `json:"cidrBlock,omitempty" name:"cidrBlock"`
	MaxPortRange *int    `json:"maxPortRange,omitempty" name:"maxPortRange"`
	MinPortRange *int    `json:"minPortRange,omitempty" name:"minPortRange"`
	Protocol     *string `json:"protocol,omitempty" name:"protocol"`
}
type StrategyrulecreatePolicies struct {
	Description  *string `json:"description,omitempty" name:"description"`
	Direction    *string `json:"direction,omitempty" name:"direction"`
	CidrBlock    *string `json:"cidrBlock,omitempty" name:"cidrBlock"`
	MaxPortRange *int    `json:"maxPortRange,omitempty" name:"maxPortRange"`
	MinPortRange *int    `json:"minPortRange,omitempty" name:"minPortRange"`
	Protocol     *string `json:"protocol,omitempty" name:"protocol"`
}


type CloudDeskreinstallRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
	ImageId    *string `json:"imageId,omitempty" name:"imageId"`
}

func (r *CloudDeskreinstallRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloudDeskreinstallResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *CloudDeskreinstallResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloudDeskreinstallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloudDeskmanageRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"instanceIds,omitempty" name:"instanceIds"`
	Action1     *string   `json:"action1,omitempty" name:"action1"`
}

func (r *CloudDeskmanageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloudDeskmanageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *CloudDeskmanageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloudDeskmanageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloudDeskeditRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
	Name       *string `json:"name,omitempty" name:"name"`
}

func (r *CloudDeskeditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloudDeskeditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *CloudDeskeditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloudDeskeditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloudDeskcreateRequest struct {
	*ksyunhttp.BaseRequest
	InstanceName    *string `json:"instanceName,omitempty" name:"instanceName"`
	InstanceType    *string `json:"instanceType,omitempty" name:"instanceType"`
	ImageId         *string `json:"imageId,omitempty" name:"imageId"`
	EdgeNodeId      *string `json:"edgeNodeId,omitempty" name:"edgeNodeId"`
	SystemDisk      *int    `json:"systemDisk,omitempty" name:"systemDisk"`
	DataDisk        *int    `json:"dataDisk,omitempty" name:"dataDisk"`
	BillType        *int    `json:"billType,omitempty" name:"billType"`
	Duration        *int    `json:"duration,omitempty" name:"duration"`
	SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
	Gpu             *string `json:"gpu,omitempty" name:"gpu"`
	Quantity        *int    `json:"quantity,omitempty" name:"quantity"`
	UniqueSuffix    *bool   `json:"uniqueSuffix,omitempty" name:"uniqueSuffix"`
}

func (r *CloudDeskcreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloudDeskcreateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		InstanceId []*string `json:"InstanceId" name:"InstanceId"`
	} `json:"Data"`
}

func (r *CloudDeskcreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloudDeskcreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloudDesklistRequest struct {
	*ksyunhttp.BaseRequest
	Page      *int    `json:"page,omitempty" name:"page"`
	Size      *int    `json:"size,omitempty" name:"size"`
	Connected *string `json:"connected,omitempty" name:"connected"`
	LabelIds  *string `json:"labelIds,omitempty" name:"labelIds"`
	Name      *string `json:"name,omitempty" name:"name"`
	UserName  *string `json:"userName,omitempty" name:"userName"`
}

func (r *CloudDesklistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloudDesklistResponse struct {
	*ksyunhttp.BaseResponse
	Message   *string `json:"message" name:"message"`
	Code      *int    `json:"code" name:"code"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		PageSize *int `json:"PageSize" name:"PageSize"`
		Records []struct {
			Id           *int      `json:"id" name:"id"`
			InstanceId   *string   `json:"instanceId" name:"instanceId"`
			UserName     *string   `json:"userName" name:"userName"`
			Name         *string   `json:"name" name:"name"`
			Region       *string   `json:"region" name:"region"`
			EcStatus     *string   `json:"ecStatus" name:"ecStatus"`
			Vcpu         *int      `json:"vcpu" name:"vcpu"`
			Vmemory      *string   `json:"vmemory" name:"vmemory"`
			VdiskList    []*string `json:"VdiskList" name:"VdiskList"`
			SysDisk      *string   `json:"sysDisk" name:"sysDisk"`
			ServerVmType *string   `json:"serverVmType" name:"serverVmType"`
			IpInfo       struct {
				PrivateIpV6 *string `json:"PrivateIpV6" name:"PrivateIpV6"`
				PrivateIpV4 *string `json:"PrivateIpV4" name:"PrivateIpV4"`
				PublicIpV6 *string `json:"PublicIpV6" name:"PublicIpV6"`
				PublicIpV4 *string `json:"PublicIpV4" name:"PublicIpV4"`
			} `json:"IpInfo"`
			CreatedTime     *string   `json:"createdTime" name:"createdTime"`
			SecurityGroupId *string   `json:"securityGroupId" name:"securityGroupId"`
			UserId          *int      `json:"userId" name:"userId"`
			BillingType     *int      `json:"billingType" name:"billingType"`
			SystemType      *string   `json:"systemType" name:"systemType"`
			AccountId       *int      `json:"accountId" name:"accountId"`
			CreateTime      *string   `json:"createTime" name:"createTime"`
			EdgeNodeId      *string   `json:"edgeNodeId" name:"edgeNodeId"`
			EdgeNodeName    *string   `json:"edgeNodeName" name:"edgeNodeName"`
			Connected       *string   `json:"connected" name:"connected"`
			Email           *string   `json:"email" name:"email"`
			Gpu             *string   `json:"gpu" name:"gpu"`
			Label           []*string `json:"Label" name:"Label"`
		} `json:"Records" name:"Records"`
		PageNum *int `json:"PageNum" name:"PageNum"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
	Detail *string `json:"detail" name:"detail"`
}

func (r *CloudDesklistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloudDesklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategyruleeditRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string                     `json:"securityGroupId,omitempty" name:"securityGroupId"`
	Policies        []*StrategyruleeditPolicies `json:"policies,omitempty" name:"policies"`
}

func (r *StrategyruleeditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategyruleeditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *StrategyruleeditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategyruleeditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategyrulecreateRequest struct {
	*ksyunhttp.BaseRequest
	Name        *string                       `json:"name,omitempty" name:"name"`
	Description *string                       `json:"description,omitempty" name:"description"`
	Policies    []*StrategyrulecreatePolicies `json:"policies,omitempty" name:"policies"`
}

func (r *StrategyrulecreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategyrulecreateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Id *int `json:"Id" name:"Id"`
	} `json:"Data"`
}

func (r *StrategyrulecreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategyrulecreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategyunboundRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
	InstanceId      *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *StrategyunboundRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategyunboundResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *StrategyunboundResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategyunboundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategyboundRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
	InstanceId      *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *StrategyboundRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategyboundResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *StrategyboundResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategyboundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategydeleteRequest struct {
	*ksyunhttp.BaseRequest
	Id []*int `json:"id,omitempty" name:"id"`
}

func (r *StrategydeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategydeleteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *StrategydeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategydeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategyeditRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
	Name            *string `json:"name,omitempty" name:"name"`
	Description     *string `json:"description,omitempty" name:"description"`
}

func (r *StrategyeditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategyeditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *StrategyeditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategyeditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategycreateRequest struct {
	*ksyunhttp.BaseRequest
	Name        *string `json:"name,omitempty" name:"name"`
	Description *string `json:"description,omitempty" name:"description"`
}

func (r *StrategycreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategycreateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Id         *int    `json:"Id" name:"Id"`
		StrategyId *string `json:"StrategyId" name:"StrategyId"`
		StrategyName *string `json:"StrategyName" name:"StrategyName"`
		AccountId  *int    `json:"AccountId" name:"AccountId"`
		Comment    *string `json:"Comment" name:"Comment"`
		InstanceNum *string `json:"InstanceNum" name:"InstanceNum"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
}

func (r *StrategycreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategycreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategylistRequest struct {
	*ksyunhttp.BaseRequest
	Size *int    `json:"size,omitempty" name:"size"`
	Page *int    `json:"page,omitempty" name:"page"`
	Name *string `json:"name,omitempty" name:"name"`
}

func (r *StrategylistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategylistResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		PageNum *int `json:"PageNum" name:"PageNum"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Records []struct {
			Id              *int    `json:"id" name:"id"`
			SecurityGroupId *string `json:"securityGroupId" name:"securityGroupId"`
			Name            *string `json:"name" name:"name"`
			Description     *string `json:"description" name:"description"`
			InstanceNum     *int    `json:"instanceNum" name:"instanceNum"`
			CreateTime      *string `json:"createTime" name:"createTime"`
			UpdateTime      *string `json:"updateTime" name:"updateTime"`
		} `json:"Records" name:"Records"`
	} `json:"Data"`
}

func (r *StrategylistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategylistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RolesdeleteRequest struct {
	*ksyunhttp.BaseRequest
	Id *int `json:"id,omitempty" name:"id"`
}

func (r *RolesdeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RolesdeleteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *RolesdeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RolesdeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RoleseditRequest struct {
	*ksyunhttp.BaseRequest
	Id           *int    `json:"id,omitempty" name:"id"`
	Name         *string `json:"name,omitempty" name:"name"`
	Description  *string `json:"description,omitempty" name:"description"`
	FileTransfer *int    `json:"fileTransfer,omitempty" name:"fileTransfer"`
	Clipboard    *int    `json:"clipboard,omitempty" name:"clipboard"`
	Disk         *int    `json:"disk,omitempty" name:"disk"`
	Usb          *int    `json:"usb,omitempty" name:"usb"`
}

func (r *RoleseditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RoleseditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *RoleseditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RoleseditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RolescreateRequest struct {
	*ksyunhttp.BaseRequest
	Name         *string `json:"name,omitempty" name:"name"`
	Description  *string `json:"description,omitempty" name:"description"`
	FileTransfer *int    `json:"fileTransfer,omitempty" name:"fileTransfer"`
	Clipboard    *int    `json:"clipboard,omitempty" name:"clipboard"`
	Disk         *int    `json:"disk,omitempty" name:"disk"`
	Usb          *int    `json:"usb,omitempty" name:"usb"`
}

func (r *RolescreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RolescreateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Id *int `json:"Id" name:"Id"`
	} `json:"Data"`
}

func (r *RolescreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RolescreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RoleslistRequest struct {
	*ksyunhttp.BaseRequest
	Size *int    `json:"size,omitempty" name:"size"`
	Page *int    `json:"page,omitempty" name:"page"`
	Name *string `json:"name,omitempty" name:"name"`
}

func (r *RoleslistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RoleslistResponse struct {
	*ksyunhttp.BaseResponse
	Message   *string `json:"message" name:"message"`
	Code      *int    `json:"code" name:"code"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		PageNum *int `json:"PageNum" name:"PageNum"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Records []struct {
			Id           *int    `json:"id" name:"id"`
			CreatedAt    *string `json:"createdAt" name:"createdAt"`
			Name         *string `json:"name" name:"name"`
			Description  *string `json:"description" name:"description"`
			UserNums     *int    `json:"userNums" name:"userNums"`
			FileTransfer *int    `json:"fileTransfer" name:"fileTransfer"`
			Clipboard    *int    `json:"clipboard" name:"clipboard"`
			Disk         *int    `json:"disk" name:"disk"`
			Usb          *int    `json:"usb" name:"usb"`
		} `json:"Records" name:"Records"`
	} `json:"Data"`
}

func (r *RoleslistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RoleslistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ImagedeleteRequest struct {
	*ksyunhttp.BaseRequest
	ImageId *string `json:"imageId,omitempty" name:"imageId"`
}

func (r *ImagedeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ImagedeleteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *ImagedeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImagedeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ImageeditRequest struct {
	*ksyunhttp.BaseRequest
	Id          *int    `json:"id,omitempty" name:"id"`
	ImageId     *string `json:"imageId,omitempty" name:"imageId"`
	ImageName   *string `json:"imageName,omitempty" name:"imageName"`
	Description *string `json:"description,omitempty" name:"description"`
}

func (r *ImageeditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ImageeditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *ImageeditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImageeditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ImagecreateRequest struct {
	*ksyunhttp.BaseRequest
	ImageName   *string `json:"imageName,omitempty" name:"imageName"`
	Description *string `json:"description,omitempty" name:"description"`
	InstanceId  *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *ImagecreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ImagecreateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		ImageID *int `json:"ImageID" name:"ImageID"`
	} `json:"Data"`
}

func (r *ImagecreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImagecreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ImagelistRequest struct {
	*ksyunhttp.BaseRequest
	Size *int    `json:"size,omitempty" name:"size"`
	Page *int    `json:"page,omitempty" name:"page"`
	Name *string `json:"name,omitempty" name:"name"`
}

func (r *ImagelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ImagelistResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		PageNum *int `json:"PageNum" name:"PageNum"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Records []struct {
			Id            *int    `json:"id" name:"id"`
			Name          *string `json:"name" name:"name"`
			AccountId     *int    `json:"accountId" name:"accountId"`
			ImageType     *string `json:"imageType" name:"imageType"`
			ImageSize     *string `json:"imageSize" name:"imageSize"`
			Description   *string `json:"description" name:"description"`
			Status        *int    `json:"status" name:"status"`
			ImageId       *string `json:"imageId" name:"imageId"`
			SystemType    *string `json:"systemType" name:"systemType"`
			SystemVersion *string `json:"systemVersion" name:"systemVersion"`
			CreateTime    *string `json:"createTime" name:"createTime"`
			UpdateTime    *string `json:"updateTime" name:"updateTime"`
		} `json:"Records" name:"Records"`
	} `json:"Data"`
}

func (r *ImagelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImagelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StrategyrulebatchEditRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" name:"securityGroupIds"`
}

func (r *StrategyrulebatchEditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StrategyrulebatchEditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *StrategyrulebatchEditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StrategyrulebatchEditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type MonitorregionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *MonitorregionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type MonitorregionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Regions []struct {
			Name  *string `json:"name" name:"name"`
			Value *string `json:"value" name:"value"`
		} `json:"Regions" name:"Regions"`
	} `json:"Data"`
}

func (r *MonitorregionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorregionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UsersinstancebindRequest struct {
	*ksyunhttp.BaseRequest
	Id         *int    `json:"id,omitempty" name:"id"`
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *UsersinstancebindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UsersinstancebindResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *UsersinstancebindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UsersinstancebindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UserspasswordresetRequest struct {
	*ksyunhttp.BaseRequest
	Id       *int    `json:"id,omitempty" name:"id"`
	Password *string `json:"password,omitempty" name:"password"`
}

func (r *UserspasswordresetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UserspasswordresetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *UserspasswordresetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserspasswordresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UsersdeleteRequest struct {
	*ksyunhttp.BaseRequest
	Id *int `json:"id,omitempty" name:"id"`
}

func (r *UsersdeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UsersdeleteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *UsersdeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UsersdeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UserseditRequest struct {
	*ksyunhttp.BaseRequest
	Id           *int    `json:"id,omitempty" name:"id"`
	Name         *string `json:"name,omitempty" name:"name"`
	PhoneOrEmail *string `json:"phoneOrEmail,omitempty" name:"phoneOrEmail"`
	NickName     *string `json:"nickName,omitempty" name:"nickName"`
	Status       *int    `json:"status,omitempty" name:"status"`
}

func (r *UserseditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UserseditResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *UserseditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserseditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UserscreateRequest struct {
	*ksyunhttp.BaseRequest
	Name         *string `json:"name,omitempty" name:"name"`
	Password     *string `json:"password,omitempty" name:"password"`
	PhoneOrEmail *string `json:"phoneOrEmail,omitempty" name:"phoneOrEmail"`
	RoleId       *int    `json:"roleId,omitempty" name:"roleId"`
}

func (r *UserscreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UserscreateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Id *int `json:"Id" name:"Id"`
	} `json:"Data"`
}

func (r *UserscreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserscreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UserslistRequest struct {
	*ksyunhttp.BaseRequest
	Size     *int    `json:"size,omitempty" name:"size"`
	Page     *int    `json:"page,omitempty" name:"page"`
	Username *string `json:"username,omitempty" name:"username"`
	Phone    *string `json:"phone,omitempty" name:"phone"`
	Email    *string `json:"email,omitempty" name:"email"`
}

func (r *UserslistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UserslistResponse struct {
	*ksyunhttp.BaseResponse
	Message   *string `json:"message" name:"message"`
	Code      *int    `json:"code" name:"code"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		PageNum *int `json:"PageNum" name:"PageNum"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Records []struct {
			Id           *int    `json:"id" name:"id"`
			CreateTime   *string `json:"createTime" name:"createTime"`
			UserName     *string `json:"userName" name:"userName"`
			NickName     *string `json:"nickName" name:"nickName"`
			Status       *string `json:"status" name:"status"`
			DeskNums     *int    `json:"deskNums" name:"deskNums"`
			RoleId       *int    `json:"roleId" name:"roleId"`
			RoleName     *string `json:"roleName" name:"roleName"`
			PhoneOrEmail *string `json:"phoneOrEmail" name:"phoneOrEmail"`
		} `json:"Records" name:"Records"`
	} `json:"Data"`
}

func (r *UserslistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserslistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloudDeskgetDesktopUrlRequest struct {
	*ksyunhttp.BaseRequest
	Token      *string `json:"token,omitempty" name:"token"`
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *CloudDeskgetDesktopUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloudDeskgetDesktopUrlResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Code *int    `json:"Code" name:"Code"`
		Url  *string `json:"Url" name:"Url"`
	} `json:"Data"`
}

func (r *CloudDeskgetDesktopUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloudDeskgetDesktopUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type QueryCloudDesksubmitShellRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds  []*string `json:"instanceIds,omitempty" name:"instanceIds"`
	Name         *string   `json:"name,omitempty" name:"name"`
	ShellContent *string   `json:"shellContent,omitempty" name:"shellContent"`
}

func (r *QueryCloudDesksubmitShellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryCloudDesksubmitShellResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		PlanId *int `json:"PlanId" name:"PlanId"`
	} `json:"Data"`
}

func (r *QueryCloudDesksubmitShellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCloudDesksubmitShellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateCloudDeskgetTokenRequest struct {
	*ksyunhttp.BaseRequest
	Username *string `json:"username,omitempty" name:"username"`
	Password *string `json:"password,omitempty" name:"password"`
}

func (r *CreateCloudDeskgetTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateCloudDeskgetTokenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		ExpiresAt *int    `json:"ExpiresAt" name:"ExpiresAt"`
		Token     *string `json:"Token" name:"Token"`
	} `json:"Data"`
}

func (r *CreateCloudDeskgetTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudDeskgetTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type QueryShellStatusRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds *string `json:"instanceIds,omitempty" name:"instanceIds"`
	PlanId      *int    `json:"planId,omitempty" name:"planId"`
}

func (r *QueryShellStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryShellStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      []struct {
		PlanUuid *int    `json:"PlanUuid" name:"PlanUuid"`
		Status   *int    `json:"Status" name:"Status"`
		Name     *string `json:"Name" name:"Name"`
		Version  *string `json:"Version" name:"Version"`
		Msg      *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *QueryShellStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryShellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetProxyIpRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"instanceIds,omitempty" name:"instanceIds"`
	Province    *string   `json:"province,omitempty" name:"province"`
	Isp         *string   `json:"isp,omitempty" name:"isp"`
	City        *string   `json:"city,omitempty" name:"city"`
}

func (r *SetProxyIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetProxyIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		PlanId *int `json:"PlanId" name:"PlanId"`
	} `json:"Data"`
}

func (r *SetProxyIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProxyIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetProxyConfigRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *GetProxyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetProxyConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      []struct {
		Name       *string `json:"Name" name:"Name"`
		CreateTime *int    `json:"CreateTime" name:"CreateTime"`
		City       *string `json:"City" name:"City"`
		Isp        *string `json:"Isp" name:"Isp"`
		PlanUuid   *string `json:"PlanUuid" name:"PlanUuid"`
		Province   *string `json:"Province" name:"Province"`
		Version    *string `json:"Version" name:"Version"`
	} `json:"Data"`
}

func (r *GetProxyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProxyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type QueryRuledetailRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
}

func (r *QueryRuledetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryRuledetailResponse struct {
	*ksyunhttp.BaseResponse
	Message   *string `json:"message" name:"message"`
	Code      *int    `json:"code" name:"code"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Name                  *string `json:"Name" name:"Name"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		Description           *string `json:"Description" name:"Description"`
		SecurityGroupPolicies []struct {
			Protocol              *string `json:"protocol" name:"protocol"`
			Description           *string `json:"description" name:"description"`
			MaxPortRange          *int    `json:"maxPortRange" name:"maxPortRange"`
			MinPortRange          *int    `json:"minPortRange" name:"minPortRange"`
			SecurityGroupPolicyId *int    `json:"securityGroupPolicyId" name:"securityGroupPolicyId"`
			Direction             *string `json:"direction" name:"direction"`
			SecurityGroupId       *string `json:"securityGroupId" name:"securityGroupId"`
		} `json:"SecurityGroupPolicies" name:"SecurityGroupPolicies"`
	} `json:"Data"`
}

func (r *QueryRuledetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRuledetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type QueryUsersinfoRequest struct {
	*ksyunhttp.BaseRequest
	Username *string `json:"username,omitempty" name:"username"`
	Phone    *string `json:"phone,omitempty" name:"phone"`
	Email    *string `json:"email,omitempty" name:"email"`
}

func (r *QueryUsersinfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryUsersinfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Detail    *string `json:"detail" name:"detail"`
	Data      struct {
		Id        *int    `json:"Id" name:"Id"`
		AccountId *int    `json:"AccountId" name:"AccountId"`
		UserName  *string `json:"UserName" name:"UserName"`
		NickName  *string `json:"NickName" name:"NickName"`
		Phone     *string `json:"Phone" name:"Phone"`
		Email     *string `json:"Email" name:"Email"`
		PhoneOrEmail *string `json:"PhoneOrEmail" name:"PhoneOrEmail"`
		DeskNums  *int    `json:"DeskNums" name:"DeskNums"`
		RoleId    *int    `json:"RoleId" name:"RoleId"`
		RoleName  *string `json:"RoleName" name:"RoleName"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
}

func (r *QueryUsersinfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUsersinfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetDetailRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *GetDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      struct {
		InstanceId        *string   `json:"InstanceId" name:"InstanceId"`
		Name              *string   `json:"Name" name:"Name"`
		Vcpu              *int      `json:"Vcpu" name:"Vcpu"`
		Vmemory           *string   `json:"Vmemory" name:"Vmemory"`
		SysDisk           *string   `json:"SysDisk" name:"SysDisk"`
		CreateTime        *string   `json:"CreateTime" name:"CreateTime"`
		EdgeNodeId        *string   `json:"EdgeNodeId" name:"EdgeNodeId"`
		EdgeNodeName      *string   `json:"EdgeNodeName" name:"EdgeNodeName"`
		EcStatus          *string   `json:"EcStatus" name:"EcStatus"`
		SecurityGroupId   *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
		ImageId           *string   `json:"ImageId" name:"ImageId"`
		Label             []*string `json:"Label" name:"Label"`
		UserName          *string   `json:"UserName" name:"UserName"`
		UserId            *int      `json:"UserId" name:"UserId"`
	} `json:"Data"`
}

func (r *GetDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListLabelRequest struct {
	*ksyunhttp.BaseRequest
	Name *string `json:"name,omitempty" name:"name"`
}

func (r *ListLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListLabelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      []struct {
		Id *int `json:"Id" name:"Id"`
		Name *string `json:"Name" name:"Name"`
	} `json:"Data"`
}

func (r *ListLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CancelInstanceLabelRequest struct {
	*ksyunhttp.BaseRequest
	LabelId    []*int  `json:"labelId,omitempty" name:"labelId"`
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *CancelInstanceLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CancelInstanceLabelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *CancelInstanceLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelInstanceLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateInstanceLabelRequest struct {
	*ksyunhttp.BaseRequest
	LabelId    []*int  `json:"labelId,omitempty" name:"labelId"`
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *UpdateInstanceLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateInstanceLabelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *UpdateInstanceLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteLabelRequest struct {
	*ksyunhttp.BaseRequest
	Id []*int `json:"id,omitempty" name:"id"`
}

func (r *DeleteLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteLabelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *DeleteLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateLabelRequest struct {
	*ksyunhttp.BaseRequest
	Id   *int    `json:"id,omitempty" name:"id"`
	Name *string `json:"name,omitempty" name:"name"`
}

func (r *UpdateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateLabelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
	Data      *bool   `json:"data" name:"data"`
}

func (r *UpdateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLabelRequest struct {
	*ksyunhttp.BaseRequest
	Name *string `json:"name,omitempty" name:"name"`
}

func (r *CreateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateLabelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Data      *bool   `json:"data" name:"data"`
	Code      *int    `json:"code" name:"code"`
	Message   *string `json:"message" name:"message"`
}

func (r *CreateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

