package v20190425

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateSecurityGroupSecurityGroupRule struct {
	SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol,omitempty" name:"SecurityGroupRuleProtocol"`
	SecurityGroupRuleName     *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}

type ModifySecurityGroupRuleSecurityGroupRule struct {
	SecurityGroupRuleId       *string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
	SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol,omitempty" name:"SecurityGroupRuleProtocol"`
	SecurityGroupRuleName     *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}

type CreateInstanceDatabaseInstanceDatabasePrivileges struct {
	InstanceAccountName *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	Privilege           *string `json:"Privilege,omitempty" name:"Privilege"`
}

type ModifyInstanceDatabasePrivilegesInstanceDatabasePrivileges struct {
	InstanceAccountName *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	Privilege           *string `json:"Privilege,omitempty" name:"Privilege"`
}

type CreateInstanceAccountInstanceAccountPrivileges struct {
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Privilege            *string `json:"Privilege,omitempty" name:"Privilege"`
}

type ModifyInstanceAccountPrivilegesInstanceAccountPrivileges struct {
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Privilege            *string `json:"Privilege,omitempty" name:"Privilege"`
}

type RestoreToCurInstanceSrcDatabases struct {
	DatabaseName  *string   `json:"DatabaseName,omitempty" name:"DatabaseName"`
	WholeDatabase *string   `json:"WholeDatabase,omitempty" name:"WholeDatabase"`
	TableNames    []*string `json:"TableNames,omitempty" name:"TableNames"`
}

type RestoreToCurInstanceDstDatabases struct {
	DatabaseName  *string   `json:"DatabaseName,omitempty" name:"DatabaseName"`
	WholeDatabase *string   `json:"WholeDatabase,omitempty" name:"WholeDatabase"`
	TableNames    []*string `json:"TableNames,omitempty" name:"TableNames"`
}

type CreateDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	Mem                 *int      `json:"Mem,omitempty" name:"Mem"`
	Disk                *int      `json:"Disk,omitempty" name:"Disk"`
	DBInstanceName      *string   `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	Engine              *string   `json:"Engine,omitempty" name:"Engine"`
	EngineVersion       *string   `json:"EngineVersion,omitempty" name:"EngineVersion"`
	DBInstanceType      *string   `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	MasterUserPassword  *string   `json:"MasterUserPassword,omitempty" name:"MasterUserPassword"`
	MasterUserName      *string   `json:"MasterUserName,omitempty" name:"MasterUserName"`
	VpcId               *string   `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId            *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	Port                *string   `json:"Port,omitempty" name:"Port"`
	PreferredBackupTime *string   `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	DBParameterGroupId  *string   `json:"DBParameterGroupId,omitempty" name:"DBParameterGroupId"`
	SecurityGroupId     *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	BillType            *string   `json:"BillType,omitempty" name:"BillType"`
	Duration            *string   `json:"Duration,omitempty" name:"Duration"`
	DurationUnit        *string   `json:"DurationUnit,omitempty" name:"DurationUnit"`
	AvailabilityZone    []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	ProjectId           *int      `json:"ProjectId,omitempty" name:"ProjectId"`
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
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime              *string   `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			VolumeSize                       *int      `json:"VolumeSize" name:"VolumeSize"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName                   *string   `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId               *string   `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList             []*string `json:"AvailabilityZoneList" name:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
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
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
	DBInstanceType       *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	DBInstanceStatus     *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`
	ExpiryDateLessThan   *string `json:"ExpiryDateLessThan,omitempty" name:"ExpiryDateLessThan"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
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
				Id    *string `json:"Id" name:"Id"`
				Vcpus *int    `json:"Vcpus" name:"Vcpus"`
				Disk  *int    `json:"Disk" name:"Disk"`
				Ram   *int    `json:"Ram" name:"Ram"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName                   *string   `json:"MasterUserName" name:"MasterUserName"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			OrderSource                      *string   `json:"OrderSource" name:"OrderSource"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone           *string   `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone            *string   `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			OrderUse                         *string   `json:"OrderUse" name:"OrderUse"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
			Tags                             []*string `json:"Tags" name:"Tags"`
		} `json:"Instances" name:"Instances"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
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
	Data struct {
		Instances []struct {
			DBInstanceClass struct {
				Id    *string `json:"Id" name:"Id"`
				Vcpus *int    `json:"Vcpus" name:"Vcpus"`
				Disk  *int    `json:"Disk" name:"Disk"`
				Ram   *int    `json:"Ram" name:"Ram"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName                   *string   `json:"MasterUserName" name:"MasterUserName"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			OrderSource                      *string   `json:"OrderSource" name:"OrderSource"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone           *string   `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone            *string   `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			OrderUse                         *string   `json:"OrderUse" name:"OrderUse"`
			Eip                              *string   `json:"Eip" name:"Eip"`
			EipPort                          *string   `json:"EipPort" name:"EipPort"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
			Tags                             []*string `json:"Tags" name:"Tags"`
		} `json:"Instances" name:"Instances"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
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
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			MiniVersion                      *string   `json:"MiniVersion" name:"MiniVersion"`
			SecurityGroups                   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			NetworkType                      *int      `json:"NetworkType" name:"NetworkType"`
			MaintenanceTime                  *string   `json:"MaintenanceTime" name:"MaintenanceTime"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BindInstances                    []*string `json:"BindInstances" name:"BindInstances"`
			ProxyNodeInfo                    []*string `json:"ProxyNodeInfo" name:"ProxyNodeInfo"`
			ProxyInfo                        []*string `json:"ProxyInfo" name:"ProxyInfo"`
			AutoSwitch                       *int      `json:"AutoSwitch" name:"AutoSwitch"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
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
	DBInstanceIdentifier     *string                                 `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
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
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
	SecurityGroupRuleAction *string                                     `json:"SecurityGroupRuleAction,omitempty" name:"SecurityGroupRuleAction"`
	SecurityGroupId         *string                                     `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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
			SecurityGroupId          *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string `json:"Created" name:"Created"`
			Instances                []struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip                  *string `json:"Vip" name:"Vip"`
				Created              *string `json:"Created" name:"Created"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId       *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName     *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created                   *string `json:"Created" name:"Created"`
			} `json:"SecurityGroupRules"`
		} `json:"SecurityGroups" name:"SecurityGroups"`
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
	SecurityGroupId       *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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

type DescribeCollationsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeCollationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCollationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCollationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCollationsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Collations []struct {
		} `json:"Collations" name:"Collations"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeCollationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCollationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier        *string                                             `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName        *string                                             `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabaseCollation   *string                                             `json:"InstanceDatabaseCollation,omitempty" name:"InstanceDatabaseCollation"`
	InstanceDatabaseDescription *string                                             `json:"InstanceDatabaseDescription,omitempty" name:"InstanceDatabaseDescription"`
	InstanceDatabasePrivileges  []*CreateInstanceDatabaseInstanceDatabasePrivileges `json:"InstanceDatabasePrivileges,omitempty" name:"InstanceDatabasePrivileges"`
}

func (r *CreateInstanceDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateInstanceDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInstanceDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabasePrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string                                                       `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName       *string                                                       `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabasePrivileges []*ModifyInstanceDatabasePrivilegesInstanceDatabasePrivileges `json:"InstanceDatabasePrivileges,omitempty" name:"InstanceDatabasePrivileges"`
}

func (r *ModifyInstanceDatabasePrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabasePrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceDatabasePrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabasePrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabasePrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabasePrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDatabasesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
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
			InstanceDatabaseName        *string `json:"InstanceDatabaseName" name:"InstanceDatabaseName"`
			InstanceDatabaseCollation   *string `json:"InstanceDatabaseCollation" name:"InstanceDatabaseCollation"`
			InstanceDatabaseDescription *string `json:"InstanceDatabaseDescription" name:"InstanceDatabaseDescription"`
			InstanceDatabaseStatus      *string `json:"InstanceDatabaseStatus" name:"InstanceDatabaseStatus"`
			Created                     *string `json:"Created" name:"Created"`
			InstanceDatabasePrivileges  []struct {
				InstanceAccountName *string `json:"InstanceAccountName" name:"InstanceAccountName"`
				Privilege           *string `json:"Privilege" name:"Privilege"`
			} `json:"InstanceDatabasePrivileges"`
		} `json:"InstanceDatabases" name:"InstanceDatabases"`
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

type CreateInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string                                           `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string                                           `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPassword    *string                                           `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string                                           `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
	InstanceAccountPrivileges  []*CreateInstanceAccountInstanceAccountPrivileges `json:"InstanceAccountPrivileges,omitempty" name:"InstanceAccountPrivileges"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeInstanceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Accounts []struct {
			InstanceAccountName        *string `json:"InstanceAccountName" name:"InstanceAccountName"`
			InstanceAccountDescription *string `json:"InstanceAccountDescription" name:"InstanceAccountDescription"`
			Created                    *string `json:"Created" name:"Created"`
			InstanceAccountType        *string `json:"InstanceAccountType" name:"InstanceAccountType"`
			InstanceAccountPrivileges  []struct {
				InstanceDatabaseName *string `json:"InstanceDatabaseName" name:"InstanceDatabaseName"`
				Privilege            *string `json:"Privilege" name:"Privilege"`
			} `json:"InstanceAccountPrivileges"`
		} `json:"Accounts" name:"Accounts"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeInstanceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier       *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName        *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountDescription *string `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
	InstanceAccountPassword    *string `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
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

type ModifyInstanceAccountPrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier      *string                                                     `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName       *string                                                     `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPrivileges []*ModifyInstanceAccountPrivilegesInstanceAccountPrivileges `json:"InstanceAccountPrivileges,omitempty" name:"InstanceAccountPrivileges"`
}

func (r *ModifyInstanceAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountPrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceAccountName  *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
}

func (r *DeleteInstanceDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteInstanceDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabaseInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier        *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName        *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabaseDescription *string `json:"InstanceDatabaseDescription,omitempty" name:"InstanceDatabaseDescription"`
}

func (r *ModifyInstanceDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceDatabaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabaseInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabaseInfoResponse) FromJsonString(s string) error {
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
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime              *string   `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			VolumeSize                       *int      `json:"VolumeSize" name:"VolumeSize"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName                   *string   `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId               *string   `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList             []*string `json:"AvailabilityZoneList" name:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
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

type RestoreDBInstanceFromDBBackupRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceName     *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
	DBBackupIdentifier *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	DBInstanceType     *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	ProjectId          *string `json:"ProjectId,omitempty" name:"ProjectId"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Duration           *int    `json:"Duration,omitempty" name:"Duration"`
	DurationUnit       *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	Port               *int    `json:"Port,omitempty" name:"Port"`
	BillType           *string `json:"BillType,omitempty" name:"BillType"`
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
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime              *string   `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			DatastoreVersionId               *string   `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList             []*string `json:"AvailabilityZoneList" name:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			DBSource                         struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName       *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType       *string `json:"DBInstanceType" name:"DBInstanceType"`
				DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
				DBBackupName         *string `json:"DBBackupName" name:"DBBackupName"`
			} `json:"DBSource"`
			ProductId        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string   `json:"ProjectName" name:"ProjectName"`
			Region           *string   `json:"Region" name:"Region"`
			ServiceEndTime   *string   `json:"ServiceEndTime" name:"ServiceEndTime"`
			ServiceStartTime *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"DBInstance" name:"DBInstance"`
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
		DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
		DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
		Engine               *string `json:"Engine" name:"Engine"`
		EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
		BackupCreateTime     *string `json:"BackupCreateTime" name:"BackupCreateTime"`
		BackupUpdatedTime    *string `json:"BackupUpdatedTime" name:"BackupUpdatedTime"`
		DBBackupName         *string `json:"DBBackupName" name:"DBBackupName"`
		Description          *string `json:"Description" name:"Description"`
		BackupType           *string `json:"BackupType" name:"BackupType"`
		Status               *string `json:"Status" name:"Status"`
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
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
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
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
	BackupType           *string `json:"BackupType,omitempty" name:"BackupType"`
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
			DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBBackupIdentifier   *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
			Engine               *string `json:"Engine" name:"Engine"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			BackupCreateTime     *string `json:"BackupCreateTime" name:"BackupCreateTime"`
			BackupUpdatedTime    *string `json:"BackupUpdatedTime" name:"BackupUpdatedTime"`
			DBBackupName         *string `json:"DBBackupName" name:"DBBackupName"`
			Description          *string `json:"Description" name:"Description"`
			BackupMode           *string `json:"BackupMode" name:"BackupMode"`
			BackupType           *string `json:"BackupType" name:"BackupType"`
			Status               *string `json:"Status" name:"Status"`
			BackupSize           *int    `json:"BackupSize" name:"BackupSize"`
			BackupLocationRef    *string `json:"BackupLocationRef" name:"BackupLocationRef"`
			RemotePath           *string `json:"RemotePath" name:"RemotePath"`
		} `json:"DBBackup" name:"DBBackup"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
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
	DBInstanceIdentifier   *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	PreferredBackupTime    *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	IncrementalBackupCycle *string `json:"IncrementalBackupCycle,omitempty" name:"IncrementalBackupCycle"`
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

type AllocateDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Port                 *string `json:"Port,omitempty" name:"Port"`
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
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"DBInstance" name:"DBInstance"`
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
				Id      *string `json:"Id" name:"Id"`
				Iops    *int    `json:"Iops" name:"Iops"`
				Vcpus   *int    `json:"Vcpus" name:"Vcpus"`
				Disk    *int    `json:"Disk" name:"Disk"`
				Ram     *int    `json:"Ram" name:"Ram"`
				Mem     *int    `json:"Mem" name:"Mem"`
				MaxConn *int    `json:"MaxConn" name:"MaxConn"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			Vip                    *string `json:"Vip" name:"Vip"`
			Port                   *int    `json:"Port" name:"Port"`
			Engine                 *string `json:"Engine" name:"Engine"`
			EngineVersion          *string `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime     *string `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName         *string `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId     *string `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                  *string `json:"VpcId" name:"VpcId"`
			SubnetId               *string `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible     *bool   `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType               *string `json:"BillType" name:"BillType"`
			OrderType              *string `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone  *bool   `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone *string `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone  *string `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			AvailabilityZoneList   []struct {
				MemberType *string `json:"MemberType" name:"MemberType"`
				AzCode     *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed                         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode                      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit                            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat                      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			ServiceStartTime                 *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId                       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups                   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"DBInstance" name:"DBInstance"`
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

type ListSlowLogsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Time                 *string `json:"Time,omitempty" name:"Time"`
	OrderBy              *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListSlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSlowLogsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SlowLogs []struct {
			TotalElapsedTime    *int    `json:"TotalElapsedTime" name:"TotalElapsedTime"`
			TotalExecutionCount *int    `json:"TotalExecutionCount" name:"TotalExecutionCount"`
			TotalLogicalReads   *int    `json:"TotalLogicalReads" name:"TotalLogicalReads"`
			TotalPhysicalReads  *int    `json:"TotalPhysicalReads" name:"TotalPhysicalReads"`
			SqlText             *string `json:"SqlText" name:"SqlText"`
			CollectTime         *string `json:"CollectTime" name:"CollectTime"`
		} `json:"SlowLogs" name:"SlowLogs"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
	} `json:"Data"`
}

func (r *ListSlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListErrorLogsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	StartTime            *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime              *string `json:"EndTime,omitempty" name:"EndTime"`
	Marker               *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords           *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *ListErrorLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListErrorLogsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		ErrorLogs []struct {
			SqlText     *string `json:"SqlText" name:"SqlText"`
			CollectTime *string `json:"CollectTime" name:"CollectTime"`
		} `json:"ErrorLogs" name:"ErrorLogs"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListErrorLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListErrorLogsResponse) FromJsonString(s string) error {
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
				Id    *string `json:"Id" name:"Id"`
				Vcpus *int    `json:"Vcpus" name:"Vcpus"`
				Disk  *int    `json:"Disk" name:"Disk"`
				Ram   *int    `json:"Ram" name:"Ram"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName                   *string   `json:"MasterUserName" name:"MasterUserName"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			OrderSource                      *string   `json:"OrderSource" name:"OrderSource"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone           *string   `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone            *string   `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			OrderUse                         *string   `json:"OrderUse" name:"OrderUse"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
			Tags                             []*string `json:"Tags" name:"Tags"`
		} `json:"Instances" name:"Instances"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
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

type DescribeImportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Keyword              *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeImportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportTaskResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		ImportTasks []struct {
			ImportTaskId     *string `json:"ImportTaskId" name:"ImportTaskId"`
			DBName           *string `json:"DBName" name:"DBName"`
			Created          *string `json:"Created" name:"Created"`
			StartTime        *string `json:"StartTime" name:"StartTime"`
			EndTime          *string `json:"EndTime" name:"EndTime"`
			ImportTaskStatus *string `json:"ImportTaskStatus" name:"ImportTaskStatus"`
		} `json:"ImportTasks" name:"ImportTasks"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportFileRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	ImportTaskId         *string `json:"ImportTaskId,omitempty" name:"ImportTaskId"`
}

func (r *DescribeImportFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeImportFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportFileResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		ImportFile []struct {
			ImportFileName *string `json:"ImportFileName" name:"ImportFileName"`
			ImportFileSize *int    `json:"ImportFileSize" name:"ImportFileSize"`
			StartTime      *string `json:"StartTime" name:"StartTime"`
			EndTime        *string `json:"EndTime" name:"EndTime"`
			ImportFileType *string `json:"ImportFileType" name:"ImportFileType"`
		} `json:"ImportFile" name:"ImportFile"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeImportFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	KS3Region            *string `json:"KS3Region,omitempty" name:"KS3Region"`
	KS3Bucket            *string `json:"KS3Bucket,omitempty" name:"KS3Bucket"`
	KS3Key               *string `json:"KS3Key,omitempty" name:"KS3Key"`
	KS3FileSize          *string `json:"KS3FileSize,omitempty" name:"KS3FileSize"`
	ImportTaskId         *string `json:"ImportTaskId,omitempty" name:"ImportTaskId"`
	DBName               *string `json:"DBName,omitempty" name:"DBName"`
}

func (r *CreateImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateImportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateImportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FinishImportTaskRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	ImportTaskId         *string `json:"ImportTaskId,omitempty" name:"ImportTaskId"`
}

func (r *FinishImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FinishImportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "FinishImportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FinishImportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *FinishImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FinishImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRestorableTimeRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBInstanceRestorableTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRestorableTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBInstanceRestorableTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRestorableTimeResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		RestorableTime struct {
			Begin *string `json:"Begin" name:"Begin"`
			End   *string `json:"End" name:"End"`
		} `json:"RestorableTime" name:"RestorableTime"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBInstanceRestorableTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRestorableTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHistoryDatabaseInfoRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	RestorableTime       *string `json:"RestorableTime,omitempty" name:"RestorableTime"`
}

