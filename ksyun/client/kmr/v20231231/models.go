package v20231231
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type CreateAutoScalePolicyExecuteRules struct {
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}


type ListInstancesRequest struct {
	*ksyunhttp.BaseRequest
	PageNumber       *int      `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize         *int      `json:"PageSize,omitempty" name:"PageSize"`
	InstanceStatus   []*string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	InstanceNameOrId *string   `json:"InstanceNameOrId,omitempty" name:"InstanceNameOrId"`
}

func (r *ListInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		List []struct {
			InstanceId       *string `json:"InstanceId" name:"InstanceId"`
			InstanceName     *string `json:"InstanceName" name:"InstanceName"`
			InstanceStatus   *string `json:"InstanceStatus" name:"InstanceStatus"`
			AccountId        *string `json:"AccountId" name:"AccountId"`
			BeginTime        *string `json:"BeginTime" name:"BeginTime"`
			Duration         *int    `json:"Duration" name:"Duration"`
			DurationUnitDic  *int    `json:"DurationUnitDic" name:"DurationUnitDic"`
			RunningTime      *int    `json:"RunningTime" name:"RunningTime"`
			AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			Region           *string `json:"Region" name:"Region"`
			VpcId            *string `json:"VpcId" name:"VpcId"`
			VpcSubnetId      *string `json:"VpcSubnetId" name:"VpcSubnetId"`
			PackageType      *string `json:"PackageType" name:"PackageType"`
			ChargeType       *string `json:"ChargeType" name:"ChargeType"`
			RunMode          *string `json:"RunMode" name:"RunMode"`
			ResourceSpec     struct {
				Fe *int `json:"Fe" name:"Fe"`
				Cn *int `json:"Cn" name:"Cn"`
			} `json:"ResourceSpec"`
		} `json:"List" name:"List"`
	} `json:"Data"`
}

func (r *ListInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetInstanceDetailRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceDetailResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		RunningTime     *int    `json:"RunningTime" name:"RunningTime"`
		InstanceName    *string `json:"InstanceName" name:"InstanceName"`
		AccountId       *string `json:"AccountId" name:"AccountId"`
		TenantId        *string `json:"TenantId" name:"TenantId"`
		Status          *string `json:"Status" name:"Status"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		BeginTime       *string `json:"BeginTime" name:"BeginTime"`
		ChargeType      *string `json:"ChargeType" name:"ChargeType"`
		SlbId           *string `json:"SlbId" name:"SlbId"`
		SlbIp           *string `json:"SlbIp" name:"SlbIp"`
		TerminalSubnetId *string `json:"TerminalSubnetId" name:"TerminalSubnetId"`
		OrderId         *string `json:"OrderId" name:"OrderId"`
		PackageType     *string `json:"PackageType" name:"PackageType"`
		Region          *string `json:"Region" name:"Region"`
		VpcId           *string `json:"VpcId" name:"VpcId"`
		VpcCidr         *string `json:"VpcCidr" name:"VpcCidr"`
		VpcSubnetId     *string `json:"VpcSubnetId" name:"VpcSubnetId"`
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		InstanceInfo    struct {
			RunMode          *string `json:"RunMode" name:"RunMode"`
			Product          *string `json:"Product" name:"Product"`
			Version          *string `json:"Version" name:"Version"`
			HighAvailability *bool   `json:"HighAvailability" name:"HighAvailability"`
			FeEndpoints      struct {
				FeQueryPort   *int    `json:"FeQueryPort" name:"FeQueryPort"`
				FeHttpPort    *int    `json:"FeHttpPort" name:"FeHttpPort"`
				FePrivateSlbIP *string `json:"FePrivateSlbIP" name:"FePrivateSlbIP"`
				FePublicSlbIP *string `json:"FePublicSlbIP" name:"FePublicSlbIP"`
				FePublicAClId *string `json:"FePublicAClId" name:"FePublicAClId"`
			} `json:"FeEndpoints"`
			CnEndpoints struct {
				CnHttpPort *int `json:"CnHttpPort" name:"CnHttpPort"`
				CnPublicSlbIP *string `json:"CnPublicSlbIP" name:"CnPublicSlbIP"`
				CnPublicAClId *string `json:"CnPublicAClId" name:"CnPublicAClId"`
			} `json:"CnEndpoints"`
			FeResourceSpec struct {
				InitVolumeSize    *int    `json:"InitVolumeSize" name:"InitVolumeSize"`
				VolumeSize        *int    `json:"VolumeSize" name:"VolumeSize"`
				VolumeType        *string `json:"VolumeType" name:"VolumeType"`
				Cu                *int    `json:"Cu" name:"Cu"`
				NodeNumber        *int    `json:"NodeNumber" name:"NodeNumber"`
				ElasticNodeNumber *int    `json:"ElasticNodeNumber" name:"ElasticNodeNumber"`
			} `json:"FeResourceSpec"`
			CnResourceSpec struct {
				InitVolumeSize    *int    `json:"InitVolumeSize" name:"InitVolumeSize"`
				VolumeSize        *int    `json:"VolumeSize" name:"VolumeSize"`
				VolumeType        *string `json:"VolumeType" name:"VolumeType"`
				Cu                *int    `json:"Cu" name:"Cu"`
				NodeNumber        *int    `json:"NodeNumber" name:"NodeNumber"`
				ElasticNodeNumber *int    `json:"ElasticNodeNumber" name:"ElasticNodeNumber"`
			} `json:"CnResourceSpec"`
			KS3Spec struct {
				Bucket *string `json:"Bucket" name:"Bucket"`
				Prefix *string `json:"Prefix" name:"Prefix"`
			} `json:"KS3Spec"`
			ManagerAccessSpec struct {
				ManagerWebAddr *string `json:"ManagerWebAddr" name:"ManagerWebAddr"`
				UserName  *string   `json:"UserName" name:"UserName"`
				AllowList []*string `json:"AllowList" name:"AllowList"`
				DenyList  []*string `json:"DenyList" name:"DenyList"`
			} `json:"ManagerAccessSpec"`
			TunaHosts *string `json:"TunaHosts" name:"TunaHosts"`
		} `json:"InstanceInfo" name:"InstanceInfo"`
	} `json:"Data"`
}

