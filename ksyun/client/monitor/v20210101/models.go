package v20210101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type ListAlarmPolicyRequest struct {
	*ksyunhttp.BaseRequest
	PageIndex *int `json:"PageIndex,omitempty" name:"PageIndex"`
	PageSize  *int `json:"PageSize,omitempty" name:"PageSize"`
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
			PolicyId          *int    `json:"policyId" name:"policyId"`
			PolicyName        *string `json:"policyName" name:"policyName"`
			ProductType       *string `json:"productType" name:"productType"`
			PolicyType        *string `json:"policyType" name:"policyType"`
			Enabled           *string `json:"enabled" name:"enabled"`
			InstanceInfoCount *int    `json:"instanceInfoCount" name:"instanceInfoCount"`
			TriggerRuleCount  *int    `json:"triggerRuleCount" name:"triggerRuleCount"`
			ContactInfoCount  *int    `json:"contactInfoCount" name:"contactInfoCount"`
			CallbackUrl       *string `json:"callbackUrl" name:"callbackUrl"`
		} `json:"AlarmPolicyList" name:"AlarmPolicyList"`
	} `json:"Data"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"requestId" name:"requestId"`
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
			PolicyId          *int    `json:"policyId" name:"policyId"`
			PolicyName        *string `json:"policyName" name:"policyName"`
			ProductType       *int    `json:"productType" name:"productType"`
			PolicyType        *int    `json:"policyType" name:"policyType"`
			Enabled           *int    `json:"enabled" name:"enabled"`
			InstanceInfoCount *int    `json:"InstanceInfoCount" name:"InstanceInfoCount"`
			TriggerRuleList   []struct {
				TriggerId    *int    `json:"TriggerId" name:"TriggerId"`
				Period       *string `json:"Period" name:"Period"`
				Method       *string `json:"Method" name:"Method"`
				Compare      *string `json:"Compare" name:"Compare"`
				TriggerValue *string `json:"TriggerValue" name:"TriggerValue"`
				ItemName     *string `json:"ItemName" name:"ItemName"`
				ItemKey      *string `json:"ItemKey" name:"ItemKey"`
				Units        *string `json:"Units" name:"Units"`
				EffectBT     *string `json:"EffectBT" name:"EffectBT"`
				EffectET     *string `json:"EffectET" name:"EffectET"`
				Tags         *string `json:"Tags" name:"Tags"`
				Interval     *string `json:"Interval" name:"Interval"`
				Points       *int    `json:"Points" name:"Points"`
				MaxCount     *int    `json:"MaxCount" name:"MaxCount"`
			} `json:"TriggerRuleList"`
			ContactInfoList []struct {
				ContactId   *int    `json:"ContactId" name:"ContactId"`
				ContactName *string `json:"ContactName" name:"ContactName"`
				ContactWay  *int    `json:"ContactWay" name:"ContactWay"`
				ContactFlag *int    `json:"ContactFlag" name:"ContactFlag"`
			} `json:"ContactInfoList"`
			CallbackUrl *string `json:"callbackUrl" name:"callbackUrl"`
		} `json:"AlarmPolicyList" name:"AlarmPolicyList"`
	} `json:"Data"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"requestId" name:"requestId"`
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
	PolicyId  *int `json:"PolicyId,omitempty" name:"PolicyId"`
	PageIndex *int `json:"PageIndex,omitempty" name:"PageIndex"`
	PageSize  *int `json:"PageSize,omitempty" name:"PageSize"`
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
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		ProductType      *int `json:"ProductType" name:"ProductType"`
		InstanceInfoList []struct {
			InstanceId   *string `json:"instanceId" name:"instanceId"`
			Ip           *string `json:"ip" name:"ip"`
			InstanceName *string `json:"instanceName" name:"instanceName"`
		} `json:"InstanceInfoList" name:"InstanceInfoList"`
	} `json:"Data"`
	TotalCount *int `json:"totalCount" name:"totalCount"`
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
		ContactId   *int    `json:"ContactId" name:"ContactId"`
		ContactName *string `json:"ContactName" name:"ContactName"`
		ContactWay  *int    `json:"ContactWay" name:"ContactWay"`
		ContactFlag *int    `json:"ContactFlag" name:"ContactFlag"`
	} `json:"ContactInfoList"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"requestId" name:"requestId"`
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
	PolicyId    *int   `json:"PolicyId,omitempty" name:"PolicyId"`
	ContactFlag *int   `json:"ContactFlag,omitempty" name:"ContactFlag"`
	ContactWay  *int   `json:"ContactWay,omitempty" name:"ContactWay"`
	ContactId   []*int `json:"ContactId,omitempty" name:"ContactId"`
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
	RequestId *string `json:"requestId" name:"requestId"`
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
	PolicyId    *int   `json:"PolicyId,omitempty" name:"PolicyId"`
	ContactFlag *int   `json:"ContactFlag,omitempty" name:"ContactFlag"`
	ContactId   []*int `json:"ContactId,omitempty" name:"ContactId"`
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
	RequestId *string `json:"requestId" name:"requestId"`
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
			UserGrpId *int    `json:"userGrpId" name:"userGrpId"`
			Name      *string `json:"name" name:"name"`
			UserCount *int    `json:"userCount" name:"userCount"`
		} `json:"UserGrpList" name:"UserGrpList"`
	} `json:"Data"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"requestId" name:"requestId"`
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
			UserId     *int    `json:"userId" name:"userId"`
			UserName   *string `json:"userName" name:"userName"`
			UserEmail  *string `json:"userEmail" name:"userEmail"`
			UserPhone  *string `json:"userPhone" name:"userPhone"`
			UserGroups []struct {
				UserGrpId *int    `json:"UserGrpId" name:"UserGrpId"`
				GroupName *string `json:"GroupName" name:"GroupName"`
			} `json:"UserGroups"`
			UserStatus *int `json:"userStatus" name:"userStatus"`
		} `json:"UserList" name:"UserList"`
	} `json:"Data"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"requestId" name:"requestId"`
}

