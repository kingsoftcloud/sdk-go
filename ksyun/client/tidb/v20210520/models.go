package v20210520
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type CreateInstanceModuleConfigs struct {
	ModuleType  *string `json:"ModuleType,omitempty" name:"ModuleType"`
	PackageCode *string `json:"PackageCode,omitempty" name:"PackageCode"`
	Replicas    *int    `json:"Replicas,omitempty" name:"Replicas"`
	Cpu         *int    `json:"Cpu,omitempty" name:"Cpu"`
	Mem         *int    `json:"Mem,omitempty" name:"Mem"`
	Tidisk      *int    `json:"Tidisk,omitempty" name:"Tidisk"`
}
type CreateInstanceBackupConfig struct {
	MaxBackups          *string `json:"MaxBackups,omitempty" name:"MaxBackups"`
	MaxReservedHours    *string `json:"MaxReservedHours,omitempty" name:"MaxReservedHours"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
}
type CreateSecurityRuleRules struct {
	Cidr []*string `json:"Cidr,omitempty" name:"Cidr"`
	Desc []*string `json:"Desc,omitempty" name:"Desc"`
}
type CreateTaskTargetMySQL struct {
	User        *string `json:"User,omitempty" name:"User"`
	Passwd      *string `json:"Passwd,omitempty" name:"Passwd"`
	WorkerCount *int    `json:"WorkerCount,omitempty" name:"WorkerCount"`
	MaxTxnRow   *int    `json:"MaxTxnRow,omitempty" name:"MaxTxnRow"`
}
type CreateTaskTargetKafka struct {
	KafkaVersion      *string   `json:"KafkaVersion,omitempty" name:"KafkaVersion"`
	TopicName         *string   `json:"TopicName,omitempty" name:"TopicName"`
	Protocol          *string   `json:"Protocol,omitempty" name:"Protocol"`
	PartitionNum      *int      `json:"PartitionNum,omitempty" name:"PartitionNum"`
	MaxMessageBytes   *int      `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
	ReplicationFacter *int      `json:"ReplicationFacter,omitempty" name:"ReplicationFacter"`
	MounterWorkerNum  *int      `json:"MounterWorkerNum,omitempty" name:"MounterWorkerNum"`
	DomainList        []*string `json:"DomainList,omitempty" name:"DomainList"`
	EndpointList      []*string `json:"EndpointList,omitempty" name:"EndpointList"`
}


type CreateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceName            *string                        `json:"InstanceName,omitempty" name:"InstanceName"`
	EnableModules           *string                        `json:"EnableModules,omitempty" name:"EnableModules"`
	ModuleConfigs           []*CreateInstanceModuleConfigs `json:"ModuleConfigs,omitempty" name:"ModuleConfigs"`
	AdminUser               *string                        `json:"AdminUser,omitempty" name:"AdminUser"`
	AdminPassword           *string                        `json:"AdminPassword,omitempty" name:"AdminPassword"`
	VpcId                   *string                        `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId                *string                        `json:"SubnetId,omitempty" name:"SubnetId"`
	BillType                *int                           `json:"BillType,omitempty" name:"BillType"`
	Duration                *string                        `json:"Duration,omitempty" name:"Duration"`
	ProductType             *int                           `json:"ProductType,omitempty" name:"ProductType"`
	ProjectId               *string                        `json:"ProjectId,omitempty" name:"ProjectId"`
	EnableAutoBackup        *bool                          `json:"EnableAutoBackup,omitempty" name:"EnableAutoBackup"`
	Engine                  *string                        `json:"Engine,omitempty" name:"Engine"`
	EngineVersion           *string                        `json:"EngineVersion,omitempty" name:"EngineVersion"`
	ClientPort              *int                           `json:"ClientPort,omitempty" name:"ClientPort"`
	Az                      *string                        `json:"Az,omitempty" name:"Az"`
	SecurityGroupId         *string                        `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	BackupConfig            *CreateInstanceBackupConfig    `json:"BackupConfig,omitempty" name:"BackupConfig"`
	BackupId                *string                        `json:"backupId,omitempty" name:"backupId"`
	BackupRestoreInstanceId *string                        `json:"backupRestoreInstanceId,omitempty" name:"backupRestoreInstanceId"`
	BackupRestoreTime       *string                        `json:"backupRestoreTime,omitempty" name:"backupRestoreTime"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		OrderId   *string `json:"OrderId" name:"OrderId"`
		AccountId *string `json:"AccountId" name:"AccountId"`
		Region    *string `json:"Region" name:"Region"`
		Az        *string `json:"Az" name:"Az"`
	} `json:"Data"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProjectIds     *string `json:"ProjectIds,omitempty" name:"ProjectIds"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	FuzzySearch    *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset         *int    `json:"Offset,omitempty" name:"Offset"`
	Limit          *int    `json:"Limit,omitempty" name:"Limit"`
	OrderBy        *string `json:"OrderBy,omitempty" name:"OrderBy"`
	ProductType    *int    `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *ListInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
	Total     *int    `json:"Total" name:"Total"`
	Data      []struct {
		InstanceId    *string `json:"InstanceId" name:"InstanceId"`
		InstanceName  *string `json:"InstanceName" name:"InstanceName"`
		AdminUser     *string `json:"AdminUser" name:"AdminUser"`
		StatusName    *string `json:"StatusName" name:"StatusName"`
		Status        *string `json:"Status" name:"Status"`
		VpcId         *string `json:"VpcId" name:"VpcId"`
		SubnetId      *string `json:"SubnetId" name:"SubnetId"`
		Engine        *string `json:"Engine" name:"Engine"`
		EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
		ProjectId     *string `json:"ProjectId" name:"ProjectId"`
		ProjectName   *string `json:"ProjectName" name:"ProjectName"`
		BillType      *int    `json:"BillType" name:"BillType"`
		ProductId     *string `json:"ProductId" name:"ProductId"`
		ProductType   *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		ProductWhat   *int    `json:"ProductWhat" name:"ProductWhat"`
		CreateDate    *string `json:"CreateDate" name:"CreateDate"`
		UpdateDate    *string `json:"UpdateDate" name:"UpdateDate"`
		Region        *string `json:"Region" name:"Region"`
		RegionName    *string `json:"RegionName" name:"RegionName"`
		Az            *string `json:"Az" name:"Az"`
		AzName        *string `json:"AzName" name:"AzName"`
		UserId        *string `json:"UserId" name:"UserId"`
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		ClientIp      *string `json:"ClientIp" name:"ClientIp"`
		ClientPort    *int    `json:"ClientPort" name:"ClientPort"`
		EnableModules *string `json:"EnableModules" name:"EnableModules"`
		ModuleConfigs []struct {
			ModuleType  *string `json:"moduleType" name:"moduleType"`
			PackageCode *string `json:"packageCode" name:"packageCode"`
			Replicas    *int    `json:"replicas" name:"replicas"`
			Cpu         *int    `json:"cpu" name:"cpu"`
			Mem         *int    `json:"mem" name:"mem"`
			Tidisk      *int    `json:"tidisk" name:"tidisk"`
			CreateTime  *string `json:"createTime" name:"createTime"`
		} `json:"ModuleConfigs" name:"ModuleConfigs"`
		EndTime *string `json:"EndTime" name:"EndTime"`
	} `json:"Data"`
}

