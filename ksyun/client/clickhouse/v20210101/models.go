package v20210101

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type ModifyInstanceAccountPrivilegesInstanceAccountPrivileges struct {
	InstanceDatabaseName *string `json:"InstanceDatabaseName,omitempty" name:"InstanceDatabaseName"`
	Privilege            *string `json:"Privilege,omitempty" name:"Privilege"`
}

type ListInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProductType *string   `json:"ProductType,omitempty" name:"ProductType"`
	TagId       *string   `json:"TagId,omitempty" name:"TagId"`
	ProjectIds  []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`
	FuzzySearch *string   `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	OrderBy     []*string `json:"OrderBy,omitempty" name:"OrderBy"`
	Offset      *int      `json:"Offset,omitempty" name:"Offset"`
	Limit       *int      `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListInstanceResponse struct {
	*ksyunhttp.BaseResponse
	One *int `json:"1" name:"1"`
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
		InstanceId        *string `json:"InstanceId" name:"InstanceId"`
		InstanceName      *string `json:"InstanceName" name:"InstanceName"`
		InstacneConfig    *string `json:"InstacneConfig" name:"InstacneConfig"`
		AdminUser         *string `json:"AdminUser" name:"AdminUser"`
		StatusName        *string `json:"StatusName" name:"StatusName"`
		Status            *string `json:"Status" name:"Status"`
		NetworkType       *string `json:"NetworkType" name:"NetworkType"`
		VpcId             *string `json:"VpcId" name:"VpcId"`
		SubnetId          *string `json:"SubnetId" name:"SubnetId"`
		Vip               *string `json:"Vip" name:"Vip"`
		Engine            *string `json:"Engine" name:"Engine"`
		EngineVersion     *string `json:"EngineVersion" name:"EngineVersion"`
		ProjectId         *string `json:"ProjectId" name:"ProjectId"`
		ProjectName       *string `json:"ProjectName" name:"ProjectName"`
		BillType          *int    `json:"BillType" name:"BillType"`
		StorageSize       *int    `json:"StorageSize" name:"StorageSize"`
		UsedStorageSize   *int    `json:"UsedStorageSize" name:"UsedStorageSize"`
		StorageType       *string `json:"StorageType" name:"StorageType"`
		Mem               *int    `json:"Mem" name:"Mem"`
		Cpu               *int    `json:"Cpu" name:"Cpu"`
		TcpPort           *int    `json:"TcpPort" name:"TcpPort"`
		HttpPort          *int    `json:"HttpPort" name:"HttpPort"`
		NodeNum           *int    `json:"NodeNum" name:"NodeNum"`
		ProductId         *string `json:"ProductId" name:"ProductId"`
		ProductType       *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName   *string `json:"ProductTypeName" name:"ProductTypeName"`
		ProductWhat       *int    `json:"ProductWhat" name:"ProductWhat"`
		CreateDate        *string `json:"CreateDate" name:"CreateDate"`
		UpdateDate        *string `json:"UpdateDate" name:"UpdateDate"`
		Region            *string `json:"Region" name:"Region"`
		Az                *string `json:"Az" name:"Az"`
		UserId            *string `json:"UserId" name:"UserId"`
		SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		SecurityGroupDesc *string `json:"SecurityGroupDesc" name:"SecurityGroupDesc"`
		ServiceEndTime    *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		ShardList         []struct {
			Id   *string `json:"id" name:"id"`
			Name *string `json:"name" name:"name"`
		} `json:"ShardList" name:"ShardList"`
		Replicas             *int      `json:"Replicas" name:"Replicas"`
		DirectConnectionVips []*string `json:"DirectConnectionVips" name:"DirectConnectionVips"`
		MultiAz              *int      `json:"MultiAz" name:"MultiAz"`
		Area                 struct {
			Master  *string `json:"Master" name:"Master"`
			Standby *string `json:"Standby" name:"Standby"`
		} `json:"Area" name:"Area"`
	} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProductType         *string `json:"ProductType,omitempty" name:"ProductType"`
	InstanceName        *string `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceType        *string `json:"InstanceType,omitempty" name:"InstanceType"`
	AdminUser           *string `json:"AdminUser,omitempty" name:"AdminUser"`
	AdminPassword       *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	VpcId               *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId            *string `json:"SubnetId,omitempty" name:"SubnetId"`
	Engine              *string `json:"Engine,omitempty" name:"Engine"`
	EngineVersion       *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	ProjectId           *string `json:"ProjectId,omitempty" name:"ProjectId"`
	BillType            *int    `json:"BillType,omitempty" name:"BillType"`
	Duration            *int    `json:"Duration,omitempty" name:"Duration"`
	EbsSize             *int    `json:"EbsSize,omitempty" name:"EbsSize"`
	EbsType             *string `json:"EbsType,omitempty" name:"EbsType"`
	Mem                 *int    `json:"Mem,omitempty" name:"Mem"`
	Cpu                 *int    `json:"Cpu,omitempty" name:"Cpu"`
	TcpPort             *int    `json:"TcpPort,omitempty" name:"TcpPort"`
	HttpPort            *int    `json:"HttpPort,omitempty" name:"HttpPort"`
	Az                  *string `json:"Az,omitempty" name:"Az"`
	NodeNum             *int    `json:"NodeNum,omitempty" name:"NodeNum"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" name:"PreferredBackupTime"`
	SecurityGroupId     *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		OrderId      *string `json:"OrderId" name:"OrderId"`
		AccountId    *string `json:"AccountId" name:"AccountId"`
		Region       *string `json:"Region" name:"Region"`
		Az           *string `json:"Az" name:"Az"`
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
	InstanceIds    []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	DeleteDirectly *bool     `json:"DeleteDirectly,omitempty" name:"DeleteDirectly"`
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
		Msg        *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds *string `json:"instanceIds,omitempty" name:"instanceIds"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestartInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg        *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
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

type ListSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	ProductType *int `json:"ProductType,omitempty" name:"ProductType"`
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
	Data      []struct {
		SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		IpVersion         *string `json:"IpVersion" name:"IpVersion"`
		Description       *string `json:"Description" name:"Description"`
		InstanceCount     *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime        *string `json:"UpdateTime" name:"UpdateTime"`
		Rules             []struct {
			RuleId      *string `json:"RuleId" name:"RuleId"`
			Cidr        *string `json:"Cidr" name:"Cidr"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			Description *string `json:"Description" name:"Description"`
		} `json:"Rules" name:"Rules"`
		Instances []struct {
			InstanceId      *string `json:"InstanceId" name:"InstanceId"`
			InstanceName    *string `json:"InstanceName" name:"InstanceName"`
			ProductType     *int    `json:"ProductType" name:"ProductType"`
			ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
			Vip             *string `json:"Vip" name:"Vip"`
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			Status          *string `json:"Status" name:"Status"`
			UpdatetTime     *string `json:"UpdatetTime" name:"UpdatetTime"`
		} `json:"Instances" name:"Instances"`
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
	ProductType     *int    `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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
		SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		IpVersion         *string `json:"IpVersion" name:"IpVersion"`
		Description       *string `json:"Description" name:"Description"`
		InstanceCount     *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime        *string `json:"UpdateTime" name:"UpdateTime"`
		Rules             []struct {
			RuleId      *string `json:"RuleId" name:"RuleId"`
			Cidr        *string `json:"Cidr" name:"Cidr"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			Description *string `json:"Description" name:"Description"`
		} `json:"Rules" name:"Rules"`
		Instances []struct {
			InstanceId      *string `json:"InstanceId" name:"InstanceId"`
			InstanceName    *string `json:"InstanceName" name:"InstanceName"`
			ProductType     *int    `json:"ProductType" name:"ProductType"`
			ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
			Vip             *string `json:"Vip" name:"Vip"`
			CreateTime      *string `json:"CreateTime" name:"CreateTime"`
			Status          *string `json:"Status" name:"Status"`
			UpdatetTime     *string `json:"UpdatetTime" name:"UpdatetTime"`
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

type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	ProductType       *int    `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
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
		SecurityGroupId   *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
		IpVersion         *string   `json:"IpVersion" name:"IpVersion"`
		Description       *string   `json:"Description" name:"Description"`
		InstanceCount     *int      `json:"InstanceCount" name:"InstanceCount"`
		CreateTime        *string   `json:"CreateTime" name:"CreateTime"`
		UpdateTime        *string   `json:"UpdateTime" name:"UpdateTime"`
		Rules             []*string `json:"Rules" name:"Rules"`
		Instances         []*string `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	ProductType      *int    `json:"ProductType,omitempty" name:"ProductType"`
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
		OperStatus      *string `json:"OperStatus" name:"OperStatus"`
		Msg             *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	ProductType       *int    `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *RenameSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenameSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		SecurityGroupId   *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string   `json:"SecurityGroupName" name:"SecurityGroupName"`
		IpVersion         *string   `json:"IpVersion" name:"IpVersion"`
		Description       *string   `json:"Description" name:"Description"`
		InstanceCount     *int      `json:"InstanceCount" name:"InstanceCount"`
		CreateTime        *string   `json:"CreateTime" name:"CreateTime"`
		UpdateTime        *string   `json:"UpdateTime" name:"UpdateTime"`
		Rules             []*string `json:"Rules" name:"Rules"`
		Instances         []*string `json:"Instances" name:"Instances"`
	} `json:"Data"`
}

func (r *RenameSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	ProductType        *int    `json:"ProductType,omitempty" name:"ProductType"`
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
		SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		IpVersion         *string `json:"IpVersion" name:"IpVersion"`
		Description       *string `json:"Description" name:"Description"`
		InstanceCount     *int    `json:"InstanceCount" name:"InstanceCount"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime        *string `json:"UpdateTime" name:"UpdateTime"`
		Rules             []struct {
			RuleId      *string `json:"RuleId" name:"RuleId"`
			Cidr        *string `json:"Cidr" name:"Cidr"`
			CreateTime  *string `json:"CreateTime" name:"CreateTime"`
			Description *string `json:"Description" name:"Description"`
		} `json:"Rules" name:"Rules"`
		Instances []*string `json:"Instances" name:"Instances"`
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
	ProductType      *int      `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	InstanceIds      *string   `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	ProductType     *int      `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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

type CreateSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	ProductType     *int      `json:"ProductType,omitempty" name:"ProductType"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Cidrs           []*string `json:"Cidrs,omitempty" name:"Cidrs"`
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

type DeleteSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	ProductType     *int    `json:"ProductType,omitempty" name:"ProductType"`
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
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	ProjectIds      []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`
	FuzzySearch     *string   `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset          *int      `json:"Offset,omitempty" name:"Offset"`
	Limit           *int      `json:"Limit,omitempty" name:"Limit"`
	OrderBy         []*string `json:"OrderBy,omitempty" name:"OrderBy"`
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
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		InstanceName    *string `json:"InstanceName" name:"InstanceName"`
		ProductType     *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		Vip             *string `json:"Vip" name:"Vip"`
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		Status          *string `json:"Status" name:"Status"`
		UpdatetTime     *string `json:"UpdatetTime" name:"UpdatetTime"`
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
	FuzzySearch *string   `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	Offset      *int      `json:"Offset,omitempty" name:"Offset"`
	Limit       *int      `json:"Limit,omitempty" name:"Limit"`
	OrderBy     []*string `json:"OrderBy,omitempty" name:"OrderBy"`
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
		InstanceId      *string `json:"InstanceId" name:"InstanceId"`
		InstanceName    *string `json:"InstanceName" name:"InstanceName"`
		ProductType     *int    `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		Vip             *string `json:"Vip" name:"Vip"`
		CreateTime      *string `json:"CreateTime" name:"CreateTime"`
		Status          *string `json:"Status" name:"Status"`
		UpdatetTime     *string `json:"UpdatetTime" name:"UpdatetTime"`
	} `json:"Data"`
}

