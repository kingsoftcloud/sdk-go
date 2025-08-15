package v20240117

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetDwsuMetricRequest struct {
	*ksyunhttp.BaseRequest
	DwsuId    *string `json:"DwsuId,omitempty" name:"DwsuId"`
	Timestamp *int    `json:"Timestamp,omitempty" name:"Timestamp"`
}

func (r *GetDwsuMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDwsuMetricResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	Data    []struct {
		Name       *string `json:"Name" name:"Name"`
		ResourceId *string `json:"ResourceId" name:"ResourceId"`
		Value      *string `json:"Value" name:"Value"`
	} `json:"Data"`
}

func (r *GetDwsuMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDwsuMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
