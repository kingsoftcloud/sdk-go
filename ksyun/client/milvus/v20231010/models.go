package v20231010

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type CreateSecurityRulesRules struct {
	Cidr   *string `json:"Cidr,omitempty" name:"Cidr"`
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}
type CreateSecurityGroupRules struct {
	Cidr   *string `json:"Cidr,omitempty" name:"Cidr"`
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}
type CreateBackupDBCollection struct {
	DbName      *string   `json:"DbName,omitempty" name:"DbName"`
	Description *string   `json:"Description,omitempty" name:"Description"`
	Collections []*string `json:"Collections,omitempty" name:"Collections"`
}

type DeleteSecurityRulesRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	RuleIds         *string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DeleteSecurityRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSecurityRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *DeleteSecurityRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRulesRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string                     `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Rules           []*CreateSecurityRulesRules `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateSecurityRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateSecurityRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *CreateSecurityRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindSecurityGroupInstancesRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstancesIds    *string `json:"InstancesIds,omitempty" name:"InstancesIds"`
}

func (r *UnbindSecurityGroupInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UnbindSecurityGroupInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *UnbindSecurityGroupInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindSecurityGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSecurityGroupInstancesRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	InstancesIds    *string `json:"InstancesIds,omitempty" name:"InstancesIds"`
}