func (r *ListInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceResponse) FromJsonString(s string) error {
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

type DescribeInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId    *string `json:"InstanceId" name:"InstanceId"`
		InstanceName  *string `json:"InstanceName" name:"InstanceName"`
		EnableModules *string `json:"EnableModules" name:"EnableModules"`
		ModuleConfigs []struct {
			ModuleType  *string `json:"moduleType" name:"moduleType"`
			PackageCode *string `json:"packageCode" name:"packageCode"`
			Replicas    *int    `json:"replicas" name:"replicas"`
			Cpu         *int    `json:"cpu" name:"cpu"`
			Mem         *int    `json:"mem" name:"mem"`
			Tidisk      *int    `json:"tidisk" name:"tidisk"`
			CreateTime  *string `json:"createTime" name:"createTime"`
		} `json:"ModuleConfigs" name:"ModuleConfigs"`
		AdminUser     *string `json:"AdminUser" name:"AdminUser"`
		StatusName    *string `json:"StatusName" name:"StatusName"`
		Status        *string `json:"Status" name:"Status"`
		VpcId         *string `json:"VpcId" name:"VpcId"`
		SubnetId      *string `json:"SubnetId" name:"SubnetId"`
		Engine        *string `json:"Engine" name:"Engine"`
		EngineVersion *string `json:"EngineVersion" name:"EngineVersion"`
		ProjectId     *string `json:"ProjectId" name:"ProjectId"`
		ProjectName   *string `json:"ProjectName" name:"ProjectName"`
		BillType      *int    `json:"BillType" name:"BillType"`
		ProductId     *string `json:"ProductId" name:"ProductId"`
		ProductType   *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		ProductWhat   *int    `json:"ProductWhat" name:"ProductWhat"`
		CreateDate    *string `json:"CreateDate" name:"CreateDate"`
		Region        *string `json:"Region" name:"Region"`
		RegionName    *string `json:"RegionName" name:"RegionName"`
		Az            *string `json:"Az" name:"Az"`
		AzName        *string `json:"AzName" name:"AzName"`
		UserId        *string `json:"UserId" name:"UserId"`
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		SecurityGroupDesc *string `json:"SecurityGroupDesc" name:"SecurityGroupDesc"`
		BackupConfig  struct {
			MaxBackups          *int    `json:"maxBackups" name:"maxBackups"`
			MaxReservedHours    *int    `json:"maxReservedHours" name:"maxReservedHours"`
			PreferredBackupTime *string `json:"preferredBackupTime" name:"preferredBackupTime"`
			Increment           *string `json:"increment" name:"increment"`
		} `json:"BackupConfig" name:"BackupConfig"`
		ClientIp             *string `json:"ClientIp" name:"ClientIp"`
		ClientPort           *int    `json:"ClientPort" name:"ClientPort"`
		EndTime              *string `json:"EndTime" name:"EndTime"`
		DashboardClientVip   *string `json:"DashboardClientVip" name:"DashboardClientVip"`
		DashboardClientVport *int    `json:"DashboardClientVport" name:"DashboardClientVport"`
		TimonitorURL         *string `json:"TimonitorURL" name:"TimonitorURL"`
	} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
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

type DeleteInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
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

type RenameInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListRegionRequest struct {
	*ksyunhttp.BaseRequest
	ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *ListRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListRegionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		RegionId   *string `json:"RegionId" name:"RegionId"`
		InnerCode  *string `json:"InnerCode" name:"InnerCode"`
		RegionCode *string `json:"RegionCode" name:"RegionCode"`
		RegionName *string `json:"RegionName" name:"RegionName"`
		RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
		Active     *bool   `json:"Active" name:"Active"`
		RegionType *int    `json:"RegionType" name:"RegionType"`
		Overseas   *bool   `json:"Overseas" name:"Overseas"`
		AzList     []struct {
			AzCode *string `json:"AzCode" name:"AzCode"`
			AzName *string `json:"AzName" name:"AzName"`
		} `json:"AzList" name:"AzList"`
		AreaCode *string `json:"AreaCode" name:"AreaCode"`
		AreaName *string `json:"AreaName" name:"AreaName"`
		AreaEnName *string `json:"AreaEnName" name:"AreaEnName"`
	} `json:"Data"`
}

func (r *ListRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescRegionRequest struct {
	*ksyunhttp.BaseRequest
	RegionCode  *string `json:"RegionCode,omitempty" name:"RegionCode"`
	ProductType *int    `json:"ProductType,omitempty" name:"ProductType"`
}

func (r *DescRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescRegionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		RegionId   *string `json:"RegionId" name:"RegionId"`
		InnerCode  *string `json:"InnerCode" name:"InnerCode"`
		RegionCode *string `json:"RegionCode" name:"RegionCode"`
		RegionName *string `json:"RegionName" name:"RegionName"`
		RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
		Active     *bool   `json:"Active" name:"Active"`
		RegionType *int    `json:"RegionType" name:"RegionType"`
		Overseas   *bool   `json:"Overseas" name:"Overseas"`
		AzList     []struct {
			AzCode *string `json:"AzCode" name:"AzCode"`
			AzName *string `json:"AzName" name:"AzName"`
		} `json:"AzList" name:"AzList"`
		AreaCode *string `json:"AreaCode" name:"AreaCode"`
		AreaName *string `json:"AreaName" name:"AreaName"`
		AreaEnName *string `json:"AreaEnName" name:"AreaEnName"`
	} `json:"Data"`
}

