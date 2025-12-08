package v20181225
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
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
			DBInstanceIdentifier  *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName        *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus      *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType        *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId    *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime   *string   `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId               *string   `json:"GroupId" name:"GroupId"`
			Vip                   *string   `json:"Vip" name:"Vip"`
			Port                  *int      `json:"Port" name:"Port"`
			Engine                *string   `json:"Engine" name:"Engine"`
			EngineVersion         *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime    *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			MasterUserName        *string   `json:"MasterUserName" name:"MasterUserName"`
			DatastoreVersionId    *string   `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                 *string   `json:"VpcId" name:"VpcId"`
			SubnetId              *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible    *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType              *string   `json:"BillType" name:"BillType"`
			OrderType             *string   `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList  []*string `json:"AvailabilityZoneList" name:"AvailabilityZoneList"`
			DiskUsed              *int      `json:"DiskUsed" name:"DiskUsed"`
			Audit                 *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId             *string   `json:"ProductId" name:"ProductId"`
			ProductWhat           *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId             *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName           *string   `json:"ProjectName" name:"ProjectName"`
			Region                *string   `json:"Region" name:"Region"`
			ServiceStartTime      *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId            *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups        []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6           *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId            *int      `json:"BillTypeId" name:"BillTypeId"`
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

type DescribeDBInstancesResponse struct {
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
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int    `json:"DiskUsed" name:"DiskUsed"`
			Eip              *string `json:"Eip" name:"Eip"`
			EipPort          *int    `json:"EipPort" name:"EipPort"`
			InnerAzCode      *string `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool   `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			MiniVersion      *string `json:"MiniVersion" name:"MiniVersion"`
			SecurityGroups   []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			NetworkType *int      `json:"NetworkType" name:"NetworkType"`
			SupportIPV6 *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BindInstances []*string `json:"BindInstances" name:"BindInstances"`
			ProxyNodeInfo []*string `json:"ProxyNodeInfo" name:"ProxyNodeInfo"`
			ProxyInfo   []*string `json:"ProxyInfo" name:"ProxyInfo"`
			AutoSwitch  *int      `json:"AutoSwitch" name:"AutoSwitch"`
			BillTypeId  *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
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

type StatisticDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Statistic struct {
			RUNNINGTASK *int `json:"RUNNING_TASK" name:"RUNNING_TASK"`
			ACTIVE      *int `json:"ACTIVE" name:"ACTIVE"`
			INVALID     *int `json:"INVALID" name:"INVALID"`
		} `json:"Statistic" name:"Statistic"`
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
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int    `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool   `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
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
	DBInstanceIdentifier     []*string                               `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	SecurityGroupRule        []*CreateSecurityGroupSecurityGroupRule `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created               *string `json:"Created" name:"Created"`
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
				DBInstanceName *string `json:"DBInstanceName" name:"DBInstanceName"`
				Vip            *string `json:"Vip" name:"Vip"`
				Created        *string `json:"Created" name:"Created"`
				DBInstanceType *string `json:"DBInstanceType" name:"DBInstanceType"`
			} `json:"Instances"`
			SecurityGroupRules []struct {
				SecurityGroupRuleId *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created             *string `json:"Created" name:"Created"`
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
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created               *string `json:"Created" name:"Created"`
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

type ModifySecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created               *string `json:"Created" name:"Created"`
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
				SecurityGroupRuleId *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created             *string `json:"Created" name:"Created"`
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
	SecurityGroupId         *string                                     `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupRuleAction *string                                     `json:"SecurityGroupRuleAction,omitempty" name:"SecurityGroupRuleAction"`
	SecurityGroupRule       []*ModifySecurityGroupRuleSecurityGroupRule `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
}

func (r *ModifySecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifySecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		SecurityGroups []struct {
			SecurityGroupId          *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName        *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
			SecurityGroupDescription *string   `json:"SecurityGroupDescription" name:"SecurityGroupDescription"`
			Created                  *string   `json:"Created" name:"Created"`
			Instances                []*string `json:"Instances" name:"Instances"`
			SecurityGroupRules       []struct {
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created               *string `json:"Created" name:"Created"`
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
				SecurityGroupRuleId   *string `json:"SecurityGroupRuleId" name:"SecurityGroupRuleId"`
				SecurityGroupRuleName *string `json:"SecurityGroupRuleName" name:"SecurityGroupRuleName"`
				SecurityGroupRuleProtocol *string `json:"SecurityGroupRuleProtocol" name:"SecurityGroupRuleProtocol"`
				Created               *string `json:"Created" name:"Created"`
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
	SecurityGroupRuleId   *string `json:"SecurityGroupRuleId,omitempty" name:"SecurityGroupRuleId"`
	SecurityGroupRuleName *string `json:"SecurityGroupRuleName,omitempty" name:"SecurityGroupRuleName"`
}

func (r *ModifySecurityGroupRuleNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DescribeDBLogFilesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DescribeDBLogFiles []struct {
			LogFileName *string `json:"LogFileName" name:"LogFileName"`
			Size        *int    `json:"Size" name:"Size"`
			RawSize     *int    `json:"RawSize" name:"RawSize"`
			Date        *string `json:"Date" name:"Date"`
			StartTime   *string `json:"StartTime" name:"StartTime"`
			EndTime     *string `json:"EndTime" name:"EndTime"`
		} `json:"DescribeDBLogFiles" name:"DescribeDBLogFiles"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		Marker *int `json:"Marker" name:"Marker"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
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

type CreateDBBackupResponse struct {
	*ksyunhttp.BaseResponse
	DBBackup struct {
		DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
		DBBackupIdentifier *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
		Engine             *string `json:"Engine" name:"Engine"`
		EngineVersion      *string `json:"EngineVersion" name:"EngineVersion"`
		BackupCreateTime   *string `json:"BackupCreateTime" name:"BackupCreateTime"`
		BackupUpdatedTime  *string `json:"BackupUpdatedTime" name:"BackupUpdatedTime"`
		DBBackupName       *string `json:"DBBackupName" name:"DBBackupName"`
		Description        *string `json:"Description" name:"Description"`
		BackupType         *string `json:"BackupType" name:"BackupType"`
		Status             *string `json:"Status" name:"Status"`
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
			BackupMode           *string `json:"BackupMode" name:"BackupMode"`
			BackupType           *string `json:"BackupType" name:"BackupType"`
			Status               *string `json:"Status" name:"Status"`
			BackupSize           *int    `json:"BackupSize" name:"BackupSize"`
			BackupLocationRef    *string `json:"BackupLocationRef" name:"BackupLocationRef"`
			RemotePath           *string `json:"RemotePath" name:"RemotePath"`
			MD5                  *string `json:"MD5" name:"MD5"`
		} `json:"DBBackup" name:"DBBackup"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
		MaxRecords *int `json:"MaxRecords" name:"MaxRecords"`
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
			DBInstanceIdentifier   *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName         *string `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus       *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType         *string `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId     *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime    *string `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId                *string `json:"GroupId" name:"GroupId"`
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int    `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool   `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
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

type CreateDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Description          *string `json:"Description" name:"Description"`
			Parameters           struct {
				AutovacuumAnalyzeScaleFactor *int `json:"AutovacuumAnalyzeScaleFactor" name:"AutovacuumAnalyzeScaleFactor"`
			} `json:"Parameters"`
			Engine *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroup" name:"DBParameterGroup"`
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

type ModifyDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Description          *string `json:"Description" name:"Description"`
			Parameters           struct {
				AutovacuumAnalyzeScaleFactor *int `json:"AutovacuumAnalyzeScaleFactor" name:"AutovacuumAnalyzeScaleFactor"`
			} `json:"Parameters"`
			Engine *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroup" name:"DBParameterGroup"`
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

type ResetDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroup struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Description          *string `json:"Description" name:"Description"`
			Parameters           struct {
				AutovacuumAnalyzeScaleFactor *int    `json:"AutovacuumAnalyzeScaleFactor" name:"AutovacuumAnalyzeScaleFactor"`
				LogTempFiles                 *int    `json:"LogTempFiles" name:"LogTempFiles"`
				AutovacuumVacuumThreshold    *int    `json:"AutovacuumVacuumThreshold" name:"AutovacuumVacuumThreshold"`
				VacuumFreezeTableAge         *int    `json:"VacuumFreezeTableAge" name:"VacuumFreezeTableAge"`
				AutovacuumFreezeMaxAge       *int    `json:"AutovacuumFreezeMaxAge" name:"AutovacuumFreezeMaxAge"`
				WalLevel                     *string `json:"WalLevel" name:"WalLevel"`
				AutovacuumVacuumCostLimit    *int    `json:"AutovacuumVacuumCostLimit" name:"AutovacuumVacuumCostLimit"`
				AutovacuumVacuumScaleFactor  *int    `json:"AutovacuumVacuumScaleFactor" name:"AutovacuumVacuumScaleFactor"`
				TrackActivityQuerySize       *int    `json:"TrackActivityQuerySize" name:"TrackActivityQuerySize"`
				AutovacuumMaxWorkers         *int    `json:"AutovacuumMaxWorkers" name:"AutovacuumMaxWorkers"`
				CheckpointTimeout            *int    `json:"CheckpointTimeout" name:"CheckpointTimeout"`
				WalKeepSegments              *int    `json:"WalKeepSegments" name:"WalKeepSegments"`
				AutovacuumVacuumCostDelay    *int    `json:"AutovacuumVacuumCostDelay" name:"AutovacuumVacuumCostDelay"`
				AutovacuumNaptime            *int    `json:"AutovacuumNaptime" name:"AutovacuumNaptime"`
				AutovacuumAnalyzeThreshold   *int    `json:"AutovacuumAnalyzeThreshold" name:"AutovacuumAnalyzeThreshold"`
				DefaultStatisticsTarget      *int    `json:"DefaultStatisticsTarget" name:"DefaultStatisticsTarget"`
				LogAutovacuumMinDuration     *int    `json:"LogAutovacuumMinDuration" name:"LogAutovacuumMinDuration"`
			} `json:"Parameters"`
			Engine *string `json:"Engine" name:"Engine"`
		} `json:"DBParameterGroup" name:"DBParameterGroup"`
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

type DescribeDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		DBParameterGroups []struct {
			DBParameterGroupId   *string `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			DBParameterGroupName *string `json:"DBParameterGroupName" name:"DBParameterGroupName"`
			EngineVersion        *string `json:"EngineVersion" name:"EngineVersion"`
			Description          *string `json:"Description" name:"Description"`
			Engine               *string `json:"Engine" name:"Engine"`
			Parameters           struct {
				AutovacuumAnalyzeScaleFactor *int    `json:"AutovacuumAnalyzeScaleFactor" name:"AutovacuumAnalyzeScaleFactor"`
				LogTempFiles                 *int    `json:"LogTempFiles" name:"LogTempFiles"`
				AutovacuumVacuumThreshold    *int    `json:"AutovacuumVacuumThreshold" name:"AutovacuumVacuumThreshold"`
				VacuumFreezeTableAge         *int    `json:"VacuumFreezeTableAge" name:"VacuumFreezeTableAge"`
				AutovacuumFreezeMaxAge       *int    `json:"AutovacuumFreezeMaxAge" name:"AutovacuumFreezeMaxAge"`
				WalLevel                     *string `json:"WalLevel" name:"WalLevel"`
				AutovacuumVacuumCostLimit    *int    `json:"AutovacuumVacuumCostLimit" name:"AutovacuumVacuumCostLimit"`
				AutovacuumVacuumScaleFactor  *int    `json:"AutovacuumVacuumScaleFactor" name:"AutovacuumVacuumScaleFactor"`
				TrackActivityQuerySize       *int    `json:"TrackActivityQuerySize" name:"TrackActivityQuerySize"`
				AutovacuumMaxWorkers         *int    `json:"AutovacuumMaxWorkers" name:"AutovacuumMaxWorkers"`
				CheckpointTimeout            *int    `json:"CheckpointTimeout" name:"CheckpointTimeout"`
				WalKeepSegments              *int    `json:"WalKeepSegments" name:"WalKeepSegments"`
				AutovacuumVacuumCostDelay    *int    `json:"AutovacuumVacuumCostDelay" name:"AutovacuumVacuumCostDelay"`
				AutovacuumNaptime            *int    `json:"AutovacuumNaptime" name:"AutovacuumNaptime"`
				AutovacuumAnalyzeThreshold   *int    `json:"AutovacuumAnalyzeThreshold" name:"AutovacuumAnalyzeThreshold"`
				DefaultStatisticsTarget      *int    `json:"DefaultStatisticsTarget" name:"DefaultStatisticsTarget"`
				LogAutovacuumMinDuration     *int    `json:"LogAutovacuumMinDuration" name:"LogAutovacuumMinDuration"`
			} `json:"Parameters"`
		} `json:"DBParameterGroups" name:"DBParameterGroups"`
		Source     *string `json:"Source" name:"Source"`
		MaxRecords *int    `json:"MaxRecords" name:"MaxRecords"`
		TotalCount *int    `json:"TotalCount" name:"TotalCount"`
		Marker     *int    `json:"Marker" name:"Marker"`
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

type DescribeEngineDefaultParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engine     *string `json:"Engine" name:"Engine"`
		EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
		Parameters struct {
			AutovacuumAnalyzeScaleFactor struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumAnalyzeScaleFactor"`
			AutovacuumVacuumThreshold struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumVacuumThreshold"`
			LogTempFiles struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"LogTempFiles"`
			AutovacuumFreezeMaxAge struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumFreezeMaxAge"`
			VacuumFreezeTableAge struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"VacuumFreezeTableAge"`
			WalLevel struct {
				DefaultField    *string   `json:"DefaultField" name:"DefaultField"`
				RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string   `json:"TypeField" name:"TypeField"`
				Enums           []*string `json:"Enums" name:"Enums"`
			} `json:"WalLevel"`
			AutovacuumVacuumCostLimit struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumVacuumCostLimit"`
			AutovacuumVacuumScaleFactor struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumVacuumScaleFactor"`
			TrackActivityQuerySize struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"TrackActivityQuerySize"`
			AutovacuumMaxWorkers struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumMaxWorkers"`
			CheckpointTimeout struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"CheckpointTimeout"`
			AutovacuumVacuumCostDelay struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumVacuumCostDelay"`
			WalKeepSegments struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"WalKeepSegments"`
			AutovacuumNaptime struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumNaptime"`
			AutovacuumAnalyzeThreshold struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"AutovacuumAnalyzeThreshold"`
			DefaultStatisticsTarget struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"DefaultStatisticsTarget"`
			LogAutovacuumMinDuration struct {
				Min             *int    `json:"Min" name:"Min"`
				DefaultField    *int    `json:"DefaultField" name:"DefaultField"`
				Max             *int    `json:"Max" name:"Max"`
				RestartRequired *bool   `json:"RestartRequired" name:"RestartRequired"`
				TypeField       *string `json:"TypeField" name:"TypeField"`
			} `json:"LogAutovacuumMinDuration"`
		} `json:"Parameters" name:"Parameters"`
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

type DescribeDBInstanceParametersResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
		Parameters struct {
			AutovacuumAnalyzeScaleFactor *int    `json:"autovacuum_analyze_scale_factor" name:"autovacuum_analyze_scale_factor"`
			LogTempFiles                 *int    `json:"log_temp_files" name:"log_temp_files"`
			AutovacuumVacuumThreshold    *int    `json:"autovacuum_vacuum_threshold" name:"autovacuum_vacuum_threshold"`
			VacuumFreezeTableAge         *int    `json:"vacuum_freeze_table_age" name:"vacuum_freeze_table_age"`
			AutovacuumFreezeMaxAge       *int    `json:"autovacuum_freeze_max_age" name:"autovacuum_freeze_max_age"`
			WalLevel                     *string `json:"wal_level" name:"wal_level"`
			AutovacuumVacuumCostLimit    *int    `json:"autovacuum_vacuum_cost_limit" name:"autovacuum_vacuum_cost_limit"`
			AutovacuumVacuumScaleFactor  *int    `json:"autovacuum_vacuum_scale_factor" name:"autovacuum_vacuum_scale_factor"`
			TrackActivityQuerySize       *int    `json:"track_activity_query_size" name:"track_activity_query_size"`
			AutovacuumMaxWorkers         *int    `json:"autovacuum_max_workers" name:"autovacuum_max_workers"`
			CheckpointTimeout            *int    `json:"checkpoint_timeout" name:"checkpoint_timeout"`
			WalKeepSegments              *int    `json:"wal_keep_segments" name:"wal_keep_segments"`
			AutovacuumVacuumCostDelay    *int    `json:"autovacuum_vacuum_cost_delay" name:"autovacuum_vacuum_cost_delay"`
			AutovacuumNaptime            *int    `json:"autovacuum_naptime" name:"autovacuum_naptime"`
			AutovacuumAnalyzeThreshold   *int    `json:"autovacuum_analyze_threshold" name:"autovacuum_analyze_threshold"`
			DefaultStatisticsTarget      *int    `json:"default_statistics_target" name:"default_statistics_target"`
			LogAutovacuumMinDuration     *int    `json:"log_autovacuum_min_duration" name:"log_autovacuum_min_duration"`
		} `json:"Parameters" name:"Parameters"`
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

