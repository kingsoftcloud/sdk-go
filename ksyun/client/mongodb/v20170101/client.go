package v20170101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2017-01-01"

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

func NewCreateMongoDBInstanceRequest() (request *CreateMongoDBInstanceRequest) {
    request = &CreateMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateMongoDBInstance")
    return
}

func NewCreateMongoDBInstanceResponse() (response *CreateMongoDBInstanceResponse) {
    response = &CreateMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateMongoDBInstance(request *CreateMongoDBInstanceRequest) (string) {
    return c.CreateMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateMongoDBInstanceWithContext(ctx context.Context, request *CreateMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewCreateMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteMongoDBInstanceRequest() (request *DeleteMongoDBInstanceRequest) {
    request = &DeleteMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteMongoDBInstance")
    return
}

func NewDeleteMongoDBInstanceResponse() (response *DeleteMongoDBInstanceResponse) {
    response = &DeleteMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteMongoDBInstance(request *DeleteMongoDBInstanceRequest) (string) {
    return c.DeleteMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteMongoDBInstanceWithContext(ctx context.Context, request *DeleteMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewDeleteMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeMongoDBInstanceRequest() (request *DescribeMongoDBInstanceRequest) {
    request = &DescribeMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBInstance")
    return
}

func NewDescribeMongoDBInstanceResponse() (response *DescribeMongoDBInstanceResponse) {
    response = &DescribeMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeMongoDBInstance(request *DescribeMongoDBInstanceRequest) (string) {
    return c.DescribeMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBInstanceWithContext(ctx context.Context, request *DescribeMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewDescribeMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeMongoDBInstancesRequest() (request *DescribeMongoDBInstancesRequest) {
    request = &DescribeMongoDBInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBInstances")
    return
}

func NewDescribeMongoDBInstancesResponse() (response *DescribeMongoDBInstancesResponse) {
    response = &DescribeMongoDBInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeMongoDBInstances(request *DescribeMongoDBInstancesRequest) (string) {
    return c.DescribeMongoDBInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBInstancesWithContext(ctx context.Context, request *DescribeMongoDBInstancesRequest) (string) {
    if request == nil {
        request = NewDescribeMongoDBInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeMongoDBInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeMongoDBInstanceNodeRequest() (request *DescribeMongoDBInstanceNodeRequest) {
    request = &DescribeMongoDBInstanceNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBInstanceNode")
    return
}

func NewDescribeMongoDBInstanceNodeResponse() (response *DescribeMongoDBInstanceNodeResponse) {
    response = &DescribeMongoDBInstanceNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeMongoDBInstanceNode(request *DescribeMongoDBInstanceNodeRequest) (string) {
    return c.DescribeMongoDBInstanceNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBInstanceNodeWithContext(ctx context.Context, request *DescribeMongoDBInstanceNodeRequest) (string) {
    if request == nil {
        request = NewDescribeMongoDBInstanceNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeMongoDBInstanceNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRenameMongoDBInstanceRequest() (request *RenameMongoDBInstanceRequest) {
    request = &RenameMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "RenameMongoDBInstance")
    return
}

func NewRenameMongoDBInstanceResponse() (response *RenameMongoDBInstanceResponse) {
    response = &RenameMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenameMongoDBInstance(request *RenameMongoDBInstanceRequest) (string) {
    return c.RenameMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RenameMongoDBInstanceWithContext(ctx context.Context, request *RenameMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewRenameMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRenameMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewResetPasswordMongoDBInstanceRequest() (request *ResetPasswordMongoDBInstanceRequest) {
    request = &ResetPasswordMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "ResetPasswordMongoDBInstance")
    return
}

func NewResetPasswordMongoDBInstanceResponse() (response *ResetPasswordMongoDBInstanceResponse) {
    response = &ResetPasswordMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ResetPasswordMongoDBInstance(request *ResetPasswordMongoDBInstanceRequest) (string) {
    return c.ResetPasswordMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) ResetPasswordMongoDBInstanceWithContext(ctx context.Context, request *ResetPasswordMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewResetPasswordMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewResetPasswordMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestartMongoDBInstanceRequest() (request *RestartMongoDBInstanceRequest) {
    request = &RestartMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "RestartMongoDBInstance")
    return
}

func NewRestartMongoDBInstanceResponse() (response *RestartMongoDBInstanceResponse) {
    response = &RestartMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestartMongoDBInstance(request *RestartMongoDBInstanceRequest) (string) {
    return c.RestartMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RestartMongoDBInstanceWithContext(ctx context.Context, request *RestartMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewRestartMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRestartMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateMongoDBSnapshotRequest() (request *CreateMongoDBSnapshotRequest) {
    request = &CreateMongoDBSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateMongoDBSnapshot")
    return
}

func NewCreateMongoDBSnapshotResponse() (response *CreateMongoDBSnapshotResponse) {
    response = &CreateMongoDBSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateMongoDBSnapshot(request *CreateMongoDBSnapshotRequest) (string) {
    return c.CreateMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) CreateMongoDBSnapshotWithContext(ctx context.Context, request *CreateMongoDBSnapshotRequest) (string) {
    if request == nil {
        request = NewCreateMongoDBSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateMongoDBSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSetMongoDBTimingSnapshotRequest() (request *SetMongoDBTimingSnapshotRequest) {
    request = &SetMongoDBTimingSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "SetMongoDBTimingSnapshot")
    return
}

func NewSetMongoDBTimingSnapshotResponse() (response *SetMongoDBTimingSnapshotResponse) {
    response = &SetMongoDBTimingSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SetMongoDBTimingSnapshot(request *SetMongoDBTimingSnapshotRequest) (string) {
    return c.SetMongoDBTimingSnapshotWithContext(context.Background(), request)
}

func (c *Client) SetMongoDBTimingSnapshotWithContext(ctx context.Context, request *SetMongoDBTimingSnapshotRequest) (string) {
    if request == nil {
        request = NewSetMongoDBTimingSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSetMongoDBTimingSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeMongoDBSnapshotRequest() (request *DescribeMongoDBSnapshotRequest) {
    request = &DescribeMongoDBSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBSnapshot")
    return
}

func NewDescribeMongoDBSnapshotResponse() (response *DescribeMongoDBSnapshotResponse) {
    response = &DescribeMongoDBSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeMongoDBSnapshot(request *DescribeMongoDBSnapshotRequest) (string) {
    return c.DescribeMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBSnapshotWithContext(ctx context.Context, request *DescribeMongoDBSnapshotRequest) (string) {
    if request == nil {
        request = NewDescribeMongoDBSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeMongoDBSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteMongoDBSnapshotRequest() (request *DeleteMongoDBSnapshotRequest) {
    request = &DeleteMongoDBSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteMongoDBSnapshot")
    return
}

func NewDeleteMongoDBSnapshotResponse() (response *DeleteMongoDBSnapshotResponse) {
    response = &DeleteMongoDBSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteMongoDBSnapshot(request *DeleteMongoDBSnapshotRequest) (string) {
    return c.DeleteMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) DeleteMongoDBSnapshotWithContext(ctx context.Context, request *DeleteMongoDBSnapshotRequest) (string) {
    if request == nil {
        request = NewDeleteMongoDBSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteMongoDBSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRenameMongoDBSnapshotRequest() (request *RenameMongoDBSnapshotRequest) {
    request = &RenameMongoDBSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "RenameMongoDBSnapshot")
    return
}

func NewRenameMongoDBSnapshotResponse() (response *RenameMongoDBSnapshotResponse) {
    response = &RenameMongoDBSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenameMongoDBSnapshot(request *RenameMongoDBSnapshotRequest) (string) {
    return c.RenameMongoDBSnapshotWithContext(context.Background(), request)
}

func (c *Client) RenameMongoDBSnapshotWithContext(ctx context.Context, request *RenameMongoDBSnapshotRequest) (string) {
    if request == nil {
        request = NewRenameMongoDBSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRenameMongoDBSnapshotResponse()
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
    request.Init().WithApiInfo("mongodb", APIVersion, "AddSecurityGroupRule")
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
func NewListSecurityGroupRulesRequest() (request *ListSecurityGroupRulesRequest) {
    request = &ListSecurityGroupRulesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "ListSecurityGroupRules")
    return
}

func NewListSecurityGroupRulesResponse() (response *ListSecurityGroupRulesResponse) {
    response = &ListSecurityGroupRulesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListSecurityGroupRules(request *ListSecurityGroupRulesRequest) (string) {
    return c.ListSecurityGroupRulesWithContext(context.Background(), request)
}

func (c *Client) ListSecurityGroupRulesWithContext(ctx context.Context, request *ListSecurityGroupRulesRequest) (string) {
    if request == nil {
        request = NewListSecurityGroupRulesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListSecurityGroupRulesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateMongoDBInstanceRequest() (request *UpdateMongoDBInstanceRequest) {
    request = &UpdateMongoDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "UpdateMongoDBInstance")
    return
}

func NewUpdateMongoDBInstanceResponse() (response *UpdateMongoDBInstanceResponse) {
    response = &UpdateMongoDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateMongoDBInstance(request *UpdateMongoDBInstanceRequest) (string) {
    return c.UpdateMongoDBInstanceWithContext(context.Background(), request)
}

func (c *Client) UpdateMongoDBInstanceWithContext(ctx context.Context, request *UpdateMongoDBInstanceRequest) (string) {
    if request == nil {
        request = NewUpdateMongoDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewUpdateMongoDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddSecondaryInstanceRequest() (request *AddSecondaryInstanceRequest) {
    request = &AddSecondaryInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "AddSecondaryInstance")
    return
}

func NewAddSecondaryInstanceResponse() (response *AddSecondaryInstanceResponse) {
    response = &AddSecondaryInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddSecondaryInstance(request *AddSecondaryInstanceRequest) (string) {
    return c.AddSecondaryInstanceWithContext(context.Background(), request)
}

func (c *Client) AddSecondaryInstanceWithContext(ctx context.Context, request *AddSecondaryInstanceRequest) (string) {
    if request == nil {
        request = NewAddSecondaryInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewAddSecondaryInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeMongoDBShardNodeRequest() (request *DescribeMongoDBShardNodeRequest) {
    request = &DescribeMongoDBShardNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongoDBShardNode")
    return
}

func NewDescribeMongoDBShardNodeResponse() (response *DescribeMongoDBShardNodeResponse) {
    response = &DescribeMongoDBShardNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeMongoDBShardNode(request *DescribeMongoDBShardNodeRequest) (string) {
    return c.DescribeMongoDBShardNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeMongoDBShardNodeWithContext(ctx context.Context, request *DescribeMongoDBShardNodeRequest) (string) {
    if request == nil {
        request = NewDescribeMongoDBShardNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeMongoDBShardNodeResponse()
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
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeValidRegion")
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
func NewAllocateEipRequest() (request *AllocateEipRequest) {
    request = &AllocateEipRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "AllocateEip")
    return
}

func NewAllocateEipResponse() (response *AllocateEipResponse) {
    response = &AllocateEipResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AllocateEip(request *AllocateEipRequest) (string) {
    return c.AllocateEipWithContext(context.Background(), request)
}

func (c *Client) AllocateEipWithContext(ctx context.Context, request *AllocateEipRequest) (string) {
    if request == nil {
        request = NewAllocateEipRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAllocateEipResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeallocateEipRequest() (request *DeallocateEipRequest) {
    request = &DeallocateEipRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DeallocateEip")
    return
}

func NewDeallocateEipResponse() (response *DeallocateEipResponse) {
    response = &DeallocateEipResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeallocateEip(request *DeallocateEipRequest) (string) {
    return c.DeallocateEipWithContext(context.Background(), request)
}

func (c *Client) DeallocateEipWithContext(ctx context.Context, request *DeallocateEipRequest) (string) {
    if request == nil {
        request = NewDeallocateEipRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeallocateEipResponse()
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
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeRegions")
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
func NewCreateMongoDBShardInstanceRequest() (request *CreateMongoDBShardInstanceRequest) {
    request = &CreateMongoDBShardInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateMongoDBShardInstance")
    return
}

func NewCreateMongoDBShardInstanceResponse() (response *CreateMongoDBShardInstanceResponse) {
    response = &CreateMongoDBShardInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateMongoDBShardInstance(request *CreateMongoDBShardInstanceRequest) (string) {
    return c.CreateMongoDBShardInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateMongoDBShardInstanceWithContext(ctx context.Context, request *CreateMongoDBShardInstanceRequest) (string) {
    if request == nil {
        request = NewCreateMongoDBShardInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateMongoDBShardInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDownloadSnapshotRequest() (request *DownloadSnapshotRequest) {
    request = &DownloadSnapshotRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DownloadSnapshot")
    return
}

func NewDownloadSnapshotResponse() (response *DownloadSnapshotResponse) {
    response = &DownloadSnapshotResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DownloadSnapshot(request *DownloadSnapshotRequest) (string) {
    return c.DownloadSnapshotWithContext(context.Background(), request)
}

func (c *Client) DownloadSnapshotWithContext(ctx context.Context, request *DownloadSnapshotRequest) (string) {
    if request == nil {
        request = NewDownloadSnapshotRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDownloadSnapshotResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCloneInstanceRequest() (request *CloneInstanceRequest) {
    request = &CloneInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "CloneInstance")
    return
}

func NewCloneInstanceResponse() (response *CloneInstanceResponse) {
    response = &CloneInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CloneInstance(request *CloneInstanceRequest) (string) {
    return c.CloneInstanceWithContext(context.Background(), request)
}

func (c *Client) CloneInstanceWithContext(ctx context.Context, request *CloneInstanceRequest) (string) {
    if request == nil {
        request = NewCloneInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCloneInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeShardNodeRequest() (request *DescribeShardNodeRequest) {
    request = &DescribeShardNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeShardNode")
    return
}

func NewDescribeShardNodeResponse() (response *DescribeShardNodeResponse) {
    response = &DescribeShardNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeShardNode(request *DescribeShardNodeRequest) (string) {
    return c.DescribeShardNodeWithContext(context.Background(), request)
}

func (c *Client) DescribeShardNodeWithContext(ctx context.Context, request *DescribeShardNodeRequest) (string) {
    if request == nil {
        request = NewDescribeShardNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeShardNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstanceStatisticRequest() (request *DescribeInstanceStatisticRequest) {
    request = &DescribeInstanceStatisticRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceStatistic")
    return
}

func NewDescribeInstanceStatisticResponse() (response *DescribeInstanceStatisticResponse) {
    response = &DescribeInstanceStatisticResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceStatistic(request *DescribeInstanceStatisticRequest) (string) {
    return c.DescribeInstanceStatisticWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceStatisticWithContext(ctx context.Context, request *DescribeInstanceStatisticRequest) (string) {
    if request == nil {
        request = NewDescribeInstanceStatisticRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeInstanceStatisticResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAddClusterNodeRequest() (request *AddClusterNodeRequest) {
    request = &AddClusterNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "AddClusterNode")
    return
}

func NewAddClusterNodeResponse() (response *AddClusterNodeResponse) {
    response = &AddClusterNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AddClusterNode(request *AddClusterNodeRequest) (string) {
    return c.AddClusterNodeWithContext(context.Background(), request)
}

func (c *Client) AddClusterNodeWithContext(ctx context.Context, request *AddClusterNodeRequest) (string) {
    if request == nil {
        request = NewAddClusterNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewAddClusterNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteClusterNodeRequest() (request *DeleteClusterNodeRequest) {
    request = &DeleteClusterNodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteClusterNode")
    return
}

func NewDeleteClusterNodeResponse() (response *DeleteClusterNodeResponse) {
    response = &DeleteClusterNodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteClusterNode(request *DeleteClusterNodeRequest) (string) {
    return c.DeleteClusterNodeWithContext(context.Background(), request)
}

func (c *Client) DeleteClusterNodeWithContext(ctx context.Context, request *DeleteClusterNodeRequest) (string) {
    if request == nil {
        request = NewDeleteClusterNodeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteClusterNodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSlowLogDetailRequest() (request *DescribeSlowLogDetailRequest) {
    request = &DescribeSlowLogDetailRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogDetail")
    return
}

func NewDescribeSlowLogDetailResponse() (response *DescribeSlowLogDetailResponse) {
    response = &DescribeSlowLogDetailResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSlowLogDetail(request *DescribeSlowLogDetailRequest) (string) {
    return c.DescribeSlowLogDetailWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogDetailWithContext(ctx context.Context, request *DescribeSlowLogDetailRequest) (string) {
    if request == nil {
        request = NewDescribeSlowLogDetailRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSlowLogDetailResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSlowLogStatisticsRequest() (request *DescribeSlowLogStatisticsRequest) {
    request = &DescribeSlowLogStatisticsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogStatistics")
    return
}

func NewDescribeSlowLogStatisticsResponse() (response *DescribeSlowLogStatisticsResponse) {
    response = &DescribeSlowLogStatisticsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSlowLogStatistics(request *DescribeSlowLogStatisticsRequest) (string) {
    return c.DescribeSlowLogStatisticsWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogStatisticsWithContext(ctx context.Context, request *DescribeSlowLogStatisticsRequest) (string) {
    if request == nil {
        request = NewDescribeSlowLogStatisticsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSlowLogStatisticsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSlowLogDatabaseRequest() (request *DescribeSlowLogDatabaseRequest) {
    request = &DescribeSlowLogDatabaseRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogDatabase")
    return
}

func NewDescribeSlowLogDatabaseResponse() (response *DescribeSlowLogDatabaseResponse) {
    response = &DescribeSlowLogDatabaseResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSlowLogDatabase(request *DescribeSlowLogDatabaseRequest) (string) {
    return c.DescribeSlowLogDatabaseWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogDatabaseWithContext(ctx context.Context, request *DescribeSlowLogDatabaseRequest) (string) {
    if request == nil {
        request = NewDescribeSlowLogDatabaseRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSlowLogDatabaseResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSlowLogLineChartRequest() (request *DescribeSlowLogLineChartRequest) {
    request = &DescribeSlowLogLineChartRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogLineChart")
    return
}

func NewDescribeSlowLogLineChartResponse() (response *DescribeSlowLogLineChartResponse) {
    response = &DescribeSlowLogLineChartResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSlowLogLineChart(request *DescribeSlowLogLineChartRequest) (string) {
    return c.DescribeSlowLogLineChartWithContext(context.Background(), request)
}

func (c *Client) DescribeSlowLogLineChartWithContext(ctx context.Context, request *DescribeSlowLogLineChartRequest) (string) {
    if request == nil {
        request = NewDescribeSlowLogLineChartRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeSlowLogLineChartResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpdateMongoDBInstanceClusterRequest() (request *UpdateMongoDBInstanceClusterRequest) {
    request = &UpdateMongoDBInstanceClusterRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "UpdateMongoDBInstanceCluster")
    return
}

func NewUpdateMongoDBInstanceClusterResponse() (response *UpdateMongoDBInstanceClusterResponse) {
    response = &UpdateMongoDBInstanceClusterResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpdateMongoDBInstanceCluster(request *UpdateMongoDBInstanceClusterRequest) (string) {
    return c.UpdateMongoDBInstanceClusterWithContext(context.Background(), request)
}

func (c *Client) UpdateMongoDBInstanceClusterWithContext(ctx context.Context, request *UpdateMongoDBInstanceClusterRequest) (string) {
    if request == nil {
        request = NewUpdateMongoDBInstanceClusterRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewUpdateMongoDBInstanceClusterResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeClusterForRestoreRequest() (request *DescribeClusterForRestoreRequest) {
    request = &DescribeClusterForRestoreRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeClusterForRestore")
    return
}

func NewDescribeClusterForRestoreResponse() (response *DescribeClusterForRestoreResponse) {
    response = &DescribeClusterForRestoreResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeClusterForRestore(request *DescribeClusterForRestoreRequest) (string) {
    return c.DescribeClusterForRestoreWithContext(context.Background(), request)
}

func (c *Client) DescribeClusterForRestoreWithContext(ctx context.Context, request *DescribeClusterForRestoreRequest) (string) {
    if request == nil {
        request = NewDescribeClusterForRestoreRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeClusterForRestoreResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


