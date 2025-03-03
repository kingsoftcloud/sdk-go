package v20190806

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2019-08-06"

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

func NewQueryPodsByInformerRequest() (request *QueryPodsByInformerRequest) {
	request = &QueryPodsByInformerRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce2", APIVersion, "QueryPodsByInformer")
	return
}

func NewQueryPodsByInformerResponse() (response *QueryPodsByInformerResponse) {
	response = &QueryPodsByInformerResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryPodsByInformer(request *QueryPodsByInformerRequest) string {
	return c.QueryPodsByInformerWithContext(context.Background(), request)
}

func (c *Client) QueryPodsByInformerWithContext(ctx context.Context, request *QueryPodsByInformerRequest) string {
	if request == nil {
		request = NewQueryPodsByInformerRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryPodsByInformerResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
