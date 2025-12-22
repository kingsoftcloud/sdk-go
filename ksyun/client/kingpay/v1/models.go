package v1

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type QueryCashWalletActionRequest struct {
	*ksyunhttp.BaseRequest
	Subject *int `json:"subject,omitempty" name:"subject"`
}

func (r *QueryCashWalletActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryCashWalletActionResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"request_id" name:"request_id"`
	Data      struct {
		CustomerId      *int    `json:"CustomerId" name:"CustomerId"`
		AvailableAmount *string `json:"AvailableAmount" name:"AvailableAmount"`
		RewardAmount    *string `json:"RewardAmount" name:"RewardAmount"`
		FrozenAmount    *string `json:"FrozenAmount" name:"FrozenAmount"`
		Currency        *string `json:"Currency" name:"Currency"`
	} `json:"Data"`
}

func (r *QueryCashWalletActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCashWalletActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
