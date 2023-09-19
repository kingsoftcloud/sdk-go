/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package http

import (
	"encoding/json"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	"io/ioutil"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"Code"`
		Message string `json:"Message"`
	} `json:"Error,omitempty"`
	RequestId string `json:"RequestId"`
}

type DeprecatedAPIErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewKsyunSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Error.Code != "" {
		return errors.NewKsyunSDKError(resp.Error.Code, resp.Error.Message, resp.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewKsyunSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewKsyunSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, response Response) string {

	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return msg
	}
	return string(body)
}
