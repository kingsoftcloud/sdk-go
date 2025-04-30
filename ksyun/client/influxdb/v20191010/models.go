package v20191010
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId     *string `json:"ProjectId,omitempty" name:"ProjectId"`
	InstanceName  *string `json:"InstanceName,omitempty" name:"InstanceName"`
	ProductType   *string `json:"ProductType,omitempty" name:"ProductType"`
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	InstanceType  *string `json:"InstanceType,omitempty" name:"InstanceType"`
	EbsType       *string `json:"EbsType,omitempty" name:"EbsType"`
	EbsSize       *int    `json:"EbsSize,omitempty" name:"EbsSize"`
	VpcId         *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId      *string `json:"SubnetId,omitempty" name:"SubnetId"`
	BillType      *int    `json:"BillType,omitempty" name:"BillType"`
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
	Error     *string `json:"Error" name:"Error"`
	Data      struct {
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceId  *string   `json:"InstanceId" name:"InstanceId"`
		OrderId     *string   `json:"OrderId" name:"OrderId"`
		SubOrderIds []*string `json:"SubOrderIds" name:"SubOrderIds"`
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
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
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
	Error     *string `json:"Error" name:"Error"`
	Data      struct {
		InstanceId       *string `json:"InstanceId" name:"InstanceId"`
		InstanceName     *string `json:"InstanceName" name:"InstanceName"`
		Region           *string `json:"Region" name:"Region"`
		RegionName       *string `json:"RegionName" name:"RegionName"`
		AvailabilityZoneName *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		AvailabilityZone *string `json:"AvailabilityZone" name:"AvailabilityZone"`
		Status           *string `json:"Status" name:"Status"`
		EngineVersion    *string `json:"EngineVersion" name:"EngineVersion"`
		BillType         *int    `json:"BillType" name:"BillType"`
		BillTypeName     *string `json:"BillTypeName" name:"BillTypeName"`
		ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
		Mode             *int    `json:"Mode" name:"Mode"`
		ModeName         *string `json:"ModeName" name:"ModeName"`
		InstanceType     *string `json:"InstanceType" name:"InstanceType"`
		EbsType          *string `json:"EbsType" name:"EbsType"`
		EbsSize          *int    `json:"EbsSize" name:"EbsSize"`
		Vip              *string `json:"Vip" name:"Vip"`
		Port             *int    `json:"Port" name:"Port"`
		VpcId            *string `json:"VpcId" name:"VpcId"`
		SubnetId         *string `json:"SubnetId" name:"SubnetId"`
		SecurityGroupId  *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		ExpirationTime   *string `json:"ExpirationTime" name:"ExpirationTime"`
		ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
		ProjectName      *string `json:"ProjectName" name:"ProjectName"`
		UsedDisk         *string `json:"UsedDisk" name:"UsedDisk"`
		MaxDisk          *string `json:"MaxDisk" name:"MaxDisk"`
		Eip              *string `json:"Eip" name:"Eip"`
		Eport            *int    `json:"Eport" name:"Eport"`
		Bandwidth        *string `json:"Bandwidth" name:"Bandwidth"`
	} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstancesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	Vip          *string `json:"Vip,omitempty" name:"Vip"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	FuzzySearch  *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
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
	Error     *string `json:"Error" name:"Error"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit *int `json:"Limit" name:"Limit"`
		Data  []struct {
			InstanceId           *string `json:"InstanceId" name:"InstanceId"`
			InstanceName         *string `json:"InstanceName" name:"InstanceName"`
			Region               *string `json:"Region" name:"Region"`
			RegionName           *string `json:"RegionName" name:"RegionName"`
			AvailabilityZoneName *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
			AvailabilityZone     *string `json:"AvailabilityZone" name:"AvailabilityZone"`
			Status               *string `json:"Status" name:"Status"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			BillType             *int    `json:"BillType" name:"BillType"`
			BillTypeName         *string `json:"BillTypeName" name:"BillTypeName"`
			ProductWhat          *int    `json:"ProductWhat" name:"ProductWhat"`
			Mode                 *int    `json:"Mode" name:"Mode"`
			ModeName             *string `json:"ModeName" name:"ModeName"`
			InstanceType         *string `json:"InstanceType" name:"InstanceType"`
			EbsType              *string `json:"EbsType" name:"EbsType"`
			EbsSize              *int    `json:"EbsSize" name:"EbsSize"`
			Vip                  *string `json:"Vip" name:"Vip"`
			Port                 *int    `json:"Port" name:"Port"`
			VpcId                *string `json:"VpcId" name:"VpcId"`
			SubnetId             *string `json:"SubnetId" name:"SubnetId"`
			SecurityGroupId      *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			CreateTime           *string `json:"CreateTime" name:"CreateTime"`
			ExpirationTime       *string `json:"ExpirationTime" name:"ExpirationTime"`
			ProjectId            *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName          *string `json:"ProjectName" name:"ProjectName"`
			UsedDisk             *string `json:"UsedDisk" name:"UsedDisk"`
			MaxDisk              *string `json:"MaxDisk" name:"MaxDisk"`
			Eip                  *string `json:"Eip" name:"Eip"`
			Eport                *int    `json:"Eport" name:"Eport"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceNodeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Error     *string `json:"Error" name:"Error"`
	Data      []struct {
		Id        *string `json:"Id" name:"Id"`
		Name      *string `json:"Name" name:"Name"`
		Status    *string `json:"Status" name:"Status"`
		Ip        *string `json:"Ip" name:"Ip"`
		Vip       *string `json:"Vip" name:"Vip"`
		Port      *int    `json:"Port" name:"Port"`
		UsedMemory *string `json:"UsedMemory" name:"UsedMemory"`
		MaxMemory *string `json:"MaxMemory" name:"MaxMemory"`
		Role      *string `json:"Role" name:"Role"`
	} `json:"Data"`
}

func (r *DescribeInstanceNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RenameInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenameInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeValidRegionsRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeValidRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeValidRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeValidRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeValidRegionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		RegionCode        *string `json:"RegionCode" name:"RegionCode"`
		RegionName        *string `json:"RegionName" name:"RegionName"`
		AvailabilityZones []struct {
			Code *string `json:"Code" name:"Code"`
			Name *string `json:"Name" name:"Name"`
		} `json:"AvailabilityZones" name:"AvailabilityZones"`
	} `json:"Data"`
}

