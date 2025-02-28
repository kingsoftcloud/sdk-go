package v20240513

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2024-05-13"

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

func NewGetProjectInstanceListNewRequest() (request *GetProjectInstanceListNewRequest) {
	request = &GetProjectInstanceListNewRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("iam", APIVersion, "GetProjectInstanceListNew")
	return
}

func NewGetProjectInstanceListNewResponse() (response *GetProjectInstanceListNewResponse) {
	response = &GetProjectInstanceListNewResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetProjectInstanceListNew(request *GetProjectInstanceListNewRequest) string {
	return c.GetProjectInstanceListNewWithContext(context.Background(), request)
}

func (c *Client) GetProjectInstanceListNewWithContext(ctx context.Context, request *GetProjectInstanceListNewRequest) string {
	if request == nil {
		request = NewGetProjectInstanceListNewRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetProjectInstanceListNewResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
