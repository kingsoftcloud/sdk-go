package v20250501
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type Strategyrulecreatepolicies struct {
    Description *string `json:"description,omitempty" name:"description"`
    Direction *string `json:"direction,omitempty" name:"direction"`
    CidrBlock *string `json:"cidrBlock,omitempty" name:"cidrBlock"`
    MaxPortRange *int `json:"maxPortRange,omitempty" name:"maxPortRange"`
    MinPortRange *int `json:"minPortRange,omitempty" name:"minPortRange"`
    Protocol *string `json:"protocol,omitempty" name:"protocol"`
}


type CloudDeskreinstallRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
    ImageId *string `json:"imageId,omitempty" name:"imageId"`
}

func (r *CloudDeskreinstallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloudDeskreinstallRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloudDeskreinstallRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CloudDeskreinstallResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    InstanceIds *string `json:"instanceIds,omitempty" name:"instanceIds"`
    Action1 *string `json:"action1,omitempty" name:"action1"`
}

func (r *CloudDeskmanageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloudDeskmanageRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloudDeskmanageRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CloudDeskmanageResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Name *string `json:"name,omitempty" name:"name"`
    Id *int `json:"id,omitempty" name:"id"`
}

func (r *CloudDeskeditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloudDeskeditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloudDeskeditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CloudDeskeditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    InstanceName *string `json:"instanceName,omitempty" name:"instanceName"`
    InstanceType *string `json:"instanceType,omitempty" name:"instanceType"`
    ImageId *string `json:"imageId,omitempty" name:"imageId"`
    EdgeNodeId *string `json:"edgeNodeId,omitempty" name:"edgeNodeId"`
    SystemDisk *int `json:"systemDisk,omitempty" name:"systemDisk"`
    DataDisk *int `json:"dataDisk,omitempty" name:"dataDisk"`
    BillType *int `json:"billType,omitempty" name:"billType"`
    Duration *int `json:"duration,omitempty" name:"duration"`
    SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
    Gpu *string `json:"gpu,omitempty" name:"gpu"`
    Quantity *int `json:"quantity,omitempty" name:"quantity"`
    UniqueSuffix *bool `json:"uniqueSuffix,omitempty" name:"uniqueSuffix"`
}

func (r *CloudDeskcreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloudDeskcreateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloudDeskcreateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CloudDeskcreateResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		InstanceId []struct {
			} `json:"InstanceId"`
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
    Page *int `json:"page,omitempty" name:"page"`
    Size *int `json:"size,omitempty" name:"size"`
    Connected *string `json:"connected,omitempty" name:"connected"`
}

