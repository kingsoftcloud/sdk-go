package v20240513
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type GetProjectInstanceListNewRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId   *int    `json:"ProjectId,omitempty" name:"ProjectId"`
	ProductLine *string `json:"ProductLine,omitempty" name:"ProductLine"`
	Ps          *int    `json:"Ps,omitempty" name:"Ps"`
	Pn          *int    `json:"Pn,omitempty" name:"Pn"`
}

func (r *GetProjectInstanceListNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetProjectInstanceListNewResponse struct {
	*ksyunhttp.BaseResponse
	ListInstanceResult struct {
		Items []struct {
			ProjectId    *int    `json:"ProjectId" name:"ProjectId"`
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			ProductGroup *int    `json:"ProductGroup" name:"ProductGroup"`
			ProductType  *int    `json:"ProductType" name:"ProductType"`
			Region       *string `json:"Region" name:"Region"`
			AccountId    *int    `json:"AccountId" name:"AccountId"`
			Status       *int    `json:"Status" name:"Status"`
		} `json:"Items" name:"Items"`
		CurrentPage *int `json:"CurrentPage" name:"CurrentPage"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		Total    *int `json:"Total" name:"Total"`
	} `json:"ListInstanceResult"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetProjectInstanceListNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectInstanceListNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

