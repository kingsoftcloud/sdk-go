package v20220101

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateAlarmPolicyRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAlarmPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		PolicyId *int `json:"PolicyId" name:"PolicyId"`
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

type DeleteAlarmPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"requestId" name:"requestId"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
