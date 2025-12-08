package v20170101
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreateMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	PayType          *string   `json:"PayType,omitempty" name:"PayType"`
	AvailabilityZone []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Name             *string   `json:"Name,omitempty" name:"Name"`
	DbVersion        *string   `json:"DbVersion,omitempty" name:"DbVersion"`
	NodeNum          *int      `json:"NodeNum,omitempty" name:"NodeNum"`
	Storage          *int      `json:"Storage,omitempty" name:"Storage"`
	Duration         *int      `json:"Duration,omitempty" name:"Duration"`
	IamProjectId     *string   `json:"IamProjectId,omitempty" name:"IamProjectId"`
	VpcId            *string   `json:"VpcId,omitempty" name:"VpcId"`
	VnetId           *string   `json:"VnetId,omitempty" name:"VnetId"`
	InstancePassword *string   `json:"InstancePassword,omitempty" name:"InstancePassword"`
	InstanceClass    *string   `json:"InstanceClass,omitempty" name:"InstanceClass"`
}

func (r *CreateMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId *string `json:"UserId" name:"UserId"`
		Region *string `json:"Region" name:"Region"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Name   *string `json:"Name" name:"Name"`
	} `json:"MongoDBInstanceResult"`
}

func (r *CreateMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *DeleteMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId        *string `json:"UserId" name:"UserId"`
		Region        *string `json:"Region" name:"Region"`
		Name          *string `json:"Name" name:"Name"`
		InstanceId    *string `json:"InstanceId" name:"InstanceId"`
		Status        *string `json:"Status" name:"Status"`
		IP            *string `json:"IP" name:"IP"`
		InstanceType  *string `json:"InstanceType" name:"InstanceType"`
		Version       *string `json:"Version" name:"Version"`
		InstanceClass *string `json:"InstanceClass" name:"InstanceClass"`
		Storage       *int    `json:"Storage" name:"Storage"`
		Port          *int    `json:"Port" name:"Port"`
		NetworkType   *string `json:"NetworkType" name:"NetworkType"`
		VpcId         *string `json:"VpcId" name:"VpcId"`
		VnetId        *string `json:"VnetId" name:"VnetId"`
		TimingSwitch  *string `json:"TimingSwitch" name:"TimingSwitch"`
		Timezone      *string `json:"Timezone" name:"Timezone"`
		TimeCycle     *string `json:"TimeCycle" name:"TimeCycle"`
		ProductId     *string `json:"ProductId" name:"ProductId"`
		ProductWhat   *int    `json:"ProductWhat" name:"ProductWhat"`
		PayType       *string `json:"PayType" name:"PayType"`
		CreateDate    *string `json:"CreateDate" name:"CreateDate"`
		ExpirationDate *string `json:"ExpirationDate" name:"ExpirationDate"`
		IamProjectId  *string `json:"IamProjectId" name:"IamProjectId"`
		IamProjectName *string `json:"IamProjectName" name:"IamProjectName"`
	} `json:"MongoDBInstanceResult"`
}

func (r *DescribeMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMongoDBInstancesRequest struct {
	*ksyunhttp.BaseRequest
	Area         *string `json:"Area,omitempty" name:"Area"`
	Vip          *string `json:"Vip,omitempty" name:"Vip"`
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	VnetId       *string `json:"VnetId,omitempty" name:"VnetId"`
	IamProjectId *string `json:"IamProjectId,omitempty" name:"IamProjectId"`
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Name         *string `json:"Name,omitempty" name:"Name"`
	Mode         *string `json:"Mode,omitempty" name:"Mode"`
	DbVersion    *string `json:"DbVersion,omitempty" name:"DbVersion"`
	Status       *string `json:"Status,omitempty" name:"Status"`
	FuzzySearch  *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	TagKey       *string `json:"TagKey,omitempty" name:"TagKey"`
	TagValue     *string `json:"TagValue,omitempty" name:"TagValue"`
	Offset       *int    `json:"Offset,omitempty" name:"Offset"`
	Limit        *int    `json:"Limit,omitempty" name:"Limit"`
	OrderBy      *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeMongoDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMongoDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	MongoDBInstancesResult []struct {
		IamProjectId  *string `json:"IamProjectId" name:"IamProjectId"`
		IamProjectName *string `json:"IamProjectName" name:"IamProjectName"`
		UserId        *string `json:"UserId" name:"UserId"`
		Region        *string `json:"Region" name:"Region"`
		Name          *string `json:"Name" name:"Name"`
		InstanceId    *string `json:"InstanceId" name:"InstanceId"`
		Status        *string `json:"Status" name:"Status"`
		IP            *string `json:"IP" name:"IP"`
		InstanceType  *string `json:"InstanceType" name:"InstanceType"`
		Version       *string `json:"Version" name:"Version"`
		InstanceClass *string `json:"InstanceClass" name:"InstanceClass"`
		Storage       *int    `json:"Storage" name:"Storage"`
		Port          *string `json:"Port" name:"Port"`
		NetworkType   *string `json:"NetworkType" name:"NetworkType"`
		VpcId         *string `json:"VpcId" name:"VpcId"`
		VnetId        *string `json:"VnetId" name:"VnetId"`
		ProductId     *string `json:"ProductId" name:"ProductId"`
		ProductWhat   *int    `json:"ProductWhat" name:"ProductWhat"`
		PayType       *string `json:"PayType" name:"PayType"`
		CreateDate    *string `json:"CreateDate" name:"CreateDate"`
		ExpirationDate *string `json:"ExpirationDate" name:"ExpirationDate"`
	} `json:"MongoDBInstancesResult"`
	Offset *int `json:"Offset" name:"Offset"`
	Limit  *int `json:"Limit" name:"Limit"`
	Total  *int `json:"Total" name:"Total"`
}

func (r *DescribeMongoDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMongoDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMongoDBInstanceNodeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeId     *string `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DescribeMongoDBInstanceNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMongoDBInstanceNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceNodeResult []struct {
		NodeId *string `json:"NodeId" name:"NodeId"`
		Name *string `json:"Name" name:"Name"`
		Role *string `json:"Role" name:"Role"`
		IP   *string `json:"IP" name:"IP"`
		Port *int    `json:"Port" name:"Port"`
		Status *string `json:"Status" name:"Status"`
	} `json:"MongoDBInstanceNodeResult"`
}

func (r *DescribeMongoDBInstanceNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMongoDBInstanceNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RenameMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Name       *string `json:"Name,omitempty" name:"Name"`
}

func (r *RenameMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenameMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	Name       *string `json:"name" name:"name"`
}

func (r *RenameMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ResetPasswordMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId       *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
}

func (r *ResetPasswordMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResetPasswordMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *ResetPasswordMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetPasswordMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RestartMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RestartMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RestartMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *RestartMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateMongoDBSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Name       *string `json:"Name,omitempty" name:"Name"`
	BackupMode *string `json:"BackupMode,omitempty" name:"BackupMode"`
}

func (r *CreateMongoDBSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMongoDBSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBSnapshotResult struct {
		SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
		Name   *string `json:"Name" name:"Name"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Status *string `json:"Status" name:"Status"`
		Create *string `json:"Create" name:"Create"`
	} `json:"MongoDBSnapshotResult"`
}

func (r *CreateMongoDBSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMongoDBSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type SetMongoDBTimingSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TimingSwitch *string `json:"TimingSwitch,omitempty" name:"TimingSwitch"`
	Timezone     *string `json:"Timezone,omitempty" name:"Timezone"`
	TimeCycle    *string `json:"TimeCycle,omitempty" name:"TimeCycle"`
}

func (r *SetMongoDBTimingSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetMongoDBTimingSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *SetMongoDBTimingSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMongoDBTimingSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMongoDBSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMongoDBSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMongoDBSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBSnapshotResult []struct {
		SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
		Name      *string `json:"Name" name:"Name"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		TypeField *string `json:"TypeField" name:"TypeField"`
		Size      *int    `json:"Size" name:"Size"`
		Status    *string `json:"Status" name:"Status"`
		BackupMode *string `json:"BackupMode" name:"BackupMode"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"MongoDBSnapshotResult"`
}

func (r *DescribeMongoDBSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMongoDBSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteMongoDBSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DeleteMongoDBSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteMongoDBSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
}

func (r *DeleteMongoDBSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMongoDBSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RenameMongoDBSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	Name       *string `json:"Name,omitempty" name:"Name"`
}

func (r *RenameMongoDBSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RenameMongoDBSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
}

func (r *RenameMongoDBSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameMongoDBSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddSecurityGroupRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Cidrs      *string `json:"Cidrs,omitempty" name:"Cidrs"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}

func (r *AddSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddSecurityGroupRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	MongoDBSecurityGroupRule []struct {
		ToPort *string `json:"ToPort" name:"ToPort"`
		Cidr   *string `json:"Cidr" name:"Cidr"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		Id     *string `json:"Id" name:"Id"`
	} `json:"MongoDBSecurityGroupRule"`
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
	Cidrs      *string `json:"Cidrs,omitempty" name:"Cidrs"`
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteSecurityGroupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteSecurityGroupRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	MongoDBSecurityGroupRule []struct {
		ToPort    *string `json:"ToPort" name:"ToPort"`
		Cidr      *string `json:"Cidr" name:"Cidr"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		Id        *string `json:"Id" name:"Id"`
		TypeField *string `json:"TypeField" name:"TypeField"`
	} `json:"MongoDBSecurityGroupRule"`
}

func (r *DeleteSecurityGroupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListSecurityGroupRulesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ListSecurityGroupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListSecurityGroupRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	MongoDBSecurityGroupRule []struct {
		ToPort *string `json:"ToPort" name:"ToPort"`
		Cidr   *string `json:"Cidr" name:"Cidr"`
		FromPort *string `json:"FromPort" name:"FromPort"`
		Protocol *string `json:"Protocol" name:"Protocol"`
		Id     *string `json:"Id" name:"Id"`
	} `json:"MongoDBSecurityGroupRule"`
}

func (r *ListSecurityGroupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateMongoDBInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	Storage       *int    `json:"Storage,omitempty" name:"Storage"`
}

func (r *UpdateMongoDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId *string `json:"UserId" name:"UserId"`
		Region *string `json:"Region" name:"Region"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Name   *string `json:"Name" name:"Name"`
	} `json:"MongoDBInstanceResult"`
}

