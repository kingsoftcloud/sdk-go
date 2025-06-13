package v20160304

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeListenersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeInstancesWithListenerFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeHealthChecksFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeLoadBalancersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeLoadBalancersTagKV struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type DescribeHostHeadersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeRulesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeBackendServerGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeBackendServersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ListPrivateLinkServerFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeAlbsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeAlbsTagKV struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateAlbListenerFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	HttpCode    *string `json:"HttpCode,omitempty" name:"HttpCode"`
}
type CreateAlbListenerRewriteConfig struct {
	HttpHost    *string `json:"HttpHost,omitempty" name:"HttpHost"`
	Url         *string `json:"Url,omitempty" name:"Url"`
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}
type DescribeAlbListenersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateAlbRuleGroupAlbRuleSetHeaderValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateAlbRuleGroupAlbRuleSetQueryValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateAlbRuleGroupAlbRuleSetCookieValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateAlbRuleGroupAlbRuleSet struct {
	AlbRuleType   *string                                    `json:"AlbRuleType,omitempty" name:"AlbRuleType"`
	AlbRuleValue  *string                                    `json:"AlbRuleValue,omitempty" name:"AlbRuleValue"`
	MethodValue   []*string                                  `json:"MethodValue,omitempty" name:"MethodValue"`
	SourceIpValue []*string                                  `json:"SourceIpValue,omitempty" name:"SourceIpValue"`
	HeaderValue   []*CreateAlbRuleGroupAlbRuleSetHeaderValue `json:"HeaderValue,omitempty" name:"HeaderValue"`
	QueryValue    []*CreateAlbRuleGroupAlbRuleSetQueryValue  `json:"QueryValue,omitempty" name:"QueryValue"`
	CookieValue   []*CreateAlbRuleGroupAlbRuleSetCookieValue `json:"CookieValue,omitempty" name:"CookieValue"`
}
type CreateAlbRuleGroupFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	HttpCode    *string `json:"HttpCode,omitempty" name:"HttpCode"`
}
type CreateAlbRuleGroupRewriteConfig struct {
	HttpHost    *string `json:"HttpHost,omitempty" name:"HttpHost"`
	Url         *string `json:"Url,omitempty" name:"Url"`
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}
type DescribeAlbRuleGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyAlbRuleGroupAlbRuleSetHeaderValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyAlbRuleGroupAlbRuleSetQueryValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyAlbRuleGroupAlbRuleSetCookieValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type ModifyAlbRuleGroupAlbRuleSet struct {
	AlbRuleType   *string                                    `json:"AlbRuleType,omitempty" name:"AlbRuleType"`
	AlbRuleValue  *string                                    `json:"AlbRuleValue,omitempty" name:"AlbRuleValue"`
	MethodValue   []*string                                  `json:"MethodValue,omitempty" name:"MethodValue"`
	SourceIpValue []*string                                  `json:"SourceIpValue,omitempty" name:"SourceIpValue"`
	HeaderValue   []*ModifyAlbRuleGroupAlbRuleSetHeaderValue `json:"HeaderValue,omitempty" name:"HeaderValue"`
	QueryValue    []*ModifyAlbRuleGroupAlbRuleSetQueryValue  `json:"QueryValue,omitempty" name:"QueryValue"`
	CookieValue   []*ModifyAlbRuleGroupAlbRuleSetCookieValue `json:"CookieValue,omitempty" name:"CookieValue"`
}
type ModifyAlbRuleGroupFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" name:"Content"`
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	HttpCode    *string `json:"HttpCode,omitempty" name:"HttpCode"`
}
type DescribeAlbListenerCertGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeAlbBackendServerGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeAlbBackendServersFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddAlbRulesHeaderValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddAlbRulesQueryValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type AddAlbRulesCookieValue struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
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

