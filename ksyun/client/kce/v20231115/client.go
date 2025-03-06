package v20231115

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2023-11-15"

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

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeCluster")
	return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCluster(request *DescribeClusterRequest) string {
	return c.DescribeClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterWithContext(ctx context.Context, request *DescribeClusterRequest) string {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateClusterDelProtectionRequest() (request *UpdateClusterDelProtectionRequest) {
	request = &UpdateClusterDelProtectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "UpdateClusterDelProtection")
	return
}

func NewUpdateClusterDelProtectionResponse() (response *UpdateClusterDelProtectionResponse) {
	response = &UpdateClusterDelProtectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateClusterDelProtection(request *UpdateClusterDelProtectionRequest) string {
	return c.UpdateClusterDelProtectionWithContext(context.Background(), request)
}

func (c *Client) UpdateClusterDelProtectionWithContext(ctx context.Context, request *UpdateClusterDelProtectionRequest) string {
	if request == nil {
		request = NewUpdateClusterDelProtectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateClusterDelProtectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
