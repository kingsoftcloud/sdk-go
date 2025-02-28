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
		ErrorMessage      *string `json:"ErrorMessage"`
		CreateTime        *string `json:"CreateTime"`
		ServiceName       *string `json:"ServiceName"`
		EventSource       *string `json:"EventSource"`
		ApiVersion        *string `json:"ApiVersion"`
		RequestParameters struct {
		} `json:"RequestParameters"`
		SourceIpAddress *string `json:"SourceIpAddress"`
		EventVersion    *string `json:"EventVersion"`
		EventType       *string `json:"EventType"`
		EventId         *string `json:"EventId"`
		EventRw         *string `json:"EventRw"`
		EventName       *string `json:"EventName"`
		UserIdentity    struct {
			RoleName  *string `json:"RoleName"`
			AccountId *string `json:"AccountId"`
			UserName  *string `json:"UserName"`
			AccessKey *string `json:"AccessKey"`
		} `json:"UserIdentity"`
		ReferencedResources []struct {
		} `json:"ReferencedResources"`
		ErrorCode    *string `json:"ErrorCode"`
		Region       *string `json:"Region"`
		RequestId    *string `json:"RequestId"`
		EventTime    *string `json:"EventTime"`
		UserAgent    *string `json:"UserAgent"`
		ResourceType *string `json:"ResourceType"`
		ResourceName *string `json:"ResourceName"`
		RegionCn     *string `json:"RegionCn"`
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
