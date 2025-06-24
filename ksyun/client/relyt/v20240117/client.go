package v20240117

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-01-17"

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

func NewGetDwsuMetricRequest() (request *GetDwsuMetricRequest) {
	request = &GetDwsuMetricRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("relyt", APIVersion, "GetDwsuMetric")
	return
}

func NewGetDwsuMetricResponse() (response *GetDwsuMetricResponse) {
	response = &GetDwsuMetricResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDwsuMetric(request *GetDwsuMetricRequest) string {
	return c.GetDwsuMetricWithContext(context.Background(), request)
}

func (c *Client) GetDwsuMetricWithContext(ctx context.Context, request *GetDwsuMetricRequest) string {
	if request == nil {
		request = NewGetDwsuMetricRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDwsuMetricResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
