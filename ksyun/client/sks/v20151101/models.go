package v20151101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeKeysFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateKeyRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *CreateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	PrivateKey *string `json:"PrivateKey" name:"PrivateKey"`
}

func (r *CreateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyName     *string `json:"KeyName,omitempty" name:"KeyName"`
	PublicKey   *string `json:"PublicKey,omitempty" name:"PublicKey"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ImportKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ImportKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ImportKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	KeyId   *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *ModifyKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults *int                  `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string               `json:"NextToken,omitempty" name:"NextToken"`
	KeyId      []*string             `json:"KeyId,omitempty" name:"KeyId"`
	Filter     []*DescribeKeysFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	KeySet     []struct {
		KeyId      *string `json:"KeyId"`
		PublicKey  *string `json:"PublicKey"`
		CreateTime *string `json:"CreateTime"`
		KeyName    *string `json:"KeyName"`
	} `json:"KeySet"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
