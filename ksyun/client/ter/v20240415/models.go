package v20240415

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

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

func (r *DescribeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesResponse struct {
	*ksyunhttp.BaseResponse
	TemplateCount *int `json:"TemplateCount" name:"TemplateCount"`
	Templates     []struct {
		TemplateId          *string `json:"TemplateId" name:"TemplateId"`
		TemplateType        *string `json:"TemplateType" name:"TemplateType"`
		TemplateName        *string `json:"TemplateName" name:"TemplateName"`
		TemplateDescription *string `json:"TemplateDescription" name:"TemplateDescription"`
		LatestVersion       *string `json:"LatestVersion" name:"LatestVersion"`
		VersionCount        *int    `json:"VersionCount" name:"VersionCount"`
		UsingCount          *int    `json:"UsingCount" name:"UsingCount"`
		CreatTime           *string `json:"CreatTime" name:"CreatTime"`
		UpdateTime          *string `json:"UpdateTime" name:"UpdateTime"`
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
