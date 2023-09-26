package v20210520
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type CreateInstanceModuleConfigs struct {
}


type CreateInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
    EnableModules *string `json:"EnableModules,omitempty" name:"EnableModules"`
    ModuleConfigs []*CreateInstanceModuleConfigs `json:"ModuleConfigs,omitempty" name:"ModuleConfigs"`
    AdminUser *string `json:"AdminUser,omitempty" name:"AdminUser"`
    AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    BillType *int `json:"BillType,omitempty" name:"BillType"`
    Duration *string `json:"Duration,omitempty" name:"Duration"`
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    BackupConfigMaxBackups *int `json:"BackupConfig.MaxBackups,omitempty" name:"BackupConfig.MaxBackups"`
    BackupConfigMaxReservedHours *int `json:"BackupConfig.MaxReservedHours,omitempty" name:"BackupConfig.MaxReservedHours"`
    BackupConfigPreferredBackupTime *string `json:"BackupConfig.PreferredBackupTime,omitempty" name:"BackupConfig.PreferredBackupTime"`
    EnableAutoBackup *bool `json:"EnableAutoBackup,omitempty" name:"EnableAutoBackup"`
    Engine *string `json:"Engine,omitempty" name:"Engine"`
    EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
    ClientPort *int `json:"ClientPort,omitempty" name:"ClientPort"`
    Az *string `json:"Az,omitempty" name:"Az"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		InstanceType *string `json:"InstanceType"`
		OrderId *string `json:"OrderId"`
		AccountId *string `json:"AccountId"`
		Region *string `json:"Region"`
		Az *string `json:"Az"`
	} `json:"Data"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ProjectIds *string `json:"ProjectIds,omitempty" name:"ProjectIds"`
    InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Offset *int `json:"Offset" name:"Offset"`
    Limit *int `json:"Limit" name:"Limit"`
    Total *int `json:"Total" name:"Total"`
	Data []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		AdminUser *string `json:"AdminUser"`
		StatusName *string `json:"StatusName"`
		Status *string `json:"Status"`
		VpcId *string `json:"VpcId"`
		SubnetId *string `json:"SubnetId"`
		Engine *string `json:"Engine"`
		EngineVersion *string `json:"EngineVersion"`
		ProjectId *string `json:"ProjectId"`
		ProjectName *string `json:"ProjectName"`
		BillType *int `json:"BillType"`
		ProductId *string `json:"ProductId"`
		ProductType *int `json:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName"`
		ProductWhat *int `json:"ProductWhat"`
		CreateDate *string `json:"CreateDate"`
		UpdateDate *string `json:"UpdateDate"`
		Region *string `json:"Region"`
		RegionName *string `json:"RegionName"`
		Az *string `json:"Az"`
		AzName *string `json:"AzName"`
		UserId *string `json:"UserId"`
		SecurityGroupId *string `json:"SecurityGroupId"`
		ClientIp *string `json:"ClientIp"`
		ClientPort *int `json:"ClientPort"`
		EnableModules *string `json:"EnableModules"`
		ModuleConfigs []struct {
					moduleType *string `json:"moduleType"`
					packageCode *string `json:"packageCode"`
					replicas *int `json:"replicas"`
					cpu *int `json:"cpu"`
					mem *int `json:"mem"`
					tidisk *int `json:"tidisk"`
					createTime *string `json:"createTime"`
			} `json:"ModuleConfigs"`
			EndTime *string `json:"EndTime"`
		} `json:"Data"`
}