func (r *BindSecurityGroupInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type BindSecurityGroupInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *BindSecurityGroupInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSecurityGroupInstancesResponse) FromJsonString(s string) error {
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
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      struct {
		Id        *string `json:"Id" name:"Id"`
		Name      *string `json:"Name" name:"Name"`
		IpVersion *string `json:"IpVersion" name:"IpVersion"`
		CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt *string `json:"UpdatedAt" name:"UpdatedAt"`
		Detail    *string `json:"Detail" name:"Detail"`
		Rules     []struct {
			Cidr      *string `json:"Cidr" name:"Cidr"`
			Detail    *string `json:"Detail" name:"Detail"`
			Id        *string `json:"Id" name:"Id"`
			CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
			UpdatedAt *string `json:"UpdatedAt" name:"UpdatedAt"`
		} `json:"Rules" name:"Rules"`
		Instances []struct {
			Id        *string `json:"Id" name:"Id"`
			Name      *string `json:"Name" name:"Name"`
			CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
			Addresses []struct {
				Host *string `json:"Host" name:"Host"`
				Type *string `json:"Type" name:"Type"`
			} `json:"Addresses"`
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
	Name      *string                     `json:"Name,omitempty" name:"Name"`
	IpVersion *string                     `json:"IpVersion,omitempty" name:"IpVersion"`
	Detail    *string                     `json:"Detail,omitempty" name:"Detail"`
	Rules     []*CreateSecurityGroupRules `json:"Rules,omitempty" name:"Rules"`
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
		Id        *string `json:"Id" name:"Id"`
		Name      *string `json:"Name" name:"Name"`
		IpVersion *string `json:"IpVersion" name:"IpVersion"`
		CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt *string `json:"UpdatedAt" name:"UpdatedAt"`
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
	Offset *int `json:"Offset,omitempty" name:"Offset"`
	Limit  *int `json:"Limit,omitempty" name:"Limit"`
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
		Id        *string `json:"Id" name:"Id"`
		Name      *string `json:"Name" name:"Name"`
		IpVersion *string `json:"IpVersion" name:"IpVersion"`
		CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt *string `json:"UpdatedAt" name:"UpdatedAt"`
		Detail    *string `json:"Detail" name:"Detail"`
		Rules     []struct {
			Id        *string `json:"Id" name:"Id"`
			Cidr      *string `json:"Cidr" name:"Cidr"`
			Detail    *string `json:"Detail" name:"Detail"`
			CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
			UpdatedAt *string `json:"UpdatedAt" name:"UpdatedAt"`
		} `json:"Rules" name:"Rules"`
		Instances []struct {
			Id        *string   `json:"Id" name:"Id"`
			Name      *string   `json:"Name" name:"Name"`
			CreatedAt *string   `json:"CreatedAt" name:"CreatedAt"`
			Addresses []*string `json:"Addresses" name:"Addresses"`
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
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		Az           *string `json:"Az" name:"Az"`
		AzName       *string `json:"AzName" name:"AzName"`
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		State        *string `json:"State" name:"State"`
		StateName    *string `json:"StateName" name:"StateName"`
		Engine       *string `json:"Engine" name:"Engine"`
		CreateDate   *string `json:"CreateDate" name:"CreateDate"`
		Addresses    []struct {
			Host *string `json:"Host" name:"Host"`
			Port *int    `json:"Port" name:"Port"`
			Type *string `json:"Type" name:"Type"`
		} `json:"Addresses" name:"Addresses"`
		Resources struct {
			CU struct {
				CU       *int    `json:"CU" name:"CU"`
				Type     *string `json:"Type" name:"Type"`
				TypeName *string `json:"TypeName" name:"TypeName"`
			} `json:"CU"`
		} `json:"Resources" name:"Resources"`
		SubnetId    *string `json:"SubnetId" name:"SubnetId"`
		VnetId      *string `json:"VnetId" name:"VnetId"`
		MonitorLink *string `json:"MonitorLink" name:"MonitorLink"`
		Users       []struct {
			Id        *string `json:"Id" name:"Id"`
			Name      *string `json:"Name" name:"Name"`
			IsAdmin   *bool   `json:"IsAdmin" name:"IsAdmin"`
			CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
		} `json:"Users" name:"Users"`
		SecurityGroupList []*string `json:"SecurityGroupList" name:"SecurityGroupList"`
		ProjectId         *string   `json:"ProjectId" name:"ProjectId"`
		ProjectName       *string   `json:"ProjectName" name:"ProjectName"`
		BillType          *int      `json:"BillType" name:"BillType"`
		ProductId         *string   `json:"ProductId" name:"ProductId"`
		ProductType       *int      `json:"ProductType" name:"ProductType"`
		ProductWhat       *int      `json:"ProductWhat" name:"ProductWhat"`
		CreateDateOrder   *string   `json:"CreateDateOrder" name:"CreateDateOrder"`
		Region            *string   `json:"Region" name:"Region"`
	} `json:"Data"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceRequest struct {
	*ksyunhttp.BaseRequest
	Offset           *string `json:"Offset,omitempty" name:"Offset"`
	Limit            *string `json:"Limit,omitempty" name:"Limit"`
	FuzzySearch      *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	SecBindingStatus *string `json:"SecBindingStatus,omitempty" name:"SecBindingStatus"`
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
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		Az           *string `json:"Az" name:"Az"`
		AzName       *string `json:"AzName" name:"AzName"`
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		State        *string `json:"State" name:"State"`
		StateName    *string `json:"StateName" name:"StateName"`
		Engine       *string `json:"Engine" name:"Engine"`
		CreateDate   *string `json:"CreateDate" name:"CreateDate"`
		Addresses    []struct {
			Host *string `json:"Host" name:"Host"`
			Port *int    `json:"Port" name:"Port"`
			Type *string `json:"Type" name:"Type"`
		} `json:"Addresses" name:"Addresses"`
		Resources struct {
			CU struct {
				CU       *int    `json:"CU" name:"CU"`
				Type     *string `json:"Type" name:"Type"`
				TypeName *string `json:"TypeName" name:"TypeName"`
			} `json:"CU"`
		} `json:"Resources" name:"Resources"`
		Users             []*string `json:"Users" name:"Users"`
		SecurityGroupList []*string `json:"SecurityGroupList" name:"SecurityGroupList"`
		ProjectId         *string   `json:"ProjectId" name:"ProjectId"`
		ProjectName       *string   `json:"ProjectName" name:"ProjectName"`
		BillType          *int      `json:"BillType" name:"BillType"`
		ProductId         *string   `json:"ProductId" name:"ProductId"`
		ProductType       *int      `json:"ProductType" name:"ProductType"`
		ProductWhat       *int      `json:"ProductWhat" name:"ProductWhat"`
		CreateDateOrder   *string   `json:"CreateDateOrder" name:"CreateDateOrder"`
		Region            *string   `json:"Region" name:"Region"`
	} `json:"Data"`
}

func (r *ListInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceName   *string   `json:"InstanceName,omitempty" name:"InstanceName"`
	AdminPassword  *string   `json:"AdminPassword,omitempty" name:"AdminPassword"`
	SubnetId       *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	VnetId         *string   `json:"VnetId,omitempty" name:"VnetId"`
	ProjectId      *string   `json:"ProjectId,omitempty" name:"ProjectId"`
	ProductType    *string   `json:"ProductType,omitempty" name:"ProductType"`
	DBInstanceType *string   `json:"DBInstanceType,omitempty" name:"DBInstanceType"`
	Az             []*string `json:"Az,omitempty" name:"Az"`
	Engine         *string   `json:"Engine,omitempty" name:"Engine"`
	EngineVersion  *string   `json:"EngineVersion,omitempty" name:"EngineVersion"`
	AdminUser      *string   `json:"AdminUser,omitempty" name:"AdminUser"`
	Cu             *string   `json:"Cu,omitempty" name:"Cu"`
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

type ReleaseDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseDBInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReleaseDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *ReleaseDBInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseDBInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateDBInstanceEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Port       *int    `json:"Port,omitempty" name:"Port"`
}

func (r *AllocateDBInstanceEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AllocateDBInstanceEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      *bool   `json:"Data" name:"Data"`
}

func (r *AllocateDBInstanceEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateDBInstanceEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBackupRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
	Offset     *int    `json:"Offset,omitempty" name:"Offset"`
	Limit      *int    `json:"Limit,omitempty" name:"Limit"`
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
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
	Total     *int    `json:"Total" name:"Total"`
	Data      []struct {
		BackupId        *string `json:"BackupId" name:"BackupId"`
		BackupName      *string `json:"BackupName" name:"BackupName"`
		BackupState     *string `json:"BackupState" name:"BackupState"`
		BackupMethod    *string `json:"BackupMethod" name:"BackupMethod"`
		BackupDimension *string `json:"BackupDimension" name:"BackupDimension"`
		BackupSize      *int    `json:"BackupSize" name:"BackupSize"`
		RetentionDays   *int    `json:"RetentionDays" name:"RetentionDays"`
		Collections     []struct {
			DbName         *string `json:"DbName" name:"DbName"`
			CollectionName *string `json:"CollectionName" name:"CollectionName"`
			LoadState      *string `json:"LoadState" name:"LoadState"`
			Size           *int    `json:"Size" name:"Size"`
			Description    *string `json:"Description" name:"Description"`
		} `json:"Collections" name:"Collections"`
		CreatedAt *string `json:"CreatedAt" name:"CreatedAt"`
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
	InstanceId   *string                     `json:"InstanceId,omitempty" name:"InstanceId"`
	BackupName   *string                     `json:"BackupName,omitempty" name:"BackupName"`
	DBCollection []*CreateBackupDBCollection `json:"DBCollection,omitempty" name:"DBCollection"`
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
	Data      struct {
		Id                *string   `json:"Id" name:"Id"`
		Name              *string   `json:"Name" name:"Name"`
		State             *string   `json:"State" name:"State"`
		Method            *string   `json:"Method" name:"Method"`
		Dimension         *string   `json:"Dimension" name:"Dimension"`
		Size              *int      `json:"Size" name:"Size"`
		RetentionDays     *int      `json:"RetentionDays" name:"RetentionDays"`
		CreatedAt         *string   `json:"CreatedAt" name:"CreatedAt"`
		CollectionBackups []*string `json:"CollectionBackups" name:"CollectionBackups"`
	} `json:"Data"`
}

func (r *CreateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
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
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Code      *string   `json:"Code" name:"Code"`
	Message   *string   `json:"Message" name:"Message"`
	Data      []*string `json:"Data" name:"Data"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCollectionsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Offset     *int    `json:"Offset,omitempty" name:"Offset"`
	Limit      *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListCollectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListCollectionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *string `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
	Total     *int    `json:"Total" name:"Total"`
	Data      []struct {
		DbName      *string   `json:"DbName" name:"DbName"`
		Collections []*string `json:"Collections" name:"Collections"`
	} `json:"Data"`
}

func (r *ListCollectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCollectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceTrialOrderRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	Duration    *int    `json:"Duration,omitempty" name:"Duration"`
	BillType    *int    `json:"BillType,omitempty" name:"BillType"`
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
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		OrderId      *string `json:"OrderId" name:"OrderId"`
		AccountId    *string `json:"AccountId" name:"AccountId"`
		Region       *string `json:"Region" name:"Region"`
		Az           *string `json:"Az" name:"Az"`
	} `json:"Data"`
}

func (r *UpdateInstanceTrialOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceTrialOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
