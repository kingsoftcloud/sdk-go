package v20160304

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type ModifyCfwAclSrcZonesSubAreas struct {
	AreaCode *string `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}
type ModifyCfwAclSrcZones struct {
	AreaCode *string                       `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string                       `json:"AreaName,omitempty" name:"AreaName"`
	SubAreas *ModifyCfwAclSrcZonesSubAreas `json:"SubAreas,omitempty" name:"SubAreas"`
}
type ModifyCfwAclDestZonesSubAreas struct {
	AreaCode *string `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}
type ModifyCfwAclDestZones struct {
	AreaCode *string                        `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string                        `json:"AreaName,omitempty" name:"AreaName"`
	SubAreas *ModifyCfwAclDestZonesSubAreas `json:"SubAreas,omitempty" name:"SubAreas"`
}
type CreateCfwAclSrcZonesSubAreas struct {
	AreaCode *string `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}
type CreateCfwAclSrcZones struct {
	AreaCode *string                       `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string                       `json:"AreaName,omitempty" name:"AreaName"`
	SubAreas *CreateCfwAclSrcZonesSubAreas `json:"SubAreas,omitempty" name:"SubAreas"`
}
type CreateCfwAclDestZonesSubAreas struct {
	AreaCode *string `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}
type CreateCfwAclDestZones struct {
	AreaCode *string                        `json:"AreaCode,omitempty" name:"AreaCode"`
	AreaName *string                        `json:"AreaName,omitempty" name:"AreaName"`
	SubAreas *CreateCfwAclDestZonesSubAreas `json:"SubAreas,omitempty" name:"SubAreas"`
}

type DeleteBatchCfwAddrbookRequest struct {
	*ksyunhttp.BaseRequest
	AddrbookIds   []*string `json:"AddrbookIds,omitempty" name:"AddrbookIds"`
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *DeleteBatchCfwAddrbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBatchCfwAddrbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteBatchCfwAddrbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchCfwAddrbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceGroupBatchRequest struct {
	*ksyunhttp.BaseRequest
	ServiceGroupIds []*string `json:"ServiceGroupIds,omitempty" name:"ServiceGroupIds"`
	CfwInstanceId   *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *DeleteServiceGroupBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteServiceGroupBatchResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteServiceGroupBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceGroupBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchHostbookRequest struct {
	*ksyunhttp.BaseRequest
	HostbookIds   []*string `json:"HostbookIds,omitempty" name:"HostbookIds"`
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *DeleteBatchHostbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteBatchHostbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteBatchHostbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchHostbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostbookRequest struct {
	*ksyunhttp.BaseRequest
	HostbookId   *string   `json:"HostbookId,omitempty" name:"HostbookId"`
	HostbookName *string   `json:"HostbookName,omitempty" name:"HostbookName"`
	HostValue    []*string `json:"HostValue,omitempty" name:"HostValue"`
	Description  *string   `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyHostbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyHostbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
	Id        *string `json:"Id" name:"Id"`
}

func (r *ModifyHostbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostbookRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	HostbookName  *string   `json:"HostbookName,omitempty" name:"HostbookName"`
	HostValue     []*string `json:"HostValue,omitempty" name:"HostValue"`
	Description   *string   `json:"Description,omitempty" name:"Description"`
}

