package v20240415
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-04-15"

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

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ter", APIVersion, "DescribeTemplates")
	return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) string {
	return c.DescribeTemplatesWithContext(context.Background(), request)
}

func (c *Client) DescribeTemplatesWithContext(ctx context.Context, request *DescribeTemplatesRequest) string {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeTemplatesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