func (r *UpdateMongoDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMongoDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddSecondaryInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeNum    *string `json:"NodeNum,omitempty" name:"NodeNum"`
}

func (r *AddSecondaryInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddSecondaryInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *AddSecondaryInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSecondaryInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMongoDBShardNodeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMongoDBShardNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMongoDBShardNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	MongosNodeResult []struct {
		NodeId      *string `json:"NodeId" name:"NodeId"`
		Name        *string `json:"Name" name:"Name"`
		Role        *string `json:"Role" name:"Role"`
		Endpoint    *string `json:"Endpoint" name:"Endpoint"`
		Status      *string `json:"Status" name:"Status"`
		Connections *int    `json:"Connections" name:"Connections"`
		InstanceClass *string `json:"InstanceClass" name:"InstanceClass"`
	} `json:"MongosNodeResult"`
	ShardNodeResult []struct {
		NodeId *string `json:"NodeId" name:"NodeId"`
		Name   *string `json:"Name" name:"Name"`
		Status *string `json:"Status" name:"Status"`
		Disk   *int    `json:"Disk" name:"Disk"`
		Iops   *int    `json:"Iops" name:"Iops"`
		InstanceClass *string `json:"InstanceClass" name:"InstanceClass"`
	} `json:"ShardNodeResult"`
}

func (r *DescribeMongoDBShardNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMongoDBShardNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeValidRegionRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeValidRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeValidRegionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		RegionList []struct {
			AreaEnName *string `json:"areaEnName" name:"areaEnName"`
			RegionCode *string `json:"regionCode" name:"regionCode"`
			AreaCode   *string `json:"areaCode" name:"areaCode"`
			InnerCode  *string `json:"innerCode" name:"innerCode"`
			RegionType *int    `json:"regionType" name:"regionType"`
			AreaName   *string `json:"areaName" name:"areaName"`
			AzList     []struct {
				AzName *string `json:"AzName" name:"AzName"`
				AzCode *string `json:"AzCode" name:"AzCode"`
			} `json:"AzList"`
			RegionName   *string `json:"regionName" name:"regionName"`
			RegionEnName *string `json:"regionEnName" name:"regionEnName"`
		} `json:"RegionList" name:"RegionList"`
	} `json:"Data"`
}

