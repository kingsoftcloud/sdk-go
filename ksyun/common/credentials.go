/**
 * @author: fengyikai@kingsoft.com
 * @date:  3/18/2022
 * @code: f844b0d56cd31661b4fa3f33787082bb
 */
package common

var creErr = "ClientErr.CredentialErr"

type Credentials interface {
	GetSecretId() string
	GetSecretKey() string
}

type Credential struct {
	SecretId  string
	SecretKey string
	Token     string
}

func NewCredential(secretId, secretKey string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

func (c *Credential) GetSecretKey() string {
	return c.SecretKey
}

func (c *Credential) GetSecretId() string {
	return c.SecretId
}
