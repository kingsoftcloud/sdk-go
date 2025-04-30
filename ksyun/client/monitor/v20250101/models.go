package v20250101
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
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

func (r *DescribeAlertingResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlertingResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
				Key        *string `json:"Key" name:"Key"`
				Value      *string `json:"Value" name:"Value"`
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