func (r *DescribeValidRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeValidRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AllocateEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Type       *string `json:"Type,omitempty" name:"Type"`
}

func (r *AllocateEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AllocateEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *AllocateEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeallocateEipRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeallocateEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeallocateEipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeallocateEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeallocateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRegionsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRegionsResponse struct {
	*ksyunhttp.BaseResponse
	Data struct {
		Regions []struct {
			Code *string `json:"Code" name:"Code"`
			Name *string `json:"Name" name:"Name"`
			AvailabilityZones []struct {
				Code *string `json:"Code" name:"Code"`
				Name *string `json:"Name" name:"Name"`
			} `json:"AvailabilityZones"`
		} `json:"Regions" name:"Regions"`
	} `json:"Data"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateMongoDBShardInstanceRequest struct {
	*ksyunhttp.BaseRequest
	PayType          *string   `json:"PayType,omitempty" name:"PayType"`
	AvailabilityZone []*string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Name             *string   `json:"Name,omitempty" name:"Name"`
	DbVersion        *string   `json:"DbVersion,omitempty" name:"DbVersion"`
	Storage          *string   `json:"Storage,omitempty" name:"Storage"`
	Duration         *int      `json:"Duration,omitempty" name:"Duration"`
	IamProjectId     *string   `json:"IamProjectId,omitempty" name:"IamProjectId"`
	VpcId            *string   `json:"VpcId,omitempty" name:"VpcId"`
	VnetId           *string   `json:"VnetId,omitempty" name:"VnetId"`
	InstancePassword *string   `json:"InstancePassword,omitempty" name:"InstancePassword"`
	ShardClass       *string   `json:"ShardClass,omitempty" name:"ShardClass"`
	ShardNum         *int      `json:"ShardNum,omitempty" name:"ShardNum"`
	MongosNum        *int      `json:"MongosNum,omitempty" name:"MongosNum"`
	MongosClass      *string   `json:"MongosClass,omitempty" name:"MongosClass"`
}

func (r *CreateMongoDBShardInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMongoDBShardInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId *string `json:"UserId" name:"UserId"`
		Region *string `json:"Region" name:"Region"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Name   *string `json:"Name" name:"Name"`
	} `json:"MongoDBInstanceResult"`
}

