package v20151101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2015-11-01"

type Client struct {
	common.Client
}

func NewClient(credential common.Credentials, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewAssumeRoleRequest() (request *AssumeRoleRequest) {
	request = &AssumeRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sts", APIVersion, "AssumeRole")
	return
}

func NewAssumeRoleResponse() (response *AssumeRoleResponse) {
	response = &AssumeRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssumeRole(request *AssumeRoleRequest) string {
	return c.AssumeRoleWithContext(context.Background(), request)
}

func (c *Client) AssumeRoleWithContext(ctx context.Context, request *AssumeRoleRequest) string {
	if request == nil {
		request = NewAssumeRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssumeRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAssumeRoleTemporarySAMLRequest() (request *AssumeRoleTemporarySAMLRequest) {
	request = &AssumeRoleTemporarySAMLRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("sts", APIVersion, "AssumeRoleTemporarySAML")
	return
}

func NewAssumeRoleTemporarySAMLResponse() (response *AssumeRoleTemporarySAMLResponse) {
	response = &AssumeRoleTemporarySAMLResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssumeRoleTemporarySAML(request *AssumeRoleTemporarySAMLRequest) string {
	return c.AssumeRoleTemporarySAMLWithContext(context.Background(), request)
}

func (c *Client) AssumeRoleTemporarySAMLWithContext(ctx context.Context, request *AssumeRoleTemporarySAMLRequest) string {
	if request == nil {
		request = NewAssumeRoleTemporarySAMLRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAssumeRoleTemporarySAMLResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