func (r *DescRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	ProductType       *int    `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	IpVersion         *string `json:"IpVersion,omitempty" name:"IpVersion"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		GroupType     *string `json:"GroupType" name:"GroupType"`
		IpVersion     *string `json:"IpVersion" name:"IpVersion"`
		Description   *string `json:"Description" name:"Description"`
		InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		Region        *string `json:"Region" name:"Region"`
		Rules         *string `json:"Rules" name:"Rules"`
		Instances     *string `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset      *int    `json:"Offset,omitempty" name:"Offset"`
	Limit       *int    `json:"Limit,omitempty" name:"Limit"`
	OrderBy     *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
	Total     *int    `json:"Total" name:"Total"`
	Data      []struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		GroupType     *string `json:"GroupType" name:"GroupType"`
		IpVersion     *string `json:"IpVersion" name:"IpVersion"`
		Description   *string `json:"Description" name:"Description"`
		InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		Region        *string `json:"Region" name:"Region"`
	} `json:"Data"`
}

func (r *ListSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId     *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	RuleFuzzySearch     *string `json:"RuleFuzzySearch,omitempty" name:"RuleFuzzySearch"`
	InstanceFuzzySearch *string `json:"InstanceFuzzySearch,omitempty" name:"InstanceFuzzySearch"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		GroupType     *string `json:"GroupType" name:"GroupType"`
		IpVersion     *string `json:"IpVersion" name:"IpVersion"`
		Description   *string `json:"Description" name:"Description"`
		InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		Rules         []struct {
			IpVersion   *string `json:"ipVersion" name:"ipVersion"`
			Protocol    *string `json:"protocol" name:"protocol"`
			Action      *string `json:"action" name:"action"`
			Priority    *int    `json:"priority" name:"priority"`
			RuleId      *string `json:"RuleId" name:"RuleId"`
			Cidr        *string `json:"Cidr" name:"Cidr"`
			Description *string `json:"Description" name:"Description"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
		} `json:"Rules" name:"Rules"`
		Instances []struct {
			InstanceId      *string `json:"InstanceId" name:"InstanceId"`
			InstanceName    *string `json:"InstanceName" name:"InstanceName"`
			ProductType     *int    `json:"ProductType" name:"ProductType"`
			ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime      *string `json:"UpdateTime" name:"UpdateTime"`
			ClientIp        *string `json:"ClientIp" name:"ClientIp"`
			ClientPort      *string `json:"ClientPort" name:"ClientPort"`
		} `json:"Instances" name:"Instances"`
	} `json:"Data"`
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

type DeleteSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		GroupType     *string `json:"GroupType" name:"GroupType"`
		IpVersion     *string `json:"IpVersion" name:"IpVersion"`
		Description   *string `json:"Description" name:"Description"`
		InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		Region        *string `json:"Region" name:"Region"`
		Rules         *string `json:"Rules" name:"Rules"`
		Instances     *string `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *UpdateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloneSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SrcSecurityGroupId *string `json:"SrcSecurityGroupId,omitempty" name:"SrcSecurityGroupId"`
	SecurityGroupName  *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	Description        *string `json:"Description,omitempty" name:"Description"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloneSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		GroupType     *string `json:"GroupType" name:"GroupType"`
		IpVersion     *string `json:"IpVersion" name:"IpVersion"`
		Description   *string `json:"Description" name:"Description"`
		InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime    *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime    *string `json:"UpdateTime" name:"UpdateTime"`
		Rules         []struct {
			IpVersion   *string `json:"ipVersion" name:"ipVersion"`
			Protocol    *string `json:"protocol" name:"protocol"`
			Action      *string `json:"action" name:"action"`
			Priority    *int    `json:"priority" name:"priority"`
			RuleId      *string `json:"RuleId" name:"RuleId"`
			Cidr        *string `json:"Cidr" name:"Cidr"`
			Description *string `json:"Description" name:"Description"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
		} `json:"Rules" name:"Rules"`
		Instances *string `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type BindSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstanceIds     *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *BindSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type BindSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *BindSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UnbindSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstanceIds     *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *UnbindSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UnbindSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *UnbindSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RebindSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RebindSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RebindSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *RebindSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebindSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string                  `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Rules           *CreateSecurityRuleRules `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateSecurityRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSecurityRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *CreateSecurityRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	RuleId      *string `json:"RuleId,omitempty" name:"RuleId"`
	Description *string `json:"Description,omitempty" name:"Description"`
	Cidr        *string `json:"Cidr,omitempty" name:"Cidr"`
}

func (r *UpdateSecurityRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateSecurityRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *UpdateSecurityRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	RuleIds         *string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DeleteSecurityRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSecurityRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *DeleteSecurityRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListSecuredInstanceRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	ProjectIds      *string `json:"ProjectIds,omitempty" name:"ProjectIds"`
	FuzzySearch     *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset          *int    `json:"Offset,omitempty" name:"Offset"`
	Limit           *int    `json:"Limit,omitempty" name:"Limit"`
	OrderBy         *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListSecuredInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListSecuredInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
	Total     *int    `json:"Total" name:"Total"`
	Data      []struct {
		InstanceId  *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		ProductType *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
		ClientIp    *string `json:"ClientIp" name:"ClientIp"`
		ClientPort  *int    `json:"ClientPort" name:"ClientPort"`
	} `json:"Data"`
}

func (r *ListSecuredInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecuredInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListUnsecuredInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProjectIds  *string `json:"ProjectIds,omitempty" name:"ProjectIds"`
	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset      *int    `json:"Offset,omitempty" name:"Offset"`
	Limit       *int    `json:"Limit,omitempty" name:"Limit"`
	OrderBy     *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListUnsecuredInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListUnsecuredInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
	Total     *int    `json:"Total" name:"Total"`
	Data      []struct {
		InstanceId  *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		ProductType *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime  *string `json:"UpdateTime" name:"UpdateTime"`
		ClientIp    *string `json:"ClientIp" name:"ClientIp"`
		ClientPort  *int    `json:"ClientPort" name:"ClientPort"`
	} `json:"Data"`
}

func (r *ListUnsecuredInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUnsecuredInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListBackupRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Keyword    *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListBackupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		BackupId     *string `json:"BackupId" name:"BackupId"`
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		BackupName   *string `json:"BackupName" name:"BackupName"`
		BackupStatus *string `json:"BackupStatus" name:"BackupStatus"`
		BackupStatusName *string `json:"BackupStatusName" name:"BackupStatusName"`
		BackupSizeReadable *string `json:"BackupSizeReadable" name:"BackupSizeReadable"`
		RunType      *int    `json:"RunType" name:"RunType"`
		RunTypeName  *string `json:"RunTypeName" name:"RunTypeName"`
		IncreaseName *string `json:"IncreaseName" name:"IncreaseName"`
		StartTime    *string `json:"StartTime" name:"StartTime"`
		ComplateTime *string `json:"ComplateTime" name:"ComplateTime"`
		Cost         *int    `json:"Cost" name:"Cost"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
	} `json:"Data"`
}

