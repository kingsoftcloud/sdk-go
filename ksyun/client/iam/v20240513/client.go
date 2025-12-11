package v20240513

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) GetProjectInstanceListNewSend(request *GetProjectInstanceListNewRequest) (*GetProjectInstanceListNewResponse, error) {
	statusCode, msg, err := c.GetProjectInstanceListNewWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetProjectInstanceListNewResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) GetProjectInstanceListNewWithContextV2(ctx context.Context, request *GetProjectInstanceListNewRequest) (int, string, error) {
	if request == nil {
		request = NewGetProjectInstanceListNewRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetProjectInstanceListNewResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
