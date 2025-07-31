package v20250101
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeAlertingResourcesRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	StartTime *int    `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *int    `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAlertingResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAlertingResourcesResponse struct {
	*ksyunhttp.BaseResponse
	ResponseMetadata struct {
		RequestID *string `json:"RequestID" name:"RequestID"`
	} `json:"ResponseMetadata"`
	AlertingProductSet []struct {
		Namespace   *string `json:"Namespace" name:"Namespace"`
		ProductType *int    `json:"ProductType" name:"ProductType"`
		ProductName *string `json:"ProductName" name:"ProductName"`
		ResourceSet []struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			InstanceName *string `json:"InstanceName" name:"InstanceName"`
			InstanceIP   *string `json:"InstanceIP" name:"InstanceIP"`
			InstanceTags []struct {
				ResourceType *string `json:"ResourceType" name:"ResourceType"`
				ResourceId *string `json:"ResourceId" name:"ResourceId"`
				Key   *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
			} `json:"InstanceTags"`
			ProjectID *int    `json:"ProjectID" name:"ProjectID"`
			Region    *string `json:"Region" name:"Region"`
			Policies  []struct {
				PolicyId   *int    `json:"PolicyId" name:"PolicyId"`
				PolicyName *string `json:"PolicyName" name:"PolicyName"`
				PolicyRule *string `json:"PolicyRule" name:"PolicyRule"`
				CurValue   *int    `json:"CurValue" name:"CurValue"`
				AlarmState *int    `json:"AlarmState" name:"AlarmState"`
				TriggerTime *string `json:"TriggerTime" name:"TriggerTime"`
				Duration   *string `json:"Duration" name:"Duration"`
			} `json:"Policies"`
			RelatedResourceSet []struct {
				Namespace *string `json:"Namespace" name:"Namespace"`
				InstanceID *string `json:"InstanceID" name:"InstanceID"`
			} `json:"RelatedResourceSet"`
		} `json:"ResourceSet" name:"ResourceSet"`
	} `json:"AlertingProductSet"`
}

func (r *DescribeAlertingResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertingResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemEventAttributesRequest struct {
	*ksyunhttp.BaseRequest
	StartTime      *int    `json:"StartTime,omitempty" name:"StartTime"`
	EndTime        *int    `json:"EndTime,omitempty" name:"EndTime"`
	Namespace      *string `json:"Namespace,omitempty" name:"Namespace"`
	EventType      *string `json:"EventType,omitempty" name:"EventType"`
	EventName      *string `json:"EventName,omitempty" name:"EventName"`
	Level          *string `json:"Level,omitempty" name:"Level"`
	Status         *string `json:"Status,omitempty" name:"Status"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" name:"SearchKeywords"`
	PageIndex      *int    `json:"PageIndex,omitempty" name:"PageIndex"`
	PageSize       *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeSystemEventAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSystemEventAttributesResponse struct {
	*ksyunhttp.BaseResponse
	RequestID    *string `json:"RequestID" name:"RequestID"`
	SystemEvents []struct {
		Time         *int      `json:"Time" name:"Time"`
		Namespace    *string   `json:"Namespace" name:"Namespace"`
		NamespaceID  *int      `json:"NamespaceID" name:"NamespaceID"`
		EventType    *string   `json:"EventType" name:"EventType"`
		EventName    *string   `json:"EventName" name:"EventName"`
		EventLevel   *string   `json:"EventLevel" name:"EventLevel"`
		EventStatus  *string   `json:"EventStatus" name:"EventStatus"`
		Tags         []*string `json:"Tags" name:"Tags"`
		Content      *string   `json:"Content" name:"Content"`
		InstanceName *string   `json:"InstanceName" name:"InstanceName"`
		InstanceID   *string   `json:"InstanceID" name:"InstanceID"`
		Region       *string   `json:"Region" name:"Region"`
	} `json:"SystemEvents"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeSystemEventAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemEventAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