func (r *ListBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateBackupRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *CreateBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateBackupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *CreateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateBackupRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId          *string `json:"InstanceId,omitempty" name:"InstanceId"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	EnableAutobackup    *bool   `json:"EnableAutobackup,omitempty" name:"EnableAutobackup"`
	EnableIncremental   *bool   `json:"EnableIncremental,omitempty" name:"EnableIncremental"`
}

func (r *UpdateBackupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateBackupRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *UpdateBackupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBackupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteBackupRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	BackupIds  *string `json:"BackupIds,omitempty" name:"BackupIds"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBackupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateRestoreRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	BackupId   *string `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *CreateRestoreRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateRestoreResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *CreateRestoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRestoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type OpenTiMonitorRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenTiMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OpenTiMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *OpenTiMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenTiMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateTaskRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string                `json:"InstanceId,omitempty" name:"InstanceId"`
	TaskName    *string                `json:"TaskName,omitempty" name:"TaskName"`
	TargetId    *string                `json:"TargetId,omitempty" name:"TargetId"`
	VpcId       *string                `json:"VpcId,omitempty" name:"VpcId"`
	VnetId      *string                `json:"VnetId,omitempty" name:"VnetId"`
	TargetType  *string                `json:"TargetType,omitempty" name:"TargetType"`
	TargetMySQL *CreateTaskTargetMySQL `json:"TargetMySQL,omitempty" name:"TargetMySQL"`
	TargetKafka *CreateTaskTargetKafka `json:"TargetKafka,omitempty" name:"TargetKafka"`
}

func (r *CreateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTaskResponse struct {
	*ksyunhttp.BaseResponse
	TaskName   *string `json:"TaskName" name:"TaskName"`
	OperStatus *string `json:"OperStatus" name:"OperStatus"`
	Msg        struct {
	} `json:"Msg"`
}

func (r *CreateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type OperationTasksRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TaskList   *string `json:"TaskList,omitempty" name:"TaskList"`
	Operation  *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *OperationTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OperationTasksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		TaskName *string `json:"TaskName" name:"TaskName"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *OperationTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CheckTaskNameRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CheckTaskNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CheckTaskNameResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Code      *string   `json:"Code" name:"Code"`
	Message   *string   `json:"Message" name:"Message"`
	Data      []*string `json:"Data" name:"Data"`
}

