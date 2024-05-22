package v20240513
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type GetProjectInstanceListNewRequest struct {
    *ksyunhttp.BaseRequest
    ProjectId *int `json:"ProjectId,omitempty" name:"ProjectId"`
    ProductLine *string `json:"ProductLine,omitempty" name:"ProductLine"`
    Ps *int `json:"Ps,omitempty" name:"Ps"`
    Pn *int `json:"Pn,omitempty" name:"Pn"`
}

func (r *GetProjectInstanceListNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProjectInstanceListNewRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetProjectInstanceListNewRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetProjectInstanceListNewResponse struct {
    *ksyunhttp.BaseResponse
	ListInstanceResult struct {
		Items []struct {
					ProjectId *int `json:"ProjectId"`
					InstanceId *string `json:"InstanceId"`
					ProductGroup *int `json:"ProductGroup"`
					ProductType *int `json:"ProductType"`
					Region *string `json:"Region"`
					AccountId *int `json:"AccountId"`
					Status *int `json:"Status"`
			} `json:"Items"`
			CurrentPage *int `json:"CurrentPage"`
			PageSize *int `json:"PageSize"`
			Total *int `json:"Total"`
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

