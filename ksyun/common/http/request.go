/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package http

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	POST       = "POST"
	GET        = "GET"
	PUT        = "PUT"
	DELETE     = "DELETE"
	HTTP       = "http"
	HTTPS      = "https"
	RootDomain = "api.ksyun.com"
	Path       = "/"
)

type Request interface {
	GetAction() string
	GetBodyReader() io.Reader
	GetScheme() string
	GetRootDomain() string
	GetServiceDomain(string) string
	GetDomain() string
	GetHttpMethod() string
	GetParams() map[string]string
	GetBody() []byte
	GetPath() string
	GetService() string
	GetUrl() string
	GetVersion() string
	GetContentType() string
	GetContext() context.Context
	SetScheme(string)
	SetRootDomain(string)
	SetDomain(string)
	SetHttpMethod(string)
	SetPath(string)
	SetContentType(string)
	SetBody([]byte)
	SetContext(context.Context)
	GetHeaders() map[string]string
	SetHeaders(map[string]string)
}

type BaseRequest struct {
	context    context.Context
	httpMethod string
	scheme     string
	rootDomain string
	domain     string
	path       string
	params     map[string]string
	formParams map[string]string
	headers    map[string]string

	service string
	version string
	action  string

	contentType string
	body        []byte
}

func (r *BaseRequest) GetAction() string {
	return r.action
}

func (r *BaseRequest) GetHttpMethod() string {
	return r.httpMethod
}

func (r *BaseRequest) GetParams() map[string]string {
	return r.params
}

func (r *BaseRequest) GetPath() string {
	return r.path
}

func (r *BaseRequest) GetDomain() string {
	return r.domain
}

func (r *BaseRequest) GetScheme() string {
	return r.scheme
}

func (r *BaseRequest) GetRootDomain() string {
	return r.rootDomain
}

func (r *BaseRequest) GetServiceDomain(service string) (domain string) {
	rootDomain := r.rootDomain
	if rootDomain == "" {
		rootDomain = RootDomain
	}
	domain = service + "." + rootDomain
	return
}

func (r *BaseRequest) GetBody() []byte {
	return r.body
}

func (r *BaseRequest) SetBody(body []byte) {
	r.body = body
}

func (r *BaseRequest) GetContentType() string {
	return r.contentType
}

func (r *BaseRequest) SetContentType(contentType string) {
	r.contentType = contentType
}

func (r *BaseRequest) SetDomain(domain string) {
	r.domain = domain
}

func (r *BaseRequest) SetScheme(scheme string) {
	scheme = strings.ToLower(scheme)
	switch scheme {
	case HTTP:
		r.scheme = HTTP
	default:
		r.scheme = HTTPS
	}
}

func (r *BaseRequest) SetRootDomain(rootDomain string) {
	r.rootDomain = rootDomain
}

func (r *BaseRequest) SetHeaders(headers map[string]string) {
	r.headers = headers
}
func (r *BaseRequest) GetHeaders() map[string]string {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	return r.headers
}

func (r *BaseRequest) SetHttpMethod(method string) {
	switch strings.ToUpper(method) {
	case POST:
		{
			r.httpMethod = POST
		}
	case GET:
		{
			r.httpMethod = GET
		}
	case PUT:
		{
			r.httpMethod = PUT
		}
	case DELETE:
		{
			r.httpMethod = DELETE
		}
	default:
		{
			r.httpMethod = GET
		}
	}
}

func (r *BaseRequest) SetPath(path string) {
	r.path = path
}

func (r *BaseRequest) GetService() string {
	return r.service
}

func (r *BaseRequest) GetUrl() string {
	if r.httpMethod == GET {
		return r.GetScheme() + "://" + r.domain + r.path + "?" + GetUrlQueriesEncoded(r.params)
	} else if r.httpMethod == POST {
		return r.GetScheme() + "://" + r.domain + r.path
	} else if r.httpMethod == PUT {
		return r.GetScheme() + "://" + r.domain + r.path
	} else if r.httpMethod == DELETE {
		return r.GetScheme() + "://" + r.domain + r.path
	} else {
		return ""
	}
}