func (r *CreateMongoDBShardInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMongoDBShardInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DownloadSnapshotRequest struct {
	*ksyunhttp.BaseRequest
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DownloadSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DownloadSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	SnapshotId *string `json:"SnapshotId" name:"SnapshotId"`
	Url        *string `json:"Url" name:"Url"`
}

func (r *DownloadSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CloneInstanceRequest struct {
	*ksyunhttp.BaseRequest
	PayType          *string `json:"PayType,omitempty" name:"PayType"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
	Name             *string `json:"Name,omitempty" name:"Name"`
	Duration         *int    `json:"Duration,omitempty" name:"Duration"`
	IamProjectId     *string `json:"IamProjectId,omitempty" name:"IamProjectId"`
	VpcId            *string `json:"VpcId,omitempty" name:"VpcId"`
	VnetId           *string `json:"VnetId,omitempty" name:"VnetId"`
	InstancePassword *string `json:"InstancePassword,omitempty" name:"InstancePassword"`
	SnapshotId       *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	InstanceId       *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloneInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CloneInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId *string `json:"UserId" name:"UserId"`
		Region *string `json:"Region" name:"Region"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Name   *string `json:"Name" name:"Name"`
	} `json:"MongoDBInstanceResult"`
}

func (r *CloneInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeShardNodeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeShardNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeShardNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	MongosNodeResult []struct {
		NodeId      *string `json:"NodeId" name:"NodeId"`
		Name        *string `json:"Name" name:"Name"`
		Role        *string `json:"Role" name:"Role"`
		Endpoint    *string `json:"Endpoint" name:"Endpoint"`
		Status      *string `json:"Status" name:"Status"`
		Connections *int    `json:"Connections" name:"Connections"`
		InstanceClass *string `json:"InstanceClass" name:"InstanceClass"`
		Ipv6Vip     *string `json:"Ipv6Vip" name:"Ipv6Vip"`
		EipEport    *string `json:"EipEport" name:"EipEport"`
	} `json:"MongosNodeResult"`
	ShardNodeResult []struct {
		NodeId   *string `json:"NodeId" name:"NodeId"`
		Name     *string `json:"Name" name:"Name"`
		Status   *string `json:"Status" name:"Status"`
		Disk     *int    `json:"Disk" name:"Disk"`
		Iops     *int    `json:"Iops" name:"Iops"`
		InstanceClass *string `json:"InstanceClass" name:"InstanceClass"`
		NodeNum  *int    `json:"NodeNum" name:"NodeNum"`
		UsedDisk *int    `json:"UsedDisk" name:"UsedDisk"`
	} `json:"ShardNodeResult"`
}

func (r *DescribeShardNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShardNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceStatisticRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeInstanceStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceStatisticResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total struct {
			Code  *string `json:"code" name:"code"`
			Name  *string `json:"name" name:"name"`
			Count *int    `json:"count" name:"count"`
			Items []struct {
				Code  *string `json:"Code" name:"Code"`
				Name  *string `json:"Name" name:"Name"`
				Count *int    `json:"Count" name:"Count"`
			} `json:"Items"`
		} `json:"Total" name:"Total"`
		Details []struct {
			Code  *string `json:"code" name:"code"`
			Name  *string `json:"name" name:"name"`
			Count *int    `json:"count" name:"count"`
			Items []struct {
				Code  *string `json:"Code" name:"Code"`
				Name  *string `json:"Name" name:"Name"`
				Count *int    `json:"Count" name:"Count"`
			} `json:"Items"`
		} `json:"Details" name:"Details"`
	} `json:"Data"`
}

func (r *DescribeInstanceStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type AddClusterNodeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeType    *string `json:"NodeType,omitempty" name:"NodeType"`
	NodeClass   *string `json:"NodeClass,omitempty" name:"NodeClass"`
	NodeStorage *int    `json:"NodeStorage,omitempty" name:"NodeStorage"`
}

func (r *AddClusterNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddClusterNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *AddClusterNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteClusterNodeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeId     *string `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteClusterNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteClusterNodeResponse struct {
	*ksyunhttp.BaseResponse
	Null *string `json:"null" name:"null"`
}

func (r *DeleteClusterNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSlowLogDetailRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeId       *string `json:"NodeId,omitempty" name:"NodeId"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	Database     *string `json:"Database,omitempty" name:"Database"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Marker       *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords   *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeSlowLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSlowLogDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		ProcessingTime *string `json:"ProcessingTime" name:"ProcessingTime"`
		NameSpace    *string `json:"NameSpace" name:"NameSpace"`
		Content      *string `json:"Content" name:"Content"`
		Client       *string `json:"Client" name:"Client"`
		User         *string `json:"User" name:"User"`
		DocsExamined *string `json:"DocsExamined" name:"DocsExamined"`
		KeysExamined *string `json:"KeysExamined" name:"KeysExamined"`
		KeysUpdates  *string `json:"KeysUpdates" name:"KeysUpdates"`
		Nreturned    *string `json:"Nreturned" name:"Nreturned"`
		ResponseLength *string `json:"ResponseLength" name:"ResponseLength"`
		Millis       *string `json:"Millis" name:"Millis"`
	} `json:"Data"`
	Marker     *int `json:"marker" name:"marker"`
	MaxRecords *int `json:"maxRecords" name:"maxRecords"`
	Total      *int `json:"Total" name:"Total"`
}

func (r *DescribeSlowLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSlowLogStatisticsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeId       *string `json:"NodeId,omitempty" name:"NodeId"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	Marker       *int    `json:"Marker,omitempty" name:"Marker"`
	MaxRecords   *int    `json:"MaxRecords,omitempty" name:"MaxRecords"`
}

