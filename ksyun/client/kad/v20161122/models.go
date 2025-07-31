package v20161122

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
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

type DeleteForwardConfResponse struct {
	*ksyunhttp.BaseResponse
	Response struct {
		RequestId      *string `json:"RequestId" name:"RequestId"`
		ForwardConfSet struct {
			Item []struct {
				ForwardConfId *string `json:"ForwardConfId" name:"ForwardConfId"`
				Return        *string `json:"Return" name:"Return"`
				Message       *string `json:"Message" name:"Message"`
			} `json:"Item"`
		} `json:"ForwardConfSet" name:"ForwardConfSet"`
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

type DescribeForwardConfResponse struct {
	*ksyunhttp.BaseResponse
	Response struct {
		RequestId      *string `json:"RequestId" name:"RequestId"`
		ForwardConfSet struct {
			Item struct {
				KadId         *string `json:"KadId" name:"KadId"`
				ForwardConfId *string `json:"ForwardConfId" name:"ForwardConfId"`
				ServicePort   *string `json:"ServicePort" name:"ServicePort"`
				Cname         *string `json:"Cname" name:"Cname"`
				Protocol      *string `json:"Protocol" name:"Protocol"`
				SourceCount   *string `json:"SourceCount" name:"SourceCount"`
				HealthMonitor struct {
					Switch *string `json:"Switch" name:"Switch"`
					Rise   *string `json:"Rise" name:"Rise"`
					Fall   *string `json:"Fall" name:"Fall"`
					Delay  *string `json:"Delay" name:"Delay"`
				} `json:"HealthMonitor" name:"HealthMonitor"`
			} `json:"Item"`
		} `json:"ForwardConfSet" name:"ForwardConfSet"`
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

type CreateForwardSourceResponse struct {
	*ksyunhttp.BaseResponse
	Response struct {
		RequestId     *string `json:"RequestId" name:"RequestId"`
		ForwardSource struct {
			ForwardConfId   *string `json:"ForwardConfId" name:"ForwardConfId"`
			ForwardSourceId *string `json:"ForwardSourceId" name:"ForwardSourceId"`
			SourceIp        *string `json:"SourceIp" name:"SourceIp"`
			SourcePort      *string `json:"SourcePort" name:"SourcePort"`
		} `json:"ForwardSource" name:"ForwardSource"`
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

type DeleteForwardSourceResponse struct {
	*ksyunhttp.BaseResponse
	DeleteForwardSourceResponse struct {
		RequestId *string `json:"RequestId" name:"RequestId"`
		Return    *string `json:"Return" name:"Return"`
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

type DescribeForwardSourceResponse struct {
	*ksyunhttp.BaseResponse
	DescribeForwardSourceResponse struct {
		RequestId        *string `json:"RequestId" name:"RequestId"`
		ForwardSourceSet struct {
			Item struct {
				ForwardConfId          *string `json:"ForwardConfId" name:"ForwardConfId"`
				ForwardSourceId        *string `json:"ForwardSourceId" name:"ForwardSourceId"`
				SourceIp               *string `json:"SourceIp" name:"SourceIp"`
				SourcePort             *string `json:"SourcePort" name:"SourcePort"`
				AutoReplace            *string `json:"AutoReplace" name:"AutoReplace"`
				RsRegion               *string `json:"RsRegion" name:"RsRegion"`
				HealthMonitorStatusSet struct {
					Item struct {
						Ip       *string `json:"Ip" name:"Ip"`
						Region   *string `json:"Region" name:"Region"`
						Status   *string `json:"Status" name:"Status"`
						LinkType *string `json:"LinkType" name:"LinkType"`
					} `json:"Item"`
				} `json:"HealthMonitorStatusSet" name:"HealthMonitorStatusSet"`
			} `json:"Item"`
		} `json:"ForwardSourceSet" name:"ForwardSourceSet"`
	} `json:"DescribeForwardSourceResponse"`
}

func (r *DescribeForwardSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttackLogRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetAttackLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetAttackLogResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetAttackLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttackLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewRequest struct {
	*ksyunhttp.BaseRequest
	KadId     *string `json:"KadId,omitempty" name:"KadId"`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeOverviewResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
