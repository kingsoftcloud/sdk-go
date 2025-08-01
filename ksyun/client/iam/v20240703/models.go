package v20240703
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type ProjectsInfoByInstanceIdsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ProjectsInfoByInstanceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ProjectsInfoByInstanceIdsResponse struct {
	*ksyunhttp.BaseResponse
	List []struct {
		ProductLine *string `json:"ProductLine" name:"ProductLine"`
		InstanceId  *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		Region      *string `json:"Region" name:"Region"`
		ProjectId   *int    `json:"ProjectId" name:"ProjectId"`
		ProjectName *string `json:"ProjectName" name:"ProjectName"`
		ProjectDesc *string `json:"ProjectDesc" name:"ProjectDesc"`
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

