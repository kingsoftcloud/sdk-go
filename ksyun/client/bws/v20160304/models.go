package v20160304

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeBandWidthSharesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeBandWidthSharesTagKV struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateBandWidthShareRequest struct {
	*ksyunhttp.BaseRequest
	LineId             *string `json:"LineId,omitempty" name:"LineId"`
	BandWidth          *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	BandWidthShareName *string `json:"BandWidthShareName,omitempty" name:"BandWidthShareName"`
	ProjectId          *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType         *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreateBandWidthShareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateBandWidthShareResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BandWidthShareId   *string `json:"BandWidthShareId" name:"BandWidthShareId"`
	BandWidth          *int    `json:"BandWidth" name:"BandWidth"`
	BandWidthShareName *string `json:"BandWidthShareName" name:"BandWidthShareName"`
	CreateTime         *string `json:"CreateTime" name:"CreateTime"`
	LineId             *string `json:"LineId" name:"LineId"`
}

func (r *CreateBandWidthShareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBandWidthShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandWidthSharesRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId        []*string                        `json:"ProjectId,omitempty" name:"ProjectId"`
	BandWidthShareId []*string                        `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	Filter           []*DescribeBandWidthSharesFilter `json:"Filter,omitempty" name:"Filter"`
	IsContainTag     *bool                            `json:"IsContainTag,omitempty" name:"IsContainTag"`
	TagKey           []*string                        `json:"TagKey,omitempty" name:"TagKey"`
	TagKV            []*DescribeBandWidthSharesTagKV  `json:"TagKV,omitempty" name:"TagKV"`
	MaxResults       *int                             `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken        *string                          `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeBandWidthSharesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBandWidthSharesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId         *string `json:"RequestId" name:"RequestId"`
	NextToken         *string `json:"NextToken" name:"NextToken"`
	BandWidthShareSet []struct {
		BandWidthShareId               *string `json:"BandWidthShareId" name:"BandWidthShareId"`
		BandWidth                      *int    `json:"BandWidth" name:"BandWidth"`
		BandWidthShareName             *string `json:"BandWidthShareName" name:"BandWidthShareName"`
		CreateTime                     *string `json:"CreateTime" name:"CreateTime"`
		LineId                         *string `json:"LineId" name:"LineId"`
		ProjectId                      *string `json:"ProjectId" name:"ProjectId"`
		LineName                       *string `json:"LineName" name:"LineName"`
		AssociateBandWidthShareInfoSet []struct {
			AllocationId *string `json:"AllocationId" name:"AllocationId"`
		} `json:"AssociateBandWidthShareInfoSet" name:"AssociateBandWidthShareInfoSet"`
		ChargeType     *string `json:"ChargeType" name:"ChargeType"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		TagSet         []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"BandWidthShareSet"`
}

func (r *DescribeBandWidthSharesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandWidthSharesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateBandWidthShareRequest struct {
	*ksyunhttp.BaseRequest
	BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	AllocationId     *string `json:"AllocationId,omitempty" name:"AllocationId"`
}

func (r *AssociateBandWidthShareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AssociateBandWidthShareResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateBandWidthShareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateBandWidthShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateBandWidthShareRequest struct {
	*ksyunhttp.BaseRequest
	BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	AllocationId     *string `json:"AllocationId,omitempty" name:"AllocationId"`
	BandWidth        *int    `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *DisassociateBandWidthShareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisassociateBandWidthShareResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DisassociateBandWidthShareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateBandWidthShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBandWidthShareRequest struct {
	*ksyunhttp.BaseRequest
	BandWidthShareId   *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	BandWidth          *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	BandWidthShareName *string `json:"BandWidthShareName,omitempty" name:"BandWidthShareName"`
}

func (r *ModifyBandWidthShareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyBandWidthShareResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	BandWidthShare struct {
		BandWidthShareId               *string `json:"BandWidthShareId" name:"BandWidthShareId"`
		BandWidth                      *int    `json:"BandWidth" name:"BandWidth"`
		BandWidthShareName             *string `json:"BandWidthShareName" name:"BandWidthShareName"`
		CreateTime                     *string `json:"CreateTime" name:"CreateTime"`
		LineId                         *string `json:"LineId" name:"LineId"`
		ProjectId                      *string `json:"ProjectId" name:"ProjectId"`
		LineName                       *string `json:"LineName" name:"LineName"`
		AssociateBandWidthShareInfoSet []struct {
			AllocationId *string `json:"AllocationId" name:"AllocationId"`
		} `json:"AssociateBandWidthShareInfoSet" name:"AssociateBandWidthShareInfoSet"`
		ChargeType     *string `json:"ChargeType" name:"ChargeType"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	} `json:"BandWidthShare"`
}

func (r *ModifyBandWidthShareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBandWidthShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBandWidthShareRequest struct {
	*ksyunhttp.BaseRequest
	BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
}

func (r *DeleteBandWidthShareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBandWidthShareResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteBandWidthShareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBandWidthShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBwsTopEipMonitorRequest struct {
	*ksyunhttp.BaseRequest
	BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
	StartTime        *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime          *string `json:"EndTime,omitempty" name:"EndTime"`
	SortType         *string `json:"SortType,omitempty" name:"SortType"`
	PublicIp         *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

func (r *QueryBwsTopEipMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryBwsTopEipMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	BwsMonitorDataList []struct {
		AllocationId *string `json:"AllocationId" name:"AllocationId"`
		PublicIp     *string `json:"PublicIp" name:"PublicIp"`
		InBound      *string `json:"InBound" name:"InBound"`
		OutBound     *string `json:"OutBound" name:"OutBound"`
		Num          *string `json:"Num" name:"Num"`
	} `json:"BwsMonitorDataList"`
}

func (r *QueryBwsTopEipMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBwsTopEipMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