func (r *BaseRequest) GetVersion() string {
	return r.version
}

func (r *BaseRequest) GetContext() context.Context {
	if r.context == nil {
		return context.Background()
	}
	return r.context
}

func (r *BaseRequest) SetContext(ctx context.Context) {
	r.context = ctx
}

func GetUrlQueriesEncoded(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	return values.Encode()
}

func (r *BaseRequest) GetBodyReader() io.Reader {
	if r.httpMethod == POST {
		if r.contentType == "application/json" {
			bodyParams := r.GetBody()
			return bytes.NewReader(bodyParams)
		} else {
			s := GetUrlQueriesEncoded(r.params)
			return strings.NewReader(s)
		}
	} else {
		return strings.NewReader("")
	}
}

func (r *BaseRequest) Init() *BaseRequest {
	r.domain = ""
	r.path = Path
	r.params = make(map[string]string)
	r.formParams = make(map[string]string)
	return r
}

func (r *BaseRequest) WithApiInfo(service, version, action string) *BaseRequest {
	r.service = service
	r.version = version
	r.action = action
	return r
}

func (r *BaseRequest) WithContentType(contentType string) *BaseRequest {
	r.contentType = contentType
	return r
}

func HandleCommonParams(request Request, region string) {
	params := request.GetParams()
	params["Service"] = request.GetService()
	params["Action"] = request.GetAction()
	params["Version"] = request.GetVersion()
}

func ConstructParams(req Request) (err error) {
	value := reflect.ValueOf(req).Elem()
	err = FlatStructure(value, req, "")
	//log.Printf("[DEBUG] params=%s", req.GetParams())
	return
}

func FlatStructure(value reflect.Value, request Request, prefix string) (err error) {
	//log.Printf("[DEBUG] reflect value: %v", value.Type())
	valueType := value.Type()
	for i := 0; i < valueType.NumField(); i++ {
		tag := valueType.Field(i).Tag
		nameTag, hasNameTag := tag.Lookup("name")
		if !hasNameTag {
			continue
		}
		field := value.Field(i)
		kind := field.Kind()
		if kind == reflect.Ptr && field.IsNil() {
			continue
		}
		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}
		key := prefix + nameTag
		if kind == reflect.String {
			s := field.String()
			if s != "" {
				request.GetParams()[key] = s
			}
		} else if kind == reflect.Bool {
			request.GetParams()[key] = strconv.FormatBool(field.Bool())
		} else if kind == reflect.Int || kind == reflect.Int64 {
			request.GetParams()[key] = strconv.FormatInt(field.Int(), 10)
		} else if kind == reflect.Uint || kind == reflect.Uint64 {
			request.GetParams()[key] = strconv.FormatUint(field.Uint(), 10)
		} else if kind == reflect.Float64 {
			request.GetParams()[key] = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		} else if kind == reflect.Slice {
			list := value.Field(i)
			for j := 0; j < list.Len(); j++ {
				vj := list.Index(j)
				key := prefix + nameTag + "." + strconv.Itoa(j)
				kind = vj.Kind()
				if kind == reflect.Ptr && vj.IsNil() {
					continue
				}
				if kind == reflect.Ptr {
					vj = vj.Elem()
					kind = vj.Kind()
				}
				if kind == reflect.String {
					request.GetParams()[key] = vj.String()
				} else if kind == reflect.Bool {
					request.GetParams()[key] = strconv.FormatBool(vj.Bool())
				} else if kind == reflect.Int || kind == reflect.Int64 {
					request.GetParams()[key] = strconv.FormatInt(vj.Int(), 10)
				} else if kind == reflect.Uint || kind == reflect.Uint64 {
					request.GetParams()[key] = strconv.FormatUint(vj.Uint(), 10)
				} else if kind == reflect.Float64 {
					request.GetParams()[key] = strconv.FormatFloat(vj.Float(), 'f', -1, 64)
				} else {
					if err = FlatStructure(vj, request, key+"."); err != nil {
						return
					}
				}
			}
		} else {
			if err = FlatStructure(reflect.ValueOf(field.Interface()), request, prefix+nameTag+"."); err != nil {
				return
			}
		}
	}
	return
}
