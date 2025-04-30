package v20200101
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeKeadRequest struct {
	*ksyunhttp.BaseRequest
	KeadId    []*string `json:"KeadId,omitempty" name:"KeadId"`
	ProjectId []*int    `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeKeadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeKeadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeadResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeKeadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeKeadIpRequest struct {
	*ksyunhttp.BaseRequest
	Ip        *string `json:"Ip,omitempty" name:"Ip"`
	ProjectId []*int  `json:"ProjectId,omitempty" name:"ProjectId"`
	PageSize  *int    `json:"PageSize,omitempty" name:"PageSize"`
	OffSet    *int    `json:"OffSet,omitempty" name:"OffSet"`
}

func (r *DescribeKeadIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeadIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeKeadIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeadIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeKeadIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeadIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeBlockIpRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeBlockIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBlockIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	BlockIpSet []struct {
		Ip           *string `json:"Ip" name:"Ip"`
		InBps        *string `json:"InBps" name:"InBps"`
		AttackType   *string `json:"AttackType" name:"AttackType"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		RegionCode   *string `json:"RegionCode" name:"RegionCode"`
		RegionName   *string `json:"RegionName" name:"RegionName"`
		RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
		Status       *string `json:"Status" name:"Status"`
		BlockTime    *string `json:"BlockTime" name:"BlockTime"`
		UnblockTime  *string `json:"UnblockTime" name:"UnblockTime"`
		PlanUnblockTime *string `json:"PlanUnblockTime" name:"PlanUnblockTime"`
	} `json:"BlockIpSet"`
	RegionSet []struct {
		RegionCode *string `json:"RegionCode" name:"RegionCode"`
		RegionName *string `json:"RegionName" name:"RegionName"`
		RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
	} `json:"RegionSet"`
	ResourceSet  []*string `json:"ResourceSet" name:"ResourceSet"`
	BlockIpCount *string   `json:"BlockIpCount" name:"BlockIpCount"`
}

func (r *DescribeBlockIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

