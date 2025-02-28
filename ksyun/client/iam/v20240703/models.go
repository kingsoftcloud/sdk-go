package v20240703

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type ProjectsInfoByInstanceIdsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ProjectsInfoByInstanceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProjectsInfoByInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ProjectsInfoByInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ProjectsInfoByInstanceIdsResponse struct {
	*ksyunhttp.BaseResponse
	List []struct {
		ProductLine  *string `json:"ProductLine"`
		InstanceId   *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		Region       *string `json:"Region"`
		ProjectId    *int    `json:"ProjectId"`
		ProjectName  *string `json:"ProjectName"`
		ProjectDesc  *string `json:"ProjectDesc"`
	} `json:"List"`
	Total     *int    `json:"Total" name:"Total"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ProjectsInfoByInstanceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProjectsInfoByInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
