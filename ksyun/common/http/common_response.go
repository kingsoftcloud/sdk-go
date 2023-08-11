/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package http

import "encoding/json"

type actionResult map[string]interface{}
type CommonResponse struct {
	*BaseResponse
	*actionResult
}

func NewCommonResponse() (response *CommonResponse) {
	response = &CommonResponse{
		BaseResponse: &BaseResponse{},
		actionResult: &actionResult{},
	}
	return
}

func (r *CommonResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r.actionResult)
}

func (r *CommonResponse) GetBody() []byte {
	raw, _ := json.Marshal(r.actionResult)
	return raw
}