func (r *ListInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListInstanceResponse) FromJsonString(s string) error {
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

func (r *DescribeInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		EnableModules *string `json:"EnableModules"`
		ModuleConfigs []struct {
					moduleType *string `json:"moduleType"`
					packageCode *string `json:"packageCode"`
					replicas *int `json:"replicas"`
					cpu *int `json:"cpu"`
					mem *int `json:"mem"`
					tidisk *int `json:"tidisk"`
					createTime *string `json:"createTime"`
			} `json:"ModuleConfigs"`
			AdminUser *string `json:"AdminUser"`
			StatusName *string `json:"StatusName"`
			Status *string `json:"Status"`
			VpcId *string `json:"VpcId"`
			SubnetId *string `json:"SubnetId"`
			Engine *string `json:"Engine"`
			EngineVersion *string `json:"EngineVersion"`
			ProjectId *string `json:"ProjectId"`
			ProjectName *string `json:"ProjectName"`
			BillType *int `json:"BillType"`
			ProductId *string `json:"ProductId"`
			ProductType *int `json:"ProductType"`
			ProductTypeName *string `json:"ProductTypeName"`
			ProductWhat *int `json:"ProductWhat"`
			CreateDate *string `json:"CreateDate"`
			Region *string `json:"Region"`
			RegionName *string `json:"RegionName"`
			Az *string `json:"Az"`
			AzName *string `json:"AzName"`
			UserId *string `json:"UserId"`
			SecurityGroupId *string `json:"SecurityGroupId"`
			SecurityGroupName *string `json:"SecurityGroupName"`
			SecurityGroupDesc *string `json:"SecurityGroupDesc"`
			BackupConfig struct {
					maxBackups *int `json:"maxBackups"`
					maxReservedHours *int `json:"maxReservedHours"`
					preferredBackupTime *string `json:"preferredBackupTime"`
			} `json:"BackupConfig"`
			ClientIp *string `json:"ClientIp"`
			ClientPort *int `json:"ClientPort"`
			EndTime *string `json:"EndTime"`
			DashboardClientVip *string `json:"DashboardClientVip"`
			DashboardClientVport *int `json:"DashboardClientVport"`
			TimonitorURL *string `json:"TimonitorURL"`
		} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenameInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RenameInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *string `json:"Data" name:"Data"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListRegionRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *ListRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListRegionRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRegionRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListRegionResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		RegionId *string `json:"RegionId"`
		InnerCode *string `json:"InnerCode"`
		RegionCode *string `json:"RegionCode"`
		RegionName *string `json:"RegionName"`
		RegionEnName *string `json:"RegionEnName"`
		Active *bool `json:"Active"`
		RegionType *int `json:"RegionType"`
		Overseas *bool `json:"Overseas"`
		AzList []struct {
					AzCode *string `json:"AzCode"`
					AzName *string `json:"AzName"`
			} `json:"AzList"`
			AreaCode *string `json:"AreaCode"`
			AreaName *string `json:"AreaName"`
			AreaEnName *string `json:"AreaEnName"`
		} `json:"Data"`
}

func (r *ListRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescRegionRequest struct {
    *ksyunhttp.BaseRequest
    RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *DescRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescRegionRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescRegionRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescRegionResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		RegionId *string `json:"RegionId"`
		InnerCode *string `json:"InnerCode"`
		RegionCode *string `json:"RegionCode"`
		RegionName *string `json:"RegionName"`
		RegionEnName *string `json:"RegionEnName"`
		Active *bool `json:"Active"`
		RegionType *int `json:"RegionType"`
		Overseas *bool `json:"Overseas"`
		AzList []struct {
					AzCode *string `json:"AzCode"`
					AzName *string `json:"AzName"`
			} `json:"AzList"`
			AreaCode *string `json:"AreaCode"`
			AreaName *string `json:"AreaName"`
			AreaEnName *string `json:"AreaEnName"`
		} `json:"Data"`
}

func (r *DescRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		GroupType *string `json:"GroupType"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Region *string `json:"Region"`
		Rules *string `json:"Rules"`
		Instances *string `json:"Instances"`
	} `json:"Data"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Offset *int `json:"Offset" name:"Offset"`
    Limit *int `json:"Limit" name:"Limit"`
    Total *int `json:"Total" name:"Total"`
	Data []struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		GroupType *string `json:"GroupType"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Region *string `json:"Region"`
	} `json:"Data"`
}

func (r *ListSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		GroupType *string `json:"GroupType"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Region *string `json:"Region"`
		Rules *string `json:"Rules"`
		Instances *string `json:"Instances"`
	} `json:"Data"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		GroupType *string `json:"GroupType"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Region *string `json:"Region"`
		Rules *string `json:"Rules"`
		Instances *string `json:"Instances"`
	} `json:"Data"`
}

func (r *UpdateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SrcSecurityGroupId *string `json:"SrcSecurityGroupId,omitempty" name:"SrcSecurityGroupId"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloneSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		GroupType *string `json:"GroupType"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Region *string `json:"Region"`
		Rules *string `json:"Rules"`
		Instances *string `json:"Instances"`
	} `json:"Data"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *BindSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BindSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type BindSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *bool `json:"Data" name:"Data"`
}

func (r *BindSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *UnbindSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnbindSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UnbindSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *bool `json:"Data" name:"Data"`
}

func (r *UnbindSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebindSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RebindSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebindSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RebindSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RebindSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *bool `json:"Data" name:"Data"`
}

func (r *RebindSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebindSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRuleRequest struct {
    *ksyunhttp.BaseRequest
    Rules []*string `json:"Rules,omitempty" name:"Rules"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *CreateSecurityRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityRuleRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityRuleRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRuleResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *bool `json:"Data" name:"Data"`
}

func (r *CreateSecurityRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSecurityRuleRequest struct {
    *ksyunhttp.BaseRequest
    RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
    Description *string `json:"Description,omitempty" name:"Description"`
    Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

func (r *UpdateSecurityRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSecurityRuleRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateSecurityRuleRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSecurityRuleResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *bool `json:"Data" name:"Data"`
}

func (r *UpdateSecurityRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSecurityRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUnsecuredInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ProjectIds *string `json:"ProjectIds,omitempty" name:"ProjectIds"`
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListUnsecuredInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUnsecuredInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListUnsecuredInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListUnsecuredInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Offset *int `json:"Offset" name:"Offset"`
    Limit *int `json:"Limit" name:"Limit"`
    Total *int `json:"Total" name:"Total"`
	Data []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		ProductType *int `json:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		ClientIp *string `json:"ClientIp"`
		ClientPort *int `json:"ClientPort"`
	} `json:"Data"`
}

func (r *ListUnsecuredInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUnsecuredInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListBackupRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListBackupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListBackupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListBackupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		BackupId *string `json:"BackupId"`
		InstanceId *string `json:"InstanceId"`
		BackupName *string `json:"BackupName"`
		BackupStatus *string `json:"BackupStatus"`
		BackupStatusName *string `json:"BackupStatusName"`
		BackupSizeReadable *string `json:"BackupSizeReadable"`
		RunType *int `json:"RunType"`
		RunTypeName *string `json:"RunTypeName"`
		IncreaseName *string `json:"IncreaseName"`
		StartTime *string `json:"StartTime"`
		ComplateTime *string `json:"ComplateTime"`
		Cost *int `json:"Cost"`
		CreateTime *string `json:"CreateTime"`
	} `json:"Data"`
}

func (r *ListBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

