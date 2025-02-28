package v20181225

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateSecurityGroupSecurityGroupRule struct {
	SecurityGroupRuleName     *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
	SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol,omitempty" name:"SecurityGroupRuleProtocol"`
}

type ModifySecurityGroupRuleSecurityGroupRule struct {
	SecurityGroupRuleId       *string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
	SecurityGroupRuleName     *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
	SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol,omitempty" name:"SecurityGroupRuleProtocol"`
}

type CreateDBParameterGroupParameters struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyDBParameterGroupParameters struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	Mem                 *int      `json:"Mem,omitempty" name:"Mem"`
	Disk                *int      `json:"Disk,omitempty" name:"Disk"`
	DBInstanceName      *string   `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	Engine              *string   `json:"Engine,omitempty" name:"Engine"`
	EngineVersion       *string   `json:"EngineVersion,omitempty" name:"EngineVersion"`
	MasterUserPassword  *string   `json:"MasterUserPassword,omitempty" name:"MasterUserPassword"`
	MasterUserName      *string   `json:"MasterUserName,omitempty" name:"MasterUserName"`
	DBInstanceType      *string   `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	VpcId               *string   `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId            *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	PreferredBackupTime *string   `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	DBParameterGroupId  *string   `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	SecurityGroupId     *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Port                *string   `json:"Port,omitempty" name:"Port"`
	BillType            *string   `json:"BillType,omitempty" name:"BillType"`
	Duration            *string   `json:"Duration,omitempty" name:"Duration"`
	DurationUnit        *string   `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone    []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ProjectId           *string   `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			Vip                   *string `json:"Vip"`
			Port                  *int    `json:"Port"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime"`
			MasterUserName        *string `json:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZoneList  []struct {
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int  `json:"DiskUsed"`
			Audit                            *bool `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier   *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBInstanceType         *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	Keyword                *string `json:"Keyword,omitempty" name:"Keyword"`
	Marker                 *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords             *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
	GroupId                *string `json:"GroupId,omitempty" name:"GroupId"`
	ProjectId              *string `json:"ProjectId,omitempty" name:"ProjectId"`
	DBInstanceIdentifierIn *string `json:"DBInstanceIdentifierIn,omitempty" name:"DBInstanceIdentifierIn"`
	DBInstanceNameIn       *string `json:"DBInstanceNameIn,omitempty" name:"DBInstanceNameIn"`
	VipIn                  *string `json:"VipIn,omitempty" name:"VipIn"`
	ExpiryDateLessThan     *int    `json:"ExpiryDateLessThan,omitempty" name:"ExpiryDateLessThan"`
	Order                  *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			Eip                              *string `json:"Eip"`
			EipPort                          *int    `json:"EipPort"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			MiniVersion      *string `json:"MiniVersion"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			NetworkType   *int  `json:"NetworkType"`
			SupportIPV6   *bool `json:"SupportIPV6"`
			BindInstances []struct {
			} `json:"BindInstances"`
			ProxyNodeInfo []struct {
			} `json:"ProxyNodeInfo"`
			ProxyInfo []struct {
			} `json:"ProxyInfo"`
			AutoSwitch *int `json:"AutoSwitch"`
			BillTypeId *int `json:"BillTypeId"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DeleteDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticDBInstancesRequest struct {
	*ksyunhttp.BaseRequest
	ExpiryDateLessThan *string `json:"ExpiryDateLessThan,omitempty" name:"ExpiryDateLessThan"`
	GroupId            *string `json:"GroupId,omitempty" name:"GroupId"`
	Keyword            *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *StatisticDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "StatisticDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StatisticDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		TotalCount *int `json:"TotalCount"`
		Statistic  struct {
			RUNNING_TASK *int `json:"RUNNING_TASK"`
			ACTIVE       *int `json:"ACTIVE"`
			INVALID      *int `json:"INVALID"`
		} `json:"Statistic"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StatisticDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	MasterUserPassword   *string `json:"MasterUserPassword,omitempty" name:"MasterUserPassword"`
	DBParameterGroupId   *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *ModifyDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupName        *string                                 `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupDescription *string                                 `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
	DBInstanceIdentifier     []*string                               `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	SecurityGroupRule        []*CreateSecurityGroupSecurityGroupRule `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				Vip                  *string `json:"Vip"`
				Created              *string `json:"Created"`
				DBInstanceType       *string `json:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId          *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName        *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
}

func (r *ModifySecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId          *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName        *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloneSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId         *string                                     `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupRuleAction *string                                     `json:"SecurityGroupRuleAction,omitempty" name:"SecurityGroupRuleAction"`
	SecurityGroupRule       []*ModifySecurityGroupRuleSecurityGroupRule `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
}

func (r *ModifySecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupRelationRequest struct {
	*ksyunhttp.BaseRequest
	RelationAction       *string   `json:"RelationAction,omitempty" name:"RelationAction"`
	SecurityGroupId      *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	DBInstanceIdentifier []*string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SecurityGroupRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SecurityGroupRelationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SecurityGroupRelationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupRelationResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string `json:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription"`
			Created                  *string `json:"Created"`
			Instances                []struct {
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SecurityGroupRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SecurityGroupRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleNameRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupRuleId   *string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
	SecurityGroupRuleName *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}

func (r *ModifySecurityGroupRuleNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRuleNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySecurityGroupRuleNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRuleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBLogType            *string `json:"DBLogType,omitempty" name:"DBLogType"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
	MaxFileSize          *int    `json:"MaxFileSize,omitempty" name:"MaxFileSize"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DescribeDBLogFiles []struct {
			LogFileName *string `json:"LogFileName"`
			Size        *int    `json:"Size"`
			RawSize     *int    `json:"RawSize"`
			Date        *string `json:"Date"`
			StartTime   *string `json:"StartTime"`
			EndTime     *string `json:"EndTime"`
		} `json:"DescribeDBLogFiles"`
		TotalCount *int `json:"TotalCount"`
		Marker     *int `json:"Marker"`
		MaxRecords *int `json:"MaxRecords"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupName         *string `json:"DBBackupName,omitempty" name:"DBBackupName"`
	Description          *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	DBBackup struct {
		DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
		DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
		Engine               *string `json:"Engine"`
		EngineVersion        *string `json:"EngineVersion"`
		BackupCreateTime     *string `json:"BackupCreateTime"`
		BackupUpdatedTime    *string `json:"BackupUpdatedTime"`
		DBBackupName         *string `json:"DBBackupName"`
		Description          *string `json:"Description"`
		BackupType           *string `json:"BackupType"`
		Status               *string `json:"Status"`
	} `json:"DBBackup"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBBackupIdentifier *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
}

func (r *DeleteDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDBBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	BackupType           *string `json:"BackupType,omitempty" name:"BackupType"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBBackup []struct {
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
			DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
			Engine               *string `json:"Engine"`
			EngineVersion        *string `json:"EngineVersion"`
			BackupCreateTime     *string `json:"BackupCreateTime"`
			BackupUpdatedTime    *string `json:"BackupUpdatedTime"`
			DBBackupName         *string `json:"DBBackupName"`
			BackupMode           *string `json:"BackupMode"`
			BackupType           *string `json:"BackupType"`
			Status               *string `json:"Status"`
			BackupSize           *int    `json:"BackupSize"`
			BackupLocationRef    *string `json:"BackupLocationRef"`
			RemotePath           *string `json:"RemotePath"`
			MD5                  *string `json:"MD5"`
		} `json:"DBBackup"`
		TotalCount *int `json:"TotalCount"`
		MaxRecords *int `json:"MaxRecords"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBBackupPolicyRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
}

func (r *ModifyDBBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBBackupPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverrideDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
}

