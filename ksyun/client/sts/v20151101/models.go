package v20151101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type AssumeRoleRequest struct {
	*ksyunhttp.BaseRequest
	RoleKrn         *string `json:"RoleKrn,omitempty" name:"RoleKrn"`
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`
	DurationSeconds *string `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
	Policy          *string `json:"Policy,omitempty" name:"Policy"`
}

func (r *AssumeRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssumeRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssumeRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssumeRoleResponse struct {
	*ksyunhttp.BaseResponse
	AssumeRoleResult struct {
		Credentials struct {
			Expiration      *string `json:"Expiration" name:"Expiration"`
			SecretAccessKey *string `json:"SecretAccessKey" name:"SecretAccessKey"`
			AccessKeyId     *string `json:"AccessKeyId" name:"AccessKeyId"`
			SecurityToken   *string `json:"SecurityToken" name:"SecurityToken"`
		} `json:"Credentials" name:"Credentials"`
		AssumedRoleUser struct {
			Krn           *string `json:"Krn" name:"Krn"`
			AssumedRoleId *string `json:"AssumedRoleId" name:"AssumedRoleId"`
		} `json:"AssumedRoleUser" name:"AssumedRoleUser"`
	} `json:"AssumeRoleResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AssumeRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssumeRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
