package v20240415
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeStackOutputsRequest struct {
	*ksyunhttp.BaseRequest
	StackId *string `json:"StackId,omitempty" name:"StackId"`
}

func (r *DescribeStackOutputsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeStackOutputsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Outputs   []struct {
		Name      *string `json:"Name" name:"Name"`
		TypeField *string `json:"TypeField" name:"TypeField"`
		Description *string `json:"Description" name:"Description"`
		Value     *string `json:"Value" name:"Value"`
		Params    *string `json:"Params" name:"Params"`
	} `json:"Outputs"`
}

func (r *DescribeStackOutputsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStackOutputsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeStackEventsRequest struct {
	*ksyunhttp.BaseRequest
	StackId    *string `json:"StackId,omitempty" name:"StackId"`
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	Offset     *int    `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeStackEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeStackEventsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Events    []struct {
		EventType        *string `json:"EventType" name:"EventType"`
		ResourceId       *string `json:"ResourceId" name:"ResourceId"`
		ResourceLogicName *string `json:"ResourceLogicName" name:"ResourceLogicName"`
		ResourceType     *string `json:"ResourceType" name:"ResourceType"`
		EventTime        *string `json:"EventTime" name:"EventTime"`
		EventDescription *string `json:"EventDescription" name:"EventDescription"`
	} `json:"Events"`
}

func (r *DescribeStackEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteTemplateRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteTemplateResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	ReturnField *bool   `json:"Return" name:"Return"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeTemplateVersionsRequest struct {
	*ksyunhttp.BaseRequest
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeTemplateVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTemplateVersionsResponse struct {
	*ksyunhttp.BaseResponse
	Versions []struct {
		VersionNumber *string `json:"VersionNumber" name:"VersionNumber"`
		VersionDescription *string `json:"VersionDescription" name:"VersionDescription"`
		VersionState  *string `json:"VersionState" name:"VersionState"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
	} `json:"Versions"`
	RequestId    *string `json:"RequestId" name:"RequestId"`
	TemplateId   *string `json:"TemplateId" name:"TemplateId"`
	VersionCount *int    `json:"VersionCount" name:"VersionCount"`
}

func (r *DescribeTemplateVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeTemplatesRequest struct {
	*ksyunhttp.BaseRequest
	MaxResults   *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	TemplateId   []*string `json:"TemplateId,omitempty" name:"TemplateId"`
	Offset       *int      `json:"Offset,omitempty" name:"Offset"`
	TemplateName []*string `json:"TemplateName,omitempty" name:"TemplateName"`
	TemplateType *string   `json:"TemplateType,omitempty" name:"TemplateType"`
}

func (r *DescribeTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTemplatesResponse struct {
	*ksyunhttp.BaseResponse
	TemplateCount *int `json:"TemplateCount" name:"TemplateCount"`
	Templates     []struct {
		TemplateId    *string `json:"TemplateId" name:"TemplateId"`
		TemplateType  *string `json:"TemplateType" name:"TemplateType"`
		TemplateName  *string `json:"TemplateName" name:"TemplateName"`
		TemplateDescription *string `json:"TemplateDescription" name:"TemplateDescription"`
		LatestVersion *string `json:"LatestVersion" name:"LatestVersion"`
		VersionCount  *int    `json:"VersionCount" name:"VersionCount"`
		UsingCount    *int    `json:"UsingCount" name:"UsingCount"`
		CreatTime     *string `json:"CreatTime" name:"CreatTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Templates"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

