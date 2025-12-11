package v20250101
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeDefaultParamsRequest struct {
	*ksyunhttp.BaseRequest
	DbVersion *float64 `json:"DbVersion,omitempty" name:"DbVersion"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDefaultParamsResponse struct {
	*ksyunhttp.BaseResponse
	DefaultParams []struct {
		DefaultValue    *string   `json:"DefaultValue" name:"DefaultValue"`
		Visible         *int      `json:"Visible" name:"Visible"`
		RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
		ParamName       *string   `json:"ParamName" name:"ParamName"`
		ParamType       *string   `json:"ParamType" name:"ParamType"`
		DataType        *string   `json:"DataType" name:"DataType"`
		Enums           []*string `json:"Enums" name:"Enums"`
	} `json:"DefaultParams"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

