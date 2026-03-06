package v20190401

import (
	"encoding/json"

	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type ListOperateLogsRequest struct {
	*ksyunhttp.BaseRequest
	EventName      *string `json:"EventName,omitempty" name:"EventName"`
	EventRw        *string `json:"EventRw,omitempty" name:"EventRw"`
	ServiceName    *string `json:"ServiceName,omitempty" name:"ServiceName"`
	UserName       *string `json:"UserName,omitempty" name:"UserName"`
	AccessKey      *string `json:"AccessKey,omitempty" name:"AccessKey"`
	EventBeginDate *string `json:"EventBeginDate,omitempty" name:"EventBeginDate"`
	EventEndDate   *string `json:"EventEndDate,omitempty" name:"EventEndDate"`
	ResourceType   *string `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceName   *string `json:"ResourceName,omitempty" name:"ResourceName"`
	Page           *string `json:"Page,omitempty" name:"Page"`
	PageSize       *string `json:"PageSize,omitempty" name:"PageSize"`
	SearchAfter    *string `json:"SearchAfter,omitempty" name:"SearchAfter"`
	Ip             *string `json:"Ip,omitempty" name:"Ip"`
	EventType      *string `json:"EventType,omitempty" name:"EventType"`
	Pre            *int    `json:"Pre,omitempty" name:"Pre"`
	RequestId      *string `json:"RequestId,omitempty" name:"RequestId"`
}

func (r *ListOperateLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListOperateLogsResponse struct {
	*ksyunhttp.BaseResponse
	Total  *int64 `json:"Total" name:"Total"`
	Events []struct {
		ErrorMessage      *string      `json:"ErrorMessage" name:"ErrorMessage"`
		CreateTime        *string      `json:"CreateTime" name:"CreateTime"`
		ServiceName       *string      `json:"ServiceName" name:"ServiceName"`
		EventSource       *string      `json:"EventSource" name:"EventSource"`
		ApiVersion        *string      `json:"ApiVersion" name:"ApiVersion"`
		RequestParameters *interface{} `json:"RequestParameters" name:"RequestParameters"`
		SourceIpAddress   *string      `json:"SourceIpAddress" name:"SourceIpAddress"`
		EventVersion      *string      `json:"EventVersion" name:"EventVersion"`
		EventType         *string      `json:"EventType" name:"EventType"`
		EventId           *string      `json:"EventId" name:"EventId"`
		EventRw           *string      `json:"EventRw" name:"EventRw"`
		EventName         *string      `json:"EventName" name:"EventName"`
		UserIdentity      struct {
			RoleName  *string                  `json:"RoleName" name:"RoleName"`
			AccountId *ksyunhttp.StringOrInt64 `json:"AccountId" name:"AccountId"`
			UserType  *string                  `json:"UserType" name:"UserType"`
			UserName  *string                  `json:"UserName" name:"UserName"`
			AccessKey *string                  `json:"AccessKey" name:"AccessKey"`
		} `json:"UserIdentity" name:"UserIdentity"`
		ErrorCode    *string `json:"ErrorCode" name:"ErrorCode"`
		Region       *string `json:"Region" name:"Region"`
		RequestId    *string `json:"RequestId" name:"RequestId"`
		EventTime    *string `json:"EventTime" name:"EventTime"`
		UserAgent    *string `json:"UserAgent" name:"UserAgent"`
		ResourceType *string `json:"ResourceType" name:"ResourceType"`
		ResourceName *string `json:"ResourceName" name:"ResourceName"`
		RegionCn     *string `json:"RegionCn" name:"RegionCn"`
	} `json:"Events"`
	SearchAfter *string `json:"SearchAfter" name:"SearchAfter"`
	RequestId   *string `json:"RequestId" name:"RequestId"`
}

func (r *ListOperateLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
