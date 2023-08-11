/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package profile

type HttpProfile struct {
	ReqMethod  string
	ReqTimeout int
	Scheme     string
	RootDomain string
	Endpoint   string
	Protocol string
}

func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:  "POST",
		ReqTimeout: 60,
		Scheme:     "HTTP",
		RootDomain: "",
		Endpoint:   "",
	}
}
