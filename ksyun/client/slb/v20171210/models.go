package v20171210
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeLoadBalancersRequest struct {
    *ksyunhttp.BaseRequest
    LoadBalancerId []*string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
    State *string `json:"State,omitempty" name:"State"`
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    Filter []*string `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    LoadBalancerDescriptions *string `json:"LoadBalancerDescriptions" name:"LoadBalancerDescriptions"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

