package v20210101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type ModifyDBParameterGroupParameters struct {
}


type ListInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
    TagId *string `json:"TagId,omitempty" name:"TagId"`
    ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    OrderBy []*string `json:"OrderBy,omitempty" name:"OrderBy"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
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
    1 *int `json:"1" name:"1"`
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
		InstacneConfig *string `json:"InstacneConfig"`
		AdminUser *string `json:"AdminUser"`
		StatusName *string `json:"StatusName"`
		Status *string `json:"Status"`
		NetworkType *string `json:"NetworkType"`
		VpcId *string `json:"VpcId"`
		SubnetId *string `json:"SubnetId"`
		Vip *string `json:"Vip"`
		Engine *string `json:"Engine"`
		EngineVersion *string `json:"EngineVersion"`
		ProjectId *string `json:"ProjectId"`
		ProjectName *string `json:"ProjectName"`
		BillType *int `json:"BillType"`
		StorageSize *int `json:"StorageSize"`
		UsedStorageSize *int `json:"UsedStorageSize"`
		StorageType *string `json:"StorageType"`
		Mem *int `json:"Mem"`
		Cpu *int `json:"Cpu"`
		TcpPort *int `json:"TcpPort"`
		HttpPort *int `json:"HttpPort"`
		NodeNum *int `json:"NodeNum"`
		ProductId *string `json:"ProductId"`
		ProductType *int `json:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName"`
		ProductWhat *int `json:"ProductWhat"`
		CreateDate *string `json:"CreateDate"`
		UpdateDate *string `json:"UpdateDate"`
		Region *string `json:"Region"`
		Az *string `json:"Az"`
		UserId *string `json:"UserId"`
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		SecurityGroupDesc *string `json:"SecurityGroupDesc"`
		ServiceEndTime *string `json:"ServiceEndTime"`
		ShardList []struct {
					id *string `json:"id"`
					name *string `json:"name"`
			} `json:"ShardList"`
			Replicas *int `json:"Replicas"`
			DirectConnectionVips []struct {
				} `json:"DirectConnectionVips"`
				MultiAz *int `json:"MultiAz"`
				Area struct {
						Master *string `json:"Master"`
						Standby *string `json:"Standby"`
				} `json:"Area"`
			} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
    InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
    AdminUser *string `json:"AdminUser,omitempty" name:"AdminUser"`
    AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    Engine *string `json:"Engine,omitempty" name:"Engine"`
    EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    BillType *int `json:"BillType,omitempty" name:"BillType"`
    Duration *int `json:"Duration,omitempty" name:"Duration"`
    EbsSize *int `json:"EbsSize,omitempty" name:"EbsSize"`
    EbsType *string `json:"EbsType,omitempty" name:"EbsType"`
    Mem *int `json:"Mem,omitempty" name:"Mem"`
    Cpu *int `json:"Cpu,omitempty" name:"Cpu"`
    TcpPort *int `json:"TcpPort,omitempty" name:"TcpPort"`
    HttpPort *int `json:"HttpPort,omitempty" name:"HttpPort"`
    Az *string `json:"Az,omitempty" name:"Az"`
    NodeNum *int `json:"NodeNum,omitempty" name:"NodeNum"`
    PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
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

type DeleteInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
    DeleteDirectly *bool `json:"DeleteDirectly,omitempty" name:"DeleteDirectly"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		InstanceId *string `json:"InstanceId"`
		OperStatus *string `json:"OperStatus"`
		Msg *string `json:"Msg"`
	} `json:"Data"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
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

type ListSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
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
	Data []struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Rules []struct {
					RuleId *string `json:"RuleId"`
					Cidr *string `json:"Cidr"`
					CreateTime *string `json:"CreateTime"`
					Description *string `json:"Description"`
			} `json:"Rules"`
			Instances []struct {
						InstanceId *string `json:"InstanceId"`
						InstanceName *string `json:"InstanceName"`
						ProductType *int `json:"ProductType"`
						ProductTypeName *string `json:"ProductTypeName"`
						Vip *string `json:"Vip"`
						CreateTime *string `json:"CreateTime"`
						Status *string `json:"Status"`
						UpdatetTime *string `json:"UpdatetTime"`
				} `json:"Instances"`
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
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
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
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Rules []struct {
					RuleId *string `json:"RuleId"`
					Cidr *string `json:"Cidr"`
					CreateTime *string `json:"CreateTime"`
					Description *string `json:"Description"`
			} `json:"Rules"`
			Instances []struct {
						InstanceId *string `json:"InstanceId"`
						InstanceName *string `json:"InstanceName"`
						ProductType *int `json:"ProductType"`
						ProductTypeName *string `json:"ProductTypeName"`
						Vip *string `json:"Vip"`
						CreateTime *string `json:"CreateTime"`
						Status *string `json:"Status"`
						UpdatetTime *string `json:"UpdatetTime"`
				} `json:"Instances"`
			} `json:"Data"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
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
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Rules []struct {
			} `json:"Rules"`
			Instances []struct {
				} `json:"Instances"`
			} `json:"Data"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		OperStatus *string `json:"OperStatus"`
		Msg *string `json:"Msg"`
	} `json:"Data"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenameSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *RenameSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameSecurityGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameSecurityGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RenameSecurityGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		SecurityGroupId *string `json:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName"`
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Rules []struct {
			} `json:"Rules"`
			Instances []struct {
				} `json:"Instances"`
			} `json:"Data"`
}

