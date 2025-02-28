package v20240703

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2024-07-03"

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

func NewProjectsInfoByInstanceIdsRequest() (request *ProjectsInfoByInstanceIdsRequest) {
	request = &ProjectsInfoByInstanceIdsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("iam", APIVersion, "ProjectsInfoByInstanceIds")
	return
}

func NewProjectsInfoByInstanceIdsResponse() (response *ProjectsInfoByInstanceIdsResponse) {
	response = &ProjectsInfoByInstanceIdsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ProjectsInfoByInstanceIds(request *ProjectsInfoByInstanceIdsRequest) string {
	return c.ProjectsInfoByInstanceIdsWithContext(context.Background(), request)
}

func (c *Client) ProjectsInfoByInstanceIdsWithContext(ctx context.Context, request *ProjectsInfoByInstanceIdsRequest) string {
	if request == nil {
		request = NewProjectsInfoByInstanceIdsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewProjectsInfoByInstanceIdsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
