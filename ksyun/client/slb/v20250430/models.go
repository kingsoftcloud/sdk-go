package v20250430
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type DescribeBackendServersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeBackendServerGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeListenersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateListenerFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	HttpCode    *string `json:"HttpCode,omitempty" name:"HttpCode"`
}
type CreateListenerRewriteConfig struct {
	HttpHost    *string `json:"HttpHost,omitempty" name:"HttpHost"`
	Url         *string `json:"Url,omitempty" name:"Url"`
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}
type DescribeLoadBalancersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeLoadBalancersTagKV struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type DescribeListenerCertGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddRulesHeaderValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddRulesQueryValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddRulesCookieValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyRuleGroupRuleSetHeaderValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyRuleGroupRuleSetQueryValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyRuleGroupRuleSetCookieValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyRuleGroupRuleSet struct {
	RuleType      *string                              `json:"RuleType,omitempty" name:"RuleType"`
	RuleValue     *string                              `json:"RuleValue,omitempty" name:"RuleValue"`
	MethodValue   []*string                            `json:"MethodValue,omitempty" name:"MethodValue"`
	SourceIpValue []*string                            `json:"SourceIpValue,omitempty" name:"SourceIpValue"`
	HeaderValue   []*ModifyRuleGroupRuleSetHeaderValue `json:"HeaderValue,omitempty" name:"HeaderValue"`
	QueryValue    []*ModifyRuleGroupRuleSetQueryValue  `json:"QueryValue,omitempty" name:"QueryValue"`
	CookieValue   []*ModifyRuleGroupRuleSetCookieValue `json:"CookieValue,omitempty" name:"CookieValue"`
}
type ModifyRuleGroupFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	HttpCode    *string `json:"HttpCode,omitempty" name:"HttpCode"`
}
type ModifyRuleGroupRewriteConfig struct {
	HttpHost    *string `json:"HttpHost,omitempty" name:"HttpHost"`
	Url         *string `json:"Url,omitempty" name:"Url"`
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}
type DescribeRuleGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateRuleGroupRuleSetHeaderValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateRuleGroupRuleSetQueryValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateRuleGroupRuleSetCookieValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateRuleGroupRuleSet struct {
	RuleType      *string                              `json:"RuleType,omitempty" name:"RuleType"`
	RuleValue     *string                              `json:"RuleValue,omitempty" name:"RuleValue"`
	MethodValue   []*string                            `json:"MethodValue,omitempty" name:"MethodValue"`
	SourceIpValue []*string                            `json:"SourceIpValue,omitempty" name:"SourceIpValue"`
	HeaderValue   []*CreateRuleGroupRuleSetHeaderValue `json:"HeaderValue,omitempty" name:"HeaderValue"`
	QueryValue    []*CreateRuleGroupRuleSetQueryValue  `json:"QueryValue,omitempty" name:"QueryValue"`
	CookieValue   []*CreateRuleGroupRuleSetCookieValue `json:"CookieValue,omitempty" name:"CookieValue"`
}
type CreateRuleGroupFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	HttpCode    *string `json:"HttpCode,omitempty" name:"HttpCode"`
}
type CreateRuleGroupRewriteConfig struct {
	HttpHost    *string `json:"HttpHost,omitempty" name:"HttpHost"`
	Url         *string `json:"Url,omitempty" name:"Url"`
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}


type DescribeBackendServersRequest struct {
	*ksyunhttp.BaseRequest
	Filter          []*DescribeBackendServersFilter `json:"Filter,omitempty" name:"Filter"`
	BackendServerId []*string                       `json:"BackendServerId,omitempty" name:"BackendServerId"`
	MaxResults      *int                            `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken       *string                         `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeBackendServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBackendServersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	BackendServerSet []struct {
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		BackendServerIp *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		BackendServerId *string `json:"BackendServerId" name:"BackendServerId"`
		Port            *int    `json:"Port" name:"Port"`
		MasterSlaveType *string `json:"MasterSlaveType" name:"MasterSlaveType"`
		BackendServerState *string `json:"BackendServerState" name:"BackendServerState"`
	} `json:"BackendServerSet"`
}

func (r *DescribeBackendServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackendServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerId *string `json:"BackendServerId,omitempty" name:"BackendServerId"`
	Weight          *int    `json:"Weight,omitempty" name:"Weight"`
	Port            *int    `json:"Port,omitempty" name:"Port"`
	MasterSlaveType *string `json:"MasterSlaveType,omitempty" name:"MasterSlaveType"`
}

