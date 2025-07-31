package v20230306
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreatePrometheusInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceName      *string `json:"InstanceName,omitempty" name:"InstanceName"`
	ChargeType        *string `json:"ChargeType,omitempty" name:"ChargeType"`
	DataRetentionTime *int    `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`
}

func (r *CreatePrometheusInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreatePrometheusInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
	Marker     *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	Search     *string   `json:"Search,omitempty" name:"Search"`
}

func (r *DescribePrometheusInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribePrometheusInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	Marker      *int    `json:"Marker" name:"Marker"`
	TotalCount  *int    `json:"TotalCount" name:"TotalCount"`
	InstanceSet []struct {
		InstanceId             *string   `json:"InstanceId" name:"InstanceId"`
		InstanceName           *string   `json:"InstanceName" name:"InstanceName"`
		InstanceStatus         *string   `json:"InstanceStatus" name:"InstanceStatus"`
		CreateTime             *string   `json:"CreateTime" name:"CreateTime"`
		ChargeType             *string   `json:"ChargeType" name:"ChargeType"`
		DataRetentionTime      *int      `json:"DataRetentionTime" name:"DataRetentionTime"`
		AuthToken              *string   `json:"AuthToken" name:"AuthToken"`
		RemoteWriteUrl         *string   `json:"RemoteWriteUrl" name:"RemoteWriteUrl"`
		HttpApiUrl             *string   `json:"HttpApiUrl" name:"HttpApiUrl"`
		GrafanaEnabled         *bool     `json:"GrafanaEnabled" name:"GrafanaEnabled"`
		GrafanaInternalUrl     *string   `json:"GrafanaInternalUrl" name:"GrafanaInternalUrl"`
		GrafanaInternetEnabled *bool     `json:"GrafanaInternetEnabled" name:"GrafanaInternetEnabled"`
		GrafanaInternetUrl     *string   `json:"GrafanaInternetUrl" name:"GrafanaInternetUrl"`
		GrafanaWhiteList       []*string `json:"GrafanaWhiteList" name:"GrafanaWhiteList"`
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
	InstanceId        *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName      *string `json:"InstanceName,omitempty" name:"InstanceName"`
	DataRetentionTime *int    `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`
}

func (r *UpdatePrometheusInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	EnableGrafana *bool   `json:"EnableGrafana,omitempty" name:"EnableGrafana"`
	Password      *string `json:"Password,omitempty" name:"Password"`
}

func (r *EnableGrafanaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	Password   *string `json:"Password,omitempty" name:"Password"`
}

func (r *UpdateGrafanaPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId     *string `json:"InstanceId,omitempty" name:"InstanceId"`
	EnableInternet *bool   `json:"EnableInternet,omitempty" name:"EnableInternet"`
}

func (r *EnableGrafanaInternetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	WhiteList  []*string `json:"WhiteList,omitempty" name:"WhiteList"`
}

func (r *UpdateGrafanaWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	ClusterId  *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AssociateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	ClusterId  *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DisassociateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	Marker     *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	ClusterId  []*string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAssociateClusterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAssociateClusterListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	ClusterSet []struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
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
	InstanceId    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId     *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorSource *string `json:"MonitorSource,omitempty" name:"MonitorSource"`
	Marker        *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults    *int    `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeMonitorListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMonitorListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	MonitorSet []struct {
		MonitorName *string `json:"MonitorName" name:"MonitorName"`
		Type *string `json:"Type" name:"Type"`
		TargetState *string `json:"TargetState" name:"TargetState"`
		MonitorSource *string `json:"MonitorSource" name:"MonitorSource"`
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
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeMonitorCollectionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMonitorCollectionConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
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
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
	ConfigYaml  *string `json:"ConfigYaml,omitempty" name:"ConfigYaml"`
}

func (r *UpdateMonitorCollectionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
	IsCollect   *bool   `json:"IsCollect,omitempty" name:"IsCollect"`
	Marker      *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int    `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeMonitorMetricsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMonitorMetricsListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	MetricsSet []struct {
		MetricName *string `json:"MetricName" name:"MetricName"`
		IsFree *bool `json:"IsFree" name:"IsFree"`
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
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
	Marker      *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int    `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeTargetsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTargetsListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	TargetSet  []struct {
		EndpointName  *string `json:"EndpointName" name:"EndpointName"`
		EndpointState *bool   `json:"EndpointState" name:"EndpointState"`
		Labels        *string `json:"Labels" name:"Labels"`
		LastScrapeTime *string `json:"LastScrapeTime" name:"LastScrapeTime"`
		LastScrapeDuration *string `json:"LastScrapeDuration" name:"LastScrapeDuration"`
		ErrorMessage  *string `json:"ErrorMessage" name:"ErrorMessage"`
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
	ClusterId  *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAgentStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAgentStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	AgentsStatus *string `json:"AgentsStatus" name:"AgentsStatus"`
	Message      *string `json:"Message" name:"Message"`
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
	ClusterId  *string `json:"ClusterId,omitempty" name:"ClusterId"`
	Type       *string `json:"Type,omitempty" name:"Type"`
	ConfigYaml *string `json:"ConfigYaml,omitempty" name:"ConfigYaml"`
}

func (r *CreateMonitorCollectionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteMonitorCollectionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId  *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string   `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string   `json:"Type,omitempty" name:"Type"`
	MetricsName []*string `json:"MetricsName,omitempty" name:"MetricsName"`
}

func (r *EnableMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	InstanceId  *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	MonitorName *string   `json:"MonitorName,omitempty" name:"MonitorName"`
	Type        *string   `json:"Type,omitempty" name:"Type"`
	MetricsName []*string `json:"MetricsName,omitempty" name:"MetricsName"`
}

func (r *DropMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

