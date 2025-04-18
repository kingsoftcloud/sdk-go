/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package common

import (
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	"strconv"
)

const iniErr = "ClientError.INIError"

type value struct {
	raw string
}

func (v *value) int() (int, error) {
	i, e := strconv.Atoi(v.raw)
	if e != nil {
		msg := fmt.Sprintf("failed to parsing %s to int: %s", v.raw, e.Error())
		e = errors.NewKsyunSDKError(iniErr, msg, "")
	}
	return i, e
}

func (v *value) int64() (int64, error) {
	i, e := strconv.ParseInt(v.raw, 10, 64)
	if e != nil {
		msg := fmt.Sprintf("failed to parsing %s to int64: %s", v.raw, e.Error())
		e = errors.NewKsyunSDKError(iniErr, msg, "")
	}
	return i, e
}

func (v *value) string() string {
	return v.raw
}

func (v *value) bool() (bool, error) {
	switch v.raw {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "y", "ON", "on", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "n", "OFF", "off", "Off":
		return false, nil
	}
	errorMsg := fmt.Sprintf("failed to parsing \"%s\" to Bool: invalid syntax", v.raw)
	return false, errors.NewKsyunSDKError(iniErr, errorMsg, "")
}

func (v *value) float32() (float32, error) {
	f, e := strconv.ParseFloat(v.raw, 32)
	if e != nil {
		msg := fmt.Sprintf("failed to parse %s to Float32: %s", v.raw, e.Error())
		e = errors.NewKsyunSDKError(iniErr, msg, "")
	}
	return float32(f), e
}
func (v *value) float64() (float64, error) {
	f, e := strconv.ParseFloat(v.raw, 64)
	if e != nil {
		msg := fmt.Sprintf("failed to parse %s to Float64: %s", v.raw, e.Error())
		e = errors.NewKsyunSDKError(iniErr, msg, "")
	}

	return f, e
}
