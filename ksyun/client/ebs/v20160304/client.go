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

func NewCreateVolumeRequest() (request *CreateVolumeRequest) {
	request = &CreateVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "CreateVolume")
	return
}

func NewCreateVolumeResponse() (response *CreateVolumeResponse) {
	response = &CreateVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVolume(request *CreateVolumeRequest) string {
	return c.CreateVolumeWithContext(context.Background(), request)
}

func (c *Client) CreateVolumeSend(request *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	statusCode, msg, err := c.CreateVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateVolumeWithContext(ctx context.Context, request *CreateVolumeRequest) string {
	if request == nil {
		request = NewCreateVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateVolumeWithContextV2(ctx context.Context, request *CreateVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewCreateVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewAttachVolumeRequest() (request *AttachVolumeRequest) {
	request = &AttachVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "AttachVolume")
	return
}

func NewAttachVolumeResponse() (response *AttachVolumeResponse) {
	response = &AttachVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) AttachVolume(request *AttachVolumeRequest) string {
	return c.AttachVolumeWithContext(context.Background(), request)
}

func (c *Client) AttachVolumeSend(request *AttachVolumeRequest) (*AttachVolumeResponse, error) {
	statusCode, msg, err := c.AttachVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct AttachVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) AttachVolumeWithContext(ctx context.Context, request *AttachVolumeRequest) string {
	if request == nil {
		request = NewAttachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) AttachVolumeWithContextV2(ctx context.Context, request *AttachVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewAttachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewAttachVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDetachVolumeRequest() (request *DetachVolumeRequest) {
	request = &DetachVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DetachVolume")
	return
}

func NewDetachVolumeResponse() (response *DetachVolumeResponse) {
	response = &DetachVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DetachVolume(request *DetachVolumeRequest) string {
	return c.DetachVolumeWithContext(context.Background(), request)
}

func (c *Client) DetachVolumeSend(request *DetachVolumeRequest) (*DetachVolumeResponse, error) {
	statusCode, msg, err := c.DetachVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DetachVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DetachVolumeWithContext(ctx context.Context, request *DetachVolumeRequest) string {
	if request == nil {
		request = NewDetachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DetachVolumeWithContextV2(ctx context.Context, request *DetachVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewDetachVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDetachVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteVolumeRequest() (request *DeleteVolumeRequest) {
	request = &DeleteVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DeleteVolume")
	return
}

func NewDeleteVolumeResponse() (response *DeleteVolumeResponse) {
	response = &DeleteVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVolume(request *DeleteVolumeRequest) string {
	return c.DeleteVolumeWithContext(context.Background(), request)
}

func (c *Client) DeleteVolumeSend(request *DeleteVolumeRequest) (*DeleteVolumeResponse, error) {
	statusCode, msg, err := c.DeleteVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteVolumeWithContext(ctx context.Context, request *DeleteVolumeRequest) string {
	if request == nil {
		request = NewDeleteVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteVolumeWithContextV2(ctx context.Context, request *DeleteVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResizeVolumeRequest() (request *ResizeVolumeRequest) {
	request = &ResizeVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ResizeVolume")
	return
}

func NewResizeVolumeResponse() (response *ResizeVolumeResponse) {
	response = &ResizeVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResizeVolume(request *ResizeVolumeRequest) string {
	return c.ResizeVolumeWithContext(context.Background(), request)
}

func (c *Client) ResizeVolumeSend(request *ResizeVolumeRequest) (*ResizeVolumeResponse, error) {
	statusCode, msg, err := c.ResizeVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResizeVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResizeVolumeWithContext(ctx context.Context, request *ResizeVolumeRequest) string {
	if request == nil {
		request = NewResizeVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResizeVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResizeVolumeWithContextV2(ctx context.Context, request *ResizeVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewResizeVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResizeVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeVolumesRequest() (request *DescribeVolumesRequest) {
	request = &DescribeVolumesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeVolumes")
	return
}

func NewDescribeVolumesResponse() (response *DescribeVolumesResponse) {
	response = &DescribeVolumesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVolumes(request *DescribeVolumesRequest) string {
	return c.DescribeVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeVolumesSend(request *DescribeVolumesRequest) (*DescribeVolumesResponse, error) {
	statusCode, msg, err := c.DescribeVolumesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeVolumesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeVolumesWithContext(ctx context.Context, request *DescribeVolumesRequest) string {
	if request == nil {
		request = NewDescribeVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVolumesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeVolumesWithContextV2(ctx context.Context, request *DescribeVolumesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeVolumesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVolumeRequest() (request *ModifyVolumeRequest) {
	request = &ModifyVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifyVolume")
	return
}

func NewModifyVolumeResponse() (response *ModifyVolumeResponse) {
	response = &ModifyVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVolume(request *ModifyVolumeRequest) string {
	return c.ModifyVolumeWithContext(context.Background(), request)
}

func (c *Client) ModifyVolumeSend(request *ModifyVolumeRequest) (*ModifyVolumeResponse, error) {
	statusCode, msg, err := c.ModifyVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVolumeWithContext(ctx context.Context, request *ModifyVolumeRequest) string {
	if request == nil {
		request = NewModifyVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVolumeWithContextV2(ctx context.Context, request *ModifyVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeEbsInstancesRequest() (request *DescribeEbsInstancesRequest) {
	request = &DescribeEbsInstancesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeEbsInstances")
	return
}

func NewDescribeEbsInstancesResponse() (response *DescribeEbsInstancesResponse) {
	response = &DescribeEbsInstancesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeEbsInstances(request *DescribeEbsInstancesRequest) string {
	return c.DescribeEbsInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeEbsInstancesSend(request *DescribeEbsInstancesRequest) (*DescribeEbsInstancesResponse, error) {
	statusCode, msg, err := c.DescribeEbsInstancesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeEbsInstancesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeEbsInstancesWithContext(ctx context.Context, request *DescribeEbsInstancesRequest) string {
	if request == nil {
		request = NewDescribeEbsInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEbsInstancesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeEbsInstancesWithContextV2(ctx context.Context, request *DescribeEbsInstancesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeEbsInstancesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeEbsInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceVolumesRequest() (request *DescribeInstanceVolumesRequest) {
	request = &DescribeInstanceVolumesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeInstanceVolumes")
	return
}

func NewDescribeInstanceVolumesResponse() (response *DescribeInstanceVolumesResponse) {
	response = &DescribeInstanceVolumesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceVolumes(request *DescribeInstanceVolumesRequest) string {
	return c.DescribeInstanceVolumesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceVolumesSend(request *DescribeInstanceVolumesRequest) (*DescribeInstanceVolumesResponse, error) {
	statusCode, msg, err := c.DescribeInstanceVolumesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceVolumesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceVolumesWithContext(ctx context.Context, request *DescribeInstanceVolumesRequest) string {
	if request == nil {
		request = NewDescribeInstanceVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceVolumesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceVolumesWithContextV2(ctx context.Context, request *DescribeInstanceVolumesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceVolumesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeInstanceVolumesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenewVolumeRequest() (request *RenewVolumeRequest) {
	request = &RenewVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "RenewVolume")
	return
}

func NewRenewVolumeResponse() (response *RenewVolumeResponse) {
	response = &RenewVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenewVolume(request *RenewVolumeRequest) string {
	return c.RenewVolumeWithContext(context.Background(), request)
}

func (c *Client) RenewVolumeSend(request *RenewVolumeRequest) (*RenewVolumeResponse, error) {
	statusCode, msg, err := c.RenewVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RenewVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenewVolumeWithContext(ctx context.Context, request *RenewVolumeRequest) string {
	if request == nil {
		request = NewRenewVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenewVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenewVolumeWithContextV2(ctx context.Context, request *RenewVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewRenewVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenewVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewUpdateVolumeProjectRequest() (request *UpdateVolumeProjectRequest) {
	request = &UpdateVolumeProjectRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "UpdateVolumeProject")
	return
}

func NewUpdateVolumeProjectResponse() (response *UpdateVolumeProjectResponse) {
	response = &UpdateVolumeProjectResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateVolumeProject(request *UpdateVolumeProjectRequest) string {
	return c.UpdateVolumeProjectWithContext(context.Background(), request)
}

func (c *Client) UpdateVolumeProjectSend(request *UpdateVolumeProjectRequest) (*UpdateVolumeProjectResponse, error) {
	statusCode, msg, err := c.UpdateVolumeProjectWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateVolumeProjectResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UpdateVolumeProjectWithContext(ctx context.Context, request *UpdateVolumeProjectRequest) string {
	if request == nil {
		request = NewUpdateVolumeProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateVolumeProjectResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UpdateVolumeProjectWithContextV2(ctx context.Context, request *UpdateVolumeProjectRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateVolumeProjectRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUpdateVolumeProjectResponse()
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
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeSnapshots")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeSnapshotsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "CreateSnapshot")
	return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) string {
	return c.CreateSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateSnapshotSend(request *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	statusCode, msg, err := c.CreateSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest) string {
	if request == nil {
		request = NewCreateSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSnapshotWithContextV2(ctx context.Context, request *CreateSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSnapshotResponse()
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
	request.Init().WithApiInfo("ebs", APIVersion, "DeleteSnapshot")
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
	request.SetContentType("application/x-www-form-urlencoded")

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
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRollbackSnapshotRequest() (request *RollbackSnapshotRequest) {
	request = &RollbackSnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "RollbackSnapshot")
	return
}

func NewRollbackSnapshotResponse() (response *RollbackSnapshotResponse) {
	response = &RollbackSnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RollbackSnapshot(request *RollbackSnapshotRequest) string {
	return c.RollbackSnapshotWithContext(context.Background(), request)
}

func (c *Client) RollbackSnapshotSend(request *RollbackSnapshotRequest) (*RollbackSnapshotResponse, error) {
	statusCode, msg, err := c.RollbackSnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RollbackSnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RollbackSnapshotWithContext(ctx context.Context, request *RollbackSnapshotRequest) string {
	if request == nil {
		request = NewRollbackSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackSnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RollbackSnapshotWithContextV2(ctx context.Context, request *RollbackSnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewRollbackSnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRollbackSnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySnapshotRequest() (request *ModifySnapshotRequest) {
	request = &ModifySnapshotRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifySnapshot")
	return
}

func NewModifySnapshotResponse() (response *ModifySnapshotResponse) {
	response = &ModifySnapshotResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySnapshot(request *ModifySnapshotRequest) string {
	return c.ModifySnapshotWithContext(context.Background(), request)
}

func (c *Client) ModifySnapshotSend(request *ModifySnapshotRequest) (*ModifySnapshotResponse, error) {
	statusCode, msg, err := c.ModifySnapshotWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySnapshotResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySnapshotWithContext(ctx context.Context, request *ModifySnapshotRequest) string {
	if request == nil {
		request = NewModifySnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySnapshotResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySnapshotWithContextV2(ctx context.Context, request *ModifySnapshotRequest) (int, string, error) {
	if request == nil {
		request = NewModifySnapshotRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySnapshotResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRecoveryVolumeRequest() (request *RecoveryVolumeRequest) {
	request = &RecoveryVolumeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "RecoveryVolume")
	return
}

func NewRecoveryVolumeResponse() (response *RecoveryVolumeResponse) {
	response = &RecoveryVolumeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RecoveryVolume(request *RecoveryVolumeRequest) string {
	return c.RecoveryVolumeWithContext(context.Background(), request)
}

func (c *Client) RecoveryVolumeSend(request *RecoveryVolumeRequest) (*RecoveryVolumeResponse, error) {
	statusCode, msg, err := c.RecoveryVolumeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RecoveryVolumeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RecoveryVolumeWithContext(ctx context.Context, request *RecoveryVolumeRequest) string {
	if request == nil {
		request = NewRecoveryVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRecoveryVolumeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RecoveryVolumeWithContextV2(ctx context.Context, request *RecoveryVolumeRequest) (int, string, error) {
	if request == nil {
		request = NewRecoveryVolumeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRecoveryVolumeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewValidateAttachInstanceRequest() (request *ValidateAttachInstanceRequest) {
	request = &ValidateAttachInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ValidateAttachInstance")
	return
}

func NewValidateAttachInstanceResponse() (response *ValidateAttachInstanceResponse) {
	response = &ValidateAttachInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ValidateAttachInstance(request *ValidateAttachInstanceRequest) string {
	return c.ValidateAttachInstanceWithContext(context.Background(), request)
}

func (c *Client) ValidateAttachInstanceSend(request *ValidateAttachInstanceRequest) (*ValidateAttachInstanceResponse, error) {
	statusCode, msg, err := c.ValidateAttachInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ValidateAttachInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ValidateAttachInstanceWithContext(ctx context.Context, request *ValidateAttachInstanceRequest) string {
	if request == nil {
		request = NewValidateAttachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewValidateAttachInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ValidateAttachInstanceWithContextV2(ctx context.Context, request *ValidateAttachInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewValidateAttachInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewValidateAttachInstanceResponse()
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
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeAvailabilityZones")
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
func NewDescribeCreateVolumePriceRequest() (request *DescribeCreateVolumePriceRequest) {
	request = &DescribeCreateVolumePriceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeCreateVolumePrice")
	return
}

func NewDescribeCreateVolumePriceResponse() (response *DescribeCreateVolumePriceResponse) {
	response = &DescribeCreateVolumePriceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCreateVolumePrice(request *DescribeCreateVolumePriceRequest) string {
	return c.DescribeCreateVolumePriceWithContext(context.Background(), request)
}

func (c *Client) DescribeCreateVolumePriceSend(request *DescribeCreateVolumePriceRequest) (*DescribeCreateVolumePriceResponse, error) {
	statusCode, msg, err := c.DescribeCreateVolumePriceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeCreateVolumePriceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeCreateVolumePriceWithContext(ctx context.Context, request *DescribeCreateVolumePriceRequest) string {
	if request == nil {
		request = NewDescribeCreateVolumePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCreateVolumePriceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeCreateVolumePriceWithContextV2(ctx context.Context, request *DescribeCreateVolumePriceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeCreateVolumePriceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeCreateVolumePriceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifySnapshotTypeRequest() (request *ModifySnapshotTypeRequest) {
	request = &ModifySnapshotTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifySnapshotType")
	return
}

func NewModifySnapshotTypeResponse() (response *ModifySnapshotTypeResponse) {
	response = &ModifySnapshotTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySnapshotType(request *ModifySnapshotTypeRequest) string {
	return c.ModifySnapshotTypeWithContext(context.Background(), request)
}

func (c *Client) ModifySnapshotTypeSend(request *ModifySnapshotTypeRequest) (*ModifySnapshotTypeResponse, error) {
	statusCode, msg, err := c.ModifySnapshotTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifySnapshotTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifySnapshotTypeWithContext(ctx context.Context, request *ModifySnapshotTypeRequest) string {
	if request == nil {
		request = NewModifySnapshotTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySnapshotTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifySnapshotTypeWithContextV2(ctx context.Context, request *ModifySnapshotTypeRequest) (int, string, error) {
	if request == nil {
		request = NewModifySnapshotTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifySnapshotTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVolumeTypeRequest() (request *ModifyVolumeTypeRequest) {
	request = &ModifyVolumeTypeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifyVolumeType")
	return
}

func NewModifyVolumeTypeResponse() (response *ModifyVolumeTypeResponse) {
	response = &ModifyVolumeTypeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVolumeType(request *ModifyVolumeTypeRequest) string {
	return c.ModifyVolumeTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyVolumeTypeSend(request *ModifyVolumeTypeRequest) (*ModifyVolumeTypeResponse, error) {
	statusCode, msg, err := c.ModifyVolumeTypeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyVolumeTypeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVolumeTypeWithContext(ctx context.Context, request *ModifyVolumeTypeRequest) string {
	if request == nil {
		request = NewModifyVolumeTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVolumeTypeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVolumeTypeWithContextV2(ctx context.Context, request *ModifyVolumeTypeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVolumeTypeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVolumeTypeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyDedicatedBlockStorageClusterAttributeRequest() (request *ModifyDedicatedBlockStorageClusterAttributeRequest) {
	request = &ModifyDedicatedBlockStorageClusterAttributeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifyDedicatedBlockStorageClusterAttribute")
	return
}

func NewModifyDedicatedBlockStorageClusterAttributeResponse() (response *ModifyDedicatedBlockStorageClusterAttributeResponse) {
	response = &ModifyDedicatedBlockStorageClusterAttributeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDedicatedBlockStorageClusterAttribute(request *ModifyDedicatedBlockStorageClusterAttributeRequest) string {
	return c.ModifyDedicatedBlockStorageClusterAttributeWithContext(context.Background(), request)
}

func (c *Client) ModifyDedicatedBlockStorageClusterAttributeSend(request *ModifyDedicatedBlockStorageClusterAttributeRequest) (*ModifyDedicatedBlockStorageClusterAttributeResponse, error) {
	statusCode, msg, err := c.ModifyDedicatedBlockStorageClusterAttributeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyDedicatedBlockStorageClusterAttributeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyDedicatedBlockStorageClusterAttributeWithContext(ctx context.Context, request *ModifyDedicatedBlockStorageClusterAttributeRequest) string {
	if request == nil {
		request = NewModifyDedicatedBlockStorageClusterAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDedicatedBlockStorageClusterAttributeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyDedicatedBlockStorageClusterAttributeWithContextV2(ctx context.Context, request *ModifyDedicatedBlockStorageClusterAttributeRequest) (int, string, error) {
	if request == nil {
		request = NewModifyDedicatedBlockStorageClusterAttributeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyDedicatedBlockStorageClusterAttributeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResizeDedicatedBlockStorageClustersRequest() (request *ResizeDedicatedBlockStorageClustersRequest) {
	request = &ResizeDedicatedBlockStorageClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ResizeDedicatedBlockStorageClusters")
	return
}

func NewResizeDedicatedBlockStorageClustersResponse() (response *ResizeDedicatedBlockStorageClustersResponse) {
	response = &ResizeDedicatedBlockStorageClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResizeDedicatedBlockStorageClusters(request *ResizeDedicatedBlockStorageClustersRequest) string {
	return c.ResizeDedicatedBlockStorageClustersWithContext(context.Background(), request)
}

func (c *Client) ResizeDedicatedBlockStorageClustersSend(request *ResizeDedicatedBlockStorageClustersRequest) (*ResizeDedicatedBlockStorageClustersResponse, error) {
	statusCode, msg, err := c.ResizeDedicatedBlockStorageClustersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResizeDedicatedBlockStorageClustersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResizeDedicatedBlockStorageClustersWithContext(ctx context.Context, request *ResizeDedicatedBlockStorageClustersRequest) string {
	if request == nil {
		request = NewResizeDedicatedBlockStorageClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResizeDedicatedBlockStorageClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResizeDedicatedBlockStorageClustersWithContextV2(ctx context.Context, request *ResizeDedicatedBlockStorageClustersRequest) (int, string, error) {
	if request == nil {
		request = NewResizeDedicatedBlockStorageClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResizeDedicatedBlockStorageClustersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDedicatedBlockStorageClustersRequest() (request *DescribeDedicatedBlockStorageClustersRequest) {
	request = &DescribeDedicatedBlockStorageClustersRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "DescribeDedicatedBlockStorageClusters")
	return
}

func NewDescribeDedicatedBlockStorageClustersResponse() (response *DescribeDedicatedBlockStorageClustersResponse) {
	response = &DescribeDedicatedBlockStorageClustersResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDedicatedBlockStorageClusters(request *DescribeDedicatedBlockStorageClustersRequest) string {
	return c.DescribeDedicatedBlockStorageClustersWithContext(context.Background(), request)
}

func (c *Client) DescribeDedicatedBlockStorageClustersSend(request *DescribeDedicatedBlockStorageClustersRequest) (*DescribeDedicatedBlockStorageClustersResponse, error) {
	statusCode, msg, err := c.DescribeDedicatedBlockStorageClustersWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDedicatedBlockStorageClustersResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDedicatedBlockStorageClustersWithContext(ctx context.Context, request *DescribeDedicatedBlockStorageClustersRequest) string {
	if request == nil {
		request = NewDescribeDedicatedBlockStorageClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedBlockStorageClustersResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDedicatedBlockStorageClustersWithContextV2(ctx context.Context, request *DescribeDedicatedBlockStorageClustersRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDedicatedBlockStorageClustersRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDescribeDedicatedBlockStorageClustersResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDedicatedBlockStorageClusterRequest() (request *CreateDedicatedBlockStorageClusterRequest) {
	request = &CreateDedicatedBlockStorageClusterRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "CreateDedicatedBlockStorageCluster")
	return
}

func NewCreateDedicatedBlockStorageClusterResponse() (response *CreateDedicatedBlockStorageClusterResponse) {
	response = &CreateDedicatedBlockStorageClusterResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDedicatedBlockStorageCluster(request *CreateDedicatedBlockStorageClusterRequest) string {
	return c.CreateDedicatedBlockStorageClusterWithContext(context.Background(), request)
}

func (c *Client) CreateDedicatedBlockStorageClusterSend(request *CreateDedicatedBlockStorageClusterRequest) (*CreateDedicatedBlockStorageClusterResponse, error) {
	statusCode, msg, err := c.CreateDedicatedBlockStorageClusterWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDedicatedBlockStorageClusterResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDedicatedBlockStorageClusterWithContext(ctx context.Context, request *CreateDedicatedBlockStorageClusterRequest) string {
	if request == nil {
		request = NewCreateDedicatedBlockStorageClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedBlockStorageClusterResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDedicatedBlockStorageClusterWithContextV2(ctx context.Context, request *CreateDedicatedBlockStorageClusterRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDedicatedBlockStorageClusterRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDedicatedBlockStorageClusterResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyVolumePresetRequest() (request *ModifyVolumePresetRequest) {
	request = &ModifyVolumePresetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "ModifyVolumePreset")
	return
}

func NewModifyVolumePresetResponse() (response *ModifyVolumePresetResponse) {
	response = &ModifyVolumePresetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVolumePreset(request *ModifyVolumePresetRequest) string {
	return c.ModifyVolumePresetWithContext(context.Background(), request)
}

func (c *Client) ModifyVolumePresetSend(request *ModifyVolumePresetRequest) (*ModifyVolumePresetResponse, error) {
	statusCode, msg, err := c.ModifyVolumePresetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyVolumePresetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyVolumePresetWithContext(ctx context.Context, request *ModifyVolumePresetRequest) string {
	if request == nil {
		request = NewModifyVolumePresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVolumePresetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyVolumePresetWithContextV2(ctx context.Context, request *ModifyVolumePresetRequest) (int, string, error) {
	if request == nil {
		request = NewModifyVolumePresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyVolumePresetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGetUpgradeVolumeTypeProcessInfoRequest() (request *GetUpgradeVolumeTypeProcessInfoRequest) {
	request = &GetUpgradeVolumeTypeProcessInfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ebs", APIVersion, "GetUpgradeVolumeTypeProcessInfo")
	return
}

func NewGetUpgradeVolumeTypeProcessInfoResponse() (response *GetUpgradeVolumeTypeProcessInfoResponse) {
	response = &GetUpgradeVolumeTypeProcessInfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetUpgradeVolumeTypeProcessInfo(request *GetUpgradeVolumeTypeProcessInfoRequest) string {
	return c.GetUpgradeVolumeTypeProcessInfoWithContext(context.Background(), request)
}

func (c *Client) GetUpgradeVolumeTypeProcessInfoSend(request *GetUpgradeVolumeTypeProcessInfoRequest) (*GetUpgradeVolumeTypeProcessInfoResponse, error) {
	statusCode, msg, err := c.GetUpgradeVolumeTypeProcessInfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetUpgradeVolumeTypeProcessInfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetUpgradeVolumeTypeProcessInfoWithContext(ctx context.Context, request *GetUpgradeVolumeTypeProcessInfoRequest) string {
	if request == nil {
		request = NewGetUpgradeVolumeTypeProcessInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUpgradeVolumeTypeProcessInfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetUpgradeVolumeTypeProcessInfoWithContextV2(ctx context.Context, request *GetUpgradeVolumeTypeProcessInfoRequest) (int, string, error) {
	if request == nil {
		request = NewGetUpgradeVolumeTypeProcessInfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetUpgradeVolumeTypeProcessInfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


