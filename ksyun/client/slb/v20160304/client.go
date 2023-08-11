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

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeListeners")
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeListeners(request *DescribeListenersRequest) (string) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (string) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeListenersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyListenersRequest() (request *ModifyListenersRequest) {
    request = &ModifyListenersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyListeners")
    return
}

func NewModifyListenersResponse() (response *ModifyListenersResponse) {
    response = &ModifyListenersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyListeners(request *ModifyListenersRequest) (string) {
    return c.ModifyListenersWithContext(context.Background(), request)
}

func (c *Client) ModifyListenersWithContext(ctx context.Context, request *ModifyListenersRequest) (string) {
    if request == nil {
        request = NewModifyListenersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyListenersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateListenersRequest() (request *CreateListenersRequest) {
    request = &CreateListenersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateListeners")
    return
}

func NewCreateListenersResponse() (response *CreateListenersResponse) {
    response = &CreateListenersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateListeners(request *CreateListenersRequest) (string) {
    return c.CreateListenersWithContext(context.Background(), request)
}

func (c *Client) CreateListenersWithContext(ctx context.Context, request *CreateListenersRequest) (string) {
    if request == nil {
        request = NewCreateListenersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateListenersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyInstancesWithListenerRequest() (request *ModifyInstancesWithListenerRequest) {
    request = &ModifyInstancesWithListenerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyInstancesWithListener")
    return
}

func NewModifyInstancesWithListenerResponse() (response *ModifyInstancesWithListenerResponse) {
    response = &ModifyInstancesWithListenerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyInstancesWithListener(request *ModifyInstancesWithListenerRequest) (string) {
    return c.ModifyInstancesWithListenerWithContext(context.Background(), request)
}

func (c *Client) ModifyInstancesWithListenerWithContext(ctx context.Context, request *ModifyInstancesWithListenerRequest) (string) {
    if request == nil {
        request = NewModifyInstancesWithListenerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyInstancesWithListenerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRegisterInstancesWithListenerRequest() (request *RegisterInstancesWithListenerRequest) {
    request = &RegisterInstancesWithListenerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "RegisterInstancesWithListener")
    return
}

func NewRegisterInstancesWithListenerResponse() (response *RegisterInstancesWithListenerResponse) {
    response = &RegisterInstancesWithListenerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RegisterInstancesWithListener(request *RegisterInstancesWithListenerRequest) (string) {
    return c.RegisterInstancesWithListenerWithContext(context.Background(), request)
}

func (c *Client) RegisterInstancesWithListenerWithContext(ctx context.Context, request *RegisterInstancesWithListenerRequest) (string) {
    if request == nil {
        request = NewRegisterInstancesWithListenerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRegisterInstancesWithListenerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeregisterInstancesFromListenerRequest() (request *DeregisterInstancesFromListenerRequest) {
    request = &DeregisterInstancesFromListenerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeregisterInstancesFromListener")
    return
}

func NewDeregisterInstancesFromListenerResponse() (response *DeregisterInstancesFromListenerResponse) {
    response = &DeregisterInstancesFromListenerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeregisterInstancesFromListener(request *DeregisterInstancesFromListenerRequest) (string) {
    return c.DeregisterInstancesFromListenerWithContext(context.Background(), request)
}

func (c *Client) DeregisterInstancesFromListenerWithContext(ctx context.Context, request *DeregisterInstancesFromListenerRequest) (string) {
    if request == nil {
        request = NewDeregisterInstancesFromListenerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeregisterInstancesFromListenerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstancesWithListenerRequest() (request *DescribeInstancesWithListenerRequest) {
    request = &DescribeInstancesWithListenerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeInstancesWithListener")
    return
}

func NewDescribeInstancesWithListenerResponse() (response *DescribeInstancesWithListenerResponse) {
    response = &DescribeInstancesWithListenerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstancesWithListener(request *DescribeInstancesWithListenerRequest) (string) {
    return c.DescribeInstancesWithListenerWithContext(context.Background(), request)
}

func (c *Client) DescribeInstancesWithListenerWithContext(ctx context.Context, request *DescribeInstancesWithListenerRequest) (string) {
    if request == nil {
        request = NewDescribeInstancesWithListenerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeInstancesWithListenerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyHealthCheckRequest() (request *ModifyHealthCheckRequest) {
    request = &ModifyHealthCheckRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyHealthCheck")
    return
}

func NewModifyHealthCheckResponse() (response *ModifyHealthCheckResponse) {
    response = &ModifyHealthCheckResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyHealthCheck(request *ModifyHealthCheckRequest) (string) {
    return c.ModifyHealthCheckWithContext(context.Background(), request)
}

func (c *Client) ModifyHealthCheckWithContext(ctx context.Context, request *ModifyHealthCheckRequest) (string) {
    if request == nil {
        request = NewModifyHealthCheckRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyHealthCheckResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeHealthChecksRequest() (request *DescribeHealthChecksRequest) {
    request = &DescribeHealthChecksRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeHealthChecks")
    return
}

func NewDescribeHealthChecksResponse() (response *DescribeHealthChecksResponse) {
    response = &DescribeHealthChecksResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeHealthChecks(request *DescribeHealthChecksRequest) (string) {
    return c.DescribeHealthChecksWithContext(context.Background(), request)
}

func (c *Client) DescribeHealthChecksWithContext(ctx context.Context, request *DescribeHealthChecksRequest) (string) {
    if request == nil {
        request = NewDescribeHealthChecksRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeHealthChecksResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewConfigureHealthCheckRequest() (request *ConfigureHealthCheckRequest) {
    request = &ConfigureHealthCheckRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ConfigureHealthCheck")
    return
}

func NewConfigureHealthCheckResponse() (response *ConfigureHealthCheckResponse) {
    response = &ConfigureHealthCheckResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ConfigureHealthCheck(request *ConfigureHealthCheckRequest) (string) {
    return c.ConfigureHealthCheckWithContext(context.Background(), request)
}

func (c *Client) ConfigureHealthCheckWithContext(ctx context.Context, request *ConfigureHealthCheckRequest) (string) {
    if request == nil {
        request = NewConfigureHealthCheckRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewConfigureHealthCheckResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeLoadBalancers")
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (string) {
    return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) (string) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeLoadBalancersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyLoadBalancerRequest() (request *ModifyLoadBalancerRequest) {
    request = &ModifyLoadBalancerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyLoadBalancer")
    return
}

func NewModifyLoadBalancerResponse() (response *ModifyLoadBalancerResponse) {
    response = &ModifyLoadBalancerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyLoadBalancer(request *ModifyLoadBalancerRequest) (string) {
    return c.ModifyLoadBalancerWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancerWithContext(ctx context.Context, request *ModifyLoadBalancerRequest) (string) {
    if request == nil {
        request = NewModifyLoadBalancerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyLoadBalancerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
    request = &CreateLoadBalancerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateLoadBalancer")
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (string) {
    return c.CreateLoadBalancerWithContext(context.Background(), request)
}

func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) (string) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateLoadBalancerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateHostHeaderRequest() (request *CreateHostHeaderRequest) {
    request = &CreateHostHeaderRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateHostHeader")
    return
}

func NewCreateHostHeaderResponse() (response *CreateHostHeaderResponse) {
    response = &CreateHostHeaderResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateHostHeader(request *CreateHostHeaderRequest) (string) {
    return c.CreateHostHeaderWithContext(context.Background(), request)
}

func (c *Client) CreateHostHeaderWithContext(ctx context.Context, request *CreateHostHeaderRequest) (string) {
    if request == nil {
        request = NewCreateHostHeaderRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateHostHeaderResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyHostHeaderRequest() (request *ModifyHostHeaderRequest) {
    request = &ModifyHostHeaderRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyHostHeader")
    return
}

func NewModifyHostHeaderResponse() (response *ModifyHostHeaderResponse) {
    response = &ModifyHostHeaderResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyHostHeader(request *ModifyHostHeaderRequest) (string) {
    return c.ModifyHostHeaderWithContext(context.Background(), request)
}

func (c *Client) ModifyHostHeaderWithContext(ctx context.Context, request *ModifyHostHeaderRequest) (string) {
    if request == nil {
        request = NewModifyHostHeaderRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyHostHeaderResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteHostHeaderRequest() (request *DeleteHostHeaderRequest) {
    request = &DeleteHostHeaderRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeleteHostHeader")
    return
}

func NewDeleteHostHeaderResponse() (response *DeleteHostHeaderResponse) {
    response = &DeleteHostHeaderResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteHostHeader(request *DeleteHostHeaderRequest) (string) {
    return c.DeleteHostHeaderWithContext(context.Background(), request)
}

func (c *Client) DeleteHostHeaderWithContext(ctx context.Context, request *DeleteHostHeaderRequest) (string) {
    if request == nil {
        request = NewDeleteHostHeaderRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteHostHeaderResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeHostHeadersRequest() (request *DescribeHostHeadersRequest) {
    request = &DescribeHostHeadersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeHostHeaders")
    return
}

func NewDescribeHostHeadersResponse() (response *DescribeHostHeadersResponse) {
    response = &DescribeHostHeadersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeHostHeaders(request *DescribeHostHeadersRequest) (string) {
    return c.DescribeHostHeadersWithContext(context.Background(), request)
}

func (c *Client) DescribeHostHeadersWithContext(ctx context.Context, request *DescribeHostHeadersRequest) (string) {
    if request == nil {
        request = NewDescribeHostHeadersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeHostHeadersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeleteRule")
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteRule(request *DeleteRuleRequest) (string) {
    return c.DeleteRuleWithContext(context.Background(), request)
}

func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) (string) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeRules")
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeRules(request *DescribeRulesRequest) (string) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (string) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeRulesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateBackendServerGroupRequest() (request *CreateBackendServerGroupRequest) {
    request = &CreateBackendServerGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateBackendServerGroup")
    return
}

func NewCreateBackendServerGroupResponse() (response *CreateBackendServerGroupResponse) {
    response = &CreateBackendServerGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateBackendServerGroup(request *CreateBackendServerGroupRequest) (string) {
    return c.CreateBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) CreateBackendServerGroupWithContext(ctx context.Context, request *CreateBackendServerGroupRequest) (string) {
    if request == nil {
        request = NewCreateBackendServerGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateBackendServerGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteBackendServerGroupRequest() (request *DeleteBackendServerGroupRequest) {
    request = &DeleteBackendServerGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeleteBackendServerGroup")
    return
}

func NewDeleteBackendServerGroupResponse() (response *DeleteBackendServerGroupResponse) {
    response = &DeleteBackendServerGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteBackendServerGroup(request *DeleteBackendServerGroupRequest) (string) {
    return c.DeleteBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteBackendServerGroupWithContext(ctx context.Context, request *DeleteBackendServerGroupRequest) (string) {
    if request == nil {
        request = NewDeleteBackendServerGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteBackendServerGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyBackendServerGroupRequest() (request *ModifyBackendServerGroupRequest) {
    request = &ModifyBackendServerGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyBackendServerGroup")
    return
}

func NewModifyBackendServerGroupResponse() (response *ModifyBackendServerGroupResponse) {
    response = &ModifyBackendServerGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyBackendServerGroup(request *ModifyBackendServerGroupRequest) (string) {
    return c.ModifyBackendServerGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyBackendServerGroupWithContext(ctx context.Context, request *ModifyBackendServerGroupRequest) (string) {
    if request == nil {
        request = NewModifyBackendServerGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyBackendServerGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeBackendServerGroupsRequest() (request *DescribeBackendServerGroupsRequest) {
    request = &DescribeBackendServerGroupsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeBackendServerGroups")
    return
}

func NewDescribeBackendServerGroupsResponse() (response *DescribeBackendServerGroupsResponse) {
    response = &DescribeBackendServerGroupsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBackendServerGroups(request *DescribeBackendServerGroupsRequest) (string) {
    return c.DescribeBackendServerGroupsWithContext(context.Background(), request)
}

func (c *Client) DescribeBackendServerGroupsWithContext(ctx context.Context, request *DescribeBackendServerGroupsRequest) (string) {
    if request == nil {
        request = NewDescribeBackendServerGroupsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeBackendServerGroupsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRegisterBackendServerRequest() (request *RegisterBackendServerRequest) {
    request = &RegisterBackendServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "RegisterBackendServer")
    return
}

func NewRegisterBackendServerResponse() (response *RegisterBackendServerResponse) {
    response = &RegisterBackendServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RegisterBackendServer(request *RegisterBackendServerRequest) (string) {
    return c.RegisterBackendServerWithContext(context.Background(), request)
}

func (c *Client) RegisterBackendServerWithContext(ctx context.Context, request *RegisterBackendServerRequest) (string) {
    if request == nil {
        request = NewRegisterBackendServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRegisterBackendServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeregisterBackendServerRequest() (request *DeregisterBackendServerRequest) {
    request = &DeregisterBackendServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeregisterBackendServer")
    return
}

func NewDeregisterBackendServerResponse() (response *DeregisterBackendServerResponse) {
    response = &DeregisterBackendServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeregisterBackendServer(request *DeregisterBackendServerRequest) (string) {
    return c.DeregisterBackendServerWithContext(context.Background(), request)
}

func (c *Client) DeregisterBackendServerWithContext(ctx context.Context, request *DeregisterBackendServerRequest) (string) {
    if request == nil {
        request = NewDeregisterBackendServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeregisterBackendServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeBackendServersRequest() (request *DescribeBackendServersRequest) {
    request = &DescribeBackendServersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeBackendServers")
    return
}

func NewDescribeBackendServersResponse() (response *DescribeBackendServersResponse) {
    response = &DescribeBackendServersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBackendServers(request *DescribeBackendServersRequest) (string) {
    return c.DescribeBackendServersWithContext(context.Background(), request)
}

func (c *Client) DescribeBackendServersWithContext(ctx context.Context, request *DescribeBackendServersRequest) (string) {
    if request == nil {
        request = NewDescribeBackendServersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeBackendServersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateLoadBalancerAclRequest() (request *CreateLoadBalancerAclRequest) {
    request = &CreateLoadBalancerAclRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateLoadBalancerAcl")
    return
}

func NewCreateLoadBalancerAclResponse() (response *CreateLoadBalancerAclResponse) {
    response = &CreateLoadBalancerAclResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateLoadBalancerAcl(request *CreateLoadBalancerAclRequest) (string) {
    return c.CreateLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) CreateLoadBalancerAclWithContext(ctx context.Context, request *CreateLoadBalancerAclRequest) (string) {
    if request == nil {
        request = NewCreateLoadBalancerAclRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateLoadBalancerAclResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteLoadBalancerAclRequest() (request *DeleteLoadBalancerAclRequest) {
    request = &DeleteLoadBalancerAclRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeleteLoadBalancerAcl")
    return
}

func NewDeleteLoadBalancerAclResponse() (response *DeleteLoadBalancerAclResponse) {
    response = &DeleteLoadBalancerAclResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteLoadBalancerAcl(request *DeleteLoadBalancerAclRequest) (string) {
    return c.DeleteLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) DeleteLoadBalancerAclWithContext(ctx context.Context, request *DeleteLoadBalancerAclRequest) (string) {
    if request == nil {
        request = NewDeleteLoadBalancerAclRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteLoadBalancerAclResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyLoadBalancerAclRequest() (request *ModifyLoadBalancerAclRequest) {
    request = &ModifyLoadBalancerAclRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyLoadBalancerAcl")
    return
}

func NewModifyLoadBalancerAclResponse() (response *ModifyLoadBalancerAclResponse) {
    response = &ModifyLoadBalancerAclResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyLoadBalancerAcl(request *ModifyLoadBalancerAclRequest) (string) {
    return c.ModifyLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancerAclWithContext(ctx context.Context, request *ModifyLoadBalancerAclRequest) (string) {
    if request == nil {
        request = NewModifyLoadBalancerAclRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyLoadBalancerAclResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateLoadBalancerAclEntryRequest() (request *CreateLoadBalancerAclEntryRequest) {
    request = &CreateLoadBalancerAclEntryRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateLoadBalancerAclEntry")
    return
}

func NewCreateLoadBalancerAclEntryResponse() (response *CreateLoadBalancerAclEntryResponse) {
    response = &CreateLoadBalancerAclEntryResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateLoadBalancerAclEntry(request *CreateLoadBalancerAclEntryRequest) (string) {
    return c.CreateLoadBalancerAclEntryWithContext(context.Background(), request)
}

func (c *Client) CreateLoadBalancerAclEntryWithContext(ctx context.Context, request *CreateLoadBalancerAclEntryRequest) (string) {
    if request == nil {
        request = NewCreateLoadBalancerAclEntryRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateLoadBalancerAclEntryResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteLoadBalancerAclEntryRequest() (request *DeleteLoadBalancerAclEntryRequest) {
    request = &DeleteLoadBalancerAclEntryRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeleteLoadBalancerAclEntry")
    return
}

func NewDeleteLoadBalancerAclEntryResponse() (response *DeleteLoadBalancerAclEntryResponse) {
    response = &DeleteLoadBalancerAclEntryResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteLoadBalancerAclEntry(request *DeleteLoadBalancerAclEntryRequest) (string) {
    return c.DeleteLoadBalancerAclEntryWithContext(context.Background(), request)
}

func (c *Client) DeleteLoadBalancerAclEntryWithContext(ctx context.Context, request *DeleteLoadBalancerAclEntryRequest) (string) {
    if request == nil {
        request = NewDeleteLoadBalancerAclEntryRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteLoadBalancerAclEntryResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAssociateLoadBalancerAclRequest() (request *AssociateLoadBalancerAclRequest) {
    request = &AssociateLoadBalancerAclRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "AssociateLoadBalancerAcl")
    return
}

func NewAssociateLoadBalancerAclResponse() (response *AssociateLoadBalancerAclResponse) {
    response = &AssociateLoadBalancerAclResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AssociateLoadBalancerAcl(request *AssociateLoadBalancerAclRequest) (string) {
    return c.AssociateLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) AssociateLoadBalancerAclWithContext(ctx context.Context, request *AssociateLoadBalancerAclRequest) (string) {
    if request == nil {
        request = NewAssociateLoadBalancerAclRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAssociateLoadBalancerAclResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDisassociateLoadBalancerAclRequest() (request *DisassociateLoadBalancerAclRequest) {
    request = &DisassociateLoadBalancerAclRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DisassociateLoadBalancerAcl")
    return
}

func NewDisassociateLoadBalancerAclResponse() (response *DisassociateLoadBalancerAclResponse) {
    response = &DisassociateLoadBalancerAclResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DisassociateLoadBalancerAcl(request *DisassociateLoadBalancerAclRequest) (string) {
    return c.DisassociateLoadBalancerAclWithContext(context.Background(), request)
}

func (c *Client) DisassociateLoadBalancerAclWithContext(ctx context.Context, request *DisassociateLoadBalancerAclRequest) (string) {
    if request == nil {
        request = NewDisassociateLoadBalancerAclRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDisassociateLoadBalancerAclResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeLoadBalancerAclsRequest() (request *DescribeLoadBalancerAclsRequest) {
    request = &DescribeLoadBalancerAclsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribeLoadBalancerAcls")
    return
}

func NewDescribeLoadBalancerAclsResponse() (response *DescribeLoadBalancerAclsResponse) {
    response = &DescribeLoadBalancerAclsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeLoadBalancerAcls(request *DescribeLoadBalancerAclsRequest) (string) {
    return c.DescribeLoadBalancerAclsWithContext(context.Background(), request)
}

func (c *Client) DescribeLoadBalancerAclsWithContext(ctx context.Context, request *DescribeLoadBalancerAclsRequest) (string) {
    if request == nil {
        request = NewDescribeLoadBalancerAclsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribeLoadBalancerAclsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateSlbRuleRequest() (request *CreateSlbRuleRequest) {
    request = &CreateSlbRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreateSlbRule")
    return
}

func NewCreateSlbRuleResponse() (response *CreateSlbRuleResponse) {
    response = &CreateSlbRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateSlbRule(request *CreateSlbRuleRequest) (string) {
    return c.CreateSlbRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSlbRuleWithContext(ctx context.Context, request *CreateSlbRuleRequest) (string) {
    if request == nil {
        request = NewCreateSlbRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateSlbRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifySlbRuleRequest() (request *ModifySlbRuleRequest) {
    request = &ModifySlbRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifySlbRule")
    return
}

func NewModifySlbRuleResponse() (response *ModifySlbRuleResponse) {
    response = &ModifySlbRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifySlbRule(request *ModifySlbRuleRequest) (string) {
    return c.ModifySlbRuleWithContext(context.Background(), request)
}

func (c *Client) ModifySlbRuleWithContext(ctx context.Context, request *ModifySlbRuleRequest) (string) {
    if request == nil {
        request = NewModifySlbRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifySlbRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreatePrivateLinkServerRequest() (request *CreatePrivateLinkServerRequest) {
    request = &CreatePrivateLinkServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "CreatePrivateLinkServer")
    return
}

func NewCreatePrivateLinkServerResponse() (response *CreatePrivateLinkServerResponse) {
    response = &CreatePrivateLinkServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreatePrivateLinkServer(request *CreatePrivateLinkServerRequest) (string) {
    return c.CreatePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) CreatePrivateLinkServerWithContext(ctx context.Context, request *CreatePrivateLinkServerRequest) (string) {
    if request == nil {
        request = NewCreatePrivateLinkServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreatePrivateLinkServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribePrivateLinkServerRequest() (request *DescribePrivateLinkServerRequest) {
    request = &DescribePrivateLinkServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribePrivateLinkServer")
    return
}

func NewDescribePrivateLinkServerResponse() (response *DescribePrivateLinkServerResponse) {
    response = &DescribePrivateLinkServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribePrivateLinkServer(request *DescribePrivateLinkServerRequest) (string) {
    return c.DescribePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) DescribePrivateLinkServerWithContext(ctx context.Context, request *DescribePrivateLinkServerRequest) (string) {
    if request == nil {
        request = NewDescribePrivateLinkServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribePrivateLinkServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeletePrivateLinkServerRequest() (request *DeletePrivateLinkServerRequest) {
    request = &DeletePrivateLinkServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeletePrivateLinkServer")
    return
}

func NewDeletePrivateLinkServerResponse() (response *DeletePrivateLinkServerResponse) {
    response = &DeletePrivateLinkServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeletePrivateLinkServer(request *DeletePrivateLinkServerRequest) (string) {
    return c.DeletePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) DeletePrivateLinkServerWithContext(ctx context.Context, request *DeletePrivateLinkServerRequest) (string) {
    if request == nil {
        request = NewDeletePrivateLinkServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeletePrivateLinkServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyPrivateLinkServerRequest() (request *ModifyPrivateLinkServerRequest) {
    request = &ModifyPrivateLinkServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyPrivateLinkServer")
    return
}

func NewModifyPrivateLinkServerResponse() (response *ModifyPrivateLinkServerResponse) {
    response = &ModifyPrivateLinkServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyPrivateLinkServer(request *ModifyPrivateLinkServerRequest) (string) {
    return c.ModifyPrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) ModifyPrivateLinkServerWithContext(ctx context.Context, request *ModifyPrivateLinkServerRequest) (string) {
    if request == nil {
        request = NewModifyPrivateLinkServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyPrivateLinkServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAssociatePrivateLinkServerRequest() (request *AssociatePrivateLinkServerRequest) {
    request = &AssociatePrivateLinkServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "AssociatePrivateLinkServer")
    return
}

func NewAssociatePrivateLinkServerResponse() (response *AssociatePrivateLinkServerResponse) {
    response = &AssociatePrivateLinkServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AssociatePrivateLinkServer(request *AssociatePrivateLinkServerRequest) (string) {
    return c.AssociatePrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) AssociatePrivateLinkServerWithContext(ctx context.Context, request *AssociatePrivateLinkServerRequest) (string) {
    if request == nil {
        request = NewAssociatePrivateLinkServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAssociatePrivateLinkServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribePrivateLinkRequest() (request *DescribePrivateLinkRequest) {
    request = &DescribePrivateLinkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DescribePrivateLink")
    return
}

func NewDescribePrivateLinkResponse() (response *DescribePrivateLinkResponse) {
    response = &DescribePrivateLinkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribePrivateLink(request *DescribePrivateLinkRequest) (string) {
    return c.DescribePrivateLinkWithContext(context.Background(), request)
}

func (c *Client) DescribePrivateLinkWithContext(ctx context.Context, request *DescribePrivateLinkRequest) (string) {
    if request == nil {
        request = NewDescribePrivateLinkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDescribePrivateLinkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeletePrivateLinkRequest() (request *DeletePrivateLinkRequest) {
    request = &DeletePrivateLinkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "DeletePrivateLink")
    return
}

func NewDeletePrivateLinkResponse() (response *DeletePrivateLinkResponse) {
    response = &DeletePrivateLinkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeletePrivateLink(request *DeletePrivateLinkRequest) (string) {
    return c.DeletePrivateLinkWithContext(context.Background(), request)
}

func (c *Client) DeletePrivateLinkWithContext(ctx context.Context, request *DeletePrivateLinkRequest) (string) {
    if request == nil {
        request = NewDeletePrivateLinkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeletePrivateLinkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyLoadBalancerAclEntryRequest() (request *ModifyLoadBalancerAclEntryRequest) {
    request = &ModifyLoadBalancerAclEntryRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ModifyLoadBalancerAclEntry")
    return
}

func NewModifyLoadBalancerAclEntryResponse() (response *ModifyLoadBalancerAclEntryResponse) {
    response = &ModifyLoadBalancerAclEntryResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyLoadBalancerAclEntry(request *ModifyLoadBalancerAclEntryRequest) (string) {
    return c.ModifyLoadBalancerAclEntryWithContext(context.Background(), request)
}

func (c *Client) ModifyLoadBalancerAclEntryWithContext(ctx context.Context, request *ModifyLoadBalancerAclEntryRequest) (string) {
    if request == nil {
        request = NewModifyLoadBalancerAclEntryRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyLoadBalancerAclEntryResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAcceptPrivateLinkRequest() (request *AcceptPrivateLinkRequest) {
    request = &AcceptPrivateLinkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "AcceptPrivateLink")
    return
}

func NewAcceptPrivateLinkResponse() (response *AcceptPrivateLinkResponse) {
    response = &AcceptPrivateLinkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AcceptPrivateLink(request *AcceptPrivateLinkRequest) (string) {
    return c.AcceptPrivateLinkWithContext(context.Background(), request)
}

func (c *Client) AcceptPrivateLinkWithContext(ctx context.Context, request *AcceptPrivateLinkRequest) (string) {
    if request == nil {
        request = NewAcceptPrivateLinkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAcceptPrivateLinkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRejectPrivateLinkRequest() (request *RejectPrivateLinkRequest) {
    request = &RejectPrivateLinkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "RejectPrivateLink")
    return
}

func NewRejectPrivateLinkResponse() (response *RejectPrivateLinkResponse) {
    response = &RejectPrivateLinkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RejectPrivateLink(request *RejectPrivateLinkRequest) (string) {
    return c.RejectPrivateLinkWithContext(context.Background(), request)
}

func (c *Client) RejectPrivateLinkWithContext(ctx context.Context, request *RejectPrivateLinkRequest) (string) {
    if request == nil {
        request = NewRejectPrivateLinkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRejectPrivateLinkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListPrivateLinkServerRequest() (request *ListPrivateLinkServerRequest) {
    request = &ListPrivateLinkServerRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "ListPrivateLinkServer")
    return
}

func NewListPrivateLinkServerResponse() (response *ListPrivateLinkServerResponse) {
    response = &ListPrivateLinkServerResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListPrivateLinkServer(request *ListPrivateLinkServerRequest) (string) {
    return c.ListPrivateLinkServerWithContext(context.Background(), request)
}

func (c *Client) ListPrivateLinkServerWithContext(ctx context.Context, request *ListPrivateLinkServerRequest) (string) {
    if request == nil {
        request = NewListPrivateLinkServerRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewListPrivateLinkServerResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRemovePrivateLinkRequest() (request *RemovePrivateLinkRequest) {
    request = &RemovePrivateLinkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "RemovePrivateLink")
    return
}

func NewRemovePrivateLinkResponse() (response *RemovePrivateLinkResponse) {
    response = &RemovePrivateLinkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RemovePrivateLink(request *RemovePrivateLinkRequest) (string) {
    return c.RemovePrivateLinkWithContext(context.Background(), request)
}

func (c *Client) RemovePrivateLinkWithContext(ctx context.Context, request *RemovePrivateLinkRequest) (string) {
    if request == nil {
        request = NewRemovePrivateLinkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRemovePrivateLinkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetEnableAlbAccessLogRequest() (request *SetEnableAlbAccessLogRequest) {
    request = &SetEnableAlbAccessLogRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "SetEnableAlbAccessLog")
    return
}

func NewSetEnableAlbAccessLogResponse() (response *SetEnableAlbAccessLogResponse) {
    response = &SetEnableAlbAccessLogResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetEnableAlbAccessLog(request *SetEnableAlbAccessLogRequest) (string) {
    return c.SetEnableAlbAccessLogWithContext(context.Background(), request)
}

func (c *Client) SetEnableAlbAccessLogWithContext(ctx context.Context, request *SetEnableAlbAccessLogRequest) (string) {
    if request == nil {
        request = NewSetEnableAlbAccessLogRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewSetEnableAlbAccessLogResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetAlbAccessLogRequest() (request *SetAlbAccessLogRequest) {
    request = &SetAlbAccessLogRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("slb", APIVersion, "SetAlbAccessLog")
    return
}

func NewSetAlbAccessLogResponse() (response *SetAlbAccessLogResponse) {
    response = &SetAlbAccessLogResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetAlbAccessLog(request *SetAlbAccessLogRequest) (string) {
    return c.SetAlbAccessLogWithContext(context.Background(), request)
}

func (c *Client) SetAlbAccessLogWithContext(ctx context.Context, request *SetAlbAccessLogRequest) (string) {
    if request == nil {
        request = NewSetAlbAccessLogRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSetAlbAccessLogResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


