package v20220101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type CreateAlarmPolicyTriggerRules struct {
    Period *string `json:"Period,omitempty" name:"Period"`
    Method *string `json:"Method,omitempty" name:"Method"`
    Compare *string `json:"Compare,omitempty" name:"Compare"`
    TriggerValue *string `json:"TriggerValue,omitempty" name:"TriggerValue"`
    ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
    ItemKey *string `json:"ItemKey,omitempty" name:"ItemKey"`
    Units *string `json:"Units,omitempty" name:"Units"`
    Points *int `json:"Points,omitempty" name:"Points"`
    EffectBT *string `json:"EffectBT,omitempty" name:"EffectBT"`
    EffectET *string `json:"EffectET,omitempty" name:"EffectET"`
    Interval *int `json:"Interval,omitempty" name:"Interval"`
    MaxCount *int `json:"MaxCount,omitempty" name:"MaxCount"`
}

type CreateAlarmPolicyUserNotice struct {
    ContactWay *int `json:"ContactWay,omitempty" name:"ContactWay"`
    ContactFlag *int `json:"ContactFlag,omitempty" name:"ContactFlag"`
    ContactId *int `json:"ContactId,omitempty" name:"ContactId"`
}


type CreateAlarmPolicyRequest struct {
    *ksyunhttp.BaseRequest
    PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
    ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
    PolicyType *int `json:"PolicyType,omitempty" name:"PolicyType"`
    ResourceBindType *int `json:"ResourceBindType,omitempty" name:"ResourceBindType"`
    ProjectId *int `json:"ProjectId,omitempty" name:"ProjectId"`
    InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
    TriggerRules []*CreateAlarmPolicyTriggerRules `json:"TriggerRules,omitempty" name:"TriggerRules"`
    UserNotice []*CreateAlarmPolicyUserNotice `json:"UserNotice,omitempty" name:"UserNotice"`
    URLNotice []*string `json:"URLNotice,omitempty" name:"URLNotice"`
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmPolicyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAlarmPolicyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		PolicyId *int `json:"PolicyId"`
	} `json:"Data"`
}

func (r *CreateAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyRequest struct {
    *ksyunhttp.BaseRequest
    PolicyIds []*int `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlarmPolicyRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyResponse struct {
    *ksyunhttp.BaseResponse
    requestId *string `json:"requestId" name:"requestId"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