func (r *RenameSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
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
		IpVersion *string `json:"IpVersion"`
		Description *string `json:"Description"`
		InstanceCount *int `json:"InstanceCount"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
		Rules []struct {
					RuleId *string `json:"RuleId"`
					Cidr *string `json:"Cidr"`
					CreateTime *string `json:"CreateTime"`
					Description *string `json:"Description"`
			} `json:"Rules"`
			Instances []struct {
				} `json:"Instances"`
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
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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

type CreateSecurityRuleRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    Cidrs []*string `json:"Cidrs,omitempty" name:"Cidrs"`
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

type DeleteSecurityRuleRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    RuleIds *string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DeleteSecurityRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityRuleRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityRuleRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityRuleResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *bool `json:"Data" name:"Data"`
}

func (r *DeleteSecurityRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSecuredInstanceRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    OrderBy []*string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListSecuredInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSecuredInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSecuredInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListSecuredInstanceResponse struct {
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
		Vip *string `json:"Vip"`
		CreateTime *string `json:"CreateTime"`
		Status *string `json:"Status"`
		UpdatetTime *string `json:"UpdatetTime"`
	} `json:"Data"`
}

func (r *ListSecuredInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSecuredInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUnsecuredInstanceRequest struct {
    *ksyunhttp.BaseRequest
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    OrderBy []*string `json:"OrderBy,omitempty" name:"OrderBy"`
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
		Vip *string `json:"Vip"`
		CreateTime *string `json:"CreateTime"`
		Status *string `json:"Status"`
		UpdatetTime *string `json:"UpdatetTime"`
	} `json:"Data"`
}

func (r *ListUnsecuredInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUnsecuredInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListRecycledInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
    ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`
    FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
    OrderBy []*string `json:"OrderBy,omitempty" name:"OrderBy"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListRecycledInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListRecycledInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRecycledInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListRecycledInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Offset *int `json:"Offset" name:"Offset"`
    Limit *int `json:"Limit" name:"Limit"`
    Total *int `json:"Total" name:"Total"`
    Data []*string `json:"Data" name:"Data"`
}

func (r *ListRecycledInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListRecycledInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecoverRecycledInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RecoverRecycledInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverRecycledInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RecoverRecycledInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RecoverRecycledInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		InstanceId *string `json:"InstanceId"`
		OperStatus *string `json:"OperStatus"`
		Msg *string `json:"Msg"`
	} `json:"Data"`
}

func (r *RecoverRecycledInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverRecycledInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropRecycledInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DropRecycledInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropRecycledInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DropRecycledInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DropRecycledInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		InstanceId *string `json:"InstanceId"`
		OperStatus *string `json:"OperStatus"`
		Msg *string `json:"Msg"`
	} `json:"Data"`
}

func (r *DropRecycledInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropRecycledInstanceResponse) FromJsonString(s string) error {
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
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
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

type UpdateSecurityRuleRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
    Description *string `json:"Description,omitempty" name:"Description"`
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

type RebindSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
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

type DescribeEngineDefaultParametersRequest struct {
    *ksyunhttp.BaseRequest
    EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
    ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *DescribeEngineDefaultParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEngineDefaultParametersRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEngineDefaultParametersRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineDefaultParametersResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		Parameters []struct {
					name *string `json:"name"`
					min *string `json:"min"`
					max *string `json:"max"`
					visible *int `json:"visible"`
					restart_required *bool `json:"restart_required"`
					type *string `json:"type"`
					config_type *string `json:"config_type"`
					enums *string `json:"enums"`
					default *string `json:"default"`
			} `json:"Parameters"`
		} `json:"Data"`
}

func (r *DescribeEngineDefaultParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEngineDefaultParametersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Parameters []*ModifyDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
    ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *ModifyDBParameterGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBParameterGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBParameterGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *string `json:"Data" name:"Data"`
}

func (r *ModifyDBParameterGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBParameterGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *DescribeDBInstanceParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceParametersRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParametersRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		Configuration struct {
				description *string `json:"description"`
				id *string `json:"id"`
				name *string `json:"name"`
			Values struct {
				Compile_aggregate_expressions *string `json:"Compile_aggregate_expressions"`
				Enable_unaligned_array_join *string `json:"Enable_unaligned_array_join"`
				Max_memory_usage_for_user *string `json:"Max_memory_usage_for_user"`
				Enable_optimize_predicate_expression *string `json:"Enable_optimize_predicate_expression"`
				Prefer_global_in_and_join *string `json:"Prefer_global_in_and_join"`
				Max_bytes_in_join *string `json:"Max_bytes_in_join"`
				Join_use_nulls *string `json:"Join_use_nulls"`
				Join_default_strictness *string `json:"Join_default_strictness"`
				Max_rows_in_join *string `json:"Max_rows_in_join"`
				Distributed_aggregation_memory_efficient *string `json:"Distributed_aggregation_memory_efficient"`
				Join_on_disk_max_files_to_merge *string `json:"Join_on_disk_max_files_to_merge"`
				Union_default_mode *string `json:"Union_default_mode"`
				Optimize_move_to_prewhere *string `json:"Optimize_move_to_prewhere"`
				Max_rows_to_group_by *string `json:"Max_rows_to_group_by"`
				Optimize_throw_if_noop *string `json:"Optimize_throw_if_noop"`
				Join_any_take_last_row *string `json:"Join_any_take_last_row"`
				Max_memory_usage *string `json:"Max_memory_usage"`
				Min_count_to_compile_aggregate_expression *string `json:"Min_count_to_compile_aggregate_expression"`
				Max_block_size *string `json:"Max_block_size"`
				Min_count_to_compile_expression *string `json:"Min_count_to_compile_expression"`
				Mutations_sync *string `json:"Mutations_sync"`
				Max_execution_time *string `json:"Max_execution_time"`
				Partial_merge_join_rows_in_right_blocks *string `json:"Partial_merge_join_rows_in_right_blocks"`
				Max_bytes_before_external_group_by *string `json:"Max_bytes_before_external_group_by"`
				Compile_expressions *string `json:"Compile_expressions"`
				Distributed_product_mode *string `json:"Distributed_product_mode"`
				Max_insert_block_size *string `json:"Max_insert_block_size"`
				Totals_mode *string `json:"Totals_mode"`
				Join_algorithm *string `json:"Join_algorithm"`
				Optimize_aggregation_in_order *string `json:"Optimize_aggregation_in_order"`
				Join_overflow_mode *string `json:"Join_overflow_mode"`
				Use_uncompressed_cache *string `json:"Use_uncompressed_cache"`
				Allow_experimental_alter_materialized_view_structure *string `json:"Allow_experimental_alter_materialized_view_structure"`
				Insert_null_as_default *string `json:"Insert_null_as_default"`
				Optimize_read_in_order *string `json:"Optimize_read_in_order"`
				Max_bytes_before_external_sort *string `json:"Max_bytes_before_external_sort"`
				Any_join_distinct_right_table_keys *string `json:"Any_join_distinct_right_table_keys"`
				Group_by_overflow_mode *string `json:"Group_by_overflow_mode"`
				Optimize_move_to_prewhere_if_final *string `json:"Optimize_move_to_prewhere_if_final"`
				Partial_merge_join_optimizations *string `json:"Partial_merge_join_optimizations"`
			} `json:"Values"`
				datastore_version_id *string `json:"datastore_version_id"`
		} `json:"Configuration"`
	} `json:"Data"`
}

func (r *DescribeDBInstanceParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceParametersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *ResetDBParameterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetDBParameterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetDBParameterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		Configuration struct {
				description *string `json:"description"`
				id *string `json:"id"`
				name *string `json:"name"`
			Values struct {
				Compile_aggregate_expressions *string `json:"Compile_aggregate_expressions"`
				Enable_unaligned_array_join *string `json:"Enable_unaligned_array_join"`
				Max_memory_usage_for_user *string `json:"Max_memory_usage_for_user"`
				Enable_optimize_predicate_expression *string `json:"Enable_optimize_predicate_expression"`
				Prefer_global_in_and_join *string `json:"Prefer_global_in_and_join"`
				Max_bytes_in_join *string `json:"Max_bytes_in_join"`
				Join_use_nulls *string `json:"Join_use_nulls"`
				Join_default_strictness *string `json:"Join_default_strictness"`
				Max_rows_in_join *string `json:"Max_rows_in_join"`
				Distributed_aggregation_memory_efficient *string `json:"Distributed_aggregation_memory_efficient"`
				Join_on_disk_max_files_to_merge *string `json:"Join_on_disk_max_files_to_merge"`
				Union_default_mode *string `json:"Union_default_mode"`
				Optimize_move_to_prewhere *string `json:"Optimize_move_to_prewhere"`
				Max_rows_to_group_by *string `json:"Max_rows_to_group_by"`
				Optimize_throw_if_noop *string `json:"Optimize_throw_if_noop"`
				Join_any_take_last_row *string `json:"Join_any_take_last_row"`
				Max_memory_usage *string `json:"Max_memory_usage"`
				Min_count_to_compile_aggregate_expression *string `json:"Min_count_to_compile_aggregate_expression"`
				Max_block_size *string `json:"Max_block_size"`
				Min_count_to_compile_expression *string `json:"Min_count_to_compile_expression"`
				Mutations_sync *string `json:"Mutations_sync"`
				Max_execution_time *string `json:"Max_execution_time"`
				Partial_merge_join_rows_in_right_blocks *string `json:"Partial_merge_join_rows_in_right_blocks"`
				Max_bytes_before_external_group_by *string `json:"Max_bytes_before_external_group_by"`
				Compile_expressions *string `json:"Compile_expressions"`
				Distributed_product_mode *string `json:"Distributed_product_mode"`
				Max_insert_block_size *string `json:"Max_insert_block_size"`
				Totals_mode *string `json:"Totals_mode"`
				Join_algorithm *string `json:"Join_algorithm"`
				Optimize_aggregation_in_order *string `json:"Optimize_aggregation_in_order"`
				Join_overflow_mode *string `json:"Join_overflow_mode"`
				Use_uncompressed_cache *string `json:"Use_uncompressed_cache"`
				Allow_experimental_alter_materialized_view_structure *string `json:"Allow_experimental_alter_materialized_view_structure"`
				Insert_null_as_default *string `json:"Insert_null_as_default"`
				Optimize_read_in_order *string `json:"Optimize_read_in_order"`
				Max_bytes_before_external_sort *string `json:"Max_bytes_before_external_sort"`
				Any_join_distinct_right_table_keys *string `json:"Any_join_distinct_right_table_keys"`
				Group_by_overflow_mode *string `json:"Group_by_overflow_mode"`
				Optimize_move_to_prewhere_if_final *string `json:"Optimize_move_to_prewhere_if_final"`
				Partial_merge_join_optimizations *string `json:"Partial_merge_join_optimizations"`
			} `json:"Values"`
				datastore_version_id *string `json:"datastore_version_id"`
		} `json:"Configuration"`
	} `json:"Data"`
}

func (r *ResetDBParameterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetDBParameterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

