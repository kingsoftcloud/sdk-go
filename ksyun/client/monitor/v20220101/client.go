package v20220101

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2022-01-01"

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

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
	request = &CreateAlarmPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
	return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
	response = &CreateAlarmPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) string {
	return c.CreateAlarmPolicyWithContext(context.Background(), request)
}

func (c *Client) CreateAlarmPolicySend(request *CreateAlarmPolicyRequest) (*CreateAlarmPolicyResponse, error) {
	statusCode, msg, err := c.CreateAlarmPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateAlarmPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAlarmPolicyWithContext(ctx context.Context, request *CreateAlarmPolicyRequest) string {
	if request == nil {
		request = NewCreateAlarmPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlarmPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAlarmPolicyWithContextV2(ctx context.Context, request *CreateAlarmPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAlarmPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAlarmPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
	request = &DeleteAlarmPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
	return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
	response = &DeleteAlarmPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) string {
	return c.DeleteAlarmPolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteAlarmPolicySend(request *DeleteAlarmPolicyRequest) (*DeleteAlarmPolicyResponse, error) {
	statusCode, msg, err := c.DeleteAlarmPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DeleteAlarmPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAlarmPolicyWithContext(ctx context.Context, request *DeleteAlarmPolicyRequest) string {
	if request == nil {
		request = NewDeleteAlarmPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteAlarmPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAlarmPolicyWithContextV2(ctx context.Context, request *DeleteAlarmPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAlarmPolicyRequest()
	}
	// 兼容字面量创建的 request，检查 BaseRequest 是否已初始化
	if request.BaseRequest == nil {
		request.BaseRequest = &ksyunhttp.BaseRequest{}
		request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteAlarmPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
