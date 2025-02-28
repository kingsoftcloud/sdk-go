package v20201231

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-12-31"

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

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCluster(request *CreateClusterRequest) string {
	return c.CreateClusterWithContext(context.Background(), request)
}

func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) string {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
