package v20220101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateAlarmPolicyRequest struct {
	*ksyunhttp.BaseRequest
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
	Data      struct {
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
