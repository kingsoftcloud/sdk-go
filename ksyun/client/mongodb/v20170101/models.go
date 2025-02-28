package v20170101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
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

func (r *CreateMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId     *string `json:"UserId"`
		Region     *string `json:"Region"`
		InstanceId *string `json:"InstanceId"`
		Name       *string `json:"Name"`
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

func (r *DeleteMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId         *string `json:"UserId"`
		Region         *string `json:"Region"`
		Name           *string `json:"Name"`
		InstanceId     *string `json:"InstanceId"`
		Status         *string `json:"Status"`
		IP             *string `json:"IP"`
		InstanceType   *string `json:"InstanceType"`
		Version        *string `json:"Version"`
		InstanceClass  *string `json:"InstanceClass"`
		Storage        *int    `json:"Storage"`
		Port           *int    `json:"Port"`
		NetworkType    *string `json:"NetworkType"`
		VpcId          *string `json:"VpcId"`
		VnetId         *string `json:"VnetId"`
		TimingSwitch   *string `json:"TimingSwitch"`
		Timezone       *string `json:"Timezone"`
		TimeCycle      *string `json:"TimeCycle"`
		ProductId      *string `json:"ProductId"`
		ProductWhat    *int    `json:"ProductWhat"`
		PayType        *string `json:"PayType"`
		CreateDate     *string `json:"CreateDate"`
		ExpirationDate *string `json:"ExpirationDate"`
		IamProjectId   *string `json:"IamProjectId"`
		IamProjectName *string `json:"IamProjectName"`
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

func (r *DescribeMongoDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMongoDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMongoDBInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	MongoDBInstancesResult []struct {
		IamProjectId   *string `json:"IamProjectId"`
		IamProjectName *string `json:"IamProjectName"`
		UserId         *string `json:"UserId"`
		Region         *string `json:"Region"`
		Name           *string `json:"Name"`
		InstanceId     *string `json:"InstanceId"`
		Status         *string `json:"Status"`
		IP             *string `json:"IP"`
		InstanceType   *string `json:"InstanceType"`
		Version        *string `json:"Version"`
		InstanceClass  *string `json:"InstanceClass"`
		Storage        *int    `json:"Storage"`
		Port           *string `json:"Port"`
		NetworkType    *string `json:"NetworkType"`
		VpcId          *string `json:"VpcId"`
		VnetId         *string `json:"VnetId"`
		ProductId      *string `json:"ProductId"`
		ProductWhat    *int    `json:"ProductWhat"`
		PayType        *string `json:"PayType"`
		CreateDate     *string `json:"CreateDate"`
		ExpirationDate *string `json:"ExpirationDate"`
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

func (r *DescribeMongoDBInstanceNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMongoDBInstanceNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMongoDBInstanceNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceNodeResult []struct {
		NodeId *string `json:"NodeId"`
		Name   *string `json:"Name"`
		Role   *string `json:"Role"`
		IP     *string `json:"IP"`
		Port   *int    `json:"Port"`
		Status *string `json:"Status"`
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

func (r *RenameMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenameMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	name       *string `json:"name" name:"name"`
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

func (r *ResetPasswordMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ResetPasswordMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *RestartMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RestartMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CreateMongoDBSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateMongoDBSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMongoDBSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBSnapshotResult struct {
		SnapshotId *string `json:"SnapshotId"`
		Name       *string `json:"Name"`
		InstanceId *string `json:"InstanceId"`
		Status     *string `json:"Status"`
		Create     *string `json:"Create"`
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

func (r *SetMongoDBTimingSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "SetMongoDBTimingSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeMongoDBSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMongoDBSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMongoDBSnapshotResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBSnapshotResult []struct {
		SnapshotId *string `json:"SnapshotId"`
		Name       *string `json:"Name"`
		InstanceId *string `json:"InstanceId"`
		Type       *string `json:"Type"`
		Size       *int    `json:"Size"`
		Status     *string `json:"Status"`
		BackupMode *string `json:"BackupMode"`
		CreateTime *string `json:"CreateTime"`
		UpdateTime *string `json:"UpdateTime"`
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

func (r *DeleteMongoDBSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteMongoDBSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *RenameMongoDBSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RenameMongoDBSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
	RequestId                *string `json:"RequestId" name:"RequestId"`
	MongoDBSecurityGroupRule []struct {
		To_port   *string `json:"To_port"`
		Cidr      *string `json:"Cidr"`
		From_port *string `json:"From_port"`
		Protocol  *string `json:"Protocol"`
		Id        *string `json:"Id"`
	} `json:"MongoDBSecurityGroupRule"`
}

func (r *AddSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSecurityGroupRuleResponse) FromJsonString(s string) error {
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

func (r *ListSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSecurityGroupRulesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	MongoDBSecurityGroupRule []struct {
		To_port   *string `json:"To_port"`
		Cidr      *string `json:"Cidr"`
		From_port *string `json:"From_port"`
		Protocol  *string `json:"Protocol"`
		Id        *string `json:"Id"`
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

func (r *UpdateMongoDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateMongoDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMongoDBInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId     *string `json:"UserId"`
		Region     *string `json:"Region"`
		InstanceId *string `json:"InstanceId"`
		Name       *string `json:"Name"`
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

func (r *AddSecondaryInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddSecondaryInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeMongoDBShardNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMongoDBShardNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMongoDBShardNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	MongosNodeResult []struct {
		NodeId        *string `json:"NodeId"`
		Name          *string `json:"Name"`
		Role          *string `json:"Role"`
		Endpoint      *string `json:"Endpoint"`
		Status        *string `json:"Status"`
		Connections   *int    `json:"Connections"`
		InstanceClass *string `json:"InstanceClass"`
	} `json:"MongosNodeResult"`
	ShardNodeResult []struct {
		NodeId        *string `json:"NodeId"`
		Name          *string `json:"Name"`
		Status        *string `json:"Status"`
		Disk          *int    `json:"Disk"`
		Iops          *int    `json:"Iops"`
		InstanceClass *string `json:"InstanceClass"`
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
	Data      struct {
		RegionList []struct {
			areaEnName *string `json:"areaEnName"`
			regionCode *string `json:"regionCode"`
			areaCode   *string `json:"areaCode"`
			innerCode  *string `json:"innerCode"`
			regionType *int    `json:"regionType"`
			areaName   *string `json:"areaName"`
			AzList     []struct {
				AzName *string `json:"AzName"`
				AzCode *string `json:"AzCode"`
			} `json:"AzList"`
			regionName   *string `json:"regionName"`
			regionEnName *string `json:"regionEnName"`
		} `json:"RegionList"`
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

func (r *AllocateEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AllocateEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeallocateEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeallocateEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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
	Data struct {
		Regions []struct {
			Code              *string `json:"Code"`
			Name              *string `json:"Name"`
			AvailabilityZones []struct {
				Code *string `json:"Code"`
				Name *string `json:"Name"`
			} `json:"AvailabilityZones"`
		} `json:"Regions"`
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

func (r *CreateMongoDBShardInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateMongoDBShardInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMongoDBShardInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId     *string `json:"UserId"`
		Region     *string `json:"Region"`
		InstanceId *string `json:"InstanceId"`
		Name       *string `json:"Name"`
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
}

func (r *DownloadSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DownloadSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *CloneInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId     *string `json:"UserId"`
		Region     *string `json:"Region"`
		InstanceId *string `json:"InstanceId"`
		Name       *string `json:"Name"`
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

func (r *DescribeShardNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeShardNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShardNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	MongosNodeResult []struct {
		NodeId        *string `json:"NodeId"`
		Name          *string `json:"Name"`
		Role          *string `json:"Role"`
		Endpoint      *string `json:"Endpoint"`
		Status        *string `json:"Status"`
		Connections   *int    `json:"Connections"`
		InstanceClass *string `json:"InstanceClass"`
		Ipv6Vip       *string `json:"Ipv6Vip"`
		EipEport      *string `json:"EipEport"`
	} `json:"MongosNodeResult"`
	ShardNodeResult []struct {
		NodeId        *string `json:"NodeId"`
		Name          *string `json:"Name"`
		Status        *string `json:"Status"`
		Disk          *int    `json:"Disk"`
		Iops          *int    `json:"Iops"`
		InstanceClass *string `json:"InstanceClass"`
		NodeNum       *int    `json:"NodeNum"`
		UsedDisk      *int    `json:"UsedDisk"`
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
	Action *string `json:"Action,omitempty" name:"Action"`
}

func (r *DescribeInstanceStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatisticResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total struct {
			code  *string `json:"code"`
			name  *string `json:"name"`
			count *int    `json:"count"`
			Items []struct {
				Code  *string `json:"Code"`
				Name  *string `json:"Name"`
				Count *int    `json:"Count"`
			} `json:"Items"`
		} `json:"Total"`
		Details []struct {
			code  *string `json:"code"`
			name  *string `json:"name"`
			count *int    `json:"count"`
			Items []struct {
				Code  *string `json:"Code"`
				Name  *string `json:"Name"`
				Count *int    `json:"Count"`
			} `json:"Items"`
		} `json:"Details"`
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

func (r *AddClusterNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddClusterNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteClusterNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteClusterNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterNodeResponse struct {
	*ksyunhttp.BaseResponse
	null *string `json:"null" name:"null"`
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

func (r *DescribeSlowLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSlowLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		ProcessingTime *string `json:"ProcessingTime"`
		NameSpace      *string `json:"NameSpace"`
		Content        *string `json:"Content"`
		Client         *string `json:"Client"`
		User           *string `json:"User"`
		DocsExamined   *string `json:"DocsExamined"`
		KeysExamined   *string `json:"KeysExamined"`
		KeysUpdates    *string `json:"KeysUpdates"`
		Nreturned      *string `json:"Nreturned"`
		ResponseLength *string `json:"ResponseLength"`
		Millis         *string `json:"Millis"`
	} `json:"Data"`
	marker     *int `json:"marker" name:"marker"`
	maxRecords *int `json:"maxRecords" name:"maxRecords"`
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

func (r *DescribeSlowLogStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSlowLogStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogStatisticsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Op              *string `json:"Op"`
		NameSpace       *string `json:"NameSpace"`
		QueryCount      *int    `json:"QueryCount"`
		MillisAvg       *int    `json:"MillisAvg"`
		MillisMax       *int    `json:"MillisMax"`
		DocsExaminedAvg *int    `json:"DocsExaminedAvg"`
		KeysExaminedAvg *int    `json:"KeysExaminedAvg"`
		KeysUpdatesAvg  *int    `json:"KeysUpdatesAvg"`
		NreturnedAvg    *int    `json:"NreturnedAvg"`
		SlowLogDetailVo *string `json:"SlowLogDetailVo"`
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

func (r *DescribeSlowLogDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSlowLogDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDatabaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string   `json:"RequestId" name:"RequestId"`
	data      []*string `json:"data" name:"data"`
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

func (r *DescribeSlowLogLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSlowLogLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogLineChartResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		ExecTime *int `json:"ExecTime"`
		Count    *int `json:"Count"`
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

func (r *UpdateMongoDBInstanceClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateMongoDBInstanceClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMongoDBInstanceClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId     *string `json:"UserId"`
		Region     *string `json:"Region"`
		InstanceId *string `json:"InstanceId"`
		Name       *string `json:"Name"`
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

func (r *DescribeClusterForRestoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterForRestoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterForRestoreResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	MongoDBInstanceResult struct {
		UserId          *string `json:"UserId"`
		Region          *string `json:"Region"`
		Name            *string `json:"Name"`
		InstanceId      *string `json:"InstanceId"`
		Status          *string `json:"Status"`
		IP              *string `json:"IP"`
		InstanceType    *string `json:"InstanceType"`
		Version         *string `json:"Version"`
		InstanceClass   *string `json:"InstanceClass"`
		Storage         *int    `json:"Storage"`
		SecurityGroupId *string `json:"SecurityGroupId"`
		Port            *int    `json:"Port"`
		NetworkType     *string `json:"NetworkType"`
		VpcId           *string `json:"VpcId"`
		VnetId          *string `json:"VnetId"`
		TimingSwitch    *string `json:"TimingSwitch"`
		Timezone        *string `json:"Timezone"`
		TimeCycle       *string `json:"TimeCycle"`
		ProductId       *string `json:"ProductId"`
		PayType         *string `json:"PayType"`
		ProductWhat     *int    `json:"ProductWhat"`
		CreateDate      *string `json:"CreateDate"`
		ExpirationDate  *string `json:"ExpirationDate"`
		IamProjectId    *string `json:"IamProjectId"`
		IamProjectName  *string `json:"IamProjectName"`
		NodeNum         *int    `json:"NodeNum"`
		MongosNum       *string `json:"MongosNum"`
		ShardNum        *string `json:"ShardNum"`
		Mode            *string `json:"Mode"`
		Config          *string `json:"Config"`
		Area            *string `json:"Area"`
		SlbaclId        *string `json:"SlbaclId"`
		Ipv6Vip         *string `json:"Ipv6Vip"`
		IpVersion       *string `json:"IpVersion"`
		Tags            []struct {
		} `json:"Tags"`
		Shards struct {
			vcpu       *int `json:"vcpu"`
			mem_size   *int `json:"mem_size"`
			shards_num *int `json:"shards_num"`
			disk_size  *int `json:"disk_size"`
		} `json:"Shards"`
		Mongos struct {
			vcpu       *int `json:"vcpu"`
			mem_size   *int `json:"mem_size"`
			mongos_num *int `json:"mongos_num"`
			disk_size  *int `json:"disk_size"`
		} `json:"Mongos"`
	} `json:"MongoDBInstanceResult"`
}

func (r *DescribeClusterForRestoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterForRestoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
