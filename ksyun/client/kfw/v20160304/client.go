package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-03-04"

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

func NewDescribeCfwAvRequest() (request *DescribeCfwAvRequest) {
	request = &DescribeCfwAvRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeCfwAv")
	return
}

func NewDescribeCfwAvResponse() (response *DescribeCfwAvResponse) {
	response = &DescribeCfwAvResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfwAv(request *DescribeCfwAvRequest) string {
	return c.DescribeCfwAvWithContext(context.Background(), request)
}

func (c *Client) DescribeCfwAvSend(request *DescribeCfwAvRequest) (*DescribeCfwAvResponse, error) {
	statusCode, msg, err := c.DescribeCfwAvWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCfwAvResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCfwAvWithContext(ctx context.Context, request *DescribeCfwAvRequest) string {
	if request == nil {
		request = NewDescribeCfwAvRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCfwAvResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCfwAvWithContextV2(ctx context.Context, request *DescribeCfwAvRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCfwAvRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCfwAvResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCloudFireWallInstanceRequest() (request *DescribeCloudFireWallInstanceRequest) {
	request = &DescribeCloudFireWallInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kfw", APIVersion, "DescribeCloudFireWallInstance")
	return
}

func NewDescribeCloudFireWallInstanceResponse() (response *DescribeCloudFireWallInstanceResponse) {
	response = &DescribeCloudFireWallInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCloudFireWallInstance(request *DescribeCloudFireWallInstanceRequest) string {
	return c.DescribeCloudFireWallInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeCloudFireWallInstanceSend(request *DescribeCloudFireWallInstanceRequest) (*DescribeCloudFireWallInstanceResponse, error) {
	statusCode, msg, err := c.DescribeCloudFireWallInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCloudFireWallInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCloudFireWallInstanceWithContext(ctx context.Context, request *DescribeCloudFireWallInstanceRequest) string {
	if request == nil {
		request = NewDescribeCloudFireWallInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCloudFireWallInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCloudFireWallInstanceWithContextV2(ctx context.Context, request *DescribeCloudFireWallInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCloudFireWallInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCloudFireWallInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
