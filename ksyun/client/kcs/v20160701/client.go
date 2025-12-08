package v20160701
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2016-07-01"

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

func NewCreateCacheClusterRequest() (request *CreateCacheClusterRequest) {
	request = &CreateCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "CreateCacheCluster")
	return
}

func NewCreateCacheClusterResponse() (response *CreateCacheClusterResponse) {
	response = &CreateCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCacheCluster(request *CreateCacheClusterRequest) string {
	return c.CreateCacheClusterWithContext(context.Background(), request)
}

func (c *Client) CreateCacheClusterSend(request *CreateCacheClusterRequest) (*CreateCacheClusterResponse, error) {
	statusCode, msg, err := c.CreateCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCacheClusterWithContext(ctx context.Context, request *CreateCacheClusterRequest) string {
	if request == nil {
		request = NewCreateCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCacheClusterWithContextV2(ctx context.Context, request *CreateCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCacheClusterRequest() (request *DeleteCacheClusterRequest) {
	request = &DeleteCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteCacheCluster")
	return
}

func NewDeleteCacheClusterResponse() (response *DeleteCacheClusterResponse) {
	response = &DeleteCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCacheCluster(request *DeleteCacheClusterRequest) string {
	return c.DeleteCacheClusterWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheClusterSend(request *DeleteCacheClusterRequest) (*DeleteCacheClusterResponse, error) {
	statusCode, msg, err := c.DeleteCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCacheClusterWithContext(ctx context.Context, request *DeleteCacheClusterRequest) string {
	if request == nil {
		request = NewDeleteCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCacheClusterWithContextV2(ctx context.Context, request *DeleteCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheClusterRequest() (request *DescribeCacheClusterRequest) {
	request = &DescribeCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheCluster")
	return
}

func NewDescribeCacheClusterResponse() (response *DescribeCacheClusterResponse) {
	response = &DescribeCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheCluster(request *DescribeCacheClusterRequest) string {
	return c.DescribeCacheClusterWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheClusterSend(request *DescribeCacheClusterRequest) (*DescribeCacheClusterResponse, error) {
	statusCode, msg, err := c.DescribeCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheClusterWithContext(ctx context.Context, request *DescribeCacheClusterRequest) string {
	if request == nil {
		request = NewDescribeCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheClusterWithContextV2(ctx context.Context, request *DescribeCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheClustersRequest() (request *DescribeCacheClustersRequest) {
	request = &DescribeCacheClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheClusters")
	return
}

func NewDescribeCacheClustersResponse() (response *DescribeCacheClustersResponse) {
	response = &DescribeCacheClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheClusters(request *DescribeCacheClustersRequest) string {
	return c.DescribeCacheClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheClustersSend(request *DescribeCacheClustersRequest) (*DescribeCacheClustersResponse, error) {
	statusCode, msg, err := c.DescribeCacheClustersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheClustersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheClustersWithContext(ctx context.Context, request *DescribeCacheClustersRequest) string {
	if request == nil {
		request = NewDescribeCacheClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheClustersWithContextV2(ctx context.Context, request *DescribeCacheClustersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheClustersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewFlushCacheClusterRequest() (request *FlushCacheClusterRequest) {
	request = &FlushCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "FlushCacheCluster")
	return
}

func NewFlushCacheClusterResponse() (response *FlushCacheClusterResponse) {
	response = &FlushCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) FlushCacheCluster(request *FlushCacheClusterRequest) string {
	return c.FlushCacheClusterWithContext(context.Background(), request)
}

func (c *Client) FlushCacheClusterSend(request *FlushCacheClusterRequest) (*FlushCacheClusterResponse, error) {
	statusCode, msg, err := c.FlushCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct FlushCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) FlushCacheClusterWithContext(ctx context.Context, request *FlushCacheClusterRequest) string {
	if request == nil {
		request = NewFlushCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewFlushCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) FlushCacheClusterWithContextV2(ctx context.Context, request *FlushCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewFlushCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewFlushCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenameCacheClusterRequest() (request *RenameCacheClusterRequest) {
	request = &RenameCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "RenameCacheCluster")
	return
}

func NewRenameCacheClusterResponse() (response *RenameCacheClusterResponse) {
	response = &RenameCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameCacheCluster(request *RenameCacheClusterRequest) string {
	return c.RenameCacheClusterWithContext(context.Background(), request)
}

func (c *Client) RenameCacheClusterSend(request *RenameCacheClusterRequest) (*RenameCacheClusterResponse, error) {
	statusCode, msg, err := c.RenameCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RenameCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenameCacheClusterWithContext(ctx context.Context, request *RenameCacheClusterRequest) string {
	if request == nil {
		request = NewRenameCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenameCacheClusterWithContextV2(ctx context.Context, request *RenameCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewRenameCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResizeCacheClusterRequest() (request *ResizeCacheClusterRequest) {
	request = &ResizeCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "ResizeCacheCluster")
	return
}

func NewResizeCacheClusterResponse() (response *ResizeCacheClusterResponse) {
	response = &ResizeCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResizeCacheCluster(request *ResizeCacheClusterRequest) string {
	return c.ResizeCacheClusterWithContext(context.Background(), request)
}

func (c *Client) ResizeCacheClusterSend(request *ResizeCacheClusterRequest) (*ResizeCacheClusterResponse, error) {
	statusCode, msg, err := c.ResizeCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResizeCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResizeCacheClusterWithContext(ctx context.Context, request *ResizeCacheClusterRequest) string {
	if request == nil {
		request = NewResizeCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResizeCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResizeCacheClusterWithContextV2(ctx context.Context, request *ResizeCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewResizeCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewResizeCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheParametersRequest() (request *DescribeCacheParametersRequest) {
	request = &DescribeCacheParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheParameters")
	return
}

func NewDescribeCacheParametersResponse() (response *DescribeCacheParametersResponse) {
	response = &DescribeCacheParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheParameters(request *DescribeCacheParametersRequest) string {
	return c.DescribeCacheParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheParametersSend(request *DescribeCacheParametersRequest) (*DescribeCacheParametersResponse, error) {
	statusCode, msg, err := c.DescribeCacheParametersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheParametersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheParametersWithContext(ctx context.Context, request *DescribeCacheParametersRequest) string {
	if request == nil {
		request = NewDescribeCacheParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheParametersWithContextV2(ctx context.Context, request *DescribeCacheParametersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheParametersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetCacheParametersRequest() (request *SetCacheParametersRequest) {
	request = &SetCacheParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "SetCacheParameters")
	return
}

func NewSetCacheParametersResponse() (response *SetCacheParametersResponse) {
	response = &SetCacheParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetCacheParameters(request *SetCacheParametersRequest) string {
	return c.SetCacheParametersWithContext(context.Background(), request)
}

func (c *Client) SetCacheParametersSend(request *SetCacheParametersRequest) (*SetCacheParametersResponse, error) {
	statusCode, msg, err := c.SetCacheParametersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetCacheParametersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetCacheParametersWithContext(ctx context.Context, request *SetCacheParametersRequest) string {
	if request == nil {
		request = NewSetCacheParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetCacheParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetCacheParametersWithContextV2(ctx context.Context, request *SetCacheParametersRequest) (int, string, error) {
	if request == nil {
		request = NewSetCacheParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetCacheParametersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheDefaultParametersRequest() (request *DescribeCacheDefaultParametersRequest) {
	request = &DescribeCacheDefaultParametersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheDefaultParameters")
	return
}

func NewDescribeCacheDefaultParametersResponse() (response *DescribeCacheDefaultParametersResponse) {
	response = &DescribeCacheDefaultParametersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheDefaultParameters(request *DescribeCacheDefaultParametersRequest) string {
	return c.DescribeCacheDefaultParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheDefaultParametersSend(request *DescribeCacheDefaultParametersRequest) (*DescribeCacheDefaultParametersResponse, error) {
	statusCode, msg, err := c.DescribeCacheDefaultParametersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheDefaultParametersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheDefaultParametersWithContext(ctx context.Context, request *DescribeCacheDefaultParametersRequest) string {
	if request == nil {
		request = NewDescribeCacheDefaultParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheDefaultParametersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheDefaultParametersWithContextV2(ctx context.Context, request *DescribeCacheDefaultParametersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheDefaultParametersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheDefaultParametersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetCacheParameterGroupRequest() (request *SetCacheParameterGroupRequest) {
	request = &SetCacheParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "SetCacheParameterGroup")
	return
}

func NewSetCacheParameterGroupResponse() (response *SetCacheParameterGroupResponse) {
	response = &SetCacheParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetCacheParameterGroup(request *SetCacheParameterGroupRequest) string {
	return c.SetCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) SetCacheParameterGroupSend(request *SetCacheParameterGroupRequest) (*SetCacheParameterGroupResponse, error) {
	statusCode, msg, err := c.SetCacheParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetCacheParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetCacheParameterGroupWithContext(ctx context.Context, request *SetCacheParameterGroupRequest) string {
	if request == nil {
		request = NewSetCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetCacheParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetCacheParameterGroupWithContextV2(ctx context.Context, request *SetCacheParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewSetCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetCacheParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateCacheParameterGroupRequest() (request *CreateCacheParameterGroupRequest) {
	request = &CreateCacheParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "CreateCacheParameterGroup")
	return
}

func NewCreateCacheParameterGroupResponse() (response *CreateCacheParameterGroupResponse) {
	response = &CreateCacheParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCacheParameterGroup(request *CreateCacheParameterGroupRequest) string {
	return c.CreateCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) CreateCacheParameterGroupSend(request *CreateCacheParameterGroupRequest) (*CreateCacheParameterGroupResponse, error) {
	statusCode, msg, err := c.CreateCacheParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateCacheParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateCacheParameterGroupWithContext(ctx context.Context, request *CreateCacheParameterGroupRequest) string {
	if request == nil {
		request = NewCreateCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCacheParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateCacheParameterGroupWithContextV2(ctx context.Context, request *CreateCacheParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCacheParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCacheParameterGroupRequest() (request *DeleteCacheParameterGroupRequest) {
	request = &DeleteCacheParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteCacheParameterGroup")
	return
}

func NewDeleteCacheParameterGroupResponse() (response *DeleteCacheParameterGroupResponse) {
	response = &DeleteCacheParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCacheParameterGroup(request *DeleteCacheParameterGroupRequest) string {
	return c.DeleteCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheParameterGroupSend(request *DeleteCacheParameterGroupRequest) (*DeleteCacheParameterGroupResponse, error) {
	statusCode, msg, err := c.DeleteCacheParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteCacheParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCacheParameterGroupWithContext(ctx context.Context, request *DeleteCacheParameterGroupRequest) string {
	if request == nil {
		request = NewDeleteCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCacheParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCacheParameterGroupWithContextV2(ctx context.Context, request *DeleteCacheParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteCacheParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyCacheParameterGroupRequest() (request *ModifyCacheParameterGroupRequest) {
	request = &ModifyCacheParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "ModifyCacheParameterGroup")
	return
}

func NewModifyCacheParameterGroupResponse() (response *ModifyCacheParameterGroupResponse) {
	response = &ModifyCacheParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCacheParameterGroup(request *ModifyCacheParameterGroupRequest) string {
	return c.ModifyCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyCacheParameterGroupSend(request *ModifyCacheParameterGroupRequest) (*ModifyCacheParameterGroupResponse, error) {
	statusCode, msg, err := c.ModifyCacheParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyCacheParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyCacheParameterGroupWithContext(ctx context.Context, request *ModifyCacheParameterGroupRequest) string {
	if request == nil {
		request = NewModifyCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCacheParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyCacheParameterGroupWithContextV2(ctx context.Context, request *ModifyCacheParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifyCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewModifyCacheParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheParameterGroupsRequest() (request *DescribeCacheParameterGroupsRequest) {
	request = &DescribeCacheParameterGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheParameterGroups")
	return
}

func NewDescribeCacheParameterGroupsResponse() (response *DescribeCacheParameterGroupsResponse) {
	response = &DescribeCacheParameterGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheParameterGroups(request *DescribeCacheParameterGroupsRequest) string {
	return c.DescribeCacheParameterGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheParameterGroupsSend(request *DescribeCacheParameterGroupsRequest) (*DescribeCacheParameterGroupsResponse, error) {
	statusCode, msg, err := c.DescribeCacheParameterGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheParameterGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheParameterGroupsWithContext(ctx context.Context, request *DescribeCacheParameterGroupsRequest) string {
	if request == nil {
		request = NewDescribeCacheParameterGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheParameterGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheParameterGroupsWithContextV2(ctx context.Context, request *DescribeCacheParameterGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheParameterGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheParameterGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheParameterGroupRequest() (request *DescribeCacheParameterGroupRequest) {
	request = &DescribeCacheParameterGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheParameterGroup")
	return
}

func NewDescribeCacheParameterGroupResponse() (response *DescribeCacheParameterGroupResponse) {
	response = &DescribeCacheParameterGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheParameterGroup(request *DescribeCacheParameterGroupRequest) string {
	return c.DescribeCacheParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheParameterGroupSend(request *DescribeCacheParameterGroupRequest) (*DescribeCacheParameterGroupResponse, error) {
	statusCode, msg, err := c.DescribeCacheParameterGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheParameterGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheParameterGroupWithContext(ctx context.Context, request *DescribeCacheParameterGroupRequest) string {
	if request == nil {
		request = NewDescribeCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheParameterGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheParameterGroupWithContextV2(ctx context.Context, request *DescribeCacheParameterGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheParameterGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheParameterGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetTimingSnapshotRequest() (request *SetTimingSnapshotRequest) {
	request = &SetTimingSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "SetTimingSnapshot")
	return
}

func NewSetTimingSnapshotResponse() (response *SetTimingSnapshotResponse) {
	response = &SetTimingSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetTimingSnapshot(request *SetTimingSnapshotRequest) string {
	return c.SetTimingSnapshotWithContext(context.Background(), request)
}

func (c *Client) SetTimingSnapshotSend(request *SetTimingSnapshotRequest) (*SetTimingSnapshotResponse, error) {
	statusCode, msg, err := c.SetTimingSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetTimingSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetTimingSnapshotWithContext(ctx context.Context, request *SetTimingSnapshotRequest) string {
	if request == nil {
		request = NewSetTimingSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetTimingSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetTimingSnapshotWithContextV2(ctx context.Context, request *SetTimingSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewSetTimingSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetTimingSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSnapshotRequest() (request *DeleteSnapshotRequest) {
	request = &DeleteSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteSnapshot")
	return
}

func NewDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
	response = &DeleteSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) string {
	return c.DeleteSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteSnapshotSend(request *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {
	statusCode, msg, err := c.DeleteSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest) string {
	if request == nil {
		request = NewDeleteSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSnapshotWithContextV2(ctx context.Context, request *DeleteSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenameSnapshotRequest() (request *RenameSnapshotRequest) {
	request = &RenameSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "RenameSnapshot")
	return
}

func NewRenameSnapshotResponse() (response *RenameSnapshotResponse) {
	response = &RenameSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameSnapshot(request *RenameSnapshotRequest) string {
	return c.RenameSnapshotWithContext(context.Background(), request)
}

func (c *Client) RenameSnapshotSend(request *RenameSnapshotRequest) (*RenameSnapshotResponse, error) {
	statusCode, msg, err := c.RenameSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RenameSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenameSnapshotWithContext(ctx context.Context, request *RenameSnapshotRequest) string {
	if request == nil {
		request = NewRenameSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenameSnapshotWithContextV2(ctx context.Context, request *RenameSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewRenameSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRenameSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreSnapshotRequest() (request *RestoreSnapshotRequest) {
	request = &RestoreSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "RestoreSnapshot")
	return
}

func NewRestoreSnapshotResponse() (response *RestoreSnapshotResponse) {
	response = &RestoreSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreSnapshot(request *RestoreSnapshotRequest) string {
	return c.RestoreSnapshotWithContext(context.Background(), request)
}

func (c *Client) RestoreSnapshotSend(request *RestoreSnapshotRequest) (*RestoreSnapshotResponse, error) {
	statusCode, msg, err := c.RestoreSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RestoreSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreSnapshotWithContext(ctx context.Context, request *RestoreSnapshotRequest) string {
	if request == nil {
		request = NewRestoreSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreSnapshotWithContextV2(ctx context.Context, request *RestoreSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRestoreSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeSnapshots")
	return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) string {
	return c.DescribeSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeSnapshotsSend(request *DescribeSnapshotsRequest) (*DescribeSnapshotsResponse, error) {
	statusCode, msg, err := c.DescribeSnapshotsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSnapshotsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) string {
	if request == nil {
		request = NewDescribeSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSnapshotsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSnapshotsWithContextV2(ctx context.Context, request *DescribeSnapshotsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeSnapshotsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDownloadSnapshotRequest() (request *DownloadSnapshotRequest) {
	request = &DownloadSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DownloadSnapshot")
	return
}

func NewDownloadSnapshotResponse() (response *DownloadSnapshotResponse) {
	response = &DownloadSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DownloadSnapshot(request *DownloadSnapshotRequest) string {
	return c.DownloadSnapshotWithContext(context.Background(), request)
}

func (c *Client) DownloadSnapshotSend(request *DownloadSnapshotRequest) (*DownloadSnapshotResponse, error) {
	statusCode, msg, err := c.DownloadSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DownloadSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DownloadSnapshotWithContext(ctx context.Context, request *DownloadSnapshotRequest) string {
	if request == nil {
		request = NewDownloadSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDownloadSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DownloadSnapshotWithContextV2(ctx context.Context, request *DownloadSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDownloadSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDownloadSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewExportSnapshotRequest() (request *ExportSnapshotRequest) {
	request = &ExportSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "ExportSnapshot")
	return
}

func NewExportSnapshotResponse() (response *ExportSnapshotResponse) {
	response = &ExportSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ExportSnapshot(request *ExportSnapshotRequest) string {
	return c.ExportSnapshotWithContext(context.Background(), request)
}

func (c *Client) ExportSnapshotSend(request *ExportSnapshotRequest) (*ExportSnapshotResponse, error) {
	statusCode, msg, err := c.ExportSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ExportSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ExportSnapshotWithContext(ctx context.Context, request *ExportSnapshotRequest) string {
	if request == nil {
		request = NewExportSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewExportSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ExportSnapshotWithContextV2(ctx context.Context, request *ExportSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewExportSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewExportSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegions(request *DescribeRegionsRequest) string {
	return c.DescribeRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionsSend(request *DescribeRegionsRequest) (*DescribeRegionsResponse, error) {
	statusCode, msg, err := c.DescribeRegionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRegionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) string {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRegionsWithContextV2(ctx context.Context, request *DescribeRegionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRegionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAvailabilityZonesRequest() (request *DescribeAvailabilityZonesRequest) {
	request = &DescribeAvailabilityZonesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeAvailabilityZones")
	return
}

func NewDescribeAvailabilityZonesResponse() (response *DescribeAvailabilityZonesResponse) {
	response = &DescribeAvailabilityZonesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAvailabilityZones(request *DescribeAvailabilityZonesRequest) string {
	return c.DescribeAvailabilityZonesWithContext(context.Background(), request)
}

func (c *Client) DescribeAvailabilityZonesSend(request *DescribeAvailabilityZonesRequest) (*DescribeAvailabilityZonesResponse, error) {
	statusCode, msg, err := c.DescribeAvailabilityZonesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAvailabilityZonesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, request *DescribeAvailabilityZonesRequest) string {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAvailabilityZonesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAvailabilityZonesWithContextV2(ctx context.Context, request *DescribeAvailabilityZonesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAvailabilityZonesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAvailabilityZonesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DescribeCacheByRoleSend(request *DescribeCacheByRoleRequest) (*DescribeCacheByRoleResponse, error) {
	statusCode, msg, err := c.DescribeCacheByRoleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheByRoleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DescribeCacheByRoleWithContextV2(ctx context.Context, request *DescribeCacheByRoleRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheByRoleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheByRoleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewStatisticDBInstancesRequest() (request *StatisticDBInstancesRequest) {
	request = &StatisticDBInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "StatisticDBInstances")
	return
}

func NewStatisticDBInstancesResponse() (response *StatisticDBInstancesResponse) {
	response = &StatisticDBInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StatisticDBInstances(request *StatisticDBInstancesRequest) string {
	return c.StatisticDBInstancesWithContext(context.Background(), request)
}

func (c *Client) StatisticDBInstancesSend(request *StatisticDBInstancesRequest) (*StatisticDBInstancesResponse, error) {
	statusCode, msg, err := c.StatisticDBInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StatisticDBInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StatisticDBInstancesWithContext(ctx context.Context, request *StatisticDBInstancesRequest) string {
	if request == nil {
		request = NewStatisticDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStatisticDBInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StatisticDBInstancesWithContextV2(ctx context.Context, request *StatisticDBInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewStatisticDBInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStatisticDBInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdatePasswordRequest() (request *UpdatePasswordRequest) {
	request = &UpdatePasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "UpdatePassword")
	return
}

func NewUpdatePasswordResponse() (response *UpdatePasswordResponse) {
	response = &UpdatePasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdatePassword(request *UpdatePasswordRequest) string {
	return c.UpdatePasswordWithContext(context.Background(), request)
}

func (c *Client) UpdatePasswordSend(request *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	statusCode, msg, err := c.UpdatePasswordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdatePasswordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdatePasswordWithContext(ctx context.Context, request *UpdatePasswordRequest) string {
	if request == nil {
		request = NewUpdatePasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdatePasswordWithContextV2(ctx context.Context, request *UpdatePasswordRequest) (int, string, error) {
	if request == nil {
		request = NewUpdatePasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdatePasswordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestartCacheClusterRequest() (request *RestartCacheClusterRequest) {
	request = &RestartCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "RestartCacheCluster")
	return
}

func NewRestartCacheClusterResponse() (response *RestartCacheClusterResponse) {
	response = &RestartCacheClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestartCacheCluster(request *RestartCacheClusterRequest) string {
	return c.RestartCacheClusterWithContext(context.Background(), request)
}

func (c *Client) RestartCacheClusterSend(request *RestartCacheClusterRequest) (*RestartCacheClusterResponse, error) {
	statusCode, msg, err := c.RestartCacheClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RestartCacheClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestartCacheClusterWithContext(ctx context.Context, request *RestartCacheClusterRequest) string {
	if request == nil {
		request = NewRestartCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestartCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestartCacheClusterWithContextV2(ctx context.Context, request *RestartCacheClusterRequest) (int, string, error) {
	if request == nil {
		request = NewRestartCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestartCacheClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateEipRequest() (request *AllocateEipRequest) {
	request = &AllocateEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AllocateEip")
	return
}

func NewAllocateEipResponse() (response *AllocateEipResponse) {
	response = &AllocateEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateEip(request *AllocateEipRequest) string {
	return c.AllocateEipWithContext(context.Background(), request)
}

func (c *Client) AllocateEipSend(request *AllocateEipRequest) (*AllocateEipResponse, error) {
	statusCode, msg, err := c.AllocateEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AllocateEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateEipWithContext(ctx context.Context, request *AllocateEipRequest) string {
	if request == nil {
		request = NewAllocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateEipWithContextV2(ctx context.Context, request *AllocateEipRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeallocateEipRequest() (request *DeallocateEipRequest) {
	request = &DeallocateEipRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeallocateEip")
	return
}

func NewDeallocateEipResponse() (response *DeallocateEipResponse) {
	response = &DeallocateEipResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeallocateEip(request *DeallocateEipRequest) string {
	return c.DeallocateEipWithContext(context.Background(), request)
}

func (c *Client) DeallocateEipSend(request *DeallocateEipRequest) (*DeallocateEipResponse, error) {
	statusCode, msg, err := c.DeallocateEipWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeallocateEipResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeallocateEipWithContext(ctx context.Context, request *DeallocateEipRequest) string {
	if request == nil {
		request = NewDeallocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateEipResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeallocateEipWithContextV2(ctx context.Context, request *DeallocateEipRequest) (int, string, error) {
	if request == nil {
		request = NewDeallocateEipRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateEipResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) string {
	return c.DescribeInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesSend(request *DescribeInstancesRequest) (*DescribeInstancesResponse, error) {
	statusCode, msg, err := c.DescribeInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) string {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstancesWithContextV2(ctx context.Context, request *DescribeInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
	request = &DeleteSecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteSecurityGroupRule")
	return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
	response = &DeleteSecurityGroupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) string {
	return c.DeleteSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupRuleSend(request *DeleteSecurityGroupRuleRequest) (*DeleteSecurityGroupRuleResponse, error) {
	statusCode, msg, err := c.DeleteSecurityGroupRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSecurityGroupRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityGroupRuleWithContext(ctx context.Context, request *DeleteSecurityGroupRuleRequest) string {
	if request == nil {
		request = NewDeleteSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityGroupRuleWithContextV2(ctx context.Context, request *DeleteSecurityGroupRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityGroupRuleRequest() (request *CreateSecurityGroupRuleRequest) {
	request = &CreateSecurityGroupRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "CreateSecurityGroupRule")
	return
}

func NewCreateSecurityGroupRuleResponse() (response *CreateSecurityGroupRuleResponse) {
	response = &CreateSecurityGroupRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroupRule(request *CreateSecurityGroupRuleRequest) string {
	return c.CreateSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupRuleSend(request *CreateSecurityGroupRuleRequest) (*CreateSecurityGroupRuleResponse, error) {
	statusCode, msg, err := c.CreateSecurityGroupRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSecurityGroupRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityGroupRuleWithContext(ctx context.Context, request *CreateSecurityGroupRuleRequest) string {
	if request == nil {
		request = NewCreateSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityGroupRuleWithContextV2(ctx context.Context, request *CreateSecurityGroupRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityGroupRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeallocateSecurityGroupRequest() (request *DeallocateSecurityGroupRequest) {
	request = &DeallocateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeallocateSecurityGroup")
	return
}

func NewDeallocateSecurityGroupResponse() (response *DeallocateSecurityGroupResponse) {
	response = &DeallocateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeallocateSecurityGroup(request *DeallocateSecurityGroupRequest) string {
	return c.DeallocateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeallocateSecurityGroupSend(request *DeallocateSecurityGroupRequest) (*DeallocateSecurityGroupResponse, error) {
	statusCode, msg, err := c.DeallocateSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeallocateSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeallocateSecurityGroupWithContext(ctx context.Context, request *DeallocateSecurityGroupRequest) string {
	if request == nil {
		request = NewDeallocateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeallocateSecurityGroupWithContextV2(ctx context.Context, request *DeallocateSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeallocateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeallocateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAllocateSecurityGroupRequest() (request *AllocateSecurityGroupRequest) {
	request = &AllocateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AllocateSecurityGroup")
	return
}

func NewAllocateSecurityGroupResponse() (response *AllocateSecurityGroupResponse) {
	response = &AllocateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AllocateSecurityGroup(request *AllocateSecurityGroupRequest) string {
	return c.AllocateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) AllocateSecurityGroupSend(request *AllocateSecurityGroupRequest) (*AllocateSecurityGroupResponse, error) {
	statusCode, msg, err := c.AllocateSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AllocateSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AllocateSecurityGroupWithContext(ctx context.Context, request *AllocateSecurityGroupRequest) string {
	if request == nil {
		request = NewAllocateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AllocateSecurityGroupWithContextV2(ctx context.Context, request *AllocateSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewAllocateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAllocateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeSecurityGroup")
	return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
	response = &DescribeSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) string {
	return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupSend(request *DescribeSecurityGroupRequest) (*DescribeSecurityGroupResponse, error) {
	statusCode, msg, err := c.DescribeSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) string {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSecurityGroupWithContextV2(ctx context.Context, request *DescribeSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
	request = &DescribeSecurityGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeSecurityGroups")
	return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
	response = &DescribeSecurityGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) string {
	return c.DescribeSecurityGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupsSend(request *DescribeSecurityGroupsRequest) (*DescribeSecurityGroupsResponse, error) {
	statusCode, msg, err := c.DescribeSecurityGroupsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeSecurityGroupsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest) string {
	if request == nil {
		request = NewDescribeSecurityGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeSecurityGroupsWithContextV2(ctx context.Context, request *DescribeSecurityGroupsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeSecurityGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSecurityGroupsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
	request = &ModifySecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "ModifySecurityGroup")
	return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
	response = &ModifySecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) string {
	return c.ModifySecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupSend(request *ModifySecurityGroupRequest) (*ModifySecurityGroupResponse, error) {
	statusCode, msg, err := c.ModifySecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySecurityGroupWithContext(ctx context.Context, request *ModifySecurityGroupRequest) string {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySecurityGroupWithContextV2(ctx context.Context, request *ModifySecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteSecurityGroup")
	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
	response = &DeleteSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) string {
	return c.DeleteSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupSend(request *DeleteSecurityGroupRequest) (*DeleteSecurityGroupResponse, error) {
	statusCode, msg, err := c.DeleteSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) string {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityGroupWithContextV2(ctx context.Context, request *DeleteSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
	request = &CloneSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "CloneSecurityGroup")
	return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
	response = &CloneSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) string {
	return c.CloneSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CloneSecurityGroupSend(request *CloneSecurityGroupRequest) (*CloneSecurityGroupResponse, error) {
	statusCode, msg, err := c.CloneSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloneSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) string {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloneSecurityGroupWithContextV2(ctx context.Context, request *CloneSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloneSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "CreateSecurityGroup")
	return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) string {
	return c.CreateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupSend(request *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	statusCode, msg, err := c.CreateSecurityGroupWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSecurityGroupResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) string {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityGroupWithContextV2(ctx context.Context, request *CreateSecurityGroupRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheReadonlyNodeRequest() (request *DescribeCacheReadonlyNodeRequest) {
	request = &DescribeCacheReadonlyNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCacheReadonlyNode")
	return
}

func NewDescribeCacheReadonlyNodeResponse() (response *DescribeCacheReadonlyNodeResponse) {
	response = &DescribeCacheReadonlyNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheReadonlyNode(request *DescribeCacheReadonlyNodeRequest) string {
	return c.DescribeCacheReadonlyNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheReadonlyNodeSend(request *DescribeCacheReadonlyNodeRequest) (*DescribeCacheReadonlyNodeResponse, error) {
	statusCode, msg, err := c.DescribeCacheReadonlyNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheReadonlyNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheReadonlyNodeWithContext(ctx context.Context, request *DescribeCacheReadonlyNodeRequest) string {
	if request == nil {
		request = NewDescribeCacheReadonlyNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheReadonlyNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheReadonlyNodeWithContextV2(ctx context.Context, request *DescribeCacheReadonlyNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheReadonlyNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeCacheReadonlyNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAddCacheSlaveNodeRequest() (request *AddCacheSlaveNodeRequest) {
	request = &AddCacheSlaveNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AddCacheSlaveNode")
	return
}

func NewAddCacheSlaveNodeResponse() (response *AddCacheSlaveNodeResponse) {
	response = &AddCacheSlaveNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AddCacheSlaveNode(request *AddCacheSlaveNodeRequest) string {
	return c.AddCacheSlaveNodeWithContext(context.Background(), request)
}

func (c *Client) AddCacheSlaveNodeSend(request *AddCacheSlaveNodeRequest) (*AddCacheSlaveNodeResponse, error) {
	statusCode, msg, err := c.AddCacheSlaveNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AddCacheSlaveNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AddCacheSlaveNodeWithContext(ctx context.Context, request *AddCacheSlaveNodeRequest) string {
	if request == nil {
		request = NewAddCacheSlaveNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddCacheSlaveNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AddCacheSlaveNodeWithContextV2(ctx context.Context, request *AddCacheSlaveNodeRequest) (int, string, error) {
	if request == nil {
		request = NewAddCacheSlaveNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewAddCacheSlaveNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeHotKeysRequest() (request *DescribeHotKeysRequest) {
	request = &DescribeHotKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeHotKeys")
	return
}

func NewDescribeHotKeysResponse() (response *DescribeHotKeysResponse) {
	response = &DescribeHotKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeHotKeys(request *DescribeHotKeysRequest) string {
	return c.DescribeHotKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeHotKeysSend(request *DescribeHotKeysRequest) (*DescribeHotKeysResponse, error) {
	statusCode, msg, err := c.DescribeHotKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeHotKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeHotKeysWithContext(ctx context.Context, request *DescribeHotKeysRequest) string {
	if request == nil {
		request = NewDescribeHotKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHotKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeHotKeysWithContextV2(ctx context.Context, request *DescribeHotKeysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeHotKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeHotKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAnalyzeHotKeysRequest() (request *AnalyzeHotKeysRequest) {
	request = &AnalyzeHotKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeHotKeys")
	return
}

func NewAnalyzeHotKeysResponse() (response *AnalyzeHotKeysResponse) {
	response = &AnalyzeHotKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AnalyzeHotKeys(request *AnalyzeHotKeysRequest) string {
	return c.AnalyzeHotKeysWithContext(context.Background(), request)
}

func (c *Client) AnalyzeHotKeysSend(request *AnalyzeHotKeysRequest) (*AnalyzeHotKeysResponse, error) {
	statusCode, msg, err := c.AnalyzeHotKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AnalyzeHotKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AnalyzeHotKeysWithContext(ctx context.Context, request *AnalyzeHotKeysRequest) string {
	if request == nil {
		request = NewAnalyzeHotKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeHotKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AnalyzeHotKeysWithContextV2(ctx context.Context, request *AnalyzeHotKeysRequest) (int, string, error) {
	if request == nil {
		request = NewAnalyzeHotKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeHotKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCloseDirectAccessToClusterRequest() (request *CloseDirectAccessToClusterRequest) {
	request = &CloseDirectAccessToClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "CloseDirectAccessToCluster")
	return
}

func NewCloseDirectAccessToClusterResponse() (response *CloseDirectAccessToClusterResponse) {
	response = &CloseDirectAccessToClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloseDirectAccessToCluster(request *CloseDirectAccessToClusterRequest) string {
	return c.CloseDirectAccessToClusterWithContext(context.Background(), request)
}

func (c *Client) CloseDirectAccessToClusterSend(request *CloseDirectAccessToClusterRequest) (*CloseDirectAccessToClusterResponse, error) {
	statusCode, msg, err := c.CloseDirectAccessToClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloseDirectAccessToClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloseDirectAccessToClusterWithContext(ctx context.Context, request *CloseDirectAccessToClusterRequest) string {
	if request == nil {
		request = NewCloseDirectAccessToClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloseDirectAccessToClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloseDirectAccessToClusterWithContextV2(ctx context.Context, request *CloseDirectAccessToClusterRequest) (int, string, error) {
	if request == nil {
		request = NewCloseDirectAccessToClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloseDirectAccessToClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewOpenDirectAccessToClusterRequest() (request *OpenDirectAccessToClusterRequest) {
	request = &OpenDirectAccessToClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "OpenDirectAccessToCluster")
	return
}

func NewOpenDirectAccessToClusterResponse() (response *OpenDirectAccessToClusterResponse) {
	response = &OpenDirectAccessToClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) OpenDirectAccessToCluster(request *OpenDirectAccessToClusterRequest) string {
	return c.OpenDirectAccessToClusterWithContext(context.Background(), request)
}

func (c *Client) OpenDirectAccessToClusterSend(request *OpenDirectAccessToClusterRequest) (*OpenDirectAccessToClusterResponse, error) {
	statusCode, msg, err := c.OpenDirectAccessToClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct OpenDirectAccessToClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) OpenDirectAccessToClusterWithContext(ctx context.Context, request *OpenDirectAccessToClusterRequest) string {
	if request == nil {
		request = NewOpenDirectAccessToClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOpenDirectAccessToClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) OpenDirectAccessToClusterWithContextV2(ctx context.Context, request *OpenDirectAccessToClusterRequest) (int, string, error) {
	if request == nil {
		request = NewOpenDirectAccessToClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewOpenDirectAccessToClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeParentBackUpsSnapshotsRequest() (request *DescribeParentBackUpsSnapshotsRequest) {
	request = &DescribeParentBackUpsSnapshotsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeParentBackUpsSnapshots")
	return
}

func NewDescribeParentBackUpsSnapshotsResponse() (response *DescribeParentBackUpsSnapshotsResponse) {
	response = &DescribeParentBackUpsSnapshotsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeParentBackUpsSnapshots(request *DescribeParentBackUpsSnapshotsRequest) string {
	return c.DescribeParentBackUpsSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DescribeParentBackUpsSnapshotsSend(request *DescribeParentBackUpsSnapshotsRequest) (*DescribeParentBackUpsSnapshotsResponse, error) {
	statusCode, msg, err := c.DescribeParentBackUpsSnapshotsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeParentBackUpsSnapshotsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeParentBackUpsSnapshotsWithContext(ctx context.Context, request *DescribeParentBackUpsSnapshotsRequest) string {
	if request == nil {
		request = NewDescribeParentBackUpsSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeParentBackUpsSnapshotsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeParentBackUpsSnapshotsWithContextV2(ctx context.Context, request *DescribeParentBackUpsSnapshotsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeParentBackUpsSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeParentBackUpsSnapshotsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBackUpsSnapshotsDetailRequest() (request *DescribeBackUpsSnapshotsDetailRequest) {
	request = &DescribeBackUpsSnapshotsDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeBackUpsSnapshotsDetail")
	return
}

func NewDescribeBackUpsSnapshotsDetailResponse() (response *DescribeBackUpsSnapshotsDetailResponse) {
	response = &DescribeBackUpsSnapshotsDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBackUpsSnapshotsDetail(request *DescribeBackUpsSnapshotsDetailRequest) string {
	return c.DescribeBackUpsSnapshotsDetailWithContext(context.Background(), request)
}

func (c *Client) DescribeBackUpsSnapshotsDetailSend(request *DescribeBackUpsSnapshotsDetailRequest) (*DescribeBackUpsSnapshotsDetailResponse, error) {
	statusCode, msg, err := c.DescribeBackUpsSnapshotsDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeBackUpsSnapshotsDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBackUpsSnapshotsDetailWithContext(ctx context.Context, request *DescribeBackUpsSnapshotsDetailRequest) string {
	if request == nil {
		request = NewDescribeBackUpsSnapshotsDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackUpsSnapshotsDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBackUpsSnapshotsDetailWithContextV2(ctx context.Context, request *DescribeBackUpsSnapshotsDetailRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBackUpsSnapshotsDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBackUpsSnapshotsDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteLevelSnapshotsRequest() (request *DeleteLevelSnapshotsRequest) {
	request = &DeleteLevelSnapshotsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteLevelSnapshots")
	return
}

func NewDeleteLevelSnapshotsResponse() (response *DeleteLevelSnapshotsResponse) {
	response = &DeleteLevelSnapshotsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLevelSnapshots(request *DeleteLevelSnapshotsRequest) string {
	return c.DeleteLevelSnapshotsWithContext(context.Background(), request)
}

func (c *Client) DeleteLevelSnapshotsSend(request *DeleteLevelSnapshotsRequest) (*DeleteLevelSnapshotsResponse, error) {
	statusCode, msg, err := c.DeleteLevelSnapshotsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteLevelSnapshotsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteLevelSnapshotsWithContext(ctx context.Context, request *DeleteLevelSnapshotsRequest) string {
	if request == nil {
		request = NewDeleteLevelSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLevelSnapshotsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteLevelSnapshotsWithContextV2(ctx context.Context, request *DeleteLevelSnapshotsRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLevelSnapshotsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteLevelSnapshotsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDownloadLevelSnapshotRequest() (request *DownloadLevelSnapshotRequest) {
	request = &DownloadLevelSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DownloadLevelSnapshot")
	return
}

func NewDownloadLevelSnapshotResponse() (response *DownloadLevelSnapshotResponse) {
	response = &DownloadLevelSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DownloadLevelSnapshot(request *DownloadLevelSnapshotRequest) string {
	return c.DownloadLevelSnapshotWithContext(context.Background(), request)
}

func (c *Client) DownloadLevelSnapshotSend(request *DownloadLevelSnapshotRequest) (*DownloadLevelSnapshotResponse, error) {
	statusCode, msg, err := c.DownloadLevelSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DownloadLevelSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DownloadLevelSnapshotWithContext(ctx context.Context, request *DownloadLevelSnapshotRequest) string {
	if request == nil {
		request = NewDownloadLevelSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDownloadLevelSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DownloadLevelSnapshotWithContextV2(ctx context.Context, request *DownloadLevelSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewDownloadLevelSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDownloadLevelSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBigKeysRequest() (request *DescribeBigKeysRequest) {
	request = &DescribeBigKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeBigKeys")
	return
}

func NewDescribeBigKeysResponse() (response *DescribeBigKeysResponse) {
	response = &DescribeBigKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBigKeys(request *DescribeBigKeysRequest) string {
	return c.DescribeBigKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeBigKeysSend(request *DescribeBigKeysRequest) (*DescribeBigKeysResponse, error) {
	statusCode, msg, err := c.DescribeBigKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeBigKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBigKeysWithContext(ctx context.Context, request *DescribeBigKeysRequest) string {
	if request == nil {
		request = NewDescribeBigKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBigKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBigKeysWithContextV2(ctx context.Context, request *DescribeBigKeysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBigKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBigKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteBigKeysAnalyseResultRequest() (request *DeleteBigKeysAnalyseResultRequest) {
	request = &DeleteBigKeysAnalyseResultRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DeleteBigKeysAnalyseResult")
	return
}

func NewDeleteBigKeysAnalyseResultResponse() (response *DeleteBigKeysAnalyseResultResponse) {
	response = &DeleteBigKeysAnalyseResultResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBigKeysAnalyseResult(request *DeleteBigKeysAnalyseResultRequest) string {
	return c.DeleteBigKeysAnalyseResultWithContext(context.Background(), request)
}

func (c *Client) DeleteBigKeysAnalyseResultSend(request *DeleteBigKeysAnalyseResultRequest) (*DeleteBigKeysAnalyseResultResponse, error) {
	statusCode, msg, err := c.DeleteBigKeysAnalyseResultWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteBigKeysAnalyseResultResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteBigKeysAnalyseResultWithContext(ctx context.Context, request *DeleteBigKeysAnalyseResultRequest) string {
	if request == nil {
		request = NewDeleteBigKeysAnalyseResultRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBigKeysAnalyseResultResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteBigKeysAnalyseResultWithContextV2(ctx context.Context, request *DeleteBigKeysAnalyseResultRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteBigKeysAnalyseResultRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBigKeysAnalyseResultResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAnalyzeBigKeysRequest() (request *AnalyzeBigKeysRequest) {
	request = &AnalyzeBigKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeBigKeys")
	return
}

func NewAnalyzeBigKeysResponse() (response *AnalyzeBigKeysResponse) {
	response = &AnalyzeBigKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AnalyzeBigKeys(request *AnalyzeBigKeysRequest) string {
	return c.AnalyzeBigKeysWithContext(context.Background(), request)
}

func (c *Client) AnalyzeBigKeysSend(request *AnalyzeBigKeysRequest) (*AnalyzeBigKeysResponse, error) {
	statusCode, msg, err := c.AnalyzeBigKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AnalyzeBigKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AnalyzeBigKeysWithContext(ctx context.Context, request *AnalyzeBigKeysRequest) string {
	if request == nil {
		request = NewAnalyzeBigKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeBigKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AnalyzeBigKeysWithContextV2(ctx context.Context, request *AnalyzeBigKeysRequest) (int, string, error) {
	if request == nil {
		request = NewAnalyzeBigKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeBigKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCreateSnapshotStatusRequest() (request *DescribeCreateSnapshotStatusRequest) {
	request = &DescribeCreateSnapshotStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeCreateSnapshotStatus")
	return
}

func NewDescribeCreateSnapshotStatusResponse() (response *DescribeCreateSnapshotStatusResponse) {
	response = &DescribeCreateSnapshotStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCreateSnapshotStatus(request *DescribeCreateSnapshotStatusRequest) string {
	return c.DescribeCreateSnapshotStatusWithContext(context.Background(), request)
}

func (c *Client) DescribeCreateSnapshotStatusSend(request *DescribeCreateSnapshotStatusRequest) (*DescribeCreateSnapshotStatusResponse, error) {
	statusCode, msg, err := c.DescribeCreateSnapshotStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCreateSnapshotStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCreateSnapshotStatusWithContext(ctx context.Context, request *DescribeCreateSnapshotStatusRequest) string {
	if request == nil {
		request = NewDescribeCreateSnapshotStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCreateSnapshotStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCreateSnapshotStatusWithContextV2(ctx context.Context, request *DescribeCreateSnapshotStatusRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCreateSnapshotStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCreateSnapshotStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetDailyAnalyzeSwitchStateRequest() (request *GetDailyAnalyzeSwitchStateRequest) {
	request = &GetDailyAnalyzeSwitchStateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "GetDailyAnalyzeSwitchState")
	return
}

func NewGetDailyAnalyzeSwitchStateResponse() (response *GetDailyAnalyzeSwitchStateResponse) {
	response = &GetDailyAnalyzeSwitchStateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDailyAnalyzeSwitchState(request *GetDailyAnalyzeSwitchStateRequest) string {
	return c.GetDailyAnalyzeSwitchStateWithContext(context.Background(), request)
}

func (c *Client) GetDailyAnalyzeSwitchStateSend(request *GetDailyAnalyzeSwitchStateRequest) (*GetDailyAnalyzeSwitchStateResponse, error) {
	statusCode, msg, err := c.GetDailyAnalyzeSwitchStateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDailyAnalyzeSwitchStateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDailyAnalyzeSwitchStateWithContext(ctx context.Context, request *GetDailyAnalyzeSwitchStateRequest) string {
	if request == nil {
		request = NewGetDailyAnalyzeSwitchStateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDailyAnalyzeSwitchStateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDailyAnalyzeSwitchStateWithContextV2(ctx context.Context, request *GetDailyAnalyzeSwitchStateRequest) (int, string, error) {
	if request == nil {
		request = NewGetDailyAnalyzeSwitchStateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDailyAnalyzeSwitchStateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAnalyzeDailyRequest() (request *AnalyzeDailyRequest) {
	request = &AnalyzeDailyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeDaily")
	return
}

func NewAnalyzeDailyResponse() (response *AnalyzeDailyResponse) {
	response = &AnalyzeDailyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AnalyzeDaily(request *AnalyzeDailyRequest) string {
	return c.AnalyzeDailyWithContext(context.Background(), request)
}

func (c *Client) AnalyzeDailySend(request *AnalyzeDailyRequest) (*AnalyzeDailyResponse, error) {
	statusCode, msg, err := c.AnalyzeDailyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AnalyzeDailyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AnalyzeDailyWithContext(ctx context.Context, request *AnalyzeDailyRequest) string {
	if request == nil {
		request = NewAnalyzeDailyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeDailyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AnalyzeDailyWithContextV2(ctx context.Context, request *AnalyzeDailyRequest) (int, string, error) {
	if request == nil {
		request = NewAnalyzeDailyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeDailyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAnalyzeSlowDailyRequest() (request *AnalyzeSlowDailyRequest) {
	request = &AnalyzeSlowDailyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeSlowDaily")
	return
}

func NewAnalyzeSlowDailyResponse() (response *AnalyzeSlowDailyResponse) {
	response = &AnalyzeSlowDailyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AnalyzeSlowDaily(request *AnalyzeSlowDailyRequest) string {
	return c.AnalyzeSlowDailyWithContext(context.Background(), request)
}

func (c *Client) AnalyzeSlowDailySend(request *AnalyzeSlowDailyRequest) (*AnalyzeSlowDailyResponse, error) {
	statusCode, msg, err := c.AnalyzeSlowDailyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AnalyzeSlowDailyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AnalyzeSlowDailyWithContext(ctx context.Context, request *AnalyzeSlowDailyRequest) string {
	if request == nil {
		request = NewAnalyzeSlowDailyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeSlowDailyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AnalyzeSlowDailyWithContextV2(ctx context.Context, request *AnalyzeSlowDailyRequest) (int, string, error) {
	if request == nil {
		request = NewAnalyzeSlowDailyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeSlowDailyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAnalyzeDailySwitchRequest() (request *AnalyzeDailySwitchRequest) {
	request = &AnalyzeDailySwitchRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "AnalyzeDailySwitch")
	return
}

func NewAnalyzeDailySwitchResponse() (response *AnalyzeDailySwitchResponse) {
	response = &AnalyzeDailySwitchResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AnalyzeDailySwitch(request *AnalyzeDailySwitchRequest) string {
	return c.AnalyzeDailySwitchWithContext(context.Background(), request)
}

func (c *Client) AnalyzeDailySwitchSend(request *AnalyzeDailySwitchRequest) (*AnalyzeDailySwitchResponse, error) {
	statusCode, msg, err := c.AnalyzeDailySwitchWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AnalyzeDailySwitchResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AnalyzeDailySwitchWithContext(ctx context.Context, request *AnalyzeDailySwitchRequest) string {
	if request == nil {
		request = NewAnalyzeDailySwitchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeDailySwitchResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AnalyzeDailySwitchWithContextV2(ctx context.Context, request *AnalyzeDailySwitchRequest) (int, string, error) {
	if request == nil {
		request = NewAnalyzeDailySwitchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAnalyzeDailySwitchResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRestoreByTimePointSwitchRequest() (request *RestoreByTimePointSwitchRequest) {
	request = &RestoreByTimePointSwitchRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "RestoreByTimePointSwitch")
	return
}

func NewRestoreByTimePointSwitchResponse() (response *RestoreByTimePointSwitchResponse) {
	response = &RestoreByTimePointSwitchResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RestoreByTimePointSwitch(request *RestoreByTimePointSwitchRequest) string {
	return c.RestoreByTimePointSwitchWithContext(context.Background(), request)
}

func (c *Client) RestoreByTimePointSwitchSend(request *RestoreByTimePointSwitchRequest) (*RestoreByTimePointSwitchResponse, error) {
	statusCode, msg, err := c.RestoreByTimePointSwitchWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RestoreByTimePointSwitchResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RestoreByTimePointSwitchWithContext(ctx context.Context, request *RestoreByTimePointSwitchRequest) string {
	if request == nil {
		request = NewRestoreByTimePointSwitchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestoreByTimePointSwitchResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RestoreByTimePointSwitchWithContextV2(ctx context.Context, request *RestoreByTimePointSwitchRequest) (int, string, error) {
	if request == nil {
		request = NewRestoreByTimePointSwitchRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRestoreByTimePointSwitchResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRestoreTimePointsRequest() (request *DescribeRestoreTimePointsRequest) {
	request = &DescribeRestoreTimePointsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeRestoreTimePoints")
	return
}

func NewDescribeRestoreTimePointsResponse() (response *DescribeRestoreTimePointsResponse) {
	response = &DescribeRestoreTimePointsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRestoreTimePoints(request *DescribeRestoreTimePointsRequest) string {
	return c.DescribeRestoreTimePointsWithContext(context.Background(), request)
}

func (c *Client) DescribeRestoreTimePointsSend(request *DescribeRestoreTimePointsRequest) (*DescribeRestoreTimePointsResponse, error) {
	statusCode, msg, err := c.DescribeRestoreTimePointsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRestoreTimePointsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRestoreTimePointsWithContext(ctx context.Context, request *DescribeRestoreTimePointsRequest) string {
	if request == nil {
		request = NewDescribeRestoreTimePointsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRestoreTimePointsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRestoreTimePointsWithContextV2(ctx context.Context, request *DescribeRestoreTimePointsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRestoreTimePointsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRestoreTimePointsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeBigHotKeysRequest() (request *DescribeBigHotKeysRequest) {
	request = &DescribeBigHotKeysRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribeBigHotKeys")
	return
}

func NewDescribeBigHotKeysResponse() (response *DescribeBigHotKeysResponse) {
	response = &DescribeBigHotKeysResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBigHotKeys(request *DescribeBigHotKeysRequest) string {
	return c.DescribeBigHotKeysWithContext(context.Background(), request)
}

func (c *Client) DescribeBigHotKeysSend(request *DescribeBigHotKeysRequest) (*DescribeBigHotKeysResponse, error) {
	statusCode, msg, err := c.DescribeBigHotKeysWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeBigHotKeysResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeBigHotKeysWithContext(ctx context.Context, request *DescribeBigHotKeysRequest) string {
	if request == nil {
		request = NewDescribeBigHotKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBigHotKeysResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeBigHotKeysWithContextV2(ctx context.Context, request *DescribeBigHotKeysRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeBigHotKeysRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBigHotKeysResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribePluginsRequest() (request *DescribePluginsRequest) {
	request = &DescribePluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "DescribePlugins")
	return
}

func NewDescribePluginsResponse() (response *DescribePluginsResponse) {
	response = &DescribePluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribePlugins(request *DescribePluginsRequest) string {
	return c.DescribePluginsWithContext(context.Background(), request)
}

func (c *Client) DescribePluginsSend(request *DescribePluginsRequest) (*DescribePluginsResponse, error) {
	statusCode, msg, err := c.DescribePluginsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribePluginsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribePluginsWithContext(ctx context.Context, request *DescribePluginsRequest) string {
	if request == nil {
		request = NewDescribePluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribePluginsWithContextV2(ctx context.Context, request *DescribePluginsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribePluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribePluginsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewInstallPluginsRequest() (request *InstallPluginsRequest) {
	request = &InstallPluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "InstallPlugins")
	return
}

func NewInstallPluginsResponse() (response *InstallPluginsResponse) {
	response = &InstallPluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) InstallPlugins(request *InstallPluginsRequest) string {
	return c.InstallPluginsWithContext(context.Background(), request)
}

func (c *Client) InstallPluginsSend(request *InstallPluginsRequest) (*InstallPluginsResponse, error) {
	statusCode, msg, err := c.InstallPluginsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct InstallPluginsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) InstallPluginsWithContext(ctx context.Context, request *InstallPluginsRequest) string {
	if request == nil {
		request = NewInstallPluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewInstallPluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) InstallPluginsWithContextV2(ctx context.Context, request *InstallPluginsRequest) (int, string, error) {
	if request == nil {
		request = NewInstallPluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewInstallPluginsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUninstallPluginsRequest() (request *UninstallPluginsRequest) {
	request = &UninstallPluginsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("kcs", APIVersion, "UninstallPlugins")
	return
}

func NewUninstallPluginsResponse() (response *UninstallPluginsResponse) {
	response = &UninstallPluginsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UninstallPlugins(request *UninstallPluginsRequest) string {
	return c.UninstallPluginsWithContext(context.Background(), request)
}

func (c *Client) UninstallPluginsSend(request *UninstallPluginsRequest) (*UninstallPluginsResponse, error) {
	statusCode, msg, err := c.UninstallPluginsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UninstallPluginsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UninstallPluginsWithContext(ctx context.Context, request *UninstallPluginsRequest) string {
	if request == nil {
		request = NewUninstallPluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUninstallPluginsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UninstallPluginsWithContextV2(ctx context.Context, request *UninstallPluginsRequest) (int, string, error) {
	if request == nil {
		request = NewUninstallPluginsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUninstallPluginsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


