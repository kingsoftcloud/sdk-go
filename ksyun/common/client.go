/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package common

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	ksyunprofile "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Client struct {
	region          string
	httpClient      *http.Client
	httpProfile     *ksyunprofile.HttpProfile
	profile         *ksyunprofile.ClientProfile
	credential      Credentials
	signMethod      string
	unsignedPayload bool
	debug           bool
}

func (c *Client) Send(request ksyunhttp.Request, response ksyunhttp.Response) (error, string) {
	if request.GetScheme() == "" {
		request.SetScheme(c.httpProfile.Scheme)
	}

	if request.GetDomain() == "" {
		domain := c.httpProfile.Endpoint
		if domain == "" {
			domain = request.GetServiceDomain(request.GetService())
		}
		request.SetDomain(domain)
	}

	if request.GetHttpMethod() == "" {
		request.SetHttpMethod(c.httpProfile.ReqMethod)
	}

	ksyunhttp.HandleCommonParams(request, c.GetRegion())

	err, msg := c.sendWithSampleSignature(request, response)
	return err, msg
}

func (c *Client) sendWithSampleSignature(request ksyunhttp.Request, response ksyunhttp.Response) (error, string) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(c.GetRegion()),
		Credentials: credentials.NewStaticCredentials(c.credential.GetSecretId(), c.credential.GetSecretKey(), ""),
	}))
	signer := v4.NewSigner(sess.Config.Credentials)
	customizeHeaders := request.GetHeaders()
	method := request.GetHttpMethod()
	contentType := request.GetContentType()
	now := time.Now().UTC()
	urlRe := request.GetUrl()

	// 公共Header设置函数
	setCommonHeaders := func(req *http.Request) {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Amz-Date", now.Format("20060102T150405Z"))
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Host", request.GetDomain())
		req.Header.Set(SDKName, SDKVersion)
		for headerK, headerV := range customizeHeaders {
			req.Header.Set(headerK, headerV)
		}
	}

	isJson := contentType == "application/json"
	isForm := contentType == "application/x-www-form-urlencoded"
	isBodyMethod := method == "POST" || method == "PUT" || method == "DELETE"

	if isBodyMethod && isJson {
		// JSON 请求
		body, _ := json.Marshal(request)
		request.SetBody(body)

		urlRe = urlRe + "?Action=" + request.GetAction() + "&Version=" + request.GetVersion() + "&Service=" + request.GetService()

		httpRequest, err := http.NewRequestWithContext(request.GetContext(), method, urlRe, request.GetBodyReader())
		if err != nil {
			return err, ""
		}
		httpRequest.ContentLength = int64(len(body))
		setCommonHeaders(httpRequest)

		_, err = signer.Sign(httpRequest, bytes.NewReader(body), request.GetService(), c.GetRegion(), now)
		if err != nil {
			return err, ""
		}

		httpResponse, err := c.sendHttp(httpRequest)
		if err != nil {
			return err, ""
		}
		res := ksyunhttp.ParseFromHttpResponse(httpResponse, response)
		return nil, res
	} else if isBodyMethod && isForm {
		// 表单请求
		value := reflect.ValueOf(request).Elem()
		err := ksyunhttp.FlatStructure(value, request, "")
		if err != nil {
			return err, ""
		}

		paramsM := request.GetParams()
		formData := url.Values{}
		for key, value := range paramsM {
			formData.Set(key, value)
		}
		formDataEncoded := formData.Encode()
		request.SetBody([]byte(formDataEncoded))

		httpRequest, err := http.NewRequestWithContext(request.GetContext(), method, urlRe, request.GetBodyReader())
		if err != nil {
			return err, ""
		}
		setCommonHeaders(httpRequest)

		_, err = signer.Sign(httpRequest, strings.NewReader(formDataEncoded), request.GetService(), c.GetRegion(), now)
		if err != nil {
			return err, ""
		}

		httpResponse, err := c.sendHttp(httpRequest)
		if err != nil {
			return err, ""
		}
		res := ksyunhttp.ParseFromHttpResponse(httpResponse, response)
		return nil, res
	} else {
		value := reflect.ValueOf(request).Elem()
		err := ksyunhttp.FlatStructure(value, request, "")
		if err != nil {
			return err, ""
		}

		httpRequest, err := http.NewRequestWithContext(request.GetContext(), method, urlRe, request.GetBodyReader())
		if err != nil {
			return err, ""
		}
		setCommonHeaders(httpRequest)

		_, err = signer.Sign(httpRequest, nil, request.GetService(), c.GetRegion(), now)
		if err != nil {
			return err, ""
		}

		httpResponse, err := c.sendHttp(httpRequest)
		if err != nil {
			return err, ""
		}
		res := ksyunhttp.ParseFromHttpResponse(httpResponse, response)
		return nil, res
	}
}

// send http request
func (c *Client) sendHttp(request *http.Request) (response *http.Response, err error) {
	if c.debug {
		outBytes, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Printf("[ERROR] dump request failed because %s", err)
			return nil, err
		}
		log.Printf("[DEBUG] http request = %s", outBytes)
	}
	response, err = c.httpClient.Do(request)
	return response, err
}

func (c *Client) GetRegion() string {
	return c.region
}

func (c *Client) Init(region string) *Client {
	c.httpClient = &http.Client{}
	c.region = region
	c.signMethod = "HMAC-SHA256"
	c.debug = false
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return c
}

func (c *Client) WithSecretId(secretId, secretKey string) *Client {
	c.credential = NewCredential(secretId, secretKey)
	return c
}
func (c *Client) WithCredential(cred Credentials) *Client {
	c.credential = cred
	return c
}
func (c *Client) WithProfile(clientProfile *ksyunprofile.ClientProfile) *Client {
	c.profile = clientProfile
	c.signMethod = clientProfile.SignMethod
	c.unsignedPayload = clientProfile.UnsignedPayload
	c.httpProfile = clientProfile.HttpProfile
	c.debug = clientProfile.Debug
	c.httpClient.Timeout = time.Duration(c.httpProfile.ReqTimeout) * time.Second
	return c
}

func (c *Client) WithSignatureMethod(method string) *Client {
	c.signMethod = method
	return c
}

func (c *Client) WithHttpTransport(transport http.RoundTripper) *Client {
	c.httpClient.Transport = transport
	return c
}

func (c *Client) WithDebug(flag bool) *Client {
	c.debug = flag
	return c
}
