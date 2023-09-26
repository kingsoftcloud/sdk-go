package v20191017
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateInstanceRequest struct {
    *ksyunhttp.BaseRequest
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
    InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
    BillType *int `json:"BillType,omitempty" name:"BillType"`
    Duration *int `json:"Duration,omitempty" name:"Duration"`
    Mode *string `json:"Mode,omitempty" name:"Mode"`
    InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
    SsdDisk *int `json:"SsdDisk,omitempty" name:"SsdDisk"`
    NodeNum *int `json:"NodeNum,omitempty" name:"NodeNum"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		InstanceType *string `json:"InstanceType"`
		SsdDisk *int `json:"SsdDisk"`
		SubOrderId *string `json:"SubOrderId"`
	} `json:"Data"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *string `json:"Data" name:"Data"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    Offset *int `json:"Offset,omitempty" name:"Offset"`
    Limit *int `json:"Limit,omitempty" name:"Limit"`
    OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		Offset *int `json:"Offset"`
		Limit *int `json:"Limit"`
		Total *int `json:"Total"`
		Instances []struct {
					UserId *string `json:"UserId"`
					Region *string `json:"Region"`
					InstanceName *string `json:"InstanceName"`
					InstanceId *string `json:"InstanceId"`
					StatusName *string `json:"StatusName"`
					ExpirationDate *string `json:"ExpirationDate"`
					Status *string `json:"Status"`
					Vip *string `json:"Vip"`
					WebVip *string `json:"WebVip"`
					InstanceType *string `json:"InstanceType"`
					SsdDisk *int `json:"SsdDisk"`
					Protocol *string `json:"Protocol"`
					SecurityGroupId *int `json:"SecurityGroupId"`
					Port *string `json:"Port"`
					NetworkType *string `json:"NetworkType"`
					VpcId *string `json:"VpcId"`
					SubnetId *string `json:"SubnetId"`
					ProductId *string `json:"ProductId"`
					BillType *string `json:"BillType"`
					CreateDate *int `json:"CreateDate"`
					ProjectId *string `json:"ProjectId"`
					ProjectName *string `json:"ProjectName"`
					NodeNum *string `json:"NodeNum"`
					AvailabilityZone *string `json:"AvailabilityZone"`
					ProductWhat *string `json:"ProductWhat"`
					Mode *string `json:"Mode"`
					ModeName *string `json:"ModeName"`
					Eip *string `json:"Eip"`
					WebEip *string `json:"WebEip"`
					EipEgress *string `json:"EipEgress"`
			} `json:"Instances"`
		} `json:"Data"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		UserId *string `json:"UserId"`
		Region *string `json:"Region"`
		InstanceName *string `json:"InstanceName"`
		InstanceId *string `json:"InstanceId"`
		StatusName *string `json:"StatusName"`
		Status *string `json:"Status"`
		Vip *string `json:"Vip"`
		WebVip *string `json:"WebVip"`
		InstanceType *string `json:"InstanceType"`
		SsdDisk *int `json:"SsdDisk"`
		Protocol *string `json:"Protocol"`
		SecurityGroupId *int `json:"SecurityGroupId"`
		Port *string `json:"Port"`
		NetworkType *string `json:"NetworkType"`
		VpcId *string `json:"VpcId"`
		SubnetId *string `json:"SubnetId"`
		ProductId *string `json:"ProductId"`
		BillType *string `json:"BillType"`
		CreateDate *int `json:"CreateDate"`
		ProjectId *string `json:"ProjectId"`
		ProjectName *string `json:"ProjectName"`
		NodeNum *string `json:"NodeNum"`
		AvailabilityZone *string `json:"AvailabilityZone"`
		ProductWhat *string `json:"ProductWhat"`
		Mode *string `json:"Mode"`
		ModeName *string `json:"ModeName"`
		Eip *string `json:"Eip"`
		WebEip *string `json:"WebEip"`
		EipEgress *string `json:"EipEgress"`
	} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodesRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		NodeId *string `json:"NodeId"`
		Name *string `json:"Name"`
		Role *string `json:"Role"`
		Ip *string `json:"Ip"`
		Port *string `json:"Port"`
		StatusName *string `json:"StatusName"`
		Status *string `json:"Status"`
	} `json:"Data"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeValidRegionRequest struct {
    *ksyunhttp.BaseRequest
    Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeValidRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeValidRegionRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeValidRegionRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeValidRegionResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data struct {
		Regions []struct {
					Name *string `json:"Name"`
					Code *string `json:"Code"`
					RegionEnName *string `json:"RegionEnName"`
					AreaCode *string `json:"AreaCode"`
					AreaName *string `json:"AreaName"`
					AreaEnName *string `json:"AreaEnName"`
				AvailabilityZones []struct {
					Code *string `json:"Code"`
					Name *string `json:"Name"`
				} `json:"AvailabilityZones"`
			} `json:"Regions"`
		} `json:"Data"`
}

func (r *DescribeValidRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeValidRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
    *ksyunhttp.BaseRequest
    Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
    *ksyunhttp.BaseResponse
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRulesRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupRulesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRulesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRulesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		Id *string `json:"Id"`
		Protocol *string `json:"Protocol"`
		FromPort *string `json:"FromPort"`
		ToPort *string `json:"ToPort"`
		Cidr *string `json:"Cidr"`
	} `json:"Data"`
}

func (r *DescribeSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSecurityGroupRuleRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Cidrs *string `json:"Cidrs,omitempty" name:"Cidrs"`
}

func (r *AddSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSecurityGroupRuleRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddSecurityGroupRuleRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddSecurityGroupRuleResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		Id *string `json:"Id"`
		Protocol *string `json:"Protocol"`
		FromPort *string `json:"FromPort"`
		ToPort *string `json:"ToPort"`
		Cidr *string `json:"Cidr"`
	} `json:"Data"`
}

func (r *AddSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSecurityGroupRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRulesRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    Cidrs *string `json:"Cidrs,omitempty" name:"Cidrs"`
}

func (r *DeleteSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupRulesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRulesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRulesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
	Data []struct {
		Id *string `json:"Id"`
		Protocol *string `json:"Protocol"`
		FromPort *string `json:"FromPort"`
		ToPort *string `json:"ToPort"`
		Cidr *string `json:"Cidr"`
	} `json:"Data"`
}

func (r *DeleteSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *string `json:"Data" name:"Data"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenameRequest struct {
    *ksyunhttp.BaseRequest
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *RenameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type RenameResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Code *string `json:"Code" name:"Code"`
    Message *string `json:"Message" name:"Message"`
    Data *string `json:"Data" name:"Data"`
}

func (r *RenameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

