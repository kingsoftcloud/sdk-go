/**
 * @author: liyanyan2@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package common

import (
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	ksyunprofile "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
	"log"
	"net/http"
	"net/http/httputil"
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

func (c *Client) Send(request ksyunhttp.Request, response ksyunhttp.Response) (err error) {
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

	return c.sendWithSampleSignature(request, response)
}

func (c *Client) sendWithSampleSignature(request ksyunhttp.Request, response ksyunhttp.Response) (err error) {
	err = ksyunhttp.ConstructParams(request)
	if err != nil {
		return err
	}

	err = signRequest(request, c.credential, c.signMethod)
	if err != nil {
		return err
	}
	httpRequest, err := http.NewRequestWithContext(request.GetContext(), request.GetHttpMethod(), request.GetUrl(), request.GetBodyReader())
	if err != nil {
		return err
	}
	httpRequest.Header.Set("Accept", "application/json")
	if request.GetHttpMethod() == "POST" {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	httpResponse, err := c.sendHttp(httpRequest)
	if err != nil {
		return err
	}
	err = ksyunhttp.ParseFromHttpResponse(httpResponse, response)
	return err
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

