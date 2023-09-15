package v20160701
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2016-07-01"

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

func NewRebootDBInstanceRequest() (request *RebootDBInstanceRequest) {
    request = &RebootDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "RebootDBInstance")
    return
}

func NewRebootDBInstanceResponse() (response *RebootDBInstanceResponse) {
    response = &RebootDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RebootDBInstance(request *RebootDBInstanceRequest) (string) {
    return c.RebootDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RebootDBInstanceWithContext(ctx context.Context, request *RebootDBInstanceRequest) (string) {
    if request == nil {
        request = NewRebootDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRebootDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDBParameterGroupRequest() (request *ModifyDBParameterGroupRequest) {
    request = &ModifyDBParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyDBParameterGroup")
    return
}

func NewModifyDBParameterGroupResponse() (response *ModifyDBParameterGroupResponse) {
    response = &ModifyDBParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDBParameterGroup(request *ModifyDBParameterGroupRequest) (string) {
    return c.ModifyDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ModifyDBParameterGroupWithContext(ctx context.Context, request *ModifyDBParameterGroupRequest) (string) {
    if request == nil {
        request = NewModifyDBParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyDBParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewResetDBParameterGroupRequest() (request *ResetDBParameterGroupRequest) {
    request = &ResetDBParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ResetDBParameterGroup")
    return
}

func NewResetDBParameterGroupResponse() (response *ResetDBParameterGroupResponse) {
    response = &ResetDBParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ResetDBParameterGroup(request *ResetDBParameterGroupRequest) (string) {
    return c.ResetDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) ResetDBParameterGroupWithContext(ctx context.Context, request *ResetDBParameterGroupRequest) (string) {
    if request == nil {
        request = NewResetDBParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewResetDBParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBParameterGroupRequest() (request *DescribeDBParameterGroupRequest) {
    request = &DescribeDBParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBParameterGroup")
    return
}

func NewDescribeDBParameterGroupResponse() (response *DescribeDBParameterGroupResponse) {
    response = &DescribeDBParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBParameterGroup(request *DescribeDBParameterGroupRequest) (string) {
    return c.DescribeDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DescribeDBParameterGroupWithContext(ctx context.Context, request *DescribeDBParameterGroupRequest) (string) {
    if request == nil {
        request = NewDescribeDBParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEngineDefaultParametersRequest() (request *DescribeEngineDefaultParametersRequest) {
    request = &DescribeEngineDefaultParametersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineDefaultParameters")
    return
}

func NewDescribeEngineDefaultParametersResponse() (response *DescribeEngineDefaultParametersResponse) {
    response = &DescribeEngineDefaultParametersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEngineDefaultParameters(request *DescribeEngineDefaultParametersRequest) (string) {
    return c.DescribeEngineDefaultParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeEngineDefaultParametersWithContext(ctx context.Context, request *DescribeEngineDefaultParametersRequest) (string) {
    if request == nil {
        request = NewDescribeEngineDefaultParametersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeEngineDefaultParametersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateDBParameterGroupRequest() (request *CreateDBParameterGroupRequest) {
    request = &CreateDBParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateDBParameterGroup")
    return
}

func NewCreateDBParameterGroupResponse() (response *CreateDBParameterGroupResponse) {
    response = &CreateDBParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateDBParameterGroup(request *CreateDBParameterGroupRequest) (string) {
    return c.CreateDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) CreateDBParameterGroupWithContext(ctx context.Context, request *CreateDBParameterGroupRequest) (string) {
    if request == nil {
        request = NewCreateDBParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCreateDBParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteDBParameterGroupRequest() (request *DeleteDBParameterGroupRequest) {
    request = &DeleteDBParameterGroupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DeleteDBParameterGroup")
    return
}

func NewDeleteDBParameterGroupResponse() (response *DeleteDBParameterGroupResponse) {
    response = &DeleteDBParameterGroupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteDBParameterGroup(request *DeleteDBParameterGroupRequest) (string) {
    return c.DeleteDBParameterGroupWithContext(context.Background(), request)
}

func (c *Client) DeleteDBParameterGroupWithContext(ctx context.Context, request *DeleteDBParameterGroupRequest) (string) {
    if request == nil {
        request = NewDeleteDBParameterGroupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteDBParameterGroupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstance")
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (string) {
    return c.CreateDBInstanceWithContext(context.Background(), request)
}

func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) (string) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCreateDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestoreDBInstanceFromDBBackupRequest() (request *RestoreDBInstanceFromDBBackupRequest) {
    request = &RestoreDBInstanceFromDBBackupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceFromDBBackup")
    return
}

func NewRestoreDBInstanceFromDBBackupResponse() (response *RestoreDBInstanceFromDBBackupResponse) {
    response = &RestoreDBInstanceFromDBBackupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestoreDBInstanceFromDBBackup(request *RestoreDBInstanceFromDBBackupRequest) (string) {
    return c.RestoreDBInstanceFromDBBackupWithContext(context.Background(), request)
}

func (c *Client) RestoreDBInstanceFromDBBackupWithContext(ctx context.Context, request *RestoreDBInstanceFromDBBackupRequest) (string) {
    if request == nil {
        request = NewRestoreDBInstanceFromDBBackupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRestoreDBInstanceFromDBBackupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
    request = &DeleteDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DeleteDBInstance")
    return
}

func NewDeleteDBInstanceResponse() (response *DeleteDBInstanceResponse) {
    response = &DeleteDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (string) {
    return c.DeleteDBInstanceWithContext(context.Background(), request)
}

func (c *Client) DeleteDBInstanceWithContext(ctx context.Context, request *DeleteDBInstanceRequest) (string) {
    if request == nil {
        request = NewDeleteDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateDBInstanceReadReplicaRequest() (request *CreateDBInstanceReadReplicaRequest) {
    request = &CreateDBInstanceReadReplicaRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateDBInstanceReadReplica")
    return
}

func NewCreateDBInstanceReadReplicaResponse() (response *CreateDBInstanceReadReplicaResponse) {
    response = &CreateDBInstanceReadReplicaResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateDBInstanceReadReplica(request *CreateDBInstanceReadReplicaRequest) (string) {
    return c.CreateDBInstanceReadReplicaWithContext(context.Background(), request)
}

func (c *Client) CreateDBInstanceReadReplicaWithContext(ctx context.Context, request *CreateDBInstanceReadReplicaRequest) (string) {
    if request == nil {
        request = NewCreateDBInstanceReadReplicaRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCreateDBInstanceReadReplicaResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestoreDBInstanceToPointInTimeRequest() (request *RestoreDBInstanceToPointInTimeRequest) {
    request = &RestoreDBInstanceToPointInTimeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "RestoreDBInstanceToPointInTime")
    return
}

func NewRestoreDBInstanceToPointInTimeResponse() (response *RestoreDBInstanceToPointInTimeResponse) {
    response = &RestoreDBInstanceToPointInTimeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestoreDBInstanceToPointInTime(request *RestoreDBInstanceToPointInTimeRequest) (string) {
    return c.RestoreDBInstanceToPointInTimeWithContext(context.Background(), request)
}

func (c *Client) RestoreDBInstanceToPointInTimeWithContext(ctx context.Context, request *RestoreDBInstanceToPointInTimeRequest) (string) {
    if request == nil {
        request = NewRestoreDBInstanceToPointInTimeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRestoreDBInstanceToPointInTimeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBInstanceRestorableTimeRequest() (request *DescribeDBInstanceRestorableTimeRequest) {
    request = &DescribeDBInstanceRestorableTimeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceRestorableTime")
    return
}

func NewDescribeDBInstanceRestorableTimeResponse() (response *DescribeDBInstanceRestorableTimeResponse) {
    response = &DescribeDBInstanceRestorableTimeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBInstanceRestorableTime(request *DescribeDBInstanceRestorableTimeRequest) (string) {
    return c.DescribeDBInstanceRestorableTimeWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceRestorableTimeWithContext(ctx context.Context, request *DescribeDBInstanceRestorableTimeRequest) (string) {
    if request == nil {
        request = NewDescribeDBInstanceRestorableTimeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBInstanceRestorableTimeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDBInstanceRequest() (request *ModifyDBInstanceRequest) {
    request = &ModifyDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstance")
    return
}

func NewModifyDBInstanceResponse() (response *ModifyDBInstanceResponse) {
    response = &ModifyDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDBInstance(request *ModifyDBInstanceRequest) (string) {
    return c.ModifyDBInstanceWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceWithContext(ctx context.Context, request *ModifyDBInstanceRequest) (string) {
    if request == nil {
        request = NewModifyDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBLogFiles")
    return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
    response = &DescribeDBLogFilesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (string) {
    return c.DescribeDBLogFilesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBLogFilesWithContext(ctx context.Context, request *DescribeDBLogFilesRequest) (string) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBLogFilesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
    request = &DescribeDBBackupsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBBackups")
    return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
    response = &DescribeDBBackupsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) (string) {
    return c.DescribeDBBackupsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBBackupsWithContext(ctx context.Context, request *DescribeDBBackupsRequest) (string) {
    if request == nil {
        request = NewDescribeDBBackupsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBBackupsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
    request = &ModifyDBInstanceSpecRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceSpec")
    return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
    response = &ModifyDBInstanceSpecResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (string) {
    return c.ModifyDBInstanceSpecWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest) (string) {
    if request == nil {
        request = NewModifyDBInstanceSpecRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyDBInstanceSpecResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (string) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (string) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewOverrideDBInstanceRequest() (request *OverrideDBInstanceRequest) {
    request = &OverrideDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstance")
    return
}

func NewOverrideDBInstanceResponse() (response *OverrideDBInstanceResponse) {
    response = &OverrideDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) OverrideDBInstance(request *OverrideDBInstanceRequest) (string) {
    return c.OverrideDBInstanceWithContext(context.Background(), request)
}

func (c *Client) OverrideDBInstanceWithContext(ctx context.Context, request *OverrideDBInstanceRequest) (string) {
    if request == nil {
        request = NewOverrideDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewOverrideDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBEngineVersionsRequest() (request *DescribeDBEngineVersionsRequest) {
    request = &DescribeDBEngineVersionsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBEngineVersions")
    return
}

func NewDescribeDBEngineVersionsResponse() (response *DescribeDBEngineVersionsResponse) {
    response = &DescribeDBEngineVersionsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBEngineVersions(request *DescribeDBEngineVersionsRequest) (string) {
    return c.DescribeDBEngineVersionsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBEngineVersionsWithContext(ctx context.Context, request *DescribeDBEngineVersionsRequest) (string) {
    if request == nil {
        request = NewDescribeDBEngineVersionsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBEngineVersionsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewUpgradeDBInstanceEngineVersionRequest() (request *UpgradeDBInstanceEngineVersionRequest) {
    request = &UpgradeDBInstanceEngineVersionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "UpgradeDBInstanceEngineVersion")
    return
}

func NewUpgradeDBInstanceEngineVersionResponse() (response *UpgradeDBInstanceEngineVersionResponse) {
    response = &UpgradeDBInstanceEngineVersionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (string) {
    return c.UpgradeDBInstanceEngineVersionWithContext(context.Background(), request)
}

func (c *Client) UpgradeDBInstanceEngineVersionWithContext(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest) (string) {
    if request == nil {
        request = NewUpgradeDBInstanceEngineVersionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewUpgradeDBInstanceEngineVersionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDBInstanceTypeRequest() (request *ModifyDBInstanceTypeRequest) {
    request = &ModifyDBInstanceTypeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceType")
    return
}

func NewModifyDBInstanceTypeResponse() (response *ModifyDBInstanceTypeResponse) {
    response = &ModifyDBInstanceTypeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDBInstanceType(request *ModifyDBInstanceTypeRequest) (string) {
    return c.ModifyDBInstanceTypeWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceTypeWithContext(ctx context.Context, request *ModifyDBInstanceTypeRequest) (string) {
    if request == nil {
        request = NewModifyDBInstanceTypeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyDBInstanceTypeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBInstanceParametersRequest() (request *DescribeDBInstanceParametersRequest) {
    request = &DescribeDBInstanceParametersRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceParameters")
    return
}

func NewDescribeDBInstanceParametersResponse() (response *DescribeDBInstanceParametersResponse) {
    response = &DescribeDBInstanceParametersResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBInstanceParameters(request *DescribeDBInstanceParametersRequest) (string) {
    return c.DescribeDBInstanceParametersWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceParametersWithContext(ctx context.Context, request *DescribeDBInstanceParametersRequest) (string) {
    if request == nil {
        request = NewDescribeDBInstanceParametersRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBInstanceParametersResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteDBBackupRequest() (request *DeleteDBBackupRequest) {
    request = &DeleteDBBackupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DeleteDBBackup")
    return
}

func NewDeleteDBBackupResponse() (response *DeleteDBBackupResponse) {
    response = &DeleteDBBackupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteDBBackup(request *DeleteDBBackupRequest) (string) {
    return c.DeleteDBBackupWithContext(context.Background(), request)
}

func (c *Client) DeleteDBBackupWithContext(ctx context.Context, request *DeleteDBBackupRequest) (string) {
    if request == nil {
        request = NewDeleteDBBackupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDeleteDBBackupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateDBBackupRequest() (request *CreateDBBackupRequest) {
    request = &CreateDBBackupRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateDBBackup")
    return
}

func NewCreateDBBackupResponse() (response *CreateDBBackupResponse) {
    response = &CreateDBBackupResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateDBBackup(request *CreateDBBackupRequest) (string) {
    return c.CreateDBBackupWithContext(context.Background(), request)
}

func (c *Client) CreateDBBackupWithContext(ctx context.Context, request *CreateDBBackupRequest) (string) {
    if request == nil {
        request = NewCreateDBBackupRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCreateDBBackupResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRenewDBInstanceRequest() (request *RenewDBInstanceRequest) {
    request = &RenewDBInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "RenewDBInstance")
    return
}

func NewRenewDBInstanceResponse() (response *RenewDBInstanceResponse) {
    response = &RenewDBInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RenewDBInstance(request *RenewDBInstanceRequest) (string) {
    return c.RenewDBInstanceWithContext(context.Background(), request)
}

func (c *Client) RenewDBInstanceWithContext(ctx context.Context, request *RenewDBInstanceRequest) (string) {
    if request == nil {
        request = NewRenewDBInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRenewDBInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSwitchDBInstanceHARequest() (request *SwitchDBInstanceHARequest) {
    request = &SwitchDBInstanceHARequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SwitchDBInstanceHA")
    return
}

func NewSwitchDBInstanceHAResponse() (response *SwitchDBInstanceHAResponse) {
    response = &SwitchDBInstanceHAResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) (string) {
    return c.SwitchDBInstanceHAWithContext(context.Background(), request)
}

func (c *Client) SwitchDBInstanceHAWithContext(ctx context.Context, request *SwitchDBInstanceHARequest) (string) {
    if request == nil {
        request = NewSwitchDBInstanceHARequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSwitchDBInstanceHAResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGenerateDBAdminURLRequest() (request *GenerateDBAdminURLRequest) {
    request = &GenerateDBAdminURLRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "GenerateDBAdminURL")
    return
}

func NewGenerateDBAdminURLResponse() (response *GenerateDBAdminURLResponse) {
    response = &GenerateDBAdminURLResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GenerateDBAdminURL(request *GenerateDBAdminURLRequest) (string) {
    return c.GenerateDBAdminURLWithContext(context.Background(), request)
}

func (c *Client) GenerateDBAdminURLWithContext(ctx context.Context, request *GenerateDBAdminURLRequest) (string) {
    if request == nil {
        request = NewGenerateDBAdminURLRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewGenerateDBAdminURLResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStatisticDBInstancesRequest() (request *StatisticDBInstancesRequest) {
    request = &StatisticDBInstancesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "StatisticDBInstances")
    return
}

func NewStatisticDBInstancesResponse() (response *StatisticDBInstancesResponse) {
    response = &StatisticDBInstancesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StatisticDBInstances(request *StatisticDBInstancesRequest) (string) {
    return c.StatisticDBInstancesWithContext(context.Background(), request)
}

func (c *Client) StatisticDBInstancesWithContext(ctx context.Context, request *StatisticDBInstancesRequest) (string) {
    if request == nil {
        request = NewStatisticDBInstancesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStatisticDBInstancesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAllocateDBInstanceEipRequest() (request *AllocateDBInstanceEipRequest) {
    request = &AllocateDBInstanceEipRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "AllocateDBInstanceEip")
    return
}

func NewAllocateDBInstanceEipResponse() (response *AllocateDBInstanceEipResponse) {
    response = &AllocateDBInstanceEipResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AllocateDBInstanceEip(request *AllocateDBInstanceEipRequest) (string) {
    return c.AllocateDBInstanceEipWithContext(context.Background(), request)
}

func (c *Client) AllocateDBInstanceEipWithContext(ctx context.Context, request *AllocateDBInstanceEipRequest) (string) {
    if request == nil {
        request = NewAllocateDBInstanceEipRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewAllocateDBInstanceEipResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewReleaseDBInstanceEipRequest() (request *ReleaseDBInstanceEipRequest) {
    request = &ReleaseDBInstanceEipRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ReleaseDBInstanceEip")
    return
}

func NewReleaseDBInstanceEipResponse() (response *ReleaseDBInstanceEipResponse) {
    response = &ReleaseDBInstanceEipResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ReleaseDBInstanceEip(request *ReleaseDBInstanceEipRequest) (string) {
    return c.ReleaseDBInstanceEipWithContext(context.Background(), request)
}

func (c *Client) ReleaseDBInstanceEipWithContext(ctx context.Context, request *ReleaseDBInstanceEipRequest) (string) {
    if request == nil {
        request = NewReleaseDBInstanceEipRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewReleaseDBInstanceEipResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDBInstanceAvailabilityZoneRequest() (request *ModifyDBInstanceAvailabilityZoneRequest) {
    request = &ModifyDBInstanceAvailabilityZoneRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyDBInstanceAvailabilityZone")
    return
}

func NewModifyDBInstanceAvailabilityZoneResponse() (response *ModifyDBInstanceAvailabilityZoneResponse) {
    response = &ModifyDBInstanceAvailabilityZoneResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDBInstanceAvailabilityZone(request *ModifyDBInstanceAvailabilityZoneRequest) (string) {
    return c.ModifyDBInstanceAvailabilityZoneWithContext(context.Background(), request)
}

func (c *Client) ModifyDBInstanceAvailabilityZoneWithContext(ctx context.Context, request *ModifyDBInstanceAvailabilityZoneRequest) (string) {
    if request == nil {
        request = NewModifyDBInstanceAvailabilityZoneRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyDBInstanceAvailabilityZoneResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBInstanceRegionsRequest() (request *DescribeDBInstanceRegionsRequest) {
    request = &DescribeDBInstanceRegionsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstanceRegions")
    return
}

func NewDescribeDBInstanceRegionsResponse() (response *DescribeDBInstanceRegionsResponse) {
    response = &DescribeDBInstanceRegionsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBInstanceRegions(request *DescribeDBInstanceRegionsRequest) (string) {
    return c.DescribeDBInstanceRegionsWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstanceRegionsWithContext(ctx context.Context, request *DescribeDBInstanceRegionsRequest) (string) {
    if request == nil {
        request = NewDescribeDBInstanceRegionsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBInstanceRegionsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeDBInstancePackagesRequest() (request *DescribeDBInstancePackagesRequest) {
    request = &DescribeDBInstancePackagesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeDBInstancePackages")
    return
}

func NewDescribeDBInstancePackagesResponse() (response *DescribeDBInstancePackagesResponse) {
    response = &DescribeDBInstancePackagesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeDBInstancePackages(request *DescribeDBInstancePackagesRequest) (string) {
    return c.DescribeDBInstancePackagesWithContext(context.Background(), request)
}

func (c *Client) DescribeDBInstancePackagesWithContext(ctx context.Context, request *DescribeDBInstancePackagesRequest) (string) {
    if request == nil {
        request = NewDescribeDBInstancePackagesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeDBInstancePackagesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeLastLogRequest() (request *DescribeLastLogRequest) {
    request = &DescribeLastLogRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeLastLog")
    return
}

func NewDescribeLastLogResponse() (response *DescribeLastLogResponse) {
    response = &DescribeLastLogResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeLastLog(request *DescribeLastLogRequest) (string) {
    return c.DescribeLastLogWithContext(context.Background(), request)
}

func (c *Client) DescribeLastLogWithContext(ctx context.Context, request *DescribeLastLogRequest) (string) {
    if request == nil {
        request = NewDescribeLastLogRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeLastLogResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStartAuditRequest() (request *StartAuditRequest) {
    request = &StartAuditRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "StartAudit")
    return
}

func NewStartAuditResponse() (response *StartAuditResponse) {
    response = &StartAuditResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StartAudit(request *StartAuditRequest) (string) {
    return c.StartAuditWithContext(context.Background(), request)
}

func (c *Client) StartAuditWithContext(ctx context.Context, request *StartAuditRequest) (string) {
    if request == nil {
        request = NewStartAuditRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStartAuditResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStopAuditRequest() (request *StopAuditRequest) {
    request = &StopAuditRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "StopAudit")
    return
}

func NewStopAuditResponse() (response *StopAuditResponse) {
    response = &StopAuditResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StopAudit(request *StopAuditRequest) (string) {
    return c.StopAuditWithContext(context.Background(), request)
}

func (c *Client) StopAuditWithContext(ctx context.Context, request *StopAuditRequest) (string) {
    if request == nil {
        request = NewStopAuditRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStopAuditResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListAuditRequest() (request *ListAuditRequest) {
    request = &ListAuditRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ListAudit")
    return
}

func NewListAuditResponse() (response *ListAuditResponse) {
    response = &ListAuditResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAudit(request *ListAuditRequest) (string) {
    return c.ListAuditWithContext(context.Background(), request)
}

func (c *Client) ListAuditWithContext(ctx context.Context, request *ListAuditRequest) (string) {
    if request == nil {
        request = NewListAuditRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListAuditResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewAuditStatisticRequest() (request *AuditStatisticRequest) {
    request = &AuditStatisticRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "AuditStatistic")
    return
}

func NewAuditStatisticResponse() (response *AuditStatisticResponse) {
    response = &AuditStatisticResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) AuditStatistic(request *AuditStatisticRequest) (string) {
    return c.AuditStatisticWithContext(context.Background(), request)
}

func (c *Client) AuditStatisticWithContext(ctx context.Context, request *AuditStatisticRequest) (string) {
    if request == nil {
        request = NewAuditStatisticRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewAuditStatisticResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetTableRestorableTimeRequest() (request *GetTableRestorableTimeRequest) {
    request = &GetTableRestorableTimeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "GetTableRestorableTime")
    return
}

func NewGetTableRestorableTimeResponse() (response *GetTableRestorableTimeResponse) {
    response = &GetTableRestorableTimeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetTableRestorableTime(request *GetTableRestorableTimeRequest) (string) {
    return c.GetTableRestorableTimeWithContext(context.Background(), request)
}

func (c *Client) GetTableRestorableTimeWithContext(ctx context.Context, request *GetTableRestorableTimeRequest) (string) {
    if request == nil {
        request = NewGetTableRestorableTimeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewGetTableRestorableTimeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewGetHistoryDatabaseInfoRequest() (request *GetHistoryDatabaseInfoRequest) {
    request = &GetHistoryDatabaseInfoRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "GetHistoryDatabaseInfo")
    return
}

func NewGetHistoryDatabaseInfoResponse() (response *GetHistoryDatabaseInfoResponse) {
    response = &GetHistoryDatabaseInfoResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetHistoryDatabaseInfo(request *GetHistoryDatabaseInfoRequest) (string) {
    return c.GetHistoryDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) GetHistoryDatabaseInfoWithContext(ctx context.Context, request *GetHistoryDatabaseInfoRequest) (string) {
    if request == nil {
        request = NewGetHistoryDatabaseInfoRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewGetHistoryDatabaseInfoResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewOverrideDBInstanceByPointInTimeRequest() (request *OverrideDBInstanceByPointInTimeRequest) {
    request = &OverrideDBInstanceByPointInTimeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "OverrideDBInstanceByPointInTime")
    return
}

func NewOverrideDBInstanceByPointInTimeResponse() (response *OverrideDBInstanceByPointInTimeResponse) {
    response = &OverrideDBInstanceByPointInTimeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) OverrideDBInstanceByPointInTime(request *OverrideDBInstanceByPointInTimeRequest) (string) {
    return c.OverrideDBInstanceByPointInTimeWithContext(context.Background(), request)
}

func (c *Client) OverrideDBInstanceByPointInTimeWithContext(ctx context.Context, request *OverrideDBInstanceByPointInTimeRequest) (string) {
    if request == nil {
        request = NewOverrideDBInstanceByPointInTimeRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewOverrideDBInstanceByPointInTimeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestoreToCurInstanceRequest() (request *RestoreToCurInstanceRequest) {
    request = &RestoreToCurInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "RestoreToCurInstance")
    return
}

func NewRestoreToCurInstanceResponse() (response *RestoreToCurInstanceResponse) {
    response = &RestoreToCurInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestoreToCurInstance(request *RestoreToCurInstanceRequest) (string) {
    return c.RestoreToCurInstanceWithContext(context.Background(), request)
}

func (c *Client) RestoreToCurInstanceWithContext(ctx context.Context, request *RestoreToCurInstanceRequest) (string) {
    if request == nil {
        request = NewRestoreToCurInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRestoreToCurInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewRestoreToSgInstanceRequest() (request *RestoreToSgInstanceRequest) {
    request = &RestoreToSgInstanceRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "RestoreToSgInstance")
    return
}

func NewRestoreToSgInstanceResponse() (response *RestoreToSgInstanceResponse) {
    response = &RestoreToSgInstanceResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) RestoreToSgInstance(request *RestoreToSgInstanceRequest) (string) {
    return c.RestoreToSgInstanceWithContext(context.Background(), request)
}

func (c *Client) RestoreToSgInstanceWithContext(ctx context.Context, request *RestoreToSgInstanceRequest) (string) {
    if request == nil {
        request = NewRestoreToSgInstanceRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewRestoreToSgInstanceResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeAuditHotCountRequest() (request *DescribeAuditHotCountRequest) {
    request = &DescribeAuditHotCountRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotCount")
    return
}

func NewDescribeAuditHotCountResponse() (response *DescribeAuditHotCountResponse) {
    response = &DescribeAuditHotCountResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeAuditHotCount(request *DescribeAuditHotCountRequest) (string) {
    return c.DescribeAuditHotCountWithContext(context.Background(), request)
}

func (c *Client) DescribeAuditHotCountWithContext(ctx context.Context, request *DescribeAuditHotCountRequest) (string) {
    if request == nil {
        request = NewDescribeAuditHotCountRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeAuditHotCountResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeAuditHotDurationRequest() (request *DescribeAuditHotDurationRequest) {
    request = &DescribeAuditHotDurationRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeAuditHotDuration")
    return
}

func NewDescribeAuditHotDurationResponse() (response *DescribeAuditHotDurationResponse) {
    response = &DescribeAuditHotDurationResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeAuditHotDuration(request *DescribeAuditHotDurationRequest) (string) {
    return c.DescribeAuditHotDurationWithContext(context.Background(), request)
}

func (c *Client) DescribeAuditHotDurationWithContext(ctx context.Context, request *DescribeAuditHotDurationRequest) (string) {
    if request == nil {
        request = NewDescribeAuditHotDurationRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeAuditHotDurationResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSqlAuditReportRequest() (request *SqlAuditReportRequest) {
    request = &SqlAuditReportRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SqlAuditReport")
    return
}

func NewSqlAuditReportResponse() (response *SqlAuditReportResponse) {
    response = &SqlAuditReportResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SqlAuditReport(request *SqlAuditReportRequest) (string) {
    return c.SqlAuditReportWithContext(context.Background(), request)
}

func (c *Client) SqlAuditReportWithContext(ctx context.Context, request *SqlAuditReportRequest) (string) {
    if request == nil {
        request = NewSqlAuditReportRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSqlAuditReportResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSqlAuditLineChartRequest() (request *SqlAuditLineChartRequest) {
    request = &SqlAuditLineChartRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SqlAuditLineChart")
    return
}

func NewSqlAuditLineChartResponse() (response *SqlAuditLineChartResponse) {
    response = &SqlAuditLineChartResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SqlAuditLineChart(request *SqlAuditLineChartRequest) (string) {
    return c.SqlAuditLineChartWithContext(context.Background(), request)
}

func (c *Client) SqlAuditLineChartWithContext(ctx context.Context, request *SqlAuditLineChartRequest) (string) {
    if request == nil {
        request = NewSqlAuditLineChartRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSqlAuditLineChartResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSlowLogReportRequest() (request *SlowLogReportRequest) {
    request = &SlowLogReportRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SlowLogReport")
    return
}

func NewSlowLogReportResponse() (response *SlowLogReportResponse) {
    response = &SlowLogReportResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SlowLogReport(request *SlowLogReportRequest) (string) {
    return c.SlowLogReportWithContext(context.Background(), request)
}

func (c *Client) SlowLogReportWithContext(ctx context.Context, request *SlowLogReportRequest) (string) {
    if request == nil {
        request = NewSlowLogReportRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSlowLogReportResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSlowLogLineChartRequest() (request *SlowLogLineChartRequest) {
    request = &SlowLogLineChartRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SlowLogLineChart")
    return
}

func NewSlowLogLineChartResponse() (response *SlowLogLineChartResponse) {
    response = &SlowLogLineChartResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SlowLogLineChart(request *SlowLogLineChartRequest) (string) {
    return c.SlowLogLineChartWithContext(context.Background(), request)
}

func (c *Client) SlowLogLineChartWithContext(ctx context.Context, request *SlowLogLineChartRequest) (string) {
    if request == nil {
        request = NewSlowLogLineChartRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSlowLogLineChartResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewSlowLogDetailRequest() (request *SlowLogDetailRequest) {
    request = &SlowLogDetailRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "SlowLogDetail")
    return
}

func NewSlowLogDetailResponse() (response *SlowLogDetailResponse) {
    response = &SlowLogDetailResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) SlowLogDetail(request *SlowLogDetailRequest) (string) {
    return c.SlowLogDetailWithContext(context.Background(), request)
}

func (c *Client) SlowLogDetailWithContext(ctx context.Context, request *SlowLogDetailRequest) (string) {
    if request == nil {
        request = NewSlowLogDetailRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewSlowLogDetailResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStartAuditDetailExportTaskRequest() (request *StartAuditDetailExportTaskRequest) {
    request = &StartAuditDetailExportTaskRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "StartAuditDetailExportTask")
    return
}

func NewStartAuditDetailExportTaskResponse() (response *StartAuditDetailExportTaskResponse) {
    response = &StartAuditDetailExportTaskResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StartAuditDetailExportTask(request *StartAuditDetailExportTaskRequest) (string) {
    return c.StartAuditDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) StartAuditDetailExportTaskWithContext(ctx context.Context, request *StartAuditDetailExportTaskRequest) (string) {
    if request == nil {
        request = NewStartAuditDetailExportTaskRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStartAuditDetailExportTaskResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListAuditDetailExportTaskRequest() (request *ListAuditDetailExportTaskRequest) {
    request = &ListAuditDetailExportTaskRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ListAuditDetailExportTask")
    return
}

func NewListAuditDetailExportTaskResponse() (response *ListAuditDetailExportTaskResponse) {
    response = &ListAuditDetailExportTaskResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListAuditDetailExportTask(request *ListAuditDetailExportTaskRequest) (string) {
    return c.ListAuditDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) ListAuditDetailExportTaskWithContext(ctx context.Context, request *ListAuditDetailExportTaskRequest) (string) {
    if request == nil {
        request = NewListAuditDetailExportTaskRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListAuditDetailExportTaskResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstanceAccountsRequest() (request *DescribeInstanceAccountsRequest) {
    request = &DescribeInstanceAccountsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceAccounts")
    return
}

func NewDescribeInstanceAccountsResponse() (response *DescribeInstanceAccountsResponse) {
    response = &DescribeInstanceAccountsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceAccounts(request *DescribeInstanceAccountsRequest) (string) {
    return c.DescribeInstanceAccountsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceAccountsWithContext(ctx context.Context, request *DescribeInstanceAccountsRequest) (string) {
    if request == nil {
        request = NewDescribeInstanceAccountsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeInstanceAccountsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyInstanceAccountInfoRequest() (request *ModifyInstanceAccountInfoRequest) {
    request = &ModifyInstanceAccountInfoRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountInfo")
    return
}

func NewModifyInstanceAccountInfoResponse() (response *ModifyInstanceAccountInfoResponse) {
    response = &ModifyInstanceAccountInfoResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyInstanceAccountInfo(request *ModifyInstanceAccountInfoRequest) (string) {
    return c.ModifyInstanceAccountInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountInfoWithContext(ctx context.Context, request *ModifyInstanceAccountInfoRequest) (string) {
    if request == nil {
        request = NewModifyInstanceAccountInfoRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyInstanceAccountInfoResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeCollationsRequest() (request *DescribeCollationsRequest) {
    request = &DescribeCollationsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeCollations")
    return
}

func NewDescribeCollationsResponse() (response *DescribeCollationsResponse) {
    response = &DescribeCollationsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeCollations(request *DescribeCollationsRequest) (string) {
    return c.DescribeCollationsWithContext(context.Background(), request)
}

func (c *Client) DescribeCollationsWithContext(ctx context.Context, request *DescribeCollationsRequest) (string) {
    if request == nil {
        request = NewDescribeCollationsRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeCollationsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateInstanceDatabaseRequest() (request *CreateInstanceDatabaseRequest) {
    request = &CreateInstanceDatabaseRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceDatabase")
    return
}

func NewCreateInstanceDatabaseResponse() (response *CreateInstanceDatabaseResponse) {
    response = &CreateInstanceDatabaseResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateInstanceDatabase(request *CreateInstanceDatabaseRequest) (string) {
    return c.CreateInstanceDatabaseWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceDatabaseWithContext(ctx context.Context, request *CreateInstanceDatabaseRequest) (string) {
    if request == nil {
        request = NewCreateInstanceDatabaseRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewCreateInstanceDatabaseResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyInstanceDatabasePrivilegesRequest() (request *ModifyInstanceDatabasePrivilegesRequest) {
    request = &ModifyInstanceDatabasePrivilegesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabasePrivileges")
    return
}

func NewModifyInstanceDatabasePrivilegesResponse() (response *ModifyInstanceDatabasePrivilegesResponse) {
    response = &ModifyInstanceDatabasePrivilegesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyInstanceDatabasePrivileges(request *ModifyInstanceDatabasePrivilegesRequest) (string) {
    return c.ModifyInstanceDatabasePrivilegesWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabasePrivilegesWithContext(ctx context.Context, request *ModifyInstanceDatabasePrivilegesRequest) (string) {
    if request == nil {
        request = NewModifyInstanceDatabasePrivilegesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyInstanceDatabasePrivilegesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstanceDatabasesRequest() (request *DescribeInstanceDatabasesRequest) {
    request = &DescribeInstanceDatabasesRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeInstanceDatabases")
    return
}

func NewDescribeInstanceDatabasesResponse() (response *DescribeInstanceDatabasesResponse) {
    response = &DescribeInstanceDatabasesResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceDatabases(request *DescribeInstanceDatabasesRequest) (string) {
    return c.DescribeInstanceDatabasesWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceDatabasesWithContext(ctx context.Context, request *DescribeInstanceDatabasesRequest) (string) {
    if request == nil {
        request = NewDescribeInstanceDatabasesRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeInstanceDatabasesResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyInstanceDatabaseInfoRequest() (request *ModifyInstanceDatabaseInfoRequest) {
    request = &ModifyInstanceDatabaseInfoRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceDatabaseInfo")
    return
}

func NewModifyInstanceDatabaseInfoResponse() (response *ModifyInstanceDatabaseInfoResponse) {
    response = &ModifyInstanceDatabaseInfoResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyInstanceDatabaseInfo(request *ModifyInstanceDatabaseInfoRequest) (string) {
    return c.ModifyInstanceDatabaseInfoWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceDatabaseInfoWithContext(ctx context.Context, request *ModifyInstanceDatabaseInfoRequest) (string) {
    if request == nil {
        request = NewModifyInstanceDatabaseInfoRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewModifyInstanceDatabaseInfoResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewStartSlowLogDetailExportTaskRequest() (request *StartSlowLogDetailExportTaskRequest) {
    request = &StartSlowLogDetailExportTaskRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "StartSlowLogDetailExportTask")
    return
}

func NewStartSlowLogDetailExportTaskResponse() (response *StartSlowLogDetailExportTaskResponse) {
    response = &StartSlowLogDetailExportTaskResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) StartSlowLogDetailExportTask(request *StartSlowLogDetailExportTaskRequest) (string) {
    return c.StartSlowLogDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) StartSlowLogDetailExportTaskWithContext(ctx context.Context, request *StartSlowLogDetailExportTaskRequest) (string) {
    if request == nil {
        request = NewStartSlowLogDetailExportTaskRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewStartSlowLogDetailExportTaskResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewListSlowLogDetailExportTaskRequest() (request *ListSlowLogDetailExportTaskRequest) {
    request = &ListSlowLogDetailExportTaskRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ListSlowLogDetailExportTask")
    return
}

func NewListSlowLogDetailExportTaskResponse() (response *ListSlowLogDetailExportTaskResponse) {
    response = &ListSlowLogDetailExportTaskResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ListSlowLogDetailExportTask(request *ListSlowLogDetailExportTaskRequest) (string) {
    return c.ListSlowLogDetailExportTaskWithContext(context.Background(), request)
}

func (c *Client) ListSlowLogDetailExportTaskWithContext(ctx context.Context, request *ListSlowLogDetailExportTaskRequest) (string) {
    if request == nil {
        request = NewListSlowLogDetailExportTaskRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewListSlowLogDetailExportTaskResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewCreateInstanceAccountActionRequest() (request *CreateInstanceAccountActionRequest) {
    request = &CreateInstanceAccountActionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "CreateInstanceAccountAction")
    return
}

func NewCreateInstanceAccountActionResponse() (response *CreateInstanceAccountActionResponse) {
    response = &CreateInstanceAccountActionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) CreateInstanceAccountAction(request *CreateInstanceAccountActionRequest) (string) {
    return c.CreateInstanceAccountActionWithContext(context.Background(), request)
}

func (c *Client) CreateInstanceAccountActionWithContext(ctx context.Context, request *CreateInstanceAccountActionRequest) (string) {
    if request == nil {
        request = NewCreateInstanceAccountActionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewCreateInstanceAccountActionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyInstanceAccountPrivilegesActionRequest() (request *ModifyInstanceAccountPrivilegesActionRequest) {
    request = &ModifyInstanceAccountPrivilegesActionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyInstanceAccountPrivilegesAction")
    return
}

func NewModifyInstanceAccountPrivilegesActionResponse() (response *ModifyInstanceAccountPrivilegesActionResponse) {
    response = &ModifyInstanceAccountPrivilegesActionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyInstanceAccountPrivilegesAction(request *ModifyInstanceAccountPrivilegesActionRequest) (string) {
    return c.ModifyInstanceAccountPrivilegesActionWithContext(context.Background(), request)
}

func (c *Client) ModifyInstanceAccountPrivilegesActionWithContext(ctx context.Context, request *ModifyInstanceAccountPrivilegesActionRequest) (string) {
    if request == nil {
        request = NewModifyInstanceAccountPrivilegesActionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyInstanceAccountPrivilegesActionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteInstanceAccountActionRequest() (request *DeleteInstanceAccountActionRequest) {
    request = &DeleteInstanceAccountActionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceAccountAction")
    return
}

func NewDeleteInstanceAccountActionResponse() (response *DeleteInstanceAccountActionResponse) {
    response = &DeleteInstanceAccountActionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteInstanceAccountAction(request *DeleteInstanceAccountActionRequest) (string) {
    return c.DeleteInstanceAccountActionWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceAccountActionWithContext(ctx context.Context, request *DeleteInstanceAccountActionRequest) (string) {
    if request == nil {
        request = NewDeleteInstanceAccountActionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteInstanceAccountActionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDeleteInstanceDatabaseActionRequest() (request *DeleteInstanceDatabaseActionRequest) {
    request = &DeleteInstanceDatabaseActionRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DeleteInstanceDatabaseAction")
    return
}

func NewDeleteInstanceDatabaseActionResponse() (response *DeleteInstanceDatabaseActionResponse) {
    response = &DeleteInstanceDatabaseActionResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DeleteInstanceDatabaseAction(request *DeleteInstanceDatabaseActionRequest) (string) {
    return c.DeleteInstanceDatabaseActionWithContext(context.Background(), request)
}

func (c *Client) DeleteInstanceDatabaseActionWithContext(ctx context.Context, request *DeleteInstanceDatabaseActionRequest) (string) {
    if request == nil {
        request = NewDeleteInstanceDatabaseActionRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewDeleteInstanceDatabaseActionResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewModifyDBNetworkRequest() (request *ModifyDBNetworkRequest) {
    request = &ModifyDBNetworkRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "ModifyDBNetwork")
    return
}

func NewModifyDBNetworkResponse() (response *ModifyDBNetworkResponse) {
    response = &ModifyDBNetworkResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) ModifyDBNetwork(request *ModifyDBNetworkRequest) (string) {
    return c.ModifyDBNetworkWithContext(context.Background(), request)
}

func (c *Client) ModifyDBNetworkWithContext(ctx context.Context, request *ModifyDBNetworkRequest) (string) {
    if request == nil {
        request = NewModifyDBNetworkRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/x-www-form-urlencoded")

    response := NewModifyDBNetworkResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeEngineParametersModifyHistoryRequest() (request *DescribeEngineParametersModifyHistoryRequest) {
    request = &DescribeEngineParametersModifyHistoryRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("krds", APIVersion, "DescribeEngineParametersModifyHistory")
    return
}

func NewDescribeEngineParametersModifyHistoryResponse() (response *DescribeEngineParametersModifyHistoryResponse) {
    response = &DescribeEngineParametersModifyHistoryResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeEngineParametersModifyHistory(request *DescribeEngineParametersModifyHistoryRequest) (string) {
    return c.DescribeEngineParametersModifyHistoryWithContext(context.Background(), request)
}

func (c *Client) DescribeEngineParametersModifyHistoryWithContext(ctx context.Context, request *DescribeEngineParametersModifyHistoryRequest) (string) {
    if request == nil {
        request = NewDescribeEngineParametersModifyHistoryRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewDescribeEngineParametersModifyHistoryResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


