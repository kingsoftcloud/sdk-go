package v20211215

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateFunctionCode struct {
	Ks3BucketName *string `json:"Ks3BucketName,omitempty" name:"Ks3BucketName"`
	Ks3ObjectName *string `json:"Ks3ObjectName,omitempty" name:"Ks3ObjectName"`
}
type CreateFunctionEnvironmentVariables struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateFunctionEnvironment struct {
	Variables []*CreateFunctionEnvironmentVariables `json:"Variables,omitempty" name:"Variables"`
}
type CreateFunctionVpcConfig struct {
	EnableVpc    *bool   `json:"EnableVpc,omitempty" name:"EnableVpc"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	SgId         *string `json:"SgId,omitempty" name:"SgId"`
	CidrBlock    *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
}
type CreateFunctionLogConfig struct {
	EnableKlog *bool   `json:"EnableKlog,omitempty" name:"EnableKlog"`
	Project    *string `json:"Project,omitempty" name:"Project"`
	LogPool    *string `json:"LogPool,omitempty" name:"LogPool"`
}
type CreateFunctionLivenessProbeConfigHTTPGet struct {
	Protocol       *string `json:"Protocol,omitempty" name:"Protocol"`
	Path           *string `json:"Path,omitempty" name:"Path"`
	Port           *int    `json:"Port,omitempty" name:"Port"`
	TimeoutSeconds *int    `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`
}
type CreateFunctionLivenessProbeConfig struct {
	HTTPGet *CreateFunctionLivenessProbeConfigHTTPGet `json:"HTTPGet,omitempty" name:"HTTPGet"`
}
type CreateFunctionReadinessProbeConfigHTTPGet struct {
	Protocol       *string `json:"Protocol,omitempty" name:"Protocol"`
	Path           *string `json:"Path,omitempty" name:"Path"`
	TimeoutSeconds *int    `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`
	Port           *int    `json:"Port,omitempty" name:"Port"`
}
type CreateFunctionReadinessProbeConfig struct {
	HTTPGet *CreateFunctionReadinessProbeConfigHTTPGet `json:"HTTPGet,omitempty" name:"HTTPGet"`
}
type CreateFunctionCustomContainerConfig struct {
	Image        *string `json:"Image,omitempty" name:"Image"`
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
}
type CreateTriggerTriggerDesc struct {
	Methods      []*string `json:"Methods,omitempty" name:"Methods"`
	AuthRequired *bool     `json:"AuthRequired,omitempty" name:"AuthRequired"`
	Path         *string   `json:"Path,omitempty" name:"Path"`
	Timeout      *int      `json:"Timeout,omitempty" name:"Timeout"`
	Protocol     *string   `json:"Protocol,omitempty" name:"Protocol"`
}
type CreateAutoScaledTriggerTriggersTriggerConfig struct {
	QpsValue        *int    `json:"QpsValue,omitempty" name:"QpsValue"`
	CronType        *string `json:"CronType,omitempty" name:"CronType"`
	StartTime       *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime         *int64  `json:"EndTime,omitempty" name:"EndTime"`
	DesiredReplicas *int    `json:"DesiredReplicas,omitempty" name:"DesiredReplicas"`
	MemoryType      *string `json:"MemoryType,omitempty" name:"MemoryType"`
	MemoryValue     *int    `json:"MemoryValue,omitempty" name:"MemoryValue"`
}
type CreateAutoScaledTriggerTriggers struct {
	Name          *string                                       `json:"Name,omitempty" name:"Name"`
	Type          *string                                       `json:"Type,omitempty" name:"Type"`
	TriggerConfig *CreateAutoScaledTriggerTriggersTriggerConfig `json:"TriggerConfig,omitempty" name:"TriggerConfig"`
}
type ModifyAutoScaledTriggerTriggerConfig struct {
	QpsValue        *int    `json:"QpsValue,omitempty" name:"QpsValue"`
	CronType        *string `json:"CronType,omitempty" name:"CronType"`
	StartTime       *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime         *int64  `json:"EndTime,omitempty" name:"EndTime"`
	DesiredReplicas *int    `json:"DesiredReplicas,omitempty" name:"DesiredReplicas"`
	MemoryType      *string `json:"MemoryType,omitempty" name:"MemoryType"`
	MemoryValue     *int    `json:"MemoryValue,omitempty" name:"MemoryValue"`
}
type ModifyFunctionCode struct {
	Ks3BucketName *string `json:"Ks3BucketName,omitempty" name:"Ks3BucketName"`
	Ks3ObjectName *string `json:"Ks3ObjectName,omitempty" name:"Ks3ObjectName"`
}
type ModifyFunctionEnvironmentVariables struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type ModifyFunctionEnvironment struct {
	Variables []*ModifyFunctionEnvironmentVariables `json:"Variables,omitempty" name:"Variables"`
}
type ModifyFunctionVpcConfig struct {
	EnableVpc    *bool   `json:"EnableVpc,omitempty" name:"EnableVpc"`
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	SgId         *string `json:"SgId,omitempty" name:"SgId"`
	CidrBlock    *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
}
type ModifyFunctionLogConfig struct {
	EnableKlog *bool   `json:"EnableKlog,omitempty" name:"EnableKlog"`
	Project    *string `json:"Project,omitempty" name:"Project"`
	LogPool    *string `json:"LogPool,omitempty" name:"LogPool"`
}
type ModifyFunctionLivenessProbeConfigHTTPGet struct {
	Protocol       *string `json:"Protocol,omitempty" name:"Protocol"`
	Path           *string `json:"Path,omitempty" name:"Path"`
	Port           *int    `json:"Port,omitempty" name:"Port"`
	TimeoutSeconds *int    `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`
}
type ModifyFunctionLivenessProbeConfig struct {
	HTTPGet *ModifyFunctionLivenessProbeConfigHTTPGet `json:"HTTPGet,omitempty" name:"HTTPGet"`
}
type ModifyFunctionReadinessProbeConfigHTTPGet struct {
	Protocol       *string `json:"Protocol,omitempty" name:"Protocol"`
	Path           *string `json:"Path,omitempty" name:"Path"`
	Port           *int    `json:"Port,omitempty" name:"Port"`
	TimeoutSeconds *int    `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`
}
type ModifyFunctionReadinessProbeConfig struct {
	HTTPGet *ModifyFunctionReadinessProbeConfigHTTPGet `json:"HTTPGet,omitempty" name:"HTTPGet"`
}
type ModifyFunctionCustomContainerConfig struct {
	Image        *string `json:"Image,omitempty" name:"Image"`
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
}

type GetLogDateRequest struct {
	*ksyunhttp.BaseRequest
	Id *int `json:"id,omitempty" name:"id"`
}

func (r *GetLogDateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetLogDateResponse struct {
	*ksyunhttp.BaseResponse
	Id *int `json:"id" name:"id"`
}

func (r *GetLogDateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFunctionRequest struct {
	*ksyunhttp.BaseRequest
	Id                        *string                              `json:"Id,omitempty" name:"Id"`
	Name                      *string                              `json:"Name,omitempty" name:"Name"`
	Namespace                 *string                              `json:"Namespace,omitempty" name:"Namespace"`
	Runtime                   *string                              `json:"Runtime,omitempty" name:"Runtime"`
	CaPort                    *int                                 `json:"CaPort,omitempty" name:"CaPort"`
	StartupCommand            []*string                            `json:"StartupCommand,omitempty" name:"StartupCommand"`
	Description               *string                              `json:"Description,omitempty" name:"Description"`
	Timeout                   *int                                 `json:"Timeout,omitempty" name:"Timeout"`
	MemorySize                *int                                 `json:"MemorySize,omitempty" name:"MemorySize"`
	SingleInstanceConcurrency *int                                 `json:"SingleInstanceConcurrency,omitempty" name:"SingleInstanceConcurrency"`
	InternetAccess            *bool                                `json:"InternetAccess,omitempty" name:"InternetAccess"`
	Code                      *CreateFunctionCode                  `json:"Code,omitempty" name:"Code"`
	Environment               *CreateFunctionEnvironment           `json:"Environment,omitempty" name:"Environment"`
	VpcConfig                 *CreateFunctionVpcConfig             `json:"VpcConfig,omitempty" name:"VpcConfig"`
	LogConfig                 *CreateFunctionLogConfig             `json:"LogConfig,omitempty" name:"LogConfig"`
	LivenessProbeConfig       *CreateFunctionLivenessProbeConfig   `json:"LivenessProbeConfig,omitempty" name:"LivenessProbeConfig"`
	ReadinessProbeConfig      *CreateFunctionReadinessProbeConfig  `json:"ReadinessProbeConfig,omitempty" name:"ReadinessProbeConfig"`
	Layers                    []*string                            `json:"Layers,omitempty" name:"Layers"`
	CodeType                  *string                              `json:"CodeType,omitempty" name:"CodeType"`
	CustomContainerConfig     *CreateFunctionCustomContainerConfig `json:"CustomContainerConfig,omitempty" name:"CustomContainerConfig"`
}

func (r *CreateFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateFunctionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ID   *string `json:"ID" name:"ID"`
		Name struct {
		} `json:"Name" name:"Name"`
	} `json:"Data"`
}

func (r *CreateFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFunctionServiceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *CheckFunctionServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CheckFunctionServiceResponse struct {
	*ksyunhttp.BaseResponse
	Data      *string `json:"Data" name:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CheckFunctionServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFunctionServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenFunctionServiceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *OpenFunctionServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OpenFunctionServiceResponse struct {
	*ksyunhttp.BaseResponse
	Data *string `json:"Data" name:"Data"`
}

func (r *OpenFunctionServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenFunctionServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFunctionRequest struct {
	*ksyunhttp.BaseRequest
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteFunctionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTriggerRequest struct {
	*ksyunhttp.BaseRequest
	FunctionId  *string                   `json:"FunctionId,omitempty" name:"FunctionId"`
	TriggerName *string                   `json:"TriggerName,omitempty" name:"TriggerName"`
	Type        *string                   `json:"Type,omitempty" name:"Type"`
	Enable      *bool                     `json:"Enable,omitempty" name:"Enable"`
	TriggerDesc *CreateTriggerTriggerDesc `json:"TriggerDesc,omitempty" name:"TriggerDesc"`
}

func (r *CreateTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTriggerRequest struct {
	*ksyunhttp.BaseRequest
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrModifyAutoScalingConfigRequest struct {
	*ksyunhttp.BaseRequest
	FunctionId                 *string `json:"FunctionId,omitempty" name:"FunctionId"`
	IdleReplicaCount           *int    `json:"IdleReplicaCount,omitempty" name:"IdleReplicaCount"`
	MaxReplicaCount            *int    `json:"MaxReplicaCount,omitempty" name:"MaxReplicaCount"`
	MinReplicaCount            *int    `json:"MinReplicaCount,omitempty" name:"MinReplicaCount"`
	CooldownPeriod             *int    `json:"CooldownPeriod,omitempty" name:"CooldownPeriod"`
	StabilizationWindowSeconds *int    `json:"StabilizationWindowSeconds,omitempty" name:"StabilizationWindowSeconds"`
}

func (r *CreateOrModifyAutoScalingConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateOrModifyAutoScalingConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateOrModifyAutoScalingConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrModifyAutoScalingConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingConfigRequest struct {
	*ksyunhttp.BaseRequest
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`
}

func (r *DescribeAutoScalingConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAutoScalingConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		IdleReplicaCount           *int `json:"IdleReplicaCount" name:"IdleReplicaCount"`
		MaxReplicaCount            *int `json:"MaxReplicaCount" name:"MaxReplicaCount"`
		MinReplicaCount            *int `json:"MinReplicaCount" name:"MinReplicaCount"`
		CooldownPeriod             *int `json:"CooldownPeriod" name:"CooldownPeriod"`
		StabilizationWindowSeconds *int `json:"StabilizationWindowSeconds" name:"StabilizationWindowSeconds"`
	} `json:"Data"`
}

func (r *DescribeAutoScalingConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoScalingConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScaledTriggerRequest struct {
	*ksyunhttp.BaseRequest
	FunctionId *string                            `json:"FunctionId,omitempty" name:"FunctionId"`
	Triggers   []*CreateAutoScaledTriggerTriggers `json:"Triggers,omitempty" name:"Triggers"`
}

func (r *CreateAutoScaledTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAutoScaledTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateAutoScaledTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoScaledTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoScaledTriggerRequest struct {
	*ksyunhttp.BaseRequest
	FucntionId    *string                               `json:"FucntionId,omitempty" name:"FucntionId"`
	Name          *string                               `json:"Name,omitempty" name:"Name"`
	TriggerConfig *ModifyAutoScaledTriggerTriggerConfig `json:"TriggerConfig,omitempty" name:"TriggerConfig"`
}

func (r *ModifyAutoScaledTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyAutoScaledTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyAutoScaledTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoScaledTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScaledTriggersRequest struct {
	*ksyunhttp.BaseRequest
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`
	Name       *string `json:"Name,omitempty" name:"Name"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAutoScaledTriggersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAutoScaledTriggersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Triggers []struct {
			Name          *string `json:"Name" name:"Name"`
			Type          *string `json:"Type" name:"Type"`
			TriggerConfig struct {
				QpsValue        *int    `json:"QpsValue" name:"QpsValue"`
				StartTime       *int    `json:"StartTime" name:"StartTime"`
				EndTime         *int    `json:"EndTime" name:"EndTime"`
				CronType        *string `json:"CronType" name:"CronType"`
				MemoryType      *string `json:"MemoryType" name:"MemoryType"`
				MemoryValue     *string `json:"MemoryValue" name:"MemoryValue"`
				DesiredReplicas *int    `json:"DesiredReplicas" name:"DesiredReplicas"`
			} `json:"TriggerConfig"`
		} `json:"Triggers" name:"Triggers"`
	} `json:"Data"`
}

func (r *DescribeAutoScaledTriggersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoScaledTriggersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoScaledTriggerRequest struct {
	*ksyunhttp.BaseRequest
	Functionid  *string `json:"Functionid,omitempty" name:"Functionid"`
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
}

func (r *DeleteAutoScaledTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteAutoScaledTriggerResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteAutoScaledTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoScaledTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFunctionsRequest struct {
	*ksyunhttp.BaseRequest
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeFunctionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFunctionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Functions  []struct {
			Id                       *string `json:"Id" name:"Id"`
			Name                     *string `json:"Name" name:"Name"`
			Description              *string `json:"Description" name:"Description"`
			Runtime                  *string `json:"Runtime" name:"Runtime"`
			MemorySize               *int    `json:"MemorySize" name:"MemorySize"`
			CodeType                 *string `json:"CodeType" name:"CodeType"`
			State                    *string `json:"State" name:"State"`
			FunctionTotalInvocations *int    `json:"FunctionTotalInvocations" name:"FunctionTotalInvocations"`
			FunctionErrors           *int    `json:"FunctionErrors" name:"FunctionErrors"`
			RequestType              *string `json:"RequestType" name:"RequestType"`
			CreatedAt                *string `json:"CreatedAt" name:"CreatedAt"`
			UpdatedAt                *string `json:"UpdatedAt" name:"UpdatedAt"`
		} `json:"Functions" name:"Functions"`
	} `json:"Data"`
}

func (r *DescribeFunctionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFunctionRequest struct {
	*ksyunhttp.BaseRequest
	Id *string `json:"id,omitempty" name:"id"`
}

func (r *DescribeFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeFunctionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id          *string `json:"Id" name:"Id"`
		Namespace   *string `json:"Namespace" name:"Namespace"`
		Name        *string `json:"Name" name:"Name"`
		Description *string `json:"Description" name:"Description"`
		Runtime     *string `json:"Runtime" name:"Runtime"`
		Code        struct {
			TempKs3ObjectName *string `json:"TempKs3ObjectName" name:"TempKs3ObjectName"`
		} `json:"Code" name:"Code"`
		Timeout                   *int  `json:"Timeout" name:"Timeout"`
		SingleInstanceConcurrency *int  `json:"SingleInstanceConcurrency" name:"SingleInstanceConcurrency"`
		MemorySize                *int  `json:"MemorySize" name:"MemorySize"`
		InternetAccess            *bool `json:"InternetAccess" name:"InternetAccess"`
		VpcConfig                 struct {
			SubnetId  *string `json:"SubnetId" name:"SubnetId"`
			SgId      *string `json:"SgId" name:"SgId"`
			EnableVpc *bool   `json:"EnableVpc" name:"EnableVpc"`
			VpcId     *string `json:"VpcId" name:"VpcId"`
			Subnets   []struct {
				SubnetId         *string `json:"SubnetId" name:"SubnetId"`
				SgId             *string `json:"SgId" name:"SgId"`
				CidrBlock        *string `json:"CidrBlock" name:"CidrBlock"`
				AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			} `json:"Subnets"`
			VpcCidrBlock *string `json:"VpcCidrBlock" name:"VpcCidrBlock"`
		} `json:"VpcConfig" name:"VpcConfig"`
		LogConfig struct {
			EnableKlog *bool   `json:"EnableKlog" name:"EnableKlog"`
			Project    *string `json:"Project" name:"Project"`
			LogPool    *string `json:"LogPool" name:"LogPool"`
		} `json:"LogConfig" name:"LogConfig"`
		LivenessProbeConfig struct {
			HTTPGet struct {
				Protocol       *string `json:"Protocol" name:"Protocol"`
				Path           *string `json:"Path" name:"Path"`
				Port           *int    `json:"Port" name:"Port"`
				TimeoutSeconds *int    `json:"TimeoutSeconds" name:"TimeoutSeconds"`
			} `json:"HTTPGet"`
		} `json:"LivenessProbeConfig" name:"LivenessProbeConfig"`
		ReadinessProbeConfig struct {
			HTTPGet struct {
				Protocol       *string `json:"Protocol" name:"Protocol"`
				Path           *string `json:"Path" name:"Path"`
				Port           *int    `json:"Port" name:"Port"`
				TimeoutSeconds *int    `json:"TimeoutSeconds" name:"TimeoutSeconds"`
			} `json:"HTTPGet"`
		} `json:"ReadinessProbeConfig" name:"ReadinessProbeConfig"`
		AsyncConfig struct {
			RetryEnable *bool `json:"RetryEnable" name:"RetryEnable"`
		} `json:"AsyncConfig" name:"AsyncConfig"`
		CaPort                *int      `json:"CaPort" name:"CaPort"`
		StartupCommand        []*string `json:"StartupCommand" name:"StartupCommand"`
		State                 *string   `json:"State" name:"State"`
		CodeType              *string   `json:"CodeType" name:"CodeType"`
		CustomContainerConfig struct {
			Image     *string `json:"Image" name:"Image"`
			ImageName *string `json:"ImageName" name:"ImageName"`
			Version   *string `json:"Version" name:"Version"`
		} `json:"CustomContainerConfig" name:"CustomContainerConfig"`
	} `json:"Data"`
}

func (r *DescribeFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFunctionRequest struct {
	*ksyunhttp.BaseRequest
	Id                        *string                              `json:"Id,omitempty" name:"Id"`
	Runtime                   *string                              `json:"Runtime,omitempty" name:"Runtime"`
	CaPort                    *int                                 `json:"CaPort,omitempty" name:"CaPort"`
	StartupCommand            []*string                            `json:"StartupCommand,omitempty" name:"StartupCommand"`
	Timeout                   *int                                 `json:"Timeout,omitempty" name:"Timeout"`
	MemorySize                *int                                 `json:"MemorySize,omitempty" name:"MemorySize"`
	SingleInstanceConcurrency *int                                 `json:"SingleInstanceConcurrency,omitempty" name:"SingleInstanceConcurrency"`
	InternetAccess            *bool                                `json:"InternetAccess,omitempty" name:"InternetAccess"`
	Code                      *ModifyFunctionCode                  `json:"Code,omitempty" name:"Code"`
	Environment               *ModifyFunctionEnvironment           `json:"Environment,omitempty" name:"Environment"`
	VpcConfig                 *ModifyFunctionVpcConfig             `json:"VpcConfig,omitempty" name:"VpcConfig"`
	LogConfig                 *ModifyFunctionLogConfig             `json:"LogConfig,omitempty" name:"LogConfig"`
	LivenessProbeConfig       *ModifyFunctionLivenessProbeConfig   `json:"LivenessProbeConfig,omitempty" name:"LivenessProbeConfig"`
	ReadinessProbeConfig      *ModifyFunctionReadinessProbeConfig  `json:"ReadinessProbeConfig,omitempty" name:"ReadinessProbeConfig"`
	Layers                    []*string                            `json:"Layers,omitempty" name:"Layers"`
	CustomContainerConfig     *ModifyFunctionCustomContainerConfig `json:"CustomContainerConfig,omitempty" name:"CustomContainerConfig"`
}

func (r *ModifyFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyFunctionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTriggersRequest struct {
	*ksyunhttp.BaseRequest
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`
}

func (r *DescribeTriggersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTriggersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Triggers   []struct {
			Id          *string `json:"Id" name:"Id"`
			FunctionId  *string `json:"FunctionId" name:"FunctionId"`
			TriggerName *string `json:"TriggerName" name:"TriggerName"`
			Type        *string `json:"Type" name:"Type"`
			UrlInternet *string `json:"UrlInternet" name:"UrlInternet"`
			UrlIntranet *string `json:"UrlIntranet" name:"UrlIntranet"`
			TriggerDesc struct {
				AuthRequired *bool     `json:"AuthRequired" name:"AuthRequired"`
				Methods      []*string `json:"Methods" name:"Methods"`
			} `json:"TriggerDesc"`
		} `json:"Triggers" name:"Triggers"`
	} `json:"Data"`
}

func (r *DescribeTriggersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTriggersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
