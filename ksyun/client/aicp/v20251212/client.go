package v20251212

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2025-12-12"

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

func NewCreateTrainJobRequest() (request *CreateTrainJobRequest) {
	request = &CreateTrainJobRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "CreateTrainJob")
	return
}

func NewCreateTrainJobResponse() (response *CreateTrainJobResponse) {
	response = &CreateTrainJobResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateTrainJob(request *CreateTrainJobRequest) string {
	return c.CreateTrainJobWithContext(context.Background(), request)
}

func (c *Client) CreateTrainJobSend(request *CreateTrainJobRequest) (*CreateTrainJobResponse, error) {
	statusCode, msg, err := c.CreateTrainJobWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct CreateTrainJobResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateTrainJobWithContext(ctx context.Context, request *CreateTrainJobRequest) string {
	if request == nil {
		request = NewCreateTrainJobRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrainJobResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateTrainJobWithContextV2(ctx context.Context, request *CreateTrainJobRequest) (int, string, error) {
	if request == nil {
		request = NewCreateTrainJobRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateTrainJobResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeTrainJobsRequest() (request *DescribeTrainJobsRequest) {
	request = &DescribeTrainJobsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aicp", APIVersion, "DescribeTrainJobs")
	return
}

func NewDescribeTrainJobsResponse() (response *DescribeTrainJobsResponse) {
	response = &DescribeTrainJobsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTrainJobs(request *DescribeTrainJobsRequest) string {
	return c.DescribeTrainJobsWithContext(context.Background(), request)
}

func (c *Client) DescribeTrainJobsSend(request *DescribeTrainJobsRequest) (*DescribeTrainJobsResponse, error) {
	statusCode, msg, err := c.DescribeTrainJobsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	if msg == "" {
		return nil, nil
	}

	var respStruct DescribeTrainJobsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeTrainJobsWithContext(ctx context.Context, request *DescribeTrainJobsRequest) string {
	if request == nil {
		request = NewDescribeTrainJobsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeTrainJobsWithContextV2(ctx context.Context, request *DescribeTrainJobsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeTrainJobsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTrainJobsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