func (r *OverrideDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverrideDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "OverrideDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OverrideDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *OverrideDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverrideDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	Engine               *string                             `json:"Engine,omitempty" name:"Engine"`
	EngineVersion        *string                             `json:"EngineVersion,omitempty" name:"EngineVersion"`
	DBParameterGroupName *string                             `json:"DBParameterGroupName,omitempty" name:"DBParameterGroupName"`
	Description          *string                             `json:"Description,omitempty" name:"Description"`
	Parameters           []*CreateDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *CreateDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Description          *string `json:"Description"`
			Parameters           struct {
				Autovacuum_analyze_scale_factor *int `json:"Autovacuum_analyze_scale_factor"`
			} `json:"Parameters"`
			Engine *string `json:"Engine"`
		} `json:"DBParameterGroup"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId   *string                             `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	DBParameterGroupName *string                             `json:"DBParameterGroupName,omitempty" name:"DBParameterGroupName"`
	Description          *string                             `json:"Description,omitempty" name:"Description"`
	Parameters           []*ModifyDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *ModifyDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Description          *string `json:"Description"`
			Parameters           struct {
				Autovacuum_analyze_scale_factor *int `json:"Autovacuum_analyze_scale_factor"`
			} `json:"Parameters"`
			Engine *string `json:"Engine"`
		} `json:"DBParameterGroup"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
}

func (r *DeleteDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
}

func (r *ResetDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Description          *string `json:"Description"`
			Parameters           struct {
				Autovacuum_analyze_scale_factor *int    `json:"Autovacuum_analyze_scale_factor"`
				Log_temp_files                  *int    `json:"Log_temp_files"`
				Autovacuum_vacuum_threshold     *int    `json:"Autovacuum_vacuum_threshold"`
				Vacuum_freeze_table_age         *int    `json:"Vacuum_freeze_table_age"`
				Autovacuum_freeze_max_age       *int    `json:"Autovacuum_freeze_max_age"`
				Wal_level                       *string `json:"Wal_level"`
				Autovacuum_vacuum_cost_limit    *int    `json:"Autovacuum_vacuum_cost_limit"`
				Autovacuum_vacuum_scale_factor  *int    `json:"Autovacuum_vacuum_scale_factor"`
				Track_activity_query_size       *int    `json:"Track_activity_query_size"`
				Autovacuum_max_workers          *int    `json:"Autovacuum_max_workers"`
				Checkpoint_timeout              *int    `json:"Checkpoint_timeout"`
				Wal_keep_segments               *int    `json:"Wal_keep_segments"`
				Autovacuum_vacuum_cost_delay    *int    `json:"Autovacuum_vacuum_cost_delay"`
				Autovacuum_naptime              *int    `json:"Autovacuum_naptime"`
				Autovacuum_analyze_threshold    *int    `json:"Autovacuum_analyze_threshold"`
				Default_statistics_target       *int    `json:"Default_statistics_target"`
				Log_autovacuum_min_duration     *int    `json:"Log_autovacuum_min_duration"`
			} `json:"Parameters"`
			Engine *string `json:"Engine"`
		} `json:"DBParameterGroup"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ResetDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	DBParameterGroupId *string `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	Marker             *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords         *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
	Source             *string `json:"Source,omitempty" name:"Source"`
}

func (r *DescribeDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParameterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBParameterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroups []struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion"`
			Description          *string `json:"Description"`
			Engine               *string `json:"Engine"`
			Parameters           struct {
				Autovacuum_analyze_scale_factor *int    `json:"Autovacuum_analyze_scale_factor"`
				Log_temp_files                  *int    `json:"Log_temp_files"`
				Autovacuum_vacuum_threshold     *int    `json:"Autovacuum_vacuum_threshold"`
				Vacuum_freeze_table_age         *int    `json:"Vacuum_freeze_table_age"`
				Autovacuum_freeze_max_age       *int    `json:"Autovacuum_freeze_max_age"`
				Wal_level                       *string `json:"Wal_level"`
				Autovacuum_vacuum_cost_limit    *int    `json:"Autovacuum_vacuum_cost_limit"`
				Autovacuum_vacuum_scale_factor  *int    `json:"Autovacuum_vacuum_scale_factor"`
				Track_activity_query_size       *int    `json:"Track_activity_query_size"`
				Autovacuum_max_workers          *int    `json:"Autovacuum_max_workers"`
				Checkpoint_timeout              *int    `json:"Checkpoint_timeout"`
				Wal_keep_segments               *int    `json:"Wal_keep_segments"`
				Autovacuum_vacuum_cost_delay    *int    `json:"Autovacuum_vacuum_cost_delay"`
				Autovacuum_naptime              *int    `json:"Autovacuum_naptime"`
				Autovacuum_analyze_threshold    *int    `json:"Autovacuum_analyze_threshold"`
				Default_statistics_target       *int    `json:"Default_statistics_target"`
				Log_autovacuum_min_duration     *int    `json:"Log_autovacuum_min_duration"`
			} `json:"Parameters"`
		} `json:"DBParameterGroups"`
		Source     *string `json:"Source"`
		MaxRecords *int    `json:"MaxRecords"`
		TotalCount *int    `json:"TotalCount"`
		Marker     *int    `json:"Marker"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineDefaultParametersRequest struct {
	*ksyunhttp.BaseRequest
	Engine        *string `json:"Engine,omitempty" name:"Engine"`
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *DescribeEngineDefaultParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineDefaultParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEngineDefaultParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineDefaultParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engine        *string `json:"Engine"`
		EngineVersion *string `json:"EngineVersion"`
		Parameters    struct {
			Autovacuum_analyze_scale_factor struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_analyze_scale_factor"`
			Autovacuum_vacuum_threshold struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_vacuum_threshold"`
			Log_temp_files struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Log_temp_files"`
			Autovacuum_freeze_max_age struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_freeze_max_age"`
			Vacuum_freeze_table_age struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Vacuum_freeze_table_age"`
			Wal_level struct {
				Default         *string `json:"Default"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
				Enums           []struct {
				} `json:"Enums"`
			} `json:"Wal_level"`
			Autovacuum_vacuum_cost_limit struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_vacuum_cost_limit"`
			Autovacuum_vacuum_scale_factor struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_vacuum_scale_factor"`
			Track_activity_query_size struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Track_activity_query_size"`
			Autovacuum_max_workers struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_max_workers"`
			Checkpoint_timeout struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Checkpoint_timeout"`
			Autovacuum_vacuum_cost_delay struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_vacuum_cost_delay"`
			Wal_keep_segments struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Wal_keep_segments"`
			Autovacuum_naptime struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_naptime"`
			Autovacuum_analyze_threshold struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Autovacuum_analyze_threshold"`
			Default_statistics_target struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Default_statistics_target"`
			Log_autovacuum_min_duration struct {
				Min             *int    `json:"Min"`
				Default         *int    `json:"Default"`
				Max             *int    `json:"Max"`
				RestartRequired *bool   `json:"RestartRequired"`
				Type            *string `json:"Type"`
			} `json:"Log_autovacuum_min_duration"`
		} `json:"Parameters"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeEngineDefaultParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineDefaultParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBInstanceParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		EngineVersion *string `json:"EngineVersion"`
		Parameters    struct {
			autovacuum_analyze_scale_factor *int    `json:"autovacuum_analyze_scale_factor"`
			log_temp_files                  *int    `json:"log_temp_files"`
			autovacuum_vacuum_threshold     *int    `json:"autovacuum_vacuum_threshold"`
			vacuum_freeze_table_age         *int    `json:"vacuum_freeze_table_age"`
			autovacuum_freeze_max_age       *int    `json:"autovacuum_freeze_max_age"`
			wal_level                       *string `json:"wal_level"`
			autovacuum_vacuum_cost_limit    *int    `json:"autovacuum_vacuum_cost_limit"`
			autovacuum_vacuum_scale_factor  *int    `json:"autovacuum_vacuum_scale_factor"`
			track_activity_query_size       *int    `json:"track_activity_query_size"`
			autovacuum_max_workers          *int    `json:"autovacuum_max_workers"`
			checkpoint_timeout              *int    `json:"checkpoint_timeout"`
			wal_keep_segments               *int    `json:"wal_keep_segments"`
			autovacuum_vacuum_cost_delay    *int    `json:"autovacuum_vacuum_cost_delay"`
			autovacuum_naptime              *int    `json:"autovacuum_naptime"`
			autovacuum_analyze_threshold    *int    `json:"autovacuum_analyze_threshold"`
			default_statistics_target       *int    `json:"default_statistics_target"`
			log_autovacuum_min_duration     *int    `json:"log_autovacuum_min_duration"`
		} `json:"Parameters"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstanceParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *RebootDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RebootDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebootDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RebootDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEngineVersionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeDBEngineVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEngineVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBEngineVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEngineVersionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engines struct {
			PostgerSQL []struct {
				Engine        *string `json:"Engine"`
				EngineVersion *string `json:"EngineVersion"`
			} `json:"PostgerSQL"`
		} `json:"Engines"`
	} `json:"Data"`
}

