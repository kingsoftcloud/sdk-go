package v20200825
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
    SecurityGroupType *string `json:"SecurityGroupType,omitempty" name:"SecurityGroupType"`
    DBInstanceIdentifier []*string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
    SecurityGroupRuleSecurityGroupRuleNameN *string `json:"SecurityGroupRule.SecurityGroupRuleName.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleName.N"`
    SecurityGroupRuleSecurityGroupRuleCidrN *string `json:"SecurityGroupRule.SecurityGroupRuleCidr.N,omitempty" name:"SecurityGroupRule.SecurityGroupRuleCidr.N"`
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
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					SecurityGroupType *string `json:"SecurityGroupType"`
					Created *string `json:"Created"`
				Instances struct {
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleName *string `json:"SecurityGroupRuleName"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupIdN *string `json:"SecurityGroupId.N,omitempty" name:"SecurityGroupId.N"`
    SecurityGroupType *string `json:"SecurityGroupType,omitempty" name:"SecurityGroupType"`
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
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					SecurityGroupType *string `json:"SecurityGroupType"`
					Created *string `json:"Created"`
				Instances []struct {
					DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
					DBInstanceName *string `json:"DBInstanceName"`
					Vip *string `json:"Vip"`
					Created *string `json:"Created"`
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleName *string `json:"SecurityGroupRuleName"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					SecurityGroupType *string `json:"SecurityGroupType"`
					Created *string `json:"Created"`
				Instances []struct {
					DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
					DBInstanceName *string `json:"DBInstanceName"`
					Vip *string `json:"Vip"`
					Created *string `json:"Created"`
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
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
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					SecurityGroupType *string `json:"SecurityGroupType"`
					Created *string `json:"Created"`
				Instances []struct {
					DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
					DBInstanceName *string `json:"DBInstanceName"`
					Vip *string `json:"Vip"`
					Created *string `json:"Created"`
					DBInstanceType *string `json:"DBInstanceType"`
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleName *string `json:"SecurityGroupRuleName"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
    SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
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
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					SecurityGroupType *string `json:"SecurityGroupType"`
					Created *string `json:"Created"`
				Instances struct {
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupRuleAction *string `json:"SecurityGroupRuleAction,omitempty" name:"SecurityGroupRuleAction"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    SecurityGroupRuleSecurityGroupRuleId []*string `json:"SecurityGroupRule.SecurityGroupRuleId,omitempty" name:"SecurityGroupRule.SecurityGroupRuleId"`
    SecurityGroupRuleSecurityGroupRuleName []*string `json:"SecurityGroupRule.SecurityGroupRuleName,omitempty" name:"SecurityGroupRule.SecurityGroupRuleName"`
    SecurityGroupRuleSecurityGroupRuleCidr []*string `json:"SecurityGroupRule.SecurityGroupRuleCidr,omitempty" name:"SecurityGroupRule.SecurityGroupRuleCidr"`
}

func (r *ModifySecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupRuleRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					SecurityGroupType *string `json:"SecurityGroupType"`
					Created *string `json:"Created"`
				Instances []struct {
					DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
					DBInstanceName *string `json:"DBInstanceName"`
					Vip *string `json:"Vip"`
					Created *string `json:"Created"`
					DBInstanceType *string `json:"DBInstanceType"`
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleName *string `json:"SecurityGroupRuleName"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					SecurityGroupRuleCidr *string `json:"SecurityGroupRuleCidr"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupRelationRequest struct {
    *ksyunhttp.BaseRequest
    RelationAction *string `json:"RelationAction,omitempty" name:"RelationAction"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SecurityGroupRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SecurityGroupRelationRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SecurityGroupRelationRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupRelationResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
					SecurityGroupId *string `json:"SecurityGroupId"`
					SecurityGroupName *string `json:"SecurityGroupName"`
					SecurityGroupDescription *string `json:"SecurityGroupDescription"`
					Created *string `json:"Created"`
				Instances []struct {
				} `json:"Instances"`
				SecurityGroupRules []struct {
					SecurityGroupRuleId *string `json:"SecurityGroupRuleId"`
					SecurityGroupRuleName *string `json:"SecurityGroupRuleName"`
					SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
					Created *string `json:"Created"`
				} `json:"SecurityGroupRules"`
			} `json:"SecurityGroups"`
		} `json:"Data"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SecurityGroupRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SecurityGroupRelationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleNameRequest struct {
    *ksyunhttp.BaseRequest
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
    SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
    SecurityGroupRuleName *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}

func (r *ModifySecurityGroupRuleNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupRuleNameRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleNameRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleNameResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupRuleNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupRuleNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

