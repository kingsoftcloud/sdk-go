package v20160901

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-09-01"

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

func NewGetDomainPidDimensionUsageDataRequest() (request *GetDomainPidDimensionUsageDataRequest) {
	request = &GetDomainPidDimensionUsageDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainPidDimensionUsageData")
	return
}

func NewGetDomainPidDimensionUsageDataResponse() (response *GetDomainPidDimensionUsageDataResponse) {
	response = &GetDomainPidDimensionUsageDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainPidDimensionUsageData(request *GetDomainPidDimensionUsageDataRequest) string {
	return c.GetDomainPidDimensionUsageDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainPidDimensionUsageDataWithContext(ctx context.Context, request *GetDomainPidDimensionUsageDataRequest) string {
	if request == nil {
		request = NewGetDomainPidDimensionUsageDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainPidDimensionUsageDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
