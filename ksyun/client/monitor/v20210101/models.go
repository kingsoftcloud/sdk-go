package v20210101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type ListAlarmPolicyRequest struct {
    *ksyunhttp.BaseRequest
    PageIndex *int `json:"PageIndex,omitempty" name:"PageIndex"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlarmPolicyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAlarmPolicyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListAlarmPolicyResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
		AlarmPolicyList []struct {
					policyId *int `json:"policyId"`
					policyName *string `json:"policyName"`
					productType *string `json:"productType"`
					policyType *string `json:"policyType"`
					enabled *string `json:"enabled"`
					instanceInfoCount *int `json:"instanceInfoCount"`
					triggerRuleCount *int `json:"triggerRuleCount"`
					contactInfoCount *int `json:"contactInfoCount"`
					callbackUrl *string `json:"callbackUrl"`
			} `json:"AlarmPolicyList"`
		} `json:"Data"`
    totalCount *int `json:"totalCount" name:"totalCount"`
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *ListAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyRequest struct {
    *ksyunhttp.BaseRequest
    PolicyId *int `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPolicyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlarmPolicyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
		AlarmPolicyList []struct {
					policyId *int `json:"policyId"`
					policyName *string `json:"policyName"`
					productType *int `json:"productType"`
					policyType *int `json:"policyType"`
					enabled *int `json:"enabled"`
					InstanceInfoCount *int `json:"InstanceInfoCount"`
				TriggerRuleList []struct {
					TriggerId *int `json:"TriggerId"`
					Period *string `json:"Period"`
					Method *string `json:"Method"`
					Compare *string `json:"Compare"`
					TriggerValue *string `json:"TriggerValue"`
					ItemName *string `json:"ItemName"`
					ItemKey *string `json:"ItemKey"`
					Units *string `json:"Units"`
					EffectBT *string `json:"EffectBT"`
					EffectET *string `json:"EffectET"`
					Tags *string `json:"Tags"`
					Interval *string `json:"Interval"`
					Points *int `json:"Points"`
					MaxCount *int `json:"MaxCount"`
				} `json:"TriggerRuleList"`
				ContactInfoList []struct {
					ContactId *int `json:"ContactId"`
					ContactName *string `json:"ContactName"`
					ContactWay *int `json:"ContactWay"`
					ContactFlag *int `json:"ContactFlag"`
				} `json:"ContactInfoList"`
					callbackUrl *string `json:"callbackUrl"`
			} `json:"AlarmPolicyList"`
		} `json:"Data"`
    totalCount *int `json:"totalCount" name:"totalCount"`
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *DescribeAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyObjectRequest struct {
    *ksyunhttp.BaseRequest
    PolicyId *int `json:"PolicyId,omitempty" name:"PolicyId"`
    PageIndex *int `json:"PageIndex,omitempty" name:"PageIndex"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribePolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyObjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePolicyObjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyObjectResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		ProductType *int `json:"ProductType"`
		InstanceInfoList []struct {
					instanceId *string `json:"instanceId"`
					ip *string `json:"ip"`
					instanceName *string `json:"instanceName"`
			} `json:"InstanceInfoList"`
		} `json:"Data"`
    totalCount *int `json:"totalCount" name:"totalCount"`
}

func (r *DescribePolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmReceivesRequest struct {
    *ksyunhttp.BaseRequest
    PolicyId *int `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmReceivesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmReceivesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlarmReceivesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmReceivesResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
	} `json:"Data"`
	ContactInfoList []struct {
		ContactId *int `json:"ContactId"`
		ContactName *string `json:"ContactName"`
		ContactWay *int `json:"ContactWay"`
		ContactFlag *int `json:"ContactFlag"`
	} `json:"ContactInfoList"`
    totalCount *int `json:"totalCount" name:"totalCount"`
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *DescribeAlarmReceivesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmReceivesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddAlarmReceivesRequest struct {
    *ksyunhttp.BaseRequest
    PolicyId *int `json:"PolicyId,omitempty" name:"PolicyId"`
    ContactFlag *int `json:"ContactFlag,omitempty" name:"ContactFlag"`
    ContactWay *int `json:"ContactWay,omitempty" name:"ContactWay"`
    ContactId []*int `json:"ContactId,omitempty" name:"ContactId"`
}

func (r *AddAlarmReceivesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddAlarmReceivesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddAlarmReceivesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddAlarmReceivesResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *AddAlarmReceivesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddAlarmReceivesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmReceivesRequest struct {
    *ksyunhttp.BaseRequest
    PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
    ContactFlag *int `json:"ContactFlag,omitempty" name:"ContactFlag"`
    ContactId []*int `json:"ContactId,omitempty" name:"ContactId"`
}

func (r *DeleteAlarmReceivesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmReceivesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlarmReceivesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmReceivesResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *DeleteAlarmReceivesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmReceivesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserGroupRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserGroupRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUserGroupRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetUserGroupResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
		UserGrpList []struct {
					userGrpId *int `json:"userGrpId"`
					name *string `json:"name"`
					userCount *int `json:"userCount"`
			} `json:"UserGrpList"`
		} `json:"Data"`
    totalCount *int `json:"totalCount" name:"totalCount"`
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *GetUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAlertUserRequest struct {
    *ksyunhttp.BaseRequest
    UserGrpId []*int `json:"UserGrpId,omitempty" name:"UserGrpId"`
}

func (r *GetAlertUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertUserRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetAlertUserRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetAlertUserResponse struct {
    *ksyunhttp.BaseResponse
	Data struct {
		UserList struct {
				userId *int `json:"userId"`
				userName *string `json:"userName"`
				userEmail *string `json:"userEmail"`
				userPhone *string `json:"userPhone"`
			UserGroups []struct {
				UserGrpId *int `json:"UserGrpId"`
				GroupName *string `json:"GroupName"`
			} `json:"UserGroups"`
				userStatus *int `json:"userStatus"`
		} `json:"UserList"`
	} `json:"Data"`
    totalCount *int `json:"totalCount" name:"totalCount"`
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *GetAlertUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