func (r *CloudDesklistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloudDesklistRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloudDesklistRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CloudDesklistResponse struct {
    *ksyunhttp.BaseResponse
    message *string `json:"message" name:"message"`
    code *int `json:"code" name:"code"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		PageSize *int `json:"PageSize"`
		Records []struct {
					id *int `json:"id"`
					instanceId *string `json:"instanceId"`
					userName *string `json:"userName"`
					name *string `json:"name"`
					region *string `json:"region"`
					ecStatus *string `json:"ecStatus"`
					vcpu *int `json:"vcpu"`
					vmemory *int `json:"vmemory"`
				VdiskList []struct {
				} `json:"VdiskList"`
					sysDisk *string `json:"sysDisk"`
					serverVmType *string `json:"serverVmType"`
				IpInfo struct {
					PrivateIpV6 *string `json:"PrivateIpV6"`
					PrivateIpV4 *string `json:"PrivateIpV4"`
					PublicIpV6 *string `json:"PublicIpV6"`
					PublicIpV4 *string `json:"PublicIpV4"`
				} `json:"IpInfo"`
					createdTime *string `json:"createdTime"`
					securityGroupId *string `json:"securityGroupId"`
					userId *int `json:"userId"`
					billingType *int `json:"billingType"`
					systemType *string `json:"systemType"`
					accountId *int `json:"accountId"`
					createTime *string `json:"createTime"`
					edgeNodeId *string `json:"edgeNodeId"`
					edgeNodeName *string `json:"edgeNodeName"`
					connected *string `json:"connected"`
					email *string `json:"email"`
					gpu *string `json:"gpu"`
			} `json:"Records"`
			PageNum *int `json:"PageNum"`
			TotalCount *int `json:"TotalCount"`
		} `json:"Data"`
    detail *string `json:"detail" name:"detail"`
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
    SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
}

func (r *StrategyruleeditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategyruleeditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategyruleeditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategyruleeditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Name *string `json:"name,omitempty" name:"name"`
    Description *string `json:"description,omitempty" name:"description"`
    Policies []*Strategyrulecreatepolicies `json:"policies,omitempty" name:"policies"`
}

func (r *StrategyrulecreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategyrulecreateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategyrulecreateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategyrulecreateResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		Id *int `json:"Id"`
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
    InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *StrategyunboundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategyunboundRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategyunboundRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategyunboundResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *StrategyboundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategyboundRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategyboundRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategyboundResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Id *int `json:"id,omitempty" name:"id"`
}

func (r *StrategydeleteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategydeleteRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategydeleteRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategydeleteResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Name *string `json:"name,omitempty" name:"name"`
    Description *string `json:"description,omitempty" name:"description"`
}

func (r *StrategyeditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategyeditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategyeditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategyeditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Name *string `json:"name,omitempty" name:"name"`
    Description *string `json:"description,omitempty" name:"description"`
}

func (r *StrategycreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategycreateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategycreateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategycreateResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		Id *int `json:"Id"`
		StrategyId *string `json:"StrategyId"`
		StrategyName *string `json:"StrategyName"`
		AccountId *int `json:"AccountId"`
		Comment *string `json:"Comment"`
		InstanceNum *string `json:"InstanceNum"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
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
    Size *int `json:"size,omitempty" name:"size"`
    Page *int `json:"page,omitempty" name:"page"`
    Name *string `json:"name,omitempty" name:"name"`
}

func (r *StrategylistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategylistRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategylistRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategylistResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		PageNum *int `json:"PageNum"`
		PageSize *int `json:"PageSize"`
		TotalCount *int `json:"TotalCount"`
		Records []struct {
					id *int `json:"id"`
					securityGroupId *string `json:"securityGroupId"`
					name *string `json:"name"`
					description *string `json:"description"`
					instanceNum *int `json:"instanceNum"`
					createTime *string `json:"createTime"`
					updateTime *string `json:"updateTime"`
			} `json:"Records"`
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

func (r *RolesdeleteRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RolesdeleteRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RolesdeleteResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Id *int `json:"id,omitempty" name:"id"`
    Name *string `json:"name,omitempty" name:"name"`
    Description *string `json:"description,omitempty" name:"description"`
    FileTransfer *int `json:"fileTransfer,omitempty" name:"fileTransfer"`
    Clipboard *int `json:"clipboard,omitempty" name:"clipboard"`
    Disk *int `json:"disk,omitempty" name:"disk"`
    Usb *int `json:"usb,omitempty" name:"usb"`
}

func (r *RoleseditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RoleseditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RoleseditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RoleseditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Name *string `json:"name,omitempty" name:"name"`
    Description *string `json:"description,omitempty" name:"description"`
    FileTransfer *int `json:"fileTransfer,omitempty" name:"fileTransfer"`
    Clipboard *int `json:"clipboard,omitempty" name:"clipboard"`
    Disk *int `json:"disk,omitempty" name:"disk"`
    Usb *int `json:"usb,omitempty" name:"usb"`
}

func (r *RolescreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RolescreateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RolescreateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RolescreateResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		Id *int `json:"Id"`
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
    Size *int `json:"size,omitempty" name:"size"`
    Page *int `json:"page,omitempty" name:"page"`
    Name *string `json:"name,omitempty" name:"name"`
}

func (r *RoleslistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RoleslistRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RoleslistRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RoleslistResponse struct {
    *ksyunhttp.BaseResponse
    message *string `json:"message" name:"message"`
    code *int `json:"code" name:"code"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		PageNum *int `json:"PageNum"`
		PageSize *int `json:"PageSize"`
		TotalCount *int `json:"TotalCount"`
		Records []struct {
					id *int `json:"id"`
					createdAt *string `json:"createdAt"`
					name *string `json:"name"`
					description *string `json:"description"`
					userNums *int `json:"userNums"`
					fileTransfer *int `json:"fileTransfer"`
					clipboard *int `json:"clipboard"`
					disk *int `json:"disk"`
					usb *int `json:"usb"`
			} `json:"Records"`
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

func (r *ImagedeleteRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImagedeleteRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ImagedeleteResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Id *int `json:"id,omitempty" name:"id"`
    ImageId *string `json:"imageId,omitempty" name:"imageId"`
    ImageName *string `json:"imageName,omitempty" name:"imageName"`
}

func (r *ImageeditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageeditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImageeditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ImageeditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    ImageName *string `json:"imageName,omitempty" name:"imageName"`
    Description *string `json:"description,omitempty" name:"description"`
    InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *ImagecreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImagecreateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImagecreateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ImagecreateResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Size *int `json:"size,omitempty" name:"size"`
    Page *int `json:"page,omitempty" name:"page"`
    Name *string `json:"name,omitempty" name:"name"`
}

func (r *ImagelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImagelistRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImagelistRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ImagelistResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		PageNum *int `json:"PageNum"`
		PageSize *int `json:"PageSize"`
		TotalCount *int `json:"TotalCount"`
		Records []struct {
					id *int `json:"id"`
					name *string `json:"name"`
					accountId *int `json:"accountId"`
					imageType *string `json:"imageType"`
					imageSize *string `json:"imageSize"`
					description *string `json:"description"`
					status *int `json:"status"`
					imageId *string `json:"imageId"`
					systemType *string `json:"systemType"`
					systemVersion *string `json:"systemVersion"`
					createTime *string `json:"createTime"`
					updateTime *string `json:"updateTime"`
			} `json:"Records"`
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
    SecurityGroupId *string `json:"securityGroupId,omitempty" name:"securityGroupId"`
}

func (r *StrategyrulebatchEditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StrategyrulebatchEditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StrategyrulebatchEditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type StrategyrulebatchEditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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

func (r *MonitorregionsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "MonitorregionsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type MonitorregionsResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		Regions []struct {
					name *string `json:"name"`
					value *string `json:"value"`
			} `json:"Regions"`
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
    Id *int `json:"id,omitempty" name:"id"`
    InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *UsersinstancebindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UsersinstancebindRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UsersinstancebindRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UsersinstancebindResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Id *int `json:"id,omitempty" name:"id"`
    Password *string `json:"password,omitempty" name:"password"`
}

func (r *UserspasswordresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UserspasswordresetRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UserspasswordresetRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UserspasswordresetResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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

func (r *UsersdeleteRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UsersdeleteRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UsersdeleteResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Id *int `json:"id,omitempty" name:"id"`
    Name *string `json:"name,omitempty" name:"name"`
    PhoneOrEmail *string `json:"phoneOrEmail,omitempty" name:"phoneOrEmail"`
    NickName *string `json:"nickName,omitempty" name:"nickName"`
    Status *int `json:"status,omitempty" name:"status"`
}

func (r *UserseditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UserseditRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UserseditRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UserseditResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
    data *bool `json:"data" name:"data"`
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
    Name *string `json:"name,omitempty" name:"name"`
    Password *string `json:"password,omitempty" name:"password"`
    PhoneOrEmail *string `json:"phoneOrEmail,omitempty" name:"phoneOrEmail"`
    RoleId *int `json:"roleId,omitempty" name:"roleId"`
}

func (r *UserscreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UserscreateRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UserscreateRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UserscreateResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
    code *int `json:"code" name:"code"`
    message *string `json:"message" name:"message"`
    detail *string `json:"detail" name:"detail"`
	Data struct {
		Id *int `json:"Id"`
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
    Size *int `json:"size,omitempty" name:"size"`
    Page *int `json:"page,omitempty" name:"page"`
}

func (r *UserslistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UserslistRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UserslistRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UserslistResponse struct {
    *ksyunhttp.BaseResponse
    message *string `json:"message" name:"message"`
    code *int `json:"code" name:"code"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		PageNum *int `json:"PageNum"`
		PageSize *int `json:"PageSize"`
		TotalCount *int `json:"TotalCount"`
		Records []struct {
					id *int `json:"id"`
					createdTime *string `json:"createdTime"`
					userName *string `json:"userName"`
					nickName *string `json:"nickName"`
					enable *int `json:"enable"`
					deskNums *int `json:"deskNums"`
					roleId *int `json:"roleId"`
					roleName *string `json:"roleName"`
					phoneOrEmail *string `json:"phoneOrEmail"`
			} `json:"Records"`
		} `json:"Data"`
}

func (r *UserslistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UserslistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

