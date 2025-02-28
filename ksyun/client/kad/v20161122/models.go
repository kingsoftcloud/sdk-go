package v20161122

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateForwardConfRequest struct {
	*ksyunhttp.BaseRequest
	KadId       *string `json:"KadId,omitempty" name:"KadId"`
	Protocol    *string `json:"Protocol,omitempty" name:"Protocol"`
	ServicePort *int    `json:"ServicePort,omitempty" name:"ServicePort"`
}

func (r *CreateForwardConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateForwardConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateForwardConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateForwardConfResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateForwardConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateForwardConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteForwardConfRequest struct {
	*ksyunhttp.BaseRequest
	ForwardConfId *string `json:"ForwardConfId,omitempty" name:"ForwardConfId"`
}

func (r *DeleteForwardConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteForwardConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteForwardConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteForwardConfResponse struct {
	*ksyunhttp.BaseResponse
	Response struct {
		RequestId      *string `json:"RequestId"`
		ForwardConfSet struct {
			Item []struct {
				ForwardConfId *string `json:"ForwardConfId"`
				Return        *string `json:"Return"`
				Message       *string `json:"Message"`
			} `json:"Item"`
		} `json:"ForwardConfSet"`
	} `json:"Response"`
}

func (r *DeleteForwardConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteForwardConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeForwardConfRequest struct {
	*ksyunhttp.BaseRequest
	KadId         *string   `json:"KadId,omitempty" name:"KadId"`
	ForwardConfId []*string `json:"ForwardConfId,omitempty" name:"ForwardConfId"`
}

func (r *DescribeForwardConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeForwardConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeForwardConfResponse struct {
	*ksyunhttp.BaseResponse
	Response struct {
		RequestId      *string `json:"RequestId"`
		ForwardConfSet struct {
			Item struct {
				KadId         *string `json:"KadId"`
				ForwardConfId *string `json:"ForwardConfId"`
				ServicePort   *string `json:"ServicePort"`
				Cname         *string `json:"Cname"`
				Protocol      *string `json:"Protocol"`
				SourceCount   *string `json:"SourceCount"`
				HealthMonitor struct {
					Switch *string `json:"Switch"`
					Rise   *string `json:"Rise"`
					Fall   *string `json:"Fall"`
					Delay  *string `json:"Delay"`
				} `json:"HealthMonitor"`
			} `json:"Item"`
		} `json:"ForwardConfSet"`
	} `json:"Response"`
}

func (r *DescribeForwardConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateForwardSourceRequest struct {
	*ksyunhttp.BaseRequest
	ForwardConfId *string `json:"ForwardConfId,omitempty" name:"ForwardConfId"`
	SourceIp      *string `json:"SourceIp,omitempty" name:"SourceIp"`
	SourcePort    *string `json:"SourcePort,omitempty" name:"SourcePort"`
}

func (r *CreateForwardSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateForwardSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateForwardSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateForwardSourceResponse struct {
	*ksyunhttp.BaseResponse
	Response struct {
		RequestId     *string `json:"RequestId"`
		ForwardSource struct {
			ForwardConfId   *string `json:"ForwardConfId"`
			ForwardSourceId *string `json:"ForwardSourceId"`
			SourceIp        *string `json:"SourceIp"`
			SourcePort      *string `json:"SourcePort"`
		} `json:"ForwardSource"`
	} `json:"Response"`
}

func (r *CreateForwardSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateForwardSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteForwardSourceRequest struct {
	*ksyunhttp.BaseRequest
	ForwardSourceId *string `json:"ForwardSourceId,omitempty" name:"ForwardSourceId"`
}

func (r *DeleteForwardSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteForwardSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteForwardSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteForwardSourceResponse struct {
	*ksyunhttp.BaseResponse
	DeleteForwardSourceResponse struct {
		RequestId *string `json:"RequestId"`
		Return    *string `json:"Return"`
	} `json:"DeleteForwardSourceResponse"`
}

func (r *DeleteForwardSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteForwardSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeForwardSourceRequest struct {
	*ksyunhttp.BaseRequest
	ForwardConfId   *string   `json:"ForwardConfId,omitempty" name:"ForwardConfId"`
	ForwardSourceId []*string `json:"ForwardSourceId,omitempty" name:"ForwardSourceId"`
}

func (r *DescribeForwardSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeForwardSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeForwardSourceResponse struct {
	*ksyunhttp.BaseResponse
	DescribeForwardSourceResponse struct {
		RequestId        *string `json:"RequestId"`
		ForwardSourceSet struct {
			Item struct {
				ForwardConfId          *string `json:"ForwardConfId"`
				ForwardSourceId        *string `json:"ForwardSourceId"`
				SourceIp               *string `json:"SourceIp"`
				SourcePort             *string `json:"SourcePort"`
				AutoReplace            *string `json:"AutoReplace"`
				RsRegion               *string `json:"RsRegion"`
				HealthMonitorStatusSet struct {
					Item struct {
						Ip       *string `json:"Ip"`
						Region   *string `json:"Region"`
						Status   *string `json:"Status"`
						LinkType *string `json:"LinkType"`
					} `json:"Item"`
				} `json:"HealthMonitorStatusSet"`
			} `json:"Item"`
		} `json:"ForwardSourceSet"`
	} `json:"DescribeForwardSourceResponse"`
}

func (r *DescribeForwardSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
