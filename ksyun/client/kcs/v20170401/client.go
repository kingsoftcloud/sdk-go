package v20170401
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2017-04-01"

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

func NewDeleteCacheSlaveNodeRequest() (request *DeleteCacheSlaveNodeRequest) {
	request = &DeleteCacheSlaveNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteCacheSlaveNode")
	return
}

func NewDeleteCacheSlaveNodeResponse() (response *DeleteCacheSlaveNodeResponse) {
	response = &DeleteCacheSlaveNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCacheSlaveNode(request *DeleteCacheSlaveNodeRequest) string {
	return c.DeleteCacheSlaveNodeWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheSlaveNodeSend(request *DeleteCacheSlaveNodeRequest) (*DeleteCacheSlaveNodeResponse, error) {
	statusCode, msg, err := c.DeleteCacheSlaveNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteCacheSlaveNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCacheSlaveNodeWithContext(ctx context.Context, request *DeleteCacheSlaveNodeRequest) string {
	if request == nil {
		request = NewDeleteCacheSlaveNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCacheSlaveNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCacheSlaveNodeWithContextV2(ctx context.Context, request *DeleteCacheSlaveNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCacheSlaveNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCacheSlaveNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


