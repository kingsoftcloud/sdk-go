package v20200114
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		BillType     *int    `json:"BillType" name:"BillType"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceType *int    `json:"InstanceType" name:"InstanceType"`
		ServiceBeginTime *string `json:"ServiceBeginTime" name:"ServiceBeginTime"`
		Status       *int    `json:"Status" name:"Status"`
	} `json:"Data"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

