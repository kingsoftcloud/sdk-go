package v20230101
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2023-01-01"

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

func NewDescribeEventLogsRequest() (request *DescribeEventLogsRequest) {
	request = &DescribeEventLogsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeEventLogs")
	return
}

func NewDescribeEventLogsResponse() (response *DescribeEventLogsResponse) {
	response = &DescribeEventLogsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEventLogs(request *DescribeEventLogsRequest) string {
	return c.DescribeEventLogsWithContext(context.Background(), request)
}

func (c *Client) DescribeEventLogsSend(request *DescribeEventLogsRequest) (*DescribeEventLogsResponse, error) {
	statusCode, msg, err := c.DescribeEventLogsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEventLogsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEventLogsWithContext(ctx context.Context, request *DescribeEventLogsRequest) string {
	if request == nil {
		request = NewDescribeEventLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEventLogsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEventLogsWithContextV2(ctx context.Context, request *DescribeEventLogsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEventLogsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEventLogsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAddonInstanceRequest() (request *CreateAddonInstanceRequest) {
	request = &CreateAddonInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "CreateAddonInstance")
	return
}

func NewCreateAddonInstanceResponse() (response *CreateAddonInstanceResponse) {
	response = &CreateAddonInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAddonInstance(request *CreateAddonInstanceRequest) string {
	return c.CreateAddonInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateAddonInstanceSend(request *CreateAddonInstanceRequest) (*CreateAddonInstanceResponse, error) {
	statusCode, msg, err := c.CreateAddonInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateAddonInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAddonInstanceWithContext(ctx context.Context, request *CreateAddonInstanceRequest) string {
	if request == nil {
		request = NewCreateAddonInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAddonInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAddonInstanceWithContextV2(ctx context.Context, request *CreateAddonInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAddonInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAddonInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAddonInstanceRequest() (request *DeleteAddonInstanceRequest) {
	request = &DeleteAddonInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DeleteAddonInstance")
	return
}

func NewDeleteAddonInstanceResponse() (response *DeleteAddonInstanceResponse) {
	response = &DeleteAddonInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAddonInstance(request *DeleteAddonInstanceRequest) string {
	return c.DeleteAddonInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteAddonInstanceSend(request *DeleteAddonInstanceRequest) (*DeleteAddonInstanceResponse, error) {
	statusCode, msg, err := c.DeleteAddonInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteAddonInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAddonInstanceWithContext(ctx context.Context, request *DeleteAddonInstanceRequest) string {
	if request == nil {
		request = NewDeleteAddonInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAddonInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAddonInstanceWithContextV2(ctx context.Context, request *DeleteAddonInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAddonInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAddonInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAddonInstancesRequest() (request *DescribeAddonInstancesRequest) {
	request = &DescribeAddonInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeAddonInstances")
	return
}

func NewDescribeAddonInstancesResponse() (response *DescribeAddonInstancesResponse) {
	response = &DescribeAddonInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAddonInstances(request *DescribeAddonInstancesRequest) string {
	return c.DescribeAddonInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeAddonInstancesSend(request *DescribeAddonInstancesRequest) (*DescribeAddonInstancesResponse, error) {
	statusCode, msg, err := c.DescribeAddonInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAddonInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAddonInstancesWithContext(ctx context.Context, request *DescribeAddonInstancesRequest) string {
	if request == nil {
		request = NewDescribeAddonInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddonInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAddonInstancesWithContextV2(ctx context.Context, request *DescribeAddonInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAddonInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddonInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAddonListRequest() (request *DescribeAddonListRequest) {
	request = &DescribeAddonListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeAddonList")
	return
}

func NewDescribeAddonListResponse() (response *DescribeAddonListResponse) {
	response = &DescribeAddonListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAddonList(request *DescribeAddonListRequest) string {
	return c.DescribeAddonListWithContext(context.Background(), request)
}

func (c *Client) DescribeAddonListSend(request *DescribeAddonListRequest) (*DescribeAddonListResponse, error) {
	statusCode, msg, err := c.DescribeAddonListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAddonListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAddonListWithContext(ctx context.Context, request *DescribeAddonListRequest) string {
	if request == nil {
		request = NewDescribeAddonListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddonListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAddonListWithContextV2(ctx context.Context, request *DescribeAddonListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAddonListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAddonListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeComponentParamsRequest() (request *DescribeComponentParamsRequest) {
	request = &DescribeComponentParamsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeComponentParams")
	return
}

func NewDescribeComponentParamsResponse() (response *DescribeComponentParamsResponse) {
	response = &DescribeComponentParamsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeComponentParams(request *DescribeComponentParamsRequest) string {
	return c.DescribeComponentParamsWithContext(context.Background(), request)
}

func (c *Client) DescribeComponentParamsSend(request *DescribeComponentParamsRequest) (*DescribeComponentParamsResponse, error) {
	statusCode, msg, err := c.DescribeComponentParamsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeComponentParamsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeComponentParamsWithContext(ctx context.Context, request *DescribeComponentParamsRequest) string {
	if request == nil {
		request = NewDescribeComponentParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentParamsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeComponentParamsWithContextV2(ctx context.Context, request *DescribeComponentParamsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeComponentParamsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentParamsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNetworkRequest() (request *DescribeNetworkRequest) {
	request = &DescribeNetworkRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeNetwork")
	return
}

func NewDescribeNetworkResponse() (response *DescribeNetworkResponse) {
	response = &DescribeNetworkResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetwork(request *DescribeNetworkRequest) string {
	return c.DescribeNetworkWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkSend(request *DescribeNetworkRequest) (*DescribeNetworkResponse, error) {
	statusCode, msg, err := c.DescribeNetworkWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNetworkResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNetworkWithContext(ctx context.Context, request *DescribeNetworkRequest) string {
	if request == nil {
		request = NewDescribeNetworkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNetworkWithContextV2(ctx context.Context, request *DescribeNetworkRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNetworkRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNetworkResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeNodeComponentsRequest() (request *DescribeNodeComponentsRequest) {
	request = &DescribeNodeComponentsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeNodeComponents")
	return
}

func NewDescribeNodeComponentsResponse() (response *DescribeNodeComponentsResponse) {
	response = &DescribeNodeComponentsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNodeComponents(request *DescribeNodeComponentsRequest) string {
	return c.DescribeNodeComponentsWithContext(context.Background(), request)
}

func (c *Client) DescribeNodeComponentsSend(request *DescribeNodeComponentsRequest) (*DescribeNodeComponentsResponse, error) {
	statusCode, msg, err := c.DescribeNodeComponentsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeNodeComponentsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeNodeComponentsWithContext(ctx context.Context, request *DescribeNodeComponentsRequest) string {
	if request == nil {
		request = NewDescribeNodeComponentsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodeComponentsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeNodeComponentsWithContextV2(ctx context.Context, request *DescribeNodeComponentsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeNodeComponentsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeNodeComponentsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeComponentListRequest() (request *DescribeComponentListRequest) {
	request = &DescribeComponentListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kce", APIVersion, "DescribeComponentList")
	return
}

func NewDescribeComponentListResponse() (response *DescribeComponentListResponse) {
	response = &DescribeComponentListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeComponentList(request *DescribeComponentListRequest) string {
	return c.DescribeComponentListWithContext(context.Background(), request)
}

func (c *Client) DescribeComponentListSend(request *DescribeComponentListRequest) (*DescribeComponentListResponse, error) {
	statusCode, msg, err := c.DescribeComponentListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeComponentListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeComponentListWithContext(ctx context.Context, request *DescribeComponentListRequest) string {
	if request == nil {
		request = NewDescribeComponentListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeComponentListWithContextV2(ctx context.Context, request *DescribeComponentListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeComponentListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeComponentListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


