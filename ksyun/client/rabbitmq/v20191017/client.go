package v20191017
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2019-10-17"

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
    request.Init().WithApiInfo("rabbitmq", APIVersion, "CreateInstance")
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateInstance(request *CreateInstanceRequest) (string) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (string) {
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
func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DeleteInstance")
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (string) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (string) {
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
func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (string) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (string) {
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
func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeInstance")
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (string) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (string) {
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
func NewDescribeInstanceNodesRequest() (request *DescribeInstanceNodesRequest) {
    request = &DescribeInstanceNodesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeInstanceNodes")
    return
}

func NewDescribeInstanceNodesResponse() (response *DescribeInstanceNodesResponse) {
    response = &DescribeInstanceNodesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceNodes(request *DescribeInstanceNodesRequest) (string) {
    return c.DescribeInstanceNodesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceNodesWithContext(ctx context.Context, request *DescribeInstanceNodesRequest) (string) {
    if request == nil {
        request = NewDescribeInstanceNodesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeInstanceNodesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeValidRegionRequest() (request *DescribeValidRegionRequest) {
    request = &DescribeValidRegionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeValidRegion")
    return
}

func NewDescribeValidRegionResponse() (response *DescribeValidRegionResponse) {
    response = &DescribeValidRegionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeValidRegion(request *DescribeValidRegionRequest) (string) {
    return c.DescribeValidRegionWithContext(context.Background(), request)
}

func (c *Client) DescribeValidRegionWithContext(ctx context.Context, request *DescribeValidRegionRequest) (string) {
    if request == nil {
        request = NewDescribeValidRegionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeValidRegionResponse()
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
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (string) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (string) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeRegionsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSecurityGroupRulesRequest() (request *DescribeSecurityGroupRulesRequest) {
    request = &DescribeSecurityGroupRulesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DescribeSecurityGroupRules")
    return
}

func NewDescribeSecurityGroupRulesResponse() (response *DescribeSecurityGroupRulesResponse) {
    response = &DescribeSecurityGroupRulesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSecurityGroupRules(request *DescribeSecurityGroupRulesRequest) (string) {
    return c.DescribeSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupRulesWithContext(ctx context.Context, request *DescribeSecurityGroupRulesRequest) (string) {
    if request == nil {
        request = NewDescribeSecurityGroupRulesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSecurityGroupRulesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddSecurityGroupRuleRequest() (request *AddSecurityGroupRuleRequest) {
    request = &AddSecurityGroupRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "AddSecurityGroupRule")
    return
}

func NewAddSecurityGroupRuleResponse() (response *AddSecurityGroupRuleResponse) {
    response = &AddSecurityGroupRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddSecurityGroupRule(request *AddSecurityGroupRuleRequest) (string) {
    return c.AddSecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) AddSecurityGroupRuleWithContext(ctx context.Context, request *AddSecurityGroupRuleRequest) (string) {
    if request == nil {
        request = NewAddSecurityGroupRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAddSecurityGroupRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteSecurityGroupRulesRequest() (request *DeleteSecurityGroupRulesRequest) {
    request = &DeleteSecurityGroupRulesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "DeleteSecurityGroupRules")
    return
}

func NewDeleteSecurityGroupRulesResponse() (response *DeleteSecurityGroupRulesResponse) {
    response = &DeleteSecurityGroupRulesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteSecurityGroupRules(request *DeleteSecurityGroupRulesRequest) (string) {
    return c.DeleteSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupRulesWithContext(ctx context.Context, request *DeleteSecurityGroupRulesRequest) (string) {
    if request == nil {
        request = NewDeleteSecurityGroupRulesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteSecurityGroupRulesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "ResetPassword")
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ResetPassword(request *ResetPasswordRequest) (string) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) (string) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewResetPasswordResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRenameRequest() (request *RenameRequest) {
    request = &RenameRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rabbitmq", APIVersion, "Rename")
    return
}

func NewRenameResponse() (response *RenameResponse) {
    response = &RenameResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) Rename(request *RenameRequest) (string) {
    return c.RenameWithContext(context.Background(), request)
}

func (c *Client) RenameWithContext(ctx context.Context, request *RenameRequest) (string) {
    if request == nil {
        request = NewRenameRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRenameResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


