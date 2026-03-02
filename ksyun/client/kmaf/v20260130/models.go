package v20260130

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CheckModerateMessage struct {
	Role        *string `json:"Role,omitempty" name:"Role"`
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *int    `json:"ContentType,omitempty" name:"ContentType"`
}
type CheckModerateHistory struct {
	Role        *string `json:"Role,omitempty" name:"Role"`
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *int    `json:"ContentType,omitempty" name:"ContentType"`
}

type QueryAnswerRequest struct {
	*ksyunhttp.BaseRequest
	AppId     *string `json:"AppId,omitempty" name:"AppId"`
	MsgId     *string `json:"MsgId,omitempty" name:"MsgId"`
	UseStream *int    `json:"UseStream,omitempty" name:"UseStream"`
}

func (r *QueryAnswerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryAnswerResponse struct {
	*ksyunhttp.BaseResponse
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Message struct {
			Role        *string `json:"Role" name:"Role"`
			Content     *string `json:"Content" name:"Content"`
			ContentType *int    `json:"ContentType" name:"ContentType"`
		} `json:"Message" name:"Message"`
		IsFinished *bool   `json:"IsFinished" name:"IsFinished"`
		MsgId      *string `json:"MsgId" name:"MsgId"`
	} `json:"Data"`
}

func (r *QueryAnswerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAnswerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckModerateRequest struct {
	*ksyunhttp.BaseRequest
	AppId     *string                 `json:"AppId,omitempty" name:"AppId"`
	MsgId     *string                 `json:"MsgId,omitempty" name:"MsgId"`
	UseStream *int                    `json:"UseStream,omitempty" name:"UseStream"`
	Message   *CheckModerateMessage   `json:"Message,omitempty" name:"Message"`
	History   []*CheckModerateHistory `json:"History,omitempty" name:"History"`
}

func (r *CheckModerateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CheckModerateResponse struct {
	*ksyunhttp.BaseResponse
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Decision struct {
			Action *string `json:"Action" name:"Action"`
			Score  *int    `json:"Score" name:"Score"`
		} `json:"Decision" name:"Decision"`
		RiskInfo struct {
			Risks []struct {
				Category *string `json:"Category" name:"Category"`
				Label    *string `json:"Label" name:"Label"`
				Score    *int    `json:"Score" name:"Score"`
				Matches  []struct {
					Word *string `json:"Word" name:"Word"`
				} `json:"Matches" name:"Matches"`
			} `json:"Risks"`
		} `json:"RiskInfo" name:"RiskInfo"`
		MsgId *string `json:"MsgId" name:"MsgId"`
	} `json:"Data"`
}

func (r *CheckModerateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckModerateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