type RebootDBInstanceResponse struct {
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string   `json:"ProjectName" name:"ProjectName"`
			Region           *string   `json:"Region" name:"Region"`
			ServiceStartTime *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"DBInstance" name:"DBInstance"`
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

type DescribeDBEngineVersionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Engines struct {
			PostgerSQL []struct {
				Engine       *string `json:"Engine" name:"Engine"`
				EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
				MinorVersion *string `json:"MinorVersion" name:"MinorVersion"`
			} `json:"PostgerSQL"`
		} `json:"Engines" name:"Engines"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int      `json:"DiskUsed" name:"DiskUsed"`
			EipPort          *int      `json:"EipPort" name:"EipPort"`
			InnerAzCode      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string   `json:"ProjectName" name:"ProjectName"`
			Region           *string   `json:"Region" name:"Region"`
			ServiceStartTime *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId       *int      `json:"BillTypeId" name:"BillTypeId"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string   `json:"ProjectName" name:"ProjectName"`
			Region           *string   `json:"Region" name:"Region"`
			ServiceStartTime *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId       *int      `json:"BillTypeId" name:"BillTypeId"`
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

type ModifyDBInstanceSpecResponse struct {
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
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int    `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool   `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
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
			DBInstanceIdentifier  *string   `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
			DBInstanceName        *string   `json:"DBInstanceName" name:"DBInstanceName"`
			DBInstanceStatus      *string   `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			DBInstanceType        *string   `json:"DBInstanceType" name:"DBInstanceType"`
			DBParameterGroupId    *string   `json:"DBParameterGroupId" name:"DBParameterGroupId"`
			PreferredBackupTime   *string   `json:"PreferredBackupTime" name:"PreferredBackupTime"`
			GroupId               *string   `json:"GroupId" name:"GroupId"`
			SecurityGroupId       *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
			Vip                   *string   `json:"Vip" name:"Vip"`
			Port                  *int      `json:"Port" name:"Port"`
			Engine                *string   `json:"Engine" name:"Engine"`
			EngineVersion         *string   `json:"EngineVersion" name:"EngineVersion"`
			InstanceCreateTime    *string   `json:"InstanceCreateTime" name:"InstanceCreateTime"`
			DatastoreVersionId    *string   `json:"DatastoreVersionId" name:"DatastoreVersionId"`
			VpcId                 *string   `json:"VpcId" name:"VpcId"`
			SubnetId              *string   `json:"SubnetId" name:"SubnetId"`
			PubliclyAccessible    *bool     `json:"PubliclyAccessible" name:"PubliclyAccessible"`
			BillType              *string   `json:"BillType" name:"BillType"`
			OrderType             *string   `json:"OrderType" name:"OrderType"`
			MultiAvailabilityZone *bool     `json:"MultiAvailabilityZone" name:"MultiAvailabilityZone"`
			AvailabilityZoneList  []*string `json:"AvailabilityZoneList" name:"AvailabilityZoneList"`
			DiskUsed              *int      `json:"DiskUsed" name:"DiskUsed"`
			Audit                 *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			DBSource              struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName     *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType     *string `json:"DBInstanceType" name:"DBInstanceType"`
				DBBackupIdentifier *string `json:"DBBackupIdentifier" name:"DBBackupIdentifier"`
				DBBackupName       *string `json:"DBBackupName" name:"DBBackupName"`
				DBInstanceStatus   *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			} `json:"DBSource"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
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


type SwitchDBInstanceHARequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *SwitchDBInstanceHARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SwitchDBInstanceHAResponse struct {
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
			SecurityGroupId        *string `json:"SecurityGroupId" name:"SecurityGroupId"`
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int    `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool   `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string `json:"ProductId" name:"ProductId"`
			ProductWhat      *int    `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int    `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string `json:"ProjectName" name:"ProjectName"`
			Region           *string `json:"Region" name:"Region"`
			ServiceStartTime *string `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []struct {
				SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
				SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
			} `json:"SecurityGroups"`
			SupportIPV6 *bool `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId  *int  `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
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

type CreateDBInstanceReadReplicaResponse struct {
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
			ReadReplicaDBInstanceIdentifiers []struct {
				Vip *string `json:"Vip" name:"Vip"`
				ReadReplicaDBInstanceIdentifier *string `json:"ReadReplicaDBInstanceIdentifier" name:"ReadReplicaDBInstanceIdentifier"`
				Id  *string `json:"Id" name:"Id"`
			} `json:"ReadReplicaDBInstanceIdentifiers"`
			DBSource struct {
				DBInstanceIdentifier *string `json:"DBInstanceIdentifier" name:"DBInstanceIdentifier"`
				DBInstanceName   *string `json:"DBInstanceName" name:"DBInstanceName"`
				DBInstanceType   *string `json:"DBInstanceType" name:"DBInstanceType"`
				DBInstanceStatus *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
			} `json:"DBSource"`
			ProductId        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string   `json:"ProjectName" name:"ProjectName"`
			Region           *string   `json:"Region" name:"Region"`
			ServiceStartTime *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string   `json:"SubOrderId" name:"SubOrderId"`
			SecurityGroups   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			SupportIPV6      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BillTypeId       *int      `json:"BillTypeId" name:"BillTypeId"`
		} `json:"DBInstance" name:"DBInstance"`
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
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AccountName          *string `json:"AccountName,omitempty" name:"AccountName"`
	AccountPassword      *string `json:"AccountPassword,omitempty" name:"AccountPassword"`
	AccountDescription   *string `json:"AccountDescription,omitempty" name:"AccountDescription"`
}

func (r *ModifyInstanceAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DescribeInstanceDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		InstanceDatabases []struct {
			InstanceDatabaseName         *string `json:"InstanceDatabaseName" name:"InstanceDatabaseName"`
			InstanceDatabaseCollation    *string `json:"InstanceDatabaseCollation" name:"InstanceDatabaseCollation"`
			InstanceDatabaseOwner        *string `json:"InstanceDatabaseOwner" name:"InstanceDatabaseOwner"`
			InstanceDatabaseCharacterSet *string `json:"InstanceDatabaseCharacterSet" name:"InstanceDatabaseCharacterSet"`
			InstanceDatabaseStatus       *string `json:"InstanceDatabaseStatus" name:"InstanceDatabaseStatus"`
			InstanceDatabaseDescription  *string `json:"InstanceDatabaseDescription" name:"InstanceDatabaseDescription"`
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


type DescribeDBInstanceExtensionsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	KeyWord              *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *DescribeDBInstanceExtensionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBInstanceExtensionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Installed []struct {
			Comment          *string `json:"Comment" name:"Comment"`
			DefaultVersion   *string `json:"DefaultVersion" name:"DefaultVersion"`
			InstalledVersion *string `json:"InstalledVersion" name:"InstalledVersion"`
			Name             *string `json:"Name" name:"Name"`
		} `json:"Installed" name:"Installed"`
		Uninstalled []struct {
			Comment        *string `json:"Comment" name:"Comment"`
			DefaultVersion *string `json:"DefaultVersion" name:"DefaultVersion"`
			Name           *string `json:"Name" name:"Name"`
		} `json:"Uninstalled" name:"Uninstalled"`
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


type DescribeCollationsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeCollationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCollationsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Collations struct {
			LATIN5 struct {
				CtypeName     []*string `json:"CtypeName" name:"CtypeName"`
				CollationName []*string `json:"CollationName" name:"CollationName"`
			} `json:"LATIN5"`
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


type ModifyInstanceDatabaseOwnerRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Owner                *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *ModifyInstanceDatabaseOwnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceDatabaseOwnerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceDatabaseOwnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceDatabaseOwnerResponse) FromJsonString(s string) error {
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


type CreateInstanceDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier             *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	InstanceDatabaseName             *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	InstanceDatabaseCollation        *string `json:"InstanceDatabaseCollation,omitempty" name:"InstanceDatabaseCollation"`
	InstanceDatabaseCharacterSet     *string `json:"InstanceDatabaseCharacterSet,omitempty" name:"InstanceDatabaseCharacterSet"`
	InstanceDatabaseCharacterSetType *string `json:"InstanceDatabaseCharacterSetType,omitempty" name:"InstanceDatabaseCharacterSetType"`
	Description                      *string `json:"Description,omitempty" name:"Description"`
	InstanceDatabaseOwner            *string `json:"InstanceDatabaseOwner,omitempty" name:"InstanceDatabaseOwner"`
}

func (r *CreateInstanceDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DescribeInstanceAccountsRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
}

func (r *DescribeInstanceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceAccountsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Accounts []struct {
			AccountName        *string `json:"AccountName" name:"AccountName"`
			AccountStatus      *string `json:"AccountStatus" name:"AccountStatus"`
			AccountAccType     *int    `json:"AccountAccType" name:"AccountAccType"`
			AccountDescription *string `json:"AccountDescription" name:"AccountDescription"`
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

type CreateInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AccountName          *string `json:"AccountName,omitempty" name:"AccountName"`
	AccountPassword      *string `json:"AccountPassword,omitempty" name:"AccountPassword"`
	AccountDescription   *string `json:"AccountDescription,omitempty" name:"AccountDescription"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type DeleteInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AccountName          *string `json:"AccountName,omitempty" name:"AccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

type ModifyDBNetworkRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	SubnetId             *string `json:"SubnetId,omitempty" name:"SubnetId"`
	VpcId                *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *ModifyDBNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBNetworkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDBInstanceVersionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	Engine               *string `json:"Engine,omitempty" name:"Engine"`
	EngineVersion        *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *UpdateDBInstanceVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDBInstanceVersionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateDBInstanceVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDBInstanceVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceAvailabilityZoneRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	AvailabilityZone1    *string `json:"AvailabilityZone.1,omitempty" name:"AvailabilityZone.1"`
	AvailabilityZone2    *string `json:"AvailabilityZone.2,omitempty" name:"AvailabilityZone.2"`
}

func (r *ModifyDBInstanceAvailabilityZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBInstanceAvailabilityZoneResponse struct {
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
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AvailabilityZoneList"`
			DiskUsed         *int      `json:"DiskUsed" name:"DiskUsed"`
			InnerAzCode      *string   `json:"InnerAzCode" name:"InnerAzCode"`
			Audit            *bool     `json:"Audit" name:"Audit"`
			ReadReplicaDBInstanceIdentifiers []*string `json:"ReadReplicaDBInstanceIdentifiers" name:"ReadReplicaDBInstanceIdentifiers"`
			ProductId        *string   `json:"ProductId" name:"ProductId"`
			ProductWhat      *int      `json:"ProductWhat" name:"ProductWhat"`
			ProjectId        *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName      *string   `json:"ProjectName" name:"ProjectName"`
			Region           *string   `json:"Region" name:"Region"`
			ServiceStartTime *string   `json:"ServiceStartTime" name:"ServiceStartTime"`
			SubOrderId       *string   `json:"SubOrderId" name:"SubOrderId"`
			MiniVersion      *string   `json:"MiniVersion" name:"MiniVersion"`
			SecurityGroups   []*string `json:"SecurityGroups" name:"SecurityGroups"`
			NetworkType      *int      `json:"NetworkType" name:"NetworkType"`
			MaintenanceTime  *string   `json:"MaintenanceTime" name:"MaintenanceTime"`
			SupportIPV6      *bool     `json:"SupportIPV6" name:"SupportIPV6"`
			BindInstances    []*string `json:"BindInstances" name:"BindInstances"`
			ProxyNodeInfo    []*string `json:"ProxyNodeInfo" name:"ProxyNodeInfo"`
			ProxyInfo        struct {
			} `json:"ProxyInfo"`
			AutoSwitch *int `json:"AutoSwitch" name:"AutoSwitch"`
			BillTypeId *int `json:"BillTypeId" name:"BillTypeId"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDBInstanceAvailabilityZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceAvailabilityZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDBInstanceOrderRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	UpdateUse            *string `json:"UpdateUse,omitempty" name:"UpdateUse"`
	Duration             *int    `json:"Duration,omitempty" name:"Duration"`
	BillType             *string `json:"BillType,omitempty" name:"BillType"`
}

func (r *UpdateDBInstanceOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateDBInstanceOrderResponse struct {
	*ksyunhttp.BaseResponse
	Status      *string   `json:"status" name:"status"`
	OrderId     *string   `json:"orderId" name:"orderId"`
	TotalMoney  *int      `json:"totalMoney" name:"totalMoney"`
	RealMoney   *int      `json:"realMoney" name:"realMoney"`
	LastMoney   *int      `json:"lastMoney" name:"lastMoney"`
	SubOrderIds []*string `json:"SubOrderIds" name:"SubOrderIds"`
}

func (r *UpdateDBInstanceOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDBInstanceOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceProtectionRequest struct {
	*ksyunhttp.BaseRequest
	DBInstanceIdentifier *string `json:"DBInstanceIdentifier,omitempty" name:"DBInstanceIdentifier"`
	ProtectionSwitch     *string `json:"ProtectionSwitch,omitempty" name:"ProtectionSwitch"`
	ProtectionReason     *string `json:"ProtectionReason,omitempty" name:"ProtectionReason"`
}

func (r *UpdateResourceProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateResourceProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateResourceProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

