package v20160304

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeCenGrantsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeCenBandWidthPackagesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeCenBandWidthPackagesTagKV struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type DescribeCenRegionBandwidthsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeCenRoutesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeNetworkInstancesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeInterAreasFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeInterRegionsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateCenRequest struct {
	*ksyunhttp.BaseRequest
	CenName     *string `json:"CenName,omitempty" name:"CenName"`
	Description *string `json:"Description,omitempty" name:"Description"`
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
	Cen       struct {
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		CenId       *string `json:"CenId" name:"CenId"`
		CenName     *string `json:"CenName" name:"CenName"`
		Description *string `json:"Description" name:"Description"`
	} `json:"Cen"`
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
	CenId       *string `json:"CenId,omitempty" name:"CenId"`
	CenName     *string `json:"CenName,omitempty" name:"CenName"`
	Description *string `json:"Description,omitempty" name:"Description"`
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
	Cen       struct {
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		CenId       *string `json:"CenId" name:"CenId"`
		CenName     *string `json:"CenName" name:"CenName"`
		Description *string `json:"Description" name:"Description"`
	} `json:"Cen"`
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
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
		CenId       *string `json:"CenId" name:"CenId"`
		CenName     *string `json:"CenName" name:"CenName"`
		Description *string `json:"Description" name:"Description"`
	} `json:"CenSet"`
}

func (r *DescribeCensResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCensResponse) FromJsonString(s string) error {
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
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		CenGrantId        *string `json:"CenGrantId" name:"CenGrantId"`
		CenId             *string `json:"CenId" name:"CenId"`
		NetworkInstanceId *string `json:"NetworkInstanceId" name:"NetworkInstanceId"`
		InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		CenAccountId      *string `json:"CenAccountId" name:"CenAccountId"`
	} `json:"CenGrantSet"`
}

