package v20200825
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-08-25"

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

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
    request = &CreateSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateSecurityGroup")
    return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
    response = &CreateSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (string) {
    return c.CreateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) (string) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCreateSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
    request = &DescribeSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeSecurityGroup")
    return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
    response = &DescribeSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (string) {
    return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) (string) {
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
func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
    request = &DeleteSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DeleteSecurityGroup")
    return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
    response = &DeleteSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (string) {
    return c.DeleteSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) (string) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
    request = &ModifySecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroup")
    return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
    response = &ModifySecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) (string) {
    return c.ModifySecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupWithContext(ctx context.Context, request *ModifySecurityGroupRequest) (string) {
    if request == nil {
        request = NewModifySecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifySecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
    request = &CloneSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CloneSecurityGroup")
    return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
    response = &CloneSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) (string) {
    return c.CloneSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) (string) {
    if request == nil {
        request = NewCloneSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCloneSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifySecurityGroupRuleRequest() (request *ModifySecurityGroupRuleRequest) {
    request = &ModifySecurityGroupRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRule")
    return
}

func NewModifySecurityGroupRuleResponse() (response *ModifySecurityGroupRuleResponse) {
    response = &ModifySecurityGroupRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifySecurityGroupRule(request *ModifySecurityGroupRuleRequest) (string) {
    return c.ModifySecurityGroupRuleWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupRuleWithContext(ctx context.Context, request *ModifySecurityGroupRuleRequest) (string) {
    if request == nil {
        request = NewModifySecurityGroupRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifySecurityGroupRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSecurityGroupRelationRequest() (request *SecurityGroupRelationRequest) {
    request = &SecurityGroupRelationRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SecurityGroupRelation")
    return
}

func NewSecurityGroupRelationResponse() (response *SecurityGroupRelationResponse) {
    response = &SecurityGroupRelationResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SecurityGroupRelation(request *SecurityGroupRelationRequest) (string) {
    return c.SecurityGroupRelationWithContext(context.Background(), request)
}

func (c *Client) SecurityGroupRelationWithContext(ctx context.Context, request *SecurityGroupRelationRequest) (string) {
    if request == nil {
        request = NewSecurityGroupRelationRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSecurityGroupRelationResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifySecurityGroupRuleNameRequest() (request *ModifySecurityGroupRuleNameRequest) {
    request = &ModifySecurityGroupRuleNameRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifySecurityGroupRuleName")
    return
}

func NewModifySecurityGroupRuleNameResponse() (response *ModifySecurityGroupRuleNameResponse) {
    response = &ModifySecurityGroupRuleNameResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifySecurityGroupRuleName(request *ModifySecurityGroupRuleNameRequest) (string) {
    return c.ModifySecurityGroupRuleNameWithContext(context.Background(), request)
}

func (c *Client) ModifySecurityGroupRuleNameWithContext(ctx context.Context, request *ModifySecurityGroupRuleNameRequest) (string) {
    if request == nil {
        request = NewModifySecurityGroupRuleNameRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifySecurityGroupRuleNameResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


