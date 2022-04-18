package v20200101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeBlockIpRequest struct {
    *ksyunhttp.BaseRequest
    SearchStr *string `json:"SearchStr,omitempty" name:"SearchStr"`
    Status *string `json:"Status,omitempty" name:"Status"`
    InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
    RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
    StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
    endTime *string `json:"endTime,omitempty" name:"endTime"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
	BlockIpSet []struct {
		Ip *string `json:"Ip"`
		InBps *string `json:"InBps"`
		AttackType *string `json:"AttackType"`
		InstanceType *string `json:"InstanceType"`
		InstanceName *string `json:"InstanceName"`
		RegionCode *string `json:"RegionCode"`
		RegionName *string `json:"RegionName"`
		RegionEnName *string `json:"RegionEnName"`
		Status *string `json:"Status"`
		BlockTime *string `json:"BlockTime"`
		UnblockTime *string `json:"UnblockTime"`
		PlanUnblockTime *string `json:"PlanUnblockTime"`
	} `json:"BlockIpSet"`
	RegionSet []struct {
		RegionCode *string `json:"RegionCode"`
		RegionName *string `json:"RegionName"`
		RegionEnName *string `json:"RegionEnName"`
	} `json:"RegionSet"`
    ResourceSet []*string `json:"ResourceSet" name:"ResourceSet"`
    BlockIpCount *string `json:"BlockIpCount" name:"BlockIpCount"`
}

func (r *DescribeBlockIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBlockIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

