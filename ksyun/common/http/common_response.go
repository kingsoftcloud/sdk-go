/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package http

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

type StringOrInt64 string

func (s *StringOrInt64) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		*s = StringOrInt64(str)
		return nil
	}
	var num int64
	if err := json.Unmarshal(b, &num); err == nil {
		*s = StringOrInt64(strconv.FormatInt(num, 10))
		return nil
	}

	return fmt.Errorf("invalid AccountId")
}