func (r *DescribeValidRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeValidRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id            *string `json:"Id" name:"Id"`
		Name          *string `json:"Name" name:"Name"`
		SecurityRules []struct {
			Id       *string `json:"Id" name:"Id"`
			Cidr     *string `json:"Cidr" name:"Cidr"`
			Protocol *string `json:"Protocol" name:"Protocol"`
			FromPort *string `json:"FromPort" name:"FromPort"`
			ToPort   *string `json:"ToPort" name:"ToPort"`
		} `json:"SecurityRules" name:"SecurityRules"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"Data"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId        *string `json:"InstanceId,omitempty" name:"InstanceId"`
	SecurityRuleCidrs *string `json:"SecurityRuleCidrs,omitempty" name:"SecurityRuleCidrs"`
}

func (r *CreateSecurityRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id       *string `json:"Id" name:"Id"`
		Cidr     *string `json:"Cidr" name:"Cidr"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		ToPort   *string `json:"ToPort" name:"ToPort"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
	} `json:"Data"`
}

func (r *CreateSecurityRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	SecurityRuleIds *string `json:"SecurityRuleIds,omitempty" name:"SecurityRuleIds"`
}

func (r *DeleteSecurityRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSecurityRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDatabasesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Offset     *int    `json:"Offset,omitempty" name:"Offset"`
	Limit      *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit *int `json:"Limit" name:"Limit"`
		Data  []struct {
			DatabaseName *string `json:"DatabaseName" name:"DatabaseName"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *DeleteDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRetentionPolicyListRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *DescribeRetentionPolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRetentionPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRetentionPolicyListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Error     *string `json:"Error" name:"Error"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit *int `json:"Limit" name:"Limit"`
		Data  []struct {
			PolicyName         *string `json:"PolicyName" name:"PolicyName"`
			Duration           *string `json:"Duration" name:"Duration"`
			DefaultPolicy      *int    `json:"DefaultPolicy" name:"DefaultPolicy"`
			ShardGroupDuration *string `json:"ShardGroupDuration" name:"ShardGroupDuration"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *DescribeRetentionPolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateRetentionPolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName  *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	PolicyName    *string `json:"PolicyName,omitempty" name:"PolicyName"`
	Duration      *string `json:"Duration,omitempty" name:"Duration"`
	DefaultPolicy *string `json:"DefaultPolicy,omitempty" name:"DefaultPolicy"`
}

func (r *CreateRetentionPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRetentionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRetentionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRetentionPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateRetentionPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRetentionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRetentionPolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	PolicyName   *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *DeleteRetentionPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRetentionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRetentionPolicyResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteRetentionPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyRetentionPolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	PolicyName   *string `json:"PolicyName,omitempty" name:"PolicyName"`
	Duration     *string `json:"Duration,omitempty" name:"Duration"`
}

func (r *ModifyRetentionPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRetentionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyRetentionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRetentionPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyRetentionPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRetentionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMeasurementsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *DescribeMeasurementsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMeasurementsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMeasurementsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMeasurementsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit *int `json:"Limit" name:"Limit"`
		Data  []struct {
			Name *string `json:"Name" name:"Name"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *DescribeMeasurementsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMeasurementsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteMeasurementRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DatabaseName    *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	MeasurementName *string `json:"MeasurementName,omitempty" name:"MeasurementName"`
}

func (r *DeleteMeasurementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMeasurementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteMeasurementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMeasurementResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteMeasurementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMeasurementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeAccountsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit *int `json:"Limit" name:"Limit"`
		Data  []struct {
			Name *string `json:"Name" name:"Name"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateAccountRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName     *string `json:"AccountName,omitempty" name:"AccountName"`
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`
}

func (r *CreateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteAccountRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeAccountPrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Offset *int `json:"Offset" name:"Offset"`
		Limit *int `json:"Limit" name:"Limit"`
		Data  []struct {
			DatabaseName     *string `json:"DatabaseName" name:"DatabaseName"`
			AccountPrivilege *string `json:"AccountPrivilege" name:"AccountPrivilege"`
		} `json:"Data" name:"Data"`
	} `json:"Data"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GrantAccountPrivilegeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId       *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName      *string `json:"AccountName,omitempty" name:"AccountName"`
	DatabaseName     *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" name:"AccountPrivilege"`
}

func (r *GrantAccountPrivilegeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantAccountPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GrantAccountPrivilegeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantAccountPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeAccountPrivilegeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName  *string `json:"AccountName,omitempty" name:"AccountName"`
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *RevokeAccountPrivilegeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokeAccountPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RevokeAccountPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RevokeAccountPrivilegeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RevokeAccountPrivilegeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokeAccountPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName     *string `json:"AccountName,omitempty" name:"AccountName"`
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountDetailListRequest struct {
	*ksyunhttp.BaseRequest
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeAccountDetailListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAccountDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountDetailListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeAccountDetailListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

