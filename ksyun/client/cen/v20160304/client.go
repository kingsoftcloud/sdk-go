package v20160304
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewDescribeCenBandWidthPackageUsageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeNetworkInstancesRequest() (request *DescribeNetworkInstancesRequest) {
	request = &DescribeNetworkInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeNetworkInstances")
	return
}

func NewDescribeNetworkInstancesResponse() (response *DescribeNetworkInstancesResponse) {
	response = &DescribeNetworkInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetworkInstances(request *DescribeNetworkInstancesRequest) string {
	return c.DescribeNetworkInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeNetworkInstancesWithContext(ctx context.Context, request *DescribeNetworkInstancesRequest) string {
	if request == nil {
		request = NewDescribeNetworkInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeNetworkInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateCenGrantRequest() (request *CreateCenGrantRequest) {
	request = &CreateCenGrantRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CreateCenGrant")
	return
}

func NewCreateCenGrantResponse() (response *CreateCenGrantResponse) {
	response = &CreateCenGrantResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCenGrant(request *CreateCenGrantRequest) string {
	return c.CreateCenGrantWithContext(context.Background(), request)
}

func (c *Client) CreateCenGrantWithContext(ctx context.Context, request *CreateCenGrantRequest) string {
	if request == nil {
		request = NewCreateCenGrantRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateCenGrantResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInterAreasRequest() (request *DescribeInterAreasRequest) {
	request = &DescribeInterAreasRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeInterAreas")
	return
}

func NewDescribeInterAreasResponse() (response *DescribeInterAreasResponse) {
	response = &DescribeInterAreasResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInterAreas(request *DescribeInterAreasRequest) string {
	return c.DescribeInterAreasWithContext(context.Background(), request)
}

func (c *Client) DescribeInterAreasWithContext(ctx context.Context, request *DescribeInterAreasRequest) string {
	if request == nil {
		request = NewDescribeInterAreasRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInterAreasResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeInterRegionsRequest() (request *DescribeInterRegionsRequest) {
	request = &DescribeInterRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DescribeInterRegions")
	return
}

func NewDescribeInterRegionsResponse() (response *DescribeInterRegionsResponse) {
	response = &DescribeInterRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInterRegions(request *DescribeInterRegionsRequest) string {
	return c.DescribeInterRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeInterRegionsWithContext(ctx context.Context, request *DescribeInterRegionsRequest) string {
	if request == nil {
		request = NewDescribeInterRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInterRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAttachNetworkInstanceRequest() (request *AttachNetworkInstanceRequest) {
	request = &AttachNetworkInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "AttachNetworkInstance")
	return
}

func NewAttachNetworkInstanceResponse() (response *AttachNetworkInstanceResponse) {
	response = &AttachNetworkInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachNetworkInstance(request *AttachNetworkInstanceRequest) string {
	return c.AttachNetworkInstanceWithContext(context.Background(), request)
}

func (c *Client) AttachNetworkInstanceWithContext(ctx context.Context, request *AttachNetworkInstanceRequest) string {
	if request == nil {
		request = NewAttachNetworkInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachNetworkInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDetachNetworkInstanceRequest() (request *DetachNetworkInstanceRequest) {
	request = &DetachNetworkInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "DetachNetworkInstance")
	return
}

func NewDetachNetworkInstanceResponse() (response *DetachNetworkInstanceResponse) {
	response = &DetachNetworkInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachNetworkInstance(request *DetachNetworkInstanceRequest) string {
	return c.DetachNetworkInstanceWithContext(context.Background(), request)
}

func (c *Client) DetachNetworkInstanceWithContext(ctx context.Context, request *DetachNetworkInstanceRequest) string {
	if request == nil {
		request = NewDetachNetworkInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachNetworkInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCenCidrPublishRequest() (request *CenCidrPublishRequest) {
	request = &CenCidrPublishRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CenCidrPublish")
	return
}

func NewCenCidrPublishResponse() (response *CenCidrPublishResponse) {
	response = &CenCidrPublishResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CenCidrPublish(request *CenCidrPublishRequest) string {
	return c.CenCidrPublishWithContext(context.Background(), request)
}

func (c *Client) CenCidrPublishWithContext(ctx context.Context, request *CenCidrPublishRequest) string {
	if request == nil {
		request = NewCenCidrPublishRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCenCidrPublishResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCenCidrDeleteRequest() (request *CenCidrDeleteRequest) {
	request = &CenCidrDeleteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cen", APIVersion, "CenCidrDelete")
	return
}

func NewCenCidrDeleteResponse() (response *CenCidrDeleteResponse) {
	response = &CenCidrDeleteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CenCidrDelete(request *CenCidrDeleteRequest) string {
	return c.CenCidrDeleteWithContext(context.Background(), request)
}

func (c *Client) CenCidrDeleteWithContext(ctx context.Context, request *CenCidrDeleteRequest) string {
	if request == nil {
		request = NewCenCidrDeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCenCidrDeleteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


