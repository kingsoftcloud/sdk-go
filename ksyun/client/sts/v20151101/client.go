package v20151101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) AssumeRoleSend(request *AssumeRoleRequest) (*AssumeRoleResponse, error) {
	statusCode, msg, err := c.AssumeRoleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssumeRoleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) AssumeRoleWithContextV2(ctx context.Context, request *AssumeRoleRequest) (int, string, error) {
	if request == nil {
		request = NewAssumeRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssumeRoleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) AssumeRoleTemporarySAMLSend(request *AssumeRoleTemporarySAMLRequest) (*AssumeRoleTemporarySAMLResponse, error) {
	statusCode, msg, err := c.AssumeRoleTemporarySAMLWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AssumeRoleTemporarySAMLResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) AssumeRoleTemporarySAMLWithContextV2(ctx context.Context, request *AssumeRoleTemporarySAMLRequest) (int, string, error) {
	if request == nil {
		request = NewAssumeRoleTemporarySAMLRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAssumeRoleTemporarySAMLResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
