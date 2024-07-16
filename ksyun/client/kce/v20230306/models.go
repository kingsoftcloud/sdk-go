package v20230306
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreatePrometheusInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    DataRetentionTime *int `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`
}

func (r *CreatePrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePrometheusInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePrometheusInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreatePrometheusInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *CreatePrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePrometheusInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePrometheusInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribePrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePrometheusInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribePrometheusInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	InstanceSet []struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		InstanceStatus *string `json:"InstanceStatus"`
		CreateTime *string `json:"CreateTime"`
		ChargeType *string `json:"ChargeType"`
		DataRetentionTime *int `json:"DataRetentionTime"`
		AuthToken *string `json:"AuthToken"`
		RemoteWriteUrl *string `json:"RemoteWriteUrl"`
		HttpApiUrl *string `json:"HttpApiUrl"`
		GrafanaEnabled *bool `json:"GrafanaEnabled"`
		GrafanaInternalUrl *string `json:"GrafanaInternalUrl"`
		GrafanaInternetEnabled *bool `json:"GrafanaInternetEnabled"`
		GrafanaInternetUrl *string `json:"GrafanaInternetUrl"`
		GrafanaWhiteList []struct {
			} `json:"GrafanaWhiteList"`
		} `json:"InstanceSet"`
}

func (r *DescribePrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePrometheusInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePrometheusInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
    DataRetentionTime *int `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`
}

func (r *UpdatePrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePrometheusInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdatePrometheusInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePrometheusInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdatePrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePrometheusInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrometheusInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeletePrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrometheusInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePrometheusInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrometheusInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeletePrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrometheusInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableGrafanaRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    EnableGrafana *bool `json:"EnableGrafana,omitempty" name:"EnableGrafana"`
    Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *EnableGrafanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableGrafanaRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EnableGrafanaRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type EnableGrafanaResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EnableGrafanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableGrafanaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGrafanaPasswordRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *UpdateGrafanaPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGrafanaPasswordRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateGrafanaPasswordRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGrafanaPasswordResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateGrafanaPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGrafanaPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableGrafanaInternetRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
}

func (r *EnableGrafanaInternetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableGrafanaInternetRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EnableGrafanaInternetRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type EnableGrafanaInternetResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EnableGrafanaInternetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableGrafanaInternetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGrafanaWhiteListRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGrafanaWhiteListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeGrafanaWhiteListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGrafanaWhiteListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    WhiteList []*string `json:"WhiteList" name:"WhiteList"`
}

func (r *DescribeGrafanaWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGrafanaWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGrafanaWhiteListRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
}

func (r *UpdateGrafanaWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGrafanaWhiteListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateGrafanaWhiteListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGrafanaWhiteListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateGrafanaWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGrafanaWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateClusterRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AssociateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateClusterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateClusterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AssociateClusterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AssociateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateClusterRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DisassociateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateClusterRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateClusterRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateClusterResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DisassociateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociateClusterListRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    ClusterId []*string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAssociateClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssociateClusterListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAssociateClusterListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociateClusterListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	ClusterSet []struct {
		ClusterId *string `json:"ClusterId"`
	} `json:"ClusterSet"`
}

func (r *DescribeAssociateClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssociateClusterListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorListRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorSource *string `json:"MonitorSource,omitempty" name:"MonitorSource"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeMonitorListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMonitorListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	MonitorSet []struct {
		MonitorName *string `json:"MonitorName"`
		Type *string `json:"Type"`
		TargetState *string `json:"TargetState"`
		MonitorSource *string `json:"MonitorSource"`
	} `json:"MonitorSet"`
}

func (r *DescribeMonitorListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorCollectionConfigRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeMonitorCollectionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorCollectionConfigRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMonitorCollectionConfigRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorCollectionConfigResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    ConfigYaml *string `json:"ConfigYaml" name:"ConfigYaml"`
}

func (r *DescribeMonitorCollectionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorCollectionConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateMonitorCollectionConfigRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    ConfigYaml *string `json:"ConfigYaml,omitempty" name:"ConfigYaml"`
}

func (r *UpdateMonitorCollectionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateMonitorCollectionConfigRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateMonitorCollectionConfigRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateMonitorCollectionConfigResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateMonitorCollectionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateMonitorCollectionConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorMetricsListRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    IsCollect *bool `json:"IsCollect,omitempty" name:"IsCollect"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeMonitorMetricsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorMetricsListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMonitorMetricsListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorMetricsListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	MetricsSet []struct {
		MetricName *string `json:"MetricName"`
		IsFree *bool `json:"IsFree"`
	} `json:"MetricsSet"`
}

func (r *DescribeMonitorMetricsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorMetricsListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsListRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeTargetsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetsListRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeTargetsListRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsListResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Marker *int `json:"Marker" name:"Marker"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
	TargetSet []struct {
		EndpointName *string `json:"EndpointName"`
		EndpointState *bool `json:"EndpointState"`
		Labels *string `json:"Labels"`
		LastScrapeTime *string `json:"LastScrapeTime"`
		LastScrapeDuration *string `json:"LastScrapeDuration"`
		ErrorMessage *string `json:"ErrorMessage"`
	} `json:"TargetSet"`
}

func (r *DescribeTargetsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetsListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentStatusRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAgentStatusRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    AgentsStatus *string `json:"AgentsStatus" name:"AgentsStatus"`
    Message *string `json:"Message" name:"Message"`
}

func (r *DescribeAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMonitorCollectionConfigRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    Type *string `json:"Type,omitempty" name:"Type"`
    ConfigYaml *string `json:"ConfigYaml,omitempty" name:"ConfigYaml"`
}

func (r *CreateMonitorCollectionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMonitorCollectionConfigRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateMonitorCollectionConfigRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateMonitorCollectionConfigResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateMonitorCollectionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMonitorCollectionConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMonitorCollectionConfigRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteMonitorCollectionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMonitorCollectionConfigRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteMonitorCollectionConfigRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMonitorCollectionConfigResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteMonitorCollectionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMonitorCollectionConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableMetricsRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    MetricsName []*string `json:"MetricsName,omitempty" name:"MetricsName"`
}

func (r *EnableMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableMetricsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "EnableMetricsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type EnableMetricsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *EnableMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropMetricsRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    MetricsName []*string `json:"MetricsName,omitempty" name:"MetricsName"`
}

func (r *DropMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropMetricsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DropMetricsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DropMetricsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DropMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