func (r *GetAlertUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAlertUserStatusRequest struct {
	*ksyunhttp.BaseRequest
	UserId     []*int `json:"UserId,omitempty" name:"UserId"`
	UserStatus *int   `json:"UserStatus,omitempty" name:"UserStatus"`
}

func (r *UpdateAlertUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAlertUserStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateAlertUserStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAlertUserStatusResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		UserList []struct {
			UserId     *int    `json:"userId" name:"userId"`
			UserName   *string `json:"userName" name:"userName"`
			UserEmail  *string `json:"userEmail" name:"userEmail"`
			UserPhone  *string `json:"userPhone" name:"userPhone"`
			GroupName  *string `json:"groupName" name:"groupName"`
			UserGrpId  *int    `json:"userGrpId" name:"userGrpId"`
			UserStatus *int    `json:"userStatus" name:"userStatus"`
		} `json:"UserList" name:"UserList"`
	} `json:"Data"`
	TotalCount *int    `json:"totalCount" name:"totalCount"`
	RequestId  *string `json:"requestId" name:"requestId"`
}

func (r *UpdateAlertUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAlertUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSysEventGroupListRequest struct {
	*ksyunhttp.BaseRequest
	Namespace  *string   `json:"Namespace,omitempty" name:"Namespace"`
	StartTime  *int      `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *int      `json:"EndTime,omitempty" name:"EndTime"`
	EventName  *string   `json:"EventName,omitempty" name:"EventName"`
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	Order      *bool     `json:"Order,omitempty" name:"Order"`
	PageIndex  *int      `json:"PageIndex,omitempty" name:"PageIndex"`
	PageSize   *int      `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeSysEventGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSysEventGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSysEventGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSysEventGroupListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"requestId" name:"requestId"`
	Code              *string `json:"code" name:"code"`
	Message           *string `json:"message" name:"message"`
	SysEventGroupList []struct {
		InstanceId  *string `json:"InstanceId" name:"InstanceId"`
		GroupId     *string `json:"GroupId" name:"GroupId"`
		Producttype *int    `json:"Producttype" name:"Producttype"`
		EventName   *string `json:"EventName" name:"EventName"`
		Status      *string `json:"Status" name:"Status"`
		CreatedAt   *int    `json:"CreatedAt" name:"CreatedAt"`
		UpdateAt    *int    `json:"UpdateAt" name:"UpdateAt"`
		PlanAt      *int    `json:"PlanAt" name:"PlanAt"`
		FinishAt    *int    `json:"FinishAt" name:"FinishAt"`
		Deadline    *int    `json:"Deadline" name:"Deadline"`
	} `json:"SysEventGroupList"`
	TotalCount *int `json:"totalCount" name:"totalCount"`
}

func (r *DescribeSysEventGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSysEventGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeMonitorProductListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorProductListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMonitorProductListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeMonitorProductListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