func (r *GetHistoryDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHistoryDatabaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetHistoryDatabaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetHistoryDatabaseInfoResponse struct {
	*ksyunhttp.BaseResponse
	Databases []struct {
		DatabaseName *string `json:"DatabaseName" name:"DatabaseName"`
		TableNames   []struct {
		} `json:"TableNames" name:"TableNames"`
	} `json:"Databases"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetHistoryDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHistoryDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreToCurInstanceRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string                             `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	DBBackupIdentifier   *string                             `json:"DBBackupIdentifier,omitempty" name:"DBBackupIdentifier"`
	RestorableTime       *string                             `json:"RestorableTime,omitempty" name:"RestorableTime"`
	SrcDatabases         []*RestoreToCurInstanceSrcDatabases `json:"SrcDatabases,omitempty" name:"SrcDatabases"`
	DstDatabases         []*RestoreToCurInstanceDstDatabases `json:"DstDatabases,omitempty" name:"DstDatabases"`
}

func (r *RestoreToCurInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreToCurInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestoreToCurInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreToCurInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RestoreToCurInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreToCurInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabaseNameRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier    *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName    *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabaseNameNew *string `json:"InstanceDatabaseNameNew,omitempty" name:"InstanceDatabaseNameNew"`
}

func (r *ModifyInstanceDatabaseNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabaseNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyInstanceDatabaseNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceDatabaseNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabaseNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabaseNameResponse) FromJsonString(s string) error {
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
		Instances []struct {
			DBInstanceClass struct {
				Id    *string `json:"Id" name:"Id"`
				Vcpus *int    `json:"Vcpus" name:"Vcpus"`
				Disk  *int    `json:"Disk" name:"Disk"`
				Ram   *int    `json:"Ram" name:"Ram"`
			} `json:"DBInstanceClass"`
			DBInstanceIdentifier             *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName                   *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus                 *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType                   *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId               *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			GroupId                          *string   `json:"GroupId" name:"GroupId"`
			Vip                              *string   `json:"Vip" name:"Vip"`
			Port                             *int      `json:"Port" name:"Port"`
			Engine                           *string   `json:"Engine" name:"Engine"`
			EngineVersion                    *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime               *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName                   *string   `json:"MasterUserName" name:"MasterUserName"`
			VpcId                            *string   `json:"VpcId" name:"VpcId"`
			SubnetId                         *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible               *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			BillType                         *string   `json:"BillType" name:"BillType"`
			OrderType                        *string   `json:"OrderType" name:"OrderType"`
			OrderSource                      *string   `json:"OrderSource" name:"OrderSource"`
			MultiAvailabilityZone            *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			MasterAvailabilityZone           *string   `json:"MasterAvailabilityZone" name:"MasterAvailabilityZone"`
			SlaveAvailabilityZone            *string   `json:"SlaveAvailabilityZone" name:"SlaveAvailabilityZone"`
			ProductId                        *string   `json:"ProductId" name:"ProductId"`
			OrderUse                         *string   `json:"OrderUse" name:"OrderUse"`
			SupportIPV6                      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			ProjectId                        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName                      *string   `json:"ProjectName" name:"ProjectName"`
			Region                           *string   `json:"Region" name:"Region"`
			BillTypeId                       *int      `json:"BillTypeId" name:"BillTypeId"`
			Tags                             []*string `json:"Tags" name:"Tags"`
		} `json:"Instances" name:"Instances"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
		Marker     *int `json:"Marker" name:"Marker"`
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

type DescribeDBBackupPolicyRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeDBBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDBBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupPolicyResponse struct {
	*ksyunhttp.BaseResponse
	BackupConfig struct {
		ExpireAfter            *int    `json:"ExpireAfter" name:"ExpireAfter"`
		AutobackupAt           *string `json:"AutobackupAt" name:"AutobackupAt"`
		Duration               *int    `json:"Duration" name:"Duration"`
		IncrementalBackupCycle *int    `json:"IncrementalBackupCycle" name:"IncrementalBackupCycle"`
		FullBackupCycle        *string `json:"FullBackupCycle" name:"FullBackupCycle"`
		GroupId                *string `json:"GroupId" name:"GroupId"`
		BinlogExpireAfter      *int    `json:"BinlogExpireAfter" name:"BinlogExpireAfter"`
		HighFrequencyBackup    *bool   `json:"HighFrequencyBackup" name:"HighFrequencyBackup"`
	} `json:"BackupConfig"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDBBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