func (r *DescribeSlowLogStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSlowLogStatisticsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Op              *string `json:"Op" name:"Op"`
		NameSpace       *string `json:"NameSpace" name:"NameSpace"`
		QueryCount      *int    `json:"QueryCount" name:"QueryCount"`
		MillisAvg       *int    `json:"MillisAvg" name:"MillisAvg"`
		MillisMax       *int    `json:"MillisMax" name:"MillisMax"`
		DocsExaminedAvg *int    `json:"DocsExaminedAvg" name:"DocsExaminedAvg"`
		KeysExaminedAvg *int    `json:"KeysExaminedAvg" name:"KeysExaminedAvg"`
		KeysUpdatesAvg  *int    `json:"KeysUpdatesAvg" name:"KeysUpdatesAvg"`
		NreturnedAvg    *int    `json:"NreturnedAvg" name:"NreturnedAvg"`
		SlowLogDetailVo *string `json:"SlowLogDetailVo" name:"SlowLogDetailVo"`
	} `json:"Data"`
	Offset *int `json:"Offset" name:"Offset"`
	Limit  *int `json:"Limit" name:"Limit"`
	Total  *int `json:"Total" name:"Total"`
}

func (r *DescribeSlowLogStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSlowLogDatabaseRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeId       *string `json:"NodeId,omitempty" name:"NodeId"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSlowLogDatabaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSlowLogDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	Data      []*string `json:"Data" name:"Data"`
}

func (r *DescribeSlowLogDatabaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSlowLogLineChartRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	NodeId       *string `json:"NodeId,omitempty" name:"NodeId"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSlowLogLineChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSlowLogLineChartResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		ExecTime *int `json:"ExecTime" name:"ExecTime"`
		Count *int `json:"Count" name:"Count"`
	} `json:"Data"`
}

func (r *DescribeSlowLogLineChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateMongoDBInstanceClusterRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	NodeType      *string `json:"NodeType,omitempty" name:"NodeType"`
	NodeId        *string `json:"NodeId,omitempty" name:"NodeId"`
	Storage       *int    `json:"Storage,omitempty" name:"Storage"`
}

func (r *UpdateMongoDBInstanceClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateMongoDBInstanceClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId *string `json:"UserId" name:"UserId"`
		Region *string `json:"Region" name:"Region"`
		InstanceId *string `json:"InstanceId" name:"InstanceId"`
		Name   *string `json:"Name" name:"Name"`
	} `json:"MongoDBInstanceResult"`
}

func (r *UpdateMongoDBInstanceClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMongoDBInstanceClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeClusterForRestoreRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId     *string `json:"InstanceId,omitempty" name:"InstanceId"`
	ResetTimePoint *string `json:"ResetTimePoint,omitempty" name:"ResetTimePoint"`
}

func (r *DescribeClusterForRestoreRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeClusterForRestoreResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId          *string   `json:"UserId" name:"UserId"`
		Region          *string   `json:"Region" name:"Region"`
		Name            *string   `json:"Name" name:"Name"`
		InstanceId      *string   `json:"InstanceId" name:"InstanceId"`
		Status          *string   `json:"Status" name:"Status"`
		IP              *string   `json:"IP" name:"IP"`
		InstanceType    *string   `json:"InstanceType" name:"InstanceType"`
		Version         *string   `json:"Version" name:"Version"`
		InstanceClass   *string   `json:"InstanceClass" name:"InstanceClass"`
		Storage         *int      `json:"Storage" name:"Storage"`
		SecurityGroupId *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
		Port            *int      `json:"Port" name:"Port"`
		NetworkType     *string   `json:"NetworkType" name:"NetworkType"`
		VpcId           *string   `json:"VpcId" name:"VpcId"`
		VnetId          *string   `json:"VnetId" name:"VnetId"`
		TimingSwitch    *string   `json:"TimingSwitch" name:"TimingSwitch"`
		Timezone        *string   `json:"Timezone" name:"Timezone"`
		TimeCycle       *string   `json:"TimeCycle" name:"TimeCycle"`
		ProductId       *string   `json:"ProductId" name:"ProductId"`
		PayType         *string   `json:"PayType" name:"PayType"`
		ProductWhat     *int      `json:"ProductWhat" name:"ProductWhat"`
		CreateDate      *string   `json:"CreateDate" name:"CreateDate"`
		ExpirationDate  *string   `json:"ExpirationDate" name:"ExpirationDate"`
		IamProjectId    *string   `json:"IamProjectId" name:"IamProjectId"`
		IamProjectName  *string   `json:"IamProjectName" name:"IamProjectName"`
		NodeNum         *int      `json:"NodeNum" name:"NodeNum"`
		MongosNum       *string   `json:"MongosNum" name:"MongosNum"`
		ShardNum        *string   `json:"ShardNum" name:"ShardNum"`
		Mode            *string   `json:"Mode" name:"Mode"`
		Config          *string   `json:"Config" name:"Config"`
		Area            *string   `json:"Area" name:"Area"`
		SlbaclId        *string   `json:"SlbaclId" name:"SlbaclId"`
		Ipv6Vip         *string   `json:"Ipv6Vip" name:"Ipv6Vip"`
		IpVersion       *string   `json:"IpVersion" name:"IpVersion"`
		Tags            []*string `json:"Tags" name:"Tags"`
		Shards          struct {
			Vcpu      *int `json:"vcpu" name:"vcpu"`
			MemSize   *int `json:"mem_size" name:"mem_size"`
			ShardsNum *int `json:"shards_num" name:"shards_num"`
			DiskSize  *int `json:"disk_size" name:"disk_size"`
		} `json:"Shards" name:"Shards"`
		Mongos struct {
			Vcpu      *int `json:"vcpu" name:"vcpu"`
			MemSize   *int `json:"mem_size" name:"mem_size"`
			MongosNum *int `json:"mongos_num" name:"mongos_num"`
			DiskSize  *int `json:"disk_size" name:"disk_size"`
		} `json:"Mongos" name:"Mongos"`
	} `json:"MongoDBInstanceResult"`
}

