/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package profile

import (
	"math"
	"time"
)

type DurationFunc func(index int) time.Duration

func ConstantDurationFunc(duration time.Duration) DurationFunc {
	return func(int) time.Duration {
		return duration
	}
}

func ExponentialBackoff(index int) time.Duration {
	seconds := math.Pow(2, float64(index))
	return time.Duration(seconds) * time.Second
}

type ClientProfile struct {
	HttpProfile                    *HttpProfile
	SignMethod                     string
	UnsignedPayload                bool
	Language                       string
	Debug                          bool
	BackupEndPoint                 string
	NetworkFailureRetryDuration    DurationFunc
	RateLimitExceededRetryDuration DurationFunc
}

func NewClientProfile() *ClientProfile {
	return &ClientProfile{
		HttpProfile:     NewHttpProfile(),
		SignMethod:      "HMAC-SHA256",
		UnsignedPayload: false,
		Language:        "zh-CN",
		Debug:           false,
		BackupEndPoint:  "",
	}
}