func (r *CheckTaskNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTaskNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeTaskRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TaskName   *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		TiCDCTaskBaseInfo struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			TaskName     *string `json:"TaskName" name:"TaskName"`
			TaskStatus   *string `json:"TaskStatus" name:"TaskStatus"`
			CreateTime   *string `json:"CreateTime" name:"CreateTime"`
			CheckPointTs *string `json:"CheckPointTs" name:"CheckPointTs"`
			ResolvedTs   *string `json:"ResolvedTs" name:"ResolvedTs"`
			DelayTime    *string `json:"DelayTime" name:"DelayTime"`
			PauseTime    *string `json:"PauseTime" name:"PauseTime"`
			FinishTime   *string `json:"FinishTime" name:"FinishTime"`
			TaskFailInfo *string `json:"TaskFailInfo" name:"TaskFailInfo"`
		} `json:"TiCDCTaskBaseInfo" name:"TiCDCTaskBaseInfo"`
		TiCDCTaskTargetInfo struct {
			TargetId *string `json:"TargetId" name:"TargetId"`
			SinkUri  *string `json:"SinkUri" name:"SinkUri"`
		} `json:"TiCDCTaskTargetInfo" name:"TiCDCTaskTargetInfo"`
	} `json:"Data"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListTasksRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ListTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListTasksResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		TaskName *string `json:"TaskName" name:"TaskName"`
		TaskStatus *string `json:"TaskStatus" name:"TaskStatus"`
		TargetType *string `json:"TargetType" name:"TargetType"`
		TargetId *string `json:"TargetId" name:"TargetId"`
		TargetAddress *string `json:"TargetAddress" name:"TargetAddress"`
		TargetPort *string `json:"TargetPort" name:"TargetPort"`
		VpcId    *string `json:"VpcId" name:"VpcId"`
	} `json:"Data"`
}

func (r *ListTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDatabasesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Code      *string   `json:"Code" name:"Code"`
	Message   *string   `json:"Message" name:"Message"`
	Data      []*string `json:"Data" name:"Data"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
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

type DescribeAccountsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		AccountName *string `json:"AccountName" name:"AccountName"`
		AccountType *string `json:"AccountType" name:"AccountType"`
		Describe   *string `json:"Describe" name:"Describe"`
		Privileges struct {
			ALL *string `json:"ALL" name:"ALL"`
		} `json:"Privileges" name:"Privileges"`
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
	Describe        *string `json:"Describe,omitempty" name:"Describe"`
	Privileges      *string `json:"Privileges,omitempty" name:"Privileges"`
}

func (r *CreateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
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

type DeleteAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyAccountInfoRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName     *string `json:"AccountName,omitempty" name:"AccountName"`
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`
	AccountType     *string `json:"AccountType,omitempty" name:"AccountType"`
}

func (r *ModifyAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyAccountInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *ModifyAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyAccountPrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	AccountName   *string `json:"AccountName,omitempty" name:"AccountName"`
	OldPrivileges *string `json:"OldPrivileges,omitempty" name:"OldPrivileges"`
	NewPrivileges *string `json:"NewPrivileges,omitempty" name:"NewPrivileges"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyAccountPrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ConfigurationInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Operation  *string `json:"Operation,omitempty" name:"Operation"`
	EipPort    *int    `json:"EipPort,omitempty" name:"EipPort"`
}

func (r *ConfigurationInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ConfigurationInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *ConfigurationInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigurationInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateInstanceTrialOrderRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	BillType    *int    `json:"BillType,omitempty" name:"BillType"`
	Duration    *int    `json:"Duration,omitempty" name:"Duration"`
}

func (r *UpdateInstanceTrialOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateInstanceTrialOrderResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		OrderId    *string `json:"OrderId" name:"OrderId"`
		AccountId  *string `json:"AccountId" name:"AccountId"`
		Region     *string `json:"Region" name:"Region"`
		Az         *string `json:"Az" name:"Az"`
	} `json:"Data"`
}

func (r *UpdateInstanceTrialOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceTrialOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

