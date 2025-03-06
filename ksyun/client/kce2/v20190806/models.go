package v20190806

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type QueryPodsByInformerRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	PodName   *string `json:"PodName,omitempty" name:"PodName"`
	NodeIp    *string `json:"NodeIp,omitempty" name:"NodeIp"`
	PodIp     *string `json:"PodIp,omitempty" name:"PodIp"`
	Labels    *string `json:"Labels,omitempty" name:"Labels"`
	Page      *int    `json:"Page,omitempty" name:"Page"`
	PageSize  *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *QueryPodsByInformerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPodsByInformerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryPodsByInformerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryPodsByInformerResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *QueryPodsByInformerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPodsByInformerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
