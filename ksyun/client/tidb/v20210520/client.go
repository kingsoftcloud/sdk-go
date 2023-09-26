package v20210520
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2021-05-20"

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
    request.Init().WithApiInfo("tidb", APIVersion, "CreateInstance")
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
func NewListInstanceRequest() (request *ListInstanceRequest) {
    request = &ListInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "ListInstance")
    return
}

func NewListInstanceResponse() (response *ListInstanceResponse) {
    response = &ListInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListInstance(request *ListInstanceRequest) (string) {
    return c.ListInstanceWithContext(context.Background(), request)
}

func (c *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest) (string) {
    if request == nil {
        request = NewListInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListInstanceResponse()
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
    request.Init().WithApiInfo("tidb", APIVersion, "DescribeInstance")
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
func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
    request = &RenameInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "RenameInstance")
    return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
    response = &RenameInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenameInstance(request *RenameInstanceRequest) (string) {
    return c.RenameInstanceWithContext(context.Background(), request)
}

func (c *Client) RenameInstanceWithContext(ctx context.Context, request *RenameInstanceRequest) (string) {
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
func NewListRegionRequest() (request *ListRegionRequest) {
    request = &ListRegionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "ListRegion")
    return
}

func NewListRegionResponse() (response *ListRegionResponse) {
    response = &ListRegionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListRegion(request *ListRegionRequest) (string) {
    return c.ListRegionWithContext(context.Background(), request)
}

func (c *Client) ListRegionWithContext(ctx context.Context, request *ListRegionRequest) (string) {
    if request == nil {
        request = NewListRegionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListRegionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescRegionRequest() (request *DescRegionRequest) {
    request = &DescRegionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "DescRegion")
    return
}

func NewDescRegionResponse() (response *DescRegionResponse) {
    response = &DescRegionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescRegion(request *DescRegionRequest) (string) {
    return c.DescRegionWithContext(context.Background(), request)
}

func (c *Client) DescRegionWithContext(ctx context.Context, request *DescRegionRequest) (string) {
    if request == nil {
        request = NewDescRegionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescRegionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
    request = &CreateSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityGroup")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListSecurityGroupRequest() (request *ListSecurityGroupRequest) {
    request = &ListSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "ListSecurityGroup")
    return
}

func NewListSecurityGroupResponse() (response *ListSecurityGroupResponse) {
    response = &ListSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListSecurityGroup(request *ListSecurityGroupRequest) (string) {
    return c.ListSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) ListSecurityGroupWithContext(ctx context.Context, request *ListSecurityGroupRequest) (string) {
    if request == nil {
        request = NewListSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListSecurityGroupResponse()
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
    request.Init().WithApiInfo("tidb", APIVersion, "DescribeSecurityGroup")
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
func NewUpdateSecurityGroupRequest() (request *UpdateSecurityGroupRequest) {
    request = &UpdateSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityGroup")
    return
}

func NewUpdateSecurityGroupResponse() (response *UpdateSecurityGroupResponse) {
    response = &UpdateSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateSecurityGroup(request *UpdateSecurityGroupRequest) (string) {
    return c.UpdateSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) UpdateSecurityGroupWithContext(ctx context.Context, request *UpdateSecurityGroupRequest) (string) {
    if request == nil {
        request = NewUpdateSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewUpdateSecurityGroupResponse()
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
    request.Init().WithApiInfo("tidb", APIVersion, "CloneSecurityGroup")
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
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCloneSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewBindSecurityGroupRequest() (request *BindSecurityGroupRequest) {
    request = &BindSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "BindSecurityGroup")
    return
}

func NewBindSecurityGroupResponse() (response *BindSecurityGroupResponse) {
    response = &BindSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) BindSecurityGroup(request *BindSecurityGroupRequest) (string) {
    return c.BindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) BindSecurityGroupWithContext(ctx context.Context, request *BindSecurityGroupRequest) (string) {
    if request == nil {
        request = NewBindSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewBindSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUnbindSecurityGroupRequest() (request *UnbindSecurityGroupRequest) {
    request = &UnbindSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "UnbindSecurityGroup")
    return
}

func NewUnbindSecurityGroupResponse() (response *UnbindSecurityGroupResponse) {
    response = &UnbindSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UnbindSecurityGroup(request *UnbindSecurityGroupRequest) (string) {
    return c.UnbindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) UnbindSecurityGroupWithContext(ctx context.Context, request *UnbindSecurityGroupRequest) (string) {
    if request == nil {
        request = NewUnbindSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewUnbindSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRebindSecurityGroupRequest() (request *RebindSecurityGroupRequest) {
    request = &RebindSecurityGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "RebindSecurityGroup")
    return
}

func NewRebindSecurityGroupResponse() (response *RebindSecurityGroupResponse) {
    response = &RebindSecurityGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RebindSecurityGroup(request *RebindSecurityGroupRequest) (string) {
    return c.RebindSecurityGroupWithContext(context.Background(), request)
}

func (c *Client) RebindSecurityGroupWithContext(ctx context.Context, request *RebindSecurityGroupRequest) (string) {
    if request == nil {
        request = NewRebindSecurityGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewRebindSecurityGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateSecurityRuleRequest() (request *CreateSecurityRuleRequest) {
    request = &CreateSecurityRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "CreateSecurityRule")
    return
}

func NewCreateSecurityRuleResponse() (response *CreateSecurityRuleResponse) {
    response = &CreateSecurityRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateSecurityRule(request *CreateSecurityRuleRequest) (string) {
    return c.CreateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) CreateSecurityRuleWithContext(ctx context.Context, request *CreateSecurityRuleRequest) (string) {
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
func NewUpdateSecurityRuleRequest() (request *UpdateSecurityRuleRequest) {
    request = &UpdateSecurityRuleRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "UpdateSecurityRule")
    return
}

func NewUpdateSecurityRuleResponse() (response *UpdateSecurityRuleResponse) {
    response = &UpdateSecurityRuleResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateSecurityRule(request *UpdateSecurityRuleRequest) (string) {
    return c.UpdateSecurityRuleWithContext(context.Background(), request)
}

func (c *Client) UpdateSecurityRuleWithContext(ctx context.Context, request *UpdateSecurityRuleRequest) (string) {
    if request == nil {
        request = NewUpdateSecurityRuleRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewUpdateSecurityRuleResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListUnsecuredInstanceRequest() (request *ListUnsecuredInstanceRequest) {
    request = &ListUnsecuredInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "ListUnsecuredInstance")
    return
}

func NewListUnsecuredInstanceResponse() (response *ListUnsecuredInstanceResponse) {
    response = &ListUnsecuredInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListUnsecuredInstance(request *ListUnsecuredInstanceRequest) (string) {
    return c.ListUnsecuredInstanceWithContext(context.Background(), request)
}

func (c *Client) ListUnsecuredInstanceWithContext(ctx context.Context, request *ListUnsecuredInstanceRequest) (string) {
    if request == nil {
        request = NewListUnsecuredInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListUnsecuredInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListBackupRequest() (request *ListBackupRequest) {
    request = &ListBackupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tidb", APIVersion, "ListBackup")
    return
}

func NewListBackupResponse() (response *ListBackupResponse) {
    response = &ListBackupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListBackup(request *ListBackupRequest) (string) {
    return c.ListBackupWithContext(context.Background(), request)
}

func (c *Client) ListBackupWithContext(ctx context.Context, request *ListBackupRequest) (string) {
    if request == nil {
        request = NewListBackupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListBackupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


