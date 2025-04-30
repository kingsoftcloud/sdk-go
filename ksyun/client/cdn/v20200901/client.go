package v20200901
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2020-09-01"

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

func NewCreateUserUsageDataExportTaskRequest() (request *CreateUserUsageDataExportTaskRequest) {
	request = &CreateUserUsageDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "CreateUserUsageDataExportTask")
	return
}

func NewCreateUserUsageDataExportTaskResponse() (response *CreateUserUsageDataExportTaskResponse) {
	response = &CreateUserUsageDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateUserUsageDataExportTask(request *CreateUserUsageDataExportTaskRequest) string {
	return c.CreateUserUsageDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) CreateUserUsageDataExportTaskWithContext(ctx context.Context, request *CreateUserUsageDataExportTaskRequest) string {
	if request == nil {
		request = NewCreateUserUsageDataExportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateUserUsageDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetUserUsageDataExportTaskRequest() (request *GetUserUsageDataExportTaskRequest) {
	request = &GetUserUsageDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetUserUsageDataExportTask")
	return
}

func NewGetUserUsageDataExportTaskResponse() (response *GetUserUsageDataExportTaskResponse) {
	response = &GetUserUsageDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUserUsageDataExportTask(request *GetUserUsageDataExportTaskRequest) string {
	return c.GetUserUsageDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) GetUserUsageDataExportTaskWithContext(ctx context.Context, request *GetUserUsageDataExportTaskRequest) string {
	if request == nil {
		request = NewGetUserUsageDataExportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUserUsageDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteUserUsageDataExportTaskRequest() (request *DeleteUserUsageDataExportTaskRequest) {
	request = &DeleteUserUsageDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "DeleteUserUsageDataExportTask")
	return
}

func NewDeleteUserUsageDataExportTaskResponse() (response *DeleteUserUsageDataExportTaskResponse) {
	response = &DeleteUserUsageDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteUserUsageDataExportTask(request *DeleteUserUsageDataExportTaskRequest) string {
	return c.DeleteUserUsageDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteUserUsageDataExportTaskWithContext(ctx context.Context, request *DeleteUserUsageDataExportTaskRequest) string {
	if request == nil {
		request = NewDeleteUserUsageDataExportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteUserUsageDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetDomainUsageDataRequest() (request *GetDomainUsageDataRequest) {
	request = &GetDomainUsageDataRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetDomainUsageData")
	return
}

func NewGetDomainUsageDataResponse() (response *GetDomainUsageDataResponse) {
	response = &GetDomainUsageDataResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDomainUsageData(request *GetDomainUsageDataRequest) string {
	return c.GetDomainUsageDataWithContext(context.Background(), request)
}

func (c *Client) GetDomainUsageDataWithContext(ctx context.Context, request *GetDomainUsageDataRequest) string {
	if request == nil {
		request = NewGetDomainUsageDataRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDomainUsageDataResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateUsageDetailDataExportTaskRequest() (request *CreateUsageDetailDataExportTaskRequest) {
	request = &CreateUsageDetailDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "CreateUsageDetailDataExportTask")
	return
}

func NewCreateUsageDetailDataExportTaskResponse() (response *CreateUsageDetailDataExportTaskResponse) {
	response = &CreateUsageDetailDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateUsageDetailDataExportTask(request *CreateUsageDetailDataExportTaskRequest) string {
	return c.CreateUsageDetailDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) CreateUsageDetailDataExportTaskWithContext(ctx context.Context, request *CreateUsageDetailDataExportTaskRequest) string {
	if request == nil {
		request = NewCreateUsageDetailDataExportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateUsageDetailDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetUsageDetailDataExportTaskRequest() (request *GetUsageDetailDataExportTaskRequest) {
	request = &GetUsageDetailDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "GetUsageDetailDataExportTask")
	return
}

func NewGetUsageDetailDataExportTaskResponse() (response *GetUsageDetailDataExportTaskResponse) {
	response = &GetUsageDetailDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUsageDetailDataExportTask(request *GetUsageDetailDataExportTaskRequest) string {
	return c.GetUsageDetailDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) GetUsageDetailDataExportTaskWithContext(ctx context.Context, request *GetUsageDetailDataExportTaskRequest) string {
	if request == nil {
		request = NewGetUsageDetailDataExportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUsageDetailDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteUsageDetailDataExportTaskRequest() (request *DeleteUsageDetailDataExportTaskRequest) {
	request = &DeleteUsageDetailDataExportTaskRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdn", APIVersion, "DeleteUsageDetailDataExportTask")
	return
}

func NewDeleteUsageDetailDataExportTaskResponse() (response *DeleteUsageDetailDataExportTaskResponse) {
	response = &DeleteUsageDetailDataExportTaskResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteUsageDetailDataExportTask(request *DeleteUsageDetailDataExportTaskRequest) string {
	return c.DeleteUsageDetailDataExportTaskWithContext(context.Background(), request)
}

func (c *Client) DeleteUsageDetailDataExportTaskWithContext(ctx context.Context, request *DeleteUsageDetailDataExportTaskRequest) string {
	if request == nil {
		request = NewDeleteUsageDetailDataExportTaskRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteUsageDetailDataExportTaskResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}


