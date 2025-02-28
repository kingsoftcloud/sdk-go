package v20160304

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeCenInstancesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeCenGrantsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeRegionGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeCenBandWidthPackagesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeCenRegionBandwidthsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeCenRoutesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeCenRegionsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateCenRequest struct {
	*ksyunhttp.BaseRequest
	CenName *string `json:"CenName,omitempty" name:"CenName"`
}

func (r *CreateCenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateCenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCenRequest struct {
	*ksyunhttp.BaseRequest
	CenId   *string `json:"CenId,omitempty" name:"CenId"`
	CenName *string `json:"CenName,omitempty" name:"CenName"`
}

func (r *ModifyCenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyCenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenRequest struct {
	*ksyunhttp.BaseRequest
	CenId *string `json:"CenId,omitempty" name:"CenId"`
}

func (r *DeleteCenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteCenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCensRequest struct {
	*ksyunhttp.BaseRequest
	CenId      []*string `json:"CenId,omitempty" name:"CenId"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCensRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCensRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCensRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCensResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	CenSet    []struct {
		CreateTime *string `json:"CreateTime"`
		CenId      *string `json:"CenId"`
		CenName    *string `json:"CenName"`
	} `json:"CenSet"`
}

func (r *DescribeCensResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCensResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachCenInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CenId             *string `json:"CenId,omitempty" name:"CenId"`
	InstanceType      *string `json:"InstanceType,omitempty" name:"InstanceType"`
	CenRegion         *string `json:"CenRegion,omitempty" name:"CenRegion"`
	CenInstanceId     *string `json:"CenInstanceId,omitempty" name:"CenInstanceId"`
	InstanceAccountId *string `json:"InstanceAccountId,omitempty" name:"InstanceAccountId"`
}

func (r *AttachCenInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCenInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachCenInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachCenInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachCenInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCenInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCenInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CenId         *string `json:"CenId,omitempty" name:"CenId"`
	CenInstanceId *string `json:"CenInstanceId,omitempty" name:"CenInstanceId"`
}

func (r *DetachCenInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCenInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachCenInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachCenInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachCenInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCenInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenInstancesRequest struct {
	*ksyunhttp.BaseRequest
	CenInstanceId []*string                     `json:"CenInstanceId,omitempty" name:"CenInstanceId"`
	MaxResults    *int                          `json:"MaxResults,omitempty" name:"MaxResults"`
	Filter        []*DescribeCenInstancesFilter `json:"Filter,omitempty" name:"Filter"`
	NextToken     *string                       `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCenInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	NextToken      *string `json:"NextToken" name:"NextToken"`
	CenInstanceSet []struct {
		CenId             *string `json:"CenId"`
		CenInstanceId     *string `json:"CenInstanceId"`
		InstanceType      *string `json:"InstanceType"`
		CenRegion         *string `json:"CenRegion"`
		InstanceAccountId *string `json:"InstanceAccountId"`
	} `json:"CenInstanceSet"`
}

func (r *DescribeCenInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatCenGrantRequest struct {
	*ksyunhttp.BaseRequest
	CenId         *string `json:"CenId,omitempty" name:"CenId"`
	InstanceType  *string `json:"InstanceType,omitempty" name:"InstanceType"`
	CenInstanceId *string `json:"CenInstanceId,omitempty" name:"CenInstanceId"`
	CenAccountId  *string `json:"CenAccountId,omitempty" name:"CenAccountId"`
}

func (r *CreatCenGrantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatCenGrantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreatCenGrantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatCenGrantResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreatCenGrantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatCenGrantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenGrantRequest struct {
	*ksyunhttp.BaseRequest
	CenGrantId *string `json:"CenGrantId,omitempty" name:"CenGrantId"`
}

func (r *DeleteCenGrantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenGrantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCenGrantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenGrantResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteCenGrantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenGrantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenGrantsRequest struct {
	*ksyunhttp.BaseRequest
	CenGrantId []*string                  `json:"CenGrantId,omitempty" name:"CenGrantId"`
	Filter     []*DescribeCenGrantsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                    `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCenGrantsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenGrantsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenGrantsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenGrantsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	NextToken   *string `json:"NextToken" name:"NextToken"`
	CenGrantSet []struct {
		CreateTime    *string `json:"CreateTime"`
		CenGrantId    *string `json:"CenGrantId"`
		CenId         *string `json:"CenId"`
		CenInstanceId *string `json:"CenInstanceId"`
		InstanceType  *string `json:"InstanceType"`
		CenAccountId  *string `json:"CenAccountId"`
	} `json:"CenGrantSet"`
}

func (r *DescribeCenGrantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenGrantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionGroupsRequest struct {
	*ksyunhttp.BaseRequest
	RegionGroupId []*string                     `json:"RegionGroupId,omitempty" name:"RegionGroupId"`
	Filter        []*DescribeRegionGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults    *int                          `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string                       `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeRegionGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRegionGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	NextToken      *string `json:"NextToken" name:"NextToken"`
	RegionGroupSet []struct {
		RegionGroupId   *string `json:"RegionGroupId"`
		RegionGroupName *string `json:"RegionGroupName"`
	} `json:"RegionGroupSet"`
}

func (r *DescribeRegionGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenBandWidthPackageRequest struct {
	*ksyunhttp.BaseRequest
	CenId          *string `json:"CenId,omitempty" name:"CenId"`
	Name           *string `json:"Name,omitempty" name:"Name"`
	RegionAGroupId *string `json:"RegionAGroupId,omitempty" name:"RegionAGroupId"`
	RegionBGroupId *string `json:"RegionBGroupId,omitempty" name:"RegionBGroupId"`
	BandWidth      *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	ProjectId      *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType     *string `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime   *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *CreateCenBandWidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenBandWidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCenBandWidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenBandWidthPackageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateCenBandWidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenBandWidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCenBandWidthPackageRequest struct {
	*ksyunhttp.BaseRequest
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	BandWidth             *int    `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *ModifyCenBandWidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCenBandWidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCenBandWidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCenBandWidthPackageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyCenBandWidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCenBandWidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenBandWidthPackageRequest struct {
	*ksyunhttp.BaseRequest
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
}

func (r *DeleteCenBandWidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenBandWidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCenBandWidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenBandWidthPackageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteCenBandWidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenBandWidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachCenBandWidthPackageRequest struct {
	*ksyunhttp.BaseRequest
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	CenId                 *string `json:"CenId,omitempty" name:"CenId"`
}

func (r *AttachCenBandWidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCenBandWidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachCenBandWidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachCenBandWidthPackageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachCenBandWidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCenBandWidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCenBandWidthPackageRequest struct {
	*ksyunhttp.BaseRequest
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	CenId                 *string `json:"CenId,omitempty" name:"CenId"`
}

func (r *DetachCenBandWidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCenBandWidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachCenBandWidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachCenBandWidthPackageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachCenBandWidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCenBandWidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenBandWidthPackagesRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId             []*string                             `json:"ProjectId,omitempty" name:"ProjectId"`
	CenBandWidthPackageId []*string                             `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	Filter                []*DescribeCenBandWidthPackagesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults            *int                                  `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken             *string                               `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCenBandWidthPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenBandWidthPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenBandWidthPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenBandWidthPackagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	NextToken              *string `json:"NextToken" name:"NextToken"`
	CenBandWidthPackageSet []struct {
		CreateTime            *string `json:"CreateTime"`
		CenBandWidthPackageId *string `json:"CenBandWidthPackageId"`
		Name                  *string `json:"Name"`
		CenId                 *string `json:"CenId"`
		ProjectId             *string `json:"ProjectId"`
		BandWidth             *int    `json:"BandWidth"`
		RegionAGroupId        *string `json:"RegionAGroupId"`
		RegionBGroupId        *string `json:"RegionBGroupId"`
		BillType              *int    `json:"BillType"`
		ProductWhat           *int    `json:"ProductWhat"`
		ServiceEndTime        *string `json:"ServiceEndTime"`
	} `json:"CenBandWidthPackageSet"`
}

func (r *DescribeCenBandWidthPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenBandWidthPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenRegionBandwidthRequest struct {
	*ksyunhttp.BaseRequest
	CenId                 *string `json:"CenId,omitempty" name:"CenId"`
	RegionA               *string `json:"RegionA,omitempty" name:"RegionA"`
	RegionB               *string `json:"RegionB,omitempty" name:"RegionB"`
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	BandWidth             *int    `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *CreateCenRegionBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenRegionBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCenRegionBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenRegionBandwidthResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateCenRegionBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenRegionBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenRegionBandwidthRequest struct {
	*ksyunhttp.BaseRequest
	CenRegionBandwidthId *string `json:"CenRegionBandwidthId,omitempty" name:"CenRegionBandwidthId"`
}

func (r *DeleteCenRegionBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenRegionBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCenRegionBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCenRegionBandwidthResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteCenRegionBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCenRegionBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCenRegionBandwidthRequest struct {
	*ksyunhttp.BaseRequest
	CenRegionBandwidthId *string `json:"CenRegionBandwidthId,omitempty" name:"CenRegionBandwidthId"`
	BandWidth            *string `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *ModifyCenRegionBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCenRegionBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCenRegionBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCenRegionBandwidthResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyCenRegionBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCenRegionBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenRegionBandwidthsRequest struct {
	*ksyunhttp.BaseRequest
	CenRegionBandwidthId []*string                            `json:"CenRegionBandwidthId,omitempty" name:"CenRegionBandwidthId"`
	Filter               []*DescribeCenRegionBandwidthsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults           *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken            *string                              `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCenRegionBandwidthsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRegionBandwidthsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenRegionBandwidthsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenRegionBandwidthsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	NextToken             *string `json:"NextToken" name:"NextToken"`
	CenRegionBandwidthSet []struct {
		CenBandWidthPackageId *string `json:"CenBandWidthPackageId"`
		CenId                 *string `json:"CenId"`
		RegionA               *string `json:"RegionA"`
		RegionB               *string `json:"RegionB"`
		BandWidth             *int    `json:"BandWidth"`
		CenRegionBandwidthId  *string `json:"CenRegionBandwidthId"`
	} `json:"CenRegionBandwidthSet"`
}

func (r *DescribeCenRegionBandwidthsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRegionBandwidthsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenRoutesRequest struct {
	*ksyunhttp.BaseRequest
	CenRouteId []*string                  `json:"CenRouteId,omitempty" name:"CenRouteId"`
	Filter     []*DescribeCenRoutesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                       `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                    `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCenRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenRoutesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	NextToken   *string `json:"NextToken" name:"NextToken"`
	CenRouteSet []struct {
		CreateTime            *string `json:"CreateTime"`
		CenRouteId            *string `json:"CenRouteId"`
		CenId                 *string `json:"CenId"`
		DestinationCidrBlock  *string `json:"DestinationCidrBlock"`
		RouteStatus           *string `json:"RouteStatus"`
		NextHopInstanceId     *string `json:"NextHopInstanceId"`
		NextHopInstanceType   *string `json:"NextHopInstanceType"`
		NextHopInstanceRegion *string `json:"NextHopInstanceRegion"`
	} `json:"CenRouteSet"`
}

func (r *DescribeCenRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenRegionsRequest struct {
	*ksyunhttp.BaseRequest
	CenRegionId []*string                   `json:"CenRegionId,omitempty" name:"CenRegionId"`
	Filter      []*DescribeCenRegionsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults  *int                        `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken   *string                     `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCenRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenRegionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	NextToken    *string `json:"NextToken" name:"NextToken"`
	CenRegionSet []struct {
		RegionGroupId *string `json:"RegionGroupId"`
		CenRegionId   *string `json:"CenRegionId"`
		CenRegionName *string `json:"CenRegionName"`
	} `json:"CenRegionSet"`
}

func (r *DescribeCenRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenBandWidthPackageUsageRequest struct {
	*ksyunhttp.BaseRequest
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
}

func (r *DescribeCenBandWidthPackageUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenBandWidthPackageUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCenBandWidthPackageUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenBandWidthPackageUsageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeCenBandWidthPackageUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenBandWidthPackageUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
