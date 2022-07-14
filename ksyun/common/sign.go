/**
 * @author: liyanyan2@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"net/url"
	"sort"
	"strings"
)

const (
	SHA256 = "HmacSHA256"
)

func Sign(s, secretKey, method string) string {
	hashed := hmac.New(sha256.New, []byte(secretKey))
	hashed.Write([]byte(s))

	return hex.EncodeToString(hashed.Sum(nil))
}
func signRequest(request ksyunhttp.Request, credential Credentials, method string) (err error) {
	checkAuthParams(request, credential, method)
	s := getStringToSign(request)
	signature := Sign(s, credential.GetSecretKey(), method)
	request.GetParams()["Signature"] = signature
	return
}

func checkAuthParams(request ksyunhttp.Request, credential Credentials, method string) {
	params := request.GetParams()
	params["Accesskey"] = credential.GetSecretId()
	params["SignatureMethod"] = method
	delete(params, "Signature")
}

func getStringToSign(request ksyunhttp.Request) string {
	var buf bytes.Buffer
	params := request.GetParams()
	// sort params
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := range keys {
		k := keys[i]
		buf.WriteString(strings.Replace(url.QueryEscape(k), "+", "%20", -1))
		buf.WriteString("=")
		buf.WriteString(strings.Replace(url.QueryEscape(params[k]), "+", "%20", -1))
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	signStr := buf.String()
	return signStr
}
