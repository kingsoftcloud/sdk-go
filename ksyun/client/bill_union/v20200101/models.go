package v20200101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeBillSummaryByPayModeRequest struct {
    *ksyunhttp.BaseRequest
    BillBeginMonth *string `json:"BillBeginMonth,omitempty" name:"BillBeginMonth"`
    BillEndMonth *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
    SubAccount *int `json:"SubAccount,omitempty" name:"SubAccount"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByPayModeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Currency *string `json:"Currency" name:"Currency"`
    RealTotalCost *string `json:"RealTotalCost" name:"RealTotalCost"`
	SummaryOverview []struct {
		PayMode *string `json:"PayMode"`
		RealTotalCost *string `json:"RealTotalCost"`
		BillMonth *string `json:"BillMonth"`
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
    BillBeginMonth *string `json:"BillBeginMonth,omitempty" name:"BillBeginMonth"`
    BillEndMonth *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
    SubAccount *int `json:"SubAccount,omitempty" name:"SubAccount"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProductRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Currency *string `json:"Currency" name:"Currency"`
    RealTotalCost *string `json:"RealTotalCost" name:"RealTotalCost"`
	SummaryOverview []struct {
		ProductCode *string `json:"ProductCode"`
		ProductName *string `json:"ProductName"`
		RealTotalCost *string `json:"RealTotalCost"`
		BillMonth *string `json:"BillMonth"`
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
    BillEndMonth *string `json:"BillEndMonth,omitempty" name:"BillEndMonth"`
    SubAccount *int `json:"SubAccount,omitempty" name:"SubAccount"`
}

func (r *DescribeBillSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProjectRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProjectRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProjectResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Currency *string `json:"Currency" name:"Currency"`
    RealTotalCost *string `json:"RealTotalCost" name:"RealTotalCost"`
	SummaryOverview []struct {
		ProjectId *string `json:"ProjectId"`
		ProjectName *string `json:"ProjectName"`
		RealTotalCost *string `json:"RealTotalCost"`
		BillMonth *string `json:"BillMonth"`
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
    BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
    ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeInstanceSummaryBillsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceSummaryBillsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInstanceSummaryBillsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSummaryBillsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Email *string `json:"Email" name:"Email"`
    PageNum *int `json:"PageNum" name:"PageNum"`
    PageSize *int `json:"PageSize" name:"PageSize"`
    TotalCount *int `json:"TotalCount" name:"TotalCount"`
    CustomerId *int `json:"CustomerId" name:"CustomerId"`
	SummaryOverview []struct {
		BillsNo *string `json:"BillsNo"`
		CustomerBillMonth *string `json:"CustomerBillMonth"`
		BillMonth *string `json:"BillMonth"`
		ProductName *string `json:"ProductName"`
		ProductSubTyeName *string `json:"ProductSubTyeName"`
		InstanceId *string `json:"InstanceId"`
		InstanceName *string `json:"InstanceName"`
		Currency *string `json:"Currency"`
		DetailBillStartTime *string `json:"DetailBillStartTime"`
		DetailBillEndTime *string `json:"DetailBillEndTime"`
		ServiceBillStartTime *string `json:"ServiceBillStartTime"`
		PayMode *string `json:"PayMode"`
		BillTypeName *string `json:"BillTypeName"`
		RegionName *string `json:"RegionName"`
		ZoneName *string `json:"ZoneName"`
		ProjectName *string `json:"ProjectName"`
		Duration *string `json:"Duration"`
		Remark *string `json:"Remark"`
		Cost *string `json:"Cost"`
		RealCost *string `json:"RealCost"`
		ConfigSet []struct {
					Key *string `json:"Key"`
					Code *string `json:"Code"`
					Value *string `json:"Value"`
			} `json:"ConfigSet"`
			ProviderSet []struct {
						Key *string `json:"Key"`
						Code *string `json:"Code"`
						Value *string `json:"Value"`
				} `json:"ProviderSet"`
				ConsumeResources []struct {
					} `json:"ConsumeResources"`
					ExtraSet []struct {
								Key *string `json:"Key"`
								Code *string `json:"Code"`
								Value *string `json:"Value"`
						} `json:"ExtraSet"`
						TagSet []struct {
									Key *string `json:"Key"`
									Value *string `json:"Value"`
							} `json:"TagSet"`
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

func (r *DescribeProductCodeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeProductCodeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductCodeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	ProductGroupSet []struct {
		Key *string `json:"Key"`
		Value *string `json:"Value"`
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
    CustomerBillMonth *int `json:"CustomerBillMonth,omitempty" name:"CustomerBillMonth"`
    ProductGroupCode *string `json:"ProductGroupCode,omitempty" name:"ProductGroupCode"`
    StatisticType *int `json:"StatisticType,omitempty" name:"StatisticType"`
    PayType *int `json:"PayType,omitempty" name:"PayType"`
    SubAccountId *int `json:"SubAccountId,omitempty" name:"SubAccountId"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *string `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeSplitItemBillDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSplitItemBillDetailsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSplitItemBillDetailsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSplitItemBillDetailsResponse struct {
    *ksyunhttp.BaseResponse
    Status *int `json:"Status" name:"Status"`
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					InstanceId *string `json:"InstanceId"`
					InstanceName *string `json:"InstanceName"`
					ProjectId *int `json:"ProjectId"`
					ProjectName *string `json:"ProjectName"`
					ProductGroupName *string `json:"ProductGroupName"`
					ProductTypeName *string `json:"ProductTypeName"`
					BillItemName *string `json:"BillItemName"`
					BillStartTime *string `json:"BillStartTime"`
					BillEndTime *string `json:"BillEndTime"`
					ServiceBeginTime *string `json:"ServiceBeginTime"`
					RegionName *string `json:"RegionName"`
					AvailabilityZone *string `json:"AvailabilityZone"`
					PayTypeName *string `json:"PayTypeName"`
					BillTypeName *string `json:"BillTypeName"`
					BillDetailTypeName *string `json:"BillDetailTypeName"`
					MeasureValue *string `json:"MeasureValue"`
					MeasureValueUnit *string `json:"MeasureValueUnit"`
					SplitItemName *string `json:"SplitItemName"`
					SplitRatio *string `json:"SplitRatio"`
					Price *string `json:"Price"`
					RealPrice *string `json:"RealPrice"`
					Cash *string `json:"Cash"`
					Reward *string `json:"Reward"`
					CloudTicketDenomination *string `json:"CloudTicketDenomination"`
					ResourceDeductValue *string `json:"ResourceDeductValue"`
					Duration *string `json:"Duration"`
					RuleRemark *string `json:"RuleRemark"`
					FinanceUnitName *string `json:"FinanceUnitName"`
				PriceFactorInfo []struct {
				} `json:"PriceFactorInfo"`
				ExtraInfo []struct {
				} `json:"ExtraInfo"`
				TagInfo []struct {
				} `json:"TagInfo"`
					UserId *int `json:"UserId"`
					UserName *string `json:"UserName"`
					BillMonth *int `json:"BillMonth"`
					CustomerBillMonth *int `json:"CustomerBillMonth"`
					CurrencyCode *string `json:"CurrencyCode"`
			} `json:"Bills"`
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

func (r *DescribeMiItemBillsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeMiItemBillsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMiItemBillsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    url *string `json:"url" name:"url"`
	Error struct {
		Code *string `json:"Code"`
		Message *string `json:"Message"`
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
    CustomerBillMonth *int `json:"CustomerBillMonth,omitempty" name:"CustomerBillMonth"`
    ProductGroupCode *string `json:"ProductGroupCode,omitempty" name:"ProductGroupCode"`
    StatisticType *int `json:"StatisticType,omitempty" name:"StatisticType"`
    PayType *int `json:"PayType,omitempty" name:"PayType"`
    SubAccountId *int `json:"SubAccountId,omitempty" name:"SubAccountId"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
}

func (r *DescribeSplitItemDayBillDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSplitItemDayBillDetailsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSplitItemDayBillDetailsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSplitItemDayBillDetailsResponse struct {
    *ksyunhttp.BaseResponse
    Status *int `json:"Status" name:"Status"`
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

func (r *ListProductGroupsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListProductGroupsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListProductGroupsResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Success *bool `json:"Success" name:"Success"`
	Data []struct {
		Id *int `json:"Id"`
		Code *string `json:"Code"`
		Name *string `json:"Name"`
		EnName *string `json:"EnName"`
	} `json:"Data"`
}

func (r *ListProductGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListProductGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

