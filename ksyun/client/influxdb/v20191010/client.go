package v20191010
import (
	"context"
	"fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2019-10-10"

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

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateInstance")
	return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstance(request *CreateInstanceRequest) string {
	return c.CreateInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceSend(request *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	statusCode, msg, err := c.CreateInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) string {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateInstanceWithContextV2(ctx context.Context, request *CreateInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteInstance(request *DeleteInstanceRequest) string {
	return c.DeleteInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceSend(request *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	statusCode, msg, err := c.DeleteInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) string {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteInstanceWithContextV2(ctx context.Context, request *DeleteInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeInstance")
	return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstance(request *DescribeInstanceRequest) string {
	return c.DescribeInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceSend(request *DescribeInstanceRequest) (*DescribeInstanceResponse, error) {
	statusCode, msg, err := c.DescribeInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) string {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceWithContextV2(ctx context.Context, request *DescribeInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceResponse()
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
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeInstances")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewDescribeInstancesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeInstanceNodeRequest() (request *DescribeInstanceNodeRequest) {
	request = &DescribeInstanceNodeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeInstanceNode")
	return
}

func NewDescribeInstanceNodeResponse() (response *DescribeInstanceNodeResponse) {
	response = &DescribeInstanceNodeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceNode(request *DescribeInstanceNodeRequest) string {
	return c.DescribeInstanceNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceNodeSend(request *DescribeInstanceNodeRequest) (*DescribeInstanceNodeResponse, error) {
	statusCode, msg, err := c.DescribeInstanceNodeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeInstanceNodeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeInstanceNodeWithContext(ctx context.Context, request *DescribeInstanceNodeRequest) string {
	if request == nil {
		request = NewDescribeInstanceNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceNodeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeInstanceNodeWithContextV2(ctx context.Context, request *DescribeInstanceNodeRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeInstanceNodeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeInstanceNodeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
	request = &RenameInstanceRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "RenameInstance")
	return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
	response = &RenameInstanceResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RenameInstance(request *RenameInstanceRequest) string {
	return c.RenameInstanceWithContext(context.Background(), request)
}

func (c *Client) RenameInstanceSend(request *RenameInstanceRequest) (*RenameInstanceResponse, error) {
	statusCode, msg, err := c.RenameInstanceWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RenameInstanceResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RenameInstanceWithContext(ctx context.Context, request *RenameInstanceRequest) string {
	if request == nil {
		request = NewRenameInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenameInstanceResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RenameInstanceWithContextV2(ctx context.Context, request *RenameInstanceRequest) (int, string, error) {
	if request == nil {
		request = NewRenameInstanceRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRenameInstanceResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeValidRegionsRequest() (request *DescribeValidRegionsRequest) {
	request = &DescribeValidRegionsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeValidRegions")
	return
}

func NewDescribeValidRegionsResponse() (response *DescribeValidRegionsResponse) {
	response = &DescribeValidRegionsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeValidRegions(request *DescribeValidRegionsRequest) string {
	return c.DescribeValidRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeValidRegionsSend(request *DescribeValidRegionsRequest) (*DescribeValidRegionsResponse, error) {
	statusCode, msg, err := c.DescribeValidRegionsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeValidRegionsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeValidRegionsWithContext(ctx context.Context, request *DescribeValidRegionsRequest) string {
	if request == nil {
		request = NewDescribeValidRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeValidRegionsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeValidRegionsWithContextV2(ctx context.Context, request *DescribeValidRegionsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeValidRegionsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeValidRegionsResponse()
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
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeSecurityGroup")
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
	request.SetContentType("application/json")

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
	request.SetContentType("application/json")

	response := NewDescribeSecurityGroupResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateSecurityRuleRequest() (request *CreateSecurityRuleRequest) {
	request = &CreateSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateSecurityRule")
	return
}

func NewCreateSecurityRuleResponse() (response *CreateSecurityRuleResponse) {
	response = &CreateSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityRule(request *CreateSecurityRuleRequest) string {
	return c.CreateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityRuleSend(request *CreateSecurityRuleRequest) (*CreateSecurityRuleResponse, error) {
	statusCode, msg, err := c.CreateSecurityRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateSecurityRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateSecurityRuleWithContext(ctx context.Context, request *CreateSecurityRuleRequest) string {
	if request == nil {
		request = NewCreateSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateSecurityRuleWithContextV2(ctx context.Context, request *CreateSecurityRuleRequest) (int, string, error) {
	if request == nil {
		request = NewCreateSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateSecurityRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteSecurityRuleRequest() (request *DeleteSecurityRuleRequest) {
	request = &DeleteSecurityRuleRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteSecurityRule")
	return
}

func NewDeleteSecurityRuleResponse() (response *DeleteSecurityRuleResponse) {
	response = &DeleteSecurityRuleResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityRule(request *DeleteSecurityRuleRequest) string {
	return c.DeleteSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityRuleSend(request *DeleteSecurityRuleRequest) (*DeleteSecurityRuleResponse, error) {
	statusCode, msg, err := c.DeleteSecurityRuleWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteSecurityRuleResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteSecurityRuleWithContext(ctx context.Context, request *DeleteSecurityRuleRequest) string {
	if request == nil {
		request = NewDeleteSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityRuleResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteSecurityRuleWithContextV2(ctx context.Context, request *DeleteSecurityRuleRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteSecurityRuleRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteSecurityRuleResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
	request = &DescribeDatabasesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeDatabases")
	return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
	response = &DescribeDatabasesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) string {
	return c.DescribeDatabasesWithContext(context.Background(), request)
}

func (c *Client) DescribeDatabasesSend(request *DescribeDatabasesRequest) (*DescribeDatabasesResponse, error) {
	statusCode, msg, err := c.DescribeDatabasesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeDatabasesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) string {
	if request == nil {
		request = NewDescribeDatabasesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDatabasesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeDatabasesWithContextV2(ctx context.Context, request *DescribeDatabasesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeDatabasesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeDatabasesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateDatabaseRequest() (request *CreateDatabaseRequest) {
	request = &CreateDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateDatabase")
	return
}

func NewCreateDatabaseResponse() (response *CreateDatabaseResponse) {
	response = &CreateDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDatabase(request *CreateDatabaseRequest) string {
	return c.CreateDatabaseWithContext(context.Background(), request)
}

func (c *Client) CreateDatabaseSend(request *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {
	statusCode, msg, err := c.CreateDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateDatabaseWithContext(ctx context.Context, request *CreateDatabaseRequest) string {
	if request == nil {
		request = NewCreateDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateDatabaseWithContextV2(ctx context.Context, request *CreateDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewCreateDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteDatabaseRequest() (request *DeleteDatabaseRequest) {
	request = &DeleteDatabaseRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteDatabase")
	return
}

func NewDeleteDatabaseResponse() (response *DeleteDatabaseResponse) {
	response = &DeleteDatabaseResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDatabase(request *DeleteDatabaseRequest) string {
	return c.DeleteDatabaseWithContext(context.Background(), request)
}

func (c *Client) DeleteDatabaseSend(request *DeleteDatabaseRequest) (*DeleteDatabaseResponse, error) {
	statusCode, msg, err := c.DeleteDatabaseWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteDatabaseResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteDatabaseWithContext(ctx context.Context, request *DeleteDatabaseRequest) string {
	if request == nil {
		request = NewDeleteDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDatabaseResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteDatabaseWithContextV2(ctx context.Context, request *DeleteDatabaseRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteDatabaseRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteDatabaseResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeRetentionPolicyListRequest() (request *DescribeRetentionPolicyListRequest) {
	request = &DescribeRetentionPolicyListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeRetentionPolicyList")
	return
}

func NewDescribeRetentionPolicyListResponse() (response *DescribeRetentionPolicyListResponse) {
	response = &DescribeRetentionPolicyListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRetentionPolicyList(request *DescribeRetentionPolicyListRequest) string {
	return c.DescribeRetentionPolicyListWithContext(context.Background(), request)
}

func (c *Client) DescribeRetentionPolicyListSend(request *DescribeRetentionPolicyListRequest) (*DescribeRetentionPolicyListResponse, error) {
	statusCode, msg, err := c.DescribeRetentionPolicyListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeRetentionPolicyListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeRetentionPolicyListWithContext(ctx context.Context, request *DescribeRetentionPolicyListRequest) string {
	if request == nil {
		request = NewDescribeRetentionPolicyListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRetentionPolicyListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeRetentionPolicyListWithContextV2(ctx context.Context, request *DescribeRetentionPolicyListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeRetentionPolicyListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeRetentionPolicyListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateRetentionPolicyRequest() (request *CreateRetentionPolicyRequest) {
	request = &CreateRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateRetentionPolicy")
	return
}

func NewCreateRetentionPolicyResponse() (response *CreateRetentionPolicyResponse) {
	response = &CreateRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateRetentionPolicy(request *CreateRetentionPolicyRequest) string {
	return c.CreateRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) CreateRetentionPolicySend(request *CreateRetentionPolicyRequest) (*CreateRetentionPolicyResponse, error) {
	statusCode, msg, err := c.CreateRetentionPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateRetentionPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateRetentionPolicyWithContext(ctx context.Context, request *CreateRetentionPolicyRequest) string {
	if request == nil {
		request = NewCreateRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateRetentionPolicyWithContextV2(ctx context.Context, request *CreateRetentionPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewCreateRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateRetentionPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteRetentionPolicyRequest() (request *DeleteRetentionPolicyRequest) {
	request = &DeleteRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteRetentionPolicy")
	return
}

func NewDeleteRetentionPolicyResponse() (response *DeleteRetentionPolicyResponse) {
	response = &DeleteRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRetentionPolicy(request *DeleteRetentionPolicyRequest) string {
	return c.DeleteRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) DeleteRetentionPolicySend(request *DeleteRetentionPolicyRequest) (*DeleteRetentionPolicyResponse, error) {
	statusCode, msg, err := c.DeleteRetentionPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteRetentionPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteRetentionPolicyWithContext(ctx context.Context, request *DeleteRetentionPolicyRequest) string {
	if request == nil {
		request = NewDeleteRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteRetentionPolicyWithContextV2(ctx context.Context, request *DeleteRetentionPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteRetentionPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewModifyRetentionPolicyRequest() (request *ModifyRetentionPolicyRequest) {
	request = &ModifyRetentionPolicyRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "ModifyRetentionPolicy")
	return
}

func NewModifyRetentionPolicyResponse() (response *ModifyRetentionPolicyResponse) {
	response = &ModifyRetentionPolicyResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyRetentionPolicy(request *ModifyRetentionPolicyRequest) string {
	return c.ModifyRetentionPolicyWithContext(context.Background(), request)
}

func (c *Client) ModifyRetentionPolicySend(request *ModifyRetentionPolicyRequest) (*ModifyRetentionPolicyResponse, error) {
	statusCode, msg, err := c.ModifyRetentionPolicyWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ModifyRetentionPolicyResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ModifyRetentionPolicyWithContext(ctx context.Context, request *ModifyRetentionPolicyRequest) string {
	if request == nil {
		request = NewModifyRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRetentionPolicyResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ModifyRetentionPolicyWithContextV2(ctx context.Context, request *ModifyRetentionPolicyRequest) (int, string, error) {
	if request == nil {
		request = NewModifyRetentionPolicyRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewModifyRetentionPolicyResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeMeasurementsRequest() (request *DescribeMeasurementsRequest) {
	request = &DescribeMeasurementsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeMeasurements")
	return
}

func NewDescribeMeasurementsResponse() (response *DescribeMeasurementsResponse) {
	response = &DescribeMeasurementsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMeasurements(request *DescribeMeasurementsRequest) string {
	return c.DescribeMeasurementsWithContext(context.Background(), request)
}

func (c *Client) DescribeMeasurementsSend(request *DescribeMeasurementsRequest) (*DescribeMeasurementsResponse, error) {
	statusCode, msg, err := c.DescribeMeasurementsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeMeasurementsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeMeasurementsWithContext(ctx context.Context, request *DescribeMeasurementsRequest) string {
	if request == nil {
		request = NewDescribeMeasurementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMeasurementsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeMeasurementsWithContextV2(ctx context.Context, request *DescribeMeasurementsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeMeasurementsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeMeasurementsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteMeasurementRequest() (request *DeleteMeasurementRequest) {
	request = &DeleteMeasurementRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteMeasurement")
	return
}

func NewDeleteMeasurementResponse() (response *DeleteMeasurementResponse) {
	response = &DeleteMeasurementResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteMeasurement(request *DeleteMeasurementRequest) string {
	return c.DeleteMeasurementWithContext(context.Background(), request)
}

func (c *Client) DeleteMeasurementSend(request *DeleteMeasurementRequest) (*DeleteMeasurementResponse, error) {
	statusCode, msg, err := c.DeleteMeasurementWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteMeasurementResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteMeasurementWithContext(ctx context.Context, request *DeleteMeasurementRequest) string {
	if request == nil {
		request = NewDeleteMeasurementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMeasurementResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteMeasurementWithContextV2(ctx context.Context, request *DeleteMeasurementRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteMeasurementRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteMeasurementResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeAccounts")
	return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) string {
	return c.DescribeAccountsWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountsSend(request *DescribeAccountsRequest) (*DescribeAccountsResponse, error) {
	statusCode, msg, err := c.DescribeAccountsWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAccountsResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) string {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountsResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAccountsWithContextV2(ctx context.Context, request *DescribeAccountsRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountsResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "CreateAccount")
	return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateAccount(request *CreateAccountRequest) string {
	return c.CreateAccountWithContext(context.Background(), request)
}

func (c *Client) CreateAccountSend(request *CreateAccountRequest) (*CreateAccountResponse, error) {
	statusCode, msg, err := c.CreateAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct CreateAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) string {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) CreateAccountWithContextV2(ctx context.Context, request *CreateAccountRequest) (int, string, error) {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewCreateAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
	request = &DeleteAccountRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DeleteAccount")
	return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
	response = &DeleteAccountResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteAccount(request *DeleteAccountRequest) string {
	return c.DeleteAccountWithContext(context.Background(), request)
}

func (c *Client) DeleteAccountSend(request *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	statusCode, msg, err := c.DeleteAccountWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DeleteAccountResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) string {
	if request == nil {
		request = NewDeleteAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAccountResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DeleteAccountWithContextV2(ctx context.Context, request *DeleteAccountRequest) (int, string, error) {
	if request == nil {
		request = NewDeleteAccountRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewDeleteAccountResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
	request = &DescribeAccountPrivilegesRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeAccountPrivileges")
	return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
	response = &DescribeAccountPrivilegesResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) string {
	return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountPrivilegesSend(request *DescribeAccountPrivilegesRequest) (*DescribeAccountPrivilegesResponse, error) {
	statusCode, msg, err := c.DescribeAccountPrivilegesWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAccountPrivilegesResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) string {
	if request == nil {
		request = NewDescribeAccountPrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountPrivilegesResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAccountPrivilegesWithContextV2(ctx context.Context, request *DescribeAccountPrivilegesRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAccountPrivilegesRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountPrivilegesResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewGrantAccountPrivilegeRequest() (request *GrantAccountPrivilegeRequest) {
	request = &GrantAccountPrivilegeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "GrantAccountPrivilege")
	return
}

func NewGrantAccountPrivilegeResponse() (response *GrantAccountPrivilegeResponse) {
	response = &GrantAccountPrivilegeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) string {
	return c.GrantAccountPrivilegeWithContext(context.Background(), request)
}

func (c *Client) GrantAccountPrivilegeSend(request *GrantAccountPrivilegeRequest) (*GrantAccountPrivilegeResponse, error) {
	statusCode, msg, err := c.GrantAccountPrivilegeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct GrantAccountPrivilegeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) GrantAccountPrivilegeWithContext(ctx context.Context, request *GrantAccountPrivilegeRequest) string {
	if request == nil {
		request = NewGrantAccountPrivilegeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGrantAccountPrivilegeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) GrantAccountPrivilegeWithContextV2(ctx context.Context, request *GrantAccountPrivilegeRequest) (int, string, error) {
	if request == nil {
		request = NewGrantAccountPrivilegeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewGrantAccountPrivilegeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewRevokeAccountPrivilegeRequest() (request *RevokeAccountPrivilegeRequest) {
	request = &RevokeAccountPrivilegeRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "RevokeAccountPrivilege")
	return
}

func NewRevokeAccountPrivilegeResponse() (response *RevokeAccountPrivilegeResponse) {
	response = &RevokeAccountPrivilegeResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) RevokeAccountPrivilege(request *RevokeAccountPrivilegeRequest) string {
	return c.RevokeAccountPrivilegeWithContext(context.Background(), request)
}

func (c *Client) RevokeAccountPrivilegeSend(request *RevokeAccountPrivilegeRequest) (*RevokeAccountPrivilegeResponse, error) {
	statusCode, msg, err := c.RevokeAccountPrivilegeWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct RevokeAccountPrivilegeResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) RevokeAccountPrivilegeWithContext(ctx context.Context, request *RevokeAccountPrivilegeRequest) string {
	if request == nil {
		request = NewRevokeAccountPrivilegeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRevokeAccountPrivilegeResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) RevokeAccountPrivilegeWithContextV2(ctx context.Context, request *RevokeAccountPrivilegeRequest) (int, string, error) {
	if request == nil {
		request = NewRevokeAccountPrivilegeRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewRevokeAccountPrivilegeResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
	request = &ResetAccountPasswordRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "ResetAccountPassword")
	return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
	response = &ResetAccountPasswordResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) string {
	return c.ResetAccountPasswordWithContext(context.Background(), request)
}

func (c *Client) ResetAccountPasswordSend(request *ResetAccountPasswordRequest) (*ResetAccountPasswordResponse, error) {
	statusCode, msg, err := c.ResetAccountPasswordWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct ResetAccountPasswordResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) string {
	if request == nil {
		request = NewResetAccountPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetAccountPasswordResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) ResetAccountPasswordWithContextV2(ctx context.Context, request *ResetAccountPasswordRequest) (int, string, error) {
	if request == nil {
		request = NewResetAccountPasswordRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/x-www-form-urlencoded")

	response := NewResetAccountPasswordResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}
func NewDescribeAccountDetailListRequest() (request *DescribeAccountDetailListRequest) {
	request = &DescribeAccountDetailListRequest{
		BaseRequest: &ksyunhttp.BaseRequest{},
	}
	request.Init().WithApiInfo("Influxdb", APIVersion, "DescribeAccountDetailList")
	return
}

func NewDescribeAccountDetailListResponse() (response *DescribeAccountDetailListResponse) {
	response = &DescribeAccountDetailListResponse{
		BaseResponse: &ksyunhttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAccountDetailList(request *DescribeAccountDetailListRequest) string {
	return c.DescribeAccountDetailListWithContext(context.Background(), request)
}

func (c *Client) DescribeAccountDetailListSend(request *DescribeAccountDetailListRequest) (*DescribeAccountDetailListResponse, error) {
	statusCode, msg, err := c.DescribeAccountDetailListWithContextV2(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:0 Err:%s] Request failed", err)
	}
	if statusCode < 200 || statusCode > 299 {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:Request failed] %s", statusCode, msg)
	}

	var respStruct DescribeAccountDetailListResponse
	err = respStruct.FromJsonString(msg)
	if err != nil {
		return nil, fmt.Errorf("[KsyunSDKError] [HttpCode:%d Err:%s] %s", statusCode, err.Error(), msg)
	}
	return &respStruct, nil
}

func (c *Client) DescribeAccountDetailListWithContext(ctx context.Context, request *DescribeAccountDetailListRequest) string {
	if request == nil {
		request = NewDescribeAccountDetailListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountDetailListResponse()
	err, msg := c.Send(request, response)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (c *Client) DescribeAccountDetailListWithContextV2(ctx context.Context, request *DescribeAccountDetailListRequest) (int, string, error) {
	if request == nil {
		request = NewDescribeAccountDetailListRequest()
	}
	request.SetContext(ctx)
	request.SetContentType("application/json")

	response := NewDescribeAccountDetailListResponse()
	statusCode, msg, err := c.SendV2(request, response)
	if err != nil {
		return statusCode, "", err
	}
	return statusCode, msg, nil
}


