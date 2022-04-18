/**
 * @author: liyanyan2@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package errors

import (
	"fmt"
)

type KsyunSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *KsyunSDKError) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[KsyunSDKError] Code=%s, Message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("[KsyunSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewKsyunSDKError(code, message, requestId string) error {
	return &KsyunSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *KsyunSDKError) GetCode() string {
	return e.Code
}

func (e *KsyunSDKError) GetMessage() string {
	return e.Message
}

func (e *KsyunSDKError) GetRequestId() string {
	return e.RequestId
}
