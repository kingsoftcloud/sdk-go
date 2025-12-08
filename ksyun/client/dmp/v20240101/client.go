package v20240101
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2024-01-01"

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

func NewDescribeDefaultMonitorItemsRequest() (request *DescribeDefaultMonitorItemsRequest) {
	request = &DescribeDefaultMonitorItemsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDefaultMonitorItems")
	return
}

func NewDescribeDefaultMonitorItemsResponse() (response *DescribeDefaultMonitorItemsResponse) {
	response = &DescribeDefaultMonitorItemsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDefaultMonitorItems(request *DescribeDefaultMonitorItemsRequest) string {
	return c.DescribeDefaultMonitorItemsWithContext(context.Background(), request)
}

func (c *Client) DescribeDefaultMonitorItemsSend(request *DescribeDefaultMonitorItemsRequest) (*DescribeDefaultMonitorItemsResponse, error) {
	statusCode, msg, err := c.DescribeDefaultMonitorItemsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDefaultMonitorItemsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDefaultMonitorItemsWithContext(ctx context.Context, request *DescribeDefaultMonitorItemsRequest) string {
	if request == nil {
		request = NewDescribeDefaultMonitorItemsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDefaultMonitorItemsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDefaultMonitorItemsWithContextV2(ctx context.Context, request *DescribeDefaultMonitorItemsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDefaultMonitorItemsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDefaultMonitorItemsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMonitorPanelRequest() (request *DeleteMonitorPanelRequest) {
	request = &DeleteMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DeleteMonitorPanel")
	return
}

func NewDeleteMonitorPanelResponse() (response *DeleteMonitorPanelResponse) {
	response = &DeleteMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMonitorPanel(request *DeleteMonitorPanelRequest) string {
	return c.DeleteMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) DeleteMonitorPanelSend(request *DeleteMonitorPanelRequest) (*DeleteMonitorPanelResponse, error) {
	statusCode, msg, err := c.DeleteMonitorPanelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteMonitorPanelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMonitorPanelWithContext(ctx context.Context, request *DeleteMonitorPanelRequest) string {
	if request == nil {
		request = NewDeleteMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMonitorPanelWithContextV2(ctx context.Context, request *DeleteMonitorPanelRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteMonitorPanelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOperateMonitorPanelRequest() (request *OperateMonitorPanelRequest) {
	request = &OperateMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "OperateMonitorPanel")
	return
}

func NewOperateMonitorPanelResponse() (response *OperateMonitorPanelResponse) {
	response = &OperateMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OperateMonitorPanel(request *OperateMonitorPanelRequest) string {
	return c.OperateMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) OperateMonitorPanelSend(request *OperateMonitorPanelRequest) (*OperateMonitorPanelResponse, error) {
	statusCode, msg, err := c.OperateMonitorPanelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct OperateMonitorPanelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OperateMonitorPanelWithContext(ctx context.Context, request *OperateMonitorPanelRequest) string {
	if request == nil {
		request = NewOperateMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperateMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OperateMonitorPanelWithContextV2(ctx context.Context, request *OperateMonitorPanelRequest) (int, string, error) {
	if request == nil {
		request = NewOperateMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewOperateMonitorPanelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMonitorPanelRequest() (request *DescribeMonitorPanelRequest) {
	request = &DescribeMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeMonitorPanel")
	return
}

func NewDescribeMonitorPanelResponse() (response *DescribeMonitorPanelResponse) {
	response = &DescribeMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorPanel(request *DescribeMonitorPanelRequest) string {
	return c.DescribeMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorPanelSend(request *DescribeMonitorPanelRequest) (*DescribeMonitorPanelResponse, error) {
	statusCode, msg, err := c.DescribeMonitorPanelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMonitorPanelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMonitorPanelWithContext(ctx context.Context, request *DescribeMonitorPanelRequest) string {
	if request == nil {
		request = NewDescribeMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMonitorPanelWithContextV2(ctx context.Context, request *DescribeMonitorPanelRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorPanelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyMonitorPanelInfoRequest() (request *ModifyMonitorPanelInfoRequest) {
	request = &ModifyMonitorPanelInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "ModifyMonitorPanelInfo")
	return
}

func NewModifyMonitorPanelInfoResponse() (response *ModifyMonitorPanelInfoResponse) {
	response = &ModifyMonitorPanelInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyMonitorPanelInfo(request *ModifyMonitorPanelInfoRequest) string {
	return c.ModifyMonitorPanelInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyMonitorPanelInfoSend(request *ModifyMonitorPanelInfoRequest) (*ModifyMonitorPanelInfoResponse, error) {
	statusCode, msg, err := c.ModifyMonitorPanelInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyMonitorPanelInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyMonitorPanelInfoWithContext(ctx context.Context, request *ModifyMonitorPanelInfoRequest) string {
	if request == nil {
		request = NewModifyMonitorPanelInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyMonitorPanelInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyMonitorPanelInfoWithContextV2(ctx context.Context, request *ModifyMonitorPanelInfoRequest) (int, string, error) {
	if request == nil {
		request = NewModifyMonitorPanelInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyMonitorPanelInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateMonitorPanelRequest() (request *CreateMonitorPanelRequest) {
	request = &CreateMonitorPanelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "CreateMonitorPanel")
	return
}

func NewCreateMonitorPanelResponse() (response *CreateMonitorPanelResponse) {
	response = &CreateMonitorPanelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateMonitorPanel(request *CreateMonitorPanelRequest) string {
	return c.CreateMonitorPanelWithContext(context.Background(), request)
}

func (c *Client) CreateMonitorPanelSend(request *CreateMonitorPanelRequest) (*CreateMonitorPanelResponse, error) {
	statusCode, msg, err := c.CreateMonitorPanelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateMonitorPanelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateMonitorPanelWithContext(ctx context.Context, request *CreateMonitorPanelRequest) string {
	if request == nil {
		request = NewCreateMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMonitorPanelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateMonitorPanelWithContextV2(ctx context.Context, request *CreateMonitorPanelRequest) (int, string, error) {
	if request == nil {
		request = NewCreateMonitorPanelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateMonitorPanelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBatchInstancesRequest() (request *DeleteBatchInstancesRequest) {
	request = &DeleteBatchInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DeleteBatchInstances")
	return
}

func NewDeleteBatchInstancesResponse() (response *DeleteBatchInstancesResponse) {
	response = &DeleteBatchInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBatchInstances(request *DeleteBatchInstancesRequest) string {
	return c.DeleteBatchInstancesWithContext(context.Background(), request)
}

func (c *Client) DeleteBatchInstancesSend(request *DeleteBatchInstancesRequest) (*DeleteBatchInstancesResponse, error) {
	statusCode, msg, err := c.DeleteBatchInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteBatchInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBatchInstancesWithContext(ctx context.Context, request *DeleteBatchInstancesRequest) string {
	if request == nil {
		request = NewDeleteBatchInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBatchInstancesWithContextV2(ctx context.Context, request *DeleteBatchInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBatchInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteBatchInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewInstanceStatisticsRequest() (request *InstanceStatisticsRequest) {
	request = &InstanceStatisticsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "InstanceStatistics")
	return
}

func NewInstanceStatisticsResponse() (response *InstanceStatisticsResponse) {
	response = &InstanceStatisticsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) InstanceStatistics(request *InstanceStatisticsRequest) string {
	return c.InstanceStatisticsWithContext(context.Background(), request)
}

func (c *Client) InstanceStatisticsSend(request *InstanceStatisticsRequest) (*InstanceStatisticsResponse, error) {
	statusCode, msg, err := c.InstanceStatisticsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct InstanceStatisticsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) InstanceStatisticsWithContext(ctx context.Context, request *InstanceStatisticsRequest) string {
	if request == nil {
		request = NewInstanceStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstanceStatisticsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) InstanceStatisticsWithContextV2(ctx context.Context, request *InstanceStatisticsRequest) (int, string, error) {
	if request == nil {
		request = NewInstanceStatisticsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewInstanceStatisticsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMonitorPanelListRequest() (request *DescribeMonitorPanelListRequest) {
	request = &DescribeMonitorPanelListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeMonitorPanelList")
	return
}

func NewDescribeMonitorPanelListResponse() (response *DescribeMonitorPanelListResponse) {
	response = &DescribeMonitorPanelListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMonitorPanelList(request *DescribeMonitorPanelListRequest) string {
	return c.DescribeMonitorPanelListWithContext(context.Background(), request)
}

func (c *Client) DescribeMonitorPanelListSend(request *DescribeMonitorPanelListRequest) (*DescribeMonitorPanelListResponse, error) {
	statusCode, msg, err := c.DescribeMonitorPanelListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMonitorPanelListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMonitorPanelListWithContext(ctx context.Context, request *DescribeMonitorPanelListRequest) string {
	if request == nil {
		request = NewDescribeMonitorPanelListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorPanelListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMonitorPanelListWithContextV2(ctx context.Context, request *DescribeMonitorPanelListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMonitorPanelListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeMonitorPanelListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
	request = &DescribeInstanceListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeInstanceList")
	return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
	response = &DescribeInstanceListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) string {
	return c.DescribeInstanceListWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceListSend(request *DescribeInstanceListRequest) (*DescribeInstanceListResponse, error) {
	statusCode, msg, err := c.DescribeInstanceListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) string {
	if request == nil {
		request = NewDescribeInstanceListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceListWithContextV2(ctx context.Context, request *DescribeInstanceListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewImportInstanceToDmpRequest() (request *ImportInstanceToDmpRequest) {
	request = &ImportInstanceToDmpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "ImportInstanceToDmp")
	return
}

func NewImportInstanceToDmpResponse() (response *ImportInstanceToDmpResponse) {
	response = &ImportInstanceToDmpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ImportInstanceToDmp(request *ImportInstanceToDmpRequest) string {
	return c.ImportInstanceToDmpWithContext(context.Background(), request)
}

func (c *Client) ImportInstanceToDmpSend(request *ImportInstanceToDmpRequest) (*ImportInstanceToDmpResponse, error) {
	statusCode, msg, err := c.ImportInstanceToDmpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ImportInstanceToDmpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ImportInstanceToDmpWithContext(ctx context.Context, request *ImportInstanceToDmpRequest) string {
	if request == nil {
		request = NewImportInstanceToDmpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImportInstanceToDmpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ImportInstanceToDmpWithContextV2(ctx context.Context, request *ImportInstanceToDmpRequest) (int, string, error) {
	if request == nil {
		request = NewImportInstanceToDmpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImportInstanceToDmpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDedicatedClustersRequest() (request *DescribeDedicatedClustersRequest) {
	request = &DescribeDedicatedClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDedicatedClusters")
	return
}

func NewDescribeDedicatedClustersResponse() (response *DescribeDedicatedClustersResponse) {
	response = &DescribeDedicatedClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedClusters(request *DescribeDedicatedClustersRequest) string {
	return c.DescribeDedicatedClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedClustersSend(request *DescribeDedicatedClustersRequest) (*DescribeDedicatedClustersResponse, error) {
	statusCode, msg, err := c.DescribeDedicatedClustersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDedicatedClustersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDedicatedClustersWithContext(ctx context.Context, request *DescribeDedicatedClustersRequest) string {
	if request == nil {
		request = NewDescribeDedicatedClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDedicatedClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDedicatedClustersWithContextV2(ctx context.Context, request *DescribeDedicatedClustersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDedicatedClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDedicatedClustersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDedicatedHostsRequest() (request *DescribeDedicatedHostsRequest) {
	request = &DescribeDedicatedHostsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDedicatedHosts")
	return
}

func NewDescribeDedicatedHostsResponse() (response *DescribeDedicatedHostsResponse) {
	response = &DescribeDedicatedHostsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedHosts(request *DescribeDedicatedHostsRequest) string {
	return c.DescribeDedicatedHostsWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedHostsSend(request *DescribeDedicatedHostsRequest) (*DescribeDedicatedHostsResponse, error) {
	statusCode, msg, err := c.DescribeDedicatedHostsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDedicatedHostsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDedicatedHostsWithContext(ctx context.Context, request *DescribeDedicatedHostsRequest) string {
	if request == nil {
		request = NewDescribeDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedHostsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDedicatedHostsWithContextV2(ctx context.Context, request *DescribeDedicatedHostsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDedicatedHostsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedHostsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDatabaseSchemaRequest() (request *DescribeDatabaseSchemaRequest) {
	request = &DescribeDatabaseSchemaRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDatabaseSchema")
	return
}

func NewDescribeDatabaseSchemaResponse() (response *DescribeDatabaseSchemaResponse) {
	response = &DescribeDatabaseSchemaResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDatabaseSchema(request *DescribeDatabaseSchemaRequest) string {
	return c.DescribeDatabaseSchemaWithContext(context.Background(), request)
}

func (c *Client) DescribeDatabaseSchemaSend(request *DescribeDatabaseSchemaRequest) (*DescribeDatabaseSchemaResponse, error) {
	statusCode, msg, err := c.DescribeDatabaseSchemaWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDatabaseSchemaResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDatabaseSchemaWithContext(ctx context.Context, request *DescribeDatabaseSchemaRequest) string {
	if request == nil {
		request = NewDescribeDatabaseSchemaRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDatabaseSchemaResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDatabaseSchemaWithContextV2(ctx context.Context, request *DescribeDatabaseSchemaRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDatabaseSchemaRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDatabaseSchemaResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDatabaseListRequest() (request *DescribeDatabaseListRequest) {
	request = &DescribeDatabaseListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeDatabaseList")
	return
}

func NewDescribeDatabaseListResponse() (response *DescribeDatabaseListResponse) {
	response = &DescribeDatabaseListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDatabaseList(request *DescribeDatabaseListRequest) string {
	return c.DescribeDatabaseListWithContext(context.Background(), request)
}

func (c *Client) DescribeDatabaseListSend(request *DescribeDatabaseListRequest) (*DescribeDatabaseListResponse, error) {
	statusCode, msg, err := c.DescribeDatabaseListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDatabaseListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDatabaseListWithContext(ctx context.Context, request *DescribeDatabaseListRequest) string {
	if request == nil {
		request = NewDescribeDatabaseListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDatabaseListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDatabaseListWithContextV2(ctx context.Context, request *DescribeDatabaseListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDatabaseListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDatabaseListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeHistorySQLRequest() (request *DescribeHistorySQLRequest) {
	request = &DescribeHistorySQLRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "DescribeHistorySQL")
	return
}

func NewDescribeHistorySQLResponse() (response *DescribeHistorySQLResponse) {
	response = &DescribeHistorySQLResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeHistorySQL(request *DescribeHistorySQLRequest) string {
	return c.DescribeHistorySQLWithContext(context.Background(), request)
}

func (c *Client) DescribeHistorySQLSend(request *DescribeHistorySQLRequest) (*DescribeHistorySQLResponse, error) {
	statusCode, msg, err := c.DescribeHistorySQLWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeHistorySQLResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeHistorySQLWithContext(ctx context.Context, request *DescribeHistorySQLRequest) string {
	if request == nil {
		request = NewDescribeHistorySQLRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHistorySQLResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeHistorySQLWithContextV2(ctx context.Context, request *DescribeHistorySQLRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeHistorySQLRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHistorySQLResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStartExecuteSQLRequest() (request *StartExecuteSQLRequest) {
	request = &StartExecuteSQLRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "StartExecuteSQL")
	return
}

func NewStartExecuteSQLResponse() (response *StartExecuteSQLResponse) {
	response = &StartExecuteSQLResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StartExecuteSQL(request *StartExecuteSQLRequest) string {
	return c.StartExecuteSQLWithContext(context.Background(), request)
}

func (c *Client) StartExecuteSQLSend(request *StartExecuteSQLRequest) (*StartExecuteSQLResponse, error) {
	statusCode, msg, err := c.StartExecuteSQLWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StartExecuteSQLResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StartExecuteSQLWithContext(ctx context.Context, request *StartExecuteSQLRequest) string {
	if request == nil {
		request = NewStartExecuteSQLRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartExecuteSQLResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StartExecuteSQLWithContextV2(ctx context.Context, request *StartExecuteSQLRequest) (int, string, error) {
	if request == nil {
		request = NewStartExecuteSQLRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStartExecuteSQLResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateInstanceDatabaseRequest() (request *UpdateInstanceDatabaseRequest) {
	request = &UpdateInstanceDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "UpdateInstanceDatabase")
	return
}

func NewUpdateInstanceDatabaseResponse() (response *UpdateInstanceDatabaseResponse) {
	response = &UpdateInstanceDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateInstanceDatabase(request *UpdateInstanceDatabaseRequest) string {
	return c.UpdateInstanceDatabaseWithContext(context.Background(), request)
}

func (c *Client) UpdateInstanceDatabaseSend(request *UpdateInstanceDatabaseRequest) (*UpdateInstanceDatabaseResponse, error) {
	statusCode, msg, err := c.UpdateInstanceDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateInstanceDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateInstanceDatabaseWithContext(ctx context.Context, request *UpdateInstanceDatabaseRequest) string {
	if request == nil {
		request = NewUpdateInstanceDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateInstanceDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateInstanceDatabaseWithContextV2(ctx context.Context, request *UpdateInstanceDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateInstanceDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateInstanceDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateDatabaseTableRequest() (request *UpdateDatabaseTableRequest) {
	request = &UpdateDatabaseTableRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "UpdateDatabaseTable")
	return
}

func NewUpdateDatabaseTableResponse() (response *UpdateDatabaseTableResponse) {
	response = &UpdateDatabaseTableResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateDatabaseTable(request *UpdateDatabaseTableRequest) string {
	return c.UpdateDatabaseTableWithContext(context.Background(), request)
}

func (c *Client) UpdateDatabaseTableSend(request *UpdateDatabaseTableRequest) (*UpdateDatabaseTableResponse, error) {
	statusCode, msg, err := c.UpdateDatabaseTableWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateDatabaseTableResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateDatabaseTableWithContext(ctx context.Context, request *UpdateDatabaseTableRequest) string {
	if request == nil {
		request = NewUpdateDatabaseTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateDatabaseTableResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateDatabaseTableWithContextV2(ctx context.Context, request *UpdateDatabaseTableRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateDatabaseTableRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateDatabaseTableResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewTestInstanceConnectionRequest() (request *TestInstanceConnectionRequest) {
	request = &TestInstanceConnectionRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmp", APIVersion, "TestInstanceConnection")
	return
}

func NewTestInstanceConnectionResponse() (response *TestInstanceConnectionResponse) {
	response = &TestInstanceConnectionResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) TestInstanceConnection(request *TestInstanceConnectionRequest) string {
	return c.TestInstanceConnectionWithContext(context.Background(), request)
}

func (c *Client) TestInstanceConnectionSend(request *TestInstanceConnectionRequest) (*TestInstanceConnectionResponse, error) {
	statusCode, msg, err := c.TestInstanceConnectionWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct TestInstanceConnectionResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) TestInstanceConnectionWithContext(ctx context.Context, request *TestInstanceConnectionRequest) string {
	if request == nil {
		request = NewTestInstanceConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTestInstanceConnectionResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) TestInstanceConnectionWithContextV2(ctx context.Context, request *TestInstanceConnectionRequest) (int, string, error) {
	if request == nil {
		request = NewTestInstanceConnectionRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewTestInstanceConnectionResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


