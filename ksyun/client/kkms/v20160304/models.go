package v20160304
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreateKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyName     *string `json:"KeyName,omitempty" name:"KeyName"`
	Description *string `json:"Description,omitempty" name:"Description"`
	KeyUsage    *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	Origin      *string `json:"Origin,omitempty" name:"Origin"`
	ChargeType  *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyId       *string `json:"KeyId,omitempty" name:"KeyId"`
	KeyName     *string `json:"KeyName,omitempty" name:"KeyName"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyKeyStateRequest struct {
	*ksyunhttp.BaseRequest
	KeyId    *string `json:"KeyId,omitempty" name:"KeyId"`
	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`
}

func (r *ModifyKeyStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyKeyStateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyKeyStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKeyStateResponse) FromJsonString(s string) error {
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


type DescribeKeysRequest struct {
	*ksyunhttp.BaseRequest
	KeyId []*string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKeysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type EncryptRequest struct {
	*ksyunhttp.BaseRequest
	KeyId     *string `json:"KeyId,omitempty" name:"KeyId"`
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
}

func (r *EncryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EncryptResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CiphertextBlob *string `json:"CiphertextBlob" name:"CiphertextBlob"`
}

func (r *EncryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DecryptRequest struct {
	*ksyunhttp.BaseRequest
	KeyId          *string `json:"KeyId,omitempty" name:"KeyId"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`
}

func (r *DecryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DecryptResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	KeyId     *string `json:"KeyId" name:"KeyId"`
}

func (r *DecryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GenerateDataKeyRequest struct {
	*ksyunhttp.BaseRequest
	KeyId         *string `json:"KeyId,omitempty" name:"KeyId"`
	KeySpec       *string `json:"KeySpec,omitempty" name:"KeySpec"`
	NumberOfBytes *int    `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`
}

func (r *GenerateDataKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GenerateDataKeyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	KeyId     *string `json:"KeyId" name:"KeyId"`
	Plaintext *string `json:"Plaintext" name:"Plaintext"`
}

func (r *GenerateDataKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

