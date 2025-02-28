package v20160304

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2016-03-04"

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

func NewCreateCenRequest() (request *CreateCenRequest) {
	request = &CreateCenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCen")
	return
}

func NewCreateCenResponse() (response *CreateCenResponse) {
	response = &CreateCenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCen(request *CreateCenRequest) string {
	return c.CreateCenWithContext(context.Background(), request)
}

func (c *Client) CreateCenWithContext(ctx context.Context, request *CreateCenRequest) string {
	if request == nil {
		request = NewCreateCenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyCenRequest() (request *ModifyCenRequest) {
	request = &ModifyCenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "ModifyCen")
	return
}

func NewModifyCenResponse() (response *ModifyCenResponse) {
	response = &ModifyCenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCen(request *ModifyCenRequest) string {
	return c.ModifyCenWithContext(context.Background(), request)
}

func (c *Client) ModifyCenWithContext(ctx context.Context, request *ModifyCenRequest) string {
	if request == nil {
		request = NewModifyCenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteCenRequest() (request *DeleteCenRequest) {
	request = &DeleteCenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCen")
	return
}

func NewDeleteCenResponse() (response *DeleteCenResponse) {
	response = &DeleteCenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCen(request *DeleteCenRequest) string {
	return c.DeleteCenWithContext(context.Background(), request)
}

func (c *Client) DeleteCenWithContext(ctx context.Context, request *DeleteCenRequest) string {
	if request == nil {
		request = NewDeleteCenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCensRequest() (request *DescribeCensRequest) {
	request = &DescribeCensRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCens")
	return
}

func NewDescribeCensResponse() (response *DescribeCensResponse) {
	response = &DescribeCensResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCens(request *DescribeCensRequest) string {
	return c.DescribeCensWithContext(context.Background(), request)
}

func (c *Client) DescribeCensWithContext(ctx context.Context, request *DescribeCensRequest) string {
	if request == nil {
		request = NewDescribeCensRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCensResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAttachCenInstanceRequest() (request *AttachCenInstanceRequest) {
	request = &AttachCenInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "AttachCenInstance")
	return
}

func NewAttachCenInstanceResponse() (response *AttachCenInstanceResponse) {
	response = &AttachCenInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachCenInstance(request *AttachCenInstanceRequest) string {
	return c.AttachCenInstanceWithContext(context.Background(), request)
}

func (c *Client) AttachCenInstanceWithContext(ctx context.Context, request *AttachCenInstanceRequest) string {
	if request == nil {
		request = NewAttachCenInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachCenInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachCenInstanceRequest() (request *DetachCenInstanceRequest) {
	request = &DetachCenInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DetachCenInstance")
	return
}

func NewDetachCenInstanceResponse() (response *DetachCenInstanceResponse) {
	response = &DetachCenInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachCenInstance(request *DetachCenInstanceRequest) string {
	return c.DetachCenInstanceWithContext(context.Background(), request)
}

func (c *Client) DetachCenInstanceWithContext(ctx context.Context, request *DetachCenInstanceRequest) string {
	if request == nil {
		request = NewDetachCenInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachCenInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenInstancesRequest() (request *DescribeCenInstancesRequest) {
	request = &DescribeCenInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenInstances")
	return
}

func NewDescribeCenInstancesResponse() (response *DescribeCenInstancesResponse) {
	response = &DescribeCenInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenInstances(request *DescribeCenInstancesRequest) string {
	return c.DescribeCenInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeCenInstancesWithContext(ctx context.Context, request *DescribeCenInstancesRequest) string {
	if request == nil {
		request = NewDescribeCenInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreatCenGrantRequest() (request *CreatCenGrantRequest) {
	request = &CreatCenGrantRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreatCenGrant")
	return
}

func NewCreatCenGrantResponse() (response *CreatCenGrantResponse) {
	response = &CreatCenGrantResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreatCenGrant(request *CreatCenGrantRequest) string {
	return c.CreatCenGrantWithContext(context.Background(), request)
}

func (c *Client) CreatCenGrantWithContext(ctx context.Context, request *CreatCenGrantRequest) string {
	if request == nil {
		request = NewCreatCenGrantRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreatCenGrantResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteCenGrantRequest() (request *DeleteCenGrantRequest) {
	request = &DeleteCenGrantRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCenGrant")
	return
}

func NewDeleteCenGrantResponse() (response *DeleteCenGrantResponse) {
	response = &DeleteCenGrantResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCenGrant(request *DeleteCenGrantRequest) string {
	return c.DeleteCenGrantWithContext(context.Background(), request)
}

func (c *Client) DeleteCenGrantWithContext(ctx context.Context, request *DeleteCenGrantRequest) string {
	if request == nil {
		request = NewDeleteCenGrantRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenGrantResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenGrantsRequest() (request *DescribeCenGrantsRequest) {
	request = &DescribeCenGrantsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenGrants")
	return
}

func NewDescribeCenGrantsResponse() (response *DescribeCenGrantsResponse) {
	response = &DescribeCenGrantsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenGrants(request *DescribeCenGrantsRequest) string {
	return c.DescribeCenGrantsWithContext(context.Background(), request)
}

func (c *Client) DescribeCenGrantsWithContext(ctx context.Context, request *DescribeCenGrantsRequest) string {
	if request == nil {
		request = NewDescribeCenGrantsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenGrantsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeRegionGroupsRequest() (request *DescribeRegionGroupsRequest) {
	request = &DescribeRegionGroupsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeRegionGroups")
	return
}

func NewDescribeRegionGroupsResponse() (response *DescribeRegionGroupsResponse) {
	response = &DescribeRegionGroupsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRegionGroups(request *DescribeRegionGroupsRequest) string {
	return c.DescribeRegionGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionGroupsWithContext(ctx context.Context, request *DescribeRegionGroupsRequest) string {
	if request == nil {
		request = NewDescribeRegionGroupsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeRegionGroupsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateCenBandWidthPackageRequest() (request *CreateCenBandWidthPackageRequest) {
	request = &CreateCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCenBandWidthPackage")
	return
}

func NewCreateCenBandWidthPackageResponse() (response *CreateCenBandWidthPackageResponse) {
	response = &CreateCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCenBandWidthPackage(request *CreateCenBandWidthPackageRequest) string {
	return c.CreateCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) CreateCenBandWidthPackageWithContext(ctx context.Context, request *CreateCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewCreateCenBandWidthPackageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyCenBandWidthPackageRequest() (request *ModifyCenBandWidthPackageRequest) {
	request = &ModifyCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "ModifyCenBandWidthPackage")
	return
}

func NewModifyCenBandWidthPackageResponse() (response *ModifyCenBandWidthPackageResponse) {
	response = &ModifyCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCenBandWidthPackage(request *ModifyCenBandWidthPackageRequest) string {
	return c.ModifyCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) ModifyCenBandWidthPackageWithContext(ctx context.Context, request *ModifyCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewModifyCenBandWidthPackageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteCenBandWidthPackageRequest() (request *DeleteCenBandWidthPackageRequest) {
	request = &DeleteCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCenBandWidthPackage")
	return
}

func NewDeleteCenBandWidthPackageResponse() (response *DeleteCenBandWidthPackageResponse) {
	response = &DeleteCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCenBandWidthPackage(request *DeleteCenBandWidthPackageRequest) string {
	return c.DeleteCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) DeleteCenBandWidthPackageWithContext(ctx context.Context, request *DeleteCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewDeleteCenBandWidthPackageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAttachCenBandWidthPackageRequest() (request *AttachCenBandWidthPackageRequest) {
	request = &AttachCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "AttachCenBandWidthPackage")
	return
}

func NewAttachCenBandWidthPackageResponse() (response *AttachCenBandWidthPackageResponse) {
	response = &AttachCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachCenBandWidthPackage(request *AttachCenBandWidthPackageRequest) string {
	return c.AttachCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) AttachCenBandWidthPackageWithContext(ctx context.Context, request *AttachCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewAttachCenBandWidthPackageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachCenBandWidthPackageRequest() (request *DetachCenBandWidthPackageRequest) {
	request = &DetachCenBandWidthPackageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DetachCenBandWidthPackage")
	return
}

func NewDetachCenBandWidthPackageResponse() (response *DetachCenBandWidthPackageResponse) {
	response = &DetachCenBandWidthPackageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachCenBandWidthPackage(request *DetachCenBandWidthPackageRequest) string {
	return c.DetachCenBandWidthPackageWithContext(context.Background(), request)
}

func (c *Client) DetachCenBandWidthPackageWithContext(ctx context.Context, request *DetachCenBandWidthPackageRequest) string {
	if request == nil {
		request = NewDetachCenBandWidthPackageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachCenBandWidthPackageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenBandWidthPackagesRequest() (request *DescribeCenBandWidthPackagesRequest) {
	request = &DescribeCenBandWidthPackagesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackages")
	return
}

func NewDescribeCenBandWidthPackagesResponse() (response *DescribeCenBandWidthPackagesResponse) {
	response = &DescribeCenBandWidthPackagesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenBandWidthPackages(request *DescribeCenBandWidthPackagesRequest) string {
	return c.DescribeCenBandWidthPackagesWithContext(context.Background(), request)
}

func (c *Client) DescribeCenBandWidthPackagesWithContext(ctx context.Context, request *DescribeCenBandWidthPackagesRequest) string {
	if request == nil {
		request = NewDescribeCenBandWidthPackagesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenBandWidthPackagesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateCenRegionBandwidthRequest() (request *CreateCenRegionBandwidthRequest) {
	request = &CreateCenRegionBandwidthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCenRegionBandwidth")
	return
}

func NewCreateCenRegionBandwidthResponse() (response *CreateCenRegionBandwidthResponse) {
	response = &CreateCenRegionBandwidthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCenRegionBandwidth(request *CreateCenRegionBandwidthRequest) string {
	return c.CreateCenRegionBandwidthWithContext(context.Background(), request)
}

func (c *Client) CreateCenRegionBandwidthWithContext(ctx context.Context, request *CreateCenRegionBandwidthRequest) string {
	if request == nil {
		request = NewCreateCenRegionBandwidthRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenRegionBandwidthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteCenRegionBandwidthRequest() (request *DeleteCenRegionBandwidthRequest) {
	request = &DeleteCenRegionBandwidthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DeleteCenRegionBandwidth")
	return
}

func NewDeleteCenRegionBandwidthResponse() (response *DeleteCenRegionBandwidthResponse) {
	response = &DeleteCenRegionBandwidthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCenRegionBandwidth(request *DeleteCenRegionBandwidthRequest) string {
	return c.DeleteCenRegionBandwidthWithContext(context.Background(), request)
}

func (c *Client) DeleteCenRegionBandwidthWithContext(ctx context.Context, request *DeleteCenRegionBandwidthRequest) string {
	if request == nil {
		request = NewDeleteCenRegionBandwidthRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteCenRegionBandwidthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyCenRegionBandwidthRequest() (request *ModifyCenRegionBandwidthRequest) {
	request = &ModifyCenRegionBandwidthRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "ModifyCenRegionBandwidth")
	return
}

func NewModifyCenRegionBandwidthResponse() (response *ModifyCenRegionBandwidthResponse) {
	response = &ModifyCenRegionBandwidthResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCenRegionBandwidth(request *ModifyCenRegionBandwidthRequest) string {
	return c.ModifyCenRegionBandwidthWithContext(context.Background(), request)
}

func (c *Client) ModifyCenRegionBandwidthWithContext(ctx context.Context, request *ModifyCenRegionBandwidthRequest) string {
	if request == nil {
		request = NewModifyCenRegionBandwidthRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyCenRegionBandwidthResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenRegionBandwidthsRequest() (request *DescribeCenRegionBandwidthsRequest) {
	request = &DescribeCenRegionBandwidthsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRegionBandwidths")
	return
}

func NewDescribeCenRegionBandwidthsResponse() (response *DescribeCenRegionBandwidthsResponse) {
	response = &DescribeCenRegionBandwidthsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenRegionBandwidths(request *DescribeCenRegionBandwidthsRequest) string {
	return c.DescribeCenRegionBandwidthsWithContext(context.Background(), request)
}

func (c *Client) DescribeCenRegionBandwidthsWithContext(ctx context.Context, request *DescribeCenRegionBandwidthsRequest) string {
	if request == nil {
		request = NewDescribeCenRegionBandwidthsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenRegionBandwidthsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenRoutesRequest() (request *DescribeCenRoutesRequest) {
	request = &DescribeCenRoutesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRoutes")
	return
}

func NewDescribeCenRoutesResponse() (response *DescribeCenRoutesResponse) {
	response = &DescribeCenRoutesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenRoutes(request *DescribeCenRoutesRequest) string {
	return c.DescribeCenRoutesWithContext(context.Background(), request)
}

func (c *Client) DescribeCenRoutesWithContext(ctx context.Context, request *DescribeCenRoutesRequest) string {
	if request == nil {
		request = NewDescribeCenRoutesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenRoutesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenRegionsRequest() (request *DescribeCenRegionsRequest) {
	request = &DescribeCenRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenRegions")
	return
}

func NewDescribeCenRegionsResponse() (response *DescribeCenRegionsResponse) {
	response = &DescribeCenRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenRegions(request *DescribeCenRegionsRequest) string {
	return c.DescribeCenRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeCenRegionsWithContext(ctx context.Context, request *DescribeCenRegionsRequest) string {
	if request == nil {
		request = NewDescribeCenRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeCenBandWidthPackageUsageRequest() (request *DescribeCenBandWidthPackageUsageRequest) {
	request = &DescribeCenBandWidthPackageUsageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeCenBandWidthPackageUsage")
	return
}

func NewDescribeCenBandWidthPackageUsageResponse() (response *DescribeCenBandWidthPackageUsageResponse) {
	response = &DescribeCenBandWidthPackageUsageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCenBandWidthPackageUsage(request *DescribeCenBandWidthPackageUsageRequest) string {
	return c.DescribeCenBandWidthPackageUsageWithContext(context.Background(), request)
}

func (c *Client) DescribeCenBandWidthPackageUsageWithContext(ctx context.Context, request *DescribeCenBandWidthPackageUsageRequest) string {
	if request == nil {
		request = NewDescribeCenBandWidthPackageUsageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCenBandWidthPackageUsageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
