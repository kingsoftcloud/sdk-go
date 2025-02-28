package v20180627

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
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

func (c *Client) DeleteCacheClusterWithContext(ctx context.Context, request *DeleteCacheClusterRequest) string {
	if request == nil {
		request = NewDeleteCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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

func (c *Client) ResizeCacheClusterWithContext(ctx context.Context, request *ResizeCacheClusterRequest) string {
	if request == nil {
		request = NewResizeCacheClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResizeCacheClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
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