func (r *CreateHostbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateHostbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	CfwAddrbook struct {
		HostbookId    *string   `json:"HostbookId" name:"HostbookId"`
		HostbookName  *string   `json:"HostbookName" name:"HostbookName"`
		HostValue     []*string `json:"HostValue" name:"HostValue"`
		Description   *string   `json:"Description" name:"Description"`
		CitationCount *int      `json:"CitationCount" name:"CitationCount"`
		CreateTime    *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAddrbook"`
}

func (r *CreateHostbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostbookRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	HostbookIds   []*string `json:"HostbookIds,omitempty" name:"HostbookIds"`
	MaxResults    *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeHostbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeHostbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	CfwHostbooks []struct {
		HostbookId    *string   `json:"HostbookId" name:"HostbookId"`
		HostbookName  *string   `json:"HostbookName" name:"HostbookName"`
		HostValue     []*string `json:"HostValue" name:"HostValue"`
		Description   *string   `json:"Description" name:"Description"`
		CitationCount *int      `json:"CitationCount" name:"CitationCount"`
		CreateTime    *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwHostbooks"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeHostbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCfwInstanceDetailRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *QueryCfwInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryCfwInstanceDetailResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *QueryCfwInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCfwInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudFireWallInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *DeleteCloudFireWallInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCloudFireWallInstanceResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DeleteCloudFireWallInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudFireWallInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOverviewDetailRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	StartTime     *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime       *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *QueryOverviewDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryOverviewDetailResponse struct {
	*ksyunhttp.BaseResponse
	CloudFireWallInstanceDetail struct {
		InstanceId     *string `json:"InstanceId" name:"InstanceId"`
		InstanceName   *string `json:"InstanceName" name:"InstanceName"`
		InstanceType   *string `json:"InstanceType" name:"InstanceType"`
		Bandwidth      *int    `json:"Bandwidth" name:"Bandwidth"`
		Status         *int    `json:"Status" name:"Status"`
		TotalEipNum    *int    `json:"TotalEipNum" name:"TotalEipNum"`
		TotalAclNum    *int    `json:"TotalAclNum" name:"TotalAclNum"`
		UsedEipNum     *int    `json:"UsedEipNum" name:"UsedEipNum"`
		UsedAclNum     *int    `json:"UsedAclNum" name:"UsedAclNum"`
		IpsStatus      *int    `json:"IpsStatus" name:"IpsStatus"`
		AvStatus       *int    `json:"AvStatus" name:"AvStatus"`
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		ChargeType     *string `json:"ChargeType" name:"ChargeType"`
		ProjectName    *string `json:"ProjectName" name:"ProjectName"`
		Region         *string `json:"Region" name:"Region"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		Duration       *int    `json:"Duration" name:"Duration"`
		FwLbId         *string `json:"FwLbId" name:"FwLbId"`
		AclDenyCount   *int    `json:"AclDenyCount" name:"AclDenyCount"`
		IpsCount       *int    `json:"IpsCount" name:"IpsCount"`
		InMax          *string `json:"InMax" name:"InMax"`
		OutMax         *string `json:"OutMax" name:"OutMax"`
	} `json:"CloudFireWallInstanceDetail"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *QueryOverviewDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOverviewDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficLogRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	StartTime     *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime       *string `json:"EndTime,omitempty" name:"EndTime"`
	QueryKeyword  *string `json:"QueryKeyword,omitempty" name:"QueryKeyword"`
	MaxResults    *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeTrafficLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeTrafficLogResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeTrafficLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	StartTime     *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime       *string `json:"EndTime,omitempty" name:"EndTime"`
	QueryKeyword  *string `json:"QueryKeyword,omitempty" name:"QueryKeyword"`
	MaxResults    *string `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAttackLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAttackLogResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeAttackLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttackLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclLogRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	StartTime     *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime       *string `json:"EndTime,omitempty" name:"EndTime"`
	QueryKeyword  *string `json:"QueryKeyword,omitempty" name:"QueryKeyword"`
	MaxResults    *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeAclLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAclLogResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeAclLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudFireWallFeatureRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	InstanceName  *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyCloudFireWallFeatureRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCloudFireWallFeatureResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *ModifyCloudFireWallFeatureResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudFireWallFeatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwAddrbookRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	AddrbookIds   []*string `json:"AddrbookIds,omitempty" name:"AddrbookIds"`
	MaxResults    *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCfwAddrbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCfwAddrbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	CfwAddrbooks []struct {
		AddrbookId    *string   `json:"AddrbookId" name:"AddrbookId"`
		AddrbookName  *string   `json:"AddrbookName" name:"AddrbookName"`
		IpAddress     []*string `json:"IpAddress" name:"IpAddress"`
		Description   *string   `json:"Description" name:"Description"`
		CitationCount *int      `json:"CitationCount" name:"CitationCount"`
		CreateTime    *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAddrbooks"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeCfwAddrbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwAddrbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCfwAddrbookRequest struct {
	*ksyunhttp.BaseRequest
	AddrbookId *string `json:"AddrbookId,omitempty" name:"AddrbookId"`
}

func (r *DeleteCfwAddrbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCfwAddrbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteCfwAddrbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfwAddrbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCfwAddrbookRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	AddrbookName  *string   `json:"AddrbookName,omitempty" name:"AddrbookName"`
	IpAddress     []*string `json:"IpAddress,omitempty" name:"IpAddress"`
	Description   *string   `json:"Description,omitempty" name:"Description"`
	IpVersion     *string   `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *ModifyCfwAddrbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCfwAddrbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ModifyCfwAddrbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCfwAddrbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfwAddrbookRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	AddrbookName  *string   `json:"AddrbookName,omitempty" name:"AddrbookName"`
	IpAddress     []*string `json:"IpAddress,omitempty" name:"IpAddress"`
	Description   *string   `json:"Description,omitempty" name:"Description"`
	IpVersion     *string   `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *CreateCfwAddrbookRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateCfwAddrbookResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	CfwAddrbook struct {
		AddrbookId    *string   `json:"AddrbookId" name:"AddrbookId"`
		AddrbookName  *string   `json:"AddrbookName" name:"AddrbookName"`
		IpAddress     []*string `json:"IpAddress" name:"IpAddress"`
		Description   *string   `json:"Description" name:"Description"`
		CitationCount *int      `json:"CitationCount" name:"CitationCount"`
		CreateTime    *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAddrbook"`
}

func (r *CreateCfwAddrbookResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfwAddrbookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlterCfwAclStatusRequest struct {
	*ksyunhttp.BaseRequest
	AclIds []*string `json:"AclIds,omitempty" name:"AclIds"`
	Status *string   `json:"Status,omitempty" name:"Status"`
}

func (r *AlterCfwAclStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AlterCfwAclStatusResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *AlterCfwAclStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlterCfwAclStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetCfwAclHitCountRequest struct {
	*ksyunhttp.BaseRequest
	AclIds []*string `json:"AclIds,omitempty" name:"AclIds"`
}

func (r *ResetCfwAclHitCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResetCfwAclHitCountResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ResetCfwAclHitCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetCfwAclHitCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlterAclPriorityRequest struct {
	*ksyunhttp.BaseRequest
	PriorityPosition *string `json:"PriorityPosition,omitempty" name:"PriorityPosition"`
	AclId            *string `json:"AclId,omitempty" name:"AclId"`
	CfwInstanceId    *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *AlterAclPriorityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AlterAclPriorityResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *AlterAclPriorityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlterAclPriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCfwAclRequest struct {
	*ksyunhttp.BaseRequest
	AclIds        []*string `json:"AclIds,omitempty" name:"AclIds"`
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
}

func (r *DeleteCfwAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCfwAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteCfwAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfwAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCfwAclRequest struct {
	*ksyunhttp.BaseRequest
	AclId         *string                  `json:"AclId,omitempty" name:"AclId"`
	CfwInstanceId *string                  `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	AclName       *string                  `json:"AclName,omitempty" name:"AclName"`
	Direction     *string                  `json:"Direction,omitempty" name:"Direction"`
	SrcType       *string                  `json:"SrcType,omitempty" name:"SrcType"`
	SrcIps        []*string                `json:"SrcIps,omitempty" name:"SrcIps"`
	SrcAddrbooks  []*string                `json:"SrcAddrbooks,omitempty" name:"SrcAddrbooks"`
	SrcZones      []*ModifyCfwAclSrcZones  `json:"SrcZones,omitempty" name:"SrcZones"`
	DestType      *string                  `json:"DestType,omitempty" name:"DestType"`
	DestIps       []*string                `json:"DestIps,omitempty" name:"DestIps"`
	DestAddrbooks []*string                `json:"DestAddrbooks,omitempty" name:"DestAddrbooks"`
	DestZones     []*ModifyCfwAclDestZones `json:"DestZones,omitempty" name:"DestZones"`
	DestHost      []*string                `json:"DestHost,omitempty" name:"DestHost"`
	DestHostbook  []*string                `json:"DestHostbook,omitempty" name:"DestHostbook"`
	ServiceType   *string                  `json:"ServiceType,omitempty" name:"ServiceType"`
	ServiceInfos  []*string                `json:"ServiceInfos,omitempty" name:"ServiceInfos"`
	ServiceGroups []*string                `json:"ServiceGroups,omitempty" name:"ServiceGroups"`
	AppType       *string                  `json:"AppType,omitempty" name:"AppType"`
	AppValue      []*string                `json:"AppValue,omitempty" name:"AppValue"`
	Policy        *string                  `json:"Policy,omitempty" name:"Policy"`
	Status        *string                  `json:"Status,omitempty" name:"Status"`
	Description   *string                  `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyCfwAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCfwAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ModifyCfwAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCfwAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwAclRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	AclIds        []*string `json:"AclIds,omitempty" name:"AclIds"`
	MaxResults    *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCfwAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCfwAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	CfwAcls   []struct {
		CfwInstanceId *string `json:"CfwInstanceId" name:"CfwInstanceId"`
		AclId         *string `json:"AclId" name:"AclId"`
		AclName       *string `json:"AclName" name:"AclName"`
		Direction     *string `json:"Direction" name:"Direction"`
		SrcType       *string `json:"SrcType" name:"SrcType"`
		SrcIps        *string `json:"SrcIps" name:"SrcIps"`
		SrcAddrbooks  *string `json:"SrcAddrbooks" name:"SrcAddrbooks"`
		SrcZones      []struct {
			AreaCode *string `json:"AreaCode" name:"AreaCode"`
			AreaName *string `json:"AreaName" name:"AreaName"`
			SubAreas *string `json:"SubAreas" name:"SubAreas"`
		} `json:"SrcZones" name:"SrcZones"`
		DestType      *string   `json:"DestType" name:"DestType"`
		DestIps       *string   `json:"DestIps" name:"DestIps"`
		DestAddrbooks []*string `json:"DestAddrbooks" name:"DestAddrbooks"`
		ServiceType   *string   `json:"ServiceType" name:"ServiceType"`
		ServiceInfo   *string   `json:"ServiceInfo" name:"ServiceInfo"`
		ServiceGroups []*string `json:"ServiceGroups" name:"ServiceGroups"`
		AppType       *string   `json:"AppType" name:"AppType"`
		AppValue      []*string `json:"AppValue" name:"AppValue"`
		Policy        *string   `json:"Policy" name:"Policy"`
		Status        *string   `json:"Status" name:"Status"`
		Priority      *int      `json:"Priority" name:"Priority"`
		Description   *string   `json:"Description" name:"Description"`
		HitCount      *int      `json:"HitCount" name:"HitCount"`
		CreateTime    *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAcls"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeCfwAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfwAclRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId    *string                  `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	AclName          *string                  `json:"AclName,omitempty" name:"AclName"`
	Direction        *string                  `json:"Direction,omitempty" name:"Direction"`
	SrcType          *string                  `json:"SrcType,omitempty" name:"SrcType"`
	SrcIps           []*string                `json:"SrcIps,omitempty" name:"SrcIps"`
	SrcAddrbooks     []*string                `json:"SrcAddrbooks,omitempty" name:"SrcAddrbooks"`
	SrcZones         []*CreateCfwAclSrcZones  `json:"SrcZones,omitempty" name:"SrcZones"`
	DestType         *string                  `json:"DestType,omitempty" name:"DestType"`
	DestIps          []*string                `json:"DestIps,omitempty" name:"DestIps"`
	DestAddrbooks    []*string                `json:"DestAddrbooks,omitempty" name:"DestAddrbooks"`
	DestZones        []*CreateCfwAclDestZones `json:"DestZones,omitempty" name:"DestZones"`
	DestHost         []*string                `json:"DestHost,omitempty" name:"DestHost"`
	DestHostbook     []*string                `json:"DestHostbook,omitempty" name:"DestHostbook"`
	ServiceType      *string                  `json:"ServiceType,omitempty" name:"ServiceType"`
	ServiceInfos     []*string                `json:"ServiceInfos,omitempty" name:"ServiceInfos"`
	ServiceGroups    []*string                `json:"ServiceGroups,omitempty" name:"ServiceGroups"`
	AppType          *string                  `json:"AppType,omitempty" name:"AppType"`
	AppValue         []*string                `json:"AppValue,omitempty" name:"AppValue"`
	Policy           *string                  `json:"Policy,omitempty" name:"Policy"`
	Status           *string                  `json:"Status,omitempty" name:"Status"`
	Description      *string                  `json:"Description,omitempty" name:"Description"`
	PriorityPosition *string                  `json:"PriorityPosition,omitempty" name:"PriorityPosition"`
}

func (r *CreateCfwAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateCfwAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	CfwAcl    struct {
		CfwInstanceId *string `json:"CfwInstanceId" name:"CfwInstanceId"`
		AclId         *string `json:"AclId" name:"AclId"`
		AclName       *string `json:"AclName" name:"AclName"`
		Direction     *string `json:"Direction" name:"Direction"`
		SrcType       *string `json:"SrcType" name:"SrcType"`
		SrcIps        *string `json:"SrcIps" name:"SrcIps"`
		SrcAddrbooks  *string `json:"SrcAddrbooks" name:"SrcAddrbooks"`
		SrcZones      []struct {
			AreaCode *string `json:"AreaCode" name:"AreaCode"`
			AreaName *string `json:"AreaName" name:"AreaName"`
			SubAreas *string `json:"SubAreas" name:"SubAreas"`
		} `json:"SrcZones" name:"SrcZones"`
		DestType      *string   `json:"DestType" name:"DestType"`
		DestIps       *string   `json:"DestIps" name:"DestIps"`
		DestAddrbooks []*string `json:"DestAddrbooks" name:"DestAddrbooks"`
		ServiceType   *string   `json:"ServiceType" name:"ServiceType"`
		ServiceInfo   *string   `json:"ServiceInfo" name:"ServiceInfo"`
		ServiceGroups []*string `json:"ServiceGroups" name:"ServiceGroups"`
		AppType       *string   `json:"AppType" name:"AppType"`
		AppValue      []*string `json:"AppValue" name:"AppValue"`
		Policy        *string   `json:"Policy" name:"Policy"`
		Status        *string   `json:"Status" name:"Status"`
		Priority      *int      `json:"Priority" name:"Priority"`
		Description   *string   `json:"Description" name:"Description"`
		HitCount      *int      `json:"HitCount" name:"HitCount"`
		CreateTime    *string   `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAcl"`
}

func (r *CreateCfwAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfwAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCfwEipProtectRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId    *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	EipIds           []*string `json:"EipIds,omitempty" name:"EipIds"`
	EipProtectStatus *string   `json:"EipProtectStatus,omitempty" name:"EipProtectStatus"`
}

func (r *ModifyCfwEipProtectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCfwEipProtectResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ModifyCfwEipProtectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCfwEipProtectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwEipsRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId *string `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	MaxResults    *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken     *string `json:"NextToken,omitempty" name:"NextToken"`
	EipAddress    *string `json:"EipAddress,omitempty" name:"EipAddress"`
}

func (r *DescribeCfwEipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCfwEipsResponse struct {
	*ksyunhttp.BaseResponse
	PreToken     *string `json:"pre_token" name:"pre_token"`
	RequestId    *string `json:"RequestId" name:"RequestId"`
	AddressesSet []struct {
		Type                 *string `json:"Type" name:"Type"`
		Ingress              *int    `json:"Ingress" name:"Ingress"`
		RouterUuid           *string `json:"RouterUuid" name:"RouterUuid"`
		LbPoolUuid           *string `json:"LbPoolUuid" name:"LbPoolUuid"`
		NatpoolId            *string `json:"NatpoolId" name:"NatpoolId"`
		FixedIpAddress       *string `json:"FixedIpAddress" name:"FixedIpAddress"`
		PortUuid             *string `json:"PortUuid" name:"PortUuid"`
		PublicIp             *string `json:"PublicIp" name:"PublicIp"`
		AllocationId         *string `json:"AllocationId" name:"AllocationId"`
		State                *string `json:"State" name:"State"`
		IpState              *string `json:"IpState" name:"IpState"`
		LineId               *string `json:"LineId" name:"LineId"`
		NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
		PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		BandWidth            *int    `json:"BandWidth" name:"BandWidth"`
		InstanceType         *string `json:"InstanceType" name:"InstanceType"`
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		InternetGatewayId    *string `json:"InternetGatewayId" name:"InternetGatewayId"`
		BandWidthShareId     *string `json:"BandWidthShareId" name:"BandWidthShareId"`
		UserTag              *string `json:"UserTag" name:"UserTag"`
		ChargeType           *string `json:"ChargeType" name:"ChargeType"`
		ProductType          *string `json:"ProductType" name:"ProductType"`
		ProductWhat          *string `json:"ProductWhat" name:"ProductWhat"`
		ServiceEndTime       *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		IpVersion            *string `json:"IpVersion" name:"IpVersion"`
		ProjectId            *string `json:"ProjectId" name:"ProjectId"`
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		Mode                 *string `json:"Mode" name:"Mode"`
		HostType             *string `json:"HostType" name:"HostType"`
		EipPoolId            *string `json:"EipPoolId" name:"EipPoolId"`
		FirewallId           *string `json:"FirewallId" name:"FirewallId"`
	} `json:"AddressesSet"`
	NextToken  *string `json:"NextToken" name:"NextToken"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
}

func (r *DescribeCfwEipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceGroupRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId   *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	ServiceGroupIds []*string `json:"ServiceGroupIds,omitempty" name:"ServiceGroupIds"`
}

func (r *DescribeServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeServiceGroupResponse struct {
	*ksyunhttp.BaseResponse
	CfwServiceGroups []struct {
		ServiceGroupId   *string   `json:"ServiceGroupId" name:"ServiceGroupId"`
		ServiceGroupName *string   `json:"ServiceGroupName" name:"ServiceGroupName"`
		ServiceInfo      []*string `json:"ServiceInfo" name:"ServiceInfo"`
		Description      *string   `json:"Description" name:"Description"`
		CitationCount    *int      `json:"CitationCount" name:"CitationCount"`
	} `json:"CfwServiceGroups"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCfwServiceGroupRequest struct {
	*ksyunhttp.BaseRequest
	ServiceGroupId   *string   `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
	ServiceGroupName *string   `json:"ServiceGroupName,omitempty" name:"ServiceGroupName"`
	ServiceInfo      []*string `json:"ServiceInfo,omitempty" name:"ServiceInfo"`
	Description      *string   `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyCfwServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCfwServiceGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ModifyCfwServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCfwServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCfwServiceGroupRequest struct {
	*ksyunhttp.BaseRequest
	ServiceGroupId *string `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
}

func (r *DeleteCfwServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCfwServiceGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteCfwServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfwServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfwServiceGroupRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceId    *string   `json:"CfwInstanceId,omitempty" name:"CfwInstanceId"`
	ServiceGroupName *string   `json:"ServiceGroupName,omitempty" name:"ServiceGroupName"`
	ServiceInfo      []*string `json:"ServiceInfo,omitempty" name:"ServiceInfo"`
	Description      *string   `json:"Description,omitempty" name:"Description"`
}

func (r *CreateCfwServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateCfwServiceGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	CfwServiceGroup struct {
		ServiceGroupId   *string   `json:"ServiceGroupId" name:"ServiceGroupId"`
		ServiceGroupName *string   `json:"ServiceGroupName" name:"ServiceGroupName"`
		ServiceInfo      []*string `json:"ServiceInfo" name:"ServiceInfo"`
		Description      *string   `json:"Description" name:"Description"`
		CitationCount    *int      `json:"CitationCount" name:"CitationCount"`
	} `json:"CfwServiceGroup"`
}

func (r *CreateCfwServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfwServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudFireWallInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceIds []*string `json:"CfwInstanceIds,omitempty" name:"CfwInstanceIds"`
}

func (r *DescribeCloudFireWallInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCloudFireWallInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"requestId" name:"requestId"`
	CloudFireWallInstanceList []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		Bandwidth    *int    `json:"Bandwidth" name:"Bandwidth"`
		Status       *int    `json:"Status" name:"Status"`
		TotalEipNum  *int    `json:"TotalEipNum" name:"TotalEipNum"`
		UsedEipNum   *int    `json:"UsedEipNum" name:"UsedEipNum"`
		IpsStatus    *int    `json:"IpsStatus" name:"IpsStatus"`
		AvStatus     *int    `json:"AvStatus" name:"AvStatus"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
	} `json:"CloudFireWallInstanceList"`
}

func (r *DescribeCloudFireWallInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudFireWallInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
