package v20250220
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type ListInstanceSupportBillTypesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *ListInstanceSupportBillTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListInstanceSupportBillTypesResponse struct {
	*ksyunhttp.BaseResponse
	Error struct {
		Code *string `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Error"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		BillTypeEnName *string `json:"BillTypeEnName" name:"BillTypeEnName"`
		BillTypeId   *int    `json:"BillTypeId" name:"BillTypeId"`
		BillTypeName *string `json:"BillTypeName" name:"BillTypeName"`
	} `json:"Data"`
}

func (r *ListInstanceSupportBillTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceSupportBillTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddTrialToBuyTaskRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId         *string `json:"instanceId,omitempty" name:"instanceId"`
	BillType           *int    `json:"billType,omitempty" name:"billType"`
	Duration           *int    `json:"duration,omitempty" name:"duration"`
	AutoTrialToBuyDate *string `json:"autoTrialToBuyDate,omitempty" name:"autoTrialToBuyDate"`
}

func (r *AddTrialToBuyTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddTrialToBuyTaskResponse struct {
	*ksyunhttp.BaseResponse
	Error struct {
		Code *string `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Error"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"success" name:"success"`
}

func (r *AddTrialToBuyTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddTrialToBuyTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteTrialToBuyTaskRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
}

func (r *DeleteTrialToBuyTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteTrialToBuyTaskResponse struct {
	*ksyunhttp.BaseResponse
	Error struct {
		Code *string `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Error"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"success" name:"success"`
}

func (r *DeleteTrialToBuyTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrialToBuyTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateTrialToBuyNowRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"instanceId,omitempty" name:"instanceId"`
	BillType   *int    `json:"billType,omitempty" name:"billType"`
	Duration   *int    `json:"duration,omitempty" name:"duration"`
}

func (r *CreateTrialToBuyNowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTrialToBuyNowResponse struct {
	*ksyunhttp.BaseResponse
	Error struct {
		Code *string `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Error"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"success" name:"success"`
}

func (r *CreateTrialToBuyNowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrialToBuyNowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