func (r *DescribeCenGrantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenGrantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenBandWidthPackageRequest struct {
	*ksyunhttp.BaseRequest
	CenId                   *string `json:"CenId,omitempty" name:"CenId"`
	CenBandWidthPackageName *string `json:"CenBandWidthPackageName,omitempty" name:"CenBandWidthPackageName"`
	LocalAreaId             *string `json:"LocalAreaId,omitempty" name:"LocalAreaId"`
	RemoteAreaId            *string `json:"RemoteAreaId,omitempty" name:"RemoteAreaId"`
	PackageBandWidth        *int    `json:"PackageBandWidth,omitempty" name:"PackageBandWidth"`
	ProjectId               *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType              *string `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime            *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
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
	RequestId           *string `json:"RequestId" name:"RequestId"`
	CenBandWidthPackage struct {
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		CenBandWidthPackageId   *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
		CenBandWidthPackageName *string `json:"CenBandWidthPackageName" name:"CenBandWidthPackageName"`
		CenId                   *string `json:"CenId" name:"CenId"`
		ProjectId               *string `json:"ProjectId" name:"ProjectId"`
		PackageBandWidth        *int    `json:"PackageBandWidth" name:"PackageBandWidth"`
		LocalAreaId             *string `json:"LocalAreaId" name:"LocalAreaId"`
		RemoteAreaId            *string `json:"RemoteAreaId" name:"RemoteAreaId"`
	} `json:"CenBandWidthPackage"`
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
	CenBandWidthPackageId   *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	PackageBandWidth        *int    `json:"PackageBandWidth,omitempty" name:"PackageBandWidth"`
	CenBandWidthPackageName *string `json:"CenBandWidthPackageName,omitempty" name:"CenBandWidthPackageName"`
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
	RequestId           *string `json:"RequestId" name:"RequestId"`
	CenBandWidthPackage struct {
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		CenBandWidthPackageId   *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
		CenBandWidthPackageName *string `json:"CenBandWidthPackageName" name:"CenBandWidthPackageName"`
		CenId                   *string `json:"CenId" name:"CenId"`
		ProjectId               *string `json:"ProjectId" name:"ProjectId"`
		PackageBandWidth        *int    `json:"PackageBandWidth" name:"PackageBandWidth"`
		LocalAreaId             *string `json:"LocalAreaId" name:"LocalAreaId"`
		RemoteAreaId            *string `json:"RemoteAreaId" name:"RemoteAreaId"`
	} `json:"CenBandWidthPackage"`
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
	TagKey                []*string                             `json:"TagKey,omitempty" name:"TagKey"`
	TagKV                 []*DescribeCenBandWidthPackagesTagKV  `json:"TagKV,omitempty" name:"TagKV"`
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
		CreateTime              *string `json:"CreateTime" name:"CreateTime"`
		CenBandWidthPackageId   *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
		CenBandWidthPackageName *string `json:"CenBandWidthPackageName" name:"CenBandWidthPackageName"`
		CenId                   *string `json:"CenId" name:"CenId"`
		ProjectId               *string `json:"ProjectId" name:"ProjectId"`
		PackageBandWidth        *int    `json:"PackageBandWidth" name:"PackageBandWidth"`
		LocalAreaId             *string `json:"LocalAreaId" name:"LocalAreaId"`
		RemoteAreaId            *string `json:"RemoteAreaId" name:"RemoteAreaId"`
		BillType                *int    `json:"BillType" name:"BillType"`
		ProductWhat             *int    `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime          *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		TagSet                  []struct {
			TagId        *int    `json:"TagId" name:"TagId"`
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
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
	LocalRegion           *string `json:"LocalRegion,omitempty" name:"LocalRegion"`
	RemoteRegion          *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`
	CenBandWidthPackageId *string `json:"CenBandWidthPackageId,omitempty" name:"CenBandWidthPackageId"`
	InterBandWidth        *int    `json:"InterBandWidth,omitempty" name:"InterBandWidth"`
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
	RequestId          *string `json:"RequestId" name:"RequestId"`
	CenRegionBandwidth struct {
		CenBandWidthPackageId *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
		CenId                 *string `json:"CenId" name:"CenId"`
		LocalRegion           *string `json:"LocalRegion" name:"LocalRegion"`
		RemoteRegion          *string `json:"RemoteRegion" name:"RemoteRegion"`
		InterBandWidth        *int    `json:"InterBandWidth" name:"InterBandWidth"`
		CenRegionBandwidthId  *string `json:"CenRegionBandwidthId" name:"CenRegionBandwidthId"`
	} `json:"CenRegionBandwidth"`
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
	InterBandWidth       *string `json:"InterBandWidth,omitempty" name:"InterBandWidth"`
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
	RequestId          *string `json:"RequestId" name:"RequestId"`
	CenRegionBandwidth struct {
		CenBandWidthPackageId *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
		CenId                 *string `json:"CenId" name:"CenId"`
		LocalRegion           *string `json:"LocalRegion" name:"LocalRegion"`
		RemoteRegion          *string `json:"RemoteRegion" name:"RemoteRegion"`
		InterBandWidth        *int    `json:"InterBandWidth" name:"InterBandWidth"`
		CenRegionBandwidthId  *string `json:"CenRegionBandwidthId" name:"CenRegionBandwidthId"`
	} `json:"CenRegionBandwidth"`
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
		CenBandWidthPackageId *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
		CenId                 *string `json:"CenId" name:"CenId"`
		LocalRegion           *string `json:"LocalRegion" name:"LocalRegion"`
		RemoteRegion          *string `json:"RemoteRegion" name:"RemoteRegion"`
		InterBandWidth        *int    `json:"InterBandWidth" name:"InterBandWidth"`
		CenRegionBandwidthId  *string `json:"CenRegionBandwidthId" name:"CenRegionBandwidthId"`
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
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		CenRouteId           *string `json:"CenRouteId" name:"CenRouteId"`
		CenId                *string `json:"CenId" name:"CenId"`
		DestinationCidrBlock *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
		NetworkInstanceId    *string `json:"NetworkInstanceId" name:"NetworkInstanceId"`
		InstanceType         *string `json:"InstanceType" name:"InstanceType"`
		InstanceRegion       *string `json:"InstanceRegion" name:"InstanceRegion"`
		InstanceAccountId    *string `json:"InstanceAccountId" name:"InstanceAccountId"`
		NetworkRouteId       *string `json:"NetworkRouteId" name:"NetworkRouteId"`
		SelfRouteId          *string `json:"SelfRouteId" name:"SelfRouteId"`
		InstanceRouteType    *string `json:"InstanceRouteType" name:"InstanceRouteType"`
	} `json:"CenRouteSet"`
}

func (r *DescribeCenRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenRoutesResponse) FromJsonString(s string) error {
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
	RequestId                *string `json:"RequestId" name:"RequestId"`
	CenBandWidthPackageUsage struct {
		PackageUsage          *int    `json:"PackageUsage" name:"PackageUsage"`
		CenBandWidthPackageId *string `json:"CenBandWidthPackageId" name:"CenBandWidthPackageId"`
	} `json:"CenBandWidthPackageUsage"`
}

func (r *DescribeCenBandWidthPackageUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenBandWidthPackageUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInstancesRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInstanceId []*string                         `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	MaxResults        *int                              `json:"MaxResults,omitempty" name:"MaxResults"`
	Filter            []*DescribeNetworkInstancesFilter `json:"Filter,omitempty" name:"Filter"`
	NextToken         *string                           `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeNetworkInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNetworkInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NextToken          *string `json:"NextToken" name:"NextToken"`
	NetworkInstanceSet []struct {
		CenId             *string `json:"CenId" name:"CenId"`
		NetworkInstanceId *string `json:"NetworkInstanceId" name:"NetworkInstanceId"`
		InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		InstanceRegion    *string `json:"InstanceRegion" name:"InstanceRegion"`
		InstanceAccountId *string `json:"InstanceAccountId" name:"InstanceAccountId"`
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
	} `json:"NetworkInstanceSet"`
}

func (r *DescribeNetworkInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenGrantRequest struct {
	*ksyunhttp.BaseRequest
	CenId             *string `json:"CenId,omitempty" name:"CenId"`
	InstanceType      *string `json:"InstanceType,omitempty" name:"InstanceType"`
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	CenAccountId      *string `json:"CenAccountId,omitempty" name:"CenAccountId"`
}

func (r *CreateCenGrantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenGrantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCenGrantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCenGrantResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	CenGrant  struct {
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		CenGrantId        *string `json:"CenGrantId" name:"CenGrantId"`
		CenId             *string `json:"CenId" name:"CenId"`
		NetworkInstanceId *string `json:"NetworkInstanceId" name:"NetworkInstanceId"`
		InstanceType      *string `json:"InstanceType" name:"InstanceType"`
		CenAccountId      *string `json:"CenAccountId" name:"CenAccountId"`
	} `json:"CenGrant"`
}

func (r *CreateCenGrantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCenGrantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInterAreasRequest struct {
	*ksyunhttp.BaseRequest
	InterAreaId []*string                   `json:"InterAreaId,omitempty" name:"InterAreaId"`
	Filter      []*DescribeInterAreasFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults  *int                        `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken   *string                     `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeInterAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInterAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInterAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInterAreasResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	NextToken    *string `json:"NextToken" name:"NextToken"`
	InterAreaSet []struct {
		InterAreaId   *string `json:"InterAreaId" name:"InterAreaId"`
		InterAreaName *string `json:"InterAreaName" name:"InterAreaName"`
	} `json:"InterAreaSet"`
}

func (r *DescribeInterAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInterAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInterRegionsRequest struct {
	*ksyunhttp.BaseRequest
	InterRegionId []*string                     `json:"InterRegionId,omitempty" name:"InterRegionId"`
	Filter        []*DescribeInterRegionsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults    *int                          `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string                       `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeInterRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInterRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInterRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInterRegionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	NextToken      *string `json:"NextToken" name:"NextToken"`
	InterRegionSet []struct {
		InterAreaId     *string `json:"InterAreaId" name:"InterAreaId"`
		InterRegionId   *string `json:"InterRegionId" name:"InterRegionId"`
		InterRegionName *string `json:"InterRegionName" name:"InterRegionName"`
	} `json:"InterRegionSet"`
}

func (r *DescribeInterRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInterRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CenId             *string `json:"CenId,omitempty" name:"CenId"`
	InstanceType      *string `json:"InstanceType,omitempty" name:"InstanceType"`
	InstanceRegion    *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	InstanceAccountId *string `json:"InstanceAccountId,omitempty" name:"InstanceAccountId"`
}

func (r *AttachNetworkInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachNetworkInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CenId             *string `json:"CenId,omitempty" name:"CenId"`
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
}

func (r *DetachNetworkInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachNetworkInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CenCidrPublishRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInstanceId *string   `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	InstanceType      *string   `json:"InstanceType,omitempty" name:"InstanceType"`
	NetworkRouteId    []*string `json:"NetworkRouteId,omitempty" name:"NetworkRouteId"`
	SelfDefineCidr    []*string `json:"SelfDefineCidr,omitempty" name:"SelfDefineCidr"`
	CenId             *string   `json:"CenId,omitempty" name:"CenId"`
}

func (r *CenCidrPublishRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CenCidrPublishRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CenCidrPublishRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CenCidrPublishResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CenCidrPublishResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CenCidrPublishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CenCidrDeleteRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInstanceId *string   `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	InstanceType      *string   `json:"InstanceType,omitempty" name:"InstanceType"`
	NetworkRouteId    []*string `json:"NetworkRouteId,omitempty" name:"NetworkRouteId"`
	SelfRouteId       []*string `json:"SelfRouteId,omitempty" name:"SelfRouteId"`
	CenId             *string   `json:"CenId,omitempty" name:"CenId"`
}

func (r *CenCidrDeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CenCidrDeleteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CenCidrDeleteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CenCidrDeleteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CenCidrDeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CenCidrDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
