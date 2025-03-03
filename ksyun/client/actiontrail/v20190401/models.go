package v20190401

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
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
}

func (r *ListOperateLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperateLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListOperateLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListOperateLogsResponse struct {
	*ksyunhttp.BaseResponse
	Total  *int `json:"Total" name:"Total"`
	Events []struct {
		ErrorMessage      *string `json:"ErrorMessage" name:"ErrorMessage"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		ServiceName       *string `json:"ServiceName" name:"ServiceName"`
		EventSource       *string `json:"EventSource" name:"EventSource"`
		ApiVersion        *string `json:"ApiVersion" name:"ApiVersion"`
		RequestParameters struct {
		} `json:"RequestParameters" name:"RequestParameters"`
		SourceIpAddress *string `json:"SourceIpAddress" name:"SourceIpAddress"`
		EventVersion    *string `json:"EventVersion" name:"EventVersion"`
		EventType       *string `json:"EventType" name:"EventType"`
		EventId         *string `json:"EventId" name:"EventId"`
		EventRw         *string `json:"EventRw" name:"EventRw"`
		EventName       *string `json:"EventName" name:"EventName"`
		UserIdentity    struct {
			RoleName  *string `json:"RoleName" name:"RoleName"`
			AccountId *string `json:"AccountId" name:"AccountId"`
			UserName  *string `json:"UserName" name:"UserName"`
			AccessKey *string `json:"AccessKey" name:"AccessKey"`
		} `json:"UserIdentity" name:"UserIdentity"`
		ReferencedResources []struct {
		} `json:"ReferencedResources" name:"ReferencedResources"`
		ErrorCode    *string `json:"ErrorCode" name:"ErrorCode"`
		Region       *string `json:"Region" name:"Region"`
		RequestId    *string `json:"RequestId" name:"RequestId"`
		EventTime    *string `json:"EventTime" name:"EventTime"`
		UserAgent    *string `json:"UserAgent" name:"UserAgent"`
		ResourceType *string `json:"ResourceType" name:"ResourceType"`
		ResourceName *string `json:"ResourceName" name:"ResourceName"`
		RegionCn     *string `json:"RegionCn" name:"RegionCn"`
	} `json:"Events"`
	SearchAfter *int    `json:"SearchAfter" name:"SearchAfter"`
	RequestId   *string `json:"RequestId" name:"RequestId"`
}

func (r *ListOperateLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
