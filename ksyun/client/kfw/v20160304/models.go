package v20160304
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeCfwAvRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeCfwAvRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCfwAvResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	CfwAvs    []struct {
		AvId       *string `json:"AvId" name:"AvId"`
		Protocol   *string `json:"Protocol" name:"Protocol"`
		ProtectType *string `json:"ProtectType" name:"ProtectType"`
		Status     *string `json:"Status" name:"Status"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAvs"`
}

func (r *DescribeCfwAvResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwAvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchCfwAddrbookRequest struct {
	*ksyunhttp.BaseRequest
	AddrbookIds   []*string `json:"AddrbookIds,omitempty" name:"AddrbookIds"`
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *DeleteBatchCfwAddrbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBatchCfwAddrbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteBatchCfwAddrbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchCfwAddrbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