func (r *ModifyBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	BackendServer struct {
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		BackendServerIp *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		BackendServerId *string `json:"BackendServerId" name:"BackendServerId"`
		Port            *int    `json:"Port" name:"Port"`
		MasterSlaveType *string `json:"MasterSlaveType" name:"MasterSlaveType"`
		BackendServerState *string `json:"BackendServerState" name:"BackendServerState"`
	} `json:"BackendServer"`
}

func (r *ModifyBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeregisterBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerId *string `json:"BackendServerId,omitempty" name:"BackendServerId"`
}

func (r *DeregisterBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeregisterBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeregisterBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RegisterBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId   *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	BackendServerIp        *string `json:"BackendServerIp,omitempty" name:"BackendServerIp"`
	Port                   *int    `json:"Port,omitempty" name:"Port"`
	Weight                 *int    `json:"Weight,omitempty" name:"Weight"`
	NetworkInterfaceId     *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	MasterSlaveType        *string `json:"MasterSlaveType,omitempty" name:"MasterSlaveType"`
}

func (r *RegisterBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RegisterBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	BackendServer struct {
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		BackendServerIp *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		BackendServerId *string `json:"BackendServerId" name:"BackendServerId"`
		Port            *int    `json:"Port" name:"Port"`
		MasterSlaveType *string `json:"MasterSlaveType" name:"MasterSlaveType"`
		BackendServerState *string `json:"BackendServerState" name:"BackendServerState"`
	} `json:"BackendServer"`
}

func (r *RegisterBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeBackendServerGroupsRequest struct {
	*ksyunhttp.BaseRequest
	Filter               []*DescribeBackendServerGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	BackendServerGroupId []*string                            `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	MaxResults           *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken            *string                              `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeBackendServerGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBackendServerGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	BackendServerGroupSet []struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Name                *string `json:"Name" name:"Name"`
		BackendServerType   *string `json:"BackendServerType" name:"BackendServerType"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		VpcId               *string `json:"VpcId" name:"VpcId"`
		Protocol            *string `json:"Protocol" name:"Protocol"`
		BackendServerNumber *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive   *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		IpVersion           *string `json:"IpVersion" name:"IpVersion"`
		Method              *string `json:"Method" name:"Method"`
		SlowStartEnabled    *bool   `json:"SlowStartEnabled" name:"SlowStartEnabled"`
		SlowStartDuration   *int    `json:"SlowStartDuration" name:"SlowStartDuration"`
	} `json:"BackendServerGroupSet"`
}

func (r *DescribeBackendServerGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackendServerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyBackendServerGroupRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId     *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	Name                     *string `json:"Name,omitempty" name:"Name"`
	UpstreamKeepalive        *string `json:"UpstreamKeepalive,omitempty" name:"UpstreamKeepalive"`
	Method                   *string `json:"Method,omitempty" name:"Method"`
	SessionState             *string `json:"SessionState,omitempty" name:"SessionState"`
	SessionPersistencePeriod *int    `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
	CookieType               *string `json:"CookieType,omitempty" name:"CookieType"`
	CookieName               *string `json:"CookieName,omitempty" name:"CookieName"`
	HealthCheckState         *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	Timeout                  *int    `json:"Timeout,omitempty" name:"Timeout"`
	Interval                 *int    `json:"Interval,omitempty" name:"Interval"`
	HealthyThreshold         *int    `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	UnhealthyThreshold       *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	UrlPath                  *string `json:"UrlPath,omitempty" name:"UrlPath"`
	HostName                 *string `json:"HostName,omitempty" name:"HostName"`
	HealthCheckConnectPort   *int    `json:"HealthCheckConnectPort,omitempty" name:"HealthCheckConnectPort"`
	HealthProtocol           *string `json:"HealthProtocol,omitempty" name:"HealthProtocol"`
	SlowStartEnabled         *bool   `json:"SlowStartEnabled,omitempty" name:"SlowStartEnabled"`
	SlowStartDuration        *int    `json:"SlowStartDuration,omitempty" name:"SlowStartDuration"`
	HttpMethod               *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	HealthCode               *string `json:"HealthCode,omitempty" name:"HealthCode"`
	HealthCheckReq           *string `json:"HealthCheckReq,omitempty" name:"HealthCheckReq"`
	HealthCheckExp           *string `json:"HealthCheckExp,omitempty" name:"HealthCheckExp"`
}

func (r *ModifyBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BackendServerGroup struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Name                *string `json:"Name" name:"Name"`
		BackendServerType   *string `json:"BackendServerType" name:"BackendServerType"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		VpcId               *string `json:"VpcId" name:"VpcId"`
		Protocol            *string `json:"Protocol" name:"Protocol"`
		BackendServerNumber *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive   *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		IpVersion           *string `json:"IpVersion" name:"IpVersion"`
		Method              *string `json:"Method" name:"Method"`
		SlowStartEnabled    *bool   `json:"SlowStartEnabled" name:"SlowStartEnabled"`
		SlowStartDuration   *int    `json:"SlowStartDuration" name:"SlowStartDuration"`
	} `json:"BackendServerGroup"`
}

