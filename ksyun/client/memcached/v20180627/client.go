package v20180627

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-06-27"

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
	request.Init().WithApiInfo("memcached", APIVersion, "CreateCacheCluster")
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
	request.Init().WithApiInfo("memcached", APIVersion, "DeleteCacheCluster")
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
func NewResizeCacheClusterRequest() (request *ResizeCacheClusterRequest) {
	request = &ResizeCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "ResizeCacheCluster")
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
func NewDescribeCacheClustersRequest() (request *DescribeCacheClustersRequest) {
	request = &DescribeCacheClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "DescribeCacheClusters")
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
func NewDescribeCacheClusterRequest() (request *DescribeCacheClusterRequest) {
	request = &DescribeCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "DescribeCacheCluster")
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
func NewFlushCacheClusterRequest() (request *FlushCacheClusterRequest) {
	request = &FlushCacheClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "FlushCacheCluster")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.Init().WithApiInfo("memcached", APIVersion, "RenameCacheCluster")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenameCacheClusterResponse()
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
	request.Init().WithApiInfo("memcached", APIVersion, "UpdatePassword")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdatePasswordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeCacheSecurityRulesRequest() (request *DescribeCacheSecurityRulesRequest) {
	request = &DescribeCacheSecurityRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "DescribeCacheSecurityRules")
	return
}

func NewDescribeCacheSecurityRulesResponse() (response *DescribeCacheSecurityRulesResponse) {
	response = &DescribeCacheSecurityRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCacheSecurityRules(request *DescribeCacheSecurityRulesRequest) string {
	return c.DescribeCacheSecurityRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeCacheSecurityRulesSend(request *DescribeCacheSecurityRulesRequest) (*DescribeCacheSecurityRulesResponse, error) {
	statusCode, msg, err := c.DescribeCacheSecurityRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCacheSecurityRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCacheSecurityRulesWithContext(ctx context.Context, request *DescribeCacheSecurityRulesRequest) string {
	if request == nil {
		request = NewDescribeCacheSecurityRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheSecurityRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCacheSecurityRulesWithContextV2(ctx context.Context, request *DescribeCacheSecurityRulesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCacheSecurityRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCacheSecurityRulesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteCacheSecurityRuleRequest() (request *DeleteCacheSecurityRuleRequest) {
	request = &DeleteCacheSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "DeleteCacheSecurityRule")
	return
}

func NewDeleteCacheSecurityRuleResponse() (response *DeleteCacheSecurityRuleResponse) {
	response = &DeleteCacheSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCacheSecurityRule(request *DeleteCacheSecurityRuleRequest) string {
	return c.DeleteCacheSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteCacheSecurityRuleSend(request *DeleteCacheSecurityRuleRequest) (*DeleteCacheSecurityRuleResponse, error) {
	statusCode, msg, err := c.DeleteCacheSecurityRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteCacheSecurityRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteCacheSecurityRuleWithContext(ctx context.Context, request *DeleteCacheSecurityRuleRequest) string {
	if request == nil {
		request = NewDeleteCacheSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCacheSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteCacheSecurityRuleWithContextV2(ctx context.Context, request *DeleteCacheSecurityRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteCacheSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCacheSecurityRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewSetCacheSecurityRulesRequest() (request *SetCacheSecurityRulesRequest) {
	request = &SetCacheSecurityRulesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("memcached", APIVersion, "SetCacheSecurityRules")
	return
}

func NewSetCacheSecurityRulesResponse() (response *SetCacheSecurityRulesResponse) {
	response = &SetCacheSecurityRulesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetCacheSecurityRules(request *SetCacheSecurityRulesRequest) string {
	return c.SetCacheSecurityRulesWithContext(context.Background(), request)
}

func (c *Client) SetCacheSecurityRulesSend(request *SetCacheSecurityRulesRequest) (*SetCacheSecurityRulesResponse, error) {
	statusCode, msg, err := c.SetCacheSecurityRulesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetCacheSecurityRulesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) SetCacheSecurityRulesWithContext(ctx context.Context, request *SetCacheSecurityRulesRequest) string {
	if request == nil {
		request = NewSetCacheSecurityRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetCacheSecurityRulesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) SetCacheSecurityRulesWithContextV2(ctx context.Context, request *SetCacheSecurityRulesRequest) (int, string, error) {
	if request == nil {
		request = NewSetCacheSecurityRulesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewSetCacheSecurityRulesResponse()
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
	request.Init().WithApiInfo("memcached", APIVersion, "DescribeRegions")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.Init().WithApiInfo("memcached", APIVersion, "DescribeAvailabilityZones")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeAvailabilityZonesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
