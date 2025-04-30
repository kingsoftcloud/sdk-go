package v20240930

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-09-30"

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

func NewUpdatePerformanceOnePosixAclRequest() (request *UpdatePerformanceOnePosixAclRequest) {
	request = &UpdatePerformanceOnePosixAclRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "UpdatePerformanceOnePosixAcl")
	return
}

func NewUpdatePerformanceOnePosixAclResponse() (response *UpdatePerformanceOnePosixAclResponse) {
	response = &UpdatePerformanceOnePosixAclResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePerformanceOnePosixAcl(request *UpdatePerformanceOnePosixAclRequest) string {
	return c.UpdatePerformanceOnePosixAclWithContext(context.Background(), request)
}

func (c *Client) UpdatePerformanceOnePosixAclWithContext(ctx context.Context, request *UpdatePerformanceOnePosixAclRequest) string {
	if request == nil {
		request = NewUpdatePerformanceOnePosixAclRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePerformanceOnePosixAclResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribePerformanceOnePosixAclListRequest() (request *DescribePerformanceOnePosixAclListRequest) {
	request = &DescribePerformanceOnePosixAclListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kpfs", APIVersion, "DescribePerformanceOnePosixAclList")
	return
}

func NewDescribePerformanceOnePosixAclListResponse() (response *DescribePerformanceOnePosixAclListResponse) {
	response = &DescribePerformanceOnePosixAclListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePerformanceOnePosixAclList(request *DescribePerformanceOnePosixAclListRequest) string {
	return c.DescribePerformanceOnePosixAclListWithContext(context.Background(), request)
}

func (c *Client) DescribePerformanceOnePosixAclListWithContext(ctx context.Context, request *DescribePerformanceOnePosixAclListRequest) string {
	if request == nil {
		request = NewDescribePerformanceOnePosixAclListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribePerformanceOnePosixAclListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
