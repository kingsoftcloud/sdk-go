package v20211215

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateFunctionCode struct {
	SourceUrl     *string `json:"SourceUrl,omitempty" name:"SourceUrl"`
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
type CreateTriggerTriggerDesc struct {
	Methods      []*string `json:"Methods,omitempty" name:"Methods"`
	AuthRequired *bool     `json:"AuthRequired,omitempty" name:"AuthRequired"`
	Path         *string   `json:"Path,omitempty" name:"Path"`
	Timeout      *int      `json:"Timeout,omitempty" name:"Timeout"`
	Protocol     *string   `json:"Protocol,omitempty" name:"Protocol"`
}
type ModifyFunctionCode struct {
	SourceUrl     *string `json:"SourceUrl,omitempty" name:"SourceUrl"`
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

type GetLogDateRequest struct {
	*ksyunhttp.BaseRequest
	Id *int `json:"id,omitempty" name:"id"`
}

func (r *GetLogDateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetLogDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
	Id                        *string                             `json:"Id,omitempty" name:"Id"`
	Name                      *string                             `json:"Name,omitempty" name:"Name"`
	Namespace                 *string                             `json:"Namespace,omitempty" name:"Namespace"`
	Runtime                   *string                             `json:"Runtime,omitempty" name:"Runtime"`
	CaPort                    *int                                `json:"CaPort,omitempty" name:"CaPort"`
	StartupCommand            []*string                           `json:"StartupCommand,omitempty" name:"StartupCommand"`
	Description               *string                             `json:"Description,omitempty" name:"Description"`
	Timeout                   *int                                `json:"Timeout,omitempty" name:"Timeout"`
	MemorySize                *int                                `json:"MemorySize,omitempty" name:"MemorySize"`
	SingleInstanceConcurrency *int                                `json:"SingleInstanceConcurrency,omitempty" name:"SingleInstanceConcurrency"`
	InternetAccess            *bool                               `json:"InternetAccess,omitempty" name:"InternetAccess"`
	Code                      *CreateFunctionCode                 `json:"Code,omitempty" name:"Code"`
	Environment               *CreateFunctionEnvironment          `json:"Environment,omitempty" name:"Environment"`
	VpcConfig                 *CreateFunctionVpcConfig            `json:"VpcConfig,omitempty" name:"VpcConfig"`
	LogConfig                 *CreateFunctionLogConfig            `json:"LogConfig,omitempty" name:"LogConfig"`
	LivenessProbeConfig       *CreateFunctionLivenessProbeConfig  `json:"LivenessProbeConfig,omitempty" name:"LivenessProbeConfig"`
	ReadinessProbeConfig      *CreateFunctionReadinessProbeConfig `json:"ReadinessProbeConfig,omitempty" name:"ReadinessProbeConfig"`
	Layers                    []*string                           `json:"Layers,omitempty" name:"Layers"`
}

func (r *CreateFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CheckFunctionServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CheckFunctionServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *OpenFunctionServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "OpenFunctionServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CreateTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type ModifyFunctionRequest struct {
	*ksyunhttp.BaseRequest
	Id                        *string                             `json:"Id,omitempty" name:"Id"`
	Runtime                   *string                             `json:"Runtime,omitempty" name:"Runtime"`
	CaPort                    *int                                `json:"CaPort,omitempty" name:"CaPort"`
	StartupCommand            []*string                           `json:"StartupCommand,omitempty" name:"StartupCommand"`
	Timeout                   *int                                `json:"Timeout,omitempty" name:"Timeout"`
	MemorySize                *int                                `json:"MemorySize,omitempty" name:"MemorySize"`
	SingleInstanceConcurrency *int                                `json:"SingleInstanceConcurrency,omitempty" name:"SingleInstanceConcurrency"`
	InternetAccess            *bool                               `json:"InternetAccess,omitempty" name:"InternetAccess"`
	Code                      *ModifyFunctionCode                 `json:"Code,omitempty" name:"Code"`
	Environment               *ModifyFunctionEnvironment          `json:"Environment,omitempty" name:"Environment"`
	VpcConfig                 *ModifyFunctionVpcConfig            `json:"VpcConfig,omitempty" name:"VpcConfig"`
	LogConfig                 *ModifyFunctionLogConfig            `json:"LogConfig,omitempty" name:"LogConfig"`
	LivenessProbeConfig       *ModifyFunctionLivenessProbeConfig  `json:"LivenessProbeConfig,omitempty" name:"LivenessProbeConfig"`
	ReadinessProbeConfig      *ModifyFunctionReadinessProbeConfig `json:"ReadinessProbeConfig,omitempty" name:"ReadinessProbeConfig"`
	Layers                    []*string                           `json:"Layers,omitempty" name:"Layers"`
}

func (r *ModifyFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeTriggersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeTriggersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
