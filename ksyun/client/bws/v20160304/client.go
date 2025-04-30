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

func NewCreateBandWidthShareRequest() (request *CreateBandWidthShareRequest) {
	request = &CreateBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "CreateBandWidthShare")
	return
}

func NewCreateBandWidthShareResponse() (response *CreateBandWidthShareResponse) {
	response = &CreateBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateBandWidthShare(request *CreateBandWidthShareRequest) string {
	return c.CreateBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) CreateBandWidthShareWithContext(ctx context.Context, request *CreateBandWidthShareRequest) string {
	if request == nil {
		request = NewCreateBandWidthShareRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDescribeBandWidthSharesRequest() (request *DescribeBandWidthSharesRequest) {
	request = &DescribeBandWidthSharesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "DescribeBandWidthShares")
	return
}

func NewDescribeBandWidthSharesResponse() (response *DescribeBandWidthSharesResponse) {
	response = &DescribeBandWidthSharesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeBandWidthShares(request *DescribeBandWidthSharesRequest) string {
	return c.DescribeBandWidthSharesWithContext(context.Background(), request)
}

func (c *Client) DescribeBandWidthSharesWithContext(ctx context.Context, request *DescribeBandWidthSharesRequest) string {
	if request == nil {
		request = NewDescribeBandWidthSharesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeBandWidthSharesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewAssociateBandWidthShareRequest() (request *AssociateBandWidthShareRequest) {
	request = &AssociateBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "AssociateBandWidthShare")
	return
}

func NewAssociateBandWidthShareResponse() (response *AssociateBandWidthShareResponse) {
	response = &AssociateBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AssociateBandWidthShare(request *AssociateBandWidthShareRequest) string {
	return c.AssociateBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) AssociateBandWidthShareWithContext(ctx context.Context, request *AssociateBandWidthShareRequest) string {
	if request == nil {
		request = NewAssociateBandWidthShareRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAssociateBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDisassociateBandWidthShareRequest() (request *DisassociateBandWidthShareRequest) {
	request = &DisassociateBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "DisassociateBandWidthShare")
	return
}

func NewDisassociateBandWidthShareResponse() (response *DisassociateBandWidthShareResponse) {
	response = &DisassociateBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DisassociateBandWidthShare(request *DisassociateBandWidthShareRequest) string {
	return c.DisassociateBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) DisassociateBandWidthShareWithContext(ctx context.Context, request *DisassociateBandWidthShareRequest) string {
	if request == nil {
		request = NewDisassociateBandWidthShareRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDisassociateBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewModifyBandWidthShareRequest() (request *ModifyBandWidthShareRequest) {
	request = &ModifyBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "ModifyBandWidthShare")
	return
}

func NewModifyBandWidthShareResponse() (response *ModifyBandWidthShareResponse) {
	response = &ModifyBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyBandWidthShare(request *ModifyBandWidthShareRequest) string {
	return c.ModifyBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) ModifyBandWidthShareWithContext(ctx context.Context, request *ModifyBandWidthShareRequest) string {
	if request == nil {
		request = NewModifyBandWidthShareRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteBandWidthShareRequest() (request *DeleteBandWidthShareRequest) {
	request = &DeleteBandWidthShareRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "DeleteBandWidthShare")
	return
}

func NewDeleteBandWidthShareResponse() (response *DeleteBandWidthShareResponse) {
	response = &DeleteBandWidthShareResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteBandWidthShare(request *DeleteBandWidthShareRequest) string {
	return c.DeleteBandWidthShareWithContext(context.Background(), request)
}

func (c *Client) DeleteBandWidthShareWithContext(ctx context.Context, request *DeleteBandWidthShareRequest) string {
	if request == nil {
		request = NewDeleteBandWidthShareRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteBandWidthShareResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewQueryBwsTopEipMonitorRequest() (request *QueryBwsTopEipMonitorRequest) {
	request = &QueryBwsTopEipMonitorRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bws", APIVersion, "QueryBwsTopEipMonitor")
	return
}

func NewQueryBwsTopEipMonitorResponse() (response *QueryBwsTopEipMonitorResponse) {
	response = &QueryBwsTopEipMonitorResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryBwsTopEipMonitor(request *QueryBwsTopEipMonitorRequest) string {
	return c.QueryBwsTopEipMonitorWithContext(context.Background(), request)
}

func (c *Client) QueryBwsTopEipMonitorWithContext(ctx context.Context, request *QueryBwsTopEipMonitorRequest) string {
	if request == nil {
		request = NewQueryBwsTopEipMonitorRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryBwsTopEipMonitorResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