func (r *DescribeDBEngineVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEngineVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Port                 *int    `json:"Port,omitempty" name:"Port"`
}

func (r *AllocateDBInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateDBInstanceEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AllocateDBInstanceEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AllocateDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			EipPort                          *int    `json:"EipPort"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AllocateDBInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateDBInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *ReleaseDBInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseDBInstanceEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReleaseDBInstanceEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ReleaseDBInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseDBInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Mem                  *int    `json:"Mem,omitempty" name:"Mem"`
	Disk                 *int    `json:"Disk,omitempty" name:"Disk"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDBInstanceFromDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBBackupIdentifier *string   `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	DBInstanceType     *string   `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	DBInstanceName     *string   `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	BillType           *string   `json:"BillType,omitempty" name:"BillType"`
	Duration           *int      `json:"Duration,omitempty" name:"Duration"`
	DurationUnit       *string   `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone   []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	VpcId              *string   `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId           *string   `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *RestoreDBInstanceFromDBBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreDBInstanceFromDBBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreDBInstanceFromDBBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDBInstanceFromDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			SecurityGroupId       *string `json:"SecurityGroupId"`
			Vip                   *string `json:"Vip"`
			Port                  *int    `json:"Port"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZoneList  []struct {
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int  `json:"DiskUsed"`
			Audit                            *bool `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier"`
				DBBackupName         *string `json:"DBBackupName"`
				DBInstanceStatus     *string `json:"DBInstanceStatus"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreDBInstanceFromDBBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreDBInstanceFromDBBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceHARequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SwitchDBInstanceHARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDBInstanceHARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SwitchDBInstanceHARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceHAResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId"`
			Vip                    *string `json:"Vip"`
			Port                   *int    `json:"Port"`
			Engine                 *string `json:"Engine"`
			EngineVersion          *string `json:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId"`
			SubnetId               *string `json:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible"`
			BillType               *string `json:"BillType"`
			OrderType              *string `json:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType"`
				AzCode     *string `json:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int    `json:"DiskUsed"`
			InnerAzCode                      *string `json:"InnerAzCode"`
			Audit                            *bool   `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId   *string `json:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *SwitchDBInstanceHAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDBInstanceHAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceReadReplicaRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string   `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBInstanceName       *string   `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	AttachedVipId        *string   `json:"AttachedVipId,omitempty" name:"AttachedVipId"`
	BillType             *string   `json:"BillType,omitempty" name:"BillType"`
	Duration             *string   `json:"Duration,omitempty" name:"Duration"`
	DurationUnit         *string   `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone     []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Vip                  *string   `json:"Vip,omitempty" name:"Vip"`
	Mem                  *int      `json:"Mem,omitempty" name:"Mem"`
	Disk                 *int      `json:"Disk,omitempty" name:"Disk"`
}

func (r *CreateDBInstanceReadReplicaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBInstanceReadReplicaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDBInstanceReadReplicaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceReadReplicaResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBInstance struct {
			DBInstanceClass struct {
				Id      *string `json:"Id"`
				Iops    *int    `json:"Iops"`
				Vcpus   *int    `json:"Vcpus"`
				Disk    *int    `json:"Disk"`
				Ram     *int    `json:"Ram"`
				Mem     *int    `json:"Mem"`
				MaxConn *int    `json:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier  *string `json:"DBInstanceIdentifier"`
			DBInstanceName        *string `json:"DBInstanceName"`
			DBInstanceStatus      *string `json:"DBInstanceStatus"`
			DBInstanceType        *string `json:"DBInstanceType"`
			DBParameterGroupId    *string `json:"DBParameterGroupId"`
			PreferredBackupTime   *string `json:"PreferredBackupTime"`
			GroupId               *string `json:"GroupId"`
			Vip                   *string `json:"Vip"`
			Port                  *int    `json:"Port"`
			Engine                *string `json:"Engine"`
			EngineVersion         *string `json:"EngineVersion"`
			InstanceCreateTime    *string `json:"InstanceCreateTime"`
			MasterUserName        *string `json:"MasterUserName"`
			DatastoreVersionId    *string `json:"DatastoreVersionId"`
			VpcId                 *string `json:"VpcId"`
			SubnetId              *string `json:"SubnetId"`
			PubliclyAccessible    *bool   `json:"PubliclyAccessible"`
			BillType              *string `json:"BillType"`
			OrderType             *string `json:"OrderType"`
			MultiAvailabilityZone *bool   `json:"MultiAvailabilityZone"`
			AvailabilityZoneList  []struct {
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int  `json:"DiskUsed"`
			Audit                            *bool `json:"Audit"`
			ReadReplicaDBInstanceIdentifiers []struct {
				Vip                             *string `json:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier"`
				Id                              *string `json:"Id"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType"`
				DBInstanceStatus     *string `json:"DBInstanceStatus"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId"`
			ProjectName      *string `json:"ProjectName"`
			Region           *string `json:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId"`
			SecurityGroups   []struct {
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId"`
		} `json:"DBInstance"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDBInstanceReadReplicaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDBInstanceReadReplicaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPassword    *string `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
}

func (r *ModifyInstanceAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDatabasesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeInstanceDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		InstanceDatabases []struct {
			InstanceDatabaseName   *string `json:"InstanceDatabaseName"`
			InstanceDatabaseStatus *string `json:"InstanceDatabaseStatus"`
		} `json:"InstanceDatabases"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeInstanceDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceExtensionsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	KeyWord              *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *DescribeDBInstanceExtensionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceExtensionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceExtensionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceExtensionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Installed []struct {
			Comment          *string `json:"Comment"`
			DefaultVersion   *string `json:"DefaultVersion"`
			InstalledVersion *string `json:"InstalledVersion"`
			Name             *string `json:"Name"`
		} `json:"Installed"`
		Uninstalled []struct {
			Comment        *string `json:"Comment"`
			DefaultVersion *string `json:"DefaultVersion"`
			Name           *string `json:"Name"`
		} `json:"Uninstalled"`
	} `json:"Data"`
}

func (r *DescribeDBInstanceExtensionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceExtensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceExtensionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string   `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName *string   `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Operation            *string   `json:"Operation,omitempty" name:"Operation"`
	Extension            []*string `json:"Extension,omitempty" name:"Extension"`
}

func (r *ModifyDBInstanceExtensionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDBInstanceExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceExtensionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceExtensionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
