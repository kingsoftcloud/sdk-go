package v20240703

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) ProjectsInfoByInstanceIdsSend(request *ProjectsInfoByInstanceIdsRequest) (*ProjectsInfoByInstanceIdsResponse, error) {
	statusCode, msg, err := c.ProjectsInfoByInstanceIdsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct ProjectsInfoByInstanceIdsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ProjectsInfoByInstanceIdsWithContextV2(ctx context.Context, request *ProjectsInfoByInstanceIdsRequest) (int, string, error) {
	if request == nil {
		request = NewProjectsInfoByInstanceIdsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewProjectsInfoByInstanceIdsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
