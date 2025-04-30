package v20240701
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-07-01"

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

func NewDescribeCacheByRoleRequest() (request *DescribeCacheByRoleRequest) {
	request = &DescribeCacheByRoleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheByRole")
	return
}

func NewDescribeCacheByRoleResponse() (response *DescribeCacheByRoleResponse) {
	response = &DescribeCacheByRoleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheByRole(request *DescribeCacheByRoleRequest) string {
	return c.DescribeCacheByRoleWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheByRoleWithContext(ctx context.Context, request *DescribeCacheByRoleRequest) string {
	if request == nil {
		request = NewDescribeCacheByRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheByRoleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


