package v20160304
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type DescribeListenersFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeInstancesWithListenerFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeHealthChecksFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeLoadBalancersFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeHostHeadersFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeRulesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeBackendServerGroupsFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeBackendServersFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}

type ListPrivateLinkServerFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}


type DescribeListenersRequest struct {
    *ksyunhttp.BaseRequest
    ListenerId []*string `json:"ListenerId,omitempty" name:"ListenerId"`
    Filter []*DescribeListenersFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	ListenerSet []struct {
		LoadBalancerId *string `json:"LoadBalancerId"`
		CreateTime *string `json:"CreateTime"`
		ListenerName *string `json:"ListenerName"`
		ListenerId *string `json:"ListenerId"`
		ListenerState *string `json:"ListenerState"`
		CertificateId *string `json:"CertificateId"`
		ListenerProtocol *string `json:"ListenerProtocol"`
		ListenerPort *int `json:"ListenerPort"`
		Method *string `json:"Method"`
		BandWidthIn *int `json:"BandWidthIn"`
		BandWidthOut *int `json:"BandWidthOut"`
		LoadBalancerAclId *string `json:"LoadBalancerAclId"`
		HttpProtocol *string `json:"HttpProtocol"`
		TlsCipherPolicy *string `json:"TlsCipherPolicy"`
		EnableHttp2 *bool `json:"EnableHttp2"`
		RedirectListenerId *string `json:"RedirectListenerId"`
		IpVersion *string `json:"IpVersion"`
		AsPrivateLinkServer *bool `json:"AsPrivateLinkServer"`
		AsPrivateLink *bool `json:"AsPrivateLink"`
		HealthStatus *bool `json:"HealthStatus"`
		BackendServerGroupIdSet []struct {
					BackendServerGroupId *string `json:"BackendServerGroupId"`
			} `json:"BackendServerGroupIdSet"`
			RealServer []struct {
						RegisterId *string `json:"RegisterId"`
						RealServerState *string `json:"RealServerState"`
						RealServerType *string `json:"RealServerType"`
						ListenerId *string `json:"ListenerId"`
						Weight *int `json:"Weight"`
						RealServerIp *string `json:"RealServerIp"`
						RealServerPort *int `json:"RealServerPort"`
						InstanceId *string `json:"InstanceId"`
						Tag *string `json:"Tag"`
						MasterSlaveType *string `json:"MasterSlaveType"`
						NetworkInterfaceId *string `json:"NetworkInterfaceId"`
				} `json:"RealServer"`
			} `json:"ListenerSet"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeListenersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyListenersRequest struct {
    *ksyunhttp.BaseRequest
    ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
    ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
    ListenerState *string `json:"ListenerState,omitempty" name:"ListenerState"`
    Method *string `json:"Method,omitempty" name:"Method"`
    BandWidthIn *int `json:"BandWidthIn,omitempty" name:"BandWidthIn"`
    BandWidthOut *int `json:"BandWidthOut,omitempty" name:"BandWidthOut"`
    HttpProtocol *string `json:"HttpProtocol,omitempty" name:"HttpProtocol"`
    TlsCipherPolicy *string `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
    EnableHttp2 *bool `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
    SessionState *string `json:"SessionState,omitempty" name:"SessionState"`
    SessionPersistencePeriod *int `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
    CookieType *string `json:"CookieType,omitempty" name:"CookieType"`
    CookieName *string `json:"CookieName,omitempty" name:"CookieName"`
    CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
    CreateTime *string `json:"CreateTime" name:"CreateTime"`
    ListenerName *string `json:"ListenerName" name:"ListenerName"`
    ListenerId *string `json:"ListenerId" name:"ListenerId"`
    ListenerState *string `json:"ListenerState" name:"ListenerState"`
    ListenerProtocol *string `json:"ListenerProtocol" name:"ListenerProtocol"`
    ListenerPort *int `json:"ListenerPort" name:"ListenerPort"`
    Method *string `json:"Method" name:"Method"`
    BandWidthIn *int `json:"BandWidthIn" name:"BandWidthIn"`
    BandWidthOut *int `json:"BandWidthOut" name:"BandWidthOut"`
    LoadBalancerAclId *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
    HttpProtocol *string `json:"HttpProtocol" name:"HttpProtocol"`
    TlsCipherPolicy *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
    EnableHttp2 *bool `json:"EnableHttp2" name:"EnableHttp2"`
    RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
	RealServer []struct {
		RegisterId *string `json:"RegisterId"`
		RealServerState *string `json:"RealServerState"`
		RealServerType *string `json:"RealServerType"`
		ListenerId *string `json:"ListenerId"`
		Weight *int `json:"Weight"`
		RealServerIp *string `json:"RealServerIp"`
		RealServerPort *int `json:"RealServerPort"`
		InstanceId *string `json:"InstanceId"`
		Tag *string `json:"Tag"`
		MasterSlaveType *string `json:"MasterSlaveType"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId"`
	} `json:"RealServer"`
    CertificateId *string `json:"CertificateId" name:"CertificateId"`
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
    LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
    ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
    ListenerState *string `json:"ListenerState,omitempty" name:"ListenerState"`
    ListenerProtocol *string `json:"ListenerProtocol,omitempty" name:"ListenerProtocol"`
    ListenerPort *int `json:"ListenerPort,omitempty" name:"ListenerPort"`
    Method *string `json:"Method,omitempty" name:"Method"`
    BandWidthIn *int `json:"BandWidthIn,omitempty" name:"BandWidthIn"`
    BandWidthOut *int `json:"BandWidthOut,omitempty" name:"BandWidthOut"`
    LoadBalancerAclId *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
    HttpProtocol *string `json:"HttpProtocol,omitempty" name:"HttpProtocol"`
    TlsCipherPolicy *string `json:"TlsCipherPolicy,omitempty" name:"TlsCipherPolicy"`
    EnableHttp2 *bool `json:"EnableHttp2,omitempty" name:"EnableHttp2"`
    RedirectListenerId *string `json:"RedirectListenerId,omitempty" name:"RedirectListenerId"`
    SessionState *string `json:"SessionState,omitempty" name:"SessionState"`
    SessionPersistencePeriod *int `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
    CookieType *string `json:"CookieType,omitempty" name:"CookieType"`
    CookieName *string `json:"CookieName,omitempty" name:"CookieName"`
    CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
    CreateTime *string `json:"CreateTime" name:"CreateTime"`
    ListenerName *string `json:"ListenerName" name:"ListenerName"`
    ListenerId *string `json:"ListenerId" name:"ListenerId"`
    ListenerState *string `json:"ListenerState" name:"ListenerState"`
    ListenerProtocol *string `json:"ListenerProtocol" name:"ListenerProtocol"`
    ListenerPort *int `json:"ListenerPort" name:"ListenerPort"`
    Method *string `json:"Method" name:"Method"`
    BandWidthIn *int `json:"BandWidthIn" name:"BandWidthIn"`
    BandWidthOut *int `json:"BandWidthOut" name:"BandWidthOut"`
    LoadBalancerAclId *string `json:"LoadBalancerAclId" name:"LoadBalancerAclId"`
    HttpProtocol *string `json:"HttpProtocol" name:"HttpProtocol"`
    TlsCipherPolicy *string `json:"TlsCipherPolicy" name:"TlsCipherPolicy"`
    EnableHttp2 *bool `json:"EnableHttp2" name:"EnableHttp2"`
    RedirectListenerId *string `json:"RedirectListenerId" name:"RedirectListenerId"`
    IpVersion *string `json:"IpVersion" name:"IpVersion"`
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
    RegisterId *string `json:"RegisterId,omitempty" name:"RegisterId"`
    Weight *int `json:"Weight,omitempty" name:"Weight"`
    RealServerPort *int `json:"RealServerPort,omitempty" name:"RealServerPort"`
    MasterSlaveType *string `json:"MasterSlaveType,omitempty" name:"MasterSlaveType"`
    Tag *string `json:"Tag,omitempty" name:"Tag"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    RegisterId *string `json:"RegisterId" name:"RegisterId"`
    RealServerState *string `json:"RealServerState" name:"RealServerState"`
    RealServerType *string `json:"RealServerType" name:"RealServerType"`
    ListenerId *string `json:"ListenerId" name:"ListenerId"`
    Weight *int `json:"Weight" name:"Weight"`
    RealServerIp *string `json:"RealServerIp" name:"RealServerIp"`
    RealServerPort *int `json:"RealServerPort" name:"RealServerPort"`
    Tag *string `json:"Tag" name:"Tag"`
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
    RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`
    ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
    Weight *int `json:"Weight,omitempty" name:"Weight"`
    RealServerIp *string `json:"RealServerIp,omitempty" name:"RealServerIp"`
    RealServerPort *int `json:"RealServerPort,omitempty" name:"RealServerPort"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Tag *string `json:"Tag,omitempty" name:"Tag"`
    MasterSlaveType *string `json:"MasterSlaveType,omitempty" name:"MasterSlaveType"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    RegisterId *string `json:"RegisterId" name:"RegisterId"`
    RealServerState *string `json:"RealServerState" name:"RealServerState"`
    RealServerType *string `json:"RealServerType" name:"RealServerType"`
    ListenerId *string `json:"ListenerId" name:"ListenerId"`
    Weight *int `json:"Weight" name:"Weight"`
    RealServerIp *string `json:"RealServerIp" name:"RealServerIp"`
    RealServerPort *int `json:"RealServerPort" name:"RealServerPort"`
    Tag *string `json:"Tag" name:"Tag"`
    MasterSlaveType *string `json:"MasterSlaveType" name:"MasterSlaveType"`
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
    Return *bool `json:"Return" name:"Return"`
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
    RegisterId []*string `json:"RegisterId,omitempty" name:"RegisterId"`
    Filter []*DescribeInstancesWithListenerFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	RealServerSet []struct {
		CreateTime *string `json:"CreateTime"`
		RegisterId *string `json:"RegisterId"`
		RealServerState *string `json:"RealServerState"`
		RealServerType *string `json:"RealServerType"`
		ListenerId *string `json:"ListenerId"`
		Weight *int `json:"Weight"`
		RealServerIp *string `json:"RealServerIp"`
		RealServerPort *int `json:"RealServerPort"`
		InstanceId *string `json:"InstanceId"`
		Tag *string `json:"Tag"`
		MasterSlaveType *string `json:"MasterSlaveType"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId"`
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
    HealthCheckId *string `json:"HealthCheckId,omitempty" name:"HealthCheckId"`
    HealthCheckState *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
    HealthCheckConnectPort *int `json:"HealthCheckConnectPort,omitempty" name:"HealthCheckConnectPort"`
    HealthyThreshold *int `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
    Interval *int `json:"Interval,omitempty" name:"Interval"`
    Timeout *int `json:"Timeout,omitempty" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
    HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
    UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    HealthCheckReq *string `json:"HealthCheckReq,omitempty" name:"HealthCheckReq"`
    HealthCheckExp *string `json:"HealthCheckExp,omitempty" name:"HealthCheckExp"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    HealthCheckState *string `json:"HealthCheckState" name:"HealthCheckState"`
    HealthCheckConnectPort *int `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
    HealthCheckId *string `json:"HealthCheckId" name:"HealthCheckId"`
    HealthyThreshold *int `json:"HealthyThreshold" name:"HealthyThreshold"`
    Interval *int `json:"Interval" name:"Interval"`
    Timeout *int `json:"Timeout" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
    UrlPath *string `json:"UrlPath" name:"UrlPath"`
    HostName *string `json:"HostName" name:"HostName"`
    HealthCheckReq *string `json:"HealthCheckReq" name:"HealthCheckReq"`
    HealthCheckExp *string `json:"HealthCheckExp" name:"HealthCheckExp"`
}

func (r *ModifyHealthCheckResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHealthCheckResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHealthChecksRequest struct {
    *ksyunhttp.BaseRequest
    HealthCheckId []*string `json:"HealthCheckId,omitempty" name:"HealthCheckId"`
    Filter []*DescribeHealthChecksFilter `json:"Filter,omitempty" name:"Filter"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    Marker *string `json:"Marker,omitempty" name:"Marker"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
	HealthCheckSet []struct {
		ListenerId *string `json:"ListenerId"`
		HealthCheckState *string `json:"HealthCheckState"`
		HealthCheckConnectPort *int `json:"HealthCheckConnectPort"`
		HealthCheckId *string `json:"HealthCheckId"`
		HealthyThreshold *int `json:"HealthyThreshold"`
		Interval *int `json:"Interval"`
		Timeout *int `json:"Timeout"`
		UnhealthyThreshold *int `json:"UnhealthyThreshold"`
		UrlPath *string `json:"UrlPath"`
		HostName *string `json:"HostName"`
		HealthCheckReq *string `json:"HealthCheckReq"`
		HealthCheckExp *string `json:"HealthCheckExp"`
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
    ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
    HealthCheckState *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
    HealthCheckConnectPort *int `json:"HealthCheckConnectPort,omitempty" name:"HealthCheckConnectPort"`
    HealthyThreshold *int `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
    Interval *int `json:"Interval,omitempty" name:"Interval"`
    Timeout *int `json:"Timeout,omitempty" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
    HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`
    UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    HealthCheckReq *string `json:"HealthCheckReq,omitempty" name:"HealthCheckReq"`
    HealthCheckExp *string `json:"HealthCheckExp,omitempty" name:"HealthCheckExp"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    HealthCheckState *string `json:"HealthCheckState" name:"HealthCheckState"`
    HealthCheckConnectPort *int `json:"HealthCheckConnectPort" name:"HealthCheckConnectPort"`
    HealthCheckId *string `json:"HealthCheckId" name:"HealthCheckId"`
    HealthyThreshold *int `json:"HealthyThreshold" name:"HealthyThreshold"`
    Interval *int `json:"Interval" name:"Interval"`
    Timeout *int `json:"Timeout" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold" name:"UnhealthyThreshold"`
    UrlPath *string `json:"UrlPath" name:"UrlPath"`
    HostName *string `json:"HostName" name:"HostName"`
    HealthCheckReq *string `json:"HealthCheckReq" name:"HealthCheckReq"`
    HealthCheckExp *string `json:"HealthCheckExp" name:"HealthCheckExp"`
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
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    LoadBalancerId []*string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
    Filter []*DescribeLoadBalancersFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
    State *string `json:"State,omitempty" name:"State"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	LoadBalancerDescriptions []struct {
		LoadBalancerId *string `json:"LoadBalancerId"`
		LoadBalancerName *string `json:"LoadBalancerName"`
		IsWaf *bool `json:"IsWaf"`
		Type *string `json:"Type"`
		CreateTime *string `json:"CreateTime"`
		ProjectId *string `json:"ProjectId"`
		VpcId *string `json:"VpcId"`
		ServiceEndTime *string `json:"ServiceEndTime"`
		PublicIp *string `json:"PublicIp"`
		State *string `json:"State"`
		IpVersion *string `json:"IpVersion"`
		LoadBalancerState *string `json:"LoadBalancerState"`
		ListenersCount *int `json:"ListenersCount"`
		ChargeType *string `json:"ChargeType"`
		LbType *string `json:"LbType"`
		LbStatus *string `json:"LbStatus"`
	} `json:"LoadBalancerDescriptions"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerRequest struct {
    *ksyunhttp.BaseRequest
    LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
    LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
    LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
    Type *string `json:"Type" name:"Type"`
    CreateTime *string `json:"CreateTime" name:"CreateTime"`
    VpcId *string `json:"VpcId" name:"VpcId"`
    PublicIp *string `json:"PublicIp" name:"PublicIp"`
    IpVersion *string `json:"IpVersion" name:"IpVersion"`
    LbType *string `json:"LbType" name:"LbType"`
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
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
    Type *string `json:"Type,omitempty" name:"Type"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
    IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
    LbType *string `json:"LbType,omitempty" name:"LbType"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    LoadBalancerId *string `json:"LoadBalancerId" name:"LoadBalancerId"`
    LoadBalancerName *string `json:"LoadBalancerName" name:"LoadBalancerName"`
    Type *string `json:"Type" name:"Type"`
    CreateTime *string `json:"CreateTime" name:"CreateTime"`
    VpcId *string `json:"VpcId" name:"VpcId"`
    PublicIp *string `json:"PublicIp" name:"PublicIp"`
    IpVersion *string `json:"IpVersion" name:"IpVersion"`
    LbType *string `json:"LbType" name:"LbType"`
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
    ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
    CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
    HostHeader *string `json:"HostHeader,omitempty" name:"HostHeader"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    HostHeaderId *string `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    HostHeaderId []*string `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
    Filter []*DescribeHostHeadersFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	HostHeaderSet []struct {
		CreateTime *string `json:"CreateTime"`
		HostHeaderId *string `json:"HostHeaderId"`
		ListenerId *string `json:"ListenerId"`
		CertificateId *string `json:"CertificateId"`
		HostHeader *string `json:"HostHeader"`
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
    Return *bool `json:"Return" name:"Return"`
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
    RuleId []*string `json:"RuleId,omitempty" name:"RuleId"`
    Filter []*DescribeRulesFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
	RuleSet []struct {
		Method *string `json:"Method"`
		BackendServerGroupId *string `json:"BackendServerGroupId"`
		Path *string `json:"Path"`
		RuleId *string `json:"RuleId"`
		ListenerSync *string `json:"ListenerSync"`
		HostHeaderId *string `json:"HostHeaderId"`
		CreateTime *string `json:"CreateTime"`
		BackendServerSet []struct {
					BackendServerIp *string `json:"BackendServerIp"`
					RegisterId *string `json:"RegisterId"`
					BackendServerPort *int `json:"BackendServerPort"`
					BackendServerState *string `json:"BackendServerState"`
			} `json:"BackendServerSet"`
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
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    BackendServerGroupName *string `json:"BackendServerGroupName,omitempty" name:"BackendServerGroupName"`
    BackendServerGroupType *string `json:"BackendServerGroupType,omitempty" name:"BackendServerGroupType"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
    HealthCheckState *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
    HealthyThreshold *int `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
    Interval *int `json:"Interval,omitempty" name:"Interval"`
    Timeout *int `json:"Timeout,omitempty" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
    UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
    Region *string `json:"Region,omitempty" name:"Region"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
    BackendServerGroupName *string `json:"BackendServerGroupName,omitempty" name:"BackendServerGroupName"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    BackendServerGroupId []*string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
    Filter []*DescribeBackendServerGroupsFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	BackendServerGroupSet []struct {
		CreateTime *string `json:"CreateTime"`
		BackendServerGroupId *string `json:"BackendServerGroupId"`
		VpcId *string `json:"VpcId"`
		BackendServerGroupName *string `json:"BackendServerGroupName"`
		BackendServerNumber *int `json:"BackendServerNumber"`
		BackendServerGroupType *string `json:"BackendServerGroupType"`
		IpVersion *string `json:"IpVersion"`
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
    BackendServerIp *string `json:"BackendServerIp,omitempty" name:"BackendServerIp"`
    BackendServerPort *int `json:"BackendServerPort,omitempty" name:"BackendServerPort"`
    Weight *int `json:"Weight,omitempty" name:"Weight"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    RegisterId []*string `json:"RegisterId,omitempty" name:"RegisterId"`
    Filter []*DescribeBackendServersFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	BackendServerSet []struct {
		CreateTime *string `json:"CreateTime"`
		BackendServerGroupId *string `json:"BackendServerGroupId"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId"`
		BackendServerIp *string `json:"BackendServerIp"`
		InstanceId *string `json:"InstanceId"`
		RegisterId *string `json:"RegisterId"`
		BackendServerPort *int `json:"BackendServerPort"`
		Weight *int `json:"Weight"`
		BackendServerState *string `json:"BackendServerState"`
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
    IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    LoadBalancerAclId *string `json:"LoadBalancerAclId,omitempty" name:"LoadBalancerAclId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
    RuleNumber *int `json:"RuleNumber,omitempty" name:"RuleNumber"`
    RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
    Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
    Description *string `json:"Description,omitempty" name:"Description"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    Return *bool `json:"Return" name:"Return"`
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
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
	LoadBalancerAclSet []struct {
		CreateTime *string `json:"CreateTime"`
		LoadBalancerAclName *string `json:"LoadBalancerAclName"`
		LoadBalancerAclId *string `json:"LoadBalancerAclId"`
		IpVersion *string `json:"IpVersion"`
		LoadBalancerAclEntrySet []struct {
					LoadBalancerAclId *string `json:"LoadBalancerAclId"`
					LoadBalancerAclEntryId *string `json:"LoadBalancerAclEntryId"`
					CidrBlock *string `json:"CidrBlock"`
					RuleNumber *int `json:"RuleNumber"`
					RuleAction *string `json:"RuleAction"`
					Protocol *string `json:"Protocol"`
					Description *string `json:"Description"`
			} `json:"LoadBalancerAclEntrySet"`
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
    Path *string `json:"Path,omitempty" name:"Path"`
    HostHeaderId *string `json:"HostHeaderId,omitempty" name:"HostHeaderId"`
    BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
    ListenerSync *string `json:"ListenerSync,omitempty" name:"ListenerSync"`
    Method *string `json:"Method,omitempty" name:"Method"`
    SessionState *string `json:"SessionState,omitempty" name:"SessionState"`
    SessionPersistencePeriod *int `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
    CookieType *string `json:"cookieType,omitempty" name:"cookieType"`
    CookieName *string `json:"CookieName,omitempty" name:"CookieName"`
    HealthCheckState *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
    HealthyThreshold *string `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
    Interval *int `json:"Interval,omitempty" name:"Interval"`
    Timeout *int `json:"Timeout,omitempty" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
    UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
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
    Path *string `json:"Path,omitempty" name:"Path"`
    RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
    BackendServerGroupId *string `json:"BackendServerGroupId,omitempty" name:"BackendServerGroupId"`
    ListenerSync *string `json:"ListenerSync,omitempty" name:"ListenerSync"`
    Method *string `json:"Method,omitempty" name:"Method"`
    SessionState *string `json:"SessionState,omitempty" name:"SessionState"`
    SessionPersistencePeriod *int `json:"SessionPersistencePeriod,omitempty" name:"SessionPersistencePeriod"`
    CookieType *string `json:"cookieType,omitempty" name:"cookieType"`
    CookieName *string `json:"CookieName,omitempty" name:"CookieName"`
    HealthCheckState *string `json:"HealthCheckState,omitempty" name:"HealthCheckState"`
    HealthyThreshold *string `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`
    Interval *int `json:"Interval,omitempty" name:"Interval"`
    Timeout *int `json:"Timeout,omitempty" name:"Timeout"`
    UnhealthyThreshold *int `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
    UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
    HostName *string `json:"HostName,omitempty" name:"HostName"`
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
    ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
    Description *string `json:"Description,omitempty" name:"Description"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *string `json:"TotalCount" name:"TotalCount"`
	PrivateLinkServerSet []struct {
		CreateTime *string `json:"CreateTime"`
		PrivateLinkServerName *string `json:"PrivateLinkServerName"`
		PrivateLinkServerId *string `json:"PrivateLinkServerId"`
		ListenerId *string `json:"ListenerId"`
		Description *string `json:"Description"`
		ProjectId *string `json:"ProjectId"`
		PrivateLinkNum *string `json:"PrivateLinkNum"`
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
    Return *bool `json:"Return" name:"Return"`
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
    PrivateLinkServerId *string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
    PrivateLinkServerName *string `json:"PrivateLinkServerName,omitempty" name:"PrivateLinkServerName"`
    Description *string `json:"Description,omitempty" name:"Description"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
    ListenerPort *int `json:"ListenerPort,omitempty" name:"ListenerPort"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
    TotalCount *string `json:"TotalCount" name:"TotalCount"`
	PrivateLinkSet []struct {
		CreateTime *string `json:"CreateTime"`
		PrivateLinkId *string `json:"PrivateLinkId"`
		PrivateLinkServerId *string `json:"PrivateLinkServerId"`
		AccountId *string `json:"AccountId"`
		ListenerId *string `json:"ListenerId"`
		ServiceAccountId *string `json:"ServiceAccountId"`
		UpdateTime *string `json:"UpdateTime"`
		ProjectId *string `json:"ProjectId"`
		ConnectionStatus *string `json:"ConnectionStatus"`
		LoadBalancerId *string `json:"LoadBalancerId"`
		ListenerPort *string `json:"ListenerPort"`
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
    Return *bool `json:"Return" name:"Return"`
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
    RuleNumber *int `json:"RuleNumber,omitempty" name:"RuleNumber"`
    RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
    Description *string `json:"Description,omitempty" name:"Description"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    Return *bool `json:"Return" name:"Return"`
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
    Return *bool `json:"Return" name:"Return"`
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
    PrivateLinkServerId *string `json:"PrivateLinkServerId,omitempty" name:"PrivateLinkServerId"`
    Filter []*ListPrivateLinkServerFilter `json:"Filter,omitempty" name:"Filter"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
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
    PrivateLinkId *string `json:"PrivateLinkId,omitempty" name:"PrivateLinkId"`
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
    Return *bool `json:"Return" name:"Return"`
}

func (r *RemovePrivateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemovePrivateLinkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetEnableAlbAccessLogRequest struct {
    *ksyunhttp.BaseRequest
    AlbId *string `json:"AlbId,omitempty" name:"AlbId"`
    EnabledLog *bool `json:"EnabledLog,omitempty" name:"EnabledLog"`
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
    Return *bool `json:"Return" name:"Return"`
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
    AlbId *string `json:"AlbId,omitempty" name:"AlbId"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SetAlbAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAlbAccessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

