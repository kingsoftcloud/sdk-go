/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package http

import (
	"encoding/json"
)

type actionParameters map[string]interface{}

type CommonRequest struct {
	*BaseRequest
	// custom header, may be overwritten
	header map[string]string
	actionParameters
}

func NewCommonRequest(service, version, action string) (request *CommonRequest) {
	request = &CommonRequest{
		BaseRequest:      &BaseRequest{},
		actionParameters: actionParameters{},
	}
	request.Init().WithApiInfo(service, version, action)
	return
}

func (cr *CommonRequest) SetHeader(header map[string]string) {
	if header == nil {
		return
	}
	cr.header = header
}

func (cr *CommonRequest) GetHeader() map[string]string {
	return cr.header
}

func (cr *CommonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(cr.actionParameters)
}