func (r *DescribeListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	NextToken   *string `json:"NextToken" name:"NextToken"`
	ListenerSet []struct {
		LoadBalancerId      *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		ListenerName        *string `json:"ListenerName" name:"ListenerName"`
		ListenerId          *string `json:"ListenerId" name:"ListenerId"`
		ListenerState       *string `json:"ListenerState" name:"ListenerState"`
		CertificateId       *string `json:"CertificateId" name:"CertificateId"`
		ListenerProtocol    *string `json:"ListenerProtocol" name:"ListenerProtocol"`
		ListenerPort        *int    `json:"ListenerPort" name:"ListenerPort"`
		Method              *string `json:"Method" name:"Method"`
		BandWidthIn         *int    `json:"BandWidthIn" name:"BandWidthIn"`
		BandWidthOut        *int    `json:"BandWidthOut" name:"BandWidthOut"`
		LoadBalancerAclId   *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
		HttpProtocol        *string `json:"HttpProtocol" name:"HttpProtocol"`
		TlsCipherPolicy     *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		EnableHttp2         *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		RedirectListenerId  *string `json:"RedirectListenerId" name:"RedirectListenerId"`
		IpVersion           *string `json:"IpVersion" name:"IpVersion"`
		AsPrivateLinkServer *bool   `json:"AsPrivateLinkServer" name:"AsPrivateLinkServer"`
		AsPrivateLink       *bool   `json:"AsPrivateLink" name:"AsPrivateLink"`
		HealthStatus        *bool   `json:"HealthStatus" name:"HealthStatus"`
		BindType            *string `json:"BindType" name:"BindType"`
		HealthCheck         struct {
			ListenerId             *string `json:"ListenerId" name:"ListenerId"`
			HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthCheckId          *string `json:"HealthCheckId" name:"HealthCheckId"`
			HttpMethod             *string `json:"HttpMethod" name:"HttpMethod"`
			HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval               *int    `json:"Interval" name:"Interval"`
			Timeout                *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath                *string `json:"UrlPath" name:"UrlPath"`
			HostName               *string `json:"HostName" name:"HostName"`
			HealthCheckReq         *string `json:"HealthCheckReq" name:"HealthCheckReq"`
			HealthCheckExp         *string `json:"HealthCheckExp" name:"HealthCheckExp"`
			HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
		BackendServerGroupIdSet []struct {
			BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		} `json:"BackendServerGroupIdSet" name:"BackendServerGroupIdSet"`
		RealServer []struct {
			RegisterId         *string `json:"RegisterId" name:"RegisterId"`
			RealServerState    *string `json:"RealServerState" name:"RealServerState"`
			RealServerType     *string `json:"RealServerType" name:"RealServerType"`
			ListenerId         *string `json:"ListenerId" name:"ListenerId"`
			Weight             *int    `json:"Weight" name:"Weight"`
			RealServerIp       *string `json:"RealServerIp" name:"RealServerIp"`
			RealServerPort     *int    `json:"RealServerPort" name:"RealServerPort"`
			InstanceId         *string `json:"InstanceId" name:"InstanceId"`
			Tag                *string `json:"Tag" name:"Tag"`
			MasterSlaveType    *string `json:"MasterSlaveType" name:"MasterSlaveType"`
			NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		} `json:"RealServer" name:"RealServer"`
		BackendServerSet []struct {
			BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
			BackendServerIp        *string `json:"BackendServerIp" name:"BackendServerIp"`
			InstanceId             *string `json:"InstanceId" name:"InstanceId"`
			RegisterId             *string `json:"RegisterId" name:"RegisterId"`
			BackendServerPort      *int    `json:"BackendServerPort" name:"BackendServerPort"`
			DirectConnectGatewayId *string `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
			Weight                 *int    `json:"Weight" name:"Weight"`
			BackendServerState     *string `json:"BackendServerState" name:"BackendServerState"`
		} `json:"BackendServerSet" name:"BackendServerSet"`
		CaCertificateId   *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled         *bool   `json:"CaEnabled" name:"CaEnabled"`
		UpstreamKeepalive *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
	} `json:"ListenerSet"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenersRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyListenersRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId               *string `json:"ListenerId,omitempty" name:"ListenerId"`
	ListenerName             *string `json:"ListenerName,omitempty" name:"ListenerName"`
	BindType                 *string `json:"BindType,omitempty" name:"BindType"`
	ListenerState            *string `json:"ListenerState,omitempty" name:"ListenerState"`
	Method                   *string `json:"Method,omitempty" name:"Method"`
	BandWidthIn              *int    `json:"BandWidthIn,omitempty" name:"BandWidthIn"`
	BandWidthOut             *int    `json:"BandWidthOut,omitempty" name:"BandWidthOut"`
	HttpProtocol             *string `json:"HttpProtocol,omitempty" name:"HttpProtocol"`
	TlsCipherPolicy          *string `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
	EnableHttp2              *bool   `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
	SessionState             *string `json:"SessionState,omitempty" name:"SessionState"`
	SessionPersistencePeriod *int    `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
	CookieType               *string `json:"CookieType,omitempty" name:"CookieType"`
	CookieName               *string `json:"CookieName,omitempty" name:"CookieName"`
	CertificateId            *string `json:"CertificateId,omitempty" name:"CertificateId"`
	RedirectListenerId       *string `json:"RedirectListenerId,omitempty" name:"RedirectListenerId"`
	CaCertificateId          *string `json:"CaCertificateId,omitempty" name:"CaCertificateId"`
	CaEnabled                *bool   `json:"CaEnabled,omitempty" name:"CaEnabled"`
	UpstreamKeepalive        *string `json:"UpstreamKeepalive,omitempty" name:"UpstreamKeepalive"`
}

func (r *ModifyListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyListenersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	LoadBalancerId     *string `json:"LoadBalancerId" name:"LoadBalancerId"`
	CreateTime         *string `json:"CreateTime" name:"CreateTime"`
	ListenerName       *string `json:"ListenerName" name:"ListenerName"`
	ListenerId         *string `json:"ListenerId" name:"ListenerId"`
	ListenerState      *string `json:"ListenerState" name:"ListenerState"`
	ListenerProtocol   *string `json:"ListenerProtocol" name:"ListenerProtocol"`
	ListenerPort       *int    `json:"ListenerPort" name:"ListenerPort"`
	Method             *string `json:"Method" name:"Method"`
	BandWidthIn        *int    `json:"BandWidthIn" name:"BandWidthIn"`
	BandWidthOut       *int    `json:"BandWidthOut" name:"BandWidthOut"`
	LoadBalancerAclId  *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
	HttpProtocol       *string `json:"HttpProtocol" name:"HttpProtocol"`
	TlsCipherPolicy    *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
	EnableHttp2        *bool   `json:"EnableHttp2" name:"EnableHttp2"`
	RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
	HealthCheck        struct {
		ListenerId             *string `json:"ListenerId" name:"ListenerId"`
		HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
		HealthCheckId          *string `json:"HealthCheckId" name:"HealthCheckId"`
		HttpMethod             *string `json:"HttpMethod" name:"HttpMethod"`
		HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
		Interval               *int    `json:"Interval" name:"Interval"`
		Timeout                *int    `json:"Timeout" name:"Timeout"`
		UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
		UrlPath                *string `json:"UrlPath" name:"UrlPath"`
		HostName               *string `json:"HostName" name:"HostName"`
		HealthCheckReq         *string `json:"HealthCheckReq" name:"HealthCheckReq"`
		HealthCheckExp         *string `json:"HealthCheckExp" name:"HealthCheckExp"`
		HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
	} `json:"HealthCheck"`
	Session struct {
		SessionState             *string `json:"SessionState" name:"SessionState"`
		SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
		CookieType               *string `json:"CookieType" name:"CookieType"`
		CookieName               *string `json:"CookieName" name:"CookieName"`
	} `json:"Session"`
	RealServer []struct {
		RegisterId         *string `json:"RegisterId" name:"RegisterId"`
		RealServerState    *string `json:"RealServerState" name:"RealServerState"`
		RealServerType     *string `json:"RealServerType" name:"RealServerType"`
		ListenerId         *string `json:"ListenerId" name:"ListenerId"`
		Weight             *int    `json:"Weight" name:"Weight"`
		RealServerIp       *string `json:"RealServerIp" name:"RealServerIp"`
		RealServerPort     *int    `json:"RealServerPort" name:"RealServerPort"`
		InstanceId         *string `json:"InstanceId" name:"InstanceId"`
		Tag                *string `json:"Tag" name:"Tag"`
		MasterSlaveType    *string `json:"MasterSlaveType" name:"MasterSlaveType"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	} `json:"RealServer"`
	CertificateId     *string `json:"CertificateId" name:"CertificateId"`
	UpstreamKeepalive *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
}

func (r *ModifyListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenersRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId           *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ListenerName             *string `json:"ListenerName,omitempty" name:"ListenerName"`
	ListenerState            *string `json:"ListenerState,omitempty" name:"ListenerState"`
	ListenerProtocol         *string `json:"ListenerProtocol,omitempty" name:"ListenerProtocol"`
	BindType                 *string `json:"BindType,omitempty" name:"BindType"`
	ListenerPort             *int    `json:"ListenerPort,omitempty" name:"ListenerPort"`
	Method                   *string `json:"Method,omitempty" name:"Method"`
	BandWidthIn              *int    `json:"BandWidthIn,omitempty" name:"BandWidthIn"`
	BandWidthOut             *int    `json:"BandWidthOut,omitempty" name:"BandWidthOut"`
	LoadBalancerAclId        *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
	HttpProtocol             *string `json:"HttpProtocol,omitempty" name:"HttpProtocol"`
	TlsCipherPolicy          *string `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
	EnableHttp2              *bool   `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
	RedirectListenerId       *string `json:"RedirectListenerId,omitempty" name:"RedirectListenerId"`
	SessionState             *string `json:"SessionState,omitempty" name:"SessionState"`
	SessionPersistencePeriod *int    `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
	CookieType               *string `json:"CookieType,omitempty" name:"CookieType"`
	CookieName               *string `json:"CookieName,omitempty" name:"CookieName"`
	CertificateId            *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CaCertificateId          *string `json:"CaCertificateId,omitempty" name:"CaCertificateId"`
	CaEnabled                *bool   `json:"CaEnabled,omitempty" name:"CaEnabled"`
	UpstreamKeepalive        *string `json:"UpstreamKeepalive,omitempty" name:"UpstreamKeepalive"`
}

func (r *CreateListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	LoadBalancerId     *string `json:"LoadBalancerId" name:"LoadBalancerId"`
	CreateTime         *string `json:"CreateTime" name:"CreateTime"`
	ListenerName       *string `json:"ListenerName" name:"ListenerName"`
	ListenerId         *string `json:"ListenerId" name:"ListenerId"`
	ListenerState      *string `json:"ListenerState" name:"ListenerState"`
	ListenerProtocol   *string `json:"ListenerProtocol" name:"ListenerProtocol"`
	ListenerPort       *int    `json:"ListenerPort" name:"ListenerPort"`
	Method             *string `json:"Method" name:"Method"`
	BandWidthIn        *int    `json:"BandWidthIn" name:"BandWidthIn"`
	BandWidthOut       *int    `json:"BandWidthOut" name:"BandWidthOut"`
	LoadBalancerAclId  *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
	HttpProtocol       *string `json:"HttpProtocol" name:"HttpProtocol"`
	TlsCipherPolicy    *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
	EnableHttp2        *bool   `json:"EnableHttp2" name:"EnableHttp2"`
	RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
	IpVersion          *string `json:"IpVersion" name:"IpVersion"`
	UpstreamKeepalive  *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
	HealthCheck        struct {
		ListenerId             *string `json:"ListenerId" name:"ListenerId"`
		HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
		HealthCheckId          *string `json:"HealthCheckId" name:"HealthCheckId"`
		HttpMethod             *string `json:"HttpMethod" name:"HttpMethod"`
		HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
		Interval               *int    `json:"Interval" name:"Interval"`
		Timeout                *int    `json:"Timeout" name:"Timeout"`
		UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
		UrlPath                *string `json:"UrlPath" name:"UrlPath"`
		HostName               *string `json:"HostName" name:"HostName"`
		HealthCheckReq         *string `json:"HealthCheckReq" name:"HealthCheckReq"`
		HealthCheckExp         *string `json:"HealthCheckExp" name:"HealthCheckExp"`
		HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
	} `json:"HealthCheck"`
	Session struct {
		SessionState             *string `json:"SessionState" name:"SessionState"`
		SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
		CookieType               *string `json:"CookieType" name:"CookieType"`
		CookieName               *string `json:"CookieName" name:"CookieName"`
	} `json:"Session"`
}

func (r *CreateListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesWithListenerRequest struct {
	*ksyunhttp.BaseRequest
	RegisterId      *string `json:"RegisterId,omitempty" name:"RegisterId"`
	Weight          *int    `json:"Weight,omitempty" name:"Weight"`
	RealServerPort  *int    `json:"RealServerPort,omitempty" name:"RealServerPort"`
	MasterSlaveType *string `json:"MasterSlaveType,omitempty" name:"MasterSlaveType"`
	Tag             *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *ModifyInstancesWithListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesWithListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstancesWithListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesWithListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	RegisterId      *string `json:"RegisterId" name:"RegisterId"`
	RealServerState *string `json:"RealServerState" name:"RealServerState"`
	RealServerType  *string `json:"RealServerType" name:"RealServerType"`
	ListenerId      *string `json:"ListenerId" name:"ListenerId"`
	Weight          *int    `json:"Weight" name:"Weight"`
	RealServerIp    *string `json:"RealServerIp" name:"RealServerIp"`
	RealServerPort  *int    `json:"RealServerPort" name:"RealServerPort"`
	Tag             *string `json:"Tag" name:"Tag"`
	MasterSlaveType *string `json:"MasterSlaveType" name:"MasterSlaveType"`
}

func (r *ModifyInstancesWithListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesWithListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterInstancesWithListenerRequest struct {
	*ksyunhttp.BaseRequest
	RealServerType     *string `json:"RealServerType,omitempty" name:"RealServerType"`
	ListenerId         *string `json:"ListenerId,omitempty" name:"ListenerId"`
	Weight             *int    `json:"Weight,omitempty" name:"Weight"`
	RealServerIp       *string `json:"RealServerIp,omitempty" name:"RealServerIp"`
	RealServerPort     *int    `json:"RealServerPort,omitempty" name:"RealServerPort"`
	InstanceId         *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Tag                *string `json:"Tag,omitempty" name:"Tag"`
	MasterSlaveType    *string `json:"MasterSlaveType,omitempty" name:"MasterSlaveType"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *RegisterInstancesWithListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterInstancesWithListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RegisterInstancesWithListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterInstancesWithListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	RegisterId         *string `json:"RegisterId" name:"RegisterId"`
	RealServerState    *string `json:"RealServerState" name:"RealServerState"`
	RealServerType     *string `json:"RealServerType" name:"RealServerType"`
	ListenerId         *string `json:"ListenerId" name:"ListenerId"`
	Weight             *int    `json:"Weight" name:"Weight"`
	RealServerIp       *string `json:"RealServerIp" name:"RealServerIp"`
	RealServerPort     *int    `json:"RealServerPort" name:"RealServerPort"`
	Tag                *string `json:"Tag" name:"Tag"`
	MasterSlaveType    *string `json:"MasterSlaveType" name:"MasterSlaveType"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
}

func (r *RegisterInstancesWithListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterInstancesWithListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterInstancesFromListenerRequest struct {
	*ksyunhttp.BaseRequest
	RegisterId *string `json:"RegisterId,omitempty" name:"RegisterId"`
}

func (r *DeregisterInstancesFromListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterInstancesFromListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeregisterInstancesFromListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterInstancesFromListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeregisterInstancesFromListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterInstancesFromListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesWithListenerRequest struct {
	*ksyunhttp.BaseRequest
	RegisterId []*string                              `json:"RegisterId,omitempty" name:"RegisterId"`
	Filter     []*DescribeInstancesWithListenerFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                                   `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                                `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeInstancesWithListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesWithListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstancesWithListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesWithListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	NextToken     *string `json:"NextToken" name:"NextToken"`
	RealServerSet []struct {
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		RegisterId         *string `json:"RegisterId" name:"RegisterId"`
		RealServerState    *string `json:"RealServerState" name:"RealServerState"`
		RealServerType     *string `json:"RealServerType" name:"RealServerType"`
		ListenerId         *string `json:"ListenerId" name:"ListenerId"`
		Weight             *int    `json:"Weight" name:"Weight"`
		RealServerIp       *string `json:"RealServerIp" name:"RealServerIp"`
		RealServerPort     *int    `json:"RealServerPort" name:"RealServerPort"`
		InstanceId         *string `json:"InstanceId" name:"InstanceId"`
		Tag                *string `json:"Tag" name:"Tag"`
		MasterSlaveType    *string `json:"MasterSlaveType" name:"MasterSlaveType"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	} `json:"RealServerSet"`
}

func (r *DescribeInstancesWithListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesWithListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHealthCheckRequest struct {
	*ksyunhttp.BaseRequest
	HealthCheckId          *string `json:"HealthCheckId,omitempty" name:"HealthCheckId"`
	HealthCheckState       *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	HealthCheckConnectPort *int    `json:"HealthCheckConnectPort,omitempty" name:"HealthCheckConnectPort"`
	HealthyThreshold       *int    `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	Interval               *int    `json:"Interval,omitempty" name:"Interval"`
	Timeout                *int    `json:"Timeout,omitempty" name:"Timeout"`
	UnhealthyThreshold     *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	HealthProtocol         *string `json:"HealthProtocol,omitempty" name:"HealthProtocol"`
	HttpMethod             *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	UrlPath                *string `json:"UrlPath,omitempty" name:"UrlPath"`
	HostName               *string `json:"HostName,omitempty" name:"HostName"`
	HealthCheckReq         *string `json:"HealthCheckReq,omitempty" name:"HealthCheckReq"`
	HealthCheckExp         *string `json:"HealthCheckExp,omitempty" name:"HealthCheckExp"`
}

func (r *ModifyHealthCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHealthCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyHealthCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHealthCheckResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
	HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
	HealthCheckId          *string `json:"HealthCheckId" name:"HealthCheckId"`
	HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
	Interval               *int    `json:"Interval" name:"Interval"`
	Timeout                *int    `json:"Timeout" name:"Timeout"`
	UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
	UrlPath                *string `json:"UrlPath" name:"UrlPath"`
	HostName               *string `json:"HostName" name:"HostName"`
	HealthCheckReq         *string `json:"HealthCheckReq" name:"HealthCheckReq"`
	HealthCheckExp         *string `json:"HealthCheckExp" name:"HealthCheckExp"`
	HttpMethod             *string `json:"HttpMethod" name:"HttpMethod"`
}

func (r *ModifyHealthCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHealthCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHealthCheckRequest struct {
	*ksyunhttp.BaseRequest
	HealthCheckId *string `json:"HealthCheckId,omitempty" name:"HealthCheckId"`
}

func (r *DeleteHealthCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHealthCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteHealthCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHealthCheckResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteHealthCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHealthCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHealthChecksRequest struct {
	*ksyunhttp.BaseRequest
	HealthCheckId []*string                     `json:"HealthCheckId,omitempty" name:"HealthCheckId"`
	Filter        []*DescribeHealthChecksFilter `json:"Filter,omitempty" name:"Filter"`
	Limit         *int                          `json:"Limit,omitempty" name:"Limit"`
	Marker        *string                       `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeHealthChecksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthChecksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeHealthChecksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHealthChecksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	HealthCheckSet []struct {
		ListenerId             *string `json:"ListenerId" name:"ListenerId"`
		HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
		HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
		HealthCheckId          *string `json:"HealthCheckId" name:"HealthCheckId"`
		HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
		Interval               *int    `json:"Interval" name:"Interval"`
		Timeout                *int    `json:"Timeout" name:"Timeout"`
		UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
		UrlPath                *string `json:"UrlPath" name:"UrlPath"`
		HostName               *string `json:"HostName" name:"HostName"`
		HealthCheckReq         *string `json:"HealthCheckReq" name:"HealthCheckReq"`
		HealthCheckExp         *string `json:"HealthCheckExp" name:"HealthCheckExp"`
		HttpMethod             *string `json:"HttpMethod" name:"HttpMethod"`
		HealthProtocol         *string `json:"HealthProtocol" name:"HealthProtocol"`
	} `json:"HealthCheckSet"`
	NextToken *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeHealthChecksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthChecksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigureHealthCheckRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId             *string `json:"ListenerId,omitempty" name:"ListenerId"`
	HealthCheckState       *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	HealthCheckConnectPort *int    `json:"HealthCheckConnectPort,omitempty" name:"HealthCheckConnectPort"`
	HealthyThreshold       *int    `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	Interval               *int    `json:"Interval,omitempty" name:"Interval"`
	Timeout                *int    `json:"Timeout,omitempty" name:"Timeout"`
	UnhealthyThreshold     *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	HealthProtocol         *string `json:"HealthProtocol,omitempty" name:"HealthProtocol"`
	HttpMethod             *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	UrlPath                *string `json:"UrlPath,omitempty" name:"UrlPath"`
	HostName               *string `json:"HostName,omitempty" name:"HostName"`
	HealthCheckReq         *string `json:"HealthCheckReq,omitempty" name:"HealthCheckReq"`
	HealthCheckExp         *string `json:"HealthCheckExp,omitempty" name:"HealthCheckExp"`
}

func (r *ConfigureHealthCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigureHealthCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ConfigureHealthCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ConfigureHealthCheckResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
	HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
	HealthCheckId          *string `json:"HealthCheckId" name:"HealthCheckId"`
	HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
	Interval               *int    `json:"Interval" name:"Interval"`
	Timeout                *int    `json:"Timeout" name:"Timeout"`
	UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
	UrlPath                *string `json:"UrlPath" name:"UrlPath"`
	HostName               *string `json:"HostName" name:"HostName"`
	HealthCheckReq         *string `json:"HealthCheckReq" name:"HealthCheckReq"`
	HealthCheckExp         *string `json:"HealthCheckExp" name:"HealthCheckExp"`
	HttpMethod             *string `json:"HttpMethod" name:"HttpMethod"`
}

func (r *ConfigureHealthCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigureHealthCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId      []*string                      `json:"ProjectId,omitempty" name:"ProjectId"`
	LoadBalancerId []*string                      `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	Filter         []*DescribeLoadBalancersFilter `json:"Filter,omitempty" name:"Filter"`
	IsContainTag   *bool                          `json:"IsContainTag,omitempty" name:"IsContainTag"`
	TagKey         []*string                      `json:"TagKey,omitempty" name:"TagKey"`
	TagKV          []*DescribeLoadBalancersTagKV  `json:"TagKV,omitempty" name:"TagKV"`
	MaxResults     *int                           `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken      *string                        `json:"NextToken,omitempty" name:"NextToken"`
	State          *string                        `json:"State,omitempty" name:"State"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	NextToken                *string `json:"NextToken" name:"NextToken"`
	TotalCount               *int    `json:"TotalCount" name:"TotalCount"`
	LoadBalancerDescriptions []struct {
		LoadBalancerId    *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		LoadBalancerName  *string `json:"LoadBalancerName" name:"LoadBalancerName"`
		IsWaf             *bool   `json:"IsWaf" name:"IsWaf"`
		Type              *string `json:"Type" name:"Type"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		ProjectId         *string `json:"ProjectId" name:"ProjectId"`
		VpcId             *string `json:"VpcId" name:"VpcId"`
		ServiceEndTime    *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		PublicIp          *string `json:"PublicIp" name:"PublicIp"`
		State             *string `json:"State" name:"State"`
		IpVersion         *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerState *string `json:"LoadBalancerState" name:"LoadBalancerState"`
		ListenersCount    *int    `json:"ListenersCount" name:"ListenersCount"`
		ChargeType        *string `json:"ChargeType" name:"ChargeType"`
		LbType            *string `json:"LbType" name:"LbType"`
		LbStatus          *string `json:"LbStatus" name:"LbStatus"`
		VnetId            *string `json:"VnetId" name:"VnetId"`
		DeleteProtection  *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection  *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet            []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"LoadBalancerDescriptions"`
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

func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type ModifyLoadBalancerRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId    *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	LoadBalancerName  *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	LoadBalancerState *string `json:"LoadBalancerState,omitempty" name:"LoadBalancerState"`
}

func (r *ModifyLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	LoadBalancerId    *string `json:"LoadBalancerId" name:"LoadBalancerId"`
	LoadBalancerName  *string `json:"LoadBalancerName" name:"LoadBalancerName"`
	Type              *string `json:"Type" name:"Type"`
	CreateTime        *string `json:"CreateTime" name:"CreateTime"`
	VpcId             *string `json:"VpcId" name:"VpcId"`
	PublicIp          *string `json:"PublicIp" name:"PublicIp"`
	IpVersion         *string `json:"IpVersion" name:"IpVersion"`
	LbType            *string `json:"LbType" name:"LbType"`
	LoadBalancerState *string `json:"LoadBalancerState" name:"LoadBalancerState"`
}

func (r *ModifyLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerRequest struct {
	*ksyunhttp.BaseRequest
	VpcId                  *string `json:"VpcId,omitempty" name:"VpcId"`
	LoadBalancerName       *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	Type                   *string `json:"Type,omitempty" name:"Type"`
	SubnetId               *string `json:"SubnetId,omitempty" name:"SubnetId"`
	PrivateIpAddress       *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	DeleteProtection       *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
	ModificationProtection *string `json:"ModificationProtection,omitempty" name:"ModificationProtection"`
	IpVersion              *string `json:"IpVersion,omitempty" name:"IpVersion"`
	LbType                 *string `json:"LbType,omitempty" name:"LbType"`
	ProjectId              *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	LoadBalancerId   *string `json:"LoadBalancerId" name:"LoadBalancerId"`
	LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
	Type             *string `json:"Type" name:"Type"`
	CreateTime       *string `json:"CreateTime" name:"CreateTime"`
	VpcId            *string `json:"VpcId" name:"VpcId"`
	PublicIp         *string `json:"PublicIp" name:"PublicIp"`
	IpVersion        *string `json:"IpVersion" name:"IpVersion"`
	LbType           *string `json:"LbType" name:"LbType"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostHeaderRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId    *string `json:"ListenerId,omitempty" name:"ListenerId"`
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	HostHeader    *string `json:"HostHeader,omitempty" name:"HostHeader"`
}

func (r *CreateHostHeaderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostHeaderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateHostHeaderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostHeaderResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	HostHeader struct {
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		HostHeaderId  *string `json:"HostHeaderId" name:"HostHeaderId"`
		ListenerId    *string `json:"ListenerId" name:"ListenerId"`
		CertificateId *string `json:"CertificateId" name:"CertificateId"`
		HostHeader    *string `json:"HostHeader" name:"HostHeader"`
	} `json:"HostHeader"`
}

func (r *CreateHostHeaderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostHeaderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostHeaderRequest struct {
	*ksyunhttp.BaseRequest
	HostHeaderId  *string `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *ModifyHostHeaderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostHeaderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyHostHeaderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostHeaderResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	HostHeader struct {
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		HostHeaderId  *string `json:"HostHeaderId" name:"HostHeaderId"`
		ListenerId    *string `json:"ListenerId" name:"ListenerId"`
		CertificateId *string `json:"CertificateId" name:"CertificateId"`
		HostHeader    *string `json:"HostHeader" name:"HostHeader"`
	} `json:"HostHeader"`
}

func (r *ModifyHostHeaderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostHeaderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostHeaderRequest struct {
	*ksyunhttp.BaseRequest
	HostHeaderId *string `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
}

func (r *DeleteHostHeaderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostHeaderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteHostHeaderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostHeaderResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteHostHeaderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostHeaderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostHeadersRequest struct {
	*ksyunhttp.BaseRequest
	HostHeaderId []*string                    `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
	Filter       []*DescribeHostHeadersFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults   *int                         `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string                      `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeHostHeadersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostHeadersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeHostHeadersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostHeadersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	NextToken     *string `json:"NextToken" name:"NextToken"`
	HostHeaderSet []struct {
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		HostHeaderId  *string `json:"HostHeaderId" name:"HostHeaderId"`
		ListenerId    *string `json:"ListenerId" name:"ListenerId"`
		CertificateId *string `json:"CertificateId" name:"CertificateId"`
		HostHeader    *string `json:"HostHeader" name:"HostHeader"`
	} `json:"HostHeaderSet"`
}

func (r *DescribeHostHeadersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostHeadersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type DescribeRulesRequest struct {
	*ksyunhttp.BaseRequest
	RuleId     []*string              `json:"RuleId,omitempty" name:"RuleId"`
	Filter     []*DescribeRulesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                   `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RuleSet   []struct {
		Method               *string `json:"Method" name:"Method"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Path                 *string `json:"Path" name:"Path"`
		RuleId               *string `json:"RuleId" name:"RuleId"`
		ListenerSync         *string `json:"ListenerSync" name:"ListenerSync"`
		HostHeaderId         *string `json:"HostHeaderId" name:"HostHeaderId"`
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		HealthCheck          struct {
			HostName           *string `json:"HostName" name:"HostName"`
			HealthCheckState   *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold   *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval           *int    `json:"Interval" name:"Interval"`
			Timeout            *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath            *string `json:"UrlPath" name:"UrlPath"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
		BackendServerSet []struct {
			BackendServerIp    *string `json:"BackendServerIp" name:"BackendServerIp"`
			RegisterId         *string `json:"RegisterId" name:"RegisterId"`
			BackendServerPort  *int    `json:"BackendServerPort" name:"BackendServerPort"`
			BackendServerState *string `json:"BackendServerState" name:"BackendServerState"`
		} `json:"BackendServerSet" name:"BackendServerSet"`
	} `json:"RuleSet"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackendServerGroupRequest struct {
	*ksyunhttp.BaseRequest
	VpcId                  *string `json:"VpcId,omitempty" name:"VpcId"`
	Protocol               *string `json:"Protocol,omitempty" name:"Protocol"`
	BackendServerGroupName *string `json:"BackendServerGroupName,omitempty" name:"BackendServerGroupName"`
	BackendServerGroupType *string `json:"BackendServerGroupType,omitempty" name:"BackendServerGroupType"`
	HostName               *string `json:"HostName,omitempty" name:"HostName"`
	HealthCheckState       *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	HealthyThreshold       *int    `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	Interval               *int    `json:"Interval,omitempty" name:"Interval"`
	Timeout                *int    `json:"Timeout,omitempty" name:"Timeout"`
	UnhealthyThreshold     *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	UrlPath                *string `json:"UrlPath,omitempty" name:"UrlPath"`
	Region                 *string `json:"Region,omitempty" name:"Region"`
	UpstreamKeepalive      *string `json:"UpstreamKeepalive,omitempty" name:"UpstreamKeepalive"`
}

func (r *CreateBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackendServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateBackendServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BackendServerGroup struct {
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		VpcId                  *string `json:"VpcId" name:"VpcId"`
		BackendServerGroupName *string `json:"BackendServerGroupName" name:"BackendServerGroupName"`
		BackendServerNumber    *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive      *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		HealthCheck            struct {
			HostName           *string `json:"HostName" name:"HostName"`
			HealthCheckState   *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold   *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval           *int    `json:"Interval" name:"Interval"`
			Timeout            *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath            *string `json:"UrlPath" name:"UrlPath"`
		} `json:"HealthCheck" name:"HealthCheck"`
	} `json:"BackendServerGroup"`
}

func (r *CreateBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackendServerGroupResponse) FromJsonString(s string) error {
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

func (r *DeleteBackendServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteBackendServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type ModifyBackendServerGroupRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId   *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	BackendServerGroupName *string `json:"BackendServerGroupName,omitempty" name:"BackendServerGroupName"`
	UpstreamKeepalive      *string `json:"UpstreamKeepalive,omitempty" name:"UpstreamKeepalive"`
}

func (r *ModifyBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackendServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyBackendServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BackendServerGroup struct {
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		VpcId                  *string `json:"VpcId" name:"VpcId"`
		BackendServerGroupName *string `json:"BackendServerGroupName" name:"BackendServerGroupName"`
		BackendServerNumber    *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive      *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		HealthCheck            struct {
			HostName           *string `json:"HostName" name:"HostName"`
			HealthCheckState   *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold   *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval           *int    `json:"Interval" name:"Interval"`
			Timeout            *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath            *string `json:"UrlPath" name:"UrlPath"`
		} `json:"HealthCheck" name:"HealthCheck"`
	} `json:"BackendServerGroup"`
}

func (r *ModifyBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackendServerGroupsRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId []*string                            `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	Filter               []*DescribeBackendServerGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults           *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken            *string                              `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeBackendServerGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackendServerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBackendServerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackendServerGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	NextToken             *string `json:"NextToken" name:"NextToken"`
	BackendServerGroupSet []struct {
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		VpcId                  *string `json:"VpcId" name:"VpcId"`
		Protocol               *string `json:"Protocol" name:"Protocol"`
		BackendServerGroupName *string `json:"BackendServerGroupName" name:"BackendServerGroupName"`
		BackendServerNumber    *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		HealthCheck            struct {
			HostName           *string `json:"HostName" name:"HostName"`
			HealthCheckState   *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold   *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval           *int    `json:"Interval" name:"Interval"`
			Timeout            *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath            *string `json:"UrlPath" name:"UrlPath"`
		} `json:"HealthCheck" name:"HealthCheck"`
		IpVersion         *string `json:"IpVersion" name:"IpVersion"`
		Type              *string `json:"Type" name:"Type"`
		UpstreamKeepalive *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
	} `json:"BackendServerGroupSet"`
}

func (r *DescribeBackendServerGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackendServerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	BackendServerIp      *string `json:"BackendServerIp,omitempty" name:"BackendServerIp"`
	BackendServerPort    *int    `json:"BackendServerPort,omitempty" name:"BackendServerPort"`
	Weight               *int    `json:"Weight,omitempty" name:"Weight"`
}

func (r *RegisterBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBackendServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RegisterBackendServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	BackendServer struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerIp      *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		RegisterId           *string `json:"RegisterId" name:"RegisterId"`
		BackendServerPort    *int    `json:"BackendServerPort" name:"BackendServerPort"`
		Weight               *int    `json:"Weight" name:"Weight"`
	} `json:"BackendServer"`
}

func (r *RegisterBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	RegisterId *string `json:"RegisterId,omitempty" name:"RegisterId"`
}

func (r *DeregisterBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterBackendServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeregisterBackendServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type DescribeBackendServersRequest struct {
	*ksyunhttp.BaseRequest
	RegisterId []*string                       `json:"RegisterId,omitempty" name:"RegisterId"`
	Filter     []*DescribeBackendServersFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                            `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                         `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeBackendServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackendServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBackendServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackendServersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	NextToken        *string `json:"NextToken" name:"NextToken"`
	BackendServerSet []struct {
	} `json:"BackendServerSet"`
}

func (r *DescribeBackendServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackendServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerAclRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclName *string `json:"LoadBalancerAclName,omitempty" name:"LoadBalancerAclName"`
	IpVersion           *string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *CreateLoadBalancerAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateLoadBalancerAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	LoadBalancerAcl struct {
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerAclName     *string `json:"LoadBalancerAclName" name:"LoadBalancerAclName"`
		LoadBalancerAclId       *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
		IpVersion               *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerAclEntrySet []struct {
			LoadBalancerAclId      *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
			LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId" name:"LoadBalancerAclEntryId"`
			CidrBlock              *string `json:"CidrBlock" name:"CidrBlock"`
			RuleNumber             *int    `json:"RuleNumber" name:"RuleNumber"`
			RuleAction             *string `json:"RuleAction" name:"RuleAction"`
			Protocol               *string `json:"Protocol" name:"Protocol"`
			Description            *string `json:"Description" name:"Description"`
		} `json:"LoadBalancerAclEntrySet" name:"LoadBalancerAclEntrySet"`
	} `json:"LoadBalancerAcl"`
}

func (r *CreateLoadBalancerAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerAclRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclId *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
}

func (r *DeleteLoadBalancerAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteLoadBalancerAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAclRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclId   *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
	LoadBalancerAclName *string `json:"LoadBalancerAclName,omitempty" name:"LoadBalancerAclName"`
}

func (r *ModifyLoadBalancerAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	LoadBalancerAcl struct {
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerAclName     *string `json:"LoadBalancerAclName" name:"LoadBalancerAclName"`
		LoadBalancerAclId       *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
		IpVersion               *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerAclEntrySet []struct {
			LoadBalancerAclId      *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
			LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId" name:"LoadBalancerAclEntryId"`
			CidrBlock              *string `json:"CidrBlock" name:"CidrBlock"`
			RuleNumber             *int    `json:"RuleNumber" name:"RuleNumber"`
			RuleAction             *string `json:"RuleAction" name:"RuleAction"`
			Protocol               *string `json:"Protocol" name:"Protocol"`
			Description            *string `json:"Description" name:"Description"`
		} `json:"LoadBalancerAclEntrySet" name:"LoadBalancerAclEntrySet"`
	} `json:"LoadBalancerAcl"`
}

func (r *ModifyLoadBalancerAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerAclEntryRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclId *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
	CidrBlock         *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	RuleNumber        *int    `json:"RuleNumber,omitempty" name:"RuleNumber"`
	RuleAction        *string `json:"RuleAction,omitempty" name:"RuleAction"`
	Protocol          *string `json:"Protocol,omitempty" name:"Protocol"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateLoadBalancerAclEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerAclEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateLoadBalancerAclEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerAclEntryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	LoadBalancerAclEntry struct {
		LoadBalancerAclId      *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
		LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId" name:"LoadBalancerAclEntryId"`
		CidrBlock              *string `json:"CidrBlock" name:"CidrBlock"`
		RuleNumber             *int    `json:"RuleNumber" name:"RuleNumber"`
		RuleAction             *string `json:"RuleAction" name:"RuleAction"`
		Protocol               *string `json:"Protocol" name:"Protocol"`
		Description            *string `json:"Description" name:"Description"`
	} `json:"LoadBalancerAclEntry"`
}

func (r *CreateLoadBalancerAclEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerAclEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerAclEntryRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId,omitempty" name:"LoadBalancerAclEntryId"`
}

func (r *DeleteLoadBalancerAclEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerAclEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerAclEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerAclEntryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteLoadBalancerAclEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerAclEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateLoadBalancerAclRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclId *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
	ListenerId        *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *AssociateLoadBalancerAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateLoadBalancerAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateLoadBalancerAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateLoadBalancerAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateLoadBalancerAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateLoadBalancerAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateLoadBalancerAclRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DisassociateLoadBalancerAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateLoadBalancerAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateLoadBalancerAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateLoadBalancerAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DisassociateLoadBalancerAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateLoadBalancerAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerAclsRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclId []*string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
	MaxResults        *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken         *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeLoadBalancerAclsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancerAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerAclsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	LoadBalancerAclSet []struct {
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		LoadBalancerAclName     *string `json:"LoadBalancerAclName" name:"LoadBalancerAclName"`
		LoadBalancerAclId       *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
		IpVersion               *string `json:"IpVersion" name:"IpVersion"`
		LoadBalancerAclEntrySet []struct {
			LoadBalancerAclId      *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
			LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId" name:"LoadBalancerAclEntryId"`
			CidrBlock              *string `json:"CidrBlock" name:"CidrBlock"`
			RuleNumber             *int    `json:"RuleNumber" name:"RuleNumber"`
			RuleAction             *string `json:"RuleAction" name:"RuleAction"`
			Protocol               *string `json:"Protocol" name:"Protocol"`
			Description            *string `json:"Description" name:"Description"`
		} `json:"LoadBalancerAclEntrySet" name:"LoadBalancerAclEntrySet"`
	} `json:"LoadBalancerAclSet"`
	NextToken *string `json:"NextToken" name:"NextToken"`
}

func (r *DescribeLoadBalancerAclsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancerAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSlbRuleRequest struct {
	*ksyunhttp.BaseRequest
	Path                     *string `json:"Path,omitempty" name:"Path"`
	HostHeaderId             *string `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
	BackendServerGroupId     *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	ListenerSync             *string `json:"ListenerSync,omitempty" name:"ListenerSync"`
	Method                   *string `json:"Method,omitempty" name:"Method"`
	SessionState             *string `json:"SessionState,omitempty" name:"SessionState"`
	SessionPersistencePeriod *int    `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
	CookieType               *string `json:"cookieType,omitempty" name:"cookieType"`
	CookieName               *string `json:"CookieName,omitempty" name:"CookieName"`
	HealthCheckState         *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	HealthyThreshold         *string `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	Interval                 *int    `json:"Interval,omitempty" name:"Interval"`
	Timeout                  *int    `json:"Timeout,omitempty" name:"Timeout"`
	UnhealthyThreshold       *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	UrlPath                  *string `json:"UrlPath,omitempty" name:"UrlPath"`
	HostName                 *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *CreateSlbRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSlbRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSlbRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSlbRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Rule      struct {
		Method               *string `json:"Method" name:"Method"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Path                 *string `json:"Path" name:"Path"`
		RuleId               *string `json:"RuleId" name:"RuleId"`
		ListenerSync         *string `json:"ListenerSync" name:"ListenerSync"`
		HostHeaderId         *string `json:"HostHeaderId" name:"HostHeaderId"`
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		HealthCheck          struct {
			HostName           *string `json:"HostName" name:"HostName"`
			HealthCheckState   *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold   *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval           *int    `json:"Interval" name:"Interval"`
			Timeout            *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath            *string `json:"UrlPath" name:"UrlPath"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
		BackendServerSet []struct {
			BackendServerIp    *string `json:"BackendServerIp" name:"BackendServerIp"`
			RegisterId         *string `json:"RegisterId" name:"RegisterId"`
			BackendServerPort  *int    `json:"BackendServerPort" name:"BackendServerPort"`
			BackendServerState *string `json:"BackendServerState" name:"BackendServerState"`
		} `json:"BackendServerSet" name:"BackendServerSet"`
	} `json:"Rule"`
}

func (r *CreateSlbRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSlbRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySlbRuleRequest struct {
	*ksyunhttp.BaseRequest
	Path                     *string `json:"Path,omitempty" name:"Path"`
	RuleId                   *string `json:"RuleId,omitempty" name:"RuleId"`
	BackendServerGroupId     *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	ListenerSync             *string `json:"ListenerSync,omitempty" name:"ListenerSync"`
	Method                   *string `json:"Method,omitempty" name:"Method"`
	SessionState             *string `json:"SessionState,omitempty" name:"SessionState"`
	SessionPersistencePeriod *int    `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
	CookieType               *string `json:"cookieType,omitempty" name:"cookieType"`
	CookieName               *string `json:"CookieName,omitempty" name:"CookieName"`
	HealthCheckState         *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
	HealthyThreshold         *string `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
	Interval                 *int    `json:"Interval,omitempty" name:"Interval"`
	Timeout                  *int    `json:"Timeout,omitempty" name:"Timeout"`
	UnhealthyThreshold       *int    `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
	UrlPath                  *string `json:"UrlPath,omitempty" name:"UrlPath"`
	HostName                 *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *ModifySlbRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySlbRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySlbRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySlbRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Rule      struct {
		Method               *string `json:"Method" name:"Method"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Path                 *string `json:"Path" name:"Path"`
		RuleId               *string `json:"RuleId" name:"RuleId"`
		ListenerSync         *string `json:"ListenerSync" name:"ListenerSync"`
		HostHeaderId         *string `json:"HostHeaderId" name:"HostHeaderId"`
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		HealthCheck          struct {
			HostName           *string `json:"HostName" name:"HostName"`
			HealthCheckState   *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold   *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval           *int    `json:"Interval" name:"Interval"`
			Timeout            *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			UrlPath            *string `json:"UrlPath" name:"UrlPath"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
		BackendServerSet []struct {
			BackendServerIp    *string `json:"BackendServerIp" name:"BackendServerIp"`
			RegisterId         *string `json:"RegisterId" name:"RegisterId"`
			BackendServerPort  *int    `json:"BackendServerPort" name:"BackendServerPort"`
			BackendServerState *string `json:"BackendServerState" name:"BackendServerState"`
		} `json:"BackendServerSet" name:"BackendServerSet"`
	} `json:"Rule"`
}

func (r *ModifySlbRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySlbRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateLinkServerRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerName *string `json:"PrivateLinkServerName,omitempty" name:"PrivateLinkServerName"`
	ListenerId            *string `json:"ListenerId,omitempty" name:"ListenerId"`
	Description           *string `json:"Description,omitempty" name:"Description"`
	ProjectId             *string `json:"ProjectId,omitempty" name:"ProjectId"`
	DeleteProtection      *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
}

func (r *CreatePrivateLinkServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateLinkServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatePrivateLinkServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateLinkServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	PrivateLinkServer struct {
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		PrivateLinkServerName *string `json:"PrivateLinkServerName" name:"PrivateLinkServerName"`
		PrivateLinkServerId   *string `json:"PrivateLinkServerId" name:"PrivateLinkServerId"`
		ListenerId            *string `json:"ListenerId" name:"ListenerId"`
		Description           *string `json:"Description" name:"Description"`
		ProjectId             *string `json:"ProjectId" name:"ProjectId"`
		PrivateLinkNum        *int    `json:"PrivateLinkNum" name:"PrivateLinkNum"`
		ServiceEndTime        *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		DeleteProtection      *string `json:"DeleteProtection" name:"DeleteProtection"`
	} `json:"PrivateLinkServer"`
}

func (r *CreatePrivateLinkServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateLinkServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateLinkServerRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerId []*string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
	ProjectId           []*string `json:"ProjectId,omitempty" name:"ProjectId"`
	MaxResults          *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken           *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribePrivateLinkServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateLinkServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePrivateLinkServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateLinkServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	NextToken            *string `json:"NextToken" name:"NextToken"`
	TotalCount           *int    `json:"TotalCount" name:"TotalCount"`
	PrivateLinkServerSet []struct {
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		PrivateLinkServerName *string `json:"PrivateLinkServerName" name:"PrivateLinkServerName"`
		PrivateLinkServerId   *string `json:"PrivateLinkServerId" name:"PrivateLinkServerId"`
		ListenerId            *string `json:"ListenerId" name:"ListenerId"`
		Description           *string `json:"Description" name:"Description"`
		ProjectId             *string `json:"ProjectId" name:"ProjectId"`
		PrivateLinkNum        *int    `json:"PrivateLinkNum" name:"PrivateLinkNum"`
		ServiceEndTime        *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		DeleteProtection      *string `json:"DeleteProtection" name:"DeleteProtection"`
	} `json:"PrivateLinkServerSet"`
}

func (r *DescribePrivateLinkServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateLinkServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateLinkServerRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerId *string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
}

func (r *DeletePrivateLinkServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateLinkServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePrivateLinkServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateLinkServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeletePrivateLinkServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateLinkServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateLinkServerRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerId   *string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
	PrivateLinkServerName *string `json:"PrivateLinkServerName,omitempty" name:"PrivateLinkServerName"`
	Description           *string `json:"Description,omitempty" name:"Description"`
	DeleteProtection      *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
}

func (r *ModifyPrivateLinkServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateLinkServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyPrivateLinkServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateLinkServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	PrivateLinkServer struct {
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		PrivateLinkServerName *string `json:"PrivateLinkServerName" name:"PrivateLinkServerName"`
		PrivateLinkServerId   *string `json:"PrivateLinkServerId" name:"PrivateLinkServerId"`
		ListenerId            *string `json:"ListenerId" name:"ListenerId"`
		Description           *string `json:"Description" name:"Description"`
		ProjectId             *string `json:"ProjectId" name:"ProjectId"`
		PrivateLinkNum        *int    `json:"PrivateLinkNum" name:"PrivateLinkNum"`
		ServiceEndTime        *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		DeleteProtection      *string `json:"DeleteProtection" name:"DeleteProtection"`
	} `json:"PrivateLinkServer"`
}

func (r *ModifyPrivateLinkServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateLinkServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociatePrivateLinkServerRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerId *string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
	LoadBalancerId      *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ListenerPort        *int    `json:"ListenerPort,omitempty" name:"ListenerPort"`
	ProjectId           *string `json:"ProjectId,omitempty" name:"ProjectId"`
	DeleteProtection    *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
}

func (r *AssociatePrivateLinkServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociatePrivateLinkServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociatePrivateLinkServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociatePrivateLinkServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	PrivateLink struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		PrivateLinkId       *string `json:"PrivateLinkId" name:"PrivateLinkId"`
		PrivateLinkServerId *string `json:"PrivateLinkServerId" name:"PrivateLinkServerId"`
		AccountId           *string `json:"AccountId" name:"AccountId"`
		ListenerId          *string `json:"ListenerId" name:"ListenerId"`
		ServiceAccountId    *string `json:"ServiceAccountId" name:"ServiceAccountId"`
		UpdateTime          *string `json:"UpdateTime" name:"UpdateTime"`
		ProjectId           *string `json:"ProjectId" name:"ProjectId"`
		ConnectionStatus    *string `json:"ConnectionStatus" name:"ConnectionStatus"`
		LoadBalancerId      *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		ListenerPort        *int    `json:"ListenerPort" name:"ListenerPort"`
		ServiceEndTime      *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	} `json:"PrivateLink"`
}

func (r *AssociatePrivateLinkServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociatePrivateLinkServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateLinkRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkId []*string `json:"PrivateLinkId,omitempty" name:"PrivateLinkId"`
	ProjectId     []*string `json:"ProjectId,omitempty" name:"ProjectId"`
	MaxResults    *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribePrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribePrivateLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateLinkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	NextToken      *string `json:"NextToken" name:"NextToken"`
	TotalCount     *int    `json:"TotalCount" name:"TotalCount"`
	PrivateLinkSet []struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		PrivateLinkId       *string `json:"PrivateLinkId" name:"PrivateLinkId"`
		PrivateLinkServerId *string `json:"PrivateLinkServerId" name:"PrivateLinkServerId"`
		AccountId           *string `json:"AccountId" name:"AccountId"`
		ListenerId          *string `json:"ListenerId" name:"ListenerId"`
		ServiceAccountId    *string `json:"ServiceAccountId" name:"ServiceAccountId"`
		UpdateTime          *string `json:"UpdateTime" name:"UpdateTime"`
		ProjectId           *string `json:"ProjectId" name:"ProjectId"`
		ConnectionStatus    *string `json:"ConnectionStatus" name:"ConnectionStatus"`
		LoadBalancerId      *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		ListenerPort        *int    `json:"ListenerPort" name:"ListenerPort"`
		ServiceEndTime      *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	} `json:"PrivateLinkSet"`
}

func (r *DescribePrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateLinkRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkId *string `json:"PrivateLinkId,omitempty" name:"PrivateLinkId"`
}

func (r *DeletePrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeletePrivateLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateLinkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeletePrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAclEntryRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId,omitempty" name:"LoadBalancerAclEntryId"`
	RuleNumber             *int    `json:"RuleNumber,omitempty" name:"RuleNumber"`
	RuleAction             *string `json:"RuleAction,omitempty" name:"RuleAction"`
	Description            *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyLoadBalancerAclEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerAclEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAclEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAclEntryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	LoadBalancerAclEntry struct {
		LoadBalancerAclId      *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
		LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId" name:"LoadBalancerAclEntryId"`
		CidrBlock              *string `json:"CidrBlock" name:"CidrBlock"`
		RuleNumber             *int    `json:"RuleNumber" name:"RuleNumber"`
		RuleAction             *string `json:"RuleAction" name:"RuleAction"`
		Protocol               *string `json:"Protocol" name:"Protocol"`
		Description            *string `json:"Description" name:"Description"`
	} `json:"LoadBalancerAclEntry"`
}

func (r *ModifyLoadBalancerAclEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerAclEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptPrivateLinkRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkId *string `json:"PrivateLinkId,omitempty" name:"PrivateLinkId"`
}

func (r *AcceptPrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptPrivateLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AcceptPrivateLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AcceptPrivateLinkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AcceptPrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptPrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectPrivateLinkRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkId *string `json:"PrivateLinkId,omitempty" name:"PrivateLinkId"`
}

func (r *RejectPrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectPrivateLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RejectPrivateLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RejectPrivateLinkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RejectPrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectPrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPrivateLinkServerRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerId *string                        `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
	Filter              []*ListPrivateLinkServerFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListPrivateLinkServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPrivateLinkServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListPrivateLinkServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListPrivateLinkServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	PrivateLinkSet struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		PrivateLinkId       *string `json:"PrivateLinkId" name:"PrivateLinkId"`
		PrivateLinkServerId *string `json:"PrivateLinkServerId" name:"PrivateLinkServerId"`
		AccountId           *string `json:"AccountId" name:"AccountId"`
		ListenerId          *string `json:"ListenerId" name:"ListenerId"`
		ServiceAccountId    *string `json:"ServiceAccountId" name:"ServiceAccountId"`
		UpdateTime          *string `json:"UpdateTime" name:"UpdateTime"`
		ProjectId           *string `json:"ProjectId" name:"ProjectId"`
		ConnectionStatus    *string `json:"ConnectionStatus" name:"ConnectionStatus"`
		LoadBalancerId      *string `json:"LoadBalancerId" name:"LoadBalancerId"`
		ListenerPort        *int    `json:"ListenerPort" name:"ListenerPort"`
		ServiceEndTime      *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	} `json:"PrivateLinkSet"`
}

func (r *ListPrivateLinkServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPrivateLinkServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemovePrivateLinkRequest struct {
	*ksyunhttp.BaseRequest
	PrivateLinkServerId *string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
	PrivateLinkId       *string `json:"PrivateLinkId,omitempty" name:"PrivateLinkId"`
}

func (r *RemovePrivateLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemovePrivateLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RemovePrivateLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemovePrivateLinkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RemovePrivateLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemovePrivateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbRequest struct {
	*ksyunhttp.BaseRequest
	AlbName                *string `json:"AlbName,omitempty" name:"AlbName"`
	AlbVersion             *string `json:"AlbVersion,omitempty" name:"AlbVersion"`
	AlbType                *string `json:"AlbType,omitempty" name:"AlbType"`
	VpcId                  *string `json:"VpcId,omitempty" name:"VpcId"`
	IpVersion              *string `json:"IpVersion,omitempty" name:"IpVersion"`
	ProjectId              *string `json:"ProjectId,omitempty" name:"ProjectId"`
	AllocationId           *string `json:"AllocationId,omitempty" name:"AllocationId"`
	ChargeType             *string `json:"ChargeType,omitempty" name:"ChargeType"`
	SubnetId               *string `json:"SubnetId,omitempty" name:"SubnetId"`
	PrivateIpAddress       *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	EnabledQuic            *bool   `json:"EnabledQuic,omitempty" name:"EnabledQuic"`
	EnableHpa              *bool   `json:"EnableHpa,omitempty" name:"EnableHpa"`
	DeleteProtection       *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
	ModificationProtection *string `json:"ModificationProtection,omitempty" name:"ModificationProtection"`
}

func (r *CreateAlbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAlbRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	ApplicationLoadBalancer struct {
		AlbId          *string `json:"AlbId" name:"AlbId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		AlbName        *string `json:"AlbName" name:"AlbName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		AlbVersion     *string `json:"AlbVersion" name:"AlbVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		AlbType        *string `json:"AlbType" name:"AlbType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		KlogInfo       struct {
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			LogpoolName *string `json:"LogpoolName" name:"LogpoolName"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
		} `json:"KlogInfo" name:"KlogInfo"`
		BillType         *int    `json:"BillType" name:"BillType"`
		ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic      *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa        *bool   `json:"EnableHpa" name:"EnableHpa"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet           []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"ApplicationLoadBalancer"`
}

func (r *CreateAlbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbRequest struct {
	*ksyunhttp.BaseRequest
	AlbId *string `json:"AlbId,omitempty" name:"AlbId"`
}

func (r *DeleteAlbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlbRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteAlbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbNameRequest struct {
	*ksyunhttp.BaseRequest
	AlbId   *string `json:"AlbId,omitempty" name:"AlbId"`
	AlbName *string `json:"AlbName,omitempty" name:"AlbName"`
}

func (r *SetAlbNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetAlbNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	ApplicationLoadBalancer struct {
		AlbId          *string `json:"AlbId" name:"AlbId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		AlbName        *string `json:"AlbName" name:"AlbName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		AlbVersion     *string `json:"AlbVersion" name:"AlbVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		AlbType        *string `json:"AlbType" name:"AlbType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		KlogInfo       struct {
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			LogpoolName *string `json:"LogpoolName" name:"LogpoolName"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
		} `json:"KlogInfo" name:"KlogInfo"`
		BillType         *int    `json:"BillType" name:"BillType"`
		ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic      *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa        *bool   `json:"EnableHpa" name:"EnableHpa"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet           []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"ApplicationLoadBalancer"`
}

func (r *SetAlbNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbStatusRequest struct {
	*ksyunhttp.BaseRequest
	AlbId *string `json:"AlbId,omitempty" name:"AlbId"`
	State *string `json:"State,omitempty" name:"State"`
}

func (r *SetAlbStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetAlbStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	ApplicationLoadBalancer struct {
		AlbId          *string `json:"AlbId" name:"AlbId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		AlbName        *string `json:"AlbName" name:"AlbName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		AlbVersion     *string `json:"AlbVersion" name:"AlbVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		AlbType        *string `json:"AlbType" name:"AlbType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		KlogInfo       struct {
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			LogpoolName *string `json:"LogpoolName" name:"LogpoolName"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
		} `json:"KlogInfo" name:"KlogInfo"`
		BillType         *int    `json:"BillType" name:"BillType"`
		ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic      *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa        *bool   `json:"EnableHpa" name:"EnableHpa"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet           []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"ApplicationLoadBalancer"`
}

func (r *SetAlbStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbsRequest struct {
	*ksyunhttp.BaseRequest
	AlbId        []*string             `json:"AlbId,omitempty" name:"AlbId"`
	Filter       []*DescribeAlbsFilter `json:"Filter,omitempty" name:"Filter"`
	IsContainTag *bool                 `json:"IsContainTag,omitempty" name:"IsContainTag"`
	TagKey       []*string             `json:"TagKey,omitempty" name:"TagKey"`
	TagKV        []*DescribeAlbsTagKV  `json:"TagKV,omitempty" name:"TagKV"`
	ProjectId    []*string             `json:"ProjectId,omitempty" name:"ProjectId"`
	MaxResults   *int                  `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string               `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAlbsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlbsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                  *string `json:"RequestId" name:"RequestId"`
	NextToken                  *string `json:"NextToken" name:"NextToken"`
	ApplicationLoadBalancerSet []struct {
		AlbId          *string `json:"AlbId" name:"AlbId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		AlbName        *string `json:"AlbName" name:"AlbName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		AlbVersion     *string `json:"AlbVersion" name:"AlbVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		AlbType        *string `json:"AlbType" name:"AlbType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		KlogInfo       struct {
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			LogpoolName *string `json:"LogpoolName" name:"LogpoolName"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
		} `json:"KlogInfo" name:"KlogInfo"`
		BillType         *int    `json:"BillType" name:"BillType"`
		ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic      *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa        *bool   `json:"EnableHpa" name:"EnableHpa"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet           []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"ApplicationLoadBalancerSet"`
}

func (r *DescribeAlbsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbListenerRequest struct {
	*ksyunhttp.BaseRequest
	AlbId                 *string                               `json:"AlbId,omitempty" name:"AlbId"`
	AlbListenerName       *string                               `json:"AlbListenerName,omitempty" name:"AlbListenerName"`
	Protocol              *string                               `json:"Protocol,omitempty" name:"Protocol"`
	Port                  *int                                  `json:"Port,omitempty" name:"Port"`
	CertificateId         *string                               `json:"CertificateId,omitempty" name:"CertificateId"`
	TlsCipherPolicy       *string                               `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
	AlbListenerAclId      *string                               `json:"AlbListenerAclId,omitempty" name:"AlbListenerAclId"`
	AlbListenerState      *string                               `json:"AlbListenerState,omitempty" name:"AlbListenerState"`
	RedirectAlbListenerId *string                               `json:"RedirectAlbListenerId,omitempty" name:"RedirectAlbListenerId"`
	RedirectHttpCode      *string                               `json:"RedirectHttpCode,omitempty" name:"RedirectHttpCode"`
	EnableHttp2           *bool                                 `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
	BackendServerGroupId  *string                               `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	FixedResponseConfig   *CreateAlbListenerFixedResponseConfig `json:"FixedResponseConfig,omitempty" name:"FixedResponseConfig"`
	RewriteConfig         *CreateAlbListenerRewriteConfig       `json:"RewriteConfig,omitempty" name:"RewriteConfig"`
	CaEnabled             *bool                                 `json:"CaEnabled,omitempty" name:"CaEnabled"`
	CaCertificateId       *string                               `json:"CaCertificateId,omitempty" name:"CaCertificateId"`
	EnableQuicUpgrade     *bool                                 `json:"EnableQuicUpgrade,omitempty" name:"EnableQuicUpgrade"`
	QuicListenerId        *string                               `json:"QuicListenerId,omitempty" name:"QuicListenerId"`
}

func (r *CreateAlbListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAlbListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	AlbListener struct {
		AlbListenerId               *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbId                       *string `json:"AlbId" name:"AlbId"`
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		AlbListenerName             *string `json:"AlbListenerName" name:"AlbListenerName"`
		Protocol                    *string `json:"Protocol" name:"Protocol"`
		Port                        *int    `json:"Port" name:"Port"`
		CertificateId               *string `json:"CertificateId" name:"CertificateId"`
		TlsCipherPolicy             *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		DefaultBackendServerGroupId *string `json:"DefaultBackendServerGroupId" name:"DefaultBackendServerGroupId"`
		AlbListenerAclId            *string `json:"AlbListenerAclId" name:"AlbListenerAclId"`
		AlbListenerState            *string `json:"AlbListenerState" name:"AlbListenerState"`
		RedirectAlbListenerId       *string `json:"RedirectAlbListenerId" name:"RedirectAlbListenerId"`
		HttpProtocol                *string `json:"HttpProtocol" name:"HttpProtocol"`
		EnableHttp2                 *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		CaCertificateId             *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled                   *bool   `json:"CaEnabled" name:"CaEnabled"`
		EnableQuicUpgrade           *bool   `json:"EnableQuicUpgrade" name:"EnableQuicUpgrade"`
		QuicListenerId              *string `json:"QuicListenerId" name:"QuicListenerId"`
	} `json:"AlbListener"`
}

func (r *CreateAlbListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbListenerRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerId     *string `json:"AlbListenerId,omitempty" name:"AlbListenerId"`
	AlbListenerName   *string `json:"AlbListenerName,omitempty" name:"AlbListenerName"`
	AlbListenerState  *string `json:"AlbListenerState,omitempty" name:"AlbListenerState"`
	CertificateId     *string `json:"CertificateId,omitempty" name:"CertificateId"`
	TlsCipherPolicy   *string `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
	AlbListenerAclId  *string `json:"AlbListenerAclId,omitempty" name:"AlbListenerAclId"`
	HttpProtocol      *string `json:"HttpProtocol,omitempty" name:"HttpProtocol"`
	EnableHttp2       *bool   `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
	CaEnabled         *bool   `json:"CaEnabled,omitempty" name:"CaEnabled"`
	CaCertificateId   *string `json:"CaCertificateId,omitempty" name:"CaCertificateId"`
	EnableQuicUpgrade *bool   `json:"EnableQuicUpgrade,omitempty" name:"EnableQuicUpgrade"`
	QuicListenerId    *string `json:"QuicListenerId,omitempty" name:"QuicListenerId"`
}

func (r *ModifyAlbListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyAlbListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	AlbListener struct {
		AlbListenerId               *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbId                       *string `json:"AlbId" name:"AlbId"`
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		AlbListenerName             *string `json:"AlbListenerName" name:"AlbListenerName"`
		Protocol                    *string `json:"Protocol" name:"Protocol"`
		Port                        *int    `json:"Port" name:"Port"`
		CertificateId               *string `json:"CertificateId" name:"CertificateId"`
		TlsCipherPolicy             *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		DefaultBackendServerGroupId *string `json:"DefaultBackendServerGroupId" name:"DefaultBackendServerGroupId"`
		AlbListenerAclId            *string `json:"AlbListenerAclId" name:"AlbListenerAclId"`
		AlbListenerState            *string `json:"AlbListenerState" name:"AlbListenerState"`
		RedirectAlbListenerId       *string `json:"RedirectAlbListenerId" name:"RedirectAlbListenerId"`
		HttpProtocol                *string `json:"HttpProtocol" name:"HttpProtocol"`
		EnableHttp2                 *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		CaCertificateId             *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled                   *bool   `json:"CaEnabled" name:"CaEnabled"`
		EnableQuicUpgrade           *bool   `json:"EnableQuicUpgrade" name:"EnableQuicUpgrade"`
		QuicListenerId              *string `json:"QuicListenerId" name:"QuicListenerId"`
	} `json:"AlbListener"`
}

func (r *ModifyAlbListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbListenerRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerId *string `json:"AlbListenerId,omitempty" name:"AlbListenerId"`
}

func (r *DeleteAlbListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlbListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteAlbListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbListenersRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerId []*string                     `json:"AlbListenerId,omitempty" name:"AlbListenerId"`
	Filter        []*DescribeAlbListenersFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults    *int                          `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string                       `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAlbListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlbListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbListenersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	NextToken      *string `json:"NextToken" name:"NextToken"`
	AlbListenerSet []struct {
		AlbListenerId               *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbId                       *string `json:"AlbId" name:"AlbId"`
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		AlbListenerName             *string `json:"AlbListenerName" name:"AlbListenerName"`
		Protocol                    *string `json:"Protocol" name:"Protocol"`
		Port                        *int    `json:"Port" name:"Port"`
		CertificateId               *string `json:"CertificateId" name:"CertificateId"`
		TlsCipherPolicy             *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
		DefaultBackendServerGroupId *string `json:"DefaultBackendServerGroupId" name:"DefaultBackendServerGroupId"`
		AlbListenerAclId            *string `json:"AlbListenerAclId" name:"AlbListenerAclId"`
		AlbListenerState            *string `json:"AlbListenerState" name:"AlbListenerState"`
		RedirectAlbListenerId       *string `json:"RedirectAlbListenerId" name:"RedirectAlbListenerId"`
		HttpProtocol                *string `json:"HttpProtocol" name:"HttpProtocol"`
		EnableHttp2                 *bool   `json:"EnableHttp2" name:"EnableHttp2"`
		CaCertificateId             *string `json:"CaCertificateId" name:"CaCertificateId"`
		CaEnabled                   *bool   `json:"CaEnabled" name:"CaEnabled"`
		EnableQuicUpgrade           *bool   `json:"EnableQuicUpgrade" name:"EnableQuicUpgrade"`
		QuicListenerId              *string `json:"QuicListenerId" name:"QuicListenerId"`
	} `json:"AlbListenerSet"`
}

func (r *DescribeAlbListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbRuleGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupName      *string                                `json:"AlbRuleGroupName,omitempty" name:"AlbRuleGroupName"`
	AlbListenerId         *string                                `json:"AlbListenerId,omitempty" name:"AlbListenerId"`
	BackendServerGroupId  *string                                `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	Type                  *string                                `json:"Type,omitempty" name:"Type"`
	AlbRuleSet            []*CreateAlbRuleGroupAlbRuleSet        `json:"AlbRuleSet,omitempty" name:"AlbRuleSet"`
	RedirectAlbListenerId *string                                `json:"RedirectAlbListenerId,omitempty" name:"RedirectAlbListenerId"`
	RedirectHttpCode      *string                                `json:"RedirectHttpCode,omitempty" name:"RedirectHttpCode"`
	FixedResponseConfig   *CreateAlbRuleGroupFixedResponseConfig `json:"FixedResponseConfig,omitempty" name:"FixedResponseConfig"`
	RewriteConfig         *CreateAlbRuleGroupRewriteConfig       `json:"RewriteConfig,omitempty" name:"RewriteConfig"`
}

func (r *CreateAlbRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAlbRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbRuleGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	AlbRuleGroup struct {
		AlbRuleGroupId       *string `json:"AlbRuleGroupId" name:"AlbRuleGroupId"`
		AlbListenerId        *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbRuleGroupName     *string `json:"AlbRuleGroupName" name:"AlbRuleGroupName"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		AlbRuleSet           []struct {
			AlbRuleType   *string   `json:"AlbRuleType" name:"AlbRuleType"`
			AlbRuleValue  *string   `json:"AlbRuleValue" name:"AlbRuleValue"`
			MethodValue   []*string `json:"MethodValue" name:"MethodValue"`
			SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
			HeaderValue   []struct {
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
		} `json:"AlbRuleSet" name:"AlbRuleSet"`
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
	} `json:"AlbRuleGroup"`
}

func (r *CreateAlbRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbRuleGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupId *string `json:"AlbRuleGroupId,omitempty" name:"AlbRuleGroupId"`
}

func (r *DeleteAlbRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlbRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbRuleGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteAlbRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbRuleGroupsRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupId []*string                      `json:"AlbRuleGroupId,omitempty" name:"AlbRuleGroupId"`
	Filter         []*DescribeAlbRuleGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults     *int                           `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken      *string                        `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAlbRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbRuleGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlbRuleGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbRuleGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	NextToken       *string `json:"NextToken" name:"NextToken"`
	AlbRuleGroupSet []struct {
		AlbRuleGroupId       *string `json:"AlbRuleGroupId" name:"AlbRuleGroupId"`
		AlbListenerId        *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbRuleGroupName     *string `json:"AlbRuleGroupName" name:"AlbRuleGroupName"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		AlbRuleSet           []struct {
			AlbRuleType   *string   `json:"AlbRuleType" name:"AlbRuleType"`
			AlbRuleValue  *string   `json:"AlbRuleValue" name:"AlbRuleValue"`
			MethodValue   []*string `json:"MethodValue" name:"MethodValue"`
			SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
			HeaderValue   []struct {
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
		} `json:"AlbRuleSet" name:"AlbRuleSet"`
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
	} `json:"AlbRuleGroupSet"`
}

func (r *DescribeAlbRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbRuleGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupId        *string                                `json:"AlbRuleGroupId,omitempty" name:"AlbRuleGroupId"`
	AlbRuleGroupName      *string                                `json:"AlbRuleGroupName,omitempty" name:"AlbRuleGroupName"`
	BackendServerGroupId  *string                                `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	Type                  *string                                `json:"Type,omitempty" name:"Type"`
	AlbRuleSet            []*ModifyAlbRuleGroupAlbRuleSet        `json:"AlbRuleSet,omitempty" name:"AlbRuleSet"`
	RedirectAlbListenerId *string                                `json:"RedirectAlbListenerId,omitempty" name:"RedirectAlbListenerId"`
	RedirectHttpCode      *string                                `json:"RedirectHttpCode,omitempty" name:"RedirectHttpCode"`
	FixedResponseConfig   *ModifyAlbRuleGroupFixedResponseConfig `json:"FixedResponseConfig,omitempty" name:"FixedResponseConfig"`
}

func (r *ModifyAlbRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyAlbRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbRuleGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	AlbRuleGroup struct {
		AlbRuleGroupId       *string `json:"AlbRuleGroupId" name:"AlbRuleGroupId"`
		AlbListenerId        *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbRuleGroupName     *string `json:"AlbRuleGroupName" name:"AlbRuleGroupName"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		AlbRuleSet           []struct {
			AlbRuleType   *string   `json:"AlbRuleType" name:"AlbRuleType"`
			AlbRuleValue  *string   `json:"AlbRuleValue" name:"AlbRuleValue"`
			MethodValue   []*string `json:"MethodValue" name:"MethodValue"`
			SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
			HeaderValue   []struct {
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
		} `json:"AlbRuleSet" name:"AlbRuleSet"`
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
	} `json:"AlbRuleGroup"`
}

func (r *ModifyAlbRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAlbRuleRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupId *string `json:"AlbRuleGroupId,omitempty" name:"AlbRuleGroupId"`
	AlbRuleType    *string `json:"AlbRuleType,omitempty" name:"AlbRuleType"`
	AlbRuleValue   *string `json:"AlbRuleValue,omitempty" name:"AlbRuleValue"`
}

func (r *AddAlbRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAlbRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddAlbRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddAlbRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	AlbRule   struct {
		AlbRuleType  *string `json:"AlbRuleType" name:"AlbRuleType"`
		AlbRuleValue *string `json:"AlbRuleValue" name:"AlbRuleValue"`
	} `json:"AlbRule"`
}

func (r *AddAlbRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAlbRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbRuleRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupId *string `json:"AlbRuleGroupId,omitempty" name:"AlbRuleGroupId"`
	AlbRuleType    *string `json:"AlbRuleType,omitempty" name:"AlbRuleType"`
	AlbRuleValue   *string `json:"AlbRuleValue,omitempty" name:"AlbRuleValue"`
}

func (r *DeleteAlbRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlbRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteAlbRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbListenerCertGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerId *string `json:"AlbListenerId,omitempty" name:"AlbListenerId"`
}

func (r *CreateAlbListenerCertGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbListenerCertGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAlbListenerCertGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbListenerCertGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	AlbListenerCertGroup struct {
		AlbListenerId          *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbListenerCertGroupId *string `json:"AlbListenerCertGroupId" name:"AlbListenerCertGroupId"`
		AlbListenerCertSet     []struct {
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			CertificateId   *string `json:"CertificateId" name:"CertificateId"`
			CertificateName *string `json:"CertificateName" name:"CertificateName"`
			CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
			CommonName      *string `json:"CommonName" name:"CommonName"`
			ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		} `json:"AlbListenerCertSet" name:"AlbListenerCertSet"`
	} `json:"AlbListenerCertGroup"`
}

func (r *CreateAlbListenerCertGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbListenerCertGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbListenerCertGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerCertGroupId *string `json:"AlbListenerCertGroupId,omitempty" name:"AlbListenerCertGroupId"`
}

func (r *DeleteAlbListenerCertGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbListenerCertGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlbListenerCertGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbListenerCertGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteAlbListenerCertGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbListenerCertGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbListenerCertGroupsRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerCertGroupId []*string                              `json:"AlbListenerCertGroupId,omitempty" name:"AlbListenerCertGroupId"`
	Filter                 []*DescribeAlbListenerCertGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults             *int                                   `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken              *string                                `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAlbListenerCertGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbListenerCertGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlbListenerCertGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbListenerCertGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	NextToken               *string `json:"NextToken" name:"NextToken"`
	AlbListenerCertGroupSet []struct {
		AlbListenerId          *string `json:"AlbListenerId" name:"AlbListenerId"`
		AlbListenerCertGroupId *string `json:"AlbListenerCertGroupId" name:"AlbListenerCertGroupId"`
		AlbListenerCertSet     []struct {
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			CertificateId   *string `json:"CertificateId" name:"CertificateId"`
			CertificateName *string `json:"CertificateName" name:"CertificateName"`
			CertAuthority   *string `json:"CertAuthority" name:"CertAuthority"`
			CommonName      *string `json:"CommonName" name:"CommonName"`
			ExpireTime      *string `json:"ExpireTime" name:"ExpireTime"`
		} `json:"AlbListenerCertSet" name:"AlbListenerCertSet"`
	} `json:"AlbListenerCertGroupSet"`
}

func (r *DescribeAlbListenerCertGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbListenerCertGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateCertificateWithGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerCertGroupId *string `json:"AlbListenerCertGroupId,omitempty" name:"AlbListenerCertGroupId"`
	CertificateId          *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *AssociateCertificateWithGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateCertificateWithGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateCertificateWithGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type DissociateCertificateWithGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerCertGroupId *string `json:"AlbListenerCertGroupId,omitempty" name:"AlbListenerCertGroupId"`
	CertificateId          *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DissociateCertificateWithGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DissociateCertificateWithGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DissociateCertificateWithGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type SetEnableAlbAccessLogRequest struct {
	*ksyunhttp.BaseRequest
	AlbId      *string `json:"AlbId,omitempty" name:"AlbId"`
	EnabledLog *bool   `json:"EnabledLog,omitempty" name:"EnabledLog"`
}

func (r *SetEnableAlbAccessLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetEnableAlbAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetEnableAlbAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetEnableAlbAccessLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetEnableAlbAccessLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetEnableAlbAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbAccessLogRequest struct {
	*ksyunhttp.BaseRequest
	AlbId       *string `json:"AlbId,omitempty" name:"AlbId"`
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogpoolName *string `json:"LogpoolName,omitempty" name:"LogpoolName"`
}

func (r *SetAlbAccessLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetAlbAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbAccessLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	ApplicationLoadBalancer struct {
		AlbId          *string `json:"AlbId" name:"AlbId"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		AlbName        *string `json:"AlbName" name:"AlbName"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		AlbVersion     *string `json:"AlbVersion" name:"AlbVersion"`
		IpVersion      *string `json:"IpVersion" name:"IpVersion"`
		AlbType        *string `json:"AlbType" name:"AlbType"`
		PublicIp       *string `json:"PublicIp" name:"PublicIp"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		State          *string `json:"State" name:"State"`
		ListenersCount *int    `json:"ListenersCount" name:"ListenersCount"`
		Status         *string `json:"Status" name:"Status"`
		EnabledLog     *bool   `json:"EnabledLog" name:"EnabledLog"`
		KlogInfo       struct {
			ProjectName *string `json:"ProjectName" name:"ProjectName"`
			LogpoolName *string `json:"LogpoolName" name:"LogpoolName"`
			AccountId   *string `json:"AccountId" name:"AccountId"`
		} `json:"KlogInfo" name:"KlogInfo"`
		BillType         *int    `json:"BillType" name:"BillType"`
		ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		EnabledQuic      *bool   `json:"EnabledQuic" name:"EnabledQuic"`
		EnableHpa        *bool   `json:"EnableHpa" name:"EnableHpa"`
		DeleteProtection *string `json:"DeleteProtection" name:"DeleteProtection"`
		ModifyProtection *string `json:"ModifyProtection" name:"ModifyProtection"`
		TagSet           []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"ApplicationLoadBalancer"`
}

func (r *SetAlbAccessLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneLoadBalancerRequest struct {
	*ksyunhttp.BaseRequest
	CloneLoadBalancerId *string   `json:"cloneLoadBalancerId,omitempty" name:"cloneLoadBalancerId"`
	LoadBalancerName    *string   `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	Type                *string   `json:"Type,omitempty" name:"Type"`
	SubnetId            *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	PrivateIpAddress    *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	IpVersion           *string   `json:"IpVersion,omitempty" name:"IpVersion"`
	LbType              *string   `json:"LbType,omitempty" name:"LbType"`
	TagId               []*string `json:"TagId,omitempty" name:"TagId"`
	ProjectId           *string   `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CloneLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloneLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneLoadBalancerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	LoadBalancerId   *string `json:"LoadBalancerId" name:"LoadBalancerId"`
	LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
	Type             *string `json:"Type" name:"Type"`
	CreateTime       *string `json:"CreateTime" name:"CreateTime"`
	VpcId            *string `json:"VpcId" name:"VpcId"`
	PublicIp         *string `json:"PublicIp" name:"PublicIp"`
	IpVersion        *string `json:"IpVersion" name:"IpVersion"`
	LbType           *string `json:"LbType" name:"LbType"`
}

func (r *CloneLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLBDeleteProtectionRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	DeleteProtection *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
}

func (r *SetLBDeleteProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLBDeleteProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetLBDeleteProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type SetLBModificationProtectionRequest struct {
	*ksyunhttp.BaseRequest
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	ModificationProtection *string `json:"ModificationProtection,omitempty" name:"ModificationProtection"`
}

func (r *SetLBModificationProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLBModificationProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetLBModificationProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type ModifyCertificateWithGroupRequest struct {
	*ksyunhttp.BaseRequest
	AlbListenerCertGroupId *string `json:"AlbListenerCertGroupId,omitempty" name:"AlbListenerCertGroupId"`
	OldCertificateId       *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
	CertificateId          *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *ModifyCertificateWithGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertificateWithGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCertificateWithGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type CreateAlbBackendServerGroupRequest struct {
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
	HttpMethod               *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	HealthCode               *string `json:"HealthCode,omitempty" name:"HealthCode"`
}

func (r *CreateAlbBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbBackendServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAlbBackendServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlbBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BackendServerGroup struct {
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Name                   *string `json:"Name" name:"Name"`
		BackendServerType      *string `json:"BackendServerType" name:"BackendServerType"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		VpcId                  *string `json:"VpcId" name:"VpcId"`
		Protocol               *string `json:"Protocol" name:"Protocol"`
		BackendServerNumber    *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive      *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		IpVersion              *string `json:"IpVersion" name:"IpVersion"`
		Method                 *string `json:"Method" name:"Method"`
		HealthCheck            struct {
			HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval               *int    `json:"Interval" name:"Interval"`
			Timeout                *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			HostName               *string `json:"HostName" name:"HostName"`
			UrlPath                *string `json:"UrlPath" name:"UrlPath"`
			HealthProtocol         *string `json:"HealthProtocol" name:"HealthProtocol"`
			HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
			HealthCode             *string `json:"HealthCode" name:"HealthCode"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
	} `json:"BackendServerGroup"`
}

func (r *CreateAlbBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlbBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbBackendServerGroupRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
}

func (r *DeleteAlbBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbBackendServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAlbBackendServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlbBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteAlbBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlbBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbBackendServerGroupRequest struct {
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
	HttpMethod               *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
	HealthCode               *string `json:"HealthCode,omitempty" name:"HealthCode"`
}

func (r *ModifyAlbBackendServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbBackendServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyAlbBackendServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbBackendServerGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BackendServerGroup struct {
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Name                   *string `json:"Name" name:"Name"`
		BackendServerType      *string `json:"BackendServerType" name:"BackendServerType"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		VpcId                  *string `json:"VpcId" name:"VpcId"`
		Protocol               *string `json:"Protocol" name:"Protocol"`
		BackendServerNumber    *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive      *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		IpVersion              *string `json:"IpVersion" name:"IpVersion"`
		Method                 *string `json:"Method" name:"Method"`
		HealthCheck            struct {
			HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval               *int    `json:"Interval" name:"Interval"`
			Timeout                *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			HostName               *string `json:"HostName" name:"HostName"`
			UrlPath                *string `json:"UrlPath" name:"UrlPath"`
			HealthProtocol         *string `json:"HealthProtocol" name:"HealthProtocol"`
			HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
			HealthCode             *string `json:"HealthCode" name:"HealthCode"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
	} `json:"BackendServerGroup"`
}

func (r *ModifyAlbBackendServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbBackendServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbBackendServerGroupsRequest struct {
	*ksyunhttp.BaseRequest
	Filter     []*DescribeAlbBackendServerGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                                    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                                 `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAlbBackendServerGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbBackendServerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlbBackendServerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbBackendServerGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	BackendServerGroupSet []struct {
		CreateTime             *string `json:"CreateTime" name:"CreateTime"`
		BackendServerGroupId   *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		Name                   *string `json:"Name" name:"Name"`
		BackendServerType      *string `json:"BackendServerType" name:"BackendServerType"`
		BackendServerGroupType *string `json:"BackendServerGroupType" name:"BackendServerGroupType"`
		VpcId                  *string `json:"VpcId" name:"VpcId"`
		Protocol               *string `json:"Protocol" name:"Protocol"`
		BackendServerNumber    *int    `json:"BackendServerNumber" name:"BackendServerNumber"`
		UpstreamKeepalive      *string `json:"UpstreamKeepalive" name:"UpstreamKeepalive"`
		IpVersion              *string `json:"IpVersion" name:"IpVersion"`
		Method                 *string `json:"Method" name:"Method"`
		HealthCheck            struct {
			HealthCheckState       *string `json:"HealthCheckState" name:"HealthCheckState"`
			HealthyThreshold       *int    `json:"HealthyThreshold" name:"HealthyThreshold"`
			Interval               *int    `json:"Interval" name:"Interval"`
			Timeout                *int    `json:"Timeout" name:"Timeout"`
			UnhealthyThreshold     *int    `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
			HostName               *string `json:"HostName" name:"HostName"`
			UrlPath                *string `json:"UrlPath" name:"UrlPath"`
			HealthProtocol         *string `json:"HealthProtocol" name:"HealthProtocol"`
			HealthCheckConnectPort *int    `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
			HealthCode             *string `json:"HealthCode" name:"HealthCode"`
		} `json:"HealthCheck" name:"HealthCheck"`
		Session struct {
			SessionState             *string `json:"SessionState" name:"SessionState"`
			SessionPersistencePeriod *int    `json:"SessionPersistencePeriod" name:"SessionPersistencePeriod"`
			CookieType               *string `json:"CookieType" name:"CookieType"`
			CookieName               *string `json:"CookieName" name:"CookieName"`
		} `json:"Session" name:"Session"`
	} `json:"BackendServerGroupSet"`
}

func (r *DescribeAlbBackendServerGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbBackendServerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterAlbBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerGroupId   *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
	BackendServerIp        *string `json:"BackendServerIp,omitempty" name:"BackendServerIp"`
	Port                   *int    `json:"Port,omitempty" name:"Port"`
	Weight                 *int    `json:"Weight,omitempty" name:"Weight"`
	NetworkInterfaceId     *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
}

func (r *RegisterAlbBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterAlbBackendServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RegisterAlbBackendServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterAlbBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	BackendServer struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		BackendServerIp      *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		BackendServerId      *string `json:"BackendServerId" name:"BackendServerId"`
		Port                 *string `json:"Port" name:"Port"`
	} `json:"BackendServer"`
}

func (r *RegisterAlbBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterAlbBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterAlbBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerId *string `json:"BackendServerId,omitempty" name:"BackendServerId"`
}

func (r *DeregisterAlbBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterAlbBackendServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeregisterAlbBackendServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterAlbBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeregisterAlbBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterAlbBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbBackendServerRequest struct {
	*ksyunhttp.BaseRequest
	BackendServerId *string `json:"BackendServerId,omitempty" name:"BackendServerId"`
	Weight          *int    `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyAlbBackendServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbBackendServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyAlbBackendServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlbBackendServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	BackendServer struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		BackendServerIp      *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		BackendServerId      *string `json:"BackendServerId" name:"BackendServerId"`
		Port                 *string `json:"Port" name:"Port"`
	} `json:"BackendServer"`
}

func (r *ModifyAlbBackendServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlbBackendServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbBackendServersRequest struct {
	*ksyunhttp.BaseRequest
	Filter     []*DescribeAlbBackendServersFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                               `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                            `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAlbBackendServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbBackendServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAlbBackendServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlbBackendServersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	BackendServerSet []struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BackendServerGroupId *string `json:"BackendServerGroupId" name:"BackendServerGroupId"`
		BackendServerIp      *string `json:"BackendServerIp" name:"BackendServerIp"`
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		BackendServerId      *string `json:"BackendServerId" name:"BackendServerId"`
		Port                 *string `json:"Port" name:"Port"`
	} `json:"BackendServerSet"`
}

func (r *DescribeAlbBackendServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlbBackendServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBackendServerGroupWithListenerRequest struct {
	*ksyunhttp.BaseRequest
	ListenerId           *string `json:"ListenerId,omitempty" name:"ListenerId"`
	BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
}

func (r *RegisterBackendServerGroupWithListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBackendServerGroupWithListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RegisterBackendServerGroupWithListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBackendServerGroupWithListenerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RegisterBackendServerGroupWithListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBackendServerGroupWithListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPrivateLinkDeleteProtectionRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId       *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DeleteProtection *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
}

func (r *SetPrivateLinkDeleteProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPrivateLinkDeleteProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetPrivateLinkDeleteProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetPrivateLinkDeleteProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetPrivateLinkDeleteProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPrivateLinkDeleteProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbDeleteProtectionRequest struct {
	*ksyunhttp.BaseRequest
	AlbId            *string `json:"albId,omitempty" name:"albId"`
	DeleteProtection *string `json:"deleteProtection,omitempty" name:"deleteProtection"`
}

func (r *SetAlbDeleteProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbDeleteProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetAlbDeleteProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbDeleteProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetAlbDeleteProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbDeleteProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbModificationProtectionRequest struct {
	*ksyunhttp.BaseRequest
	AlbId                  *string `json:"albId,omitempty" name:"albId"`
	ModificationProtection *string `json:"modificationProtection,omitempty" name:"modificationProtection"`
}

func (r *SetAlbModificationProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbModificationProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetAlbModificationProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAlbModificationProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *SetAlbModificationProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAlbModificationProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAlbRulesRequest struct {
	*ksyunhttp.BaseRequest
	AlbRuleGroupId *string                   `json:"AlbRuleGroupId,omitempty" name:"AlbRuleGroupId"`
	AlbRuleType    *string                   `json:"AlbRuleType,omitempty" name:"AlbRuleType"`
	AlbRuleValue   *string                   `json:"AlbRuleValue,omitempty" name:"AlbRuleValue"`
	MethodValue    []*string                 `json:"MethodValue,omitempty" name:"MethodValue"`
	SourceIpValue  []*string                 `json:"SourceIpValue,omitempty" name:"SourceIpValue"`
	HeaderValue    []*AddAlbRulesHeaderValue `json:"HeaderValue,omitempty" name:"HeaderValue"`
	QueryValue     []*AddAlbRulesQueryValue  `json:"QueryValue,omitempty" name:"QueryValue"`
	CookieValue    []*AddAlbRulesCookieValue `json:"CookieValue,omitempty" name:"CookieValue"`
}

func (r *AddAlbRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAlbRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddAlbRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddAlbRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	AlbRule   struct {
		AlbRuleType   *string   `json:"AlbRuleType" name:"AlbRuleType"`
		AlbRuleValue  *string   `json:"AlbRuleValue" name:"AlbRuleValue"`
		MethodValue   []*string `json:"MethodValue" name:"MethodValue"`
		SourceIpValue []*string `json:"SourceIpValue" name:"SourceIpValue"`
		HeaderValue   []struct {
			Key   *string   `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"HeaderValue" name:"HeaderValue"`
		QueryValue []struct {
			Key   *string   `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"QueryValue" name:"QueryValue"`
		CookieValue []struct {
			Key   *string   `json:"Key" name:"Key"`
			Value []*string `json:"Value" name:"Value"`
		} `json:"CookieValue" name:"CookieValue"`
	} `json:"AlbRule"`
}

func (r *AddAlbRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAlbRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
