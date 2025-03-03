package v20210101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type ModifyDBParameterGroupParameters struct {
}

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

func (r *ListInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceResponse struct {
	*ksyunhttp.BaseResponse
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
		Replicas             *int `json:"Replicas" name:"Replicas"`
		DirectConnectionVips []struct {
		} `json:"DirectConnectionVips" name:"DirectConnectionVips"`
		MultiAz *int `json:"MultiAz" name:"MultiAz"`
		Area    struct {
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

func (r *ListSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
		} `json:"Rules" name:"Rules"`
		Instances []struct {
		} `json:"Instances" name:"Instances"`
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

func (r *RenameSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenameSecurityGroupResponse struct {
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
		} `json:"Rules" name:"Rules"`
		Instances []struct {
		} `json:"Instances" name:"Instances"`
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
		} `json:"Instances" name:"Instances"`
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

func (r *BindSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BindSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *UnbindSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnbindSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListSecuredInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSecuredInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListUnsecuredInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListUnsecuredInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListRecycledInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRecycledInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *RecoverRecycledInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RecoverRecycledInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DropRecycledInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DropRecycledInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *ListRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *UpdateSecurityRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateSecurityRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *RebindSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RebindSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
	InstanceId *string                             `json:"InstanceId,omitempty" name:"InstanceId"`
	Parameters []*ModifyDBParameterGroupParameters `json:"Parameters,omitempty" name:"Parameters"`
	ConfigType *string                             `json:"ConfigType,omitempty" name:"ConfigType"`
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

func (r *ResetDBParameterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetDBParameterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeBucketsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBucketsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *OperateHotAndColdSeparationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "OperateHotAndColdSeparationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