func (r *ModifyBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteBackendServerGroupRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
}

func (r *DeleteBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateBackendServerGroupRequest struct {
	*ksyunhttp.BaseRequest
	VpcId                    *string `json:"VpcId,omitempty" name:"VpcId"`
	Name                     *string `json:"Name,omitempty" name:"Name"`
	BackendServerType        *string `json:"BackendServerType,omitempty" name:"BackendServerType"`
	Method                   *string `json:"Method,omitempty" name:"Method"`
	SessionState             *string `json:"SessionState,omitempty" name:"SessionState"`
	SessionPersistencePeriod *int    `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
	CookieType               *string `json:"CookieType,omitempty" name:"CookieType"`
	CookieName               *string `json:"CookieName,omitempty" name:"CookieName"`
	UpstreamKeepalive        *string `json:"UpstreamKeepalive,omitempty" name:"UpstreamKeepalive"`
	Protocol                 *string `json:"Protocol,omitempty" name:"Protocol"`
	HealthCheckState         *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	Timeout                  *int    `json:"Timeout,omitempty" name:"Timeout"`
	Interval                 *int    `json:"Interval,omitempty" name:"Interval"`
	HealthyThreshold         *int    `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	UnhealthyThreshold       *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	UrlPath                  *string `json:"UrlPath,omitempty" name:"UrlPath"`
	HostName                 *string `json:"HostName,omitempty" name:"HostName"`
	HealthCheckConnectPort   *int    `json:"HealthCheckConnectPort,omitempty" name:"HealthCheckConnectPort"`
	HealthProtocol           *string `json:"HealthProtocol,omitempty" name:"HealthProtocol"`
	SlowStartEnabled         *bool   `json:"SlowStartEnabled,omitempty" name:"SlowStartEnabled"`
	SlowStartDuration        *int    `json:"SlowStartDuration,omitempty" name:"SlowStartDuration"`
	HttpMethod               *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	HealthCheckReq           *string `json:"HealthCheckReq,omitempty" name:"HealthCheckReq"`
	HealthCheckExp           *string `json:"HealthCheckExp,omitempty" name:"HealthCheckExp"`
	HealthCode               *string `json:"HealthCode,omitempty" name:"HealthCode"`
}

func (r *CreateBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BackendServerGroup struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Name                *string `json:"Name" name:"Name"`
		BackendServerType   *string `json:"BackendServerType" name:"BackendServerType"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		VpcId               *string `json:"VpcId" name:"VpcId"`
		Protocol            *string `json:"Protocol" name:"Protocol"`
		BackendServerNumber *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive   *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		IpVersion           *string `json:"IpVersion" name:"IpVersion"`
		Method              *string `json:"Method" name:"Method"`
		SlowStartEnabled    *bool   `json:"SlowStartEnabled" name:"SlowStartEnabled"`
		SlowStartDuration   *int    `json:"SlowStartDuration" name:"SlowStartDuration"`
	} `json:"BackendServerGroup"`
}

func (r *CreateBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeListenersRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId []*string                  `json:"ListenerId,omitempty" name:"ListenerId"`
	Filter     []*DescribeListenersFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                    `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeListenersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	NextToken   *string `json:"NextToken" name:"NextToken"`
	ListenerSet []struct {
		ListenerId         *string `json:"ListenerId" name:"ListenerId"`
		LoadBalancerId     *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		ListenerName       *string `json:"ListenerName" name:"ListenerName"`
		Protocol           *string `json:"Protocol" name:"Protocol"`
		Port               *int    `json:"Port" name:"Port"`
		CertificateId      *string `json:"CertificateId" name:"CertificateId"`
		TlsCipherPolicy    *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		DefaultBackendServerGroupId *string `json:"DefaultBackendServerGroupId" name:"DefaultBackendServerGroupId"`
		ListenerAclId      *string `json:"ListenerAclId" name:"ListenerAclId"`
		ListenerState      *string `json:"ListenerState" name:"ListenerState"`
		RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		RedirectListenerName *string `json:"RedirectListenerName" name:"RedirectListenerName"`
		HttpProtocol       *string `json:"HttpProtocol" name:"HttpProtocol"`
		EnableHttp2        *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		CaCertificateId    *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled          *bool   `json:"CaEnabled" name:"CaEnabled"`
		EnableQuicUpgrade  *bool   `json:"EnableQuicUpgrade" name:"EnableQuicUpgrade"`
		QuicListenerId     *string `json:"QuicListenerId" name:"QuicListenerId"`
		IdleTimeout        *int    `json:"IdleTimeout" name:"IdleTimeout"`
		ServerGroupId      *string `json:"ServerGroupId" name:"ServerGroupId"`
	} `json:"ListenerSet"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyListenerRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId        *string `json:"ListenerId,omitempty" name:"ListenerId"`
	ListenerName      *string `json:"ListenerName,omitempty" name:"ListenerName"`
	ListenerState     *string `json:"ListenerState,omitempty" name:"ListenerState"`
	CertificateId     *string `json:"CertificateId,omitempty" name:"CertificateId"`
	TlsCipherPolicy   *string `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
	ListenerAclId     *string `json:"ListenerAclId,omitempty" name:"ListenerAclId"`
	HttpProtocol      *string `json:"HttpProtocol,omitempty" name:"HttpProtocol"`
	EnableHttp2       *bool   `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
	CaEnabled         *bool   `json:"CaEnabled,omitempty" name:"CaEnabled"`
	CaCertificateId   *string `json:"CaCertificateId,omitempty" name:"CaCertificateId"`
	EnableQuicUpgrade *bool   `json:"EnableQuicUpgrade,omitempty" name:"EnableQuicUpgrade"`
	QuicListenerId    *string `json:"QuicListenerId,omitempty" name:"QuicListenerId"`
	IdleTimeout       *int    `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
	ServerGroupId     *string `json:"ServerGroupId,omitempty" name:"ServerGroupId"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Listener  struct {
		ListenerId         *string `json:"ListenerId" name:"ListenerId"`
		LoadBalancerId     *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		ListenerName       *string `json:"ListenerName" name:"ListenerName"`
		Protocol           *string `json:"Protocol" name:"Protocol"`
		Port               *int    `json:"Port" name:"Port"`
		CertificateId      *string `json:"CertificateId" name:"CertificateId"`
		TlsCipherPolicy    *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		DefaultBackendServerGroupId *string `json:"DefaultBackendServerGroupId" name:"DefaultBackendServerGroupId"`
		ListenerAclId      *string `json:"ListenerAclId" name:"ListenerAclId"`
		ListenerState      *string `json:"ListenerState" name:"ListenerState"`
		RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		RedirectListenerName *string `json:"RedirectListenerName" name:"RedirectListenerName"`
		HttpProtocol       *string `json:"HttpProtocol" name:"HttpProtocol"`
		EnableHttp2        *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		CaCertificateId    *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled          *bool   `json:"CaEnabled" name:"CaEnabled"`
		EnableQuicUpgrade  *bool   `json:"EnableQuicUpgrade" name:"EnableQuicUpgrade"`
		QuicListenerId     *string `json:"QuicListenerId" name:"QuicListenerId"`
		IdleTimeout        *int    `json:"IdleTimeout" name:"IdleTimeout"`
		ServerGroupId      *string `json:"ServerGroupId" name:"ServerGroupId"`
	} `json:"Listener"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteListenerRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateListenerRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId       *string                            `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ListenerName         *string                            `json:"ListenerName,omitempty" name:"ListenerName"`
	Protocol             *string                            `json:"Protocol,omitempty" name:"Protocol"`
	Port                 *int                               `json:"Port,omitempty" name:"Port"`
	CertificateId        *string                            `json:"CertificateId,omitempty" name:"CertificateId"`
	TlsCipherPolicy      *string                            `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
	ListenerAclId        *string                            `json:"ListenerAclId,omitempty" name:"ListenerAclId"`
	ListenerState        *string                            `json:"ListenerState,omitempty" name:"ListenerState"`
	RedirectListenerId   *string                            `json:"RedirectListenerId,omitempty" name:"RedirectListenerId"`
	RedirectHttpCode     *string                            `json:"RedirectHttpCode,omitempty" name:"RedirectHttpCode"`
	EnableHttp2          *bool                              `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
	BackendServerGroupId *string                            `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	FixedResponseConfig  *CreateListenerFixedResponseConfig `json:"FixedResponseConfig,omitempty" name:"FixedResponseConfig"`
	RewriteConfig        *CreateListenerRewriteConfig       `json:"RewriteConfig,omitempty" name:"RewriteConfig"`
	CaEnabled            *bool                              `json:"CaEnabled,omitempty" name:"CaEnabled"`
	CaCertificateId      *string                            `json:"CaCertificateId,omitempty" name:"CaCertificateId"`
	EnableQuicUpgrade    *bool                              `json:"EnableQuicUpgrade,omitempty" name:"EnableQuicUpgrade"`
	QuicListenerId       *string                            `json:"QuicListenerId,omitempty" name:"QuicListenerId"`
	IdleTimeout          *int                               `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
	ServerGroupId        *string                            `json:"ServerGroupId,omitempty" name:"ServerGroupId"`
}

func (r *CreateListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Listener  struct {
		ListenerId         *string `json:"ListenerId" name:"ListenerId"`
		LoadBalancerId     *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		ListenerName       *string `json:"ListenerName" name:"ListenerName"`
		Protocol           *string `json:"Protocol" name:"Protocol"`
		Port               *int    `json:"Port" name:"Port"`
		CertificateId      *string `json:"CertificateId" name:"CertificateId"`
		TlsCipherPolicy    *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		DefaultBackendServerGroupId *string `json:"DefaultBackendServerGroupId" name:"DefaultBackendServerGroupId"`
		ListenerAclId      *string `json:"ListenerAclId" name:"ListenerAclId"`
		ListenerState      *string `json:"ListenerState" name:"ListenerState"`
		RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		RedirectListenerName *string `json:"RedirectListenerName" name:"RedirectListenerName"`
		HttpProtocol       *string `json:"HttpProtocol" name:"HttpProtocol"`
		EnableHttp2        *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		CaCertificateId    *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled          *bool   `json:"CaEnabled" name:"CaEnabled"`
		EnableQuicUpgrade  *bool   `json:"EnableQuicUpgrade" name:"EnableQuicUpgrade"`
		QuicListenerId     *string `json:"QuicListenerId" name:"QuicListenerId"`
		IdleTimeout        *int    `json:"IdleTimeout" name:"IdleTimeout"`
		ServerGroupId      *string `json:"ServerGroupId" name:"ServerGroupId"`
	} `json:"Listener"`
}

func (r *CreateListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetAccessLogRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ProjectName    *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogpoolName    *string `json:"LogpoolName,omitempty" name:"LogpoolName"`
}

func (r *SetAccessLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetAccessLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	LoadBalancer struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancer"`
}

func (r *SetAccessLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetEnableAccessLogRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	EnabledLog     *bool   `json:"EnabledLog,omitempty" name:"EnabledLog"`
}

func (r *SetEnableAccessLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetEnableAccessLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	LoadBalancer struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancer"`
}

func (r *SetEnableAccessLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetEnableAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetLbProtocolLayersRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ProtocolLayers *string `json:"ProtocolLayers,omitempty" name:"ProtocolLayers"`
}

func (r *SetLbProtocolLayersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetLbProtocolLayersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	LoadBalancer struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancer"`
}

func (r *SetLbProtocolLayersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLbProtocolLayersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetLoadBalancerStatusRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	State          *string `json:"State,omitempty" name:"State"`
}

func (r *SetLoadBalancerStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetLoadBalancerStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	LoadBalancer struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancer"`
}

func (r *SetLoadBalancerStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetLoadBalancerNameRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
}

func (r *SetLoadBalancerNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetLoadBalancerNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	LoadBalancer struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancer"`
}

func (r *SetLoadBalancerNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeLoadBalancersRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId []*string                      `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	Filter         []*DescribeLoadBalancersFilter `json:"Filter,omitempty" name:"Filter"`
	IsContainTag   *bool                          `json:"IsContainTag,omitempty" name:"IsContainTag"`
	TagKey         []*string                      `json:"TagKey,omitempty" name:"TagKey"`
	TagKV          []*DescribeLoadBalancersTagKV  `json:"TagKV,omitempty" name:"TagKV"`
	ProjectId      []*string                      `json:"ProjectId,omitempty" name:"ProjectId"`
	MaxResults     *int                           `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken      *string                        `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeLoadBalancersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	NextToken       *string `json:"NextToken" name:"NextToken"`
	LoadBalancerSet []struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancerSet"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteLoadBalancerRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteLoadBalancerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateLoadBalancerRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerName       *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	LoadBalancerVersion    *string `json:"LoadBalancerVersion,omitempty" name:"LoadBalancerVersion"`
	LoadBalancerType       *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	VpcId                  *string `json:"VpcId,omitempty" name:"VpcId"`
	IpVersion              *string `json:"IpVersion,omitempty" name:"IpVersion"`
	ProjectId              *string `json:"ProjectId,omitempty" name:"ProjectId"`
	AllocationId           *string `json:"AllocationId,omitempty" name:"AllocationId"`
	ChargeType             *string `json:"ChargeType,omitempty" name:"ChargeType"`
	SubnetId               *string `json:"SubnetId,omitempty" name:"SubnetId"`
	PrivateIpAddress       *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	EnabledQuic            *bool   `json:"EnabledQuic,omitempty" name:"EnabledQuic"`
	EnableHpa              *bool   `json:"EnableHpa,omitempty" name:"EnableHpa"`
	ProtocolLayers         *string `json:"ProtocolLayers,omitempty" name:"ProtocolLayers"`
	DeleteProtection       *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
	ModificationProtection *string `json:"ModificationProtection,omitempty" name:"ModificationProtection"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateLoadBalancerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	LoadBalancer struct {
		LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		LoadBalancerVersion *string `json:"LoadBalancerVersion" name:"LoadBalancerVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerType *string `json:"LoadBalancerType" name:"LoadBalancerType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		BillType       *int    `json:"BillType" name:"BillType"`
		ProductWhat    *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId       *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic    *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa      *bool   `json:"EnableHpa" name:"EnableHpa"`
		BindWafStatus  *string `json:"BindWafStatus" name:"BindWafStatus"`
		WafInfo        struct {
			WafId *string `json:"WafId" name:"WafId"`
		} `json:"WafInfo" name:"WafInfo"`
		ProtocolLayers *string `json:"ProtocolLayers" name:"ProtocolLayers"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancer"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyCertificateWithGroupRequest struct {
	*ksyunhttp.BaseRequest
	ListenerCertGroupId *string `json:"ListenerCertGroupId,omitempty" name:"ListenerCertGroupId"`
	OldCertificateId    *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
	CertificateId       *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *ModifyCertificateWithGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCertificateWithGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyCertificateWithGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertificateWithGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DissociateCertificateWithGroupRequest struct {
	*ksyunhttp.BaseRequest
	ListenerCertGroupId *string `json:"ListenerCertGroupId,omitempty" name:"ListenerCertGroupId"`
	CertificateId       *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DissociateCertificateWithGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DissociateCertificateWithGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DissociateCertificateWithGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DissociateCertificateWithGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AssociateCertificateWithGroupRequest struct {
	*ksyunhttp.BaseRequest
	ListenerCertGroupId *string `json:"ListenerCertGroupId,omitempty" name:"ListenerCertGroupId"`
	CertificateId       *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *AssociateCertificateWithGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AssociateCertificateWithGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateCertificateWithGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateCertificateWithGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeListenerCertGroupsRequest struct {
	*ksyunhttp.BaseRequest
	ListenerCertGroupId []*string                           `json:"ListenerCertGroupId,omitempty" name:"ListenerCertGroupId"`
	Filter              []*DescribeListenerCertGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults          *int                                `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken           *string                             `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeListenerCertGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeListenerCertGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	NextToken            *string `json:"NextToken" name:"NextToken"`
	ListenerCertGroupSet []struct {
		ListenerId      *string `json:"ListenerId" name:"ListenerId"`
		ListenerCertGroupId *string `json:"ListenerCertGroupId" name:"ListenerCertGroupId"`
		ListenerCertSet []struct {
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			CertificateId   *string `json:"CertificateId" name:"CertificateId"`
			CertificateName *string `json:"CertificateName" name:"CertificateName"`
			CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
			CommonName      *string `json:"CommonName" name:"CommonName"`
			ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		} `json:"ListenerCertSet" name:"ListenerCertSet"`
	} `json:"ListenerCertGroupSet"`
}

func (r *DescribeListenerCertGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenerCertGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteListenerCertGroupRequest struct {
	*ksyunhttp.BaseRequest
	ListenerCertGroupId *string `json:"ListenerCertGroupId,omitempty" name:"ListenerCertGroupId"`
}

func (r *DeleteListenerCertGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteListenerCertGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteListenerCertGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenerCertGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateListenerCertGroupRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *CreateListenerCertGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateListenerCertGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	ListenerCertGroup struct {
		ListenerId      *string `json:"ListenerId" name:"ListenerId"`
		ListenerCertGroupId *string `json:"ListenerCertGroupId" name:"ListenerCertGroupId"`
		ListenerCertSet []struct {
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			CertificateId   *string `json:"CertificateId" name:"CertificateId"`
			CertificateName *string `json:"CertificateName" name:"CertificateName"`
			CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
			CommonName      *string `json:"CommonName" name:"CommonName"`
			ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		} `json:"ListenerCertSet" name:"ListenerCertSet"`
	} `json:"ListenerCertGroup"`
}

func (r *CreateListenerCertGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerCertGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddRulesRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupId   *string                `json:"RuleGroupId,omitempty" name:"RuleGroupId"`
	RuleType      *string                `json:"RuleType,omitempty" name:"RuleType"`
	RuleValue     *string                `json:"RuleValue,omitempty" name:"RuleValue"`
	MethodValue   []*string              `json:"MethodValue,omitempty" name:"MethodValue"`
	SourceIpValue []*string              `json:"SourceIpValue,omitempty" name:"SourceIpValue"`
	HeaderValue   []*AddRulesHeaderValue `json:"HeaderValue,omitempty" name:"HeaderValue"`
	QueryValue    []*AddRulesQueryValue  `json:"QueryValue,omitempty" name:"QueryValue"`
	CookieValue   []*AddRulesCookieValue `json:"CookieValue,omitempty" name:"CookieValue"`
}

func (r *AddRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Rule      struct {
		RuleType      *string   `json:"RuleType" name:"RuleType"`
		RuleValue     *string   `json:"RuleValue" name:"RuleValue"`
		MethodValue   []*string `json:"MethodValue" name:"MethodValue"`
		SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
		HeaderValue   []struct {
			Key *string `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"HeaderValue" name:"HeaderValue"`
		QueryValue []struct {
			Key *string `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"QueryValue" name:"QueryValue"`
		CookieValue []struct {
			Key *string `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"CookieValue" name:"CookieValue"`
	} `json:"Rule"`
}

func (r *AddRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupId *string `json:"RuleGroupId,omitempty" name:"RuleGroupId"`
	RuleType    *string `json:"RuleType,omitempty" name:"RuleType"`
	RuleValue   *string `json:"RuleValue,omitempty" name:"RuleValue"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupId *string `json:"RuleGroupId,omitempty" name:"RuleGroupId"`
	RuleType    *string `json:"RuleType,omitempty" name:"RuleType"`
	RuleValue   *string `json:"RuleValue,omitempty" name:"RuleValue"`
}

func (r *AddRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Rule      struct {
		RuleType      *string   `json:"RuleType" name:"RuleType"`
		RuleValue     *string   `json:"RuleValue" name:"RuleValue"`
		MethodValue   []*string `json:"MethodValue" name:"MethodValue"`
		SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
		HeaderValue   []struct {
			Key *string `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"HeaderValue" name:"HeaderValue"`
		QueryValue []struct {
			Key *string `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"QueryValue" name:"QueryValue"`
		CookieValue []struct {
			Key *string `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"CookieValue" name:"CookieValue"`
	} `json:"Rule"`
}

func (r *AddRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyRuleGroupRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupId          *string                             `json:"RuleGroupId,omitempty" name:"RuleGroupId"`
	RuleGroupName        *string                             `json:"RuleGroupName,omitempty" name:"RuleGroupName"`
	BackendServerGroupId *string                             `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	Type                 *string                             `json:"Type,omitempty" name:"Type"`
	RuleSet              []*ModifyRuleGroupRuleSet           `json:"RuleSet,omitempty" name:"RuleSet"`
	RedirectListenerId   *string                             `json:"RedirectListenerId,omitempty" name:"RedirectListenerId"`
	RedirectHttpCode     *string                             `json:"RedirectHttpCode,omitempty" name:"RedirectHttpCode"`
	FixedResponseConfig  *ModifyRuleGroupFixedResponseConfig `json:"FixedResponseConfig,omitempty" name:"FixedResponseConfig"`
	RewriteConfig        *ModifyRuleGroupRewriteConfig       `json:"RewriteConfig,omitempty" name:"RewriteConfig"`
}

func (r *ModifyRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyRuleGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	RuleGroup struct {
		RuleGroupId *string `json:"RuleGroupId" name:"RuleGroupId"`
		ListenerId  *string `json:"ListenerId" name:"ListenerId"`
		RuleGroupName *string `json:"RuleGroupName" name:"RuleGroupName"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		RuleSet     []struct {
			RuleType    *string `json:"RuleType" name:"RuleType"`
			RuleValue   *string `json:"RuleValue" name:"RuleValue"`
			MethodValue []*string `json:"MethodValue" name:"MethodValue"`
			SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
			HeaderValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"HeaderValue"`
			QueryValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"QueryValue"`
			CookieValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"CookieValue"`
		} `json:"RuleSet" name:"RuleSet"`
		RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		RedirectHttpCode *string `json:"RedirectHttpCode" name:"RedirectHttpCode"`
		RewriteConfig struct {
			HttpHost    *string `json:"HttpHost" name:"HttpHost"`
			Url         *string `json:"Url" name:"Url"`
			QueryString *string `json:"QueryString" name:"QueryString"`
		} `json:"RewriteConfig" name:"RewriteConfig"`
		FixedResponseConfig struct {
			Content     *string `json:"Content" name:"Content"`
			ContentType *string `json:"ContentType" name:"ContentType"`
			HttpCode    *string `json:"HttpCode" name:"HttpCode"`
		} `json:"FixedResponseConfig" name:"FixedResponseConfig"`
	} `json:"RuleGroup"`
}

func (r *ModifyRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRuleGroupsRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupId []*string                   `json:"RuleGroupId,omitempty" name:"RuleGroupId"`
	Filter      []*DescribeRuleGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults  *int                        `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken   *string                     `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRuleGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	NextToken    *string `json:"NextToken" name:"NextToken"`
	RuleGroupSet []struct {
		RuleGroupId *string `json:"RuleGroupId" name:"RuleGroupId"`
		ListenerId  *string `json:"ListenerId" name:"ListenerId"`
		RuleGroupName *string `json:"RuleGroupName" name:"RuleGroupName"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		RuleSet     []struct {
			RuleType    *string `json:"RuleType" name:"RuleType"`
			RuleValue   *string `json:"RuleValue" name:"RuleValue"`
			MethodValue []*string `json:"MethodValue" name:"MethodValue"`
			SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
			HeaderValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"HeaderValue"`
			QueryValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"QueryValue"`
			CookieValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"CookieValue"`
		} `json:"RuleSet" name:"RuleSet"`
		RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		RedirectHttpCode *string `json:"RedirectHttpCode" name:"RedirectHttpCode"`
		RewriteConfig struct {
			HttpHost    *string `json:"HttpHost" name:"HttpHost"`
			Url         *string `json:"Url" name:"Url"`
			QueryString *string `json:"QueryString" name:"QueryString"`
		} `json:"RewriteConfig" name:"RewriteConfig"`
		FixedResponseConfig struct {
			Content     *string `json:"Content" name:"Content"`
			ContentType *string `json:"ContentType" name:"ContentType"`
			HttpCode    *string `json:"HttpCode" name:"HttpCode"`
		} `json:"FixedResponseConfig" name:"FixedResponseConfig"`
	} `json:"RuleGroupSet"`
}

func (r *DescribeRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRuleGroupRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupId *string `json:"RuleGroupId,omitempty" name:"RuleGroupId"`
}

func (r *DeleteRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRuleGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateRuleGroupRequest struct {
	*ksyunhttp.BaseRequest
	RuleGroupName        *string                             `json:"RuleGroupName,omitempty" name:"RuleGroupName"`
	ListenerId           *string                             `json:"ListenerId,omitempty" name:"ListenerId"`
	BackendServerGroupId *string                             `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	Type                 *string                             `json:"Type,omitempty" name:"Type"`
	RuleSet              []*CreateRuleGroupRuleSet           `json:"RuleSet,omitempty" name:"RuleSet"`
	RedirectListenerId   *string                             `json:"RedirectListenerId,omitempty" name:"RedirectListenerId"`
	RedirectHttpCode     *string                             `json:"RedirectHttpCode,omitempty" name:"RedirectHttpCode"`
	FixedResponseConfig  *CreateRuleGroupFixedResponseConfig `json:"FixedResponseConfig,omitempty" name:"FixedResponseConfig"`
	RewriteConfig        *CreateRuleGroupRewriteConfig       `json:"RewriteConfig,omitempty" name:"RewriteConfig"`
}

func (r *CreateRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateRuleGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	RuleGroup struct {
		RuleGroupId *string `json:"RuleGroupId" name:"RuleGroupId"`
		ListenerId  *string `json:"ListenerId" name:"ListenerId"`
		RuleGroupName *string `json:"RuleGroupName" name:"RuleGroupName"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		RuleSet     []struct {
			RuleType    *string `json:"RuleType" name:"RuleType"`
			RuleValue   *string `json:"RuleValue" name:"RuleValue"`
			MethodValue []*string `json:"MethodValue" name:"MethodValue"`
			SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
			HeaderValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"HeaderValue"`
			QueryValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"QueryValue"`
			CookieValue []struct {
				Key   *string   `json:"Key" name:"Key"`
				Value []*string `json:"Value" name:"Value"`
			} `json:"CookieValue"`
		} `json:"RuleSet" name:"RuleSet"`
		RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		RedirectHttpCode *string `json:"RedirectHttpCode" name:"RedirectHttpCode"`
		RewriteConfig struct {
			HttpHost    *string `json:"HttpHost" name:"HttpHost"`
			Url         *string `json:"Url" name:"Url"`
			QueryString *string `json:"QueryString" name:"QueryString"`
		} `json:"RewriteConfig" name:"RewriteConfig"`
		FixedResponseConfig struct {
			Content     *string `json:"Content" name:"Content"`
			ContentType *string `json:"ContentType" name:"ContentType"`
			HttpCode    *string `json:"HttpCode" name:"HttpCode"`
		} `json:"FixedResponseConfig" name:"FixedResponseConfig"`
	} `json:"RuleGroup"`
}

func (r *CreateRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetLBModificationProtectionRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ModificationProtection *bool   `json:"ModificationProtection,omitempty" name:"ModificationProtection"`
}

func (r *SetLBModificationProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetLBModificationProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetLBModificationProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLBModificationProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetLBDeleteProtectionRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	DeleteProtection *bool   `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
}

func (r *SetLBDeleteProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetLBDeleteProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetLBDeleteProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLBDeleteProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