func (r *GetInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	TunaHosts  []*string `json:"TunaHosts,omitempty" name:"TunaHosts"`
}

func (r *ModifyHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsResponse struct {
	*ksyunhttp.BaseResponse
	Status *int `json:"status" name:"status"`
	Data   struct {
		Code      *int    `json:"Code" name:"Code"`
		Message   *string `json:"Message" name:"Message"`
		RequestId *string `json:"RequestId" name:"RequestId"`
		Data      struct {
		} `json:"Data" name:"Data"`
	} `json:"Data"`
	Message *string `json:"message" name:"message"`
}

func (r *ModifyHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListAutoScaleHistoryRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ExecAtStart *string `json:"ExecAtStart,omitempty" name:"ExecAtStart"`
	ExecAtEnd   *string `json:"ExecAtEnd,omitempty" name:"ExecAtEnd"`
	PolicyName  *string `json:"PolicyName,omitempty" name:"PolicyName"`
	PageNumber  *int    `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize    *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListAutoScaleHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAutoScaleHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAutoScaleHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAutoScaleHistoryResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		List []struct {
			Id                     *int    `json:"Id" name:"Id"`
			ScaleHistoryId         *string `json:"ScaleHistoryId" name:"ScaleHistoryId"`
			InstanceId             *string `json:"InstanceId" name:"InstanceId"`
			PolicyId               *string `json:"PolicyId" name:"PolicyId"`
			PolicyName             *string `json:"PolicyName" name:"PolicyName"`
			ChargeType             *string `json:"ChargeType" name:"ChargeType"`
			ScaleVirtualInstanceId *string `json:"ScaleVirtualInstanceId" name:"ScaleVirtualInstanceId"`
			ScaleOrderId           *string `json:"ScaleOrderId" name:"ScaleOrderId"`
			NodeType               *string `json:"NodeType" name:"NodeType"`
			ScaleAction            *string `json:"ScaleAction" name:"ScaleAction"`
			ScaleNumber            *int    `json:"ScaleNumber" name:"ScaleNumber"`
			Status                 *string `json:"Status" name:"Status"`
			ScaleinHistoryId       *string `json:"ScaleinHistoryId" name:"ScaleinHistoryId"`
			ExecAt                 *string `json:"ExecAt" name:"ExecAt"`
			CompleteAt             *string `json:"CompleteAt" name:"CompleteAt"`
		} `json:"List" name:"List"`
	} `json:"Data"`
}

func (r *ListAutoScaleHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAutoScaleHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateAutoScalePolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string                            `json:"InstanceId,omitempty" name:"InstanceId"`
	PolicyName   *string                            `json:"PolicyName,omitempty" name:"PolicyName"`
	ChargeType   *string                            `json:"ChargeType,omitempty" name:"ChargeType"`
	ExecuteCycle *string                            `json:"ExecuteCycle,omitempty" name:"ExecuteCycle"`
	ExecuteRules *CreateAutoScalePolicyExecuteRules `json:"ExecuteRules,omitempty" name:"ExecuteRules"`
}

func (r *CreateAutoScalePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoScalePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAutoScalePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalePolicyResponse struct {
	*ksyunhttp.BaseResponse
	InstanceId   *string `json:"InstanceId" name:"InstanceId"`
	PolicyName   *string `json:"PolicyName" name:"PolicyName"`
	ChargeType   *string `json:"ChargeType" name:"ChargeType"`
	ExecuteCycle *string `json:"ExecuteCycle" name:"ExecuteCycle"`
	ExecuteRules struct {
		StartTime *string `json:"StartTime" name:"StartTime"`
		EndTime  *string `json:"EndTime" name:"EndTime"`
		ScaleNum *int    `json:"ScaleNum" name:"ScaleNum"`
	} `json:"ExecuteRules"`
	NodeType *string `json:"NodeType" name:"NodeType"`
}

func (r *CreateAutoScalePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoScalePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListAutoScalePolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ListAutoScalePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAutoScalePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAutoScalePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAutoScalePolicyResponse struct {
	*ksyunhttp.BaseResponse
	Status *int `json:"status" name:"status"`
	Data   struct {
		Code    *int    `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
		RequestId *string `json:"RequestId" name:"RequestId"`
		Data    []struct {
			Id           *int    `json:"Id" name:"Id"`
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			PolicyId     *string `json:"PolicyId" name:"PolicyId"`
			PolicyName   *string `json:"PolicyName" name:"PolicyName"`
			ChargeType   *string `json:"ChargeType" name:"ChargeType"`
			ExecuteCycle *string `json:"ExecuteCycle" name:"ExecuteCycle"`
			NodeType     *string `json:"NodeType" name:"NodeType"`
			Status       *int    `json:"Status" name:"Status"`
			ExecuteRules struct {
				StartTime *string `json:"StartTime" name:"StartTime"`
				EndTime  *string `json:"EndTime" name:"EndTime"`
				ScaleNum *int    `json:"ScaleNum" name:"ScaleNum"`
			} `json:"ExecuteRules"`
			CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
			UpdatedAt *string `json:"UpdatedAt" name:"UpdatedAt"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
	Message *string `json:"message" name:"message"`
}

func (r *ListAutoScalePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAutoScalePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteAutoScalePolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	PolicyId   *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteAutoScalePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoScalePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAutoScalePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoScalePolicyResponse struct {
	*ksyunhttp.BaseResponse
	Status *int `json:"status" name:"status"`
	Data   struct {
		Code    *int    `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
		RequestId *string `json:"RequestId" name:"RequestId"`
		Data    struct {
			InstanceId *string `json:"InstanceId" name:"InstanceId"`
			PolicyId   *string `json:"PolicyId" name:"PolicyId"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
	Message *string `json:"message" name:"message"`
}

func (r *DeleteAutoScalePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoScalePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

