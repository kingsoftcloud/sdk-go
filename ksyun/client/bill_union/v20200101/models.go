package v20200101
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DescribeBillSummaryByPayModeRequest struct {
	*ksyunhttp.BaseRequest
	BillBeginMonth *string `json:"BillBeginMonth,omitempty" name:"BillBeginMonth"`
	BillEndMonth   *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
	SubAccount     *int    `json:"SubAccount,omitempty" name:"SubAccount"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBillSummaryByPayModeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	Currency        *string `json:"Currency" name:"Currency"`
	RealTotalCost   *string `json:"RealTotalCost" name:"RealTotalCost"`
	SummaryOverview []struct {
		PayMode   *string `json:"PayMode" name:"PayMode"`
		RealTotalCost *string `json:"RealTotalCost" name:"RealTotalCost"`
		BillMonth *string `json:"BillMonth" name:"BillMonth"`
	} `json:"SummaryOverview"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeBillSummaryByProductRequest struct {
	*ksyunhttp.BaseRequest
	BillBeginMonth              *string `json:"BillBeginMonth,omitempty" name:"BillBeginMonth"`
	BillEndMonth                *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
	SubAccount                  *int    `json:"SubAccount,omitempty" name:"SubAccount"`
	FetchAllFinanceRelationData *bool   `json:"FetchAllFinanceRelationData,omitempty" name:"FetchAllFinanceRelationData"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBillSummaryByProductResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	Currency        *string `json:"Currency" name:"Currency"`
	RealTotalCost   *string `json:"RealTotalCost" name:"RealTotalCost"`
	SummaryOverview []struct {
		ProductCode *string `json:"ProductCode" name:"ProductCode"`
		ProductName *string `json:"ProductName" name:"ProductName"`
		RealTotalCost *string `json:"RealTotalCost" name:"RealTotalCost"`
		BillMonth   *string `json:"BillMonth" name:"BillMonth"`
	} `json:"SummaryOverview"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeBillSummaryByProjectRequest struct {
	*ksyunhttp.BaseRequest
	BillBeginMonth *string `json:"BillBeginMonth,omitempty" name:"BillBeginMonth"`
	BillEndMonth   *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
	SubAccount     *int    `json:"SubAccount,omitempty" name:"SubAccount"`
}

func (r *DescribeBillSummaryByProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeBillSummaryByProjectResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	Currency        *string `json:"Currency" name:"Currency"`
	RealTotalCost   *string `json:"RealTotalCost" name:"RealTotalCost"`
	SummaryOverview []struct {
		ProjectId   *string `json:"ProjectId" name:"ProjectId"`
		ProjectName *string `json:"ProjectName" name:"ProjectName"`
		RealTotalCost *string `json:"RealTotalCost" name:"RealTotalCost"`
		BillMonth   *string `json:"BillMonth" name:"BillMonth"`
	} `json:"SummaryOverview"`
}

func (r *DescribeBillSummaryByProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceSummaryBillsRequest struct {
	*ksyunhttp.BaseRequest
	BillMonth   *string `json:"BillMonth,omitempty" name:"BillMonth"`
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	Page        *int    `json:"Page,omitempty" name:"Page"`
	Size        *int    `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeInstanceSummaryBillsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceSummaryBillsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	Email           *string `json:"Email" name:"Email"`
	PageNum         *int    `json:"PageNum" name:"PageNum"`
	PageSize        *int    `json:"PageSize" name:"PageSize"`
	TotalCount      *int    `json:"TotalCount" name:"TotalCount"`
	CustomerId      *int    `json:"CustomerId" name:"CustomerId"`
	SummaryOverview []struct {
		BillsNo             *string `json:"BillsNo" name:"BillsNo"`
		CustomerBillMonth   *string `json:"CustomerBillMonth" name:"CustomerBillMonth"`
		BillMonth           *string `json:"BillMonth" name:"BillMonth"`
		ProductName         *string `json:"ProductName" name:"ProductName"`
		ProductSubTyeName   *string `json:"ProductSubTyeName" name:"ProductSubTyeName"`
		InstanceId          *string `json:"InstanceId" name:"InstanceId"`
		InstanceName        *string `json:"InstanceName" name:"InstanceName"`
		Currency            *string `json:"Currency" name:"Currency"`
		DetailBillStartTime *string `json:"DetailBillStartTime" name:"DetailBillStartTime"`
		DetailBillEndTime   *string `json:"DetailBillEndTime" name:"DetailBillEndTime"`
		ServiceBillStartTime *string `json:"ServiceBillStartTime" name:"ServiceBillStartTime"`
		PayMode             *string `json:"PayMode" name:"PayMode"`
		BillTypeName        *string `json:"BillTypeName" name:"BillTypeName"`
		RegionName          *string `json:"RegionName" name:"RegionName"`
		ZoneName            *string `json:"ZoneName" name:"ZoneName"`
		ProjectName         *string `json:"ProjectName" name:"ProjectName"`
		Duration            *string `json:"Duration" name:"Duration"`
		Remark              *string `json:"Remark" name:"Remark"`
		Cost                *string `json:"Cost" name:"Cost"`
		RealCost            *string `json:"RealCost" name:"RealCost"`
		ConfigSet           []struct {
			Key   *string `json:"Key" name:"Key"`
			Code  *string `json:"Code" name:"Code"`
			Value *string `json:"Value" name:"Value"`
		} `json:"ConfigSet" name:"ConfigSet"`
		ProviderSet []struct {
			Key   *string `json:"Key" name:"Key"`
			Code  *string `json:"Code" name:"Code"`
			Value *string `json:"Value" name:"Value"`
		} `json:"ProviderSet" name:"ProviderSet"`
		ConsumeResources []*string `json:"ConsumeResources" name:"ConsumeResources"`
		ExtraSet         []struct {
			Key   *string `json:"Key" name:"Key"`
			Code  *string `json:"Code" name:"Code"`
			Value *string `json:"Value" name:"Value"`
		} `json:"ExtraSet" name:"ExtraSet"`
		TagSet []struct {
			Key   *string `json:"Key" name:"Key"`
			Value *string `json:"Value" name:"Value"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"SummaryOverview"`
}

func (r *DescribeInstanceSummaryBillsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSummaryBillsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeProductCodeRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeProductCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeProductCodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	ProductGroupSet []struct {
		Key *string `json:"Key" name:"Key"`
		Value *string `json:"Value" name:"Value"`
	} `json:"ProductGroupSet"`
}

func (r *DescribeProductCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSplitItemBillDetailsRequest struct {
	*ksyunhttp.BaseRequest
	CustomerBillMonth *int    `json:"CustomerBillMonth,omitempty" name:"CustomerBillMonth"`
	ProductGroupCode  *string `json:"ProductGroupCode,omitempty" name:"ProductGroupCode"`
	StatisticType     *int    `json:"StatisticType,omitempty" name:"StatisticType"`
	PayType           *int    `json:"PayType,omitempty" name:"PayType"`
	SubAccountId      *int    `json:"SubAccountId,omitempty" name:"SubAccountId"`
	Page              *int    `json:"Page,omitempty" name:"Page"`
	Size              *string `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeSplitItemBillDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSplitItemBillDetailsResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"Status" name:"Status"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			InstanceId              *string   `json:"InstanceId" name:"InstanceId"`
			InstanceName            *string   `json:"InstanceName" name:"InstanceName"`
			ProjectId               *int      `json:"ProjectId" name:"ProjectId"`
			ProjectName             *string   `json:"ProjectName" name:"ProjectName"`
			ProductGroupName        *string   `json:"ProductGroupName" name:"ProductGroupName"`
			ProductTypeName         *string   `json:"ProductTypeName" name:"ProductTypeName"`
			BillItemName            *string   `json:"BillItemName" name:"BillItemName"`
			BillStartTime           *string   `json:"BillStartTime" name:"BillStartTime"`
			BillEndTime             *string   `json:"BillEndTime" name:"BillEndTime"`
			ServiceBeginTime        *string   `json:"ServiceBeginTime" name:"ServiceBeginTime"`
			RegionName              *string   `json:"RegionName" name:"RegionName"`
			AvailabilityZone        *string   `json:"AvailabilityZone" name:"AvailabilityZone"`
			PayTypeName             *string   `json:"PayTypeName" name:"PayTypeName"`
			BillTypeName            *string   `json:"BillTypeName" name:"BillTypeName"`
			BillDetailTypeName      *string   `json:"BillDetailTypeName" name:"BillDetailTypeName"`
			MeasureValue            *string   `json:"MeasureValue" name:"MeasureValue"`
			MeasureValueUnit        *string   `json:"MeasureValueUnit" name:"MeasureValueUnit"`
			SplitItemName           *string   `json:"SplitItemName" name:"SplitItemName"`
			SplitRatio              *string   `json:"SplitRatio" name:"SplitRatio"`
			Price                   *string   `json:"Price" name:"Price"`
			RealPrice               *string   `json:"RealPrice" name:"RealPrice"`
			Cash                    *string   `json:"Cash" name:"Cash"`
			Reward                  *string   `json:"Reward" name:"Reward"`
			CloudTicketDenomination *string   `json:"CloudTicketDenomination" name:"CloudTicketDenomination"`
			ResourceDeductValue     *string   `json:"ResourceDeductValue" name:"ResourceDeductValue"`
			Duration                *string   `json:"Duration" name:"Duration"`
			RuleRemark              *string   `json:"RuleRemark" name:"RuleRemark"`
			FinanceUnitName         *string   `json:"FinanceUnitName" name:"FinanceUnitName"`
			PriceFactorInfo         []*string `json:"PriceFactorInfo" name:"PriceFactorInfo"`
			ExtraInfo               []*string `json:"ExtraInfo" name:"ExtraInfo"`
			TagInfo                 []*string `json:"TagInfo" name:"TagInfo"`
			UserId                  *int      `json:"UserId" name:"UserId"`
			UserName                *string   `json:"UserName" name:"UserName"`
			BillMonth               *int      `json:"BillMonth" name:"BillMonth"`
			CustomerBillMonth       *int      `json:"CustomerBillMonth" name:"CustomerBillMonth"`
			CurrencyCode            *string   `json:"CurrencyCode" name:"CurrencyCode"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"Error" name:"Error"`
}

func (r *DescribeSplitItemBillDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSplitItemBillDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeMiItemBillsRequest struct {
	*ksyunhttp.BaseRequest
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
}

func (r *DescribeMiItemBillsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMiItemBillsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Url       *string `json:"url" name:"url"`
	Error     struct {
		Code *string `json:"Code" name:"Code"`
		Message *string `json:"Message" name:"Message"`
	} `json:"Error"`
}

func (r *DescribeMiItemBillsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiItemBillsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeSplitItemDayBillDetailsRequest struct {
	*ksyunhttp.BaseRequest
	CustomerBillMonth *int    `json:"CustomerBillMonth,omitempty" name:"CustomerBillMonth"`
	ProductGroupCode  *string `json:"ProductGroupCode,omitempty" name:"ProductGroupCode"`
	StatisticType     *int    `json:"StatisticType,omitempty" name:"StatisticType"`
	PayType           *int    `json:"PayType,omitempty" name:"PayType"`
	SubAccountId      *int    `json:"SubAccountId,omitempty" name:"SubAccountId"`
	Page              *int    `json:"Page,omitempty" name:"Page"`
	Size              *int    `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeSplitItemDayBillDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeSplitItemDayBillDetailsResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"Status" name:"Status"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeSplitItemDayBillDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSplitItemDayBillDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductGroupsRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ListProductGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListProductGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"Success" name:"Success"`
	Data      []struct {
		Id   *int    `json:"Id" name:"Id"`
		Code *string `json:"Code" name:"Code"`
		Name *string `json:"Name" name:"Name"`
		EnName *string `json:"EnName" name:"EnName"`
	} `json:"Data"`
}

func (r *ListProductGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

