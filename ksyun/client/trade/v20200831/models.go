package v20200831
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type SetRenewalRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds   *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	RenewStrategy *int    `json:"RenewStrategy,omitempty" name:"RenewStrategy"`
	RenewDuration *int    `json:"RenewDuration,omitempty" name:"RenewDuration"`
}

func (r *SetRenewalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetRenewalResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"success" name:"success"`
}

func (r *SetRenewalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRenewalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