func (r *ListUnsecuredInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUnsecuredInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecycledInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ProductType *string   `json:"ProductType,omitempty" name:"ProductType"`
	ProjectIds  []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`
	FuzzySearch *string   `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	OrderBy     []*string `json:"OrderBy,omitempty" name:"OrderBy"`
	Offset      *int      `json:"Offset,omitempty" name:"Offset"`
	Limit       *int      `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListRecycledInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListRecycledInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Code      *string   `json:"Code" name:"Code"`
	Message   *string   `json:"Message" name:"Message"`
	Offset    *int      `json:"Offset" name:"Offset"`
	Limit     *int      `json:"Limit" name:"Limit"`
	Total     *int      `json:"Total" name:"Total"`
	Data      []*string `json:"Data" name:"Data"`
}

func (r *ListRecycledInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRecycledInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverRecycledInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RecoverRecycledInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RecoverRecycledInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg        *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *RecoverRecycledInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverRecycledInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DropRecycledInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DropRecycledInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DropRecycledInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		OperStatus *string `json:"OperStatus" name:"OperStatus"`
		Msg        *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DropRecycledInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DropRecycledInstanceResponse) FromJsonString(s string) error {
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
		RegionId     *string `json:"RegionId" name:"RegionId"`
		InnerCode    *string `json:"InnerCode" name:"InnerCode"`
		RegionCode   *string `json:"RegionCode" name:"RegionCode"`
		RegionName   *string `json:"RegionName" name:"RegionName"`
		RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
		Active       *bool   `json:"Active" name:"Active"`
		RegionType   *int    `json:"RegionType" name:"RegionType"`
		Overseas     *bool   `json:"Overseas" name:"Overseas"`
		AzList       []struct {
			AzCode *string `json:"AzCode" name:"AzCode"`
			AzName *string `json:"AzName" name:"AzName"`
		} `json:"AzList" name:"AzList"`
		AreaCode   *string `json:"AreaCode" name:"AreaCode"`
		AreaName   *string `json:"AreaName" name:"AreaName"`
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
	ProductType *int    `json:"ProductType,omitempty" name:"ProductType"`
	RegionCode  *string `json:"RegionCode,omitempty" name:"RegionCode"`
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
		RegionId     *string `json:"RegionId" name:"RegionId"`
		InnerCode    *string `json:"InnerCode" name:"InnerCode"`
		RegionCode   *string `json:"RegionCode" name:"RegionCode"`
		RegionName   *string `json:"RegionName" name:"RegionName"`
		RegionEnName *string `json:"RegionEnName" name:"RegionEnName"`
		Active       *bool   `json:"Active" name:"Active"`
		RegionType   *int    `json:"RegionType" name:"RegionType"`
		Overseas     *bool   `json:"Overseas" name:"Overseas"`
		AzList       []struct {
			AzCode *string `json:"AzCode" name:"AzCode"`
			AzName *string `json:"AzName" name:"AzName"`
		} `json:"AzList" name:"AzList"`
		AreaCode   *string `json:"AreaCode" name:"AreaCode"`
		AreaName   *string `json:"AreaName" name:"AreaName"`
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

type UpdateSecurityRuleRequest struct {
	*ksyunhttp.BaseRequest
	ProductType *int    `json:"ProductType,omitempty" name:"ProductType"`
	RuleId      *string `json:"RuleId,omitempty" name:"RuleId"`
	Description *string `json:"Description,omitempty" name:"Description"`
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

type RebindSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	ProductType     *int    `json:"ProductType,omitempty" name:"ProductType"`
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

type DescribeEngineDefaultParametersRequest struct {
	*ksyunhttp.BaseRequest
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	ConfigType    *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *DescribeEngineDefaultParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEngineDefaultParametersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Parameters []struct {
			Name            *string `json:"name" name:"name"`
			Min             *string `json:"min" name:"min"`
			Max             *string `json:"max" name:"max"`
			Visible         *int    `json:"visible" name:"visible"`
			RestartRequired *bool   `json:"restart_required" name:"restart_required"`
			Type            *string `json:"type" name:"type"`
			ConfigType      *string `json:"config_type" name:"config_type"`
			Enums           *string `json:"enums" name:"enums"`
			Default         *string `json:"default" name:"default"`
		} `json:"Parameters" name:"Parameters"`
	} `json:"Data"`
}

func (r *DescribeEngineDefaultParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineDefaultParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParameterGroupRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *ModifyDBParameterGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDBParameterGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *ModifyDBParameterGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParameterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceParametersRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *DescribeDBInstanceParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDBInstanceParametersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Configuration struct {
			Description *string `json:"description" name:"description"`
			Id          *string `json:"id" name:"id"`
			Name        *string `json:"name" name:"name"`
			Values      struct {
				CompileAggregateExpressions                     *string `json:"CompileAggregateExpressions" name:"CompileAggregateExpressions"`
				EnableUnalignedArrayJoin                        *string `json:"EnableUnalignedArrayJoin" name:"EnableUnalignedArrayJoin"`
				MaxMemoryUsageForUser                           *string `json:"MaxMemoryUsageForUser" name:"MaxMemoryUsageForUser"`
				EnableOptimizePredicateExpression               *string `json:"EnableOptimizePredicateExpression" name:"EnableOptimizePredicateExpression"`
				PreferGlobalInAndJoin                           *string `json:"PreferGlobalInAndJoin" name:"PreferGlobalInAndJoin"`
				MaxBytesInJoin                                  *string `json:"MaxBytesInJoin" name:"MaxBytesInJoin"`
				JoinUseNulls                                    *string `json:"JoinUseNulls" name:"JoinUseNulls"`
				JoinDefaultStrictness                           *string `json:"JoinDefaultStrictness" name:"JoinDefaultStrictness"`
				MaxRowsInJoin                                   *string `json:"MaxRowsInJoin" name:"MaxRowsInJoin"`
				DistributedAggregationMemoryEfficient           *string `json:"DistributedAggregationMemoryEfficient" name:"DistributedAggregationMemoryEfficient"`
				JoinOnDiskMaxFilesToMerge                       *string `json:"JoinOnDiskMaxFilesToMerge" name:"JoinOnDiskMaxFilesToMerge"`
				UnionDefaultMode                                *string `json:"UnionDefaultMode" name:"UnionDefaultMode"`
				OptimizeMoveToPrewhere                          *string `json:"OptimizeMoveToPrewhere" name:"OptimizeMoveToPrewhere"`
				MaxRowsToGroupBy                                *string `json:"MaxRowsToGroupBy" name:"MaxRowsToGroupBy"`
				OptimizeThrowIfNoop                             *string `json:"OptimizeThrowIfNoop" name:"OptimizeThrowIfNoop"`
				JoinAnyTakeLastRow                              *string `json:"JoinAnyTakeLastRow" name:"JoinAnyTakeLastRow"`
				MaxMemoryUsage                                  *string `json:"MaxMemoryUsage" name:"MaxMemoryUsage"`
				MinCountToCompileAggregateExpression            *string `json:"MinCountToCompileAggregateExpression" name:"MinCountToCompileAggregateExpression"`
				MaxBlockSize                                    *string `json:"MaxBlockSize" name:"MaxBlockSize"`
				MinCountToCompileExpression                     *string `json:"MinCountToCompileExpression" name:"MinCountToCompileExpression"`
				MutationsSync                                   *string `json:"MutationsSync" name:"MutationsSync"`
				MaxExecutionTime                                *string `json:"MaxExecutionTime" name:"MaxExecutionTime"`
				PartialMergeJoinRowsInRightBlocks               *string `json:"PartialMergeJoinRowsInRightBlocks" name:"PartialMergeJoinRowsInRightBlocks"`
				MaxBytesBeforeExternalGroupBy                   *string `json:"MaxBytesBeforeExternalGroupBy" name:"MaxBytesBeforeExternalGroupBy"`
				CompileExpressions                              *string `json:"CompileExpressions" name:"CompileExpressions"`
				DistributedProductMode                          *string `json:"DistributedProductMode" name:"DistributedProductMode"`
				MaxInsertBlockSize                              *string `json:"MaxInsertBlockSize" name:"MaxInsertBlockSize"`
				TotalsMode                                      *string `json:"TotalsMode" name:"TotalsMode"`
				JoinAlgorithm                                   *string `json:"JoinAlgorithm" name:"JoinAlgorithm"`
				OptimizeAggregationInOrder                      *string `json:"OptimizeAggregationInOrder" name:"OptimizeAggregationInOrder"`
				JoinOverflowMode                                *string `json:"JoinOverflowMode" name:"JoinOverflowMode"`
				UseUncompressedCache                            *string `json:"UseUncompressedCache" name:"UseUncompressedCache"`
				AllowExperimentalAlterMaterializedViewStructure *string `json:"AllowExperimentalAlterMaterializedViewStructure" name:"AllowExperimentalAlterMaterializedViewStructure"`
				InsertNullAsDefault                             *string `json:"InsertNullAsDefault" name:"InsertNullAsDefault"`
				OptimizeReadInOrder                             *string `json:"OptimizeReadInOrder" name:"OptimizeReadInOrder"`
				MaxBytesBeforeExternalSort                      *string `json:"MaxBytesBeforeExternalSort" name:"MaxBytesBeforeExternalSort"`
				AnyJoinDistinctRightTableKeys                   *string `json:"AnyJoinDistinctRightTableKeys" name:"AnyJoinDistinctRightTableKeys"`
				GroupByOverflowMode                             *string `json:"GroupByOverflowMode" name:"GroupByOverflowMode"`
				OptimizeMoveToPrewhereIfFinal                   *string `json:"OptimizeMoveToPrewhereIfFinal" name:"OptimizeMoveToPrewhereIfFinal"`
				PartialMergeJoinOptimizations                   *string `json:"PartialMergeJoinOptimizations" name:"PartialMergeJoinOptimizations"`
			} `json:"Values"`
			DatastoreVersionId *string `json:"datastore_version_id" name:"datastore_version_id"`
		} `json:"Configuration" name:"Configuration"`
	} `json:"Data"`
}

func (r *DescribeDBInstanceParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDBParameterRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *ResetDBParameterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResetDBParameterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Configuration struct {
			Description *string `json:"description" name:"description"`
			Id          *string `json:"id" name:"id"`
			Name        *string `json:"name" name:"name"`
			Values      struct {
				CompileAggregateExpressions                     *string `json:"CompileAggregateExpressions" name:"CompileAggregateExpressions"`
				EnableUnalignedArrayJoin                        *string `json:"EnableUnalignedArrayJoin" name:"EnableUnalignedArrayJoin"`
				MaxMemoryUsageForUser                           *string `json:"MaxMemoryUsageForUser" name:"MaxMemoryUsageForUser"`
				EnableOptimizePredicateExpression               *string `json:"EnableOptimizePredicateExpression" name:"EnableOptimizePredicateExpression"`
				PreferGlobalInAndJoin                           *string `json:"PreferGlobalInAndJoin" name:"PreferGlobalInAndJoin"`
				MaxBytesInJoin                                  *string `json:"MaxBytesInJoin" name:"MaxBytesInJoin"`
				JoinUseNulls                                    *string `json:"JoinUseNulls" name:"JoinUseNulls"`
				JoinDefaultStrictness                           *string `json:"JoinDefaultStrictness" name:"JoinDefaultStrictness"`
				MaxRowsInJoin                                   *string `json:"MaxRowsInJoin" name:"MaxRowsInJoin"`
				DistributedAggregationMemoryEfficient           *string `json:"DistributedAggregationMemoryEfficient" name:"DistributedAggregationMemoryEfficient"`
				JoinOnDiskMaxFilesToMerge                       *string `json:"JoinOnDiskMaxFilesToMerge" name:"JoinOnDiskMaxFilesToMerge"`
				UnionDefaultMode                                *string `json:"UnionDefaultMode" name:"UnionDefaultMode"`
				OptimizeMoveToPrewhere                          *string `json:"OptimizeMoveToPrewhere" name:"OptimizeMoveToPrewhere"`
				MaxRowsToGroupBy                                *string `json:"MaxRowsToGroupBy" name:"MaxRowsToGroupBy"`
				OptimizeThrowIfNoop                             *string `json:"OptimizeThrowIfNoop" name:"OptimizeThrowIfNoop"`
				JoinAnyTakeLastRow                              *string `json:"JoinAnyTakeLastRow" name:"JoinAnyTakeLastRow"`
				MaxMemoryUsage                                  *string `json:"MaxMemoryUsage" name:"MaxMemoryUsage"`
				MinCountToCompileAggregateExpression            *string `json:"MinCountToCompileAggregateExpression" name:"MinCountToCompileAggregateExpression"`
				MaxBlockSize                                    *string `json:"MaxBlockSize" name:"MaxBlockSize"`
				MinCountToCompileExpression                     *string `json:"MinCountToCompileExpression" name:"MinCountToCompileExpression"`
				MutationsSync                                   *string `json:"MutationsSync" name:"MutationsSync"`
				MaxExecutionTime                                *string `json:"MaxExecutionTime" name:"MaxExecutionTime"`
				PartialMergeJoinRowsInRightBlocks               *string `json:"PartialMergeJoinRowsInRightBlocks" name:"PartialMergeJoinRowsInRightBlocks"`
				MaxBytesBeforeExternalGroupBy                   *string `json:"MaxBytesBeforeExternalGroupBy" name:"MaxBytesBeforeExternalGroupBy"`
				CompileExpressions                              *string `json:"CompileExpressions" name:"CompileExpressions"`
				DistributedProductMode                          *string `json:"DistributedProductMode" name:"DistributedProductMode"`
				MaxInsertBlockSize                              *string `json:"MaxInsertBlockSize" name:"MaxInsertBlockSize"`
				TotalsMode                                      *string `json:"TotalsMode" name:"TotalsMode"`
				JoinAlgorithm                                   *string `json:"JoinAlgorithm" name:"JoinAlgorithm"`
				OptimizeAggregationInOrder                      *string `json:"OptimizeAggregationInOrder" name:"OptimizeAggregationInOrder"`
				JoinOverflowMode                                *string `json:"JoinOverflowMode" name:"JoinOverflowMode"`
				UseUncompressedCache                            *string `json:"UseUncompressedCache" name:"UseUncompressedCache"`
				AllowExperimentalAlterMaterializedViewStructure *string `json:"AllowExperimentalAlterMaterializedViewStructure" name:"AllowExperimentalAlterMaterializedViewStructure"`
				InsertNullAsDefault                             *string `json:"InsertNullAsDefault" name:"InsertNullAsDefault"`
				OptimizeReadInOrder                             *string `json:"OptimizeReadInOrder" name:"OptimizeReadInOrder"`
				MaxBytesBeforeExternalSort                      *string `json:"MaxBytesBeforeExternalSort" name:"MaxBytesBeforeExternalSort"`
				AnyJoinDistinctRightTableKeys                   *string `json:"AnyJoinDistinctRightTableKeys" name:"AnyJoinDistinctRightTableKeys"`
				GroupByOverflowMode                             *string `json:"GroupByOverflowMode" name:"GroupByOverflowMode"`
				OptimizeMoveToPrewhereIfFinal                   *string `json:"OptimizeMoveToPrewhereIfFinal" name:"OptimizeMoveToPrewhereIfFinal"`
				PartialMergeJoinOptimizations                   *string `json:"PartialMergeJoinOptimizations" name:"PartialMergeJoinOptimizations"`
			} `json:"Values"`
			DatastoreVersionId *string `json:"datastore_version_id" name:"datastore_version_id"`
		} `json:"Configuration" name:"Configuration"`
	} `json:"Data"`
}

func (r *ResetDBParameterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDBParameterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBucketsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeBucketsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBucketsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Code      *string   `json:"Code" name:"Code"`
	Message   *string   `json:"Message" name:"Message"`
	Data      []*string `json:"Data" name:"Data"`
}

func (r *DescribeBucketsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBucketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateHotAndColdSeparationRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *OperateHotAndColdSeparationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type OperateHotAndColdSeparationResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *OperateHotAndColdSeparationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateHotAndColdSeparationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountRequest struct {
	*ksyunhttp.BaseRequest
	Name        *string `json:"Name,omitempty" name:"Name"`
	Password    *string `json:"Password,omitempty" name:"Password"`
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		ModifyPrivileges *bool `json:"ModifyPrivileges" name:"ModifyPrivileges"`
		CreateAccount    *bool `json:"CreateAccount" name:"CreateAccount"`
	} `json:"Data"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountPrivilegesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId                *string                                                     `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceAccountName       *string                                                     `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPrivileges []*ModifyInstanceAccountPrivilegesInstanceAccountPrivileges `json:"InstanceAccountPrivileges,omitempty" name:"InstanceAccountPrivileges"`
}

