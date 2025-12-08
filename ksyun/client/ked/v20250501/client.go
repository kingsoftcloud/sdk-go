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

func (c *Client) CloudDeskreinstallSend(request *CloudDeskreinstallRequest) (*CloudDeskreinstallResponse, error) {
	statusCode, msg, err := c.CloudDeskreinstallWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloudDeskreinstallResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CloudDeskreinstallWithContextV2(ctx context.Context, request *CloudDeskreinstallRequest) (int, string, error) {
	if request == nil {
		request = NewCloudDeskreinstallRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskreinstallResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CloudDeskmanageSend(request *CloudDeskmanageRequest) (*CloudDeskmanageResponse, error) {
	statusCode, msg, err := c.CloudDeskmanageWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloudDeskmanageResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CloudDeskmanageWithContextV2(ctx context.Context, request *CloudDeskmanageRequest) (int, string, error) {
	if request == nil {
		request = NewCloudDeskmanageRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskmanageResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CloudDeskeditSend(request *CloudDeskeditRequest) (*CloudDeskeditResponse, error) {
	statusCode, msg, err := c.CloudDeskeditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloudDeskeditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CloudDeskeditWithContextV2(ctx context.Context, request *CloudDeskeditRequest) (int, string, error) {
	if request == nil {
		request = NewCloudDeskeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskeditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CloudDeskcreateSend(request *CloudDeskcreateRequest) (*CloudDeskcreateResponse, error) {
	statusCode, msg, err := c.CloudDeskcreateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloudDeskcreateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CloudDeskcreateWithContextV2(ctx context.Context, request *CloudDeskcreateRequest) (int, string, error) {
	if request == nil {
		request = NewCloudDeskcreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskcreateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CloudDesklistSend(request *CloudDesklistRequest) (*CloudDesklistResponse, error) {
	statusCode, msg, err := c.CloudDesklistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloudDesklistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CloudDesklistWithContext(ctx context.Context, request *CloudDesklistRequest) string {
	if request == nil {
		request = NewCloudDesklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDesklistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CloudDesklistWithContextV2(ctx context.Context, request *CloudDesklistRequest) (int, string, error) {
	if request == nil {
		request = NewCloudDesklistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDesklistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategyruleeditSend(request *StrategyruleeditRequest) (*StrategyruleeditResponse, error) {
	statusCode, msg, err := c.StrategyruleeditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategyruleeditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategyruleeditWithContextV2(ctx context.Context, request *StrategyruleeditRequest) (int, string, error) {
	if request == nil {
		request = NewStrategyruleeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyruleeditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategyrulecreateSend(request *StrategyrulecreateRequest) (*StrategyrulecreateResponse, error) {
	statusCode, msg, err := c.StrategyrulecreateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategyrulecreateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategyrulecreateWithContextV2(ctx context.Context, request *StrategyrulecreateRequest) (int, string, error) {
	if request == nil {
		request = NewStrategyrulecreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyrulecreateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategyunboundSend(request *StrategyunboundRequest) (*StrategyunboundResponse, error) {
	statusCode, msg, err := c.StrategyunboundWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategyunboundResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategyunboundWithContextV2(ctx context.Context, request *StrategyunboundRequest) (int, string, error) {
	if request == nil {
		request = NewStrategyunboundRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyunboundResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategyboundSend(request *StrategyboundRequest) (*StrategyboundResponse, error) {
	statusCode, msg, err := c.StrategyboundWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategyboundResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategyboundWithContextV2(ctx context.Context, request *StrategyboundRequest) (int, string, error) {
	if request == nil {
		request = NewStrategyboundRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyboundResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategydeleteSend(request *StrategydeleteRequest) (*StrategydeleteResponse, error) {
	statusCode, msg, err := c.StrategydeleteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategydeleteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategydeleteWithContextV2(ctx context.Context, request *StrategydeleteRequest) (int, string, error) {
	if request == nil {
		request = NewStrategydeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategydeleteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategyeditSend(request *StrategyeditRequest) (*StrategyeditResponse, error) {
	statusCode, msg, err := c.StrategyeditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategyeditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategyeditWithContextV2(ctx context.Context, request *StrategyeditRequest) (int, string, error) {
	if request == nil {
		request = NewStrategyeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyeditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategycreateSend(request *StrategycreateRequest) (*StrategycreateResponse, error) {
	statusCode, msg, err := c.StrategycreateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategycreateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategycreateWithContextV2(ctx context.Context, request *StrategycreateRequest) (int, string, error) {
	if request == nil {
		request = NewStrategycreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategycreateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategylistSend(request *StrategylistRequest) (*StrategylistResponse, error) {
	statusCode, msg, err := c.StrategylistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategylistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) StrategylistWithContext(ctx context.Context, request *StrategylistRequest) string {
	if request == nil {
		request = NewStrategylistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategylistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) StrategylistWithContextV2(ctx context.Context, request *StrategylistRequest) (int, string, error) {
	if request == nil {
		request = NewStrategylistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategylistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) RolesdeleteSend(request *RolesdeleteRequest) (*RolesdeleteResponse, error) {
	statusCode, msg, err := c.RolesdeleteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RolesdeleteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) RolesdeleteWithContextV2(ctx context.Context, request *RolesdeleteRequest) (int, string, error) {
	if request == nil {
		request = NewRolesdeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRolesdeleteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) RoleseditSend(request *RoleseditRequest) (*RoleseditResponse, error) {
	statusCode, msg, err := c.RoleseditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RoleseditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) RoleseditWithContextV2(ctx context.Context, request *RoleseditRequest) (int, string, error) {
	if request == nil {
		request = NewRoleseditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRoleseditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) RolescreateSend(request *RolescreateRequest) (*RolescreateResponse, error) {
	statusCode, msg, err := c.RolescreateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RolescreateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) RolescreateWithContextV2(ctx context.Context, request *RolescreateRequest) (int, string, error) {
	if request == nil {
		request = NewRolescreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRolescreateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) RoleslistSend(request *RoleslistRequest) (*RoleslistResponse, error) {
	statusCode, msg, err := c.RoleslistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RoleslistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RoleslistWithContext(ctx context.Context, request *RoleslistRequest) string {
	if request == nil {
		request = NewRoleslistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRoleslistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RoleslistWithContextV2(ctx context.Context, request *RoleslistRequest) (int, string, error) {
	if request == nil {
		request = NewRoleslistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewRoleslistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ImagedeleteSend(request *ImagedeleteRequest) (*ImagedeleteResponse, error) {
	statusCode, msg, err := c.ImagedeleteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ImagedeleteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ImagedeleteWithContextV2(ctx context.Context, request *ImagedeleteRequest) (int, string, error) {
	if request == nil {
		request = NewImagedeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImagedeleteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ImageeditSend(request *ImageeditRequest) (*ImageeditResponse, error) {
	statusCode, msg, err := c.ImageeditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ImageeditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ImageeditWithContextV2(ctx context.Context, request *ImageeditRequest) (int, string, error) {
	if request == nil {
		request = NewImageeditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImageeditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ImagecreateSend(request *ImagecreateRequest) (*ImagecreateResponse, error) {
	statusCode, msg, err := c.ImagecreateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ImagecreateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) ImagecreateWithContextV2(ctx context.Context, request *ImagecreateRequest) (int, string, error) {
	if request == nil {
		request = NewImagecreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImagecreateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ImagelistSend(request *ImagelistRequest) (*ImagelistResponse, error) {
	statusCode, msg, err := c.ImagelistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ImagelistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ImagelistWithContext(ctx context.Context, request *ImagelistRequest) string {
	if request == nil {
		request = NewImagelistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImagelistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ImagelistWithContextV2(ctx context.Context, request *ImagelistRequest) (int, string, error) {
	if request == nil {
		request = NewImagelistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewImagelistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) StrategyrulebatchEditSend(request *StrategyrulebatchEditRequest) (*StrategyrulebatchEditResponse, error) {
	statusCode, msg, err := c.StrategyrulebatchEditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct StrategyrulebatchEditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) StrategyrulebatchEditWithContextV2(ctx context.Context, request *StrategyrulebatchEditRequest) (int, string, error) {
	if request == nil {
		request = NewStrategyrulebatchEditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewStrategyrulebatchEditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) MonitorregionsSend(request *MonitorregionsRequest) (*MonitorregionsResponse, error) {
	statusCode, msg, err := c.MonitorregionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct MonitorregionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) MonitorregionsWithContext(ctx context.Context, request *MonitorregionsRequest) string {
	if request == nil {
		request = NewMonitorregionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewMonitorregionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) MonitorregionsWithContextV2(ctx context.Context, request *MonitorregionsRequest) (int, string, error) {
	if request == nil {
		request = NewMonitorregionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewMonitorregionsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UsersinstancebindSend(request *UsersinstancebindRequest) (*UsersinstancebindResponse, error) {
	statusCode, msg, err := c.UsersinstancebindWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UsersinstancebindResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UsersinstancebindWithContextV2(ctx context.Context, request *UsersinstancebindRequest) (int, string, error) {
	if request == nil {
		request = NewUsersinstancebindRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUsersinstancebindResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UserspasswordresetSend(request *UserspasswordresetRequest) (*UserspasswordresetResponse, error) {
	statusCode, msg, err := c.UserspasswordresetWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UserspasswordresetResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UserspasswordresetWithContextV2(ctx context.Context, request *UserspasswordresetRequest) (int, string, error) {
	if request == nil {
		request = NewUserspasswordresetRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserspasswordresetResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UsersdeleteSend(request *UsersdeleteRequest) (*UsersdeleteResponse, error) {
	statusCode, msg, err := c.UsersdeleteWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UsersdeleteResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UsersdeleteWithContextV2(ctx context.Context, request *UsersdeleteRequest) (int, string, error) {
	if request == nil {
		request = NewUsersdeleteRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUsersdeleteResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UserseditSend(request *UserseditRequest) (*UserseditResponse, error) {
	statusCode, msg, err := c.UserseditWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UserseditResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UserseditWithContextV2(ctx context.Context, request *UserseditRequest) (int, string, error) {
	if request == nil {
		request = NewUserseditRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserseditResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UserscreateSend(request *UserscreateRequest) (*UserscreateResponse, error) {
	statusCode, msg, err := c.UserscreateWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UserscreateResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UserscreateWithContextV2(ctx context.Context, request *UserscreateRequest) (int, string, error) {
	if request == nil {
		request = NewUserscreateRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserscreateResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UserslistSend(request *UserslistRequest) (*UserslistResponse, error) {
	statusCode, msg, err := c.UserslistWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UserslistResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) UserslistWithContext(ctx context.Context, request *UserslistRequest) string {
	if request == nil {
		request = NewUserslistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserslistResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) UserslistWithContextV2(ctx context.Context, request *UserslistRequest) (int, string, error) {
	if request == nil {
		request = NewUserslistRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUserslistResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CloudDeskgetDesktopUrlSend(request *CloudDeskgetDesktopUrlRequest) (*CloudDeskgetDesktopUrlResponse, error) {
	statusCode, msg, err := c.CloudDeskgetDesktopUrlWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CloudDeskgetDesktopUrlResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CloudDeskgetDesktopUrlWithContextV2(ctx context.Context, request *CloudDeskgetDesktopUrlRequest) (int, string, error) {
	if request == nil {
		request = NewCloudDeskgetDesktopUrlRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCloudDeskgetDesktopUrlResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) QueryCloudDesksubmitShellSend(request *QueryCloudDesksubmitShellRequest) (*QueryCloudDesksubmitShellResponse, error) {
	statusCode, msg, err := c.QueryCloudDesksubmitShellWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryCloudDesksubmitShellResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) QueryCloudDesksubmitShellWithContextV2(ctx context.Context, request *QueryCloudDesksubmitShellRequest) (int, string, error) {
	if request == nil {
		request = NewQueryCloudDesksubmitShellRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryCloudDesksubmitShellResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateCloudDeskgetTokenSend(request *CreateCloudDeskgetTokenRequest) (*CreateCloudDeskgetTokenResponse, error) {
	statusCode, msg, err := c.CreateCloudDeskgetTokenWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateCloudDeskgetTokenResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateCloudDeskgetTokenWithContextV2(ctx context.Context, request *CreateCloudDeskgetTokenRequest) (int, string, error) {
	if request == nil {
		request = NewCreateCloudDeskgetTokenRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateCloudDeskgetTokenResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) QueryShellStatusSend(request *QueryShellStatusRequest) (*QueryShellStatusResponse, error) {
	statusCode, msg, err := c.QueryShellStatusWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryShellStatusResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryShellStatusWithContext(ctx context.Context, request *QueryShellStatusRequest) string {
	if request == nil {
		request = NewQueryShellStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryShellStatusResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryShellStatusWithContextV2(ctx context.Context, request *QueryShellStatusRequest) (int, string, error) {
	if request == nil {
		request = NewQueryShellStatusRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryShellStatusResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) SetProxyIpSend(request *SetProxyIpRequest) (*SetProxyIpResponse, error) {
	statusCode, msg, err := c.SetProxyIpWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct SetProxyIpResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) SetProxyIpWithContextV2(ctx context.Context, request *SetProxyIpRequest) (int, string, error) {
	if request == nil {
		request = NewSetProxyIpRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewSetProxyIpResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetProxyConfigSend(request *GetProxyConfigRequest) (*GetProxyConfigResponse, error) {
	statusCode, msg, err := c.GetProxyConfigWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetProxyConfigResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetProxyConfigWithContext(ctx context.Context, request *GetProxyConfigRequest) string {
	if request == nil {
		request = NewGetProxyConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetProxyConfigResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetProxyConfigWithContextV2(ctx context.Context, request *GetProxyConfigRequest) (int, string, error) {
	if request == nil {
		request = NewGetProxyConfigRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetProxyConfigResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) QueryRuledetailSend(request *QueryRuledetailRequest) (*QueryRuledetailResponse, error) {
	statusCode, msg, err := c.QueryRuledetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryRuledetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) QueryRuledetailWithContext(ctx context.Context, request *QueryRuledetailRequest) string {
	if request == nil {
		request = NewQueryRuledetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryRuledetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) QueryRuledetailWithContextV2(ctx context.Context, request *QueryRuledetailRequest) (int, string, error) {
	if request == nil {
		request = NewQueryRuledetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryRuledetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) QueryUsersinfoSend(request *QueryUsersinfoRequest) (*QueryUsersinfoResponse, error) {
	statusCode, msg, err := c.QueryUsersinfoWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct QueryUsersinfoResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) QueryUsersinfoWithContextV2(ctx context.Context, request *QueryUsersinfoRequest) (int, string, error) {
	if request == nil {
		request = NewQueryUsersinfoRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewQueryUsersinfoResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) GetDetailSend(request *GetDetailRequest) (*GetDetailResponse, error) {
	statusCode, msg, err := c.GetDetailWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GetDetailResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GetDetailWithContext(ctx context.Context, request *GetDetailRequest) string {
	if request == nil {
		request = NewGetDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDetailResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GetDetailWithContextV2(ctx context.Context, request *GetDetailRequest) (int, string, error) {
	if request == nil {
		request = NewGetDetailRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewGetDetailResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) ListLabelSend(request *ListLabelRequest) (*ListLabelResponse, error) {
	statusCode, msg, err := c.ListLabelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ListLabelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ListLabelWithContext(ctx context.Context, request *ListLabelRequest) string {
	if request == nil {
		request = NewListLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListLabelResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ListLabelWithContextV2(ctx context.Context, request *ListLabelRequest) (int, string, error) {
	if request == nil {
		request = NewListLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewListLabelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CancelInstanceLabelSend(request *CancelInstanceLabelRequest) (*CancelInstanceLabelResponse, error) {
	statusCode, msg, err := c.CancelInstanceLabelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CancelInstanceLabelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CancelInstanceLabelWithContextV2(ctx context.Context, request *CancelInstanceLabelRequest) (int, string, error) {
	if request == nil {
		request = NewCancelInstanceLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCancelInstanceLabelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateInstanceLabelSend(request *UpdateInstanceLabelRequest) (*UpdateInstanceLabelResponse, error) {
	statusCode, msg, err := c.UpdateInstanceLabelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateInstanceLabelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateInstanceLabelWithContextV2(ctx context.Context, request *UpdateInstanceLabelRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateInstanceLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateInstanceLabelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) DeleteLabelSend(request *DeleteLabelRequest) (*DeleteLabelResponse, error) {
	statusCode, msg, err := c.DeleteLabelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteLabelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) DeleteLabelWithContextV2(ctx context.Context, request *DeleteLabelRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDeleteLabelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) UpdateLabelSend(request *UpdateLabelRequest) (*UpdateLabelResponse, error) {
	statusCode, msg, err := c.UpdateLabelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct UpdateLabelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) UpdateLabelWithContextV2(ctx context.Context, request *UpdateLabelRequest) (int, string, error) {
	if request == nil {
		request = NewUpdateLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewUpdateLabelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
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

func (c *Client) CreateLabelSend(request *CreateLabelRequest) (*CreateLabelResponse, error) {
	statusCode, msg, err := c.CreateLabelWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateLabelResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
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

func (c *Client) CreateLabelWithContextV2(ctx context.Context, request *CreateLabelRequest) (int, string, error) {
	if request == nil {
		request = NewCreateLabelRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewCreateLabelResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


