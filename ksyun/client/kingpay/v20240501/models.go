package v20240501
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type QueryCashWalletActionRequest struct {
    *ksyunhttp.BaseRequest
    Subject *int `json:"subject,omitempty" name:"subject"`
}

func (r *QueryCashWalletActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryCashWalletActionRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryCashWalletActionRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type QueryCashWalletActionResponse struct {
    *ksyunhttp.BaseResponse
    status *int `json:"status" name:"status"`
    request_id *string `json:"request_id" name:"request_id"`
	Data struct {
		CustomerId *int `json:"CustomerId"`
		AvailableAmount *string `json:"AvailableAmount"`
		RewardAmount *string `json:"RewardAmount"`
		FrozenAmount *string `json:"FrozenAmount"`
		Currency *string `json:"Currency"`
	} `json:"Data"`
}

func (r *QueryCashWalletActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryCashWalletActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

