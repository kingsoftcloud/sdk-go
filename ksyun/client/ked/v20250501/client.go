package v20250501

import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "V1"

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

func NewCloudDeskreinstallRequest() (request *CloudDeskreinstallRequest) {
	request = &CloudDeskreinstallRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CloudDeskreinstall")
	return
}

func NewCloudDeskreinstallResponse() (response *CloudDeskreinstallResponse) {
	response = &CloudDeskreinstallResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloudDeskreinstall(request *CloudDeskreinstallRequest) string {
	return c.CloudDeskreinstallWithContext(context.Background(), request)
}

func (c *Client) CloudDeskreinstallWithContext(ctx context.Context, request *CloudDeskreinstallRequest) string {
	if request == nil {
		request = NewCloudDeskreinstallRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskreinstallResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCloudDeskmanageRequest() (request *CloudDeskmanageRequest) {
	request = &CloudDeskmanageRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CloudDeskmanage")
	return
}

func NewCloudDeskmanageResponse() (response *CloudDeskmanageResponse) {
	response = &CloudDeskmanageResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloudDeskmanage(request *CloudDeskmanageRequest) string {
	return c.CloudDeskmanageWithContext(context.Background(), request)
}

func (c *Client) CloudDeskmanageWithContext(ctx context.Context, request *CloudDeskmanageRequest) string {
	if request == nil {
		request = NewCloudDeskmanageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskmanageResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCloudDeskeditRequest() (request *CloudDeskeditRequest) {
	request = &CloudDeskeditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CloudDeskedit")
	return
}

func NewCloudDeskeditResponse() (response *CloudDeskeditResponse) {
	response = &CloudDeskeditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloudDeskedit(request *CloudDeskeditRequest) string {
	return c.CloudDeskeditWithContext(context.Background(), request)
}

func (c *Client) CloudDeskeditWithContext(ctx context.Context, request *CloudDeskeditRequest) string {
	if request == nil {
		request = NewCloudDeskeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskeditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCloudDeskcreateRequest() (request *CloudDeskcreateRequest) {
	request = &CloudDeskcreateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CloudDeskcreate")
	return
}

func NewCloudDeskcreateResponse() (response *CloudDeskcreateResponse) {
	response = &CloudDeskcreateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloudDeskcreate(request *CloudDeskcreateRequest) string {
	return c.CloudDeskcreateWithContext(context.Background(), request)
}

func (c *Client) CloudDeskcreateWithContext(ctx context.Context, request *CloudDeskcreateRequest) string {
	if request == nil {
		request = NewCloudDeskcreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskcreateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCloudDesklistRequest() (request *CloudDesklistRequest) {
	request = &CloudDesklistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CloudDesklist")
	return
}

func NewCloudDesklistResponse() (response *CloudDesklistResponse) {
	response = &CloudDesklistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloudDesklist(request *CloudDesklistRequest) string {
	return c.CloudDesklistWithContext(context.Background(), request)
}

func (c *Client) CloudDesklistWithContext(ctx context.Context, request *CloudDesklistRequest) string {
	if request == nil {
		request = NewCloudDesklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCloudDesklistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategyruleeditRequest() (request *StrategyruleeditRequest) {
	request = &StrategyruleeditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategyruleedit")
	return
}

func NewStrategyruleeditResponse() (response *StrategyruleeditResponse) {
	response = &StrategyruleeditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategyruleedit(request *StrategyruleeditRequest) string {
	return c.StrategyruleeditWithContext(context.Background(), request)
}

func (c *Client) StrategyruleeditWithContext(ctx context.Context, request *StrategyruleeditRequest) string {
	if request == nil {
		request = NewStrategyruleeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyruleeditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategyrulecreateRequest() (request *StrategyrulecreateRequest) {
	request = &StrategyrulecreateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategyrulecreate")
	return
}

func NewStrategyrulecreateResponse() (response *StrategyrulecreateResponse) {
	response = &StrategyrulecreateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategyrulecreate(request *StrategyrulecreateRequest) string {
	return c.StrategyrulecreateWithContext(context.Background(), request)
}

func (c *Client) StrategyrulecreateWithContext(ctx context.Context, request *StrategyrulecreateRequest) string {
	if request == nil {
		request = NewStrategyrulecreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyrulecreateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategyunboundRequest() (request *StrategyunboundRequest) {
	request = &StrategyunboundRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategyunbound")
	return
}

func NewStrategyunboundResponse() (response *StrategyunboundResponse) {
	response = &StrategyunboundResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategyunbound(request *StrategyunboundRequest) string {
	return c.StrategyunboundWithContext(context.Background(), request)
}

func (c *Client) StrategyunboundWithContext(ctx context.Context, request *StrategyunboundRequest) string {
	if request == nil {
		request = NewStrategyunboundRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyunboundResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategyboundRequest() (request *StrategyboundRequest) {
	request = &StrategyboundRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategybound")
	return
}

func NewStrategyboundResponse() (response *StrategyboundResponse) {
	response = &StrategyboundResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategybound(request *StrategyboundRequest) string {
	return c.StrategyboundWithContext(context.Background(), request)
}

func (c *Client) StrategyboundWithContext(ctx context.Context, request *StrategyboundRequest) string {
	if request == nil {
		request = NewStrategyboundRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyboundResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategydeleteRequest() (request *StrategydeleteRequest) {
	request = &StrategydeleteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategydelete")
	return
}

func NewStrategydeleteResponse() (response *StrategydeleteResponse) {
	response = &StrategydeleteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategydelete(request *StrategydeleteRequest) string {
	return c.StrategydeleteWithContext(context.Background(), request)
}

func (c *Client) StrategydeleteWithContext(ctx context.Context, request *StrategydeleteRequest) string {
	if request == nil {
		request = NewStrategydeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategydeleteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategyeditRequest() (request *StrategyeditRequest) {
	request = &StrategyeditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategyedit")
	return
}

func NewStrategyeditResponse() (response *StrategyeditResponse) {
	response = &StrategyeditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategyedit(request *StrategyeditRequest) string {
	return c.StrategyeditWithContext(context.Background(), request)
}

func (c *Client) StrategyeditWithContext(ctx context.Context, request *StrategyeditRequest) string {
	if request == nil {
		request = NewStrategyeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyeditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategycreateRequest() (request *StrategycreateRequest) {
	request = &StrategycreateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategycreate")
	return
}

func NewStrategycreateResponse() (response *StrategycreateResponse) {
	response = &StrategycreateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategycreate(request *StrategycreateRequest) string {
	return c.StrategycreateWithContext(context.Background(), request)
}

func (c *Client) StrategycreateWithContext(ctx context.Context, request *StrategycreateRequest) string {
	if request == nil {
		request = NewStrategycreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategycreateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategylistRequest() (request *StrategylistRequest) {
	request = &StrategylistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Strategylist")
	return
}

func NewStrategylistResponse() (response *StrategylistResponse) {
	response = &StrategylistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Strategylist(request *StrategylistRequest) string {
	return c.StrategylistWithContext(context.Background(), request)
}

func (c *Client) StrategylistWithContext(ctx context.Context, request *StrategylistRequest) string {
	if request == nil {
		request = NewStrategylistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewStrategylistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRolesdeleteRequest() (request *RolesdeleteRequest) {
	request = &RolesdeleteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Rolesdelete")
	return
}

func NewRolesdeleteResponse() (response *RolesdeleteResponse) {
	response = &RolesdeleteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Rolesdelete(request *RolesdeleteRequest) string {
	return c.RolesdeleteWithContext(context.Background(), request)
}

func (c *Client) RolesdeleteWithContext(ctx context.Context, request *RolesdeleteRequest) string {
	if request == nil {
		request = NewRolesdeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRolesdeleteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRoleseditRequest() (request *RoleseditRequest) {
	request = &RoleseditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Rolesedit")
	return
}

func NewRoleseditResponse() (response *RoleseditResponse) {
	response = &RoleseditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Rolesedit(request *RoleseditRequest) string {
	return c.RoleseditWithContext(context.Background(), request)
}

func (c *Client) RoleseditWithContext(ctx context.Context, request *RoleseditRequest) string {
	if request == nil {
		request = NewRoleseditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRoleseditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRolescreateRequest() (request *RolescreateRequest) {
	request = &RolescreateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Rolescreate")
	return
}

func NewRolescreateResponse() (response *RolescreateResponse) {
	response = &RolescreateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Rolescreate(request *RolescreateRequest) string {
	return c.RolescreateWithContext(context.Background(), request)
}

func (c *Client) RolescreateWithContext(ctx context.Context, request *RolescreateRequest) string {
	if request == nil {
		request = NewRolescreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRolescreateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewRoleslistRequest() (request *RoleslistRequest) {
	request = &RoleslistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Roleslist")
	return
}

func NewRoleslistResponse() (response *RoleslistResponse) {
	response = &RoleslistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Roleslist(request *RoleslistRequest) string {
	return c.RoleslistWithContext(context.Background(), request)
}

func (c *Client) RoleslistWithContext(ctx context.Context, request *RoleslistRequest) string {
	if request == nil {
		request = NewRoleslistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRoleslistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImagedeleteRequest() (request *ImagedeleteRequest) {
	request = &ImagedeleteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Imagedelete")
	return
}

func NewImagedeleteResponse() (response *ImagedeleteResponse) {
	response = &ImagedeleteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Imagedelete(request *ImagedeleteRequest) string {
	return c.ImagedeleteWithContext(context.Background(), request)
}

func (c *Client) ImagedeleteWithContext(ctx context.Context, request *ImagedeleteRequest) string {
	if request == nil {
		request = NewImagedeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImagedeleteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImageeditRequest() (request *ImageeditRequest) {
	request = &ImageeditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Imageedit")
	return
}

func NewImageeditResponse() (response *ImageeditResponse) {
	response = &ImageeditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Imageedit(request *ImageeditRequest) string {
	return c.ImageeditWithContext(context.Background(), request)
}

func (c *Client) ImageeditWithContext(ctx context.Context, request *ImageeditRequest) string {
	if request == nil {
		request = NewImageeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImageeditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImagecreateRequest() (request *ImagecreateRequest) {
	request = &ImagecreateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Imagecreate")
	return
}

func NewImagecreateResponse() (response *ImagecreateResponse) {
	response = &ImagecreateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Imagecreate(request *ImagecreateRequest) string {
	return c.ImagecreateWithContext(context.Background(), request)
}

func (c *Client) ImagecreateWithContext(ctx context.Context, request *ImagecreateRequest) string {
	if request == nil {
		request = NewImagecreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImagecreateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewImagelistRequest() (request *ImagelistRequest) {
	request = &ImagelistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Imagelist")
	return
}

func NewImagelistResponse() (response *ImagelistResponse) {
	response = &ImagelistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Imagelist(request *ImagelistRequest) string {
	return c.ImagelistWithContext(context.Background(), request)
}

func (c *Client) ImagelistWithContext(ctx context.Context, request *ImagelistRequest) string {
	if request == nil {
		request = NewImagelistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewImagelistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewStrategyrulebatchEditRequest() (request *StrategyrulebatchEditRequest) {
	request = &StrategyrulebatchEditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "StrategyrulebatchEdit")
	return
}

func NewStrategyrulebatchEditResponse() (response *StrategyrulebatchEditResponse) {
	response = &StrategyrulebatchEditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) StrategyrulebatchEdit(request *StrategyrulebatchEditRequest) string {
	return c.StrategyrulebatchEditWithContext(context.Background(), request)
}

func (c *Client) StrategyrulebatchEditWithContext(ctx context.Context, request *StrategyrulebatchEditRequest) string {
	if request == nil {
		request = NewStrategyrulebatchEditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyrulebatchEditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewMonitorregionsRequest() (request *MonitorregionsRequest) {
	request = &MonitorregionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Monitorregions")
	return
}

func NewMonitorregionsResponse() (response *MonitorregionsResponse) {
	response = &MonitorregionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Monitorregions(request *MonitorregionsRequest) string {
	return c.MonitorregionsWithContext(context.Background(), request)
}

func (c *Client) MonitorregionsWithContext(ctx context.Context, request *MonitorregionsRequest) string {
	if request == nil {
		request = NewMonitorregionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewMonitorregionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUsersinstancebindRequest() (request *UsersinstancebindRequest) {
	request = &UsersinstancebindRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Usersinstancebind")
	return
}

func NewUsersinstancebindResponse() (response *UsersinstancebindResponse) {
	response = &UsersinstancebindResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Usersinstancebind(request *UsersinstancebindRequest) string {
	return c.UsersinstancebindWithContext(context.Background(), request)
}

func (c *Client) UsersinstancebindWithContext(ctx context.Context, request *UsersinstancebindRequest) string {
	if request == nil {
		request = NewUsersinstancebindRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUsersinstancebindResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUserspasswordresetRequest() (request *UserspasswordresetRequest) {
	request = &UserspasswordresetRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Userspasswordreset")
	return
}

func NewUserspasswordresetResponse() (response *UserspasswordresetResponse) {
	response = &UserspasswordresetResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Userspasswordreset(request *UserspasswordresetRequest) string {
	return c.UserspasswordresetWithContext(context.Background(), request)
}

func (c *Client) UserspasswordresetWithContext(ctx context.Context, request *UserspasswordresetRequest) string {
	if request == nil {
		request = NewUserspasswordresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserspasswordresetResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUsersdeleteRequest() (request *UsersdeleteRequest) {
	request = &UsersdeleteRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Usersdelete")
	return
}

func NewUsersdeleteResponse() (response *UsersdeleteResponse) {
	response = &UsersdeleteResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Usersdelete(request *UsersdeleteRequest) string {
	return c.UsersdeleteWithContext(context.Background(), request)
}

func (c *Client) UsersdeleteWithContext(ctx context.Context, request *UsersdeleteRequest) string {
	if request == nil {
		request = NewUsersdeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUsersdeleteResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUserseditRequest() (request *UserseditRequest) {
	request = &UserseditRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Usersedit")
	return
}

func NewUserseditResponse() (response *UserseditResponse) {
	response = &UserseditResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Usersedit(request *UserseditRequest) string {
	return c.UserseditWithContext(context.Background(), request)
}

func (c *Client) UserseditWithContext(ctx context.Context, request *UserseditRequest) string {
	if request == nil {
		request = NewUserseditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserseditResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUserscreateRequest() (request *UserscreateRequest) {
	request = &UserscreateRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Userscreate")
	return
}

func NewUserscreateResponse() (response *UserscreateResponse) {
	response = &UserscreateResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Userscreate(request *UserscreateRequest) string {
	return c.UserscreateWithContext(context.Background(), request)
}

func (c *Client) UserscreateWithContext(ctx context.Context, request *UserscreateRequest) string {
	if request == nil {
		request = NewUserscreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserscreateResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUserslistRequest() (request *UserslistRequest) {
	request = &UserslistRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "Userslist")
	return
}

func NewUserslistResponse() (response *UserslistResponse) {
	response = &UserslistResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) Userslist(request *UserslistRequest) string {
	return c.UserslistWithContext(context.Background(), request)
}

func (c *Client) UserslistWithContext(ctx context.Context, request *UserslistRequest) string {
	if request == nil {
		request = NewUserslistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewUserslistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCloudDeskgetDesktopUrlRequest() (request *CloudDeskgetDesktopUrlRequest) {
	request = &CloudDeskgetDesktopUrlRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CloudDeskgetDesktopUrl")
	return
}

func NewCloudDeskgetDesktopUrlResponse() (response *CloudDeskgetDesktopUrlResponse) {
	response = &CloudDeskgetDesktopUrlResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CloudDeskgetDesktopUrl(request *CloudDeskgetDesktopUrlRequest) string {
	return c.CloudDeskgetDesktopUrlWithContext(context.Background(), request)
}

func (c *Client) CloudDeskgetDesktopUrlWithContext(ctx context.Context, request *CloudDeskgetDesktopUrlRequest) string {
	if request == nil {
		request = NewCloudDeskgetDesktopUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskgetDesktopUrlResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewQueryCloudDesksubmitShellRequest() (request *QueryCloudDesksubmitShellRequest) {
	request = &QueryCloudDesksubmitShellRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "QueryCloudDesksubmitShell")
	return
}

func NewQueryCloudDesksubmitShellResponse() (response *QueryCloudDesksubmitShellResponse) {
	response = &QueryCloudDesksubmitShellResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryCloudDesksubmitShell(request *QueryCloudDesksubmitShellRequest) string {
	return c.QueryCloudDesksubmitShellWithContext(context.Background(), request)
}

func (c *Client) QueryCloudDesksubmitShellWithContext(ctx context.Context, request *QueryCloudDesksubmitShellRequest) string {
	if request == nil {
		request = NewQueryCloudDesksubmitShellRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryCloudDesksubmitShellResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateCloudDeskgetTokenRequest() (request *CreateCloudDeskgetTokenRequest) {
	request = &CreateCloudDeskgetTokenRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CreateCloudDeskgetToken")
	return
}

func NewCreateCloudDeskgetTokenResponse() (response *CreateCloudDeskgetTokenResponse) {
	response = &CreateCloudDeskgetTokenResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateCloudDeskgetToken(request *CreateCloudDeskgetTokenRequest) string {
	return c.CreateCloudDeskgetTokenWithContext(context.Background(), request)
}

func (c *Client) CreateCloudDeskgetTokenWithContext(ctx context.Context, request *CreateCloudDeskgetTokenRequest) string {
	if request == nil {
		request = NewCreateCloudDeskgetTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCloudDeskgetTokenResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewQueryShellStatusRequest() (request *QueryShellStatusRequest) {
	request = &QueryShellStatusRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "QueryShellStatus")
	return
}

func NewQueryShellStatusResponse() (response *QueryShellStatusResponse) {
	response = &QueryShellStatusResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryShellStatus(request *QueryShellStatusRequest) string {
	return c.QueryShellStatusWithContext(context.Background(), request)
}

func (c *Client) QueryShellStatusWithContext(ctx context.Context, request *QueryShellStatusRequest) string {
	if request == nil {
		request = NewQueryShellStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryShellStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewSetProxyIpRequest() (request *SetProxyIpRequest) {
	request = &SetProxyIpRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "SetProxyIp")
	return
}

func NewSetProxyIpResponse() (response *SetProxyIpResponse) {
	response = &SetProxyIpResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) SetProxyIp(request *SetProxyIpRequest) string {
	return c.SetProxyIpWithContext(context.Background(), request)
}

func (c *Client) SetProxyIpWithContext(ctx context.Context, request *SetProxyIpRequest) string {
	if request == nil {
		request = NewSetProxyIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetProxyIpResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetProxyConfigRequest() (request *GetProxyConfigRequest) {
	request = &GetProxyConfigRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "GetProxyConfig")
	return
}

func NewGetProxyConfigResponse() (response *GetProxyConfigResponse) {
	response = &GetProxyConfigResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetProxyConfig(request *GetProxyConfigRequest) string {
	return c.GetProxyConfigWithContext(context.Background(), request)
}

func (c *Client) GetProxyConfigWithContext(ctx context.Context, request *GetProxyConfigRequest) string {
	if request == nil {
		request = NewGetProxyConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetProxyConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewQueryRuledetailRequest() (request *QueryRuledetailRequest) {
	request = &QueryRuledetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "QueryRuledetail")
	return
}

func NewQueryRuledetailResponse() (response *QueryRuledetailResponse) {
	response = &QueryRuledetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryRuledetail(request *QueryRuledetailRequest) string {
	return c.QueryRuledetailWithContext(context.Background(), request)
}

func (c *Client) QueryRuledetailWithContext(ctx context.Context, request *QueryRuledetailRequest) string {
	if request == nil {
		request = NewQueryRuledetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewQueryRuledetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewQueryUsersinfoRequest() (request *QueryUsersinfoRequest) {
	request = &QueryUsersinfoRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "QueryUsersinfo")
	return
}

func NewQueryUsersinfoResponse() (response *QueryUsersinfoResponse) {
	response = &QueryUsersinfoResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) QueryUsersinfo(request *QueryUsersinfoRequest) string {
	return c.QueryUsersinfoWithContext(context.Background(), request)
}

func (c *Client) QueryUsersinfoWithContext(ctx context.Context, request *QueryUsersinfoRequest) string {
	if request == nil {
		request = NewQueryUsersinfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryUsersinfoResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewGetDetailRequest() (request *GetDetailRequest) {
	request = &GetDetailRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "GetDetail")
	return
}

func NewGetDetailResponse() (response *GetDetailResponse) {
	response = &GetDetailResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GetDetail(request *GetDetailRequest) string {
	return c.GetDetailWithContext(context.Background(), request)
}

func (c *Client) GetDetailWithContext(ctx context.Context, request *GetDetailRequest) string {
	if request == nil {
		request = NewGetDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGetDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewListLabelRequest() (request *ListLabelRequest) {
	request = &ListLabelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "ListLabel")
	return
}

func NewListLabelResponse() (response *ListLabelResponse) {
	response = &ListLabelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ListLabel(request *ListLabelRequest) string {
	return c.ListLabelWithContext(context.Background(), request)
}

func (c *Client) ListLabelWithContext(ctx context.Context, request *ListLabelRequest) string {
	if request == nil {
		request = NewListLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewListLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCancelInstanceLabelRequest() (request *CancelInstanceLabelRequest) {
	request = &CancelInstanceLabelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CancelInstanceLabel")
	return
}

func NewCancelInstanceLabelResponse() (response *CancelInstanceLabelResponse) {
	response = &CancelInstanceLabelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CancelInstanceLabel(request *CancelInstanceLabelRequest) string {
	return c.CancelInstanceLabelWithContext(context.Background(), request)
}

func (c *Client) CancelInstanceLabelWithContext(ctx context.Context, request *CancelInstanceLabelRequest) string {
	if request == nil {
		request = NewCancelInstanceLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelInstanceLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateInstanceLabelRequest() (request *UpdateInstanceLabelRequest) {
	request = &UpdateInstanceLabelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "UpdateInstanceLabel")
	return
}

func NewUpdateInstanceLabelResponse() (response *UpdateInstanceLabelResponse) {
	response = &UpdateInstanceLabelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateInstanceLabel(request *UpdateInstanceLabelRequest) string {
	return c.UpdateInstanceLabelWithContext(context.Background(), request)
}

func (c *Client) UpdateInstanceLabelWithContext(ctx context.Context, request *UpdateInstanceLabelRequest) string {
	if request == nil {
		request = NewUpdateInstanceLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateInstanceLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewDeleteLabelRequest() (request *DeleteLabelRequest) {
	request = &DeleteLabelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "DeleteLabel")
	return
}

func NewDeleteLabelResponse() (response *DeleteLabelResponse) {
	response = &DeleteLabelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteLabel(request *DeleteLabelRequest) string {
	return c.DeleteLabelWithContext(context.Background(), request)
}

func (c *Client) DeleteLabelWithContext(ctx context.Context, request *DeleteLabelRequest) string {
	if request == nil {
		request = NewDeleteLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewUpdateLabelRequest() (request *UpdateLabelRequest) {
	request = &UpdateLabelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "UpdateLabel")
	return
}

func NewUpdateLabelResponse() (response *UpdateLabelResponse) {
	response = &UpdateLabelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) UpdateLabel(request *UpdateLabelRequest) string {
	return c.UpdateLabelWithContext(context.Background(), request)
}

func (c *Client) UpdateLabelWithContext(ctx context.Context, request *UpdateLabelRequest) string {
	if request == nil {
		request = NewUpdateLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
func NewCreateLabelRequest() (request *CreateLabelRequest) {
	request = &CreateLabelRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ked", APIVersion, "CreateLabel")
	return
}

func NewCreateLabelResponse() (response *CreateLabelResponse) {
	response = &CreateLabelResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateLabel(request *CreateLabelRequest) string {
	return c.CreateLabelWithContext(context.Background(), request)
}

func (c *Client) CreateLabelWithContext(ctx context.Context, request *CreateLabelRequest) string {
	if request == nil {
		request = NewCreateLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}