func (r *ModifyInstanceAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceAccountPrivilegesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
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
	InstanceId          *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceAccountName *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceAccountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Msg *string `json:"Msg" name:"Msg"`
	} `json:"Data"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceAccountsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Accounts []struct {
			Status    *string `json:"status" name:"status"`
			Databases []struct {
				Name       *string `json:"Name" name:"Name"`
				Privileges *string `json:"Privileges" name:"Privileges"`
			} `json:"Databases"`
			Name    *string `json:"name" name:"name"`
			AccType *int    `json:"acc_type" name:"acc_type"`
		} `json:"Accounts" name:"Accounts"`
	} `json:"Data"`
}

func (r *DescribeInstanceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDatabasesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceDatabasesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Databases []struct {
			Name     *string `json:"name" name:"name"`
			Accounts []struct {
				Name       *string `json:"Name" name:"Name"`
				Privileges *string `json:"Privileges" name:"Privileges"`
			} `json:"Accounts"`
			Status *string `json:"status" name:"status"`
		} `json:"Databases" name:"Databases"`
	} `json:"Data"`
}

func (r *DescribeInstanceDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountInfoRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId                 *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceAccountName        *string `json:"InstanceAccountName,omitempty" name:"InstanceAccountName"`
	InstanceAccountPassword    *string `json:"InstanceAccountPassword,omitempty" name:"InstanceAccountPassword"`
	InstanceAccountDescription *string `json:"InstanceAccountDescription,omitempty" name:"InstanceAccountDescription"`
}

func (r *ModifyInstanceAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceAccountInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *ModifyInstanceAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceShardInfoRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceShardInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceShardInfoResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Zookeeper []struct {
			Id     *string `json:"Id" name:"Id"`
			Vcpus  *int    `json:"Vcpus" name:"Vcpus"`
			Ram    *int    `json:"Ram" name:"Ram"`
			Disk   *int    `json:"Disk" name:"Disk"`
			Name   *string `json:"Name" name:"Name"`
			ZkName *string `json:"ZkName" name:"ZkName"`
		} `json:"Zookeeper" name:"Zookeeper"`
		Clickhouse struct {
			ClusterName *string `json:"ClusterName" name:"ClusterName"`
			ClusterId   *string `json:"ClusterId" name:"ClusterId"`
			ClusterVip  *string `json:"ClusterVip" name:"ClusterVip"`
			ColdUsage   *string `json:"ColdUsage" name:"ColdUsage"`
			Shards      []struct {
			} `json:"Shards"`
		} `json:"Clickhouse" name:"Clickhouse"`
	} `json:"Data"`
}

func (r *DescribeInstanceShardInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceShardInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
