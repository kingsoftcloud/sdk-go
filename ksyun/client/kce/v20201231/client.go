package v20201231

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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

func (c *Client) CreateClusterSend(request *CreateClusterRequest) (*CreateClusterResponse, error) {
	statusCode, msg, err := c.CreateClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) string {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kce", APIVersion, "CreateCluster")
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

func (c *Client) CreateClusterWithContextV2(ctx context.Context, request *CreateClusterRequest) (int, string, error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("kce", APIVersion, "CreateCluster")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