func (r *DescribeClusterForRestoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterForRestoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeDefaultParamsRequest struct {
	*ksyunhttp.BaseRequest
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDefaultParamsResponse struct {
	*ksyunhttp.BaseResponse
	DefaultParams []struct {
		DefaultValue    *string   `json:"DefaultValue" name:"DefaultValue"`
		Visible         *int      `json:"Visible" name:"Visible"`
		RestartRequired *bool     `json:"RestartRequired" name:"RestartRequired"`
		ParamName       *string   `json:"ParamName" name:"ParamName"`
		ParamType       *string   `json:"ParamType" name:"ParamType"`
		DataType        *string   `json:"DataType" name:"DataType"`
		Enums           []*string `json:"Enums" name:"Enums"`
	} `json:"DefaultParams"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateParamGroupRequest struct {
	*ksyunhttp.BaseRequest
	ParamGroupName *string  `json:"ParamGroupName,omitempty" name:"ParamGroupName"`
	Description    *string  `json:"Description,omitempty" name:"Description"`
	DbVersion      *float64 `json:"DbVersion,omitempty" name:"DbVersion"`
	Params         *string  `json:"Params,omitempty" name:"Params"`
}

func (r *CreateParamGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateParamGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ParamGroupId *string `json:"ParamGroupId" name:"ParamGroupId"`
	} `json:"Data"`
}

func (r *CreateParamGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateParamGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParamGroupListRequest struct {
	*ksyunhttp.BaseRequest
	DbVersion  *string `json:"DbVersion,omitempty" name:"DbVersion"`
	NameSearch *string `json:"NameSearch,omitempty" name:"NameSearch"`
	Offset     *string `json:"Offset,omitempty" name:"Offset"`
	Limit      *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeParamGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeParamGroupListResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount     *int `json:"TotalCount" name:"TotalCount"`
	ParamGroupList []struct {
		DbVersion      *string `json:"DbVersion" name:"DbVersion"`
		ParamGroupName *string `json:"ParamGroupName" name:"ParamGroupName"`
		Description    *string `json:"Description" name:"Description"`
		ParamGroupId   *string `json:"ParamGroupId" name:"ParamGroupId"`
	} `json:"ParamGroupList"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Offset    *int    `json:"Offset" name:"Offset"`
	Limit     *int    `json:"Limit" name:"Limit"`
}

func (r *DescribeParamGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeParamGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParamGroupInfoRequest struct {
	*ksyunhttp.BaseRequest
	ParamGroupId *string `json:"ParamGroupId,omitempty" name:"ParamGroupId"`
}

func (r *DescribeParamGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeParamGroupInfoResponse struct {
	*ksyunhttp.BaseResponse
	ParamGroupInfo struct {
		Params struct {
			SetParameterFailIndexKeyTooInt      *string `json:"setParameter.failIndexKeyTooInt" name:"setParameter.failIndexKeyTooInt"`
			OperationProfilingSlowOpThresholdMs *int    `json:"operationProfiling.slowOpThresholdMs" name:"operationProfiling.slowOpThresholdMs"`
			ReplicationOplogSizeMB              *string `json:"replication.oplogSizeMB" name:"replication.oplogSizeMB"`
			OperationProfilingMode              *string `json:"operationProfiling.mode" name:"operationProfiling.mode"`
			SetParameterCursorTimeoutMillis     *int    `json:"setParameter.cursorTimeoutMillis" name:"setParameter.cursorTimeoutMillis"`
		} `json:"Params" name:"Params"`
		ParamGroupName *string `json:"ParamGroupName" name:"ParamGroupName"`
		Description    *string `json:"Description" name:"Description"`
		ParamGroupId   *string `json:"ParamGroupId" name:"ParamGroupId"`
	} `json:"ParamGroupInfo"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeParamGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeParamGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModifyHistoryRequest struct {
	*ksyunhttp.BaseRequest
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	Offset    *int    `json:"Offset,omitempty" name:"Offset"`
	Limit     *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeModifyHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeModifyHistoryResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount  *int      `json:"TotalCount" name:"TotalCount"`
	HistoryInfo []*string `json:"HistoryInfo" name:"HistoryInfo"`
	RequestId   *string   `json:"RequestId" name:"RequestId"`
	Offset      *int      `json:"Offset" name:"Offset"`
	Limit       *int      `json:"Limit" name:"Limit"`
}

func (r *DescribeModifyHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModifyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceParamsResponse struct {
	*ksyunhttp.BaseResponse
	RuntimeParams struct {
		SetParameter struct {
			MaxTransactionLockRequestTimeoutMillis *int `json:"maxTransactionLockRequestTimeoutMillis" name:"maxTransactionLockRequestTimeoutMillis"`
			TransactionLifetimeLimitSeconds        *int `json:"transactionLifetimeLimitSeconds" name:"transactionLifetimeLimitSeconds"`
			CursorTimeoutMillis                    *int `json:"cursorTimeoutMillis" name:"cursorTimeoutMillis"`
		} `json:"SetParameter" name:"SetParameter"`
		OperationProfiling struct {
			SlowOpThresholdMs *int    `json:"slowOpThresholdMs" name:"slowOpThresholdMs"`
			Mode              *string `json:"mode" name:"mode"`
		} `json:"OperationProfiling" name:"OperationProfiling"`
		Replication struct {
			OplogSizeMB *string `json:"oplogSizeMB" name:"oplogSizeMB"`
		} `json:"Replication" name:"Replication"`
	} `json:"RuntimeParams"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyParamGroupRequest struct {
	*ksyunhttp.BaseRequest
	NewParamGroupName *string `json:"NewParamGroupName,omitempty" name:"NewParamGroupName"`
	NewDescription    *string `json:"NewDescription,omitempty" name:"NewDescription"`
}

func (r *ModifyParamGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyParamGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *ModifyParamGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyParamGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteParamGroupRequest struct {
	*ksyunhttp.BaseRequest
	ParamGroupId *string `json:"ParamGroupId,omitempty" name:"ParamGroupId"`
}

func (r *DeleteParamGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteParamGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *DeleteParamGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteParamGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

